package points

import (
	date "go-receipt-processor/Date"
	receipt "go-receipt-processor/Receipt"
	receiptitem "go-receipt-processor/Receipt/ReceiptItem"
	utils "go-receipt-processor/TestingUtils"
	time "go-receipt-processor/Time"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_CalculatePoints(t *testing.T) {
	var testCases []utils.CreationTestingData[receipt.Receipt, int64] = []utils.CreationTestingData[receipt.Receipt, int64]{
		{Argument: receipt.Receipt{Retailer: "Target", PurchaseDate: date.Date{Year: 2022, Month: 01, Day: 01}, PurchaseTime: time.Time{Hour: 13, Minute: 01}, Total: 35.35,
			Items: []receiptitem.ReceiptItem{
				{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
				{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
				{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
				{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00}}},
			ExpectedResult: 28,
			ExpectedErr:    nil,
		},
		{Argument: receipt.Receipt{Retailer: "M&M Corner Market", PurchaseDate: date.Date{Year: 2022, Month: 03, Day: 20}, PurchaseTime: time.Time{Hour: 14, Minute: 33}, Total: 9.00,
			Items: []receiptitem.ReceiptItem{
				{ShortDescription: "Gatorade", Price: 2.25},
				{ShortDescription: "Gatorade", Price: 2.25},
				{ShortDescription: "Gatorade", Price: 2.25},
				{ShortDescription: "Gatorade", Price: 2.25},
			}},
			ExpectedResult: 109,
		},
	}
	for _, testCase := range testCases {
		result := CalculatePoints(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("calculate points ( %+v ): expected result ( %+v ) got result ( %+v )", testCase.Argument, testCase.ExpectedResult, result)
		}

	}
}

func Test_getPointsForAlphanumericalCharacters(t *testing.T) {
	var testCases []utils.CreationTestingData[string, int64] = []utils.CreationTestingData[string, int64]{
		{Argument: "", ExpectedResult: 0},
		{Argument: "Target", ExpectedResult: 6},
		{Argument: "P2W Games", ExpectedResult: 8},
		{Argument: "Firehouse Subs", ExpectedResult: 13},
		{Argument: "Waffle-House", ExpectedResult: 11},
		{Argument: "???   Game'n' - sSTop3", ExpectedResult: 11},
		{Argument: "?!____=-@=-=^^-?", ExpectedResult: 0},
	}
	for _, testCase := range testCases {
		result := getPointsForAlphanumericalCharacters(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("get points for alphanumeric characters ( %+v ): expected result ( %+v ) got result ( %+v ), 1 point per alphanumeric character ( i.e. 0-9, a-z, A-Z )", testCase.Argument, testCase.ExpectedResult, result)
		}

	}
}

func Test_getPointsForRoundDollarAmount(t *testing.T) {
	var testCases []utils.CreationTestingData[float64, int64] = []utils.CreationTestingData[float64, int64]{
		{Argument: 0.00, ExpectedResult: 50},
		{Argument: 1.00, ExpectedResult: 50},
		{Argument: 5.25, ExpectedResult: 0},
		{Argument: 10.50, ExpectedResult: 0},
		{Argument: 5.00, ExpectedResult: 50},
		{Argument: 3.76, ExpectedResult: 0},
		{Argument: 20.29, ExpectedResult: 0},
		{Argument: 100.50, ExpectedResult: 0},
		{Argument: 0.50, ExpectedResult: 00},
		{Argument: -6.30, ExpectedResult: 0},
		{Argument: -9.00, ExpectedResult: 50},
	}
	for _, testCase := range testCases {
		result := getPointsForRoundDollarAmount(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("get points for alphanumeric characters ( %+v ): expected result ( %+v ) got result ( %+v ), should get 50 points if it is a round dollar amount, and 0 points otherwise", testCase.Argument, testCase.ExpectedResult, result)
		}

	}
}

func Test_getPointsForNumberOfItems(t *testing.T) {
	var testCase []utils.CreationTestingData[[]receiptitem.ReceiptItem, int64] = []utils.CreationTestingData[[]receiptitem.ReceiptItem, int64]{
		{Argument: []receiptitem.ReceiptItem{}, ExpectedResult: 0},
		{Argument: []receiptitem.ReceiptItem{{}}, ExpectedResult: 0},
		{Argument: []receiptitem.ReceiptItem{{}, {}}, ExpectedResult: 5},
		{Argument: []receiptitem.ReceiptItem{{}, {}, {}}, ExpectedResult: 5},
		{Argument: []receiptitem.ReceiptItem{{}, {}, {}, {}}, ExpectedResult: 10},
		{Argument: []receiptitem.ReceiptItem{{}, {}, {}, {}, {}}, ExpectedResult: 10},
		{Argument: []receiptitem.ReceiptItem{{}, {}, {}, {}, {}, {}}, ExpectedResult: 15},
		{Argument: []receiptitem.ReceiptItem{{}, {}, {}, {}, {}, {}, {}}, ExpectedResult: 15},
		{Argument: []receiptitem.ReceiptItem{{}, {}, {}, {}, {}, {}, {}, {}}, ExpectedResult: 20},
	}

	for _, testCase := range testCase {
		result := getPointsForNumberOfItems(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("get points for number of items ( %+v ): expected result ( %+v ) got result ( %+v ), should be 5 points for every 2 items", testCase.Argument, testCase.ExpectedResult, result)
		}

	}
}

func Test_getPointsForItemDescriptionAndPrice(t *testing.T) {
	var testCases []utils.CreationTestingData[receiptitem.ReceiptItem, int64] = []utils.CreationTestingData[receiptitem.ReceiptItem, int64]{
		{Argument: receiptitem.DefaultReceiptItem(), ExpectedResult: 0},
		{Argument: receiptitem.ReceiptItem{}, ExpectedResult: 0},
		{Argument: receiptitem.ReceiptItem{ShortDescription: "Mountain Dew 12PK", Price: 6.49}, ExpectedResult: 0},
		{Argument: receiptitem.ReceiptItem{ShortDescription: "Emils Cheese Pizza", Price: 12.25}, ExpectedResult: 3},
		{Argument: receiptitem.ReceiptItem{ShortDescription: "Knorr Creamy Chicken", Price: 1.26}, ExpectedResult: 0},
		{Argument: receiptitem.ReceiptItem{ShortDescription: "Doritos Nacho Cheese", Price: 3.35}, ExpectedResult: 0},
		{Argument: receiptitem.ReceiptItem{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00}, ExpectedResult: 3},
		{Argument: receiptitem.ReceiptItem{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: -12.00}, ExpectedResult: -2}, //because price was never set to only be positive, the points that this receipt can return can actually be negative
	}
	for _, testCase := range testCases {
		result := getPointsForItemDescriptionAndPrice(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("get points for item description and price ( %+v ): expected result ( %+v ) got result ( %+v ), if the trimmed description length is a multiple of 3, return points equal to the price * 0.2 rounded up to the nearest integer", testCase.Argument, testCase.ExpectedResult, result)
		}

	}
}

func Test_getPointsForOddPurchaseDate(t *testing.T) {
	var testCases []utils.CreationTestingData[date.Date, int64] = []utils.CreationTestingData[date.Date, int64]{
		{Argument: date.Date{}, ExpectedResult: 0},
		{Argument: date.Date{Year: 2001, Month: 3, Day: 15}, ExpectedResult: 6},
		{Argument: date.Date{Year: 1995, Month: 5, Day: 7}, ExpectedResult: 6},
		{Argument: date.Date{Year: 2007, Month: 15, Day: 18}, ExpectedResult: 0},
		{Argument: date.Date{Year: 7, Month: 1, Day: 1}, ExpectedResult: 6},
	}
	for _, testCase := range testCases {
		result := getPointsForOddPurchaseDate(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("get points for odd purchase date ( %+v ): expected result ( %+v ) got result ( %+v ), if the day of the purchase date is odd, return 6 points, and 0 points otherwise", testCase.Argument, testCase.ExpectedResult, result)
		}
	}
}

func Test_getPointsForTimeOfDay(t *testing.T) {
	var testCases []utils.CreationTestingData[time.Time, int64] = []utils.CreationTestingData[time.Time, int64]{
		{Argument: time.Time{}, ExpectedResult: 0},
		{Argument: time.Time{}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 2, Minute: 31}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 3, Minute: 27}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 4, Minute: 59}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 9, Minute: 50}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 14, Minute: 00}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 14, Minute: 01}, ExpectedResult: 10},
		{Argument: time.Time{Hour: 15, Minute: 00}, ExpectedResult: 10},
		{Argument: time.Time{Hour: 15, Minute: 01}, ExpectedResult: 10},
		{Argument: time.Time{Hour: 15, Minute: 59}, ExpectedResult: 10},
		{Argument: time.Time{Hour: 16, Minute: 00}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 16, Minute: 31}, ExpectedResult: 0},
		{Argument: time.Time{Hour: 26, Minute: 31}, ExpectedResult: 0}, //all values that return 10 points are valid times, any invalid time, i.e. 50:106, will return 0
	}
	for _, testCase := range testCases {
		result := getPointsForTimeOfDay(testCase.Argument)
		if !cmp.Equal(result, testCase.ExpectedResult) {
			t.Fatalf("get points time of day ( %+v ): expected result ( %+v ) got result ( %+v ), if the purchase time is after 2:00 pm or 14:00 and before 4:00 pm or 16:00, return 10 points, and 0 points otherwise", testCase.Argument, testCase.ExpectedResult, result)
		}
	}
}
