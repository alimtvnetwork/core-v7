package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// DynamicGetters — value extraction and type checks
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_01_Dynamic_Data(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	// Act
	got := d.Data()
	// Assert
	actual := args.Map{"result": got != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_C31_02_Dynamic_Value(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C31_03_Dynamic_Length_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil receiver", actual)
}

func Test_C31_04_Dynamic_Length_Slice(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	actual := args.Map{"result": d.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_05_Dynamic_StructStringPtr_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.StructStringPtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C31_06_Dynamic_StructStringPtr_Cached(t *testing.T) {
	d := coredynamic.NewDynamicPtr("test", true)
	ptr1 := d.StructStringPtr()
	ptr2 := d.StructStringPtr()
	actual := args.Map{"result": ptr1 != ptr2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached pointer to be same", actual)
}

func Test_C31_07_Dynamic_String_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.String() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string", actual)
}

func Test_C31_08_Dynamic_StructString_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.StructString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string", actual)
}

func Test_C31_09_Dynamic_IsNull(t *testing.T) {
	d := coredynamic.NewDynamic(nil, false)
	actual := args.Map{"result": d.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsNull true", actual)
}

func Test_C31_10_Dynamic_IsValid_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	actual := args.Map{"result": d.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": d.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid true", actual)
}

func Test_C31_11_Dynamic_IsPointer_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsPointer()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C31_12_Dynamic_IsPointer_True(t *testing.T) {
	val := "hello"
	d := coredynamic.NewDynamicPtr(&val, true)
	actual := args.Map{"result": d.IsPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsPointer true", actual)
	// call again to test cached path
	actual := args.Map{"result": d.IsPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsPointer true on second call", actual)
}

func Test_C31_13_Dynamic_IsValueType_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsValueType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C31_14_Dynamic_IsValueType_True(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"result": d.IsValueType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsValueType true", actual)
}

func Test_C31_15_Dynamic_IsStructStringNullOrEmpty(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStructStringNullOrEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	d2 := coredynamic.NewDynamicPtr(nil, false)
	actual := args.Map{"result": d2.IsStructStringNullOrEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for null data", actual)
}

func Test_C31_16_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStructStringNullOrEmptyOrWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
}

func Test_C31_17_Dynamic_IsPrimitive(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsPrimitive()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"result": d2.IsPrimitive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for int", actual)
}

func Test_C31_18_Dynamic_IsNumber(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsNumber()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr(3.14, true)
	actual := args.Map{"result": d2.IsNumber()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for float64", actual)
}

func Test_C31_19_Dynamic_IsStringType(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStringType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr("hi", true)
	actual := args.Map{"result": d2.IsStringType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for string", actual)
	d3 := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"result": d3.IsStringType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for int", actual)
}

func Test_C31_20_Dynamic_IsStruct(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStruct()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	type S struct{ X int }
	d2 := coredynamic.NewDynamicPtr(S{X: 1}, true)
	actual := args.Map{"result": d2.IsStruct()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for struct", actual)
}

func Test_C31_21_Dynamic_IsFunc(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsFunc()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr(func() {}, true)
	actual := args.Map{"result": d2.IsFunc()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for func", actual)
}

func Test_C31_22_Dynamic_IsSliceOrArray(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsSliceOrArray()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr([]int{1}, true)
	actual := args.Map{"result": d2.IsSliceOrArray()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for slice", actual)
}

func Test_C31_23_Dynamic_IsSliceOrArrayOrMap(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsSliceOrArrayOrMap()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"result": d2.IsSliceOrArrayOrMap()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for map", actual)
}

func Test_C31_24_Dynamic_IsMap(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsMap()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	d2 := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"result": d2.IsMap()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for map", actual)
}

func Test_C31_25_Dynamic_IntDefault(t *testing.T) {
	var d *coredynamic.Dynamic
	val, ok := d.IntDefault(99)
	actual := args.Map{"result": ok || val != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected default 99, ok=false", actual)
	d2 := coredynamic.NewDynamicPtr("42", true)
	val2, ok2 := d2.IntDefault(0)
	actual := args.Map{"result": ok2 || val2 != 42}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42, ok=true", actual)
	d3 := coredynamic.NewDynamicPtr("abc", true)
	val3, ok3 := d3.IntDefault(7)
	actual := args.Map{"result": ok3 || val3 != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected default 7, ok=false", actual)
}

func Test_C31_26_Dynamic_Float64(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Float64()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	d2 := coredynamic.NewDynamicPtr("3.14", true)
	val, err2 := d2.Float64()
	actual := args.Map{"result": err2 != nil || val != 3.14}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3.14, got err=", actual)
	d3 := coredynamic.NewDynamicPtr("notfloat", true)
	_, err3 := d3.Float64()
	actual := args.Map{"result": err3 == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected parse error", actual)
}

func Test_C31_27_Dynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.ValueInt() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	d2 := coredynamic.NewDynamicValid("notint")
	actual := args.Map{"result": d2.ValueInt() == 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid value", actual)
}

func Test_C31_28_Dynamic_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(uint(10))
	actual := args.Map{"result": d.ValueUInt() != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
	d2 := coredynamic.NewDynamicValid("x")
	actual := args.Map{"result": d2.ValueUInt() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_29_Dynamic_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	actual := args.Map{"result": len(d.ValueStrings()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	d2 := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d2.ValueStrings() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C31_30_Dynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamicValid(true)
	actual := args.Map{"result": d.ValueBool()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	d2 := coredynamic.NewDynamicValid("x")
	actual := args.Map{"result": d2.ValueBool()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C31_31_Dynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamicValid(int64(100))
	actual := args.Map{"result": d.ValueInt64() != 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
	d2 := coredynamic.NewDynamicValid("x")
	actual := args.Map{"result": d2.ValueInt64() == 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_C31_32_Dynamic_ValueNullErr(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil receiver", actual)
	d2 := coredynamic.NewDynamicPtr(nil, false)
	actual := args.Map{"result": d2.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for null data", actual)
	d3 := coredynamic.NewDynamicPtr("ok", true)
	actual := args.Map{"result": d3.ValueNullErr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil error", actual)
}

func Test_C31_33_Dynamic_ValueString(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string", actual)
	d2 := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"result": d2.ValueString() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	d3 := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"result": d3.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

func Test_C31_34_Dynamic_Bytes(t *testing.T) {
	var d *coredynamic.Dynamic
	b, ok := d.Bytes()
	actual := args.Map{"result": ok || b != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, false for nil receiver", actual)
	d2 := coredynamic.NewDynamicPtr([]byte{1, 2}, true)
	b2, ok2 := d2.Bytes()
	actual := args.Map{"result": ok2 || len(b2) != 2}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	d3 := coredynamic.NewDynamicPtr("str", true)
	_, ok3 := d3.Bytes()
	actual := args.Map{"result": ok3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for string", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicJson
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_35_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"result": err != nil || len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty bytes for null", actual)
}

func Test_C31_36_Dynamic_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json bytes", actual)
}

func Test_C31_37_Dynamic_JsonString(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	s, err := d.JsonString()
	actual := args.Map{"result": err != nil || s != "42"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected '42', got ''", actual)
}

func Test_C31_38_Dynamic_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hi", true)
	s := d.JsonStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C31_39_Dynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamicValid(99)
	actual := args.Map{"result": d.JsonModel() != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
	actual := args.Map{"result": d.JsonModelAny() != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_C31_40_Dynamic_Json_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	j := d.Json()
	actual := args.Map{"result": j.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	jp := d.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C31_41_Dynamic_ValueMarshal_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C31_42_Dynamic_Deserialize_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`{}`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C31_43_Dynamic_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	b := d.JsonPayloadMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicReflect
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_44_Dynamic_ReflectSetTo_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil receiver", actual)
}

func Test_C31_45_Dynamic_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	kv, err := d.MapToKeyVal()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": kv.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 entries", actual)
}

func Test_C31_46_Dynamic_ReflectTypeName(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	name := d.ReflectTypeName()
	actual := args.Map{"result": name == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty type name", actual)
}

func Test_C31_47_Dynamic_IsReflectTypeOf(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"result": d.IsReflectTypeOf(reflect.TypeOf(""))}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for string type", actual)
}

func Test_C31_48_Dynamic_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]string{"a", "b", "c"}, true)
	actual := args.Map{"result": d.ItemUsingIndex(1) != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
	rv := d.ItemReflectValueUsingIndex(0)
	actual := args.Map{"result": rv.String() != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_C31_49_Dynamic_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"x": 5}, true)
	actual := args.Map{"result": d.ItemUsingKey("x") != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
	rv := d.ItemReflectValueUsingKey("x")
	actual := args.Map{"result": rv.Interface() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C31_50_Dynamic_Loop_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.Loop(func(i int, item any) bool { return false })
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil data", actual)
}

func Test_C31_51_Dynamic_Loop_WithBreak(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return i == 1
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 iterations", actual)
}

func Test_C31_52_Dynamic_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	actual := args.Map{"result": called}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected called", actual)
}

func Test_C31_53_Dynamic_LoopMap_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C31_54_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4, 5}, true)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt()%2 == 0, false
	})
	actual := args.Map{"result": result.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 even numbers", actual)
}

func Test_C31_55_Dynamic_FilterAsDynamicCollection_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_56_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4, 5}, true)
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 2
	})
	actual := args.Map{"result": result.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// Dynamic Clone / Ptr / NonPtr
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_57_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	c := d.Clone()
	actual := args.Map{"result": c.Value() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_C31_58_Dynamic_ClonePtr_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C31_59_Dynamic_NonPtr_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	np := d.NonPtr()
	actual := args.Map{"result": np.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	p := d.Ptr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// Collection[T] — generic
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_60_Collection_Basic(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(10).Add(20).Add(30)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.First() != 10 || c.Last() != 30}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "wrong first/last", actual)
	actual := args.Map{"result": c.At(1) != 20}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 20 at index 1", actual)
}

func Test_C31_61_Collection_FirstOrDefault_LastOrDefault(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	f, ok := c.FirstOrDefault()
	actual := args.Map{"result": ok || f != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, false for empty", actual)
	l, ok := c.LastOrDefault()
	actual := args.Map{"result": ok || l != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, false for empty", actual)
	c.Add(5)
	f2, ok2 := c.FirstOrDefault()
	actual := args.Map{"result": ok2 || *f2 != 5}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C31_62_Collection_Items_Nil(t *testing.T) {
	var c *coredynamic.Collection[int]
	items := c.Items()
	actual := args.Map{"result": len(items) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C31_63_Collection_Skip_Take_Limit(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	actual := args.Map{"result": len(c.Skip(2)) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": len(c.Take(2)) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(c.Limit(3)) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_64_Collection_SkipCollection_TakeCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	sc := c.SkipCollection(3)
	actual := args.Map{"result": sc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	tc := c.TakeCollection(2)
	actual := args.Map{"result": tc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	lc := c.LimitCollection(3)
	actual := args.Map{"result": lc.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_65_Collection_SafeLimitCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sl := c.SafeLimitCollection(10)
	actual := args.Map{"result": sl.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 (capped)", actual)
}

func Test_C31_66_Collection_AddMany_AddNonNil(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddMany(1, 2, 3)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	val := 99
	c.AddNonNil(&val)
	c.AddNonNil(nil)
	actual := args.Map{"result": c.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C31_67_Collection_RemoveAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": c.RemoveAt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.RemoveAt(99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for invalid index", actual)
}

func Test_C31_68_Collection_Clear_Dispose(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.Clear()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	c.Add(1)
	c.Dispose()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 after dispose", actual)
}

func Test_C31_69_Collection_Loop(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	sum := 0
	c.Loop(func(i int, item int) bool {
		sum += item
		return false
	})
	actual := args.Map{"result": sum != 60}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 60", actual)
}

func Test_C31_70_Collection_Loop_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	called := false
	c.Loop(func(i int, item int) bool {
		called = true
		return false
	})
	actual := args.Map{"result": called}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be called on empty", actual)
}

func Test_C31_71_Collection_Loop_Break(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	count := 0
	c.Loop(func(i int, item int) bool {
		count++
		return i == 2
	})
	actual := args.Map{"result": count != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_72_Collection_Filter(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6})
	evens := c.Filter(func(v int) bool { return v%2 == 0 })
	actual := args.Map{"result": evens.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 evens", actual)
}

func Test_C31_73_Collection_Filter_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	result := c.Filter(func(v int) bool { return true })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_74_Collection_LoopAsync(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.LoopAsync(func(i int, item int) {
		// just ensure it doesn't panic
	})
}

func Test_C31_75_Collection_LoopAsync_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.LoopAsync(func(i int, item int) {
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// Collection paging
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_76_Collection_GetPagesSize(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	actual := args.Map{"result": c.GetPagesSize(2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
	actual := args.Map{"result": c.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for invalid page size", actual)
	actual := args.Map{"result": c.GetPagesSize(-1) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for negative page size", actual)
}

func Test_C31_77_Collection_GetSinglePageCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	page := c.GetSinglePageCollection(3, 2)
	actual := args.Map{"result": page.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_78_Collection_GetSinglePageCollection_SmallCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	page := c.GetSinglePageCollection(5, 1)
	actual := args.Map{"result": page.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_79_Collection_GetPagedCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6, 7})
	pages := c.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
}

func Test_C31_80_Collection_GetPagedCollection_SmallCollection(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	pages := c.GetPagedCollection(5)
	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// Collection serialization
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_81_Collection_MarshalUnmarshalJSON(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	b, err := c.MarshalJSON()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "marshal failed", actual)
	c2 := coredynamic.EmptyCollection[int]()
	err = c2.UnmarshalJSON(b)
	actual := args.Map{"result": err != nil || c2.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unmarshal failed", actual)
}

func Test_C31_82_Collection_JsonString(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	s, err := c.JsonString()
	actual := args.Map{"result": err != nil || s != "[1,2]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [1,2]", actual)
}

func Test_C31_83_Collection_JsonStringMust(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	s := c.JsonStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C31_84_Collection_Strings(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	strs := c.Strings()
	actual := args.Map{"result": len(strs) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_85_Collection_String(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	s := c.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionMethods
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_86_Collection_AddIf(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIf(true, 1)
	c.AddIf(false, 2)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C31_87_Collection_AddManyIf(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(true, 1, 2, 3)
	c.AddManyIf(false, 4, 5)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_88_Collection_AddCollection(t *testing.T) {
	c1 := coredynamic.CollectionFrom([]int{1, 2})
	c2 := coredynamic.CollectionFrom([]int{3, 4})
	c1.AddCollection(c2)
	actual := args.Map{"result": c1.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	c1.AddCollection(nil)
	actual := args.Map{"result": c1.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 4", actual)
}

func Test_C31_89_Collection_AddCollections(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1})
	c2 := coredynamic.CollectionFrom([]int{2, 3})
	c3 := coredynamic.CollectionFrom([]int{4})
	c.AddCollections(c2, nil, c3)
	actual := args.Map{"result": c.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C31_90_Collection_ConcatNew(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	c2 := c.ConcatNew(3, 4)
	actual := args.Map{"result": c2.Length() != 4 || c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ConcatNew should not mutate original", actual)
}

func Test_C31_91_Collection_Clone(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c2 := c.Clone()
	actual := args.Map{"result": c2.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	var nilC *coredynamic.Collection[int]
	c3 := nilC.Clone()
	actual := args.Map{"result": c3.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_92_Collection_Capacity(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	actual := args.Map{"result": c.Capacity() < 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected capacity >= 10", actual)
	var nilC *coredynamic.Collection[int]
	actual := args.Map{"result": nilC.Capacity() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_93_Collection_AddCapacity(t *testing.T) {
	c := coredynamic.NewCollection[int](5)
	c.AddCapacity(10)
	actual := args.Map{"result": c.Capacity() < 15}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 15", actual)
	c.AddCapacity(0)
	c.AddCapacity(-1)
}

func Test_C31_94_Collection_Resize(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.Resize(100)
	actual := args.Map{"result": c.Capacity() < 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected >= 100", actual)
	// no-op resize
	c.Resize(5)
}

func Test_C31_95_Collection_Reverse(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	c.Reverse()
	actual := args.Map{"result": c.At(0) != 4 || c.At(3) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	// single element
	c2 := coredynamic.CollectionFrom([]int{1})
	c2.Reverse()
	actual := args.Map{"result": c2.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single element unchanged", actual)
}

func Test_C31_96_Collection_InsertAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 5})
	c.InsertAt(2, 3, 4)
	actual := args.Map{"result": c.Length() != 5 || c.At(2) != 3 || c.At(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "InsertAt failed", actual)
	// no items
	c.InsertAt(0)
}

func Test_C31_97_Collection_IndexOfFunc_ContainsFunc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	idx := c.IndexOfFunc(func(v int) bool { return v == 20 })
	actual := args.Map{"result": idx != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": c.ContainsFunc(func(v int) bool { return v == 30 })}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected contains", actual)
	actual := args.Map{"result": c.ContainsFunc(func(v int) bool { return v == 99 })}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not contains", actual)
}

func Test_C31_98_Collection_SafeAt(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": c.SafeAt(1) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": c.SafeAt(99) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected zero value for invalid index", actual)
}

func Test_C31_99_Collection_SprintItems(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	items := c.SprintItems("[%v]")
	actual := args.Map{"result": items[0] != "[1]" || items[1] != "[2]"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "format mismatch", actual)
}

func Test_C31_100_Collection_HasIndex(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": c.HasIndex(2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": c.HasIndex(3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": c.HasIndex(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for negative", actual)
}

func Test_C31_101_Collection_Count_Alias(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	actual := args.Map{"result": c.Count() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_102_Collection_HasAnyItem(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	c.Add(1)
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionLock
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_103_CollectionLock_LengthLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	actual := args.Map{"result": c.LengthLock() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_104_CollectionLock_IsEmptyLock(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": c.IsEmptyLock()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C31_105_CollectionLock_AddLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddLock(1).AddLock(2)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_106_CollectionLock_AddsLock_AddManyLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddsLock(1, 2, 3)
	c.AddManyLock(4, 5)
	actual := args.Map{"result": c.Length() != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_C31_107_CollectionLock_AddCollectionLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1})
	c2 := coredynamic.CollectionFrom([]int{2, 3})
	c.AddCollectionLock(c2)
	c.AddCollectionLock(nil)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_108_CollectionLock_AddCollectionsLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1})
	c2 := coredynamic.CollectionFrom([]int{2})
	c3 := coredynamic.CollectionFrom([]int{3})
	c.AddCollectionsLock(c2, nil, c3)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_109_CollectionLock_AddIfLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIfLock(true, 1)
	c.AddIfLock(false, 2)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C31_110_CollectionLock_RemoveAtLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": c.RemoveAtLock(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
	actual := args.Map{"result": c.RemoveAtLock(99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C31_111_CollectionLock_ClearLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.ClearLock()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_112_CollectionLock_ItemsLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	items := c.ItemsLock()
	actual := args.Map{"result": len(items) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_113_CollectionLock_FirstLock_LastLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	actual := args.Map{"result": c.FirstLock() != 10 || c.LastLock() != 30}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "wrong first/last lock", actual)
}

func Test_C31_114_CollectionLock_LoopLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sum := 0
	c.LoopLock(func(i int, item int) bool {
		sum += item
		return false
	})
	actual := args.Map{"result": sum != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
}

func Test_C31_115_CollectionLock_LoopLock_Break(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	count := 0
	c.LoopLock(func(i int, item int) bool {
		count++
		return i == 0
	})
	actual := args.Map{"result": count != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C31_116_CollectionLock_FilterLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	evens := c.FilterLock(func(v int) bool { return v%2 == 0 })
	actual := args.Map{"result": evens.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_117_CollectionLock_StringsLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	strs := c.StringsLock()
	actual := args.Map{"result": len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSearch
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_118_Contains_IndexOf(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})
	actual := args.Map{"result": coredynamic.Contains(c, "b")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected contains b", actual)
	actual := args.Map{"result": coredynamic.IndexOf(c, "d") != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C31_119_Has_HasAll(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": coredynamic.Has(c, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has", actual)
	actual := args.Map{"result": coredynamic.HasAll(c, 1, 2, 3)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has all", actual)
	actual := args.Map{"result": coredynamic.HasAll(c, 1, 4)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	empty := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": coredynamic.HasAll(empty, 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
}

func Test_C31_120_LastIndexOf(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 2, 1})
	actual := args.Map{"result": coredynamic.LastIndexOf(c, 2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": coredynamic.LastIndexOf(c, 99) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C31_121_Count(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 2, 3, 2})
	actual := args.Map{"result": coredynamic.Count(c, 2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_122_ContainsLock_IndexOfLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{10, 20, 30})
	actual := args.Map{"result": coredynamic.ContainsLock(c, 20)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": coredynamic.IndexOfLock(c, 40) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSort
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_123_SortFunc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"result": c.At(0) != 1 || c.At(2) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "sort failed", actual)
}

func Test_C31_124_SortAsc_SortDesc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAsc(c)
	actual := args.Map{"result": c.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 first", actual)
	coredynamic.SortDesc(c)
	actual := args.Map{"result": c.At(0) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 first", actual)
}

func Test_C31_125_SortAscLock_SortDescLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAscLock(c)
	actual := args.Map{"result": c.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	coredynamic.SortDescLock(c)
	actual := args.Map{"result": c.At(0) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_126_SortedAsc_SortedDesc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	asc := coredynamic.SortedAsc(c)
	actual := args.Map{"result": asc.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": c.At(0) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "original should be unchanged", actual)
	desc := coredynamic.SortedDesc(c)
	actual := args.Map{"result": desc.At(0) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_127_IsSorted_IsSortedAsc_IsSortedDesc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": coredynamic.IsSortedAsc(c)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected sorted asc", actual)
	actual := args.Map{"result": coredynamic.IsSortedDesc(c)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not sorted desc", actual)
	single := coredynamic.CollectionFrom([]int{1})
	actual := args.Map{"result": single.IsSorted(func(a, b int) bool { return a < b })}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "single element is always sorted", actual)
}

func Test_C31_128_SortFuncLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{5, 3, 1})
	c.SortFuncLock(func(a, b int) bool { return a < b })
	actual := args.Map{"result": c.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C31_129_SortedFunc(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	sorted := c.SortedFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"result": sorted.At(0) != 1 || c.At(0) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SortedFunc should not mutate original", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionMap
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_130_Map(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	mapped := coredynamic.Map(c, func(v int) string {
		return "x"
	})
	actual := args.Map{"result": mapped.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_131_Map_Nil(t *testing.T) {
	result := coredynamic.Map[int, string](nil, func(v int) string { return "" })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_132_FlatMap(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2})
	result := coredynamic.FlatMap(c, func(v int) []string {
		return []string{"a", "b"}
	})
	actual := args.Map{"result": result.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C31_133_FlatMap_Nil(t *testing.T) {
	result := coredynamic.FlatMap[int, string](nil, func(v int) []string { return nil })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_134_Reduce(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	sum := coredynamic.Reduce(c, 0, func(acc int, item int) int {
		return acc + item
	})
	actual := args.Map{"result": sum != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
}

func Test_C31_135_Reduce_Nil(t *testing.T) {
	result := coredynamic.Reduce[int, int](nil, 42, func(acc int, item int) int { return acc })
	actual := args.Map{"result": result != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected initial value 42", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionDistinct
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_136_Distinct(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 2, 3, 3, 3})
	d := coredynamic.Distinct(c)
	actual := args.Map{"result": d.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C31_137_Distinct_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	d := coredynamic.Distinct(c)
	actual := args.Map{"result": d.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_138_Unique(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b", "a"})
	u := coredynamic.Unique(c)
	actual := args.Map{"result": u.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_139_DistinctLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 1, 2})
	d := coredynamic.DistinctLock(c)
	actual := args.Map{"result": d.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_140_DistinctCount(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 1, 2, 3, 3})
	actual := args.Map{"result": coredynamic.DistinctCount(c) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	empty := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": coredynamic.DistinctCount(empty) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_141_IsDistinct(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": coredynamic.IsDistinct(c)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected distinct", actual)
	c2 := coredynamic.CollectionFrom([]int{1, 2, 2})
	actual := args.Map{"result": coredynamic.IsDistinct(c2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not distinct", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionGroupBy
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_142_GroupBy(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5, 6})
	groups := coredynamic.GroupBy(c, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})
	actual := args.Map{"result": len(groups) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 groups", actual)
	actual := args.Map{"result": groups["even"].Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 evens", actual)
}

func Test_C31_143_GroupBy_Nil(t *testing.T) {
	groups := coredynamic.GroupBy[int, string](nil, func(v int) string { return "" })
	actual := args.Map{"result": len(groups) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_C31_144_GroupByLock(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	groups := coredynamic.GroupByLock(c, func(v int) int { return v % 2 })
	actual := args.Map{"result": len(groups) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 groups", actual)
}

func Test_C31_145_GroupByCount(t *testing.T) {
	c := coredynamic.CollectionFrom([]string{"a", "b", "a", "c", "b", "a"})
	counts := coredynamic.GroupByCount(c, func(v string) string { return v })
	actual := args.Map{"result": counts["a"] != 3 || counts["b"] != 2 || counts["c"] != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "wrong counts", actual)
}

func Test_C31_146_GroupByCount_Nil(t *testing.T) {
	counts := coredynamic.GroupByCount[string, string](nil, func(v string) string { return v })
	actual := args.Map{"result": len(counts) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_147_LeftRight_IsEmpty_HasAnyItem(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}
	actual := args.Map{"result": lr.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	lr2 := &coredynamic.LeftRight{Left: "a", Right: nil}
	actual := args.Map{"result": lr2.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	actual := args.Map{"result": lr2.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has any", actual)
}

func Test_C31_148_LeftRight_HasLeft_HasRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "x", Right: nil}
	actual := args.Map{"result": lr.HasLeft()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasLeft true", actual)
	actual := args.Map{"result": lr.HasRight()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasRight false", actual)
}

func Test_C31_149_LeftRight_IsLeftEmpty_IsRightEmpty(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: "y"}
	actual := args.Map{"result": lr.IsLeftEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected left empty", actual)
	actual := args.Map{"result": lr.IsRightEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected right not empty", actual)
}

func Test_C31_150_LeftRight_LeftToDynamic_RightToDynamic(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ld := lr.LeftToDynamic()
	actual := args.Map{"result": ld.Value() != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	rd := lr.RightToDynamic()
	actual := args.Map{"result": rd.Value() != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
}

func Test_C31_151_LeftRight_NilReceiver(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"result": lr.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	actual := args.Map{"result": lr.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	actual := args.Map{"result": lr.HasLeft()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	actual := args.Map{"result": lr.HasRight()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	actual := args.Map{"result": lr.IsLeftEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	actual := args.Map{"result": lr.IsRightEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	actual := args.Map{"result": lr.LeftToDynamic() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": lr.RightToDynamic() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C31_152_LeftRight_DeserializeLeft_DeserializeRight(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "hello", Right: 42}
	l := lr.DeserializeLeft()
	actual := args.Map{"result": l == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	r := lr.DeserializeRight()
	actual := args.Map{"result": r == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	var nilLR *coredynamic.LeftRight
	actual := args.Map{"result": nilLR.DeserializeLeft() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilLR.DeserializeRight() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C31_153_LeftRight_LeftReflectSet_RightReflectSet(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "hello", Right: "world"}
	var leftTarget string
	err := lr.LeftReflectSet(&leftTarget)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	var rightTarget string
	err = lr.RightReflectSet(&rightTarget)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
}

func Test_C31_154_LeftRight_ReflectSet_NilReceiver(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"result": lr.LeftReflectSet(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
	actual := args.Map{"result": lr.RightReflectSet(nil) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_C31_155_LeftRight_TypeStatus(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()
	actual := args.Map{"result": ts.IsSame}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same type", actual)
	var nilLR *coredynamic.LeftRight
	ts2 := nilLR.TypeStatus()
	_ = ts2
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_156_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("test msg")
	c := ds.Clone()
	actual := args.Map{"result": c.Message != "test msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test msg", actual)
}

func Test_C31_157_DynamicStatus_ClonePtr(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	cp := ds.ClonePtr()
	actual := args.Map{"result": cp == nil || cp.Message != "msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected msg", actual)
	var nilDS *coredynamic.DynamicStatus
	actual := args.Map{"result": nilDS.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C31_158_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	actual := args.Map{"result": ds.Message != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty message", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionTypes — factory functions
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_159_CollectionTypes_Factories(t *testing.T) {
	sc := coredynamic.NewStringCollection(2)
	sc.Add("a")
	actual := args.Map{"result": sc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	es := coredynamic.EmptyStringCollection()
	actual := args.Map{"result": es.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	ic := coredynamic.NewIntCollection(2)
	ic.Add(1)
	actual := args.Map{"result": ic.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	eic := coredynamic.EmptyIntCollection()
	actual := args.Map{"result": eic.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	i64c := coredynamic.NewInt64Collection(2)
	i64c.Add(int64(1))
	bc := coredynamic.NewByteCollection(2)
	bc.Add(byte(1))
	boolC := coredynamic.NewBoolCollection(2)
	boolC.Add(true)
	f64c := coredynamic.NewFloat64Collection(2)
	f64c.Add(1.0)
	amc := coredynamic.NewAnyMapCollection(2)
	amc.Add(map[string]any{"a": 1})
	smc := coredynamic.NewStringMapCollection(2)
	smc.Add(map[string]string{"a": "b"})
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionFrom / CollectionClone
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_160_CollectionFrom_NilSlice(t *testing.T) {
	c := coredynamic.CollectionFrom[int](nil)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_161_CollectionClone(t *testing.T) {
	original := []int{1, 2, 3}
	c := coredynamic.CollectionClone(original)
	original[0] = 99
	actual := args.Map{"result": c.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected deep copy", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// Collection AddWithWgLock
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_162_CollectionLock_AddWithWgLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, 42)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C31_163_DynamicCollection_Strings_String(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.Add(coredynamic.NewDynamicValid("a"))
	dc.Add(coredynamic.NewDynamicValid("b"))
	strs := dc.Strings()
	actual := args.Map{"result": len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	s := dc.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C31_164_DynamicCollection_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAnyNonNull(nil, true)
	actual := args.Map{"result": dc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
	dc.AddAnyNonNull("hello", true)
	actual := args.Map{"result": dc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C31_165_DynamicCollection_AddPtr_AddManyPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	d1 := coredynamic.NewDynamicPtr("a", true)
	dc.AddPtr(d1)
	dc.AddPtr(nil)
	actual := args.Map{"result": dc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	d2 := coredynamic.NewDynamicPtr("b", true)
	dc.AddManyPtr(d2, nil)
	actual := args.Map{"result": dc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_166_DynamicCollection_RemoveAt(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	actual := args.Map{"result": dc.RemoveAt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
	actual := args.Map{"result": dc.RemoveAt(99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C31_167_DynamicCollection_ListStrings(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(42, true)
	strs := dc.ListStrings()
	actual := args.Map{"result": len(strs) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C31_168_DynamicCollection_AnyItems(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	items := dc.AnyItems()
	actual := args.Map{"result": len(items) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"result": len(empty.AnyItems()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_169_DynamicCollection_AnyItemsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	ac := dc.AnyItemsCollection()
	actual := args.Map{"result": ac.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"result": empty.AnyItemsCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_170_DynamicCollection_AddAnyMany(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyMany("a", "b", "c")
	actual := args.Map{"result": dc.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	dc.AddAnyMany()
}

func Test_C31_171_DynamicCollection_JsonModel_JsonModelAny(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	m := dc.JsonModel()
	actual := args.Map{"result": len(m.Items) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 item in model", actual)
	ma := dc.JsonModelAny()
	actual := args.Map{"result": ma == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C31_172_DynamicCollection_Json_JsonPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	j := dc.Json()
	_ = j
	jp := dc.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C31_173_DynamicCollection_Paging(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagesSize(3)
	actual := args.Map{"result": pages != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	actual := args.Map{"result": dc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_174_DynamicCollection_Items_NilReceiver(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	items := dc.Items()
	actual := args.Map{"result": len(items) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C31_175_DynamicCollection_Length_NilReceiver(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"result": dc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C31_176_DynamicCollection_IsEmpty_NilReceiver(t *testing.T) {
	var dc *coredynamic.DynamicCollection
	actual := args.Map{"result": dc.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C31_177_DynamicCollection_Loop(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("a", true).AddAny("b", true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_178_DynamicCollection_Loop_Break(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return i == 1
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C31_179_DynamicCollection_HasIndex(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	actual := args.Map{"result": dc.HasIndex(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": dc.HasIndex(1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}
