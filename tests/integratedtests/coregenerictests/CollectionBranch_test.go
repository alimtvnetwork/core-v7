package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: Collection — ForEach
// ==========================================================================

func Test_Collection_ForEach_VisitsAll(t *testing.T) {
	// Arrange
	tc := collectionForEachVisitsAllTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	visited := 0
	var firstEntry, lastEntry string

	col.ForEach(func(index int, item int) {
		if visited == 0 {
			firstEntry = fmt.Sprintf("%d:%d", index, item)
		}
		lastEntry = fmt.Sprintf("%d:%d", index, item)
		visited++
	})

	// Act
	actual := args.Map{
		"visited":    visited,
		"firstEntry": firstEntry,
		"lastEntry":  lastEntry,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_ForEach_Empty(t *testing.T) {
	// Arrange
	tc := collectionForEachEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	visited := 0
	col.ForEach(func(index int, item int) { visited++ })

	// Act
	actual := args.Map{"visited": visited}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — ForEachBreak
// ==========================================================================

func Test_Collection_ForEachBreak_Stops(t *testing.T) {
	// Arrange
	tc := collectionForEachBreakStopsTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
	visited := 0
	col.ForEachBreak(func(index int, item int) bool {
		visited++
		return item >= 3
	})

	// Act
	actual := args.Map{"visited": visited}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_ForEachBreak_VisitsAll(t *testing.T) {
	// Arrange
	tc := collectionForEachBreakVisitsAllTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5)
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
// Test: Collection — SortFunc
// ==========================================================================

func Test_Collection_SortFunc_Asc(t *testing.T) {
	// Arrange
	tc := collectionSortFuncAscTestCase
	col := coregeneric.New.Collection.Int.Items(3, 1, 5, 2, 4)
	col.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"first": col.First(),
		"last":  col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SortFunc_Desc(t *testing.T) {
	// Arrange
	tc := collectionSortFuncDescTestCase
	col := coregeneric.New.Collection.Int.Items(3, 1, 5, 2, 4)
	col.SortFunc(func(a, b int) bool { return a > b })

	// Act
	actual := args.Map{
		"first": col.First(),
		"last":  col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SortFunc_Single(t *testing.T) {
	// Arrange
	tc := collectionSortFuncSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	col.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"first": col.First(),
		"last":  col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddIfMany
// ==========================================================================

func Test_Collection_AddIfMany_True(t *testing.T) {
	// Arrange
	tc := collectionAddIfManyTrueTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(true, 10, 20, 30)

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_AddIfMany_False(t *testing.T) {
	// Arrange
	tc := collectionAddIfManyFalseTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(false, 10, 20, 30)

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddFunc
// ==========================================================================

func Test_Collection_AddFunc(t *testing.T) {
	// Arrange
	tc := collectionAddFuncTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddFunc(func() int { return 42 })

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddCollections
// ==========================================================================

func Test_Collection_AddCollections_Merge(t *testing.T) {
	// Arrange
	tc := collectionAddCollectionsMergeTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	c2 := coregeneric.New.Collection.Int.Items(4, 5)
	c3 := coregeneric.New.Collection.Int.Items(6)
	col.AddCollections(c2, c3)

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_AddCollections_WithNil(t *testing.T) {
	// Arrange
	tc := collectionAddCollectionsNilTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	empty := coregeneric.EmptyCollection[int]()
	col.AddCollections(empty)

	// Act
	actual := args.Map{
		"length": col.Length(),
		"first":  col.First(),
		"last":   col.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Clone edge cases
// ==========================================================================

func Test_Collection_Clone_Empty(t *testing.T) {
	// Arrange
	tc := collectionCloneEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	cloned := col.Clone()

	// Act
	actual := args.Map{
		"length":  cloned.Length(),
		"isEmpty": cloned.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Skip/Take boundary
// ==========================================================================

func Test_Collection_SkipAll(t *testing.T) {
	// Arrange
	tc := collectionSkipAllTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	skipped := col.Skip(10)

	// Act
	actual := args.Map{"length": len(skipped)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_TakeMore(t *testing.T) {
	// Arrange
	tc := collectionTakeMoreTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	taken := col.Take(100)

	// Act
	actual := args.Map{"length": len(taken)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_SkipZeroTakeZero(t *testing.T) {
	// Arrange
	tc := collectionSkipZeroTakeZeroTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	skipZero := col.Skip(0)
	takeZero := col.Take(0)

	// Act
	actual := args.Map{
		"skipLength": len(skipZero),
		"takeLength": len(takeZero),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Filter edge cases
// ==========================================================================

func Test_Collection_Filter_NoMatch(t *testing.T) {
	// Arrange
	tc := collectionFilterNoMatchTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	filtered := col.Filter(func(item int) bool { return item > 100 })

	// Act
	actual := args.Map{
		"length":  filtered.Length(),
		"isEmpty": filtered.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Filter_AllMatch(t *testing.T) {
	// Arrange
	tc := collectionFilterAllMatchTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	filtered := col.Filter(func(item int) bool { return item > 0 })

	// Act
	actual := args.Map{"length": filtered.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Filter_Empty(t *testing.T) {
	// Arrange
	tc := collectionFilterEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	filtered := col.Filter(func(item int) bool { return true })

	// Act
	actual := args.Map{
		"length":  filtered.Length(),
		"isEmpty": filtered.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — CountFunc edge cases
// ==========================================================================

func Test_Collection_CountFunc_NoMatch(t *testing.T) {
	// Arrange
	tc := collectionCountFuncNoMatchTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	count := col.CountFunc(func(item int) bool { return item > 100 })

	// Act
	actual := args.Map{"count": count}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_CountFunc_Empty(t *testing.T) {
	// Arrange
	tc := collectionCountFuncEmptyTestCase
	col := coregeneric.EmptyCollection[int]()
	count := col.CountFunc(func(item int) bool { return true })

	// Act
	actual := args.Map{"count": count}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — String output
// ==========================================================================

func Test_Collection_String_Populated(t *testing.T) {
	// Arrange
	tc := collectionStringPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)

	// Act
	actual := args.Map{"result": col.String()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_String_Empty(t *testing.T) {
	// Arrange
	tc := collectionStringEmptyTestCase
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": col.String()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Lock variants
// ==========================================================================

func Test_Collection_Lock_Variants(t *testing.T) {
	// Arrange
	tc := collectionLockVariantsTestCase
	col := coregeneric.EmptyCollection[int]()
	col.AddLock(1)
	col.AddsLock(2, 3)

	// Act
	actual := args.Map{
		"lengthLock":  col.LengthLock(),
		"isEmptyLock": col.IsEmptyLock(),
		"length":      col.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — Metadata methods
// ==========================================================================

func Test_Collection_Metadata_Populated(t *testing.T) {
	// Arrange
	tc := collectionMetadataPopulatedTestCase
	col := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{
		"hasAnyItem": col.HasAnyItem(),
		"hasItems":   col.HasItems(),
		"hasIndex2":  col.HasIndex(2),
		"hasIndex5":  col.HasIndex(5),
		"lastIndex":  col.LastIndex(),
		"count":      col.Count(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Collection_Metadata_Empty(t *testing.T) {
	// Arrange
	tc := collectionMetadataEmptyTestCase
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{
		"hasAnyItem": col.HasAnyItem(),
		"hasItems":   col.HasItems(),
		"hasIndex0":  col.HasIndex(0),
		"lastIndex":  col.LastIndex(),
		"count":      col.Count(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — RemoveAt single item
// ==========================================================================

func Test_Collection_RemoveAt_Single(t *testing.T) {
	// Arrange
	tc := collectionRemoveAtSingleTestCase
	col := coregeneric.New.Collection.Int.Items(42)
	removed := col.RemoveAt(0)

	// Act
	actual := args.Map{
		"removed": removed,
		"length":  col.Length(),
		"isEmpty": col.IsEmpty(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — AddCollection with empty
// ==========================================================================

func Test_Collection_AddCollection_Empty(t *testing.T) {
	// Arrange
	tc := collectionAddCollectionEmptyTestCase
	col := coregeneric.New.Collection.Int.Items(1, 2, 3)
	empty := coregeneric.EmptyCollection[int]()
	col.AddCollection(empty)

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Hashmap — IsEquals
// ==========================================================================

func Test_Hashmap_IsEquals_SameKeys(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsSameKeysTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)
	hm1.Set("b", 2)
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm2.Set("a", 99)
	hm2.Set("b", 100)

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DiffKeys(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsDiffKeysTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)
	hm1.Set("b", 2)
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm2.Set("x", 1)
	hm2.Set("y", 2)

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_DiffLength(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsDiffLengthTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm2.Set("a", 1)
	hm2.Set("b", 2)

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_BothNil(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsBothNilTestCase
	var hm1 *coregeneric.Hashmap[string, int]
	var hm2 *coregeneric.Hashmap[string, int]

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_NilVsNonNil(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsNilVsNonNilTestCase
	var hm1 *coregeneric.Hashmap[string, int]
	hm2 := coregeneric.New.Hashmap.StringInt.Cap(5)

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm2)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Hashmap_IsEquals_SamePtr(t *testing.T) {
	// Arrange
	tc := hashmapIsEqualsSamePtrTestCase
	hm1 := coregeneric.New.Hashmap.StringInt.Cap(5)
	hm1.Set("a", 1)

	// Act
	actual := args.Map{"isEquals": hm1.IsEquals(hm1)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Collection — CollectionLenCap
// ==========================================================================

func Test_Collection_LenCap(t *testing.T) {
	// Arrange
	tc := collectionLenCapTestCase
	col := coregeneric.CollectionLenCap[int](3, 10)

	// Act
	actual := args.Map{
		"length":   col.Length(),
		"capacity": col.Capacity(),
		"first":    col.First(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
