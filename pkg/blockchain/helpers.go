package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// FetchBlockNumber retrieves the block number for a given transaction hash.
func FetchBlockNumber(rpcURL string, txHash string) (int64, error) {
	client, err := ethclient.Dial(rpcURL)

	if err != nil {
		return 0, err
	}
	defer client.Close()

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))

	if err != nil {
		return 0, err
	}

	return receipt.BlockNumber.Int64(), nil
}

// FetchBlockTimestamp queries an RPC endpoint for the timestamp of a given block number.
func FetchBlockTimestamp(rpcURL string, blockNumber int64) (int64, error) {
	client, err := ethclient.Dial(rpcURL)

	if err != nil {
		return 0, err
	}
	defer client.Close()

	header, err := client.HeaderByNumber(context.Background(), big.NewInt(blockNumber))

	if err != nil {
		return 0, err
	}

	return int64(header.Time), nil
}
