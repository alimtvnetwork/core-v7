package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ========================================
// S15: SimpleSlice extended methods
//   Transpile, Join variants, Concat, CSV,
//   Sort, Reverse, JSON, Collection, PrependAppend
// ========================================

func Test_C75_SimpleSlice_Transpile(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Transpile", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		if result.Length() != 3 {
			t.Errorf("expected 3, got %d", result.Length())
		}
		if result.First() != "A" {
			t.Errorf("expected 'A', got '%s'", result.First())
		}
	})
}

func Test_C75_SimpleSlice_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Transpile_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		if result.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_TranspileJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.TranspileJoin(strings.ToUpper, ",")

		// Assert
		if result != "A,B" {
			t.Errorf("expected 'A,B', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "a")

		// Act
		hs := ss.Hashset()

		// Assert
		if hs.Length() != 2 {
			t.Errorf("expected 2 distinct, got %d", hs.Length())
		}
	})
}

func Test_C75_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Join", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.Join(",")

		// Assert
		if result != "a,b,c" {
			t.Errorf("expected 'a,b,c', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_Join_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Join_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		result := ss.Join(",")

		// Assert
		if result != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_JoinLine(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x", "y")

		// Act
		result := ss.JoinLine()

		// Assert
		if !strings.Contains(result, "x") || !strings.Contains(result, "y") {
			t.Errorf("expected joined lines, got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLine_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.JoinLine() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		if result == "" {
			t.Error("expected non-empty")
		}
		if !strings.HasSuffix(result, "\n") {
			t.Error("expected trailing newline")
		}
	})
}

func Test_C75_SimpleSlice_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLineEofLine_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.JoinLineEofLine() != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_JoinLineEofLine_AlreadyHasSuffix(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLineEofLine_AlreadyHasSuffix", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b\n")

		// Act
		result := ss.JoinLineEofLine()

		// Assert — should not double-add newline
		if strings.HasSuffix(result, "\n\n") {
			t.Error("should not double newline")
		}
	})
}

func Test_C75_SimpleSlice_JoinSpace(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinSpace", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("hello", "world")

		// Act
		result := ss.JoinSpace()

		// Assert
		if result != "hello world" {
			t.Errorf("expected 'hello world', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_JoinComma(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinComma", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinComma()

		// Assert
		if result != "a,b" {
			t.Errorf("expected 'a,b', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_JoinCsv(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsv", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinCsv()

		// Assert
		if !strings.Contains(result, "\"a\"") {
			t.Errorf("expected quoted csv, got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsvLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinCsvLine()

		// Assert
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C75_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_CsvStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("hello", "world")

		// Act
		result := ss.CsvStrings()

		// Assert
		if len(result) != 2 {
			t.Errorf("expected 2, got %d", len(result))
		}
	})
}

func Test_C75_SimpleSlice_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_CsvStrings_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		result := ss.CsvStrings()

		// Assert
		if len(result) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsvString", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinCsvString(";")

		// Assert
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C75_SimpleSlice_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsvString_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.JoinCsvString(",") != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_JoinWith(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinWith", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinWith(" - ")

		// Assert
		if result != " - a - b" {
			t.Errorf("expected ' - a - b', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinWith_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.JoinWith(",") != "" {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a:b", "c:d")

		// Act
		result := ss.EachItemSplitBy(":")

		// Assert
		if result.Length() != 4 {
			t.Errorf("expected 4, got %d", result.Length())
		}
	})
}

func Test_C75_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "d")

		// Act
		result := ss.PrependJoin(",", "a", "b")

		// Assert
		if result != "a,b,c,d" {
			t.Errorf("expected 'a,b,c,d', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.AppendJoin(",", "c", "d")

		// Assert
		if result != "a,b,c,d" {
			t.Errorf("expected 'a,b,c,d', got '%s'", result)
		}
	})
}

func Test_C75_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b")

		// Act
		result := ss.PrependAppend([]string{"a"}, []string{"c"})

		// Assert
		if result.Length() != 3 {
			t.Errorf("expected 3, got %d", result.Length())
		}
		if result.First() != "a" || result.Last() != "c" {
			t.Errorf("expected [a,b,c], got %v", result.Strings())
		}
	})
}

func Test_C75_SimpleSlice_PrependAppend_EmptyPrepend(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependAppend_EmptyPrepend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.PrependAppend(nil, []string{"b"})

		// Assert
		if ss.Length() != 2 {
			t.Errorf("expected 2, got %d", ss.Length())
		}
	})
}

func Test_C75_SimpleSlice_PrependAppend_EmptyAppend(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependAppend_EmptyAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.PrependAppend([]string{"z"}, nil)

		// Assert
		if ss.Length() != 2 {
			t.Errorf("expected 2, got %d", ss.Length())
		}
	})
}

func Test_C75_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Collection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		col := ss.Collection(false)

		// Assert
		if col.Length() != 2 {
			t.Errorf("expected 2, got %d", col.Length())
		}
	})
}

func Test_C75_SimpleSlice_ToCollection(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ToCollection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x", "y")

		// Act
		col := ss.ToCollection(true)

		// Assert
		if col.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_C75_SimpleSlice_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_NonPtr_Ptr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		nonPtr := ss.NonPtr()
		ptr := ss.Ptr()

		// Assert
		if nonPtr.Length() != 1 {
			t.Error("nonPtr mismatch")
		}
		if ptr.Length() != 1 {
			t.Error("ptr mismatch")
		}
	})
}

func Test_C75_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_String", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.String()

		// Assert
		if !strings.Contains(result, "a") {
			t.Error("expected to contain 'a'")
		}
	})
}

func Test_C75_SimpleSlice_String_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_String_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.String() != "" {
			t.Error("expected empty string")
		}
	})
}

func Test_C75_SimpleSlice_ConcatNewSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNewSimpleSlices", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a")
		ss2 := corestr.New.SimpleSlice.Lines("b")

		// Act
		result := ss1.ConcatNewSimpleSlices(ss2)

		// Assert
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C75_SimpleSlice_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNewStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.ConcatNewStrings("b", "c")

		// Assert
		if len(result) != 3 {
			t.Errorf("expected 3, got %d", len(result))
		}
	})
}

func Test_C75_SimpleSlice_ConcatNewStrings_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNewStrings_NilReceiver", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.ConcatNewStrings("a")

		// Assert
		if len(result) != 1 {
			t.Errorf("expected 1, got %d", len(result))
		}
	})
}

func Test_C75_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNew", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.ConcatNew("b")

		// Assert
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C75_SimpleSlice_Sort(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Sort", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")

		// Act
		ss.Sort()

		// Assert
		if ss.First() != "a" {
			t.Errorf("expected first 'a', got '%s'", ss.First())
		}
	})
}

func Test_C75_SimpleSlice_Reverse_3Items(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_3Items", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		ss.Reverse()

		// Assert
		if ss.First() != "c" || ss.Last() != "a" {
			t.Errorf("expected reversed, got %v", ss.Strings())
		}
	})
}

func Test_C75_SimpleSlice_Reverse_2Items(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_2Items", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		ss.Reverse()

		// Assert
		if ss.First() != "b" || ss.Last() != "a" {
			t.Errorf("expected reversed")
		}
	})
}

func Test_C75_SimpleSlice_Reverse_1Item(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_1Item", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.Reverse()

		// Assert
		if ss.First() != "a" {
			t.Error("expected same")
		}
	})
}

func Test_C75_SimpleSlice_Reverse_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.Reverse()

		// Assert
		if ss.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C75_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JsonModel", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		model := ss.JsonModel()

		// Assert
		if len(model) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C75_SimpleSlice_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JsonModelAny", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		modelAny := ss.JsonModelAny()

		// Assert
		if modelAny == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C75_SimpleSlice_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_MarshalUnmarshalJSON", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		bytes, err := ss.MarshalJSON()
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		target := corestr.New.SimpleSlice.Empty()
		err = target.UnmarshalJSON(bytes)

		// Assert
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}
		if target.Length() != 2 {
			t.Errorf("expected 2, got %d", target.Length())
		}
	})
}

func Test_C75_SimpleSlice_Json_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Json_JsonPtr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x")

		// Act
		jsonResult := ss.Json()
		jsonPtrResult := ss.JsonPtr()

		// Assert
		if jsonResult.HasError() {
			t.Error("json error")
		}
		if jsonPtrResult.HasError() {
			t.Error("jsonPtr error")
		}
	})
}

func Test_C75_SimpleSlice_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ParseInjectUsingJson", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		jsonResult := ss.JsonPtr()
		target := corestr.New.SimpleSlice.Empty()

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C75_SimpleSlice_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ParseInjectUsingJsonMust", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x")
		jsonResult := ss.JsonPtr()
		target := corestr.New.SimpleSlice.Empty()

		// Act
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Assert
		if result.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_C75_SimpleSlice_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsonContractsBinder", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.AsJsonContractsBinder() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C75_SimpleSlice_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsoner", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.AsJsoner() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C75_SimpleSlice_ToPtr_ToNonPtr(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ToPtr_ToNonPtr", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ptr := ss.ToPtr()
		nonPtr := ss.ToNonPtr()

		// Assert
		if ptr == nil {
			t.Error("expected non-nil ptr")
		}
		if nonPtr.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C75_SimpleSlice_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JsonParseSelfInject", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		jsonResult := ss.JsonPtr()
		target := corestr.New.SimpleSlice.Empty()

		// Act
		err := target.JsonParseSelfInject(jsonResult)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

func Test_C75_SimpleSlice_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsonParseSelfInjector", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.AsJsonParseSelfInjector() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C75_SimpleSlice_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsonMarshaller", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.AsJsonMarshaller() == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C75_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Clear", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.Clear()

		// Assert
		if result.Length() != 0 {
			t.Error("expected empty after clear")
		}
	})
}

func Test_C75_SimpleSlice_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Clear_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.Clear()

		// Assert
		if result != nil {
			t.Error("expected nil")
		}
	})
}

func Test_C75_SimpleSlice_Dispose(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Dispose", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.Dispose()

		// Assert
		if ss.Length() != 0 {
			t.Error("expected empty after dispose")
		}
	})
}

func Test_C75_SimpleSlice_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Dispose_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act — should not panic
		ss.Dispose()
	})
}

func Test_C75_SimpleSlice_Serialize(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Serialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		bytes, err := ss.Serialize()

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(bytes) == 0 {
			t.Error("expected non-empty bytes")
		}
	})
}

func Test_C75_SimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Deserialize", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		var target []string
		err := ss.Deserialize(&target)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if len(target) != 2 {
			t.Errorf("expected 2, got %d", len(target))
		}
	})
}

func Test_C75_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_SafeStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		empty := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if len(ss.SafeStrings()) != 1 {
			t.Error("expected 1")
		}
		if len(empty.SafeStrings()) != 0 {
			t.Error("expected 0")
		}
	})
}

// --- newSimpleSliceCreator ---

func Test_C75_NewSimpleSlice_Cap(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Cap", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Cap(10)

		// Assert
		if ss == nil || ss.Length() != 0 {
			t.Error("expected empty with capacity")
		}
	})
}

func Test_C75_NewSimpleSlice_Cap_Negative(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Cap_Negative", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Cap(-5)

		// Assert
		if ss == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C75_NewSimpleSlice_Default(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Default", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Default()

		// Assert
		if ss == nil || ss.Length() != 0 {
			t.Error("expected empty default")
		}
	})
}

func Test_C75_NewSimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Deserialize", func() {
		// Arrange
		input := []byte(`["a","b"]`)

		// Act
		ss, err := corestr.New.SimpleSlice.Deserialize(input)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if ss.Length() != 2 {
			t.Errorf("expected 2, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_Deserialize_Invalid(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Deserialize_Invalid", func() {
		// Arrange
		input := []byte(`not json`)

		// Act
		ss, err := corestr.New.SimpleSlice.Deserialize(input)

		// Assert
		if err == nil {
			t.Error("expected error")
		}
		if ss.Length() != 0 {
			t.Error("expected empty on error")
		}
	})
}

func Test_C75_NewSimpleSlice_UsingLines_Clone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLines_Clone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLines(true, "a", "b")

		// Assert
		if ss.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_C75_NewSimpleSlice_UsingLines_NoClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLines_NoClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLines(false, "x")

		// Assert
		if ss.Length() != 1 {
			t.Errorf("expected 1")
		}
	})
}

func Test_C75_NewSimpleSlice_UsingLines_Nil(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLines_Nil", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLines(false)

		// Assert — nil variadic returns empty
		if ss.Length() != 0 {
			t.Errorf("expected 0, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_Lines(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Lines", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Assert
		if ss.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C75_NewSimpleSlice_Split(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Split", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Split("a:b:c", ":")

		// Assert
		if ss.Length() != 3 {
			t.Errorf("expected 3, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_SplitLines(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_SplitLines", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.SplitLines("a\nb\nc")

		// Assert
		if ss.Length() != 3 {
			t.Errorf("expected 3, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_SpreadStrings", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.SpreadStrings("a", "b")

		// Assert
		if ss.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C75_NewSimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Hashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		ss := corestr.New.SimpleSlice.Hashset(hs)

		// Assert
		if ss.Length() != 2 {
			t.Errorf("expected 2, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_Hashset_Empty(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Hashset_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		ss := corestr.New.SimpleSlice.Hashset(hs)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C75_NewSimpleSlice_Create(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Create", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Create([]string{"a"})

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsPtr(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsPtr", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsPtr([]string{"a"})

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsPtr_Empty", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsPtr([]string{})

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsOptions_Clone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsOptions_Clone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsOptions(true, []string{"a", "b"})

		// Assert
		if ss.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsOptions_NoClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsOptions_NoClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsOptions(false, []string{"x"})

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsOptions_Empty(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsOptions_Empty", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsOptions(false, []string{})

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsClone([]string{"a", "b"})

		// Assert
		if ss.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C75_NewSimpleSlice_StringsClone_Nil(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsClone_Nil", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsClone(nil)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C75_NewSimpleSlice_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Direct_Clone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Direct(true, []string{"a", "b"})

		// Assert
		if ss.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C75_NewSimpleSlice_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Direct_NoClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Direct(false, []string{"x"})

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C75_NewSimpleSlice_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Direct_Nil", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Direct(false, nil)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C75_NewSimpleSlice_UsingSeparatorLine(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingSeparatorLine", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingSeparatorLine(":", "a:b:c")

		// Assert
		if ss.Length() != 3 {
			t.Errorf("expected 3, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_UsingLine(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLine", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLine("a\nb")

		// Assert
		if ss.Length() < 2 {
			t.Errorf("expected at least 2, got %d", ss.Length())
		}
	})
}

func Test_C75_NewSimpleSlice_ByLen(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_ByLen", func() {
		// Arrange
		input := []string{"a", "b", "c"}

		// Act
		ss := corestr.New.SimpleSlice.ByLen(input)

		// Assert
		if ss == nil {
			t.Error("expected non-nil")
		}
	})
}
