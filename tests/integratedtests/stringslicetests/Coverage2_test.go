package stringslicetests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// NonEmptySlice
// ============================================================================

func Test_Cov2_NonEmptySlice_Mixed(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "", "b", ""})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice filters empty -- mixed", actual)
}

func Test_Cov2_NonEmptySlice_Empty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- empty input", actual)
}

// ============================================================================
// NonEmptyStrings
// ============================================================================

func Test_Cov2_NonEmptyStrings_Mixed(t *testing.T) {
	result := stringslice.NonEmptyStrings([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings filters empty -- mixed", actual)
}

func Test_Cov2_NonEmptyStrings_Nil(t *testing.T) {
	result := stringslice.NonEmptyStrings(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- nil", actual)
}

func Test_Cov2_NonEmptyStrings_Empty(t *testing.T) {
	result := stringslice.NonEmptyStrings([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptyStrings returns empty -- empty", actual)
}

// ============================================================================
// NonWhitespace
// ============================================================================

func Test_Cov2_NonWhitespace_Mixed(t *testing.T) {
	result := stringslice.NonWhitespace([]string{"a", "  ", "b", ""})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespace filters ws and empty -- mixed", actual)
}

func Test_Cov2_NonWhitespace_Nil(t *testing.T) {
	result := stringslice.NonWhitespace(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- nil", actual)
}

func Test_Cov2_NonWhitespace_Empty(t *testing.T) {
	result := stringslice.NonWhitespace([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonWhitespace returns empty -- empty", actual)
}

// ============================================================================
// Clone
// ============================================================================

func Test_Cov2_Clone_NonEmpty(t *testing.T) {
	result := stringslice.Clone([]string{"a", "b"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "Clone copies slice -- non-empty", actual)
}

func Test_Cov2_Clone_Empty(t *testing.T) {
	result := stringslice.Clone([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- empty", actual)
}

// ============================================================================
// FirstLastDefault
// ============================================================================

func Test_Cov2_FirstLastDefault_Empty(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns empty -- empty", actual)
}

func Test_Cov2_FirstLastDefault_Single(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{"only"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "only", "last": ""}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns first only -- single", actual)
}

func Test_Cov2_FirstLastDefault_Multiple(t *testing.T) {
	first, last := stringslice.FirstLastDefault([]string{"a", "b", "c"})
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefault returns first and last -- multiple", actual)
}

// ============================================================================
// FirstLastDefaultStatus
// ============================================================================

func Test_Cov2_FirstLastDefaultStatus_Empty(t *testing.T) {
	s := stringslice.FirstLastDefaultStatus([]string{})
	actual := args.Map{"valid": s.IsValid, "hasFirst": s.HasFirst, "hasLast": s.HasLast}
	expected := args.Map{"valid": false, "hasFirst": false, "hasLast": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns invalid -- empty", actual)
}

func Test_Cov2_FirstLastDefaultStatus_Single(t *testing.T) {
	s := stringslice.FirstLastDefaultStatus([]string{"only"})
	actual := args.Map{"valid": s.IsValid, "hasFirst": s.HasFirst, "hasLast": s.HasLast, "first": s.First}
	expected := args.Map{"valid": false, "hasFirst": true, "hasLast": false, "first": "only"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns partial -- single", actual)
}

func Test_Cov2_FirstLastDefaultStatus_Multiple(t *testing.T) {
	s := stringslice.FirstLastDefaultStatus([]string{"a", "b"})
	actual := args.Map{"valid": s.IsValid, "hasFirst": s.HasFirst, "hasLast": s.HasLast, "first": s.First, "last": s.Last}
	expected := args.Map{"valid": true, "hasFirst": true, "hasLast": true, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatus returns valid -- multiple", actual)
}

// ============================================================================
// SafeIndexAt
// ============================================================================

func Test_Cov2_SafeIndexAt_Valid(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a", "b"}, 1)}
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns element -- valid index", actual)
}

func Test_Cov2_SafeIndexAt_OutOfBounds(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, 5)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- out of bounds", actual)
}

func Test_Cov2_SafeIndexAt_Negative(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{"a"}, -1)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- negative index", actual)
}

func Test_Cov2_SafeIndexAt_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAt([]string{}, 0)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "SafeIndexAt returns empty -- empty slice", actual)
}

// ============================================================================
// SafeIndexAtWith
// ============================================================================

func Test_Cov2_SafeIndexAtWith_Valid(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")}
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns element -- valid index", actual)
}

func Test_Cov2_SafeIndexAtWith_Default(t *testing.T) {
	actual := args.Map{"result": stringslice.SafeIndexAtWith([]string{"a"}, 5, "def")}
	expected := args.Map{"result": "def"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWith returns default -- out of bounds", actual)
}

// ============================================================================
// NonEmptyJoin
// ============================================================================

func Test_Cov2_NonEmptyJoin_Mixed(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin filters and joins -- mixed", actual)
}

func Test_Cov2_NonEmptyJoin_Nil(t *testing.T) {
	actual := args.Map{"result": stringslice.NonEmptyJoin(nil, ",")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- nil", actual)
}

func Test_Cov2_NonEmptyJoin_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.NonEmptyJoin([]string{}, ",")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- empty", actual)
}

// ============================================================================
// NonWhitespaceJoin
// ============================================================================

func Test_Cov2_NonWhitespaceJoin_Mixed(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin filters ws and joins -- mixed", actual)
}

func Test_Cov2_NonWhitespaceJoin_Nil(t *testing.T) {
	actual := args.Map{"result": stringslice.NonWhitespaceJoin(nil, ",")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- nil", actual)
}

func Test_Cov2_NonWhitespaceJoin_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.NonWhitespaceJoin([]string{}, ",")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns empty -- empty", actual)
}

// ============================================================================
// PrependNew / PrependLineNew
// ============================================================================

func Test_Cov2_PrependNew(t *testing.T) {
	result := stringslice.PrependNew([]string{"b", "c"}, "a")
	actual := args.Map{"len": len(*result), "first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "PrependNew prepends items -- one prepend", actual)
}

func Test_Cov2_PrependNew_Empty(t *testing.T) {
	result := stringslice.PrependNew(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "PrependNew returns empty -- nil both", actual)
}

func Test_Cov2_PrependLineNew(t *testing.T) {
	result := stringslice.PrependLineNew("first", []string{"second"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "first"}
	expected.ShouldBeEqual(t, 0, "PrependLineNew prepends line -- one line", actual)
}

// ============================================================================
// TrimmedEachWords
// ============================================================================

func Test_Cov2_TrimmedEachWords_Mixed(t *testing.T) {
	result := stringslice.TrimmedEachWords([]string{"  a  ", "  ", " b "})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords trims and filters -- mixed", actual)
}

func Test_Cov2_TrimmedEachWords_Nil(t *testing.T) {
	result := stringslice.TrimmedEachWords(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns empty -- nil", actual)
}

func Test_Cov2_TrimmedEachWords_Empty(t *testing.T) {
	result := stringslice.TrimmedEachWords([]string{})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWords returns empty -- empty", actual)
}

// ============================================================================
// ExpandByFunc
// ============================================================================

func Test_Cov2_ExpandByFunc_NonEmpty(t *testing.T) {
	result := stringslice.ExpandByFunc([]string{"a,b", "c"}, func(line string) []string {
		return strings.Split(line, ",")
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc expands -- comma split", actual)
}

func Test_Cov2_ExpandByFunc_Empty(t *testing.T) {
	result := stringslice.ExpandByFunc([]string{}, func(line string) []string { return nil })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc returns empty -- empty input", actual)
}

func Test_Cov2_ExpandByFunc_NilReturn(t *testing.T) {
	result := stringslice.ExpandByFunc([]string{"a"}, func(line string) []string { return nil })
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandByFunc skips nil returns -- nil expand", actual)
}

// ============================================================================
// ExpandBySplits
// ============================================================================

func Test_Cov2_ExpandBySplits_NonEmpty(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{"a,b"}, ",", ";")
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits expands -- comma and semi", actual)
}

func Test_Cov2_ExpandBySplits_Empty(t *testing.T) {
	result := stringslice.ExpandBySplits([]string{}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplits returns empty -- empty input", actual)
}

// ============================================================================
// SplitTrimmedNonEmpty
// ============================================================================

func Test_Cov2_SplitTrimmedNonEmpty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty(" a , b , ", ",", -1)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty trims and filters -- comma", actual)
}

// ============================================================================
// SplitContentsByWhitespace
// ============================================================================

func Test_Cov2_SplitContentsByWhitespace(t *testing.T) {
	result := stringslice.SplitContentsByWhitespace("  hello  world  ")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[1]}
	expected := args.Map{"len": 2, "first": "hello", "last": "world"}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespace splits by ws -- two words", actual)
}

// ============================================================================
// NonNullStrings
// ============================================================================

func Test_Cov2_NonNullStrings_Nil(t *testing.T) {
	result := stringslice.NonNullStrings(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonNullStrings returns empty -- nil", actual)
}

func Test_Cov2_NonNullStrings_NonNil(t *testing.T) {
	result := stringslice.NonNullStrings([]string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NonNullStrings passes through -- non-nil", actual)
}

// ============================================================================
// CloneIf — JoinWith / Joins
// ============================================================================

func Test_Cov2_JoinWith_NonEmpty(t *testing.T) {
	result := stringslice.JoinWith(",", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ",a,b"}
	expected.ShouldBeEqual(t, 0, "JoinWith prepends joiner -- comma", actual)
}

func Test_Cov2_JoinWith_Empty(t *testing.T) {
	result := stringslice.JoinWith(",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "JoinWith returns empty -- no items", actual)
}

func Test_Cov2_Joins_NonEmpty(t *testing.T) {
	result := stringslice.Joins(",", "a", "b")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "Joins joins items -- comma", actual)
}

func Test_Cov2_Joins_Empty(t *testing.T) {
	result := stringslice.Joins(",")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Joins returns empty -- no items", actual)
}

// ============================================================================
// CloneIf — more branches
// ============================================================================

func Test_Cov2_CloneIf_NilNoClone(t *testing.T) {
	result := stringslice.CloneIf(false, 0, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CloneIf nil no-clone returns empty -- nil+false", actual)
}

func Test_Cov2_CloneIf_Clone(t *testing.T) {
	result := stringslice.CloneIf(true, 5, []string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CloneIf clone returns copy -- true", actual)
}

// ============================================================================
// NonEmptySlicePtr
// ============================================================================

func Test_Cov2_NonEmptySlicePtr_Mixed(t *testing.T) {
	result := stringslice.NonEmptySlicePtr([]string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr filters empty -- mixed", actual)
}

// ============================================================================
// NonEmptySlice — extra
// ============================================================================

func Test_Cov2_NonEmptySlice_AllEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"", ""})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- all empty", actual)
}

// ============================================================================
// FirstOrDefault / FirstOrDefaultPtr / LastOrDefault / LastOrDefaultPtr
// ============================================================================

func Test_Cov2_FirstOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringslice.FirstOrDefault([]string{"a", "b"})}
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns first -- non-empty", actual)
}

func Test_Cov2_FirstOrDefault_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.FirstOrDefault(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns empty -- nil", actual)
}

func Test_Cov2_LastOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringslice.LastOrDefault([]string{"a", "b"})}
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns last -- non-empty", actual)
}

func Test_Cov2_LastOrDefault_Empty(t *testing.T) {
	actual := args.Map{"result": stringslice.LastOrDefault(nil)}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault returns empty -- nil", actual)
}
