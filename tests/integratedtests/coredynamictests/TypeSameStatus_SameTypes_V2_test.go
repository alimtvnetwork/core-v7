package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// TypeSameStatus
// ==========================================================================

func Test_TypeSameStatus_SameTypes(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", "world")

	// Act
	actual := args.Map{
		"isSame":       ts.IsSame,
		"leftNotNull":  !ts.IsLeftUnknownNull,
		"rightNotNull": !ts.IsRightUnknownNull,
		"leftPtr":      ts.IsLeftPointer,
		"rightPtr":     ts.IsRightPointer,
	}

	// Assert
	expected := args.Map{
		"isSame": true, "leftNotNull": true, "rightNotNull": true,
		"leftPtr": false, "rightPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus same types -- string string", actual)
}

func Test_TypeSameStatus_DiffTypes(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", 42)

	// Act
	actual := args.Map{"isSame": ts.IsSame}

	// Assert
	expected := args.Map{"isSame": false}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus diff types -- string int", actual)
}

func Test_TypeSameStatus_NilLeft(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, "hello")

	// Act
	actual := args.Map{
		"leftNull": ts.IsLeftUnknownNull,
		"rightNull": ts.IsRightUnknownNull,
	}

	// Assert
	expected := args.Map{
		"leftNull": true,
		"rightNull": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus nil left -- nil string", actual)
}

func Test_TypeSameStatus_Pointers(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, &s)

	// Act
	actual := args.Map{
		"leftPtr": ts.IsLeftPointer,
		"rightPtr": ts.IsRightPointer,
		"isSame": ts.IsSame,
	}

	// Assert
	expected := args.Map{
		"leftPtr": true,
		"rightPtr": true,
		"isSame": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus pointers -- same ptr type", actual)
}

// ==========================================================================
// TypeNotEqualErr
// ==========================================================================

func Test_TypeNotEqualErr_Same_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", "b")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr same types no error -- string string", actual)
}

func Test_TypeNotEqualErr_Diff(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr("a", 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypeNotEqualErr diff types has error -- string int", actual)
}

// ==========================================================================
// TypesIndexOf
// ==========================================================================

func Test_TypesIndexOf_Found_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)

	// Act
	actual := args.Map{"index": coredynamic.TypesIndexOf(intType, strType, intType)}

	// Assert
	expected := args.Map{"index": 1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf finds at index 1 -- int in [str,int]", actual)
}

func Test_TypesIndexOf_NotFound_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	boolType := reflect.TypeOf(true)

	// Act
	actual := args.Map{"index": coredynamic.TypesIndexOf(boolType, strType)}

	// Assert
	expected := args.Map{"index": -1}
	expected.ShouldBeEqual(t, 0, "TypesIndexOf returns -1 -- not found", actual)
}

// ==========================================================================
// NotAcceptedTypesErr
// ==========================================================================

func Test_NotAcceptedTypesErr_Match_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr match -- string", actual)
}

func Test_NotAcceptedTypesErr_NoMatch_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr("hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "NotAcceptedTypesErr no match -- string vs int", actual)
}

// ==========================================================================
// ReflectKindValidation
// ==========================================================================

func Test_ReflectKindValidation_Match_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.String, "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation match -- string", actual)
}

func Test_ReflectKindValidation_NoMatch(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.Int, "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectKindValidation no match -- string vs int", actual)
}

// ==========================================================================
// ReflectTypeValidation
// ==========================================================================

func Test_ReflectTypeValidation_Match_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation match -- string", actual)
}

func Test_ReflectTypeValidation_NoMatch(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(0), "hello")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation no match -- string vs int", actual)
}

func Test_ReflectTypeValidation_NilNotExpected_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(""), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation nil not expected -- nil", actual)
}

func Test_ReflectTypeValidation_NilAllowed_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectTypeValidation nil type mismatch -- nil allowed but mismatched", actual)
}

// ==========================================================================
// ReflectInterfaceVal
// ==========================================================================

func Test_ReflectInterfaceVal_Value_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	result := coredynamic.ReflectInterfaceVal("hello")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal value type -- string", actual)
}

func Test_ReflectInterfaceVal_Pointer_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	s := "hello"
	result := coredynamic.ReflectInterfaceVal(&s)

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal pointer -- derefs ptr", actual)
}

// ==========================================================================
// ZeroSetAny
// ==========================================================================

func Test_ZeroSetAny_Nil_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	// should not panic
	coredynamic.ZeroSetAny(nil)

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny nil no panic -- nil", actual)
}

func Test_ZeroSetAny_Pointer(t *testing.T) {
	// Arrange
	type s struct{ A int }
	val := &s{A: 42}
	coredynamic.ZeroSetAny(val)

	// Act
	actual := args.Map{"a": val.A}

	// Assert
	expected := args.Map{"a": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSetAny zeros struct -- pointer", actual)
}

// ==========================================================================
// ZeroSet
// ==========================================================================

func Test_ZeroSet_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	type s struct{ A int }
	val := &s{A: 42}
	coredynamic.ZeroSet(reflect.ValueOf(val))

	// Act
	actual := args.Map{"a": val.A}

	// Assert
	expected := args.Map{"a": 0}
	expected.ShouldBeEqual(t, 0, "ZeroSet zeros struct -- pointer rv", actual)
}

// ==========================================================================
// Type
// ==========================================================================

func Test_Type_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	rt := coredynamic.Type("hello")

	// Act
	actual := args.Map{"name": rt.String()}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "Type returns reflect.Type -- string", actual)
}

// ==========================================================================
// AnyToReflectVal
// ==========================================================================

func Test_AnyToReflectVal_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	rv := coredynamic.AnyToReflectVal(42)

	// Act
	actual := args.Map{
		"kind": rv.Kind() == reflect.Int,
		"val": int(rv.Int()),
	}

	// Assert
	expected := args.Map{
		"kind": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "AnyToReflectVal returns value -- int", actual)
}

// ==========================================================================
// KeyVal — comprehensive
// ==========================================================================

func Test_KeyVal_Basics(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "myKey", Value: "myVal"}

	// Act
	actual := args.Map{
		"keyString":   kv.KeyString(),
		"valueString": kv.ValueString(),
		"string":      kv.String() != "",
		"isKeyNull":   kv.IsKeyNull(),
		"isValNull":   kv.IsValueNull(),
	}

	// Assert
	expected := args.Map{
		"keyString": "myKey", "valueString": "myVal",
		"string": true, "isKeyNull": false, "isValNull": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal basics return expected -- string kv", actual)
}

func Test_KeyVal_TypedGetters(t *testing.T) {
	// Arrange
	kvInt := &coredynamic.KeyVal{Key: "k", Value: 42}
	kvUint := &coredynamic.KeyVal{Key: "k", Value: uint(10)}
	kvBool := &coredynamic.KeyVal{Key: "k", Value: true}
	kvI64 := &coredynamic.KeyVal{Key: "k", Value: int64(100)}
	kvStrings := &coredynamic.KeyVal{Key: "k", Value: []string{"a"}}

	// Act
	actual := args.Map{
		"int":     kvInt.ValueInt(),
		"uint":    int(kvUint.ValueUInt()),
		"bool":    kvBool.ValueBool(),
		"int64":   int(kvI64.ValueInt64()),
		"strings": len(kvStrings.ValueStrings()),
	}

	// Assert
	expected := args.Map{
		"int": 42, "uint": 10, "bool": true, "int64": 100, "strings": 1,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal typed getters -- various types", actual)
}

func Test_KeyVal_TypedGetters_Fail(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "notInt"}

	// Act
	actual := args.Map{
		"int":     kv.ValueInt(),
		"uint":    int(kv.ValueUInt()),
		"bool":    kv.ValueBool(),
		"int64":   int(kv.ValueInt64()),
		"strings": kv.ValueStrings() == nil,
	}

	// Assert
	expected := args.Map{
		"int": -1, "uint": 0, "bool": false, "int64": -1, "strings": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal typed getters fail -- wrong type", actual)
}

func Test_KeyVal_Dynamic(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	kd := kv.KeyDynamic()
	vd := kv.ValueDynamic()
	kdp := kv.KeyDynamicPtr()
	vdp := kv.ValueDynamicPtr()

	// Act
	actual := args.Map{
		"kd": kd.Data(), "vd": vd.Data(),
		"kdp": kdp.Data(), "vdp": vdp.Data(),
	}

	// Assert
	expected := args.Map{
		"kd": "k", "vd": "v", "kdp": "k", "vdp": "v",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal Dynamic methods -- string kv", actual)
}

func Test_KeyVal_NullErrors(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{
		"valueNullErr": kv.ValueNullErr() == nil,
		"keyNullErr":   kv.KeyNullErr() == nil,
	}

	// Assert
	expected := args.Map{
		"valueNullErr": true,
		"keyNullErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal null errors nil -- non-null kv", actual)
}

func Test_KeyVal_NullErrors_Nil(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{
		"valueNullErr": kv.ValueNullErr() != nil,
		"keyNullErr":   kv.KeyNullErr() != nil,
	}

	// Assert
	expected := args.Map{
		"valueNullErr": true,
		"keyNullErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal null errors return err -- nil receiver", actual)
}

func Test_KeyVal_NilStrings(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{
		"keyString":   kv.KeyString(),
		"valueString": kv.ValueString(),
	}

	// Assert
	expected := args.Map{
		"keyString": "",
		"valueString": "",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal nil receiver strings empty -- nil", actual)
}

func Test_KeyVal_ReflectSetMethods(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var keyTarget string
	var valTarget string
	keyErr := kv.KeyReflectSet(&keyTarget)
	valErr := kv.ValueReflectSet(&valTarget)

	// Act
	actual := args.Map{
		"keyErr": keyErr == nil, "valErr": valErr == nil,
		"keyTarget": keyTarget, "valTarget": valTarget,
	}

	// Assert
	expected := args.Map{
		"keyErr": true, "valErr": true,
		"keyTarget": "hello", "valTarget": "world",
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSet methods -- string kv", actual)
}

func Test_KeyVal_ReflectSetMethods_Nil(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{
		"keyErr":    kv.KeyReflectSet(nil) != nil,
		"valErr":    kv.ValueReflectSet(nil) != nil,
		"setToErr":  kv.ReflectSetTo(nil) != nil,
		"setKeyErr": kv.ReflectSetKey(nil) != nil,
	}

	// Assert
	expected := args.Map{
		"keyErr": true,
		"valErr": true,
		"setToErr": true,
		"setKeyErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal ReflectSet nil receiver returns errors -- nil", actual)
}

func Test_KeyVal_CastKeyVal_Nil_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal
	err := kv.CastKeyVal(nil, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "KeyVal CastKeyVal nil receiver -- nil", actual)
}

func Test_KeyVal_Json_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	jp := kv.JsonPtr()

	// Act
	actual := args.Map{
		"model":    kv.JsonModel() != nil,
		"modelAny": kv.JsonModelAny() != nil,
		"jNotNil":  j.Bytes != nil,
		"jpNotNil": jp != nil,
	}

	// Assert
	expected := args.Map{
		"model": true,
		"modelAny": true,
		"jNotNil": true,
		"jpNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal Json methods -- valid kv", actual)
}

func Test_KeyVal_Serialize_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	bytes, err := kv.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "KeyVal Serialize returns bytes -- valid", actual)
}

func Test_KeyVal_ValueReflectValue_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()

	// Act
	actual := args.Map{"kind": rv.Kind() == reflect.Int}

	// Assert
	expected := args.Map{"kind": true}
	expected.ShouldBeEqual(t, 0, "KeyVal ValueReflectValue kind -- int", actual)
}

// ==========================================================================
// SimpleRequest — coverage
// ==========================================================================

func Test_SimpleRequest_Constructors(t *testing.T) {
	// Arrange
	r1 := coredynamic.InvalidSimpleRequestNoMessage()
	r2 := coredynamic.InvalidSimpleRequest("error")
	r3 := coredynamic.NewSimpleRequest("data", true, "")
	r4 := coredynamic.NewSimpleRequestValid("data")

	// Act
	actual := args.Map{
		"r1Valid": r1.IsValid(), "r1Msg": r1.Message(),
		"r2Valid": r2.IsValid(), "r2Msg": r2.Message(),
		"r3Valid": r3.IsValid(), "r3Request": r3.Request(),
		"r4Valid": r4.IsValid(), "r4Value": r4.Value(),
	}

	// Assert
	expected := args.Map{
		"r1Valid": false, "r1Msg": "",
		"r2Valid": false, "r2Msg": "error",
		"r3Valid": true, "r3Request": "data",
		"r4Valid": true, "r4Value": "data",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest constructors -- all", actual)
}

func Test_SimpleRequest_InvalidError(t *testing.T) {
	// Arrange
	r1 := coredynamic.NewSimpleRequestValid("ok")
	r2 := coredynamic.InvalidSimpleRequest("err")

	// Act
	actual := args.Map{
		"r1Err": r1.InvalidError() == nil,
		"r2Err": r2.InvalidError() != nil,
	}

	// Assert
	expected := args.Map{
		"r1Err": true,
		"r2Err": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest InvalidError -- valid and invalid", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid("hello")
	matchErr := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	mismatchErr := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)

	// Act
	actual := args.Map{
		"match":    matchErr == nil,
		"mismatch": mismatchErr != nil,
	}

	// Assert
	expected := args.Map{
		"match": true,
		"mismatch": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest GetErrorOnTypeMismatch -- match and mismatch", actual)
}

func Test_SimpleRequest_IsPointer(t *testing.T) {
	// Arrange
	s := "hello"
	r := coredynamic.NewSimpleRequestValid(&s)

	// Act
	actual := args.Map{"isPtr": r.IsPointer()}

	// Assert
	expected := args.Map{"isPtr": true}
	expected.ShouldBeEqual(t, 0, "SimpleRequest IsPointer -- pointer value", actual)
}

// ==========================================================================
// MapAsKeyValSlice — error path
// ==========================================================================

func Test_MapAsKeyValSlice_NotMap_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf("notAMap")
	_, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice not map returns error -- string", actual)
}

func Test_MapAsKeyValSlice_Valid_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	kvc, err := coredynamic.MapAsKeyValSlice(rv)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"notNil": kvc != nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MapAsKeyValSlice valid map -- 2 entries", actual)
}

// ==========================================================================
// TypeStatus — pointer branches
// ==========================================================================

func Test_TypeStatus_PointerBranches(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "hello")

	// Act
	actual := args.Map{
		"isAnyPtr":         ts.IsAnyPointer(),
		"isBothPtr":        ts.IsBothPointer(),
		"sameRegardless":   ts.IsSameRegardlessPointer(),
		"nonPointerLeft":   ts.NonPointerLeft().String(),
		"nonPointerRight":  ts.NonPointerRight().String(),
	}

	// Assert
	expected := args.Map{
		"isAnyPtr": true, "isBothPtr": false,
		"sameRegardless": true,
		"nonPointerLeft": "string", "nonPointerRight": "string",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus pointer branches -- ptr vs value", actual)
}

func Test_TypeStatus_NullNames(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, nil)

	// Act
	actual := args.Map{
		"leftName":  ts.LeftName(),
		"rightName": ts.RightName(),
		"leftFull":  ts.LeftFullName(),
		"rightFull": ts.RightFullName(),
	}

	// Assert
	expected := args.Map{
		"leftName": "<nil>", "rightName": "<nil>",
		"leftFull": "<nil>", "rightFull": "<nil>",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus null names -- nil nil", actual)
}

func Test_TypeStatus_IsEqual_RightDiff(t *testing.T) {
	// Arrange
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	ts1 := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: strType}
	ts2 := &coredynamic.TypeStatus{IsSame: true, Left: strType, Right: intType}

	// Act
	actual := args.Map{"equal": ts1.IsEqual(ts2)}

	// Assert
	expected := args.Map{"equal": false}
	expected.ShouldBeEqual(t, 0, "TypeStatus IsEqual diff right -- different Right", actual)
}

// ==========================================================================
// CastTo — additional paths
// ==========================================================================

func Test_CastTo_MatchingType_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{
		"isValid":   result.IsValid,
		"matching":  result.IsMatchingAcceptedType,
		"hasAnyIss": result.HasAnyIssues(),
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"matching": true,
		"hasAnyIss": false,
	}
	expected.ShouldBeEqual(t, 0, "CastTo matching type -- string", actual)
}

func Test_CastTo_NotMatching_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{
		"matching": result.IsMatchingAcceptedType,
		"hasErr":   result.HasError(),
	}

	// Assert
	expected := args.Map{
		"matching": false,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "CastTo not matching -- string vs int", actual)
}

// ==========================================================================
// CastedResult — additional methods
// ==========================================================================

func Test_CastedResult_Methods(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{
		IsValid: true, IsNull: false, IsPointer: false,
		IsMatchingAcceptedType: true, SourceKind: reflect.String,
	}

	// Act
	actual := args.Map{
		"isNotNull":     cr.IsNotNull(),
		"isNotPtr":      cr.IsNotPointer(),
		"isNotMismatch": cr.IsNotMatchingAcceptedType(),
		"isSourceKind":  cr.IsSourceKind(reflect.String),
		"hasAnyIssues":  cr.HasAnyIssues(),
	}

	// Assert
	expected := args.Map{
		"isNotNull": true, "isNotPtr": true,
		"isNotMismatch": false, "isSourceKind": true,
		"hasAnyIssues": false,
	}
	expected.ShouldBeEqual(t, 0, "CastedResult methods -- valid result", actual)
}

// ==========================================================================
// Dynamic — Clone, NonPtr, Ptr
// ==========================================================================

func Test_Dynamic_Clone_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	cloned := d.Clone()
	clonedPtr := d.ClonePtr()
	var nilD *coredynamic.Dynamic
	nilClone := nilD.ClonePtr()

	// Act
	actual := args.Map{
		"cloneData":   cloned.Data(),
		"ptrNotNil":   clonedPtr != nil,
		"nilClone":    nilClone == nil,
		"nonPtrData":  d.NonPtr().Data(),
		"ptrNotNil2":  d.Ptr() != nil,
	}

	// Assert
	expected := args.Map{
		"cloneData": "hello", "ptrNotNil": true,
		"nilClone": true, "nonPtrData": "hello", "ptrNotNil2": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic Clone/NonPtr/Ptr -- string", actual)
}

// ==========================================================================
// Dynamic — InvalidDynamic constructor
// ==========================================================================

func Test_Dynamic_InvalidDynamic_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()

	// Act
	actual := args.Map{
		"isValid": d.IsValid(),
		"isNull": d.IsNull(),
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"isNull": true,
	}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic returns invalid -- no data", actual)
}

func Test_Dynamic_InvalidDynamicPtr_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamicPtr()

	// Act
	actual := args.Map{
		"notNil": d != nil,
		"isValid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isValid": false,
	}
	expected.ShouldBeEqual(t, 0, "InvalidDynamicPtr returns ptr -- invalid", actual)
}
