package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ── Collection CRUD ──

func Test_C37_Collection_Add(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Add", func() {
		c := corestr.New.Collection.Empty()
		c.Add("a")
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmpty("")
		c.AddNonEmpty("a")
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyWhitespace("   ")
		c.AddNonEmptyWhitespace("a")
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddIf", func() {
		c := corestr.New.Collection.Empty()
		c.AddIf(false, "skip")
		c.AddIf(true, "keep")
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Empty()
		c.AddIfMany(false, "a", "b")
		c.AddIfMany(true, "c", "d")
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddError", func() {
		c := corestr.New.Collection.Empty()
		c.AddError(nil)
		if c.Length() != 0 { t.Fatal("expected 0") }
		c.AddError(errForTest)
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_Adds(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Adds", func() {
		c := corestr.New.Collection.Empty()
		c.Adds("a", "b", "c")
		if c.Length() != 3 { t.Fatal("expected 3") }
	})
}

func Test_C37_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddStrings([]string{"x", "y"})
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddLock("a")
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddsLock", func() {
		c := corestr.New.Collection.Empty()
		c.AddsLock("a", "b")
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddFunc", func() {
		c := corestr.New.Collection.Empty()
		c.AddFunc(func() string { return "hello" })
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddFuncErr", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AddFuncErr(func() (string, error) { return "", errForTest }, func(e error) {})
		if c.Length() != 1 { t.Fatal("expected 1 still") }
	})
}

func Test_C37_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddCollection", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c1.AddCollection(c2)
		if c1.Length() != 2 { t.Fatal("expected 2") }
		c1.AddCollection(corestr.Empty.Collection())
		if c1.Length() != 2 { t.Fatal("expected 2 still") }
	})
}

func Test_C37_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddCollections", func() {
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AddCollections(c1, c2, corestr.Empty.Collection())
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c1)
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddHashmapsValues(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddHashmapsValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("k", "v")
		c.AddHashmapsValues(h)
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AddHashmapsValues(nil)
		if c.Length() != 1 { t.Fatal("expected 1 after nil") }
	})
}

func Test_C37_Collection_AddHashmapsKeys(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddHashmapsKeys", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("mykey", "myval")
		c.AddHashmapsKeys(h)
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AddHashmapsKeys(nil)
	})
}

func Test_C37_Collection_AddHashmapsKeysValues(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddHashmapsKeysValues", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValues(h)
		if c.Length() != 2 { t.Fatal("expected 2") }
		c.AddHashmapsKeysValues(nil)
	})
}

// ── Collection query methods ──

func Test_C37_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasAnyItem() { t.Fatal("expected true") }
	})
}

func Test_C37_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.HasIndex(0) { t.Fatal("expected true") }
		if !c.HasIndex(1) { t.Fatal("expected true") }
		if c.HasIndex(2) { t.Fatal("expected false") }
		if c.HasIndex(-1) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_C37_Collection_LastIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.LastIndex() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		if c.Capacity() < 10 { t.Fatal("expected >= 10") }
		empty := &corestr.Collection{}
		_ = empty.Capacity()
	})
}

func Test_C37_Collection_Count(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Count", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.Count() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Length_Nil", func() {
		var c *corestr.Collection
		if c.Length() != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.LengthLock() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEmpty", func() {
		if !corestr.Empty.Collection().IsEmpty() { t.Fatal("expected empty") }
	})
}

func Test_C37_Collection_IsEmptyLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEmptyLock", func() {
		if !corestr.Empty.Collection().IsEmptyLock() { t.Fatal("expected empty") }
	})
}

func Test_C37_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasItems() { t.Fatal("expected has items") }
	})
}

func Test_C37_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_C37_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.RemoveAt(1) { t.Fatal("expected true") }
		if c.Length() != 2 { t.Fatal("expected 2") }
		if c.RemoveAt(-1) { t.Fatal("expected false") }
		if c.RemoveAt(100) { t.Fatal("expected false") }
	})
}

// ── Collection equality ──

func Test_C37_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEquals", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		if !a.IsEquals(b) { t.Fatal("expected equal") }
	})
}

func Test_C37_Collection_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEquals_Nil", func() {
		var a *corestr.Collection
		var b *corestr.Collection
		if !a.IsEquals(b) { t.Fatal("expected true nil==nil") }
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.IsEquals(nil) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_IsEquals_Self(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEquals_Self", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		if !a.IsEquals(a) { t.Fatal("expected true for self") }
	})
}

func Test_C37_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEquals_BothEmpty", func() {
		a := corestr.Empty.Collection()
		b := corestr.Empty.Collection()
		if !a.IsEquals(b) { t.Fatal("expected true") }
	})
}

func Test_C37_Collection_IsEquals_DiffLen(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEquals_DiffLen", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		if a.IsEquals(b) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_IsEqualsWithSensitive_Insensitive(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEqualsWithSensitive_Insensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})
		if !a.IsEqualsWithSensitive(false, b) { t.Fatal("expected equal insensitive") }
		if a.IsEqualsWithSensitive(true, b) { t.Fatal("expected not equal sensitive") }
	})
}

func Test_C37_Collection_IsEqualsWithSensitive_Mismatch(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsEqualsWithSensitive_Mismatch", func() {
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"World"})
		if a.IsEqualsWithSensitive(false, b) { t.Fatal("expected not equal") }
	})
}

// ── Collection accessors ──

func Test_C37_Collection_First(t *testing.T) {
	safeTest(t, "Test_C37_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.First() != "a" { t.Fatal("expected a") }
	})
}

func Test_C37_Collection_Last(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Last() != "b" { t.Fatal("expected b") }
	})
}

func Test_C37_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_C37_Collection_FirstOrDefault", func() {
		if corestr.Empty.Collection().FirstOrDefault() != "" { t.Fatal("expected empty") }
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.FirstOrDefault() != "a" { t.Fatal("expected a") }
	})
}

func Test_C37_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_C37_Collection_LastOrDefault", func() {
		if corestr.Empty.Collection().LastOrDefault() != "" { t.Fatal("expected empty") }
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.LastOrDefault() != "a" { t.Fatal("expected a") }
	})
}

func Test_C37_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.IndexAt(1) != "b" { t.Fatal("expected b") }
	})
}

func Test_C37_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.SafeIndexAtUsingLength("def", 1, 0) != "a" { t.Fatal("expected a") }
		if c.SafeIndexAtUsingLength("def", 1, 5) != "def" { t.Fatal("expected def") }
	})
}

func Test_C37_Collection_List_Items_ListStrings(t *testing.T) {
	safeTest(t, "Test_C37_Collection_List_Items_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.List()) != 1 { t.Fatal() }
		if len(c.Items()) != 1 { t.Fatal() }
		if len(c.ListStrings()) != 1 { t.Fatal() }
		if len(c.ListStringsPtr()) != 1 { t.Fatal() }
		if len(c.ListPtr()) != 1 { t.Fatal() }
	})
}

// ── Collection transform ──

func Test_C37_Collection_Take(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)
		if taken.Length() != 2 { t.Fatal("expected 2") }
		same := c.Take(10)
		if same != c { t.Fatal("expected same ptr") }
		empty := c.Take(0)
		if empty.Length() != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)
		if skipped.Length() != 2 { t.Fatal("expected 2") }
		same := c.Skip(0)
		if same != c { t.Fatal("expected same ptr") }
	})
}

func Test_C37_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		if c.First() != "c" { t.Fatal("expected c") }
		// Two items
		c2 := corestr.New.Collection.Strings([]string{"x", "y"})
		c2.Reverse()
		if c2.First() != "y" { t.Fatal("expected y") }
		// Single
		c3 := corestr.New.Collection.Strings([]string{"z"})
		c3.Reverse()
		if c3.First() != "z" { t.Fatal("expected z") }
	})
}

func Test_C37_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		sorted := c.SortedListAsc()
		if sorted[0] != "a" { t.Fatal("expected a first") }
		if corestr.Empty.Collection().SortedListAsc() == nil { t.Fatal("unexpected nil") }
	})
}

func Test_C37_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAsc()
		if c.First() != "a" { t.Fatal("expected a") }
		corestr.Empty.Collection().SortedAsc()
	})
}

func Test_C37_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()
		if c.First() != "a" { t.Fatal("expected a") }
	})
}

func Test_C37_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		dsc := c.SortedListDsc()
		if dsc[0] != "c" { t.Fatal("expected c") }
	})
}

func Test_C37_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_C37_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		u := c.UniqueList()
		if len(u) != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		u := c.UniqueListLock()
		if len(u) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_C37_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		m := c.UniqueBoolMap()
		if len(m) != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		m := c.UniqueBoolMapLock()
		if len(m) != 1 { t.Fatal("expected 1") }
	})
}

// ── Collection filter ──

func Test_C37_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if len(result) != 2 { t.Fatal("expected 2") }
		// empty
		result2 := corestr.Empty.Collection().Filter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result2) != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Filter_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})
		if len(result) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result) != 2 { t.Fatal("expected 2") }
		// empty
		result2 := corestr.Empty.Collection().FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result2) != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_C37_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, s == "a", false
		})
		if fc.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_FilteredCollectionLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if fc.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 2 { t.Fatal("expected 2") }
		// empty
		result2 := corestr.Empty.Collection().FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result2) != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 1 { t.Fatal("expected 1") }
		// empty
		result2 := corestr.Empty.Collection().FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result2) != 0 { t.Fatal("expected 0") }
	})
}

// ── Collection search ──

func Test_C37_Collection_Has(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.Has("a") { t.Fatal("expected true") }
		if c.Has("z") { t.Fatal("expected false") }
		if corestr.Empty.Collection().Has("a") { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasLock("a") { t.Fatal("expected true") }
	})
}

func Test_C37_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		if !c.HasPtr(&s) { t.Fatal("expected true") }
		if c.HasPtr(nil) { t.Fatal("expected false") }
		if corestr.Empty.Collection().HasPtr(&s) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.HasAll("a", "b") { t.Fatal("expected true") }
		if c.HasAll("a", "z") { t.Fatal("expected false") }
		if corestr.Empty.Collection().HasAll("a") { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		if !c.HasUsingSensitivity("hello", false) { t.Fatal("expected true") }
		if c.HasUsingSensitivity("hello", true) { t.Fatal("expected false") }
		if c.HasUsingSensitivity("world", false) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsContainsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		if !c.IsContainsPtr(&s) { t.Fatal("expected true") }
		if c.IsContainsPtr(nil) { t.Fatal("expected false") }
		miss := "z"
		if c.IsContainsPtr(&miss) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAll("a", "b") { t.Fatal("expected true") }
		if c.IsContainsAll(nil...) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAllSlice([]string{"a"}) { t.Fatal("expected true") }
		if c.IsContainsAllSlice([]string{}) { t.Fatal("expected false on empty") }
		if corestr.Empty.Collection().IsContainsAllSlice([]string{"a"}) { t.Fatal("expected false") }
	})
}

func Test_C37_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.IsContainsAllLock("a") { t.Fatal("expected true") }
		if c.IsContainsAllLock(nil...) { t.Fatal("expected false nil") }
	})
}

func Test_C37_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_C37_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, ok := c.GetHashsetPlusHasAll([]string{"a", "b"})
		if !ok { t.Fatal("expected true") }
		if hs.Length() != 2 { t.Fatal("expected 2") }
		_, ok2 := c.GetHashsetPlusHasAll(nil)
		if ok2 { t.Fatal("expected false") }
	})
}

// ── Collection string ops ──

func Test_C37_Collection_String(t *testing.T) {
	safeTest(t, "Test_C37_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.String() == "" { t.Fatal("expected non-empty") }
		if corestr.Empty.Collection().String() == "" { t.Fatal("should have no elements text") }
	})
}

func Test_C37_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.StringLock() == "" { t.Fatal("expected non-empty") }
	})
}

func Test_C37_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_C37_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.StringJSON() == "" { t.Fatal("expected non-empty") }
	})
}

func Test_C37_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_C37_Collection_JsonString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.JsonString() == "" { t.Fatal() }
		if c.JsonStringMust() == "" { t.Fatal() }
	})
}

func Test_C37_Collection_Join(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Join(",") != "a,b" { t.Fatal("expected a,b") }
		if corestr.Empty.Collection().Join(",") != "" { t.Fatal("expected empty") }
	})
}

func Test_C37_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_C37_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.JoinLine()
		if r == "" { t.Fatal("expected non-empty") }
		if corestr.Empty.Collection().JoinLine() != "" { t.Fatal("expected empty") }
	})
}

func Test_C37_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.Joins(",") == "" { t.Fatal() }
		if c.Joins(",", "b") == "" { t.Fatal() }
	})
}

func Test_C37_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyJoins(",")
		if r == "" { t.Fatal("expected non-empty") }
	})
}

func Test_C37_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		r := c.NonWhitespaceJoins(",")
		if r == "" { t.Fatal("expected non-empty") }
	})
}

// ── Collection CSV ──

func Test_C37_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Csv() == "" { t.Fatal() }
		if corestr.Empty.Collection().Csv() != "" { t.Fatal() }
	})
}

func Test_C37_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_C37_Collection_CsvOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.CsvOptions(true) == "" { t.Fatal() }
		if corestr.Empty.Collection().CsvOptions(false) != "" { t.Fatal() }
	})
}

func Test_C37_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_C37_Collection_CsvLines", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		lines := c.CsvLines()
		if len(lines) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_C37_Collection_CsvLinesOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		lines := c.CsvLinesOptions(true)
		if len(lines) != 1 { t.Fatal("expected 1") }
	})
}

// ── Collection pages ──

func Test_C37_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_C37_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		if c.GetPagesSize(2) != 3 { t.Fatal("expected 3") }
		if c.GetPagesSize(0) != 0 { t.Fatal("expected 0") }
		if c.GetPagesSize(-1) != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_C37_Collection_GetPagedCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		paged := c.GetPagedCollection(2)
		if paged.Length() < 1 { t.Fatal("expected at least 1 page") }
	})
}

func Test_C37_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_C37_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		page := c.GetSinglePageCollection(2, 1)
		if page.Length() != 2 { t.Fatal("expected 2") }
		page2 := c.GetSinglePageCollection(2, 2)
		if page2.Length() != 2 { t.Fatal("expected 2") }
		// small collection
		small := corestr.New.Collection.Strings([]string{"a"})
		same := small.GetSinglePageCollection(10, 1)
		if same != small { t.Fatal("expected same ptr") }
	})
}

// ── Collection hashset/map ──

func Test_C37_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()
		if hs.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.HashsetWithDoubleLength()
	})
}

func Test_C37_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()
		if hs.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_C37_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"abc", "axy", "bcd"})
		ccm := c.CharCollectionMap()
		if ccm.Length() < 2 { t.Fatal("expected >= 2 groups") }
	})
}

// ── Collection non-empty ──

func Test_C37_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		ne := c.NonEmptyList()
		if len(ne) != 2 { t.Fatal("expected 2") }
		if corestr.Empty.Collection().NonEmptyList() == nil { t.Fatal() }
	})
}

func Test_C37_Collection_NonEmptyListPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyListPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", ""})
		ne := c.NonEmptyListPtr()
		if len(*ne) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"a", ""})
		if len(c.NonEmptyItems()) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_NonEmptyItemsPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyItemsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", ""})
		_ = c.NonEmptyItemsPtr()
	})
}

func Test_C37_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  "})
		r := c.NonEmptyItemsOrNonWhitespace()
		if len(r) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_NonEmptyItemsOrNonWhitespacePtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_NonEmptyItemsOrNonWhitespacePtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_ = c.NonEmptyItemsOrNonWhitespacePtr()
	})
}

// ── Collection add non-empty ──

func Test_C37_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")
		if c.Length() != 2 { t.Fatal("expected 2") }
		c.AddsNonEmpty()
	})
}

func Test_C37_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Empty()
		s := "hello"
		c.AddsNonEmptyPtrLock(&s, nil)
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "b")
		if c.Length() != 2 { t.Fatal("expected 2") }
		c.AddNonEmptyStrings()
	})
}

func Test_C37_Collection_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddNonEmptyStringsSlice", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a"})
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AddNonEmptyStringsSlice(nil)
	})
}

// ── Collection misc ──

func Test_C37_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_C37_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a,b", "c"})
		split := c.EachItemSplitBy(",")
		if len(split) != 3 { t.Fatal("expected 3") }
	})
}

func Test_C37_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		cn := c.ConcatNew(0, "b")
		if cn.Length() != 2 { t.Fatal("expected 2") }
		cn2 := c.ConcatNew(0)
		if cn2.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_ToError(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ToError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		if c.ToError(",") == nil { t.Fatal("expected non-nil") }
	})
}

func Test_C37_Collection_ToDefaultError(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ToDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err1"})
		if c.ToDefaultError() == nil { t.Fatal("expected non-nil") }
	})
}

func Test_C37_Collection_AsError(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AsError", func() {
		c := corestr.New.Collection.Strings([]string{"err"})
		if c.AsError(",") == nil { t.Fatal("expected error") }
		if corestr.Empty.Collection().AsError(",") != nil { t.Fatal("expected nil") }
	})
}

func Test_C37_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"e"})
		if c.AsDefaultError() == nil { t.Fatal("expected error") }
	})
}

func Test_C37_Collection_InsertAt(t *testing.T) {
	safeTest(t, "Test_C37_Collection_InsertAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c"})
		c.InsertAt(0, "b")
		if c.Length() < 3 { t.Fatal("expected >= 3") }
	})
}

func Test_C37_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.ChainRemoveAt(1)
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_C37_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)
		if c.Length() != 2 { t.Fatal("expected 2") }
		// nil indexes
		c.RemoveItemsIndexes(true)
	})
}

func Test_C37_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AppendCollectionPtr(c2)
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Empty()
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections(c1, corestr.Empty.Collection())
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AppendCollections()
	})
}

func Test_C37_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys(42, "hello", nil)
		if c.Length() != 2 { t.Fatal("expected 2") }
		c.AppendAnys()
	})
}

func Test_C37_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock(42)
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AppendAnysLock()
	})
}

func Test_C37_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys(42, nil)
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AppendNonEmptyAnys(nil)
	})
}

func Test_C37_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendAnysUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, false },
			42,
		)
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AppendAnysUsingFilter(
			func(s string, i int) (string, bool, bool) { return s, true, false },
		)
	})
}

func Test_C37_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AppendAnysUsingFilterLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysUsingFilterLock(
			func(s string, i int) (string, bool, bool) { return s, true, false },
			42,
		)
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddsAsync(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddsAsync", func() {
		c := corestr.New.Collection.Cap(10)
		wg := sync.WaitGroup{}
		wg.Add(1)
		c.AddsAsync(&wg, "a", "b")
		wg.Wait()
		// may or may not be 2 due to race, but should not panic
	})
}

func Test_C37_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Cap(2)
		wg := sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(&wg, "a")
		wg.Wait()
		if c.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking(
			[]string{"a", "", "b"},
			func(s string) bool { return s != "" },
		)
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd(
			[]string{"a,b", "c"},
			func(s string) []string { return []string{s} },
		)
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_C37_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		if c.Length() != 2 { t.Fatal("expected 2") }
	})
}

func Test_C37_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_C37_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.GetAllExcept([]string{"b"})
		if len(r) != 2 { t.Fatal("expected 2") }
		r2 := c.GetAllExcept(nil)
		if len(r2) != 3 { t.Fatal("expected 3") }
	})
}

func Test_C37_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_C37_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExceptCollection(nil)
		if len(r) != 2 { t.Fatal("expected 2") }
		exc := corestr.New.Collection.Strings([]string{"a"})
		r2 := c.GetAllExceptCollection(exc)
		if len(r2) != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "x" })
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AddFuncResult()
	})
}

func Test_C37_Collection_New(t *testing.T) {
	safeTest(t, "Test_C37_Collection_New", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		n := c.New("x", "y")
		if n.Length() != 2 { t.Fatal("expected 2") }
		n2 := c.New()
		if n2.Length() != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		l := c.ListCopyPtrLock()
		if len(l) != 1 { t.Fatal("expected 1") }
		l2 := corestr.Empty.Collection().ListCopyPtrLock()
		if len(l2) != 0 { t.Fatal("expected 0") }
	})
}

func Test_C37_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryString(1)
		if s == "" { t.Fatal("expected non-empty") }
	})
}

func Test_C37_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_C37_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.SummaryStringWithHeader("H") == "" { t.Fatal() }
		if corestr.Empty.Collection().SummaryStringWithHeader("H") == "" { t.Fatal() }
	})
}

// ── Collection JSON ──

func Test_C37_Collection_Json(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		j := c.Json()
		if j.HasError() { t.Fatal(j.Error) }
	})
}

func Test_C37_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C37_Collection_JsonPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jp := c.JsonPtr()
		if jp.HasError() { t.Fatal(jp.Error) }
	})
}

func Test_C37_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_C37_Collection_JsonModel", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.JsonModel()) != 1 { t.Fatal() }
	})
}

func Test_C37_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_C37_Collection_JsonModelAny", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.JsonModelAny() == nil { t.Fatal() }
	})
}

func Test_C37_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_C37_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.MarshalJSON()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal("expected bytes") }
	})
}

func Test_C37_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_C37_Collection_UnmarshalJSON", func() {
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil { t.Fatal(err) }
		if c.Length() != 2 { t.Fatal("expected 2") }
		// invalid
		err2 := c.UnmarshalJSON([]byte(`{invalid`))
		if err2 == nil { t.Fatal("expected error") }
	})
}

func Test_C37_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ParseInjectUsingJson", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := &corestr.Collection{}
		result, err := c2.ParseInjectUsingJson(jr)
		if err != nil { t.Fatal(err) }
		if result.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_C37_Collection_ParseInjectUsingJsonMust", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := &corestr.Collection{}
		result := c2.ParseInjectUsingJsonMust(jr)
		if result.Length() != 1 { t.Fatal("expected 1") }
	})
}

func Test_C37_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_C37_Collection_JsonParseSelfInject", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := &corestr.Collection{}
		err := c2.JsonParseSelfInject(jr)
		if err != nil { t.Fatal(err) }
	})
}

func Test_C37_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()
		if err != nil { t.Fatal(err) }
		if len(b) == 0 { t.Fatal("expected bytes") }
	})
}

func Test_C37_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		var target []string
		err := c.Deserialize(&target)
		if err != nil { t.Fatal(err) }
	})
}

	// ── Collection interface casts ──

func Test_C37_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.AsJsonMarshaller() == nil { t.Fatal() }
	})
}

func Test_C37_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.AsJsonContractsBinder() == nil { t.Fatal() }
	})
}

	// ── Collection resize/capacity ──

func Test_C37_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(10)
		if c.Capacity() < 12 { t.Fatal("expected >= 12") }
		c.AddCapacity()
	})
}

func Test_C37_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Resize", func() {
		c := corestr.New.Collection.Cap(2)
		c.Resize(100)
		if c.Capacity() < 100 { t.Fatal() }
		c.Resize(1) // no-op
	})
}

	// ── Collection clear/dispose ──

func Test_C37_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Clear()
		if c.Length() != 0 { t.Fatal("expected 0") }
		var nilC *corestr.Collection
		if nilC.Clear() != nil { t.Fatal("expected nil") }
	})
}

func Test_C37_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_C37_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		if c.Length() != 0 { t.Fatal("expected 0") }
		var nilC *corestr.Collection
		nilC.Dispose() // should not panic
	})
}

	// ── Collection Hashmap filter ──

func Test_C37_Collection_AddHashmapsKeysValuesUsingFilter(t *testing.T) {
	safeTest(t, "Test_C37_Collection_AddHashmapsKeysValuesUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		h := corestr.New.Hashmap.Cap(2)
		h.AddOrUpdate("k", "v")
		c.AddHashmapsKeysValuesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) {
				return p.Key + "=" + p.Value, true, false
			},
			h,
		)
		if c.Length() != 1 { t.Fatal("expected 1") }
		c.AddHashmapsKeysValuesUsingFilter(
			func(p corestr.KeyValuePair) (string, bool, bool) { return "", true, false },
			nil,
		)
	})
}

	// ── newCollectionCreator ──

func Test_C37_NewCollectionCreator_Methods(t *testing.T) {
	safeTest(t, "Test_C37_NewCollectionCreator_Methods", func() {
		// Cap
		c1 := corestr.New.Collection.Cap(5)
		if c1.Capacity() < 5 { t.Fatal() }
		// CloneStrings
		c2 := corestr.New.Collection.CloneStrings([]string{"a"})
		if c2.Length() != 1 { t.Fatal() }
		// Create
		c3 := corestr.New.Collection.Create([]string{"a"})
		if c3.Length() != 1 { t.Fatal() }
		// StringsOptions clone
		c4 := corestr.New.Collection.StringsOptions(true, []string{"a"})
		if c4.Length() != 1 { t.Fatal() }
		// StringsOptions no clone empty
		c5 := corestr.New.Collection.StringsOptions(false, nil)
		if c5.Length() != 0 { t.Fatal() }
		// LineUsingSep
		c6 := corestr.New.Collection.LineUsingSep(",", "a,b,c")
		if c6.Length() != 3 { t.Fatal() }
		// LineDefault
		c7 := corestr.New.Collection.LineDefault("a\nb")
		if c7.Length() != 2 { t.Fatal() }
		// StringsPlusCap
		c8 := corestr.New.Collection.StringsPlusCap(5, []string{"a"})
		if c8.Length() != 1 { t.Fatal() }
		c9 := corestr.New.Collection.StringsPlusCap(0, []string{"a"})
		if c9.Length() != 1 { t.Fatal() }
		// CapStrings
		c10 := corestr.New.Collection.CapStrings(5, []string{"a"})
		if c10.Length() != 1 { t.Fatal() }
		c11 := corestr.New.Collection.CapStrings(0, []string{"a"})
		if c11.Length() != 1 { t.Fatal() }
		// LenCap
		c12 := corestr.New.Collection.LenCap(3, 10)
		if c12.Length() != 3 { t.Fatal() }
		// Empty
		c13 := corestr.New.Collection.Empty()
		if c13.Length() != 0 { t.Fatal() }
	})
}

	// ── emptyCreator ──

func Test_C37_EmptyCreator_All(t *testing.T) {
	safeTest(t, "Test_C37_EmptyCreator_All", func() {
		if corestr.Empty.Collection().Length() != 0 { t.Fatal() }
		if corestr.Empty.LinkedList().Length() != 0 { t.Fatal() }
		if corestr.Empty.SimpleSlice().Length() != 0 { t.Fatal() }
		if corestr.Empty.KeyAnyValuePair() == nil { t.Fatal() }
		if corestr.Empty.KeyValuePair() == nil { t.Fatal() }
		if corestr.Empty.KeyValueCollection().Length() != 0 { t.Fatal() }
		if corestr.Empty.LinkedCollections().Length() != 0 { t.Fatal() }
		if corestr.Empty.LeftRight() == nil { t.Fatal() }
		_ = corestr.Empty.SimpleStringOnce()
		if corestr.Empty.SimpleStringOncePtr() == nil { t.Fatal() }
		if corestr.Empty.Hashset().Length() != 0 { t.Fatal() }
		if corestr.Empty.HashsetsCollection().Length() != 0 { t.Fatal() }
		if corestr.Empty.Hashmap().Length() != 0 { t.Fatal() }
		if corestr.Empty.CharCollectionMap().Length() != 0 { t.Fatal() }
		if corestr.Empty.KeyValuesCollection().Length() != 0 { t.Fatal() }
		if corestr.Empty.CollectionsOfCollection().Length() != 0 { t.Fatal() }
		if corestr.Empty.CharHashsetMap().Length() != 0 { t.Fatal() }
	})
}
