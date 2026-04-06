package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — constructors
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_BytesConverter_NewBytesConverter(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	actual := args.Map{"notNil": bc != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "NewBytesConverter returns correct value -- with args", actual)
}

func Test_I17_BytesConverter_NewUsingJsonResult(t *testing.T) {
	jr := corejson.NewPtr("test")
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(jr)
	actual := args.Map{"noErr": err == nil, "notNil": bc != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult returns correct value -- with args", actual)
}

func Test_I17_BytesConverter_NewUsingJsonResult_Error(t *testing.T) {
	jr := &corejson.Result{} // empty/invalid
	bc, err := coredynamic.NewBytesConverterUsingJsonResult(jr)
	actual := args.Map{"hasErr": err != nil, "nil": bc == nil}
	expected := args.Map{"hasErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "NewBytesConverterUsingJsonResult returns error -- error", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — Deserialize
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_BytesConverter_Deserialize(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var target string
	err := bc.Deserialize(&target)
	actual := args.Map{"noErr": err == nil, "val": target}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- Deserialize", actual)
}

func Test_I17_BytesConverter_DeserializeMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`42`))
	var target int
	bc.DeserializeMust(&target)
	actual := args.Map{"val": target}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- DeserializeMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — ToBool
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_BytesConverter_ToBool(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`true`))
	val, err := bc.ToBool()
	actual := args.Map{"noErr": err == nil, "val": val}
	expected := args.Map{"noErr": true, "val": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBool", actual)
}

func Test_I17_BytesConverter_ToBoolMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`false`))
	val := bc.ToBoolMust()
	actual := args.Map{"val": val}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBoolMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — String methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_BytesConverter_SafeCastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`raw text`))
	actual := args.Map{"val": bc.SafeCastString()}
	expected := args.Map{"val": "raw text"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- SafeCastString", actual)
}

func Test_I17_BytesConverter_SafeCastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	actual := args.Map{"val": bc.SafeCastString()}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns empty -- SafeCastString empty", actual)
}

func Test_I17_BytesConverter_CastString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`raw text`))
	val, err := bc.CastString()
	actual := args.Map{"noErr": err == nil, "val": val}
	expected := args.Map{"noErr": true, "val": "raw text"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- CastString", actual)
}

func Test_I17_BytesConverter_CastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte{})
	_, err := bc.CastString()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns empty -- CastString empty", actual)
}

func Test_I17_BytesConverter_ToString(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	val, err := bc.ToString()
	actual := args.Map{"noErr": err == nil, "val": val}
	expected := args.Map{"noErr": true, "val": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToString", actual)
}

func Test_I17_BytesConverter_ToStringMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"world"`))
	val := bc.ToStringMust()
	actual := args.Map{"val": val}
	expected := args.Map{"val": "world"}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringMust", actual)
}

func Test_I17_BytesConverter_ToStrings(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b","c"]`))
	val, err := bc.ToStrings()
	actual := args.Map{"noErr": err == nil, "len": len(val)}
	expected := args.Map{"noErr": true, "len": 3}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStrings", actual)
}

func Test_I17_BytesConverter_ToStringsMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x","y"]`))
	val := bc.ToStringsMust()
	actual := args.Map{"len": len(val)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToStringsMust", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — numeric
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_BytesConverter_ToInt64(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`999`))
	val, err := bc.ToInt64()
	actual := args.Map{"noErr": err == nil, "val": val}
	expected := args.Map{"noErr": true, "val": int64(999)}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64", actual)
}

func Test_I17_BytesConverter_ToInt64Must(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`123`))
	val := bc.ToInt64Must()
	actual := args.Map{"val": val}
	expected := args.Map{"val": int64(123)}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToInt64Must", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BytesConverter — complex type deserialization
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_BytesConverter_ToHashmap(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":"1","b":"2"}`))
	hm, err := bc.ToHashmap()
	actual := args.Map{"noErr": err == nil, "notNil": hm != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToHashmap", actual)
}

func Test_I17_BytesConverter_ToHashmap_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	hm, err := bc.ToHashmap()
	actual := args.Map{"hasErr": err != nil, "nil": hm == nil}
	expected := args.Map{"hasErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToHashmap invalid", actual)
}

func Test_I17_BytesConverter_ToHashmapMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"x":"y"}`))
	hm := bc.ToHashmapMust()
	actual := args.Map{"notNil": hm != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToHashmapMust", actual)
}

func Test_I17_BytesConverter_ToHashset(t *testing.T) {
	// Hashset internal is map[string]bool, so JSON must be object not array
	bc := coredynamic.NewBytesConverter([]byte(`{"a":true,"b":true}`))
	hs, err := bc.ToHashset()
	actual := args.Map{"noErr": err == nil, "notNil": hs != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToHashset", actual)
}

func Test_I17_BytesConverter_ToHashset_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	hs, err := bc.ToHashset()
	actual := args.Map{"hasErr": err != nil, "nil": hs == nil}
	expected := args.Map{"hasErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToHashset invalid", actual)
}

func Test_I17_BytesConverter_ToHashsetMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"x":true}`))
	hs := bc.ToHashsetMust()
	actual := args.Map{"notNil": hs != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToHashsetMust", actual)
}

func Test_I17_BytesConverter_ToCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	c, err := bc.ToCollection()
	actual := args.Map{"noErr": err == nil, "notNil": c != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToCollection", actual)
}

func Test_I17_BytesConverter_ToCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	c, err := bc.ToCollection()
	actual := args.Map{"hasErr": err != nil, "nil": c == nil}
	expected := args.Map{"hasErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToCollection invalid", actual)
}

func Test_I17_BytesConverter_ToCollectionMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	c := bc.ToCollectionMust()
	actual := args.Map{"notNil": c != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToCollectionMust", actual)
}

func Test_I17_BytesConverter_ToSimpleSlice(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	ss, err := bc.ToSimpleSlice()
	actual := args.Map{"noErr": err == nil, "notNil": ss != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToSimpleSlice", actual)
}

func Test_I17_BytesConverter_ToSimpleSlice_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	ss, err := bc.ToSimpleSlice()
	actual := args.Map{"hasErr": err != nil, "nil": ss == nil}
	expected := args.Map{"hasErr": true, "nil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToSimpleSlice invalid", actual)
}

func Test_I17_BytesConverter_ToSimpleSliceMust(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	ss := bc.ToSimpleSliceMust()
	actual := args.Map{"notNil": ss != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToSimpleSliceMust", actual)
}

func Test_I17_BytesConverter_ToKeyValCollection(t *testing.T) {
	kvc := coredynamic.EmptyKeyValCollection()
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	bytes, _ := json.Marshal(kvc)
	bc := coredynamic.NewBytesConverter(bytes)
	result, err := bc.ToKeyValCollection()
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToKeyValCollection", actual)
}

func Test_I17_BytesConverter_ToKeyValCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	_, err := bc.ToKeyValCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToKeyValCollection invalid", actual)
}

func Test_I17_BytesConverter_ToAnyCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	ac, err := bc.ToAnyCollection()
	actual := args.Map{"noErr": err == nil, "notNil": ac != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToAnyCollection", actual)
}

func Test_I17_BytesConverter_ToAnyCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	_, err := bc.ToAnyCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToAnyCollection invalid", actual)
}

func Test_I17_BytesConverter_ToMapAnyItems(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":1,"b":2}`))
	m, err := bc.ToMapAnyItems()
	actual := args.Map{"noErr": err == nil, "notNil": m != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToMapAnyItems", actual)
}

func Test_I17_BytesConverter_ToMapAnyItems_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	_, err := bc.ToMapAnyItems()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToMapAnyItems invalid", actual)
}

func Test_I17_BytesConverter_ToDynamicCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"Items":[]}`))
	result, err := bc.ToDynamicCollection()
	actual := args.Map{"noErr": err == nil, "notNil": result != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToDynamicCollection", actual)
}

func Test_I17_BytesConverter_ToDynamicCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	_, err := bc.ToDynamicCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToDynamicCollection invalid", actual)
}

func Test_I17_BytesConverter_ToJsonResultCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"JsonResultsCollection":[]}`))
	rc, err := bc.ToJsonResultCollection()
	actual := args.Map{"noErr": err == nil, "notNil": rc != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToJsonResultCollection", actual)
}

func Test_I17_BytesConverter_ToJsonResultCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	_, err := bc.ToJsonResultCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToJsonResultCollection invalid", actual)
}

func Test_I17_BytesConverter_ToJsonMapResults(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`{"a":1}`))
	mr, err := bc.ToJsonMapResults()
	actual := args.Map{"noErr": err == nil, "notNil": mr != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToJsonMapResults", actual)
}

func Test_I17_BytesConverter_ToJsonMapResults_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`not json`))
	_, err := bc.ToJsonMapResults()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns error -- ToJsonMapResults invalid", actual)
}

func Test_I17_BytesConverter_ToBytesCollection(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["YQ==","Yg=="]`))
	bColl, err := bc.ToBytesCollection()
	actual := args.Map{"noErr": err == nil, "notNil": bColl != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- ToBytesCollection", actual)
}
func Test_I17_SimpleResult_Clone(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("hello")
	cloned := sr.Clone()
	actual := args.Map{"valid": cloned.IsValid(), "msg": cloned.Message}
	expected := args.Map{"valid": true, "msg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- Clone", actual)
}

func Test_I17_SimpleResult_Clone_Nil(t *testing.T) {
	var sr *coredynamic.SimpleResult
	cloned := sr.Clone()
	actual := args.Map{"valid": cloned.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- Clone nil", actual)
}
	sr := coredynamic.InvalidSimpleResult("some error")
	err := sr.InvalidError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- InvalidError with message", actual)
}

func Test_I17_SimpleResult_InvalidError_Cached(t *testing.T) {
	sr := coredynamic.InvalidSimpleResult("cached error")
	err1 := sr.InvalidError()
	err2 := sr.InvalidError()
	actual := args.Map{"same": err1 == err2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- InvalidError cached", actual)
}

func Test_I17_SimpleResult_InvalidError_Nil(t *testing.T) {
	var sr *coredynamic.SimpleResult
	actual := args.Map{"nil": sr.InvalidError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- InvalidError nil", actual)
}

func Test_I17_SimpleResult_InvalidError_EmptyMessage(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("ok")
	actual := args.Map{"nil": sr.InvalidError() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns empty -- InvalidError empty msg", actual)
}

func Test_I17_SimpleResult_GetErrorOnTypeMismatch_Match(t *testing.T) {
	sr := coredynamic.NewSimpleResultValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- GetErrorOnTypeMismatch match", actual)
}

func Test_I17_SimpleResult_GetErrorOnTypeMismatch_Mismatch_ExcludeMsg(t *testing.T) {
	sr := coredynamic.NewSimpleResult("hello", true, "msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- GetErrorOnTypeMismatch exclude msg", actual)
}

func Test_I17_SimpleResult_GetErrorOnTypeMismatch_Mismatch_IncludeMsg(t *testing.T) {
	sr := coredynamic.NewSimpleResult("hello", true, "detail msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- GetErrorOnTypeMismatch include msg", actual)
}

func Test_I17_SimpleResult_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var sr *coredynamic.SimpleResult
	actual := args.Map{"nil": sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false) == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns nil -- GetErrorOnTypeMismatch nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SimpleRequest — nil receiver, cached error, type mismatch include msg
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_SimpleRequest_NilReceiver_Message(t *testing.T) {
	var sr *coredynamic.SimpleRequest
	actual := args.Map{"msg": sr.Message()}
	expected := args.Map{"msg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- nil Message", actual)
}

func Test_I17_SimpleRequest_NilReceiver_Request(t *testing.T) {
	var sr *coredynamic.SimpleRequest
	actual := args.Map{"nil": sr.Request() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- nil Request", actual)
}

func Test_I17_SimpleRequest_NilReceiver_Value(t *testing.T) {
	var sr *coredynamic.SimpleRequest
	actual := args.Map{"nil": sr.Value() == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns nil -- nil Value", actual)
}

func Test_I17_SimpleRequest_InvalidError_Cached(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequest("cached")
	err1 := sr.InvalidError()
	err2 := sr.InvalidError()
	actual := args.Map{"same": err1 == err2, "hasErr": err1 != nil}
	expected := args.Map{"same": true, "hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- InvalidError cached", actual)
}

func Test_I17_SimpleRequest_GetErrorOnTypeMismatch_IncludeMsg(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("hello", true, "detail")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- GetErrorOnTypeMismatch include msg", actual)
}

func Test_I17_SimpleRequest_GetErrorOnTypeMismatch_ExcludeMsg(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("hello", true, "detail")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- GetErrorOnTypeMismatch exclude msg", actual)
}

func Test_I17_SimpleRequest_IsPointer_Cached(t *testing.T) {
	x := 42
	sr := coredynamic.NewSimpleRequestValid(&x)
	p1 := sr.IsPointer()
	p2 := sr.IsPointer() // cached
	actual := args.Map{"p1": p1, "p2": p2}
	expected := args.Map{"p1": true, "p2": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- IsPointer cached", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MapAsKeyValSlice
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_MapAsKeyValSlice_Success(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	rv := reflect.ValueOf(m)
	kvc, err := coredynamic.MapAsKeyValSlice(rv)
	actual := args.Map{"noErr": err == nil, "len": kvc.Length()}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns correct value -- success", actual)
}

func Test_I17_MapAsKeyValSlice_NotMap(t *testing.T) {
	rv := reflect.ValueOf("not a map")
	_, err := coredynamic.MapAsKeyValSlice(rv)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice returns correct value -- not map", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr / MustBeAcceptedTypes
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_NotAcceptedTypesErr_Match(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(""), reflect.TypeOf(0))
	actual := args.Map{"nil": err == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns error -- match", actual)
}

func Test_I17_NotAcceptedTypesErr_NoMatch(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(0), reflect.TypeOf(true))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr returns empty -- no match", actual)
}

func Test_I17_MustBeAcceptedTypes_Success(t *testing.T) {
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(""))
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes returns correct value -- success", actual)
}

func Test_I17_MustBeAcceptedTypes_Panic(t *testing.T) {
	defer func() {
		r := recover()
		actual := args.Map{"panicked": r != nil}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "MustBeAcceptedTypes panics -- panic", actual)
	}()
	coredynamic.MustBeAcceptedTypes("hello", reflect.TypeOf(0))
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToReflectVal / ReflectInterfaceVal
// ══════════════════════════════════════════════════════════════════════════════

func Test_I17_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal(42)
	actual := args.Map{"valid": rv.IsValid(), "val": rv.Interface()}
	expected := args.Map{"valid": true, "val": 42}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns correct value -- with args", actual)
}

func Test_I17_ReflectInterfaceVal_NonPointer(t *testing.T) {
	val := coredynamic.ReflectInterfaceVal(42)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns non-empty -- non-pointer", actual)
}

func Test_I17_ReflectInterfaceVal_Pointer(t *testing.T) {
	x := 42
	val := coredynamic.ReflectInterfaceVal(&x)
	actual := args.Map{"val": val}
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- pointer", actual)
}
