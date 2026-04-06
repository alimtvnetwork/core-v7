package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Dynamic — constructors and core
// ═══════════════════════════════════════════

func Test_Cov11_Dynamic_InvalidDynamic(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	actual := args.Map{"valid": d.IsValid(), "null": d.IsNull()}
	expected := args.Map{"valid": false, "null": true}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic returns error -- with args", actual)
}

func Test_Cov11_Dynamic_InvalidDynamicPtr(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	actual := args.Map{"valid": d.IsValid(), "null": d.IsNull(), "nn": d != nil}
	expected := args.Map{"valid": false, "null": true, "nn": true}
	expected.ShouldBeEqual(t, 0, "InvalidDynamicPtr returns error -- with args", actual)
}

func Test_Cov11_Dynamic_NewDynamicValid(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"data": d.Data(), "value": d.Value(), "valid": d.IsValid(), "invalid": d.IsInvalid()}
	expected := args.Map{"data": "hello", "value": "hello", "valid": true, "invalid": false}
	expected.ShouldBeEqual(t, 0, "NewDynamicValid returns non-empty -- with args", actual)
}

func Test_Cov11_Dynamic_NewDynamic(t *testing.T) {
	d := coredynamic.NewDynamic(42, false)
	actual := args.Map{"data": d.Data(), "valid": d.IsValid()}
	expected := args.Map{"data": 42, "valid": false}
	expected.ShouldBeEqual(t, 0, "NewDynamic returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	cloned := d.Clone()
	actual := args.Map{"data": cloned.Data(), "valid": cloned.IsValid()}
	expected := args.Map{"data": "hello", "valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Clone", actual)
}
func Test_Cov11_Dynamic_NonPtrPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	np := d.NonPtr()
	p := d.Ptr()
	actual := args.Map{"npData": np.Data(), "pNN": p != nil}
	expected := args.Map{"npData": "hello", "pNN": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- NonPtr/Ptr", actual)
}

// ═══════════════════════════════════════════
// DynamicGetters — all methods
// ═══════════════════════════════════════════

func Test_Cov11_Dynamic_Length_Slice(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	actual := args.Map{"len": d.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Length returns correct value -- slice", actual)
}

func Test_Cov11_Dynamic_Length_Nil(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	actual := args.Map{"len": d.Length()}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_Length_Map(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	actual := args.Map{"len": d.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Length returns correct value -- map", actual)
}

func Test_Cov11_Dynamic_StructString(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	s1 := d.StructString()
	s2 := d.String()
	sp := d.StructStringPtr()
	actual := args.Map{"s1NE": s1 != "", "s2NE": s2 != "", "spNN": sp != nil}
	expected := args.Map{"s1NE": true, "s2NE": true, "spNN": true}
	expected.ShouldBeEqual(t, 0, "StructString returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_StructString_Cached(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	_ = d.StructStringPtr() // first call
	sp := d.StructStringPtr() // cached
	actual := args.Map{"spNN": sp != nil}
	expected := args.Map{"spNN": true}
	expected.ShouldBeEqual(t, 0, "StructString returns correct value -- cached", actual)
}

func Test_Cov11_Dynamic_IsPointer(t *testing.T) {
	s := "hello"
	dPtr := coredynamic.NewDynamicPtr(&s, true)
	dVal := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"ptr": dPtr.IsPointer(), "val": dVal.IsPointer(), "valType": dVal.IsValueType()}
	expected := args.Map{"ptr": true, "val": false, "valType": true}
	expected.ShouldBeEqual(t, 0, "IsPointer returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsPointer_Cached(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	_ = d.IsPointer()
	actual := args.Map{"v": d.IsPointer()}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsPointer returns correct value -- cached", actual)
}

func Test_Cov11_Dynamic_IsStructStringNull(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	actual := args.Map{"v": d.IsStructStringNullOrEmpty()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmpty returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	actual := args.Map{"v": d.IsStructStringNullOrEmptyOrWhitespace()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmptyOrWhitespace returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_IsPrimitive(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	di := coredynamic.NewDynamicPtr(42, true)
	actual := args.Map{"str": d.IsPrimitive(), "int": di.IsPrimitive()}
	expected := args.Map{"str": true, "int": true}
	expected.ShouldBeEqual(t, 0, "IsPrimitive returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsNumber(t *testing.T) {
	di := coredynamic.NewDynamicPtr(42, true)
	ds := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"int": di.IsNumber(), "str": ds.IsNumber()}
	expected := args.Map{"int": true, "str": false}
	expected.ShouldBeEqual(t, 0, "IsNumber returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsStringType(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"v": d.IsStringType()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringType returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsStruct(t *testing.T) {
	type ts struct{ X int }
	d := coredynamic.NewDynamicPtr(ts{X: 1}, true)
	actual := args.Map{"v": d.IsStruct()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStruct returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsFunc(t *testing.T) {
	d := coredynamic.NewDynamicPtr(func() {}, true)
	actual := args.Map{"v": d.IsFunc()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsFunc returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsSliceOrArray(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1}, true)
	actual := args.Map{"v": d.IsSliceOrArray()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsSliceOrArray returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsSliceOrArrayOrMap(t *testing.T) {
	dm := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"v": dm.IsSliceOrArrayOrMap()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsSliceOrArrayOrMap returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsMap(t *testing.T) {
	dm := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	actual := args.Map{"v": dm.IsMap()}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMap returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IntDefault_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "IntDefault returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_IntDefault_ParseOK(t *testing.T) {
	d := coredynamic.NewDynamicPtr(42, true)
	val, ok := d.IntDefault(0)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "IntDefault returns correct value -- parse ok", actual)
}

func Test_Cov11_Dynamic_IntDefault_ParseFail(t *testing.T) {
	d := coredynamic.NewDynamicPtr("abc", true)
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "IntDefault returns correct value -- parse fail", actual)
}

func Test_Cov11_Dynamic_Float64_Nil(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_Float64_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr(3.14, true)
	val, err := d.Float64()
	actual := args.Map{"gt3": val > 3.0, "noErr": err == nil}
	expected := args.Map{"gt3": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns non-empty -- valid", actual)
}

func Test_Cov11_Dynamic_Float64_ParseFail(t *testing.T) {
	d := coredynamic.NewDynamicPtr("abc", true)
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns correct value -- parse fail", actual)
}

func Test_Cov11_Dynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	ds := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"int": d.ValueInt(), "str": ds.ValueInt()}
	expected := args.Map{"int": 42, "str": -1}
	expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(uint(42))
	ds := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"uint": d.ValueUInt(), "str": ds.ValueUInt()}
	expected := args.Map{"uint": uint(42), "str": uint(0)}
	expected.ShouldBeEqual(t, 0, "ValueUInt returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	ds := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"len": len(d.ValueStrings()), "nilStr": ds.ValueStrings() == nil}
	expected := args.Map{"len": 2, "nilStr": true}
	expected.ShouldBeEqual(t, 0, "ValueStrings returns non-empty -- with args", actual)
}

func Test_Cov11_Dynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamicValid(true)
	ds := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"bool": d.ValueBool(), "str": ds.ValueBool()}
	expected := args.Map{"bool": true, "str": false}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamicValid(int64(99))
	ds := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"int64": d.ValueInt64(), "str": ds.ValueInt64()}
	expected := args.Map{"int64": int64(99), "str": int64(-1)}
	expected.ShouldBeEqual(t, 0, "ValueInt64 returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ValueNullErr(t *testing.T) {
	var nilD *coredynamic.Dynamic
	errNil := nilD.ValueNullErr()
	d := coredynamic.NewDynamicPtr(nil, false)
	errNull := d.ValueNullErr()
	dv := coredynamic.NewDynamicPtr("hello", true)
	errOK := dv.ValueNullErr()
	actual := args.Map{"nil": errNil != nil, "null": errNull != nil, "ok": errOK == nil}
	expected := args.Map{"nil": true, "null": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "ValueNullErr returns error -- with args", actual)
}

func Test_Cov11_Dynamic_ValueString(t *testing.T) {
	var nilD *coredynamic.Dynamic
	vs1 := nilD.ValueString()
	d := coredynamic.NewDynamicPtr("hello", true)
	vs2 := d.ValueString()
	di := coredynamic.NewDynamicPtr(42, true)
	vs3 := di.ValueString()
	actual := args.Map{"nil": vs1, "str": vs2, "intNE": vs3 != ""}
	expected := args.Map{"nil": "", "str": "hello", "intNE": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- with args", actual)
}

func Test_Cov11_Dynamic_Bytes(t *testing.T) {
	var nilD *coredynamic.Dynamic
	b1, ok1 := nilD.Bytes()
	d := coredynamic.NewDynamicPtr([]byte("hello"), true)
	b2, ok2 := d.Bytes()
	ds := coredynamic.NewDynamicPtr("hello", true)
	_, ok3 := ds.Bytes()
	actual := args.Map{"nilB": b1 == nil, "nilOK": ok1, "bLen": len(b2) > 0, "bOK": ok2, "sOK": ok3}
	expected := args.Map{"nilB": true, "nilOK": false, "bLen": true, "bOK": true, "sOK": false}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// DynamicJson — all methods
// ═══════════════════════════════════════════

func Test_Cov11_Dynamic_JsonBytes(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	jb, err := d.JsonBytes()
	actual := args.Map{"len": len(jb) > 0, "noErr": err == nil}
	expected := args.Map{"len": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonBytes returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	jb, err := d.JsonBytesPtr()
	actual := args.Map{"empty": len(jb) == 0, "noErr": err == nil}
	expected := args.Map{"empty": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonBytesPtr returns correct value -- null", actual)
}

func Test_Cov11_Dynamic_JsonString(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	js, err := d.JsonString()
	actual := args.Map{"ne": js != "", "noErr": err == nil}
	expected := args.Map{"ne": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	js := d.JsonStringMust()
	actual := args.Map{"ne": js != ""}
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_MarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.MarshalJSON()
	actual := args.Map{"len": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"len": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ValueMarshal(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.ValueMarshal()
	actual := args.Map{"len": len(b) > 0, "noErr": err == nil}
	expected := args.Map{"len": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ValueMarshal_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	b := d.JsonPayloadMust()
	actual := args.Map{"len": len(b) > 0}
	expected := args.Map{"len": true}
	expected.ShouldBeEqual(t, 0, "JsonPayloadMust returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_JsonModelAny(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"v": d.JsonModel(), "any": d.JsonModelAny()}
	expected := args.Map{"v": "hello", "any": "hello"}
	expected.ShouldBeEqual(t, 0, "JsonModel returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_Json(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	j := d.Json()
	jp := d.JsonPtr()
	actual := args.Map{"jLen": j.Length() > 0, "jpNN": jp != nil}
	expected := args.Map{"jLen": true, "jpNN": true}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`"hello"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_UnmarshalJSON_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`"hello"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns nil -- nil", actual)
}

// ═══════════════════════════════════════════
// DynamicReflect — all methods
// ═══════════════════════════════════════════

func Test_Cov11_Dynamic_ReflectValue(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	rv := d.ReflectValue()
	rv2 := d.ReflectValue() // cached
	actual := args.Map{"nn": rv != nil, "same": rv == rv2}
	expected := args.Map{"nn": true, "same": true}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ReflectKind(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"kind": d.ReflectKind() == reflect.String}
	expected := args.Map{"kind": true}
	expected.ShouldBeEqual(t, 0, "ReflectKind returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ReflectTypeName(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"ne": d.ReflectTypeName() != ""}
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeName returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ReflectType(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	rt := d.ReflectType()
	rt2 := d.ReflectType() // cached
	actual := args.Map{"name": rt.Name(), "same": rt == rt2}
	expected := args.Map{"name": "string", "same": true}
	expected.ShouldBeEqual(t, 0, "ReflectType returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsReflectTypeOf(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"v": d.IsReflectTypeOf(reflect.TypeOf(""))}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsReflectTypeOf returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_IsReflectKind(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	actual := args.Map{"v": d.IsReflectKind(reflect.String)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsReflectKind returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]string{"a", "b", "c"}, true)
	rv := d.ItemReflectValueUsingIndex(1)
	item := d.ItemUsingIndex(1)
	actual := args.Map{"rv": rv.String(), "item": item}
	expected := args.Map{"rv": "b", "item": "b"}
	expected.ShouldBeEqual(t, 0, "ItemUsingIndex returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"x": 42}, true)
	rv := d.ItemReflectValueUsingKey("x")
	item := d.ItemUsingKey("x")
	actual := args.Map{"rv": int(rv.Int()), "item": item}
	expected := args.Map{"rv": 42, "item": 42}
	expected.ShouldBeEqual(t, 0, "ItemUsingKey returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ReflectSetTo(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	var target string
	err := d.ReflectSetTo(&target)
	actual := args.Map{"noErr": err == nil, "target": target}
	expected := args.Map{"noErr": true, "target": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_ReflectSetTo_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns nil -- nil", actual)
}

func Test_Cov11_Dynamic_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	kvc, err := d.MapToKeyVal()
	actual := args.Map{"noErr": err == nil, "len": kvc.Length()}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "MapToKeyVal returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_Loop(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{10, 20, 30}, true)
	sum := 0
	called := d.Loop(func(index int, item any) bool {
		sum += item.(int)
		return false
	})
	actual := args.Map{"sum": sum, "called": called}
	expected := args.Map{"sum": 60, "called": true}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_Loop_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.Loop(func(index int, item any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Loop returns empty -- empty", actual)
}

func Test_Cov11_Dynamic_Loop_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	called := d.Loop(func(index int, item any) bool {
		count++
		return index == 1
	})
	actual := args.Map{"count": count, "called": called}
	expected := args.Map{"count": 2, "called": true}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- break", actual)
}

func Test_Cov11_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt()%2 == 0, false
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_FilterAsDynamicCollection_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Filter returns empty -- empty", actual)
}

func Test_Cov11_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, index == 1
	})
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Filter returns correct value -- break", actual)
}

func Test_Cov11_Dynamic_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	count := 0
	called := d.LoopMap(func(index int, key, value any) bool {
		count++
		return false
	})
	actual := args.Map{"count": count, "called": called}
	expected := args.Map{"count": 1, "called": true}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- with args", actual)
}

func Test_Cov11_Dynamic_LoopMap_Empty(t *testing.T) {
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.LoopMap(func(index int, key, value any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LoopMap returns empty -- empty", actual)
}

func Test_Cov11_Dynamic_LoopMap_Break(t *testing.T) {
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2, "c": 3}, true)
	count := 0
	d.LoopMap(func(index int, key, value any) bool {
		count++
		return true
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- break", actual)
}

func Test_Cov11_Dynamic_ConvertUsingFunc(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	converter := func(in any, typeMust reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(in)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(""))
	actual := args.Map{"valid": result.IsValid(), "result": result.Result}
	expected := args.Map{"valid": true, "result": "hello"}
	expected.ShouldBeEqual(t, 0, "ConvertUsingFunc returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// DynamicStatus — all methods
// ═══════════════════════════════════════════

func Test_Cov11_DynamicStatus_Invalid(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	ds2 := coredynamic.InvalidDynamicStatus("err")
	actual := args.Map{
		"valid": ds.IsValid(), "msg": ds.Message,
		"valid2": ds2.IsValid(), "msg2": ds2.Message,
	}
	expected := args.Map{"valid": false, "msg": "", "valid2": false, "msg2": "err"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns error -- invalid", actual)
}
func Test_Cov11_LengthOfReflect(t *testing.T) {
	s := reflect.ValueOf([]int{1, 2, 3})
	a := reflect.ValueOf([2]int{1, 2})
	m := reflect.ValueOf(map[string]int{"a": 1})
	str := reflect.ValueOf("hello")
	actual := args.Map{"slice": coredynamic.LengthOfReflect(s), "arr": coredynamic.LengthOfReflect(a), "map": coredynamic.LengthOfReflect(m), "str": coredynamic.LengthOfReflect(str)}
	expected := args.Map{"slice": 3, "arr": 2, "map": 1, "str": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- with args", actual)
}

func Test_Cov11_LengthOfReflect_Ptr(t *testing.T) {
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)
	actual := args.Map{"v": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"v": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- pointer", actual)
}

func Test_Cov11_ReflectInterfaceVal_Value(t *testing.T) {
	v := coredynamic.ReflectInterfaceVal("hello")
	actual := args.Map{"v": v}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- value", actual)
}

func Test_Cov11_ReflectInterfaceVal_Ptr(t *testing.T) {
	s := "hello"
	v := coredynamic.ReflectInterfaceVal(&s)
	actual := args.Map{"v": v}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- ptr", actual)
}

func Test_Cov11_SafeTypeName(t *testing.T) {
	actual := args.Map{"str": coredynamic.SafeTypeName("hello"), "nil": coredynamic.SafeTypeName(nil)}
	expected := args.Map{"str": "string", "nil": ""}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- with args", actual)
}

func Test_Cov11_ZeroSetAny(t *testing.T) {
	type ts struct{ Name string }
	v := &ts{Name: "hello"}
	coredynamic.ZeroSetAny(v)
	actual := args.Map{"name": v.Name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns correct value -- with args", actual)
}

func Test_Cov11_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns nil -- nil", actual)
}

func Test_Cov11_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal("hello")
	actual := args.Map{"kind": rv.Kind() == reflect.String}
	expected := args.Map{"kind": true}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- with args", actual)
}

func Test_Cov11_CastTo_Matching(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	actual := args.Map{"valid": result.IsValid, "match": result.IsMatchingAcceptedType, "null": result.IsNull}
	expected := args.Map{"valid": true, "match": true, "null": false}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- matching", actual)
}

func Test_Cov11_CastTo_NotMatching(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(42))
	actual := args.Map{"match": result.IsMatchingAcceptedType, "hasErr": result.HasError()}
	expected := args.Map{"match": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- not matching", actual)
}
