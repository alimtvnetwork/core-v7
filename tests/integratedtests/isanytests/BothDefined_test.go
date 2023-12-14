package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

func Test_BothDefined_Verification(t *testing.T) {
	for caseIndex, testCase := range bothDefinedTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.Two)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second

			actualSlice.AppendFmt(
				defaultCaseIndexBoolStringFmt,
				i,
				isany.DefinedBoth(f, s),
				corecsv.AnyToTypesCsvDefault(f, s),
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
