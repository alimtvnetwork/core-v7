package conditionaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// BoolByOrder
// ═══════════════════════════════════════════

func Test_BoolByOrder_FirstTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder(true, false, false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns true -- first true", actual)
}

func Test_BoolByOrder_AllFalse(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder(false, false, false)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns false -- all false", actual)
}

func Test_BoolByOrder_LastTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder(false, false, true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns true -- last true", actual)
}

// ═══════════════════════════════════════════
// BoolFunctionsByOrder
// ═══════════════════════════════════════════

func Test_BoolFunctionsByOrder_FirstTrue_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	f1 := func() bool { return true }
	f2 := func() bool { return false }

	// Act
	actual := args.Map{"result": conditional.BoolFunctionsByOrder(f1, f2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns true -- first func true", actual)
}

func Test_BoolFunctionsByOrder_AllFalse_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	f1 := func() bool { return false }
	f2 := func() bool { return false }

	// Act
	actual := args.Map{"result": conditional.BoolFunctionsByOrder(f1, f2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns false -- all false", actual)
}

// ═══════════════════════════════════════════
// ErrorFunc
// ═══════════════════════════════════════════

func Test_ErrorFunc_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	errTrue := func() error { return errors.New("true-err") }
	errFalse := func() error { return errors.New("false-err") }
	result := conditional.ErrorFunc(true, errTrue, errFalse)

	// Act
	actual := args.Map{"err": result().Error()}

	// Assert
	expected := args.Map{"err": "true-err"}
	expected.ShouldBeEqual(t, 0, "ErrorFunc returns trueFunc -- condition true", actual)
}

func Test_ErrorFunc_False_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	errTrue := func() error { return errors.New("true-err") }
	errFalse := func() error { return errors.New("false-err") }
	result := conditional.ErrorFunc(false, errTrue, errFalse)

	// Act
	actual := args.Map{"err": result().Error()}

	// Assert
	expected := args.Map{"err": "false-err"}
	expected.ShouldBeEqual(t, 0, "ErrorFunc returns falseFunc -- condition false", actual)
}

// ═══════════════════════════════════════════
// ErrorFunctionResult
// ═══════════════════════════════════════════

func Test_ErrorFunctionResult_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	errTrue := func() error { return errors.New("true-err") }
	errFalse := func() error { return nil }
	result := conditional.ErrorFunctionResult(true, errTrue, errFalse)

	// Act
	actual := args.Map{
		"hasErr": result != nil,
		"err": result.Error(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"err": "true-err",
	}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionResult returns error -- condition true", actual)
}

func Test_ErrorFunctionResult_False_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	errTrue := func() error { return errors.New("true-err") }
	errFalse := func() error { return nil }
	result := conditional.ErrorFunctionResult(false, errTrue, errFalse)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionResult returns nil -- condition false", actual)
}

// ═══════════════════════════════════════════
// ErrorFunctionsExecuteResults
// ═══════════════════════════════════════════

func Test_ErrorFunctionsExecuteResults_True(t *testing.T) {
	// Arrange
	trueFuncs := []func() error{
		func() error { return nil },
		func() error { return errors.New("err1") },
	}
	falseFuncs := []func() error{func() error { return nil }}
	result := conditional.ErrorFunctionsExecuteResults(true, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"hasErr": result != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults returns error -- one func errors", actual)
}

func Test_ErrorFunctionsExecuteResults_False(t *testing.T) {
	// Arrange
	trueFuncs := []func() error{func() error { return errors.New("err1") }}
	falseFuncs := []func() error{func() error { return nil }}
	result := conditional.ErrorFunctionsExecuteResults(false, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults returns nil -- false path all ok", actual)
}

func Test_ErrorFunctionsExecuteResults_EmptyFuncs(t *testing.T) {
	// Arrange
	result := conditional.ErrorFunctionsExecuteResults(true, []func() error{}, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults returns nil -- empty funcs", actual)
}

func Test_ErrorFunctionsExecuteResults_NilFuncInSlice(t *testing.T) {
	// Arrange
	funcs := []func() error{nil, func() error { return nil }}
	result := conditional.ErrorFunctionsExecuteResults(true, funcs, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults skips nil -- nil func in slice", actual)
}

// ═══════════════════════════════════════════
// StringsIndexVal
// ═══════════════════════════════════════════

func Test_StringsIndexVal_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.StringsIndexVal(true, []string{"a", "b", "c"}, 0, 2)}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "StringsIndexVal returns trueValue index -- condition true", actual)
}

func Test_StringsIndexVal_False_FromBoolByOrderFirstTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.StringsIndexVal(false, []string{"a", "b", "c"}, 0, 2)}

	// Assert
	expected := args.Map{"result": "c"}
	expected.ShouldBeEqual(t, 0, "StringsIndexVal returns falseValue index -- condition false", actual)
}

// ═══════════════════════════════════════════
// Functions (generic)
// ═══════════════════════════════════════════

func Test_Functions_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "t1", true, false },
	}
	falseFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "f1", true, false },
	}
	result := conditional.Functions[string](true, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Functions returns trueFuncs -- condition true", actual)
}

func Test_Functions_False(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "t1", true, false },
	}
	falseFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "f1", true, false },
		func() (string, bool, bool) { return "f2", true, false },
	}
	result := conditional.Functions[string](false, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Functions returns falseFuncs -- condition false", actual)
}

// ═══════════════════════════════════════════
// FunctionsExecuteResults (generic)
// ═══════════════════════════════════════════

func Test_FunctionsExecuteResults_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "t1", true, false },
		func() (string, bool, bool) { return "t2", true, false },
	}
	result := conditional.FunctionsExecuteResults[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "t1",
	}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults returns results -- true path", actual)
}

func Test_FunctionsExecuteResults_Break_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	funcs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "a", true, false },
		func() (string, bool, bool) { return "b", true, true }, // break
		func() (string, bool, bool) { return "c", true, false },
	}
	result := conditional.FunctionsExecuteResults[string](true, funcs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults stops at break -- isTake+isBreak", actual)
}

func Test_FunctionsExecuteResults_SkipNotTake(t *testing.T) {
	// Arrange
	funcs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "a", false, false }, // skip
		func() (string, bool, bool) { return "b", true, false },
	}
	result := conditional.FunctionsExecuteResults[string](true, funcs, nil)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "b",
	}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults skips isTake=false -- selective take", actual)
}

func Test_FunctionsExecuteResults_NilFunc_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	funcs := []func() (string, bool, bool){
		nil,
		func() (string, bool, bool) { return "a", true, false },
	}
	result := conditional.FunctionsExecuteResults[string](true, funcs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults skips nil func -- nil in slice", actual)
}

func Test_FunctionsExecuteResults_Empty_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	result := conditional.FunctionsExecuteResults[string](true, nil, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults returns nil -- empty funcs", actual)
}

// ═══════════════════════════════════════════
// AnyFunctions
// ═══════════════════════════════════════════

func Test_AnyFunctions_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "t1", true, false },
	}
	result := conditional.AnyFunctions(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctions returns trueFuncs -- condition true", actual)
}

func Test_AnyFunctions_False_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	falseFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "f1", true, false },
	}
	result := conditional.AnyFunctions(false, nil, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctions returns falseFuncs -- condition false", actual)
}

// ═══════════════════════════════════════════
// AnyFunctionsExecuteResults
// ═══════════════════════════════════════════

func Test_AnyFunctionsExecuteResults_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "r1", true, false },
		func() (any, bool, bool) { return "r2", true, false },
	}
	result := conditional.AnyFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults returns results -- true path", actual)
}

func Test_AnyFunctionsExecuteResults_Break(t *testing.T) {
	// Arrange
	funcs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, true },
		func() (any, bool, bool) { return "b", true, false },
	}
	result := conditional.AnyFunctionsExecuteResults(true, funcs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults stops at break -- isBreak", actual)
}

func Test_AnyFunctionsExecuteResults_Empty_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	result := conditional.AnyFunctionsExecuteResults(true, nil, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults returns nil -- empty funcs", actual)
}

// ═══════════════════════════════════════════
// VoidFunctions
// ═══════════════════════════════════════════

func Test_VoidFunctions_True_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	called := false
	trueFuncs := []func(){func() { called = true }}
	falseFuncs := []func(){}
	conditional.VoidFunctions(true, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"trueCalled": called}

	// Assert
	expected := args.Map{"trueCalled": true}
	expected.ShouldBeEqual(t, 0, "VoidFunctions executes true funcs -- condition true", actual)
}

func Test_VoidFunctions_False_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	called := false
	trueFuncs := []func(){}
	falseFuncs := []func(){func() { called = true }}
	conditional.VoidFunctions(false, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"falseCalled": called}

	// Assert
	expected := args.Map{"falseCalled": true}
	expected.ShouldBeEqual(t, 0, "VoidFunctions executes false funcs -- condition false", actual)
}

func Test_VoidFunctions_NilFunc(t *testing.T) {
	// Arrange
	called := false
	funcs := []func(){nil, func() { called = true }}
	conditional.VoidFunctions(true, funcs, nil)

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "VoidFunctions skips nil -- nil func in slice", actual)
}

// ═══════════════════════════════════════════
// TypedErrorFunctionsExecuteResults
// ═══════════════════════════════════════════

func Test_TypedErrorFunctionsExecuteResults_True(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, error){
		func() (string, error) { return "r1", nil },
		func() (string, error) { return "", errors.New("err1") },
		func() (string, error) { return "r3", nil },
	}
	results, err := conditional.TypedErrorFunctionsExecuteResults[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"resultLen": len(results),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"resultLen": 2,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults returns results+error -- mixed", actual)
}

func Test_TypedErrorFunctionsExecuteResults_Empty_FromBoolByOrderFirstTrue(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[string](true, nil, nil)

	// Act
	actual := args.Map{
		"isNil": results == nil,
		"errNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults returns nil -- empty", actual)
}

func Test_TypedErrorFunctionsExecuteResults_NilFunc(t *testing.T) {
	// Arrange
	funcs := []func() (int, error){
		nil,
		func() (int, error) { return 42, nil },
	}
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](true, funcs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"errNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"errNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults skips nil -- nil in slice", actual)
}

// ═══════════════════════════════════════════
// Typed wrappers — bool
// ═══════════════════════════════════════════

func Test_IfFuncBool(t *testing.T) {
	// Arrange
	trueF := func() bool { return true }
	falseF := func() bool { return false }

	// Act
	actual := args.Map{
		"true":  conditional.IfFuncBool(true, trueF, falseF),
		"false": conditional.IfFuncBool(false, trueF, falseF),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": false,
	}
	expected.ShouldBeEqual(t, 0, "IfFuncBool returns bool -- both conditions", actual)
}

func Test_IfTrueFuncBool(t *testing.T) {
	// Arrange
	trueF := func() bool { return true }

	// Act
	actual := args.Map{
		"true":  conditional.IfTrueFuncBool(true, trueF),
		"false": conditional.IfTrueFuncBool(false, trueF),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": false,
	}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncBool returns bool -- both conditions", actual)
}

func Test_NilDefBool(t *testing.T) {
	// Arrange
	v := true

	// Act
	actual := args.Map{
		"nonNil": conditional.NilDefBool(&v, false),
		"nil":    conditional.NilDefBool(nil, true),
	}

	// Assert
	expected := args.Map{
		"nonNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "NilDefBool returns bool -- nil and non-nil", actual)
}

func Test_NilDefPtrBool(t *testing.T) {
	// Arrange
	v := true
	nonNilResult := conditional.NilDefPtrBool(&v, false)
	nilResult := conditional.NilDefPtrBool(nil, true)

	// Act
	actual := args.Map{
		"nonNilVal": *nonNilResult,
		"nilVal":    *nilResult,
	}

	// Assert
	expected := args.Map{
		"nonNilVal": true,
		"nilVal": true,
	}
	expected.ShouldBeEqual(t, 0, "NilDefPtrBool returns *bool -- nil and non-nil", actual)
}

func Test_ValueOrZeroBool(t *testing.T) {
	// Arrange
	v := true

	// Act
	actual := args.Map{
		"nonNil": conditional.ValueOrZeroBool(&v),
		"nil":    conditional.ValueOrZeroBool(nil),
	}

	// Assert
	expected := args.Map{
		"nonNil": true,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroBool returns bool -- nil and non-nil", actual)
}

func Test_PtrOrZeroBool(t *testing.T) {
	// Arrange
	v := true
	nonNilResult := conditional.PtrOrZeroBool(&v)
	nilResult := conditional.PtrOrZeroBool(nil)

	// Act
	actual := args.Map{
		"nonNilVal": *nonNilResult,
		"nilVal":    *nilResult,
	}

	// Assert
	expected := args.Map{
		"nonNilVal": true,
		"nilVal": false,
	}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroBool returns *bool -- nil and non-nil", actual)
}

func Test_NilValBool(t *testing.T) {
	// Arrange
	v := true

	// Act
	actual := args.Map{
		"nonNil": conditional.NilValBool(&v, false, true),
		"nil":    conditional.NilValBool(nil, true, false),
	}

	// Assert
	expected := args.Map{
		"nonNil": true,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "NilValBool returns bool -- nil and non-nil", actual)
}

func Test_NilValPtrBool(t *testing.T) {
	// Arrange
	v := true
	nonNilResult := conditional.NilValPtrBool(&v, false, true)
	nilResult := conditional.NilValPtrBool(nil, true, false)

	// Act
	actual := args.Map{
		"nonNilVal": *nonNilResult,
		"nilVal": *nilResult,
	}

	// Assert
	expected := args.Map{
		"nonNilVal": true,
		"nilVal": true,
	}
	expected.ShouldBeEqual(t, 0, "NilValPtrBool returns *bool -- nil and non-nil", actual)
}

func Test_IfSliceBool(t *testing.T) {
	// Arrange
	trueSlice := []bool{true}
	falseSlice := []bool{false}

	// Act
	actual := args.Map{
		"trueLen":  len(conditional.IfSliceBool(true, trueSlice, falseSlice)),
		"falseLen": len(conditional.IfSliceBool(false, trueSlice, falseSlice)),
	}

	// Assert
	expected := args.Map{
		"trueLen": 1,
		"falseLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "IfSliceBool returns slice -- both conditions", actual)
}

func Test_IfPtrBool(t *testing.T) {
	// Arrange
	v1 := true
	v2 := false
	result := conditional.IfPtrBool(true, &v1, &v2)

	// Act
	actual := args.Map{"val": *result}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "IfPtrBool returns *bool -- condition true", actual)
}

// ═══════════════════════════════════════════
// Typed wrappers — byte
// ═══════════════════════════════════════════

func Test_IfByte(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfByte(true, byte(1), byte(2)),
		"false": conditional.IfByte(false, byte(1), byte(2)),
	}

	// Assert
	expected := args.Map{
		"true": byte(1),
		"false": byte(2),
	}
	expected.ShouldBeEqual(t, 0, "IfByte returns byte -- both conditions", actual)
}

func Test_IfFuncByte(t *testing.T) {
	// Arrange
	trueF := func() byte { return 1 }
	falseF := func() byte { return 2 }

	// Act
	actual := args.Map{"true": conditional.IfFuncByte(true, trueF, falseF)}

	// Assert
	expected := args.Map{"true": byte(1)}
	expected.ShouldBeEqual(t, 0, "IfFuncByte returns byte -- condition true", actual)
}

func Test_IfTrueFuncByte(t *testing.T) {
	// Arrange
	trueF := func() byte { return 42 }

	// Act
	actual := args.Map{
		"true":  conditional.IfTrueFuncByte(true, trueF),
		"false": conditional.IfTrueFuncByte(false, trueF),
	}

	// Assert
	expected := args.Map{
		"true": byte(42),
		"false": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncByte returns byte -- both conditions", actual)
}

func Test_NilDefByte(t *testing.T) {
	// Arrange
	v := byte(42)

	// Act
	actual := args.Map{
		"nonNil": conditional.NilDefByte(&v, 0),
		"nil":    conditional.NilDefByte(nil, 99),
	}

	// Assert
	expected := args.Map{
		"nonNil": byte(42),
		"nil": byte(99),
	}
	expected.ShouldBeEqual(t, 0, "NilDefByte returns byte -- nil and non-nil", actual)
}

func Test_ValueOrZeroByte(t *testing.T) {
	// Arrange
	v := byte(42)

	// Act
	actual := args.Map{
		"nonNil": conditional.ValueOrZeroByte(&v),
		"nil":    conditional.ValueOrZeroByte(nil),
	}

	// Assert
	expected := args.Map{
		"nonNil": byte(42),
		"nil": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroByte returns byte -- nil and non-nil", actual)
}
