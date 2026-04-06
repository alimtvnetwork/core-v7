package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══ SimpleSlice comprehensive ═══

func Test_C39_SimpleSlice_Add_Adds(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Add_Adds", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.Add("a")
		ss.Adds("b", "c")
		ss.Append("d")

		// Act
		actual := args.Map{"result": ss.Length() != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		ss.Adds()
		ss.Append()
	})
}

func Test_C39_SimpleSlice_AddIf_AddsIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddIf_AddsIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddIf(false, "skip")
		ss.AddIf(true, "keep")
		ss.AddsIf(false, "x")
		ss.AddsIf(true, "y")

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)
		ss.AddSplit("a,b,c", ",")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddError(nil)
		ss.AddError(errForTest)

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddStruct_AddPointer(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddStruct_AddPointer", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddStruct(false, "hello")
		ss.AddStruct(true, nil)
		ss.AddPointer(false, "world")
		ss.AddPointer(true, nil)

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AppendFmt("hello %d", 42)
		ss.AppendFmt("", )

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AppendFmtIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AppendFmtIf(true, "x=%d", 1)
		ss.AppendFmtIf(false, "y=%d", 2)

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsTitleValue("Name", "Alice")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsTitleValueIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsTitleValueIf(true, "N", "A")
		ss.AddAsTitleValueIf(false, "N", "B")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsCurlyTitleWrap("Key", "Val")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddAsCurlyTitleWrapIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(2)
		ss.AddAsCurlyTitleWrapIf(true, "K", "V")
		ss.AddAsCurlyTitleWrapIf(false, "K", "V")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "c")
		ss.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		ss.InsertAt(-1, "x") // no-op
		ss.InsertAt(100, "y") // no-op
	})
}

func Test_C39_SimpleSlice_Accessors(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Accessors", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.FirstOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.LastOrDefault() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.FirstDynamic() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.LastDynamic() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.FirstOrDefaultDynamic() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.LastOrDefaultDynamic() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.Count() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.LastIndex() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.HasIndex(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// empty
		es := corestr.Empty.SimpleSlice()
		actual = args.Map{"result": es.FirstOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": es.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilSS *corestr.SimpleSlice
		actual = args.Map{"result": nilSS.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_Skip_Take_Limit(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Skip_Take_Limit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		s := ss.Skip(1)

		// Act
		actual := args.Map{"result": len(s) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s2 := ss.Skip(10)
		actual = args.Map{"result": len(s2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		tk := ss.Take(2)
		actual = args.Map{"result": len(tk) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		tk2 := ss.Take(10)
		actual = args.Map{"result": len(tk2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		lm := ss.Limit(1)
		actual = args.Map{"result": len(lm) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IsContains("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("abc", "xyz")

		// Act
		actual := args.Map{"result": ss.IsContainsFunc("abc", func(a, b string) bool { return a == b })}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IsContainsFunc("a", func(a, b string) bool { return true })}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IndexOf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IndexOf("b") != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.IndexOf("z") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IndexOf("a") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("abc")

		// Act
		actual := args.Map{"result": ss.IndexOfFunc("abc", func(a, b string) bool { return a == b }) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IndexOfFunc("a", func(a, b string) bool { return true }) != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")
		count := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{"result": count != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().CountFunc(func(i int, s string) bool { return true }) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		ts := ss.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"result": ts.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		tj := ss.TranspileJoin(func(s string) string { return s }, ",")
		actual = args.Map{"result": tj == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// empty transpile
		_ = corestr.Empty.SimpleSlice().Transpile(func(s string) string { return s })
	})
}

func Test_C39_SimpleSlice_Join_Methods(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Join_Methods", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.Join(",") == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinLine() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinSpace() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinComma() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinCsv() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinCsvLine() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinWith(",") == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.JoinCsvString(",") == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().Join(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinWith(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinCsvString(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		r := ss.JoinLineEofLine()

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// already has suffix
		ss2 := corestr.New.SimpleSlice.Lines("a\n")
		_ = ss2.JoinLineEofLine()
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinLineEofLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_PrependJoin_AppendJoin(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_PrependJoin_AppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b")
		r := ss.PrependJoin(",", "a")

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := ss.AppendJoin(",", "c")
		actual = args.Map{"result": r2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b")
		ss.PrependAppend([]string{"a"}, []string{"c"})

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a,b", "c")
		split := ss.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": split.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsEqual(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilSS *corestr.SimpleSlice
		actual = args.Map{"result": nilSS.IsEqual(nil)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": a.IsEqual(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		c := corestr.New.SimpleSlice.Lines("a")
		actual = args.Map{"result": a.IsEqual(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// empty
		e1 := corestr.New.SimpleSlice.Empty()
		e2 := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": e1.IsEqual(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsEqualLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.IsEqualLines([]string{"a"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.IsEqualLines([]string{"a", "c"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualUnorderedLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsUnorderedEqualRaw(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsUnorderedEqualRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqualRaw(true, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.IsUnorderedEqualRaw(false, "a", "b")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsUnorderedEqual", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		other := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsUnorderedEqual(true, other)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		e1 := corestr.New.SimpleSlice.Empty()
		e2 := corestr.New.SimpleSlice.Empty()
		actual = args.Map{"result": e1.IsUnorderedEqual(true, e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsDistinctEqual", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsDistinctEqualRaw("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		other := corestr.New.SimpleSlice.Lines("a", "b")
		actual = args.Map{"result": ss.IsDistinctEqual(other)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualByFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.IsEqualByFuncLinesSplit(false, "\n", "a\nb", func(i int, l, r string) bool { return l == r })}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_Sort_Reverse(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Sort_Reverse", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")
		ss.Sort()

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		ss.Reverse()
		actual = args.Map{"result": ss.First() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// 2 items
		ss2 := corestr.New.SimpleSlice.Lines("b", "a")
		ss2.Reverse()
		actual = args.Map{"result": ss2.First() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// 1 item
		ss3 := corestr.New.SimpleSlice.Lines("a")
		ss3.Reverse()
	})
}

func Test_C39_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_ConcatNew", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		cn := ss.ConcatNew("b", "c")

		// Act
		actual := args.Map{"result": cn.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cns := ss.ConcatNewStrings("d")
		actual = args.Map{"result": len(cns) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_ConcatNewSimpleSlices", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		ss2 := corestr.New.SimpleSlice.Lines("b")
		cn := ss.ConcatNewSimpleSlices(ss2)

		// Act
		actual := args.Map{"result": cn.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_Clone(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Clone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		c := ss.Clone(true)

		// Act
		actual := args.Map{"result": c.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cp := ss.ClonePtr(true)
		actual = args.Map{"result": cp.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_ = ss.DeepClone()
		_ = ss.ShadowClone()
		var nilSS *corestr.SimpleSlice
		actual = args.Map{"result": nilSS.ClonePtr(true) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_Collection_Hashset(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Collection_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.Collection(false).Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.ToCollection(false).Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.Hashset().Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Strings_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(ss.Strings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(ss.List()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(ss.SafeStrings()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(corestr.Empty.SimpleSlice().SafeStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_String", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_CsvStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": len(ss.CsvStrings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(corestr.Empty.SimpleSlice().CsvStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AsError(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AsError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("err")

		// Act
		actual := args.Map{"result": ss.AsError(",") == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.AsDefaultError() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().AsError(",") != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_RemoveIndexes", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		newSS, err := ss.RemoveIndexes(1)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": newSS.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_, err2 := corestr.Empty.SimpleSlice().RemoveIndexes(0)
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C39_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_DistinctDiff", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		r := ss.DistinctDiffRaw("b", "c")

		// Act
		actual := args.Map{"result": len(r) < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		other := corestr.New.SimpleSlice.Lines("b", "c")
		r2 := ss.DistinctDiff(other)
		actual = args.Map{"result": len(r2) < 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_AddedRemovedLinesDiff", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := ss.AddedRemovedLinesDiff("b", "c")

		// Act
		actual := args.Map{"result": len(added) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected added", actual)
		actual = args.Map{"result": len(removed) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected removed", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		j := ss.Json()

		// Act
		actual := args.Map{"hasError": j.HasError()}

		// Assert
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		jp := ss.JsonPtr()
		actual = args.Map{"hasError": jp.HasError()}
		expected = args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns no error", actual)
		actual = args.Map{"result": ss.JsonModelAny() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		b, err := ss.MarshalJSON()
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		ss2 := corestr.New.SimpleSlice.Empty()
		err2 := ss2.UnmarshalJSON(b)
		actual = args.Map{"result": err2}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
		err3 := ss2.UnmarshalJSON([]byte(`{bad`))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_ParseInjectUsingJson", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Empty()
		result, err := ss2.ParseInjectUsingJson(jr)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": result.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_InterfaceCasts", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		actual := args.Map{"result": ss.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.AsJsoner() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.AsJsonParseSelfInjector() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": ss.AsJsonMarshaller() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SimpleSlice_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Clear_Dispose", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Clear()
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		ss2 := corestr.New.SimpleSlice.Lines("a")
		ss2.Dispose()
		var nilSS *corestr.SimpleSlice
		actual = args.Map{"result": nilSS.Clear() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		nilSS.Dispose()
	})
}

func Test_C39_SimpleSlice_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C39_SimpleSlice_Serialize_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		b, err := ss.Serialize()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": len(b) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var target []string
		err2 := ss.Deserialize(&target)
		actual = args.Map{"result": err2}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
	})
}

	// ── newSimpleSliceCreator ──

func Test_C39_NewSimpleSliceCreator(t *testing.T) {
	safeTest(t, "Test_C39_NewSimpleSliceCreator", func() {
		s1 := corestr.New.SimpleSlice.Cap(5)
		actual := args.Map{"result": s1.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s2 := corestr.New.SimpleSlice.Cap(-1)
		actual = args.Map{"result": s2.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s3 := corestr.New.SimpleSlice.Default()
		actual = args.Map{"result": s3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s4 := corestr.New.SimpleSlice.Lines("a", "b")
		actual = args.Map{"result": s4.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s5 := corestr.New.SimpleSlice.SpreadStrings("a")
		actual = args.Map{"result": s5.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s6 := corestr.New.SimpleSlice.Create([]string{"a"})
		actual = args.Map{"result": s6.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s7 := corestr.New.SimpleSlice.StringsPtr([]string{"a"})
		actual = args.Map{"result": s7.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s8 := corestr.New.SimpleSlice.StringsPtr(nil)
		actual = args.Map{"result": s8.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s9 := corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		actual = args.Map{"result": s9.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s10 := corestr.New.SimpleSlice.StringsOptions(false, []string{"a"})
		actual = args.Map{"result": s10.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s11 := corestr.New.SimpleSlice.StringsClone([]string{"a"})
		actual = args.Map{"result": s11.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s12 := corestr.New.SimpleSlice.StringsClone(nil)
		actual = args.Map{"result": s12.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s13 := corestr.New.SimpleSlice.Direct(true, []string{"a"})
		actual = args.Map{"result": s13.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s14 := corestr.New.SimpleSlice.Direct(false, []string{"a"})
		actual = args.Map{"result": s14.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s15 := corestr.New.SimpleSlice.Direct(false, nil)
		actual = args.Map{"result": s15.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s16 := corestr.New.SimpleSlice.UsingLines(true, "a")
		actual = args.Map{"result": s16.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s17 := corestr.New.SimpleSlice.UsingLines(false, "a")
		actual = args.Map{"result": s17.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s18 := corestr.New.SimpleSlice.Split("a,b", ",")
		actual = args.Map{"result": s18.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s19 := corestr.New.SimpleSlice.SplitLines("a\nb")
		actual = args.Map{"result": s19.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s20 := corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b")
		actual = args.Map{"result": s20.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s21 := corestr.New.SimpleSlice.UsingLine("a\nb")
		actual = args.Map{"result": s21.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s22 := corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.StringsSpreadItems("a"))
		actual = args.Map{"result": s22.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s23 := corestr.New.SimpleSlice.Hashset(corestr.Empty.Hashset())
		actual = args.Map{"result": s23.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s24 := corestr.New.SimpleSlice.Map(map[string]string{"a": "1"})
		actual = args.Map{"result": s24.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s25 := corestr.New.SimpleSlice.ByLen([]string{"a", "b"})
		actual = args.Map{"result": s25 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

	// ═══ SimpleStringOnce ═══

func Test_C39_SSO_SetGet(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetGet", func() {
		sso := corestr.New.SimpleStringOnce.Create("hello", true)
		actual := args.Map{"result": sso.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsDefined()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsUninitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsInvalid()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetOnUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		err := sso.SetOnUninitialized("val")
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		err2 := sso.SetOnUninitialized("val2")
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C39_SSO_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_C39_SSO_GetSetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetSetOnce("first")
		actual := args.Map{"result": v != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2 := sso.GetSetOnce("second")
		actual = args.Map{"result": v2 != "first"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_GetOnce(t *testing.T) {
	safeTest(t, "Test_C39_SSO_GetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetOnce()
		actual := args.Map{"result": v != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C39_SSO_GetOnceFunc", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetOnceFunc(func() string { return "computed" })
		actual := args.Map{"result": v != "computed"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2 := sso.GetOnceFunc(func() string { return "other" })
		actual = args.Map{"result": v2 != "computed"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_C39_SSO_SetOnceIfUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		actual := args.Map{"result": sso.SetOnceIfUninitialized("v")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.SetOnceIfUninitialized("v2")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Invalidate_Reset(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Invalidate_Reset", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		sso.Invalidate()
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		sso2 := corestr.New.SimpleStringOnce.Init("world")
		sso2.Reset()
		actual = args.Map{"result": sso2.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_NumericConversions(t *testing.T) {
	safeTest(t, "Test_C39_SSO_NumericConversions", func() {
		sso := corestr.New.SimpleStringOnce.Init("42")
		actual := args.Map{"result": sso.Int() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.ValueDefInt() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.ValueInt(0) != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.Byte() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.ValueDefByte() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.ValueByte(0) != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.Int16() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.Int32() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v, ok := sso.Uint16()
		actual = args.Map{"result": ok || v != 42}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2, ok2 := sso.Uint32()
		actual = args.Map{"result": ok2 || v2 != 42}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// non-numeric
		bad := corestr.New.SimpleStringOnce.Init("abc")
		actual = args.Map{"result": bad.Int() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": bad.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": bad.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": bad.Int32() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Float64(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Float64", func() {
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		actual := args.Map{"result": sso.ValueFloat64(0) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.ValueDefFloat64() == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		bad := corestr.New.SimpleStringOnce.Init("abc")
		actual = args.Map{"result": bad.ValueFloat64(1.0) != 1.0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Boolean(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Boolean", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		actual := args.Map{"result": sso.Boolean(false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.BooleanDefault()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		sso2 := corestr.New.SimpleStringOnce.Init("true")
		actual = args.Map{"result": sso2.Boolean(false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		bad := corestr.New.SimpleStringOnce.Init("xyz")
		actual = args.Map{"result": bad.Boolean(false)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		uninit := corestr.New.SimpleStringOnce.Empty()
		actual = args.Map{"result": uninit.Boolean(true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_IsSetter(t *testing.T) {
	safeTest(t, "Test_C39_SSO_IsSetter", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		actual := args.Map{"result": sso.IsSetter(false).IsTrue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		uninit := corestr.New.SimpleStringOnce.Empty()
		actual = args.Map{"result": uninit.IsSetter(true).IsTrue()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		bad := corestr.New.SimpleStringOnce.Init("xyz")
		_ = bad.IsSetter(false)
		ssoTrue := corestr.New.SimpleStringOnce.Init("true")
		actual = args.Map{"result": ssoTrue.IsSetter(false).IsTrue()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_WithinRange(t *testing.T) {
	safeTest(t, "Test_C39_SSO_WithinRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("50")
		v, ok := sso.WithinRange(true, 0, 100)
		actual := args.Map{"result": ok || v != 50}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v2, ok2 := sso.WithinRange(true, 60, 100)
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v2 != 60}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v3, ok3 := sso.WithinRange(true, 0, 10)
		actual = args.Map{"result": ok3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": v3 != 10}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		v4, ok4 := sso.WithinRange(false, 60, 100)
		actual = args.Map{"result": ok4 || v4 != 50}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		_, ok5 := sso.WithinRangeDefault(0, 100)
		actual = args.Map{"result": ok5}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_StringMethods(t *testing.T) {
	safeTest(t, "Test_C39_SSO_StringMethods", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{"result": sso.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.StringPtr() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.SafeValue() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.Trim() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.HasValidNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.HasValidNonWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.Is("hello")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsAnyOf("hello", "world")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsContains("ell")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsAnyContains("xyz", "ell")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsAnyContains("xyz")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsEqualNonSensitive("HELLO")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// nil ptr
		var nilSSO *corestr.SimpleStringOnce
		actual = args.Map{"result": nilSSO.String() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": nilSSO.StringPtr() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// uninit safe value
		uninit := corestr.New.SimpleStringOnce.Empty()
		actual = args.Map{"result": uninit.SafeValue() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		// IsAnyOf empty
		actual = args.Map{"result": sso.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_ValueBytes(t *testing.T) {
	safeTest(t, "Test_C39_SSO_ValueBytes", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		actual := args.Map{"result": len(sso.ValueBytes()) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": len(sso.ValueBytesPtr()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C39_SSO_ConcatNew", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		cn := sso.ConcatNew(" world")
		actual := args.Map{"result": cn.Value() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C39_SSO_ConcatNewUsingStrings", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		cn := sso.ConcatNewUsingStrings(",", "world")
		actual := args.Map{"result": cn.Value() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Clone(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Clone", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		c := sso.Clone()
		actual := args.Map{"result": c.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cp := sso.ClonePtr()
		actual = args.Map{"result": cp.Value() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		cu := sso.CloneUsingNewVal("world")
		actual = args.Map{"result": cu.Value() != "world"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		var nilSSO *corestr.SimpleStringOnce
		actual = args.Map{"result": nilSSO.ClonePtr() != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
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
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		sso.SetUnInit()
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Split(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Split", func() {
		sso := corestr.New.SimpleStringOnce.Init("a,b,c")
		actual := args.Map{"result": len(sso.Split(",")) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		l, r := sso.SplitLeftRight(",")
		actual = args.Map{"result": l != "a" || r == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		l2, r2 := sso.SplitLeftRightTrim(",")
		actual = args.Map{"result": l2 != "a" || r2 == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
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
		actual := args.Map{"result": l != "nosep" || r != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		l2, r2 := sso.SplitLeftRightTrim(",")
		actual = args.Map{"result": l2 != "nosep" || r2 != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Regex(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Regex", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello123")
		actual := args.Map{"result": sso.IsRegexMatches(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r, ok := sso.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": ok || len(r) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		r2 := sso.RegexFindAllStrings(nil, -1)
		actual = args.Map{"result": len(r2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
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
		actual := args.Map{"hasError": j.HasError()}
		expected := args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "Json returns no error", actual)
		jp := sso.JsonPtr()
		actual = args.Map{"hasError": jp.HasError()}
		expected = args.Map{"hasError": false}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns no error", actual)
		actual = args.Map{"result": sso.JsonModelAny() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		b, err := sso.MarshalJSON()
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		sso2 := corestr.New.SimpleStringOnce.Empty()
		err2 := sso2.UnmarshalJSON(b)
		actual = args.Map{"result": err2}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err2", actual)
	})
}

func Test_C39_SSO_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_C39_SSO_InterfaceCasts", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		actual := args.Map{"result": sso.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.AsJsoner() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.AsJsonParseSelfInjector() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		actual = args.Map{"result": sso.AsJsonMarshaller() == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

func Test_C39_SSO_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_C39_SSO_Serialize_Deserialize", func() {
		sso := corestr.New.SimpleStringOnce.Init("hi")
		b, err := sso.Serialize()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		actual = args.Map{"result": len(b) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}

	// ── newSimpleStringOnceCreator ──

func Test_C39_NewSSOCreator(t *testing.T) {
	safeTest(t, "Test_C39_NewSSOCreator", func() {
		s1 := corestr.New.SimpleStringOnce.Init("hi")
		actual := args.Map{"result": s1.Value() != "hi"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s2 := corestr.New.SimpleStringOnce.InitPtr("hi")
		actual = args.Map{"result": s2.Value() != "hi"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s3 := corestr.New.SimpleStringOnce.Create("v", true)
		actual = args.Map{"result": s3.Value() != "v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s4 := corestr.New.SimpleStringOnce.CreatePtr("v", true)
		actual = args.Map{"result": s4.Value() != "v"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s5 := corestr.New.SimpleStringOnce.Uninitialized("val")
		actual = args.Map{"result": s5.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s6 := corestr.New.SimpleStringOnce.Empty()
		actual = args.Map{"result": s6.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
		s7 := corestr.New.SimpleStringOnce.Any(false, 42, true)
		actual = args.Map{"result": s7.Value() == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "assertion", actual)
	})
}
