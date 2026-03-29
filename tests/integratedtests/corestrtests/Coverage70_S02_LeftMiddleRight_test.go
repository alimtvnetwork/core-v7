package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — Constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov70_InvalidLeftMiddleRightNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov70_InvalidLeftMiddleRightNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{"isValid": lmr.IsValid, "left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"isValid": false, "left": "", "middle": "", "right": ""}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRightNoMessage returns invalid -- no message", actual)
	})
}

func Test_Cov70_InvalidLeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_Cov70_InvalidLeftMiddleRight", func() {
		lmr := corestr.InvalidLeftMiddleRight("err msg")
		actual := args.Map{"isValid": lmr.IsValid, "hasMsg": lmr.Message != ""}
		expected := args.Map{"isValid": false, "hasMsg": true}
		expected.ShouldBeEqual(t, 0, "InvalidLeftMiddleRight returns invalid -- with message", actual)
	})
}

func Test_Cov70_NewLeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_Cov70_NewLeftMiddleRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"isValid": lmr.IsValid, "left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"isValid": true, "left": "a", "middle": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight returns valid -- three strings", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — Bytes methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov70_LMR_LeftBytes(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_LeftBytes", func() {
		lmr := corestr.NewLeftMiddleRight("abc", "d", "e")
		actual := args.Map{"len": len(lmr.LeftBytes())}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "LeftBytes returns bytes -- valid left", actual)
	})
}

func Test_Cov70_LMR_RightBytes(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_RightBytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "cde")
		actual := args.Map{"len": len(lmr.RightBytes())}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "RightBytes returns bytes -- valid right", actual)
	})
}

func Test_Cov70_LMR_MiddleBytes(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_MiddleBytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "mid", "c")
		actual := args.Map{"len": len(lmr.MiddleBytes())}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "MiddleBytes returns bytes -- valid middle", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — Trim methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov70_LMR_LeftTrim(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_LeftTrim", func() {
		lmr := corestr.NewLeftMiddleRight(" x ", "m", "r")
		actual := args.Map{"val": lmr.LeftTrim()}
		expected := args.Map{"val": "x"}
		expected.ShouldBeEqual(t, 0, "LeftTrim returns trimmed -- whitespace left", actual)
	})
}

func Test_Cov70_LMR_RightTrim(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_RightTrim", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", " y ")
		actual := args.Map{"val": lmr.RightTrim()}
		expected := args.Map{"val": "y"}
		expected.ShouldBeEqual(t, 0, "RightTrim returns trimmed -- whitespace right", actual)
	})
}

func Test_Cov70_LMR_MiddleTrim(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_MiddleTrim", func() {
		lmr := corestr.NewLeftMiddleRight("l", " z ", "r")
		actual := args.Map{"val": lmr.MiddleTrim()}
		expected := args.Map{"val": "z"}
		expected.ShouldBeEqual(t, 0, "MiddleTrim returns trimmed -- whitespace middle", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — Empty/Whitespace checks
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov70_LMR_IsLeftEmpty(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsLeftEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("", "m", "r")
		actual := args.Map{"empty": lmr.IsLeftEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsLeftEmpty returns true -- empty left", actual)
	})
}

func Test_Cov70_LMR_IsRightEmpty(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsRightEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "")
		actual := args.Map{"empty": lmr.IsRightEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsRightEmpty returns true -- empty right", actual)
	})
}

func Test_Cov70_LMR_IsMiddleEmpty(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsMiddleEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("l", "", "r")
		actual := args.Map{"empty": lmr.IsMiddleEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsMiddleEmpty returns true -- empty middle", actual)
	})
}

func Test_Cov70_LMR_IsMiddleWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsMiddleWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("l", "  ", "r")
		actual := args.Map{"ws": lmr.IsMiddleWhitespace()}
		expected := args.Map{"ws": true}
		expected.ShouldBeEqual(t, 0, "IsMiddleWhitespace returns true -- whitespace middle", actual)
	})
}

func Test_Cov70_LMR_IsLeftWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsLeftWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "m", "r")
		actual := args.Map{"ws": lmr.IsLeftWhitespace()}
		expected := args.Map{"ws": true}
		expected.ShouldBeEqual(t, 0, "IsLeftWhitespace returns true -- whitespace left", actual)
	})
}

func Test_Cov70_LMR_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsRightWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "  ")
		actual := args.Map{"ws": lmr.IsRightWhitespace()}
		expected := args.Map{"ws": true}
		expected.ShouldBeEqual(t, 0, "IsRightWhitespace returns true -- whitespace right", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — HasValid* methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov70_LMR_HasValidNonEmptyLeft(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonEmptyLeft", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyLeft returns true -- valid non-empty", actual)
	})
}

func Test_Cov70_LMR_HasValidNonEmptyLeft_Empty(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonEmptyLeft_Empty", func() {
		lmr := corestr.NewLeftMiddleRight("", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyLeft returns false -- empty left", actual)
	})
}

func Test_Cov70_LMR_HasValidNonEmptyRight(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonEmptyRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyRight returns true -- valid non-empty", actual)
	})
}

func Test_Cov70_LMR_HasValidNonEmptyMiddle(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonEmptyMiddle", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyMiddle()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyMiddle returns true -- valid non-empty", actual)
	})
}

func Test_Cov70_LMR_HasValidNonEmptyMiddle_Empty(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonEmptyMiddle_Empty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyMiddle()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonEmptyMiddle returns false -- empty middle", actual)
	})
}

func Test_Cov70_LMR_HasValidNonWhitespaceLeft(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonWhitespaceLeft", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonWhitespaceLeft()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceLeft returns true -- non-ws", actual)
	})
}

func Test_Cov70_LMR_HasValidNonWhitespaceLeft_Ws(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonWhitespaceLeft_Ws", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonWhitespaceLeft()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceLeft returns false -- ws left", actual)
	})
}

func Test_Cov70_LMR_HasValidNonWhitespaceRight(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonWhitespaceRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonWhitespaceRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceRight returns true -- non-ws", actual)
	})
}

func Test_Cov70_LMR_HasValidNonWhitespaceMiddle(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonWhitespaceMiddle", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonWhitespaceMiddle()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceMiddle returns true -- non-ws", actual)
	})
}

func Test_Cov70_LMR_HasValidNonWhitespaceMiddle_Ws(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasValidNonWhitespaceMiddle_Ws", func() {
		lmr := corestr.NewLeftMiddleRight("a", " ", "c")
		actual := args.Map{"result": lmr.HasValidNonWhitespaceMiddle()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasValidNonWhitespaceMiddle returns false -- ws middle", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LeftMiddleRight — HasSafeNonEmpty, IsAll, Is, Clone, ToLeftRight, Clear, Dispose
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov70_LMR_HasSafeNonEmpty_True(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasSafeNonEmpty_True", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns true -- all non-empty", actual)
	})
}

func Test_Cov70_LMR_HasSafeNonEmpty_EmptyMiddle(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_HasSafeNonEmpty_EmptyMiddle", func() {
		lmr := corestr.NewLeftMiddleRight("a", "", "c")
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "HasSafeNonEmpty returns false -- empty middle", actual)
	})
}

func Test_Cov70_LMR_IsAll(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"match": lmr.IsAll("a", "b", "c"), "noMatch": lmr.IsAll("a", "x", "c")}
		expected := args.Map{"match": true, "noMatch": false}
		expected.ShouldBeEqual(t, 0, "IsAll returns correct -- match and mismatch", actual)
	})
}

func Test_Cov70_LMR_Is(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_Is", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"match": lmr.Is("a", "c"), "noMatch": lmr.Is("a", "x")}
		expected := args.Map{"match": true, "noMatch": false}
		expected.ShouldBeEqual(t, 0, "Is returns correct -- left+right match and mismatch", actual)
	})
}

func Test_Cov70_LMR_Clone(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()
		actual := args.Map{"left": c.Left, "middle": c.Middle, "right": c.Right, "notSame": c != lmr}
		expected := args.Map{"left": "a", "middle": "b", "right": "c", "notSame": true}
		expected.ShouldBeEqual(t, 0, "Clone returns copy -- valid LMR", actual)
	})
}

func Test_Cov70_LMR_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		actual := args.Map{"left": lr.Left, "right": lr.Right, "isValid": lr.IsValid}
		expected := args.Map{"left": "a", "right": "c", "isValid": true}
		expected.ShouldBeEqual(t, 0, "ToLeftRight returns LR -- drops middle", actual)
	})
}

func Test_Cov70_LMR_Clear(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_Clear", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()
		actual := args.Map{"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "", "middle": "", "right": ""}
		expected.ShouldBeEqual(t, 0, "Clear zeroes fields -- valid LMR", actual)
	})
}

func Test_Cov70_LMR_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_Clear_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Clear() // must not panic
		actual := args.Map{"noPanic": true}
		expected := args.Map{"noPanic": true}
		expected.ShouldBeEqual(t, 0, "Clear returns safely -- nil receiver", actual)
	})
}

func Test_Cov70_LMR_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_Dispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()
		actual := args.Map{"left": lmr.Left, "middle": lmr.Middle, "right": lmr.Right}
		expected := args.Map{"left": "", "middle": "", "right": ""}
		expected.ShouldBeEqual(t, 0, "Dispose clears fields -- valid LMR", actual)
	})
}

func Test_Cov70_LMR_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov70_LMR_Dispose_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Dispose() // must not panic
		actual := args.Map{"noPanic": true}
		expected := args.Map{"noPanic": true}
		expected.ShouldBeEqual(t, 0, "Dispose returns safely -- nil receiver", actual)
	})
}
