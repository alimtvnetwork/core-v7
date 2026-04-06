package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Clone
// ==========================================

func Test_StringSlice_Clone_NonEmpty(t *testing.T) {
	src := []string{"a", "b", "c"}
	result := stringslice.Clone(src)
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone: expected 3", actual)
	// independence
	result[0] = "z"
	actual := args.Map{"result": src[0] == "z"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should produce independent copy", actual)
}

func Test_StringSlice_Clone_Empty(t *testing.T) {
	result := stringslice.Clone([]string{})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone empty: expected 0", actual)
}

func Test_StringSlice_Clone_Nil(t *testing.T) {
	result := stringslice.Clone(nil)
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone nil should return non-nil empty slice", actual)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone nil: expected 0", actual)
}

// ==========================================
// CloneUsingCap
// ==========================================

func Test_StringSlice_CloneUsingCap_AddsCapacity(t *testing.T) {
	src := []string{"a"}
	result := stringslice.CloneUsingCap(10, src)
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap len: expected 1", actual)
	actual := args.Map{"result": cap(result) < 11}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap cap: expected >= 11", actual)
}

func Test_StringSlice_CloneUsingCap_Empty(t *testing.T) {
	result := stringslice.CloneUsingCap(5, []string{})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// FirstOrDefault / LastOrDefault
// ==========================================

func Test_StringSlice_FirstOrDefault_NonEmpty(t *testing.T) {
	result := stringslice.FirstOrDefault([]string{"x", "y"})
	actual := args.Map{"result": result != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x', got ''", actual)
}

func Test_StringSlice_FirstOrDefault_Empty(t *testing.T) {
	result := stringslice.FirstOrDefault([]string{})
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_LastOrDefault_NonEmpty(t *testing.T) {
	result := stringslice.LastOrDefault([]string{"a", "b", "c"})
	actual := args.Map{"result": result != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'c', got ''", actual)
}

func Test_StringSlice_LastOrDefault_Empty(t *testing.T) {
	result := stringslice.LastOrDefault([]string{})
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_LastOrDefault_Single(t *testing.T) {
	result := stringslice.LastOrDefault([]string{"only"})
	actual := args.Map{"result": result != "only"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'only', got ''", actual)
}

// ==========================================
// SafeIndexAt
// ==========================================

func Test_StringSlice_SafeIndexAt_Valid(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a", "b", "c"}, 1)
	actual := args.Map{"result": result != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'b', got ''", actual)
}

func Test_StringSlice_SafeIndexAt_OutOfBounds(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a"}, 5)
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_SafeIndexAt_NegativeIndex(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{"a"}, -1)
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_SafeIndexAt_EmptySlice(t *testing.T) {
	result := stringslice.SafeIndexAt([]string{}, 0)
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

// ==========================================
// InPlaceReverse
// ==========================================

func Test_StringSlice_InPlaceReverse_Multiple(t *testing.T) {
	s := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&s)
	r := *result
	actual := args.Map{"result": r[0] != "d" || r[1] != "c" || r[2] != "b" || r[3] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [d c b a]", actual)
}

func Test_StringSlice_InPlaceReverse_Two(t *testing.T) {
	s := []string{"x", "y"}
	result := stringslice.InPlaceReverse(&s)
	r := *result
	actual := args.Map{"result": r[0] != "y" || r[1] != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [y x]", actual)
}

func Test_StringSlice_InPlaceReverse_Single(t *testing.T) {
	s := []string{"only"}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"result": (*result)[0] != "only"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single element should remain unchanged", actual)
}

func Test_StringSlice_InPlaceReverse_Empty(t *testing.T) {
	s := []string{}
	result := stringslice.InPlaceReverse(&s)
	actual := args.Map{"result": len(*result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should remain empty", actual)
}

func Test_StringSlice_InPlaceReverse_Nil(t *testing.T) {
	result := stringslice.InPlaceReverse(nil)
	actual := args.Map{"result": result == nil || len(*result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice ptr", actual)
}

// ==========================================
// MergeNew
// ==========================================

func Test_StringSlice_MergeNew_BothNonEmpty(t *testing.T) {
	result := stringslice.MergeNew([]string{"a", "b"}, "c", "d")
	actual := args.Map{"result": len(result) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": result[0] != "a" || result[3] != "d"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected order:", actual)
}

func Test_StringSlice_MergeNew_EmptyFirst(t *testing.T) {
	result := stringslice.MergeNew([]string{}, "x")
	actual := args.Map{"result": len(result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_StringSlice_MergeNew_NoAdditional(t *testing.T) {
	result := stringslice.MergeNew([]string{"a", "b"})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringSlice_MergeNew_BothEmpty(t *testing.T) {
	result := stringslice.MergeNew([]string{})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// NonEmptySlice
// ==========================================

func Test_StringSlice_NonEmpty_FiltersEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "", "b", "", "c"})
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_StringSlice_NonEmpty_AllEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"", "", ""})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_StringSlice_NonEmpty_NoneEmpty(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{"a", "b"})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringSlice_NonEmpty_EmptySlice(t *testing.T) {
	result := stringslice.NonEmptySlice([]string{})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// NonWhitespace
// ==========================================

func Test_StringSlice_NonWhitespace_FiltersWhitespace(t *testing.T) {
	result := stringslice.NonWhitespace([]string{"a", "  ", "b", "\t", "c"})
	actual := args.Map{"result": len(result) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_StringSlice_NonWhitespace_Nil(t *testing.T) {
	result := stringslice.NonWhitespace(nil)
	actual := args.Map{"result": result == nil || len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
}

func Test_StringSlice_NonWhitespace_Empty(t *testing.T) {
	result := stringslice.NonWhitespace([]string{})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// IsEmpty / HasAnyItem
// ==========================================

func Test_StringSlice_IsEmpty_True(t *testing.T) {
	actual := args.Map{"result": stringslice.IsEmpty([]string{})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty slice should be empty", actual)
}

func Test_StringSlice_IsEmpty_False(t *testing.T) {
	actual := args.Map{"result": stringslice.IsEmpty([]string{"a"})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty slice should not be empty", actual)
}

func Test_StringSlice_HasAnyItem_True(t *testing.T) {
	actual := args.Map{"result": stringslice.HasAnyItem([]string{"a"})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_StringSlice_HasAnyItem_False(t *testing.T) {
	actual := args.Map{"result": stringslice.HasAnyItem([]string{})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have items", actual)
}

// ==========================================
// SortIf
// ==========================================

func Test_StringSlice_SortIf_True(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)
	actual := args.Map{"result": result[0] != "a" || result[1] != "b" || result[2] != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_StringSlice_SortIf_False(t *testing.T) {
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)
	actual := args.Map{"result": result[0] != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unsorted", actual)
}

// ==========================================
// SafeRangeItems
// ==========================================

func Test_StringSlice_SafeRangeItems_ValidRange(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	result := stringslice.SafeRangeItems(s, 1, 3)
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": result[0] != "b" || result[1] != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [b c]", actual)
}

func Test_StringSlice_SafeRangeItems_Nil(t *testing.T) {
	result := stringslice.SafeRangeItems(nil, 0, 1)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil: expected 0", actual)
}

func Test_StringSlice_SafeRangeItems_Empty(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{}, 0, 1)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty: expected 0", actual)
}

func Test_StringSlice_SafeRangeItems_StartBeyondLength(t *testing.T) {
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "start beyond: expected 0", actual)
}

// ==========================================
// ExpandByFunc
// ==========================================

func Test_StringSlice_ExpandByFunc_Basic(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{"a,b", "c,d"},
		func(line string) []string {
			return []string{line + "-1", line + "-2"}
		},
	)
	actual := args.Map{"result": len(result) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_StringSlice_ExpandByFunc_Empty(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{},
		func(line string) []string { return []string{line} },
	)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_StringSlice_ExpandByFunc_SkipsNilReturn(t *testing.T) {
	result := stringslice.ExpandByFunc(
		[]string{"a", "skip", "b"},
		func(line string) []string {
			if line == "skip" {
				return nil
			}
			return []string{line}
		},
	)
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (skip nil return)", actual)
}
