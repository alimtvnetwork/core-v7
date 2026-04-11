package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// === Collection uncovered ===

func Test_Collection_LengthLock(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": c.LengthLock() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Collection_IsEmptyLock(t *testing.T) {
	// Arrange
	c := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": c.IsEmptyLock()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Collection_AddLock(t *testing.T) {
	// Arrange
	c := coregeneric.EmptyCollection[int]()
	c.AddLock(1)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_AddsLock(t *testing.T) {
	// Arrange
	c := coregeneric.EmptyCollection[int]()
	c.AddsLock(1, 2)

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Collection_AddIfMany(t *testing.T) {
	// Arrange
	c := coregeneric.EmptyCollection[int]()
	c.AddIfMany(false, 1, 2)

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	c.AddIfMany(true, 1, 2)
	actual = args.Map{"result": c.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Collection_ForEachBreak(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	c.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Collection_CountFunc(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	n := c.CountFunc(func(v int) bool { return v > 2 })

	// Act
	actual := args.Map{"result": n != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Collection_SortFunc(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"result": c.First() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Collection_Reverse(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{1, 2, 3})
	c.Reverse()

	// Act
	actual := args.Map{"result": c.First() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Collection_ConcatNew(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{1})
	n := c.ConcatNew(2, 3)

	// Act
	actual := args.Map{"result": n.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Collection_String(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"result": c.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Collection_CollectionLenCap(t *testing.T) {
	// Arrange
	c := coregeneric.CollectionLenCap[int](3, 10)

	// Act
	actual := args.Map{"result": c.Length() != 3 || c.Capacity() < 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// === LinkedList uncovered ===

func Test_LinkedList_LengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": ll.LengthLock() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_LinkedList_IsEmptyLock(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{"result": ll.IsEmptyLock()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_LinkedList_AddLock(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(1)

	// Act
	actual := args.Map{"result": ll.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_LinkedList_AddsIf(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(false, 1, 2)

	// Act
	actual := args.Map{"result": ll.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	ll.AddsIf(true, 1, 2)
	actual = args.Map{"result": ll.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_LinkedList_AppendChainOfNodes(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1)
	chain := coregeneric.LinkedListFrom([]int{2, 3})
	ll.AppendChainOfNodes(chain.Head())

	// Act
	actual := args.Map{"result": ll.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	chain := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendChainOfNodes(chain.Head())

	// Act
	actual := args.Map{"result": ll.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_LinkedList_ForEachBreak(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_LinkedList_ForEachBreak_FirstItem(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return true
	})

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_LinkedList_IndexAt(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(2)

	// Act
	actual := args.Map{"result": node == nil || node.Element != 30}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 30", actual)
	actual = args.Map{"result": ll.IndexAt(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for negative", actual)
	actual = args.Map{"result": ll.IndexAt(10) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out of range", actual)
}

func Test_LinkedList_Collection(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	c := ll.Collection()

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_LinkedList_String(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1})

	// Act
	actual := args.Map{"result": ll.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// === Numeric funcs uncovered ===

func Test_CompareNumeric(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.CompareNumeric(1, 2) != corecomparator.LeftLess}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
	actual = args.Map{"result": coregeneric.CompareNumeric(2, 1) != corecomparator.LeftGreater}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
	actual = args.Map{"result": coregeneric.CompareNumeric(1, 1) != corecomparator.Equal}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Clamp(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Clamp(5, 1, 10) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "in range", actual)
	actual = args.Map{"result": coregeneric.Clamp(-1, 0, 10) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "below min", actual)
	actual = args.Map{"result": coregeneric.Clamp(20, 0, 10) != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "above max", actual)
}

func Test_Abs(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Abs(-5) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_AbsDiff(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.AbsDiff(3, 5) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Sign(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Sign(-5) != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual = args.Map{"result": coregeneric.Sign(0) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual = args.Map{"result": coregeneric.Sign(5) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_SafeDiv(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SafeDiv(10, 0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for div by zero", actual)
	actual = args.Map{"result": coregeneric.SafeDiv(10, 2) != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_SafeDivOrDefault(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SafeDivOrDefault(10, 0, -1) != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_MinOfSlice(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.MinOfSlice([]int{3, 1, 2}) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_MaxOfSlice(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.MaxOfSlice([]int{3, 1, 2}) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_InRangeExclusive(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.InRangeExclusive(5, 0, 10)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": coregeneric.InRangeExclusive(0, 0, 10)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for boundary", actual)
}

func Test_IsNegative(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsNegative(-1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_IsNonNegative(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsNonNegative(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}
