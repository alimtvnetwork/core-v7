package simplewraptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_WithStartEnd_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithStartEnd("[", "hello")

	// Assert
	actual := args.Map{"result": strings.Contains(result, "[")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain bracket wrapper", actual)
}

func Test_WithBracketsQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithBracketsQuotation("hello")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_WithCurlyQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithCurlyQuotation("hello")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_WithParenthesisQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithParenthesisQuotation("hello")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_TitleSquareCsvMeta_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.TitleSquareCsvMeta("title", "a", "b")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_TitleQuotationMeta_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_DoubleQuoteWrapElements_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 elements", actual)
}
