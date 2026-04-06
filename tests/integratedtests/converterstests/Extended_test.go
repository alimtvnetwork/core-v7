package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
)

// TestStringTo_IntegerWithDefault verifies integer conversion with default.
func TestStringTo_IntegerWithDefault(t *testing.T) {
	for _, tc := range stringToIntegerWithDefaultCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			val, ok := converters.StringTo.IntegerWithDefault(tc.input, tc.defaultVal)

			// Assert
			if val != tc.expectedVal {
				t.Errorf("expected %d, got %d", tc.expectedVal, val)
			}
			if ok != tc.expectedOk {
				t.Errorf("expected ok=%v, got %v", tc.expectedOk, ok)
			}
		})
	}
}

// TestStringTo_Integer verifies integer conversion with error.
func TestStringTo_Integer(t *testing.T) {
	// Act
	val, err := converters.StringTo.Integer("42")

	// Assert
	if err != nil || val != 42 {
		t.Errorf("expected 42, got %d, err=%v", val, err)
	}

	_, err = converters.StringTo.Integer("abc")
	if err == nil {
		t.Error("expected error for non-numeric")
	}
}

// TestStringTo_IntegerDefault verifies default integer conversion.
func TestStringTo_IntegerDefault(t *testing.T) {
	if converters.StringTo.IntegerDefault("10") != 10 {
		t.Error("expected 10")
	}
	if converters.StringTo.IntegerDefault("abc") != 0 {
		t.Error("expected 0 for invalid")
	}
}

// TestStringTo_Float64 verifies float64 conversion.
func TestStringTo_Float64(t *testing.T) {
	val, err := converters.StringTo.Float64("3.14")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val < 3.13 || val > 3.15 {
		t.Errorf("expected ~3.14, got %f", val)
	}

	_, err = converters.StringTo.Float64("abc")
	if err == nil {
		t.Error("expected error for non-numeric")
	}
}

// TestStringTo_Float64Default verifies default float conversion.
func TestStringTo_Float64Default(t *testing.T) {
	val, ok := converters.StringTo.Float64Default("2.5", 0.0)
	if !ok || val != 2.5 {
		t.Errorf("expected 2.5, got %f", val)
	}
	val, ok = converters.StringTo.Float64Default("abc", 9.9)
	if ok || val != 9.9 {
		t.Errorf("expected 9.9 default, got %f", val)
	}
}

// TestStringTo_Float64Conditional verifies deprecated conditional.
func TestStringTo_Float64Conditional(t *testing.T) {
	val, ok := converters.StringTo.Float64Conditional("2.5", 0.0)
	if !ok || val != 2.5 {
		t.Errorf("expected 2.5, got %f", val)
	}
}

// TestStringTo_Byte verifies byte conversion.
func TestStringTo_Byte(t *testing.T) {
	for _, tc := range stringToByteCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := converters.StringTo.Byte(tc.input)
			if tc.expectErr && err == nil {
				t.Error("expected error")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tc.expectErr && val != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, val)
			}
		})
	}
}

// TestStringTo_ByteWithDefault verifies byte with default.
func TestStringTo_ByteWithDefault(t *testing.T) {
	val, ok := converters.StringTo.ByteWithDefault("100", 0)
	if !ok || val != 100 {
		t.Errorf("expected 100, got %d", val)
	}
	val, ok = converters.StringTo.ByteWithDefault("abc", 55)
	if ok || val != 55 {
		t.Errorf("expected 55 default, got %d", val)
	}
}

// TestStringTo_IntegersWithDefaults verifies multi-integer parsing.
func TestStringTo_IntegersWithDefaults(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("1,2,abc", ",", -1)
	if len(result.Values) != 3 {
		t.Errorf("expected 3 values, got %d", len(result.Values))
	}
	if result.Values[2] != -1 {
		t.Errorf("expected default -1 for invalid, got %d", result.Values[2])
	}
	if result.CombinedError == nil {
		t.Error("expected combined error")
	}
}

// TestStringTo_IntegersWithDefaults_Empty verifies empty input.
func TestStringTo_IntegersWithDefaults_Empty(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("", ",", -1)
	if len(result.Values) != 0 {
		t.Errorf("expected 0 values, got %d", len(result.Values))
	}
}

// TestStringTo_IntegersConditional verifies conditional integer parsing.
func TestStringTo_IntegersConditional(t *testing.T) {
	result := converters.StringTo.IntegersConditional("1,2,3", ",", func(in string) (int, bool, bool) {
		if in == "2" {
			return 0, false, false
		}
		return len(in), true, false
	})
	if len(result) != 2 {
		t.Errorf("expected 2 items, got %d", len(result))
	}
}

// TestStringTo_IntegersConditional_Empty verifies empty input.
func TestStringTo_IntegersConditional_Empty(t *testing.T) {
	result := converters.StringTo.IntegersConditional("", ",", func(in string) (int, bool, bool) {
		return 0, true, false
	})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// TestStringTo_BytesConditional verifies conditional bytes parsing.
func TestStringTo_BytesConditional(t *testing.T) {
	result := converters.StringTo.BytesConditional("a,b", ",", func(in string) (byte, bool, bool) {
		return in[0], true, false
	})
	if len(result) != 2 {
		t.Errorf("expected 2 items, got %d", len(result))
	}
}

// TestStringTo_BytesConditional_Empty verifies empty input.
func TestStringTo_BytesConditional_Empty(t *testing.T) {
	result := converters.StringTo.BytesConditional("", ",", func(in string) (byte, bool, bool) {
		return 0, true, false
	})
	if len(result) != 0 {
		t.Errorf("expected 0, got %d", len(result))
	}
}

// TestStringTo_JsonBytes verifies JSON bytes wrapping.
func TestStringTo_JsonBytes(t *testing.T) {
	result := converters.StringTo.JsonBytes("hello")
	if string(result) != `"hello"` {
		t.Errorf("expected '\"hello\"', got '%s'", string(result))
	}
}

// TestBytesTo_String verifies bytes-to-string conversion.
func TestBytesTo_String(t *testing.T) {
	if converters.BytesTo.String([]byte("hello")) != "hello" {
		t.Error("expected 'hello'")
	}
	if converters.BytesTo.String(nil) != "" {
		t.Error("expected empty for nil")
	}
	if converters.BytesTo.String([]byte{}) != "" {
		t.Error("expected empty for empty slice")
	}
}

// TestBytesTo_PtrString verifies bytes-to-string via PtrString.
func TestBytesTo_PtrString(t *testing.T) {
	if converters.BytesTo.PtrString([]byte("test")) != "test" {
		t.Error("expected 'test'")
	}
}

// TestBytesTo_PointerToBytes verifies pointer-to-bytes safe copy.
func TestBytesTo_PointerToBytes(t *testing.T) {
	result := converters.BytesTo.PointerToBytes(nil)
	if len(result) != 0 {
		t.Error("expected empty for nil")
	}
	result = converters.BytesTo.PointerToBytes([]byte{1, 2})
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

// TestUnsafeBytesToStringWithErr verifies unsafe conversion.
func TestUnsafeBytesToStringWithErr(t *testing.T) {
	s, err := converters.UnsafeBytesToStringWithErr([]byte("hello"))
	if err != nil || s != "hello" {
		t.Errorf("expected 'hello', got '%s', err=%v", s, err)
	}
	_, err = converters.UnsafeBytesToStringWithErr(nil)
	if err == nil {
		t.Error("expected error for nil")
	}
}

// TestUnsafeBytesToString verifies unsafe conversion without error.
func TestUnsafeBytesToString(t *testing.T) {
	if converters.UnsafeBytesToString(nil) != "" {
		t.Error("expected empty for nil")
	}
	if converters.UnsafeBytesToString([]byte("test")) != "test" {
		t.Error("expected 'test'")
	}
}

// TestUnsafeBytesToStrings verifies safe byte-to-strings.
func TestUnsafeBytesToStrings(t *testing.T) {
	result := converters.UnsafeBytesToStrings(nil)
	if result != nil {
		t.Error("expected nil for nil")
	}
	result = converters.UnsafeBytesToStrings([]byte{65, 66})
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

// TestUnsafeBytesToStringPtr verifies nil and non-nil.
func TestUnsafeBytesToStringPtr(t *testing.T) {
	if converters.UnsafeBytesToStringPtr(nil) != nil {
		t.Error("expected nil for nil")
	}
	ptr := converters.UnsafeBytesToStringPtr([]byte("ok"))
	if ptr == nil {
		t.Error("expected non-nil")
	}
}

// TestUnsafeBytesPtrToStringPtr verifies pointer-based unsafe conversion.
func TestUnsafeBytesPtrToStringPtr(t *testing.T) {
	if converters.UnsafeBytesPtrToStringPtr(nil) != nil {
		t.Error("expected nil for nil")
	}
	ptr := converters.UnsafeBytesPtrToStringPtr([]byte("ok"))
	if ptr == nil {
		t.Error("expected non-nil")
	}
}

// TestAnyTo_ToString verifies any-to-string.
func TestAnyTo_ToString(t *testing.T) {
	if converters.AnyTo.ToString(false, nil) != "" {
		t.Error("expected empty for nil")
	}
	r := converters.AnyTo.ToString(false, "hello")
	if r == "" {
		t.Error("expected non-empty")
	}
	r = converters.AnyTo.ToString(true, "hello")
	if r == "" {
		t.Error("expected non-empty for full name")
	}
}

// TestAnyTo_String verifies String method.
func TestAnyTo_String(t *testing.T) {
	if converters.AnyTo.String(nil) != "" {
		t.Error("nil should return empty")
	}
	if converters.AnyTo.String(42) == "" {
		t.Error("expected non-empty")
	}
}

// TestAnyTo_FullString verifies FullString.
func TestAnyTo_FullString(t *testing.T) {
	if converters.AnyTo.FullString(nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_StringWithType verifies type-included string.
func TestAnyTo_StringWithType(t *testing.T) {
	if converters.AnyTo.StringWithType(nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_ToSafeSerializedString verifies safe serialization.
func TestAnyTo_ToSafeSerializedString(t *testing.T) {
	if converters.AnyTo.ToSafeSerializedString(nil) != "" {
		t.Error("nil should return empty")
	}
	r := converters.AnyTo.ToSafeSerializedString([]byte("test"))
	if r != "test" {
		t.Errorf("expected 'test', got '%s'", r)
	}
	r = converters.AnyTo.ToSafeSerializedString(42)
	if r == "" {
		t.Error("expected non-empty for int")
	}
}

// TestAnyTo_ToSafeSerializedStringSprintValue verifies sprint value.
func TestAnyTo_ToSafeSerializedStringSprintValue(t *testing.T) {
	r := converters.AnyTo.ToSafeSerializedStringSprintValue("test")
	if r == "" {
		t.Error("expected non-empty")
	}
}

// TestAnyTo_Bytes verifies byte conversion.
func TestAnyTo_Bytes(t *testing.T) {
	r := converters.AnyTo.Bytes([]byte{1, 2})
	if len(r) != 2 {
		t.Error("expected 2 bytes")
	}
	r = converters.AnyTo.Bytes("hello")
	if string(r) != "hello" {
		t.Error("expected 'hello'")
	}
	r = converters.AnyTo.Bytes(42)
	if len(r) == 0 {
		t.Error("expected non-empty for int JSON")
	}
	r = converters.AnyTo.Bytes([]byte(nil))
	if len(r) != 0 {
		t.Error("expected empty for nil bytes")
	}
}

// TestAnyTo_ToPrettyJson verifies pretty JSON.
func TestAnyTo_ToPrettyJson(t *testing.T) {
	if converters.AnyTo.ToPrettyJson(nil) != "" {
		t.Error("nil should return empty")
	}
	r := converters.AnyTo.ToPrettyJson(map[string]int{"a": 1})
	if r == "" {
		t.Error("expected non-empty JSON")
	}
}

// TestAnyTo_ValueString verifies ValueString.
func TestAnyTo_ValueString(t *testing.T) {
	if converters.AnyTo.ValueString(nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_ToValueString verifies ToValueString.
func TestAnyTo_ToValueString(t *testing.T) {
	if converters.AnyTo.ToValueString(nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_ToValueStringWithType verifies type-included value string.
func TestAnyTo_ToValueStringWithType(t *testing.T) {
	r := converters.AnyTo.ToValueStringWithType(nil)
	if r == "" {
		t.Error("nil should return type format")
	}
	r = converters.AnyTo.ToValueStringWithType(42)
	if r == "" {
		t.Error("expected non-empty")
	}
}

// TestAnyTo_ToFullNameValueString verifies full name value string.
func TestAnyTo_ToFullNameValueString(t *testing.T) {
	if converters.AnyTo.ToFullNameValueString(nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_ItemsJoin verifies items join.
func TestAnyTo_ItemsJoin(t *testing.T) {
	if converters.AnyTo.ItemsJoin(",", nil...) != "" {
		t.Error("nil should return empty")
	}
	r := converters.AnyTo.ItemsJoin(",", "a", "b")
	if r == "" {
		t.Error("expected non-empty")
	}
}

// TestAnyTo_ToItemsThenJoin verifies items then join.
func TestAnyTo_ToItemsThenJoin(t *testing.T) {
	if converters.AnyTo.ToItemsThenJoin(true, ",", nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_SmartString verifies smart string.
func TestAnyTo_SmartString(t *testing.T) {
	if converters.AnyTo.SmartString(nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestAnyTo_SmartStringsOf verifies smart strings.
func TestAnyTo_SmartStringsOf(t *testing.T) {
	if converters.AnyTo.SmartStringsOf() != "" {
		t.Error("empty should return empty")
	}
}

// TestStringsTo_Hashset verifies hashset creation.
func TestStringsTo_Hashset(t *testing.T) {
	result := converters.StringsTo.Hashset([]string{"a", "b"})
	if len(result) != 2 {
		t.Errorf("expected 2, got %d", len(result))
	}
}

// TestStringsTo_PointerStrings verifies pointer strings.
func TestStringsTo_PointerStrings(t *testing.T) {
	result := converters.StringsTo.PointerStrings(nil)
	if result == nil || len(*result) != 0 {
		t.Error("nil input should return empty pointer slice")
	}
	input := []string{"a", "b"}
	result = converters.StringsTo.PointerStrings(&input)
	if len(*result) != 2 {
		t.Errorf("expected 2, got %d", len(*result))
	}
}

// TestStringsTo_PointerStringsCopy verifies copy pointer strings.
func TestStringsTo_PointerStringsCopy(t *testing.T) {
	result := converters.StringsTo.PointerStringsCopy(nil)
	if result == nil || len(*result) != 0 {
		t.Error("nil input should return empty pointer slice")
	}
	input := []string{"x"}
	result = converters.StringsTo.PointerStringsCopy(&input)
	if len(*result) != 1 {
		t.Errorf("expected 1, got %d", len(*result))
	}
}

// TestStringsTo_IntegersWithDefaults verifies multi-integer defaults.
func TestStringsTo_IntegersWithDefaults(t *testing.T) {
	r := converters.StringsTo.IntegersWithDefaults(-1, "1", "abc", "3")
	if len(r.Values) != 3 {
		t.Errorf("expected 3, got %d", len(r.Values))
	}
	if r.Values[1] != -1 {
		t.Errorf("expected -1, got %d", r.Values[1])
	}
}

// TestStringsTo_IntegersConditional verifies conditional processing.
func TestStringsTo_IntegersConditional(t *testing.T) {
	r := converters.StringsTo.IntegersConditional(func(in string) (int, bool, bool) {
		return len(in), true, false
	}, "a", "bb")
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

// TestStringsTo_IntegersSkipErrors verifies skip errors.
func TestStringsTo_IntegersSkipErrors(t *testing.T) {
	r := converters.StringsTo.IntegersSkipErrors("1", "abc", "3")
	if len(r) != 3 {
		t.Errorf("expected 3, got %d", len(r))
	}
}

// TestStringsTo_IntegersSkipAndDefaultValue verifies skip and default.
func TestStringsTo_IntegersSkipAndDefaultValue(t *testing.T) {
	r := converters.StringsTo.IntegersSkipAndDefaultValue(-1, "-", "1", "-", "abc")
	if len(r) != 3 {
		t.Errorf("expected 3, got %d", len(r))
	}
}

// TestStringsTo_IntegersSkipMapAndDefaultValue verifies skip map.
func TestStringsTo_IntegersSkipMapAndDefaultValue(t *testing.T) {
	skipMap := map[string]bool{"-": true}
	r := converters.StringsTo.IntegersSkipMapAndDefaultValue(-1, skipMap, "1", "-", "abc")
	if len(r) != 3 {
		t.Errorf("expected 3, got %d", len(r))
	}
}

// TestStringsTo_BytesWithDefaults verifies byte defaults.
func TestStringsTo_BytesWithDefaults(t *testing.T) {
	r := converters.StringsTo.BytesWithDefaults(0, "1", "abc", "300")
	if len(r.Values) != 3 {
		t.Errorf("expected 3, got %d", len(r.Values))
	}
}

// TestStringsTo_BytesConditional verifies conditional bytes.
func TestStringsTo_BytesConditional(t *testing.T) {
	r := converters.StringsTo.BytesConditional(func(in string) (byte, bool, bool) {
		return in[0], true, false
	}, []string{"a", "b"})
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

// TestStringsTo_Csv verifies CSV generation.
func TestStringsTo_Csv(t *testing.T) {
	r := converters.StringsTo.Csv(false, "a", "b")
	if r == "" {
		t.Error("expected non-empty CSV")
	}
}

// TestStringsTo_CsvUsingPtrStrings verifies nil-safe CSV.
func TestStringsTo_CsvUsingPtrStrings(t *testing.T) {
	if converters.StringsTo.CsvUsingPtrStrings(false, nil) != "" {
		t.Error("nil should return empty")
	}
}

// TestStringsTo_CsvWithIndexes verifies indexed CSV.
func TestStringsTo_CsvWithIndexes(t *testing.T) {
	r := converters.StringsTo.CsvWithIndexes([]string{"a", "b"})
	if r == "" {
		t.Error("expected non-empty")
	}
}

// TestStringsTo_MapConverter verifies map converter.
func TestStringsTo_MapConverter(t *testing.T) {
	mc := converters.StringsTo.MapConverter("a:1", "b:2")
	if mc.Length() != 2 {
		t.Errorf("expected 2, got %d", mc.Length())
	}
}

// TestStringsToMapConverter_Methods verifies StringsToMapConverter methods.
func TestStringsToMapConverter_Methods(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a:1", "b:2"})
	if mc.IsEmpty() {
		t.Error("should not be empty")
	}
	if !mc.HasAnyItem() {
		t.Error("should have items")
	}
	if mc.LastIndex() != 1 {
		t.Error("last index should be 1")
	}
	ss := mc.SafeStrings()
	if len(ss) != 2 {
		t.Errorf("expected 2, got %d", len(ss))
	}

	var nilMc *converters.StringsToMapConverter
	if nilMc.Length() != 0 {
		t.Error("nil length should be 0")
	}
}

// TestStringsTo_Float64sConditional verifies conditional float parsing.
func TestStringsTo_Float64sConditional(t *testing.T) {
	r := converters.StringsTo.Float64sConditional(func(in string) (float64, bool, bool) {
		return 1.0, true, false
	}, []string{"a", "b"})
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}
