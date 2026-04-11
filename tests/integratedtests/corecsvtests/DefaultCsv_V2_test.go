package corecsvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── DefaultCsv / DefaultCsvStrings ──

func Test_DefaultCsv(t *testing.T) {
	// Arrange
	result := corecsv.DefaultCsv("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultCsv returns correct value -- with args", actual)
}

func Test_DefaultCsvStrings(t *testing.T) {
	// Arrange
	result := corecsv.DefaultCsvStrings("a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultCsvStrings returns correct value -- with args", actual)
}

func Test_DefaultCsvUsingJoiner(t *testing.T) {
	// Arrange
	result := corecsv.DefaultCsvUsingJoiner(" | ", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultCsvUsingJoiner returns correct value -- with args", actual)
}

// ── DefaultAnyCsv / DefaultAnyCsvStrings ──

func Test_DefaultAnyCsv(t *testing.T) {
	// Arrange
	result := corecsv.DefaultAnyCsv("a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsv returns correct value -- with args", actual)
}

func Test_DefaultAnyCsvStrings(t *testing.T) {
	// Arrange
	result := corecsv.DefaultAnyCsvStrings("a", 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvStrings returns correct value -- with args", actual)
}

func Test_DefaultAnyCsvUsingJoiner(t *testing.T) {
	// Arrange
	result := corecsv.DefaultAnyCsvUsingJoiner(" | ", "a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DefaultAnyCsvUsingJoiner returns correct value -- with args", actual)
}

// ── StringsToCsvStrings all quote branches ──

func Test_StringsToCsvStrings_SingleQuote(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStrings(true, true, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SingleQuote returns correct value -- with args", actual)
}

func Test_StringsToCsvStrings_NoQuote(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStrings(false, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NoQuote returns correct value -- with args", actual)
}

func Test_StringsToCsvStringsDefault(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStringsDefault("a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStringsDefault returns correct value -- with args", actual)
}

func Test_StringsToStringDefault(t *testing.T) {
	// Arrange
	result := corecsv.StringsToStringDefault("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToStringDefault returns correct value -- with args", actual)
}

// ── AnyItemsToCsvStrings all quote branches ──

func Test_AnyItemsToCsvStrings_SingleQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvStrings(true, true, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItems_SingleQuote returns correct value -- with args", actual)
}

func Test_AnyItemsToCsvStrings_NoQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvStrings(false, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItems_NoQuote returns correct value -- with args", actual)
}

func Test_AnyItemsToStringDefault(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToStringDefault("a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToStringDefault returns correct value -- with args", actual)
}

// ── AnyToTypesCsvStrings all branches ──

func Test_AnyToTypesCsvStrings_SingleQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(true, true, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Types_SingleQuote returns correct value -- with args", actual)
}

func Test_AnyToTypesCsvStrings_DoubleQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(true, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Types_DoubleQuote returns correct value -- with args", actual)
}

func Test_AnyToTypesCsvDefault(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvDefault("a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvDefault returns correct value -- with args", actual)
}

// ── AnyToValuesTypeStrings ──

func Test_AnyToValuesTypeStrings(t *testing.T) {
	// Arrange
	result := corecsv.AnyToValuesTypeStrings("a", 1)
	empty := corecsv.AnyToValuesTypeStrings()

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings returns non-empty -- with args", actual)
}

func Test_AnyToValuesTypeString(t *testing.T) {
	// Arrange
	result := corecsv.AnyToValuesTypeString("a")
	empty := corecsv.AnyToValuesTypeString()

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"empty": empty,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeString returns non-empty -- with args", actual)
}

// ── Stringers ──

type covTestStringer struct{ val string }

func (s covTestStringer) String() string { return s.val }

func Test_StringersToCsvStrings_AllBranches(t *testing.T) {
	// Arrange
	s := covTestStringer{val: "x"}
	single := corecsv.StringersToCsvStrings(true, true, s)
	double := corecsv.StringersToCsvStrings(true, false, s)
	noQuote := corecsv.StringersToCsvStrings(false, false, s)
	empty := corecsv.StringersToCsvStrings(false, false)

	// Act
	actual := args.Map{
		"singleLen":  len(single),
		"doubleLen":  len(double),
		"noQuoteLen": len(noQuote),
		"emptyLen":   len(empty),
	}

	// Assert
	expected := args.Map{
		"singleLen":  1,
		"doubleLen":  1,
		"noQuoteLen": 1,
		"emptyLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "StringersToCsvStrings returns correct value -- with args", actual)
}

func Test_StringersToString(t *testing.T) {
	// Arrange
	s := covTestStringer{val: "x"}
	result := corecsv.StringersToString(", ", false, false, s)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString returns correct value -- with args", actual)
}

func Test_StringersToStringDefault(t *testing.T) {
	// Arrange
	s := covTestStringer{val: "x"}
	result := corecsv.StringersToStringDefault(s)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToStringDefault returns correct value -- with args", actual)
}

// ── CompileStringers ──

func Test_CompileStringersToCsvStrings_AllBranches(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	single := corecsv.CompileStringersToCsvStrings(true, true, f)
	double := corecsv.CompileStringersToCsvStrings(true, false, f)
	noQuote := corecsv.CompileStringersToCsvStrings(false, false, f)
	empty := corecsv.CompileStringersToCsvStrings(false, false)

	// Act
	actual := args.Map{
		"singleLen":  len(single),
		"doubleLen":  len(double),
		"noQuoteLen": len(noQuote),
		"emptyLen":   len(empty),
	}

	// Assert
	expected := args.Map{
		"singleLen":  1,
		"doubleLen":  1,
		"noQuoteLen": 1,
		"emptyLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "CompileStringersToCsvStrings returns correct value -- with args", actual)
}

func Test_CompileStringersToString(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	result := corecsv.CompileStringersToString(", ", false, false, f)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString returns correct value -- with args", actual)
}

func Test_CompileStringersToStringDefault(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	result := corecsv.CompileStringersToStringDefault(f)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToStringDefault returns correct value -- with args", actual)
}

// ── StringFunctionsToString all branches ──

func Test_StringFunctionsToString_AllBranches(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	single := corecsv.StringFunctionsToString(true, true, f)
	double := corecsv.StringFunctionsToString(true, false, f)
	noQuote := corecsv.StringFunctionsToString(false, false, f)
	empty := corecsv.StringFunctionsToString(false, false)

	// Act
	actual := args.Map{
		"singleLen":  len(single),
		"doubleLen":  len(double),
		"noQuoteLen": len(noQuote),
		"emptyLen":   len(empty),
	}

	// Assert
	expected := args.Map{
		"singleLen":  1,
		"doubleLen":  1,
		"noQuoteLen": 1,
		"emptyLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "StringFunctionsToString returns correct value -- with args", actual)
}

// ── RangeNames ──

func Test_RangeNamesWithValuesIndexesCsvString(t *testing.T) {
	// Arrange
	result := corecsv.RangeNamesWithValuesIndexesCsvString("A", "B")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesCsvString returns non-empty -- with args", actual)
}

func Test_RangeNamesWithValuesIndexesString(t *testing.T) {
	// Arrange
	result := corecsv.RangeNamesWithValuesIndexesString(" | ", "A", "B")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexesString returns non-empty -- with args", actual)
}

// ── Empty-items branch coverage ──

func Test_AnyItemsToCsvString_EmptyItems(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvString(", ", true, false)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns empty -- no items", actual)
}

func Test_StringsToCsvString_EmptyItems(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvString(", ", true, false)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns empty -- no items", actual)
}

func Test_RangeNamesWithValuesIndexes_EmptyItems(t *testing.T) {
	// Arrange
	result := corecsv.RangeNamesWithValuesIndexes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns empty -- no items", actual)
}

func Test_AnyToValuesTypeStrings_EmptyStringItem(t *testing.T) {
	// Arrange
	// Covers the finalString == "" branch inside AnyToValuesTypeStrings
	result := corecsv.AnyToValuesTypeStrings(nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"notPanicked": true,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"notPanicked": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyToValuesTypeStrings covers empty-string item -- nil input", actual)
}
