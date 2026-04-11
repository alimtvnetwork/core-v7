package coreuniquetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreunique/intunique"
)

// ============================================================================
// GetMap with duplicates
// ============================================================================

func Test_IntUnique_GetMap_WithDuplicates_Ext(t *testing.T) {
	tc := extGetMapWithDuplicatesTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)
	clone := make([]int, len(slice))
	copy(clone, slice)

	// Act
	result := intunique.GetMap(&clone)

	// Assert
	actual := args.Map{
		"isNil":  result == nil,
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// GetMap nil
// ============================================================================

func Test_IntUnique_GetMap_Nil_Ext(t *testing.T) {
	tc := extGetMapNilTestCase

	// Act
	result := intunique.GetMap(nil)

	// Assert
	actual := args.Map{
		"isNil": result == nil,
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// GetMap empty
// ============================================================================

func Test_IntUnique_GetMap_Empty_Ext(t *testing.T) {
	tc := extGetMapEmptyTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)

	// Act
	result := intunique.GetMap(&slice)

	// Assert
	actual := args.Map{
		"isNil":  result == nil,
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// Get empty slice
// ============================================================================

func Test_IntUnique_Get_EmptySlice_Ext(t *testing.T) {
	tc := extGetEmptySliceTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)

	// Act
	result := intunique.Get(&slice)

	// Assert
	actual := args.Map{
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}

// ============================================================================
// Get single element
// ============================================================================

func Test_IntUnique_Get_SingleElement_Ext(t *testing.T) {
	tc := extGetSingleElementTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	inputVal, _ := input.Get("input")
	slice := inputVal.([]int)

	// Act
	result := intunique.Get(&slice)

	// Assert
	actual := args.Map{
		"length": len(*result),
	}
	tc.ShouldBeEqualMap(t, 0, actual)
}
