package receiptitem

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrEmptyPriceString   error = errors.New("attempting to parse item price variable from empty string")
	ErrParsingReceiptItem error = errors.New("parsing receipt item")
	ErrParsingPrice       error = errors.New("parsing receipt item price")
	//ErrInvalidReceiptItem error = errors.New("invalid receipt item")
)

type UnparsedReceiptItem struct {
	ShortDescription string
	Price            string
}

type ReceiptItem struct {
	ShortDescription string
	Price            float64
}

func DefaultReceiptItem() ReceiptItem {
	return ReceiptItem{
		ShortDescription: "",
		Price:            0,
	}
}

func ParseReceiptItem(unparsedItem UnparsedReceiptItem) (ReceiptItem, error) {
	receiptItem := DefaultReceiptItem()
	if strings.ReplaceAll(unparsedItem.Price, " ", "") == "" {
		return receiptItem, ErrEmptyPriceString
	}
	receiptItem.ShortDescription = unparsedItem.ShortDescription

	var parsingPriceErr error = nil
	priceValue, err := strconv.ParseFloat(unparsedItem.Price, 64)
	if err == nil {
		receiptItem.Price = priceValue
	} else {
		parsingPriceErr = fmt.Errorf("%w from \"%s\" ... %s", ErrParsingPrice, unparsedItem.Price, err.Error())
		return receiptItem, fmt.Errorf("%w from %+v ... %s", ErrParsingReceiptItem, unparsedItem, parsingPriceErr.Error())
	}

	return receiptItem, nil
}
