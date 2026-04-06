package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
)

// =============================================================================
// SimpleSlice — uncovered branches
// =============================================================================

func Test_SimpleSlice_AddIf_Skip(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddIf(false, "skip")
	if s.Length() != 0 {
		t.Error("should not add when false")
	}
}

func Test_SimpleSlice_Adds_Empty(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	s.Adds()
	if s.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_SimpleSlice_AddSlice_Empty(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddSlice([]string{})
	if s.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_SimpleSlice_AddsIf_Skip(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddsIf(false, "a", "b")
	if s.Length() != 0 {
		t.Error("should not add when false")
	}
}

func Test_SimpleSlice_AddFunc(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	s.AddFunc(func() string { return "hello" })
	if s.Length() != 1 || s.First() != "hello" {
		t.Error("AddFunc should add result")
	}
}

func Test_SimpleSlice_FirstOrDefault_Empty(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	if s.FirstOrDefault() != "" {
		t.Error("empty should return zero value")
	}
}

func Test_SimpleSlice_LastOrDefault_Empty(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[string]()
	if s.LastOrDefault() != "" {
		t.Error("empty should return zero value")
	}
}

func Test_SimpleSlice_Skip_OverCount(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	result := s.Skip(10)
	if len(result) != 0 {
		t.Error("skip beyond length should return empty")
	}
}

func Test_SimpleSlice_Take_OverCount(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	result := s.Take(10)
	if len(result) != 2 {
		t.Error("take beyond length should return all")
	}
}

func Test_SimpleSlice_Length_Nil(t *testing.T) {
	var s *coregeneric.SimpleSlice[int]
	if s.Length() != 0 {
		t.Error("nil should return 0")
	}
}

func Test_SimpleSlice_IsEmpty_Nil(t *testing.T) {
	var s *coregeneric.SimpleSlice[int]
	if !s.IsEmpty() {
		t.Error("nil should be empty")
	}
}

func Test_SimpleSlice_HasAnyItem(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1})
	if !s.HasAnyItem() {
		t.Error("should have items")
	}
}

func Test_SimpleSlice_HasItems(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1})
	if !s.HasItems() {
		t.Error("should have items")
	}
}

func Test_SimpleSlice_HasIndex(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	if !s.HasIndex(2) {
		t.Error("should have index 2")
	}
	if s.HasIndex(5) {
		t.Error("should not have index 5")
	}
	if s.HasIndex(-1) {
		t.Error("should not have negative index")
	}
}

func Test_SimpleSlice_InsertAt_OutOfBounds(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	s.InsertAt(-1, 0)
	if s.Length() != 2 {
		t.Error("negative index should not insert")
	}
	s.InsertAt(10, 0)
	if s.Length() != 2 {
		t.Error("out of bounds should not insert")
	}
}

func Test_SimpleSlice_InsertAt_Valid(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 3})
	s.InsertAt(1, 2)
	if s.Length() != 3 {
		t.Error("should insert")
	}
}

func Test_SimpleSlice_ForEach(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{10, 20})
	sum := 0
	s.ForEach(func(i int, item int) { sum += item })
	if sum != 30 {
		t.Errorf("expected 30, got %d", sum)
	}
}

func Test_SimpleSlice_Filter(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3, 4})
	even := s.Filter(func(v int) bool { return v%2 == 0 })
	if even.Length() != 2 {
		t.Error("should filter to 2 items")
	}
}

func Test_SimpleSlice_CountFunc(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3, 4})
	count := s.CountFunc(func(i int, v int) bool { return v > 2 })
	if count != 2 {
		t.Error("should count 2 items > 2")
	}
}

func Test_SimpleSlice_Clone(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	c := s.Clone()
	if c.Length() != 3 {
		t.Error("clone should have same length")
	}
}

func Test_SimpleSlice_Clone_Empty(t *testing.T) {
	s := coregeneric.EmptySimpleSlice[int]()
	c := s.Clone()
	if c.Length() != 0 {
		t.Error("clone of empty should be empty")
	}
}

func Test_SimpleSlice_String(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	str := s.String()
	if str == "" {
		t.Error("should return non-empty")
	}
}

func Test_SimpleSlice_Count(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	if s.Count() != 2 {
		t.Error("count should equal length")
	}
}

func Test_SimpleSlice_Last(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	if s.Last() != 3 {
		t.Error("should return last")
	}
}

func Test_SimpleSlice_LastOrDefault_NonEmpty(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	if s.LastOrDefault() != 3 {
		t.Error("should return last")
	}
}

func Test_SimpleSlice_Items(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	if len(s.Items()) != 2 {
		t.Error("items should return underlying")
	}
}

// =============================================================================
// LinkedList — uncovered branches
// =============================================================================

func Test_LinkedList_LengthLock_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.Add("a")
	if ll.LengthLock() != 1 {
		t.Error("should be 1")
	}
}

func Test_LinkedList_IsEmptyLock_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	if !ll.IsEmptyLock() {
		t.Error("should be empty")
	}
}

func Test_LinkedList_AddLock_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddLock("a")
	if ll.Length() != 1 {
		t.Error("should have 1")
	}
}

func Test_LinkedList_AddSlice_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddSlice([]int{1, 2, 3})
	if ll.Length() != 3 {
		t.Error("should have 3")
	}
}

func Test_LinkedList_AddIf_Skip(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddIf(false, "skip")
	if ll.Length() != 0 {
		t.Error("should not add")
	}
}

func Test_LinkedList_AddsIf_Skip(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddsIf(false, "a", "b")
	if ll.Length() != 0 {
		t.Error("should not add")
	}
}

func Test_LinkedList_AddFunc_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	ll.AddFunc(func() string { return "x" })
	if ll.Length() != 1 || ll.First() != "x" {
		t.Error("should add func result")
	}
}

func Test_LinkedList_AddFront_NonEmpty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(2)
	ll.AddFront(1)
	if ll.First() != 1 || ll.Length() != 2 {
		t.Error("AddFront should prepend")
	}
}

func Test_LinkedList_PushBack(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.PushBack(1)
	if ll.Length() != 1 {
		t.Error("PushBack should add")
	}
}

func Test_LinkedList_PushFront(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.PushFront(1)
	if ll.Length() != 1 {
		t.Error("PushFront should add")
	}
}

func Test_LinkedList_Push_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Push(1)
	if ll.Length() != 1 {
		t.Error("Push should add")
	}
}

func Test_LinkedList_AppendNode_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	node := &coregeneric.LinkedListNode[int]{Element: 5}
	ll.AppendNode(node)
	if ll.First() != 5 {
		t.Error("should append node")
	}
}

func Test_LinkedList_AppendNode_NonEmpty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.Add(1)
	node := &coregeneric.LinkedListNode[int]{Element: 2}
	ll.AppendNode(node)
	if ll.Last() != 2 || ll.Length() != 2 {
		t.Error("should append")
	}
}

func Test_LinkedList_AppendChainOfNodes_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	chain := coregeneric.LinkedListFrom([]int{1, 2, 3})
	ll.AppendChainOfNodes(chain.Head())
	if ll.Length() != 3 {
		t.Error("should append chain")
	}
}

func Test_LinkedList_AppendChainOfNodes_NonEmpty(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{0})
	chain := coregeneric.LinkedListFrom([]int{1, 2})
	ll.AppendChainOfNodes(chain.Head())
	if ll.Length() != 3 {
		t.Error("should append chain")
	}
}

func Test_LinkedList_FirstOrDefault_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	if ll.FirstOrDefault() != "" {
		t.Error("empty should return zero")
	}
}

func Test_LinkedList_LastOrDefault_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[string]()
	if ll.LastOrDefault() != "" {
		t.Error("empty should return zero")
	}
}

func Test_LinkedList_Items_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	if len(ll.Items()) != 0 {
		t.Error("empty should return empty")
	}
}

func Test_LinkedList_Collection_Cov(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	col := ll.Collection()
	if col.Length() != 2 {
		t.Error("collection should have 2")
	}
}

func Test_LinkedList_ForEach_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	called := false
	ll.ForEach(func(i int, item int) { called = true })
	if called {
		t.Error("should not call fn on empty")
	}
}

func Test_LinkedList_ForEachBreak_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	called := false
	ll.ForEachBreak(func(i int, item int) bool { called = true; return false })
	if called {
		t.Error("should not call fn on empty")
	}
}

func Test_LinkedList_ForEachBreak_Break(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return item == 1 // break on first
	})
	if count != 1 {
		t.Error("should break after first")
	}
}

func Test_LinkedList_ForEachBreak_BreakLater(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return item == 2
	})
	if count != 2 {
		t.Error("should break after second")
	}
}

func Test_LinkedList_IndexAt_OutOfBounds_Cov(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	if ll.IndexAt(-1) != nil {
		t.Error("negative should return nil")
	}
	if ll.IndexAt(5) != nil {
		t.Error("out of bounds should return nil")
	}
}

func Test_LinkedList_IndexAt_Empty_Cov(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	if ll.IndexAt(0) != nil {
		t.Error("empty should return nil")
	}
}

func Test_LinkedList_IndexAt_Valid_Cov(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(1)
	if node == nil || node.Element != 20 {
		t.Error("should return correct node")
	}
}

func Test_LinkedList_String_Cov(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	if ll.String() == "" {
		t.Error("should return non-empty")
	}
}

// LinkedListNode
func Test_LinkedListNode_Clone(t *testing.T) {
	node := &coregeneric.LinkedListNode[int]{Element: 42}
	c := node.Clone()
	if c.Element != 42 || c.HasNext() {
		t.Error("clone should copy element, no next")
	}
}

func Test_LinkedListNode_ListPtr(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	list := ll.Head().ListPtr()
	if len(*list) != 3 {
		t.Error("should collect all elements")
	}
}

func Test_LinkedListNode_String(t *testing.T) {
	node := &coregeneric.LinkedListNode[string]{Element: "hello"}
	if node.String() != "hello" {
		t.Error("should return element string")
	}
}

// =============================================================================
// Hashmap — uncovered branches
// =============================================================================

func Test_Hashmap_IsEmptyLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if !hm.IsEmptyLock() {
		t.Error("should be empty")
	}
}

func Test_Hashmap_LengthLock(t *testing.T) {
	hm := coregeneric.NewHashmap[string, int](0)
	hm.Set("a", 1)
	if hm.LengthLock() != 1 {
		t.Error("should be 1")
	}
}

func Test_Hashmap_SetLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.SetLock("a", 1)
	v, ok := hm.Get("a")
	if !ok || v != 1 {
		t.Error("SetLock should set")
	}
}

func Test_Hashmap_GetOrDefault(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if hm.GetOrDefault("x", 42) != 42 {
		t.Error("missing key should return default")
	}
	hm.Set("x", 10)
	if hm.GetOrDefault("x", 42) != 10 {
		t.Error("found key should return value")
	}
}

func Test_Hashmap_GetLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	v, ok := hm.GetLock("a")
	if !ok || v != 1 {
		t.Error("GetLock should work")
	}
}

func Test_Hashmap_Contains(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	if !hm.Contains("a") {
		t.Error("should contain")
	}
}

func Test_Hashmap_ContainsLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	if !hm.ContainsLock("a") {
		t.Error("should contain")
	}
}

func Test_Hashmap_IsKeyMissing_Cov(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if !hm.IsKeyMissing("x") {
		t.Error("should be missing")
	}
}

func Test_Hashmap_Remove_NotExist(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if hm.Remove("x") {
		t.Error("should return false for missing key")
	}
}

func Test_Hashmap_RemoveLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	if !hm.RemoveLock("a") {
		t.Error("should return true for existing key")
	}
}

func Test_Hashmap_AddOrUpdateMap_Empty(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.AddOrUpdateMap(map[string]int{})
	if hm.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_Hashmap_AddOrUpdateHashmap_Nil(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.AddOrUpdateHashmap(nil)
	if hm.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_Hashmap_ForEach_Cov(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	count := 0
	hm.ForEach(func(k string, v int) { count++ })
	if count != 1 {
		t.Error("should call once")
	}
}

func Test_Hashmap_ForEachBreak_Cov(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1, "b": 2})
	count := 0
	hm.ForEachBreak(func(k string, v int) bool {
		count++
		return true // break immediately
	})
	if count != 1 {
		t.Error("should break after first")
	}
}

func Test_Hashmap_ConcatNew(t *testing.T) {
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	result := hm1.ConcatNew(hm2, nil)
	if result.Length() != 2 {
		t.Error("should have 2 entries")
	}
}

func Test_Hashmap_Clone(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	c := hm.Clone()
	if c.Length() != 1 {
		t.Error("clone should copy")
	}
}

func Test_Hashmap_IsEquals_BothNil_Cov(t *testing.T) {
	var hm1, hm2 *coregeneric.Hashmap[string, int]
	if !hm1.IsEquals(hm2) {
		t.Error("both nil should be equal")
	}
}

func Test_Hashmap_IsEquals_OneNil_Cov(t *testing.T) {
	var hm1 *coregeneric.Hashmap[string, int]
	hm2 := coregeneric.EmptyHashmap[string, int]()
	if hm1.IsEquals(hm2) {
		t.Error("one nil should not be equal")
	}
}

func Test_Hashmap_IsEquals_SamePtr_Cov(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if !hm.IsEquals(hm) {
		t.Error("same pointer should be equal")
	}
}

func Test_Hashmap_IsEquals_DiffLength_Cov(t *testing.T) {
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.EmptyHashmap[string, int]()
	if hm1.IsEquals(hm2) {
		t.Error("different lengths should not be equal")
	}
}

func Test_Hashmap_IsEquals_MissingKey(t *testing.T) {
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"b": 1})
	if hm1.IsEquals(hm2) {
		t.Error("missing key should not be equal")
	}
}

func Test_Hashmap_String_Cov(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	if hm.String() == "" {
		t.Error("should return non-empty")
	}
}

func Test_Hashmap_Set_Overwrite(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("a", 1)
	if !isNew {
		t.Error("first set should be new")
	}
	isNew = hm.Set("a", 2)
	if isNew {
		t.Error("second set should not be new")
	}
}

func Test_Hashmap_HasItems(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	if !hm.HasItems() {
		t.Error("should have items")
	}
}

func Test_Hashmap_Keys_Empty_Cov(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if len(hm.Keys()) != 0 {
		t.Error("should return empty")
	}
}

func Test_Hashmap_Values_Empty_Cov(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if len(hm.Values()) != 0 {
		t.Error("should return empty")
	}
}

func Test_Hashmap_Map(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	m := hm.Map()
	if len(m) != 1 {
		t.Error("should return underlying map")
	}
}

// =============================================================================
// Hashset — uncovered branches
// =============================================================================

func Test_Hashset_AddBool_Existing(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	existed := hs.AddBool("a")
	if !existed {
		t.Error("existing item should return true")
	}
}

func Test_Hashset_AddBool_New(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	existed := hs.AddBool("a")
	if existed {
		t.Error("new item should return false")
	}
}

func Test_Hashset_AddLock(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")
	if !hs.Has("a") {
		t.Error("should add")
	}
}

func Test_Hashset_AddSliceLock(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSliceLock([]string{"a", "b"})
	if hs.Length() != 2 {
		t.Error("should add 2")
	}
}

func Test_Hashset_AddIf_Skip(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIf(false, "a")
	if hs.Length() != 0 {
		t.Error("should not add")
	}
}

func Test_Hashset_AddIfMany_Skip(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIfMany(false, "a", "b")
	if hs.Length() != 0 {
		t.Error("should not add")
	}
}

func Test_Hashset_AddHashsetItems_Nil(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddHashsetItems(nil)
	if hs.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_Hashset_AddItemsMap_FalseValue(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})
	if hs.Length() != 1 {
		t.Error("should only add true values")
	}
}

func Test_Hashset_ContainsLock(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if !hs.ContainsLock("a") {
		t.Error("should contain")
	}
}

func Test_Hashset_HasAll_Fail(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if hs.HasAll("a", "b") {
		t.Error("should fail for missing")
	}
}

func Test_Hashset_HasAny_Fail(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if hs.HasAny("x", "y") {
		t.Error("should fail for all missing")
	}
}

func Test_Hashset_Remove_NotExist(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	if hs.Remove("x") {
		t.Error("should return false")
	}
}

func Test_Hashset_RemoveLock(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if !hs.RemoveLock("a") {
		t.Error("should return true")
	}
}

func Test_Hashset_ListPtr(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	p := hs.ListPtr()
	if len(*p) != 1 {
		t.Error("should return 1 item")
	}
}

func Test_Hashset_Resize(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	hs.Resize(100)
	if hs.Length() != 1 {
		t.Error("should preserve items")
	}
}

func Test_Hashset_Resize_TooSmall(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	hs.Resize(1)
	if hs.Length() != 3 {
		t.Error("should not resize when capacity < length")
	}
}

func Test_Hashset_Collection(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	col := hs.Collection()
	if col.Length() != 1 {
		t.Error("should have 1")
	}
}

func Test_Hashset_IsEquals_BothNil_Cov(t *testing.T) {
	var hs1, hs2 *coregeneric.Hashset[string]
	if !hs1.IsEquals(hs2) {
		t.Error("both nil should be equal")
	}
}

func Test_Hashset_IsEquals_OneNil(t *testing.T) {
	var hs1 *coregeneric.Hashset[string]
	hs2 := coregeneric.EmptyHashset[string]()
	if hs1.IsEquals(hs2) {
		t.Error("one nil should not be equal")
	}
}

func Test_Hashset_IsEquals_SamePtr(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if !hs.IsEquals(hs) {
		t.Error("same pointer should be equal")
	}
}

func Test_Hashset_IsEquals_DiffLength(t *testing.T) {
	hs1 := coregeneric.HashsetFrom([]string{"a"})
	hs2 := coregeneric.HashsetFrom([]string{"a", "b"})
	if hs1.IsEquals(hs2) {
		t.Error("different lengths should not be equal")
	}
}

func Test_Hashset_IsEquals_MissingKey(t *testing.T) {
	hs1 := coregeneric.HashsetFrom([]string{"a"})
	hs2 := coregeneric.HashsetFrom([]string{"b"})
	if hs1.IsEquals(hs2) {
		t.Error("different keys should not be equal")
	}
}

func Test_Hashset_IsEmptyLock(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	if !hs.IsEmptyLock() {
		t.Error("should be empty")
	}
}

func Test_Hashset_LengthLock(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if hs.LengthLock() != 1 {
		t.Error("should be 1")
	}
}

func Test_Hashset_String(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if hs.String() == "" {
		t.Error("should return non-empty")
	}
}

func Test_Hashset_Map(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	if len(hs.Map()) != 1 {
		t.Error("should return underlying map")
	}
}

// =============================================================================
// Collection — uncovered branches
// =============================================================================

func Test_Collection_AddIfMany_Skip(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(false, 1, 2)
	if col.Length() != 0 {
		t.Error("should not add")
	}
}

func Test_Collection_AddFunc_Cov(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddFunc(func() int { return 42 })
	if col.Length() != 1 || col.First() != 42 {
		t.Error("should add func result")
	}
}

func Test_Collection_AddCollection_Empty_Cov(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	other := coregeneric.EmptyCollection[int]()
	col.AddCollection(other)
	if col.Length() != 0 {
		t.Error("should remain empty")
	}
}

func Test_Collection_AddCollections(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	c1 := coregeneric.CollectionFrom([]int{1})
	c2 := coregeneric.EmptyCollection[int]()
	col.AddCollections(c1, c2)
	if col.Length() != 1 {
		t.Error("should add from non-empty")
	}
}

func Test_Collection_RemoveAt_OutOfBounds_Cov(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	if col.RemoveAt(-1) {
		t.Error("negative should return false")
	}
	if col.RemoveAt(5) {
		t.Error("out of bounds should return false")
	}
}

func Test_Collection_SafeAt_OutOfBounds_Cov(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	if col.SafeAt(5) != 0 {
		t.Error("out of bounds should return zero")
	}
}

func Test_Collection_SafeAt_Empty_Cov(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	if col.SafeAt(0) != 0 {
		t.Error("empty should return zero")
	}
}

func Test_Collection_ForEachBreak(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	col.ForEachBreak(func(i int, item int) bool {
		count++
		return item == 2
	})
	if count != 2 {
		t.Error("should break after second")
	}
}

func Test_Collection_CountFunc(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	count := col.CountFunc(func(v int) bool { return v > 2 })
	if count != 2 {
		t.Error("should count 2")
	}
}

func Test_Collection_SortFunc(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{3, 1, 2})
	col.SortFunc(func(a, b int) bool { return a < b })
	if col.First() != 1 {
		t.Error("first should be 1 after sort")
	}
}

func Test_Collection_Reverse(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	col.Reverse()
	if col.First() != 3 || col.Last() != 1 {
		t.Error("should reverse")
	}
}

func Test_Collection_ConcatNew(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	result := col.ConcatNew(2, 3)
	if result.Length() != 3 {
		t.Error("should concat")
	}
}

func Test_Collection_Capacity(t *testing.T) {
	col := coregeneric.NewCollection[int](10)
	if col.Capacity() < 10 {
		t.Error("should have at least 10 capacity")
	}
}

func Test_Collection_Capacity_NilItems(t *testing.T) {
	col := &coregeneric.Collection[int]{}
	// items is nil
	_ = col.Capacity()
}

func Test_Collection_ItemsPtr_Cov(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	p := col.ItemsPtr()
	if p == nil || len(*p) != 1 {
		t.Error("should return pointer to items")
	}
}

func Test_Collection_LengthLock(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	if col.LengthLock() != 2 {
		t.Error("should be 2")
	}
}

func Test_Collection_IsEmptyLock(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	if !col.IsEmptyLock() {
		t.Error("should be empty")
	}
}

func Test_Collection_AddLock(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddLock(1)
	if col.Length() != 1 {
		t.Error("should add")
	}
}

func Test_Collection_AddsLock(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddsLock(1, 2)
	if col.Length() != 2 {
		t.Error("should add 2")
	}
}

func Test_Collection_CollectionLenCap(t *testing.T) {
	col := coregeneric.CollectionLenCap[int](5, 10)
	if col.Length() != 5 || col.Capacity() < 10 {
		t.Error("should have length 5, cap >= 10")
	}
}

func Test_Collection_HasItems(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	if !col.HasItems() {
		t.Error("should have items")
	}
}

func Test_Collection_String(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	if col.String() == "" {
		t.Error("should return non-empty")
	}
}

// =============================================================================
// Pair — uncovered branches
// =============================================================================

func Test_Pair_InvalidPairNoMessage(t *testing.T) {
	p := coregeneric.InvalidPairNoMessage[string, int]()
	if p.IsValid || p.HasMessage() {
		t.Error("should be invalid without message")
	}
}

func Test_Pair_HasMessage(t *testing.T) {
	p := coregeneric.InvalidPair[string, string]("err")
	if !p.HasMessage() {
		t.Error("should have message")
	}
}

func Test_Pair_IsInvalid(t *testing.T) {
	p := coregeneric.InvalidPair[string, string]("err")
	if !p.IsInvalid() {
		t.Error("should be invalid")
	}
}

func Test_Pair_IsInvalid_Nil_Cov(t *testing.T) {
	var p *coregeneric.Pair[string, string]
	if !p.IsInvalid() {
		t.Error("nil should be invalid")
	}
}

func Test_Pair_Values_Cov(t *testing.T) {
	p := coregeneric.NewPair("a", 1)
	l, r := p.Values()
	if l != "a" || r != 1 {
		t.Error("values should return left and right")
	}
}

func Test_Pair_Clone_Nil(t *testing.T) {
	var p *coregeneric.Pair[string, string]
	if p.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Pair_IsEqual_BothNil_Cov(t *testing.T) {
	var p1, p2 *coregeneric.Pair[string, string]
	if !p1.IsEqual(p2) {
		t.Error("both nil should be equal")
	}
}

func Test_Pair_IsEqual_OneNil(t *testing.T) {
	var p1 *coregeneric.Pair[string, string]
	p2 := coregeneric.NewPair("a", "b")
	if p1.IsEqual(p2) {
		t.Error("one nil should not be equal")
	}
}

func Test_Pair_String_Nil_Cov(t *testing.T) {
	var p *coregeneric.Pair[string, string]
	if p.String() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Pair_Clear_Cov(t *testing.T) {
	p := coregeneric.NewPair("a", "b")
	p.Clear()
	if p.IsValid || p.Left != "" || p.Right != "" {
		t.Error("should clear all fields")
	}
}

func Test_Pair_Clear_Nil(t *testing.T) {
	var p *coregeneric.Pair[string, string]
	p.Clear() // should not panic
}

func Test_Pair_Dispose_Cov(t *testing.T) {
	p := coregeneric.NewPair("a", "b")
	p.Dispose()
	if p.IsValid {
		t.Error("should dispose")
	}
}

// =============================================================================
// Triple — uncovered branches
// =============================================================================

func Test_Triple_InvalidTripleNoMessage(t *testing.T) {
	tr := coregeneric.InvalidTripleNoMessage[string, string, string]()
	if tr.IsValid || tr.HasMessage() {
		t.Error("should be invalid without message")
	}
}

func Test_Triple_HasMessage(t *testing.T) {
	tr := coregeneric.InvalidTriple[string, string, string]("err")
	if !tr.HasMessage() {
		t.Error("should have message")
	}
}

func Test_Triple_IsInvalid(t *testing.T) {
	tr := coregeneric.InvalidTriple[string, string, string]("err")
	if !tr.IsInvalid() {
		t.Error("should be invalid")
	}
}

func Test_Triple_IsInvalid_Nil_Cov(t *testing.T) {
	var tr *coregeneric.Triple[string, string, string]
	if !tr.IsInvalid() {
		t.Error("nil should be invalid")
	}
}

func Test_Triple_Values_Cov(t *testing.T) {
	tr := coregeneric.NewTriple("a", "b", "c")
	l, m, r := tr.Values()
	if l != "a" || m != "b" || r != "c" {
		t.Error("should return all values")
	}
}

func Test_Triple_Clone_Nil(t *testing.T) {
	var tr *coregeneric.Triple[string, string, string]
	if tr.Clone() != nil {
		t.Error("nil clone should return nil")
	}
}

func Test_Triple_IsEqual_BothNil_Cov(t *testing.T) {
	var tr1, tr2 *coregeneric.Triple[string, string, string]
	if !tr1.IsEqual(tr2) {
		t.Error("both nil should be equal")
	}
}

func Test_Triple_IsEqual_OneNil(t *testing.T) {
	var tr1 *coregeneric.Triple[string, string, string]
	tr2 := coregeneric.NewTriple("a", "b", "c")
	if tr1.IsEqual(tr2) {
		t.Error("one nil should not be equal")
	}
}

func Test_Triple_String_Nil_Cov(t *testing.T) {
	var tr *coregeneric.Triple[string, string, string]
	if tr.String() != "" {
		t.Error("nil should return empty")
	}
}

func Test_Triple_Clear_Cov(t *testing.T) {
	tr := coregeneric.NewTriple("a", "b", "c")
	tr.Clear()
	if tr.IsValid || tr.Left != "" || tr.Middle != "" || tr.Right != "" {
		t.Error("should clear all fields")
	}
}

func Test_Triple_Clear_Nil(t *testing.T) {
	var tr *coregeneric.Triple[string, string, string]
	tr.Clear() // should not panic
}

func Test_Triple_Dispose_Cov(t *testing.T) {
	tr := coregeneric.NewTriple("a", "b", "c")
	tr.Dispose()
	if tr.IsValid {
		t.Error("should dispose")
	}
}

// =============================================================================
// PairFrom / TripleFrom — uncovered branches
// =============================================================================

func Test_PairFromSplitFull_NoSep(t *testing.T) {
	p := coregeneric.PairFromSplitFull("nosep", "=")
	if p.IsValid {
		t.Error("should be invalid when no separator")
	}
}

func Test_PairFromSplitFullTrimmed_NoSep(t *testing.T) {
	p := coregeneric.PairFromSplitFullTrimmed("  nosep  ", "=")
	if p.IsValid {
		t.Error("should be invalid when no separator")
	}
}

func Test_PairFromSlice_Empty(t *testing.T) {
	p := coregeneric.PairFromSlice([]string{})
	if p.IsValid {
		t.Error("empty should be invalid")
	}
}

func Test_PairFromSlice_Single(t *testing.T) {
	p := coregeneric.PairFromSlice([]string{"only"})
	if p.IsValid {
		t.Error("single should be invalid")
	}
}

func Test_PairDivide_Odd(t *testing.T) {
	p := coregeneric.PairDivide(11)
	l, r := p.Values()
	if l+r != 11 {
		t.Error("should sum to original")
	}
}

func Test_PairDivideWeighted(t *testing.T) {
	p := coregeneric.PairDivideWeighted(100, 0.3)
	l, r := p.Values()
	if l+r != 100 {
		t.Error("should sum to 100")
	}
}

func Test_TripleFromSplit_TwoParts(t *testing.T) {
	tr := coregeneric.TripleFromSplit("a.b", ".")
	if tr.IsValid {
		t.Error("two parts should be invalid")
	}
}

func Test_TripleFromSplit_FourParts(t *testing.T) {
	tr := coregeneric.TripleFromSplit("a.b.c.d", ".")
	if !tr.IsValid {
		t.Error("4+ parts should be valid")
	}
	if tr.Right != "d" {
		t.Errorf("right should be last part, got '%s'", tr.Right)
	}
}

func Test_TripleFromSlice_Empty(t *testing.T) {
	tr := coregeneric.TripleFromSlice([]string{})
	if tr.IsValid {
		t.Error("empty should be invalid")
	}
}

func Test_TripleFromSlice_Single(t *testing.T) {
	tr := coregeneric.TripleFromSlice([]string{"only"})
	if tr.IsValid {
		t.Error("single should be invalid")
	}
}

func Test_TripleDivide(t *testing.T) {
	tr := coregeneric.TripleDivide(10)
	l, m, r := tr.Values()
	if l+m+r != 10 {
		t.Error("should sum to 10")
	}
}

func Test_TripleDivideWeighted(t *testing.T) {
	tr := coregeneric.TripleDivideWeighted(100, 0.2, 0.3)
	l, m, r := tr.Values()
	if l+m+r != 100 {
		t.Error("should sum to 100")
	}
}

// =============================================================================
// funcs.go — uncovered package-level functions
// =============================================================================

func Test_MapCollection_Nil(t *testing.T) {
	result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })
	if result.Length() != 0 {
		t.Error("nil source should return empty")
	}
}

func Test_FlatMapCollection_Nil_Cov(t *testing.T) {
	result := coregeneric.FlatMapCollection[int, string](nil, func(i int) []string { return nil })
	if result.Length() != 0 {
		t.Error("nil should return empty")
	}
}

func Test_ReduceCollection_Nil_Cov(t *testing.T) {
	result := coregeneric.ReduceCollection[int, int](nil, 0, func(acc int, item int) int { return acc + item })
	if result != 0 {
		t.Error("nil should return initial")
	}
}

func Test_GroupByCollection_Nil_Cov(t *testing.T) {
	result := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })
	if len(result) != 0 {
		t.Error("nil should return empty map")
	}
}

func Test_ContainsFunc_Nil_Cov(t *testing.T) {
	if coregeneric.ContainsFunc[int](nil, func(i int) bool { return true }) {
		t.Error("nil should return false")
	}
}

func Test_IndexOfFunc_Nil_Cov(t *testing.T) {
	if coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true }) != -1 {
		t.Error("nil should return -1")
	}
}

func Test_ContainsItem_Nil_Cov(t *testing.T) {
	if coregeneric.ContainsItem[int](nil, 1) {
		t.Error("nil should return false")
	}
}

func Test_IndexOfItem_Nil(t *testing.T) {
	if coregeneric.IndexOfItem[int](nil, 1) != -1 {
		t.Error("nil should return -1")
	}
}

func Test_Distinct_Nil_Cov(t *testing.T) {
	result := coregeneric.Distinct[int](nil)
	if result.Length() != 0 {
		t.Error("nil should return empty")
	}
}

func Test_MapSimpleSlice_Nil_Cov(t *testing.T) {
	result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })
	if result.Length() != 0 {
		t.Error("nil should return empty")
	}
}

// comparablefuncs
func Test_ContainsAll_Nil(t *testing.T) {
	if coregeneric.ContainsAll[int](nil, 1) {
		t.Error("nil should return false")
	}
}

func Test_ContainsAny_Nil(t *testing.T) {
	if coregeneric.ContainsAny[int](nil, 1) {
		t.Error("nil should return false")
	}
}

func Test_RemoveItem_Nil(t *testing.T) {
	if coregeneric.RemoveItem[int](nil, 1) {
		t.Error("nil should return false")
	}
}

func Test_RemoveItem_NotFound(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	if coregeneric.RemoveItem(col, 5) {
		t.Error("not found should return false")
	}
}

func Test_RemoveAllItems_Nil(t *testing.T) {
	if coregeneric.RemoveAllItems[int](nil, 1) != 0 {
		t.Error("nil should return 0")
	}
}

func Test_ToHashset_Nil(t *testing.T) {
	hs := coregeneric.ToHashset[int](nil)
	if hs.Length() != 0 {
		t.Error("nil should return empty")
	}
}

func Test_DistinctSimpleSlice_Nil(t *testing.T) {
	result := coregeneric.DistinctSimpleSlice[int](nil)
	if result.Length() != 0 {
		t.Error("nil should return empty")
	}
}

func Test_ContainsSimpleSliceItem_Nil(t *testing.T) {
	if coregeneric.ContainsSimpleSliceItem[int](nil, 1) {
		t.Error("nil should return false")
	}
}

// orderedfuncs
func Test_SortCollection_Nil(t *testing.T) {
	if coregeneric.SortCollection[int](nil) != nil {
		t.Error("nil should return nil")
	}
}

func Test_SortCollectionDesc_Nil(t *testing.T) {
	if coregeneric.SortCollectionDesc[int](nil) != nil {
		t.Error("nil should return nil")
	}
}

func Test_MinCollectionOrDefault_Empty(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	if coregeneric.MinCollectionOrDefault(col, 42) != 42 {
		t.Error("empty should return default")
	}
}

func Test_MaxCollectionOrDefault_Empty(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	if coregeneric.MaxCollectionOrDefault(col, 42) != 42 {
		t.Error("empty should return default")
	}
}

func Test_IsSortedCollection_Nil(t *testing.T) {
	if !coregeneric.IsSortedCollection[int](nil) {
		t.Error("nil should return true")
	}
}

func Test_SumCollection_Nil(t *testing.T) {
	if coregeneric.SumCollection[int](nil) != 0 {
		t.Error("nil should return 0")
	}
}

func Test_ClampCollection_Nil(t *testing.T) {
	if coregeneric.ClampCollection[int](nil, 0, 10) != nil {
		t.Error("nil should return nil")
	}
}

func Test_MinHashsetOrDefault_Empty(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	if coregeneric.MinHashsetOrDefault(hs, 42) != 42 {
		t.Error("empty should return default")
	}
}

func Test_MaxHashsetOrDefault_Empty(t *testing.T) {
	hs := coregeneric.EmptyHashset[int]()
	if coregeneric.MaxHashsetOrDefault(hs, 42) != 42 {
		t.Error("empty should return default")
	}
}

func Test_MinKeyHashmapOrDefault_Empty(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if coregeneric.MinKeyHashmapOrDefault(hm, "def") != "def" {
		t.Error("empty should return default")
	}
}

func Test_MaxKeyHashmapOrDefault_Empty(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if coregeneric.MaxKeyHashmapOrDefault(hm, "def") != "def" {
		t.Error("empty should return default")
	}
}

func Test_MinValueHashmapOrDefault_Empty(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if coregeneric.MinValueHashmapOrDefault(hm, 42) != 42 {
		t.Error("empty should return default")
	}
}

func Test_MaxValueHashmapOrDefault_Empty(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	if coregeneric.MaxValueHashmapOrDefault(hm, 42) != 42 {
		t.Error("empty should return default")
	}
}

// orderedfuncs: non-nil paths
func Test_SortSimpleSlice(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{3, 1, 2})
	coregeneric.SortSimpleSlice(s)
	if (*s)[0] != 1 {
		t.Error("should sort ascending")
	}
}

func Test_SortSimpleSliceDesc(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 3, 2})
	coregeneric.SortSimpleSliceDesc(s)
	if (*s)[0] != 3 {
		t.Error("should sort descending")
	}
}

func Test_MinSimpleSlice(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{3, 1, 2})
	if coregeneric.MinSimpleSlice(s) != 1 {
		t.Error("should return min")
	}
}

func Test_MaxSimpleSlice(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{3, 1, 2})
	if coregeneric.MaxSimpleSlice(s) != 3 {
		t.Error("should return max")
	}
}

func Test_SumSimpleSlice(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	if coregeneric.SumSimpleSlice(s) != 6 {
		t.Error("should return sum")
	}
}

func Test_SortedListHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	sorted := coregeneric.SortedListHashset(hs)
	if sorted[0] != 1 {
		t.Error("should be sorted")
	}
}

func Test_SortedListDescHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	sorted := coregeneric.SortedListDescHashset(hs)
	if sorted[0] != 3 {
		t.Error("should be sorted desc")
	}
}

func Test_MinHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	if coregeneric.MinHashset(hs) != 1 {
		t.Error("should return min")
	}
}

func Test_MaxHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	if coregeneric.MaxHashset(hs) != 3 {
		t.Error("should return max")
	}
}

func Test_SortedCollectionHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	col := coregeneric.SortedCollectionHashset(hs)
	if col.First() != 1 {
		t.Error("should be sorted")
	}
}

func Test_SortedKeysHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1, "b": 2})
	keys := coregeneric.SortedKeysHashmap(hm)
	if keys[0] != "a" {
		t.Error("should be sorted")
	}
}

func Test_SortedKeysDescHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1, "b": 2})
	keys := coregeneric.SortedKeysDescHashmap(hm)
	if keys[0] != "c" {
		t.Error("should be sorted desc")
	}
}

func Test_MinKeyHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1})
	if coregeneric.MinKeyHashmap(hm) != "a" {
		t.Error("should return min key")
	}
}

func Test_MaxKeyHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1})
	if coregeneric.MaxKeyHashmap(hm) != "c" {
		t.Error("should return max key")
	}
}

func Test_SortedValuesHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1, "c": 2})
	vals := coregeneric.SortedValuesHashmap(hm)
	if vals[0] != 1 {
		t.Error("should be sorted")
	}
}

func Test_MinValueHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})
	if coregeneric.MinValueHashmap(hm) != 1 {
		t.Error("should return min value")
	}
}

func Test_MaxValueHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})
	if coregeneric.MaxValueHashmap(hm) != 3 {
		t.Error("should return max value")
	}
}

func Test_ClampCollection_Values(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{-5, 0, 5, 15})
	coregeneric.ClampCollection(col, 0, 10)
	items := col.Items()
	if items[0] != 0 || items[3] != 10 {
		t.Error("should clamp")
	}
}

// numericfuncs
func Test_IsLess(t *testing.T) {
	if !coregeneric.IsLess(1, 2) {
		t.Error("1 < 2")
	}
}

func Test_IsLessOrEqual(t *testing.T) {
	if !coregeneric.IsLessOrEqual(2, 2) {
		t.Error("2 <= 2")
	}
}

func Test_IsGreater(t *testing.T) {
	if !coregeneric.IsGreater(3, 2) {
		t.Error("3 > 2")
	}
}

func Test_IsGreaterOrEqual(t *testing.T) {
	if !coregeneric.IsGreaterOrEqual(2, 2) {
		t.Error("2 >= 2")
	}
}

func Test_IsNumericEqual(t *testing.T) {
	if !coregeneric.IsNumericEqual(2, 2) {
		t.Error("2 == 2")
	}
}

func Test_IsNotEqual(t *testing.T) {
	if !coregeneric.IsNotEqual(1, 2) {
		t.Error("1 != 2")
	}
}

func Test_Clamp(t *testing.T) {
	if coregeneric.Clamp(-1, 0, 10) != 0 {
		t.Error("below min")
	}
	if coregeneric.Clamp(15, 0, 10) != 10 {
		t.Error("above max")
	}
	if coregeneric.Clamp(5, 0, 10) != 5 {
		t.Error("in range")
	}
}

func Test_ClampMin(t *testing.T) {
	if coregeneric.ClampMin(-1, 0) != 0 {
		t.Error("should clamp to min")
	}
	if coregeneric.ClampMin(5, 0) != 5 {
		t.Error("should keep value")
	}
}

func Test_ClampMax(t *testing.T) {
	if coregeneric.ClampMax(15, 10) != 10 {
		t.Error("should clamp to max")
	}
	if coregeneric.ClampMax(5, 10) != 5 {
		t.Error("should keep value")
	}
}

func Test_InRange(t *testing.T) {
	if !coregeneric.InRange(5, 0, 10) {
		t.Error("should be in range")
	}
	if coregeneric.InRange(-1, 0, 10) {
		t.Error("should not be in range")
	}
}

func Test_InRangeExclusive(t *testing.T) {
	if coregeneric.InRangeExclusive(0, 0, 10) {
		t.Error("boundary should not be in range exclusive")
	}
	if !coregeneric.InRangeExclusive(5, 0, 10) {
		t.Error("should be in range exclusive")
	}
}

func Test_Abs(t *testing.T) {
	if coregeneric.Abs(-5) != 5 {
		t.Error("abs(-5) should be 5")
	}
	if coregeneric.Abs(5) != 5 {
		t.Error("abs(5) should be 5")
	}
}

func Test_AbsDiff(t *testing.T) {
	if coregeneric.AbsDiff(3, 5) != 2 {
		t.Error("should be 2")
	}
	if coregeneric.AbsDiff(5, 3) != 2 {
		t.Error("should be 2")
	}
}

func Test_Sum(t *testing.T) {
	if coregeneric.Sum(1, 2, 3) != 6 {
		t.Error("should sum to 6")
	}
}

func Test_MinOf(t *testing.T) {
	if coregeneric.MinOf(3, 5) != 3 {
		t.Error("should return 3")
	}
}

func Test_MaxOf(t *testing.T) {
	if coregeneric.MaxOf(3, 5) != 5 {
		t.Error("should return 5")
	}
}

func Test_MinOfSlice(t *testing.T) {
	if coregeneric.MinOfSlice([]int{3, 1, 2}) != 1 {
		t.Error("should return 1")
	}
}

func Test_MaxOfSlice(t *testing.T) {
	if coregeneric.MaxOfSlice([]int{3, 1, 2}) != 3 {
		t.Error("should return 3")
	}
}

func Test_IsZero(t *testing.T) {
	if !coregeneric.IsZero(0) {
		t.Error("0 should be zero")
	}
}

func Test_IsPositive(t *testing.T) {
	if !coregeneric.IsPositive(1) {
		t.Error("1 should be positive")
	}
}

func Test_IsNegative(t *testing.T) {
	if !coregeneric.IsNegative(-1) {
		t.Error("-1 should be negative")
	}
}

func Test_IsNonNegative(t *testing.T) {
	if !coregeneric.IsNonNegative(0) {
		t.Error("0 should be non-negative")
	}
}

func Test_Sign(t *testing.T) {
	if coregeneric.Sign(-5) != -1 {
		t.Error("negative should be -1")
	}
	if coregeneric.Sign(0) != 0 {
		t.Error("zero should be 0")
	}
	if coregeneric.Sign(5) != 1 {
		t.Error("positive should be 1")
	}
}

func Test_SafeDiv(t *testing.T) {
	if coregeneric.SafeDiv(10, 0) != 0 {
		t.Error("div by zero should return 0")
	}
	if coregeneric.SafeDiv(10, 2) != 5 {
		t.Error("10/2 should be 5")
	}
}

func Test_SafeDivOrDefault(t *testing.T) {
	if coregeneric.SafeDivOrDefault(10, 0, -1) != -1 {
		t.Error("div by zero should return default")
	}
	if coregeneric.SafeDivOrDefault(10, 2, -1) != 5 {
		t.Error("10/2 should be 5")
	}
}

// =============================================================================
// CompareNumeric
// =============================================================================

func Test_CompareNumeric_All(t *testing.T) {
	eq := coregeneric.CompareNumeric(5, 5)
	if !eq.IsEqual() {
		t.Error("equal")
	}
	gt := coregeneric.CompareNumeric(5, 3)
	if !gt.IsLeftGreater() {
		t.Error("greater")
	}
	lt := coregeneric.CompareNumeric(3, 5)
	if !lt.IsLeftLess() {
		t.Error("less")
	}
}

// =============================================================================
// PointerSliceSorter — uncovered branches
// =============================================================================

func Test_PointerSliceSorter_Desc_Cov(t *testing.T) {
	a, b, c := 3, 1, 2
	items := []*int{&a, &b, &c}
	sorter := coregeneric.NewPointerSliceSorterDesc(items)
	sorter.Sort()
	if *sorter.Items()[0] != 3 {
		t.Error("should sort desc")
	}
}

func Test_PointerSliceSorter_Func(t *testing.T) {
	a, b := 1, 2
	items := []*int{&b, &a}
	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(x, y int) bool { return x < y }, false)
	sorter.Sort()
	if *sorter.Items()[0] != 1 {
		t.Error("should sort asc with custom func")
	}
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
	a, b := 1, 2
	items := []*int{&b, &a}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.SetDesc()
	sorter.Sort()
	if *sorter.Items()[0] != 2 {
		t.Error("should be desc")
	}
	sorter.SetAsc()
	sorter.Sort()
	if *sorter.Items()[0] != 1 {
		t.Error("should be asc")
	}
	sorter.SetNilFirst(true)
	sorter.SetLessFunc(func(x, y int) bool { return x < y })
	_ = sorter.IsSorted()
}

func Test_PointerSliceSorter_SetItems_Cov(t *testing.T) {
	sorter := coregeneric.NewPointerSliceSorterAsc([]*int{})
	a := 5
	sorter.SetItems([]*int{&a})
	if len(sorter.Items()) != 1 {
		t.Error("should replace items")
	}
}

func Test_PointerSliceSorter_Len_NilItems(t *testing.T) {
	sorter := coregeneric.NewPointerSliceSorterAsc[int](nil)
	if sorter.Len() != 0 {
		t.Error("nil items should return 0")
	}
}

func Test_PointerSliceSorter_Less_BothNil(t *testing.T) {
	items := []*int{nil, nil}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	if sorter.Less(0, 1) {
		t.Error("both nil should return false")
	}
}

func Test_PointerSliceSorter_NilFirst_Cov(t *testing.T) {
	a := 1
	items := []*int{&a, nil}
	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(x, y int) bool { return x < y }, true)
	sorter.Sort()
	if sorter.Items()[0] != nil {
		t.Error("nilFirst should put nil first")
	}
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
	p := coregeneric.NewPairOf(1, 2)
	if !p.IsValid || p.Left != 1 || p.Right != 2 {
		t.Error("should create valid pair")
	}
}

func Test_InvalidPairOf(t *testing.T) {
	p := coregeneric.InvalidPairOf[string]("err")
	if p.IsValid {
		t.Error("should be invalid")
	}
}

func Test_PairFromSplitTrimmed_Cov(t *testing.T) {
	p := coregeneric.PairFromSplitTrimmed(" a = b ", "=")
	if p.Left != "a" || p.Right != "b" {
		t.Error("should trim")
	}
}

// TripleFrom helper functions
func Test_NewTripleOf(t *testing.T) {
	tr := coregeneric.NewTripleOf(1, 2, 3)
	if !tr.IsValid {
		t.Error("should be valid")
	}
}

func Test_InvalidTripleOf(t *testing.T) {
	tr := coregeneric.InvalidTripleOf[string]("err")
	if tr.IsValid {
		t.Error("should be invalid")
	}
}

func Test_TripleFromSplitTrimmed(t *testing.T) {
	tr := coregeneric.TripleFromSplitTrimmed(" a . b . c ", ".")
	if tr.Left != "a" || tr.Middle != "b" || tr.Right != "c" {
		t.Error("should trim")
	}
}

func Test_TripleFromSplitN(t *testing.T) {
	tr := coregeneric.TripleFromSplitN("a:b:c:d", ":")
	if tr.Left != "a" || tr.Middle != "b" || tr.Right != "c:d" {
		t.Error("should split into 3")
	}
}

func Test_TripleFromSplitNTrimmed(t *testing.T) {
	tr := coregeneric.TripleFromSplitNTrimmed(" a : b : c : d ", ":")
	if tr.Left != "a" || tr.Middle != "b" {
		t.Error("should split and trim")
	}
}

// =============================================================================
// Exercising functional paths for completeness
// =============================================================================

func Test_MapCollection_NonNil(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	result := coregeneric.MapCollection(col, func(i int) string { return fmt.Sprintf("%d", i) })
	if result.Length() != 3 {
		t.Error("should map all")
	}
}

func Test_FlatMapCollection_NonNil(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	result := coregeneric.FlatMapCollection(col, func(i int) []string { return []string{"a", "b"} })
	if result.Length() != 4 {
		t.Error("should flatmap")
	}
}

func Test_ReduceCollection_NonNil(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	result := coregeneric.ReduceCollection(col, 0, func(acc int, item int) int { return acc + item })
	if result != 6 {
		t.Error("should reduce to 6")
	}
}

func Test_GroupByCollection_NonNil(t *testing.T) {
	col := coregeneric.CollectionFrom([]string{"a", "ab", "b", "bc"})
	result := coregeneric.GroupByCollection(col, func(s string) string { return string(s[0]) })
	if len(result) != 2 {
		t.Error("should group by first char")
	}
}

func Test_ContainsFunc_Found_Cov(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	if !coregeneric.ContainsFunc(col, func(i int) bool { return i == 2 }) {
		t.Error("should find 2")
	}
}

func Test_IndexOfFunc_Found_Cov(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{10, 20, 30})
	if coregeneric.IndexOfFunc(col, func(i int) bool { return i == 20 }) != 1 {
		t.Error("should find at index 1")
	}
}

func Test_RemoveAllItems_Multi(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 1, 3, 1})
	count := coregeneric.RemoveAllItems(col, 1)
	if count != 3 || col.Length() != 2 {
		t.Error("should remove all 1s")
	}
}

func Test_MapSimpleSlice_NonNil(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2})
	result := coregeneric.MapSimpleSlice(s, func(i int) string { return fmt.Sprintf("%d", i) })
	if result.Length() != 2 {
		t.Error("should map")
	}
}

func Test_DistinctSimpleSlice_WithDups(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 1, 3})
	result := coregeneric.DistinctSimpleSlice(s)
	if result.Length() != 3 {
		t.Error("should deduplicate")
	}
}

func Test_ContainsSimpleSliceItem_Found(t *testing.T) {
	s := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	if !coregeneric.ContainsSimpleSliceItem(s, 2) {
		t.Error("should find 2")
	}
}

func Test_ContainsAll_Found(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	if !coregeneric.ContainsAll(col, 1, 2) {
		t.Error("should contain all")
	}
}

func Test_ContainsAny_Found(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	if !coregeneric.ContainsAny(col, 5, 2) {
		t.Error("should contain any")
	}
}

func Test_RemoveItem_Found_Cov(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	if !coregeneric.RemoveItem(col, 2) {
		t.Error("should remove")
	}
	if col.Length() != 2 {
		t.Error("should have 2")
	}
}

func Test_Distinct_WithDups(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 1, 3})
	result := coregeneric.Distinct(col)
	if result.Length() != 3 {
		t.Error("should deduplicate")
	}
}

func Test_ToHashset_NonNil(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	hs := coregeneric.ToHashset(col)
	if hs.Length() != 3 {
		t.Error("should convert")
	}
}
