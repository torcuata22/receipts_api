package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []struct {
		ShortDescription string `json:"shortDescription"`
		Price            string `json:"price"`
	} `json:"items"`
	Total string `json:"total"`
}

var receiptStore = make(map[string]int) //used to store receipts in memory, not necessary if I were using DB

func processReceipt(c *gin.Context) {
	var receipt Receipt

	// Parse JSON body
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidJSON.Error()})
		return
	}

	// Calculate points (placeholder logic)
	points := calculatePoints(receipt)

	// Generate a unique receipt ID
	id := uuid.New().String()

	// Store receipt ID and points
	receiptStore[id] = points

	// Return the receipt ID
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func calculatePoints(r Receipt) int {
	points := 0

	fmt.Println("Rule 1: ", countAlphanumeric(r.Retailer))

	//Rule 1: add point for every alphanumeric character in retailers name
	points += countAlphanumeric(r.Retailer)

	//Rule2: add 50 points if total is a round dollar amount
	fmt.Println("Rule 2: ", isRoundDollarAmount(r.Total))
	if isRoundDollarAmount(r.Total) {
		points += 50
	}

	//Rule3: add 25 points if total is a multiple of 0.25
	fmt.Println("Rule 3: ", isMultipleOfQuarter(r.Total))
	if isMultipleOfQuarter(r.Total) {
		points += 25
	}

	//Rule4: add 5 points for every two items on the receipt
	fmt.Println("Rule 4: ", points)
	points += (len(r.Items) / 2) * 5

	//Rule5: price X 0.2 if length of description is a multiple of 3
	fmt.Println("Rule 5: ", points)
	for _, item := range r.Items {
		descLen := len(strings.TrimSpace(item.ShortDescription))
		if descLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	//Rule6: add 6 points if the day of the purchase date is odd
	fmt.Println("Rule 6: ", isPurchaseDayOdd(r.PurchaseDate))
	if isPurchaseDayOdd(r.PurchaseDate) {
		points += 6
	}

	//Rule7: add 10 points if the time of purchase is after 2:00pm and before 4:00pm
	fmt.Println("Rule 7: ", isTimeBetween2And4(r.PurchaseTime))
	if isTimeBetween2And4(r.PurchaseTime) {
		points += 10
	}
	return points
}

func countAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			count++
		}
	}
	return count
}

func isRoundDollarAmount(total string) bool {
	price, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return price == float64(int(price))
}

func isMultipleOfQuarter(total string) bool {
	price, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	const epsilon = 0.00001
	return math.Abs(math.Mod(price, 0.25)) < epsilon
}

func isPurchaseDayOdd(purchaseDate string) bool {
	// Parse the purchase date string into a time.Time object
	t, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return false
	}
	return t.Day()%2 != 0
}

func isTimeBetween2And4(purchaseTime string) bool {
	// Parse the purchase time string into a time.Time object
	parsedTime, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return false
	}
	return parsedTime.Hour() == 14 || parsedTime.Hour() == 15 || (parsedTime.Hour() == 16 && parsedTime.Minute() == 0)
}

func getPoints(c *gin.Context) {

	id := c.Param("id")

	points, exists := receiptStore[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": ErrReceiptNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"points": points,
	})

}

func main() {
	r := gin.Default()
	r.POST("receipts/process", processReceipt)
	r.GET("receipts/:id/points", getPoints)

	r.Run(":8080")
}
