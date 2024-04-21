package api

import (
	receipt "go-receipt-processor/Receipt"
	receiptitem "go-receipt-processor/Receipt/ReceiptItem"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServerTestCase struct {
	unparsedReceipt     receipt.UnparsedReceipt
	expectedStatusCodeA int
	expectedStatusCodeB int
	expectedPoints      int
}

func TestServer(t *testing.T) {

	var testCases []ServerTestCase = []ServerTestCase{
		{
			unparsedReceipt: receipt.UnparsedReceipt{
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
			expectedStatusCodeA: 200,
			expectedStatusCodeB: 200,
			expectedPoints:      28,
		},
		{
			unparsedReceipt: receipt.UnparsedReceipt{
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
			expectedStatusCodeA: 200,
			expectedStatusCodeB: 200,
			expectedPoints:      109,
		},
	}
	url := "http://localhost:8080/receipts/"

	for _, testCase := range testCases {
		server := NewServer()
		w := httptest.NewRecorder()
		unparsedReceiptJson, _ := json.Marshal(testCase.unparsedReceipt)
		body := bytes.NewReader(unparsedReceiptJson)
		r, _ := http.NewRequest("POST", url+"process", body)
		server.Router.ServeHTTP(w, r)
		var id idResponse

		switch statusCode := w.Result().StatusCode; statusCode {
		case 200:
			json.NewDecoder(w.Body).Decode(&id)
		case testCase.expectedStatusCodeA:
			continue
		default:
			t.Fatalf("test server:\n    expected status code: \"%d\"\n    actual status code: \"%d\"\n", testCase.expectedStatusCodeA, statusCode)
		}

		w = httptest.NewRecorder()

		r, _ = http.NewRequest("GET", url+id.Id, body)
		server.Router.ServeHTTP(w, r)

		statusCode := w.Result().StatusCode
		var points pointsResponse
		switch {
		case statusCode == 200 && statusCode == testCase.expectedStatusCodeB:
			json.NewDecoder(w.Body).Decode(&points)
		case statusCode == testCase.expectedStatusCodeB:
			continue
		default:
			t.Fatalf("test server:\n    expected status code: \"%d\"\n    actual status code: \"%d\"\n", testCase.expectedStatusCodeA, statusCode)
		}
		if points.Points != int64(testCase.expectedPoints) {
			t.Fatalf("    expected: %d\n    got: %d\n", testCase.expectedPoints, points.Points)
		}
	}
}
