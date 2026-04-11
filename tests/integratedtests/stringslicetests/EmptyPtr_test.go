package stringslicetests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Deprecated Ptr variants — ensure every wrapper is exercised
// ══════════════════════════════════════════════════════════════════════════════

// ── EmptyPtr ──

func Test_EmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.EmptyPtr()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyPtr returns empty -- returns empty slice", actual)
}

// ── MakePtr ──

func Test_MakePtr(t *testing.T) {
	// Arrange
	result := stringslice.MakePtr(3, 5)

	// Act
	actual := args.Map{
		"len": len(result),
		"capGe5": cap(result) >= 5,
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"capGe5": true,
	}
	expected.ShouldBeEqual(t, 0, "MakePtr returns non-empty -- returns slice with length and capacity", actual)
}

// ── MakeLenPtr ──

func Test_MakeLenPtr(t *testing.T) {
	// Arrange
	result := stringslice.MakeLenPtr(4)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "MakeLenPtr returns non-empty -- returns slice with length", actual)
}

// ── MakeDefaultPtr ──

func Test_MakeDefaultPtr(t *testing.T) {
	// Arrange
	result := stringslice.MakeDefaultPtr(5)

	// Act
	actual := args.Map{
		"len": len(result),
		"capGe5": cap(result) >= 5,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"capGe5": true,
	}
	expected.ShouldBeEqual(t, 0, "MakeDefaultPtr returns non-empty -- returns zero-len slice with capacity", actual)
}

// ── FirstPtr ──

func Test_FirstPtr_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.FirstPtr([]string{"x", "y"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "FirstPtr returns correct value -- returns first element", actual)
}

// ── LastPtr ──

func Test_LastPtr_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.LastPtr([]string{"x", "y"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "y"}
	expected.ShouldBeEqual(t, 0, "LastPtr returns correct value -- returns last element", actual)
}

// ── FirstOrDefaultPtr ──

func Test_FirstOrDefaultPtr_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.FirstOrDefaultPtr(nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr returns empty -- nil", actual)
}

func Test_FirstOrDefaultPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstOrDefaultPtr([]string{"a"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr returns first -- non-empty", actual)
}

// ── LastOrDefaultPtr ──

func Test_LastOrDefaultPtr_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.LastOrDefaultPtr(nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr returns empty -- nil", actual)
}

func Test_LastOrDefaultPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LastOrDefaultPtr([]string{"a", "b"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr returns last -- non-empty", actual)
}

// ── LastIndexPtr ──

func Test_LastIndexPtr(t *testing.T) {
	// Arrange
	result := stringslice.LastIndexPtr([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexPtr returns correct value -- returns last index", actual)
}

// ── LastSafeIndexPtr ──

func Test_LastSafeIndexPtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.LastSafeIndexPtr(nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns -1 -- nil", actual)
}

func Test_LastSafeIndexPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LastSafeIndexPtr([]string{"a", "b"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns last index -- non-empty", actual)
}

// ── IsEmptyPtr ──

func Test_IsEmptyPtr_True(t *testing.T) {
	// Arrange
	result := stringslice.IsEmptyPtr(nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns true -- nil", actual)
}

func Test_IsEmptyPtr_False(t *testing.T) {
	// Arrange
	result := stringslice.IsEmptyPtr([]string{"a"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns false -- non-empty", actual)
}

// ── HasAnyItemPtr ──

func Test_HasAnyItemPtr_True(t *testing.T) {
	// Arrange
	result := stringslice.HasAnyItemPtr([]string{"a"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr returns true -- has items", actual)
}

func Test_HasAnyItemPtr_False(t *testing.T) {
	// Arrange
	result := stringslice.HasAnyItemPtr(nil)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr returns false -- nil", actual)
}

// ── LengthOfPointer ──

func Test_LengthOfPointer_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.LengthOfPointer([]string{"a", "b"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer returns correct value -- returns length", actual)
}

// ── TrimmedEachWordsPtr ──

func Test_TrimmedEachWordsPtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsPtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns empty -- nil", actual)
}

func Test_TrimmedEachWordsPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsPtr([]string{"  a ", " b "})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns trimmed -- non-empty", actual)
}

// ── NonWhitespacePtr ──

func Test_NonWhitespacePtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespacePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns empty -- nil", actual)
}

func Test_NonWhitespacePtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespacePtr([]string{"a", " ", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns filtered -- non-empty", actual)
}

// ── NonWhitespaceJoinPtr ──

func Test_NonWhitespaceJoinPtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoinPtr(nil, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns empty -- nil", actual)
}

func Test_NonWhitespaceJoinPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoinPtr([]string{"a", " ", "b"}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns joined -- non-empty", actual)
}

// ── NonEmptyJoinPtr ──

func Test_NonEmptyJoinPtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoinPtr(nil, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns empty -- nil", actual)
}

func Test_NonEmptyJoinPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoinPtr([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns joined -- non-empty", actual)
}

// ── SafeRangeItemsPtr ──

func Test_SafeRangeItemsPtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItemsPtr(nil, 0, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns empty -- nil", actual)
}

func Test_SafeRangeItemsPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItemsPtr([]string{"a", "b", "c"}, 0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr delegates to SafeRangeItems -- non-empty", actual)
}

// ── SafeIndexesPtr — empty ──

func Test_SafeIndexesPtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesPtr(nil, 0)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr returns default len slice -- nil input", actual)
}

// ── SafeIndexAtWithPtr — valid ──

func Test_SafeIndexAtWithPtr_Valid_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWithPtr([]string{"a", "b"}, 1, "def")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns element -- valid index", actual)
}

func Test_SafeIndexAtWithPtr_OOB(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWithPtr([]string{"a"}, 5, "def")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns default -- OOB", actual)
}

// ── SlicePtr ──

func Test_SlicePtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SlicePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns empty -- nil", actual)
}

func Test_SlicePtr_NonEmpty(t *testing.T) {
	// Arrange
	input := []string{"a", "b"}
	result := stringslice.SlicePtr(input)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns same slice -- non-empty", actual)
}

// ── ClonePtr ──

func Test_ClonePtr_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns empty -- nil", actual)
}

func Test_ClonePtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns cloned -- non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Core functions — uncovered branches
// ══════════════════════════════════════════════════════════════════════════════

// ── FirstLastDefaultPtr — non-empty ──

func Test_FirstLastDefaultPtr_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	f, l := stringslice.FirstLastDefaultPtr(nil)

	// Act
	actual := args.Map{
		"first": f,
		"last": l,
	}

	// Assert
	expected := args.Map{
		"first": "",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr returns empty -- nil", actual)
}

func Test_FirstLastDefaultPtr_NonEmpty(t *testing.T) {
	// Arrange
	f, l := stringslice.FirstLastDefaultPtr([]string{"a", "b"})

	// Act
	actual := args.Map{
		"first": f,
		"last": l,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr returns first/last -- non-empty", actual)
}

// ── FirstLastDefaultStatusPtr ──

func Test_FirstLastDefaultStatusPtr_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatusPtr(nil)

	// Act
	actual := args.Map{"isValid": result.IsValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr returns invalid -- nil", actual)
}

func Test_FirstLastDefaultStatusPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatusPtr([]string{"a", "b"})

	// Act
	actual := args.Map{
		"isValid": result.IsValid,
		"first": result.First,
		"last": result.Last,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr returns valid -- non-empty", actual)
}

// ── FirstOrDefaultWith ──

func Test_FirstOrDefaultWith_Empty(t *testing.T) {
	// Arrange
	result, ok := stringslice.FirstOrDefaultWith(nil, "fallback")

	// Act
	actual := args.Map{
		"val": result,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "fallback",
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns default -- nil", actual)
}

func Test_FirstOrDefaultWith_NonEmpty(t *testing.T) {
	// Arrange
	result, ok := stringslice.FirstOrDefaultWith([]string{"a"}, "fallback")

	// Act
	actual := args.Map{
		"val": result,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "a",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns first -- non-empty", actual)
}

// ── FirstLastDefault — single element ──

func Test_FirstLastDefault_SingleElem(t *testing.T) {
	// Arrange
	f, l := stringslice.FirstLastDefault([]string{"only"})

	// Act
	actual := args.Map{
		"first": f,
		"last": l,
	}

	// Assert
	expected := args.Map{
		"first": "only",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns first only -- single element", actual)
}

// ── FirstLastDefaultStatus — all branches ──

func Test_FirstLastDefaultStatus_Empty(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatus(nil)

	// Act
	actual := args.Map{
		"isValid": result.IsValid,
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"hasFirst": false,
		"hasLast": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns invalid -- empty", actual)
}

func Test_FirstLastDefaultStatus_Single_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatus([]string{"only"})

	// Act
	actual := args.Map{
		"isValid": result.IsValid,
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
		"first": result.First,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"hasFirst": true,
		"hasLast": false,
		"first": "only",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns partial -- single elem", actual)
}

func Test_FirstLastDefaultStatus_Multi(t *testing.T) {
	// Arrange
	result := stringslice.FirstLastDefaultStatus([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"isValid": result.IsValid,
		"hasFirst": result.HasFirst,
		"hasLast": result.HasLast,
		"first": result.First,
		"last": result.Last,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"hasFirst": true,
		"hasLast": true,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns valid -- multiple elems", actual)
}

// ── InPlaceReverse — all branches ──

func Test_InPlaceReverse_Nil(t *testing.T) {
	// Arrange
	result := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns empty -- nil", actual)
}

func Test_InPlaceReverse_Single(t *testing.T) {
	// Arrange
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns same -- single", actual)
}

func Test_InPlaceReverse_Two_FromEmptyPtr(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[1],
	}

	// Assert
	expected := args.Map{
		"first": "b",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse swaps -- two elements", actual)
}

func Test_InPlaceReverse_Three(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"mid": (*result)[1],
		"last": (*result)[2],
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"mid": "b",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- three elements", actual)
}

func Test_InPlaceReverse_Four(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[3],
	}

	// Assert
	expected := args.Map{
		"first": "d",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- four elements (even)", actual)
}

// ── IndexesDefault — empty slice ──

func Test_IndexesDefault_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault(nil, 0)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns empty -- nil slice", actual)
}

func Test_IndexesDefault_NoIndexes(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault([]string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns empty -- no indexes", actual)
}

// ── SafeIndexesDefaultWithDetail — missing indexes ──

func Test_SafeIndexesDefaultWithDetail_SomeMissing(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b"}, 0, 5, -1)

	// Act
	actual := args.Map{
		"valuesLen":  len(result.Values),
		"missingLen": len(result.MissingIndexes),
		"anyMissing": result.IsAnyMissing,
		"isValid":    result.IsValid,
	}

	// Assert
	expected := args.Map{
		"valuesLen": 1,
		"missingLen": 2,
		"anyMissing": true,
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail reports missing -- some OOB", actual)
}

func Test_SafeIndexesDefaultWithDetail_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail(nil, 0)

	// Act
	actual := args.Map{
		"isValid": result.IsValid,
		"anyMissing": result.IsAnyMissing,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"anyMissing": true,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns invalid -- nil", actual)
}

// ── SafeIndexRanges — negative requestLength ──

func Test_SafeIndexRanges_NegativeRange(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, 3, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns empty -- start > end", actual)
}

func Test_SafeIndexRanges_OOBIndexes(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, -1, 5)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns padded -- OOB handled with empty strings", actual)
}

// ── SplitTrimmedNonEmpty ──

func Test_SplitTrimmedNonEmpty_Basic(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty(" a , b , c ", ",", -1)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns trimmed -- basic", actual)
}

func Test_SplitTrimmedNonEmpty_Limited(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", 2)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b,c",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns limited -- n=2", actual)
}

// ── SplitTrimmedNonEmptyAll ──

func Test_SplitTrimmedNonEmptyAll_Basic(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmptyAll(" x | y | z ", "|")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "x",
		"last": "z",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll returns trimmed -- basic", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_RegexTrimmedSplitNonEmptyAll_FromEmptyPtr(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`[,;]+`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, " a , b ;; c ")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns trimmed -- regex split", actual)
}

func Test_RegexTrimmedSplitNonEmptyAll_AllEmpty(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`.`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "abc")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns empty -- all split to empty", actual)
}

// ── ExpandByFunc — with empty expansion ──

func Test_ExpandByFunc_SomeEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc([]string{"a", "b"}, func(line string) []string {
		if line == "a" {
			return []string{"a1", "a2"}
		}
		return nil
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a1",
		"last": "a2",
	}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns nil -- skips nil expansions", actual)
}

func Test_ExpandByFunc_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(nil, func(line string) []string { return nil })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns empty -- nil input", actual)
}

// ── ExpandBySplit ──

func Test_ExpandBySplit_Basic(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{"a,b", "c,d"}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns expanded -- basic", actual)
}

func Test_ExpandBySplit_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit(nil, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- nil", actual)
}

// ── ExpandBySplits — multiple splitters, no splitters ──

func Test_ExpandBySplits_MultipleSplitters(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a,b;c"}, ",", ";")

	// Act
	actual := args.Map{"len": len(result)}
	// "a,b;c" split by "," => ["a", "b;c"], then "a,b;c" split by ";" => ["a,b", "c"]
	// total = 4

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns expanded -- multiple splitters", actual)
}

func Test_ExpandBySplits_NoSplitters(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- no splitters", actual)
}

// ── AppendStringsWithMainSlice — all branches ──

func Test_AppendStringsWithMainSlice_NoItems(t *testing.T) {
	// Arrange
	input := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(false, input)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns same -- no items", actual)
}

func Test_AppendStringsWithMainSlice_SkipEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"}, "b", "", "c")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice skips empty -- isSkipEmpty true", actual)
}

func Test_AppendStringsWithMainSlice_IncludeEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithMainSlice(false, []string{"a"}, "b", "", "c")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice includes empty -- isSkipEmpty false", actual)
}

// ── AppendStringsWithAnyItems — clone=true branches ──

func Test_AppendStringsWithAnyItems_Clone(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(true, false, []any{"x"}, "a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems clones and appends -- clone true", actual)
}

func Test_AppendStringsWithAnyItems_NoAppend(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"x"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns same -- no items to append", actual)
}

// ── AppendAnyItemsWithStrings — nil item ──

func Test_AppendAnyItemsWithStrings_NilItem(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"}, nil, "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings skips nil -- nil item in list", actual)
}

func Test_AppendAnyItemsWithStrings_Clone(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(true, false, []string{"a"}, "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings clones -- clone true", actual)
}

// ── CloneSimpleSliceToPointers — non-empty ──

func Test_CloneSimpleSliceToPointers_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.CloneSimpleSliceToPointers([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns cloned ptr -- non-empty", actual)
}

// ── LinesProcess — break mid-iteration ──

func Test_LinesProcess_BreakAfterTake(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, s string) (string, bool, bool) {
		return s, true, i == 1
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "LinesProcess stops early -- break after second take", actual)
}

// ── LinesSimpleProcess — identity ──

func Test_LinesSimpleProcess_Identity(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcess([]string{"a", "b"}, func(s string) string { return s })

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns identity -- passthrough", actual)
}

func Test_LinesSimpleProcess_Empty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcess(nil, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns empty -- nil", actual)
}

// ── LinesSimpleProcessNoEmpty — mixed ──

func Test_LinesSimpleProcessNoEmpty_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "b", "c"}, func(s string) string {
		if s == "b" {
			return ""
		}
		return s + "!"
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a!",
		"last": "c!",
	}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty filters empties -- mixed", actual)
}

// ── LinesAsyncProcess — multiple items ──

func Test_LinesAsyncProcess_Multi(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess([]string{"a", "b", "c"}, func(i int, s string) string {
		return fmt.Sprintf("%d:%s", i, s)
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "0:a",
		"last": "2:c",
	}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess processes all -- multiple items", actual)
}

func Test_LinesAsyncProcess_Empty(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess(nil, func(i int, s string) string { return s })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns empty -- nil", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor — not-slice input, nil, empty ──

func Test_AnyLinesProcessAsync_NotSlice(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor("not-a-slice", func(i int, item any) string {
		return ""
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- non-slice input", actual)
}

func Test_AnyLinesProcessAsync_Nil(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(nil, func(i int, item any) string {
		return ""
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- nil", actual)
}

func Test_AnyLinesProcessAsync_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]int{}, func(i int, item any) string {
		return ""
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- empty slice", actual)
}

// ── ProcessAsync — empty ──

func Test_ProcessAsync_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ProcessAsync(func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns empty -- no items", actual)
}

// ── ProcessOptionAsync — isSkipOnNil false ──

func Test_ProcessOptionAsync_NoSkip(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(false, func(i int, item any) string {
		return ""
	}, "a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns all including empty -- isSkipOnNil false", actual)
}

func Test_ProcessOptionAsync_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(true, func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- no items", actual)
}

// ── NonEmptyJoin — all empty strings ──

func Test_NonEmptyJoin_AllEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"", "", ""}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- all empty strings", actual)
}

func Test_NonEmptyJoin_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin(nil, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- nil", actual)
}

func Test_NonEmptyJoin_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns joined -- mixed", actual)
}

// ── NonWhitespaceJoin ──

func Test_NonWhitespaceJoin_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin(nil, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- nil", actual)
}

func Test_NonWhitespaceJoin_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- empty slice", actual)
}

func Test_NonWhitespaceJoin_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{"a", " ", "b"}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns joined -- mixed", actual)
}

// ── MergeNewSimple ──

func Test_MergeNewSimple_Empty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns empty -- no args", actual)
}

func Test_MergeNewSimple_WithEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, nil, []string{"b"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple skips nil -- mixed", actual)
}

// ── SortIf ──

func Test_SortIf_True(t *testing.T) {
	// Arrange
	result := stringslice.SortIf(true, []string{"c", "a", "b"})

	// Act
	actual := args.Map{
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SortIf sorts -- isSort true", actual)
}

func Test_SortIf_False(t *testing.T) {
	// Arrange
	result := stringslice.SortIf(false, []string{"c", "a", "b"})

	// Act
	actual := args.Map{
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "SortIf no-op -- isSort false", actual)
}

// ── JoinWith / Joins — empty ──

func Test_JoinWith_Empty(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- no items", actual)
}

func Test_Joins_Empty(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- no items", actual)
}

func Test_Joins_Multi(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",", "a", "b", "c")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b,c"}
	expected.ShouldBeEqual(t, 0, "Joins returns joined -- multiple items", actual)
}

func Test_JoinWith_Multi(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",", "a", "b")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ",a,b"}
	expected.ShouldBeEqual(t, 0, "JoinWith returns prepended join -- multiple items", actual)
}

// ── CloneIf — non-clone branches ──

func Test_CloneIf_NilNoClone(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf returns empty -- nil no clone", actual)
}

func Test_CloneIf_NonNilNoClone(t *testing.T) {
	// Arrange
	input := []string{"a", "b"}
	result := stringslice.CloneIf(false, 0, input)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "CloneIf returns original -- non-nil no clone", actual)
}

// ── AnyItemsCloneIf — non-clone branches ──

func Test_AnyItemsCloneIf_NilNoClone(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns empty -- nil no clone", actual)
}

func Test_AnyItemsCloneIf_NonNilNoClone(t *testing.T) {
	// Arrange
	input := []any{"a", "b"}
	result := stringslice.AnyItemsCloneIf(false, 0, input)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns original -- non-nil no clone", actual)
}

func Test_AnyItemsCloneIf_NonNilClone(t *testing.T) {
	// Arrange
	input := []any{"a", "b"}
	result := stringslice.AnyItemsCloneIf(true, 3, input)

	// Act
	actual := args.Map{
		"len": len(result),
		"capGe5": cap(result) >= 5,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"capGe5": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf clones with cap -- clone true non-nil", actual)
}

// ── SafeIndexAt — all branches ──

func Test_SafeIndexAt_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{"a", "b"}, 1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns element -- valid", actual)
}

func Test_SafeIndexAt_NegIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{"a"}, -1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- negative index", actual)
}

func Test_SafeIndexAt_OOB(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{"a"}, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- OOB", actual)
}

func Test_SafeIndexAt_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt(nil, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- nil", actual)
}

// ── SafeIndexAtWith — all branches ──

func Test_SafeIndexAtWith_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns element -- valid", actual)
}

func Test_SafeIndexAtWith_NegIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWith([]string{"a"}, -1, "def")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns default -- negative", actual)
}

// ── NonWhitespace — nil ──

func Test_NonWhitespace_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- nil", actual)
}

func Test_NonWhitespace_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- empty slice", actual)
}

// ── NonEmptyStrings — nil ──

func Test_NonEmptyStrings_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- nil", actual)
}

func Test_NonEmptyStrings_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- empty slice", actual)
}

// ── AllElemLengthSlices — nil among slices, no args ──

func Test_AllElemLengthSlices_NoArgs(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices()

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns 0 -- no args", actual)
}

func Test_AllElemLengthSlices_NilAmong(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices([]string{"a"}, nil, []string{"b", "c"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 3}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices skips nil -- mixed", actual)
}

// ── PrependLineNew ──

func Test_PrependLineNew_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.PrependLineNew("first", []string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "first",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "PrependLineNew returns correct value -- prepends single line", actual)
}

// ── AppendLineNew ──

func Test_AppendLineNew_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{"a"}, "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendLineNew returns correct value -- appends single line", actual)
}

// ── Clone ──

func Test_Clone_NonEmpty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.Clone([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns deep copy -- non-empty", actual)
}

func Test_Clone_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.Clone(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- nil", actual)
}

// ── Simple accessors: First, Last, IndexAt, Empty, IsEmpty, HasAnyItem ──

func Test_First_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.First([]string{"x", "y"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "First returns correct value -- returns first element", actual)
}

func Test_Last_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.Last([]string{"x", "y"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "y"}
	expected.ShouldBeEqual(t, 0, "Last returns correct value -- returns last element", actual)
}

func Test_IndexAt_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.IndexAt([]string{"a", "b", "c"}, 2)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "c"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns correct value -- returns element at index", actual)
}

func Test_Empty_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.Empty()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- returns empty slice", actual)
}

func Test_IsEmpty_True(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.IsEmpty(nil)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty true -- nil", actual)
}

func Test_IsEmpty_False(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.IsEmpty([]string{"a"})}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty false -- has items", actual)
}

func Test_HasAnyItem_True(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.HasAnyItem([]string{"a"})}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem true -- has items", actual)
}

func Test_HasAnyItem_False(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.HasAnyItem(nil)}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem false -- nil", actual)
}

// ── FirstOrDefault ──

func Test_FirstOrDefault_Empty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefault(nil)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- nil", actual)
}

func Test_FirstOrDefault_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{"a"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns first -- non-empty", actual)
}

// ── LastOrDefault ──

func Test_LastOrDefault_Empty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefault(nil)}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- nil", actual)
}

func Test_LastOrDefault_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefault([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns last -- non-empty", actual)
}

// ── MakeDefault / MakeLen / Make ──

func Test_MakeDefault_FromEmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.MakeDefault(10)

	// Act
	actual := args.Map{
		"len": len(result),
		"capGe10": cap(result) >= 10,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"capGe10": true,
	}
	expected.ShouldBeEqual(t, 0, "MakeDefault returns non-empty -- returns zero-len with capacity", actual)
}

func Test_MakeLen(t *testing.T) {
	// Arrange
	result := stringslice.MakeLen(3)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLen returns non-empty -- returns slice with length", actual)
}

func Test_Make(t *testing.T) {
	// Arrange
	result := stringslice.Make(2, 5)

	// Act
	actual := args.Map{
		"len": len(result),
		"capGe5": cap(result) >= 5,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"capGe5": true,
	}
	expected.ShouldBeEqual(t, 0, "Make returns non-empty -- returns slice with length and capacity", actual)
}

// ── SafeIndexes — OOB and negative ──

func Test_SafeIndexes_WithOOB(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes([]string{"a", "b"}, 0, 5, -1, 1)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"second": result[1],
		"third": result[2],
		"fourth": result[3],
	}

	// Assert
	expected := args.Map{
		"len": 4,
		"first": "a",
		"second": "",
		"third": "",
		"fourth": "b",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns empty -- handles OOB and negative with empty strings", actual)
}

func Test_SafeIndexes_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes(nil, 0)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns empty -- returns default for empty slice", actual)
}

// ── NonEmptySlice — with empty strings ──

func Test_NonEmptySlice_MixedEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "", "b", ""})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice filters empties -- mixed", actual)
}

func Test_NonEmptySlice_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- nil", actual)
}

// ── NonNullStrings — nil ──

func Test_NonNullStrings_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns empty -- nil", actual)
}

// ── SplitContentsByWhitespace — with content ──

func Test_SplitContentsByWhitespace_Multi(t *testing.T) {
	// Arrange
	result := stringslice.SplitContentsByWhitespace("  a  b  c  ")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace returns fields -- whitespace separated", actual)
}
