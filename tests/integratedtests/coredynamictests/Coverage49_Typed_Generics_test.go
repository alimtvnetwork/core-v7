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

func Test_Cov49_TD_NewTypedDynamic(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	actual := args.Map{"valid": d.IsValid(), "data": d.Data(), "val": d.Value(), "invalid": d.IsInvalid()}
	expected := args.Map{"valid": true, "data": "hello", "val": "hello", "invalid": false}
	expected.ShouldBeEqual(t, 0, "NewTypedDynamic", actual)
}

func Test_Cov49_TD_NewTypedDynamicValid(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid(42)
	actual := args.Map{"valid": d.IsValid(), "data": d.Data()}
	expected := args.Map{"valid": true, "data": 42}
	expected.ShouldBeEqual(t, 0, "NewTypedDynamicValid", actual)
}

func Test_Cov49_TD_NewTypedDynamicPtr(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr("x", true)
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedDynamicPtr", actual)
}

func Test_Cov49_TD_InvalidTypedDynamic(t *testing.T) {
	d := coredynamic.InvalidTypedDynamic[string]()
	actual := args.Map{"valid": d.IsValid(), "invalid": d.IsInvalid()}
	expected := args.Map{"valid": false, "invalid": true}
	expected.ShouldBeEqual(t, 0, "InvalidTypedDynamic", actual)
}

func Test_Cov49_TD_InvalidTypedDynamicPtr(t *testing.T) {
	d := coredynamic.InvalidTypedDynamicPtr[int]()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": false}
	expected.ShouldBeEqual(t, 0, "InvalidTypedDynamicPtr", actual)
}

func Test_Cov49_TD_String(t *testing.T) {
	d := coredynamic.NewTypedDynamic(42, true)
	actual := args.Map{"r": d.String()}
	expected := args.Map{"r": "42"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic String", actual)
}

// =============================================================================
// TypedDynamic — JSON
// =============================================================================

func Test_Cov49_TD_JsonBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	b, err := d.JsonBytes()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonBytes", actual)
}

func Test_Cov49_TD_JsonResult(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	r := d.JsonResult()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonResult", actual)
}

func Test_Cov49_TD_JsonString(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	s, err := d.JsonString()
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": `"hello"`}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonString", actual)
}

func Test_Cov49_TD_MarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	b, err := json.Marshal(d)
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic MarshalJSON", actual)
}

func Test_Cov49_TD_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := json.Unmarshal([]byte(`"world"`), d)
	actual := args.Map{"noErr": err == nil, "valid": d.IsValid(), "data": d.Data()}
	expected := args.Map{"noErr": true, "valid": true, "data": "world"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic UnmarshalJSON", actual)
}

func Test_Cov49_TD_ValueMarshal(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	b, err := d.ValueMarshal()
	actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"noErr": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueMarshal", actual)
}

func Test_Cov49_TD_Bytes_IsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic([]byte("hello"), true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "hasBytes": len(b) > 0}
	expected := args.Map{"ok": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Bytes isBytes", actual)
}

func Test_Cov49_TD_Bytes_NotBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "hasBytes": len(b) > 0}
	expected := args.Map{"ok": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Bytes notBytes", actual)
}

// =============================================================================
// TypedDynamic — GetAs* / Value*
// =============================================================================

func Test_Cov49_TD_GetAsString(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	v, ok := d.GetAsString()
	actual := args.Map{"ok": ok, "v": v}
	expected := args.Map{"ok": true, "v": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsString", actual)
}

func Test_Cov49_TD_GetAsInt(t *testing.T) {
	d := coredynamic.NewTypedDynamic(42, true)
	v, ok := d.GetAsInt()
	actual := args.Map{"ok": ok, "v": v}
	expected := args.Map{"ok": true, "v": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsInt", actual)
}

func Test_Cov49_TD_GetAsInt64(t *testing.T) {
	d := coredynamic.NewTypedDynamic(int64(99), true)
	v, ok := d.GetAsInt64()
	actual := args.Map{"ok": ok, "v": v}
	expected := args.Map{"ok": true, "v": int64(99)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsInt64", actual)
}

func Test_Cov49_TD_GetAsUint(t *testing.T) {
	d := coredynamic.NewTypedDynamic(uint(5), true)
	v, ok := d.GetAsUint()
	actual := args.Map{"ok": ok, "v": v}
	expected := args.Map{"ok": true, "v": uint(5)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsUint", actual)
}

func Test_Cov49_TD_GetAsFloat64(t *testing.T) {
	d := coredynamic.NewTypedDynamic(3.14, true)
	v, ok := d.GetAsFloat64()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsFloat64", actual)
	_ = v
}

func Test_Cov49_TD_GetAsFloat32(t *testing.T) {
	d := coredynamic.NewTypedDynamic(float32(1.5), true)
	_, ok := d.GetAsFloat32()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsFloat32", actual)
}

func Test_Cov49_TD_GetAsBool(t *testing.T) {
	d := coredynamic.NewTypedDynamic(true, true)
	v, ok := d.GetAsBool()
	actual := args.Map{"ok": ok, "v": v}
	expected := args.Map{"ok": true, "v": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsBool", actual)
}

func Test_Cov49_TD_GetAsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic([]byte("hi"), true)
	_, ok := d.GetAsBytes()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsBytes", actual)
}

func Test_Cov49_TD_GetAsStrings(t *testing.T) {
	d := coredynamic.NewTypedDynamic([]string{"a"}, true)
	_, ok := d.GetAsStrings()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAsStrings", actual)
}

func Test_Cov49_TD_ValueString_IsString(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	actual := args.Map{"r": d.ValueString()}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueString isString", actual)
}

func Test_Cov49_TD_ValueString_NotString(t *testing.T) {
	d := coredynamic.NewTypedDynamic(42, true)
	actual := args.Map{"r": d.ValueString()}
	expected := args.Map{"r": "42"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueString notString", actual)
}

func Test_Cov49_TD_ValueInt(t *testing.T) {
	d := coredynamic.NewTypedDynamic(42, true)
	actual := args.Map{"r": d.ValueInt()}
	expected := args.Map{"r": 42}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueInt", actual)
}

func Test_Cov49_TD_ValueInt_Fail(t *testing.T) {
	d := coredynamic.NewTypedDynamic("x", true)
	actual := args.Map{"r": d.ValueInt()}
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueInt fail", actual)
}

func Test_Cov49_TD_ValueInt64(t *testing.T) {
	d := coredynamic.NewTypedDynamic(int64(99), true)
	actual := args.Map{"r": d.ValueInt64()}
	expected := args.Map{"r": int64(99)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueInt64", actual)
}

func Test_Cov49_TD_ValueBool(t *testing.T) {
	d := coredynamic.NewTypedDynamic(true, true)
	actual := args.Map{"r": d.ValueBool()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueBool", actual)
}

func Test_Cov49_TD_ValueBool_Fail(t *testing.T) {
	d := coredynamic.NewTypedDynamic("x", true)
	actual := args.Map{"r": d.ValueBool()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ValueBool fail", actual)
}

// =============================================================================
// TypedDynamic — Clone / NonPtr / Ptr / ToDynamic / Deserialize
// =============================================================================

func Test_Cov49_TD_Clone(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	c := d.Clone()
	actual := args.Map{"data": c.Data()}
	expected := args.Map{"data": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Clone", actual)
}
	d := coredynamic.NewTypedDynamic("x", true)
	actual := args.Map{"data": d.NonPtr().Data()}
	expected := args.Map{"data": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic NonPtr", actual)
}

func Test_Cov49_TD_ToDynamic(t *testing.T) {
	d := coredynamic.NewTypedDynamic("hello", true)
	dyn := d.ToDynamic()
	actual := args.Map{"valid": dyn.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ToDynamic", actual)
}

func Test_Cov49_TD_Deserialize_Valid(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := d.Deserialize([]byte(`"world"`))
	actual := args.Map{"noErr": err == nil, "valid": d.IsValid(), "data": d.Data()}
	expected := args.Map{"noErr": true, "valid": true, "data": "world"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize valid", actual)
}

func Test_Cov49_TD_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.TypedDynamic[string]
	err := d.Deserialize([]byte(`"x"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize nil", actual)
}

func Test_Cov49_TD_Deserialize_Invalid(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := d.Deserialize([]byte(`bad`))
	actual := args.Map{"hasErr": err != nil, "valid": d.IsValid()}
	expected := args.Map{"hasErr": true, "valid": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize invalid", actual)
}

func Test_Cov49_TD_JsonModel(t *testing.T) {
	d := coredynamic.NewTypedDynamic("x", true)
	actual := args.Map{"r": d.JsonModel()}
	expected := args.Map{"r": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonModel", actual)
}

func Test_Cov49_TD_JsonModelAny(t *testing.T) {
	d := coredynamic.NewTypedDynamic("x", true)
	actual := args.Map{"r": d.JsonModelAny()}
	expected := args.Map{"r": "x"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonModelAny", actual)
}

func Test_Cov49_TD_Json(t *testing.T) {
	d := coredynamic.NewTypedDynamic("x", true)
	r := d.Json()
	actual := args.Map{"noErr": !r.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Json", actual)
}

func Test_Cov49_TD_JsonPtr(t *testing.T) {
	d := coredynamic.NewTypedDynamic("x", true)
	actual := args.Map{"notNil": d.JsonPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonPtr", actual)
}

// =============================================================================
// TypedSimpleResult
// =============================================================================

func Test_Cov49_TSR_NewValid(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("ok")
	actual := args.Map{"valid": r.IsValid(), "data": r.Data(), "result": r.Result(), "msg": r.Message()}
	expected := args.Map{"valid": true, "data": "ok", "result": "ok", "msg": ""}
	expected.ShouldBeEqual(t, 0, "TSR NewValid", actual)
}

func Test_Cov49_TSR_Invalid(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleResult[string]("fail")
	actual := args.Map{"valid": r.IsValid(), "invalid": r.IsInvalid(), "msg": r.Message()}
	expected := args.Map{"valid": false, "invalid": true, "msg": "fail"}
	expected.ShouldBeEqual(t, 0, "TSR Invalid", actual)
}

func Test_Cov49_TSR_InvalidNoMessage(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleResultNoMessage[int]()
	actual := args.Map{"valid": r.IsValid(), "msg": r.Message()}
	expected := args.Map{"valid": false, "msg": ""}
	expected.ShouldBeEqual(t, 0, "TSR InvalidNoMessage", actual)
}

func Test_Cov49_TSR_IsValid_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{"valid": r.IsValid(), "invalid": r.IsInvalid()}
	expected := args.Map{"valid": false, "invalid": true}
	expected.ShouldBeEqual(t, 0, "TSR IsValid nil", actual)
}

func Test_Cov49_TSR_Message_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{"msg": r.Message()}
	expected := args.Map{"msg": ""}
	expected.ShouldBeEqual(t, 0, "TSR Message nil", actual)
}

func Test_Cov49_TSR_String_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{"r": r.String()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "TSR String nil", actual)
}

func Test_Cov49_TSR_String_Valid(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid(42)
	actual := args.Map{"r": r.String()}
	expected := args.Map{"r": "42"}
	expected.ShouldBeEqual(t, 0, "TSR String valid", actual)
}

func Test_Cov49_TSR_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{"noErr": r.InvalidError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TSR InvalidError nil", actual)
}

func Test_Cov49_TSR_InvalidError_NoMessage(t *testing.T) {
	r := coredynamic.NewTypedSimpleResult("x", false, "")
	actual := args.Map{"noErr": r.InvalidError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TSR InvalidError no message", actual)
}

func Test_Cov49_TSR_InvalidError_WithMessage(t *testing.T) {
	r := coredynamic.NewTypedSimpleResult("x", false, "fail")
	e1 := r.InvalidError()
	e2 := r.InvalidError() // cached
	actual := args.Map{"hasErr": e1 != nil, "same": e1 == e2}
	expected := args.Map{"hasErr": true, "same": true}
	expected.ShouldBeEqual(t, 0, "TSR InvalidError with message cached", actual)
}

func Test_Cov49_TSR_Clone_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	c := r.Clone()
	actual := args.Map{"valid": c.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR Clone nil", actual)
}
	var r *coredynamic.TypedSimpleResult[string]
	sr := r.ToSimpleResult()
	actual := args.Map{"valid": sr.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR ToSimpleResult nil", actual)
}

func Test_Cov49_TSR_ToTypedDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	td := r.ToTypedDynamic()
	actual := args.Map{"valid": td.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR ToTypedDynamic nil", actual)
}

func Test_Cov49_TSR_ToDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	d := r.ToDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSR ToDynamic nil", actual)
}

func Test_Cov49_TSR_ToSimpleResult_Valid(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("ok")
	sr := r.ToSimpleResult()
	actual := args.Map{"valid": sr.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TSR ToSimpleResult valid", actual)
}

func Test_Cov49_TSR_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	s, ok := r.GetAsString()
	actual := args.Map{"ok": ok, "val": s}
	expected := args.Map{"ok": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "TSR GetAsString", actual)
}

func Test_Cov49_TSR_JSON(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	b, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, merr := r.MarshalJSON()
	actual := args.Map{
		"noErr": err == nil, "hasBytes": len(b) > 0,
		"jrOk": !jr.HasError(), "jOk": !j.HasError(),
		"jpNotNil": jp != nil,
		"mNoErr": merr == nil, "mHasBytes": len(mb) > 0,
		"model":  r.JsonModel(), "modelAny": r.JsonModelAny(),
	}
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

func Test_Cov49_TSReq_NewValid(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("input")
	actual := args.Map{"valid": r.IsValid(), "data": r.Data(), "req": r.Request(), "val": r.Value()}
	expected := args.Map{"valid": true, "data": "input", "req": "input", "val": "input"}
	expected.ShouldBeEqual(t, 0, "TSReq NewValid", actual)
}

func Test_Cov49_TSReq_Invalid(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleRequest[string]("bad")
	actual := args.Map{"valid": r.IsValid(), "invalid": r.IsInvalid(), "msg": r.Message()}
	expected := args.Map{"valid": false, "invalid": true, "msg": "bad"}
	expected.ShouldBeEqual(t, 0, "TSReq Invalid", actual)
}

func Test_Cov49_TSReq_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{"valid": r.IsValid(), "invalid": r.IsInvalid(), "msg": r.Message(), "str": r.String()}
	expected := args.Map{"valid": false, "invalid": true, "msg": "", "str": ""}
	expected.ShouldBeEqual(t, 0, "TSReq nil accessors", actual)
}

func Test_Cov49_TSReq_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{"noErr": r.InvalidError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "TSReq InvalidError nil", actual)
}

func Test_Cov49_TSReq_InvalidError_Cached(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequest("x", false, "fail")
	e1 := r.InvalidError()
	e2 := r.InvalidError()
	actual := args.Map{"hasErr": e1 != nil, "same": e1 == e2}
	expected := args.Map{"hasErr": true, "same": true}
	expected.ShouldBeEqual(t, 0, "TSReq InvalidError cached", actual)
}

func Test_Cov49_TSReq_Clone_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{"isNil": r.Clone() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "TSReq Clone nil", actual)
}

func Test_Cov49_TSReq_Clone_Valid(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("x")
	c := r.Clone()
	actual := args.Map{"data": c.Data()}
	expected := args.Map{"data": "x"}
	expected.ShouldBeEqual(t, 0, "TSReq Clone valid", actual)
}

func Test_Cov49_TSReq_ToSimpleRequest_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	sr := r.ToSimpleRequest()
	actual := args.Map{"valid": sr.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSReq ToSimpleRequest nil", actual)
}

func Test_Cov49_TSReq_ToSimpleRequest_Valid(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("x")
	sr := r.ToSimpleRequest()
	actual := args.Map{"valid": sr.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "TSReq ToSimpleRequest valid", actual)
}

func Test_Cov49_TSReq_ToTypedDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	td := r.ToTypedDynamic()
	actual := args.Map{"valid": td.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSReq ToTypedDynamic nil", actual)
}

func Test_Cov49_TSReq_ToDynamic_Nil(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	d := r.ToDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "TSReq ToDynamic nil", actual)
}

func Test_Cov49_TSReq_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	s, ok := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, i64Ok := r.GetAsInt64()
	_, f64Ok := r.GetAsFloat64()
	_, f32Ok := r.GetAsFloat32()
	_, bOk := r.GetAsBool()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()
	actual := args.Map{"sOk": ok, "s": s, "iOk": iOk, "i64Ok": i64Ok, "f64Ok": f64Ok, "f32Ok": f32Ok, "bOk": bOk, "byOk": byOk, "ssOk": ssOk}
	expected := args.Map{"sOk": true, "s": "hello", "iOk": false, "i64Ok": false, "f64Ok": false, "f32Ok": false, "bOk": false, "byOk": false, "ssOk": false}
	expected.ShouldBeEqual(t, 0, "TSReq GetAs all", actual)
}

// =============================================================================
// SimpleRequest
// =============================================================================

func Test_Cov49_SR_NewValid(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("input")
	actual := args.Map{"valid": r.IsValid(), "msg": r.Message()}
	expected := args.Map{"valid": true, "msg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleRequest NewValid", actual)
}

func Test_Cov49_SR_Invalid(t *testing.T) {
	r := coredynamic.InvalidSimpleRequest("fail")
	actual := args.Map{"valid": r.IsValid(), "msg": r.Message()}
	expected := args.Map{"valid": false, "msg": "fail"}
	expected.ShouldBeEqual(t, 0, "SimpleRequest Invalid", actual)
}

func Test_Cov49_SR_InvalidNoMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleRequestNoMessage()
	actual := args.Map{"valid": r.IsValid(), "msg": r.Message()}
	expected := args.Map{"valid": false, "msg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidNoMessage", actual)
}

func Test_Cov49_SR_Request_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	actual := args.Map{"isNil": r.Request() == nil, "msg": r.Message()}
	expected := args.Map{"isNil": true, "msg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleRequest Request nil", actual)
}

func Test_Cov49_SR_Value_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	actual := args.Map{"isNil": r.Value() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest Value nil", actual)
}

func Test_Cov49_SR_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch nil", actual)
}

func Test_Cov49_SR_GetErrorOnTypeMismatch_Match(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("hello")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch match", actual)
}

func Test_Cov49_SR_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch mismatch", actual)
}

func Test_Cov49_SR_GetErrorOnTypeMismatch_Include(t *testing.T) {
	r := coredynamic.NewSimpleRequest(42, true, "extra")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch include", actual)
}

func Test_Cov49_SR_IsReflectKind_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	actual := args.Map{"r": r.IsReflectKind(reflect.String)}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsReflectKind nil", actual)
}

func Test_Cov49_SR_IsPointer_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	actual := args.Map{"r": r.IsPointer()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer nil", actual)
}

func Test_Cov49_SR_IsPointer_Value(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("hello")
	actual := args.Map{"r": r.IsPointer()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer value", actual)
}

func Test_Cov49_SR_IsPointer_Cached(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("hello")
	r.IsPointer() // first call caches
	actual := args.Map{"r": r.IsPointer()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer cached", actual)
}

func Test_Cov49_SR_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	actual := args.Map{"noErr": r.InvalidError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError nil", actual)
}

func Test_Cov49_SR_InvalidError_NoMessage(t *testing.T) {
	r := coredynamic.NewSimpleRequest(nil, false, "")
	actual := args.Map{"noErr": r.InvalidError() == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError no message", actual)
}

func Test_Cov49_SR_InvalidError_WithMessage(t *testing.T) {
	r := coredynamic.NewSimpleRequest(nil, false, "fail")
	e1 := r.InvalidError()
	e2 := r.InvalidError() // cached
	actual := args.Map{"hasErr": e1 != nil, "same": e1 == e2}
	expected := args.Map{"hasErr": true, "same": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError cached", actual)
}
