package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── StringTo additional coverage ──

func Test_StringTo_Integer(t *testing.T) {
	// Arrange
	result, err := converters.StringTo.Integer("42")

	// Act
	actual := args.Map{
		"value": result,
		"hasError": err != nil,
	}

	// Assert
	expected := args.Map{
		"value": 42,
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Integer returns correct value -- with args", actual)
}

func Test_StringTo_Integer_Invalid_FromStringToInteger(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Integer("abc")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Integer returns error -- invalid", actual)
}

func Test_StringTo_Float64(t *testing.T) {
	// Arrange
	result, err := converters.StringTo.Float64("3.14")

	// Act
	actual := args.Map{
		"gt3": result > 3.0,
		"hasError": err != nil,
	}

	// Assert
	expected := args.Map{
		"gt3": true,
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64 returns correct value -- with args", actual)
}

func Test_StringTo_Float64_Invalid_FromStringToInteger(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Float64("abc")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64 returns error -- invalid", actual)
}

func Test_StringTo_Byte(t *testing.T) {
	// Arrange
	result, err := converters.StringTo.Byte("42")

	// Act
	actual := args.Map{
		"value": int(result),
		"hasError": err != nil,
	}

	// Assert
	expected := args.Map{
		"value": 42,
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- with args", actual)
}

func Test_StringTo_Byte_Invalid(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Byte("abc")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns error -- invalid", actual)
}

func Test_StringTo_Byte_Overflow(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Byte("300")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- overflow", actual)
}

// ── AnyTo additional coverage ──

func Test_AnyTo_SmartStrings(t *testing.T) {
	// Arrange
	result := converters.AnyTo.SmartStringsOf("a", 42, true)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartStrings returns correct value -- with args", actual)
}

func Test_AnyTo_SmartStrings_Empty(t *testing.T) {
	// Arrange
	result := converters.AnyTo.SmartStringsOf()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartStrings returns empty -- empty", actual)
}

func Test_AnyTo_ToStringsUsingProcessor_WithBreak(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return "x", true, index >= 0 // break on first
	}, []any{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor returns non-empty -- with break", actual)
}

func Test_AnyTo_ToStringsUsingProcessor_Skip(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return "", false, false // skip all
	}, []any{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor returns correct value -- skip all", actual)
}

func Test_AnyTo_ToNonNullItems_WithValues(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToNonNullItems(true, []any{nil, 42, nil, "hello"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ToNonNullItems returns non-empty -- with values", actual)
}

func Test_AnyTo_ToStringsUsingSimpleProcessor_WithValues(t *testing.T) {
	// Arrange
	result := converters.AnyTo.ToStringsUsingSimpleProcessor(true, func(index int, in any) string {
		return "item"
	}, []any{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingSimpleProcessor returns non-empty -- with values", actual)
}

// ── StringsTo additional ──

func Test_StringsTo_Csv_WithTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Csv(true, "  a  ", "  b  ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsTo.Csv returns non-empty -- with trim", actual)
}

func Test_StringsTo_HashmapOptions_NoTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapOptions(false, "=", "a=1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapOptions returns empty -- no trim", actual)
}

func Test_StringsTo_HashmapUsingFuncOptions_NoTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapUsingFuncOptions(false, func(line string) (string, string) {
		return "k", "v"
	}, "line1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapUsingFuncOptions returns empty -- no trim", actual)
}

// ── StringsToMapConverter additional ──

func Test_StringsToMapConverter_Length(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"length":    mc.Length(),
		"isEmpty":   mc.IsEmpty(),
		"hasAny":    mc.HasAnyItem(),
		"lastIndex": mc.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"length":    3,
		"isEmpty":   false,
		"hasAny":    true,
		"lastIndex": 2,
	}
	expected.ShouldBeEqual(t, 0, "StringsToMapConverter returns correct value -- collection methods", actual)
}

func Test_StringsToMapConverter_Empty(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{})

	// Act
	actual := args.Map{
		"length":  mc.Length(),
		"isEmpty": mc.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"length":  0,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsToMapConverter returns empty -- empty", actual)
}

// ── BytesTo ──

func Test_BytesTo_String(t *testing.T) {
	// Arrange
	result := converters.BytesTo.String([]byte("hello"))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesTo.String returns correct value -- with args", actual)
}

func Test_BytesTo_String_Empty_FromStringToInteger(t *testing.T) {
	// Arrange
	result := converters.BytesTo.String(nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "BytesTo.String returns nil -- nil", actual)
}

func Test_UnsafeBytesTo_String(t *testing.T) {
	// Arrange
	result := converters.UnsafeBytesToString([]byte("hello"))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesTo.String returns correct value -- with args", actual)
}

func Test_UnsafeBytesTo_String_Empty(t *testing.T) {
	// Arrange
	result := converters.UnsafeBytesToString(nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesTo.String returns nil -- nil", actual)
}

// ── MapStringAnyUsingFunc trim ──

func Test_StringsTo_MapStringAnyUsingFunc_Trim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.MapStringAnyUsingFunc(true, func(line string) (string, any) {
		return "  k  ", "v"
	}, "line1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringAnyUsingFunc returns correct value -- trim", actual)
}

func Test_StringsTo_MapStringIntegerUsingFunc_NoTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.MapStringIntegerUsingFunc(false, func(line string) (string, int) {
		return "k", 1
	}, "line1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringIntegerUsingFunc returns empty -- no trim", actual)
}
