package handlers

import (
	"net/http"
	"strconv"

	"receipt-points/helpers"
	"receipt-points/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receiptStore = make(map[string]int)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	points := CalculatePoints(receipt)
	id := uuid.New().String()
	receiptStore[id] = points

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func CalculatePoints(r models.Receipt) int {
	points := 0
	points += helpers.CountAlphanumeric(r.Retailer)
	if helpers.IsRoundDollarAmount(r.Total) {
		points += 50
	}
	if helpers.IsMultipleOfQuarter(r.Total) {
		points += 25
	}
	points += (len(r.Items) / 2) * 5

	for _, item := range r.Items {
		descLen := len(item.ShortDescription)
		if descLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(price * 0.2)
		}
	}
	if helpers.IsPurchaseDayOdd(r.PurchaseDate) {
		points += 6
	}
	if helpers.IsTimeBetween2And4(r.PurchaseTime) {
		points += 10
	}
	return points
}
