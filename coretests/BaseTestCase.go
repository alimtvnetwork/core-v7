package coretests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
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
	arrangeInputActualType := reflect.TypeOf(it.ArrangeInput)
	actualInputActualType := reflect.TypeOf(it.ActualInput)
	expectedInputActualType := reflect.TypeOf(it.ExpectedInput)

	if reflectinternal.IsNotNull(it.ArrangeInput) && arrangeInputActualType != it.ArrangeExpectedType {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Arrange Type Mismatch",
				it.ArrangeExpectedType,
				arrangeInputActualType))
	}

	if reflectinternal.IsNotNull(it.ActualInput) && actualInputActualType != it.ActualExpectedType {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Actual Type Mismatch",
				it.ActualExpectedType,
				actualInputActualType))
	}

	if reflectinternal.IsNotNull(it.ExpectedInput) && expectedInputActualType != it.ExpectedTypeOfExpected {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
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

func (it *BaseTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual interface{},
) {
	it.ShouldBeExplicit(
		true,
		caseIndex,
		t,
		it.Title,
		actual,
		assert,
		it.Expected())
}

func (it *BaseTestCase) ShouldBeExplicit(
	isValidateType bool,
	caseIndex int,
	t *testing.T,
	title string,
	actual interface{},
	assert convey.Assertion,
	expected interface{},
) {
	it.SetActual(actual)

	convey.Convey(title, t, func() {
		convey.SoMsg(it.String(caseIndex), actual, assert, expected)
	})

	if !isValidateType {
		return
	}

	err := it.TypeValidationError()
	errHeader := fmt.Sprintf(
		"case %d : test case type validation must passes",
		caseIndex)

	if err != nil {
		err = errors.New(errHeader + err.Error() + ", case title : " + title)
	}

	convey.Convey(errHeader, t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func (it *BaseTestCase) AsSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}
