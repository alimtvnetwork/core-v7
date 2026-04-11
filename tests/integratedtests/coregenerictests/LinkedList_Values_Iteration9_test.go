package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedList.Values() — multi-element iteration
// Covers LinkedListIter.go L36 (current = current.next)
// ══════════════════════════════════════════════════════════════════════════════

func Test_LinkedList_Values_MultiElement(t *testing.T) {
	// Arrange
	tc := cov5ValuesIterMultiElementTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	// Act
	collected := make([]int, 0, 3)
	for v := range ll.Values() {
		collected = append(collected, v)
	}

	actual := args.Map{
		"count":  len(collected),
		"first":  collected[0],
		"second": collected[1],
		"third":  collected[2],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LinkedList_Values_BreakEarly_FromLinkedListValuesIter(t *testing.T) {
	// Arrange
	tc := cov5ValuesIterBreakEarlyTestCase
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})

	// Act
	collected := make([]int, 0, 1)
	for v := range ll.Values() {
		collected = append(collected, v)
		break
	}

	actual := args.Map{
		"count": len(collected),
		"first": collected[0],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
