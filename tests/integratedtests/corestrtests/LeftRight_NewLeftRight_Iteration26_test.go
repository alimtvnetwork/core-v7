package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_NewLeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("key", "value")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftRight returns correct value -- with args", actual)
	})
}

func Test_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_InvalidNoMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage returns error -- with args", actual)
	})
}

func Test_LeftRight_InvalidWithMessage(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_InvalidWithMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("bad")

		// Act
		actual := args.Map{
			"valid": lr.IsValid,
			"msg": lr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "bad",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight returns error -- with args", actual)
	})
}

func Test_LeftRight_UsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlice_Two", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- two", actual)
	})
}

func Test_LeftRight_UsingSlice_One(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlice_One", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"only"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "only",
			"right": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- one", actual)
	})
}

func Test_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns empty -- empty", actual)
	})
}

func Test_LeftRight_UsingSlicePtr_Empty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlicePtr_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns empty -- empty", actual)
	})
}

func Test_LeftRight_UsingSlicePtr_NonEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlicePtr_NonEmpty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns empty -- non-empty", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_Two", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- two", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_Nil", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns nil -- nil", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns empty -- empty", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_One", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" only "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "only",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- one", actual)
	})
}

func Test_LeftRight_Bytes(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Bytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc", "xyz")

		// Act
		actual := args.Map{
			"leftLen": len(lr.LeftBytes()),
			"rightLen": len(lr.RightBytes()),
		}

		// Assert
		expected := args.Map{
			"leftLen": 3,
			"rightLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Bytes", actual)
	})
}

func Test_LeftRight_Trim(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Trim", func() {
		// Arrange
		lr := corestr.NewLeftRight(" a ", " b ")

		// Act
		actual := args.Map{
			"left": lr.LeftTrim(),
			"right": lr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Trim", actual)
	})
}

func Test_LeftRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_IsEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "x")

		// Act
		actual := args.Map{
			"leftEmpty": lr.IsLeftEmpty(),
			"rightEmpty": lr.IsRightEmpty(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty": true,
			"rightEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns empty -- IsEmpty", actual)
	})
}

func Test_LeftRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_IsWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("  ", "x")

		// Act
		actual := args.Map{
			"leftWs": lr.IsLeftWhitespace(),
			"rightWs": lr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftWs": true,
			"rightWs": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- IsWhitespace", actual)
	})
}

func Test_LeftRight_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_HasValidNonEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"left":  lr.HasValidNonEmptyLeft(),
			"right": lr.HasValidNonEmptyRight(),
			"safe":  lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"left": true,
			"right": true,
			"safe": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns empty -- HasValidNonEmpty", actual)
	})
}

func Test_LeftRight_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_HasValidNonWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", " ")

		// Act
		actual := args.Map{
			"leftNWS":  lr.HasValidNonWhitespaceLeft(),
			"rightNWS": lr.HasValidNonWhitespaceRight(),
		}

		// Assert
		expected := args.Map{
			"leftNWS": true,
			"rightNWS": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns non-empty -- HasValidNonWhitespace", actual)
	})
}

func Test_LeftRight_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_NonPtr_Ptr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		p := lr.Ptr()

		// Act
		actual := args.Map{
			"npLeft": np.Left,
			"pSame": p == lr,
		}

		// Assert
		expected := args.Map{
			"npLeft": "a",
			"pSame": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- NonPtr/Ptr", actual)
	})
}

func Test_LeftRight_RegexMatch(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_RegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "xyz")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"leftMatch":  lr.IsLeftRegexMatch(re),
			"rightMatch": lr.IsRightRegexMatch(re),
			"nilRegex":   lr.IsLeftRegexMatch(nil),
		}

		// Assert
		expected := args.Map{
			"leftMatch": true,
			"rightMatch": false,
			"nilRegex": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- RegexMatch", actual)
	})
}

func Test_LeftRight_Is_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Is_IsKey_IsVal", func() {
		// Arrange
		lr := corestr.NewLeftRight("k", "v")

		// Act
		actual := args.Map{
			"is":     lr.Is("k", "v"),
			"isNot":  lr.Is("k", "x"),
			"isLeft": lr.IsLeft("k"),
			"isRight": lr.IsRight("v"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"isNot": false,
			"isLeft": true,
			"isRight": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Is/IsLeft/IsRight", actual)
	})
}

func Test_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("a", "c")

		// Act
		actual := args.Map{
			"same":    lr1.IsEqual(lr2),
			"diff":    lr1.IsEqual(lr3),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
			"nilOne":  lr1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"same": true,
			"diff": false,
			"nilBoth": true,
			"nilOne": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- IsEqual", actual)
	})
}

func Test_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()

		// Act
		actual := args.Map{
			"left": cloned.Left,
			"right": cloned.Right,
			"notSame": cloned != lr,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"notSame": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clone", actual)
	})
}

func Test_LeftRight_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Clear_Dispose", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clear", actual)

		lr2 := corestr.NewLeftRight("x", "y")
		lr2.Dispose()
		actual2 := args.Map{"left": lr2.Left}
		expected2 := args.Map{"left": ""}
		expected2.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Dispose", actual2)

		// nil paths
		(*corestr.LeftRight)(nil).Clear()
		(*corestr.LeftRight)(nil).Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplitTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplitFull", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull returns correct value -- with args", actual)
	})
}

func Test_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplitFullTrimmed", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_New", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
			"valid": lmr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Invalid", func() {
		// Arrange
		lmr1 := corestr.InvalidLeftMiddleRightNoMessage()
		lmr2 := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{
			"v1": lmr1.IsValid,
			"v2": lmr2.IsValid,
			"msg": lmr2.Message,
		}

		// Assert
		expected := args.Map{
			"v1": false,
			"v2": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight returns error -- with args", actual)
	})
}

func Test_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Bytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("ab", "cd", "ef")

		// Act
		actual := args.Map{
			"lLen": len(lmr.LeftBytes()),
			"mLen": len(lmr.MiddleBytes()),
			"rLen": len(lmr.RightBytes()),
		}

		// Assert
		expected := args.Map{
			"lLen": 2,
			"mLen": 2,
			"rLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Bytes", actual)
	})
}

func Test_LeftMiddleRight_Trim(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Trim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")

		// Act
		actual := args.Map{
			"l": lmr.LeftTrim(),
			"m": lmr.MiddleTrim(),
			"r": lmr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"l": "a",
			"m": "b",
			"r": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Trim", actual)
	})
}

func Test_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_IsEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "x", "")

		// Act
		actual := args.Map{
			"lEmpty": lmr.IsLeftEmpty(),
			"mEmpty": lmr.IsMiddleEmpty(),
			"rEmpty": lmr.IsRightEmpty(),
		}

		// Assert
		expected := args.Map{
			"lEmpty": true,
			"mEmpty": false,
			"rEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns empty -- IsEmpty", actual)
	})
}

func Test_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_IsWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("  ", "x", "  ")

		// Act
		actual := args.Map{
			"lWs": lmr.IsLeftWhitespace(),
			"mWs": lmr.IsMiddleWhitespace(),
			"rWs": lmr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"lWs": true,
			"mWs": false,
			"rWs": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- IsWhitespace", actual)
	})
}

func Test_LeftMiddleRight_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_HasValidNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"l": lmr.HasValidNonEmptyLeft(), "m": lmr.HasValidNonEmptyMiddle(), "r": lmr.HasValidNonEmptyRight(),
			"safe": lmr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"l": true,
			"m": true,
			"r": true,
			"safe": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns empty -- HasValidNonEmpty", actual)
	})
}

func Test_LeftMiddleRight_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_HasValidNonWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", " ", "c")

		// Act
		actual := args.Map{
			"l": lmr.HasValidNonWhitespaceLeft(), "m": lmr.HasValidNonWhitespaceMiddle(), "r": lmr.HasValidNonWhitespaceRight(),
		}

		// Assert
		expected := args.Map{
			"l": true,
			"m": false,
			"r": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns non-empty -- HasValidNonWhitespace", actual)
	})
}

func Test_LeftMiddleRight_IsAll_Is(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_IsAll_Is", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"isAll": lmr.IsAll("a", "b", "c"),
			"is": lmr.Is("a", "c"),
		}

		// Assert
		expected := args.Map{
			"isAll": true,
			"is": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- IsAll/Is", actual)
	})
}

func Test_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()

		// Act
		actual := args.Map{
			"left": cloned.Left,
			"notSame": cloned != lmr,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"notSame": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Clone", actual)
	})
}

func Test_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- ToLeftRight", actual)
	})
}

func Test_LeftMiddleRight_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Clear_Dispose", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()

		// Act
		actual := args.Map{"left": lmr.Left}

		// Assert
		expected := args.Map{"left": ""}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Clear", actual)

		lmr2 := corestr.NewLeftMiddleRight("x", "y", "z")
		lmr2.Dispose()
		(*corestr.LeftMiddleRight)(nil).Clear()
		(*corestr.LeftMiddleRight)(nil).Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplitTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplitN", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c:d:e",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplitNTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Basic", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "name", Value: "alice"}

		// Act
		actual := args.Map{
			"key": kv.KeyName(), "varName": kv.VariableName(), "val": kv.ValueString(),
			"isVarEq": kv.IsVariableNameEqual("name"), "isValEq": kv.IsValueEqual("alice"),
		}

		// Assert
		expected := args.Map{
			"key": "name",
			"varName": "name",
			"val": "alice",
			"isVarEq": true,
			"isValEq": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- basic", actual)
	})
}

func Test_KeyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Json", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		jr := kv.Json()

		// Act
		actual := args.Map{"noErr": !jr.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Json", actual)
	})
}

func Test_KeyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_JsonPtr", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"notNil": kv.JsonPtr() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- JsonPtr", actual)
	})
}

func Test_KeyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Serialize", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Serialize", actual)
	})
}

func Test_KeyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_SerializeMust", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- SerializeMust", actual)
	})
}

func Test_KeyValuePair_Compile_String(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Compile_String", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"compile": kv.Compile(),
			"str": kv.String(),
		}

		// Assert
		expected := args.Map{
			"compile": "{k:v}",
			"str": "{k:v}",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Compile/String", actual)
	})
}

func Test_KeyValuePair_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_EmptyChecks", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "", Value: ""}

		// Act
		actual := args.Map{
			"keyEmpty": kv.IsKeyEmpty(), "valEmpty": kv.IsValueEmpty(),
			"hasKey": kv.HasKey(), "hasVal": kv.HasValue(),
			"kvEmpty": kv.IsKeyValueEmpty(), "kvAnyEmpty": kv.IsKeyValueAnyEmpty(),
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true, "valEmpty": true,
			"hasKey": false, "hasVal": false,
			"kvEmpty": true, "kvAnyEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns empty -- EmptyChecks", actual)
	})
}

func Test_KeyValuePair_Trim(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Trim", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}

		// Act
		actual := args.Map{
			"key": kv.TrimKey(),
			"val": kv.TrimValue(),
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"val": "v",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Trim", actual)
	})
}

func Test_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueBool", func() {
		// Arrange
		kv1 := corestr.KeyValuePair{Value: "true"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		kv3 := corestr.KeyValuePair{Value: ""}

		// Act
		actual := args.Map{
			"t": kv1.ValueBool(),
			"f": kv2.ValueBool(),
			"empty": kv3.ValueBool(),
		}

		// Assert
		expected := args.Map{
			"t": true,
			"f": false,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueBool", actual)
	})
}

func Test_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueInt", func() {
		// Arrange
		kv1 := corestr.KeyValuePair{Value: "42"}
		kv2 := corestr.KeyValuePair{Value: "abc"}

		// Act
		actual := args.Map{
			"val": kv1.ValueInt(0),
			"def": kv2.ValueInt(99),
			"defInt": kv1.ValueDefInt(),
		}

		// Assert
		expected := args.Map{
			"val": 42,
			"def": 99,
			"defInt": 42,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueInt", actual)
	})
}

func Test_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueByte", func() {
		// Arrange
		kv1 := corestr.KeyValuePair{Value: "100"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		kv3 := corestr.KeyValuePair{Value: "300"}

		// Act
		actual := args.Map{
			"val": kv1.ValueByte(0),
			"def": kv2.ValueByte(7),
			"overflow": kv3.ValueByte(5),
		}

		// Assert
		expected := args.Map{
			"val": byte(100),
			"def": byte(7),
			"overflow": byte(5),
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueByte", actual)
	})
}

func Test_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueDefByte", func() {
		// Arrange
		kv1 := corestr.KeyValuePair{Value: "50"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		kv3 := corestr.KeyValuePair{Value: "999"}

		// Act
		actual := args.Map{
			"val": kv1.ValueDefByte(),
			"err": kv2.ValueDefByte(),
			"over": kv3.ValueDefByte(),
		}

		// Assert
		expected := args.Map{
			"val": byte(50),
			"err": byte(0),
			"over": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueDefByte", actual)
	})
}

func Test_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueFloat64", func() {
		// Arrange
		kv1 := corestr.KeyValuePair{Value: "3.14"}
		kv2 := corestr.KeyValuePair{Value: "abc"}

		// Act
		actual := args.Map{
			"close": kv1.ValueFloat64(0) > 3.1,
			"def": kv2.ValueFloat64(1.0),
			"defFloat": kv1.ValueDefFloat64() > 3.1,
		}

		// Assert
		expected := args.Map{
			"close": true,
			"def": 1.0,
			"defFloat": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueFloat64", actual)
	})
}

func Test_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueValid", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "test"}
		vv := kv.ValueValid()

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "test",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueValid", actual)
	})
}

func Test_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueValidOptions", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: "test"}
		vv := kv.ValueValidOptions(false, "bad")

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "bad",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueValidOptions", actual)
	})
}

func Test_KeyValuePair_Is_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Is_IsKey_IsVal", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"is": kv.Is("k", "v"),
			"isKey": kv.IsKey("k"),
			"isVal": kv.IsVal("v"),
		}

		// Assert
		expected := args.Map{
			"is": true,
			"isKey": true,
			"isVal": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Is/IsKey/IsVal", actual)
	})
}

func Test_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_FormatString", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "name", Value: "bob"}

		// Act
		actual := args.Map{"fmt": kv.FormatString("%s=%s")}

		// Assert
		expected := args.Map{"fmt": "name=bob"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- FormatString", actual)
	})
}

func Test_KeyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Clear_Dispose", func() {
		// Arrange
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()

		// Act
		actual := args.Map{
			"key": kv.Key,
			"val": kv.Value,
		}

		// Assert
		expected := args.Map{
			"key": "",
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Clear", actual)

		kv2 := &corestr.KeyValuePair{Key: "x", Value: "y"}
		kv2.Dispose()
		(*corestr.KeyValuePair)(nil).Clear()
		(*corestr.KeyValuePair)(nil).Dispose()
	})
}

func Test_KeyValuePair_NilChecks(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_NilChecks", func() {
		// Arrange
		var kv *corestr.KeyValuePair

		// Act
		actual := args.Map{"anyEmpty": kv.IsKeyValueAnyEmpty()}

		// Assert
		expected := args.Map{"anyEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns nil -- nil checks", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_Valid(t *testing.T) {
	safeTest(t, "Test_I26_TextWithLineNumber_Valid", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"hasLine": tl.HasLineNumber(), "invalid": tl.IsInvalidLineNumber(),
			"len": tl.Length(), "empty": tl.IsEmpty(), "emptyText": tl.IsEmptyText(),
			"emptyBoth": tl.IsEmptyTextLineBoth(),
		}

		// Assert
		expected := args.Map{
			"hasLine": true, "invalid": false,
			"len": 5, "empty": false, "emptyText": false,
			"emptyBoth": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- valid", actual)
	})
}

func Test_TextWithLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_I26_TextWithLineNumber_Nil", func() {
		// Arrange
		var tl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"hasLine": tl.HasLineNumber(), "invalid": tl.IsInvalidLineNumber(),
			"len": tl.Length(), "empty": tl.IsEmpty(), "emptyText": tl.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"hasLine": false, "invalid": true,
			"len": 0, "empty": true, "emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns nil -- nil", actual)
	})
}

func Test_TextWithLineNumber_EmptyText(t *testing.T) {
	safeTest(t, "Test_I26_TextWithLineNumber_EmptyText", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: ""}

		// Act
		actual := args.Map{
			"empty": tl.IsEmpty(),
			"emptyText": tl.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns empty -- empty text", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_CloneSlice_Normal(t *testing.T) {
	safeTest(t, "Test_I26_CloneSlice_Normal", func() {
		// Arrange
		s := []string{"a", "b", "c"}
		cloned := corestr.CloneSlice(s)

		// Act
		actual := args.Map{
			"len": len(cloned),
			"first": cloned[0],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- normal", actual)
	})
}

func Test_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I26_CloneSlice_Empty", func() {
		// Arrange
		cloned := corestr.CloneSlice([]string{})

		// Act
		actual := args.Map{"len": len(cloned)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- empty", actual)
	})
}

func Test_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_I26_CloneSliceIf_Clone", func() {
		// Arrange
		cloned := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{
			"len": len(cloned),
			"first": cloned[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- clone", actual)
	})
}

func Test_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_I26_CloneSliceIf_NoClone", func() {
		// Arrange
		result := corestr.CloneSliceIf(false, "a", "b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- no clone", actual)
	})
}

func Test_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_I26_CloneSliceIf_Empty", func() {
		// Arrange
		result := corestr.CloneSliceIf(true)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_I26_AnyToString_WithFieldName", func() {
		// Arrange
		result := corestr.AnyToString(true, 42)

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field name", actual)
	})
}

func Test_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_I26_AnyToString_WithoutFieldName", func() {
		// Arrange
		result := corestr.AnyToString(false, 42)

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- without field name", actual)
	})
}

func Test_AnyToString_EmptyString(t *testing.T) {
	safeTest(t, "Test_I26_AnyToString_EmptyString", func() {
		// Arrange
		result := corestr.AnyToString(false, "")

		// Act
		actual := args.Map{"empty": result}

		// Assert
		expected := args.Map{"empty": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty string", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength / AllIndividualsLengthOfSimpleSlices
// ══════════════════════════════════════════════════════════════════════════════

func Test_AllIndividualStringsOfStringsLength_Valid(t *testing.T) {
	safeTest(t, "Test_I26_AllIndividualStringsOfStringsLength_Valid", func() {
		// Arrange
		items := [][]string{{"a", "b"}, {"c"}}

		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&items)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns non-empty -- valid", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_I26_AllIndividualStringsOfStringsLength_Nil", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns nil -- nil", actual)
	})
}

func Test_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_I26_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Act
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices(nil)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns nil -- nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_I26_ValueStatus_Invalid", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("bad")

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"msg": vs.ValueValid.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "bad",
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns error -- invalid", actual)
	})
}

func Test_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_I26_ValueStatus_InvalidNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{"valid": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns empty -- invalid no msg", actual)
	})
}

func Test_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_I26_ValueStatus_Clone", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("test")
		cloned := vs.Clone()

		// Act
		actual := args.Map{
			"msg": cloned.ValueValid.Message,
			"notSame": cloned != vs,
		}

		// Assert
		expected := args.Map{
			"msg": "test",
			"notSame": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- Clone", actual)
	})
}
