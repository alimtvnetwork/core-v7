package coredynamictests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════
// Collection[T] — core methods
// ═══════════════════════════════════════════

func Test_Collection_NewAndEmpty(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](10)
	e := coredynamic.EmptyCollection[string]()

	// Act
	actual := args.Map{
		"cLen": c.Length(),
		"cEmpty": c.IsEmpty(),
		"eLen": e.Length(),
		"eEmpty": e.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"cLen": 0,
		"cEmpty": true,
		"eLen": 0,
		"eEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "New/Empty returns empty -- with args", actual)
}

func Test_Collection_From_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom[string](nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CollectionFrom returns nil -- nil", actual)
}

func Test_Collection_From(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": c.Length(),
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "CollectionFrom returns correct value -- with args", actual)
}

func Test_Collection_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionClone([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": c.Length(),
		"first": c.First(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "CollectionClone returns correct value -- with args", actual)
}

func Test_Collection_At_SafeAt(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"at": c.At(1),
		"safe": c.SafeAt(1),
		"oob": c.SafeAt(99),
	}

	// Assert
	expected := args.Map{
		"at": "b",
		"safe": "b",
		"oob": "",
	}
	expected.ShouldBeEqual(t, 0, "At/SafeAt returns correct value -- with args", actual)
}

func Test_Collection_FirstOrDefault(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})
	e := coredynamic.EmptyCollection[string]()
	f, fOK := c.FirstOrDefault()
	l, lOK := c.LastOrDefault()
	ef, efOK := e.FirstOrDefault()
	el, elOK := e.LastOrDefault()

	// Act
	actual := args.Map{
		"f": *f,
		"fOK": fOK,
		"l": *l,
		"lOK": lOK,
		"efNil": ef == nil,
		"efOK": efOK,
		"elNil": el == nil,
		"elOK": elOK,
	}

	// Assert
	expected := args.Map{
		"f": "a",
		"fOK": true,
		"l": "a",
		"lOK": true,
		"efNil": true,
		"efOK": false,
		"elNil": true,
		"elOK": false,
	}
	expected.ShouldBeEqual(t, 0, "FirstOrDefault returns correct value -- with args", actual)
}

func Test_Collection_Items_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]

	// Act
	actual := args.Map{"len": len(c.Items())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Items returns nil -- nil", actual)
}

func Test_Collection_Length_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]

	// Act
	actual := args.Map{
		"len": c.Length(),
		"count": c.Count(),
		"empty": c.IsEmpty(),
		"has": c.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"count": 0,
		"empty": true,
		"has": false,
	}
	expected.ShouldBeEqual(t, 0, "Length returns nil -- nil", actual)
}

func Test_Collection_HasIndex(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"v0": c.HasIndex(0),
		"v1": c.HasIndex(1),
		"v2": c.HasIndex(2),
		"neg": c.HasIndex(-1),
		"last": c.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"v0": true,
		"v1": true,
		"v2": false,
		"neg": false,
		"last": 1,
	}
	expected.ShouldBeEqual(t, 0, "HasIndex returns correct value -- with args", actual)
}

func Test_Collection_SkipTakeLimitSlice(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c", "d"})

	// Act
	actual := args.Map{
		"skipLen": len(c.Skip(2)), "takeLen": len(c.Take(2)), "limitLen": len(c.Limit(2)),
		"skipCLen": c.SkipCollection(2).Length(), "takeCLen": c.TakeCollection(2).Length(),
		"limitCLen": c.LimitCollection(2).Length(), "safeLimitLen": c.SafeLimitCollection(2).Length(),
	}

	// Assert
	expected := args.Map{
		"skipLen": 2, "takeLen": 2, "limitLen": 2,
		"skipCLen": 2, "takeCLen": 2, "limitCLen": 2, "safeLimitLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "Skip/Take/Limit returns correct value -- with args", actual)
}

func Test_Collection_Add_AddMany_AddNonNil(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.Add("a").AddMany("b", "c")
	s := "d"
	c.AddNonNil(&s)
	c.AddNonNil(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "Add/AddMany/AddNonNil returns nil -- with args", actual)
}

func Test_Collection_RemoveAt(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})
	ok := c.RemoveAt(1)
	bad := c.RemoveAt(99)

	// Act
	actual := args.Map{
		"ok": ok,
		"bad": bad,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"bad": false,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAt returns correct value -- with args", actual)
}

func Test_Collection_ClearDispose(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	c.Clear()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clear returns correct value -- with args", actual)
	c2 := coredynamic.CollectionFrom([]string{"a"})
	c2.Dispose()
	actual2 := args.Map{"items": len(c2.Items())}
	expected2 := args.Map{"items": 0}
	expected2.ShouldBeEqual(t, 0, "Dispose returns correct value -- with args", actual2)
}

func Test_Collection_Loop(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})
	count := 0
	c.Loop(func(i int, item string) bool { count++; return false })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Loop returns correct value -- with args", actual)
}

func Test_Collection_Loop_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	count := 0
	c.Loop(func(i int, item string) bool { count++; return false })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "Loop returns empty -- empty", actual)
}

func Test_Collection_LoopAsync(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	var mu sync.Mutex
	sum := 0
	c.LoopAsync(func(i int, item int) {
		mu.Lock()
		sum += item
		mu.Unlock()
	})

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "LoopAsync returns correct value -- with args", actual)
}

func Test_Collection_LoopAsync_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.LoopAsync(func(i int, item int) {})

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "LoopAsync returns empty -- empty", actual)
}

func Test_Collection_Filter(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	filtered := c.Filter(func(i int) bool { return i%2 == 0 })

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Filter returns correct value -- with args", actual)
}

func Test_Collection_Filter_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	filtered := c.Filter(func(i int) bool { return true })

	// Act
	actual := args.Map{"len": filtered.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Filter returns empty -- empty", actual)
}

func Test_Collection_Paging(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	pages := c.GetPagesSize(2)
	zero := c.GetPagesSize(0)

	// Act
	actual := args.Map{
		"pages": pages,
		"zero": zero,
	}

	// Assert
	expected := args.Map{
		"pages": 3,
		"zero": 0,
	}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns correct value -- with args", actual)
}

func Test_Collection_GetPagedCollection(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	paged := c.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(paged)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- with args", actual)
}

func Test_Collection_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	paged := c.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(paged)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns correct value -- small", actual)
}

func Test_Collection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4, 5})
	page := c.GetSinglePageCollection(2, 1)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection returns correct value -- with args", actual)
}

func Test_Collection_Json(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	js, err := c.JsonString()
	jm := c.JsonStringMust()
	b, merr := c.MarshalJSON()

	// Act
	actual := args.Map{
		"jsNE": js != "",
		"noErr": err == nil,
		"jmNE": jm != "",
		"bLen": len(b) > 0,
		"merrNil": merr == nil,
	}

	// Assert
	expected := args.Map{
		"jsNE": true,
		"noErr": true,
		"jmNE": true,
		"bLen": true,
		"merrNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- Json", actual)
}

func Test_Collection_UnmarshalJSON(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	err := c.UnmarshalJSON([]byte(`["x","y"]`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns correct value -- with args", actual)
}

func Test_Collection_Strings(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	strs := c.Strings()
	str := c.String()

	// Act
	actual := args.Map{
		"len": len(strs),
		"strNE": str != "",
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"strNE": true,
	}
	expected.ShouldBeEqual(t, 0, "Strings returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// CollectionMethods — AddIf, AddCollection, ConcatNew, Clone, Capacity, Reverse, InsertAt, IndexOfFunc, SprintItems
// ═══════════════════════════════════════════

func Test_CollectionMethods_AddIf(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.AddIf(true, "a").AddIf(false, "b")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddIf returns correct value -- with args", actual)
}

func Test_CollectionMethods_AddManyIf(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.AddManyIf(true, "a", "b").AddManyIf(false, "c")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddManyIf returns correct value -- with args", actual)
}

func Test_CollectionMethods_AddCollection(t *testing.T) {
	// Arrange
	c1 := coredynamic.CollectionFrom([]string{"a"})
	c2 := coredynamic.CollectionFrom([]string{"b", "c"})
	c1.AddCollection(c2)
	c1.AddCollection(nil)

	// Act
	actual := args.Map{"len": c1.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AddCollection returns correct value -- with args", actual)
}

func Test_CollectionMethods_AddCollections(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c1 := coredynamic.CollectionFrom([]string{"a"})
	c.AddCollections(c1, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddCollections returns correct value -- with args", actual)
}

func Test_CollectionMethods_ConcatNew(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})
	c2 := c.ConcatNew("b", "c")

	// Act
	actual := args.Map{
		"origLen": c.Length(),
		"newLen": c2.Length(),
	}

	// Assert
	expected := args.Map{
		"origLen": 1,
		"newLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns correct value -- with args", actual)
}

func Test_CollectionMethods_Clone_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]
	cloned := c.Clone()

	// Act
	actual := args.Map{"len": cloned.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns nil -- nil", actual)
}

func Test_CollectionMethods_Capacity(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](10)
	var nilC *coredynamic.Collection[string]

	// Act
	actual := args.Map{
		"cap": c.Capacity() >= 10,
		"nilCap": nilC.Capacity(),
	}

	// Assert
	expected := args.Map{
		"cap": true,
		"nilCap": 0,
	}
	expected.ShouldBeEqual(t, 0, "Capacity returns correct value -- with args", actual)
}

func Test_CollectionMethods_AddCapacity(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.AddCapacity(10)
	c.AddCapacity(0) // no-op

	// Act
	actual := args.Map{"capGTE15": c.Capacity() >= 15}

	// Assert
	expected := args.Map{"capGTE15": true}
	expected.ShouldBeEqual(t, 0, "AddCapacity returns correct value -- with args", actual)
}

func Test_CollectionMethods_Resize(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.Resize(20)
	c.Resize(5) // no-op

	// Act
	actual := args.Map{"capGTE20": c.Capacity() >= 20}

	// Assert
	expected := args.Map{"capGTE20": true}
	expected.ShouldBeEqual(t, 0, "Resize returns correct value -- with args", actual)
}

func Test_CollectionMethods_Reverse(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})
	c.Reverse()

	// Act
	actual := args.Map{
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"last": "a",
	}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- with args", actual)
}

func Test_CollectionMethods_Reverse_Single(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})
	c.Reverse()

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- single", actual)
}

func Test_CollectionMethods_InsertAt(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "c"})
	c.InsertAt(1, "b")

	// Act
	actual := args.Map{
		"len": c.Length(),
		"at1": c.At(1),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"at1": "b",
	}
	expected.ShouldBeEqual(t, 0, "InsertAt returns correct value -- with args", actual)
}

func Test_CollectionMethods_InsertAt_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})
	c.InsertAt(0)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "InsertAt returns empty -- empty", actual)
}

func Test_CollectionMethods_IndexOfFunc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})
	idx := c.IndexOfFunc(func(s string) bool { return s == "b" })
	missing := c.IndexOfFunc(func(s string) bool { return s == "z" })
	has := c.ContainsFunc(func(s string) bool { return s == "b" })

	// Act
	actual := args.Map{
		"idx": idx,
		"missing": missing,
		"has": has,
	}

	// Assert
	expected := args.Map{
		"idx": 1,
		"missing": -1,
		"has": true,
	}
	expected.ShouldBeEqual(t, 0, "IndexOfFunc returns correct value -- with args", actual)
}

func Test_CollectionMethods_SprintItems(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	items := c.SprintItems("%d")

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SprintItems returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// CollectionSearch — Contains, IndexOf, Has, HasAll, LastIndexOf, Count, Lock variants
// ═══════════════════════════════════════════

func Test_CollectionSearch_Contains(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{
		"has":   coredynamic.Contains(c, "b"),
		"miss":  coredynamic.Contains(c, "z"),
		"idx":   coredynamic.IndexOf(c, "b"),
		"alias": coredynamic.Has(c, "b"),
		"all":   coredynamic.HasAll(c, "a", "b"),
		"allM":  coredynamic.HasAll(c, "a", "z"),
		"last":  coredynamic.LastIndexOf(c, "b"),
		"lastM": coredynamic.LastIndexOf(c, "z"),
		"count": coredynamic.Count(c, "b"),
	}

	// Assert
	expected := args.Map{
		"has": true, "miss": false, "idx": 1, "alias": true,
		"all": true, "allM": false, "last": 1, "lastM": -1, "count": 1,
	}
	expected.ShouldBeEqual(t, 0, "Search returns correct value -- with args", actual)
}

func Test_CollectionSearch_HasAll_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()

	// Act
	actual := args.Map{"v": coredynamic.HasAll(c, "a")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "HasAll returns empty -- empty", actual)
}

func Test_CollectionSearch_Lock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"contains": coredynamic.ContainsLock(c, "a"),
		"indexOf":  coredynamic.IndexOfLock(c, "a"),
	}

	// Assert
	expected := args.Map{
		"contains": true,
		"indexOf": 0,
	}
	expected.ShouldBeEqual(t, 0, "Search returns correct value -- Lock", actual)
}

// ═══════════════════════════════════════════
// CollectionSort
// ═══════════════════════════════════════════

func Test_CollectionSort_SortFunc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortFunc returns correct value -- with args", actual)
}

func Test_CollectionSort_SortFunc_Single(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFunc returns correct value -- single", actual)
}

func Test_CollectionSort_SortFuncLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	c.SortFuncLock(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFuncLock returns correct value -- with args", actual)
}

func Test_CollectionSort_SortedFunc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	sorted := c.SortedFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"origFirst": c.First(),
		"sortedFirst": sorted.First(),
	}

	// Assert
	expected := args.Map{
		"origFirst": 3,
		"sortedFirst": 1,
	}
	expected.ShouldBeEqual(t, 0, "SortedFunc returns correct value -- with args", actual)
}

func Test_CollectionSort_PackageLevelSort(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAsc(c)

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortAsc returns correct value -- with args", actual)
	coredynamic.SortDesc(c)
	actual2 := args.Map{"first": c.First()}
	expected2 := args.Map{"first": 3}
	expected2.ShouldBeEqual(t, 0, "SortDesc returns correct value -- with args", actual2)
}

func Test_CollectionSort_SortAscDescLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAscLock(c)

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortAscLock returns correct value -- with args", actual)
	coredynamic.SortDescLock(c)
	actual2 := args.Map{"first": c.First()}
	expected2 := args.Map{"first": 3}
	expected2.ShouldBeEqual(t, 0, "SortDescLock returns correct value -- with args", actual2)
}

func Test_CollectionSort_SortedAscDesc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	asc := coredynamic.SortedAsc(c)
	desc := coredynamic.SortedDesc(c)

	// Act
	actual := args.Map{
		"ascFirst": asc.First(),
		"descFirst": desc.First(),
	}

	// Assert
	expected := args.Map{
		"ascFirst": 1,
		"descFirst": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortedAsc/Desc returns correct value -- with args", actual)
}

func Test_CollectionSort_IsSorted(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	single := coredynamic.CollectionFrom([]int{1})

	// Act
	actual := args.Map{
		"asc":       coredynamic.IsSortedAsc(c),
		"desc":      coredynamic.IsSortedDesc(c),
		"singleAsc": coredynamic.IsSortedAsc(single),
	}

	// Assert
	expected := args.Map{
		"asc": true,
		"desc": false,
		"singleAsc": true,
	}
	expected.ShouldBeEqual(t, 0, "IsSorted returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// CollectionDistinct
// ═══════════════════════════════════════════

func Test_CollectionDistinct(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "a", "c", "b"})
	d := coredynamic.Distinct(c)
	u := coredynamic.Unique(c)
	dl := coredynamic.DistinctLock(c)
	dc := coredynamic.DistinctCount(c)
	id := coredynamic.IsDistinct(c)

	// Act
	actual := args.Map{
		"dLen": d.Length(),
		"uLen": u.Length(),
		"dlLen": dl.Length(),
		"dc": dc,
		"id": id,
	}

	// Assert
	expected := args.Map{
		"dLen": 3,
		"uLen": 3,
		"dlLen": 3,
		"dc": 3,
		"id": false,
	}
	expected.ShouldBeEqual(t, 0, "Distinct returns correct value -- with args", actual)
}

func Test_CollectionDistinct_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	d := coredynamic.Distinct(c)
	dc := coredynamic.DistinctCount(c)

	// Act
	actual := args.Map{
		"len": d.Length(),
		"dc": dc,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"dc": 0,
	}
	expected.ShouldBeEqual(t, 0, "Distinct returns empty -- empty", actual)
}

// ═══════════════════════════════════════════
// CollectionGroupBy
// ═══════════════════════════════════════════

func Test_CollectionGroupBy(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"apple", "avocado", "banana"})
	groups := coredynamic.GroupBy(c, func(s string) byte { return s[0] })

	// Act
	actual := args.Map{"keys": len(groups)}

	// Assert
	expected := args.Map{"keys": 2}
	expected.ShouldBeEqual(t, 0, "GroupBy returns correct value -- with args", actual)
}

func Test_CollectionGroupBy_Nil(t *testing.T) {
	// Arrange
	groups := coredynamic.GroupBy[string, byte](nil, func(s string) byte { return s[0] })

	// Act
	actual := args.Map{"keys": len(groups)}

	// Assert
	expected := args.Map{"keys": 0}
	expected.ShouldBeEqual(t, 0, "GroupBy returns nil -- nil", actual)
}

func Test_CollectionGroupBy_Lock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	groups := coredynamic.GroupByLock(c, func(s string) string { return s })

	// Act
	actual := args.Map{"keys": len(groups)}

	// Assert
	expected := args.Map{"keys": 2}
	expected.ShouldBeEqual(t, 0, "GroupByLock returns correct value -- with args", actual)
}

func Test_CollectionGroupBy_Count(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "a"})
	counts := coredynamic.GroupByCount(c, func(s string) string { return s })

	// Act
	actual := args.Map{
		"a": counts["a"],
		"b": counts["b"],
	}

	// Assert
	expected := args.Map{
		"a": 2,
		"b": 1,
	}
	expected.ShouldBeEqual(t, 0, "GroupByCount returns correct value -- with args", actual)
}

func Test_CollectionGroupBy_Count_Nil(t *testing.T) {
	// Arrange
	counts := coredynamic.GroupByCount[string, string](nil, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(counts)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupByCount returns nil -- nil", actual)
}

// ═══════════════════════════════════════════
// CollectionMap — Map, FlatMap, Reduce
// ═══════════════════════════════════════════

func Test_CollectionMap_Map(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	mapped := coredynamic.Map(c, func(i int) string { return "x" })

	// Act
	actual := args.Map{"len": mapped.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- with args", actual)
}

func Test_CollectionMap_Map_Nil(t *testing.T) {
	// Arrange
	mapped := coredynamic.Map[int, string](nil, func(i int) string { return "x" })

	// Act
	actual := args.Map{"len": mapped.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- nil", actual)
}

func Test_CollectionMap_FlatMap(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"ab", "cd"})
	flat := coredynamic.FlatMap(c, func(s string) []byte { return []byte(s) })

	// Act
	actual := args.Map{"len": flat.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "FlatMap returns correct value -- with args", actual)
}

func Test_CollectionMap_FlatMap_Nil(t *testing.T) {
	// Arrange
	flat := coredynamic.FlatMap[string, byte](nil, func(s string) []byte { return []byte(s) })

	// Act
	actual := args.Map{"len": flat.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "FlatMap returns nil -- nil", actual)
}

func Test_CollectionMap_Reduce(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sum := coredynamic.Reduce(c, 0, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "Reduce returns correct value -- with args", actual)
}

func Test_CollectionMap_Reduce_Nil(t *testing.T) {
	// Arrange
	sum := coredynamic.Reduce[int, int](nil, 99, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 99}
	expected.ShouldBeEqual(t, 0, "Reduce returns nil -- nil", actual)
}

// ═══════════════════════════════════════════
// CollectionLock — all Lock methods
// ═══════════════════════════════════════════

func Test_CollectionLock_LengthLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"len": c.LengthLock(),
		"empty": c.IsEmptyLock(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"empty": false,
	}
	expected.ShouldBeEqual(t, 0, "LengthLock returns correct value -- with args", actual)
}

func Test_CollectionLock_AddLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.AddLock("a")
	c.AddsLock("b", "c")
	c.AddManyLock("d")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "AddLock returns correct value -- with args", actual)
}

func Test_CollectionLock_AddCollectionLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c2 := coredynamic.CollectionFrom([]string{"a", "b"})
	c.AddCollectionLock(c2)
	c.AddCollectionLock(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddCollectionLock returns correct value -- with args", actual)
}

func Test_CollectionLock_AddCollectionsLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c1 := coredynamic.CollectionFrom([]string{"a"})
	c.AddCollectionsLock(c1, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddCollectionsLock returns correct value -- with args", actual)
}

func Test_CollectionLock_AddIfLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	c.AddIfLock(true, "a").AddIfLock(false, "b")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddIfLock returns correct value -- with args", actual)
}

func Test_CollectionLock_RemoveAtLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	ok := c.RemoveAtLock(0)
	bad := c.RemoveAtLock(99)

	// Act
	actual := args.Map{
		"ok": ok,
		"bad": bad,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"bad": false,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAtLock returns correct value -- with args", actual)
}

func Test_CollectionLock_ClearLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})
	c.ClearLock()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClearLock returns correct value -- with args", actual)
}

func Test_CollectionLock_ItemsLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	items := c.ItemsLock()

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ItemsLock returns correct value -- with args", actual)
}

func Test_CollectionLock_FirstLastLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"first": c.FirstLock(),
		"last": c.LastLock(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "FirstLastLock returns correct value -- with args", actual)
}

func Test_CollectionLock_AddWithWgLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](5)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, "a")
	wg.Wait()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddWithWgLock returns non-empty -- with args", actual)
}

func Test_CollectionLock_LoopLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	count := 0
	c.LoopLock(func(i int, item string) bool { count++; return false })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "LoopLock returns correct value -- with args", actual)
}

func Test_CollectionLock_FilterLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	f := c.FilterLock(func(i int) bool { return i%2 == 0 })

	// Act
	actual := args.Map{"len": f.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterLock returns correct value -- with args", actual)
}

func Test_CollectionLock_StringsLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})
	strs := c.StringsLock()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringsLock returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// CollectionTypes — factory shortcuts
// ═══════════════════════════════════════════

func Test_CollectionTypes(t *testing.T) {
	// Arrange
	sc := coredynamic.NewStringCollection(5)
	esc := coredynamic.EmptyStringCollection()
	ic := coredynamic.NewIntCollection(5)
	eic := coredynamic.EmptyIntCollection()
	i64c := coredynamic.NewInt64Collection(5)
	bc := coredynamic.NewByteCollection(5)
	boc := coredynamic.NewBoolCollection(5)
	fc := coredynamic.NewFloat64Collection(5)
	amc := coredynamic.NewAnyMapCollection(5)
	smc := coredynamic.NewStringMapCollection(5)

	// Act
	actual := args.Map{
		"sc": sc != nil, "esc": esc != nil, "ic": ic != nil, "eic": eic != nil,
		"i64c": i64c != nil, "bc": bc != nil, "boc": boc != nil, "fc": fc != nil,
		"amc": amc != nil, "smc": smc != nil,
	}

	// Assert
	expected := args.Map{
		"sc": true, "esc": true, "ic": true, "eic": true,
		"i64c": true, "bc": true, "boc": true, "fc": true,
		"amc": true, "smc": true,
	}
	expected.ShouldBeEqual(t, 0, "CollectionTypes returns correct value -- with args", actual)
}

// ═══════════════════════════════════════════
// newCreator — New.Collection pattern
// ═══════════════════════════════════════════

func Test_NewCreator_Collection(t *testing.T) {
	// Arrange
	sc := coredynamic.New.Collection.String.Cap(5)
	se := coredynamic.New.Collection.String.Empty()
	sf := coredynamic.New.Collection.String.From([]string{"a"})
	scl := coredynamic.New.Collection.String.Clone([]string{"a"})
	si := coredynamic.New.Collection.String.Items("a", "b")
	slc := coredynamic.New.Collection.String.LenCap(3, 10)
	scr := coredynamic.New.Collection.String.Create([]string{"a", "b"})

	// Act
	actual := args.Map{
		"scNN": sc != nil, "seNN": se != nil, "sfLen": sf.Length(), "sclLen": scl.Length(),
		"siLen": si.Length(), "slcLen": slc.Length(), "scrLen": scr.Length(),
	}

	// Assert
	expected := args.Map{
		"scNN": true, "seNN": true, "sfLen": 1, "sclLen": 1,
		"siLen": 2, "slcLen": 3, "scrLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "New.Collection.String returns correct value -- with args", actual)
}

func Test_NewCreator_Int(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.LenCap(2, 5)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int.LenCap returns correct value -- with args", actual)
}

func Test_NewCreator_Int64(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int64.LenCap(2, 5)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Int64.LenCap returns correct value -- with args", actual)
}

func Test_NewCreator_Byte(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Byte.LenCap(2, 5)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "New.Collection.Byte.LenCap returns correct value -- with args", actual)
}

func Test_NewCreator_Any(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Any.Cap(5)

	// Act
	actual := args.Map{"nn": c != nil}

	// Assert
	expected := args.Map{"nn": true}
	expected.ShouldBeEqual(t, 0, "New.Collection.Any returns correct value -- with args", actual)
}

func Test_NewCreator_Other(t *testing.T) {
	// Arrange
	bs := coredynamic.New.Collection.ByteSlice.Empty()
	bo := coredynamic.New.Collection.Bool.Empty()
	f32 := coredynamic.New.Collection.Float32.Empty()
	f64 := coredynamic.New.Collection.Float64.Empty()
	am := coredynamic.New.Collection.AnyMap.Empty()
	sm := coredynamic.New.Collection.StringMap.Empty()
	im := coredynamic.New.Collection.IntMap.Empty()

	// Act
	actual := args.Map{
		"bs": bs != nil, "bo": bo != nil, "f32": f32 != nil, "f64": f64 != nil,
		"am": am != nil, "sm": sm != nil, "im": im != nil,
	}

	// Assert
	expected := args.Map{
		"bs": true, "bo": true, "f32": true, "f64": true,
		"am": true, "sm": true, "im": true,
	}
	expected.ShouldBeEqual(t, 0, "New.Collection.Others returns correct value -- with args", actual)
}
