package argstests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_FuncWrap_Creation_Verification(t *testing.T) {
	for caseIndex, testCase := range funWrapCreationTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.ThreeFunc)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(30)
		toStringsConv := coretests.GetAssert.ToStrings

		// Act
		output, err := input.InvokeWithValidArgs()

		actualSlice.Adds(toStringsConv(output)...)

		if err != nil {
			errLines := coretests.
				GetAssert.
				ErrorToLinesWithSpacesDefault(err)

			actualSlice.Add(
				"error : ",
			)

			actualSlice.Adds(
				errLines...,
			)
		}

		finalActLines := actualSlice.Strings()
		actualSlice.Dispose()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
