package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/keymk"
)

// Test_Cov5_AppendStringsWithBaseAnyItems_SkipEmpty covers
// keymk/appendStringsWithBaseAnyItems.go L13-14: skip empty entry with IsSkipEmptyEntry.
func Test_AppendStringsWithBaseAnyItems_SkipEmpty(t *testing.T) {
	// Arrange
	key := keymk.NewKey.Create(
		keymk.JoinerOption, // IsSkipEmptyEntry=true
		"root",
	)

	// Act — pass empty string as any item; it should be skipped
	result := key.Compile("a", "", "b")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": "root-a-b"}
	expected.ShouldBeEqual(t, 0, "appendStringsWithBaseAnyItems returns empty -- skip empty", actual)
}
