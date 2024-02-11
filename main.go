package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func packOrder(packSizes []int, orderSize int) []int {
	sort.Ints(packSizes) // Sort pack sizes in ascending order for dynamic programming

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

	return bestFitCombination
}

func maxPackSize(packSizes []int) int {
	max := 0
	for _, size := range packSizes {
		if size > max {
			max = size
		}
	}
	return max
}

type Packs struct {
	sync.Mutex
	Sizes []int `json:"packSizes"`
}

var packets Packs

func postHandler(c *gin.Context) {
	packets.Lock()
	if err := c.BindJSON(&packets); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	packets.Unlock()

	c.JSON(200, gin.H{"packs": packets.Sizes})
}

func getHandler(c *gin.Context) {
	orderSizeStr := c.Query("orderSize")
	orderSize, err := strconv.Atoi(orderSizeStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid orderSize parameter"})
		return
	}

	if len(packets.Sizes) == 0 {
		c.JSON(400, gin.H{"error": "empty packet sizes"})
		return
	}

	packs := packOrder(packets.Sizes, orderSize)
	c.JSON(200, gin.H{"packs": packs, "packSizes": packets.Sizes})
}

func main() {
	r := gin.Default()
	r.POST("/pack", postHandler)
	r.GET("/pack", getHandler)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
