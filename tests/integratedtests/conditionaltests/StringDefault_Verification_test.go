package conditionaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_StringDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range stringDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueValue, _ := input.GetAsString("trueValue")

		// Act
		result := conditional.StringDefault(isTrue, trueValue)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilOrEmptyStr_Verification(t *testing.T) {
	for caseIndex, testCase := range nilOrEmptyStrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNilOrEmpty, _ := input.GetAsString("onNilOrEmpty")
		onNonNil, _ := input.GetAsString("onNonNil")

		var strPtr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			strPtr = &val
		}

		// Act
		result := conditional.NilOrEmptyStr(strPtr, onNilOrEmpty, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilOrEmptyStrPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range nilOrEmptyStrPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNilOrEmpty, _ := input.GetAsString("onNilOrEmpty")
		onNonNil, _ := input.GetAsString("onNonNil")

		var strPtr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			strPtr = &val
		}

		// Act
		result := conditional.NilOrEmptyStrPtr(strPtr, onNilOrEmpty, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_NilDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range nilDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsString("defVal")

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.NilDefPtr[string](ptr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_NilVal_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		var ptr *string
		if !isNil {
			val := "something"
			ptr = &val
		}

		// Act
		result := conditional.NilVal[string](ptr, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_NilValPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range nilValPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		onNil, _ := input.GetAsString("onNil")
		onNonNil, _ := input.GetAsString("onNonNil")

		var ptr *string
		if !isNil {
			val := "something"
			ptr = &val
		}

		// Act
		result := conditional.NilValPtr[string](ptr, onNil, onNonNil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_ValueOrZero_Verification(t *testing.T) {
	for caseIndex, testCase := range valueOrZeroTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := conditional.ValueOrZero[string](ptr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PtrOrZero_Verification(t *testing.T) {
	for caseIndex, testCase := range ptrOrZeroTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		var ptr *int
		if !isNil {
			val, _ := input.GetAsInt("value")
			ptr = &val
		}

		// Act
		result := conditional.PtrOrZero[int](ptr)

		actual := args.Map{
			"isNil": result == nil,
			"value": *result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IfPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range ifPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true
		trueVal, _ := input.GetAsString("trueValue")
		falseVal, _ := input.GetAsString("falseValue")

		truePtr := &trueVal
		falsePtr := &falseVal

		// Act
		result := conditional.IfPtr[string](isTrue, truePtr, falsePtr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, *result)
	}
}

func Test_Func_Verification(t *testing.T) {
	for caseIndex, testCase := range funcTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isTrueVal, _ := input.Get("isTrue")
		isTrue := isTrueVal == true

		trueFunc := func() any { return "true-result" }
		falseFunc := func() any { return "false-result" }

		// Act
		resultFunc := conditional.Func(isTrue, trueFunc, falseFunc)
		result := resultFunc()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result.(string))
	}
}
