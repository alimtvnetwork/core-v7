package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IntegerWithin_ToInt32_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToInt32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt32(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToUnsignedInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToUint16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToUnsignedInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToUnsignedInt32_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToUint32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToUnsignedInt32(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToUnsignedInt64_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToUint64TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToUnsignedInt64(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToInt8_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToInt8TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToInt8(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToInt32_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToInt32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToInt32(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToUnsignedInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToUint16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToUnsignedInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToUnsignedInt64_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToUint64TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToUnsignedInt64(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MaxFloat32_Verification(t *testing.T) {
	for caseIndex, tc := range maxFloat32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		aRaw, _ := input.Get("a")
		bRaw, _ := input.Get("b")
		a := float32(aRaw.(float64))
		b := float32(bRaw.(float64))

		// Act
		result := coremath.MaxFloat32(a, b)

		// Assert
		actual := args.Map{"result": float64(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinFloat32_Verification(t *testing.T) {
	for caseIndex, tc := range minFloat32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		aRaw, _ := input.Get("a")
		bRaw, _ := input.Get("b")
		a := float32(aRaw.(float64))
		b := float32(bRaw.(float64))

		// Act
		result := coremath.MinFloat32(a, b)

		// Assert
		actual := args.Map{"result": float64(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
