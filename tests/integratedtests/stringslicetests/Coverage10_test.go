package stringslicetests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov10_FirstLastDefaultStatus_Multi(t *testing.T) {
	result := stringslice.FirstLastDefaultStatus([]string{"a", "b", "c"})
	actual := args.Map{"isValid": result.IsValid, "hasFirst": result.HasFirst, "hasLast": result.HasLast, "first": result.First, "last": result.Last}
	expected := args.Map{"isValid": true, "hasFirst": true, "hasLast": true, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns valid -- multiple elems", actual)
}

// ── InPlaceReverse — all branches ──

func Test_Cov10_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns empty -- nil", actual)
}

func Test_Cov10_InPlaceReverse_Single(t *testing.T) {
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"len": len(*result), "first": (*result)[0]}
	expected := args.Map{"len": 1, "first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns same -- single", actual)
}

func Test_Cov10_InPlaceReverse_Two(t *testing.T) {
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[1]}
	expected := args.Map{"first": "b", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse swaps -- two elements", actual)
}

func Test_Cov10_InPlaceReverse_Three(t *testing.T) {
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "mid": (*result)[1], "last": (*result)[2]}
	expected := args.Map{"first": "c", "mid": "b", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- three elements", actual)
}

func Test_Cov10_InPlaceReverse_Four(t *testing.T) {
	s := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[3]}
	expected := args.Map{"first": "d", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- four elements (even)", actual)
}

// ── IndexesDefault — empty slice ──

func Test_Cov10_IndexesDefault_EmptySlice(t *testing.T) {
	result := stringslice.IndexesDefault(nil, 0)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns empty -- nil slice", actual)
}

func Test_Cov10_IndexesDefault_NoIndexes(t *testing.T) {
	result := stringslice.IndexesDefault([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns empty -- no indexes", actual)
}

// ── SafeIndexesDefaultWithDetail — missing indexes ──

func Test_Cov10_SafeIndexesDefaultWithDetail_SomeMissing(t *testing.T) {
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b"}, 0, 5, -1)
	actual := args.Map{
		"valuesLen":  len(result.Values),
		"missingLen": len(result.MissingIndexes),
		"anyMissing": result.IsAnyMissing,
		"isValid":    result.IsValid,
	}
	expected := args.Map{"valuesLen": 1, "missingLen": 2, "anyMissing": true, "isValid": true}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail reports missing -- some OOB", actual)
}

func Test_Cov10_SafeIndexesDefaultWithDetail_Empty(t *testing.T) {
	result := stringslice.SafeIndexesDefaultWithDetail(nil, 0)
	actual := args.Map{"isValid": result.IsValid, "anyMissing": result.IsAnyMissing}
	expected := args.Map{"isValid": false, "anyMissing": true}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns invalid -- nil", actual)
}

// ── SafeIndexRanges — negative requestLength ──

func Test_Cov10_SafeIndexRanges_NegativeRange(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, 3, 1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns empty -- start > end", actual)
}

func Test_Cov10_SafeIndexRanges_OOBIndexes(t *testing.T) {
	result := stringslice.SafeIndexRanges([]string{"a", "b"}, -1, 5)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns padded -- OOB handled with empty strings", actual)
}

// ── SplitTrimmedNonEmpty ──

func Test_Cov10_SplitTrimmedNonEmpty_Basic(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty(" a , b , c ", ",", -1)
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns trimmed -- basic", actual)
}

func Test_Cov10_SplitTrimmedNonEmpty_Limited(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", 2)
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b,c"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns limited -- n=2", actual)
}

// ── SplitTrimmedNonEmptyAll ──

func Test_Cov10_SplitTrimmedNonEmptyAll_Basic(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll(" x | y | z ", "|")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "x", "last": "z"}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll returns trimmed -- basic", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_Cov10_RegexTrimmedSplitNonEmptyAll(t *testing.T) {
	re := regexp.MustCompile(`[,;]+`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, " a , b ;; c ")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns trimmed -- regex split", actual)
}

func Test_Cov10_RegexTrimmedSplitNonEmptyAll_AllEmpty(t *testing.T) {
	re := regexp.MustCompile(`.`)
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "abc")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns empty -- all split to empty", actual)
}

// ── ExpandByFunc — with empty expansion ──

func Test_Cov10_ExpandByFunc_SomeEmpty(t *testing.T) {
	result := stringslice.ExpandByFunc([]string{"a", "b"}, func(line string) []string {
		if line == "a" {
			return []string{"a1", "a2"}
		}
		return nil
	})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a1", "last": "a2"}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns nil -- skips nil expansions", actual)
}

func Test_Cov10_ExpandByFunc_Empty(t *testing.T) {
	result := stringslice.ExpandByFunc(nil, func(line string) []string { return nil })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns empty -- nil input", actual)
}

// ── ExpandBySplit ──

func Test_Cov10_ExpandBySplit_Basic(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{"a,b", "c,d"}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns expanded -- basic", actual)
}

func Test_Cov10_ExpandBySplit_Empty(t *testing.T) {
	result := stringslice.ExpandBySplit(nil, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- nil", actual)
}

// ── ExpandBySplits — multiple splitters, no splitters ──

func Test_Cov10_ExpandBySplits_MultipleSplitters(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a,b;c"}, ",", ";")
	actual := args.Map{"len": len(result)}
	// "a,b;c" split by "," => ["a", "b;c"], then "a,b;c" split by ";" => ["a,b", "c"]
	// total = 4
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns expanded -- multiple splitters", actual)
}

func Test_Cov10_ExpandBySplits_NoSplitters(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- no splitters", actual)
}

// ── AppendStringsWithMainSlice — all branches ──

func Test_Cov10_AppendStringsWithMainSlice_NoItems(t *testing.T) {
	input := []string{"a"}
	result := stringslice.AppendStringsWithMainSlice(false, input)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns same -- no items", actual)
}

func Test_Cov10_AppendStringsWithMainSlice_SkipEmpty(t *testing.T) {
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"}, "b", "", "c")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice skips empty -- isSkipEmpty true", actual)
}

func Test_Cov10_AppendStringsWithMainSlice_IncludeEmpty(t *testing.T) {
	result := stringslice.AppendStringsWithMainSlice(false, []string{"a"}, "b", "", "c")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice includes empty -- isSkipEmpty false", actual)
}

// ── AppendStringsWithAnyItems — clone=true branches ──

func Test_Cov10_AppendStringsWithAnyItems_Clone(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(true, false, []any{"x"}, "a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems clones and appends -- clone true", actual)
}

func Test_Cov10_AppendStringsWithAnyItems_NoAppend(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(false, false, []any{"x"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns same -- no items to append", actual)
}

// ── AppendAnyItemsWithStrings — nil item ──

func Test_Cov10_AppendAnyItemsWithStrings_NilItem(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"}, nil, "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings skips nil -- nil item in list", actual)
}

func Test_Cov10_AppendAnyItemsWithStrings_Clone(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(true, false, []string{"a"}, "b")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings clones -- clone true", actual)
}

// ── CloneSimpleSliceToPointers — non-empty ──

func Test_Cov10_CloneSimpleSliceToPointers_NonEmpty(t *testing.T) {
	result := stringslice.CloneSimpleSliceToPointers([]string{"a", "b"})
	actual := args.Map{"len": len(*result), "first": (*result)[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns cloned ptr -- non-empty", actual)
}

// ── LinesProcess — break mid-iteration ──

func Test_Cov10_LinesProcess_BreakAfterTake(t *testing.T) {
	result := stringslice.LinesProcess([]string{"a", "b", "c"}, func(i int, s string) (string, bool, bool) {
		return s, true, i == 1
	})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "LinesProcess stops early -- break after second take", actual)
}

// ── LinesSimpleProcess — identity ──

func Test_Cov10_LinesSimpleProcess_Identity(t *testing.T) {
	result := stringslice.LinesSimpleProcess([]string{"a", "b"}, func(s string) string { return s })
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns identity -- passthrough", actual)
}

func Test_Cov10_LinesSimpleProcess_Empty(t *testing.T) {
	result := stringslice.LinesSimpleProcess(nil, func(s string) string { return s })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess returns empty -- nil", actual)
}

// ── LinesSimpleProcessNoEmpty — mixed ──

func Test_Cov10_LinesSimpleProcessNoEmpty_Mixed(t *testing.T) {
	result := stringslice.LinesSimpleProcessNoEmpty([]string{"a", "b", "c"}, func(s string) string {
		if s == "b" {
			return ""
		}
		return s + "!"
	})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a!", "last": "c!"}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty filters empties -- mixed", actual)
}

// ── LinesAsyncProcess — multiple items ──

func Test_Cov10_LinesAsyncProcess_Multi(t *testing.T) {
	result := stringslice.LinesAsyncProcess([]string{"a", "b", "c"}, func(i int, s string) string {
		return fmt.Sprintf("%d:%s", i, s)
	})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "0:a", "last": "2:c"}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess processes all -- multiple items", actual)
}

func Test_Cov10_LinesAsyncProcess_Empty(t *testing.T) {
	result := stringslice.LinesAsyncProcess(nil, func(i int, s string) string { return s })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LinesAsyncProcess returns empty -- nil", actual)
}

// ── AnyLinesProcessAsyncUsingProcessor — not-slice input, nil, empty ──

func Test_Cov10_AnyLinesProcessAsync_NotSlice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor("not-a-slice", func(i int, item any) string {
		return ""
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- non-slice input", actual)
}

func Test_Cov10_AnyLinesProcessAsync_Nil(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(nil, func(i int, item any) string {
		return ""
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- nil", actual)
}

func Test_Cov10_AnyLinesProcessAsync_EmptySlice(t *testing.T) {
	result := stringslice.AnyLinesProcessAsyncUsingProcessor([]int{}, func(i int, item any) string {
		return ""
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyLinesProcessAsync returns empty -- empty slice", actual)
}

// ── ProcessAsync — empty ──

func Test_Cov10_ProcessAsync_Empty(t *testing.T) {
	result := stringslice.ProcessAsync(func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessAsync returns empty -- no items", actual)
}

// ── ProcessOptionAsync — isSkipOnNil false ──

func Test_Cov10_ProcessOptionAsync_NoSkip(t *testing.T) {
	result := stringslice.ProcessOptionAsync(false, func(i int, item any) string {
		return ""
	}, "a", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns all including empty -- isSkipOnNil false", actual)
}

func Test_Cov10_ProcessOptionAsync_Empty(t *testing.T) {
	result := stringslice.ProcessOptionAsync(true, func(i int, item any) string { return "" })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ProcessOptionAsync returns empty -- no items", actual)
}

// ── NonEmptyJoin — all empty strings ──

func Test_Cov10_NonEmptyJoin_AllEmpty(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"", "", ""}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- all empty strings", actual)
}

func Test_Cov10_NonEmptyJoin_Nil(t *testing.T) {
	result := stringslice.NonEmptyJoin(nil, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- nil", actual)
}

func Test_Cov10_NonEmptyJoin_Mixed(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns joined -- mixed", actual)
}

// ── NonWhitespaceJoin ──

func Test_Cov10_NonWhitespaceJoin_Nil(t *testing.T) {
	result := stringslice.NonWhitespaceJoin(nil, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- nil", actual)
}

func Test_Cov10_NonWhitespaceJoin_Empty(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- empty slice", actual)
}

func Test_Cov10_NonWhitespaceJoin_Mixed(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{"a", " ", "b"}, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns joined -- mixed", actual)
}

// ── MergeNewSimple ──

func Test_Cov10_MergeNewSimple_Empty(t *testing.T) {
	result := stringslice.MergeNewSimple()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns empty -- no args", actual)
}

func Test_Cov10_MergeNewSimple_WithEmpty(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, nil, []string{"b"})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple skips nil -- mixed", actual)
}

// ── SortIf ──

func Test_Cov10_SortIf_True(t *testing.T) {
	result := stringslice.SortIf(true, []string{"c", "a", "b"})
	actual := args.Map{"first": result[0], "last": result[2]}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf sorts -- isSort true", actual)
}

func Test_Cov10_SortIf_False(t *testing.T) {
	result := stringslice.SortIf(false, []string{"c", "a", "b"})
	actual := args.Map{"first": result[0], "last": result[2]}
	expected := args.Map{"first": "c", "last": "b"}
	expected.ShouldBeEqual(t, 0, "SortIf no-op -- isSort false", actual)
}

// ── JoinWith / Joins — empty ──

func Test_Cov10_JoinWith_Empty(t *testing.T) {
	result := stringslice.JoinWith(",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- no items", actual)
}

func Test_Cov10_Joins_Empty(t *testing.T) {
	result := stringslice.Joins(",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- no items", actual)
}

func Test_Cov10_Joins_Multi(t *testing.T) {
	result := stringslice.Joins(",", "a", "b", "c")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a,b,c"}
	expected.ShouldBeEqual(t, 0, "Joins returns joined -- multiple items", actual)
}

func Test_Cov10_JoinWith_Multi(t *testing.T) {
	result := stringslice.JoinWith(",", "a", "b")
	actual := args.Map{"val": result}
	expected := args.Map{"val": ",a,b"}
	expected.ShouldBeEqual(t, 0, "JoinWith returns prepended join -- multiple items", actual)
}

// ── CloneIf — non-clone branches ──

func Test_Cov10_CloneIf_NilNoClone(t *testing.T) {
	result := stringslice.CloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf returns empty -- nil no clone", actual)
}

func Test_Cov10_CloneIf_NonNilNoClone(t *testing.T) {
	input := []string{"a", "b"}
	result := stringslice.CloneIf(false, 0, input)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "CloneIf returns original -- non-nil no clone", actual)
}

// ── AnyItemsCloneIf — non-clone branches ──

func Test_Cov10_AnyItemsCloneIf_NilNoClone(t *testing.T) {
	result := stringslice.AnyItemsCloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns empty -- nil no clone", actual)
}

func Test_Cov10_AnyItemsCloneIf_NonNilNoClone(t *testing.T) {
	input := []any{"a", "b"}
	result := stringslice.AnyItemsCloneIf(false, 0, input)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns original -- non-nil no clone", actual)
}

func Test_Cov10_AnyItemsCloneIf_NonNilClone(t *testing.T) {
	input := []any{"a", "b"}
	result := stringslice.AnyItemsCloneIf(true, 3, input)
	actual := args.Map{"len": len(result), "capGe5": cap(result) >= 5}
	expected := args.Map{"len": 2, "capGe5": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf clones with cap -- clone true non-nil", actual)
}

// ── SafeIndexAt — all branches ──

func Test_Cov10_SafeIndexAt_Valid(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a", "b"}, 1)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns element -- valid", actual)
}

func Test_Cov10_SafeIndexAt_NegIndex(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a"}, -1)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- negative index", actual)
}

func Test_Cov10_SafeIndexAt_OOB(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a"}, 5)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- OOB", actual)
}

func Test_Cov10_SafeIndexAt_Empty(t *testing.T) {
	result := stringslice.SafeIndexAt(nil, 0)
	actual := args.Map{"val": result}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- nil", actual)
}

// ── SafeIndexAtWith — all branches ──

func Test_Cov10_SafeIndexAtWith_Valid(t *testing.T) {
	result := stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns element -- valid", actual)
}

func Test_Cov10_SafeIndexAtWith_NegIndex(t *testing.T) {
	result := stringslice.SafeIndexAtWith([]string{"a"}, -1, "def")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns default -- negative", actual)
}

// ── NonWhitespace — nil ──

func Test_Cov10_NonWhitespace_Nil(t *testing.T) {
	result := stringslice.NonWhitespace(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- nil", actual)
}

func Test_Cov10_NonWhitespace_EmptySlice(t *testing.T) {
	result := stringslice.NonWhitespace([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- empty slice", actual)
}

// ── NonEmptyStrings — nil ──

func Test_Cov10_NonEmptyStrings_Nil(t *testing.T) {
	result := stringslice.NonEmptyStrings(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- nil", actual)
}

func Test_Cov10_NonEmptyStrings_EmptySlice(t *testing.T) {
	result := stringslice.NonEmptyStrings([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- empty slice", actual)
}

// ── AllElemLengthSlices — nil among slices, no args ──

func Test_Cov10_AllElemLengthSlices_NoArgs(t *testing.T) {
	result := stringslice.AllElemLengthSlices()
	actual := args.Map{"val": result}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns 0 -- no args", actual)
}

func Test_Cov10_AllElemLengthSlices_NilAmong(t *testing.T) {
	result := stringslice.AllElemLengthSlices([]string{"a"}, nil, []string{"b", "c"})
	actual := args.Map{"val": result}
	expected := args.Map{"val": 3}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices skips nil -- mixed", actual)
}

// ── PrependLineNew ──

func Test_Cov10_PrependLineNew(t *testing.T) {
	result := stringslice.PrependLineNew("first", []string{"a", "b"})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "first", "last": "b"}
	expected.ShouldBeEqual(t, 0, "PrependLineNew returns correct value -- prepends single line", actual)
}

// ── AppendLineNew ──

func Test_Cov10_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew returns correct value -- appends single line", actual)
}

// ── Clone ──

func Test_Cov10_Clone_NonEmpty(t *testing.T) {
	result := stringslice.Clone([]string{"a", "b"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "Clone returns deep copy -- non-empty", actual)
}

func Test_Cov10_Clone_Empty(t *testing.T) {
	result := stringslice.Clone(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- nil", actual)
}

// ── Simple accessors: First, Last, IndexAt, Empty, IsEmpty, HasAnyItem ──

func Test_Cov10_First(t *testing.T) {
	result := stringslice.First([]string{"x", "y"})
	actual := args.Map{"val": result}
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "First returns correct value -- returns first element", actual)
}

func Test_Cov10_Last(t *testing.T) {
	result := stringslice.Last([]string{"x", "y"})
	actual := args.Map{"val": result}
	expected := args.Map{"val": "y"}
	expected.ShouldBeEqual(t, 0, "Last returns correct value -- returns last element", actual)
}

func Test_Cov10_IndexAt(t *testing.T) {
	result := stringslice.IndexAt([]string{"a", "b", "c"}, 2)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "c"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns correct value -- returns element at index", actual)
}

func Test_Cov10_Empty(t *testing.T) {
	result := stringslice.Empty()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- returns empty slice", actual)
}

func Test_Cov10_IsEmpty_True(t *testing.T) {
	actual := args.Map{"val": stringslice.IsEmpty(nil)}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty true -- nil", actual)
}

func Test_Cov10_IsEmpty_False(t *testing.T) {
	actual := args.Map{"val": stringslice.IsEmpty([]string{"a"})}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty false -- has items", actual)
}

func Test_Cov10_HasAnyItem_True(t *testing.T) {
	actual := args.Map{"val": stringslice.HasAnyItem([]string{"a"})}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem true -- has items", actual)
}

func Test_Cov10_HasAnyItem_False(t *testing.T) {
	actual := args.Map{"val": stringslice.HasAnyItem(nil)}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem false -- nil", actual)
}

// ── FirstOrDefault ──

func Test_Cov10_FirstOrDefault_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefault(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- nil", actual)
}

func Test_Cov10_FirstOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{"a"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns first -- non-empty", actual)
}

// ── LastOrDefault ──

func Test_Cov10_LastOrDefault_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefault(nil)}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- nil", actual)
}

func Test_Cov10_LastOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefault([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns last -- non-empty", actual)
}

// ── MakeDefault / MakeLen / Make ──

func Test_Cov10_MakeDefault(t *testing.T) {
	result := stringslice.MakeDefault(10)
	actual := args.Map{"len": len(result), "capGe10": cap(result) >= 10}
	expected := args.Map{"len": 0, "capGe10": true}
	expected.ShouldBeEqual(t, 0, "MakeDefault returns non-empty -- returns zero-len with capacity", actual)
}

func Test_Cov10_MakeLen(t *testing.T) {
	result := stringslice.MakeLen(3)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLen returns non-empty -- returns slice with length", actual)
}

func Test_Cov10_Make(t *testing.T) {
	result := stringslice.Make(2, 5)
	actual := args.Map{"len": len(result), "capGe5": cap(result) >= 5}
	expected := args.Map{"len": 2, "capGe5": true}
	expected.ShouldBeEqual(t, 0, "Make returns non-empty -- returns slice with length and capacity", actual)
}

// ── SafeIndexes — OOB and negative ──

func Test_Cov10_SafeIndexes_WithOOB(t *testing.T) {
	result := stringslice.SafeIndexes([]string{"a", "b"}, 0, 5, -1, 1)
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1], "third": result[2], "fourth": result[3]}
	expected := args.Map{"len": 4, "first": "a", "second": "", "third": "", "fourth": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns empty -- handles OOB and negative with empty strings", actual)
}

func Test_Cov10_SafeIndexes_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexes(nil, 0)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns empty -- returns default for empty slice", actual)
}

// ── NonEmptySlice — with empty strings ──

func Test_Cov10_NonEmptySlice_MixedEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "", "b", ""})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice filters empties -- mixed", actual)
}

func Test_Cov10_NonEmptySlice_Empty(t *testing.T) {
	result := stringslice.NonEmptySlice(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- nil", actual)
}

// ── NonNullStrings — nil ──

func Test_Cov10_NonNullStrings_Nil(t *testing.T) {
	result := stringslice.NonNullStrings(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns empty -- nil", actual)
}

// ── SplitContentsByWhitespace — with content ──

func Test_Cov10_SplitContentsByWhitespace_Multi(t *testing.T) {
	result := stringslice.SplitContentsByWhitespace("  a  b  c  ")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace returns fields -- whitespace separated", actual)
}
