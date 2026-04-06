package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_IsCompareEqualLogically_Fallthrough(t *testing.T) {
	// Arrange
	// Cover the final `return false` branch in IsCompareEqualLogically
	c := corecomparator.Equal
	result := c.IsCompareEqualLogically(corecomparator.Inconclusive)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for inconclusive", actual)
}
