package isanytests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

func Test_Ext_NotNull_Verification(t *testing.T) {
	for caseIndex, tc := range extNotNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := isany.NotNull(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_AllNull_Verification(t *testing.T) {
	for caseIndex, tc := range extAllNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AllNull(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_AnyNull_Verification(t *testing.T) {
	for caseIndex, tc := range extAnyNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AnyNull(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_Zero_Verification(t *testing.T) {
	for caseIndex, tc := range extZeroTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.Get("input")

		// Act
		result := isany.Zero(inputVal)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_AllZero_Verification(t *testing.T) {
	for caseIndex, tc := range extAllZeroTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AllZero(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_AnyZero_Verification(t *testing.T) {
	for caseIndex, tc := range extAnyZeroTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.AnyZero(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_Pointer_Verification(t *testing.T) {
	for caseIndex, tc := range extPointerTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		usePtrRaw, _ := input.Get("usePointer")
		usePtr := usePtrRaw.(bool)

		// Act
		var result bool
		if usePtr {
			val := 42
			result = isany.Pointer(&val)
		} else {
			inputVal, _ := input.Get("input")
			result = isany.Pointer(inputVal)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_Function_Verification(t *testing.T) {
	for caseIndex, tc := range extFunctionTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useFuncRaw, _ := input.Get("useFunc")
		useFunc := useFuncRaw.(bool)

		// Act
		var result bool
		if useFunc {
			isFunc, _ := isany.Function(func() {})
			result = isFunc
		} else {
			inputVal, _ := input.Get("input")
			isFunc, _ := isany.Function(inputVal)
			result = isFunc
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_TypeSame_Verification(t *testing.T) {
	for caseIndex, tc := range extTypeSameTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.Get("left")
		right, _ := input.Get("right")

		// Act
		result := isany.TypeSame(left, right)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_StringEqual_Verification(t *testing.T) {
	for caseIndex, tc := range extStringEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.GetAsString("left")
		right, _ := input.GetAsString("right")

		// Act
		result := isany.StringEqual(left, right)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_DefinedAllOf_Verification(t *testing.T) {
	for caseIndex, tc := range extDefinedAllOfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.DefinedAllOf(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_DefinedAnyOf_Verification(t *testing.T) {
	for caseIndex, tc := range extDefinedAnyOfTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		result := isany.DefinedAnyOf(inputs...)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_ReflectNull_Verification(t *testing.T) {
	for caseIndex, tc := range extReflectNullTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useNilPtrRaw, _ := input.Get("useNilPtr")
		useNilPtr := useNilPtrRaw.(bool)

		// Act
		var result bool
		if useNilPtr {
			var p *int
			rv := reflect.ValueOf(&p).Elem()
			result = isany.ReflectNull(rv)
		} else {
			val := 42
			rv := reflect.ValueOf(&val)
			result = isany.ReflectNull(rv)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_DefinedLeftRight_Verification(t *testing.T) {
	for caseIndex, tc := range extDefinedLeftRightTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		leftNilRaw, hasLeftNil := input.Get("leftNil")

		var left, right any
		if hasLeftNil && leftNilRaw == true {
			left = nil
		} else {
			left, _ = input.Get("left")
		}
		right, _ = input.Get("right")

		// Act
		leftDef, rightDef := isany.DefinedLeftRight(left, right)

		// Assert
		actual := args.Map{
			"leftDefined":  fmt.Sprintf("%v", leftDef),
			"rightDefined": fmt.Sprintf("%v", rightDef),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_DefinedItems_Verification(t *testing.T) {
	for caseIndex, tc := range extDefinedItemsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputsRaw, _ := input.Get("inputs")
		inputs := inputsRaw.([]any)

		// Act
		_, result := isany.DefinedItems(inputs...)

		// Assert
		actual := args.Map{
			"count": len(result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_NotDeepEqual_Verification(t *testing.T) {
	for caseIndex, tc := range extNotDeepEqualTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		left, _ := input.Get("left")
		right, _ := input.Get("right")

		// Act
		result := isany.NotDeepEqual(left, right)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_Conclusive_Verification(t *testing.T) {
	for caseIndex, tc := range extConclusiveTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		leftRaw, _ := input.Get("left")
		rightRaw, _ := input.Get("right")

		// Act
		isEqual, isConclusive := isany.Conclusive(leftRaw, rightRaw)
		result := fmt.Sprintf("%v %v", isEqual, isConclusive)
		_ = result

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_Ext_FuncOnly_Verification(t *testing.T) {
	for caseIndex, tc := range extFuncOnlyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		useFuncRaw, _ := input.Get("useFunc")
		useFunc := useFuncRaw.(bool)

		// Act
		var result bool
		if useFunc {
			result = isany.FuncOnly(func() {})
		} else {
			result = isany.FuncOnly(42)
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
