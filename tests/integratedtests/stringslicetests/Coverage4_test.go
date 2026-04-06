package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coretests/args"
)
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
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "SafeIndexRanges returns 4 -- range 1 to 4 inclusive", actual)
}
