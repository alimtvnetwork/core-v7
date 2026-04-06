package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage31 — LinkedCollection & LinkedList tests (Iteration 28)
//
// API Reference (verified from source):
//   - corestr.New.LinkedCollection.Create()    (singular, NOT LinkedCollections)
//   - corestr.New.Collection.Strings([]string{...})  (takes a slice)
//   - lc.IsEqualsPtr(other)                    (NOT IsChainEqual)
//   - LinkedCollections.SafeIndexAt(index)     (no SafeIndexAtLock on LinkedCollections)
//   - LinkedList.SafeIndexAtLock(index)        (exists on LinkedList)
// ══════════════════════════════════════════════════════════════════════════════

// ---------- LinkedCollection: SafeIndexAt out-of-range ----------

func Test_I28_LinkedCollections_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_SafeIndexAt_OutOfRange", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c1)

		// Act
		node := lc.SafeIndexAt(999)

		// Assert
		actual := args.Map{"isNil": node == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt returns nil -- out-of-range index", actual)
	})
}

// ---------- LinkedCollection: SafeIndexAt with Lock out-of-range ----------

func Test_I28_LinkedCollections_SafeIndexAtLock_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_SafeIndexAtLock_OutOfRange", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c1)

		// Act — manually lock/unlock since SafeIndexAtLock may not exist
		var node *corestr.LinkedCollectionNode
		func() {
			lc.Lock()
			defer lc.Unlock()
			node = lc.SafeIndexAt(999)
		}()

		// Assert
		actual := args.Map{"isNil": node == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtLock returns nil -- out-of-range index", actual)
	})
}

// ---------- LinkedCollection: IsEqualsPtr both empty ----------

func Test_I28_LinkedCollections_IsChainEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_IsChainEqual_BothEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc2 := corestr.New.LinkedCollection.Create()

		// Act — IsEqualsPtr (NOT IsChainEqual, which doesn't exist)
		result := lc1.IsEqualsPtr(lc2)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr returns true -- both empty", actual)
	})
}

// ---------- LinkedCollection: IsEqualsPtr one empty ----------

func Test_I28_LinkedCollections_IsChainEqual_OneEmpty(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_IsChainEqual_OneEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		result := lc1.IsEqualsPtr(lc2)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr returns false -- one empty", actual)
	})
}

// ---------- LinkedCollection: ToCollection with items ----------

func Test_I28_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_ToCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		result := lc.ToCollection(0)

		// Assert
		actual := args.Map{"length": result.Length()}
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "ToCollection merges all -- with items", actual)
	})
}

// ---------- LinkedCollection: ToCollectionsOfCollection with items ----------

func Test_I28_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_ToCollectionsOfCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		lc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		result := lc.ToCollectionsOfCollection(0)

		// Assert
		actual := args.Map{"length": result.Length()}
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "ToCollectionsOfCollection returns collections -- with items", actual)
	})
}

// ---------- LinkedCollection: AddCollections all nil ----------

func Test_I28_LinkedCollections_AddCollections_AllNil(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_AddCollections_AllNil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		result := lc.AddCollections([]*corestr.Collection{nil, nil, nil})

		// Assert
		actual := args.Map{"isEmpty": result.IsEmpty()}
		expected := args.Map{"isEmpty": true}
		expected.ShouldBeEqual(t, 0, "AddCollections returns self -- all nil", actual)
	})
}

// ---------- LinkedListNode: IsChainEqual both nil ----------

func Test_I28_LinkedListNode_IsChainEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_I28_LinkedListNode_IsChainEqual_BothNil", func() {
		// Arrange
		var n1 *corestr.LinkedListNode
		var n2 *corestr.LinkedListNode

		// Act
		result := n1.IsChainEqual(n2, true)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": true}
		expected.ShouldBeEqual(t, 0, "IsChainEqual returns true -- both nil", actual)
	})
}

// ---------- LinkedListNode: IsChainEqual one nil ----------

func Test_I28_LinkedListNode_IsChainEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_I28_LinkedListNode_IsChainEqual_OneNil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		n1 := ll.Head()

		// Act
		result := n1.IsChainEqual(nil, true)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsChainEqual returns false -- one nil", actual)
	})
}

// ---------- LinkedList: SafeIndexAt out-of-range ----------

func Test_I28_LinkedList_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_LinkedList_SafeIndexAt_OutOfRange", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		node := ll.SafeIndexAt(999)

		// Assert
		actual := args.Map{"isNil": node == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt returns nil -- out-of-range", actual)
	})
}

// ---------- LinkedList: SafeIndexAtLock out-of-range ----------

func Test_I28_LinkedList_SafeIndexAtLock_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_LinkedList_SafeIndexAtLock_OutOfRange", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		node := ll.SafeIndexAtLock(999)

		// Assert
		actual := args.Map{"isNil": node == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtLock returns nil -- out-of-range", actual)
	})
}

// ---------- CollectionsOfCollection: AllIndividualItemsLength with empty collection ----------

func Test_I28_CollectionsOfCollection_AllIndividualItemsLength_EmptyCollectionSkip(t *testing.T) {
	safeTest(t, "Test_I28_CollectionsOfCollection_AllIndividualItemsLength_EmptyCollectionSkip", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(4)
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		coc.Add(corestr.New.Collection.Cap(0))
		coc.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		result := coc.AllIndividualItemsLength()

		// Assert
		actual := args.Map{"sum": result}
		expected := args.Map{"sum": 3}
		expected.ShouldBeEqual(t, 0, "AllIndividualItemsLength skips empty -- mixed collections", actual)
	})
}

// ---------- CollectionsOfCollection: List nil/empty skip ----------

func Test_I28_CollectionsOfCollection_List_NilEmptySkip(t *testing.T) {
	safeTest(t, "Test_I28_CollectionsOfCollection_List_NilEmptySkip", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(4)
		coc.Add(corestr.New.Collection.Strings([]string{"a"}))
		coc.Add(nil)
		coc.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		result := coc.List(0)

		// Assert
		actual := args.Map{"length": len(result)}
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "List skips nil -- mixed items", actual)
	})
}
