package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_SquareWrapIf_Wraps_All_Without_Existing_Condition_Checking_Can_Have_DuplicateSquareBrackets(
	t *testing.T,
) {
	for caseIndex, testCase := range squareBracketWrapTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.SquareWrapIf(
					true,
					input,
				),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_SquareWrapIf_Disabled_Wraps_Nothing(
	t *testing.T,
) {
	for caseIndex, testCase := range squareBracketWrapDisabledTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.SquareWrapIf(
					false,
					input,
				),
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
