package integratedtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
)

func Test_AnyToDoubleQuoteLines_Verification(t *testing.T) {
	for caseIndex, testCase := range anyToDoubleQuoteLinesTestCases {
		// Arrange
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(20)
		asserter := coretests.GetAssert
		params := testCase.Parameters
		actFunc := asserter.AnyToDoubleQuoteLines

		// Act
		outputs := actFunc(
			params.First.(int),
			params.Second,
		)

		actualSlice.Adds(outputs...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
