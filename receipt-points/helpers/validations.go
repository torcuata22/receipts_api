package helpers

import (
	"time"
)

// IsPurchaseDayOdd checks if the purchase date's day is odd.
func IsPurchaseDayOdd(purchaseDate string) bool {
	t, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return false
	}
	return t.Day()%2 != 0
}

// IsTimeBetween2And4 checks if the purchase time is between 2:00pm and 4:00pm.
func IsTimeBetween2And4(purchaseTime string) bool {
	parsedTime, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return false
	}
	return parsedTime.Hour() == 14 || parsedTime.Hour() == 15 || (parsedTime.Hour() == 16 && parsedTime.Minute() == 0)
}
