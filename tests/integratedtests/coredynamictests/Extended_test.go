package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
)

// ==========================================
// Dynamic constructors
// ==========================================

func Test_Dynamic_InvalidDynamic(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	if d.IsValid() {
		t.Error("should be invalid")
	}
}

func Test_Dynamic_InvalidDynamicPtr(t *testing.T) {
	d := coredynamic.InvalidDynamicPtr()
	if d == nil {
		t.Error("should not be nil")
	}
	if d.IsValid() {
		t.Error("should be invalid")
	}
}

func Test_Dynamic_NewDynamicValid(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	if !d.IsValid() {
		t.Error("should be valid")
	}
	if d.Data() != "hello" {
		t.Error("data mismatch")
	}
}

func Test_Dynamic_NewDynamic(t *testing.T) {
	d := coredynamic.NewDynamic("data", true)
	if !d.IsValid() {
		t.Error("should be valid")
	}
}

func Test_Dynamic_NewDynamicPtr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("data", true)
	if d == nil {
		t.Error("should not be nil")
	}
}

// ==========================================
// Clone
// ==========================================

func Test_Dynamic_Clone_Ext(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	cloned := d.Clone()
	if cloned.Data() != "hello" {
		t.Error("clone data mismatch")
	}
}

func Test_Dynamic_ClonePtr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	cloned := d.ClonePtr()
	if cloned == nil || cloned.Data() != "hello" {
		t.Error("clonePtr data mismatch")
	}
}

func Test_Dynamic_ClonePtr_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	cloned := d.ClonePtr()
	if cloned != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Dynamic_NonPtr_Ext(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	n := d.NonPtr()
	if n.Data() != "hello" {
		t.Error("NonPtr should return same value")
	}
}

func Test_Dynamic_Ptr(t *testing.T) {
	d := coredynamic.NewDynamicPtr("hello", true)
	p := d.Ptr()
	if p == nil {
		t.Error("Ptr should return non-nil pointer")
	}
	if p.Data() != d.Data() {
		t.Error("Ptr should return pointer with same data")
	}
}

// ==========================================
// SimpleRequest
// ==========================================

func Test_SimpleRequest_InvalidNoMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequestNoMessage()
	if sr.IsValid() {
		t.Error("should be invalid")
	}
	if sr.Message() != "" {
		t.Error("should have empty message")
	}
}

func Test_SimpleRequest_Invalid(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequest("err msg")
	if sr.IsValid() {
		t.Error("should be invalid")
	}
	if sr.Message() != "err msg" {
		t.Error("message mismatch")
	}
}

func Test_SimpleRequest_New(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("data", true, "msg")
	if !sr.IsValid() {
		t.Error("should be valid")
	}
	if sr.Request() != "data" {
		t.Error("request data mismatch")
	}
	if sr.Value() != "data" {
		t.Error("value data mismatch")
	}
}

func Test_SimpleRequest_Valid(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("data")
	if !sr.IsValid() {
		t.Error("should be valid")
	}
}

func Test_SimpleRequest_IsReflectKind(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	if !sr.IsReflectKind(reflect.String) {
		t.Error("should be string kind")
	}
}

func Test_SimpleRequest_IsPointer_NonPointer(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	if sr.IsPointer() {
		t.Error("string should not be pointer")
	}
}

func Test_SimpleRequest_IsPointer_Pointer(t *testing.T) {
	val := "hello"
	sr := coredynamic.NewSimpleRequestValid(&val)
	if !sr.IsPointer() {
		t.Error("should be pointer")
	}
}

func Test_SimpleRequest_InvalidError_WithMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequest("error message")
	err := sr.InvalidError()
	if err == nil {
		t.Error("should return error")
	}
	if err.Error() != "error message" {
		t.Errorf("expected 'error message', got '%s'", err.Error())
	}
	// Second call should return cached error
	err2 := sr.InvalidError()
	if err != err2 {
		t.Error("should return cached error")
	}
}

func Test_SimpleRequest_InvalidError_EmptyMessage(t *testing.T) {
	sr := coredynamic.InvalidSimpleRequestNoMessage()
	err := sr.InvalidError()
	if err != nil {
		t.Error("empty message should return nil")
	}
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Match(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err != nil {
		t.Error("matching type should return nil")
	}
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	sr := coredynamic.NewSimpleRequestValid("hello")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err == nil {
		t.Error("mismatching type should return error")
	}
}

func Test_SimpleRequest_GetErrorOnTypeMismatch_MismatchWithMessage(t *testing.T) {
	sr := coredynamic.NewSimpleRequest("hello", true, "custom msg")
	err := sr.GetErrorOnTypeMismatch(reflect.TypeOf(0), true)
	if err == nil {
		t.Error("should return error with message")
	}
}

// ==========================================
// SimpleResult
// ==========================================

func Test_SimpleResult(t *testing.T) {
	sr := coredynamic.NewSimpleResult("data", true, "")
	if sr.Result != "data" {
		t.Error("value mismatch")
	}
	if sr.InvalidError() != nil {
		t.Error("should not have error")
	}
}

// ==========================================
// KeyVal
// ==========================================

func Test_KeyVal(t *testing.T) {
	kv := coredynamic.KeyVal{Key: "key", Value: "val"}
	if kv.Key != "key" {
		t.Error("key mismatch")
	}
	if kv.Value != "val" {
		t.Error("val mismatch")
	}
}

// ==========================================
// LeftRight
// ==========================================

func Test_LeftRight_IsLeftEmpty(t *testing.T) {
	lr := coredynamic.LeftRight{Right: "right"}
	if !lr.IsLeftEmpty() {
		t.Error("left should be empty")
	}
}

func Test_LeftRight_IsRightEmpty(t *testing.T) {
	lr := coredynamic.LeftRight{Left: "left"}
	if !lr.IsRightEmpty() {
		t.Error("right should be empty")
	}
}

// ==========================================
// TypeSameStatus
// ==========================================

func Test_TypeSameStatus(t *testing.T) {
	ts := coredynamic.TypeSameStatus("hello", "world")
	if !ts.IsSame {
		t.Error("same types should be same")
	}
}

// ==========================================
// CastTo
// ==========================================

func Test_CastTo_Match(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(""))
	if result.Error != nil {
		t.Errorf("should not error: %v", result.Error)
	}
	if !result.IsMatchingAcceptedType {
		t.Error("should match accepted type")
	}
}

func Test_CastTo_Mismatch(t *testing.T) {
	result := coredynamic.CastTo(false, "hello", reflect.TypeOf(0))
	if result.Error == nil {
		t.Error("mismatching cast should return error")
	}
}
