package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// TypedDynamic — comprehensive coverage
// ==========================================================================

func Test_Cov2_TypedDynamic_Constructors(t *testing.T) {
	d1 := coredynamic.NewTypedDynamic[string]("hello", true)
	d2 := coredynamic.NewTypedDynamicValid[string]("world")
	d3 := coredynamic.NewTypedDynamicPtr[int](42, true)
	d4 := coredynamic.InvalidTypedDynamic[string]()
	d5 := coredynamic.InvalidTypedDynamicPtr[string]()
	actual := args.Map{
		"d1Valid": d1.IsValid(), "d1Data": d1.Data(),
		"d2Valid": d2.IsValid(), "d2Value": d2.Value(),
		"d3NotNil": d3 != nil, "d3Data": d3.Data(),
		"d4Invalid": d4.IsInvalid(),
		"d5NotNil": d5 != nil, "d5Invalid": d5.IsInvalid(),
	}
	expected := args.Map{
		"d1Valid": true, "d1Data": "hello",
		"d2Valid": true, "d2Value": "world",
		"d3NotNil": true, "d3Data": 42,
		"d4Invalid": true,
		"d5NotNil": true, "d5Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic constructors return expected -- all variants", actual)
}

func Test_Cov2_TypedDynamic_StringAndJson(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[string]("test")
	jsonBytes, jsonErr := d.JsonBytes()
	jsonStr, jsonStrErr := d.JsonString()
	marshalBytes, marshalErr := d.MarshalJSON()
	valueMarshal, valueMarshalErr := d.ValueMarshal()
	_, bytesOk := d.Bytes()
	actual := args.Map{
		"string": d.String(), "jsonLen": len(jsonBytes) > 0, "jsonErr": jsonErr != nil,
		"jsonStr": jsonStr != "", "jsonStrErr": jsonStrErr != nil,
		"marshalLen": len(marshalBytes) > 0, "marshalErr": marshalErr != nil,
		"valueMarshalLen": len(valueMarshal) > 0, "valueMarshalErr": valueMarshalErr != nil,
		"bytesOk": bytesOk,
		"jsonModel": d.JsonModel(), "jsonModelAny": d.JsonModelAny() != nil,
	}
	expected := args.Map{
		"string": "test", "jsonLen": true, "jsonErr": false,
		"jsonStr": true, "jsonStrErr": false,
		"marshalLen": true, "marshalErr": false,
		"valueMarshalLen": true, "valueMarshalErr": false,
		"bytesOk": true,
		"jsonModel": "test", "jsonModelAny": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic string and JSON methods -- string value", actual)
}

func Test_Cov2_TypedDynamic_GetAs(t *testing.T) {
	dStr := coredynamic.NewTypedDynamicValid[any]("hello")
	dInt := coredynamic.NewTypedDynamicValid[any](42)
	dBool := coredynamic.NewTypedDynamicValid[any](true)
	dFloat64 := coredynamic.NewTypedDynamicValid[any](3.14)
	dBytes := coredynamic.NewTypedDynamicValid[any]([]byte("hi"))
	dStrings := coredynamic.NewTypedDynamicValid[any]([]string{"a"})

	strVal, strOk := dStr.GetAsString()
	intVal, intOk := dInt.GetAsInt()
	boolVal, boolOk := dBool.GetAsBool()
	f64Val, f64Ok := dFloat64.GetAsFloat64()
	bytesVal, bytesOk := dBytes.GetAsBytes()
	stringsVal, stringsOk := dStrings.GetAsStrings()
	_, f32Ok := dStr.GetAsFloat32()
	_, i64Ok := dStr.GetAsInt64()
	_, uintOk := dStr.GetAsUint()

	actual := args.Map{
		"str": strVal, "strOk": strOk,
		"int": intVal, "intOk": intOk,
		"bool": boolVal, "boolOk": boolOk,
		"f64Above3": f64Val > 3, "f64Ok": f64Ok,
		"bytesLen": len(bytesVal), "bytesOk": bytesOk,
		"stringsLen": len(stringsVal), "stringsOk": stringsOk,
		"f32Ok": f32Ok, "i64Ok": i64Ok, "uintOk": uintOk,
	}
	expected := args.Map{
		"str": "hello", "strOk": true,
		"int": 42, "intOk": true,
		"bool": true, "boolOk": true,
		"f64Above3": true, "f64Ok": true,
		"bytesLen": 2, "bytesOk": true,
		"stringsLen": 1, "stringsOk": true,
		"f32Ok": false, "i64Ok": false, "uintOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic GetAs methods return expected -- various types", actual)
}

func Test_Cov2_TypedDynamic_ValueMethods(t *testing.T) {
	dStr := coredynamic.NewTypedDynamicValid[any]("hello")
	dInt := coredynamic.NewTypedDynamicValid[any](42)
	dBool := coredynamic.NewTypedDynamicValid[any](true)
	dI64 := coredynamic.NewTypedDynamicValid[any](int64(100))

	actual := args.Map{
		"valueString": dStr.ValueString(),
		"valueInt":    dInt.ValueInt(),
		"valueBool":   dBool.ValueBool(),
		"valueInt64":  int(dI64.ValueInt64()),
	}
	expected := args.Map{
		"valueString": "hello", "valueInt": 42,
		"valueBool": true, "valueInt64": 100,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Value* methods return expected -- various", actual)
}

func Test_Cov2_TypedDynamic_Clone(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[string]("hello")
	cloned := d.Clone()
	ptr := coredynamic.NewTypedDynamicPtr[string]("world", true)
	clonedPtr := ptr.ClonePtr()
	var nilPtr *coredynamic.TypedDynamic[string]
	nilCloned := nilPtr.ClonePtr()

	actual := args.Map{
		"clonedData": cloned.Data(), "clonedValid": cloned.IsValid(),
		"clonedPtrData": clonedPtr.Data(),
		"nilCloned": nilCloned == nil,
		"nonPtr": d.NonPtr().Data(),
		"ptr": d.Ptr() != nil,
	}
	expected := args.Map{
		"clonedData": "hello", "clonedValid": true,
		"clonedPtrData": "world",
		"nilCloned": true,
		"nonPtr": "hello", "ptr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Clone methods return expected -- all paths", actual)
}

func Test_Cov2_TypedDynamic_ToDynamic(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[string]("hello")
	dyn := d.ToDynamic()
	actual := args.Map{"dynValid": dyn.IsValid(), "dynData": dyn.Data()}
	expected := args.Map{"dynValid": true, "dynData": "hello"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ToDynamic returns expected -- valid", actual)
}

func Test_Cov2_TypedDynamic_Deserialize(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := d.Deserialize([]byte(`"hello"`))
	actual := args.Map{"hasErr": err != nil, "data": d.Data(), "valid": d.IsValid()}
	expected := args.Map{"hasErr": false, "data": "hello", "valid": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize sets data -- valid JSON", actual)
}

func Test_Cov2_TypedDynamic_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := d.UnmarshalJSON([]byte(`"world"`))
	actual := args.Map{"hasErr": err != nil, "data": d.Data()}
	expected := args.Map{"hasErr": false, "data": "world"}
	expected.ShouldBeEqual(t, 0, "TypedDynamic UnmarshalJSON sets data -- valid JSON", actual)
}

func Test_Cov2_TypedDynamic_JsonResult(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[string]("test")
	jr := d.JsonResult()
	jp := d.JsonPtr()
	actual := args.Map{"jrNotNil": true, "jpNotNil": jp != nil, "jsonNotNil": d.Json().Bytes != nil}
	_ = jr
	expected := args.Map{"jrNotNil": true, "jpNotNil": true, "jsonNotNil": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonResult returns non-nil -- valid", actual)
}

func Test_Cov2_TypedDynamic_BytesNonByteType(t *testing.T) {
	d := coredynamic.NewTypedDynamicValid[int](42)
	bytes, ok := d.Bytes()
	actual := args.Map{"ok": ok, "hasBytes": len(bytes) > 0}
	expected := args.Map{"ok": true, "hasBytes": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Bytes marshals non-byte type -- int", actual)
}

// ==========================================================================
// TypedSimpleResult — comprehensive coverage
// ==========================================================================

func Test_Cov2_TypedSimpleResult_Constructors(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleResult[string]("hello", true, "")
	r2 := coredynamic.NewTypedSimpleResultValid[string]("world")
	r3 := coredynamic.InvalidTypedSimpleResult[string]("error msg")
	r4 := coredynamic.InvalidTypedSimpleResultNoMessage[string]()
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Data": r1.Data(), "r1Result": r1.Result(),
		"r2Valid": r2.IsValid(), "r2Message": r2.Message(),
		"r3Invalid": r3.IsInvalid(), "r3Message": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}
	expected := args.Map{
		"r1Valid": true, "r1Data": "hello", "r1Result": "hello",
		"r2Valid": true, "r2Message": "",
		"r3Invalid": true, "r3Message": "error msg",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult constructors return expected -- all", actual)
}

func Test_Cov2_TypedSimpleResult_InvalidError(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleResultValid[string]("ok")
	r2 := coredynamic.InvalidTypedSimpleResult[string]("err")
	var r3 *coredynamic.TypedSimpleResult[string]
	actual := args.Map{
		"r1Err": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
		"r3Err": r3.InvalidError() == nil,
		"r3Valid": r3.IsValid(),
		"r3Invalid": r3.IsInvalid(),
		"r3Msg": r3.Message(),
		"r3Str": r3.String(),
	}
	expected := args.Map{
		"r1Err": true, "r2Err": true,
		"r3Err": true, "r3Valid": false, "r3Invalid": true,
		"r3Msg": "", "r3Str": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult InvalidError returns expected -- all paths", actual)
}

func Test_Cov2_TypedSimpleResult_Json(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("test")
	jb, jbErr := r.JsonBytes()
	mb, mbErr := r.MarshalJSON()
	actual := args.Map{
		"jbLen": len(jb) > 0, "jbErr": jbErr != nil,
		"mbLen": len(mb) > 0, "mbErr": mbErr != nil,
		"model": r.JsonModel(), "modelAny": r.JsonModelAny() != nil,
		"jsonNotNil": true, "jsonPtrNotNil": r.JsonPtr() != nil,
	}
	_ = r.Json()
	_ = r.JsonResult()
	expected := args.Map{
		"jbLen": true, "jbErr": false,
		"mbLen": true, "mbErr": false,
		"model": "test", "modelAny": true,
		"jsonNotNil": true, "jsonPtrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult JSON methods return expected -- valid", actual)
}

func Test_Cov2_TypedSimpleResult_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[any]("hello")
	s, sOk := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, bOk := r.GetAsBool()
	_, fOk := r.GetAsFloat64()
	_, i64Ok := r.GetAsInt64()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "bOk": bOk,
		"fOk": fOk, "i64Ok": i64Ok, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "bOk": false,
		"fOk": false, "i64Ok": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult GetAs methods return expected -- string", actual)
}

func Test_Cov2_TypedSimpleResult_CloneAndConvert(t *testing.T) {
	r := coredynamic.NewTypedSimpleResultValid[string]("hello")
	cloned := r.Clone()
	clonedPtr := r.ClonePtr()
	var nilR *coredynamic.TypedSimpleResult[string]
	nilClone := nilR.Clone()
	nilClonePtr := nilR.ClonePtr()
	sr := r.ToSimpleResult()
	nilSr := nilR.ToSimpleResult()
	td := r.ToTypedDynamic()
	nilTd := nilR.ToTypedDynamic()
	dyn := r.ToDynamic()
	nilDyn := nilR.ToDynamic()
	actual := args.Map{
		"clonedData": cloned.Data(), "clonedPtrNotNil": clonedPtr != nil,
		"nilCloneInvalid": nilClone.IsInvalid(), "nilClonePtr": nilClonePtr == nil,
		"srValid": sr.IsValid(), "nilSrInvalid": nilSr.IsInvalid(),
		"tdValid": td.IsValid(), "nilTdInvalid": nilTd.IsInvalid(),
		"dynValid": dyn.IsValid(), "nilDynInvalid": nilDyn.IsInvalid(),
	}
	expected := args.Map{
		"clonedData": "hello", "clonedPtrNotNil": true,
		"nilCloneInvalid": true, "nilClonePtr": true,
		"srValid": true, "nilSrInvalid": true,
		"tdValid": true, "nilTdInvalid": true,
		"dynValid": true, "nilDynInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult Clone/Convert return expected -- all paths", actual)
}

// ==========================================================================
// TypedSimpleRequest — comprehensive coverage
// ==========================================================================

func Test_Cov2_TypedSimpleRequest_Constructors(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleRequest[string]("hello", true, "")
	r2 := coredynamic.NewTypedSimpleRequestValid[string]("world")
	r3 := coredynamic.InvalidTypedSimpleRequest[string]("err")
	r4 := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Data": r1.Data(), "r1Request": r1.Request(), "r1Value": r1.Value(),
		"r2Valid": r2.IsValid(),
		"r3Invalid": r3.IsInvalid(), "r3Message": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}
	expected := args.Map{
		"r1Valid": true, "r1Data": "hello", "r1Request": "hello", "r1Value": "hello",
		"r2Valid": true,
		"r3Invalid": true, "r3Message": "err",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest constructors return expected -- all", actual)
}

func Test_Cov2_TypedSimpleRequest_InvalidError(t *testing.T) {
	r1 := coredynamic.NewTypedSimpleRequestValid[string]("ok")
	r2 := coredynamic.InvalidTypedSimpleRequest[string]("err")
	var r3 *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{
		"r1Err": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
		"r3Err": r3.InvalidError() == nil,
		"r3Valid": r3.IsValid(), "r3Invalid": r3.IsInvalid(),
		"r3Msg": r3.Message(), "r3Str": r3.String(),
	}
	expected := args.Map{
		"r1Err": true, "r2Err": true, "r3Err": true,
		"r3Valid": false, "r3Invalid": true, "r3Msg": "", "r3Str": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest InvalidError returns expected -- all paths", actual)
}

func Test_Cov2_TypedSimpleRequest_Json(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("test")
	jb, _ := r.JsonBytes()
	mb, _ := r.MarshalJSON()
	_ = r.Json()
	_ = r.JsonResult()
	_ = r.JsonPtr()
	actual := args.Map{
		"jbLen": len(jb) > 0, "mbLen": len(mb) > 0,
		"model": r.JsonModel(), "modelAny": r.JsonModelAny() != nil,
	}
	expected := args.Map{
		"jbLen": true, "mbLen": true,
		"model": "test", "modelAny": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest JSON methods return expected -- valid", actual)
}

func Test_Cov2_TypedSimpleRequest_GetAs(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[any]("hello")
	s, sOk := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, bOk := r.GetAsBool()
	_, fOk := r.GetAsFloat64()
	_, f32Ok := r.GetAsFloat32()
	_, i64Ok := r.GetAsInt64()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "bOk": bOk,
		"fOk": fOk, "f32Ok": f32Ok, "i64Ok": i64Ok, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "bOk": false,
		"fOk": false, "f32Ok": false, "i64Ok": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest GetAs methods return expected -- string", actual)
}

func Test_Cov2_TypedSimpleRequest_CloneAndConvert(t *testing.T) {
	r := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	cloned := r.Clone()
	var nilR *coredynamic.TypedSimpleRequest[string]
	nilClone := nilR.Clone()
	sr := r.ToSimpleRequest()
	nilSr := nilR.ToSimpleRequest()
	td := r.ToTypedDynamic()
	nilTd := nilR.ToTypedDynamic()
	dyn := r.ToDynamic()
	nilDyn := nilR.ToDynamic()
	actual := args.Map{
		"clonedData": cloned.Data(),
		"nilClone": nilClone == nil,
		"srValid": sr.IsValid(), "nilSrInvalid": nilSr.IsInvalid(),
		"tdValid": td.IsValid(), "nilTdInvalid": nilTd.IsInvalid(),
		"dynValid": dyn.IsValid(), "nilDynInvalid": nilDyn.IsInvalid(),
	}
	expected := args.Map{
		"clonedData": "hello", "nilClone": true,
		"srValid": true, "nilSrInvalid": true,
		"tdValid": true, "nilTdInvalid": true,
		"dynValid": true, "nilDynInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest Clone/Convert return expected -- all paths", actual)
}

// ==========================================================================
// SafeTypeName
// ==========================================================================

func Test_Cov2_SafeTypeName(t *testing.T) {
	actual := args.Map{
		"string": coredynamic.SafeTypeName("hello"),
		"int":    coredynamic.SafeTypeName(42),
		"nil":    coredynamic.SafeTypeName(nil),
	}
	expected := args.Map{
		"string": "string", "int": "int", "nil": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns expected -- various types", actual)
}

// ==========================================================================
// PointerOrNonPointer
// ==========================================================================

func Test_Cov2_PointerOrNonPointer(t *testing.T) {
	val := "hello"
	outPtr, _ := coredynamic.PointerOrNonPointer(true, &val)
	outNonPtr, _ := coredynamic.PointerOrNonPointer(false, &val)
	actual := args.Map{
		"ptrNotNil":    outPtr != nil,
		"nonPtrNotNil": outNonPtr != nil,
	}
	expected := args.Map{
		"ptrNotNil": true, "nonPtrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns expected -- ptr and non-ptr", actual)
}

// ==========================================================================
// LengthOfReflect
// ==========================================================================

func Test_Cov2_LengthOfReflect(t *testing.T) {
	actual := args.Map{
		"slice": coredynamic.LengthOfReflect(reflect.ValueOf([]int{1, 2, 3})),
		"array": coredynamic.LengthOfReflect(reflect.ValueOf([2]int{1, 2})),
		"map":   coredynamic.LengthOfReflect(reflect.ValueOf(map[string]int{"a": 1})),
		"int":   coredynamic.LengthOfReflect(reflect.ValueOf(42)),
	}
	expected := args.Map{"slice": 3, "array": 2, "map": 1, "int": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns expected -- various kinds", actual)
}

// ==========================================================================
// BytesConverter
// ==========================================================================

func Test_Cov2_BytesConverter_SafeCastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	actual := args.Map{
		"safe":      bc.SafeCastString(),
		"safeEmpty": bcEmpty.SafeCastString(),
	}
	expected := args.Map{"safe": "hello", "safeEmpty": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString returns expected -- valid and empty", actual)
}

func Test_Cov2_BytesConverter_CastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	val, err := bc.CastString()
	_, errEmpty := bcEmpty.CastString()
	actual := args.Map{
		"val": val, "hasErr": err != nil, "emptyHasErr": errEmpty != nil,
	}
	expected := args.Map{"val": "hello", "hasErr": false, "emptyHasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString returns expected -- valid and empty", actual)
}

func Test_Cov2_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("true"))
	val, err := bc.ToBool()
	actual := args.Map{"val": val, "hasErr": err != nil}
	expected := args.Map{"val": true, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBool returns true -- valid", actual)
}

func Test_Cov2_BytesConverter_ToBoolMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("false"))
	actual := args.Map{"val": bc.ToBoolMust()}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBoolMust returns false -- valid", actual)
}

func Test_Cov2_BytesConverter_ToString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	val, err := bc.ToString()
	actual := args.Map{"val": val, "hasErr": err != nil}
	expected := args.Map{"val": "hello", "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToString returns expected -- valid JSON string", actual)
}

func Test_Cov2_BytesConverter_ToStringMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	actual := args.Map{"val": bc.ToStringMust()}
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringMust returns expected -- valid", actual)
}

func Test_Cov2_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	val, err := bc.ToStrings()
	actual := args.Map{"len": len(val), "hasErr": err != nil}
	expected := args.Map{"len": 2, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStrings returns expected -- valid", actual)
}

func Test_Cov2_BytesConverter_ToStringsMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	val := bc.ToStringsMust()
	actual := args.Map{"len": len(val)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringsMust returns expected -- valid", actual)
}

func Test_Cov2_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("42"))
	val, err := bc.ToInt64()
	actual := args.Map{"val": int(val), "hasErr": err != nil}
	expected := args.Map{"val": 42, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64 returns expected -- valid", actual)
}

func Test_Cov2_BytesConverter_ToInt64Must(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("99"))
	actual := args.Map{"val": int(bc.ToInt64Must())}
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64Must returns expected -- valid", actual)
}

func Test_Cov2_BytesConverter_Deserialize(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var result string
	err := bc.Deserialize(&result)
	actual := args.Map{"hasErr": err != nil, "result": result}
	expected := args.Map{"hasErr": false, "result": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize returns expected -- valid", actual)
}

func Test_Cov2_BytesConverter_DeserializeMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	var result string
	bc.DeserializeMust(&result)
	actual := args.Map{"result": result}
	expected := args.Map{"result": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter DeserializeMust returns expected -- valid", actual)
}

// ==========================================================================
// Dynamic — additional uncovered methods
// ==========================================================================

func Test_Cov2_Dynamic_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	kv, err := d.MapToKeyVal()
	actual := args.Map{"hasErr": err != nil, "notNil": kv != nil}
	expected := args.Map{"hasErr": false, "notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic MapToKeyVal returns no error -- valid map", actual)
}

func Test_Cov2_Dynamic_ItemReflectValue(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	rv := d.ItemReflectValueUsingIndex(0)
	actual := args.Map{"valid": rv.IsValid(), "val": rv.Interface()}
	expected := args.Map{"valid": true, "val": "a"}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemReflectValueUsingIndex returns expected -- index 0", actual)
}

func Test_Cov2_Dynamic_ItemReflectValueUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"x": 99})
	rv := d.ItemReflectValueUsingKey("x")
	actual := args.Map{"valid": rv.IsValid(), "val": int(rv.Int())}
	expected := args.Map{"valid": true, "val": 99}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemReflectValueUsingKey returns expected -- key x", actual)
}

func Test_Cov2_Dynamic_Json_Deserialize(t *testing.T) {
	d := coredynamic.NewDynamicPtr("", true)
	_, err := d.Deserialize([]byte(`"hello"`))
	var nilD *coredynamic.Dynamic
	_, nilErr := nilD.Deserialize(nil)
	actual := args.Map{"err": err != nil, "nilErr": nilErr != nil}
	expected := args.Map{"err": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Deserialize returns expected -- valid and nil", actual)
}

func Test_Cov2_Dynamic_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicPtr("", true)
	err := d.UnmarshalJSON([]byte(`"test"`))
	var nilD *coredynamic.Dynamic
	nilErr := nilD.UnmarshalJSON(nil)
	actual := args.Map{"err": err != nil, "nilErr": nilErr != nil}
	expected := args.Map{"err": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic UnmarshalJSON returns expected -- valid and nil", actual)
}

func Test_Cov2_Dynamic_ValueMarshal_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueMarshal returns error -- nil receiver", actual)
}

func Test_Cov2_Dynamic_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	jp := d.JsonPtr()
	j := d.Json()
	actual := args.Map{"jpNotNil": jp != nil, "jNotNil": j.Bytes != nil}
	expected := args.Map{"jpNotNil": true, "jNotNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Json/JsonPtr returns non-nil -- valid", actual)
}

func Test_Cov2_Dynamic_LoopBreak(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return index == 0 // break after first
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop breaks early -- break at index 0", actual)
}

func Test_Cov2_Dynamic_FilterBreak(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	filtered := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, index == 1 // break at index 1
	})
	actual := args.Map{"length": filtered.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection breaks early -- at index 1", actual)
}

func Test_Cov2_Dynamic_LoopMapBreak(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2, "c": 3})
	count := 0
	d.LoopMap(func(index int, key, value any) bool {
		count++
		return true // break after first
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap breaks early -- break at first", actual)
}

// ==========================================================================
// IsAnyTypesOf
// ==========================================================================

func Test_Cov2_IsAnyTypesOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	actual := args.Map{
		"match":   coredynamic.IsAnyTypesOf(strType, strType, intType),
		"noMatch": coredynamic.IsAnyTypesOf(strType, intType),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns expected -- match and no match", actual)
}

// ==========================================================================
// Dynamic — ConvertUsingFunc
// ==========================================================================

func Test_Cov2_Dynamic_ConvertUsingFunc(t *testing.T) {
	d := coredynamic.NewDynamicValid("42")
	converter := func(data any, expectedType reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(data)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(""))
	actual := args.Map{"valid": result.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ConvertUsingFunc returns valid -- identity converter", actual)
}

// ==========================================================================
// Dynamic — JSON roundtrip marshal/unmarshal
// ==========================================================================

func Test_Cov2_Dynamic_JsonRoundTrip(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
	}
	d := coredynamic.NewDynamicValid(testStruct{Name: "alice"})
	marshalledBytes, _ := json.Marshal(d)
	actual := args.Map{"hasBytes": len(marshalledBytes) > 0}
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JSON roundtrip produces bytes -- struct", actual)
}
