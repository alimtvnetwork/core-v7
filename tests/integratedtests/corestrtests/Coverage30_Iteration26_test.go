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

func Test_I26_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_NewLeftRight", func() {
		lr := corestr.NewLeftRight("key", "value")
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "key", "right": "value", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewLeftRight returns correct value -- with args", actual)
	})
}

func Test_I26_LeftRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_InvalidNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRightNoMessage returns error -- with args", actual)
	})
}

func Test_I26_LeftRight_InvalidWithMessage(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_InvalidWithMessage", func() {
		lr := corestr.InvalidLeftRight("bad")
		actual := args.Map{"valid": lr.IsValid, "msg": lr.Message}
		expected := args.Map{"valid": false, "msg": "bad"}
		expected.ShouldBeEqual(t, 0, "InvalidLeftRight returns error -- with args", actual)
	})
}

func Test_I26_LeftRight_UsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlice_Two", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "b", "valid": true}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- two", actual)
	})
}

func Test_I26_LeftRight_UsingSlice_One(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlice_One", func() {
		lr := corestr.LeftRightUsingSlice([]string{"only"})
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "only", "right": "", "valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns correct value -- one", actual)
	})
}

func Test_I26_LeftRight_UsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice([]string{})
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlice returns empty -- empty", actual)
	})
}

func Test_I26_LeftRight_UsingSlicePtr_Empty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlicePtr_Empty", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{})
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns empty -- empty", actual)
	})
}

func Test_I26_LeftRight_UsingSlicePtr_NonEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_UsingSlicePtr_NonEmpty", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRightUsingSlicePtr returns empty -- non-empty", actual)
	})
}

func Test_I26_LeftRight_TrimmedUsingSlice_Two(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_Two", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- two", actual)
	})
}

func Test_I26_LeftRight_TrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_Nil", func() {
		lr := corestr.LeftRightTrimmedUsingSlice(nil)
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns nil -- nil", actual)
	})
}

func Test_I26_LeftRight_TrimmedUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_Empty", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})
		actual := args.Map{"valid": lr.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns empty -- empty", actual)
	})
}

func Test_I26_LeftRight_TrimmedUsingSlice_One(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_TrimmedUsingSlice_One", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" only "})
		actual := args.Map{"left": lr.Left, "valid": lr.IsValid}
		expected := args.Map{"left": "only", "valid": false}
		expected.ShouldBeEqual(t, 0, "LeftRightTrimmedUsingSlice returns correct value -- one", actual)
	})
}

func Test_I26_LeftRight_Bytes(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Bytes", func() {
		lr := corestr.NewLeftRight("abc", "xyz")
		actual := args.Map{"leftLen": len(lr.LeftBytes()), "rightLen": len(lr.RightBytes())}
		expected := args.Map{"leftLen": 3, "rightLen": 3}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Bytes", actual)
	})
}

func Test_I26_LeftRight_Trim(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Trim", func() {
		lr := corestr.NewLeftRight(" a ", " b ")
		actual := args.Map{"left": lr.LeftTrim(), "right": lr.RightTrim()}
		expected := args.Map{"left": "a", "right": "b"}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Trim", actual)
	})
}

func Test_I26_LeftRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_IsEmpty", func() {
		lr := corestr.NewLeftRight("", "x")
		actual := args.Map{"leftEmpty": lr.IsLeftEmpty(), "rightEmpty": lr.IsRightEmpty()}
		expected := args.Map{"leftEmpty": true, "rightEmpty": false}
		expected.ShouldBeEqual(t, 0, "LeftRight returns empty -- IsEmpty", actual)
	})
}

func Test_I26_LeftRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_IsWhitespace", func() {
		lr := corestr.NewLeftRight("  ", "x")
		actual := args.Map{"leftWs": lr.IsLeftWhitespace(), "rightWs": lr.IsRightWhitespace()}
		expected := args.Map{"leftWs": true, "rightWs": false}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- IsWhitespace", actual)
	})
}

func Test_I26_LeftRight_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_HasValidNonEmpty", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{
			"left":  lr.HasValidNonEmptyLeft(),
			"right": lr.HasValidNonEmptyRight(),
			"safe":  lr.HasSafeNonEmpty(),
		}
		expected := args.Map{"left": true, "right": true, "safe": true}
		expected.ShouldBeEqual(t, 0, "LeftRight returns empty -- HasValidNonEmpty", actual)
	})
}

func Test_I26_LeftRight_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_HasValidNonWhitespace", func() {
		lr := corestr.NewLeftRight("a", " ")
		actual := args.Map{
			"leftNWS":  lr.HasValidNonWhitespaceLeft(),
			"rightNWS": lr.HasValidNonWhitespaceRight(),
		}
		expected := args.Map{"leftNWS": true, "rightNWS": false}
		expected.ShouldBeEqual(t, 0, "LeftRight returns non-empty -- HasValidNonWhitespace", actual)
	})
}

func Test_I26_LeftRight_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_NonPtr_Ptr", func() {
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		p := lr.Ptr()
		actual := args.Map{"npLeft": np.Left, "pSame": p == lr}
		expected := args.Map{"npLeft": "a", "pSame": true}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- NonPtr/Ptr", actual)
	})
}

func Test_I26_LeftRight_RegexMatch(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_RegexMatch", func() {
		lr := corestr.NewLeftRight("abc123", "xyz")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{
			"leftMatch":  lr.IsLeftRegexMatch(re),
			"rightMatch": lr.IsRightRegexMatch(re),
			"nilRegex":   lr.IsLeftRegexMatch(nil),
		}
		expected := args.Map{"leftMatch": true, "rightMatch": false, "nilRegex": false}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- RegexMatch", actual)
	})
}

func Test_I26_LeftRight_Is_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Is_IsKey_IsVal", func() {
		lr := corestr.NewLeftRight("k", "v")
		actual := args.Map{
			"is":     lr.Is("k", "v"),
			"isNot":  lr.Is("k", "x"),
			"isLeft": lr.IsLeft("k"),
			"isRight": lr.IsRight("v"),
		}
		expected := args.Map{"is": true, "isNot": false, "isLeft": true, "isRight": true}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Is/IsLeft/IsRight", actual)
	})
}

func Test_I26_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		lr3 := corestr.NewLeftRight("a", "c")
		actual := args.Map{
			"same":    lr1.IsEqual(lr2),
			"diff":    lr1.IsEqual(lr3),
			"nilBoth": (*corestr.LeftRight)(nil).IsEqual(nil),
			"nilOne":  lr1.IsEqual(nil),
		}
		expected := args.Map{"same": true, "diff": false, "nilBoth": true, "nilOne": false}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- IsEqual", actual)
	})
}

func Test_I26_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		cloned := lr.Clone()
		actual := args.Map{"left": cloned.Left, "right": cloned.Right, "notSame": cloned != lr}
		expected := args.Map{"left": "a", "right": "b", "notSame": true}
		expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clone", actual)
	})
}

func Test_I26_LeftRight_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I26_LeftRight_Clear_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "", "right": ""}
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

func Test_I26_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_I26_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_I26_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b:c:d"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFull returns correct value -- with args", actual)
	})
}

func Test_I26_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftRightFromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "a", "right": "b : c"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplitFullTrimmed returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right, "valid": lmr.IsValid}
		expected := args.Map{"left": "a", "mid": "b", "right": "c", "valid": true}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight returns correct value -- with args", actual)
	})
}

func Test_I26_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Invalid", func() {
		lmr1 := corestr.InvalidLeftMiddleRightNoMessage()
		lmr2 := corestr.InvalidLeftMiddleRight("err")
		actual := args.Map{"v1": lmr1.IsValid, "v2": lmr2.IsValid, "msg": lmr2.Message}
		expected := args.Map{"v1": false, "v2": false, "msg": "err"}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight returns error -- with args", actual)
	})
}

func Test_I26_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Bytes", func() {
		lmr := corestr.NewLeftMiddleRight("ab", "cd", "ef")
		actual := args.Map{"lLen": len(lmr.LeftBytes()), "mLen": len(lmr.MiddleBytes()), "rLen": len(lmr.RightBytes())}
		expected := args.Map{"lLen": 2, "mLen": 2, "rLen": 2}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Bytes", actual)
	})
}

func Test_I26_LeftMiddleRight_Trim(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Trim", func() {
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
		actual := args.Map{"l": lmr.LeftTrim(), "m": lmr.MiddleTrim(), "r": lmr.RightTrim()}
		expected := args.Map{"l": "a", "m": "b", "r": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Trim", actual)
	})
}

func Test_I26_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_IsEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("", "x", "")
		actual := args.Map{"lEmpty": lmr.IsLeftEmpty(), "mEmpty": lmr.IsMiddleEmpty(), "rEmpty": lmr.IsRightEmpty()}
		expected := args.Map{"lEmpty": true, "mEmpty": false, "rEmpty": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns empty -- IsEmpty", actual)
	})
}

func Test_I26_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_IsWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "x", "  ")
		actual := args.Map{"lWs": lmr.IsLeftWhitespace(), "mWs": lmr.IsMiddleWhitespace(), "rWs": lmr.IsRightWhitespace()}
		expected := args.Map{"lWs": true, "mWs": false, "rWs": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- IsWhitespace", actual)
	})
}

func Test_I26_LeftMiddleRight_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_HasValidNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{
			"l": lmr.HasValidNonEmptyLeft(), "m": lmr.HasValidNonEmptyMiddle(), "r": lmr.HasValidNonEmptyRight(),
			"safe": lmr.HasSafeNonEmpty(),
		}
		expected := args.Map{"l": true, "m": true, "r": true, "safe": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns empty -- HasValidNonEmpty", actual)
	})
}

func Test_I26_LeftMiddleRight_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_HasValidNonWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("a", " ", "c")
		actual := args.Map{
			"l": lmr.HasValidNonWhitespaceLeft(), "m": lmr.HasValidNonWhitespaceMiddle(), "r": lmr.HasValidNonWhitespaceRight(),
		}
		expected := args.Map{"l": true, "m": false, "r": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns non-empty -- HasValidNonWhitespace", actual)
	})
}

func Test_I26_LeftMiddleRight_IsAll_Is(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_IsAll_Is", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"isAll": lmr.IsAll("a", "b", "c"), "is": lmr.Is("a", "c")}
		expected := args.Map{"isAll": true, "is": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- IsAll/Is", actual)
	})
}

func Test_I26_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()
		actual := args.Map{"left": cloned.Left, "notSame": cloned != lmr}
		expected := args.Map{"left": "a", "notSame": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- Clone", actual)
	})
}

func Test_I26_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		actual := args.Map{"left": lr.Left, "right": lr.Right, "valid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "c", "valid": true}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns correct value -- ToLeftRight", actual)
	})
}

func Test_I26_LeftMiddleRight_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRight_Clear_Dispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()
		actual := args.Map{"left": lmr.Left}
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

func Test_I26_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns correct value -- with args", actual)
	})
}

func Test_I26_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplitTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitTrimmed returns correct value -- with args", actual)
	})
}

func Test_I26_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplitN", func() {
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "a", "mid": "b", "right": "c:d:e"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitN returns correct value -- with args", actual)
	})
}

func Test_I26_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_I26_LeftMiddleRightFromSplitNTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
		actual := args.Map{"left": lmr.Left, "mid": lmr.Middle}
		expected := args.Map{"left": "a", "mid": "b"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplitNTrimmed returns correct value -- with args", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValuePair
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Basic", func() {
		kv := corestr.KeyValuePair{Key: "name", Value: "alice"}
		actual := args.Map{
			"key": kv.KeyName(), "varName": kv.VariableName(), "val": kv.ValueString(),
			"isVarEq": kv.IsVariableNameEqual("name"), "isValEq": kv.IsValueEqual("alice"),
		}
		expected := args.Map{"key": "name", "varName": "name", "val": "alice", "isVarEq": true, "isValEq": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- basic", actual)
	})
}

func Test_I26_KeyValuePair_Json(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Json", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		jr := kv.Json()
		actual := args.Map{"noErr": !jr.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Json", actual)
	})
}

func Test_I26_KeyValuePair_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_JsonPtr", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"notNil": kv.JsonPtr() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- JsonPtr", actual)
	})
}

func Test_I26_KeyValuePair_Serialize(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Serialize", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b, err := kv.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Serialize", actual)
	})
}

func Test_I26_KeyValuePair_SerializeMust(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_SerializeMust", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		b := kv.SerializeMust()
		actual := args.Map{"hasBytes": len(b) > 0}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- SerializeMust", actual)
	})
}

func Test_I26_KeyValuePair_Compile_String(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Compile_String", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"compile": kv.Compile(), "str": kv.String()}
		expected := args.Map{"compile": "{k:v}", "str": "{k:v}"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Compile/String", actual)
	})
}

func Test_I26_KeyValuePair_EmptyChecks(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_EmptyChecks", func() {
		kv := corestr.KeyValuePair{Key: "", Value: ""}
		actual := args.Map{
			"keyEmpty": kv.IsKeyEmpty(), "valEmpty": kv.IsValueEmpty(),
			"hasKey": kv.HasKey(), "hasVal": kv.HasValue(),
			"kvEmpty": kv.IsKeyValueEmpty(), "kvAnyEmpty": kv.IsKeyValueAnyEmpty(),
		}
		expected := args.Map{
			"keyEmpty": true, "valEmpty": true,
			"hasKey": false, "hasVal": false,
			"kvEmpty": true, "kvAnyEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns empty -- EmptyChecks", actual)
	})
}

func Test_I26_KeyValuePair_Trim(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Trim", func() {
		kv := corestr.KeyValuePair{Key: " k ", Value: " v "}
		actual := args.Map{"key": kv.TrimKey(), "val": kv.TrimValue()}
		expected := args.Map{"key": "k", "val": "v"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Trim", actual)
	})
}

func Test_I26_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueBool", func() {
		kv1 := corestr.KeyValuePair{Value: "true"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		kv3 := corestr.KeyValuePair{Value: ""}
		actual := args.Map{"t": kv1.ValueBool(), "f": kv2.ValueBool(), "empty": kv3.ValueBool()}
		expected := args.Map{"t": true, "f": false, "empty": false}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueBool", actual)
	})
}

func Test_I26_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueInt", func() {
		kv1 := corestr.KeyValuePair{Value: "42"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		actual := args.Map{"val": kv1.ValueInt(0), "def": kv2.ValueInt(99), "defInt": kv1.ValueDefInt()}
		expected := args.Map{"val": 42, "def": 99, "defInt": 42}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueInt", actual)
	})
}

func Test_I26_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueByte", func() {
		kv1 := corestr.KeyValuePair{Value: "100"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		kv3 := corestr.KeyValuePair{Value: "300"}
		actual := args.Map{"val": kv1.ValueByte(0), "def": kv2.ValueByte(7), "overflow": kv3.ValueByte(5)}
		expected := args.Map{"val": byte(100), "def": byte(7), "overflow": byte(5)}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueByte", actual)
	})
}

func Test_I26_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueDefByte", func() {
		kv1 := corestr.KeyValuePair{Value: "50"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		kv3 := corestr.KeyValuePair{Value: "999"}
		actual := args.Map{"val": kv1.ValueDefByte(), "err": kv2.ValueDefByte(), "over": kv3.ValueDefByte()}
		expected := args.Map{"val": byte(50), "err": byte(0), "over": byte(0)}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueDefByte", actual)
	})
}

func Test_I26_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueFloat64", func() {
		kv1 := corestr.KeyValuePair{Value: "3.14"}
		kv2 := corestr.KeyValuePair{Value: "abc"}
		actual := args.Map{"close": kv1.ValueFloat64(0) > 3.1, "def": kv2.ValueFloat64(1.0), "defFloat": kv1.ValueDefFloat64() > 3.1}
		expected := args.Map{"close": true, "def": 1.0, "defFloat": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- ValueFloat64", actual)
	})
}

func Test_I26_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueValid", func() {
		kv := corestr.KeyValuePair{Value: "test"}
		vv := kv.ValueValid()
		actual := args.Map{"val": vv.Value, "valid": vv.IsValid}
		expected := args.Map{"val": "test", "valid": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueValid", actual)
	})
}

func Test_I26_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Value: "test"}
		vv := kv.ValueValidOptions(false, "bad")
		actual := args.Map{"valid": vv.IsValid, "msg": vv.Message}
		expected := args.Map{"valid": false, "msg": "bad"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns non-empty -- ValueValidOptions", actual)
	})
}

func Test_I26_KeyValuePair_Is_IsKey_IsVal(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Is_IsKey_IsVal", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"is": kv.Is("k", "v"), "isKey": kv.IsKey("k"), "isVal": kv.IsVal("v")}
		expected := args.Map{"is": true, "isKey": true, "isVal": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Is/IsKey/IsVal", actual)
	})
}

func Test_I26_KeyValuePair_FormatString(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_FormatString", func() {
		kv := &corestr.KeyValuePair{Key: "name", Value: "bob"}
		actual := args.Map{"fmt": kv.FormatString("%s=%s")}
		expected := args.Map{"fmt": "name=bob"}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- FormatString", actual)
	})
}

func Test_I26_KeyValuePair_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_Clear_Dispose", func() {
		kv := &corestr.KeyValuePair{Key: "k", Value: "v"}
		kv.Clear()
		actual := args.Map{"key": kv.Key, "val": kv.Value}
		expected := args.Map{"key": "", "val": ""}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns correct value -- Clear", actual)

		kv2 := &corestr.KeyValuePair{Key: "x", Value: "y"}
		kv2.Dispose()
		(*corestr.KeyValuePair)(nil).Clear()
		(*corestr.KeyValuePair)(nil).Dispose()
	})
}

func Test_I26_KeyValuePair_NilChecks(t *testing.T) {
	safeTest(t, "Test_I26_KeyValuePair_NilChecks", func() {
		var kv *corestr.KeyValuePair
		actual := args.Map{"anyEmpty": kv.IsKeyValueAnyEmpty()}
		expected := args.Map{"anyEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns nil -- nil checks", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_TextWithLineNumber_Valid(t *testing.T) {
	safeTest(t, "Test_I26_TextWithLineNumber_Valid", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		actual := args.Map{
			"hasLine": tl.HasLineNumber(), "invalid": tl.IsInvalidLineNumber(),
			"len": tl.Length(), "empty": tl.IsEmpty(), "emptyText": tl.IsEmptyText(),
			"emptyBoth": tl.IsEmptyTextLineBoth(),
		}
		expected := args.Map{
			"hasLine": true, "invalid": false,
			"len": 5, "empty": false, "emptyText": false,
			"emptyBoth": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns non-empty -- valid", actual)
	})
}

func Test_I26_TextWithLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_I26_TextWithLineNumber_Nil", func() {
		var tl *corestr.TextWithLineNumber
		actual := args.Map{
			"hasLine": tl.HasLineNumber(), "invalid": tl.IsInvalidLineNumber(),
			"len": tl.Length(), "empty": tl.IsEmpty(), "emptyText": tl.IsEmptyText(),
		}
		expected := args.Map{
			"hasLine": false, "invalid": true,
			"len": 0, "empty": true, "emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns nil -- nil", actual)
	})
}

func Test_I26_TextWithLineNumber_EmptyText(t *testing.T) {
	safeTest(t, "Test_I26_TextWithLineNumber_EmptyText", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: ""}
		actual := args.Map{"empty": tl.IsEmpty(), "emptyText": tl.IsEmptyText()}
		expected := args.Map{"empty": true, "emptyText": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns empty -- empty text", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// CloneSlice / CloneSliceIf
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_CloneSlice_Normal(t *testing.T) {
	safeTest(t, "Test_I26_CloneSlice_Normal", func() {
		s := []string{"a", "b", "c"}
		cloned := corestr.CloneSlice(s)
		actual := args.Map{"len": len(cloned), "first": cloned[0]}
		expected := args.Map{"len": 3, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns correct value -- normal", actual)
	})
}

func Test_I26_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I26_CloneSlice_Empty", func() {
		cloned := corestr.CloneSlice([]string{})
		actual := args.Map{"len": len(cloned)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty -- empty", actual)
	})
}

func Test_I26_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_I26_CloneSliceIf_Clone", func() {
		cloned := corestr.CloneSliceIf(true, "a", "b")
		actual := args.Map{"len": len(cloned), "first": cloned[0]}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns correct value -- clone", actual)
	})
}

func Test_I26_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_I26_CloneSliceIf_NoClone", func() {
		result := corestr.CloneSliceIf(false, "a", "b")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- no clone", actual)
	})
}

func Test_I26_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_I26_CloneSliceIf_Empty", func() {
		result := corestr.CloneSliceIf(true)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns empty -- empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_I26_AnyToString_WithFieldName", func() {
		result := corestr.AnyToString(true, 42)
		actual := args.Map{"notEmpty": result != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- with field name", actual)
	})
}

func Test_I26_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_I26_AnyToString_WithoutFieldName", func() {
		result := corestr.AnyToString(false, 42)
		actual := args.Map{"notEmpty": result != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString returns non-empty -- without field name", actual)
	})
}

func Test_I26_AnyToString_EmptyString(t *testing.T) {
	safeTest(t, "Test_I26_AnyToString_EmptyString", func() {
		result := corestr.AnyToString(false, "")
		actual := args.Map{"empty": result}
		expected := args.Map{"empty": ""}
		expected.ShouldBeEqual(t, 0, "AnyToString returns empty -- empty string", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength / AllIndividualsLengthOfSimpleSlices
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_AllIndividualStringsOfStringsLength_Valid(t *testing.T) {
	safeTest(t, "Test_I26_AllIndividualStringsOfStringsLength_Valid", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(&items)}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns non-empty -- valid", actual)
	})
}

func Test_I26_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_I26_AllIndividualStringsOfStringsLength_Nil", func() {
		actual := args.Map{"len": corestr.AllIndividualStringsOfStringsLength(nil)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength returns nil -- nil", actual)
	})
}

func Test_I26_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_I26_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		actual := args.Map{"len": corestr.AllIndividualsLengthOfSimpleSlices(nil)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualsLengthOfSimpleSlices returns nil -- nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_I26_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("bad")
		actual := args.Map{"valid": vs.ValueValid.IsValid, "msg": vs.ValueValid.Message}
		expected := args.Map{"valid": false, "msg": "bad"}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns error -- invalid", actual)
	})
}

func Test_I26_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_I26_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"valid": vs.ValueValid.IsValid}
		expected := args.Map{"valid": false}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns empty -- invalid no msg", actual)
	})
}

func Test_I26_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_I26_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("test")
		cloned := vs.Clone()
		actual := args.Map{"msg": cloned.ValueValid.Message, "notSame": cloned != vs}
		expected := args.Map{"msg": "test", "notSame": true}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- Clone", actual)
	})
}
