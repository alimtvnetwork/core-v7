package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// DynamicGetters — value extraction, type checks
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_01_Dynamic_Data(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.Data() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_02_Dynamic_Value(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"result": d.Value() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_C34_03_Dynamic_Length_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_04_Dynamic_Length_Slice(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	actual := args.Map{"result": d.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_05_Dynamic_StructStringPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	ptr := d.StructStringPtr()
	actual := args.Map{"result": ptr == nil || *ptr == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty string ptr", actual)
	// second call should return cached
	ptr2 := d.StructStringPtr()
	actual := args.Map{"result": ptr != ptr2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached pointer", actual)
}

func Test_C34_06_Dynamic_StructStringPtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.StructStringPtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C34_07_Dynamic_String(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C34_08_Dynamic_String_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.String() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_09_Dynamic_StructString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.StructString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_10_Dynamic_IsNull(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	actual := args.Map{"result": d.IsNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected null", actual)
}

func Test_C34_11_Dynamic_IsValid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

func Test_C34_12_Dynamic_IsInvalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	actual := args.Map{"result": d.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_C34_13_Dynamic_IsPointer_True(t *testing.T) {
	val := 42
	d := coredynamic.NewDynamicValid(&val)
	actual := args.Map{"result": d.IsPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected pointer", actual)
}

func Test_C34_14_Dynamic_IsPointer_False(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.IsPointer()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-pointer", actual)
}

func Test_C34_15_Dynamic_IsPointer_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsPointer()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C34_16_Dynamic_IsValueType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.IsValueType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected value type", actual)
}

func Test_C34_17_Dynamic_IsValueType_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsValueType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C34_18_Dynamic_IsStructStringNullOrEmpty(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	actual := args.Map{"result": d.IsStructStringNullOrEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil data", actual)
}

func Test_C34_19_Dynamic_IsStructStringNullOrEmpty_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStructStringNullOrEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil receiver", actual)
}

func Test_C34_20_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	actual := args.Map{"result": d.IsStructStringNullOrEmptyOrWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_21_Dynamic_IsStructStringNullOrEmptyOrWhitespace_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStructStringNullOrEmptyOrWhitespace()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
}

func Test_C34_22_Dynamic_IsPrimitive(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.IsPrimitive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected primitive", actual)
}

func Test_C34_23_Dynamic_IsPrimitive_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsPrimitive()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C34_24_Dynamic_IsNumber(t *testing.T) {
	d := coredynamic.NewDynamicValid(3.14)
	actual := args.Map{"result": d.IsNumber()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected number", actual)
}

func Test_C34_25_Dynamic_IsNumber_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsNumber()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_26_Dynamic_IsStringType(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"result": d.IsStringType()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_27_Dynamic_IsStringType_NotString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.IsStringType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_28_Dynamic_IsStringType_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStringType()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_29_Dynamic_IsStruct(t *testing.T) {
	type s struct{ X int }
	d := coredynamic.NewDynamicValid(s{X: 1})
	actual := args.Map{"result": d.IsStruct()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected struct", actual)
}

func Test_C34_30_Dynamic_IsStruct_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsStruct()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_31_Dynamic_IsFunc(t *testing.T) {
	d := coredynamic.NewDynamicValid(func() {})
	actual := args.Map{"result": d.IsFunc()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected func", actual)
}

func Test_C34_32_Dynamic_IsFunc_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsFunc()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_33_Dynamic_IsSliceOrArray(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1})
	actual := args.Map{"result": d.IsSliceOrArray()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_34_Dynamic_IsSliceOrArray_Array(t *testing.T) {
	d := coredynamic.NewDynamicValid([2]int{1, 2})
	actual := args.Map{"result": d.IsSliceOrArray()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for array", actual)
}

func Test_C34_35_Dynamic_IsSliceOrArray_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsSliceOrArray()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_36_Dynamic_IsSliceOrArrayOrMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	actual := args.Map{"result": d.IsSliceOrArrayOrMap()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_37_Dynamic_IsSliceOrArrayOrMap_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsSliceOrArrayOrMap()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_38_Dynamic_IsMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	actual := args.Map{"result": d.IsMap()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_39_Dynamic_IsMap_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.IsMap()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_40_Dynamic_IntDefault_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	v, ok := d.IntDefault(0)
	actual := args.Map{"result": ok || v != 42}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42/true, got/", actual)
}

func Test_C34_41_Dynamic_IntDefault_Invalid(t *testing.T) {
	d := coredynamic.NewDynamicValid("not a number")
	v, ok := d.IntDefault(99)
	actual := args.Map{"result": ok || v != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99/false, got/", actual)
}

func Test_C34_42_Dynamic_IntDefault_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	v, ok := d.IntDefault(5)
	actual := args.Map{"result": ok || v != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected default for nil", actual)
}

func Test_C34_43_Dynamic_Float64_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(3.14)
	v, err := d.Float64()
	actual := args.Map{"result": err != nil || v != 3.14}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3.14, got err=", actual)
}

func Test_C34_44_Dynamic_Float64_Invalid(t *testing.T) {
	d := coredynamic.NewDynamicValid("abc")
	_, err := d.Float64()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_45_Dynamic_Float64_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Float64()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C34_46_Dynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.ValueInt() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_47_Dynamic_ValueInt_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid("str")
	v := d.ValueInt()
	actual := args.Map{"result": v == 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be 42", actual)
}

func Test_C34_48_Dynamic_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(uint(10))
	actual := args.Map{"result": d.ValueUInt() != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
}

func Test_C34_49_Dynamic_ValueUInt_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid("str")
	actual := args.Map{"result": d.ValueUInt() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_50_Dynamic_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	v := d.ValueStrings()
	actual := args.Map{"result": len(v) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_51_Dynamic_ValueStrings_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.ValueStrings() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C34_52_Dynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamicValid(true)
	actual := args.Map{"result": d.ValueBool()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_53_Dynamic_ValueBool_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.ValueBool()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_54_Dynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamicValid(int64(100))
	actual := args.Map{"result": d.ValueInt64() != 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_C34_55_Dynamic_ValueInt64_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	v := d.ValueInt64()
	_ = v // just ensure no panic
}

func Test_C34_56_Dynamic_ValueNullErr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ValueNullErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_57_Dynamic_ValueNullErr_NullData(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	err := d.ValueNullErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for null data", actual)
}

func Test_C34_58_Dynamic_ValueNullErr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	err := d.ValueNullErr()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C34_59_Dynamic_ValueString_String(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"result": d.ValueString() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_C34_60_Dynamic_ValueString_NonString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.ValueString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C34_61_Dynamic_ValueString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_62_Dynamic_Bytes_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid([]byte{1, 2, 3})
	b, ok := d.Bytes()
	actual := args.Map{"result": ok || len(b) != 3}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_63_Dynamic_Bytes_WrongType(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	_, ok := d.Bytes()
	actual := args.Map{"result": ok}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_64_Dynamic_Bytes_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	b, ok := d.Bytes()
	actual := args.Map{"result": ok || b != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil/false", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicJson — serialization/deserialization
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_65_Dynamic_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`{}`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C34_66_Dynamic_ValueMarshal(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.ValueMarshal()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_67_Dynamic_ValueMarshal_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_68_Dynamic_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b := d.JsonPayloadMust()
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_69_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamic(nil, true)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error for null", actual)
	actual := args.Map{"result": len(b) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty bytes", actual)
}

func Test_C34_70_Dynamic_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.JsonBytesPtr()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_71_Dynamic_MarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicValid("test")
	b, err := d.MarshalJSON()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_72_Dynamic_UnmarshalJSON_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`42`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C34_73_Dynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	actual := args.Map{"result": d.JsonModel() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_74_Dynamic_JsonModelAny(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	actual := args.Map{"result": d.JsonModelAny() != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x", actual)
}

func Test_C34_75_Dynamic_Json(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	j := d.Json()
	_ = j
}

func Test_C34_76_Dynamic_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_77_Dynamic_JsonBytes(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.JsonBytes()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_78_Dynamic_JsonString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s, err := d.JsonString()
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_C34_79_Dynamic_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.JsonStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C34_80_Dynamic_ParseInjectUsingJson(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	d2 := coredynamic.NewDynamicValid(0)
	_, err := d2.ParseInjectUsingJson(jp)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error from Dynamic json round-trip with untyped destination", actual)
}

func Test_C34_81_Dynamic_ParseInjectUsingJsonMust(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	d2 := coredynamic.NewDynamicValid(0)

	didPanic := false
	func() {
		defer func() {
			if recover() != nil {
				didPanic = true
			}
		}()
		_ = d2.ParseInjectUsingJsonMust(jp)
	}()

	actual := args.Map{"result": didPanic}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic from ParseInjectUsingJsonMust when unmarshal fails", actual)
}

func Test_C34_82_Dynamic_JsonParseSelfInject(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	jp := d.JsonPtr()
	d2 := coredynamic.NewDynamicValid(0)
	err := d2.JsonParseSelfInject(jp)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error from JsonParseSelfInject with untyped destination", actual)
}

func Test_C34_83_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	c := d.Clone()
	actual := args.Map{"result": c.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_84_Dynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	cp := d.ClonePtr()
	actual := args.Map{"result": cp == nil || cp.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned ptr", actual)
}

func Test_C34_85_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C34_86_Dynamic_NonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	np := d.NonPtr()
	actual := args.Map{"result": np.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_87_Dynamic_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	p := d.Ptr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// Collection[T] — base, Lock, Search, Distinct, GroupBy, Map
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_88_CollectionFrom(t *testing.T) {
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_89_CollectionFrom_Nil(t *testing.T) {
	c := coredynamic.CollectionFrom[int](nil)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_90_CollectionClone(t *testing.T) {
	c := coredynamic.CollectionClone([]string{"a", "b"})
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_91_Collection_FirstOrDefault_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	v, ok := c.FirstOrDefault()
	actual := args.Map{"result": ok || v != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil/false", actual)
}

func Test_C34_92_Collection_FirstOrDefault_NonEmpty(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	v, ok := c.FirstOrDefault()
	actual := args.Map{"result": ok || *v != 42}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42/true", actual)
}

func Test_C34_93_Collection_LastOrDefault_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	v, ok := c.LastOrDefault()
	actual := args.Map{"result": ok || v != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil/false", actual)
}

func Test_C34_94_Collection_LastOrDefault_NonEmpty(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	v, ok := c.LastOrDefault()
	actual := args.Map{"result": ok || *v != 2}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 2/true", actual)
}

func Test_C34_95_Collection_Items_Nil(t *testing.T) {
	var c *coredynamic.Collection[int]
	items := c.Items()
	actual := args.Map{"result": len(items) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_96_Collection_Count(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": c.Count() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_97_Collection_HasAnyItem(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	c.Add(1)
	actual := args.Map{"result": c.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_98_Collection_HasIndex(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": c.HasIndex(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": c.HasIndex(2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": c.HasIndex(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for -1", actual)
}

func Test_C34_99_Collection_Skip(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	s := c.Skip(1)
	actual := args.Map{"result": len(s) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_100_Collection_Take(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	s := c.Take(2)
	actual := args.Map{"result": len(s) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_101_Collection_Limit(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": len(c.Limit(1)) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_102_Collection_SkipCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	sc := c.SkipCollection(2)
	actual := args.Map{"result": sc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_103_Collection_TakeCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	tc := c.TakeCollection(2)
	actual := args.Map{"result": tc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_104_Collection_LimitCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	lc := c.LimitCollection(1)
	actual := args.Map{"result": lc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_105_Collection_SafeLimitCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	lc := c.SafeLimitCollection(100)
	actual := args.Map{"result": lc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_106_Collection_AddMany(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddMany(1, 2, 3)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_107_Collection_AddNonNil(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	val := 42
	c.AddNonNil(&val)
	c.AddNonNil(nil)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_108_Collection_RemoveAt_Valid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	ok := c.RemoveAt(1)
	actual := args.Map{"result": ok || c.Length() != 2}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_C34_109_Collection_RemoveAt_Invalid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	actual := args.Map{"result": c.RemoveAt(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_110_Collection_Clear(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	c.Clear()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_111_Collection_Dispose(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.Dispose()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_112_Collection_Loop(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	sum := 0
	c.Loop(func(i int, v int) bool {
		sum += v
		return false
	})
	actual := args.Map{"result": sum != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
}

func Test_C34_113_Collection_Loop_Break(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	count := 0
	c.Loop(func(i int, v int) bool {
		count++
		return i == 1
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_114_Collection_Loop_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.Loop(func(i int, v int) bool {
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
		return false
	})
}

func Test_C34_115_Collection_LoopAsync(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	var mu sync.Mutex
	sum := 0
	c.LoopAsync(func(i int, v int) {
		mu.Lock()
		sum += v
		mu.Unlock()
	})
	actual := args.Map{"result": sum != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
}

func Test_C34_116_Collection_LoopAsync_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.LoopAsync(func(i int, v int) {
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
	})
}

func Test_C34_117_Collection_Filter(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3).Add(4)
	evens := c.Filter(func(v int) bool { return v%2 == 0 })
	actual := args.Map{"result": evens.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_118_Collection_Filter_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	result := c.Filter(func(v int) bool { return true })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_119_Collection_GetPagesSize(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	for i := 0; i < 10; i++ {
		c.Add(i)
	}
	actual := args.Map{"result": c.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C34_120_Collection_GetPagesSize_Zero(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	actual := args.Map{"result": c.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_121_Collection_GetSinglePageCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	for i := 0; i < 10; i++ {
		c.Add(i)
	}
	page := c.GetSinglePageCollection(3, 1)
	actual := args.Map{"result": page.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_122_Collection_GetSinglePageCollection_Small(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	page := c.GetSinglePageCollection(10, 1)
	actual := args.Map{"result": page.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
}

func Test_C34_123_Collection_GetPagedCollection(t *testing.T) {
	c := coredynamic.NewCollection[int](10)
	for i := 0; i < 10; i++ {
		c.Add(i)
	}
	pages := c.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C34_124_Collection_GetPagedCollection_Small(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	pages := c.GetPagedCollection(10)
	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_125_Collection_MarshalJSON(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	b, err := c.MarshalJSON()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_126_Collection_UnmarshalJSON(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	err := c.UnmarshalJSON([]byte(`[1,2,3]`))
	actual := args.Map{"result": err != nil || c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 items", actual)
}

func Test_C34_127_Collection_JsonString(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	s, err := c.JsonString()
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_C34_128_Collection_JsonStringMust(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	s := c.JsonStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C34_129_Collection_Strings(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	strs := c.Strings()
	actual := args.Map{"result": len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_130_Collection_String(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	s := c.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionLock — thread-safe variants
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_131_Collection_LengthLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": c.LengthLock() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_132_Collection_IsEmptyLock(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": c.IsEmptyLock()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_133_Collection_AddLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddLock(42)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_134_Collection_AddsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddsLock(1, 2, 3)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_135_Collection_AddManyLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyLock(1, 2)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_136_Collection_AddCollectionLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c2 := coredynamic.NewCollection[int](4)
	c2.Add(2).Add(3)
	c.AddCollectionLock(c2)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_137_Collection_AddCollectionLock_Nil(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.AddCollectionLock(nil)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_138_Collection_AddCollectionsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c2 := coredynamic.NewCollection[int](4)
	c2.Add(1)
	c3 := coredynamic.NewCollection[int](4)
	c3.Add(2)
	c.AddCollectionsLock(c2, nil, c3)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_139_Collection_AddIfLock_True(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIfLock(true, 42)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_140_Collection_AddIfLock_False(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddIfLock(false, 42)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_141_Collection_RemoveAtLock_Valid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": c.RemoveAtLock(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_142_Collection_RemoveAtLock_Invalid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	actual := args.Map{"result": c.RemoveAtLock(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_143_Collection_ClearLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	c.ClearLock()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_144_Collection_ItemsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	items := c.ItemsLock()
	actual := args.Map{"result": len(items) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_145_Collection_FirstLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	actual := args.Map{"result": c.FirstLock() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_146_Collection_LastLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": c.LastLock() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_147_Collection_AddWithWgLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, 42)
	wg.Wait()
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_148_Collection_LoopLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	sum := 0
	c.LoopLock(func(i int, v int) bool {
		sum += v
		return false
	})
	actual := args.Map{"result": sum != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_149_Collection_LoopLock_Break(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	count := 0
	c.LoopLock(func(i int, v int) bool {
		count++
		return true
	})
	actual := args.Map{"result": count != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_150_Collection_FilterLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	result := c.FilterLock(func(v int) bool { return v > 1 })
	actual := args.Map{"result": result.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_151_Collection_StringsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	strs := c.StringsLock()
	actual := args.Map{"result": len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSearch — Contains, IndexOf, Has, HasAll, LastIndexOf, Count
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_152_Contains(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	actual := args.Map{"result": coredynamic.Contains(c, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": coredynamic.Contains(c, 99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_153_IndexOf(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("b").Add("c")
	actual := args.Map{"result": coredynamic.IndexOf(c, "b") != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": coredynamic.IndexOf(c, "z") != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C34_154_Has(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	actual := args.Map{"result": coredynamic.Has(c, 42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_155_HasAll(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	actual := args.Map{"result": coredynamic.HasAll(c, 1, 3)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": coredynamic.HasAll(c, 1, 99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C34_156_HasAll_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": coredynamic.HasAll(c, 1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
}

func Test_C34_157_LastIndexOf(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1)
	actual := args.Map{"result": coredynamic.LastIndexOf(c, 1) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": coredynamic.LastIndexOf(c, 99) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C34_158_Count(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1)
	actual := args.Map{"result": coredynamic.Count(c, 1) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_159_ContainsLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	actual := args.Map{"result": coredynamic.ContainsLock(c, 42)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_160_IndexOfLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	actual := args.Map{"result": coredynamic.IndexOfLock(c, 2) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionDistinct
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_161_Distinct(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1).Add(3).Add(2)
	d := coredynamic.Distinct(c)
	actual := args.Map{"result": d.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_162_Distinct_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	d := coredynamic.Distinct(c)
	actual := args.Map{"result": d.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_163_Unique(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("b").Add("a")
	u := coredynamic.Unique(c)
	actual := args.Map{"result": u.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_164_DistinctLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(1).Add(2)
	d := coredynamic.DistinctLock(c)
	actual := args.Map{"result": d.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_165_DistinctCount(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(1)
	actual := args.Map{"result": coredynamic.DistinctCount(c) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_166_DistinctCount_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	actual := args.Map{"result": coredynamic.DistinctCount(c) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_167_IsDistinct_True(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	actual := args.Map{"result": coredynamic.IsDistinct(c)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C34_168_IsDistinct_False(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(1)
	actual := args.Map{"result": coredynamic.IsDistinct(c)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionGroupBy
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_169_GroupBy(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3).Add(4)
	groups := coredynamic.GroupBy(c, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})
	actual := args.Map{"result": len(groups) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 groups", actual)
	actual := args.Map{"result": groups["even"].Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 evens", actual)
}

func Test_C34_170_GroupBy_Nil(t *testing.T) {
	groups := coredynamic.GroupBy[int, string](nil, func(v int) string { return "" })
	actual := args.Map{"result": len(groups) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_171_GroupBy_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	groups := coredynamic.GroupBy(c, func(v int) string { return "" })
	actual := args.Map{"result": len(groups) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_172_GroupByLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	groups := coredynamic.GroupByLock(c, func(v int) int { return v % 2 })
	actual := args.Map{"result": len(groups) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 groups", actual)
}

func Test_C34_173_GroupByCount(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("b").Add("a")
	counts := coredynamic.GroupByCount(c, func(s string) string { return s })
	actual := args.Map{"result": counts["a"] != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_174_GroupByCount_Nil(t *testing.T) {
	counts := coredynamic.GroupByCount[string, string](nil, func(s string) string { return s })
	actual := args.Map{"result": len(counts) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionMap — Map, FlatMap, Reduce
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_175_Map(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	result := coredynamic.Map(c, func(v int) int { return v * 2 })
	actual := args.Map{"result": result.Length() != 3 || result.At(0) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected doubled", actual)
}

func Test_C34_176_Map_Nil(t *testing.T) {
	result := coredynamic.Map[int, int](nil, func(v int) int { return v })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_177_Map_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	result := coredynamic.Map(c, func(v int) int { return v })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_178_FlatMap(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a,b").Add("c,d")
	result := coredynamic.FlatMap(c, func(s string) []string {
		return []string{s + "1", s + "2"}
	})
	actual := args.Map{"result": result.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C34_179_FlatMap_Nil(t *testing.T) {
	result := coredynamic.FlatMap[int, int](nil, func(v int) []int { return nil })
	actual := args.Map{"result": result.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C34_180_Reduce(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	sum := coredynamic.Reduce(c, 0, func(acc int, v int) int { return acc + v })
	actual := args.Map{"result": sum != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6", actual)
}

func Test_C34_181_Reduce_Nil(t *testing.T) {
	result := coredynamic.Reduce[int, int](nil, 42, func(acc int, v int) int { return acc + v })
	actual := args.Map{"result": result != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected initial value", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// AnyCollection — extended methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C34_182_AnyCollection_AtAsDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(42)
	d := ac.AtAsDynamic(0)
	actual := args.Map{"result": d.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C34_183_AnyCollection_DynamicItems(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	items := ac.DynamicItems()
	actual := args.Map{"result": len(items) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_184_AnyCollection_DynamicItems_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	items := ac.DynamicItems()
	actual := args.Map{"result": len(items) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_185_AnyCollection_DynamicCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	dc := ac.DynamicCollection()
	actual := args.Map{"result": dc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_186_AnyCollection_DynamicCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	dc := ac.DynamicCollection()
	actual := args.Map{"result": dc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_187_AnyCollection_ReflectSetAt(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(42)
	var target int
	err := ac.ReflectSetAt(0, &target)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_C34_188_AnyCollection_ListStrings(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("hello").Add(42)
	strs := ac.ListStrings(false)
	actual := args.Map{"result": len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_189_AnyCollection_ListStringsPtr(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("x")
	strs := ac.ListStringsPtr(true)
	actual := args.Map{"result": len(strs) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_190_AnyCollection_Loop_Sync(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_191_AnyCollection_Loop_Sync_Break(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return true
	})
	actual := args.Map{"result": count != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_192_AnyCollection_Loop_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	ac.Loop(true, func(i int, item any) bool {
		return false
	})
	// just ensure no panic
}

func Test_C34_193_AnyCollection_Loop_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.Loop(false, func(i int, item any) bool {
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
		return false
	})
}

func Test_C34_194_AnyCollection_LoopDynamic_Sync(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	count := 0
	ac.LoopDynamic(false, func(i int, d coredynamic.Dynamic) bool {
		count++
		return false
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_195_AnyCollection_LoopDynamic_Sync_Break(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	count := 0
	ac.LoopDynamic(false, func(i int, d coredynamic.Dynamic) bool {
		count++
		return true
	})
	actual := args.Map{"result": count != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_196_AnyCollection_LoopDynamic_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	ac.LoopDynamic(true, func(i int, d coredynamic.Dynamic) bool {
		return false
	})
}

func Test_C34_197_AnyCollection_LoopDynamic_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	ac.LoopDynamic(false, func(i int, d coredynamic.Dynamic) bool {
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
		return false
	})
}

func Test_C34_198_AnyCollection_AddAny(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAny(42, true)
	actual := args.Map{"result": ac.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_199_AnyCollection_AddNonNull(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddNonNull(nil)
	ac.AddNonNull(42)
	actual := args.Map{"result": ac.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_200_AnyCollection_AddNonNullDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddNonNullDynamic(nil, true)
	ac.AddNonNullDynamic(42, true)
	actual := args.Map{"result": ac.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_201_AnyCollection_AddAnyManyDynamic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnyManyDynamic(1, 2, 3)
	actual := args.Map{"result": ac.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_202_AnyCollection_AddAnyManyDynamic_Nil(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnyManyDynamic(nil)
	// nil variadic should skip
}

func Test_C34_203_AnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnySliceFromSingleItem([]int{1, 2, 3})
	actual := args.Map{"result": ac.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_204_AnyCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnySliceFromSingleItem(nil)
	actual := args.Map{"result": ac.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_205_AnyCollection_AddMany(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddMany(1, nil, 3)
	actual := args.Map{"result": ac.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
}

func Test_C34_206_AnyCollection_AddMany_NilVariadic(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.AddMany(nil)
	// nil variadic should skip
}

func Test_C34_207_AnyCollection_AddAnyWithTypeValidation_Valid(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(0), 42)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_C34_208_AnyCollection_AddAnyWithTypeValidation_Invalid(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(0), "str")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_209_AnyCollection_AddAnyItemsWithTypeValidation_Continue(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(0), 1, "bad", 3)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_210_AnyCollection_AddAnyItemsWithTypeValidation_Stop(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0), 1, "bad", 3)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_211_AnyCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0))
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C34_212_AnyCollection_Paging(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	actual := args.Map{"result": ac.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	actual := args.Map{"result": ac.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for zero page size", actual)
}

func Test_C34_213_AnyCollection_GetSinglePageCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	page := ac.GetSinglePageCollection(3, 1)
	actual := args.Map{"result": page.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C34_214_AnyCollection_GetSinglePageCollection_Small(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	page := ac.GetSinglePageCollection(10, 1)
	actual := args.Map{"result": page.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
}

func Test_C34_215_AnyCollection_GetPagedCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	pages := ac.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pages", actual)
}

func Test_C34_216_AnyCollection_GetPagedCollection_Small(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	pages := ac.GetPagedCollection(10)
	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_217_AnyCollection_Json(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	j := ac.Json()
	_ = j
}

func Test_C34_218_AnyCollection_JsonPtr(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	actual := args.Map{"result": jp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_219_AnyCollection_JsonModel(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	m := ac.JsonModel()
	actual := args.Map{"result": len(m) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C34_220_AnyCollection_JsonModelAny(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	v := ac.JsonModelAny()
	actual := args.Map{"result": v == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_221_AnyCollection_MarshalJSON(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(42)
	b, err := ac.MarshalJSON()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C34_222_AnyCollection_UnmarshalJSON(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.UnmarshalJSON([]byte(`[1,2,3]`))
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
}

func Test_C34_223_AnyCollection_UnmarshalJSON_Bad(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	err := ac.UnmarshalJSON([]byte(`not json`))
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_224_AnyCollection_JsonResultsCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	rc := ac.JsonResultsCollection()
	actual := args.Map{"result": rc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_225_AnyCollection_JsonResultsCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	rc := ac.JsonResultsCollection()
	actual := args.Map{"result": rc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_226_AnyCollection_JsonResultsPtrCollection(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	rc := ac.JsonResultsPtrCollection()
	actual := args.Map{"result": rc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_227_AnyCollection_JsonResultsPtrCollection_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	rc := ac.JsonResultsPtrCollection()
	actual := args.Map{"result": rc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C34_228_AnyCollection_ParseInjectUsingJson(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	ac2 := coredynamic.NewAnyCollection(4)
	_, err := ac2.ParseInjectUsingJson(jp)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for AnyCollection JSON payload {}", actual)
}

func Test_C34_229_AnyCollection_ParseInjectUsingJson_Bad(t *testing.T) {
	badJson := corejson.NewPtr("not an any collection")
	ac := coredynamic.NewAnyCollection(4)
	_, err := ac.ParseInjectUsingJson(badJson)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C34_230_AnyCollection_ParseInjectUsingJsonMust(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	ac2 := coredynamic.NewAnyCollection(4)
	panicked := false

	func() {
		defer func() {
			if recover() != nil {
				panicked = true
			}
		}()

		_ = ac2.ParseInjectUsingJsonMust(jp)
	}()

	actual := args.Map{"result": panicked}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected panic for AnyCollection JSON payload {}", actual)
}

func Test_C34_231_AnyCollection_JsonParseSelfInject(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	jp := ac.JsonPtr()
	ac2 := coredynamic.NewAnyCollection(4)
	err := ac2.JsonParseSelfInject(jp)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for AnyCollection JSON payload {}", actual)
}

func Test_C34_232_AnyCollection_Strings(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add("x")
	strs := ac.Strings()
	actual := args.Map{"result": len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C34_233_AnyCollection_Strings_Empty(t *testing.T) {
	ac := coredynamic.EmptyAnyCollection()
	strs := ac.Strings()
	actual := args.Map{"result": len(strs) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C34_234_AnyCollection_String(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1)
	s := ac.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C34_235_AnyCollection_GetPagingInfo(t *testing.T) {
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	info := ac.GetPagingInfo(3, 1)
	actual := args.Map{"result": info.TotalPages != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 total pages", actual)
}
