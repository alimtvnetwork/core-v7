package simplewraptests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/simplewrap"
)

type testStringer struct{}

func (s testStringer) String() string { return "stringer" }

func Test_TitleCurlyMeta(t *testing.T) {
	// Arrange
	r := simplewrap.TitleCurlyMeta("title", "val", "meta")

	// Act
	actual := args.Map{
		"notEmpty": r != "",
		"containsTitle": strings.Contains(r, "title"),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"containsTitle": true,
	}
	expected.ShouldBeEqual(t, 0, "TitleCurlyMeta returns correct value -- with args", actual)
}

func Test_TitleSquareMeta(t *testing.T) {
	// Arrange
	r := simplewrap.TitleSquareMeta("title", "val", "meta")

	// Act
	actual := args.Map{
		"notEmpty": r != "",
		"containsTitle": strings.Contains(r, "title"),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"containsTitle": true,
	}
	expected.ShouldBeEqual(t, 0, "TitleSquareMeta returns correct value -- with args", actual)
}

func Test_TitleQuotationMeta(t *testing.T) {
	// Arrange
	r := simplewrap.TitleQuotationMeta("title", "val", "meta")

	// Act
	actual := args.Map{
		"notEmpty": r != "",
		"containsTitle": strings.Contains(r, "title"),
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"containsTitle": true,
	}
	expected.ShouldBeEqual(t, 0, "TitleQuotationMeta returns correct value -- with args", actual)
}

func Test_TitleSquareCsvMeta(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.TitleSquareCsvMeta("title", "val", "a", "b") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareCsvMeta returns correct value -- with args", actual)
}

func Test_TitleSquareMetaUsingFmt(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.TitleSquareMetaUsingFmt(testStringer{}, testStringer{}, testStringer{}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquareMetaUsingFmt returns correct value -- with args", actual)
}

func Test_WithBracketsQuotation_FromTitleCurlyMeta(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.WithBracketsQuotation("hello") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBracketsQuotation returns non-empty -- with args", actual)
}

func Test_WithCurlyQuotation_FromTitleCurlyMeta(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.WithCurlyQuotation("hello") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithCurlyQuotation returns non-empty -- with args", actual)
}

func Test_WithParenthesisQuotation_FromTitleCurlyMeta(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.WithParenthesisQuotation("hello") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesisQuotation returns non-empty -- with args", actual)
}

func Test_CurlyWrapOption(t *testing.T) {
	// Act
	actual := args.Map{
		"skipIfExists":    simplewrap.CurlyWrapOption(true, "{hello}"),
		"noSkip":          simplewrap.CurlyWrapOption(false, "hello"),
		"skipNotPresent":  simplewrap.CurlyWrapOption(true, "hello"),
	}

	// Assert
	expected := args.Map{
		"skipIfExists":    "{hello}",
		"noSkip":          "{hello}",
		"skipNotPresent":  "{hello}",
	}
	expected.ShouldBeEqual(t, 0, "CurlyWrapOption returns correct value -- with args", actual)
}

func Test_DoubleQuoteWrapElements_Nil_FromTitleCurlyMeta(t *testing.T) {
	// Arrange
	r := simplewrap.DoubleQuoteWrapElements(false, nil...)

	// Act
	actual := args.Map{"isNotNil": r != nil}

	// Assert
	expected := args.Map{"isNotNil": true}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements_Nil returns nil -- with args", actual)
}

func Test_DoubleQuoteWrapElements_EmptySlice(t *testing.T) {
	// Act
	actual := args.Map{"len": len(simplewrap.DoubleQuoteWrapElements(false))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements_EmptySlice returns empty -- with args", actual)
}

func Test_DoubleQuoteWrapElements_SkipExistence_FromTitleCurlyMeta(t *testing.T) {
	// Act
	actual := args.Map{"len": len(simplewrap.DoubleQuoteWrapElements(true, `"already"`, "naked"))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements_SkipExistence returns correct value -- with args", actual)
}

func Test_DoubleQuoteWrapElementsWithIndexes_FromTitleCurlyMeta(t *testing.T) {
	// Act
	actual := args.Map{
		"nilNotNil":     simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...) != nil,
		"emptyLen":      len(simplewrap.DoubleQuoteWrapElementsWithIndexes()),
		"itemsLen":      len(simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")),
		"containsIndex": strings.Contains(simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")[0], "[0]"),
	}

	// Assert
	expected := args.Map{
		"nilNotNil":     true,
		"emptyLen":      0,
		"itemsLen":      2,
		"containsIndex": true,
	}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes returns non-empty -- with args", actual)
}

func Test_WithDoubleQuote_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.WithDoubleQuote("")}

	// Assert
	expected := args.Map{"result": `""`}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuote_Empty returns empty -- with args", actual)
}

func Test_WithDoubleQuoteAny_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.WithDoubleQuoteAny(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithDoubleQuoteAny_Int returns non-empty -- with args", actual)
}

func Test_WithSingleQuote_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.WithSingleQuote("")}

	// Assert
	expected := args.Map{"result": "''"}
	expected.ShouldBeEqual(t, 0, "WithSingleQuote_Empty returns empty -- with args", actual)
}

func Test_ToJsonName_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.ToJsonName(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToJsonName_Int returns correct value -- with args", actual)
}

func Test_WithCurly_Int(t *testing.T) {
	// Act
	actual := args.Map{"contains42": strings.Contains(simplewrap.WithCurly(42), "42")}

	// Assert
	expected := args.Map{"contains42": true}
	expected.ShouldBeEqual(t, 0, "WithCurly_Int returns non-empty -- with args", actual)
}

func Test_With_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.With("", "", "")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "With_Empty returns empty -- with args", actual)
}

func Test_WithStartEnd_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.WithStartEnd("", "")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "WithStartEnd_Empty returns empty -- with args", actual)
}

func Test_MsgWrapNumber_Int64(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.MsgWrapNumber("total", int64(100)) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgWrapNumber_Int64 returns correct value -- with args", actual)
}

func Test_MsgCsvItems_Empty_FromTitleCurlyMeta(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.MsgCsvItems("msg") != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MsgCsvItems_Empty returns empty -- with args", actual)
}

func Test_ConditionalWrapWith_BothPresent2Char(t *testing.T) {
	// Act
	actual := args.Map{"result": simplewrap.ConditionalWrapWith('{', "{}", '}')}

	// Assert
	expected := args.Map{"result": "{}"}
	expected.ShouldBeEqual(t, 0, "ConditionalWrapWith_BothPresent2Char returns non-empty -- with args", actual)
}

func Test_CurlyWrap_Int(t *testing.T) {
	// Act
	actual := args.Map{"contains42": strings.Contains(simplewrap.CurlyWrap(42), "42")}

	// Assert
	expected := args.Map{"contains42": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrap_Int returns correct value -- with args", actual)
}

func Test_SquareWrap_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.SquareWrap(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SquareWrap_Int returns correct value -- with args", actual)
}

func Test_ParenthesisWrap_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.ParenthesisWrap(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ParenthesisWrap_Int returns correct value -- with args", actual)
}

func Test_TitleCurlyWrap_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.TitleCurlyWrap("t", 42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleCurlyWrap_Int returns correct value -- with args", actual)
}

func Test_TitleSquare_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.TitleSquare("t", 42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TitleSquare_Int returns correct value -- with args", actual)
}

func Test_WithBrackets_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.WithBrackets(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithBrackets_Int returns non-empty -- with args", actual)
}

func Test_WithParenthesis_Int(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.WithParenthesis(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "WithParenthesis_Int returns non-empty -- with args", actual)
}

func Test_CurlyWrapIf_Stringer(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": simplewrap.CurlyWrapIf(true, testStringer{}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf_Stringer returns correct value -- with args", actual)
}

func Test_CurlyWrapIf_FmtStringer(t *testing.T) {
	// Arrange
	var s fmt.Stringer = testStringer{}

	// Act
	actual := args.Map{"notEmpty": simplewrap.CurlyWrapIf(true, s) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf_FmtStringer returns correct value -- with args", actual)
}
