package simplewraptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── ConditionalWrapWith exhaustive branch coverage ──

func Test_ConditionalWrapWith_BothMissing(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "hello", '}')}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- both missing", actual)
}

func Test_ConditionalWrapWith_LeftMissing(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "hello}", '}')}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- left missing", actual)
}

func Test_ConditionalWrapWith_RightMissing(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{hello", '}')}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- right missing", actual)
}

func Test_ConditionalWrapWith_BothPresent(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{hello}", '}')}
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- both present", actual)
}

func Test_ConditionalWrapWith_SingleCharMatchStart(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{", '}')}
	expected := args.Map{"result": "{}" }
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- single char start", actual)
}

func Test_ConditionalWrapWith_Empty(t *testing.T) {
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "", '}')}
	expected := args.Map{"result": "{}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns empty -- empty", actual)
}

// ── MsgWrapMsg all branches ──

func Test_MsgWrapMsg_BothEmpty(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("", "")}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns empty -- both empty", actual)
}

func Test_MsgWrapMsg_MsgEmpty(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("", "wrapped")}
	expected := args.Map{"result": "wrapped"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns empty -- msg empty", actual)
}

func Test_MsgWrapMsg_WrappedEmpty(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("msg", "")}
	expected := args.Map{"result": "msg"}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns empty -- wrapped empty", actual)
}

func Test_MsgWrapMsg_BothPresent(t *testing.T) {
	result := simplewrap.MsgWrapMsg("msg", "wrapped")
	actual := args.Map{
		"containsMsg": strings.Contains(result, "msg"),
		"containsWrapped": strings.Contains(result, "wrapped"),
	}
	expected := args.Map{
		"containsMsg": true,
		"containsWrapped": true,
	}
	expected.ShouldBeEqual(t, 0, "MsgWrapMsg returns correct value -- both present", actual)
}

// ── DoubleQuoteWrapElements branches ──

func Test_DoubleQuoteWrapElements_NilInput(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)
	actual := args.Map{
		"notNil": result != nil,
		"len": len(result),
	}
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns nil -- nil input", actual)
}

func Test_DoubleQuoteWrapElements_SkipExistence(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElements(true, `"already"`, "naked")
	actual := args.Map{"len": len(result), "alreadyWrapped": result[0] == `"already"`, "nakedWrapped": strings.HasPrefix(result[1], `"`)}
	expected := args.Map{
		"len": 2,
		"alreadyWrapped": true,
		"nakedWrapped": true,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns correct value -- skip existence", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes branches ──

func Test_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)
	actual := args.Map{
		"notNil": result != nil,
		"len": len(result),
	}
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns nil -- nil", actual)
}

func Test_DoubleQuoteWrapElementsWithIndexes_Empty(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns empty -- empty", actual)
}

// ── WithPtr nil combinations ──

func Test_WithPtr_AllNil(t *testing.T) {
	result := simplewrap.WithPtr(nil, nil, nil)
	actual := args.Map{"result": *result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "WithPtr returns nil -- all nil", actual)
}

// ── CurlyWrapOption branches ──

func Test_CurlyWrapOption_SkipAlreadyWrapped(t *testing.T) {
	actual := args.Map{"result": simplewrap.CurlyWrapOption(true, "{test}")}
	expected := args.Map{"result": "{test}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption returns correct value -- skip already wrapped", actual)
}

func Test_CurlyWrapOption_NoSkip(t *testing.T) {
	actual := args.Map{"result": simplewrap.CurlyWrapOption(false, "test")}
	expected := args.Map{"result": "{test}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption returns empty -- no skip", actual)
}

// ── WithBracketsQuotation / WithCurlyQuotation / WithParenthesisQuotation ──

func Test_WithBracketsQuotation(t *testing.T) {
	r := simplewrap.WithBracketsQuotation("test")
	actual := args.Map{"containsBracket": strings.Contains(r, "["), "containsQuote": strings.Contains(r, `"`)}
	expected := args.Map{
		"containsBracket": true,
		"containsQuote": true,
	}
	expected.ShouldBeEqual(t, 0, "WithBracketsQuotation returns non-empty -- with args", actual)
}

func Test_WithCurlyQuotation(t *testing.T) {
	r := simplewrap.WithCurlyQuotation("test")
	actual := args.Map{"containsCurly": strings.Contains(r, "{"), "containsQuote": strings.Contains(r, `"`)}
	expected := args.Map{
		"containsCurly": true,
		"containsQuote": true,
	}
	expected.ShouldBeEqual(t, 0, "WithCurlyQuotation returns non-empty -- with args", actual)
}

func Test_WithParenthesisQuotation(t *testing.T) {
	r := simplewrap.WithParenthesisQuotation("test")
	actual := args.Map{"containsParen": strings.Contains(r, "("), "containsQuote": strings.Contains(r, `"`)}
	expected := args.Map{
		"containsParen": true,
		"containsQuote": true,
	}
	expected.ShouldBeEqual(t, 0, "WithParenthesisQuotation returns non-empty -- with args", actual)
}
