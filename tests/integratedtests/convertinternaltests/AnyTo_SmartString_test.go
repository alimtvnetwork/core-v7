package convertinternaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

// ── AnyTo extra branches ──

func Test_AnyTo_SmartString_Stringer_FromAnyToSmartString(t *testing.T) {
	// Arrange
	type myStringer struct{ val string }
	// Use error interface as a proxy for Stringer-like
	err := errors.New("test-err")
	result := convertinternal.AnyTo.SmartString(err)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "test-err"}
	expected.ShouldBeEqual(t, 0, "SmartString_Error returns error -- with args", actual)
}

func Test_AnyTo_SmartString_StringSlice_FromAnyToSmartString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartString([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SmartString_StringSlice returns correct value -- with args", actual)
}

func Test_AnyTo_SmartString_AnySlice_FromAnyToSmartString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartString([]any{"a", 1})
	emptySlice := convertinternal.AnyTo.SmartString([]any{})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"empty": emptySlice,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "SmartString_AnySlice returns correct value -- with args", actual)
}

func Test_AnyTo_SmartJson_Error_FromAnyToSmartString(t *testing.T) {
	// Arrange
	err := errors.New("test")
	result := convertinternal.AnyTo.SmartJson(err)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "SmartJson_Error returns error -- with args", actual)
}

func Test_AnyTo_SmartJson_NilError(t *testing.T) {
	// Arrange
	var err error
	result := convertinternal.AnyTo.SmartJson(err)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "SmartJson_NilError returns nil -- with args", actual)
}

func Test_AnyTo_SmartPrettyJsonLines_FromAnyToSmartString(t *testing.T) {
	// Arrange
	strResult := convertinternal.AnyTo.SmartPrettyJsonLines("hello\nworld")
	sliceResult := convertinternal.AnyTo.SmartPrettyJsonLines([]string{"a", "b"})
	nilResult := convertinternal.AnyTo.SmartPrettyJsonLines(nil)

	// Act
	actual := args.Map{
		"strLen":   len(strResult),
		"sliceLen": len(sliceResult),
		"nilLen":   len(nilResult),
	}

	// Assert
	expected := args.Map{
		"strLen":   2,
		"sliceLen": 2,
		"nilLen":   0,
	}
	expected.ShouldBeEqual(t, 0, "SmartPrettyJsonLines returns correct value -- with args", actual)
}

func Test_AnyTo_Strings_MapTypes(t *testing.T) {
	// Arrange
	mapAny := convertinternal.AnyTo.Strings(map[string]any{"a": 1})
	mapAnyAny := convertinternal.AnyTo.Strings(map[any]any{"a": 1})
	mapStrInt := convertinternal.AnyTo.Strings(map[string]int{"a": 1})
	mapIntStr := convertinternal.AnyTo.Strings(map[int]string{1: "a"})

	// Act
	actual := args.Map{
		"mapAnyLen":    len(mapAny),
		"mapAnyAnyLen": len(mapAnyAny),
		"mapStrIntLen": len(mapStrInt),
		"mapIntStrLen": len(mapIntStr),
	}

	// Assert
	expected := args.Map{
		"mapAnyLen":    1,
		"mapAnyAnyLen": 1,
		"mapStrIntLen": 1,
		"mapIntStrLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "Strings_MapTypes returns correct value -- with args", actual)
}

func Test_AnyTo_Strings_SliceTypes(t *testing.T) {
	// Arrange
	int64s := convertinternal.AnyTo.Strings([]int64{1, 2})
	float64s := convertinternal.AnyTo.Strings([]float64{1.1, 2.2})
	bytes := convertinternal.AnyTo.Strings([]byte{1, 2})

	// Act
	actual := args.Map{
		"int64sLen":   len(int64s),
		"float64sLen": len(float64s),
		"bytesLen":    len(bytes),
	}

	// Assert
	expected := args.Map{
		"int64sLen":   2,
		"float64sLen": 2,
		"bytesLen":    2,
	}
	expected.ShouldBeEqual(t, 0, "Strings_SliceTypes returns correct value -- with args", actual)
}

func Test_AnyTo_String_NilPtr_FromAnyToSmartString(t *testing.T) {
	// Arrange
	var s *string
	result := convertinternal.AnyTo.String(s)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "String_NilPtr returns nil -- with args", actual)
}

func Test_AnyTo_String_Error_FromAnyToSmartString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.String(errors.New("e"))
	var nilErr error
	nilResult := convertinternal.AnyTo.String(nilErr)

	// Act
	actual := args.Map{
		"val": result,
		"nil": nilResult,
	}

	// Assert
	expected := args.Map{
		"val": "e",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "String_Error returns error -- with args", actual)
}

func Test_AnyTo_FullPropertyString_FromAnyToSmartString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.FullPropertyString(42)
	nilResult := convertinternal.AnyTo.FullPropertyString(nil)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"nil": nilResult,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "FullPropertyString returns correct value -- with args", actual)
}

func Test_AnyTo_TypeName_FromAnyToSmartString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.TypeName(42)
	nilResult := convertinternal.AnyTo.TypeName(nil)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"nil": nilResult,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "TypeName returns correct value -- with args", actual)
}

// ── Map extra coverage ──

func Test_Map_KeysValues(t *testing.T) {
	// Arrange
	k, v, err := convertinternal.Map.KeysValues(map[string]string{"a": "1"})

	// Act
	actual := args.Map{
		"kLen": len(k),
		"vLen": len(v),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"kLen": 1,
		"vLen": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeysValues returns non-empty -- with args", actual)
}

func Test_Map_KeysValues_BadType(t *testing.T) {
	// Arrange
	_, _, err := convertinternal.Map.KeysValues(42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeysValues_BadType returns non-empty -- with args", actual)
}

func Test_Map_SortedKeys(t *testing.T) {
	// Arrange
	result, err := convertinternal.Map.SortedKeys(map[string]string{"b": "2", "a": "1"})

	// Act
	actual := args.Map{
		"first": result[0],
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SortedKeys returns correct value -- with args", actual)
}

func Test_Map_SortedKeysValues(t *testing.T) {
	// Arrange
	k, v, err := convertinternal.Map.SortedKeysValues(map[string]string{"b": "2", "a": "1"})

	// Act
	actual := args.Map{
		"firstK": k[0],
		"firstV": v[0],
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"firstK": "a",
		"firstV": "1",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SortedKeysValues returns non-empty -- with args", actual)
}

func Test_Map_Values_MoreTypes(t *testing.T) {
	// Arrange
	v1, _ := convertinternal.Map.Values(map[string]any{"a": 1})
	v2, _ := convertinternal.Map.Values(map[int]any{1: "a"})
	v3, _ := convertinternal.Map.Values(map[string]int{"a": 1})
	v4, _ := convertinternal.Map.Values(map[int]string{1: "a"})
	v5, _ := convertinternal.Map.Values(map[float64]any{1.0: "a"})
	v6, _ := convertinternal.Map.Values(map[any]any{"a": 1})
	v7, _ := convertinternal.Map.Values(map[any]string{"a": "b"})

	// Act
	actual := args.Map{
		"v1": len(v1), "v2": len(v2), "v3": len(v3),
		"v4": len(v4), "v5": len(v5), "v6": len(v6), "v7": len(v7),
	}

	// Assert
	expected := args.Map{
		"v1": 1, "v2": 1, "v3": 1,
		"v4": 1, "v5": 1, "v6": 1, "v7": 1,
	}
	expected.ShouldBeEqual(t, 0, "Values_MoreTypes returns non-empty -- with args", actual)
}

func Test_Map_CombineMapStringAny(t *testing.T) {
	// Arrange
	result := convertinternal.Map.CombineMapStringAny(false, map[string]any{"a": 1}, map[string]any{"b": 2})
	emptyResult := convertinternal.Map.CombineMapStringAny(false, nil, nil)
	skipResult := convertinternal.Map.CombineMapStringAny(true, map[string]any{"a": ""}, map[string]any{"b": ""})

	// Act
	actual := args.Map{
		"len":      len(result),
		"emptyLen": len(emptyResult),
		"skipLen":  len(skipResult),
	}

	// Assert
	expected := args.Map{
		"len":      2,
		"emptyLen": 0,
		"skipLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "CombineMapStringAny returns correct value -- with args", actual)
}

func Test_Map_StringAnyToStringString_Skip(t *testing.T) {
	// Arrange
	result := convertinternal.Map.StringAnyToStringString(true, map[string]any{"a": ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringAnyToStringString_Skip returns correct value -- with args", actual)
}

// ── Map.Keys extra types ──

func Test_Map_Keys_MoreTypes(t *testing.T) {
	// Arrange
	k1, _ := convertinternal.Map.Keys(map[string]any{"a": 1})
	k2, _ := convertinternal.Map.Keys(map[int]any{1: "a"})
	k3, _ := convertinternal.Map.Keys(map[float64]any{1.0: "a"})
	k4, _ := convertinternal.Map.Keys(map[any]any{"a": 1})
	k5, _ := convertinternal.Map.Keys(map[any]string{"a": "b"})

	// Act
	actual := args.Map{
		"k1": len(k1), "k2": len(k2), "k3": len(k3), "k4": len(k4), "k5": len(k5),
	}

	// Assert
	expected := args.Map{
		"k1": 1, "k2": 1, "k3": 1, "k4": 1, "k5": 1,
	}
	expected.ShouldBeEqual(t, 0, "Keys_MoreTypes returns correct value -- with args", actual)
}

// ── KeyValuesTo ──

func Test_KeyValuesTo_ToMap(t *testing.T) {
	// Arrange
	result := convertinternal.KeyValuesTo.ToMap([]string{"a", "b"}, []string{"1", "2"})
	nilResult := convertinternal.KeyValuesTo.ToMap(nil, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "KeyValuesTo_ToMap returns non-empty -- with args", actual)
}

func Test_KeyValuesTo_ToMapPtr(t *testing.T) {
	// Arrange
	keys := []string{"a"}
	vals := []string{"1"}
	result := convertinternal.KeyValuesTo.ToMapPtr(&keys, &vals)
	nilResult := convertinternal.KeyValuesTo.ToMapPtr(nil, nil)

	// Act
	actual := args.Map{
		"len": len(*result),
		"nilIsEmpty": *nilResult == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"nilIsEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValuesTo_ToMapPtr returns non-empty -- with args", actual)
}

// ── CodeFormatter ──

func Test_CodeFormatter_GolangRaw_FromAnyToSmartString(t *testing.T) {
	// Arrange
	emptyResult, err := convertinternal.CodeFormatter.GolangRaw([]byte{})

	// Act
	actual := args.Map{
		"emptyLen": len(emptyResult),
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"emptyLen": 0,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GolangRaw_Empty returns empty -- with args", actual)
}

func Test_CodeFormatter_Golang_FromAnyToSmartString(t *testing.T) {
	// Arrange
	emptyResult, err := convertinternal.CodeFormatter.Golang("")

	// Act
	actual := args.Map{
		"empty": emptyResult,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"empty": "",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Golang_Empty returns empty -- with args", actual)
}

func Test_CodeFormatter_Golang_Invalid_FromAnyToSmartString(t *testing.T) {
	// Arrange
	_, err := convertinternal.CodeFormatter.Golang("invalid go code {{{}}")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Golang_Invalid returns error -- with args", actual)
}

// ── Util.Strings.PrependWithSpaces ──

func Test_UtilStrings_PrependWithSpaces(t *testing.T) {
	result := convertinternal.Util.Strings.PrependWithSpaces(2, []string{"a"}, 4, "header")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "PrependWithSpaces returns non-empty -- with args", actual)
}

func Test_UtilString_PrependWithSpaces(t *testing.T) {
	result := convertinternal.Util.String.PrependWithSpaces(", ", 2, []string{"a"}, 0, "header")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringPrependWithSpaces returns non-empty -- with args", actual)
}

func Test_UtilString_PrependWithSpacesDefault(t *testing.T) {
	result := convertinternal.Util.String.PrependWithSpacesDefault(0, []string{"a"}, 0, "header")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependWithSpacesDefault returns non-empty -- with args", actual)
}

// ── Integers.FromIntegersToMap ──

func Test_Integers_FromIntegersToMap_FromAnyToSmartString(t *testing.T) {
	result := convertinternal.Integers.FromIntegersToMap(1, 2)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FromIntegersToMap returns correct value -- with args", actual)
}

// ── Int8ToMapBool empty ──

func Test_Integers_Int8ToMapBool_Empty(t *testing.T) {
	result := convertinternal.Integers.Int8ToMapBool()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Int8ToMapBool_Empty returns empty -- with args", actual)
}
