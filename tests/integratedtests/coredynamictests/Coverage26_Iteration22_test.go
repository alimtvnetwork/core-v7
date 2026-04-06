package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CollectionTypes — factory shortcuts
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_NewStringCollection(t *testing.T) {
	// Arrange
	col := coredynamic.NewStringCollection(5)
	col.Add("a").Add("b")

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewStringCollection returns correct value -- with args", actual)
}

func Test_I22_EmptyStringCollection(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()

	// Act
	actual := args.Map{"empty": col.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyStringCollection returns empty -- with args", actual)
}

func Test_I22_NewIntCollection(t *testing.T) {
	// Arrange
	col := coredynamic.NewIntCollection(3)
	col.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "NewIntCollection returns correct value -- with args", actual)
}

func Test_I22_EmptyIntCollection(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()

	// Act
	actual := args.Map{"empty": col.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "EmptyIntCollection returns empty -- with args", actual)
}

func Test_I22_NewInt64Collection(t *testing.T) {
	// Arrange
	col := coredynamic.NewInt64Collection(2)
	col.Add(int64(99))

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewInt64Collection returns correct value -- with args", actual)
}

func Test_I22_NewByteCollection(t *testing.T) {
	// Arrange
	col := coredynamic.NewByteCollection(2)
	col.Add(byte(0x41))

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewByteCollection returns correct value -- with args", actual)
}

func Test_I22_NewBoolCollection(t *testing.T) {
	// Arrange
	col := coredynamic.NewBoolCollection(2)
	col.Add(true).Add(false)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewBoolCollection returns correct value -- with args", actual)
}

func Test_I22_NewFloat64Collection(t *testing.T) {
	// Arrange
	col := coredynamic.NewFloat64Collection(2)
	col.Add(3.14)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewFloat64Collection returns correct value -- with args", actual)
}

func Test_I22_NewAnyMapCollection(t *testing.T) {
	// Arrange
	col := coredynamic.NewAnyMapCollection(2)
	col.Add(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewAnyMapCollection returns correct value -- with args", actual)
}

func Test_I22_NewStringMapCollection(t *testing.T) {
	// Arrange
	col := coredynamic.NewStringMapCollection(2)
	col.Add(map[string]string{"k": "v"})

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewStringMapCollection returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionDistinct
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_Distinct_Duplicates(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a").Add("c").Add("b")
	result := coredynamic.Distinct(col)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Distinct returns correct value -- duplicates", actual)
}

func Test_I22_Distinct_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	result := coredynamic.Distinct(col)

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Distinct returns empty -- empty", actual)
}

func Test_I22_Unique_Alias(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(1)
	result := coredynamic.Unique(col)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Unique returns correct value -- alias", actual)
}

func Test_I22_DistinctLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("x").Add("x").Add("y")
	result := coredynamic.DistinctLock(col)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DistinctLock returns correct value -- with args", actual)
}

func Test_I22_DistinctCount(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a")

	// Act
	actual := args.Map{"count": coredynamic.DistinctCount(col)}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DistinctCount returns correct value -- with args", actual)
}

func Test_I22_DistinctCount_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()

	// Act
	actual := args.Map{"count": coredynamic.DistinctCount(col)}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "DistinctCount returns empty -- empty", actual)
}

func Test_I22_IsDistinct_True(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"distinct": coredynamic.IsDistinct(col)}

	// Assert
	expected := args.Map{"distinct": true}
	expected.ShouldBeEqual(t, 0, "IsDistinct returns non-empty -- true", actual)
}

func Test_I22_IsDistinct_False(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(1)

	// Act
	actual := args.Map{"distinct": coredynamic.IsDistinct(col)}

	// Assert
	expected := args.Map{"distinct": false}
	expected.ShouldBeEqual(t, 0, "IsDistinct returns non-empty -- false", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionMap — Map, FlatMap, Reduce
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_Map_Transform(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	result := coredynamic.Map(col, func(i int) string {
		return "x"
	})

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- transform", actual)
}

func Test_I22_Map_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	result := coredynamic.Map(col, func(i int) string { return "" })

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- empty", actual)
}

func Test_I22_Map_Nil(t *testing.T) {
	// Arrange
	result := coredynamic.Map[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- nil", actual)
}

func Test_I22_FlatMap(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyCollection[[]string]()
	col.Add([]string{"a", "b"}).Add([]string{"c"})
	result := coredynamic.FlatMap(col, func(s []string) []string { return s })

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "FlatMap returns correct value -- with args", actual)
}

func Test_I22_FlatMap_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyCollection[[]string]()
	result := coredynamic.FlatMap(col, func(s []string) []string { return s })

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FlatMap returns empty -- empty", actual)
}

func Test_I22_FlatMap_Nil(t *testing.T) {
	// Arrange
	result := coredynamic.FlatMap[[]string, string](nil, func(s []string) []string { return s })

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "FlatMap returns nil -- nil", actual)
}

func Test_I22_Reduce_Sum(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)
	sum := coredynamic.Reduce(col, 0, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "Reduce returns correct value -- sum", actual)
}

func Test_I22_Reduce_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	sum := coredynamic.Reduce(col, 10, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 10}
	expected.ShouldBeEqual(t, 0, "Reduce returns empty -- empty returns initial", actual)
}

func Test_I22_Reduce_Nil(t *testing.T) {
	// Arrange
	sum := coredynamic.Reduce[int, int](nil, 42, func(acc int, item int) int { return acc + item })

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 42}
	expected.ShouldBeEqual(t, 0, "Reduce returns nil -- nil returns initial", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionSearch — Contains, IndexOf, Has, HasAll, LastIndexOf, Count, Lock variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_Contains_Found(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b")

	// Act
	actual := args.Map{"found": coredynamic.Contains(col, "b")}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "Contains returns correct value -- found", actual)
}

func Test_I22_Contains_NotFound(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a")

	// Act
	actual := args.Map{"found": coredynamic.Contains(col, "z")}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "Contains returns correct value -- not found", actual)
}

func Test_I22_IndexOf_Found(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("x").Add("y").Add("z")

	// Act
	actual := args.Map{"idx": coredynamic.IndexOf(col, "y")}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "IndexOf returns correct value -- found", actual)
}

func Test_I22_IndexOf_NotFound(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("x")

	// Act
	actual := args.Map{"idx": coredynamic.IndexOf(col, "z")}

	// Assert
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "IndexOf returns correct value -- not found", actual)
}

func Test_I22_Has_Alias(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2)

	// Act
	actual := args.Map{"has": coredynamic.Has(col, 2)}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Has returns correct value -- alias", actual)
}

func Test_I22_HasAll_True(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"all": coredynamic.HasAll(col, 1, 3)}

	// Assert
	expected := args.Map{"all": true}
	expected.ShouldBeEqual(t, 0, "HasAll returns non-empty -- true", actual)
}

func Test_I22_HasAll_False(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2)

	// Act
	actual := args.Map{"all": coredynamic.HasAll(col, 1, 9)}

	// Assert
	expected := args.Map{"all": false}
	expected.ShouldBeEqual(t, 0, "HasAll returns non-empty -- false", actual)
}

func Test_I22_HasAll_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()

	// Act
	actual := args.Map{"all": coredynamic.HasAll(col, 1)}

	// Assert
	expected := args.Map{"all": false}
	expected.ShouldBeEqual(t, 0, "HasAll returns empty -- empty", actual)
}

func Test_I22_LastIndexOf_Found(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a")

	// Act
	actual := args.Map{"idx": coredynamic.LastIndexOf(col, "a")}

	// Assert
	expected := args.Map{"idx": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexOf returns correct value -- found", actual)
}

func Test_I22_LastIndexOf_NotFound(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a")

	// Act
	actual := args.Map{"idx": coredynamic.LastIndexOf(col, "z")}

	// Assert
	expected := args.Map{"idx": -1}
	expected.ShouldBeEqual(t, 0, "LastIndexOf returns correct value -- not found", actual)
}

func Test_I22_Count_Occurrences(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a").Add("a")

	// Act
	actual := args.Map{"count": coredynamic.Count(col, "a")}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "Count returns correct value -- occurrences", actual)
}

func Test_I22_ContainsLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("x")

	// Act
	actual := args.Map{"found": coredynamic.ContainsLock(col, "x")}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "ContainsLock returns correct value -- with args", actual)
}

func Test_I22_IndexOfLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b")

	// Act
	actual := args.Map{"idx": coredynamic.IndexOfLock(col, "b")}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "IndexOfLock returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionSort
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_SortFunc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	col.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"first": col.First(),
		"last": col.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortFunc returns correct value -- with args", actual)
}

func Test_I22_SortFunc_Single(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1)
	col.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFunc returns correct value -- single", actual)
}

func Test_I22_SortFuncLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	col.SortFuncLock(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFuncLock returns correct value -- with args", actual)
}

func Test_I22_SortedFunc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	sorted := col.SortedFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"origFirst": col.First(),
		"sortedFirst": sorted.First(),
	}

	// Assert
	expected := args.Map{
		"origFirst": 3,
		"sortedFirst": 1,
	}
	expected.ShouldBeEqual(t, 0, "SortedFunc returns correct value -- does not mutate", actual)
}

func Test_I22_SortAsc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	coredynamic.SortAsc(col)

	// Act
	actual := args.Map{
		"first": col.First(),
		"last": col.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortAsc returns correct value -- with args", actual)
}

func Test_I22_SortDesc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(3).Add(2)
	coredynamic.SortDesc(col)

	// Act
	actual := args.Map{
		"first": col.First(),
		"last": col.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 3,
		"last": 1,
	}
	expected.ShouldBeEqual(t, 0, "SortDesc returns correct value -- with args", actual)
}

func Test_I22_SortAscLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1)
	coredynamic.SortAscLock(col)

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortAscLock returns correct value -- with args", actual)
}

func Test_I22_SortDescLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(3)
	coredynamic.SortDescLock(col)

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortDescLock returns correct value -- with args", actual)
}

func Test_I22_SortedAsc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1).Add(2)
	sorted := coredynamic.SortedAsc(col)

	// Act
	actual := args.Map{
		"origFirst": col.First(),
		"sortedFirst": sorted.First(),
	}

	// Assert
	expected := args.Map{
		"origFirst": 3,
		"sortedFirst": 1,
	}
	expected.ShouldBeEqual(t, 0, "SortedAsc returns correct value -- with args", actual)
}

func Test_I22_SortedDesc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(3).Add(2)
	sorted := coredynamic.SortedDesc(col)

	// Act
	actual := args.Map{
		"origFirst": col.First(),
		"sortedFirst": sorted.First(),
	}

	// Assert
	expected := args.Map{
		"origFirst": 1,
		"sortedFirst": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortedDesc returns correct value -- with args", actual)
}

func Test_I22_IsSorted_True(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"sorted": col.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"sorted": true}
	expected.ShouldBeEqual(t, 0, "IsSorted returns non-empty -- true", actual)
}

func Test_I22_IsSorted_False(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(1)

	// Act
	actual := args.Map{"sorted": col.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"sorted": false}
	expected.ShouldBeEqual(t, 0, "IsSorted returns non-empty -- false", actual)
}

func Test_I22_IsSorted_Single(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1)

	// Act
	actual := args.Map{"sorted": col.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"sorted": true}
	expected.ShouldBeEqual(t, 0, "IsSorted returns correct value -- single", actual)
}

func Test_I22_IsSortedAsc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"asc": coredynamic.IsSortedAsc(col)}

	// Assert
	expected := args.Map{"asc": true}
	expected.ShouldBeEqual(t, 0, "IsSortedAsc returns correct value -- with args", actual)
}

func Test_I22_IsSortedDesc(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(3).Add(2).Add(1)

	// Act
	actual := args.Map{"desc": coredynamic.IsSortedDesc(col)}

	// Assert
	expected := args.Map{"desc": true}
	expected.ShouldBeEqual(t, 0, "IsSortedDesc returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CollectionGroupBy
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_GroupBy(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("ab").Add("ac").Add("bd")
	groups := coredynamic.GroupBy(col, func(s string) string { return string(s[0]) })

	// Act
	actual := args.Map{
		"groups": len(groups),
		"aLen": groups["a"].Length(),
	}

	// Assert
	expected := args.Map{
		"groups": 2,
		"aLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "GroupBy returns correct value -- with args", actual)
}

func Test_I22_GroupBy_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	groups := coredynamic.GroupBy(col, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(groups)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupBy returns empty -- empty", actual)
}

func Test_I22_GroupBy_Nil(t *testing.T) {
	// Arrange
	groups := coredynamic.GroupBy[string, string](nil, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(groups)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupBy returns nil -- nil", actual)
}

func Test_I22_GroupByLock(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyIntCollection()
	col.Add(1).Add(2).Add(3).Add(4)
	groups := coredynamic.GroupByLock(col, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})

	// Act
	actual := args.Map{
		"even": groups["even"].Length(),
		"odd": groups["odd"].Length(),
	}

	// Assert
	expected := args.Map{
		"even": 2,
		"odd": 2,
	}
	expected.ShouldBeEqual(t, 0, "GroupByLock returns correct value -- with args", actual)
}

func Test_I22_GroupByCount(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	col.Add("a").Add("b").Add("a")
	counts := coredynamic.GroupByCount(col, func(s string) string { return s })

	// Act
	actual := args.Map{
		"a": counts["a"],
		"b": counts["b"],
	}

	// Assert
	expected := args.Map{
		"a": 2,
		"b": 1,
	}
	expected.ShouldBeEqual(t, 0, "GroupByCount returns correct value -- with args", actual)
}

func Test_I22_GroupByCount_Empty(t *testing.T) {
	// Arrange
	col := coredynamic.EmptyStringCollection()
	counts := coredynamic.GroupByCount(col, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(counts)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupByCount returns empty -- empty", actual)
}

func Test_I22_GroupByCount_Nil(t *testing.T) {
	// Arrange
	counts := coredynamic.GroupByCount[string, string](nil, func(s string) string { return s })

	// Act
	actual := args.Map{"len": len(counts)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupByCount returns nil -- nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReflectSetFromTo — deeper paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_I22_ReflectSetFromTo_BothNil(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo(nil, nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- both nil", actual)
}

func Test_I22_ReflectSetFromTo_SameType(t *testing.T) {
	// Arrange
	from := "hello"
	var to string
	err := coredynamic.ReflectSetFromTo(from, &to)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": to,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- same type", actual)
}

func Test_I22_ReflectSetFromTo_SamePointerType(t *testing.T) {
	// Arrange
	from := new(string)
	*from = "hello"
	to := new(string)
	err := coredynamic.ReflectSetFromTo(from, to)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": *to,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- same ptr type", actual)
}

func Test_I22_ReflectSetFromTo_ToNonPointer(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("hello", "notpointer")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns non-empty -- to non-pointer", actual)
}

func Test_I22_ReflectSetFromTo_ToNil(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectSetFromTo("hello", nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- to nil", actual)
}

func Test_I22_ReflectSetFromTo_FromNilPtr(t *testing.T) {
	// Arrange
	var from *string
	var to string
	err := coredynamic.ReflectSetFromTo(from, &to)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns nil -- from nil ptr", actual)
}

func Test_I22_ReflectSetFromTo_BytesToStruct(t *testing.T) {
	// Arrange
	type Simple struct {
		Name string `json:"name"`
	}
	from := []byte(`{"name":"test"}`)
	var to Simple
	err := coredynamic.ReflectSetFromTo(from, &to)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name": to.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- bytes to struct", actual)
}

func Test_I22_ReflectSetFromTo_StructToBytes(t *testing.T) {
	// Arrange
	type Simple struct {
		Name string `json:"name"`
	}
	from := Simple{Name: "test"}
	var to []byte
	err := coredynamic.ReflectSetFromTo(from, &to)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(to) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- struct to bytes", actual)
}

func Test_I22_ReflectSetFromTo_TypeMismatch(t *testing.T) {
	// Arrange
	from := "hello"
	var to int
	err := coredynamic.ReflectSetFromTo(from, to)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- type mismatch", actual)
}

func Test_I22_ReflectSetFromTo_IntValue(t *testing.T) {
	// Arrange
	from := 42
	var to int
	err := coredynamic.ReflectSetFromTo(from, &to)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": to,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- int value", actual)
}

func Test_I22_ReflectSetFromTo_ReflectType(t *testing.T) {
	// Arrange
	_ = reflect.TypeOf("")
	from := true
	var to bool
	err := coredynamic.ReflectSetFromTo(from, &to)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": to,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct value -- bool", actual)
}
