package coregeneric

import "testing"

// Test_IndexAt_ReturnsNil_WhenInternalChainShorterThanLength
// covers the defensive guard at LinkedList.go:319-321 where
// the node chain ends before reaching the requested index
// (length field is inconsistent with actual chain).
func Test_IndexAt_ReturnsNil_WhenInternalChainShorterThanLength(t *testing.T) {
	// Arrange
	ll := EmptyLinkedList[string]()
	ll.Add("a")
	// Corrupt internal state: set length larger than actual chain
	ll.length = 5

	// Act
	result := ll.IndexAt(3)

	// Assert
	if result != nil {
		t.Errorf("expected nil when chain is shorter than length, got %v", result)
	}
}
