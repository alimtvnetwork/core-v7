package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S11a — SimpleSlice.go Lines 1-600 — Add, Query, Join, Wrap
// ══════════════════════════════════════════════════════════════

func Test_S11_01_SimpleSlice_Add(t *testing.T) {
	safeTest(t, "Test_S11_01_SimpleSlice_Add", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Add("a")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_02_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_S11_02_SimpleSlice_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddSplit("a,b,c", ",")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S11_03_SimpleSlice_AddIf(t *testing.T) {
	safeTest(t, "Test_S11_03_SimpleSlice_AddIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddIf(true, "yes")
		ss.AddIf(false, "no")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_04_SimpleSlice_Adds(t *testing.T) {
	safeTest(t, "Test_S11_04_SimpleSlice_Adds", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Adds("a", "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_05_SimpleSlice_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_S11_05_SimpleSlice_Adds_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Adds()

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_06_SimpleSlice_Append(t *testing.T) {
	safeTest(t, "Test_S11_06_SimpleSlice_Append", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Append("a")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_07_SimpleSlice_Append_Empty(t *testing.T) {
	safeTest(t, "Test_S11_07_SimpleSlice_Append_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Append()

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_08_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_S11_08_SimpleSlice_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("hello %s", "world")

		// Assert
		actual := args.Map{"result": ss.Length() != 1 || ss.First() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello world'", actual)
	})
}

func Test_S11_09_SimpleSlice_AppendFmt_EmptySkip(t *testing.T) {
	safeTest(t, "Test_S11_09_SimpleSlice_AppendFmt_EmptySkip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("")

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_10_SimpleSlice_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_S11_10_SimpleSlice_AppendFmtIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmtIf(true, "val=%d", 42)
		ss.AppendFmtIf(false, "skip")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_11_SimpleSlice_AppendFmtIf_EmptyFormatSkip(t *testing.T) {
	safeTest(t, "Test_S11_11_SimpleSlice_AppendFmtIf_EmptyFormatSkip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmtIf(true, "")

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_12_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_S11_12_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsTitleValue("Name", "John")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_13_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_S11_13_SimpleSlice_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsCurlyTitleWrap("Key", "Val")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_14_SimpleSlice_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_S11_14_SimpleSlice_AddAsCurlyTitleWrapIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsCurlyTitleWrapIf(true, "K", "V")
		ss.AddAsCurlyTitleWrapIf(false, "K2", "V2")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_15_SimpleSlice_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_S11_15_SimpleSlice_AddAsTitleValueIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsTitleValueIf(true, "K", "V")
		ss.AddAsTitleValueIf(false, "K2", "V2")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_16_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_S11_16_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})

		// Act
		ss.InsertAt(1, "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S11_17_SimpleSlice_InsertAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_S11_17_SimpleSlice_InsertAt_OutOfRange", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		ss.InsertAt(-1, "x")
		ss.InsertAt(100, "y")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 — out of range skipped", actual)
	})
}

func Test_S11_18_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_S11_18_SimpleSlice_AddStruct", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddStruct(true, "hello")
		ss.AddStruct(false, nil) // nil skipped

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_19_SimpleSlice_AddPointer(t *testing.T) {
	safeTest(t, "Test_S11_19_SimpleSlice_AddPointer", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddPointer(true, "hello")
		ss.AddPointer(false, nil) // nil skipped

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_20_SimpleSlice_AddsIf(t *testing.T) {
	safeTest(t, "Test_S11_20_SimpleSlice_AddsIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddsIf(true, "a", "b")
		ss.AddsIf(false, "c")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_21_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_S11_21_SimpleSlice_AddError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddError(nil)
		ss.AddError(&testErr{})

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_22_SimpleSlice_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_S11_22_SimpleSlice_AsDefaultError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"err1", "err2"})

		// Act
		err := ss.AsDefaultError()

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_S11_23_SimpleSlice_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_S11_23_SimpleSlice_AsError_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		err := ss.AsError(",")

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_S11_24_SimpleSlice_FirstAndLast(t *testing.T) {
	safeTest(t, "Test_S11_24_SimpleSlice_FirstAndLast", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ss.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		actual = args.Map{"result": ss.FirstDynamic().(string) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a dynamic", actual)
		actual = args.Map{"result": ss.LastDynamic().(string) != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c dynamic", actual)
	})
}

func Test_S11_25_SimpleSlice_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_S11_25_SimpleSlice_FirstOrDefault", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})

		// Act & Assert
		actual := args.Map{"result": ss.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ss.FirstOrDefaultDynamic().(string) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty dynamic", actual)
		actual = args.Map{"result": ss2.FirstOrDefault() != "x"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_S11_26_SimpleSlice_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_S11_26_SimpleSlice_LastOrDefault", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})

		// Act & Assert
		actual := args.Map{"result": ss.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ss.LastOrDefaultDynamic().(string) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty dynamic", actual)
		actual = args.Map{"result": ss2.LastOrDefault() != "x"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_S11_27_SimpleSlice_Skip(t *testing.T) {
	safeTest(t, "Test_S11_27_SimpleSlice_Skip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		skipped := ss.Skip(1)
		skippedAll := ss.Skip(10)

		// Assert
		actual := args.Map{"result": len(skipped) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(skippedAll) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_28_SimpleSlice_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_S11_28_SimpleSlice_SkipDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.SkipDynamic(1)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_29_SimpleSlice_Take(t *testing.T) {
	safeTest(t, "Test_S11_29_SimpleSlice_Take", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		taken := ss.Take(2)
		takenAll := ss.Take(10)

		// Assert
		actual := args.Map{"result": len(taken) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(takenAll) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S11_30_SimpleSlice_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_S11_30_SimpleSlice_TakeDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.TakeDynamic(1)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_S11_31_SimpleSlice_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_S11_31_SimpleSlice_LimitDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.LimitDynamic(1) == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual = args.Map{"result": len(ss.Limit(1)) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_32_SimpleSlice_Length_Count(t *testing.T) {
	safeTest(t, "Test_S11_32_SimpleSlice_Length_Count", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": ss.Length() != 1 || ss.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_33_SimpleSlice_Length_Nil(t *testing.T) {
	safeTest(t, "Test_S11_33_SimpleSlice_Length_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_34_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_S11_34_SimpleSlice_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "bb", "ccc"})

		// Act
		count := ss.CountFunc(func(index int, item string) bool {
			return len(item) > 1
		})

		// Assert
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_35_SimpleSlice_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_S11_35_SimpleSlice_CountFunc_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		count := ss.CountFunc(func(index int, item string) bool { return true })

		// Assert
		actual := args.Map{"result": count != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_36_SimpleSlice_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S11_36_SimpleSlice_IsEmpty", func() {
		// Act & Assert
		actual := args.Map{"result": corestr.Empty.SimpleSlice().IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_37_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_S11_37_SimpleSlice_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.IsContains("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IsContains("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_S11_38_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_S11_38_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc"})

		// Act & Assert
		actual := args.Map{"result": ss.IsContainsFunc("ab", strings.Contains)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IsContainsFunc("a", strings.Contains)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_S11_39_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_S11_39_SimpleSlice_IndexOf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.IndexOf("b") != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ss.IndexOf("z") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IndexOf("a") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1 for empty", actual)
	})
}

func Test_S11_40_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_S11_40_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc", "def"})

		// Act
		idx := ss.IndexOfFunc("de", strings.Contains)

		// Assert
		actual := args.Map{"result": idx != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IndexOfFunc("a", strings.Contains) != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1 for empty", actual)
	})
}

func Test_S11_41_SimpleSlice_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S11_41_SimpleSlice_HasAnyItem", func() {
		// Act & Assert
		actual := args.Map{"result": corestr.New.SimpleSlice.Strings([]string{"a"}).HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_S11_42_SimpleSlice_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_S11_42_SimpleSlice_LastIndex_HasIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ss.HasIndex(0) || !ss.HasIndex(1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.HasIndex(2) || ss.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_S11_43_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_S11_43_SimpleSlice_Strings_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(ss.Strings()) != 1 || len(ss.List()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_44_SimpleSlice_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_S11_44_SimpleSlice_WrapDoubleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": ss.WrapDoubleQuote().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_45_SimpleSlice_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_S11_45_SimpleSlice_WrapSingleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapSingleQuote().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_46_SimpleSlice_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_S11_46_SimpleSlice_WrapTildaQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapTildaQuote().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_47_SimpleSlice_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_S11_47_SimpleSlice_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapDoubleQuoteIfMissing().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_48_SimpleSlice_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_S11_48_SimpleSlice_WrapSingleQuoteIfMissing", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapSingleQuoteIfMissing().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_S11_49_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_S11_49_SimpleSlice_Transpile", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.First() != "A"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected A", actual)
	})
}

func Test_S11_50_SimpleSlice_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_S11_50_SimpleSlice_Transpile_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_S11_51_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_S11_51_SimpleSlice_TranspileJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.TranspileJoin(strings.ToUpper, ",")

		// Assert
		actual := args.Map{"result": result != "A,B"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'A,B', got ''", actual)
	})
}

func Test_S11_52_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_S11_52_SimpleSlice_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		hs := ss.Hashset()

		// Assert
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_S11_53_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_S11_53_SimpleSlice_Join", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().Join(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_54_SimpleSlice_JoinLine(t *testing.T) {
	safeTest(t, "Test_S11_54_SimpleSlice_JoinLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": ss.JoinLine() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_55_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_S11_55_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		actual := args.Map{"result": strings.HasSuffix(result, "\n")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected newline at end", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinLineEofLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_S11_56_SimpleSlice_JoinLineEofLine_AlreadyHas(t *testing.T) {
	safeTest(t, "Test_S11_56_SimpleSlice_JoinLineEofLine_AlreadyHas", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a\n"})

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S11_57_SimpleSlice_JoinSpace(t *testing.T) {
	safeTest(t, "Test_S11_57_SimpleSlice_JoinSpace", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.JoinSpace() != "a b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a b'", actual)
	})
}

func Test_S11_58_SimpleSlice_JoinComma(t *testing.T) {
	safeTest(t, "Test_S11_58_SimpleSlice_JoinComma", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.JoinComma() != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b'", actual)
	})
}

func Test_S11_59_SimpleSlice_JoinCsv(t *testing.T) {
	safeTest(t, "Test_S11_59_SimpleSlice_JoinCsv", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.JoinCsv()

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S11_60_SimpleSlice_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_S11_60_SimpleSlice_JoinCsvLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.JoinCsvLine() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_S11_61_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_S11_61_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a,b", "c,d"})

		// Act
		result := ss.EachItemSplitBy(",")

		// Assert
		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_S11_62_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_S11_62_SimpleSlice_PrependJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		result := ss.PrependJoin(",", "a")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_S11_63_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_S11_63_SimpleSlice_AppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.AppendJoin(",", "b")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_S11_64_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_S11_64_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})

		// Act
		ss.PrependAppend([]string{"a"}, []string{"c"})

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_S11_65_SimpleSlice_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_S11_65_SimpleSlice_PrependAppend_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})

		// Act
		ss.PrependAppend(nil, nil)

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
