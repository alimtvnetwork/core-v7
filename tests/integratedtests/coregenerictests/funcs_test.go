package coregenerictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================================================
// Test: MapCollection
// ==========================================================================

func Test_MapCollection_IntToString(t *testing.T) {
	// Arrange
	tc := mapCollectionIntToStringTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := coregeneric.MapCollection(src, func(i int) string { return fmt.Sprintf("v%d", i) })

	// Act
	actual := args.Map{
		"length": result.Length(),
		"first":  result.First(),
		"last":   result.Last(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapCollection_NilSource(t *testing.T) {
	// Arrange
	tc := mapCollectionNilSourceTestCase
	result := coregeneric.MapCollection[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"isEmpty": result.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapCollection_EmptySource(t *testing.T) {
	// Arrange
	tc := mapCollectionEmptySourceTestCase
	src := coregeneric.EmptyCollection[int]()
	result := coregeneric.MapCollection(src, func(i int) string { return "" })

	// Act
	actual := args.Map{"isEmpty": result.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: FlatMapCollection
// ==========================================================================

func Test_FlatMapCollection_Flattens(t *testing.T) {
	// Arrange
	tc := flatMapCollectionFlattensTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)
	result := coregeneric.FlatMapCollection(src, func(i int) []int { return []int{i, i * 10} })

	// Act
	actual := args.Map{"length": result.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_FlatMapCollection_Nil(t *testing.T) {
	// Arrange
	tc := flatMapCollectionNilTestCase
	result := coregeneric.FlatMapCollection[int, int](nil, func(i int) []int { return nil })

	// Act
	actual := args.Map{"isEmpty": result.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ReduceCollection
// ==========================================================================

func Test_ReduceCollection_Sum(t *testing.T) {
	// Arrange
	tc := reduceCollectionSumTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4)
	sum := coregeneric.ReduceCollection(src, 0, func(a, b int) int { return a + b })

	// Act
	actual := args.Map{"result": sum}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ReduceCollection_Nil(t *testing.T) {
	// Arrange
	tc := reduceCollectionNilTestCase
	result := coregeneric.ReduceCollection[int, int](nil, 99, func(a, b int) int { return a + b })

	// Act
	actual := args.Map{"result": result}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ReduceCollection_Concat(t *testing.T) {
	// Arrange
	tc := reduceCollectionConcatTestCase
	src := coregeneric.New.Collection.String.Items("a", "b", "c")
	result := coregeneric.ReduceCollection(src, "", func(a, b string) string { return a + b })

	// Act
	actual := args.Map{"result": result}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: GroupByCollection
// ==========================================================================

func Test_GroupByCollection_Groups(t *testing.T) {
	// Arrange
	tc := groupByCollectionGroupsTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3, 4, 5, 6)
	groups := coregeneric.GroupByCollection(src, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})

	// Act
	actual := args.Map{
		"groupCount": len(groups),
		"evenCount":  groups["even"].Length(),
		"oddCount":   groups["odd"].Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_GroupByCollection_Nil(t *testing.T) {
	// Arrange
	tc := groupByCollectionNilTestCase
	groups := coregeneric.GroupByCollection[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"groupCount": len(groups)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ContainsFunc
// ==========================================================================

func Test_ContainsFunc_Found(t *testing.T) {
	// Arrange
	tc := containsFuncFoundTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)

	// Act
	actual := args.Map{"result": coregeneric.ContainsFunc(src, func(i int) bool { return i == 2 })}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsFunc_NotFound(t *testing.T) {
	// Arrange
	tc := containsFuncNotFoundTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)

	// Act
	actual := args.Map{"result": coregeneric.ContainsFunc(src, func(i int) bool { return i == 99 })}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsFunc_Nil(t *testing.T) {
	// Arrange
	tc := containsFuncNilTestCase

	// Act
	actual := args.Map{"result": coregeneric.ContainsFunc[int](nil, func(i int) bool { return true })}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: ContainsItem
// ==========================================================================

func Test_ContainsItem_Found(t *testing.T) {
	// Arrange
	tc := containsItemFoundTestCase
	src := coregeneric.New.Collection.String.Items("a", "b", "c")

	// Act
	actual := args.Map{"result": coregeneric.ContainsItem(src, "b")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsItem_NotFound(t *testing.T) {
	// Arrange
	tc := containsItemNotFoundTestCase
	src := coregeneric.New.Collection.String.Items("a", "b")

	// Act
	actual := args.Map{"result": coregeneric.ContainsItem(src, "z")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_ContainsItem_Nil(t *testing.T) {
	// Arrange
	tc := containsItemNilTestCase

	// Act
	actual := args.Map{"result": coregeneric.ContainsItem[string](nil, "x")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IndexOfFunc
// ==========================================================================

func Test_IndexOfFunc_Found(t *testing.T) {
	// Arrange
	tc := indexOfFuncFoundTestCase
	src := coregeneric.New.Collection.Int.Items(10, 20, 30)

	// Act
	actual := args.Map{"index": coregeneric.IndexOfFunc(src, func(i int) bool { return i == 20 })}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IndexOfFunc_NotFound(t *testing.T) {
	// Arrange
	tc := indexOfFuncNotFoundTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 3)

	// Act
	actual := args.Map{"index": coregeneric.IndexOfFunc(src, func(i int) bool { return i == 99 })}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IndexOfFunc_Nil(t *testing.T) {
	// Arrange
	tc := indexOfFuncNilTestCase

	// Act
	actual := args.Map{"index": coregeneric.IndexOfFunc[int](nil, func(i int) bool { return true })}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: IndexOfItem
// ==========================================================================

func Test_IndexOfItem_Found(t *testing.T) {
	// Arrange
	tc := indexOfItemFoundTestCase
	src := coregeneric.New.Collection.String.Items("x", "y", "z")

	// Act
	actual := args.Map{"index": coregeneric.IndexOfItem(src, "z")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_IndexOfItem_NotFound(t *testing.T) {
	// Arrange
	tc := indexOfItemNotFoundTestCase
	src := coregeneric.New.Collection.String.Items("a")

	// Act
	actual := args.Map{"index": coregeneric.IndexOfItem(src, "q")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: Distinct
// ==========================================================================

func Test_Distinct_RemovesDuplicates(t *testing.T) {
	// Arrange
	tc := distinctRemovesDuplicatesTestCase
	src := coregeneric.New.Collection.Int.Items(1, 2, 2, 3, 1, 3)

	// Act
	actual := args.Map{"length": coregeneric.Distinct(src).Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Distinct_Nil(t *testing.T) {
	// Arrange
	tc := distinctNilTestCase

	// Act
	actual := args.Map{"isEmpty": coregeneric.Distinct[int](nil).IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Distinct_NoDuplicates(t *testing.T) {
	// Arrange
	tc := distinctNoDuplicatesTestCase
	src := coregeneric.New.Collection.String.Items("a", "b", "c")

	// Act
	actual := args.Map{"length": coregeneric.Distinct(src).Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================================================
// Test: MapSimpleSlice
// ==========================================================================

func Test_MapSimpleSlice_Transforms(t *testing.T) {
	// Arrange
	tc := mapSimpleSliceTransformsTestCase
	src := coregeneric.SimpleSliceFrom([]int{1, 2, 3})
	result := coregeneric.MapSimpleSlice(src, func(i int) string { return fmt.Sprintf("%d", i) })

	// Act
	actual := args.Map{"length": result.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapSimpleSlice_Nil_Fromfuncs(t *testing.T) {
	// Arrange
	tc := mapSimpleSliceNilTestCase
	result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"isEmpty": result.IsEmpty()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
