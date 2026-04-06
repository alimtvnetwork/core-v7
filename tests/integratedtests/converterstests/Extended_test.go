package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestStringTo_IntegerWithDefault verifies integer conversion with default.
func TestStringTo_IntegerWithDefault(t *testing.T) {
	for _, tc := range stringToIntegerWithDefaultCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			val, ok := converters.StringTo.IntegerWithDefault(tc.input, tc.defaultVal)

			// Assert
			actual := args.Map{"result": val != tc.expectedVal}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
			actual := args.Map{"result": ok != tc.expectedOk}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected ok=", actual)
		})
	}
}

// TestStringTo_Integer verifies integer conversion with error.
func TestStringTo_Integer(t *testing.T) {
	// Act
	val, err := converters.StringTo.Integer("42")

	// Assert
	actual := args.Map{"result": err != nil || val != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42, got, err=", actual)

	_, err = converters.StringTo.Integer("abc")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-numeric", actual)
}

// TestStringTo_IntegerDefault verifies default integer conversion.
func TestStringTo_IntegerDefault(t *testing.T) {
	actual := args.Map{"result": converters.StringTo.IntegerDefault("10") != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
	actual := args.Map{"result": converters.StringTo.IntegerDefault("abc") != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for invalid", actual)
}

// TestStringTo_Float64 verifies float64 conversion.
func TestStringTo_Float64(t *testing.T) {
	val, err := converters.StringTo.Float64("3.14")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": val < 3.13 || val > 3.15}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ~3.14", actual)

	_, err = converters.StringTo.Float64("abc")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-numeric", actual)
}

// TestStringTo_Float64Default verifies default float conversion.
func TestStringTo_Float64Default(t *testing.T) {
	val, ok := converters.StringTo.Float64Default("2.5", 0.0)
	actual := args.Map{"result": ok || val != 2.5}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	val, ok = converters.StringTo.Float64Default("abc", 9.9)
	actual := args.Map{"result": ok || val != 9.9}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 9.9 default", actual)
}

// TestStringTo_Float64Conditional verifies deprecated conditional.
func TestStringTo_Float64Conditional(t *testing.T) {
	val, ok := converters.StringTo.Float64Conditional("2.5", 0.0)
	actual := args.Map{"result": ok || val != 2.5}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
}

// TestStringTo_Byte verifies byte conversion.
func TestStringTo_Byte(t *testing.T) {
	for _, tc := range stringToByteCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := converters.StringTo.Byte(tc.input)
			actual := args.Map{"result": tc.expectErr && err == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected error", actual)
			actual := args.Map{"result": tc.expectErr && err != nil}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
			actual := args.Map{"result": tc.expectErr && val != tc.expected}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestStringTo_ByteWithDefault verifies byte with default.
func TestStringTo_ByteWithDefault(t *testing.T) {
	val, ok := converters.StringTo.ByteWithDefault("100", 0)
	actual := args.Map{"result": ok || val != 100}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
	val, ok = converters.StringTo.ByteWithDefault("abc", 55)
	actual := args.Map{"result": ok || val != 55}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 55 default", actual)
}

// TestStringTo_IntegersWithDefaults verifies multi-integer parsing.
func TestStringTo_IntegersWithDefaults(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("1,2,abc", ",", -1)
	actual := args.Map{"result": len(result.Values) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 values", actual)
	actual := args.Map{"result": result.Values[2] != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected default -1 for invalid", actual)
	actual := args.Map{"result": result.CombinedError == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected combined error", actual)
}

// TestStringTo_IntegersWithDefaults_Empty verifies empty input.
func TestStringTo_IntegersWithDefaults_Empty(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("", ",", -1)
	actual := args.Map{"result": len(result.Values) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 values", actual)
}

// TestStringTo_IntegersConditional verifies conditional integer parsing.
func TestStringTo_IntegersConditional(t *testing.T) {
	result := converters.StringTo.IntegersConditional("1,2,3", ",", func(in string) (int, bool, bool) {
		if in == "2" {
			return 0, false, false
		}
		return len(in), true, false
	})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

// TestStringTo_IntegersConditional_Empty verifies empty input.
func TestStringTo_IntegersConditional_Empty(t *testing.T) {
	result := converters.StringTo.IntegersConditional("", ",", func(in string) (int, bool, bool) {
		return 0, true, false
	})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// TestStringTo_BytesConditional verifies conditional bytes parsing.
func TestStringTo_BytesConditional(t *testing.T) {
	result := converters.StringTo.BytesConditional("a,b", ",", func(in string) (byte, bool, bool) {
		return in[0], true, false
	})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
}

// TestStringTo_BytesConditional_Empty verifies empty input.
func TestStringTo_BytesConditional_Empty(t *testing.T) {
	result := converters.StringTo.BytesConditional("", ",", func(in string) (byte, bool, bool) {
		return 0, true, false
	})
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// TestStringTo_JsonBytes verifies JSON bytes wrapping.
func TestStringTo_JsonBytes(t *testing.T) {
	result := converters.StringTo.JsonBytes("hello")
	actual := args.Map{"result": string(result) != `"hello"`}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '\"hello\"', got ''", actual)
}

// TestBytesTo_String verifies bytes-to-string conversion.
func TestBytesTo_String(t *testing.T) {
	actual := args.Map{"result": converters.BytesTo.String([]byte("hello")) != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello'", actual)
	actual := args.Map{"result": converters.BytesTo.String(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	actual := args.Map{"result": converters.BytesTo.String([]byte{}) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for empty slice", actual)
}

// TestBytesTo_PtrString verifies bytes-to-string via PtrString.
func TestBytesTo_PtrString(t *testing.T) {
	actual := args.Map{"result": converters.BytesTo.PtrString([]byte("test")) != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test'", actual)
}

// TestBytesTo_PointerToBytes verifies pointer-to-bytes safe copy.
func TestBytesTo_PointerToBytes(t *testing.T) {
	result := converters.BytesTo.PointerToBytes(nil)
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	result = converters.BytesTo.PointerToBytes([]byte{1, 2})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestUnsafeBytesToStringWithErr verifies unsafe conversion.
func TestUnsafeBytesToStringWithErr(t *testing.T) {
	s, err := converters.UnsafeBytesToStringWithErr([]byte("hello"))
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello', got '', err=", actual)
	_, err = converters.UnsafeBytesToStringWithErr(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// TestUnsafeBytesToString verifies unsafe conversion without error.
func TestUnsafeBytesToString(t *testing.T) {
	actual := args.Map{"result": converters.UnsafeBytesToString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	actual := args.Map{"result": converters.UnsafeBytesToString([]byte("test")) != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test'", actual)
}

// TestUnsafeBytesToStrings verifies safe byte-to-strings.
func TestUnsafeBytesToStrings(t *testing.T) {
	result := converters.UnsafeBytesToStrings(nil)
	actual := args.Map{"result": result != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil", actual)
	result = converters.UnsafeBytesToStrings([]byte{65, 66})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestUnsafeBytesToStringPtr verifies nil and non-nil.
func TestUnsafeBytesToStringPtr(t *testing.T) {
	actual := args.Map{"result": converters.UnsafeBytesToStringPtr(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil", actual)
	ptr := converters.UnsafeBytesToStringPtr([]byte("ok"))
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// TestUnsafeBytesPtrToStringPtr verifies pointer-based unsafe conversion.
func TestUnsafeBytesPtrToStringPtr(t *testing.T) {
	actual := args.Map{"result": converters.UnsafeBytesPtrToStringPtr(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil", actual)
	ptr := converters.UnsafeBytesPtrToStringPtr([]byte("ok"))
	actual := args.Map{"result": ptr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// TestAnyTo_ToString verifies any-to-string.
func TestAnyTo_ToString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToString(false, nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	r := converters.AnyTo.ToString(false, "hello")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	r = converters.AnyTo.ToString(true, "hello")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty for full name", actual)
}

// TestAnyTo_String verifies String method.
func TestAnyTo_String(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.String(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	actual := args.Map{"result": converters.AnyTo.String(42) == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// TestAnyTo_FullString verifies FullString.
func TestAnyTo_FullString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.FullString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_StringWithType verifies type-included string.
func TestAnyTo_StringWithType(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.StringWithType(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_ToSafeSerializedString verifies safe serialization.
func TestAnyTo_ToSafeSerializedString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToSafeSerializedString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	r := converters.AnyTo.ToSafeSerializedString([]byte("test"))
	actual := args.Map{"result": r != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'test', got ''", actual)
	r = converters.AnyTo.ToSafeSerializedString(42)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty for int", actual)
}

// TestAnyTo_ToSafeSerializedStringSprintValue verifies sprint value.
func TestAnyTo_ToSafeSerializedStringSprintValue(t *testing.T) {
	r := converters.AnyTo.ToSafeSerializedStringSprintValue("test")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// TestAnyTo_Bytes verifies byte conversion.
func TestAnyTo_Bytes(t *testing.T) {
	r := converters.AnyTo.Bytes([]byte{1, 2})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 bytes", actual)
	r = converters.AnyTo.Bytes("hello")
	actual := args.Map{"result": string(r) != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'hello'", actual)
	r = converters.AnyTo.Bytes(42)
	actual := args.Map{"result": len(r) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty for int JSON", actual)
	r = converters.AnyTo.Bytes([]byte(nil))
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil bytes", actual)
}

// TestAnyTo_ToPrettyJson verifies pretty JSON.
func TestAnyTo_ToPrettyJson(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToPrettyJson(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	r := converters.AnyTo.ToPrettyJson(map[string]int{"a": 1})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty JSON", actual)
}

// TestAnyTo_ValueString verifies ValueString.
func TestAnyTo_ValueString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ValueString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_ToValueString verifies ToValueString.
func TestAnyTo_ToValueString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToValueString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_ToValueStringWithType verifies type-included value string.
func TestAnyTo_ToValueStringWithType(t *testing.T) {
	r := converters.AnyTo.ToValueStringWithType(nil)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return type format", actual)
	r = converters.AnyTo.ToValueStringWithType(42)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// TestAnyTo_ToFullNameValueString verifies full name value string.
func TestAnyTo_ToFullNameValueString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToFullNameValueString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_ItemsJoin verifies items join.
func TestAnyTo_ItemsJoin(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ItemsJoin(",", nil...) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	r := converters.AnyTo.ItemsJoin(",", "a", "b")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// TestAnyTo_ToItemsThenJoin verifies items then join.
func TestAnyTo_ToItemsThenJoin(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToItemsThenJoin(true, ",", nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_SmartString verifies smart string.
func TestAnyTo_SmartString(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.SmartString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestAnyTo_SmartStringsOf verifies smart strings.
func TestAnyTo_SmartStringsOf(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.SmartStringsOf() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty", actual)
}

// TestStringsTo_Hashset verifies hashset creation.
func TestStringsTo_Hashset(t *testing.T) {
	result := converters.StringsTo.Hashset([]string{"a", "b"})
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestStringsTo_PointerStrings verifies pointer strings.
func TestStringsTo_PointerStrings(t *testing.T) {
	result := converters.StringsTo.PointerStrings(nil)
	actual := args.Map{"result": result == nil || len(*result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil input should return empty pointer slice", actual)
	input := []string{"a", "b"}
	result = converters.StringsTo.PointerStrings(&input)
	actual := args.Map{"result": len(*result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestStringsTo_PointerStringsCopy verifies copy pointer strings.
func TestStringsTo_PointerStringsCopy(t *testing.T) {
	result := converters.StringsTo.PointerStringsCopy(nil)
	actual := args.Map{"result": result == nil || len(*result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil input should return empty pointer slice", actual)
	input := []string{"x"}
	result = converters.StringsTo.PointerStringsCopy(&input)
	actual := args.Map{"result": len(*result) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// TestStringsTo_IntegersWithDefaults verifies multi-integer defaults.
func TestStringsTo_IntegersWithDefaults(t *testing.T) {
	r := converters.StringsTo.IntegersWithDefaults(-1, "1", "abc", "3")
	actual := args.Map{"result": len(r.Values) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": r.Values[1] != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

// TestStringsTo_IntegersConditional verifies conditional processing.
func TestStringsTo_IntegersConditional(t *testing.T) {
	r := converters.StringsTo.IntegersConditional(func(in string) (int, bool, bool) {
		return len(in), true, false
	}, "a", "bb")
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestStringsTo_IntegersSkipErrors verifies skip errors.
func TestStringsTo_IntegersSkipErrors(t *testing.T) {
	r := converters.StringsTo.IntegersSkipErrors("1", "abc", "3")
	actual := args.Map{"result": len(r) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// TestStringsTo_IntegersSkipAndDefaultValue verifies skip and default.
func TestStringsTo_IntegersSkipAndDefaultValue(t *testing.T) {
	r := converters.StringsTo.IntegersSkipAndDefaultValue(-1, "-", "1", "-", "abc")
	actual := args.Map{"result": len(r) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// TestStringsTo_IntegersSkipMapAndDefaultValue verifies skip map.
func TestStringsTo_IntegersSkipMapAndDefaultValue(t *testing.T) {
	skipMap := map[string]bool{"-": true}
	r := converters.StringsTo.IntegersSkipMapAndDefaultValue(-1, skipMap, "1", "-", "abc")
	actual := args.Map{"result": len(r) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// TestStringsTo_BytesWithDefaults verifies byte defaults.
func TestStringsTo_BytesWithDefaults(t *testing.T) {
	r := converters.StringsTo.BytesWithDefaults(0, "1", "abc", "300")
	actual := args.Map{"result": len(r.Values) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// TestStringsTo_BytesConditional verifies conditional bytes.
func TestStringsTo_BytesConditional(t *testing.T) {
	r := converters.StringsTo.BytesConditional(func(in string) (byte, bool, bool) {
		return in[0], true, false
	}, []string{"a", "b"})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestStringsTo_Csv verifies CSV generation.
func TestStringsTo_Csv(t *testing.T) {
	r := converters.StringsTo.Csv(false, "a", "b")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty CSV", actual)
}

// TestStringsTo_CsvUsingPtrStrings verifies nil-safe CSV.
func TestStringsTo_CsvUsingPtrStrings(t *testing.T) {
	actual := args.Map{"result": converters.StringsTo.CsvUsingPtrStrings(false, nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// TestStringsTo_CsvWithIndexes verifies indexed CSV.
func TestStringsTo_CsvWithIndexes(t *testing.T) {
	r := converters.StringsTo.CsvWithIndexes([]string{"a", "b"})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// TestStringsTo_MapConverter verifies map converter.
func TestStringsTo_MapConverter(t *testing.T) {
	mc := converters.StringsTo.MapConverter("a:1", "b:2")
	actual := args.Map{"result": mc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// TestStringsToMapConverter_Methods verifies StringsToMapConverter methods.
func TestStringsToMapConverter_Methods(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a:1", "b:2"})
	actual := args.Map{"result": mc.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
	actual := args.Map{"result": mc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
	actual := args.Map{"result": mc.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "last index should be 1", actual)
	ss := mc.SafeStrings()
	actual := args.Map{"result": len(ss) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	var nilMc *converters.StringsToMapConverter
	actual := args.Map{"result": nilMc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil length should be 0", actual)
}

// TestStringsTo_Float64sConditional verifies conditional float parsing.
func TestStringsTo_Float64sConditional(t *testing.T) {
	r := converters.StringsTo.Float64sConditional(func(in string) (float64, bool, bool) {
		return 1.0, true, false
	}, []string{"a", "b"})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}
