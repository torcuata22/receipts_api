package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCountAlphanumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Target", 6},
		{"M&M Corner Market", 14},
		{"!@#$%^&*()", 0},
		{"123_##)", 3},
	}

	for _, test := range tests {
		result := countAlphanumeric(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestIsRoundDollarAmount(t *testing.T) {
	tests := []struct {
		total    string
		expected bool
	}{
		{"35", true},
		{"35.35", false},
		{"10.01", false},
		{"10.00", true},
	}

	for _, test := range tests {
		result := isRoundDollarAmount(test.total)
		if result != test.expected {
			t.Errorf("For input %s, expected %t, got %t", test.total, test.expected, result)
		}
	}
}

func TestIsMultipleOfQuarter(t *testing.T) {
	tests := []struct {
		total    string
		expected bool
	}{
		{"35", true},
		{"35.50", true},
		{"10.01", false},
		{"10.57", false},
	}

	for _, test := range tests {
		result := isMultipleOfQuarter(test.total)
		if result != test.expected {
			t.Errorf("For input %s, expected %t, got %t", test.total, test.expected, result)
		}
	}
}

func TestIsPurchaseDayOdd(t *testing.T) {
	tests := []struct {
		purchaseDate string
		expected     bool
	}{
		{"2024-05-01", true},
		{"2024-05-02", false},
		{"2024-11-03", true},
		{"2024-11-04", false},
	}

	for _, test := range tests {
		result := isPurchaseDayOdd(test.purchaseDate)
		if result != test.expected {
			t.Errorf("For input %s, expected %t, got %t", test.purchaseDate, test.expected, result)
		}
	}
}

func TestIsTimeBetween2And4(t *testing.T) {
	tests := []struct {
		purchaseTime string
		expected     bool
	}{
		{"12:00", false},
		{"12:59", false},
		{"13:00", false},
		{"13:59", false},
		{"14:00", true},
		{"14:01", true},
		{"14:59", true},
		{"15:00", true},
		{"16:00", true},
		{"16:01", false},
		{"16:59", false},
		{"19:30", false},
	}

	for _, test := range tests {
		result := isTimeBetween2And4(test.purchaseTime)
		if result != test.expected {
			t.Errorf("For input %s, expected %t, got %t", test.purchaseTime, test.expected, result)
		}
	}
}

func TestProcessReceipt(t *testing.T) {
	// Create a new gin engine for testing
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register the routes
	r.POST("receipts/process", processReceipt)
	r.GET("receipts/:id/points", getPoints)

	// Create a sample receipt
	receipt := Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2024-05-01",
		PurchaseTime: "14:30",
		Items: []struct {
			ShortDescription string `json:"shortDescription"`
			Price            string `json:"price"`
		}{
			{"Item 1", "10.00"},
			{"Item 2", "20.00"},
		},
		Total: "30.00",
	}

	// Marshal the receipt to JSON
	jsonBytes, err := json.Marshal(receipt)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	w := httptest.NewRecorder()

	// Call the processReceipt function
	r.ServeHTTP(w, req)

	// Check if the response status code is 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Unmarshal the response to get the receipt ID
	var resp map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	// Get the points for the receipt
	req, err = http.NewRequest("GET", "/receipts/"+resp["id"]+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check if the response status code is 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	// Unmarshal the response to get the points
	var pointsResp map[string]int
	err = json.Unmarshal(w.Body.Bytes(), &pointsResp)
	if err != nil {
		t.Fatal(err)
	}

	// Calculate the expected points
	expectedPoints := calculatePoints(receipt)

	// Check if the points are correct
	if pointsResp["points"] != expectedPoints {
		t.Errorf("Expected points %d, got %d", expectedPoints, pointsResp["points"])
	}
}
