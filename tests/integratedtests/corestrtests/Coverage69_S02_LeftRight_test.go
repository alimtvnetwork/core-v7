package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — Constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov69_InvalidLeftRightNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov69_InvalidLeftRightNoMessage", func() {
		// Arrange & Act
		lr := corestr.InvalidLeftRightNoMessage()

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": false,
			"left": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage returns invalid -- no message", actual)
	})
}

func Test_Cov69_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov69_InvalidLeftRight", func() {
		// Arrange & Act
		lr := corestr.InvalidLeftRight("test error")

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"hasMsg": lr.Message != "",
		}
		expected := args.Map{
			"isValid": false,
			"hasMsg": true,
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight returns invalid -- with message", actual)
	})
}

func Test_Cov69_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov69_NewLeftRight", func() {
		// Arrange & Act
		lr := corestr.NewLeftRight("a", "b")

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": true,
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "NewLeftRight returns valid -- two strings", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — LeftRightUsingSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov69_LeftRightUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlice_Empty", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlice([]string{})

		// Assert
		actual := args.Map{"isValid": lr.IsValid}
		expected := args.Map{"isValid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns invalid -- empty slice", actual)
	})
}

func Test_Cov69_LeftRightUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlice_One", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlice([]string{"only"})

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": false,
			"left": "only",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns invalid -- one element", actual)
	})
}

func Test_Cov69_LeftRightUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlice_Two", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": true,
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns valid -- two elements", actual)
	})
}

func Test_Cov69_LeftRightUsingSlice_Three(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlice_Three", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlice([]string{"a", "b", "c"})

		// Assert — length != 2 so invalid, takes first and last
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": false,
			"left": "a",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns invalid -- three elements", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — LeftRightUsingSlicePtr (deprecated)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov69_LeftRightUsingSlicePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlicePtr_Nil", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlicePtr(nil)

		// Assert
		actual := args.Map{"isValid": lr.IsValid}
		expected := args.Map{"isValid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns invalid -- nil", actual)
	})
}

func Test_Cov69_LeftRightUsingSlicePtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlicePtr_Empty", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlicePtr([]string{})

		// Assert
		actual := args.Map{"isValid": lr.IsValid}
		expected := args.Map{"isValid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns invalid -- empty", actual)
	})
}

func Test_Cov69_LeftRightUsingSlicePtr_Two(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightUsingSlicePtr_Two", func() {
		// Arrange & Act
		lr := corestr.LeftRightUsingSlicePtr([]string{"x", "y"})

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": true,
			"left": "x",
			"right": "y",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns valid -- two elements", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — LeftRightTrimmedUsingSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov69_LeftRightTrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightTrimmedUsingSlice_Nil", func() {
		// Arrange & Act
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Assert
		actual := args.Map{"isValid": lr.IsValid}
		expected := args.Map{"isValid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns invalid -- nil", actual)
	})
}

func Test_Cov69_LeftRightTrimmedUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightTrimmedUsingSlice_Empty", func() {
		// Arrange & Act
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Assert
		actual := args.Map{"isValid": lr.IsValid}
		expected := args.Map{"isValid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns invalid -- empty", actual)
	})
}

func Test_Cov69_LeftRightTrimmedUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightTrimmedUsingSlice_One", func() {
		// Arrange & Act
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" only "})

		// Assert — single element IS trimmed (source line 56 calls strings.TrimSpace)
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
		}
		expected := args.Map{
			"isValid": false,
			"left": "only",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns invalid -- one element", actual)
	})
}

func Test_Cov69_LeftRightTrimmedUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Cov69_LeftRightTrimmedUsingSlice_Two", func() {
		// Arrange & Act
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Assert
		actual := args.Map{
			"isValid": lr.IsValid,
			"left": lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{
			"isValid": true,
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns valid trimmed -- two elements", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — String-specific methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov69_LR_LeftBytes(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_LeftBytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc", "def")

		// Act
		actual := args.Map{
			"len": len(lr.LeftBytes()),
			"first": lr.LeftBytes()[0],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"first": byte('a'),
		}
		expected.ShouldBeEqual(t, 0, "LeftBytes returns bytes -- valid left", actual)
	})
}

func Test_Cov69_LR_RightBytes(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_RightBytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc", "def")

		// Act
		actual := args.Map{
			"len": len(lr.RightBytes()),
			"first": lr.RightBytes()[0],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"first": byte('d'),
		}
		expected.ShouldBeEqual(t, 0, "RightBytes returns bytes -- valid right", actual)
	})
}

func Test_Cov69_LR_LeftTrim(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_LeftTrim", func() {
		// Arrange
		lr := corestr.NewLeftRight(" x ", "y")

		// Act
		actual := args.Map{"trimmed": lr.LeftTrim()}

		// Assert
		expected := args.Map{"trimmed": "x"}
		expected.ShouldBeEqual(t, 0, "LeftTrim returns trimmed -- whitespace left", actual)
	})
}

func Test_Cov69_LR_RightTrim(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_RightTrim", func() {
		// Arrange
		lr := corestr.NewLeftRight("x", " y ")

		// Act
		actual := args.Map{"trimmed": lr.RightTrim()}

		// Assert
		expected := args.Map{"trimmed": "y"}
		expected.ShouldBeEqual(t, 0, "RightTrim returns trimmed -- whitespace right", actual)
	})
}

func Test_Cov69_LR_IsLeftEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsLeftEmpty_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "b")

		// Act
		actual := args.Map{"empty": lr.IsLeftEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsLeftEmpty returns true -- empty left", actual)
	})
}

func Test_Cov69_LR_IsRightEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsRightEmpty_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "")

		// Act
		actual := args.Map{"empty": lr.IsRightEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsRightEmpty returns true -- empty right", actual)
	})
}

func Test_Cov69_LR_IsLeftWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsLeftWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("  ", "b")

		// Act
		actual := args.Map{"ws": lr.IsLeftWhitespace()}

		// Assert
		expected := args.Map{"ws": true}
		expected.ShouldBeEqual(t, 0, "IsLeftWhitespace returns true -- whitespace left", actual)
	})
}

func Test_Cov69_LR_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsRightWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "  ")

		// Act
		actual := args.Map{"ws": lr.IsRightWhitespace()}

		// Assert
		expected := args.Map{"ws": true}
		expected.ShouldBeEqual(t, 0, "IsRightWhitespace returns true -- whitespace right", actual)
	})
}

func Test_Cov69_LR_HasValidNonEmptyLeft_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasValidNonEmptyLeft_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonEmptyLeft()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyLeft returns true -- valid non-empty", actual)
	})
}

func Test_Cov69_LR_HasValidNonEmptyLeft_EmptyLeft(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasValidNonEmptyLeft_EmptyLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonEmptyLeft()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyLeft returns false -- empty left", actual)
	})
}

func Test_Cov69_LR_HasValidNonEmptyRight_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasValidNonEmptyRight_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonEmptyRight()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyRight returns true -- valid non-empty", actual)
	})
}

func Test_Cov69_LR_HasValidNonWhitespaceLeft_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasValidNonWhitespaceLeft_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonWhitespaceLeft()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceLeft returns true -- non-ws", actual)
	})
}

func Test_Cov69_LR_HasValidNonWhitespaceLeft_Ws(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasValidNonWhitespaceLeft_Ws", func() {
		// Arrange
		lr := corestr.NewLeftRight("  ", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonWhitespaceLeft()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceLeft returns false -- whitespace left", actual)
	})
}

func Test_Cov69_LR_HasValidNonWhitespaceRight_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasValidNonWhitespaceRight_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonWhitespaceRight()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceRight returns true -- non-ws", actual)
	})
}

func Test_Cov69_LR_HasSafeNonEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasSafeNonEmpty_True", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns true -- both non-empty", actual)
	})
}

func Test_Cov69_LR_HasSafeNonEmpty_EmptyRight(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_HasSafeNonEmpty_EmptyRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "")

		// Act
		actual := args.Map{"result": lr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns false -- empty right", actual)
	})
}

func Test_Cov69_LR_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_NonPtr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		v := lr.NonPtr()

		// Act
		actual := args.Map{
			"left": v.Left,
			"right": v.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "NonPtr returns copy -- valid LR", actual)
	})
}

func Test_Cov69_LR_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Ptr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		p := lr.Ptr()

		// Act
		actual := args.Map{"same": p == lr}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Ptr returns self -- valid LR", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — Regex, Is, IsEqual, Clone, Clear, Dispose
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov69_LR_IsLeftRegexMatch_Match(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsLeftRegexMatch_Match", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "x")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"match": lr.IsLeftRegexMatch(re)}

		// Assert
		expected := args.Map{"match": true}
		expected.ShouldBeEqual(t, 0, "IsLeftRegexMatch returns true -- digits in left", actual)
	})
}

func Test_Cov69_LR_IsLeftRegexMatch_Nil(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsLeftRegexMatch_Nil", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc", "x")

		// Act
		actual := args.Map{"match": lr.IsLeftRegexMatch(nil)}

		// Assert
		expected := args.Map{"match": false}
		expected.ShouldBeEqual(t, 0, "IsLeftRegexMatch returns false -- nil regex", actual)
	})
}

func Test_Cov69_LR_IsRightRegexMatch_Match(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsRightRegexMatch_Match", func() {
		// Arrange
		lr := corestr.NewLeftRight("x", "abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"match": lr.IsRightRegexMatch(re)}

		// Assert
		expected := args.Map{"match": true}
		expected.ShouldBeEqual(t, 0, "IsRightRegexMatch returns true -- digits in right", actual)
	})
}

func Test_Cov69_LR_IsRightRegexMatch_Nil(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsRightRegexMatch_Nil", func() {
		// Arrange
		lr := corestr.NewLeftRight("x", "abc")

		// Act
		actual := args.Map{"match": lr.IsRightRegexMatch(nil)}

		// Assert
		expected := args.Map{"match": false}
		expected.ShouldBeEqual(t, 0, "IsRightRegexMatch returns false -- nil regex", actual)
	})
}

func Test_Cov69_LR_IsLeft(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello", "world")

		// Act
		actual := args.Map{
			"yes": lr.IsLeft("hello"),
			"no": lr.IsLeft("x"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsLeft returns correct -- match and mismatch", actual)
	})
}

func Test_Cov69_LR_IsRight(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("hello", "world")

		// Act
		actual := args.Map{
			"yes": lr.IsRight("world"),
			"no": lr.IsRight("x"),
		}

		// Assert
		expected := args.Map{
			"yes": true,
			"no": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRight returns correct -- match and mismatch", actual)
	})
}

func Test_Cov69_LR_Is(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Is", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"match": lr.Is("a", "b"),
			"noMatch": lr.Is("a", "x"),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"noMatch": false,
		}
		expected.ShouldBeEqual(t, 0, "Is returns correct -- both match and mismatch", actual)
	})
}

func Test_Cov69_LR_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsEqual_BothNil", func() {
		// Arrange
		var a, b *corestr.LeftRight

		// Act
		actual := args.Map{"equal": a.IsEqual(b)}

		// Assert
		expected := args.Map{"equal": true}
		expected.ShouldBeEqual(t, 0, "IsEqual returns true -- both nil", actual)
	})
}

func Test_Cov69_LR_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsEqual_OneNil", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"equal": lr.IsEqual(nil)}

		// Assert
		expected := args.Map{"equal": false}
		expected.ShouldBeEqual(t, 0, "IsEqual returns false -- one nil", actual)
	})
}

func Test_Cov69_LR_IsEqual_Match(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsEqual_Match", func() {
		// Arrange
		a := corestr.NewLeftRight("a", "b")
		b := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"equal": a.IsEqual(b)}

		// Assert
		expected := args.Map{"equal": true}
		expected.ShouldBeEqual(t, 0, "IsEqual returns true -- same values", actual)
	})
}

func Test_Cov69_LR_IsEqual_Mismatch(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_IsEqual_Mismatch", func() {
		// Arrange
		a := corestr.NewLeftRight("a", "b")
		b := corestr.NewLeftRight("a", "c")

		// Act
		actual := args.Map{"equal": a.IsEqual(b)}

		// Assert
		expected := args.Map{"equal": false}
		expected.ShouldBeEqual(t, 0, "IsEqual returns false -- different right", actual)
	})
}

func Test_Cov69_LR_Clone(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"right": c.Right,
			"notSame": c != lr,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"notSame": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone returns copy -- valid LR", actual)
	})
}

func Test_Cov69_LR_Clear(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Clear", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "Clear zeroes fields -- valid LR", actual)
	})
}

func Test_Cov69_LR_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Clear_Nil", func() {
		// Arrange
		var lr *corestr.LeftRight
		lr.Clear() // must not panic

		// Act
		actual := args.Map{"noPanic": true}

		// Assert
		expected := args.Map{"noPanic": true}
		expected.ShouldBeEqual(t, 0, "Clear returns safely -- nil receiver", actual)
	})
}

func Test_Cov69_LR_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Dispose", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Dispose()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "Dispose clears fields -- valid LR", actual)
	})
}

func Test_Cov69_LR_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov69_LR_Dispose_Nil", func() {
		// Arrange
		var lr *corestr.LeftRight
		lr.Dispose() // must not panic

		// Act
		actual := args.Map{"noPanic": true}

		// Assert
		expected := args.Map{"noPanic": true}
		expected.ShouldBeEqual(t, 0, "Dispose returns safely -- nil receiver", actual)
	})
}
