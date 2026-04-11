package csvinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

// ── StringsToCsvStrings all branches ──

func Test_StringsToCsvStrings_SingleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvStrings(true, true, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SingleQuote returns correct value -- with args", actual)
}

func Test_StringsToCsvStrings_NoQuote(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvStrings(false, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NoQuote returns correct value -- with args", actual)
}

func Test_StringsToCsvStrings_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvStrings(false, false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- with args", actual)
}

func Test_StringsToCsvStringsDefault(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvStringsDefault("a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Default returns correct value -- with args", actual)
}

func Test_StringsToStringDefault(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToStringDefault("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToStringDefault returns correct value -- with args", actual)
}

func Test_StringsToStringDefaultNoQuotations(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToStringDefaultNoQuotations("a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NoQuotations returns correct value -- with args", actual)
}

// ── AnyItemsToCsvStrings all branches ──

func Test_AnyItemsToCsvStrings_SingleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvStrings(true, true, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Any_SingleQuote returns correct value -- with args", actual)
}

func Test_AnyItemsToCsvStrings_DoubleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvStrings(true, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Any_DoubleQuote returns correct value -- with args", actual)
}

func Test_AnyItemsToCsvStrings_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvStrings(false, false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Any_Empty returns empty -- with args", actual)
}

func Test_AnyItemsToStringDefault(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToStringDefault("a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToStringDefault returns correct value -- with args", actual)
}

// ── Stringers all branches ──

type cov2Stringer struct{ val string }

func (s cov2Stringer) String() string { return s.val }

func Test_StringersToCsvStrings_AllBranches(t *testing.T) {
	// Arrange
	s := cov2Stringer{val: "x"}
	single := csvinternal.StringersToCsvStrings(true, true, s)
	double := csvinternal.StringersToCsvStrings(true, false, s)
	noQuote := csvinternal.StringersToCsvStrings(false, false, s)
	empty := csvinternal.StringersToCsvStrings(false, false)

	// Act
	actual := args.Map{
		"singleLen": len(single), "doubleLen": len(double),
		"noQuoteLen": len(noQuote), "emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"singleLen": 1, "doubleLen": 1,
		"noQuoteLen": 1, "emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringersToCsvStrings returns correct value -- with args", actual)
}

func Test_StringersToStringDefault(t *testing.T) {
	// Arrange
	s := cov2Stringer{val: "x"}
	result := csvinternal.StringersToStringDefault(s)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToStringDefault returns correct value -- with args", actual)
}

// ── CompileStringers all branches ──

func Test_CompileStringersToCsvStrings_AllBranches(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	single := csvinternal.CompileStringersToCsvStrings(true, true, f)
	double := csvinternal.CompileStringersToCsvStrings(true, false, f)
	noQuote := csvinternal.CompileStringersToCsvStrings(false, false, f)
	empty := csvinternal.CompileStringersToCsvStrings(false, false)

	// Act
	actual := args.Map{
		"singleLen": len(single), "doubleLen": len(double),
		"noQuoteLen": len(noQuote), "emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"singleLen": 1, "doubleLen": 1,
		"noQuoteLen": 1, "emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "CompileStringersCsvStrings returns correct value -- with args", actual)
}

func Test_CompileStringersToStringDefault(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToStringDefault(f)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToStringDefault returns correct value -- with args", actual)
}

// ── RangeNames ──

func Test_RangeNamesWithValuesIndexesCsvString(t *testing.T) {
	// Arrange
	result := csvinternal.RangeNamesWithValuesIndexesCsvString("A", "B")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsvString returns correct value -- with args", actual)
}

// ── AnyItemsToCsvString single quote ──

func Test_AnyItemsToCsvString_SingleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvString(", ", true, true, "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_SingleQuote returns correct value -- with args", actual)
}

// ── StringsToCsvString single quote ──

func Test_StringsToCsvString_SingleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvString(", ", true, true, "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_SingleQuote returns correct value -- with args", actual)
}

