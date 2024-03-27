package processor

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/simonovic86/superform/config"
	"github.com/simonovic86/superform/internal/model"
	"github.com/simonovic86/superform/pkg/api"
	"github.com/simonovic86/superform/pkg/blockchain"
	"github.com/simonovic86/superform/pkg/blockchain/superform"
)

// ProcessTransactions filters transactions based on source and destination chain IDs, AMBs, and confirms withdrawals
func ProcessTransactions(transactions []model.Transaction, txType int, sourceChainID, destinationChainID int, ambs []string, chains []model.Chain) ([]model.ProcessedTransaction, error) {
	var filteredTxs []model.ProcessedTransaction

	for _, tx := range transactions {
		// Ensure the transaction is a withdrawal and matches the source and destination chain IDs
		if tx.TxType == txType && tx.FromChain == sourceChainID && tx.ToChain == destinationChainID {
			fmt.Printf("processing transaction: %s\n", tx.TransactionHash)

			detailedTx, err := api.FetchDetailedTransaction(tx.TransactionHash)

			if err != nil {
				return nil, err
			}

			if includesAMB(detailedTx, ambs) {
				// fetch confirmation timestamp and original output amounts
				var rpc string
				var timestamp int64
				var dstPayloadId int
				var blockNumber int64

				timestamp, rpc, dstPayloadId, blockNumber, err = fetchDstData(detailedTx, chains)

				if err != nil {
					return nil, err
				}

				if timestamp != -1 {
					// the transaction is processed
					detailedTx.CreationTimestamp = tx.CreationTimestamp
					detailedTx.ProcessedTimestamp = timestamp

					// fetch original output amount

					var client *ethclient.Client
					client, err = ethclient.Dial(rpc)
					defer client.Close()

					if err != nil {
						return nil, err
					}

					var payloadHelper *superform.PayloadHelper
					payloadHelper, err = superform.NewPayloadHelper(common.HexToAddress(config.PayloadHelperContractAddress), client)

					if err != nil {
						return nil, err
					}

					var payload superform.IPayloadHelperDecodedDstPayload
					payload, err = payloadHelper.DecodeCoreStateRegistryPayload(nil, big.NewInt(int64(dstPayloadId)))

					if err != nil {
						return nil, err
					}

					detailedTx.InitialOutputAmount = payload.OutputAmounts[0].String()

					amount := payload.Amounts[0]
					detailedTx.Amount = amount.String()

					superformAddress := fetchDstSuperformAddress(detailedTx)

					var superformContract *superform.Superform
					superformContract, err = superform.NewSuperform(common.HexToAddress(superformAddress), client)

					if err != nil {
						return nil, err
					}

					var delayedOutputAmount *big.Int
					delayedOutputAmount, err = superformContract.PreviewWithdrawFrom(&bind.CallOpts{BlockNumber: big.NewInt(blockNumber)}, amount)

					if err != nil {
						return nil, err
					}

					detailedTx.DelayedOutputAmount = delayedOutputAmount.String()

					filteredTxs = append(filteredTxs, model.ProcessedTransaction{
						Type:                detailedTx.Type,
						ID:                  detailedTx.ID,
						Hash:                detailedTx.Hash,
						Chain:               detailedTx.Chain,
						Status:              detailedTx.Status,
						SuperformId:         detailedTx.SuperformID,
						InitialTimestamp:    detailedTx.CreationTimestamp / 1000,
						DelayedTimestamp:    detailedTx.ProcessedTimestamp,
						Amount:              detailedTx.Amount,
						InitialOutputAmount: detailedTx.InitialOutputAmount,
						DelayedOutputAmount: detailedTx.DelayedOutputAmount,
					})
				}
			}
		}
	}

	return filteredTxs, nil
}

// includesAMB checks if the transaction used a combination of 2 specified AMBs
func includesAMB(tx *model.DetailedTransaction, ambs []string) bool {
	ambCount := make(map[string]int)
	for _, phase := range tx.Phases {
		for _, row := range phase.Rows {
			for _, amb := range ambs {
				if row.Name == amb {
					ambCount[amb]++
				}
			}
		}
	}

	// Count how many different AMBs are present
	differentAMBs := 0
	for _, count := range ambCount {
		if count > 0 {
			differentAMBs++
		}
	}

	return differentAMBs >= 2 // True if at least two different specified AMBs are found
}

// fetchDstData retrieves the RPC, payloadId and a confirmation timestamp for a withdrawal transaction (checks if the transaction has a "Processing Withdrawal" phase with "Update Hash" row)
func fetchDstData(tx *model.DetailedTransaction, chains []model.Chain) (int64, string, int, int64, error) {
	for _, phase := range tx.Phases {
		if phase.Name == "Processing Withdrawal" {
			for _, row := range phase.Rows {
				if row.Component == "hash" && row.Label == "Update Hash" {
					chainId := row.ChainID

					var rpc string
					for _, chain := range chains {
						if chain.ChainID == chainId {
							for _, r := range chain.Rpcs {
								if strings.Contains(r, "wss") {
									continue // skip websocket endpoints
								}

								if strings.Contains(r, "$") {
									continue // skip endpoints with placeholders for now
									// TODO: implement placeholder replacement (infura, alchemy)
								}

								rpc = r
							}
							break
						}
					}

					blockNumber, err := blockchain.FetchBlockNumber(rpc, row.Hash)

					if err != nil {
						return -1, "", -1, -1, err
					}

					timestamp, err := blockchain.FetchBlockTimestamp(rpc, blockNumber)

					if err != nil {
						return -1, "", -1, -1, err
					}

					dstPayloadId, err := fetchDstPayloadId(tx)

					if err != nil {
						return -1, "", -1, -1, err
					}

					return timestamp, rpc, dstPayloadId, blockNumber, nil
				}
			}
		}
	}

	return -1, "", -1, -1, nil
}

// fetchDstPayloadId retrieves the destination payload ID
func fetchDstPayloadId(tx *model.DetailedTransaction) (int, error) {
	for _, phase := range tx.Phases {
		if phase.Name == "Initiation" {
			for _, row := range phase.Rows {
				if row.Component == "value" && row.Label == "Destination Payload ID" {
					id, err := strconv.Atoi(row.Value.(string))

					if err != nil {
						return -1, err
					}

					return id, nil
				}
			}
		}
	}

	return -1, nil
}

// fetchDstSuperformAddress retrieves the address of the destination Superform contract
func fetchDstSuperformAddress(tx *model.DetailedTransaction) string {
	for _, row := range tx.Rows {
		if row.Component == "superform" && row.Label == "Target Superform" {
			return row.Address
		}
	}

	return ""
}
