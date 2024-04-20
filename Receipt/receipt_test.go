package receipt

import (
	date "go-receipt-processor/Date"
	receiptitem "go-receipt-processor/Receipt/ReceiptItem"
	utils "go-receipt-processor/TestingUtils"
	time "go-receipt-processor/Time"
	"strconv"
	"testing"
)

func Test_ParseReceipt(t *testing.T) {

	// Note: most individual structures are tested on their own ( i.e. Time, Date, and Receipt Item ) so these tests are able to be more general in making sure that the errors are still being carried threw from the internal structs to the Receipt structs
	var testCases []utils.CreationTestingData[UnparsedReceipt, Receipt] = []utils.CreationTestingData[UnparsedReceipt, Receipt]{
		{ //Tests a couple of examples that should not throw errors
			Argument: UnparsedReceipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Total:        "35.35",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
			},
			ExpectedResult: Receipt{
				Id:           "0",
				Retailer:     "Target",
				PurchaseDate: date.Date{Year: 2022, Month: 01, Day: 01},
				PurchaseTime: time.Time{Hour: 13, Minute: 01},
				Total:        35.35,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
					{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
					{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
					{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
				}},
		},
		{
			Argument: UnparsedReceipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Total:        "9.00",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
			},

			ExpectedResult: Receipt{
				Id:           "1",
				Retailer:     "M&M Corner Market",
				PurchaseDate: date.Date{Year: 2022, Month: 03, Day: 20},
				PurchaseTime: time.Time{Hour: 14, Minute: 33},
				Total:        9.00,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Gatorade", Price: 2.25},
					{ShortDescription: "Gatorade", Price: 2.25},
					{ShortDescription: "Gatorade", Price: 2.25},
					{ShortDescription: "Gatorade", Price: 2.25},
				}},
		},
		{
			Argument: UnparsedReceipt{
				Retailer:     "Walmart",
				PurchaseDate: "2024-02-29",
				PurchaseTime: "14:30",
				Total:        "10.37",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Banana", Price: "1.12"},
					{ShortDescription: "Sandwich", Price: "5.50"},
					{ShortDescription: "Chips", Price: "1.50"},
				}},
			ExpectedResult: Receipt{
				Id:           "2",
				Retailer:     "Walmart",
				PurchaseDate: date.Date{Year: 2024, Month: 2, Day: 29},
				PurchaseTime: time.Time{Hour: 14, Minute: 30},
				Total:        10.37,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Gatorade", Price: 2.25},
					{ShortDescription: "Banana", Price: 1.12},
					{ShortDescription: "Sandwich", Price: 5.50},
					{ShortDescription: "Chips", Price: 1.50},
				}},
		},
		{
			Argument: UnparsedReceipt{
				Retailer:     "Target",
				PurchaseDate: "2021-04-05",
				PurchaseTime: "12:00",
				Total:        "8.03",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Water", Price: "1.00"},
					{ShortDescription: "Apple", Price: "0.75"},
					{ShortDescription: "Sandwich", Price: "6.28"},
				}},
			ExpectedResult: Receipt{
				Id:           "3",
				Retailer:     "Target",
				PurchaseDate: date.Date{Year: 2021, Month: 04, Day: 05},
				PurchaseTime: time.Time{Hour: 12, Minute: 0},
				Total:        8.03,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Water", Price: 1.00},
					{ShortDescription: "Apple", Price: 0.75},
					{ShortDescription: "Sandwich", Price: 6.28},
				}},
		},
		{
			Argument: UnparsedReceipt{
				Retailer:     "Best Buy",
				PurchaseDate: "2012-11-20",
				PurchaseTime: "15:45",
				Total:        "50.00",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Headphones", Price: "50.00"},
				},
			},
			ExpectedResult: Receipt{
				Id:           "4",
				Retailer:     "Best Buy",
				PurchaseDate: date.Date{Year: 2012, Month: 11, Day: 20},
				PurchaseTime: time.Time{Hour: 15, Minute: 45},
				Total:        50.00,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Headphones", Price: 50.00},
				},
			},
		},
		{
			Argument: UnparsedReceipt{
				Retailer:     "Costco",
				PurchaseDate: "1999-09-30",
				PurchaseTime: "17:30",
				Total:        "25.50",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Toilet Paper", Price: "15.00"},
					{ShortDescription: "Paper Towels", Price: "10.50"},
				},
			},
			ExpectedResult: Receipt{
				Id:           "5",
				Retailer:     "Costco",
				PurchaseDate: date.Date{Year: 1999, Month: 9, Day: 30},
				PurchaseTime: time.Time{Hour: 17, Minute: 30},
				Total:        25.50,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Toilet Paper", Price: 15.00},
					{ShortDescription: "Paper Towels", Price: 10.50},
				},
			},
		},
		{ //Tests if receipts that with items that don't match the price are given
			Argument: UnparsedReceipt{
				Retailer:     "Walmart",
				PurchaseDate: "2024-02-29",
				PurchaseTime: "14:30",
				Total:        "10.00",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Banana", Price: "1.12"},
					{ShortDescription: "Sandwich", Price: "5.50"},
				}},
			ExpectedResult: Receipt{
				Id:           "6",
				Retailer:     "Walmart",
				PurchaseDate: date.Date{Year: 2024, Month: 2, Day: 29},
				PurchaseTime: time.Time{Hour: 14, Minute: 30},
				Total:        10.00,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Gatorade", Price: 2.25},
					{ShortDescription: "Banana", Price: 1.12},
					{ShortDescription: "Sandwich", Price: 5.50},
				}},
			ExpectedErr: ErrInvalidReceipt,
		},
		{ //Tests if errors with time will still be caught
			Argument: UnparsedReceipt{
				Retailer:     "Costco",
				PurchaseDate: "1999-09-30",
				PurchaseTime: "1730",
				Total:        "25.50",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Toilet Paper", Price: "15.00"},
					{ShortDescription: "Paper Towels", Price: "10.50"},
				},
			},
			ExpectedResult: Receipt{
				Id:           "7",
				Retailer:     "Costco",
				PurchaseDate: date.Date{Year: 1999, Month: 9, Day: 30},
				PurchaseTime: time.Time{},
				Total:        25.50,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Toilet Paper", Price: 15.00},
					{ShortDescription: "Paper Towels", Price: 10.50},
				},
			},
			ExpectedErr: ErrParsingReceipt,
		},
		{ //Tests receipt with leap year date
			Argument: UnparsedReceipt{
				Retailer:     "Walmart",
				PurchaseDate: "2024-02-29",
				PurchaseTime: "14:30",
				Total:        "8.87",
				Items: []receiptitem.UnparsedReceiptItem{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Banana", Price: "1.12"},
					{ShortDescription: "Sandwich", Price: "5.50"},
				}},
			ExpectedResult: Receipt{
				Id:           "8",
				Retailer:     "Walmart",
				PurchaseDate: date.Date{Year: 2024, Month: 2, Day: 29},
				PurchaseTime: time.Time{Hour: 14, Minute: 30},
				Total:        8.87,
				Items: []receiptitem.ReceiptItem{
					{ShortDescription: "Gatorade", Price: 2.25},
					{ShortDescription: "Banana", Price: 1.12},
					{ShortDescription: "Sandwich", Price: 5.50},
				}},
			ExpectedErr: nil,
		},
	}
	for i, testCase := range testCases {
		result, err := ParseReceipt(strconv.Itoa(i), testCase.Argument, true)
		errCheck := testCase.CheckTestCase("parse receipt", result, err, false)
		if errCheck != nil {
			t.Fatalf("%s", errCheck.Error())
		}
	}
}
