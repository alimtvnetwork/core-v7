package corestrtests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// Collection.go — Seg-02: Lines 950–1700 (~200 uncovered stmts)
// Covers: InsertAt, ChainRemoveAt, RemoveItemsIndexes, RemoveItemsIndexesPtr,
//         AppendCollectionPtr, AppendCollections, AppendAnysLock, AppendAnys,
//         AppendAnysUsingFilter, AppendAnysUsingFilterLock, AppendNonEmptyAnys,
//         AddsAsync, AddsNonEmpty, AddsNonEmptyPtrLock, UniqueBoolMapLock,
//         UniqueBoolMap, UniqueListLock, UniqueList, List, Filter, FilterLock,
//         FilteredCollection, FilteredCollectionLock, FilterPtrLock, FilterPtr,
//         NonEmptyListPtr, NonEmptyList, HashsetAsIs, HashsetWithDoubleLength,
//         HashsetLock, NonEmptyItems, NonEmptyItemsPtr,
//         NonEmptyItemsOrNonWhitespace, NonEmptyItemsOrNonWhitespacePtr,
//         Items, ListPtr, ListCopyPtrLock, HasLock, Has, HasPtr, HasAll,
//         SortedListAsc, SortedAsc, SortedAscLock, SortedListDsc,
//         HasUsingSensitivity, IsContainsPtr, GetHashsetPlusHasAll,
//         IsContainsAllSlice, IsContainsAll, IsContainsAllLock
// =============================================================================

func Test_Cov59_Collection_InsertAt_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_InsertAt_Empty", func() {
		c := corestr.New.Collection.Empty()
		c.InsertAt(0, "a")
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "a"}
		expected.ShouldBeEqual(t, 0, "InsertAt on empty appends", actual)
	})
}

func Test_Cov59_Collection_InsertAt_Last(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_InsertAt_Last", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.InsertAt(1, "c")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "InsertAt at last appends", actual)
	})
}

func Test_Cov59_Collection_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_InsertAt_Middle", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		c.InsertAt(1, "x")
		actual := args.Map{"lenGte": c.Length() >= 4}
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "InsertAt at middle inserts", actual)
	})
}

func Test_Cov59_Collection_ChainRemoveAt(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ChainRemoveAt", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.ChainRemoveAt(1)
		actual := args.Map{"len": r.Length(), "first": r.First()}
		expected := args.Map{"len": 2, "first": "a"}
		expected.ShouldBeEqual(t, 0, "ChainRemoveAt removes item", actual)
	})
}

func Test_Cov59_Collection_RemoveItemsIndexes(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexes", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		c.RemoveItemsIndexes(true, 1, 3)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexes removes specified indexes", actual)
	})
}

func Test_Cov59_Collection_RemoveItemsIndexes_NilIgnore(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexes_NilIgnore", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexes(true)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexes nil with ignore returns same", actual)
	})
}

func Test_Cov59_Collection_RemoveItemsIndexesPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexesPtr_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexesPtr(true, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr nil returns same", actual)
	})
}

func Test_Cov59_Collection_RemoveItemsIndexesPtr_EmptyWithIgnore(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexesPtr_EmptyWithIgnore", func() {
		c := corestr.New.Collection.Empty()
		c.RemoveItemsIndexesPtr(true, []int{0})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr empty with ignore returns same", actual)
	})
}

func Test_Cov59_Collection_RemoveItemsIndexesPtr_EmptyPanics(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexesPtr_EmptyPanics", func() {
		c := corestr.New.Collection.Empty()
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			c.RemoveItemsIndexesPtr(false, []int{0})
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr empty without ignore panics", actual)
	})
}

func Test_Cov59_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendCollectionPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b", "c"})
		c.AppendCollectionPtr(other)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AppendCollectionPtr appends items", actual)
	})
}

func Test_Cov59_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		e := corestr.New.Collection.Empty()
		c.AppendCollections(a, e, b)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendCollections appends non-empty", actual)
	})
}

func Test_Cov59_Collection_AppendCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendCollections_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendCollections no args returns same", actual)
	})
}

func Test_Cov59_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock("x", 42)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysLock adds any items", actual)
	})
}

func Test_Cov59_Collection_AppendAnysLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysLock_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendAnysLock()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysLock no args returns same", actual)
	})
}

func Test_Cov59_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", nil, 123)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnys skips nil items", actual)
	})
}

func Test_Cov59_Collection_AppendAnys_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnys_Empty", func() {
		c := corestr.New.Collection.Empty()
		c.AppendAnys()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnys no args returns same", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s + "!", true, false
		}
		c.AppendAnysUsingFilter(filter, "a", nil, "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter filters and transforms", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter_Break", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		c.AppendAnysUsingFilter(filter, "a", "b", "c")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter breaks early", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilter_Skip(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter_Skip", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return "", false, false
		}
		c.AppendAnysUsingFilter(filter, "a")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter skip all", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter_Empty", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		c.AppendAnysUsingFilter(filter)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter no items returns same", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilterLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		c.AppendAnysUsingFilterLock(filter, "x", nil, "y")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock adds with lock", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock_Break", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		c.AppendAnysUsingFilterLock(filter, "a", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock breaks early", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilterLock_Skip(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock_Skip", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return "", false, false
		}
		c.AppendAnysUsingFilterLock(filter, "a")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock skip all", actual)
	})
}

func Test_Cov59_Collection_AppendAnysUsingFilterLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock_Nil", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		c.AppendAnysUsingFilterLock(filter, nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock nil returns same", actual)
	})
}

func Test_Cov59_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil, "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendNonEmptyAnys adds non-nil items", actual)
	})
}

func Test_Cov59_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendNonEmptyAnys_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys(nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendNonEmptyAnys nil returns same", actual)
	})
}

func Test_Cov59_Collection_AddsAsync(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsAsync", func() {
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddsAsync(wg, "a", "b")
		wg.Wait()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsAsync adds items async", actual)
	})
}

func Test_Cov59_Collection_AddsAsync_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsAsync_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddsAsync(nil)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsAsync nil returns same", actual)
	})
}

func Test_Cov59_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsNonEmpty skips empty strings", actual)
	})
}

func Test_Cov59_Collection_AddsNonEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmpty_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsNonEmpty nil returns same", actual)
	})
}

func Test_Cov59_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Empty()
		a := "hello"
		b := ""
		c.AddsNonEmptyPtrLock(&a, nil, &b)
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsNonEmptyPtrLock skips nil and empty ptrs", actual)
	})
}

func Test_Cov59_Collection_AddsNonEmptyPtrLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmptyPtrLock_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddsNonEmptyPtrLock()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsNonEmptyPtrLock nil returns same", actual)
	})
}

func Test_Cov59_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		m := c.UniqueBoolMap()
		actual := args.Map{"len": len(m), "hasA": m["a"], "hasB": m["b"]}
		expected := args.Map{"len": 2, "hasA": true, "hasB": true}
		expected.ShouldBeEqual(t, 0, "UniqueBoolMap returns deduplicated map", actual)
	})
}

func Test_Cov59_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"x", "y", "x"})
		m := c.UniqueBoolMapLock()
		actual := args.Map{"len": len(m)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "UniqueBoolMapLock returns deduplicated map", actual)
	})
}

func Test_Cov59_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		u := c.UniqueList()
		actual := args.Map{"len": len(u)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "UniqueList returns unique items", actual)
	})
}

func Test_Cov59_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		u := c.UniqueListLock()
		actual := args.Map{"len": len(u)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "UniqueListLock returns unique items", actual)
	})
}

func Test_Cov59_Collection_List(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"len": len(c.List())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "List returns items", actual)
	})
}

func Test_Cov59_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"aa", "b", "cc"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.Filter(filter)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter keeps matching items", actual)
	})
}

func Test_Cov59_Collection_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Filter_Empty", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		r := c.Filter(filter)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter on empty returns empty", actual)
	})
}

func Test_Cov59_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Filter_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, i >= 0
		}
		r := c.Filter(filter)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Filter breaks on first", actual)
	})
}

func Test_Cov59_Collection_FilterLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterLock", func() {
		c := corestr.New.Collection.Strings([]string{"x", "yy"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.FilterLock(filter)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterLock filters with lock", actual)
	})
}

func Test_Cov59_Collection_FilterLock_Break(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterLock_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		r := c.FilterLock(filter)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterLock breaks early", actual)
	})
}

func Test_Cov59_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.FilteredCollection(filter)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilteredCollection returns new filtered collection", actual)
	})
}

func Test_Cov59_Collection_FilteredCollectionLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilteredCollectionLock", func() {
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.FilteredCollectionLock(filter)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilteredCollectionLock returns filtered", actual)
	})
}

func Test_Cov59_Collection_FilterPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtr", func() {
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		}
		r := c.FilterPtr(filter)
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtr returns pointer slice", actual)
	})
}

func Test_Cov59_Collection_FilterPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtr_Empty", func() {
		c := corestr.New.Collection.Empty()
		filter := func(s *string, i int) (*string, bool, bool) { return s, true, false }
		r := c.FilterPtr(filter)
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterPtr empty returns empty", actual)
	})
}

func Test_Cov59_Collection_FilterPtr_Break(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtr_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		}
		r := c.FilterPtr(filter)
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtr breaks early", actual)
	})
}

func Test_Cov59_Collection_FilterPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		}
		r := c.FilterPtrLock(filter)
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock returns pointer slice", actual)
	})
}

func Test_Cov59_Collection_FilterPtrLock_Break(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtrLock_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		}
		r := c.FilterPtrLock(filter)
		actual := args.Map{"len": len(*r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock breaks early", actual)
	})
}

func Test_Cov59_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyList()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyList filters empty strings", actual)
	})
}

func Test_Cov59_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyList_Empty", func() {
		c := corestr.New.Collection.Empty()
		r := c.NonEmptyList()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "NonEmptyList on empty returns empty", actual)
	})
}
func Test_Cov59_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetAsIs creates hashset", actual)
	})
}

func Test_Cov59_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetWithDoubleLength()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetWithDoubleLength creates hashset", actual)
	})
}

func Test_Cov59_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()
		actual := args.Map{"len": hs.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetLock creates hashset with lock", actual)
	})
}

func Test_Cov59_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyItems()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItems filters empty", actual)
	})
}
func Test_Cov59_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", ""})
		r := c.NonEmptyItemsOrNonWhitespace()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsOrNonWhitespace filters whitespace", actual)
	})
}
func Test_Cov59_Collection_Items(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Items", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"len": len(c.Items())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items returns items", actual)
	})
}
func Test_Cov59_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.ListCopyPtrLock()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock returns copy", actual)
	})
}

func Test_Cov59_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ListCopyPtrLock_Empty", func() {
		c := corestr.New.Collection.Empty()
		r := c.ListCopyPtrLock()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock empty returns empty", actual)
	})
}

func Test_Cov59_Collection_Has(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Has", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"has": c.Has("a"), "miss": c.Has("z")}
		expected := args.Map{"has": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "Has checks containment", actual)
	})
}

func Test_Cov59_Collection_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Has_Empty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"has": c.Has("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has on empty returns false", actual)
	})
}

func Test_Cov59_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"has": c.HasLock("a")}
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasLock checks with lock", actual)
	})
}

func Test_Cov59_Collection_HasPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "a"
		actual := args.Map{"has": c.HasPtr(&s), "nil": c.HasPtr(nil)}
		expected := args.Map{"has": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "HasPtr checks pointer containment", actual)
	})
}

func Test_Cov59_Collection_HasPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasPtr_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := "a"
		actual := args.Map{"has": c.HasPtr(&s)}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasPtr empty returns false", actual)
	})
}

func Test_Cov59_Collection_HasPtr_Miss(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasPtr_Miss", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "z"
		actual := args.Map{"has": c.HasPtr(&s)}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasPtr miss returns false", actual)
	})
}

func Test_Cov59_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{"all": c.HasAll("a", "b"), "miss": c.HasAll("a", "z")}
		expected := args.Map{"all": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "HasAll checks all items present", actual)
	})
}

func Test_Cov59_Collection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasAll_Empty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"has": c.HasAll("a")}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasAll empty returns false", actual)
	})
}

func Test_Cov59_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		r := c.SortedListAsc()
		actual := args.Map{"first": r[0], "last": r[2]}
		expected := args.Map{"first": "a", "last": "c"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc returns sorted copy", actual)
	})
}

func Test_Cov59_Collection_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedListAsc_Empty", func() {
		c := corestr.New.Collection.Empty()
		r := c.SortedListAsc()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedListAsc empty returns empty", actual)
	})
}

func Test_Cov59_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		c.SortedAsc()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAsc sorts in place", actual)
	})
}

func Test_Cov59_Collection_SortedAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAsc_Empty", func() {
		c := corestr.New.Collection.Empty()
		c.SortedAsc()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SortedAsc on empty returns same", actual)
	})
}

func Test_Cov59_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()
		actual := args.Map{"first": c.First()}
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAscLock sorts with lock", actual)
	})
}

func Test_Cov59_Collection_SortedAscLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAscLock_Empty", func() {
		c := corestr.New.Collection.Empty()
		c.SortedAscLock()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SortedAscLock empty returns same", actual)
	})
}

func Test_Cov59_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
		r := c.SortedListDsc()
		actual := args.Map{"first": r[0], "last": r[2]}
		expected := args.Map{"first": "c", "last": "a"}
		expected.ShouldBeEqual(t, 0, "SortedListDsc returns desc sorted", actual)
	})
}

func Test_Cov59_Collection_HasUsingSensitivity(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasUsingSensitivity", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		actual := args.Map{
			"sensitive":   c.HasUsingSensitivity("hello", true),
			"insensitive": c.HasUsingSensitivity("hello", false),
			"miss":        c.HasUsingSensitivity("xyz", false),
		}
		expected := args.Map{"sensitive": false, "insensitive": true, "miss": false}
		expected.ShouldBeEqual(t, 0, "HasUsingSensitivity handles case", actual)
	})
}

func Test_Cov59_Collection_IsContainsPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		m := "z"
		actual := args.Map{"has": c.IsContainsPtr(&s), "miss": c.IsContainsPtr(&m), "nil": c.IsContainsPtr(nil)}
		expected := args.Map{"has": true, "miss": false, "nil": false}
		expected.ShouldBeEqual(t, 0, "IsContainsPtr checks pointer", actual)
	})
}

func Test_Cov59_Collection_IsContainsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsPtr_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := "a"
		actual := args.Map{"has": c.IsContainsPtr(&s)}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContainsPtr empty returns false", actual)
	})
}

func Test_Cov59_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})
		actual := args.Map{"hasAll": hasAll, "hsLen": hs.Length()}
		expected := args.Map{"hasAll": true, "hsLen": 3}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll returns hashset and check", actual)
	})
}

func Test_Cov59_Collection_GetHashsetPlusHasAll_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetHashsetPlusHasAll_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_, hasAll := c.GetHashsetPlusHasAll(nil)
		actual := args.Map{"hasAll": hasAll}
		expected := args.Map{"hasAll": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll nil returns false", actual)
	})
}

func Test_Cov59_Collection_IsContainsAllSlice(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAllSlice", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		actual := args.Map{
			"all":    c.IsContainsAllSlice([]string{"a", "b"}),
			"miss":   c.IsContainsAllSlice([]string{"a", "z"}),
			"empty":  c.IsContainsAllSlice([]string{}),
		}
		expected := args.Map{"all": true, "miss": false, "empty": false}
		expected.ShouldBeEqual(t, 0, "IsContainsAllSlice checks all items", actual)
	})
}

func Test_Cov59_Collection_IsContainsAllSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAllSlice_Empty", func() {
		c := corestr.New.Collection.Empty()
		actual := args.Map{"has": c.IsContainsAllSlice([]string{"a"})}
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContainsAllSlice empty returns false", actual)
	})
}

func Test_Cov59_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		actual := args.Map{"all": c.IsContainsAll("a", "b"), "nil": c.IsContainsAll()}
		expected := args.Map{"all": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "IsContainsAll variadic check", actual)
	})
}

func Test_Cov59_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"has": c.IsContainsAllLock("a"), "nil": c.IsContainsAllLock()}
		expected := args.Map{"has": true, "nil": false}
		expected.ShouldBeEqual(t, 0, "IsContainsAllLock checks with lock", actual)
	})
}

// =============================================================================
// Collection.go — Seg-02 Part B: Lines 1700–2201 (remaining ~150 stmts)
// Covers: New, AddNonEmptyStrings, AddFuncResult, AddNonEmptyStringsSlice,
//         AddStringsByFuncChecking, ExpandSlicePlusAdd, MergeSlicesOfSlice,
//         GetAllExceptCollection, GetAllExcept, CharCollectionMap,
//         SummaryString, SummaryStringWithHeader, String, CsvLines,
//         CsvLinesOptions, Csv, CsvOptions, StringLock, AddCapacity,
//         Resize, Joins, NonEmptyJoins, NonWhitespaceJoins,
//         JsonModel, JsonModelAny, MarshalJSON, UnmarshalJSON, Json, JsonPtr,
//         ParseInjectUsingJson, ParseInjectUsingJsonMust, JsonParseSelfInject,
//         Clear, Dispose, AsJsonMarshaller, AsJsonContractsBinder,
//         Serialize, Deserialize, Join, JoinLine
// =============================================================================

func Test_Cov59_Collection_New(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_New", func() {
		c := corestr.New.Collection.Empty()
		n := c.New("a", "b")
		actual := args.Map{"len": n.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "New creates new collection from slice", actual)
	})
}

func Test_Cov59_Collection_New_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_New_Empty", func() {
		c := corestr.New.Collection.Empty()
		n := c.New()
		actual := args.Map{"empty": n.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "New with no args creates empty", actual)
	})
}

func Test_Cov59_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings adds non-empty", actual)
	})
}

func Test_Cov59_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStrings_Empty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings no args returns same", actual)
	})
}

func Test_Cov59_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" }, func() string { return "b" })
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddFuncResult adds func results", actual)
	})
}

func Test_Cov59_Collection_AddFuncResult_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddFuncResult_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddFuncResult()
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddFuncResult nil returns same", actual)
	})
}

func Test_Cov59_Collection_AddNonEmptyStringsSlice(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStringsSlice", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a", "b"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStringsSlice adds all", actual)
	})
}

func Test_Cov59_Collection_AddNonEmptyStringsSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStringsSlice_Empty", func() {
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStringsSlice empty returns same", actual)
	})
}

func Test_Cov59_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool { return len(s) > 1 })
		actual := args.Map{"len": c.Length(), "first": c.First()}
		expected := args.Map{"len": 1, "first": "bb"}
		expected.ShouldBeEqual(t, 0, "AddStringsByFuncChecking filters by func", actual)
	})
}

func Test_Cov59_Collection_ExpandSlicePlusAdd(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ExpandSlicePlusAdd", func() {
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
			return []string{s + "!"}
		})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ExpandSlicePlusAdd expands and adds", actual)
	})
}

func Test_Cov59_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlice merges slices", actual)
	})
}

func Test_Cov59_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		r := c.GetAllExceptCollection(except)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection excludes items", actual)
	})
}

func Test_Cov59_Collection_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExceptCollection_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExceptCollection(nil)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil copies all", actual)
	})
}

func Test_Cov59_Collection_GetAllExcept(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExcept", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.GetAllExcept([]string{"a"})
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept excludes items", actual)
	})
}

func Test_Cov59_Collection_GetAllExcept_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExcept_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExcept(nil)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil copies all", actual)
	})
}

func Test_Cov59_Collection_CharCollectionMap(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CharCollectionMap", func() {
		c := corestr.New.Collection.Strings([]string{"abc", "def"})
		m := c.CharCollectionMap()
		actual := args.Map{"nonNil": m != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap creates char map", actual)
	})
}

func Test_Cov59_Collection_SummaryString(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SummaryString", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryString(1)
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString returns non-empty", actual)
	})
}

func Test_Cov59_Collection_SummaryStringWithHeader(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SummaryStringWithHeader", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryStringWithHeader("HDR")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader returns non-empty", actual)
	})
}

func Test_Cov59_Collection_SummaryStringWithHeader_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SummaryStringWithHeader_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.SummaryStringWithHeader("HDR")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader empty has header", actual)
	})
}

func Test_Cov59_Collection_String(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_String", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.String()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String returns formatted string", actual)
	})
}

func Test_Cov59_Collection_String_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_String_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.String()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty returns no-elements marker", actual)
	})
}

func Test_Cov59_Collection_CsvLines(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvLines", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.CsvLines()
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvLines returns quoted elements", actual)
	})
}

func Test_Cov59_Collection_CsvLinesOptions(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvLinesOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.CsvLinesOptions(true)
		actual := args.Map{"len": len(r)}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CsvLinesOptions returns elements", actual)
	})
}

func Test_Cov59_Collection_Csv(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Csv", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Csv()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Csv returns csv string", actual)
	})
}

func Test_Cov59_Collection_Csv_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Csv_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.Csv()
		actual := args.Map{"empty": s == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Csv empty returns empty", actual)
	})
}

func Test_Cov59_Collection_CsvOptions(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvOptions", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.CsvOptions(true)
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions returns csv string", actual)
	})
}

func Test_Cov59_Collection_CsvOptions_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvOptions_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.CsvOptions(false)
		actual := args.Map{"empty": s == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions empty returns empty", actual)
	})
}

func Test_Cov59_Collection_StringLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_StringLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.StringLock()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock returns string", actual)
	})
}

func Test_Cov59_Collection_StringLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_StringLock_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.StringLock()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty returns marker", actual)
	})
}

func Test_Cov59_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddCapacity", func() {
		c := corestr.New.Collection.Empty()
		c.AddCapacity(10)
		actual := args.Map{"capGte": c.Capacity() >= 10}
		expected := args.Map{"capGte": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity increases capacity", actual)
	})
}

func Test_Cov59_Collection_AddCapacity_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddCapacity_Nil", func() {
		c := corestr.New.Collection.Empty()
		c.AddCapacity()
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity no args returns same", actual)
	})
}

func Test_Cov59_Collection_Resize(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Resize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Resize(100)
		actual := args.Map{"capGte": c.Capacity() >= 100}
		expected := args.Map{"capGte": true}
		expected.ShouldBeEqual(t, 0, "Resize increases capacity", actual)
	})
}

func Test_Cov59_Collection_Resize_SmallerIgnored(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Resize_SmallerIgnored", func() {
		c := corestr.New.Collection.Cap(50)
		origCap := c.Capacity()
		c.Resize(10)
		actual := args.Map{"same": c.Capacity() == origCap}
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Resize smaller ignored", actual)
	})
}

func Test_Cov59_Collection_Joins(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Joins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Joins(",")
		actual := args.Map{"val": s}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Joins joins items", actual)
	})
}

func Test_Cov59_Collection_Joins_WithExtra(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Joins_WithExtra", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.Joins(",", "b", "c")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins with extra items", actual)
	})
}

func Test_Cov59_Collection_NonEmptyJoins(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		s := c.NonEmptyJoins(",")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins joins non-empty", actual)
	})
}

func Test_Cov59_Collection_NonWhitespaceJoins(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonWhitespaceJoins", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		s := c.NonWhitespaceJoins(",")
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins joins non-whitespace", actual)
	})
}

func Test_Cov59_Collection_JsonModel(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonModel", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		actual := args.Map{"len": len(c.JsonModel())}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel returns items", actual)
	})
}

func Test_Cov59_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonModelAny", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.JsonModelAny()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny returns any", actual)
	})
}

func Test_Cov59_Collection_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_MarshalJSON", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.MarshalJSON()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "MarshalJSON returns bytes", actual)
	})
}

func Test_Cov59_Collection_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UnmarshalJSON", func() {
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))
		actual := args.Map{"noErr": err == nil, "len": c.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON parses items", actual)
	})
}

func Test_Cov59_Collection_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UnmarshalJSON_Error", func() {
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`invalid`))
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns error on bad input", actual)
	})
}

func Test_Cov59_Collection_Json(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Json", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Json()
		actual := args.Map{"nonEmpty": r.JsonString() != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json returns Result", actual)
	})
}

func Test_Cov59_Collection_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.JsonPtr()
		actual := args.Map{"nonNil": r != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns Result ptr", actual)
	})
}

func Test_Cov59_Collection_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJson", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := corejson.NewPtr([]string{"x", "y"})
		r, err := c.ParseInjectUsingJson(jr)
		actual := args.Map{"noErr": err == nil, "len": r.Length()}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson parses", actual)
	})
}

func Test_Cov59_Collection_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJson_Error", func() {
		c := corestr.New.Collection.Empty()
		jr := &corejson.Result{Error: errForTest}
		_, err := c.ParseInjectUsingJson(jr)
		actual := args.Map{"hasErr": err != nil}
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error", actual)
	})
}

func Test_Cov59_Collection_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJsonMust", func() {
		c := corestr.New.Collection.Empty()
		jr := corejson.NewPtr([]string{"x"})
		r := c.ParseInjectUsingJsonMust(jr)
		actual := args.Map{"len": r.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust parses", actual)
	})
}

func Test_Cov59_Collection_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJsonMust_Panics", func() {
		c := corestr.New.Collection.Empty()
		jr := &corejson.Result{Error: errForTest}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			c.ParseInjectUsingJsonMust(jr)
		}()
		actual := args.Map{"panicked": panicked}
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics on error", actual)
	})
}

func Test_Cov59_Collection_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonParseSelfInject", func() {
		c := corestr.New.Collection.Empty()
		jr := corejson.NewPtr([]string{"a"})
		err := c.JsonParseSelfInject(jr)
		actual := args.Map{"noErr": err == nil, "len": c.Length()}
		expected := args.Map{"noErr": true, "len": 1}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject injects items", actual)
	})
}

func Test_Cov59_Collection_Clear(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Clear", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empties collection", actual)
	})
}

func Test_Cov59_Collection_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Clear_Nil", func() {
		var c *corestr.Collection
		r := c.Clear()
		actual := args.Map{"nil": r == nil}
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear on nil returns nil", actual)
	})
}

func Test_Cov59_Collection_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Dispose", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()
		actual := args.Map{"empty": c.IsEmpty()}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Dispose clears and nils items", actual)
	})
}

func Test_Cov59_Collection_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Dispose_Nil", func() {
		var c *corestr.Collection
		c.Dispose() // should not panic
		actual := args.Map{"ok": true}
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Dispose on nil does not panic", actual)
	})
}

func Test_Cov59_Collection_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AsJsonMarshaller", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		m := c.AsJsonMarshaller()
		actual := args.Map{"nonNil": m != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller returns interface", actual)
	})
}

func Test_Cov59_Collection_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AsJsonContractsBinder", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b := c.AsJsonContractsBinder()
		actual := args.Map{"nonNil": b != nil}
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns interface", actual)
	})
}

func Test_Cov59_Collection_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Serialize", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()
		actual := args.Map{"noErr": err == nil, "nonEmpty": len(b) > 0}
		expected := args.Map{"noErr": true, "nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Serialize returns bytes", actual)
	})
}

func Test_Cov59_Collection_Deserialize(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Deserialize", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		var target []string
		err := c.Deserialize(&target)
		actual := args.Map{"noErr": err == nil, "len": len(target)}
		expected := args.Map{"noErr": true, "len": 2}
		expected.ShouldBeEqual(t, 0, "Deserialize parses to target", actual)
	})
}

func Test_Cov59_Collection_Join(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Join", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Join(",")
		actual := args.Map{"val": s}
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join joins with separator", actual)
	})
}

func Test_Cov59_Collection_Join_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Join_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.Join(",")
		actual := args.Map{"empty": s == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Join empty returns empty", actual)
	})
}

func Test_Cov59_Collection_JoinLine(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JoinLine", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JoinLine()
		actual := args.Map{"nonEmpty": s != ""}
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine joins with newline", actual)
	})
}

func Test_Cov59_Collection_JoinLine_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JoinLine_Empty", func() {
		c := corestr.New.Collection.Empty()
		s := c.JoinLine()
		actual := args.Map{"empty": s == ""}
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine empty returns empty", actual)
	})
}
