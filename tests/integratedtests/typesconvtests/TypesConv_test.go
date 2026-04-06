package typesconvtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/typesconv"
)

// ==========================================================================
// StringToBool
// ==========================================================================

func Test_StringToBool_Verification(t *testing.T) {
	for caseIndex, tc := range stringToBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := typesconv.StringToBool(inputStr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// IntPtr / IntPtrToSimple / IntPtrToSimpleDef / IntPtrToDefPtr
// ==========================================================================

func Test_IntPtr(t *testing.T) {
	// Arrange
	tc := intPtrTestCase
	ptr := typesconv.IntPtr(42)

	// Act
	actual := args.Map{
		"value": *ptr,
		"isNil": ptr == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IntPtrToSimple_Verification(t *testing.T) {
	for caseIndex, tc := range intPtrToSimpleTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int

		if !isNil {
			v, _ := input.GetAsInt("value")
			ptr = &v
		}

		result := typesconv.IntPtrToSimple(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, tc := range intPtrToSimpleDefTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsInt("defVal")

		// Act
		var ptr *int

		if !isNil {
			v, _ := input.GetAsInt("value")
			ptr = &v
		}

		result := typesconv.IntPtrToSimpleDef(ptr, defVal)

		// Assert
		actual := args.Map{
			"result":      result,
			"defaultUsed": isNil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, tc := range intPtrToDefPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *int

		if !isNil {
			v, _ := input.GetAsInt("value")
			ptr = &v
		}

		result := typesconv.IntPtrToDefPtr(ptr, 77)

		// Assert
		actual := args.Map{"value": *result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// BoolPtr / BoolPtrToSimple / BoolPtrToSimpleDef / BoolPtrToDefPtr
// ==========================================================================

func Test_BoolPtr(t *testing.T) {
	// Arrange
	tc := boolPtrTestCase
	ptr := typesconv.BoolPtr(true)

	// Act
	actual := args.Map{
		"value": *ptr,
		"isNil": ptr == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BoolPtrToSimple_Verification(t *testing.T) {
	for caseIndex, tc := range boolPtrToSimpleTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *bool

		if !isNil {
			v, _ := input.GetAsBool("value")
			ptr = &v
		}

		result := typesconv.BoolPtrToSimple(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, tc := range boolPtrToSimpleDefTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsBool("defVal")

		// Act
		var ptr *bool

		if !isNil {
			v, _ := input.GetAsBool("value")
			ptr = &v
		}

		result := typesconv.BoolPtrToSimpleDef(ptr, defVal)

		// Assert
		actual := args.Map{
			"result":      result,
			"defaultUsed": isNil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BoolPtrToDefPtr_Verification(t *testing.T) {
	for caseIndex, tc := range boolPtrToDefPtrTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *bool

		if !isNil {
			v, _ := input.GetAsBool("value")
			ptr = &v
		}

		result := typesconv.BoolPtrToDefPtr(ptr, false)

		// Assert
		actual := args.Map{"value": *result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// BytePtr / BytePtrToSimple / BytePtrToSimpleDef
// ==========================================================================

func Test_BytePtr(t *testing.T) {
	// Arrange
	tc := bytePtrTestCase
	ptr := typesconv.BytePtr(42)

	// Act
	actual := args.Map{
		"value": int(*ptr),
		"isNil": ptr == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_BytePtrToSimple_Verification(t *testing.T) {
	for caseIndex, tc := range bytePtrToSimpleTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *byte

		if !isNil {
			v, _ := input.GetAsInt("value")
			b := byte(v)
			ptr = &b
		}

		result := typesconv.BytePtrToSimple(ptr)

		// Assert
		actual := args.Map{"result": int(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_BytePtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, tc := range bytePtrToSimpleDefTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsInt("defVal")

		// Act
		var ptr *byte

		if !isNil {
			v, _ := input.GetAsInt("value")
			b := byte(v)
			ptr = &b
		}

		result := typesconv.BytePtrToSimpleDef(ptr, byte(defVal))

		// Assert
		actual := args.Map{
			"result":      int(result),
			"defaultUsed": isNil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// FloatPtr / FloatPtrToSimple / FloatPtrToSimpleDef
// ==========================================================================

func Test_FloatPtr(t *testing.T) {
	// Arrange
	tc := floatPtrTestCase
	ptr := typesconv.FloatPtr(1.5)

	// Act
	actual := args.Map{
		"value": *ptr,
		"isNil": ptr == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FloatPtrToSimple_Verification(t *testing.T) {
	for caseIndex, tc := range floatPtrToSimpleTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *float32

		if !isNil {
			vAny, _ := input.Get("value")
			v := float32(vAny.(float64))
			ptr = &v
		}

		result := typesconv.FloatPtrToSimple(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_FloatPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, tc := range floatPtrToSimpleDefTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defAny, _ := input.Get("defVal")
		defVal := float32(defAny.(float64))

		// Act
		var ptr *float32

		if !isNil {
			vAny, _ := input.Get("value")
			v := float32(vAny.(float64))
			ptr = &v
		}

		result := typesconv.FloatPtrToSimpleDef(ptr, defVal)

		// Assert
		actual := args.Map{
			"result":      result,
			"defaultUsed": isNil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// StringPtr / StringPtrToSimple / StringPtrToSimpleDef / StringPointerToBool
// ==========================================================================

func Test_StringPtr(t *testing.T) {
	// Arrange
	tc := stringPtrTestCase
	ptr := typesconv.StringPtr("test")

	// Act
	actual := args.Map{
		"value": *ptr,
		"isNil": ptr == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_StringPtrToSimple_Verification(t *testing.T) {
	for caseIndex, tc := range stringPtrToSimpleTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *string

		if !isNil {
			v, _ := input.GetAsString("value")
			ptr = &v
		}

		result := typesconv.StringPtrToSimple(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPtrToSimpleDef_Verification(t *testing.T) {
	for caseIndex, tc := range stringPtrToSimpleDefTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true
		defVal, _ := input.GetAsString("defVal")

		// Act
		var ptr *string

		if !isNil {
			v, _ := input.GetAsString("value")
			ptr = &v
		}

		result := typesconv.StringPtrToSimpleDef(ptr, defVal)

		// Assert
		actual := args.Map{
			"result":      result,
			"defaultUsed": isNil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringPointerToBool_Verification(t *testing.T) {
	for caseIndex, tc := range stringPointerToBoolTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		var ptr *string

		if !isNil {
			v, _ := input.GetAsString("value")
			ptr = &v
		}

		result := typesconv.StringPointerToBool(ptr)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
