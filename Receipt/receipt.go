package receipt

import (
	date "go-receipt-processor/Date"
	receiptitem "go-receipt-processor/Receipt/ReceiptItem"
	time "go-receipt-processor/Time"
	"errors"
	"fmt"
	"strconv"
)

var (

	//sentinel errors
	ErrParsingReceipt     error = errors.New("parsing receipt")
	ErrParsingReceiptItem error = errors.New("parsing receipt")
	ErrParsingTotal       error = errors.New("parsing receipt total")
	ErrInvalidTotal       error = errors.New("invalid receipt total")
	ErrInvalidReceipt     error = errors.New("invalid receipt")
	// all invalid syntax provided will be considered a parsing error
)

type UnparsedReceipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []receiptitem.UnparsedReceiptItem
	Total        string
}

type Receipt struct {
	Id           string
	Retailer     string
	PurchaseDate date.Date
	PurchaseTime time.Time
	Items        []receiptitem.ReceiptItem
	Total        float64
}

func (r Receipt) isValid() error {

	dateValidation := r.PurchaseDate.IsValid()
	timeValidation := r.PurchaseTime.IsValid()

	calculatedTotal := calculateSumOfReceiptItems(r.Items)
	diff := r.Total - calculatedTotal
	var totalValidation error = nil
	if diff >= 0.001 || diff <= -0.001 {
		totalValidation = fmt.Errorf("%w ... calculated total, %f, does not match parsed total, %f", ErrInvalidTotal, calculatedTotal, r.Total)
	}

	joinedValidationErr := errors.Join(dateValidation, timeValidation, totalValidation)
	if joinedValidationErr != nil {
		return fmt.Errorf("%w given %+v ... %v", ErrInvalidReceipt, r, joinedValidationErr)
	}
	return nil
}

func calculateSumOfReceiptItems(items []receiptitem.ReceiptItem) float64 {
	var total float64 = 0.0
	for _, item := range items {
		total += item.Price
	}
	return total
}

func ParseReceipt(id string, unparsedReceipt UnparsedReceipt, validateResults bool) (Receipt, error) {
	receipt := Receipt{
		Id:       id,
		Retailer: unparsedReceipt.Retailer,
	}
	purchaseDate, purchaseDateErr := date.ParseDate(unparsedReceipt.PurchaseDate, false)
	receipt.PurchaseDate = purchaseDate

	purchaseTime, purchaseTimeErr := time.ParseTime(unparsedReceipt.PurchaseTime, false)
	receipt.PurchaseTime = purchaseTime

	receiptItems, receiptItemsErr := ParseAllReceiptItems(unparsedReceipt.Items)
	receipt.Items = receiptItems

	var parseTotalErr error = nil
	total, err := strconv.ParseFloat(unparsedReceipt.Total, 64)
	if err != nil {
		parseTotalErr = fmt.Errorf("%w given \"%s\" ... %s", ErrParsingTotal, unparsedReceipt.Total, err.Error())
	} else {
		receipt.Total = total
	}

	joinedErrs := errors.Join(purchaseDateErr, purchaseTimeErr, receiptItemsErr, parseTotalErr)
	if joinedErrs != nil {
		return receipt, fmt.Errorf("( %s ) %w from %+v ... %s", id, ErrParsingReceipt, unparsedReceipt, joinedErrs.Error())
	}
	if validateResults {
		return receipt, receipt.isValid()
	}
	return receipt, nil
}

func ParseAllReceiptItems(unparsedItems []receiptitem.UnparsedReceiptItem) ([]receiptitem.ReceiptItem, error) {
	parsedItems := []receiptitem.ReceiptItem{}
	errs := []error{}
	for _, item := range unparsedItems {
		parsedItem, err := receiptitem.ParseReceiptItem(item)
		if err == nil {
			parsedItems = append(parsedItems, parsedItem)
		} else {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return parsedItems, fmt.Errorf("errors found when parsing receipt items ... %w", errors.Join(errs...))
	}
	return parsedItems, nil
}
