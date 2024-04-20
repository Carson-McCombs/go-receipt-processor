package receiptitem

import (
	utils "go-receipt-processor/TestingUtils"

	"testing"
)

func Test_ParseReceiptItem(t *testing.T) {
	var testCases []utils.CreationTestingData[UnparsedReceiptItem, ReceiptItem] = []utils.CreationTestingData[UnparsedReceiptItem, ReceiptItem]{
		{Argument: UnparsedReceiptItem{ShortDescription: "", Price: ""}, ExpectedResult: DefaultReceiptItem(), ExpectedErr: ErrEmptyPriceString},
		{Argument: UnparsedReceiptItem{ShortDescription: "Gatorade", Price: "2.79"}, ExpectedResult: ReceiptItem{ShortDescription: "Gatorade", Price: 2.79}, ExpectedErr: nil},
		{Argument: UnparsedReceiptItem{ShortDescription: "Powerade", Price: "1.50"}, ExpectedResult: ReceiptItem{ShortDescription: "Powerade", Price: 1.50}, ExpectedErr: nil},
		{Argument: UnparsedReceiptItem{ShortDescription: "Gatorade Refund", Price: "-2.79"}, ExpectedResult: ReceiptItem{ShortDescription: "Gatorade Refund", Price: -2.79}, ExpectedErr: nil},
	}
	for _, testCase := range testCases {
		result, err := ParseReceiptItem(testCase.Argument)
		errCheck := testCase.CheckTestCase("parse receipt item", result, err, false)
		if errCheck != nil {
			t.Fatalf("%s", errCheck.Error())
		}

	}
}
