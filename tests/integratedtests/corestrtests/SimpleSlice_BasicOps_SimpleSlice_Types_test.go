package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleSlice_BasicOps_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_BasicOps", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{"result": ss.Length() != 3 || ss.IsEmpty() || !ss.HasAnyItem() || ss.LastIndex() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
		actual = args.Map{"result": ss.Count() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "count wrong", actual)
	})
}

func Test_SimpleSlice_AddMethods_FromSimpleSliceBasicOpsS(t *testing.T) {
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

func Test_SimpleSlice_FirstLast_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_FirstLast", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.First() != "a" || ss.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		actual = args.Map{"result": ss.FirstOrDefault() != "a" || ss.LastOrDefault() != "b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		_ = ss.FirstDynamic()
		_ = ss.LastDynamic()
		_ = ss.FirstOrDefaultDynamic()
		_ = ss.LastOrDefaultDynamic()
		e := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": e.FirstOrDefault() != "" || e.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_SkipTakeLimit_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_SkipTakeLimit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{"result": len(ss.Skip(1)) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong skip", actual)
		actual = args.Map{"result": len(ss.Take(2)) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong take", actual)
		actual = args.Map{"result": len(ss.Limit(2)) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong limit", actual)
		_ = ss.SkipDynamic(1)
		_ = ss.TakeDynamic(2)
		_ = ss.LimitDynamic(2)
	})
}

func Test_SimpleSlice_IsContains_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsContains("x")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_IsContainsFunc_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("abc", "def")

		// Act
		actual := args.Map{"result": ss.IsContainsFunc("abc", func(a, b string) bool { return a == b })}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleSlice_IndexOf_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IndexOf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IndexOf("b") != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong index", actual)
		actual = args.Map{"result": ss.IndexOf("x") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_SimpleSlice_IndexOfFunc_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		idx := ss.IndexOfFunc("b", func(a, b string) bool { return a == b })

		// Act
		actual := args.Map{"result": idx != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong index", actual)
	})
}

func Test_SimpleSlice_HasIndex_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_HasIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.HasIndex(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_InsertAt_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		ss.InsertAt(-1, "x") // out of range, skip
		ss.InsertAt(100, "y") // out of range, skip
	})
}

func Test_SimpleSlice_WrapQuotes_FromSimpleSliceBasicOpsS(t *testing.T) {
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

func Test_SimpleSlice_Join_FromSimpleSliceBasicOpsS(t *testing.T) {
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

func Test_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Strings_List", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.Strings()
		_ = ss.List()
	})
}

func Test_SimpleSlice_Transpile_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Transpile", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"result": result.IsContains("a!")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected transpiled", actual)
	})
}

func Test_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_TranspileJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		s := ss.TranspileJoin(func(s string) string { return s + "!" }, ",")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SimpleSlice_EachItemSplitBy_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a,b", "c")
		result := ss.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": result.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SimpleSlice_Concat(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Concat", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		newSS := ss.ConcatNew("b", "c")

		// Act
		actual := args.Map{"result": newSS.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_ = ss.ConcatNewStrings("d")
		ss2 := corestr.New.SimpleSlice.Lines("e")
		_ = ss.ConcatNewSimpleSlices(ss2)
	})
}

func Test_SimpleSlice_PrependAppend_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b")
		ss.PrependAppend([]string{"a"}, []string{"c"})

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SimpleSlice_PrependAppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_PrependAppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b")
		s := ss.PrependJoin(",", "a")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		s2 := ss.AppendJoin(",", "c")
		actual = args.Map{"result": s2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SimpleSlice_IsEqual_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqualLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsEqualLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqualUnorderedLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_SimpleSlice_Collection_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Collection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		c := ss.Collection(false)

		// Act
		actual := args.Map{"result": c.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = ss.ToCollection(true)
	})
}

func Test_SimpleSlice_CountFunc_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		cnt := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"result": cnt != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_AsError_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_AsError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("e1", "e2")
		err := ss.AsError(",")

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		_ = ss.AsDefaultError()
		empty := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": empty.AsError(",") != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_SimpleSlice_String_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_String", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		empty := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": empty.String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_Sort_Reverse_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Sort_Reverse", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")
		ss.Sort()

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong sort", actual)
		ss.Reverse()
		actual = args.Map{"result": ss.First() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong reverse", actual)
	})
}

func Test_SimpleSlice_Hashset_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		hs := ss.Hashset()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleSlice_CsvStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		csv := ss.CsvStrings()

		// Act
		actual := args.Map{"result": len(csv) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_JsonModel(t *testing.T) {
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

func Test_CollectionsOfCollection_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Empty()

		// Act
		actual := args.Map{"result": cc.IsEmpty() || cc.HasItems() || cc.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		cc.Add(c1).Add(c2)
		actual = args.Map{"result": cc.Length() != 2 || cc.IsEmpty() || !cc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "filled checks failed", actual)
		actual = args.Map{"result": cc.AllIndividualItemsLength() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		list := cc.List(0)
		actual = args.Map{"result": len(list) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		col := cc.ToCollection()
		actual = args.Map{"result": col.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_ = cc.Items()
		_ = cc.String()
	})
}

func Test_CollectionsOfCollection_AddStrings_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_AddStrings", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Empty()
		cc.AddStrings(true, []string{"a", "b"})

		// Act
		actual := args.Map{"result": cc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CollectionsOfCollection_AddsStringsOfStrings_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_AddsStringsOfStrings", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Empty()
		cc.AddsStringsOfStrings(true, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"result": cc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CollectionsOfCollection_AddCollections_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_AddCollections", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Empty()
		c := *corestr.New.Collection.Strings([]string{"a"})
		cc.AddCollections(c)
		cc.Adds(c)

		// Act
		actual := args.Map{"result": cc.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CollectionsOfCollection_JsonOps(t *testing.T) {
	safeTest(t, "Test_Cov23_CollectionsOfCollection_JsonOps", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Empty()
		cc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = cc.JsonModel()
		_ = cc.JsonModelAny()
		b, err := cc.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
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

func Test_HashmapDiff_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_HashmapDiff", func() {
		// Arrange
		hd := corestr.HashmapDiff{"a": "1", "b": "2"}

		// Act
		actual := args.Map{"result": hd.IsEmpty() || !hd.HasAnyItem() || hd.Length() != 2 || hd.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
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

func Test_KeyValuePair_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.Key != "k" || kv.Value != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
	})
}

func Test_KeyAnyValuePair_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_KeyAnyValuePair", func() {
		// Arrange
		kav := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		s := kav.ValueString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleStringOnce_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_SimpleStringOnce", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": sso.IsEmpty() || !sso.IsDefined() || sso.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
		sso2 := corestr.New.SimpleStringOnce.Empty()
		actual = args.Map{"result": sso2.IsEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		sso2.GetSetOnce("world")
		actual = args.Map{"result": sso2.Value() != "world"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "wrong", actual)
		sso2.GetSetOnce("again") // should not change
		actual = args.Map{"result": sso2.Value() != "world"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not change", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NewCreator methods — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewCreator(t *testing.T) {
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

func Test_HashsetsCollection_FromSimpleSliceBasicOpsS(t *testing.T) {
	safeTest(t, "Test_Cov23_HashsetsCollection", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()

		// Act
		actual := args.Map{"result": hc.IsEmpty() || hc.HasItems() || hc.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "basic checks failed", actual)
		hs := corestr.New.Hashset.Cap(2)
		hs.Add("a")
		hc.Add(hs)
		actual = args.Map{"result": hc.Length() != 1 || hc.IsEmpty() || !hc.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "filled checks failed", actual)
	})
}
