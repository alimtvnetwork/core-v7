package corestrtests

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── SimpleSlice: Add/Adds/Append/AddIf/AddsIf ──

func Test_Cov37_Add(t *testing.T) {
	safeTest(t, "Test_Cov37_Add", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.Add("a")
		s.Add("b")
		if s.Length() != 2 {
			t.Errorf("expected 2, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddSplit(t *testing.T) {
	safeTest(t, "Test_Cov37_AddSplit", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddSplit("a,b,c", ",")
		if s.Length() != 3 {
			t.Errorf("expected 3, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddIf_True", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddIf(true, "yes")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddIf_False", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddIf(false, "no")
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

func Test_Cov37_Adds(t *testing.T) {
	safeTest(t, "Test_Cov37_Adds", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.Adds("a", "b", "c")
		if s.Length() != 3 {
			t.Errorf("expected 3, got %d", s.Length())
		}
	})
}

func Test_Cov37_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Adds_Empty", func() {
		s := corestr.New.SimpleSlice.Lines("x")
		s.Adds()
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_Append(t *testing.T) {
	safeTest(t, "Test_Cov37_Append", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.Append("a", "b")
		if s.Length() != 2 {
			t.Errorf("expected 2, got %d", s.Length())
		}
	})
}

func Test_Cov37_Append_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Append_Empty", func() {
		s := corestr.New.SimpleSlice.Lines("x")
		s.Append()
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddsIf_True", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddsIf(true, "a", "b")
		if s.Length() != 2 {
			t.Errorf("expected 2, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddsIf_False", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddsIf(false, "a", "b")
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

// ── AppendFmt / AppendFmtIf ──

func Test_Cov37_AppendFmt(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmt", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmt("hello %s", "world")
		if s.First() != "hello world" {
			t.Errorf("unexpected: %s", s.First())
		}
	})
}

func Test_Cov37_AppendFmt_EmptySkip(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmt_EmptySkip", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmt("")
		// empty format with no values still appends (fmt.Sprintf("") == "")
		// Actually the code checks: format == "" && len(v) == 0 → skip
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

func Test_Cov37_AppendFmtIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmtIf_True", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmtIf(true, "val=%d", 42)
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AppendFmtIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmtIf_False", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmtIf(false, "val=%d", 42)
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

func Test_Cov37_AppendFmtIf_EmptyFormat(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendFmtIf_EmptyFormat", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AppendFmtIf(true, "")
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

// ── AddAsTitleValue / AddAsCurlyTitleWrap / If variants ──

func Test_Cov37_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsTitleValue", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsTitleValue("Key", "Val")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddAsTitleValueIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsTitleValueIf_True", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsTitleValueIf(true, "K", "V")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddAsTitleValueIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsTitleValueIf_False", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsTitleValueIf(false, "K", "V")
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsCurlyTitleWrap", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsCurlyTitleWrap("K", "V")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddAsCurlyTitleWrapIf_True(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsCurlyTitleWrapIf_True", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsCurlyTitleWrapIf(true, "K", "V")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddAsCurlyTitleWrapIf_False(t *testing.T) {
	safeTest(t, "Test_Cov37_AddAsCurlyTitleWrapIf_False", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddAsCurlyTitleWrapIf(false, "K", "V")
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

// ── InsertAt ──

func Test_Cov37_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov37_InsertAt_Middle", func() {
		s := corestr.New.SimpleSlice.Lines("a", "c")
		s.InsertAt(1, "b")
		if s.Length() != 3 || (*s)[1] != "b" {
			t.Errorf("unexpected: %v", *s)
		}
	})
}

func Test_Cov37_InsertAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov37_InsertAt_OutOfRange", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.InsertAt(-1, "x")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
		s.InsertAt(99, "x")
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_InsertAt_End(t *testing.T) {
	safeTest(t, "Test_Cov37_InsertAt_End", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.InsertAt(2, "c")
		if s.Length() != 3 || s.Last() != "c" {
			t.Errorf("unexpected: %v", *s)
		}
	})
}

// ── AddStruct / AddPointer ──

func Test_Cov37_AddStruct(t *testing.T) {
	safeTest(t, "Test_Cov37_AddStruct", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddStruct(true, struct{ Name string }{Name: "test"})
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddStruct_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddStruct_Nil", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddStruct(true, nil)
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

func Test_Cov37_AddPointer_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddPointer_Nil", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddPointer(true, nil)
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

// ── AddError ──

func Test_Cov37_AddError_NonNil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddError_NonNil", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddError(errors.New("test error"))
	})
}

func Test_Cov37_AddError_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddError_Nil", func() {
		s := corestr.New.SimpleSlice.Empty()
		s.AddError(nil)
		if s.Length() != 0 {
			t.Errorf("expected 0, got %d", s.Length())
		}
	})
}

// ── AsError / AsDefaultError ──

func Test_Cov37_AsError_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_AsError_NonEmpty", func() {
		s := corestr.New.SimpleSlice.Lines("err1", "err2")
		err := s.AsError("; ")
		if err == nil || !strings.Contains(err.Error(), "err1") {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func Test_Cov37_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_AsError_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.AsError("; ") != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov37_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_Cov37_AsDefaultError", func() {
		s := corestr.New.SimpleSlice.Lines("e1")
		err := s.AsDefaultError()
		if err == nil {
			t.Error("expected error")
		}
	})
}

// ── First/Last/FirstOrDefault/LastOrDefault ──

func Test_Cov37_First(t *testing.T) {
	safeTest(t, "Test_Cov37_First", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.First() != "a" {
			t.Errorf("expected a")
		}
	})
}

func Test_Cov37_Last(t *testing.T) {
	safeTest(t, "Test_Cov37_Last", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.Last() != "b" {
			t.Errorf("expected b")
		}
	})
}

func Test_Cov37_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("x")
		val := s.FirstDynamic()
		if val != "x" {
			t.Errorf("unexpected: %v", val)
		}
	})
}

func Test_Cov37_LastDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_LastDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("x", "y")
		val := s.LastDynamic()
		if val != "y" {
			t.Errorf("unexpected: %v", val)
		}
	})
}

func Test_Cov37_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstOrDefault_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.FirstOrDefault() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstOrDefault_NonEmpty", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		if s.FirstOrDefault() != "a" {
			t.Error("expected a")
		}
	})
}

func Test_Cov37_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_FirstOrDefaultDynamic", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.FirstOrDefaultDynamic() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_LastOrDefault_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.LastOrDefault() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_LastOrDefault_NonEmpty", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.LastOrDefault() != "b" {
			t.Error("expected b")
		}
	})
}

func Test_Cov37_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_LastOrDefaultDynamic", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.LastOrDefaultDynamic() != "" {
			t.Error("expected empty")
		}
	})
}

// ── Skip/Take/Limit ──

func Test_Cov37_Skip(t *testing.T) {
	safeTest(t, "Test_Cov37_Skip", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := s.Skip(1)
		if len(result) != 2 {
			t.Errorf("expected 2, got %d", len(result))
		}
	})
}

func Test_Cov37_Skip_ExceedsLen(t *testing.T) {
	safeTest(t, "Test_Cov37_Skip_ExceedsLen", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.Skip(5)
		if len(result) != 0 {
			t.Errorf("expected 0, got %d", len(result))
		}
	})
}

func Test_Cov37_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_SkipDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.SkipDynamic(1)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_SkipDynamic_ExceedsLen(t *testing.T) {
	safeTest(t, "Test_Cov37_SkipDynamic_ExceedsLen", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.SkipDynamic(5)
		slice, ok := result.([]string)
		if !ok || len(slice) != 0 {
			t.Errorf("expected empty slice")
		}
	})
}

func Test_Cov37_Take(t *testing.T) {
	safeTest(t, "Test_Cov37_Take", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := s.Take(2)
		if len(result) != 2 {
			t.Errorf("expected 2, got %d", len(result))
		}
	})
}

func Test_Cov37_Take_ExceedsLen(t *testing.T) {
	safeTest(t, "Test_Cov37_Take_ExceedsLen", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.Take(5)
		if len(result) != 1 {
			t.Errorf("expected 1, got %d", len(result))
		}
	})
}

func Test_Cov37_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_TakeDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.TakeDynamic(1)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_Limit(t *testing.T) {
	safeTest(t, "Test_Cov37_Limit", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := s.Limit(2)
		if len(result) != 2 {
			t.Errorf("expected 2, got %d", len(result))
		}
	})
}

func Test_Cov37_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_Cov37_LimitDynamic", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.LimitDynamic(1)
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

// ── Length/Count/CountFunc/IsEmpty/HasAnyItem/LastIndex/HasIndex ──

func Test_Cov37_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Length_Nil", func() {
		var s *corestr.SimpleSlice
		if s.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov37_Count(t *testing.T) {
	safeTest(t, "Test_Cov37_Count", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.Count() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov37_CountFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := s.CountFunc(func(i int, item string) bool {
			return len(item) > 1
		})
		if count != 2 {
			t.Errorf("expected 2, got %d", count)
		}
	})
}

func Test_Cov37_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_CountFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		count := s.CountFunc(func(i int, item string) bool { return true })
		if count != 0 {
			t.Errorf("expected 0, got %d", count)
		}
	})
}

func Test_Cov37_IsEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEmpty_True", func() {
		s := corestr.New.SimpleSlice.Empty()
		if !s.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov37_HasAnyItem", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		if !s.HasAnyItem() {
			t.Error("expected HasAnyItem")
		}
	})
}

func Test_Cov37_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov37_LastIndex", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.LastIndex() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov37_HasIndex", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if !s.HasIndex(1) {
			t.Error("expected true")
		}
		if s.HasIndex(2) {
			t.Error("expected false")
		}
		if s.HasIndex(-1) {
			t.Error("expected false for negative")
		}
	})
}

// ── IsContains / IsContainsFunc / IndexOf / IndexOfFunc ──

func Test_Cov37_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContains", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if !s.IsContains("a") {
			t.Error("expected true")
		}
		if s.IsContains("z") {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContains_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.IsContains("a") {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContainsFunc", func() {
		s := corestr.New.SimpleSlice.Lines("Hello", "World")
		found := s.IsContainsFunc("hello", func(item, searching string) bool {
			return strings.EqualFold(item, searching)
		})
		if !found {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsContainsFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		found := s.IsContainsFunc("x", func(item, searching string) bool { return true })
		if found {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOf", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		if s.IndexOf("b") != 1 {
			t.Errorf("expected 1")
		}
		if s.IndexOf("z") != -1 {
			t.Errorf("expected -1")
		}
	})
}

func Test_Cov37_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOf_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.IndexOf("x") != -1 {
			t.Errorf("expected -1")
		}
	})
}

func Test_Cov37_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOfFunc", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		idx := s.IndexOfFunc("B", func(item, searching string) bool {
			return strings.EqualFold(item, searching)
		})
		if idx != 1 {
			t.Errorf("expected 1, got %d", idx)
		}
	})
}

func Test_Cov37_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IndexOfFunc_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		idx := s.IndexOfFunc("x", func(item, searching string) bool { return true })
		if idx != -1 {
			t.Errorf("expected -1")
		}
	})
}

// ── Strings/List/Hashset ──

func Test_Cov37_Strings(t *testing.T) {
	safeTest(t, "Test_Cov37_Strings", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if len(s.Strings()) != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_List(t *testing.T) {
	safeTest(t, "Test_Cov37_List", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		if len(s.List()) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov37_Hashset", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "a")
		hs := s.Hashset()
		if hs.Length() != 2 {
			t.Errorf("expected 2 unique, got %d", hs.Length())
		}
	})
}

// ── Wrap variants ──

func Test_Cov37_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapDoubleQuote", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapDoubleQuote()
		if !strings.Contains(result.First(), "\"") {
			t.Errorf("expected double quotes")
		}
	})
}

func Test_Cov37_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapSingleQuote", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapSingleQuote()
		if !strings.Contains(result.First(), "'") {
			t.Errorf("expected single quotes")
		}
	})
}

func Test_Cov37_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapTildaQuote", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapTildaQuote()
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapDoubleQuoteIfMissing", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapDoubleQuoteIfMissing()
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov37_WrapSingleQuoteIfMissing", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.WrapSingleQuoteIfMissing()
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── Transpile / TranspileJoin ──

func Test_Cov37_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov37_Transpile", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.Transpile(strings.ToUpper)
		if result.First() != "A" {
			t.Errorf("expected A, got %s", result.First())
		}
	})
}

func Test_Cov37_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Transpile_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		result := s.Transpile(strings.ToUpper)
		if !result.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_Cov37_TranspileJoin", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.TranspileJoin(strings.ToUpper, ",")
		if result != "A,B" {
			t.Errorf("expected A,B, got %s", result)
		}
	})
}

// ── Join variants ──

func Test_Cov37_Join(t *testing.T) {
	safeTest(t, "Test_Cov37_Join", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.Join(",") != "a,b" {
			t.Errorf("unexpected")
		}
	})
}

func Test_Cov37_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Join_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.Join(",") != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLine", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinLine()
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov37_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLine_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.JoinLine() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_JoinLineEofLine_NoSuffix(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLineEofLine_NoSuffix", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinLineEofLine()
		if !strings.HasSuffix(result, "\n") {
			t.Error("expected newline suffix")
		}
	})
}

func Test_Cov37_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLineEofLine_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.JoinLineEofLine() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_JoinSpace(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinSpace", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.JoinSpace() != "a b" {
			t.Errorf("unexpected: %s", s.JoinSpace())
		}
	})
}

func Test_Cov37_JoinComma(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinComma", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.JoinComma() != "a,b" {
			t.Errorf("unexpected")
		}
	})
}

func Test_Cov37_JoinCsv(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsv", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinCsv()
		if !strings.Contains(result, "\"a\"") {
			t.Errorf("expected quoted, got %s", result)
		}
	})
}

func Test_Cov37_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsvLine", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinCsvLine()
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_Cov37_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsvString", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.JoinCsvString(",")
		if !strings.Contains(result, "\"a\"") {
			t.Errorf("expected quoted, got %s", result)
		}
	})
}

func Test_Cov37_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinCsvString_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.JoinCsvString(",") != "" {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_JoinWith(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinWith", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.JoinWith(", ")
		if !strings.HasPrefix(result, ", ") {
			t.Errorf("expected prefix, got %s", result)
		}
	})
}

func Test_Cov37_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinWith_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.JoinWith(",") != "" {
			t.Error("expected empty")
		}
	})
}

// ── EachItemSplitBy ──

func Test_Cov37_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov37_EachItemSplitBy", func() {
		s := corestr.New.SimpleSlice.Lines("a,b", "c,d")
		result := s.EachItemSplitBy(",")
		if result.Length() != 4 {
			t.Errorf("expected 4, got %d", result.Length())
		}
	})
}

// ── PrependJoin / AppendJoin / PrependAppend ──

func Test_Cov37_PrependJoin(t *testing.T) {
	safeTest(t, "Test_Cov37_PrependJoin", func() {
		s := corestr.New.SimpleSlice.Lines("b", "c")
		result := s.PrependJoin(",", "a")
		if result != "a,b,c" {
			t.Errorf("expected a,b,c, got %s", result)
		}
	})
}

func Test_Cov37_AppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov37_AppendJoin", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		result := s.AppendJoin(",", "c")
		if result != "a,b,c" {
			t.Errorf("expected a,b,c, got %s", result)
		}
	})
}

func Test_Cov37_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Cov37_PrependAppend", func() {
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend([]string{"a"}, []string{"c"})
		if s.Length() != 3 {
			t.Errorf("expected 3, got %d", s.Length())
		}
	})
}

func Test_Cov37_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_PrependAppend_Empty", func() {
		s := corestr.New.SimpleSlice.Lines("b")
		s.PrependAppend(nil, nil)
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

// ── IsEqual / IsEqualLines / IsEqualUnorderedLines / IsEqualUnorderedLinesClone ──

func Test_Cov37_IsEqual_Same(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_Same", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		if !a.IsEqual(b) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_BothNil", func() {
		var a, b *corestr.SimpleSlice
		if !a.IsEqual(b) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_OneNil", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		if a.IsEqual(nil) {
			t.Error("expected not equal")
		}
	})
}

func Test_Cov37_IsEqual_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_DiffLen", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		if a.IsEqual(b) {
			t.Error("expected not equal")
		}
	})
}

func Test_Cov37_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqual_BothEmpty", func() {
		a := corestr.New.SimpleSlice.Empty()
		b := corestr.New.SimpleSlice.Empty()
		if !a.IsEqual(b) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqualLines_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualLines_Mismatch", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		if a.IsEqualLines([]string{"a", "c"}) {
			t.Error("expected not equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLines", func() {
		a := corestr.New.SimpleSlice.Lines("b", "a")
		if !a.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLines_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLines_DiffLen", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		if a.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Error("expected not equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLines_BothNil", func() {
		var a *corestr.SimpleSlice
		if !a.IsEqualUnorderedLines(nil) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone", func() {
		a := corestr.New.SimpleSlice.Lines("b", "a")
		if !a.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone_DiffLen", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		if a.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Error("expected not equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone_BothNil", func() {
		var a *corestr.SimpleSlice
		if !a.IsEqualUnorderedLinesClone(nil) {
			t.Error("expected equal")
		}
	})
}

func Test_Cov37_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualUnorderedLinesClone_BothEmpty", func() {
		a := corestr.New.SimpleSlice.Empty()
		if !a.IsEqualUnorderedLinesClone([]string{}) {
			t.Error("expected equal")
		}
	})
}

// ── IsEqualByFunc / IsEqualByFuncLinesSplit ──

func Test_Cov37_IsEqualByFunc_Match(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_Match", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		result := a.IsEqualByFunc(func(i int, l, r string) bool {
			return strings.EqualFold(l, r)
		}, "A", "B")
		if !result {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsEqualByFunc_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_DiffLen", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		if a.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b") {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_Empty", func() {
		a := corestr.New.SimpleSlice.Empty()
		if !a.IsEqualByFunc(func(i int, l, r string) bool { return true }) {
			t.Error("expected true for both empty")
		}
	})
}

func Test_Cov37_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFunc_Mismatch", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		if a.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "X") {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_Match(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_Match", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		result := a.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool {
			return l == r
		})
		if !result {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_Trim", func() {
		a := corestr.New.SimpleSlice.Lines(" a ", " b ")
		result := a.IsEqualByFuncLinesSplit(true, ",", "a,b", func(i int, l, r string) bool {
			return l == r
		})
		if !result {
			t.Error("expected true with trim")
		}
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_DiffLen", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		if a.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true }) {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsEqualByFuncLinesSplit_Empty", func() {
		a := corestr.New.SimpleSlice.Empty()
		if a.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }) {
			// "" split by "," yields [""] which has len 1, not 0.
			t.Error("expected false for empty slice vs split-empty string")
		}
	})
}

// ── Collection / ToCollection / NonPtr / Ptr ──

func Test_Cov37_Collection(t *testing.T) {
	safeTest(t, "Test_Cov37_Collection", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		col := s.Collection(false)
		if col.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov37_ToCollection", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		col := s.ToCollection(true)
		if col.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_NonPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		np := s.NonPtr()
		if np.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov37_Ptr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		p := s.Ptr()
		if p.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_ToPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_ToPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		p := s.ToPtr()
		if p == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_ToNonPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		np := s.ToNonPtr()
		if np.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── String ──

func Test_Cov37_String(t *testing.T) {
	safeTest(t, "Test_Cov37_String", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if !strings.Contains(s.String(), "a") {
			t.Error("expected a")
		}
	})
}

func Test_Cov37_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_String_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if s.String() != "" {
			t.Error("expected empty")
		}
	})
}

// ── ConcatNew / ConcatNewStrings / ConcatNewSimpleSlices ──

func Test_Cov37_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNew", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNew("b", "c")
		if result.Length() != 3 {
			t.Errorf("expected 3, got %d", result.Length())
		}
	})
}

func Test_Cov37_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNewStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.ConcatNewStrings("b")
		if len(result) != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_ConcatNewStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNewStrings_Nil", func() {
		var s *corestr.SimpleSlice
		result := s.ConcatNewStrings("a")
		if len(result) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov37_ConcatNewSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Lines("a")
		s2 := corestr.New.SimpleSlice.Lines("b")
		result := s1.ConcatNewSimpleSlices(s2)
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

// ── CsvStrings ──

func Test_Cov37_CsvStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_CsvStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		csv := s.CsvStrings()
		if len(csv) != 1 || !strings.Contains(csv[0], "\"") {
			t.Errorf("expected quoted")
		}
	})
}

func Test_Cov37_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_CsvStrings_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		csv := s.CsvStrings()
		if len(csv) != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ── Sort / Reverse ──

func Test_Cov37_Sort(t *testing.T) {
	safeTest(t, "Test_Cov37_Sort", func() {
		s := corestr.New.SimpleSlice.Lines("c", "a", "b")
		s.Sort()
		if s.First() != "a" {
			t.Errorf("expected a first")
		}
	})
}

func Test_Cov37_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov37_Reverse", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		s.Reverse()
		if s.First() != "c" || s.Last() != "a" {
			t.Errorf("unexpected: %v", *s)
		}
	})
}

func Test_Cov37_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_Cov37_Reverse_Two", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Reverse()
		if s.First() != "b" {
			t.Errorf("expected b first")
		}
	})
}

func Test_Cov37_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Cov37_Reverse_Single", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.Reverse()
		if s.First() != "a" {
			t.Errorf("expected a")
		}
	})
}

// ── JSON / Serialize / Deserialize ──

func Test_Cov37_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov37_MarshalJSON", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		data, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), "\"a\"") {
			t.Errorf("unexpected: %s", data)
		}
	})
}

func Test_Cov37_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov37_UnmarshalJSON", func() {
		s := corestr.New.SimpleSlice.Empty()
		err := json.Unmarshal([]byte(`["x","y"]`), s)
		if err != nil {
			t.Fatal(err)
		}
		if s.Length() != 2 {
			t.Errorf("expected 2, got %d", s.Length())
		}
	})
}

func Test_Cov37_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov37_UnmarshalJSON_Invalid", func() {
		s := corestr.New.SimpleSlice.Empty()
		err := json.Unmarshal([]byte(`not-json`), s)
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_Cov37_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonModel", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		model := s.JsonModel()
		if len(model) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonModelAny", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		model := s.JsonModelAny()
		if model == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_Json(t *testing.T) {
	safeTest(t, "Test_Cov37_Json", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.Json()
		if result.HasError() {
			t.Errorf("unexpected error: %v", result.Error)
		}
	})
}

func Test_Cov37_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonPtr", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.JsonPtr()
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov37_Serialize", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		data, err := s.Serialize()
		if err != nil {
			t.Fatal(err)
		}
		if len(data) == 0 {
			t.Error("expected non-empty bytes")
		}
	})
}

func Test_Cov37_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov37_Deserialize", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		var target []string
		err := s.Deserialize(&target)
		if err != nil {
			t.Fatal(err)
		}
		if len(target) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov37_ParseInjectUsingJson", func() {
		s := corestr.New.SimpleSlice.Empty()
		src := corestr.New.SimpleSlice.Lines("a", "b")
		jsonResult := src.JsonPtr()
		result, err := s.ParseInjectUsingJson(jsonResult)
		if err != nil {
			t.Fatal(err)
		}
		if result.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov37_ParseInjectUsingJsonMust", func() {
		s := corestr.New.SimpleSlice.Empty()
		src := corestr.New.SimpleSlice.Lines("a")
		jsonResult := src.JsonPtr()
		result := s.ParseInjectUsingJsonMust(jsonResult)
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov37_JsonParseSelfInject", func() {
		s := corestr.New.SimpleSlice.Empty()
		src := corestr.New.SimpleSlice.Lines("x")
		err := s.JsonParseSelfInject(src.JsonPtr())
		if err != nil {
			t.Fatal(err)
		}
	})
}

// ── AsJsoner / AsJsonContractsBinder / AsJsonParseSelfInjector / AsJsonMarshaller ──

func Test_Cov37_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsoner", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsoner()
		if j == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsonContractsBinder", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsonContractsBinder()
		if j == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsonParseSelfInjector", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsonParseSelfInjector()
		if j == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov37_AsJsonMarshaller", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		j := s.AsJsonMarshaller()
		if j == nil {
			t.Error("expected non-nil")
		}
	})
}

// ── Clear / Dispose ──

func Test_Cov37_Clear(t *testing.T) {
	safeTest(t, "Test_Cov37_Clear", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		s.Clear()
		if s.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov37_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Clear_Nil", func() {
		var s *corestr.SimpleSlice
		result := s.Clear()
		if result != nil {
			t.Error("expected nil")
		}
	})
}

func Test_Cov37_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov37_Dispose", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		s.Dispose()
		if s.Length() != 0 {
			t.Errorf("expected 0 after dispose")
		}
	})
}
func Test_Cov37_Clone_Deep(t *testing.T) {
	safeTest(t, "Test_Cov37_Clone_Deep", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		cloned := s.Clone(true)
		if cloned.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Clone_Shallow(t *testing.T) {
	safeTest(t, "Test_Cov37_Clone_Shallow", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.Clone(false)
		if cloned.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}
	safeTest(t, "Test_Cov37_DeepClone", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.DeepClone()
		if cloned.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_ShadowClone(t *testing.T) {
	safeTest(t, "Test_Cov37_ShadowClone", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		cloned := s.ShadowClone()
		if cloned.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── IsDistinctEqual / IsDistinctEqualRaw ──

func Test_Cov37_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_Cov37_IsDistinctEqualRaw", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "a")
		if !s.IsDistinctEqualRaw("a", "b") {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_Cov37_IsDistinctEqual", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("b", "a")
		if !a.IsDistinctEqual(b) {
			t.Error("expected true")
		}
	})
}

// ── IsUnorderedEqualRaw / IsUnorderedEqual ──

func Test_Cov37_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_Clone", func() {
		s := corestr.New.SimpleSlice.Lines("b", "a")
		if !s.IsUnorderedEqualRaw(true, "a", "b") {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_NoClone", func() {
		s := corestr.New.SimpleSlice.Lines("b", "a")
		if !s.IsUnorderedEqualRaw(false, "a", "b") {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsUnorderedEqualRaw_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_DiffLen", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		if s.IsUnorderedEqualRaw(true, "a", "b") {
			t.Error("expected false")
		}
	})
}

func Test_Cov37_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqualRaw_BothEmpty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if !s.IsUnorderedEqualRaw(true) {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqual_BothEmpty", func() {
		a := corestr.New.SimpleSlice.Empty()
		b := corestr.New.SimpleSlice.Empty()
		if !a.IsUnorderedEqual(true, b) {
			t.Error("expected true")
		}
	})
}

func Test_Cov37_IsUnorderedEqual_RightNil(t *testing.T) {
	safeTest(t, "Test_Cov37_IsUnorderedEqual_RightNil", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		if a.IsUnorderedEqual(true, nil) {
			t.Error("expected false")
		}
	})
}

// ── DistinctDiffRaw / DistinctDiff ──

func Test_Cov37_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiffRaw_BothNil", func() {
		var s *corestr.SimpleSlice
		result := s.DistinctDiffRaw()
		if len(result) != 0 {
			t.Errorf("expected empty")
		}
	})
}

func Test_Cov37_DistinctDiffRaw_LeftNilRightNot(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiffRaw_LeftNilRightNot", func() {
		var s *corestr.SimpleSlice
		result := s.DistinctDiffRaw("a")
		if len(result) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiffRaw_RightNil", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		result := s.DistinctDiffRaw()
		if len(result) == 0 {
			t.Errorf("expected non-empty")
		}
	})
}

func Test_Cov37_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiff_BothNil", func() {
		var a *corestr.SimpleSlice
		result := a.DistinctDiff(nil)
		if len(result) != 0 {
			t.Errorf("expected empty")
		}
	})
}

func Test_Cov37_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiff_LeftNil", func() {
		var a *corestr.SimpleSlice
		b := corestr.New.SimpleSlice.Lines("x")
		result := a.DistinctDiff(b)
		if len(result) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_Cov37_DistinctDiff_RightNil", func() {
		a := corestr.New.SimpleSlice.Lines("x")
		result := a.DistinctDiff(nil)
		if len(result) != 1 {
			t.Errorf("expected 1")
		}
	})
}

// ── AddedRemovedLinesDiff ──

func Test_Cov37_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_Cov37_AddedRemovedLinesDiff", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := s.AddedRemovedLinesDiff("b", "c")
		if len(added) == 0 {
			t.Error("expected added items")
		}
		if len(removed) == 0 {
			t.Error("expected removed items")
		}
	})
}

func Test_Cov37_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov37_AddedRemovedLinesDiff_BothNil", func() {
		var s *corestr.SimpleSlice
		added, removed := s.AddedRemovedLinesDiff()
		if added != nil || removed != nil {
			t.Error("expected both nil")
		}
	})
}

// ── RemoveIndexes ──

func Test_Cov37_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_Cov37_RemoveIndexes", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := s.RemoveIndexes(1)
		if err != nil {
			t.Fatal(err)
		}
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_Cov37_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_RemoveIndexes_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		_, err := s.RemoveIndexes(0)
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_Cov37_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_Cov37_RemoveIndexes_InvalidIndex", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		_, err := s.RemoveIndexes(5)
		if err == nil {
			t.Error("expected error for invalid index")
		}
	})
}

// ── SafeStrings ──

func Test_Cov37_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_SafeStrings", func() {
		s := corestr.New.SimpleSlice.Lines("a")
		if len(s.SafeStrings()) != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_SafeStrings_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if len(s.SafeStrings()) != 0 {
			t.Errorf("expected 0")
		}
	})
}

// ── newSimpleSliceCreator factory methods ──

func Test_Cov37_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Cap", func() {
		s := corestr.New.SimpleSlice.Cap(5)
		if s.Length() != 0 {
			t.Errorf("expected 0")
		}
	})
}

func Test_Cov37_Creator_Cap_Zero(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Cap_Zero", func() {
		s := corestr.New.SimpleSlice.Cap(0)
		if s == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_Cov37_Creator_Default(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Default", func() {
		s := corestr.New.SimpleSlice.Default()
		if s == nil || s.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_Creator_Lines(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Lines", func() {
		s := corestr.New.SimpleSlice.Lines("a", "b")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_Split(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Split", func() {
		s := corestr.New.SimpleSlice.Split("a,b", ",")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_SplitLines(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_SplitLines", func() {
		s := corestr.New.SimpleSlice.SplitLines("a\nb")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_SpreadStrings", func() {
		s := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Hashset", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		s := corestr.New.SimpleSlice.Hashset(hs)
		if s.Length() != 2 {
			t.Errorf("expected 2, got %d", s.Length())
		}
	})
}

func Test_Cov37_Creator_Hashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Hashset_Empty", func() {
		hs := corestr.New.Hashset.Empty()
		s := corestr.New.SimpleSlice.Hashset(hs)
		if !s.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_Creator_Create(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Create", func() {
		s := corestr.New.SimpleSlice.Create([]string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_Strings(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Strings", func() {
		s := corestr.New.SimpleSlice.Strings([]string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}
	safeTest(t, "Test_Cov37_Creator_StringsOptions_Clone", func() {
		s := corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_StringsOptions_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsOptions_NoClone", func() {
		s := corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_StringsOptions_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsOptions_Empty", func() {
		s := corestr.New.SimpleSlice.StringsOptions(false, []string{})
		if !s.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_Creator_StringsClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsClone", func() {
		s := corestr.New.SimpleSlice.StringsClone([]string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_StringsClone_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_StringsClone_Nil", func() {
		s := corestr.New.SimpleSlice.StringsClone(nil)
		if !s.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_Creator_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Direct_Clone", func() {
		s := corestr.New.SimpleSlice.Direct(true, []string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Direct_NoClone", func() {
		s := corestr.New.SimpleSlice.Direct(false, []string{"a"})
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Direct_Nil", func() {
		s := corestr.New.SimpleSlice.Direct(true, nil)
		if !s.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_Creator_UsingLines_Clone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingLines_Clone", func() {
		s := corestr.New.SimpleSlice.UsingLines(true, "a", "b")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_UsingLines_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingLines_NoClone", func() {
		s := corestr.New.SimpleSlice.UsingLines(false, "a")
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_UsingSeparatorLine(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingSeparatorLine", func() {
		s := corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_UsingLine(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_UsingLine", func() {
		s := corestr.New.SimpleSlice.UsingLine("a\nb")
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Empty", func() {
		s := corestr.New.SimpleSlice.Empty()
		if !s.IsEmpty() {
			t.Error("expected empty")
		}
	})
}

func Test_Cov37_Creator_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Deserialize", func() {
		data, _ := json.Marshal([]string{"a", "b"})
		s, err := corestr.New.SimpleSlice.Deserialize(data)
		if err != nil {
			t.Fatal(err)
		}
		if s.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_Cov37_Creator_Deserialize_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Deserialize_Invalid", func() {
		_, err := corestr.New.SimpleSlice.Deserialize([]byte("bad"))
		if err == nil {
			t.Error("expected error")
		}
	})
}

func Test_Cov37_Creator_DeserializeJsoner(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_DeserializeJsoner", func() {
		src := corestr.New.SimpleSlice.Lines("a")
		jsoner := src.AsJsoner()
		s, err := corestr.New.SimpleSlice.DeserializeJsoner(jsoner)
		if err != nil {
			t.Fatal(err)
		}
		if s.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_Cov37_Creator_Map(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_Map", func() {
		m := map[string]string{"a": "1", "b": "2"}
		s := corestr.New.SimpleSlice.Map(m)
		if s.Length() != 2 {
			t.Errorf("expected 2, got %d", s.Length())
		}
	})
}

func Test_Cov37_Creator_ByLen(t *testing.T) {
	safeTest(t, "Test_Cov37_Creator_ByLen", func() {
		s := corestr.New.SimpleSlice.ByLen([]string{"a", "b", "c"})
		if s == nil {
			t.Error("expected non-nil")
		}
	})
}

// Fix the broken AddError test
func Test_Cov37_AddError_WithError(t *testing.T) {
	safeTest(t, "Test_Cov37_AddError_WithError", func() {
		s := corestr.New.SimpleSlice.Empty()
		err := json.Unmarshal([]byte("bad"), &struct{}{})
		s.AddError(err)
		if s.Length() != 1 {
			t.Errorf("expected 1, got %d", s.Length())
		}
	})
}

// ── JoinLineEofLine with already-suffixed ──

func Test_Cov37_JoinLineEofLine_AlreadySuffixed(t *testing.T) {
	safeTest(t, "Test_Cov37_JoinLineEofLine_AlreadySuffixed", func() {
		s := corestr.New.SimpleSlice.Lines("a\n")
		result := s.JoinLineEofLine()
		if !strings.HasSuffix(result, "\n") {
			t.Error("expected newline suffix")
		}
	})
}

// Ensure the unused import is used
var _ = corejson.Result{}
