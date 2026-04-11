package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── NilCheck ──

func Test_NilCheck_Nil_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.NilCheck(nil, "default", "nonnil")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "default"}
	expected.ShouldBeEqual(t, 0, "NilCheck nil -- default", actual)
}

func Test_NilCheck_NonNil_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.NilCheck("val", "default", "nonnil")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "nonnil"}
	expected.ShouldBeEqual(t, 0, "NilCheck nonnil -- nonnil", actual)
}

// ── DefOnNil ──

func Test_DefOnNil_Nil_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.DefOnNil(nil, "fallback")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "fallback"}
	expected.ShouldBeEqual(t, 0, "DefOnNil nil -- fallback", actual)
}

func Test_DefOnNil_NonNil_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.DefOnNil("actual", "fallback")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "actual"}
	expected.ShouldBeEqual(t, 0, "DefOnNil nonnil -- actual", actual)
}

// ── NilOrEmptyStr ──

func Test_NilOrEmptyStr_Nil_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.NilOrEmptyStr(nil, "empty", "notempty")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns nil -- nil", actual)
}

func Test_NilOrEmptyStr_Empty_FromNilCheckNil(t *testing.T) {
	// Arrange
	s := ""
	result := conditional.NilOrEmptyStr(&s, "empty", "notempty")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns nil -- empty string", actual)
}

func Test_NilOrEmptyStr_NonEmpty_FromNilCheckNil(t *testing.T) {
	// Arrange
	s := "hello"
	result := conditional.NilOrEmptyStr(&s, "empty", "notempty")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "notempty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns nil -- non-empty", actual)
}

// ── NilOrEmptyStrPtr ──

func Test_NilOrEmptyStrPtr_Nil_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.NilOrEmptyStrPtr(nil, "empty", "notempty")

	// Act
	actual := args.Map{"val": *result}

	// Assert
	expected := args.Map{"val": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStrPtr returns nil -- nil", actual)
}

func Test_NilOrEmptyStrPtr_NonEmpty_FromNilCheckNil(t *testing.T) {
	// Arrange
	s := "hello"
	result := conditional.NilOrEmptyStrPtr(&s, "empty", "notempty")

	// Act
	actual := args.Map{"val": *result}

	// Assert
	expected := args.Map{"val": "notempty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStrPtr returns nil -- non-empty", actual)
}

// ── StringDefault ──

func Test_StringDefault_True_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.StringDefault(true, "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringDefault returns non-empty -- true", actual)
}

func Test_StringDefault_False_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.StringDefault(false, "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StringDefault false -- empty", actual)
}

// ── BoolByOrder ──

func Test_BoolByOrder_AllFalse_FromNilCheckNil(t *testing.T) {
	// Arrange
	result := conditional.BoolByOrder(false, false, false)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns non-empty -- all false", actual)
}

func Test_BoolByOrder_SecondTrue(t *testing.T) {
	// Arrange
	result := conditional.BoolByOrder(false, true, false)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns non-empty -- second true", actual)
}

// ── Func ──

func Test_Func_True_FromNilCheckNil(t *testing.T) {
	// Arrange
	f := conditional.Func(true, func() any { return "t" }, func() any { return "f" })

	// Act
	actual := args.Map{"result": f()}

	// Assert
	expected := args.Map{"result": "t"}
	expected.ShouldBeEqual(t, 0, "Func returns non-empty -- true", actual)
}

func Test_Func_False_FromNilCheckNil(t *testing.T) {
	// Arrange
	f := conditional.Func(false, func() any { return "t" }, func() any { return "f" })

	// Act
	actual := args.Map{"result": f()}

	// Assert
	expected := args.Map{"result": "f"}
	expected.ShouldBeEqual(t, 0, "Func returns non-empty -- false", actual)
}

// ── Generic If / IfFunc / IfTrueFunc / IfSlice / NilDef / NilDefPtr / IfPtr ──

func Test_Generic_If(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.If[int](true, 1, 2),
		"false": conditional.If[int](false, 1, 2),
	}

	// Assert
	expected := args.Map{
		"true": 1,
		"false": 2,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- If", actual)
}

func Test_Generic_IfFunc(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfFunc[string](true, func() string { return "a" }, func() string { return "b" }),
		"false": conditional.IfFunc[string](false, func() string { return "a" }, func() string { return "b" }),
	}

	// Assert
	expected := args.Map{
		"true": "a",
		"false": "b",
	}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- IfFunc", actual)
}

func Test_Generic_IfTrueFunc(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfTrueFunc[int](true, func() int { return 42 }),
		"false": conditional.IfTrueFunc[int](false, func() int { return 42 }),
	}

	// Assert
	expected := args.Map{
		"true": 42,
		"false": 0,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns non-empty -- IfTrueFunc", actual)
}

func Test_Generic_IfSlice(t *testing.T) {
	// Arrange
	a := []int{1, 2}
	b := []int{3, 4}

	// Act
	actual := args.Map{
		"trueLen":  len(conditional.IfSlice[int](true, a, b)),
		"falseLen": len(conditional.IfSlice[int](false, a, b)),
	}

	// Assert
	expected := args.Map{
		"trueLen": 2,
		"falseLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- IfSlice", actual)
}

func Test_Generic_NilDef(t *testing.T) {
	// Arrange
	v := 42

	// Act
	actual := args.Map{
		"nil":    conditional.NilDef[int](nil, 99),
		"nonNil": conditional.NilDef[int](&v, 99),
	}

	// Assert
	expected := args.Map{
		"nil": 99,
		"nonNil": 42,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilDef", actual)
}

func Test_Generic_NilDefPtr(t *testing.T) {
	// Arrange
	v := 42
	result := conditional.NilDefPtr[int](nil, 99)
	result2 := conditional.NilDefPtr[int](&v, 99)

	// Act
	actual := args.Map{
		"nilVal": *result,
		"nonNilVal": *result2,
	}

	// Assert
	expected := args.Map{
		"nilVal": 99,
		"nonNilVal": 42,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilDefPtr", actual)
}

func Test_Generic_IfPtr(t *testing.T) {
	// Arrange
	a := 1
	b := 2

	// Act
	actual := args.Map{
		"true":  *conditional.IfPtr[int](true, &a, &b),
		"false": *conditional.IfPtr[int](false, &a, &b),
	}

	// Assert
	expected := args.Map{
		"true": 1,
		"false": 2,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- IfPtr", actual)
}

func Test_Generic_NilVal(t *testing.T) {
	// Arrange
	v := "hello"

	// Act
	actual := args.Map{
		"nil":    conditional.NilVal[string](nil, "default", "has"),
		"nonNil": conditional.NilVal[string](&v, "default", "has"),
	}

	// Assert
	expected := args.Map{
		"nil": "default",
		"nonNil": "has",
	}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilVal", actual)
}

func Test_Generic_NilValPtr(t *testing.T) {
	// Arrange
	v := "hello"
	r1 := conditional.NilValPtr[string](nil, "d", "h")
	r2 := conditional.NilValPtr[string](&v, "d", "h")

	// Act
	actual := args.Map{
		"nil": *r1,
		"nonNil": *r2,
	}

	// Assert
	expected := args.Map{
		"nil": "d",
		"nonNil": "h",
	}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilValPtr", actual)
}

func Test_Generic_ValueOrZero(t *testing.T) {
	// Arrange
	v := 42

	// Act
	actual := args.Map{
		"nil":    conditional.ValueOrZero[int](nil),
		"nonNil": conditional.ValueOrZero[int](&v),
	}

	// Assert
	expected := args.Map{
		"nil": 0,
		"nonNil": 42,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- ValueOrZero", actual)
}

func Test_Generic_PtrOrZero(t *testing.T) {
	// Arrange
	v := 42
	r1 := conditional.PtrOrZero[int](nil)
	r2 := conditional.PtrOrZero[int](&v)

	// Act
	actual := args.Map{
		"nilVal": *r1,
		"nonNilVal": *r2,
	}

	// Assert
	expected := args.Map{
		"nilVal": 0,
		"nonNilVal": 42,
	}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- PtrOrZero", actual)
}
