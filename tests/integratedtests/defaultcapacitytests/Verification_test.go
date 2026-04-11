package defaultcapacitytests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/defaultcapacity"
)

func Test_OfSplits_Verification(t *testing.T) {
	for caseIndex, testCase := range ofSplitsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		wholeLength := input.GetAsIntDefault("wholeLength", 0)
		limit := input.GetAsIntDefault("limit", 0)

		// Act
		result := defaultcapacity.OfSplits(wholeLength, limit)

		actual := args.Map{}
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		if _, has := expected["isPositive"]; has {
			actual["isPositive"] = result > 0
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PredictiveDefaultSmall_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveDefaultSmallTestCases {
		// Arrange
		input := testCase.ArrangeInput.(int)

		// Act
		result := defaultcapacity.PredictiveDefaultSmall(input)

		actual := args.Map{
			"isPositive": result > 0,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PredictiveFiftyPercentIncrement_Verification(t *testing.T) {
	for caseIndex, testCase := range predictiveFiftyPercentTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		possibleLen := input.GetAsIntDefault("possibleLen", 0)
		additionalCap := input.GetAsIntDefault("additionalCap", 0)

		// Act
		result := defaultcapacity.PredictiveFiftyPercentIncrement(possibleLen, additionalCap)

		actual := args.Map{}
		expected := testCase.ExpectedInput.(args.Map)

		if _, has := expected["result"]; has {
			actual["result"] = result
		}
		if _, has := expected["isPositive"]; has {
			actual["isPositive"] = result > 0
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
