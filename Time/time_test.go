package time

import (
	utils "GoReceiptProcessor/TestingUtils"
	"testing"
)

func Test_ParseTime(t *testing.T) {
	var testCases []utils.CreationTestingData[string, Time] = []utils.CreationTestingData[string, Time]{
		{Argument: "01:01", ExpectedResult: Time{Hour: 1, Minute: 1}, ExpectedErr: nil},
		{Argument: "23:59", ExpectedResult: Time{Hour: 23, Minute: 59}, ExpectedErr: nil},
		{Argument: "00:00", ExpectedResult: Time{Hour: 0, Minute: 0}, ExpectedErr: nil},

		//Tests Dates that don't have a leading zero to make up a two digit hour ( I considered this acceptable, although two digits is still required for the minute)
		{Argument: "5:51", ExpectedResult: Time{Hour: 5, Minute: 51}, ExpectedErr: nil},
		{Argument: "3:51", ExpectedResult: Time{Hour: 3, Minute: 51}, ExpectedErr: nil},
		//Tests Dates that are "Invalid" or *could* be parsed correctly but lie outside of the acceptable range
		{Argument: "13:", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax},                 //Tests if the minute is 0 digits long, that an error would correctly be thrown
		{Argument: "13:1", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax},                //Tests if the minute is 1 digit long, that an error would correctly be thrown
		{Argument: "13:001", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax},              //Tests if the minute is more than 2 digits long, than an error would correctly be thrown
		{Argument: "24:00", ExpectedResult: Time{Hour: 24, Minute: 0}, ExpectedErr: ErrInvalidTime},  //Tests hour range ( 0 <= hour <= 23)
		{Argument: "37:12", ExpectedResult: Time{Hour: 37, Minute: 12}, ExpectedErr: ErrInvalidTime}, //Tests hour range ( 0 <= hour <= 23)
		{Argument: "12:60", ExpectedResult: Time{Hour: 12, Minute: 60}, ExpectedErr: ErrInvalidTime}, //Tests minute range ( 0 <= minute < 60)
		{Argument: "18:61", ExpectedResult: Time{Hour: 18, Minute: 61}, ExpectedErr: ErrInvalidTime}, //Tests minute range ( 0 <= minute < 60)
		{Argument: "13:99", ExpectedResult: Time{Hour: 13, Minute: 99}, ExpectedErr: ErrInvalidTime}, //Tests minute range ( 0 <= hour <= 60)
		//Tests inputs that have an invalid syntax will correctly throw an error
		{Argument: "asdaf", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax}, //Tests if alphabetic characters are given in input
		{Argument: "a2s55", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax}, //Tests if alpha-numeric characers are given in input
		{Argument: "hr:mn", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax}, //Tests if given the correct format and string length, if alphabetic characters were provided instead of numeric ones, if an error would correctly be thrown
		{Argument: "h1:2n", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax}, //Tests if given correct digit count, if alpha-numeric characters in input would correctly throw an error
		{Argument: "1234", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax},  //Tests if an input with a missing colon would correctly throw an error
		{Argument: "12345", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax}, //Tests if an input with the correct total lenght and only numbers would correctly still throw an error
		{Argument: "-5:51", ExpectedResult: Time{}, ExpectedErr: ErrInvalidTimeSyntax}, //Tests if negative times would correctly throw an error
	}
	for _, testCase := range testCases {
		result, err := ParseTime(testCase.Argument, true)
		errCheck := testCase.CheckTestCase("parse time", result, err, false)
		if errCheck != nil {
			t.Fatalf("%s", errCheck.Error())
		}
	}
}
