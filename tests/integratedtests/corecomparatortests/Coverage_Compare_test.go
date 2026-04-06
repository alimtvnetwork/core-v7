package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Cov_Compare_IsCompareEqualLogically_LeftLessEqual(t *testing.T) {
	result := corecomparator.LeftLess.IsCompareEqualLogically(corecomparator.LeftLessEqual)
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftLess should match LeftLessEqual logically", actual)
}

func Test_Cov_Compare_IsCompareEqualLogically_Fallthrough(t *testing.T) {
	// Inconclusive compared with LeftGreater should return false
	result := corecomparator.Inconclusive.IsCompareEqualLogically(corecomparator.LeftGreater)
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}
