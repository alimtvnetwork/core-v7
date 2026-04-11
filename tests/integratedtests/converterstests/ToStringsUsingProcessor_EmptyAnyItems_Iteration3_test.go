package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov6_ToStringsUsingProcessor_EmptyAnyItems covers
// converters/anyItemConverter.go L147-149: when ToAnyItems returns empty slice.
func Test_ToStringsUsingProcessor_EmptyAnyItems(t *testing.T) {
	// Arrange
	processor := func(index int, in any) (string, bool, bool) {
		return "", false, false
	}

	// Act
	result := converters.AnyTo.ToStringsUsingProcessor(
		true,
		processor,
		[]int{},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor returns empty -- empty anyItems", actual)
}
