package coreindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreindexes"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestHasIndex verifies index existence.
func TestHasIndex(t *testing.T) {
	indexes := []int{1, 3, 5}
	actual := args.Map{"result": coreindexes.HasIndex(indexes, 3)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3 should be found", actual)
	actual := args.Map{"result": coreindexes.HasIndex(indexes, 2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "2 should not be found", actual)
}

// TestIsInvalidIndex verifies invalid index.
func TestIsInvalidIndex(t *testing.T) {
	actual := args.Map{"result": coreindexes.IsInvalidIndex(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "-1 should be invalid", actual)
	actual := args.Map{"result": coreindexes.IsInvalidIndex(0)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "0 should be valid", actual)
}

// TestIsWithinIndexRange verifies index range.
func TestIsWithinIndexRange(t *testing.T) {
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(2, 5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "index 2 should be within length 5", actual)
	actual := args.Map{"result": coreindexes.IsWithinIndexRange(5, 5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "index 5 should not be within length 5", actual)
}

// TestLastIndex verifies last index.
func TestLastIndex(t *testing.T) {
	actual := args.Map{"result": coreindexes.LastIndex(5) != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "last index of 5 should be 4", actual)
	actual := args.Map{"result": coreindexes.LastIndex(0) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "last index of 0 should be -1", actual)
}

// TestNameByIndex verifies name lookup.
func TestNameByIndex(t *testing.T) {
	actual := args.Map{"result": coreindexes.NameByIndex(0) != "First"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'First', got ''", actual)
	actual := args.Map{"result": coreindexes.NameByIndex(9) != "Tenth"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Tenth', got ''", actual)
	actual := args.Map{"result": coreindexes.NameByIndex(99) != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of range should return empty", actual)
}

// TestOf verifies index-of search.
func TestOf(t *testing.T) {
	indexes := []int{10, 20, 30}
	actual := args.Map{"result": coreindexes.Of(indexes, 20) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "20 should be at position 1", actual)
	actual := args.Map{"result": coreindexes.Of(indexes, 99) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "99 should return -1", actual)
}

// TestSafeEndingIndex verifies safe ending.
func TestSafeEndingIndex(t *testing.T) {
	actual := args.Map{"result": coreindexes.SafeEndingIndex(5, 3) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "within range should return lastTaking", actual)
	actual := args.Map{"result": coreindexes.SafeEndingIndex(3, 5) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "exceeding should return lastIndex", actual)
}

// TestHasIndexPlusRemoveIndex verifies find-and-remove.
func TestHasIndexPlusRemoveIndex(t *testing.T) {
	indexes := []int{1, 2, 3}
	actual := args.Map{"result": coreindexes.HasIndexPlusRemoveIndex(&indexes, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2 should be found", actual)
	actual := args.Map{"result": len(indexes) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 items", actual)
	actual := args.Map{"result": coreindexes.HasIndexPlusRemoveIndex(&indexes, 99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "99 should not be found", actual)
}

// TestConstants verifies index constants.
func TestConstants(t *testing.T) {
	actual := args.Map{"result": coreindexes.First != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "First should be 0", actual)
	actual := args.Map{"result": coreindexes.Second != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Second should be 1", actual)
	actual := args.Map{"result": coreindexes.Tenth != 9}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Tenth should be 9", actual)
	actual := args.Map{"result": coreindexes.Index0 != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Index0 should be 0", actual)
	actual := args.Map{"result": coreindexes.Index20 != 20}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Index20 should be 20", actual)
}
