package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_NewLeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR New", Expected: true, Actual: lr.IsValid && lr.Left == "a" && lr.Right == "b", Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_InvalidLeftRight", func() {
		lr := corestr.InvalidLeftRight("err")
		tc := caseV1Compat{Name: "LR Invalid", Expected: false, Actual: lr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		tc := caseV1Compat{Name: "LR InvalidNoMsg", Expected: false, Actual: lr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_LeftBytes(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_LeftBytes", func() {
		lr := corestr.NewLeftRight("ab", "cd")
		tc := caseV1Compat{Name: "LR LeftBytes", Expected: 2, Actual: len(lr.LeftBytes()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_RightBytes(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_RightBytes", func() {
		lr := corestr.NewLeftRight("ab", "cd")
		tc := caseV1Compat{Name: "LR RightBytes", Expected: 2, Actual: len(lr.RightBytes()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_LeftTrim(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_LeftTrim", func() {
		lr := corestr.NewLeftRight(" a ", "b")
		tc := caseV1Compat{Name: "LR LeftTrim", Expected: "a", Actual: lr.LeftTrim(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_RightTrim(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_RightTrim", func() {
		lr := corestr.NewLeftRight("a", " b ")
		tc := caseV1Compat{Name: "LR RightTrim", Expected: "b", Actual: lr.RightTrim(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsLeftEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsLeftEmpty", func() {
		lr := corestr.NewLeftRight("", "b")
		tc := caseV1Compat{Name: "LR IsLeftEmpty", Expected: true, Actual: lr.IsLeftEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsRightEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsRightEmpty", func() {
		lr := corestr.NewLeftRight("a", "")
		tc := caseV1Compat{Name: "LR IsRightEmpty", Expected: true, Actual: lr.IsRightEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsLeftWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsLeftWhitespace", func() {
		lr := corestr.NewLeftRight("  ", "b")
		tc := caseV1Compat{Name: "LR IsLeftWhitespace", Expected: true, Actual: lr.IsLeftWhitespace(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsRightWhitespace", func() {
		lr := corestr.NewLeftRight("a", "  ")
		tc := caseV1Compat{Name: "LR IsRightWhitespace", Expected: true, Actual: lr.IsRightWhitespace(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_HasValidNonEmptyLeft(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_HasValidNonEmptyLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR HasValidNonEmptyLeft", Expected: true, Actual: lr.HasValidNonEmptyLeft(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_HasValidNonEmptyRight(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_HasValidNonEmptyRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR HasValidNonEmptyRight", Expected: true, Actual: lr.HasValidNonEmptyRight(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_HasValidNonWhitespaceLeft(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_HasValidNonWhitespaceLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR HasValidNonWSLeft", Expected: true, Actual: lr.HasValidNonWhitespaceLeft(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_HasValidNonWhitespaceRight(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_HasValidNonWhitespaceRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR HasValidNonWSRight", Expected: true, Actual: lr.HasValidNonWhitespaceRight(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_HasSafeNonEmpty", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR HasSafeNonEmpty", Expected: true, Actual: lr.HasSafeNonEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_Is(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_Is", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR Is", Expected: true, Actual: lr.Is("a", "b"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsLeft(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR IsLeft", Expected: true, Actual: lr.IsLeft("a"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsRight(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR IsRight", Expected: true, Actual: lr.IsRight("b"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR IsEqual", Expected: true, Actual: lr1.IsEqual(lr2), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsEqual_Nil(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsEqual_Nil", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR IsEqual nil", Expected: false, Actual: lr.IsEqual(nil), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()
		tc := caseV1Compat{Name: "LR Clone", Expected: true, Actual: c.IsEqual(lr), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_NonPtr", func() {
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		tc := caseV1Compat{Name: "LR NonPtr", Expected: "a", Actual: np.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_Ptr", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR Ptr", Expected: true, Actual: lr.Ptr() == lr, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsLeftRegexMatch(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsLeftRegexMatch", func() {
		lr := corestr.NewLeftRight("abc123", "b")
		re := regexp.MustCompile(`\d+`)
		tc := caseV1Compat{Name: "LR IsLeftRegexMatch", Expected: true, Actual: lr.IsLeftRegexMatch(re), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsLeftRegexMatch_Nil(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsLeftRegexMatch_Nil", func() {
		lr := corestr.NewLeftRight("a", "b")
		tc := caseV1Compat{Name: "LR IsLeftRegexMatch nil", Expected: false, Actual: lr.IsLeftRegexMatch(nil), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_IsRightRegexMatch(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_IsRightRegexMatch", func() {
		lr := corestr.NewLeftRight("a", "xyz123")
		re := regexp.MustCompile(`\d+`)
		tc := caseV1Compat{Name: "LR IsRightRegexMatch", Expected: true, Actual: lr.IsRightRegexMatch(re), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_Clear(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_Clear", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
		tc := caseV1Compat{Name: "LR Clear", Expected: "", Actual: lr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Dispose()
		tc := caseV1Compat{Name: "LR Dispose", Expected: "", Actual: lr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_UsingSlice(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_UsingSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		tc := caseV1Compat{Name: "LR UsingSlice", Expected: true, Actual: lr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_UsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_UsingSlice_Single", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a"})
		tc := caseV1Compat{Name: "LR UsingSlice single", Expected: false, Actual: lr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_UsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice([]string{})
		tc := caseV1Compat{Name: "LR UsingSlice empty", Expected: false, Actual: lr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRight_TrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRight_TrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		tc := caseV1Compat{Name: "LR TrimmedUsingSlice", Expected: "a", Actual: lr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// LeftRightFromSplit
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		tc := caseV1Compat{Name: "LRFromSplit", Expected: "key", Actual: lr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRightFromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		tc := caseV1Compat{Name: "LRFromSplitTrimmed", Expected: "key", Actual: lr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRightFromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		tc := caseV1Compat{Name: "LRFromSplitFull", Expected: "b:c:d", Actual: lr.Right, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov53_LeftRightFromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		tc := caseV1Compat{Name: "LRFromSplitFullTrimmed", Expected: "a", Actual: lr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// LeftMiddleRight
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_LMR_New(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := caseV1Compat{Name: "LMR New", Expected: true, Actual: lmr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		tc := caseV1Compat{Name: "LMR Invalid", Expected: false, Actual: lmr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_InvalidNoMsg(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_InvalidNoMsg", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		tc := caseV1Compat{Name: "LMR InvalidNoMsg", Expected: false, Actual: lmr.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_LeftBytes(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_LeftBytes", func() {
		lmr := corestr.NewLeftMiddleRight("ab", "cd", "ef")
		tc := caseV1Compat{Name: "LMR LeftBytes", Expected: 2, Actual: len(lmr.LeftBytes()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_MiddleBytes(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_MiddleBytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "bc", "d")
		tc := caseV1Compat{Name: "LMR MiddleBytes", Expected: 2, Actual: len(lmr.MiddleBytes()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_RightBytes(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_RightBytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "cd")
		tc := caseV1Compat{Name: "LMR RightBytes", Expected: 2, Actual: len(lmr.RightBytes()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Trims(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Trims", func() {
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
		tc := caseV1Compat{Name: "LMR LeftTrim", Expected: "a", Actual: lmr.LeftTrim(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "LMR MiddleTrim", Expected: "b", Actual: lmr.MiddleTrim(), Args: args.Map{}}
		tc2.ShouldBeEqual(t)
		tc3 := caseV1Compat{Name: "LMR RightTrim", Expected: "c", Actual: lmr.RightTrim(), Args: args.Map{}}
		tc3.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_IsEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("", "b", "")
		tc := caseV1Compat{Name: "LMR IsLeftEmpty", Expected: true, Actual: lmr.IsLeftEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "LMR IsMiddleEmpty", Expected: false, Actual: lmr.IsMiddleEmpty(), Args: args.Map{}}
		tc2.ShouldBeEqual(t)
		tc3 := caseV1Compat{Name: "LMR IsRightEmpty", Expected: true, Actual: lmr.IsRightEmpty(), Args: args.Map{}}
		tc3.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Whitespace(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Whitespace", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")
		tc := caseV1Compat{Name: "LMR IsLeftWS", Expected: true, Actual: lmr.IsLeftWhitespace(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "LMR IsMiddleWS", Expected: true, Actual: lmr.IsMiddleWhitespace(), Args: args.Map{}}
		tc2.ShouldBeEqual(t)
		tc3 := caseV1Compat{Name: "LMR IsRightWS", Expected: true, Actual: lmr.IsRightWhitespace(), Args: args.Map{}}
		tc3.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_HasValidNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := caseV1Compat{Name: "LMR HasValidNonEmptyLeft", Expected: true, Actual: lmr.HasValidNonEmptyLeft(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "LMR HasValidNonEmptyMiddle", Expected: true, Actual: lmr.HasValidNonEmptyMiddle(), Args: args.Map{}}
		tc2.ShouldBeEqual(t)
		tc3 := caseV1Compat{Name: "LMR HasValidNonEmptyRight", Expected: true, Actual: lmr.HasValidNonEmptyRight(), Args: args.Map{}}
		tc3.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_HasValidNonWS(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_HasValidNonWS", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := caseV1Compat{Name: "LMR HasValidNonWSLeft", Expected: true, Actual: lmr.HasValidNonWhitespaceLeft(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "LMR HasValidNonWSMiddle", Expected: true, Actual: lmr.HasValidNonWhitespaceMiddle(), Args: args.Map{}}
		tc2.ShouldBeEqual(t)
		tc3 := caseV1Compat{Name: "LMR HasValidNonWSRight", Expected: true, Actual: lmr.HasValidNonWhitespaceRight(), Args: args.Map{}}
		tc3.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_HasSafeNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := caseV1Compat{Name: "LMR HasSafeNonEmpty", Expected: true, Actual: lmr.HasSafeNonEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_IsAll(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := caseV1Compat{Name: "LMR IsAll", Expected: true, Actual: lmr.IsAll("a", "b", "c"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Is(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Is", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		tc := caseV1Compat{Name: "LMR Is", Expected: true, Actual: lmr.Is("a", "c"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Clone(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()
		tc := caseV1Compat{Name: "LMR Clone", Expected: true, Actual: c.IsAll("a", "b", "c"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		tc := caseV1Compat{Name: "LMR ToLeftRight", Expected: true, Actual: lr.Is("a", "c"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Clear(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Clear", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()
		tc := caseV1Compat{Name: "LMR Clear", Expected: "", Actual: lmr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMR_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov53_LMR_Dispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()
		tc := caseV1Compat{Name: "LMR Dispose", Expected: "", Actual: lmr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_LMRFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov53_LMRFromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		tc := caseV1Compat{Name: "LMRFromSplit", Expected: "b", Actual: lmr.Middle, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMRFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov53_LMRFromSplitTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		tc := caseV1Compat{Name: "LMRFromSplitTrimmed", Expected: "b", Actual: lmr.Middle, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMRFromSplitN(t *testing.T) {
	safeTest(t, "Test_Cov53_LMRFromSplitN", func() {
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		tc := caseV1Compat{Name: "LMRFromSplitN right", Expected: "c:d:e", Actual: lmr.Right, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_LMRFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov53_LMRFromSplitNTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
		tc := caseV1Compat{Name: "LMRFromSplitNTrimmed", Expected: "a", Actual: lmr.Left, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// ValidValue
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_New", func() {
		vv := corestr.NewValidValue("hello")
		tc := caseV1Compat{Name: "VV New", Expected: true, Actual: vv.IsValid && vv.Value == "hello", Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_NewEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_NewEmpty", func() {
		vv := corestr.NewValidValueEmpty()
		tc := caseV1Compat{Name: "VV NewEmpty", Expected: true, Actual: vv.IsValid && vv.Value == "", Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Invalid", func() {
		vv := corestr.InvalidValidValue("err")
		tc := caseV1Compat{Name: "VV Invalid", Expected: false, Actual: vv.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_InvalidNoMsg(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_InvalidNoMsg", func() {
		vv := corestr.InvalidValidValueNoMessage()
		tc := caseV1Compat{Name: "VV InvalidNoMsg", Expected: false, Actual: vv.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_IsEmpty", func() {
		vv := corestr.NewValidValue("")
		tc := caseV1Compat{Name: "VV IsEmpty", Expected: true, Actual: vv.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_IsWhitespace", func() {
		vv := corestr.NewValidValue("  ")
		tc := caseV1Compat{Name: "VV IsWhitespace", Expected: true, Actual: vv.IsWhitespace(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Trim", func() {
		vv := corestr.NewValidValue(" abc ")
		tc := caseV1Compat{Name: "VV Trim", Expected: "abc", Actual: vv.Trim(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		tc := caseV1Compat{Name: "VV HasValidNonEmpty", Expected: true, Actual: vv.HasValidNonEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_HasValidNonWhitespace", func() {
		vv := corestr.NewValidValue("x")
		tc := caseV1Compat{Name: "VV HasValidNonWS", Expected: true, Actual: vv.HasValidNonWhitespace(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_HasSafeNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		tc := caseV1Compat{Name: "VV HasSafeNonEmpty", Expected: true, Actual: vv.HasSafeNonEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		tc := caseV1Compat{Name: "VV ValueBool", Expected: true, Actual: vv.ValueBool(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_ValueInt", func() {
		vv := corestr.NewValidValue("42")
		tc := caseV1Compat{Name: "VV ValueInt", Expected: 42, Actual: vv.ValueInt(0), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_ValueDefInt", func() {
		vv := corestr.NewValidValue("10")
		tc := caseV1Compat{Name: "VV ValueDefInt", Expected: 10, Actual: vv.ValueDefInt(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_ValueFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		tc := caseV1Compat{Name: "VV ValueFloat64", Expected: 3.14, Actual: vv.ValueFloat64(0), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Is", func() {
		vv := corestr.NewValidValue("x")
		tc := caseV1Compat{Name: "VV Is", Expected: true, Actual: vv.Is("x"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		tc := caseV1Compat{Name: "VV IsAnyOf", Expected: true, Actual: vv.IsAnyOf("a", "b", "c"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		tc := caseV1Compat{Name: "VV IsContains", Expected: true, Actual: vv.IsContains("world"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		tc := caseV1Compat{Name: "VV IsEqualNonSensitive", Expected: true, Actual: vv.IsEqualNonSensitive("hello"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_IsRegexMatches", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		tc := caseV1Compat{Name: "VV IsRegexMatches", Expected: true, Actual: vv.IsRegexMatches(re), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		tc := caseV1Compat{Name: "VV Split", Expected: 3, Actual: len(vv.Split(",")), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("x")
		c := vv.Clone()
		tc := caseV1Compat{Name: "VV Clone", Expected: "x", Actual: c.Value, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_String", func() {
		vv := corestr.NewValidValue("x")
		tc := caseV1Compat{Name: "VV String", Expected: "x", Actual: vv.String(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("x")
		tc := caseV1Compat{Name: "VV FullString", Expected: true, Actual: len(vv.FullString()) > 0, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_ValueBytesOnce", func() {
		vv := corestr.NewValidValue("ab")
		tc := caseV1Compat{Name: "VV ValueBytesOnce", Expected: 2, Actual: len(vv.ValueBytesOnce()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Clear", func() {
		vv := corestr.NewValidValue("x")
		vv.Clear()
		tc := caseV1Compat{Name: "VV Clear", Expected: "", Actual: vv.Value, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("x")
		data, err := vv.Serialize()
		tc := caseV1Compat{Name: "VV Serialize", Expected: true, Actual: err == nil && len(data) > 0, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// ValidValues
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_ValidValues_New(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_New", func() {
		vvs := corestr.NewValidValues(5)
		tc := caseV1Compat{Name: "VVs New", Expected: true, Actual: vvs.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_Add", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		tc := caseV1Compat{Name: "VVs Add", Expected: 1, Actual: vvs.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_Count(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_Count", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		tc := caseV1Compat{Name: "VVs Count", Expected: 1, Actual: vvs.Count(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_HasAnyItem", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		tc := caseV1Compat{Name: "VVs HasAnyItem", Expected: true, Actual: vvs.HasAnyItem(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		tc := caseV1Compat{Name: "VVs SafeValueAt", Expected: "x", Actual: vvs.SafeValueAt(0), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_SafeValueAt_OOB(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_SafeValueAt_OOB", func() {
		vvs := corestr.NewValidValues(5)
		tc := caseV1Compat{Name: "VVs SafeValueAt oob", Expected: "", Actual: vvs.SafeValueAt(0), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		tc := caseV1Compat{Name: "VVs Strings", Expected: 2, Actual: len(vvs.Strings()), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_String", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		tc := caseV1Compat{Name: "VVs String", Expected: true, Actual: len(vvs.String()) > 0, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_Hashmap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "k", "v")
		hm := vvs.Hashmap()
		tc := caseV1Compat{Name: "VVs Hashmap", Expected: true, Actual: hm.Has("k"), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_Map", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "k", "v")
		m := vvs.Map()
		tc := caseV1Compat{Name: "VVs Map", Expected: "v", Actual: m["k"], Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov53_ValidValues_Empty", func() {
		vvs := corestr.EmptyValidValues()
		tc := caseV1Compat{Name: "VVs Empty", Expected: true, Actual: vvs.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov53_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("err")
		tc := caseV1Compat{Name: "VS Invalid", Expected: false, Actual: vs.ValueValid.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValueStatus_InvalidNoMsg(t *testing.T) {
	safeTest(t, "Test_Cov53_ValueStatus_InvalidNoMsg", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		tc := caseV1Compat{Name: "VS InvalidNoMsg", Expected: false, Actual: vs.ValueValid.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_Cov53_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("err")
		c := vs.Clone()
		tc := caseV1Compat{Name: "VS Clone", Expected: false, Actual: c.ValueValid.IsValid, Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════

func Test_Cov53_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov53_TextWithLineNumber_HasLineNumber", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		tc := caseV1Compat{Name: "TWL HasLineNumber", Expected: true, Actual: twl.HasLineNumber(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov53_TextWithLineNumber_IsInvalidLineNumber", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}
		tc := caseV1Compat{Name: "TWL IsInvalidLineNumber", Expected: true, Actual: twl.IsInvalidLineNumber(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_Cov53_TextWithLineNumber_Length", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		tc := caseV1Compat{Name: "TWL Length", Expected: 5, Actual: twl.Length(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Cov53_TextWithLineNumber_IsEmpty", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		tc := caseV1Compat{Name: "TWL IsEmpty", Expected: true, Actual: twl.IsEmpty(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_Cov53_TextWithLineNumber_IsEmptyText", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		tc := caseV1Compat{Name: "TWL IsEmptyText", Expected: true, Actual: twl.IsEmptyText(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}

func Test_Cov53_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_Cov53_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		twl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		tc := caseV1Compat{Name: "TWL IsEmptyTextLineBoth", Expected: true, Actual: twl.IsEmptyTextLineBoth(), Args: args.Map{}}
		tc.ShouldBeEqual(t)
	})
}
