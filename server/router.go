package server

import (
	"net/http"
	"strconv"
	"sync"

	binpacking "repack/lib"

	"github.com/gin-gonic/gin"
)

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

func optionsHandler(c *gin.Context) {
	// Set CORS headers
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Return ok status
	c.AbortWithStatus(http.StatusOK)
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

	packs, err := binpacking.PackOrder(packets.Sizes, orderSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"packs": packs, "packSizes": packets.Sizes})
}
