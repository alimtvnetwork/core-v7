package coreuniquetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreunique/intunique"
)

// ═══════════════════════════════════════════
// GetMap
// ═══════════════════════════════════════════

func Test_GetMap_NilSlice(t *testing.T) {
	// Arrange
	result := intunique.GetMap(nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "GetMap returns nil -- nil input", actual)
}

func Test_GetMap_EmptySlice(t *testing.T) {
	// Arrange
	input := []int{}
	result := intunique.GetMap(&input)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"len": len(*result),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetMap returns empty map -- empty slice", actual)
}

func Test_GetMap_WithDuplicates(t *testing.T) {
	// Arrange
	input := []int{1, 2, 2, 3, 3, 3}
	result := intunique.GetMap(&input)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"len":    len(*result),
		"has1":   (*result)[1],
		"has2":   (*result)[2],
		"has3":   (*result)[3],
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len":    3,
		"has1":   true,
		"has2":   true,
		"has3":   true,
	}
	expected.ShouldBeEqual(t, 0, "GetMap returns unique map -- with duplicates", actual)
}

// ═══════════════════════════════════════════
// Get — additional branch coverage
// ═══════════════════════════════════════════

func Test_Get_NilSlice(t *testing.T) {
	// Arrange
	result := intunique.Get(nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Get returns nil -- nil input", actual)
}

func Test_Get_SingleElement(t *testing.T) {
	// Arrange
	input := []int{42}
	result := intunique.Get(&input)

	// Act
	actual := args.Map{
		"len": len(*result),
		"first": (*result)[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": 42,
	}
	expected.ShouldBeEqual(t, 0, "Get returns same -- single element", actual)
}
