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
	// Arrange
	r := conditional.If[int](true, 2, 7)

	// Act
	actual := args.Map{"result": r != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_If_Int_False(t *testing.T) {
	// Arrange
	r := conditional.If[int](false, 2, 7)

	// Act
	actual := args.Map{"result": r != 7}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 7", actual)
}

func Test_If_String_True(t *testing.T) {
	// Arrange
	r := conditional.If[string](true, "yes", "no")

	// Act
	actual := args.Map{"result": r != "yes"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'yes', got ''", actual)
}

// ==========================================
// IfFunc
// ==========================================

func Test_IfFunc_True(t *testing.T) {
	// Arrange
	r := conditional.IfFunc[string](
		true,
		func() string { return "true" },
		func() string { return "false" },
	)

	// Act
	actual := args.Map{"result": r != "true"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'true', got ''", actual)
}

func Test_IfFunc_False(t *testing.T) {
	// Arrange
	r := conditional.IfFunc[string](
		false,
		func() string { return "true" },
		func() string { return "false" },
	)

	// Act
	actual := args.Map{"result": r != "false"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'false', got ''", actual)
}

// ==========================================
// IfTrueFunc
// ==========================================

func Test_IfTrueFunc_True(t *testing.T) {
	// Arrange
	r := conditional.IfTrueFunc[int](true, func() int { return 42 })

	// Act
	actual := args.Map{"result": r != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_IfTrueFunc_False(t *testing.T) {
	// Arrange
	r := conditional.IfTrueFunc[int](false, func() int { return 42 })

	// Act
	actual := args.Map{"result": r != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ==========================================
// IfSlice
// ==========================================

func Test_IfSlice_True(t *testing.T) {
	// Arrange
	r := conditional.IfSlice[int](true, []int{1, 2}, []int{3})

	// Act
	actual := args.Map{"result": len(r) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_IfSlice_False(t *testing.T) {
	// Arrange
	r := conditional.IfSlice[int](false, []int{1, 2}, []int{3})

	// Act
	actual := args.Map{"result": len(r) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ==========================================
// NilDef / NilDefPtr
// ==========================================

func Test_NilDef_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilDef[int](nil, 99)

	// Act
	actual := args.Map{"result": r != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_NilDef_NonNil(t *testing.T) {
	// Arrange
	v := 42
	r := conditional.NilDef[int](&v, 99)

	// Act
	actual := args.Map{"result": r != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_NilDefPtr_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilDefPtr[int](nil, 99)

	// Act
	actual := args.Map{"result": *r != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
}

func Test_NilDefPtr_NonNil(t *testing.T) {
	// Arrange
	v := 42
	r := conditional.NilDefPtr[int](&v, 99)

	// Act
	actual := args.Map{"result": *r != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

// ==========================================
// NilVal / NilValPtr
// ==========================================

func Test_NilVal_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilVal[string](nil, "nil", "nonnil")

	// Act
	actual := args.Map{"result": r != "nil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nil', got ''", actual)
}

func Test_NilVal_NonNil(t *testing.T) {
	// Arrange
	v := "hello"
	r := conditional.NilVal[string](&v, "nil", "nonnil")

	// Act
	actual := args.Map{"result": r != "nonnil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nonnil', got ''", actual)
}

func Test_NilValPtr_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilValPtr[string](nil, "nil", "nonnil")

	// Act
	actual := args.Map{"result": *r != "nil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nil', got ''", actual)
}

func Test_NilValPtr_NonNil(t *testing.T) {
	// Arrange
	v := "hello"
	r := conditional.NilValPtr[string](&v, "nil", "nonnil")

	// Act
	actual := args.Map{"result": *r != "nonnil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'nonnil', got ''", actual)
}

// ==========================================
// ValueOrZero / PtrOrZero
// ==========================================

func Test_ValueOrZero_Nil(t *testing.T) {
	// Arrange
	r := conditional.ValueOrZero[int](nil)

	// Act
	actual := args.Map{"result": r != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_ValueOrZero_NonNil(t *testing.T) {
	// Arrange
	v := 42
	r := conditional.ValueOrZero[int](&v)

	// Act
	actual := args.Map{"result": r != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_PtrOrZero_Nil(t *testing.T) {
	// Arrange
	r := conditional.PtrOrZero[int](nil)

	// Act
	actual := args.Map{"result": r == nil || *r != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pointer to 0", actual)
}

func Test_PtrOrZero_NonNil(t *testing.T) {
	// Arrange
	v := 42
	r := conditional.PtrOrZero[int](&v)

	// Act
	actual := args.Map{"result": *r != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

// ==========================================
// IfPtr
// ==========================================

func Test_IfPtr_True(t *testing.T) {
	// Arrange
	a, b := 1, 2
	r := conditional.IfPtr[int](true, &a, &b)

	// Act
	actual := args.Map{"result": *r != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_IfPtr_False(t *testing.T) {
	// Arrange
	a, b := 1, 2
	r := conditional.IfPtr[int](false, &a, &b)

	// Act
	actual := args.Map{"result": *r != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

// ==========================================
// NilCheck (deprecated but still needs coverage)
// ==========================================

func Test_NilCheck_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilCheck(nil, "onNil", "onNonNil")

	// Act
	actual := args.Map{"result": r != "onNil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'onNil', got ''", actual)
}

func Test_NilCheck_NonNil(t *testing.T) {
	// Arrange
	r := conditional.NilCheck("val", "onNil", "onNonNil")

	// Act
	actual := args.Map{"result": r != "onNonNil"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'onNonNil', got ''", actual)
}

// ==========================================
// DefOnNil
// ==========================================

func Test_DefOnNil_Nil(t *testing.T) {
	// Arrange
	r := conditional.DefOnNil(nil, "default")

	// Act
	actual := args.Map{"result": r != "default"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'default', got ''", actual)
}

func Test_DefOnNil_NonNil(t *testing.T) {
	// Arrange
	r := conditional.DefOnNil("value", "default")

	// Act
	actual := args.Map{"result": r != "value"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
}

// ==========================================
// NilOrEmptyStr / NilOrEmptyStrPtr
// ==========================================

func Test_NilOrEmptyStr_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilOrEmptyStr(nil, "empty", "notempty")

	// Act
	actual := args.Map{"result": r != "empty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'empty', got ''", actual)
}

func Test_NilOrEmptyStr_Empty(t *testing.T) {
	// Arrange
	s := ""
	r := conditional.NilOrEmptyStr(&s, "empty", "notempty")

	// Act
	actual := args.Map{"result": r != "empty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'empty', got ''", actual)
}

func Test_NilOrEmptyStr_NonEmpty(t *testing.T) {
	// Arrange
	s := "hello"
	r := conditional.NilOrEmptyStr(&s, "empty", "notempty")

	// Act
	actual := args.Map{"result": r != "notempty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'notempty', got ''", actual)
}

func Test_NilOrEmptyStrPtr_Nil(t *testing.T) {
	// Arrange
	r := conditional.NilOrEmptyStrPtr(nil, "empty", "notempty")

	// Act
	actual := args.Map{"result": *r != "empty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'empty', got ''", actual)
}

func Test_NilOrEmptyStrPtr_NonEmpty(t *testing.T) {
	// Arrange
	s := "hello"
	r := conditional.NilOrEmptyStrPtr(&s, "empty", "notempty")

	// Act
	actual := args.Map{"result": *r != "notempty"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'notempty', got ''", actual)
}

// ==========================================
// StringDefault
// ==========================================

func Test_StringDefault_True(t *testing.T) {
	// Arrange
	r := conditional.StringDefault(true, "value")

	// Act
	actual := args.Map{"result": r != "value"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'value', got ''", actual)
}

func Test_StringDefault_False(t *testing.T) {
	// Arrange
	r := conditional.StringDefault(false, "value")

	// Act
	actual := args.Map{"result": r != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty, got ''", actual)
}

// ==========================================
// BoolByOrder
// ==========================================

func Test_BoolByOrder_FirstTrue_FromIfInt(t *testing.T) {
	// Arrange
	r := conditional.BoolByOrder(true, false)

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true", actual)
}

func Test_BoolByOrder_AllFalse_FromIfInt(t *testing.T) {
	// Arrange
	r := conditional.BoolByOrder(false, false, false)

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_BoolByOrder_LastTrue_FromIfInt(t *testing.T) {
	// Arrange
	r := conditional.BoolByOrder(false, false, true)

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true", actual)
}

func Test_BoolByOrder_Empty(t *testing.T) {
	// Arrange
	r := conditional.BoolByOrder()

	// Act
	actual := args.Map{"result": r}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return false", actual)
}

// ==========================================
// Func
// ==========================================

func Test_Func_True(t *testing.T) {
	// Arrange
	trueF := func() any { return "true" }
	falseF := func() any { return "false" }
	r := conditional.Func(true, trueF, falseF)

	// Act
	actual := args.Map{"result": r() != "true"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return true func", actual)
}

func Test_Func_False(t *testing.T) {
	// Arrange
	trueF := func() any { return "true" }
	falseF := func() any { return "false" }
	r := conditional.Func(false, trueF, falseF)

	// Act
	actual := args.Map{"result": r() != "false"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false func", actual)
}

// ==========================================
// StringsIndexVal
// ==========================================

func Test_StringsIndexVal_True_FromIfInt(t *testing.T) {
	// Arrange
	r := conditional.StringsIndexVal(true, []string{"a", "b", "c"}, 0, 2)

	// Act
	actual := args.Map{"result": r != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'a', got ''", actual)
}

func Test_StringsIndexVal_False_FromIfInt(t *testing.T) {
	// Arrange
	r := conditional.StringsIndexVal(false, []string{"a", "b", "c"}, 0, 2)

	// Act
	actual := args.Map{"result": r != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'c', got ''", actual)
}
