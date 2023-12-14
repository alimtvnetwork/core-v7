package reflectinternaltests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"

	"gitlab.com/auk-go/core/internal/reflectinternal"
)

func Test_PascalFuncName_Verification(t *testing.T) {
	for caseIndex, testCase := range pascalFuncNameTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.One)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(40)
		inArgString := input.First.(string)

		// Act
		actFuncPascalFuncName := reflectinternal.GetFunc.PascalFuncName
		result := actFuncPascalFuncName(inArgString)

		actualSlice.AppendFmt(
			"%d : %s -> %s",
			caseIndex,
			inArgString,
			result,
		)

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
