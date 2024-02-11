package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func packOrder(packSizes []int, orderSize int) ([]int, error) {
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

type Packs struct {
	sync.Mutex
	Sizes []int `json:"packSizes"`
}

var packets Packs

func postHandler(c *gin.Context) {
	var tempPacks Packs
	if err := c.BindJSON(&tempPacks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, size := range tempPacks.Sizes {
		if size <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "pack sizes must be positive integers"})
			return
		}
	}

	packets.Lock()
	packets.Sizes = tempPacks.Sizes
	packets.Unlock()

	c.JSON(http.StatusOK, gin.H{"packs": packets.Sizes})
}

func getHandler(c *gin.Context) {
	orderSizeStr := c.Query("orderSize")
	orderSize, err := strconv.Atoi(orderSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid orderSize parameter"})
		return
	}

	packets.Lock()
	defer packets.Unlock()

	if len(packets.Sizes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty packet sizes"})
		return
	}

	packs, err := packOrder(packets.Sizes, orderSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"packs": packs, "packSizes": packets.Sizes})
}

func main() {
	r := gin.Default()
	r.POST("/pack", postHandler)
	r.GET("/pack", getHandler)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
