package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/simonovic86/superform/config"
	"github.com/simonovic86/superform/internal/model"
)

// FetchSupportedChains fetches supported chains from the Superform API
func FetchSupportedChains() ([]model.Chain, error) {
	url := fmt.Sprintf("%s/supported/chains", config.BaseApiURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var chains []model.Chain
	if err := json.NewDecoder(resp.Body).Decode(&chains); err != nil {
		return nil, err
	}

	return chains, nil
}

// FetchTransactions fetches transactions from the Superform API
func FetchTransactions(offset int, limit int) ([]model.Transaction, error) {
	url := fmt.Sprintf("%s/explorer/transactions?offset=%d&limit=%d", config.BaseApiURL, offset, limit)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var transactions []model.Transaction

	if err = json.NewDecoder(resp.Body).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

// FetchDetailedTransaction fetches a detailed view of a transaction by hash
func FetchDetailedTransaction(txHash string) (*model.DetailedTransaction, error) {
	url := fmt.Sprintf("%s/explorer/transaction/%s", config.BaseApiURL, txHash)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Since the API response is an array, we unmarshal into a slice of DetailedTransaction
	var detailedTxs []model.DetailedTransaction
	if err = json.Unmarshal(body, &detailedTxs); err != nil {
		return nil, err
	}

	// Check if we received at least one detailed transaction
	if len(detailedTxs) == 0 {
		return nil, fmt.Errorf("no detailed transaction data found for hash: %s", txHash)
	}

	// Return the first item of the array
	return &detailedTxs[0], nil
}
