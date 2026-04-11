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

func Test_SimpleResult_InvalidNoMessage_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResultNoMessage()

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"msg": r.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidNoMessage", actual)
}

func Test_SimpleResult_Invalid(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResult("fail reason")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"msg": r.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "fail reason",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult Invalid", actual)
}

func Test_SimpleResult_NewValid(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid(42)

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"result": r.Result,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"result": 42,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult NewValid", actual)
}

func Test_SimpleResult_New(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("data", true, "")

	// Act
	actual := args.Map{
		"valid": r.IsValid(),
		"result": r.Result,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"result": "data",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult New", actual)
}

// =============================================================================
// SimpleResult — GetErrorOnTypeMismatch
// =============================================================================

func Test_SimpleResult_GetErrorOnTypeMismatch_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch nil", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch_Match_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("data", true, "")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch match", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch_NoInclude(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult(42, true, "msg")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch no include", actual)
}

func Test_SimpleResult_GetErrorOnTypeMismatch_Include(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult(42, true, "extra msg")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult GetErrorOnTypeMismatch include", actual)
}

// =============================================================================
// SimpleResult — InvalidError
// =============================================================================

func Test_SimpleResult_InvalidError_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	err := r.InvalidError()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError nil", actual)
}

func Test_SimpleResult_InvalidError_NoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult(nil, false, "")
	err := r.InvalidError()

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError no message", actual)
}

func Test_SimpleResult_InvalidError_WithMessage_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult(nil, false, "bad input")
	err := r.InvalidError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError with message", actual)
}

func Test_SimpleResult_InvalidError_Cached_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult(nil, false, "bad")
	e1 := r.InvalidError()
	e2 := r.InvalidError()

	// Act
	actual := args.Map{"same": e1 == e2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult InvalidError cached", actual)
}

// =============================================================================
// SimpleResult — Clone / ClonePtr
// =============================================================================

func Test_SimpleResult_Clone_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	c := r.Clone()

	// Act
	actual := args.Map{"valid": c.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SimpleResult Clone nil", actual)
}

func Test_SimpleResult_Clone_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("data", true, "msg")
	c := r.Clone()

	// Act
	actual := args.Map{
		"valid": c.IsValid(),
		"msg": c.Message,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"msg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult Clone valid", actual)
}

func Test_SimpleResult_ClonePtr_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult

	// Act
	actual := args.Map{"isNil": r.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SimpleResult ClonePtr nil", actual)
}

func Test_SimpleResult_ClonePtr_Valid(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("data", true, "")
	c := r.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"valid": c.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "SimpleResult ClonePtr valid", actual)
}

// =============================================================================
// Dynamic — Clone / ClonePtr / NonPtr / Ptr / Constructors
// =============================================================================

func Test_Dynamic_Clone_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)
	c := d.Clone()

	// Act
	actual := args.Map{"valid": c.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Clone", actual)
}

func Test_Dynamic_ClonePtr_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	actual := args.Map{"isNil": d.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic ClonePtr nil", actual)
}

func Test_Dynamic_ClonePtr_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	c := d.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"valid": c.IsValid(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"valid": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic ClonePtr valid", actual)
}

func Test_Dynamic_NonPtr_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	np := d.NonPtr()

	// Act
	actual := args.Map{"valid": np.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "Dynamic NonPtr", actual)
}

func Test_Dynamic_Ptr_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(42, true)
	p := d.Ptr()

	// Act
	actual := args.Map{"notNil": p != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Dynamic Ptr", actual)
}

func Test_Dynamic_InvalidDynamic_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "InvalidDynamic", actual)
}

func Test_Dynamic_NewDynamicValid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewDynamicValid", actual)
}

// =============================================================================
// DynamicStatus — constructors / clone
// =============================================================================

func Test_DynamicStatus_InvalidNoMessage_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatusNoMessage()

	// Act
	actual := args.Map{
		"valid": ds.IsValid(),
		"msg": ds.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus InvalidNoMessage", actual)
}

func Test_DynamicStatus_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("fail")

	// Act
	actual := args.Map{
		"valid": ds.IsValid(),
		"msg": ds.Message,
	}

	// Assert
	expected := args.Map{
		"valid": false,
		"msg": "fail",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus Invalid", actual)
}

func Test_DynamicStatus_Clone_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("msg")
	c := ds.Clone()

	// Act
	actual := args.Map{"msg": c.Message}

	// Assert
	expected := args.Map{"msg": "msg"}
	expected.ShouldBeEqual(t, 0, "DynamicStatus Clone", actual)
}

func Test_DynamicStatus_ClonePtr_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var ds *coredynamic.DynamicStatus

	// Act
	actual := args.Map{"isNil": ds.ClonePtr() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicStatus ClonePtr nil", actual)
}

func Test_DynamicStatus_ClonePtr_Valid(t *testing.T) {
	// Arrange
	ds := coredynamic.InvalidDynamicStatus("msg")
	c := ds.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"msg": c.Message,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"msg": "msg",
	}
	expected.ShouldBeEqual(t, 0, "DynamicStatus ClonePtr valid", actual)
}

// =============================================================================
// LeftRight — all branches
// =============================================================================

func Test_LeftRight_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"r": lr.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEmpty nil", actual)
}

func Test_LeftRight_IsEmpty_BothNil(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: nil}

	// Act
	actual := args.Map{"r": lr.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEmpty both nil", actual)
}

func Test_LeftRight_IsEmpty_HasData(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}

	// Act
	actual := args.Map{"r": lr.IsEmpty()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LeftRight IsEmpty has data", actual)
}

func Test_LeftRight_HasAnyItem(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}

	// Act
	actual := args.Map{"r": lr.HasAnyItem()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight HasAnyItem", actual)
}

func Test_LeftRight_HasLeft_True(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}

	// Act
	actual := args.Map{"r": lr.HasLeft()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight HasLeft true", actual)
}

func Test_LeftRight_HasLeft_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"r": lr.HasLeft()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LeftRight HasLeft nil", actual)
}

func Test_LeftRight_HasRight_True(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: "b"}

	// Act
	actual := args.Map{"r": lr.HasRight()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight HasRight true", actual)
}

func Test_LeftRight_HasRight_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"r": lr.HasRight()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "LeftRight HasRight nil", actual)
}

func Test_LeftRight_IsLeftEmpty_True(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: "b"}

	// Act
	actual := args.Map{"r": lr.IsLeftEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsLeftEmpty true", actual)
}

func Test_LeftRight_IsLeftEmpty_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"r": lr.IsLeftEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsLeftEmpty nil", actual)
}

func Test_LeftRight_IsRightEmpty_True(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}

	// Act
	actual := args.Map{"r": lr.IsRightEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsRightEmpty true", actual)
}

func Test_LeftRight_IsRightEmpty_Nil(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"r": lr.IsRightEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "LeftRight IsRightEmpty nil", actual)
}

func Test_LeftRight_LeftReflectSet_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	err := lr.LeftReflectSet(nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftReflectSet nil", actual)
}

func Test_LeftRight_RightReflectSet_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	err := lr.RightReflectSet(nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightReflectSet nil", actual)
}

func Test_LeftRight_DeserializeLeft_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.DeserializeLeft() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft nil", actual)
}

func Test_LeftRight_DeserializeLeft_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "data", Right: nil}
	r := lr.DeserializeLeft()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeLeft valid", actual)
}

func Test_LeftRight_DeserializeRight_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.DeserializeRight() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeRight nil", actual)
}

func Test_LeftRight_DeserializeRight_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: "data"}
	r := lr.DeserializeRight()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight DeserializeRight valid", actual)
}

func Test_LeftRight_LeftToDynamic_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.LeftToDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight LeftToDynamic nil", actual)
}

func Test_LeftRight_LeftToDynamic_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: nil}
	d := lr.LeftToDynamic()

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
	expected.ShouldBeEqual(t, 0, "LeftRight LeftToDynamic valid", actual)
}

func Test_LeftRight_RightToDynamic_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight

	// Act
	actual := args.Map{"isNil": lr.RightToDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LeftRight RightToDynamic nil", actual)
}

func Test_LeftRight_RightToDynamic_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: nil, Right: "b"}
	d := lr.RightToDynamic()

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
	expected.ShouldBeEqual(t, 0, "LeftRight RightToDynamic valid", actual)
}

func Test_LeftRight_TypeStatus_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var lr *coredynamic.LeftRight
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"notZero": ts.IsSame}

	// Assert
	expected := args.Map{"notZero": true}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus nil", actual)
}

func Test_LeftRight_TypeStatus_Valid(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: "b"}
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"same": ts.IsSame}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus valid same type", actual)
}

func Test_LeftRight_TypeStatus_DiffType(t *testing.T) {
	// Arrange
	lr := &coredynamic.LeftRight{Left: "a", Right: 42}
	ts := lr.TypeStatus()

	// Act
	actual := args.Map{"same": ts.IsSame}

	// Assert
	expected := args.Map{"same": false}
	expected.ShouldBeEqual(t, 0, "LeftRight TypeStatus diff type", actual)
}

// =============================================================================
// CastedResult — all branches
// =============================================================================

func Test_CastedResult_IsInvalid_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"r": cr.IsInvalid()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsInvalid nil", actual)
}

func Test_CastedResult_IsInvalid_Invalid(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: false}

	// Act
	actual := args.Map{"r": cr.IsInvalid()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsInvalid invalid", actual)
}

func Test_CastedResult_IsInvalid_Valid(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: true}

	// Act
	actual := args.Map{"r": cr.IsInvalid()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsInvalid valid", actual)
}

func Test_CastedResult_IsNotNull_Nil_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"r": cr.IsNotNull()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotNull nil", actual)
}

func Test_CastedResult_IsNotNull_True(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsNull: false}

	// Act
	actual := args.Map{"r": cr.IsNotNull()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotNull true", actual)
}

func Test_CastedResult_IsNotPointer_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"r": cr.IsNotPointer()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotPointer nil", actual)
}

func Test_CastedResult_IsNotPointer_True(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsPointer: false}

	// Act
	actual := args.Map{"r": cr.IsNotPointer()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotPointer true", actual)
}

func Test_CastedResult_IsNotMatchingAcceptedType_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"r": cr.IsNotMatchingAcceptedType()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotMatchingAcceptedType nil", actual)
}

func Test_CastedResult_IsNotMatchingAcceptedType_True(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsMatchingAcceptedType: false}

	// Act
	actual := args.Map{"r": cr.IsNotMatchingAcceptedType()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsNotMatchingAcceptedType true", actual)
}

func Test_CastedResult_IsSourceKind_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"r": cr.IsSourceKind(reflect.String)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult IsSourceKind nil", actual)
}

func Test_CastedResult_IsSourceKind_Match(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{SourceKind: reflect.String}

	// Act
	actual := args.Map{"r": cr.IsSourceKind(reflect.String)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult IsSourceKind match", actual)
}

func Test_CastedResult_HasError_Nil(t *testing.T) {
	// Arrange
	var cr *coredynamic.CastedResult

	// Act
	actual := args.Map{"r": cr.HasError()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult HasError nil", actual)
}

func Test_CastedResult_HasError_NoError(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{}

	// Act
	actual := args.Map{"r": cr.HasError()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult HasError no error", actual)
}

func Test_CastedResult_HasAnyIssues_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: false}

	// Act
	actual := args.Map{"r": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues invalid", actual)
}

func Test_CastedResult_HasAnyIssues_Null(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: true}

	// Act
	actual := args.Map{"r": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues null", actual)
}

func Test_CastedResult_HasAnyIssues_NotMatching(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: false}

	// Act
	actual := args.Map{"r": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues not matching", actual)
}

func Test_CastedResult_HasAnyIssues_AllGood(t *testing.T) {
	// Arrange
	cr := &coredynamic.CastedResult{IsValid: true, IsNull: false, IsMatchingAcceptedType: true}

	// Act
	actual := args.Map{"r": cr.HasAnyIssues()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "CastedResult HasAnyIssues all good", actual)
}

// =============================================================================
// BytesConverter — basic conversions
// =============================================================================

func Test_BytesConverter_SafeCastString_Empty_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter(nil)

	// Act
	actual := args.Map{"r": bc.SafeCastString()}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString empty", actual)
}

func Test_BytesConverter_SafeCastString_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))

	// Act
	actual := args.Map{"r": bc.SafeCastString()}

	// Assert
	expected := args.Map{"r": "hello"}
	expected.ShouldBeEqual(t, 0, "BytesConverter SafeCastString valid", actual)
}

func Test_BytesConverter_CastString_Empty_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter(nil)
	_, err := bc.CastString()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString empty", actual)
}

func Test_BytesConverter_CastString_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("hello"))
	s, err := bc.CastString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter CastString valid", actual)
}

func Test_BytesConverter_ToBool_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("true"))
	r, err := bc.ToBool()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": r,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBool valid", actual)
}

func Test_BytesConverter_ToBoolMust_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("false"))

	// Act
	actual := args.Map{"r": bc.ToBoolMust()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBoolMust valid", actual)
}

func Test_BytesConverter_ToString_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hello"`))
	s, err := bc.ToString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"r": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"r": "hello",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToString valid", actual)
}

func Test_BytesConverter_ToStringMust_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"hi"`))

	// Act
	actual := args.Map{"r": bc.ToStringMust()}

	// Assert
	expected := args.Map{"r": "hi"}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringMust valid", actual)
}

func Test_BytesConverter_ToStrings_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["a","b"]`))
	s, err := bc.ToStrings()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(s),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStrings valid", actual)
}

func Test_BytesConverter_ToStringsMust_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`["x"]`))
	s := bc.ToStringsMust()

	// Act
	actual := args.Map{"len": len(s)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToStringsMust valid", actual)
}

func Test_BytesConverter_ToInt64_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("42"))
	v, err := bc.ToInt64()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"v": v,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"v": int64(42),
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64 valid", actual)
}

func Test_BytesConverter_ToInt64Must_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("99"))

	// Act
	actual := args.Map{"v": bc.ToInt64Must()}

	// Assert
	expected := args.Map{"v": int64(99)}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToInt64Must valid", actual)
}

func Test_BytesConverter_Deserialize_Valid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte(`"data"`))
	var out string
	err := bc.Deserialize(&out)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": out,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "data",
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize valid", actual)
}

func Test_BytesConverter_Deserialize_Invalid(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("not json"))
	var out string
	err := bc.Deserialize(&out)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter Deserialize invalid", actual)
}

func Test_BytesConverter_ToHashmap_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	data := map[string]string{"a": "1"}
	b, _ := json.Marshal(data)
	bc := coredynamic.NewBytesConverter(b)
	hm, err := bc.ToHashmap()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": hm != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap valid", actual)
}

func Test_BytesConverter_ToHashmap_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToHashmap()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashmap invalid", actual)
}

func Test_BytesConverter_ToHashset_Valid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	data := map[string]bool{"a": true, "b": true}
	b, _ := json.Marshal(data)
	bc := coredynamic.NewBytesConverter(b)
	hs, err := bc.ToHashset()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": hs != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset valid", actual)
}

func Test_BytesConverter_ToHashset_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToHashset()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToHashset invalid", actual)
}

func Test_BytesConverter_ToCollection_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToCollection invalid", actual)
}

func Test_BytesConverter_ToSimpleSlice_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToSimpleSlice()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToSimpleSlice invalid", actual)
}

func Test_BytesConverter_ToKeyValCollection_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToKeyValCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToKeyValCollection invalid", actual)
}

func Test_BytesConverter_ToAnyCollection_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToAnyCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToAnyCollection invalid", actual)
}

func Test_BytesConverter_ToMapAnyItems_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToMapAnyItems()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToMapAnyItems invalid", actual)
}

func Test_BytesConverter_ToMapAnyItems_Valid(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(map[string]any{"Items": map[string]any{"a": 1}})
	bc := coredynamic.NewBytesConverter(b)
	m, err := bc.ToMapAnyItems()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": m.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToMapAnyItems valid", actual)
}

func Test_BytesConverter_ToDynamicCollection_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToDynamicCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToDynamicCollection invalid", actual)
}

func Test_BytesConverter_ToJsonResultCollection_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToJsonResultCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonResultCollection invalid", actual)
}

func Test_BytesConverter_ToJsonMapResults_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToJsonMapResults()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToJsonMapResults invalid", actual)
}

func Test_BytesConverter_ToBytesCollection_Invalid_FromSimpleResultInvalidN(t *testing.T) {
	// Arrange
	bc := coredynamic.NewBytesConverter([]byte("bad"))
	_, err := bc.ToBytesCollection()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "BytesConverter ToBytesCollection invalid", actual)
}
