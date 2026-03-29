package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 8: Add variants, accessors, search, wrap (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovSS1_01_Add(t *testing.T) {
	safeTest(t, "Test_CovSS1_01_Add", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.Add("a")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_02_AddSplit(t *testing.T) {
	safeTest(t, "Test_CovSS1_02_AddSplit", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddSplit("a,b,c", ",")
		if ss.Length() != 3 {
			t.Fatalf("expected 3, got %d", ss.Length())
		}
	})
}

func Test_CovSS1_03_AddIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_03_AddIf", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
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

func Test_CovSS1_04_Adds_Append(t *testing.T) {
	safeTest(t, "Test_CovSS1_04_Adds_Append", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.Adds("a", "b")
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
		ss.Adds()
		ss.Append("c")
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		ss.Append()
	})
}

func Test_CovSS1_05_AppendFmt(t *testing.T) {
	safeTest(t, "Test_CovSS1_05_AppendFmt", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AppendFmt("hello %s", "world")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty format and no values
		ss.AppendFmt("")
	})
}

func Test_CovSS1_06_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_06_AppendFmtIf", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AppendFmtIf(false, "skip %s", "x")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AppendFmtIf(true, "keep %s", "x")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		// empty format
		ss.AppendFmtIf(true, "")
	})
}

func Test_CovSS1_07_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_CovSS1_07_AddAsTitleValue", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsTitleValue("Name", "Alice")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_08_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_CovSS1_08_AddAsCurlyTitleWrap", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsCurlyTitleWrap("Name", "Alice")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_09_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_09_AddAsCurlyTitleWrapIf", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsCurlyTitleWrapIf(false, "skip", "x")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddAsCurlyTitleWrapIf(true, "keep", "x")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_10_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_10_AddAsTitleValueIf", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsTitleValueIf(false, "skip", "x")
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddAsTitleValueIf(true, "keep", "x")
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_11_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovSS1_11_InsertAt", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})
		ss.InsertAt(1, "b")
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		// out of range
		ss.InsertAt(-1, "x")
		ss.InsertAt(100, "x")
	})
}

func Test_CovSS1_12_AddStruct(t *testing.T) {
	safeTest(t, "Test_CovSS1_12_AddStruct", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddStruct(true, struct{ Name string }{"Alice"})
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss.AddStruct(true, nil)
		if ss.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_CovSS1_13_AddPointer(t *testing.T) {
	safeTest(t, "Test_CovSS1_13_AddPointer", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		v := "hello"
		ss.AddPointer(false, &v)
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
		ss.AddPointer(false, nil)
	})
}

func Test_CovSS1_14_AddsIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_14_AddsIf", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
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

func Test_CovSS1_15_AddError(t *testing.T) {
	safeTest(t, "Test_CovSS1_15_AddError", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddError(nil)
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
		ss.AddError(fmt.Errorf("oops"))
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_16_AsDefaultError_AsError(t *testing.T) {
	safeTest(t, "Test_CovSS1_16_AsDefaultError_AsError", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"err1", "err2"})
		e := ss.AsDefaultError()
		if e == nil {
			t.Fatal("expected error")
		}
		e2 := ss.AsError(",")
		if e2 == nil {
			t.Fatal("expected error")
		}
		// empty
		empty := corestr.New.SimpleSlice.Strings([]string{})
		if empty.AsError(",") != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_CovSS1_17_First_Last_Dynamic(t *testing.T) {
	safeTest(t, "Test_CovSS1_17_First_Last_Dynamic", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		if ss.First() != "a" {
			t.Fatal("expected a")
		}
		if ss.Last() != "c" {
			t.Fatal("expected c")
		}
		_ = ss.FirstDynamic()
		_ = ss.LastDynamic()
	})
}

func Test_CovSS1_18_FirstOrDefault_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_CovSS1_18_FirstOrDefault_LastOrDefault", func() {
		empty := corestr.New.SimpleSlice.Strings([]string{})
		if empty.FirstOrDefault() != "" {
			t.Fatal("expected empty")
		}
		if empty.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
		_ = empty.FirstOrDefaultDynamic()
		_ = empty.LastOrDefaultDynamic()

		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.FirstOrDefault() != "a" {
			t.Fatal("expected a")
		}
		if ss.LastOrDefault() != "a" {
			t.Fatal("expected a")
		}
	})
}

func Test_CovSS1_19_Skip_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_CovSS1_19_Skip_SkipDynamic", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		r := ss.Skip(1)
		if len(r) != 2 {
			t.Fatal("expected 2")
		}
		// skip all
		r2 := ss.Skip(10)
		if len(r2) != 0 {
			t.Fatal("expected 0")
		}
		_ = ss.SkipDynamic(1)
		_ = ss.SkipDynamic(10)
	})
}

func Test_CovSS1_20_Take_TakeDynamic_Limit(t *testing.T) {
	safeTest(t, "Test_CovSS1_20_Take_TakeDynamic_Limit", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		r := ss.Take(2)
		if len(r) != 2 {
			t.Fatal("expected 2")
		}
		// take all
		r2 := ss.Take(10)
		if len(r2) != 3 {
			t.Fatal("expected 3")
		}
		_ = ss.TakeDynamic(2)
		_ = ss.TakeDynamic(10)
		_ = ss.Limit(1)
		_ = ss.LimitDynamic(1)
	})
}

func Test_CovSS1_21_Length_Count_CountFunc(t *testing.T) {
	safeTest(t, "Test_CovSS1_21_Length_Count_CountFunc", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "bb", "ccc"})
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		if ss.Count() != 3 {
			t.Fatal("expected 3")
		}
		c := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
		if c != 2 {
			t.Fatal("expected 2")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.CountFunc(func(i int, s string) bool { return true }) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSS1_22_IsEmpty_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovSS1_22_IsEmpty_HasAnyItem", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{})
		if !ss.IsEmpty() {
			t.Fatal("expected empty")
		}
		if ss.HasAnyItem() {
			t.Fatal("expected no items")
		}
		ss.Add("a")
		if ss.IsEmpty() {
			t.Fatal("expected not empty")
		}
		if !ss.HasAnyItem() {
			t.Fatal("expected items")
		}
	})
}

func Test_CovSS1_23_IsContains(t *testing.T) {
	safeTest(t, "Test_CovSS1_23_IsContains", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsContains("a") {
			t.Fatal("expected true")
		}
		if ss.IsContains("z") {
			t.Fatal("expected false")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.IsContains("a") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovSS1_24_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_CovSS1_24_IsContainsFunc", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"abc", "def"})
		found := ss.IsContainsFunc("abc", func(item, searching string) bool {
			return item == searching
		})
		if !found {
			t.Fatal("expected true")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.IsContainsFunc("x", func(a, b string) bool { return a == b }) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovSS1_25_IndexOf_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_CovSS1_25_IndexOf_IndexOfFunc", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		if ss.IndexOf("b") != 1 {
			t.Fatal("expected 1")
		}
		if ss.IndexOf("z") != -1 {
			t.Fatal("expected -1")
		}
		idx := ss.IndexOfFunc("b", func(item, searching string) bool {
			return item == searching
		})
		if idx != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.IndexOf("a") != -1 {
			t.Fatal("expected -1")
		}
		if e.IndexOfFunc("a", func(a, b string) bool { return a == b }) != -1 {
			t.Fatal("expected -1")
		}
	})
}

func Test_CovSS1_26_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovSS1_26_LastIndex_HasIndex", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if ss.LastIndex() != 1 {
			t.Fatal("expected 1")
		}
		if !ss.HasIndex(0) {
			t.Fatal("expected true")
		}
		if ss.HasIndex(5) {
			t.Fatal("expected false")
		}
		if ss.HasIndex(-1) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovSS1_27_Strings_List(t *testing.T) {
	safeTest(t, "Test_CovSS1_27_Strings_List", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if len(ss.Strings()) != 1 {
			t.Fatal("expected 1")
		}
		if len(ss.List()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovSS1_28_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_CovSS1_28_WrapQuotes", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := ss.WrapDoubleQuote()
		if (*r)[0] != `"a"` {
			t.Fatal("expected wrapped")
		}
		ss2 := corestr.New.SimpleSlice.Strings([]string{"a"})
		r2 := ss2.WrapSingleQuote()
		if (*r2)[0] != "'a'" {
			t.Fatal("expected wrapped")
		}
		ss3 := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss3.WrapTildaQuote()
		ss4 := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss4.WrapDoubleQuoteIfMissing()
		ss5 := corestr.New.SimpleSlice.Strings([]string{"a"})
		_ = ss5.WrapSingleQuoteIfMissing()
	})
}

func Test_CovSS1_29_Transpile_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_CovSS1_29_Transpile_TranspileJoin", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.Transpile(func(s string) string { return s + "!" })
		if (*r)[0] != "a!" {
			t.Fatal("expected a!")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		_ = e.Transpile(func(s string) string { return s })
		// TranspileJoin
		s := ss.TranspileJoin(func(s string) string { return s }, ",")
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovSS1_30_Hashset(t *testing.T) {
	safeTest(t, "Test_CovSS1_30_Hashset", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "a"})
		hs := ss.Hashset()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovSS1_31_Join_JoinLine_JoinSpace_JoinComma(t *testing.T) {
	safeTest(t, "Test_CovSS1_31_Join_JoinLine_JoinSpace_JoinComma", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if ss.Join(",") != "a,b" {
			t.Fatal("expected a,b")
		}
		_ = ss.JoinLine()
		_ = ss.JoinSpace()
		_ = ss.JoinComma()
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.Join(",") != "" {
			t.Fatal("expected empty")
		}
		if e.JoinLine() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovSS1_32_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_CovSS1_32_JoinLineEofLine", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.JoinLineEofLine()
		if r == "" {
			t.Fatal("expected non-empty")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.JoinLineEofLine() != "" {
			t.Fatal("expected empty")
		}
		// already has suffix
		ss2 := corestr.New.SimpleSlice.Strings([]string{"a\n"})
		_ = ss2.JoinLineEofLine()
	})
}

func Test_CovSS1_33_JoinCsv_JoinCsvLine_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_CovSS1_33_JoinCsv_JoinCsvLine_JoinCsvString", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		_ = ss.JoinCsv()
		_ = ss.JoinCsvLine()
		s := ss.JoinCsvString(",")
		if s == "" {
			t.Fatal("expected non-empty")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.JoinCsvString(",") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovSS1_34_CsvStrings(t *testing.T) {
	safeTest(t, "Test_CovSS1_34_CsvStrings", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := ss.CsvStrings()
		if len(r) != 1 {
			t.Fatal("expected 1")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if len(e.CsvStrings()) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovSS1_35_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_CovSS1_35_EachItemSplitBy", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a,b", "c,d"})
		r := ss.EachItemSplitBy(",")
		if r.Length() != 4 {
			t.Fatalf("expected 4, got %d", r.Length())
		}
	})
}

func Test_CovSS1_36_PrependJoin_AppendJoin(t *testing.T) {
	safeTest(t, "Test_CovSS1_36_PrependJoin_AppendJoin", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		s := ss.PrependJoin(",", "a")
		if s != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", s)
		}
		s2 := ss.AppendJoin(",", "c")
		if s2 != "b,c" {
			t.Fatalf("expected 'b,c', got '%s'", s2)
		}
	})
}

func Test_CovSS1_37_PrependAppend(t *testing.T) {
	safeTest(t, "Test_CovSS1_37_PrependAppend", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		ss.PrependAppend([]string{"a"}, []string{"c"})
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		// empty prepend/append
		ss.PrependAppend(nil, nil)
	})
}

func Test_CovSS1_38_JoinWith(t *testing.T) {
	safeTest(t, "Test_CovSS1_38_JoinWith", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.JoinWith(",")
		if r == "" {
			t.Fatal("expected non-empty")
		}
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if e.JoinWith(",") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovSS1_39_IsEqual(t *testing.T) {
	safeTest(t, "Test_CovSS1_39_IsEqual", func() {
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
		// nil
		if a.IsEqual(nil) {
			t.Fatal("expected false")
		}
		// diff length
		c := corestr.New.SimpleSlice.Strings([]string{"a"})
		if a.IsEqual(c) {
			t.Fatal("expected false")
		}
		// both empty
		e1 := corestr.New.SimpleSlice.Strings([]string{})
		e2 := corestr.New.SimpleSlice.Strings([]string{})
		if !e1.IsEqual(e2) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSS1_40_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_CovSS1_40_IsEqualLines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if !ss.IsEqualLines([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if ss.IsEqualLines([]string{"a", "c"}) {
			t.Fatal("expected false")
		}
		if ss.IsEqualLines([]string{"a"}) {
			t.Fatal("expected false (diff length)")
		}
	})
}

func Test_CovSS1_41_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_CovSS1_41_IsEqualUnorderedLines", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if ss.IsEqualUnorderedLines([]string{"a", "c"}) {
			t.Fatal("expected false")
		}
		// diff length
		if ss.IsEqualUnorderedLines([]string{"a"}) {
			t.Fatal("expected false")
		}
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if !e.IsEqualUnorderedLines([]string{}) {
			t.Fatal("expected true")
		}
	})
}

func Test_CovSS1_42_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_CovSS1_42_IsEqualUnorderedLinesClone", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})
		if !ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if ss.IsEqualUnorderedLinesClone([]string{"a", "c"}) {
			t.Fatal("expected false")
		}
		// diff length
		if ss.IsEqualUnorderedLinesClone([]string{"a"}) {
			t.Fatal("expected false")
		}
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		if !e.IsEqualUnorderedLinesClone([]string{}) {
			t.Fatal("expected true")
		}
	})
}
