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

func Test_TypedDynamic_String_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hello", true)

	// Act
	actual := args.Map{
		"data": d.Data(), "value": d.Value(), "valid": d.IsValid(),
		"invalid": d.IsInvalid(), "str": d.String(),
	}

	// Assert
	expected := args.Map{
		"data": "hello", "value": "hello", "valid": true,
		"invalid": false, "str": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- string valid", actual)
}

func Test_TypedDynamic_Invalid_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidTypedDynamic[int]()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"invalid": d.IsInvalid(),
		"data": d.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
		"data": 0,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- invalid int", actual)
}

func Test_TypedDynamic_NewValidPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[string]("ok")
	dp := coredynamic.NewTypedDynamicPtr[int](42, true)
	dip := coredynamic.InvalidTypedDynamicPtr[string]()

	// Act
	actual := args.Map{
		"validStr": d.IsValid(), "ptrValid": dp.IsValid(), "ptrData": dp.Data(),
		"invPtr": dip.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"validStr": true,
		"ptrValid": true,
		"ptrData": 42,
		"invPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- constructors", actual)
}

func Test_TypedDynamic_JsonBytes_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("test", true)
	b, err := d.JsonBytes()

	// Act
	actual := args.Map{
		"err": err == nil,
		"bytes": string(b),
	}

	// Assert
	expected := args.Map{
		"err": true,
		"bytes": "\"test\"",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonBytes", actual)
}

func Test_TypedDynamic_JsonResult_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](5, true)
	jr := d.JsonResult()
	j := d.Json()
	jp := d.JsonPtr()

	// Act
	actual := args.Map{
		"hasJR": jr.Bytes != nil,
		"hasJ": j.Bytes != nil,
		"hasJP": jp != nil,
	}

	// Assert
	expected := args.Map{
		"hasJR": true,
		"hasJ": true,
		"hasJP": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonResult/Json/JsonPtr", actual)
}

func Test_TypedDynamic_JsonString_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	s, err := d.JsonString()

	// Act
	actual := args.Map{
		"err": err == nil,
		"str": s,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"str": "\"hello\"",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonString", actual)
}

func Test_TypedDynamic_MarshalJSON_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](99, true)
	b, err := d.MarshalJSON()

	// Act
	actual := args.Map{
		"err": err == nil,
		"val": string(b),
	}

	// Assert
	expected := args.Map{
		"err": true,
		"val": "99",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- MarshalJSON", actual)
}

func Test_TypedDynamic_UnmarshalJSON_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](0, false)
	err := d.UnmarshalJSON([]byte("42"))

	// Act
	actual := args.Map{
		"err": err == nil,
		"data": d.Data(),
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"err": true,
		"data": 42,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- UnmarshalJSON", actual)
}

func Test_TypedDynamic_ValueMarshal_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("x", true)
	b, err := d.ValueMarshal()

	// Act
	actual := args.Map{
		"err": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueMarshal", actual)
}

func Test_TypedDynamic_Bytes_IsBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[[]byte]([]byte("hi"), true)
	b, ok := d.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"val": string(b),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"val": "hi",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes is bytes", actual)
}

func Test_TypedDynamic_Bytes_NotBytes(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hi", true)
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
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes marshalled", actual)
}

func Test_TypedDynamic_GetAsString_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("ok", true)
	v, ok := d.GetAsString()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": "ok",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsString", actual)
}

func Test_TypedDynamic_GetAsInt_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](7, true)
	v, ok := d.GetAsInt()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": 7,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt", actual)
}

func Test_TypedDynamic_GetAsInt64_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int64](int64(100), true)
	v, ok := d.GetAsInt64()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": int64(100),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsInt64", actual)
}

func Test_TypedDynamic_GetAsUint_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[uint](uint(3), true)
	v, ok := d.GetAsUint()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": uint(3),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsUint", actual)
}

func Test_TypedDynamic_GetAsFloat64_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[float64](1.5, true)
	v, ok := d.GetAsFloat64()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": 1.5,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat64", actual)
}

func Test_TypedDynamic_GetAsFloat32_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[float32](float32(2.5), true)
	v, ok := d.GetAsFloat32()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": float32(2.5),
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsFloat32", actual)
}

func Test_TypedDynamic_GetAsBool_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[bool](true, true)
	v, ok := d.GetAsBool()

	// Act
	actual := args.Map{
		"v": v,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"v": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBool", actual)
}

func Test_TypedDynamic_GetAsBytes_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[[]byte]([]byte{1, 2}, true)
	v, ok := d.GetAsBytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(v),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsBytes", actual)
}

func Test_TypedDynamic_GetAsStrings_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[[]string]([]string{"a"}, true)
	v, ok := d.GetAsStrings()

	// Act
	actual := args.Map{
		"ok": ok,
		"len": len(v),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAsStrings", actual)
}

func Test_TypedDynamic_ValueString_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ds := coredynamic.NewTypedDynamic[string]("hi", true)
	di := coredynamic.NewTypedDynamic[int](5, true)

	// Act
	actual := args.Map{
		"str": ds.ValueString(),
		"int": di.ValueString(),
	}

	// Assert
	expected := args.Map{
		"str": "hi",
		"int": "5",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueString", actual)
}

func Test_TypedDynamic_ValueInt_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](10, true)
	dBad := coredynamic.NewTypedDynamic[string]("x", true)

	// Act
	actual := args.Map{
		"v": d.ValueInt(),
		"bad": dBad.ValueInt(),
	}

	// Assert
	expected := args.Map{
		"v": 10,
		"bad": -1,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt", actual)
}

func Test_TypedDynamic_ValueInt64_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int64](int64(20), true)
	dBad := coredynamic.NewTypedDynamic[string]("x", true)

	// Act
	actual := args.Map{
		"v": d.ValueInt64(),
		"bad": dBad.ValueInt64(),
	}

	// Assert
	expected := args.Map{
		"v": int64(20),
		"bad": int64(-1),
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueInt64", actual)
}

func Test_TypedDynamic_ValueBool_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[bool](true, true)
	dBad := coredynamic.NewTypedDynamic[string]("x", true)

	// Act
	actual := args.Map{
		"v": d.ValueBool(),
		"bad": dBad.ValueBool(),
	}

	// Assert
	expected := args.Map{
		"v": true,
		"bad": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ValueBool", actual)
}

func Test_TypedDynamic_Clone_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("x", true)
	c := d.Clone()

	// Act
	actual := args.Map{
		"data": c.Data(),
		"valid": c.IsValid(),
	}

	// Assert
	expected := args.Map{
		"data": "x",
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Clone", actual)
}

func Test_TypedDynamic_ClonePtr_Nil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var d *coredynamic.TypedDynamic[string]
	cp := d.ClonePtr()

	// Act
	actual := args.Map{"nil": cp == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns nil -- nil ClonePtr", actual)
}

func Test_TypedDynamic_ClonePtr_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](5, true)
	cp := d.ClonePtr()

	// Act
	actual := args.Map{
		"data": cp.Data(),
		"valid": cp.IsValid(),
	}

	// Assert
	expected := args.Map{
		"data": 5,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ClonePtr", actual)
}

func Test_TypedDynamic_NonPtr_Ptr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[int](1, true)
	np := d.NonPtr()
	p := d.Ptr()

	// Act
	actual := args.Map{
		"npData": np.Data(),
		"pData": p.Data(),
	}

	// Assert
	expected := args.Map{
		"npData": 1,
		"pData": 1,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- NonPtr/Ptr", actual)
}

func Test_TypedDynamic_ToDynamic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	dyn := d.ToDynamic()

	// Act
	actual := args.Map{
		"valid": dyn.IsValid(),
		"data": dyn.Data(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"data": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ToDynamic", actual)
}

func Test_TypedDynamic_Deserialize_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[int](0, false)
	err := d.Deserialize([]byte("77"))

	// Act
	actual := args.Map{
		"err": err == nil,
		"data": d.Data(),
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"err": true,
		"data": 77,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Deserialize", actual)
}

func Test_TypedDynamic_Deserialize_Nil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var d *coredynamic.TypedDynamic[int]
	err := d.Deserialize([]byte("1"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns error -- nil Deserialize", actual)
}

func Test_TypedDynamic_JsonModel_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic[string]("m", true)

	// Act
	actual := args.Map{
		"model": d.JsonModel(),
		"any": d.JsonModelAny(),
	}

	// Assert
	expected := args.Map{
		"model": "m",
		"any": "m",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModel", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleRequest — constructors, accessors, JSON, GetAs*, Clone, Convert
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypedSimpleRequest_Constructors_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleRequest[string]("data", true, "msg")
	r2 := coredynamic.NewTypedSimpleRequestValid[int](10)
	r3 := coredynamic.InvalidTypedSimpleRequest[string]("err")
	r4 := coredynamic.InvalidTypedSimpleRequestNoMessage[int]()

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message(), "r1Data": r1.Data(),
		"r2Valid": r2.IsValid(), "r2Data": r2.Data(),
		"r3Invalid": r3.IsInvalid(), "r3Msg": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"r1Valid": true, "r1Msg": "msg", "r1Data": "data",
		"r2Valid": true, "r2Data": 10,
		"r3Invalid": true, "r3Msg": "err",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- constructors", actual)
}

func Test_TypedSimpleRequest_NilReceiver_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleRequest[string]

	// Act
	actual := args.Map{
		"valid": r.IsValid(), "invalid": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
		"msg": "",
		"str": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns default -- nil receiver", actual)
}

func Test_TypedSimpleRequest_RequestValue(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("hello")

	// Act
	actual := args.Map{
		"req": r.Request(),
		"val": r.Value(),
		"str": r.String(),
	}

	// Assert
	expected := args.Map{
		"req": "hello",
		"val": "hello",
		"str": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Request/Value/String", actual)
}

func Test_TypedSimpleRequest_InvalidError_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleRequestValid[int](1)
	r2 := coredynamic.InvalidTypedSimpleRequest[int]("bad")
	var r3 *coredynamic.TypedSimpleRequest[int]

	// Act
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
		"r3Nil": r3.InvalidError() == nil,
	}

	// Assert
	expected := args.Map{
		"r1Nil": true,
		"r2Err": true,
		"r3Nil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- InvalidError", actual)
}

func Test_TypedSimpleRequest_InvalidError_Cached_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidTypedSimpleRequest[int]("cached")
	e1 := r.InvalidError()
	e2 := r.InvalidError()

	// Act
	actual := args.Map{"same": e1 == e2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns cached -- InvalidError cached", actual)
}

func Test_TypedSimpleRequest_Json_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("hi")
	b, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, me := r.MarshalJSON()

	// Act
	actual := args.Map{
		"bytesErr": err == nil, "hasBytes": len(b) > 0,
		"hasJR": jr.Bytes != nil, "hasJ": j.Bytes != nil, "hasJP": jp != nil,
		"marshalErr": me == nil, "marshalLen": len(mb) > 0,
	}

	// Assert
	expected := args.Map{
		"bytesErr": true, "hasBytes": true,
		"hasJR": true, "hasJ": true, "hasJP": true,
		"marshalErr": true, "marshalLen": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JSON methods", actual)
}

func Test_TypedSimpleRequest_JsonModel_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("m")

	// Act
	actual := args.Map{
		"model": r.JsonModel(),
		"any": r.JsonModelAny(),
	}

	// Assert
	expected := args.Map{
		"model": "m",
		"any": "m",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JsonModel", actual)
}

func Test_TypedSimpleRequest_GetAs_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{
		"sv": sv, "sok": sok, "iv": iv, "iok": iok,
		"i64v": i64v, "i64ok": i64ok, "f64v": f64v, "f64ok": f64ok,
		"f32v": f32v, "f32ok": f32ok, "bv": bv, "bok": bok,
		"byLen": len(byv), "byok": byok, "ssLen": len(ssv), "ssok": ssok,
	}

	// Assert
	expected := args.Map{
		"sv": "s", "sok": true, "iv": 5, "iok": true,
		"i64v": int64(6), "i64ok": true, "f64v": 1.1, "f64ok": true,
		"f32v": float32(2.2), "f32ok": true, "bv": true, "bok": true,
		"byLen": 1, "byok": true, "ssLen": 1, "ssok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAs methods", actual)
}

func Test_TypedSimpleRequest_Clone_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("c")
	c := r.Clone()
	var rn *coredynamic.TypedSimpleRequest[string]
	cn := rn.Clone()

	// Act
	actual := args.Map{
		"data": c.Data(),
		"valid": c.IsValid(),
		"nilClone": cn == nil,
	}

	// Assert
	expected := args.Map{
		"data": "c",
		"valid": true,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Clone", actual)
}

func Test_TypedSimpleRequest_ToSimpleRequest_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("sr")
	sr := r.ToSimpleRequest()
	var rn *coredynamic.TypedSimpleRequest[string]
	srn := rn.ToSimpleRequest()

	// Act
	actual := args.Map{
		"valid": sr.IsValid(),
		"nilValid": srn.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"nilValid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToSimpleRequest", actual)
}

func Test_TypedSimpleRequest_ToTypedDynamic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("td")
	td := r.ToTypedDynamic()
	var rn *coredynamic.TypedSimpleRequest[string]
	tdn := rn.ToTypedDynamic()

	// Act
	actual := args.Map{
		"data": td.Data(),
		"valid": td.IsValid(),
		"nilInvalid": tdn.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"data": "td",
		"valid": true,
		"nilInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToTypedDynamic", actual)
}

func Test_TypedSimpleRequest_ToDynamic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("dyn")
	d := r.ToDynamic()
	var rn *coredynamic.TypedSimpleRequest[string]
	dn := rn.ToDynamic()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"nilInvalid": dn.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"nilInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- ToDynamic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypedSimpleResult — constructors, accessors, JSON, GetAs*, Clone, Convert
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypedSimpleResult_Constructors_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleResult[string]("ok", true, "")
	r2 := coredynamic.NewTypedSimpleResultValid[int](42)
	r3 := coredynamic.InvalidTypedSimpleResult[string]("bad")
	r4 := coredynamic.InvalidTypedSimpleResultNoMessage[int]()

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Data": r1.Data(),
		"r2Valid": r2.IsValid(), "r2Result": r2.Result(),
		"r3Invalid": r3.IsInvalid(), "r3Msg": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"r1Valid": true, "r1Data": "ok",
		"r2Valid": true, "r2Result": 42,
		"r3Invalid": true, "r3Msg": "bad",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- constructors", actual)
}

func Test_TypedSimpleResult_NilReceiver_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var r *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{
		"valid": r.IsValid(), "invalid": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
		"msg": "",
		"str": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns default -- nil receiver", actual)
}

func Test_TypedSimpleResult_InvalidError_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleResultValid[int](1)
	r2 := coredynamic.InvalidTypedSimpleResult[int]("err")
	var r3 *coredynamic.TypedSimpleResult[int]
	e2a := r2.InvalidError()
	e2b := r2.InvalidError() // cached

	// Act
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil, "r2Err": e2a != nil,
		"r3Nil": r3.InvalidError() == nil, "cached": e2a == e2b,
	}

	// Assert
	expected := args.Map{
		"r1Nil": true,
		"r2Err": true,
		"r3Nil": true,
		"cached": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- InvalidError", actual)
}

func Test_TypedSimpleResult_Json_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("x")
	b, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, me := r.MarshalJSON()

	// Act
	actual := args.Map{
		"err": err == nil, "hasBytes": len(b) > 0,
		"hasJR": jr.Bytes != nil, "hasJ": j.Bytes != nil, "hasJP": jp != nil,
		"marshalErr": me == nil, "marshalLen": len(mb) > 0,
	}

	// Assert
	expected := args.Map{
		"err": true, "hasBytes": true,
		"hasJR": true, "hasJ": true, "hasJP": true,
		"marshalErr": true, "marshalLen": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JSON methods", actual)
}

func Test_TypedSimpleResult_JsonModel_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("m")

	// Act
	actual := args.Map{
		"model": r.JsonModel(),
		"any": r.JsonModelAny(),
	}

	// Assert
	expected := args.Map{
		"model": "m",
		"any": "m",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JsonModel", actual)
}

func Test_TypedSimpleResult_GetAs_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{
		"sv": sv, "sok": sok, "iv": iv, "iok": iok,
		"i64v": i64v, "i64ok": i64ok, "f64v": f64v, "f64ok": f64ok,
		"bv": bv, "bok": bok,
		"byLen": len(byv), "byok": byok, "ssLen": len(ssv), "ssok": ssok,
	}

	// Assert
	expected := args.Map{
		"sv": "s", "sok": true, "iv": 5, "iok": true,
		"i64v": int64(6), "i64ok": true, "f64v": 1.1, "f64ok": true,
		"bv": true, "bok": true,
		"byLen": 1, "byok": true, "ssLen": 1, "ssok": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAs methods", actual)
}

func Test_TypedSimpleResult_Clone_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("c")
	c := r.Clone()
	var rn *coredynamic.TypedSimpleResult[string]
	cn := rn.Clone()

	// Act
	actual := args.Map{
		"data": c.Data(),
		"valid": c.IsValid(),
		"nilData": cn.Data(),
	}

	// Assert
	expected := args.Map{
		"data": "c",
		"valid": true,
		"nilData": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Clone", actual)
}

func Test_TypedSimpleResult_ClonePtr_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("cp")
	cp := r.ClonePtr()
	var rn *coredynamic.TypedSimpleResult[string]
	cpn := rn.ClonePtr()

	// Act
	actual := args.Map{
		"data": cp.Data(),
		"nilClone": cpn == nil,
	}

	// Assert
	expected := args.Map{
		"data": "cp",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ClonePtr", actual)
}

func Test_TypedSimpleResult_ToSimpleResult_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("sr")
	sr := r.ToSimpleResult()
	var rn *coredynamic.TypedSimpleResult[string]
	srn := rn.ToSimpleResult()

	// Act
	actual := args.Map{
		"valid": sr.IsValid(),
		"nilValid": srn.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"nilValid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToSimpleResult", actual)
}

func Test_TypedSimpleResult_ToTypedDynamic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("td")
	td := r.ToTypedDynamic()
	var rn *coredynamic.TypedSimpleResult[string]
	tdn := rn.ToTypedDynamic()

	// Act
	actual := args.Map{
		"data": td.Data(),
		"valid": td.IsValid(),
		"nilInvalid": tdn.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"data": "td",
		"valid": true,
		"nilInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToTypedDynamic", actual)
}

func Test_TypedSimpleResult_ToDynamic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("dyn")
	d := r.ToDynamic()
	var rn *coredynamic.TypedSimpleResult[string]
	dn := rn.ToDynamic()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"nilInvalid": dn.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"nilInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- ToDynamic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — SafeCastString, CastString, ToBool, ToInt64, ToStrings
// ══════════════════════════════════════════════════════════════════════════════

func Test_BytesConverter_SafeCastString_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})

	// Act
	actual := args.Map{
		"val": bc.SafeCastString(),
		"empty": bcEmpty.SafeCastString(),
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- SafeCastString", actual)
}

func Test_BytesConverter_CastString_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("ok"))
	s, err := bc.CastString()
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	_, errE := bcEmpty.CastString()

	// Act
	actual := args.Map{
		"val": s,
		"err": err == nil,
		"emptyErr": errE != nil,
	}

	// Assert
	expected := args.Map{
		"val": "ok",
		"err": true,
		"emptyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- CastString", actual)
}

func Test_BytesConverter_ToBool_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("true"))
	v, err := bc.ToBool()

	// Act
	actual := args.Map{
		"val": v,
		"err": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"err": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBool", actual)
}

func Test_BytesConverter_ToBoolMust_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("false"))
	v := bc.ToBoolMust()

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBoolMust", actual)
}

func Test_BytesConverter_ToString_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("\"hello\""))
	v, err := bc.ToString()

	// Act
	actual := args.Map{
		"val": v,
		"err": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"err": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToString", actual)
}

func Test_BytesConverter_ToStringMust_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("\"world\""))
	v := bc.ToStringMust()

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringMust", actual)
}

func Test_BytesConverter_ToStrings_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("[\"a\",\"b\"]"))
	v, err := bc.ToStrings()

	// Act
	actual := args.Map{
		"len": len(v),
		"err": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"err": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStrings", actual)
}

func Test_BytesConverter_ToStringsMust_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("[\"x\"]"))
	v := bc.ToStringsMust()

	// Act
	actual := args.Map{"len": len(v)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringsMust", actual)
}

func Test_BytesConverter_ToInt64_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("99"))
	v, err := bc.ToInt64()

	// Act
	actual := args.Map{
		"val": v,
		"err": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int64(99),
		"err": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64", actual)
}

func Test_BytesConverter_ToInt64Must_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("50"))
	v := bc.ToInt64Must()

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": int64(50)}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64Must", actual)
}

func Test_BytesConverter_Deserialize_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	bc := coredynamic.NewBytesConverter([]byte(`{"name":"test"}`))
	var s sample
	err := bc.Deserialize(&s)

	// Act
	actual := args.Map{
		"err": err == nil,
		"name": s.Name,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- Deserialize", actual)
}

func Test_BytesConverter_DeserializeMust_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	type sample struct {
		Age int `json:"age"`
	}
	bc := coredynamic.NewBytesConverter([]byte(`{"age":25}`))
	var s sample
	bc.DeserializeMust(&s)

	// Act
	actual := args.Map{"age": s.Age}

	// Assert
	expected := args.Map{"age": 25}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- DeserializeMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TypeStatus — methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_TypeStatus_IsValid_NilReceiver(t *testing.T) {
	// Arrange
	var ts *coredynamic.TypeStatus

	// Act
	actual := args.Map{
		"valid": ts.IsValid(),
		"invalid": ts.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns false -- nil IsValid", actual)
}

func Test_TypeStatus_IsValid_Cached_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	v1 := ts.IsValid()
	v2 := ts.IsValid() // cached

	// Act
	actual := args.Map{
		"v1": v1,
		"v2": v2,
	}

	// Assert
	expected := args.Map{
		"v1": true,
		"v2": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns true -- IsValid cached", actual)
}

func Test_TypeStatus_Methods_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{
		"same":    ts.IsSame,
		"notSame": ts.IsNotSame(),
		"notEq":   ts.IsNotEqualTypes(),
		"anyPtr":  ts.IsAnyPointer(),
		"bothPtr": ts.IsBothPointer(),
	}

	// Assert
	expected := args.Map{
		"same": true, "notSame": false, "notEq": false,
		"anyPtr": false, "bothPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- methods", actual)
}

func Test_TypeStatus_Names_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 5)

	// Act
	actual := args.Map{
		"leftName":      ts.LeftName(),
		"rightName":     ts.RightName(),
		"leftFullName":  ts.LeftFullName(),
		"rightFullName": ts.RightFullName(),
	}

	// Assert
	expected := args.Map{
		"leftName": "string", "rightName": "int",
		"leftFullName": "string", "rightFullName": "int",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- names", actual)
}

func Test_TypeStatus_NilNames(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, nil)

	// Act
	actual := args.Map{
		"leftName": ts.LeftName(), "rightName": ts.RightName(),
		"leftFull": ts.LeftFullName(), "rightFull": ts.RightFullName(),
	}

	// Assert
	expected := args.Map{
		"leftName": "<nil>", "rightName": "<nil>",
		"leftFull": "<nil>", "rightFull": "<nil>",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- nil names", actual)
}

func Test_TypeStatus_NotMatchMessage_Same_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	msg := ts.NotMatchMessage("left", "right")

	// Act
	actual := args.Map{"empty": msg == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns empty -- NotMatchMessage same", actual)
}

func Test_TypeStatus_NotMatchMessage_Diff_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 5)
	msg := ts.NotMatchMessage("left", "right")

	// Act
	actual := args.Map{"hasMsg": msg != ""}

	// Assert
	expected := args.Map{"hasMsg": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns message -- NotMatchMessage diff", actual)
}

func Test_TypeStatus_NotMatchErr_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	tsSame := coredynamic.TypeSameStatus("a", "b")
	tsDiff := coredynamic.TypeSameStatus("a", 5)

	// Act
	actual := args.Map{
		"sameNil": tsSame.NotMatchErr("l", "r") == nil,
		"diffErr": tsDiff.NotMatchErr("l", "r") != nil,
	}

	// Assert
	expected := args.Map{
		"sameNil": true,
		"diffErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- NotMatchErr", actual)
}

func Test_TypeStatus_ValidationError_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	tsSame := coredynamic.TypeSameStatus("a", "b")
	tsDiff := coredynamic.TypeSameStatus("a", 5)

	// Act
	actual := args.Map{
		"sameNil": tsSame.ValidationError() == nil,
		"diffErr": tsDiff.ValidationError() != nil,
	}

	// Assert
	expected := args.Map{
		"sameNil": true,
		"diffErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- ValidationError", actual)
}

func Test_TypeStatus_NotEqualSrcDestination(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 5)
	msg := ts.NotEqualSrcDestinationMessage()
	err := ts.NotEqualSrcDestinationErr()

	// Act
	actual := args.Map{
		"hasMsg": msg != "",
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasMsg": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- SrcDestination", actual)
}

func Test_TypeStatus_MustBeSame_NoPanic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.MustBeSame() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus no panic -- MustBeSame same types", actual)
}

func Test_TypeStatus_MustBeSame_Panics_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 5)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- MustBeSame different types", actual)
	}()
	ts.MustBeSame()
}

func Test_TypeStatus_SrcDestinationMustBeSame_NoPanic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")
	ts.SrcDestinationMustBeSame() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus no panic -- SrcDestinationMustBeSame", actual)
}

func Test_TypeStatus_SrcDestinationMustBeSame_Panics_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", 5)
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeStatus panics -- SrcDestinationMustBeSame", actual)
	}()
	ts.SrcDestinationMustBeSame()
}

func Test_TypeStatus_IsSameRegardlessPointer(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("a", "b")

	// Act
	actual := args.Map{"same": ts.IsSameRegardlessPointer()}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns true -- IsSameRegardlessPointer", actual)
}

func Test_TypeStatus_IsEqual_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", "b")
	ts3 := coredynamic.TypeSameStatus("a", 5)
	var tsNil *coredynamic.TypeStatus

	// Act
	actual := args.Map{
		"same":    ts1.IsEqual(&ts2),
		"diff":    ts1.IsEqual(&ts3),
		"nilNil":  tsNil.IsEqual(nil),
		"nilOther": tsNil.IsEqual(&ts1),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
		"nilNil": true,
		"nilOther": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ValueStatus — constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValueStatus_Constructors(t *testing.T) {
	// Arrange
	vs1 := coredynamic.InvalidValueStatusNoMessage()
	vs2 := coredynamic.InvalidValueStatus("oops")

	// Act
	actual := args.Map{
		"vs1Valid": vs1.IsValid, "vs1Msg": vs1.Message,
		"vs2Valid": vs2.IsValid, "vs2Msg": vs2.Message,
	}

	// Assert
	expected := args.Map{
		"vs1Valid": false, "vs1Msg": "",
		"vs2Valid": false, "vs2Msg": "oops",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns correct value -- constructors", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CastedResult — nil receiver methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_CastedResult_NilReceiver_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{
		"invalid":  cr.IsInvalid(),
		"notNull":  cr.IsNotNull(),
		"notPtr":   cr.IsNotPointer(),
		"notMatch": cr.IsNotMatchingAcceptedType(),
		"srcKind":  cr.IsSourceKind(reflect.String),
		"hasErr":   cr.HasError(),
		"issues":   cr.HasAnyIssues(),
	}

	// Assert
	expected := args.Map{
		"invalid": true, "notNull": false, "notPtr": false,
		"notMatch": false, "srcKind": false, "hasErr": false, "issues": true,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult returns default -- nil receiver", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Utility functions — SafeTypeName, ZeroSetAny, TypeNotEqualErr, TypeMustBeSame
// ══════════════════════════════════════════════════════════════════════════════

func Test_SafeTypeName_FromTypedDynamicStringIt(t *testing.T) {
	// Act
	actual := args.Map{
		"str": coredynamic.SafeTypeName("hello"),
		"int": coredynamic.SafeTypeName(5),
		"nil": coredynamic.SafeTypeName(nil),
	}

	// Assert
	expected := args.Map{
		"str": "string",
		"int": "int",
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- various types", actual)
}

func Test_ZeroSetAny_Nil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	coredynamic.ZeroSetAny(nil) // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny no panic -- nil input", actual)
}

func Test_ZeroSetAny_Struct(t *testing.T) {
	// Arrange
	type sample struct {
		Name string
		Age  int
	}
	s := &sample{Name: "test", Age: 5}
	coredynamic.ZeroSetAny(s)

	// Act
	actual := args.Map{
		"name": s.Name,
		"age": s.Age,
	}

	// Assert
	expected := args.Map{
		"name": "",
		"age": 0,
	}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns zeroed -- struct", actual)
}

func Test_TypeNotEqualErr_Same_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", "b")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns nil -- same types", actual)
}

func Test_TypeNotEqualErr_Different_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", 5)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- different types", actual)
}

func Test_TypeMustBeSame_NoPanic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	coredynamic.TypeMustBeSame("a", "b") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeMustBeSame no panic -- same types", actual)
}

func Test_TypeMustBeSame_Panics_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"panicked": r != nil}

	// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- different types", actual)
	}()
	coredynamic.TypeMustBeSame("a", 5)
}

func Test_IsAnyTypesOf_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{
		"found":    coredynamic.IsAnyTypesOf(strType, strType, intType),
		"notFound": coredynamic.IsAnyTypesOf(strType, intType),
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- found/not found", actual)
}

func Test_TypesIndexOf_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{
		"found":    coredynamic.TypesIndexOf(strType, intType, strType),
		"notFound": coredynamic.TypesIndexOf(strType, intType),
	}

	// Assert
	expected := args.Map{
		"found": 1,
		"notFound": -1,
	}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- index", actual)
}

func Test_AnyToReflectVal_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	rv := coredynamic.AnyToReflectVal("test")

	// Act
	actual := args.Map{"kind": rv.Kind().String()}

	// Assert
	expected := args.Map{"kind": "string"}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- string", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectKindValidation, ReflectTypeValidation
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectKindValidation_Match_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.String, "hi")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns nil -- match", actual)
}

func Test_ReflectKindValidation_Mismatch_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.Int, "hi")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns error -- mismatch", actual)
}

func Test_ReflectTypeValidation_Match_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns nil -- match", actual)
}

func Test_ReflectTypeValidation_NilNotAllowed_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns error -- nil not allowed", actual)
}

func Test_ReflectTypeValidation_TypeMismatch_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 5)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns error -- type mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectInterfaceVal
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectInterfaceVal_Value_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	result := coredynamic.ReflectInterfaceVal("hello")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- value type", actual)
}

func Test_ReflectInterfaceVal_Pointer_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	s := "hello"
	result := coredynamic.ReflectInterfaceVal(&s)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- pointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ══════════════════════════════════════════════════════════════════════════════

func Test_LengthOfReflect_Slice_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]string{"a", "b"})

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- slice", actual)
}

func Test_LengthOfReflect_Array_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([3]int{1, 2, 3})

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- array", actual)
}

func Test_LengthOfReflect_Map_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(map[string]int{"a": 1})

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- map", actual)
}

func Test_LengthOfReflect_Other_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("str")

	// Act
	actual := args.Map{"len": coredynamic.LengthOfReflect(rv)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 0 -- other kind", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr, MustBeAcceptedTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_NotAcceptedTypesErr_Match_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hi", reflect.TypeOf(""))

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns nil -- match", actual)
}

func Test_NotAcceptedTypesErr_NoMatch_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hi", reflect.TypeOf(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- no match", actual)
}

func Test_MustBeAcceptedTypes_NoPanic_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	coredynamic.MustBeAcceptedTypes("hi", reflect.TypeOf(""))

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes no panic -- match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicStatus — constructors, Clone
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicStatus_Constructors(t *testing.T) {
	// Arrange
	ds1 := coredynamic.InvalidDynamicStatusNoMessage()
	ds2 := coredynamic.InvalidDynamicStatus("err")

	// Act
	actual := args.Map{
		"ds1Valid": ds1.IsValid(), "ds1Msg": ds1.Message,
		"ds2Valid": ds2.IsValid(), "ds2Msg": ds2.Message,
	}

	// Assert
	expected := args.Map{
		"ds1Valid": false, "ds1Msg": "",
		"ds2Valid": false, "ds2Msg": "err",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- constructors", actual)
}

func Test_DynamicStatus_Clone_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("clone")
	c := ds.Clone()

	// Act
	actual := args.Map{
		"msg": c.Message,
		"valid": c.IsValid(),
	}

	// Assert
	expected := args.Map{
		"msg": "clone",
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- Clone", actual)
}

func Test_DynamicStatus_ClonePtr_Nil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var ds *coredynamic.DynamicStatus
	cp := ds.ClonePtr()

	// Act
	actual := args.Map{"nil": cp == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns nil -- nil ClonePtr", actual)
}

func Test_DynamicStatus_ClonePtr_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("cp")
	cp := ds.ClonePtr()

	// Act
	actual := args.Map{
		"msg": cp.Message,
		"valid": cp.IsValid(),
	}

	// Assert
	expected := args.Map{
		"msg": "cp",
		"valid": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- ClonePtr", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleRequest — constructors, methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleRequest_Constructors_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.InvalidSimpleRequestNoMessage()
	r2 := coredynamic.InvalidSimpleRequest("msg")
	r3 := coredynamic.NewSimpleRequest("data", true, "")
	r4 := coredynamic.NewSimpleRequestValid("ok")

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message(),
		"r2Valid": r2.IsValid(), "r2Msg": r2.Message(),
		"r3Valid": r3.IsValid(), "r3Data": r3.Request(),
		"r4Valid": r4.IsValid(), "r4Data": r4.Value(),
	}

	// Assert
	expected := args.Map{
		"r1Valid": false, "r1Msg": "",
		"r2Valid": false, "r2Msg": "msg",
		"r3Valid": true, "r3Data": "data",
		"r4Valid": true, "r4Data": "ok",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- constructors", actual)
}

func Test_SimpleRequest_NilReceiver(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{
		"msg": r.Message(), "req": r.Request(), "val": r.Value(),
		"pointer": r.IsPointer(), "kind": r.IsReflectKind(reflect.String),
	}

	// Assert
	expected := args.Map{
		"msg": "", "req": nil, "val": nil,
		"pointer": false, "kind": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns default -- nil receiver", actual)
}

func Test_SimpleRequest_InvalidError_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewSimpleRequestValid("ok")
	r2 := coredynamic.InvalidSimpleRequest("bad")
	var r3 *coredynamic.SimpleRequest
	e2a := r2.InvalidError()
	e2b := r2.InvalidError() // cached

	// Act
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil, "r2Err": e2a != nil,
		"r3Nil": r3.InvalidError() == nil, "cached": e2a == e2b,
	}

	// Assert
	expected := args.Map{
		"r1Nil": true,
		"r2Err": true,
		"r3Nil": true,
		"cached": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- InvalidError", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- nil GetErrorOnTypeMismatch", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Match_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- type match", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_NoMsg(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- type mismatch no msg", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_WithMsg(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequest("hi", true, "extra")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- type mismatch with msg", actual)
}

func Test_SimpleRequest_IsPointer_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	s := "hello"
	r := coredynamic.NewSimpleRequestValid(&s)
	rV := coredynamic.NewSimpleRequestValid("val")
	p1 := r.IsPointer()
	p2 := r.IsPointer() // cached

	// Act
	actual := args.Map{
		"ptr": p1,
		"cached": p2,
		"nonPtr": rV.IsPointer(),
	}

	// Assert
	expected := args.Map{
		"ptr": true,
		"cached": true,
		"nonPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- IsPointer", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleResult — constructors, methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_SimpleResult_Constructors_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.InvalidSimpleResultNoMessage()
	r2 := coredynamic.InvalidSimpleResult("msg")
	r3 := coredynamic.NewSimpleResultValid("ok")
	r4 := coredynamic.NewSimpleResult("data", true, "info")

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message,
		"r2Valid": r2.IsValid(), "r2Msg": r2.Message,
		"r3Valid": r3.IsValid(), "r3Result": r3.Result,
		"r4Valid": r4.IsValid(), "r4Msg": r4.Message,
	}

	// Assert
	expected := args.Map{
		"r1Valid": false, "r1Msg": "",
		"r2Valid": false, "r2Msg": "msg",
		"r3Valid": true, "r3Result": "ok",
		"r4Valid": true, "r4Msg": "info",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- constructors", actual)
}

func Test_SimpleResult_InvalidError_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewSimpleResultValid("ok")
	r2 := coredynamic.InvalidSimpleResult("bad")
	var r3 *coredynamic.SimpleResult
	e2a := r2.InvalidError()
	e2b := r2.InvalidError() // cached

	// Act
	actual := args.Map{
		"r1Nil": r1.InvalidError() == nil, "r2Err": e2a != nil,
		"r3Nil": r3.InvalidError() == nil, "cached": e2a == e2b,
	}

	// Assert
	expected := args.Map{
		"r1Nil": true,
		"r2Err": true,
		"r3Nil": true,
		"cached": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- InvalidError", actual)
}

func Test_SimpleResult_Clone_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("c", true, "m")
	c := r.Clone()
	var rn *coredynamic.SimpleResult
	cn := rn.Clone()

	// Act
	actual := args.Map{
		"msg": c.Message,
		"valid": c.IsValid(),
		"nilMsg": cn.Message,
	}

	// Assert
	expected := args.Map{
		"msg": "m",
		"valid": true,
		"nilMsg": "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- Clone", actual)
}

func Test_SimpleResult_ClonePtr_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("c", true, "m")
	cp := r.ClonePtr()
	var rn *coredynamic.SimpleResult
	cpn := rn.ClonePtr()

	// Act
	actual := args.Map{
		"msg": cp.Message,
		"nilClone": cpn == nil,
	}

	// Assert
	expected := args.Map{
		"msg": "m",
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- ClonePtr", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch_Nil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- nil GetErrorOnTypeMismatch", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch_Match_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- type match", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch_NoMsg(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid("hi")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- type mismatch no msg", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAnyItemDiff — various methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapAnyItemDiff_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItemDiff

	// Act
	actual := args.Map{
		"len": m.Length(),
		"raw": len(m.Raw()),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"raw": 0,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns 0 -- nil receiver", actual)
}

func Test_MapAnyItemDiff_BasicMethods(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1, "b": 2}

	// Act
	actual := args.Map{
		"len": m.Length(), "empty": m.IsEmpty(), "hasAny": m.HasAnyItem(),
		"lastIdx": m.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"empty": false,
		"hasAny": true,
		"lastIdx": 1,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- basic methods", actual)
}

func Test_MapAnyItemDiff_AllKeysSorted_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"b": 2, "a": 1}
	keys := m.AllKeysSorted()

	// Act
	actual := args.Map{
		"first": keys[0],
		"second": keys[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- AllKeysSorted", actual)
}

func Test_MapAnyItemDiff_IsRawEqual_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{
		"same": m.IsRawEqual(false, map[string]any{"a": 1}),
		"diff": m.IsRawEqual(false, map[string]any{"a": 2}),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- IsRawEqual", actual)
}

func Test_MapAnyItemDiff_HasAnyChanges_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}

	// Act
	actual := args.Map{
		"noChanges":  m.HasAnyChanges(false, map[string]any{"a": 1}),
		"hasChanges": m.HasAnyChanges(false, map[string]any{"a": 2}),
	}

	// Assert
	expected := args.Map{
		"noChanges": false,
		"hasChanges": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- HasAnyChanges", actual)
}

func Test_MapAnyItemDiff_MapAnyItems_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	mai := m.MapAnyItems()

	// Act
	actual := args.Map{"notNil": mai != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- MapAnyItems", actual)
}

func Test_MapAnyItemDiff_Clear_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	cleared := m.Clear()

	// Act
	actual := args.Map{"len": len(cleared)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- Clear", actual)
}

func Test_MapAnyItemDiff_Clear_Nil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItemDiff
	cleared := m.Clear()

	// Act
	actual := args.Map{"len": len(cleared)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- nil Clear", actual)
}

func Test_MapAnyItemDiff_Json_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	m := coredynamic.MapAnyItemDiff{"a": 1}
	j := m.Json()
	jp := m.JsonPtr()
	pj := m.PrettyJsonString()

	// Act
	actual := args.Map{
		"hasJ": j.Bytes != nil,
		"hasJP": jp != nil,
		"hasPJ": pj != "",
	}

	// Assert
	expected := args.Map{
		"hasJ": true,
		"hasJP": true,
		"hasPJ": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAnyItemDiff returns correct value -- JSON methods", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectSetFromTo — various paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReflectSetFromTo_BothNil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(nil, nil)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- both nil", actual)
}

func Test_ReflectSetFromTo_SameType_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	src := "hello"
	dst := ""
	err := coredynamic.ReflectSetFromTo(&src, &dst)

	// Act
	actual := args.Map{
		"err": err == nil,
		"dst": dst,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"dst": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- same pointer type", actual)
}

func Test_ReflectSetFromTo_NonPtrToPtrSameBase(t *testing.T) {
	// Arrange
	dst := ""
	err := coredynamic.ReflectSetFromTo("world", &dst)

	// Act
	actual := args.Map{
		"err": err == nil,
		"dst": dst,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"dst": "world",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- non-ptr to ptr", actual)
}

func Test_ReflectSetFromTo_BytesToStruct_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	var s sample
	err := coredynamic.ReflectSetFromTo([]byte(`{"name":"test"}`), &s)

	// Act
	actual := args.Map{
		"err": err == nil,
		"name": s.Name,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- bytes to struct", actual)
}

func Test_ReflectSetFromTo_StructToBytes_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	s := sample{Name: "hi"}
	var b []byte
	err := coredynamic.ReflectSetFromTo(s, &b)

	// Act
	actual := args.Map{
		"err": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"hasBytes": true,
	}
	// verify bytes are valid JSON
	var parsed sample
	json.Unmarshal(b, &parsed)
	actual["name"] = parsed.Name
	expected["name"] = "hi"
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- struct to bytes", actual)
}

func Test_ReflectSetFromTo_DestNotPointer_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("from", "notPtr")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns error -- dest not pointer", actual)
}

func Test_ReflectSetFromTo_DestNil_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("from", nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns error -- dest nil", actual)
}

func Test_ReflectSetFromTo_TypeMismatch_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	var dst int
	err := coredynamic.ReflectSetFromTo("from", &dst)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns error -- type mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CastTo
// ══════════════════════════════════════════════════════════════════════════════

func Test_CastTo_Match_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	cr := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{
		"valid": cr.IsValid, "match": cr.IsMatchingAcceptedType,
		"null": cr.IsNull, "hasErr": cr.HasError(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"match": true,
		"null": false,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- match", actual)
}

func Test_CastTo_NoMatch(t *testing.T) {
	// Arrange
	cr := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{
		"match": cr.IsMatchingAcceptedType,
		"hasErr": cr.HasError(),
	}

	// Assert
	expected := args.Map{
		"match": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo returns correct value -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapAsKeyValSlice_Map(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1, "b": 2}
	kv, err := coredynamic.MapAsKeyValSlice(reflect.ValueOf(m))

	// Act
	actual := args.Map{
		"err": err == nil,
		"len": kv.Length(),
	}

	// Assert
	expected := args.Map{
		"err": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns correct value -- map", actual)
}

func Test_MapAsKeyValSlice_NotMap_FromTypedDynamicStringIt(t *testing.T) {
	// Arrange
	_, err := coredynamic.MapAsKeyValSlice(reflect.ValueOf("not a map"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns error -- not map", actual)
}
