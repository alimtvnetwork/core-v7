package coretests

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests/args"
)

// SimpleTestCase
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
type SimpleTestCase struct {
	Title         string      `json:",omitempty"` // consider as header
	ArrangeInput  interface{} `json:",omitempty"` // preparing input, initial input
	ActualInput   interface{} `json:",omitempty"` // (dynamically set) : must be set after running Act, using SetActual
	ExpectedInput interface{} `json:",omitempty"` // expectation set from the test
	Params        args.Map
}

func (it SimpleTestCase) CaseTitle() string {
	return it.Title
}

// ArrangeString
//
//	returns ArrangeInput in string
//	format using constants.SprintValueFormat
func (it SimpleTestCase) ArrangeString() string {
	return GetAssert.ToString(it.ArrangeInput)
}

// Input returns ArrangeInput
func (it SimpleTestCase) Input() interface{} {
	return it.ArrangeInput
}

func (it SimpleTestCase) Expected() interface{} {
	return it.ExpectedInput
}

func (it SimpleTestCase) ExpectedString() string {
	return GetAssert.ToString(it.ExpectedInput)
}

func (it SimpleTestCase) Actual() interface{} {
	return it.ActualInput
}

func (it SimpleTestCase) ActualString() string {
	return GetAssert.ToString(it.ActualInput)
}

func (it SimpleTestCase) SetActual(actual interface{}) {
	it.ActualInput = actual
}

// String
//
//	returns a string format using GetAssertMessageUsingSimpleTestCaseWrapper
//	- https://prnt.sc/lxUV0eYk_qlg
func (it SimpleTestCase) String(caseIndex int) string {
	return GetAssert.
		SimpleTestCaseWrapper.
		String(
			caseIndex,
			it,
		)
}

func (it SimpleTestCase) LinesString(caseIndex int) string {
	return GetAssert.
		SimpleTestCaseWrapper.
		CaseLinesUsingDoubleQuoteLinesToString(
			caseIndex,
			it,
		)
}

func (it SimpleTestCase) noPrintAssert(
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

func (it SimpleTestCase) FormTitle(caseIndex int) string {
	return fmt.Sprintf(
		skippedMsgFormat,
		caseIndex,
		it.Title,
	)
}

func (it SimpleTestCase) CustomTitle(caseIndex int, title string) string {
	return fmt.Sprintf(
		skippedMsgFormat,
		caseIndex,
		title,
	)
}

func (it SimpleTestCase) ShouldBeEqual(
	caseIndex int,
	t *testing.T,
	actual interface{},
) {
	it.SetActual(actual)

	it.ShouldBeExplicit(
		caseIndex,
		t,
		it.Title,
		actual,
		convey.ShouldEqual,
		it.Expected(),
	)
}

func (it SimpleTestCase) ShouldHaveNoError(
	caseIndex int,
	t *testing.T,
	err error,
) {
	it.SetActual(err)

	it.ShouldBeExplicit(
		caseIndex,
		t,
		it.Title,
		err,
		convey.ShouldBeNil,
		it.Expected(),
	)
}

func (it SimpleTestCase) ShouldContains(
	caseIndex int,
	t *testing.T,
	actual interface{},
) {
	it.SetActual(actual)

	it.ShouldBeExplicit(
		caseIndex,
		t,
		it.Title,
		actual,
		convey.ShouldContain,
		it.Expected(),
	)
}

func (it SimpleTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual interface{},
) {
	it.SetActual(actual)

	it.ShouldBeExplicit(
		caseIndex,
		t,
		it.Title,
		actual,
		assert,
		it.Expected(),
	)
}

func (it SimpleTestCase) ShouldBeExplicit(
	caseIndex int,
	t *testing.T,
	title string,
	actual interface{},
	assert convey.Assertion,
	expected interface{},
) {
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
}

func (it SimpleTestCase) AsSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}
