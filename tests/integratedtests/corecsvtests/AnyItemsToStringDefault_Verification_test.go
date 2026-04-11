package corecsvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_AnyItemsToStringDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range anyItemsToStringDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]any)

		// Act
		result := corecsv.AnyItemsToStringDefault(items...)

		actual := args.Map{
			"notEmpty": result != "",
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_CompileStringersToCsvStrings_Verification(t *testing.T) {
	for caseIndex, testCase := range compileStringersToCsvStringsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 0)
		includeQuote := input.GetAsBoolDefault("includeQuote", false)
		singleQuote := input.GetAsBoolDefault("singleQuote", false)

		var funcs []func() string
		for i := 0; i < count; i++ {
			funcs = append(funcs, func() string { return "val" })
		}

		// Act
		result := corecsv.CompileStringersToCsvStrings(includeQuote, singleQuote, funcs...)

		actual := args.Map{
			"length": len(result),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
