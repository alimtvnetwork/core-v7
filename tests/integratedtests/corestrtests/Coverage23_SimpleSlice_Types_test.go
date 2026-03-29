package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_SimpleSlice_BasicOps(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_BasicOps", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		if ss.Length() != 3 || ss.IsEmpty() || !ss.HasAnyItem() || ss.LastIndex() != 2 {
			t.Fatal("basic checks failed")
		}
		if ss.Count() != 3 {
			t.Fatal("count wrong")
		}
	})
}

func Test_Cov23_SimpleSlice_AddMethods(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_AddMethods", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.Add("a")
		ss.AddIf(false, "skip")
		ss.AddIf(true, "b")
		ss.Adds("c", "d")
		ss.Append("e")
		ss.AppendFmt("%s-%d", "f", 1)
		ss.AppendFmt("literal-no-fmt") // no format directives, single arg
		ss.AppendFmtIf(false, "%s", "skip")
		ss.AppendFmtIf(true, "%s", "g")
		ss.AddAsTitleValue("title", "value")
		ss.AddAsCurlyTitleWrap("title", "value")
		ss.AddAsCurlyTitleWrapIf(false, "t", "v")
		ss.AddAsCurlyTitleWrapIf(true, "t", "v")
		ss.AddAsTitleValueIf(false, "t", "v")
		ss.AddAsTitleValueIf(true, "t", "v")
		ss.AddsIf(false, "x")
		ss.AddsIf(true, "h")
		ss.AddError(nil)
		ss.AddError(errors.New("e"))
		ss.AddSplit("x,y", ",")
		ss.AddStruct(false, map[string]int{"a": 1})
		ss.AddStruct(false, nil)
		ss.AddPointer(false, nil)
		ss.AddPointer(false, "hello")
	})
}

func Test_Cov23_SimpleSlice_FirstLast(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_FirstLast", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.First() != "a" || ss.Last() != "b" {
			t.Fatal("wrong")
		}
		if ss.FirstOrDefault() != "a" || ss.LastOrDefault() != "b" {
			t.Fatal("wrong")
		}
		_ = ss.FirstDynamic()
		_ = ss.LastDynamic()
		_ = ss.FirstOrDefaultDynamic()
		_ = ss.LastOrDefaultDynamic()
		e := corestr.New.SimpleSlice.Empty()
		if e.FirstOrDefault() != "" || e.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_Cov23_SimpleSlice_SkipTakeLimit(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_SkipTakeLimit", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		if len(ss.Skip(1)) != 2 {
			t.Fatal("wrong skip")
		}
		if len(ss.Take(2)) != 2 {
			t.Fatal("wrong take")
		}
		if len(ss.Limit(2)) != 2 {
			t.Fatal("wrong limit")
		}
		_ = ss.SkipDynamic(1)
		_ = ss.TakeDynamic(2)
		_ = ss.LimitDynamic(2)
	})
}

func Test_Cov23_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsContains", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsContains("a") {
			t.Fatal("expected true")
		}
		if ss.IsContains("x") {
			t.Fatal("expected false")
		}
	})
}

func Test_Cov23_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsContainsFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("abc", "def")
		if !ss.IsContainsFunc("abc", func(a, b string) bool { return a == b }) {
			t.Fatal("expected true")
		}
	})
}

func Test_Cov23_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IndexOf", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if ss.IndexOf("b") != 1 {
			t.Fatal("wrong index")
		}
		if ss.IndexOf("x") != -1 {
			t.Fatal("expected -1")
		}
	})
}

func Test_Cov23_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IndexOfFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		idx := ss.IndexOfFunc("b", func(a, b string) bool { return a == b })
		if idx != 1 {
			t.Fatal("wrong index")
		}
	})
}

func Test_Cov23_SimpleSlice_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_HasIndex", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if !ss.HasIndex(0) {
			t.Fatal("expected true")
		}
		if ss.HasIndex(5) {
			t.Fatal("expected false")
		}
	})
}

func Test_Cov23_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_InsertAt", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
		ss.InsertAt(-1, "x") // out of range, skip
		ss.InsertAt(100, "y") // out of range, skip
	})
}

func Test_Cov23_SimpleSlice_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_WrapQuotes", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.WrapDoubleQuote()
		ss2 := corestr.New.SimpleSlice.Lines("b")
		_ = ss2.WrapSingleQuote()
		ss3 := corestr.New.SimpleSlice.Lines("c")
		_ = ss3.WrapTildaQuote()
		ss4 := corestr.New.SimpleSlice.Lines("d")
		_ = ss4.WrapDoubleQuoteIfMissing()
		ss5 := corestr.New.SimpleSlice.Lines("e")
		_ = ss5.WrapSingleQuoteIfMissing()
	})
}

func Test_Cov23_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Join", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.Join(",")
		_ = ss.JoinLine()
		_ = ss.JoinSpace()
		_ = ss.JoinComma()
		_ = ss.JoinCsv()
		_ = ss.JoinCsvLine()
		_ = ss.JoinWith(",")
		_ = ss.JoinLineEofLine()
		_ = ss.JoinCsvString(",")
	})
}

func Test_Cov23_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Strings_List", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.Strings()
		_ = ss.List()
	})
}

func Test_Cov23_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Transpile", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.Transpile(func(s string) string { return s + "!" })
		if !result.IsContains("a!") {
			t.Fatal("expected transpiled")
		}
	})
}

func Test_Cov23_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_TranspileJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		s := ss.TranspileJoin(func(s string) string { return s + "!" }, ",")
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov23_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_EachItemSplitBy", func() {
		ss := corestr.New.SimpleSlice.Lines("a,b", "c")
		result := ss.EachItemSplitBy(",")
		if result.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov23_SimpleSlice_Concat(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Concat", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		newSS := ss.ConcatNew("b", "c")
		if newSS.Length() != 3 {
			t.Fatal("expected 3")
		}
		_ = ss.ConcatNewStrings("d")
		ss2 := corestr.New.SimpleSlice.Lines("e")
		_ = ss.ConcatNewSimpleSlices(ss2)
	})
}

func Test_Cov23_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_PrependAppend", func() {
		ss := corestr.New.SimpleSlice.Lines("b")
		ss.PrependAppend([]string{"a"}, []string{"c"})
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_Cov23_SimpleSlice_PrependAppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_PrependAppendJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("b")
		s := ss.PrependJoin(",", "a")
		if s == "" {
			t.Fatal("expected non-empty")
		}
		s2 := ss.AppendJoin(",", "c")
		if s2 == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov23_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqual", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		if !a.IsEqual(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov23_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqualLines", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		if !ss.IsEqualLines([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov23_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqualUnorderedLines", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		if !ss.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov23_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		if !ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Fatal("expected equal")
		}
	})
}

func Test_Cov23_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Collection", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		c := ss.Collection(false)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
		_ = ss.ToCollection(true)
	})
}

func Test_Cov23_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_CountFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		cnt := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
		if cnt != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov23_SimpleSlice_AsError(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_AsError", func() {
		ss := corestr.New.SimpleSlice.Lines("e1", "e2")
		err := ss.AsError(",")
		if err == nil {
			t.Fatal("expected error")
		}
		_ = ss.AsDefaultError()
		empty := corestr.New.SimpleSlice.Empty()
		if empty.AsError(",") != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_Cov23_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_String", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		if ss.String() == "" {
			t.Fatal("expected non-empty")
		}
		empty := corestr.New.SimpleSlice.Empty()
		if empty.String() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_Cov23_SimpleSlice_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Sort_Reverse", func() {
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")
		ss.Sort()
		if ss.First() != "a" {
			t.Fatal("wrong sort")
		}
		ss.Reverse()
		if ss.First() != "c" {
			t.Fatal("wrong reverse")
		}
	})
}

func Test_Cov23_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Hashset", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		hs := ss.Hashset()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov23_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_CsvStrings", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		csv := ss.CsvStrings()
		if len(csv) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov23_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_JsonModel", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.JsonModel()
		_ = ss.NonPtr()
		_ = ss.Ptr()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_CollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection", func() {
		cc := corestr.New.CollectionsOfCollection.Empty()
		if !cc.IsEmpty() || cc.HasItems() || cc.Length() != 0 {
			t.Fatal("basic checks failed")
		}
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		cc.Add(c1).Add(c2)
		if cc.Length() != 2 || cc.IsEmpty() || !cc.HasItems() {
			t.Fatal("filled checks failed")
		}
		if cc.AllIndividualItemsLength() != 3 {
			t.Fatal("expected 3")
		}
		list := cc.List(0)
		if len(list) != 3 {
			t.Fatal("expected 3")
		}
		col := cc.ToCollection()
		if col.Length() != 3 {
			t.Fatal("expected 3")
		}
		_ = cc.Items()
		_ = cc.String()
	})
}

func Test_Cov23_CollectionsOfCollection_AddStrings(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_AddStrings", func() {
		cc := corestr.New.CollectionsOfCollection.Empty()
		cc.AddStrings(true, []string{"a", "b"})
		if cc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_Cov23_CollectionsOfCollection_AddsStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_AddsStringsOfStrings", func() {
		cc := corestr.New.CollectionsOfCollection.Empty()
		cc.AddsStringsOfStrings(true, []string{"a"}, []string{"b"})
		if cc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov23_CollectionsOfCollection_AddCollections(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_AddCollections", func() {
		cc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})
		cc.AddCollections(c)
		cc.Adds(c)
		if cc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_Cov23_CollectionsOfCollection_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_JsonOps", func() {
		cc := corestr.New.CollectionsOfCollection.Empty()
		cc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = cc.JsonModel()
		_ = cc.JsonModelAny()
		b, err := cc.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected")
		}
		cc2 := corestr.New.CollectionsOfCollection.Empty()
		_ = cc2.UnmarshalJSON(b)
		_ = cc.Json()
		_ = cc.JsonPtr()
		_ = cc.AsJsonContractsBinder()
		_ = cc.AsJsoner()
		_ = cc.AsJsonParseSelfInjector()
		_ = cc.AsJsonMarshaller()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashmapDiff — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_HashmapDiff(t *testing.T) {
	safeTest(t, "Test_Cov23_HashmapDiff", func() {
		hd := corestr.HashmapDiff{"a": "1", "b": "2"}
		if hd.IsEmpty() || !hd.HasAnyItem() || hd.Length() != 2 || hd.LastIndex() != 1 {
			t.Fatal("basic checks failed")
		}
		_ = hd.Raw()
		_ = hd.AllKeysSorted()
		_ = hd.MapAnyItems()
		_ = hd.RawMapStringAnyDiff()
		_ = hd.IsRawEqual(map[string]string{"a": "1", "b": "2"})
		_ = hd.HasAnyChanges(map[string]string{"a": "1"})
		_ = hd.HashmapDiffUsingRaw(map[string]string{"a": "1"})
		_ = hd.DiffRaw(map[string]string{"a": "1"})
		_ = hd.DiffJsonMessage(map[string]string{"a": "1"})
		_ = hd.ShouldDiffMessage("title", map[string]string{"a": "1"})
		_ = hd.LogShouldDiffMessage("title", map[string]string{"a": "1"})
		diff := hd.DiffRaw(map[string]string{"a": "1"})
		_ = hd.ToStringsSliceOfDiffMap(diff)
		_, _ = hd.Serialize()
		var target map[string]string
		_ = hd.Deserialize(&target)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue / ValidValues / ValueStatus / KeyValuePair / KeyAnyValuePair / TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov23_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		if kv.Key != "k" || kv.Value != "v" {
			t.Fatal("wrong")
		}
	})
}

func Test_Cov23_KeyAnyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov23_KeyAnyValuePair", func() {
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kav.ValueString()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleStringOnce", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		if sso.IsEmpty() || !sso.IsDefined() || sso.Value() != "hello" {
			t.Fatal("basic checks failed")
		}
		sso2 := corestr.New.SimpleStringOnce.Empty()
		if !sso2.IsEmpty() {
			t.Fatal("expected empty")
		}
		sso2.GetSetOnce("world")
		if sso2.Value() != "world" {
			t.Fatal("wrong")
		}
		sso2.GetSetOnce("again") // should not change
		if sso2.Value() != "world" {
			t.Fatal("should not change")
		}
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NewCreator methods — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_NewCreator(t *testing.T) {
	safeTest(t, "Test_Cov23_NewCreator", func() {
		_ = corestr.New.Collection.Empty()
		_ = corestr.New.Collection.Cap(5)
		_ = corestr.New.Collection.Strings([]string{"a"})
		_ = corestr.New.Collection.StringsOptions(true, []string{"a"})
		_ = corestr.New.Hashmap.Cap(5)
		_ = corestr.New.Hashmap.UsingMap(map[string]string{"a": "1"})
		_ = corestr.New.Hashset.Cap(5)
		_ = corestr.New.Hashset.Empty()
		_ = corestr.New.Hashset.Strings([]string{"a"})
		_ = corestr.New.SimpleSlice.Empty()
		_ = corestr.New.SimpleSlice.Lines("a", "b")
		_ = corestr.New.CollectionsOfCollection.Empty()
		_ = corestr.New.CollectionsOfCollection.Cap(5)
		_ = corestr.New.CollectionsOfCollection.Strings([]string{"a"})
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetsCollection — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov23_HashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Cov23_HashsetsCollection", func() {
		hc := corestr.New.HashsetsCollection.Empty()
		if !hc.IsEmpty() || hc.HasItems() || hc.Length() != 0 {
			t.Fatal("basic checks failed")
		}
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")
		hc.Add(hs)
		if hc.Length() != 1 || hc.IsEmpty() || !hc.HasItems() {
			t.Fatal("filled checks failed")
		}
	})
}
