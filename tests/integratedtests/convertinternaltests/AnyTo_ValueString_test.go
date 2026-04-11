package convertinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

// ── AnyTo ──

func Test_AnyTo_ValueString(t *testing.T) {
	// Act
	actual := args.Map{
		"int":    convertinternal.AnyTo.ValueString(42),
		"nil":    convertinternal.AnyTo.ValueString(nil),
		"string": convertinternal.AnyTo.ValueString("hello"),
	}
	expected := args.Map{
		"int":    "42",
		"nil":    "",
		"string": "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_ValueString returns non-empty -- with args", actual)
}

func Test_AnyTo_SmartString(t *testing.T) {
	// Act
	actual := args.Map{
		"string": convertinternal.AnyTo.SmartString("hello"),
		"nil":    convertinternal.AnyTo.SmartString(nil),
		"int":    convertinternal.AnyTo.SmartString(42),
	}
	expected := args.Map{
		"string": "hello",
		"nil":    "",
		"int":    "42",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_SmartString returns correct value -- with args", actual)
}

func Test_AnyTo_SmartJson(t *testing.T) {
	// Act
	actual := args.Map{
		"nil":    convertinternal.AnyTo.SmartJson(nil),
		"string": convertinternal.AnyTo.SmartJson("hello"),
		"int":    convertinternal.AnyTo.SmartJson(42),
	}
	expected := args.Map{
		"nil":    "",
		"string": "hello",
		"int":    "42",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_SmartJson returns correct value -- with args", actual)
}

func Test_AnyTo_String(t *testing.T) {
	// Act
	str := "hello"
	actual := args.Map{
		"string":    convertinternal.AnyTo.String("hello"),
		"ptrString": convertinternal.AnyTo.String(&str),
		"nil":       convertinternal.AnyTo.String(nil),
		"int":       convertinternal.AnyTo.String(42),
		"bool":      convertinternal.AnyTo.String(true),
	}
	expected := args.Map{
		"string":    "hello",
		"ptrString": "hello",
		"nil":       "",
		"int":       "42",
		"bool":      "true",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_String returns correct value -- with args", actual)
}

func Test_AnyTo_Strings(t *testing.T) {
	// Act
	strSlice := convertinternal.AnyTo.Strings([]string{"a", "b"})
	intSlice := convertinternal.AnyTo.Strings([]int{1, 2, 3})
	boolSlice := convertinternal.AnyTo.Strings([]bool{true, false})
	nilResult := convertinternal.AnyTo.Strings(nil)
	mapResult := convertinternal.AnyTo.Strings(map[string]string{"a": "1"})

	actual := args.Map{
		"strLen":  len(strSlice),
		"intLen":  len(intSlice),
		"boolLen": len(boolSlice),
		"nilLen":  len(nilResult),
		"mapLen":  len(mapResult),
	}
	expected := args.Map{
		"strLen":  2,
		"intLen":  3,
		"boolLen": 2,
		"nilLen":  0,
		"mapLen":  1,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_Strings returns correct value -- with args", actual)
}

func Test_AnyTo_PrettyJsonLines_FromAnyToValueString(t *testing.T) {
	// Act
	result := convertinternal.AnyTo.PrettyJsonLines(map[string]string{"a": "1"})
	nilResult := convertinternal.AnyTo.PrettyJsonLines(nil)

	actual := args.Map{
		"resultLen": len(result) > 0,
		"nilLen":    len(nilResult),
	}
	expected := args.Map{
		"resultLen": true,
		"nilLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo_PrettyJsonLines returns correct value -- with args", actual)
}

// ── Integers ──

func Test_Integers_ToMapBool_FromAnyToValueString(t *testing.T) {
	// Act
	result := convertinternal.Integers.ToMapBool(1, 2, 3)
	empty := convertinternal.Integers.ToMapBool()
	int8Map := convertinternal.Integers.Int8ToMapBool(1, 2)

	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(empty),
		"int8Len":   len(int8Map),
	}
	expected := args.Map{
		"resultLen": 3,
		"emptyLen":  0,
		"int8Len":   2,
	}
	expected.ShouldBeEqual(t, 0, "Integers_ToMapBool returns correct value -- with args", actual)
}

func Test_Integers_IntegersToStrings_FromAnyToValueString(t *testing.T) {
	// Act
	result := convertinternal.Integers.IntegersToStrings([]int{1, 2, 3})

	actual := args.Map{
		"len":   len(result),
		"first": result[0],
	}
	expected := args.Map{
		"len":   3,
		"first": "1",
	}
	expected.ShouldBeEqual(t, 0, "Integers_IntegersToStrings returns correct value -- with args", actual)
}

// ── StringsTo ──

func Test_StringsTo_WithSpaces(t *testing.T) {
	// Act
	result := convertinternal.StringsTo.WithSpaces(2, "a", "b")
	empty := convertinternal.StringsTo.WithSpaces(2)

	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(empty),
	}
	expected := args.Map{
		"resultLen": 2,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "StringsTo_WithSpaces returns non-empty -- with args", actual)
}

// ── Map ──

func Test_Map_Keys(t *testing.T) {
	// Act
	keys, err := convertinternal.Map.Keys(map[string]string{"a": "1", "b": "2"})
	intKeys, err2 := convertinternal.Map.Keys(map[int]string{1: "a"})
	_, errBad := convertinternal.Map.Keys(42)

	actual := args.Map{
		"keysLen":    len(keys),
		"noErr":      err == nil,
		"intKeysLen": len(intKeys),
		"noErr2":     err2 == nil,
		"hasErrBad":  errBad != nil,
	}
	expected := args.Map{
		"keysLen":    2,
		"noErr":      true,
		"intKeysLen": 1,
		"noErr2":     true,
		"hasErrBad":  true,
	}
	expected.ShouldBeEqual(t, 0, "Map_Keys returns correct value -- with args", actual)
}

func Test_Map_Values(t *testing.T) {
	// Act
	vals, err := convertinternal.Map.Values(map[string]string{"a": "1"})
	_, errBad := convertinternal.Map.Values(42)

	actual := args.Map{
		"valsLen":   len(vals),
		"noErr":     err == nil,
		"hasErrBad": errBad != nil,
	}
	expected := args.Map{
		"valsLen":   1,
		"noErr":     true,
		"hasErrBad": true,
	}
	expected.ShouldBeEqual(t, 0, "Map_Values returns non-empty -- with args", actual)
}

func Test_Map_StringAnyToStringString(t *testing.T) {
	// Act
	result := convertinternal.Map.StringAnyToStringString(false, map[string]any{"a": 1})
	empty := convertinternal.Map.StringAnyToStringString(false, map[string]any{})

	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(empty),
	}
	expected := args.Map{
		"resultLen": 1,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "Map_StringAnyToStringString returns correct value -- with args", actual)
}

func Test_Map_FromIntegersToMap(t *testing.T) {
	// Act
	result := convertinternal.Map.FromIntegersToMap(1, 2, 3)
	empty := convertinternal.Map.FromIntegersToMap()

	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(empty),
	}
	expected := args.Map{
		"resultLen": 3,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "Map_FromIntegersToMap returns correct value -- with args", actual)
}

// ── CodeFormatter / Util ──

func Test_Util_StringIndexToPosition(t *testing.T) {
	// Act
	actual := args.Map{
		"0": convertinternal.Util.String.IndexToPosition(0),
		"1": convertinternal.Util.String.IndexToPosition(1),
		"2": convertinternal.Util.String.IndexToPosition(2),
		"3": convertinternal.Util.String.IndexToPosition(3),
	}
	expected := args.Map{
		"0": "1st",
		"1": "2nd",
		"2": "3rd",
		"3": "4th",
	}
	expected.ShouldBeEqual(t, 0, "Util_StringIndexToPosition returns correct value -- with args", actual)
}

func Test_Util_StringPascalCase(t *testing.T) {
	// Act
	actual := args.Map{
		"simple":     convertinternal.Util.String.PascalCase("hello"),
		"underscore": convertinternal.Util.String.PascalCase("hello_world"),
		"empty":      convertinternal.Util.String.PascalCase(""),
	}
	expected := args.Map{
		"simple":     "Hello",
		"underscore": "HelloWorld",
		"empty":      "",
	}
	expected.ShouldBeEqual(t, 0, "Util_StringPascalCase returns correct value -- with args", actual)
}

func Test_Util_StringCamelCase(t *testing.T) {
	// Act
	actual := args.Map{
		"simple":     convertinternal.Util.String.CamelCase("Hello"),
		"underscore": convertinternal.Util.String.CamelCase("Hello_World"),
		"empty":      convertinternal.Util.String.CamelCase(""),
	}
	expected := args.Map{
		"simple":     "hello",
		"underscore": "helloWorld",
		"empty":      "",
	}
	expected.ShouldBeEqual(t, 0, "Util_StringCamelCase returns correct value -- with args", actual)
}
