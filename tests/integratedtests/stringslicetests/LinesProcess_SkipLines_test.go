package stringslicetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── LinesProcess — isTake false branch ──

func Test_LinesProcess_SkipLines(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, s string) (string, bool, bool) {
		// skip even indexes, take odd
		return s, i%2 == 1, false
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns filtered -- isTake false skips", actual)
}

func Test_LinesProcess_BreakWithoutTake(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, s string) (string, bool, bool) {
		return s, false, i == 0
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns empty -- break immediately without take", actual)
}

func Test_LinesProcess_TakeAll(t *testing.T) {
	// Arrange
	result := stringslice.LinesProcess([]string{"a", "b"}, func(i int, s string) (string, bool, bool) {
		return s + "!", true, false
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
		"last": "b!",
	}
	expected.ShouldBeEqual(t, 0, "LinesProcess returns all processed -- take all no break", actual)
}

// ── SafeRangeItems — InvalidValue branches ──

func Test_SafeRangeItems_EndInvalidValue(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, constants.InvalidValue)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns clipped -- end is InvalidValue", actual)
}

func Test_SafeRangeItems_StartInvalidValue(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, constants.InvalidValue, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns from start -- start is InvalidValue", actual)
}

func Test_SafeRangeItems_ValidRange(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c", "d"}, 1, 3)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "b",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns range -- valid start and end", actual)
}

// ── SafeIndexesDefaultWithDetail — all valid indexes ──

func Test_SafeIndexesDefaultWithDetail_AllValid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b", "c"}, 0, 1, 2)

	// Act
	actual := args.Map{
		"valuesLen":  len(result.Values),
		"anyMissing": result.IsAnyMissing,
		"isValid":    result.IsValid,
	}

	// Assert
	expected := args.Map{
		"valuesLen": 3,
		"anyMissing": false,
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns all valid -- no missing indexes", actual)
}

func Test_SafeIndexesDefaultWithDetail_NoIndexes(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b"})

	// Act
	actual := args.Map{
		"valuesLen":  len(result.Values),
		"anyMissing": result.IsAnyMissing,
		"isValid":    result.IsValid,
	}

	// Assert
	expected := args.Map{
		"valuesLen": 0,
		"anyMissing": false,
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns empty values -- no indexes requested", actual)
}

// ── InvalidIndexValuesDetail ──

func Test_InvalidIndexValuesDetail_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.InvalidIndexValuesDetail()

	// Act
	actual := args.Map{
		"valuesLen":    len(result.Values),
		"missingLen":   len(result.MissingIndexes),
		"isAnyMissing": result.IsAnyMissing,
		"isValid":      result.IsValid,
	}

	// Assert
	expected := args.Map{
		"valuesLen": 0,
		"missingLen": 0,
		"isAnyMissing": true,
		"isValid": false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidIndexValuesDetail returns invalid -- default", actual)
}

// ── InvalidFirstLastStatus / InvalidFirstLastStatusForLast ──

func Test_InvalidFirstLastStatus_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.InvalidFirstLastStatus()

	// Act
	actual := args.Map{
		"first": result.First, "last": result.Last,
		"isValid": result.IsValid, "hasFirst": result.HasFirst, "hasLast": result.HasLast,
	}

	// Assert
	expected := args.Map{
		"first": "",
		"last": "",
		"isValid": false,
		"hasFirst": false,
		"hasLast": false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidFirstLastStatus returns all false -- default", actual)
}

func Test_InvalidFirstLastStatusForLast_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.InvalidFirstLastStatusForLast("hello")

	// Act
	actual := args.Map{
		"first": result.First, "last": result.Last,
		"isValid": result.IsValid, "hasFirst": result.HasFirst, "hasLast": result.HasLast,
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"last": "",
		"isValid": false,
		"hasFirst": true,
		"hasLast": false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidFirstLastStatusForLast returns partial -- first only", actual)
}

// ── MergeSlicesOfSlices — all nil slices ──

func Test_MergeSlicesOfSlices_AllNil(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices(nil, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- all nil slices", actual)
}

// ── TrimmedEachWords — all items trim to whitespace ──

func Test_TrimmedEachWords_AllWhitespace(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{"  ", "\t", "\n"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns empty -- all whitespace", actual)
}

func Test_TrimmedEachWords_NilReturnsNil(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords(nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns nil -- nil input", actual)
}

// ── NonEmptySlice — all items non-empty ──

func Test_NonEmptySlice_AllNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "b", "c"})

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
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns all -- all non-empty", actual)
}

// ── NonEmptyStrings — all items empty ──

func Test_NonEmptyStrings_AllEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyStrings([]string{"", "", ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- all empty strings", actual)
}

// ── NonNullStrings — all items empty ──

func Test_NonNullStrings_AllEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings([]string{"", ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns empty -- all empty strings", actual)
}

// ── NonWhitespace — all whitespace ──

func Test_NonWhitespace_AllWhitespace(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{"  ", "\t", "\n", ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- all whitespace", actual)
}

// ── SafeIndexAtUsingLastIndexPtr — distinct branches ──

func Test_SafeIndexAtUsingLastIndexPtr_LastIndexZero_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, 0, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns empty -- lastIndex is 0", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_LastIndexNeg(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a"}, -1, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns empty -- lastIndex negative", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_LastIndexLessThanIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a", "b"}, 1, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns empty -- lastIndex less than index", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_IndexNeg(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a", "b"}, 1, -1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns empty -- index negative", actual)
}

func Test_SafeIndexAtUsingLastIndexPtr_Valid_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndexPtr([]string{"a", "b", "c"}, 2, 1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndexPtr returns element -- valid", actual)
}

// ── SafeIndexAtUsingLastIndex — all condition branches ──

func Test_SafeIndexAtUsingLastIndex_LastIndexZero(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- lastIndex zero", actual)
}

func Test_SafeIndexAtUsingLastIndex_Valid_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a", "b", "c"}, 2, 1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns element -- valid", actual)
}

// ── SafeIndexAtWithPtr — all branches ──

func Test_SafeIndexAtWithPtr_NilSlice(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWithPtr(nil, 0, "def")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns default -- nil slice", actual)
}

func Test_SafeIndexAtWithPtr_NegIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWithPtr([]string{"a"}, -1, "def")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns default -- negative index", actual)
}

// ── SafeIndexesPtr — non-empty slice ──

func Test_SafeIndexesPtr_NonEmpty_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexesPtr([]string{"a", "b", "c"}, 0, 2)

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
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexesPtr returns correct -- valid indexes", actual)
}

// ── AppendAnyItemsWithStrings — isSkipOnEmpty with empty fmt val ──

func Test_AppendAnyItemsWithStrings_SkipEmptyFmt(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, true, []string{"a"}, "")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings skips empty val -- isSkipOnEmpty true", actual)
}

func Test_AppendAnyItemsWithStrings_NoSkipEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"}, "")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings includes empty val -- isSkipOnEmpty false", actual)
}

// ── AppendStringsWithAnyItems — isSkipOnEmpty true and false ──

func Test_AppendStringsWithAnyItems_SkipEmpty_FromLinesProcessSkipLine(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, true, []any{"x"}, "a", "")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems skips empty -- isSkipOnEmpty true", actual)
}

func Test_AppendStringsWithAnyItems_IncludeEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"x"}, "a", "")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems includes empty -- isSkipOnEmpty false", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor — int slice ──

func Test_AnyLinesProcessAsync_IntSlice(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]int{1, 2, 3}, func(i int, item any) string {
		return fmt.Sprintf("%v", item)
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
		"first": "1",
		"last": "3",
	}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns processed -- int slice", actual)
}

// ── LinesAsyncProcess — single item ──

func Test_LinesAsyncProcess_Single(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess([]string{"only"}, func(i int, s string) string { return s + "!" })

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "only!",
	}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns processed -- single item", actual)
}

// ── ProcessAsync — single item ──

func Test_ProcessAsync_Single(t *testing.T) {
	// Arrange
	result := stringslice.ProcessAsync(func(i int, item any) string {
		return item.(string) + "!"
	}, "only")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "only!",
	}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns processed -- single item", actual)
}

// ── ProcessOptionAsync — all empty with skip ──

func Test_ProcessOptionAsync_AllEmptySkip(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(true, func(i int, item any) string {
		return ""
	}, "a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- all empty with skip", actual)
}

// ── CloneUsingCap — nil input ──

func Test_CloneUsingCap_Nil(t *testing.T) {
	// Arrange
	result := stringslice.CloneUsingCap(5, nil)

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
	expected.ShouldBeEqual(t, 0, "CloneUsingCap returns empty with cap -- nil input", actual)
}

// ── AnyItemsCloneUsingCap — nil input ──

func Test_AnyItemsCloneUsingCap_Nil(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneUsingCap(3, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns empty -- nil input", actual)
}

// ── AnyItemsCloneIf — nil with clone ──

func Test_AnyItemsCloneIf_NilClone(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneIf(true, 3, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns empty -- nil with clone", actual)
}

// ── SafeIndexRanges — same start and end (single element) ──

func Test_SafeIndexRanges_SingleElement(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b", "c"}, 1, 1)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns single -- same start and end", actual)
}

// ── SafeIndexes — all valid indexes ──

func Test_SafeIndexes_AllValid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes([]string{"a", "b", "c"}, 0, 1, 2)

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
		"third": result[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
		"third": "c",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns all -- all valid indexes", actual)
}

// ── IndexesDefault — single index ──

func Test_IndexesDefault_SingleIndex(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 1)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns one -- single index", actual)
}

// ── AllElemLengthSlices — single non-nil ──

func Test_AllElemLengthSlices_Single(t *testing.T) {
	// Arrange
	result := stringslice.AllElemLengthSlices([]string{"a", "b"})

	// Act
	actual := args.Map{"count": result}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns 2 -- single slice", actual)
}

// ── NonEmptyIf — with nil slice ──

func Test_NonEmptyIf_NilTrue(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(true, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf returns empty -- nil with true", actual)
}

func Test_NonEmptyIf_NilFalse(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(false, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf returns empty -- nil with false", actual)
}

// ── TrimmedEachWordsIf — nil input ──

func Test_TrimmedEachWordsIf_NilTrue(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(true, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf returns nil -- nil with true", actual)
}

func Test_TrimmedEachWordsIf_NilFalse(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(false, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf returns empty -- nil with false", actual)
}

// ── MergeNew — nil first slice ──

func Test_MergeNew_NilFirstWithItems(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew(nil, "a", "b")

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
	expected.ShouldBeEqual(t, 0, "MergeNew returns items only -- nil first slice", actual)
}

func Test_MergeNew_NilBoth(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNew returns empty -- nil both", actual)
}

// ── PrependNew — nil both ──

func Test_PrependNew_NilSlice(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew(nil, "a")

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
	expected.ShouldBeEqual(t, 0, "PrependNew returns prepended only -- nil second slice", actual)
}

// ── CloneSimpleSliceToPointers — nil ──

func Test_CloneSimpleSliceToPointers_Nil(t *testing.T) {
	// Arrange
	result := stringslice.CloneSimpleSliceToPointers(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns empty ptr -- nil input", actual)
}

// ── NonEmptySlicePtr — nil ──

func Test_NonEmptySlicePtr_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlicePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr returns empty -- nil input", actual)
}

// ── LinesSimpleProcessNoEmpty — all empty output ──

func Test_LinesSimpleProcessNoEmpty_AllEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "b"}, func(lineIn string) string {
		return ""
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns empty -- all processor returns empty", actual)
}

// ── ExpandBySplits — single splitter ──

func Test_ExpandBySplits_SingleSplitter(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a,b,c"}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns expanded -- single splitter", actual)
}

// ── SplitContentsByWhitespace — empty string ──

func Test_SplitContentsByWhitespace_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SplitContentsByWhitespace("")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace returns empty -- empty input", actual)
}

// ── Joins / JoinWith — single item ──

func Test_JoinWith_SingleItem(t *testing.T) {
	// Arrange
	result := stringslice.JoinWith(",", "a")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ",a"}
	expected.ShouldBeEqual(t, 0, "JoinWith returns prepended -- single item", actual)
}

func Test_Joins_SingleItem(t *testing.T) {
	// Arrange
	result := stringslice.Joins(",", "a")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "Joins returns item -- single item", actual)
}

// ── CloneIf — nil with clone ──

func Test_CloneIf_NilWithClone(t *testing.T) {
	// Arrange
	result := stringslice.CloneIf(true, 5, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf returns empty -- nil with clone true", actual)
}
