package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Generic If
// ==========================================

func Test_If_Int_True(t *testing.T) {
	r := conditional.If[int](true, 2, 7)
	actual := args.Map{"result": r != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_If_Int_False(t *testing.T) {
	r := conditional.If[int](false, 2, 7)
	actual := args.Map{"result": r != 7}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_If_String_True(t *testing.T) {
	r := conditional.If[string](true, "yes", "no")
	actual := args.Map{"result": r != "yes"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'yes', got ''", actual)
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
	actual := args.Map{"result": r != "true"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'true', got ''", actual)
}

func Test_IfFunc_False(t *testing.T) {
	r := conditional.IfFunc[string](
		false,
		func() string { return "true" },
		func() string { return "false" },
	)
	actual := args.Map{"result": r != "false"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'false', got ''", actual)
}

// ==========================================
// IfTrueFunc
// ==========================================

func Test_IfTrueFunc_True(t *testing.T) {
	r := conditional.IfTrueFunc[int](true, func() int { return 42 })
	actual := args.Map{"result": r != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_IfTrueFunc_False(t *testing.T) {
	r := conditional.IfTrueFunc[int](false, func() int { return 42 })
	actual := args.Map{"result": r != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// IfSlice
// ==========================================

func Test_IfSlice_True(t *testing.T) {
	r := conditional.IfSlice[int](true, []int{1, 2}, []int{3})
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_IfSlice_False(t *testing.T) {
	r := conditional.IfSlice[int](false, []int{1, 2}, []int{3})
	actual := args.Map{"result": len(r) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ==========================================
// NilDef / NilDefPtr
// ==========================================

func Test_NilDef_Nil(t *testing.T) {
	r := conditional.NilDef[int](nil, 99)
	actual := args.Map{"result": r != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_NilDef_NonNil(t *testing.T) {
	v := 42
	r := conditional.NilDef[int](&v, 99)
	actual := args.Map{"result": r != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_NilDefPtr_Nil(t *testing.T) {
	r := conditional.NilDefPtr[int](nil, 99)
	actual := args.Map{"result": *r != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_NilDefPtr_NonNil(t *testing.T) {
	v := 42
	r := conditional.NilDefPtr[int](&v, 99)
	actual := args.Map{"result": *r != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

// ==========================================
// NilVal / NilValPtr
// ==========================================

func Test_NilVal_Nil(t *testing.T) {
	r := conditional.NilVal[string](nil, "nil", "nonnil")
	actual := args.Map{"result": r != "nil"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nil', got ''", actual)
}

func Test_NilVal_NonNil(t *testing.T) {
	v := "hello"
	r := conditional.NilVal[string](&v, "nil", "nonnil")
	actual := args.Map{"result": r != "nonnil"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nonnil', got ''", actual)
}

func Test_NilValPtr_Nil(t *testing.T) {
	r := conditional.NilValPtr[string](nil, "nil", "nonnil")
	actual := args.Map{"result": *r != "nil"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nil', got ''", actual)
}

func Test_NilValPtr_NonNil(t *testing.T) {
	v := "hello"
	r := conditional.NilValPtr[string](&v, "nil", "nonnil")
	actual := args.Map{"result": *r != "nonnil"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nonnil', got ''", actual)
}

// ==========================================
// ValueOrZero / PtrOrZero
// ==========================================

func Test_ValueOrZero_Nil(t *testing.T) {
	r := conditional.ValueOrZero[int](nil)
	actual := args.Map{"result": r != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_ValueOrZero_NonNil(t *testing.T) {
	v := 42
	r := conditional.ValueOrZero[int](&v)
	actual := args.Map{"result": r != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_PtrOrZero_Nil(t *testing.T) {
	r := conditional.PtrOrZero[int](nil)
	actual := args.Map{"result": r == nil || *r != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pointer to 0", actual)
}

func Test_PtrOrZero_NonNil(t *testing.T) {
	v := 42
	r := conditional.PtrOrZero[int](&v)
	actual := args.Map{"result": *r != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

// ==========================================
// IfPtr
// ==========================================

func Test_IfPtr_True(t *testing.T) {
	a, b := 1, 2
	r := conditional.IfPtr[int](true, &a, &b)
	actual := args.Map{"result": *r != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_IfPtr_False(t *testing.T) {
	a, b := 1, 2
	r := conditional.IfPtr[int](false, &a, &b)
	actual := args.Map{"result": *r != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ==========================================
// NilCheck (deprecated but still needs coverage)
// ==========================================

func Test_NilCheck_Nil(t *testing.T) {
	r := conditional.NilCheck(nil, "onNil", "onNonNil")
	actual := args.Map{"result": r != "onNil"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'onNil', got ''", actual)
}

func Test_NilCheck_NonNil(t *testing.T) {
	r := conditional.NilCheck("val", "onNil", "onNonNil")
	actual := args.Map{"result": r != "onNonNil"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'onNonNil', got ''", actual)
}

// ==========================================
// DefOnNil
// ==========================================

func Test_DefOnNil_Nil(t *testing.T) {
	r := conditional.DefOnNil(nil, "default")
	actual := args.Map{"result": r != "default"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'default', got ''", actual)
}

func Test_DefOnNil_NonNil(t *testing.T) {
	r := conditional.DefOnNil("value", "default")
	actual := args.Map{"result": r != "value"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
}

// ==========================================
// NilOrEmptyStr / NilOrEmptyStrPtr
// ==========================================

func Test_NilOrEmptyStr_Nil(t *testing.T) {
	r := conditional.NilOrEmptyStr(nil, "empty", "notempty")
	actual := args.Map{"result": r != "empty"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'empty', got ''", actual)
}

func Test_NilOrEmptyStr_Empty(t *testing.T) {
	s := ""
	r := conditional.NilOrEmptyStr(&s, "empty", "notempty")
	actual := args.Map{"result": r != "empty"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'empty', got ''", actual)
}

func Test_NilOrEmptyStr_NonEmpty(t *testing.T) {
	s := "hello"
	r := conditional.NilOrEmptyStr(&s, "empty", "notempty")
	actual := args.Map{"result": r != "notempty"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'notempty', got ''", actual)
}

func Test_NilOrEmptyStrPtr_Nil(t *testing.T) {
	r := conditional.NilOrEmptyStrPtr(nil, "empty", "notempty")
	actual := args.Map{"result": *r != "empty"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'empty', got ''", actual)
}

func Test_NilOrEmptyStrPtr_NonEmpty(t *testing.T) {
	s := "hello"
	r := conditional.NilOrEmptyStrPtr(&s, "empty", "notempty")
	actual := args.Map{"result": *r != "notempty"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'notempty', got ''", actual)
}

// ==========================================
// StringDefault
// ==========================================

func Test_StringDefault_True(t *testing.T) {
	r := conditional.StringDefault(true, "value")
	actual := args.Map{"result": r != "value"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
}

func Test_StringDefault_False(t *testing.T) {
	r := conditional.StringDefault(false, "value")
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty, got ''", actual)
}

// ==========================================
// BoolByOrder
// ==========================================

func Test_BoolByOrder_FirstTrue(t *testing.T) {
	r := conditional.BoolByOrder(true, false)
	actual := args.Map{"result": r}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true", actual)
}

func Test_BoolByOrder_AllFalse(t *testing.T) {
	r := conditional.BoolByOrder(false, false, false)
	actual := args.Map{"result": r}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_BoolByOrder_LastTrue(t *testing.T) {
	r := conditional.BoolByOrder(false, false, true)
	actual := args.Map{"result": r}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true", actual)
}

func Test_BoolByOrder_Empty(t *testing.T) {
	r := conditional.BoolByOrder()
	actual := args.Map{"result": r}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return false", actual)
}

// ==========================================
// Func
// ==========================================

func Test_Func_True(t *testing.T) {
	trueF := func() any { return "true" }
	falseF := func() any { return "false" }
	r := conditional.Func(true, trueF, falseF)
	actual := args.Map{"result": r() != "true"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return true func", actual)
}

func Test_Func_False(t *testing.T) {
	trueF := func() any { return "true" }
	falseF := func() any { return "false" }
	r := conditional.Func(false, trueF, falseF)
	actual := args.Map{"result": r() != "false"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false func", actual)
}

// ==========================================
// StringsIndexVal
// ==========================================

func Test_StringsIndexVal_True(t *testing.T) {
	r := conditional.StringsIndexVal(true, []string{"a", "b", "c"}, 0, 2)
	actual := args.Map{"result": r != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
}

func Test_StringsIndexVal_False(t *testing.T) {
	r := conditional.StringsIndexVal(false, []string{"a", "b", "c"}, 0, 2)
	actual := args.Map{"result": r != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'c', got ''", actual)
}
