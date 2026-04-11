package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 8: Add variants, accessors, search, wrap (L1-700)
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovSS1_01_Add(t *testing.T) {
	safeTest(t, "Test_CovSS1_01_Add", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.Add("a")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_02_AddSplit(t *testing.T) {
	safeTest(t, "Test_CovSS1_02_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddSplit("a,b,c", ",")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovSS1_03_AddIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_03_AddIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddIf(false, "skip")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddIf(true, "keep")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_04_Adds_Append(t *testing.T) {
	safeTest(t, "Test_CovSS1_04_Adds_Append", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.Adds("a", "b")

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss.Adds()
		ss.Append("c")
		actual = args.Map{"result": ss.Length() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		ss.Append()
	})
}

func Test_CovSS1_05_AppendFmt(t *testing.T) {
	safeTest(t, "Test_CovSS1_05_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AppendFmt("hello %s", "world")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty format and no values
		ss.AppendFmt("")
	})
}

func Test_CovSS1_06_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_06_AppendFmtIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AppendFmtIf(false, "skip %s", "x")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AppendFmtIf(true, "keep %s", "x")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty format
		ss.AppendFmtIf(true, "")
	})
}

func Test_CovSS1_07_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_CovSS1_07_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsTitleValue("Name", "Alice")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_08_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_CovSS1_08_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsCurlyTitleWrap("Name", "Alice")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_09_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_09_AddAsCurlyTitleWrapIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsCurlyTitleWrapIf(false, "skip", "x")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddAsCurlyTitleWrapIf(true, "keep", "x")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_10_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_10_AddAsTitleValueIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddAsTitleValueIf(false, "skip", "x")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddAsTitleValueIf(true, "keep", "x")
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_11_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovSS1_11_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})
		ss.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// out of range
		ss.InsertAt(-1, "x")
		ss.InsertAt(100, "x")
	})
}

func Test_CovSS1_12_AddStruct(t *testing.T) {
	safeTest(t, "Test_CovSS1_12_AddStruct", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddStruct(true, struct{ Name string }{"Alice"})

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss.AddStruct(true, nil)
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected still 1", actual)
	})
}

func Test_CovSS1_13_AddPointer(t *testing.T) {
	safeTest(t, "Test_CovSS1_13_AddPointer", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		v := "hello"
		ss.AddPointer(false, &v)

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		ss.AddPointer(false, nil)
	})
}

func Test_CovSS1_14_AddsIf(t *testing.T) {
	safeTest(t, "Test_CovSS1_14_AddsIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddsIf(false, "a", "b")

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddsIf(true, "a", "b")
		actual = args.Map{"result": ss.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS1_15_AddError(t *testing.T) {
	safeTest(t, "Test_CovSS1_15_AddError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})
		ss.AddError(nil)

		// Act
		actual := args.Map{"result": ss.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		ss.AddError(fmt.Errorf("oops"))
		actual = args.Map{"result": ss.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_16_AsDefaultError_AsError(t *testing.T) {
	safeTest(t, "Test_CovSS1_16_AsDefaultError_AsError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"err1", "err2"})
		e := ss.AsDefaultError()

		// Act
		actual := args.Map{"result": e == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		e2 := ss.AsError(",")
		actual = args.Map{"result": e2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		// empty
		empty := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": empty.AsError(",") != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_CovSS1_17_First_Last_Dynamic(t *testing.T) {
	safeTest(t, "Test_CovSS1_17_First_Last_Dynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ss.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		_ = ss.FirstDynamic()
		_ = ss.LastDynamic()
	})
}

func Test_CovSS1_18_FirstOrDefault_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_CovSS1_18_FirstOrDefault_LastOrDefault", func() {
		// Arrange
		empty := corestr.New.SimpleSlice.Strings([]string{})

		// Act
		actual := args.Map{"result": empty.FirstOrDefault() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": empty.LastOrDefault() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		_ = empty.FirstOrDefaultDynamic()
		_ = empty.LastOrDefaultDynamic()

		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual = args.Map{"result": ss.FirstOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ss.LastOrDefault() != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_CovSS1_19_Skip_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_CovSS1_19_Skip_SkipDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		r := ss.Skip(1)

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// skip all
		r2 := ss.Skip(10)
		actual = args.Map{"result": len(r2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		_ = ss.SkipDynamic(1)
		_ = ss.SkipDynamic(10)
	})
}

func Test_CovSS1_20_Take_TakeDynamic_Limit(t *testing.T) {
	safeTest(t, "Test_CovSS1_20_Take_TakeDynamic_Limit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
		r := ss.Take(2)

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// take all
		r2 := ss.Take(10)
		actual = args.Map{"result": len(r2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_ = ss.TakeDynamic(2)
		_ = ss.TakeDynamic(10)
		_ = ss.Limit(1)
		_ = ss.LimitDynamic(1)
	})
}

func Test_CovSS1_21_Length_Count_CountFunc(t *testing.T) {
	safeTest(t, "Test_CovSS1_21_Length_Count_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "bb", "ccc"})

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": ss.Count() != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		c := ss.CountFunc(func(i int, s string) bool { return len(s) > 1 })
		actual = args.Map{"result": c != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.CountFunc(func(i int, s string) bool { return true }) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSS1_22_IsEmpty_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_CovSS1_22_IsEmpty_HasAnyItem", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{})

		// Act
		actual := args.Map{"result": ss.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ss.HasAnyItem()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		ss.Add("a")
		actual = args.Map{"result": ss.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
		actual = args.Map{"result": ss.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected items", actual)
	})
}

func Test_CovSS1_23_IsContains(t *testing.T) {
	safeTest(t, "Test_CovSS1_23_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsContains("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovSS1_24_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_CovSS1_24_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc", "def"})
		found := ss.IsContainsFunc("abc", func(item, searching string) bool {
			return item == searching
		})

		// Act
		actual := args.Map{"result": found}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsContainsFunc("x", func(a, b string) bool { return a == b })}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovSS1_25_IndexOf_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_CovSS1_25_IndexOf_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{"result": ss.IndexOf("b") != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ss.IndexOf("z") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
		idx := ss.IndexOfFunc("b", func(item, searching string) bool {
			return item == searching
		})
		actual = args.Map{"result": idx != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IndexOf("a") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
		actual = args.Map{"result": e.IndexOfFunc("a", func(a, b string) bool { return a == b }) != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_CovSS1_26_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_CovSS1_26_LastIndex_HasIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ss.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.HasIndex(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": ss.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovSS1_27_Strings_List(t *testing.T) {
	safeTest(t, "Test_CovSS1_27_Strings_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": len(ss.Strings()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(ss.List()) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSS1_28_WrapQuotes(t *testing.T) {
	safeTest(t, "Test_CovSS1_28_WrapQuotes", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := ss.WrapDoubleQuote()

		// Act
		actual := args.Map{"result": (*r)[0] != `"a"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		ss2 := corestr.New.SimpleSlice.Strings([]string{"a"})
		r2 := ss2.WrapSingleQuote()
		actual = args.Map{"result": (*r2)[0] != "'a'"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
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
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.Transpile(func(s string) string { return s + "!" })

		// Act
		actual := args.Map{"result": (*r)[0] != "a!"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a!", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		_ = e.Transpile(func(s string) string { return s })
		// TranspileJoin
		s := ss.TranspileJoin(func(s string) string { return s }, ",")
		actual = args.Map{"result": s == ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovSS1_30_Hashset(t *testing.T) {
	safeTest(t, "Test_CovSS1_30_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "a"})
		hs := ss.Hashset()

		// Act
		actual := args.Map{"result": hs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSS1_31_Join_JoinLine_JoinSpace_JoinComma(t *testing.T) {
	safeTest(t, "Test_CovSS1_31_Join_JoinLine_JoinSpace_JoinComma", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		_ = ss.JoinLine()
		_ = ss.JoinSpace()
		_ = ss.JoinComma()
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.Join(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": e.JoinLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovSS1_32_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_CovSS1_32_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.JoinLineEofLine()

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.JoinLineEofLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		// already has suffix
		ss2 := corestr.New.SimpleSlice.Strings([]string{"a\n"})
		_ = ss2.JoinLineEofLine()
	})
}

func Test_CovSS1_33_JoinCsv_JoinCsvLine_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_CovSS1_33_JoinCsv_JoinCsvLine_JoinCsvString", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		_ = ss.JoinCsv()
		_ = ss.JoinCsvLine()
		s := ss.JoinCsvString(",")

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.JoinCsvString(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovSS1_34_CsvStrings(t *testing.T) {
	safeTest(t, "Test_CovSS1_34_CsvStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		r := ss.CsvStrings()

		// Act
		actual := args.Map{"result": len(r) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": len(e.CsvStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSS1_35_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_CovSS1_35_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a,b", "c,d"})
		r := ss.EachItemSplitBy(",")

		// Act
		actual := args.Map{"result": r.Length() != 4}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_CovSS1_36_PrependJoin_AppendJoin(t *testing.T) {
	safeTest(t, "Test_CovSS1_36_PrependJoin_AppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		s := ss.PrependJoin(",", "a")

		// Act
		actual := args.Map{"result": s != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
		s2 := ss.AppendJoin(",", "c")
		actual = args.Map{"result": s2 != "b,c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'b,c', got ''", actual)
	})
}

func Test_CovSS1_37_PrependAppend(t *testing.T) {
	safeTest(t, "Test_CovSS1_37_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		ss.PrependAppend([]string{"a"}, []string{"c"})

		// Act
		actual := args.Map{"result": ss.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// empty prepend/append
		ss.PrependAppend(nil, nil)
	})
}

func Test_CovSS1_38_JoinWith(t *testing.T) {
	safeTest(t, "Test_CovSS1_38_JoinWith", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		r := ss.JoinWith(",")

		// Act
		actual := args.Map{"result": r == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		// empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.JoinWith(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovSS1_39_IsEqual(t *testing.T) {
	safeTest(t, "Test_CovSS1_39_IsEqual", func() {
		// Arrange
		a := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		b := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": a.IsEqual(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
		// nil
		actual = args.Map{"result": a.IsEqual(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		c := corestr.New.SimpleSlice.Strings([]string{"a"})
		actual = args.Map{"result": a.IsEqual(c)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e1 := corestr.New.SimpleSlice.Strings([]string{})
		e2 := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e1.IsEqual(e2)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSS1_40_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_CovSS1_40_IsEqualLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.IsEqualLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsEqualLines([]string{"a", "c"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": ss.IsEqualLines([]string{"a"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false (diff length)", actual)
	})
}

func Test_CovSS1_41_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_CovSS1_41_IsEqualUnorderedLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsEqualUnorderedLines([]string{"a", "c"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		actual = args.Map{"result": ss.IsEqualUnorderedLines([]string{"a"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsEqualUnorderedLines([]string{})}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSS1_42_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_CovSS1_42_IsEqualUnorderedLinesClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b", "a"})

		// Act
		actual := args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "b"})}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a", "c"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// diff length
		actual = args.Map{"result": ss.IsEqualUnorderedLinesClone([]string{"a"})}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// both empty
		e := corestr.New.SimpleSlice.Strings([]string{})
		actual = args.Map{"result": e.IsEqualUnorderedLinesClone([]string{})}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}
