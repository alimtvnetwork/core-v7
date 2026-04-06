package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Clone / ClonePtr / CloneUsingCap ──

func Test_Cov5_Clone(t *testing.T) {
	result := stringslice.Clone([]string{"a", "b"})
	nilResult := stringslice.Clone(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns correct -- 2 items", actual)
}
}
func Test_Cov5_FirstOrDefault(t *testing.T) {
	actual := args.Map{"found": stringslice.FirstOrDefault([]string{"a"}), "empty": stringslice.FirstOrDefault(nil)}
	expected := args.Map{"found": "a", "empty": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns correct -- found and empty", actual)
}
func Test_Cov5_Last(t *testing.T) {
	actual := args.Map{"val": stringslice.Last([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Last returns last -- 2 items", actual)
}
func Test_Cov5_LastOrDefault(t *testing.T) {
	actual := args.Map{"found": stringslice.LastOrDefault([]string{"a", "b"}), "empty": stringslice.LastOrDefault(nil)}
	expected := args.Map{"found": "b", "empty": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns correct -- found and empty", actual)
}

// ── FirstLastDefault / FirstLastDefaultStatus ──

func Test_Cov5_FirstLastDefault_Single(t *testing.T) {
	f, l := stringslice.FirstLastDefault([]string{"a"})
	actual := args.Map{"first": f, "last": l}
	expected := args.Map{"first": "a", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault single -- first only", actual)
}

func Test_Cov5_FirstLastDefaultStatus_Single(t *testing.T) {
	s := stringslice.FirstLastDefaultStatus([]string{"a"})
	actual := args.Map{"first": s.First, "hasFirst": s.HasFirst, "hasLast": s.HasLast, "isValid": s.IsValid}
	expected := args.Map{"first": "a", "hasFirst": true, "hasLast": false, "isValid": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus single -- first only", actual)
}
func Test_Cov5_HasAnyItem(t *testing.T) {
	actual := args.Map{"yes": stringslice.HasAnyItem([]string{"a"}), "no": stringslice.HasAnyItem(nil)}
	expected := args.Map{"yes": true, "no": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns correct value -- returns correct", actual)
}
	expected.ShouldBeEqual(t, 0, "IndexAt returns correct -- index 1", actual)
}

func Test_Cov5_SafeIndexAt(t *testing.T) {
	actual := args.Map{
		"valid":   stringslice.SafeIndexAt([]string{"a", "b"}, 1),
		"invalid": stringslice.SafeIndexAt([]string{"a"}, 5),
		"neg":     stringslice.SafeIndexAt([]string{"a"}, -1),
		"empty":   stringslice.SafeIndexAt(nil, 0),
	}
	expected := args.Map{"valid": "b", "invalid": "", "neg": "", "empty": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns correct -- all branches", actual)
}
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns 2 -- skip empty", actual)
}

func Test_Cov5_NonEmptyIf(t *testing.T) {
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})
	result2 := stringslice.NonEmptyIf(false, []string{"a", "", "b"})
	actual := args.Map{"trueLen": len(result), "falseLen": len(result2)}
	expected := args.Map{"trueLen": 2, "falseLen": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf returns correct -- true and false", actual)
}

func Test_Cov5_NonEmptyStrings(t *testing.T) {
	result := stringslice.NonEmptyStrings([]string{"a", "", "b"})
	nilResult := stringslice.NonEmptyStrings(nil)
	emptyResult := stringslice.NonEmptyStrings([]string{})
	actual := args.Map{"len": len(result), "nilLen": len(nilResult), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "nilLen": 0, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns correct -- all branches", actual)
}

func Test_Cov5_NonNullStrings(t *testing.T) {
	result := stringslice.NonNullStrings([]string{"a", "", "b"})
	nilResult := stringslice.NonNullStrings(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns correct value -- returns correct", actual)
}

// ── NonWhitespace / NonWhitespaceJoin ──

func Test_Cov5_NonWhitespace(t *testing.T) {
	result := stringslice.NonWhitespace([]string{"a", "  ", "", "b"})
	nilResult := stringslice.NonWhitespace(nil)
	emptyResult := stringslice.NonWhitespace([]string{})
	actual := args.Map{"len": len(result), "nilLen": len(nilResult), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "nilLen": 0, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns correct -- all branches", actual)
}

func Test_Cov5_NonWhitespaceJoin(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ",")
	nilResult := stringslice.NonWhitespaceJoin(nil, ",")
	emptyResult := stringslice.NonWhitespaceJoin([]string{}, ",")
	actual := args.Map{"val": result, "nilVal": nilResult, "emptyVal": emptyResult}
	expected := args.Map{"val": "a,b", "nilVal": "", "emptyVal": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns correct value -- returns correct", actual)
}

func Test_Cov5_NonEmptyJoin(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	nilResult := stringslice.NonEmptyJoin(nil, ",")
	emptyResult := stringslice.NonEmptyJoin([]string{}, ",")
	actual := args.Map{"val": result, "nilVal": nilResult, "emptyVal": emptyResult}
	expected := args.Map{"val": "a,b", "nilVal": "", "emptyVal": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- returns correct", actual)
}

// ── InPlaceReverse ──

func Test_Cov5_InPlaceReverse(t *testing.T) {
	items := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&items)
	two := []string{"x", "y"}
	result2 := stringslice.InPlaceReverse(&two)
	single := []string{"z"}
	result3 := stringslice.InPlaceReverse(&single)
	result4 := stringslice.InPlaceReverse(nil)
	actual := args.Map{
		"first": (*result)[0], "last": (*result)[2],
		"twoFirst": (*result2)[0], "singleFirst": (*result3)[0],
		"nilLen": len(*result4),
	}
	expected := args.Map{"first": "c", "last": "a", "twoFirst": "y", "singleFirst": "z", "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns correct -- all branches", actual)
}

// ── TrimmedEachWords / TrimmedEachWordsIf ──

func Test_Cov5_TrimmedEachWords(t *testing.T) {
	result := stringslice.TrimmedEachWords([]string{"  a  ", "  ", " b "})
	nilResult := stringslice.TrimmedEachWords(nil)
	emptyResult := stringslice.TrimmedEachWords([]string{})
	actual := args.Map{"len": len(result), "first": result[0], "nilNil": nilResult == nil, "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "first": "a", "nilNil": true, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns correct value -- returns correct", actual)
}

func Test_Cov5_TrimmedEachWordsIf(t *testing.T) {
	result := stringslice.TrimmedEachWordsIf(true, []string{"  a  ", " "})
	result2 := stringslice.TrimmedEachWordsIf(false, []string{"  a  ", " "})
	actual := args.Map{"trueLen": len(result), "falseLen": len(result2)}
	expected := args.Map{"trueLen": 1, "falseLen": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf returns correct value -- returns correct", actual)
}

// ── MergeNew / MergeNewSimple / MergeSlicesOfSlices / PrependNew / AppendLineNew ──

func Test_Cov5_MergeNew(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"}, "b", "c")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeNew returns 3 -- 1+2", actual)
}

func Test_Cov5_MergeNewSimple(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"}, nil)
	emptyResult := stringslice.MergeNewSimple()
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 3, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns correct value -- returns correct", actual)
}

func Test_Cov5_MergeSlicesOfSlices(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, nil, []string{"b"})
	emptyResult := stringslice.MergeSlicesOfSlices()
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns correct value -- returns correct", actual)
}

func Test_Cov5_PrependNew(t *testing.T) {
	result := stringslice.PrependNew([]string{"b"}, "a")
	actual := args.Map{"first": (*result)[0], "len": len(*result)}
	expected := args.Map{"first": "a", "len": 2}
	expected.ShouldBeEqual(t, 0, "PrependNew returns correct -- prepend a", actual)
}

func Test_Cov5_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result), "last": result[1]}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew returns correct -- append b", actual)
}

// ── AppendStringsWithMainSlice ──

func Test_Cov5_AppendStringsWithMainSlice(t *testing.T) {
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"}, "b", "", "c")
	noSkip := stringslice.AppendStringsWithMainSlice(false, []string{"a"}, "b", "", "c")
	noAppend := stringslice.AppendStringsWithMainSlice(false, []string{"a"})
	actual := args.Map{"skipLen": len(result), "noSkipLen": len(noSkip), "noAppendLen": len(noAppend)}
	expected := args.Map{"skipLen": 3, "noSkipLen": 4, "noAppendLen": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice returns non-empty -- returns correct", actual)
}

// ── AppendStringsWithAnyItems / AppendAnyItemsWithStrings ──

func Test_Cov5_AppendStringsWithAnyItems_Func(t *testing.T) {
	result := stringslice.AppendStringsWithAnyItems(true, true, []any{1}, "a", "", "b")
	noAppend := stringslice.AppendStringsWithAnyItems(false, false, []any{1})
	actual := args.Map{"len": len(result), "noAppendLen": len(noAppend)}
	expected := args.Map{"len": 3, "noAppendLen": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns non-empty -- returns correct", actual)
}

func Test_Cov5_AppendAnyItemsWithStrings_Func(t *testing.T) {
	result := stringslice.AppendAnyItemsWithStrings(true, true, []string{"a"}, 1, nil, 2)
	noAppend := stringslice.AppendAnyItemsWithStrings(false, false, []string{"a"})
	actual := args.Map{"len": len(result), "noAppendLen": len(noAppend)}
	expected := args.Map{"len": 3, "noAppendLen": 1}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns non-empty -- returns correct", actual)
}

// ── AnyItemsCloneIf / AnyItemsCloneUsingCap ──

func Test_Cov5_AnyItemsCloneIf(t *testing.T) {
	result := stringslice.AnyItemsCloneIf(true, 5, []any{1, 2})
	noClone := stringslice.AnyItemsCloneIf(false, 0, []any{1})
	nilNoClone := stringslice.AnyItemsCloneIf(false, 0, nil)
	actual := args.Map{"len": len(result), "noCloneLen": len(noClone), "nilLen": len(nilNoClone)}
	expected := args.Map{"len": 2, "noCloneLen": 1, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsCloneIf returns correct value -- returns correct", actual)
}

// ── AllElemLengthSlices ──

func Test_Cov5_AllElemLengthSlices(t *testing.T) {
	result := stringslice.AllElemLengthSlices([]string{"a"}, nil, []string{"b", "c"})
	emptyResult := stringslice.AllElemLengthSlices()
	actual := args.Map{"count": result, "emptyCount": emptyResult}
	expected := args.Map{"count": 3, "emptyCount": 0}
	expected.ShouldBeEqual(t, 0, "AllElemLengthSlices returns correct value -- returns correct", actual)
}

// ── SafeRangeItems ──

func Test_Cov5_SafeRangeItems(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a", "b", "c", "d"}, 1, 3)
	nilResult := stringslice.SafeRangeItems(nil, 0, 1)
	emptyResult := stringslice.SafeRangeItems([]string{}, 0, 1)
	outResult := stringslice.SafeRangeItems([]string{"a"}, 5, 6)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult), "emptyLen": len(emptyResult), "outLen": len(outResult)}
	expected := args.Map{"len": 2, "nilLen": 0, "emptyLen": 0, "outLen": 0}
	expected.ShouldBeEqual(t, 0, "SafeRangeItems returns correct value -- returns correct", actual)
}

// ── SafeIndexesDefaultWithDetail ──

func Test_Cov5_SafeIndexesDefaultWithDetail(t *testing.T) {
	result := stringslice.SafeIndexesDefaultWithDetail([]string{"a", "b", "c"}, 0, 2, 99)
	emptyResult := stringslice.SafeIndexesDefaultWithDetail(nil, 0)
	actual := args.Map{
		"valuesLen":  len(result.Values),
		"anyMissing": result.IsAnyMissing,
		"isValid":    result.IsValid,
		"emptyValid": emptyResult.IsValid,
	}
	expected := args.Map{"valuesLen": 2, "anyMissing": true, "isValid": true, "emptyValid": false}
	expected.ShouldBeEqual(t, 0, "SafeIndexesDefaultWithDetail returns non-empty -- returns correct", actual)
}

// ── IndexesDefault ──

func Test_Cov5_IndexesDefault(t *testing.T) {
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 0, 2)
	emptyResult := stringslice.IndexesDefault(nil, 0)
	noIndexResult := stringslice.IndexesDefault([]string{"a"})
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult), "noIdxLen": len(noIndexResult)}
	expected := args.Map{"len": 2, "emptyLen": 0, "noIdxLen": 0}
	expected.ShouldBeEqual(t, 0, "IndexesDefault returns correct value -- returns correct", actual)
}

// ── SplitTrimmedNonEmpty ──

func Test_Cov5_SplitTrimmedNonEmpty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("a, , b", ",", -1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty returns 2 -- skip empty", actual)
}

// ── RegexTrimmedSplitNonEmptyAll ──

func Test_Cov5_RegexTrimmedSplitNonEmptyAll(t *testing.T) {
	re := regexp.MustCompile("[,;]")
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "a, ;b; c")
	actual := args.Map{"gt0": len(result) > 0}
	expected := args.Map{"gt0": true}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll returns empty -- returns items", actual)
}

// ── ExpandByFunc / ExpandBySplit / ExpandBySplits ──

func Test_Cov5_ExpandByFunc(t *testing.T) {
	result := stringslice.ExpandByFunc([]string{"a,b", "c,d"}, func(s string) []string { return []string{s + "!"} })
	emptyResult := stringslice.ExpandByFunc(nil, nil)
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns correct value -- returns correct", actual)
}

func Test_Cov5_ExpandBySplit(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{"a,b", "c,d"}, ",")
	emptyResult := stringslice.ExpandBySplit(nil, ",")
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 4, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns correct value -- returns correct", actual)
}
}

// ── Joins / JoinWith (from CloneIf.go) ──

func Test_Cov5_JoinWith(t *testing.T) {
	result := stringslice.JoinWith(",", "a", "b")
	emptyResult := stringslice.JoinWith(",")
	actual := args.Map{"val": result, "empty": emptyResult}
	expected := args.Map{"val": ",a,b", "empty": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns non-empty -- returns correct", actual)
}

func Test_Cov5_Joins(t *testing.T) {
	result := stringslice.Joins(",", "a", "b")
	emptyResult := stringslice.Joins(",")
	actual := args.Map{"val": result, "empty": emptyResult}
	expected := args.Map{"val": "a,b", "empty": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns correct value -- returns correct", actual)
}

// ── SplitContentsByWhitespace ──

func Test_Cov5_SplitContentsByWhitespace(t *testing.T) {
	result := stringslice.SplitContentsByWhitespace("  hello  world  ")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "hello"}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace returns correct value -- returns correct", actual)
}

// ── PrependLineNew ──

func Test_Cov5_PrependLineNew(t *testing.T) {
	result := stringslice.PrependLineNew("first", []string{"second"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "first"}
	expected.ShouldBeEqual(t, 0, "PrependLineNew returns correct value -- returns correct", actual)
}
