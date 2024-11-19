package utils

import (
	"math"
	"strconv"
	"strings"
	"time"

	"receipt-service/models"
)

func CalculatePoints(receipt *models.Receipt) int {
	points := 0

	// Rule 1: One point per alphanumeric character in retailer name
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount
	if isRoundDollar(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if isMultiple(receipt.Total, 0.25) {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points based on item description length
	for _, item := range receipt.Items {
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the purchase day is odd
	if isOddDay(receipt.PurchaseDate) {
		points += 6
	}

	// Rule 7: 10 points if time is between 2:00 PM and 4:00 PM
	if isAfternoon(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

// Helper function to check if a character is alphanumeric
func isAlphanumeric(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

// Helper function to check if the total is a round dollar amount
func isRoundDollar(total string) bool {
	val, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return val == float64(int(val))
}

// Helper function to check if the total is a multiple of a given factor
func isMultiple(total string, factor float64) bool {
	val, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return math.Mod(val, factor) == 0
}

// Helper function to check if the purchase day is odd
func isOddDay(date string) bool {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	return t.Day()%2 != 0
}

// Helper function to check if the time is between 2:00 PM and 4:00 PM
func isAfternoon(timeStr string) bool {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	return t.Hour() >= 14 && t.Hour() < 16
}
