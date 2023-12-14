package coreversiontests

import (
	"testing"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coreversion"
)

func Test_Comparison_Verification(t *testing.T) {
	for caseIndex, testCase := range comparisonStringTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.LeftRight)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		creatorFunc := coreversion.New.Default

		// Act
		for index, input := range inputs {
			l := input.Left.(string)
			r := input.Right.(string)
			expectation := input.Expect.(corecomparator.Compare)

			left := creatorFunc(l)
			right := creatorFunc(r)
			isMatch := left.IsExpectedComparison(
				expectation,
				&right,
			)

			actualSlice.AppendFmt(
				comparisonFmt,
				index,
				l,
				expectation.OperatorSymbol(),
				r,
				expectation,
				isMatch,
			)
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
