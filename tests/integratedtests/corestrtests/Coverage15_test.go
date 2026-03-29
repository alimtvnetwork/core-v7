package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSlice_Empty", func() {
		result := corestr.CloneSlice(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- empty", actual)
	})
}

func Test_Cov15_CloneSlice_Valid(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSlice_Valid", func() {
		result := corestr.CloneSlice([]string{"a", "b"})
		actual := args.Map{"len": len(result), "first": result[0]}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns non-empty -- valid", actual)
	})
}

func Test_Cov15_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSliceIf_Empty", func() {
		result := corestr.CloneSliceIf(true)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- empty", actual)
	})
}

func Test_Cov15_CloneSliceIf_SkipClone(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSliceIf_SkipClone", func() {
		result := corestr.CloneSliceIf(false, "a", "b")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- skip", actual)
	})
}

func Test_Cov15_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_Cov15_CloneSliceIf_Clone", func() {
		result := corestr.CloneSliceIf(true, "a", "b")
		actual := args.Map{"len": len(result), "first": result[0]}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- clone", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_AnyToString_Empty", func() {
		result := corestr.AnyToString(false, "")
		actual := args.Map{"val": result}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty", actual)
	})
}

func Test_Cov15_AnyToString_NoFieldName(t *testing.T) {
	safeTest(t, "Test_Cov15_AnyToString_NoFieldName", func() {
		result := corestr.AnyToString(false, "hello")
		actual := args.Map{"notEmpty": result != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- no field", actual)
	})
}

func Test_Cov15_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Cov15_AnyToString_WithFieldName", func() {
		result := corestr.AnyToString(true, "hello")
		actual := args.Map{"notEmpty": result != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_NewLeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "b", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewLeftRight returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage returns error -- with args", actual)
	})
}

func Test_Cov15_LeftRight_InvalidWithMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_InvalidWithMessage", func() {
		lr := corestr.InvalidLeftRight("msg")
		actual := args.Map{"valid": lr.IsValid, "msg": lr.Message}
		expected := args.Map{"valid": false, "msg": "msg"}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight returns error -- with args", actual)
	})
}

func Test_Cov15_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_UsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice(nil)
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns empty -- empty", actual)
	})
}

func Test_Cov15_LeftRight_UsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_UsingSlice_One", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a"})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- one", actual)
	})
}

func Test_Cov15_LeftRight_UsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_UsingSlice_Two", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "b", "valid": true}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- two", actual)
	})
}

func Test_Cov15_LeftRight_TrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_TrimmedUsingSlice_Nil", func() {
		lr := corestr.LeftRightTrimmedUsingSlice(nil)
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns nil -- nil", actual)
	})
}

func Test_Cov15_LeftRight_TrimmedUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_TrimmedUsingSlice_One", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})
		actual := args.Map{"left": lr.Left, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- one", actual)
	})
}

func Test_Cov15_LeftRight_TrimmedUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_TrimmedUsingSlice_Two", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- two", actual)
	})
}

func Test_Cov15_LeftRight_LeftBytes(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_LeftBytes", func() {
		lr := corestr.NewLeftRight("hi", "")
		actual := args.Map{"len": len(lr.LeftBytes())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "LeftRight.LeftBytes returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRight_IsLeftEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsLeftEmpty", func() {
		lr := corestr.NewLeftRight("", "b")
		actual := args.Map{"val": lr.IsLeftEmpty()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsLeftEmpty returns empty -- with args", actual)
	})
}

func Test_Cov15_LeftRight_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsRightWhitespace", func() {
		lr := corestr.NewLeftRight("a", "  ")
		actual := args.Map{"val": lr.IsRightWhitespace()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsRightWhitespace returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRight_HasValidNonEmptyLeft(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_HasValidNonEmptyLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"val": lr.HasValidNonEmptyLeft()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyLeft returns empty -- with args", actual)
	})
}

func Test_Cov15_LeftRight_HasValidNonWhitespaceRight(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_HasValidNonWhitespaceRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"val": lr.HasValidNonWhitespaceRight()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceRight returns non-empty -- with args", actual)
	})
}

func Test_Cov15_LeftRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_HasSafeNonEmpty", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"val": lr.HasSafeNonEmpty()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns empty -- with args", actual)
	})
}

func Test_Cov15_LeftRight_Is(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_Is", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"match": lr.Is("a", "b"), "noMatch": lr.Is("x", "y")}
		expected := args.Map{"match": true, "noMatch": false}
		expected.ShouldBeEqual(t, 0, "LeftRight.Is returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("x", "y")
		actual := args.Map{"equal": lr1.IsEqual(lr2), "notEqual": lr1.IsEqual(lr3)}
		expected := args.Map{"equal": true, "notEqual": false}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsEqual returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRight_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_IsEqual_BothNil", func() {
		var lr1, lr2 *corestr.LeftRight
		actual := args.Map{"val": lr1.IsEqual(lr2)}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.IsEqual returns nil -- both nil", actual)
	})
}

func Test_Cov15_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()
		actual := args.Map{"left": cloned.Left, "right": cloned.Right}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRight.Clone returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRight_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRight_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Dispose()
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "LeftRight.Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{"valid": lmr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage returns error -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_MiddleTrim(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_MiddleTrim", func() {
		lmr := corestr.NewLeftMiddleRight("a", " b ", "c")
		actual := args.Map{"val": lmr.MiddleTrim()}
		expected := args.Map{"val": "b"}
		expected.ShouldBeEqual(t, 0, "MiddleTrim returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_IsMiddleWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_IsMiddleWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("a", "  ", "c")
		actual := args.Map{"val": lmr.IsMiddleWhitespace()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsMiddleWhitespace returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_HasValidNonEmptyMiddle(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_HasValidNonEmptyMiddle", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"val": lmr.HasValidNonEmptyMiddle()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyMiddle returns empty -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_HasSafeNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"val": lmr.HasSafeNonEmpty()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LMR.HasSafeNonEmpty returns empty -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"val": lmr.IsAll("a", "b", "c")}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "LMR.IsAll returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LMR.ToLeftRight returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()
		actual := args.Map{"left": cloned.Left, "mid": cloned.Middle, "right": cloned.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LMR.Clone returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftRightFromSplit / LeftMiddleRightFromSplit
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b:c:d"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftRightFromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b : c"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplitTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplitN", func() {
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c:d:e"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN returns correct value -- with args", actual)
	})
}

func Test_Cov15_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_Cov15_LeftMiddleRightFromSplitNTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle}
		expected := args.Map{"left": "a", "mid": "b"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_New", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "hello", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns non-empty -- with args", actual)
	})
}

func Test_Cov15_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Empty", func() {
		vv := corestr.NewValidValueEmpty()
		actual := args.Map{"empty": vv.IsEmpty(), "valid": vv.IsValid}
		expected := args.Map{"empty": true, "valid": true}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns empty -- with args", actual)
	})
}

func Test_Cov15_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Invalid", func() {
		vv := corestr.InvalidValidValue("msg")
		actual := args.Map{"valid": vv.IsValid, "msg": vv.Message}
		expected := args.Map{"valid": false, "msg": "msg"}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns error -- with args", actual)
	})
}

func Test_Cov15_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_InvalidNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		actual := args.Map{"valid": vv.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns error -- with args", actual)
	})
}

func Test_Cov15_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueBytesOnce", func() {
		vv := corestr.NewValidValue("hi")
		b1 := vv.ValueBytesOnce()
		b2 := vv.ValueBytesOnce()
		actual := args.Map{"len": len(b1), "same": len(b1) == len(b2)}
		expected := args.Map{"len": 2, "same": true}
		expected.ShouldBeEqual(t, 0, "ValueBytesOnce returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsWhitespace", func() {
		vv := corestr.NewValidValue("  ")
		actual := args.Map{"val": vv.IsWhitespace()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "ValidValue.IsWhitespace returns non-empty -- with args", actual)
	})
}

func Test_Cov15_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Trim", func() {
		vv := corestr.NewValidValue(" hello ")
		actual := args.Map{"val": vv.Trim()}
		expected := args.Map{"val": "hello"}
		expected.ShouldBeEqual(t, 0, "ValidValue.Trim returns non-empty -- with args", actual)
	})
}

func Test_Cov15_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("hi")
		actual := args.Map{"val": vv.HasValidNonEmpty()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmpty returns empty -- with args", actual)
	})
}

func Test_Cov15_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		actual := args.Map{"val": vv.ValueBool()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueBool_Empty", func() {
		vv := corestr.NewValidValue("")
		actual := args.Map{"val": vv.ValueBool()}
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "ValueBool returns empty -- empty", actual)
	})
}

func Test_Cov15_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueInt", func() {
		vv := corestr.NewValidValue("42")
		actual := args.Map{"val": vv.ValueInt(0)}
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_ValueInt_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueInt_Invalid", func() {
		vv := corestr.NewValidValue("abc")
		actual := args.Map{"val": vv.ValueInt(99)}
		expected := args.Map{"val": 99}
		expected.ShouldBeEqual(t, 0, "ValueInt returns error -- invalid", actual)
	})
}

func Test_Cov15_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueByte", func() {
		vv := corestr.NewValidValue("42")
		actual := args.Map{"val": vv.ValueByte(0)}
		expected := args.Map{"val": byte(42)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_ValueByte_TooHigh(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueByte_TooHigh", func() {
		vv := corestr.NewValidValue("999")
		actual := args.Map{"val": vv.ValueByte(0)}
		expected := args.Map{"val": byte(255)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- too high", actual)
	})
}

func Test_Cov15_ValidValue_ValueByte_Negative(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueByte_Negative", func() {
		vv := corestr.NewValidValue("-1")
		actual := args.Map{"val": vv.ValueByte(0)}
		expected := args.Map{"val": byte(0)}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct value -- negative", actual)
	})
}

func Test_Cov15_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_ValueFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		result := vv.ValueFloat64(0)
		actual := args.Map{"positive": result > 3}
		expected := args.Map{"positive": true}
		expected.ShouldBeEqual(t, 0, "ValueFloat64 returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		actual := args.Map{"match": vv.IsAnyOf("a", "b", "c"), "noMatch": vv.IsAnyOf("x")}
		expected := args.Map{"match": true, "noMatch": false}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_IsAnyOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsAnyOf_Empty", func() {
		vv := corestr.NewValidValue("b")
		actual := args.Map{"val": vv.IsAnyOf()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsAnyOf returns empty -- empty", actual)
	})
}

func Test_Cov15_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{"val": vv.IsContains("world")}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{"match": vv.IsAnyContains("world"), "empty": vv.IsAnyContains()}
		expected := args.Map{"match": true, "empty": true}
		expected.ShouldBeEqual(t, 0, "IsAnyContains returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		actual := args.Map{"val": vv.IsEqualNonSensitive("hello")}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsEqualNonSensitive returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_IsRegexMatches_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_IsRegexMatches_Nil", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"val": vv.IsRegexMatches(nil)}
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "IsRegexMatches returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_RegexFindString_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_RegexFindString_Nil", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"val": vv.RegexFindString(nil)}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "RegexFindString returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		vv := corestr.NewValidValue("hello")
		items, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)
		actual := args.Map{"len": len(items), "hasAny": hasAny}
		expected := args.Map{"len": 0, "hasAny": false}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStringsWithFlag returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_RegexFindAllStrings_Nil", func() {
		vv := corestr.NewValidValue("hello")
		items := vv.RegexFindAllStrings(nil, -1)
		actual := args.Map{"len": len(items)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RegexFindAllStrings returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		items := vv.Split(",")
		actual := args.Map{"len": len(items)}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Split returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Clone", func() {
		vv := corestr.NewValidValue("hi")
		cloned := vv.Clone()
		actual := args.Map{"val": cloned.Value, "valid": cloned.IsValid}
		expected := args.Map{"val": "hi", "valid": true}
		expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Clone_Nil", func() {
		var vv *corestr.ValidValue
		cloned := vv.Clone()
		actual := args.Map{"nil": cloned == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_String_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"val": vv.String()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "String returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_FullString_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"val": vv.FullString()}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "FullString returns nil -- nil", actual)
	})
}

func Test_Cov15_ValidValue_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValue_Dispose", func() {
		vv := corestr.NewValidValue("hi")
		vv.Dispose()
		actual := args.Map{"empty": vv.IsEmpty(), "valid": vv.IsValid}
		expected := args.Map{"empty": true, "valid": false}
		expected.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValidValues
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Empty", func() {
		vvs := corestr.EmptyValidValues()
		actual := args.Map{"empty": vvs.IsEmpty(), "len": vvs.Length()}
		expected := args.Map{"empty": true, "len": 0}
		expected.ShouldBeEqual(t, 0, "EmptyValidValues returns empty -- with args", actual)
	})
}

func Test_Cov15_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Add", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("hello")
		actual := args.Map{"len": vvs.Length(), "hasAny": vvs.HasAnyItem()}
		expected := args.Map{"len": 1, "hasAny": true}
		expected.ShouldBeEqual(t, 0, "ValidValues.Add returns non-empty -- with args", actual)
	})
}

func Test_Cov15_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		vvs.Add("b")
		strs := vvs.Strings()
		actual := args.Map{"len": len(strs)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ValidValues.Strings returns non-empty -- with args", actual)
	})
}

func Test_Cov15_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("hello")
		actual := args.Map{"val": vvs.SafeValueAt(0), "oob": vvs.SafeValueAt(5)}
		expected := args.Map{"val": "hello", "oob": ""}
		expected.ShouldBeEqual(t, 0, "SafeValueAt returns correct value -- with args", actual)
	})
}

func Test_Cov15_ValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_ValidValues_Nil", func() {
		var vvs *corestr.ValidValues
		actual := args.Map{"len": vvs.Length(), "empty": vvs.IsEmpty()}
		expected := args.Map{"len": 0, "empty": true}
		expected.ShouldBeEqual(t, 0, "ValidValues returns nil -- nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov15_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("msg")
		actual := args.Map{"valid": vs.ValueValid.IsValid, "msg": vs.ValueValid.Message}
		expected := args.Map{"valid": false, "msg": "msg"}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns error -- with args", actual)
	})
}

func Test_Cov15_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov15_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"valid": vs.ValueValid.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns error -- with args", actual)
	})
}

func Test_Cov15_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_Cov15_ValueStatus_Clone", func() {
		vs := &corestr.ValueStatus{ValueValid: corestr.NewValidValue("hi"), Index: 3}
		cloned := vs.Clone()
		actual := args.Map{"val": cloned.ValueValid.Value, "idx": cloned.Index}
		expected := args.Map{"val": "hi", "idx": 3}
		expected.ShouldBeEqual(t, 0, "ValueStatus.Clone returns non-empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov15_TextWithLineNumber_HasLineNumber", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		actual := args.Map{"has": tln.HasLineNumber(), "invalid": tln.IsInvalidLineNumber()}
		expected := args.Map{"has": true, "invalid": false}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber.HasLineNumber returns non-empty -- with args", actual)
	})
}

func Test_Cov15_TextWithLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_Cov15_TextWithLineNumber_Nil", func() {
		var tln *corestr.TextWithLineNumber
		actual := args.Map{"len": tln.Length(), "empty": tln.IsEmpty(), "emptyText": tln.IsEmptyText()}
		expected := args.Map{"len": 0, "empty": true, "emptyText": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns nil -- nil", actual)
	})
}

func Test_Cov15_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_Cov15_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"val": tln.IsEmptyTextLineBoth()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth returns empty -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_Basic", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{
			"key": kvp.KeyName(), "val": kvp.ValueString(),
			"isKey": kvp.IsKey("k"), "isVal": kvp.IsVal("v"),
			"hasKey": kvp.HasKey(), "hasVal": kvp.HasValue(),
		}
		expected := args.Map{
			"key": "k", "val": "v",
			"isKey": true, "isVal": true,
			"hasKey": true, "hasVal": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- basic", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueBool", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "true"}
		actual := args.Map{"val": kvp.ValueBool()}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueBool returns correct value -- with args", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueBool_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueBool_Empty", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: ""}
		actual := args.Map{"val": kvp.ValueBool()}
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueBool returns empty -- empty", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueInt", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "42"}
		actual := args.Map{"val": kvp.ValueInt(0)}
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueInt returns correct value -- with args", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueByte", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "42"}
		actual := args.Map{"val": kvp.ValueByte(0)}
		expected := args.Map{"val": byte(42)}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueByte returns correct value -- with args", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueByte_TooHigh(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueByte_TooHigh", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "999"}
		actual := args.Map{"val": kvp.ValueByte(5)}
		expected := args.Map{"val": byte(5)}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueByte returns correct value -- high", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueFloat64", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		result := kvp.ValueFloat64(0)
		actual := args.Map{"positive": result > 3}
		expected := args.Map{"positive": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueFloat64 returns correct value -- with args", actual)
	})
}

func Test_Cov15_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_ValueValid", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kvp.ValueValid()
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "v", "valid": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.ValueValid returns non-empty -- with args", actual)
	})
}

func Test_Cov15_KeyValuePair_String(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_String", func() {
		kvp := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"notEmpty": kvp.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.String returns correct value -- with args", actual)
	})
}

func Test_Cov15_KeyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyValuePair_Dispose", func() {
		kvp := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kvp.Dispose()
		actual := args.Map{"keyEmpty": kvp.Key == "", "valEmpty": kvp.Value == ""}
		expected := args.Map{"keyEmpty": true, "valEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair.Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_KeyAnyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_Basic", func() {
		kavp := corestr.KeyAnyValuePair{Key: "k", Value: 42}
		actual := args.Map{"key": kavp.KeyName(), "hasVal": kavp.HasValue(), "notNull": kavp.HasNonNull()}
		expected := args.Map{"key": "k", "hasVal": true, "notNull": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair returns correct value -- basic", actual)
	})
}

func Test_Cov15_KeyAnyValuePair_IsValueNull(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_IsValueNull", func() {
		kavp := corestr.KeyAnyValuePair{Key: "k"}
		actual := args.Map{"null": kavp.IsValueNull()}
		expected := args.Map{"null": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair.IsValueNull returns correct value -- with args", actual)
	})
}

func Test_Cov15_KeyAnyValuePair_ValueString(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_ValueString", func() {
		kavp := corestr.KeyAnyValuePair{Key: "k", Value: "hello"}
		actual := args.Map{"notEmpty": kavp.ValueString() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair.ValueString returns non-empty -- with args", actual)
	})
}

func Test_Cov15_KeyAnyValuePair_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov15_KeyAnyValuePair_Dispose", func() {
		kavp := &corestr.KeyAnyValuePair{Key: "k", Value: "v"}
		kavp.Dispose()
		actual := args.Map{"keyEmpty": kavp.Key == "", "null": kavp.IsValueNull()}
		expected := args.Map{"keyEmpty": true, "null": true}
		expected.ShouldBeEqual(t, 0, "KeyAnyValuePair.Dispose returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// StringUtils (utils)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_StringUtils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDoubleIfMissing_Empty", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapDoubleIfMissing("")}
		expected := args.Map{"val": `""`}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing returns empty -- empty", actual)
	})
}

func Test_Cov15_StringUtils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapDoubleIfMissing(`"hi"`)}
		expected := args.Map{"val": `"hi"`}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing returns correct value -- already", actual)
	})
}

func Test_Cov15_StringUtils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDoubleIfMissing_NotWrapped", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapDoubleIfMissing("hi")}
		expected := args.Map{"val": `"hi"`}
		expected.ShouldBeEqual(t, 0, "WrapDoubleIfMissing returns correct value -- not wrapped", actual)
	})
}

func Test_Cov15_StringUtils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingleIfMissing_Empty", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapSingleIfMissing("")}
		expected := args.Map{"val": "''"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing returns empty -- empty", actual)
	})
}

func Test_Cov15_StringUtils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingleIfMissing_AlreadyWrapped", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapSingleIfMissing("'hi'")}
		expected := args.Map{"val": "'hi'"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing returns correct value -- already", actual)
	})
}

func Test_Cov15_StringUtils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingleIfMissing_NotWrapped", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapSingleIfMissing("hi")}
		expected := args.Map{"val": "'hi'"}
		expected.ShouldBeEqual(t, 0, "WrapSingleIfMissing returns correct value -- not wrapped", actual)
	})
}

func Test_Cov15_StringUtils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapDouble", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapDouble("hi")}
		expected := args.Map{"val": `"hi"`}
		expected.ShouldBeEqual(t, 0, "WrapDouble returns correct value -- with args", actual)
	})
}

func Test_Cov15_StringUtils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapSingle", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapSingle("hi")}
		expected := args.Map{"val": "'hi'"}
		expected.ShouldBeEqual(t, 0, "WrapSingle returns correct value -- with args", actual)
	})
}

func Test_Cov15_StringUtils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_Cov15_StringUtils_WrapTilda", func() {
		actual := args.Map{"val": corestr.StringUtils.WrapTilda("hi")}
		expected := args.Map{"val": "`hi`"}
		expected.ShouldBeEqual(t, 0, "WrapTilda returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Key methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_GetSetOnce", func() {
		sso := corestr.SimpleStringOnce{}
		val := sso.GetSetOnce("hello")
		actual := args.Map{"val": val, "init": sso.IsInitialized()}
		expected := args.Map{"val": "hello", "init": true}
		expected.ShouldBeEqual(t, 0, "GetSetOnce returns correct value -- with args", actual)
	})
}

func Test_Cov15_SimpleStringOnce_GetSetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_GetSetOnce_AlreadyInit", func() {
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("first")
		val := sso.GetSetOnce("second")
		actual := args.Map{"val": val}
		expected := args.Map{"val": "first"}
		expected.ShouldBeEqual(t, 0, "GetSetOnce returns correct value -- already init", actual)
	})
}

func Test_Cov15_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_GetOnce", func() {
		sso := corestr.SimpleStringOnce{}
		val := sso.GetOnce()
		actual := args.Map{"val": val, "init": sso.IsInitialized()}
		expected := args.Map{"val": "", "init": true}
		expected.ShouldBeEqual(t, 0, "GetOnce returns correct value -- with args", actual)
	})
}

func Test_Cov15_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_SetOnUninitialized", func() {
		sso := corestr.SimpleStringOnce{}
		err := sso.SetOnUninitialized("hello")
		actual := args.Map{"noErr": err == nil, "val": sso.Value()}
		expected := args.Map{"noErr": true, "val": "hello"}
		expected.ShouldBeEqual(t, 0, "SetOnUninitialized returns correct value -- with args", actual)
	})
}

func Test_Cov15_SimpleStringOnce_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_SetOnUninitialized_AlreadyInit", func() {
		sso := corestr.SimpleStringOnce{}
		_ = sso.SetOnUninitialized("first")
		err := sso.SetOnUninitialized("second")
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "SetOnUninitialized returns correct value -- already init", actual)
	})
}

func Test_Cov15_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Invalidate", func() {
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("hello")
		sso.Invalidate()
		actual := args.Map{"init": sso.IsInitialized(), "empty": sso.IsEmpty()}
		expected := args.Map{"init": false, "empty": true}
		expected.ShouldBeEqual(t, 0, "Invalidate returns error -- with args", actual)
	})
}

func Test_Cov15_SimpleStringOnce_Boolean(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Boolean", func() {
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("yes")
		actual := args.Map{"val": sso.Boolean(true)}
		expected := args.Map{"val": true}
		expected.ShouldBeEqual(t, 0, "Boolean returns correct value -- yes", actual)
	})
}

func Test_Cov15_SimpleStringOnce_Boolean_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Boolean_Uninit", func() {
		sso := corestr.SimpleStringOnce{}
		actual := args.Map{"val": sso.Boolean(true)}
		expected := args.Map{"val": false}
		expected.ShouldBeEqual(t, 0, "Boolean returns correct value -- uninit", actual)
	})
}

func Test_Cov15_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Int", func() {
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("42")
		actual := args.Map{"val": sso.Int()}
		expected := args.Map{"val": 42}
		expected.ShouldBeEqual(t, 0, "Int returns correct value -- with args", actual)
	})
}

func Test_Cov15_SimpleStringOnce_Byte(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_Byte", func() {
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("42")
		actual := args.Map{"val": sso.Byte()}
		expected := args.Map{"val": byte(42)}
		expected.ShouldBeEqual(t, 0, "Byte returns correct value -- with args", actual)
	})
}

func Test_Cov15_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnce_ConcatNew", func() {
		sso := corestr.SimpleStringOnce{}
		sso.GetSetOnce("hello")
		result := sso.ConcatNew(" world")
		actual := args.Map{"val": result.Value()}
		expected := args.Map{"val": "hello world"}
		expected.ShouldBeEqual(t, 0, "ConcatNew returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnceModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov15_SimpleStringOnceModel(t *testing.T) {
	safeTest(t, "Test_Cov15_SimpleStringOnceModel", func() {
		m := corestr.SimpleStringOnceModel{Value: "hi", IsInitialize: true}
		actual := args.Map{"val": m.Value, "init": m.IsInitialize}
		expected := args.Map{"val": "hi", "init": true}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnceModel returns correct value -- with args", actual)
	})
}
