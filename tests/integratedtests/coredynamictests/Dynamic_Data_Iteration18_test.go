package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// DynamicGetters — type checks, value extraction
// ══════════════════════════════════════════════════════════════════════════════

func Test_Dynamic_Data(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"val": d.Data()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Data", actual)
}

func Test_Dynamic_Value(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)

	// Act
	actual := args.Map{"val": d.Value()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Value", actual)
}

func Test_Dynamic_Length_Slice(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1, 2, 3}, true)

	// Act
	actual := args.Map{"len": d.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Length slice", actual)
}

func Test_Dynamic_Length_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"len": d.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- Length nil", actual)
}

func Test_Dynamic_StructStringPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	ptr := d.StructStringPtr()

	// Act
	actual := args.Map{"notNil": ptr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- StructStringPtr", actual)
}

func Test_Dynamic_StructStringPtr_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"nil": d.StructStringPtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- StructStringPtr nil", actual)
}

func Test_Dynamic_String(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("world", true)

	// Act
	actual := args.Map{"val": d.String()}

	// Assert
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- String", actual)
}

func Test_Dynamic_String_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"val": d.String()}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- String nil", actual)
}

func Test_Dynamic_IsNull(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(nil, false)

	// Act
	actual := args.Map{"null": d.IsNull()}

	// Assert
	expected := args.Map{"null": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsNull", actual)
}

func Test_Dynamic_IsValid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("ok", true)

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"invalid": d.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- IsValid", actual)
}

func Test_Dynamic_IsPointer(t *testing.T) {
	// Arrange
	x := 42
	d := coredynamic.NewDynamic(&x, true)

	// Act
	actual := args.Map{"ptr": d.IsPointer()}

	// Assert
	expected := args.Map{"ptr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsPointer", actual)
}

func Test_Dynamic_IsPointer_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"ptr": d.IsPointer()}

	// Assert
	expected := args.Map{"ptr": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IsPointer nil", actual)
}

func Test_Dynamic_IsValueType(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)

	// Act
	actual := args.Map{"vt": d.IsValueType()}

	// Assert
	expected := args.Map{"vt": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsValueType", actual)
}

func Test_Dynamic_IsValueType_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"vt": d.IsValueType()}

	// Assert
	expected := args.Map{"vt": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IsValueType nil", actual)
}

func Test_Dynamic_IsStructStringNullOrEmpty(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(nil, false)

	// Act
	actual := args.Map{"empty": d.IsStructStringNullOrEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns empty -- IsStructStringNullOrEmpty", actual)
}

func Test_Dynamic_IsStructStringNullOrEmpty_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"empty": d.IsStructStringNullOrEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IsStructStringNullOrEmpty nil", actual)
}

func Test_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("   ", true)

	// Act
	actual := args.Map{"ws": d.IsStructStringNullOrEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"ws": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns empty -- IsStructStringNullOrEmptyOrWhitespace", actual)
}

func Test_Dynamic_IsPrimitive(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)

	// Act
	actual := args.Map{"prim": d.IsPrimitive()}

	// Assert
	expected := args.Map{"prim": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsPrimitive", actual)
}

func Test_Dynamic_IsPrimitive_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"prim": d.IsPrimitive()}

	// Assert
	expected := args.Map{"prim": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IsPrimitive nil", actual)
}

func Test_Dynamic_IsNumber(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(3.14, true)

	// Act
	actual := args.Map{"num": d.IsNumber()}

	// Assert
	expected := args.Map{"num": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsNumber", actual)
}

func Test_Dynamic_IsNumber_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"num": d.IsNumber()}

	// Assert
	expected := args.Map{"num": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IsNumber nil", actual)
}

func Test_Dynamic_IsStringType(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)

	// Act
	actual := args.Map{"str": d.IsStringType()}

	// Assert
	expected := args.Map{"str": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsStringType", actual)
}

func Test_Dynamic_IsStringType_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"str": d.IsStringType()}

	// Assert
	expected := args.Map{"str": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IsStringType nil", actual)
}

func Test_Dynamic_IsStruct(t *testing.T) {
	// Arrange
	type s struct{ X int }
	d := coredynamic.NewDynamic(s{X: 1}, true)

	// Act
	actual := args.Map{"st": d.IsStruct()}

	// Assert
	expected := args.Map{"st": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsStruct", actual)
}

func Test_Dynamic_IsFunc(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(func() {}, true)

	// Act
	actual := args.Map{"fn": d.IsFunc()}

	// Assert
	expected := args.Map{"fn": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsFunc", actual)
}

func Test_Dynamic_IsSliceOrArray(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1}, true)

	// Act
	actual := args.Map{"sa": d.IsSliceOrArray()}

	// Assert
	expected := args.Map{"sa": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsSliceOrArray", actual)
}

func Test_Dynamic_IsSliceOrArrayOrMap(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]int{"a": 1}, true)

	// Act
	actual := args.Map{"sam": d.IsSliceOrArrayOrMap()}

	// Assert
	expected := args.Map{"sam": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsSliceOrArrayOrMap", actual)
}

func Test_Dynamic_IsMap(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]int{"a": 1}, true)

	// Act
	actual := args.Map{"m": d.IsMap()}

	// Assert
	expected := args.Map{"m": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsMap", actual)
}

func Test_Dynamic_IntDefault(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("42", true)
	val, ok := d.IntDefault(0)

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IntDefault", actual)
}

func Test_Dynamic_IntDefault_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	val, ok := d.IntDefault(99)

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 99,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- IntDefault nil", actual)
}

func Test_Dynamic_IntDefault_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)
	val, ok := d.IntDefault(7)

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 7,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- IntDefault invalid", actual)
}

func Test_Dynamic_Float64(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("3.14", true)
	val, err := d.Float64()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"close": val > 3.13 && val < 3.15,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"close": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Float64", actual)
}

func Test_Dynamic_Float64_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, err := d.Float64()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- Float64 nil", actual)
}

func Test_Dynamic_Float64_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)
	_, err := d.Float64()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- Float64 invalid", actual)
}

func Test_Dynamic_ValueInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)

	// Act
	actual := args.Map{"val": d.ValueInt()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueInt", actual)
}

func Test_Dynamic_ValueInt_NotInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)

	// Act
	actual := args.Map{"val": d.ValueInt()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueInt not int", actual)
}

func Test_Dynamic_ValueUInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(uint(5), true)

	// Act
	actual := args.Map{"val": d.ValueUInt()}

	// Assert
	expected := args.Map{"val": uint(5)}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueUInt", actual)
}

func Test_Dynamic_ValueStrings(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]string{"a", "b"}, true)

	// Act
	actual := args.Map{"len": len(d.ValueStrings())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- ValueStrings", actual)
}

func Test_Dynamic_ValueBool(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(true, true)

	// Act
	actual := args.Map{"val": d.ValueBool()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueBool", actual)
}

func Test_Dynamic_ValueInt64(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(int64(999), true)

	// Act
	actual := args.Map{"val": d.ValueInt64()}

	// Assert
	expected := args.Map{"val": int64(999)}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueInt64", actual)
}

func Test_Dynamic_ValueNullErr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("ok", true)

	// Act
	actual := args.Map{"nil": d.ValueNullErr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- ValueNullErr", actual)
}

func Test_Dynamic_ValueNullErr_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- ValueNullErr nil", actual)
}

func Test_Dynamic_ValueString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	actual := args.Map{"val": d.ValueString()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- ValueString", actual)
}

func Test_Dynamic_ValueString_NonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	s := d.ValueString()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- ValueString non-string", actual)
}

func Test_Dynamic_Bytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]byte{1, 2, 3}, true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(b),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Bytes", actual)
}

func Test_Dynamic_Bytes_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, ok := d.Bytes()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- Bytes nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicReflect — reflect operations, loops, filters
// ══════════════════════════════════════════════════════════════════════════════

func Test_Dynamic_ReflectValue(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	rv := d.ReflectValue()

	// Act
	actual := args.Map{
		"valid": rv.IsValid(),
		"kind": rv.Kind().String(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"kind": "int",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ReflectValue", actual)
}

func Test_Dynamic_ReflectKind(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)

	// Act
	actual := args.Map{"kind": d.ReflectKind()}

	// Assert
	expected := args.Map{"kind": reflect.String}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ReflectKind", actual)
}

func Test_Dynamic_ReflectType(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	rt := d.ReflectType()

	// Act
	actual := args.Map{"name": rt.String()}

	// Assert
	expected := args.Map{"name": "int"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ReflectType", actual)
}

func Test_Dynamic_IsReflectTypeOf(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)

	// Act
	actual := args.Map{"match": d.IsReflectTypeOf(reflect.TypeOf(""))}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsReflectTypeOf", actual)
}

func Test_Dynamic_IsReflectKind(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)

	// Act
	actual := args.Map{"match": d.IsReflectKind(reflect.Int)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsReflectKind", actual)
}

func Test_Dynamic_ItemUsingIndex(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]string{"a", "b", "c"}, true)

	// Act
	actual := args.Map{"val": d.ItemUsingIndex(1)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ItemUsingIndex", actual)
}

func Test_Dynamic_ItemUsingKey(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]int{"x": 5}, true)

	// Act
	actual := args.Map{"val": d.ItemUsingKey("x")}

	// Assert
	expected := args.Map{"val": 5}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ItemUsingKey", actual)
}

func Test_Dynamic_ReflectSetTo(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	var target string
	err := d.ReflectSetTo(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ReflectSetTo", actual)
}

func Test_Dynamic_ReflectSetTo_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- ReflectSetTo nil", actual)
}

func Test_Dynamic_ConvertUsingFunc(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	converter := func(val any, expectedType reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(val)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(""))

	// Act
	actual := args.Map{"valid": result.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ConvertUsingFunc", actual)
}

func Test_Dynamic_Loop(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{10, 20, 30}, true)
	sum := 0
	called := d.Loop(func(i int, item any) bool {
		sum += item.(int)
		return false
	})

	// Act
	actual := args.Map{
		"called": called,
		"sum": sum,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"sum": 60,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Loop", actual)
}

func Test_Dynamic_Loop_Break(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{10, 20, 30}, true)
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return i == 0
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Loop break", actual)
}

func Test_Dynamic_Loop_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(nil, false)
	called := d.Loop(func(i int, item any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- Loop invalid", actual)
}

func Test_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1, 2, 3, 4, 5}, true)
	filtered := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt() > 2, false
	})

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- FilterAsDynamicCollection", actual)
}

func Test_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic([]int{1, 2, 3, 4}, true)
	filtered := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 1
	})

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- FilterAsDynamicCollection break", actual)
}

func Test_Dynamic_LoopMap(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]int{"a": 1, "b": 2}, true)
	count := 0
	called := d.LoopMap(func(i int, k, v any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{
		"called": called,
		"count": count,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"count": 2,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- LoopMap", actual)
}

func Test_Dynamic_LoopMap_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(nil, false)
	called := d.LoopMap(func(i int, k, v any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- LoopMap invalid", actual)
}

func Test_Dynamic_MapToKeyVal(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]int{"a": 1}, true)
	kvc, err := d.MapToKeyVal()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": kvc != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- MapToKeyVal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicJson — JSON methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_Dynamic_ValueMarshal(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	b, err := d.ValueMarshal()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ValueMarshal", actual)
}

func Test_Dynamic_ValueMarshal_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- ValueMarshal nil", actual)
}

func Test_Dynamic_JsonPayloadMust(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("test", true)
	b := d.JsonPayloadMust()

	// Act
	actual := args.Map{"hasBytes": len(b) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonPayloadMust", actual)
}

func Test_Dynamic_JsonBytesPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("abc", true)
	b, err := d.JsonBytesPtr()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonBytesPtr", actual)
}

func Test_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(nil, false)
	b, err := d.JsonBytesPtr()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"empty": len(b) == 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonBytesPtr null", actual)
}

func Test_Dynamic_MarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	b, err := d.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- MarshalJSON", actual)
}

func Test_Dynamic_JsonModel(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("val", true)

	// Act
	actual := args.Map{"val": d.JsonModel()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonModel", actual)
}

func Test_Dynamic_JsonModelAny(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)

	// Act
	actual := args.Map{"val": d.JsonModelAny()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonModelAny", actual)
}

func Test_Dynamic_Json(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("test", true)
	jr := d.Json()

	// Act
	actual := args.Map{"notNil": &jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Json", actual)
}

func Test_Dynamic_JsonPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("test", true)
	jr := d.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonPtr", actual)
}

func Test_Dynamic_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	seed := "initial"
	d := coredynamic.NewDynamic(&seed, true)
	jr := corejson.NewPtr("updated")
	result, err := d.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ParseInjectUsingJson", actual)
}

func Test_Dynamic_JsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	s, err := d.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonString", actual)
}

func Test_Dynamic_JsonStringMust(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("world", true)
	s := d.JsonStringMust()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- JsonStringMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionLock — thread-safe operations
// ══════════════════════════════════════════════════════════════════════════════

func Test_CollectionLock_LengthLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b"})

	// Act
	actual := args.Map{"len": c.LengthLock()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- LengthLock", actual)
}

func Test_CollectionLock_IsEmptyLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()

	// Act
	actual := args.Map{"empty": c.IsEmptyLock()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns empty -- IsEmptyLock", actual)
}

func Test_CollectionLock_AddLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddLock("x")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- AddLock", actual)
}

func Test_CollectionLock_AddsLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddsLock("a", "b", "c")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- AddsLock", actual)
}

func Test_CollectionLock_AddCollectionLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c1 := coredynamic.CollectionFrom[string]([]string{"a"})
	c2 := coredynamic.CollectionFrom[string]([]string{"b", "c"})
	c1.AddCollectionLock(c2)

	// Act
	actual := args.Map{"len": c1.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- AddCollectionLock", actual)
}

func Test_CollectionLock_AddCollectionLock_Nil(t *testing.T) {
	// Arrange
	c1 := coredynamic.CollectionFrom[string]([]string{"a"})
	c1.AddCollectionLock(nil)

	// Act
	actual := args.Map{"len": c1.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns nil -- AddCollectionLock nil", actual)
}

func Test_CollectionLock_AddCollectionsLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c1 := coredynamic.CollectionFrom[string]([]string{"a"})
	c2 := coredynamic.CollectionFrom[string]([]string{"b"})
	c.AddCollectionsLock(c1, nil, c2)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- AddCollectionsLock", actual)
}

func Test_CollectionLock_AddIfLock_True(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddIfLock(true, "x")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns non-empty -- AddIfLock true", actual)
}

func Test_CollectionLock_AddIfLock_False(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddIfLock(false, "x")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns non-empty -- AddIfLock false", actual)
}

func Test_CollectionLock_RemoveAtLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b", "c"})
	ok := c.RemoveAtLock(1)

	// Act
	actual := args.Map{
		"ok": ok,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- RemoveAtLock", actual)
}

func Test_CollectionLock_RemoveAtLock_Invalid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a"})
	ok := c.RemoveAtLock(5)

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns error -- RemoveAtLock invalid", actual)
}

func Test_CollectionLock_ClearLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b"})
	c.ClearLock()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- ClearLock", actual)
}

func Test_CollectionLock_ItemsLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b"})
	items := c.ItemsLock()

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- ItemsLock", actual)
}

func Test_CollectionLock_FirstLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b"})

	// Act
	actual := args.Map{"val": c.FirstLock()}

	// Assert
	expected := args.Map{"val": "a"}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- FirstLock", actual)
}

func Test_CollectionLock_LastLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b"})

	// Act
	actual := args.Map{"val": c.LastLock()}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- LastLock", actual)
}

func Test_CollectionLock_AddWithWgLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, "x")
	wg.Wait()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns non-empty -- AddWithWgLock", actual)
}

func Test_CollectionLock_LoopLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b", "c"})
	count := 0
	c.LoopLock(func(i int, item string) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- LoopLock", actual)
}

func Test_CollectionLock_LoopLock_Break(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b", "c"})
	count := 0
	c.LoopLock(func(i int, item string) bool {
		count++
		return true
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- LoopLock break", actual)
}

func Test_CollectionLock_FilterLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "bb", "ccc"})
	filtered := c.FilterLock(func(s string) bool {
		return len(s) > 1
	})

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- FilterLock", actual)
}

func Test_CollectionLock_StringsLock_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string]([]string{"a", "b"})
	strs := c.StringsLock()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CollectionLock returns correct value -- StringsLock", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAnyItems — Add, Get, Paging, JSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapAnyItems_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{
		"empty": m.IsEmpty(),
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- Empty", actual)
}

func Test_MapAnyItems_NewUsingItems(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"len": m.Length(),
		"has": m.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- NewUsingItems", actual)
}

func Test_MapAnyItems_NewUsingItems_Nil(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(nil)

	// Act
	actual := args.Map{"empty": m.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- NewUsingItems nil", actual)
}

func Test_MapAnyItems_HasKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})

	// Act
	actual := args.Map{
		"has": m.HasKey("k"),
		"miss": m.HasKey("z"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- HasKey", actual)
}

func Test_MapAnyItems_HasKey_Nil(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"has": m.HasKey("x")}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- HasKey nil", actual)
}

func Test_MapAnyItems_Add(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.Add("k", "v")
	isNew2 := m.Add("k", "v2")

	// Act
	actual := args.Map{
		"isNew": isNew,
		"isNew2": isNew2,
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"isNew": true,
		"isNew2": false,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Add", actual)
}

func Test_MapAnyItems_Set(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	isNew := m.Set("k", "v")

	// Act
	actual := args.Map{"isNew": isNew}

	// Assert
	expected := args.Map{"isNew": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Set", actual)
}

func Test_MapAnyItems_GetValue(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": 42})

	// Act
	actual := args.Map{
		"val": m.GetValue("k"),
		"nil": m.GetValue("z"),
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"nil": nil,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetValue", actual)
}

func Test_MapAnyItems_Get_FromDynamicDataIteration(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": "v"})
	v, has := m.Get("k")
	_, miss := m.Get("z")

	// Act
	actual := args.Map{
		"val": v,
		"has": has,
		"miss": miss,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"has": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Get", actual)
}

func Test_MapAnyItems_AddMapResult(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.AddMapResult(map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AddMapResult", actual)
}

func Test_MapAnyItems_AddMapResult_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	m.AddMapResult(nil)

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- AddMapResult empty", actual)
}

func Test_MapAnyItems_GetPagesSize(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2, "c": 3})

	// Act
	actual := args.Map{"pages": m.GetPagesSize(2)}

	// Assert
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetPagesSize", actual)
}

func Test_MapAnyItems_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()

	// Act
	actual := args.Map{"pages": m.GetPagesSize(0)}

	// Assert
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetPagesSize zero", actual)
}

func Test_MapAnyItems_JsonString(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	s, err := m.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonString", actual)
}

func Test_MapAnyItems_JsonStringMust(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	s := m.JsonStringMust()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- JsonStringMust", actual)
}

func Test_MapAnyItems_AllKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"b": 2, "a": 1})
	keys := m.AllKeysSorted()

	// Act
	actual := args.Map{
		"len": len(keys),
		"first": keys[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- AllKeysSorted", actual)
}

func Test_MapAnyItems_AllValues(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	vals := m.AllValues()

	// Act
	actual := args.Map{"len": len(vals)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AllValues", actual)
}

func Test_MapAnyItems_Clear(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m.Clear()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Clear", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2, "c": 3})
	sub := m.GetNewMapUsingKeys(false, "a", "c")

	// Act
	actual := args.Map{"len": sub.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- GetNewMapUsingKeys", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys_Empty(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	sub := m.GetNewMapUsingKeys(false)

	// Act
	actual := args.Map{"empty": sub.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns empty -- GetNewMapUsingKeys empty", actual)
}

func Test_MapAnyItems_AddWithValidation_Match(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddWithValidation(reflect.TypeOf(""), "k", "v")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AddWithValidation match", actual)
}

func Test_MapAnyItems_AddWithValidation_Mismatch(t *testing.T) {
	// Arrange
	m := coredynamic.EmptyMapAnyItems()
	err := m.AddWithValidation(reflect.TypeOf(0), "k", "v")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns non-empty -- AddWithValidation mismatch", actual)
}

func Test_MapAnyItems_Nil_Length(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns nil -- nil Length", actual)
}
