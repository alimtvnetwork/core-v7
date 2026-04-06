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
	d := coredynamic.InvalidDynamic()
	actual := args.Map{"result": d.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Dynamic_InvalidDynamicPtr(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	actual := args.Map{"result": d == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
	actual := args.Map{"result": d.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Dynamic_NewDynamicValid(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	actual := args.Map{"result": d.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
	actual := args.Map{"result": d.Data() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "data mismatch", actual)
}

func Test_Dynamic_NewDynamic(t *testing.T) {
	d := coredynamic.NewDynamic("data", true)
	actual := args.Map{"result": d.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_Dynamic_NewDynamicPtr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("data", true)
	actual := args.Map{"result": d == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}

// ==========================================
// Clone
// ==========================================

func Test_Dynamic_Clone_Ext(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	cloned := d.Clone()
	actual := args.Map{"result": cloned.Data() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone data mismatch", actual)
}

func Test_Dynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	cloned := d.ClonePtr()
	actual := args.Map{"result": cloned == nil || cloned.Data() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clonePtr data mismatch", actual)
}

func Test_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	cloned := d.ClonePtr()
	actual := args.Map{"result": cloned != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Dynamic_NonPtr_Ext(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	n := d.NonPtr()
	actual := args.Map{"result": n.Data() != "hello"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NonPtr should return same value", actual)
}

func Test_Dynamic_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	p := d.Ptr()
	actual := args.Map{"result": p == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Ptr should return non-nil pointer", actual)
	actual := args.Map{"result": p.Data() != d.Data()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Ptr should return pointer with same data", actual)
}

// ==========================================
// SimpleRequest
// ==========================================

func Test_SimpleRequest_InvalidNoMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequestNoMessage()
	actual := args.Map{"result": sr.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	actual := args.Map{"result": sr.Message() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have empty message", actual)
}

func Test_SimpleRequest_Invalid(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequest("err msg")
	actual := args.Map{"result": sr.IsValid()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	actual := args.Map{"result": sr.Message() != "err msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "message mismatch", actual)
}

func Test_SimpleRequest_New(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("data", true, "msg")
	actual := args.Map{"result": sr.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
	actual := args.Map{"result": sr.Request() != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "request data mismatch", actual)
	actual := args.Map{"result": sr.Value() != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "value data mismatch", actual)
}

func Test_SimpleRequest_Valid(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("data")
	actual := args.Map{"result": sr.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_SimpleRequest_IsReflectKind(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	actual := args.Map{"result": sr.IsReflectKind(reflect.String)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be string kind", actual)
}

func Test_SimpleRequest_IsPointer_NonPointer(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	actual := args.Map{"result": sr.IsPointer()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "string should not be pointer", actual)
}

func Test_SimpleRequest_IsPointer_Pointer(t *testing.T) {
	val := "hello"
	sr := coredynamic.NewSimpleRequestValid(&val)
	actual := args.Map{"result": sr.IsPointer()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be pointer", actual)
}

func Test_SimpleRequest_InvalidError_WithMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequest("error message")
	err := sr.InvalidError()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error", actual)
	actual := args.Map{"result": err.Error() != "error message"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'error message', got ''", actual)
	// Second call should return cached error
	err2 := sr.InvalidError()
	actual := args.Map{"result": err != err2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return cached error", actual)
}

func Test_SimpleRequest_InvalidError_EmptyMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequestNoMessage()
	err := sr.InvalidError()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty message should return nil", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Match(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching type should return nil", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatching type should return error", actual)
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_MismatchWithMessage(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("hello", true, "custom msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return error with message", actual)
}

// ==========================================
// SimpleResult
// ==========================================

func Test_SimpleResult(t *testing.T) {
	sr := coredynamic.NewSimpleResult("data", true, "")
	actual := args.Map{"result": sr.Result != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "value mismatch", actual)
	actual := args.Map{"result": sr.InvalidError() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have error", actual)
}

// ==========================================
// KeyVal
// ==========================================

func Test_KeyVal(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "key", Value: "val"}
	actual := args.Map{"result": kv.Key != "key"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "key mismatch", actual)
	actual := args.Map{"result": kv.Value != "val"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "val mismatch", actual)
}

// ==========================================
// LeftRight
// ==========================================

func Test_LeftRight_IsLeftEmpty(t *testing.T) {
	lr := coredynamic.LeftRight{Right: "right"}
	actual := args.Map{"result": lr.IsLeftEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "left should be empty", actual)
}

func Test_LeftRight_IsRightEmpty(t *testing.T) {
	lr := coredynamic.LeftRight{Left: "left"}
	actual := args.Map{"result": lr.IsRightEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "right should be empty", actual)
}

// ==========================================
// TypeSameStatus
// ==========================================

func Test_TypeSameStatus(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", "world")
	actual := args.Map{"result": ts.IsSame}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same types should be same", actual)
}

// ==========================================
// CastTo
// ==========================================

func Test_CastTo_Match(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	actual := args.Map{"result": result.Error != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not error:", actual)
	actual := args.Map{"result": result.IsMatchingAcceptedType}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match accepted type", actual)
}

func Test_CastTo_Mismatch(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))
	actual := args.Map{"result": result.Error == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatching cast should return error", actual)
}
