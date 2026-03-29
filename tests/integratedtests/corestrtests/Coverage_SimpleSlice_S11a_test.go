package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_02_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_S11_02_SimpleSlice_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddSplit("a,b,c", ",")

		// Assert
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_04_SimpleSlice_Adds(t *testing.T) {
	safeTest(t, "Test_S11_04_SimpleSlice_Adds", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Adds("a", "b")

		// Assert
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_05_SimpleSlice_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_S11_05_SimpleSlice_Adds_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Adds()

		// Assert
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_06_SimpleSlice_Append(t *testing.T) {
	safeTest(t, "Test_S11_06_SimpleSlice_Append", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Append("a")

		// Assert
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_07_SimpleSlice_Append_Empty(t *testing.T) {
	safeTest(t, "Test_S11_07_SimpleSlice_Append_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Append()

		// Assert
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_08_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_S11_08_SimpleSlice_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("hello %s", "world")

		// Assert
		if ss.Length() != 1 || ss.First() != "hello world" {
			t.Fatal("expected 'hello world'")
		}
	})
}

func Test_S11_09_SimpleSlice_AppendFmt_EmptySkip(t *testing.T) {
	safeTest(t, "Test_S11_09_SimpleSlice_AppendFmt_EmptySkip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("")

		// Assert
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_11_SimpleSlice_AppendFmtIf_EmptyFormatSkip(t *testing.T) {
	safeTest(t, "Test_S11_11_SimpleSlice_AppendFmtIf_EmptyFormatSkip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmtIf(true, "")

		// Assert
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_12_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_S11_12_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsTitleValue("Name", "John")

		// Assert
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_13_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_S11_13_SimpleSlice_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsCurlyTitleWrap("Key", "Val")

		// Assert
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_16_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_S11_16_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})

		// Act
		ss.InsertAt(1, "b")

		// Assert
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1 — out of range skipped")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if ss.Length() != 2 {
			t.Fatal("expected 2")
		}
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
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_22_SimpleSlice_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_S11_22_SimpleSlice_AsDefaultError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"err1", "err2"})

		// Act
		err := ss.AsDefaultError()

		// Assert
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_S11_23_SimpleSlice_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_S11_23_SimpleSlice_AsError_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		err := ss.AsError(",")

		// Assert
		if err != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_S11_24_SimpleSlice_FirstAndLast(t *testing.T) {
	safeTest(t, "Test_S11_24_SimpleSlice_FirstAndLast", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		if ss.First() != "a" {
			t.Fatal("expected a")
		}
		if ss.Last() != "c" {
			t.Fatal("expected c")
		}
		if ss.FirstDynamic().(string) != "a" {
			t.Fatal("expected a dynamic")
		}
		if ss.LastDynamic().(string) != "c" {
			t.Fatal("expected c dynamic")
		}
	})
}

func Test_S11_25_SimpleSlice_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_S11_25_SimpleSlice_FirstOrDefault", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})

		// Act & Assert
		if ss.FirstOrDefault() != "" {
			t.Fatal("expected empty")
		}
		if ss.FirstOrDefaultDynamic().(string) != "" {
			t.Fatal("expected empty dynamic")
		}
		if ss2.FirstOrDefault() != "x" {
			t.Fatal("expected x")
		}
	})
}

func Test_S11_26_SimpleSlice_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_S11_26_SimpleSlice_LastOrDefault", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})

		// Act & Assert
		if ss.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
		if ss.LastOrDefaultDynamic().(string) != "" {
			t.Fatal("expected empty dynamic")
		}
		if ss2.LastOrDefault() != "x" {
			t.Fatal("expected x")
		}
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
		if len(skipped) != 2 {
			t.Fatal("expected 2")
		}
		if len(skippedAll) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_28_SimpleSlice_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_S11_28_SimpleSlice_SkipDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.SkipDynamic(1)

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
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
		if len(taken) != 2 {
			t.Fatal("expected 2")
		}
		if len(takenAll) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S11_30_SimpleSlice_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_S11_30_SimpleSlice_TakeDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.TakeDynamic(1)

		// Assert
		if result == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_S11_31_SimpleSlice_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_S11_31_SimpleSlice_LimitDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		if ss.LimitDynamic(1) == nil {
			t.Fatal("expected non-nil")
		}
		if len(ss.Limit(1)) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_32_SimpleSlice_Length_Count(t *testing.T) {
	safeTest(t, "Test_S11_32_SimpleSlice_Length_Count", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		if ss.Length() != 1 || ss.Count() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_33_SimpleSlice_Length_Nil(t *testing.T) {
	safeTest(t, "Test_S11_33_SimpleSlice_Length_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		if ss.Length() != 0 {
			t.Fatal("expected 0")
		}
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
		if count != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_35_SimpleSlice_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_S11_35_SimpleSlice_CountFunc_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		count := ss.CountFunc(func(index int, item string) bool { return true })

		// Assert
		if count != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_36_SimpleSlice_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S11_36_SimpleSlice_IsEmpty", func() {
		// Act & Assert
		if !corestr.Empty.SimpleSlice().IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_37_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_S11_37_SimpleSlice_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		if !ss.IsContains("a") {
			t.Fatal("expected true")
		}
		if ss.IsContains("z") {
			t.Fatal("expected false")
		}
		if corestr.Empty.SimpleSlice().IsContains("a") {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_S11_38_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_S11_38_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc"})

		// Act & Assert
		if !ss.IsContainsFunc("ab", strings.Contains) {
			t.Fatal("expected true")
		}
		if corestr.Empty.SimpleSlice().IsContainsFunc("a", strings.Contains) {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_S11_39_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_S11_39_SimpleSlice_IndexOf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		if ss.IndexOf("b") != 1 {
			t.Fatal("expected 1")
		}
		if ss.IndexOf("z") != -1 {
			t.Fatal("expected -1")
		}
		if corestr.Empty.SimpleSlice().IndexOf("a") != -1 {
			t.Fatal("expected -1 for empty")
		}
	})
}

func Test_S11_40_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_S11_40_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc", "def"})

		// Act
		idx := ss.IndexOfFunc("de", strings.Contains)

		// Assert
		if idx != 1 {
			t.Fatalf("expected 1, got %d", idx)
		}
		if corestr.Empty.SimpleSlice().IndexOfFunc("a", strings.Contains) != -1 {
			t.Fatal("expected -1 for empty")
		}
	})
}

func Test_S11_41_SimpleSlice_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_S11_41_SimpleSlice_HasAnyItem", func() {
		// Act & Assert
		if !corestr.New.SimpleSlice.Strings([]string{"a"}).HasAnyItem() {
			t.Fatal("expected true")
		}
	})
}

func Test_S11_42_SimpleSlice_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_S11_42_SimpleSlice_LastIndex_HasIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		if ss.LastIndex() != 1 {
			t.Fatal("expected 1")
		}
		if !ss.HasIndex(0) || !ss.HasIndex(1) {
			t.Fatal("expected true")
		}
		if ss.HasIndex(2) || ss.HasIndex(-1) {
			t.Fatal("expected false")
		}
	})
}

func Test_S11_43_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_S11_43_SimpleSlice_Strings_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		if len(ss.Strings()) != 1 || len(ss.List()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_44_SimpleSlice_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_S11_44_SimpleSlice_WrapDoubleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		if ss.WrapDoubleQuote().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_45_SimpleSlice_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_S11_45_SimpleSlice_WrapSingleQuote", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.WrapSingleQuote().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_46_SimpleSlice_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_S11_46_SimpleSlice_WrapTildaQuote", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.WrapTildaQuote().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_47_SimpleSlice_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_S11_47_SimpleSlice_WrapDoubleQuoteIfMissing", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.WrapDoubleQuoteIfMissing().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_48_SimpleSlice_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_S11_48_SimpleSlice_WrapSingleQuoteIfMissing", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.WrapSingleQuoteIfMissing().Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_S11_49_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_S11_49_SimpleSlice_Transpile", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		if result.First() != "A" {
			t.Fatal("expected A")
		}
	})
}

func Test_S11_50_SimpleSlice_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_S11_50_SimpleSlice_Transpile_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		if result.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_S11_51_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_S11_51_SimpleSlice_TranspileJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.TranspileJoin(strings.ToUpper, ",")

		// Assert
		if result != "A,B" {
			t.Fatalf("expected 'A,B', got '%s'", result)
		}
	})
}

func Test_S11_52_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_S11_52_SimpleSlice_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		hs := ss.Hashset()

		// Assert
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_S11_53_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_S11_53_SimpleSlice_Join", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		if ss.Join(",") != "a,b" {
			t.Fatal("expected a,b")
		}
		if corestr.Empty.SimpleSlice().Join(",") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_54_SimpleSlice_JoinLine(t *testing.T) {
	safeTest(t, "Test_S11_54_SimpleSlice_JoinLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		if ss.JoinLine() == "" {
			t.Fatal("expected non-empty")
		}
		if corestr.Empty.SimpleSlice().JoinLine() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_55_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_S11_55_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		if !strings.HasSuffix(result, "\n") {
			t.Fatal("expected newline at end")
		}
		if corestr.Empty.SimpleSlice().JoinLineEofLine() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_S11_56_SimpleSlice_JoinLineEofLine_AlreadyHas(t *testing.T) {
	safeTest(t, "Test_S11_56_SimpleSlice_JoinLineEofLine_AlreadyHas", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a\n"})

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S11_57_SimpleSlice_JoinSpace(t *testing.T) {
	safeTest(t, "Test_S11_57_SimpleSlice_JoinSpace", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if ss.JoinSpace() != "a b" {
			t.Fatal("expected 'a b'")
		}
	})
}

func Test_S11_58_SimpleSlice_JoinComma(t *testing.T) {
	safeTest(t, "Test_S11_58_SimpleSlice_JoinComma", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})
		if ss.JoinComma() != "a,b" {
			t.Fatal("expected 'a,b'")
		}
	})
}

func Test_S11_59_SimpleSlice_JoinCsv(t *testing.T) {
	safeTest(t, "Test_S11_59_SimpleSlice_JoinCsv", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.JoinCsv()
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S11_60_SimpleSlice_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_S11_60_SimpleSlice_JoinCsvLine", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		if ss.JoinCsvLine() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_S11_61_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_S11_61_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a,b", "c,d"})

		// Act
		result := ss.EachItemSplitBy(",")

		// Assert
		if result.Length() != 4 {
			t.Fatalf("expected 4, got %d", result.Length())
		}
	})
}

func Test_S11_62_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_S11_62_SimpleSlice_PrependJoin", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		result := ss.PrependJoin(",", "a")
		if result != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", result)
		}
	})
}

func Test_S11_63_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_S11_63_SimpleSlice_AppendJoin", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.AppendJoin(",", "b")
		if result != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", result)
		}
	})
}

func Test_S11_64_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_S11_64_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})

		// Act
		ss.PrependAppend([]string{"a"}, []string{"c"})

		// Assert
		if ss.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_S11_65_SimpleSlice_PrependAppend_Empty(t *testing.T) {
	safeTest(t, "Test_S11_65_SimpleSlice_PrependAppend_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})

		// Act
		ss.PrependAppend(nil, nil)

		// Assert
		if ss.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}
