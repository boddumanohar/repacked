package binpacking

import (
	"errors"
	"math"
	"sort"
)

func PackOrder(packSizes []int, orderSize int) ([]int, error) {
	if len(packSizes) == 0 || orderSize < 1 {
		return nil, errors.New("invalid input: packSizes must be non-empty and orderSize must be positive")
	}

	sort.Ints(packSizes) // Sort pack sizes in ascending order for dynamic programming

	if orderSize < packSizes[0] {
		return nil, errors.New("order size is smaller than the smallest pack size")
	}

	dp := make([]int, orderSize+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	bestCombination := make([][]int, len(dp))
	for i := range bestCombination {
		bestCombination[i] = make([]int, len(packSizes))
	}

	// Dynamic Programming to find the minimum number of packs
	for i := 1; i <= orderSize; i++ {
		for j, size := range packSizes {
			if size <= i && dp[i-size]+1 < dp[i] {
				dp[i] = dp[i-size] + 1
				copy(bestCombination[i], bestCombination[i-size])
				bestCombination[i][j]++
			}
		}
	}

	if dp[orderSize] == math.MaxInt32 {
		return nil, errors.New("no combination of packs can satisfy the order size")
	}

	return bestCombination[orderSize], nil
}
