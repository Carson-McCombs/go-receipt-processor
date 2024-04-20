package date

import (
	utils "go-receipt-processor/TestingUtils"
	"testing"
)

func Test_ParseDate(t *testing.T) {
	var testCases []utils.CreationTestingData[string, Date] = []utils.CreationTestingData[string, Date]{
		{Argument: "", ExpectedResult: Date{}, ExpectedErr: ErrEmptyDateString},
		{Argument: "2000-59-92", ExpectedResult: Date{Year: 2000, Month: 59, Day: 92}, ExpectedErr: ErrInvalidDate},
		{Argument: "1999-12-92", ExpectedResult: Date{Year: 1999, Month: 12, Day: 92}, ExpectedErr: ErrInvalidDate},
		{Argument: "2004-02-29", ExpectedResult: Date{Year: 2004, Month: 02, Day: 29}, ExpectedErr: nil},
		{Argument: "ndaifsa", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "20-05-02-01", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "2005-02", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "2015-05-", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "2004-12-20-", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "2005-02-29a", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "2005e-02-29a", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
		{Argument: "06-23-2007", ExpectedResult: Date{}, ExpectedErr: ErrInvalidDateSyntax},
	}
	for _, testCase := range testCases {
		result, err := ParseDate(testCase.Argument, true)

		errCheck := testCase.CheckTestCase("parse date", result, err, false)
		if errCheck != nil {
			t.Fatalf("%s", errCheck.Error())
		}

	}
}
