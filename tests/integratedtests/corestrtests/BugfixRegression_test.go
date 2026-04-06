package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: Hashset.AddNonEmpty — regression for no-op bug
// ==========================================================================

func Test_Hashset_AddNonEmpty_Regression(t *testing.T) {
	safeTest(t, "Test_Hashset_AddNonEmpty_Regression", func() {
		// Arrange
		// Case 0: non-empty adds
		{
			tc := hashsetAddNonEmptyAddsTestCase
			hs := corestr.New.Hashset.Empty()
			hs.AddNonEmpty("hello")

		// Act
			actual := args.Map{
				"length":       fmt.Sprintf("%d", hs.Length()),
				"containsItem": fmt.Sprintf("%v", hs.Has("hello")),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: empty string skipped
		{
			tc := hashsetAddNonEmptySkipsEmptyTestCase
			hs := corestr.New.Hashset.Empty()
			hs.AddNonEmpty("")
			tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", hs.Length()))
		}

		// Case 2: chained
		{
			tc := hashsetAddNonEmptyChainedTestCase
			hs := corestr.New.Hashset.Empty()
			hs.AddNonEmpty("a").AddNonEmpty("b").AddNonEmpty("c")
			actual := args.Map{
				"length":        fmt.Sprintf("%d", hs.Length()),
				"containsItem1": fmt.Sprintf("%v", hs.Has("a")),
				"containsItem2": fmt.Sprintf("%v", hs.Has("b")),
				"containsItem3": fmt.Sprintf("%v", hs.Has("c")),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: SimpleSlice.InsertAt — regression for not-persisting + no bounds
// ==========================================================================

func Test_SimpleSlice_InsertAt_Regression(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_InsertAt_Regression", func() {
		// Arrange
		// Case 0: middle insert
		{
			tc := simpleSliceInsertAtMiddleTestCase
			ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
			ss.InsertAt(1, "X")
			items := ss.List()

		// Act
			actual := args.Map{
				"length": fmt.Sprintf("%d", ss.Length()),
				"item0":  items[0],
				"item1":  items[1],
				"item2":  items[2],
				"item3":  items[3],
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: prepend
		{
			tc := simpleSliceInsertAtPrependTestCase
			ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
			ss.InsertAt(0, "X")
			items := ss.List()
			actual := args.Map{
				"length": fmt.Sprintf("%d", ss.Length()),
				"item0":  items[0],
				"item1":  items[1],
				"item2":  items[2],
				"item3":  items[3],
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: append at end
		{
			tc := simpleSliceInsertAtAppendTestCase
			ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
			ss.InsertAt(3, "X")
			items := ss.List()
			actual := args.Map{
				"length": fmt.Sprintf("%d", ss.Length()),
				"item0":  items[0],
				"item1":  items[1],
				"item2":  items[2],
				"item3":  items[3],
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 3: negative index — no change
		{
			tc := simpleSliceInsertAtNegativeTestCase
			ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
			ss.InsertAt(-1, "X")
			items := ss.List()
			actual := args.Map{
				"length": fmt.Sprintf("%d", ss.Length()),
				"item0":  items[0],
				"item1":  items[1],
				"item2":  items[2],
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 4: out-of-bounds — no change
		{
			tc := simpleSliceInsertAtOutOfBoundsTestCase
			ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})
			ss.InsertAt(100, "X")
			items := ss.List()
			actual := args.Map{
				"length": fmt.Sprintf("%d", ss.Length()),
				"item0":  items[0],
				"item1":  items[1],
				"item2":  items[2],
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: Collection.RemoveAt — regression for inverted guard
// ==========================================================================

func Test_Collection_RemoveAt_Regression(t *testing.T) {
	safeTest(t, "Test_Collection_RemoveAt_Regression", func() {
		// Arrange
		// Case 0: middle
		{
			tc := collectionRemoveAtMiddleTestCase
			col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
			ok := col.RemoveAt(1)

		// Act
			actual := args.Map{
				"isRemoved":       fmt.Sprintf("%v", ok),
				"remainingLength": fmt.Sprintf("%d", col.Length()),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: first
		{
			tc := collectionRemoveAtFirstTestCase
			col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
			ok := col.RemoveAt(0)
			actual := args.Map{
				"isRemoved":       fmt.Sprintf("%v", ok),
				"remainingLength": fmt.Sprintf("%d", col.Length()),
				"newFirstItem":    col.First(),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: last
		{
			tc := collectionRemoveAtLastTestCase
			col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
			ok := col.RemoveAt(2)
			actual := args.Map{
				"isRemoved":       fmt.Sprintf("%v", ok),
				"remainingLength": fmt.Sprintf("%d", col.Length()),
				"lastItem":        col.Last(),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 3: negative
		{
			tc := collectionRemoveAtNegativeTestCase
			col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
			ok := col.RemoveAt(-1)
			actual := args.Map{
				"isRemoved":       fmt.Sprintf("%v", ok),
				"remainingLength": fmt.Sprintf("%d", col.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 4: out-of-bounds
		{
			tc := collectionRemoveAtOutOfBoundsTestCase
			col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
			ok := col.RemoveAt(100)
			actual := args.Map{
				"isRemoved":       fmt.Sprintf("%v", ok),
				"remainingLength": fmt.Sprintf("%d", col.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 5: empty
		{
			tc := collectionRemoveAtEmptyTestCase
			col := corestr.New.Collection.Empty()
			ok := col.RemoveAt(0)
			actual := args.Map{
				"isRemoved":       fmt.Sprintf("%v", ok),
				"remainingLength": fmt.Sprintf("%d", col.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: Hashmap.IsEqualPtr — regression for inverted comparison
// ==========================================================================

func Test_Hashmap_IsEqualPtr_Regression(t *testing.T) {
	safeTest(t, "Test_Hashmap_IsEqualPtr_Regression", func() {
		// Case 0: same keys same values
		{
			tc := hashmapIsEqualPtrSameTestCase
			hm1 := corestr.New.Hashmap.Empty()
			hm1.Set("a", "1")
			hm1.Set("b", "2")
			hm2 := corestr.New.Hashmap.Empty()
			hm2.Set("a", "1")
			hm2.Set("b", "2")

		// Assert
			tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
		}

		// Case 1: same keys different values
		{
			tc := hashmapIsEqualPtrDiffValTestCase
			hm1 := corestr.New.Hashmap.Empty()
			hm1.Set("a", "1")
			hm2 := corestr.New.Hashmap.Empty()
			hm2.Set("a", "DIFFERENT")
			tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
		}

		// Case 2: different keys
		{
			tc := hashmapIsEqualPtrDiffKeysTestCase
			hm1 := corestr.New.Hashmap.Empty()
			hm1.Set("a", "1")
			hm2 := corestr.New.Hashmap.Empty()
			hm2.Set("z", "1")
			tc.ShouldBeEqual(t, 2, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
		}

		// Case 3: both empty
		{
			tc := hashmapIsEqualPtrBothEmptyTestCase
			hm1 := corestr.New.Hashmap.Empty()
			hm2 := corestr.New.Hashmap.Empty()
			tc.ShouldBeEqual(t, 3, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
		}

		// Case 4: nil vs non-nil
		{
			tc := hashmapIsEqualPtrNilVsNonNilTestCase
			var hm1 *corestr.Hashmap
			hm2 := corestr.New.Hashmap.Empty()
			tc.ShouldBeEqual(t, 4, fmt.Sprintf("%v", hm1.IsEqualPtr(hm2)))
		}
	})
}

// ==========================================================================
// Test: Caching removal — IsEmpty/Length on fresh instances
// ==========================================================================

func Test_Caching_Removal_Regression(t *testing.T) {
	safeTest(t, "Test_Caching_Removal_Regression", func() {
		// Arrange
		// Case 0: fresh Hashset
		{
			tc := cachingRemovalFreshHashsetTestCase
			hs := corestr.New.Hashset.Empty()

		// Act
			actual := args.Map{
				"isEmpty": fmt.Sprintf("%v", hs.IsEmpty()),
				"length":  fmt.Sprintf("%d", hs.Length()),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: Hashset after Add
		{
			tc := cachingRemovalHashsetAfterAddTestCase
			hs := corestr.New.Hashset.Empty()
			hs.Add("a").Add("b")
			actual := args.Map{
				"isEmpty": fmt.Sprintf("%v", hs.IsEmpty()),
				"length":  fmt.Sprintf("%d", hs.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: fresh Hashmap
		{
			tc := cachingRemovalFreshHashmapTestCase
			hm := corestr.New.Hashmap.Empty()
			actual := args.Map{
				"isEmpty": fmt.Sprintf("%v", hm.IsEmpty()),
				"length":  fmt.Sprintf("%d", hm.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 3: Hashmap after Set
		{
			tc := cachingRemovalHashmapAfterSetTestCase
			hm := corestr.New.Hashmap.Empty()
			hm.Set("x", "1")
			hm.Set("y", "2")
			actual := args.Map{
				"isEmpty": fmt.Sprintf("%v", hm.IsEmpty()),
				"length":  fmt.Sprintf("%d", hm.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: SimpleSlice.Skip/Take — regression for bounds protection
// ==========================================================================

func Test_SimpleSlice_SkipTake_Regression(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_SkipTake_Regression", func() {
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Case 0: Skip beyond length
		{
			tc := simpleSliceSkipBeyondTestCase
			result := ss.Skip(100)

		// Assert
			tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", len(result)))
		}

		// Case 1: Take beyond length
		{
			tc := simpleSliceTakeBeyondTestCase
			result := ss.Take(100)
			tc.ShouldBeEqual(t, 1, fmt.Sprintf("%d", len(result)))
		}

		// Case 2: Skip 0
		{
			tc := simpleSliceSkipZeroTestCase
			result := ss.Skip(0)
			tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", len(result)))
		}

		// Case 3: Take 0
		{
			tc := simpleSliceTakeZeroTestCase
			result := ss.Take(0)
			tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", len(result)))
		}
	})
}

// ==========================================================================
// Test: HasIndex — regression for negative index guard
// ==========================================================================

func Test_HasIndex_Negative_Regression(t *testing.T) {
	safeTest(t, "Test_HasIndex_Negative_Regression", func() {
		// Case 0: SimpleSlice
		{
			tc := hasIndexNegativeSimpleSliceTestCase
			ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Assert
			tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", ss.HasIndex(-1)))
		}

		// Case 1: Collection
		{
			tc := hasIndexNegativeCollectionTestCase
			col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
			tc.ShouldBeEqual(t, 1, fmt.Sprintf("%v", col.HasIndex(-1)))
		}
	})
}

// ==========================================================================
// Test: Hashmap.Clear nil safety — regression for nil panic
// ==========================================================================

func Test_Hashmap_Clear_NilSafety_Regression(t *testing.T) {
	safeTest(t, "Test_Hashmap_Clear_NilSafety_Regression", func() {
		// Arrange
		// Case 0: nil receiver
		{
			tc := hashmapClearNilReceiverTestCase
			var hm *corestr.Hashmap
			result := hm.Clear()

		// Assert
			tc.ShouldBeEqual(t, 0, fmt.Sprintf("%v", result == nil))
		}

		// Case 1: populated hashmap clears to empty
		{
			tc := hashmapClearPopulatedTestCase
			hm := corestr.New.Hashmap.Empty()
			hm.Set("a", "1")
			hm.Set("b", "2")
			hm.Clear()

		// Act
			actual := args.Map{
				"length":  fmt.Sprintf("%d", hm.Length()),
				"isEmpty": fmt.Sprintf("%v", hm.IsEmpty()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: chainability after Clear
		{
			tc := hashmapClearChainableTestCase
			hm := corestr.New.Hashmap.Empty()
			hm.Set("x", "old")
			hm.Clear().Set("y", "new")
			actual := args.Map{
				"lengthAfterClear": fmt.Sprintf("%d", hm.Length()),
				"lengthAfterReAdd": fmt.Sprintf("%d", len(hm.ValuesList())),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: Hashset.AddBool cache invalidation — regression for stale cache
// ==========================================================================

func Test_Hashset_AddBool_CacheInvalidation_Regression(t *testing.T) {
	safeTest(t, "Test_Hashset_AddBool_CacheInvalidation_Regression", func() {
		// Arrange
		// Case 0: new item invalidates cache
		{
			tc := hashsetAddBoolNewItemTestCase
			hs := corestr.New.Hashset.Empty()
			isExist := hs.AddBool("hello")
			// Force cache rebuild by calling Items
			items := hs.Items()

		// Act
			actual := args.Map{
				"existedBefore": fmt.Sprintf("%v", isExist),
				"lengthAfter":   fmt.Sprintf("%d", len(items)),
				"itemsContains": fmt.Sprintf("%v", hs.Has("hello")),
			}

		// Assert
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 1: existing item returns true, no length change
		{
			tc := hashsetAddBoolExistingTestCase
			hs := corestr.New.Hashset.Empty()
			hs.Add("hello")
			isExist := hs.AddBool("hello")
			actual := args.Map{
				"existedBefore": fmt.Sprintf("%v", isExist),
				"lengthAfter":   fmt.Sprintf("%d", hs.Length()),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: multiple new items all reflected in Items()
		{
			tc := hashsetAddBoolMultipleTestCase
			hs := corestr.New.Hashset.Empty()
			hs.AddBool("a")
			hs.AddBool("b")
			hs.AddBool("c")
			actual := args.Map{
				"length":        fmt.Sprintf("%d", hs.Length()),
				"containsItem1": fmt.Sprintf("%v", hs.Has("a")),
				"containsItem2": fmt.Sprintf("%v", hs.Has("b")),
				"containsItem3": fmt.Sprintf("%v", hs.Has("c")),
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}
	})
}

// ==========================================================================
// Test: Hashmap.AddOrUpdateCollection length mismatch — regression
// ==========================================================================

func Test_Hashmap_AddOrUpdateCollection_LengthMismatch_Regression(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateCollection_LengthMismatch_Regression", func() {
		// Arrange
		// Case 0: mismatched lengths
		{
			tc := hashmapAddOrUpdateMismatchedTestCase
			hm := corestr.New.Hashmap.Empty()
			keys := corestr.New.Collection.Strings([]string{"k1", "k2", "k3"})
			values := corestr.New.Collection.Strings([]string{"v1", "v2"})
			hm.AddOrUpdateCollection(keys, values)

		// Assert
			tc.ShouldBeEqual(t, 0, fmt.Sprintf("%d", hm.Length()))
		}

		// Case 1: equal lengths adds all
		{
			tc := hashmapAddOrUpdateEqualTestCase
			hm := corestr.New.Hashmap.Empty()
			keys := corestr.New.Collection.Strings([]string{"k1", "k2"})
			values := corestr.New.Collection.Strings([]string{"v1", "v2"})
			hm.AddOrUpdateCollection(keys, values)
			v1, _ := hm.Get("k1")
			v2, _ := hm.Get("k2")

		// Act
			actual := args.Map{
				"length": fmt.Sprintf("%d", hm.Length()),
				"value1": v1,
				"value2": v2,
			}
			tc.ShouldBeEqualMapFirst(t, actual)
		}

		// Case 2: nil keys
		{
			tc := hashmapAddOrUpdateNilKeysTestCase
			hm := corestr.New.Hashmap.Empty()
			values := corestr.New.Collection.Strings([]string{"v1"})
			hm.AddOrUpdateCollection(nil, values)
			tc.ShouldBeEqual(t, 2, fmt.Sprintf("%d", hm.Length()))
		}

		// Case 3: empty keys
		{
			tc := hashmapAddOrUpdateEmptyKeysTestCase
			hm := corestr.New.Hashmap.Empty()
			keys := corestr.New.Collection.Empty()
			values := corestr.New.Collection.Strings([]string{"v1"})
			hm.AddOrUpdateCollection(keys, values)
			tc.ShouldBeEqual(t, 3, fmt.Sprintf("%d", hm.Length()))
		}
	})
}
