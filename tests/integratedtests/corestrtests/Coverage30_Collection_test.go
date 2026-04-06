package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Collection basic methods ──

func Test_C30_Collection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Length_Nil", func() {
		var c *corestr.Collection
		actual := args.Map{"result": c.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C30_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.HasAnyItem()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		c.Add("x")
		actual := args.Map{"result": c.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_C30_Collection_LastIndex", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.LastIndex() != -1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C30_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.HasIndex(1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.HasIndex(5)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ListStringsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.ListStringsPtr()
	})
}

func Test_C30_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.ListStrings()
	})
}

func Test_C30_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_C30_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.StringJSON()
	})
}

func Test_C30_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_C30_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"result": c.RemoveAt(1)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.RemoveAt(-1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual := args.Map{"result": c.RemoveAt(99)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_Count(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Count", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		actual := args.Map{"result": c.Capacity() < 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected >= 10", actual)
	})
}

func Test_C30_Collection_Capacity_Nil(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Capacity_Nil", func() {
		c := &corestr.Collection{}
		_ = c.Capacity()
	})
}

func Test_C30_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.LengthLock() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_C30_Collection_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEqualsWithSensitive", func() {
		a := corestr.New.Collection.Strings([]string{"A"})
		b := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": a.IsEqualsWithSensitive(true, b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
		actual := args.Map{"result": a.IsEqualsWithSensitive(false, b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal case insensitive", actual)
	})
}

func Test_C30_Collection_IsEquals_NilBoth(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals_NilBoth", func() {
		var a, b *corestr.Collection
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_IsEquals_OneNil(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals_OneNil", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": a.IsEquals(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals_BothEmpty", func() {
		a := corestr.New.Collection.Empty()
		b := corestr.New.Collection.Empty()
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals_OneEmpty", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Empty()
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals_DiffLen", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": a.IsEquals(b)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEquals_SamePtr", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": a.IsEquals(a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for same pointer", actual)
	})
}

func Test_C30_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEmptyLock", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.IsEmptyLock()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsEmpty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.HasItems()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddLock("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("  ")
		c.AddNonEmptyWhitespace("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_Add(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Add", func() {
		c := corestr.New.Collection.Empty()
		c.Add("a")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddError", func() {
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		c.AddError(errors.New("err"))
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"result": c.AsDefaultError() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		c.Add("err1")
		actual := args.Map{"result": c.AsDefaultError() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C30_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AsError", func() {
		c := corestr.New.Collection.Empty()
		_ = c.AsError(",")
	})
}

func Test_C30_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddIf", func() {
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "add")
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C30_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c"})
		r := c.EachItemSplitBy(",")
		actual := args.Map{"result": len(r)}
		expected := args.Map{"result": 3}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C30_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.ConcatNew(0, "b", "c")
		actual := args.Map{"result": n.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C30_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ConcatNew_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.ConcatNew(0)
		actual := args.Map{"result": n.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.ToError(",")
	})
}

func Test_C30_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.ToDefaultError()
	})
}

func Test_C30_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		c.AddIfMany(true, "a", "b")
		actual := args.Map{"result": c.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C30_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddFunc", func() {
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "x" })
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddFuncErr", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(func() (string, error) { return "x", nil }, func(e error) {})
		c.AddFuncErr(func() (string, error) { return "", errors.New("e") }, func(e error) {})
		actual := args.Map{"result": c.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C30_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddsLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")
	})
}

func Test_C30_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Adds", func() {
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b")
	})
}

func Test_C30_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"a"})
	})
}

func Test_C30_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddCollection", func() {
		c := corestr.New.Collection.Empty()
		other := corestr.New.Collection.Strings([]string{"a"})
		c.AddCollection(other)
		c.AddCollection(corestr.New.Collection.Empty())
	})
}

func Test_C30_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddCollections", func() {
		c := corestr.New.Collection.Empty()
		c.AddCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Empty())
	})
}

func Test_C30_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddPointerCollectionsLock(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_C30_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsValues(hm)
		c.AddHashmapsValues(nil)
	})
}

func Test_C30_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeys(hm)
		c.AddHashmapsKeys(nil)
	})
}

func Test_C30_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(hm)
		c.AddHashmapsKeysValues(nil)
	})
}

func Test_C30_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		hm := corestr.New.Hashmap.Empty()
		hm.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValuesUsingFilter(func(kv corestr.KeyValuePair) (string, bool, bool) {
			return kv.Key + "=" + kv.Value, true, false
		}, hm)
		c.AddHashmapsKeysValuesUsingFilter(func(kv corestr.KeyValuePair) (string, bool, bool) {
			return "", false, true
		}, hm)
		c.AddHashmapsKeysValuesUsingFilter(nil, nil)
	})
}

func Test_C30_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()
	})
}

func Test_C30_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.IndexAt(1) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C30_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.SafeIndexAtUsingLength("def", 1, 0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual := args.Map{"result": c.SafeIndexAtUsingLength("def", 0, 1) != "def"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_C30_Collection_First(t *testing.T) {
	safeTest(t, "Test_C30_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C30_Collection_Last(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.Last() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C30_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C30_Collection_LastOrDefault", func() {
		c := corestr.New.Collection.Empty()
		_ = c.LastOrDefault()
		c.Add("x")
		_ = c.LastOrDefault()
	})
}

func Test_C30_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C30_Collection_FirstOrDefault", func() {
		c := corestr.New.Collection.Empty()
		_ = c.FirstOrDefault()
		c.Add("x")
		_ = c.FirstOrDefault()
	})
}

func Test_C30_Collection_Take(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		_ = c.Take(2)
		_ = c.Take(0)
		_ = c.Take(5)
	})
}

func Test_C30_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		_ = c.Skip(1)
		_ = c.Skip(0)
	})
}

func Test_C30_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		c2 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2.Reverse()
		c3 := corestr.New.Collection.Strings([]string{"a"})
		c3.Reverse()
	})
}

func Test_C30_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		actual := args.Map{"result": c.GetPagesSize(2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		actual := args.Map{"result": c.GetPagesSize(0) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C30_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetPagedCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagedCollection(2)
		_ = pages
	})
}

func Test_C30_Collection_GetPagedCollection_Small(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetPagedCollection_Small", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.GetPagedCollection(5)
	})
}

func Test_C30_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		_ = c.GetSinglePageCollection(2, 1)
		_ = c.GetSinglePageCollection(2, 3)
	})
}

func Test_C30_Collection_GetSinglePageCollection_Small(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetSinglePageCollection_Small", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.GetSinglePageCollection(5, 1)
	})
}

func Test_C30_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_C30_Collection_InsertAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.InsertAt(1, "x")
		c2 := corestr.New.Collection.Empty()
		c2.InsertAt(0, "x")
	})
}

func Test_C30_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)
	})
}

func Test_C30_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_C30_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c2.RemoveItemsIndexes(true)
	})
}

func Test_C30_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Empty()
		c.AppendCollectionPtr(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_C30_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Empty()
		c.AppendCollections(corestr.New.Collection.Strings([]string{"a"}), corestr.New.Collection.Empty())
		c.AppendCollections()
	})
}

func Test_C30_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock("a", 42)
		c.AppendAnysLock()
	})
}

func Test_C30_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", nil, 42)
		c.AppendAnys()
	})
}

func Test_C30_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendAnysUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "a", nil, 42)
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return "", false, true
		}, "x")
		c.AppendAnysUsingFilter(nil)
	})
}

func Test_C30_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendAnysUsingFilterLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "a", nil)
		c.AppendAnysUsingFilterLock(func(s string, i int) (string, bool, bool) {
			return "", false, true
		}, "x")
		c.AppendAnysUsingFilterLock(nil)
	})
}

func Test_C30_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil)
		c.AppendNonEmptyAnys(nil)
	})
}

func Test_C30_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("", "a")
		c.AddsNonEmpty(nil...)
	})
}

func Test_C30_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Empty()
		s := "a"
		c.AddsNonEmptyPtrLock(nil, &s)
	})
}

func Test_C30_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_C30_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		m := c.UniqueBoolMap()
		actual := args.Map{"result": len(m) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C30_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.UniqueBoolMapLock()
	})
}

func Test_C30_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_C30_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		_ = c.UniqueList()
	})
}

func Test_C30_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.UniqueListLock()
	})
}

func Test_C30_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.Filter(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		c2 := corestr.New.Collection.Empty()
		r2 := c2.Filter(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(r2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C30_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.FilterLock(func(s string, i int) (string, bool, bool) { return s, true, false })
	})
}

func Test_C30_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_C30_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.FilteredCollection(func(s string, i int) (string, bool, bool) { return s, true, false })
	})
}

func Test_C30_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_FilteredCollectionLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) { return s, true, false })
	})
}

func Test_C30_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.FilterPtr(func(s *string, i int) (*string, bool, bool) { return s, true, false })
		c2 := corestr.New.Collection.Empty()
		_ = c2.FilterPtr(func(s *string, i int) (*string, bool, bool) { return s, true, false })
	})
}

func Test_C30_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) { return s, true, false })
	})
}

func Test_C30_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_C30_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"", "a"})
		r := c.NonEmptyList()
		actual := args.Map{"result": len(r) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		_ = c.NonEmptyListPtr()
	})
}

func Test_C30_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.HashsetAsIs()
	})
}

func Test_C30_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.HashsetWithDoubleLength()
	})
}

func Test_C30_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.HashsetLock()
	})
}

func Test_C30_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_C30_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"", "a"})
		_ = c.NonEmptyItems()
		_ = c.NonEmptyItemsPtr()
	})
}

func Test_C30_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C30_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"  ", "a"})
		_ = c.NonEmptyItemsOrNonWhitespace()
		_ = c.NonEmptyItemsOrNonWhitespacePtr()
	})
}

func Test_C30_Collection_Items(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Items", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.Items()
		_ = c.ListPtr()
	})
}

func Test_C30_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.ListCopyPtrLock()
		c2 := corestr.New.Collection.Empty()
		_ = c2.ListCopyPtrLock()
	})
}

func Test_C30_Collection_Has(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.Has("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		c2 := corestr.New.Collection.Empty()
		actual := args.Map{"result": c2.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.HasLock("a")
	})
}

func Test_C30_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		actual := args.Map{"result": c.HasPtr(&s)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.HasPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C30_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		c2 := corestr.New.Collection.Empty()
		actual := args.Map{"result": c2.HasAll("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		r := c.SortedListAsc()
		actual := args.Map{"result": r[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected sorted", actual)
		c2 := corestr.New.Collection.Empty()
		_ = c2.SortedListAsc()
	})
}

func Test_C30_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAsc()
		c2 := corestr.New.Collection.Empty()
		c2.SortedAsc()
	})
}

func Test_C30_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()
		c2 := corestr.New.Collection.Empty()
		c2.SortedAscLock()
	})
}

func Test_C30_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.SortedListDsc()
	})
}

func Test_C30_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_C30_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"A"})
		actual := args.Map{"result": c.HasUsingSensitivity("A", true)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.HasUsingSensitivity("a", false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true case insensitive", actual)
	})
}

func Test_C30_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsContainsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		actual := args.Map{"result": c.IsContainsPtr(&s)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_, has := c.GetHashsetPlusHasAll([]string{"a"})
		actual := args.Map{"result": has}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		_, has2 := c.GetHashsetPlusHasAll(nil)
		actual := args.Map{"result": has2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": c.IsContainsAllSlice([]string{"a"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": c.IsContainsAllSlice([]string{})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C30_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": c.IsContainsAll("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C30_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.IsContainsAllLock("a")
	})
}

func Test_C30_Collection_New(t *testing.T) {
	safeTest(t, "Test_C30_Collection_New", func() {
		c := corestr.New.Collection.Empty()
		n := c.New("a", "b")
		actual := args.Map{"result": n.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		n2 := c.New()
		actual := args.Map{"result": n2.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C30_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("", "a")
		c.AddNonEmptyStrings()
	})
}

func Test_C30_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" })
		c.AddFuncResult(nil...)
	})
}

func Test_C30_Collection_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddNonEmptyStringsSlice", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a"})
		c.AddNonEmptyStringsSlice([]string{})
	})
}

func Test_C30_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "bb"}, func(s string) bool { return len(s) == 1 })
	})
}

func Test_C30_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string { return []string{s} })
	})
}

func Test_C30_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_C30_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
	})
}

func Test_C30_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExceptCollection(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": len(r) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		r2 := c.GetAllExceptCollection(nil)
		actual := args.Map{"result": len(r2) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C30_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_C30_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.GetAllExcept([]string{"a"})
		_ = c.GetAllExcept(nil)
	})
}

func Test_C30_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_C30_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"apple", "banana"})
		_ = c.CharCollectionMap()
	})
}

func Test_C30_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.SummaryString(1)
	})
}

func Test_C30_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_C30_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Empty()
		_ = c.SummaryStringWithHeader("Header")
		c2 := corestr.New.Collection.Strings([]string{"a"})
		_ = c2.SummaryStringWithHeader("Header")
	})
}

func Test_C30_Collection_String(t *testing.T) {
	safeTest(t, "Test_C30_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.String()
		c2 := corestr.New.Collection.Empty()
		_ = c2.String()
	})
}

func Test_C30_Collection_CsvMethods(t *testing.T) {
	safeTest(t, "Test_C30_Collection_CsvMethods", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.CsvLines()
		_ = c.CsvLinesOptions(true)
		_ = c.Csv()
		_ = c.CsvOptions(true)
		c2 := corestr.New.Collection.Empty()
		_ = c2.Csv()
	})
}

func Test_C30_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_C30_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.StringLock()
		c2 := corestr.New.Collection.Empty()
		_ = c2.StringLock()
	})
}

func Test_C30_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Empty()
		c.AddCapacity(10)
		c.AddCapacity()
	})
}

func Test_C30_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Resize", func() {
		c := corestr.New.Collection.Cap(5)
		c.Resize(10)
		c.Resize(1)
	})
}

func Test_C30_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.Joins(",")
		_ = c.Joins(",", "c")
	})
}

func Test_C30_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_C30_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		_ = c.NonEmptyJoins(",")
	})
}

func Test_C30_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_C30_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		_ = c.NonWhitespaceJoins(",")
	})
}

func Test_C30_Collection_JsonMethods(t *testing.T) {
	safeTest(t, "Test_C30_Collection_JsonMethods", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.JsonModel()
		_ = c.JsonModelAny()
		_, _ = c.MarshalJSON()
		_ = c.Json()
		_ = c.JsonPtr()
		_ = c.JsonString()
		_ = c.JsonStringMust()
		_ = c.AsJsonMarshaller()
		_ = c.AsJsonContractsBinder()
		_, _ = c.Serialize()
	})
}

func Test_C30_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C30_Collection_UnmarshalJSON", func() {
		c := &corestr.Collection{}
		_ = c.UnmarshalJSON([]byte(`["a","b"]`))
		_ = c.UnmarshalJSON([]byte(`invalid`))
	})
}

func Test_C30_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ParseInjectUsingJson", func() {
		c := corestr.New.Collection.Empty()
		r := corejson.New([]string{"a"})
		_, _ = c.ParseInjectUsingJson(&r)
	})
}

func Test_C30_Collection_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ParseInjectUsingJson_Error", func() {
		c := corestr.New.Collection.Empty()
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := c.ParseInjectUsingJson(bad)
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_C30_Collection_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_C30_Collection_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		c := corestr.New.Collection.Empty()
		bad := corejson.NewResult.UsingString(`invalid`)
		c.ParseInjectUsingJsonMust(bad)
	})
}

func Test_C30_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C30_Collection_JsonParseSelfInject", func() {
		c := corestr.New.Collection.Empty()
		r := corejson.New([]string{"a"})
		_ = c.JsonParseSelfInject(&r)
	})
}

func Test_C30_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()
		var c2 *corestr.Collection
		c2.Clear()
	})
}

func Test_C30_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		var c2 *corestr.Collection
		c2.Dispose()
	})
}

func Test_C30_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		var target []string
		_ = c.Deserialize(&target)
	})
}

func Test_C30_Collection_Join(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		_ = c.Join(",")
		c2 := corestr.New.Collection.Empty()
		_ = c2.Join(",")
	})
}

func Test_C30_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_C30_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.JoinLine()
		c2 := corestr.New.Collection.Empty()
		_ = c2.JoinLine()
	})
}

func Test_C30_Collection_Single(t *testing.T) {
	safeTest(t, "Test_C30_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.Single()
	})
}

func Test_C30_Collection_List(t *testing.T) {
	safeTest(t, "Test_C30_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.List()
	})
}

// ── newCollectionCreator ──

func Test_C30_NCC_Empty(t *testing.T) { _ = corestr.New.Collection.Empty() }
func Test_C30_NCC_Cap(t *testing.T) { _ = corestr.New.Collection.Cap(5) }
func Test_C30_NCC_CloneStrings(t *testing.T) { _ = corestr.New.Collection.CloneStrings([]string{"a"}) }
func Test_C30_NCC_Create(t *testing.T) { _ = corestr.New.Collection.Create([]string{"a"}) }
func Test_C30_NCC_Strings(t *testing.T) { _ = corestr.New.Collection.Strings([]string{"a"}) }
func Test_C30_NCC_StringsOptions(t *testing.T) {
	safeTest(t, "Test_C30_NCC_StringsOptions", func() {
		_ = corestr.New.Collection.StringsOptions(true, []string{"a"})
		_ = corestr.New.Collection.StringsOptions(false, []string{})
		_ = corestr.New.Collection.StringsOptions(false, []string{"a"})
	})
}
func Test_C30_NCC_LineUsingSep(t *testing.T) { _ = corestr.New.Collection.LineUsingSep(",", "a,b") }
func Test_C30_NCC_LineDefault(t *testing.T) { _ = corestr.New.Collection.LineDefault("a|b") }
func Test_C30_NCC_StringsPlusCap(t *testing.T) {
	safeTest(t, "Test_C30_NCC_StringsPlusCap", func() {
		_ = corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		_ = corestr.New.Collection.StringsPlusCap(0, []string{"a"})
	})
}
func Test_C30_NCC_CapStrings(t *testing.T) {
	safeTest(t, "Test_C30_NCC_CapStrings", func() {
		_ = corestr.New.Collection.CapStrings(5, []string{"a"})
		_ = corestr.New.Collection.CapStrings(0, []string{"a"})
	})
}
func Test_C30_NCC_LenCap(t *testing.T) { _ = corestr.New.Collection.LenCap(0, 5) }

// ── Collection AddStringsAsync ──

func Test_C30_Collection_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddStringsAsync", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		c.AddStringsAsync(wg, []string{"a", "b"})
		wg.Wait()
		c.AddStringsAsync(wg, []string{})
	})
}

// ── Collection AddsAsync ──

func Test_C30_Collection_AddsAsync(t *testing.T) {
	safeTest(t, "Test_C30_Collection_AddsAsync", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddsAsync(wg, "a", "b")
		wg.Wait()
		c.AddsAsync(wg, nil...)
	})
}
