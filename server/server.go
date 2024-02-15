package server

import (
	"fmt"
	"repack/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	utils.InitializeLogger()
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/pack", postHandler)
	r.GET("/pack", getHandler)
	r.OPTIONS("/pack", optionsHandler)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
