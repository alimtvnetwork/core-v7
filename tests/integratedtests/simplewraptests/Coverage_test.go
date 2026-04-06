package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── WithDoubleQuote / WithDoubleQuoteAny / WithSingleQuote ──

func Test_WithDoubleQuote_Coverage(t *testing.T) {
	result := simplewrap.WithDoubleQuote("hello")
	actual := args.Map{"result": result != `"hello"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected \"hello\", actual)
}

func Test_WithDoubleQuoteAny_Coverage(t *testing.T) {
	result := simplewrap.WithDoubleQuoteAny(42)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_WithSingleQuote_Coverage(t *testing.T) {
	result := simplewrap.WithSingleQuote("hello")
	actual := args.Map{"result": result != `'hello'`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello'", actual)
}

// ── CurlyWrap / CurlyWrapIf ──

func Test_CurlyWrap_Coverage(t *testing.T) {
	result := simplewrap.CurlyWrap("hello")
	actual := args.Map{"result": result != "{hello}"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected {hello}", actual)
}

func Test_CurlyWrapIf_Coverage(t *testing.T) {
	wrapped := simplewrap.CurlyWrapIf(true, "hello")
	actual := args.Map{"result": wrapped != "{hello}"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected {hello}", actual)

	notWrapped := simplewrap.CurlyWrapIf(false, "hello")
	actual := args.Map{"result": notWrapped != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

// ── SquareWrap / SquareWrapIf ──

func Test_SquareWrap_Coverage(t *testing.T) {
	result := simplewrap.SquareWrap("hello")
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [hello]", actual)
}

func Test_SquareWrapIf_Coverage(t *testing.T) {
	wrapped := simplewrap.SquareWrapIf(true, "hello")
	actual := args.Map{"result": wrapped != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [hello]", actual)

	notWrapped := simplewrap.SquareWrapIf(false, "hello")
	actual := args.Map{"result": notWrapped != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

// ── ParenthesisWrap / ParenthesisWrapIf ──

func Test_ParenthesisWrap_Coverage(t *testing.T) {
	result := simplewrap.ParenthesisWrap("hello")
	actual := args.Map{"result": result != "(hello)"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected (hello)", actual)
}

func Test_ParenthesisWrapIf_Coverage(t *testing.T) {
	wrapped := simplewrap.ParenthesisWrapIf(true, "hello")
	actual := args.Map{"result": wrapped != "(hello)"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected (hello)", actual)

	notWrapped := simplewrap.ParenthesisWrapIf(false, "hello")
	actual := args.Map{"result": notWrapped != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

// ── With / WithPtr / WithStartEnd / WithStartEndPtr ──

func Test_With_Coverage(t *testing.T) {
	result := simplewrap.With("[", "hello", "]")
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [hello]", actual)
}

func Test_WithPtr_Coverage(t *testing.T) {
	start, source, end := "[", "hello", "]"
	result := simplewrap.WithPtr(&start, &source, &end)
	actual := args.Map{"result": *result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [hello]", actual)

	// Nil cases
	resultNil := simplewrap.WithPtr(nil, &source, nil)
	actual := args.Map{"result": *resultNil != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)

	resultNilSrc := simplewrap.WithPtr(&start, nil, &end)
	actual := args.Map{"result": *resultNilSrc != "[]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected []", actual)
}

func Test_WithStartEnd_Coverage(t *testing.T) {
	result := simplewrap.WithStartEnd("'", "hello")
	actual := args.Map{"result": result != "'hello'"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello'", actual)
}

func Test_WithStartEndPtr_Coverage(t *testing.T) {
	wrapper, source := "'", "hello"
	result := simplewrap.WithStartEndPtr(&wrapper, &source)
	actual := args.Map{"result": *result != "'hello'"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello'", actual)
}

// ── WithBrackets / WithCurly / WithParenthesis ──

func Test_WithBrackets_Coverage(t *testing.T) {
	result := simplewrap.WithBrackets("hello")
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [hello]", actual)
}

func Test_WithCurly_Coverage(t *testing.T) {
	result := simplewrap.WithCurly("hello")
	actual := args.Map{"result": result != "{hello}"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected {hello}", actual)
}

func Test_WithParenthesis_Coverage(t *testing.T) {
	result := simplewrap.WithParenthesis("hello")
	actual := args.Map{"result": result != "(hello)"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected (hello)", actual)
}

// ── TitleCurlyWrap / TitleSquare ──

func Test_TitleCurlyWrap_Coverage(t *testing.T) {
	result := simplewrap.TitleCurlyWrap("title", "value")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_TitleSquare_Coverage(t *testing.T) {
	result := simplewrap.TitleSquare("title", "value")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

// ── MsgWrapMsg / MsgWrapNumber / MsgCsvItems ──

func Test_MsgWrapMsg_Coverage(t *testing.T) {
	actual := args.Map{"result": simplewrap.MsgWrapMsg("", "") != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both empty should be empty", actual)
	actual := args.Map{"result": simplewrap.MsgWrapMsg("", "wrapped") != "wrapped"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty msg should return wrapped", actual)
	actual := args.Map{"result": simplewrap.MsgWrapMsg("msg", "") != "msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty wrapped should return msg", actual)
	result := simplewrap.MsgWrapMsg("msg", "wrapped")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both non-empty should not be empty", actual)
}

func Test_MsgWrapNumber_Coverage(t *testing.T) {
	result := simplewrap.MsgWrapNumber("count", 42)
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_MsgCsvItems_Coverage(t *testing.T) {
	result := simplewrap.MsgCsvItems("items", "a", "b", "c")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

// ── ToJsonName ──

func Test_ToJsonName_Coverage(t *testing.T) {
	result := simplewrap.ToJsonName("hello")
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

// ── ConditionalWrapWith ──

func Test_ConditionalWrapWith_Coverage(t *testing.T) {
	// Both present — return as-is
	result := simplewrap.ConditionalWrapWith('[', "[hello]", ']')
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both present: expected [hello]", actual)

	// Empty input — wrap
	result = simplewrap.ConditionalWrapWith('[', "", ']')
	actual := args.Map{"result": result != "[]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty: expected []", actual)

	// Both missing — add both
	result = simplewrap.ConditionalWrapWith('[', "hello", ']')
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both missing: expected [hello]", actual)

	// Right missing
	result = simplewrap.ConditionalWrapWith('[', "[hello", ']')
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right missing: expected [hello]", actual)

	// Left missing
	result = simplewrap.ConditionalWrapWith('[', "hello]", ']')
	actual := args.Map{"result": result != "[hello]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "left missing: expected [hello]", actual)

	// Single char that matches start
	result = simplewrap.ConditionalWrapWith('[', "[", ']')
	actual := args.Map{"result": result != "[]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single char: expected []", actual)
}

// ── DoubleQuoteWrapElements ──

func Test_DoubleQuoteWrapElements_Coverage(t *testing.T) {
	// Normal
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2 items", actual)
	actual := args.Map{"result": result[0] != `"a"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected \"a\", actual)

	// Nil input
	result = simplewrap.DoubleQuoteWrapElements(false, )
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty input should return empty", actual)

	// With skip
	result = simplewrap.DoubleQuoteWrapElements(true, "a", "b")
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2 items", actual)
}

// ── DoubleQuoteWrapElementsWithIndexes ──

func Test_DoubleQuoteWrapElementsWithIndexes_Coverage(t *testing.T) {
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(
		"a", "b")
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2 items", actual)
}
