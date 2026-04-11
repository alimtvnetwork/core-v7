package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

// ── bytesToPrettyConvert ──

func Test_Pretty_Bytes_Safe(t *testing.T) {
	// Arrange
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.Safe("", input)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Safe returns correct value -- with args", actual)
}

func Test_Pretty_Bytes_SafeDefault(t *testing.T) {
	// Arrange
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.SafeDefault(input)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.SafeDefault returns correct value -- with args", actual)
}

func Test_Pretty_Bytes_Prefix(t *testing.T) {
	// Arrange
	input := []byte(`{"key":"value"}`)
	result, err := jsoninternal.Pretty.Bytes.Prefix("  ", input)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Prefix returns correct value -- with args", actual)
}

func Test_Pretty_Bytes_Indent(t *testing.T) {
	// Arrange
	input := []byte(`{"key":"value"}`)
	result, err := jsoninternal.Pretty.Bytes.Indent("", "\t", input)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.Indent returns correct value -- with args", actual)
}

func Test_Pretty_Bytes_PrefixMust(t *testing.T) {
	// Arrange
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.PrefixMust("", input)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.PrefixMust returns correct value -- with args", actual)
}

func Test_Pretty_Bytes_DefaultMust(t *testing.T) {
	// Arrange
	input := []byte(`{"key":"value"}`)
	result := jsoninternal.Pretty.Bytes.DefaultMust(input)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.Bytes.DefaultMust returns correct value -- with args", actual)
}

// ── stringToPrettyConvert ──

func Test_Pretty_String_Safe(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.Safe("", `{"key":"value"}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Safe returns correct value -- with args", actual)
}

func Test_Pretty_String_SafeDefault(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.SafeDefault(`{"key":"value"}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.SafeDefault returns correct value -- with args", actual)
}

func Test_Pretty_String_Prefix(t *testing.T) {
	// Arrange
	result, err := jsoninternal.Pretty.String.Prefix("", `{"key":"value"}`)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Prefix returns correct value -- with args", actual)
}

func Test_Pretty_String_Indent(t *testing.T) {
	// Arrange
	result, err := jsoninternal.Pretty.String.Indent("", "\t", `{"key":"value"}`)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty.String.Indent returns correct value -- with args", actual)
}

func Test_Pretty_String_PrefixMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.PrefixMust("", `{"key":"value"}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.PrefixMust returns correct value -- with args", actual)
}

func Test_Pretty_String_DefaultMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.DefaultMust(`{"key":"value"}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Pretty.String.DefaultMust returns correct value -- with args", actual)
}

// ── anyToConvert ──

func Test_AnyTo_SafeString_FromPrettyBytes(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.AnyTo.SafeString(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeString returns correct value -- with args", actual)
}

func Test_AnyTo_String_FromPrettyBytes(t *testing.T) {
	// Arrange
	result, err := jsoninternal.Pretty.AnyTo.String(map[string]string{"k": "v"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.String returns correct value -- with args", actual)
}

func Test_AnyTo_PrettyString_FromPrettyBytes(t *testing.T) {
	// Arrange
	result, err := jsoninternal.Pretty.AnyTo.PrettyString("", map[string]string{"k": "v"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyString returns correct value -- with args", actual)
}

func Test_AnyTo_PrettyStringIndent_FromPrettyBytes(t *testing.T) {
	// Arrange
	result, err := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "\t", map[string]string{"k": "v"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringIndent returns correct value -- with args", actual)
}

func Test_AnyTo_SafePrettyString_FromPrettyBytes(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.AnyTo.SafePrettyString("", map[string]string{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafePrettyString returns correct value -- with args", actual)
}

func Test_AnyTo_PrettyStringDefault_FromPrettyBytes(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefault(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringDefault returns correct value -- with args", actual)
}

func Test_AnyTo_PrettyStringDefaultMust_FromPrettyBytes(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefaultMust(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringDefaultMust returns correct value -- with args", actual)
}

// ── stringJsonConverter ──

func Test_StringJson_SafeDefault_FromPrettyBytes(t *testing.T) {
	// Arrange
	result := jsoninternal.String.SafeDefault(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringJson.SafeDefault returns correct value -- with args", actual)
}

func Test_StringJson_Default_FromPrettyBytes(t *testing.T) {
	// Arrange
	result, err := jsoninternal.String.Default(map[string]string{"k": "v"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringJson.Default returns correct value -- with args", actual)
}

func Test_StringJson_Pretty_FromPrettyBytes(t *testing.T) {
	// Arrange
	result, err := jsoninternal.String.Pretty(map[string]string{"k": "v"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringJson.Pretty returns correct value -- with args", actual)
}

func Test_StringJson_StringValue_FromPrettyBytes(t *testing.T) {
	// Arrange
	result := jsoninternal.String.StringValue("test")

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringJson.StringValue returns correct value -- with args", actual)
}
