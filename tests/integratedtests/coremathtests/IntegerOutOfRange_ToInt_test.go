package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coremath"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IntegerOutOfRange_ToInt(t *testing.T) {
	// Act
	actual := args.Map{"result": coremath.IsOutOfRange.Integer.ToInt(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 should be in range", actual)
}
