package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// PrependNew
// ==========================================

func Test_StringSlice_PrependNew_BothNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"c", "d"}, "a", "b")
	r := *result

	// Act
	actual := args.Map{"result": len(r) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual = args.Map{"result": r[0] != "a" || r[1] != "b" || r[2] != "c" || r[3] != "d"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected order:", actual)
}

func Test_StringSlice_PrependNew_EmptyPrepend(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{"a", "b"})
	r := *result

	// Act
	actual := args.Map{"result": len(r) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringSlice_PrependNew_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{}, "x")
	r := *result

	// Act
	actual := args.Map{"result": len(r) != 1 || r[0] != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [x]", actual)
}

func Test_StringSlice_PrependNew_BothEmpty(t *testing.T) {
	// Arrange
	result := stringslice.PrependNew([]string{})
	r := *result

	// Act
	actual := args.Map{"result": len(r) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// AppendLineNew
// ==========================================

func Test_StringSlice_AppendLineNew_Basic(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{"a"}, "b")

	// Act
	actual := args.Map{"result": len(result) != 2 || result[1] != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [a b]", actual)
}

func Test_StringSlice_AppendLineNew_EmptySlice(t *testing.T) {
	// Arrange
	result := stringslice.AppendLineNew([]string{}, "x")

	// Act
	actual := args.Map{"result": len(result) != 1 || result[0] != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [x]", actual)
}

// ==========================================
// PrependLineNew
// ==========================================

func Test_StringSlice_PrependLineNew_Basic(t *testing.T) {
	// Arrange
	result := stringslice.PrependLineNew("a", []string{"b", "c"})

	// Act
	actual := args.Map{"result": len(result) != 3 || result[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [a b c]", actual)
}

// ==========================================
// MakeDefault / Make / MakeLen / MakePtr / MakeLenPtr / MakeDefaultPtr
// ==========================================

func Test_StringSlice_MakeDefault(t *testing.T) {
	// Arrange
	result := stringslice.MakeDefault(5)

	// Act
	actual := args.Map{"result": len(result) != 0 || cap(result) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len=0 cap=5, got len= cap=", actual)
}

func Test_StringSlice_Make(t *testing.T) {
	// Arrange
	result := stringslice.Make(3, 5)

	// Act
	actual := args.Map{"result": len(result) != 3 || cap(result) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len=3 cap=5, got len= cap=", actual)
}

func Test_StringSlice_MakeLen(t *testing.T) {
	// Arrange
	result := stringslice.MakeLen(3)

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len=3", actual)
}

// ==========================================
// Empty / EmptyPtr
// ==========================================

func Test_StringSlice_Empty(t *testing.T) {
	// Arrange
	result := stringslice.Empty()

	// Act
	actual := args.Map{"result": result == nil || len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty should return non-nil empty slice", actual)
}

func Test_StringSlice_EmptyPtr(t *testing.T) {
	// Arrange
	result := stringslice.EmptyPtr()

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EmptyPtr should return empty slice", actual)
}

// ==========================================
// IsEmpty / IsEmptyPtr / HasAnyItem / HasAnyItemPtr
// ==========================================

func Test_StringSlice_IsEmptyPtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.IsEmptyPtr(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil ptr should be empty", actual)
}

func Test_StringSlice_IsEmptyPtr_NonNil(t *testing.T) {
	// Arrange
	s := []string{"a"}

	// Act
	actual := args.Map{"result": stringslice.IsEmptyPtr(s)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-empty should not be empty", actual)
}

func Test_StringSlice_HasAnyItemPtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.HasAnyItemPtr(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should not have items", actual)
}

func Test_StringSlice_HasAnyItemPtr_NonNil(t *testing.T) {
	// Arrange
	s := []string{"a"}

	// Act
	actual := args.Map{"result": stringslice.HasAnyItemPtr(s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "non-empty should have items", actual)
}

// ==========================================
// CloneIf
// ==========================================

func Test_StringSlice_CloneIf_True(t *testing.T) {
	// Arrange
	src := []string{"a", "b"}
	result := stringslice.CloneIf(true, 0, src)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	result[0] = "z"
	actual = args.Map{"result": src[0] == "z"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CloneIf true should produce independent copy", actual)
}

func Test_StringSlice_CloneIf_False(t *testing.T) {
	// Arrange
	src := []string{"a", "b"}
	result := stringslice.CloneIf(false, 0, src)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ==========================================
// ClonePtr
// ==========================================

func Test_StringSlice_ClonePtr_Nil(t *testing.T) {
	// Arrange
	result := stringslice.ClonePtr(nil)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
}

func Test_StringSlice_ClonePtr_NonEmpty(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}
	result := stringslice.ClonePtr(s)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ==========================================
// MergeNewSimple
// ==========================================

func Test_StringSlice_MergeNewSimple_Basic(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_StringSlice_MergeNewSimple_BothEmpty(t *testing.T) {
	// Arrange
	result := stringslice.MergeNewSimple([]string{}, []string{})

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// NonEmptyIf
// ==========================================

func Test_StringSlice_NonEmptyIf_True(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_StringSlice_NonEmptyIf_False(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (NonNullStrings filters empty)", actual)
}

// ==========================================
// NonEmptyJoin / NonEmptyJoinPtr
// ==========================================

func Test_StringSlice_NonEmptyJoin(t *testing.T) {
	// Arrange
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"result": result != "a,b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
}

// ==========================================
// NonWhitespaceJoin
// ==========================================

func Test_StringSlice_NonWhitespaceJoin(t *testing.T) {
	// Arrange
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ",")

	// Act
	actual := args.Map{"result": result != "a,b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
}

// ==========================================
// SlicePtr / LengthOfPointer
// ==========================================

func Test_StringSlice_SlicePtr(t *testing.T) {
	// Arrange
	s := []string{"a", "b"}
	result := stringslice.SlicePtr(s)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SlicePtr should return slice", actual)
}

func Test_StringSlice_LengthOfPointer_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": stringslice.LengthOfPointer(nil) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_StringSlice_LengthOfPointer_NonNil(t *testing.T) {
	// Arrange
	s := []string{"a", "b", "c"}

	// Act
	actual := args.Map{"result": stringslice.LengthOfPointer(s) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// ==========================================
// IndexesDefault / SafeIndexes
// ==========================================

func Test_StringSlice_IndexesDefault(t *testing.T) {
	// Arrange
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 0, 2)

	// Act
	actual := args.Map{"result": len(result) != 2 || result[0] != "a" || result[1] != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected result:", actual)
}

// ==========================================
// SplitTrimmedNonEmpty
// ==========================================

func Test_StringSlice_SplitTrimmedNonEmpty(t *testing.T) {
	// Arrange
	result := stringslice.SplitTrimmedNonEmpty("a, b, , c", ",", -1)

	// Act
	actual := args.Map{"result": len(result) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3, got:", actual)
}

// ==========================================
// FirstLastDefaultStatus / FirstLastStatus
// ==========================================

func Test_StringSlice_FirstLastDefaultStatus_NonEmpty(t *testing.T) {
	// Arrange
	status := stringslice.FirstLastDefaultStatus([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"result": status.First != "a" || status.Last != "c" || !status.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected: first= last= isValid=", actual)
}

func Test_StringSlice_FirstLastDefaultStatus_Empty(t *testing.T) {
	// Arrange
	status := stringslice.FirstLastDefaultStatus([]string{})

	// Act
	actual := args.Map{"result": status.First != "" || status.Last != "" || status.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty strings and false", actual)
}

// ==========================================
// NonNullStrings
// ==========================================

func Test_StringSlice_NonNullStrings(t *testing.T) {
	// Arrange
	result := stringslice.NonNullStrings([]string{"a", "", "b"})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items (filters empty strings)", actual)
}

// ==========================================
// TrimmedEachWords
// ==========================================

func Test_StringSlice_TrimmedEachWords(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords([]string{" a ", " b "})

	// Act
	actual := args.Map{"result": result[0] != "a" || result[1] != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [a b]", actual)
}

func Test_StringSlice_TrimmedEachWords_Nil(t *testing.T) {
	// Arrange
	result := stringslice.TrimmedEachWords(nil)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty slice", actual)
}

// ==========================================
// SafeIndexAtWith
// ==========================================

func Test_StringSlice_SafeIndexAtWith_Valid(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")

	// Act
	actual := args.Map{"result": result != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'b', got ''", actual)
}

func Test_StringSlice_SafeIndexAtWith_OutOfBounds(t *testing.T) {
	// Arrange
	result := stringslice.SafeIndexAtWith([]string{"a"}, 5, "def")

	// Act
	actual := args.Map{"result": result != "def"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'def', got ''", actual)
}

// ==========================================
// FirstOrDefaultWith
// ==========================================

func Test_StringSlice_FirstOrDefaultWith_NonEmpty(t *testing.T) {
	// Arrange
	result, _ := stringslice.FirstOrDefaultWith([]string{"x"}, "def")

	// Act
	actual := args.Map{"result": result != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'x', got ''", actual)
}

func Test_StringSlice_FirstOrDefaultWith_Empty(t *testing.T) {
	// Arrange
	result, _ := stringslice.FirstOrDefaultWith([]string{}, "def")

	// Act
	actual := args.Map{"result": result != "def"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'def', got ''", actual)
}
