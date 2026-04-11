package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TypedDynamic.JsonString valid (line 117) ──

func Test_TypedDynamic_JsonString_Valid_I29(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[string]("hello", true)

	// Act
	jsonStr, err := td.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "TypedDynamic JsonString valid", expected)
}
