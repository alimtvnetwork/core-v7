package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
)

// ==========================================
// Generic If
// ==========================================

func Test_If_Int_True(t *testing.T) {
	r := conditional.If[int](true, 2, 7)
	if r != 2 {
		t.Errorf("expected 2, got %d", r)
	}
}

func Test_If_Int_False(t *testing.T) {
	r := conditional.If[int](false, 2, 7)
	if r != 7 {
		t.Errorf("expected 7, got %d", r)
	}
}

func Test_If_String_True(t *testing.T) {
	r := conditional.If[string](true, "yes", "no")
	if r != "yes" {
		t.Errorf("expected 'yes', got '%s'", r)
	}
}

// ==========================================
// IfFunc
// ==========================================

func Test_IfFunc_True(t *testing.T) {
	r := conditional.IfFunc[string](
		true,
		func() string { return "true" },
		func() string { return "false" },
	)
	if r != "true" {
		t.Errorf("expected 'true', got '%s'", r)
	}
}

func Test_IfFunc_False(t *testing.T) {
	r := conditional.IfFunc[string](
		false,
		func() string { return "true" },
		func() string { return "false" },
	)
	if r != "false" {
		t.Errorf("expected 'false', got '%s'", r)
	}
}

// ==========================================
// IfTrueFunc
// ==========================================

func Test_IfTrueFunc_True(t *testing.T) {
	r := conditional.IfTrueFunc[int](true, func() int { return 42 })
	if r != 42 {
		t.Errorf("expected 42, got %d", r)
	}
}

func Test_IfTrueFunc_False(t *testing.T) {
	r := conditional.IfTrueFunc[int](false, func() int { return 42 })
	if r != 0 {
		t.Errorf("expected 0, got %d", r)
	}
}

// ==========================================
// IfSlice
// ==========================================

func Test_IfSlice_True(t *testing.T) {
	r := conditional.IfSlice[int](true, []int{1, 2}, []int{3})
	if len(r) != 2 {
		t.Errorf("expected 2, got %d", len(r))
	}
}

func Test_IfSlice_False(t *testing.T) {
	r := conditional.IfSlice[int](false, []int{1, 2}, []int{3})
	if len(r) != 1 {
		t.Errorf("expected 1, got %d", len(r))
	}
}

// ==========================================
// NilDef / NilDefPtr
// ==========================================

func Test_NilDef_Nil(t *testing.T) {
	r := conditional.NilDef[int](nil, 99)
	if r != 99 {
		t.Errorf("expected 99, got %d", r)
	}
}

func Test_NilDef_NonNil(t *testing.T) {
	v := 42
	r := conditional.NilDef[int](&v, 99)
	if r != 42 {
		t.Errorf("expected 42, got %d", r)
	}
}

func Test_NilDefPtr_Nil(t *testing.T) {
	r := conditional.NilDefPtr[int](nil, 99)
	if *r != 99 {
		t.Errorf("expected 99, got %d", *r)
	}
}

func Test_NilDefPtr_NonNil(t *testing.T) {
	v := 42
	r := conditional.NilDefPtr[int](&v, 99)
	if *r != 42 {
		t.Errorf("expected 42, got %d", *r)
	}
}

// ==========================================
// NilVal / NilValPtr
// ==========================================

func Test_NilVal_Nil(t *testing.T) {
	r := conditional.NilVal[string](nil, "nil", "nonnil")
	if r != "nil" {
		t.Errorf("expected 'nil', got '%s'", r)
	}
}

func Test_NilVal_NonNil(t *testing.T) {
	v := "hello"
	r := conditional.NilVal[string](&v, "nil", "nonnil")
	if r != "nonnil" {
		t.Errorf("expected 'nonnil', got '%s'", r)
	}
}

func Test_NilValPtr_Nil(t *testing.T) {
	r := conditional.NilValPtr[string](nil, "nil", "nonnil")
	if *r != "nil" {
		t.Errorf("expected 'nil', got '%s'", *r)
	}
}

func Test_NilValPtr_NonNil(t *testing.T) {
	v := "hello"
	r := conditional.NilValPtr[string](&v, "nil", "nonnil")
	if *r != "nonnil" {
		t.Errorf("expected 'nonnil', got '%s'", *r)
	}
}

// ==========================================
// ValueOrZero / PtrOrZero
// ==========================================

func Test_ValueOrZero_Nil(t *testing.T) {
	r := conditional.ValueOrZero[int](nil)
	if r != 0 {
		t.Errorf("expected 0, got %d", r)
	}
}

func Test_ValueOrZero_NonNil(t *testing.T) {
	v := 42
	r := conditional.ValueOrZero[int](&v)
	if r != 42 {
		t.Errorf("expected 42, got %d", r)
	}
}

func Test_PtrOrZero_Nil(t *testing.T) {
	r := conditional.PtrOrZero[int](nil)
	if r == nil || *r != 0 {
		t.Error("expected pointer to 0")
	}
}

func Test_PtrOrZero_NonNil(t *testing.T) {
	v := 42
	r := conditional.PtrOrZero[int](&v)
	if *r != 42 {
		t.Errorf("expected 42, got %d", *r)
	}
}

// ==========================================
// IfPtr
// ==========================================

func Test_IfPtr_True(t *testing.T) {
	a, b := 1, 2
	r := conditional.IfPtr[int](true, &a, &b)
	if *r != 1 {
		t.Errorf("expected 1, got %d", *r)
	}
}

func Test_IfPtr_False(t *testing.T) {
	a, b := 1, 2
	r := conditional.IfPtr[int](false, &a, &b)
	if *r != 2 {
		t.Errorf("expected 2, got %d", *r)
	}
}

// ==========================================
// NilCheck (deprecated but still needs coverage)
// ==========================================

func Test_NilCheck_Nil(t *testing.T) {
	r := conditional.NilCheck(nil, "onNil", "onNonNil")
	if r != "onNil" {
		t.Errorf("expected 'onNil', got '%v'", r)
	}
}

func Test_NilCheck_NonNil(t *testing.T) {
	r := conditional.NilCheck("val", "onNil", "onNonNil")
	if r != "onNonNil" {
		t.Errorf("expected 'onNonNil', got '%v'", r)
	}
}

// ==========================================
// DefOnNil
// ==========================================

func Test_DefOnNil_Nil(t *testing.T) {
	r := conditional.DefOnNil(nil, "default")
	if r != "default" {
		t.Errorf("expected 'default', got '%v'", r)
	}
}

func Test_DefOnNil_NonNil(t *testing.T) {
	r := conditional.DefOnNil("value", "default")
	if r != "value" {
		t.Errorf("expected 'value', got '%v'", r)
	}
}

// ==========================================
// NilOrEmptyStr / NilOrEmptyStrPtr
// ==========================================

func Test_NilOrEmptyStr_Nil(t *testing.T) {
	r := conditional.NilOrEmptyStr(nil, "empty", "notempty")
	if r != "empty" {
		t.Errorf("expected 'empty', got '%s'", r)
	}
}

func Test_NilOrEmptyStr_Empty(t *testing.T) {
	s := ""
	r := conditional.NilOrEmptyStr(&s, "empty", "notempty")
	if r != "empty" {
		t.Errorf("expected 'empty', got '%s'", r)
	}
}

func Test_NilOrEmptyStr_NonEmpty(t *testing.T) {
	s := "hello"
	r := conditional.NilOrEmptyStr(&s, "empty", "notempty")
	if r != "notempty" {
		t.Errorf("expected 'notempty', got '%s'", r)
	}
}

func Test_NilOrEmptyStrPtr_Nil(t *testing.T) {
	r := conditional.NilOrEmptyStrPtr(nil, "empty", "notempty")
	if *r != "empty" {
		t.Errorf("expected 'empty', got '%s'", *r)
	}
}

func Test_NilOrEmptyStrPtr_NonEmpty(t *testing.T) {
	s := "hello"
	r := conditional.NilOrEmptyStrPtr(&s, "empty", "notempty")
	if *r != "notempty" {
		t.Errorf("expected 'notempty', got '%s'", *r)
	}
}

// ==========================================
// StringDefault
// ==========================================

func Test_StringDefault_True(t *testing.T) {
	r := conditional.StringDefault(true, "value")
	if r != "value" {
		t.Errorf("expected 'value', got '%s'", r)
	}
}

func Test_StringDefault_False(t *testing.T) {
	r := conditional.StringDefault(false, "value")
	if r != "" {
		t.Errorf("expected empty, got '%s'", r)
	}
}

// ==========================================
// BoolByOrder
// ==========================================

func Test_BoolByOrder_FirstTrue(t *testing.T) {
	r := conditional.BoolByOrder(true, false)
	if !r {
		t.Error("should return true")
	}
}

func Test_BoolByOrder_AllFalse(t *testing.T) {
	r := conditional.BoolByOrder(false, false, false)
	if r {
		t.Error("should return false")
	}
}

func Test_BoolByOrder_LastTrue(t *testing.T) {
	r := conditional.BoolByOrder(false, false, true)
	if !r {
		t.Error("should return true")
	}
}

func Test_BoolByOrder_Empty(t *testing.T) {
	r := conditional.BoolByOrder()
	if r {
		t.Error("empty should return false")
	}
}

// ==========================================
// Func
// ==========================================

func Test_Func_True(t *testing.T) {
	trueF := func() any { return "true" }
	falseF := func() any { return "false" }
	r := conditional.Func(true, trueF, falseF)
	if r() != "true" {
		t.Error("should return true func")
	}
}

func Test_Func_False(t *testing.T) {
	trueF := func() any { return "true" }
	falseF := func() any { return "false" }
	r := conditional.Func(false, trueF, falseF)
	if r() != "false" {
		t.Error("should return false func")
	}
}

// ==========================================
// StringsIndexVal
// ==========================================

func Test_StringsIndexVal_True(t *testing.T) {
	r := conditional.StringsIndexVal(true, []string{"a", "b", "c"}, 0, 2)
	if r != "a" {
		t.Errorf("expected 'a', got '%s'", r)
	}
}

func Test_StringsIndexVal_False(t *testing.T) {
	r := conditional.StringsIndexVal(false, []string{"a", "b", "c"}, 0, 2)
	if r != "c" {
		t.Errorf("expected 'c', got '%s'", r)
	}
}
