package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 2: Take, Skip, Reverse, Page, Sort, Filter, JSON, etc.
// Covers ~200 uncovered statements from Collection.go lines 700-2201
// ══════════════════════════════════════════════════════════════════════════════

func Test_CovP2_01_Take(t *testing.T) {
	safeTest(t, "Test_CovP2_01_Take", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		// take less than length
		taken := col.Take(2)
		if taken.Length() != 2 {
			t.Fatal("expected 2")
		}
		// take >= length
		same := col.Take(10)
		if same.Length() != 4 {
			t.Fatal("expected 4")
		}
		// take 0
		zero := col.Take(0)
		if zero.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_02_Skip(t *testing.T) {
	safeTest(t, "Test_CovP2_02_Skip", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		skipped := col.Skip(2)
		if skipped.Length() != 2 {
			t.Fatal("expected 2")
		}
		// skip 0
		same := col.Skip(0)
		if same.Length() != 4 {
			t.Fatal("expected 4")
		}
		// skip > length => panic
		defer func() { recover() }()
		col.Skip(100)
		t.Fatal("expected panic")
	})
}

func Test_CovP2_03_Reverse(t *testing.T) {
	safeTest(t, "Test_CovP2_03_Reverse", func() {
		// single element
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col1.Reverse()
		if col1.First() != "a" {
			t.Fatal("expected a")
		}

		// two elements
		col2 := corestr.New.Collection.Strings([]string{"a", "b"})
		col2.Reverse()
		if col2.First() != "b" {
			t.Fatal("expected b first")
		}

		// three+ elements
		col3 := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		col3.Reverse()
		if col3.First() != "d" || col3.Last() != "a" {
			t.Fatal("expected reversed")
		}
	})
}

func Test_CovP2_04_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_CovP2_04_GetPagesSize", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := col.GetPagesSize(2)
		if pages != 3 {
			t.Fatalf("expected 3, got %d", pages)
		}
		// zero page size
		if col.GetPagesSize(0) != 0 {
			t.Fatal("expected 0")
		}
		// negative
		if col.GetPagesSize(-1) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_05_GetPagedCollection(t *testing.T) {
	safeTest(t, "Test_CovP2_05_GetPagedCollection", func() {
		items := make([]string, 25)
		for i := range items {
			items[i] = "x"
		}
		col := corestr.New.Collection.Strings(items)
		paged := col.GetPagedCollection(10)
		if paged.Length() != 3 {
			t.Fatalf("expected 3 pages, got %d", paged.Length())
		}

		// less than page size
		small := corestr.New.Collection.Strings([]string{"a", "b"})
		pagedSmall := small.GetPagedCollection(10)
		if pagedSmall.Length() != 1 {
			t.Fatal("expected 1 page")
		}
	})
}

func Test_CovP2_06_GetSinglePageCollection(t *testing.T) {
	safeTest(t, "Test_CovP2_06_GetSinglePageCollection", func() {
		items := make([]string, 25)
		for i := range items {
			items[i] = "x"
		}
		col := corestr.New.Collection.Strings(items)
		// page 1
		page1 := col.GetSinglePageCollection(10, 1)
		if page1.Length() != 10 {
			t.Fatalf("expected 10, got %d", page1.Length())
		}
		// page 3 (partial)
		page3 := col.GetSinglePageCollection(10, 3)
		if page3.Length() != 5 {
			t.Fatalf("expected 5, got %d", page3.Length())
		}
		// small collection
		small := corestr.New.Collection.Strings([]string{"a"})
		same := small.GetSinglePageCollection(10, 1)
		if same.Length() != 1 {
			t.Fatal("expected 1")
		}
		// negative page index
		defer func() { recover() }()
		col.GetSinglePageCollection(10, 0)
		t.Fatal("expected panic")
	})
}

func Test_CovP2_07_AddStringsAsync(t *testing.T) {
	safeTest(t, "Test_CovP2_07_AddStringsAsync", func() {
		col := corestr.New.Collection.Cap(10)
		wg := &sync.WaitGroup{}
		col.AddStringsAsync(wg, []string{"a", "b", "c"})
		wg.Wait()
		// empty
		col.AddStringsAsync(wg, []string{})
	})
}

func Test_CovP2_08_InsertAt(t *testing.T) {
	safeTest(t, "Test_CovP2_08_InsertAt", func() {
		col := corestr.New.Collection.Strings([]string{"a", "c"})
		col.InsertAt(0, "x") // at last index  
		if col.Length() != 3 {
			t.Fatal("expected 3")
		}
		// empty collection
		empty := corestr.Empty.Collection()
		empty.InsertAt(0, "first")
	})
}

func Test_CovP2_09_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_CovP2_09_ChainRemoveAt", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		col.ChainRemoveAt(1)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_10_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_CovP2_10_RemoveItemsIndexes", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		col.RemoveItemsIndexes(true, 1, 3)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil indexes with ignore
		col.RemoveItemsIndexes(true)
	})
}

func Test_CovP2_11_RemoveItemsIndexesPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_11_RemoveItemsIndexesPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		col.RemoveItemsIndexesPtr(false, []int{0})
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil indexes
		col.RemoveItemsIndexesPtr(false, nil)

		// panic on empty with validate
		defer func() { recover() }()
		empty := corestr.Empty.Collection()
		empty.RemoveItemsIndexesPtr(false, []int{0})
	})
}

func Test_CovP2_12_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_12_AppendCollectionPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b", "c"})
		col.AppendCollectionPtr(other)
		if col.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovP2_13_AppendCollections(t *testing.T) {
	safeTest(t, "Test_CovP2_13_AppendCollections", func() {
		col := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		col.AppendCollections(c1, c2, corestr.Empty.Collection())
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty
		col.AppendCollections()
	})
}

func Test_CovP2_14_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_CovP2_14_AppendAnysLock", func() {
		col := corestr.New.Collection.Cap(5)
		col.AppendAnysLock("a", 123, nil)
		if col.Length() != 2 {
			t.Fatal("expected 2 (nil skipped)")
		}
		col.AppendAnysLock()
	})
}

func Test_CovP2_15_AppendAnys(t *testing.T) {
	safeTest(t, "Test_CovP2_15_AppendAnys", func() {
		col := corestr.New.Collection.Cap(5)
		col.AppendAnys("a", 42, nil, "b")
		if col.Length() != 3 {
			t.Fatal("expected 3")
		}
		col.AppendAnys()
	})
}

func Test_CovP2_16_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_CovP2_16_AppendAnysUsingFilter", func() {
		col := corestr.New.Collection.Cap(10)
		col.AppendAnysUsingFilter(
			func(str string, i int) (string, bool, bool) {
				return str, true, false
			},
			"a", nil, "b",
		)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// empty
		col.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) { return s, true, false })
		// break
		col2 := corestr.New.Collection.Cap(10)
		col2.AppendAnysUsingFilter(
			func(str string, i int) (string, bool, bool) { return str, true, true },
			"x", "y",
		)
		// not keep
		col3 := corestr.New.Collection.Cap(10)
		col3.AppendAnysUsingFilter(
			func(str string, i int) (string, bool, bool) { return str, false, false },
			"x",
		)
	})
}

func Test_CovP2_17_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_CovP2_17_AppendAnysUsingFilterLock", func() {
		col := corestr.New.Collection.Cap(10)
		col.AppendAnysUsingFilterLock(
			func(str string, i int) (string, bool, bool) { return str, true, false },
			"a", nil, "b",
		)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil items
		col.AppendAnysUsingFilterLock(
			func(str string, i int) (string, bool, bool) { return str, true, false },
		)
		// break + skip
		col2 := corestr.New.Collection.Cap(10)
		col2.AppendAnysUsingFilterLock(
			func(str string, i int) (string, bool, bool) { return str, true, true },
			"x",
		)
		col3 := corestr.New.Collection.Cap(10)
		col3.AppendAnysUsingFilterLock(
			func(str string, i int) (string, bool, bool) { return str, false, false },
			"x",
		)
	})
}

func Test_CovP2_18_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_CovP2_18_AppendNonEmptyAnys", func() {
		col := corestr.New.Collection.Cap(10)
		col.AppendNonEmptyAnys("a", nil, "", "b")
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil
		col.AppendNonEmptyAnys()
	})
}

func Test_CovP2_19_AddsAsync(t *testing.T) {
	safeTest(t, "Test_CovP2_19_AddsAsync", func() {
		col := corestr.New.Collection.Cap(10)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col.AddsAsync(wg, "a", "b")
		wg.Wait()
		// nil
		col.AddsAsync(wg)
	})
}

func Test_CovP2_20_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovP2_20_AddsNonEmpty", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddsNonEmpty("a", "", "b")
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// nil
		col.AddsNonEmpty()
	})
}

func Test_CovP2_21_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_CovP2_21_AddsNonEmptyPtrLock", func() {
		col := corestr.New.Collection.Cap(5)
		a := "hello"
		empty := ""
		col.AddsNonEmptyPtrLock(&a, nil, &empty)
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
		col.AddsNonEmptyPtrLock()
	})
}

func Test_CovP2_22_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_CovP2_22_UniqueBoolMapLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		m := col.UniqueBoolMapLock()
		if len(m) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_23_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_CovP2_23_UniqueBoolMap", func() {
		col := corestr.New.Collection.Strings([]string{"x", "y", "x"})
		m := col.UniqueBoolMap()
		if len(m) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_24_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_CovP2_24_UniqueListLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		list := col.UniqueListLock()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_25_UniqueList(t *testing.T) {
	safeTest(t, "Test_CovP2_25_UniqueList", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		list := col.UniqueList()
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_26_Filter(t *testing.T) {
	safeTest(t, "Test_CovP2_26_Filter", func() {
		col := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := col.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		// empty
		empty := corestr.Empty.Collection()
		r := empty.Filter(func(s string, i int) (string, bool, bool) { return s, true, false })
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
		// break
		col2 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r2 := col2.Filter(func(s string, i int) (string, bool, bool) { return s, true, true })
		if len(r2) != 1 {
			t.Fatal("expected 1 (break on first)")
		}
	})
}

func Test_CovP2_27_FilterLock(t *testing.T) {
	safeTest(t, "Test_CovP2_27_FilterLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := col.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		// break
		r2 := col.FilterLock(func(s string, i int) (string, bool, bool) { return s, true, true })
		if len(r2) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP2_28_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_CovP2_28_FilteredCollection", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := col.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, s == "a", false
		})
		if fc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP2_29_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovP2_29_FilteredCollectionLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := col.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		if fc.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_30_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_CovP2_30_FilterPtrLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := col.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		if len(*result) != 3 {
			t.Fatal("expected 3")
		}
		// break
		r2 := col.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		})
		if len(*r2) != 1 {
			t.Fatal("expected 1")
		}
		// empty
		empty := corestr.Empty.Collection()
		r3 := empty.FilterPtrLock(func(s *string, i int) (*string, bool, bool) { return s, true, false })
		if len(*r3) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_31_FilterPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_31_FilterPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		result := col.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		if len(*result) != 2 {
			t.Fatal("expected 2")
		}
		// break
		r2 := col.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		})
		if len(*r2) != 1 {
			t.Fatal("expected 1")
		}
		// empty
		empty := corestr.Empty.Collection()
		r3 := empty.FilterPtr(func(s *string, i int) (*string, bool, bool) { return s, true, false })
		if len(*r3) != 0 {
			t.Fatal("expected 0")
		}
	})
}
func Test_CovP2_33_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_CovP2_33_NonEmptyList", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := col.NonEmptyList()
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		empty := corestr.Empty.Collection()
		r := empty.NonEmptyList()
		if len(r) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_34_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_CovP2_34_HashsetAsIs", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := col.HashsetAsIs()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_35_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_CovP2_35_HashsetWithDoubleLength", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := col.HashsetWithDoubleLength()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_36_HashsetLock(t *testing.T) {
	safeTest(t, "Test_CovP2_36_HashsetLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := col.HashsetLock()
		if hs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_37_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_CovP2_37_NonEmptyItems", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := col.NonEmptyItems()
		if len(r) != 2 {
			t.Fatal("expected 2")
		}
	})
}
func Test_CovP2_39_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_CovP2_39_NonEmptyItemsOrNonWhitespace", func() {
		col := corestr.New.Collection.Strings([]string{"a", "   ", "b"})
		r := col.NonEmptyItemsOrNonWhitespace()
		if len(r) != 2 {
			t.Fatal("expected 2")
		}
	})
}
func Test_CovP2_41_Items(t *testing.T) {
	safeTest(t, "Test_CovP2_41_Items", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if len(col.Items()) != 1 {
			t.Fatal("expected 1")
		}
	})
}
func Test_CovP2_43_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_CovP2_43_ListCopyPtrLock", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		r := col.ListCopyPtrLock()
		if len(r) != 1 {
			t.Fatal("expected 1")
		}
		empty := corestr.Empty.Collection()
		r2 := empty.ListCopyPtrLock()
		if len(r2) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_44_HasLock(t *testing.T) {
	safeTest(t, "Test_CovP2_44_HasLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if !col.HasLock("a") {
			t.Fatal("expected true")
		}
	})
}

func Test_CovP2_45_Has(t *testing.T) {
	safeTest(t, "Test_CovP2_45_Has", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if !col.Has("a") {
			t.Fatal("expected true")
		}
		if col.Has("z") {
			t.Fatal("expected false")
		}
		empty := corestr.Empty.Collection()
		if empty.Has("a") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP2_46_HasPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_46_HasPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "a"
		if !col.HasPtr(&s) {
			t.Fatal("expected true")
		}
		if col.HasPtr(nil) {
			t.Fatal("expected false for nil")
		}
		miss := "z"
		if col.HasPtr(&miss) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP2_47_HasAll(t *testing.T) {
	safeTest(t, "Test_CovP2_47_HasAll", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !col.HasAll("a", "b") {
			t.Fatal("expected true")
		}
		if col.HasAll("a", "z") {
			t.Fatal("expected false")
		}
		empty := corestr.Empty.Collection()
		if empty.HasAll("a") {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP2_48_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_CovP2_48_SortedListAsc", func() {
		col := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		sorted := col.SortedListAsc()
		if sorted[0] != "a" {
			t.Fatal("expected a first")
		}
		empty := corestr.Empty.Collection()
		r := empty.SortedListAsc()
		if len(r) != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovP2_49_SortedAsc(t *testing.T) {
	safeTest(t, "Test_CovP2_49_SortedAsc", func() {
		col := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		col.SortedAsc()
		if col.First() != "a" {
			t.Fatal("expected a first")
		}
		empty := corestr.Empty.Collection()
		empty.SortedAsc()
	})
}

func Test_CovP2_50_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_CovP2_50_SortedAscLock", func() {
		col := corestr.New.Collection.Strings([]string{"c", "a"})
		col.SortedAscLock()
		if col.First() != "a" {
			t.Fatal("expected a")
		}
		empty := corestr.Empty.Collection()
		empty.SortedAscLock()
	})
}

func Test_CovP2_51_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_CovP2_51_SortedListDsc", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		sorted := col.SortedListDsc()
		if sorted[0] != "c" {
			t.Fatal("expected c first")
		}
	})
}

func Test_CovP2_52_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_CovP2_52_HasUsingSensitivity", func() {
		col := corestr.New.Collection.Strings([]string{"Hello", "World"})
		if !col.HasUsingSensitivity("hello", false) {
			t.Fatal("expected true case-insensitive")
		}
		if col.HasUsingSensitivity("hello", true) {
			t.Fatal("expected false case-sensitive")
		}
		if col.HasUsingSensitivity("missing", false) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP2_53_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_53_IsContainsPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		a := "a"
		if !col.IsContainsPtr(&a) {
			t.Fatal("expected true")
		}
		if col.IsContainsPtr(nil) {
			t.Fatal("expected false for nil")
		}
		miss := "z"
		if col.IsContainsPtr(&miss) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP2_54_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_CovP2_54_GetHashsetPlusHasAll", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		hs, ok := col.GetHashsetPlusHasAll([]string{"a", "b"})
		if !ok {
			t.Fatal("expected true")
		}
		if hs.Length() != 3 {
			t.Fatal("expected 3")
		}
		// nil items
		_, ok2 := col.GetHashsetPlusHasAll(nil)
		if ok2 {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_CovP2_55_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_CovP2_55_IsContainsAllSlice", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		if !col.IsContainsAllSlice([]string{"a", "b"}) {
			t.Fatal("expected true")
		}
		if col.IsContainsAllSlice([]string{}) {
			t.Fatal("expected false for empty")
		}
		if col.IsContainsAllSlice([]string{"z"}) {
			t.Fatal("expected false")
		}
	})
}

func Test_CovP2_56_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_CovP2_56_IsContainsAll", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if !col.IsContainsAll("a", "b") {
			t.Fatal("expected true")
		}
		if col.IsContainsAll() {
			t.Fatal("expected false for nil variadic")
		}
	})
}

func Test_CovP2_57_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_CovP2_57_IsContainsAllLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if !col.IsContainsAllLock("a") {
			t.Fatal("expected true")
		}
		if col.IsContainsAllLock() {
			t.Fatal("expected false for nil")
		}
	})
}

func Test_CovP2_58_New(t *testing.T) {
	safeTest(t, "Test_CovP2_58_New", func() {
		col := corestr.New.Collection.Strings([]string{"x"})
		newCol := col.New("a", "b")
		if newCol.Length() != 2 {
			t.Fatal("expected 2")
		}
		emptyNew := col.New()
		if emptyNew.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_59_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_CovP2_59_AddNonEmptyStrings", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddNonEmptyStrings("a", "", "b")
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		col.AddNonEmptyStrings()
	})
}

func Test_CovP2_60_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_CovP2_60_AddFuncResult", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddFuncResult(
			func() string { return "a" },
			func() string { return "b" },
		)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		col.AddFuncResult()
	})
}

func Test_CovP2_61_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_CovP2_61_AddNonEmptyStringsSlice", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddNonEmptyStringsSlice([]string{"a", "b"})
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		col.AddNonEmptyStringsSlice([]string{})
	})
}

func Test_CovP2_62_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_CovP2_62_AddStringsByFuncChecking", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddStringsByFuncChecking(
			[]string{"abc", "a", "abcd"},
			func(line string) bool { return len(line) > 2 },
		)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_63_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_CovP2_63_ExpandSlicePlusAdd", func() {
		col := corestr.New.Collection.Cap(10)
		col.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string { return []string{line + "!"} },
		)
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_64_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_CovP2_64_MergeSlicesOfSlice", func() {
		col := corestr.New.Collection.Cap(10)
		col.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_65_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_CovP2_65_GetAllExceptCollection", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := col.GetAllExceptCollection(except)
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		// nil
		r2 := col.GetAllExceptCollection(nil)
		if len(r2) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovP2_66_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_CovP2_66_GetAllExcept", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := col.GetAllExcept([]string{"a"})
		if len(result) != 2 {
			t.Fatal("expected 2")
		}
		r2 := col.GetAllExcept(nil)
		if len(r2) != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_CovP2_67_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_CovP2_67_CharCollectionMap", func() {
		col := corestr.New.Collection.Strings([]string{"abc", "bcd"})
		ccm := col.CharCollectionMap()
		if ccm == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_CovP2_68_SummaryString(t *testing.T) {
	safeTest(t, "Test_CovP2_68_SummaryString", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.SummaryString(1)
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovP2_69_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_CovP2_69_SummaryStringWithHeader", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		s := col.SummaryStringWithHeader("Header:")
		if s == "" {
			t.Fatal("expected non-empty")
		}
		empty := corestr.Empty.Collection()
		s2 := empty.SummaryStringWithHeader("Header:")
		if s2 == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovP2_70_String(t *testing.T) {
	safeTest(t, "Test_CovP2_70_String", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		if col.String() == "" {
			t.Fatal("expected non-empty")
		}
		empty := corestr.Empty.Collection()
		_ = empty.String()
	})
}

func Test_CovP2_71_CsvLines(t *testing.T) {
	safeTest(t, "Test_CovP2_71_CsvLines", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lines := col.CsvLines()
		if len(lines) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_CovP2_72_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_CovP2_72_CsvLinesOptions", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.CsvLinesOptions(true)
	})
}

func Test_CovP2_73_Csv(t *testing.T) {
	safeTest(t, "Test_CovP2_73_Csv", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		csv := col.Csv()
		if csv == "" {
			t.Fatal("expected non-empty")
		}
		empty := corestr.Empty.Collection()
		if empty.Csv() != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovP2_74_CsvOptions(t *testing.T) {
	safeTest(t, "Test_CovP2_74_CsvOptions", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.CsvOptions(true)
		empty := corestr.Empty.Collection()
		_ = empty.CsvOptions(false)
	})
}

func Test_CovP2_75_StringLock(t *testing.T) {
	safeTest(t, "Test_CovP2_75_StringLock", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		s := col.StringLock()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		empty := corestr.Empty.Collection()
		_ = empty.StringLock()
	})
}

func Test_CovP2_76_AddCapacity(t *testing.T) {
	safeTest(t, "Test_CovP2_76_AddCapacity", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddCapacity(10, 20)
		if col.Capacity() < 30 {
			t.Fatal("expected at least 30")
		}
		col.AddCapacity()
	})
}

func Test_CovP2_77_Resize(t *testing.T) {
	safeTest(t, "Test_CovP2_77_Resize", func() {
		col := corestr.New.Collection.Cap(5)
		col.Resize(100)
		if col.Capacity() < 100 {
			t.Fatal("expected at least 100")
		}
		// no resize needed
		col.Resize(5)
	})
}

func Test_CovP2_78_Joins(t *testing.T) {
	safeTest(t, "Test_CovP2_78_Joins", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.Joins(",")
		if s != "a,b" {
			t.Fatalf("expected 'a,b', got '%s'", s)
		}
		s2 := col.Joins(",", "c")
		if s2 == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_CovP2_79_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_CovP2_79_NonEmptyJoins", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		s := col.NonEmptyJoins(",")
		_ = s
	})
}

func Test_CovP2_80_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_CovP2_80_NonWhitespaceJoins", func() {
		col := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		_ = col.NonWhitespaceJoins(",")
	})
}

func Test_CovP2_81_JsonModel(t *testing.T) {
	safeTest(t, "Test_CovP2_81_JsonModel", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		m := col.JsonModel()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP2_82_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_CovP2_82_JsonModelAny", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.JsonModelAny()
	})
}

func Test_CovP2_83_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovP2_83_MarshalJSON", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		data, err := col.MarshalJSON()
		if err != nil {
			t.Fatal("unexpected error")
		}
		if len(data) == 0 {
			t.Fatal("expected data")
		}
	})
}

func Test_CovP2_84_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovP2_84_UnmarshalJSON", func() {
		col := corestr.New.Collection.Cap(5)
		err := col.UnmarshalJSON([]byte(`["a","b"]`))
		if err != nil {
			t.Fatal("unexpected error")
		}
		if col.Length() != 2 {
			t.Fatal("expected 2")
		}
		// invalid json
		err2 := col.UnmarshalJSON([]byte(`invalid`))
		if err2 == nil {
			t.Fatal("expected error")
		}
	})
}

func Test_CovP2_85_Json(t *testing.T) {
	safeTest(t, "Test_CovP2_85_Json", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.Json()
	})
}

func Test_CovP2_86_JsonPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_86_JsonPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.JsonPtr()
	})
}

func Test_CovP2_87_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_CovP2_87_ParseInjectUsingJson", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := col.JsonPtr()
		newCol := corestr.New.Collection.Cap(5)
		result, err := newCol.ParseInjectUsingJson(jsonResult)
		if err != nil {
			t.Fatal("unexpected error")
		}
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP2_88_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovP2_88_ParseInjectUsingJsonMust", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := col.JsonPtr()
		newCol := corestr.New.Collection.Cap(5)
		result := newCol.ParseInjectUsingJsonMust(jsonResult)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_CovP2_89_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovP2_89_JsonParseSelfInject", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := col.JsonPtr()
		newCol := corestr.New.Collection.Cap(5)
		err := newCol.JsonParseSelfInject(jsonResult)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovP2_90_Clear(t *testing.T) {
	safeTest(t, "Test_CovP2_90_Clear", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		col.Clear()
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_91_Dispose(t *testing.T) {
	safeTest(t, "Test_CovP2_91_Dispose", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		col.Dispose()
		if col.Length() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_CovP2_92_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_CovP2_92_AsJsonMarshaller", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.AsJsonMarshaller()
	})
}

func Test_CovP2_93_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_CovP2_93_AsJsonContractsBinder", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_ = col.AsJsonContractsBinder()
	})
}

func Test_CovP2_94_Serialize(t *testing.T) {
	safeTest(t, "Test_CovP2_94_Serialize", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		_, err := col.Serialize()
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovP2_95_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovP2_95_Deserialize", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		var target []string
		err := col.Deserialize(&target)
		if err != nil {
			t.Fatal("unexpected error")
		}
	})
}

func Test_CovP2_96_Join(t *testing.T) {
	safeTest(t, "Test_CovP2_96_Join", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		if col.Join(",") != "a,b" {
			t.Fatal("expected a,b")
		}
		empty := corestr.Empty.Collection()
		if empty.Join(",") != "" {
			t.Fatal("expected empty")
		}
	})
}

func Test_CovP2_97_JoinLine(t *testing.T) {
	safeTest(t, "Test_CovP2_97_JoinLine", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.JoinLine()
		if s == "" {
			t.Fatal("expected non-empty")
		}
		empty := corestr.Empty.Collection()
		if empty.JoinLine() != "" {
			t.Fatal("expected empty")
		}
	})
}
