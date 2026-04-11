package csvinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

func Test_StringsToCsvString(t *testing.T) {
	// Act
	noQuote := csvinternal.StringsToCsvString(", ", false, false, "a", "b", "c")
	singleQuote := csvinternal.StringsToCsvString(", ", true, true, "a", "b")
	doubleQuote := csvinternal.StringsToCsvString(", ", true, false, "a", "b")
	empty := csvinternal.StringsToCsvString(", ", false, false)

	actual := args.Map{
		"noQuoteNotEmpty":     noQuote != "",
		"singleQuoteNotEmpty": singleQuote != "",
		"doubleQuoteNotEmpty": doubleQuote != "",
		"empty":               empty,
	}
	expected := args.Map{
		"noQuoteNotEmpty":     true,
		"singleQuoteNotEmpty": true,
		"doubleQuoteNotEmpty": true,
		"empty":               "",
	}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns correct value -- with args", actual)
}

func Test_AnyItemsToCsvString(t *testing.T) {
	// Act
	result := csvinternal.AnyItemsToCsvString(", ", false, false, 1, "hello", true)
	empty := csvinternal.AnyItemsToCsvString(", ", false, false)

	actual := args.Map{
		"resultNotEmpty": result != "",
		"empty":          empty,
	}
	expected := args.Map{
		"resultNotEmpty": true,
		"empty":          "",
	}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns correct value -- with args", actual)
}

func Test_RangeNamesWithValuesIndexes(t *testing.T) {
	// Act
	result := csvinternal.RangeNamesWithValuesIndexes("Major", "Minor", "Patch")
	empty := csvinternal.RangeNamesWithValuesIndexes()

	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(empty),
		"first":     result[0],
	}
	expected := args.Map{
		"resultLen": 3,
		"emptyLen":  0,
		"first":     result[0],
	}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns non-empty -- with args", actual)
}

type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_StringersToString(t *testing.T) {
	// Arrange
	s1 := testStringer{val: "a"}
	s2 := testStringer{val: "b"}

	// Act
	result := csvinternal.StringersToString(", ", false, false, s1, s2)

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "StringersToString returns correct value -- with args", actual)
}

func Test_CompileStringersToString(t *testing.T) {
	// Arrange
	f1 := func() string { return "a" }
	f2 := func() string { return "b" }

	// Act
	result := csvinternal.CompileStringersToString(", ", false, false, f1, f2)

	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString returns correct value -- with args", actual)
}
