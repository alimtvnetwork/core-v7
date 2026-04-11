package conditionaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ── AnyFunctions ──

func Test_AnyFunctions_True(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, false },
	}
	result := conditional.AnyFunctions(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctions true -- returns trueFuncs", actual)
}

func Test_AnyFunctions_False(t *testing.T) {
	// Arrange
	falseFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "b", true, false },
	}
	result := conditional.AnyFunctions(false, nil, falseFuncs)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctions false -- returns falseFuncs", actual)
}

// ── AnyFunctionsExecuteResults ──

func Test_AnyFunctionsExecuteResults_True(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, false },
		func() (any, bool, bool) { return "b", true, false },
	}
	result := conditional.AnyFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults true -- 2 items", actual)
}

func Test_AnyFunctionsExecuteResults_WithBreak(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "a", true, true },
		func() (any, bool, bool) { return "b", true, false },
	}
	result := conditional.AnyFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults break -- 1 item", actual)
}

func Test_AnyFunctionsExecuteResults_Empty(t *testing.T) {
	// Arrange
	result := conditional.AnyFunctionsExecuteResults(true, nil, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults empty -- nil", actual)
}

func Test_AnyFunctionsExecuteResults_NilFunc(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		nil,
		func() (any, bool, bool) { return "a", true, false },
	}
	result := conditional.AnyFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults nil func -- skipped", actual)
}

func Test_AnyFunctionsExecuteResults_NotTaken(t *testing.T) {
	// Arrange
	trueFuncs := []func() (any, bool, bool){
		func() (any, bool, bool) { return "skip", false, false },
		func() (any, bool, bool) { return "take", true, false },
	}
	result := conditional.AnyFunctionsExecuteResults(true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyFunctionsExecuteResults not taken -- 1 item", actual)
}

// ── VoidFunctions ──

func Test_VoidFunctions_True(t *testing.T) {
	// Arrange
	called := false
	conditional.VoidFunctions(true,
		[]func(){func() { called = true }},
		nil,
	)

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "VoidFunctions true -- true funcs called", actual)
}

func Test_VoidFunctions_False(t *testing.T) {
	// Arrange
	called := false
	conditional.VoidFunctions(false,
		nil,
		[]func(){func() { called = true }},
	)

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "VoidFunctions false -- false funcs called", actual)
}

func Test_VoidFunctions_NilFuncs(t *testing.T) {
	// Arrange
	conditional.VoidFunctions(true,
		[]func(){nil, func() {}},
		nil,
	)

	// Act
	actual := args.Map{"noPanic": true}

	// Assert
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "VoidFunctions nil func -- skipped", actual)
}

// ── Setter / SetterDefault ──

func Test_Setter_True(t *testing.T) {
	// Arrange
	result := conditional.Setter(true, issetter.True, issetter.False)

	// Act
	actual := args.Map{"isTrue": result.IsTrue()}

	// Assert
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "Setter true -- returns trueValue", actual)
}

func Test_Setter_False(t *testing.T) {
	// Arrange
	result := conditional.Setter(false, issetter.True, issetter.False)

	// Act
	actual := args.Map{"isFalse": result.IsFalse()}

	// Assert
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "Setter false -- returns falseValue", actual)
}

func Test_SetterDefault_Unset(t *testing.T) {
	// Arrange
	result := conditional.SetterDefault(issetter.Uninitialized, issetter.True)

	// Act
	actual := args.Map{"isTrue": result.IsTrue()}

	// Assert
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "SetterDefault unset -- returns default", actual)
}

func Test_SetterDefault_Set(t *testing.T) {
	// Arrange
	result := conditional.SetterDefault(issetter.False, issetter.True)

	// Act
	actual := args.Map{"isFalse": result.IsFalse()}

	// Assert
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "SetterDefault set -- returns current", actual)
}

// ── StringsIndexVal ──

func Test_StringsIndexVal_True(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}
	result := conditional.StringsIndexVal(true, slice, 0, 2)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a"}
	expected.ShouldBeEqual(t, 0, "StringsIndexVal true -- trueValue index", actual)
}

func Test_StringsIndexVal_False(t *testing.T) {
	// Arrange
	slice := []string{"a", "b", "c"}
	result := conditional.StringsIndexVal(false, slice, 0, 2)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "c"}
	expected.ShouldBeEqual(t, 0, "StringsIndexVal false -- falseValue index", actual)
}

// ── ErrorFunc / ErrorFunctionResult ──

func Test_ErrorFunc_True(t *testing.T) {
	// Arrange
	f := conditional.ErrorFunc(true,
		func() error { return errors.New("true") },
		func() error { return errors.New("false") },
	)

	// Act
	actual := args.Map{"msg": f().Error()}

	// Assert
	expected := args.Map{"msg": "true"}
	expected.ShouldBeEqual(t, 0, "ErrorFunc true -- returns true func", actual)
}

func Test_ErrorFunc_False(t *testing.T) {
	// Arrange
	f := conditional.ErrorFunc(false,
		func() error { return errors.New("true") },
		func() error { return errors.New("false") },
	)

	// Act
	actual := args.Map{"msg": f().Error()}

	// Assert
	expected := args.Map{"msg": "false"}
	expected.ShouldBeEqual(t, 0, "ErrorFunc false -- returns false func", actual)
}

func Test_ErrorFunctionResult_True(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionResult(true,
		func() error { return nil },
		func() error { return errors.New("false") },
	)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionResult true -- nil", actual)
}

func Test_ErrorFunctionResult_False(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionResult(false,
		func() error { return nil },
		func() error { return errors.New("false") },
	)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionResult false -- error", actual)
}

// ── ErrorFunctionsExecuteResults ──

func Test_ErrorFunctionsExecuteResults_True_NoErrors(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionsExecuteResults(true,
		[]func() error{func() error { return nil }},
		nil,
	)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults true -- no errors", actual)
}

func Test_ErrorFunctionsExecuteResults_WithErrors(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionsExecuteResults(true,
		[]func() error{
			func() error { return errors.New("e1") },
			nil,
			func() error { return errors.New("e2") },
		},
		nil,
	)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults errors -- aggregated", actual)
}

func Test_ErrorFunctionsExecuteResults_Empty(t *testing.T) {
	// Arrange
	err := conditional.ErrorFunctionsExecuteResults(true, nil, nil)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ErrorFunctionsExecuteResults empty -- nil", actual)
}

// ── Functions / FunctionsExecuteResults (generic) ──

func Test_Functions_True(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, bool, bool){
		func() (string, bool, bool) { return "a", true, false },
	}
	result := conditional.Functions[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Functions generic true -- returns trueFuncs", actual)
}

func Test_FunctionsExecuteResults_True(t *testing.T) {
	// Arrange
	trueFuncs := []func() (int, bool, bool){
		func() (int, bool, bool) { return 1, true, false },
		func() (int, bool, bool) { return 2, true, false },
	}
	result := conditional.FunctionsExecuteResults[int](true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults true -- 2 items", actual)
}

func Test_FunctionsExecuteResults_Break(t *testing.T) {
	// Arrange
	trueFuncs := []func() (int, bool, bool){
		func() (int, bool, bool) { return 1, true, true },
		func() (int, bool, bool) { return 2, true, false },
	}
	result := conditional.FunctionsExecuteResults[int](true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults break -- 1 item", actual)
}

func Test_FunctionsExecuteResults_NilFunc(t *testing.T) {
	// Arrange
	trueFuncs := []func() (int, bool, bool){
		nil,
		func() (int, bool, bool) { return 1, true, false },
	}
	result := conditional.FunctionsExecuteResults[int](true, trueFuncs, nil)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults nil func -- skipped", actual)
}

func Test_FunctionsExecuteResults_Empty(t *testing.T) {
	// Arrange
	result := conditional.FunctionsExecuteResults[int](true, nil, nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FunctionsExecuteResults empty -- nil", actual)
}

// ── TypedErrorFunctionsExecuteResults ──

func Test_TypedErrorFunctionsExecuteResults_True_NoErrors(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, error){
		func() (string, error) { return "a", nil },
		func() (string, error) { return "b", nil },
	}
	results, err := conditional.TypedErrorFunctionsExecuteResults[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults returns empty -- no errors", actual)
}

func Test_TypedErrorFunctionsExecuteResults_WithErrors(t *testing.T) {
	// Arrange
	trueFuncs := []func() (string, error){
		func() (string, error) { return "a", nil },
		func() (string, error) { return "", errors.New("fail") },
		nil,
	}
	results, err := conditional.TypedErrorFunctionsExecuteResults[string](true, trueFuncs, nil)

	// Act
	actual := args.Map{
		"len": len(results),
		"hasError": err != nil,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults returns error -- with errors", actual)
}

func Test_TypedErrorFunctionsExecuteResults_Empty_FromAnyFunctionsTrue(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "TypedErrorFunctionsExecuteResults returns empty -- empty", actual)
}

// ── BoolFunctionsByOrder ──

func Test_BoolFunctionsByOrder_FirstTrue(t *testing.T) {
	// Arrange
	result := conditional.BoolFunctionsByOrder(
		func() bool { return true },
		func() bool { return false },
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns non-empty -- first true", actual)
}

func Test_BoolFunctionsByOrder_AllFalse(t *testing.T) {
	// Arrange
	result := conditional.BoolFunctionsByOrder(
		func() bool { return false },
		func() bool { return false },
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder returns non-empty -- all false", actual)
}

func Test_BoolFunctionsByOrder_Empty(t *testing.T) {
	// Arrange
	result := conditional.BoolFunctionsByOrder()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BoolFunctionsByOrder empty -- false", actual)
}

// ── typed wrappers ──

func Test_IfTrueFuncStrings_True(t *testing.T) {
	// Arrange
	result := conditional.IfTrueFuncStrings(true, func() []string { return []string{"a"} })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncStrings true -- returns value", actual)
}

func Test_IfTrueFuncStrings_False(t *testing.T) {
	// Arrange
	result := conditional.IfTrueFuncStrings(false, func() []string { return []string{"a"} })

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncStrings false -- nil", actual)
}

func Test_IfTrueFuncBytes_True(t *testing.T) {
	// Arrange
	result := conditional.IfTrueFuncBytes(true, func() []byte { return []byte{1} })

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncBytes true -- returns value", actual)
}

func Test_IfSliceAny_True(t *testing.T) {
	// Arrange
	result := conditional.IfSliceAny(true, []any{1}, []any{2})

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "IfSliceAny true -- trueValue", actual)
}

func Test_IfAny_True(t *testing.T) {
	// Arrange
	result := conditional.IfAny(true, "yes", "no")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "yes"}
	expected.ShouldBeEqual(t, 0, "IfAny returns non-empty -- true", actual)
}

func Test_IfFuncAny_True(t *testing.T) {
	// Arrange
	result := conditional.IfFuncAny(true,
		func() any { return "yes" },
		func() any { return "no" },
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "yes"}
	expected.ShouldBeEqual(t, 0, "IfFuncAny returns non-empty -- true", actual)
}
