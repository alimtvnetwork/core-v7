package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — value accessor methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyVal_KeyDynamic_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "name", Value: 42}
	d := kv.KeyDynamic()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"val": d.ValueString(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "name",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyDynamic", actual)
}

func Test_KeyVal_ValueDynamic_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "hello"}
	d := kv.ValueDynamic()

	// Act
	actual := args.Map{
		"valid": d.IsValid(),
		"val": d.ValueString(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueDynamic", actual)
}

func Test_KeyVal_KeyDynamicPtr(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyDynamicPtr", actual)
}

func Test_KeyVal_KeyDynamicPtr_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"nil": kv.KeyDynamicPtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyDynamicPtr nil", actual)
}

func Test_KeyVal_ValueDynamicPtr(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueDynamicPtr", actual)
}

func Test_KeyVal_ValueDynamicPtr_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"nil": kv.ValueDynamicPtr() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueDynamicPtr nil", actual)
}

func Test_KeyVal_IsKeyNull_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: nil, Value: "v"}

	// Act
	actual := args.Map{"isNull": kv.IsKeyNull()}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsKeyNull", actual)
}

func Test_KeyVal_IsKeyNull_NotNull(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"isNull": kv.IsKeyNull()}

	// Assert
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsKeyNull not null", actual)
}

func Test_KeyVal_IsValueNull_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: nil}

	// Act
	actual := args.Map{"isNull": kv.IsValueNull()}

	// Assert
	expected := args.Map{"isNull": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsValueNull", actual)
}

func Test_KeyVal_IsValueNull_NotNull(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"isNull": kv.IsValueNull()}

	// Assert
	expected := args.Map{"isNull": false}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- IsValueNull not null", actual)
}

func Test_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "", Value: "v"}

	// Act
	actual := args.Map{"empty": kv.IsKeyNullOrEmptyString()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns empty -- IsKeyNullOrEmptyString", actual)
}

func Test_KeyVal_String(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "name", Value: 42}
	s := kv.String()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- String", actual)
}

func Test_KeyVal_String_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"empty": kv.String() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- String nil", actual)
}

func Test_KeyVal_ValueReflectValue(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()

	// Act
	actual := args.Map{
		"valid": rv.IsValid(),
		"val": rv.Interface(),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueReflectValue", actual)
}

func Test_KeyVal_ValueInt(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"val": kv.ValueInt()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt", actual)
}

func Test_KeyVal_ValueInt_NotInt(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}

	// Act
	actual := args.Map{"val": kv.ValueInt()}

	// Assert
	expected := args.Map{"val": -1}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt not int", actual)
}

func Test_KeyVal_ValueUInt(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: uint(7)}

	// Act
	actual := args.Map{"val": kv.ValueUInt()}

	// Assert
	expected := args.Map{"val": uint(7)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueUInt", actual)
}

func Test_KeyVal_ValueUInt_NotUint(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}

	// Act
	actual := args.Map{"val": kv.ValueUInt()}

	// Assert
	expected := args.Map{"val": uint(0)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueUInt not uint", actual)
}

func Test_KeyVal_ValueStrings(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: []string{"a", "b"}}

	// Act
	actual := args.Map{"len": len(kv.ValueStrings())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyVal returns non-empty -- ValueStrings", actual)
}

func Test_KeyVal_ValueStrings_NotSlice(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"nil": kv.ValueStrings() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns non-empty -- ValueStrings not slice", actual)
}

func Test_KeyVal_ValueBool(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: true}

	// Act
	actual := args.Map{"val": kv.ValueBool()}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueBool", actual)
}

func Test_KeyVal_ValueBool_NotBool(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"val": kv.ValueBool()}

	// Assert
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueBool not bool", actual)
}

func Test_KeyVal_ValueInt64(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: int64(999)}

	// Act
	actual := args.Map{"val": kv.ValueInt64()}

	// Assert
	expected := args.Map{"val": int64(999)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt64", actual)
}

func Test_KeyVal_ValueInt64_NotInt64(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "str"}

	// Act
	actual := args.Map{"val": kv.ValueInt64()}

	// Assert
	expected := args.Map{"val": int64(-1)}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueInt64 not int64", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — nil receiver error methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyVal_ValueNullErr_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"hasErr": kv.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueNullErr nil", actual)
}

func Test_KeyVal_ValueNullErr_NullValue_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: nil}

	// Act
	actual := args.Map{"hasErr": kv.ValueNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- ValueNullErr null value", actual)
}

func Test_KeyVal_ValueNullErr_Valid_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"noErr": kv.ValueNullErr() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- ValueNullErr valid", actual)
}

func Test_KeyVal_KeyNullErr_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"hasErr": kv.KeyNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyNullErr nil", actual)
}

func Test_KeyVal_KeyNullErr_NullKey_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: nil, Value: 42}

	// Act
	actual := args.Map{"hasErr": kv.KeyNullErr() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- KeyNullErr null key", actual)
}

func Test_KeyVal_KeyNullErr_Valid_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"noErr": kv.KeyNullErr() == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns error -- KeyNullErr valid", actual)
}

func Test_KeyVal_KeyString(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "mykey", Value: 1}

	// Act
	actual := args.Map{"notEmpty": kv.KeyString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyString", actual)
}

func Test_KeyVal_KeyString_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"empty": kv.KeyString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyString nil", actual)
}

func Test_KeyVal_ValueString(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "hello"}

	// Act
	actual := args.Map{"notEmpty": kv.ValueString() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns non-empty -- ValueString", actual)
}

func Test_KeyVal_ValueString_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"empty": kv.ValueString() == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueString nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — Reflect set methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyVal_ReflectSetKey(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "mykey", Value: 42}
	var target string
	err := kv.ReflectSetKey(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "mykey",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ReflectSetKey", actual)
}

func Test_KeyVal_ReflectSetKey_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"hasErr": kv.ReflectSetKey(nil) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ReflectSetKey nil", actual)
}

func Test_KeyVal_KeyReflectSet(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "abc", Value: 1}
	var target string
	err := kv.KeyReflectSet(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "abc",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- KeyReflectSet", actual)
}

func Test_KeyVal_KeyReflectSet_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"hasErr": kv.KeyReflectSet(nil) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- KeyReflectSet nil", actual)
}

func Test_KeyVal_ValueReflectSet(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	var target int
	err := kv.ValueReflectSet(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ValueReflectSet", actual)
}

func Test_KeyVal_ValueReflectSet_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"hasErr": kv.ValueReflectSet(nil) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ValueReflectSet nil", actual)
}

func Test_KeyVal_ReflectSetTo(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "world"}
	var target string
	err := kv.ReflectSetTo(&target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "world",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ReflectSetTo", actual)
}

func Test_KeyVal_ReflectSetTo_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"hasErr": kv.ReflectSetTo(nil) != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- ReflectSetTo nil", actual)
}

func Test_KeyVal_ReflectSetToMust_Success(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	var target int
	kv.ReflectSetToMust(&target)

	// Act
	actual := args.Map{"val": target}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ReflectSetToMust success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyVal — JSON methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyVal_JsonModel(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"notNil": kv.JsonModel() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonModel", actual)
}

func Test_KeyVal_JsonModelAny(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"notNil": kv.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonModelAny", actual)
}

func Test_KeyVal_Json(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.Json()

	// Act
	actual := args.Map{"noErr": !jr.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- Json", actual)
}

func Test_KeyVal_JsonPtr(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	jr := kv.JsonPtr()

	// Act
	actual := args.Map{
		"notNil": jr != nil,
		"noErr": !jr.HasError(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonPtr", actual)
}

func Test_KeyVal_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{}
	original := coredynamic.KeyVal{Key: "pk", Value: "pv"}
	jr := corejson.NewPtr(original)
	result, err := kv.ParseInjectUsingJson(jr)

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
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ParseInjectUsingJson", actual)
}

func Test_KeyVal_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{}
	original := coredynamic.KeyVal{Key: "pk", Value: "pv"}
	jr := corejson.NewPtr(original)
	result := kv.ParseInjectUsingJsonMust(jr)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- ParseInjectUsingJsonMust", actual)
}

func Test_KeyVal_JsonParseSelfInject(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{}
	original := coredynamic.KeyVal{Key: "pk", Value: "pv"}
	jr := corejson.NewPtr(original)
	err := kv.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- JsonParseSelfInject", actual)
}

func Test_KeyVal_Serialize(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	bytes, err := kv.Serialize()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": len(bytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- Serialize", actual)
}

func Test_KeyVal_CastKeyVal_Nil_FromKeyValKeyDynamicIter(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	err := kv.CastKeyVal(nil, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal returns nil -- CastKeyVal nil", actual)
}

func Test_KeyVal_CastKeyVal_Success(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "mykey", Value: 42}
	var k string
	var v int
	err := kv.CastKeyVal(&k, &v)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"k": k,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"k": "mykey",
		"v": 42,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal returns correct value -- CastKeyVal success", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// New.Collection creators — Generic, String, Int, Int64, Byte, Any
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewCollection_String_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Empty returns empty -- with args", actual)
}

func Test_NewCollection_String_Cap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Cap(10)

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Cap returns correct value -- with args", actual)
}

func Test_NewCollection_String_From(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.From([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.From returns correct value -- with args", actual)
}

func Test_NewCollection_String_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Clone([]string{"x", "y"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Clone returns correct value -- with args", actual)
}

func Test_NewCollection_String_Items(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Items("a", "b")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Items returns correct value -- with args", actual)
}

func Test_NewCollection_String_Create(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Create([]string{"x"})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.Create returns correct value -- with args", actual)
}

func Test_NewCollection_String_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.LenCap(3, 10)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.String.LenCap returns correct value -- with args", actual)
}

func Test_NewCollection_Int_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Empty returns empty -- with args", actual)
}

func Test_NewCollection_Int_Cap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.Cap(5)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Cap returns correct value -- with args", actual)
}

func Test_NewCollection_Int_From(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.From([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.From returns correct value -- with args", actual)
}

func Test_NewCollection_Int_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.LenCap(2, 8)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.LenCap returns correct value -- with args", actual)
}

func Test_NewCollection_Int64_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int64.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.Empty returns empty -- with args", actual)
}

func Test_NewCollection_Int64_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int64.LenCap(1, 4)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.LenCap returns correct value -- with args", actual)
}

func Test_NewCollection_Byte_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Byte.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.Empty returns empty -- with args", actual)
}

func Test_NewCollection_Byte_LenCap(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Byte.LenCap(5, 10)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.LenCap returns correct value -- with args", actual)
}

func Test_NewCollection_Any_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Any.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.Empty returns empty -- with args", actual)
}

func Test_NewCollection_Any_From(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Any.From([]any{"a", 1, true})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any.From returns correct value -- with args", actual)
}

func Test_NewCollection_Bool_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Bool.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Bool.Empty returns empty -- with args", actual)
}

func Test_NewCollection_Float64_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Float64.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float64.Empty returns empty -- with args", actual)
}

func Test_NewCollection_Float32_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Float32.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.Float32.Empty returns empty -- with args", actual)
}

func Test_NewCollection_AnyMap_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.AnyMap.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.AnyMap.Empty returns empty -- with args", actual)
}

func Test_NewCollection_StringMap_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.StringMap.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.StringMap.Empty returns empty -- with args", actual)
}

func Test_NewCollection_IntMap_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.IntMap.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.IntMap.Empty returns empty -- with args", actual)
}

func Test_NewCollection_ByteSlice_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.ByteSlice.Empty()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.Empty returns empty -- with args", actual)
}

func Test_NewCollection_ByteSlice_From(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.ByteSlice.From([][]byte{{1, 2}, {3, 4}})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.ByteSlice.From returns correct value -- with args", actual)
}

func Test_NewCollection_Int_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.Clone([]int{10, 20, 30})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Clone returns correct value -- with args", actual)
}

func Test_NewCollection_Int_Items(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.Items(1, 2, 3, 4)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.Items returns correct value -- with args", actual)
}
