package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ValidValue ──

func Test_ValidValue_Valid(t *testing.T) {
	safeTest(t, "Test_Cov8_ValidValue_Valid", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"value":   vv.Value,
			"isValid": vv.IsValid,
			"isEmpty": vv.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"value": "hello", "isValid": true, "isEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns valid -- non-empty string", actual)
	})
}

func Test_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_Cov8_ValidValue_Empty", func() {
		// Arrange
		// NewValidValue("") sets IsValid: true per implementation
		vv := corestr.NewValidValue("")

		// Act
		actual := args.Map{
			"isValid": vv.IsValid,
			"isEmpty": vv.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isValid": true,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns valid-empty -- empty string", actual)
	})
}

func Test_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov8_ValidValue_Invalid", func() {
		// Arrange
		vv := corestr.InvalidValidValue("bad input")

		// Act
		actual := args.Map{
			"isValid": vv.IsValid,
			"message": vv.Message,
		}

		// Assert
		expected := args.Map{
			"isValid": false,
			"message": "bad input",
		}
		expected.ShouldBeEqual(t, 0, "ValidValue returns invalid -- error message", actual)
	})
}

// ── LeftRight ──

func Test_LeftRight(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("l", "r")

		// Act
		actual := args.Map{
			"left":  lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "l",
			"right": "r",
		}
		expected.ShouldBeEqual(t, 0, "LeftRight returns fields -- struct access", actual)
	})
}

// ── LeftMiddleRight ──

func Test_LeftMiddleRight(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftMiddleRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")

		// Act
		actual := args.Map{
			"left":   lmr.Left,
			"middle": lmr.Middle,
			"right":  lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "l",
			"middle": "m",
			"right": "r",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRight returns fields -- struct access", actual)
	})
}

// ── LeftRightFromSplit ──

func Test_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{
			"left":  lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "key",
			"right": "value",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns split -- equals separator", actual)
	})
}

func Test_LeftRightFromSplit_NoSep(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftRightFromSplit_NoSep", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("nosep", "=")

		// Act
		actual := args.Map{
			"left": lr.Left,
			"right": lr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "nosep",
			"right": "",
		}
		expected.ShouldBeEqual(t, 0, "LeftRightFromSplit returns left-only -- no separator", actual)
	})
}

// ── LeftMiddleRightFromSplit ──

func Test_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_Cov8_LeftMiddleRightFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a:b:c", ":")

		// Act
		actual := args.Map{
			"left":   lmr.Left,
			"middle": lmr.Middle,
			"right":  lmr.Right,
		}

		// Assert
		expected := args.Map{
			"left": "a",
			"middle": "b",
			"right": "c",
		}
		expected.ShouldBeEqual(t, 0, "LeftMiddleRightFromSplit returns three parts -- colon separator", actual)
	})
}

// ── ValueStatus ──

func Test_ValueStatus(t *testing.T) {
	safeTest(t, "Test_Cov8_ValueStatus", func() {
		// Arrange
		vs := corestr.ValueStatus{
			ValueValid: &corestr.ValidValue{Value: "test", IsValid: true},
			Index:      0,
		}

		// Act
		actual := args.Map{
			"value":   vs.ValueValid.Value,
			"isValid": vs.ValueValid.IsValid,
		}

		// Assert
		expected := args.Map{
			"value": "test",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus returns fields -- struct access", actual)
	})
}

// ── TextWithLineNumber ──

func Test_TextWithLineNumber(t *testing.T) {
	safeTest(t, "Test_Cov8_TextWithLineNumber", func() {
		// Arrange
		tln := corestr.TextWithLineNumber{
			LineNumber: 5,
			Text:       "hello",
		}

		// Act
		actual := args.Map{
			"lineNo": tln.LineNumber,
			"text":   tln.Text,
		}

		// Assert
		expected := args.Map{
			"lineNo": 5,
			"text": "hello",
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber returns fields -- struct access", actual)
	})
}

// ── KeyValuePair ──

func Test_KeyValuePair(t *testing.T) {
	safeTest(t, "Test_Cov8_KeyValuePair", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{
			"key":      kv.Key,
			"value":    kv.Value,
			"notEmpty": kv.String() != "",
		}

		// Assert
		expected := args.Map{
			"key": "k",
			"value": "v",
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "KeyValuePair returns fields -- struct access", actual)
	})
}

// ── CloneSlice ──

func Test_CloneSlice(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSlice", func() {
		// Arrange
		original := []string{"a", "b", "c"}
		cloned := corestr.CloneSlice(original)

		// Act
		actual := args.Map{
			"len":   len(cloned),
			"first": cloned[0],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns copy -- valid input", actual)
	})
}

func Test_CloneSlice_Nil(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSlice_Nil", func() {
		// Arrange
		// CloneSlice(nil) returns []string{} not nil
		cloned := corestr.CloneSlice(nil)

		// Act
		actual := args.Map{"len": len(cloned)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "CloneSlice returns empty slice -- nil input", actual)
	})
}

// ── CloneSliceIf ──

func Test_CloneSliceIf_True(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSliceIf_True", func() {
		// Arrange
		original := []string{"a", "b"}
		cloned := corestr.CloneSliceIf(true, original...)

		// Act
		actual := args.Map{"len": len(cloned)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns cloned -- true flag", actual)
	})
}

func Test_CloneSliceIf_False(t *testing.T) {
	safeTest(t, "Test_Cov8_CloneSliceIf_False", func() {
		// Arrange
		original := []string{"a", "b"}
		result := corestr.CloneSliceIf(false, original...)

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf returns original -- false flag", actual)
	})
}

// ── SimpleStringOnce via New.SimpleStringOnce.Init ──

func Test_SimpleStringOnce(t *testing.T) {
	safeTest(t, "Test_Cov8_SimpleStringOnce", func() {
		// Arrange
		so := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{
			"value":         so.Value(),
			"string":        so.Value(),
			"isEmpty":       so.IsEmpty(),
			"isInitialized": so.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"value": "hello", "string": "hello", "isEmpty": false, "isInitialized": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns value -- Init", actual)
	})
}

func Test_SimpleStringOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_Cov8_SimpleStringOnce_Uninitialized", func() {
		// Arrange
		so := corestr.New.SimpleStringOnce.Uninitialized("pending")

		// Act
		actual := args.Map{
			"value":         so.Value(),
			"isInitialized": so.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"value": "pending",
			"isInitialized": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns uninitialized -- Uninitialized", actual)
	})
}

func Test_SimpleStringOnce_Empty(t *testing.T) {
	safeTest(t, "Test_Cov8_SimpleStringOnce_Empty", func() {
		// Arrange
		so := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{
			"isEmpty":       so.IsEmpty(),
			"isInitialized": so.IsInitialized(),
		}

		// Assert
		expected := args.Map{
			"isEmpty": true,
			"isInitialized": false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce returns empty -- Empty creator", actual)
	})
}
