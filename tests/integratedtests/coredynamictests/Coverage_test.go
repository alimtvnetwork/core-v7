package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Dynamic — DynamicGetters.go coverage
// ==========================================================================

func Test_Cov_Dynamic_Getters_String(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{
		"data": d.Data(), "value": d.Value(),
		"string": d.String(), "structString": d.StructString(),
		"structStringPtrNotNil": d.StructStringPtr() != nil,
		"valueString": d.ValueString(),
		"isNull": d.IsNull(), "isValid": d.IsValid(), "isInvalid": d.IsInvalid(),
	}
	expected := args.Map{
		"data": "hello", "value": "hello",
		"string": "hello", "structString": "hello",
		"structStringPtrNotNil": true, "valueString": "hello",
		"isNull": false, "isValid": true, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Getters returns expected -- string value", actual)
}

func Test_Cov_Dynamic_Getters_TypeChecks(t *testing.T) {
	dStr := coredynamic.NewDynamicValid("hello")
	dInt := coredynamic.NewDynamicValid(42)
	dSlice := coredynamic.NewDynamicValid([]int{1, 2, 3})
	dMap := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	dStruct := coredynamic.NewDynamicValid(struct{}{})
	dFunc := coredynamic.NewDynamicValid(func() {})
	ptr := "hello"
	dPtr := coredynamic.NewDynamicValid(&ptr)
	actual := args.Map{
		"isStringType":     dStr.IsStringType(),
		"isNumber":         dInt.IsNumber(),
		"isPrimitive":      dInt.IsPrimitive(),
		"isSliceOrArray":   dSlice.IsSliceOrArray(),
		"isSliceOrArrayOrMap": dSlice.IsSliceOrArrayOrMap(),
		"isMap":            dMap.IsMap(),
		"isStruct":         dStruct.IsStruct(),
		"isFunc":           dFunc.IsFunc(),
		"isPointer":        dPtr.IsPointer(),
		"isValueType":      dStr.IsValueType(),
		"length":           dSlice.Length(),
	}
	expected := args.Map{
		"isStringType": true, "isNumber": true, "isPrimitive": true,
		"isSliceOrArray": true, "isSliceOrArrayOrMap": true,
		"isMap": true, "isStruct": true, "isFunc": true,
		"isPointer": true, "isValueType": true, "length": 3,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic TypeChecks returns expected -- various types", actual)
}

func Test_Cov_Dynamic_Getters_ValueExtraction(t *testing.T) {
	dInt := coredynamic.NewDynamicValid(42)
	dUint := coredynamic.NewDynamicValid(uint(10))
	dBool := coredynamic.NewDynamicValid(true)
	dInt64 := coredynamic.NewDynamicValid(int64(100))
	dStrings := coredynamic.NewDynamicValid([]string{"a", "b"})
	dBytes := coredynamic.NewDynamicValid([]byte("hi"))
	actual := args.Map{
		"valueInt":     dInt.ValueInt(),
		"valueUInt":    int(dUint.ValueUInt()),
		"valueBool":    dBool.ValueBool(),
		"valueInt64":   int(dInt64.ValueInt64()),
		"stringsLen":   len(dStrings.ValueStrings()),
	}
	rawBytes, bytesOk := dBytes.Bytes()
	actual["bytesLen"] = len(rawBytes)
	actual["bytesOk"] = bytesOk
	expected := args.Map{
		"valueInt": 42, "valueUInt": 10, "valueBool": true,
		"valueInt64": 100, "stringsLen": 2,
		"bytesLen": 2, "bytesOk": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueExtraction returns expected -- various types", actual)
}

func Test_Cov_Dynamic_Getters_IntDefault(t *testing.T) {
	d := coredynamic.NewDynamicValid("42")
	val, ok := d.IntDefault(0)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 42, "ok": true}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault returns 42 -- string 42", actual)
}

func Test_Cov_Dynamic_Getters_IntDefault_Invalid(t *testing.T) {
	d := coredynamic.NewDynamicValid("abc")
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault returns default -- invalid string", actual)
}

func Test_Cov_Dynamic_Getters_IntDefault_Nil(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	val, ok := d.IntDefault(99)
	actual := args.Map{"val": val, "ok": ok}
	expected := args.Map{"val": 99, "ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault returns default -- nil data", actual)
}

func Test_Cov_Dynamic_Getters_Float64_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid("3.14")
	val, err := d.Float64()
	actual := args.Map{"hasErr": err != nil, "above3": val > 3.0}
	expected := args.Map{"hasErr": false, "above3": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 returns value -- valid float string", actual)
}

func Test_Cov_Dynamic_Getters_Float64_Nil(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 returns error -- nil data", actual)
}

func Test_Cov_Dynamic_Getters_Float64_Invalid(t *testing.T) {
	d := coredynamic.NewDynamicValid("abc")
	_, err := d.Float64()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 returns error -- invalid string", actual)
}

func Test_Cov_Dynamic_Getters_ValueNullErr(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr returns nil -- valid data", actual)
}

func Test_Cov_Dynamic_Getters_ValueNullErr_Nil(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr returns error -- nil data", actual)
}

func Test_Cov_Dynamic_Getters_ValueString_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"result": d.ValueString()}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueString returns empty -- nil receiver", actual)
}

func Test_Cov_Dynamic_Getters_Bytes_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	_, ok := d.Bytes()
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Bytes returns false -- nil receiver", actual)
}

func Test_Cov_Dynamic_Getters_StructStringNull(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	actual := args.Map{
		"isNullOrEmpty":     d.IsStructStringNullOrEmpty(),
		"isNullOrEmptyOrWs": d.IsStructStringNullOrEmptyOrWhitespace(),
	}
	expected := args.Map{"isNullOrEmpty": true, "isNullOrEmptyOrWs": true}
	expected.ShouldBeEqual(t, 0, "Dynamic StructString null checks return true -- nil data", actual)
}

// ==========================================================================
// Dynamic — DynamicReflect.go coverage
// ==========================================================================

func Test_Cov_Dynamic_Reflect_Methods(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{
		"reflectValueNotNil": d.ReflectValue() != nil,
		"reflectKind":        d.ReflectKind() == reflect.String,
		"reflectTypeName":    d.ReflectTypeName() != "",
		"reflectTypeNotNil":  d.ReflectType() != nil,
		"isReflectTypeOf":    d.IsReflectTypeOf(reflect.TypeOf("")),
		"isReflectKind":      d.IsReflectKind(reflect.String),
	}
	expected := args.Map{
		"reflectValueNotNil": true, "reflectKind": true,
		"reflectTypeName": true, "reflectTypeNotNil": true,
		"isReflectTypeOf": true, "isReflectKind": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Reflect Methods returns expected -- string", actual)
}

func Test_Cov_Dynamic_Reflect_Loop(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return false
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop iterates all -- 3 items", actual)
}

func Test_Cov_Dynamic_Reflect_Loop_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	called := d.Loop(func(index int, item any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop returns false -- invalid", actual)
}

func Test_Cov_Dynamic_Reflect_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	called := d.LoopMap(func(index int, key, value any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap iterates map -- 1 entry", actual)
}

func Test_Cov_Dynamic_Reflect_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
	filtered := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueString() == "b", false
	})
	actual := args.Map{"length": filtered.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection returns 1 -- filter b", actual)
}

func Test_Cov_Dynamic_Reflect_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	actual := args.Map{"first": d.ItemUsingIndex(0)}
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemUsingIndex returns first -- index 0", actual)
}

func Test_Cov_Dynamic_Reflect_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"key": 42})
	actual := args.Map{"val": d.ItemUsingKey("key")}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemUsingKey returns 42 -- key lookup", actual)
}

func Test_Cov_Dynamic_Reflect_ReflectSetTo(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	var target string
	err := d.ReflectSetTo(&target)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectSetTo returns no error -- valid", actual)
}

func Test_Cov_Dynamic_Reflect_ReflectSetTo_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectSetTo returns error -- nil receiver", actual)
}

// ==========================================================================
// Dynamic — DynamicJson.go coverage
// ==========================================================================

func Test_Cov_Dynamic_Json_Methods(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	jsonBytes, jsonErr := d.JsonBytesPtr()
	jsonStr, jsonStrErr := d.JsonString()
	jsonMust := d.JsonStringMust()
	marshalBytes, marshalErr := d.MarshalJSON()
	valMarshalBytes, valMarshalErr := d.ValueMarshal()
	payloadMust := d.JsonPayloadMust()
	actual := args.Map{
		"jsonBytesLen": len(jsonBytes) > 0, "jsonErr": jsonErr != nil,
		"jsonStrNotEmpty": jsonStr != "", "jsonStrErr": jsonStrErr != nil,
		"jsonMust": jsonMust != "", "marshalLen": len(marshalBytes) > 0,
		"marshalErr": marshalErr != nil, "valMarshalLen": len(valMarshalBytes) > 0,
		"valMarshalErr": valMarshalErr != nil, "payloadLen": len(payloadMust) > 0,
	}
	expected := args.Map{
		"jsonBytesLen": true, "jsonErr": false,
		"jsonStrNotEmpty": true, "jsonStrErr": false,
		"jsonMust": true, "marshalLen": true,
		"marshalErr": false, "valMarshalLen": true,
		"valMarshalErr": false, "payloadLen": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Json Methods returns expected -- string value", actual)
}

func Test_Cov_Dynamic_Json_Null(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	jsonBytes, jsonErr := d.JsonBytesPtr()
	actual := args.Map{"len": len(jsonBytes), "hasErr": jsonErr != nil}
	expected := args.Map{"len": 0, "hasErr": false}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonBytesPtr returns empty -- nil data", actual)
}

func Test_Cov_Dynamic_Json_JsonModel(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"notNil": d.JsonModel() != nil, "anyNotNil": d.JsonModelAny() != nil}
	expected := args.Map{"notNil": true, "anyNotNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonModel returns non-nil -- valid", actual)
}

// ==========================================================================
// DynamicStatus coverage
// ==========================================================================

func Test_Cov_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	actual := args.Map{"isValid": ds.IsValid(), "message": ds.Message}
	expected := args.Map{"isValid": false, "message": ""}
	expected.ShouldBeEqual(t, 0, "DynamicStatus InvalidNoMessage returns invalid -- no message", actual)
}

func Test_Cov_DynamicStatus_InvalidWithMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("error")
	actual := args.Map{"isValid": ds.IsValid(), "message": ds.Message}
	expected := args.Map{"isValid": false, "message": "error"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus InvalidWithMessage returns invalid -- with message", actual)
}

func Test_Cov_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("error")
	cloned := ds.Clone()
	actual := args.Map{"isValid": cloned.IsValid(), "message": cloned.Message}
	expected := args.Map{"isValid": false, "message": "error"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus Clone returns copy -- with message", actual)
}

func Test_Cov_DynamicStatus_ClonePtr(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("error")
	clonedPtr := ds.ClonePtr()
	var nilDs *coredynamic.DynamicStatus
	actual := args.Map{"notNil": clonedPtr != nil, "nilClone": nilDs.ClonePtr() == nil}
	expected := args.Map{"notNil": true, "nilClone": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus ClonePtr returns expected -- valid and nil", actual)
}

// ==========================================================================
// ValueStatus coverage
// ==========================================================================

func Test_Cov_ValueStatus_InvalidNoMessage(t *testing.T) {
	vs := coredynamic.InvalidValueStatusNoMessage()
	actual := args.Map{"isValid": vs.IsValid, "message": vs.Message}
	expected := args.Map{"isValid": false, "message": ""}
	expected.ShouldBeEqual(t, 0, "ValueStatus InvalidNoMessage returns invalid -- no message", actual)
}

func Test_Cov_ValueStatus_InvalidWithMessage(t *testing.T) {
	vs := coredynamic.InvalidValueStatus("error")
	actual := args.Map{"isValid": vs.IsValid, "message": vs.Message, "valNil": vs.Value == nil}
	expected := args.Map{"isValid": false, "message": "error", "valNil": true}
	expected.ShouldBeEqual(t, 0, "ValueStatus InvalidWithMessage returns invalid -- with message", actual)
}

// ==========================================================================
// TypeStatus coverage
// ==========================================================================

func Test_Cov_TypeStatus_Methods(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	ts := &coredynamic.TypeStatus{
		IsSame: false, Left: strType, Right: intType,
		IsLeftPointer: false, IsRightPointer: false,
	}
	actual := args.Map{
		"isValid": ts.IsValid(), "isInvalid": ts.IsInvalid(),
		"isNotSame": ts.IsNotSame(), "isNotEqual": ts.IsNotEqualTypes(),
		"isAnyPointer": ts.IsAnyPointer(), "isBothPointer": ts.IsBothPointer(),
		"isSameRegardless": ts.IsSameRegardlessPointer(),
		"leftName": ts.LeftName(), "rightName": ts.RightName(),
		"leftFullName": ts.LeftFullName(), "rightFullName": ts.RightFullName(),
		"notMatchMsg": ts.NotMatchMessage("l", "r") != "",
		"notMatchErr": ts.NotMatchErr("l", "r") != nil,
		"validationErr": ts.ValidationError() != nil,
		"srcDestMsg": ts.NotEqualSrcDestinationMessage() != "",
		"srcDestErr": ts.NotEqualSrcDestinationErr() != nil,
	}
	expected := args.Map{
		"isValid": true, "isInvalid": false,
		"isNotSame": true, "isNotEqual": true,
		"isAnyPointer": false, "isBothPointer": false,
		"isSameRegardless": false,
		"leftName": "string", "rightName": "int",
		"leftFullName": "string", "rightFullName": "int",
		"notMatchMsg": true, "notMatchErr": true,
		"validationErr": true, "srcDestMsg": true, "srcDestErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus Methods returns expected -- different types", actual)
}

func Test_Cov_TypeStatus_Same(t *testing.T) {
	strType := reflect.TypeOf("")
	ts := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}
	actual := args.Map{
		"isValid":        ts.IsValid(),
		"validationErr":  ts.ValidationError() == nil,
		"notMatchMsg":    ts.NotMatchMessage("l", "r"),
		"notMatchErr":    ts.NotMatchErr("l", "r") == nil,
		"sameRegardless": ts.IsSameRegardlessPointer(),
	}
	expected := args.Map{
		"isValid": true, "validationErr": true,
		"notMatchMsg": "", "notMatchErr": true, "sameRegardless": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus Same returns expected -- same types", actual)
}

func Test_Cov_TypeStatus_NilReceiver(t *testing.T) {
	var ts *coredynamic.TypeStatus
	actual := args.Map{"isValid": ts.IsValid(), "isInvalid": ts.IsInvalid()}
	expected := args.Map{"isValid": false, "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus nil returns safe defaults -- nil receiver", actual)
}

func Test_Cov_TypeStatus_IsEqual(t *testing.T) {
	strType := reflect.TypeOf("")
	ts1 := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}
	ts2 := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}
	ts3 := &coredynamic.TypeStatus{IsSame: false, Left: strType, Right: strType}
	actual := args.Map{
		"sameValues": ts1.IsEqual(ts2), "diffIsSame": ts1.IsEqual(ts3),
		"bothNil": (*coredynamic.TypeStatus)(nil).IsEqual(nil),
		"leftNil": (*coredynamic.TypeStatus)(nil).IsEqual(ts1),
	}
	expected := args.Map{"sameValues": true, "diffIsSame": false, "bothNil": true, "leftNil": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual returns expected -- various", actual)
}

// ==========================================================================
// SimpleResult — additional coverage
// ==========================================================================

func Test_Cov_SimpleResult_InvalidNoMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleResultNoMessage()
	actual := args.Map{"isValid": sr.IsValid(), "message": sr.Message}
	expected := args.Map{"isValid": false, "message": ""}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidNoMessage returns invalid -- no message", actual)
}

func Test_Cov_SimpleResult_InvalidWithMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleResult("error")
	actual := args.Map{"isValid": sr.IsValid(), "message": sr.Message, "invalidErr": sr.InvalidError() != nil}
	expected := args.Map{"isValid": false, "message": "error", "invalidErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidWithMessage returns invalid -- with message", actual)
}

func Test_Cov_SimpleResult_Valid(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("data")
	actual := args.Map{"isValid": sr.IsValid(), "result": sr.Result}
	expected := args.Map{"isValid": true, "result": "data"}
	expected.ShouldBeEqual(t, 0, "SimpleResult Valid returns valid -- with data", actual)
}

func Test_Cov_SimpleResult_Clone(t *testing.T) {
	sr := coredynamic.NewSimpleResult("data", true, "")
	cloned := sr.Clone()
	actual := args.Map{"result": cloned.Result, "isValid": cloned.IsValid()}
	expected := args.Map{"result": "data", "isValid": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult Clone returns copy -- valid", actual)
}

func Test_Cov_SimpleResult_ClonePtr(t *testing.T) {
	sr := coredynamic.NewSimpleResult("data", true, "")
	cloned := sr.ClonePtr()
	var nilSr *coredynamic.SimpleResult
	actual := args.Map{"notNil": cloned != nil, "nilClone": nilSr.ClonePtr() == nil}
	expected := args.Map{"notNil": true, "nilClone": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult ClonePtr returns expected -- valid and nil", actual)
}

func Test_Cov_SimpleResult_GetErrorOnTypeMismatch(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("hello")
	errMatch := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	errMismatch := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"matchErr": errMatch != nil, "mismatchErr": errMismatch != nil}
	expected := args.Map{"matchErr": false, "mismatchErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch returns expected -- string vs int", actual)
}
