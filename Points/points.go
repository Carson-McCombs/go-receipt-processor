package points

import (
	date "GoReceiptProcessor/Date"
	receipt "GoReceiptProcessor/Receipt"
	receiptitem "GoReceiptProcessor/Receipt/ReceiptItem"
	time "GoReceiptProcessor/Time"
	"math"
	"regexp"
	"strings"
)

var alphanumericRegex = regexp.MustCompile("[[:alnum:]]")

func CalculatePoints(receipt receipt.Receipt) int64 {
	var points int64 = 0
	points += getPointsForAlphanumericalCharacters(receipt.Retailer)
	points += getPointsForRoundDollarAmount(receipt.Total)
	points += getPointsForMultipleOf25Cents(receipt.Total)
	points += getPointsForNumberOfItems(receipt.Items)
	points += getSumOfPointsForItemDescriptionAndPrice(receipt.Items)
	points += getPointsForOddPurchaseDate(receipt.PurchaseDate)
	points += getPointsForTimeOfDay(receipt.PurchaseTime)
	return points
}

func getPointsForAlphanumericalCharacters(str string) int64 {
	alphanumericString := alphanumericRegex.FindAllString(str, -1)
	alphanumericCount := len(alphanumericString)
	return int64(alphanumericCount)
}

func isMultipleOfFloat(value float64, multiple float64) bool {
	modValue := math.Mod(value, multiple)
	return modValue == 0
}

func isMultipleOfUint(value uint, multiple uint) bool {
	remainder := (value - uint(value/multiple)*multiple)
	return remainder == 0
}
func getPointsForRoundDollarAmount(currency float64) int64 {
	if isMultipleOfFloat(currency, 1) {
		return 50
	}
	return 0
}

func getPointsForMultipleOf25Cents(currency float64) int64 {
	if isMultipleOfFloat(currency, 0.25) {
		return 25
	}
	return 0
}

func getPointsForNumberOfItems(items []receiptitem.ReceiptItem) int64 {
	points := int64(len(items)/2) * 5
	return points
}

func getPointsForItemDescriptionAndPrice(item receiptitem.ReceiptItem) int64 {
	trimmedDescription := strings.Trim(item.ShortDescription, " ")
	if isMultipleOfUint(uint(len(trimmedDescription)), 3) {
		points := int64(math.Ceil(item.Price * 0.2))
		return points
	}
	return 0
}

func getSumOfPointsForItemDescriptionAndPrice(items []receiptitem.ReceiptItem) int64 {
	var points int64 = 0
	for _, item := range items {
		points += getPointsForItemDescriptionAndPrice(item)
	}
	return points
}

func getPointsForOddPurchaseDate(date date.Date) int64 {
	if !isMultipleOfUint(uint(date.Day), 2) {
		return 6
	}
	return 0
}

func getPointsForTimeOfDay(time time.Time) int64 {
	if time.Hour < 14 || time.Hour >= 16 {
		return 0
	}
	if time.Hour == 14 && time.Minute == 0 {
		return 0
	}
	return 10
}
