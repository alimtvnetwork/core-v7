package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ValidValue ──

func Test_Cov8_ValidValue_Valid(t *testing.T) {
	safeTest(t, "Test_Cov8_ValidValue_Valid", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{
			"value":   vv.Value,
			"isValid": vv.IsValid,
			"isEmpty": vv.IsEmpty(),
		}
		expected := args.Map{
			"value": "hello", "isValid": true, "isEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns valid -- non-empty string", actual)
	})
}

func Test_Cov8_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov8_ValidValue_Empty", func() {
		// NewValidValue("") sets IsValid: true per implementation
		vv := corestr.NewValidValue("")
		actual := args.Map{
			"isValid": vv.IsValid,
			"isEmpty": vv.IsEmpty(),
		}
		expected := args.Map{"isValid": true, "isEmpty": true}
		expected.ShouldBeEqual(t, 0, "ValidValue returns valid-empty -- empty string", actual)
	})
}

func Test_Cov8_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov8_ValidValue_Invalid", func() {
		vv := corestr.InvalidValidValue("bad input")
		actual := args.Map{
			"isValid": vv.IsValid,
			"message": vv.Message,
		}
		expected := args.Map{"isValid": false, "message": "bad input"}
		expected.ShouldBeEqual(t, 0, "ValidValue returns invalid -- error message", actual)
	})
}

// ── LeftRight ──

func Test_Cov8_LeftRight(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftRight", func() {
		lr := corestr.NewLeftRight("l", "r")
		actual := args.Map{
			"left":  lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{"left": "l", "right": "r"}
		expected.ShouldBeEqual(t, 0, "LeftRight returns fields -- struct access", actual)
	})
}

// ── LeftMiddleRight ──

func Test_Cov8_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftMiddleRight", func() {
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")
		actual := args.Map{
			"left":   lmr.Left,
			"middle": lmr.Middle,
			"right":  lmr.Right,
		}
		expected := args.Map{"left": "l", "middle": "m", "right": "r"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns fields -- struct access", actual)
	})
}

// ── LeftRightFromSplit ──

func Test_Cov8_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftRightFromSplit", func() {
		lr := corestr.LeftRightFromSplit("key=value", "=")
		actual := args.Map{
			"left":  lr.Left,
			"right": lr.Right,
		}
		expected := args.Map{"left": "key", "right": "value"}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns split -- equals separator", actual)
	})
}

func Test_Cov8_LeftRightFromSplit_NoSep(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftRightFromSplit_NoSep", func() {
		lr := corestr.LeftRightFromSplit("nosep", "=")
		actual := args.Map{"left": lr.Left, "right": lr.Right}
		expected := args.Map{"left": "nosep", "right": ""}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns left-only -- no separator", actual)
	})
}

// ── LeftMiddleRightFromSplit ──

func Test_Cov8_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftMiddleRightFromSplit", func() {
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")
		actual := args.Map{
			"left":   lmr.Left,
			"middle": lmr.Middle,
			"right":  lmr.Right,
		}
		expected := args.Map{"left": "a", "middle": "b", "right": "c"}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns three parts -- colon separator", actual)
	})
}

// ── ValueStatus ──

func Test_Cov8_ValueStatus(t *testing.T) {
	safeTest(t, "Test_Cov8_ValueStatus", func() {
		vs := corestr.ValueStatus{
			ValueValid: &corestr.ValidValue{Value: "test", IsValid: true},
			Index:      0,
		}
		actual := args.Map{
			"value":   vs.ValueValid.Value,
			"isValid": vs.ValueValid.IsValid,
		}
		expected := args.Map{"value": "test", "isValid": true}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns fields -- struct access", actual)
	})
}

// ── TextWithLineNumber ──

func Test_Cov8_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov8_TextWithLineNumber", func() {
		tln := corestr.TextWithLineNumber{
			LineNumber: 5,
			Text:       "hello",
		}
		actual := args.Map{
			"lineNo": tln.LineNumber,
			"text":   tln.Text,
		}
		expected := args.Map{"lineNo": 5, "text": "hello"}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns fields -- struct access", actual)
	})
}

// ── KeyValuePair ──

func Test_Cov8_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov8_KeyValuePair", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{
			"key":      kv.Key,
			"value":    kv.Value,
			"notEmpty": kv.String() != "",
		}
		expected := args.Map{"key": "k", "value": "v", "notEmpty": true}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns fields -- struct access", actual)
	})
}

// ── CloneSlice ──

func Test_Cov8_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSlice", func() {
		original := []string{"a", "b", "c"}
		cloned := corestr.CloneSlice(original)
		actual := args.Map{
			"len":   len(cloned),
			"first": cloned[0],
		}
		expected := args.Map{"len": 3, "first": "a"}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns copy -- valid input", actual)
	})
}

func Test_Cov8_CloneSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSlice_Nil", func() {
		// CloneSlice(nil) returns []string{} not nil
		cloned := corestr.CloneSlice(nil)
		actual := args.Map{"len": len(cloned)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty slice -- nil input", actual)
	})
}

// ── CloneSliceIf ──

func Test_Cov8_CloneSliceIf_True(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSliceIf_True", func() {
		original := []string{"a", "b"}
		cloned := corestr.CloneSliceIf(true, original...)
		actual := args.Map{"len": len(cloned)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns cloned -- true flag", actual)
	})
}

func Test_Cov8_CloneSliceIf_False(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSliceIf_False", func() {
		original := []string{"a", "b"}
		result := corestr.CloneSliceIf(false, original...)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns original -- false flag", actual)
	})
}

// ── SimpleStringOnce via New.SimpleStringOnce.Init ──

func Test_Cov8_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov8_SimpleStringOnce", func() {
		so := corestr.New.SimpleStringOnce.Init("hello")
		actual := args.Map{
			"value":         so.Value(),
			"string":        so.Value(),
			"isEmpty":       so.IsEmpty(),
			"isInitialized": so.IsInitialized(),
		}
		expected := args.Map{
			"value": "hello", "string": "hello", "isEmpty": false, "isInitialized": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns value -- Init", actual)
	})
}

func Test_Cov8_SimpleStringOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_Cov8_SimpleStringOnce_Uninitialized", func() {
		so := corestr.New.SimpleStringOnce.Uninitialized("pending")
		actual := args.Map{
			"value":         so.Value(),
			"isInitialized": so.IsInitialized(),
		}
		expected := args.Map{"value": "pending", "isInitialized": false}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns uninitialized -- Uninitialized", actual)
	})
}

func Test_Cov8_SimpleStringOnce_Empty(t *testing.T) {
	safeTest(t, "Test_Cov8_SimpleStringOnce_Empty", func() {
		so := corestr.New.SimpleStringOnce.Empty()
		actual := args.Map{
			"isEmpty":       so.IsEmpty(),
			"isInitialized": so.IsInitialized(),
		}
		expected := args.Map{"isEmpty": true, "isInitialized": false}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns empty -- Empty creator", actual)
	})
}
