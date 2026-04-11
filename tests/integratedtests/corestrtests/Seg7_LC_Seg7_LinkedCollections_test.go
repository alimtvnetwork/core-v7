package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollections — Segment 7f
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_IsEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{
			"empty": lc.IsEmpty(),
			"hasItems": lc.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg7_LC_Add(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Add", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		lc.Add(c1).Add(c2)

		// Act
		actual := args.Map{
			"len": lc.Length(),
			"allLen": lc.AllIndividualItemsLength(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"allLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 collections 3 items", actual)
	})
}

func Test_Seg7_LC_AddLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AddLock(c)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddLock -- 1", actual)
	})
}

func Test_Seg7_LC_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings("a", "b")

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 1 collection", actual)
	})
}

func Test_Seg7_LC_AddStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddStrings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStrings()

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_Seg7_LC_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddStringsLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsLock("a")

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStringsLock -- 1", actual)
	})
}

func Test_Seg7_LC_AddFront(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddFront", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"b"})
		c2 := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c1).AddFront(c2)

		// Act
		actual := args.Map{
			"head": lc.Head().Element.List()[0],
			"len": lc.Length(),
		}

		// Assert
		expected := args.Map{
			"head": "a",
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "AddFront -- prepended", actual)
	})
}

func Test_Seg7_LC_AddFront_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddFront_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AddFront(c)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddFront empty -- added", actual)
	})
}

func Test_Seg7_LC_AddFrontLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddFrontLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AddFrontLock(c)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddFrontLock -- 1", actual)
	})
}

func Test_Seg7_LC_Push_PushBack_PushFront(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Push_PushBack_PushFront", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c3 := corestr.New.Collection.Strings([]string{"c"})
		lc.Push(c1).PushBack(c2).PushFront(c3)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Push/PushBack/PushFront -- 3", actual)
	})
}

func Test_Seg7_LC_PushBackLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_PushBackLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.PushBackLock(c)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "PushBackLock -- 1", actual)
	})
}

func Test_Seg7_LC_AddCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollection(c)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollection -- 1", actual)
	})
}

func Test_Seg7_LC_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddCollection_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollection(nil)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil -- no change", actual)
	})
}

func Test_Seg7_LC_AddAnother(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddAnother", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc1.Add(c1)
		lc2 := corestr.New.LinkedCollection.Create()
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc2.Add(c2)
		lc1.AddAnother(lc2)

		// Act
		actual := args.Map{"len": lc1.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddAnother -- merged", actual)
	})
}

func Test_Seg7_LC_AddAnother_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddAnother_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddAnother(nil)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddAnother nil -- no change", actual)
	})
}

func Test_Seg7_LC_AddStringsOfStrings(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddStringsOfStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsOfStrings -- 2 collections", actual)
	})
}

func Test_Seg7_LC_AddStringsOfStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddStringsOfStrings_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddStringsOfStrings(false)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsOfStrings empty -- no change", actual)
	})
}

// ── Accessors ───────────────────────────────────────────────────────────────

func Test_Seg7_LC_FirstLastSingle(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_FirstLastSingle", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{
			"first":  lc.First().List()[0],
			"last":   lc.Last().List()[0],
			"single": lc.Single().List()[0],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "a",
			"single": "a",
		}
		expected.ShouldBeEqual(t, 0, "First/Last/Single -- correct", actual)
	})
}

func Test_Seg7_LC_FirstOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_FirstOrDefault_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"empty": lc.FirstOrDefault().IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- empty collection", actual)
	})
}

func Test_Seg7_LC_LastOrDefault_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_LastOrDefault_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"empty": lc.LastOrDefault().IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- empty collection", actual)
	})
}

func Test_Seg7_LC_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_LengthLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": lc.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg7_LC_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_IsEmptyLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"empty": lc.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_Seg7_LC_SafeIndexAt(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_SafeIndexAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{
			"at0":  lc.SafeIndexAt(0) != nil,
			"neg":  lc.SafeIndexAt(-1) == nil,
			"over": lc.SafeIndexAt(5) == nil,
		}

		// Assert
		expected := args.Map{
			"at0": true,
			"neg": true,
			"over": true,
		}
		expected.ShouldBeEqual(t, 0, "SafeIndexAt -- various", actual)
	})
}

func Test_Seg7_LC_SafePointerIndexAt(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_SafePointerIndexAt", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{
			"at0": lc.SafePointerIndexAt(0) != nil,
			"nil": lc.SafePointerIndexAt(5) == nil,
		}

		// Assert
		expected := args.Map{
			"at0": true,
			"nil": true,
		}
		expected.ShouldBeEqual(t, 0, "SafePointerIndexAt -- correct", actual)
	})
}

func Test_Seg7_LC_GetNextNodes(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_GetNextNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		nodes := lc.GetNextNodes(1)

		// Act
		actual := args.Map{"len": len(nodes)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetNextNodes -- 1 node", actual)
	})
}

func Test_Seg7_LC_GetAllLinkedNodes(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_GetAllLinkedNodes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		nodes := lc.GetAllLinkedNodes()

		// Act
		actual := args.Map{"len": len(nodes)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllLinkedNodes -- 1", actual)
	})
}

// ── IsEquals ────────────────────────────────────────────────────────────────

func Test_Seg7_LC_IsEqualsPtr(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_IsEqualsPtr", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc1.Add(c1)
		lc2 := corestr.New.LinkedCollection.Create()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		lc2.Add(c2)

		// Act
		actual := args.Map{
			"eq":   lc1.IsEqualsPtr(lc2),
			"self": lc1.IsEqualsPtr(lc1),
			"nil":  lc1.IsEqualsPtr(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"self": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr -- various", actual)
	})
}

func Test_Seg7_LC_IsEqualsPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_IsEqualsPtr_BothEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc2 := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"eq": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr both empty -- true", actual)
	})
}

func Test_Seg7_LC_IsEqualsPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_IsEqualsPtr_OneEmpty", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc1.Add(c)
		lc2 := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"eq": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr one empty -- false", actual)
	})
}

func Test_Seg7_LC_IsEqualsPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_IsEqualsPtr_DiffLen", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		lc1.Add(c1)
		lc2 := corestr.New.LinkedCollection.Create()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c3 := corestr.New.Collection.Strings([]string{"b"})
		lc2.Add(c2).Add(c3)

		// Act
		actual := args.Map{"eq": lc1.IsEqualsPtr(lc2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualsPtr diff len -- false", actual)
	})
}

// ── ToCollection / ToStrings / ToCollectionsOfCollection ────────────────────

func Test_Seg7_LC_ToCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ToCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c)
		result := lc.ToCollection(0)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollection -- 2 items", actual)
	})
}

func Test_Seg7_LC_ToCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ToCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		result := lc.ToCollection(0)

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "ToCollection empty -- empty", actual)
	})
}

func Test_Seg7_LC_ToCollectionSimple(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ToCollectionSimple", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": lc.ToCollectionSimple().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ToCollectionSimple -- 1", actual)
	})
}

func Test_Seg7_LC_ToStrings(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ToStrings", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": len(lc.ToStrings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ToStrings -- 1", actual)
	})
}

func Test_Seg7_LC_ToCollectionsOfCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ToCollectionsOfCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		result := lc.ToCollectionsOfCollection(0)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ToCollectionsOfCollection -- 1", actual)
	})
}

func Test_Seg7_LC_ToCollectionsOfCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ToCollectionsOfCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		result := lc.ToCollectionsOfCollection(0)

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "ToCollectionsOfCollection empty -- empty", actual)
	})
}

func Test_Seg7_LC_ItemsOfItems(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ItemsOfItems", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": len(lc.ItemsOfItems())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ItemsOfItems -- 1", actual)
	})
}

func Test_Seg7_LC_ItemsOfItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ItemsOfItems_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"len": len(lc.ItemsOfItems())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ItemsOfItems empty -- 0", actual)
	})
}

func Test_Seg7_LC_ItemsOfItemsCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ItemsOfItemsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": len(lc.ItemsOfItemsCollection())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ItemsOfItemsCollection -- 1", actual)
	})
}

func Test_Seg7_LC_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_SimpleSlice", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": lc.SimpleSlice().Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice -- 1", actual)
	})
}

func Test_Seg7_LC_List(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_List", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": len(lc.List())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "List -- 1", actual)
	})
}

func Test_Seg7_LC_List_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_List_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"len": len(lc.List())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty -- 0", actual)
	})
}

// ── Loop / Filter ───────────────────────────────────────────────────────────

func Test_Seg7_LC_Loop(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Loop", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 2}
		expected.ShouldBeEqual(t, 0, "Loop -- visits all", actual)
	})
}

func Test_Seg7_LC_Loop_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Loop_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 0}
		expected.ShouldBeEqual(t, 0, "Loop empty -- no visits", actual)
	})
}

func Test_Seg7_LC_Loop_Break(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Loop_Break", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		count := 0
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return true
		})

		// Act
		actual := args.Map{"count": count}

		// Assert
		expected := args.Map{"count": 1}
		expected.ShouldBeEqual(t, 0, "Loop break -- 1", actual)
	})
}

func Test_Seg7_LC_Filter(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Filter", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		result := lc.Filter(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter -- 2 nodes", actual)
	})
}

func Test_Seg7_LC_FilterAsCollection(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_FilterAsCollection", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c)
		result := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		}, 0)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterAsCollection -- 2 items", actual)
	})
}

func Test_Seg7_LC_FilterAsCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_FilterAsCollection_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		result := lc.FilterAsCollection(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: false, IsBreak: false}
		}, 0)

		// Act
		actual := args.Map{"empty": result.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "FilterAsCollection none kept -- empty", actual)
	})
}

func Test_Seg7_LC_FilterAsCollections(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_FilterAsCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		result := lc.FilterAsCollections(func(arg *corestr.LinkedCollectionFilterParameter) *corestr.LinkedCollectionFilterResult {
			return &corestr.LinkedCollectionFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterAsCollections -- 1", actual)
	})
}

// ── String / Join / JSON ────────────────────────────────────────────────────

func Test_Seg7_LC_String(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_String", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"nonEmpty": lc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg7_LC_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_String_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"nonEmpty": lc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- non-empty (NoElements)", actual)
	})
}

func Test_Seg7_LC_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_StringLock", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"nonEmpty": lc.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg7_LC_Join(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Join", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c)

		// Act
		actual := args.Map{"val": lc.Join(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join -- comma separated", actual)
	})
}

func Test_Seg7_LC_Joins(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Joins", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"nonEmpty": lc.Joins(",", "b") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins -- combined", actual)
	})
}

func Test_Seg7_LC_Json(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Json", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		j := lc.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg7_LC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_MarshalJSON", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		b, err := lc.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg7_LC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_UnmarshalJSON", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		b, _ := lc.MarshalJSON()
		lc2 := corestr.New.LinkedCollection.Create()
		err := lc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg7_LC_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_UnmarshalJSON_Invalid", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		err := lc.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg7_LC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_JsonModel", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)

		// Act
		actual := args.Map{"len": len(lc.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1", actual)
	})
}

func Test_Seg7_LC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_JsonModelAny", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{"notNil": lc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg7_LC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ParseInjectUsingJson", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		jr := lc.JsonPtr()
		lc2 := corestr.New.LinkedCollection.Create()
		_, err := lc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg7_LC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ParseInjectUsingJsonMust", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		jr := lc.JsonPtr()
		lc2 := corestr.New.LinkedCollection.Create()
		result := lc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg7_LC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_JsonParseSelfInject", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		jr := lc.JsonPtr()
		lc2 := corestr.New.LinkedCollection.Create()
		err := lc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg7_LC_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_InterfaceCasts", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()

		// Act
		actual := args.Map{
			"jsoner":   lc.AsJsoner() != nil,
			"binder":   lc.AsJsonContractsBinder() != nil,
			"injector": lc.AsJsonParseSelfInjector() != nil,
			"marsh":    lc.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
			"marsh": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

// ── Remove / Clear ──────────────────────────────────────────────────────────

func Test_Seg7_LC_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNodeByIndex", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c3 := corestr.New.Collection.Strings([]string{"c"})
		lc.Add(c1).Add(c2).Add(c3)
		lc.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex -- removed", actual)
	})
}

func Test_Seg7_LC_RemoveNodeByIndex_First(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNodeByIndex_First", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		lc.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex first -- removed", actual)
	})
}

func Test_Seg7_LC_RemoveNodeByIndex_Last(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNodeByIndex_Last", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		lc.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndex last -- removed", actual)
	})
}

func Test_Seg7_LC_RemoveNodeByIndexes(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNodeByIndexes", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c3 := corestr.New.Collection.Strings([]string{"c"})
		lc.Add(c1).Add(c2).Add(c3)
		lc.RemoveNodeByIndexes(false, 0, 2)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes -- removed 2", actual)
	})
}

func Test_Seg7_LC_RemoveNodeByIndexes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		lc.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNodeByIndexes empty -- no change", actual)
	})
}

func Test_Seg7_LC_RemoveNode(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNode", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		node := lc.SafeIndexAt(1)
		lc.RemoveNode(node)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNode -- removed", actual)
	})
}

func Test_Seg7_LC_RemoveNode_First(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveNode_First", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		lc.RemoveNode(lc.Head())

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveNode first -- removed", actual)
	})
}

func Test_Seg7_LC_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ConcatNew", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc1.Add(c)
		lc2 := corestr.New.LinkedCollection.Create()
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc2.Add(c2)
		result := lc1.ConcatNew(false, lc2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- combined", actual)
	})
}

func Test_Seg7_LC_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ConcatNew_EmptyClone", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		result := lc.ConcatNew(true)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty clone -- cloned", actual)
	})
}

func Test_Seg7_LC_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_ConcatNew_EmptyNoClone", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		result := lc.ConcatNew(false)

		// Act
		actual := args.Map{"same": result == lc}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty no clone -- returns self", actual)
	})
}

func Test_Seg7_LC_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AppendCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.AppendCollections(false, c1, c2)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendCollections -- 2", actual)
	})
}

func Test_Seg7_LC_AppendCollections_SkipNil(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AppendCollections_SkipNil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AppendCollections(true, c, nil)

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendCollections skip nil -- 1", actual)
	})
}

func Test_Seg7_LC_AddCollections(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddCollections", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollections([]*corestr.Collection{c})

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollections -- 1", actual)
	})
}

func Test_Seg7_LC_AddCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddCollections_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.AddCollections([]*corestr.Collection{})

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollections empty -- no change", actual)
	})
}

func Test_Seg7_LC_AddCollectionsPtr(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_AddCollectionsPtr", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.AddCollectionsPtr([]*corestr.Collection{c})

		// Act
		actual := args.Map{"len": lc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollectionsPtr -- 1", actual)
	})
}

func Test_Seg7_LC_Clear(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Clear", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		lc.Clear()

		// Act
		actual := args.Map{"empty": lc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg7_LC_Clear_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_Clear_Empty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		result := lc.Clear()

		// Act
		actual := args.Map{"same": result == lc}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Clear empty -- returns self", actual)
	})
}

func Test_Seg7_LC_RemoveAll(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_RemoveAll", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		lc.RemoveAll()

		// Act
		actual := args.Map{"empty": lc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "RemoveAll -- emptied", actual)
	})
}

func Test_Seg7_LC_GetCompareSummary(t *testing.T) {
	safeTest(t, "Test_Seg7_LC_GetCompareSummary", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc1.Add(c)
		lc2 := corestr.New.LinkedCollection.Create()
		c2 := corestr.New.Collection.Strings([]string{"a"})
		lc2.Add(c2)

		// Act
		actual := args.Map{"nonEmpty": lc1.GetCompareSummary(lc2, "left", "right") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "GetCompareSummary -- non-empty", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedListNode — Segment 7g
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LLN_Basic(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_Basic", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "hello"}

		// Act
		actual := args.Map{
			"elem":    node.Element,
			"str":     node.String(),
			"hasNext": node.HasNext(),
			"next":    node.Next() == nil,
		}

		// Assert
		expected := args.Map{
			"elem": "hello",
			"str": "hello",
			"hasNext": false,
			"next": true,
		}
		expected.ShouldBeEqual(t, 0, "Basic -- correct", actual)
	})
}

func Test_Seg7_LLN_Clone(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_Clone", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "hello"}
		c := node.Clone()

		// Act
		actual := args.Map{
			"elem": c.Element,
			"diff": c != node,
			"hasNext": c.HasNext(),
		}

		// Assert
		expected := args.Map{
			"elem": "hello",
			"diff": true,
			"hasNext": false,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- new copy no next", actual)
	})
}

func Test_Seg7_LLN_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_EndOfChain", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		end, length := ll.Head().EndOfChain()

		// Act
		actual := args.Map{
			"end": end.Element,
			"len": length,
		}

		// Assert
		expected := args.Map{
			"end": "c",
			"len": 3,
		}
		expected.ShouldBeEqual(t, 0, "EndOfChain -- end node + length", actual)
	})
}

func Test_Seg7_LLN_List(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		list := ll.Head().List()

		// Act
		actual := args.Map{
			"len": len(list),
			"first": list[0],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "List -- from node onwards", actual)
	})
}

func Test_Seg7_LLN_Join(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"val": ll.Head().Join(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join -- comma separated", actual)
	})
}

func Test_Seg7_LLN_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_IsEqual", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("a", "b")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("a", "b")

		// Act
		actual := args.Map{
			"eq":      ll1.Head().IsEqual(ll2.Head()),
			"self":    ll1.Head().IsEqual(ll1.Head()),
			"nilBoth": (*corestr.LinkedListNode)(nil).IsEqual(nil),
			"nilOne":  ll1.Head().IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"self": true,
			"nilBoth": true,
			"nilOne": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqual -- various", actual)
	})
}

func Test_Seg7_LLN_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_IsEqualValue", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "a"}

		// Act
		actual := args.Map{
			"eq": node.IsEqualValue("a"),
			"neq": node.IsEqualValue("b"),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualValue -- correct", actual)
	})
}

func Test_Seg7_LLN_IsEqualValueSensitive(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_IsEqualValueSensitive", func() {
		// Arrange
		node := &corestr.LinkedListNode{Element: "ABC"}

		// Act
		actual := args.Map{
			"sensitive":   node.IsEqualValueSensitive("abc", true),
			"insensitive": node.IsEqualValueSensitive("abc", false),
		}

		// Assert
		expected := args.Map{
			"sensitive": false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualValueSensitive -- case matters", actual)
	})
}

func Test_Seg7_LLN_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_IsChainEqual", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("a", "b", "c")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("a", "b", "c")

		// Act
		actual := args.Map{
			"eq":      ll1.Head().IsChainEqual(ll2.Head(), true),
			"nilBoth": (*corestr.LinkedListNode)(nil).IsChainEqual(nil, true),
			"nilOne":  ll1.Head().IsChainEqual(nil, true),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"nilBoth": true,
			"nilOne": false,
		}
		expected.ShouldBeEqual(t, 0, "IsChainEqual -- correct", actual)
	})
}

func Test_Seg7_LLN_IsEqualSensitive(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_IsEqualSensitive", func() {
		// Arrange
		ll1 := corestr.New.LinkedList.Create()
		ll1.Adds("ABC")
		ll2 := corestr.New.LinkedList.Create()
		ll2.Adds("abc")

		// Act
		actual := args.Map{
			"sensitive":   ll1.Head().IsEqualSensitive(ll2.Head(), true),
			"insensitive": ll1.Head().IsEqualSensitive(ll2.Head(), false),
			"nilBoth":     (*corestr.LinkedListNode)(nil).IsEqualSensitive(nil, true),
			"nilOne":      ll1.Head().IsEqualSensitive(nil, true),
			"self":        ll1.Head().IsEqualSensitive(ll1.Head(), true),
		}

		// Assert
		expected := args.Map{
			"sensitive": false,
			"insensitive": true,
			"nilBoth": true,
			"nilOne": false,
			"self": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualSensitive -- various", actual)
	})
}

func Test_Seg7_LLN_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_CreateLinkedList", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		newLL := ll.Head().CreateLinkedList()

		// Act
		actual := args.Map{"len": newLL.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CreateLinkedList -- from chain", actual)
	})
}

func Test_Seg7_LLN_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_LoopEndOfChain", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		count := 0
		end, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{
			"end": end.Element,
			"len": length,
			"count": count,
		}

		// Assert
		expected := args.Map{
			"end": "b",
			"len": 2,
			"count": 2,
		}
		expected.ShouldBeEqual(t, 0, "LoopEndOfChain -- all visited", actual)
	})
}

func Test_Seg7_LLN_LoopEndOfChain_Break(t *testing.T) {
	safeTest(t, "Test_Seg7_LLN_LoopEndOfChain_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		_, length := ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return true
		})

		// Act
		actual := args.Map{"len": length}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LoopEndOfChain break -- 1", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// LinkedCollectionNode — Segment 7h
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_LCN_Basic(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_Basic", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		node := &corestr.LinkedCollectionNode{Element: c}

		// Act
		actual := args.Map{
			"hasElem": node.HasElement(),
			"hasNext": node.HasNext(),
			"isEmpty": node.IsEmpty(),
			"str":     node.String() != "",
		}

		// Assert
		expected := args.Map{
			"hasElem": true,
			"hasNext": false,
			"isEmpty": false,
			"str": true,
		}
		expected.ShouldBeEqual(t, 0, "Basic -- correct", actual)
	})
}

func Test_Seg7_LCN_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_IsEmpty_Nil", func() {
		// Arrange
		var node *corestr.LinkedCollectionNode

		// Act
		actual := args.Map{"empty": node.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmpty nil -- true", actual)
	})
}

func Test_Seg7_LCN_Clone(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_Clone", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		node := &corestr.LinkedCollectionNode{Element: c}
		cl := node.Clone()

		// Act
		actual := args.Map{
			"diff": cl != node,
			"hasNext": cl.HasNext(),
		}

		// Assert
		expected := args.Map{
			"diff": true,
			"hasNext": false,
		}
		expected.ShouldBeEqual(t, 0, "Clone -- new copy", actual)
	})
}

func Test_Seg7_LCN_EndOfChain(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_EndOfChain", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		end, length := lc.Head().EndOfChain()

		// Act
		actual := args.Map{
			"len": length,
			"notNil": end != nil,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "EndOfChain -- end + length", actual)
	})
}

func Test_Seg7_LCN_List(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_List", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c)
		list := lc.Head().List()

		// Act
		actual := args.Map{"len": len(list)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List -- 2 items", actual)
	})
}

func Test_Seg7_LCN_Join(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_Join", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		lc.Add(c)

		// Act
		actual := args.Map{"val": lc.Head().Join(",")}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join -- comma separated", actual)
	})
}

func Test_Seg7_LCN_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_IsEqual", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"a"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		n2 := &corestr.LinkedCollectionNode{Element: c2}

		// Act
		actual := args.Map{
			"eq":      n1.IsEqual(n2),
			"self":    n1.IsEqual(n1),
			"nilBoth": (*corestr.LinkedCollectionNode)(nil).IsEqual(nil),
			"nilOne":  n1.IsEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"self": true,
			"nilBoth": true,
			"nilOne": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqual -- various", actual)
	})
}

func Test_Seg7_LCN_IsEqualValue(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_IsEqualValue", func() {
		// Arrange
		c1 := corestr.New.Collection.Strings([]string{"a"})
		node := &corestr.LinkedCollectionNode{Element: c1}
		c2 := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"eq":  node.IsEqualValue(c2),
			"nil": node.IsEqualValue(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualValue -- correct", actual)
	})
}

func Test_Seg7_LCN_IsChainEqual(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_IsChainEqual", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc1.Add(c1).Add(c2)
		lc2 := corestr.New.LinkedCollection.Create()
		c3 := corestr.New.Collection.Strings([]string{"a"})
		c4 := corestr.New.Collection.Strings([]string{"b"})
		lc2.Add(c3).Add(c4)

		// Act
		actual := args.Map{
			"eq":      lc1.Head().IsChainEqual(lc2.Head()),
			"nilBoth": (*corestr.LinkedCollectionNode)(nil).IsChainEqual(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"nilBoth": true,
		}
		expected.ShouldBeEqual(t, 0, "IsChainEqual -- correct", actual)
	})
}

func Test_Seg7_LCN_CreateLinkedList(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_CreateLinkedList", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c := corestr.New.Collection.Strings([]string{"a"})
		lc.Add(c)
		newLC := lc.Head().CreateLinkedList()

		// Act
		actual := args.Map{"len": newLC.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CreateLinkedList -- from chain", actual)
	})
}

func Test_Seg7_LCN_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_Seg7_LCN_LoopEndOfChain", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		lc.Add(c1).Add(c2)
		count := 0
		_, length := lc.Head().LoopEndOfChain(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{
			"len": length,
			"count": count,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"count": 2,
		}
		expected.ShouldBeEqual(t, 0, "LoopEndOfChain -- all visited", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NonChainedLinkedListNodes — Segment 7i
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_NCLLN_Basic(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_Basic", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)

		// Act
		actual := args.Map{
			"empty": ncl.IsEmpty(),
			"len": ncl.Length(),
			"hasItems": ncl.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "Basic -- empty", actual)
	})
}

func Test_Seg7_NCLLN_Adds(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_Adds", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		ncl.Adds(n1, n2)

		// Act
		actual := args.Map{
			"len":   ncl.Length(),
			"first": ncl.First().Element,
			"last":  ncl.Last().Element,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
			"last": "b",
		}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 nodes", actual)
	})
}

func Test_Seg7_NCLLN_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_Adds_Nil", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		ncl.Adds(nil...)

		// Act
		actual := args.Map{"len": ncl.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_Seg7_NCLLN_FirstLastOrDefault(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_FirstLastOrDefault", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)

		// Act
		actual := args.Map{
			"firstNil": ncl.FirstOrDefault() == nil,
			"lastNil": ncl.LastOrDefault() == nil,
		}

		// Assert
		expected := args.Map{
			"firstNil": true,
			"lastNil": true,
		}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault/LastOrDefault empty -- nil", actual)
	})
}

func Test_Seg7_NCLLN_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_ApplyChaining", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		ncl.Adds(n1, n2)
		ncl.ApplyChaining()

		// Act
		actual := args.Map{
			"chained": ncl.IsChainingApplied(),
			"hasNext": n1.HasNext(),
		}

		// Assert
		expected := args.Map{
			"chained": true,
			"hasNext": true,
		}
		expected.ShouldBeEqual(t, 0, "ApplyChaining -- linked", actual)
	})
}

func Test_Seg7_NCLLN_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_ApplyChaining_Empty", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		ncl.ApplyChaining()

		// Act
		actual := args.Map{"len": ncl.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ApplyChaining empty -- no change", actual)
	})
}

func Test_Seg7_NCLLN_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_ToChainedNodes", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		n1 := &corestr.LinkedListNode{Element: "a"}
		n2 := &corestr.LinkedListNode{Element: "b"}
		ncl.Adds(n1, n2)
		result := ncl.ToChainedNodes()

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes -- not nil", actual)
	})
}

func Test_Seg7_NCLLN_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_ToChainedNodes_Empty", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		result := ncl.ToChainedNodes()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes empty -- 0", actual)
	})
}

func Test_Seg7_NCLLN_Items(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLLN_Items", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedListNodes(4)
		n := &corestr.LinkedListNode{Element: "a"}
		ncl.Adds(n)

		// Act
		actual := args.Map{"len": len(ncl.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// NonChainedLinkedCollectionNodes — Segment 7j
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg7_NCLN_Basic(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_Basic", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)

		// Act
		actual := args.Map{
			"empty": ncl.IsEmpty(),
			"len": ncl.Length(),
			"hasItems": ncl.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"len": 0,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "Basic -- empty", actual)
	})
}

func Test_Seg7_NCLN_Adds(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_Adds", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		n2 := &corestr.LinkedCollectionNode{Element: c2}
		ncl.Adds(n1, n2)

		// Act
		actual := args.Map{
			"len": ncl.Length(),
			"first": ncl.First() != nil,
			"last": ncl.Last() != nil,
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": true,
			"last": true,
		}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 nodes", actual)
	})
}

func Test_Seg7_NCLN_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_Adds_Nil", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		ncl.Adds(nil...)

		// Act
		actual := args.Map{"len": ncl.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_Seg7_NCLN_FirstLastOrDefault(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_FirstLastOrDefault", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)

		// Act
		actual := args.Map{
			"firstNil": ncl.FirstOrDefault() == nil,
			"lastNil": ncl.LastOrDefault() == nil,
		}

		// Assert
		expected := args.Map{
			"firstNil": true,
			"lastNil": true,
		}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault/LastOrDefault empty -- nil", actual)
	})
}

func Test_Seg7_NCLN_ApplyChaining(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_ApplyChaining", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		n2 := &corestr.LinkedCollectionNode{Element: c2}
		ncl.Adds(n1, n2)
		ncl.ApplyChaining()

		// Act
		actual := args.Map{
			"chained": ncl.IsChainingApplied(),
			"hasNext": n1.HasNext(),
		}

		// Assert
		expected := args.Map{
			"chained": true,
			"hasNext": true,
		}
		expected.ShouldBeEqual(t, 0, "ApplyChaining -- linked", actual)
	})
}

func Test_Seg7_NCLN_ApplyChaining_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_ApplyChaining_Empty", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		ncl.ApplyChaining()

		// Act
		actual := args.Map{"len": ncl.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ApplyChaining empty -- no change", actual)
	})
}

func Test_Seg7_NCLN_ApplyChaining_AlreadyApplied(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_ApplyChaining_AlreadyApplied", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		c := corestr.New.Collection.Strings([]string{"a"})
		n := &corestr.LinkedCollectionNode{Element: c}
		ncl.Adds(n)
		ncl.ApplyChaining()
		ncl.ApplyChaining() // should be no-op

		// Act
		actual := args.Map{"chained": ncl.IsChainingApplied()}

		// Assert
		expected := args.Map{"chained": true}
		expected.ShouldBeEqual(t, 0, "ApplyChaining already applied -- no-op", actual)
	})
}

func Test_Seg7_NCLN_ToChainedNodes(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_ToChainedNodes", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		n1 := &corestr.LinkedCollectionNode{Element: c1}
		n2 := &corestr.LinkedCollectionNode{Element: c2}
		ncl.Adds(n1, n2)
		result := ncl.ToChainedNodes()

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes -- not nil", actual)
	})
}

func Test_Seg7_NCLN_ToChainedNodes_Empty(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_ToChainedNodes_Empty", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		result := ncl.ToChainedNodes()

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ToChainedNodes empty -- 0", actual)
	})
}

func Test_Seg7_NCLN_Items(t *testing.T) {
	safeTest(t, "Test_Seg7_NCLN_Items", func() {
		// Arrange
		ncl := corestr.NewNonChainedLinkedCollectionNodes(4)
		c := corestr.New.Collection.Strings([]string{"a"})
		n := &corestr.LinkedCollectionNode{Element: c}
		ncl.Adds(n)

		// Act
		actual := args.Map{"len": len(ncl.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1", actual)
	})
}

// suppress unused import warning
var _ = fmt.Sprintf

