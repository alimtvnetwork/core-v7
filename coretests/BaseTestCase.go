package coretests

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type BaseTestCase struct {
	Title                                    string // consider as header
	ArrangeInput, ActualInput, ExpectedInput interface{}
	HasError                                 bool
	IsValidateError                          bool
}

func (it *BaseTestCase) CaseTitle() string {
	return it.Title
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
