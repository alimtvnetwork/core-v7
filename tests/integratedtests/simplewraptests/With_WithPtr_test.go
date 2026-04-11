package simplewraptests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── Basic wrap functions ──

func Test_With_FromWithWithPtr(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.With("[", "hello", "]")}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "With returns non-empty -- with args", actual)
}

func Test_WithPtr_AllPresent(t *testing.T) {
	// Arrange
	s, src, e := "[", "hello", "]"
	result := simplewrap.WithPtr(&s, &src, &e)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns non-empty -- all present", actual)
}

func Test_WithStartEnd_FromWithWithPtr(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.WithStartEnd("|", "hello")}

	// Assert
	expected := args.Map{"result": "|hello|"}
	expected.ShouldBeEqual(t, 0, "WithStartEnd returns non-empty -- with args", actual)
}

func Test_WithStartEndPtr_FromWithWithPtr(t *testing.T) {
	// Arrange
	w, src := "|", "hello"
	result := simplewrap.WithStartEndPtr(&w, &src)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "|hello|"}
	expected.ShouldBeEqual(t, 0, "WithStartEndPtr returns non-empty -- with args", actual)
}

func Test_WithDoubleQuote_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.WithDoubleQuote("hello")

	// Act
	actual := args.Map{"contains": strings.Contains(result, "hello")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote returns non-empty -- with args", actual)
}

func Test_WithSingleQuote_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.WithSingleQuote("hello")

	// Act
	actual := args.Map{"contains": strings.Contains(result, "hello")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote returns non-empty -- with args", actual)
}

func Test_WithDoubleQuoteAny_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.WithDoubleQuoteAny(42)

	// Act
	actual := args.Map{"contains": strings.Contains(result, "42")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny returns non-empty -- with args", actual)
}

func Test_ToJsonName_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.ToJsonName("test")

	// Act
	actual := args.Map{"contains": strings.Contains(result, "test")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName returns correct value -- with args", actual)
}

// ── Bracket/Curly/Parenthesis wraps ──

func Test_WithBrackets_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.WithBrackets("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithBrackets returns non-empty -- with args", actual)
}

func Test_WithCurly_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.WithCurly("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "WithCurly returns non-empty -- with args", actual)
}

func Test_WithParenthesis_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.WithParenthesis("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "WithParenthesis returns non-empty -- with args", actual)
}

func Test_CurlyWrap_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrap("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "{hello}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrap returns correct value -- with args", actual)
}

func Test_SquareWrap_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.SquareWrap("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "SquareWrap returns correct value -- with args", actual)
}

func Test_ParenthesisWrap_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.ParenthesisWrap("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "(hello)"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap returns correct value -- with args", actual)
}

// ── If variants ──

func Test_CurlyWrapIf_True(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.CurlyWrapIf(true, "x")}

	// Assert
	expected := args.Map{"result": "{x}"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf returns correct value -- true", actual)
}

func Test_CurlyWrapIf_False(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.CurlyWrapIf(false, "x")}

	// Assert
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf returns correct value -- false", actual)
}

func Test_SquareWrapIf_True(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.SquareWrapIf(true, "x")}

	// Assert
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "SquareWrapIf returns correct value -- true", actual)
}

func Test_SquareWrapIf_False(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.SquareWrapIf(false, "x")}

	// Assert
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "SquareWrapIf returns correct value -- false", actual)
}

func Test_ParenthesisWrapIf_True(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ParenthesisWrapIf(true, "x")}

	// Assert
	expected := args.Map{"result": "(x)"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrapIf returns correct value -- true", actual)
}

func Test_ParenthesisWrapIf_False(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ParenthesisWrapIf(false, "x")}

	// Assert
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrapIf returns correct value -- false", actual)
}

// ── Title wraps ──

func Test_TitleCurlyWrap_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleCurlyWrap("title", "value")

	// Act
	actual := args.Map{
		"containsTitle": strings.Contains(result, "title"),
		"containsVal": strings.Contains(result, "value"),
	}

	// Assert
	expected := args.Map{
		"containsTitle": true,
		"containsVal": true,
	}
	expected.ShouldBeEqual(t, 0, "TitleCurlyWrap returns correct value -- with args", actual)
}

func Test_TitleSquare_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquare("title", "value")

	// Act
	actual := args.Map{
		"containsTitle": strings.Contains(result, "title"),
		"containsVal": strings.Contains(result, "value"),
	}

	// Assert
	expected := args.Map{
		"containsTitle": true,
		"containsVal": true,
	}
	expected.ShouldBeEqual(t, 0, "TitleSquare returns correct value -- with args", actual)
}

func Test_TitleCurlyMeta_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleCurlyMeta("title", "value", "meta")

	// Act
	actual := args.Map{"containsMeta": strings.Contains(result, "meta")}

	// Assert
	expected := args.Map{"containsMeta": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyMeta returns correct value -- with args", actual)
}

func Test_TitleSquareMeta_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquareMeta("title", "value", "meta")

	// Act
	actual := args.Map{"containsMeta": strings.Contains(result, "meta")}

	// Assert
	expected := args.Map{"containsMeta": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMeta returns correct value -- with args", actual)
}

func Test_TitleSquareCsvMeta_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquareCsvMeta("title", "value", "m1", "m2")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareCsvMeta returns correct value -- with args", actual)
}

type cov4Stringer struct{ val string }
func (s cov4Stringer) String() string { return s.val }

func Test_TitleSquareMetaUsingFmt_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquareMetaUsingFmt(
		cov4Stringer{"title"},
		cov4Stringer{"value"},
		cov4Stringer{"meta"},
	)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMetaUsingFmt returns correct value -- with args", actual)
}

func Test_TitleQuotationMeta_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")

	// Act
	actual := args.Map{"containsQuote": strings.Contains(result, `"`)}

	// Assert
	expected := args.Map{"containsQuote": true}
	expected.ShouldBeEqual(t, 0, "TitleQuotationMeta returns correct value -- with args", actual)
}

// ── MsgWrapNumber ──

func Test_MsgWrapNumber_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapNumber("count", 42)

	// Act
	actual := args.Map{"contains": strings.Contains(result, "42")}

	// Assert
	expected := args.Map{"contains": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber returns correct value -- with args", actual)
}

// ── MsgCsvItems ──

func Test_MsgCsvItems_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.MsgCsvItems("msg", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems returns correct value -- with args", actual)
}

// ── DoubleQuoteWrapElements ──

func Test_DoubleQuoteWrapElements_Normal(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")

	// Act
	actual := args.Map{"len": len(result), "firstQuoted": strings.HasPrefix(result[0], fmt.Sprintf("%c", '"'))}

	// Assert
	expected := args.Map{
		"len": 2,
		"firstQuoted": true,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns correct value -- normal", actual)
}

func Test_DoubleQuoteWrapElements_Empty_FromWithWithPtr(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns empty -- empty", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes ──

func Test_DoubleQuoteWrapElementsWithIndexes_Normal(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"containsIdx": strings.Contains(result[0], "[0]"),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"containsIdx": true,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns non-empty -- normal", actual)
}
