package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/simonovic86/superform/config"
	"github.com/simonovic86/superform/internal/model"
	"github.com/simonovic86/superform/pkg/analytics"
	"github.com/simonovic86/superform/pkg/api"
	"github.com/simonovic86/superform/pkg/processor"
)

func main() {
	// Custom usage message
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("\nAMBs should be two out of the three values: Layerzero, Hyperlane, Wormhole")
	}

	// Define the flags
	sourceChainID := flag.Int("sourceChainID", 42161, "Source chain ID")
	destinationChainID := flag.Int("destinationChainID", 10, "Destination chain ID")
	ambsInput := flag.String("ambs", "Layerzero,Hyperlane", "Comma-separated list of two AMBs out of Layerzero, Hyperlane, Wormhole")
	outputPath := flag.String("output", "transactions.json", "Path of the output file for processed transactions (optional)")

	// Parse the flags
	flag.Parse()

	// Validate AMBs
	validAmbs := map[string]bool{"Layerzero": true, "Hyperlane": true, "Wormhole": true}
	ambs := strings.Split(*ambsInput, ",")

	if len(ambs) != 2 || !validAmbs[ambs[0]] || !validAmbs[ambs[1]] || ambs[0] == ambs[1] {
		fmt.Println("error: AMBs must be two different values out of Layerzero, Hyperlane, Wormhole")
		os.Exit(1)
	}

	// Output the parsed values
	fmt.Printf("source Chain ID: %d\n", *sourceChainID)
	fmt.Printf("destination Chain ID: %d\n", *destinationChainID)
	fmt.Printf("AMBs: %v\n", ambs)

	if *outputPath != "" {
		fmt.Printf("processed transactions will be written to: %s\n", *outputPath)
	}

	txType := 1 // "withdrawal" filtering for withdrawals

	chains, err := api.FetchSupportedChains()

	if err != nil {
		log.Fatalf("error fetching supported chains: %v\n", err)
	}

	offset := 0
	limit := 1000
	transactions := make([]model.Transaction, 0)

	fmt.Printf("fetching maximum %d transactions for demonstration purposes...\n", config.MaxProcessedTransactions)

	for {
		var result []model.Transaction
		result, err = api.FetchTransactions(offset, limit)

		if err != nil {
			log.Fatalf("error fetching transactions: %v\n", err)
		}

		if len(result) == 0 {
			break
		}

		transactions = append(transactions, result...)

		if len(transactions) >= config.MaxProcessedTransactions {
			break
		}

		offset += limit
	}

	fmt.Printf("processing %d transactions...\n", len(transactions))

	processedTxs, err := processor.ProcessTransactions(transactions, txType, *sourceChainID, *destinationChainID, ambs, chains)

	if err != nil {
		log.Fatalf("error processing transactions: %v\n", err)
	}

	// Calculate the total delay time and the average
	averageTime := analytics.CalculateAverageTime(processedTxs)

	// Calculate the Time Adjustment Factor
	timeAdjustmentFactor, err := analytics.CalculateTimeAdjustmentFactor(processedTxs)

	if err != nil {
		log.Fatalf("error calculating time adjustment factor: %v\n", err)
	}

	fmt.Printf("\naverage time until the transaction can be processed is %v and the time adjustment factor: %.2f%%\n", averageTime, timeAdjustmentFactor)

	// Write the processed transactions to a file
	file, _ := json.MarshalIndent(processedTxs, "", " ")

	_ = ioutil.WriteFile(*outputPath, file, 0644)
}
