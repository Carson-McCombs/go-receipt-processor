package date

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	monthNames        = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	validDaysPerMonth = []uint8{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	ErrEmptyDateString = errors.New("attempting to parse date variable from empty string")
	ErrParsingDate     = errors.New("parsing date")
	ErrParsingYear     = errors.New("parsing year")
	ErrParsingMonth    = errors.New("parsing month")
	ErrParsingDay      = errors.New("parsing day")

	ErrInvalidDateSyntax = errors.New("invalid date syntax")

	ErrInvalidDate = errors.New("invalid date")
)

type Date struct {
	Year  uint16
	Month uint8
	Day   uint8
}

func (d *Date) Equals(other Date) bool {
	return (d.Year == other.Year && d.Month == other.Month && d.Day == other.Day)
}

func GetMonthName(month uint8) string {
	if month < 1 || month > 12 {
		return "invalid month"
	}
	return monthNames[month-1]
}

func GetValidDaysInMonth(year uint16, month uint8) uint8 {
	if month < 1 || month > 12 {
		return 0 //month outside of bounds
	}
	if year%4 == 0 && month == 2 {
		return 29 // February has 29 days on a Leap Year
	}
	return validDaysPerMonth[month-1]
}

// Checks if the date is valid, including the bounds of the month and the day, but not the bounds of the year.
func (d Date) IsValid() error {

	if d.Month < 1 || d.Month > 12 {
		return fmt.Errorf("%w provided month value: %d ( valid values range inclusively from 1 to 12 )", ErrInvalidDate, d.Month)
	}
	validDaysInMonth := GetValidDaysInMonth(d.Year, d.Month)
	if d.Day < 1 || d.Day > validDaysInMonth {
		return fmt.Errorf("%w provided day value: %+v ( given the month and year, valid values range inclusively from 1 to %d )", ErrInvalidDate, d, validDaysInMonth)
	}

	return nil
}

func inspectSyntaxAndSplitDateString(dateString string) (string, string, string, error) {
	dateWithNoSpaces := strings.ReplaceAll(dateString, " ", "")
	if len(dateWithNoSpaces) != 10 { // format should be 10 characters, YYYY-MM-DD
		return "", "", "", fmt.Errorf("%w given \"%s\", too many characters ( valid format is YYYY-MM-DD )", ErrInvalidDateSyntax, dateString)
	}
	if strings.Count(dateWithNoSpaces, "-") != 2 { // format should only have 2 dashes
		return "", "", "", fmt.Errorf("%w given \"%s\", invalid number of dashes ( valid format is YYYY-MM-DD )", ErrInvalidDateSyntax, dateString)
	}
	splitStrings := strings.Split(dateWithNoSpaces, "-")

	for i := 0; i < len(splitStrings); i++ {
		isNumeric := regexp.MustCompile(`\d`).MatchString(splitStrings[i])
		if !isNumeric {
			return "", "", "", fmt.Errorf("%w given \"%s\", only numeric characters are allowed ( valid format is YYYY-MM-DD )", ErrInvalidDateSyntax, dateString)
		}
	}
	switch {
	case len(splitStrings[0]) != 4:
		return "", "", "", fmt.Errorf("%w given \"%s\", invalid number of digits for year ( valid format is YYYY-MM-DD )", ErrInvalidDateSyntax, dateString)
	case len(splitStrings[1]) != 2:
		return "", "", "", fmt.Errorf("%w given \"%s\", invalid number of digits for month ( valid format is YYYY-MM-DD )", ErrInvalidDateSyntax, dateString)
	case len(splitStrings[2]) != 2:
		return "", "", "", fmt.Errorf("%w given \"%s\", invalid number of digits for day ( valid format is YYYY-MM-DD )", ErrInvalidDateSyntax, dateString)
	default:
		return splitStrings[0], splitStrings[1], splitStrings[2], nil

	}

}

func ParseDate(dateString string, validateResults bool) (Date, error) {
	date := Date{}
	if strings.Trim(dateString, " ") == "" {
		return date, ErrEmptyDateString
	}

	yearString, monthString, dayString, syntaxErr := inspectSyntaxAndSplitDateString(dateString)
	if syntaxErr != nil {
		return date, syntaxErr
	}

	var parsingYearErr error = nil
	yearValue, err := strconv.ParseUint(yearString, 10, 16)
	if err == nil {
		date.Year = uint16(yearValue)
	} else {
		parsingYearErr = fmt.Errorf("%w ... given %s ... %s", ErrParsingYear, yearString, err.Error())
	}

	var parsingMonthErr error = nil
	monthValue, err := strconv.ParseUint(monthString, 10, 8)
	if err == nil {
		date.Month = uint8(monthValue)
	} else {
		parsingMonthErr = fmt.Errorf("%w given %s ... %s", ErrParsingMonth, monthString, err.Error())
	}

	var parsingDayErr error = nil
	dayValue, err := strconv.ParseUint(dayString, 10, 8)

	if err == nil {
		date.Day = uint8(dayValue)
	} else {
		parsingDayErr = fmt.Errorf("%w given %s ... %s", ErrParsingDay, dayString, err.Error())
	}

	if parsingYearErr != nil || parsingMonthErr != nil || parsingDayErr != nil {
		return date, fmt.Errorf("%w given %s ... %s", ErrParsingDate, dateString, errors.Join(parsingYearErr, parsingMonthErr, parsingDayErr).Error())
	}
	if validateResults {
		return date, date.IsValid()
	}
	return date, nil

}
