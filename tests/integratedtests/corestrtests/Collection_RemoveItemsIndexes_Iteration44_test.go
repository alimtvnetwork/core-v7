package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════
// Collection — remove, append, filter, unique, sort, search
// ═══════════════════════════════════════════════════════════════

func Test_Collection_RemoveItemsIndexes_Valid(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_RemoveItemsIndexes_Valid", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		c.RemoveItemsIndexes(true, 1)
		tc := caseV1Compat{Name: "RemoveItemsIndexes", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_RemoveItemsIndexes_NilIgnore(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_RemoveItemsIndexes_NilIgnore", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.RemoveItemsIndexes(true)
		tc := caseV1Compat{Name: "RemoveItemsIndexes nil ignore", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendCollectionPtr(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendCollectionPtr", func() {
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b", "c"})
		c1.AppendCollectionPtr(c2)
		tc := caseV1Compat{Name: "AppendCollectionPtr", Expected: 3, Actual: c1.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendCollections(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendCollections", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Strings([]string{"b"})
		c3 := corestr.New.Collection.Strings([]string{"c"})
		c.AppendCollections(c2, c3)
		tc := caseV1Compat{Name: "AppendCollections", Expected: 3, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendCollections_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendCollections_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		c.AppendCollections()
		tc := caseV1Compat{Name: "AppendCollections empty", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnys(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnys", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys("hello", 42)
		tc := caseV1Compat{Name: "AppendAnys", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnys_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnys_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys()
		tc := caseV1Compat{Name: "AppendAnys empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnys_NilSkip(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnys_NilSkip", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnys(nil, "ok")
		tc := caseV1Compat{Name: "AppendAnys nil skip", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnysLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnysLock", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock("hello")
		tc := caseV1Compat{Name: "AppendAnysLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnysLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnysLock_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysLock()
		tc := caseV1Compat{Name: "AppendAnysLock empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnysUsingFilter(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnysUsingFilter", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		}, "a", "b")
		tc := caseV1Compat{Name: "AppendAnysUsingFilter", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnysUsingFilter_Break(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnysUsingFilter_Break", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilter(func(s string, i int) (string, bool, bool) {
			return s, true, true
		}, "a", "b")
		tc := caseV1Compat{Name: "AppendAnysUsingFilter break", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendAnysUsingFilter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendAnysUsingFilter_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendAnysUsingFilter(nil)
		tc := caseV1Compat{Name: "AppendAnysUsingFilter empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendNonEmptyAnys(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendNonEmptyAnys", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys("hello", nil, "world")
		tc := caseV1Compat{Name: "AppendNonEmptyAnys", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AppendNonEmptyAnys_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AppendNonEmptyAnys_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AppendNonEmptyAnys(nil)
		tc := caseV1Compat{Name: "AppendNonEmptyAnys nil only", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddsNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddsNonEmpty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty("a", "", "b")
		tc := caseV1Compat{Name: "AddsNonEmpty", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddsNonEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddsNonEmpty_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddsNonEmpty()
		tc := caseV1Compat{Name: "AddsNonEmpty nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddsNonEmptyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddsNonEmptyPtrLock", func() {
		c := corestr.New.Collection.Cap(5)
		s := "hello"
		c.AddsNonEmptyPtrLock(&s, nil)
		tc := caseV1Compat{Name: "AddsNonEmptyPtrLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_UniqueBoolMap(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_UniqueBoolMap", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		m := c.UniqueBoolMap()
		tc := caseV1Compat{Name: "UniqueBoolMap", Expected: 2, Actual: len(m), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_UniqueBoolMapLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_UniqueBoolMapLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "a"})
		m := c.UniqueBoolMapLock()
		tc := caseV1Compat{Name: "UniqueBoolMapLock", Expected: 1, Actual: len(m), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_UniqueList(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_UniqueList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "a"})
		list := c.UniqueList()
		tc := caseV1Compat{Name: "UniqueList", Expected: 2, Actual: len(list), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_UniqueListLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_UniqueListLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "b"})
		list := c.UniqueListLock()
		tc := caseV1Compat{Name: "UniqueListLock", Expected: 2, Actual: len(list), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Filter(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Filter", func() {
		c := corestr.New.Collection.Strings([]string{"abc", "def", "ab"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 3, false
		})
		tc := caseV1Compat{Name: "Filter", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Filter_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Filter_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})
		tc := caseV1Compat{Name: "Filter empty", Expected: 0, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Filter_Break(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Filter_Break", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		result := c.Filter(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})
		tc := caseV1Compat{Name: "Filter break", Expected: 1, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_FilteredCollection(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_FilteredCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "bb", "c"})
		result := c.FilteredCollection(func(s string, i int) (string, bool, bool) {
			return s, len(s) == 1, false
		})
		tc := caseV1Compat{Name: "FilteredCollection", Expected: 2, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Has_Found(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Has_Found", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "Has found", Expected: true, Actual: c.Has("b"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Has_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Has_NotFound", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "Has not found", Expected: false, Actual: c.Has("z"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Has_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Has_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "Has empty", Expected: false, Actual: c.Has("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "HasLock", Expected: true, Actual: c.HasLock("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasPtr_Found(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasPtr_Found", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		s := "a"
		tc := caseV1Compat{Name: "HasPtr found", Expected: true, Actual: c.HasPtr(&s), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasPtr_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "HasPtr nil", Expected: false, Actual: c.HasPtr(nil), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasAll(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc := caseV1Compat{Name: "HasAll", Expected: true, Actual: c.HasAll("a", "c"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasAll_Missing(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasAll_Missing", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "HasAll missing", Expected: false, Actual: c.HasAll("a", "z"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasAll_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "HasAll empty", Expected: false, Actual: c.HasAll("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_SortedListAsc(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_SortedListAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		result := c.SortedListAsc()
		tc := caseV1Compat{Name: "SortedListAsc first", Expected: "a", Actual: result[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_SortedListAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_SortedListAsc_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		result := c.SortedListAsc()
		tc := caseV1Compat{Name: "SortedListAsc empty", Expected: 0, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_SortedAsc(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_SortedAsc", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a", "b"})
		c.SortedAsc()
		tc := caseV1Compat{Name: "SortedAsc", Expected: "a", Actual: c.First(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_SortedAsc_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_SortedAsc_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		c.SortedAsc()
		tc := caseV1Compat{Name: "SortedAsc empty", Expected: true, Actual: c.IsEmpty(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_SortedAscLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_SortedAscLock", func() {
		c := corestr.New.Collection.Strings([]string{"c", "a"})
		c.SortedAscLock()
		tc := caseV1Compat{Name: "SortedAscLock", Expected: "a", Actual: c.First(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_SortedListDsc(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_SortedListDsc", func() {
		c := corestr.New.Collection.Strings([]string{"a", "c", "b"})
		result := c.SortedListDsc()
		tc := caseV1Compat{Name: "SortedListDsc first", Expected: "c", Actual: result[0], Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasUsingSensitivity_CaseSensitive(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasUsingSensitivity_CaseSensitive", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		tc := caseV1Compat{Name: "HasUsingSensitivity sensitive", Expected: false, Actual: c.HasUsingSensitivity("hello", true), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HasUsingSensitivity_CaseInsensitive(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HasUsingSensitivity_CaseInsensitive", func() {
		c := corestr.New.Collection.Strings([]string{"Hello"})
		tc := caseV1Compat{Name: "HasUsingSensitivity insensitive", Expected: true, Actual: c.HasUsingSensitivity("hello", false), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_IsContainsPtr_Found(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_IsContainsPtr_Found", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		s := "b"
		tc := caseV1Compat{Name: "IsContainsPtr found", Expected: true, Actual: c.IsContainsPtr(&s), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_IsContainsPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_IsContainsPtr_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "IsContainsPtr nil", Expected: false, Actual: c.IsContainsPtr(nil), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_IsContainsAll(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_IsContainsAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		tc := caseV1Compat{Name: "IsContainsAll", Expected: true, Actual: c.IsContainsAll("a", "b"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_IsContainsAll_Missing(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_IsContainsAll_Missing", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "IsContainsAll missing", Expected: false, Actual: c.IsContainsAll("a", "z"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_IsContainsAllSlice_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_IsContainsAllSlice_Empty", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "IsContainsAllSlice empty", Expected: false, Actual: c.IsContainsAllSlice([]string{}), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_IsContainsAllLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_IsContainsAllLock", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "IsContainsAllLock", Expected: true, Actual: c.IsContainsAllLock("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_GetHashsetPlusHasAll(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_GetHashsetPlusHasAll", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs, hasAll := c.GetHashsetPlusHasAll([]string{"a", "b"})
		tc := caseV1Compat{Name: "GetHashsetPlusHasAll", Expected: true, Actual: hasAll, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "GetHashsetPlusHasAll hs", Expected: true, Actual: hs != nil, Args: args.Map{}}
		tc2.ShouldBeEqual(t)
	})
}

func Test_Collection_GetHashsetPlusHasAll_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_GetHashsetPlusHasAll_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		_, hasAll := c.GetHashsetPlusHasAll(nil)
		tc := caseV1Compat{Name: "GetHashsetPlusHasAll nil", Expected: false, Actual: hasAll, Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_List(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_List", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		tc := caseV1Compat{Name: "List", Expected: 2, Actual: len(c.List()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_Items(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_Items", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "Items", Expected: 1, Actual: len(c.Items()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_ListPtr(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_ListPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "ListPtr", Expected: 1, Actual: len(c.ListPtr()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_ListCopyPtrLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_ListCopyPtrLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		tc := caseV1Compat{Name: "ListCopyPtrLock", Expected: 1, Actual: len(c.ListCopyPtrLock()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_ListCopyPtrLock_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_ListCopyPtrLock_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		tc := caseV1Compat{Name: "ListCopyPtrLock empty", Expected: 0, Actual: len(c.ListCopyPtrLock()), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_NonEmptyList(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_NonEmptyList", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := c.NonEmptyList()
		tc := caseV1Compat{Name: "NonEmptyList", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_NonEmptyList_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_NonEmptyList_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		result := c.NonEmptyList()
		tc := caseV1Compat{Name: "NonEmptyList empty", Expected: 0, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_NonEmptyListPtr(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_NonEmptyListPtr", func() {
		c := corestr.New.Collection.Strings([]string{"a", ""})
		result := c.NonEmptyListPtr()
		tc := caseV1Compat{Name: "NonEmptyListPtr", Expected: 1, Actual: len(*result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HashsetAsIs(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HashsetAsIs", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		hs := c.HashsetAsIs()
		tc := caseV1Compat{Name: "HashsetAsIs", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HashsetWithDoubleLength(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HashsetWithDoubleLength", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetWithDoubleLength()
		tc := caseV1Compat{Name: "HashsetWithDoubleLength", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_HashsetLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_HashsetLock", func() {
		c := corestr.New.Collection.Strings([]string{"a"})
		hs := c.HashsetLock()
		tc := caseV1Compat{Name: "HashsetLock", Expected: true, Actual: hs.Has("a"), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_NonEmptyItems(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_NonEmptyItems", func() {
		c := corestr.New.Collection.Strings([]string{"a", "", "b"})
		result := c.NonEmptyItems()
		tc := caseV1Compat{Name: "NonEmptyItems", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_NonEmptyItemsOrNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_NonEmptyItemsOrNonWhitespace", func() {
		c := corestr.New.Collection.Strings([]string{"a", "  ", "b"})
		result := c.NonEmptyItemsOrNonWhitespace()
		tc := caseV1Compat{Name: "NonEmptyItemsOrNonWhitespace", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_New(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_New", func() {
		c := corestr.New.Collection.Cap(0)
		result := c.New("a", "b")
		tc := caseV1Compat{Name: "New", Expected: 2, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_New_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_New_Empty", func() {
		c := corestr.New.Collection.Cap(0)
		result := c.New()
		tc := caseV1Compat{Name: "New empty", Expected: 0, Actual: result.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddNonEmptyStrings(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddNonEmptyStrings", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyStrings("a", "", "b")
		tc := caseV1Compat{Name: "AddNonEmptyStrings", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddNonEmptyStrings_Empty(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddNonEmptyStrings_Empty", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyStrings()
		tc := caseV1Compat{Name: "AddNonEmptyStrings empty", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddFuncResult(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddFuncResult", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFuncResult(func() string { return "hello" })
		tc := caseV1Compat{Name: "AddFuncResult", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddFuncResult_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddFuncResult_Nil", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddFuncResult()
		tc := caseV1Compat{Name: "AddFuncResult nil", Expected: 0, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddStringsByFuncChecking(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddStringsByFuncChecking", func() {
		c := corestr.New.Collection.Cap(5)
		c.AddStringsByFuncChecking([]string{"ok", "bad", "ok2"}, func(s string) bool {
			return s != "bad"
		})
		tc := caseV1Compat{Name: "AddStringsByFuncChecking", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_MergeSlicesOfSlice(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_MergeSlicesOfSlice", func() {
		c := corestr.New.Collection.Cap(5)
		c.MergeSlicesOfSlice([]string{"a"}, []string{"b"})
		tc := caseV1Compat{Name: "MergeSlicesOfSlice", Expected: 2, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_GetAllExceptCollection(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_GetAllExceptCollection", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b", "c"})
		except := corestr.New.Collection.Strings([]string{"b"})
		result := c.GetAllExceptCollection(except)
		tc := caseV1Compat{Name: "GetAllExceptCollection", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_GetAllExceptCollection_Nil(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_GetAllExceptCollection_Nil", func() {
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		result := c.GetAllExceptCollection(nil)
		tc := caseV1Compat{Name: "GetAllExceptCollection nil", Expected: 2, Actual: len(result), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_Collection_AddPointerCollectionsLock(t *testing.T) {
	safeTest(t, "Test_Cov48_Collection_AddPointerCollectionsLock", func() {
		c := corestr.New.Collection.Cap(5)
		c2 := corestr.New.Collection.Strings([]string{"a"})
		c.AddPointerCollectionsLock(c2)
		tc := caseV1Compat{Name: "AddPointerCollectionsLock", Expected: 1, Actual: c.Length(), Args: args.Map{}}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
