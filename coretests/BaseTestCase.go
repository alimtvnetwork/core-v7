package coretests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/issetter"
)

// BaseTestCase
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
//   - Will verify type using VerifyTypeOf
type BaseTestCase struct {
	Title           string         `json:",omitempty"` // consider as header
	ArrangeInput    interface{}    `json:",omitempty"` // preparing input, initial input
	ActualInput     interface{}    `json:",omitempty"` // (dynamically set) : must be set after running Act, using SetActual
	ExpectedInput   interface{}    `json:",omitempty"` // expectation set from the test
	Additional      interface{}    `json:",omitempty"` // additional input to do
	CustomFormat    string         `json:",omitempty"` // custom format for the test case
	VerifyTypeOf    *VerifyTypeOf  `json:",omitempty"` // Setting this creates the verify auto, verifies ArrangeInput, ActualInput, ExpectedInput type
	Parameters      *args.Holder   `json:",omitempty"` // If Act function / or any function requires more parameters it can be defined in the Holder.
	IsEnable        issetter.Value `json:",omitempty"` // Only false makes it disabled.
	HasError        bool           `json:",omitempty"`
	HasPanic        bool           `json:",omitempty"`
	IsValidateError bool           `json:",omitempty"`
}

func (it *BaseTestCase) CaseTitle() string {
	return it.Title
}

func (it *BaseTestCase) ArrangeTypeName() string {
	return reflectinternal.TypeName(it.ArrangeInput)
}

func (it *BaseTestCase) TypesValidationMustPasses(t *testing.T) {
	err := it.TypeValidationError()

	if err != nil {
		t.Error(
			"any one of the type validation failed",
			err.Error(),
		)
	}
}

func (it *BaseTestCase) IsTypeInvalidOrSkipVerify() bool {
	return it == nil ||
		it.VerifyTypeOf == nil ||
		it.VerifyTypeOf.IsInvalidOrSkipVerify()
}

func (it *BaseTestCase) HasParameters() bool {
	return it != nil &&
		it.Parameters != nil
}

func (it *BaseTestCase) IsInvalidParameters() bool {
	return it == nil || it.Parameters == nil
}

func (it *BaseTestCase) FirstParam() interface{} {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.First
}

func (it *BaseTestCase) SecondParam() interface{} {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Second
}

func (it *BaseTestCase) ThirdParam() interface{} {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Third
}

func (it *BaseTestCase) FourthParam() interface{} {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Fourth
}

func (it *BaseTestCase) FifthParam() interface{} {
	if it.IsInvalidParameters() {
		return nil
	}

	return it.Parameters.Fifth
}

func (it *BaseTestCase) HashmapParam() (hasMapItem bool, hashMap map[string]interface{}) {
	if it.IsInvalidParameters() {
		return false, map[string]interface{}{}
	}

	hashMap = it.Parameters.Hashmap

	return len(hashMap) > 0, hashMap
}

func (it *BaseTestCase) HasValidHashmapParam() bool {
	if it.IsInvalidParameters() {
		return false
	}

	hashMap := it.Parameters.Hashmap

	return len(hashMap) > 0
}

func (it *BaseTestCase) IsVerifyType() bool {
	return it != nil && !it.IsTypeInvalidOrSkipVerify()
}

// TypeValidationError
//
// must use SetActual to set the actual,
// what received from the act method,
// set it using SetActual
func (it *BaseTestCase) TypeValidationError() error {
	if it.IsTypeInvalidOrSkipVerify() {
		return nil
	}

	var sliceErr []string
	arrangeInputActualType := reflect.TypeOf(it.ArrangeInput)
	actualInputActualType := reflect.TypeOf(it.ActualInput)
	expectedInputActualType := reflect.TypeOf(it.ExpectedInput)
	verifyOf := it.VerifyTypeOf

	if reflectinternal.Is.Defined(it.ArrangeInput) && arrangeInputActualType != verifyOf.ArrangeInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Arrange Type Mismatch",
				verifyOf.ArrangeInput,
				arrangeInputActualType,
			),
		)
	}

	if reflectinternal.Is.Defined(it.ActualInput) && actualInputActualType != verifyOf.ActualInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Actual Type Mismatch",
				verifyOf.ActualInput,
				actualInputActualType,
			),
		)
	}

	if reflectinternal.Is.Defined(it.ExpectedInput) && expectedInputActualType != verifyOf.ExpectedInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Expected Type Mismatch",
				verifyOf.ExpectedInput,
				expectedInputActualType,
			),
		)
	}

	if len(sliceErr) > 0 {
		var newSlice []string

		newSlice = append(
			newSlice,
			it.Title,
		)
		sliceErr = append(
			newSlice,
			sliceErr...,
		)
	}

	return errcore.SliceToError(sliceErr)
}

// ArrangeString
//
//	returns ArrangeInput in string
//	format using constants.SprintValueFormat
func (it *BaseTestCase) ArrangeString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ArrangeInput,
	)
}

// Input returns ArrangeInput
func (it *BaseTestCase) Input() interface{} {
	return it.ArrangeInput
}

func (it *BaseTestCase) Expected() interface{} {
	return it.ExpectedInput
}

func (it *BaseTestCase) ExpectedString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ExpectedInput,
	)
}

func (it *BaseTestCase) Actual() interface{} {
	return it.ActualInput
}

func (it *BaseTestCase) ActualLines() []string {
	return GetAssert.ToStrings(it.ActualInput)
}

func (it *BaseTestCase) ExpectedLines() []string {
	return GetAssert.ToStrings(it.ExpectedInput)
}

func (it *BaseTestCase) ActualString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ActualInput,
	)
}

func (it *BaseTestCase) SetActual(actual interface{}) {
	it.ActualInput = actual
}

// String
//
//	returns a string format using GetAssertMessageUsingSimpleTestCaseWrapper
//	- https://prnt.sc/lxUV0eYk_qlg
func (it *BaseTestCase) String(caseIndex int) string {
	return GetAssert.SimpleTestCaseWrapper.String(
		caseIndex, it,
	)
}

func (it *BaseTestCase) LinesString(caseIndex int) string {
	return GetAssert.SimpleTestCaseWrapper.CaseLinesUsingDoubleQuoteLinesToString(
		caseIndex, it,
	)
}

func (it *BaseTestCase) IsDisabled() bool {
	return it.IsEnable.IsFalse()
}

func (it *BaseTestCase) IsSkipWithLog(caseIndex int) bool {
	if it.IsDisabled() {
		fmt.Printf(
			"Header : %s (%d), skipped: Disabled.",
			it.Title,
			caseIndex,
		)

		return true
	}

	return false
}

// ShouldBe
//
// Disabled testcases will not be executed.
func (it *BaseTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual interface{},
) {
	if it.IsEnable.IsFalse() {
		it.noPrintAssert(caseIndex, t, assert, actual)

		return
	}

	it.ShouldBeExplicit(
		true,
		caseIndex,
		t,
		it.Title,
		actual,
		assert,
		it.Expected(),
	)
}

func (it *BaseTestCase) noPrintAssert(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual interface{},
) {
	toTile := it.FormTitle(caseIndex)

	it.SetActual(actual)

	convey.Convey(
		toTile, t, func() {
			convey.SoMsg(
				toTile,
				actual,
				assert,
				it.ExpectedInput,
			)
		},
	)
}

func (it *BaseTestCase) FormTitle(caseIndex int) string {
	return fmt.Sprintf(
		skippedMsgFormat,
		caseIndex,
		it.Title,
	)
}

func (it *BaseTestCase) CustomTitle(caseIndex int, title string) string {
	return fmt.Sprintf(
		skippedMsgFormat,
		caseIndex,
		title,
	)
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
	if it.IsEnable.IsFalse() {
		it.noPrintAssert(caseIndex, t, assert, actual)

		return
	}

	it.SetActual(actual)
	headerTitle := it.CustomTitle(caseIndex, title)
	actualLines := GetAssert.ToStrings(actual)
	expectedLines := GetAssert.ToStrings(expected)
	compare := assert(actualLines, expectedLines)
	isFailed := compare != ""

	convey.Convey(
		headerTitle, t, func() {
			if isFailed {
				toString := it.LinesString(caseIndex)

				fmt.Println(toString)
			}

			convey.SoMsg(
				headerTitle,
				actualLines,
				assert,
				expectedLines,
			)
		},
	)

	if !isValidateType {
		return
	}

	it.TypeShouldMatch(t, caseIndex, title)
}

func (it *BaseTestCase) TypeShouldMatch(
	t *testing.T,
	caseIndex int,
	title string,
) {
	err := it.TypeValidationError()

	if err == nil {
		return
	}

	errHeader := fmt.Sprintf(
		"%d : %s - type verification failed",
		caseIndex,
		title,
	)

	var finalError error

	if err != nil {
		finalError = errors.New(errHeader + err.Error())
	}

	convey.Convey(
		errHeader, t, func() {
			convey.So(
				finalError,
				convey.ShouldBeNil,
			)
		},
	)
}

func (it *BaseTestCase) AsSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}

func (it *BaseTestCase) AsBaseTestCaseWrapper() BaseTestCaseWrapper {
	return it
}
