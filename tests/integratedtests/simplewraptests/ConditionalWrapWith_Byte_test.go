package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── ConditionalWrapWith (startChar byte, input string, endChar byte) ──

func Test_ConditionalWrapWith_Wrapped(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('[', "x", ']')}

	// Assert
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith wraps -- not already wrapped", actual)
}

func Test_ConditionalWrapWith_AlreadyWrapped(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('[', "[x]", ']')}

	// Assert
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith no-op -- already wrapped", actual)
}

func Test_ConditionalWrapWith_Empty_FromConditionalWrapWithB(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith empty -- just brackets", actual)
}

// ── MsgWrapMsg ──

func Test_MsgWrapMsg(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapMsg("hello", "world")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg -- not empty", actual)
}

// ── CurlyWrapOption (isSkipIfExists bool, source any) ──

func Test_CurlyWrapOption_NonEmpty(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrapOption(false, "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption non-empty -- wrapped", actual)
}

func Test_CurlyWrapOption_SkipIfExists(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrapOption(true, "{hello}")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption skip if exists -- no double wrap", actual)
}

func Test_CurlyWrapOption_Empty(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrapOption(false, "")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption empty -- just curlies", actual)
}

// ── WithBracketsQuotation ──

func Test_WithBracketsQuotation_FromConditionalWrapWithB(t *testing.T) {
	// Arrange
	result := simplewrap.WithBracketsQuotation("x")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBracketsQuotation -- wrapped", actual)
}

// ── WithCurlyQuotation ──

func Test_WithCurlyQuotation_FromConditionalWrapWithB(t *testing.T) {
	// Arrange
	result := simplewrap.WithCurlyQuotation("x")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithCurlyQuotation -- wrapped", actual)
}

// ── WithParenthesisQuotation ──

func Test_WithParenthesisQuotation_FromConditionalWrapWithB(t *testing.T) {
	// Arrange
	result := simplewrap.WithParenthesisQuotation("x")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesisQuotation -- wrapped", actual)
}

// ── MsgCsvItems ──

func Test_MsgCsvItems_Empty(t *testing.T) {
	// Arrange
	result := simplewrap.MsgCsvItems("msg")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems no items -- msg only", actual)
}
