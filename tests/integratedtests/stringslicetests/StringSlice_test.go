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
	// Arrange
	src := []string{"a", "b", "c"}
	result := stringslice.Clone(src)

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone: expected 3", actual)
	// independence
	result[0] = "z"
	actual = args.Map{"result": src[0] == "z"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone should produce independent copy", actual)
}

func Test_StringSlice_Clone_Empty(t *testing.T) {
	// Arrange
	result := stringslice.Clone([]string{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone empty: expected 0", actual)
}

func Test_StringSlice_Clone_Nil(t *testing.T) {
	// Arrange
	result := stringslice.Clone(nil)

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone nil should return non-nil empty slice", actual)
	actual = args.Map{"result": len(result) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Clone nil: expected 0", actual)
}

// ==========================================
// CloneUsingCap
// ==========================================

func Test_StringSlice_CloneUsingCap_AddsCapacity(t *testing.T) {
	// Arrange
	src := []string{"a"}
	result := stringslice.CloneUsingCap(10, src)

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap len: expected 1", actual)
	actual = args.Map{"result": cap(result) < 11}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap cap: expected >= 11", actual)
}

func Test_StringSlice_CloneUsingCap_Empty(t *testing.T) {
	// Arrange
	result := stringslice.CloneUsingCap(5, []string{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// FirstOrDefault / LastOrDefault
// ==========================================

func Test_StringSlice_FirstOrDefault_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstOrDefault([]string{"x", "y"})

	// Act
	actual := args.Map{"result": result != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x', got ''", actual)
}

func Test_StringSlice_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	result := stringslice.FirstOrDefault([]string{})

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_LastOrDefault_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LastOrDefault([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"result": result != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'c', got ''", actual)
}

func Test_StringSlice_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	result := stringslice.LastOrDefault([]string{})

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_LastOrDefault_Single(t *testing.T) {
	// Arrange
	result := stringslice.LastOrDefault([]string{"only"})

	// Act
	actual := args.Map{"result": result != "only"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'only', got ''", actual)
}

// ==========================================
// SafeIndexAt
// ==========================================

func Test_StringSlice_SafeIndexAt_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{"a", "b", "c"}, 1)

	// Act
	actual := args.Map{"result": result != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'b', got ''", actual)
}

func Test_StringSlice_SafeIndexAt_OutOfBounds(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{"a"}, 5)

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_SafeIndexAt_NegativeIndex(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{"a"}, -1)

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

func Test_StringSlice_SafeIndexAt_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAt([]string{}, 0)

	// Act
	actual := args.Map{"result": result != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string, got ''", actual)
}

// ==========================================
// InPlaceReverse
// ==========================================

func Test_StringSlice_InPlaceReverse_Multiple(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c", "d"}
	result := stringslice.InPlaceReverse(&s)
	r := *result

	// Act
	actual := args.Map{"result": r[0] != "d" || r[1] != "c" || r[2] != "b" || r[3] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [d c b a]", actual)
}

func Test_StringSlice_InPlaceReverse_Two(t *testing.T) {
	// Arrange
	s := []string{"x", "y"}
	result := stringslice.InPlaceReverse(&s)
	r := *result

	// Act
	actual := args.Map{"result": r[0] != "y" || r[1] != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [y x]", actual)
}

func Test_StringSlice_InPlaceReverse_Single(t *testing.T) {
	// Arrange
	s := []string{"only"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{"result": (*result)[0] != "only"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single element should remain unchanged", actual)
}

func Test_StringSlice_InPlaceReverse_Empty(t *testing.T) {
	// Arrange
	s := []string{}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{"result": len(*result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should remain empty", actual)
}

func Test_StringSlice_InPlaceReverse_Nil(t *testing.T) {
	// Arrange
	result := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{"result": result == nil || len(*result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice ptr", actual)
}

// ==========================================
// MergeNew
// ==========================================

func Test_StringSlice_MergeNew_BothNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a", "b"}, "c", "d")

	// Act
	actual := args.Map{"result": len(result) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual = args.Map{"result": result[0] != "a" || result[3] != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected order:", actual)
}

func Test_StringSlice_MergeNew_EmptyFirst(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{}, "x")

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_StringSlice_MergeNew_NoAdditional(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a", "b"})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringSlice_MergeNew_BothEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// NonEmptySlice
// ==========================================

func Test_StringSlice_NonEmpty_FiltersEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "", "b", "", "c"})

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_StringSlice_NonEmpty_AllEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"", "", ""})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_StringSlice_NonEmpty_NoneEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{"a", "b"})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringSlice_NonEmpty_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptySlice([]string{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// NonWhitespace
// ==========================================

func Test_StringSlice_NonWhitespace_FiltersWhitespace(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{"a", "  ", "b", "\t", "c"})

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_StringSlice_NonWhitespace_Nil(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace(nil)

	// Act
	actual := args.Map{"result": result == nil || len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
}

func Test_StringSlice_NonWhitespace_Empty(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespace([]string{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// IsEmpty / HasAnyItem
// ==========================================

func Test_StringSlice_IsEmpty_True(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmpty([]string{})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty slice should be empty", actual)
}

func Test_StringSlice_IsEmpty_False(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmpty([]string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty slice should not be empty", actual)
}

func Test_StringSlice_HasAnyItem_True(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItem([]string{"a"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_StringSlice_HasAnyItem_False(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItem([]string{})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should not have items", actual)
}

// ==========================================
// SortIf
// ==========================================

func Test_StringSlice_SortIf_True(t *testing.T) {
	// Arrange
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)

	// Act
	actual := args.Map{"result": result[0] != "a" || result[1] != "b" || result[2] != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_StringSlice_SortIf_False(t *testing.T) {
	// Arrange
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)

	// Act
	actual := args.Map{"result": result[0] != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected unsorted", actual)
}

// ==========================================
// SafeRangeItems
// ==========================================

func Test_StringSlice_SafeRangeItems_ValidRange(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c", "d", "e"}
	result := stringslice.SafeRangeItems(s, 1, 3)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": result[0] != "b" || result[1] != "c"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [b c]", actual)
}

func Test_StringSlice_SafeRangeItems_Nil(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems(nil, 0, 1)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil: expected 0", actual)
}

func Test_StringSlice_SafeRangeItems_Empty(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{}, 0, 1)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty: expected 0", actual)
}

func Test_StringSlice_SafeRangeItems_StartBeyondLength(t *testing.T) {
	// Arrange
	result := stringslice.SafeRangeItems([]string{"a"}, 5, 10)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "start beyond: expected 0", actual)
}

// ==========================================
// ExpandByFunc
// ==========================================

func Test_StringSlice_ExpandByFunc_Basic(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(
		[]string{"a,b", "c,d"},
		func(line string) []string {
			return []string{line + "-1", line + "-2"}
		},
	)

	// Act
	actual := args.Map{"result": len(result) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_StringSlice_ExpandByFunc_Empty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(
		[]string{},
		func(line string) []string { return []string{line} },
	)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_StringSlice_ExpandByFunc_SkipsNilReturn(t *testing.T) {
	// Arrange
	result := stringslice.ExpandByFunc(
		[]string{"a", "skip", "b"},
		func(line string) []string {
			if line == "skip" {
				return nil
			}
			return []string{line}
		},
	)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (skip nil return)", actual)
}
