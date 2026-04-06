package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// First / Last
// ============================================================================

func Test_Cov_First(t *testing.T) {
	actual := args.Map{"result": stringslice.First([]string{"a", "b"})}
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "First returns first element -- non-empty", actual)
}

func Test_Cov_Last(t *testing.T) {
	actual := args.Map{"result": stringslice.Last([]string{"a", "b", "c"})}
	expected := args.Map{"result": "c"}
	expected.ShouldBeEqual(t, 0, "Last returns last element -- non-empty", actual)
}

// ============================================================================
// FirstOrDefaultWith
// ============================================================================

func Test_Cov_FirstOrDefaultWith_NonEmpty(t *testing.T) {
	result, isSuccess := stringslice.FirstOrDefaultWith([]string{"x", "y"}, "def")
	actual := args.Map{"result": result, "isSuccess": isSuccess}
	expected := args.Map{"result": "x", "isSuccess": true}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns first -- non-empty", actual)
}

func Test_Cov_FirstOrDefaultWith_Empty(t *testing.T) {
	result, isSuccess := stringslice.FirstOrDefaultWith([]string{}, "def")
	actual := args.Map{"result": result, "isSuccess": isSuccess}
	expected := args.Map{"result": "def", "isSuccess": false}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns default -- empty", actual)
}

// ============================================================================
// Make / MakeLen / Empty
// ============================================================================

func Test_Cov_Make(t *testing.T) {
	s := stringslice.Make(0, 5)
	actual := args.Map{"len": len(s), "cap": cap(s)}
	expected := args.Map{"len": 0, "cap": 5}
	expected.ShouldBeEqual(t, 0, "Make creates slice with cap -- 0,5", actual)
}

func Test_Cov_MakeLen(t *testing.T) {
	s := stringslice.MakeLen(3)
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLen creates slice with length -- 3", actual)
}

func Test_Cov_Empty(t *testing.T) {
	s := stringslice.Empty()
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty slice -- zero length", actual)
}

// ============================================================================
// HasAnyItem / IsEmpty
// ============================================================================

func Test_Cov_HasAnyItem_True(t *testing.T) {
	actual := args.Map{"result": stringslice.HasAnyItem([]string{"a"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns true -- non-empty", actual)
}

func Test_Cov_HasAnyItem_False(t *testing.T) {
	actual := args.Map{"result": stringslice.HasAnyItem([]string{})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns false -- empty", actual)
}

func Test_Cov_IsEmpty_True(t *testing.T) {
	actual := args.Map{"result": stringslice.IsEmpty([]string{})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- empty", actual)
}

func Test_Cov_IsEmpty_False(t *testing.T) {
	actual := args.Map{"result": stringslice.IsEmpty([]string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns false -- non-empty", actual)
}

// ============================================================================
// IndexAt
// ============================================================================

func Test_Cov_IndexAt(t *testing.T) {
	actual := args.Map{"result": stringslice.IndexAt([]string{"a", "b", "c"}, 1)}
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns element at index -- index 1", actual)
}

// ============================================================================
// AppendLineNew
// ============================================================================

func Test_Cov_AppendLineNew(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	actual := args.Map{"len": len(result), "last": result[len(result)-1]}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew appends item -- one item", actual)
}

func Test_Cov_AppendLineNew_Empty(t *testing.T) {
	result := stringslice.AppendLineNew(nil, "b")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 1, "first": "b"}
	expected.ShouldBeEqual(t, 0, "AppendLineNew appends to nil -- nil slice", actual)
}

// ============================================================================
// MergeNew
// ============================================================================

func Test_Cov_MergeNew_BothEmpty(t *testing.T) {
	result := stringslice.MergeNew(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNew returns empty -- both nil", actual)
}

func Test_Cov_MergeNew_NonEmpty(t *testing.T) {
	result := stringslice.MergeNew([]string{"a"}, "b", "c")
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "MergeNew merges slices -- both non-empty", actual)
}

// ============================================================================
// MergeNewSimple
// ============================================================================

func Test_Cov_MergeNewSimple_Empty(t *testing.T) {
	result := stringslice.MergeNewSimple()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns empty -- no args", actual)
}

func Test_Cov_MergeNewSimple_NonEmpty(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})
	actual := args.Map{"len": len(result), "last": result[2]}
	expected := args.Map{"len": 3, "last": "c"}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple merges slices -- two slices", actual)
}

func Test_Cov_MergeNewSimple_WithEmpty(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{}, []string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple skips empty -- empty first", actual)
}

// ============================================================================
// ClonePtr
// ============================================================================

func Test_Cov_ClonePtr_NonEmpty(t *testing.T) {
	result := stringslice.ClonePtr([]string{"a", "b"})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns copy -- non-empty", actual)
}

func Test_Cov_ClonePtr_Empty(t *testing.T) {
	result := stringslice.ClonePtr(nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns empty -- nil", actual)
}

// ============================================================================
// AppendStringsWithMainSlice
// ============================================================================

func Test_Cov_AppendStringsWithMainSlice_SkipEmpty(t *testing.T) {
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"}, "", "b")
	actual := args.Map{"len": len(result), "last": result[len(result)-1]}
	expected := args.Map{"len": 2, "last": "b"}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice skips empty -- isSkipEmpty", actual)
}

func Test_Cov_AppendStringsWithMainSlice_NoSkip(t *testing.T) {
	result := stringslice.AppendStringsWithMainSlice(false, []string{"a"}, "", "b")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice includes empty -- no skip", actual)
}

func Test_Cov_AppendStringsWithMainSlice_NoItems(t *testing.T) {
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice unchanged -- no items", actual)
}

// ============================================================================
// InPlaceReverse
// ============================================================================

func Test_Cov_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	actual := args.Map{"len": len(*result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns empty -- nil", actual)
}

func Test_Cov_InPlaceReverse_Single(t *testing.T) {
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0]}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse unchanged -- single item", actual)
}

func Test_Cov_InPlaceReverse_Two(t *testing.T) {
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[1]}
	expected := args.Map{"first": "b", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse swaps -- two items", actual)
}

func Test_Cov_InPlaceReverse_Three(t *testing.T) {
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"first": (*result)[0], "last": (*result)[2]}
	expected := args.Map{"first": "c", "last": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- three items", actual)
}

// ============================================================================
// SortIf
// ============================================================================

func Test_Cov_SortIf_True(t *testing.T) {
	result := stringslice.SortIf(true, []string{"c", "a", "b"})
	actual := args.Map{"first": result[0], "last": result[2]}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf sorts -- isSort true", actual)
}

func Test_Cov_SortIf_False(t *testing.T) {
	result := stringslice.SortIf(false, []string{"c", "a", "b"})
	actual := args.Map{"first": result[0]}
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf no-op -- isSort false", actual)
}

// ============================================================================
// ExpandBySplit
// ============================================================================

func Test_Cov_ExpandBySplit_NonEmpty(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{"a,b", "c,d"}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit expands -- comma split", actual)
}

func Test_Cov_ExpandBySplit_Empty(t *testing.T) {
	result := stringslice.ExpandBySplit([]string{}, ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- empty input", actual)
}

// ============================================================================
// NonEmptyIf
// ============================================================================

func Test_Cov_NonEmptyIf_True(t *testing.T) {
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf filters empty -- isNonEmpty true", actual)
}

func Test_Cov_NonEmptyIf_False(t *testing.T) {
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf false calls NonNullStrings -- filters empty", actual)
}

// ============================================================================
// MergeSlicesOfSlices
// ============================================================================

func Test_Cov_MergeSlicesOfSlices_Empty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- no input", actual)
}

func Test_Cov_MergeSlicesOfSlices_NonEmpty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{"b", "c"})
	actual := args.Map{"len": len(result), "first": result[0], "last": result[2]}
	expected := args.Map{"len": 3, "first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices merges -- two slices", actual)
}

func Test_Cov_MergeSlicesOfSlices_WithEmpty(t *testing.T) {
	result := stringslice.MergeSlicesOfSlices([]string{}, []string{"a"})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices skips empty -- one empty", actual)
}
