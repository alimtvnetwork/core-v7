package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRightFromSplit_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftRightFromSplit_Basic", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"isValid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns pair -- key=value", actual)
	})
}

func Test_LeftRightFromSplit_NoSep(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftRightFromSplit_NoSep", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("nosep", "=")

		// Act
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
		}

		// Assert
		expected := args.Map{
			"isValid": false,
			"left": "nosep",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns invalid -- no separator", actual)
	})
}

func Test_LeftRightFromSplitTrimmed_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftRightFromSplitTrimmed_Basic", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"isValid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed returns trimmed pair -- with spaces", actual)
	})
}

func Test_LeftRightFromSplitFull_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftRightFromSplitFull_Basic", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b:c:d",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull returns first split -- colon separated", actual)
	})
}

func Test_LeftRightFromSplitFullTrimmed_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftRightFromSplitFullTrimmed_Basic", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b : c",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed returns trimmed -- with spaces", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftMiddleRightFromSplit_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftMiddleRightFromSplit_Basic", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
			"isValid": lmr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns triple -- dot separated", actual)
	})
}

func Test_LeftMiddleRightFromSplit_TwoParts(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftMiddleRightFromSplit_TwoParts", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a.b", ".")

		// Act
		actual := args.Map{"isValid": lmr.IsValid}

		// Assert
		expected := args.Map{"isValid": false}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns invalid -- only two parts", actual)
	})
}

func Test_LeftMiddleRightFromSplitTrimmed_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftMiddleRightFromSplitTrimmed_Basic", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed returns trimmed -- with spaces", actual)
	})
}

func Test_LeftMiddleRightFromSplitN_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftMiddleRightFromSplitN_Basic", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c:d:e",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN returns triple -- remainder in right", actual)
	})
}

func Test_LeftMiddleRightFromSplitNTrimmed_Basic(t *testing.T) {
	safeTest(t, "Test_Cov71_LeftMiddleRightFromSplitNTrimmed_Basic", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"middle": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c : d",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed returns trimmed -- with spaces", actual)
	})
}
