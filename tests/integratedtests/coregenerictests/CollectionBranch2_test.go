package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: Collection — RemoveAt edge cases
// ==========================================================================

func Test_Collection_RemoveAt_Middle(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtMiddleTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(2)

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
		"first":   col.First(),
		"last":    col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_First(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtFirstTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(0)

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
		"first":   col.First(),
		"last":    col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_Last(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtLastTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(4)

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
		"first":   col.First(),
		"last":    col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_Negative(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtNegativeTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(-1)

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_OutOfBounds(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtOutOfBoundsTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	ok := col.RemoveAt(100)

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_RemoveAt_Empty(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	ok := col.RemoveAt(0)

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Reverse
// ==========================================================================

func Test_Collection_Reverse_Populated(t *testing.T) {
	// Arrange
	tc := collectionReversePopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	col.Reverse()

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Reverse_Single(t *testing.T) {
	// Arrange
	tc := collectionReverseSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	col.Reverse()

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Reverse_Empty(t *testing.T) {
	// Arrange
	tc := collectionReverseEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	col.Reverse()

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — FirstOrDefault
// ==========================================================================

func Test_Collection_FirstOrDefault_Populated(t *testing.T) {
	// Arrange
	tc := collectionFirstOrDefaultPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{"result": col.FirstOrDefault()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := collectionFirstOrDefaultEmptyTestCase
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": col.FirstOrDefault()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — LastOrDefault
// ==========================================================================

func Test_Collection_LastOrDefault_Populated(t *testing.T) {
	// Arrange
	tc := collectionLastOrDefaultPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{"result": col.LastOrDefault()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	tc := collectionLastOrDefaultEmptyTestCase
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": col.LastOrDefault()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — SafeAt
// ==========================================================================

func Test_Collection_SafeAt_Valid(t *testing.T) {
	// Arrange
	tc := collectionSafeAtValidTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{"result": col.SafeAt(1)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SafeAt_Negative(t *testing.T) {
	// Arrange
	tc := collectionSafeAtNegativeTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{"result": col.SafeAt(-1)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SafeAt_OutOfBounds(t *testing.T) {
	// Arrange
	tc := collectionSafeAtOutOfBoundsTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{"result": col.SafeAt(100)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SafeAt_Empty(t *testing.T) {
	// Arrange
	tc := collectionSafeAtEmptyTestCase
	empty := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": empty.SafeAt(0)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — ConcatNew
// ==========================================================================

func Test_Collection_ConcatNew_Populated(t *testing.T) {
	// Arrange
	tc := collectionConcatNewPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := col.ConcatNew(4, 5)

	// Act
	actual := args.Map{
		"resultLength": result.Length(),
		"resultFirst":  result.First(),
		"resultLast":   result.Last(),
		"origLength":   col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_ConcatNew_Empty(t *testing.T) {
	// Arrange
	tc := collectionConcatNewEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	result := col.ConcatNew(10, 20)

	// Act
	actual := args.Map{
		"length": result.Length(),
		"first":  result.First(),
		"last":   result.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddIf
// ==========================================================================

func Test_Collection_AddIf_True(t *testing.T) {
	// Arrange
	tc := collectionAddIfTrueTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(true, 42)

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_AddIf_False(t *testing.T) {
	// Arrange
	tc := collectionAddIfFalseTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(false, 42)

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — ForEachBreak on empty
// ==========================================================================

func Test_Collection_ForEachBreak_Empty(t *testing.T) {
	// Arrange
	tc := collectionForEachBreakEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	visited := 0
	col.ForEachBreak(func(index int, item int) bool {
		visited++
		return false
	})

	// Act
	actual := args.Map{"visited": visited}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddSlice
// ==========================================================================

func Test_Collection_AddSlice_Populated(t *testing.T) {
	// Arrange
	tc := collectionAddSlicePopulatedTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddSlice([]int{10, 20, 30})

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_AddSlice_Empty(t *testing.T) {
	// Arrange
	tc := collectionAddSliceEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddSlice([]int{})

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Items / ItemsPtr
// ==========================================================================

func Test_Collection_Items_Slice(t *testing.T) {
	// Arrange
	tc := collectionItemsSliceTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	items := col.Items()

	// Act
	actual := args.Map{
		"length": len(items),
		"first":  items[0],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_ItemsPtr(t *testing.T) {
	// Arrange
	tc := collectionItemsPtrTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	ptr := col.ItemsPtr()

	// Act
	actual := args.Map{"isNotNil": ptr != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — RemoveAtLock
// ==========================================================================

func Test_Collection_RemoveAtLock(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtLockTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	col.Lock()
	ok := col.RemoveAt(1)
	col.Unlock()

	// Act
	actual := args.Map{
		"removed": ok,
		"length":  col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
