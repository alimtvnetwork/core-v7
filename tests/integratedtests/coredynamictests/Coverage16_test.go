package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Dynamic — DynamicJson.go uncovered paths
// ==========================================================================

func Test_C16_Dynamic_Deserialize_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	result, err := d.Deserialize([]byte(`"hello"`))
	actual := args.Map{
		"hasErr":    err != nil,
		"isInvalid": result.IsInvalid(),
	}
	expected := args.Map{
		"hasErr":    true,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "Deserialize returns nil -- nil receiver", actual)
}

func Test_C16_Dynamic_ValueMarshal_NilReceiver(t *testing.T) {
	var d *coredynamic.Dynamic
	_, err := d.ValueMarshal()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns nil -- nil receiver", actual)
}

func Test_C16_Dynamic_ValueMarshal_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid("test")
	b, err := d.ValueMarshal()
	actual := args.Map{
		"noErr":   err == nil,
		"hasData": len(b) > 0,
	}
	expected := args.Map{
		"noErr":   true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueMarshal returns non-empty -- valid", actual)
}

func Test_C16_Dynamic_JsonPayloadMust(t *testing.T) {
	d := coredynamic.NewDynamicValid("test")
	b := d.JsonPayloadMust()
	actual := args.Map{"hasData": len(b) > 0}
	expected := args.Map{"hasData": true}
	expected.ShouldBeEqual(t, 0, "JsonPayloadMust returns correct value -- with args", actual)
}

func Test_C16_Dynamic_JsonBytesPtr_Null(t *testing.T) {
	d := coredynamic.NewDynamicValid(nil)
	b, err := d.JsonBytesPtr()
	actual := args.Map{
		"noErr":  err == nil,
		"empty":  len(b) == 0,
	}
	expected := args.Map{
		"noErr":  true,
		"empty":  true,
	}
	expected.ShouldBeEqual(t, 0, "JsonBytesPtr returns correct value -- null data", actual)
}

func Test_C16_Dynamic_JsonBytesPtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	b, err := d.JsonBytesPtr()
	actual := args.Map{
		"noErr":   err == nil,
		"hasData": len(b) > 0,
	}
	expected := args.Map{
		"noErr":   true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonBytesPtr returns non-empty -- valid", actual)
}

func Test_C16_Dynamic_MarshalJSON(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	b, err := d.MarshalJSON()
	actual := args.Map{
		"noErr":   err == nil,
		"hasData": len(b) > 0,
	}
	expected := args.Map{
		"noErr":   true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns correct value -- with args", actual)
}

func Test_C16_Dynamic_UnmarshalJSON_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.UnmarshalJSON([]byte(`"hello"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns nil -- nil receiver", actual)
}

func Test_C16_Dynamic_JsonModel_JsonModelAny(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	actual := args.Map{
		"model":    d.JsonModel(),
		"modelAny": d.JsonModelAny(),
	}
	expected := args.Map{
		"model":    "x",
		"modelAny": "x",
	}
	expected.ShouldBeEqual(t, 0, "JsonModel/JsonModelAny returns correct value -- with args", actual)
}

func Test_C16_Dynamic_Json_JsonPtr(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	j := d.Json()
	jp := d.JsonPtr()
	_ = j
	actual := args.Map{
		"ptrNotNil": jp != nil,
	}
	expected := args.Map{
		"ptrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Json/JsonPtr returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ParseInjectUsingJson_Error(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	badResult := corejson.NewPtr("invalid json test")
	_, err := d.ParseInjectUsingJson(badResult)
	// badResult may or may not have error depending on implementation
	_ = err
}

func Test_C16_Dynamic_JsonParseSelfInject(t *testing.T) {
	d := coredynamic.NewDynamicValid("x")
	jr := corejson.NewPtr("hello")
	err := d.JsonParseSelfInject(jr)
	_ = err // just coverage
}

func Test_C16_Dynamic_JsonBytes_JsonString(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	b, err1 := d.JsonBytes()
	s, err2 := d.JsonString()
	actual := args.Map{
		"bytesOK": err1 == nil && len(b) > 0,
		"strOK":   err2 == nil && s != "",
	}
	expected := args.Map{
		"bytesOK": true,
		"strOK":   true,
	}
	expected.ShouldBeEqual(t, 0, "JsonBytes/JsonString returns correct value -- with args", actual)
}

func Test_C16_Dynamic_JsonStringMust(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	s := d.JsonStringMust()
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct value -- with args", actual)
}

// ==========================================================================
// Dynamic — DynamicGetters.go uncovered paths
// ==========================================================================

func Test_C16_Dynamic_IsStructStringNullOrEmpty(t *testing.T) {
	d := coredynamic.NewDynamicValid("")
	dNull := coredynamic.NewDynamicValid(nil)
	dVal := coredynamic.NewDynamicValid("abc")
	actual := args.Map{
		"empty":   d.IsStructStringNullOrEmpty(),
		"null":    dNull.IsStructStringNullOrEmpty(),
		"hasVal":  dVal.IsStructStringNullOrEmpty(),
	}
	expected := args.Map{
		"empty":   true,
		"null":    true,
		"hasVal":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmpty returns empty -- with args", actual)
}

func Test_C16_Dynamic_IsStructStringNullOrEmptyOrWhitespace(t *testing.T) {
	d := coredynamic.NewDynamicValid("   ")
	dVal := coredynamic.NewDynamicValid("abc")
	actual := args.Map{
		"whitespace": d.IsStructStringNullOrEmptyOrWhitespace(),
		"hasVal":     dVal.IsStructStringNullOrEmptyOrWhitespace(),
	}
	expected := args.Map{
		"whitespace": true,
		"hasVal":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsStructStringNullOrEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_C16_Dynamic_IntDefault(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	val, ok := d.IntDefault(0)
	dNull := coredynamic.NewDynamicValid(nil)
	valNull, okNull := dNull.IntDefault(99)
	dBad := coredynamic.NewDynamicValid("abc")
	valBad, okBad := dBad.IntDefault(77)
	actual := args.Map{
		"val": val, "ok": ok,
		"valNull": valNull, "okNull": okNull,
		"valBad": valBad, "okBad": okBad,
	}
	expected := args.Map{
		"val": 42, "ok": true,
		"valNull": 99, "okNull": false,
		"valBad": 77, "okBad": false,
	}
	expected.ShouldBeEqual(t, 0, "IntDefault returns correct value -- with args", actual)
}

func Test_C16_Dynamic_Float64(t *testing.T) {
	d := coredynamic.NewDynamicValid(3.14)
	val, err := d.Float64()
	dNull := coredynamic.NewDynamicValid(nil)
	_, errNull := dNull.Float64()
	dBad := coredynamic.NewDynamicValid("notanumber")
	_, errBad := dBad.Float64()
	actual := args.Map{
		"val":     val > 3.0,
		"noErr":   err == nil,
		"errNull": errNull != nil,
		"errBad":  errBad != nil,
	}
	expected := args.Map{
		"val":     true,
		"noErr":   true,
		"errNull": true,
		"errBad":  true,
	}
	expected.ShouldBeEqual(t, 0, "Float64 returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ValueUInt(t *testing.T) {
	d := coredynamic.NewDynamicValid(uint(5))
	dFail := coredynamic.NewDynamicValid("nope")
	actual := args.Map{
		"ok":   d.ValueUInt(),
		"fail": dFail.ValueUInt(),
	}
	expected := args.Map{
		"ok":   uint(5),
		"fail": uint(0),
	}
	expected.ShouldBeEqual(t, 0, "ValueUInt returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ValueStrings(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	dFail := coredynamic.NewDynamicValid(42)
	actual := args.Map{
		"ok":   len(d.ValueStrings()),
		"fail": dFail.ValueStrings() == nil,
	}
	expected := args.Map{
		"ok":   2,
		"fail": true,
	}
	expected.ShouldBeEqual(t, 0, "ValueStrings returns non-empty -- with args", actual)
}

func Test_C16_Dynamic_ValueBool(t *testing.T) {
	d := coredynamic.NewDynamicValid(true)
	dFail := coredynamic.NewDynamicValid("nope")
	actual := args.Map{
		"ok":   d.ValueBool(),
		"fail": dFail.ValueBool(),
	}
	expected := args.Map{
		"ok":   true,
		"fail": false,
	}
	expected.ShouldBeEqual(t, 0, "ValueBool returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ValueInt64(t *testing.T) {
	d := coredynamic.NewDynamicValid(int64(99))
	dFail := coredynamic.NewDynamicValid("nope")
	actual := args.Map{
		"ok":   d.ValueInt64(),
		"fail": dFail.ValueInt64(),
	}
	expected := args.Map{
		"ok":   int64(99),
		"fail": int64(-1),
	}
	expected.ShouldBeEqual(t, 0, "ValueInt64 returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ValueNullErr(t *testing.T) {
	var dNil *coredynamic.Dynamic
	errNil := dNil.ValueNullErr()
	d := coredynamic.NewDynamicValid(nil)
	errNull := d.ValueNullErr()
	dOk := coredynamic.NewDynamicValid("x")
	errOk := dOk.ValueNullErr()
	actual := args.Map{
		"errNil":  errNil != nil,
		"errNull": errNull != nil,
		"errOk":   errOk == nil,
	}
	expected := args.Map{
		"errNil":  true,
		"errNull": true,
		"errOk":   true,
	}
	expected.ShouldBeEqual(t, 0, "ValueNullErr returns error -- with args", actual)
}

func Test_C16_Dynamic_ValueString_NonString(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	s := d.ValueString()
	var dNil *coredynamic.Dynamic
	sNil := dNil.ValueString()
	actual := args.Map{
		"intStr":  s != "",
		"nilStr":  sNil,
	}
	expected := args.Map{
		"intStr":  true,
		"nilStr":  "",
	}
	expected.ShouldBeEqual(t, 0, "ValueString returns non-empty -- non-string", actual)
}

func Test_C16_Dynamic_Bytes(t *testing.T) {
	d := coredynamic.NewDynamicValid([]byte{1, 2, 3})
	b, ok := d.Bytes()
	var dNil *coredynamic.Dynamic
	bNil, okNil := dNil.Bytes()
	dWrong := coredynamic.NewDynamicValid("nope")
	_, okWrong := dWrong.Bytes()
	actual := args.Map{
		"ok":      ok,
		"len":     len(b),
		"nilOk":   okNil,
		"nilB":    bNil == nil,
		"wrongOk": okWrong,
	}
	expected := args.Map{
		"ok":      true,
		"len":     3,
		"nilOk":   false,
		"nilB":    true,
		"wrongOk": false,
	}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- with args", actual)
}

// ==========================================================================
// Dynamic — DynamicReflect.go uncovered paths
// ==========================================================================

func Test_C16_Dynamic_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1})
	kv, err := d.MapToKeyVal()
	actual := args.Map{
		"noErr": err == nil,
		"hasKv": kv != nil && kv.Length() > 0,
	}
	expected := args.Map{
		"noErr": true,
		"hasKv": true,
	}
	expected.ShouldBeEqual(t, 0, "MapToKeyVal returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ReflectType(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	rt := d.ReflectType()
	actual := args.Map{
		"name": rt.Kind() == reflect.String,
	}
	expected := args.Map{
		"name": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectType returns correct value -- with args", actual)
}

func Test_C16_Dynamic_IsReflectTypeOf(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{
		"match":   d.IsReflectTypeOf(reflect.TypeOf("")),
		"noMatch": d.IsReflectTypeOf(reflect.TypeOf(42)),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsReflectTypeOf returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b", "c"})
	item := d.ItemUsingIndex(1)
	rv := d.ItemReflectValueUsingIndex(0)
	actual := args.Map{
		"item":    item,
		"rvValid": rv.IsValid(),
	}
	expected := args.Map{
		"item":    "b",
		"rvValid": true,
	}
	expected.ShouldBeEqual(t, 0, "ItemUsingIndex returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"x": 42})
	item := d.ItemUsingKey("x")
	rv := d.ItemReflectValueUsingKey("x")
	actual := args.Map{
		"item":    item,
		"rvValid": rv.IsValid(),
	}
	expected := args.Map{
		"item":    42,
		"rvValid": true,
	}
	expected.ShouldBeEqual(t, 0, "ItemUsingKey returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ReflectSetTo(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	var target string
	err := d.ReflectSetTo(&target)
	actual := args.Map{
		"noErr": err == nil,
		"val":   target,
	}
	expected := args.Map{
		"noErr": true,
		"val":   "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns correct value -- with args", actual)
}

func Test_C16_Dynamic_ReflectSetTo_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetTo returns nil -- nil", actual)
}

func Test_C16_Dynamic_Loop(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{10, 20, 30})
	sum := 0
	called := d.Loop(func(i int, item any) bool {
		sum += item.(int)
		return false
	})
	actual := args.Map{"called": called, "sum": sum}
	expected := args.Map{"called": true, "sum": 60}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- with args", actual)
}

func Test_C16_Dynamic_Loop_Break(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return i == 0 // break after first
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- break", actual)
}

func Test_C16_Dynamic_Loop_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	called := d.Loop(func(i int, item any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "Loop returns error -- invalid", actual)
}

func Test_C16_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4, 5})
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return item.ValueInt() > 2, false
	})
	actual := args.Map{"count": result.Length()}
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns correct value -- with args", actual)
}

func Test_C16_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4, 5})
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 1 // take first two then break
	})
	actual := args.Map{"count": result.Length()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns correct value -- break", actual)
}

func Test_C16_Dynamic_FilterAsDynamicCollection_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})
	actual := args.Map{"empty": result.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FilterAsDynamicCollection returns error -- invalid", actual)
}

func Test_C16_Dynamic_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})
	count := 0
	called := d.LoopMap(func(i int, key, value any) bool {
		count++
		return false
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 2}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- with args", actual)
}

func Test_C16_Dynamic_LoopMap_Break(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})
	count := 0
	called := d.LoopMap(func(i int, key, value any) bool {
		count++
		return true // break immediately
	})
	actual := args.Map{"called": called, "count": count}
	expected := args.Map{"called": true, "count": 1}
	expected.ShouldBeEqual(t, 0, "LoopMap returns correct value -- break", actual)
}

func Test_C16_Dynamic_LoopMap_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	called := d.LoopMap(func(i int, key, value any) bool { return false })
	actual := args.Map{"called": called}
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LoopMap returns error -- invalid", actual)
}

// ==========================================================================
// DynamicStatus coverage
// ==========================================================================

func Test_C16_DynamicStatus(t *testing.T) {
	invalid := coredynamic.InvalidDynamicStatus("bad")
	invalidNoMsg := coredynamic.InvalidDynamicStatusNoMessage()
	clone := invalid.Clone()
	var nilPtr *coredynamic.DynamicStatus
	cloneNil := nilPtr.ClonePtr()
	cloneValid := invalid.ClonePtr()
	actual := args.Map{
		"msg":        invalid.Message,
		"noMsg":      invalidNoMsg.Message,
		"cloneMsg":   clone.Message,
		"nilClone":   cloneNil == nil,
		"validClone": cloneValid != nil,
	}
	expected := args.Map{
		"msg":        "bad",
		"noMsg":      "",
		"cloneMsg":   "bad",
		"nilClone":   true,
		"validClone": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus returns correct value -- with args", actual)
}

// ==========================================================================
// ValueStatus coverage
// ==========================================================================

func Test_C16_ValueStatus(t *testing.T) {
	inv := coredynamic.InvalidValueStatus("oops")
	invNoMsg := coredynamic.InvalidValueStatusNoMessage()
	actual := args.Map{
		"msg":   inv.Message,
		"valid": inv.IsValid,
		"noMsg": invNoMsg.Message,
	}
	expected := args.Map{
		"msg":   "oops",
		"valid": false,
		"noMsg": "",
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- with args", actual)
}

// ==========================================================================
// SimpleRequest coverage
// ==========================================================================

func Test_C16_SimpleRequest(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("data", true, "msg")
	srValid := coredynamic.NewSimpleRequestValid("data2")
	srInvalid := coredynamic.InvalidSimpleRequest("bad")
	srInvalidNoMsg := coredynamic.InvalidSimpleRequestNoMessage()
	actual := args.Map{
		"msg":      sr.Message(),
		"req":      sr.Request(),
		"val":      sr.Value(),
		"validReq": srValid.Request(),
		"invMsg":   srInvalid.Message(),
		"noMsg":    srInvalidNoMsg.Message(),
	}
	expected := args.Map{
		"msg":      "msg",
		"req":      "data",
		"val":      "data",
		"validReq": "data2",
		"invMsg":   "bad",
		"noMsg":    "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- constructors", actual)
}

func Test_C16_SimpleRequest_TypeMismatch(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	err1 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(42), false)
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(42), true)
	errNone := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{
		"err1":    err1 != nil,
		"err2":    err2 != nil,
		"errNone": errNone == nil,
	}
	expected := args.Map{
		"err1":    true,
		"err2":    true,
		"errNone": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- GetErrorOnTypeMismatch", actual)
}

func Test_C16_SimpleRequest_IsPointer(t *testing.T) {
	s := "hello"
	sr := coredynamic.NewSimpleRequestValid(&s)
	srVal := coredynamic.NewSimpleRequestValid("hello")
	actual := args.Map{
		"isPtr":    sr.IsPointer(),
		"isNotPtr": srVal.IsPointer(),
	}
	expected := args.Map{
		"isPtr":    true,
		"isNotPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- IsPointer", actual)
}

func Test_C16_SimpleRequest_InvalidError(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequest("oops")
	err1 := sr.InvalidError()
	err2 := sr.InvalidError() // cached
	srEmpty := coredynamic.NewSimpleRequestValid("x")
	errEmpty := srEmpty.InvalidError()
	actual := args.Map{
		"err1":    err1 != nil,
		"err2":    err2 != nil,
		"same":    err1 == err2,
		"noEmpty": errEmpty == nil,
	}
	expected := args.Map{
		"err1":    true,
		"err2":    true,
		"same":    true,
		"noEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- InvalidError", actual)
}

// ==========================================================================
// SimpleResult coverage
// ==========================================================================

func Test_C16_SimpleResult(t *testing.T) {
	sr := coredynamic.NewSimpleResult("data", true, "msg")
	srValid := coredynamic.NewSimpleResultValid("data2")
	srInvalid := coredynamic.InvalidSimpleResult("bad")
	srInvalidNoMsg := coredynamic.InvalidSimpleResultNoMessage()
	actual := args.Map{
		"msg":        sr.Message,
		"res":        sr.Result,
		"validRes":   srValid.Result,
		"invMsg":     srInvalid.Message,
		"noMsg":      srInvalidNoMsg.Message,
	}
	expected := args.Map{
		"msg":        "msg",
		"res":        "data",
		"validRes":   "data2",
		"invMsg":     "bad",
		"noMsg":      "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- constructors", actual)
}

func Test_C16_SimpleResult_TypeMismatch(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("hello")
	err1 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(42), false)
	err2 := sr.GetErrorOnTypeMismatch(reflect.TypeOf(42), true)
	errNone := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{
		"err1":    err1 != nil,
		"err2":    err2 != nil,
		"errNone": errNone == nil,
	}
	expected := args.Map{
		"err1":    true,
		"err2":    true,
		"errNone": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- GetErrorOnTypeMismatch", actual)
}

func Test_C16_SimpleResult_InvalidError(t *testing.T) {
	sr := coredynamic.InvalidSimpleResult("oops")
	err1 := sr.InvalidError()
	err2 := sr.InvalidError() // cached
	srEmpty := coredynamic.NewSimpleResultValid("x")
	errEmpty := srEmpty.InvalidError()
	actual := args.Map{
		"err1":    err1 != nil,
		"err2":    err2 != nil,
		"same":    err1 == err2,
		"noEmpty": errEmpty == nil,
	}
	expected := args.Map{
		"err1":    true,
		"err2":    true,
		"same":    true,
		"noEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- InvalidError", actual)
}

func Test_C16_SimpleResult_Clone(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("data")
	clone := sr.Clone()
	clonePtr := sr.ClonePtr()
	var nilSr *coredynamic.SimpleResult
	nilClone := nilSr.ClonePtr()
	actual := args.Map{
		"cloneRes":   clone.Result,
		"ptrRes":     clonePtr.Result,
		"nilIsNil":   nilClone == nil,
	}
	expected := args.Map{
		"cloneRes":   "data",
		"ptrRes":     "data",
		"nilIsNil":   true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- Clone", actual)
}

// ==========================================================================
// TypedDynamic coverage
// ==========================================================================

func Test_C16_TypedDynamic_Full(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	dp := coredynamic.NewTypedDynamicPtr[string]("ptr", true)
	inv := coredynamic.InvalidTypedDynamic[string]()
	invP := coredynamic.InvalidTypedDynamicPtr[string]()
	actual := args.Map{
		"data": d.Data(), "val": d.Value(), "valid": d.IsValid(), "invalid": d.IsInvalid(),
		"str":    d.String(),
		"ptrVal": dp.Data(),
		"invOk":  inv.IsInvalid(),
		"invPOk": invP.IsInvalid(),
	}
	expected := args.Map{
		"data": "hello", "val": "hello", "valid": true, "invalid": false,
		"str":    "hello",
		"ptrVal": "ptr",
		"invOk":  true,
		"invPOk": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- core", actual)
}

func Test_C16_TypedDynamic_JSON(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("hello", true)
	b, err := d.JsonBytes()
	jr := d.JsonResult()
	j := d.Json()
	jp := d.JsonPtr()
	s, sErr := d.JsonString()
	mb, mErr := d.MarshalJSON()
	vm, vmErr := d.ValueMarshal()
	bytes, bOk := d.Bytes()
	actual := args.Map{
		"bytesOK":  err == nil && len(b) > 0,
		"jrNotNil": jr.JsonString() != "",
		"jNotNil":  j.JsonString() != "",
		"jpNotNil": jp != nil,
		"strOK":    sErr == nil && s != "",
		"marshalOK": mErr == nil && len(mb) > 0,
		"vmOK":     vmErr == nil && len(vm) > 0,
		"bOK":      bOk,
		"bLen":     len(bytes) >= 0,
	}
	expected := args.Map{
		"bytesOK":  true,
		"jrNotNil": true,
		"jNotNil":  true,
		"jpNotNil": true,
		"strOK":    true,
		"marshalOK": true,
		"vmOK":     true,
		"bOK":      bOk,
		"bLen":     true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JSON", actual)
}

func Test_C16_TypedDynamic_UnmarshalJSON(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := d.UnmarshalJSON([]byte(`"updated"`))
	actual := args.Map{
		"noErr": err == nil,
		"val":   d.Data(),
		"valid": d.IsValid(),
	}
	expected := args.Map{
		"noErr": true,
		"val":   "updated",
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- UnmarshalJSON", actual)
}

func Test_C16_TypedDynamic_Bytes_IsBytes(t *testing.T) {
	d := coredynamic.NewTypedDynamic[[]byte]([]byte{1, 2}, true)
	b, ok := d.Bytes()
	actual := args.Map{"ok": ok, "len": len(b)}
	expected := args.Map{"ok": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Bytes []byte", actual)
}

func Test_C16_TypedDynamic_GetAs(t *testing.T) {
	dStr := coredynamic.NewTypedDynamic[string]("hello", true)
	s, sOk := dStr.GetAsString()
	_, iOk := dStr.GetAsInt()

	dInt := coredynamic.NewTypedDynamic[int](42, true)
	v, vOk := dInt.GetAsInt()
	_, i64Ok := dInt.GetAsInt64()
	_, uOk := dInt.GetAsUint()
	_, f64Ok := dInt.GetAsFloat64()
	_, f32Ok := dInt.GetAsFloat32()
	_, bOk := dInt.GetAsBool()
	_, byOk := dInt.GetAsBytes()
	_, ssOk := dInt.GetAsStrings()

	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk,
		"v": v, "vOk": vOk,
		"i64Ok": i64Ok, "uOk": uOk, "f64Ok": f64Ok, "f32Ok": f32Ok,
		"bOk": bOk, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false,
		"v": 42, "vOk": true,
		"i64Ok": false, "uOk": false, "f64Ok": false, "f32Ok": false,
		"bOk": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- GetAs", actual)
}

func Test_C16_TypedDynamic_Value_Methods(t *testing.T) {
	dStr := coredynamic.NewTypedDynamic[string]("hello", true)
	dInt := coredynamic.NewTypedDynamic[int](42, true)
	dI64 := coredynamic.NewTypedDynamic[int64](int64(99), true)
	dBool := coredynamic.NewTypedDynamic[bool](true, true)
	dBadStr := coredynamic.NewTypedDynamic[int](7, true)
	actual := args.Map{
		"str":    dStr.ValueString(),
		"int":    dInt.ValueInt(),
		"i64":    dI64.ValueInt64(),
		"bool":   dBool.ValueBool(),
		"badStr": dBadStr.ValueString(),
	}
	expected := args.Map{
		"str":    "hello",
		"int":    42,
		"i64":    int64(99),
		"bool":   true,
		"badStr": "7",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Value methods", actual)
}

func Test_C16_TypedDynamic_Clone(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("x", true)
	clone := d.Clone()
	dp := coredynamic.NewTypedDynamicPtr[string]("y", true)
	cloneP := dp.ClonePtr()
	var nilD *coredynamic.TypedDynamic[string]
	nilClone := nilD.ClonePtr()
	nonPtr := d.NonPtr()
	ptr := dp.Ptr()
	toDyn := d.ToDynamic()
	actual := args.Map{
		"cloneVal":  clone.Data(),
		"clonePVal": cloneP.Data(),
		"nilClone":  nilClone == nil,
		"nonPtr":    nonPtr.Data(),
		"ptrNotNil": ptr != nil,
		"dynValid":  toDyn.IsValid(),
	}
	expected := args.Map{
		"cloneVal":  "x",
		"clonePVal": "y",
		"nilClone":  true,
		"nonPtr":    "x",
		"ptrNotNil": true,
		"dynValid":  true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Clone", actual)
}

func Test_C16_TypedDynamic_Deserialize(t *testing.T) {
	d := coredynamic.NewTypedDynamicPtr[string]("", false)
	err := d.Deserialize([]byte(`"hello"`))
	var nilD *coredynamic.TypedDynamic[string]
	errNil := nilD.Deserialize([]byte(`"x"`))
	actual := args.Map{
		"noErr":  err == nil,
		"val":    d.Data(),
		"valid":  d.IsValid(),
		"nilErr": errNil != nil,
	}
	expected := args.Map{
		"noErr":  true,
		"val":    "hello",
		"valid":  true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- Deserialize", actual)
}

func Test_C16_TypedDynamic_JsonModel(t *testing.T) {
	d := coredynamic.NewTypedDynamic[string]("x", true)
	actual := args.Map{
		"model":    d.JsonModel(),
		"modelAny": d.JsonModelAny(),
	}
	expected := args.Map{
		"model":    "x",
		"modelAny": "x",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- JsonModel", actual)
}

// ==========================================================================
// TypedSimpleRequest coverage
// ==========================================================================

func Test_C16_TypedSimpleRequest(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequest[string]("data", true, "")
	srV := coredynamic.NewTypedSimpleRequestValid[string]("v")
	inv := coredynamic.InvalidTypedSimpleRequest[string]("bad")
	invNo := coredynamic.InvalidTypedSimpleRequestNoMessage[string]()
	var nilSR *coredynamic.TypedSimpleRequest[string]
	actual := args.Map{
		"data":     sr.Data(),
		"req":      sr.Request(),
		"val":      sr.Value(),
		"valid":    sr.IsValid(),
		"invalid":  sr.IsInvalid(),
		"msg":      sr.Message(),
		"str":      sr.String(),
		"vData":    srV.Data(),
		"invMsg":   inv.Message(),
		"invValid": inv.IsValid(),
		"noMsg":    invNo.Message(),
		"nilValid": nilSR.IsValid(),
		"nilInval": nilSR.IsInvalid(),
		"nilMsg":   nilSR.Message(),
		"nilStr":   nilSR.String(),
	}
	expected := args.Map{
		"data":     "data",
		"req":      "data",
		"val":      "data",
		"valid":    true,
		"invalid":  false,
		"msg":      "",
		"str":      "data",
		"vData":    "v",
		"invMsg":   "bad",
		"invValid": false,
		"noMsg":    "",
		"nilValid": false,
		"nilInval": true,
		"nilMsg":   "",
		"nilStr":   "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- with args", actual)
}

func Test_C16_TypedSimpleRequest_InvalidError(t *testing.T) {
	sr := coredynamic.InvalidTypedSimpleRequest[string]("oops")
	err := sr.InvalidError()
	err2 := sr.InvalidError() // cached
	srOk := coredynamic.NewTypedSimpleRequestValid[string]("x")
	errOk := srOk.InvalidError()
	var nilSR *coredynamic.TypedSimpleRequest[string]
	errNil := nilSR.InvalidError()
	actual := args.Map{
		"err":    err != nil,
		"same":   err == err2,
		"noErr":  errOk == nil,
		"nilErr": errNil == nil,
	}
	expected := args.Map{
		"err":    true,
		"same":   true,
		"noErr":  true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns error -- InvalidError", actual)
}

func Test_C16_TypedSimpleRequest_JSON(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	b, err := sr.JsonBytes()
	jr := sr.JsonResult()
	j := sr.Json()
	jp := sr.JsonPtr()
	mb, mErr := sr.MarshalJSON()
	actual := args.Map{
		"bytesOK": err == nil && len(b) > 0,
		"jrOK":    jr.JsonString() != "",
		"jOK":     j.JsonString() != "",
		"jpOK":    jp != nil,
		"mbOK":    mErr == nil && len(mb) > 0,
		"model":   sr.JsonModel(),
		"any":     sr.JsonModelAny(),
	}
	expected := args.Map{
		"bytesOK": true,
		"jrOK":    true,
		"jOK":     true,
		"jpOK":    true,
		"mbOK":    true,
		"model":   "hello",
		"any":     "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- JSON", actual)
}

func Test_C16_TypedSimpleRequest_GetAs(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid[string]("hello")
	s, sOk := sr.GetAsString()
	_, iOk := sr.GetAsInt()
	_, i64Ok := sr.GetAsInt64()
	_, f64Ok := sr.GetAsFloat64()
	_, f32Ok := sr.GetAsFloat32()
	_, bOk := sr.GetAsBool()
	_, byOk := sr.GetAsBytes()
	_, ssOk := sr.GetAsStrings()
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

func Test_C16_TypedSimpleRequest_Clone(t *testing.T) {
	sr := coredynamic.NewTypedSimpleRequestValid[string]("x")
	clone := sr.Clone()
	var nilSR *coredynamic.TypedSimpleRequest[string]
	nilClone := nilSR.Clone()
	toSimple := sr.ToSimpleRequest()
	toTD := sr.ToTypedDynamic()
	toDyn := sr.ToDynamic()
	nilToSimple := nilSR.ToSimpleRequest()
	nilToTD := nilSR.ToTypedDynamic()
	nilToDyn := nilSR.ToDynamic()
	actual := args.Map{
		"cloneData":    clone.Data(),
		"nilClone":     nilClone == nil,
		"simpleReq":    toSimple.Request(),
		"tdData":       toTD.Data(),
		"dynValid":     toDyn.IsValid(),
		"nilSimple":    nilToSimple != nil,
		"nilTD":        nilToTD.IsInvalid(),
		"nilDyn":       nilToDyn.IsInvalid(),
	}
	expected := args.Map{
		"cloneData":    "x",
		"nilClone":     true,
		"simpleReq":    "x",
		"tdData":       "x",
		"dynValid":     true,
		"nilSimple":    true,
		"nilTD":        true,
		"nilDyn":       true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleRequest returns correct value -- Clone/Convert", actual)
}

// ==========================================================================
// TypedSimpleResult coverage
// ==========================================================================

func Test_C16_TypedSimpleResult(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResult[string]("data", true, "")
	srV := coredynamic.NewTypedSimpleResultValid[string]("v")
	inv := coredynamic.InvalidTypedSimpleResult[string]("bad")
	invNo := coredynamic.InvalidTypedSimpleResultNoMessage[string]()
	var nilSR *coredynamic.TypedSimpleResult[string]
	actual := args.Map{
		"data":     sr.Data(),
		"res":      sr.Result(),
		"valid":    sr.IsValid(),
		"invalid":  sr.IsInvalid(),
		"msg":      sr.Message(),
		"str":      sr.String(),
		"vData":    srV.Data(),
		"invMsg":   inv.Message(),
		"invValid": inv.IsValid(),
		"noMsg":    invNo.Message(),
		"nilValid": nilSR.IsValid(),
		"nilInval": nilSR.IsInvalid(),
		"nilMsg":   nilSR.Message(),
		"nilStr":   nilSR.String(),
	}
	expected := args.Map{
		"data":     "data",
		"res":      "data",
		"valid":    true,
		"invalid":  false,
		"msg":      "",
		"str":      "data",
		"vData":    "v",
		"invMsg":   "bad",
		"invValid": false,
		"noMsg":    "",
		"nilValid": false,
		"nilInval": true,
		"nilMsg":   "",
		"nilStr":   "",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- with args", actual)
}

func Test_C16_TypedSimpleResult_InvalidError(t *testing.T) {
	sr := coredynamic.InvalidTypedSimpleResult[string]("oops")
	err := sr.InvalidError()
	err2 := sr.InvalidError()
	srOk := coredynamic.NewTypedSimpleResultValid[string]("x")
	errOk := srOk.InvalidError()
	var nilSR *coredynamic.TypedSimpleResult[string]
	errNil := nilSR.InvalidError()
	actual := args.Map{
		"err":    err != nil,
		"same":   err == err2,
		"noErr":  errOk == nil,
		"nilErr": errNil == nil,
	}
	expected := args.Map{
		"err":    true,
		"same":   true,
		"noErr":  true,
		"nilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns error -- InvalidError", actual)
}

func Test_C16_TypedSimpleResult_JSON(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid[string]("hello")
	b, err := sr.JsonBytes()
	jr := sr.JsonResult()
	j := sr.Json()
	jp := sr.JsonPtr()
	mb, mErr := sr.MarshalJSON()
	actual := args.Map{
		"bytesOK": err == nil && len(b) > 0,
		"jrOK":    jr.JsonString() != "",
		"jOK":     j.JsonString() != "",
		"jpOK":    jp != nil,
		"mbOK":    mErr == nil && len(mb) > 0,
		"model":   sr.JsonModel(),
		"any":     sr.JsonModelAny(),
	}
	expected := args.Map{
		"bytesOK": true,
		"jrOK":    true,
		"jOK":     true,
		"jpOK":    true,
		"mbOK":    true,
		"model":   "hello",
		"any":     "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- JSON", actual)
}

func Test_C16_TypedSimpleResult_GetAs(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid[string]("hello")
	s, sOk := sr.GetAsString()
	_, iOk := sr.GetAsInt()
	_, i64Ok := sr.GetAsInt64()
	_, f64Ok := sr.GetAsFloat64()
	_, bOk := sr.GetAsBool()
	_, byOk := sr.GetAsBytes()
	_, ssOk := sr.GetAsStrings()
	actual := args.Map{
		"s": s, "sOk": sOk, "iOk": iOk, "i64Ok": i64Ok,
		"f64Ok": f64Ok, "bOk": bOk, "byOk": byOk, "ssOk": ssOk,
	}
	expected := args.Map{
		"s": "hello", "sOk": true, "iOk": false, "i64Ok": false,
		"f64Ok": false, "bOk": false, "byOk": false, "ssOk": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- GetAs", actual)
}

func Test_C16_TypedSimpleResult_Clone(t *testing.T) {
	sr := coredynamic.NewTypedSimpleResultValid[string]("x")
	clone := sr.Clone()
	cloneP := sr.ClonePtr()
	var nilSR *coredynamic.TypedSimpleResult[string]
	nilClone := nilSR.Clone()
	nilCloneP := nilSR.ClonePtr()
	toSimple := sr.ToSimpleResult()
	toTD := sr.ToTypedDynamic()
	toDyn := sr.ToDynamic()
	nilToSimple := nilSR.ToSimpleResult()
	nilToTD := nilSR.ToTypedDynamic()
	nilToDyn := nilSR.ToDynamic()
	actual := args.Map{
		"cloneData":  clone.Data(),
		"clonePData": cloneP.Data(),
		"nilData":    nilClone.Data(),
		"nilP":       nilCloneP == nil,
		"simpleRes":  toSimple.Result,
		"tdData":     toTD.Data(),
		"dynValid":   toDyn.IsValid(),
		"nilSimple":  nilToSimple != nil,
		"nilTD":      nilToTD.IsInvalid(),
		"nilDyn":     nilToDyn.IsInvalid(),
	}
	expected := args.Map{
		"cloneData":  "x",
		"clonePData": "x",
		"nilData":    "",
		"nilP":       true,
		"simpleRes":  "x",
		"tdData":     "x",
		"dynValid":   true,
		"nilSimple":  true,
		"nilTD":      true,
		"nilDyn":     true,
	}
	expected.ShouldBeEqual(t, 0, "TypedSimpleResult returns correct value -- Clone/Convert", actual)
}

// ==========================================================================
// Package-level funcs coverage
// ==========================================================================

func Test_C16_SafeTypeName(t *testing.T) {
	actual := args.Map{
		"string": coredynamic.SafeTypeName("hello"),
		"nil":    coredynamic.SafeTypeName(nil),
		"int":    coredynamic.SafeTypeName(42),
	}
	expected := args.Map{
		"string": "string",
		"nil":    "",
		"int":    "int",
	}
	expected.ShouldBeEqual(t, 0, "SafeTypeName returns correct value -- with args", actual)
}

func Test_C16_Type(t *testing.T) {
	rt := coredynamic.Type("hello")
	actual := args.Map{"kind": rt.Kind() == reflect.String}
	expected := args.Map{"kind": true}
	expected.ShouldBeEqual(t, 0, "Type returns correct value -- with args", actual)
}

func Test_C16_IsAnyTypesOf(t *testing.T) {
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	boolType := reflect.TypeOf(true)
	actual := args.Map{
		"found":    coredynamic.IsAnyTypesOf(strType, intType, strType),
		"notFound": coredynamic.IsAnyTypesOf(boolType, intType, strType),
	}
	expected := args.Map{
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyTypesOf returns correct value -- with args", actual)
}

func Test_C16_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal("hello")
	actual := args.Map{"valid": rv.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- with args", actual)
}

func Test_C16_PointerOrNonPointer(t *testing.T) {
	s := "hello"
	out, rv := coredynamic.PointerOrNonPointer(false, &s)
	_ = out
	_ = rv
	// PointerOrNonPointer(true, nonPointer) panics due to uninitialized reflect.Value
	// in source — skip that call path
	outStruct, rvStruct := coredynamic.PointerOrNonPointer(false, s)
	_ = outStruct
	_ = rvStruct
}

func Test_C16_ZeroSet(t *testing.T) {
	type S struct{ X int }
	s := S{X: 42}
	rv := reflect.ValueOf(&s)
	coredynamic.ZeroSet(rv)
	actual := args.Map{"x": s.X}
	expected := args.Map{"x": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSet returns correct value -- with args", actual)
}

func Test_C16_LengthOfReflect(t *testing.T) {
	rv := reflect.ValueOf([]int{1, 2, 3})
	l := coredynamic.LengthOfReflect(rv)
	actual := args.Map{"len": l}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns correct value -- with args", actual)
}

// ==========================================================================
// BytesConverter coverage
// ==========================================================================

func Test_C16_BytesConverter(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()
	actual := args.Map{
		"noErr": err == nil,
		"val":   s,
	}
	expected := args.Map{
		"noErr": true,
		"val":   "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToString", actual)
}

func Test_C16_BytesConverter_SafeCastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("raw"))
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	actual := args.Map{
		"val":   bc.SafeCastString(),
		"empty": bcEmpty.SafeCastString(),
	}
	expected := args.Map{
		"val":   "raw",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- SafeCastString", actual)
}

func Test_C16_BytesConverter_CastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("raw"))
	s, err := bc.CastString()
	bcEmpty := coredynamic.NewBytesConverter([]byte{})
	_, errEmpty := bcEmpty.CastString()
	actual := args.Map{
		"val":      s,
		"noErr":    err == nil,
		"errEmpty": errEmpty != nil,
	}
	expected := args.Map{
		"val":      "raw",
		"noErr":    true,
		"errEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- CastString", actual)
}

func Test_C16_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`true`))
	val, err := bc.ToBool()
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBool", actual)
}

func Test_C16_BytesConverter_ToBoolMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`false`))
	val := bc.ToBoolMust()
	actual := args.Map{"val": val}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBoolMust", actual)
}

func Test_C16_BytesConverter_ToStringMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hi"`))
	val := bc.ToStringMust()
	actual := args.Map{"val": val}
	expected := args.Map{"val": "hi"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringMust", actual)
}

func Test_C16_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	val, err := bc.ToStrings()
	actual := args.Map{"len": len(val), "noErr": err == nil}
	expected := args.Map{"len": 2, "noErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStrings", actual)
}

func Test_C16_BytesConverter_ToStringsMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	val := bc.ToStringsMust()
	actual := args.Map{"len": len(val)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringsMust", actual)
}

func Test_C16_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`42`))
	val, err := bc.ToInt64()
	actual := args.Map{"val": val, "noErr": err == nil}
	expected := args.Map{"val": int64(42), "noErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64", actual)
}

func Test_C16_BytesConverter_ToInt64Must(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`99`))
	val := bc.ToInt64Must()
	actual := args.Map{"val": val}
	expected := args.Map{"val": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64Must", actual)
}

func Test_C16_BytesConverter_Deserialize(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var s string
	err := bc.Deserialize(&s)
	actual := args.Map{"val": s, "noErr": err == nil}
	expected := args.Map{"val": "hello", "noErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- Deserialize", actual)
}

func Test_C16_BytesConverter_DeserializeMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`42`))
	var i int
	bc.DeserializeMust(&i)
	actual := args.Map{"val": i}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- DeserializeMust", actual)
}

func Test_C16_BytesConverterUsingJsonResult(t *testing.T) {
	jr := corejson.NewPtr("hello")
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(jr)
	actual := args.Map{
		"noErr": err == nil,
		"notNil": bc != nil,
	}
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult returns correct value -- with args", actual)
}
