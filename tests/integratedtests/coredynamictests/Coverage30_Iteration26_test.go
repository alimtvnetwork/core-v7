package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// TypedDynamic — constructors, accessors, JSON, GetAs*, Value*, Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_TypedDynamic_String_Valid(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	actual := args.Map{
		"data": d.Data(), "value": d.Value(), "valid": d.IsValid(),
		"invalid": d.IsInvalid(), "str": d.String(),
	}
	expected := args.Map{
		"data": "hello", "value": "hello", "valid": true,
		"invalid": false, "str": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- string valid", actual)
}

func Test_I26_TypedDynamic_Invalid(t *testing.T) {
	d := coredynamic.InvalidTypedDynamic[int]()
	actual := args.Map{"valid": d.IsValid(), "invalid": d.IsInvalid(), "data": d.Data()}
	expected := args.Map{"valid": false, "invalid": true, "data": 0}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- invalid int", actual)
}

func Test_I26_TypedDynamic_NewValidPtr(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[string]("ok")
	dp := coredynamic.NewTypedDynamicPtr[int](42, true)
	dip := coredynamic.InvalidTypedDynamicPtr[string]()
	actual := args.Map{
		"validStr": d.IsValid(), "ptrValid": dp.IsValid(), "ptrData": dp.Data(),
		"invPtr": dip.IsInvalid(),
	}
	expected := args.Map{"validStr": true, "ptrValid": true, "ptrData": 42, "invPtr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- constructors", actual)
}

func Test_I26_TypedDynamic_JsonBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("test", true)
	b, err := d.JsonBytes()
	actual := args.Map{"err": err == nil, "bytes": string(b)}
	expected := args.Map{"err": true, "bytes": "\"test\""}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonBytes", actual)
}

func Test_I26_TypedDynamic_JsonResult(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](5, true)
	jr := d.JsonResult()
	j := d.Json()
	jp := d.JsonPtr()
	actual := args.Map{"hasJR": jr.Bytes != nil, "hasJ": j.Bytes != nil, "hasJP": jp != nil}
	expected := args.Map{"hasJR": true, "hasJ": true, "hasJP": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonResult/Json/JsonPtr", actual)
}

func Test_I26_TypedDynamic_JsonString(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	s, err := d.JsonString()
	actual := args.Map{"err": err == nil, "str": s}
	expected := args.Map{"err": true, "str": "\"hello\""}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonString", actual)
}

func Test_I26_TypedDynamic_MarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](99, true)
	b, err := d.MarshalJSON()
	actual := args.Map{"err": err == nil, "val": string(b)}
	expected := args.Map{"err": true, "val": "99"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- MarshalJSON", actual)
}

func Test_I26_TypedDynamic_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](0, false)
	err := d.UnmarshalJSON([]byte("42"))
	actual := args.Map{"err": err == nil, "data": d.Data(), "valid": d.IsValid()}
	expected := args.Map{"err": true, "data": 42, "valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- UnmarshalJSON", actual)
}

func Test_I26_TypedDynamic_ValueMarshal(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("x", true)
	b, err := d.ValueMarshal()
	actual := args.Map{"err": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"err": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueMarshal", actual)
}

func Test_I26_TypedDynamic_Bytes_IsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]byte]([]byte("hi"), true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "val": string(b)}
	expected := args.Map{"ok": true, "val": "hi"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes is bytes", actual)
}

func Test_I26_TypedDynamic_Bytes_NotBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hi", true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "hasBytes": len(b) > 0}
	expected := args.Map{"ok": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes marshalled", actual)
}

func Test_I26_TypedDynamic_GetAsString(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("ok", true)
	v, ok := d.GetAsString()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": "ok", "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsString", actual)
}

func Test_I26_TypedDynamic_GetAsInt(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](7, true)
	v, ok := d.GetAsInt()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": 7, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt", actual)
}

func Test_I26_TypedDynamic_GetAsInt64(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int64](int64(100), true)
	v, ok := d.GetAsInt64()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": int64(100), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt64", actual)
}

func Test_I26_TypedDynamic_GetAsUint(t *testing.T) {
	d := coredynamic.NewTypedDynamic[uint](uint(3), true)
	v, ok := d.GetAsUint()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": uint(3), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsUint", actual)
}

func Test_I26_TypedDynamic_GetAsFloat64(t *testing.T) {
	d := coredynamic.NewTypedDynamic[float64](1.5, true)
	v, ok := d.GetAsFloat64()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": 1.5, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat64", actual)
}

func Test_I26_TypedDynamic_GetAsFloat32(t *testing.T) {
	d := coredynamic.NewTypedDynamic[float32](float32(2.5), true)
	v, ok := d.GetAsFloat32()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": float32(2.5), "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat32", actual)
}

func Test_I26_TypedDynamic_GetAsBool(t *testing.T) {
	d := coredynamic.NewTypedDynamic[bool](true, true)
	v, ok := d.GetAsBool()
	actual := args.Map{"v": v, "ok": ok}
	expected := args.Map{"v": true, "ok": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBool", actual)
}

func Test_I26_TypedDynamic_GetAsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]byte]([]byte{1, 2}, true)
	v, ok := d.GetAsBytes()
	actual := args.Map{"ok": ok, "len": len(v)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBytes", actual)
}

func Test_I26_TypedDynamic_GetAsStrings(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]string]([]string{"a"}, true)
	v, ok := d.GetAsStrings()
	actual := args.Map{"ok": ok, "len": len(v)}
	expected := args.Map{"ok": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsStrings", actual)
}

func Test_I26_TypedDynamic_ValueString(t *testing.T) {
	ds := coredynamic.NewTypedDynamic[string]("hi", true)
	di := coredynamic.NewTypedDynamic[int](5, true)
	actual := args.Map{"str": ds.ValueString(), "int": di.ValueString()}
	expected := args.Map{"str": "hi", "int": "5"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueString", actual)
}

func Test_I26_TypedDynamic_ValueInt(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](10, true)
	dBad := coredynamic.NewTypedDynamic[string]("x", true)
	actual := args.Map{"v": d.ValueInt(), "bad": dBad.ValueInt()}
	expected := args.Map{"v": 10, "bad": -1}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt", actual)
}

func Test_I26_TypedDynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int64](int64(20), true)
	dBad := coredynamic.NewTypedDynamic[string]("x", true)
	actual := args.Map{"v": d.ValueInt64(), "bad": dBad.ValueInt64()}
	expected := args.Map{"v": int64(20), "bad": int64(-1)}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt64", actual)
}

func Test_I26_TypedDynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewTypedDynamic[bool](true, true)
	dBad := coredynamic.NewTypedDynamic[string]("x", true)
	actual := args.Map{"v": d.ValueBool(), "bad": dBad.ValueBool()}
	expected := args.Map{"v": true, "bad": false}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueBool", actual)
}

func Test_I26_TypedDynamic_Clone(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("x", true)
	c := d.Clone()
	actual := args.Map{"data": c.Data(), "valid": c.IsValid()}
	expected := args.Map{"data": "x", "valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Clone", actual)
}

func Test_I26_TypedDynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.TypedDynamic[string]
	cp := d.ClonePtr()
	actual := args.Map{"nil": cp == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns nil -- nil ClonePtr", actual)
}

func Test_I26_TypedDynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](5, true)
	cp := d.ClonePtr()
	actual := args.Map{"data": cp.Data(), "valid": cp.IsValid()}
	expected := args.Map{"data": 5, "valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ClonePtr", actual)
}

func Test_I26_TypedDynamic_NonPtr_Ptr(t *testing.T) {
	d := coredynamic.NewTypedDynamic[int](1, true)
	np := d.NonPtr()
	p := d.Ptr()
	actual := args.Map{"npData": np.Data(), "pData": p.Data()}
	expected := args.Map{"npData": 1, "pData": 1}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- NonPtr/Ptr", actual)
}

func Test_I26_TypedDynamic_ToDynamic(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	dyn := d.ToDynamic()
	actual := args.Map{"valid": dyn.IsValid(), "data": dyn.Data()}
	expected := args.Map{"valid": true, "data": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ToDynamic", actual)
}

func Test_I26_TypedDynamic_Deserialize(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[int](0, false)
	err := d.Deserialize([]byte("77"))
	actual := args.Map{"err": err == nil, "data": d.Data(), "valid": d.IsValid()}
	expected := args.Map{"err": true, "data": 77, "valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Deserialize", actual)
}

func Test_I26_TypedDynamic_Deserialize_Nil(t *testing.T) {
	var d *coredynamic.TypedDynamic[int]
	err := d.Deserialize([]byte("1"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns error -- nil Deserialize", actual)
}

func Test_I26_TypedDynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("m", true)
	actual := args.Map{"model": d.JsonModel(), "any": d.JsonModelAny()}
	expected := args.Map{"model": "m", "any": "m"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModel", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleRequest — constructors, accessors, JSON, GetAs*, Clone, Convert
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_TypedSimpleRequest_Constructors(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleRequest[string]("data", true, "msg")
	r2 := coredynamic.NewTypedSimpleRequestValid[int](10)
	r3 := coredynamic.InvalidTypedSimpleRequest[string]("err")
	r4 := coredynamic.InvalidTypedSimpleRequestNoMessage[int]()
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message(), "r1Data": r1.Data(),
		"r2Valid": r2.IsValid(), "r2Data": r2.Data(),
		"r3Invalid": r3.IsInvalid(), "r3Msg": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}
	expected := args.Map{
		"r1Valid": true, "r1Msg": "msg", "r1Data": "data",
		"r2Valid": true, "r2Data": 10,
		"r3Invalid": true, "r3Msg": "err",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- constructors", actual)
}

func Test_I26_TypedSimpleRequest_NilReceiver(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{
		"valid": r.IsValid(), "invalid": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(),
	}
	expected := args.Map{"valid": false, "invalid": true, "msg": "", "str": ""}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns default -- nil receiver", actual)
}

func Test_I26_TypedSimpleRequest_RequestValue(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	actual := args.Map{"req": r.Request(), "val": r.Value(), "str": r.String()}
	expected := args.Map{"req": "hello", "val": "hello", "str": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Request/Value/String", actual)
}

func Test_I26_TypedSimpleRequest_InvalidError(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleRequestValid[int](1)
	r2 := coredynamic.InvalidTypedSimpleRequest[int]("bad")
	var r3 *coredynamic.TypedSimpleRequest[int]
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
		"r3Nil": r3.InvalidError() == nil,
	}
	expected := args.Map{"r1Nil": true, "r2Err": true, "r3Nil": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- InvalidError", actual)
}

func Test_I26_TypedSimpleRequest_InvalidError_Cached(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleRequest[int]("cached")
	e1 := r.InvalidError()
	e2 := r.InvalidError()
	actual := args.Map{"same": e1 == e2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns cached -- InvalidError cached", actual)
}

func Test_I26_TypedSimpleRequest_Json(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hi")
	b, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, me := r.MarshalJSON()
	actual := args.Map{
		"bytesErr": err == nil, "hasBytes": len(b) > 0,
		"hasJR": jr.Bytes != nil, "hasJ": j.Bytes != nil, "hasJP": jp != nil,
		"marshalErr": me == nil, "marshalLen": len(mb) > 0,
	}
	expected := args.Map{
		"bytesErr": true, "hasBytes": true,
		"hasJR": true, "hasJ": true, "hasJP": true,
		"marshalErr": true, "marshalLen": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JSON methods", actual)
}

func Test_I26_TypedSimpleRequest_JsonModel(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("m")
	actual := args.Map{"model": r.JsonModel(), "any": r.JsonModelAny()}
	expected := args.Map{"model": "m", "any": "m"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonModel", actual)
}

func Test_I26_TypedSimpleRequest_GetAs(t *testing.T) {
	rs := coredynamic.NewTypedSimpleRequestValid[string]("s")
	ri := coredynamic.NewTypedSimpleRequestValid[int](5)
	ri64 := coredynamic.NewTypedSimpleRequestValid[int64](int64(6))
	rf64 := coredynamic.NewTypedSimpleRequestValid[float64](1.1)
	rf32 := coredynamic.NewTypedSimpleRequestValid[float32](float32(2.2))
	rb := coredynamic.NewTypedSimpleRequestValid[bool](true)
	rby := coredynamic.NewTypedSimpleRequestValid[[]byte]([]byte{1})
	rss := coredynamic.NewTypedSimpleRequestValid[[]string]([]string{"a"})

	sv, sok := rs.GetAsString()
	iv, iok := ri.GetAsInt()
	i64v, i64ok := ri64.GetAsInt64()
	f64v, f64ok := rf64.GetAsFloat64()
	f32v, f32ok := rf32.GetAsFloat32()
	bv, bok := rb.GetAsBool()
	byv, byok := rby.GetAsBytes()
	ssv, ssok := rss.GetAsStrings()

	actual := args.Map{
		"sv": sv, "sok": sok, "iv": iv, "iok": iok,
		"i64v": i64v, "i64ok": i64ok, "f64v": f64v, "f64ok": f64ok,
		"f32v": f32v, "f32ok": f32ok, "bv": bv, "bok": bok,
		"byLen": len(byv), "byok": byok, "ssLen": len(ssv), "ssok": ssok,
	}
	expected := args.Map{
		"sv": "s", "sok": true, "iv": 5, "iok": true,
		"i64v": int64(6), "i64ok": true, "f64v": 1.1, "f64ok": true,
		"f32v": float32(2.2), "f32ok": true, "bv": true, "bok": true,
		"byLen": 1, "byok": true, "ssLen": 1, "ssok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAs methods", actual)
}

func Test_I26_TypedSimpleRequest_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("c")
	c := r.Clone()
	var rn *coredynamic.TypedSimpleRequest[string]
	cn := rn.Clone()
	actual := args.Map{"data": c.Data(), "valid": c.IsValid(), "nilClone": cn == nil}
	expected := args.Map{"data": "c", "valid": true, "nilClone": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Clone", actual)
}

func Test_I26_TypedSimpleRequest_ToSimpleRequest(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("sr")
	sr := r.ToSimpleRequest()
	var rn *coredynamic.TypedSimpleRequest[string]
	srn := rn.ToSimpleRequest()
	actual := args.Map{"valid": sr.IsValid(), "nilValid": srn.IsInvalid()}
	expected := args.Map{"valid": true, "nilValid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToSimpleRequest", actual)
}

func Test_I26_TypedSimpleRequest_ToTypedDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("td")
	td := r.ToTypedDynamic()
	var rn *coredynamic.TypedSimpleRequest[string]
	tdn := rn.ToTypedDynamic()
	actual := args.Map{"data": td.Data(), "valid": td.IsValid(), "nilInvalid": tdn.IsInvalid()}
	expected := args.Map{"data": "td", "valid": true, "nilInvalid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToTypedDynamic", actual)
}

func Test_I26_TypedSimpleRequest_ToDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("dyn")
	d := r.ToDynamic()
	var rn *coredynamic.TypedSimpleRequest[string]
	dn := rn.ToDynamic()
	actual := args.Map{"valid": d.IsValid(), "nilInvalid": dn.IsInvalid()}
	expected := args.Map{"valid": true, "nilInvalid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToDynamic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleResult — constructors, accessors, JSON, GetAs*, Clone, Convert
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_TypedSimpleResult_Constructors(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleResult[string]("ok", true, "")
	r2 := coredynamic.NewTypedSimpleResultValid[int](42)
	r3 := coredynamic.InvalidTypedSimpleResult[string]("bad")
	r4 := coredynamic.InvalidTypedSimpleResultNoMessage[int]()
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Data": r1.Data(),
		"r2Valid": r2.IsValid(), "r2Result": r2.Result(),
		"r3Invalid": r3.IsInvalid(), "r3Msg": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}
	expected := args.Map{
		"r1Valid": true, "r1Data": "ok",
		"r2Valid": true, "r2Result": 42,
		"r3Invalid": true, "r3Msg": "bad",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- constructors", actual)
}

func Test_I26_TypedSimpleResult_NilReceiver(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{
		"valid": r.IsValid(), "invalid": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(),
	}
	expected := args.Map{"valid": false, "invalid": true, "msg": "", "str": ""}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns default -- nil receiver", actual)
}

func Test_I26_TypedSimpleResult_InvalidError(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleResultValid[int](1)
	r2 := coredynamic.InvalidTypedSimpleResult[int]("err")
	var r3 *coredynamic.TypedSimpleResult[int]
	e2a := r2.InvalidError()
	e2b := r2.InvalidError() // cached
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil, "r2Err": e2a != nil,
		"r3Nil": r3.InvalidError() == nil, "cached": e2a == e2b,
	}
	expected := args.Map{"r1Nil": true, "r2Err": true, "r3Nil": true, "cached": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- InvalidError", actual)
}

func Test_I26_TypedSimpleResult_Json(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("x")
	b, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, me := r.MarshalJSON()
	actual := args.Map{
		"err": err == nil, "hasBytes": len(b) > 0,
		"hasJR": jr.Bytes != nil, "hasJ": j.Bytes != nil, "hasJP": jp != nil,
		"marshalErr": me == nil, "marshalLen": len(mb) > 0,
	}
	expected := args.Map{
		"err": true, "hasBytes": true,
		"hasJR": true, "hasJ": true, "hasJP": true,
		"marshalErr": true, "marshalLen": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JSON methods", actual)
}

func Test_I26_TypedSimpleResult_JsonModel(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("m")
	actual := args.Map{"model": r.JsonModel(), "any": r.JsonModelAny()}
	expected := args.Map{"model": "m", "any": "m"}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonModel", actual)
}

func Test_I26_TypedSimpleResult_GetAs(t *testing.T) {
	rs := coredynamic.NewTypedSimpleResultValid[string]("s")
	ri := coredynamic.NewTypedSimpleResultValid[int](5)
	ri64 := coredynamic.NewTypedSimpleResultValid[int64](int64(6))
	rf64 := coredynamic.NewTypedSimpleResultValid[float64](1.1)
	rb := coredynamic.NewTypedSimpleResultValid[bool](true)
	rby := coredynamic.NewTypedSimpleResultValid[[]byte]([]byte{1})
	rss := coredynamic.NewTypedSimpleResultValid[[]string]([]string{"a"})

	sv, sok := rs.GetAsString()
	iv, iok := ri.GetAsInt()
	i64v, i64ok := ri64.GetAsInt64()
	f64v, f64ok := rf64.GetAsFloat64()
	bv, bok := rb.GetAsBool()
	byv, byok := rby.GetAsBytes()
	ssv, ssok := rss.GetAsStrings()

	actual := args.Map{
		"sv": sv, "sok": sok, "iv": iv, "iok": iok,
		"i64v": i64v, "i64ok": i64ok, "f64v": f64v, "f64ok": f64ok,
		"bv": bv, "bok": bok,
		"byLen": len(byv), "byok": byok, "ssLen": len(ssv), "ssok": ssok,
	}
	expected := args.Map{
		"sv": "s", "sok": true, "iv": 5, "iok": true,
		"i64v": int64(6), "i64ok": true, "f64v": 1.1, "f64ok": true,
		"bv": true, "bok": true,
		"byLen": 1, "byok": true, "ssLen": 1, "ssok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAs methods", actual)
}

func Test_I26_TypedSimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("c")
	c := r.Clone()
	var rn *coredynamic.TypedSimpleResult[string]
	cn := rn.Clone()
	actual := args.Map{"data": c.Data(), "valid": c.IsValid(), "nilData": cn.Data()}
	expected := args.Map{"data": "c", "valid": true, "nilData": ""}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Clone", actual)
}

func Test_I26_TypedSimpleResult_ClonePtr(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("cp")
	cp := r.ClonePtr()
	var rn *coredynamic.TypedSimpleResult[string]
	cpn := rn.ClonePtr()
	actual := args.Map{"data": cp.Data(), "nilClone": cpn == nil}
	expected := args.Map{"data": "cp", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ClonePtr", actual)
}

func Test_I26_TypedSimpleResult_ToSimpleResult(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("sr")
	sr := r.ToSimpleResult()
	var rn *coredynamic.TypedSimpleResult[string]
	srn := rn.ToSimpleResult()
	actual := args.Map{"valid": sr.IsValid(), "nilValid": srn.IsInvalid()}
	expected := args.Map{"valid": true, "nilValid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToSimpleResult", actual)
}

func Test_I26_TypedSimpleResult_ToTypedDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("td")
	td := r.ToTypedDynamic()
	var rn *coredynamic.TypedSimpleResult[string]
	tdn := rn.ToTypedDynamic()
	actual := args.Map{"data": td.Data(), "valid": td.IsValid(), "nilInvalid": tdn.IsInvalid()}
	expected := args.Map{"data": "td", "valid": true, "nilInvalid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToTypedDynamic", actual)
}

func Test_I26_TypedSimpleResult_ToDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("dyn")
	d := r.ToDynamic()
	var rn *coredynamic.TypedSimpleResult[string]
	dn := rn.ToDynamic()
	actual := args.Map{"valid": d.IsValid(), "nilInvalid": dn.IsInvalid()}
	expected := args.Map{"valid": true, "nilInvalid": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToDynamic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — SafeCastString, CastString, ToBool, ToInt64, ToStrings
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_BytesConverter_SafeCastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	actual := args.Map{"val": bc.SafeCastString(), "empty": bcEmpty.SafeCastString()}
	expected := args.Map{"val": "hello", "empty": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- SafeCastString", actual)
}

func Test_I26_BytesConverter_CastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("ok"))
	s, err := bc.CastString()
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	_, errE := bcEmpty.CastString()
	actual := args.Map{"val": s, "err": err == nil, "emptyErr": errE != nil}
	expected := args.Map{"val": "ok", "err": true, "emptyErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- CastString", actual)
}

func Test_I26_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("true"))
	v, err := bc.ToBool()
	actual := args.Map{"val": v, "err": err == nil}
	expected := args.Map{"val": true, "err": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBool", actual)
}

func Test_I26_BytesConverter_ToBoolMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("false"))
	v := bc.ToBoolMust()
	actual := args.Map{"val": v}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBoolMust", actual)
}

func Test_I26_BytesConverter_ToString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("\"hello\""))
	v, err := bc.ToString()
	actual := args.Map{"val": v, "err": err == nil}
	expected := args.Map{"val": "hello", "err": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToString", actual)
}

func Test_I26_BytesConverter_ToStringMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("\"world\""))
	v := bc.ToStringMust()
	actual := args.Map{"val": v}
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringMust", actual)
}

func Test_I26_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("[\"a\",\"b\"]"))
	v, err := bc.ToStrings()
	actual := args.Map{"len": len(v), "err": err == nil}
	expected := args.Map{"len": 2, "err": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStrings", actual)
}

func Test_I26_BytesConverter_ToStringsMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("[\"x\"]"))
	v := bc.ToStringsMust()
	actual := args.Map{"len": len(v)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringsMust", actual)
}

func Test_I26_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("99"))
	v, err := bc.ToInt64()
	actual := args.Map{"val": v, "err": err == nil}
	expected := args.Map{"val": int64(99), "err": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64", actual)
}

func Test_I26_BytesConverter_ToInt64Must(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("50"))
	v := bc.ToInt64Must()
	actual := args.Map{"val": v}
	expected := args.Map{"val": int64(50)}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64Must", actual)
}

func Test_I26_BytesConverter_Deserialize(t *testing.T) {
	type sample struct {
		Name string `json:"name"`
	}
	bc := coredynamic.NewBytesConverter([]byte(`{"name":"test"}`))
	var s sample
	err := bc.Deserialize(&s)
	actual := args.Map{"err": err == nil, "name": s.Name}
	expected := args.Map{"err": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- Deserialize", actual)
}

func Test_I26_BytesConverter_DeserializeMust(t *testing.T) {
	type sample struct {
		Age int `json:"age"`
	}
	bc := coredynamic.NewBytesConverter([]byte(`{"age":25}`))
	var s sample
	bc.DeserializeMust(&s)
	actual := args.Map{"age": s.Age}
	expected := args.Map{"age": 25}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- DeserializeMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypeStatus — methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_TypeStatus_IsValid_NilReceiver(t *testing.T) {
	var ts *coredynamic.TypeStatus
	actual := args.Map{"valid": ts.IsValid(), "invalid": ts.IsInvalid()}
	expected := args.Map{"valid": false, "invalid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns false -- nil IsValid", actual)
}

func Test_I26_TypeStatus_IsValid_Cached(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	v1 := ts.IsValid()
	v2 := ts.IsValid() // cached
	actual := args.Map{"v1": v1, "v2": v2}
	expected := args.Map{"v1": true, "v2": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns true -- IsValid cached", actual)
}

func Test_I26_TypeStatus_Methods(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	actual := args.Map{
		"same":    ts.IsSame,
		"notSame": ts.IsNotSame(),
		"notEq":   ts.IsNotEqualTypes(),
		"anyPtr":  ts.IsAnyPointer(),
		"bothPtr": ts.IsBothPointer(),
	}
	expected := args.Map{
		"same": true, "notSame": false, "notEq": false,
		"anyPtr": false, "bothPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- methods", actual)
}

func Test_I26_TypeStatus_Names(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", 5)
	actual := args.Map{
		"leftName":      ts.LeftName(),
		"rightName":     ts.RightName(),
		"leftFullName":  ts.LeftFullName(),
		"rightFullName": ts.RightFullName(),
	}
	expected := args.Map{
		"leftName": "string", "rightName": "int",
		"leftFullName": "string", "rightFullName": "int",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- names", actual)
}

func Test_I26_TypeStatus_NilNames(t *testing.T) {
	ts := coredynamic.TypeSameStatus(nil, nil)
	actual := args.Map{
		"leftName": ts.LeftName(), "rightName": ts.RightName(),
		"leftFull": ts.LeftFullName(), "rightFull": ts.RightFullName(),
	}
	expected := args.Map{
		"leftName": "<nil>", "rightName": "<nil>",
		"leftFull": "<nil>", "rightFull": "<nil>",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- nil names", actual)
}

func Test_I26_TypeStatus_NotMatchMessage_Same(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	msg := ts.NotMatchMessage("left", "right")
	actual := args.Map{"empty": msg == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns empty -- NotMatchMessage same", actual)
}

func Test_I26_TypeStatus_NotMatchMessage_Diff(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", 5)
	msg := ts.NotMatchMessage("left", "right")
	actual := args.Map{"hasMsg": msg != ""}
	expected := args.Map{"hasMsg": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns message -- NotMatchMessage diff", actual)
}

func Test_I26_TypeStatus_NotMatchErr(t *testing.T) {
	tsSame := coredynamic.TypeSameStatus("a", "b")
	tsDiff := coredynamic.TypeSameStatus("a", 5)
	actual := args.Map{
		"sameNil": tsSame.NotMatchErr("l", "r") == nil,
		"diffErr": tsDiff.NotMatchErr("l", "r") != nil,
	}
	expected := args.Map{"sameNil": true, "diffErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotMatchErr", actual)
}

func Test_I26_TypeStatus_ValidationError(t *testing.T) {
	tsSame := coredynamic.TypeSameStatus("a", "b")
	tsDiff := coredynamic.TypeSameStatus("a", 5)
	actual := args.Map{
		"sameNil": tsSame.ValidationError() == nil,
		"diffErr": tsDiff.ValidationError() != nil,
	}
	expected := args.Map{"sameNil": true, "diffErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- ValidationError", actual)
}

func Test_I26_TypeStatus_NotEqualSrcDestination(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", 5)
	msg := ts.NotEqualSrcDestinationMessage()
	err := ts.NotEqualSrcDestinationErr()
	actual := args.Map{"hasMsg": msg != "", "hasErr": err != nil}
	expected := args.Map{"hasMsg": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- SrcDestination", actual)
}

func Test_I26_TypeStatus_MustBeSame_NoPanic(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.MustBeSame() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus no panic -- MustBeSame same types", actual)
}

func Test_I26_TypeStatus_MustBeSame_Panics(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", 5)
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- MustBeSame different types", actual)
	}()
	ts.MustBeSame()
}

func Test_I26_TypeStatus_SrcDestinationMustBeSame_NoPanic(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.SrcDestinationMustBeSame() // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus no panic -- SrcDestinationMustBeSame", actual)
}

func Test_I26_TypeStatus_SrcDestinationMustBeSame_Panics(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", 5)
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- SrcDestinationMustBeSame", actual)
	}()
	ts.SrcDestinationMustBeSame()
}

func Test_I26_TypeStatus_IsSameRegardlessPointer(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	actual := args.Map{"same": ts.IsSameRegardlessPointer()}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns true -- IsSameRegardlessPointer", actual)
}

func Test_I26_TypeStatus_IsEqual(t *testing.T) {
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", "b")
	ts3 := coredynamic.TypeSameStatus("a", 5)
	var tsNil *coredynamic.TypeStatus
	actual := args.Map{
		"same":    ts1.IsEqual(&ts2),
		"diff":    ts1.IsEqual(&ts3),
		"nilNil":  tsNil.IsEqual(nil),
		"nilOther": tsNil.IsEqual(&ts1),
	}
	expected := args.Map{"same": true, "diff": false, "nilNil": true, "nilOther": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_ValueStatus_Constructors(t *testing.T) {
	vs1 := coredynamic.InvalidValueStatusNoMessage()
	vs2 := coredynamic.InvalidValueStatus("oops")
	actual := args.Map{
		"vs1Valid": vs1.IsValid, "vs1Msg": vs1.Message,
		"vs2Valid": vs2.IsValid, "vs2Msg": vs2.Message,
	}
	expected := args.Map{
		"vs1Valid": false, "vs1Msg": "",
		"vs2Valid": false, "vs2Msg": "oops",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns correct value -- constructors", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CastedResult — nil receiver methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_CastedResult_NilReceiver(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{
		"invalid":  cr.IsInvalid(),
		"notNull":  cr.IsNotNull(),
		"notPtr":   cr.IsNotPointer(),
		"notMatch": cr.IsNotMatchingAcceptedType(),
		"srcKind":  cr.IsSourceKind(reflect.String),
		"hasErr":   cr.HasError(),
		"issues":   cr.HasAnyIssues(),
	}
	expected := args.Map{
		"invalid": true, "notNull": false, "notPtr": false,
		"notMatch": false, "srcKind": false, "hasErr": false, "issues": true,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult returns default -- nil receiver", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Utility functions — SafeTypeName, ZeroSetAny, TypeNotEqualErr, TypeMustBeSame
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_SafeTypeName(t *testing.T) {
	actual := args.Map{
		"str": coredynamic.SafeTypeName("hello"),
		"int": coredynamic.SafeTypeName(5),
		"nil": coredynamic.SafeTypeName(nil),
	}
	expected := args.Map{"str": "string", "int": "int", "nil": ""}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- various types", actual)
}

func Test_I26_ZeroSetAny_Nil(t *testing.T) {
	coredynamic.ZeroSetAny(nil) // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny no panic -- nil input", actual)
}

func Test_I26_ZeroSetAny_Struct(t *testing.T) {
	type sample struct {
		Name string
		Age  int
	}
	s := &sample{Name: "test", Age: 5}
	coredynamic.ZeroSetAny(s)
	actual := args.Map{"name": s.Name, "age": s.Age}
	expected := args.Map{"name": "", "age": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns zeroed -- struct", actual)
}

func Test_I26_TypeNotEqualErr_Same(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", "b")
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns nil -- same types", actual)
}

func Test_I26_TypeNotEqualErr_Different(t *testing.T) {
	err := coredynamic.TypeNotEqualErr("a", 5)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- different types", actual)
}

func Test_I26_TypeMustBeSame_NoPanic(t *testing.T) {
	coredynamic.TypeMustBeSame("a", "b") // should not panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame no panic -- same types", actual)
}

func Test_I26_TypeMustBeSame_Panics(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- different types", actual)
	}()
	coredynamic.TypeMustBeSame("a", 5)
}

func Test_I26_IsAnyTypesOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{
		"found":    coredynamic.IsAnyTypesOf(strType, strType, intType),
		"notFound": coredynamic.IsAnyTypesOf(strType, intType),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- found/not found", actual)
}

func Test_I26_TypesIndexOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{
		"found":    coredynamic.TypesIndexOf(strType, intType, strType),
		"notFound": coredynamic.TypesIndexOf(strType, intType),
	}
	expected := args.Map{"found": 1, "notFound": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- index", actual)
}

func Test_I26_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal("test")
	actual := args.Map{"kind": rv.Kind().String()}
	expected := args.Map{"kind": "string"}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- string", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectKindValidation, ReflectTypeValidation
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_ReflectKindValidation_Match(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.String, "hi")
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns nil -- match", actual)
}

func Test_I26_ReflectKindValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.Int, "hi")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns error -- mismatch", actual)
}

func Test_I26_ReflectTypeValidation_Match(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), "hello")
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- match", actual)
}

func Test_I26_ReflectTypeValidation_NilNotAllowed(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns error -- nil not allowed", actual)
}

func Test_I26_ReflectTypeValidation_TypeMismatch(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 5)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns error -- type mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectInterfaceVal
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_ReflectInterfaceVal_Value(t *testing.T) {
	result := coredynamic.ReflectInterfaceVal("hello")
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- value type", actual)
}

func Test_I26_ReflectInterfaceVal_Pointer(t *testing.T) {
	s := "hello"
	result := coredynamic.ReflectInterfaceVal(&s)
	actual := args.Map{"val": result}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_LengthOfReflect_Slice(t *testing.T) {
	rv := reflect.ValueOf([]string{"a", "b"})
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- slice", actual)
}

func Test_I26_LengthOfReflect_Array(t *testing.T) {
	rv := reflect.ValueOf([3]int{1, 2, 3})
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- array", actual)
}

func Test_I26_LengthOfReflect_Map(t *testing.T) {
	rv := reflect.ValueOf(map[string]int{"a": 1})
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- map", actual)
}

func Test_I26_LengthOfReflect_Other(t *testing.T) {
	rv := reflect.ValueOf("str")
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 0 -- other kind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr, MustBeAcceptedTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_NotAcceptedTypesErr_Match(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hi", reflect.TypeOf(""))
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns nil -- match", actual)
}

func Test_I26_NotAcceptedTypesErr_NoMatch(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hi", reflect.TypeOf(0))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- no match", actual)
}

func Test_I26_MustBeAcceptedTypes_NoPanic(t *testing.T) {
	coredynamic.MustBeAcceptedTypes("hi", reflect.TypeOf(""))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes no panic -- match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicStatus — constructors, Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_DynamicStatus_Constructors(t *testing.T) {
	ds1 := coredynamic.InvalidDynamicStatusNoMessage()
	ds2 := coredynamic.InvalidDynamicStatus("err")
	actual := args.Map{
		"ds1Valid": ds1.IsValid(), "ds1Msg": ds1.Message,
		"ds2Valid": ds2.IsValid(), "ds2Msg": ds2.Message,
	}
	expected := args.Map{
		"ds1Valid": false, "ds1Msg": "",
		"ds2Valid": false, "ds2Msg": "err",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- constructors", actual)
}

func Test_I26_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("clone")
	c := ds.Clone()
	actual := args.Map{"msg": c.Message, "valid": c.IsValid()}
	expected := args.Map{"msg": "clone", "valid": false}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- Clone", actual)
}

func Test_I26_DynamicStatus_ClonePtr_Nil(t *testing.T) {
	var ds *coredynamic.DynamicStatus
	cp := ds.ClonePtr()
	actual := args.Map{"nil": cp == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns nil -- nil ClonePtr", actual)
}

func Test_I26_DynamicStatus_ClonePtr(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("cp")
	cp := ds.ClonePtr()
	actual := args.Map{"msg": cp.Message, "valid": cp.IsValid()}
	expected := args.Map{"msg": "cp", "valid": false}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- ClonePtr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleRequest — constructors, methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_SimpleRequest_Constructors(t *testing.T) {
	r1 := coredynamic.InvalidSimpleRequestNoMessage()
	r2 := coredynamic.InvalidSimpleRequest("msg")
	r3 := coredynamic.NewSimpleRequest("data", true, "")
	r4 := coredynamic.NewSimpleRequestValid("ok")
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message(),
		"r2Valid": r2.IsValid(), "r2Msg": r2.Message(),
		"r3Valid": r3.IsValid(), "r3Data": r3.Request(),
		"r4Valid": r4.IsValid(), "r4Data": r4.Value(),
	}
	expected := args.Map{
		"r1Valid": false, "r1Msg": "",
		"r2Valid": false, "r2Msg": "msg",
		"r3Valid": true, "r3Data": "data",
		"r4Valid": true, "r4Data": "ok",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- constructors", actual)
}

func Test_I26_SimpleRequest_NilReceiver(t *testing.T) {
	var r *coredynamic.SimpleRequest
	actual := args.Map{
		"msg": r.Message(), "req": r.Request(), "val": r.Value(),
		"pointer": r.IsPointer(), "kind": r.IsReflectKind(reflect.String),
	}
	expected := args.Map{
		"msg": "", "req": nil, "val": nil,
		"pointer": false, "kind": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns default -- nil receiver", actual)
}

func Test_I26_SimpleRequest_InvalidError(t *testing.T) {
	r1 := coredynamic.NewSimpleRequestValid("ok")
	r2 := coredynamic.InvalidSimpleRequest("bad")
	var r3 *coredynamic.SimpleRequest
	e2a := r2.InvalidError()
	e2b := r2.InvalidError() // cached
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil, "r2Err": e2a != nil,
		"r3Nil": r3.InvalidError() == nil, "cached": e2a == e2b,
	}
	expected := args.Map{"r1Nil": true, "r2Err": true, "r3Nil": true, "cached": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- InvalidError", actual)
}

func Test_I26_SimpleRequest_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- nil GetErrorOnTypeMismatch", actual)
}

func Test_I26_SimpleRequest_GetErrorOnTypeMismatch_Match(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- type match", actual)
}

func Test_I26_SimpleRequest_GetErrorOnTypeMismatch_NoMsg(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- type mismatch no msg", actual)
}

func Test_I26_SimpleRequest_GetErrorOnTypeMismatch_WithMsg(t *testing.T) {
	r := coredynamic.NewSimpleRequest("hi", true, "extra")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- type mismatch with msg", actual)
}

func Test_I26_SimpleRequest_IsPointer(t *testing.T) {
	s := "hello"
	r := coredynamic.NewSimpleRequestValid(&s)
	rV := coredynamic.NewSimpleRequestValid("val")
	p1 := r.IsPointer()
	p2 := r.IsPointer() // cached
	actual := args.Map{"ptr": p1, "cached": p2, "nonPtr": rV.IsPointer()}
	expected := args.Map{"ptr": true, "cached": true, "nonPtr": false}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- IsPointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleResult — constructors, methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_SimpleResult_Constructors(t *testing.T) {
	r1 := coredynamic.InvalidSimpleResultNoMessage()
	r2 := coredynamic.InvalidSimpleResult("msg")
	r3 := coredynamic.NewSimpleResultValid("ok")
	r4 := coredynamic.NewSimpleResult("data", true, "info")
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message,
		"r2Valid": r2.IsValid(), "r2Msg": r2.Message,
		"r3Valid": r3.IsValid(), "r3Result": r3.Result,
		"r4Valid": r4.IsValid(), "r4Msg": r4.Message,
	}
	expected := args.Map{
		"r1Valid": false, "r1Msg": "",
		"r2Valid": false, "r2Msg": "msg",
		"r3Valid": true, "r3Result": "ok",
		"r4Valid": true, "r4Msg": "info",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- constructors", actual)
}

func Test_I26_SimpleResult_InvalidError(t *testing.T) {
	r1 := coredynamic.NewSimpleResultValid("ok")
	r2 := coredynamic.InvalidSimpleResult("bad")
	var r3 *coredynamic.SimpleResult
	e2a := r2.InvalidError()
	e2b := r2.InvalidError() // cached
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil, "r2Err": e2a != nil,
		"r3Nil": r3.InvalidError() == nil, "cached": e2a == e2b,
	}
	expected := args.Map{"r1Nil": true, "r2Err": true, "r3Nil": true, "cached": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- InvalidError", actual)
}

func Test_I26_SimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewSimpleResult("c", true, "m")
	c := r.Clone()
	var rn *coredynamic.SimpleResult
	cn := rn.Clone()
	actual := args.Map{"msg": c.Message, "valid": c.IsValid(), "nilMsg": cn.Message}
	expected := args.Map{"msg": "m", "valid": true, "nilMsg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- Clone", actual)
}

func Test_I26_SimpleResult_ClonePtr(t *testing.T) {
	r := coredynamic.NewSimpleResult("c", true, "m")
	cp := r.ClonePtr()
	var rn *coredynamic.SimpleResult
	cpn := rn.ClonePtr()
	actual := args.Map{"msg": cp.Message, "nilClone": cpn == nil}
	expected := args.Map{"msg": "m", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- ClonePtr", actual)
}

func Test_I26_SimpleResult_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- nil GetErrorOnTypeMismatch", actual)
}

func Test_I26_SimpleResult_GetErrorOnTypeMismatch_Match(t *testing.T) {
	r := coredynamic.NewSimpleResultValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- type match", actual)
}

func Test_I26_SimpleResult_GetErrorOnTypeMismatch_NoMsg(t *testing.T) {
	r := coredynamic.NewSimpleResultValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- type mismatch no msg", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAnyItemDiff — various methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_MapAnyItemDiff_NilReceiver(t *testing.T) {
	var m *coredynamic.MapAnyItemDiff
	actual := args.Map{"len": m.Length(), "raw": len(m.Raw())}
	expected := args.Map{"len": 0, "raw": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns 0 -- nil receiver", actual)
}

func Test_I26_MapAnyItemDiff_BasicMethods(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}
	actual := args.Map{
		"len": m.Length(), "empty": m.IsEmpty(), "hasAny": m.HasAnyItem(),
		"lastIdx": m.LastIndex(),
	}
	expected := args.Map{"len": 2, "empty": false, "hasAny": true, "lastIdx": 1}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- basic methods", actual)
}

func Test_I26_MapAnyItemDiff_AllKeysSorted(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"b": 2, "a": 1}
	keys := m.AllKeysSorted()
	actual := args.Map{"first": keys[0], "second": keys[1]}
	expected := args.Map{"first": "a", "second": "b"}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- AllKeysSorted", actual)
}

func Test_I26_MapAnyItemDiff_IsRawEqual(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{
		"same": m.IsRawEqual(false, map[string]any{"a": 1}),
		"diff": m.IsRawEqual(false, map[string]any{"a": 2}),
	}
	expected := args.Map{"same": true, "diff": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual", actual)
}

func Test_I26_MapAnyItemDiff_HasAnyChanges(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	actual := args.Map{
		"noChanges":  m.HasAnyChanges(false, map[string]any{"a": 1}),
		"hasChanges": m.HasAnyChanges(false, map[string]any{"a": 2}),
	}
	expected := args.Map{"noChanges": false, "hasChanges": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyChanges", actual)
}

func Test_I26_MapAnyItemDiff_MapAnyItems(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	mai := m.MapAnyItems()
	actual := args.Map{"notNil": mai != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- MapAnyItems", actual)
}

func Test_I26_MapAnyItemDiff_Clear(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	cleared := m.Clear()
	actual := args.Map{"len": len(cleared)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Clear", actual)
}

func Test_I26_MapAnyItemDiff_Clear_Nil(t *testing.T) {
	var m *coredynamic.MapAnyItemDiff
	cleared := m.Clear()
	actual := args.Map{"len": len(cleared)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- nil Clear", actual)
}

func Test_I26_MapAnyItemDiff_Json(t *testing.T) {
	m := coredynamic.MapAnyItemDiff{"a": 1}
	j := m.Json()
	jp := m.JsonPtr()
	pj := m.PrettyJsonString()
	actual := args.Map{"hasJ": j.Bytes != nil, "hasJP": jp != nil, "hasPJ": pj != ""}
	expected := args.Map{"hasJ": true, "hasJP": true, "hasPJ": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- JSON methods", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectSetFromTo — various paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_ReflectSetFromTo_BothNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo(nil, nil)
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- both nil", actual)
}

func Test_I26_ReflectSetFromTo_SameType(t *testing.T) {
	src := "hello"
	dst := ""
	err := coredynamic.ReflectSetFromTo(&src, &dst)
	actual := args.Map{"err": err == nil, "dst": dst}
	expected := args.Map{"err": true, "dst": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- same pointer type", actual)
}

func Test_I26_ReflectSetFromTo_NonPtrToPtrSameBase(t *testing.T) {
	dst := ""
	err := coredynamic.ReflectSetFromTo("world", &dst)
	actual := args.Map{"err": err == nil, "dst": dst}
	expected := args.Map{"err": true, "dst": "world"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- non-ptr to ptr", actual)
}

func Test_I26_ReflectSetFromTo_BytesToStruct(t *testing.T) {
	type sample struct {
		Name string `json:"name"`
	}
	var s sample
	err := coredynamic.ReflectSetFromTo([]byte(`{"name":"test"}`), &s)
	actual := args.Map{"err": err == nil, "name": s.Name}
	expected := args.Map{"err": true, "name": "test"}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- bytes to struct", actual)
}

func Test_I26_ReflectSetFromTo_StructToBytes(t *testing.T) {
	type sample struct {
		Name string `json:"name"`
	}
	s := sample{Name: "hi"}
	var b []byte
	err := coredynamic.ReflectSetFromTo(s, &b)
	actual := args.Map{"err": err == nil, "hasBytes": len(b) > 0}
	expected := args.Map{"err": true, "hasBytes": true}
	// verify bytes are valid JSON
	var parsed sample
	json.Unmarshal(b, &parsed)
	actual["name"] = parsed.Name
	expected["name"] = "hi"
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- struct to bytes", actual)
}

func Test_I26_ReflectSetFromTo_DestNotPointer(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("from", "notPtr")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns error -- dest not pointer", actual)
}

func Test_I26_ReflectSetFromTo_DestNil(t *testing.T) {
	err := coredynamic.ReflectSetFromTo("from", nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns error -- dest nil", actual)
}

func Test_I26_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	var dst int
	err := coredynamic.ReflectSetFromTo("from", &dst)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns error -- type mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CastTo
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_CastTo_Match(t *testing.T) {
	cr := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	actual := args.Map{
		"valid": cr.IsValid, "match": cr.IsMatchingAcceptedType,
		"null": cr.IsNull, "hasErr": cr.HasError(),
	}
	expected := args.Map{"valid": true, "match": true, "null": false, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- match", actual)
}

func Test_I26_CastTo_NoMatch(t *testing.T) {
	cr := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))
	actual := args.Map{"match": cr.IsMatchingAcceptedType, "hasErr": cr.HasError()}
	expected := args.Map{"match": false, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_I26_MapAsKeyValSlice_Map(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	kv, err := coredynamic.MapAsKeyValSlice(reflect.ValueOf(m))
	actual := args.Map{"err": err == nil, "len": kv.Length()}
	expected := args.Map{"err": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns correct value -- map", actual)
}

func Test_I26_MapAsKeyValSlice_NotMap(t *testing.T) {
	_, err := coredynamic.MapAsKeyValSlice(reflect.ValueOf("not a map"))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns error -- not map", actual)
}
