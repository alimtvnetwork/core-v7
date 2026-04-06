package corestrtests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — IsEmpty / HasItems / Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_I30_Hashset_IsEmpty_New(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEmpty_New", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"empty": hs.IsEmpty(), "items": hs.HasItems(), "len": hs.Length(), "hasAny": hs.HasAnyItem()}
		expected := args.Map{"empty": true, "items": false, "len": 0, "hasAny": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- empty", actual)
	})
}

func Test_I30_Hashset_Length_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Length_Nil", func() {
		var hs *corestr.Hashset
		actual := args.Map{"len": hs.Length(), "empty": hs.IsEmpty()}
		expected := args.Map{"len": 0, "empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- nil length", actual)
	})
}

func Test_I30_Hashset_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEmptyLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"empty": hs.IsEmptyLock()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEmptyLock", actual)
	})
}

func Test_I30_Hashset_LengthLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_LengthLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("a")
		actual := args.Map{"len": hs.LengthLock()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- LengthLock", actual)
	})
}

func Test_I30_Hashset_Add(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Add", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Add("a")
		actual := args.Map{"has": hs.Has("a"), "len": hs.Length()}
		expected := args.Map{"has": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Add", actual)
	})
}

func Test_I30_Hashset_AddBool(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddBool", func() {
		hs := corestr.New.Hashset.Cap(5)
		existed1 := hs.AddBool("a")
		existed2 := hs.AddBool("a")
		actual := args.Map{"existed1": existed1, "existed2": existed2}
		expected := args.Map{"existed1": false, "existed2": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddBool", actual)
	})
}

func Test_I30_Hashset_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddNonEmpty", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddNonEmpty("")
		hs.AddNonEmpty("a")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddNonEmpty", actual)
	})
}

func Test_I30_Hashset_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddNonEmptyWhitespace", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddNonEmptyWhitespace("   ")
		hs.AddNonEmptyWhitespace("a")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddNonEmptyWhitespace", actual)
	})
}

func Test_I30_Hashset_AddIf(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddIf", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddIf(true, "a")
		hs.AddIf(false, "b")
		actual := args.Map{"hasA": hs.Has("a"), "hasB": hs.Has("b")}
		expected := args.Map{"hasA": true, "hasB": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIf", actual)
	})
}

func Test_I30_Hashset_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddIfMany", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddIfMany(true, "a", "b")
		hs.AddIfMany(false, "c")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIfMany", actual)
	})
}

func Test_I30_Hashset_AddFunc(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddFunc", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddFunc(func() string { return "computed" })
		actual := args.Map{"has": hs.Has("computed")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddFunc", actual)
	})
}

func Test_I30_Hashset_AddFuncErr_NoErr(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddFuncErr_NoErr", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		actual := args.Map{"has": hs.Has("ok")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddFuncErr no err", actual)
	})
}

func Test_I30_Hashset_AddFuncErr_Err(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddFuncErr_Err", func() {
		hs := corestr.New.Hashset.Cap(5)
		called := false
		hs.AddFuncErr(func() (string, error) { return "", fmt.Errorf("fail") }, func(e error) { called = true })
		actual := args.Map{"empty": hs.IsEmpty(), "called": called}
		expected := args.Map{"empty": true, "called": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- AddFuncErr err", actual)
	})
}

func Test_I30_Hashset_AddLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddLock("a")
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddLock", actual)
	})
}

func Test_I30_Hashset_AddPtr(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddPtr", func() {
		hs := corestr.New.Hashset.Cap(5)
		s := "hello"
		hs.AddPtr(&s)
		actual := args.Map{"has": hs.Has("hello")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddPtr", actual)
	})
}

func Test_I30_Hashset_AddPtrLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddPtrLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		s := "hello"
		hs.AddPtrLock(&s)
		actual := args.Map{"has": hs.Has("hello")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddPtrLock", actual)
	})
}

func Test_I30_Hashset_Adds(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Adds", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Adds("a", "b")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Adds", actual)
	})
}

func Test_I30_Hashset_Adds_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Adds_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.Adds(nil...)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- Adds nil", actual)
	})
}

func Test_I30_Hashset_AddStrings(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStrings", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStrings([]string{"a", "b"})
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddStrings", actual)
	})
}

func Test_I30_Hashset_AddStrings_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStrings_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStrings(nil)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddStrings nil", actual)
	})
}

func Test_I30_Hashset_AddStringsLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStringsLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStringsLock([]string{"a"})
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddStringsLock", actual)
	})
}

func Test_I30_Hashset_AddStringsLock_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStringsLock_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddStringsLock(nil)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddStringsLock nil", actual)
	})
}

func Test_I30_Hashset_AddHashsetItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetItems", func() {
		hs := corestr.New.Hashset.Cap(5)
		other := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs.AddHashsetItems(other)
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddHashsetItems", actual)
	})
}

func Test_I30_Hashset_AddHashsetItems_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetItems_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddHashsetItems(nil)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddHashsetItems nil", actual)
	})
}

func Test_I30_Hashset_AddItemsMap(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMap", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddItemsMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"hasA": hs.Has("a"), "hasB": hs.Has("b")}
		expected := args.Map{"hasA": true, "hasB": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddItemsMap", actual)
	})
}

func Test_I30_Hashset_AddItemsMap_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMap_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddItemsMap(nil)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddItemsMap nil", actual)
	})
}

func Test_I30_Hashset_AddCollection(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollection", func() {
		hs := corestr.New.Hashset.Cap(5)
		coll := corestr.New.Collection.Strings([]string{"a"})
		hs.AddCollection(coll)
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCollection", actual)
	})
}

func Test_I30_Hashset_AddCollection_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollection_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddCollection(nil)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddCollection nil", actual)
	})
}

func Test_I30_Hashset_AddCollections(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollections", func() {
		hs := corestr.New.Hashset.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		hs.AddCollections(c1, nil, c2)
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCollections", actual)
	})
}

func Test_I30_Hashset_AddCollections_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCollections_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddCollections(nil...)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddCollections nil", actual)
	})
}

func Test_I30_Hashset_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddWithWgLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddWithWgLock("a", wg)
		wg.Wait()
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- AddWithWgLock", actual)
	})
}

func Test_I30_Hashset_AddSimpleSlice(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddSimpleSlice", func() {
		hs := corestr.New.Hashset.Cap(5)
		ss := corestr.SimpleSlice([]string{"a", "b"})
		hs.AddSimpleSlice(&ss)
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddSimpleSlice", actual)
	})
}

func Test_I30_Hashset_Has_Contains(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Has_Contains", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"has": hs.Has("a"), "contains": hs.Contains("a"), "missing": hs.IsMissing("b")}
		expected := args.Map{"has": true, "contains": true, "missing": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Has/Contains/IsMissing", actual)
	})
}

func Test_I30_Hashset_HasLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasLock", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"hl": hs.HasLock("a"), "hwl": hs.HasWithLock("a"), "ml": hs.IsMissingLock("z")}
		expected := args.Map{"hl": true, "hwl": true, "ml": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- lock variants", actual)
	})
}

func Test_I30_Hashset_HasAllStrings(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAllStrings", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"all": hs.HasAllStrings([]string{"a", "b"}), "miss": hs.HasAllStrings([]string{"a", "c"})}
		expected := args.Map{"all": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAllStrings", actual)
	})
}

func Test_I30_Hashset_HasAll(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAll", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"all": hs.HasAll("a", "b"), "miss": hs.HasAll("a", "c")}
		expected := args.Map{"all": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAll", actual)
	})
}

func Test_I30_Hashset_HasAny(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAny", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"any": hs.HasAny("z", "a"), "none": hs.HasAny("x", "y")}
		expected := args.Map{"any": true, "none": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAny", actual)
	})
}

func Test_I30_Hashset_IsAllMissing(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsAllMissing", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"allMiss": hs.IsAllMissing("x", "y"), "notAll": hs.IsAllMissing("a", "x")}
		expected := args.Map{"allMiss": true, "notAll": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsAllMissing", actual)
	})
}

func Test_I30_Hashset_HasAllCollectionItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_HasAllCollectionItems", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		coll := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"has": hs.HasAllCollectionItems(coll), "nil": hs.HasAllCollectionItems(nil)}
		expected := args.Map{"has": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- HasAllCollectionItems", actual)
	})
}

func Test_I30_Hashset_List(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_List", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(hs.List())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- List", actual)
	})
}
func Test_I30_Hashset_ListCopyLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListCopyLock", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(hs.ListCopyLock())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListCopyLock", actual)
	})
}

func Test_I30_Hashset_Items(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Items", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"len": len(hs.Items())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Items", actual)
	})
}

func Test_I30_Hashset_Collection(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Collection", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		coll := hs.Collection()
		actual := args.Map{"len": coll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Collection", actual)
	})
}

func Test_I30_Hashset_SortedList(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SortedList", func() {
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		sorted := hs.SortedList()
		actual := args.Map{"first": sorted[0], "second": sorted[1]}
		expected := args.Map{"first": "a", "second": "b"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SortedList", actual)
	})
}

func Test_I30_Hashset_OrderedList(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_OrderedList", func() {
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		ol := hs.OrderedList()
		actual := args.Map{"len": len(ol)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- OrderedList", actual)
	})
}

func Test_I30_Hashset_OrderedList_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_OrderedList_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		ol := hs.OrderedList()
		actual := args.Map{"len": len(ol)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- OrderedList empty", actual)
	})
}

func Test_I30_Hashset_ListPtrSortedAsc(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListPtrSortedAsc", func() {
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		sorted := hs.ListPtrSortedAsc()
		actual := args.Map{"first": sorted[0]}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtrSortedAsc", actual)
	})
}

func Test_I30_Hashset_ListPtrSortedDsc(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ListPtrSortedDsc", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		sorted := hs.ListPtrSortedDsc()
		actual := args.Map{"first": sorted[0]}
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ListPtrSortedDsc", actual)
	})
}

func Test_I30_Hashset_SafeStrings(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SafeStrings", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"len": len(hs.SafeStrings())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- SafeStrings empty", actual)
	})
}

func Test_I30_Hashset_Lines(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Lines", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"len": len(hs.Lines())}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- Lines empty", actual)
	})
}

func Test_I30_Hashset_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SimpleSlice", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		ss := hs.SimpleSlice()
		actual := args.Map{"len": ss.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SimpleSlice", actual)
	})
}

func Test_I30_Hashset_SimpleSlice_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SimpleSlice_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		ss := hs.SimpleSlice()
		actual := args.Map{"empty": ss.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- SimpleSlice empty", actual)
	})
}

func Test_I30_Hashset_MapStringAny(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MapStringAny", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		m := hs.MapStringAny()
		actual := args.Map{"len": len(m)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- MapStringAny", actual)
	})
}

func Test_I30_Hashset_MapStringAny_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MapStringAny_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		m := hs.MapStringAny()
		actual := args.Map{"len": len(m)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- MapStringAny empty", actual)
	})
}

func Test_I30_Hashset_MapStringAnyDiff(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MapStringAnyDiff", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		d := hs.MapStringAnyDiff()
		actual := args.Map{"notNil": d != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- MapStringAnyDiff", actual)
	})
}

func Test_I30_Hashset_Resize(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Resize", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Resize(100)
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize", actual)
	})
}

func Test_I30_Hashset_Resize_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Resize_AlreadyLarger", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		hs.Resize(1)
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize already larger", actual)
	})
}

func Test_I30_Hashset_ResizeLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ResizeLock", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.ResizeLock(100)
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ResizeLock", actual)
	})
}

func Test_I30_Hashset_ResizeLock_AlreadyLarger(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ResizeLock_AlreadyLarger", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b", "c"})
		hs.ResizeLock(1)
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ResizeLock already larger", actual)
	})
}

func Test_I30_Hashset_AddCapacities(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacities", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacities(10, 20)
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCapacities", actual)
	})
}

func Test_I30_Hashset_AddCapacities_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacities_Empty", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacities()
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddCapacities empty", actual)
	})
}

func Test_I30_Hashset_AddCapacitiesLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacitiesLock", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacitiesLock(10)
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddCapacitiesLock", actual)
	})
}

func Test_I30_Hashset_AddCapacitiesLock_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddCapacitiesLock_Empty", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.AddCapacitiesLock()
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- AddCapacitiesLock empty", actual)
	})
}

func Test_I30_Hashset_ConcatNewHashsets_NoArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewHashsets_NoArgs", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		c := hs.ConcatNewHashsets(true)
		actual := args.Map{"has": c.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- ConcatNewHashsets no args", actual)
	})
}

func Test_I30_Hashset_ConcatNewHashsets_WithArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewHashsets_WithArgs", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		other := corestr.New.Hashset.Strings([]string{"b"})
		c := hs.ConcatNewHashsets(true, other, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- ConcatNewHashsets with args", actual)
	})
}

func Test_I30_Hashset_ConcatNewStrings_NoArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewStrings_NoArgs", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		c := hs.ConcatNewStrings(true)
		actual := args.Map{"has": c.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- ConcatNewStrings no args", actual)
	})
}

func Test_I30_Hashset_ConcatNewStrings_WithArgs(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ConcatNewStrings_WithArgs", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		c := hs.ConcatNewStrings(true, []string{"b", "c"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- ConcatNewStrings with args", actual)
	})
}

func Test_I30_Hashset_Filter(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Filter", func() {
		hs := corestr.New.Hashset.Strings([]string{"abc", "x"})
		filtered := hs.Filter(func(s string) bool { return len(s) > 1 })
		actual := args.Map{"len": filtered.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Filter", actual)
	})
}

func Test_I30_Hashset_GetFilteredItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredItems", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredItems", actual)
	})
}

func Test_I30_Hashset_GetFilteredItems_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredItems_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- GetFilteredItems empty", actual)
	})
}

func Test_I30_Hashset_GetFilteredItems_Break(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredItems_Break", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.GetFilteredItems(func(s string, i int) (string, bool, bool) { return s, true, true })
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredItems break", actual)
	})
}

func Test_I30_Hashset_GetFilteredCollection(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredCollection", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"len": coll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredCollection", actual)
	})
}

func Test_I30_Hashset_GetFilteredCollection_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredCollection_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"empty": coll.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- GetFilteredCollection empty", actual)
	})
}

func Test_I30_Hashset_GetFilteredCollection_Break(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetFilteredCollection_Break", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		coll := hs.GetFilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, true })
		actual := args.Map{"len": coll.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetFilteredCollection break", actual)
	})
}

func Test_I30_Hashset_GetAllExceptHashset(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptHashset", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		exc := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptHashset(exc)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExceptHashset", actual)
	})
}

func Test_I30_Hashset_GetAllExceptHashset_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptHashset_Nil", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptHashset(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExceptHashset nil", actual)
	})
}

func Test_I30_Hashset_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExcept", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.GetAllExcept([]string{"a"})
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExcept", actual)
	})
}

func Test_I30_Hashset_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExcept_Nil", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExcept(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExcept nil", actual)
	})
}

func Test_I30_Hashset_GetAllExceptSpread(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptSpread", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.GetAllExceptSpread("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExceptSpread", actual)
	})
}

func Test_I30_Hashset_GetAllExceptSpread_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptSpread_Nil", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptSpread(nil...)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExceptSpread nil", actual)
	})
}

func Test_I30_Hashset_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptCollection", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		coll := corestr.New.Collection.Strings([]string{"a"})
		result := hs.GetAllExceptCollection(coll)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- GetAllExceptCollection", actual)
	})
}

func Test_I30_Hashset_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_GetAllExceptCollection_Nil", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.GetAllExceptCollection(nil)
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- GetAllExceptCollection nil", actual)
	})
}

func Test_I30_Hashset_AddsUsingFilter(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsUsingFilter", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, "a", "b")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsUsingFilter", actual)
	})
}

func Test_I30_Hashset_AddsUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsUsingFilter_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(nil, nil...)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsUsingFilter nil", actual)
	})
}

func Test_I30_Hashset_AddsUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsUsingFilter_Break", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsUsingFilter break", actual)
	})
}

func Test_I30_Hashset_AddsAnyUsingFilter_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilter_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(nil, nil...)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsAnyUsingFilter nil", actual)
	})
}

func Test_I30_Hashset_AddsAnyUsingFilter_NilItem(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilter_NilItem", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false }, nil, "hello")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsAnyUsingFilter nil item", actual)
	})
}

func Test_I30_Hashset_AddsAnyUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilter_Break", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsAnyUsingFilter break", actual)
	})
}

func Test_I30_Hashset_AddsAnyUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilterLock_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilterLock(nil, nil...)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddsAnyUsingFilterLock nil", actual)
	})
}

func Test_I30_Hashset_AddsAnyUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddsAnyUsingFilterLock_Break", func() {
		hs := corestr.New.Hashset.Cap(5)
		hs.AddsAnyUsingFilterLock(func(s string, i int) (string, bool, bool) { return s, true, true }, "a", "b")
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddsAnyUsingFilterLock break", actual)
	})
}

func Test_I30_Hashset_Remove(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Remove", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hs.Remove("a")
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Remove", actual)
	})
}

func Test_I30_Hashset_SafeRemove(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_SafeRemove", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.SafeRemove("a")
		hs.SafeRemove("missing")
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- SafeRemove", actual)
	})
}

func Test_I30_Hashset_RemoveWithLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_RemoveWithLock", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.RemoveWithLock("a")
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- RemoveWithLock", actual)
	})
}

func Test_I30_Hashset_Clear(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Clear", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Clear()
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Clear", actual)
	})
}

func Test_I30_Hashset_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Clear_Nil", func() {
		var hs *corestr.Hashset
		result := hs.Clear()
		actual := args.Map{"nil": result == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- Clear nil", actual)
	})
}

func Test_I30_Hashset_Dispose(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Dispose", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hs.Dispose()
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Dispose", actual)
	})
}

func Test_I30_Hashset_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Dispose_Nil", func() {
		var hs *corestr.Hashset
		hs.Dispose() // should not panic
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- Dispose nil", actual)
	})
}

func Test_I30_Hashset_IsEquals_Same(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_Same", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"same": hs.IsEquals(hs)}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals same ptr", actual)
	})
}

func Test_I30_Hashset_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_BothNil", func() {
		var hs *corestr.Hashset
		actual := args.Map{"eq": hs.IsEquals(nil)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- IsEquals both nil", actual)
	})
}

func Test_I30_Hashset_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_OneNil", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"eq": hs.IsEquals(nil)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- IsEquals one nil", actual)
	})
}

func Test_I30_Hashset_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_BothEmpty", func() {
		hs1 := corestr.New.Hashset.Cap(5)
		hs2 := corestr.New.Hashset.Cap(5)
		actual := args.Map{"eq": hs1.IsEquals(hs2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- IsEquals both empty", actual)
	})
}

func Test_I30_Hashset_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_DiffLen", func() {
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"a", "b"})
		actual := args.Map{"eq": hs1.IsEquals(hs2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals diff len", actual)
	})
}

func Test_I30_Hashset_IsEquals_DiffItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEquals_DiffItems", func() {
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"b"})
		actual := args.Map{"eq": hs1.IsEquals(hs2)}
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEquals diff items", actual)
	})
}

func Test_I30_Hashset_IsEqual(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEqual", func() {
		hs1 := corestr.New.Hashset.Strings([]string{"a"})
		hs2 := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"eq": hs1.IsEqual(hs2)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEqual", actual)
	})
}

func Test_I30_Hashset_IsEqualsLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_IsEqualsLock", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"eq": hs.IsEqualsLock(hs)}
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- IsEqualsLock", actual)
	})
}

func Test_I30_Hashset_ToLowerSet(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ToLowerSet", func() {
		hs := corestr.New.Hashset.Strings([]string{"ABC"})
		lower := hs.ToLowerSet()
		actual := args.Map{"has": lower.Has("abc")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ToLowerSet", actual)
	})
}

func Test_I30_Hashset_String_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_String_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notEmpty": hs.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- String empty", actual)
	})
}

func Test_I30_Hashset_String_WithItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_String_WithItems", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notEmpty": hs.String() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- String with items", actual)
	})
}

func Test_I30_Hashset_StringLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_StringLock", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notEmpty": hs.StringLock() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- StringLock empty", actual)
	})
}

func Test_I30_Hashset_StringLock_WithItems(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_StringLock_WithItems", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notEmpty": hs.StringLock() != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- StringLock with items", actual)
	})
}

func Test_I30_Hashset_Join(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Join", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"val": hs.Join(",")}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Join", actual)
	})
}

func Test_I30_Hashset_JoinLine(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JoinLine", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"val": hs.JoinLine()}
		expected := args.Map{"val": "a"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JoinLine", actual)
	})
}

func Test_I30_Hashset_JoinSorted_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JoinSorted_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"val": hs.JoinSorted(",")}
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- JoinSorted empty", actual)
	})
}

func Test_I30_Hashset_JoinSorted(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JoinSorted", func() {
		hs := corestr.New.Hashset.Strings([]string{"b", "a"})
		actual := args.Map{"val": hs.JoinSorted(",")}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JoinSorted", actual)
	})
}

func Test_I30_Hashset_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_NonEmptyJoins", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notEmpty": hs.NonEmptyJoins(",") != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- NonEmptyJoins", actual)
	})
}

func Test_I30_Hashset_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_NonWhitespaceJoins", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notEmpty": hs.NonWhitespaceJoins(",") != ""}
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- NonWhitespaceJoins", actual)
	})
}

func Test_I30_Hashset_JsonModel(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonModel", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jm := hs.JsonModel()
		actual := args.Map{"len": len(jm)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonModel", actual)
	})
}

func Test_I30_Hashset_JsonModel_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonModel_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		jm := hs.JsonModel()
		actual := args.Map{"len": len(jm)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- JsonModel empty", actual)
	})
}

func Test_I30_Hashset_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonModelAny", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		actual := args.Map{"notNil": hs.JsonModelAny() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonModelAny", actual)
	})
}

func Test_I30_Hashset_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_MarshalJSON", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		b, err := hs.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- MarshalJSON", actual)
	})
}

func Test_I30_Hashset_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_UnmarshalJSON", func() {
		hs := corestr.New.Hashset.Cap(5)
		err := hs.UnmarshalJSON([]byte(`{"a":true}`))
		actual := args.Map{"noErr": err == nil, "has": hs.Has("a")}
		expected := args.Map{"noErr": true, "has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- UnmarshalJSON", actual)
	})
}

func Test_I30_Hashset_UnmarshalJSON_Err(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_UnmarshalJSON_Err", func() {
		hs := corestr.New.Hashset.Cap(5)
		err := hs.UnmarshalJSON([]byte(`{invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- UnmarshalJSON err", actual)
	})
}

func Test_I30_Hashset_Json(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Json", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		j := hs.Json()
		actual := args.Map{"hasBytes": j.HasBytes()}
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Json", actual)
	})
}

func Test_I30_Hashset_JsonPtr(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonPtr", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jp := hs.JsonPtr()
		actual := args.Map{"notNil": jp != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonPtr", actual)
	})
}

func Test_I30_Hashset_Serialize(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Serialize", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		b, err := hs.Serialize()
		actual := args.Map{"noErr": err == nil, "hasBytes": len(b) > 0}
		expected := args.Map{"noErr": true, "hasBytes": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Serialize", actual)
	})
}

func Test_I30_Hashset_Deserialize(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Deserialize", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		target := map[string]bool{}
		err := hs.Deserialize(&target)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Deserialize", actual)
	})
}

func Test_I30_Hashset_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ParseInjectUsingJson", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Cap(5)
		result, err := hs2.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "has": result.Has("a")}
		expected := args.Map{"noErr": true, "has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- ParseInjectUsingJson", actual)
	})
}

func Test_I30_Hashset_ParseInjectUsingJson_Err(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_ParseInjectUsingJson_Err", func() {
		hs := corestr.New.Hashset.Cap(5)
		badJson := corejson.NewPtr(42)
		_, err := hs.ParseInjectUsingJson(badJson)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns error -- ParseInjectUsingJson err", actual)
	})
}

func Test_I30_Hashset_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_JsonParseSelfInject", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		jr := hs.JsonPtr()
		hs2 := corestr.New.Hashset.Cap(5)
		err := hs2.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil}
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- JsonParseSelfInject", actual)
	})
}

func Test_I30_Hashset_AsJsoner(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsoner", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsoner() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsoner", actual)
	})
}

func Test_I30_Hashset_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsonContractsBinder", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsonContractsBinder() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsonContractsBinder", actual)
	})
}

func Test_I30_Hashset_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsonParseSelfInjector", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsonParseSelfInjector() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsonParseSelfInjector", actual)
	})
}

func Test_I30_Hashset_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AsJsonMarshaller", func() {
		hs := corestr.New.Hashset.Cap(5)
		actual := args.Map{"notNil": hs.AsJsonMarshaller() != nil}
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AsJsonMarshaller", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLinesRaw_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_BothEmpty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLinesRaw()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- DistinctDiffLinesRaw both empty", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLinesRaw_LeftOnly(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_LeftOnly", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.DistinctDiffLinesRaw()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLinesRaw left only", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLinesRaw_RightOnly(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_RightOnly", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLinesRaw("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLinesRaw right only", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLinesRaw_Both(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLinesRaw_Both", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.DistinctDiffLinesRaw("b", "c")
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLinesRaw both", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLines_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_BothEmpty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLines()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- DistinctDiffLines both empty", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLines_LeftOnly(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_LeftOnly", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.DistinctDiffLines()
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLines left only", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLines_RightOnly(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_RightOnly", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.DistinctDiffLines("a")
		actual := args.Map{"len": len(result)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLines right only", actual)
	})
}

func Test_I30_Hashset_DistinctDiffLines_Both(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffLines_Both", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		result := hs.DistinctDiffLines("b", "c")
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffLines both", actual)
	})
}

func Test_I30_Hashset_DistinctDiffHashset(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_DistinctDiffHashset", func() {
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		other := corestr.New.Hashset.Strings([]string{"b", "c"})
		result := hs.DistinctDiffHashset(other)
		actual := args.Map{"hasItems": len(result) > 0}
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- DistinctDiffHashset", actual)
	})
}

func Test_I30_Hashset_WrapDoubleQuote(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapDoubleQuote", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapDoubleQuote()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapDoubleQuote", actual)
	})
}

func Test_I30_Hashset_WrapDoubleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapDoubleQuoteIfMissing", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapDoubleQuoteIfMissing()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapDoubleQuoteIfMissing", actual)
	})
}

func Test_I30_Hashset_WrapSingleQuote(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapSingleQuote", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapSingleQuote()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapSingleQuote", actual)
	})
}

func Test_I30_Hashset_WrapSingleQuoteIfMissing(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_WrapSingleQuoteIfMissing", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		result := hs.WrapSingleQuoteIfMissing()
		actual := args.Map{"hasAny": result.HasAnyItem()}
		expected := args.Map{"hasAny": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- WrapSingleQuoteIfMissing", actual)
	})
}

func Test_I30_Hashset_Transpile_Empty(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_Transpile_Empty", func() {
		hs := corestr.New.Hashset.Cap(5)
		result := hs.Transpile(func(s string) string { return s })
		actual := args.Map{"empty": result.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns empty -- Transpile empty", actual)
	})
}

func Test_I30_Hashset_AddStringsPtrWgLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddStringsPtrWgLock", func() {
		hs := corestr.New.Hashset.Cap(200)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddStringsPtrWgLock([]string{"a", "b"}, wg)
		wg.Wait()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddStringsPtrWgLock", actual)
	})
}

func Test_I30_Hashset_AddHashsetWgLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetWgLock", func() {
		hs := corestr.New.Hashset.Cap(200)
		other := corestr.New.Hashset.Strings([]string{"a"})
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddHashsetWgLock(other, wg)
		wg.Wait()
		actual := args.Map{"has": hs.Has("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddHashsetWgLock", actual)
	})
}

func Test_I30_Hashset_AddHashsetWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddHashsetWgLock_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		hs.AddHashsetWgLock(nil, wg)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddHashsetWgLock nil", actual)
	})
}

func Test_I30_Hashset_AddItemsMapWgLock(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMapWgLock", func() {
		hs := corestr.New.Hashset.Cap(200)
		m := map[string]bool{"a": true, "b": false}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		hs.AddItemsMapWgLock(&m, wg)
		wg.Wait()
		actual := args.Map{"hasA": hs.Has("a"), "hasB": hs.Has("b")}
		expected := args.Map{"hasA": true, "hasB": false}
		expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddItemsMapWgLock", actual)
	})
}

func Test_I30_Hashset_AddItemsMapWgLock_Nil(t *testing.T) {
	safeTest(t, "Test_I30_Hashset_AddItemsMapWgLock_Nil", func() {
		hs := corestr.New.Hashset.Cap(5)
		wg := &sync.WaitGroup{}
		hs.AddItemsMapWgLock(nil, wg)
		actual := args.Map{"empty": hs.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashset returns nil -- AddItemsMapWgLock nil", actual)
	})
}
