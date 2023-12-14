package simplewraptests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/simplewrap"
)

func Test_ParenthesisWrapIf_Wraps_All_Without_Existing_Condition_Checking_Can_Have_DuplicateParenthesis(t *testing.T) {
	for caseIndex, testCase := range parenthesisValidTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for _, input := range inputs {
			actualSlice.Add(
				simplewrap.ParenthesisWrapIf(
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

func Test_ParenthesisWrapIf_Disabled_Wraps_All_Without_Existing_Condition_Checking_Can_Have_DuplicateParenthesis(t *testing.T) {

	for caseIndex, testCase := range parenthesisDisabledRemainsAsItIsTestCases {
		// Arrange
		inputs := testCase.Arrange()
		actualSlice := corestr.New.SimpleSlice.Cap(len(inputs))

		for _, input := range inputs {
			actualSlice.Add(simplewrap.ParenthesisWrapIf(false, input))
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
