package typesconvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/typesconv"
)

func Test_Ext_BoolPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extBoolPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsBoolDefault("value", false)

		// Act
		result := typesconv.BoolPtr(val)

		actual := args.Map{
			"notNil": result != nil,
			"deref": *result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BoolPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range extBoolPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrToSimple(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BoolPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range extBoolPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)
		defVal := input.GetAsBoolDefault("defVal", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrToSimpleDef(ptr, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BoolPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extBoolPtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)
		defVal := input.GetAsBoolDefault("defVal", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrToDefPtr(ptr, defVal)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BoolPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range extBoolPtrDefValFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *bool
		if !isNil {
			val := input.GetAsBoolDefault("value", false)
			ptr = &val
		}

		// Act
		result := typesconv.BoolPtrDefValFunc(ptr, func() bool { return true })

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_IntPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extIntPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val := input.GetAsIntDefault("value", 0)

		// Act
		result := typesconv.IntPtr(val)

		actual := args.Map{
			"notNil": result != nil,
			"deref": *result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_IntPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range extIntPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *int
		if !isNil {
			val := input.GetAsIntDefault("value", 0)
			ptr = &val
		}

		// Act
		result := typesconv.IntPtrToSimple(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_IntPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range extIntPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal := input.GetAsIntDefault("defVal", 0)

		// Act
		result := typesconv.IntPtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_IntPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extIntPtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal := input.GetAsIntDefault("defVal", 0)

		// Act
		result := typesconv.IntPtrToDefPtr(nil, defVal)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_IntPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range extIntPtrDefValFuncTestCases {
		// Arrange
		_ = testCase.ArrangeInput

		// Act
		result := typesconv.IntPtrDefValFunc(nil, func() int { return 55 })

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")

		// Act
		result := typesconv.StringPtr(val)

		actual := args.Map{
			"notNil": result != nil,
			"deref": *result,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPtrToSimpleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := typesconv.StringPtrToSimple(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal, _ := input.GetAsString("defVal")

		// Act
		result := typesconv.StringPtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		defVal, _ := input.GetAsString("defVal")

		// Act
		result := typesconv.StringPtrToDefPtr(nil, defVal)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPtrDefValFuncTestCases {
		// Arrange
		_ = testCase.ArrangeInput

		// Act
		result := typesconv.StringPtrDefValFunc(nil, func() string { return "generated" })

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringToBool_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")

		// Act
		result := typesconv.StringToBool(val)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPointerToBool_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPointerToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := typesconv.StringPointerToBool(ptr)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringPointerToBoolPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringPointerToBoolPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil := input.GetAsBoolDefault("isNil", false)

		var ptr *string
		if !isNil {
			val, _ := input.GetAsString("value")
			ptr = &val
		}

		// Act
		result := typesconv.StringPointerToBoolPtr(ptr)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_StringToBoolPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extStringToBoolPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		val, _ := input.GetAsString("value")

		// Act
		result := typesconv.StringToBoolPtr(val)

		actual := args.Map{"deref": *result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BytePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extBytePtrTestCases {
		// Arrange
		result := typesconv.BytePtr(5)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BytePtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range extBytePtrToSimpleTestCases {
		// Act
		result := typesconv.BytePtrToSimple(nil)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BytePtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range extBytePtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawDef, _ := input.Get("defVal")
		defVal := rawDef.(byte)

		// Act
		result := typesconv.BytePtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BytePtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extBytePtrToDefPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawDef, _ := input.Get("defVal")
		defVal := rawDef.(byte)

		// Act
		result := typesconv.BytePtrToDefPtr(nil, defVal)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_BytePtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range extBytePtrDefValFuncTestCases {
		// Act
		result := typesconv.BytePtrDefValFunc(nil, func() byte { return 7 })

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_FloatPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extFloatPtrTestCases {
		// Act
		result := typesconv.FloatPtr(3.14)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_FloatPtrToSimple_Verification(t *testing.T) {
	for caseIndex, testCase := range extFloatPtrToSimpleTestCases {
		// Act
		result := typesconv.FloatPtrToSimple(nil)

		actual := args.Map{"isZero": result == 0}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_FloatPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, testCase := range extFloatPtrToSimpleDefTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		rawDef, _ := input.Get("defVal")
		defVal := rawDef.(float32)

		// Act
		result := typesconv.FloatPtrToSimpleDef(nil, defVal)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_FloatPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extFloatPtrToDefPtrTestCases {
		// Act
		result := typesconv.FloatPtrToDefPtr(nil, 2.5)

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Ext_FloatPtrDefValFunc_Verification(t *testing.T) {
	for caseIndex, testCase := range extFloatPtrDefValFuncTestCases {
		// Act
		result := typesconv.FloatPtrDefValFunc(nil, func() float32 { return 9.9 })

		actual := args.Map{"notNil": result != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
