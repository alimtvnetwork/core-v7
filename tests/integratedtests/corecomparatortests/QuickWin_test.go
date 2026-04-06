package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_IsCompareEqualLogically_Fallthrough(t *testing.T) {
	// Cover the final `return false` branch in IsCompareEqualLogically
	c := corecomparator.Equal
	result := c.IsCompareEqualLogically(corecomparator.Inconclusive)
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for inconclusive", actual)
}
