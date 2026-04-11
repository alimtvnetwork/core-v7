package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// KeyVal — Dynamic accessors
// =============================================================================

func Test_KeyVal_KeyDynamic_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "mykey", Value: 42}
	d := kv.KeyDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamic", actual)
}

func Test_KeyVal_ValueDynamic_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.ValueDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamic", actual)
}

func Test_KeyVal_KeyDynamicPtr_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"isNil": kv.KeyDynamicPtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamicPtr nil", actual)
}

func Test_KeyVal_KeyDynamicPtr_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	d := kv.KeyDynamicPtr()

	// Act
	actual := args.Map{
		"notNil": d != nil,
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyDynamicPtr valid", actual)
}

func Test_KeyVal_ValueDynamicPtr_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"isNil": kv.ValueDynamicPtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamicPtr nil", actual)
}

func Test_KeyVal_ValueDynamicPtr_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 99}
	d := kv.ValueDynamicPtr()

	// Act
	actual := args.Map{
		"notNil": d != nil,
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueDynamicPtr valid", actual)
}

// =============================================================================
// KeyVal — Null checks
// =============================================================================

func Test_KeyVal_IsKeyNull_True(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}

	// Act
	actual := args.Map{"r": kv.IsKeyNull()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNull true", actual)
}

func Test_KeyVal_IsKeyNull_False(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"r": kv.IsKeyNull()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNull false", actual)
}

func Test_KeyVal_IsKeyNullOrEmptyString_Null(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}

	// Act
	actual := args.Map{"r": kv.IsKeyNullOrEmptyString()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString null", actual)
}

func Test_KeyVal_IsKeyNullOrEmptyString_Empty(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "", Value: "v"}

	// Act
	actual := args.Map{"r": kv.IsKeyNullOrEmptyString()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString empty", actual)
}

func Test_KeyVal_IsKeyNullOrEmptyString_NonEmpty(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"r": kv.IsKeyNullOrEmptyString()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal IsKeyNullOrEmptyString non-empty", actual)
}

func Test_KeyVal_IsValueNull_True(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: nil}

	// Act
	actual := args.Map{"r": kv.IsValueNull()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal IsValueNull true", actual)
}

func Test_KeyVal_IsValueNull_False(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"r": kv.IsValueNull()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal IsValueNull false", actual)
}

// =============================================================================
// KeyVal — String, KeyString, ValueString
// =============================================================================

func Test_KeyVal_String_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"r": kv.String()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal String nil", actual)
}

func Test_KeyVal_String_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	s := kv.String()

	// Act
	actual := args.Map{"nonEmpty": s != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal String valid", actual)
}

func Test_KeyVal_KeyString_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"r": kv.KeyString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString nil receiver", actual)
}

func Test_KeyVal_KeyString_NilKey(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: nil, Value: "v"}

	// Act
	actual := args.Map{"r": kv.KeyString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString nil key", actual)
}

func Test_KeyVal_KeyString_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "myKey", Value: "v"}

	// Act
	actual := args.Map{"r": kv.KeyString()}

	// Assert
	expected := args.Map{"r": "myKey"}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyString valid", actual)
}

func Test_KeyVal_ValueString_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"r": kv.ValueString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString nil", actual)
}

func Test_KeyVal_ValueString_NilValue(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}

	// Act
	actual := args.Map{"r": kv.ValueString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString nil value", actual)
}

func Test_KeyVal_ValueString_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "hello"}

	// Act
	actual := args.Map{"r": kv.ValueString()}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueString valid", actual)
}

// =============================================================================
// KeyVal — Value typed accessors
// =============================================================================

func Test_KeyVal_ValueInt_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"r": kv.ValueInt()}

	// Assert
	expected := args.Map{"r": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt valid", actual)
}

func Test_KeyVal_ValueInt_Invalid(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "notint"}

	// Act
	actual := args.Map{"r": kv.ValueInt()}

	// Assert
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt invalid", actual)
}

func Test_KeyVal_ValueUInt_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: uint(10)}

	// Act
	actual := args.Map{"r": kv.ValueUInt()}

	// Assert
	expected := args.Map{"r": uint(10)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt valid", actual)
}

func Test_KeyVal_ValueUInt_Invalid(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "notuint"}

	// Act
	actual := args.Map{"r": kv.ValueUInt()}

	// Assert
	expected := args.Map{"r": uint(0)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueUInt invalid", actual)
}

func Test_KeyVal_ValueBool_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: true}

	// Act
	actual := args.Map{"r": kv.ValueBool()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool valid", actual)
}

func Test_KeyVal_ValueBool_Invalid(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "notbool"}

	// Act
	actual := args.Map{"r": kv.ValueBool()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueBool invalid", actual)
}

func Test_KeyVal_ValueInt64_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: int64(999)}

	// Act
	actual := args.Map{"r": kv.ValueInt64()}

	// Assert
	expected := args.Map{"r": int64(999)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64 valid", actual)
}

func Test_KeyVal_ValueInt64_Invalid(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "notint64"}

	// Act
	actual := args.Map{"r": kv.ValueInt64()}

	// Assert
	expected := args.Map{"r": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueInt64 invalid", actual)
}

func Test_KeyVal_ValueStrings_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}

	// Act
	actual := args.Map{"len": len(kv.ValueStrings())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings valid", actual)
}

func Test_KeyVal_ValueStrings_Invalid(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "notslice"}

	// Act
	actual := args.Map{"isNil": kv.ValueStrings() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueStrings invalid", actual)
}

// =============================================================================
// KeyVal — Null error methods
// =============================================================================

func Test_KeyVal_ValueNullErr_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	err := kv.ValueNullErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr nil", actual)
}

func Test_KeyVal_ValueNullErr_NullValue_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}
	err := kv.ValueNullErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr null value", actual)
}

func Test_KeyVal_ValueNullErr_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"noErr": kv.ValueNullErr() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueNullErr valid", actual)
}

func Test_KeyVal_KeyNullErr_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	err := kv.KeyNullErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr nil", actual)
}

func Test_KeyVal_KeyNullErr_NullKey_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: nil, Value: 42}
	err := kv.KeyNullErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr null key", actual)
}

func Test_KeyVal_KeyNullErr_Valid_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"noErr": kv.KeyNullErr() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyNullErr valid", actual)
}

// =============================================================================
// KeyVal — CastKeyVal
// =============================================================================

func Test_KeyVal_CastKeyVal_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	var k, v string
	err := kv.CastKeyVal(&k, &v)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal nil", actual)
}

func Test_KeyVal_CastKeyVal_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "mykey", Value: "myval"}
	var k, v string
	err := kv.CastKeyVal(&k, &v)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": v,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "myval",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal valid", actual)
}

// =============================================================================
// KeyVal — ReflectSet methods
// =============================================================================

func Test_KeyVal_ReflectSetKey_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	var k string
	err := kv.ReflectSetKey(&k)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetKey nil", actual)
}

func Test_KeyVal_KeyReflectSet_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	var k string
	err := kv.KeyReflectSet(&k)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal KeyReflectSet nil", actual)
}

func Test_KeyVal_ValueReflectSet_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	var v string
	err := kv.ValueReflectSet(&v)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectSet nil", actual)
}

func Test_KeyVal_ReflectSetTo_Nil_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	var v string
	err := kv.ReflectSetTo(&v)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetTo nil", actual)
}

func Test_KeyVal_ReflectSetTo_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "hello"}
	var v string
	err := kv.ReflectSetTo(&v)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": "hello",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetTo valid", actual)
}

func Test_KeyVal_ReflectSetToMust_Valid(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "world"}
	var v string
	kv.ReflectSetToMust(&v)

	// Act
	actual := args.Map{"v": v}

	// Assert
	expected := args.Map{"v": "world"}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSetToMust valid", actual)
}

func Test_KeyVal_ValueReflectValue_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()

	// Act
	actual := args.Map{"valid": rv.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectValue", actual)
}

// =============================================================================
// KeyVal — JSON methods
// =============================================================================

func Test_KeyVal_JsonModel_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"notNil": kv.JsonModel() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonModel", actual)
}

func Test_KeyVal_JsonModelAny_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"notNil": kv.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonModelAny", actual)
}

func Test_KeyVal_Json_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	r := kv.Json()

	// Act
	actual := args.Map{"hasBytes": r.HasAnyItem()}

	// Assert
	expected := args.Map{"hasBytes": true}
	expected.ShouldBeEqual(t, 0, "KeyVal Json", actual)
}

func Test_KeyVal_JsonPtr_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"notNil": kv.JsonPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonPtr", actual)
}

func Test_KeyVal_Serialize_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal Serialize", actual)
}

func Test_KeyVal_ParseInjectUsingJson_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &coredynamic.KeyVal{}
	result, err := kv2.ParseInjectUsingJson(&jr)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJson", actual)
}

func Test_KeyVal_ParseInjectUsingJsonMust_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &coredynamic.KeyVal{}
	result := kv2.ParseInjectUsingJsonMust(&jr)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ParseInjectUsingJsonMust", actual)
}

func Test_KeyVal_JsonParseSelfInject_FromKeyValKeyDynamicKeyV(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()
	kv2 := &coredynamic.KeyVal{}
	err := kv2.JsonParseSelfInject(&jr)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal JsonParseSelfInject", actual)
}
