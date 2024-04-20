package time

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrEmptyTimeString = errors.New("attempting to parse time variable from empty string")

	ErrParsingTime   = errors.New("parsing time")
	ErrParsingHour   = errors.New("parsing hour")
	ErrParsingMinute = errors.New("parsing minute")

	ErrInvalidTimeSyntax = errors.New("invalid time syntax")

	ErrInvalidTime   = errors.New("invalid time")
	ErrInvalidHour   = errors.New("invalid hour")
	ErrInvalidMinute = errors.New("invalid minute")
)

type Time struct { //assuming 24 hour time
	Hour   uint8
	Minute uint8
}

func (t *Time) Equals(other Time) bool {
	return (t.Hour == other.Hour && t.Minute == other.Minute)
}

func (t Time) IsValid() error {
	isHourValid := t.Hour < 24
	isMinuteValid := t.Minute < 60

	if isHourValid && isMinuteValid {
		return nil
	} else if isHourValid && !isMinuteValid {
		return fmt.Errorf("%w invalid minute values provided: %+v ( valid values range inclusively from 0 to 59 )", ErrInvalidTime, t)
	} else if !isHourValid && isMinuteValid {
		return fmt.Errorf("%w invalid hour values provided: %+v ( valid values range inclusively from 0 to 23 )", ErrInvalidTime, t)
	} else {
		return fmt.Errorf("%w invalid hour and minute values provided: %+v ( valid hour values range inclusively from 0 to 23 and valid minute values range inclusively from 0 to 59 )", ErrInvalidTime, t)
	}
}

func inspectSyntaxAndSplitTimeString(timeString string) (string, string, error) {
	timeWithNoSpaces := strings.ReplaceAll(timeString, " ", "")
	if len(timeWithNoSpaces) > 5 { // format should be 5 characters, HH:MM
		return "", "", fmt.Errorf("%w given \"%s\", too many characters ( valid format is HH:MM )", ErrInvalidTimeSyntax, timeWithNoSpaces)
	}
	if strings.Count(timeWithNoSpaces, ":") != 1 { // format should only have 1 colon
		return "", "", fmt.Errorf("%w given \"%s\", there should be exactly 1 colon ( valid format is HH:MM )", ErrInvalidTimeSyntax, timeWithNoSpaces)
	}
	splitStrings := strings.Split(timeWithNoSpaces, ":")

	for _, splitStr := range splitStrings {
		isNumeric := regexp.MustCompile(`^\d+$`).MatchString(splitStr)
		if !isNumeric {
			return "", "", fmt.Errorf("%w given \"%s\", only postive numeric characters are allowed ( valid format is HH:MM )", ErrInvalidTimeSyntax, timeWithNoSpaces)
		}
	}
	hourLength := len(splitStrings[0])
	minuteLength := len(splitStrings[1])
	switch {
	case hourLength == 0 || hourLength > 2: // allows for hour to only be 1 digit long without throwing an error
		return "", "", fmt.Errorf("%w given \"%s\", invalid number of digits for hour ( valid format is HH:MM )", ErrInvalidTimeSyntax, timeWithNoSpaces)
	case minuteLength != 2:
		return "", "", fmt.Errorf("%w given \"%s\", invalid number of digits for minute ( valid format is HH:MM )", ErrInvalidTimeSyntax, timeWithNoSpaces)
	default:
		return splitStrings[0], splitStrings[1], nil
	}

}

func ParseTime(timeString string, validateResults bool) (Time, error) {
	time := Time{}
	if strings.Trim(timeString, " ") == "" {
		return time, ErrEmptyTimeString
	}

	hourString, minuteString, syntaxErr := inspectSyntaxAndSplitTimeString(timeString)
	if syntaxErr != nil {
		return time, syntaxErr
	}

	var parsingHourErr error = nil
	hourValue, err := strconv.ParseUint(hourString, 10, 8)
	if err == nil {
		time.Hour = uint8(hourValue)
	} else {
		parsingHourErr = fmt.Errorf("%w given %s ... %s", ErrParsingHour, hourString, err.Error())
	}

	var parsingMinuteErr error = nil
	minuteValue, err := strconv.ParseUint(minuteString, 10, 8)
	if err == nil {
		time.Minute = uint8(minuteValue)
	} else {
		parsingMinuteErr = fmt.Errorf("%w given %s ... %s", ErrParsingMinute, minuteString, err.Error())
	}

	if parsingHourErr != nil || parsingMinuteErr != nil {
		return time, fmt.Errorf("%w from %s ... %s", ErrParsingTime, timeString, errors.Join(parsingHourErr, parsingMinuteErr).Error())
	}
	if validateResults {
		return time, time.IsValid()
	}
	return time, nil
}
