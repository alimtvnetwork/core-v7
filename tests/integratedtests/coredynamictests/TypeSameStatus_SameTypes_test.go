package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// TypeSameStatus
// ═══════════════════════════════════════════

func Test_TypeSameStatus_SameTypes_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", "world")

	// Act
	actual := args.Map{
		"isSame":    ts.IsSame,
		"isNotSame": ts.IsNotSame(),
		"isValid":   ts.IsValid(),
		"leftNN":    ts.Left != nil,
		"rightNN":   ts.Right != nil,
	}

	// Assert
	expected := args.Map{
		"isSame": true, "isNotSame": false,
		"isValid": true, "leftNN": true, "rightNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- same types", actual)
}

func Test_TypeSameStatus_DiffTypes_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", 42)

	// Act
	actual := args.Map{
		"isSame":  ts.IsSame,
		"notSame": ts.IsNotSame(),
		"isNEq":   ts.IsNotEqualTypes(),
	}

	// Assert
	expected := args.Map{
		"isSame": false,
		"notSame": true,
		"isNEq": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- diff types", actual)
}

func Test_TypeSameStatus_Pointers_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "world")

	// Act
	actual := args.Map{
		"isAnyPtr":  ts.IsAnyPointer(),
		"isBothPtr": ts.IsBothPointer(),
		"leftPtr":   ts.IsLeftPointer,
		"rightPtr":  ts.IsRightPointer,
	}

	// Assert
	expected := args.Map{
		"isAnyPtr": true, "isBothPtr": false,
		"leftPtr": true, "rightPtr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns correct value -- pointers", actual)
}

func Test_TypeSameStatus_NilInput(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus(nil, "hello")

	// Act
	actual := args.Map{
		"leftNull":  ts.IsLeftUnknownNull,
		"rightNull": ts.IsRightUnknownNull,
		"isSame":    ts.IsSame,
	}

	// Assert
	expected := args.Map{
		"leftNull": true,
		"rightNull": false,
		"isSame": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeSameStatus returns nil -- nil input", actual)
}

// ═══════════════════════════════════════════
// TypeStatus — methods
// ═══════════════════════════════════════════

func Test_TypeStatus_Methods_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", "world")

	// Act
	actual := args.Map{
		"isValid":       ts.IsValid(),
		"isInvalid":     ts.IsInvalid(),
		"sameRegardless": ts.IsSameRegardlessPointer(),
		"leftName":      ts.LeftName(),
		"rightName":     ts.RightName(),
		"leftFullName":  ts.LeftFullName(),
		"rightFullName": ts.RightFullName(),
	}

	// Assert
	expected := args.Map{
		"isValid": true, "isInvalid": false,
		"sameRegardless": true,
		"leftName": "string", "rightName": "string",
		"leftFullName": "string", "rightFullName": "string",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- methods", actual)
}

func Test_TypeStatus_NilReceiver_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	var nilTS *coredynamic.TypeStatus

	// Act
	actual := args.Map{
		"isValid":   nilTS.IsValid(),
		"isInvalid": nilTS.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns nil -- nil receiver", actual)
}

func Test_TypeStatus_NotMatchMessage(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", 42)
	msg := ts.NotMatchMessage("left", "right")
	sameTS := coredynamic.TypeSameStatus("a", "b")
	sameMsg := sameTS.NotMatchMessage("l", "r")

	// Act
	actual := args.Map{
		"msgNE":   msg != "",
		"sameMsg": sameMsg,
	}

	// Assert
	expected := args.Map{
		"msgNE": true,
		"sameMsg": "",
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- not match message", actual)
}

func Test_TypeStatus_NotMatchErr(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", 42)
	err := ts.NotMatchErr("l", "r")
	sameTS := coredynamic.TypeSameStatus("a", "b")
	sameErr := sameTS.NotMatchErr("l", "r")

	// Act
	actual := args.Map{
		"hasErr":  err != nil,
		"noErr":   sameErr == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- not match err", actual)
}

func Test_TypeStatus_ValidationError(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", 42)
	err := ts.ValidationError()
	sameTS := coredynamic.TypeSameStatus("a", "b")
	sameErr := sameTS.ValidationError()

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"noErr": sameErr == nil,
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns error -- validation error", actual)
}

func Test_TypeStatus_NotEqualSrcDest(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", 42)
	msg := ts.NotEqualSrcDestinationMessage()
	err := ts.NotEqualSrcDestinationErr()

	// Act
	actual := args.Map{
		"msgNE": msg != "",
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"msgNE": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- src dest", actual)
}

func Test_TypeStatus_MustBeSame(t *testing.T) {
	// Arrange
	sameTS := coredynamic.TypeSameStatus("a", "b")
	sameTS.MustBeSame() // should not panic
	sameTS.SrcDestinationMustBeSame() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- must be same", actual)
}

func Test_TypeStatus_IsEqual_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	ts1 := coredynamic.TypeSameStatus("a", "b")
	ts2 := coredynamic.TypeSameStatus("a", "b")
	ts3 := coredynamic.TypeSameStatus("a", 42)
	var nilTS *coredynamic.TypeStatus

	// Act
	actual := args.Map{
		"equal":    ts1.IsEqual(&ts2),
		"notEqual": ts1.IsEqual(&ts3),
		"bothNil":  nilTS.IsEqual(nil),
		"oneNil":   ts1.IsEqual(nil),
	}

	// Assert
	expected := args.Map{
		"equal": true, "notEqual": false,
		"bothNil": true, "oneNil": false,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- is equal", actual)
}

func Test_TypeStatus_NonPointerLeftRight(t *testing.T) {
	// Arrange
	s := "hello"
	ts := coredynamic.TypeSameStatus(&s, "world")

	// Act
	actual := args.Map{
		"nonPtrLeftName":  ts.NonPointerLeft().Name(),
		"nonPtrRightName": ts.NonPointerRight().Name(),
		"sameRegardless":  ts.IsSameRegardlessPointer(),
	}

	// Assert
	expected := args.Map{
		"nonPtrLeftName": "string", "nonPtrRightName": "string",
		"sameRegardless": true,
	}
	expected.ShouldBeEqual(t, 0, "TypeStatus returns non-empty -- non-pointer left/right", actual)
}

func Test_TypeStatus_NullNames_FromTypeSameStatusSameTy(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "TypeStatus returns correct value -- null names", actual)
}

// ═══════════════════════════════════════════
// TypedDynamic
// ═══════════════════════════════════════════

func Test_TypedDynamic_Constructors(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	dv := coredynamic.NewTypedDynamicValid("world")
	dp := coredynamic.NewTypedDynamicPtr("ptr", true)
	inv := coredynamic.InvalidTypedDynamic[string]()
	invP := coredynamic.InvalidTypedDynamicPtr[string]()

	// Act
	actual := args.Map{
		"data":     d.Data(),
		"value":    d.Value(),
		"valid":    d.IsValid(),
		"dvData":   dv.Data(),
		"dvValid":  dv.IsValid(),
		"dpNN":     dp != nil,
		"dpData":   dp.Data(),
		"invValid": inv.IsValid(),
		"invInv":   inv.IsInvalid(),
		"invPNN":   invP != nil,
	}

	// Assert
	expected := args.Map{
		"data": "hello", "value": "hello", "valid": true,
		"dvData": "world", "dvValid": true,
		"dpNN": true, "dpData": "ptr",
		"invValid": false, "invInv": true, "invPNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- constructors", actual)
}

func Test_TypedDynamic_GetAs(t *testing.T) {
	// Arrange
	ds := coredynamic.NewTypedDynamic("hello", true)
	di := coredynamic.NewTypedDynamic(42, true)
	db := coredynamic.NewTypedDynamic(true, true)
	df := coredynamic.NewTypedDynamic(3.14, true)
	strVal, strOK := ds.GetAsString()
	intVal, intOK := di.GetAsInt()
	boolVal, boolOK := db.GetAsBool()
	fVal, fOK := df.GetAsFloat64()
	_, f32OK := df.GetAsFloat32()
	_, i64OK := di.GetAsInt64()
	_, uiOK := di.GetAsUint()
	_, bOK := ds.GetAsBytes()
	_, ssOK := ds.GetAsStrings()

	// Act
	actual := args.Map{
		"str": strVal, "strOK": strOK,
		"int": intVal, "intOK": intOK,
		"bool": boolVal, "boolOK": boolOK,
		"float": fVal > 3.0, "fOK": fOK,
		"f32OK": f32OK, "i64OK": i64OK, "uiOK": uiOK,
		"bOK": bOK, "ssOK": ssOK,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "strOK": true,
		"int": 42, "intOK": true,
		"bool": true, "boolOK": true,
		"float": true, "fOK": true,
		"f32OK": false, "i64OK": false, "uiOK": false,
		"bOK": false, "ssOK": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- getAs", actual)
}

func Test_TypedDynamic_Value(t *testing.T) {
	// Arrange
	ds := coredynamic.NewTypedDynamic("hello", true)
	di := coredynamic.NewTypedDynamic(42, true)
	db := coredynamic.NewTypedDynamic(true, true)
	d64 := coredynamic.NewTypedDynamic(int64(99), true)

	// Act
	actual := args.Map{
		"valStr":  ds.ValueString(),
		"valInt":  di.ValueInt(),
		"valBool": db.ValueBool(),
		"valI64":  d64.ValueInt64(),
		"strStr":  ds.String(),
	}

	// Assert
	expected := args.Map{
		"valStr": "hello", "valInt": 42,
		"valBool": true, "valI64": int64(99),
		"strStr": "hello",
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- value", actual)
}

func Test_TypedDynamic_Json(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	jb, jErr := d.JsonBytes()
	jr := d.JsonResult()
	j := d.Json()
	jp := d.JsonPtr()
	js, jsErr := d.JsonString()
	mb, mErr := d.MarshalJSON()
	vm, vmErr := d.ValueMarshal()
	model := d.JsonModel()
	modelAny := d.JsonModelAny()

	// Act
	actual := args.Map{
		"jbLen": len(jb) > 0, "jErr": jErr == nil,
		"jrLen": jr.Length() > 0,
		"jLen":  j.Length() > 0,
		"jpNN":  jp != nil,
		"jsNE":  js != "", "jsErr": jsErr == nil,
		"mbLen": len(mb) > 0, "mErr": mErr == nil,
		"vmLen": len(vm) > 0, "vmErr": vmErr == nil,
		"model": model, "modelAnyNN": modelAny != nil,
	}

	// Assert
	expected := args.Map{
		"jbLen": true, "jErr": true,
		"jrLen": true, "jLen": true, "jpNN": true,
		"jsNE": true, "jsErr": true,
		"mbLen": true, "mErr": true,
		"vmLen": true, "vmErr": true,
		"model": "hello", "modelAnyNN": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- json", actual)
}

func Test_TypedDynamic_Bytes(t *testing.T) {
	// Arrange
	db := coredynamic.NewTypedDynamic([]byte("hello"), true)
	ds := coredynamic.NewTypedDynamic("hello", true)
	bytesB, okB := db.Bytes()
	bytesS, okS := ds.Bytes()

	// Act
	actual := args.Map{
		"bLen": len(bytesB) > 0, "bOK": okB,
		"sLen": len(bytesS) > 0, "sOK": okS,
	}

	// Assert
	expected := args.Map{
		"bLen": true, "bOK": true,
		"sLen": true, "sOK": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- bytes", actual)
}

func Test_TypedDynamic_ClonePtrNonPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamic("hello", true)
	cloned := d.Clone()
	dp := coredynamic.NewTypedDynamicPtr("test", true)
	clonedPtr := dp.ClonePtr()
	var nilDP *coredynamic.TypedDynamic[string]
	nilClone := nilDP.ClonePtr()
	nonPtr := d.NonPtr()
	ptr := dp.Ptr()
	toDyn := d.ToDynamic()

	// Act
	actual := args.Map{
		"clonedData": cloned.Data(),
		"clonePtrNN": clonedPtr != nil,
		"nilCloneNil": nilClone == nil,
		"nonPtrData":  nonPtr.Data(),
		"ptrNN":       ptr != nil,
		"toDynValid":  toDyn.IsValid(),
	}

	// Assert
	expected := args.Map{
		"clonedData": "hello", "clonePtrNN": true,
		"nilCloneNil": true, "nonPtrData": "hello",
		"ptrNN": true, "toDynValid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- clone/ptr", actual)
}

func Test_TypedDynamic_UnmarshalDeserialize(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("", false)
	err := d.UnmarshalJSON([]byte(`"hello"`))

	// Act
	actual := args.Map{
		"errNil":  err == nil,
		"data":    d.Data(),
		"isValid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"data": "hello",
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- unmarshal", actual)
}

func Test_TypedDynamic_UnmarshalBadJSON(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr(0, false)
	err := d.UnmarshalJSON([]byte(`"not-an-int"`))

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"isValid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"isValid": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- unmarshal bad json", actual)
}

func Test_TypedDynamic_Deserialize(t *testing.T) {
	// Arrange
	d := coredynamic.NewTypedDynamicPtr("", false)
	jsonBytes, _ := json.Marshal("world")
	err := d.Deserialize(jsonBytes)

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"data": d.Data(),
		"valid": d.IsValid(),
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"data": "world",
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns correct value -- deserialize", actual)
}

func Test_TypedDynamic_DeserializeNilReceiver(t *testing.T) {
	// Arrange
	var nilD *coredynamic.TypedDynamic[string]
	err := nilD.Deserialize([]byte(`"hello"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TypedDynamic returns nil -- deserialize nil", actual)
}

// ═══════════════════════════════════════════
// SimpleResult
// ═══════════════════════════════════════════

func Test_SimpleResult_Constructors(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("hello")
	srInv := coredynamic.InvalidSimpleResult("err-msg")
	srInvNM := coredynamic.InvalidSimpleResultNoMessage()
	srFull := coredynamic.NewSimpleResult("val", false, "msg")

	// Act
	actual := args.Map{
		"srValid":    sr.IsValid(),
		"srResult":   sr.Result,
		"invValid":   srInv.IsValid(),
		"invMsg":     srInv.Message,
		"invNMValid": srInvNM.IsValid(),
		"fullResult": srFull.Result,
		"fullMsg":    srFull.Message,
	}

	// Assert
	expected := args.Map{
		"srValid": true, "srResult": "hello",
		"invValid": false, "invMsg": "err-msg",
		"invNMValid": false, "fullResult": "val", "fullMsg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- constructors", actual)
}

func Test_SimpleResult_InvalidError(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("hello")
	srInv := coredynamic.InvalidSimpleResult("err-msg")

	// Act
	actual := args.Map{
		"validErr": sr.InvalidError() == nil,
		"invErr":   srInv.InvalidError() != nil,
	}

	// Assert
	expected := args.Map{
		"validErr": true,
		"invErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns error -- invalid error", actual)
}

func Test_SimpleResult_TypeMismatch(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("hello")
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	nilErr := sr.GetErrorOnTypeMismatch(strType, false)
	hasErr := sr.GetErrorOnTypeMismatch(intType, false)
	hasErrMsg := sr.GetErrorOnTypeMismatch(intType, true)

	// Act
	actual := args.Map{
		"nilErr":  nilErr == nil,
		"hasErr":  hasErr != nil,
		"hasMsg":  hasErrMsg != nil,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"hasErr": true,
		"hasMsg": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- type mismatch", actual)
}

func Test_SimpleResult_Clone_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResultValid("hello")
	cloned := sr.Clone()
	clonedPtr := sr.ClonePtr()
	var nilSR *coredynamic.SimpleResult
	nilClone := nilSR.ClonePtr()

	// Act
	actual := args.Map{
		"clonedResult": cloned.Result,
		"ptrNN":        clonedPtr != nil,
		"nilCloneNil":  nilClone == nil,
	}

	// Assert
	expected := args.Map{
		"clonedResult": "hello", "ptrNN": true, "nilCloneNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult returns correct value -- clone", actual)
}

// ═══════════════════════════════════════════
// SimpleRequest
// ═══════════════════════════════════════════

func Test_SimpleRequest_Constructors_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")
	srInv := coredynamic.InvalidSimpleRequest("err-msg")
	srInvNM := coredynamic.InvalidSimpleRequestNoMessage()
	srFull := coredynamic.NewSimpleRequest("val", false, "msg")

	// Act
	actual := args.Map{
		"srValid":    sr.IsValid(),
		"srRequest":  sr.Request(),
		"srValue":    sr.Value(),
		"srMsg":      sr.Message(),
		"invValid":   srInv.IsValid(),
		"invMsg":     srInv.Message(),
		"invNMValid": srInvNM.IsValid(),
		"fullMsg":    srFull.Message(),
	}

	// Assert
	expected := args.Map{
		"srValid": true, "srRequest": "hello", "srValue": "hello", "srMsg": "",
		"invValid": false, "invMsg": "err-msg",
		"invNMValid": false, "fullMsg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- constructors", actual)
}

func Test_SimpleRequest_InvalidError_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")
	srInv := coredynamic.InvalidSimpleRequest("err-msg")

	// Act
	actual := args.Map{
		"validErr": sr.InvalidError() == nil,
		"invErr":   srInv.InvalidError() != nil,
	}

	// Assert
	expected := args.Map{
		"validErr": true,
		"invErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns error -- invalid error", actual)
}

func Test_SimpleRequest_TypeMismatch(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")
	strType := reflect.TypeOf("")
	intType := reflect.TypeOf(0)
	nilErr := sr.GetErrorOnTypeMismatch(strType, false)
	hasErr := sr.GetErrorOnTypeMismatch(intType, true)

	// Act
	actual := args.Map{
		"nilErr": nilErr == nil,
		"hasErr": hasErr != nil,
	}

	// Assert
	expected := args.Map{
		"nilErr": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- type mismatch", actual)
}

func Test_SimpleRequest_IsPointer_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	s := "hello"
	sr := coredynamic.NewSimpleRequestValid(&s)
	srNonPtr := coredynamic.NewSimpleRequestValid("hello")

	// Act
	actual := args.Map{
		"isPtr":     sr.IsPointer(),
		"isNonPtr":  srNonPtr.IsPointer(),
		"isKindStr": srNonPtr.IsReflectKind(reflect.String),
	}

	// Assert
	expected := args.Map{
		"isPtr": true,
		"isNonPtr": false,
		"isKindStr": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleRequest returns correct value -- is pointer", actual)
}

// ═══════════════════════════════════════════
// ValueStatus (coredynamic)
// ═══════════════════════════════════════════

func Test_ValueStatus_Basic(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatusNoMessage()
	vs2 := coredynamic.InvalidValueStatus("err")

	// Act
	actual := args.Map{
		"vsValid":  vs.IsValid,
		"vs2Valid": vs2.IsValid,
		"vs2Msg":   vs2.Message,
		"vs2Index": vs2.Index,
	}

	// Assert
	expected := args.Map{
		"vsValid": false, "vs2Valid": false,
		"vs2Msg": "err", "vs2Index": -1,
	}
	expected.ShouldBeEqual(t, 0, "ValueStatus returns non-empty -- basic", actual)
}

// ═══════════════════════════════════════════
// ZeroSet / SafeZeroSet
// ═══════════════════════════════════════════

func Test_ZeroSet_FromTypeSameStatusSameTy_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	type testStruct struct{ Name string }
	ts := testStruct{Name: "hello"}
	rv := reflect.ValueOf(&ts)
	coredynamic.ZeroSet(rv)

	// Act
	actual := args.Map{"name": ts.Name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "ZeroSet returns correct value -- with args", actual)
}

func Test_SafeZeroSet_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	type testStruct struct{ Name string }
	ts := testStruct{Name: "hello"}
	rv := reflect.ValueOf(&ts)
	coredynamic.SafeZeroSet(rv)

	// Act
	actual := args.Map{"name": ts.Name}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// BytesConverter
// ═══════════════════════════════════════════

func Test_BytesConverter_Basic(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	str, err := bc.ToString()
	strMust := bc.ToStringMust()
	castStr := bc.SafeCastString()
	castStr2, castErr := bc.CastString()

	// Act
	actual := args.Map{
		"str":      str,
		"errNil":   err == nil,
		"must":     strMust,
		"cast":     castStr != "",
		"cast2":    castStr2 != "",
		"castErr":  castErr == nil,
	}

	// Assert
	expected := args.Map{
		"str": "hello", "errNil": true, "must": "hello",
		"cast": true, "cast2": true, "castErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- basic", actual)
}

func Test_BytesConverter_EmptyBytes(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte{})
	castStr := bc.SafeCastString()
	_, castErr := bc.CastString()

	// Act
	actual := args.Map{
		"cast":    castStr,
		"castErr": castErr != nil,
	}

	// Assert
	expected := args.Map{
		"cast": "",
		"castErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns empty -- empty", actual)
}

func Test_BytesConverter_Bool(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`true`))
	b, err := bc.ToBool()
	bMust := bc.ToBoolMust()

	// Act
	actual := args.Map{
		"val": b,
		"errNil": err == nil,
		"must": bMust,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"errNil": true,
		"must": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- bool", actual)
}

func Test_BytesConverter_Strings(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	strs, err := bc.ToStrings()
	must := bc.ToStringsMust()

	// Act
	actual := args.Map{
		"len": len(strs), "errNil": err == nil, "mustLen": len(must),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"errNil": true,
		"mustLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- strings", actual)
}

func Test_BytesConverter_Int64(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`42`))
	val, err := bc.ToInt64()
	must := bc.ToInt64Must()

	// Act
	actual := args.Map{
		"val": val,
		"errNil": err == nil,
		"must": must,
	}

	// Assert
	expected := args.Map{
		"val": int64(42),
		"errNil": true,
		"must": int64(42),
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- int64", actual)
}

func Test_BytesConverter_Deserialize_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	var target string
	err := bc.Deserialize(&target)

	// Act
	actual := args.Map{
		"errNil": err == nil,
		"target": target,
	}

	// Assert
	expected := args.Map{
		"errNil": true,
		"target": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter returns correct value -- deserialize", actual)
}

// ═══════════════════════════════════════════
// Type function
// ═══════════════════════════════════════════

func Test_Type_FromTypeSameStatusSameTy_FromTypeSameStatusSameTy(t *testing.T) {
	// Arrange
	rt := coredynamic.Type("hello")

	// Act
	actual := args.Map{"name": rt.Name()}

	// Assert
	expected := args.Map{"name": "string"}
	expected.ShouldBeEqual(t, 0, "Type returns correct value -- function", actual)
}
