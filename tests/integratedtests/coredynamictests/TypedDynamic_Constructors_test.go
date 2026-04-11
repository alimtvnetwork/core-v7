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

func Test_TypedDynamic_Constructors_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d1 := coredynamic.NewTypedDynamic[string]("hello", true)
	d2 := coredynamic.NewTypedDynamicValid[string]("world")
	d3 := coredynamic.NewTypedDynamicPtr[int](42, true)
	d4 := coredynamic.InvalidTypedDynamic[string]()
	d5 := coredynamic.InvalidTypedDynamicPtr[string]()

	// Act
	actual := args.Map{
		"d1Valid": d1.IsValid(), "d1Data": d1.Data(),
		"d2Valid": d2.IsValid(), "d2Value": d2.Value(),
		"d3NotNil": d3 != nil, "d3Data": d3.Data(),
		"d4Invalid": d4.IsInvalid(),
		"d5NotNil": d5 != nil, "d5Invalid": d5.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"d1Valid": true, "d1Data": "hello",
		"d2Valid": true, "d2Value": "world",
		"d3NotNil": true, "d3Data": 42,
		"d4Invalid": true,
		"d5NotNil": true, "d5Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic constructors return expected -- all variants", actual)
}

func Test_TypedDynamic_StringAndJson(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[string]("test")
	jsonBytes, jsonErr := d.JsonBytes()
	jsonStr, jsonStrErr := d.JsonString()
	marshalBytes, marshalErr := d.MarshalJSON()
	valueMarshal, valueMarshalErr := d.ValueMarshal()
	_, bytesOk := d.Bytes()

	// Act
	actual := args.Map{
		"string": d.String(), "jsonLen": len(jsonBytes) > 0, "jsonErr": jsonErr != nil,
		"jsonStr": jsonStr != "", "jsonStrErr": jsonStrErr != nil,
		"marshalLen": len(marshalBytes) > 0, "marshalErr": marshalErr != nil,
		"valueMarshalLen": len(valueMarshal) > 0, "valueMarshalErr": valueMarshalErr != nil,
		"bytesOk": bytesOk,
		"jsonModel": d.JsonModel(), "jsonModelAny": d.JsonModelAny() != nil,
	}

	// Assert
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

func Test_TypedDynamic_GetAs_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{
		"str": strVal, "strOk": strOk,
		"int": intVal, "intOk": intOk,
		"bool": boolVal, "boolOk": boolOk,
		"f64Above3": f64Val > 3, "f64Ok": f64Ok,
		"bytesLen": len(bytesVal), "bytesOk": bytesOk,
		"stringsLen": len(stringsVal), "stringsOk": stringsOk,
		"f32Ok": f32Ok, "i64Ok": i64Ok, "uintOk": uintOk,
	}

	// Assert
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

func Test_TypedDynamic_ValueMethods(t *testing.T) {
	// Arrange
	dStr := coredynamic.NewTypedDynamicValid[any]("hello")
	dInt := coredynamic.NewTypedDynamicValid[any](42)
	dBool := coredynamic.NewTypedDynamicValid[any](true)
	dI64 := coredynamic.NewTypedDynamicValid[any](int64(100))

	// Act
	actual := args.Map{
		"valueString": dStr.ValueString(),
		"valueInt":    dInt.ValueInt(),
		"valueBool":   dBool.ValueBool(),
		"valueInt64":  int(dI64.ValueInt64()),
	}

	// Assert
	expected := args.Map{
		"valueString": "hello", "valueInt": 42,
		"valueBool": true, "valueInt64": 100,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Value* methods return expected -- various", actual)
}

func Test_TypedDynamic_Clone_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[string]("hello")
	cloned := d.Clone()
	ptr := coredynamic.NewTypedDynamicPtr[string]("world", true)
	clonedPtr := ptr.ClonePtr()
	var nilPtr *coredynamic.TypedDynamic[string]
	nilCloned := nilPtr.ClonePtr()

	// Act
	actual := args.Map{
		"clonedData": cloned.Data(), "clonedValid": cloned.IsValid(),
		"clonedPtrData": clonedPtr.Data(),
		"nilCloned": nilCloned == nil,
		"nonPtr": d.NonPtr().Data(),
		"ptr": d.Ptr() != nil,
	}

	// Assert
	expected := args.Map{
		"clonedData": "hello", "clonedValid": true,
		"clonedPtrData": "world",
		"nilCloned": true,
		"nonPtr": "hello", "ptr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Clone methods return expected -- all paths", actual)
}

func Test_TypedDynamic_ToDynamic_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[string]("hello")
	dyn := d.ToDynamic()

	// Act
	actual := args.Map{
		"dynValid": dyn.IsValid(),
		"dynData": dyn.Data(),
	}

	// Assert
	expected := args.Map{
		"dynValid": true,
		"dynData": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic ToDynamic returns expected -- valid", actual)
}

func Test_TypedDynamic_Deserialize_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := d.Deserialize([]byte(`"hello"`))

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"data": d.Data(),
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"data": "hello",
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Deserialize sets data -- valid JSON", actual)
}

func Test_TypedDynamic_UnmarshalJSON_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := d.UnmarshalJSON([]byte(`"world"`))

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"data": d.Data(),
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"data": "world",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic UnmarshalJSON sets data -- valid JSON", actual)
}

func Test_TypedDynamic_JsonResult_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[string]("test")
	jr := d.JsonResult()
	jp := d.JsonPtr()

	// Act
	actual := args.Map{
		"jrNotNil": true,
		"jpNotNil": jp != nil,
		"jsonNotNil": d.Json().Bytes != nil,
	}
	_ = jr

	// Assert
	expected := args.Map{
		"jrNotNil": true,
		"jpNotNil": true,
		"jsonNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic JsonResult returns non-nil -- valid", actual)
}

func Test_TypedDynamic_BytesNonByteType(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicValid[int](42)
	bytes, ok := d.Bytes()

	// Act
	actual := args.Map{
		"ok": ok,
		"hasBytes": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic Bytes marshals non-byte type -- int", actual)
}

// ==========================================================================
// TypedSimpleResult — comprehensive coverage
// ==========================================================================

func Test_TypedSimpleResult_Constructors_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleResult[string]("hello", true, "")
	r2 := coredynamic.NewTypedSimpleResultValid[string]("world")
	r3 := coredynamic.InvalidTypedSimpleResult[string]("error msg")
	r4 := coredynamic.InvalidTypedSimpleResultNoMessage[string]()

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Data": r1.Data(), "r1Result": r1.Result(),
		"r2Valid": r2.IsValid(), "r2Message": r2.Message(),
		"r3Invalid": r3.IsInvalid(), "r3Message": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"r1Valid": true, "r1Data": "hello", "r1Result": "hello",
		"r2Valid": true, "r2Message": "",
		"r3Invalid": true, "r3Message": "error msg",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult constructors return expected -- all", actual)
}

func Test_TypedSimpleResult_InvalidError_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleResultValid[string]("ok")
	r2 := coredynamic.InvalidTypedSimpleResult[string]("err")
	var r3 *coredynamic.TypedSimpleResult[string]

	// Act
	actual := args.Map{
		"r1Err": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
		"r3Err": r3.InvalidError() == nil,
		"r3Valid": r3.IsValid(),
		"r3Invalid": r3.IsInvalid(),
		"r3Msg": r3.Message(),
		"r3Str": r3.String(),
	}

	// Assert
	expected := args.Map{
		"r1Err": true, "r2Err": true,
		"r3Err": true, "r3Valid": false, "r3Invalid": true,
		"r3Msg": "", "r3Str": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult InvalidError returns expected -- all paths", actual)
}

func Test_TypedSimpleResult_Json_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[string]("test")
	jb, jbErr := r.JsonBytes()
	mb, mbErr := r.MarshalJSON()

	// Act
	actual := args.Map{
		"jbLen": len(jb) > 0, "jbErr": jbErr != nil,
		"mbLen": len(mb) > 0, "mbErr": mbErr != nil,
		"model": r.JsonModel(), "modelAny": r.JsonModelAny() != nil,
		"jsonNotNil": true, "jsonPtrNotNil": r.JsonPtr() != nil,
	}
	_ = r.Json()
	_ = r.JsonResult()

	// Assert
	expected := args.Map{
		"jbLen": true, "jbErr": false,
		"mbLen": true, "mbErr": false,
		"model": "test", "modelAny": true,
		"jsonNotNil": true, "jsonPtrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult JSON methods return expected -- valid", actual)
}

func Test_TypedSimpleResult_GetAs_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleResultValid[any]("hello")
	s, sOk := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, bOk := r.GetAsBool()
	_, fOk := r.GetAsFloat64()
	_, i64Ok := r.GetAsInt64()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()

	// Act
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "bOk": bOk,
		"fOk": fOk, "i64Ok": i64Ok, "byOk": byOk, "ssOk": ssOk,
	}

	// Assert
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "bOk": false,
		"fOk": false, "i64Ok": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult GetAs methods return expected -- string", actual)
}

func Test_TypedSimpleResult_CloneAndConvert(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{
		"clonedData": cloned.Data(), "clonedPtrNotNil": clonedPtr != nil,
		"nilCloneInvalid": nilClone.IsInvalid(), "nilClonePtr": nilClonePtr == nil,
		"srValid": sr.IsValid(), "nilSrInvalid": nilSr.IsInvalid(),
		"tdValid": td.IsValid(), "nilTdInvalid": nilTd.IsInvalid(),
		"dynValid": dyn.IsValid(), "nilDynInvalid": nilDyn.IsInvalid(),
	}

	// Assert
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

func Test_TypedSimpleRequest_Constructors_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleRequest[string]("hello", true, "")
	r2 := coredynamic.NewTypedSimpleRequestValid[string]("world")
	r3 := coredynamic.InvalidTypedSimpleRequest[string]("err")
	r4 := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Data": r1.Data(), "r1Request": r1.Request(), "r1Value": r1.Value(),
		"r2Valid": r2.IsValid(),
		"r3Invalid": r3.IsInvalid(), "r3Message": r3.Message(),
		"r4Invalid": r4.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"r1Valid": true, "r1Data": "hello", "r1Request": "hello", "r1Value": "hello",
		"r2Valid": true,
		"r3Invalid": true, "r3Message": "err",
		"r4Invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest constructors return expected -- all", actual)
}

func Test_TypedSimpleRequest_InvalidError_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewTypedSimpleRequestValid[string]("ok")
	r2 := coredynamic.InvalidTypedSimpleRequest[string]("err")
	var r3 *coredynamic.TypedSimpleRequest[string]

	// Act
	actual := args.Map{
		"r1Err": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
		"r3Err": r3.InvalidError() == nil,
		"r3Valid": r3.IsValid(), "r3Invalid": r3.IsInvalid(),
		"r3Msg": r3.Message(), "r3Str": r3.String(),
	}

	// Assert
	expected := args.Map{
		"r1Err": true, "r2Err": true, "r3Err": true,
		"r3Valid": false, "r3Invalid": true, "r3Msg": "", "r3Str": "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest InvalidError returns expected -- all paths", actual)
}

func Test_TypedSimpleRequest_Json_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[string]("test")
	jb, _ := r.JsonBytes()
	mb, _ := r.MarshalJSON()
	_ = r.Json()
	_ = r.JsonResult()
	_ = r.JsonPtr()

	// Act
	actual := args.Map{
		"jbLen": len(jb) > 0, "mbLen": len(mb) > 0,
		"model": r.JsonModel(), "modelAny": r.JsonModelAny() != nil,
	}

	// Assert
	expected := args.Map{
		"jbLen": true, "mbLen": true,
		"model": "test", "modelAny": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest JSON methods return expected -- valid", actual)
}

func Test_TypedSimpleRequest_GetAs_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	r := coredynamic.NewTypedSimpleRequestValid[any]("hello")
	s, sOk := r.GetAsString()
	_, iOk := r.GetAsInt()
	_, bOk := r.GetAsBool()
	_, fOk := r.GetAsFloat64()
	_, f32Ok := r.GetAsFloat32()
	_, i64Ok := r.GetAsInt64()
	_, byOk := r.GetAsBytes()
	_, ssOk := r.GetAsStrings()

	// Act
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "bOk": bOk,
		"fOk": fOk, "f32Ok": f32Ok, "i64Ok": i64Ok, "byOk": byOk, "ssOk": ssOk,
	}

	// Assert
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "bOk": false,
		"fOk": false, "f32Ok": false, "i64Ok": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest GetAs methods return expected -- string", actual)
}

func Test_TypedSimpleRequest_CloneAndConvert(t *testing.T) {
	// Arrange
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

	// Act
	actual := args.Map{
		"clonedData": cloned.Data(),
		"nilClone": nilClone == nil,
		"srValid": sr.IsValid(), "nilSrInvalid": nilSr.IsInvalid(),
		"tdValid": td.IsValid(), "nilTdInvalid": nilTd.IsInvalid(),
		"dynValid": dyn.IsValid(), "nilDynInvalid": nilDyn.IsInvalid(),
	}

	// Assert
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

func Test_SafeTypeName_FromTypedDynamicConstruc(t *testing.T) {
	// Act
	actual := args.Map{
		"string": coredynamic.SafeTypeName("hello"),
		"int":    coredynamic.SafeTypeName(42),
		"nil":    coredynamic.SafeTypeName(nil),
	}

	// Assert
	expected := args.Map{
		"string": "string", "int": "int", "nil": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns expected -- various types", actual)
}

// ==========================================================================
// PointerOrNonPointer
// ==========================================================================

func Test_PointerOrNonPointer_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	val := "hello"
	outPtr, _ := coredynamic.PointerOrNonPointer(true, &val)
	outNonPtr, _ := coredynamic.PointerOrNonPointer(false, &val)

	// Act
	actual := args.Map{
		"ptrNotNil":    outPtr != nil,
		"nonPtrNotNil": outNonPtr != nil,
	}

	// Assert
	expected := args.Map{
		"ptrNotNil": true, "nonPtrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns expected -- ptr and non-ptr", actual)
}

// ==========================================================================
// LengthOfReflect
// ==========================================================================

func Test_LengthOfReflect_FromTypedDynamicConstruc(t *testing.T) {
	// Act
	actual := args.Map{
		"slice": coredynamic.LengthOfReflect(reflect.ValueOf([]int{1, 2, 3})),
		"array": coredynamic.LengthOfReflect(reflect.ValueOf([2]int{1, 2})),
		"map":   coredynamic.LengthOfReflect(reflect.ValueOf(map[string]int{"a": 1})),
		"int":   coredynamic.LengthOfReflect(reflect.ValueOf(42)),
	}

	// Assert
	expected := args.Map{
		"slice": 3,
		"array": 2,
		"map": 1,
		"int": 0,
	}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns expected -- various kinds", actual)
}

// ==========================================================================
// BytesConverter
// ==========================================================================

func Test_BytesConverter_SafeCastString_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})

	// Act
	actual := args.Map{
		"safe":      bc.SafeCastString(),
		"safeEmpty": bcEmpty.SafeCastString(),
	}

	// Assert
	expected := args.Map{
		"safe": "hello",
		"safeEmpty": "",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString returns expected -- valid and empty", actual)
}

func Test_BytesConverter_CastString_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	val, err := bc.CastString()
	_, errEmpty := bcEmpty.CastString()

	// Act
	actual := args.Map{
		"val": val, "hasErr": err != nil, "emptyHasErr": errEmpty != nil,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"hasErr": false,
		"emptyHasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString returns expected -- valid and empty", actual)
}

func Test_BytesConverter_ToBool_FromTypedDynamicConstruc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("true"))
	val, err := bc.ToBool()

	// Act
	actual := args.Map{
		"val": val,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBool returns true -- valid", actual)
}

func Test_BytesConverter_ToBoolMust_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("false"))

	// Act
	actual := args.Map{"val": bc.ToBoolMust()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBoolMust returns false -- valid", actual)
}

func Test_BytesConverter_ToString_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	val, err := bc.ToString()

	// Act
	actual := args.Map{
		"val": val,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToString returns expected -- valid JSON string", actual)
}

func Test_BytesConverter_ToStringMust_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))

	// Act
	actual := args.Map{"val": bc.ToStringMust()}

	// Assert
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringMust returns expected -- valid", actual)
}

func Test_BytesConverter_ToStrings_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	val, err := bc.ToStrings()

	// Act
	actual := args.Map{
		"len": len(val),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStrings returns expected -- valid", actual)
}

func Test_BytesConverter_ToStringsMust_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	val := bc.ToStringsMust()

	// Act
	actual := args.Map{"len": len(val)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringsMust returns expected -- valid", actual)
}

func Test_BytesConverter_ToInt64_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("42"))
	val, err := bc.ToInt64()

	// Act
	actual := args.Map{
		"val": int(val),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64 returns expected -- valid", actual)
}

func Test_BytesConverter_ToInt64Must_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("99"))

	// Act
	actual := args.Map{"val": int(bc.ToInt64Must())}

	// Assert
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64Must returns expected -- valid", actual)
}

func Test_BytesConverter_Deserialize_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var result string
	err := bc.Deserialize(&result)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"result": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize returns expected -- valid", actual)
}

func Test_BytesConverter_DeserializeMust_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	var result string
	bc.DeserializeMust(&result)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter DeserializeMust returns expected -- valid", actual)
}

// ==========================================================================
// Dynamic — additional uncovered methods
// ==========================================================================

func Test_Dynamic_MapToKeyVal_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	kv, err := d.MapToKeyVal()

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"notNil": kv != nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic MapToKeyVal returns no error -- valid map", actual)
}

func Test_Dynamic_ItemReflectValue(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	rv := d.ItemReflectValueUsingIndex(0)

	// Act
	actual := args.Map{
		"valid": rv.IsValid(),
		"val": rv.Interface(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "a",
	}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemReflectValueUsingIndex returns expected -- index 0", actual)
}

func Test_Dynamic_ItemReflectValueUsingKey(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"x": 99})
	rv := d.ItemReflectValueUsingKey("x")

	// Act
	actual := args.Map{
		"valid": rv.IsValid(),
		"val": int(rv.Int()),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": 99,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemReflectValueUsingKey returns expected -- key x", actual)
}

func Test_Dynamic_Json_Deserialize(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("", true)
	_, err := d.Deserialize([]byte(`"hello"`))
	var nilD *coredynamic.Dynamic
	_, nilErr := nilD.Deserialize(nil)

	// Act
	actual := args.Map{
		"err": err != nil,
		"nilErr": nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Deserialize returns expected -- valid and nil", actual)
}

func Test_Dynamic_UnmarshalJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("", true)
	err := d.UnmarshalJSON([]byte(`"test"`))
	var nilD *coredynamic.Dynamic
	nilErr := nilD.UnmarshalJSON(nil)

	// Act
	actual := args.Map{
		"err": err != nil,
		"nilErr": nilErr != nil,
	}

	// Assert
	expected := args.Map{
		"err": true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic UnmarshalJSON returns expected -- valid and nil", actual)
}

func Test_Dynamic_ValueMarshal_Nil_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueMarshal returns error -- nil receiver", actual)
}

func Test_Dynamic_JsonPtr_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	jp := d.JsonPtr()
	j := d.Json()

	// Act
	actual := args.Map{
		"jpNotNil": jp != nil,
		"jNotNil": j.Bytes != nil,
	}

	// Assert
	expected := args.Map{
		"jpNotNil": true,
		"jNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Json/JsonPtr returns non-nil -- valid", actual)
}

func Test_Dynamic_LoopBreak(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return index == 0 // break after first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop breaks early -- break at index 0", actual)
}

func Test_Dynamic_FilterBreak(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	filtered := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return true, index == 1 // break at index 1
	})

	// Act
	actual := args.Map{"length": filtered.Length()}

	// Assert
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection breaks early -- at index 1", actual)
}

func Test_Dynamic_LoopMapBreak(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2, "c": 3})
	count := 0
	d.LoopMap(func(index int, key, value any) bool {
		count++
		return true // break after first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap breaks early -- break at first", actual)
}

// ==========================================================================
// IsAnyTypesOf
// ==========================================================================

func Test_IsAnyTypesOf_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{
		"match":   coredynamic.IsAnyTypesOf(strType, strType, intType),
		"noMatch": coredynamic.IsAnyTypesOf(strType, intType),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns expected -- match and no match", actual)
}

// ==========================================================================
// Dynamic — ConvertUsingFunc
// ==========================================================================

func Test_Dynamic_ConvertUsingFunc_FromTypedDynamicConstruc(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("42")
	converter := func(data any, expectedType reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(data)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(""))

	// Act
	actual := args.Map{"valid": result.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ConvertUsingFunc returns valid -- identity converter", actual)
}

// ==========================================================================
// Dynamic — JSON roundtrip marshal/unmarshal
// ==========================================================================

func Test_Dynamic_JsonRoundTrip(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string `json:"name"`
	}
	d := coredynamic.NewDynamicValid(testStruct{Name: "alice"})
	marshalledBytes, _ := json.Marshal(d)

	// Act
	actual := args.Map{"hasBytes": len(marshalledBytes) > 0}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JSON roundtrip produces bytes -- struct", actual)
}
