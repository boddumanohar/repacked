package binpacking

import (
	"errors"
	"math"
	"repack/utils"
	"sort"

	"go.uber.org/zap"
)

func maxPackSize(packSizes []int) int {
	max := 0
	for _, size := range packSizes {
		if size > max {
			max = size
		}
	}
	return max
}

func PackOrder(packSizes []int, orderSize int) ([]int, error) {
	utils.Logger.Info("", zap.Int("orderSize", orderSize))
	if len(packSizes) == 0 || orderSize < 1 {
		return nil, errors.New("invalid input: packSizes must be non-empty and orderSize must be positive")
	}

	sort.Ints(packSizes) // Sort pack sizes in ascending order for dynamic programming

	if orderSize < packSizes[0] {
		return nil, errors.New("ordersize smaller than the smallest item")
	}

	// Initialize DP array with max value
	dp := make([]int, orderSize+maxPackSize(packSizes)+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	bestCombination := make([][]int, len(dp))
	for i := range bestCombination {
		bestCombination[i] = make([]int, len(packSizes))
	}

	// Dynamic Programming to find the minimum number of packs
	for i := 1; i < len(dp); i++ {
		for j, size := range packSizes {
			if size <= i && dp[i-size]+1 < dp[i] {
				dp[i] = dp[i-size] + 1
				copy(bestCombination[i], bestCombination[i-size])
				bestCombination[i][j]++
			}
		}
	}

	// Find the combination that fulfills the order size or goes over by the smallest amount
	bestFit := math.MaxInt32
	var bestFitCombination []int
	for i := orderSize; i < len(dp); i++ {
		if dp[i] != math.MaxInt32 && i-bestFit < 0 {
			bestFit = i
			bestFitCombination = bestCombination[i]
		}
	}

	return bestFitCombination, nil
}
