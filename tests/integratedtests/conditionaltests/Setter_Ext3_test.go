package conditionaltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ============================================================================
// Setter
// ============================================================================

func Test_Setter_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extSetterTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueVal := issetter.GetBool(true)
		falseVal := issetter.GetBool(false)

		// Act
		result := conditional.Setter(isTrue, trueVal, falseVal)

		// Assert
		actual := args.Map{
			"isTrue": result.IsTrue(),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// SetterDefault
// ============================================================================

func Test_SetterDefault_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extSetterDefaultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useUnsetVal, _ := input.Get("useUnset")
		useUnset := useUnsetVal == true

		defVal := issetter.GetBool(true)
		var current issetter.Value
		if useUnset {
			current = issetter.Uninitialized
		} else {
			current = issetter.GetBool(false)
		}

		// Act
		result := conditional.SetterDefault(current, defVal)

		// Assert
		actual := args.Map{
			"isTrue": result.IsTrue(),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// BoolFunctionsByOrder
// ============================================================================

func Test_BoolFunctionsByOrder_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extBoolFuncsByOrderTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		valuesRaw, _ := input.Get("values")
		values := valuesRaw.([]bool)

		var funcs []func() bool
		for _, v := range values {
			val := v
			funcs = append(funcs, func() bool { return val })
		}

		// Act
		result := conditional.BoolFunctionsByOrder(funcs...)

		// Assert
		actual := args.Map{
			"result": result,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// VoidFunctions
// ============================================================================

func Test_VoidFunctions_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extVoidFunctionsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueCount := 0
		falseCount := 0

		trueFuncs := []func(){
			func() { trueCount++ },
			func() { trueCount++ },
		}
		falseFuncs := []func(){
			func() { falseCount++ },
		}

		// Act
		conditional.VoidFunctions(isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"trueCount":  trueCount,
			"falseCount": falseCount,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// ErrorFunc
// ============================================================================

func Test_ErrorFunc_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extErrorFuncTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueFunc := func() error { return errors.New("true-err") }
		falseFunc := func() error { return errors.New("false-err") }

		// Act
		resultFunc := conditional.ErrorFunc(isTrue, trueFunc, falseFunc)
		err := resultFunc()

		// Assert
		actual := args.Map{
			"result": err.Error(),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// ErrorFunctionResult
// ============================================================================

func Test_ErrorFunctionResult_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extErrorFunctionResultTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueFunc := func() error { return errors.New("true-err") }
		falseFunc := func() error { return errors.New("false-err") }

		// Act
		err := conditional.ErrorFunctionResult(isTrue, trueFunc, falseFunc)

		// Assert
		actual := args.Map{
			"result": err.Error(),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// ErrorFunctionsExecuteResults
// ============================================================================

func Test_ErrorFunctionsExecuteResults_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extErrorFunctionsExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		hasErrorVal, _ := input.Get("hasError")
		hasError := hasErrorVal == true
		emptyVal, _ := input.Get("empty")
		isEmpty := emptyVal == true

		var trueFuncs []func() error
		var falseFuncs []func() error

		if isEmpty {
			trueFuncs = []func() error{}
			falseFuncs = []func() error{}
		} else if hasError {
			trueFuncs = []func() error{
				func() error { return errors.New("err1") },
			}
			falseFuncs = []func() error{}
		} else {
			trueFuncs = []func() error{
				func() error { return nil },
			}
			falseFuncs = []func() error{}
		}

		// Act
		err := conditional.ErrorFunctionsExecuteResults(isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// AnyFunctions
// ============================================================================

func Test_AnyFunctions_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extAnyFunctionsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueFuncs := []func() (any, bool, bool){
			func() (any, bool, bool) { return "a", true, false },
			func() (any, bool, bool) { return "b", true, false },
		}
		falseFuncs := []func() (any, bool, bool){
			func() (any, bool, bool) { return "c", true, false },
		}

		// Act
		result := conditional.AnyFunctions(isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// AnyFunctionsExecuteResults
// ============================================================================

func Test_AnyFunctionsExecuteResults_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extAnyFunctionsExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		scenario, _ := input.GetAsString("scenario")

		var trueFuncs []func() (any, bool, bool)
		falseFuncs := []func() (any, bool, bool){}

		switch scenario {
		case "take-all":
			trueFuncs = []func() (any, bool, bool){
				func() (any, bool, bool) { return "a", true, false },
				func() (any, bool, bool) { return "b", true, false },
			}
		case "break-early":
			trueFuncs = []func() (any, bool, bool){
				func() (any, bool, bool) { return "a", true, true },
				func() (any, bool, bool) { return "b", true, false },
			}
		case "with-nil":
			trueFuncs = []func() (any, bool, bool){
				nil,
				func() (any, bool, bool) { return "a", true, false },
			}
		case "empty":
			trueFuncs = []func() (any, bool, bool){}
		}

		// Act
		result := conditional.AnyFunctionsExecuteResults(isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// FunctionsExecuteResults (generic)
// ============================================================================

func Test_FunctionsExecuteResults_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extFunctionsExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		scenario, _ := input.GetAsString("scenario")

		trueFuncs := []func() (string, bool, bool){
			func() (string, bool, bool) { return "a", true, false },
			func() (string, bool, bool) { return "b", true, false },
		}
		falseFuncs := []func() (string, bool, bool){
			func() (string, bool, bool) { return "c", true, false },
		}

		if scenario == "break-early" {
			trueFuncs = []func() (string, bool, bool){
				func() (string, bool, bool) { return "a", true, true },
				func() (string, bool, bool) { return "b", true, false },
			}
		}

		// Act
		result := conditional.FunctionsExecuteResults[string](isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Functions (generic selector)
// ============================================================================

func Test_Functions_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extFunctionsSelectorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueFuncs := []func() (string, bool, bool){
			func() (string, bool, bool) { return "a", true, false },
			func() (string, bool, bool) { return "b", true, false },
		}
		falseFuncs := []func() (string, bool, bool){
			func() (string, bool, bool) { return "c", true, false },
		}

		// Act
		result := conditional.Functions[string](isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"length": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// TypedErrorFunctionsExecuteResults
// ============================================================================

func Test_TypedErrorFunctionsExecuteResults_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extTypedErrorFunctionsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		scenario, _ := input.GetAsString("scenario")

		var trueFuncs []func() (int, error)
		falseFuncs := []func() (int, error){}

		switch scenario {
		case "success":
			trueFuncs = []func() (int, error){
				func() (int, error) { return 1, nil },
				func() (int, error) { return 2, nil },
			}
		case "mixed":
			trueFuncs = []func() (int, error){
				func() (int, error) { return 1, nil },
				func() (int, error) { return 0, errors.New("fail") },
			}
		case "empty":
			trueFuncs = []func() (int, error){}
		case "with-nil":
			trueFuncs = []func() (int, error){
				nil,
				func() (int, error) { return 1, nil },
			}
		}

		// Act
		results, err := conditional.TypedErrorFunctionsExecuteResults[int](isTrue, trueFuncs, falseFuncs)

		// Assert
		actual := args.Map{
			"resultLen": len(results),
			"hasError":  err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Typed convenience wrappers
// ============================================================================

func Test_NilDefString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilDefStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsString("defVal")

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.NilDefString(ptr, defVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilDefPtrString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilDefPtrStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsString("defVal")

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.NilDefPtrString(ptr, defVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_ValueOrZeroString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extValueOrZeroStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.ValueOrZeroString(ptr)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PtrOrZeroString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extPtrOrZeroStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.PtrOrZeroString(ptr)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_NilValString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilValStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.NilValString(ptr, onNil, onNonNil)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilValPtrString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilValPtrStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.NilValPtrString(ptr, onNil, onNonNil)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_NilDefBool_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilDefBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defValRaw, _ := input.Get("defVal")
		defVal := defValRaw == true

		var ptr *bool
		if !isNil {
			valRaw, _ := input.Get("value")
			v := valRaw == true
			ptr = &v
		}

		// Act
		result := conditional.NilDefBool(ptr, defVal)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilDefInt_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilDefIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsInt("defVal")

		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		// Act
		result := conditional.NilDefInt(ptr, defVal)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilDefByte_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilDefByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defValRaw, _ := input.Get("defVal")
		defVal := defValRaw.(byte)

		var ptr *byte
		if !isNil {
			valRaw, _ := input.Get("value")
			v := valRaw.(byte)
			ptr = &v
		}

		// Act
		result := conditional.NilDefByte(ptr, defVal)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ValueOrZeroBool_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extValueOrZeroBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *bool
		if !isNil {
			valRaw, _ := input.Get("value")
			v := valRaw == true
			ptr = &v
		}

		// Act
		result := conditional.ValueOrZeroBool(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ValueOrZeroInt_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extValueOrZeroIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		// Act
		result := conditional.ValueOrZeroInt(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ValueOrZeroByte_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extValueOrZeroByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *byte
		if !isNil {
			valRaw, _ := input.Get("value")
			v := valRaw.(byte)
			ptr = &v
		}

		// Act
		result := conditional.ValueOrZeroByte(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilValBool_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilValBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNilRaw, _ := input.Get("onNil")
		onNil := onNilRaw == true
		onNonNilRaw, _ := input.Get("onNonNil")
		onNonNil := onNonNilRaw == true

		var ptr *bool
		if !isNil {
			valRaw, _ := input.Get("value")
			v := valRaw == true
			ptr = &v
		}

		// Act
		result := conditional.NilValBool(ptr, onNil, onNonNil)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilValInt_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilValIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsInt("onNil")
		onNonNil, _ := input.GetAsInt("onNonNil")

		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		// Act
		result := conditional.NilValInt(ptr, onNil, onNonNil)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NilValByte_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extNilValByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNilRaw, _ := input.Get("onNil")
		onNil := onNilRaw.(byte)
		onNonNilRaw, _ := input.Get("onNonNil")
		onNonNil := onNonNilRaw.(byte)

		var ptr *byte
		if !isNil {
			valRaw, _ := input.Get("value")
			v := valRaw.(byte)
			ptr = &v
		}

		// Act
		result := conditional.NilValByte(ptr, onNil, onNonNil)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfString_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extIfStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsString("trueValue")
		falseVal, _ := input.GetAsString("falseValue")

		// Act
		result := conditional.IfString(isTrue, trueVal, falseVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_IfTrueFuncInt_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extIfTrueFuncIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsInt("trueValue")

		// Act
		result := conditional.IfTrueFuncInt(isTrue, func() int { return trueVal })

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfFuncByte_Ext3_Verification(t *testing.T) {
	for caseIndex, tc := range extIfFuncByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValRaw, _ := input.Get("trueValue")
		trueVal := trueValRaw.(byte)
		falseValRaw, _ := input.Get("falseValue")
		falseVal := falseValRaw.(byte)

		// Act
		result := conditional.IfFuncByte(isTrue, func() byte { return trueVal }, func() byte { return falseVal })

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Additional typed wrappers: NilDefPtrBool, NilDefPtrInt, NilDefPtrByte,
// PtrOrZeroBool, PtrOrZeroInt, PtrOrZeroByte, NilValPtrBool, NilValPtrInt, NilValPtrByte
// ============================================================================

func Test_NilDefPtrBool_Ext3_Verification(t *testing.T) {
	// Arrange - nil
	resultNil := conditional.NilDefPtrBool(nil, true)
	actual := args.Map{"result": resultNil == nil || *resultNil != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilDefPtrBool nil case failed", actual)

	// Arrange - non-nil
	v := false
	resultNonNil := conditional.NilDefPtrBool(&v, true)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != false}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilDefPtrBool non-nil case failed", actual)
}

func Test_NilDefPtrInt_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.NilDefPtrInt(nil, 42)

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilDefPtrInt nil case failed", actual)

	v := 7
	resultNonNil := conditional.NilDefPtrInt(&v, 42)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != 7}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilDefPtrInt non-nil case failed", actual)
}

func Test_NilDefPtrByte_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.NilDefPtrByte(nil, byte(0xFF))

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != byte(0xFF)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilDefPtrByte nil case failed", actual)

	v := byte(0x0A)
	resultNonNil := conditional.NilDefPtrByte(&v, byte(0xFF))
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != byte(0x0A)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilDefPtrByte non-nil case failed", actual)
}

func Test_PtrOrZeroBool_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.PtrOrZeroBool(nil)

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != false}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroBool nil case failed", actual)

	v := true
	resultNonNil := conditional.PtrOrZeroBool(&v)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != true}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroBool non-nil case failed", actual)
}

func Test_PtrOrZeroInt_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.PtrOrZeroInt(nil)

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroInt nil case failed", actual)

	v := 99
	resultNonNil := conditional.PtrOrZeroInt(&v)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != 99}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroInt non-nil case failed", actual)
}

func Test_PtrOrZeroByte_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.PtrOrZeroByte(nil)

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != byte(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroByte nil case failed", actual)

	v := byte(42)
	resultNonNil := conditional.PtrOrZeroByte(&v)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != byte(42)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PtrOrZeroByte non-nil case failed", actual)
}

func Test_NilValPtrBool_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.NilValPtrBool(nil, true, false)

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != true}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilValPtrBool nil case failed", actual)

	v := true
	resultNonNil := conditional.NilValPtrBool(&v, true, false)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != false}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilValPtrBool non-nil case failed", actual)
}

func Test_NilValPtrInt_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.NilValPtrInt(nil, -1, 1)

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilValPtrInt nil case failed", actual)

	v := 50
	resultNonNil := conditional.NilValPtrInt(&v, -1, 1)
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilValPtrInt non-nil case failed", actual)
}

func Test_NilValPtrByte_Ext3_Verification(t *testing.T) {
	// Arrange
	resultNil := conditional.NilValPtrByte(nil, byte(0), byte(1))

	// Act
	actual := args.Map{"result": resultNil == nil || *resultNil != byte(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilValPtrByte nil case failed", actual)

	v := byte(5)
	resultNonNil := conditional.NilValPtrByte(&v, byte(0), byte(1))
	actual = args.Map{"result": resultNonNil == nil || *resultNonNil != byte(1)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NilValPtrByte non-nil case failed", actual)
}

// ============================================================================
// IfTrueFuncByte
// ============================================================================

func Test_IfTrueFuncByte_Ext3_Verification(t *testing.T) {
	// Arrange
	// true case
	result := conditional.IfTrueFuncByte(true, func() byte { return byte(42) })

	// Act
	actual := args.Map{"result": result != byte(42)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncByte true case failed", actual)

	// false case
	result2 := conditional.IfTrueFuncByte(false, func() byte { return byte(42) })
	actual = args.Map{"result": result2 != byte(0)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IfTrueFuncByte false case failed", actual)
}
