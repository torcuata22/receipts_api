package main

import (
	"receipt-points/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("receipts/process", handlers.ProcessReceipt)
	r.GET("receipts/:id/points", handlers.GetPoints)

	r.Run(":8080")
}
