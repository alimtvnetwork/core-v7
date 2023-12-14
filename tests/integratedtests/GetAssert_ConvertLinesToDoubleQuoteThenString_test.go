package integratedtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
)

func Test_ConvertLinesToDoubleQuoteThenString_Verification(t *testing.T) {
	for caseIndex, testCase := range convertLinesToDoubleQuoteThenStringTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert
		params := testCase.Parameters
		actFunc := asserter.ConvertLinesToDoubleQuoteThenString

		// Act
		outputLine := actFunc(
			params.First.(int),
			params.Second.([]string),
		)

		actualSlice.Add(outputLine)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_AnyToStringDoubleQuoteLine_Verification(t *testing.T) {
	for caseIndex, testCase := range convertLinesToDoubleQuoteThenStringTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert
		params := testCase.Parameters
		actFunc := asserter.AnyToStringDoubleQuoteLine

		// Act
		outputLine := actFunc(
			params.First.(int),
			params.Second.([]string),
		)

		actualSlice.Add(outputLine)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
