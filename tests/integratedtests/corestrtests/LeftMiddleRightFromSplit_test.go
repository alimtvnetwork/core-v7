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
		// Arrange
		tc := leftMiddleRightFromSplitNormalTestCase
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_TwoParts(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_TwoParts", func() {
		// Arrange
		tc := leftMiddleRightFromSplitTwoPartsTestCase
		lmr := corestr.LeftMiddleRightFromSplit("a.b", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_SinglePart(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_SinglePart", func() {
		// Arrange
		tc := leftMiddleRightFromSplitSinglePartTestCase
		lmr := corestr.LeftMiddleRightFromSplit("hello", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_FourPlus(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_FourPlus", func() {
		// Arrange
		tc := leftMiddleRightFromSplitFourPlusTestCase
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c.d", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_Empty(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_Empty", func() {
		// Arrange
		tc := leftMiddleRightFromSplitEmptyTestCase
		lmr := corestr.LeftMiddleRightFromSplit("", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplit_Edges(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplit_Edges", func() {
		// Arrange
		tc := leftMiddleRightFromSplitEdgesTestCase
		lmr := corestr.LeftMiddleRightFromSplit("..", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitTrimmed — trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitTrimmed_All(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitTrimmed_All", func() {
		// Arrange
		tc := leftMiddleRightFromSplitTrimmedAllTestCase
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  .  c  ", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitTrimmed_Two(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitTrimmed_Two", func() {
		// Arrange
		tc := leftMiddleRightFromSplitTrimmedTwoTestCase
		lmr := corestr.LeftMiddleRightFromSplitTrimmed("  a  .  b  ", ".")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

func Test_LeftMiddleRightFromSplitN_Remainder(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_Remainder", func() {
		// Arrange
		tc := leftMiddleRightFromSplitNRemainderTestCase
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitN_Exact3(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_Exact3", func() {
		// Arrange
		tc := leftMiddleRightFromSplitNExact3TestCase
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c", ":")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitN_TwoOnly(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_TwoOnly", func() {
		// Arrange
		tc := leftMiddleRightFromSplitNTwoOnlyTestCase
		lmr := corestr.LeftMiddleRightFromSplitN("a:b", ":")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitN_MissingSep(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitN_MissingSep", func() {
		// Arrange
		tc := leftMiddleRightFromSplitNMissingSepTestCase
		lmr := corestr.LeftMiddleRightFromSplitN("nosep", ":")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

// ==========================================================================
// Test: LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftMiddleRightFromSplitNTrimmed_Remainder(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitNTrimmed_Remainder", func() {
		// Arrange
		tc := leftMiddleRightFromSplitNTrimmedRemainderTestCase
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d : e ", ":")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_LeftMiddleRightFromSplitNTrimmed_Two(t *testing.T) {
	safeTest(t, "Test_LeftMiddleRightFromSplitNTrimmed_Two", func() {
		// Arrange
		tc := leftMiddleRightFromSplitNTrimmedTwoTestCase
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b ", ":")

		// Act
		actual := args.Map{
			"left":    lmr.Left,
			"middle":  lmr.Middle,
			"right":   lmr.Right,
			"isValid": fmt.Sprintf("%v", lmr.IsValid),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
