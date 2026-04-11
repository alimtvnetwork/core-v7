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

func Test_Dynamic_Getters_String(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"data": d.Data(), "value": d.Value(),
		"string": d.String(), "structString": d.StructString(),
		"structStringPtrNotNil": d.StructStringPtr() != nil,
		"valueString": d.ValueString(),
		"isNull": d.IsNull(), "isValid": d.IsValid(), "isInvalid": d.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"data": "hello", "value": "hello",
		"string": "hello", "structString": "hello",
		"structStringPtrNotNil": true, "valueString": "hello",
		"isNull": false, "isValid": true, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Getters returns expected -- string value", actual)
}

func Test_Dynamic_Getters_TypeChecks(t *testing.T) {
	// Arrange
	dStr := coredynamic.NewDynamicValid("hello")
	dInt := coredynamic.NewDynamicValid(42)
	dSlice := coredynamic.NewDynamicValid([]int{1, 2, 3})
	dMap := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	dStruct := coredynamic.NewDynamicValid(struct{}{})
	dFunc := coredynamic.NewDynamicValid(func() {})
	ptr := "hello"
	dPtr := coredynamic.NewDynamicValid(&ptr)

	// Act
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

	// Assert
	expected := args.Map{
		"isStringType": true, "isNumber": true, "isPrimitive": true,
		"isSliceOrArray": true, "isSliceOrArrayOrMap": true,
		"isMap": true, "isStruct": true, "isFunc": true,
		"isPointer": true, "isValueType": true, "length": 3,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic TypeChecks returns expected -- various types", actual)
}

func Test_Dynamic_Getters_ValueExtraction(t *testing.T) {
	// Arrange
	dInt := coredynamic.NewDynamicValid(42)
	dUint := coredynamic.NewDynamicValid(uint(10))
	dBool := coredynamic.NewDynamicValid(true)
	dInt64 := coredynamic.NewDynamicValid(int64(100))
	dStrings := coredynamic.NewDynamicValid([]string{"a", "b"})
	dBytes := coredynamic.NewDynamicValid([]byte("hi"))

	// Act
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

	// Assert
	expected := args.Map{
		"valueInt": 42, "valueUInt": 10, "valueBool": true,
		"valueInt64": 100, "stringsLen": 2,
		"bytesLen": 2, "bytesOk": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueExtraction returns expected -- various types", actual)
}

func Test_Dynamic_Getters_IntDefault(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("42")
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
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault returns 42 -- string 42", actual)
}

func Test_Dynamic_Getters_IntDefault_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("abc")
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
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault returns default -- invalid string", actual)
}

func Test_Dynamic_Getters_IntDefault_Nil(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(nil)
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
	expected.ShouldBeEqual(t, 0, "Dynamic IntDefault returns default -- nil data", actual)
}

func Test_Dynamic_Getters_Float64_Valid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("3.14")
	val, err := d.Float64()

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"above3": val > 3.0,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"above3": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 returns value -- valid float string", actual)
}

func Test_Dynamic_Getters_Float64_Nil(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(nil)
	_, err := d.Float64()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 returns error -- nil data", actual)
}

func Test_Dynamic_Getters_Float64_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("abc")
	_, err := d.Float64()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Float64 returns error -- invalid string", actual)
}

func Test_Dynamic_Getters_ValueNullErr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr returns nil -- valid data", actual)
}

func Test_Dynamic_Getters_ValueNullErr_Nil(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(nil)

	// Act
	actual := args.Map{"hasErr": d.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueNullErr returns error -- nil data", actual)
}

func Test_Dynamic_Getters_ValueString_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"result": d.ValueString()}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic ValueString returns empty -- nil receiver", actual)
}

func Test_Dynamic_Getters_Bytes_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	_, ok := d.Bytes()

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Bytes returns false -- nil receiver", actual)
}

func Test_Dynamic_Getters_StructStringNull(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(nil)

	// Act
	actual := args.Map{
		"isNullOrEmpty":     d.IsStructStringNullOrEmpty(),
		"isNullOrEmptyOrWs": d.IsStructStringNullOrEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isNullOrEmpty": true,
		"isNullOrEmptyOrWs": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic StructString null checks return true -- nil data", actual)
}

// ==========================================================================
// Dynamic — DynamicReflect.go coverage
// ==========================================================================

func Test_Dynamic_Reflect_Methods(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"reflectValueNotNil": d.ReflectValue() != nil,
		"reflectKind":        d.ReflectKind() == reflect.String,
		"reflectTypeName":    d.ReflectTypeName() != "",
		"reflectTypeNotNil":  d.ReflectType() != nil,
		"isReflectTypeOf":    d.IsReflectTypeOf(reflect.TypeOf("")),
		"isReflectKind":      d.IsReflectKind(reflect.String),
	}

	// Assert
	expected := args.Map{
		"reflectValueNotNil": true, "reflectKind": true,
		"reflectTypeName": true, "reflectTypeNotNil": true,
		"isReflectTypeOf": true, "isReflectKind": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Reflect Methods returns expected -- string", actual)
}

func Test_Dynamic_Reflect_Loop(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
	count := 0
	d.Loop(func(index int, item any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop iterates all -- 3 items", actual)
}

func Test_Dynamic_Reflect_Loop_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()
	called := d.Loop(func(index int, item any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Dynamic Loop returns false -- invalid", actual)
}

func Test_Dynamic_Reflect_LoopMap(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	called := d.LoopMap(func(index int, key, value any) bool { return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "Dynamic LoopMap iterates map -- 1 entry", actual)
}

func Test_Dynamic_Reflect_FilterAsDynamicCollection(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
	filtered := d.FilterAsDynamicCollection(func(index int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueString() == "b", false
	})

	// Act
	actual := args.Map{"length": filtered.Length()}

	// Assert
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "Dynamic FilterAsDynamicCollection returns 1 -- filter b", actual)
}

func Test_Dynamic_Reflect_ItemUsingIndex(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]string{"a", "b"})

	// Act
	actual := args.Map{"first": d.ItemUsingIndex(0)}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemUsingIndex returns first -- index 0", actual)
}

func Test_Dynamic_Reflect_ItemUsingKey(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"key": 42})

	// Act
	actual := args.Map{"val": d.ItemUsingKey("key")}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Dynamic ItemUsingKey returns 42 -- key lookup", actual)
}

func Test_Dynamic_Reflect_ReflectSetTo(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	var target string
	err := d.ReflectSetTo(&target)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectSetTo returns no error -- valid", actual)
}

func Test_Dynamic_Reflect_ReflectSetTo_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ReflectSetTo returns error -- nil receiver", actual)
}

// ==========================================================================
// Dynamic — DynamicJson.go coverage
// ==========================================================================

func Test_Dynamic_Json_Methods(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	jsonBytes, jsonErr := d.JsonBytesPtr()
	jsonStr, jsonStrErr := d.JsonString()
	jsonMust := d.JsonStringMust()
	marshalBytes, marshalErr := d.MarshalJSON()
	valMarshalBytes, valMarshalErr := d.ValueMarshal()
	payloadMust := d.JsonPayloadMust()

	// Act
	actual := args.Map{
		"jsonBytesLen": len(jsonBytes) > 0, "jsonErr": jsonErr != nil,
		"jsonStrNotEmpty": jsonStr != "", "jsonStrErr": jsonStrErr != nil,
		"jsonMust": jsonMust != "", "marshalLen": len(marshalBytes) > 0,
		"marshalErr": marshalErr != nil, "valMarshalLen": len(valMarshalBytes) > 0,
		"valMarshalErr": valMarshalErr != nil, "payloadLen": len(payloadMust) > 0,
	}

	// Assert
	expected := args.Map{
		"jsonBytesLen": true, "jsonErr": false,
		"jsonStrNotEmpty": true, "jsonStrErr": false,
		"jsonMust": true, "marshalLen": true,
		"marshalErr": false, "valMarshalLen": true,
		"valMarshalErr": false, "payloadLen": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Json Methods returns expected -- string value", actual)
}

func Test_Dynamic_Json_Null(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(nil)
	jsonBytes, jsonErr := d.JsonBytesPtr()

	// Act
	actual := args.Map{
		"len": len(jsonBytes),
		"hasErr": jsonErr != nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonBytesPtr returns empty -- nil data", actual)
}

func Test_Dynamic_Json_JsonModel(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{
		"notNil": d.JsonModel() != nil,
		"anyNotNil": d.JsonModelAny() != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"anyNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic JsonModel returns non-nil -- valid", actual)
}

// ==========================================================================
// DynamicStatus coverage
// ==========================================================================

func Test_DynamicStatus_InvalidNoMessage_FromDynamicGetters(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatusNoMessage()

	// Act
	actual := args.Map{
		"isValid": ds.IsValid(),
		"message": ds.Message,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus InvalidNoMessage returns invalid -- no message", actual)
}

func Test_DynamicStatus_InvalidWithMessage(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("error")

	// Act
	actual := args.Map{
		"isValid": ds.IsValid(),
		"message": ds.Message,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "error",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus InvalidWithMessage returns invalid -- with message", actual)
}

func Test_DynamicStatus_Clone_FromDynamicGetters(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("error")
	cloned := ds.Clone()

	// Act
	actual := args.Map{
		"isValid": cloned.IsValid(),
		"message": cloned.Message,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "error",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus Clone returns copy -- with message", actual)
}

func Test_DynamicStatus_ClonePtr_FromDynamicGetters(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("error")
	clonedPtr := ds.ClonePtr()
	var nilDs *coredynamic.DynamicStatus

	// Act
	actual := args.Map{
		"notNil": clonedPtr != nil,
		"nilClone": nilDs.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus ClonePtr returns expected -- valid and nil", actual)
}

// ==========================================================================
// ValueStatus coverage
// ==========================================================================

func Test_ValueStatus_InvalidNoMessage_FromDynamicGetters(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatusNoMessage()

	// Act
	actual := args.Map{
		"isValid": vs.IsValid,
		"message": vs.Message,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus InvalidNoMessage returns invalid -- no message", actual)
}

func Test_ValueStatus_InvalidWithMessage(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatus("error")

	// Act
	actual := args.Map{
		"isValid": vs.IsValid,
		"message": vs.Message,
		"valNil": vs.Value == nil,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "error",
		"valNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus InvalidWithMessage returns invalid -- with message", actual)
}

// ==========================================================================
// TypeStatus coverage
// ==========================================================================

func Test_TypeStatus_Methods(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	ts := &coredynamic.TypeStatus{
		IsSame: false, Left: strType, Right: intType,
		IsLeftPointer: false, IsRightPointer: false,
	}

	// Act
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

	// Assert
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

func Test_TypeStatus_Same(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	ts := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}

	// Act
	actual := args.Map{
		"isValid":        ts.IsValid(),
		"validationErr":  ts.ValidationError() == nil,
		"notMatchMsg":    ts.NotMatchMessage("l", "r"),
		"notMatchErr":    ts.NotMatchErr("l", "r") == nil,
		"sameRegardless": ts.IsSameRegardlessPointer(),
	}

	// Assert
	expected := args.Map{
		"isValid": true, "validationErr": true,
		"notMatchMsg": "", "notMatchErr": true, "sameRegardless": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus Same returns expected -- same types", actual)
}

func Test_TypeStatus_NilReceiver(t *testing.T) {
	// Arrange
	var ts *coredynamic.TypeStatus

	// Act
	actual := args.Map{
		"isValid": ts.IsValid(),
		"isInvalid": ts.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus nil returns safe defaults -- nil receiver", actual)
}

func Test_TypeStatus_IsEqual(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	ts1 := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}
	ts2 := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}
	ts3 := &coredynamic.TypeStatus{IsSame: false, Left: strType, Right: strType}

	// Act
	actual := args.Map{
		"sameValues": ts1.IsEqual(ts2), "diffIsSame": ts1.IsEqual(ts3),
		"bothNil": (*coredynamic.TypeStatus)(nil).IsEqual(nil),
		"leftNil": (*coredynamic.TypeStatus)(nil).IsEqual(ts1),
	}

	// Assert
	expected := args.Map{
		"sameValues": true,
		"diffIsSame": false,
		"bothNil": true,
		"leftNil": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual returns expected -- various", actual)
}

// ==========================================================================
// SimpleResult — additional coverage
// ==========================================================================

func Test_SimpleResult_InvalidNoMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidSimpleResultNoMessage()

	// Act
	actual := args.Map{
		"isValid": sr.IsValid(),
		"message": sr.Message,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidNoMessage returns invalid -- no message", actual)
}

func Test_SimpleResult_InvalidWithMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidSimpleResult("error")

	// Act
	actual := args.Map{
		"isValid": sr.IsValid(),
		"message": sr.Message,
		"invalidErr": sr.InvalidError() != nil,
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"message": "error",
		"invalidErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidWithMessage returns invalid -- with message", actual)
}

func Test_SimpleResult_Valid(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("data")

	// Act
	actual := args.Map{
		"isValid": sr.IsValid(),
		"result": sr.Result,
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"result": "data",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult Valid returns valid -- with data", actual)
}

func Test_SimpleResult_Clone_FromDynamicGetters(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResult("data", true, "")
	cloned := sr.Clone()

	// Act
	actual := args.Map{
		"result": cloned.Result,
		"isValid": cloned.IsValid(),
	}

	// Assert
	expected := args.Map{
		"result": "data",
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult Clone returns copy -- valid", actual)
}

func Test_SimpleResult_ClonePtr_FromDynamicGetters(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResult("data", true, "")
	cloned := sr.ClonePtr()
	var nilSr *coredynamic.SimpleResult

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"nilClone": nilSr.ClonePtr() == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilClone": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult ClonePtr returns expected -- valid and nil", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("hello")
	errMatch := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	errMismatch := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)

	// Act
	actual := args.Map{
		"matchErr": errMatch != nil,
		"mismatchErr": errMismatch != nil,
	}

	// Assert
	expected := args.Map{
		"matchErr": false,
		"mismatchErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch returns expected -- string vs int", actual)
}
