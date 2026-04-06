package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── NilCheck ──

func Test_Cov7_NilCheck_Nil(t *testing.T) {
	result := conditional.NilCheck(nil, "default", "nonnil")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "default"}
	expected.ShouldBeEqual(t, 0, "NilCheck nil -- default", actual)
}

func Test_Cov7_NilCheck_NonNil(t *testing.T) {
	result := conditional.NilCheck("val", "default", "nonnil")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "nonnil"}
	expected.ShouldBeEqual(t, 0, "NilCheck nonnil -- nonnil", actual)
}

// ── DefOnNil ──

func Test_Cov7_DefOnNil_Nil(t *testing.T) {
	result := conditional.DefOnNil(nil, "fallback")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "fallback"}
	expected.ShouldBeEqual(t, 0, "DefOnNil nil -- fallback", actual)
}

func Test_Cov7_DefOnNil_NonNil(t *testing.T) {
	result := conditional.DefOnNil("actual", "fallback")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "actual"}
	expected.ShouldBeEqual(t, 0, "DefOnNil nonnil -- actual", actual)
}

// ── NilOrEmptyStr ──

func Test_Cov7_NilOrEmptyStr_Nil(t *testing.T) {
	result := conditional.NilOrEmptyStr(nil, "empty", "notempty")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns nil -- nil", actual)
}

func Test_Cov7_NilOrEmptyStr_Empty(t *testing.T) {
	s := ""
	result := conditional.NilOrEmptyStr(&s, "empty", "notempty")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns nil -- empty string", actual)
}

func Test_Cov7_NilOrEmptyStr_NonEmpty(t *testing.T) {
	s := "hello"
	result := conditional.NilOrEmptyStr(&s, "empty", "notempty")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "notempty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns nil -- non-empty", actual)
}

// ── NilOrEmptyStrPtr ──

func Test_Cov7_NilOrEmptyStrPtr_Nil(t *testing.T) {
	result := conditional.NilOrEmptyStrPtr(nil, "empty", "notempty")
	actual := args.Map{"val": *result}
	expected := args.Map{"val": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStrPtr returns nil -- nil", actual)
}

func Test_Cov7_NilOrEmptyStrPtr_NonEmpty(t *testing.T) {
	s := "hello"
	result := conditional.NilOrEmptyStrPtr(&s, "empty", "notempty")
	actual := args.Map{"val": *result}
	expected := args.Map{"val": "notempty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStrPtr returns nil -- non-empty", actual)
}

// ── StringDefault ──

func Test_Cov7_StringDefault_True(t *testing.T) {
	result := conditional.StringDefault(true, "hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringDefault returns non-empty -- true", actual)
}

func Test_Cov7_StringDefault_False(t *testing.T) {
	result := conditional.StringDefault(false, "hello")
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StringDefault false -- empty", actual)
}

// ── BoolByOrder ──

func Test_Cov7_BoolByOrder_AllFalse(t *testing.T) {
	result := conditional.BoolByOrder(false, false, false)
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns non-empty -- all false", actual)
}

func Test_Cov7_BoolByOrder_SecondTrue(t *testing.T) {
	result := conditional.BoolByOrder(false, true, false)
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns non-empty -- second true", actual)
}

// ── Func ──

func Test_Cov7_Func_True(t *testing.T) {
	f := conditional.Func(true, func() any { return "t" }, func() any { return "f" })
	actual := args.Map{"result": f()}
	expected := args.Map{"result": "t"}
	expected.ShouldBeEqual(t, 0, "Func returns non-empty -- true", actual)
}

func Test_Cov7_Func_False(t *testing.T) {
	f := conditional.Func(false, func() any { return "t" }, func() any { return "f" })
	actual := args.Map{"result": f()}
	expected := args.Map{"result": "f"}
	expected.ShouldBeEqual(t, 0, "Func returns non-empty -- false", actual)
}

// ── Generic If / IfFunc / IfTrueFunc / IfSlice / NilDef / NilDefPtr / IfPtr ──

func Test_Cov7_Generic_If(t *testing.T) {
	actual := args.Map{
		"true":  conditional.If[int](true, 1, 2),
		"false": conditional.If[int](false, 1, 2),
	}
	expected := args.Map{"true": 1, "false": 2}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- If", actual)
}

func Test_Cov7_Generic_IfFunc(t *testing.T) {
	actual := args.Map{
		"true":  conditional.IfFunc[string](true, func() string { return "a" }, func() string { return "b" }),
		"false": conditional.IfFunc[string](false, func() string { return "a" }, func() string { return "b" }),
	}
	expected := args.Map{"true": "a", "false": "b"}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- IfFunc", actual)
}

func Test_Cov7_Generic_IfTrueFunc(t *testing.T) {
	actual := args.Map{
		"true":  conditional.IfTrueFunc[int](true, func() int { return 42 }),
		"false": conditional.IfTrueFunc[int](false, func() int { return 42 }),
	}
	expected := args.Map{"true": 42, "false": 0}
	expected.ShouldBeEqual(t, 0, "Generic returns non-empty -- IfTrueFunc", actual)
}

func Test_Cov7_Generic_IfSlice(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	actual := args.Map{
		"trueLen":  len(conditional.IfSlice[int](true, a, b)),
		"falseLen": len(conditional.IfSlice[int](false, a, b)),
	}
	expected := args.Map{"trueLen": 2, "falseLen": 2}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- IfSlice", actual)
}

func Test_Cov7_Generic_NilDef(t *testing.T) {
	v := 42
	actual := args.Map{
		"nil":    conditional.NilDef[int](nil, 99),
		"nonNil": conditional.NilDef[int](&v, 99),
	}
	expected := args.Map{"nil": 99, "nonNil": 42}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilDef", actual)
}

func Test_Cov7_Generic_NilDefPtr(t *testing.T) {
	v := 42
	result := conditional.NilDefPtr[int](nil, 99)
	result2 := conditional.NilDefPtr[int](&v, 99)
	actual := args.Map{"nilVal": *result, "nonNilVal": *result2}
	expected := args.Map{"nilVal": 99, "nonNilVal": 42}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilDefPtr", actual)
}

func Test_Cov7_Generic_IfPtr(t *testing.T) {
	a := 1
	b := 2
	actual := args.Map{
		"true":  *conditional.IfPtr[int](true, &a, &b),
		"false": *conditional.IfPtr[int](false, &a, &b),
	}
	expected := args.Map{"true": 1, "false": 2}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- IfPtr", actual)
}

func Test_Cov7_Generic_NilVal(t *testing.T) {
	v := "hello"
	actual := args.Map{
		"nil":    conditional.NilVal[string](nil, "default", "has"),
		"nonNil": conditional.NilVal[string](&v, "default", "has"),
	}
	expected := args.Map{"nil": "default", "nonNil": "has"}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilVal", actual)
}

func Test_Cov7_Generic_NilValPtr(t *testing.T) {
	v := "hello"
	r1 := conditional.NilValPtr[string](nil, "d", "h")
	r2 := conditional.NilValPtr[string](&v, "d", "h")
	actual := args.Map{"nil": *r1, "nonNil": *r2}
	expected := args.Map{"nil": "d", "nonNil": "h"}
	expected.ShouldBeEqual(t, 0, "Generic returns nil -- NilValPtr", actual)
}

func Test_Cov7_Generic_ValueOrZero(t *testing.T) {
	v := 42
	actual := args.Map{
		"nil":    conditional.ValueOrZero[int](nil),
		"nonNil": conditional.ValueOrZero[int](&v),
	}
	expected := args.Map{"nil": 0, "nonNil": 42}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- ValueOrZero", actual)
}

func Test_Cov7_Generic_PtrOrZero(t *testing.T) {
	v := 42
	r1 := conditional.PtrOrZero[int](nil)
	r2 := conditional.PtrOrZero[int](&v)
	actual := args.Map{"nilVal": *r1, "nonNilVal": *r2}
	expected := args.Map{"nilVal": 0, "nonNilVal": 42}
	expected.ShouldBeEqual(t, 0, "Generic returns correct value -- PtrOrZero", actual)
}
