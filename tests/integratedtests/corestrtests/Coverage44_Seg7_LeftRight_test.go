package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight — Segment 7a
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LR_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_NewLeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("left", "right")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "left",
			"right": "right",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftRight -- valid pair", actual)
	})
}

func Test_Seg7_LR_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_InvalidLeftRight", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("err")

		// Act
		actual := args.Map{
			"valid": lr.IsValid,
			"msg": lr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight -- invalid with message", actual)
	})
}

func Test_Seg7_LR_InvalidLeftRightNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_InvalidLeftRightNoMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage -- invalid", actual)
	})
}

func Test_Seg7_LR_LeftRightUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightUsingSlice_Two", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice 2 -- valid", actual)
	})
}

func Test_Seg7_LR_LeftRightUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightUsingSlice_One", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice 1 -- invalid", actual)
	})
}

func Test_Seg7_LR_LeftRightUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightUsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice empty -- invalid", actual)
	})
}

func Test_Seg7_LR_LeftRightUsingSlice_Three(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightUsingSlice_Three", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b", "c"})

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
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice 3 -- invalid takes last", actual)
	})
}

func Test_Seg7_LR_LeftRightUsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightUsingSlicePtr", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr -- delegates", actual)
	})
}

func Test_Seg7_LR_LeftRightUsingSlicePtr_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightUsingSlicePtr_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr empty -- invalid", actual)
	})
}

func Test_Seg7_LR_LeftRightTrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightTrimmedUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

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
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice -- trimmed", actual)
	})
}

func Test_Seg7_LR_LeftRightTrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightTrimmedUsingSlice_Nil", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice nil -- invalid", actual)
	})
}

func Test_Seg7_LR_LeftRightTrimmedUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightTrimmedUsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice empty -- invalid", actual)
	})
}

func Test_Seg7_LR_LeftRightTrimmedUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_LeftRightTrimmedUsingSlice_One", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{"a"})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice one -- invalid", actual)
	})
}

func Test_Seg7_LR_Bytes(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Bytes", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftBytes/RightBytes -- correct length", actual)
	})
}

func Test_Seg7_LR_Trim(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Trim", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftTrim/RightTrim -- trimmed", actual)
	})
}

func Test_Seg7_LR_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_EmptyChecks", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "")

		// Act
		actual := args.Map{
			"leftEmpty":  lr.IsLeftEmpty(),
			"rightEmpty": lr.IsRightEmpty(),
			"leftWS":     lr.IsLeftWhitespace(),
			"rightWS":    lr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty": false,
			"rightEmpty": true,
			"leftWS": false,
			"rightWS": true,
		}
		expected.ShouldBeEqual(t, 0, "Empty checks -- correct", actual)
	})
}

func Test_Seg7_LR_ValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_ValidNonEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"validLeft":   lr.HasValidNonEmptyLeft(),
			"validRight":  lr.HasValidNonEmptyRight(),
			"validWSLeft": lr.HasValidNonWhitespaceLeft(),
			"validWSR":    lr.HasValidNonWhitespaceRight(),
			"safe":        lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"validLeft": true, "validRight": true,
			"validWSLeft": true, "validWSR": true, "safe": true,
		}
		expected.ShouldBeEqual(t, 0, "Valid non-empty -- all true", actual)
	})
}

func Test_Seg7_LR_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_NonPtr_Ptr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"nonPtrLeft": lr.NonPtr().Left,
			"ptrSame": lr.Ptr() == lr,
		}

		// Assert
		expected := args.Map{
			"nonPtrLeft": "a",
			"ptrSame": true,
		}
		expected.ShouldBeEqual(t, 0, "NonPtr/Ptr -- correct", actual)
	})
}

func Test_Seg7_LR_Regex(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Regex", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "xyz")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{
			"leftMatch":  lr.IsLeftRegexMatch(re),
			"rightMatch": lr.IsRightRegexMatch(re),
			"nilLeft":    lr.IsLeftRegexMatch(nil),
			"nilRight":   lr.IsRightRegexMatch(nil),
		}

		// Assert
		expected := args.Map{
			"leftMatch": true,
			"rightMatch": false,
			"nilLeft": false,
			"nilRight": false,
		}
		expected.ShouldBeEqual(t, 0, "Regex -- matches", actual)
	})
}

func Test_Seg7_LR_Is(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Is", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"isLeft":  lr.IsLeft("a"),
			"isRight": lr.IsRight("b"),
			"is":      lr.Is("a", "b"),
			"isNot":   lr.Is("a", "c"),
		}

		// Assert
		expected := args.Map{
			"isLeft": true,
			"isRight": true,
			"is": true,
			"isNot": false,
		}
		expected.ShouldBeEqual(t, 0, "Is -- checks", actual)
	})
}

func Test_Seg7_LR_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")

		// Act
		actual := args.Map{
			"eq":      lr1.IsEqual(lr2),
			"neq":     lr1.IsEqual(lr3),
			"self":    lr1.IsEqual(lr1),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
			"nilOne":  lr1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
			"self": true,
			"nilBoth": true,
			"nilOne": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqual -- various", actual)
	})
}

func Test_Seg7_LR_Clone(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"right": c.Right,
			"diff": c != lr,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
			"diff": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- new copy", actual)
	})
}

func Test_Seg7_LR_Clear(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Clear", func() {
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
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg7_LR_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Dispose", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Dispose()

		// Act
		actual := args.Map{"left": lr.Left}

		// Assert
		expected := args.Map{"left": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleared", actual)
	})
}

func Test_Seg7_LR_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Clear_Nil", func() {
		var lr *corestr.LeftRight
		lr.Clear() // should not panic
	})
}

func Test_Seg7_LR_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LR_Dispose_Nil", func() {
		var lr *corestr.LeftRight
		lr.Dispose() // should not panic
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRightFromSplit — Segment 7b
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LRFS_FromSplit(t *testing.T) {
	safeTest(t, "Test_Seg7_LRFS_FromSplit", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit -- split", actual)
	})
}

func Test_Seg7_LRFS_FromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg7_LRFS_FromSplitTrimmed", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed -- trimmed", actual)
	})
}

func Test_Seg7_LRFS_FromSplitFull(t *testing.T) {
	safeTest(t, "Test_Seg7_LRFS_FromSplitFull", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull -- first sep only", actual)
	})
}

func Test_Seg7_LRFS_FromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg7_LRFS_FromSplitFullTrimmed", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed -- trimmed", actual)
	})
}

func Test_Seg7_LRFS_FromSplit_NoSep(t *testing.T) {
	safeTest(t, "Test_Seg7_LRFS_FromSplit_NoSep", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("nosep", "=")

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit no sep -- invalid", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — Segment 7c
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LMR_New(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_New", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right, "valid": lmr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight -- valid", actual)
	})
}

func Test_Seg7_LMR_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Invalid", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{
			"valid": lmr.IsValid,
			"msg": lmr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "err",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight -- invalid", actual)
	})
}

func Test_Seg7_LMR_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"valid": lmr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage -- invalid", actual)
	})
}

func Test_Seg7_LMR_Bytes(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Bytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("ab", "cd", "ef")

		// Act
		actual := args.Map{
			"leftLen": len(lmr.LeftBytes()), "midLen": len(lmr.MiddleBytes()), "rightLen": len(lmr.RightBytes()),
		}

		// Assert
		expected := args.Map{
			"leftLen": 2,
			"midLen": 2,
			"rightLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Bytes -- correct", actual)
	})
}

func Test_Seg7_LMR_Trim(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Trim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")

		// Act
		actual := args.Map{
			"left": lmr.LeftTrim(),
			"mid": lmr.MiddleTrim(),
			"right": lmr.RightTrim(),
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "Trim -- trimmed", actual)
	})
}

func Test_Seg7_LMR_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_EmptyChecks", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "", "c")

		// Act
		actual := args.Map{
			"leftEmpty": lmr.IsLeftEmpty(), "midEmpty": lmr.IsMiddleEmpty(), "rightEmpty": lmr.IsRightEmpty(),
			"leftWS": lmr.IsLeftWhitespace(), "midWS": lmr.IsMiddleWhitespace(), "rightWS": lmr.IsRightWhitespace(),
		}

		// Assert
		expected := args.Map{
			"leftEmpty": false, "midEmpty": true, "rightEmpty": false,
			"leftWS": false, "midWS": true, "rightWS": false,
		}
		expected.ShouldBeEqual(t, 0, "Empty checks -- correct", actual)
	})
}

func Test_Seg7_LMR_ValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_ValidNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"validLeft":  lmr.HasValidNonEmptyLeft(),
			"validMid":   lmr.HasValidNonEmptyMiddle(),
			"validRight": lmr.HasValidNonEmptyRight(),
			"validWSL":   lmr.HasValidNonWhitespaceLeft(),
			"validWSM":   lmr.HasValidNonWhitespaceMiddle(),
			"validWSR":   lmr.HasValidNonWhitespaceRight(),
			"safe":       lmr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"validLeft": true, "validMid": true, "validRight": true,
			"validWSL": true, "validWSM": true, "validWSR": true, "safe": true,
		}
		expected.ShouldBeEqual(t, 0, "Valid non-empty -- all true", actual)
	})
}

func Test_Seg7_LMR_IsAll(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_IsAll", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{
			"isAll":    lmr.IsAll("a", "b", "c"),
			"isAllNot": lmr.IsAll("a", "x", "c"),
			"is":      lmr.Is("a", "c"),
			"isNot":   lmr.Is("a", "x"),
		}

		// Assert
		expected := args.Map{
			"isAll": true,
			"isAllNot": false,
			"is": true,
			"isNot": false,
		}
		expected.ShouldBeEqual(t, 0, "IsAll/Is -- checks", actual)
	})
}

func Test_Seg7_LMR_Clone(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()

		// Act
		actual := args.Map{
			"left": c.Left,
			"mid": c.Middle,
			"right": c.Right,
			"diff": c != lmr,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
			"diff": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- new copy", actual)
	})
}

func Test_Seg7_LMR_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_ToLeftRight", func() {
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
		expected.ShouldBeEqual(t, 0, "ToLeftRight -- drops middle", actual)
	})
}

func Test_Seg7_LMR_Clear(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Clear", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()

		// Act
		actual := args.Map{
			"left": lmr.Left,
			"mid": lmr.Middle,
			"right": lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "",
			"mid": "",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg7_LMR_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Dispose", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()

		// Act
		actual := args.Map{"left": lmr.Left}

		// Assert
		expected := args.Map{"left": ""}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleared", actual)
	})
}

func Test_Seg7_LMR_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Clear_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Clear()
	})
}

func Test_Seg7_LMR_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LMR_Dispose_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Dispose()
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit — Segment 7d
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LMRFS_FromSplit(t *testing.T) {
	safeTest(t, "Test_Seg7_LMRFS_FromSplit", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit -- split", actual)
	})
}

func Test_Seg7_LMRFS_FromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg7_LMRFS_FromSplitTrimmed", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed -- trimmed", actual)
	})
}

func Test_Seg7_LMRFS_FromSplitN(t *testing.T) {
	safeTest(t, "Test_Seg7_LMRFS_FromSplitN", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN -- 3 parts", actual)
	})
}

func Test_Seg7_LMRFS_FromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Seg7_LMRFS_FromSplitNTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")

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
			"right": "c : d",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed -- trimmed 3 parts", actual)
	})
}

func Test_Seg7_LMRFS_FromSplit_NoSep(t *testing.T) {
	safeTest(t, "Test_Seg7_LMRFS_FromSplit_NoSep", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("nosep", ".")

		// Act
		actual := args.Map{"valid": lmr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit no sep -- invalid", actual)
	})
}
