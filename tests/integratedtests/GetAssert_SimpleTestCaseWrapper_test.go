package integratedtests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_SimpleTestCaseWrapper_String_Verification(t *testing.T) {
	for caseIndex, testCase := range stringTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert.SimpleTestCaseWrapper
		actFunc := asserter.String
		caseV1 := testCase.ArrangeInput.(coretestcases.CaseV1)
		simplerWrapper := caseV1.AsSimpleTestCaseWrapper()

		// Act
		output := actFunc(
			caseIndex,
			simplerWrapper,
		)

		actualSlice.Adds(strings.Split(output, "\n")...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_SimpleTestCaseWrapper_Lines_Verification(t *testing.T) {
	for caseIndex, testCase := range linesTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert.SimpleTestCaseWrapper
		caseV1 := testCase.ArrangeInput.(coretestcases.CaseV1)
		simplerWrapper := caseV1.AsSimpleTestCaseWrapper()
		prefixWithSpaceFunc := coretests.GetAssert.ToStringsWithSpace
		actFunc := asserter.Lines

		// Act
		arrange, expected := actFunc(
			simplerWrapper,
		)

		actualSlice.Add("Title: " + simplerWrapper.CaseTitle())
		actualSlice.Add("Arrange Lines:")
		actualSlice.Adds(prefixWithSpaceFunc(4, arrange)...)
		actualSlice.Add("Expected Lines:")
		actualSlice.Adds(prefixWithSpaceFunc(4, expected)...)

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
