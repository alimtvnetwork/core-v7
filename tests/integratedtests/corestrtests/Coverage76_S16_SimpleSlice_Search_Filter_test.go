package corestrtests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ========================================
// S16: SimpleSlice search/filter/comparison
//   IsEqual variants, Clone, DistinctDiff,
//   RemoveIndexes, IsEqualByFunc, AddedRemovedLinesDiff
// ========================================

func Test_C76_SimpleSlice_IsEqual_BothEqual(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqual_BothEqual", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a", "b")
		ss2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if !ss1.IsEqual(ss2) {
			t.Error("expected equal")
		}
	})
}

func Test_C76_SimpleSlice_IsEqual_DiffContent(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqual_DiffContent", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a", "b")
		ss2 := corestr.New.SimpleSlice.Lines("a", "c")

		// Act & Assert
		if ss1.IsEqual(ss2) {
			t.Error("expected not equal")
		}
	})
}

func Test_C76_SimpleSlice_IsEqual_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqual_DiffLength", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a")
		ss2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if ss1.IsEqual(ss2) {
			t.Error("expected not equal due to length")
		}
	})
}

func Test_C76_SimpleSlice_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqual_BothNil", func() {
		// Arrange
		var ss1 *corestr.SimpleSlice
		var ss2 *corestr.SimpleSlice

		// Act & Assert
		if !ss1.IsEqual(ss2) {
			t.Error("expected true for both nil")
		}
	})
}

func Test_C76_SimpleSlice_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqual_OneNil", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a")
		var ss2 *corestr.SimpleSlice

		// Act & Assert
		if ss1.IsEqual(ss2) {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqual_BothEmpty", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Empty()
		ss2 := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss1.IsEqual(ss2) {
			t.Error("expected true for both empty")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualLines(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if !ss.IsEqualLines([]string{"a", "b"}) {
			t.Error("expected true")
		}
		if ss.IsEqualLines([]string{"a", "c"}) {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualLines_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsEqualLines([]string{"a", "b"}) {
			t.Error("expected false for diff length")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualLines_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualLines_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		if !ss.IsEqualLines(nil) {
			t.Error("expected true for both nil")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualLines_OneNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualLines_OneNil", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsEqualLines(nil) {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLines(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLines", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act & Assert
		if !ss.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Error("expected true for unordered equal")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLines_Mismatch(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLines_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if ss.IsEqualUnorderedLines([]string{"a", "c"}) {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLines_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLines_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsEqualUnorderedLines([]string{"a", "b"}) {
			t.Error("expected false for diff length")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLines_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLines_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		if !ss.IsEqualUnorderedLines(nil) {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLines_BothEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss.IsEqualUnorderedLines([]string{}) {
			t.Error("expected true for both empty")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLinesClone(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLinesClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("c", "a", "b")

		// Act & Assert
		if !ss.IsEqualUnorderedLinesClone([]string{"b", "a", "c"}) {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsEqualUnorderedLinesClone([]string{"a", "b"}) {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if ss.IsEqualUnorderedLinesClone([]string{"a", "c"}) {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		if !ss.IsEqualUnorderedLinesClone(nil) {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualUnorderedLinesClone_BothEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss.IsEqualUnorderedLinesClone([]string{}) {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_Clone_Deep(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Clone_Deep", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		cloned := ss.Clone(true)

		// Assert
		if cloned.Length() != 2 {
			t.Errorf("expected 2, got %d", cloned.Length())
		}
	})
}

func Test_C76_SimpleSlice_Clone_Shallow(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Clone_Shallow", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		cloned := ss.Clone(false)

		// Assert
		if cloned.Length() != 1 {
			t.Error("expected 1")
		}
	})
}
	safeTest(t, "Test_C76_SimpleSlice_DeepClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x", "y")

		// Act
		cloned := ss.DeepClone()

		// Assert
		if cloned.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C76_SimpleSlice_ShadowClone(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_ShadowClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x")

		// Act
		cloned := ss.ShadowClone()

		// Assert
		if cloned.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_IsDistinctEqualRaw(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsDistinctEqualRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "a")

		// Act & Assert
		if !ss.IsDistinctEqualRaw("a", "b") {
			t.Error("expected true")
		}
		if ss.IsDistinctEqualRaw("a", "c") {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsDistinctEqual(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsDistinctEqual", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a", "b", "a")
		ss2 := corestr.New.SimpleSlice.Lines("b", "a")

		// Act & Assert
		if !ss1.IsDistinctEqual(ss2) {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqualRaw_Clone(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqualRaw_Clone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act & Assert
		if !ss.IsUnorderedEqualRaw(true, "a", "b") {
			t.Error("expected true with clone")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqualRaw_NoClone(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqualRaw_NoClone", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("b", "a")

		// Act & Assert
		if !ss.IsUnorderedEqualRaw(false, "a", "b") {
			t.Error("expected true without clone")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqualRaw_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqualRaw_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsUnorderedEqualRaw(false, "a", "b") {
			t.Error("expected false for diff length")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqualRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqualRaw_BothEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss.IsUnorderedEqualRaw(false) {
			t.Error("expected true for both empty")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqual_Clone(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqual_Clone", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("b", "a")
		ss2 := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if !ss1.IsUnorderedEqual(true, ss2) {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqual_BothEmpty", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Empty()
		ss2 := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss1.IsUnorderedEqual(false, ss2) {
			t.Error("expected true for both empty")
		}
	})
}

func Test_C76_SimpleSlice_IsUnorderedEqual_NilRight(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsUnorderedEqual_NilRight", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsUnorderedEqual(false, nil) {
			t.Error("expected false for nil right")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFunc(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("hello", "world")

		// Act
		result := ss.IsEqualByFunc(func(index int, left, right string) bool {
			return strings.EqualFold(left, right)
		}, "HELLO", "WORLD")

		// Assert
		if !result {
			t.Error("expected true for case-insensitive match")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFunc_Mismatch(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFunc_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.IsEqualByFunc(func(index int, left, right string) bool {
			return left == right
		}, "a", "c")

		// Assert
		if result {
			t.Error("expected false for mismatch")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFunc_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFunc_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsEqualByFunc(func(i int, l, r string) bool { return true }, "a", "b") {
			t.Error("expected false for diff length")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFunc_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFunc_BothEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss.IsEqualByFunc(func(i int, l, r string) bool { return true }) {
			t.Error("expected true for both empty")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFuncLinesSplit(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFuncLinesSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool {
			return l == r
		})

		// Assert
		if !result {
			t.Error("expected true")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_WithTrim(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_WithTrim", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines(" a ", " b ")

		// Act
		result := ss.IsEqualByFuncLinesSplit(true, ",", " a , b ", func(i int, l, r string) bool {
			return l == r
		})

		// Assert
		if !result {
			t.Error("expected true with trim")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_DiffLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.IsEqualByFuncLinesSplit(false, ",", "a,b", func(i int, l, r string) bool { return true }) {
			t.Error("expected false for diff length")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		// strings.Split("", ",") returns [""] (length 1), empty slice has length 0 → not equal
		if ss.IsEqualByFuncLinesSplit(false, ",", "", func(i int, l, r string) bool { return true }) {
			t.Error("expected false for empty vs single-element split")
		}
	})
}

func Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsEqualByFuncLinesSplit_Mismatch", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.IsEqualByFuncLinesSplit(false, ",", "a,c", func(i int, l, r string) bool {
			return l == r
		})

		// Assert
		if result {
			t.Error("expected false for mismatch")
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiffRaw(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiffRaw", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		diff := ss.DistinctDiffRaw("b", "c", "d")

		// Assert
		if len(diff) == 0 {
			t.Error("expected non-empty diff")
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiffRaw_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiffRaw_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		diff := ss.DistinctDiffRaw()

		// Assert
		if len(diff) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiffRaw_LeftNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiffRaw_LeftNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		diff := ss.DistinctDiffRaw("a", "b")

		// Assert
		if len(diff) != 2 {
			t.Errorf("expected 2, got %d", len(diff))
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiffRaw_RightNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiffRaw_RightNil", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		diff := ss.DistinctDiffRaw()

		// Assert
		if len(diff) != 2 {
			t.Errorf("expected 2, got %d", len(diff))
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiff(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiff", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("a", "b")
		ss2 := corestr.New.SimpleSlice.Lines("b", "c")

		// Act
		diff := ss1.DistinctDiff(ss2)

		// Assert
		if len(diff) == 0 {
			t.Error("expected non-empty diff")
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiff_BothNil", func() {
		// Arrange
		var ss1 *corestr.SimpleSlice
		var ss2 *corestr.SimpleSlice

		// Act
		diff := ss1.DistinctDiff(ss2)

		// Assert
		if len(diff) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiff_LeftNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiff_LeftNil", func() {
		// Arrange
		var ss1 *corestr.SimpleSlice
		ss2 := corestr.New.SimpleSlice.Lines("a")

		// Act
		diff := ss1.DistinctDiff(ss2)

		// Assert
		if len(diff) != 1 {
			t.Errorf("expected 1, got %d", len(diff))
		}
	})
}

func Test_C76_SimpleSlice_DistinctDiff_RightNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DistinctDiff_RightNil", func() {
		// Arrange
		ss1 := corestr.New.SimpleSlice.Lines("x")

		// Act
		diff := ss1.DistinctDiff(nil)

		// Assert
		if len(diff) != 1 {
			t.Errorf("expected 1, got %d", len(diff))
		}
	})
}

func Test_C76_SimpleSlice_AddedRemovedLinesDiff(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddedRemovedLinesDiff", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		added, removed := ss.AddedRemovedLinesDiff("b", "c")

		// Assert
		if len(added) == 0 {
			t.Error("expected some added lines")
		}
		if len(removed) == 0 {
			t.Error("expected some removed lines")
		}
	})
}

func Test_C76_SimpleSlice_AddedRemovedLinesDiff_BothNil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddedRemovedLinesDiff_BothNil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		added, removed := ss.AddedRemovedLinesDiff()

		// Assert
		if added != nil || removed != nil {
			t.Error("expected nil for both nil inputs")
		}
	})
}

func Test_C76_SimpleSlice_RemoveIndexes(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_RemoveIndexes", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c", "d")

		// Act
		result, err := ss.RemoveIndexes(1, 3)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C76_SimpleSlice_RemoveIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_RemoveIndexes_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		_, err := ss.RemoveIndexes(0)

		// Assert
		if err == nil {
			t.Error("expected error for empty slice")
		}
	})
}

func Test_C76_SimpleSlice_RemoveIndexes_InvalidIndex(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_RemoveIndexes_InvalidIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result, err := ss.RemoveIndexes(5)

		// Assert
		if err == nil {
			t.Error("expected error for invalid index")
		}
		if result.Length() != 2 {
			t.Errorf("expected 2 (all kept), got %d", result.Length())
		}
	})
}

// --- Additional Add methods ---

func Test_C76_SimpleSlice_AddSplit(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddSplit("a:b:c", ":")

		// Assert
		if ss.Length() != 3 {
			t.Errorf("expected 3, got %d", ss.Length())
		}
	})
}

func Test_C76_SimpleSlice_AddIf_True(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddIf_True", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddIf(true, "x")

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddIf_False(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddIf_False", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddIf(false, "x")

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Adds_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.Adds()

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_Append_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Append_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.Append()

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_AppendFmt(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AppendFmt("hello %s", "world")

		// Assert
		if ss.Length() != 1 || ss.First() != "hello world" {
			t.Errorf("expected 'hello world', got '%s'", ss.First())
		}
	})
}

func Test_C76_SimpleSlice_AppendFmt_EmptyFormat(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AppendFmt_EmptyFormat", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AppendFmt("")

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0 for empty format with no args")
		}
	})
}

func Test_C76_SimpleSlice_AppendFmtIf_True(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AppendFmtIf_True", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AppendFmtIf(true, "val=%d", 42)

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AppendFmtIf_False(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AppendFmtIf_False", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AppendFmtIf(false, "val=%d", 42)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_AddAsTitleValue(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddAsTitleValue("Name", "Alice")

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddAsCurlyTitleWrap("Key", "Val")

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddAsCurlyTitleWrapIf_True(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddAsCurlyTitleWrapIf_True", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddAsCurlyTitleWrapIf(true, "K", "V")

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddAsCurlyTitleWrapIf_False(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddAsCurlyTitleWrapIf_False", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddAsCurlyTitleWrapIf(false, "K", "V")

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_AddAsTitleValueIf_True(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddAsTitleValueIf_True", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddAsTitleValueIf(true, "T", "V")

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddAsTitleValueIf_False(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddAsTitleValueIf_False", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddAsTitleValueIf(false, "T", "V")

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_InsertAt(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "c")

		// Act
		ss.InsertAt(1, "b")

		// Assert
		if ss.Length() != 3 {
			t.Errorf("expected 3, got %d", ss.Length())
		}
		strs := ss.Strings()
		if strs[1] != "b" {
			t.Errorf("expected 'b' at index 1, got '%s'", strs[1])
		}
	})
}

func Test_C76_SimpleSlice_InsertAt_NegativeIndex(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_InsertAt_NegativeIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.InsertAt(-1, "x")

		// Assert — should not add
		if ss.Length() != 1 {
			t.Error("expected 1, negative index should be ignored")
		}
	})
}

func Test_C76_SimpleSlice_InsertAt_BeyondLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_InsertAt_BeyondLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		ss.InsertAt(5, "x")

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddStruct(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddStruct", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()
		type testStruct struct{ Name string }

		// Act
		ss.AddStruct(false, testStruct{Name: "test"})

		// Assert
		if ss.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_AddStruct_Nil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddStruct_Nil", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddStruct(false, nil)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_AddPointer_Nil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddPointer_Nil", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddPointer(false, nil)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_AddsIf_True(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddsIf_True", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddsIf(true, "a", "b")

		// Assert
		if ss.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C76_SimpleSlice_AddsIf_False(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddsIf_False", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddsIf(false, "a", "b")

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_AddError(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AddError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		ss.AddError(nil)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0 for nil error")
		}
	})
}

func Test_C76_SimpleSlice_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AsDefaultError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("err1", "err2")

		// Act
		err := ss.AsDefaultError()

		// Assert
		if err == nil {
			t.Error("expected non-nil error")
		}
	})
}

func Test_C76_SimpleSlice_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_AsError_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		err := ss.AsError(",")

		// Assert
		if err != nil {
			t.Error("expected nil for empty")
		}
	})
}

func Test_C76_SimpleSlice_CountFunc(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "bb", "ccc")

		// Act
		count := ss.CountFunc(func(index int, item string) bool {
			return len(item) > 1
		})

		// Assert
		if count != 2 {
			t.Errorf("expected 2, got %d", count)
		}
	})
}

func Test_C76_SimpleSlice_CountFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_CountFunc_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		count := ss.CountFunc(func(i int, s string) bool { return true })

		// Assert
		if count != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C76_SimpleSlice_IsContains(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if !ss.IsContains("a") {
			t.Error("expected true")
		}
		if ss.IsContains("z") {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_IsContains_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsContains_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.IsContains("a") {
			t.Error("expected false for empty")
		}
	})
}

func Test_C76_SimpleSlice_IsContainsFunc(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("Hello", "World")

		// Act
		found := ss.IsContainsFunc("hello", func(item, searching string) bool {
			return strings.EqualFold(item, searching)
		})

		// Assert
		if !found {
			t.Error("expected true for case-insensitive search")
		}
	})
}

func Test_C76_SimpleSlice_IsContainsFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IsContainsFunc_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.IsContainsFunc("a", func(i, s string) bool { return true }) {
			t.Error("expected false for empty")
		}
	})
}

func Test_C76_SimpleSlice_IndexOfFunc(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		idx := ss.IndexOfFunc("b", func(item, searching string) bool {
			return item == searching
		})

		// Assert
		if idx != 1 {
			t.Errorf("expected 1, got %d", idx)
		}
	})
}

func Test_C76_SimpleSlice_IndexOfFunc_NotFound(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IndexOfFunc_NotFound", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		idx := ss.IndexOfFunc("z", func(item, searching string) bool {
			return item == searching
		})

		// Assert
		if idx != -1 {
			t.Errorf("expected -1, got %d", idx)
		}
	})
}

func Test_C76_SimpleSlice_IndexOfFunc_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IndexOfFunc_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		idx := ss.IndexOfFunc("a", func(i, s string) bool { return true })

		// Assert
		if idx != -1 {
			t.Errorf("expected -1, got %d", idx)
		}
	})
}

func Test_C76_SimpleSlice_IndexOf(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IndexOf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x", "y", "z")

		// Act & Assert
		if ss.IndexOf("y") != 1 {
			t.Error("expected 1")
		}
		if ss.IndexOf("w") != -1 {
			t.Error("expected -1")
		}
	})
}

func Test_C76_SimpleSlice_IndexOf_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_IndexOf_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.IndexOf("a") != -1 {
			t.Error("expected -1")
		}
	})
}

func Test_C76_SimpleSlice_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_HasAnyItem", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")
		empty := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if !ss.HasAnyItem() {
			t.Error("expected true")
		}
		if empty.HasAnyItem() {
			t.Error("expected false")
		}
	})
}

func Test_C76_SimpleSlice_HasIndex(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_HasIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if !ss.HasIndex(0) {
			t.Error("expected true for 0")
		}
		if !ss.HasIndex(1) {
			t.Error("expected true for 1")
		}
		if ss.HasIndex(2) {
			t.Error("expected false for 2")
		}
		if ss.HasIndex(-1) {
			t.Error("expected false for -1")
		}
	})
}

func Test_C76_SimpleSlice_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_WrapDoubleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.WrapDoubleQuote()

		// Assert
		if result.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C76_SimpleSlice_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_WrapSingleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.WrapSingleQuote()

		// Assert
		if result.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_WrapTildaQuote(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_WrapTildaQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.WrapTildaQuote()

		// Assert
		if result.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C76_SimpleSlice_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", `"b"`)

		// Act
		result := ss.WrapDoubleQuoteIfMissing()

		// Assert
		if result.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C76_SimpleSlice_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_WrapSingleQuoteIfMissing", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "'b'")

		// Act
		result := ss.WrapSingleQuoteIfMissing()

		// Assert
		if result.Length() != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C76_SimpleSlice_FirstDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_FirstDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("first", "second")

		// Act
		result := ss.FirstDynamic()

		// Assert
		if result != "first" {
			t.Error("expected 'first'")
		}
	})
}

func Test_C76_SimpleSlice_LastDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_LastDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("first", "last")

		// Act
		result := ss.LastDynamic()

		// Assert
		if result != "last" {
			t.Error("expected 'last'")
		}
	})
}

func Test_C76_SimpleSlice_FirstOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_FirstOrDefault_NonEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act & Assert
		if ss.FirstOrDefault() != "a" {
			t.Error("expected 'a'")
		}
	})
}

func Test_C76_SimpleSlice_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_FirstOrDefault_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.FirstOrDefault() != "" {
			t.Error("expected empty string")
		}
	})
}

func Test_C76_SimpleSlice_FirstOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_FirstOrDefaultDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x")

		// Act
		result := ss.FirstOrDefaultDynamic()

		// Assert
		if result != "x" {
			t.Error("expected 'x'")
		}
	})
}

func Test_C76_SimpleSlice_LastOrDefault_NonEmpty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_LastOrDefault_NonEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act & Assert
		if ss.LastOrDefault() != "b" {
			t.Error("expected 'b'")
		}
	})
}

func Test_C76_SimpleSlice_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_LastOrDefault_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act & Assert
		if ss.LastOrDefault() != "" {
			t.Error("expected empty string")
		}
	})
}

func Test_C76_SimpleSlice_LastOrDefaultDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_LastOrDefaultDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("x")

		// Act
		result := ss.LastOrDefaultDynamic()

		// Assert
		if result != "x" {
			t.Error("expected 'x'")
		}
	})
}

func Test_C76_SimpleSlice_SkipDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_SkipDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.SkipDynamic(1)

		// Assert
		asSlice, ok := result.([]string)
		if !ok {
			t.Error("expected []string type")
		}
		if len(asSlice) != 2 {
			t.Errorf("expected 2, got %d", len(asSlice))
		}
	})
}

func Test_C76_SimpleSlice_SkipDynamic_BeyondLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_SkipDynamic_BeyondLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.SkipDynamic(5)

		// Assert
		asSlice, ok := result.([]string)
		if !ok {
			t.Error("expected []string type")
		}
		if len(asSlice) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C76_SimpleSlice_Skip(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Skip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.Skip(2)

		// Assert
		if len(result) != 1 || result[0] != "c" {
			t.Errorf("expected ['c'], got %v", result)
		}
	})
}

func Test_C76_SimpleSlice_TakeDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_TakeDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.TakeDynamic(2)

		// Assert
		asSlice, ok := result.(corestr.SimpleSlice)
		if !ok {
			// May also return as []string
			asStrSlice, ok2 := result.([]string)
			if !ok2 {
				t.Error("expected SimpleSlice or []string type")
			} else if len(asStrSlice) != 2 {
				t.Errorf("expected 2, got %d", len(asStrSlice))
			}
		} else if asSlice.Length() != 2 {
			t.Errorf("expected 2, got %d", asSlice.Length())
		}
	})
}

func Test_C76_SimpleSlice_Take(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Take", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.Take(2)

		// Assert
		if len(result) != 2 {
			t.Errorf("expected 2, got %d", len(result))
		}
	})
}

func Test_C76_SimpleSlice_Take_BeyondLength(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Take_BeyondLength", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		result := ss.Take(5)

		// Assert
		if len(result) != 1 {
			t.Errorf("expected 1, got %d", len(result))
		}
	})
}

func Test_C76_SimpleSlice_LimitDynamic(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_LimitDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")

		// Act
		result := ss.LimitDynamic(1)

		// Assert
		if result == nil {
			t.Error("expected non-nil")
		}
	})
}

func Test_C76_SimpleSlice_Limit(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Limit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		result := ss.Limit(1)

		// Assert
		if len(result) != 1 {
			t.Errorf("expected 1, got %d", len(result))
		}
	})
}

func Test_C76_SimpleSlice_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Length_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		if ss.Length() != 0 {
			t.Error("expected 0 for nil")
		}
	})
}

func Test_C76_SimpleSlice_Strings_List(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Strings_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		strs := ss.Strings()
		list := ss.List()

		// Assert
		if len(strs) != 2 || len(list) != 2 {
			t.Error("expected 2 for both")
		}
	})
}

func Test_C76_SimpleSlice_DeserializeJsoner(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_DeserializeJsoner", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		jsoner := ss.AsJsoner()

		// Act
		result, err := corestr.New.SimpleSlice.DeserializeJsoner(jsoner)

		// Assert
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C76_SimpleSlice_Map(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Map", func() {
		// Arrange
		m := map[string]int{"a": 1, "b": 2}

		// Act
		ss := corestr.New.SimpleSlice.Map(m)

		// Assert
		if ss.Length() != 2 {
			t.Errorf("expected 2, got %d", ss.Length())
		}
	})
}

func Test_C76_SimpleSlice_Map_Empty(t *testing.T) {
	safeTest(t, "Test_C76_SimpleSlice_Map_Empty", func() {
		// Arrange
		m := map[string]int{}

		// Act
		ss := corestr.New.SimpleSlice.Map(m)

		// Assert
		if ss.Length() != 0 {
			t.Error("expected 0")
		}
	})
}
