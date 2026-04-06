package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": taken.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// take >= length
		same := col.Take(10)
		actual := args.Map{"result": same.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
		// take 0
		zero := col.Take(0)
		actual := args.Map{"result": zero.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_02_Skip(t *testing.T) {
	safeTest(t, "Test_CovP2_02_Skip", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		skipped := col.Skip(2)
		actual := args.Map{"result": skipped.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// skip 0
		same := col.Skip(0)
		actual := args.Map{"result": same.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
		// skip > length => panic
		defer func() { recover() }()
		col.Skip(100)
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	})
}

func Test_CovP2_03_Reverse(t *testing.T) {
	safeTest(t, "Test_CovP2_03_Reverse", func() {
		// single element
		col1 := corestr.New.Collection.Strings([]string{"a"})
		col1.Reverse()
		actual := args.Map{"result": col1.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)

		// two elements
		col2 := corestr.New.Collection.Strings([]string{"a", "b"})
		col2.Reverse()
		actual := args.Map{"result": col2.First() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b first", actual)

		// three+ elements
		col3 := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		col3.Reverse()
		actual := args.Map{"result": col3.First() != "d" || col3.Last() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reversed", actual)
	})
}

func Test_CovP2_04_GetPagesSize(t *testing.T) {
	safeTest(t, "Test_CovP2_04_GetPagesSize", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d", "e"})
		pages := col.GetPagesSize(2)
		actual := args.Map{"result": pages != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// zero page size
		actual := args.Map{"result": col.GetPagesSize(0) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// negative
		actual := args.Map{"result": col.GetPagesSize(-1) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": paged.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)

		// less than page size
		small := corestr.New.Collection.Strings([]string{"a", "b"})
		pagedSmall := small.GetPagedCollection(10)
		actual := args.Map{"result": pagedSmall.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 page", actual)
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
		actual := args.Map{"result": page1.Length() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		// page 3 (partial)
		page3 := col.GetSinglePageCollection(10, 3)
		actual := args.Map{"result": page3.Length() != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
		// small collection
		small := corestr.New.Collection.Strings([]string{"a"})
		same := small.GetSinglePageCollection(10, 1)
		actual := args.Map{"result": same.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// negative page index
		defer func() { recover() }()
		col.GetSinglePageCollection(10, 0)
		actual := args.Map{"result": false}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
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
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// empty collection
		empty := corestr.Empty.Collection()
		empty.InsertAt(0, "first")
	})
}

func Test_CovP2_09_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_CovP2_09_ChainRemoveAt", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		col.ChainRemoveAt(1)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_10_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_CovP2_10_RemoveItemsIndexes", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		col.RemoveItemsIndexes(true, 1, 3)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil indexes with ignore
		col.RemoveItemsIndexes(true)
	})
}

func Test_CovP2_11_RemoveItemsIndexesPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_11_RemoveItemsIndexesPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		col.RemoveItemsIndexesPtr(false, []int{0})
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovP2_13_AppendCollections(t *testing.T) {
	safeTest(t, "Test_CovP2_13_AppendCollections", func() {
		col := corestr.New.Collection.Cap(10)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		col.AppendCollections(c1, c2, corestr.Empty.Collection())
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		col.AppendCollections()
	})
}

func Test_CovP2_14_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_CovP2_14_AppendAnysLock", func() {
		col := corestr.New.Collection.Cap(5)
		col.AppendAnysLock("a", 123, nil)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
		col.AppendAnysLock()
	})
}

func Test_CovP2_15_AppendAnys(t *testing.T) {
	safeTest(t, "Test_CovP2_15_AppendAnys", func() {
		col := corestr.New.Collection.Cap(5)
		col.AppendAnys("a", 42, nil, "b")
		actual := args.Map{"result": col.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		col.AddsNonEmptyPtrLock()
	})
}

func Test_CovP2_22_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_CovP2_22_UniqueBoolMapLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		m := col.UniqueBoolMapLock()
		actual := args.Map{"result": len(m) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_23_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_CovP2_23_UniqueBoolMap", func() {
		col := corestr.New.Collection.Strings([]string{"x", "y", "x"})
		m := col.UniqueBoolMap()
		actual := args.Map{"result": len(m) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_24_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_CovP2_24_UniqueListLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "a", "b"})
		list := col.UniqueListLock()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_25_UniqueList(t *testing.T) {
	safeTest(t, "Test_CovP2_25_UniqueList", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		list := col.UniqueList()
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_26_Filter(t *testing.T) {
	safeTest(t, "Test_CovP2_26_Filter", func() {
		col := corestr.New.Collection.Strings([]string{"a", "bb", "ccc"})
		result := col.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// empty
		empty := corestr.Empty.Collection()
		r := empty.Filter(func(s string, i int) (string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(r) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// break
		col2 := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r2 := col2.Filter(func(s string, i int) (string, bool, bool) { return s, true, true })
		actual := args.Map{"result": len(r2) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break on first)", actual)
	})
}

func Test_CovP2_27_FilterLock(t *testing.T) {
	safeTest(t, "Test_CovP2_27_FilterLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "bb"})
		result := col.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// break
		r2 := col.FilterLock(func(s string, i int) (string, bool, bool) { return s, true, true })
		actual := args.Map{"result": len(r2) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP2_28_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_CovP2_28_FilteredCollection", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := col.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, s == "a", false
		})
		actual := args.Map{"result": fc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP2_29_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_CovP2_29_FilteredCollectionLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		fc := col.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": fc.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_30_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_CovP2_30_FilterPtrLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := col.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": len(*result) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// break
		r2 := col.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		})
		actual := args.Map{"result": len(*r2) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		empty := corestr.Empty.Collection()
		r3 := empty.FilterPtrLock(func(s *string, i int) (*string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(*r3) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_31_FilterPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_31_FilterPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		result := col.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})
		actual := args.Map{"result": len(*result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// break
		r2 := col.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		})
		actual := args.Map{"result": len(*r2) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		// empty
		empty := corestr.Empty.Collection()
		r3 := empty.FilterPtr(func(s *string, i int) (*string, bool, bool) { return s, true, false })
		actual := args.Map{"result": len(*r3) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_32_NonEmptyListPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_32_NonEmptyListPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := col.NonEmptyListPtr()
		actual := args.Map{"result": len(*result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_33_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_CovP2_33_NonEmptyList", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := col.NonEmptyList()
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		empty := corestr.Empty.Collection()
		r := empty.NonEmptyList()
		actual := args.Map{"result": len(r) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_34_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_CovP2_34_HashsetAsIs", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := col.HashsetAsIs()
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_35_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_CovP2_35_HashsetWithDoubleLength", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := col.HashsetWithDoubleLength()
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_36_HashsetLock(t *testing.T) {
	safeTest(t, "Test_CovP2_36_HashsetLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := col.HashsetLock()
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_37_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_CovP2_37_NonEmptyItems", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := col.NonEmptyItems()
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_38_NonEmptyItemsPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_38_NonEmptyItemsPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := col.NonEmptyItemsPtr()
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_39_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_CovP2_39_NonEmptyItemsOrNonWhitespace", func() {
		col := corestr.New.Collection.Strings([]string{"a", "   ", "b"})
		r := col.NonEmptyItemsOrNonWhitespace()
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_40_NonEmptyItemsOrNonWhitespacePtr(t *testing.T) {
	safeTest(t, "Test_CovP2_40_NonEmptyItemsOrNonWhitespacePtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "   ", "b"})
		r := col.NonEmptyItemsOrNonWhitespacePtr()
		actual := args.Map{"result": len(r) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_41_Items(t *testing.T) {
	safeTest(t, "Test_CovP2_41_Items", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(col.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP2_42_ListPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_42_ListPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": len(col.ListPtr()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP2_43_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_CovP2_43_ListCopyPtrLock", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		r := col.ListCopyPtrLock()
		actual := args.Map{"result": len(r) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		empty := corestr.Empty.Collection()
		r2 := empty.ListCopyPtrLock()
		actual := args.Map{"result": len(r2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_44_HasLock(t *testing.T) {
	safeTest(t, "Test_CovP2_44_HasLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": col.HasLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovP2_45_Has(t *testing.T) {
	safeTest(t, "Test_CovP2_45_Has", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": col.Has("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.Has("z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.Has("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP2_46_HasPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_46_HasPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "a"
		actual := args.Map{"result": col.HasPtr(&s)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.HasPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
		miss := "z"
		actual := args.Map{"result": col.HasPtr(&miss)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP2_47_HasAll(t *testing.T) {
	safeTest(t, "Test_CovP2_47_HasAll", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"result": col.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.HasAll("a", "z")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.HasAll("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP2_48_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_CovP2_48_SortedListAsc", func() {
		col := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		sorted := col.SortedListAsc()
		actual := args.Map{"result": sorted[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		empty := corestr.Empty.Collection()
		r := empty.SortedListAsc()
		actual := args.Map{"result": len(r) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovP2_49_SortedAsc(t *testing.T) {
	safeTest(t, "Test_CovP2_49_SortedAsc", func() {
		col := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		col.SortedAsc()
		actual := args.Map{"result": col.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a first", actual)
		empty := corestr.Empty.Collection()
		empty.SortedAsc()
	})
}

func Test_CovP2_50_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_CovP2_50_SortedAscLock", func() {
		col := corestr.New.Collection.Strings([]string{"c", "a"})
		col.SortedAscLock()
		actual := args.Map{"result": col.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		empty := corestr.Empty.Collection()
		empty.SortedAscLock()
	})
}

func Test_CovP2_51_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_CovP2_51_SortedListDsc", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		sorted := col.SortedListDsc()
		actual := args.Map{"result": sorted[0] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c first", actual)
	})
}

func Test_CovP2_52_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_CovP2_52_HasUsingSensitivity", func() {
		col := corestr.New.Collection.Strings([]string{"Hello", "World"})
		actual := args.Map{"result": col.HasUsingSensitivity("hello", false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true case-insensitive", actual)
		actual := args.Map{"result": col.HasUsingSensitivity("hello", true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false case-sensitive", actual)
		actual := args.Map{"result": col.HasUsingSensitivity("missing", false)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP2_53_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_CovP2_53_IsContainsPtr", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		a := "a"
		actual := args.Map{"result": col.IsContainsPtr(&a)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.IsContainsPtr(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
		miss := "z"
		actual := args.Map{"result": col.IsContainsPtr(&miss)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP2_54_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_CovP2_54_GetHashsetPlusHasAll", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		hs, ok := col.GetHashsetPlusHasAll([]string{"a", "b"})
		actual := args.Map{"result": ok}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": hs.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		// nil items
		_, ok2 := col.GetHashsetPlusHasAll(nil)
		actual := args.Map{"result": ok2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_CovP2_55_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_CovP2_55_IsContainsAllSlice", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"result": col.IsContainsAllSlice([]string{"a", "b"})}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.IsContainsAllSlice([]string{})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
		actual := args.Map{"result": col.IsContainsAllSlice([]string{"z"})}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovP2_56_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_CovP2_56_IsContainsAll", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": col.IsContainsAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.IsContainsAll()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil variadic", actual)
	})
}

func Test_CovP2_57_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_CovP2_57_IsContainsAllLock", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": col.IsContainsAllLock("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": col.IsContainsAllLock()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_CovP2_58_New(t *testing.T) {
	safeTest(t, "Test_CovP2_58_New", func() {
		col := corestr.New.Collection.Strings([]string{"x"})
		newCol := col.New("a", "b")
		actual := args.Map{"result": newCol.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		emptyNew := col.New()
		actual := args.Map{"result": emptyNew.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_59_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_CovP2_59_AddNonEmptyStrings", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddNonEmptyStrings("a", "", "b")
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		col.AddFuncResult()
	})
}

func Test_CovP2_61_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_CovP2_61_AddNonEmptyStringsSlice", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddNonEmptyStringsSlice([]string{"a", "b"})
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_63_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_CovP2_63_ExpandSlicePlusAdd", func() {
		col := corestr.New.Collection.Cap(10)
		col.ExpandSlicePlusAdd(
			[]string{"a,b", "c,d"},
			func(line string) []string { return []string{line + "!"} },
		)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_64_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_CovP2_64_MergeSlicesOfSlice", func() {
		col := corestr.New.Collection.Cap(10)
		col.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovP2_65_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_CovP2_65_GetAllExceptCollection", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := col.GetAllExceptCollection(except)
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// nil
		r2 := col.GetAllExceptCollection(nil)
		actual := args.Map{"result": len(r2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovP2_66_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_CovP2_66_GetAllExcept", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := col.GetAllExcept([]string{"a"})
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		r2 := col.GetAllExcept(nil)
		actual := args.Map{"result": len(r2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_CovP2_67_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_CovP2_67_CharCollectionMap", func() {
		col := corestr.New.Collection.Strings([]string{"abc", "bcd"})
		ccm := col.CharCollectionMap()
		actual := args.Map{"result": ccm == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovP2_68_SummaryString(t *testing.T) {
	safeTest(t, "Test_CovP2_68_SummaryString", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.SummaryString(1)
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovP2_69_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_CovP2_69_SummaryStringWithHeader", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		s := col.SummaryStringWithHeader("Header:")
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		empty := corestr.Empty.Collection()
		s2 := empty.SummaryStringWithHeader("Header:")
		actual := args.Map{"result": s2 == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_CovP2_70_String(t *testing.T) {
	safeTest(t, "Test_CovP2_70_String", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"result": col.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		empty := corestr.Empty.Collection()
		_ = empty.String()
	})
}

func Test_CovP2_71_CsvLines(t *testing.T) {
	safeTest(t, "Test_CovP2_71_CsvLines", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		lines := col.CsvLines()
		actual := args.Map{"result": len(lines) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
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
		actual := args.Map{"result": csv == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.Csv() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
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
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		empty := corestr.Empty.Collection()
		_ = empty.StringLock()
	})
}

func Test_CovP2_76_AddCapacity(t *testing.T) {
	safeTest(t, "Test_CovP2_76_AddCapacity", func() {
		col := corestr.New.Collection.Cap(5)
		col.AddCapacity(10, 20)
		actual := args.Map{"result": col.Capacity() < 30}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 30", actual)
		col.AddCapacity()
	})
}

func Test_CovP2_77_Resize(t *testing.T) {
	safeTest(t, "Test_CovP2_77_Resize", func() {
		col := corestr.New.Collection.Cap(5)
		col.Resize(100)
		actual := args.Map{"result": col.Capacity() < 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 100", actual)
		// no resize needed
		col.Resize(5)
	})
}

func Test_CovP2_78_Joins(t *testing.T) {
	safeTest(t, "Test_CovP2_78_Joins", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.Joins(",")
		actual := args.Map{"result": s != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
		s2 := col.Joins(",", "c")
		actual := args.Map{"result": s2 == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
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
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": len(data) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected data", actual)
	})
}

func Test_CovP2_84_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_CovP2_84_UnmarshalJSON", func() {
		col := corestr.New.Collection.Cap(5)
		err := col.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": col.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		// invalid json
		err2 := col.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"result": err2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP2_88_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovP2_88_ParseInjectUsingJsonMust", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := col.JsonPtr()
		newCol := corestr.New.Collection.Cap(5)
		result := newCol.ParseInjectUsingJsonMust(jsonResult)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovP2_89_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovP2_89_JsonParseSelfInject", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		jsonResult := col.JsonPtr()
		newCol := corestr.New.Collection.Cap(5)
		err := newCol.JsonParseSelfInject(jsonResult)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovP2_90_Clear(t *testing.T) {
	safeTest(t, "Test_CovP2_90_Clear", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		col.Clear()
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovP2_91_Dispose(t *testing.T) {
	safeTest(t, "Test_CovP2_91_Dispose", func() {
		col := corestr.New.Collection.Strings([]string{"a"})
		col.Dispose()
		actual := args.Map{"result": col.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
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
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovP2_95_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovP2_95_Deserialize", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		var target []string
		err := col.Deserialize(&target)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovP2_96_Join(t *testing.T) {
	safeTest(t, "Test_CovP2_96_Join", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"result": col.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.Join(",") != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovP2_97_JoinLine(t *testing.T) {
	safeTest(t, "Test_CovP2_97_JoinLine", func() {
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		s := col.JoinLine()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		empty := corestr.Empty.Collection()
		actual := args.Map{"result": empty.JoinLine() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}
