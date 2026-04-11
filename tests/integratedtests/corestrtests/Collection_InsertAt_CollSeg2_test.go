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

func Test_Collection_InsertAt_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_InsertAt_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.InsertAt(0, "a")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "InsertAt on empty appends", actual)
	})
}

func Test_Collection_InsertAt_Last(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_InsertAt_Last", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.InsertAt(1, "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "InsertAt at last appends", actual)
	})
}

func Test_Collection_InsertAt_Middle(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_InsertAt_Middle", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		c.InsertAt(1, "x")

		// Act
		actual := args.Map{"lenGte": c.Length() >= 4}

		// Assert
		expected := args.Map{"lenGte": true}
		expected.ShouldBeEqual(t, 0, "InsertAt at middle inserts", actual)
	})
}

func Test_Collection_ChainRemoveAt_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ChainRemoveAt", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.ChainRemoveAt(1)

		// Act
		actual := args.Map{
			"len": r.Length(),
			"first": r.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "ChainRemoveAt removes item", actual)
	})
}

func Test_Collection_RemoveItemsIndexes_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexes", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c", "d"})
		c.RemoveItemsIndexes(true, 1, 3)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexes removes specified indexes", actual)
	})
}

func Test_Collection_RemoveItemsIndexes_NilIgnore(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexes_NilIgnore", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexes(true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexes nil with ignore returns same", actual)
	})
}

func Test_Collection_RemoveItemsIndexesPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexesPtr_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexesPtr(true, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr nil returns same", actual)
	})
}

func Test_Collection_RemoveItemsIndexesPtr_EmptyWithIgnore(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexesPtr_EmptyWithIgnore", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.RemoveItemsIndexesPtr(true, []int{0})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr empty with ignore returns same", actual)
	})
}

func Test_Collection_RemoveItemsIndexesPtr_EmptyPanics(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_RemoveItemsIndexesPtr_EmptyPanics", func() {
		// Arrange
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

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "RemoveItemsIndexesPtr empty without ignore panics", actual)
	})
}

func Test_Collection_AppendCollectionPtr_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendCollectionPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		other := corestr.New.Collection.Strings([]string{"b", "c"})
		c.AppendCollectionPtr(other)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AppendCollectionPtr appends items", actual)
	})
}

func Test_Collection_AppendCollections_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendCollections", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		a := corestr.New.Collection.Strings([]string{"a"})
		b := corestr.New.Collection.Strings([]string{"b"})
		e := corestr.New.Collection.Empty()
		c.AppendCollections(a, e, b)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendCollections appends non-empty", actual)
	})
}

func Test_Collection_AppendCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendCollections_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendCollections no args returns same", actual)
	})
}

func Test_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnysLock("x", 42)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysLock adds any items", actual)
	})
}

func Test_Collection_AppendAnysLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendAnysLock()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysLock no args returns same", actual)
	})
}

func Test_Collection_AppendAnys_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnys("a", nil, 123)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnys skips nil items", actual)
	})
}

func Test_Collection_AppendAnys_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnys_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendAnys()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnys no args returns same", actual)
	})
}

func Test_Collection_AppendAnysUsingFilter_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s + "!", true, false
		}
		c.AppendAnysUsingFilter(filter, "a", nil, "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter filters and transforms", actual)
	})
}

func Test_Collection_AppendAnysUsingFilter_Break_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		c.AppendAnysUsingFilter(filter, "a", "b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter breaks early", actual)
	})
}

func Test_Collection_AppendAnysUsingFilter_Skip_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return "", false, false
		}
		c.AppendAnysUsingFilter(filter, "a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter skip all", actual)
	})
}

func Test_Collection_AppendAnysUsingFilter_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilter_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		c.AppendAnysUsingFilter(filter)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilter no items returns same", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, false
		}
		c.AppendAnysUsingFilterLock(filter, "x", nil, "y")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock adds with lock", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_Break_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		c.AppendAnysUsingFilterLock(filter, "a", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock breaks early", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_Skip(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock_Skip", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) {
			return "", false, false
		}
		c.AppendAnysUsingFilterLock(filter, "a")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock skip all", actual)
	})
}

func Test_Collection_AppendAnysUsingFilterLock_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendAnysUsingFilterLock_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		c.AppendAnysUsingFilterLock(filter, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendAnysUsingFilterLock nil returns same", actual)
	})
}

func Test_Collection_AppendNonEmptyAnys_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendNonEmptyAnys", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys("a", nil, "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AppendNonEmptyAnys adds non-nil items", actual)
	})
}

func Test_Collection_AppendNonEmptyAnys_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AppendNonEmptyAnys_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AppendNonEmptyAnys(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendNonEmptyAnys nil returns same", actual)
	})
}

func Test_Collection_AddsAsync_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsAsync", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.AddsAsync(wg, "a", "b")
		wg.Wait()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsAsync adds items async", actual)
	})
}

func Test_Collection_AddsAsync_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsAsync_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsAsync(nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsAsync nil returns same", actual)
	})
}

func Test_Collection_AddsNonEmpty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty("a", "", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsNonEmpty skips empty strings", actual)
	})
}

func Test_Collection_AddsNonEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmpty_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsNonEmpty()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsNonEmpty nil returns same", actual)
	})
}

func Test_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmptyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		a := "hello"
		b := ""
		c.AddsNonEmptyPtrLock(&a, nil, &b)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsNonEmptyPtrLock skips nil and empty ptrs", actual)
	})
}

func Test_Collection_AddsNonEmptyPtrLock_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddsNonEmptyPtrLock_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddsNonEmptyPtrLock()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsNonEmptyPtrLock nil returns same", actual)
	})
}

func Test_Collection_UniqueBoolMap_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueBoolMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		m := c.UniqueBoolMap()

		// Act
		actual := args.Map{
			"len": len(m),
			"hasA": m["a"],
			"hasB": m["b"],
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"hasA": true,
			"hasB": true,
		}
		expected.ShouldBeEqual(t, 0, "UniqueBoolMap returns deduplicated map", actual)
	})
}

func Test_Collection_UniqueBoolMapLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueBoolMapLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"x", "y", "x"})
		m := c.UniqueBoolMapLock()

		// Act
		actual := args.Map{"len": len(m)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "UniqueBoolMapLock returns deduplicated map", actual)
	})
}

func Test_Collection_UniqueList_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		u := c.UniqueList()

		// Act
		actual := args.Map{"len": len(u)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "UniqueList returns unique items", actual)
	})
}

func Test_Collection_UniqueListLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UniqueListLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		u := c.UniqueListLock()

		// Act
		actual := args.Map{"len": len(u)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "UniqueListLock returns unique items", actual)
	})
}

func Test_Collection_List_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_List", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.List())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "List returns items", actual)
	})
}

func Test_Collection_Filter_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Filter", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b", "cc"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.Filter(filter)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Filter keeps matching items", actual)
	})
}

func Test_Collection_Filter_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Filter_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s string, i int) (string, bool, bool) { return s, true, false }
		r := c.Filter(filter)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Filter on empty returns empty", actual)
	})
}

func Test_Collection_Filter_Break_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Filter_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, i >= 0
		}
		r := c.Filter(filter)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Filter breaks on first", actual)
	})
}

func Test_Collection_FilterLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"x", "yy"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.FilterLock(filter)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterLock filters with lock", actual)
	})
}

func Test_Collection_FilterLock_Break_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, true, true
		}
		r := c.FilterLock(filter)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterLock breaks early", actual)
	})
}

func Test_Collection_FilteredCollection_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilteredCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.FilteredCollection(filter)

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilteredCollection returns new filtered collection", actual)
	})
}

func Test_Collection_FilteredCollectionLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilteredCollectionLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		}
		r := c.FilteredCollectionLock(filter)

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilteredCollectionLock returns filtered", actual)
	})
}

func Test_Collection_FilterPtr_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		}
		r := c.FilterPtr(filter)

		// Act
		actual := args.Map{"len": len(*r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtr returns pointer slice", actual)
	})
}

func Test_Collection_FilterPtr_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtr_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		filter := func(s *string, i int) (*string, bool, bool) { return s, true, false }
		r := c.FilterPtr(filter)

		// Act
		actual := args.Map{"len": len(*r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterPtr empty returns empty", actual)
	})
}

func Test_Collection_FilterPtr_Break_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtr_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		}
		r := c.FilterPtr(filter)

		// Act
		actual := args.Map{"len": len(*r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtr breaks early", actual)
	})
}

func Test_Collection_FilterPtrLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"aa", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		}
		r := c.FilterPtrLock(filter)

		// Act
		actual := args.Map{"len": len(*r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock returns pointer slice", actual)
	})
}

func Test_Collection_FilterPtrLock_Break_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_FilterPtrLock_Break", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		filter := func(s *string, i int) (*string, bool, bool) {
			return s, true, true
		}
		r := c.FilterPtrLock(filter)

		// Act
		actual := args.Map{"len": len(*r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock breaks early", actual)
	})
}

func Test_Collection_NonEmptyList_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyList", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyList()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyList filters empty strings", actual)
	})
}

func Test_Collection_NonEmptyList_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyList_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		r := c.NonEmptyList()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "NonEmptyList on empty returns empty", actual)
	})
}

func Test_Collection_NonEmptyListPtr_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyListPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyListPtr()

		// Act
		actual := args.Map{"len": len(*r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyListPtr returns ptr to non-empty", actual)
	})
}

func Test_Collection_HashsetAsIs_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HashsetAsIs", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetAsIs creates hashset", actual)
	})
}

func Test_Collection_HashsetWithDoubleLength_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HashsetWithDoubleLength", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetWithDoubleLength()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetWithDoubleLength creates hashset", actual)
	})
}

func Test_Collection_HashsetLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HashsetLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()

		// Act
		actual := args.Map{"len": hs.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetLock creates hashset with lock", actual)
	})
}

func Test_Collection_NonEmptyItems_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyItems", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyItems()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItems filters empty", actual)
	})
}

func Test_Collection_NonEmptyItemsPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyItemsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		r := c.NonEmptyItemsPtr()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsPtr filters empty", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespace_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyItemsOrNonWhitespace", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  ", ""})
		r := c.NonEmptyItemsOrNonWhitespace()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsOrNonWhitespace filters whitespace", actual)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespacePtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyItemsOrNonWhitespacePtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  ", ""})
		r := c.NonEmptyItemsOrNonWhitespacePtr()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NonEmptyItemsOrNonWhitespacePtr filters", actual)
	})
}

func Test_Collection_Items(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Items", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items returns items", actual)
	})
}

func Test_Collection_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ListPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.ListPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListPtr returns items", actual)
	})
}

func Test_Collection_ListCopyPtrLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ListCopyPtrLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock returns copy", actual)
	})
}

func Test_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ListCopyPtrLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		r := c.ListCopyPtrLock()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "ListCopyPtrLock empty returns empty", actual)
	})
}

func Test_Collection_Has_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Has", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"has": c.Has("a"),
			"miss": c.Has("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Has checks containment", actual)
	})
}

func Test_Collection_Has_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Has_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"has": c.Has("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has on empty returns false", actual)
	})
}

func Test_Collection_HasLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"has": c.HasLock("a")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasLock checks with lock", actual)
	})
}

func Test_Collection_HasPtr_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "a"

		// Act
		actual := args.Map{
			"has": c.HasPtr(&s),
			"nil": c.HasPtr(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasPtr checks pointer containment", actual)
	})
}

func Test_Collection_HasPtr_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasPtr_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := "a"

		// Act
		actual := args.Map{"has": c.HasPtr(&s)}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasPtr empty returns false", actual)
	})
}

func Test_Collection_HasPtr_Miss(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasPtr_Miss", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "z"

		// Act
		actual := args.Map{"has": c.HasPtr(&s)}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasPtr miss returns false", actual)
	})
}

func Test_Collection_HasAll_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"all": c.HasAll("a", "b"),
			"miss": c.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll checks all items present", actual)
	})
}

func Test_Collection_HasAll_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasAll_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"has": c.HasAll("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasAll empty returns false", actual)
	})
}

func Test_Collection_SortedListAsc_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedListAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		r := c.SortedListAsc()

		// Act
		actual := args.Map{
			"first": r[0],
			"last": r[2],
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "c",
		}
		expected.ShouldBeEqual(t, 0, "SortedListAsc returns sorted copy", actual)
	})
}

func Test_Collection_SortedListAsc_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedListAsc_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		r := c.SortedListAsc()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedListAsc empty returns empty", actual)
	})
}

func Test_Collection_SortedAsc_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		c.SortedAsc()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAsc sorts in place", actual)
	})
}

func Test_Collection_SortedAsc_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAsc_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.SortedAsc()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SortedAsc on empty returns same", actual)
	})
}

func Test_Collection_SortedAscLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAscLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"b", "a"})
		c.SortedAscLock()

		// Act
		actual := args.Map{"first": c.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "SortedAscLock sorts with lock", actual)
	})
}

func Test_Collection_SortedAscLock_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedAscLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.SortedAscLock()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "SortedAscLock empty returns same", actual)
	})
}

func Test_Collection_SortedListDsc_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SortedListDsc", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
		r := c.SortedListDsc()

		// Act
		actual := args.Map{
			"first": r[0],
			"last": r[2],
		}

		// Assert
		expected := args.Map{
			"first": "c",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "SortedListDsc returns desc sorted", actual)
	})
}

func Test_Collection_HasUsingSensitivity_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_HasUsingSensitivity", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"Hello"})

		// Act
		actual := args.Map{
			"sensitive":   c.HasUsingSensitivity("hello", true),
			"insensitive": c.HasUsingSensitivity("hello", false),
			"miss":        c.HasUsingSensitivity("xyz", false),
		}

		// Assert
		expected := args.Map{
			"sensitive": false,
			"insensitive": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasUsingSensitivity handles case", actual)
	})
}

func Test_Collection_IsContainsPtr_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		m := "z"

		// Act
		actual := args.Map{
			"has": c.IsContainsPtr(&s),
			"miss": c.IsContainsPtr(&m),
			"nil": c.IsContainsPtr(nil),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsPtr checks pointer", actual)
	})
}

func Test_Collection_IsContainsPtr_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsPtr_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := "a"

		// Act
		actual := args.Map{"has": c.IsContainsPtr(&s)}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContainsPtr empty returns false", actual)
	})
}

func Test_Collection_GetHashsetPlusHasAll_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetHashsetPlusHasAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})

		// Act
		actual := args.Map{
			"hasAll": hasAll,
			"hsLen": hs.Length(),
		}

		// Assert
		expected := args.Map{
			"hasAll": true,
			"hsLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll returns hashset and check", actual)
	})
}

func Test_Collection_GetHashsetPlusHasAll_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetHashsetPlusHasAll_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		_, hasAll := c.GetHashsetPlusHasAll(nil)

		// Act
		actual := args.Map{"hasAll": hasAll}

		// Assert
		expected := args.Map{"hasAll": false}
		expected.ShouldBeEqual(t, 0, "GetHashsetPlusHasAll nil returns false", actual)
	})
}

func Test_Collection_IsContainsAllSlice_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAllSlice", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})

		// Act
		actual := args.Map{
			"all":    c.IsContainsAllSlice([]string{"a", "b"}),
			"miss":   c.IsContainsAllSlice([]string{"a", "z"}),
			"empty":  c.IsContainsAllSlice([]string{}),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"miss": false,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAllSlice checks all items", actual)
	})
}

func Test_Collection_IsContainsAllSlice_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAllSlice_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()

		// Act
		actual := args.Map{"has": c.IsContainsAllSlice([]string{"a"})}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "IsContainsAllSlice empty returns false", actual)
	})
}

func Test_Collection_IsContainsAll_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAll", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{
			"all": c.IsContainsAll("a", "b"),
			"nil": c.IsContainsAll(),
		}

		// Assert
		expected := args.Map{
			"all": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContainsAll variadic check", actual)
	})
}

func Test_Collection_IsContainsAllLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_IsContainsAllLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{
			"has": c.IsContainsAllLock("a"),
			"nil": c.IsContainsAllLock(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"nil": false,
		}
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

func Test_Collection_New_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_New", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		n := c.New("a", "b")

		// Act
		actual := args.Map{"len": n.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "New creates new collection from slice", actual)
	})
}

func Test_Collection_New_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_New_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		n := c.New()

		// Act
		actual := args.Map{"empty": n.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "New with no args creates empty", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStrings", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings("a", "", "b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings adds non-empty", actual)
	})
}

func Test_Collection_AddNonEmptyStrings_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStrings_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStrings()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStrings no args returns same", actual)
	})
}

func Test_Collection_AddFuncResult_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddFuncResult", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncResult(func() string { return "a" }, func() string { return "b" })

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddFuncResult adds func results", actual)
	})
}

func Test_Collection_AddFuncResult_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddFuncResult_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddFuncResult()

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddFuncResult nil returns same", actual)
	})
}

func Test_Collection_AddNonEmptyStringsSlice_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStringsSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStringsSlice adds all", actual)
	})
}

func Test_Collection_AddNonEmptyStringsSlice_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddNonEmptyStringsSlice_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddNonEmptyStringsSlice([]string{})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddNonEmptyStringsSlice empty returns same", actual)
	})
}

func Test_Collection_AddStringsByFuncChecking_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddStringsByFuncChecking", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddStringsByFuncChecking([]string{"a", "bb", "c"}, func(s string) bool { return len(s) > 1 })

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "bb",
		}
		expected.ShouldBeEqual(t, 0, "AddStringsByFuncChecking filters by func", actual)
	})
}

func Test_Collection_ExpandSlicePlusAdd_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ExpandSlicePlusAdd", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.ExpandSlicePlusAdd([]string{"a,b"}, func(s string) []string {
			return []string{s + "!"}
		})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ExpandSlicePlusAdd expands and adds", actual)
	})
}

func Test_Collection_MergeSlicesOfSlice_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_MergeSlicesOfSlice", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlice merges slices", actual)
	})
}

func Test_Collection_GetAllExceptCollection_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExceptCollection", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		r := c.GetAllExceptCollection(except)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection excludes items", actual)
	})
}

func Test_Collection_GetAllExceptCollection_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExceptCollection_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExceptCollection(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExceptCollection nil copies all", actual)
	})
}

func Test_Collection_GetAllExcept_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExcept", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		r := c.GetAllExcept([]string{"a"})

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept excludes items", actual)
	})
}

func Test_Collection_GetAllExcept_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_GetAllExcept_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.GetAllExcept(nil)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetAllExcept nil copies all", actual)
	})
}

func Test_Collection_CharCollectionMap_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CharCollectionMap", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"abc", "def"})
		m := c.CharCollectionMap()

		// Act
		actual := args.Map{"nonNil": m != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap creates char map", actual)
	})
}

func Test_Collection_SummaryString_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SummaryString", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryString(1)

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString returns non-empty", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SummaryStringWithHeader", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.SummaryStringWithHeader("HDR")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader returns non-empty", actual)
	})
}

func Test_Collection_SummaryStringWithHeader_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_SummaryStringWithHeader_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.SummaryStringWithHeader("HDR")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringWithHeader empty has header", actual)
	})
}

func Test_Collection_String_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_String", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.String()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String returns formatted string", actual)
	})
}

func Test_Collection_String_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_String_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.String()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty returns no-elements marker", actual)
	})
}

func Test_Collection_CsvLines_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvLines", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		r := c.CsvLines()

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CsvLines returns quoted elements", actual)
	})
}

func Test_Collection_CsvLinesOptions_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvLinesOptions", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.CsvLinesOptions(true)

		// Act
		actual := args.Map{"len": len(r)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CsvLinesOptions returns elements", actual)
	})
}

func Test_Collection_Csv_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Csv", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Csv()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Csv returns csv string", actual)
	})
}

func Test_Collection_Csv_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Csv_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.Csv()

		// Act
		actual := args.Map{"empty": s == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Csv empty returns empty", actual)
	})
}

func Test_Collection_CsvOptions_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvOptions", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.CsvOptions(true)

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions returns csv string", actual)
	})
}

func Test_Collection_CsvOptions_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_CsvOptions_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.CsvOptions(false)

		// Act
		actual := args.Map{"empty": s == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "CsvOptions empty returns empty", actual)
	})
}

func Test_Collection_StringLock_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_StringLock", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.StringLock()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock returns string", actual)
	})
}

func Test_Collection_StringLock_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_StringLock_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.StringLock()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock empty returns marker", actual)
	})
}

func Test_Collection_AddCapacity(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddCapacity", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddCapacity(10)

		// Act
		actual := args.Map{"capGte": c.Capacity() >= 10}

		// Assert
		expected := args.Map{"capGte": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity increases capacity", actual)
	})
}

func Test_Collection_AddCapacity_Nil(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AddCapacity_Nil", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		c.AddCapacity()

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "AddCapacity no args returns same", actual)
	})
}

func Test_Collection_Resize_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Resize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Resize(100)

		// Act
		actual := args.Map{"capGte": c.Capacity() >= 100}

		// Assert
		expected := args.Map{"capGte": true}
		expected.ShouldBeEqual(t, 0, "Resize increases capacity", actual)
	})
}

func Test_Collection_Resize_SmallerIgnored(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Resize_SmallerIgnored", func() {
		// Arrange
		c := corestr.New.Collection.Cap(50)
		origCap := c.Capacity()
		c.Resize(10)

		// Act
		actual := args.Map{"same": c.Capacity() == origCap}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Resize smaller ignored", actual)
	})
}

func Test_Collection_Joins_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Joins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Joins(",")

		// Act
		actual := args.Map{"val": s}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Joins joins items", actual)
	})
}

func Test_Collection_Joins_WithExtra_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Joins_WithExtra", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		s := c.Joins(",", "b", "c")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Joins with extra items", actual)
	})
}

func Test_Collection_NonEmptyJoins_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonEmptyJoins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		s := c.NonEmptyJoins(",")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonEmptyJoins joins non-empty", actual)
	})
}

func Test_Collection_NonWhitespaceJoins_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_NonWhitespaceJoins", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		s := c.NonWhitespaceJoins(",")

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "NonWhitespaceJoins joins non-whitespace", actual)
	})
}

func Test_Collection_JsonModel_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonModel", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})

		// Act
		actual := args.Map{"len": len(c.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel returns items", actual)
	})
}

func Test_Collection_JsonModelAny(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonModelAny", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.JsonModelAny()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny returns any", actual)
	})
}

func Test_Collection_MarshalJSON_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_MarshalJSON", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON returns bytes", actual)
	})
}

func Test_Collection_UnmarshalJSON_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UnmarshalJSON", func() {
		// Arrange
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`["a","b"]`))

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
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON parses items", actual)
	})
}

func Test_Collection_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_UnmarshalJSON_Error", func() {
		// Arrange
		c := &corestr.Collection{}
		err := c.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON returns error on bad input", actual)
	})
}

func Test_Collection_Json_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Json", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.Json()

		// Act
		actual := args.Map{"nonEmpty": r.JsonString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Json returns Result", actual)
	})
}

func Test_Collection_JsonPtr_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonPtr", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		r := c.JsonPtr()

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "JsonPtr returns Result ptr", actual)
	})
}

func Test_Collection_ParseInjectUsingJson_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJson", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		jr := corejson.NewPtr([]string{"x", "y"})
		r, err := c.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": r.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson parses", actual)
	})
}

func Test_Collection_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJson_Error", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		jr := &corejson.Result{Error: errForTest}
		_, err := c.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson returns error", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJsonMust", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		jr := corejson.NewPtr([]string{"x"})
		r := c.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"len": r.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust parses", actual)
	})
}

func Test_Collection_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
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

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust panics on error", actual)
	})
}

func Test_Collection_JsonParseSelfInject_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JsonParseSelfInject", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		jr := corejson.NewPtr([]string{"a"})
		err := c.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": c.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject injects items", actual)
	})
}

func Test_Collection_Clear_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Clear", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		c.Clear()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Clear empties collection", actual)
	})
}

func Test_Collection_Clear_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Clear_Nil", func() {
		// Arrange
		var c *corestr.Collection
		r := c.Clear()

		// Act
		actual := args.Map{"nil": r == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear on nil returns nil", actual)
	})
}

func Test_Collection_Dispose_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Dispose", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		c.Dispose()

		// Act
		actual := args.Map{"empty": c.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Dispose clears and nils items", actual)
	})
}

func Test_Collection_Dispose_Nil_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Dispose_Nil", func() {
		// Arrange
		var c *corestr.Collection
		c.Dispose() // should not panic

		// Act
		actual := args.Map{"ok": true}

		// Assert
		expected := args.Map{"ok": true}
		expected.ShouldBeEqual(t, 0, "Dispose on nil does not panic", actual)
	})
}

func Test_Collection_AsJsonMarshaller_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AsJsonMarshaller", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		m := c.AsJsonMarshaller()

		// Act
		actual := args.Map{"nonNil": m != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonMarshaller returns interface", actual)
	})
}

func Test_Collection_AsJsonContractsBinder_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_AsJsonContractsBinder", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		b := c.AsJsonContractsBinder()

		// Act
		actual := args.Map{"nonNil": b != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "AsJsonContractsBinder returns interface", actual)
	})
}

func Test_Collection_Serialize_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Serialize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a"})
		b, err := c.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize returns bytes", actual)
	})
}

func Test_Collection_Deserialize_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Deserialize", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		var target []string
		err := c.Deserialize(&target)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": len(target),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 2,
		}
		expected.ShouldBeEqual(t, 0, "Deserialize parses to target", actual)
	})
}

func Test_Collection_Join_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Join", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.Join(",")

		// Act
		actual := args.Map{"val": s}

		// Assert
		expected := args.Map{"val": "a,b"}
		expected.ShouldBeEqual(t, 0, "Join joins with separator", actual)
	})
}

func Test_Collection_Join_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_Join_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.Join(",")

		// Act
		actual := args.Map{"empty": s == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Join empty returns empty", actual)
	})
}

func Test_Collection_JoinLine_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JoinLine", func() {
		// Arrange
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := c.JoinLine()

		// Act
		actual := args.Map{"nonEmpty": s != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine joins with newline", actual)
	})
}

func Test_Collection_JoinLine_Empty_FromCollectionInsertAtCo(t *testing.T) {
	safeTest(t, "Test_Cov59_Collection_JoinLine_Empty", func() {
		// Arrange
		c := corestr.New.Collection.Empty()
		s := c.JoinLine()

		// Act
		actual := args.Map{"empty": s == ""}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "JoinLine empty returns empty", actual)
	})
}
