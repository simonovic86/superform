package analytics

import (
	"fmt"
	"math/big"
	"time"

	"github.com/simonovic86/superform/internal/model"
)

// CalculateAverageTime calculates the average delay time for transactions
func CalculateAverageTime(transactions []model.ProcessedTransaction) time.Duration {
	var totalProcessingTime time.Duration

	for _, tx := range transactions {
		processingTime := time.Duration(tx.DelayedTimestamp-tx.InitialTimestamp) * time.Second
		totalProcessingTime += processingTime
	}

	return totalProcessingTime / time.Duration(len(transactions))
}

// CalculateTimeAdjustmentFactor calculates the time-weighted average percentage change in output amount.
func CalculateTimeAdjustmentFactor(transactions []model.ProcessedTransaction) (*big.Float, error) {
	totalTimeWeightedPercentageChange := big.NewFloat(0)
	totalTime := int64(0)

	for _, tx := range transactions {
		if tx.InitialOutputAmount == "0" {
			continue // Avoid division by zero
		}

		timeDelay := tx.DelayedTimestamp - tx.InitialTimestamp
		if timeDelay <= 0 {
			continue // Ensure we have a positive time delay
		}

		delayedOutputAmount, ok := new(big.Float).SetString(tx.DelayedOutputAmount)
		if !ok {
			continue // Skip if we can't parse the delayed output amount
		}

		initialOutputAmount, ok := new(big.Float).SetString(tx.InitialOutputAmount)
		if !ok {
			continue // Skip if we can't parse the initial output amount
		}

		change := new(big.Float).Sub(delayedOutputAmount, initialOutputAmount)
		percentageChange := new(big.Float).Quo(change, initialOutputAmount)
		timeWeightedChange := new(big.Float).Mul(percentageChange, new(big.Float).SetInt64(timeDelay))

		totalTimeWeightedPercentageChange.Add(totalTimeWeightedPercentageChange, timeWeightedChange)
		totalTime += timeDelay
	}

	if totalTime == 0 {
		return nil, fmt.Errorf("no valid transactions for calculation")
	}

	averageTimeWeightedPercentageChange := new(big.Float).Quo(totalTimeWeightedPercentageChange, new(big.Float).SetInt64(totalTime))
	timeAdjustmentFactor := new(big.Float).Mul(averageTimeWeightedPercentageChange, big.NewFloat(100))

	return timeAdjustmentFactor, nil
}
