package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// SimpleResult — constructors
// =============================================================================

func Test_Cov42_SimpleResult_InvalidNoMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleResultNoMessage()
	actual := args.Map{"valid": r.IsValid(), "msg": r.Message}
	expected := args.Map{"valid": false, "msg": ""}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidNoMessage", actual)
}

func Test_Cov42_SimpleResult_Invalid(t *testing.T) {
	r := coredynamic.InvalidSimpleResult("fail reason")
	actual := args.Map{"valid": r.IsValid(), "msg": r.Message}
	expected := args.Map{"valid": false, "msg": "fail reason"}
	expected.ShouldBeEqual(t, 0, "SimpleResult Invalid", actual)
}

func Test_Cov42_SimpleResult_NewValid(t *testing.T) {
	r := coredynamic.NewSimpleResultValid(42)
	actual := args.Map{"valid": r.IsValid(), "result": r.Result}
	expected := args.Map{"valid": true, "result": 42}
	expected.ShouldBeEqual(t, 0, "SimpleResult NewValid", actual)
}

func Test_Cov42_SimpleResult_New(t *testing.T) {
	r := coredynamic.NewSimpleResult("data", true, "")
	actual := args.Map{"valid": r.IsValid(), "result": r.Result}
	expected := args.Map{"valid": true, "result": "data"}
	expected.ShouldBeEqual(t, 0, "SimpleResult New", actual)
}

// =============================================================================
// SimpleResult — GetErrorOnTypeMismatch
// =============================================================================

func Test_Cov42_SimpleResult_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch nil", actual)
}

func Test_Cov42_SimpleResult_GetErrorOnTypeMismatch_Match(t *testing.T) {
	r := coredynamic.NewSimpleResult("data", true, "")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch match", actual)
}

func Test_Cov42_SimpleResult_GetErrorOnTypeMismatch_NoInclude(t *testing.T) {
	r := coredynamic.NewSimpleResult(42, true, "msg")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch no include", actual)
}

func Test_Cov42_SimpleResult_GetErrorOnTypeMismatch_Include(t *testing.T) {
	r := coredynamic.NewSimpleResult(42, true, "extra msg")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch include", actual)
}

// =============================================================================
// SimpleResult — InvalidError
// =============================================================================

func Test_Cov42_SimpleResult_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	err := r.InvalidError()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError nil", actual)
}

func Test_Cov42_SimpleResult_InvalidError_NoMessage(t *testing.T) {
	r := coredynamic.NewSimpleResult(nil, false, "")
	err := r.InvalidError()
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError no message", actual)
}

func Test_Cov42_SimpleResult_InvalidError_WithMessage(t *testing.T) {
	r := coredynamic.NewSimpleResult(nil, false, "bad input")
	err := r.InvalidError()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError with message", actual)
}

func Test_Cov42_SimpleResult_InvalidError_Cached(t *testing.T) {
	r := coredynamic.NewSimpleResult(nil, false, "bad")
	e1 := r.InvalidError()
	e2 := r.InvalidError()
	actual := args.Map{"same": e1 == e2}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError cached", actual)
}

// =============================================================================
// SimpleResult — Clone / ClonePtr
// =============================================================================

func Test_Cov42_SimpleResult_Clone_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	c := r.Clone()
	actual := args.Map{"valid": c.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SimpleResult Clone nil", actual)
}

func Test_Cov42_SimpleResult_Clone_Valid(t *testing.T) {
	r := coredynamic.NewSimpleResult("data", true, "msg")
	c := r.Clone()
	actual := args.Map{"valid": c.IsValid(), "msg": c.Message}
	expected := args.Map{"valid": true, "msg": "msg"}
	expected.ShouldBeEqual(t, 0, "SimpleResult Clone valid", actual)
}

func Test_Cov42_SimpleResult_ClonePtr_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	actual := args.Map{"isNil": r.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult ClonePtr nil", actual)
}

func Test_Cov42_SimpleResult_ClonePtr_Valid(t *testing.T) {
	r := coredynamic.NewSimpleResult("data", true, "")
	c := r.ClonePtr()
	actual := args.Map{"notNil": c != nil, "valid": c.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult ClonePtr valid", actual)
}

// =============================================================================
// Dynamic — Clone / ClonePtr / NonPtr / Ptr / Constructors
// =============================================================================

func Test_Cov42_Dynamic_Clone(t *testing.T) {
	d := coredynamic.NewDynamic("hello", true)
	c := d.Clone()
	actual := args.Map{"valid": c.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Clone", actual)
}

func Test_Cov42_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	actual := args.Map{"isNil": d.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ClonePtr nil", actual)
}

func Test_Cov42_Dynamic_ClonePtr_Valid(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	c := d.ClonePtr()
	actual := args.Map{"notNil": c != nil, "valid": c.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ClonePtr valid", actual)
}

func Test_Cov42_Dynamic_NonPtr(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	np := d.NonPtr()
	actual := args.Map{"valid": np.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic NonPtr", actual)
}

func Test_Cov42_Dynamic_Ptr(t *testing.T) {
	d := coredynamic.NewDynamic(42, true)
	p := d.Ptr()
	actual := args.Map{"notNil": p != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Ptr", actual)
}

func Test_Cov42_Dynamic_InvalidDynamic(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic", actual)
}

func Test_Cov42_Dynamic_NewDynamicValid(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"valid": d.IsValid()}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewDynamicValid", actual)
}

// =============================================================================
// DynamicStatus — constructors / clone
// =============================================================================

func Test_Cov42_DynamicStatus_InvalidNoMessage(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatusNoMessage()
	actual := args.Map{"valid": ds.IsValid(), "msg": ds.Message}
	expected := args.Map{"valid": false, "msg": ""}
	expected.ShouldBeEqual(t, 0, "DynamicStatus InvalidNoMessage", actual)
}

func Test_Cov42_DynamicStatus_Invalid(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("fail")
	actual := args.Map{"valid": ds.IsValid(), "msg": ds.Message}
	expected := args.Map{"valid": false, "msg": "fail"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus Invalid", actual)
}

func Test_Cov42_DynamicStatus_Clone(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	c := ds.Clone()
	actual := args.Map{"msg": c.Message}
	expected := args.Map{"msg": "msg"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus Clone", actual)
}

func Test_Cov42_DynamicStatus_ClonePtr_Nil(t *testing.T) {
	var ds *coredynamic.DynamicStatus
	actual := args.Map{"isNil": ds.ClonePtr() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus ClonePtr nil", actual)
}

func Test_Cov42_DynamicStatus_ClonePtr_Valid(t *testing.T) {
	ds := coredynamic.InvalidDynamicStatus("msg")
	c := ds.ClonePtr()
	actual := args.Map{"notNil": c != nil, "msg": c.Message}
	expected := args.Map{"notNil": true, "msg": "msg"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus ClonePtr valid", actual)
}

// =============================================================================
// LeftRight — all branches
// =============================================================================

func Test_Cov42_LeftRight_IsEmpty_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"r": lr.IsEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEmpty nil", actual)
}

func Test_Cov42_LeftRight_IsEmpty_BothNil(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}
	actual := args.Map{"r": lr.IsEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEmpty both nil", actual)
}

func Test_Cov42_LeftRight_IsEmpty_HasData(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}
	actual := args.Map{"r": lr.IsEmpty()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEmpty has data", actual)
}

func Test_Cov42_LeftRight_HasAnyItem(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}
	actual := args.Map{"r": lr.HasAnyItem()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight HasAnyItem", actual)
}

func Test_Cov42_LeftRight_HasLeft_True(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}
	actual := args.Map{"r": lr.HasLeft()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight HasLeft true", actual)
}

func Test_Cov42_LeftRight_HasLeft_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"r": lr.HasLeft()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LeftRight HasLeft nil", actual)
}

func Test_Cov42_LeftRight_HasRight_True(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: "b"}
	actual := args.Map{"r": lr.HasRight()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight HasRight true", actual)
}

func Test_Cov42_LeftRight_HasRight_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"r": lr.HasRight()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LeftRight HasRight nil", actual)
}

func Test_Cov42_LeftRight_IsLeftEmpty_True(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: "b"}
	actual := args.Map{"r": lr.IsLeftEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsLeftEmpty true", actual)
}

func Test_Cov42_LeftRight_IsLeftEmpty_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"r": lr.IsLeftEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsLeftEmpty nil", actual)
}

func Test_Cov42_LeftRight_IsRightEmpty_True(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}
	actual := args.Map{"r": lr.IsRightEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsRightEmpty true", actual)
}

func Test_Cov42_LeftRight_IsRightEmpty_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"r": lr.IsRightEmpty()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsRightEmpty nil", actual)
}

func Test_Cov42_LeftRight_LeftReflectSet_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	err := lr.LeftReflectSet(nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftReflectSet nil", actual)
}

func Test_Cov42_LeftRight_RightReflectSet_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	err := lr.RightReflectSet(nil)
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightReflectSet nil", actual)
}

func Test_Cov42_LeftRight_DeserializeLeft_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"isNil": lr.DeserializeLeft() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft nil", actual)
}

func Test_Cov42_LeftRight_DeserializeLeft_Valid(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "data", Right: nil}
	r := lr.DeserializeLeft()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft valid", actual)
}

func Test_Cov42_LeftRight_DeserializeRight_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"isNil": lr.DeserializeRight() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeRight nil", actual)
}

func Test_Cov42_LeftRight_DeserializeRight_Valid(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: "data"}
	r := lr.DeserializeRight()
	actual := args.Map{"notNil": r != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeRight valid", actual)
}

func Test_Cov42_LeftRight_LeftToDynamic_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"isNil": lr.LeftToDynamic() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftToDynamic nil", actual)
}

func Test_Cov42_LeftRight_LeftToDynamic_Valid(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}
	d := lr.LeftToDynamic()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftToDynamic valid", actual)
}

func Test_Cov42_LeftRight_RightToDynamic_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	actual := args.Map{"isNil": lr.RightToDynamic() == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightToDynamic nil", actual)
}

func Test_Cov42_LeftRight_RightToDynamic_Valid(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: nil, Right: "b"}
	d := lr.RightToDynamic()
	actual := args.Map{"notNil": d != nil, "valid": d.IsValid()}
	expected := args.Map{"notNil": true, "valid": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightToDynamic valid", actual)
}

func Test_Cov42_LeftRight_TypeStatus_Nil(t *testing.T) {
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()
	actual := args.Map{"notZero": ts.IsSame}
	expected := args.Map{"notZero": true}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus nil", actual)
}

func Test_Cov42_LeftRight_TypeStatus_Valid(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()
	actual := args.Map{"same": ts.IsSame}
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus valid same type", actual)
}

func Test_Cov42_LeftRight_TypeStatus_DiffType(t *testing.T) {
	lr := &coredynamic.LeftRight{Left: "a", Right: 42}
	ts := lr.TypeStatus()
	actual := args.Map{"same": ts.IsSame}
	expected := args.Map{"same": false}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus diff type", actual)
}

// =============================================================================
// CastedResult — all branches
// =============================================================================

func Test_Cov42_CastedResult_IsInvalid_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"r": cr.IsInvalid()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsInvalid nil", actual)
}

func Test_Cov42_CastedResult_IsInvalid_Invalid(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: false}
	actual := args.Map{"r": cr.IsInvalid()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsInvalid invalid", actual)
}

func Test_Cov42_CastedResult_IsInvalid_Valid(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: true}
	actual := args.Map{"r": cr.IsInvalid()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsInvalid valid", actual)
}

func Test_Cov42_CastedResult_IsNotNull_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"r": cr.IsNotNull()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotNull nil", actual)
}

func Test_Cov42_CastedResult_IsNotNull_True(t *testing.T) {
	cr := &coredynamic.CastedResult{IsNull: false}
	actual := args.Map{"r": cr.IsNotNull()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotNull true", actual)
}

func Test_Cov42_CastedResult_IsNotPointer_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"r": cr.IsNotPointer()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotPointer nil", actual)
}

func Test_Cov42_CastedResult_IsNotPointer_True(t *testing.T) {
	cr := &coredynamic.CastedResult{IsPointer: false}
	actual := args.Map{"r": cr.IsNotPointer()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotPointer true", actual)
}

func Test_Cov42_CastedResult_IsNotMatchingAcceptedType_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"r": cr.IsNotMatchingAcceptedType()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotMatchingAcceptedType nil", actual)
}

func Test_Cov42_CastedResult_IsNotMatchingAcceptedType_True(t *testing.T) {
	cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}
	actual := args.Map{"r": cr.IsNotMatchingAcceptedType()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotMatchingAcceptedType true", actual)
}

func Test_Cov42_CastedResult_IsSourceKind_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"r": cr.IsSourceKind(reflect.String)}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsSourceKind nil", actual)
}

func Test_Cov42_CastedResult_IsSourceKind_Match(t *testing.T) {
	cr := &coredynamic.CastedResult{SourceKind: reflect.String}
	actual := args.Map{"r": cr.IsSourceKind(reflect.String)}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsSourceKind match", actual)
}

func Test_Cov42_CastedResult_HasError_Nil(t *testing.T) {
	var cr *coredynamic.CastedResult
	actual := args.Map{"r": cr.HasError()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult HasError nil", actual)
}

func Test_Cov42_CastedResult_HasError_NoError(t *testing.T) {
	cr := &coredynamic.CastedResult{}
	actual := args.Map{"r": cr.HasError()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult HasError no error", actual)
}

func Test_Cov42_CastedResult_HasAnyIssues_Invalid(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: false}
	actual := args.Map{"r": cr.HasAnyIssues()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues invalid", actual)
}

func Test_Cov42_CastedResult_HasAnyIssues_Null(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: true}
	actual := args.Map{"r": cr.HasAnyIssues()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues null", actual)
}

func Test_Cov42_CastedResult_HasAnyIssues_NotMatching(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: false}
	actual := args.Map{"r": cr.HasAnyIssues()}
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues not matching", actual)
}

func Test_Cov42_CastedResult_HasAnyIssues_AllGood(t *testing.T) {
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: true}
	actual := args.Map{"r": cr.HasAnyIssues()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues all good", actual)
}

// =============================================================================
// BytesConverter — basic conversions
// =============================================================================

func Test_Cov42_BytesConverter_SafeCastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter(nil)
	actual := args.Map{"r": bc.SafeCastString()}
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString empty", actual)
}

func Test_Cov42_BytesConverter_SafeCastString_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	actual := args.Map{"r": bc.SafeCastString()}
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString valid", actual)
}

func Test_Cov42_BytesConverter_CastString_Empty(t *testing.T) {
	bc := coredynamic.NewBytesConverter(nil)
	_, err := bc.CastString()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString empty", actual)
}

func Test_Cov42_BytesConverter_CastString_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	s, err := bc.CastString()
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString valid", actual)
}

func Test_Cov42_BytesConverter_ToBool_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("true"))
	r, err := bc.ToBool()
	actual := args.Map{"noErr": err == nil, "r": r}
	expected := args.Map{"noErr": true, "r": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBool valid", actual)
}

func Test_Cov42_BytesConverter_ToBoolMust_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("false"))
	actual := args.Map{"r": bc.ToBoolMust()}
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBoolMust valid", actual)
}

func Test_Cov42_BytesConverter_ToString_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()
	actual := args.Map{"noErr": err == nil, "r": s}
	expected := args.Map{"noErr": true, "r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToString valid", actual)
}

func Test_Cov42_BytesConverter_ToStringMust_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"hi"`))
	actual := args.Map{"r": bc.ToStringMust()}
	expected := args.Map{"r": "hi"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringMust valid", actual)
}

func Test_Cov42_BytesConverter_ToStrings_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	s, err := bc.ToStrings()
	actual := args.Map{"noErr": err == nil, "len": len(s)}
	expected := args.Map{"noErr": true, "len": 2}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStrings valid", actual)
}

func Test_Cov42_BytesConverter_ToStringsMust_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	s := bc.ToStringsMust()
	actual := args.Map{"len": len(s)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringsMust valid", actual)
}

func Test_Cov42_BytesConverter_ToInt64_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("42"))
	v, err := bc.ToInt64()
	actual := args.Map{"noErr": err == nil, "v": v}
	expected := args.Map{"noErr": true, "v": int64(42)}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64 valid", actual)
}

func Test_Cov42_BytesConverter_ToInt64Must_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("99"))
	actual := args.Map{"v": bc.ToInt64Must()}
	expected := args.Map{"v": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64Must valid", actual)
}

func Test_Cov42_BytesConverter_Deserialize_Valid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte(`"data"`))
	var out string
	err := bc.Deserialize(&out)
	actual := args.Map{"noErr": err == nil, "val": out}
	expected := args.Map{"noErr": true, "val": "data"}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize valid", actual)
}

func Test_Cov42_BytesConverter_Deserialize_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("not json"))
	var out string
	err := bc.Deserialize(&out)
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize invalid", actual)
}

func Test_Cov42_BytesConverter_ToHashmap_Valid(t *testing.T) {
	data := map[string]string{"a": "1"}
	b, _ := json.Marshal(data)
	bc := coredynamic.NewBytesConverter(b)
	hm, err := bc.ToHashmap()
	actual := args.Map{"noErr": err == nil, "notNil": hm != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap valid", actual)
}

func Test_Cov42_BytesConverter_ToHashmap_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToHashmap()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap invalid", actual)
}

func Test_Cov42_BytesConverter_ToHashset_Valid(t *testing.T) {
	data := map[string]bool{"a": true, "b": true}
	b, _ := json.Marshal(data)
	bc := coredynamic.NewBytesConverter(b)
	hs, err := bc.ToHashset()
	actual := args.Map{"noErr": err == nil, "notNil": hs != nil}
	expected := args.Map{"noErr": true, "notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset valid", actual)
}

func Test_Cov42_BytesConverter_ToHashset_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToHashset()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset invalid", actual)
}

func Test_Cov42_BytesConverter_ToCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToCollection invalid", actual)
}

func Test_Cov42_BytesConverter_ToSimpleSlice_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToSimpleSlice()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToSimpleSlice invalid", actual)
}

func Test_Cov42_BytesConverter_ToKeyValCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToKeyValCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToKeyValCollection invalid", actual)
}

func Test_Cov42_BytesConverter_ToAnyCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToAnyCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToAnyCollection invalid", actual)
}

func Test_Cov42_BytesConverter_ToMapAnyItems_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToMapAnyItems()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToMapAnyItems invalid", actual)
}

func Test_Cov42_BytesConverter_ToMapAnyItems_Valid(t *testing.T) {
	b, _ := json.Marshal(map[string]any{"Items": map[string]any{"a": 1}})
	bc := coredynamic.NewBytesConverter(b)
	m, err := bc.ToMapAnyItems()
	actual := args.Map{"noErr": err == nil, "len": m.Length()}
	expected := args.Map{"noErr": true, "len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToMapAnyItems valid", actual)
}

func Test_Cov42_BytesConverter_ToDynamicCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToDynamicCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToDynamicCollection invalid", actual)
}

func Test_Cov42_BytesConverter_ToJsonResultCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToJsonResultCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonResultCollection invalid", actual)
}

func Test_Cov42_BytesConverter_ToJsonMapResults_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToJsonMapResults()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonMapResults invalid", actual)
}

func Test_Cov42_BytesConverter_ToBytesCollection_Invalid(t *testing.T) {
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToBytesCollection()
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBytesCollection invalid", actual)
}
