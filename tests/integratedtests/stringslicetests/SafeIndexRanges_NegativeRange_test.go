package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SafeIndexRanges — uncovered branches ──

func Test_SafeIndexRanges_NegativeRange_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, 3, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns empty -- negative range", actual)
}

func Test_SafeIndexRanges_EmptySlice_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{}, 0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns zeroed slice -- empty input", actual)
}

func Test_SafeIndexRanges_PartialOutOfBounds(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexRanges([]string{"a", "b", "c"}, -1, 4)

	// Act
	actual := args.Map{
		"len": len(result),
		"idx1": result[1],
		"idx2": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 6,
		"idx1": "a",
		"idx2": "b",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns partial -- out of bounds edges", actual)
}

// ── SafeRangeItems — uncovered branches ──

func Test_SafeRangeItems_NilSlice(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems(nil, 0, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns empty -- nil slice", actual)
}

func Test_SafeRangeItems_StartBeyondLast(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns empty -- start beyond last", actual)
}

func Test_SafeRangeItems_EndBeyondLast(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, 100)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns clipped -- end beyond last", actual)
}

func Test_SafeRangeItems_InvalidStart(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, -1, 2)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns from start -- invalid start index", actual)
}

// ── SafeRangeItemsPtr — uncovered branches ──

func Test_SafeRangeItemsPtr_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItemsPtr([]string{}, 0, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns empty -- empty slice", actual)
}

func Test_SafeRangeItemsPtr_Valid_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItemsPtr([]string{"a", "b"}, 0, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns items -- valid range", actual)
}

// ── SafeIndexes — uncovered branches ──

func Test_SafeIndexes_OutOfBounds_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes([]string{"a", "b"}, 0, 5, -1)

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
		"third": result[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "",
		"third": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns partial -- out of bounds indexes", actual)
}

func Test_SafeIndexes_EmptySlice_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexes([]string{}, 0, 1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns zeroed -- empty input slice", actual)
}

// ── SplitTrimmedNonEmpty — uncovered branches ──

func Test_SplitTrimmedNonEmpty_Content(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty("  a , b , c  ", ",", -1)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns trimmed -- comma separated", actual)
}

// ── SplitTrimmedNonEmptyAll ──

func Test_SplitTrimmedNonEmptyAll_Content(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmptyAll("  x | y  ", "|")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll returns trimmed -- pipe separated", actual)
}

// ── TrimmedEachWordsPtr ──

func Test_TrimmedEachWordsPtr_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsPtr([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns empty -- empty input", actual)
}

func Test_TrimmedEachWordsPtr_Items(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsPtr([]string{" a ", " b "})

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns trimmed -- whitespace items", actual)
}

// ── NonWhitespacePtr ──

func Test_NonWhitespacePtr_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespacePtr([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns empty -- empty input", actual)
}

func Test_NonWhitespacePtr_Items(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespacePtr([]string{"a", "  ", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns filtered -- whitespace removed", actual)
}

// ── NonWhitespaceJoinPtr ──

func Test_NonWhitespaceJoinPtr_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoinPtr([]string{}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns empty -- empty input", actual)
}

func Test_NonWhitespaceJoinPtr_Items(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoinPtr([]string{"a", "  ", "b"}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns joined -- whitespace removed", actual)
}

// ── NonEmptyJoinPtr ──

func Test_NonEmptyJoinPtr_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoinPtr([]string{}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns empty -- empty input", actual)
}

func Test_NonEmptyJoinPtr_Items(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoinPtr([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns joined -- empty removed", actual)
}

// ── LastSafeIndexPtr ──

func Test_LastSafeIndexPtr_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.LastSafeIndexPtr([]string{})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns -1 -- empty input", actual)
}

func Test_LastSafeIndexPtr_Items(t *testing.T) {
	// Arrange
	result := stringslice.LastSafeIndexPtr([]string{"a", "b"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns 1 -- two items", actual)
}

// ── SafeIndexAtUsingLastIndex ──

func Test_SafeIndexAtUsingLastIndex_OutOfRange_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 5)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- index out of range", actual)
}

func Test_SafeIndexAtUsingLastIndex_NegativeIndex_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, -1)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- negative index", actual)
}

func Test_SafeIndexAtUsingLastIndex_ZeroLastIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 0)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- lastIndex is 0", actual)
}

// ── CloneSimpleSliceToPointers ──

func Test_CloneSimpleSliceToPointers_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.CloneSimpleSliceToPointers([]string{})

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns empty ptr -- empty input", actual)
}

func Test_CloneSimpleSliceToPointers_Items(t *testing.T) {
	// Arrange
	result := stringslice.CloneSimpleSliceToPointers([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns ptr -- two items", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_RegexTrimmedSplitNonEmptyAll_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`[,;]`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, " a , b ; c ")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns trimmed -- regex split", actual)
}

// ── ExpandBySplits — uncovered branches ──

func Test_ExpandBySplits_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- empty input", actual)
}

func Test_ExpandBySplits_NoSplitters_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a,b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- no splitters", actual)
}

func Test_ExpandBySplits_Items(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplits([]string{"a,b", "c;d"}, ",", ";")

	// Act
	actual := args.Map{"hasItems": len(result) > 2}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns expanded -- multiple splitters", actual)
}

// ── LinesSimpleProcess — uncovered branches ──

func Test_LinesSimpleProcess_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcess([]string{}, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns empty -- empty input", actual)
}

func Test_LinesSimpleProcess_Items(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcess([]string{"a", "b"}, func(s string) string { return s + "!" })

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "a!",
		"second": "b!",
	}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns processed -- append exclamation", actual)
}

// ── LinesSimpleProcessNoEmpty — uncovered branches ──

func Test_LinesSimpleProcessNoEmpty_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcessNoEmpty([]string{}, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns empty -- empty input", actual)
}

func Test_LinesSimpleProcessNoEmpty_SkipsEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "", "b"}, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns filtered -- empty removed", actual)
}

// ── MergeSlicesOfSlices — uncovered branches ──

func Test_MergeSlicesOfSlices_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- no input", actual)
}

func Test_MergeSlicesOfSlices_AllEmpty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{}, []string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- all empty slices", actual)
}

func Test_MergeSlicesOfSlices_Mixed(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{}, []string{"b", "c"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns merged -- mixed slices", actual)
}

// ── AnyItemsCloneIf — uncovered branches ──

func Test_AnyItemsCloneIf_NilNoClone_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneIf(false, 0, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns empty -- nil no clone", actual)
}

func Test_AnyItemsCloneIf_NoClone(t *testing.T) {
	// Arrange
	input := []any{"a", "b"}
	result := stringslice.AnyItemsCloneIf(false, 0, input)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns original -- no clone", actual)
}

func Test_AnyItemsCloneIf_Clone(t *testing.T) {
	// Arrange
	input := []any{"a"}
	result := stringslice.AnyItemsCloneIf(true, 2, input)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns cloned -- clone mode", actual)
}

// ── AnyItemsCloneUsingCap — uncovered branches ──

func Test_AnyItemsCloneUsingCap_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneUsingCap(5, []any{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns empty -- empty input", actual)
}

func Test_AnyItemsCloneUsingCap_Items(t *testing.T) {
	// Arrange
	result := stringslice.AnyItemsCloneUsingCap(2, []any{"x", "y"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns cloned -- two items", actual)
}

// ── ProcessAsync ──

func Test_ProcessAsync_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.ProcessAsync(func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns empty -- no items", actual)
}

func Test_ProcessAsync_Items(t *testing.T) {
	// Arrange
	result := stringslice.ProcessAsync(func(i int, item any) string {
		return item.(string) + "!"
	}, "a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a!",
	}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns processed -- two items", actual)
}

// ── ProcessOptionAsync ──

func Test_ProcessOptionAsync_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(false, func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- no items", actual)
}

func Test_ProcessOptionAsync_SkipNil(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(true, func(i int, item any) string {
		if item == "skip" {
			return ""
		}
		return item.(string)
	}, "a", "skip", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns filtered -- skip empty", actual)
}

func Test_ProcessOptionAsync_ReturnAll_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.ProcessOptionAsync(false, func(i int, item any) string {
		if item == "skip" {
			return ""
		}
		return item.(string)
	}, "a", "skip", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns all -- no skip", actual)
}

// ── LinesAsyncProcess ──

func Test_LinesAsyncProcess_Empty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess([]string{}, func(i int, s string) string { return s })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns empty -- empty input", actual)
}

func Test_LinesAsyncProcess_Items(t *testing.T) {
	// Arrange
	result := stringslice.LinesAsyncProcess([]string{"a", "b"}, func(i int, s string) string { return s + "!" })

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a!",
	}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns processed -- two items", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor ──

func Test_AnyLinesProcessAsync_Nil_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(nil, func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- nil input", actual)
}

func Test_AnyLinesProcessAsync_NotSlice_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor("notslice", func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- non-slice input", actual)
}

func Test_AnyLinesProcessAsync_EmptySlice_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]int{}, func(i int, item any) string { return "" })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- empty slice", actual)
}

func Test_AnyLinesProcessAsync_Items(t *testing.T) {
	// Arrange
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]string{"x", "y"}, func(i int, item any) string {
		return item.(string) + "!"
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "x!",
	}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns processed -- string slice", actual)
}

// ── AppendAnyItemsWithStrings — uncovered branches ──

func Test_AppendAnyItemsWithStrings_NilItem_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{}, nil, "hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings skips nil -- nil and string items", actual)
}

func Test_AppendAnyItemsWithStrings_SkipEmpty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AppendAnyItemsWithStrings(false, true, []string{}, "")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns empty -- skip empty string", actual)
}

// ── AppendStringsWithAnyItems — uncovered branches ──

func Test_AppendStringsWithAnyItems_SkipEmpty_FromSafeIndexRangesNegat(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, true, []any{}, "hello", "")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns filtered -- skip empty", actual)
}

func Test_AppendStringsWithAnyItems_NoItems(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"existing"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns original -- no appending items", actual)
}
