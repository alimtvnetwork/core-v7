package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_NotContains_CaseSensitive(t *testing.T) {
	// Arrange
	nc := stringcompareas.NotContains
	// Case-sensitive: "Hello World" does NOT contain "hello" (different case)
	// so NotContains returns true (it is indeed not contained).
	result := nc.IsCompareSuccess(false, "Hello World", "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "case sensitive: 'hello' not in 'Hello World', NotContains should be true", actual)
}

func Test_NotContains_CaseInsensitive(t *testing.T) {
	// Arrange
	nc := stringcompareas.NotContains
	result := nc.IsCompareSuccess(true, "Hello World", "hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "case insensitive: should contain", actual)
}
