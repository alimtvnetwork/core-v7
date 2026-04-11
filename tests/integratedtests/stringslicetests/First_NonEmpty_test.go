package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// First / FirstPtr / Last / LastPtr
// ============================================================================

func Test_First_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.First([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "First returns first -- non-empty", actual)
}

func Test_FirstPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.FirstPtr([]string{"x"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "FirstPtr returns first -- non-empty", actual)
}

func Test_Last_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.Last([]string{"a", "b", "c"})}

	// Assert
	expected := args.Map{"val": "c"}
	expected.ShouldBeEqual(t, 0, "Last returns last -- non-empty", actual)
}

func Test_LastPtr_NonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.LastPtr([]string{"x", "y"})

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "y"}
	expected.ShouldBeEqual(t, 0, "LastPtr returns last -- non-empty", actual)
}

// ============================================================================
// FirstOrDefault / LastOrDefault
// ============================================================================

func Test_FirstOrDefault_Empty_FromFirstNonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{})}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- empty", actual)
}

func Test_FirstOrDefault_NonEmpty_FromFirstNonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.FirstOrDefault([]string{"a"})}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault non-empty -- a", actual)
}

func Test_LastOrDefault_Empty_FromFirstNonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefault([]string{})}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- empty", actual)
}

func Test_LastOrDefault_NonEmpty_FromFirstNonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.LastOrDefault([]string{"a", "b"})}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "LastOrDefault non-empty -- b", actual)
}

// ============================================================================
// InPlaceReverse
// ============================================================================

func Test_InPlaceReverse_FromFirstNonEmpty(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "InPlaceReverse reverses -- a,b,c", actual)
}

func Test_InPlaceReverse_Nil_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.InPlaceReverse(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "InPlaceReverse nil -- empty", actual)
}

// ============================================================================
// IndexAt
// ============================================================================

func Test_IndexAt_Valid(t *testing.T) {
	// Act
	actual := args.Map{"val": stringslice.IndexAt([]string{"a", "b"}, 1)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "IndexAt returns element -- index 1", actual)
}

// ============================================================================
// HasAnyItem / IsEmpty
// ============================================================================

func Test_HasAnyItem_FromFirstNonEmpty(t *testing.T) {
	// Act
	actual := args.Map{
		"nonEmpty": stringslice.HasAnyItem([]string{"a"}),
		"empty":    stringslice.HasAnyItem([]string{}),
		"nil":      stringslice.HasAnyItem(nil),
	}

	// Assert
	expected := args.Map{
		"nonEmpty": true,
		"empty": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyItem -- various", actual)
}

func Test_IsEmpty(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    stringslice.IsEmpty([]string{}),
		"nonEmpty": stringslice.IsEmpty([]string{"a"}),
		"nil":      stringslice.IsEmpty(nil),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"nonEmpty": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty -- various", actual)
}

// ============================================================================
// MergeNew / MergeNewSimple
// ============================================================================

func Test_MergeNew_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNew([]string{"a"}, "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MergeNew merges -- slice+item", actual)
}

func Test_MergeNewSimple_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeNewSimple merges -- 2 slices", actual)
}

// ============================================================================
// Make / MakeLen / MakeDefault / Empty
// ============================================================================

func Test_Make_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.Make(0, 5)

	// Act
	actual := args.Map{
		"cap": cap(result) >= 5,
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"cap": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "Make creates slice -- cap 5", actual)
}

func Test_MakeLen_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MakeLen(3)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLen creates slice with length -- len 3", actual)
}

func Test_MakeDefault_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MakeDefault(10)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MakeDefault creates slice -- default", actual)
}

func Test_Empty_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.Empty()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- empty", actual)
}

// ============================================================================
// SortIf
// ============================================================================

func Test_SortIf_True_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(true, s)

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SortIf true sorts -- c,a,b", actual)
}

func Test_SortIf_False_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	s := []string{"c", "a", "b"}
	result := stringslice.SortIf(false, s)

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortIf false no sort -- c,a,b", actual)
}

// ============================================================================
// ClonePtr / CloneUsingCap
// ============================================================================

func Test_ClonePtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ClonePtr clones -- 2 items", actual)
}

func Test_CloneUsingCap_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.CloneUsingCap(10, []string{"a"})

	// Act
	actual := args.Map{
		"len": len(result),
		"capAbove": cap(result) >= 10,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"capAbove": true,
	}
	expected.ShouldBeEqual(t, 0, "CloneUsingCap clones with cap -- cap 10", actual)
}

// ============================================================================
// AppendLineNew
// ============================================================================

func Test_AppendLineNew_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{"a"}, "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"last": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AppendLineNew appends -- b", actual)
}

// ============================================================================
// SlicePtr / EmptyPtr
// ============================================================================

func Test_SlicePtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	s := []string{"a"}
	result := stringslice.SlicePtr(s)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SlicePtr returns slice -- 1 item", actual)
}

func Test_EmptyPtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.EmptyPtr()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "EmptyPtr returns empty -- empty", actual)
}

// ============================================================================
// LengthOfPointer
// ============================================================================

func Test_LengthOfPointer_Valid(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}

	// Act
	actual := args.Map{"len": stringslice.LengthOfPointer(s)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer returns length -- 2", actual)
}

func Test_LengthOfPointer_Nil(t *testing.T) {
	// Act
	actual := args.Map{"len": stringslice.LengthOfPointer(nil)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfPointer nil returns 0 -- nil", actual)
}

// ============================================================================
// MakePtr / MakeLenPtr / MakeDefaultPtr
// ============================================================================

func Test_MakePtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MakePtr(0, 5)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MakePtr returns slice -- cap 5", actual)
}

func Test_MakeLenPtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MakeLenPtr(3)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MakeLenPtr returns slice -- len 3", actual)
}

func Test_MakeDefaultPtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MakeDefaultPtr(10)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MakeDefaultPtr returns slice -- default", actual)
}

// ============================================================================
// FirstLastStatus
// ============================================================================

func Test_FirstLastStatus_Multiple(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}

	// Act
	actual := args.Map{"result": len(s) < 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected at least 2 items", actual)
	actual = args.Map{
		"first": s[0],
		"last": s[len(s)-1],
	}
	expected = args.Map{
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastStatus returns both -- 2 items", actual)
}

func Test_FirstLastStatus_Empty(t *testing.T) {
	// Arrange
	s := []string{}

	// Act
	actual := args.Map{"empty": len(s) == 0}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FirstLastStatus empty -- empty", actual)
}

// ============================================================================
// HasAnyItemPtr / IsEmptyPtr
// ============================================================================

func Test_HasAnyItemPtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	s := []string{"a"}

	// Act
	actual := args.Map{
		"has":    stringslice.HasAnyItemPtr(s),
		"nilPtr": stringslice.HasAnyItemPtr(nil),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"nilPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "HasAnyItemPtr -- valid and nil", actual)
}

func Test_IsEmptyPtr_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	s := []string{"a"}

	// Act
	actual := args.Map{
		"notEmpty": stringslice.IsEmptyPtr(s),
		"nilPtr":   stringslice.IsEmptyPtr(nil),
	}

	// Assert
	expected := args.Map{
		"notEmpty": false,
		"nilPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr -- valid and nil", actual)
}

// ============================================================================
// ExpandBySplit
// ============================================================================

func Test_ExpandBySplit_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.ExpandBySplit([]string{"a,b", "c"}, ",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "ExpandBySplit splits -- comma", actual)
}

// ============================================================================
// NonEmptyIf
// ============================================================================

func Test_NonEmptyIf_True_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptyIf true filters -- mixed", actual)
}

func Test_NonEmptyIf_False_FromFirstNonEmpty(t *testing.T) {
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

func Test_MergeSlicesOfSlices_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeSlicesOfSlices([]string{"a"}, []string{"b", "c"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices merges -- 2 slices", actual)
}

// ============================================================================
// TrimmedEachWordsIf
// ============================================================================

func Test_TrimmedEachWordsIf_True_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(true, []string{"  a  ", "  ", " b "})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf true -- trims", actual)
}

func Test_TrimmedEachWordsIf_False_FromFirstNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWordsIf(false, []string{"  a  ", "  "})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsIf false -- no trim", actual)
}
