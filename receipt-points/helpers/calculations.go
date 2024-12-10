package helpers

import (
	"math"
	"strconv"
)

// CountAlphanumeric counts the alphanumeric characters in a string.
func CountAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			count++
		}
	}
	return count
}

// IsRoundDollarAmount checks if the total is a round dollar amount.
func IsRoundDollarAmount(total string) bool {
	price, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return price == float64(int(price))
}

// IsMultipleOfQuarter checks if the total is a multiple of 0.25.
func IsMultipleOfQuarter(total string) bool {
	price, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	const epsilon = 0.00001
	return math.Abs(math.Mod(price, 0.25)) < epsilon
}
