package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: LeftMiddleRightFromSplit — edge cases
// ==========================================================================

func Test_LeftMiddleRightFromSplit_Normal(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_Normal", func() {
		tc := leftMiddleRightFromSplitNormalTestCase
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_TwoParts(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_TwoParts", func() {
		tc := leftMiddleRightFromSplitTwoPartsTestCase
		lmr := corestr.LeftMiddleRightFromSplit("a.b", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_SinglePart(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_SinglePart", func() {
		tc := leftMiddleRightFromSplitSinglePartTestCase
		lmr := corestr.LeftMiddleRightFromSplit("hello", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_FourPlus(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_FourPlus", func() {
		tc := leftMiddleRightFromSplitFourPlusTestCase
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c.d", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_Empty(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_Empty", func() {
		tc := leftMiddleRightFromSplitEmptyTestCase
		lmr := corestr.LeftMiddleRightFromSplit("", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_Edges(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_Edges", func() {
		tc := leftMiddleRightFromSplitEdgesTestCase
		lmr := corestr.LeftMiddleRightFromSplit("..", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitTrimmed — trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitTrimmed_All(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitTrimmed_All", func() {
		tc := leftMiddleRightFromSplitTrimmedAllTestCase
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  .  c  ", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitTrimmed_Two(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitTrimmed_Two", func() {
		tc := leftMiddleRightFromSplitTrimmedTwoTestCase
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  ", ".")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

func Test_LeftMiddleRightFromSplitN_Remainder(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_Remainder", func() {
		tc := leftMiddleRightFromSplitNRemainderTestCase
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitN_Exact3(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_Exact3", func() {
		tc := leftMiddleRightFromSplitNExact3TestCase
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitN_TwoOnly(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_TwoOnly", func() {
		tc := leftMiddleRightFromSplitNTwoOnlyTestCase
		lmr := corestr.LeftMiddleRightFromSplitN("a:b", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitN_MissingSep(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_MissingSep", func() {
		tc := leftMiddleRightFromSplitNMissingSepTestCase
		lmr := corestr.LeftMiddleRightFromSplitN("nosep", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitNTrimmed_Remainder(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitNTrimmed_Remainder", func() {
		tc := leftMiddleRightFromSplitNTrimmedRemainderTestCase
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d : e ", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitNTrimmed_Two(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitNTrimmed_Two", func() {
		tc := leftMiddleRightFromSplitNTrimmedTwoTestCase
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b ", ":")
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
