package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// TypedDynamic — constructors / accessors
// =============================================================================

func Test_TD_NewTypedDynamic(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"data": d.Data(),
		"val": d.Value(),
		"invalid": d.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"data": "hello",
		"val": "hello",
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedDynamic", actual)
}

func Test_TD_NewTypedDynamicValid(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid(42)

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"data": d.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"data": 42,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedDynamicValid", actual)
}

func Test_TD_NewTypedDynamicPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("x", true)

	// Act
	actual := args.Map{
		"notNil": d != nil,
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedDynamicPtr", actual)
}

func Test_TD_InvalidTypedDynamic(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidTypedDynamic[string]()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"invalid": d.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidTypedDynamic", actual)
}

func Test_TD_InvalidTypedDynamicPtr(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidTypedDynamicPtr[int]()

	// Act
	actual := args.Map{
		"notNil": d != nil,
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidTypedDynamicPtr", actual)
}

func Test_TD_String(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(42, true)

	// Act
	actual := args.Map{"r": d.String()}

	// Assert
	expected := args.Map{"r": "42"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic String", actual)
}

// =============================================================================
// TypedDynamic — JSON
// =============================================================================

func Test_TD_JsonBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	b, err := d.JsonBytes()

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
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonBytes", actual)
}

func Test_TD_JsonResult(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	r := d.JsonResult()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonResult", actual)
}

func Test_TD_JsonString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	s, err := d.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": `"hello"`,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonString", actual)
}

func Test_TD_MarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	b, err := json.Marshal(d)

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
	expected.ShouldBeEqual(t, 0, "TypedDynamic MarshalJSON", actual)
}

func Test_TD_UnmarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := json.Unmarshal([]byte(`"world"`), d)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"valid": d.IsValid(),
		"data": d.Data(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"valid": true,
		"data": "world",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic UnmarshalJSON", actual)
}

func Test_TD_ValueMarshal(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
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
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueMarshal", actual)
}

func Test_TD_Bytes_IsBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic([]byte("hello"), true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Bytes isBytes", actual)
}

func Test_TD_Bytes_NotBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Bytes notBytes", actual)
}

// =============================================================================
// TypedDynamic — GetAs* / Value*
// =============================================================================

func Test_TD_GetAsString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	v, ok := d.GetAsString()

	// Act
	actual := args.Map{
		"ok": ok,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"v": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsString", actual)
}

func Test_TD_GetAsInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(42, true)
	v, ok := d.GetAsInt()

	// Act
	actual := args.Map{
		"ok": ok,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"v": 42,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsInt", actual)
}

func Test_TD_GetAsInt64(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(int64(99), true)
	v, ok := d.GetAsInt64()

	// Act
	actual := args.Map{
		"ok": ok,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"v": int64(99),
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsInt64", actual)
}

func Test_TD_GetAsUint(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(uint(5), true)
	v, ok := d.GetAsUint()

	// Act
	actual := args.Map{
		"ok": ok,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"v": uint(5),
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsUint", actual)
}

func Test_TD_GetAsFloat64(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(3.14, true)
	v, ok := d.GetAsFloat64()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsFloat64", actual)
	_ = v
}

func Test_TD_GetAsFloat32(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(float32(1.5), true)
	_, ok := d.GetAsFloat32()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsFloat32", actual)
}

func Test_TD_GetAsBool(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(true, true)
	v, ok := d.GetAsBool()

	// Act
	actual := args.Map{
		"ok": ok,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"v": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsBool", actual)
}

func Test_TD_GetAsBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic([]byte("hi"), true)
	_, ok := d.GetAsBytes()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsBytes", actual)
}

func Test_TD_GetAsStrings(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic([]string{"a"}, true)
	_, ok := d.GetAsStrings()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsStrings", actual)
}

func Test_TD_ValueString_IsString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)

	// Act
	actual := args.Map{"r": d.ValueString()}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueString isString", actual)
}

func Test_TD_ValueString_NotString(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(42, true)

	// Act
	actual := args.Map{"r": d.ValueString()}

	// Assert
	expected := args.Map{"r": "42"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueString notString", actual)
}

func Test_TD_ValueInt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(42, true)

	// Act
	actual := args.Map{"r": d.ValueInt()}

	// Assert
	expected := args.Map{"r": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueInt", actual)
}

func Test_TD_ValueInt_Fail(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)

	// Act
	actual := args.Map{"r": d.ValueInt()}

	// Assert
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueInt fail", actual)
}

func Test_TD_ValueInt64(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(int64(99), true)

	// Act
	actual := args.Map{"r": d.ValueInt64()}

	// Assert
	expected := args.Map{"r": int64(99)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueInt64", actual)
}

func Test_TD_ValueBool(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic(true, true)

	// Act
	actual := args.Map{"r": d.ValueBool()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueBool", actual)
}

func Test_TD_ValueBool_Fail(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)

	// Act
	actual := args.Map{"r": d.ValueBool()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueBool fail", actual)
}

// =============================================================================
// TypedDynamic — Clone / NonPtr / Ptr / ToDynamic / Deserialize
// =============================================================================

func Test_TD_Clone(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	c := d.Clone()

	// Act
	actual := args.Map{"data": c.Data()}

	// Assert
	expected := args.Map{"data": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Clone", actual)
}

func Test_TD_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.TypedDynamic[string]

	// Act
	actual := args.Map{"isNil": d.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ClonePtr nil", actual)
}

func Test_TD_ClonePtr_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("x", true)
	c := d.ClonePtr()

	// Act
	actual := args.Map{"notNil": c != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ClonePtr valid", actual)
}

func Test_TD_NonPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)

	// Act
	actual := args.Map{"data": d.NonPtr().Data()}

	// Assert
	expected := args.Map{"data": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic NonPtr", actual)
}

func Test_TD_ToDynamic(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	dyn := d.ToDynamic()

	// Act
	actual := args.Map{"valid": dyn.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ToDynamic", actual)
}

func Test_TD_Deserialize_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := d.Deserialize([]byte(`"world"`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"valid": d.IsValid(),
		"data": d.Data(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"valid": true,
		"data": "world",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize valid", actual)
}

func Test_TD_Deserialize_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.TypedDynamic[string]
	err := d.Deserialize([]byte(`"x"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize nil", actual)
}

func Test_TD_Deserialize_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := d.Deserialize([]byte(`bad`))

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize invalid", actual)
}

func Test_TD_JsonModel(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)

	// Act
	actual := args.Map{"r": d.JsonModel()}

	// Assert
	expected := args.Map{"r": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonModel", actual)
}

func Test_TD_JsonModelAny(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)

	// Act
	actual := args.Map{"r": d.JsonModelAny()}

	// Assert
	expected := args.Map{"r": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonModelAny", actual)
}

func Test_TD_Json(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)
	r := d.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Json", actual)
}

func Test_TD_JsonPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("x", true)

	// Act
	actual := args.Map{"notNil": d.JsonPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonPtr", actual)
}

// =============================================================================
// TypedSimpleResult
// =============================================================================

func Test_TSR_NewValid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid("ok")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"data": r.Data(),
		"result": r.Result(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"data": "ok",
		"result": "ok",
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "TSR NewValid", actual)
}

func Test_TSR_Invalid(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleResult[string]("fail")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"invalid": r.IsInvalid(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
		"msg": "fail",
	}
	expected.ShouldBeEqual(t, 0, "TSR Invalid", actual)
}

func Test_TSR_InvalidNoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleResultNoMessage[int]()

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "TSR InvalidNoMessage", actual)
}

func Test_TSR_IsValid_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"invalid": r.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TSR IsValid nil", actual)
}

func Test_TSR_Message_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{"msg": r.Message()}

	// Assert
	expected := args.Map{"msg": ""}
	expected.ShouldBeEqual(t, 0, "TSR Message nil", actual)
}

func Test_TSR_String_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{"r": r.String()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "TSR String nil", actual)
}

func Test_TSR_String_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid(42)

	// Act
	actual := args.Map{"r": r.String()}

	// Assert
	expected := args.Map{"r": "42"}
	expected.ShouldBeEqual(t, 0, "TSR String valid", actual)
}

func Test_TSR_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{"noErr": r.InvalidError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TSR InvalidError nil", actual)
}

func Test_TSR_InvalidError_NoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResult("x", false, "")

	// Act
	actual := args.Map{"noErr": r.InvalidError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TSR InvalidError no message", actual)
}

func Test_TSR_InvalidError_WithMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResult("x", false, "fail")
	e1 := r.InvalidError()
	e2 := r.InvalidError() // cached

	// Act
	actual := args.Map{
		"hasErr": e1 != nil,
		"same": e1 == e2,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "TSR InvalidError with message cached", actual)
}

func Test_TSR_Clone_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]
	c := r.Clone()

	// Act
	actual := args.Map{"valid": c.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR Clone nil", actual)
}

func Test_TSR_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{"isNil": r.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TSR ClonePtr nil", actual)
}

func Test_TSR_ClonePtr_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid("ok")
	c := r.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"data": c.Data(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"data": "ok",
	}
	expected.ShouldBeEqual(t, 0, "TSR ClonePtr valid", actual)
}

func Test_TSR_ToSimpleResult_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]
	sr := r.ToSimpleResult()

	// Act
	actual := args.Map{"valid": sr.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR ToSimpleResult nil", actual)
}

func Test_TSR_ToTypedDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]
	td := r.ToTypedDynamic()

	// Act
	actual := args.Map{"valid": td.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR ToTypedDynamic nil", actual)
}

func Test_TSR_ToDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]
	d := r.ToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR ToDynamic nil", actual)
}

func Test_TSR_ToSimpleResult_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid("ok")
	sr := r.ToSimpleResult()

	// Act
	actual := args.Map{"valid": sr.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TSR ToSimpleResult valid", actual)
}

func Test_TSR_GetAs(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid("hello")
	s, ok := r.GetAsString()

	// Act
	actual := args.Map{
		"ok": ok,
		"val": s,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TSR GetAsString", actual)
}

func Test_TSR_JSON(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid("hello")
	b, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, merr := r.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil, "hasBytes": len(b) > 0,
		"jrOk": !jr.HasError(), "jOk": !j.HasError(),
		"jpNotNil": jp != nil,
		"mNoErr": merr == nil, "mHasBytes": len(mb) > 0,
		"model":  r.JsonModel(), "modelAny": r.JsonModelAny(),
	}

	// Assert
	expected := args.Map{
		"noErr": true, "hasBytes": true,
		"jrOk": true, "jOk": true,
		"jpNotNil": true,
		"mNoErr": true, "mHasBytes": true,
		"model": "hello", "modelAny": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TSR JSON methods", actual)
}

// =============================================================================
// TypedSimpleRequest
// =============================================================================

func Test_TSReq_NewValid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid("input")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"data": r.Data(),
		"req": r.Request(),
		"val": r.Value(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"data": "input",
		"req": "input",
		"val": "input",
	}
	expected.ShouldBeEqual(t, 0, "TSReq NewValid", actual)
}

func Test_TSReq_Invalid(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleRequest[string]("bad")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"invalid": r.IsInvalid(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
		"msg": "bad",
	}
	expected.ShouldBeEqual(t, 0, "TSReq Invalid", actual)
}

func Test_TSReq_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"invalid": r.IsInvalid(),
		"msg": r.Message(),
		"str": r.String(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
		"msg": "",
		"str": "",
	}
	expected.ShouldBeEqual(t, 0, "TSReq nil accessors", actual)
}

func Test_TSReq_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]

	// Act
	actual := args.Map{"noErr": r.InvalidError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TSReq InvalidError nil", actual)
}

func Test_TSReq_InvalidError_Cached(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequest("x", false, "fail")
	e1 := r.InvalidError()
	e2 := r.InvalidError()

	// Act
	actual := args.Map{
		"hasErr": e1 != nil,
		"same": e1 == e2,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "TSReq InvalidError cached", actual)
}

func Test_TSReq_Clone_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]

	// Act
	actual := args.Map{"isNil": r.Clone() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TSReq Clone nil", actual)
}

func Test_TSReq_Clone_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid("x")
	c := r.Clone()

	// Act
	actual := args.Map{"data": c.Data()}

	// Assert
	expected := args.Map{"data": "x"}
	expected.ShouldBeEqual(t, 0, "TSReq Clone valid", actual)
}

func Test_TSReq_ToSimpleRequest_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]
	sr := r.ToSimpleRequest()

	// Act
	actual := args.Map{"valid": sr.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSReq ToSimpleRequest nil", actual)
}

func Test_TSReq_ToSimpleRequest_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid("x")
	sr := r.ToSimpleRequest()

	// Act
	actual := args.Map{"valid": sr.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TSReq ToSimpleRequest valid", actual)
}

func Test_TSReq_ToTypedDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]
	td := r.ToTypedDynamic()

	// Act
	actual := args.Map{"valid": td.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSReq ToTypedDynamic nil", actual)
}

func Test_TSReq_ToDynamic_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]
	d := r.ToDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSReq ToDynamic nil", actual)
}

func Test_TSReq_GetAs(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	s, ok := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, i64Ok := r.GetAsInt64()
	_, f64Ok := r.GetAsFloat64()
	_, f32Ok := r.GetAsFloat32()
	_, bOk := r.GetAsBool()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()

	// Act
	actual := args.Map{
		"sOk": ok,
		"s": s,
		"iOk": iOk,
		"i64Ok": i64Ok,
		"f64Ok": f64Ok,
		"f32Ok": f32Ok,
		"bOk": bOk,
		"byOk": byOk,
		"ssOk": ssOk,
	}

	// Assert
	expected := args.Map{
		"sOk": true,
		"s": "hello",
		"iOk": false,
		"i64Ok": false,
		"f64Ok": false,
		"f32Ok": false,
		"bOk": false,
		"byOk": false,
		"ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TSReq GetAs all", actual)
}

// =============================================================================
// SimpleRequest
// =============================================================================

func Test_SR_NewValid(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("input")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest NewValid", actual)
}

func Test_SR_Invalid(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequest("fail")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "fail",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest Invalid", actual)
}

func Test_SR_InvalidNoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequestNoMessage()

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidNoMessage", actual)
}

func Test_SR_Request_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{
		"isNil": r.Request() == nil,
		"msg": r.Message(),
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest Request nil", actual)
}

func Test_SR_Value_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"isNil": r.Value() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest Value nil", actual)
}

func Test_SR_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch nil", actual)
}

func Test_SR_GetErrorOnTypeMismatch_Match(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("hello")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch match", actual)
}

func Test_SR_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch mismatch", actual)
}

func Test_SR_GetErrorOnTypeMismatch_Include(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequest(42, true, "extra")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch include", actual)
}

func Test_SR_IsReflectKind_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"r": r.IsReflectKind(reflect.String)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsReflectKind nil", actual)
}

func Test_SR_IsPointer_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"r": r.IsPointer()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer nil", actual)
}

func Test_SR_IsPointer_Value(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("hello")

	// Act
	actual := args.Map{"r": r.IsPointer()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer value", actual)
}

func Test_SR_IsPointer_Cached(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("hello")
	r.IsPointer() // first call caches

	// Act
	actual := args.Map{"r": r.IsPointer()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer cached", actual)
}

func Test_SR_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"noErr": r.InvalidError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError nil", actual)
}

func Test_SR_InvalidError_NoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequest(nil, false, "")

	// Act
	actual := args.Map{"noErr": r.InvalidError() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError no message", actual)
}

func Test_SR_InvalidError_WithMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequest(nil, false, "fail")
	e1 := r.InvalidError()
	e2 := r.InvalidError() // cached

	// Act
	actual := args.Map{
		"hasErr": e1 != nil,
		"same": e1 == e2,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError cached", actual)
}
