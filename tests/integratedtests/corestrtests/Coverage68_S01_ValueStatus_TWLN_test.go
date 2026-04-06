package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ValueStatus ──

func Test_S01_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValueStatus", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("bad")

		// Act
		actual := args.Map{
			"valid":   vs.ValueValid.IsValid,
			"msg":     vs.ValueValid.Message,
			"index":   vs.Index,
			"valEmpty": vs.ValueValid.Value == "",
		}

		// Assert
		expected := args.Map{
			"valid":   false,
			"msg":     "bad",
			"index":   -1,
			"valEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatus returns correct value -- with message", actual)
	})
}

func Test_S01_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValueStatusNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{"valid": vs.ValueValid.IsValid, "msg": vs.ValueValid.Message}

		// Assert
		expected := args.Map{"valid": false, "msg": ""}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns correct value -- no message", actual)
	})
}

func Test_S01_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_S01_ValueStatus_Clone", func() {
		// Arrange
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("hello"),
			Index:      5,
		}
		c := vs.Clone()

		// Act
		actual := args.Map{"val": c.ValueValid.Value, "index": c.Index, "samePtr": vs == c}

		// Assert
		expected := args.Map{"val": "hello", "index": 5, "samePtr": false}
		expected.ShouldBeEqual(t, 0, "ValueStatus.Clone returns correct value -- deep copy", actual)
	})
}

// ── TextWithLineNumber ──

func Test_S01_TWLN_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_HasLineNumber", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		var tw3 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"valid": tw.HasLineNumber(), "invalid": tw2.HasLineNumber(), "nil": tw3.HasLineNumber()}

		// Assert
		expected := args.Map{"valid": true, "invalid": false, "nil": false}
		expected.ShouldBeEqual(t, 0, "HasLineNumber returns correct value -- valid, invalid, nil", actual)
	})
}

func Test_S01_TWLN_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsInvalidLineNumber", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		var tw3 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"valid": tw.IsInvalidLineNumber(), "invalid": tw2.IsInvalidLineNumber(), "nil": tw3.IsInvalidLineNumber()}

		// Assert
		expected := args.Map{"valid": false, "invalid": true, "nil": true}
		expected.ShouldBeEqual(t, 0, "IsInvalidLineNumber returns correct value -- valid, invalid, nil", actual)
	})
}

func Test_S01_TWLN_Length(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_Length", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		var tw2 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"len": tw.Length(), "nilLen": tw2.Length()}

		// Assert
		expected := args.Map{"len": 5, "nilLen": 0}
		expected.ShouldBeEqual(t, 0, "Length returns correct value -- normal and nil", actual)
	})
}

func Test_S01_TWLN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsEmpty", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		tw3 := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		var tw4 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"validText": tw.IsEmpty(),
			"invalidLn": tw2.IsEmpty(),
			"emptyText": tw3.IsEmpty(),
			"nil":       tw4.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"validText": false,
			"invalidLn": true,
			"emptyText": true,
			"nil":       true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns correct value -- various cases", actual)
	})
}

func Test_S01_TWLN_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsEmptyText", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		var tw3 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"hasText": tw.IsEmptyText(), "empty": tw2.IsEmptyText(), "nil": tw3.IsEmptyText()}

		// Assert
		expected := args.Map{"hasText": false, "empty": true, "nil": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyText returns correct value -- text, empty, nil", actual)
	})
}

func Test_S01_TWLN_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsEmptyTextLineBoth", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}

		// Act
		actual := args.Map{"both": tw.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"both": false}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth returns correct value -- delegates to IsEmpty", actual)
	})
}
