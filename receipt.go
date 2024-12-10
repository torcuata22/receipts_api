// package main

// import (
// 	"github.com/google/uuid"
// 	"receipts_api/errors"
// 		)

// type Receipt struct {
// 	Retailer     string `json:"retailer"`
// 	PurchaseDate string `json:"purchaseDate"`
// 	PurchaseTime string `json:"purchaseTime"`
// 	Items        []struct {
// 		ShortDescription string `json:"shortDescription"`
// 		Price            string `json:"price"`
// 	} `json:"items"`
// 	Total string `json:"total"`
// }

// var receiptStore = make(map[string]int) //used to store receipts in memory, not necessary if I were using DB

// func ProcessReceipt(c *gin.Context, receipt Receipt) (string, error) {
// 	// Calculate points (placeholder logic)
// 	points := calculatePoints(receipt)

// 	// Generate a unique receipt ID
// 	id := uuid.New().String()

// 	// Store receipt ID and points
// 	receiptStore[id] = points

// 	return id, nil
// }