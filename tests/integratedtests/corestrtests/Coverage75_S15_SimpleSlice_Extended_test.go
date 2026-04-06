package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": result.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": result.First() != "A"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'A', got ''", actual)
	})
}

func Test_C75_SimpleSlice_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Transpile_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_TranspileJoin(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_TranspileJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.TranspileJoin(strings.ToUpper, ",")

		// Assert
		actual := args.Map{"result": result != "A,B"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'A,B', got ''", actual)
	})
}

func Test_C75_SimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "a")

		// Act
		hs := ss.Hashset()

		// Assert
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 distinct", actual)
	})
}

func Test_C75_SimpleSlice_Join(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Join", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.Join(",")

		// Assert
		actual := args.Map{"result": result != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b,c', got ''", actual)
	})
}

func Test_C75_SimpleSlice_Join_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Join_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		result := ss.Join(",")

		// Assert
		actual := args.Map{"result": result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_JoinLine(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x", "y")

		// Act
		result := ss.JoinLine()

		// Assert
		actual := args.Map{"result": strings.Contains(result, "x") || !strings.Contains(result, "y")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected joined lines, got ''", actual)
	})
}

func Test_C75_SimpleSlice_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLine_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		actual := args.Map{"result": ss.JoinLine() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_JoinLineEofLine(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": strings.HasSuffix(result, "\n")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected trailing newline", actual)
	})
}

func Test_C75_SimpleSlice_JoinLineEofLine_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLineEofLine_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		actual := args.Map{"result": ss.JoinLineEofLine() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_JoinLineEofLine_AlreadyHasSuffix(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinLineEofLine_AlreadyHasSuffix", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b\n")

		// Act
		result := ss.JoinLineEofLine()

		// Assert — should not double-add newline
		actual := args.Map{"result": strings.HasSuffix(result, "\n\n")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not double newline", actual)
	})
}

func Test_C75_SimpleSlice_JoinSpace(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinSpace", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("hello", "world")

		// Act
		result := ss.JoinSpace()

		// Assert
		actual := args.Map{"result": result != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello world', got ''", actual)
	})
}

func Test_C75_SimpleSlice_JoinComma(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinComma", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinComma()

		// Assert
		actual := args.Map{"result": result != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_C75_SimpleSlice_JoinCsv(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsv", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinCsv()

		// Assert
		actual := args.Map{"result": strings.Contains(result, "\"a\"")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected quoted csv, got ''", actual)
	})
}

func Test_C75_SimpleSlice_JoinCsvLine(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsvLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinCsvLine()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C75_SimpleSlice_CsvStrings(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_CsvStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("hello", "world")

		// Act
		result := ss.CsvStrings()

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_CsvStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_CsvStrings_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		result := ss.CsvStrings()

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_JoinCsvString(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsvString", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinCsvString(";")

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C75_SimpleSlice_JoinCsvString_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinCsvString_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		actual := args.Map{"result": ss.JoinCsvString(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_JoinWith(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinWith", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.JoinWith(" - ")

		// Assert
		actual := args.Map{"result": result != " - a - b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected ' - a - b', got ''", actual)
	})
}

func Test_C75_SimpleSlice_JoinWith_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JoinWith_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		actual := args.Map{"result": ss.JoinWith(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a:b", "c:d")

		// Act
		result := ss.EachItemSplitBy(":")

		// Assert
		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_C75_SimpleSlice_PrependJoin(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "d")

		// Act
		result := ss.PrependJoin(",", "a", "b")

		// Assert
		actual := args.Map{"result": result != "a,b,c,d"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b,c,d', got ''", actual)
	})
}

func Test_C75_SimpleSlice_AppendJoin(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.AppendJoin(",", "c", "d")

		// Assert
		actual := args.Map{"result": result != "a,b,c,d"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b,c,d', got ''", actual)
	})
}

func Test_C75_SimpleSlice_PrependAppend(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b")

		// Act
		result := ss.PrependAppend([]string{"a"}, []string{"c"})

		// Assert
		actual := args.Map{"result": result.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual = args.Map{"result": result.First() != "a" || result.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected [a,b,c]", actual)
	})
}

func Test_C75_SimpleSlice_PrependAppend_EmptyPrepend(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependAppend_EmptyPrepend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.PrependAppend(nil, []string{"b"})

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_PrependAppend_EmptyAppend(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_PrependAppend_EmptyAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.PrependAppend([]string{"z"}, nil)

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_Collection(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Collection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		col := ss.Collection(false)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_ToCollection(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ToCollection", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x", "y")

		// Act
		col := ss.ToCollection(true)

		// Assert
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": nonPtr.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nonPtr mismatch", actual)
		actual = args.Map{"result": ptr.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ptr mismatch", actual)
	})
}

func Test_C75_SimpleSlice_String(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_String", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.String()

		// Assert
		actual := args.Map{"result": strings.Contains(result, "a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected to contain 'a'", actual)
	})
}

func Test_C75_SimpleSlice_String_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_String_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		actual := args.Map{"result": ss.String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
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
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNewStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.ConcatNewStrings("b", "c")

		// Assert
		actual := args.Map{"result": len(result) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C75_SimpleSlice_ConcatNewStrings_NilReceiver(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNewStrings_NilReceiver", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.ConcatNewStrings("a")

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_SimpleSlice_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_ConcatNew", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.ConcatNew("b")

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_Sort(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Sort", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")

		// Act
		ss.Sort()

		// Assert
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first 'a', got ''", actual)
	})
}

func Test_C75_SimpleSlice_Reverse_3Items(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_3Items", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		ss.Reverse()

		// Assert
		actual := args.Map{"result": ss.First() != "c" || ss.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_C75_SimpleSlice_Reverse_2Items(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_2Items", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		ss.Reverse()

		// Assert
		actual := args.Map{"result": ss.First() != "b" || ss.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_C75_SimpleSlice_Reverse_1Item(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_1Item", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.Reverse()

		// Assert
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_C75_SimpleSlice_Reverse_Empty(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Reverse_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.Reverse()

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C75_SimpleSlice_JsonModel(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JsonModel", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		model := ss.JsonModel()

		// Assert
		actual := args.Map{"result": len(model) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_JsonModelAny", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		modelAny := ss.JsonModelAny()

		// Assert
		actual := args.Map{"result": modelAny == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C75_SimpleSlice_MarshalUnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_MarshalUnmarshalJSON", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		bytes, err := ss.MarshalJSON()
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "marshal error:", actual)

		target := corestr.New.SimpleSlice.Empty()
		err = target.UnmarshalJSON(bytes)

		// Assert
		actual = args.Map{"result": err != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unmarshal error:", actual)
		actual = args.Map{"result": target.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error", actual)
		actual = args.Map{"result": jsonPtrResult.HasError()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "jsonPtr error", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual = args.Map{"result": result.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_SimpleSlice_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsonContractsBinder", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		actual := args.Map{"result": ss.AsJsonContractsBinder() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C75_SimpleSlice_AsJsoner(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsoner", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		actual := args.Map{"result": ss.AsJsoner() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
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
		actual := args.Map{"result": ptr == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil ptr", actual)
		actual = args.Map{"result": nonPtr.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

func Test_C75_SimpleSlice_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsonParseSelfInjector", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		actual := args.Map{"result": ss.AsJsonParseSelfInjector() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C75_SimpleSlice_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_AsJsonMarshaller", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		actual := args.Map{"result": ss.AsJsonMarshaller() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C75_SimpleSlice_Clear(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Clear", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.Clear()

		// Assert
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty after clear", actual)
	})
}

func Test_C75_SimpleSlice_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Clear_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.Clear()

		// Assert
		actual := args.Map{"result": result != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C75_SimpleSlice_Dispose(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_Dispose", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.Dispose()

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty after dispose", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual = args.Map{"result": len(bytes) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual = args.Map{"result": len(target) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_SimpleSlice_SafeStrings(t *testing.T) {
	safeTest(t, "Test_C75_SimpleSlice_SafeStrings", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		empty := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		actual := args.Map{"result": len(ss.SafeStrings()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(empty.SafeStrings()) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// --- newSimpleSliceCreator ---

func Test_C75_NewSimpleSlice_Cap(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Cap", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Cap(10)

		// Assert
		actual := args.Map{"result": ss == nil || ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty with capacity", actual)
	})
}

func Test_C75_NewSimpleSlice_Cap_Negative(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Cap_Negative", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Cap(-5)

		// Assert
		actual := args.Map{"result": ss == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C75_NewSimpleSlice_Default(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Default", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Default()

		// Assert
		actual := args.Map{"result": ss == nil || ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty default", actual)
	})
}

func Test_C75_NewSimpleSlice_Deserialize(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Deserialize", func() {
		// Arrange
		input := []byte(`["a","b"]`)

		// Act
		ss, err := corestr.New.SimpleSlice.Deserialize(input)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual = args.Map{"result": ss.Length() != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_Deserialize_Invalid(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Deserialize_Invalid", func() {
		// Arrange
		input := []byte(`not json`)

		// Act
		ss, err := corestr.New.SimpleSlice.Deserialize(input)

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
		actual = args.Map{"result": ss.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty on error", actual)
	})
}

func Test_C75_NewSimpleSlice_UsingLines_Clone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLines_Clone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLines(true, "a", "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_UsingLines_NoClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLines_NoClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLines(false, "x")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_NewSimpleSlice_UsingLines_Nil(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLines_Nil", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLines(false)

		// Assert — nil variadic returns empty
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C75_NewSimpleSlice_Lines(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Lines", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_Split(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Split", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Split("a:b:c", ":")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C75_NewSimpleSlice_SplitLines(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_SplitLines", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.SplitLines("a\nb\nc")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C75_NewSimpleSlice_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_SpreadStrings", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.SpreadStrings("a", "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_Hashset(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Hashset", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		ss := corestr.New.SimpleSlice.Hashset(hs)

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_Hashset_Empty(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Hashset_Empty", func() {
		// Arrange
		hs := corestr.New.Hashset.Empty()

		// Act
		ss := corestr.New.SimpleSlice.Hashset(hs)

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C75_NewSimpleSlice_Create(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Create", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Create([]string{"a"})

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsPtr(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsPtr", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsPtr([]string{"a"})

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsPtr_Empty", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsPtr([]string{})

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsOptions_Clone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsOptions_Clone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsOptions(true, []string{"a", "b"})

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsOptions_NoClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsOptions_NoClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsOptions(false, []string{"x"})

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsOptions_Empty(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsOptions_Empty", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsOptions(false, []string{})

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsClone([]string{"a", "b"})

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_StringsClone_Nil(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_StringsClone_Nil", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.StringsClone(nil)

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C75_NewSimpleSlice_Direct_Clone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Direct_Clone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Direct(true, []string{"a", "b"})

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C75_NewSimpleSlice_Direct_NoClone(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Direct_NoClone", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Direct(false, []string{"x"})

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C75_NewSimpleSlice_Direct_Nil(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_Direct_Nil", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.Direct(false, nil)

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C75_NewSimpleSlice_UsingSeparatorLine(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingSeparatorLine", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingSeparatorLine(":", "a:b:c")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C75_NewSimpleSlice_UsingLine(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_UsingLine", func() {
		// Arrange & Act
		ss := corestr.New.SimpleSlice.UsingLine("a\nb")

		// Assert
		actual := args.Map{"result": ss.Length() < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C75_NewSimpleSlice_ByLen(t *testing.T) {
	safeTest(t, "Test_C75_NewSimpleSlice_ByLen", func() {
		// Arrange
		input := []string{"a", "b", "c"}

		// Act
		ss := corestr.New.SimpleSlice.ByLen(input)

		// Assert
		actual := args.Map{"result": ss == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}
