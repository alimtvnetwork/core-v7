package conditionaltests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ============================================================================
// BoolByOrder
// ============================================================================

func Test_BoolByOrder_AllFalse_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder(false, false, false)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns false -- all false", actual)
}

func Test_BoolByOrder_SecondTrue_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder(false, true, false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns true -- second true", actual)
}

func Test_BoolByOrder_Empty_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolByOrder()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolByOrder returns false -- empty input", actual)
}

// ============================================================================
// BoolFunctionsByOrder
// ============================================================================

func Test_BoolFunctionsByOrder_AllFalse_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolFunctionsByOrder(
		func() bool { return false },
		func() bool { return false },
	)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns false -- all false funcs", actual)
}

func Test_BoolFunctionsByOrder_SecondTrue_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolFunctionsByOrder(
		func() bool { return false },
		func() bool { return true },
	)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns true -- second func true", actual)
}

func Test_BoolFunctionsByOrder_Empty_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.BoolFunctionsByOrder()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns false -- no funcs", actual)
}

// ============================================================================
// Setter / SetterDefault
// ============================================================================

func Test_Setter_True_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", conditional.Setter(true, issetter.True, issetter.False))}

	// Assert
	expected := args.Map{"result": "True"}
	expected.ShouldBeEqual(t, 0, "Setter returns trueValue -- isTrue", actual)
}

func Test_Setter_False_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", conditional.Setter(false, issetter.True, issetter.False))}

	// Assert
	expected := args.Map{"result": "False"}
	expected.ShouldBeEqual(t, 0, "Setter returns falseValue -- isFalse", actual)
}

func Test_SetterDefault_Uninitialized_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", conditional.SetterDefault(issetter.Uninitialized, issetter.True))}

	// Assert
	expected := args.Map{"result": "True"}
	expected.ShouldBeEqual(t, 0, "SetterDefault returns defVal -- uninitialized", actual)
}

func Test_SetterDefault_Initialized_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", conditional.SetterDefault(issetter.False, issetter.True))}

	// Assert
	expected := args.Map{"result": "False"}
	expected.ShouldBeEqual(t, 0, "SetterDefault returns current -- initialized", actual)
}

// ============================================================================
// StringDefault
// ============================================================================

func Test_StringDefault_True_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.StringDefault(true, "hello")}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "StringDefault returns value -- true", actual)
}

func Test_StringDefault_False_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.StringDefault(false, "hello")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "StringDefault returns empty -- false", actual)
}

// ============================================================================
// StringsIndexVal
// ============================================================================

func Test_StringsIndexVal_True_Cov5(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	actual := args.Map{"result": conditional.StringsIndexVal(true, slice, 0, 2)}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "StringsIndexVal returns trueIndex -- true", actual)
}

func Test_StringsIndexVal_False_Cov5(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}

	// Act
	actual := args.Map{"result": conditional.StringsIndexVal(false, slice, 0, 2)}

	// Assert
	expected := args.Map{"result": "c"}
	expected.ShouldBeEqual(t, 0, "StringsIndexVal returns falseIndex -- false", actual)
}

// ============================================================================
// NilCheck
// ============================================================================

func Test_NilCheck_Nil_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilCheck(nil, "nil-val", "non-nil-val")}

	// Assert
	expected := args.Map{"result": "nil-val"}
	expected.ShouldBeEqual(t, 0, "NilCheck returns onNil -- nil input", actual)
}

func Test_NilCheck_NonNil_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilCheck("something", "nil-val", "non-nil-val")}

	// Assert
	expected := args.Map{"result": "non-nil-val"}
	expected.ShouldBeEqual(t, 0, "NilCheck returns onNonNil -- non-nil input", actual)
}

// ============================================================================
// DefOnNil
// ============================================================================

func Test_DefOnNil_Nil_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.DefOnNil(nil, "default")}

	// Assert
	expected := args.Map{"result": "default"}
	expected.ShouldBeEqual(t, 0, "DefOnNil returns default -- nil input", actual)
}

func Test_DefOnNil_NonNil_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.DefOnNil("val", "default")}

	// Assert
	expected := args.Map{"result": "val"}
	expected.ShouldBeEqual(t, 0, "DefOnNil returns value -- non-nil input", actual)
}

// ============================================================================
// NilOrEmptyStr / NilOrEmptyStrPtr
// ============================================================================

func Test_NilOrEmptyStr_Nil_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilOrEmptyStr(nil, "empty", "full")}

	// Assert
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns onNilOrEmpty -- nil", actual)
}

func Test_NilOrEmptyStr_Empty_Cov5(t *testing.T) {
	// Arrange
	s := ""

	// Act
	actual := args.Map{"result": conditional.NilOrEmptyStr(&s, "empty", "full")}

	// Assert
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns onNilOrEmpty -- empty string", actual)
}

func Test_NilOrEmptyStr_NonEmpty_Cov5(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	actual := args.Map{"result": conditional.NilOrEmptyStr(&s, "empty", "full")}

	// Assert
	expected := args.Map{"result": "full"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStr returns onNonNilOrNonEmpty -- non-empty", actual)
}

func Test_NilOrEmptyStrPtr_Nil_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilOrEmptyStrPtr(nil, "empty", "full")

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "empty"}
	expected.ShouldBeEqual(t, 0, "NilOrEmptyStrPtr returns ptr to onNilOrEmpty -- nil", actual)
}

// ============================================================================
// ErrorFunc
// ============================================================================

func Test_ErrorFunc_True_Cov5(t *testing.T) {
	// Arrange
	trueF := func() error { return errors.New("true-err") }
	falseF := func() error { return errors.New("false-err") }
	f := conditional.ErrorFunc(true, trueF, falseF)

	// Act
	actual := args.Map{"result": f().Error()}

	// Assert
	expected := args.Map{"result": "true-err"}
	expected.ShouldBeEqual(t, 0, "ErrorFunc returns trueFunc -- true", actual)
}

func Test_ErrorFunc_False_Cov5(t *testing.T) {
	// Arrange
	trueF := func() error { return errors.New("true-err") }
	falseF := func() error { return errors.New("false-err") }
	f := conditional.ErrorFunc(false, trueF, falseF)

	// Act
	actual := args.Map{"result": f().Error()}

	// Assert
	expected := args.Map{"result": "false-err"}
	expected.ShouldBeEqual(t, 0, "ErrorFunc returns falseFunc -- false", actual)
}

// ============================================================================
// ErrorFunctionResult
// ============================================================================

func Test_ErrorFunctionResult_True_Cov5(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionResult(true,
		func() error { return errors.New("true-err") },
		func() error { return nil })

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"msg": err.Error(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"msg": "true-err",
	}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionResult returns trueFunc result -- true", actual)
}

func Test_ErrorFunctionResult_False_Cov5(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionResult(false,
		func() error { return errors.New("true-err") },
		func() error { return nil })

	// Act
	actual := args.Map{"hasErr": err == nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionResult returns falseFunc result -- false", actual)
}

// ============================================================================
// Func
// ============================================================================

func Test_Func_True_Cov5(t *testing.T) {
	// Arrange
	f := conditional.Func(true,
		func() any { return "true-val" },
		func() any { return "false-val" })

	// Act
	actual := args.Map{"result": f()}

	// Assert
	expected := args.Map{"result": "true-val"}
	expected.ShouldBeEqual(t, 0, "Func returns trueFunc -- true", actual)
}

func Test_Func_False_Cov5(t *testing.T) {
	// Arrange
	f := conditional.Func(false,
		func() any { return "true-val" },
		func() any { return "false-val" })

	// Act
	actual := args.Map{"result": f()}

	// Assert
	expected := args.Map{"result": "false-val"}
	expected.ShouldBeEqual(t, 0, "Func returns falseFunc -- false", actual)
}

// ============================================================================
// AnyFunctions
// ============================================================================

func Test_AnyFunctions_True_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, false },
	}
	falseFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "b", true, false },
	}
	result := conditional.AnyFunctions(true, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctions returns trueFuncs -- true", actual)
}

func Test_AnyFunctions_False_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, false },
	}
	falseFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "b", true, false },
		func() (any, bool, bool) { return "c", true, false },
	}
	result := conditional.AnyFunctions(false, trueFuncs, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyFunctions returns falseFuncs -- false", actual)
}

// ============================================================================
// AnyFunctionsExecuteResults
// ============================================================================

func Test_AnyFunctionsExecuteResults_True_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, false },
		func() (any, bool, bool) { return "b", true, false },
	}
	results := conditional.AnyFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"first": results[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults returns true results -- true", actual)
}

func Test_AnyFunctionsExecuteResults_Empty_Cov5(t *testing.T) {
	// Arrange
	results := conditional.AnyFunctionsExecuteResults(true, nil, nil)

	// Act
	actual := args.Map{"isNil": results == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults returns nil -- empty", actual)
}

func Test_AnyFunctionsExecuteResults_Break_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, true },
		func() (any, bool, bool) { return "b", true, false },
	}
	results := conditional.AnyFunctionsExecuteResults(true, funcs, nil)

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults stops on break -- isBreak", actual)
}

func Test_AnyFunctionsExecuteResults_SkipNil_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() (any, bool, bool){
		nil,
		func() (any, bool, bool) { return "a", true, false },
	}
	results := conditional.AnyFunctionsExecuteResults(true, funcs, nil)

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults skips nil funcs -- nil entry", actual)
}

func Test_AnyFunctionsExecuteResults_NotTake_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", false, false },
	}
	results := conditional.AnyFunctionsExecuteResults(true, funcs, nil)

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults skips not-take -- isTake false", actual)
}

// ============================================================================
// FunctionsExecuteResults (generic)
// ============================================================================

func Test_FunctionsExecuteResults_True_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "a", true, false },
		func() (string, bool, bool) { return "b", true, false },
	}
	results := conditional.FunctionsExecuteResults[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"first": results[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults returns results -- true", actual)
}

func Test_FunctionsExecuteResults_Empty_Cov5(t *testing.T) {
	// Arrange
	results := conditional.FunctionsExecuteResults[int](true, nil, nil)

	// Act
	actual := args.Map{"isNil": results == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults returns nil -- empty funcs", actual)
}

func Test_FunctionsExecuteResults_Break_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() (int, bool, bool){
		func() (int, bool, bool) { return 1, true, true },
		func() (int, bool, bool) { return 2, true, false },
	}
	results := conditional.FunctionsExecuteResults[int](true, funcs, nil)

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults stops on break -- isBreak", actual)
}

func Test_FunctionsExecuteResults_SkipNil_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() (int, bool, bool){
		nil,
		func() (int, bool, bool) { return 1, true, false },
	}
	results := conditional.FunctionsExecuteResults[int](true, funcs, nil)

	// Act
	actual := args.Map{"len": len(results)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults skips nil -- nil entry", actual)
}

// ============================================================================
// Functions (generic selector)
// ============================================================================

func Test_Functions_True_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "a", true, false },
	}
	result := conditional.Functions[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Functions returns trueFuncs -- true", actual)
}

func Test_Functions_False_Cov5(t *testing.T) {
	// Arrange
	falseFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "b", true, false },
	}
	result := conditional.Functions[string](false, nil, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Functions returns falseFuncs -- false", actual)
}

// ============================================================================
// ErrorFunctionsExecuteResults
// ============================================================================

func Test_ErrorFunctionsExecuteResults_True_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() error{
		func() error { return errors.New("err1") },
	}
	err := conditional.ErrorFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults returns error -- true with errs", actual)
}

func Test_ErrorFunctionsExecuteResults_NoErr_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() error{
		func() error { return nil },
	}
	err := conditional.ErrorFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults returns nil -- true no errs", actual)
}

func Test_ErrorFunctionsExecuteResults_Empty_Cov5(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionsExecuteResults(true, nil, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults returns nil -- empty funcs", actual)
}

func Test_ErrorFunctionsExecuteResults_SkipNil_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() error{
		nil,
		func() error { return nil },
	}
	err := conditional.ErrorFunctionsExecuteResults(true, funcs, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults skips nil funcs -- nil entry", actual)
}

// ============================================================================
// TypedErrorFunctionsExecuteResults
// ============================================================================

func Test_TypedErrorFunctionsExecuteResults_True_Cov5(t *testing.T) {
	// Arrange
	trueFuncs := []func() (int, error){
		func() (int, error) { return 1, nil },
		func() (int, error) { return 2, errors.New("err") },
		func() (int, error) { return 3, nil },
	}
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults collects results and errors -- mixed", actual)
}

func Test_TypedErrorFunctionsExecuteResults_Empty_Cov5(t *testing.T) {
	// Arrange
	results, err := conditional.TypedErrorFunctionsExecuteResults[int](true, nil, nil)

	// Act
	actual := args.Map{
		"isNil": results == nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults returns nil -- empty", actual)
}

func Test_TypedErrorFunctionsExecuteResults_SkipNil_Cov5(t *testing.T) {
	// Arrange
	funcs := []func() (string, error){
		nil,
		func() (string, error) { return "a", nil },
	}
	results, err := conditional.TypedErrorFunctionsExecuteResults[string](true, funcs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults skips nil -- nil entry", actual)
}

// ============================================================================
// VoidFunctions
// ============================================================================

func Test_VoidFunctions_True_Cov5(t *testing.T) {
	// Arrange
	counter := 0
	trueFuncs := []func(){func() { counter++ }}
	falseFuncs := []func(){func() { counter += 10 }}
	conditional.VoidFunctions(true, trueFuncs, falseFuncs)
	// VoidFunctions executes both when true (true funcs then false funcs)

	// Act
	actual := args.Map{"counter": counter}

	// Assert
	expected := args.Map{"counter": 11}
	expected.ShouldBeEqual(t, 0, "VoidFunctions executes both -- true", actual)
}

func Test_VoidFunctions_False_Cov5(t *testing.T) {
	// Arrange
	counter := 0
	trueFuncs := []func(){func() { counter++ }}
	falseFuncs := []func(){func() { counter += 10 }}
	conditional.VoidFunctions(false, trueFuncs, falseFuncs)
	// When false, only falseFuncs are executed

	// Act
	actual := args.Map{"counter": counter}

	// Assert
	expected := args.Map{"counter": 10}
	expected.ShouldBeEqual(t, 0, "VoidFunctions executes false only -- false", actual)
}

func Test_VoidFunctions_NilEntry_Cov5(t *testing.T) {
	// Arrange
	counter := 0
	funcs := []func(){nil, func() { counter++ }}
	conditional.VoidFunctions(false, nil, funcs)

	// Act
	actual := args.Map{"counter": counter}

	// Assert
	expected := args.Map{"counter": 1}
	expected.ShouldBeEqual(t, 0, "VoidFunctions skips nil entries -- nil entry", actual)
}

// ============================================================================
// Typed wrappers: byte
// ============================================================================

func Test_IfByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfByte(true, 1, 2),
		"false": conditional.IfByte(false, 1, 2),
	}

	// Assert
	expected := args.Map{
		"true": byte(1),
		"false": byte(2),
	}
	expected.ShouldBeEqual(t, 0, "IfByte returns correct value -- true/false", actual)
}

func Test_IfFuncByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"true": conditional.IfFuncByte(true, func() byte { return 10 }, func() byte { return 20 }),
	}

	// Assert
	expected := args.Map{"true": byte(10)}
	expected.ShouldBeEqual(t, 0, "IfFuncByte returns trueFunc result -- true", actual)
}

func Test_IfTrueFuncByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfTrueFuncByte(true, func() byte { return 5 }),
		"false": conditional.IfTrueFuncByte(false, func() byte { return 5 }),
	}

	// Assert
	expected := args.Map{
		"true": byte(5),
		"false": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncByte returns value or zero -- true/false", actual)
}

func Test_IfSliceByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"len": len(conditional.IfSliceByte(true, []byte{1, 2}, []byte{3})),
	}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "IfSliceByte returns trueSlice -- true", actual)
}

func Test_IfPtrByte_Cov5(t *testing.T) {
	// Arrange
	a := byte(1)
	b := byte(2)
	result := conditional.IfPtrByte(true, &a, &b)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": byte(1)}
	expected.ShouldBeEqual(t, 0, "IfPtrByte returns truePtr -- true", actual)
}

func Test_NilDefByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"nil":    conditional.NilDefByte(nil, 99),
		"nonNil": conditional.NilDefByte(func() *byte { v := byte(5); return &v }(), 99),
	}

	// Assert
	expected := args.Map{
		"nil": byte(99),
		"nonNil": byte(5),
	}
	expected.ShouldBeEqual(t, 0, "NilDefByte returns default or value -- nil/nonNil", actual)
}

func Test_NilDefPtrByte_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilDefPtrByte(nil, 42)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": byte(42)}
	expected.ShouldBeEqual(t, 0, "NilDefPtrByte returns ptr to default -- nil", actual)
}

func Test_ValueOrZeroByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.ValueOrZeroByte(nil)}

	// Assert
	expected := args.Map{"result": byte(0)}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroByte returns zero -- nil", actual)
}

func Test_PtrOrZeroByte_Cov5(t *testing.T) {
	// Arrange
	result := conditional.PtrOrZeroByte(nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": byte(0)}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroByte returns ptr to zero -- nil", actual)
}

func Test_NilValByte_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"nil":    conditional.NilValByte(nil, 10, 20),
		"nonNil": conditional.NilValByte(func() *byte { v := byte(5); return &v }(), 10, 20),
	}

	// Assert
	expected := args.Map{
		"nil": byte(10),
		"nonNil": byte(20),
	}
	expected.ShouldBeEqual(t, 0, "NilValByte returns onNil or onNonNil -- nil/nonNil", actual)
}

func Test_NilValPtrByte_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilValPtrByte(nil, 10, 20)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": byte(10)}
	expected.ShouldBeEqual(t, 0, "NilValPtrByte returns ptr to onNil -- nil", actual)
}

// ============================================================================
// Typed wrappers: bool
// ============================================================================

func Test_NilDefBool_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilDefBool(nil, true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NilDefBool returns default -- nil", actual)
}

func Test_NilDefPtrBool_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilDefPtrBool(nil, true)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NilDefPtrBool returns ptr to default -- nil", actual)
}

func Test_ValueOrZeroBool_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.ValueOrZeroBool(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroBool returns zero -- nil", actual)
}

func Test_PtrOrZeroBool_Cov5(t *testing.T) {
	// Arrange
	result := conditional.PtrOrZeroBool(nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroBool returns ptr to zero -- nil", actual)
}

func Test_NilValBool_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilValBool(nil, true, false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NilValBool returns onNil -- nil", actual)
}

func Test_NilValPtrBool_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilValPtrBool(nil, true, false)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NilValPtrBool returns ptr to onNil -- nil", actual)
}

// ============================================================================
// Typed wrappers: int (uncovered functions)
// ============================================================================

func Test_IfFuncInt_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"result": conditional.IfFuncInt(true, func() int { return 10 }, func() int { return 20 }),
	}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "IfFuncInt returns trueFunc result -- true", actual)
}

func Test_IfTrueFuncInt_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfTrueFuncInt(true, func() int { return 5 }),
		"false": conditional.IfTrueFuncInt(false, func() int { return 5 }),
	}

	// Assert
	expected := args.Map{
		"true": 5,
		"false": 0,
	}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncInt returns value or zero -- true/false", actual)
}

func Test_IfSliceInt_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"len": len(conditional.IfSliceInt(true, []int{1, 2}, []int{3}))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "IfSliceInt returns trueSlice -- true", actual)
}

func Test_IfPtrInt_Cov5(t *testing.T) {
	// Arrange
	a, b := 1, 2
	result := conditional.IfPtrInt(true, &a, &b)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "IfPtrInt returns truePtr -- true", actual)
}

func Test_NilDefInt_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilDefInt(nil, 99)}

	// Assert
	expected := args.Map{"result": 99}
	expected.ShouldBeEqual(t, 0, "NilDefInt returns default -- nil", actual)
}

func Test_NilDefPtrInt_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilDefPtrInt(nil, 42)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "NilDefPtrInt returns ptr to default -- nil", actual)
}

func Test_ValueOrZeroInt_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.ValueOrZeroInt(nil)}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroInt returns zero -- nil", actual)
}

func Test_PtrOrZeroInt_Cov5(t *testing.T) {
	// Arrange
	result := conditional.PtrOrZeroInt(nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroInt returns ptr to zero -- nil", actual)
}

func Test_NilValInt_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilValInt(nil, 10, 20)}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "NilValInt returns onNil -- nil", actual)
}

func Test_NilValPtrInt_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilValPtrInt(nil, 10, 20)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": 10}
	expected.ShouldBeEqual(t, 0, "NilValPtrInt returns ptr to onNil -- nil", actual)
}

// ============================================================================
// Typed wrappers: string (uncovered functions)
// ============================================================================

func Test_IfFuncString_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"result": conditional.IfFuncString(true, func() string { return "a" }, func() string { return "b" }),
	}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "IfFuncString returns trueFunc result -- true", actual)
}

func Test_IfTrueFuncString_Cov5(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  conditional.IfTrueFuncString(true, func() string { return "val" }),
		"false": conditional.IfTrueFuncString(false, func() string { return "val" }),
	}

	// Assert
	expected := args.Map{
		"true": "val",
		"false": "",
	}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncString returns value or zero -- true/false", actual)
}

func Test_IfSliceString_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"len": len(conditional.IfSliceString(true, []string{"a", "b"}, []string{"c"}))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "IfSliceString returns trueSlice -- true", actual)
}

func Test_IfPtrString_Cov5(t *testing.T) {
	// Arrange
	a, b := "x", "y"
	result := conditional.IfPtrString(true, &a, &b)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "x"}
	expected.ShouldBeEqual(t, 0, "IfPtrString returns truePtr -- true", actual)
}

func Test_NilDefString_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilDefString(nil, "default")}

	// Assert
	expected := args.Map{"result": "default"}
	expected.ShouldBeEqual(t, 0, "NilDefString returns default -- nil", actual)
}

func Test_NilDefPtrString_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilDefPtrString(nil, "def")

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "def"}
	expected.ShouldBeEqual(t, 0, "NilDefPtrString returns ptr to default -- nil", actual)
}

func Test_ValueOrZeroString_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.ValueOrZeroString(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ValueOrZeroString returns zero -- nil", actual)
}

func Test_PtrOrZeroString_Cov5(t *testing.T) {
	// Arrange
	result := conditional.PtrOrZeroString(nil)

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroString returns ptr to zero -- nil", actual)
}

func Test_NilValString_Cov5(t *testing.T) {
	// Act
	actual := args.Map{"result": conditional.NilValString(nil, "nil", "nonnil")}

	// Assert
	expected := args.Map{"result": "nil"}
	expected.ShouldBeEqual(t, 0, "NilValString returns onNil -- nil", actual)
}

func Test_NilValPtrString_Cov5(t *testing.T) {
	// Arrange
	result := conditional.NilValPtrString(nil, "nil", "nonnil")

	// Act
	actual := args.Map{"result": *result}

	// Assert
	expected := args.Map{"result": "nil"}
	expected.ShouldBeEqual(t, 0, "NilValPtrString returns ptr to onNil -- nil", actual)
}
