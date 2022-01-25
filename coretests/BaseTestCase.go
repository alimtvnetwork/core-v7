package coretests

import (
	"fmt"
	"reflect"
	"testing"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

type BaseTestCase struct {
	Title                                                           string // consider as header
	ArrangeInput, ActualInput, ExpectedInput                        interface{}
	ArrangeExpectedType, ActualExpectedType, ExpectedTypeOfExpected reflect.Type
	HasError                                                        bool
	IsValidateError                                                 bool
}

func (it *BaseTestCase) CaseTitle() string {
	return it.Title
}

func (it *BaseTestCase) TypesValidationMustPasses(t *testing.T) {
	err := it.TypeValidationError()

	if err != nil {
		t.Error(
			"any one of the type validation failed",
			err.Error())
	}
}

func (it *BaseTestCase) TypeValidationError() error {
	var sliceErr []string
	arrangeInputActualType := reflect.TypeOf(it.ArrangeExpectedType)
	actualInputActualType := reflect.TypeOf(it.ActualInput)
	expectedInputActualType := reflect.TypeOf(it.ExpectedInput)

	if arrangeInputActualType != it.ArrangeExpectedType {
		sliceErr = append(
			sliceErr,
			errcore.Expecting(
				"Arrange Type Mismatch",
				it.ArrangeExpectedType,
				arrangeInputActualType))
	}

	if actualInputActualType != it.ActualExpectedType {
		sliceErr = append(
			sliceErr,
			errcore.Expecting(
				"Actual Type Mismatch",
				it.ActualExpectedType,
				actualInputActualType))
	}

	if expectedInputActualType != it.ExpectedTypeOfExpected {
		sliceErr = append(
			sliceErr,
			errcore.Expecting(
				"Expected Type Mismatch",
				it.ExpectedTypeOfExpected,
				expectedInputActualType))
	}

	return errcore.SliceToError(sliceErr)
}

func (it *BaseTestCase) ArrangeString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ArrangeInput)
}

func (it *BaseTestCase) Input() interface{} {
	return it.ArrangeInput
}

func (it *BaseTestCase) Expected() interface{} {
	return it.ExpectedInput
}

func (it *BaseTestCase) ExpectedString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ExpectedInput)
}

func (it *BaseTestCase) Actual() interface{} {
	return it.ActualInput
}

func (it *BaseTestCase) ActualString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ActualInput)
}

func (it *BaseTestCase) SetActual(actual interface{}) {
	it.ActualInput = actual
}

func (it *BaseTestCase) String(caseIndex int) string {
	return GetAssertMessageUsingSimpleTestCaseWrapper(
		caseIndex, it)
}

func (it *BaseTestCase) AsSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}
