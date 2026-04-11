package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_IsOutOfRange_Integer_ToUnsignedInt32(t *testing.T) {
	// Arrange
	result := coremath.IsOutOfRange.Integer.ToUnsignedInt32(-1)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for negative value", actual)
	result2 := coremath.IsOutOfRange.Integer.ToUnsignedInt32(100)
	actual = args.Map{"result": result2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for valid value", actual)
}
