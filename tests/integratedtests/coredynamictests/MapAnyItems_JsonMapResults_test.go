package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MapAnyItems.JsonMapResults (line 903) ──

func Test_MapAnyItems_JsonMapResults_Valid_I29(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(3)
	m.Add("key1", "val1")
	m.Add("key2", 42)

	// Act
	results, err := m.JsonMapResults()

	// Assert
	actual := args.Map{
		"notNil":   results != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "MapAnyItems JsonMapResults valid", expected)
}

// ── MapAnyItems.GetUsingUnmarshallAt type mismatch (line 350) ──

func Test_MapAnyItems_GetUsingUnmarshallAt_TypeMismatch_I29(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(3)
	m.Add("key1", "stringValue")

	var target int

	// Act
	err := m.GetUsingUnmarshallAt("key1", &target)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "MapAnyItems GetUsingUnmarshallAt type mismatch", expected)
}
