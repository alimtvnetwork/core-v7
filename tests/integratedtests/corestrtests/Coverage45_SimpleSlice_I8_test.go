package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================================================
// SimpleSlice — Core operations
// =============================================================================

func Test_I8_SS01_Add(t *testing.T) {
	safeTest(t, "Test_I8_SS01_Add", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Add("b")
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS02_AddSplit(t *testing.T) {
	safeTest(t, "Test_I8_SS02_AddSplit", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddSplit("a.b.c", ".")
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_I8_SS03_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_SS03_AddIf", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddIf(false, "skip")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddIf(true, "keep")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS04_Adds(t *testing.T) {
	safeTest(t, "Test_I8_SS04_Adds", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Adds("a", "b")
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS05_Append(t *testing.T) {
	safeTest(t, "Test_I8_SS05_Append", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Append("a", "b")
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS06_AppendFmt(t *testing.T) {
	safeTest(t, "Test_I8_SS06_AppendFmt", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AppendFmt("hello %s", "world")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty format
		ss.AppendFmt("")
		if ss.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_I8_SS07_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_I8_SS07_AppendFmtIf", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AppendFmtIf(false, "skip %s", "x")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AppendFmtIf(true, "keep %s", "x")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS08_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_I8_SS08_AddAsTitleValue", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsTitleValue("Title", "Val")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS09_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_I8_SS09_AddAsCurlyTitleWrap", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsCurlyTitleWrap("Title", "Val")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS10_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_I8_SS10_AddAsCurlyTitleWrapIf", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsCurlyTitleWrapIf(false, "T", "V")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddAsCurlyTitleWrapIf(true, "T", "V")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS11_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_I8_SS11_AddAsTitleValueIf", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddAsTitleValueIf(false, "T", "V")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddAsTitleValueIf(true, "T", "V")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS12_InsertAt(t *testing.T) {
	safeTest(t, "Test_I8_SS12_InsertAt", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		// out of range
		ss.InsertAt(-1, "x")
		if ss.Length() != 3 {
			t.Fatal("expected still 3")
		}
	})
}

func Test_I8_SS13_AddStruct(t *testing.T) {
	safeTest(t, "Test_I8_SS13_AddStruct", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		type testS struct{ Name string }
		ss.AddStruct(true, testS{Name: "hello"})
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		// nil
		ss.AddStruct(true, nil)
		if ss.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_I8_SS14_AddPointer(t *testing.T) {
	safeTest(t, "Test_I8_SS14_AddPointer", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		s := "hello"
		ss.AddPointer(true, &s)
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss.AddPointer(true, nil)
		if ss.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_I8_SS15_AddsIf(t *testing.T) {
	safeTest(t, "Test_I8_SS15_AddsIf", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddsIf(false, "a", "b")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddsIf(true, "a", "b")
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS16_AddError(t *testing.T) {
	safeTest(t, "Test_I8_SS16_AddError", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddError(errors.New("err"))
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss.AddError(nil)
		if ss.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_I8_SS17_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_I8_SS17_AsDefaultError", func() {
		ss := corestr.New.SimpleSlice.Lines("err1")
		err := ss.AsDefaultError()
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_I8_SS18_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_I8_SS18_AsError_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		err := ss.AsError(",")
		if err != nil {
			t.Fatal("expected nil")
		}
	})
}

// =============================================================================
// SimpleSlice — Access, Skip, Take
// =============================================================================

func Test_I8_SS19_First(t *testing.T) {
	safeTest(t, "Test_I8_SS19_First", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.First() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_SS20_Last(t *testing.T) {
	safeTest(t, "Test_I8_SS20_Last", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.Last() != "b" {
			t.Fatal("expected 'b'")
		}
	})
}

func Test_I8_SS21_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_SS21_FirstOrDefault", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.FirstOrDefault() != "a" {
			t.Fatal("expected 'a'")
		}
		empty := corestr.New.SimpleSlice.Cap(0)
		if empty.FirstOrDefault() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_SS22_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_SS22_LastOrDefault", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.LastOrDefault() != "a" {
			t.Fatal("expected 'a'")
		}
		empty := corestr.New.SimpleSlice.Cap(0)
		if empty.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_SS23_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS23_FirstDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.FirstDynamic() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_SS24_LastDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS24_LastDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.LastDynamic() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_SS25_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS25_FirstOrDefaultDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.FirstOrDefaultDynamic()
	})
}

func Test_I8_SS26_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS26_LastOrDefaultDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.LastOrDefaultDynamic()
	})
}

func Test_I8_SS27_Skip(t *testing.T) {
	safeTest(t, "Test_I8_SS27_Skip", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := ss.Skip(1)
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		// skip more than length
		result2 := ss.Skip(10)
		if len(result2) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_SS28_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS28_SkipDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.SkipDynamic(1)
		_ = ss.SkipDynamic(10)
	})
}

func Test_I8_SS29_Take(t *testing.T) {
	safeTest(t, "Test_I8_SS29_Take", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result := ss.Take(2)
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		// take more
		result2 := ss.Take(10)
		if len(result2) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_I8_SS30_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS30_TakeDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.TakeDynamic(1)
		_ = ss.TakeDynamic(10)
	})
}

func Test_I8_SS31_Limit(t *testing.T) {
	safeTest(t, "Test_I8_SS31_Limit", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.Limit(1)
		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS32_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_I8_SS32_LimitDynamic", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.LimitDynamic(1)
	})
}

func Test_I8_SS33_Count(t *testing.T) {
	safeTest(t, "Test_I8_SS33_Count", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.Count() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS34_CountFunc(t *testing.T) {
	safeTest(t, "Test_I8_SS34_CountFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		c := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
		if c != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS35_IsContains(t *testing.T) {
	safeTest(t, "Test_I8_SS35_IsContains", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsContains("a") {
			t.Fatal("expected true")
		}
		if ss.IsContains("z") {
			t.Fatal("expected false")
		}
	})
}

// =============================================================================
// SimpleSlice — JSON, String, Sort
// =============================================================================

func Test_I8_SS36_Json(t *testing.T) {
	safeTest(t, "Test_I8_SS36_Json", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		j := ss.Json()
		if j.JsonString() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_SS37_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_SS37_ParseInjectUsingJson", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Cap(1)
		_, err := ss2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_I8_SS38_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_SS38_ParseInjectUsingJson_Error", func() {
		ss := corestr.New.SimpleSlice.Cap(1)
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := ss.ParseInjectUsingJson(bad)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_I8_SS39_String(t *testing.T) {
	safeTest(t, "Test_I8_SS39_String", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.String() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_SS40_Join(t *testing.T) {
	safeTest(t, "Test_I8_SS40_Join", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.Join(",") != "a,b" {
			t.Fatal("expected 'a,b'")
		}
	})
}

func Test_I8_SS41_List(t *testing.T) {
	safeTest(t, "Test_I8_SS41_List", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if len(ss.List()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_SS42_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_I8_SS42_RemoveIndexes", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		result, err := ss.RemoveIndexes(1)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS43_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_SS43_MarshalJSON", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		data, err := ss.MarshalJSON()
		if err != nil || len(data) == 0 {
			t.Fatal("expected marshal")
		}
	})
}

func Test_I8_SS44_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_SS44_UnmarshalJSON", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		err := ss.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil {
			t.Fatal("unexpected error")
		}
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_SS45_SafeStrings(t *testing.T) {
	safeTest(t, "Test_I8_SS45_SafeStrings", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "", "b")
		result := ss.SafeStrings()
		if len(result) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_I8_SS46_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_SS46_Serialize", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		bytes, err := ss.Serialize()
		if err != nil || len(bytes) == 0 {
			t.Fatal("expected serialization")
		}
	})
}
