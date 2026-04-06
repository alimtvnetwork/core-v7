package converterstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── StringTo: IntegerWithDefault ──

func Test_Cov3_StringTo_IntegerWithDefault_Valid(t *testing.T) {
	val, ok := converters.StringTo.IntegerWithDefault("42", -1)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "IntegerWithDefault returns non-empty -- valid", actual)
}

func Test_Cov3_StringTo_IntegerWithDefault_Empty(t *testing.T) {
	val, ok := converters.StringTo.IntegerWithDefault("", -1)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": -1, "ok": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithDefault returns empty -- empty", actual)
}

func Test_Cov3_StringTo_IntegerWithDefault_Invalid(t *testing.T) {
	val, ok := converters.StringTo.IntegerWithDefault("abc", -1)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": -1, "ok": false}
	expected.ShouldBeEqual(t, 0, "IntegerWithDefault returns error -- invalid", actual)
}

// ── StringTo: IntegersWithDefaults ──

func Test_Cov3_StringTo_IntegersWithDefaults_Valid(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("1,2,3", ",", -1)
	actual := args.Map{"len": len(result.Values), "noErr": result.CombinedError == nil}
	expected := args.Map{"len": 3, "noErr": true}
	expected.ShouldBeEqual(t, 0, "IntegersWithDefaults returns non-empty -- valid", actual)
}

func Test_Cov3_StringTo_IntegersWithDefaults_Empty(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("", ",", -1)
	actual := args.Map{"len": len(result.Values), "noErr": result.CombinedError == nil}
	expected := args.Map{"len": 0, "noErr": true}
	expected.ShouldBeEqual(t, 0, "IntegersWithDefaults returns empty -- empty", actual)
}

func Test_Cov3_StringTo_IntegersWithDefaults_WithErrors(t *testing.T) {
	result := converters.StringTo.IntegersWithDefaults("1,abc,3", ",", -1)
	actual := args.Map{"len": len(result.Values), "hasErr": result.CombinedError != nil}
	expected := args.Map{"len": 3, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "IntegersWithDefaults returns error -- with errors", actual)
}

// ── StringTo: IntegersConditional ──

func Test_Cov3_StringTo_IntegersConditional_Empty(t *testing.T) {
	result := converters.StringTo.IntegersConditional("", ",", func(in string) (int, bool, bool) {
		return 0, true, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IntegersConditional returns empty -- empty", actual)
}

func Test_Cov3_StringTo_IntegersConditional_TakeAndBreak(t *testing.T) {
	result := converters.StringTo.IntegersConditional("1,2,3", ",", func(in string) (int, bool, bool) {
		return 99, true, in == "2"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "IntegersConditional returns correct value -- take and break", actual)
}

func Test_Cov3_StringTo_IntegersConditional_Skip(t *testing.T) {
	result := converters.StringTo.IntegersConditional("1,2,3", ",", func(in string) (int, bool, bool) {
		return 0, false, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "IntegersConditional returns correct value -- skip all", actual)
}

// ── StringTo: IntegerMust ──

func Test_Cov3_StringTo_IntegerMust_Valid(t *testing.T) {
	result := converters.StringTo.IntegerMust("42")
	actual := args.Map{"val": result}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "IntegerMust returns non-empty -- valid", actual)
}

// ── StringTo: IntegerDefault ──

func Test_Cov3_StringTo_IntegerDefault_Valid(t *testing.T) {
	actual := args.Map{
		"valid":   converters.StringTo.IntegerDefault("42"),
		"invalid": converters.StringTo.IntegerDefault("abc"),
	}
	expected := args.Map{
		"valid":   42,
		"invalid": 0,
	}
	expected.ShouldBeEqual(t, 0, "IntegerDefault returns correct value -- with args", actual)
}

// ── StringTo: Float64Must ──

func Test_Cov3_StringTo_Float64Must_Valid(t *testing.T) {
	result := converters.StringTo.Float64Must("3.14")
	actual := args.Map{"gt3": result > 3.0}
	expected := args.Map{"gt3": true}
	expected.ShouldBeEqual(t, 0, "Float64Must returns non-empty -- valid", actual)
}

// ── StringTo: Float64Default ──

func Test_Cov3_StringTo_Float64Default_Valid(t *testing.T) {
	val, ok := converters.StringTo.Float64Default("3.14", 0.0)
	actual := args.Map{"gt3": val > 3.0, "ok": ok}
	expected := args.Map{"gt3": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "Float64Default returns non-empty -- valid", actual)
}

func Test_Cov3_StringTo_Float64Default_Invalid(t *testing.T) {
	val, ok := converters.StringTo.Float64Default("abc", 9.9)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 9.9, "ok": false}
	expected.ShouldBeEqual(t, 0, "Float64Default returns error -- invalid", actual)
}

// ── StringTo: Float64Conditional (deprecated alias) ──

func Test_Cov3_StringTo_Float64Conditional(t *testing.T) {
	val, ok := converters.StringTo.Float64Conditional("3.14", 0.0)
	actual := args.Map{"gt3": val > 3.0, "ok": ok}
	expected := args.Map{"gt3": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "Float64Conditional returns non-empty -- valid", actual)
}

// ── StringTo: ByteWithDefault ──

func Test_Cov3_StringTo_ByteWithDefault_Valid(t *testing.T) {
	val, ok := converters.StringTo.ByteWithDefault("42", 0)
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "ByteWithDefault returns non-empty -- valid", actual)
}

func Test_Cov3_StringTo_ByteWithDefault_Invalid(t *testing.T) {
	val, ok := converters.StringTo.ByteWithDefault("abc", 99)
	actual := args.Map{"val": int(val), "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "ByteWithDefault returns error -- invalid", actual)
}

// ── StringTo: BytesConditional ──

func Test_Cov3_StringTo_BytesConditional_Empty(t *testing.T) {
	result := converters.StringTo.BytesConditional("", ",", func(in string) (byte, bool, bool) {
		return 0, true, false
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesConditional returns empty -- empty", actual)
}

func Test_Cov3_StringTo_BytesConditional_WithBreak(t *testing.T) {
	result := converters.StringTo.BytesConditional("a,b,c", ",", func(in string) (byte, bool, bool) {
		return in[0], true, in == "b"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesConditional returns non-empty -- with break", actual)
}

// ── StringTo: Byte additional branches ──

func Test_Cov3_StringTo_Byte_Zero(t *testing.T) {
	val, err := converters.StringTo.Byte("0")
	actual := args.Map{"val": int(val), "noErr": err == nil}
	expected := args.Map{"val": 0, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- zero", actual)
}

func Test_Cov3_StringTo_Byte_One(t *testing.T) {
	val, err := converters.StringTo.Byte("1")
	actual := args.Map{"val": int(val), "noErr": err == nil}
	expected := args.Map{"val": 1, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- one", actual)
}

func Test_Cov3_StringTo_Byte_Empty(t *testing.T) {
	_, err := converters.StringTo.Byte("")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Byte returns empty -- empty", actual)
}

func Test_Cov3_StringTo_Byte_Negative(t *testing.T) {
	_, err := converters.StringTo.Byte("-1")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Byte returns correct value -- negative", actual)
}

// ── StringTo: JsonBytes ──

func Test_Cov3_StringTo_JsonBytes(t *testing.T) {
	result := converters.StringTo.JsonBytes("hello")
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonBytes returns correct value -- with args", actual)
}

// ── AnyTo: ToString ──

func Test_Cov3_AnyTo_ToString(t *testing.T) {
	actual := args.Map{
		"withFull": converters.AnyTo.ToString(true, "hello") != "",
		"noFull":   converters.AnyTo.ToString(false, "hello") != "",
		"nilVal":   converters.AnyTo.ToString(false, nil),
	}
	expected := args.Map{
		"withFull": true,
		"noFull":   true,
		"nilVal":   "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToString returns correct value -- with args", actual)
}

// ── AnyTo: String ──

func Test_Cov3_AnyTo_String(t *testing.T) {
	actual := args.Map{
		"val": converters.AnyTo.String("hello") != "",
		"nil": converters.AnyTo.String(nil),
	}
	expected := args.Map{
		"val": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.String returns correct value -- with args", actual)
}

// ── AnyTo: FullString ──

func Test_Cov3_AnyTo_FullString(t *testing.T) {
	actual := args.Map{
		"val": converters.AnyTo.FullString("hello") != "",
		"nil": converters.AnyTo.FullString(nil),
	}
	expected := args.Map{
		"val": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.FullString returns correct value -- with args", actual)
}

// ── AnyTo: StringWithType ──

func Test_Cov3_AnyTo_StringWithType(t *testing.T) {
	actual := args.Map{
		"val": converters.AnyTo.StringWithType(42) != "",
		"nil": converters.AnyTo.StringWithType(nil),
	}
	expected := args.Map{
		"val": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.StringWithType returns non-empty -- with args", actual)
}

// ── AnyTo: ToSafeSerializedString ──

func Test_Cov3_AnyTo_ToSafeSerializedString(t *testing.T) {
	actual := args.Map{
		"bytes":  converters.AnyTo.ToSafeSerializedString([]byte("hi")),
		"struct": converters.AnyTo.ToSafeSerializedString(struct{ N int }{42}) != "",
		"nil":    converters.AnyTo.ToSafeSerializedString(nil),
	}
	expected := args.Map{
		"bytes":  "hi",
		"struct": true,
		"nil":    "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToSafeSerializedString returns correct value -- with args", actual)
}

// ── AnyTo: ToSafeSerializedStringSprintValue ──

func Test_Cov3_AnyTo_ToSafeSerializedStringSprintValue(t *testing.T) {
	actual := args.Map{
		"notEmpty": converters.AnyTo.ToSafeSerializedStringSprintValue("hello") != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToSafeSerializedStringSprintValue returns correct value -- with args", actual)
}

// ── AnyTo: ToValueString ──

func Test_Cov3_AnyTo_ToValueString(t *testing.T) {
	actual := args.Map{
		"val": converters.AnyTo.ToValueString(42) != "",
		"nil": converters.AnyTo.ToValueString(nil),
	}
	expected := args.Map{
		"val": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToValueString returns non-empty -- with args", actual)
}

// ── AnyTo: ToValueStringWithType ──

func Test_Cov3_AnyTo_ToValueStringWithType(t *testing.T) {
	actual := args.Map{
		"val":    converters.AnyTo.ToValueStringWithType(42) != "",
		"nilVal": converters.AnyTo.ToValueStringWithType(nil) != "",
	}
	expected := args.Map{
		"val":    true,
		"nilVal": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToValueStringWithType returns non-empty -- with args", actual)
}

// ── AnyTo: ToAnyItems ──

func Test_Cov3_AnyTo_ToAnyItems(t *testing.T) {
	result := converters.AnyTo.ToAnyItems(true, []any{1, 2})
	nilResult := converters.AnyTo.ToAnyItems(true, nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToAnyItems returns correct value -- with args", actual)
}

// ── AnyTo: ItemsToStringsSkipOnNil ──

func Test_Cov3_AnyTo_ItemsToStringsSkipOnNil(t *testing.T) {
	result := converters.AnyTo.ItemsToStringsSkipOnNil("a", 42, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo.ItemsToStringsSkipOnNil returns nil -- with args", actual)
}

// ── AnyTo: ItemsJoin ──

func Test_Cov3_AnyTo_ItemsJoin(t *testing.T) {
	result := converters.AnyTo.ItemsJoin(", ", "a", "b")
	nilResult := converters.AnyTo.ItemsJoin(", ")
	actual := args.Map{"val": result, "nil": nilResult}
	expected := args.Map{"val": "a, b", "nil": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo.ItemsJoin returns correct value -- with args", actual)
}

// ── AnyTo: ToItemsThenJoin ──

func Test_Cov3_AnyTo_ToItemsThenJoin(t *testing.T) {
	result := converters.AnyTo.ToItemsThenJoin(true, ", ", []any{"a", "b"})
	nilResult := converters.AnyTo.ToItemsThenJoin(true, ", ", nil)
	actual := args.Map{"notEmpty": result != "", "nil": nilResult}
	expected := args.Map{"notEmpty": true, "nil": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToItemsThenJoin returns correct value -- with args", actual)
}

// ── AnyTo: ToFullNameValueString ──

func Test_Cov3_AnyTo_ToFullNameValueString(t *testing.T) {
	actual := args.Map{
		"val": converters.AnyTo.ToFullNameValueString(42) != "",
		"nil": converters.AnyTo.ToFullNameValueString(nil),
	}
	expected := args.Map{
		"val": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToFullNameValueString returns non-empty -- with args", actual)
}

// ── AnyTo: ToPrettyJson ──

func Test_Cov3_AnyTo_ToPrettyJson(t *testing.T) {
	actual := args.Map{
		"struct": converters.AnyTo.ToPrettyJson(struct{ N int }{42}) != "",
		"nil":    converters.AnyTo.ToPrettyJson(nil),
	}
	expected := args.Map{
		"struct": true,
		"nil":    "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToPrettyJson returns correct value -- with args", actual)
}

// ── AnyTo: Bytes ──

func Test_Cov3_AnyTo_Bytes(t *testing.T) {
	fromBytes := converters.AnyTo.Bytes([]byte("hi"))
	fromString := converters.AnyTo.Bytes("hello")
	fromNilBytes := converters.AnyTo.Bytes([]byte(nil))
	fromStruct := converters.AnyTo.Bytes(struct{ N int }{1})
	actual := args.Map{
		"fromBytesLen":    len(fromBytes),
		"fromStringLen":   len(fromString),
		"nilBytesLen":     len(fromNilBytes),
		"fromStructEmpty": len(fromStruct) > 0,
	}
	expected := args.Map{
		"fromBytesLen":    2,
		"fromStringLen":   5,
		"nilBytesLen":     0,
		"fromStructEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.Bytes returns correct value -- with args", actual)
}

// ── AnyTo: ValueString ──

func Test_Cov3_AnyTo_ValueString(t *testing.T) {
	actual := args.Map{
		"val": converters.AnyTo.ValueString(42) != "",
		"nil": converters.AnyTo.ValueString(nil),
	}
	expected := args.Map{
		"val": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.ValueString returns non-empty -- with args", actual)
}

// ── AnyTo: SmartString ──

func Test_Cov3_AnyTo_SmartString(t *testing.T) {
	actual := args.Map{
		"str": converters.AnyTo.SmartString("hello"),
		"int": converters.AnyTo.SmartString(42) != "",
		"nil": converters.AnyTo.SmartString(nil),
	}
	expected := args.Map{
		"str": "hello",
		"int": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartString returns correct value -- with args", actual)
}

// ── AnyTo: SmartStringsJoiner ──

func Test_Cov3_AnyTo_SmartStringsJoiner(t *testing.T) {
	result := converters.AnyTo.SmartStringsJoiner(", ", "a", 42)
	emptyResult := converters.AnyTo.SmartStringsJoiner(", ")
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartStringsJoiner returns correct value -- with args", actual)
}

// ── AnyTo: SmartStringsOf ──

func Test_Cov3_AnyTo_SmartStringsOf(t *testing.T) {
	result := converters.AnyTo.SmartStringsOf("a", 42)
	emptyResult := converters.AnyTo.SmartStringsOf()
	actual := args.Map{"notEmpty": result != "", "empty": emptyResult}
	expected := args.Map{"notEmpty": true, "empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo.SmartStringsOf returns correct value -- with args", actual)
}

// ── AnyTo: ToStrings ──

func Test_Cov3_AnyTo_ToStrings(t *testing.T) {
	result := converters.AnyTo.ToStrings(true, []any{"a", "b"})
	nilResult := converters.AnyTo.ToStrings(true, nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": len(result), "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToStrings returns correct value -- with args", actual)
}

// ── AnyTo: ToNonNullItems nil input ──

func Test_Cov3_AnyTo_ToNonNullItems_Nil(t *testing.T) {
	result := converters.AnyTo.ToNonNullItems(true, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo.ToNonNullItems returns nil -- nil", actual)
}

// ── StringsTo: Hashset ──

func Test_Cov3_StringsTo_Hashset(t *testing.T) {
	result := converters.StringsTo.Hashset([]string{"a", "b", "a"})
	actual := args.Map{"len": len(result), "hasA": result["a"], "hasB": result["b"]}
	expected := args.Map{"len": 2, "hasA": true, "hasB": true}
	expected.ShouldBeEqual(t, 0, "StringsTo.Hashset returns correct value -- with args", actual)
}

// ── StringsTo: PointerStrings ──

func Test_Cov3_StringsTo_PointerStrings(t *testing.T) {
	items := []string{"a", "b"}
	result := converters.StringsTo.PointerStrings(&items)
	nilResult := converters.StringsTo.PointerStrings(nil)
	actual := args.Map{"len": len(*result), "nilLen": len(*nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "StringsTo.PointerStrings returns correct value -- with args", actual)
}

// ── StringsTo: PointerStringsCopy ──

func Test_Cov3_StringsTo_PointerStringsCopy(t *testing.T) {
	items := []string{"a", "b"}
	result := converters.StringsTo.PointerStringsCopy(&items)
	nilResult := converters.StringsTo.PointerStringsCopy(nil)
	actual := args.Map{"len": len(*result), "nilLen": len(*nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "StringsTo.PointerStringsCopy returns correct value -- with args", actual)
}

// ── StringsTo: IntegersConditional ──

func Test_Cov3_StringsTo_IntegersConditional(t *testing.T) {
	result := converters.StringsTo.IntegersConditional(func(in string) (int, bool, bool) {
		return len(in), true, false
	}, "a", "bb")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsTo.IntegersConditional returns correct value -- with args", actual)
}

func Test_Cov3_StringsTo_IntegersConditional_Break(t *testing.T) {
	result := converters.StringsTo.IntegersConditional(func(in string) (int, bool, bool) {
		return 0, true, true
	}, "a", "b", "c")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringsTo.IntegersConditional returns correct value -- break", actual)
}

// ── StringsTo: IntegersWithDefaults ──

func Test_Cov3_StringsTo_IntegersWithDefaults(t *testing.T) {
	result := converters.StringsTo.IntegersWithDefaults(-1, "1", "abc", "3")
	actual := args.Map{"len": len(result.Values), "hasErr": result.CombinedError != nil}
	expected := args.Map{"len": 3, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "StringsTo.IntegersWithDefaults returns non-empty -- with args", actual)
}

// ── StringsTo: IntegersOptionPanic (no panic) ──

func Test_Cov3_StringsTo_IntegersOptionPanic_NoPanic(t *testing.T) {
	result := converters.StringsTo.IntegersOptionPanic(false, "1", "abc", "3")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "IntegersOptionPanic panics -- noPanic", actual)
}

// ── StringsTo: IntegersSkipErrors ──

func Test_Cov3_StringsTo_IntegersSkipErrors(t *testing.T) {
	result := converters.StringsTo.IntegersSkipErrors("1", "abc", "3")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "IntegersSkipErrors returns error -- with args", actual)
}

// ── StringsTo: BytesMust ──

func Test_Cov3_StringsTo_BytesMust(t *testing.T) {
	result := converters.StringsTo.BytesMust("0", "1", "255")
	actual := args.Map{"len": len(result), "first": int(result[0]), "last": int(result[2])}
	expected := args.Map{"len": 3, "first": 0, "last": 255}
	expected.ShouldBeEqual(t, 0, "BytesMust returns correct value -- with args", actual)
}

// ── BytesTo: PtrString ──

func Test_Cov3_BytesTo_PtrString(t *testing.T) {
	actual := args.Map{
		"val": converters.BytesTo.PtrString([]byte("hello")),
		"nil": converters.BytesTo.PtrString(nil),
	}
	expected := args.Map{
		"val": "hello",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.PtrString returns correct value -- with args", actual)
}

// ── BytesTo: PointerToBytes ──

func Test_Cov3_BytesTo_PointerToBytes(t *testing.T) {
	result := converters.BytesTo.PointerToBytes([]byte("hi"))
	nilResult := converters.BytesTo.PointerToBytes(nil)
	actual := args.Map{"len": len(result), "nilLen": len(nilResult)}
	expected := args.Map{"len": 2, "nilLen": 0}
	expected.ShouldBeEqual(t, 0, "BytesTo.PointerToBytes returns correct value -- with args", actual)
}

// ── UnsafeBytesTo (package-level functions) ──

func Test_Cov3_UnsafeBytesToStringWithErr(t *testing.T) {
	val, err := converters.UnsafeBytesToStringWithErr([]byte("hi"))
	_, nilErr := converters.UnsafeBytesToStringWithErr(nil)
	actual := args.Map{"val": val, "noErr": err == nil, "nilErr": nilErr != nil}
	expected := args.Map{"val": "hi", "noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesToStringWithErr returns error -- with args", actual)
}

func Test_Cov3_UnsafeBytesToString(t *testing.T) {
	actual := args.Map{
		"val": converters.UnsafeBytesToString([]byte("hi")),
		"nil": converters.UnsafeBytesToString(nil),
	}
	expected := args.Map{
		"val": "hi",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesToString returns correct value -- with args", actual)
}

func Test_Cov3_UnsafeBytesToStrings(t *testing.T) {
	result := converters.UnsafeBytesToStrings([]byte{65, 66})
	nilResult := converters.UnsafeBytesToStrings(nil)
	actual := args.Map{"len": len(result), "nilIsNil": nilResult == nil}
	expected := args.Map{"len": 2, "nilIsNil": true}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesToStrings returns correct value -- with args", actual)
}

func Test_Cov3_UnsafeBytesToStringPtr(t *testing.T) {
	result := converters.UnsafeBytesToStringPtr([]byte("hi"))
	nilResult := converters.UnsafeBytesToStringPtr(nil)
	actual := args.Map{"notNil": result != nil, "nilIsNil": nilResult == nil}
	expected := args.Map{"notNil": true, "nilIsNil": true}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesToStringPtr returns correct value -- with args", actual)
}

func Test_Cov3_UnsafeBytesPtrToStringPtr(t *testing.T) {
	result := converters.UnsafeBytesPtrToStringPtr([]byte("hi"))
	nilResult := converters.UnsafeBytesPtrToStringPtr(nil)
	actual := args.Map{"notNil": result != nil, "nilIsNil": nilResult == nil}
	expected := args.Map{"notNil": true, "nilIsNil": true}
	expected.ShouldBeEqual(t, 0, "UnsafeBytesPtrToStringPtr returns correct value -- with args", actual)
}

// ── StringsToMapConverter additional methods ──

func Test_Cov3_StringsToMapConverter_LineSplitMap(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a:1", "b:2"})
	result := mc.LineSplitMap(":")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsToMapConverter.LineSplitMap returns correct value -- with args", actual)
}

func Test_Cov3_StringsToMapConverter_LineProcessorMapStringIntegerTrim(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapStringIntegerTrim(func(line string) (string, int) {
		return "k", 1
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringIntegerTrim returns correct value -- with args", actual)
}

func Test_Cov3_StringsToMapConverter_LineProcessorMapStringAnyTrim(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapStringAnyTrim(func(line string) (string, any) {
		return "k", "v"
	})
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringAnyTrim returns correct value -- with args", actual)
}

func Test_Cov3_StringsToMapConverter_Strings(t *testing.T) {
	mc := converters.StringsToMapConverter([]string{"a", "b"})
	actual := args.Map{"len": len(mc.Strings())}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsToMapConverter.Strings returns correct value -- with args", actual)
}

// ── AnyTo: ToStringsUsingProcessor nil ──

func Test_Cov3_AnyTo_ToStringsUsingProcessor_Nil(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return fmt.Sprintf("%v", in), true, false
	}, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor returns nil -- nil", actual)
}

// ── AnyTo: ToStringsUsingSimpleProcessor nil ──

func Test_Cov3_AnyTo_ToStringsUsingSimpleProcessor_Nil(t *testing.T) {
	result := converters.AnyTo.ToStringsUsingSimpleProcessor(true, func(index int, in any) string {
		return "x"
	}, nil)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingSimpleProcessor returns nil -- nil", actual)
}
