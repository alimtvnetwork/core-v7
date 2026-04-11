package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: LeftRightFromSplit — edge cases
// ==========================================================================

func Test_LeftRightFromSplit_FromLeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_LeftRightFromSplit", func() {
		// Arrange
		// Case 0: Normal key=value split
		{
			tc := leftRightFromSplitNormalTestCase
			lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
			actual := args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: Missing separator
		{
			tc := leftRightFromSplitMissingSepTestCase
			lr := corestr.LeftRightFromSplit("no-separator-here", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: Empty input
		{
			tc := leftRightFromSplitEmptyTestCase
			lr := corestr.LeftRightFromSplit("", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 3: Separator at start
		{
			tc := leftRightFromSplitSepAtStartTestCase
			lr := corestr.LeftRightFromSplit("=value", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 4: Separator at end
		{
			tc := leftRightFromSplitSepAtEndTestCase
			lr := corestr.LeftRightFromSplit("key=", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 5: Multiple separators
		{
			tc := leftRightFromSplitMultipleSepTestCase
			lr := corestr.LeftRightFromSplit("a=b=c", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

func Test_LeftRightFromSplitTrimmed_FromLeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_LeftRightFromSplitTrimmed", func() {
		// Arrange
		// Case 0: Trims whitespace
		{
			tc := leftRightFromSplitTrimmedTrimsTestCase
			lr := corestr.LeftRightFromSplitTrimmed("  key  =  value  ", "=")

		// Act
			actual := args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: No separator
		{
			tc := leftRightFromSplitTrimmedNoSepTestCase
			lr := corestr.LeftRightFromSplitTrimmed("  hello  ", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: Whitespace-only parts
		{
			tc := leftRightFromSplitTrimmedWhitespaceTestCase
			lr := corestr.LeftRightFromSplitTrimmed("   =   ", "=")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: LeftRightFromSplitFull — remainder handling
// ==========================================================================

func Test_LeftRightFromSplitFull_FromLeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_LeftRightFromSplitFull", func() {
		// Arrange
		// Case 0: Remainder in right
		{
			tc := leftRightFromSplitFullRemainderTestCase
			lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Act
			actual := args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: Single separator
		{
			tc := leftRightFromSplitFullSingleSepTestCase
			lr := corestr.LeftRightFromSplitFull("key:value", ":")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: Missing separator
		{
			tc := leftRightFromSplitFullMissingSepTestCase
			lr := corestr.LeftRightFromSplitFull("nosep", ":")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

func Test_LeftRightFromSplitFullTrimmed_FromLeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_LeftRightFromSplitFullTrimmed", func() {
		// Arrange
		// Case 0: Remainder trimmed
		{
			tc := leftRightFromSplitFullTrimmedRemainderTestCase
			lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c : d ", ":")

		// Act
			actual := args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: Missing separator trimmed
		{
			tc := leftRightFromSplitFullTrimmedMissingSepTestCase
			lr := corestr.LeftRightFromSplitFullTrimmed("  hello  ", ":")
			actual = args.Map{
				"left":    lr.Left,
				"right":   lr.Right,
				"isValid": fmt.Sprintf("%v", lr.IsValid),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}
