package stringslicetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage14 — stringslice final coverage gaps
// ══════════════════════════════════════════════════════════════════════════════

// --- ProcessAsync ---

func Test_ProcessAsync_Valid(t *testing.T) {
	// Arrange
	items := []any{"a", "b", "c"}

	// Act
	result := stringslice.ProcessAsync(func(index int, item any) string {
		return item.(string) + "!"
	}, items...)

	// Assert
	convey.Convey("ProcessAsync processes all items", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
		convey.So(result[0], convey.ShouldEqual, "a!")
	})
}

func Test_ProcessAsync_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.ProcessAsync(func(index int, item any) string {
		return ""
	})

	// Assert
	convey.Convey("ProcessAsync empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- ProcessOptionAsync ---

func Test_ProcessOptionAsync_SkipOnNil(t *testing.T) {
	// Arrange
	items := []any{"a", "b", ""}

	// Act
	result := stringslice.ProcessOptionAsync(
		true,
		func(index int, item any) string {
			return item.(string)
		},
		items...,
	)

	// Assert
	convey.Convey("ProcessOptionAsync skips empty strings", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_ProcessOptionAsync_NoSkip_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	items := []any{"a", "", "c"}

	// Act
	result := stringslice.ProcessOptionAsync(
		false,
		func(index int, item any) string {
			return item.(string)
		},
		items...,
	)

	// Assert
	convey.Convey("ProcessOptionAsync no skip returns all", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_ProcessOptionAsync_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.ProcessOptionAsync(
		true,
		func(index int, item any) string { return "" },
	)

	// Assert
	convey.Convey("ProcessOptionAsync empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- LinesAsyncProcess ---

func Test_LinesAsyncProcess_Valid(t *testing.T) {
	// Arrange
	lines := []string{"hello", "world"}

	// Act
	result := stringslice.LinesAsyncProcess(lines, func(index int, lineIn string) string {
		return lineIn + "!"
	})

	// Assert
	convey.Convey("LinesAsyncProcess processes all lines", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
		convey.So(result[0], convey.ShouldEqual, "hello!")
	})
}

func Test_LinesAsyncProcess_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.LinesAsyncProcess([]string{}, func(index int, lineIn string) string {
		return lineIn
	})

	// Assert
	convey.Convey("LinesAsyncProcess empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- AnyLinesProcessAsyncUsingProcessor ---

func Test_AnyLinesProcessAsync_Valid(t *testing.T) {
	// Arrange
	lines := []int{1, 2, 3}

	// Act
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		lines,
		func(index int, lineIn any) string {
			return "x"
		},
	)

	// Assert
	convey.Convey("AnyLinesProcessAsync processes all", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_AnyLinesProcessAsync_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		nil,
		func(index int, lineIn any) string { return "" },
	)

	// Assert
	convey.Convey("AnyLinesProcessAsync nil returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_AnyLinesProcessAsync_NotSlice_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		"not-a-slice",
		func(index int, lineIn any) string { return "" },
	)

	// Assert
	convey.Convey("AnyLinesProcessAsync non-slice returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_AnyLinesProcessAsync_EmptySlice_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.AnyLinesProcessAsyncUsingProcessor(
		[]string{},
		func(index int, lineIn any) string { return "" },
	)

	// Assert
	convey.Convey("AnyLinesProcessAsync empty slice returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- ExpandByFunc ---

func Test_ExpandByFunc_Valid(t *testing.T) {
	// Arrange
	slice := []string{"a-b", "c-d"}

	// Act
	result := stringslice.ExpandByFunc(slice, func(line string) []string {
		return []string{line, line + "!"}
	})

	// Assert
	convey.Convey("ExpandByFunc expands items", t, func() {
		convey.So(len(result), convey.ShouldEqual, 4)
	})
}

func Test_ExpandByFunc_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.ExpandByFunc([]string{}, func(line string) []string {
		return []string{line}
	})

	// Assert
	convey.Convey("ExpandByFunc empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_ExpandByFunc_ExpandReturnsEmpty(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.ExpandByFunc(slice, func(line string) []string {
		return []string{}
	})

	// Assert
	convey.Convey("ExpandByFunc with empty expand returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- ExpandBySplit ---

func Test_ExpandBySplit_Valid(t *testing.T) {
	// Arrange
	slice := []string{"a-b", "c-d"}

	// Act
	result := stringslice.ExpandBySplit(slice, "-")

	// Assert
	convey.Convey("ExpandBySplit splits and expands", t, func() {
		convey.So(len(result), convey.ShouldEqual, 4)
	})
}

func Test_ExpandBySplit_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.ExpandBySplit([]string{}, "-")

	// Assert
	convey.Convey("ExpandBySplit empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- ExpandBySplits ---

func Test_ExpandBySplits_MultiSplitters(t *testing.T) {
	// Arrange
	slice := []string{"a-b.c"}

	// Act
	result := stringslice.ExpandBySplits(slice, "-", ".")

	// Assert
	convey.Convey("ExpandBySplits multi-splitter expands", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 1)
	})
}

func Test_ExpandBySplits_NoSplitters_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a-b"}

	// Act
	result := stringslice.ExpandBySplits(slice)

	// Assert
	convey.Convey("ExpandBySplits no splitters returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- SafeIndexRanges ---

func Test_SafeIndexRanges_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c", "d", "e"}

	// Act
	result := stringslice.SafeIndexRanges(slice, 1, 3)

	// Assert
	convey.Convey("SafeIndexRanges returns correct range", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
		convey.So(result[0], convey.ShouldEqual, "b")
	})
}

func Test_SafeIndexRanges_NegativeRange_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexRanges(slice, 3, 1) // start > end

	// Assert
	convey.Convey("SafeIndexRanges negative range returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_SafeIndexRanges_OutOfBounds(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexRanges(slice, 0, 5)

	// Assert
	convey.Convey("SafeIndexRanges out of bounds fills empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 6)
	})
}

// --- SafeIndexes ---

func Test_SafeIndexes_OutOfBounds(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexes(slice, 0, 5, -1)

	// Assert
	convey.Convey("SafeIndexes skips out of bounds", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
		convey.So(result[0], convey.ShouldEqual, "a")
		convey.So(result[1], convey.ShouldBeEmpty) // 5 is out
	})
}

// --- SafeIndexesDefaultWithDetail ---

func Test_SafeIndexesDefaultWithDetail_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	result := stringslice.SafeIndexesDefaultWithDetail(slice, 0, 2)

	// Assert
	convey.Convey("SafeIndexesDefaultWithDetail returns valid", t, func() {
		convey.So(result.IsValid, convey.ShouldBeTrue)
		convey.So(len(result.Values), convey.ShouldEqual, 2)
	})
}

func Test_SafeIndexesDefaultWithDetail_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SafeIndexesDefaultWithDetail([]string{}, 0, 1)

	// Assert
	convey.Convey("SafeIndexesDefaultWithDetail empty returns invalid", t, func() {
		convey.So(result.IsValid, convey.ShouldBeFalse)
	})
}

func Test_SafeIndexesDefaultWithDetail_OutOfBounds(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexesDefaultWithDetail(slice, 0, 5, -1)

	// Assert
	convey.Convey("SafeIndexesDefaultWithDetail tracks missing", t, func() {
		convey.So(result.IsAnyMissing, convey.ShouldBeTrue)
		convey.So(len(result.MissingIndexes), convey.ShouldBeGreaterThan, 0)
	})
}

// --- SafeRangeItems ---

func Test_SafeRangeItems_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c", "d"}

	// Act
	result := stringslice.SafeRangeItems(slice, 1, 3)

	// Assert
	convey.Convey("SafeRangeItems returns correct range", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_SafeRangeItems_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SafeRangeItems(nil, 0, 2)

	// Assert
	convey.Convey("SafeRangeItems nil returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_SafeRangeItems_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SafeRangeItems([]string{}, 0, 2)

	// Assert
	convey.Convey("SafeRangeItems empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_SafeRangeItems_StartBeyondLastIndex(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeRangeItems(slice, 5, 10)

	// Assert
	convey.Convey("SafeRangeItems start beyond last returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_SafeRangeItems_EndBeyondLastIndex(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	result := stringslice.SafeRangeItems(slice, 0, 100)

	// Assert
	convey.Convey("SafeRangeItems clamps end to last index", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_SafeRangeItems_InvalidStartValue(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	result := stringslice.SafeRangeItems(slice, -1, 2)

	// Assert
	convey.Convey("SafeRangeItems invalid start uses 0", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

// --- SafeRangeItemsPtr ---

func Test_SafeRangeItemsPtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SafeRangeItemsPtr([]string{}, 0, 2)

	// Assert
	convey.Convey("SafeRangeItemsPtr empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- SafeIndexesPtr ---

func Test_SafeIndexesPtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SafeIndexesPtr([]string{}, 0)

	// Assert
	convey.Convey("SafeIndexesPtr empty returns default", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

// --- IndexesDefault ---

func Test_IndexesDefault_Valid(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	result := stringslice.IndexesDefault(slice, 0, 2)

	// Assert
	convey.Convey("IndexesDefault returns values at indexes", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
		convey.So(result[0], convey.ShouldEqual, "a")
	})
}

func Test_IndexesDefault_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.IndexesDefault([]string{}, 0)

	// Assert
	convey.Convey("IndexesDefault empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_IndexesDefault_NoIndexes_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.IndexesDefault([]string{"a"})

	// Assert
	convey.Convey("IndexesDefault no indexes returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NonEmptyIf ---

func Test_NonEmptyIf_True_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "", "b"}

	// Act
	result := stringslice.NonEmptyIf(true, slice)

	// Assert
	convey.Convey("NonEmptyIf true filters empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_NonEmptyIf_False_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "", "b"}

	// Act
	result := stringslice.NonEmptyIf(false, slice)

	// Assert
	convey.Convey("NonEmptyIf false returns NonNullStrings", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

// --- NonEmptyJoin ---

func Test_NonEmptyJoin_Valid(t *testing.T) {
	// Arrange
	slice := []string{"a", "", "b"}

	// Act
	result := stringslice.NonEmptyJoin(slice, ",")

	// Assert
	convey.Convey("NonEmptyJoin joins non-empty", t, func() {
		convey.So(result, convey.ShouldEqual, "a,b")
	})
}

func Test_NonEmptyJoin_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonEmptyJoin(nil, ",")

	// Assert
	convey.Convey("NonEmptyJoin nil returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_NonEmptyJoin_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonEmptyJoin([]string{}, ",")

	// Assert
	convey.Convey("NonEmptyJoin empty returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- NonEmptyJoinPtr ---

func Test_NonEmptyJoinPtr_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "", "b"}

	// Act
	result := stringslice.NonEmptyJoinPtr(slice, ",")

	// Assert
	convey.Convey("NonEmptyJoinPtr joins non-empty", t, func() {
		convey.So(result, convey.ShouldEqual, "a,b")
	})
}

func Test_NonEmptyJoinPtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonEmptyJoinPtr([]string{}, ",")

	// Assert
	convey.Convey("NonEmptyJoinPtr empty returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- NonWhitespaceJoin ---

func Test_NonWhitespaceJoin_Valid(t *testing.T) {
	// Arrange
	slice := []string{"a", "   ", "b"}

	// Act
	result := stringslice.NonWhitespaceJoin(slice, ",")

	// Assert
	convey.Convey("NonWhitespaceJoin joins non-whitespace", t, func() {
		convey.So(result, convey.ShouldEqual, "a,b")
	})
}

func Test_NonWhitespaceJoin_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonWhitespaceJoin(nil, ",")

	// Assert
	convey.Convey("NonWhitespaceJoin nil returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_NonWhitespaceJoin_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonWhitespaceJoin([]string{}, ",")

	// Assert
	convey.Convey("NonWhitespaceJoin empty returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- NonWhitespaceJoinPtr ---

func Test_NonWhitespaceJoinPtr_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "   ", "b"}

	// Act
	result := stringslice.NonWhitespaceJoinPtr(slice, ",")

	// Assert
	convey.Convey("NonWhitespaceJoinPtr joins non-whitespace", t, func() {
		convey.So(result, convey.ShouldEqual, "a,b")
	})
}

func Test_NonWhitespaceJoinPtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonWhitespaceJoinPtr([]string{}, ",")

	// Assert
	convey.Convey("NonWhitespaceJoinPtr empty returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- TrimmedEachWordsIf ---

func Test_TrimmedEachWordsIf_True_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{" a ", "  ", " b "}

	// Act
	result := stringslice.TrimmedEachWordsIf(true, slice)

	// Assert
	convey.Convey("TrimmedEachWordsIf true trims", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_TrimmedEachWordsIf_False_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "", "b"}

	// Act
	result := stringslice.TrimmedEachWordsIf(false, slice)

	// Assert
	convey.Convey("TrimmedEachWordsIf false returns NonNullStrings", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

// --- TrimmedEachWordsPtr ---

func Test_TrimmedEachWordsPtr_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{" a ", " b "}

	// Act
	result := stringslice.TrimmedEachWordsPtr(slice)

	// Assert
	convey.Convey("TrimmedEachWordsPtr trims all", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_TrimmedEachWordsPtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.TrimmedEachWordsPtr([]string{})

	// Assert
	convey.Convey("TrimmedEachWordsPtr empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- TrimmedEachWords nil input ---

func Test_TrimmedEachWords_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.TrimmedEachWords(nil)

	// Assert
	convey.Convey("TrimmedEachWords nil returns nil", t, func() {
		convey.So(result, convey.ShouldBeNil)
	})
}

// --- PrependLineNew ---

func Test_PrependLineNew_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"b", "c"}

	// Act
	result := stringslice.PrependLineNew("a", slice)

	// Assert
	convey.Convey("PrependLineNew prepends line", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
		convey.So(result[0], convey.ShouldEqual, "a")
	})
}

// --- AppendLineNew ---

func Test_AppendLineNew_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.AppendLineNew(slice, "c")

	// Assert
	convey.Convey("AppendLineNew appends line", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

// --- LinesProcess ---

func Test_LinesProcess_Valid(t *testing.T) {
	// Arrange
	lines := []string{"a", "b", "c"}

	// Act
	result := stringslice.LinesProcess(lines, func(index int, lineIn string) (string, bool, bool) {
		return lineIn + "!", true, false
	})

	// Assert
	convey.Convey("LinesProcess processes all lines", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

func Test_LinesProcess_Break(t *testing.T) {
	// Arrange
	lines := []string{"a", "b", "c"}

	// Act
	result := stringslice.LinesProcess(lines, func(index int, lineIn string) (string, bool, bool) {
		return lineIn, true, index == 0
	})

	// Assert
	convey.Convey("LinesProcess breaks early", t, func() {
		convey.So(len(result), convey.ShouldEqual, 1)
	})
}

func Test_LinesProcess_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.LinesProcess([]string{}, func(index int, lineIn string) (string, bool, bool) {
		return lineIn, true, false
	})

	// Assert
	convey.Convey("LinesProcess empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_LinesProcess_SkipFalse(t *testing.T) {
	// Arrange
	lines := []string{"a", "b", "c"}

	// Act
	result := stringslice.LinesProcess(lines, func(index int, lineIn string) (string, bool, bool) {
		return lineIn, index%2 == 0, false // skip odd indexes
	})

	// Assert
	convey.Convey("LinesProcess skips when isTake=false", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

// --- LinesSimpleProcess ---

func Test_LinesSimpleProcess_Valid(t *testing.T) {
	// Arrange
	lines := []string{"a", "b"}

	// Act
	result := stringslice.LinesSimpleProcess(lines, func(lineIn string) string {
		return lineIn + "!"
	})

	// Assert
	convey.Convey("LinesSimpleProcess processes all", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
		convey.So(result[0], convey.ShouldEqual, "a!")
	})
}

func Test_LinesSimpleProcess_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.LinesSimpleProcess([]string{}, func(lineIn string) string {
		return lineIn
	})

	// Assert
	convey.Convey("LinesSimpleProcess empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- LinesSimpleProcessNoEmpty ---

func Test_LinesSimpleProcessNoEmpty_Valid(t *testing.T) {
	// Arrange
	lines := []string{"a", "b", ""}

	// Act
	result := stringslice.LinesSimpleProcessNoEmpty(lines, func(lineIn string) string {
		return lineIn
	})

	// Assert
	convey.Convey("LinesSimpleProcessNoEmpty filters empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_LinesSimpleProcessNoEmpty_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.LinesSimpleProcessNoEmpty([]string{}, func(lineIn string) string {
		return lineIn
	})

	// Assert
	convey.Convey("LinesSimpleProcessNoEmpty empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NonEmptySlice ---

func Test_NonEmptySlice_Valid(t *testing.T) {
	// Arrange
	slice := []string{"a", "", "b"}

	// Act
	result := stringslice.NonEmptySlice(slice)

	// Assert
	convey.Convey("NonEmptySlice filters empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}

func Test_NonEmptySlice_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonEmptySlice([]string{})

	// Assert
	convey.Convey("NonEmptySlice empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NonEmptyStrings ---

func Test_NonEmptyStrings_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonEmptyStrings(nil)

	// Assert
	convey.Convey("NonEmptyStrings nil returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NonWhitespace ---

func Test_NonWhitespace_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonWhitespace(nil)

	// Assert
	convey.Convey("NonWhitespace nil returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_NonWhitespace_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonWhitespace([]string{})

	// Assert
	convey.Convey("NonWhitespace empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NonWhitespacePtr ---

func Test_NonWhitespacePtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonWhitespacePtr([]string{})

	// Assert
	convey.Convey("NonWhitespacePtr empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- NonNullStrings ---

func Test_NonNullStrings_Nil_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonNullStrings(nil)

	// Assert
	convey.Convey("NonNullStrings nil returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- SplitTrimmedNonEmpty ---

func Test_SplitTrimmedNonEmpty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SplitTrimmedNonEmpty("a , b , c", ",", -1)

	// Assert
	convey.Convey("SplitTrimmedNonEmpty splits and trims", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
		convey.So(result[0], convey.ShouldEqual, "a")
	})
}

// --- SplitTrimmedNonEmptyAll ---

func Test_SplitTrimmedNonEmptyAll_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.SplitTrimmedNonEmptyAll("a , b , c", ",")

	// Assert
	convey.Convey("SplitTrimmedNonEmptyAll splits and trims all", t, func() {
		convey.So(len(result), convey.ShouldEqual, 3)
	})
}

// --- SafeIndexAtUsingLastIndex ---

func Test_SafeIndexAtUsingLastIndex_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	result := stringslice.SafeIndexAtUsingLastIndex(slice, 2, 1)

	// Assert
	convey.Convey("SafeIndexAtUsingLastIndex returns correct value", t, func() {
		convey.So(result, convey.ShouldEqual, "b")
	})
}

func Test_SafeIndexAtUsingLastIndex_OutOfBounds(t *testing.T) {
	// Arrange
	slice := []string{"a"}

	// Act
	result := stringslice.SafeIndexAtUsingLastIndex(slice, 0, 1)

	// Assert
	convey.Convey("SafeIndexAtUsingLastIndex returns empty for out of bounds", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_SafeIndexAtUsingLastIndex_NegativeIndex_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a"}

	// Act
	result := stringslice.SafeIndexAtUsingLastIndex(slice, 0, -1)

	// Assert
	convey.Convey("SafeIndexAtUsingLastIndex returns empty for negative", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- SafeIndexAtUsingLastIndexPtr ---

func Test_SafeIndexAtUsingLastIndexPtr_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexAtUsingLastIndexPtr(slice, 1, 0)

	// Assert
	convey.Convey("SafeIndexAtUsingLastIndexPtr returns value", t, func() {
		convey.So(result, convey.ShouldEqual, "a")
	})
}

func Test_SafeIndexAtUsingLastIndexPtr_Zero(t *testing.T) {
	// Arrange
	slice := []string{"a"}

	// Act
	result := stringslice.SafeIndexAtUsingLastIndexPtr(slice, 0, 0)

	// Assert
	convey.Convey("SafeIndexAtUsingLastIndexPtr zero lastIndex returns empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// --- SafeIndexAtWith ---

func Test_SafeIndexAtWith_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexAtWith(slice, 0, "default")

	// Assert
	convey.Convey("SafeIndexAtWith returns value", t, func() {
		convey.So(result, convey.ShouldEqual, "a")
	})
}

func Test_SafeIndexAtWith_Default_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a"}

	// Act
	result := stringslice.SafeIndexAtWith(slice, 5, "default")

	// Assert
	convey.Convey("SafeIndexAtWith returns default", t, func() {
		convey.So(result, convey.ShouldEqual, "default")
	})
}

// --- SafeIndexAtWithPtr ---

func Test_SafeIndexAtWithPtr_Valid_FromProcessAsyncValidFin(t *testing.T) {
	// Arrange
	slice := []string{"a", "b"}

	// Act
	result := stringslice.SafeIndexAtWithPtr(slice, 0, "default")

	// Assert
	convey.Convey("SafeIndexAtWithPtr returns value", t, func() {
		convey.So(result, convey.ShouldEqual, "a")
	})
}

// --- PrependNew ---

func Test_PrependNew_Valid(t *testing.T) {
	// Arrange
	slice := []string{"b", "c"}

	// Act
	result := stringslice.PrependNew(slice, "a")

	// Assert
	convey.Convey("PrependNew prepends items", t, func() {
		convey.So(len(*result), convey.ShouldEqual, 3)
		convey.So((*result)[0], convey.ShouldEqual, "a")
	})
}

func Test_PrependNew_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.PrependNew([]string{})

	// Assert
	convey.Convey("PrependNew empty returns empty", t, func() {
		convey.So(len(*result), convey.ShouldEqual, 0)
	})
}

// --- NonEmptySlicePtr ---

func Test_NonEmptySlicePtr_Empty_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.NonEmptySlicePtr([]string{})

	// Assert
	convey.Convey("NonEmptySlicePtr empty returns empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

// --- MakeDefault, MakeDefaultPtr, MakeLenPtr, MakePtr ---

func Test_MakeDefault_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.MakeDefault(10)

	// Assert
	convey.Convey("MakeDefault creates empty with capacity", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
		convey.So(cap(result), convey.ShouldBeGreaterThan, 0)
	})
}

func Test_MakeDefaultPtr_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.MakeDefaultPtr(10)

	// Assert
	convey.Convey("MakeDefaultPtr creates empty", t, func() {
		convey.So(len(result), convey.ShouldEqual, 0)
	})
}

func Test_MakeLenPtr_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.MakeLenPtr(5)

	// Assert
	convey.Convey("MakeLenPtr creates with length", t, func() {
		convey.So(len(result), convey.ShouldEqual, 5)
	})
}

func Test_MakePtr_FromProcessAsyncValidFin(t *testing.T) {
	// Act
	result := stringslice.MakePtr(2, 10)

	// Assert
	convey.Convey("MakePtr creates with length and capacity", t, func() {
		convey.So(len(result), convey.ShouldEqual, 2)
	})
}
