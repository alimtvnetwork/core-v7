package corecsvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Default convenience functions
// ═══════════════════════════════════════════

func Test_DefaultCsv_FromDefaultCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.DefaultCsv("a", "b")}

	// Assert
	expected := args.Map{
		"result": `"a", "b"`,
	}
	expected.ShouldBeEqual(t, 0, "DefaultCsv returns double-quoted csv -- two items", actual)
}

func Test_DefaultCsv_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.DefaultCsv()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "DefaultCsv returns empty -- no items", actual)
}

func Test_DefaultAnyCsv_FromDefaultCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.DefaultAnyCsv("x", 1)}

	// Assert
	expected := args.Map{
		"result": `"x", "1"`,
	}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsv returns double-quoted csv -- mixed types", actual)
}

func Test_AnyItemsToStringDefault_FromDefaultCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.AnyItemsToStringDefault("hello")}

	// Assert
	expected := args.Map{"result": `"hello"`}
	expected.ShouldBeEqual(t, 0, "AnyItemsToStringDefault returns double-quoted -- single item", actual)
}

func Test_DefaultCsvStrings_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.DefaultCsvStrings("a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": `"a"`,
	}
	expected.ShouldBeEqual(t, 0, "DefaultCsvStrings returns double-quoted slice -- two items", actual)
}

func Test_DefaultAnyCsvStrings_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.DefaultAnyCsvStrings("x", 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvStrings returns slice -- two items", actual)
}

func Test_DefaultCsvUsingJoiner_FromDefaultCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.DefaultCsvUsingJoiner(" | ", "a", "b")}

	// Assert
	expected := args.Map{"result": `"a" | "b"`}
	expected.ShouldBeEqual(t, 0, "DefaultCsvUsingJoiner returns joined -- pipe joiner", actual)
}

func Test_DefaultAnyCsvUsingJoiner_FromDefaultCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.DefaultAnyCsvUsingJoiner(" | ", "a", "b")}

	// Assert
	expected := args.Map{"result": `"a" | "b"`}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvUsingJoiner returns joined -- pipe joiner", actual)
}

func Test_StringsToCsvStringsDefault_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStringsDefault("a")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": `"a"`,
	}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStringsDefault returns double-quoted -- one item", actual)
}

func Test_StringsToStringDefault_FromDefaultCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": corecsv.StringsToStringDefault("a", "b")}

	// Assert
	expected := args.Map{
		"result": `"a", "b"`,
	}
	expected.ShouldBeEqual(t, 0, "StringsToStringDefault returns joined csv -- two items", actual)
}

// ═══════════════════════════════════════════
// CompileStringers
// ═══════════════════════════════════════════

func Test_CompileStringersToStringDefault_FromDefaultCsv(t *testing.T) {
	// Arrange
	fn1 := func() string { return "hello" }
	fn2 := func() string { return "world" }
	result := corecsv.CompileStringersToStringDefault(fn1, fn2)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToStringDefault returns non-empty -- two funcs", actual)
}

func Test_CompileStringersToString_FromDefaultCsv(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.CompileStringersToString(" | ", true, true, fn)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString returns non-empty -- single quote", actual)
}

func Test_CompileStringersToCsvStrings_SingleQuote(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.CompileStringersToCsvStrings(true, true, fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings returns slice -- single quote", actual)
}

func Test_CompileStringersToCsvStrings_DoubleQuote(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.CompileStringersToCsvStrings(true, false, fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings returns slice -- double quote", actual)
}

func Test_CompileStringersToCsvStrings_NoQuote(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.CompileStringersToCsvStrings(false, false, fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings returns slice -- no quote", actual)
}

func Test_CompileStringersToCsvStrings_Empty(t *testing.T) {
	// Arrange
	result := corecsv.CompileStringersToCsvStrings(true, true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings returns empty -- no funcs", actual)
}

// ═══════════════════════════════════════════
// StringFunctionsToString — all 3 branches
// ═══════════════════════════════════════════

func Test_StringFunctionsToString_SingleQuote(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.StringFunctionsToString(true, true, fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString returns slice -- single quote", actual)
}

func Test_StringFunctionsToString_DoubleQuote(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.StringFunctionsToString(true, false, fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString returns slice -- double quote", actual)
}

func Test_StringFunctionsToString_NoQuote(t *testing.T) {
	// Arrange
	fn := func() string { return "val" }
	result := corecsv.StringFunctionsToString(false, false, fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString returns slice -- no quote", actual)
}

func Test_StringFunctionsToString_Empty(t *testing.T) {
	// Arrange
	result := corecsv.StringFunctionsToString(true, true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString returns empty -- no funcs", actual)
}

// ═══════════════════════════════════════════
// AnyToTypesCsvStrings — all 3 branches
// ═══════════════════════════════════════════

func Test_AnyToTypesCsvStrings_SingleQuote_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(true, true, "hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns slice -- single quote", actual)
}

func Test_AnyToTypesCsvStrings_DoubleQuote_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(true, false, "hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns slice -- double quote", actual)
}

func Test_AnyToTypesCsvStrings_NoQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(false, false, "hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns slice -- no quote", actual)
}

func Test_AnyToTypesCsvStrings_Empty(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(true, true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns empty -- no items", actual)
}

func Test_AnyToTypesCsvDefault_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvDefault("hello", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvDefault returns non-empty -- two items", actual)
}

// ═══════════════════════════════════════════
// AnyToValuesTypeStrings / AnyToValuesTypeString
// ═══════════════════════════════════════════

func Test_AnyToValuesTypeStrings_NonEmpty(t *testing.T) {
	// Arrange
	result := corecsv.AnyToValuesTypeStrings("hello", 42)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings returns slice -- two items", actual)
}

func Test_AnyToValuesTypeStrings_Empty(t *testing.T) {
	// Arrange
	result := corecsv.AnyToValuesTypeStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings returns empty -- no items", actual)
}

func Test_AnyToValuesTypeString_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.AnyToValuesTypeString("hello", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeString returns non-empty -- two items", actual)
}

func Test_AnyToValuesTypeString_Empty(t *testing.T) {
	// Arrange
	result := corecsv.AnyToValuesTypeString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeString returns empty -- no items", actual)
}

// ═══════════════════════════════════════════
// RangeNamesWithValuesIndexes*
// ═══════════════════════════════════════════

func Test_RangeNamesWithValuesIndexesString_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.RangeNamesWithValuesIndexesString(" | ", "A", "B")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesString returns non-empty -- two items", actual)
}

// ═══════════════════════════════════════════
// StringersToCsvStrings — remaining branches
// ═══════════════════════════════════════════

func Test_StringsToCsvStrings_SingleQuote_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStrings(true, true, "a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "'a'",
	}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStrings returns single-quoted -- two items", actual)
}

func Test_AnyItemsToCsvStrings_SingleQuote_FromDefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvStrings(true, true, "a", 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvStrings returns single-quoted -- two items", actual)
}
