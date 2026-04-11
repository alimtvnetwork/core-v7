package anycmptests

import (
	"testing"

	"github.com/alimtvnetwork/core/anycmp"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Coverage test for anycmp.Cmp to ensure all branches reach 100%.
// Existing CmpBranch tests cover the main cases; these cover edge paths
// in compound conditions.

var extCmpEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "Cmp nil slice vs nil slice returns Equal (both reflection-null)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: ([]int)(nil), Second: ([]string)(nil)},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp nil map vs nil map returns Equal (both reflection-null)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: (map[string]int)(nil), Second: (map[int]int)(nil)},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp nil func vs nil func returns Equal (both reflection-null)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: (func())(nil), Second: (func(int))(nil)},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp bool true vs true returns Equal (== match)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: true, Second: true},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp bool true vs false returns Inconclusive",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: true, Second: false},
		},
		ExpectedInput: args.Map{"name": "Inconclusive"},
	},
	{
		Title: "Cmp float64 same returns Equal",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: 3.14, Second: 3.14},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
}

func Test_Cmp_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range extCmpEdgeCaseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
