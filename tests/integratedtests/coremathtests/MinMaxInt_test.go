package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_MaxInt_Verification(t *testing.T) {
	for caseIndex, tc := range maxIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MaxInt(a, b)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinInt_Verification(t *testing.T) {
	for caseIndex, tc := range minIntTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MinInt(a, b)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MaxByte_Verification(t *testing.T) {
	for caseIndex, tc := range maxByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MaxByte(byte(a), byte(b))

		// Assert
		actual := args.Map{"result": int(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MinByte_Verification(t *testing.T) {
	for caseIndex, tc := range minByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		a, _ := input.GetAsInt("a")
		b, _ := input.GetAsInt("b")

		// Act
		result := coremath.MinByte(byte(a), byte(b))

		// Assert
		actual := args.Map{"result": int(result)}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToByte_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToByte(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToInt8_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToInt8TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt8(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOutOfRange_ToByte_Verification(t *testing.T) {
	for caseIndex, tc := range integerOutOfRangeToByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsOutOfRange.Integer.ToByte(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerWithin_ToInt16_Verification(t *testing.T) {
	for caseIndex, tc := range integerWithinToInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		value, _ := input.GetAsInt("value")

		// Act
		result := coremath.IsRangeWithin.Integer.ToInt16(value)

		// Assert
		actual := args.Map{"result": result}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
