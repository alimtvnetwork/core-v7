package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

// ── AnyTo.PrettyStringDefaultMust ──

func Test_AnyTo_PrettyStringDefaultMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefaultMust(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- PrettyStringDefaultMust", actual)
}

func Test_AnyTo_PrettyStringDefaultMust_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "AnyTo panics -- PrettyStringDefaultMust panic", actual)
	}()
	// channels can't be marshalled
	ch := make(chan int)
	jsoninternal.Pretty.AnyTo.PrettyStringDefaultMust(ch)
}

// ── Bytes.PrefixMust / DefaultMust ──

func Test_Bytes_PrefixMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.Bytes.PrefixMust("", []byte(`{"a":1}`))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- PrefixMust", actual)
}

func Test_Bytes_PrefixMust_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Bytes panics -- PrefixMust panic", actual)
	}()
	jsoninternal.Pretty.Bytes.PrefixMust("", []byte(`invalid json`))
}

func Test_Bytes_DefaultMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.Bytes.DefaultMust([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- DefaultMust", actual)
}

func Test_Bytes_DefaultMust_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "Bytes panics -- DefaultMust panic", actual)
	}()
	jsoninternal.Pretty.Bytes.DefaultMust([]byte(`invalid`))
}

// ── String converter ──

func Test_StringJson_Default(t *testing.T) {
	// Arrange
	result, err := jsoninternal.String.Default(map[string]int{"a": 1})

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
	expected.ShouldBeEqual(t, 0, "StringJson returns correct value -- Default", actual)
}

func Test_StringJson_SafeDefault(t *testing.T) {
	// Arrange
	result := jsoninternal.String.SafeDefault(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringJson returns correct value -- SafeDefault", actual)
}

func Test_StringJson_Pretty(t *testing.T) {
	// Arrange
	result, err := jsoninternal.String.Pretty(map[string]int{"a": 1})

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
	expected.ShouldBeEqual(t, 0, "StringJson returns correct value -- Pretty", actual)
}

func Test_StringJson_StringValue(t *testing.T) {
	// Arrange
	result := jsoninternal.String.StringValue("hello")

	// Act
	actual := args.Map{"result": string(result)}

	// Assert
	expected := args.Map{"result": `"hello"`}
	expected.ShouldBeEqual(t, 0, "StringJson returns correct value -- StringValue", actual)
}

// ── String to Pretty converter ──

func Test_StringToPretty_Safe(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.Safe("", `{"a":1}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty returns correct value -- Safe", actual)
}

func Test_StringToPretty_SafeDefault(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.SafeDefault(`{"a":1}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty returns correct value -- SafeDefault", actual)
}

func Test_StringToPretty_Indent(t *testing.T) {
	// Arrange
	result, err := jsoninternal.Pretty.String.Indent("", "  ", `{"a":1}`)

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
	expected.ShouldBeEqual(t, 0, "StringToPretty returns correct value -- Indent", actual)
}

func Test_StringToPretty_PrefixMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.PrefixMust("", `{"a":1}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty returns correct value -- PrefixMust", actual)
}

func Test_StringToPretty_PrefixMust_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "StringToPretty panics -- PrefixMust panic", actual)
	}()
	jsoninternal.Pretty.String.PrefixMust("", `invalid`)
}

func Test_StringToPretty_DefaultMust(t *testing.T) {
	// Arrange
	result := jsoninternal.Pretty.String.DefaultMust(`{"a":1}`)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringToPretty returns correct value -- DefaultMust", actual)
}

func Test_StringToPretty_DefaultMust_Panic(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "StringToPretty panics -- DefaultMust panic", actual)
	}()
	jsoninternal.Pretty.String.DefaultMust(`invalid`)
}
