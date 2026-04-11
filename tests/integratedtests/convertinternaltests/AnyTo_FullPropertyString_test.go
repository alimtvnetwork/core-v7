package convertinternaltests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/convertinternal"
)

// ── AnyTo extended ──

func Test_AnyTo_FullPropertyString(t *testing.T) {
	// Act
	actual := args.Map{
		"int":    convertinternal.AnyTo.FullPropertyString(42) != "",
		"nil":    convertinternal.AnyTo.FullPropertyString(nil),
		"string": convertinternal.AnyTo.FullPropertyString("hi") != "",
	}

	// Assert
	expected := args.Map{
		"int": true,
		"nil": "",
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- FullPropertyString", actual)
}

func Test_AnyTo_TypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"int":    convertinternal.AnyTo.TypeName(42) != "",
		"nil":    convertinternal.AnyTo.TypeName(nil),
		"string": convertinternal.AnyTo.TypeName("hi") != "",
	}

	// Assert
	expected := args.Map{
		"int": true,
		"nil": "",
		"string": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- TypeName", actual)
}

func Test_AnyTo_SmartString_Error(t *testing.T) {
	// Arrange
	err := errors.New("test error")
	var nilErr error

	// Act
	actual := args.Map{
		"error":   convertinternal.AnyTo.SmartString(err),
		"nilErr":  convertinternal.AnyTo.SmartString(nilErr),
	}

	// Assert
	expected := args.Map{
		"error": "test error",
		"nilErr": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns error -- SmartString error", actual)
}

func Test_AnyTo_SmartString_Stringer(t *testing.T) {
	// Arrange
	type myStringer struct{ val string }
	// Use a type that implements fmt.Stringer
	result := convertinternal.AnyTo.SmartString(fmt.Errorf("stringer"))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartString stringer", actual)
}

func Test_AnyTo_SmartString_StringSlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartString([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartString string slice", actual)
}

func Test_AnyTo_SmartString_AnySlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartString([]any{"a", 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartString any slice", actual)
}

func Test_AnyTo_SmartString_EmptyAnySlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartString([]any{})

	// Act
	actual := args.Map{"empty": result}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo returns empty -- SmartString empty any slice", actual)
}

func Test_AnyTo_SmartJson_Error(t *testing.T) {
	// Arrange
	err := errors.New("json err")
	var nilErr error

	// Act
	actual := args.Map{
		"error":  convertinternal.AnyTo.SmartJson(err),
		"nilErr": convertinternal.AnyTo.SmartJson(nilErr),
	}

	// Assert
	expected := args.Map{
		"error": "json err",
		"nilErr": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns error -- SmartJson error", actual)
}

func Test_AnyTo_SmartJson_StringSlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartJson([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartJson string slice", actual)
}

func Test_AnyTo_SmartJson_Bool(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.SmartJson(true)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "true"}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartJson bool", actual)
}

func Test_AnyTo_SmartJson_Struct(t *testing.T) {
	// Arrange
	type s struct{ A int }
	result := convertinternal.AnyTo.SmartJson(s{A: 1})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartJson struct", actual)
}

func Test_AnyTo_SmartPrettyJsonLines(t *testing.T) {
	// Arrange
	strResult := convertinternal.AnyTo.SmartPrettyJsonLines("a\nb")
	sliceResult := convertinternal.AnyTo.SmartPrettyJsonLines([]string{"a", "b"})
	nilResult := convertinternal.AnyTo.SmartPrettyJsonLines(nil)
	structResult := convertinternal.AnyTo.SmartPrettyJsonLines(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"strLen":    len(strResult),
		"sliceLen":  len(sliceResult),
		"nilLen":    len(nilResult),
		"structGt0": len(structResult) > 0,
	}

	// Assert
	expected := args.Map{
		"strLen": 2,
		"sliceLen": 2,
		"nilLen": 0,
		"structGt0": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SmartPrettyJsonLines", actual)
}

func Test_AnyTo_PrettyJsonLines(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.PrettyJsonLines(map[string]int{"a": 1})
	nilResult := convertinternal.AnyTo.PrettyJsonLines(nil)

	// Act
	actual := args.Map{
		"gt0": len(result) > 0,
		"nilLen": len(nilResult),
	}

	// Assert
	expected := args.Map{
		"gt0": true,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- PrettyJsonLines", actual)
}

func Test_AnyTo_Strings_MapAny(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(map[string]any{"a": 1})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings map[string]any", actual)
}

func Test_AnyTo_Strings_MapAnyAny(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(map[any]any{"a": 1})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings map[any]any", actual)
}

func Test_AnyTo_Strings_MapIntString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(map[int]string{1: "a"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings map[int]string", actual)
}

func Test_AnyTo_Strings_MapStringInt(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(map[string]int{"a": 1})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings map[string]int", actual)
}

func Test_AnyTo_Strings_AnySlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings([]any{"a", 1})
	emptyResult := convertinternal.AnyTo.Strings([]any{})

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings any slice", actual)
}

func Test_AnyTo_Strings_Error(t *testing.T) {
	// Arrange
	err := errors.New("line1\nline2")
	var nilErr error
	result := convertinternal.AnyTo.Strings(err)
	nilResult := convertinternal.AnyTo.Strings(nilErr)

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
	expected.ShouldBeEqual(t, 0, "AnyTo returns error -- Strings error", actual)
}

func Test_AnyTo_Strings_EmptyString(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings("")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyTo returns empty -- Strings empty string", actual)
}

func Test_AnyTo_Strings_Int64Slice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings([]int64{1, 2})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings int64 slice", actual)
}

func Test_AnyTo_Strings_Float64Slice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings([]float64{1.1, 2.2})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings float64 slice", actual)
}

func Test_AnyTo_Strings_ByteSlice(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings([]byte{1, 2, 3})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings byte slice", actual)
}

func Test_AnyTo_Strings_Stringer(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(fmt.Errorf("a\nb"))

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings stringer", actual)
}

func Test_AnyTo_Strings_Bool(t *testing.T) {
	// Arrange
	result := convertinternal.AnyTo.Strings(true)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "true",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- Strings bool", actual)
}

func Test_AnyTo_String_NilPtr(t *testing.T) {
	// Arrange
	var p *string
	result := convertinternal.AnyTo.String(p)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyTo returns nil -- String nil ptr", actual)
}

func Test_AnyTo_String_Error(t *testing.T) {
	// Arrange
	err := errors.New("test")
	var nilErr error

	// Act
	actual := args.Map{
		"error":  convertinternal.AnyTo.String(err),
		"nilErr": convertinternal.AnyTo.String(nilErr),
	}

	// Assert
	expected := args.Map{
		"error": "test",
		"nilErr": "",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns error -- String error", actual)
}

// ── CodeFormatter ──

func Test_CodeFormatter_GolangRaw(t *testing.T) {
	// Arrange
	code := []byte("package main\nfunc main(){}")
	result, err := convertinternal.CodeFormatter.GolangRaw(code)
	emptyResult, emptyErr := convertinternal.CodeFormatter.GolangRaw([]byte{})

	// Act
	actual := args.Map{
		"notEmpty": len(result) > 0, "noErr": err == nil,
		"emptyLen": len(emptyResult), "emptyNoErr": emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
		"emptyLen": 0,
		"emptyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CodeFormatter returns correct value -- GolangRaw", actual)
}

func Test_CodeFormatter_Golang(t *testing.T) {
	// Arrange
	result, err := convertinternal.CodeFormatter.Golang("package main\nfunc main(){}")
	emptyResult, emptyErr := convertinternal.CodeFormatter.Golang("")

	// Act
	actual := args.Map{
		"notEmpty": result != "", "noErr": err == nil,
		"emptyResult": emptyResult, "emptyNoErr": emptyErr == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
		"emptyResult": "",
		"emptyNoErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CodeFormatter returns correct value -- Golang", actual)
}

func Test_CodeFormatter_Golang_Invalid(t *testing.T) {
	// Arrange
	_, err := convertinternal.CodeFormatter.Golang("not valid go code {{{")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CodeFormatter returns error -- Golang invalid", actual)
}

// ── Integers ──

func Test_Integers_ToMapBool(t *testing.T) {
	result := convertinternal.Integers.ToMapBool(1, 2, 3)
	emptyResult := convertinternal.Integers.ToMapBool()
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
		"has1": result[1],
	}
	expected := args.Map{
		"len": 3,
		"emptyLen": 0,
		"has1": true,
	}
	expected.ShouldBeEqual(t, 0, "Integers returns correct value -- ToMapBool", actual)
}

func Test_Integers_Int8ToMapBool(t *testing.T) {
	result := convertinternal.Integers.Int8ToMapBool(1, 2)
	emptyResult := convertinternal.Integers.Int8ToMapBool()
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Integers returns correct value -- Int8ToMapBool", actual)
}

func Test_Integers_FromIntegersToMap(t *testing.T) {
	result := convertinternal.Integers.FromIntegersToMap(1, 2)
	emptyResult := convertinternal.Integers.FromIntegersToMap()
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "Integers returns correct value -- FromIntegersToMap", actual)
}

func Test_Integers_IntegersToStrings(t *testing.T) {
	result := convertinternal.Integers.IntegersToStrings([]int{1, 2})
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}
	expected := args.Map{
		"len": 2,
		"first": "1",
	}
	expected.ShouldBeEqual(t, 0, "Integers returns correct value -- IntegersToStrings", actual)
}

// ── KeyValues ──

func Test_KeyValues_ToMap(t *testing.T) {
	result := convertinternal.KeyValuesTo.ToMap([]string{"a", "b"}, []string{"1", "2"})
	nilResult := convertinternal.KeyValuesTo.ToMap(nil, nil)
	actual := args.Map{
		"len": len(result),
		"nilLen": len(nilResult),
		"a": result["a"],
	}
	expected := args.Map{
		"len": 2,
		"nilLen": 0,
		"a": "1",
	}
	expected.ShouldBeEqual(t, 0, "KeyValues returns non-empty -- ToMap", actual)
}

func Test_KeyValues_ToMapPtr(t *testing.T) {
	keys := []string{"a"}
	vals := []string{"1"}
	result := convertinternal.KeyValuesTo.ToMapPtr(&keys, &vals)
	nilResult := convertinternal.KeyValuesTo.ToMapPtr(nil, nil)
	actual := args.Map{
		"notNil": result != nil,
		"nilNotNil": nilResult != nil,
	}
	expected := args.Map{
		"notNil": true,
		"nilNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValues returns non-empty -- ToMapPtr", actual)
}

// ── Map ──

func Test_Map_Keys_StringString(t *testing.T) {
	keys, err := convertinternal.Map.Keys(map[string]string{"a": "1", "b": "2"})
	actual := args.Map{
		"len": len(keys),
		"noErr": err == nil,
	}
	expected := args.Map{
		"len": 2,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Keys string string", actual)
}

func Test_Map_Keys_StringAny(t *testing.T) {
	keys, err := convertinternal.Map.Keys(map[string]any{"a": 1})
	actual := args.Map{
		"len": len(keys),
		"noErr": err == nil,
	}
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Keys string any", actual)
}

func Test_Map_Keys_IntAny(t *testing.T) {
	keys, err := convertinternal.Map.Keys(map[int]any{1: "a"})
	actual := args.Map{
		"len": len(keys),
		"noErr": err == nil,
	}
	expected := args.Map{
		"len": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Keys int any", actual)
}

func Test_Map_Keys_Unsupported(t *testing.T) {
	_, err := convertinternal.Map.Keys("not a map")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Keys unsupported", actual)
}

func Test_Map_KeysValues_StringString(t *testing.T) {
	keys, vals, err := convertinternal.Map.KeysValues(map[string]string{"a": "1"})
	actual := args.Map{
		"kLen": len(keys),
		"vLen": len(vals),
		"noErr": err == nil,
	}
	expected := args.Map{
		"kLen": 1,
		"vLen": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns non-empty -- KeysValues string string", actual)
}

func Test_Map_KeysValues_StringAny(t *testing.T) {
	keys, vals, err := convertinternal.Map.KeysValues(map[string]any{"a": 1})
	actual := args.Map{
		"kLen": len(keys),
		"vLen": len(vals),
		"noErr": err == nil,
	}
	expected := args.Map{
		"kLen": 1,
		"vLen": 1,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns non-empty -- KeysValues string any", actual)
}

func Test_Map_KeysValues_Unsupported(t *testing.T) {
	_, _, err := convertinternal.Map.KeysValues("not a map")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Map returns non-empty -- KeysValues unsupported", actual)
}
