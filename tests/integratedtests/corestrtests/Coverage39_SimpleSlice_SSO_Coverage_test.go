package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ═══ SimpleSlice comprehensive ═══

func Test_C39_SimpleSlice_Add_Adds(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Add_Adds", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Add("a")
		ss.Adds("b", "c")
		ss.Append("d")
		if ss.Length() != 4 { t.Fatal() }
		ss.Adds()
		ss.Append()
	})
}

func Test_C39_SimpleSlice_AddIf_AddsIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddIf_AddsIf", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddIf(false, "skip")
		ss.AddIf(true, "keep")
		ss.AddsIf(false, "x")
		ss.AddsIf(true, "y")
		if ss.Length() != 2 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddSplit", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddSplit("a,b,c", ",")
		if ss.Length() != 3 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddError", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddError(nil)
		ss.AddError(errForTest)
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddStruct_AddPointer(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddStruct_AddPointer", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddStruct(false, "hello")
		ss.AddStruct(true, nil)
		ss.AddPointer(false, "world")
		ss.AddPointer(true, nil)
		if ss.Length() != 2 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AppendFmt", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AppendFmt("hello %d", 42)
		ss.AppendFmt("", )
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AppendFmtIf", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AppendFmtIf(true, "x=%d", 1)
		ss.AppendFmtIf(false, "y=%d", 2)
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsTitleValue", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsTitleValue("Name", "Alice")
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsTitleValueIf", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsTitleValueIf(true, "N", "A")
		ss.AddAsTitleValueIf(false, "N", "B")
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsCurlyTitleWrap", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsCurlyTitleWrap("Key", "Val")
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsCurlyTitleWrapIf", func() {
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsCurlyTitleWrapIf(true, "K", "V")
		ss.AddAsCurlyTitleWrapIf(false, "K", "V")
		if ss.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_InsertAt", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")
		if ss.Length() != 3 { t.Fatal() }
		ss.InsertAt(-1, "x") // no-op
		ss.InsertAt(100, "y") // no-op
	})
}

func Test_C39_SimpleSlice_Accessors(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Accessors", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		if ss.First() != "a" { t.Fatal() }
		if ss.Last() != "c" { t.Fatal() }
		if ss.FirstOrDefault() != "a" { t.Fatal() }
		if ss.LastOrDefault() != "c" { t.Fatal() }
		if ss.FirstDynamic() != "a" { t.Fatal() }
		if ss.LastDynamic() != "c" { t.Fatal() }
		if ss.FirstOrDefaultDynamic() != "a" { t.Fatal() }
		if ss.LastOrDefaultDynamic() != "c" { t.Fatal() }
		if ss.Length() != 3 { t.Fatal() }
		if ss.Count() != 3 { t.Fatal() }
		if ss.LastIndex() != 2 { t.Fatal() }
		if !ss.HasIndex(0) { t.Fatal() }
		if ss.HasIndex(5) { t.Fatal() }
		if !ss.HasAnyItem() { t.Fatal() }
		// empty
		es := corestr.Empty.SimpleSlice()
		if es.FirstOrDefault() != "" { t.Fatal() }
		if es.LastOrDefault() != "" { t.Fatal() }
		var nilSS *corestr.SimpleSlice
		if nilSS.Length() != 0 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_Skip_Take_Limit(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Skip_Take_Limit", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		s := ss.Skip(1)
		if len(s) != 2 { t.Fatal() }
		s2 := ss.Skip(10)
		if len(s2) != 0 { t.Fatal() }
		tk := ss.Take(2)
		if len(tk) != 2 { t.Fatal() }
		tk2 := ss.Take(10)
		if len(tk2) != 3 { t.Fatal() }
		lm := ss.Limit(1)
		if len(lm) != 1 { t.Fatal() }
		// Dynamic
		_ = ss.SkipDynamic(1)
		_ = ss.SkipDynamic(100)
		_ = ss.TakeDynamic(1)
		_ = ss.TakeDynamic(100)
		_ = ss.LimitDynamic(1)
	})
}

func Test_C39_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsContains", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsContains("a") { t.Fatal() }
		if ss.IsContains("z") { t.Fatal() }
		if corestr.Empty.SimpleSlice().IsContains("a") { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsContainsFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("abc", "xyz")
		if !ss.IsContainsFunc("abc", func(a, b string) bool { return a == b }) { t.Fatal() }
		if corestr.Empty.SimpleSlice().IsContainsFunc("a", func(a, b string) bool { return true }) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IndexOf", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.IndexOf("b") != 1 { t.Fatal() }
		if ss.IndexOf("z") != -1 { t.Fatal() }
		if corestr.Empty.SimpleSlice().IndexOf("a") != -1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IndexOfFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("abc")
		if ss.IndexOfFunc("abc", func(a, b string) bool { return a == b }) != 0 { t.Fatal() }
		if corestr.Empty.SimpleSlice().IndexOfFunc("a", func(a, b string) bool { return true }) != -1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_CountFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
		if count != 2 { t.Fatal() }
		if corestr.Empty.SimpleSlice().CountFunc(func(i int, s string) bool { return true }) != 0 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_WrapQuotes", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.WrapDoubleQuote()
		ss2 := corestr.New.SimpleSlice.Lines("a")
		_ = ss2.WrapSingleQuote()
		ss3 := corestr.New.SimpleSlice.Lines("a")
		_ = ss3.WrapTildaQuote()
		ss4 := corestr.New.SimpleSlice.Lines("a")
		_ = ss4.WrapDoubleQuoteIfMissing()
		ss5 := corestr.New.SimpleSlice.Lines("a")
		_ = ss5.WrapSingleQuoteIfMissing()
	})
}

func Test_C39_SimpleSlice_Transpile_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Transpile_TranspileJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		ts := ss.Transpile(func(s string) string { return s + "!" })
		if ts.Length() != 2 { t.Fatal() }
		tj := ss.TranspileJoin(func(s string) string { return s }, ",")
		if tj == "" { t.Fatal() }
		// empty transpile
		_ = corestr.Empty.SimpleSlice().Transpile(func(s string) string { return s })
	})
}

func Test_C39_SimpleSlice_Join_Methods(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Join_Methods", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.Join(",") == "" { t.Fatal() }
		if ss.JoinLine() == "" { t.Fatal() }
		if ss.JoinSpace() == "" { t.Fatal() }
		if ss.JoinComma() == "" { t.Fatal() }
		if ss.JoinCsv() == "" { t.Fatal() }
		if ss.JoinCsvLine() == "" { t.Fatal() }
		if ss.JoinWith(",") == "" { t.Fatal() }
		if ss.JoinCsvString(",") == "" { t.Fatal() }
		if corestr.Empty.SimpleSlice().Join(",") != "" { t.Fatal() }
		if corestr.Empty.SimpleSlice().JoinLine() != "" { t.Fatal() }
		if corestr.Empty.SimpleSlice().JoinWith(",") != "" { t.Fatal() }
		if corestr.Empty.SimpleSlice().JoinCsvString(",") != "" { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_JoinLineEofLine", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		r := ss.JoinLineEofLine()
		if r == "" { t.Fatal() }
		// already has suffix
		ss2 := corestr.New.SimpleSlice.Lines("a\n")
		_ = ss2.JoinLineEofLine()
		if corestr.Empty.SimpleSlice().JoinLineEofLine() != "" { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_PrependJoin_AppendJoin(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_PrependJoin_AppendJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("b")
		r := ss.PrependJoin(",", "a")
		if r == "" { t.Fatal() }
		r2 := ss.AppendJoin(",", "c")
		if r2 == "" { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_PrependAppend", func() {
		ss := corestr.New.SimpleSlice.Lines("b")
		ss.PrependAppend([]string{"a"}, []string{"c"})
		if ss.Length() != 3 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_EachItemSplitBy", func() {
		ss := corestr.New.SimpleSlice.Lines("a,b", "c")
		split := ss.EachItemSplitBy(",")
		if split.Length() != 3 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqual", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		if !a.IsEqual(b) { t.Fatal() }
		var nilSS *corestr.SimpleSlice
		if !nilSS.IsEqual(nil) { t.Fatal() }
		if a.IsEqual(nil) { t.Fatal() }
		c := corestr.New.SimpleSlice.Lines("a")
		if a.IsEqual(c) { t.Fatal() }
		// empty
		e1 := corestr.New.SimpleSlice.Empty()
		e2 := corestr.New.SimpleSlice.Empty()
		if !e1.IsEqual(e2) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualLines", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsEqualLines([]string{"a", "b"}) { t.Fatal() }
		if ss.IsEqualLines([]string{"a"}) { t.Fatal() }
		if ss.IsEqualLines([]string{"a", "c"}) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualUnorderedLines", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		if !ss.IsEqualUnorderedLines([]string{"a", "b"}) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		if !ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsUnorderedEqualRaw(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsUnorderedEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		if !ss.IsUnorderedEqualRaw(true, "a", "b") { t.Fatal() }
		if !ss.IsUnorderedEqualRaw(false, "a", "b") { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsUnorderedEqual", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		other := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsUnorderedEqual(true, other) { t.Fatal() }
		e1 := corestr.New.SimpleSlice.Empty()
		e2 := corestr.New.SimpleSlice.Empty()
		if !e1.IsUnorderedEqual(true, e2) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsDistinctEqual", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsDistinctEqualRaw("a", "b") { t.Fatal() }
		other := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsDistinctEqual(other) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualByFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b") { t.Fatal() }
		if ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a") { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsEqualByFuncLinesSplit(false, "\n", "a\nb", func(i int, l, r string) bool { return l == r }) { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Sort_Reverse", func() {
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")
		ss.Sort()
		if ss.First() != "a" { t.Fatal() }
		ss.Reverse()
		if ss.First() != "c" { t.Fatal() }
		// 2 items
		ss2 := corestr.New.SimpleSlice.Lines("b", "a")
		ss2.Reverse()
		if ss2.First() != "a" { t.Fatal() }
		// 1 item
		ss3 := corestr.New.SimpleSlice.Lines("a")
		ss3.Reverse()
	})
}

func Test_C39_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_ConcatNew", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		cn := ss.ConcatNew("b", "c")
		if cn.Length() != 3 { t.Fatal() }
		cns := ss.ConcatNewStrings("d")
		if len(cns) != 2 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_ConcatNewSimpleSlices", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss2 := corestr.New.SimpleSlice.Lines("b")
		cn := ss.ConcatNewSimpleSlices(ss2)
		if cn.Length() != 2 { t.Fatal() }
	})
}
func Test_C39_SimpleSlice_Collection_Hashset(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Collection_Hashset", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.Collection(false).Length() != 1 { t.Fatal() }
		if ss.ToCollection(false).Length() != 1 { t.Fatal() }
		if ss.Hashset().Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Strings_List", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if len(ss.Strings()) != 1 { t.Fatal() }
		if len(ss.List()) != 1 { t.Fatal() }
		if len(ss.SafeStrings()) != 1 { t.Fatal() }
		if len(corestr.Empty.SimpleSlice().SafeStrings()) != 0 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_String", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.String() == "" { t.Fatal() }
		if corestr.Empty.SimpleSlice().String() != "" { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_CsvStrings", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if len(ss.CsvStrings()) != 1 { t.Fatal() }
		if len(corestr.Empty.SimpleSlice().CsvStrings()) != 0 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AsError(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AsError", func() {
		ss := corestr.New.SimpleSlice.Lines("err")
		if ss.AsError(",") == nil { t.Fatal() }
		if ss.AsDefaultError() == nil { t.Fatal() }
		if corestr.Empty.SimpleSlice().AsError(",") != nil { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_RemoveIndexes", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		newSS, err := ss.RemoveIndexes(1)
		if err != nil { t.Fatal(err) }
		if newSS.Length() != 2 { t.Fatal() }
		_, err2 := corestr.Empty.SimpleSlice().RemoveIndexes(0)
		if err2 == nil { t.Fatal("expected error") }
	})
}

func Test_C39_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_DistinctDiff", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		r := ss.DistinctDiffRaw("b", "c")
		if len(r) < 1 { t.Fatal() }
		other := corestr.New.SimpleSlice.Lines("b", "c")
		r2 := ss.DistinctDiff(other)
		if len(r2) < 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddedRemovedLinesDiff", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := ss.AddedRemovedLinesDiff("b", "c")
		if len(added) == 0 { t.Fatal("expected added") }
		if len(removed) == 0 { t.Fatal("expected removed") }
	})
}

func Test_C39_SimpleSlice_NonPtr_Ptr_ToPtr_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_NonPtr_Ptr_ToPtr_ToNonPtr", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.NonPtr()
		_ = ss.Ptr()
		_ = ss.ToPtr()
		_ = ss.ToNonPtr()
	})
}

func Test_C39_SimpleSlice_JSON(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_JSON", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		j := ss.Json()
		if j.HasError() { t.Fatal(j.Error) }
		jp := ss.JsonPtr()
		if jp.HasError() { t.Fatal(jp.Error) }
		if ss.JsonModelAny() == nil { t.Fatal() }
		b, err := ss.MarshalJSON()
		if err != nil { t.Fatal(err) }
		ss2 := corestr.New.SimpleSlice.Empty()
		err2 := ss2.UnmarshalJSON(b)
		if err2 != nil { t.Fatal(err2) }
		err3 := ss2.UnmarshalJSON([]byte(`{bad`))
		if err3 == nil { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_ParseInjectUsingJson", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Empty()
		result, err := ss2.ParseInjectUsingJson(jr)
		if err != nil { t.Fatal(err) }
		if result.Length() != 1 { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_InterfaceCasts", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.AsJsonContractsBinder() == nil { t.Fatal() }
		if ss.AsJsoner() == nil { t.Fatal() }
		if ss.AsJsonParseSelfInjector() == nil { t.Fatal() }
		if ss.AsJsonMarshaller() == nil { t.Fatal() }
	})
}

func Test_C39_SimpleSlice_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Clear_Dispose", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Clear()
		if ss.Length() != 0 { t.Fatal() }
		ss2 := corestr.New.SimpleSlice.Lines("a")
		ss2.Dispose()
		var nilSS *corestr.SimpleSlice
		if nilSS.Clear() != nil { t.Fatal() }
		nilSS.Dispose()
	})
}

func Test_C39_SimpleSlice_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Serialize_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		b, err := ss.Serialize()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal() }
		var target []string
		err2 := ss.Deserialize(&target)
		if err2 != nil { t.Fatal(err2) }
	})
}

	// ── newSimpleSliceCreator ──
func Test_C39_SSO_SetGet(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetGet", func() {
		sso := corestr.New.SimpleStringOnce.Create("hello", true)
		if sso.Value() != "hello" { t.Fatal() }
		if !sso.IsInitialized() { t.Fatal() }
		if !sso.IsDefined() { t.Fatal() }
		if sso.IsUninitialized() { t.Fatal() }
		if sso.IsInvalid() { t.Fatal() }
	})
}

func Test_C39_SSO_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetOnUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		err := sso.SetOnUninitialized("val")
		if err != nil { t.Fatal(err) }
		err2 := sso.SetOnUninitialized("val2")
		if err2 == nil { t.Fatal("expected error") }
	})
}

func Test_C39_SSO_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_C39_SSO_GetSetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetSetOnce("first")
		if v != "first" { t.Fatal() }
		v2 := sso.GetSetOnce("second")
		if v2 != "first" { t.Fatal() }
	})
}

func Test_C39_SSO_GetOnce(t *testing.T) {
	safeTest(t, "Test_C39_SSO_GetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetOnce()
		if v != "" { t.Fatal() }
	})
}

func Test_C39_SSO_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C39_SSO_GetOnceFunc", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetOnceFunc(func() string { return "computed" })
		if v != "computed" { t.Fatal() }
		v2 := sso.GetOnceFunc(func() string { return "other" })
		if v2 != "computed" { t.Fatal() }
	})
}

func Test_C39_SSO_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetOnceIfUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		if !sso.SetOnceIfUninitialized("v") { t.Fatal() }
		if sso.SetOnceIfUninitialized("v2") { t.Fatal() }
	})
}

func Test_C39_SSO_Invalidate_Reset(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Invalidate_Reset", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		sso.Invalidate()
		if sso.IsInitialized() { t.Fatal() }
		sso2 := corestr.New.SimpleStringOnce.Init("world")
		sso2.Reset()
		if sso2.IsInitialized() { t.Fatal() }
	})
}

func Test_C39_SSO_NumericConversions(t *testing.T) {
	safeTest(t, "Test_C39_SSO_NumericConversions", func() {
		sso := corestr.New.SimpleStringOnce.Init("42")
		if sso.Int() != 42 { t.Fatal() }
		if sso.ValueDefInt() != 42 { t.Fatal() }
		if sso.ValueInt(0) != 42 { t.Fatal() }
		if sso.Byte() != 42 { t.Fatal() }
		if sso.ValueDefByte() != 42 { t.Fatal() }
		if sso.ValueByte(0) != 42 { t.Fatal() }
		if sso.Int16() != 42 { t.Fatal() }
		if sso.Int32() != 42 { t.Fatal() }
		v, ok := sso.Uint16()
		if !ok || v != 42 { t.Fatal() }
		v2, ok2 := sso.Uint32()
		if !ok2 || v2 != 42 { t.Fatal() }
		// non-numeric
		bad := corestr.New.SimpleStringOnce.Init("abc")
		if bad.Int() != 0 { t.Fatal() }
		if bad.Byte() != 0 { t.Fatal() }
		if bad.Int16() != 0 { t.Fatal() }
		if bad.Int32() != 0 { t.Fatal() }
	})
}

func Test_C39_SSO_Float64(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Float64", func() {
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		if sso.ValueFloat64(0) == 0 { t.Fatal() }
		if sso.ValueDefFloat64() == 0 { t.Fatal() }
		bad := corestr.New.SimpleStringOnce.Init("abc")
		if bad.ValueFloat64(1.0) != 1.0 { t.Fatal() }
	})
}

func Test_C39_SSO_Boolean(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Boolean", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		if !sso.Boolean(false) { t.Fatal() }
		if !sso.BooleanDefault() { t.Fatal() }
		if !sso.IsValueBool() { t.Fatal() }
		sso2 := corestr.New.SimpleStringOnce.Init("true")
		if !sso2.Boolean(false) { t.Fatal() }
		bad := corestr.New.SimpleStringOnce.Init("xyz")
		if bad.Boolean(false) { t.Fatal() }
		uninit := corestr.New.SimpleStringOnce.Empty()
		if uninit.Boolean(true) { t.Fatal() }
	})
}

func Test_C39_SSO_IsSetter(t *testing.T) {
	safeTest(t, "Test_C39_SSO_IsSetter", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		if !sso.IsSetter(false).IsTrue() { t.Fatal() }
		uninit := corestr.New.SimpleStringOnce.Empty()
		if uninit.IsSetter(true).IsTrue() { t.Fatal() }
		bad := corestr.New.SimpleStringOnce.Init("xyz")
		_ = bad.IsSetter(false)
		ssoTrue := corestr.New.SimpleStringOnce.Init("true")
		if !ssoTrue.IsSetter(false).IsTrue() { t.Fatal() }
	})
}

func Test_C39_SSO_WithinRange(t *testing.T) {
	safeTest(t, "Test_C39_SSO_WithinRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("50")
		v, ok := sso.WithinRange(true, 0, 100)
		if !ok || v != 50 { t.Fatal() }
		v2, ok2 := sso.WithinRange(true, 60, 100)
		if ok2 { t.Fatal() }
		if v2 != 60 { t.Fatal() } // boundary
		v3, ok3 := sso.WithinRange(true, 0, 10)
		if ok3 { t.Fatal() }
		if v3 != 10 { t.Fatal() } // boundary
		v4, ok4 := sso.WithinRange(false, 60, 100)
		if ok4 || v4 != 50 { t.Fatal() }
		_, ok5 := sso.WithinRangeDefault(0, 100)
		if !ok5 { t.Fatal() }
	})
}

func Test_C39_SSO_StringMethods(t *testing.T) {
	safeTest(t, "Test_C39_SSO_StringMethods", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		if sso.String() == "" { t.Fatal() }
		if sso.StringPtr() == nil { t.Fatal() }
		if sso.SafeValue() == "" { t.Fatal() }
		if sso.Trim() != "hello" { t.Fatal() }
		if sso.IsEmpty() { t.Fatal() }
		if sso.IsWhitespace() { t.Fatal() }
		if !sso.HasValidNonEmpty() { t.Fatal() }
		if !sso.HasValidNonWhitespace() { t.Fatal() }
		if !sso.HasSafeNonEmpty() { t.Fatal() }
		if !sso.Is("hello") { t.Fatal() }
		if !sso.IsAnyOf("hello", "world") { t.Fatal() }
		if !sso.IsContains("ell") { t.Fatal() }
		if !sso.IsAnyContains("xyz", "ell") { t.Fatal() }
		if sso.IsAnyContains("xyz") { t.Fatal() }
		if !sso.IsEqualNonSensitive("HELLO") { t.Fatal() }
		// nil ptr
		var nilSSO *corestr.SimpleStringOnce
		if nilSSO.String() != "" { t.Fatal() }
		if nilSSO.StringPtr() == nil { t.Fatal() }
		// uninit safe value
		uninit := corestr.New.SimpleStringOnce.Empty()
		if uninit.SafeValue() != "" { t.Fatal() }
		// IsAnyOf empty
		if !sso.IsAnyOf() { t.Fatal() }
		if !sso.IsAnyContains() { t.Fatal() }
	})
}

func Test_C39_SSO_ValueBytes(t *testing.T) {
	safeTest(t, "Test_C39_SSO_ValueBytes", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		if len(sso.ValueBytes()) != 2 { t.Fatal() }
		if len(sso.ValueBytesPtr()) != 2 { t.Fatal() }
	})
}

func Test_C39_SSO_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C39_SSO_ConcatNew", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		cn := sso.ConcatNew(" world")
		if cn.Value() != "hello world" { t.Fatal() }
	})
}

func Test_C39_SSO_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C39_SSO_ConcatNewUsingStrings", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		cn := sso.ConcatNewUsingStrings(",", "world")
		if cn.Value() == "" { t.Fatal() }
	})
}
func Test_C39_SSO_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_C39_SSO_NonPtr_Ptr", func() {
		sso := corestr.New.SimpleStringOnce.Init("h")
		_ = sso.NonPtr()
		_ = sso.Ptr()
	})
}

func Test_C39_SSO_SetInit_SetUnInit(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetInit_SetUnInit", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		sso.SetInitialize()
		if !sso.IsInitialized() { t.Fatal() }
		sso.SetUnInit()
		if sso.IsInitialized() { t.Fatal() }
	})
}

func Test_C39_SSO_Split(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Split", func() {
		sso := corestr.New.SimpleStringOnce.Init("a,b,c")
		if len(sso.Split(",")) != 3 { t.Fatal() }
		l, r := sso.SplitLeftRight(",")
		if l != "a" || r == "" { t.Fatal() }
		l2, r2 := sso.SplitLeftRightTrim(",")
		if l2 != "a" || r2 == "" { t.Fatal() }
		_ = sso.SplitNonEmpty(",")
		_ = sso.SplitTrimNonWhitespace(",")
		_ = sso.LinesSimpleSlice()
		_ = sso.SimpleSlice(",")
	})
}

func Test_C39_SSO_SplitLeftRight_SingleItem(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SplitLeftRight_SingleItem", func() {
		sso := corestr.New.SimpleStringOnce.Init("nosep")
		l, r := sso.SplitLeftRight(",")
		if l != "nosep" || r != "" { t.Fatal() }
		l2, r2 := sso.SplitLeftRightTrim(",")
		if l2 != "nosep" || r2 != "" { t.Fatal() }
	})
}

func Test_C39_SSO_Regex(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Regex", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello123")
		if sso.IsRegexMatches(nil) { t.Fatal() }
		if sso.RegexFindString(nil) != "" { t.Fatal() }
		r, ok := sso.RegexFindAllStringsWithFlag(nil, -1)
		if ok || len(r) != 0 { t.Fatal() }
		r2 := sso.RegexFindAllStrings(nil, -1)
		if len(r2) != 0 { t.Fatal() }
	})
}

func Test_C39_SSO_Dispose(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Dispose", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		sso.Dispose()
		var nilSSO *corestr.SimpleStringOnce
		nilSSO.Dispose()
	})
}

func Test_C39_SSO_JSON(t *testing.T) {
	safeTest(t, "Test_C39_SSO_JSON", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		j := sso.Json()
		if j.HasError() { t.Fatal(j.Error) }
		jp := sso.JsonPtr()
		if jp.HasError() { t.Fatal(jp.Error) }
		if sso.JsonModelAny() == nil { t.Fatal() }
		b, err := sso.MarshalJSON()
		if err != nil { t.Fatal(err) }
		sso2 := corestr.New.SimpleStringOnce.Empty()
		err2 := sso2.UnmarshalJSON(b)
		if err2 != nil { t.Fatal(err2) }
	})
}

func Test_C39_SSO_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C39_SSO_InterfaceCasts", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		if sso.AsJsonContractsBinder() == nil { t.Fatal() }
		if sso.AsJsoner() == nil { t.Fatal() }
		if sso.AsJsonParseSelfInjector() == nil { t.Fatal() }
		if sso.AsJsonMarshaller() == nil { t.Fatal() }
	})
}

func Test_C39_SSO_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Serialize_Deserialize", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		b, err := sso.Serialize()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal() }
	})
}

	// ── newSimpleStringOnceCreator ──

func Test_C39_NewSSOCreator(t *testing.T) {
	safeTest(t, "Test_C39_NewSSOCreator", func() {
		s1 := corestr.New.SimpleStringOnce.Init("hi")
		if s1.Value() != "hi" { t.Fatal() }
		s2 := corestr.New.SimpleStringOnce.InitPtr("hi")
		if s2.Value() != "hi" { t.Fatal() }
		s3 := corestr.New.SimpleStringOnce.Create("v", true)
		if s3.Value() != "v" { t.Fatal() }
		s4 := corestr.New.SimpleStringOnce.CreatePtr("v", true)
		if s4.Value() != "v" { t.Fatal() }
		s5 := corestr.New.SimpleStringOnce.Uninitialized("val")
		if s5.IsInitialized() { t.Fatal() }
		s6 := corestr.New.SimpleStringOnce.Empty()
		if s6.IsInitialized() { t.Fatal() }
		s7 := corestr.New.SimpleStringOnce.Any(false, 42, true)
		if s7.Value() == "" { t.Fatal() }
	})
}
