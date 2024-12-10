package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, exists := receiptStore[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points})
}
