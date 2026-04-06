package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 5b
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg5_HS_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEmpty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"empty": h.IsEmpty(), "hasItems": h.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_Seg5_HS_Add(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Add", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Add("a").Add("b")
		actual := args.Map{"len": h.Length(), "has": h.Has("a")}
		expected := args.Map{"len": 2, "has": true}
		expected.ShouldBeEqual(t, 0, "Add -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddBool(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddBool", func() {
		h := corestr.New.Hashset.Cap(2)
		existed1 := h.AddBool("a")
		existed2 := h.AddBool("a")
		actual := args.Map{"existed1": existed1, "existed2": existed2}
		expected := args.Map{"existed1": false, "existed2": true}
		expected.ShouldBeEqual(t, 0, "AddBool -- new then existing", actual)
	})
}

func Test_Seg5_HS_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddNonEmpty", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddNonEmpty("a").AddNonEmpty("")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty -- skips empty", actual)
	})
}

func Test_Seg5_HS_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddNonEmptyWhitespace", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("  ")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyWhitespace -- skips whitespace", actual)
	})
}

func Test_Seg5_HS_AddIf(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddIf", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddIf(true, "a").AddIf(false, "b")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true", actual)
	})
}

func Test_Seg5_HS_AddIfMany(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddIfMany", func() {
		h := corestr.New.Hashset.Cap(4)
		h.AddIfMany(true, "a", "b").AddIfMany(false, "c")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddIfMany -- only true batch", actual)
	})
}

func Test_Seg5_HS_AddFunc(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddFunc", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddFunc(func() string { return "a" })
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddFunc -- added", actual)
	})
}

func Test_Seg5_HS_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddFuncErr", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddFuncErr(func() (string, error) { return "a", nil }, func(err error) {})
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddFuncErr -- added", actual)
	})
}

func Test_Seg5_HS_AddFuncErr_Error(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddFuncErr_Error", func() {
		h := corestr.New.Hashset.Cap(2)
		var handledErr error
		h.AddFuncErr(
			func() (string, error) { return "", &testErr{} },
			func(err error) { handledErr = err },
		)
		actual := args.Map{"handled": handledErr != nil, "len": h.Length()}
		expected := args.Map{"handled": true, "len": 0}
		expected.ShouldBeEqual(t, 0, "AddFuncErr error -- handled", actual)
	})
}

// testErr is defined in shared_compat_helpers.go

func Test_Seg5_HS_Adds(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Adds", func() {
		h := corestr.New.Hashset.Cap(4)
		h.Adds("a", "b", "c")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Adds -- 3 items", actual)
	})
}

func Test_Seg5_HS_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Adds_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.Adds(nil...)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddStrings(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStrings", func() {
		h := corestr.New.Hashset.Cap(4)
		h.AddStrings([]string{"a", "b"})
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStrings_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddStrings(nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddPtr(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddPtr", func() {
		h := corestr.New.Hashset.Cap(2)
		s := "test"
		h.AddPtr(&s)
		actual := args.Map{"has": h.Has("test")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddPtr -- added", actual)
	})
}

func Test_Seg5_HS_AddLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddLock", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddLock("a")
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLock -- added", actual)
	})
}

func Test_Seg5_HS_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddPtrLock", func() {
		h := corestr.New.Hashset.Cap(2)
		s := "test"
		h.AddPtrLock(&s)
		actual := args.Map{"has": h.Has("test")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddPtrLock -- added", actual)
	})
}

func Test_Seg5_HS_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddWithWgLock", func() {
		h := corestr.New.Hashset.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddWithWgLock("a", &wg)
		wg.Wait()
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddWithWgLock -- added", actual)
	})
}

func Test_Seg5_HS_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetItems", func() {
		h := corestr.New.Hashset.Cap(4)
		h2 := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.AddHashsetItems(h2)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetItems_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddHashsetItems(nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetItems nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMap", func() {
		h := corestr.New.Hashset.Cap(4)
		h.AddItemsMap(map[string]bool{"a": true, "b": false, "c": true})
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddItemsMap -- only true values", actual)
	})
}

func Test_Seg5_HS_AddItemsMap_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMap_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddItemsMap(nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMap nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMapWgLock", func() {
		h := corestr.New.Hashset.Cap(4)
		m := map[string]bool{"a": true, "b": false}
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddItemsMapWgLock(&m, &wg)
		wg.Wait()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddItemsMapWgLock -- only true", actual)
	})
}

func Test_Seg5_HS_AddItemsMapWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddItemsMapWgLock_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddItemsMapWgLock(nil, nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddItemsMapWgLock nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetWgLock", func() {
		h := corestr.New.Hashset.Cap(4)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddHashsetWgLock(h2, &wg)
		wg.Wait()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetWgLock -- 1 item", actual)
	})
}

func Test_Seg5_HS_AddHashsetWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddHashsetWgLock_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddHashsetWgLock(nil, nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetWgLock nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStringsPtrWgLock", func() {
		h := corestr.New.Hashset.Cap(4)
		wg := sync.WaitGroup{}
		wg.Add(1)
		h.AddStringsPtrWgLock([]string{"a", "b"}, &wg)
		wg.Wait()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsPtrWgLock -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStringsLock", func() {
		h := corestr.New.Hashset.Cap(4)
		h.AddStringsLock([]string{"a", "b"})
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddStringsLock -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddStringsLock_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddStringsLock(nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStringsLock nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollection", func() {
		h := corestr.New.Hashset.Cap(4)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		h.AddCollection(c)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollection -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollection_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddCollection(nil)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollection nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddCollections(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollections", func() {
		h := corestr.New.Hashset.Cap(4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		h.AddCollections(c1, c2)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddCollections -- 2 items", actual)
	})
}

func Test_Seg5_HS_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCollections_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddCollections(nil...)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollections nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddSimpleSlice", func() {
		h := corestr.New.Hashset.Cap(4)
		ss := corestr.SimpleSlice{"a", "b"}
		h.AddSimpleSlice(&ss)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSimpleSlice -- 2 items", actual)
	})
}

// ── Has / Contains / Missing ────────────────────────────────────────────────

func Test_Seg5_HS_Has_Contains(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Has_Contains", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{
			"has":      h.Has("a"),
			"contains": h.Contains("a"),
			"missing":  h.IsMissing("z"),
			"hasLock":  h.HasLock("a"),
		}
		expected := args.Map{
			"has":      true,
			"contains": true,
			"missing":  true,
			"hasLock":  true,
		}
		expected.ShouldBeEqual(t, 0, "Has/Contains/Missing -- correct", actual)
	})
}

func Test_Seg5_HS_IsMissingLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsMissingLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"missing": h.IsMissingLock("z")}
		expected := args.Map{"missing": true}
		expected.ShouldBeEqual(t, 0, "IsMissingLock -- true", actual)
	})
}

func Test_Seg5_HS_HasWithLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasWithLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"has": h.HasWithLock("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasWithLock -- true", actual)
	})
}

func Test_Seg5_HS_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAllStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"all": h.HasAllStrings([]string{"a", "b"}), "miss": h.HasAllStrings([]string{"a", "z"})}
		expected := args.Map{"all": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "HasAllStrings -- all and missing", actual)
	})
}

func Test_Seg5_HS_HasAll(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAll", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"all": h.HasAll("a", "b"), "miss": h.HasAll("a", "z")}
		expected := args.Map{"all": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "HasAll -- all and missing", actual)
	})
}

func Test_Seg5_HS_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAllCollectionItems", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"has": h.HasAllCollectionItems(c), "nil": h.HasAllCollectionItems(nil)}
		expected := args.Map{"has": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "HasAllCollectionItems -- found and nil", actual)
	})
}

func Test_Seg5_HS_HasAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"any": h.HasAny("z", "a"), "none": h.HasAny("x", "y")}
		expected := args.Map{"any": true, "none": false}
		expected.ShouldBeEqual(t, 0, "HasAny -- found and none", actual)
	})
}

func Test_Seg5_HS_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_HasAnyItem", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"has": h.HasAnyItem()}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyItem -- true", actual)
	})
}

func Test_Seg5_HS_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsAllMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"miss": h.IsAllMissing("x", "y"), "notMiss": h.IsAllMissing("a", "y")}
		expected := args.Map{"miss": true, "notMiss": false}
		expected.ShouldBeEqual(t, 0, "IsAllMissing -- all missing and partial", actual)
	})
}

// ── List / Items / Lines / Sorted ───────────────────────────────────────────

func Test_Seg5_HS_List(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_List", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.List())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "List -- 1 item", actual)
	})
}

func Test_Seg5_HS_Items(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Items", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.Items())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1 item", actual)
	})
}

func Test_Seg5_HS_SafeStrings(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SafeStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.SafeStrings())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SafeStrings -- 1 item", actual)
	})
}

func Test_Seg5_HS_SafeStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SafeStrings_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"len": len(h.SafeStrings())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeStrings empty -- 0", actual)
	})
}

func Test_Seg5_HS_Lines(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Lines", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.Lines())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Lines -- 1 item", actual)
	})
}

func Test_Seg5_HS_Lines_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Lines_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"len": len(h.Lines())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Lines empty -- 0", actual)
	})
}

func Test_Seg5_HS_OrderedList(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_OrderedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		result := h.OrderedList()
		actual := args.Map{"first": result[0]}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "OrderedList -- sorted asc", actual)
	})
}

func Test_Seg5_HS_OrderedList_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_OrderedList_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"len": len(h.OrderedList())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "OrderedList empty -- 0", actual)
	})
}

func Test_Seg5_HS_SortedList(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SortedList", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a", "b"})
		result := h.SortedList()
		actual := args.Map{"first": result[0]}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedList -- sorted", actual)
	})
}

func Test_Seg5_HS_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ListPtrSortedAsc", func() {
		h := corestr.New.Hashset.Strings([]string{"c", "a"})
		result := h.ListPtrSortedAsc()
		actual := args.Map{"first": result[0]}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "ListPtrSortedAsc -- sorted", actual)
	})
}

func Test_Seg5_HS_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ListPtrSortedDsc", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "c"})
		result := h.ListPtrSortedDsc()
		actual := args.Map{"first": result[0]}
		expected := args.Map{"first": "c"}
		expected.ShouldBeEqual(t, 0, "ListPtrSortedDsc -- descending", actual)
	})
}

func Test_Seg5_HS_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ListCopyLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.ListCopyLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListCopyLock -- 1", actual)
	})
}

func Test_Seg5_HS_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SimpleSlice", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": h.SimpleSlice().Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice -- 1 item", actual)
	})
}

func Test_Seg5_HS_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SimpleSlice_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"empty": h.SimpleSlice().IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice empty -- empty", actual)
	})
}

func Test_Seg5_HS_Collection(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Collection", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": h.Collection().Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection -- 1 item", actual)
	})
}

func Test_Seg5_HS_MapStringAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MapStringAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.MapStringAny())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapStringAny -- 1 item", actual)
	})
}

func Test_Seg5_HS_MapStringAny_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MapStringAny_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"len": len(h.MapStringAny())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "MapStringAny empty -- 0", actual)
	})
}

func Test_Seg5_HS_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MapStringAnyDiff", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.MapStringAnyDiff())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapStringAnyDiff -- 1 item", actual)
	})
}

// ── Filter ──────────────────────────────────────────────────────────────────

func Test_Seg5_HS_Filter(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Filter", func() {
		h := corestr.New.Hashset.Strings([]string{"aa", "b", "cc"})
		result := h.Filter(func(s string) bool { return len(s) > 1 })
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter -- 2 match", actual)
	})
}

func Test_Seg5_HS_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredItems", func() {
		h := corestr.New.Hashset.Strings([]string{"aa", "b"})
		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems -- 1 match", actual)
	})
}

func Test_Seg5_HS_GetFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredItems_Empty", func() {
		h := corestr.New.Hashset.Empty()
		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems empty -- 0", actual)
	})
}

func Test_Seg5_HS_GetFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredItems_Break", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		result := h.GetFilteredItems(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredItems break -- stops early", actual)
	})
}

func Test_Seg5_HS_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredCollection", func() {
		h := corestr.New.Hashset.Strings([]string{"aa", "b"})
		result := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection -- 1 match", actual)
	})
}

func Test_Seg5_HS_GetFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredCollection_Empty", func() {
		h := corestr.New.Hashset.Empty()
		result := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection empty -- empty", actual)
	})
}

func Test_Seg5_HS_GetFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetFilteredCollection_Break", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := h.GetFilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, true, true
		})
		actual := args.Map{"hasItems": result.HasAnyItem()}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "GetFilteredCollection break -- stops", actual)
	})
}

func Test_Seg5_HS_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddsUsingFilter", func() {
		h := corestr.New.Hashset.Cap(4)
		h.AddsUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, len(s) > 1, false },
			"a", "bb", "c",
		)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter -- 1 kept", actual)
	})
}

func Test_Seg5_HS_AddsUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddsUsingFilter_Nil", func() {
		h := corestr.New.Hashset.Cap(2)
		h.AddsUsingFilter(nil, nil...)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter nil -- no change", actual)
	})
}

func Test_Seg5_HS_AddsUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddsUsingFilter_Break", func() {
		h := corestr.New.Hashset.Cap(4)
		h.AddsUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, true },
			"a", "b",
		)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsUsingFilter break -- 1 item", actual)
	})
}

// ── Except ──────────────────────────────────────────────────────────────────

func Test_Seg5_HS_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptHashset", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		except := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptHashset(except)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset -- 1 remaining", actual)
	})
}

func Test_Seg5_HS_GetAllExceptHashset_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptHashset_Nil", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptHashset(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptHashset nil -- all items", actual)
	})
}

func Test_Seg5_HS_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExcept", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := h.GetAllExcept([]string{"a"})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExcept -- 1 remaining", actual)
	})
}

func Test_Seg5_HS_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExcept_Nil", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExcept(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil -- all items", actual)
	})
}

func Test_Seg5_HS_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptSpread", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := h.GetAllExceptSpread("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptSpread -- 1 remaining", actual)
	})
}

func Test_Seg5_HS_GetAllExceptSpread_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptSpread_Nil", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptSpread(nil...)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptSpread nil -- all items", actual)
	})
}

func Test_Seg5_HS_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptCollection", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		c := corestr.New.Collection.Strings([]string{"a"})
		result := h.GetAllExceptCollection(c)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection -- 1 remaining", actual)
	})
}

func Test_Seg5_HS_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_GetAllExceptCollection_Nil", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.GetAllExceptCollection(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil -- all items", actual)
	})
}

// ── Resize / AddCapacities ──────────────────────────────────────────────────

func Test_Seg5_HS_Resize(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Resize", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Resize(10)
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Resize -- preserved items", actual)
	})
}

func Test_Seg5_HS_Resize_SmallerThanLen(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Resize_SmallerThanLen", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Resize(1)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Resize smaller -- no change", actual)
	})
}

func Test_Seg5_HS_ResizeLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ResizeLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.ResizeLock(10)
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ResizeLock -- preserved items", actual)
	})
}

func Test_Seg5_HS_AddCapacities(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacities", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacities(10, 20)
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCapacities -- preserved items", actual)
	})
}

func Test_Seg5_HS_AddCapacities_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacities_Empty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacities()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCapacities empty -- no change", actual)
	})
}

func Test_Seg5_HS_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacitiesLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacitiesLock(10)
		actual := args.Map{"has": h.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCapacitiesLock -- preserved", actual)
	})
}

func Test_Seg5_HS_AddCapacitiesLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_AddCapacitiesLock_Empty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.AddCapacitiesLock()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCapacitiesLock empty -- no change", actual)
	})
}

// ── Concat ──────────────────────────────────────────────────────────────────

func Test_Seg5_HS_ConcatNewHashsets(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewHashsets", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		result := h.ConcatNewHashsets(true, h2)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets -- merged", actual)
	})
}

func Test_Seg5_HS_ConcatNewHashsets_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewHashsets_Empty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewHashsets(true)
		actual := args.Map{"has": result.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewHashsets empty -- cloned", actual)
	})
}

func Test_Seg5_HS_ConcatNewStrings(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewStrings", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true, []string{"b", "c"})
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings -- merged", actual)
	})
}

func Test_Seg5_HS_ConcatNewStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ConcatNewStrings_Empty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.ConcatNewStrings(true)
		actual := args.Map{"has": result.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ConcatNewStrings empty -- cloned", actual)
	})
}

// ── IsEquals / IsEqual ──────────────────────────────────────────────────────

func Test_Seg5_HS_IsEquals(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		h3 := corestr.New.Hashset.Strings([]string{"b"})
		actual := args.Map{
			"eq":      h1.IsEquals(h2),
			"neq":     h1.IsEquals(h3),
			"same":    h1.IsEquals(h1),
			"nilBoth": (*corestr.Hashset)(nil).IsEquals(nil),
			"nilOne":  h1.IsEquals(nil),
		}
		expected := args.Map{
			"eq":      true,
			"neq":     false,
			"same":    true,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_Seg5_HS_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals_DiffLen", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"eq": h1.IsEquals(h2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_Seg5_HS_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals_BothEmpty", func() {
		h1 := corestr.New.Hashset.Empty()
		h2 := corestr.New.Hashset.Empty()
		actual := args.Map{"eq": h1.IsEquals(h2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_Seg5_HS_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEquals_OneEmpty", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Empty()
		actual := args.Map{"eq": h1.IsEquals(h2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_Seg5_HS_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEqual", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"eq": h1.IsEqual(h2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual -- delegates to IsEquals", actual)
	})
}

func Test_Seg5_HS_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEqualsLock", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"eq": h1.IsEqualsLock(h2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock -- true", actual)
	})
}

// ── Remove ──────────────────────────────────────────────────────────────────

func Test_Seg5_HS_Remove(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Remove", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h.Remove("a")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Remove -- removed", actual)
	})
}

func Test_Seg5_HS_SafeRemove(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_SafeRemove", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.SafeRemove("a").SafeRemove("z")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeRemove -- safe", actual)
	})
}

func Test_Seg5_HS_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_RemoveWithLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.RemoveWithLock("a")
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveWithLock -- removed", actual)
	})
}

// ── String / Join ───────────────────────────────────────────────────────────

func Test_Seg5_HS_String(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_String", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"nonEmpty": h.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg5_HS_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_String_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"nonEmpty": h.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- has NoElements text", actual)
	})
}

func Test_Seg5_HS_StringLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_StringLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"nonEmpty": h.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_Seg5_HS_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_StringLock_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"nonEmpty": h.StringLock() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty -- NoElements", actual)
	})
}

func Test_Seg5_HS_Join(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Join", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"val": h.Join(",")}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Join -- value", actual)
	})
}

func Test_Seg5_HS_JoinSorted(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JoinSorted", func() {
		h := corestr.New.Hashset.Strings([]string{"b", "a"})
		actual := args.Map{"val": h.JoinSorted(",")}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "JoinSorted -- sorted", actual)
	})
}

func Test_Seg5_HS_JoinSorted_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JoinSorted_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"val": h.JoinSorted(",")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "JoinSorted empty -- empty", actual)
	})
}

func Test_Seg5_HS_JoinLine(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JoinLine", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"nonEmpty": h.JoinLine() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine -- non-empty", actual)
	})
}

func Test_Seg5_HS_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_NonEmptyJoins", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"nonEmpty": h.NonEmptyJoins(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins -- non-empty", actual)
	})
}

func Test_Seg5_HS_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_NonWhitespaceJoins", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"nonEmpty": h.NonWhitespaceJoins(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins -- non-empty", actual)
	})
}

// ── ToLowerSet ──────────────────────────────────────────────────────────────

func Test_Seg5_HS_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ToLowerSet", func() {
		h := corestr.New.Hashset.Strings([]string{"ABC"})
		result := h.ToLowerSet()
		actual := args.Map{"has": result.Has("abc")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "ToLowerSet -- lowered", actual)
	})
}

// ── Length / IsEmpty Lock ───────────────────────────────────────────────────

func Test_Seg5_HS_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Length_Nil", func() {
		var h *corestr.Hashset
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_Seg5_HS_LengthLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_LengthLock", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": h.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_Seg5_HS_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_IsEmptyLock", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"empty": h.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_Seg5_HS_Clear(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Clear", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Clear()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_Seg5_HS_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Clear_Nil", func() {
		var h *corestr.Hashset
		actual := args.Map{"nil": h.Clear() == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

func Test_Seg5_HS_Dispose(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Dispose", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		h.Dispose()
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_Seg5_HS_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Dispose_Nil", func() {
		var h *corestr.Hashset
		h.Dispose() // should not panic
	})
}

// ── Wrap / Transpile ────────────────────────────────────────────────────────

func Test_Seg5_HS_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapDoubleQuote", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapDoubleQuote()
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuote -- 1 item", actual)
	})
}

func Test_Seg5_HS_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapSingleQuote", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapSingleQuote()
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuote -- 1 item", actual)
	})
}

func Test_Seg5_HS_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapDoubleQuoteIfMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapDoubleQuoteIfMissing()
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapDoubleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_Seg5_HS_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_WrapSingleQuoteIfMissing", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.WrapSingleQuoteIfMissing()
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "WrapSingleQuoteIfMissing -- 1 item", actual)
	})
}

func Test_Seg5_HS_Transpile(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Transpile", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		result := h.Transpile(func(s string) string { return s + "!" })
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Transpile -- 1 item", actual)
	})
}

func Test_Seg5_HS_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Transpile_Empty", func() {
		h := corestr.New.Hashset.Empty()
		result := h.Transpile(func(s string) string { return s })
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Transpile empty -- empty", actual)
	})
}

// ── DistinctDiff ────────────────────────────────────────────────────────────

func Test_Seg5_HS_DistinctDiffLinesRaw(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		diff := h.DistinctDiffLinesRaw("b", "c")
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw -- 2 diff items", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw_BothEmpty", func() {
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLinesRaw()
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw both empty -- 0", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLinesRaw_LeftEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw_LeftEmpty", func() {
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLinesRaw("a")
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw left empty -- right items", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLinesRaw_RightEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLinesRaw_RightEmpty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		diff := h.DistinctDiffLinesRaw()
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLinesRaw right empty -- left items", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLines(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		diff := h.DistinctDiffLines("b", "c")
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines -- 2 diff items", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines_BothEmpty", func() {
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLines()
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines both empty -- 0", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLines_LeftNotEmpty_RightEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines_LeftNotEmpty_RightEmpty", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		diff := h.DistinctDiffLines()
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines left only -- left items", actual)
	})
}

func Test_Seg5_HS_DistinctDiffLines_LeftEmpty_RightNotEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffLines_LeftEmpty_RightNotEmpty", func() {
		h := corestr.New.Hashset.Empty()
		diff := h.DistinctDiffLines("a")
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "DistinctDiffLines right only -- right items", actual)
	})
}

func Test_Seg5_HS_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DistinctDiffHashset", func() {
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		h2 := corestr.New.Hashset.Strings([]string{"b", "c"})
		diff := h.DistinctDiffHashset(h2)
		actual := args.Map{"len": len(diff)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "DistinctDiffHashset -- 2 diff items", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_Seg5_HS_Json(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Json", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		j := h.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg5_HS_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_MarshalJSON", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		b, err := h.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg5_HS_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_UnmarshalJSON", func() {
		h := corestr.New.Hashset.Empty()
		err := h.UnmarshalJSON([]byte(`{"a":true}`))
		actual := args.Map{"noErr": err == nil, "len": h.Length()}
		expected := args.Map{"noErr": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg5_HS_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_UnmarshalJSON_Invalid", func() {
		h := corestr.New.Hashset.Empty()
		err := h.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg5_HS_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonModel", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(h.JsonModel())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_Seg5_HS_JsonModel_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonModel_Empty", func() {
		h := corestr.New.Hashset.Empty()
		actual := args.Map{"len": len(h.JsonModel())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "JsonModel empty -- 0", actual)
	})
}

func Test_Seg5_HS_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonModelAny", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notNil": h.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg5_HS_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ParseInjectUsingJson", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Empty()
		result, err := h2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "len": result.Length()}
		expected := args.Map{"noErr": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- round trip", actual)
	})
}

func Test_Seg5_HS_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_ParseInjectUsingJsonMust", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Empty()
		result := h2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg5_HS_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_JsonParseSelfInject", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		jr := h.JsonPtr()
		h2 := corestr.New.Hashset.Empty()
		err := h2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

func Test_Seg5_HS_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Serialize", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		b, err := h.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg5_HS_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_Deserialize", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		var dest map[string]bool
		err := h.Deserialize(&dest)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Seg5_HS_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_InterfaceCasts", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{
			"jsoner":   h.AsJsoner() != nil,
			"binder":   h.AsJsonContractsBinder() != nil,
			"injector": h.AsJsonParseSelfInjector() != nil,
			"marsh":    h.AsJsonMarshaller() != nil,
		}
		expected := args.Map{
			"jsoner":   true,
			"binder":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetDataModel
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg5_HS_DataModel(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DataModel", func() {
		dm := &corestr.HashsetDataModel{Items: map[string]bool{"a": true}}
		h := corestr.NewHashsetUsingDataModel(dm)
		actual := args.Map{"len": h.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetUsingDataModel -- 1 item", actual)
	})
}

func Test_Seg5_HS_DataModel_Reverse(t *testing.T) {
	safeTest(t, "Test_Seg5_HS_DataModel_Reverse", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		dm := corestr.NewHashsetsDataModelUsing(h)
		actual := args.Map{"len": len(dm.Items)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetsDataModelUsing -- 1 item", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// HashsetsCollection
// ══════════════════════════════════════════════════════════════════════════════

func Test_Seg5_HSC_IsEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEmpty", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)
		actual := args.Map{"empty": hsc.IsEmpty(), "hasItems": hsc.HasItems()}
		expected := args.Map{"empty": true, "hasItems": false}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection IsEmpty -- true", actual)
	})
}

func Test_Seg5_HSC_Add(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Add", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Add -- 1 hashset", actual)
	})
}

func Test_Seg5_HSC_AddNonNil(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddNonNil", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.AddNonNil(h).AddNonNil(nil)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonNil -- skips nil", actual)
	})
}

func Test_Seg5_HSC_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddNonEmpty", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		empty := corestr.New.Hashset.Empty()
		hsc.AddNonEmpty(h).AddNonEmpty(empty)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty -- skips empty", actual)
	})
}

func Test_Seg5_HSC_Adds(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Adds", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 4)
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc.Adds(h1, h2)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 hashsets", actual)
	})
}

func Test_Seg5_HSC_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Adds_Nil", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		hsc.Adds(nil...)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_Seg5_HSC_Length_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Length_Nil", func() {
		var hsc *corestr.HashsetsCollection
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_Seg5_HSC_LastIndex(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_LastIndex", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		actual := args.Map{"last": hsc.LastIndex()}
		expected := args.Map{"last": 0}
		expected.ShouldBeEqual(t, 0, "LastIndex -- 0", actual)
	})
}
func Test_Seg5_HSC_StringsList(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_StringsList", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(h)
		actual := args.Map{"len": len(hsc.StringsList())}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "StringsList -- 2 strings", actual)
	})
}

func Test_Seg5_HSC_StringsList_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_StringsList_Empty", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)
		actual := args.Map{"len": len(hsc.StringsList())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "StringsList empty -- 0", actual)
	})
}

func Test_Seg5_HSC_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ListDirectPtr", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		actual := args.Map{"len": len(*hsc.ListDirectPtr())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListDirectPtr -- 1 item", actual)
	})
}

func Test_Seg5_HSC_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddHashsetsCollection", func() {
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc2.Add(h2)
		hsc1.AddHashsetsCollection(hsc2)
		actual := args.Map{"len": hsc1.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetsCollection -- merged", actual)
	})
}

func Test_Seg5_HSC_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddHashsetsCollection_Nil", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		hsc.AddHashsetsCollection(nil)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetsCollection nil -- no change", actual)
	})
}

func Test_Seg5_HSC_ConcatNew(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ConcatNew", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc2.Add(h2)
		result := hsc.ConcatNew(hsc2)
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- merged", actual)
	})
}

func Test_Seg5_HSC_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ConcatNew_Empty", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		result := hsc.ConcatNew()
		actual := args.Map{"len": result.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty -- cloned", actual)
	})
}

func Test_Seg5_HSC_HasAll(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_HasAll", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(h)
		actual := args.Map{"has": hsc.HasAll("a", "b"), "miss": hsc.HasAll("a", "z")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "HasAll -- found and missing", actual)
	})
}

func Test_Seg5_HSC_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_HasAll_Empty", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)
		actual := args.Map{"has": hsc.HasAll("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasAll empty -- false", actual)
	})
}

func Test_Seg5_HSC_IsEqualPtr(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr", func() {
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		hsc1.Add(h1)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		hsc2.Add(h2)
		actual := args.Map{
			"eq":      hsc1.IsEqualPtr(hsc2),
			"same":    hsc1.IsEqualPtr(hsc1),
			"nilBoth": (*corestr.HashsetsCollection)(nil).IsEqualPtr(nil),
			"nilOne":  hsc1.IsEqualPtr(nil),
		}
		expected := args.Map{
			"eq":      true,
			"same":    true,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr -- various", actual)
	})
}

func Test_Seg5_HSC_IsEqualPtr_DiffLen(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_DiffLen", func() {
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr diff len -- false", actual)
	})
}

func Test_Seg5_HSC_IsEqualPtr_BothEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_BothEmpty", func() {
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 0)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr both empty -- true", actual)
	})
}

func Test_Seg5_HSC_IsEqualPtr_OneEmpty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_OneEmpty", func() {
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr one empty -- false", actual)
	})
}

func Test_Seg5_HSC_IsEqualPtr_Different(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_Different", func() {
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h1)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc2.Add(h2)
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr different -- false", actual)
	})
}

func Test_Seg5_HSC_IsEqual(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqual", func() {
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := *corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		hsc2.Add(h2)
		actual := args.Map{"eq": hsc1.IsEqual(hsc2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual -- delegates to IsEqualPtr", actual)
	})
}

func Test_Seg5_HSC_String(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_String", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		actual := args.Map{"nonEmpty": hsc.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_Seg5_HSC_String_Empty(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_String_Empty", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)
		actual := args.Map{"nonEmpty": hsc.String() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- has NoElements", actual)
	})
}

func Test_Seg5_HSC_Join(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Join", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		actual := args.Map{"nonEmpty": hsc.Join(",") != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Join -- non-empty", actual)
	})
}

func Test_Seg5_HSC_Json(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Json", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		j := hsc.Json()
		actual := args.Map{"noErr": !j.HasError()}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_Seg5_HSC_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_MarshalJSON", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		b, err := hsc.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_Seg5_HSC_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_UnmarshalJSON", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		b, _ := hsc.MarshalJSON()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		err := hsc2.UnmarshalJSON(b)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_Seg5_HSC_UnmarshalJSON_Invalid(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_UnmarshalJSON_Invalid", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)
		err := hsc.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_Seg5_HSC_Serialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Serialize", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		b, err := hsc.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_Seg5_HSC_Deserialize(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Deserialize", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		var dest interface{}
		err := hsc.Deserialize(&dest)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_Seg5_HSC_JsonModel(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_JsonModel", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		actual := args.Map{"notNil": hsc.JsonModel() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel -- non-nil", actual)
	})
}

func Test_Seg5_HSC_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_JsonModelAny", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		actual := args.Map{"notNil": hsc.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_Seg5_HSC_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ParseInjectUsingJson", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		_, err := hsc2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_Seg5_HSC_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ParseInjectUsingJsonMust", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		result := hsc2.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"notNil": result != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_Seg5_HSC_InterfaceCasts(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_InterfaceCasts", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		actual := args.Map{
			"jsoner":   hsc.AsJsoner() != nil,
			"binder":   hsc.AsJsonContractsBinder() != nil,
			"injector": hsc.AsJsonParseSelfInjector() != nil,
			"marsh":    hsc.AsJsonMarshaller() != nil,
		}
		expected := args.Map{
			"jsoner":   true,
			"binder":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_Seg5_HSC_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_JsonParseSelfInject", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		err := hsc2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── HashsetsCollectionDataModel ─────────────────────────────────────────────

func Test_Seg5_HSCDM_DataModel(t *testing.T) {
	safeTest(t, "Test_Seg5_HSCDM_DataModel", func() {
		h := corestr.New.Hashset.Strings([]string{"a"})
		dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{h}}
		hsc := corestr.NewHashsetsCollectionUsingDataModel(dm)
		actual := args.Map{"len": hsc.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionUsingDataModel -- 1 item", actual)
	})
}

func Test_Seg5_HSCDM_DataModel_Reverse(t *testing.T) {
	safeTest(t, "Test_Seg5_HSCDM_DataModel_Reverse", func() {
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		actual := args.Map{"len": len(dm.Items)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionDataModelUsing -- 1 item", actual)
	})
}
