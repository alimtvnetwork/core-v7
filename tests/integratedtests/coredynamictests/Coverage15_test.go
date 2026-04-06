package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// TypedSimpleRequest[T] — all methods
// ═══════════════════════════════════════════

func Test_Cov15_TypedSimpleRequest_Constructors(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequest("hello", true, "")
	rv := coredynamic.NewTypedSimpleRequestValid("world")
	inv := coredynamic.InvalidTypedSimpleRequest[string]("err")
	invNM := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()
	actual := args.Map{
		"rData":   r.Data(), "rReq": r.Request(), "rVal": r.Value(),
		"rValid":  r.IsValid(), "rInvalid": r.IsInvalid(), "rMsg": r.Message(),
		"rvValid": rv.IsValid(), "rvData": rv.Data(),
		"invValid": inv.IsValid(), "invMsg": inv.Message(),
		"invNMValid": invNM.IsValid(),
	}
	expected := args.Map{
		"rData": "hello", "rReq": "hello", "rVal": "hello",
		"rValid": true, "rInvalid": false, "rMsg": "",
		"rvValid": true, "rvData": "world",
		"invValid": false, "invMsg": "err",
		"invNMValid": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- constructors", actual)
}

func Test_Cov15_TypedSimpleRequest_NilReceiver(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{
		"valid": r.IsValid(), "invalid": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(),
		"err": r.InvalidError() == nil,
	}
	expected := args.Map{
		"valid": false, "invalid": true,
		"msg": "", "str": "", "err": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns nil -- nil", actual)
}

func Test_Cov15_TypedSimpleRequest_InvalidError(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	inv := coredynamic.InvalidTypedSimpleRequest[string]("err")
	err1 := r.InvalidError()
	err2 := inv.InvalidError()
	err3 := inv.InvalidError() // cached
	actual := args.Map{"noErr": err1 == nil, "hasErr": err2 != nil, "cached": err3 == err2}
	expected := args.Map{"noErr": true, "hasErr": true, "cached": true}
	expected.ShouldBeEqual(t, 0, "InvalidError returns error -- with args", actual)
}

func Test_Cov15_TypedSimpleRequest_String(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	actual := args.Map{"v": r.String()}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleRequest_Json(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	jb, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, merr := r.MarshalJSON()
	jm := r.JsonModel()
	jma := r.JsonModelAny()
	actual := args.Map{
		"jbLen": len(jb) > 0, "noErr": err == nil,
		"jrLen": jr.Length() > 0, "jLen": j.Length() > 0, "jpNN": jp != nil,
		"mbLen": len(mb) > 0, "merrNil": merr == nil,
		"jm": jm, "jmaNN": jma != nil,
	}
	expected := args.Map{
		"jbLen": true, "noErr": true,
		"jrLen": true, "jLen": true, "jpNN": true,
		"mbLen": true, "merrNil": true,
		"jm": "hello", "jmaNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Json", actual)
}

func Test_Cov15_TypedSimpleRequest_GetAs(t *testing.T) {
	rs := coredynamic.NewTypedSimpleRequestValid("hello")
	ri := coredynamic.NewTypedSimpleRequestValid(42)
	rb := coredynamic.NewTypedSimpleRequestValid(true)
	rf := coredynamic.NewTypedSimpleRequestValid(3.14)
	str, strOK := rs.GetAsString()
	i, iOK := ri.GetAsInt()
	i64, i64OK := ri.GetAsInt64()
	_, f64OK := rf.GetAsFloat64()
	f32, f32OK := rf.GetAsFloat32()
	b, bOK := rb.GetAsBool()
	_, bytesOK := rs.GetAsBytes()
	_, strsOK := rs.GetAsStrings()
	actual := args.Map{
		"str": str, "strOK": strOK, "i": i, "iOK": iOK,
		"i64": i64, "i64OK": i64OK, "f64OK": f64OK,
		"f32": f32, "f32OK": f32OK, "b": b, "bOK": bOK,
		"bytesOK": bytesOK, "strsOK": strsOK,
	}
	expected := args.Map{
		"str": "hello", "strOK": true, "i": 42, "iOK": true,
		"i64": int64(0), "i64OK": false, "f64OK": true,
		"f32": float32(0), "f32OK": false, "b": true, "bOK": true,
		"bytesOK": false, "strsOK": false,
	}
	expected.ShouldBeEqual(t, 0, "GetAs returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleRequest_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	cloned := r.Clone()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilClone := nilR.Clone()
	actual := args.Map{"cData": cloned.Data(), "nilNil": nilClone == nil}
	expected := args.Map{"cData": "hello", "nilNil": true}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleRequest_ToSimpleRequest(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	sr := r.ToSimpleRequest()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilSR := nilR.ToSimpleRequest()
	actual := args.Map{"valid": sr.IsValid(), "nilValid": nilSR.IsValid()}
	expected := args.Map{"valid": true, "nilValid": false}
	expected.ShouldBeEqual(t, 0, "ToSimpleRequest returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleRequest_ToTypedDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	td := r.ToTypedDynamic()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilTD := nilR.ToTypedDynamic()
	actual := args.Map{"data": td.Data(), "valid": td.IsValid(), "nilValid": nilTD.IsValid()}
	expected := args.Map{"data": "hello", "valid": true, "nilValid": false}
	expected.ShouldBeEqual(t, 0, "ToTypedDynamic returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleRequest_ToDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid("hello")
	d := r.ToDynamic()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilD := nilR.ToDynamic()
	actual := args.Map{"data": d.Data(), "valid": d.IsValid(), "nilValid": nilD.IsValid()}
	expected := args.Map{"data": "hello", "valid": true, "nilValid": false}
	expected.ShouldBeEqual(t, 0, "ToDynamic returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// TypedSimpleResult[T] — all methods
// ═══════════════════════════════════════════

func Test_Cov15_TypedSimpleResult_Constructors(t *testing.T) {
	r := coredynamic.NewTypedSimpleResult("hello", true, "")
	rv := coredynamic.NewTypedSimpleResultValid("world")
	inv := coredynamic.InvalidTypedSimpleResult[string]("err")
	invNM := coredynamic.InvalidTypedSimpleResultNoMessage[string]()
	actual := args.Map{
		"rData": r.Data(), "rResult": r.Result(),
		"rValid": r.IsValid(), "rInvalid": r.IsInvalid(), "rMsg": r.Message(),
		"rvValid": rv.IsValid(),
		"invValid": inv.IsValid(), "invMsg": inv.Message(),
		"invNMValid": invNM.IsValid(),
	}
	expected := args.Map{
		"rData": "hello", "rResult": "hello",
		"rValid": true, "rInvalid": false, "rMsg": "",
		"rvValid": true,
		"invValid": false, "invMsg": "err",
		"invNMValid": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- constructors", actual)
}

func Test_Cov15_TypedSimpleResult_NilReceiver(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{
		"valid": r.IsValid(), "invalid": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(),
		"err": r.InvalidError() == nil,
	}
	expected := args.Map{
		"valid": false, "invalid": true,
		"msg": "", "str": "", "err": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns nil -- nil", actual)
}

func Test_Cov15_TypedSimpleResult_InvalidError(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	inv := coredynamic.InvalidTypedSimpleResult[string]("err")
	err1 := r.InvalidError()
	err2 := inv.InvalidError()
	err3 := inv.InvalidError() // cached
	actual := args.Map{"noErr": err1 == nil, "hasErr": err2 != nil, "cached": err3 == err2}
	expected := args.Map{"noErr": true, "hasErr": true, "cached": true}
	expected.ShouldBeEqual(t, 0, "InvalidError returns error -- with args", actual)
}

func Test_Cov15_TypedSimpleResult_String(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	actual := args.Map{"v": r.String()}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleResult_Json(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	jb, err := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mb, merr := r.MarshalJSON()
	jm := r.JsonModel()
	jma := r.JsonModelAny()
	actual := args.Map{
		"jbLen": len(jb) > 0, "noErr": err == nil,
		"jrLen": jr.Length() > 0, "jLen": j.Length() > 0, "jpNN": jp != nil,
		"mbLen": len(mb) > 0, "merrNil": merr == nil,
		"jm": jm, "jmaNN": jma != nil,
	}
	expected := args.Map{
		"jbLen": true, "noErr": true,
		"jrLen": true, "jLen": true, "jpNN": true,
		"mbLen": true, "merrNil": true,
		"jm": "hello", "jmaNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Json", actual)
}

func Test_Cov15_TypedSimpleResult_GetAs(t *testing.T) {
	rs := coredynamic.NewTypedSimpleResultValid("hello")
	ri := coredynamic.NewTypedSimpleResultValid(42)
	rb := coredynamic.NewTypedSimpleResultValid(true)
	rf := coredynamic.NewTypedSimpleResultValid(3.14)
	str, strOK := rs.GetAsString()
	i, iOK := ri.GetAsInt()
	i64, i64OK := ri.GetAsInt64()
	_, f64OK := rf.GetAsFloat64()
	b, bOK := rb.GetAsBool()
	_, bytesOK := rs.GetAsBytes()
	_, strsOK := rs.GetAsStrings()
	actual := args.Map{
		"str": str, "strOK": strOK, "i": i, "iOK": iOK,
		"i64": i64, "i64OK": i64OK, "f64OK": f64OK,
		"b": b, "bOK": bOK,
		"bytesOK": bytesOK, "strsOK": strsOK,
	}
	expected := args.Map{
		"str": "hello", "strOK": true, "i": 42, "iOK": true,
		"i64": int64(0), "i64OK": false, "f64OK": true,
		"b": true, "bOK": true,
		"bytesOK": false, "strsOK": false,
	}
	expected.ShouldBeEqual(t, 0, "GetAs returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	cloned := r.Clone()
	clonedPtr := r.ClonePtr()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilClone := nilR.ClonePtr()
	nilCloneVal := nilR.Clone()
	actual := args.Map{
		"cData": cloned.Data(), "cpNN": clonedPtr != nil,
		"nilNil": nilClone == nil, "nilValid": nilCloneVal.IsValid(),
	}
	expected := args.Map{
		"cData": "hello", "cpNN": true,
		"nilNil": true, "nilValid": false,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleResult_ToSimpleResult(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	sr := r.ToSimpleResult()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilSR := nilR.ToSimpleResult()
	actual := args.Map{"valid": sr.IsValid(), "nilValid": nilSR.IsValid()}
	expected := args.Map{"valid": true, "nilValid": false}
	expected.ShouldBeEqual(t, 0, "ToSimpleResult returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleResult_ToTypedDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	td := r.ToTypedDynamic()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilTD := nilR.ToTypedDynamic()
	actual := args.Map{"data": td.Data(), "valid": td.IsValid(), "nilValid": nilTD.IsValid()}
	expected := args.Map{"data": "hello", "valid": true, "nilValid": false}
	expected.ShouldBeEqual(t, 0, "ToTypedDynamic returns correct value -- with args", actual)
}

func Test_Cov15_TypedSimpleResult_ToDynamic(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid("hello")
	d := r.ToDynamic()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilD := nilR.ToDynamic()
	actual := args.Map{"data": d.Data(), "valid": d.IsValid(), "nilValid": nilD.IsValid()}
	expected := args.Map{"data": "hello", "valid": true, "nilValid": false}
	expected.ShouldBeEqual(t, 0, "ToDynamic returns correct value -- with args", actual)
}
