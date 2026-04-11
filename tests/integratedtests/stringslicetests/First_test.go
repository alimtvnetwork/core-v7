package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// First / Last
// ============================================================================

func Test_First_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.First([]string{"a", "b"})}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "First returns first element -- non-empty", actual)
}

func Test_Last_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.Last([]string{"a", "b", "c"})}

	// Assert
	expected := args.Map{"result": "c"}
	expected.ShouldBeEqual(t, 0, "Last returns last element -- non-empty", actual)
}

// ============================================================================
// FirstOrDefaultWith
// ============================================================================

func Test_FirstOrDefaultWith_NonEmpty_FromFirst(t *testing.T) {
	// Arrange
	result, isSuccess := stringslice.FirstOrDefaultWith([]string{"x", "y"}, "def")

	// Act
	actual := args.Map{
		"result": result,
		"isSuccess": isSuccess,
	}

	// Assert
	expected := args.Map{
		"result": "x",
		"isSuccess": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns first -- non-empty", actual)
}

func Test_FirstOrDefaultWith_Empty_FromFirst(t *testing.T) {
	// Arrange
	result, isSuccess := stringslice.FirstOrDefaultWith([]string{}, "def")

	// Act
	actual := args.Map{
		"result": result,
		"isSuccess": isSuccess,
	}

	// Assert
	expected := args.Map{
		"result": "def",
		"isSuccess": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultWith returns default -- empty", actual)
}

// ============================================================================
// Make / MakeLen / Empty
// ============================================================================

func Test_Make_FromFirst(t *testing.T) {
	// Arrange
	s := stringslice.Make(0, 5)

	// Act
	actual := args.Map{
		"len": len(s),
		"cap": cap(s),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"cap": 5,
	}
	expected.ShouldBeEqual(t, 0, "Make creates slice with cap -- 0,5", actual)
}

func Test_MakeLen_FromFirst(t *testing.T) {
	// Arrange
	s := stringslice.MakeLen(3)

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLen creates slice with length -- 3", actual)
}

func Test_Empty_FromFirst(t *testing.T) {
	// Arrange
	s := stringslice.Empty()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty slice -- zero length", actual)
}

// ============================================================================
// HasAnyItem / IsEmpty
// ============================================================================

func Test_HasAnyItem_True_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItem([]string{"a"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns true -- non-empty", actual)
}

func Test_HasAnyItem_False_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItem([]string{})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasAnyItem returns false -- empty", actual)
}

func Test_IsEmpty_True_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmpty([]string{})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- empty", actual)
}

func Test_IsEmpty_False_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmpty([]string{"a"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns false -- non-empty", actual)
}

// ============================================================================
// IndexAt
// ============================================================================

func Test_IndexAt_FromFirst(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IndexAt([]string{"a", "b", "c"}, 1)}

	// Assert
	expected := args.Map{"result": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns element at index -- index 1", actual)
}

// ============================================================================
// AppendLineNew
// ============================================================================

func Test_AppendLineNew_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{"a"}, "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[len(result)-1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendLineNew appends item -- one item", actual)
}

func Test_AppendLineNew_Empty(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew(nil, "b")

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
	expected.ShouldBeEqual(t, 0, "AppendLineNew appends to nil -- nil slice", actual)
}

// ============================================================================
// MergeNew
// ============================================================================

func Test_MergeNew_BothEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNew returns empty -- both nil", actual)
}

func Test_MergeNew_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a"}, "b", "c")

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
	expected.ShouldBeEqual(t, 0, "MergeNew merges slices -- both non-empty", actual)
}

// ============================================================================
// MergeNewSimple
// ============================================================================

func Test_MergeNewSimple_Empty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple returns empty -- no args", actual)
}

func Test_MergeNewSimple_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple merges slices -- two slices", actual)
}

func Test_MergeNewSimple_WithEmpty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{}, []string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple skips empty -- empty first", actual)
}

// ============================================================================
// ClonePtr
// ============================================================================

func Test_ClonePtr_NonEmpty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr([]string{"a", "b"})

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
	expected.ShouldBeEqual(t, 0, "ClonePtr returns copy -- non-empty", actual)
}

func Test_ClonePtr_Empty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns empty -- nil", actual)
}

// ============================================================================
// AppendStringsWithMainSlice
// ============================================================================

func Test_AppendStringsWithMainSlice_SkipEmpty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"}, "", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[len(result)-1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice skips empty -- isSkipEmpty", actual)
}

func Test_AppendStringsWithMainSlice_NoSkip(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithMainSlice(false, []string{"a"}, "", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice includes empty -- no skip", actual)
}

func Test_AppendStringsWithMainSlice_NoItems_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.AppendStringsWithMainSlice(true, []string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithMainSlice unchanged -- no items", actual)
}

// ============================================================================
// InPlaceReverse
// ============================================================================

func Test_InPlaceReverse_Nil_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse returns empty -- nil", actual)
}

func Test_InPlaceReverse_Single_FromFirst(t *testing.T) {
	// Arrange
	s := []string{"a"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{"first": (*result)[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse unchanged -- single item", actual)
}

func Test_InPlaceReverse_Two_FromFirst(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[1],
	}

	// Assert
	expected := args.Map{
		"first": "b",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse swaps -- two items", actual)
}

func Test_InPlaceReverse_Three_FromFirst(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c"}
	result := stringslice.InPlaceReverse(&s)

	// Act
	actual := args.Map{
		"first": (*result)[0],
		"last": (*result)[2],
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- three items", actual)
}

// ============================================================================
// SortIf
// ============================================================================

func Test_SortIf_True_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.SortIf(true, []string{"c", "a", "b"})

	// Act
	actual := args.Map{
		"first": result[0],
		"last": result[2],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SortIf sorts -- isSort true", actual)
}

func Test_SortIf_False_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.SortIf(false, []string{"c", "a", "b"})

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf no-op -- isSort false", actual)
}

// ============================================================================
// ExpandBySplit
// ============================================================================

func Test_ExpandBySplit_NonEmpty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{"a,b", "c,d"}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit expands -- comma split", actual)
}

func Test_ExpandBySplit_Empty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit returns empty -- empty input", actual)
}

// ============================================================================
// NonEmptyIf
// ============================================================================

func Test_NonEmptyIf_True_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf filters empty -- isNonEmpty true", actual)
}

func Test_NonEmptyIf_False_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf false calls NonNullStrings -- filters empty", actual)
}

// ============================================================================
// MergeSlicesOfSlices
// ============================================================================

func Test_MergeSlicesOfSlices_Empty_FromFirst(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices returns empty -- no input", actual)
}

func Test_MergeSlicesOfSlices_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{"b", "c"})

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
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices merges -- two slices", actual)
}

func Test_MergeSlicesOfSlices_WithEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{}, []string{"a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices skips empty -- one empty", actual)
}
