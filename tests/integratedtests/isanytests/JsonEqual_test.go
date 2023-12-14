package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

func Test_JsonEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range jsonEqualTestCases {
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
				defaultCaseIndexBoolStringStringFmt,
				i,
				isany.JsonEqual(f, s),
				corejson.Serialize.ToString(f),
				corejson.Serialize.ToString(s),
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
