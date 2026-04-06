package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection — uncovered branches ──

func Test_Cov3_Collection_LengthLock(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	actual := args.Map{"len": col.LengthLock()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LengthLock", actual)
}

func Test_Cov3_Collection_IsEmptyLock(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	actual := args.Map{"empty": col.IsEmptyLock()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEmptyLock", actual)
}

func Test_Cov3_Collection_HasItems(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	empty := coregeneric.EmptyCollection[int]()
	actual := args.Map{"has": col.HasItems(), "empty": empty.HasItems()}
	expected := args.Map{"has": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasItems", actual)
}

func Test_Cov3_Collection_AddLock(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddLock(42)
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 42}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddLock", actual)
}

func Test_Cov3_Collection_AddsLock(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddsLock(1, 2, 3)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsLock", actual)
}

func Test_Cov3_Collection_AddIf_True(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(true, 42)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddIf true", actual)
}

func Test_Cov3_Collection_AddIfMany_True(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(true, 1, 2)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddIfMany true", actual)
}

func Test_Cov3_Collection_AddCollection(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	other := coregeneric.CollectionFrom([]int{2, 3})
	emptyOther := coregeneric.EmptyCollection[int]()
	col.AddCollection(emptyOther)
	col.AddCollection(other)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollection", actual)
}

func Test_Cov3_Collection_AddCollections(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	col.AddCollections(
		coregeneric.CollectionFrom([]int{1}),
		coregeneric.EmptyCollection[int](),
		coregeneric.CollectionFrom([]int{2}),
	)
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollections", actual)
}

func Test_Cov3_Collection_RemoveAt(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	ok := col.RemoveAt(1)
	fail := col.RemoveAt(-1)
	outOfRange := col.RemoveAt(100)
	actual := args.Map{"ok": ok, "fail": fail, "outOfRange": outOfRange, "len": col.Length()}
	expected := args.Map{"ok": true, "fail": false, "outOfRange": false, "len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt", actual)
}

func Test_Cov3_Collection_FirstOrDefault_Empty(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	actual := args.Map{"val": col.FirstOrDefault()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- FirstOrDefault empty", actual)
}

func Test_Cov3_Collection_LastOrDefault_Empty(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	actual := args.Map{"val": col.LastOrDefault()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- LastOrDefault empty", actual)
}

func Test_Cov3_Collection_SafeAt(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{10, 20})
	actual := args.Map{"valid": col.SafeAt(1), "invalid": col.SafeAt(99)}
	expected := args.Map{"valid": 20, "invalid": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SafeAt", actual)
}

func Test_Cov3_Collection_Skip_Take(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	actual := args.Map{
		"skipLen":    len(col.Skip(2)),
		"skipAll":    len(col.Skip(100)),
		"takeLen":    len(col.Take(2)),
		"takeAll":    len(col.Take(100)),
	}
	expected := args.Map{"skipLen": 2, "skipAll": 0, "takeLen": 2, "takeAll": 4}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Skip/Take", actual)
}

func Test_Cov3_Collection_ForEachBreak(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	col.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ForEachBreak", actual)
}

func Test_Cov3_Collection_Filter(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	filtered := col.Filter(func(v int) bool { return v > 2 })
	actual := args.Map{"len": filtered.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Filter", actual)
}

func Test_Cov3_Collection_Clone_Empty(t *testing.T) {
	col := coregeneric.EmptyCollection[int]()
	cloned := col.Clone()
	actual := args.Map{"empty": cloned.IsEmpty()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- Clone empty", actual)
}

func Test_Cov3_Collection_SortFunc(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{3, 1, 2})
	col.SortFunc(func(a, b int) bool { return a < b })
	actual := args.Map{"first": col.First(), "last": col.Last()}
	expected := args.Map{"first": 1, "last": 3}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SortFunc", actual)
}

func Test_Cov3_Collection_String(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	actual := args.Map{"notEmpty": col.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- String", actual)
}

func Test_Cov3_Collection_Count(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	actual := args.Map{"count": col.Count()}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Count", actual)
}

func Test_Cov3_Collection_ItemsPtr(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1})
	actual := args.Map{"notNil": col.ItemsPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ItemsPtr", actual)
}

func Test_Cov3_Collection_HasIndex(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2})
	actual := args.Map{"valid": col.HasIndex(1), "invalid": col.HasIndex(5), "neg": col.HasIndex(-1)}
	expected := args.Map{"valid": true, "invalid": false, "neg": false}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasIndex", actual)
}

// ── Hashset — uncovered branches ──

func Test_Cov3_Hashset_LengthLock(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	actual := args.Map{"len": hs.LengthLock()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- LengthLock", actual)
}

func Test_Cov3_Hashset_IsEmptyLock(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	actual := args.Map{"empty": hs.IsEmptyLock()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEmptyLock", actual)
}

func Test_Cov3_Hashset_HasItems(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	actual := args.Map{"has": hs.HasItems()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasItems", actual)
}

func Test_Cov3_Hashset_AddLock(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")
	actual := args.Map{"has": hs.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddLock", actual)
}

func Test_Cov3_Hashset_AddSliceLock(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSliceLock([]string{"a", "b"})
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddSliceLock", actual)
}

func Test_Cov3_Hashset_ContainsLock(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	actual := args.Map{"has": hs.ContainsLock("a"), "miss": hs.ContainsLock("b")}
	expected := args.Map{"has": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ContainsLock", actual)
}

func Test_Cov3_Hashset_HasAll_HasAny(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a", "b"})
	actual := args.Map{
		"all":      hs.HasAll("a", "b"),
		"notAll":   hs.HasAll("a", "c"),
		"any":      hs.HasAny("c", "a"),
		"notAny":   hs.HasAny("c", "d"),
	}
	expected := args.Map{"all": true, "notAll": false, "any": true, "notAny": false}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAll/HasAny", actual)
}

func Test_Cov3_Hashset_RemoveLock(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a", "b"})
	ok := hs.RemoveLock("a")
	miss := hs.RemoveLock("c")
	actual := args.Map{"ok": ok, "miss": miss, "len": hs.Length()}
	expected := args.Map{"ok": true, "miss": false, "len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- RemoveLock", actual)
}

func Test_Cov3_Hashset_ListPtr(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	actual := args.Map{"notNil": hs.ListPtr() != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtr", actual)
}

func Test_Cov3_Hashset_Map(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	actual := args.Map{"len": len(hs.Map())}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Map", actual)
}

func Test_Cov3_Hashset_Collection(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a", "b"})
	col := hs.Collection()
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Collection", actual)
}

func Test_Cov3_Hashset_IsEquals(t *testing.T) {
	hs1 := coregeneric.HashsetFrom([]string{"a", "b"})
	hs2 := coregeneric.HashsetFrom([]string{"a", "b"})
	hs3 := coregeneric.HashsetFrom([]string{"a", "c"})
	var nilHs *coregeneric.Hashset[string]
	actual := args.Map{
		"equal":     hs1.IsEquals(hs2),
		"notEqual":  hs1.IsEquals(hs3),
		"sameRef":   hs1.IsEquals(hs1),
		"nilBoth":   nilHs.IsEquals(nilHs),
		"nilLeft":   nilHs.IsEquals(hs1),
	}
	expected := args.Map{"equal": true, "notEqual": false, "sameRef": true, "nilBoth": true, "nilLeft": false}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals", actual)
}

func Test_Cov3_Hashset_String(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	actual := args.Map{"notEmpty": hs.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- String", actual)
}

func Test_Cov3_Hashset_AddHashsetItems_Nil(t *testing.T) {
	hs := coregeneric.HashsetFrom([]string{"a"})
	hs.AddHashsetItems(nil)
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddHashsetItems nil", actual)
}

func Test_Cov3_Hashset_AddIf(t *testing.T) {
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIf(true, "a")
	hs.AddIf(false, "b")
	actual := args.Map{"len": hs.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIf", actual)
}

// ── Hashmap — uncovered branches ──

func Test_Cov3_Hashmap_HasItems(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	actual := args.Map{"has": hm.HasItems()}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasItems", actual)
}

func Test_Cov3_Hashmap_IsEmptyLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	actual := args.Map{"empty": hm.IsEmptyLock()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- IsEmptyLock", actual)
}

func Test_Cov3_Hashmap_LengthLock(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	actual := args.Map{"len": hm.LengthLock()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- LengthLock", actual)
}

func Test_Cov3_Hashmap_SetLock(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.SetLock("a", 1)
	actual := args.Map{"has": hm.Has("a")}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- SetLock", actual)
}

func Test_Cov3_Hashmap_GetOrDefault(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	actual := args.Map{"found": hm.GetOrDefault("a", -1), "missing": hm.GetOrDefault("b", -1)}
	expected := args.Map{"found": 1, "missing": -1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetOrDefault", actual)
}

func Test_Cov3_Hashmap_GetLock(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	val, found := hm.GetLock("a")
	actual := args.Map{"val": val, "found": found}
	expected := args.Map{"val": 1, "found": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetLock", actual)
}

func Test_Cov3_Hashmap_ContainsLock(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	actual := args.Map{"has": hm.ContainsLock("a"), "miss": hm.ContainsLock("b")}
	expected := args.Map{"has": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ContainsLock", actual)
}

func Test_Cov3_Hashmap_IsKeyMissing(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	actual := args.Map{"miss": hm.IsKeyMissing("b"), "has": hm.IsKeyMissing("a")}
	expected := args.Map{"miss": true, "has": false}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsKeyMissing", actual)
}

func Test_Cov3_Hashmap_RemoveLock(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	ok := hm.RemoveLock("a")
	miss := hm.RemoveLock("b")
	actual := args.Map{"ok": ok, "miss": miss}
	expected := args.Map{"ok": true, "miss": false}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- RemoveLock", actual)
}

func Test_Cov3_Hashmap_AddOrUpdateMap(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.AddOrUpdateMap(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{})
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateMap", actual)
}

func Test_Cov3_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	hm := coregeneric.EmptyHashmap[string, int]()
	other := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(other)
	hm.AddOrUpdateHashmap(nil)
	actual := args.Map{"len": hm.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateHashmap", actual)
}

func Test_Cov3_Hashmap_Clone(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	cloned := hm.Clone()
	actual := args.Map{"len": cloned.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Clone", actual)
}

func Test_Cov3_Hashmap_IsEquals(t *testing.T) {
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm3 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	var nilHm *coregeneric.Hashmap[string, int]
	actual := args.Map{
		"equal":    hm1.IsEquals(hm2),
		"notEqual": hm1.IsEquals(hm3),
		"sameRef":  hm1.IsEquals(hm1),
		"nilBoth":  nilHm.IsEquals(nilHm),
		"nilLeft":  nilHm.IsEquals(hm1),
	}
	expected := args.Map{"equal": true, "notEqual": false, "sameRef": true, "nilBoth": true, "nilLeft": false}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEquals", actual)
}

func Test_Cov3_Hashmap_String(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	actual := args.Map{"notEmpty": hm.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- String", actual)
}

// ── LinkedList — uncovered branches ──

func Test_Cov3_LinkedList_LengthLock(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	actual := args.Map{"len": ll.LengthLock()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- LengthLock", actual)
}

func Test_Cov3_LinkedList_IsEmptyLock(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	actual := args.Map{"empty": ll.IsEmptyLock()}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- IsEmptyLock", actual)
}

func Test_Cov3_LinkedList_AddLock(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(42)
	actual := args.Map{"first": ll.First()}
	expected := args.Map{"first": 42}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddLock", actual)
}

func Test_Cov3_LinkedList_AddsIf(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(true, 1, 2)
	ll.AddsIf(false, 3)
	actual := args.Map{"len": ll.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddsIf", actual)
}

func Test_Cov3_LinkedList_AddFunc(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddFunc(func() int { return 99 })
	actual := args.Map{"first": ll.First()}
	expected := args.Map{"first": 99}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddFunc", actual)
}

func Test_Cov3_LinkedList_PushBackFrontPush(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	ll.PushBack(1)
	ll.PushFront(0)
	ll.Push(2)
	actual := args.Map{"first": ll.First(), "last": ll.Last(), "len": ll.Length()}
	expected := args.Map{"first": 0, "last": 2, "len": 3}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- PushBack/PushFront/Push", actual)
}

func Test_Cov3_LinkedList_AppendNode(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	node := &coregeneric.LinkedListNode[int]{Element: 42}
	ll.AppendNode(node)
	actual := args.Map{"first": ll.First(), "len": ll.Length()}
	expected := args.Map{"first": 42, "len": 1}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- AppendNode empty", actual)
}

func Test_Cov3_LinkedList_FirstOrDefault_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	actual := args.Map{"val": ll.FirstOrDefault()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- FirstOrDefault empty", actual)
}

func Test_Cov3_LinkedList_LastOrDefault_Empty(t *testing.T) {
	ll := coregeneric.EmptyLinkedList[int]()
	actual := args.Map{"val": ll.LastOrDefault()}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- LastOrDefault empty", actual)
}

func Test_Cov3_LinkedList_ForEachBreak(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- ForEachBreak", actual)
}

func Test_Cov3_LinkedList_IndexAt(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(1)
	nilNode := ll.IndexAt(-1)
	outNode := ll.IndexAt(100)
	actual := args.Map{"val": node.Element, "nil1": nilNode == nil, "nil2": outNode == nil}
	expected := args.Map{"val": 20, "nil1": true, "nil2": true}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- IndexAt", actual)
}

func Test_Cov3_LinkedList_String(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1})
	actual := args.Map{"notEmpty": ll.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- String", actual)
}

func Test_Cov3_LinkedList_Collection(t *testing.T) {
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	col := ll.Collection()
	actual := args.Map{"len": col.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Collection", actual)
}

// ── SimpleSlice — uncovered branches ──

func Test_Cov3_SimpleSlice_AddIf(t *testing.T) {
	ss := coregeneric.EmptySimpleSlice[int]()
	ss.AddIf(true, 1)
	ss.AddIf(false, 2)
	actual := args.Map{"len": ss.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddIf", actual)
}

func Test_Cov3_SimpleSlice_AddsIf(t *testing.T) {
	ss := coregeneric.EmptySimpleSlice[int]()
	ss.AddsIf(true, 1, 2)
	ss.AddsIf(false, 3)
	actual := args.Map{"len": ss.Length()}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddsIf", actual)
}

func Test_Cov3_SimpleSlice_AddFunc(t *testing.T) {
	ss := coregeneric.EmptySimpleSlice[int]()
	ss.AddFunc(func() int { return 99 })
	actual := args.Map{"first": ss.First()}
	expected := args.Map{"first": 99}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddFunc", actual)
}

func Test_Cov3_SimpleSlice_InsertAt(t *testing.T) {
	ss := coregeneric.SimpleSliceFrom([]int{1, 3})
	ss.InsertAt(1, 2)
	outOfRange := coregeneric.SimpleSliceFrom([]int{1})
	outOfRange.InsertAt(-1, 0)
	actual := args.Map{"len": ss.Length(), "outLen": outOfRange.Length()}
	expected := args.Map{"len": 3, "outLen": 1}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- InsertAt", actual)
}

func Test_Cov3_SimpleSlice_CountFunc(t *testing.T) {
	ss := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	count := ss.CountFunc(func(i int, item int) bool { return item > 1 })
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- CountFunc", actual)
}

func Test_Cov3_SimpleSlice_Clone(t *testing.T) {
	ss := coregeneric.SimpleSliceFrom([]int{1, 2})
	cloned := ss.Clone()
	emptyClone := coregeneric.EmptySimpleSlice[int]().Clone()
	actual := args.Map{"len": cloned.Length(), "emptyLen": emptyClone.Length()}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Clone", actual)
}

func Test_Cov3_SimpleSlice_String(t *testing.T) {
	ss := coregeneric.SimpleSliceFrom([]int{1})
	actual := args.Map{"notEmpty": ss.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- String", actual)
}

// ── orderedfuncs — uncovered branches ──

func Test_Cov3_SortCollectionDesc(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 3, 2})
	coregeneric.SortCollectionDesc(col)
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortCollectionDesc returns correct value -- with args", actual)
}

func Test_Cov3_SortCollectionDesc_Nil(t *testing.T) {
	result := coregeneric.SortCollectionDesc[int](nil)
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SortCollectionDesc returns nil -- nil", actual)
}

func Test_Cov3_MinMaxCollectionOrDefault(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{3, 1, 2})
	empty := coregeneric.EmptyCollection[int]()
	actual := args.Map{
		"min":         coregeneric.MinCollectionOrDefault(col, -1),
		"max":         coregeneric.MaxCollectionOrDefault(col, -1),
		"minEmpty":    coregeneric.MinCollectionOrDefault(empty, -1),
		"maxEmpty":    coregeneric.MaxCollectionOrDefault(empty, -1),
	}
	expected := args.Map{"min": 1, "max": 3, "minEmpty": -1, "maxEmpty": -1}
	expected.ShouldBeEqual(t, 0, "MinMax returns correct value -- CollectionOrDefault", actual)
}

func Test_Cov3_IsSortedCollection(t *testing.T) {
	sorted := coregeneric.CollectionFrom([]int{1, 2, 3})
	unsorted := coregeneric.CollectionFrom([]int{3, 1, 2})
	actual := args.Map{"sorted": coregeneric.IsSortedCollection(sorted), "unsorted": coregeneric.IsSortedCollection(unsorted), "nil": coregeneric.IsSortedCollection[int](nil)}
	expected := args.Map{"sorted": true, "unsorted": false, "nil": true}
	expected.ShouldBeEqual(t, 0, "IsSortedCollection returns correct value -- with args", actual)
}

func Test_Cov3_SumCollection(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{"sum": coregeneric.SumCollection(col), "nil": coregeneric.SumCollection[int](nil)}
	expected := args.Map{"sum": 6, "nil": 0}
	expected.ShouldBeEqual(t, 0, "SumCollection returns correct value -- with args", actual)
}

func Test_Cov3_ClampCollection(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{-1, 5, 15})
	coregeneric.ClampCollection(col, 0, 10)
	actual := args.Map{"first": col.First(), "last": col.Last(), "nil": coregeneric.ClampCollection[int](nil, 0, 10) == nil}
	expected := args.Map{"first": 0, "last": 10, "nil": true}
	expected.ShouldBeEqual(t, 0, "ClampCollection returns correct value -- with args", actual)
}

func Test_Cov3_SortSimpleSliceDesc(t *testing.T) {
	ss := coregeneric.SimpleSliceFrom([]int{1, 3, 2})
	coregeneric.SortSimpleSliceDesc(ss)
	actual := args.Map{"first": ss.First()}
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortSimpleSliceDesc returns correct value -- with args", actual)
}

func Test_Cov3_SumSimpleSlice(t *testing.T) {
	ss := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	actual := args.Map{"sum": coregeneric.SumSimpleSlice(ss)}
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "SumSimpleSlice returns correct value -- with args", actual)
}

// ── orderedfuncs Hashset ──

func Test_Cov3_SortedListDescHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	sorted := coregeneric.SortedListDescHashset(hs)
	actual := args.Map{"first": sorted[0]}
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortedListDescHashset returns correct value -- with args", actual)
}

func Test_Cov3_MinMaxHashsetOrDefault(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	empty := coregeneric.EmptyHashset[int]()
	actual := args.Map{
		"min":      coregeneric.MinHashsetOrDefault(hs, -1),
		"max":      coregeneric.MaxHashsetOrDefault(hs, -1),
		"minEmpty": coregeneric.MinHashsetOrDefault(empty, -1),
		"maxEmpty": coregeneric.MaxHashsetOrDefault(empty, -1),
	}
	expected := args.Map{"min": 1, "max": 3, "minEmpty": -1, "maxEmpty": -1}
	expected.ShouldBeEqual(t, 0, "MinMaxHashsetOrDefault returns correct value -- with args", actual)
}

func Test_Cov3_SortedCollectionHashset(t *testing.T) {
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	col := coregeneric.SortedCollectionHashset(hs)
	actual := args.Map{"first": col.First()}
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortedCollectionHashset returns correct value -- with args", actual)
}

// ── orderedfuncs Hashmap ──

func Test_Cov3_SortedKeysDescHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1, "b": 2})
	keys := coregeneric.SortedKeysDescHashmap(hm)
	actual := args.Map{"first": keys[0]}
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortedKeysDescHashmap returns correct value -- with args", actual)
}

func Test_Cov3_MinMaxKeyHashmapOrDefault(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1})
	empty := coregeneric.EmptyHashmap[string, int]()
	actual := args.Map{
		"min":      coregeneric.MinKeyHashmapOrDefault(hm, "z"),
		"max":      coregeneric.MaxKeyHashmapOrDefault(hm, "z"),
		"minEmpty": coregeneric.MinKeyHashmapOrDefault(empty, "z"),
		"maxEmpty": coregeneric.MaxKeyHashmapOrDefault(empty, "z"),
	}
	expected := args.Map{"min": "a", "max": "c", "minEmpty": "z", "maxEmpty": "z"}
	expected.ShouldBeEqual(t, 0, "MinMaxKeyHashmapOrDefault returns correct value -- with args", actual)
}

func Test_Cov3_SortedValuesHashmap(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})
	vals := coregeneric.SortedValuesHashmap(hm)
	actual := args.Map{"first": vals[0]}
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortedValuesHashmap returns non-empty -- with args", actual)
}

func Test_Cov3_MinMaxValueHashmapOrDefault(t *testing.T) {
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})
	empty := coregeneric.EmptyHashmap[string, int]()
	actual := args.Map{
		"min":      coregeneric.MinValueHashmapOrDefault(hm, -1),
		"max":      coregeneric.MaxValueHashmapOrDefault(hm, -1),
		"minEmpty": coregeneric.MinValueHashmapOrDefault(empty, -1),
		"maxEmpty": coregeneric.MaxValueHashmapOrDefault(empty, -1),
	}
	expected := args.Map{"min": 1, "max": 3, "minEmpty": -1, "maxEmpty": -1}
	expected.ShouldBeEqual(t, 0, "MinMaxValueHashmapOrDefault returns correct value -- with args", actual)
}

// ── comparablefuncs — uncovered branches ──

func Test_Cov3_ContainsAll(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{
		"all":    coregeneric.ContainsAll(col, 1, 2),
		"notAll": coregeneric.ContainsAll(col, 1, 5),
		"nil":    coregeneric.ContainsAll[int](nil, 1),
	}
	expected := args.Map{"all": true, "notAll": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "ContainsAll returns correct value -- with args", actual)
}

func Test_Cov3_ContainsAny(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	actual := args.Map{
		"any":    coregeneric.ContainsAny(col, 5, 2),
		"notAny": coregeneric.ContainsAny(col, 5, 6),
		"nil":    coregeneric.ContainsAny[int](nil, 1),
	}
	expected := args.Map{"any": true, "notAny": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "ContainsAny returns correct value -- with args", actual)
}

func Test_Cov3_RemoveItem(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	ok := coregeneric.RemoveItem(col, 2)
	miss := coregeneric.RemoveItem(col, 99)
	nilRm := coregeneric.RemoveItem[int](nil, 1)
	actual := args.Map{"ok": ok, "miss": miss, "nil": nilRm, "len": col.Length()}
	expected := args.Map{"ok": true, "miss": false, "nil": false, "len": 2}
	expected.ShouldBeEqual(t, 0, "RemoveItem returns correct value -- with args", actual)
}

func Test_Cov3_RemoveAllItems(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 1, 3})
	removed := coregeneric.RemoveAllItems(col, 1)
	nilRm := coregeneric.RemoveAllItems[int](nil, 1)
	actual := args.Map{"removed": removed, "len": col.Length(), "nil": nilRm}
	expected := args.Map{"removed": 2, "len": 2, "nil": 0}
	expected.ShouldBeEqual(t, 0, "RemoveAllItems returns correct value -- with args", actual)
}

func Test_Cov3_ToHashset(t *testing.T) {
	col := coregeneric.CollectionFrom([]int{1, 2, 2})
	hs := coregeneric.ToHashset(col)
	nilHs := coregeneric.ToHashset[int](nil)
	actual := args.Map{"len": hs.Length(), "nilEmpty": nilHs.IsEmpty()}
	expected := args.Map{"len": 2, "nilEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToHashset returns correct value -- with args", actual)
}

// ── numericfuncs — uncovered branches ──

func Test_Cov3_CompareNumeric(t *testing.T) {
	actual := args.Map{
		"equal":   coregeneric.CompareNumeric(5, 5),
		"greater": coregeneric.CompareNumeric(7, 5),
		"less":    coregeneric.CompareNumeric(3, 5),
	}
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"greater": corecomparator.LeftGreater,
		"less":    corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "CompareNumeric returns correct value -- with args", actual)
}

func Test_Cov3_Clamp(t *testing.T) {
	actual := args.Map{
		"below":  coregeneric.Clamp(-1, 0, 10),
		"above":  coregeneric.Clamp(15, 0, 10),
		"inside": coregeneric.Clamp(5, 0, 10),
	}
	expected := args.Map{"below": 0, "above": 10, "inside": 5}
	expected.ShouldBeEqual(t, 0, "Clamp returns correct value -- with args", actual)
}

func Test_Cov3_ClampMinMax(t *testing.T) {
	actual := args.Map{
		"clampMin": coregeneric.ClampMin(-1, 0),
		"clampMax": coregeneric.ClampMax(15, 10),
		"okMin":    coregeneric.ClampMin(5, 0),
		"okMax":    coregeneric.ClampMax(5, 10),
	}
	expected := args.Map{"clampMin": 0, "clampMax": 10, "okMin": 5, "okMax": 5}
	expected.ShouldBeEqual(t, 0, "ClampMin/ClampMax returns correct value -- with args", actual)
}

func Test_Cov3_Abs_AbsDiff(t *testing.T) {
	actual := args.Map{
		"absNeg":  coregeneric.Abs(-5),
		"absPos":  coregeneric.Abs(5),
		"diff":    coregeneric.AbsDiff(3, 7),
		"diffRev": coregeneric.AbsDiff(7, 3),
	}
	expected := args.Map{"absNeg": 5, "absPos": 5, "diff": 4, "diffRev": 4}
	expected.ShouldBeEqual(t, 0, "Abs/AbsDiff returns correct value -- with args", actual)
}

func Test_Cov3_MinMaxOfSlice(t *testing.T) {
	actual := args.Map{
		"min": coregeneric.MinOfSlice([]int{3, 1, 2}),
		"max": coregeneric.MaxOfSlice([]int{3, 1, 2}),
	}
	expected := args.Map{"min": 1, "max": 3}
	expected.ShouldBeEqual(t, 0, "MinOfSlice/MaxOfSlice returns correct value -- with args", actual)
}

func Test_Cov3_IsZero_IsPositive_IsNegative(t *testing.T) {
	actual := args.Map{
		"zero":     coregeneric.IsZero(0),
		"notZero":  coregeneric.IsZero(1),
		"positive": coregeneric.IsPositive(1),
		"notPos":   coregeneric.IsPositive(0),
		"negative": coregeneric.IsNegative(-1),
		"notNeg":   coregeneric.IsNegative(1),
	}
	expected := args.Map{"zero": true, "notZero": false, "positive": true, "notPos": false, "negative": true, "notNeg": false}
	expected.ShouldBeEqual(t, 0, "IsZero/IsPositive/IsNegative returns correct value -- with args", actual)
}

func Test_Cov3_SafeDiv(t *testing.T) {
	actual := args.Map{
		"normal": coregeneric.SafeDiv(10, 3),
		"zero":   coregeneric.SafeDiv(10, 0),
	}
	expected := args.Map{"normal": 3, "zero": 0}
	expected.ShouldBeEqual(t, 0, "SafeDiv returns correct value -- with args", actual)
}

// ── Relational predicates ──

func Test_Cov3_Relational(t *testing.T) {
	actual := args.Map{
		"less":     coregeneric.IsLess(3, 5),
		"lessEq":   coregeneric.IsLessOrEqual(5, 5),
		"greater":  coregeneric.IsGreater(7, 5),
		"greaterEq": coregeneric.IsGreaterOrEqual(5, 5),
		"inRange":  coregeneric.InRange(5, 1, 10),
	}
	expected := args.Map{"less": true, "lessEq": true, "greater": true, "greaterEq": true, "inRange": true}
	expected.ShouldBeEqual(t, 0, "Relational returns correct value -- predicates", actual)
}
