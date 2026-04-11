package coregenerictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── numericfuncs uncovered branches ──

func Test_InRangeExclusive_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{
		"atMin":    coregeneric.InRangeExclusive(1, 1, 10),
		"atMax":    coregeneric.InRangeExclusive(10, 1, 10),
		"inside":   coregeneric.InRangeExclusive(5, 1, 10),
		"below":    coregeneric.InRangeExclusive(0, 1, 10),
	}

	// Assert
	expected := args.Map{
		"atMin":    false,
		"atMax":    false,
		"inside":   true,
		"below":    false,
	}
	expected.ShouldBeEqual(t, 0, "InRangeExclusive returns correct value -- with args", actual)
}

func Test_SafeDivOrDefault_FromInRangeExclusiveOutO(t *testing.T) {
	// Act
	actual := args.Map{
		"normal": coregeneric.SafeDivOrDefault(10, 3, -1),
		"zero":   coregeneric.SafeDivOrDefault(10, 0, -1),
	}

	// Assert
	expected := args.Map{
		"normal": 3,
		"zero":   -1,
	}
	expected.ShouldBeEqual(t, 0, "SafeDivOrDefault returns correct value -- with args", actual)
}

func Test_IsNonNegative_FromInRangeExclusiveOutO(t *testing.T) {
	// Act
	actual := args.Map{
		"positive": coregeneric.IsNonNegative(5),
		"zero":     coregeneric.IsNonNegative(0),
		"negative": coregeneric.IsNonNegative(-1),
	}

	// Assert
	expected := args.Map{
		"positive": true,
		"zero":     true,
		"negative": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNonNegative returns correct value -- with args", actual)
}

func Test_Sign_FromInRangeExclusiveOutO(t *testing.T) {
	// Act
	actual := args.Map{
		"negative": coregeneric.Sign(-5),
		"zero":     coregeneric.Sign(0),
		"positive": coregeneric.Sign(5),
	}

	// Assert
	expected := args.Map{
		"negative": -1,
		"zero":     0,
		"positive": 1,
	}
	expected.ShouldBeEqual(t, 0, "Sign returns correct value -- with args", actual)
}

func Test_IsNotEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"same": coregeneric.IsNotEqual(5, 5),
		"diff": coregeneric.IsNotEqual(5, 6),
	}

	// Assert
	expected := args.Map{
		"same": false,
		"diff": true,
	}
	expected.ShouldBeEqual(t, 0, "IsNotEqual returns correct value -- with args", actual)
}

func Test_IsNumericEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"same": coregeneric.IsNumericEqual(5, 5),
		"diff": coregeneric.IsNumericEqual(5, 6),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNumericEqual returns correct value -- with args", actual)
}

// ── Collection uncovered branches ──

func Test_Collection_Capacity_Nil(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()

	// Act
	actual := args.Map{"cap": col.Capacity()}

	// Assert
	expected := args.Map{"cap": 0}
	expected.ShouldBeEqual(t, 0, "Collection returns empty -- Capacity empty", actual)
}

func Test_Collection_HasAnyItem(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1})

	// Act
	actual := args.Map{"result": col.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Collection returns correct value -- HasAnyItem", actual)
}

func Test_Collection_AddIfMany_Skip(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddIfMany(false, 1, 2, 3)

	// Act
	actual := args.Map{"len": col.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AddIfMany returns correct value -- skip", actual)
}

func Test_Collection_AddFunc_FromInRangeExclusiveOutO(t *testing.T) {
	// Arrange
	col := coregeneric.EmptyCollection[int]()
	col.AddFunc(func() int { return 42 })

	// Act
	actual := args.Map{"first": col.First()}

	// Assert
	expected := args.Map{"first": 42}
	expected.ShouldBeEqual(t, 0, "AddFunc returns correct value -- with args", actual)
}

func Test_Collection_CountFunc_FromInRangeExclusiveOutO(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3, 4, 5})
	count := col.CountFunc(func(v int) bool { return v > 3 })

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "CountFunc returns correct value -- with args", actual)
}

func Test_Collection_ConcatNew_FromInRangeExclusiveOutO(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2})
	result := col.ConcatNew(3, 4)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 4}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns correct value -- with args", actual)
}

func Test_Collection_Reverse_FromInRangeExclusiveOutO(t *testing.T) {
	// Arrange
	col := coregeneric.CollectionFrom([]int{1, 2, 3})
	col.Reverse()

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
	expected.ShouldBeEqual(t, 0, "Reverse returns correct value -- with args", actual)
}

// ── Hashmap uncovered branches ──

func Test_Hashmap_Set_ReturnsBool(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	isNew := hm.Set("a", 1)
	isUpdate := hm.Set("a", 2)

	// Act
	actual := args.Map{
		"isNew": isNew,
		"isUpdate": isUpdate,
	}
	// Set returns true if newly added

	// Assert
	expected := args.Map{
		"isNew": true,
		"isUpdate": false,
	}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- Set return", actual)
}

func Test_Hashmap_ForEachBreak_FromInRangeExclusiveOutO(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	hm.Set("b", 2)
	count := 0
	hm.ForEachBreak(func(k string, v int) bool {
		count++
		return true // break immediately
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns correct value -- ForEachBreak", actual)
}

func Test_Hashmap_ConcatNew_NilOther(t *testing.T) {
	// Arrange
	hm := coregeneric.EmptyHashmap[string, int]()
	hm.Set("a", 1)
	result := hm.ConcatNew(nil)

	// Act
	actual := args.Map{"len": result.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashmap returns nil -- ConcatNew nil other", actual)
}

// ── Hashset uncovered branches ──

func Test_Hashset_AddBool_FromInRangeExclusiveOutO(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	existed1 := hs.AddBool("a")
	existed2 := hs.AddBool("a")

	// Act
	actual := args.Map{
		"first": existed1,
		"second": existed2,
	}

	// Assert
	expected := args.Map{
		"first": false,
		"second": true,
	}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddBool", actual)
}

func Test_Hashset_AddIfMany_Skip(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddIfMany(false, "a", "b")

	// Act
	actual := args.Map{"len": hs.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- AddIfMany skip", actual)
}

func Test_Hashset_AddItemsMap_FalseValue(t *testing.T) {
	// Arrange
	hs := coregeneric.EmptyHashset[string]()
	hs.AddItemsMap(map[string]bool{"a": true, "b": false})

	// Act
	actual := args.Map{"len": hs.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Hashset returns non-empty -- AddItemsMap false value", actual)
}

func Test_Hashset_Resize_TooSmall(t *testing.T) {
	// Arrange
	hs := coregeneric.HashsetFrom([]string{"a", "b", "c"})
	hs.Resize(1) // smaller than current, should not resize

	// Act
	actual := args.Map{"len": hs.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "Hashset returns correct value -- Resize too small", actual)
}

// ── MapSimpleSlice nil ──

func Test_MapSimpleSlice_Nil(t *testing.T) {
	// Arrange
	result := coregeneric.MapSimpleSlice[int, string](nil, func(i int) string { return "" })

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "MapSimpleSlice returns nil -- nil", actual)
}

// ── DistinctSimpleSlice nil ──

func Test_DistinctSimpleSlice_Nil(t *testing.T) {
	// Arrange
	result := coregeneric.DistinctSimpleSlice[int](nil)

	// Act
	actual := args.Map{"empty": result.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DistinctSimpleSlice returns nil -- nil", actual)
}

// ── ContainsSimpleSliceItem nil ──

func Test_ContainsSimpleSliceItem_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": coregeneric.ContainsSimpleSliceItem[int](nil, 1)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ContainsSimpleSliceItem returns nil -- nil", actual)
}
