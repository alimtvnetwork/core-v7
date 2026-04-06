package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
)

// ==========================================
// PrependNew
// ==========================================

func Test_StringSlice_PrependNew_BothNonEmpty(t *testing.T) {
	result := stringslice.PrependNew([]string{"c", "d"}, "a", "b")
	r := *result
	if len(r) != 4 {
		t.Errorf("expected 4, got %d", len(r))
	}
	if r[0] != "a" || r[1] != "b" || r[2] != "c" || r[3] != "d" {
		t.Errorf("unexpected order: %v", r)
	}
}

func Test_StringSlice_PrependNew_EmptyPrepend(t *testing.T) {
	result := stringslice.PrependNew([]string{"a", "b"})
	r := *result
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

func Test_StringSlice_PrependNew_EmptySlice(t *testing.T) {
	result := stringslice.PrependNew([]string{}, "x")
	r := *result
	if len(r) != 1 || r[0] != "x" {
		t.Errorf("expected [x], got %v", r)
	}
}

func Test_StringSlice_PrependNew_BothEmpty(t *testing.T) {
	result := stringslice.PrependNew([]string{})
	r := *result
	if len(r) != 0 {
		t.Errorf("expected 0, got %d", len(r))
	}
}

// ==========================================
// AppendLineNew
// ==========================================

func Test_StringSlice_AppendLineNew_Basic(t *testing.T) {
	result := stringslice.AppendLineNew([]string{"a"}, "b")
	if len(result) != 2 || result[1] != "b" {
		t.Errorf("expected [a b], got %v", result)
	}
}

func Test_StringSlice_AppendLineNew_EmptySlice(t *testing.T) {
	result := stringslice.AppendLineNew([]string{}, "x")
	if len(result) != 1 || result[0] != "x" {
		t.Errorf("expected [x], got %v", result)
	}
}

// ==========================================
// PrependLineNew
// ==========================================

func Test_StringSlice_PrependLineNew_Basic(t *testing.T) {
	result := stringslice.PrependLineNew("a", []string{"b", "c"})
	if len(result) != 3 || result[0] != "a" {
		t.Errorf("expected [a b c], got %v", result)
	}
}

// ==========================================
// MakeDefault / Make / MakeLen / MakePtr / MakeLenPtr / MakeDefaultPtr
// ==========================================

func Test_StringSlice_MakeDefault(t *testing.T) {
	result := stringslice.MakeDefault(5)
	if len(result) != 0 || cap(result) != 5 {
		t.Errorf("expected len=0 cap=5, got len=%d cap=%d", len(result), cap(result))
	}
}

func Test_StringSlice_Make(t *testing.T) {
	result := stringslice.Make(3, 5)
	if len(result) != 3 || cap(result) != 5 {
		t.Errorf("expected len=3 cap=5, got len=%d cap=%d", len(result), cap(result))
	}
}

func Test_StringSlice_MakeLen(t *testing.T) {
	result := stringslice.MakeLen(3)
	if len(result) != 3 {
		t.Errorf("expected len=3, got %d", len(result))
	}
}

// ==========================================
// Empty / EmptyPtr
// ==========================================

func Test_StringSlice_Empty(t *testing.T) {
	result := stringslice.Empty()
	if result == nil || len(result) != 0 {
		t.Error("Empty should return non-nil empty slice")
	}
}

func Test_StringSlice_EmptyPtr(t *testing.T) {
	result := stringslice.EmptyPtr()
	if len(result) != 0 {
		t.Error("EmptyPtr should return empty slice")
	}
}

// ==========================================
// IsEmpty / IsEmptyPtr / HasAnyItem / HasAnyItemPtr
// ==========================================

func Test_StringSlice_IsEmptyPtr_Nil(t *testing.T) {
	if !stringslice.IsEmptyPtr(nil) {
		t.Error("nil ptr should be empty")
	}
}

func Test_StringSlice_IsEmptyPtr_NonNil(t *testing.T) {
	s := []string{"a"}
	if stringslice.IsEmptyPtr(s) {
		t.Error("non-empty should not be empty")
	}
}

func Test_StringSlice_HasAnyItemPtr_Nil(t *testing.T) {
	if stringslice.HasAnyItemPtr(nil) {
		t.Error("nil should not have items")
	}
}

func Test_StringSlice_HasAnyItemPtr_NonNil(t *testing.T) {
	s := []string{"a"}
	if !stringslice.HasAnyItemPtr(s) {
		t.Error("non-empty should have items")
	}
}

// ==========================================
// CloneIf
// ==========================================

func Test_StringSlice_CloneIf_True(t *testing.T) {
	src := []string{"a", "b"}
	result := stringslice.CloneIf(true, 0, src)
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
	result[0] = "z"
	if src[0] == "z" {
		t.Error("CloneIf true should produce independent copy")
	}
}

func Test_StringSlice_CloneIf_False(t *testing.T) {
	src := []string{"a", "b"}
	result := stringslice.CloneIf(false, 0, src)
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

// ==========================================
// ClonePtr
// ==========================================

func Test_StringSlice_ClonePtr_Nil(t *testing.T) {
	result := stringslice.ClonePtr(nil)
	if len(result) != 0 {
		t.Error("nil should return empty slice")
	}
}

func Test_StringSlice_ClonePtr_NonEmpty(t *testing.T) {
	s := []string{"a", "b"}
	result := stringslice.ClonePtr(s)
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

// ==========================================
// MergeNewSimple
// ==========================================

func Test_StringSlice_MergeNewSimple_Basic(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{"a"}, []string{"b", "c"})
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func Test_StringSlice_MergeNewSimple_BothEmpty(t *testing.T) {
	result := stringslice.MergeNewSimple([]string{}, []string{})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// ==========================================
// NonEmptyIf
// ==========================================

func Test_StringSlice_NonEmptyIf_True(t *testing.T) {
	result := stringslice.NonEmptyIf(true, []string{"a", "", "b"})
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

func Test_StringSlice_NonEmptyIf_False(t *testing.T) {
	result := stringslice.NonEmptyIf(false, []string{"a", "", "b"})
	if len(result) != 2 {
		t.Errorf("expected 2 (NonNullStrings filters empty), got %d", len(result))
	}
}

// ==========================================
// NonEmptyJoin / NonEmptyJoinPtr
// ==========================================

func Test_StringSlice_NonEmptyJoin(t *testing.T) {
	result := stringslice.NonEmptyJoin([]string{"a", "", "b"}, ",")
	if result != "a,b" {
		t.Errorf("expected 'a,b', got '%s'", result)
	}
}

// ==========================================
// NonWhitespaceJoin
// ==========================================

func Test_StringSlice_NonWhitespaceJoin(t *testing.T) {
	result := stringslice.NonWhitespaceJoin([]string{"a", "  ", "b"}, ",")
	if result != "a,b" {
		t.Errorf("expected 'a,b', got '%s'", result)
	}
}

// ==========================================
// SlicePtr / LengthOfPointer
// ==========================================

func Test_StringSlice_SlicePtr(t *testing.T) {
	s := []string{"a", "b"}
	result := stringslice.SlicePtr(s)
	if len(result) != 2 {
		t.Error("SlicePtr should return slice")
	}
}

func Test_StringSlice_LengthOfPointer_Nil(t *testing.T) {
	if stringslice.LengthOfPointer(nil) != 0 {
		t.Error("nil should return 0")
	}
}

func Test_StringSlice_LengthOfPointer_NonNil(t *testing.T) {
	s := []string{"a", "b", "c"}
	if stringslice.LengthOfPointer(s) != 3 {
		t.Errorf("expected 3")
	}
}

// ==========================================
// IndexesDefault / SafeIndexes
// ==========================================

func Test_StringSlice_IndexesDefault(t *testing.T) {
	result := stringslice.IndexesDefault([]string{"a", "b", "c"}, 0, 2)
	if len(result) != 2 || result[0] != "a" || result[1] != "c" {
		t.Errorf("unexpected result: %v", result)
	}
}

// ==========================================
// SplitTrimmedNonEmpty
// ==========================================

func Test_StringSlice_SplitTrimmedNonEmpty(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmpty("a, b, , c", ",", -1)
	if len(result) != 3 {
		t.Errorf("expected 3, got %d: %v", len(result), result)
	}
}

// ==========================================
// FirstLastDefaultStatus / FirstLastStatus
// ==========================================

func Test_StringSlice_FirstLastDefaultStatus_NonEmpty(t *testing.T) {
	status := stringslice.FirstLastDefaultStatus([]string{"a", "b", "c"})
	if status.First != "a" || status.Last != "c" || !status.IsValid {
		t.Errorf("unexpected: first=%s last=%s isValid=%v", status.First, status.Last, status.IsValid)
	}
}

func Test_StringSlice_FirstLastDefaultStatus_Empty(t *testing.T) {
	status := stringslice.FirstLastDefaultStatus([]string{})
	if status.First != "" || status.Last != "" || status.IsValid {
		t.Error("empty should return empty strings and false")
	}
}

// ==========================================
// NonNullStrings
// ==========================================

func Test_StringSlice_NonNullStrings(t *testing.T) {
	result := stringslice.NonNullStrings([]string{"a", "", "b"})
	if len(result) != 2 {
		t.Errorf("expected 2 items (filters empty strings), got %v", result)
	}
}

// ==========================================
// TrimmedEachWords
// ==========================================

func Test_StringSlice_TrimmedEachWords(t *testing.T) {
	result := stringslice.TrimmedEachWords([]string{" a ", " b "})
	if result[0] != "a" || result[1] != "b" {
		t.Errorf("expected [a b], got %v", result)
	}
}

func Test_StringSlice_TrimmedEachWords_Nil(t *testing.T) {
	result := stringslice.TrimmedEachWords(nil)
	if len(result) != 0 {
		t.Error("nil should return empty slice")
	}
}

// ==========================================
// SafeIndexAtWith
// ==========================================

func Test_StringSlice_SafeIndexAtWith_Valid(t *testing.T) {
	result := stringslice.SafeIndexAtWith([]string{"a", "b"}, 1, "def")
	if result != "b" {
		t.Errorf("expected 'b', got '%s'", result)
	}
}

func Test_StringSlice_SafeIndexAtWith_OutOfBounds(t *testing.T) {
	result := stringslice.SafeIndexAtWith([]string{"a"}, 5, "def")
	if result != "def" {
		t.Errorf("expected 'def', got '%s'", result)
	}
}

// ==========================================
// FirstOrDefaultWith
// ==========================================

func Test_StringSlice_FirstOrDefaultWith_NonEmpty(t *testing.T) {
	result, _ := stringslice.FirstOrDefaultWith([]string{"x"}, "def")
	if result != "x" {
		t.Errorf("expected 'x', got '%s'", result)
	}
}

func Test_StringSlice_FirstOrDefaultWith_Empty(t *testing.T) {
	result, _ := stringslice.FirstOrDefaultWith([]string{}, "def")
	if result != "def" {
		t.Errorf("expected 'def', got '%s'", result)
	}
}
