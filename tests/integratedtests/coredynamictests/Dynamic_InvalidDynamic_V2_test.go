package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Dynamic constructors
// ==========================================

func Test_Dynamic_InvalidDynamic(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()

	// Act
	actual := args.Map{"result": d.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Dynamic_InvalidDynamicPtr(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamicPtr()

	// Act
	actual := args.Map{"result": d == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual = args.Map{"result": d.IsValid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Dynamic_NewDynamicValid(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{"result": d.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
	actual = args.Map{"result": d.Data() != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "data mismatch", actual)
}

func Test_Dynamic_NewDynamic(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("data", true)

	// Act
	actual := args.Map{"result": d.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_Dynamic_NewDynamicPtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("data", true)

	// Act
	actual := args.Map{"result": d == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ==========================================
// Clone
// ==========================================

func Test_Dynamic_Clone_Ext(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	cloned := d.Clone()

	// Act
	actual := args.Map{"result": cloned.Data() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone data mismatch", actual)
}

func Test_Dynamic_ClonePtr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	cloned := d.ClonePtr()

	// Act
	actual := args.Map{"result": cloned == nil || cloned.Data() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clonePtr data mismatch", actual)
}

func Test_Dynamic_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	cloned := d.ClonePtr()

	// Act
	actual := args.Map{"result": cloned != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Dynamic_NonPtr_Ext(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")
	n := d.NonPtr()

	// Act
	actual := args.Map{"result": n.Data() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NonPtr should return same value", actual)
}

func Test_Dynamic_Ptr(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicPtr("hello", true)
	p := d.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Ptr should return non-nil pointer", actual)
	actual = args.Map{"result": p.Data() != d.Data()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Ptr should return pointer with same data", actual)
}

// ==========================================
// SimpleRequest
// ==========================================

func Test_SimpleRequest_InvalidNoMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidSimpleRequestNoMessage()

	// Act
	actual := args.Map{"result": sr.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	actual = args.Map{"result": sr.Message() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have empty message", actual)
}

func Test_SimpleRequest_Invalid(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidSimpleRequest("err msg")

	// Act
	actual := args.Map{"result": sr.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	actual = args.Map{"result": sr.Message() != "err msg"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "message mismatch", actual)
}

func Test_SimpleRequest_New(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequest("data", true, "msg")

	// Act
	actual := args.Map{"result": sr.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
	actual = args.Map{"result": sr.Request() != "data"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "request data mismatch", actual)
	actual = args.Map{"result": sr.Value() != "data"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "value data mismatch", actual)
}

func Test_SimpleRequest_Valid(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("data")

	// Act
	actual := args.Map{"result": sr.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_SimpleRequest_IsReflectKind(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")

	// Act
	actual := args.Map{"result": sr.IsReflectKind(reflect.String)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be string kind", actual)
}

func Test_SimpleRequest_IsPointer_NonPointer(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")

	// Act
	actual := args.Map{"result": sr.IsPointer()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be pointer", actual)
}

func Test_SimpleRequest_IsPointer_Pointer(t *testing.T) {
	// Arrange
	val := "hello"
	sr := coredynamic.NewSimpleRequestValid(&val)

	// Act
	actual := args.Map{"result": sr.IsPointer()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be pointer", actual)
}

func Test_SimpleRequest_InvalidError_WithMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidSimpleRequest("error message")
	err := sr.InvalidError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
	actual = args.Map{"result": err.Error() != "error message"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'error message', got ''", actual)
	// Second call should return cached error
	err2 := sr.InvalidError()
	actual = args.Map{"result": err != err2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return cached error", actual)
}

func Test_SimpleRequest_InvalidError_EmptyMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.InvalidSimpleRequestNoMessage()
	err := sr.InvalidError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty message should return nil", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Match(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching type should return nil", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatching type should return error", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_MismatchWithMessage(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleRequest("hello", true, "custom msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error with message", actual)
}

// ==========================================
// SimpleResult
// ==========================================

func Test_SimpleResult(t *testing.T) {
	// Arrange
	sr := coredynamic.NewSimpleResult("data", true, "")

	// Act
	actual := args.Map{"result": sr.Result != "data"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "value mismatch", actual)
	actual = args.Map{"result": sr.InvalidError() != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error", actual)
}

// ==========================================
// KeyVal
// ==========================================

func Test_KeyVal(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "key", Value: "val"}

	// Act
	actual := args.Map{"result": kv.Key != "key"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "key mismatch", actual)
	actual = args.Map{"result": kv.Value != "val"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "val mismatch", actual)
}

// ==========================================
// LeftRight
// ==========================================

func Test_LeftRight_IsLeftEmpty(t *testing.T) {
	// Arrange
	lr := coredynamic.LeftRight{Right: "right"}

	// Act
	actual := args.Map{"result": lr.IsLeftEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left should be empty", actual)
}

func Test_LeftRight_IsRightEmpty(t *testing.T) {
	// Arrange
	lr := coredynamic.LeftRight{Left: "left"}

	// Act
	actual := args.Map{"result": lr.IsRightEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "right should be empty", actual)
}

// ==========================================
// TypeSameStatus
// ==========================================

func Test_TypeSameStatus(t *testing.T) {
	// Arrange
	ts := coredynamic.TypeSameStatus("hello", "world")

	// Act
	actual := args.Map{"result": ts.IsSame}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same types should be same", actual)
}

// ==========================================
// CastTo
// ==========================================

func Test_CastTo_Match(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))

	// Act
	actual := args.Map{"result": result.Error != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not error:", actual)
	actual = args.Map{"result": result.IsMatchingAcceptedType}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match accepted type", actual)
}

func Test_CastTo_Mismatch(t *testing.T) {
	// Arrange
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))

	// Act
	actual := args.Map{"result": result.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatching cast should return error", actual)
}
