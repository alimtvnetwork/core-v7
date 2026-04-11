package convertinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

// ── StringsTo.WithSpaces ──

func Test_StringsTo_WithSpaces_FromStringsToWithSpaces(t *testing.T) {
	// Act
	result := convertinternal.StringsTo.WithSpaces(2, "a", "b")
	emptyResult := convertinternal.StringsTo.WithSpaces(2)

	// Assert
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo returns non-empty -- WithSpaces", actual)
}

// ── Util.String methods ──

func Test_StringUtil_IndexToPosition(t *testing.T) {
	// Act
	actual := args.Map{
		"first":  convertinternal.Util.String.IndexToPosition(0),
		"second": convertinternal.Util.String.IndexToPosition(1),
		"third":  convertinternal.Util.String.IndexToPosition(2),
		"fourth": convertinternal.Util.String.IndexToPosition(3),
	}

	// Assert
	expected := args.Map{
		"first": "1st",
		"second": "2nd",
		"third": "3rd",
		"fourth": "4th",
	}
	expected.ShouldBeEqual(t, 0, "IndexToPosition returns correct value -- with args", actual)
}

func Test_StringUtil_PascalCase(t *testing.T) {
	// Act
	actual := args.Map{
		"simple":     convertinternal.Util.String.PascalCase("hello"),
		"underscore": convertinternal.Util.String.PascalCase("hello_world"),
		"empty":      convertinternal.Util.String.PascalCase(""),
		"single":     convertinternal.Util.String.PascalCase("a"),
	}

	// Assert
	expected := args.Map{
		"simple": "Hello",
		"underscore": "HelloWorld",
		"empty": "",
		"single": "A",
	}
	expected.ShouldBeEqual(t, 0, "PascalCase returns correct value -- with args", actual)
}

func Test_StringUtil_CamelCase(t *testing.T) {
	// Act
	actual := args.Map{
		"simple":     convertinternal.Util.String.CamelCase("Hello"),
		"underscore": convertinternal.Util.String.CamelCase("hello_world"),
		"empty":      convertinternal.Util.String.CamelCase(""),
		"single":     convertinternal.Util.String.CamelCase("A"),
	}

	// Assert
	expected := args.Map{
		"simple": "hello",
		"underscore": "helloWorld",
		"empty": "",
		"single": "a",
	}
	expected.ShouldBeEqual(t, 0, "CamelCase returns correct value -- with args", actual)
}

func Test_StringUtil_PrependWithSpaces(t *testing.T) {
	// Act
	result := convertinternal.Util.String.PrependWithSpaces(
		", ",
		2, []string{"existing"},
		4, "prepend",
	)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependWithSpaces returns non-empty -- with args", actual)
}

func Test_StringUtil_PrependWithSpacesDefault(t *testing.T) {
	// Act
	result := convertinternal.Util.String.PrependWithSpacesDefault(
		2, []string{"existing"},
		4, "prepend",
	)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependWithSpacesDefault returns non-empty -- with args", actual)
}

// ── Util.Strings.PrependWithSpaces — no-space branches ──

func Test_StringsUtil_PrependWithSpaces_NoSpaces(t *testing.T) {
	// Act
	result := convertinternal.Util.Strings.PrependWithSpaces(
		0, []string{"existing"},
		0, "prepend",
	)

	// Assert
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "PrependWithSpaces returns empty -- no spaces", actual)
}

// ── AnyTo.String — various types ──

func Test_AnyTo_String_Bool(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  convertinternal.AnyTo.String(true),
		"false": convertinternal.AnyTo.String(false),
	}

	// Assert
	expected := args.Map{
		"true": "true",
		"false": "false",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- String bool", actual)
}

func Test_AnyTo_String_Int(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.String(42)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- String int", actual)
}

func Test_AnyTo_String_Fallback(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.String(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- String fallback", actual)
}

// ── AnyTo.Strings — map[string]string ──

func Test_AnyTo_Strings_MapStringString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(map[string]string{"a": "1"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings map[string]string", actual)
}

// ── AnyTo.Strings — []int / []bool ──

func Test_AnyTo_Strings_IntSlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings int slice", actual)
}

func Test_AnyTo_Strings_BoolSlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings([]bool{true, false})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings bool slice", actual)
}

// ── AnyTo.SmartString — Namer (not easily testable without namer) ──

func Test_AnyTo_ValueString_FromStringsToWithSpaces(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.ValueString(nil)
	nonNil := convertinternal.AnyTo.ValueString(42)

	// Act
	actual := args.Map{
		"nil": result,
		"nonNilNotEmpty": nonNil != "",
	}

	// Assert
	expected := args.Map{
		"nil": "",
		"nonNilNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns non-empty -- ValueString", actual)
}

// ── AnyTo.SmartJson — int/uint types ──

func Test_AnyTo_SmartJson_Int(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartJson(42)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "42"}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartJson int", actual)
}

func Test_AnyTo_SmartJson_Nil(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartJson(nil)

	// Act
	actual := args.Map{"empty": result}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo returns nil -- SmartJson nil", actual)
}

func Test_AnyTo_SmartJson_Default(t *testing.T) {
	// Arrange
	type s struct{ A int }
	result := convertinternal.AnyTo.SmartJson(s{A: 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartJson default fallback", actual)
}
