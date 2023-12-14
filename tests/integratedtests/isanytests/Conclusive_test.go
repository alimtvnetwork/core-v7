package isanytests

import (
	"testing"

	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/isany"
)

func Test_Conclusive_Verification(t *testing.T) {
	for caseIndex, testCase := range conclusiveTestCases {
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
			isEqual, isConclusive := isany.Conclusive(f, s)
			values := corecsv.AnyToValuesTypeString(f, s)
			conclusive := conditional.String(
				isConclusive,
				"Conclusive",
				"Inconclusive",
			)

			actualSlice.AppendFmt(
				conclusiveFormat,
				i,
				isEqual,
				conclusive,
				values,
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
