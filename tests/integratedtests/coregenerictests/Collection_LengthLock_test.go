package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection — uncovered branches ──

func Test_Collection_LengthLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"len": col.LengthLock()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- LengthLock", actual)
}

func Test_Collection_IsEmptyLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"empty": col.IsEmptyLock()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- IsEmptyLock", actual)
}

func Test_Collection_HasItems(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})
	empty := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{
		"has": col.HasItems(),
		"empty": empty.HasItems(),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"empty": false,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasItems", actual)
}

func Test_Collection_AddLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddLock(42)

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 42}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddLock", actual)
}

func Test_Collection_AddsLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddsLock(1, 2, 3)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddsLock", actual)
}

func Test_Collection_AddIf_True_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddIf(true, 42)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddIf true", actual)
}

func Test_Collection_AddIfMany_True_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(true, 1, 2)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns non-empty -- AddIfMany true", actual)
}

func Test_Collection_AddCollection(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})
	other := coregeneric.CollectionFrom([]int{2, 3})
	emptyOther := coregeneric.EmptyCollection[int]()
	col.AddCollection(emptyOther)
	col.AddCollection(other)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollection", actual)
}

func Test_Collection_AddCollections(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddCollections(
		coregeneric.CollectionFrom([]int{1}),
		coregeneric.EmptyCollection[int](),
		coregeneric.CollectionFrom([]int{2}),
	)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- AddCollections", actual)
}

func Test_Collection_RemoveAt(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	ok := col.RemoveAt(1)
	fail := col.RemoveAt(-1)
	outOfRange := col.RemoveAt(100)

	// Act
	actual := args.Map{
		"ok": ok,
		"fail": fail,
		"outOfRange": outOfRange,
		"len": col.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"fail": false,
		"outOfRange": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- RemoveAt", actual)
}

func Test_Collection_FirstOrDefault_Empty_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"val": col.FirstOrDefault()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- FirstOrDefault empty", actual)
}

func Test_Collection_LastOrDefault_Empty_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"val": col.LastOrDefault()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- LastOrDefault empty", actual)
}

func Test_Collection_SafeAt(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{10, 20})

	// Act
	actual := args.Map{
		"valid": col.SafeAt(1),
		"invalid": col.SafeAt(99),
	}

	// Assert
	expected := args.Map{
		"valid": 20,
		"invalid": 0,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SafeAt", actual)
}

func Test_Collection_Skip_Take(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4})

	// Act
	actual := args.Map{
		"skipLen":    len(col.Skip(2)),
		"skipAll":    len(col.Skip(100)),
		"takeLen":    len(col.Take(2)),
		"takeAll":    len(col.Take(100)),
	}

	// Assert
	expected := args.Map{
		"skipLen": 2,
		"skipAll": 0,
		"takeLen": 2,
		"takeAll": 4,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Skip/Take", actual)
}

func Test_Collection_ForEachBreak_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	count := 0
	col.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ForEachBreak", actual)
}

func Test_Collection_Filter(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4})
	filtered := col.Filter(func(v int) bool { return v > 2 })

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Filter", actual)
}

func Test_Collection_Clone_Empty_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	cloned := col.Clone()

	// Act
	actual := args.Map{"empty": cloned.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- Clone empty", actual)
}

func Test_Collection_SortFunc_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{3, 1, 2})
	col.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"first": col.First(),
		"last": col.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- SortFunc", actual)
}

func Test_Collection_String_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"notEmpty": col.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- String", actual)
}

func Test_Collection_Count(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"count": col.Count()}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Count", actual)
}

func Test_Collection_ItemsPtr_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"notNil": col.ItemsPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- ItemsPtr", actual)
}

func Test_Collection_HasIndex(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{
		"valid": col.HasIndex(1),
		"invalid": col.HasIndex(5),
		"neg": col.HasIndex(-1),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
		"neg": false,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasIndex", actual)
}

// ── Hashset — uncovered branches ──

func Test_Hashset_LengthLock(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"len": hs.LengthLock()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- LengthLock", actual)
}

func Test_Hashset_IsEmptyLock(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()

	// Act
	actual := args.Map{"empty": hs.IsEmptyLock()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEmptyLock", actual)
}

func Test_Hashset_HasItems(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"has": hs.HasItems()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasItems", actual)
}

func Test_Hashset_AddLock(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddLock("a")

	// Act
	actual := args.Map{"has": hs.Has("a")}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddLock", actual)
}

func Test_Hashset_AddSliceLock(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddSliceLock([]string{"a", "b"})

	// Act
	actual := args.Map{"len": hs.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddSliceLock", actual)
}

func Test_Hashset_ContainsLock(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{
		"has": hs.ContainsLock("a"),
		"miss": hs.ContainsLock("b"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ContainsLock", actual)
}

func Test_Hashset_HasAll_HasAny(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"all":      hs.HasAll("a", "b"),
		"notAll":   hs.HasAll("a", "c"),
		"any":      hs.HasAny("c", "a"),
		"notAny":   hs.HasAny("c", "d"),
	}

	// Assert
	expected := args.Map{
		"all": true,
		"notAll": false,
		"any": true,
		"notAny": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAll/HasAny", actual)
}

func Test_Hashset_RemoveLock(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a", "b"})
	ok := hs.RemoveLock("a")
	miss := hs.RemoveLock("c")

	// Act
	actual := args.Map{
		"ok": ok,
		"miss": miss,
		"len": hs.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"miss": false,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- RemoveLock", actual)
}

func Test_Hashset_ListPtr(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"notNil": hs.ListPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtr", actual)
}

func Test_Hashset_Map(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"len": len(hs.Map())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Map", actual)
}

func Test_Hashset_Collection(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a", "b"})
	col := hs.Collection()

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Collection", actual)
}

func Test_Hashset_IsEquals(t *testing.T) {
	// Arrange
	hs1 := coregeneric.HashsetFrom([]string{"a", "b"})
	hs2 := coregeneric.HashsetFrom([]string{"a", "b"})
	hs3 := coregeneric.HashsetFrom([]string{"a", "c"})
	var nilHs *coregeneric.Hashset[string]

	// Act
	actual := args.Map{
		"equal":     hs1.IsEquals(hs2),
		"notEqual":  hs1.IsEquals(hs3),
		"sameRef":   hs1.IsEquals(hs1),
		"nilBoth":   nilHs.IsEquals(nilHs),
		"nilLeft":   nilHs.IsEquals(hs1),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"sameRef": true,
		"nilBoth": true,
		"nilLeft": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals", actual)
}

func Test_Hashset_String(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})

	// Act
	actual := args.Map{"notEmpty": hs.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- String", actual)
}

func Test_Hashset_AddHashsetItems_Nil(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a"})
	hs.AddHashsetItems(nil)

	// Act
	actual := args.Map{"len": hs.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddHashsetItems nil", actual)
}

func Test_Hashset_AddIf(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIf(true, "a")
	hs.AddIf(false, "b")

	// Act
	actual := args.Map{"len": hs.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIf", actual)
}

// ── Hashmap — uncovered branches ──

func Test_Hashmap_HasItems(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"has": hm.HasItems()}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- HasItems", actual)
}

func Test_Hashmap_IsEmptyLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{"empty": hm.IsEmptyLock()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns empty -- IsEmptyLock", actual)
}

func Test_Hashmap_LengthLock(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"len": hm.LengthLock()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- LengthLock", actual)
}

func Test_Hashmap_SetLock(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.SetLock("a", 1)

	// Act
	actual := args.Map{"has": hm.Has("a")}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- SetLock", actual)
}

func Test_Hashmap_GetOrDefault(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"found": hm.GetOrDefault("a", -1),
		"missing": hm.GetOrDefault("b", -1),
	}

	// Assert
	expected := args.Map{
		"found": 1,
		"missing": -1,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetOrDefault", actual)
}

func Test_Hashmap_GetLock(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	val, found := hm.GetLock("a")

	// Act
	actual := args.Map{
		"val": val,
		"found": found,
	}

	// Assert
	expected := args.Map{
		"val": 1,
		"found": true,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- GetLock", actual)
}

func Test_Hashmap_ContainsLock(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"has": hm.ContainsLock("a"),
		"miss": hm.ContainsLock("b"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ContainsLock", actual)
}

func Test_Hashmap_IsKeyMissing(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"miss": hm.IsKeyMissing("b"),
		"has": hm.IsKeyMissing("a"),
	}

	// Assert
	expected := args.Map{
		"miss": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsKeyMissing", actual)
}

func Test_Hashmap_RemoveLock(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	ok := hm.RemoveLock("a")
	miss := hm.RemoveLock("b")

	// Act
	actual := args.Map{
		"ok": ok,
		"miss": miss,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- RemoveLock", actual)
}

func Test_Hashmap_AddOrUpdateMap(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.AddOrUpdateMap(map[string]int{"a": 1})
	hm.AddOrUpdateMap(map[string]int{})

	// Act
	actual := args.Map{"len": hm.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateMap", actual)
}

func Test_Hashmap_AddOrUpdateHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	other := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm.AddOrUpdateHashmap(other)
	hm.AddOrUpdateHashmap(nil)

	// Act
	actual := args.Map{"len": hm.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- AddOrUpdateHashmap", actual)
}

func Test_Hashmap_Clone(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})
	cloned := hm.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Clone", actual)
}

func Test_Hashmap_IsEquals(t *testing.T) {
	// Arrange
	hm1 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm2 := coregeneric.HashmapFrom(map[string]int{"a": 1})
	hm3 := coregeneric.HashmapFrom(map[string]int{"b": 2})
	var nilHm *coregeneric.Hashmap[string, int]

	// Act
	actual := args.Map{
		"equal":    hm1.IsEquals(hm2),
		"notEqual": hm1.IsEquals(hm3),
		"sameRef":  hm1.IsEquals(hm1),
		"nilBoth":  nilHm.IsEquals(nilHm),
		"nilLeft":  nilHm.IsEquals(hm1),
	}

	// Assert
	expected := args.Map{
		"equal": true,
		"notEqual": false,
		"sameRef": true,
		"nilBoth": true,
		"nilLeft": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- IsEquals", actual)
}

func Test_Hashmap_String(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 1})

	// Act
	actual := args.Map{"notEmpty": hm.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- String", actual)
}

// ── LinkedList — uncovered branches ──

func Test_LinkedList_LengthLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})

	// Act
	actual := args.Map{"len": ll.LengthLock()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- LengthLock", actual)
}

func Test_LinkedList_IsEmptyLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{"empty": ll.IsEmptyLock()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- IsEmptyLock", actual)
}

func Test_LinkedList_AddLock_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddLock(42)

	// Act
	actual := args.Map{"first": ll.First()}

	// Assert
	expected := args.Map{"first": 42}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddLock", actual)
}

func Test_LinkedList_AddsIf_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddsIf(true, 1, 2)
	ll.AddsIf(false, 3)

	// Act
	actual := args.Map{"len": ll.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddsIf", actual)
}

func Test_LinkedList_AddFunc(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.AddFunc(func() int { return 99 })

	// Act
	actual := args.Map{"first": ll.First()}

	// Assert
	expected := args.Map{"first": 99}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- AddFunc", actual)
}

func Test_LinkedList_PushBackFrontPush(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	ll.PushBack(1)
	ll.PushFront(0)
	ll.Push(2)

	// Act
	actual := args.Map{
		"first": ll.First(),
		"last": ll.Last(),
		"len": ll.Length(),
	}

	// Assert
	expected := args.Map{
		"first": 0,
		"last": 2,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- PushBack/PushFront/Push", actual)
}

func Test_LinkedList_AppendNode(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()
	node := &coregeneric.LinkedListNode[int]{Element: 42}
	ll.AppendNode(node)

	// Act
	actual := args.Map{
		"first": ll.First(),
		"len": ll.Length(),
	}

	// Assert
	expected := args.Map{
		"first": 42,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- AppendNode empty", actual)
}

func Test_LinkedList_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{"val": ll.FirstOrDefault()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- FirstOrDefault empty", actual)
}

func Test_LinkedList_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	ll := coregeneric.EmptyLinkedList[int]()

	// Act
	actual := args.Map{"val": ll.LastOrDefault()}

	// Assert
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "LinkedList returns empty -- LastOrDefault empty", actual)
}

func Test_LinkedList_ForEachBreak_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2, 3})
	count := 0
	ll.ForEachBreak(func(i int, item int) bool {
		count++
		return i == 1
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- ForEachBreak", actual)
}

func Test_LinkedList_IndexAt_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{10, 20, 30})
	node := ll.IndexAt(1)
	nilNode := ll.IndexAt(-1)
	outNode := ll.IndexAt(100)

	// Act
	actual := args.Map{
		"val": node.Element,
		"nil1": nilNode == nil,
		"nil2": outNode == nil,
	}

	// Assert
	expected := args.Map{
		"val": 20,
		"nil1": true,
		"nil2": true,
	}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- IndexAt", actual)
}

func Test_LinkedList_String_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1})

	// Act
	actual := args.Map{"notEmpty": ll.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- String", actual)
}

func Test_LinkedList_Collection_FromCollectionLengthLock(t *testing.T) {
	// Arrange
	ll := coregeneric.LinkedListFrom([]int{1, 2})
	col := ll.Collection()

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "LinkedList returns correct value -- Collection", actual)
}

// ── SimpleSlice — uncovered branches ──

func Test_SimpleSlice_AddIf(t *testing.T) {
	// Arrange
	ss := coregeneric.EmptySimpleSlice[int]()
	ss.AddIf(true, 1)
	ss.AddIf(false, 2)

	// Act
	actual := args.Map{"len": ss.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddIf", actual)
}

func Test_SimpleSlice_AddsIf(t *testing.T) {
	// Arrange
	ss := coregeneric.EmptySimpleSlice[int]()
	ss.AddsIf(true, 1, 2)
	ss.AddsIf(false, 3)

	// Act
	actual := args.Map{"len": ss.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddsIf", actual)
}

func Test_SimpleSlice_AddFunc(t *testing.T) {
	// Arrange
	ss := coregeneric.EmptySimpleSlice[int]()
	ss.AddFunc(func() int { return 99 })

	// Act
	actual := args.Map{"first": ss.First()}

	// Assert
	expected := args.Map{"first": 99}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- AddFunc", actual)
}

func Test_SimpleSlice_InsertAt(t *testing.T) {
	// Arrange
	ss := coregeneric.SimpleSliceFrom([]int{1, 3})
	ss.InsertAt(1, 2)
	outOfRange := coregeneric.SimpleSliceFrom([]int{1})
	outOfRange.InsertAt(-1, 0)

	// Act
	actual := args.Map{
		"len": ss.Length(),
		"outLen": outOfRange.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"outLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- InsertAt", actual)
}

func Test_SimpleSlice_CountFunc(t *testing.T) {
	// Arrange
	ss := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	count := ss.CountFunc(func(i int, item int) bool { return item > 1 })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- CountFunc", actual)
}

func Test_SimpleSlice_Clone(t *testing.T) {
	// Arrange
	ss := coregeneric.SimpleSliceFrom([]int{1, 2})
	cloned := ss.Clone()
	emptyClone := coregeneric.EmptySimpleSlice[int]().Clone()

	// Act
	actual := args.Map{
		"len": cloned.Length(),
		"emptyLen": emptyClone.Length(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- Clone", actual)
}

func Test_SimpleSlice_String(t *testing.T) {
	// Arrange
	ss := coregeneric.SimpleSliceFrom([]int{1})

	// Act
	actual := args.Map{"notEmpty": ss.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SimpleSlice returns correct value -- String", actual)
}

// ── orderedfuncs — uncovered branches ──

func Test_SortCollectionDesc(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 3, 2})
	coregeneric.SortCollectionDesc(col)

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortCollectionDesc returns correct value -- with args", actual)
}

func Test_SortCollectionDesc_Nil(t *testing.T) {
	// Arrange
	result := coregeneric.SortCollectionDesc[int](nil)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SortCollectionDesc returns nil -- nil", actual)
}

func Test_MinMaxCollectionOrDefault(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{3, 1, 2})
	empty := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{
		"min":         coregeneric.MinCollectionOrDefault(col, -1),
		"max":         coregeneric.MaxCollectionOrDefault(col, -1),
		"minEmpty":    coregeneric.MinCollectionOrDefault(empty, -1),
		"maxEmpty":    coregeneric.MaxCollectionOrDefault(empty, -1),
	}

	// Assert
	expected := args.Map{
		"min": 1,
		"max": 3,
		"minEmpty": -1,
		"maxEmpty": -1,
	}
	expected.ShouldBeEqual(t, 0, "MinMax returns correct value -- CollectionOrDefault", actual)
}

func Test_IsSortedCollection(t *testing.T) {
	// Arrange
	sorted := coregeneric.CollectionFrom([]int{1, 2, 3})
	unsorted := coregeneric.CollectionFrom([]int{3, 1, 2})

	// Act
	actual := args.Map{
		"sorted": coregeneric.IsSortedCollection(sorted),
		"unsorted": coregeneric.IsSortedCollection(unsorted),
		"nil": coregeneric.IsSortedCollection[int](nil),
	}

	// Assert
	expected := args.Map{
		"sorted": true,
		"unsorted": false,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "IsSortedCollection returns correct value -- with args", actual)
}

func Test_SumCollection(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{
		"sum": coregeneric.SumCollection(col),
		"nil": coregeneric.SumCollection[int](nil),
	}

	// Assert
	expected := args.Map{
		"sum": 6,
		"nil": 0,
	}
	expected.ShouldBeEqual(t, 0, "SumCollection returns correct value -- with args", actual)
}

func Test_ClampCollection(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{-1, 5, 15})
	coregeneric.ClampCollection(col, 0, 10)

	// Act
	actual := args.Map{
		"first": col.First(),
		"last": col.Last(),
		"nil": coregeneric.ClampCollection[int](nil, 0, 10) == nil,
	}

	// Assert
	expected := args.Map{
		"first": 0,
		"last": 10,
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "ClampCollection returns correct value -- with args", actual)
}

func Test_SortSimpleSliceDesc(t *testing.T) {
	// Arrange
	ss := coregeneric.SimpleSliceFrom([]int{1, 3, 2})
	coregeneric.SortSimpleSliceDesc(ss)

	// Act
	actual := args.Map{"first": ss.First()}

	// Assert
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortSimpleSliceDesc returns correct value -- with args", actual)
}

func Test_SumSimpleSlice(t *testing.T) {
	// Arrange
	ss := coregeneric.SimpleSliceFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"sum": coregeneric.SumSimpleSlice(ss)}

	// Assert
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "SumSimpleSlice returns correct value -- with args", actual)
}

// ── orderedfuncs Hashset ──

func Test_SortedListDescHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	sorted := coregeneric.SortedListDescHashset(hs)

	// Act
	actual := args.Map{"first": sorted[0]}

	// Assert
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortedListDescHashset returns correct value -- with args", actual)
}

func Test_MinMaxHashsetOrDefault(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	empty := coregeneric.EmptyHashset[int]()

	// Act
	actual := args.Map{
		"min":      coregeneric.MinHashsetOrDefault(hs, -1),
		"max":      coregeneric.MaxHashsetOrDefault(hs, -1),
		"minEmpty": coregeneric.MinHashsetOrDefault(empty, -1),
		"maxEmpty": coregeneric.MaxHashsetOrDefault(empty, -1),
	}

	// Assert
	expected := args.Map{
		"min": 1,
		"max": 3,
		"minEmpty": -1,
		"maxEmpty": -1,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxHashsetOrDefault returns correct value -- with args", actual)
}

func Test_SortedCollectionHashset(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]int{3, 1, 2})
	col := coregeneric.SortedCollectionHashset(hs)

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortedCollectionHashset returns correct value -- with args", actual)
}

// ── orderedfuncs Hashmap ──

func Test_SortedKeysDescHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1, "b": 2})
	keys := coregeneric.SortedKeysDescHashmap(hm)

	// Act
	actual := args.Map{"first": keys[0]}

	// Assert
	expected := args.Map{"first": "c"}
	expected.ShouldBeEqual(t, 0, "SortedKeysDescHashmap returns correct value -- with args", actual)
}

func Test_MinMaxKeyHashmapOrDefault(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"c": 3, "a": 1})
	empty := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{
		"min":      coregeneric.MinKeyHashmapOrDefault(hm, "z"),
		"max":      coregeneric.MaxKeyHashmapOrDefault(hm, "z"),
		"minEmpty": coregeneric.MinKeyHashmapOrDefault(empty, "z"),
		"maxEmpty": coregeneric.MaxKeyHashmapOrDefault(empty, "z"),
	}

	// Assert
	expected := args.Map{
		"min": "a",
		"max": "c",
		"minEmpty": "z",
		"maxEmpty": "z",
	}
	expected.ShouldBeEqual(t, 0, "MinMaxKeyHashmapOrDefault returns correct value -- with args", actual)
}

func Test_SortedValuesHashmap(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})
	vals := coregeneric.SortedValuesHashmap(hm)

	// Act
	actual := args.Map{"first": vals[0]}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortedValuesHashmap returns non-empty -- with args", actual)
}

func Test_MinMaxValueHashmapOrDefault(t *testing.T) {
	// Arrange
	hm := coregeneric.HashmapFrom(map[string]int{"a": 3, "b": 1})
	empty := coregeneric.EmptyHashmap[string, int]()

	// Act
	actual := args.Map{
		"min":      coregeneric.MinValueHashmapOrDefault(hm, -1),
		"max":      coregeneric.MaxValueHashmapOrDefault(hm, -1),
		"minEmpty": coregeneric.MinValueHashmapOrDefault(empty, -1),
		"maxEmpty": coregeneric.MaxValueHashmapOrDefault(empty, -1),
	}

	// Assert
	expected := args.Map{
		"min": 1,
		"max": 3,
		"minEmpty": -1,
		"maxEmpty": -1,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxValueHashmapOrDefault returns correct value -- with args", actual)
}

// ── comparablefuncs — uncovered branches ──

func Test_ContainsAll(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{
		"all":    coregeneric.ContainsAll(col, 1, 2),
		"notAll": coregeneric.ContainsAll(col, 1, 5),
		"nil":    coregeneric.ContainsAll[int](nil, 1),
	}

	// Assert
	expected := args.Map{
		"all": true,
		"notAll": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "ContainsAll returns correct value -- with args", actual)
}

func Test_ContainsAny(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{
		"any":    coregeneric.ContainsAny(col, 5, 2),
		"notAny": coregeneric.ContainsAny(col, 5, 6),
		"nil":    coregeneric.ContainsAny[int](nil, 1),
	}

	// Assert
	expected := args.Map{
		"any": true,
		"notAny": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "ContainsAny returns correct value -- with args", actual)
}

func Test_RemoveItem(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	ok := coregeneric.RemoveItem(col, 2)
	miss := coregeneric.RemoveItem(col, 99)
	nilRm := coregeneric.RemoveItem[int](nil, 1)

	// Act
	actual := args.Map{
		"ok": ok,
		"miss": miss,
		"nil": nilRm,
		"len": col.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"miss": false,
		"nil": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "RemoveItem returns correct value -- with args", actual)
}

func Test_RemoveAllItems(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 1, 3})
	removed := coregeneric.RemoveAllItems(col, 1)
	nilRm := coregeneric.RemoveAllItems[int](nil, 1)

	// Act
	actual := args.Map{
		"removed": removed,
		"len": col.Length(),
		"nil": nilRm,
	}

	// Assert
	expected := args.Map{
		"removed": 2,
		"len": 2,
		"nil": 0,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAllItems returns correct value -- with args", actual)
}

func Test_ToHashset(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 2})
	hs := coregeneric.ToHashset(col)
	nilHs := coregeneric.ToHashset[int](nil)

	// Act
	actual := args.Map{
		"len": hs.Length(),
		"nilEmpty": nilHs.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"nilEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "ToHashset returns correct value -- with args", actual)
}

// ── numericfuncs — uncovered branches ──

func Test_CompareNumeric_FromCollectionLengthLock(t *testing.T) {
	// Act
	actual := args.Map{
		"equal":   coregeneric.CompareNumeric(5, 5),
		"greater": coregeneric.CompareNumeric(7, 5),
		"less":    coregeneric.CompareNumeric(3, 5),
	}

	// Assert
	expected := args.Map{
		"equal":   corecomparator.Equal,
		"greater": corecomparator.LeftGreater,
		"less":    corecomparator.LeftLess,
	}
	expected.ShouldBeEqual(t, 0, "CompareNumeric returns correct value -- with args", actual)
}

func Test_Clamp_FromCollectionLengthLock(t *testing.T) {
	// Act
	actual := args.Map{
		"below":  coregeneric.Clamp(-1, 0, 10),
		"above":  coregeneric.Clamp(15, 0, 10),
		"inside": coregeneric.Clamp(5, 0, 10),
	}

	// Assert
	expected := args.Map{
		"below": 0,
		"above": 10,
		"inside": 5,
	}
	expected.ShouldBeEqual(t, 0, "Clamp returns correct value -- with args", actual)
}

func Test_ClampMinMax(t *testing.T) {
	// Act
	actual := args.Map{
		"clampMin": coregeneric.ClampMin(-1, 0),
		"clampMax": coregeneric.ClampMax(15, 10),
		"okMin":    coregeneric.ClampMin(5, 0),
		"okMax":    coregeneric.ClampMax(5, 10),
	}

	// Assert
	expected := args.Map{
		"clampMin": 0,
		"clampMax": 10,
		"okMin": 5,
		"okMax": 5,
	}
	expected.ShouldBeEqual(t, 0, "ClampMin/ClampMax returns correct value -- with args", actual)
}

func Test_Abs_AbsDiff(t *testing.T) {
	// Act
	actual := args.Map{
		"absNeg":  coregeneric.Abs(-5),
		"absPos":  coregeneric.Abs(5),
		"diff":    coregeneric.AbsDiff(3, 7),
		"diffRev": coregeneric.AbsDiff(7, 3),
	}

	// Assert
	expected := args.Map{
		"absNeg": 5,
		"absPos": 5,
		"diff": 4,
		"diffRev": 4,
	}
	expected.ShouldBeEqual(t, 0, "Abs/AbsDiff returns correct value -- with args", actual)
}

func Test_MinMaxOfSlice(t *testing.T) {
	// Act
	actual := args.Map{
		"min": coregeneric.MinOfSlice([]int{3, 1, 2}),
		"max": coregeneric.MaxOfSlice([]int{3, 1, 2}),
	}

	// Assert
	expected := args.Map{
		"min": 1,
		"max": 3,
	}
	expected.ShouldBeEqual(t, 0, "MinOfSlice/MaxOfSlice returns correct value -- with args", actual)
}

func Test_IsZero_IsPositive_IsNegative(t *testing.T) {
	// Act
	actual := args.Map{
		"zero":     coregeneric.IsZero(0),
		"notZero":  coregeneric.IsZero(1),
		"positive": coregeneric.IsPositive(1),
		"notPos":   coregeneric.IsPositive(0),
		"negative": coregeneric.IsNegative(-1),
		"notNeg":   coregeneric.IsNegative(1),
	}

	// Assert
	expected := args.Map{
		"zero": true,
		"notZero": false,
		"positive": true,
		"notPos": false,
		"negative": true,
		"notNeg": false,
	}
	expected.ShouldBeEqual(t, 0, "IsZero/IsPositive/IsNegative returns correct value -- with args", actual)
}

func Test_SafeDiv_FromCollectionLengthLock(t *testing.T) {
	// Act
	actual := args.Map{
		"normal": coregeneric.SafeDiv(10, 3),
		"zero":   coregeneric.SafeDiv(10, 0),
	}

	// Assert
	expected := args.Map{
		"normal": 3,
		"zero": 0,
	}
	expected.ShouldBeEqual(t, 0, "SafeDiv returns correct value -- with args", actual)
}

// ── Relational predicates ──

func Test_Relational(t *testing.T) {
	// Act
	actual := args.Map{
		"less":     coregeneric.IsLess(3, 5),
		"lessEq":   coregeneric.IsLessOrEqual(5, 5),
		"greater":  coregeneric.IsGreater(7, 5),
		"greaterEq": coregeneric.IsGreaterOrEqual(5, 5),
		"inRange":  coregeneric.InRange(5, 1, 10),
	}

	// Assert
	expected := args.Map{
		"less": true,
		"lessEq": true,
		"greater": true,
		"greaterEq": true,
		"inRange": true,
	}
	expected.ShouldBeEqual(t, 0, "Relational returns correct value -- predicates", actual)
}
