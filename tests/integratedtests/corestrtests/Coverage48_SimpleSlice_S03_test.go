package corestrtests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// SimpleSlice — Segment 03: Uncovered methods
// ═══════════════════════════════════════════════════════════════

// --- Wrap methods ---

func Test_Cov48_SS_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_WrapDoubleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		// Act
		result := ss.WrapDoubleQuote()
		// Assert
		tc := caseV1Compat{Name: "WrapDoubleQuote", Expected: 2, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_WrapSingleQuote", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.WrapSingleQuote()
		tc := caseV1Compat{Name: "WrapSingleQuote", Expected: 1, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_WrapTildaQuote", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.WrapTildaQuote()
		tc := caseV1Compat{Name: "WrapTildaQuote", Expected: 1, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_WrapDoubleQuoteIfMissing", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.WrapDoubleQuoteIfMissing()
		tc := caseV1Compat{Name: "WrapDoubleQuoteIfMissing", Expected: 1, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_WrapSingleQuoteIfMissing", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.WrapSingleQuoteIfMissing()
		tc := caseV1Compat{Name: "WrapSingleQuoteIfMissing", Expected: 1, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Transpile ---

func Test_Cov48_SS_Transpile(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Transpile", func() {
		ss := corestr.New.SimpleSlice.Lines("abc", "def")
		result := ss.Transpile(func(s string) string { return s + "!" })
		tc := caseV1Compat{Name: "Transpile first", Expected: "abc!", Actual: result.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Transpile_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		result := ss.Transpile(func(s string) string { return s })
		tc := caseV1Compat{Name: "Transpile empty", Expected: true, Actual: result.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_TranspileJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.TranspileJoin(func(s string) string { return s + "!" }, ",")
		tc := caseV1Compat{Name: "TranspileJoin", Expected: "a!,b!", Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Hashset ---

func Test_Cov48_SS_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Hashset", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "a")
		hs := ss.Hashset()
		tc := caseV1Compat{Name: "Hashset length", Expected: 2, Actual: hs.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Join variants ---

func Test_Cov48_SS_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Join_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "Join empty", Expected: "", Actual: ss.Join(","), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinLine", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.JoinLine()
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov48_SS_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinLine_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "JoinLine empty", Expected: "", Actual: ss.JoinLine(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinLineEofLine", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.JoinLineEofLine()
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov48_SS_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinLineEofLine_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "JoinLineEofLine empty", Expected: "", Actual: ss.JoinLineEofLine(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinSpace(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinSpace", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "JoinSpace", Expected: "a b", Actual: ss.JoinSpace(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinComma(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinComma", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "JoinComma", Expected: "a,b", Actual: ss.JoinComma(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinCsv(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinCsv", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.JoinCsv()
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov48_SS_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinCsvLine", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.JoinCsvLine()
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov48_SS_JoinWith(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinWith", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.JoinWith("-")
		tc := caseV1Compat{Name: "JoinWith", Expected: "-a-b", Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinWith_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "JoinWith empty", Expected: "", Actual: ss.JoinWith("-"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinCsvString", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.JoinCsvString(",")
		if result == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_Cov48_SS_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JoinCsvString_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "JoinCsvString empty", Expected: "", Actual: ss.JoinCsvString(","), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- EachItemSplitBy ---

func Test_Cov48_SS_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_EachItemSplitBy", func() {
		ss := corestr.New.SimpleSlice.Lines("a.b", "c.d")
		result := ss.EachItemSplitBy(".")
		tc := caseV1Compat{Name: "EachItemSplitBy", Expected: 4, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- PrependJoin, AppendJoin, PrependAppend ---

func Test_Cov48_SS_PrependJoin(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_PrependJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "c")
		result := ss.PrependJoin(",", "a")
		tc := caseV1Compat{Name: "PrependJoin", Expected: "a,b,c", Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_AppendJoin(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AppendJoin", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.AppendJoin(",", "c")
		tc := caseV1Compat{Name: "AppendJoin", Expected: "a,b,c", Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_PrependAppend(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_PrependAppend", func() {
		ss := corestr.New.SimpleSlice.Lines("b")
		ss.PrependAppend([]string{"a"}, []string{"c"})
		tc := caseV1Compat{Name: "PrependAppend length", Expected: 3, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_PrependAppend_EmptyBoth(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_PrependAppend_EmptyBoth", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.PrependAppend(nil, nil)
		tc := caseV1Compat{Name: "PrependAppend empty", Expected: 1, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsEqual variants ---

func Test_Cov48_SS_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqual", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "IsEqual true", Expected: true, Actual: a.IsEqual(b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqual_DiffLength(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqual_DiffLength", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "IsEqual diff len", Expected: false, Actual: a.IsEqual(b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqual_BothNil", func() {
		var a *corestr.SimpleSlice
		var b *corestr.SimpleSlice
		tc := caseV1Compat{Name: "IsEqual both nil", Expected: true, Actual: a.IsEqual(b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqual_OneNil", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		var b *corestr.SimpleSlice
		tc := caseV1Compat{Name: "IsEqual one nil", Expected: false, Actual: a.IsEqual(b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqual_BothEmpty", func() {
		a := corestr.New.SimpleSlice.Cap(0)
		b := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IsEqual both empty", Expected: true, Actual: a.IsEqual(b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualLines", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "IsEqualLines true", Expected: true, Actual: ss.IsEqualLines([]string{"a", "b"}), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualLines_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualLines_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "IsEqualLines mismatch", Expected: false, Actual: ss.IsEqualLines([]string{"a", "c"}), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualUnorderedLines", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		tc := caseV1Compat{Name: "IsEqualUnorderedLines", Expected: true, Actual: ss.IsEqualUnorderedLines([]string{"a", "b"}), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualUnorderedLines_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualUnorderedLines_DiffLen", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "IsEqualUnorderedLines diff", Expected: false, Actual: ss.IsEqualUnorderedLines([]string{"a", "b"}), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualUnorderedLinesClone", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		tc := caseV1Compat{Name: "IsEqualUnorderedLinesClone", Expected: true, Actual: ss.IsEqualUnorderedLinesClone([]string{"a", "b"}), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualUnorderedLinesClone_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualUnorderedLinesClone_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "IsEqualUnorderedLinesClone mismatch", Expected: false, Actual: ss.IsEqualUnorderedLinesClone([]string{"x", "y"}), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsContainsFunc, IndexOfFunc, IndexOf ---

func Test_Cov48_SS_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsContainsFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("abc", "def")
		found := ss.IsContainsFunc("ab", func(item, search string) bool {
			return len(item) >= 2 && item[:2] == search
		})
		tc := caseV1Compat{Name: "IsContainsFunc", Expected: true, Actual: found, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsContainsFunc_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IsContainsFunc empty", Expected: false, Actual: ss.IsContainsFunc("x", func(a, b string) bool { return a == b }), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IndexOfFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("abc", "def")
		idx := ss.IndexOfFunc("def", func(item, search string) bool { return item == search })
		tc := caseV1Compat{Name: "IndexOfFunc", Expected: 1, Actual: idx, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IndexOfFunc_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IndexOfFunc_NotFound", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		idx := ss.IndexOfFunc("z", func(item, search string) bool { return item == search })
		tc := caseV1Compat{Name: "IndexOfFunc not found", Expected: -1, Actual: idx, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IndexOfFunc_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IndexOfFunc empty", Expected: -1, Actual: ss.IndexOfFunc("x", func(a, b string) bool { return a == b }), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IndexOf(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IndexOf", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		tc := caseV1Compat{Name: "IndexOf", Expected: 1, Actual: ss.IndexOf("b"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IndexOf_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IndexOf_NotFound", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "IndexOf not found", Expected: -1, Actual: ss.IndexOf("z"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- HasIndex, HasAnyItem ---

func Test_Cov48_SS_HasIndex(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_HasIndex", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "HasIndex true", Expected: true, Actual: ss.HasIndex(1), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_HasIndex_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_HasIndex_OutOfRange", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "HasIndex oob", Expected: false, Actual: ss.HasIndex(5), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_HasIndex_Negative(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_HasIndex_Negative", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "HasIndex neg", Expected: false, Actual: ss.HasIndex(-1), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_HasAnyItem", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "HasAnyItem", Expected: true, Actual: ss.HasAnyItem(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_HasAnyItem_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_HasAnyItem_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "HasAnyItem empty", Expected: false, Actual: ss.HasAnyItem(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Strings, NonPtr, Ptr, ToPtr, ToNonPtr ---

func Test_Cov48_SS_Strings(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Strings", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "Strings", Expected: 1, Actual: len(ss.Strings()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_NonPtr", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		np := ss.NonPtr()
		tc := caseV1Compat{Name: "NonPtr length", Expected: 1, Actual: np.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Ptr", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		p := ss.Ptr()
		if p == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_ToPtr(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ToPtr", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		p := ss.ToPtr()
		if p == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ToNonPtr", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		np := ss.ToNonPtr()
		tc := caseV1Compat{Name: "ToNonPtr", Expected: 1, Actual: np.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- String (empty) ---

func Test_Cov48_SS_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_String_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "String empty", Expected: "", Actual: ss.String(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- ConcatNew variants ---

func Test_Cov48_SS_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ConcatNew", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.ConcatNew("b", "c")
		tc := caseV1Compat{Name: "ConcatNew", Expected: 3, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ConcatNewStrings", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.ConcatNewStrings("b")
		tc := caseV1Compat{Name: "ConcatNewStrings", Expected: 2, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ConcatNewSimpleSlices", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		b := corestr.New.SimpleSlice.Lines("b")
		result := a.ConcatNewSimpleSlices(b)
		tc := caseV1Compat{Name: "ConcatNewSimpleSlices", Expected: 2, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Collection, ToCollection ---

func Test_Cov48_SS_Collection(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Collection", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		c := ss.Collection(false)
		tc := caseV1Compat{Name: "Collection", Expected: 2, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_ToCollection(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ToCollection", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		c := ss.ToCollection(true)
		tc := caseV1Compat{Name: "ToCollection", Expected: 1, Actual: c.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- CsvStrings ---

func Test_Cov48_SS_CsvStrings(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_CsvStrings", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.CsvStrings()
		tc := caseV1Compat{Name: "CsvStrings", Expected: 1, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_CsvStrings_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		result := ss.CsvStrings()
		tc := caseV1Compat{Name: "CsvStrings empty", Expected: 0, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Sort, Reverse ---

func Test_Cov48_SS_Sort(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Sort", func() {
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")
		ss.Sort()
		tc := caseV1Compat{Name: "Sort first", Expected: "a", Actual: ss.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Reverse(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Reverse", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		ss.Reverse()
		tc := caseV1Compat{Name: "Reverse first", Expected: "c", Actual: ss.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Reverse_Two", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		ss.Reverse()
		tc := caseV1Compat{Name: "Reverse two", Expected: "b", Actual: ss.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Reverse_Single(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Reverse_Single", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Reverse()
		tc := caseV1Compat{Name: "Reverse single", Expected: "a", Actual: ss.First(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- JsonModel, JsonModelAny, AsJson* ---

func Test_Cov48_SS_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JsonModel", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		m := ss.JsonModel()
		tc := caseV1Compat{Name: "JsonModel", Expected: 1, Actual: len(m), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JsonModelAny", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		a := ss.JsonModelAny()
		if a == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AsJsonContractsBinder", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		b := ss.AsJsonContractsBinder()
		if b == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AsJsoner", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		j := ss.AsJsoner()
		if j == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AsJsonParseSelfInjector", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		inj := ss.AsJsonParseSelfInjector()
		if inj == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AsJsonMarshaller", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		m := ss.AsJsonMarshaller()
		if m == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_Cov48_SS_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_JsonParseSelfInject", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Cap(0)
		err := ss2.JsonParseSelfInject(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_Cov48_SS_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ParseInjectUsingJsonMust", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		jr := ss.JsonPtr()
		ss2 := corestr.New.SimpleSlice.Cap(0)
		result := ss2.ParseInjectUsingJsonMust(jr)
		tc := caseV1Compat{Name: "ParseInjectUsingJsonMust", Expected: 1, Actual: result.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- MarshalJSON / UnmarshalJSON invalid ---

func Test_Cov48_SS_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_UnmarshalJSON_Invalid", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		err := json.Unmarshal([]byte("invalid"), ss)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

// --- Clear, Dispose ---

func Test_Cov48_SS_Clear(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Clear", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		ss.Clear()
		tc := caseV1Compat{Name: "Clear", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Clear_Nil", func() {
		var ss *corestr.SimpleSlice
		result := ss.Clear()
		if result != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_Cov48_SS_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Dispose", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Dispose()
		tc := caseV1Compat{Name: "Dispose", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}
func Test_Cov48_SS_Clone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Clone", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		cloned := ss.Clone(true)
		tc := caseV1Compat{Name: "Clone length", Expected: 2, Actual: cloned.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}
	safeTest(t, "Test_Cov48_SS_DeepClone", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		cloned := ss.DeepClone()
		tc := caseV1Compat{Name: "DeepClone", Expected: 1, Actual: cloned.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_ShadowClone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_ShadowClone", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		cloned := ss.ShadowClone()
		tc := caseV1Compat{Name: "ShadowClone", Expected: 1, Actual: cloned.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsDistinctEqual, IsDistinctEqualRaw ---

func Test_Cov48_SS_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsDistinctEqualRaw", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "a")
		tc := caseV1Compat{Name: "IsDistinctEqualRaw", Expected: true, Actual: ss.IsDistinctEqualRaw("b", "a"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsDistinctEqual", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("b", "a")
		tc := caseV1Compat{Name: "IsDistinctEqual", Expected: true, Actual: a.IsDistinctEqual(b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsUnorderedEqual, IsUnorderedEqualRaw ---

func Test_Cov48_SS_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqualRaw_Clone", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		tc := caseV1Compat{Name: "IsUnorderedEqualRaw clone", Expected: true, Actual: ss.IsUnorderedEqualRaw(true, "a", "b"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqualRaw_NoClone", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		tc := caseV1Compat{Name: "IsUnorderedEqualRaw no clone", Expected: true, Actual: ss.IsUnorderedEqualRaw(false, "a", "b"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsUnorderedEqualRaw_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqualRaw_DiffLen", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "IsUnorderedEqualRaw diff len", Expected: false, Actual: ss.IsUnorderedEqualRaw(false, "a", "b"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsUnorderedEqualRaw_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqualRaw_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IsUnorderedEqualRaw empty", Expected: true, Actual: ss.IsUnorderedEqualRaw(false), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsUnorderedEqual(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqual", func() {
		a := corestr.New.SimpleSlice.Lines("b", "a")
		b := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "IsUnorderedEqual", Expected: true, Actual: a.IsUnorderedEqual(true, b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqual_BothEmpty", func() {
		a := corestr.New.SimpleSlice.Cap(0)
		b := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IsUnorderedEqual empty", Expected: true, Actual: a.IsUnorderedEqual(false, b), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsUnorderedEqual_NilRight(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsUnorderedEqual_NilRight", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "IsUnorderedEqual nil right", Expected: false, Actual: a.IsUnorderedEqual(false, nil), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsEqualByFunc ---

func Test_Cov48_SS_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFunc", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")
		tc := caseV1Compat{Name: "IsEqualByFunc", Expected: true, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFunc_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "x")
		tc := caseV1Compat{Name: "IsEqualByFunc mismatch", Expected: false, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualByFunc_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFunc_DiffLen", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r }, "a", "b")
		tc := caseV1Compat{Name: "IsEqualByFunc diff len", Expected: false, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualByFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFunc_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		result := ss.IsEqualByFunc(func(i int, l, r string) bool { return l == r })
		tc := caseV1Compat{Name: "IsEqualByFunc empty", Expected: true, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsEqualByFuncLinesSplit ---

func Test_Cov48_SS_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFuncLinesSplit", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })
		tc := caseV1Compat{Name: "IsEqualByFuncLinesSplit", Expected: true, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualByFuncLinesSplit_Trim(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFuncLinesSplit_Trim", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.IsEqualByFuncLinesSplit(true, ",", " a , b ", func(i int, l, r string) bool { return l == r })
		tc := caseV1Compat{Name: "IsEqualByFuncLinesSplit trim", Expected: true, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualByFuncLinesSplit_DiffLen(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFuncLinesSplit_DiffLen", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return l == r })
		tc := caseV1Compat{Name: "IsEqualByFuncLinesSplit diff len", Expected: false, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEqualByFuncLinesSplit_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEqualByFuncLinesSplit_Mismatch", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,x", func(i int, l, r string) bool { return l == r })
		tc := caseV1Compat{Name: "IsEqualByFuncLinesSplit mismatch", Expected: false, Actual: result, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- DistinctDiffRaw, DistinctDiff ---

func Test_Cov48_SS_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiffRaw", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result := ss.DistinctDiffRaw("b", "c")
		if len(result) == 0 {
			t.Fatal("expected diff items")
		}
	})
}

func Test_Cov48_SS_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiffRaw_BothNil", func() {
		var ss *corestr.SimpleSlice
		result := ss.DistinctDiffRaw()
		tc := caseV1Compat{Name: "DistinctDiffRaw nil", Expected: 0, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_DistinctDiffRaw_NilLeft(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiffRaw_NilLeft", func() {
		var ss *corestr.SimpleSlice
		result := ss.DistinctDiffRaw("a")
		tc := caseV1Compat{Name: "DistinctDiffRaw nil left", Expected: 1, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_DistinctDiffRaw_NilRight(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiffRaw_NilRight", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		result := ss.DistinctDiffRaw()
		tc := caseV1Compat{Name: "DistinctDiffRaw nil right", Expected: 1, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiff", func() {
		a := corestr.New.SimpleSlice.Lines("a", "b")
		b := corestr.New.SimpleSlice.Lines("b", "c")
		result := a.DistinctDiff(b)
		if len(result) == 0 {
			t.Fatal("expected diff items")
		}
	})
}

func Test_Cov48_SS_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiff_BothNil", func() {
		var a *corestr.SimpleSlice
		var b *corestr.SimpleSlice
		result := a.DistinctDiff(b)
		tc := caseV1Compat{Name: "DistinctDiff nil", Expected: 0, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_DistinctDiff_NilLeft(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiff_NilLeft", func() {
		var a *corestr.SimpleSlice
		b := corestr.New.SimpleSlice.Lines("x")
		result := a.DistinctDiff(b)
		tc := caseV1Compat{Name: "DistinctDiff nil left", Expected: 1, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_DistinctDiff_NilRight(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_DistinctDiff_NilRight", func() {
		a := corestr.New.SimpleSlice.Lines("a")
		result := a.DistinctDiff(nil)
		tc := caseV1Compat{Name: "DistinctDiff nil right", Expected: 1, Actual: len(result), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- AddedRemovedLinesDiff ---

func Test_Cov48_SS_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AddedRemovedLinesDiff", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		added, removed := ss.AddedRemovedLinesDiff("b", "c")
		if len(added) == 0 {
			t.Fatal("expected added items")
		}
		if len(removed) == 0 {
			t.Fatal("expected removed items")
		}
	})
}

func Test_Cov48_SS_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_AddedRemovedLinesDiff_BothNil", func() {
		var ss *corestr.SimpleSlice
		added, removed := ss.AddedRemovedLinesDiff()
		if added != nil || removed != nil {
			t.Fatal("expected nil")
		}
	})
}

// --- RemoveIndexes edge cases ---

func Test_Cov48_SS_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_RemoveIndexes_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		_, err := ss.RemoveIndexes(0)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_Cov48_SS_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_RemoveIndexes_InvalidIndex", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		result, err := ss.RemoveIndexes(99)
		if err == nil {
			t.Fatal("expected error for invalid index")
		}
		if result.Length() != 2 {
			t.Fatal("expected 2 items still")
		}
	})
}

// --- Deserialize ---

func Test_Cov48_SS_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Deserialize", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		var target []string
		err := ss.Deserialize(&target)
		if err != nil {
			t.Fatal("unexpected error")
		}
		tc := caseV1Compat{Name: "Deserialize", Expected: 2, Actual: len(target), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- CountFunc empty ---

func Test_Cov48_SS_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_CountFunc_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		c := ss.CountFunc(func(i int, s string) bool { return true })
		tc := caseV1Compat{Name: "CountFunc empty", Expected: 0, Actual: c, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsContains empty ---

func Test_Cov48_SS_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsContains_Empty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IsContains empty", Expected: false, Actual: ss.IsContains("x"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- Length nil ---

func Test_Cov48_SS_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Length_Nil", func() {
		var ss *corestr.SimpleSlice
		tc := caseV1Compat{Name: "Length nil", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- IsEmpty ---

func Test_Cov48_SS_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEmpty", func() {
		ss := corestr.New.SimpleSlice.Cap(0)
		tc := caseV1Compat{Name: "IsEmpty true", Expected: true, Actual: ss.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_IsEmpty_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_IsEmpty_NonEmpty", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		tc := caseV1Compat{Name: "IsEmpty false", Expected: false, Actual: ss.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// --- newSimpleSliceCreator coverage ---

func Test_Cov48_SS_Creator_Cap(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Cap", func() {
		ss := corestr.New.SimpleSlice.Cap(5)
		tc := caseV1Compat{Name: "Creator Cap", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Cap_Negative(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Cap_Negative", func() {
		ss := corestr.New.SimpleSlice.Cap(-1)
		tc := caseV1Compat{Name: "Creator Cap neg", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Default(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Default", func() {
		ss := corestr.New.SimpleSlice.Default()
		tc := caseV1Compat{Name: "Creator Default", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Empty", func() {
		ss := corestr.New.SimpleSlice.Empty()
		tc := caseV1Compat{Name: "Creator Empty", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Lines(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Lines", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		tc := caseV1Compat{Name: "Creator Lines", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_SpreadStrings", func() {
		ss := corestr.New.SimpleSlice.SpreadStrings("a", "b")
		tc := caseV1Compat{Name: "Creator SpreadStrings", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Split(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Split", func() {
		ss := corestr.New.SimpleSlice.Split("a,b,c", ",")
		tc := caseV1Compat{Name: "Creator Split", Expected: 3, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_SplitLines(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_SplitLines", func() {
		ss := corestr.New.SimpleSlice.SplitLines("a\nb")
		tc := caseV1Compat{Name: "Creator SplitLines", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Create(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Create", func() {
		ss := corestr.New.SimpleSlice.Create([]string{"a"})
		tc := caseV1Compat{Name: "Creator Create", Expected: 1, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_StringsClone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_StringsClone", func() {
		original := []string{"a", "b"}
		ss := corestr.New.SimpleSlice.StringsClone(original)
		tc := caseV1Compat{Name: "Creator StringsClone", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_StringsClone_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_StringsClone_Nil", func() {
		ss := corestr.New.SimpleSlice.StringsClone(nil)
		tc := caseV1Compat{Name: "Creator StringsClone nil", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Direct_Clone", func() {
		ss := corestr.New.SimpleSlice.Direct(true, []string{"a"})
		tc := caseV1Compat{Name: "Creator Direct clone", Expected: 1, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Direct_NoClone", func() {
		ss := corestr.New.SimpleSlice.Direct(false, []string{"a"})
		tc := caseV1Compat{Name: "Creator Direct no clone", Expected: 1, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Direct_Nil", func() {
		ss := corestr.New.SimpleSlice.Direct(true, nil)
		tc := caseV1Compat{Name: "Creator Direct nil", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_UsingLines_Clone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_UsingLines_Clone", func() {
		ss := corestr.New.SimpleSlice.UsingLines(true, "a", "b")
		tc := caseV1Compat{Name: "Creator UsingLines clone", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_UsingLines_NoClone(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_UsingLines_NoClone", func() {
		ss := corestr.New.SimpleSlice.UsingLines(false, "a")
		tc := caseV1Compat{Name: "Creator UsingLines no clone", Expected: 1, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_UsingLine(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_UsingLine", func() {
		ss := corestr.New.SimpleSlice.UsingLine("a\nb")
		if ss.Length() < 1 {
			t.Fatal("expected at least 1")
		}
	})
}

func Test_Cov48_SS_Creator_UsingSeparatorLine(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_UsingSeparatorLine", func() {
		ss := corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b,c")
		tc := caseV1Compat{Name: "Creator UsingSeparatorLine", Expected: 3, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Hashset(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Hashset", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		ss := corestr.New.SimpleSlice.Hashset(hs)
		tc := caseV1Compat{Name: "Creator Hashset", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Hashset_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Hashset_Empty", func() {
		hs := corestr.New.Hashset.Empty()
		ss := corestr.New.SimpleSlice.Hashset(hs)
		tc := caseV1Compat{Name: "Creator Hashset empty", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Map(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Map", func() {
		m := map[string]int{"a": 1, "b": 2}
		ss := corestr.New.SimpleSlice.Map(m)
		tc := caseV1Compat{Name: "Creator Map", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_ByLen(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_ByLen", func() {
		source := []string{"a", "b", "c"}
		ss := corestr.New.SimpleSlice.ByLen(source)
		tc := caseV1Compat{Name: "Creator ByLen", Expected: 0, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Deserialize", func() {
		data, _ := json.Marshal([]string{"a", "b"})
		ss, err := corestr.New.SimpleSlice.Deserialize(data)
		if err != nil {
			t.Fatal("unexpected error")
		}
		tc := caseV1Compat{Name: "Creator Deserialize", Expected: 2, Actual: ss.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov48_SS_Creator_Deserialize_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov48_SS_Creator_Deserialize_Invalid", func() {
		_, err := corestr.New.SimpleSlice.Deserialize([]byte("invalid"))
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
