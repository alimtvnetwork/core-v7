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
func Test_StringSlice_Empty(t *testing.T) {
	result := stringslice.Empty()
	if result == nil || len(result) != 0 {
		t.Error("Empty should return non-nil empty slice")
	}
}
		t.Error("CloneIf true should produce independent copy")
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
