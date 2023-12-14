package integratedtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_GetAssert_SortedArray_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedArrayTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		sortedArray := asserter.SortedArray

		// Act
		outputs := sortedArray(
			input["isPrint"].(bool),
			input["isSort"].(bool),
			input["message"].(string),
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
