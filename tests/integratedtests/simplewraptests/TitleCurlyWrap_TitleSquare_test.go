package simplewraptests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

// ── TitleCurlyWrap ──

func Test_TitleCurlyWrap(t *testing.T) {
	// Arrange
	result := simplewrap.TitleCurlyWrap("title", "value")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyWrap returns correct value -- with args", actual)
}

// ── TitleSquare ──

func Test_TitleSquare(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquare("title", "value")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquare returns correct value -- with args", actual)
}

// ── TitleSquareMeta ──

func Test_TitleSquareMeta_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquareMeta("title", "value", "meta")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMeta returns correct value -- with args", actual)
}

// ── TitleSquareMetaUsingFmt ──

type cov6Stringer struct{ val string }

func (s cov6Stringer) String() string { return s.val }

func Test_TitleSquareMetaUsingFmt_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquareMetaUsingFmt(
		cov6Stringer{"t"}, cov6Stringer{"v"}, cov6Stringer{"m"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMetaUsingFmt returns correct value -- with args", actual)
}

// ── TitleSquareCsvMeta ──

func Test_TitleSquareCsvMeta_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.TitleSquareCsvMeta("title", "value", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareCsvMeta returns correct value -- with args", actual)
}

// ── ToJsonName ──

func Test_ToJsonName(t *testing.T) {
	// Arrange
	result := simplewrap.ToJsonName("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName returns correct value -- with args", actual)
}

// ── MsgWrapNumber ──

func Test_MsgWrapNumber(t *testing.T) {
	// Arrange
	result := simplewrap.MsgWrapNumber("count", 42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber returns correct value -- with args", actual)
}

// ── With / WithPtr ──

func Test_With(t *testing.T) {
	// Arrange
	result := simplewrap.With("[", "hello", "]")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "With returns non-empty -- with args", actual)
}

func Test_WithPtr(t *testing.T) {
	// Arrange
	s, e, v := "[", "]", "hello"
	result := simplewrap.WithPtr(&s, &v, &e)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "[hello]"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns non-empty -- with args", actual)
}

func Test_WithPtr_Nils(t *testing.T) {
	// Arrange
	v := "hello"
	result := simplewrap.WithPtr(nil, &v, nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns nil -- nils", actual)
}

func Test_WithPtr_NilSource(t *testing.T) {
	// Arrange
	s, e := "[", "]"
	result := simplewrap.WithPtr(&s, nil, &e)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "[]"}
	expected.ShouldBeEqual(t, 0, "WithPtr returns nil -- nil source", actual)
}

// ── WithStartEnd / WithStartEndPtr ──

func Test_WithStartEnd(t *testing.T) {
	// Arrange
	result := simplewrap.WithStartEnd("'", "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WithStartEnd returns non-empty -- with args", actual)
}

func Test_WithStartEndPtr(t *testing.T) {
	// Arrange
	w, v := "'", "hello"
	result := simplewrap.WithStartEndPtr(&w, &v)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "'hello'"}
	expected.ShouldBeEqual(t, 0, "WithStartEndPtr returns non-empty -- with args", actual)
}

// ── WithDoubleQuoteAny ──

func Test_WithDoubleQuoteAny(t *testing.T) {
	// Arrange
	result := simplewrap.WithDoubleQuoteAny(42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny returns non-empty -- with args", actual)
}

// ── WithSingleQuote ──

func Test_WithSingleQuote(t *testing.T) {
	// Arrange
	result := simplewrap.WithSingleQuote("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote returns non-empty -- with args", actual)
}

// ── WithDoubleQuote ──

func Test_WithDoubleQuote(t *testing.T) {
	// Arrange
	result := simplewrap.WithDoubleQuote("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote returns non-empty -- with args", actual)
}

// ── WithBrackets ──

func Test_WithBrackets(t *testing.T) {
	// Arrange
	result := simplewrap.WithBrackets("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBrackets returns non-empty -- with args", actual)
}

// ── WithCurly ──

func Test_WithCurly(t *testing.T) {
	// Arrange
	result := simplewrap.WithCurly("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithCurly returns non-empty -- with args", actual)
}

// ── WithParenthesis ──

func Test_WithParenthesis(t *testing.T) {
	// Arrange
	result := simplewrap.WithParenthesis("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesis returns non-empty -- with args", actual)
}

// ── CurlyWrap ──

func Test_CurlyWrap(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrap("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrap returns correct value -- with args", actual)
}

// ── CurlyWrapIf ──

func Test_CurlyWrapIf(t *testing.T) {
	// Arrange
	result := simplewrap.CurlyWrapIf(true, "hello")
	noWrap := simplewrap.CurlyWrapIf(false, "hello")

	// Act
	actual := args.Map{
		"wrapped": result != "",
		"noWrap": noWrap != "",
	}

	// Assert
	expected := args.Map{
		"wrapped": true,
		"noWrap": true,
	}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf returns correct value -- with args", actual)
}

// ── ParenthesisWrap ──

func Test_ParenthesisWrap(t *testing.T) {
	// Arrange
	result := simplewrap.ParenthesisWrap("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap returns correct value -- with args", actual)
}

// ── ParenthesisWrapIf ──

func Test_ParenthesisWrapIf(t *testing.T) {
	// Arrange
	result := simplewrap.ParenthesisWrapIf(true, "hello")
	noWrap := simplewrap.ParenthesisWrapIf(false, "hello")

	// Act
	actual := args.Map{
		"wrapped": result != "",
		"noWrap": noWrap != "",
	}

	// Assert
	expected := args.Map{
		"wrapped": true,
		"noWrap": true,
	}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrapIf returns correct value -- with args", actual)
}

// ── SquareWrap ──

func Test_SquareWrap(t *testing.T) {
	// Arrange
	result := simplewrap.SquareWrap("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SquareWrap returns correct value -- with args", actual)
}

// ── SquareWrapIf ──

func Test_SquareWrapIf(t *testing.T) {
	// Arrange
	result := simplewrap.SquareWrapIf(true, "hello")
	noWrap := simplewrap.SquareWrapIf(false, "hello")

	// Act
	actual := args.Map{
		"wrapped": result != "",
		"noWrap": noWrap != "",
	}

	// Assert
	expected := args.Map{
		"wrapped": true,
		"noWrap": true,
	}
	expected.ShouldBeEqual(t, 0, "SquareWrapIf returns correct value -- with args", actual)
}

// ── TitleCurlyMeta ──

func Test_TitleCurlyMeta_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.TitleCurlyMeta("title", "value", "meta")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyMeta returns correct value -- with args", actual)
}

// ── TitleQuotationMeta ──

func Test_TitleQuotationMeta_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleQuotationMeta returns correct value -- with args", actual)
}

// ── DoubleQuoteWrapElements — skip on existence ──

func Test_DoubleQuoteWrapElements_SkipOnExistence_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(true, `"hello"`, "world")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns correct value -- skip on existence", actual)
}

func Test_DoubleQuoteWrapElements_Nil_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns nil -- nil", actual)
}

func Test_DoubleQuoteWrapElements_Empty(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements returns empty -- empty", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes ──

func Test_DoubleQuoteWrapElementsWithIndexes_Nil_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns nil -- nil", actual)
}

func Test_DoubleQuoteWrapElementsWithIndexes_Empty_FromTitleCurlyWrapTitleS(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns empty -- empty", actual)
}

func Test_DoubleQuoteWrapElementsWithIndexes_Items(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")

	// Act
	actual := args.Map{
		"len": len(result),
		"notEmpty": result[0] != "",
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns non-empty -- items", actual)
}

// ── ConditionalWrapWith — missing left/right ──

func Test_ConditionalWrapWith_MissingLeft(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "x]", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- missing left", actual)
}

func Test_ConditionalWrapWith_MissingRight(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "[x", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "[x]"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- missing right", actual)
}

func Test_ConditionalWrapWith_SingleCharPresent(t *testing.T) {
	// Arrange
	result := simplewrap.ConditionalWrapWith('[', "[", ']')

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": fmt.Sprintf("[%c", ']')}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith returns non-empty -- single char present", actual)
}

// ── MsgCsvItems with items ──

func Test_MsgCsvItems_WithItems(t *testing.T) {
	// Arrange
	result := simplewrap.MsgCsvItems("msg", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems returns non-empty -- with items", actual)
}
