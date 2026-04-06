package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SafeIndexRanges — uncovered branches ──

func Test_Cov8_SafeIndexRanges_NegativeRange(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, 3, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns empty -- negative range", actual)
}

func Test_Cov8_SafeIndexRanges_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{}, 0, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns zeroed slice -- empty input", actual)
}

func Test_Cov8_SafeIndexRanges_PartialOutOfBounds(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a", "b", "c"}, -1, 4)
	actual := args.Map{"len": len(result), "idx1": result[1], "idx2": result[2]}
	expected := args.Map{"len": 6, "idx1": "a", "idx2": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns partial -- out of bounds edges", actual)
}

// ── SafeRangeItems — uncovered branches ──

func Test_Cov8_SafeRangeItems_NilSlice(t *testing.T) {
	result := stringslice.SafeRangeItems(nil, 0, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns empty -- nil slice", actual)
}

func Test_Cov8_SafeRangeItems_StartBeyondLast(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns empty -- start beyond last", actual)
}

func Test_Cov8_SafeRangeItems_EndBeyondLast(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, 0, 100)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns clipped -- end beyond last", actual)
}

func Test_Cov8_SafeRangeItems_InvalidStart(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a", "b", "c"}, -1, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns from start -- invalid start index", actual)
}

// ── SafeRangeItemsPtr — uncovered branches ──

func Test_Cov8_SafeRangeItemsPtr_Empty(t *testing.T) {
	result := stringslice.SafeRangeItemsPtr([]string{}, 0, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns empty -- empty slice", actual)
}

func Test_Cov8_SafeRangeItemsPtr_Valid(t *testing.T) {
	result := stringslice.SafeRangeItemsPtr([]string{"a", "b"}, 0, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns items -- valid range", actual)
}

// ── SafeIndexes — uncovered branches ──

func Test_Cov8_SafeIndexes_OutOfBounds(t *testing.T) {
	result := stringslice.SafeIndexes([]string{"a", "b"}, 0, 5, -1)
	actual := args.Map{"first": result[0], "second": result[1], "third": result[2]}
	expected := args.Map{"first": "a", "second": "", "third": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns partial -- out of bounds indexes", actual)
}

func Test_Cov8_SafeIndexes_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexes([]string{}, 0, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns zeroed -- empty input slice", actual)
}

// ── SplitTrimmedNonEmpty — uncovered branches ──

func Test_Cov8_SplitTrimmedNonEmpty_Content(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("  a , b , c  ", ",", -1)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "a"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns trimmed -- comma separated", actual)
}

// ── SplitTrimmedNonEmptyAll ──

func Test_Cov8_SplitTrimmedNonEmptyAll_Content(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll("  x | y  ", "|")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "x"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll returns trimmed -- pipe separated", actual)
}

// ── TrimmedEachWordsPtr ──

func Test_Cov8_TrimmedEachWordsPtr_Empty(t *testing.T) {
	result := stringslice.TrimmedEachWordsPtr([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns empty -- empty input", actual)
}

func Test_Cov8_TrimmedEachWordsPtr_Items(t *testing.T) {
	result := stringslice.TrimmedEachWordsPtr([]string{" a ", " b "})
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr returns trimmed -- whitespace items", actual)
}

// ── NonWhitespacePtr ──

func Test_Cov8_NonWhitespacePtr_Empty(t *testing.T) {
	result := stringslice.NonWhitespacePtr([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns empty -- empty input", actual)
}

func Test_Cov8_NonWhitespacePtr_Items(t *testing.T) {
	result := stringslice.NonWhitespacePtr([]string{"a", "  ", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns filtered -- whitespace removed", actual)
}

// ── NonWhitespaceJoinPtr ──

func Test_Cov8_NonWhitespaceJoinPtr_Empty(t *testing.T) {
	result := stringslice.NonWhitespaceJoinPtr([]string{}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns empty -- empty input", actual)
}

func Test_Cov8_NonWhitespaceJoinPtr_Items(t *testing.T) {
	result := stringslice.NonWhitespaceJoinPtr([]string{"a", "  ", "b"}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns joined -- whitespace removed", actual)
}

// ── NonEmptyJoinPtr ──

func Test_Cov8_NonEmptyJoinPtr_Empty(t *testing.T) {
	result := stringslice.NonEmptyJoinPtr([]string{}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns empty -- empty input", actual)
}

func Test_Cov8_NonEmptyJoinPtr_Items(t *testing.T) {
	result := stringslice.NonEmptyJoinPtr([]string{"a", "", "b"}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr returns joined -- empty removed", actual)
}

// ── LastSafeIndexPtr ──

func Test_Cov8_LastSafeIndexPtr_Empty(t *testing.T) {
	result := stringslice.LastSafeIndexPtr([]string{})
	actual := args.Map{"val": result}
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns -1 -- empty input", actual)
}

func Test_Cov8_LastSafeIndexPtr_Items(t *testing.T) {
	result := stringslice.LastSafeIndexPtr([]string{"a", "b"})
	actual := args.Map{"val": result}
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns 1 -- two items", actual)
}

// ── SafeIndexAtUsingLastIndex ──

func Test_Cov8_SafeIndexAtUsingLastIndex_OutOfRange(t *testing.T) {
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- index out of range", actual)
}

func Test_Cov8_SafeIndexAtUsingLastIndex_NegativeIndex(t *testing.T) {
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, -1)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- negative index", actual)
}

func Test_Cov8_SafeIndexAtUsingLastIndex_ZeroLastIndex(t *testing.T) {
	result := stringslice.SafeIndexAtUsingLastIndex([]string{"a"}, 0, 0)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns empty -- lastIndex is 0", actual)
}

// ── CloneSimpleSliceToPointers ──

func Test_Cov8_CloneSimpleSliceToPointers_Empty(t *testing.T) {
	result := stringslice.CloneSimpleSliceToPointers([]string{})
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns empty ptr -- empty input", actual)
}

func Test_Cov8_CloneSimpleSliceToPointers_Items(t *testing.T) {
	result := stringslice.CloneSimpleSliceToPointers([]string{"a", "b"})
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns ptr -- two items", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_Cov8_RegexTrimmedSplitNonEmptyAll(t *testing.T) {
	re := regexp.MustCompile(`[,;]`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, " a , b ; c ")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "a"}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns trimmed -- regex split", actual)
}

// ── ExpandBySplits — uncovered branches ──

func Test_Cov8_ExpandBySplits_Empty(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- empty input", actual)
}

func Test_Cov8_ExpandBySplits_NoSplitters(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a,b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- no splitters", actual)
}

func Test_Cov8_ExpandBySplits_Items(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a,b", "c;d"}, ",", ";")
	actual := args.Map{"hasItems": len(result) > 2}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns expanded -- multiple splitters", actual)
}

// ── LinesSimpleProcess — uncovered branches ──

func Test_Cov8_LinesSimpleProcess_Empty(t *testing.T) {
	result := stringslice.LinesSimpleProcess([]string{}, func(s string) string { return s })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns empty -- empty input", actual)
}

func Test_Cov8_LinesSimpleProcess_Items(t *testing.T) {
	result := stringslice.LinesSimpleProcess([]string{"a", "b"}, func(s string) string { return s + "!" })
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "a!", "second": "b!"}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns processed -- append exclamation", actual)
}

// ── LinesSimpleProcessNoEmpty — uncovered branches ──

func Test_Cov8_LinesSimpleProcessNoEmpty_Empty(t *testing.T) {
	result := stringslice.LinesSimpleProcessNoEmpty([]string{}, func(s string) string { return s })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns empty -- empty input", actual)
}

func Test_Cov8_LinesSimpleProcessNoEmpty_SkipsEmpty(t *testing.T) {
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "", "b"}, func(s string) string { return s })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty returns filtered -- empty removed", actual)
}

// ── MergeSlicesOfSlices — uncovered branches ──

func Test_Cov8_MergeSlicesOfSlices_Empty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- no input", actual)
}

func Test_Cov8_MergeSlicesOfSlices_AllEmpty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{}, []string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- all empty slices", actual)
}

func Test_Cov8_MergeSlicesOfSlices_Mixed(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{}, []string{"b", "c"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns merged -- mixed slices", actual)
}

// ── AnyItemsCloneIf — uncovered branches ──

func Test_Cov8_AnyItemsCloneIf_NilNoClone(t *testing.T) {
	result := stringslice.AnyItemsCloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns empty -- nil no clone", actual)
}

func Test_Cov8_AnyItemsCloneIf_NoClone(t *testing.T) {
	input := []any{"a", "b"}
	result := stringslice.AnyItemsCloneIf(false, 0, input)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns original -- no clone", actual)
}

func Test_Cov8_AnyItemsCloneIf_Clone(t *testing.T) {
	input := []any{"a"}
	result := stringslice.AnyItemsCloneIf(true, 2, input)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns cloned -- clone mode", actual)
}

// ── AnyItemsCloneUsingCap — uncovered branches ──

func Test_Cov8_AnyItemsCloneUsingCap_Empty(t *testing.T) {
	result := stringslice.AnyItemsCloneUsingCap(5, []any{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns empty -- empty input", actual)
}

func Test_Cov8_AnyItemsCloneUsingCap_Items(t *testing.T) {
	result := stringslice.AnyItemsCloneUsingCap(2, []any{"x", "y"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneUsingCap returns cloned -- two items", actual)
}

// ── ProcessAsync ──

func Test_Cov8_ProcessAsync_Empty(t *testing.T) {
	result := stringslice.ProcessAsync(func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns empty -- no items", actual)
}

func Test_Cov8_ProcessAsync_Items(t *testing.T) {
	result := stringslice.ProcessAsync(func(i int, item any) string {
		return item.(string) + "!"
	}, "a", "b")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a!"}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns processed -- two items", actual)
}

// ── ProcessOptionAsync ──

func Test_Cov8_ProcessOptionAsync_Empty(t *testing.T) {
	result := stringslice.ProcessOptionAsync(false, func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- no items", actual)
}

func Test_Cov8_ProcessOptionAsync_SkipNil(t *testing.T) {
	result := stringslice.ProcessOptionAsync(true, func(i int, item any) string {
		if item == "skip" {
			return ""
		}
		return item.(string)
	}, "a", "skip", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns filtered -- skip empty", actual)
}

func Test_Cov8_ProcessOptionAsync_ReturnAll(t *testing.T) {
	result := stringslice.ProcessOptionAsync(false, func(i int, item any) string {
		if item == "skip" {
			return ""
		}
		return item.(string)
	}, "a", "skip", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns all -- no skip", actual)
}

// ── LinesAsyncProcess ──

func Test_Cov8_LinesAsyncProcess_Empty(t *testing.T) {
	result := stringslice.LinesAsyncProcess([]string{}, func(i int, s string) string { return s })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns empty -- empty input", actual)
}

func Test_Cov8_LinesAsyncProcess_Items(t *testing.T) {
	result := stringslice.LinesAsyncProcess([]string{"a", "b"}, func(i int, s string) string { return s + "!" })
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a!"}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns processed -- two items", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor ──

func Test_Cov8_AnyLinesProcessAsync_Nil(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(nil, func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- nil input", actual)
}

func Test_Cov8_AnyLinesProcessAsync_NotSlice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor("notslice", func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- non-slice input", actual)
}

func Test_Cov8_AnyLinesProcessAsync_EmptySlice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]int{}, func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- empty slice", actual)
}

func Test_Cov8_AnyLinesProcessAsync_Items(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]string{"x", "y"}, func(i int, item any) string {
		return item.(string) + "!"
	})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "x!"}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns processed -- string slice", actual)
}

// ── AppendAnyItemsWithStrings — uncovered branches ──

func Test_Cov8_AppendAnyItemsWithStrings_NilItem(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{}, nil, "hello")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings skips nil -- nil and string items", actual)
}

func Test_Cov8_AppendAnyItemsWithStrings_SkipEmpty(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(false, true, []string{}, "")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns empty -- skip empty string", actual)
}

// ── AppendStringsWithAnyItems — uncovered branches ──

func Test_Cov8_AppendStringsWithAnyItems_SkipEmpty(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(false, true, []any{}, "hello", "")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns filtered -- skip empty", actual)
}

func Test_Cov8_AppendStringsWithAnyItems_NoItems(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"existing"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns original -- no appending items", actual)
}
