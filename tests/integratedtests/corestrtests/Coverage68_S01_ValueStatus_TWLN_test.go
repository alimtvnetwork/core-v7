package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ValueStatus ──

func Test_S01_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_S01_InvalidValueStatus", func() {
		vs := corestr.InvalidValueStatus("bad")
		actual := args.Map{
			"valid":   vs.ValueValid.IsValid,
			"msg":     vs.ValueValid.Message,
			"index":   vs.Index,
			"valEmpty": vs.ValueValid.Value == "",
		}
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
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"valid": vs.ValueValid.IsValid, "msg": vs.ValueValid.Message}
		expected := args.Map{"valid": false, "msg": ""}
		expected.ShouldBeEqual(t, 0, "InvalidValueStatusNoMessage returns correct value -- no message", actual)
	})
}

func Test_S01_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_S01_ValueStatus_Clone", func() {
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("hello"),
			Index:      5,
		}
		c := vs.Clone()
		actual := args.Map{"val": c.ValueValid.Value, "index": c.Index, "samePtr": vs == c}
		expected := args.Map{"val": "hello", "index": 5, "samePtr": false}
		expected.ShouldBeEqual(t, 0, "ValueStatus.Clone returns correct value -- deep copy", actual)
	})
}

// ── TextWithLineNumber ──

func Test_S01_TWLN_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_HasLineNumber", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		var tw3 *corestr.TextWithLineNumber
		actual := args.Map{"valid": tw.HasLineNumber(), "invalid": tw2.HasLineNumber(), "nil": tw3.HasLineNumber()}
		expected := args.Map{"valid": true, "invalid": false, "nil": false}
		expected.ShouldBeEqual(t, 0, "HasLineNumber returns correct value -- valid, invalid, nil", actual)
	})
}

func Test_S01_TWLN_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsInvalidLineNumber", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		var tw3 *corestr.TextWithLineNumber
		actual := args.Map{"valid": tw.IsInvalidLineNumber(), "invalid": tw2.IsInvalidLineNumber(), "nil": tw3.IsInvalidLineNumber()}
		expected := args.Map{"valid": false, "invalid": true, "nil": true}
		expected.ShouldBeEqual(t, 0, "IsInvalidLineNumber returns correct value -- valid, invalid, nil", actual)
	})
}

func Test_S01_TWLN_Length(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_Length", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		var tw2 *corestr.TextWithLineNumber
		actual := args.Map{"len": tw.Length(), "nilLen": tw2.Length()}
		expected := args.Map{"len": 5, "nilLen": 0}
		expected.ShouldBeEqual(t, 0, "Length returns correct value -- normal and nil", actual)
	})
}

func Test_S01_TWLN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsEmpty", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		tw3 := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		var tw4 *corestr.TextWithLineNumber
		actual := args.Map{
			"validText": tw.IsEmpty(),
			"invalidLn": tw2.IsEmpty(),
			"emptyText": tw3.IsEmpty(),
			"nil":       tw4.IsEmpty(),
		}
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
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		var tw3 *corestr.TextWithLineNumber
		actual := args.Map{"hasText": tw.IsEmptyText(), "empty": tw2.IsEmptyText(), "nil": tw3.IsEmptyText()}
		expected := args.Map{"hasText": false, "empty": true, "nil": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyText returns correct value -- text, empty, nil", actual)
	})
}

func Test_S01_TWLN_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_S01_TWLN_IsEmptyTextLineBoth", func() {
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		actual := args.Map{"both": tw.IsEmptyTextLineBoth()}
		expected := args.Map{"both": false}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth returns correct value -- delegates to IsEmpty", actual)
	})
}
