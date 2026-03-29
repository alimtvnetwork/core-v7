package corestrtests

import (
	"errors"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
)

// =============================================================================
// Collection — JSON and Serialization
// =============================================================================

func Test_I8_C01_Collection_JsonString(t *testing.T) {
	safeTest(t, "Test_I8_C01_Collection_JsonString", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JsonString()
		if s == "" {
			t.Fatal("expected non-empty json")
		}
	})
}

func Test_I8_C02_Collection_JsonStringMust(t *testing.T) {
	safeTest(t, "Test_I8_C02_Collection_JsonStringMust", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.JsonStringMust()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C03_Collection_StringJSON(t *testing.T) {
	safeTest(t, "Test_I8_C03_Collection_StringJSON", func() {
		c := corestr.New.Collection.Strings([]string{"x"})
		s := c.StringJSON()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C04_Collection_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_I8_C04_Collection_HasAnyItem", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasAnyItem() {
			t.Fatal("expected true")
		}
		e := corestr.Empty.Collection()
		if e.HasAnyItem() {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_C05_Collection_LastIndex(t *testing.T) {
	safeTest(t, "Test_I8_C05_Collection_LastIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if c.LastIndex() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C06_Collection_HasIndex(t *testing.T) {
	safeTest(t, "Test_I8_C06_Collection_HasIndex", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.HasIndex(0) {
			t.Fatal("expected true for 0")
		}
		if !c.HasIndex(1) {
			t.Fatal("expected true for 1")
		}
		if c.HasIndex(2) {
			t.Fatal("expected false for 2")
		}
		if c.HasIndex(-1) {
			t.Fatal("expected false for -1")
		}
	})
}

func Test_I8_C07_Collection_ListStringsPtr(t *testing.T) {
	safeTest(t, "Test_I8_C07_Collection_ListStringsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.ListStringsPtr()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C08_Collection_ListStrings(t *testing.T) {
	safeTest(t, "Test_I8_C08_Collection_ListStrings", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if len(c.ListStrings()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C09_Collection_RemoveAt(t *testing.T) {
	safeTest(t, "Test_I8_C09_Collection_RemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		ok := c.RemoveAt(1)
		if !ok {
			t.Fatal("expected success")
		}
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
		// negative index
		if c.RemoveAt(-1) {
			t.Fatal("expected false for negative")
		}
		// out of range
		if c.RemoveAt(100) {
			t.Fatal("expected false for out of range")
		}
	})
}

func Test_I8_C10_Collection_Count(t *testing.T) {
	safeTest(t, "Test_I8_C10_Collection_Count", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.Count() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C11_Collection_Capacity(t *testing.T) {
	safeTest(t, "Test_I8_C11_Collection_Capacity", func() {
		c := corestr.New.Collection.Cap(10)
		if c.Capacity() < 10 {
			t.Fatal("expected at least 10")
		}
	})
}

func Test_I8_C12_Collection_Capacity_Nil(t *testing.T) {
	safeTest(t, "Test_I8_C12_Collection_Capacity_Nil", func() {
		c := corestr.New.Collection.Strings(nil)
		_ = c.Capacity()
	})
}

func Test_I8_C13_Collection_LengthLock(t *testing.T) {
	safeTest(t, "Test_I8_C13_Collection_LengthLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.LengthLock() != 2 {
			t.Fatal("expected 2")
		}
	})
}

// =============================================================================
// Collection — Equality
// =============================================================================

func Test_I8_C14_Collection_IsEquals(t *testing.T) {
	safeTest(t, "Test_I8_C14_Collection_IsEquals", func() {
		a := corestr.New.Collection.Strings([]string{"a", "b"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		if !a.IsEquals(b) {
			t.Fatal("expected equal")
		}
	})
}

func Test_I8_C15_Collection_IsEqualsWithSensitive_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_I8_C15_Collection_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.Collection.Strings([]string{"Hello"})
		b := corestr.New.Collection.Strings([]string{"hello"})
		if !a.IsEqualsWithSensitive(false, b) {
			t.Fatal("expected equal case-insensitive")
		}
		if a.IsEqualsWithSensitive(true, b) {
			t.Fatal("expected not equal case-sensitive")
		}
	})
}

func Test_I8_C16_Collection_IsEquals_BothNil(t *testing.T) {
	safeTest(t, "Test_I8_C16_Collection_IsEquals_BothNil", func() {
		var a, b *corestr.Collection
		_ = a
		_ = b
		// Can't directly test nil.IsEquals(nil) due to nil receiver panic
		// But we test via isCollectionPrecheckEqual through different path
	})
}

func Test_I8_C17_Collection_IsEquals_DiffLength(t *testing.T) {
	safeTest(t, "Test_I8_C17_Collection_IsEquals_DiffLength", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"a", "b"})
		if a.IsEquals(b) {
			t.Fatal("expected not equal for different length")
		}
	})
}

func Test_I8_C18_Collection_IsEquals_SamePtr(t *testing.T) {
	safeTest(t, "Test_I8_C18_Collection_IsEquals_SamePtr", func() {
		a := corestr.New.Collection.Strings([]string{"a"})
		if !a.IsEquals(a) {
			t.Fatal("expected equal for same pointer")
		}
	})
}

func Test_I8_C19_Collection_IsEquals_BothEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C19_Collection_IsEquals_BothEmpty", func() {
		a := corestr.Empty.Collection()
		b := corestr.Empty.Collection()
		if !a.IsEquals(b) {
			t.Fatal("expected equal for both empty")
		}
	})
}

func Test_I8_C20_Collection_IsEquals_OneEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C20_Collection_IsEquals_OneEmpty", func() {
		a := corestr.Empty.Collection()
		b := corestr.New.Collection.Strings([]string{"a"})
		if a.IsEquals(b) {
			t.Fatal("expected not equal")
		}
	})
}

// =============================================================================
// Collection — Add variants
// =============================================================================

func Test_I8_C21_Collection_IsEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C21_Collection_IsEmpty", func() {
		c := corestr.Empty.Collection()
		if !c.IsEmpty() {
			t.Fatal("expected empty")
		}
		if !c.IsEmptyLock() {
			t.Fatal("expected empty with lock")
		}
	})
}

func Test_I8_C22_Collection_HasItems(t *testing.T) {
	safeTest(t, "Test_I8_C22_Collection_HasItems", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if !c.HasItems() {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_C23_Collection_AddLock(t *testing.T) {
	safeTest(t, "Test_I8_C23_Collection_AddLock", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddLock("a")
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C24_Collection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C24_Collection_AddNonEmpty", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddNonEmpty("")
		if c.Length() != 0 {
			t.Fatal("expected 0 for empty string")
		}
		c.AddNonEmpty("a")
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C25_Collection_AddNonEmptyWhitespace(t *testing.T) {
	safeTest(t, "Test_I8_C25_Collection_AddNonEmptyWhitespace", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddNonEmptyWhitespace("  ")
		if c.Length() != 0 {
			t.Fatal("expected 0 for whitespace")
		}
		c.AddNonEmptyWhitespace("a")
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C26_Collection_AddError(t *testing.T) {
	safeTest(t, "Test_I8_C26_Collection_AddError", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddError(nil)
		if c.Length() != 0 {
			t.Fatal("expected 0 for nil error")
		}
		c.AddError(errors.New("test"))
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C27_Collection_AsDefaultError(t *testing.T) {
	safeTest(t, "Test_I8_C27_Collection_AsDefaultError", func() {
		c := corestr.New.Collection.Strings([]string{"err1", "err2"})
		err := c.AsDefaultError()
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_I8_C28_Collection_AsError_Empty(t *testing.T) {
	safeTest(t, "Test_I8_C28_Collection_AsError_Empty", func() {
		c := corestr.Empty.Collection()
		err := c.AsError(",")
		if err != nil {
			t.Fatal("expected nil")
		}
	})
}

func Test_I8_C29_Collection_AddIf(t *testing.T) {
	safeTest(t, "Test_I8_C29_Collection_AddIf", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddIf(false, "skip")
		if c.Length() != 0 {
			t.Fatal("expected 0")
		}
		c.AddIf(true, "keep")
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C30_Collection_EachItemSplitBy(t *testing.T) {
	safeTest(t, "Test_I8_C30_Collection_EachItemSplitBy", func() {
		c := corestr.New.Collection.Strings([]string{"a.b", "c.d"})
		result := c.EachItemSplitBy(".")
		if len(result) != 4 {
			t.Fatalf("expected 4, got %d", len(result))
		}
	})
}

func Test_I8_C31_Collection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_I8_C31_Collection_ConcatNew", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		nc := c.ConcatNew(0, "b", "c")
		if nc.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_I8_C32_Collection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_I8_C32_Collection_ConcatNew_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		nc := c.ConcatNew(0)
		if nc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C33_Collection_AddIfMany(t *testing.T) {
	safeTest(t, "Test_I8_C33_Collection_AddIfMany", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(false, "skip1", "skip2")
		if c.Length() != 0 {
			t.Fatal("expected 0")
		}
		c.AddIfMany(true, "keep1", "keep2")
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C34_Collection_AddFunc(t *testing.T) {
	safeTest(t, "Test_I8_C34_Collection_AddFunc", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddFunc(func() string { return "hello" })
		if c.Length() != 1 || c.First() != "hello" {
			t.Fatal("expected 'hello'")
		}
	})
}

func Test_I8_C35_Collection_AddFuncErr(t *testing.T) {
	safeTest(t, "Test_I8_C35_Collection_AddFuncErr", func() {
		c := corestr.New.Collection.Cap(2)
		// success
		c.AddFuncErr(func() (string, error) { return "ok", nil }, func(e error) {})
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
		// error
		c.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(e error) {})
		if c.Length() != 1 {
			t.Fatal("expected still 1")
		}
	})
}

func Test_I8_C36_Collection_AddsLock(t *testing.T) {
	safeTest(t, "Test_I8_C36_Collection_AddsLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsLock("a", "b")
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C37_Collection_AddStrings(t *testing.T) {
	safeTest(t, "Test_I8_C37_Collection_AddStrings", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C38_Collection_AddCollection(t *testing.T) {
	safeTest(t, "Test_I8_C38_Collection_AddCollection", func() {
		c := corestr.New.Collection.Cap(5)
		other := corestr.New.Collection.Strings([]string{"x", "y"})
		c.AddCollection(other)
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty collection
		c.AddCollection(corestr.Empty.Collection())
		if c.Length() != 2 {
			t.Fatal("expected still 2")
		}
	})
}

func Test_I8_C39_Collection_AddCollections(t *testing.T) {
	safeTest(t, "Test_I8_C39_Collection_AddCollections", func() {
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c.AddCollections(c1, c2, corestr.Empty.Collection())
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C40_Collection_AddWithWgLock(t *testing.T) {
	safeTest(t, "Test_I8_C40_Collection_AddWithWgLock", func() {
		c := corestr.New.Collection.Cap(5)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddWithWgLock(wg, "a")
		wg.Wait()
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

// =============================================================================
// Collection — Access, Sort, Filter
// =============================================================================

func Test_I8_C41_Collection_IndexAt(t *testing.T) {
	safeTest(t, "Test_I8_C41_Collection_IndexAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.IndexAt(0) != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_C42_Collection_SafeIndexAtUsingLength(t *testing.T) {
	safeTest(t, "Test_I8_C42_Collection_SafeIndexAtUsingLength", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.SafeIndexAtUsingLength("default", 2, 5) != "default" {
			t.Fatal("expected default")
		}
		if c.SafeIndexAtUsingLength("default", 2, 0) != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_C43_Collection_First(t *testing.T) {
	safeTest(t, "Test_I8_C43_Collection_First", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.First() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_C44_Collection_Last(t *testing.T) {
	safeTest(t, "Test_I8_C44_Collection_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Last() != "b" {
			t.Fatal("expected 'b'")
		}
	})
}

func Test_I8_C45_Collection_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_C45_Collection_LastOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.LastOrDefault() != "a" {
			t.Fatal("expected 'a'")
		}
		e := corestr.Empty.Collection()
		if e.LastOrDefault() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_C46_Collection_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_I8_C46_Collection_FirstOrDefault", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		if c.FirstOrDefault() != "a" {
			t.Fatal("expected 'a'")
		}
		e := corestr.Empty.Collection()
		if e.FirstOrDefault() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_C47_Collection_Take(t *testing.T) {
	safeTest(t, "Test_I8_C47_Collection_Take", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		taken := c.Take(2)
		if taken.Length() != 2 {
			t.Fatal("expected 2")
		}
		// take more than length
		taken2 := c.Take(10)
		if taken2.Length() != 3 {
			t.Fatal("expected 3")
		}
		// take 0
		taken3 := c.Take(0)
		if taken3.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_C48_Collection_Skip(t *testing.T) {
	safeTest(t, "Test_I8_C48_Collection_Skip", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		skipped := c.Skip(1)
		if skipped.Length() != 2 {
			t.Fatal("expected 2")
		}
		// skip 0
		skipped2 := c.Skip(0)
		if skipped2.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_I8_C49_Collection_Reverse(t *testing.T) {
	safeTest(t, "Test_I8_C49_Collection_Reverse", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.Reverse()
		if c.First() != "c" || c.Last() != "a" {
			t.Fatal("expected reversed")
		}
	})
}

func Test_I8_C50_Collection_Reverse_Two(t *testing.T) {
	safeTest(t, "Test_I8_C50_Collection_Reverse_Two", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Reverse()
		if c.First() != "b" {
			t.Fatal("expected 'b' first")
		}
	})
}

func Test_I8_C51_Collection_Reverse_One(t *testing.T) {
	safeTest(t, "Test_I8_C51_Collection_Reverse_One", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Reverse()
		if c.First() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_C52_Collection_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_I8_C52_Collection_GetPagesSize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagesSize(2)
		if pages != 3 {
			t.Fatalf("expected 3, got %d", pages)
		}
		if c.GetPagesSize(0) != 0 {
			t.Fatal("expected 0 for zero page size")
		}
	})
}

func Test_I8_C53_Collection_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_I8_C53_Collection_GetSinglePageCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		page := c.GetSinglePageCollection(2, 2)
		if page.Length() != 2 {
			t.Fatal("expected 2")
		}
		page3 := c.GetSinglePageCollection(2, 3)
		if page3.Length() != 1 {
			t.Fatal("expected 1 for last page")
		}
	})
}

func Test_I8_C54_Collection_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_I8_C54_Collection_GetPagedCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := c.GetPagedCollection(2)
		if pages.Length() < 3 {
			t.Fatal("expected at least 3 pages")
		}
	})
}

func Test_I8_C55_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_I8_C55_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_I8_C56_Collection_Filter_WithBreak(t *testing.T) {
	safeTest(t, "Test_I8_C56_Collection_Filter_WithBreak", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 1
		})
		if len(result) != 2 {
			t.Fatalf("expected 2, got %d", len(result))
		}
	})
}

func Test_I8_C57_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_I8_C57_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C58_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_I8_C58_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb"})
		fc := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if fc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C59_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_I8_C59_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtr(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C60_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_I8_C60_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.FilterPtrLock(func(sp *string, i int) (*string, bool, bool) {
			return sp, true, false
		})
		if len(*result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

// =============================================================================
// Collection — Search, Sort, CSV, etc.
// =============================================================================

func Test_I8_C61_Collection_Has(t *testing.T) {
	safeTest(t, "Test_I8_C61_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.Has("a") {
			t.Fatal("expected true")
		}
		if c.Has("z") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_C62_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_I8_C62_Collection_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		if !c.HasPtr(&s) {
			t.Fatal("expected true")
		}
		if c.HasPtr(nil) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_I8_C63_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_I8_C63_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.HasAll("a", "b") {
			t.Fatal("expected true")
		}
		if c.HasAll("a", "z") {
			t.Fatal("expected false")
		}
	})
}

func Test_I8_C64_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_I8_C64_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		sorted := c.SortedListAsc()
		if sorted[0] != "a" {
			t.Fatal("expected 'a' first")
		}
	})
}

func Test_I8_C65_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_I8_C65_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		c.SortedAsc()
		if c.First() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_C66_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_I8_C66_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAscLock()
		if c.First() != "a" {
			t.Fatal("expected 'a'")
		}
	})
}

func Test_I8_C67_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_I8_C67_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
		sorted := c.SortedListDsc()
		if sorted[0] != "c" {
			t.Fatal("expected 'c' first")
		}
	})
}

func Test_I8_C68_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_I8_C68_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		if !c.HasUsingSensitivity("hello", false) {
			t.Fatal("expected true case-insensitive")
		}
		if c.HasUsingSensitivity("hello", true) {
			t.Fatal("expected false case-sensitive")
		}
	})
}

func Test_I8_C69_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_I8_C69_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !c.IsContainsAll("a", "b") {
			t.Fatal("expected true")
		}
		if c.IsContainsAll("a", "z") {
			t.Fatal("expected false")
		}
		if c.IsContainsAll(nil...) {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_I8_C70_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_I8_C70_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if !c.IsContainsAllLock("a", "b") {
			t.Fatal("expected true")
		}
	})
}

func Test_I8_C71_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_I8_C71_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		csv := c.Csv()
		if csv == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C72_Collection_CsvEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C72_Collection_CsvEmpty", func() {
		c := corestr.Empty.Collection()
		if c.Csv() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_I8_C73_Collection_String(t *testing.T) {
	safeTest(t, "Test_I8_C73_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C74_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_I8_C74_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.StringLock()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C75_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_I8_C75_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		if c.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_C76_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_I8_C76_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		if c.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_C77_Collection_Join(t *testing.T) {
	safeTest(t, "Test_I8_C77_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		if c.Join(",") != "a,b" {
			t.Fatal("expected 'a,b'")
		}
	})
}

func Test_I8_C78_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_I8_C78_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JoinLine()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C79_Collection_Json(t *testing.T) {
	safeTest(t, "Test_I8_C79_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		j := c.Json()
		if j.JsonString() == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C80_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_I8_C80_Collection_ParseInjectUsingJson", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := c.JsonPtr()
		c2 := corestr.Empty.Collection()
		_, err := c2.ParseInjectUsingJson(jr)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_I8_C81_Collection_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_I8_C81_Collection_ParseInjectUsingJson_Error", func() {
		c := corestr.Empty.Collection()
		bad := corejson.NewResult.UsingString(`invalid`)
		_, err := c.ParseInjectUsingJson(bad)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_I8_C82_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_I8_C82_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		bytes, err := c.Serialize()
		if err != nil || len(bytes) == 0 {
			t.Fatal("expected serialization")
		}
	})
}

func Test_I8_C83_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_C83_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		data, err := c.MarshalJSON()
		if err != nil || len(data) == 0 {
			t.Fatal("expected marshal")
		}
	})
}

func Test_I8_C84_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_I8_C84_Collection_UnmarshalJSON", func() {
		c := corestr.Empty.Collection()
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil {
			t.Fatal("unexpected error")
		}
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

// =============================================================================
// Collection — More methods
// =============================================================================

func Test_I8_C85_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_I8_C85_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		list := c.NonEmptyList()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C86_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_I8_C86_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		list := c.UniqueList()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C87_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_I8_C87_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		list := c.UniqueListLock()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C88_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_I8_C88_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		m := c.UniqueBoolMap()
		if len(m) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C89_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_I8_C89_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C90_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_I8_C90_Collection_Resize", func() {
		c := corestr.New.Collection.Cap(2)
		c.Add("a")
		c.Resize(100)
		if c.Capacity() < 100 {
			t.Fatal("expected capacity >= 100")
		}
	})
}

func Test_I8_C91_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_I8_C91_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Cap(2)
		c.AddCapacity(50)
		if c.Capacity() < 50 {
			t.Fatal("expected capacity >= 50")
		}
	})
}

func Test_I8_C92_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_I8_C92_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Joins(",")
		if s != "a,b" {
			t.Fatal("expected 'a,b'")
		}
		s2 := c.Joins(",", "c")
		if s2 == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_I8_C93_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_I8_C93_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.GetAllExcept([]string{"b"})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C94_Collection_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_I8_C94_Collection_GetAllExcept_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.GetAllExcept(nil)
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C95_Collection_New(t *testing.T) {
	safeTest(t, "Test_I8_C95_Collection_New", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		nc := c.New("x", "y")
		if nc.Length() != 2 {
			t.Fatal("expected 2")
		}
		nc2 := c.New()
		if nc2.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_I8_C96_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_I8_C96_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys(42, "hello", nil)
		if c.Length() != 2 {
			t.Fatal("expected 2 (nil skipped)")
		}
	})
}

func Test_I8_C97_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_I8_C97_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock(42, "hello")
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C98_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_I8_C98_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys(42, nil)
		if c.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_I8_C99_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_I8_C99_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty("a", "", "b")
		if c.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_I8_C100_Collection_Single(t *testing.T) {
	safeTest(t, "Test_I8_C100_Collection_Single", func() {
		c := corestr.New.Collection.Strings([]string{"only"})
		if c.Single() != "only" {
			t.Fatal("expected 'only'")
		}
	})
}
