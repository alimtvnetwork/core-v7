package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── WithDoubleQuote ──

func Test_WithDoubleQuote_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithDoubleQuote("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": `"hello"`}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote returns non-empty -- with args", actual)
}

// ── WithSingleQuote ──

func Test_WithSingleQuote_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithSingleQuote("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote returns non-empty -- with args", actual)
}

// ── With ──

func Test_With_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.With("[", "hello", "]")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "With returns non-empty -- with args", actual)
}

// ── WithPtr ──

func Test_WithPtr_AllNonNil(t *testing.T) {
	// Arrange
	s, e, src := "[", "]", "hello"
	result := simplewrap.WithPtr(&s, &src, &e)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns nil -- all non-nil", actual)
}

func Test_WithPtr_NilSource_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	s, e := "[", "]"
	result := simplewrap.WithPtr(&s, nil, &e)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns nil -- nil source", actual)
}

func Test_WithPtr_NilStartEnd(t *testing.T) {
	// Arrange
	src := "hello"
	result := simplewrap.WithPtr(nil, &src, nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns nil -- nil start/end", actual)
}

// ── WithStartEnd ──

func Test_WithStartEnd_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithStartEnd("*", "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "*hello*"}
	expected.ShouldBeEqual(t, 0, "WithStartEnd returns non-empty -- with args", actual)
}

// ── WithBrackets ──

func Test_WithBrackets_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithBrackets("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithBrackets returns non-empty -- with args", actual)
}

// ── WithCurly ──

func Test_WithCurly_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithCurly("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "WithCurly returns non-empty -- with args", actual)
}

// ── WithParenthesis ──

func Test_WithParenthesis_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithParenthesis("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "WithParenthesis returns non-empty -- with args", actual)
}

// ── WithDoubleQuoteAny ──

func Test_WithDoubleQuoteAny_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.WithDoubleQuoteAny(42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny returns non-empty -- with args", actual)
}

// ── CurlyWrap ──

func Test_CurlyWrap_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrap("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrap returns correct value -- with args", actual)
}

// ── SquareWrap ──

func Test_SquareWrap_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.SquareWrap("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "SquareWrap returns correct value -- with args", actual)
}

// ── ParenthesisWrap ──

func Test_ParenthesisWrap_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ParenthesisWrap("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap returns correct value -- with args", actual)
}

// ── ToJsonName ──

func Test_ToJsonName_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ToJsonName("name")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName returns correct value -- with args", actual)
}

// ── MsgWrapMsg ──

func Test_MsgWrapMsg_BothEmpty_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapMsg("", "")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns empty -- both empty", actual)
}

func Test_MsgWrapMsg_EmptyMsg(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapMsg("", "wrapped")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "wrapped"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns empty -- empty msg", actual)
}

func Test_MsgWrapMsg_EmptyWrapped(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapMsg("msg", "")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns empty -- empty wrapped", actual)
}

func Test_MsgWrapMsg_Both(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapMsg("msg", "wrapped")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns correct value -- both", actual)
}

// ── MsgWrapNumber ──

func Test_MsgWrapNumber_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapNumber("count", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber returns correct value -- with args", actual)
}

// ── MsgCsvItems ──

func Test_MsgCsvItems(t *testing.T) {
	// Arrange
	result := simplewrap.MsgCsvItems("msg", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems returns correct value -- with args", actual)
}

// ── ConditionalWrapWith ──

func Test_ConditionalWrapWith_Empty_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns empty -- empty", actual)
}

func Test_ConditionalWrapWith_AlreadyWrapped_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "[hello]", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- already wrapped", actual)
}

func Test_ConditionalWrapWith_MissingRight_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "[hello", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- missing right", actual)
}

func Test_ConditionalWrapWith_MissingLeft_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "hello]", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- missing left", actual)
}

func Test_ConditionalWrapWith_BothMissing_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "hello", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- both missing", actual)
}

func Test_ConditionalWrapWith_SingleCharMatch(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "[", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- single char match", actual)
}

// ── DoubleQuoteWrapElements ──

func Test_DoubleQuoteWrapElements_Nil_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns nil -- nil", actual)
}

func Test_DoubleQuoteWrapElements_Empty_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns empty -- empty", actual)
}

func Test_DoubleQuoteWrapElements_NonEmpty(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": `"a"`,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns empty -- non-empty", actual)
}

func Test_DoubleQuoteWrapElements_SkipOnExistence_FromWithQuoteVariants(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(true, "a", `"b"`)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns correct value -- skip on existence", actual)
}
