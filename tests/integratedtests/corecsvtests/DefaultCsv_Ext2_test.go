package corecsvtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// DefaultCsv
// ============================================================================

func Test_DefaultCsv_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2DefaultCsvTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.DefaultCsv(items...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
			"hasComma": strings.Contains(result, ","),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// DefaultCsvStrings
// ============================================================================

func Test_DefaultCsvStrings_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2DefaultCsvStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.DefaultCsvStrings(items...)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// DefaultCsvUsingJoiner
// ============================================================================

func Test_DefaultCsvUsingJoiner_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2DefaultCsvUsingJoinerTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.DefaultCsvUsingJoiner(joiner, items...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// DefaultAnyCsv
// ============================================================================

func Test_DefaultAnyCsv_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2DefaultAnyCsvTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]any)

		// Act
		result := corecsv.DefaultAnyCsv(items...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// DefaultAnyCsvStrings
// ============================================================================

func Test_DefaultAnyCsvStrings_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2DefaultAnyCsvStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]any)

		// Act
		result := corecsv.DefaultAnyCsvStrings(items...)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// DefaultAnyCsvUsingJoiner
// ============================================================================

func Test_DefaultAnyCsvUsingJoiner_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2DefaultAnyCsvUsingJoinerTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")
		rawItems, _ := input.Get("items")
		items := rawItems.([]any)

		// Act
		result := corecsv.DefaultAnyCsvUsingJoiner(joiner, items...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// StringsToCsvStringsDefault
// ============================================================================

func Test_StringsToCsvStringsDefault_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2StringsToCsvStringsDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.StringsToCsvStringsDefault(items...)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// StringsToStringDefault
// ============================================================================

func Test_StringsToStringDefault_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2StringsToStringDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.StringsToStringDefault(items...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// StringsToCsvString -- all three quote branches
// ============================================================================

func Test_StringsToCsvString_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2StringsToCsvStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		joiner, _ := input.GetAsString("joiner")
		quote := input.GetAsBoolDefault("quote", false)
		singleQuote := input.GetAsBoolDefault("singleQuote", false)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.StringsToCsvString(joiner, quote, singleQuote, items...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// StringsToCsvStrings -- all three quote branches
// ============================================================================

func Test_StringsToCsvStrings_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2StringsToCsvStringsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		quote := input.GetAsBoolDefault("quote", false)
		singleQuote := input.GetAsBoolDefault("singleQuote", false)
		rawItems, _ := input.Get("items")
		items := rawItems.([]string)

		// Act
		result := corecsv.StringsToCsvStrings(quote, singleQuote, items...)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// CompileStringersToString
// ============================================================================

func Test_CompileStringersToString_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2CompileStringersToStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 0)
		joiner, _ := input.GetAsString("joiner")

		var funcs []func() string
		for i := 0; i < count; i++ {
			funcs = append(funcs, func() string { return "val" })
		}

		// Act
		result := corecsv.CompileStringersToString(joiner, true, false, funcs...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// CompileStringersToStringDefault
// ============================================================================

func Test_CompileStringersToStringDefault_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2CompileStringersToStringDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		count := input.GetAsIntDefault("count", 0)

		var funcs []func() string
		for i := 0; i < count; i++ {
			funcs = append(funcs, func() string { return "val" })
		}

		// Act
		result := corecsv.CompileStringersToStringDefault(funcs...)

		// Assert
		actual := args.Map{
			"notEmpty": result != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// StringFunctionsToString -- all three branches
// ============================================================================

func Test_StringFunctionsToString_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2StringFunctionsToStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		quote := input.GetAsBoolDefault("quote", false)
		singleQuote := input.GetAsBoolDefault("singleQuote", false)
		count := input.GetAsIntDefault("count", 0)

		var funcs []func() string
		for i := 0; i < count; i++ {
			funcs = append(funcs, func() string { return "val" })
		}

		// Act
		result := corecsv.StringFunctionsToString(quote, singleQuote, funcs...)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
