package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection.JsonString valid (line 355) ──

func Test_Collection_JsonString_Valid_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](3)
	coll.Add("a")
	coll.Add("b")

	// Act
	jsonStr, err := coll.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "Collection JsonString valid", expected)
}

// ── Collection.JsonStringMust valid (line 364) ──

func Test_Collection_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](2)
	coll.Add("test")

	// Act
	result := coll.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "Collection JsonStringMust valid", expected)
}
