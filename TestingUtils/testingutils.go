package testingutils

import (
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

type CreationTestingData[A any, T any] struct {
	Argument       A
	ExpectedResult T
	ExpectedErr    error
}

func errorsEqual(errA error, errB error) bool {
	if (errA == nil) && (errB == nil) {
		return true
	} else if (errA == nil) != (errB == nil) {
		return false
	}
	return errors.Is(errA, errB)
}

func (testCase *CreationTestingData[A, T]) CheckTestCase(functionName string, actualResult T, actualErr error, debugAll bool) error {
	equalResult := cmp.Equal(testCase.ExpectedResult, actualResult)
	equalError := errorsEqual(actualErr, testCase.ExpectedErr)
	startingString := fmt.Sprintf("\n%s:\n    argument:\n        %+v\n", functionName, testCase.Argument)
	expectedResultString := fmt.Sprintf("    expected result:\n        %+v\n", testCase.ExpectedResult)
	actualResultString := fmt.Sprintf("    actual result:\n        %+v\n", actualResult)

	expectedErrString := fmt.Sprintf("    expected error:\n        %v\n", testCase.ExpectedErr)
	actualErrString := fmt.Sprintf("    actual error:\n        %v\n", actualErr)

	if !equalResult || !equalError {
		if debugAll {
			return fmt.Errorf(startingString + expectedResultString + expectedErrString + actualResultString + actualErrString)
		} else if !equalResult && !equalError {
			return fmt.Errorf(startingString + expectedErrString + actualErrString)
		} else if !equalResult {
			return fmt.Errorf(startingString + expectedResultString + actualResultString)
		} else if !equalError {
			return fmt.Errorf(startingString + expectedErrString + actualErrString)
		}
	}

	return nil
}
