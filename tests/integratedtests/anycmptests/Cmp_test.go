package anycmptests

import (
	"testing"

	"gitlab.com/auk-go/core/anycmp"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_Cmp_Verification(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// Arrange
		inputs := testCase.ArrangeInput.([]args.Two)
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			actualSlice.AppendFmt(
				"%d : %s (%T, %T)",
				i,
				anycmp.Cmp(
					parameter.First,
					parameter.Second,
				).
					String(),
				parameter.First,
				parameter.Second,
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
