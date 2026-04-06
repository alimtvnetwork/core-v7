package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

func intPtr(v int) *int       { return &v }
func strPtr(v string) *string { return &v }

func ptrStr[T any](p *T) string {
	if p == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%v", *p)
}

func ptrSliceToStrings[T any](items []*T) []string {
	result := make([]string, len(items))

	for i, p := range items {
		result[i] = ptrStr(p)
	}

	return result
}

// ==========================================================================
// Test: Ascending sort
// ==========================================================================

func Test_PointerSliceSorter_Asc_Int(t *testing.T) {
	tc := ptrSorterAscIntTestCase
	items := []*int{intPtr(3), intPtr(1), intPtr(5), intPtr(2), intPtr(4)}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

func Test_PointerSliceSorter_Asc_String(t *testing.T) {
	tc := ptrSorterAscStringTestCase
	items := []*string{strPtr("cherry"), strPtr("apple"), strPtr("banana")}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

// ==========================================================================
// Test: Descending sort
// ==========================================================================

func Test_PointerSliceSorter_Desc(t *testing.T) {
	tc := ptrSorterDescIntTestCase
	items := []*int{intPtr(3), intPtr(1), intPtr(5), intPtr(2), intPtr(4)}
	sorter := coregeneric.NewPointerSliceSorterDesc(items)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

// ==========================================================================
// Test: Nil handling
// ==========================================================================

func Test_PointerSliceSorter_NilsToEnd(t *testing.T) {
	tc := ptrSorterNilsToEndTestCase
	items := []*int{nil, intPtr(3), intPtr(1), nil, intPtr(5)}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

func Test_PointerSliceSorter_NilFirst(t *testing.T) {
	tc := ptrSorterNilFirstTestCase
	items := []*int{intPtr(3), nil, intPtr(1), nil, intPtr(5)}
	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(a, b int) bool {
		return a < b
	}, true)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

func Test_PointerSliceSorter_AllNil(t *testing.T) {
	tc := ptrSorterAllNilTestCase
	items := []*int{nil, nil, nil}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

// ==========================================================================
// Test: Custom Less function
// ==========================================================================

func Test_PointerSliceSorter_CustomLess(t *testing.T) {
	tc := ptrSorterCustomLessTestCase
	items := []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5)}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	sorter := coregeneric.NewPointerSliceSorterFunc(items, func(a, b int) bool {
		return abs(a-3) < abs(b-3)
	}, false)
	sorter.Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}

// ==========================================================================
// Test: SetAsc / SetDesc switching
// ==========================================================================

func Test_PointerSliceSorter_Switch(t *testing.T) {
	// Arrange
	tc := ptrSorterSwitchTestCase
	items := []*int{intPtr(3), intPtr(1), intPtr(5), intPtr(2), intPtr(4)}

	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()
	firstAfterAsc := ptrStr(sorter.Items()[0])
	lastAfterAsc := ptrStr(sorter.Items()[4])

	sorter.SetDesc().Sort()
	firstAfterDesc := ptrStr(sorter.Items()[0])
	lastAfterDesc := ptrStr(sorter.Items()[4])

	// Act
	actual := args.Map{
		"firstAfterAsc":  firstAfterAsc,
		"lastAfterAsc":   lastAfterAsc,
		"firstAfterDesc": firstAfterDesc,
		"lastAfterDesc":  lastAfterDesc,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IsSorted
// ==========================================================================

func Test_PointerSliceSorter_IsSorted(t *testing.T) {
	// Arrange
	tc := ptrSorterIsSortedTestCase
	items := []*int{intPtr(3), intPtr(1), intPtr(5)}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)

	beforeSort := sorter.IsSorted()
	sorter.Sort()
	afterSort := sorter.IsSorted()

	// Act
	actual := args.Map{
		"beforeSort": beforeSort,
		"afterSort":  afterSort,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Edge cases
// ==========================================================================

func Test_PointerSliceSorter_Empty(t *testing.T) {
	// Arrange
	tc := ptrSorterEmptyTestCase
	items := []*int{}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()

	// Act
	actual := args.Map{
		"length":   sorter.Len(),
		"isSorted": sorter.IsSorted(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PointerSliceSorter_Single(t *testing.T) {
	// Arrange
	tc := ptrSorterSingleTestCase
	items := []*int{intPtr(42)}
	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.Sort()

	// Act
	actual := args.Map{
		"length":   sorter.Len(),
		"isSorted": sorter.IsSorted(),
		"value":    ptrStr(sorter.Items()[0]),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_PointerSliceSorter_NilSlice(t *testing.T) {
	tc := ptrSorterNilSliceTestCase
	sorter := coregeneric.NewPointerSliceSorterAsc[int](nil)

	// Assert
	tc.ShouldBeEqualFirst(t, fmt.Sprintf("%d", sorter.Len()))
}

// ==========================================================================
// Test: SetItems / Items
// ==========================================================================

func Test_PointerSliceSorter_SetItems(t *testing.T) {
	// Arrange
	tc := ptrSorterSetItemsTestCase
	sorter := coregeneric.NewPointerSliceSorterAsc([]*int{intPtr(5), intPtr(1)})
	sorter.Sort()

	newItems := []*int{intPtr(30), intPtr(10), intPtr(20)}
	sorter.SetItems(newItems).Sort()

	sortedStrs := ptrSliceToStrings(sorter.Items())

	// Act
	actual := args.Map{
		"length": sorter.Len(),
		"item0":  sortedStrs[0],
		"item1":  sortedStrs[1],
		"item2":  sortedStrs[2],
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Chaining
// ==========================================================================

func Test_PointerSliceSorter_Chaining(t *testing.T) {
	tc := ptrSorterChainingTestCase
	items := []*int{intPtr(3), nil, intPtr(1), intPtr(5)}

	sorter := coregeneric.NewPointerSliceSorterAsc(items)
	sorter.SetDesc().SetNilFirst(true).Sort()

	// Assert
	tc.ShouldBeEqualFirst(t, ptrSliceToStrings(sorter.Items())...)
}
