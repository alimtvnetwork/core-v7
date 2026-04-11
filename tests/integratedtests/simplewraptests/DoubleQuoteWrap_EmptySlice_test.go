package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov8_DoubleQuoteWrapElements_EmptyNonNilSlice tests the length==0 branch
// with a non-nil empty slice.
func Test_DoubleQuoteWrapElements_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	actual := simplewrap.DoubleQuoteWrapElements(false, input...)

	// Assert
	actual := args.Map{"result": len(actual) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements with empty slice: got len, want 0", actual)
}

// Test_Cov8_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice tests length==0 branch.
func Test_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	actual := simplewrap.DoubleQuoteWrapElementsWithIndexes(input...)

	// Assert
	actual := args.Map{"result": len(actual) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes with empty slice: got len, want 0", actual)
}
