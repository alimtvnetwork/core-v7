package integratedtests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_GetAssert_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range quickTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		quickFunc := asserter.Quick

		// Act
		output := quickFunc(
			input.When(),                        // when
			input.Actual(),                      // actual
			input.Expect(),                      // expected
			input.GetAsIntDefault("counter", 0), // counter
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
