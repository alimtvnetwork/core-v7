package coredynamictests

import (
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// CollectionMethods — AddIf / AddManyIf
// =============================================================================

func Test_ColM_AddIf_True(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddIf(true, "a")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddIf true", actual)
}

func Test_ColM_AddIf_False(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[string]()
	c.AddIf(false, "a")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddIf false", actual)
}

func Test_ColM_AddManyIf_True(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddManyIf(true, 1, 2)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddManyIf true", actual)
}

func Test_ColM_AddManyIf_False(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddManyIf(false, 1, 2)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddManyIf false", actual)
}

func Test_ColM_AddManyIf_EmptyItems(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddManyIf(true)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddManyIf empty items", actual)
}

// =============================================================================
// AddCollection / AddCollections / ConcatNew
// =============================================================================

func Test_ColM_AddCollection_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	c.AddCollection(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddCollection nil", actual)
}

func Test_ColM_AddCollection_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	other := coredynamic.CollectionFrom([]int{2, 3})
	c.AddCollection(other)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AddCollection valid", actual)
}

func Test_ColM_AddCollections(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	a := coredynamic.CollectionFrom([]int{1})
	b := coredynamic.CollectionFrom([]int{2})
	c.AddCollections(a, nil, b)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddCollections", actual)
}

func Test_ColM_ConcatNew(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	n := c.ConcatNew(3, 4)

	// Act
	actual := args.Map{
		"origLen": c.Length(),
		"newLen": n.Length(),
	}

	// Assert
	expected := args.Map{
		"origLen": 2,
		"newLen": 4,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew", actual)
}

// =============================================================================
// Clone / Capacity / AddCapacity / Resize
// =============================================================================

func Test_ColM_Clone_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[int]
	cl := c.Clone()

	// Act
	actual := args.Map{"len": cl.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone nil", actual)
}

func Test_ColM_Clone_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	cl := c.Clone()
	cl.Add(3)

	// Act
	actual := args.Map{
		"origLen": c.Length(),
		"cloneLen": cl.Length(),
	}

	// Assert
	expected := args.Map{
		"origLen": 2,
		"cloneLen": 3,
	}
	expected.ShouldBeEqual(t, 0, "Clone valid", actual)
}

func Test_ColM_Capacity_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[int]

	// Act
	actual := args.Map{"r": c.Capacity()}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "Capacity nil", actual)
}

func Test_ColM_AddCapacity_Zero(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](5)
	cap1 := c.Capacity()
	c.AddCapacity(0)

	// Act
	actual := args.Map{"same": c.Capacity() == cap1}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "AddCapacity zero", actual)
}

func Test_ColM_AddCapacity_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](2)
	c.AddCapacity(10)

	// Act
	actual := args.Map{"grew": c.Capacity() >= 12}

	// Assert
	expected := args.Map{"grew": true}
	expected.ShouldBeEqual(t, 0, "AddCapacity valid", actual)
}

func Test_ColM_Resize_NoGrow(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](10)
	cap1 := c.Capacity()
	c.Resize(5)

	// Act
	actual := args.Map{"same": c.Capacity() == cap1}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "Resize no grow", actual)
}

func Test_ColM_Resize_Grow(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](2)
	c.Resize(20)

	// Act
	actual := args.Map{"grew": c.Capacity() >= 20}

	// Assert
	expected := args.Map{"grew": true}
	expected.ShouldBeEqual(t, 0, "Resize grow", actual)
}

// =============================================================================
// Reverse / InsertAt
// =============================================================================

func Test_ColM_Reverse_Single(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	c.Reverse()

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "Reverse single", actual)
}

func Test_ColM_Reverse_Multi(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	c.Reverse()

	// Act
	actual := args.Map{
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 3,
		"last": 1,
	}
	expected.ShouldBeEqual(t, 0, "Reverse multi", actual)
}

func Test_ColM_InsertAt_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	c.InsertAt(1)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "InsertAt empty items", actual)
}

func Test_ColM_InsertAt_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 3})
	c.InsertAt(1, 2)

	// Act
	actual := args.Map{
		"len": c.Length(),
		"at1": c.At(1),
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"at1": 2,
	}
	expected.ShouldBeEqual(t, 0, "InsertAt valid", actual)
}

// =============================================================================
// IndexOfFunc / ContainsFunc / SafeAt / SprintItems
// =============================================================================

func Test_ColM_IndexOfFunc_Found(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"r": c.IndexOfFunc(func(s string) bool { return s == "b" })}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "IndexOfFunc found", actual)
}

func Test_ColM_IndexOfFunc_NotFound(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})

	// Act
	actual := args.Map{"r": c.IndexOfFunc(func(s string) bool { return s == "z" })}

	// Assert
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "IndexOfFunc not found", actual)
}

func Test_ColM_ContainsFunc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"found": c.ContainsFunc(func(i int) bool { return i == 2 })}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "ContainsFunc", actual)
}

func Test_ColM_SafeAt_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{"r": c.SafeAt(1)}

	// Assert
	expected := args.Map{"r": "b"}
	expected.ShouldBeEqual(t, 0, "SafeAt valid", actual)
}

func Test_ColM_SafeAt_Invalid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a"})

	// Act
	actual := args.Map{"r": c.SafeAt(5)}

	// Assert
	expected := args.Map{"r": ""}
	expected.ShouldBeEqual(t, 0, "SafeAt invalid", actual)
}

func Test_ColM_SprintItems(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{42})
	s := c.SprintItems("[%d]")

	// Act
	actual := args.Map{"first": s[0]}

	// Assert
	expected := args.Map{"first": "[42]"}
	expected.ShouldBeEqual(t, 0, "SprintItems", actual)
}

// =============================================================================
// CollectionSearch — Contains / IndexOf / Has / HasAll / LastIndexOf / Count
// =============================================================================

func Test_ColS_Contains(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b"})

	// Act
	actual := args.Map{
		"found": coredynamic.Contains(c, "b"),
		"notFound": coredynamic.Contains(c, "z"),
	}

	// Assert
	expected := args.Map{
		"found": true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "Contains", actual)
}

func Test_ColS_IndexOf(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{10, 20, 30})

	// Act
	actual := args.Map{"r": coredynamic.IndexOf(c, 20)}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "IndexOf", actual)
}

func Test_ColS_Has(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"x"})

	// Act
	actual := args.Map{"r": coredynamic.Has(c, "x")}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "Has", actual)
}

func Test_ColS_HasAll_True(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": coredynamic.HasAll(c, 1, 3)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "HasAll true", actual)
}

func Test_ColS_HasAll_False(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"r": coredynamic.HasAll(c, 1, 9)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "HasAll false", actual)
}

func Test_ColS_HasAll_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()

	// Act
	actual := args.Map{"r": coredynamic.HasAll(c, 1)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "HasAll empty", actual)
}

func Test_ColS_LastIndexOf(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 1, 3})

	// Act
	actual := args.Map{"r": coredynamic.LastIndexOf(c, 1)}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "LastIndexOf", actual)
}

func Test_ColS_LastIndexOf_NotFound(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"r": coredynamic.LastIndexOf(c, 9)}

	// Assert
	expected := args.Map{"r": -1}
	expected.ShouldBeEqual(t, 0, "LastIndexOf not found", actual)
}

func Test_ColS_Count(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 1, 1})

	// Act
	actual := args.Map{"r": coredynamic.Count(c, 1)}

	// Assert
	expected := args.Map{"r": 3}
	expected.ShouldBeEqual(t, 0, "Count", actual)
}

func Test_ColS_ContainsLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"r": coredynamic.ContainsLock(c, 1)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "ContainsLock", actual)
}

func Test_ColS_IndexOfLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{5, 10})

	// Act
	actual := args.Map{"r": coredynamic.IndexOfLock(c, 10)}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "IndexOfLock", actual)
}

// =============================================================================
// CollectionSort
// =============================================================================

func Test_ColSort_SortFunc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{
		"first": c.First(),
		"last": c.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 1,
		"last": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortFunc", actual)
}

func Test_ColSort_SortFunc_Single(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFunc single", actual)
}

func Test_ColSort_SortAsc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	coredynamic.SortAsc(c)

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortAsc", actual)
}

func Test_ColSort_SortDesc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 3, 2})
	coredynamic.SortDesc(c)

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 3}
	expected.ShouldBeEqual(t, 0, "SortDesc", actual)
}

func Test_ColSort_SortedAsc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 1, 2})
	s := coredynamic.SortedAsc(c)

	// Act
	actual := args.Map{
		"origFirst": c.First(),
		"sortedFirst": s.First(),
	}

	// Assert
	expected := args.Map{
		"origFirst": 3,
		"sortedFirst": 1,
	}
	expected.ShouldBeEqual(t, 0, "SortedAsc", actual)
}

func Test_ColSort_SortedDesc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 3, 2})
	s := coredynamic.SortedDesc(c)

	// Act
	actual := args.Map{"sortedFirst": s.First()}

	// Assert
	expected := args.Map{"sortedFirst": 3}
	expected.ShouldBeEqual(t, 0, "SortedDesc", actual)
}

func Test_ColSort_SortFuncLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{2, 1})
	c.SortFuncLock(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortFuncLock", actual)
}

func Test_ColSort_SortAscLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{2, 1})
	coredynamic.SortAscLock(c)

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 1}
	expected.ShouldBeEqual(t, 0, "SortAscLock", actual)
}

func Test_ColSort_SortDescLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	coredynamic.SortDescLock(c)

	// Act
	actual := args.Map{"first": c.First()}

	// Assert
	expected := args.Map{"first": 2}
	expected.ShouldBeEqual(t, 0, "SortDescLock", actual)
}

func Test_ColSort_IsSorted_True(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": c.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsSorted true", actual)
}

func Test_ColSort_IsSorted_Single(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"r": c.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsSorted single", actual)
}

func Test_ColSort_IsSortedAsc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": coredynamic.IsSortedAsc(c)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsSortedAsc", actual)
}

func Test_ColSort_IsSortedDesc(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{3, 2, 1})

	// Act
	actual := args.Map{"r": coredynamic.IsSortedDesc(c)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsSortedDesc", actual)
}

// =============================================================================
// CollectionMap — Map / FlatMap / Reduce
// =============================================================================

func Test_ColMap_Map_Nil(t *testing.T) {
	// Arrange
	r := coredynamic.Map[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map nil", actual)
}

func Test_ColMap_Map_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	r := coredynamic.Map(c, func(i int) int { return i * 2 })

	// Act
	actual := args.Map{
		"first": r.First(),
		"last": r.Last(),
	}

	// Assert
	expected := args.Map{
		"first": 2,
		"last": 6,
	}
	expected.ShouldBeEqual(t, 0, "Map valid", actual)
}

func Test_ColMap_FlatMap_Nil(t *testing.T) {
	// Arrange
	r := coredynamic.FlatMap[int, string](nil, func(i int) []string { return nil })

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "FlatMap nil", actual)
}

func Test_ColMap_FlatMap_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"ab", "cd"})
	r := coredynamic.FlatMap(c, func(s string) []string { return []string{s, s} })

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "FlatMap valid", actual)
}

func Test_ColMap_Reduce_Nil(t *testing.T) {
	// Arrange
	r := coredynamic.Reduce[int, int](nil, 0, func(acc, i int) int { return acc + i })

	// Act
	actual := args.Map{"r": r}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "Reduce nil", actual)
}

func Test_ColMap_Reduce_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	r := coredynamic.Reduce(c, 0, func(acc, i int) int { return acc + i })

	// Act
	actual := args.Map{"r": r}

	// Assert
	expected := args.Map{"r": 6}
	expected.ShouldBeEqual(t, 0, "Reduce valid", actual)
}

// =============================================================================
// CollectionDistinct
// =============================================================================

func Test_ColDist_Distinct_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()

	// Act
	actual := args.Map{"len": coredynamic.Distinct(c).Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Distinct empty", actual)
}

func Test_ColDist_Distinct_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 1, 3, 2})

	// Act
	actual := args.Map{"len": coredynamic.Distinct(c).Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Distinct valid", actual)
}

func Test_ColDist_Unique(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 1})

	// Act
	actual := args.Map{"len": coredynamic.Unique(c).Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Unique", actual)
}

func Test_ColDist_DistinctLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 1})

	// Act
	actual := args.Map{"len": coredynamic.DistinctLock(c).Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DistinctLock", actual)
}

func Test_ColDist_DistinctCount_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()

	// Act
	actual := args.Map{"r": coredynamic.DistinctCount(c)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "DistinctCount empty", actual)
}

func Test_ColDist_DistinctCount_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 1})

	// Act
	actual := args.Map{"r": coredynamic.DistinctCount(c)}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "DistinctCount valid", actual)
}

func Test_ColDist_IsDistinct_True(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})

	// Act
	actual := args.Map{"r": coredynamic.IsDistinct(c)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsDistinct true", actual)
}

func Test_ColDist_IsDistinct_False(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 1})

	// Act
	actual := args.Map{"r": coredynamic.IsDistinct(c)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "IsDistinct false", actual)
}

// =============================================================================
// CollectionGroupBy
// =============================================================================

func Test_ColGB_GroupBy_Nil(t *testing.T) {
	// Arrange
	r := coredynamic.GroupBy[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"len": len(r)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupBy nil", actual)
}

func Test_ColGB_GroupBy_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	r := coredynamic.GroupBy(c, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})

	// Act
	actual := args.Map{
		"groups": len(r),
		"evenLen": r["even"].Length(),
	}

	// Assert
	expected := args.Map{
		"groups": 2,
		"evenLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "GroupBy valid", actual)
}

func Test_ColGB_GroupByLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	r := coredynamic.GroupByLock(c, func(i int) int { return i })

	// Act
	actual := args.Map{"groups": len(r)}

	// Assert
	expected := args.Map{"groups": 2}
	expected.ShouldBeEqual(t, 0, "GroupByLock", actual)
}

func Test_ColGB_GroupByCount_Nil(t *testing.T) {
	// Arrange
	r := coredynamic.GroupByCount[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"len": len(r)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GroupByCount nil", actual)
}

func Test_ColGB_GroupByCount_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]string{"a", "b", "a"})
	r := coredynamic.GroupByCount(c, func(s string) string { return s })

	// Act
	actual := args.Map{
		"a": r["a"],
		"b": r["b"],
	}

	// Assert
	expected := args.Map{
		"a": 2,
		"b": 1,
	}
	expected.ShouldBeEqual(t, 0, "GroupByCount valid", actual)
}

// =============================================================================
// CollectionLock — thread-safe methods
// =============================================================================

func Test_ColLock_LengthLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{"r": c.LengthLock()}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "LengthLock", actual)
}

func Test_ColLock_IsEmptyLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()

	// Act
	actual := args.Map{"r": c.IsEmptyLock()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyLock", actual)
}

func Test_ColLock_AddLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddLock(1)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddLock", actual)
}

func Test_ColLock_AddsLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddsLock(1, 2)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddsLock", actual)
}

func Test_ColLock_AddManyLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddManyLock(1, 2, 3)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AddManyLock", actual)
}

func Test_ColLock_AddCollectionLock_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	c.AddCollectionLock(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddCollectionLock nil", actual)
}

func Test_ColLock_AddCollectionLock_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	other := coredynamic.CollectionFrom([]int{2})
	c.AddCollectionLock(other)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AddCollectionLock valid", actual)
}

func Test_ColLock_AddCollectionsLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	a := coredynamic.CollectionFrom([]int{1})
	c.AddCollectionsLock(a, nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddCollectionsLock", actual)
}

func Test_ColLock_AddIfLock_True(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddIfLock(true, 1)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddIfLock true", actual)
}

func Test_ColLock_AddIfLock_False(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	c.AddIfLock(false, 1)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddIfLock false", actual)
}

func Test_ColLock_RemoveAtLock_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})

	// Act
	actual := args.Map{
		"ok": c.RemoveAtLock(0),
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAtLock valid", actual)
}

func Test_ColLock_RemoveAtLock_Invalid(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"ok": c.RemoveAtLock(5)}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "RemoveAtLock invalid", actual)
}

func Test_ColLock_ClearLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	c.ClearLock()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ClearLock", actual)
}

func Test_ColLock_ItemsLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2})
	items := c.ItemsLock()

	// Act
	actual := args.Map{"len": len(items)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ItemsLock", actual)
}

func Test_ColLock_FirstLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{5})

	// Act
	actual := args.Map{"r": c.FirstLock()}

	// Assert
	expected := args.Map{"r": 5}
	expected.ShouldBeEqual(t, 0, "FirstLock", actual)
}

func Test_ColLock_LastLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{5, 10})

	// Act
	actual := args.Map{"r": c.LastLock()}

	// Assert
	expected := args.Map{"r": 10}
	expected.ShouldBeEqual(t, 0, "LastLock", actual)
}

func Test_ColLock_AddWithWgLock(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyCollection[int]()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c.AddWithWgLock(wg, 42)
	wg.Wait()

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AddWithWgLock", actual)
}

func Test_ColLock_LoopLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	sum := 0
	c.LoopLock(func(i int, item int) bool { sum += item; return false })

	// Act
	actual := args.Map{"sum": sum}

	// Assert
	expected := args.Map{"sum": 6}
	expected.ShouldBeEqual(t, 0, "LoopLock", actual)
}

func Test_ColLock_LoopLock_Break(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3})
	count := 0
	c.LoopLock(func(i int, item int) bool { count++; return item == 2 })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "LoopLock break", actual)
}

func Test_ColLock_FilterLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1, 2, 3, 4})
	f := c.FilterLock(func(i int) bool { return i > 2 })

	// Act
	actual := args.Map{"len": f.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "FilterLock", actual)
}

func Test_ColLock_StringsLock(t *testing.T) {
	// Arrange
	c := coredynamic.CollectionFrom([]int{1})
	s := c.StringsLock()

	// Act
	actual := args.Map{"first": s[0]}

	// Assert
	expected := args.Map{"first": "1"}
	expected.ShouldBeEqual(t, 0, "StringsLock", actual)
}

// =============================================================================
// CollectionTypes — factory shortcuts
// =============================================================================

func Test_ColTypes_Factories(t *testing.T) {
	// Arrange
	sc := coredynamic.NewStringCollection(5)
	esc := coredynamic.EmptyStringCollection()
	ic := coredynamic.NewIntCollection(5)
	eic := coredynamic.EmptyIntCollection()
	i64c := coredynamic.NewInt64Collection(5)
	bc := coredynamic.NewByteCollection(5)
	boolc := coredynamic.NewBoolCollection(5)
	f64c := coredynamic.NewFloat64Collection(5)
	amc := coredynamic.NewAnyMapCollection(5)
	smc := coredynamic.NewStringMapCollection(5)

	// Act
	actual := args.Map{
		"sc": sc.Length(), "esc": esc.Length(),
		"ic": ic.Length(), "eic": eic.Length(),
		"i64c": i64c.Length(), "bc": bc.Length(),
		"boolc": boolc.Length(), "f64c": f64c.Length(),
		"amc": amc.Length(), "smc": smc.Length(),
	}

	// Assert
	expected := args.Map{
		"sc": 0, "esc": 0,
		"ic": 0, "eic": 0,
		"i64c": 0, "bc": 0,
		"boolc": 0, "f64c": 0,
		"amc": 0, "smc": 0,
	}
	expected.ShouldBeEqual(t, 0, "CollectionTypes factories", actual)
}
