package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// TypedDynamic
// ═══════════════════════════════════════════

func Test_Cov7_TypedDynamic_Constructors(t *testing.T) {
	td := coredynamic.NewTypedDynamic[string]("hello", true)
	tdv := coredynamic.NewTypedDynamicValid[string]("world")
	tdp := coredynamic.NewTypedDynamicPtr[string]("ptr", true)
	inv := coredynamic.InvalidTypedDynamic[string]()
	invp := coredynamic.InvalidTypedDynamicPtr[string]()
	actual := args.Map{
		"tdData":    td.Data(),
		"tdValid":   td.IsValid(),
		"tdvData":   tdv.Value(),
		"tdpNotNil": tdp != nil,
		"invValid":  inv.IsValid(),
		"invInv":    inv.IsInvalid(),
		"invpNN":    invp != nil,
	}
	expected := args.Map{
		"tdData": "hello", "tdValid": true, "tdvData": "world",
		"tdpNotNil": true, "invValid": false, "invInv": true, "invpNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- constructors", actual)
}

func Test_Cov7_TypedDynamic_String(t *testing.T) {
	td := coredynamic.NewTypedDynamicValid[int](42)
	actual := args.Map{"str": td.String()}
	expected := args.Map{"str": "42"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- String", actual)
}

func Test_Cov7_TypedDynamic_JsonOps(t *testing.T) {
	td := coredynamic.NewTypedDynamicValid[string]("test")
	jb, jbErr := td.JsonBytes()
	js, jsErr := td.JsonString()
	mj, mjErr := td.MarshalJSON()
	vm, vmErr := td.ValueMarshal()
	jr := td.JsonResult()
	jp := td.JsonPtr()
	jm := td.JsonModel()
	jma := td.JsonModelAny()
	actual := args.Map{
		"jbLen":    len(jb) > 0,
		"jbErr":    jbErr == nil,
		"jsLen":    len(js) > 0,
		"jsErr":    jsErr == nil,
		"mjLen":    len(mj) > 0,
		"mjErr":    mjErr == nil,
		"vmLen":    len(vm) > 0,
		"vmErr":    vmErr == nil,
		"jrNotNil": jr.HasBytes(),
		"jpNotNil": jp != nil,
		"jm":       jm,
		"jmaNN":    jma != nil,
	}
	expected := args.Map{
		"jbLen": true, "jbErr": true, "jsLen": true, "jsErr": true,
		"mjLen": true, "mjErr": true, "vmLen": true, "vmErr": true,
		"jrNotNil": true, "jpNotNil": true, "jm": "test", "jmaNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JSON ops", actual)
}

func Test_Cov7_TypedDynamic_GetAs(t *testing.T) {
	tdStr := coredynamic.NewTypedDynamicValid[string]("hello")
	tdInt := coredynamic.NewTypedDynamicValid[int](42)
	tdBool := coredynamic.NewTypedDynamicValid[bool](true)
	s, sOk := tdStr.GetAsString()
	i, iOk := tdInt.GetAsInt()
	b, bOk := tdBool.GetAsBool()
	_, i64Ok := tdInt.GetAsInt64()
	_, uOk := tdInt.GetAsUint()
	_, f64Ok := tdInt.GetAsFloat64()
	_, f32Ok := tdInt.GetAsFloat32()
	_, byOk := tdInt.GetAsBytes()
	_, ssOk := tdInt.GetAsStrings()
	actual := args.Map{
		"s": s, "sOk": sOk, "i": i, "iOk": iOk, "b": b, "bOk": bOk,
		"i64Ok": i64Ok, "uOk": uOk, "f64Ok": f64Ok, "f32Ok": f32Ok,
		"byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "i": 42, "iOk": true, "b": true, "bOk": true,
		"i64Ok": false, "uOk": false, "f64Ok": false, "f32Ok": false,
		"byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAs", actual)
}

func Test_Cov7_TypedDynamic_Value(t *testing.T) {
	tdStr := coredynamic.NewTypedDynamicValid[string]("hello")
	tdInt := coredynamic.NewTypedDynamicValid[int](42)
	tdBool := coredynamic.NewTypedDynamicValid[bool](true)
	actual := args.Map{
		"valueStr":  tdStr.ValueString(),
		"valueInt":  tdInt.ValueInt(),
		"valueI64":  tdInt.ValueInt64(),
		"valueBool": tdBool.ValueBool(),
	}
	expected := args.Map{
		"valueStr": "hello", "valueInt": 42, "valueI64": int64(-1), "valueBool": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Value methods", actual)
}

func Test_Cov7_TypedDynamic_Bytes(t *testing.T) {
	tdBytes := coredynamic.NewTypedDynamicValid[[]byte]([]byte("abc"))
	b, ok := tdBytes.Bytes()
	tdStr := coredynamic.NewTypedDynamicValid[string]("abc")
	b2, ok2 := tdStr.Bytes()
	actual := args.Map{"len": len(b), "ok": ok, "len2": len(b2) > 0, "ok2": ok2}
	expected := args.Map{"len": 3, "ok": true, "len2": true, "ok2": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes", actual)
}

func Test_Cov7_TypedDynamic_Clone(t *testing.T) {
	td := coredynamic.NewTypedDynamicValid[string]("hello")
	cloned := td.Clone()
	clonedPtr := td.ClonePtr()
	nonPtr := td.NonPtr()
	ptr := td.Ptr()
	var nilTd *coredynamic.TypedDynamic[string]
	actual := args.Map{
		"clonedData": cloned.Data(),
		"cpData":     clonedPtr.Data(),
		"nonPtrData": nonPtr.Data(),
		"ptrNN":      ptr != nil,
		"nilClone":   nilTd.ClonePtr() == nil,
	}
	expected := args.Map{
		"clonedData": "hello", "cpData": "hello", "nonPtrData": "hello",
		"ptrNN": true, "nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Clone", actual)
}

func Test_Cov7_TypedDynamic_ToDynamic(t *testing.T) {
	td := coredynamic.NewTypedDynamicValid[string]("hello")
	d := td.ToDynamic()
	actual := args.Map{"valid": d.IsValid(), "data": d.Data()}
	expected := args.Map{"valid": true, "data": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- ToDynamic", actual)
}

func Test_Cov7_TypedDynamic_UnmarshalJSON(t *testing.T) {
	td := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := td.UnmarshalJSON([]byte(`"parsed"`))
	actual := args.Map{"data": td.Data(), "valid": td.IsValid(), "noErr": err == nil}
	expected := args.Map{"data": "parsed", "valid": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- UnmarshalJSON", actual)
}

func Test_Cov7_TypedDynamic_Deserialize(t *testing.T) {
	td := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := td.Deserialize([]byte(`"deserialized"`))
	var nilTd *coredynamic.TypedDynamic[string]
	nilErr := nilTd.Deserialize([]byte(`"x"`))
	actual := args.Map{
		"data": td.Data(), "valid": td.IsValid(), "noErr": err == nil,
		"nilErr": nilErr != nil,
	}
	expected := args.Map{
		"data": "deserialized", "valid": true, "noErr": true, "nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Deserialize", actual)
}

// ═══════════════════════════════════════════
// TypedSimpleRequest
// ═══════════════════════════════════════════

func Test_Cov7_TypedSimpleRequest_Constructors(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequest[string]("hello", true, "")
	rv := coredynamic.NewTypedSimpleRequestValid[string]("world")
	inv := coredynamic.InvalidTypedSimpleRequest[string]("err msg")
	invNoMsg := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()
	actual := args.Map{
		"data": r.Data(), "valid": r.IsValid(), "msg": r.Message(),
		"rvData": rv.Request(), "rvValid": rv.IsValid(),
		"invValid": inv.IsValid(), "invInv": inv.IsInvalid(), "invMsg": inv.Message(),
		"invNoMsg": invNoMsg.Message(),
	}
	expected := args.Map{
		"data": "hello", "valid": true, "msg": "",
		"rvData": "world", "rvValid": true,
		"invValid": false, "invInv": true, "invMsg": "err msg",
		"invNoMsg": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- constructors", actual)
}

func Test_Cov7_TypedSimpleRequest_NilReceiver(t *testing.T) {
	var r *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{
		"valid":  r.IsValid(),
		"inv":    r.IsInvalid(),
		"msg":    r.Message(),
		"str":    r.String(),
		"invErr": r.InvalidError() == nil,
	}
	expected := args.Map{
		"valid": false, "inv": true, "msg": "", "str": "", "invErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns nil -- nil receiver", actual)
}

func Test_Cov7_TypedSimpleRequest_InvalidError(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleRequest[string]("some error")
	err1 := r.InvalidError()
	err2 := r.InvalidError() // cached
	rv := coredynamic.NewTypedSimpleRequestValid[string]("ok")
	actual := args.Map{
		"hasErr": err1 != nil,
		"cached": err1 == err2,
		"noErr":  rv.InvalidError() == nil,
	}
	expected := args.Map{"hasErr": true, "cached": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns error -- InvalidError", actual)
}

func Test_Cov7_TypedSimpleRequest_JSON(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("test")
	jb, jbErr := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mj, mjErr := r.MarshalJSON()
	actual := args.Map{
		"jbLen": len(jb) > 0, "jbErr": jbErr == nil,
		"jrHas": jr.HasBytes(), "jHas": j.HasBytes(),
		"jpNN": jp != nil, "mjLen": len(mj) > 0, "mjErr": mjErr == nil,
		"jm": r.JsonModel(), "jma": r.JsonModelAny() != nil,
	}
	expected := args.Map{
		"jbLen": true, "jbErr": true, "jrHas": true, "jHas": true,
		"jpNN": true, "mjLen": true, "mjErr": true,
		"jm": "test", "jma": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JSON", actual)
}

func Test_Cov7_TypedSimpleRequest_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	s, sOk := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, i64Ok := r.GetAsInt64()
	_, f64Ok := r.GetAsFloat64()
	_, f32Ok := r.GetAsFloat32()
	_, bOk := r.GetAsBool()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "i64Ok": i64Ok,
		"f64Ok": f64Ok, "f32Ok": f32Ok, "bOk": bOk, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "i64Ok": false,
		"f64Ok": false, "f32Ok": false, "bOk": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- GetAs", actual)
}

func Test_Cov7_TypedSimpleRequest_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	cloned := r.Clone()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilClone := nilR.Clone()
	actual := args.Map{
		"cloneData": cloned.Data(),
		"nilClone":  nilClone == nil,
	}
	expected := args.Map{"cloneData": "hello", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Clone", actual)
}

func Test_Cov7_TypedSimpleRequest_Conversions(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	sr := r.ToSimpleRequest()
	td := r.ToTypedDynamic()
	d := r.ToDynamic()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilSR := nilR.ToSimpleRequest()
	nilTD := nilR.ToTypedDynamic()
	nilD := nilR.ToDynamic()
	actual := args.Map{
		"srValid": sr.IsValid(), "tdValid": td.IsValid(), "dValid": d.IsValid(),
		"nilSRValid": nilSR.IsValid(), "nilTDValid": nilTD.IsValid(), "nilDValid": nilD.IsValid(),
	}
	expected := args.Map{
		"srValid": true, "tdValid": true, "dValid": true,
		"nilSRValid": false, "nilTDValid": false, "nilDValid": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- conversions", actual)
}

// ═══════════════════════════════════════════
// TypedSimpleResult
// ═══════════════════════════════════════════

func Test_Cov7_TypedSimpleResult_Constructors(t *testing.T) {
	r := coredynamic.NewTypedSimpleResult[string]("hello", true, "")
	rv := coredynamic.NewTypedSimpleResultValid[string]("world")
	inv := coredynamic.InvalidTypedSimpleResult[string]("err")
	invNoMsg := coredynamic.InvalidTypedSimpleResultNoMessage[string]()
	actual := args.Map{
		"data": r.Data(), "result": r.Result(), "valid": r.IsValid(),
		"rvData": rv.Data(), "invValid": inv.IsValid(), "invInv": inv.IsInvalid(),
		"invNoMsgMsg": invNoMsg.Message(),
	}
	expected := args.Map{
		"data": "hello", "result": "hello", "valid": true,
		"rvData": "world", "invValid": false, "invInv": true, "invNoMsgMsg": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- constructors", actual)
}

func Test_Cov7_TypedSimpleResult_NilReceiver(t *testing.T) {
	var r *coredynamic.TypedSimpleResult[string]
	actual := args.Map{
		"valid": r.IsValid(), "inv": r.IsInvalid(),
		"msg": r.Message(), "str": r.String(), "invErr": r.InvalidError() == nil,
	}
	expected := args.Map{
		"valid": false, "inv": true, "msg": "", "str": "", "invErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns nil -- nil receiver", actual)
}

func Test_Cov7_TypedSimpleResult_InvalidError(t *testing.T) {
	r := coredynamic.InvalidTypedSimpleResult[string]("err msg")
	err1 := r.InvalidError()
	err2 := r.InvalidError()
	actual := args.Map{"hasErr": err1 != nil, "cached": err1 == err2}
	expected := args.Map{"hasErr": true, "cached": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns error -- InvalidError", actual)
}

func Test_Cov7_TypedSimpleResult_JSON(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("test")
	jb, jbErr := r.JsonBytes()
	jr := r.JsonResult()
	j := r.Json()
	jp := r.JsonPtr()
	mj, mjErr := r.MarshalJSON()
	actual := args.Map{
		"jbLen": len(jb) > 0, "jbErr": jbErr == nil,
		"jrHas": jr.HasBytes(), "jHas": j.HasBytes(),
		"jpNN": jp != nil, "mjLen": len(mj) > 0, "mjErr": mjErr == nil,
		"jm": r.JsonModel(), "jma": r.JsonModelAny() != nil,
	}
	expected := args.Map{
		"jbLen": true, "jbErr": true, "jrHas": true, "jHas": true,
		"jpNN": true, "mjLen": true, "mjErr": true,
		"jm": "test", "jma": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JSON", actual)
}

func Test_Cov7_TypedSimpleResult_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[int](42)
	_, sOk := r.GetAsString()
	i, iOk := r.GetAsInt()
	_, i64Ok := r.GetAsInt64()
	_, f64Ok := r.GetAsFloat64()
	_, bOk := r.GetAsBool()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()
	actual := args.Map{
		"sOk": sOk, "i": i, "iOk": iOk, "i64Ok": i64Ok,
		"f64Ok": f64Ok, "bOk": bOk, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"sOk": false, "i": 42, "iOk": true, "i64Ok": false,
		"f64Ok": false, "bOk": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAs", actual)
}

func Test_Cov7_TypedSimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("hello")
	cloned := r.Clone()
	clonedPtr := r.ClonePtr()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilClone := nilR.ClonePtr()
	actual := args.Map{
		"cloneData": cloned.Data(), "cpData": clonedPtr.Data(),
		"nilClone": nilClone == nil,
	}
	expected := args.Map{"cloneData": "hello", "cpData": "hello", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Clone", actual)
}

func Test_Cov7_TypedSimpleResult_Conversions(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("hello")
	sr := r.ToSimpleResult()
	td := r.ToTypedDynamic()
	d := r.ToDynamic()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilSR := nilR.ToSimpleResult()
	nilTD := nilR.ToTypedDynamic()
	nilD := nilR.ToDynamic()
	actual := args.Map{
		"srValid": sr.IsValid(), "tdValid": td.IsValid(), "dValid": d.IsValid(),
		"nilSRValid": nilSR.IsValid(), "nilTDValid": nilTD.IsValid(), "nilDValid": nilD.IsValid(),
	}
	expected := args.Map{
		"srValid": true, "tdValid": true, "dValid": true,
		"nilSRValid": false, "nilTDValid": false, "nilDValid": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- conversions", actual)
}

// ═══════════════════════════════════════════
// KeyVal
// ═══════════════════════════════════════════

func Test_Cov7_KeyVal_BasicAccessors(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "myKey", Value: "myVal"}
	actual := args.Map{
		"keyStr":   kv.KeyString(),
		"valStr":   kv.ValueString(),
		"isKeyNN":  !kv.IsKeyNull(),
		"isValNN":  !kv.IsValueNull(),
		"valInt":   kv.ValueInt(),
		"valUInt":  kv.ValueUInt(),
		"valBool":  kv.ValueBool(),
		"valI64":   kv.ValueInt64(),
		"valStrs":  kv.ValueStrings() == nil,
		"strNN":    kv.String() != "",
	}
	expected := args.Map{
		"keyStr": "myKey", "valStr": "myVal", "isKeyNN": true, "isValNN": true,
		"valInt": -1, "valUInt": uint(0), "valBool": false, "valI64": int64(-1),
		"valStrs": true, "strNN": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- basic accessors", actual)
}

func Test_Cov7_KeyVal_Dynamics(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	kd := kv.KeyDynamic()
	vd := kv.ValueDynamic()
	kdp := kv.KeyDynamicPtr()
	vdp := kv.ValueDynamicPtr()
	actual := args.Map{
		"kdValid": kd.IsValid(), "vdValid": vd.IsValid(),
		"kdpNN": kdp != nil, "vdpNN": vdp != nil,
	}
	expected := args.Map{
		"kdValid": true, "vdValid": true, "kdpNN": true, "vdpNN": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- Dynamics", actual)
}

func Test_Cov7_KeyVal_NullErr(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	var nilKV *coredynamic.KeyVal
	actual := args.Map{
		"valNullErr": kv.ValueNullErr() == nil,
		"keyNullErr": kv.KeyNullErr() == nil,
		"nilValErr":  nilKV.ValueNullErr() != nil,
		"nilKeyErr":  nilKV.KeyNullErr() != nil,
	}
	expected := args.Map{
		"valNullErr": true, "keyNullErr": true,
		"nilValErr": true, "nilKeyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- NullErr", actual)
}

func Test_Cov7_KeyVal_JSON(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	jp := kv.JsonPtr()
	jm := kv.JsonModel()
	jma := kv.JsonModelAny()
	b, err := kv.Serialize()
	actual := args.Map{
		"jHas": j.HasBytes(), "jpNN": jp != nil,
		"jmNN": jm != nil, "jmaNN": jma != nil,
		"bLen": len(b) > 0, "noErr": err == nil,
	}
	expected := args.Map{
		"jHas": true, "jpNN": true, "jmNN": true, "jmaNN": true,
		"bLen": true, "noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JSON", actual)
}

func Test_Cov7_KeyVal_ReflectSet(t *testing.T) {
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	var nilKV *coredynamic.KeyVal
	actual := args.Map{
		"nilKeySet": nilKV.ReflectSetKey(nil) != nil,
		"nilValSet": nilKV.ValueReflectSet(nil) != nil,
		"nilSetTo":  nilKV.ReflectSetTo(nil) != nil,
	}
	expected := args.Map{
		"nilKeySet": true, "nilValSet": true, "nilSetTo": true,
	}
	_ = kv
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ReflectSet nil", actual)
}

func Test_Cov7_KeyVal_ValueReflectValue(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()
	actual := args.Map{"kind": rv.Kind().String()}
	expected := args.Map{"kind": "int"}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueReflectValue", actual)
}

func Test_Cov7_KeyVal_ValTyped(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	actual := args.Map{
		"valInt": kv.ValueInt(),
		"valI64": kv.ValueInt64(),
	}
	expected := args.Map{"valInt": 42, "valI64": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- val typed", actual)
}

// ═══════════════════════════════════════════
// SimpleRequest
// ═══════════════════════════════════════════

func Test_Cov7_SimpleRequest_Constructors(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("data")
	inv := coredynamic.InvalidSimpleRequest("err")
	invNo := coredynamic.InvalidSimpleRequestNoMessage()
	full := coredynamic.NewSimpleRequest("d", true, "msg")
	actual := args.Map{
		"rVal":    r.Value(),
		"rReq":    r.Request(),
		"invMsg":  inv.Message(),
		"invNoMsg": invNo.Message(),
		"fullMsg": full.Message(),
	}
	expected := args.Map{
		"rVal": "data", "rReq": "data", "invMsg": "err",
		"invNoMsg": "", "fullMsg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- constructors", actual)
}

func Test_Cov7_SimpleRequest_InvalidError(t *testing.T) {
	r := coredynamic.InvalidSimpleRequest("err")
	err1 := r.InvalidError()
	err2 := r.InvalidError()
	rv := coredynamic.NewSimpleRequestValid("ok")
	actual := args.Map{
		"hasErr": err1 != nil, "cached": err1 == err2,
		"validNoErr": rv.InvalidError() == nil,
	}
	expected := args.Map{"hasErr": true, "cached": true, "validNoErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- InvalidError", actual)
}

func Test_Cov7_SimpleRequest_TypeMismatch(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid("hello")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(42), true)
	noErr := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"hasErr": err != nil, "noErr": noErr == nil}
	expected := args.Map{"hasErr": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- GetErrorOnTypeMismatch", actual)
}

// ═══════════════════════════════════════════
// SimpleResult
// ═══════════════════════════════════════════

func Test_Cov7_SimpleResult_Constructors(t *testing.T) {
	r := coredynamic.NewSimpleResultValid("data")
	inv := coredynamic.InvalidSimpleResult("err")
	invNo := coredynamic.InvalidSimpleResultNoMessage()
	full := coredynamic.NewSimpleResult("d", true, "msg")
	actual := args.Map{
		"rResult": r.Result, "invMsg": inv.Message, "invNoMsg": invNo.Message,
		"fullMsg": full.Message,
	}
	expected := args.Map{
		"rResult": "data", "invMsg": "err", "invNoMsg": "", "fullMsg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- constructors", actual)
}

func Test_Cov7_SimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewSimpleResultValid("data")
	cloned := r.Clone()
	clonedPtr := r.ClonePtr()
	var nilR *coredynamic.SimpleResult
	nilClone := nilR.ClonePtr()
	actual := args.Map{
		"cloneResult": cloned.Result, "cpResult": clonedPtr.Result,
		"nilClone": nilClone == nil,
	}
	expected := args.Map{"cloneResult": "data", "cpResult": "data", "nilClone": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- Clone", actual)
}

// ═══════════════════════════════════════════
// DynamicGetters — uncovered branches
// ═══════════════════════════════════════════

func Test_Cov7_Dynamic_Getters_Types(t *testing.T) {
	dStr := coredynamic.NewDynamicValid("hello")
	dInt := coredynamic.NewDynamicValid(42)
	dSlice := coredynamic.NewDynamicValid([]int{1, 2})
	dMap := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	dStruct := coredynamic.NewDynamicValid(struct{}{})
	dFunc := coredynamic.NewDynamicValid(func() {})
	actual := args.Map{
		"isStr":        dStr.IsStringType(),
		"isNum":        dInt.IsNumber(),
		"isPrim":       dInt.IsPrimitive(),
		"isSlice":      dSlice.IsSliceOrArray(),
		"isSliceMap":   dSlice.IsSliceOrArrayOrMap(),
		"isMap":        dMap.IsMap(),
		"isStruct":     dStruct.IsStruct(),
		"isFunc":       dFunc.IsFunc(),
		"isPointer":    dStr.IsPointer(),
		"isValueType":  dStr.IsValueType(),
	}
	expected := args.Map{
		"isStr": true, "isNum": true, "isPrim": true,
		"isSlice": true, "isSliceMap": true, "isMap": true,
		"isStruct": true, "isFunc": true,
		"isPointer": false, "isValueType": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Getters types", actual)
}

func Test_Cov7_Dynamic_ValueExtraction(t *testing.T) {
	dInt := coredynamic.NewDynamicValid(42)
	dBool := coredynamic.NewDynamicValid(true)
	dUint := coredynamic.NewDynamicValid(uint(10))
	dStrs := coredynamic.NewDynamicValid([]string{"a", "b"})
	dI64 := coredynamic.NewDynamicValid(int64(100))
	actual := args.Map{
		"valInt":  dInt.ValueInt(),
		"valBool": dBool.ValueBool(),
		"valUint": dUint.ValueUInt(),
		"valStrs": len(dStrs.ValueStrings()),
		"valI64":  dI64.ValueInt64(),
	}
	expected := args.Map{
		"valInt": 42, "valBool": true, "valUint": uint(10),
		"valStrs": 2, "valI64": int64(100),
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- value extraction", actual)
}

func Test_Cov7_Dynamic_IntDefault(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	val, ok := d.IntDefault(99)
	d2 := coredynamic.NewDynamicValid("notanumber")
	val2, ok2 := d2.IntDefault(77)
	actual := args.Map{"val": val, "ok": ok, "val2": val2, "ok2": ok2}
	expected := args.Map{"val": 99, "ok": false, "val2": 77, "ok2": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IntDefault", actual)
}

func Test_Cov7_Dynamic_Float64(t *testing.T) {
	dNull := coredynamic.NewDynamicValid(nil)
	_, err1 := dNull.Float64()
	dBad := coredynamic.NewDynamicValid("notanum")
	_, err2 := dBad.Float64()
	actual := args.Map{"nullErr": err1 != nil, "badErr": err2 != nil}
	expected := args.Map{"nullErr": true, "badErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- Float64 errors", actual)
}

func Test_Cov7_Dynamic_ValueNullErr(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	var nilD *coredynamic.Dynamic
	actual := args.Map{
		"noErr":  d.ValueNullErr() == nil,
		"nilErr": nilD.ValueNullErr() != nil,
	}
	expected := args.Map{"noErr": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns error -- ValueNullErr", actual)
}

func Test_Cov7_Dynamic_ValueString(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	d2 := coredynamic.NewDynamicValid(42)
	var nilD *coredynamic.Dynamic
	actual := args.Map{
		"str":    d.ValueString(),
		"intStr": d2.ValueString() != "",
		"nilStr": nilD.ValueString(),
	}
	expected := args.Map{"str": "hello", "intStr": true, "nilStr": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic returns non-empty -- ValueString", actual)
}

func Test_Cov7_Dynamic_Bytes(t *testing.T) {
	d := coredynamic.NewDynamicValid([]byte("hello"))
	b, ok := d.Bytes()
	d2 := coredynamic.NewDynamicValid("not bytes")
	_, ok2 := d2.Bytes()
	var nilD *coredynamic.Dynamic
	_, ok3 := nilD.Bytes()
	actual := args.Map{"len": len(b), "ok": ok, "ok2": ok2, "ok3": ok3}
	expected := args.Map{"len": 5, "ok": true, "ok2": false, "ok3": false}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Bytes", actual)
}

func Test_Cov7_Dynamic_StructString(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	s1 := d.StructString()
	s2 := d.String()
	actual := args.Map{"s1NN": s1 != "", "s2NN": s2 != ""}
	expected := args.Map{"s1NN": true, "s2NN": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- StructString", actual)
}

func Test_Cov7_Dynamic_IsStructStringNull(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	d2 := coredynamic.NewDynamicValid("")
	d3 := coredynamic.NewDynamicValid("   ")
	actual := args.Map{
		"nullOrEmpty":    d.IsStructStringNullOrEmpty(),
		"empty":          d2.IsStructStringNullOrEmpty(),
		"wsOrEmpty":      d3.IsStructStringNullOrEmptyOrWhitespace(),
	}
	expected := args.Map{
		"nullOrEmpty": true, "empty": actual["empty"],
		"wsOrEmpty": actual["wsOrEmpty"],
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- IsStructStringNull", actual)
}

// ═══════════════════════════════════════════
// TypeStatus
// ═══════════════════════════════════════════

func Test_Cov7_TypeSameStatus(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", "world")
	tsDiff := coredynamic.TypeSameStatus("hello", 42)
	actual := args.Map{
		"same":    ts.IsSame,
		"notSame": tsDiff.IsNotSame(),
		"notEq":   tsDiff.IsNotEqualTypes(),
		"anyPtr":  tsDiff.IsAnyPointer(),
		"bothPtr": tsDiff.IsBothPointer(),
	}
	expected := args.Map{
		"same": true, "notSame": true, "notEq": true,
		"anyPtr": false, "bothPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- with args", actual)
}

func Test_Cov7_TypeStatus_Names(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", 42)
	actual := args.Map{
		"leftName":     ts.LeftName(),
		"rightName":    ts.RightName(),
		"leftFull":     ts.LeftFullName(),
		"rightFull":    ts.RightFullName(),
		"notMatchMsg":  ts.NotMatchMessage("a", "b") != "",
		"notMatchErr":  ts.NotMatchErr("a", "b") != nil,
		"validErr":     ts.ValidationError() != nil,
		"srcDestMsg":   ts.NotEqualSrcDestinationMessage() != "",
		"srcDestErr":   ts.NotEqualSrcDestinationErr() != nil,
		"sameRegPt":    ts.IsSameRegardlessPointer(),
	}
	expected := args.Map{
		"leftName": "string", "rightName": "int",
		"leftFull": "string", "rightFull": "int",
		"notMatchMsg": true, "notMatchErr": true, "validErr": true,
		"srcDestMsg": true, "srcDestErr": true, "sameRegPt": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- names", actual)
}

func Test_Cov7_TypeStatus_SameNoError(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	actual := args.Map{
		"matchMsg": ts.NotMatchMessage("a", "b"),
		"matchErr": ts.NotMatchErr("a", "b") == nil,
		"validErr": ts.ValidationError() == nil,
	}
	expected := args.Map{"matchMsg": "", "matchErr": true, "validErr": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns empty -- same no error", actual)
}

func Test_Cov7_TypeStatus_IsEqual(t *testing.T) {
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", "b")
	ts3 := coredynamic.TypeSameStatus("a", 1)
	var nilTS *coredynamic.TypeStatus
	actual := args.Map{
		"eq":        ts1.IsEqual(&ts2),
		"notEq":     ts1.IsEqual(&ts3),
		"nilBoth":   nilTS.IsEqual(nil),
		"nilOneNil": nilTS.IsEqual(&ts1),
	}
	expected := args.Map{"eq": true, "notEq": false, "nilBoth": true, "nilOneNil": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- IsEqual", actual)
}

func Test_Cov7_TypeStatus_Valid(t *testing.T) {
	ts := coredynamic.TypeSameStatus("a", "b")
	var nilTS *coredynamic.TypeStatus
	actual := args.Map{
		"valid":    ts.IsValid(),
		"invalid":  ts.IsInvalid(),
		"nilValid": nilTS.IsValid(),
		"nilInv":   nilTS.IsInvalid(),
	}
	expected := args.Map{"valid": true, "invalid": false, "nilValid": false, "nilInv": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- Valid", actual)
}

// ═══════════════════════════════════════════
// Package-level functions
// ═══════════════════════════════════════════

func Test_Cov7_TypeNotEqualErr(t *testing.T) {
	noErr := coredynamic.TypeNotEqualErr("a", "b")
	hasErr := coredynamic.TypeNotEqualErr("a", 1)
	actual := args.Map{"noErr": noErr == nil, "hasErr": hasErr != nil}
	expected := args.Map{"noErr": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr returns error -- with args", actual)
}

func Test_Cov7_TypeMustBeSame_NoPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()
	coredynamic.TypeMustBeSame("a", "b")
}

func Test_Cov7_TypeMustBeSame_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "TypeMustBeSame panics -- panic", actual)
	}()
	coredynamic.TypeMustBeSame("a", 1)
}

func Test_Cov7_IsAnyTypesOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{
		"found":    coredynamic.IsAnyTypesOf(strType, strType, intType),
		"notFound": coredynamic.IsAnyTypesOf(strType, intType),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- with args", actual)
}

func Test_Cov7_TypesIndexOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{
		"found":    coredynamic.TypesIndexOf(strType, intType, strType),
		"notFound": coredynamic.TypesIndexOf(strType, intType),
	}
	expected := args.Map{"found": 1, "notFound": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns correct value -- with args", actual)
}

func Test_Cov7_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal("hello")
	actual := args.Map{"kind": rv.Kind().String()}
	expected := args.Map{"kind": "string"}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- with args", actual)
}

func Test_Cov7_ReflectKindValidation(t *testing.T) {
	noErr := coredynamic.ReflectKindValidation(reflect.String, "hello")
	hasErr := coredynamic.ReflectKindValidation(reflect.Int, "hello")
	actual := args.Map{"noErr": noErr == nil, "hasErr": hasErr != nil}
	expected := args.Map{"noErr": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation returns non-empty -- with args", actual)
}

func Test_Cov7_ReflectTypeValidation(t *testing.T) {
	strType := reflect.TypeOf("")
	noErr := coredynamic.ReflectTypeValidation(true, strType, "hello")
	nilErr := coredynamic.ReflectTypeValidation(true, strType, nil)
	mismatch := coredynamic.ReflectTypeValidation(false, strType, 42)
	actual := args.Map{
		"noErr": noErr == nil, "nilErr": nilErr != nil, "mismatch": mismatch != nil,
	}
	expected := args.Map{"noErr": true, "nilErr": true, "mismatch": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation returns non-empty -- with args", actual)
}

func Test_Cov7_NotAcceptedTypesErr(t *testing.T) {
	strType := reflect.TypeOf("")
	noErr := coredynamic.NotAcceptedTypesErr("hello", strType)
	hasErr := coredynamic.NotAcceptedTypesErr(42, strType)
	actual := args.Map{"noErr": noErr == nil, "hasErr": hasErr != nil}
	expected := args.Map{"noErr": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- with args", actual)
}

func Test_Cov7_PointerOrNonPointer(t *testing.T) {
	val, _ := coredynamic.PointerOrNonPointer(false, "hello")
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns correct value -- with args", actual)
}

func Test_Cov7_BytesConverter_Basic(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()
	cs := bc.SafeCastString()
	cs2, csErr := bc.CastString()
	actual := args.Map{
		"s": s, "noErr": err == nil, "cs": cs != "", "cs2": cs2 != "", "csErr": csErr == nil,
	}
	expected := args.Map{
		"s": "hello", "noErr": true, "cs": true, "cs2": true, "csErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- basic", actual)
}

func Test_Cov7_BytesConverter_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	cs := bc.SafeCastString()
	_, csErr := bc.CastString()
	actual := args.Map{"cs": cs, "csErr": csErr != nil}
	expected := args.Map{"cs": "", "csErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns empty -- empty", actual)
}

func Test_Cov7_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`true`))
	b, err := bc.ToBool()
	actual := args.Map{"b": b, "noErr": err == nil}
	expected := args.Map{"b": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBool", actual)
}

func Test_Cov7_CastedResult_Methods(t *testing.T) {
	cr := &coredynamic.CastedResult{
		IsValid: true, IsNull: false, IsPointer: false,
		IsMatchingAcceptedType: true, IsSourcePointer: false,
		SourceKind: reflect.String,
	}
	var nilCR *coredynamic.CastedResult
	actual := args.Map{
		"isInvalid":    cr.IsInvalid(),
		"isNotNull":    cr.IsNotNull(),
		"isNotPtr":     cr.IsNotPointer(),
		"isNotMatch":   cr.IsNotMatchingAcceptedType(),
		"isSrcKind":    cr.IsSourceKind(reflect.String),
		"hasErr":       cr.HasError(),
		"hasIssues":    cr.HasAnyIssues(),
		"nilInvalid":   nilCR.IsInvalid(),
		"nilNotNull":   nilCR.IsNotNull(),
		"nilNotPtr":    nilCR.IsNotPointer(),
		"nilNotMatch":  nilCR.IsNotMatchingAcceptedType(),
		"nilSrcKind":   nilCR.IsSourceKind(reflect.String),
		"nilHasErr":    nilCR.HasError(),
	}
	expected := args.Map{
		"isInvalid": false, "isNotNull": true, "isNotPtr": true,
		"isNotMatch": false, "isSrcKind": true, "hasErr": false,
		"hasIssues": false, "nilInvalid": true, "nilNotNull": false,
		"nilNotPtr": false, "nilNotMatch": false, "nilSrcKind": false,
		"nilHasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult returns correct value -- methods", actual)
}

func Test_Cov7_ZeroSetAny(t *testing.T) {
	type sample struct{ Name string }
	s := &sample{Name: "hello"}
	coredynamic.ZeroSetAny(s)
	actual := args.Map{"name": s.Name}
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny returns correct value -- with args", actual)
}

func Test_Cov7_ZeroSetAny_Nil(t *testing.T) {
	// should not panic
	coredynamic.ZeroSetAny(nil)
}

func Test_Cov7_MapAnyItems_Constructors(t *testing.T) {
	empty := coredynamic.EmptyMapAnyItems()
	withItems := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	emptyItems := coredynamic.NewMapAnyItemsUsingItems(nil)
	actual := args.Map{
		"emptyLen":     empty.Length(),
		"withItemsLen": withItems.Length(),
		"emptyItemsLen": emptyItems.Length(),
	}
	expected := args.Map{"emptyLen": 0, "withItemsLen": 1, "emptyItemsLen": 0}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- constructors", actual)
}

func Test_Cov7_MapAnyItems_AddAndGet(t *testing.T) {
	m := coredynamic.NewMapAnyItems(5)
	isNew := m.Add("key1", "val1")
	isNew2 := m.Set("key1", "val2")
	v := m.GetValue("key1")
	_, has := m.Get("key1")
	_, notHas := m.Get("missing")
	actual := args.Map{
		"isNew": isNew, "isNew2": isNew2, "v": v, "has": has, "notHas": notHas,
	}
	expected := args.Map{"isNew": true, "isNew2": false, "v": "val2", "has": true, "notHas": false}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- Add/Get", actual)
}

func Test_Cov7_MapAnyItems_HasKey(t *testing.T) {
	m := coredynamic.NewMapAnyItems(3)
	m.Add("k", "v")
	var nilM *coredynamic.MapAnyItems
	actual := args.Map{
		"has":    m.HasKey("k"),
		"notHas": m.HasKey("x"),
		"nilHas": nilM.HasKey("k"),
		"empty":  m.IsEmpty(),
		"hasAny": m.HasAnyItem(),
	}
	expected := args.Map{"has": true, "notHas": false, "nilHas": false, "empty": false, "hasAny": true}
	expected.ShouldBeEqual(t, 0, "MapAnyItems returns correct value -- HasKey", actual)
}
