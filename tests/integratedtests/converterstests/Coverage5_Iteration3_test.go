package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════
// StringsToMapConverter — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C5_STMC_SafeStrings_Empty(t *testing.T) {
	var c converters.StringsToMapConverter
	r := c.SafeStrings()
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_SafeStrings_NonEmpty(t *testing.T) {
	c := converters.StringsToMapConverter{"a", "b"}
	r := c.SafeStrings()
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_Strings(t *testing.T) {
	c := converters.StringsToMapConverter{"a"}
	r := c.Strings()
	actual := args.Map{"result": len(r) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_Length_Nil(t *testing.T) {
	var c *converters.StringsToMapConverter
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_Length(t *testing.T) {
	c := converters.StringsToMapConverter{"a", "b"}
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_IsEmpty(t *testing.T) {
	c := converters.StringsToMapConverter{}
	actual := args.Map{"result": c.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_HasAnyItem(t *testing.T) {
	c := converters.StringsToMapConverter{"x"}
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LastIndex(t *testing.T) {
	c := converters.StringsToMapConverter{"a", "b"}
	actual := args.Map{"result": c.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LineSplitMapOptions_Trim(t *testing.T) {
	c := converters.StringsToMapConverter{" k : v "}
	m := c.LineSplitMapOptions(true, ":")
	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_STMC_LineSplitMapOptions_NoTrim(t *testing.T) {
	c := converters.StringsToMapConverter{"k:v"}
	m := c.LineSplitMapOptions(false, ":")
	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_STMC_LineProcessorMapOptions(t *testing.T) {
	c := converters.StringsToMapConverter{"hello"}
	m := c.LineProcessorMapOptions(true, func(line string) (string, string) {
		return line, "val"
	})
	actual := args.Map{"result": m["hello"] != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LineProcessorMapStringIntegerTrim(t *testing.T) {
	c := converters.StringsToMapConverter{"hello"}
	m := c.LineProcessorMapStringIntegerTrim(func(line string) (string, int) {
		return line, 42
	})
	actual := args.Map{"result": m["hello"] != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LineProcessorMapStringIntegerOptions(t *testing.T) {
	c := converters.StringsToMapConverter{"hello"}
	m := c.LineProcessorMapStringIntegerOptions(false, func(line string) (string, int) {
		return line, 7
	})
	actual := args.Map{"result": m["hello"] != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LineProcessorMapStringAnyTrim(t *testing.T) {
	c := converters.StringsToMapConverter{"hello"}
	m := c.LineProcessorMapStringAnyTrim(func(line string) (string, any) {
		return line, true
	})
	actual := args.Map{"result": m["hello"] != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LineProcessorMapStringAnyOptions(t *testing.T) {
	c := converters.StringsToMapConverter{"hello"}
	m := c.LineProcessorMapStringAnyOptions(false, func(line string) (string, any) {
		return line, 99
	})
	actual := args.Map{"result": m["hello"] != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_STMC_LineSplitMapTrim(t *testing.T) {
	c := converters.StringsToMapConverter{" a : b "}
	m := c.LineSplitMapTrim(":")
	actual := args.Map{"result": m["a"] != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_STMC_LineSplitMap(t *testing.T) {
	c := converters.StringsToMapConverter{"a:b"}
	m := c.LineSplitMap(":")
	actual := args.Map{"result": m["a"] != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

// ═══════════════════════════════════════════════
// anyItemConverter — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C5_AIC_ToString_Nil(t *testing.T) {
	r := converters.AnyTo.ToString(false, nil)
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToString_WithFullName(t *testing.T) {
	r := converters.AnyTo.ToString(true, 42)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToString_WithoutFullName(t *testing.T) {
	r := converters.AnyTo.ToString(false, 42)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_String_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.String(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_String_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.String("hello") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_FullString_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.FullString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_FullString_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.FullString("hello") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_StringWithType_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.StringWithType(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_StringWithType_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.StringWithType("hello") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToSafeSerializedString_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToSafeSerializedString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToSafeSerializedString_Bytes(t *testing.T) {
	r := converters.AnyTo.ToSafeSerializedString([]byte("hello"))
	actual := args.Map{"result": r != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToSafeSerializedString_Struct(t *testing.T) {
	r := converters.AnyTo.ToSafeSerializedString(map[string]int{"a": 1})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToSafeSerializedStringSprintValue(t *testing.T) {
	r := converters.AnyTo.ToSafeSerializedStringSprintValue("hello")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToStrings(t *testing.T) {
	r := converters.AnyTo.ToStrings(true, []string{"a", "b"})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToStrings_Nil(t *testing.T) {
	r := converters.AnyTo.ToStrings(true, nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToStringsUsingProcessor_Nil(t *testing.T) {
	r := converters.AnyTo.ToStringsUsingProcessor(true,
		func(i int, in any) (string, bool, bool) { return "", false, false }, nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToStringsUsingProcessor_WithBreak(t *testing.T) {
	r := converters.AnyTo.ToStringsUsingProcessor(false,
		func(i int, in any) (string, bool, bool) {
			return "x", true, i >= 0
		}, []string{"a", "b"})
	actual := args.Map{"result": len(r) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C5_AIC_ToStringsUsingProcessor_NoTake(t *testing.T) {
	r := converters.AnyTo.ToStringsUsingProcessor(false,
		func(i int, in any) (string, bool, bool) {
			return "", false, false
		}, []string{"a"})
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToStringsUsingSimpleProcessor_Nil(t *testing.T) {
	r := converters.AnyTo.ToStringsUsingSimpleProcessor(true,
		func(i int, in any) string { return "" }, nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToStringsUsingSimpleProcessor_Valid(t *testing.T) {
	r := converters.AnyTo.ToStringsUsingSimpleProcessor(false,
		func(i int, in any) string { return "mapped" }, []string{"a"})
	actual := args.Map{"result": len(r) != 1 || r[0] != "mapped"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToValueString_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToValueString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToValueString_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToValueString(42) == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToValueStringWithType_Nil(t *testing.T) {
	r := converters.AnyTo.ToValueStringWithType(nil)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should produce formatted nil string", actual)
}

func Test_C5_AIC_ToValueStringWithType_Valid(t *testing.T) {
	r := converters.AnyTo.ToValueStringWithType(42)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToAnyItems(t *testing.T) {
	r := converters.AnyTo.ToAnyItems(true, nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	r2 := converters.AnyTo.ToAnyItems(false, []int{1, 2})
	actual := args.Map{"result": len(r2) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToNonNullItems_Nil(t *testing.T) {
	r := converters.AnyTo.ToNonNullItems(true, nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToNonNullItems_Valid(t *testing.T) {
	r := converters.AnyTo.ToNonNullItems(false, []int{1})
	actual := args.Map{"result": len(r) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ItemsToStringsSkipOnNil(t *testing.T) {
	r := converters.AnyTo.ItemsToStringsSkipOnNil("a", nil, "b")
	actual := args.Map{"result": len(r) < 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ItemsJoin_Nil(t *testing.T) {
	r := converters.AnyTo.ItemsJoin(", ")
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ItemsJoin_Valid(t *testing.T) {
	r := converters.AnyTo.ItemsJoin(", ", "a", "b")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToItemsThenJoin_Nil(t *testing.T) {
	r := converters.AnyTo.ToItemsThenJoin(true, ", ", nil)
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToItemsThenJoin_Valid(t *testing.T) {
	r := converters.AnyTo.ToItemsThenJoin(false, ", ", []string{"a", "b"})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToFullNameValueString_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToFullNameValueString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToFullNameValueString_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToFullNameValueString(42) == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToPrettyJson_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ToPrettyJson(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToPrettyJson_Valid(t *testing.T) {
	r := converters.AnyTo.ToPrettyJson(map[string]int{"a": 1})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ToPrettyJson_ErrorSwallowed(t *testing.T) {
	// channel can't be marshalled
	r := converters.AnyTo.ToPrettyJson(make(chan int))
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_Bytes_ByteSlice(t *testing.T) {
	r := converters.AnyTo.Bytes([]byte("hello"))
	actual := args.Map{"result": string(r) != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_Bytes_NilByteSlice(t *testing.T) {
	r := converters.AnyTo.Bytes([]byte(nil))
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_Bytes_String(t *testing.T) {
	r := converters.AnyTo.Bytes("test")
	actual := args.Map{"result": string(r) != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_Bytes_Other(t *testing.T) {
	r := converters.AnyTo.Bytes(42)
	actual := args.Map{"result": string(r) != "42"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ValueString_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ValueString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_ValueString_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.ValueString(42) == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_SmartString_Nil(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.SmartString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_SmartString_Valid(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.SmartString("hello") == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_SmartStringsJoiner_Empty(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.SmartStringsJoiner(", ") != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_SmartStringsJoiner_Valid(t *testing.T) {
	r := converters.AnyTo.SmartStringsJoiner(", ", "a", 1)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_SmartStringsOf_Empty(t *testing.T) {
	actual := args.Map{"result": converters.AnyTo.SmartStringsOf() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_AIC_SmartStringsOf_Valid(t *testing.T) {
	r := converters.AnyTo.SmartStringsOf("a", "b")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// bytesTo — all uncovered methods
// ═══════════════════════════════════════════════

func Test_C5_BytesTo_PtrString_Empty(t *testing.T) {
	actual := args.Map{"result": converters.BytesTo.PtrString(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_BytesTo_PtrString_Valid(t *testing.T) {
	actual := args.Map{"result": converters.BytesTo.PtrString([]byte("hello")) != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_BytesTo_String_Empty(t *testing.T) {
	actual := args.Map{"result": converters.BytesTo.String(nil) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_BytesTo_String_Valid(t *testing.T) {
	actual := args.Map{"result": converters.BytesTo.String([]byte("test")) != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_BytesTo_PointerToBytes_Nil(t *testing.T) {
	r := converters.BytesTo.PointerToBytes(nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_BytesTo_PointerToBytes_Valid(t *testing.T) {
	r := converters.BytesTo.PointerToBytes([]byte{1, 2})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// stringTo — uncovered methods
// ═══════════════════════════════════════════════

func Test_C5_StringTo_IntegerWithDefault_Empty(t *testing.T) {
	v, ok := converters.StringTo.IntegerWithDefault("", 99)
	actual := args.Map{"result": ok || v != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_IntegersWithDefaults_Empty(t *testing.T) {
	r := converters.StringTo.IntegersWithDefaults("", ",", 0)
	actual := args.Map{"result": r.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_IntegersWithDefaults_WithError(t *testing.T) {
	r := converters.StringTo.IntegersWithDefaults("1,abc,3", ",", 0)
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	actual := args.Map{"result": r.Values[1] != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_IntegersConditional_Empty(t *testing.T) {
	r := converters.StringTo.IntegersConditional("", ",",
		func(in string) (int, bool, bool) { return 0, true, false })
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_IntegersConditional_WithBreak(t *testing.T) {
	r := converters.StringTo.IntegersConditional("1,2,3", ",",
		func(in string) (int, bool, bool) {
			v := converters.StringTo.IntegerDefault(in)
			return v, true, v >= 2
		})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C5_StringTo_IntegerMust_Success(t *testing.T) {
	actual := args.Map{"result": converters.StringTo.IntegerMust("42") != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_IntegerMust_Panic(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	converters.StringTo.IntegerMust("abc")
}

func Test_C5_StringTo_IntegerDefault(t *testing.T) {
	actual := args.Map{"result": converters.StringTo.IntegerDefault("abc") != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_Integer_Error(t *testing.T) {
	_, err := converters.StringTo.Integer("abc")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_Float64Must_Panic(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	converters.StringTo.Float64Must("abc")
}

func Test_C5_StringTo_Float64Default_Fail(t *testing.T) {
	_, ok := converters.StringTo.Float64Default("abc", 1.5)
	actual := args.Map{"result": ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_Float64Conditional(t *testing.T) {
	v, ok := converters.StringTo.Float64Conditional("3.14", 0)
	actual := args.Map{"result": ok || v != 3.14}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_ByteWithDefault_Fail(t *testing.T) {
	_, ok := converters.StringTo.ByteWithDefault("abc", 0)
	actual := args.Map{"result": ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_BytesConditional_Break(t *testing.T) {
	r := converters.StringTo.BytesConditional("1,2,3", ",",
		func(in string) (byte, bool, bool) {
			return 0, true, true
		})
	actual := args.Map{"result": len(r) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_BytesConditional_Empty(t *testing.T) {
	r := converters.StringTo.BytesConditional("", ",",
		func(in string) (byte, bool, bool) { return 0, false, false })
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_Byte_Edge(t *testing.T) {
	_, err := converters.StringTo.Byte("-1")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
	_, err = converters.StringTo.Byte("256")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringTo_JsonBytes(t *testing.T) {
	b := converters.StringTo.JsonBytes("test")
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// stringsTo — uncovered methods
// ═══════════════════════════════════════════════

func Test_C5_StringsTo_Hashset(t *testing.T) {
	m := converters.StringsTo.Hashset([]string{"a", "b"})
	actual := args.Map{"result": m["a"] || !m["b"]}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_HashmapTrimColon(t *testing.T) {
	m := converters.StringsTo.HashmapTrimColon(" k : v ")
	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_StringsTo_HashmapTrimHyphen(t *testing.T) {
	m := converters.StringsTo.HashmapTrimHyphen(" k - v ")
	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_StringsTo_HashmapOptions(t *testing.T) {
	m := converters.StringsTo.HashmapOptions(true, "=", "k = v")
	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_StringsTo_HashmapTrim(t *testing.T) {
	m := converters.StringsTo.HashmapTrim(":", []string{" k : v "})
	actual := args.Map{"result": m["k"] != "v"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_StringsTo_HashmapUsingFuncOptions(t *testing.T) {
	m := converters.StringsTo.HashmapUsingFuncOptions(true,
		func(line string) (string, string) { return line, "val" },
		"hello",
	)
	actual := args.Map{"result": m["hello"] != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_HashmapUsingFuncTrim(t *testing.T) {
	m := converters.StringsTo.HashmapUsingFuncTrim(
		func(line string) (string, string) { return line, "val" },
		"hello",
	)
	actual := args.Map{"result": m["hello"] != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_MapStringIntegerUsingFunc(t *testing.T) {
	m := converters.StringsTo.MapStringIntegerUsingFunc(true,
		func(line string) (string, int) { return line, 1 },
		"hello",
	)
	actual := args.Map{"result": m["hello"] != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_MapStringAnyUsingFunc(t *testing.T) {
	m := converters.StringsTo.MapStringAnyUsingFunc(true,
		func(line string) (string, any) { return line, true },
		"hello",
	)
	actual := args.Map{"result": m["hello"] != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_MapConverter(t *testing.T) {
	mc := converters.StringsTo.MapConverter("a:b")
	actual := args.Map{"result": mc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PointerStrings_Nil(t *testing.T) {
	r := converters.StringsTo.PointerStrings(nil)
	actual := args.Map{"result": r == nil || len(*r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PointerStrings_Valid(t *testing.T) {
	sl := []string{"a", "b"}
	r := converters.StringsTo.PointerStrings(&sl)
	actual := args.Map{"result": len(*r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PointerStringsCopy_Nil(t *testing.T) {
	r := converters.StringsTo.PointerStringsCopy(nil)
	actual := args.Map{"result": r == nil || len(*r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PointerStringsCopy_Valid(t *testing.T) {
	sl := []string{"a", "b"}
	r := converters.StringsTo.PointerStringsCopy(&sl)
	actual := args.Map{"result": len(*r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_IntegersConditional(t *testing.T) {
	r := converters.StringsTo.IntegersConditional(
		func(in string) (int, bool, bool) { return 1, true, false },
		"a", "b",
	)
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_IntegersWithDefaults_WithError(t *testing.T) {
	r := converters.StringsTo.IntegersWithDefaults(0, "abc", "2")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_IntegersOptionPanic_NoPanic(t *testing.T) {
	r := converters.StringsTo.IntegersOptionPanic(false, "1", "abc")
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_IntegersSkipErrors(t *testing.T) {
	r := converters.StringsTo.IntegersSkipErrors("1", "abc", "3")
	actual := args.Map{"result": len(r) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_IntegersSkipMapAndDefaultValue(t *testing.T) {
	skip := map[string]bool{"skip": true}
	r := converters.StringsTo.IntegersSkipMapAndDefaultValue(0, skip, "1", "skip", "abc")
	actual := args.Map{"result": r[0] != 1 || r[1] != 0 || r[2] != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_StringsTo_IntegersSkipAndDefaultValue(t *testing.T) {
	r := converters.StringsTo.IntegersSkipAndDefaultValue(0, "skip", "1", "skip", "abc")
	actual := args.Map{"result": r[0] != 1 || r[1] != 0 || r[2] != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "got", actual)
}

func Test_C5_StringsTo_BytesConditional(t *testing.T) {
	r := converters.StringsTo.BytesConditional(
		func(in string) (byte, bool, bool) { return 1, true, false },
		[]string{"a", "b"},
	)
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_BytesWithDefaults_Valid(t *testing.T) {
	r := converters.StringsTo.BytesWithDefaults(0, "1", "2")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_BytesWithDefaults_ParseError(t *testing.T) {
	r := converters.StringsTo.BytesWithDefaults(0, "abc")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_BytesWithDefaults_OutOfRange(t *testing.T) {
	r := converters.StringsTo.BytesWithDefaults(0, "256", "-1")
	actual := args.Map{"result": r.HasError()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_Csv(t *testing.T) {
	r := converters.StringsTo.Csv(false, "a", "b")
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_CsvUsingPtrStrings_Nil(t *testing.T) {
	r := converters.StringsTo.CsvUsingPtrStrings(false, nil)
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_CsvUsingPtrStrings_Valid(t *testing.T) {
	sl := []string{"a", "b"}
	r := converters.StringsTo.CsvUsingPtrStrings(false, &sl)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_CsvWithIndexes(t *testing.T) {
	r := converters.StringsTo.CsvWithIndexes([]string{"a", "b"})
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_BytesMust(t *testing.T) {
	r := converters.StringsTo.BytesMust("1", "2")
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_Float64sMust(t *testing.T) {
	r := converters.StringsTo.Float64sMust("1.5", "2.5")
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_Float64sConditional(t *testing.T) {
	r := converters.StringsTo.Float64sConditional(
		func(in string) (float64, bool, bool) { return 1.0, true, false },
		[]string{"a", "b"},
	)
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PtrOfPtrToPtrStrings_Nil(t *testing.T) {
	r := converters.StringsTo.PtrOfPtrToPtrStrings(nil)
	actual := args.Map{"result": r == nil || len(*r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PtrOfPtrToPtrStrings_WithNil(t *testing.T) {
	s := "hello"
	sl := []*string{&s, nil}
	r := converters.StringsTo.PtrOfPtrToPtrStrings(&sl)
	actual := args.Map{"result": len(*r) != 2 || (*r)[0] != "hello" || (*r)[1] != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PtrOfPtrToMapStringBool_Nil(t *testing.T) {
	r := converters.StringsTo.PtrOfPtrToMapStringBool(nil)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_PtrOfPtrToMapStringBool_WithNil(t *testing.T) {
	s := "hello"
	sl := []*string{&s, nil}
	r := converters.StringsTo.PtrOfPtrToMapStringBool(&sl)
	actual := args.Map{"result": r["hello"]}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_CloneIf_Clone(t *testing.T) {
	r := converters.StringsTo.CloneIf(true, "a", "b")
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_CloneIf_NoClone(t *testing.T) {
	r := converters.StringsTo.CloneIf(false, "a", "b")
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_StringsTo_CloneIf_Empty(t *testing.T) {
	r := converters.StringsTo.CloneIf(true)
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

// ═══════════════════════════════════════════════
// unsafeBytesTo — all uncovered functions
// ═══════════════════════════════════════════════

func Test_C5_UnsafeBytesToStringWithErr_Nil(t *testing.T) {
	_, err := converters.UnsafeBytesToStringWithErr(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToStringWithErr_Valid(t *testing.T) {
	s, err := converters.UnsafeBytesToStringWithErr([]byte("hello"))
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToStrings_Nil(t *testing.T) {
	r := converters.UnsafeBytesToStrings(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToStrings_Valid(t *testing.T) {
	r := converters.UnsafeBytesToStrings([]byte{65, 66})
	actual := args.Map{"result": len(r) != 2 || r[0] != "A"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToStringPtr_Nil(t *testing.T) {
	r := converters.UnsafeBytesToStringPtr(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToStringPtr_Valid(t *testing.T) {
	r := converters.UnsafeBytesToStringPtr([]byte("test"))
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToString_Nil(t *testing.T) {
	r := converters.UnsafeBytesToString(nil)
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesToString_Valid(t *testing.T) {
	r := converters.UnsafeBytesToString([]byte("hello"))
	actual := args.Map{"result": r != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesPtrToStringPtr_Nil(t *testing.T) {
	r := converters.UnsafeBytesPtrToStringPtr(nil)
	actual := args.Map{"result": r != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}

func Test_C5_UnsafeBytesPtrToStringPtr_Valid(t *testing.T) {
	r := converters.UnsafeBytesPtrToStringPtr([]byte("test"))
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "assertion", actual)
}
