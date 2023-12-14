package integratedtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_GetAssert_ToStrings_Verification(t *testing.T) {
	for caseIndex, testCase := range toStringsTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		actFunc := asserter.ToStrings

		// Act
		outputs := actFunc(
			input["any"],
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
