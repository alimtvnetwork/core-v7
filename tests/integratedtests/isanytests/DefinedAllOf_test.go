package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

func Test_DefinedAllOf_Verification(t *testing.T) {
	for caseIndex, testCase := range definedAllOfTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([][]interface{})
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringFmt,
				i,
				isany.DefinedAllOf(input...),
				corecsv.AnyToTypesCsvDefault(input...),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
