package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Compare_IsCompareEqualLogically_LeftLessEqual(t *testing.T) {
	// Arrange
	result := corecomparator.LeftLess.IsCompareEqualLogically(corecomparator.LeftLessEqual)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftLess should match LeftLessEqual logically", actual)
}

func Test_Compare_IsCompareEqualLogically_Fallthrough(t *testing.T) {
	// Arrange
	// Inconclusive compared with LeftGreater should return false
	result := corecomparator.Inconclusive.IsCompareEqualLogically(corecomparator.LeftGreater)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}
