package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_DoubleQuoteWrapElements_Nil(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false, nil...)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_DoubleQuoteWrapElementsWithIndexes_Nil(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(nil...)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}
