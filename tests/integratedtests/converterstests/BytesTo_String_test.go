package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BytesTo.String ──

func Test_BytesTo_String_Empty(t *testing.T) {
	// Arrange
	result := converters.BytesTo.String([]byte{})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "BytesTo.String returns empty -- empty", actual)
}

func Test_BytesTo_String_NonEmpty(t *testing.T) {
	// Arrange
	result := converters.BytesTo.String([]byte("hello"))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesTo.String returns empty -- non-empty", actual)
}

// ── BytesTo.PtrString ──

func Test_BytesTo_PtrString_Empty(t *testing.T) {
	// Arrange
	result := converters.BytesTo.PtrString([]byte{})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "BytesTo.PtrString returns empty -- empty", actual)
}

func Test_BytesTo_PtrString_NonEmpty(t *testing.T) {
	// Arrange
	result := converters.BytesTo.PtrString([]byte("world"))

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "world"}
	expected.ShouldBeEqual(t, 0, "BytesTo.PtrString returns empty -- non-empty", actual)
}

// ── BytesTo.PointerToBytes ──

func Test_BytesTo_PointerToBytes_Nil(t *testing.T) {
	// Arrange
	result := converters.BytesTo.PointerToBytes(nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesTo.PointerToBytes returns nil -- nil", actual)
}

func Test_BytesTo_PointerToBytes_NonNil(t *testing.T) {
	// Arrange
	result := converters.BytesTo.PointerToBytes([]byte{1, 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": byte(1),
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.PointerToBytes returns nil -- non-nil", actual)
}

// ── StringTo.IntegerDefault ──

func Test_StringTo_IntegerDefault_Valid(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegerDefault("42")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegerDefault returns non-empty -- valid", actual)
}

func Test_StringTo_IntegerDefault_Invalid(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegerDefault("abc")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegerDefault returns error -- invalid", actual)
}

// ── StringTo.Integer ──

func Test_StringTo_Integer_Valid(t *testing.T) {
	// Arrange
	val, err := converters.StringTo.Integer("123")

	// Act
	actual := args.Map{
		"val": val,
		"isNilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": 123,
		"isNilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Integer returns non-empty -- valid", actual)
}

func Test_StringTo_Integer_Invalid(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Integer("abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Integer returns error -- invalid", actual)
}

// ── StringTo.IntegerMust ──

func Test_StringTo_IntegerMust_Valid(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegerMust("99")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": 99}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegerMust returns non-empty -- valid", actual)
}

func Test_StringTo_IntegerMust_Panics(t *testing.T) {
	// Arrange
	var didPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		converters.StringTo.IntegerMust("abc")
	}()

	// Act
	actual := args.Map{"panicked": didPanic}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegerMust panics -- panics", actual)
}

// ── StringTo.IntegersConditional ──

func Test_StringTo_IntegersConditional_Empty(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegersConditional("", ",", func(in string) (int, bool, bool) {
		return 0, false, false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegersConditional returns empty -- empty", actual)
}

func Test_StringTo_IntegersConditional_WithBreak(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegersConditional("1,2,3", ",", func(in string) (int, bool, bool) {
		if in == "2" {
			return 0, false, true
		}
		return 1, true, false
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegersConditional returns non-empty -- with break", actual)
}

// ── StringTo.IntegersWithDefaults ──

func Test_StringTo_IntegersWithDefaults_Empty(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegersWithDefaults("", ",", -1)

	// Act
	actual := args.Map{
		"len": len(result.Values),
		"nilErr": result.CombinedError == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegersWithDefaults returns empty -- empty", actual)
}

func Test_StringTo_IntegersWithDefaults_MixedValid(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegersWithDefaults("1,abc,3", ",", -1)

	// Act
	actual := args.Map{
		"len": len(result.Values),
		"first": result.Values[0],
		"second": result.Values[1],
		"third": result.Values[2],
		"hasErr": result.CombinedError != nil,
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
		"second": -1,
		"third": 3,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.IntegersWithDefaults returns non-empty -- mixed", actual)
}

// ── StringTo.Float64 ──

func Test_StringTo_Float64_Valid(t *testing.T) {
	// Arrange
	val, err := converters.StringTo.Float64("3.14")

	// Act
	actual := args.Map{
		"isNilErr": err == nil,
		"above3": val > 3.0,
	}

	// Assert
	expected := args.Map{
		"isNilErr": true,
		"above3": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64 returns non-empty -- valid", actual)
}

func Test_StringTo_Float64_Invalid(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Float64("abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64 returns error -- invalid", actual)
}

// ── StringTo.Float64Default ──

func Test_StringTo_Float64Default_Valid(t *testing.T) {
	// Arrange
	val, ok := converters.StringTo.Float64Default("2.5", -1.0)

	// Act
	actual := args.Map{
		"ok": ok,
		"above2": val > 2.0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"above2": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64Default returns non-empty -- valid", actual)
}

func Test_StringTo_Float64Default_Invalid(t *testing.T) {
	// Arrange
	val, ok := converters.StringTo.Float64Default("abc", -1.0)

	// Act
	actual := args.Map{
		"ok": ok,
		"val": val,
	}

	// Assert
	expected := args.Map{
		"ok": false,
		"val": -1.0,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64Default returns error -- invalid", actual)
}

// ── StringTo.Float64Must ──

func Test_StringTo_Float64Must_Valid(t *testing.T) {
	// Arrange
	result := converters.StringTo.Float64Must("1.5")

	// Act
	actual := args.Map{"above1": result > 1.0}

	// Assert
	expected := args.Map{"above1": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Float64Must returns non-empty -- valid", actual)
}

// ── StringTo.Byte ──

func Test_StringTo_Byte_Empty(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Byte("")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns empty -- empty", actual)
}

func Test_StringTo_Byte_Zero(t *testing.T) {
	// Arrange
	val, err := converters.StringTo.Byte("0")

	// Act
	actual := args.Map{
		"val": val,
		"isNilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": byte(0),
		"isNilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- zero", actual)
}

func Test_StringTo_Byte_One(t *testing.T) {
	// Arrange
	val, err := converters.StringTo.Byte("1")

	// Act
	actual := args.Map{
		"val": val,
		"isNilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": byte(1),
		"isNilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- one", actual)
}

func Test_StringTo_Byte_Valid(t *testing.T) {
	// Arrange
	val, err := converters.StringTo.Byte("200")

	// Act
	actual := args.Map{
		"val": val,
		"isNilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": byte(200),
		"isNilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns non-empty -- valid", actual)
}

func Test_StringTo_Byte_Negative(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Byte("-1")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- negative", actual)
}

func Test_StringTo_Byte_Over255(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Byte("256")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- over 255", actual)
}

func Test_StringTo_Byte_NotNumber(t *testing.T) {
	// Arrange
	_, err := converters.StringTo.Byte("abc")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringTo.Byte returns correct value -- not number", actual)
}

// ── StringTo.ByteWithDefault ──

func Test_StringTo_ByteWithDefault_Valid(t *testing.T) {
	// Arrange
	val, ok := converters.StringTo.ByteWithDefault("100", 0)

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": byte(100),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.ByteWithDefault returns non-empty -- valid", actual)
}

func Test_StringTo_ByteWithDefault_Invalid(t *testing.T) {
	// Arrange
	val, ok := converters.StringTo.ByteWithDefault("abc", 99)

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": byte(99),
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "StringTo.ByteWithDefault returns error -- invalid", actual)
}

// ── StringTo.JsonBytes ──

func Test_StringTo_JsonBytes(t *testing.T) {
	// Arrange
	result := converters.StringTo.JsonBytes("name")

	// Act
	actual := args.Map{"notEmpty": len(result) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringTo.JsonBytes returns correct value -- with args", actual)
}

// ── StringsTo.Hashset ──

func Test_StringsTo_Hashset(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Hashset([]string{"a", "b", "a"})

	// Act
	actual := args.Map{
		"len": len(result),
		"hasA": result["a"],
		"hasB": result["b"],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasA": true,
		"hasB": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.Hashset returns correct value -- with args", actual)
}

// ── StringsTo.IntegersSkipErrors ──

func Test_StringsTo_IntegersSkipErrors(t *testing.T) {
	// Arrange
	result := converters.StringsTo.IntegersSkipErrors("1", "abc", "3")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"third": result[2],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
		"third": 3,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.IntegersSkipErrors returns error -- with args", actual)
}

// ── StringsTo.IntegersSkipAndDefaultValue ──

func Test_StringsTo_IntegersSkipAndDefaultValue(t *testing.T) {
	// Arrange
	result := converters.StringsTo.IntegersSkipAndDefaultValue(-1, "skip", "1", "skip", "abc")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.IntegersSkipAndDefaultValue returns correct value -- with args", actual)
}

// ── StringsTo.Csv ──

func Test_StringsTo_Csv(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Csv(false, "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsTo.Csv returns correct value -- with args", actual)
}

// ── StringsTo.CsvUsingPtrStrings ──

func Test_StringsTo_CsvUsingPtrStrings_Nil(t *testing.T) {
	// Arrange
	result := converters.StringsTo.CsvUsingPtrStrings(false, nil)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StringsTo.CsvUsingPtrStrings returns nil -- nil", actual)
}

// ── StringsTo.MapConverter ──

func Test_StringsTo_MapConverter(t *testing.T) {
	// Arrange
	result := converters.StringsTo.MapConverter("a:b", "c:d")

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsTo.MapConverter returns correct value -- with args", actual)
}

// ── StringsTo.PointerStrings ──

func Test_StringsTo_PointerStrings_Nil(t *testing.T) {
	// Arrange
	result := converters.StringsTo.PointerStrings(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsTo.PointerStrings returns nil -- nil", actual)
}

func Test_StringsTo_PointerStrings_NonNil(t *testing.T) {
	// Arrange
	lines := []string{"a", "b"}
	result := converters.StringsTo.PointerStrings(&lines)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsTo.PointerStrings returns nil -- non-nil", actual)
}

// ── StringsTo.PointerStringsCopy ──

func Test_StringsTo_PointerStringsCopy_Nil(t *testing.T) {
	// Arrange
	result := converters.StringsTo.PointerStringsCopy(nil)

	// Act
	actual := args.Map{"len": len(*result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsTo.PointerStringsCopy returns nil -- nil", actual)
}

func Test_StringsTo_PointerStringsCopy_NonNil(t *testing.T) {
	// Arrange
	lines := []string{"x", "y"}
	result := converters.StringsTo.PointerStringsCopy(&lines)

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": *(*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "x",
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.PointerStringsCopy returns nil -- non-nil", actual)
}

// ── StringsTo.HashmapTrimColon ──

func Test_StringsTo_HashmapTrimColon(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapTrimColon("key:val")

	// Act
	actual := args.Map{"val": result["key"]}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "StringsTo.HashmapTrimColon returns correct value -- with args", actual)
}

// ── StringsTo.HashmapTrimHyphen ──

func Test_StringsTo_HashmapTrimHyphen(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapTrimHyphen("key-val")

	// Act
	actual := args.Map{"val": result["key"]}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "StringsTo.HashmapTrimHyphen returns correct value -- with args", actual)
}

// ── StringsTo.HashmapOptions ──

func Test_StringsTo_HashmapOptions(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapOptions(true, "=", "k=v")

	// Act
	actual := args.Map{"val": result["k"]}

	// Assert
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "StringsTo.HashmapOptions returns correct value -- with args", actual)
}

// ── StringsTo.HashmapTrim ──

func Test_StringsTo_HashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapTrim(":", []string{"a:b"})

	// Act
	actual := args.Map{"val": result["a"]}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "StringsTo.HashmapTrim returns correct value -- with args", actual)
}

// ── StringsTo.BytesWithDefaults ──

func Test_StringsTo_BytesWithDefaults_Valid(t *testing.T) {
	// Arrange
	result := converters.StringsTo.BytesWithDefaults(0, "10", "20")

	// Act
	actual := args.Map{
		"len": len(result.Values),
		"first": result.Values[0],
		"second": result.Values[1],
		"nilErr": result.CombinedError == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": byte(10),
		"second": byte(20),
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.BytesWithDefaults returns non-empty -- valid", actual)
}

func Test_StringsTo_BytesWithDefaults_Invalid(t *testing.T) {
	// Arrange
	result := converters.StringsTo.BytesWithDefaults(0, "abc")

	// Act
	actual := args.Map{
		"first": result.Values[0],
		"hasErr": result.CombinedError != nil,
	}

	// Assert
	expected := args.Map{
		"first": byte(0),
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.BytesWithDefaults returns error -- invalid", actual)
}

func Test_StringsTo_BytesWithDefaults_OutOfRange(t *testing.T) {
	// Arrange
	result := converters.StringsTo.BytesWithDefaults(0, "300")

	// Act
	actual := args.Map{
		"first": result.Values[0],
		"hasErr": result.CombinedError != nil,
	}

	// Assert
	expected := args.Map{
		"first": byte(0),
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo.BytesWithDefaults returns non-empty -- out of range", actual)
}
