package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// === Collection uncovered ===

func Test_Cov_Collection_LengthLock(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2})
	actual := args.Map{"result": c.LengthLock() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_IsEmptyLock(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	actual := args.Map{"result": c.IsEmptyLock()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov_Collection_AddLock(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	c.AddLock(1)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_Collection_AddsLock(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	c.AddsLock(1, 2)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_AddIfMany(t *testing.T) {
	c := coregeneric.EmptyCollection[int]()
	c.AddIfMany(false, 1, 2)
	actual := args.Map{"result": c.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	c.AddIfMany(true, 1, 2)
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_ForEachBreak(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	c.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_CountFunc(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	n := c.CountFunc(func(v int) bool { return v > 2 })
	actual := args.Map{"result": n != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Collection_SortFunc(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"result": c.First() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_Collection_Reverse(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1, 2, 3})
	c.Reverse()
	actual := args.Map{"result": c.First() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov_Collection_ConcatNew(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1})
	n := c.ConcatNew(2, 3)
	actual := args.Map{"result": n.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov_Collection_String(t *testing.T) {
	c := coregeneric.CollectionFrom([]int{1})
	actual := args.Map{"result": c.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_Cov_Collection_CollectionLenCap(t *testing.T) {
	c := coregeneric.CollectionLenCap[int](3, 10)
	actual := args.Map{"result": c.Length() != 3 || c.Capacity() < 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// === LinkedList uncovered ===

func Test_Cov_LinkedList_LengthLock(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	actual := args.Map{"result": ll.LengthLock() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_LinkedList_IsEmptyLock(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	actual := args.Map{"result": ll.IsEmptyLock()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_Cov_LinkedList_AddLock(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(1)
	actual := args.Map{"result": ll.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_LinkedList_AddsIf(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(false, 1, 2)
	actual := args.Map{"result": ll.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	ll.AddsIf(true, 1, 2)
	actual := args.Map{"result": ll.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_LinkedList_AppendChainOfNodes(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1)
	chain := coregeneric.LinkedListFrom([]int{2, 3})
	ll.AppendChainOfNodes(chain.Head())
	actual := args.Map{"result": ll.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	chain := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendChainOfNodes(chain.Head())
	actual := args.Map{"result": ll.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_LinkedList_ForEachBreak(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_LinkedList_ForEachBreak_FirstItem(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return true
	})
	actual := args.Map{"result": count != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_LinkedList_IndexAt(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(2)
	actual := args.Map{"result": node == nil || node.Element != 30}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 30", actual)
	actual := args.Map{"result": ll.IndexAt(-1) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for negative", actual)
	actual := args.Map{"result": ll.IndexAt(10) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for out of range", actual)
}

func Test_Cov_LinkedList_Collection(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	c := ll.Collection()
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_LinkedList_String(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1})
	actual := args.Map{"result": ll.String() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// === Numeric funcs uncovered ===

func Test_Cov_CompareNumeric(t *testing.T) {
	actual := args.Map{"result": coregeneric.CompareNumeric(1, 2) != corecomparator.LeftLess}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftLess", actual)
	actual := args.Map{"result": coregeneric.CompareNumeric(2, 1) != corecomparator.LeftGreater}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected LeftGreater", actual)
	actual := args.Map{"result": coregeneric.CompareNumeric(1, 1) != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Cov_Clamp(t *testing.T) {
	actual := args.Map{"result": coregeneric.Clamp(5, 1, 10) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "in range", actual)
	actual := args.Map{"result": coregeneric.Clamp(-1, 0, 10) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "below min", actual)
	actual := args.Map{"result": coregeneric.Clamp(20, 0, 10) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "above max", actual)
}

func Test_Cov_Abs(t *testing.T) {
	actual := args.Map{"result": coregeneric.Abs(-5) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_Cov_AbsDiff(t *testing.T) {
	actual := args.Map{"result": coregeneric.AbsDiff(3, 5) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Cov_Sign(t *testing.T) {
	actual := args.Map{"result": coregeneric.Sign(-5) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
	actual := args.Map{"result": coregeneric.Sign(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": coregeneric.Sign(5) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_SafeDiv(t *testing.T) {
	actual := args.Map{"result": coregeneric.SafeDiv(10, 0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for div by zero", actual)
	actual := args.Map{"result": coregeneric.SafeDiv(10, 2) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)
}

func Test_Cov_SafeDivOrDefault(t *testing.T) {
	actual := args.Map{"result": coregeneric.SafeDivOrDefault(10, 0, -1) != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_Cov_MinOfSlice(t *testing.T) {
	actual := args.Map{"result": coregeneric.MinOfSlice([]int{3, 1, 2}) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_Cov_MaxOfSlice(t *testing.T) {
	actual := args.Map{"result": coregeneric.MaxOfSlice([]int{3, 1, 2}) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Cov_InRangeExclusive(t *testing.T) {
	actual := args.Map{"result": coregeneric.InRangeExclusive(5, 0, 10)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": coregeneric.InRangeExclusive(0, 0, 10)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for boundary", actual)
}

func Test_Cov_IsNegative(t *testing.T) {
	actual := args.Map{"result": coregeneric.IsNegative(-1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Cov_IsNonNegative(t *testing.T) {
	actual := args.Map{"result": coregeneric.IsNonNegative(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}
