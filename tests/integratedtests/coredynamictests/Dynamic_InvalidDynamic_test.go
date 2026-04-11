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

func Test_Dynamic_InvalidDynamic_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"null": d.IsNull(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"null": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic returns error -- with args", actual)
}

func Test_Dynamic_InvalidDynamicPtr_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamicPtr()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"null": d.IsNull(),
		"nn": d != nil,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"null": true,
		"nn": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidDynamicPtr returns error -- with args", actual)
}

func Test_Dynamic_NewDynamicValid_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"data": d.Data(),
		"value": d.Value(),
		"valid": d.IsValid(),
		"invalid": d.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"data": "hello",
		"value": "hello",
		"valid": true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "NewDynamicValid returns non-empty -- with args", actual)
}

func Test_Dynamic_NewDynamic_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, false)

	// Act
	actual := args.Map{
		"data": d.Data(),
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"data": 42,
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "NewDynamic returns correct value -- with args", actual)
}

func Test_Dynamic_Clone(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	cloned := d.Clone()

	// Act
	actual := args.Map{
		"data": cloned.Data(),
		"valid": cloned.IsValid(),
	}

	// Assert
	expected := args.Map{
		"data": "hello",
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Clone", actual)
}

func Test_Dynamic_ClonePtr_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	cp := d.ClonePtr()
	var nilD *coredynamic.Dynamic
	nilCp := nilD.ClonePtr()

	// Act
	actual := args.Map{
		"nn": cp != nil,
		"data": cp.Data(),
		"nilNil": nilCp == nil,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"data": "hello",
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- ClonePtr", actual)
}

func Test_Dynamic_NonPtrPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	np := d.NonPtr()
	p := d.Ptr()

	// Act
	actual := args.Map{
		"npData": np.Data(),
		"pNN": p != nil,
	}

	// Assert
	expected := args.Map{
		"npData": "hello",
		"pNN": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- NonPtr/Ptr", actual)
}

// ═══════════════════════════════════════════
// DynamicGetters — all methods
// ═══════════════════════════════════════════

func Test_Dynamic_Length_Slice_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)

	// Act
	actual := args.Map{"len": d.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Length returns correct value -- slice", actual)
}

func Test_Dynamic_Length_Nil_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)

	// Act
	actual := args.Map{"len": d.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Dynamic_Length_Map(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)

	// Act
	actual := args.Map{"len": d.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Length returns correct value -- map", actual)
}

func Test_Dynamic_StructString(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	s1 := d.StructString()
	s2 := d.String()
	sp := d.StructStringPtr()

	// Act
	actual := args.Map{
		"s1NE": s1 != "",
		"s2NE": s2 != "",
		"spNN": sp != nil,
	}

	// Assert
	expected := args.Map{
		"s1NE": true,
		"s2NE": true,
		"spNN": true,
	}
	expected.ShouldBeEqual(t, 0, "StructString returns correct value -- with args", actual)
}

func Test_Dynamic_StructString_Cached(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	_ = d.StructStringPtr() // first call
	sp := d.StructStringPtr() // cached

	// Act
	actual := args.Map{"spNN": sp != nil}

	// Assert
	expected := args.Map{"spNN": true}
	expected.ShouldBeEqual(t, 0, "StructString returns correct value -- cached", actual)
}

func Test_Dynamic_IsPointer_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	s := "hello"
	dPtr := coredynamic.NewDynamicPtr(&s, true)
	dVal := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{
		"ptr": dPtr.IsPointer(),
		"val": dVal.IsPointer(),
		"valType": dVal.IsValueType(),
	}

	// Assert
	expected := args.Map{
		"ptr": true,
		"val": false,
		"valType": true,
	}
	expected.ShouldBeEqual(t, 0, "IsPointer returns correct value -- with args", actual)
}

func Test_Dynamic_IsPointer_Cached(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	_ = d.IsPointer()

	// Act
	actual := args.Map{"v": d.IsPointer()}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsPointer returns correct value -- cached", actual)
}

func Test_Dynamic_IsStructStringNull(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)

	// Act
	actual := args.Map{"v": d.IsStructStringNullOrEmpty()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmpty returns nil -- nil", actual)
}

func Test_Dynamic_IsStructStringNullOrEmptyOrWhitespace_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)

	// Act
	actual := args.Map{"v": d.IsStructStringNullOrEmptyOrWhitespace()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmptyOrWhitespace returns nil -- nil", actual)
}

func Test_Dynamic_IsPrimitive_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	di := coredynamic.NewDynamicPtr(42, true)

	// Act
	actual := args.Map{
		"str": d.IsPrimitive(),
		"int": di.IsPrimitive(),
	}

	// Assert
	expected := args.Map{
		"str": true,
		"int": true,
	}
	expected.ShouldBeEqual(t, 0, "IsPrimitive returns correct value -- with args", actual)
}

func Test_Dynamic_IsNumber_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	di := coredynamic.NewDynamicPtr(42, true)
	ds := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{
		"int": di.IsNumber(),
		"str": ds.IsNumber(),
	}

	// Assert
	expected := args.Map{
		"int": true,
		"str": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNumber returns correct value -- with args", actual)
}

func Test_Dynamic_IsStringType_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{"v": d.IsStringType()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStringType returns correct value -- with args", actual)
}

func Test_Dynamic_IsStruct_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	type ts struct{ X int }
	d := coredynamic.NewDynamicPtr(ts{X: 1}, true)

	// Act
	actual := args.Map{"v": d.IsStruct()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStruct returns correct value -- with args", actual)
}

func Test_Dynamic_IsFunc_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(func() {}, true)

	// Act
	actual := args.Map{"v": d.IsFunc()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsFunc returns correct value -- with args", actual)
}

func Test_Dynamic_IsSliceOrArray_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]int{1}, true)

	// Act
	actual := args.Map{"v": d.IsSliceOrArray()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsSliceOrArray returns correct value -- with args", actual)
}

func Test_Dynamic_IsSliceOrArrayOrMap_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	dm := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)

	// Act
	actual := args.Map{"v": dm.IsSliceOrArrayOrMap()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsSliceOrArrayOrMap returns correct value -- with args", actual)
}

func Test_Dynamic_IsMap_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	dm := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)

	// Act
	actual := args.Map{"v": dm.IsMap()}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMap returns correct value -- with args", actual)
}

func Test_Dynamic_IntDefault_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)
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
	expected.ShouldBeEqual(t, 0, "IntDefault returns nil -- nil", actual)
}

func Test_Dynamic_IntDefault_ParseOK(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(42, true)
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
	expected.ShouldBeEqual(t, 0, "IntDefault returns correct value -- parse ok", actual)
}

func Test_Dynamic_IntDefault_ParseFail(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("abc", true)
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
	expected.ShouldBeEqual(t, 0, "IntDefault returns correct value -- parse fail", actual)
}

func Test_Dynamic_Float64_Nil_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)
	_, err := d.Float64()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns nil -- nil", actual)
}

func Test_Dynamic_Float64_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(3.14, true)
	val, err := d.Float64()

	// Act
	actual := args.Map{
		"gt3": val > 3.0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"gt3": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Float64 returns non-empty -- valid", actual)
}

func Test_Dynamic_Float64_ParseFail(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("abc", true)
	_, err := d.Float64()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Float64 returns correct value -- parse fail", actual)
}

func Test_Dynamic_ValueInt_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)
	ds := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"int": d.ValueInt(),
		"str": ds.ValueInt(),
	}

	// Assert
	expected := args.Map{
		"int": 42,
		"str": -1,
	}
	expected.ShouldBeEqual(t, 0, "ValueInt returns correct value -- with args", actual)
}

func Test_Dynamic_ValueUInt_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(uint(42))
	ds := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"uint": d.ValueUInt(),
		"str": ds.ValueUInt(),
	}

	// Assert
	expected := args.Map{
		"uint": uint(42),
		"str": uint(0),
	}
	expected.ShouldBeEqual(t, 0, "ValueUInt returns correct value -- with args", actual)
}

func Test_Dynamic_ValueStrings_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	ds := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"len": len(d.ValueStrings()),
		"nilStr": ds.ValueStrings() == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilStr": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueStrings returns non-empty -- with args", actual)
}

func Test_Dynamic_ValueBool_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(true)
	ds := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"bool": d.ValueBool(),
		"str": ds.ValueBool(),
	}

	// Assert
	expected := args.Map{
		"bool": true,
		"str": false,
	}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
}

func Test_Dynamic_ValueInt64_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(int64(99))
	ds := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"int64": d.ValueInt64(),
		"str": ds.ValueInt64(),
	}

	// Assert
	expected := args.Map{
		"int64": int64(99),
		"str": int64(-1),
	}
	expected.ShouldBeEqual(t, 0, "ValueInt64 returns correct value -- with args", actual)
}

func Test_Dynamic_ValueNullErr_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	var nilD *coredynamic.Dynamic
	errNil := nilD.ValueNullErr()
	d := coredynamic.NewDynamicPtr(nil, false)
	errNull := d.ValueNullErr()
	dv := coredynamic.NewDynamicPtr("hello", true)
	errOK := dv.ValueNullErr()

	// Act
	actual := args.Map{
		"nil": errNil != nil,
		"null": errNull != nil,
		"ok": errOK == nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"null": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueNullErr returns error -- with args", actual)
}

func Test_Dynamic_ValueString_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	var nilD *coredynamic.Dynamic
	vs1 := nilD.ValueString()
	d := coredynamic.NewDynamicPtr("hello", true)
	vs2 := d.ValueString()
	di := coredynamic.NewDynamicPtr(42, true)
	vs3 := di.ValueString()

	// Act
	actual := args.Map{
		"nil": vs1,
		"str": vs2,
		"intNE": vs3 != "",
	}

	// Assert
	expected := args.Map{
		"nil": "",
		"str": "hello",
		"intNE": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- with args", actual)
}

func Test_Dynamic_Bytes_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	var nilD *coredynamic.Dynamic
	b1, ok1 := nilD.Bytes()
	d := coredynamic.NewDynamicPtr([]byte("hello"), true)
	b2, ok2 := d.Bytes()
	ds := coredynamic.NewDynamicPtr("hello", true)
	_, ok3 := ds.Bytes()

	// Act
	actual := args.Map{
		"nilB": b1 == nil,
		"nilOK": ok1,
		"bLen": len(b2) > 0,
		"bOK": ok2,
		"sOK": ok3,
	}

	// Assert
	expected := args.Map{
		"nilB": true,
		"nilOK": false,
		"bLen": true,
		"bOK": true,
		"sOK": false,
	}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// DynamicJson — all methods
// ═══════════════════════════════════════════

func Test_Dynamic_JsonBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	jb, err := d.JsonBytes()

	// Act
	actual := args.Map{
		"len": len(jb) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonBytes returns correct value -- with args", actual)
}

func Test_Dynamic_JsonBytesPtr_Null_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)
	jb, err := d.JsonBytesPtr()

	// Act
	actual := args.Map{
		"empty": len(jb) == 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonBytesPtr returns correct value -- null", actual)
}

func Test_Dynamic_JsonString_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	js, err := d.JsonString()

	// Act
	actual := args.Map{
		"ne": js != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"ne": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct value -- with args", actual)
}

func Test_Dynamic_JsonStringMust_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	js := d.JsonStringMust()

	// Act
	actual := args.Map{"ne": js != ""}

	// Assert
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct value -- with args", actual)
}

func Test_Dynamic_MarshalJSON_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.MarshalJSON()

	// Act
	actual := args.Map{
		"len": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns correct value -- with args", actual)
}

func Test_Dynamic_ValueMarshal_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	b, err := d.ValueMarshal()

	// Act
	actual := args.Map{
		"len": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns correct value -- with args", actual)
}

func Test_Dynamic_ValueMarshal_Nil_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns nil -- nil", actual)
}

func Test_Dynamic_JsonPayloadMust_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	b := d.JsonPayloadMust()

	// Act
	actual := args.Map{"len": len(b) > 0}

	// Assert
	expected := args.Map{"len": true}
	expected.ShouldBeEqual(t, 0, "JsonPayloadMust returns correct value -- with args", actual)
}

func Test_Dynamic_JsonModelAny_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"v": d.JsonModel(),
		"any": d.JsonModelAny(),
	}

	// Assert
	expected := args.Map{
		"v": "hello",
		"any": "hello",
	}
	expected.ShouldBeEqual(t, 0, "JsonModel returns correct value -- with args", actual)
}

func Test_Dynamic_Json_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	j := d.Json()
	jp := d.JsonPtr()

	// Act
	actual := args.Map{
		"jLen": j.Length() > 0,
		"jpNN": jp != nil,
	}

	// Assert
	expected := args.Map{
		"jLen": true,
		"jpNN": true,
	}
	expected.ShouldBeEqual(t, 0, "Json returns correct value -- with args", actual)
}

func Test_Dynamic_Deserialize_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, err := d.Deserialize([]byte(`"hello"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Deserialize returns nil -- nil", actual)
}

func Test_Dynamic_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`"hello"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns nil -- nil", actual)
}

// ═══════════════════════════════════════════
// DynamicReflect — all methods
// ═══════════════════════════════════════════

func Test_Dynamic_ReflectValue_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	rv := d.ReflectValue()
	rv2 := d.ReflectValue() // cached

	// Act
	actual := args.Map{
		"nn": rv != nil,
		"same": rv == rv2,
	}

	// Assert
	expected := args.Map{
		"nn": true,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- with args", actual)
}

func Test_Dynamic_ReflectKind_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{"kind": d.ReflectKind() == reflect.String}

	// Assert
	expected := args.Map{"kind": true}
	expected.ShouldBeEqual(t, 0, "ReflectKind returns correct value -- with args", actual)
}

func Test_Dynamic_ReflectTypeName(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{"ne": d.ReflectTypeName() != ""}

	// Assert
	expected := args.Map{"ne": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeName returns correct value -- with args", actual)
}

func Test_Dynamic_ReflectType_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	rt := d.ReflectType()
	rt2 := d.ReflectType() // cached

	// Act
	actual := args.Map{
		"name": rt.Name(),
		"same": rt == rt2,
	}

	// Assert
	expected := args.Map{
		"name": "string",
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectType returns correct value -- with args", actual)
}

func Test_Dynamic_IsReflectTypeOf_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{"v": d.IsReflectTypeOf(reflect.TypeOf(""))}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsReflectTypeOf returns correct value -- with args", actual)
}

func Test_Dynamic_IsReflectKind_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)

	// Act
	actual := args.Map{"v": d.IsReflectKind(reflect.String)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsReflectKind returns correct value -- with args", actual)
}

func Test_Dynamic_ItemUsingIndex_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]string{"a", "b", "c"}, true)
	rv := d.ItemReflectValueUsingIndex(1)
	item := d.ItemUsingIndex(1)

	// Act
	actual := args.Map{
		"rv": rv.String(),
		"item": item,
	}

	// Assert
	expected := args.Map{
		"rv": "b",
		"item": "b",
	}
	expected.ShouldBeEqual(t, 0, "ItemUsingIndex returns correct value -- with args", actual)
}

func Test_Dynamic_ItemUsingKey_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(map[string]int{"x": 42}, true)
	rv := d.ItemReflectValueUsingKey("x")
	item := d.ItemUsingKey("x")

	// Act
	actual := args.Map{
		"rv": int(rv.Int()),
		"item": item,
	}

	// Assert
	expected := args.Map{
		"rv": 42,
		"item": 42,
	}
	expected.ShouldBeEqual(t, 0, "ItemUsingKey returns correct value -- with args", actual)
}

func Test_Dynamic_ReflectSetTo_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	var target string
	err := d.ReflectSetTo(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"target": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"target": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- with args", actual)
}

func Test_Dynamic_ReflectSetTo_Nil_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns nil -- nil", actual)
}

func Test_Dynamic_MapToKeyVal_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2}, true)
	kvc, err := d.MapToKeyVal()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": kvc.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapToKeyVal returns correct value -- with args", actual)
}

func Test_Dynamic_Loop_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]int{10, 20, 30}, true)
	sum := 0
	called := d.Loop(func(index int, item any) bool {
		sum += item.(int)
		return false
	})

	// Act
	actual := args.Map{
		"sum": sum,
		"called": called,
	}

	// Assert
	expected := args.Map{
		"sum": 60,
		"called": true,
	}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- with args", actual)
}

func Test_Dynamic_Loop_Empty(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.Loop(func(index int, item any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Loop returns empty -- empty", actual)
}

func Test_Dynamic_Loop_Break_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3}, true)
	count := 0
	called := d.Loop(func(index int, item any) bool {
		count++
		return index == 1
	})

	// Act
	actual := args.Map{
		"count": count,
		"called": called,
	}

	// Assert
	expected := args.Map{
		"count": 2,
		"called": true,
	}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- break", actual)
}

func Test_Dynamic_FilterAsDynamicCollection_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt()%2 == 0, false
	})

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns correct value -- with args", actual)
}

func Test_Dynamic_FilterAsDynamicCollection_Empty(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Filter returns empty -- empty", actual)
}

func Test_Dynamic_FilterAsDynamicCollection_Break_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr([]int{1, 2, 3, 4}, true)
	result := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, index == 1
	})

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Filter returns correct value -- break", actual)
}

func Test_Dynamic_LoopMap_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1}, true)
	count := 0
	called := d.LoopMap(func(index int, key, value any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{
		"count": count,
		"called": called,
	}

	// Assert
	expected := args.Map{
		"count": 1,
		"called": true,
	}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- with args", actual)
}

func Test_Dynamic_LoopMap_Empty(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(nil, false)
	called := d.LoopMap(func(index int, key, value any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LoopMap returns empty -- empty", actual)
}

func Test_Dynamic_LoopMap_Break(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr(map[string]int{"a": 1, "b": 2, "c": 3}, true)
	count := 0
	d.LoopMap(func(index int, key, value any) bool {
		count++
		return true
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- break", actual)
}

func Test_Dynamic_ConvertUsingFunc_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	converter := func(in any, typeMust reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(in)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(""))

	// Act
	actual := args.Map{
		"valid": result.IsValid(),
		"result": result.Result,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ConvertUsingFunc returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// DynamicStatus — all methods
// ═══════════════════════════════════════════

func Test_DynamicStatus_Invalid_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	ds2 := coredynamic.InvalidDynamicStatus("err")

	// Act
	actual := args.Map{
		"valid": ds.IsValid(), "msg": ds.Message,
		"valid2": ds2.IsValid(), "msg2": ds2.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
		"valid2": false,
		"msg2": "err",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns error -- invalid", actual)
}

func Test_DynamicStatus_Clone_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("err")
	cloned := ds.Clone()
	clonedPtr := ds.ClonePtr()
	var nilDS *coredynamic.DynamicStatus
	nilClone := nilDS.ClonePtr()

	// Act
	actual := args.Map{
		"msg": cloned.Message, "ptrNN": clonedPtr != nil, "nilNil": nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"msg": "err",
		"ptrNN": true,
		"nilNil": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- clone", actual)
}

// ═══════════════════════════════════════════
// Standalone functions
// ═══════════════════════════════════════════

func Test_LengthOfReflect(t *testing.T) {
	// Arrange
	s := reflect.ValueOf([]int{1, 2, 3})
	a := reflect.ValueOf([2]int{1, 2})
	m := reflect.ValueOf(map[string]int{"a": 1})
	str := reflect.ValueOf("hello")

	// Act
	actual := args.Map{
		"slice": coredynamic.LengthOfReflect(s),
		"arr": coredynamic.LengthOfReflect(a),
		"map": coredynamic.LengthOfReflect(m),
		"str": coredynamic.LengthOfReflect(str),
	}

	// Assert
	expected := args.Map{
		"slice": 3,
		"arr": 2,
		"map": 1,
		"str": 0,
	}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- with args", actual)
}

func Test_LengthOfReflect_Ptr(t *testing.T) {
	// Arrange
	s := []int{1, 2}
	rv := reflect.ValueOf(&s)

	// Act
	actual := args.Map{"v": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"v": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- pointer", actual)
}

func Test_ReflectInterfaceVal_Value(t *testing.T) {
	// Arrange
	v := coredynamic.ReflectInterfaceVal("hello")

	// Act
	actual := args.Map{"v": v}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- value", actual)
}

func Test_ReflectInterfaceVal_Ptr(t *testing.T) {
	// Arrange
	s := "hello"
	v := coredynamic.ReflectInterfaceVal(&s)

	// Act
	actual := args.Map{"v": v}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- ptr", actual)
}

func Test_SafeTypeName(t *testing.T) {
	// Act
	actual := args.Map{
		"str": coredynamic.SafeTypeName("hello"),
		"nil": coredynamic.SafeTypeName(nil),
	}

	// Assert
	expected := args.Map{
		"str": "string",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- with args", actual)
}

func Test_ZeroSetAny_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	type ts struct{ Name string }
	v := &ts{Name: "hello"}
	coredynamic.ZeroSetAny(v)

	// Act
	actual := args.Map{"name": v.Name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns correct value -- with args", actual)
}

func Test_ZeroSetAny_Nil_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	coredynamic.ZeroSetAny(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns nil -- nil", actual)
}

func Test_AnyToReflectVal_FromDynamicInvalidDynami(t *testing.T) {
	// Arrange
	rv := coredynamic.AnyToReflectVal("hello")

	// Act
	actual := args.Map{"kind": rv.Kind() == reflect.String}

	// Assert
	expected := args.Map{"kind": true}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- with args", actual)
}

func Test_CastTo_Matching(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{
		"valid": result.IsValid,
		"match": result.IsMatchingAcceptedType,
		"null": result.IsNull,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"match": true,
		"null": false,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- matching", actual)
}

func Test_CastTo_NotMatching(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(42))

	// Act
	actual := args.Map{
		"match": result.IsMatchingAcceptedType,
		"hasErr": result.HasError(),
	}

	// Assert
	expected := args.Map{
		"match": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- not matching", actual)
}
