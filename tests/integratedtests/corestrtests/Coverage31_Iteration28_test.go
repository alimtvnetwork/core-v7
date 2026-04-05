package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage31 — corestr remaining gaps (Iteration 28)
//
// Targets:
//   - CharCollectionMap: AddHashmapsValues nil, AddHashmapsKeys nil, AddHashmapsAll nil
//   - CharCollectionMap: AllLengthsSumLock with items, ToHashsetsCollection empty collection skip
//   - CharHashsetMap: AddLock nil items init, Add nil items init, AddAll nil/empty
//   - CharHashsetMap: GetHashset nil items init, GetHashsetLock nil items init
//   - CharHashsetMap: efficientAddOfLargeItems
//   - Collection: resizeForHashmaps nil, resizeForCollections empty, resizeForItems nil, resizeForAnys empty
//   - CollectionsOfCollection: AllLengthsSum empty collection skip, ToStrings nil/empty skip
//   - LinkedCollectionNode: IsChainEqual nil branches, isNextChainEqual nil branches
//   - LinkedCollections: incrementLengthLock, IsChainEqual nil branches
//   - LinkedCollections: SafeIndexAt out-of-range, SafeIndexAtLock out-of-range
//   - LinkedCollections: AddFirstNonNil all nil, ToCollection/ToCollectionsOfCollection nil node
//   - LinkedCollections: ToNonChainedNodes empty result
//   - LinkedList: IsChainEqual nil branches, SafeIndexAt/SafeIndexAtLock out-of-range
//   - LinkedListNode: IsChainEqual nil branches, isNextChainEqual nil branches
//   - NonChainedLinkedCollectionNodes: Length nil
//   - NonChainedLinkedListNodes: Length nil items
//   - SimpleSlice: IsEqualUnorderedLines nil, IsEqualUnorderedLinesClone nil
//   - SimpleSlice: IsEqualByFuncLinesSplit empty match
//   - SimpleStringOnce: ToIntBounded fallthrough
//   - ValidValue: ParseInjectUsingJson error path
//   - Hashmap: safeWaitGroupDone nil
// ══════════════════════════════════════════════════════════════════════════════

// ---------- CharCollectionMap: AddHashmapsValues nil ----------

func Test_I28_CharCollectionMap_AddHashmapsValues_Nil(t *testing.T) {
	safeTest(t, "Test_I28_CharCollectionMap_AddHashmapsValues_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Cap(4)

		// Act
		result := ccm.AddHashmapsValues(nil)

		// Assert
		actual := args.Map{"isNil": result == nil}
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues returns self -- nil input", actual)
	})
}

// ---------- CharCollectionMap: AddHashmapsKeysOrValuesBothUsingFilter nil ----------

func Test_I28_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_I28_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Cap(4)

		// Act
		result := ccm.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)

		// Assert
		actual := args.Map{"isNil": result == nil}
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter returns self -- nil input", actual)
	})
}

// ---------- CharCollectionMap: AddHashmapsKeysValuesBoth nil ----------

func Test_I28_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil(t *testing.T) {
	safeTest(t, "Test_I28_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Cap(4)

		// Act
		result := ccm.AddHashmapsKeysValuesBoth(nil)

		// Assert
		actual := args.Map{"isNil": result == nil}
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesBoth returns self -- nil input", actual)
	})
}

// ---------- CharHashsetMap: AddLock triggers nil items init ----------

func Test_I28_CharHashsetMap_AddLock_NilItemsInit(t *testing.T) {
	safeTest(t, "Test_I28_CharHashsetMap_AddLock_NilItemsInit", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4)

		// Act
		chm.AddLock("apple")

		// Assert
		actual := args.Map{"has": chm.Has("apple")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLock initializes items -- nil items", actual)
	})
}

// ---------- CharHashsetMap: Add triggers nil items init ----------

func Test_I28_CharHashsetMap_Add_NilItemsInit(t *testing.T) {
	safeTest(t, "Test_I28_CharHashsetMap_Add_NilItemsInit", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4)

		// Act
		chm.Add("banana")

		// Assert
		actual := args.Map{"has": chm.Has("banana")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Add initializes items -- nil items", actual)
	})
}

// ---------- CharHashsetMap: AddAll empty/nil ----------

func Test_I28_CharHashsetMap_AddAll_Empty(t *testing.T) {
	safeTest(t, "Test_I28_CharHashsetMap_AddAll_Empty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(4)

		// Act
		result := chm.AddAll(0, nil)

		// Assert
		actual := args.Map{"length": result.Length()}
		expected := args.Map{"length": 0}
		expected.ShouldBeEqual(t, 0, "AddAll returns empty -- nil input", actual)
	})
}

// ---------- SimpleSlice: IsEqualUnorderedLines nil receiver ----------

func Test_I28_SimpleSlice_IsEqualUnorderedLines_NilReceiver(t *testing.T) {
	safeTest(t, "Test_I28_SimpleSlice_IsEqualUnorderedLines_NilReceiver", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.IsEqualUnorderedLines([]string{"a"})

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLines returns false -- nil receiver", actual)
	})
}

// ---------- SimpleSlice: IsEqualUnorderedLinesClone nil receiver ----------

func Test_I28_SimpleSlice_IsEqualUnorderedLinesClone_NilReceiver(t *testing.T) {
	safeTest(t, "Test_I28_SimpleSlice_IsEqualUnorderedLinesClone_NilReceiver", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act
		result := ss.IsEqualUnorderedLinesClone([]string{"a"})

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualUnorderedLinesClone returns false -- nil receiver", actual)
	})
}

// ---------- SimpleSlice: IsEqualByFuncLinesSplit both empty ----------

func Test_I28_SimpleSlice_IsEqualByFuncLinesSplit_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I28_SimpleSlice_IsEqualByFuncLinesSplit_BothEmpty", func() {
		// Arrange
		ss := corestr.SimpleSlice{}

		// Act — empty left with empty-split right
		result := ss.IsEqualByFuncLinesSplit(
			false,
			",",
			"",
			func(index int, left, right string) bool {
				return left == right
			},
		)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsEqualByFuncLinesSplit returns false -- empty vs single empty string", actual)
	})
}

// ---------- ValidValue: ParseInjectUsingJson error path ----------

func Test_I28_ValidValue_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I28_ValidValue_ParseInjectUsingJson_Error", func() {
		// Arrange
		vv := &corestr.ValidValue{}
		badJson := corejson.NewPtr([]byte("not-valid-json{{{"))

		// Act
		result, err := vv.ParseInjectUsingJson(badJson)

		// Assert
		actual := args.Map{
			"resultNil": result == nil,
			"hasErr":    err != nil,
		}
		expected := args.Map{
			"resultNil": true,
			"hasErr":    true,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns nil,err -- invalid json", actual)
	})
}

// ---------- LinkedCollections: SafeIndexAt out-of-range ----------

func Test_I28_LinkedCollections_SafeIndexAt_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_SafeIndexAt_OutOfRange", func() {
		// Arrange
		lc := corestr.New.LinkedCollections.Create()
		c1 := corestr.New.Collection.Strings("a", "b")
		lc.Add(c1)

		// Act
		node := lc.SafeIndexAt(999)

		// Assert
		actual := args.Map{"isNil": node == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt returns nil -- out-of-range index", actual)
	})
}

// ---------- LinkedCollections: SafeIndexAtLock out-of-range ----------

func Test_I28_LinkedCollections_SafeIndexAtLock_OutOfRange(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_SafeIndexAtLock_OutOfRange", func() {
		// Arrange
		lc := corestr.New.LinkedCollections.Create()
		c1 := corestr.New.Collection.Strings("a", "b")
		lc.Add(c1)

		// Act
		node := lc.SafeIndexAtLock(999)

		// Assert
		actual := args.Map{"isNil": node == nil}
		expected := args.Map{"isNil": true}
		expected.ShouldBeEqual(t, 0, "SafeIndexAtLock returns nil -- out-of-range index", actual)
	})
}

// ---------- LinkedCollections: IsChainEqual both nil heads ----------

func Test_I28_LinkedCollections_IsChainEqual_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_IsChainEqual_BothEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollections.Create()
		lc2 := corestr.New.LinkedCollections.Create()

		// Act
		result := lc1.IsChainEqual(lc2)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": true}
		expected.ShouldBeEqual(t, 0, "IsChainEqual returns true -- both empty", actual)
	})
}

// ---------- LinkedCollections: IsChainEqual one nil ----------

func Test_I28_LinkedCollections_IsChainEqual_OneEmpty(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_IsChainEqual_OneEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollections.Create()
		lc2 := corestr.New.LinkedCollections.Create()
		lc2.Add(corestr.New.Collection.Strings("a"))

		// Act
		result := lc1.IsChainEqual(lc2)

		// Assert
		actual := args.Map{"isEqual": result}
		expected := args.Map{"isEqual": false}
		expected.ShouldBeEqual(t, 0, "IsChainEqual returns false -- one empty", actual)
	})
}

// ---------- LinkedCollections: ToCollection with items ----------

func Test_I28_LinkedCollections_ToCollection(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_ToCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollections.Create()
		lc.Add(corestr.New.Collection.Strings("a", "b"))
		lc.Add(corestr.New.Collection.Strings("c"))

		// Act
		result := lc.ToCollection(0)

		// Assert
		actual := args.Map{"length": result.Length()}
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "ToCollection merges all -- with items", actual)
	})
}

// ---------- LinkedCollections: ToCollectionsOfCollection with items ----------

func Test_I28_LinkedCollections_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_ToCollectionsOfCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollections.Create()
		lc.Add(corestr.New.Collection.Strings("a", "b"))
		lc.Add(corestr.New.Collection.Strings("c"))

		// Act
		result := lc.ToCollectionsOfCollection(0)

		// Assert
		actual := args.Map{"length": result.Length()}
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "ToCollectionsOfCollection returns collections -- with items", actual)
	})
}

// ---------- LinkedCollections: AddCollections all nil ----------

func Test_I28_LinkedCollections_AddCollections_AllNil(t *testing.T) {
	safeTest(t, "Test_I28_LinkedCollections_AddCollections_AllNil", func() {
		// Arrange
		lc := corestr.New.LinkedCollections.Create()

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

// ---------- CollectionsOfCollection: AllLengthsSum with empty collection ----------

func Test_I28_CollectionsOfCollection_AllLengthsSum_EmptyCollectionSkip(t *testing.T) {
	safeTest(t, "Test_I28_CollectionsOfCollection_AllLengthsSum_EmptyCollectionSkip", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(4)
		coc.Add(corestr.New.Collection.Strings("a", "b"))
		coc.Add(corestr.New.Collection.Cap(0)) // empty
		coc.Add(corestr.New.Collection.Strings("c"))

		// Act
		result := coc.AllLengthsSum()

		// Assert
		actual := args.Map{"sum": result}
		expected := args.Map{"sum": 3}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum skips empty -- mixed collections", actual)
	})
}

// ---------- CollectionsOfCollection: ToStrings nil/empty skip ----------

func Test_I28_CollectionsOfCollection_ToStrings_NilEmptySkip(t *testing.T) {
	safeTest(t, "Test_I28_CollectionsOfCollection_ToStrings_NilEmptySkip", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(4)
		coc.Add(corestr.New.Collection.Strings("a"))
		coc.Add(nil) // nil
		coc.Add(corestr.New.Collection.Strings("b"))

		// Act
		result := coc.ToStrings()

		// Assert
		actual := args.Map{"length": len(result)}
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "ToStrings skips nil -- mixed items", actual)
	})
}
