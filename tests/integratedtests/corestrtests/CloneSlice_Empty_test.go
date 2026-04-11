package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSlice_Empty", func() {
		// Arrange
		result := corestr.CloneSlice(nil)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- empty", actual)
	})
}

func Test_CloneSlice_Valid(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSlice_Valid", func() {
		// Arrange
		result := corestr.CloneSlice([]string{"a", "b"})

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns non-empty -- valid", actual)
	})
}

func Test_CloneSliceIf_Empty_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSliceIf_Empty", func() {
		// Arrange
		result := corestr.CloneSliceIf(true)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- empty", actual)
	})
}

func Test_CloneSliceIf_SkipClone(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSliceIf_SkipClone", func() {
		// Arrange
		result := corestr.CloneSliceIf(false, "a", "b")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- skip", actual)
	})
}

func Test_CloneSliceIf_Clone_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSliceIf_Clone", func() {
		// Arrange
		result := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- clone", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyToString_Empty_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_AnyToString_Empty", func() {
		// Arrange
		result := corestr.AnyToString(false, "")

		// Act
		actual := args.Map{"val": result}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty", actual)
	})
}

func Test_AnyToString_NoFieldName(t *testing.T) {
	safeTest(t, "Test_Cov15_AnyToString_NoFieldName", func() {
		// Arrange
		result := corestr.AnyToString(false, "hello")

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- no field", actual)
	})
}

func Test_AnyToString_WithFieldName_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_AnyToString_WithFieldName", func() {
		// Arrange
		result := corestr.AnyToString(true, "hello")

		// Act
		actual := args.Map{"notEmpty": result != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_NewLeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

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
		expected.ShouldBeEqual(t, 0, "NewLeftRight returns correct value -- with args", actual)
	})
}

func Test_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_InvalidNoMessage", func() {
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
	safeTest(t, "Test_Cov15_LeftRight_InvalidWithMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("msg")

		// Act
		actual := args.Map{
			"valid": lr.IsValid,
			"msg": lr.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight returns error -- with args", actual)
	})
}

func Test_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_UsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice(nil)

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns empty -- empty", actual)
	})
}

func Test_LeftRight_UsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_UsingSlice_One", func() {
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
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- one", actual)
	})
}

func Test_LeftRight_UsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_UsingSlice_Two", func() {
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

func Test_LeftRight_TrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_TrimmedUsingSlice_Nil", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Act
		actual := args.Map{"valid": lr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns nil -- nil", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_TrimmedUsingSlice_One", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})

		// Act
		actual := args.Map{
			"left": lr.Left,
			"valid": lr.IsValid,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- one", actual)
	})
}

func Test_LeftRight_TrimmedUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_TrimmedUsingSlice_Two", func() {
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

func Test_LeftRight_LeftBytes(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_LeftBytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("hi", "")

		// Act
		actual := args.Map{"len": len(lr.LeftBytes())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LeftRight.LeftBytes returns correct value -- with args", actual)
	})
}

func Test_LeftRight_IsLeftEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsLeftEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "b")

		// Act
		actual := args.Map{"val": lr.IsLeftEmpty()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsLeftEmpty returns empty -- with args", actual)
	})
}

func Test_LeftRight_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsRightWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "  ")

		// Act
		actual := args.Map{"val": lr.IsRightWhitespace()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsRightWhitespace returns correct value -- with args", actual)
	})
}

func Test_LeftRight_HasValidNonEmptyLeft(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_HasValidNonEmptyLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"val": lr.HasValidNonEmptyLeft()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyLeft returns empty -- with args", actual)
	})
}

func Test_LeftRight_HasValidNonWhitespaceRight(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_HasValidNonWhitespaceRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"val": lr.HasValidNonWhitespaceRight()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceRight returns non-empty -- with args", actual)
	})
}

func Test_LeftRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_HasSafeNonEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"val": lr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns empty -- with args", actual)
	})
}

func Test_LeftRight_Is(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_Is", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{
			"match": lr.Is("a", "b"),
			"noMatch": lr.Is("x", "y"),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"noMatch": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight.Is returns correct value -- with args", actual)
	})
}

func Test_LeftRight_IsEqual_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")

		// Act
		actual := args.Map{
			"equal": lr1.IsEqual(lr2),
			"notEqual": lr1.IsEqual(lr3),
		}

		// Assert
		expected := args.Map{
			"equal": true,
			"notEqual": false,
		}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsEqual returns correct value -- with args", actual)
	})
}

func Test_LeftRight_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsEqual_BothNil", func() {
		// Arrange
		var lr1, lr2 *corestr.LeftRight

		// Act
		actual := args.Map{"val": lr1.IsEqual(lr2)}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsEqual returns nil -- both nil", actual)
	})
}

func Test_LeftRight_Clone_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()

		// Act
		actual := args.Map{
			"left": cloned.Left,
			"right": cloned.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "b",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight.Clone returns correct value -- with args", actual)
	})
}

func Test_LeftRight_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_Dispose", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		lr.Dispose()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_New", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

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
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"valid": lmr.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage returns error -- with args", actual)
	})
}

func Test_LeftMiddleRight_MiddleTrim(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_MiddleTrim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", " b ", "c")

		// Act
		actual := args.Map{"val": lmr.MiddleTrim()}

		// Assert
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "MiddleTrim returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_IsMiddleWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_IsMiddleWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "  ", "c")

		// Act
		actual := args.Map{"val": lmr.IsMiddleWhitespace()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsMiddleWhitespace returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_HasValidNonEmptyMiddle(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_HasValidNonEmptyMiddle", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"val": lmr.HasValidNonEmptyMiddle()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyMiddle returns empty -- with args", actual)
	})
}

func Test_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_HasSafeNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"val": lmr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LMR.HasSafeNonEmpty returns empty -- with args", actual)
	})
}

func Test_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_IsAll", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"val": lmr.IsAll("a", "b", "c")}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LMR.IsAll returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_ToLeftRight_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LMR.ToLeftRight returns correct value -- with args", actual)
	})
}

func Test_LeftMiddleRight_Clone_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()

		// Act
		actual := args.Map{
			"left": cloned.Left,
			"mid": cloned.Middle,
			"right": cloned.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"mid": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LMR.Clone returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRightFromSplit / LeftMiddleRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRightFromSplit_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplit", func() {
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

func Test_LeftRightFromSplitTrimmed_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplitTrimmed", func() {
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

func Test_LeftRightFromSplitFull_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplitFull", func() {
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

func Test_LeftRightFromSplitFullTrimmed_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplitFullTrimmed", func() {
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

func Test_LeftMiddleRightFromSplit_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplit", func() {
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

func Test_LeftMiddleRightFromSplitTrimmed_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplitTrimmed", func() {
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

func Test_LeftMiddleRightFromSplitN_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplitN", func() {
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

func Test_LeftMiddleRightFromSplitNTrimmed_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplitNTrimmed", func() {
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
// ValidValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_New", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Empty", func() {
		// Arrange
		vv := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{
			"empty": vv.IsEmpty(),
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns empty -- with args", actual)
	})
}

func Test_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Invalid", func() {
		// Arrange
		vv := corestr.InvalidValidValue("msg")

		// Act
		actual := args.Map{
			"valid": vv.IsValid,
			"msg": vv.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns error -- with args", actual)
	})
}

func Test_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_InvalidNoMessage", func() {
		// Arrange
		vv := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{"valid": vv.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns error -- with args", actual)
	})
}

func Test_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("hi")
		b1 := vv.ValueBytesOnce()
		b2 := vv.ValueBytesOnce()

		// Act
		actual := args.Map{
			"len": len(b1),
			"same": len(b1) == len(b2),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"same": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("  ")

		// Act
		actual := args.Map{"val": vv.IsWhitespace()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "ValidValue.IsWhitespace returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue(" hello ")

		// Act
		actual := args.Map{"val": vv.Trim()}

		// Assert
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "ValidValue.Trim returns non-empty -- with args", actual)
	})
}

func Test_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{"val": vv.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns empty -- with args", actual)
	})
}

func Test_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueBool", func() {
		// Arrange
		vv := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"val": vv.ValueBool()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
	})
}

func Test_ValidValue_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueBool_Empty", func() {
		// Arrange
		vv := corestr.NewValidValue("")

		// Act
		actual := args.Map{"val": vv.ValueBool()}

		// Assert
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns empty -- empty", actual)
	})
}

func Test_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueInt", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"val": vv.ValueInt(0)}

		// Assert
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
	})
}

func Test_ValidValue_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueInt_Invalid", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{"val": vv.ValueInt(99)}

		// Assert
		expected := args.Map{"val": 99}
		expected.ShouldBeEqual(t, 0, "ValueInt returns error -- invalid", actual)
	})
}

func Test_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueByte", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"val": vv.ValueByte(0)}

		// Assert
		expected := args.Map{"val": byte(42)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- with args", actual)
	})
}

func Test_ValidValue_ValueByte_TooHigh(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueByte_TooHigh", func() {
		// Arrange
		vv := corestr.NewValidValue("999")

		// Act
		actual := args.Map{"val": vv.ValueByte(0)}

		// Assert
		expected := args.Map{"val": byte(255)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- too high", actual)
	})
}

func Test_ValidValue_ValueByte_Negative(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueByte_Negative", func() {
		// Arrange
		vv := corestr.NewValidValue("-1")

		// Act
		actual := args.Map{"val": vv.ValueByte(0)}

		// Assert
		expected := args.Map{"val": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- negative", actual)
	})
}

func Test_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")
		result := vv.ValueFloat64(0)

		// Act
		actual := args.Map{"positive": result > 3}

		// Assert
		expected := args.Map{"positive": true}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act
		actual := args.Map{
			"match": vv.IsAnyOf("a", "b", "c"),
			"noMatch": vv.IsAnyOf("x"),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"noMatch": false,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsAnyOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsAnyOf_Empty", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act
		actual := args.Map{"val": vv.IsAnyOf()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns empty -- empty", actual)
	})
}

func Test_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"val": vv.IsContains("world")}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{
			"match": vv.IsAnyContains("world"),
			"empty": vv.IsAnyContains(),
		}

		// Assert
		expected := args.Map{
			"match": true,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "IsAnyContains returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{"val": vv.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive returns correct value -- with args", actual)
	})
}

func Test_ValidValue_IsRegexMatches_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsRegexMatches_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"val": vv.IsRegexMatches(nil)}

		// Assert
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches returns nil -- nil", actual)
	})
}

func Test_ValidValue_RegexFindString_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_RegexFindString_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"val": vv.RegexFindString(nil)}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "RegexFindString returns nil -- nil", actual)
	})
}

func Test_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		items, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)

		// Act
		actual := args.Map{
			"len": len(items),
			"hasAny": hasAny,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag returns nil -- nil", actual)
	})
}

func Test_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_RegexFindAllStrings_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		items := vv.RegexFindAllStrings(nil, -1)

		// Act
		actual := args.Map{"len": len(items)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings returns nil -- nil", actual)
	})
}

func Test_ValidValue_Split_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")
		items := vv.Split(",")

		// Act
		actual := args.Map{"len": len(items)}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Split returns correct value -- with args", actual)
	})
}

func Test_ValidValue_Clone_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("hi")
		cloned := vv.Clone()

		// Act
		actual := args.Map{
			"val": cloned.Value,
			"valid": cloned.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "hi",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
	})
}

func Test_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Clone_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue
		cloned := vv.Clone()

		// Act
		actual := args.Map{"nil": cloned == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
	})
}

func Test_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_String_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"val": vv.String()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "String returns nil -- nil", actual)
	})
}

func Test_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_FullString_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"val": vv.FullString()}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "FullString returns nil -- nil", actual)
	})
}

func Test_ValidValue_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("hi")
		vv.Dispose()

		// Act
		actual := args.Map{
			"empty": vv.IsEmpty(),
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"valid": false,
		}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{
			"empty": vvs.IsEmpty(),
			"len": vvs.Length(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "EmptyValidValues returns empty -- with args", actual)
	})
}

func Test_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Add", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("hello")

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"hasAny": vvs.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"hasAny": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues.Add returns non-empty -- with args", actual)
	})
}

func Test_ValidValues_Strings_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Strings", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vvs.Add("b")
		strs := vvs.Strings()

		// Act
		actual := args.Map{"len": len(strs)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues.Strings returns non-empty -- with args", actual)
	})
}

func Test_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_SafeValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("hello")

		// Act
		actual := args.Map{
			"val": vvs.SafeValueAt(0),
			"oob": vvs.SafeValueAt(5),
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"oob": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValueAt returns correct value -- with args", actual)
	})
}

func Test_ValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Nil", func() {
		// Arrange
		var vvs *corestr.ValidValues

		// Act
		actual := args.Map{
			"len": vvs.Length(),
			"empty": vvs.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov15_ValueStatus_Invalid", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("msg")

		// Act
		actual := args.Map{
			"valid": vs.ValueValid.IsValid,
			"msg": vs.ValueValid.Message,
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"msg": "msg",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns error -- with args", actual)
	})
}

func Test_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_ValueStatus_InvalidNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{"valid": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns error -- with args", actual)
	})
}

func Test_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_Cov15_ValueStatus_Clone", func() {
		// Arrange
		vs := &corestr.ValueStatus{ValueValid: corestr.NewValidValue("hi"), Index: 3}
		cloned := vs.Clone()

		// Act
		actual := args.Map{
			"val": cloned.ValueValid.Value,
			"idx": cloned.Index,
		}

		// Assert
		expected := args.Map{
			"val": "hi",
			"idx": 3,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus.Clone returns non-empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov15_TextWithLineNumber_HasLineNumber", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}

		// Act
		actual := args.Map{
			"has": tln.HasLineNumber(),
			"invalid": tln.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber.HasLineNumber returns non-empty -- with args", actual)
	})
}

func Test_TextWithLineNumber_Nil_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_TextWithLineNumber_Nil", func() {
		// Arrange
		var tln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"len": tln.Length(),
			"empty": tln.IsEmpty(),
			"emptyText": tln.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
			"emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns nil -- nil", actual)
	})
}

func Test_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_Cov15_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"val": tln.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth returns empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_Basic", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"key": kvp.KeyName(), "val": kvp.ValueString(),
			"isKey": kvp.IsKey("k"), "isVal": kvp.IsVal("v"),
			"hasKey": kvp.HasKey(), "hasVal": kvp.HasValue(),
		}

		// Assert
		expected := args.Map{
			"key": "k", "val": "v",
			"isKey": true, "isVal": true,
			"hasKey": true, "hasVal": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- basic", actual)
	})
}

func Test_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueBool", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "true"}

		// Act
		actual := args.Map{"val": kvp.ValueBool()}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueBool returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueBool_Empty", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act
		actual := args.Map{"val": kvp.ValueBool()}

		// Assert
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueBool returns empty -- empty", actual)
	})
}

func Test_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueInt", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "42"}

		// Act
		actual := args.Map{"val": kvp.ValueInt(0)}

		// Assert
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueInt returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueByte", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "42"}

		// Act
		actual := args.Map{"val": kvp.ValueByte(0)}

		// Assert
		expected := args.Map{"val": byte(42)}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueByte returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_ValueByte_TooHigh(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueByte_TooHigh", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "999"}

		// Act
		actual := args.Map{"val": kvp.ValueByte(5)}

		// Assert
		expected := args.Map{"val": byte(5)}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueByte returns correct value -- high", actual)
	})
}

func Test_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueFloat64", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		result := kvp.ValueFloat64(0)

		// Act
		actual := args.Map{"positive": result > 3}

		// Assert
		expected := args.Map{"positive": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueFloat64 returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_ValueValid_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueValid", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValid()

		// Act
		actual := args.Map{
			"val": vv.Value,
			"valid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "v",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueValid returns non-empty -- with args", actual)
	})
}

func Test_KeyValuePair_String(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_String", func() {
		// Arrange
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"notEmpty": kvp.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.String returns correct value -- with args", actual)
	})
}

func Test_KeyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_Dispose", func() {
		// Arrange
		kvp := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvp.Dispose()

		// Act
		actual := args.Map{
			"keyEmpty": kvp.Key == "",
			"valEmpty": kvp.Value == "",
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true,
			"valEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_Basic", func() {
		// Arrange
		kavp := corestr.KeyAnyValuePair{Key: "k", Value: 42}

		// Act
		actual := args.Map{
			"key": kavp.KeyName(),
			"hasVal": kavp.HasValue(),
			"notNull": kavp.HasNonNull(),
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"hasVal": true,
			"notNull": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- basic", actual)
	})
}

func Test_KeyAnyValuePair_IsValueNull(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_IsValueNull", func() {
		// Arrange
		kavp := corestr.KeyAnyValuePair{Key: "k"}

		// Act
		actual := args.Map{"null": kavp.IsValueNull()}

		// Assert
		expected := args.Map{"null": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair.IsValueNull returns correct value -- with args", actual)
	})
}

func Test_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_ValueString", func() {
		// Arrange
		kavp := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}

		// Act
		actual := args.Map{"notEmpty": kavp.ValueString() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair.ValueString returns non-empty -- with args", actual)
	})
}

func Test_KeyAnyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_Dispose", func() {
		// Arrange
		kavp := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kavp.Dispose()

		// Act
		actual := args.Map{
			"keyEmpty": kavp.Key == "",
			"null": kavp.IsValueNull(),
		}

		// Assert
		expected := args.Map{
			"keyEmpty": true,
			"null": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair.Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// StringUtils (utils)
// ══════════════════════════════════════════════════════════════════════════════

func Test_StringUtils_WrapDoubleIfMissing_Empty_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDoubleIfMissing_Empty", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapDoubleIfMissing("")}

		// Assert
		expected := args.Map{"val": `""`}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing returns empty -- empty", actual)
	})
}

func Test_StringUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapDoubleIfMissing(`"hi"`)}

		// Assert
		expected := args.Map{"val": `"hi"`}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing returns correct value -- already", actual)
	})
}

func Test_StringUtils_WrapDoubleIfMissing_NotWrapped_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDoubleIfMissing_NotWrapped", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapDoubleIfMissing("hi")}

		// Assert
		expected := args.Map{"val": `"hi"`}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing returns correct value -- not wrapped", actual)
	})
}

func Test_StringUtils_WrapSingleIfMissing_Empty_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingleIfMissing_Empty", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapSingleIfMissing("")}

		// Assert
		expected := args.Map{"val": "''"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing returns empty -- empty", actual)
	})
}

func Test_StringUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapSingleIfMissing("'hi'")}

		// Assert
		expected := args.Map{"val": "'hi'"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing returns correct value -- already", actual)
	})
}

func Test_StringUtils_WrapSingleIfMissing_NotWrapped_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingleIfMissing_NotWrapped", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapSingleIfMissing("hi")}

		// Assert
		expected := args.Map{"val": "'hi'"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing returns correct value -- not wrapped", actual)
	})
}

func Test_StringUtils_WrapDouble_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDouble", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapDouble("hi")}

		// Assert
		expected := args.Map{"val": `"hi"`}
		expected.ShouldBeEqual(t, 0, "WrapDouble returns correct value -- with args", actual)
	})
}

func Test_StringUtils_WrapSingle_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingle", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapSingle("hi")}

		// Assert
		expected := args.Map{"val": "'hi'"}
		expected.ShouldBeEqual(t, 0, "WrapSingle returns correct value -- with args", actual)
	})
}

func Test_StringUtils_WrapTilda_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapTilda", func() {
		// Act
		actual := args.Map{"val": corestr.StringUtils.WrapTilda("hi")}

		// Assert
		expected := args.Map{"val": "`hi`"}
		expected.ShouldBeEqual(t, 0, "WrapTilda returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Key methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleStringOnce_GetSetOnce_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_GetSetOnce", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		val := sso.GetSetOnce("hello")

		// Act
		actual := args.Map{
			"val": val,
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "GetSetOnce returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_GetSetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_GetSetOnce_AlreadyInit", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("first")
		val := sso.GetSetOnce("second")

		// Act
		actual := args.Map{"val": val}

		// Assert
		expected := args.Map{"val": "first"}
		expected.ShouldBeEqual(t, 0, "GetSetOnce returns correct value -- already init", actual)
	})
}

func Test_SimpleStringOnce_GetOnce_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_GetOnce", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		val := sso.GetOnce()

		// Act
		actual := args.Map{
			"val": val,
			"init": sso.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"val": "",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "GetOnce returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_SetOnUninitialized", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		err := sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"val": sso.Value(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "SetOnUninitialized returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_SetOnUninitialized_AlreadyInit_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("first")
		err := sso.SetOnUninitialized("second")

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "SetOnUninitialized returns correct value -- already init", actual)
	})
}

func Test_SimpleStringOnce_Invalidate_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Invalidate", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("hello")
		sso.Invalidate()

		// Act
		actual := args.Map{
			"init": sso.IsInitialized(),
			"empty": sso.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"init": false,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Invalidate returns error -- with args", actual)
	})
}

func Test_SimpleStringOnce_Boolean_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Boolean", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("yes")

		// Act
		actual := args.Map{"val": sso.Boolean(true)}

		// Assert
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "Boolean returns correct value -- yes", actual)
	})
}

func Test_SimpleStringOnce_Boolean_Uninit_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Boolean_Uninit", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"val": sso.Boolean(true)}

		// Assert
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "Boolean returns correct value -- uninit", actual)
	})
}

func Test_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Int", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("42")

		// Act
		actual := args.Map{"val": sso.Int()}

		// Assert
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "Int returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_Byte(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Byte", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("42")

		// Act
		actual := args.Map{"val": sso.Byte()}

		// Assert
		expected := args.Map{"val": byte(42)}
		expected.ShouldBeEqual(t, 0, "Byte returns correct value -- with args", actual)
	})
}

func Test_SimpleStringOnce_ConcatNew_FromCloneSliceEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_ConcatNew", func() {
		// Arrange
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("hello")
		result := sso.ConcatNew(" world")

		// Act
		actual := args.Map{"val": result.Value()}

		// Assert
		expected := args.Map{"val": "hello world"}
		expected.ShouldBeEqual(t, 0, "ConcatNew returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnceModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleStringOnceModel(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnceModel", func() {
		// Arrange
		m := corestr.SimpleStringOnceModel{Value: "hi", IsInitialize: true}

		// Act
		actual := args.Map{
			"val": m.Value,
			"init": m.IsInitialize,
		}

		// Assert
		expected := args.Map{
			"val": "hi",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnceModel returns correct value -- with args", actual)
	})
}
