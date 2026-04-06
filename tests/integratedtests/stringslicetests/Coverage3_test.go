package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// First / FirstPtr / Last / LastPtr
// ============================================================================

func Test_Cov3_First_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.First([]string{"a", "b"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "First returns first -- non-empty", actual)
}

func Test_Cov3_FirstPtr_NonEmpty(t *testing.T) {
	result := stringslice.FirstPtr([]string{"x"})
	actual := args.Map{"val": result}
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "FirstPtr returns first -- non-empty", actual)
}

func Test_Cov3_Last_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.Last([]string{"a", "b", "c"})}
	expected := args.Map{"val": "c"}
	expected.ShouldBeEqual(t, 0, "Last returns last -- non-empty", actual)
}

func Test_Cov3_LastPtr_NonEmpty(t *testing.T) {
	result := stringslice.LastPtr([]string{"x", "y"})
	actual := args.Map{"val": result}
	expected := args.Map{"val": "y"}
	expected.ShouldBeEqual(t, 0, "LastPtr returns last -- non-empty", actual)
}

// ============================================================================
// FirstOrDefault / LastOrDefault
// ============================================================================

func Test_Cov3_FirstOrDefault_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{})}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- empty", actual)
}

func Test_Cov3_FirstOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{"a"})}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault non-empty -- a", actual)
}

func Test_Cov3_LastOrDefault_Empty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefault([]string{})}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- empty", actual)
}

func Test_Cov3_LastOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"val": stringslice.LastOrDefault([]string{"a", "b"})}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault non-empty -- b", actual)
}

// ============================================================================
// InPlaceReverse
// ============================================================================

func Test_Cov3_InPlaceReverse(t *testing.T) {
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- a,b,c", actual)
}

func Test_Cov3_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse nil -- empty", actual)
}

// ============================================================================
// IndexAt
// ============================================================================

func Test_Cov3_IndexAt_Valid(t *testing.T) {
	actual := args.Map{"val": stringslice.IndexAt([]string{"a", "b"}, 1)}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns element -- index 1", actual)
}

// ============================================================================
// HasAnyItem / IsEmpty
// ============================================================================

func Test_Cov3_HasAnyItem(t *testing.T) {
	actual := args.Map{
		"nonEmpty": stringslice.HasAnyItem([]string{"a"}),
		"empty":    stringslice.HasAnyItem([]string{}),
		"nil":      stringslice.HasAnyItem(nil),
	}
	expected := args.Map{"nonEmpty": true, "empty": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem -- various", actual)
}

func Test_Cov3_IsEmpty(t *testing.T) {
	actual := args.Map{
		"empty":    stringslice.IsEmpty([]string{}),
		"nonEmpty": stringslice.IsEmpty([]string{"a"}),
		"nil":      stringslice.IsEmpty(nil),
	}
	expected := args.Map{"empty": true, "nonEmpty": false, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty -- various", actual)
}

// ============================================================================
// MergeNew / MergeNewSimple
// ============================================================================

func Test_Cov3_MergeNew(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MergeNew merges -- slice+item", actual)
}

func Test_Cov3_MergeNewSimple(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple merges -- 2 slices", actual)
}

// ============================================================================
// Make / MakeLen / MakeDefault / Empty
// ============================================================================

func Test_Cov3_Make(t *testing.T) {
	result := stringslice.Make(0, 5)
	actual := args.Map{"cap": cap(result) >= 5, "len": len(result)}
	expected := args.Map{"cap": true, "len": 0}
	expected.ShouldBeEqual(t, 0, "Make creates slice -- cap 5", actual)
}

func Test_Cov3_MakeLen(t *testing.T) {
	result := stringslice.MakeLen(3)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLen creates slice with length -- len 3", actual)
}

func Test_Cov3_MakeDefault(t *testing.T) {
	result := stringslice.MakeDefault(10)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MakeDefault creates slice -- default", actual)
}

func Test_Cov3_Empty(t *testing.T) {
	result := stringslice.Empty()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- empty", actual)
}

// ============================================================================
// SortIf
// ============================================================================

func Test_Cov3_SortIf_True(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SortIf true sorts -- c,a,b", actual)
}

func Test_Cov3_SortIf_False(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf false no sort -- c,a,b", actual)
}

// ============================================================================
// ClonePtr / CloneUsingCap
// ============================================================================

func Test_Cov3_ClonePtr(t *testing.T) {
	result := stringslice.ClonePtr([]string{"a", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ClonePtr clones -- 2 items", actual)
}

func Test_Cov3_CloneUsingCap(t *testing.T) {
	result := stringslice.CloneUsingCap(10, []string{"a"})
	actual := args.Map{"len": len(result), "capAbove": cap(result) >= 10}
	expected := args.Map{"len": 1, "capAbove": true}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap clones with cap -- cap 10", actual)
}

// ============================================================================
// AppendLineNew
// ============================================================================

func Test_Cov3_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result), "last": result[1]}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew appends -- b", actual)
}

// ============================================================================
// SlicePtr / EmptyPtr
// ============================================================================

func Test_Cov3_SlicePtr(t *testing.T) {
	s := []string{"a"}
	result := stringslice.SlicePtr(s)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns slice -- 1 item", actual)
}

func Test_Cov3_EmptyPtr(t *testing.T) {
	result := stringslice.EmptyPtr()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyPtr returns empty -- empty", actual)
}

// ============================================================================
// LengthOfPointer
// ============================================================================

func Test_Cov3_LengthOfPointer_Valid(t *testing.T) {
	s := []string{"a", "b"}
	actual := args.Map{"len": stringslice.LengthOfPointer(s)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer returns length -- 2", actual)
}

func Test_Cov3_LengthOfPointer_Nil(t *testing.T) {
	actual := args.Map{"len": stringslice.LengthOfPointer(nil)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer nil returns 0 -- nil", actual)
}

// ============================================================================
// MakePtr / MakeLenPtr / MakeDefaultPtr
// ============================================================================

func Test_Cov3_MakePtr(t *testing.T) {
	result := stringslice.MakePtr(0, 5)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MakePtr returns slice -- cap 5", actual)
}

func Test_Cov3_MakeLenPtr(t *testing.T) {
	result := stringslice.MakeLenPtr(3)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLenPtr returns slice -- len 3", actual)
}

func Test_Cov3_MakeDefaultPtr(t *testing.T) {
	result := stringslice.MakeDefaultPtr(10)
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MakeDefaultPtr returns slice -- default", actual)
}

// ============================================================================
// FirstLastStatus
// ============================================================================

func Test_Cov3_FirstLastStatus_Multiple(t *testing.T) {
	s := []string{"a", "b"}
	if len(s) < 2 {
		t.Error("expected at least 2 items")
	}
	actual := args.Map{"first": s[0], "last": s[len(s)-1]}
	expected := args.Map{"first": "a", "last": "b"}
	expected.ShouldBeEqual(t, 0, "FirstLastStatus returns both -- 2 items", actual)
}

func Test_Cov3_FirstLastStatus_Empty(t *testing.T) {
	s := []string{}
	actual := args.Map{"empty": len(s) == 0}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FirstLastStatus empty -- empty", actual)
}

// ============================================================================
// HasAnyItemPtr / IsEmptyPtr
// ============================================================================

func Test_Cov3_HasAnyItemPtr(t *testing.T) {
	s := []string{"a"}
	actual := args.Map{
		"has":    stringslice.HasAnyItemPtr(s),
		"nilPtr": stringslice.HasAnyItemPtr(nil),
	}
	expected := args.Map{"has": true, "nilPtr": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr -- valid and nil", actual)
}

func Test_Cov3_IsEmptyPtr(t *testing.T) {
	s := []string{"a"}
	actual := args.Map{
		"notEmpty": stringslice.IsEmptyPtr(s),
		"nilPtr":   stringslice.IsEmptyPtr(nil),
	}
	expected := args.Map{"notEmpty": false, "nilPtr": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr -- valid and nil", actual)
}

// ============================================================================
// ExpandBySplit
// ============================================================================

func Test_Cov3_ExpandBySplit(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{"a,b", "c"}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit splits -- comma", actual)
}

// ============================================================================
// NonEmptyIf
// ============================================================================

func Test_Cov3_NonEmptyIf_True(t *testing.T) {
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf true filters -- mixed", actual)
}

func Test_Cov3_NonEmptyIf_False(t *testing.T) {
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf false calls NonNullStrings -- filters empty", actual)
}

// ============================================================================
// MergeSlicesOfSlices
// ============================================================================

func Test_Cov3_MergeSlicesOfSlices(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{"b", "c"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices merges -- 2 slices", actual)
}

// ============================================================================
// TrimmedEachWordsIf
// ============================================================================

func Test_Cov3_TrimmedEachWordsIf_True(t *testing.T) {
	result := stringslice.TrimmedEachWordsIf(true, []string{"  a  ", "  ", " b "})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf true -- trims", actual)
}

func Test_Cov3_TrimmedEachWordsIf_False(t *testing.T) {
	result := stringslice.TrimmedEachWordsIf(false, []string{"  a  ", "  "})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf false -- no trim", actual)
}
