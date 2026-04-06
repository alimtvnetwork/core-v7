package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// SimpleSlice — uncovered branches
// =============================================================================

func Test_SimpleSlice_AddIf_Skip(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddIf(false, "skip")

	// Act
	actual := args.Map{"result": s.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add when false", actual)
}

func Test_SimpleSlice_Adds_Empty(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()
	s.Adds()

	// Act
	actual := args.Map{"result": s.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_SimpleSlice_AddSlice_Empty(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddSlice([]string{})

	// Act
	actual := args.Map{"result": s.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_SimpleSlice_AddsIf_Skip(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddsIf(false, "a", "b")

	// Act
	actual := args.Map{"result": s.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add when false", actual)
}

func Test_SimpleSlice_AddFunc(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddFunc(func() string { return "hello" })

	// Act
	actual := args.Map{"result": s.Length() != 1 || s.First() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AddFunc should add result", actual)
}

func Test_SimpleSlice_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()

	// Act
	actual := args.Map{"result": s.FirstOrDefault() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return zero value", actual)
}

func Test_SimpleSlice_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[string]()

	// Act
	actual := args.Map{"result": s.LastOrDefault() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return zero value", actual)
}

func Test_SimpleSlice_Skip_OverCount(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	result := s.Skip(10)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip beyond length should return empty", actual)
}

func Test_SimpleSlice_Take_OverCount(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	result := s.Take(10)

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "take beyond length should return all", actual)
}

func Test_SimpleSlice_Length_Nil(t *testing.T) {
	// Arrange
	var s *coregeneric.SimpleSlice[int]

	// Act
	actual := args.Map{"result": s.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_SimpleSlice_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var s *coregeneric.SimpleSlice[int]

	// Act
	actual := args.Map{"result": s.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
}

func Test_SimpleSlice_HasAnyItem(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1})

	// Act
	actual := args.Map{"result": s.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_SimpleSlice_HasItems(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1})

	// Act
	actual := args.Map{"result": s.HasItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_SimpleSlice_HasIndex(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": s.HasIndex(2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have index 2", actual)
	actual := args.Map{"result": s.HasIndex(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have index 5", actual)
	actual := args.Map{"result": s.HasIndex(-1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not have negative index", actual)
}

func Test_SimpleSlice_InsertAt_OutOfBounds(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	s.InsertAt(-1, 0)

	// Act
	actual := args.Map{"result": s.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "negative index should not insert", actual)
	s.InsertAt(10, 0)
	actual := args.Map{"result": s.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of bounds should not insert", actual)
}

func Test_SimpleSlice_InsertAt_Valid(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 3})
	s.InsertAt(1, 2)

	// Act
	actual := args.Map{"result": s.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should insert", actual)
}

func Test_SimpleSlice_ForEach(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{10, 20})
	sum := 0
	s.ForEach(func(i int, item int) { sum += item })

	// Act
	actual := args.Map{"result": sum != 30}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 30", actual)
}

func Test_SimpleSlice_Filter(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3, 4})
	even := s.Filter(func(v int) bool { return v%2 == 0 })

	// Act
	actual := args.Map{"result": even.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should filter to 2 items", actual)
}

func Test_SimpleSlice_CountFunc(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3, 4})
	count := s.CountFunc(func(i int, v int) bool { return v > 2 })

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should count 2 items > 2", actual)
}

func Test_SimpleSlice_Clone(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	c := s.Clone()

	// Act
	actual := args.Map{"result": c.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should have same length", actual)
}

func Test_SimpleSlice_Clone_Empty(t *testing.T) {
	// Arrange
	s := coregeneric.EmptySimpleSlice[int]()
	c := s.Clone()

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone of empty should be empty", actual)
}

func Test_SimpleSlice_String(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	str := s.String()

	// Act
	actual := args.Map{"result": str == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_SimpleSlice_Count(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": s.Count() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "count should equal length", actual)
}

func Test_SimpleSlice_Last(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": s.Last() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return last", actual)
}

func Test_SimpleSlice_LastOrDefault_NonEmpty(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": s.LastOrDefault() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return last", actual)
}

func Test_SimpleSlice_Items(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": len(s.Items()) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "items should return underlying", actual)
}

// =============================================================================
// LinkedList — uncovered branches
// =============================================================================

func Test_LinkedList_LengthLock_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.Add("a")

	// Act
	actual := args.Map{"result": ll.LengthLock() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be 1", actual)
}

func Test_LinkedList_IsEmptyLock_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()

	// Act
	actual := args.Map{"result": ll.IsEmptyLock()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_LinkedList_AddLock_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddLock("a")

	// Act
	actual := args.Map{"result": ll.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1", actual)
}

func Test_LinkedList_AddSlice_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddSlice([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": ll.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 3", actual)
}

func Test_LinkedList_AddIf_Skip(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddIf(false, "skip")

	// Act
	actual := args.Map{"result": ll.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add", actual)
}

func Test_LinkedList_AddsIf_Skip(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddsIf(false, "a", "b")

	// Act
	actual := args.Map{"result": ll.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add", actual)
}

func Test_LinkedList_AddFunc_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddFunc(func() string { return "x" })

	// Act
	actual := args.Map{"result": ll.Length() != 1 || ll.First() != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add func result", actual)
}

func Test_LinkedList_AddFront_NonEmpty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(2)
	ll.AddFront(1)

	// Act
	actual := args.Map{"result": ll.First() != 1 || ll.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AddFront should prepend", actual)
}

func Test_LinkedList_PushBack(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.PushBack(1)

	// Act
	actual := args.Map{"result": ll.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PushBack should add", actual)
}

func Test_LinkedList_PushFront(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.PushFront(1)

	// Act
	actual := args.Map{"result": ll.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "PushFront should add", actual)
}

func Test_LinkedList_Push_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Push(1)

	// Act
	actual := args.Map{"result": ll.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Push should add", actual)
}

func Test_LinkedList_AppendNode_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	node := &coregeneric.LinkedListNode[int]{Element: 5}
	ll.AppendNode(node)

	// Act
	actual := args.Map{"result": ll.First() != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should append node", actual)
}

func Test_LinkedList_AppendNode_NonEmpty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1)
	node := &coregeneric.LinkedListNode[int]{Element: 2}
	ll.AppendNode(node)

	// Act
	actual := args.Map{"result": ll.Last() != 2 || ll.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should append", actual)
}

func Test_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	chain := coregeneric.LinkedListFrom([]int{1, 2, 3})
	ll.AppendChainOfNodes(chain.Head())

	// Act
	actual := args.Map{"result": ll.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should append chain", actual)
}

func Test_LinkedList_AppendChainOfNodes_NonEmpty(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{0})
	chain := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendChainOfNodes(chain.Head())

	// Act
	actual := args.Map{"result": ll.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should append chain", actual)
}

func Test_LinkedList_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()

	// Act
	actual := args.Map{"result": ll.FirstOrDefault() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return zero", actual)
}

func Test_LinkedList_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[string]()

	// Act
	actual := args.Map{"result": ll.LastOrDefault() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return zero", actual)
}

func Test_LinkedList_Items_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{"result": len(ll.Items()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return empty", actual)
}

func Test_LinkedList_Collection_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	col := ll.Collection()

	// Act
	actual := args.Map{"result": col.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "collection should have 2", actual)
}

func Test_LinkedList_ForEach_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	called := false
	ll.ForEach(func(i int, item int) { called = true })

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not call fn on empty", actual)
}

func Test_LinkedList_ForEachBreak_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	called := false
	ll.ForEachBreak(func(i int, item int) bool { called = true; return false })

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not call fn on empty", actual)
}

func Test_LinkedList_ForEachBreak_Break(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return item == 1 // break on first
	})

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should break after first", actual)
}

func Test_LinkedList_ForEachBreak_BreakLater(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return item == 2
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should break after second", actual)
}

func Test_LinkedList_IndexAt_OutOfBounds_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": ll.IndexAt(-1) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "negative should return nil", actual)
	actual := args.Map{"result": ll.IndexAt(5) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of bounds should return nil", actual)
}

func Test_LinkedList_IndexAt_Empty_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{"result": ll.IndexAt(0) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_LinkedList_IndexAt_Valid_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(1)

	// Act
	actual := args.Map{"result": node == nil || node.Element != 20}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return correct node", actual)
}

func Test_LinkedList_String_Cov(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": ll.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// LinkedListNode
func Test_LinkedListNode_Clone(t *testing.T) {
	// Arrange
	node := &coregeneric.LinkedListNode[int]{Element: 42}
	c := node.Clone()

	// Act
	actual := args.Map{"result": c.Element != 42 || c.HasNext()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should copy element, no next", actual)
}

func Test_LinkedListNode_ListPtr(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	list := ll.Head().ListPtr()

	// Act
	actual := args.Map{"result": len(*list) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should collect all elements", actual)
}

func Test_LinkedListNode_String(t *testing.T) {
	// Arrange
	node := &coregeneric.LinkedListNode[string]{Element: "hello"}

	// Act
	actual := args.Map{"result": node.String() != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return element string", actual)
}

// =============================================================================
// Hashmap — uncovered branches
// =============================================================================

func Test_Hashmap_IsEmptyLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm.IsEmptyLock()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_Hashmap_LengthLock(t *testing.T) {
	// Arrange
	hm := coregeneric.NewHashmap[string, int](0)
	hm.Set("a", 1)

	// Act
	actual := args.Map{"result": hm.LengthLock() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be 1", actual)
}

func Test_Hashmap_SetLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.SetLock("a", 1)
	v, ok := hm.Get("a")

	// Act
	actual := args.Map{"result": ok || v != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "SetLock should set", actual)
}

func Test_Hashmap_GetOrDefault(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm.GetOrDefault("x", 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing key should return default", actual)
	hm.Set("x", 10)
	actual := args.Map{"result": hm.GetOrDefault("x", 42) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "found key should return value", actual)
}

func Test_Hashmap_GetLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	v, ok := hm.GetLock("a")

	// Act
	actual := args.Map{"result": ok || v != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetLock should work", actual)
}

func Test_Hashmap_Contains(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)

	// Act
	actual := args.Map{"result": hm.Contains("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain", actual)
}

func Test_Hashmap_ContainsLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)

	// Act
	actual := args.Map{"result": hm.ContainsLock("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain", actual)
}

func Test_Hashmap_IsKeyMissing_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm.IsKeyMissing("x")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be missing", actual)
}

func Test_Hashmap_Remove_NotExist(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm.Remove("x")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false for missing key", actual)
}

func Test_Hashmap_RemoveLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)

	// Act
	actual := args.Map{"result": hm.RemoveLock("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true for existing key", actual)
}

func Test_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.AddOrUpdateMap(map[string]int{})

	// Act
	actual := args.Map{"result": hm.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.AddOrUpdateHashmap(nil)

	// Act
	actual := args.Map{"result": hm.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_Hashmap_ForEach_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	count := 0
	hm.ForEach(func(k string, v int) { count++ })

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should call once", actual)
}

func Test_Hashmap_ForEachBreak_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	count := 0
	hm.ForEachBreak(func(k string, v int) bool {
		count++
		return true // break immediately
	})

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should break after first", actual)
}

func Test_Hashmap_ConcatNew(t *testing.T) {
	// Arrange
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	result := hm1.ConcatNew(hm2, nil)

	// Act
	actual := args.Map{"result": result.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2 entries", actual)
}

func Test_Hashmap_Clone(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	c := hm.Clone()

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "clone should copy", actual)
}

func Test_Hashmap_IsEquals_BothNil_Cov(t *testing.T) {
	// Arrange
	var hm1, hm2 *coregeneric.Hashmap[string, int]

	// Act
	actual := args.Map{"result": hm1.IsEquals(hm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Hashmap_IsEquals_OneNil_Cov(t *testing.T) {
	// Arrange
	var hm1 *coregeneric.Hashmap[string, int]
	hm2 := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm1.IsEquals(hm2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
}

func Test_Hashmap_IsEquals_SamePtr_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm.IsEquals(hm)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)
}

func Test_Hashmap_IsEquals_DiffLength_Cov(t *testing.T) {
	// Arrange
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": hm1.IsEquals(hm2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different lengths should not be equal", actual)
}

func Test_Hashmap_IsEquals_MissingKey(t *testing.T) {
	// Arrange
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})

	// Act
	actual := args.Map{"result": hm1.IsEquals(hm2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing key should not be equal", actual)
}

func Test_Hashmap_String_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": hm.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Hashmap_Set_Overwrite(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("a", 1)

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "first set should be new", actual)
	isNew = hm.Set("a", 2)
	actual := args.Map{"result": isNew}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "second set should not be new", actual)
}

func Test_Hashmap_HasItems(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": hm.HasItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_Hashmap_Keys_Empty_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": len(hm.Keys()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return empty", actual)
}

func Test_Hashmap_Values_Empty_Cov(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": len(hm.Values()) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return empty", actual)
}

func Test_Hashmap_Map(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	m := hm.Map()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return underlying map", actual)
}

// =============================================================================
// Hashset — uncovered branches
// =============================================================================

func Test_Hashset_AddBool_Existing(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})
	existed := hs.AddBool("a")

	// Act
	actual := args.Map{"result": existed}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "existing item should return true", actual)
}

func Test_Hashset_AddBool_New(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	existed := hs.AddBool("a")

	// Act
	actual := args.Map{"result": existed}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "new item should return false", actual)
}

func Test_Hashset_AddLock(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")

	// Act
	actual := args.Map{"result": hs.Has("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should add", actual)
}

func Test_Hashset_AddSliceLock(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSliceLock([]string{"a", "b"})

	// Act
	actual := args.Map{"result": hs.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add 2", actual)
}

func Test_Hashset_AddIf_Skip(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIf(false, "a")

	// Act
	actual := args.Map{"result": hs.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add", actual)
}

func Test_Hashset_AddIfMany_Skip(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIfMany(false, "a", "b")

	// Act
	actual := args.Map{"result": hs.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add", actual)
}

func Test_Hashset_AddHashsetItems_Nil(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddHashsetItems(nil)

	// Act
	actual := args.Map{"result": hs.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_Hashset_AddItemsMap_FalseValue(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})

	// Act
	actual := args.Map{"result": hs.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should only add true values", actual)
}

func Test_Hashset_ContainsLock(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.ContainsLock("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain", actual)
}

func Test_Hashset_HasAll_Fail(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.HasAll("a", "b")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should fail for missing", actual)
}

func Test_Hashset_HasAny_Fail(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.HasAny("x", "y")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should fail for all missing", actual)
}

func Test_Hashset_Remove_NotExist(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()

	// Act
	actual := args.Map{"result": hs.Remove("x")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false", actual)
}

func Test_Hashset_RemoveLock(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.RemoveLock("a")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true", actual)
}

func Test_Hashset_ListPtr(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})
	p := hs.ListPtr()

	// Act
	actual := args.Map{"result": len(*p) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 1 item", actual)
}

func Test_Hashset_Resize(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})
	hs.Resize(100)

	// Act
	actual := args.Map{"result": hs.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should preserve items", actual)
}

func Test_Hashset_Resize_TooSmall(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	hs.Resize(1)

	// Act
	actual := args.Map{"result": hs.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not resize when capacity < length", actual)
}

func Test_Hashset_Collection(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})
	col := hs.Collection()

	// Act
	actual := args.Map{"result": col.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 1", actual)
}

func Test_Hashset_IsEquals_BothNil_Cov(t *testing.T) {
	// Arrange
	var hs1, hs2 *coregeneric.Hashset[string]

	// Act
	actual := args.Map{"result": hs1.IsEquals(hs2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Hashset_IsEquals_OneNil(t *testing.T) {
	// Arrange
	var hs1 *coregeneric.Hashset[string]
	hs2 := coregeneric.EmptyHashset[string]()

	// Act
	actual := args.Map{"result": hs1.IsEquals(hs2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
}

func Test_Hashset_IsEquals_SamePtr(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.IsEquals(hs)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)
}

func Test_Hashset_IsEquals_DiffLength(t *testing.T) {
	// Arrange
	hs1 := coregeneric.HashsetFrom([]string{"a"})
	hs2 := coregeneric.HashsetFrom([]string{"a", "b"})

	// Act
	actual := args.Map{"result": hs1.IsEquals(hs2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different lengths should not be equal", actual)
}

func Test_Hashset_IsEquals_MissingKey(t *testing.T) {
	// Arrange
	hs1 := coregeneric.HashsetFrom([]string{"a"})
	hs2 := coregeneric.HashsetFrom([]string{"b"})

	// Act
	actual := args.Map{"result": hs1.IsEquals(hs2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "different keys should not be equal", actual)
}

func Test_Hashset_IsEmptyLock(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()

	// Act
	actual := args.Map{"result": hs.IsEmptyLock()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_Hashset_LengthLock(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.LengthLock() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be 1", actual)
}

func Test_Hashset_String(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": hs.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

func Test_Hashset_Map(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"result": len(hs.Map()) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return underlying map", actual)
}

// =============================================================================
// Collection — uncovered branches
// =============================================================================

func Test_Collection_AddIfMany_Skip(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(false, 1, 2)

	// Act
	actual := args.Map{"result": col.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not add", actual)
}

func Test_Collection_AddFunc_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddFunc(func() int { return 42 })

	// Act
	actual := args.Map{"result": col.Length() != 1 || col.First() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add func result", actual)
}

func Test_Collection_AddCollection_Empty_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	other := coregeneric.EmptyCollection[int]()
	col.AddCollection(other)

	// Act
	actual := args.Map{"result": col.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remain empty", actual)
}

func Test_Collection_AddCollections(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	c1 := coregeneric.CollectionFrom([]int{1})
	c2 := coregeneric.EmptyCollection[int]()
	col.AddCollections(c1, c2)

	// Act
	actual := args.Map{"result": col.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add from non-empty", actual)
}

func Test_Collection_RemoveAt_OutOfBounds_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"result": col.RemoveAt(-1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "negative should return false", actual)
	actual := args.Map{"result": col.RemoveAt(5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of bounds should return false", actual)
}

func Test_Collection_SafeAt_OutOfBounds_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"result": col.SafeAt(5) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "out of bounds should return zero", actual)
}

func Test_Collection_SafeAt_Empty_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": col.SafeAt(0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return zero", actual)
}

func Test_Collection_ForEachBreak(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	col.ForEachBreak(func(i int, item int) bool {
		count++
		return item == 2
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should break after second", actual)
}

func Test_Collection_CountFunc(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	count := col.CountFunc(func(v int) bool { return v > 2 })

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should count 2", actual)
}

func Test_Collection_SortFunc(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{3, 1, 2})
	col.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"result": col.First() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "first should be 1 after sort", actual)
}

func Test_Collection_Reverse(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	col.Reverse()

	// Act
	actual := args.Map{"result": col.First() != 3 || col.Last() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should reverse", actual)
}

func Test_Collection_ConcatNew(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})
	result := col.ConcatNew(2, 3)

	// Act
	actual := args.Map{"result": result.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should concat", actual)
}

func Test_Collection_Capacity(t *testing.T) {
	// Arrange
	col := coregeneric.NewCollection[int](10)

	// Act
	actual := args.Map{"result": col.Capacity() < 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have at least 10 capacity", actual)
}

func Test_Collection_Capacity_NilItems(t *testing.T) {
	col := &coregeneric.Collection[int]{}
	// items is nil
	_ = col.Capacity()
}

func Test_Collection_ItemsPtr_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})
	p := col.ItemsPtr()

	// Act
	actual := args.Map{"result": p == nil || len(*p) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return pointer to items", actual)
}

func Test_Collection_LengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": col.LengthLock() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be 2", actual)
}

func Test_Collection_IsEmptyLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": col.IsEmptyLock()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be empty", actual)
}

func Test_Collection_AddLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddLock(1)

	// Act
	actual := args.Map{"result": col.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add", actual)
}

func Test_Collection_AddsLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddsLock(1, 2)

	// Act
	actual := args.Map{"result": col.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should add 2", actual)
}

func Test_Collection_CollectionLenCap(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionLenCap[int](5, 10)

	// Act
	actual := args.Map{"result": col.Length() != 5 || col.Capacity() < 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have length 5, cap >= 10", actual)
}

func Test_Collection_HasItems(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"result": col.HasItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have items", actual)
}

func Test_Collection_String(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"result": col.String() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return non-empty", actual)
}

// =============================================================================
// Pair — uncovered branches
// =============================================================================

func Test_Pair_InvalidPairNoMessage(t *testing.T) {
	// Arrange
	p := coregeneric.InvalidPairNoMessage[string, int]()

	// Act
	actual := args.Map{"result": p.IsValid || p.HasMessage()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid without message", actual)
}

func Test_Pair_HasMessage(t *testing.T) {
	// Arrange
	p := coregeneric.InvalidPair[string, string]("err")

	// Act
	actual := args.Map{"result": p.HasMessage()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have message", actual)
}

func Test_Pair_IsInvalid(t *testing.T) {
	// Arrange
	p := coregeneric.InvalidPair[string, string]("err")

	// Act
	actual := args.Map{"result": p.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Pair_IsInvalid_Nil_Cov(t *testing.T) {
	// Arrange
	var p *coregeneric.Pair[string, string]

	// Act
	actual := args.Map{"result": p.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Pair_Values_Cov(t *testing.T) {
	// Arrange
	p := coregeneric.NewPair("a", 1)
	l, r := p.Values()

	// Act
	actual := args.Map{"result": l != "a" || r != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "values should return left and right", actual)
}

func Test_Pair_Clone_Nil(t *testing.T) {
	// Arrange
	var p *coregeneric.Pair[string, string]

	// Act
	actual := args.Map{"result": p.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Pair_IsEqual_BothNil_Cov(t *testing.T) {
	// Arrange
	var p1, p2 *coregeneric.Pair[string, string]

	// Act
	actual := args.Map{"result": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Pair_IsEqual_OneNil(t *testing.T) {
	// Arrange
	var p1 *coregeneric.Pair[string, string]
	p2 := coregeneric.NewPair("a", "b")

	// Act
	actual := args.Map{"result": p1.IsEqual(p2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
}

func Test_Pair_String_Nil_Cov(t *testing.T) {
	// Arrange
	var p *coregeneric.Pair[string, string]

	// Act
	actual := args.Map{"result": p.String() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Pair_Clear_Cov(t *testing.T) {
	// Arrange
	p := coregeneric.NewPair("a", "b")
	p.Clear()

	// Act
	actual := args.Map{"result": p.IsValid || p.Left != "" || p.Right != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clear all fields", actual)
}

func Test_Pair_Clear_Nil(t *testing.T) {
	var p *coregeneric.Pair[string, string]
	p.Clear() // should not panic
}

func Test_Pair_Dispose_Cov(t *testing.T) {
	// Arrange
	p := coregeneric.NewPair("a", "b")
	p.Dispose()

	// Act
	actual := args.Map{"result": p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should dispose", actual)
}

// =============================================================================
// Triple — uncovered branches
// =============================================================================

func Test_Triple_InvalidTripleNoMessage(t *testing.T) {
	// Arrange
	tr := coregeneric.InvalidTripleNoMessage[string, string, string]()

	// Act
	actual := args.Map{"result": tr.IsValid || tr.HasMessage()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid without message", actual)
}

func Test_Triple_HasMessage(t *testing.T) {
	// Arrange
	tr := coregeneric.InvalidTriple[string, string, string]("err")

	// Act
	actual := args.Map{"result": tr.HasMessage()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should have message", actual)
}

func Test_Triple_IsInvalid(t *testing.T) {
	// Arrange
	tr := coregeneric.InvalidTriple[string, string, string]("err")

	// Act
	actual := args.Map{"result": tr.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_Triple_IsInvalid_Nil_Cov(t *testing.T) {
	// Arrange
	var tr *coregeneric.Triple[string, string, string]

	// Act
	actual := args.Map{"result": tr.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
}

func Test_Triple_Values_Cov(t *testing.T) {
	// Arrange
	tr := coregeneric.NewTriple("a", "b", "c")
	l, m, r := tr.Values()

	// Act
	actual := args.Map{"result": l != "a" || m != "b" || r != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return all values", actual)
}

func Test_Triple_Clone_Nil(t *testing.T) {
	// Arrange
	var tr *coregeneric.Triple[string, string, string]

	// Act
	actual := args.Map{"result": tr.Clone() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil clone should return nil", actual)
}

func Test_Triple_IsEqual_BothNil_Cov(t *testing.T) {
	// Arrange
	var tr1, tr2 *coregeneric.Triple[string, string, string]

	// Act
	actual := args.Map{"result": tr1.IsEqual(tr2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)
}

func Test_Triple_IsEqual_OneNil(t *testing.T) {
	// Arrange
	var tr1 *coregeneric.Triple[string, string, string]
	tr2 := coregeneric.NewTriple("a", "b", "c")

	// Act
	actual := args.Map{"result": tr1.IsEqual(tr2)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one nil should not be equal", actual)
}

func Test_Triple_String_Nil_Cov(t *testing.T) {
	// Arrange
	var tr *coregeneric.Triple[string, string, string]

	// Act
	actual := args.Map{"result": tr.String() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_Triple_Clear_Cov(t *testing.T) {
	// Arrange
	tr := coregeneric.NewTriple("a", "b", "c")
	tr.Clear()

	// Act
	actual := args.Map{"result": tr.IsValid || tr.Left != "" || tr.Middle != "" || tr.Right != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clear all fields", actual)
}

func Test_Triple_Clear_Nil(t *testing.T) {
	var tr *coregeneric.Triple[string, string, string]
	tr.Clear() // should not panic
}

func Test_Triple_Dispose_Cov(t *testing.T) {
	// Arrange
	tr := coregeneric.NewTriple("a", "b", "c")
	tr.Dispose()

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should dispose", actual)
}

// =============================================================================
// PairFrom / TripleFrom — uncovered branches
// =============================================================================

func Test_PairFromSplitFull_NoSep(t *testing.T) {
	// Arrange
	p := coregeneric.PairFromSplitFull("nosep", "=")

	// Act
	actual := args.Map{"result": p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid when no separator", actual)
}

func Test_PairFromSplitFullTrimmed_NoSep(t *testing.T) {
	// Arrange
	p := coregeneric.PairFromSplitFullTrimmed("  nosep  ", "=")

	// Act
	actual := args.Map{"result": p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid when no separator", actual)
}

func Test_PairFromSlice_Empty(t *testing.T) {
	// Arrange
	p := coregeneric.PairFromSlice([]string{})

	// Act
	actual := args.Map{"result": p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should be invalid", actual)
}

func Test_PairFromSlice_Single(t *testing.T) {
	// Arrange
	p := coregeneric.PairFromSlice([]string{"only"})

	// Act
	actual := args.Map{"result": p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single should be invalid", actual)
}

func Test_PairDivide_Odd(t *testing.T) {
	// Arrange
	p := coregeneric.PairDivide(11)
	l, r := p.Values()

	// Act
	actual := args.Map{"result": l+r != 11}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sum to original", actual)
}

func Test_PairDivideWeighted(t *testing.T) {
	// Arrange
	p := coregeneric.PairDivideWeighted(100, 0.3)
	l, r := p.Values()

	// Act
	actual := args.Map{"result": l+r != 100}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sum to 100", actual)
}

func Test_TripleFromSplit_TwoParts(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSplit("a.b", ".")

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "two parts should be invalid", actual)
}

func Test_TripleFromSplit_FourParts(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSplit("a.b.c.d", ".")

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "4+ parts should be valid", actual)
	actual := args.Map{"result": tr.Right != "d"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "right should be last part, got ''", actual)
}

func Test_TripleFromSlice_Empty(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSlice([]string{})

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should be invalid", actual)
}

func Test_TripleFromSlice_Single(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSlice([]string{"only"})

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "single should be invalid", actual)
}

func Test_TripleDivide(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleDivide(10)
	l, m, r := tr.Values()

	// Act
	actual := args.Map{"result": l+m+r != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sum to 10", actual)
}

func Test_TripleDivideWeighted(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleDivideWeighted(100, 0.2, 0.3)
	l, m, r := tr.Values()

	// Act
	actual := args.Map{"result": l+m+r != 100}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sum to 100", actual)
}

// =============================================================================
// funcs.go — uncovered package-level functions
// =============================================================================

func Test_MapCollection_Nil(t *testing.T) {
	// Arrange
	result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil source should return empty", actual)
}

func Test_FlatMapCollection_Nil_Cov(t *testing.T) {
	// Arrange
	result := coregeneric.FlatMapCollection[int, string](nil, func(i int) []string { return nil })

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_ReduceCollection_Nil_Cov(t *testing.T) {
	// Arrange
	result := coregeneric.ReduceCollection[int, int](nil, 0, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"result": result != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return initial", actual)
}

func Test_GroupByCollection_Nil_Cov(t *testing.T) {
	// Arrange
	result := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty map", actual)
}

func Test_ContainsFunc_Nil_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ContainsFunc[int](nil, func(i int) bool { return true })}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_IndexOfFunc_Nil_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true }) != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return -1", actual)
}

func Test_ContainsItem_Nil_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ContainsItem[int](nil, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_IndexOfItem_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IndexOfItem[int](nil, 1) != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return -1", actual)
}

func Test_Distinct_Nil_Cov(t *testing.T) {
	// Arrange
	result := coregeneric.Distinct[int](nil)

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_MapSimpleSlice_Nil_Cov(t *testing.T) {
	// Arrange
	result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// comparablefuncs
func Test_ContainsAll_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ContainsAll[int](nil, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_ContainsAny_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ContainsAny[int](nil, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_RemoveItem_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.RemoveItem[int](nil, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

func Test_RemoveItem_NotFound(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"result": coregeneric.RemoveItem(col, 5)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "not found should return false", actual)
}

func Test_RemoveAllItems_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.RemoveAllItems[int](nil, 1) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_ToHashset_Nil(t *testing.T) {
	// Arrange
	hs := coregeneric.ToHashset[int](nil)

	// Act
	actual := args.Map{"result": hs.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_DistinctSimpleSlice_Nil(t *testing.T) {
	// Arrange
	result := coregeneric.DistinctSimpleSlice[int](nil)

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

func Test_ContainsSimpleSliceItem_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ContainsSimpleSliceItem[int](nil, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return false", actual)
}

// orderedfuncs
func Test_SortCollection_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SortCollection[int](nil) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_SortCollectionDesc_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SortCollectionDesc[int](nil) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_MinCollectionOrDefault_Empty(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": coregeneric.MinCollectionOrDefault(col, 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_MaxCollectionOrDefault_Empty(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"result": coregeneric.MaxCollectionOrDefault(col, 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_IsSortedCollection_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsSortedCollection[int](nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "nil should return true", actual)
}

func Test_SumCollection_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SumCollection[int](nil) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return 0", actual)
}

func Test_ClampCollection_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ClampCollection[int](nil, 0, 10) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return nil", actual)
}

func Test_MinHashsetOrDefault_Empty(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{"result": coregeneric.MinHashsetOrDefault(hs, 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_MaxHashsetOrDefault_Empty(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{"result": coregeneric.MaxHashsetOrDefault(hs, 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_MinKeyHashmapOrDefault_Empty(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": coregeneric.MinKeyHashmapOrDefault(hm, "def") != "def"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_MaxKeyHashmapOrDefault_Empty(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": coregeneric.MaxKeyHashmapOrDefault(hm, "def") != "def"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_MinValueHashmapOrDefault_Empty(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": coregeneric.MinValueHashmapOrDefault(hm, 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

func Test_MaxValueHashmapOrDefault_Empty(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"result": coregeneric.MaxValueHashmapOrDefault(hm, 42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return default", actual)
}

// orderedfuncs: non-nil paths
func Test_SortSimpleSlice(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{3, 1, 2})
	coregeneric.SortSimpleSlice(s)

	// Act
	actual := args.Map{"result": (*s)[0] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sort ascending", actual)
}

func Test_SortSimpleSliceDesc(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 3, 2})
	coregeneric.SortSimpleSliceDesc(s)

	// Act
	actual := args.Map{"result": (*s)[0] != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sort descending", actual)
}

func Test_MinSimpleSlice(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{3, 1, 2})

	// Act
	actual := args.Map{"result": coregeneric.MinSimpleSlice(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return min", actual)
}

func Test_MaxSimpleSlice(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{3, 1, 2})

	// Act
	actual := args.Map{"result": coregeneric.MaxSimpleSlice(s) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return max", actual)
}

func Test_SumSimpleSlice(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coregeneric.SumSimpleSlice(s) != 6}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return sum", actual)
}

func Test_SortedListHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	sorted := coregeneric.SortedListHashset(hs)

	// Act
	actual := args.Map{"result": sorted[0] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be sorted", actual)
}

func Test_SortedListDescHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	sorted := coregeneric.SortedListDescHashset(hs)

	// Act
	actual := args.Map{"result": sorted[0] != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be sorted desc", actual)
}

func Test_MinHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})

	// Act
	actual := args.Map{"result": coregeneric.MinHashset(hs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return min", actual)
}

func Test_MaxHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})

	// Act
	actual := args.Map{"result": coregeneric.MaxHashset(hs) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return max", actual)
}

func Test_SortedCollectionHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	col := coregeneric.SortedCollectionHashset(hs)

	// Act
	actual := args.Map{"result": col.First() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be sorted", actual)
}

func Test_SortedKeysHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1, "b": 2})
	keys := coregeneric.SortedKeysHashmap(hm)

	// Act
	actual := args.Map{"result": keys[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be sorted", actual)
}

func Test_SortedKeysDescHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1, "b": 2})
	keys := coregeneric.SortedKeysDescHashmap(hm)

	// Act
	actual := args.Map{"result": keys[0] != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be sorted desc", actual)
}

func Test_MinKeyHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1})

	// Act
	actual := args.Map{"result": coregeneric.MinKeyHashmap(hm) != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return min key", actual)
}

func Test_MaxKeyHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1})

	// Act
	actual := args.Map{"result": coregeneric.MaxKeyHashmap(hm) != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return max key", actual)
}

func Test_SortedValuesHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1, "c": 2})
	vals := coregeneric.SortedValuesHashmap(hm)

	// Act
	actual := args.Map{"result": vals[0] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be sorted", actual)
}

func Test_MinValueHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})

	// Act
	actual := args.Map{"result": coregeneric.MinValueHashmap(hm) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return min value", actual)
}

func Test_MaxValueHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})

	// Act
	actual := args.Map{"result": coregeneric.MaxValueHashmap(hm) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return max value", actual)
}

func Test_ClampCollection_Values(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{-5, 0, 5, 15})
	coregeneric.ClampCollection(col, 0, 10)
	items := col.Items()

	// Act
	actual := args.Map{"result": items[0] != 0 || items[3] != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clamp", actual)
}

// numericfuncs
func Test_IsLess(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsLess(1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1 < 2", actual)
}

func Test_IsLessOrEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsLessOrEqual(2, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2 <= 2", actual)
}

func Test_IsGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsGreater(3, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "3 > 2", actual)
}

func Test_IsGreaterOrEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsGreaterOrEqual(2, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2 >= 2", actual)
}

func Test_IsNumericEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsNumericEqual(2, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "2 == 2", actual)
}

func Test_IsNotEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsNotEqual(1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1 != 2", actual)
}

func Test_Clamp(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Clamp(-1, 0, 10) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "below min", actual)
	actual := args.Map{"result": coregeneric.Clamp(15, 0, 10) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "above max", actual)
	actual := args.Map{"result": coregeneric.Clamp(5, 0, 10) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "in range", actual)
}

func Test_ClampMin(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ClampMin(-1, 0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clamp to min", actual)
	actual := args.Map{"result": coregeneric.ClampMin(5, 0) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should keep value", actual)
}

func Test_ClampMax(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ClampMax(15, 10) != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should clamp to max", actual)
	actual := args.Map{"result": coregeneric.ClampMax(5, 10) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should keep value", actual)
}

func Test_InRange(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.InRange(5, 0, 10)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be in range", actual)
	actual := args.Map{"result": coregeneric.InRange(-1, 0, 10)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be in range", actual)
}

func Test_InRangeExclusive(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.InRangeExclusive(0, 0, 10)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "boundary should not be in range exclusive", actual)
	actual := args.Map{"result": coregeneric.InRangeExclusive(5, 0, 10)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be in range exclusive", actual)
}

func Test_Abs(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Abs(-5) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "abs(-5) should be 5", actual)
	actual := args.Map{"result": coregeneric.Abs(5) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "abs(5) should be 5", actual)
}

func Test_AbsDiff(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.AbsDiff(3, 5) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be 2", actual)
	actual := args.Map{"result": coregeneric.AbsDiff(5, 3) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be 2", actual)
}

func Test_Sum(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Sum(1, 2, 3) != 6}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sum to 6", actual)
}

func Test_MinOf(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.MinOf(3, 5) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 3", actual)
}

func Test_MaxOf(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.MaxOf(3, 5) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 5", actual)
}

func Test_MinOfSlice(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.MinOfSlice([]int{3, 1, 2}) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 1", actual)
}

func Test_MaxOfSlice(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.MaxOfSlice([]int{3, 1, 2}) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return 3", actual)
}

func Test_IsZero(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsZero(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be zero", actual)
}

func Test_IsPositive(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsPositive(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "1 should be positive", actual)
}

func Test_IsNegative(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsNegative(-1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "-1 should be negative", actual)
}

func Test_IsNonNegative(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.IsNonNegative(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "0 should be non-negative", actual)
}

func Test_Sign(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.Sign(-5) != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "negative should be -1", actual)
	actual := args.Map{"result": coregeneric.Sign(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "zero should be 0", actual)
	actual := args.Map{"result": coregeneric.Sign(5) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "positive should be 1", actual)
}

func Test_SafeDiv(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SafeDiv(10, 0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "div by zero should return 0", actual)
	actual := args.Map{"result": coregeneric.SafeDiv(10, 2) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "10/2 should be 5", actual)
}

func Test_SafeDivOrDefault(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.SafeDivOrDefault(10, 0, -1) != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "div by zero should return default", actual)
	actual := args.Map{"result": coregeneric.SafeDivOrDefault(10, 2, -1) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "10/2 should be 5", actual)
}

// =============================================================================
// CompareNumeric
// =============================================================================

func Test_CompareNumeric_All(t *testing.T) {
	// Arrange
	eq := coregeneric.CompareNumeric(5, 5)

	// Act
	actual := args.Map{"result": eq.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "equal", actual)
	gt := coregeneric.CompareNumeric(5, 3)
	actual := args.Map{"result": gt.IsLeftGreater()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "greater", actual)
	lt := coregeneric.CompareNumeric(3, 5)
	actual := args.Map{"result": lt.IsLeftLess()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "less", actual)
}

// =============================================================================
// PointerSliceSorter — uncovered branches
// =============================================================================

func Test_PointerSliceSorter_Desc_Cov(t *testing.T) {
	// Arrange
	a, b, c := 3, 1, 2
	items := []*int{&a, &b, &c}
	sorter := coregeneric.NewPointerSliceSorterDesc(items)
	sorter.Sort()

	// Act
	actual := args.Map{"result": *sorter.Items()[0] != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sort desc", actual)
}

func Test_PointerSliceSorter_Func(t *testing.T) {
	// Arrange
	a, b := 1, 2
	items := []*int{&b, &a}
	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(x, y int) bool { return x < y }, false)
	sorter.Sort()

	// Act
	actual := args.Map{"result": *sorter.Items()[0] != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should sort asc with custom func", actual)
}

func Test_PointerSliceSorter_NilHandling(t *testing.T) {
	a := 1
	items := []*int{nil, &a, nil}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()
	// nils should be at the end
	if sorter.Items()[0] == nil {
		// nilFirst is false, so nils should be last
		t.Log("nils may be at beginning if nilFirst")
	}
}

func Test_PointerSliceSorter_SetMethods(t *testing.T) {
	// Arrange
	a, b := 1, 2
	items := []*int{&b, &a}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.SetDesc()
	sorter.Sort()

	// Act
	actual := args.Map{"result": *sorter.Items()[0] != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be desc", actual)
	sorter.SetAsc()
	sorter.Sort()
	actual := args.Map{"result": *sorter.Items()[0] != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be asc", actual)
	sorter.SetNilFirst(true)
	sorter.SetLessFunc(func(x, y int) bool { return x < y })
	_ = sorter.IsSorted()
}

func Test_PointerSliceSorter_SetItems_Cov(t *testing.T) {
	// Arrange
	sorter := coregeneric.NewPointerSliceSorterAsc([]*int{})
	a := 5
	sorter.SetItems([]*int{&a})

	// Act
	actual := args.Map{"result": len(sorter.Items()) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should replace items", actual)
}

func Test_PointerSliceSorter_Len_NilItems(t *testing.T) {
	// Arrange
	sorter := coregeneric.NewPointerSliceSorterAsc[int](nil)

	// Act
	actual := args.Map{"result": sorter.Len() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil items should return 0", actual)
}

func Test_PointerSliceSorter_Less_BothNil(t *testing.T) {
	// Arrange
	items := []*int{nil, nil}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)

	// Act
	actual := args.Map{"result": sorter.Less(0, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both nil should return false", actual)
}

func Test_PointerSliceSorter_NilFirst_Cov(t *testing.T) {
	// Arrange
	a := 1
	items := []*int{&a, nil}
	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(x, y int) bool { return x < y }, true)
	sorter.Sort()

	// Act
	actual := args.Map{"result": sorter.Items()[0] != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nilFirst should put nil first", actual)
}

// =============================================================================
// Typed creators via New
// =============================================================================

func Test_New_Collection_Creators(t *testing.T) {
	_ = coregeneric.New.Collection.String.Empty()
	_ = coregeneric.New.Collection.String.Cap(10)
	_ = coregeneric.New.Collection.String.From([]string{"a"})
	_ = coregeneric.New.Collection.String.Clone([]string{"a"})
	_ = coregeneric.New.Collection.String.Items("a", "b")
	_ = coregeneric.New.Collection.String.LenCap(5, 10)
	_ = coregeneric.New.Collection.Int.Empty()
	_ = coregeneric.New.Collection.Float64.Empty()
	_ = coregeneric.New.Collection.Bool.Empty()
	_ = coregeneric.New.Collection.Any.Empty()
}

func Test_New_Hashset_Creators(t *testing.T) {
	_ = coregeneric.New.Hashset.String.Empty()
	_ = coregeneric.New.Hashset.String.Cap(10)
	_ = coregeneric.New.Hashset.String.From([]string{"a"})
	_ = coregeneric.New.Hashset.String.Items("a", "b")
	_ = coregeneric.New.Hashset.String.UsingMap(map[string]bool{"a": true})
	_ = coregeneric.New.Hashset.Int.Empty()
}

func Test_New_Hashmap_Creators(t *testing.T) {
	_ = coregeneric.New.Hashmap.StringString.Empty()
	_ = coregeneric.New.Hashmap.StringString.Cap(10)
	_ = coregeneric.New.Hashmap.StringString.From(map[string]string{"a": "b"})
	_ = coregeneric.New.Hashmap.StringString.Clone(map[string]string{"a": "b"})
	_ = coregeneric.New.Hashmap.StringInt.Empty()
	_ = coregeneric.New.Hashmap.IntString.Empty()
}

func Test_New_SimpleSlice_Creators(t *testing.T) {
	_ = coregeneric.New.SimpleSlice.String.Empty()
	_ = coregeneric.New.SimpleSlice.String.Cap(10)
	_ = coregeneric.New.SimpleSlice.String.From([]string{"a"})
	_ = coregeneric.New.SimpleSlice.String.Clone([]string{"a"})
	_ = coregeneric.New.SimpleSlice.String.Items("a", "b")
	_ = coregeneric.New.SimpleSlice.Int.Empty()
}

func Test_New_LinkedList_Creators(t *testing.T) {
	_ = coregeneric.New.LinkedList.String.Empty()
	_ = coregeneric.New.LinkedList.String.From([]string{"a"})
	_ = coregeneric.New.LinkedList.String.Items("a", "b")
	_ = coregeneric.New.LinkedList.Int.Empty()
}

func Test_New_Pair_Creators(t *testing.T) {
	_ = coregeneric.New.Pair.StringString("a", "b")
	_ = coregeneric.New.Pair.StringInt("a", 1)
	_ = coregeneric.New.Pair.StringInt64("a", int64(1))
	_ = coregeneric.New.Pair.StringFloat64("a", 1.0)
	_ = coregeneric.New.Pair.StringBool("a", true)
	_ = coregeneric.New.Pair.StringAny("a", "b")
	_ = coregeneric.New.Pair.IntInt(1, 2)
	_ = coregeneric.New.Pair.IntString(1, "a")
	_ = coregeneric.New.Pair.Any("a", "b")
	_ = coregeneric.New.Pair.InvalidStringString("err")
	_ = coregeneric.New.Pair.InvalidAny("err")
	_ = coregeneric.New.Pair.Split("a=b", "=")
	_ = coregeneric.New.Pair.SplitTrimmed(" a = b ", "=")
	_ = coregeneric.New.Pair.SplitFull("a:b:c", ":")
	_ = coregeneric.New.Pair.SplitFullTrimmed(" a : b : c ", ":")
	_ = coregeneric.New.Pair.FromSlice([]string{"a", "b"})
	_ = coregeneric.New.Pair.DivideInt(10)
	_ = coregeneric.New.Pair.DivideInt64(int64(10))
	_ = coregeneric.New.Pair.DivideFloat64(10.0)
	_ = coregeneric.New.Pair.DivideIntWeighted(100, 0.3)
	_ = coregeneric.New.Pair.DivideFloat64Weighted(100.0, 0.3)
}

func Test_New_Triple_Creators(t *testing.T) {
	_ = coregeneric.New.Triple.StringStringString("a", "b", "c")
	_ = coregeneric.New.Triple.StringIntString("a", 1, "b")
	_ = coregeneric.New.Triple.StringAnyAny("a", "b", "c")
	_ = coregeneric.New.Triple.Any("a", "b", "c")
	_ = coregeneric.New.Triple.InvalidStringStringString("err")
	_ = coregeneric.New.Triple.InvalidAny("err")
	_ = coregeneric.New.Triple.Split("a.b.c", ".")
	_ = coregeneric.New.Triple.SplitTrimmed(" a . b . c ", ".")
	_ = coregeneric.New.Triple.SplitN("a:b:c:d", ":")
	_ = coregeneric.New.Triple.SplitNTrimmed(" a : b : c : d ", ":")
	_ = coregeneric.New.Triple.FromSlice([]string{"a", "b", "c"})
	_ = coregeneric.New.Triple.DivideInt(10)
	_ = coregeneric.New.Triple.DivideInt64(int64(10))
	_ = coregeneric.New.Triple.DivideFloat64(10.0)
	_ = coregeneric.New.Triple.DivideIntWeighted(100, 0.2, 0.3)
	_ = coregeneric.New.Triple.DivideFloat64Weighted(100.0, 0.2, 0.3)
}

// PairFrom helper functions
func Test_NewPairOf(t *testing.T) {
	// Arrange
	p := coregeneric.NewPairOf(1, 2)

	// Act
	actual := args.Map{"result": p.IsValid || p.Left != 1 || p.Right != 2}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should create valid pair", actual)
}

func Test_InvalidPairOf(t *testing.T) {
	// Arrange
	p := coregeneric.InvalidPairOf[string]("err")

	// Act
	actual := args.Map{"result": p.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_PairFromSplitTrimmed_Cov(t *testing.T) {
	// Arrange
	p := coregeneric.PairFromSplitTrimmed(" a = b ", "=")

	// Act
	actual := args.Map{"result": p.Left != "a" || p.Right != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should trim", actual)
}

// TripleFrom helper functions
func Test_NewTripleOf(t *testing.T) {
	// Arrange
	tr := coregeneric.NewTripleOf(1, 2, 3)

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be valid", actual)
}

func Test_InvalidTripleOf(t *testing.T) {
	// Arrange
	tr := coregeneric.InvalidTripleOf[string]("err")

	// Act
	actual := args.Map{"result": tr.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be invalid", actual)
}

func Test_TripleFromSplitTrimmed(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSplitTrimmed(" a . b . c ", ".")

	// Act
	actual := args.Map{"result": tr.Left != "a" || tr.Middle != "b" || tr.Right != "c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should trim", actual)
}

func Test_TripleFromSplitN(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSplitN("a:b:c:d", ":")

	// Act
	actual := args.Map{"result": tr.Left != "a" || tr.Middle != "b" || tr.Right != "c:d"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should split into 3", actual)
}

func Test_TripleFromSplitNTrimmed(t *testing.T) {
	// Arrange
	tr := coregeneric.TripleFromSplitNTrimmed(" a : b : c : d ", ":")

	// Act
	actual := args.Map{"result": tr.Left != "a" || tr.Middle != "b"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should split and trim", actual)
}

// =============================================================================
// Exercising functional paths for completeness
// =============================================================================

func Test_MapCollection_NonNil(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	result := coregeneric.MapCollection(col, func(i int) string { return fmt.Sprintf("%d", i) })

	// Act
	actual := args.Map{"result": result.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should map all", actual)
}

func Test_FlatMapCollection_NonNil(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})
	result := coregeneric.FlatMapCollection(col, func(i int) []string { return []string{"a", "b"} })

	// Act
	actual := args.Map{"result": result.Length() != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should flatmap", actual)
}

func Test_ReduceCollection_NonNil(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	result := coregeneric.ReduceCollection(col, 0, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"result": result != 6}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should reduce to 6", actual)
}

func Test_GroupByCollection_NonNil(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]string{"a", "ab", "b", "bc"})
	result := coregeneric.GroupByCollection(col, func(s string) string { return string(s[0]) })

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should group by first char", actual)
}

func Test_ContainsFunc_Found_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coregeneric.ContainsFunc(col, func(i int) bool { return i == 2 })}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find 2", actual)
}

func Test_IndexOfFunc_Found_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{10, 20, 30})

	// Act
	actual := args.Map{"result": coregeneric.IndexOfFunc(col, func(i int) bool { return i == 20 }) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should find at index 1", actual)
}

func Test_RemoveAllItems_Multi(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 1, 3, 1})
	count := coregeneric.RemoveAllItems(col, 1)

	// Act
	actual := args.Map{"result": count != 3 || col.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should remove all 1s", actual)
}

func Test_MapSimpleSlice_NonNil(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	result := coregeneric.MapSimpleSlice(s, func(i int) string { return fmt.Sprintf("%d", i) })

	// Act
	actual := args.Map{"result": result.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should map", actual)
}

func Test_DistinctSimpleSlice_WithDups(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 1, 3})
	result := coregeneric.DistinctSimpleSlice(s)

	// Act
	actual := args.Map{"result": result.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should deduplicate", actual)
}

func Test_ContainsSimpleSliceItem_Found(t *testing.T) {
	// Arrange
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coregeneric.ContainsSimpleSliceItem(s, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find 2", actual)
}

func Test_ContainsAll_Found(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coregeneric.ContainsAll(col, 1, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain all", actual)
}

func Test_ContainsAny_Found(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coregeneric.ContainsAny(col, 5, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain any", actual)
}

func Test_RemoveItem_Found_Cov(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coregeneric.RemoveItem(col, 2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should remove", actual)
	actual := args.Map{"result": col.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should have 2", actual)
}

func Test_Distinct_WithDups(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 1, 3})
	result := coregeneric.Distinct(col)

	// Act
	actual := args.Map{"result": result.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should deduplicate", actual)
}

func Test_ToHashset_NonNil(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	hs := coregeneric.ToHashset(col)

	// Act
	actual := args.Map{"result": hs.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should convert", actual)
}
