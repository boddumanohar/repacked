package main

import (
	"fmt"
	"math"
	"sort"
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

func packOrder(packSizes []int, orderSize int) []int {
	sort.Ints(packSizes) // Sort in asc order

	// Initialize Dynamic programming array with max value
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

	return bestFitCombination
}

func main() {
	packSizes := []int{23, 31, 53}
	orderSize := 263

	packs := packOrder(packSizes, orderSize)
	fmt.Printf("Order size %d packed in: ", orderSize)
	for i, num := range packs {
		if num > 0 {
			fmt.Printf("%d x %d ", num, packSizes[i])
		}
	}
	fmt.Println()
}
