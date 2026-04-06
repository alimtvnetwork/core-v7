package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── CloneSimpleSliceToPointers ──

func Test_Cov4_CloneSimpleSliceToPointers(t *testing.T) {
	original := []string{"a", "b", "c"}
	cloned := stringslice.CloneSimpleSliceToPointers(original)
	actual := args.Map{
		"notNil": cloned != nil,
		"len":    len(*cloned),
	}
	expected := args.Map{
		"notNil": true,
		"len":    3,
	}
	expected.ShouldBeEqual(t, 0, "CloneSimpleSliceToPointers returns correct -- 3 items", actual)
}

// ── FirstOrDefaultPtr ──

func Test_Cov4_FirstOrDefaultPtr_HasItems(t *testing.T) {
	items := []string{"hello", "world"}
	result := stringslice.FirstOrDefaultPtr(items)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr returns first -- has items", actual)
}

func Test_Cov4_FirstOrDefaultPtr_Empty(t *testing.T) {
	result := stringslice.FirstOrDefaultPtr([]string{})
	actual := args.Map{"isEmpty": result == ""}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "FirstOrDefaultPtr returns empty -- empty", actual)
}

// ── LastOrDefaultPtr ──

func Test_Cov4_LastOrDefaultPtr_HasItems(t *testing.T) {
	items := []string{"hello", "world"}
	result := stringslice.LastOrDefaultPtr(items)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr returns last -- has items", actual)
}

func Test_Cov4_LastOrDefaultPtr_Empty(t *testing.T) {
	result := stringslice.LastOrDefaultPtr([]string{})
	actual := args.Map{"isEmpty": result == ""}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "LastOrDefaultPtr returns empty -- empty", actual)
}

// ── LastIndexPtr ──

func Test_Cov4_LastIndexPtr_HasItems(t *testing.T) {
	items := []string{"a", "b", "c"}
	result := stringslice.LastIndexPtr(items)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexPtr returns 2 -- 3 items", actual)
}

// ── LastSafeIndexPtr ──

func Test_Cov4_LastSafeIndexPtr_HasItems(t *testing.T) {
	items := []string{"a", "b"}
	result := stringslice.LastSafeIndexPtr(items)
	actual := args.Map{"val": result}
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "LastSafeIndexPtr returns 1 -- 2 items", actual)
}

// ── SafeIndexAtWithPtr ──

func Test_Cov4_SafeIndexAtWithPtr_Valid(t *testing.T) {
	items := []string{"a", "b", "c"}
	result := stringslice.SafeIndexAtWithPtr(items, 1, "")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns correct -- index 1", actual)
}

func Test_Cov4_SafeIndexAtWithPtr_OutOfRange(t *testing.T) {
	items := []string{"a"}
	result := stringslice.SafeIndexAtWithPtr(items, 5, "default")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "default"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtWithPtr returns default -- out of range", actual)
}

// ── SafeIndexAtUsingLastIndex ──

func Test_Cov4_SafeIndexAtUsingLastIndex_Valid(t *testing.T) {
	items := []string{"a", "b", "c"}
	result := stringslice.SafeIndexAtUsingLastIndex(items, 2, 2)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "c"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns correct -- last index", actual)
}

func Test_Cov4_SafeIndexAtUsingLastIndex_OutOfRange(t *testing.T) {
	items := []string{"a"}
	result := stringslice.SafeIndexAtUsingLastIndex(items, 5, 0)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "SafeIndexAtUsingLastIndex returns first -- lastIndex exceeds slice", actual)
}

// ── SafeRangeItemsPtr ──

func Test_Cov4_SafeRangeItemsPtr_Valid(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	result := stringslice.SafeRangeItemsPtr(items, 1, 3)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SafeRangeItemsPtr returns 2 items -- range 1 to 3", actual)
}

// ── NonWhitespacePtr ──

func Test_Cov4_NonWhitespacePtr_Mixed(t *testing.T) {
	items := []string{"hello", "  ", "world", ""}
	result := stringslice.NonWhitespacePtr(items)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespacePtr returns 2 -- skip whitespace and empty", actual)
}

// ── NonEmptyJoinPtr ──

func Test_Cov4_NonEmptyJoinPtr(t *testing.T) {
	items := []string{"hello", "", "world", ""}
	result := stringslice.NonEmptyJoinPtr(items, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello,world"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoinPtr joins non-empty -- comma sep", actual)
}

// ── NonWhitespaceJoinPtr ──

func Test_Cov4_NonWhitespaceJoinPtr(t *testing.T) {
	items := []string{"hello", "  ", "world"}
	result := stringslice.NonWhitespaceJoinPtr(items, ",")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello,world"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr joins non-whitespace -- comma sep", actual)
}

// ── AppendStringsWithAnyItems ──

func Test_Cov4_AppendStringsWithAnyItems(t *testing.T) {
	mainSlice := []any{"hello"}
	result := stringslice.AppendStringsWithAnyItems(false, false, mainSlice, "world", "!")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendStringsWithAnyItems returns 3 -- 1 + 2", actual)
}

// ── AppendAnyItemsWithStrings ──

func Test_Cov4_AppendAnyItemsWithStrings(t *testing.T) {
	mainSlice := []string{"hello"}
	result := stringslice.AppendAnyItemsWithStrings(false, false, mainSlice, 42, "world")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AppendAnyItemsWithStrings returns 3 -- 1 + 2", actual)
}

// ── LinesSimpleProcess ──

func Test_Cov4_LinesSimpleProcess(t *testing.T) {
	lines := []string{"hello", "world"}
	result := stringslice.LinesSimpleProcess(lines, func(lineIn string) string {
		return lineIn + "!"
	})
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "hello!"}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcess processes all -- 2 lines", actual)
}

// ── LinesSimpleProcessNoEmpty ──

func Test_Cov4_LinesSimpleProcessNoEmpty(t *testing.T) {
	lines := []string{"hello", "", "world", "   "}
	result := stringslice.LinesSimpleProcessNoEmpty(lines, func(lineIn string) string {
		return lineIn
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LinesSimpleProcessNoEmpty skips empty -- 3 non-empty returned", actual)
}

// ── TrimmedEachWordsPtr ──

func Test_Cov4_TrimmedEachWordsPtr(t *testing.T) {
	items := []string{"  hello  ", " world "}
	result := stringslice.TrimmedEachWordsPtr(items)
	actual := args.Map{"first": result[0], "last": result[1]}
	expected := args.Map{"first": "hello", "last": "world"}
	expected.ShouldBeEqual(t, 0, "TrimmedEachWordsPtr trims all -- 2 items", actual)
}

// ── SplitTrimmedNonEmptyAll ──

func Test_Cov4_SplitTrimmedNonEmptyAll(t *testing.T) {
	result := stringslice.SplitTrimmedNonEmptyAll("hello, , world", ",")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll returns 2 -- skip empty", actual)
}

// ── FirstLastDefaultPtr ──

func Test_Cov4_FirstLastDefaultPtr_HasItems(t *testing.T) {
	items := []string{"a", "b", "c"}
	first, last := stringslice.FirstLastDefaultPtr(items)
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr returns first and last -- 3 items", actual)
}

func Test_Cov4_FirstLastDefaultPtr_Empty(t *testing.T) {
	first, last := stringslice.FirstLastDefaultPtr([]string{})
	actual := args.Map{"firstEmpty": first == "", "lastEmpty": last == ""}
	expected := args.Map{"firstEmpty": true, "lastEmpty": true}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultPtr returns empty -- empty", actual)
}

// ── FirstLastDefaultStatusPtr ──

func Test_Cov4_FirstLastDefaultStatusPtr_HasItems(t *testing.T) {
	items := []string{"a", "b"}
	result := stringslice.FirstLastDefaultStatusPtr(items)
	actual := args.Map{
		"first":    result.First,
		"last":     result.Last,
		"hasFirst": result.HasFirst,
	}
	expected := args.Map{
		"first":    "a",
		"last":     "b",
		"hasFirst": true,
	}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr returns correct -- 2 items", actual)
}

func Test_Cov4_FirstLastDefaultStatusPtr_Empty(t *testing.T) {
	result := stringslice.FirstLastDefaultStatusPtr([]string{})
	actual := args.Map{"isValid": result.IsValid}
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "FirstLastDefaultStatusPtr returns invalid -- empty", actual)
}

// ── SafeIndexes ──

func Test_Cov4_SafeIndexes(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	result := stringslice.SafeIndexes(items, 0, 2, 99)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "SafeIndexes returns 3 -- preallocated with empty for out of range", actual)
}

// ── SafeIndexRanges ──

func Test_Cov4_SafeIndexRanges(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}
	result := stringslice.SafeIndexRanges(items, 1, 4)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns 4 -- range 1 to 4 inclusive", actual)
}
