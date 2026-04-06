package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// newCreator factories — New.Collection.String/Int/Byte/Any etc.
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_01_NewCreator_String_Empty(t *testing.T) {
	c := coredynamic.New.Collection.String.Empty()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_02_NewCreator_String_Cap(t *testing.T) {
	c := coredynamic.New.Collection.String.Cap(10)
	actual := args.Map{"result": c.Capacity() < 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cap >= 10", actual)
}

func Test_C35_03_NewCreator_String_From(t *testing.T) {
	c := coredynamic.New.Collection.String.From([]string{"a", "b"})
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_04_NewCreator_String_Clone(t *testing.T) {
	c := coredynamic.New.Collection.String.Clone([]string{"a", "b"})
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_05_NewCreator_String_Items(t *testing.T) {
	c := coredynamic.New.Collection.String.Items("a", "b", "c")
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C35_06_NewCreator_String_Create(t *testing.T) {
	c := coredynamic.New.Collection.String.Create([]string{"x"})
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_07_NewCreator_String_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.String.LenCap(3, 10)
	actual := args.Map{"result": c.Length() != 3 || c.Capacity() < 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len=3 cap>=10", actual)
}

func Test_C35_08_NewCreator_Int_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int.Empty()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_09_NewCreator_Int_Cap(t *testing.T) {
	c := coredynamic.New.Collection.Int.Cap(5)
	c.Add(42)
	actual := args.Map{"result": c.At(0) != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C35_10_NewCreator_Int_Items(t *testing.T) {
	c := coredynamic.New.Collection.Int.Items(1, 2, 3)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C35_11_NewCreator_Int_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int.LenCap(2, 10)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_12_NewCreator_Int64_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Int64.Empty()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_13_NewCreator_Int64_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Int64.LenCap(1, 5)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_14_NewCreator_Byte_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Byte.Empty()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_15_NewCreator_Byte_LenCap(t *testing.T) {
	c := coredynamic.New.Collection.Byte.LenCap(2, 8)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_16_NewCreator_Any_Empty(t *testing.T) {
	c := coredynamic.New.Collection.Any.Empty()
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_17_NewCreator_Any_Items(t *testing.T) {
	c := coredynamic.New.Collection.Any.Items(1, "two", 3.0)
	actual := args.Map{"result": c.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// KeyVal — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_18_KeyVal_KeyDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.KeyDynamic()
	actual := args.Map{"result": d.Value() != "k"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected k", actual)
}

func Test_C35_19_KeyVal_ValueDynamic(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	d := kv.ValueDynamic()
	actual := args.Map{"result": d.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C35_20_KeyVal_KeyDynamicPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	dp := kv.KeyDynamicPtr()
	actual := args.Map{"result": dp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C35_21_KeyVal_KeyDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"result": kv.KeyDynamicPtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_22_KeyVal_ValueDynamicPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	dp := kv.ValueDynamicPtr()
	actual := args.Map{"result": dp == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C35_23_KeyVal_ValueDynamicPtr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"result": kv.ValueDynamicPtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_24_KeyVal_IsKeyNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: 1}
	actual := args.Map{"result": kv.IsKeyNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_25_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "", Value: 1}
	actual := args.Map{"result": kv.IsKeyNullOrEmptyString()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_26_KeyVal_IsValueNull(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"result": kv.IsValueNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_27_KeyVal_String(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	s := kv.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_28_KeyVal_String_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"result": kv.String() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_29_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	actual := args.Map{"result": rv.Kind() != reflect.Int}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected int", actual)
}

func Test_C35_30_KeyVal_ValueInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"result": kv.ValueInt() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C35_31_KeyVal_ValueInt_Wrong(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}
	v := kv.ValueInt()
	_ = v // just no panic
}

func Test_C35_32_KeyVal_ValueUInt(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: uint(10)}
	actual := args.Map{"result": kv.ValueUInt() != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
}

func Test_C35_33_KeyVal_ValueStrings(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}
	actual := args.Map{"result": len(kv.ValueStrings()) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_34_KeyVal_ValueStrings_Wrong(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{"result": kv.ValueStrings() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_35_KeyVal_ValueBool(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: true}
	actual := args.Map{"result": kv.ValueBool()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_36_KeyVal_ValueInt64(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: int64(100)}
	actual := args.Map{"result": kv.ValueInt64() != 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_C35_37_KeyVal_CastKeyVal_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.CastKeyVal(nil, nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C35_38_KeyVal_ReflectSetKey_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ReflectSetKey(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_39_KeyVal_ValueNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueNullErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_40_KeyVal_ValueNullErr_NullValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: nil}
	err := kv.ValueNullErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_41_KeyVal_ValueNullErr_Valid(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	err := kv.ValueNullErr()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_42_KeyVal_KeyNullErr_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyNullErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_43_KeyVal_KeyNullErr_NullKey(t *testing.T) {
	kv := coredynamic.KeyVal{Key: nil, Value: 42}
	err := kv.KeyNullErr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_44_KeyVal_KeyString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "hello", Value: 1}
	actual := args.Map{"result": kv.KeyString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_45_KeyVal_KeyString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"result": kv.KeyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_46_KeyVal_ValueString(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "world"}
	actual := args.Map{"result": kv.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_47_KeyVal_ValueString_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	actual := args.Map{"result": kv.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_48_KeyVal_KeyReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.KeyReflectSet(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_49_KeyVal_ValueReflectSet_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ValueReflectSet(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_50_KeyVal_ReflectSetTo_Nil(t *testing.T) {
	var kv *coredynamic.KeyVal
	err := kv.ReflectSetTo(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_51_KeyVal_Json(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	_ = j
}

func Test_C35_52_KeyVal_JsonPtr(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"result": kv.JsonPtr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C35_53_KeyVal_JsonModel(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"result": kv.JsonModel() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C35_54_KeyVal_Serialize(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C35_55_KeyVal_ParseInjectUsingJson(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jp := kv.JsonPtr()
	kv2 := &coredynamic.KeyVal{}
	_, err := kv2.ParseInjectUsingJson(jp)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_C35_56_KeyVal_JsonParseSelfInject(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jp := kv.JsonPtr()
	kv2 := &coredynamic.KeyVal{}
	err := kv2.JsonParseSelfInject(jp)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValCollection — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_57_KeyValCollection_Basic(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	actual := args.Map{"result": kvc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": kvc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_58_KeyValCollection_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"result": kvc.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_59_KeyValCollection_AddPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	kvc.AddPtr(kv)
	kvc.AddPtr(nil)
	actual := args.Map{"result": kvc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_60_KeyValCollection_AddMany(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)
	actual := args.Map{"result": kvc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_61_KeyValCollection_AddManyPtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	k1 := &coredynamic.KeyVal{Key: "a", Value: 1}
	kvc.AddManyPtr(k1, nil)
	actual := args.Map{"result": kvc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_62_KeyValCollection_Items_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	actual := args.Map{"result": kvc.Items() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_63_KeyValCollection_MapAnyItems(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	m := kvc.MapAnyItems()
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_64_KeyValCollection_MapAnyItems_Empty(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	m := kvc.MapAnyItems()
	actual := args.Map{"result": m.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_65_KeyValCollection_AllKeys(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := kvc.AllKeys()
	actual := args.Map{"result": len(keys) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_66_KeyValCollection_AllKeysSorted(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 2})
	keys := kvc.AllKeysSorted()
	actual := args.Map{"result": len(keys) != 2 || keys[0] > keys[1]}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_C35_67_KeyValCollection_AllValues(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	vals := kvc.AllValues()
	actual := args.Map{"result": len(vals) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_68_KeyValCollection_String(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_69_KeyValCollection_String_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	actual := args.Map{"result": kvc.String() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_70_KeyValCollection_Json(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	j := kvc.Json()
	_ = j
}

func Test_C35_71_KeyValCollection_Serialize(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	b, err := kvc.Serialize()
	actual := args.Map{"result": err != nil || len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_C35_72_KeyValCollection_JsonString(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s, err := kvc.JsonString()
	actual := args.Map{"result": err != nil || s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_C35_73_KeyValCollection_JsonStringMust(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.JsonStringMust()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_74_KeyValCollection_Clone(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	cloned := kvc.Clone()
	actual := args.Map{"result": cloned.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_75_KeyValCollection_ClonePtr(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	cp := kvc.ClonePtr()
	actual := args.Map{"result": cp == nil || cp.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned", actual)
}

func Test_C35_76_KeyValCollection_ClonePtr_Nil(t *testing.T) {
	var kvc *coredynamic.KeyValCollection
	actual := args.Map{"result": kvc.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_77_KeyValCollection_Paging(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	actual := args.Map{"result": kvc.GetPagesSize(3) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
	actual := args.Map{"result": kvc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_78_KeyValCollection_GetSinglePageCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	page := kvc.GetSinglePageCollection(3, 1)
	actual := args.Map{"result": page.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C35_79_KeyValCollection_GetPagedCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(3)
	actual := args.Map{"result": len(pages) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C35_80_KeyValCollection_GetPagedCollection_Small(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 1})
	pages := kvc.GetPagedCollection(10)
	actual := args.Map{"result": len(pages) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_81_KeyValCollection_JsonResultsCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	rc := kvc.JsonResultsCollection()
	actual := args.Map{"result": rc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C35_82_KeyValCollection_JsonResultsPtrCollection(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	rc := kvc.JsonResultsPtrCollection()
	actual := args.Map{"result": rc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C35_83_KeyValCollection_JsonMapResults(t *testing.T) {
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: 42})
	mr, err := kvc.JsonMapResults()
	actual := args.Map{"result": err != nil || mr == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItems — Add, Get, Deserialize, validation, paging
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_84_MapAnyItems_Add(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.Add("key", 42)
	actual := args.Map{"result": isNew}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
	isNew2 := m.Add("key", 99)
	actual := args.Map{"result": isNew2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not new", actual)
}

func Test_C35_85_MapAnyItems_Set(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Set("k", 1)
	actual := args.Map{"result": m.GetValue("k") != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_86_MapAnyItems_HasKey(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 1)
	actual := args.Map{"result": m.HasKey("k")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": m.HasKey("nope")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C35_87_MapAnyItems_HasKey_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItems
	actual := args.Map{"result": m.HasKey("k")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C35_88_MapAnyItems_Get(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	v, has := m.Get("k")
	actual := args.Map{"result": has || v != 42}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_, has2 := m.Get("nope")
	actual := args.Map{"result": has2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C35_89_MapAnyItems_GetValue(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "hello")
	actual := args.Map{"result": m.GetValue("k") != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
	actual := args.Map{"result": m.GetValue("nope") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_90_MapAnyItems_ReflectSetTo(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	err := m.ReflectSetTo("k", &target)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_C35_91_MapAnyItems_ReflectSetTo_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	var target int
	err := m.ReflectSetTo("nope", &target)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_92_MapAnyItems_Deserialize(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", 42)
	var target int
	err := m.Deserialize("k", &target)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual := args.Map{"result": target != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C35_93_MapAnyItems_Deserialize_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	var target int
	err := m.Deserialize("nope", &target)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_94_MapAnyItems_AddKeyAny(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: 42})
	actual := args.Map{"result": isNew}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
}

func Test_C35_95_MapAnyItems_AddKeyAnyWithValidation_Valid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddKeyAnyWithValidation(reflect.TypeOf(0), corejson.KeyAny{Key: "k", AnyInf: 42})
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_C35_96_MapAnyItems_AddKeyAnyWithValidation_Invalid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddKeyAnyWithValidation(reflect.TypeOf(0), corejson.KeyAny{Key: "k", AnyInf: "str"})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_97_MapAnyItems_AddWithValidation_Valid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddWithValidation(reflect.TypeOf(0), "k", 42)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
}

func Test_C35_98_MapAnyItems_AddWithValidation_Invalid(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	err := m.AddWithValidation(reflect.TypeOf(0), "k", "str")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_99_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	jr := corejson.NewPtr(42)
	m.AddJsonResultPtr("k", jr)
	actual := args.Map{"result": m.HasKey("k")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected key", actual)
}

func Test_C35_100_MapAnyItems_AddJsonResultPtr_Nil(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.AddJsonResultPtr("k", nil)
	actual := args.Map{"result": m.HasKey("k")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no key", actual)
}

func Test_C35_101_MapAnyItems_GetPagesSize(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	actual := args.Map{"result": m.GetPagesSize(2) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 pages", actual)
	actual := args.Map{"result": m.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_102_MapAnyItems_GetFieldsMap(t *testing.T) {
	inner := map[string]any{"x": 1}
	b, _ := json.Marshal(inner)
	var stored any
	json.Unmarshal(b, &stored)
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", stored)
	fm, err, found := m.GetFieldsMap("k")
	actual := args.Map{"result": found}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found", actual)
	_ = fm
	_ = err
}

func Test_C35_103_MapAnyItems_GetFieldsMap_Missing(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	_, _, found := m.GetFieldsMap("nope")
	actual := args.Map{"result": found}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C35_104_MapAnyItems_GetSafeFieldsMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	_, found := m.GetSafeFieldsMap("nope")
	actual := args.Map{"result": found}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C35_105_NewMapAnyItemsUsingItems(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": 1})
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_106_NewMapAnyItemsUsingItems_Empty(t *testing.T) {
	m := coredynamic.NewMapAnyItemsUsingItems(nil)
	actual := args.Map{"result": m.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_107_NewMapAnyItemsUsingAnyTypeMap(t *testing.T) {
	_, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

func Test_C35_108_NewMapAnyItemsUsingAnyTypeMap_Valid(t *testing.T) {
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]any{"k": 1})
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItemDiff
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_109_MapAnyItemDiff_Length(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	actual := args.Map{"result": d.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C35_110_MapAnyItemDiff_Length_Nil(t *testing.T) {
	var d *coredynamic.MapAnyItemDiff
	actual := args.Map{"result": d.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_111_MapAnyItemDiff_IsEmpty(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{}
	actual := args.Map{"result": d.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_112_MapAnyItemDiff_HasAnyItem(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{"result": d.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_113_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"b": 1, "a": 2}
	keys := d.AllKeysSorted()
	actual := args.Map{"result": len(keys) != 2 || keys[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted keys", actual)
}

func Test_C35_114_MapAnyItemDiff_Raw(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	r := d.Raw()
	actual := args.Map{"result": len(r) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_115_MapAnyItemDiff_Raw_Nil(t *testing.T) {
	var d *coredynamic.MapAnyItemDiff
	actual := args.Map{"result": len(d.Raw()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_116_MapAnyItemDiff_MapAnyItems(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	m := d.MapAnyItems()
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_117_MapAnyItemDiff_RawMapDiffer(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	rd := d.RawMapDiffer()
	actual := args.Map{"result": len(rd) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C35_118_MapAnyItemDiff_IsRawEqual(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	actual := args.Map{"result": d.IsRawEqual(false, map[string]any{"k": 1})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_C35_119_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	actual := args.Map{"result": d.HasAnyChanges(false, map[string]any{"k": 1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no changes", actual)
	actual := args.Map{"result": d.HasAnyChanges(false, map[string]any{"k": 2})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected changes", actual)
}

func Test_C35_120_MapAnyItemDiff_Clear(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	cleared := d.Clear()
	actual := args.Map{"result": len(cleared) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_121_MapAnyItemDiff_Clear_Nil(t *testing.T) {
	var d *coredynamic.MapAnyItemDiff
	cleared := d.Clear()
	actual := args.Map{"result": len(cleared) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C35_122_MapAnyItemDiff_Json(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	j := d.Json()
	_ = j
}

func Test_C35_123_MapAnyItemDiff_PrettyJsonString(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	s := d.PrettyJsonString()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_124_MapAnyItemDiff_LogPrettyJsonString(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{"k": 1}
	d.LogPrettyJsonString() // just no panic
}

func Test_C35_125_MapAnyItemDiff_LogPrettyJsonString_Empty(t *testing.T) {
	d := coredynamic.MapAnyItemDiff{}
	d.LogPrettyJsonString() // no panic
}

// ═══════════════════════════════════════════════════════════════════════
// TypeStatus — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_126_TypeStatus_IsValid(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

func Test_C35_127_TypeStatus_IsValid_Nil(t *testing.T) {
	var st *coredynamic.TypeStatus
	actual := args.Map{"result": st.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C35_128_TypeStatus_IsInvalid(t *testing.T) {
	var st *coredynamic.TypeStatus
	actual := args.Map{"result": st.IsInvalid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_129_TypeStatus_IsNotSame(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "hello")
	actual := args.Map{"result": st.IsNotSame()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected not same", actual)
}

func Test_C35_130_TypeStatus_IsNotEqualTypes(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	actual := args.Map{"result": st.IsNotEqualTypes()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_131_TypeStatus_IsAnyPointer(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(&v, 42)
	actual := args.Map{"result": st.IsAnyPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_132_TypeStatus_IsBothPointer(t *testing.T) {
	v1, v2 := 42, 100
	st := coredynamic.TypeSameStatus(&v1, &v2)
	actual := args.Map{"result": st.IsBothPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_133_TypeStatus_NonPointerLeft(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(&v, 100)
	npl := st.NonPointerLeft()
	actual := args.Map{"result": npl.Kind() != reflect.Int}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected int", actual)
}

func Test_C35_134_TypeStatus_NonPointerRight(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(100, &v)
	npr := st.NonPointerRight()
	actual := args.Map{"result": npr.Kind() != reflect.Int}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected int", actual)
}

func Test_C35_135_TypeStatus_IsSameRegardlessPointer(t *testing.T) {
	v := 42
	st := coredynamic.TypeSameStatus(&v, 100)
	actual := args.Map{"result": st.IsSameRegardlessPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_136_TypeStatus_LeftName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.LeftName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_137_TypeStatus_LeftName_Nil(t *testing.T) {
	st := coredynamic.TypeSameStatus(nil, 42)
	n := st.LeftName()
	actual := args.Map{"result": n == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected something", actual)
}

func Test_C35_138_TypeStatus_RightName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.RightName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_139_TypeStatus_LeftFullName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.LeftFullName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_140_TypeStatus_RightFullName(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.RightFullName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C35_141_TypeStatus_NotMatchMessage_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.NotMatchMessage("l", "r") != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for same", actual)
}

func Test_C35_142_TypeStatus_NotMatchMessage_Different(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	msg := st.NotMatchMessage("l", "r")
	actual := args.Map{"result": msg == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C35_143_TypeStatus_NotMatchErr_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.NotMatchErr("l", "r") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_144_TypeStatus_NotMatchErr_Different(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	actual := args.Map{"result": st.NotMatchErr("l", "r") == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_145_TypeStatus_MustBeSame_Same(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not panic", actual)
	}()
	st := coredynamic.TypeSameStatus(42, 100)
	st.MustBeSame()
}

func Test_C35_146_TypeStatus_MustBeSame_Different(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()
	st := coredynamic.TypeSameStatus(42, "x")
	st.MustBeSame()
}

func Test_C35_147_TypeStatus_ValidationError(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	actual := args.Map{"result": st.ValidationError() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_148_TypeStatus_ValidationError_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.ValidationError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_149_TypeStatus_NotEqualSrcDestinationMessage(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	msg := st.NotEqualSrcDestinationMessage()
	actual := args.Map{"result": msg == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C35_150_TypeStatus_NotEqualSrcDestinationErr(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "x")
	actual := args.Map{"result": st.NotEqualSrcDestinationErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_151_TypeStatus_SrcDestinationMustBeSame_Same(t *testing.T) {
	defer func() {
		actual := args.Map{"result": r := recover(); r != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not panic", actual)
	}()
	st := coredynamic.TypeSameStatus(42, 100)
	st.SrcDestinationMustBeSame()
}

func Test_C35_152_TypeStatus_IsEqual_BothNil(t *testing.T) {
	var a, b *coredynamic.TypeStatus
	actual := args.Map{"result": a.IsEqual(b)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_153_TypeStatus_IsEqual_OneNil(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st.IsEqual(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C35_154_TypeStatus_IsEqual_Same(t *testing.T) {
	st1 := coredynamic.TypeSameStatus(42, 100)
	st2 := coredynamic.TypeSameStatus(42, 100)
	actual := args.Map{"result": st1.IsEqual(&st2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_155_TypeStatus_IsEqual_Different(t *testing.T) {
	st1 := coredynamic.TypeSameStatus(42, 100)
	st2 := coredynamic.TypeSameStatus(42, "x")
	actual := args.Map{"result": st1.IsEqual(&st2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicStatus, ValueStatus, CastedResult
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_156_DynamicStatus_Invalid(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	actual := args.Map{"result": ds.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": ds.Message != "msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected msg", actual)
}

func Test_C35_157_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	actual := args.Map{"result": ds.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_C35_158_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	cloned := ds.Clone()
	actual := args.Map{"result": cloned.Message != "msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected msg", actual)
}

func Test_C35_159_DynamicStatus_ClonePtr(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	cp := ds.ClonePtr()
	actual := args.Map{"result": cp == nil || cp.Message != "msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned", actual)
}

func Test_C35_160_DynamicStatus_ClonePtr_Nil(t *testing.T) {
	var ds *coredynamic.DynamicStatus
	actual := args.Map{"result": ds.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C35_161_ValueStatus_Invalid(t *testing.T) {
	vs := coredynamic.InvalidValueStatus("msg")
	actual := args.Map{"result": vs.IsValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_C35_162_ValueStatus_InvalidNoMessage(t *testing.T) {
	vs := coredynamic.InvalidValueStatusNoMessage()
	actual := args.Map{"result": vs.IsValid}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// BytesConverter — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C35_163_BytesConverter_SafeCastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	actual := args.Map{"result": bc.SafeCastString() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_C35_164_BytesConverter_SafeCastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	actual := args.Map{"result": bc.SafeCastString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C35_165_BytesConverter_CastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	s, err := bc.CastString()
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello", actual)
}

func Test_C35_166_BytesConverter_CastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	_, err := bc.CastString()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_167_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("true"))
	v, err := bc.ToBool()
	actual := args.Map{"result": err != nil || !v}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C35_168_BytesConverter_Deserialize(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`42`))
	var target int
	err := bc.Deserialize(&target)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	actual := args.Map{"result": target != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C35_169_BytesConverter_ToString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()
	actual := args.Map{"result": err != nil || s != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hello, got err=", actual)
}

func Test_C35_170_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	strs, err := bc.ToStrings()
	actual := args.Map{"result": err != nil || len(strs) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 strings", actual)
}

func Test_C35_171_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`100`))
	v, err := bc.ToInt64()
	actual := args.Map{"result": err != nil || v != 100}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 100", actual)
}

func Test_C35_172_BytesConverter_ToHashmap(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":"1"}`))
	hm, err := bc.ToHashmap()
	actual := args.Map{"result": err != nil || hm == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hashmap", actual)
}

func Test_C35_173_BytesConverter_ToHashmap_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashmap()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_174_BytesConverter_ToHashset(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":true,"b":true}`))
	hs, err := bc.ToHashset()
	actual := args.Map{"result": err != nil || hs == nil || hs.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected hashset with 2 items", actual)
}

func Test_C35_175_BytesConverter_ToHashset_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToHashset()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_176_BytesConverter_ToCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	c, err := bc.ToCollection()
	actual := args.Map{"result": err != nil || c == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected collection", actual)
}

func Test_C35_177_BytesConverter_ToCollection_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`invalid`))
	_, err := bc.ToCollection()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_178_BytesConverter_ToSimpleSlice(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x","y"]`))
	ss, err := bc.ToSimpleSlice()
	actual := args.Map{"result": err != nil || ss == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected simple slice", actual)
}

func Test_C35_179_BytesConverter_ToSimpleSlice_Bad(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`bad`))
	_, err := bc.ToSimpleSlice()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C35_180_NewBytesConverterUsingJsonResult(t *testing.T) {
	jr := corejson.NewPtr(42)
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(jr)
	actual := args.Map{"result": err != nil || bc == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected converter", actual)
}
