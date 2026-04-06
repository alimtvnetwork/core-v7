package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// CollectionMethods — AddIf, AddManyIf, AddCollection, ConcatNew, etc.
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_01_Collection_AddIf_True(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.AddIf(true, "b")

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_02_Collection_AddIf_False(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.AddIf(false, "x")

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_03_Collection_AddManyIf_True(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(true, 1, 2, 3)

	// Act
	actual := args.Map{"result": c.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C33_04_Collection_AddManyIf_False(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(false, 1, 2, 3)

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_05_Collection_AddManyIf_EmptyItems(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(true)

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_06_Collection_AddCollection(t *testing.T) {
	// Arrange
	c1 := coredynamic.NewCollection[string](4)
	c1.Add("a").Add("b")
	c2 := coredynamic.NewCollection[string](4)
	c2.Add("c").Add("d")
	c1.AddCollection(c2)

	// Act
	actual := args.Map{"result": c1.Length() != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C33_07_Collection_AddCollection_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.AddCollection(nil)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_08_Collection_AddCollection_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.AddCollection(coredynamic.EmptyCollection[string]())

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_09_Collection_AddCollections(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c2 := coredynamic.NewCollection[int](2)
	c2.Add(2).Add(3)
	c3 := coredynamic.NewCollection[int](2)
	c3.Add(4)
	c.AddCollections(c2, nil, c3)

	// Act
	actual := args.Map{"result": c.Length() != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C33_10_Collection_ConcatNew(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](2)
	c.Add("a").Add("b")
	c2 := c.ConcatNew("c", "d")

	// Act
	actual := args.Map{"result": c2.Length() != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	// original unchanged
	actual := args.Map{"result": c.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected original 2", actual)
}

func Test_C33_11_Collection_Clone(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(10).Add(20)
	cloned := c.Clone()
	cloned.Add(30)

	// Act
	actual := args.Map{"result": c.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected original 2", actual)
	actual := args.Map{"result": cloned.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned 3", actual)
}

func Test_C33_12_Collection_Clone_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[int]
	cloned := c.Clone()

	// Act
	actual := args.Map{"result": cloned.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_13_Collection_Capacity(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](10)

	// Act
	actual := args.Map{"result": c.Capacity() < 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected capacity >= 10", actual)
}

func Test_C33_14_Collection_Capacity_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.Collection[string]

	// Act
	actual := args.Map{"result": c.Capacity() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_15_Collection_AddCapacity(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](2)
	c.Add(1).Add(2)
	oldCap := c.Capacity()
	c.AddCapacity(10)

	// Act
	actual := args.Map{"result": c.Capacity() < oldCap+10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected capacity growth", actual)
}

func Test_C33_16_Collection_AddCapacity_Zero(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddCapacity(0)
	// no panic
}

func Test_C33_17_Collection_Resize(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](2)
	c.Add(1)
	c.Resize(20)

	// Act
	actual := args.Map{"result": c.Capacity() < 20}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected capacity >= 20", actual)
	actual := args.Map{"result": c.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected length 1", actual)
}

func Test_C33_18_Collection_Resize_Smaller(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](20)
	c.Resize(5) // should be no-op

	// Act
	actual := args.Map{"result": c.Capacity() < 20}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected capacity unchanged", actual)
}

func Test_C33_19_Collection_Reverse(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	c.Reverse()

	// Act
	actual := args.Map{"result": c.At(0) != 3 || c.At(1) != 2 || c.At(2) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected reversed [3,2,1]", actual)
}

func Test_C33_20_Collection_Reverse_SingleItem(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.Reverse()

	// Act
	actual := args.Map{"result": c.At(0) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [1]", actual)
}

func Test_C33_21_Collection_Reverse_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.Reverse() // no panic
}

func Test_C33_22_Collection_InsertAt(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("c")
	c.InsertAt(1, "b")

	// Act
	actual := args.Map{"result": c.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": c.At(1) != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b at index 1", actual)
}

func Test_C33_23_Collection_InsertAt_EmptyItems(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.InsertAt(0) // no items

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_24_Collection_IndexOfFunc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(10).Add(20).Add(30)
	idx := c.IndexOfFunc(func(v int) bool { return v == 20 })

	// Act
	actual := args.Map{"result": idx != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_25_Collection_IndexOfFunc_NotFound(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(10)
	idx := c.IndexOfFunc(func(v int) bool { return v == 99 })

	// Act
	actual := args.Map{"result": idx != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_C33_26_Collection_ContainsFunc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[string](4)
	c.Add("hello").Add("world")

	// Act
	actual := args.Map{"result": c.ContainsFunc(func(s string) bool { return s == "world" })}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": c.ContainsFunc(func(s string) bool { return s == "nope" })}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C33_27_Collection_SafeAt_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(42)

	// Act
	actual := args.Map{"result": c.SafeAt(0) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_28_Collection_SafeAt_OutOfRange(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	v := c.SafeAt(99)

	// Act
	actual := args.Map{"result": v != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected zero value", actual)
}

func Test_C33_29_Collection_SprintItems(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	result := c.SprintItems("[%d]")

	// Act
	actual := args.Map{"result": len(result) != 2 || result[0] != "[1]"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected result", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSort — SortFunc, SortAsc, SortDesc, IsSorted, etc.
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_30_Collection_SortFunc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"result": c.At(0) != 1 || c.At(1) != 2 || c.At(2) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted asc", actual)
}

func Test_C33_31_Collection_SortFunc_SingleItem(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.SortFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"result": c.At(0) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected [1]", actual)
}

func Test_C33_32_Collection_SortFuncLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	c.SortFuncLock(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"result": c.At(0) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_C33_33_Collection_SortedFunc_NoMutate(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	sorted := c.SortedFunc(func(a, b int) bool { return a < b })

	// Act
	actual := args.Map{"result": c.At(0) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "original should be unchanged", actual)
	actual := args.Map{"result": sorted.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "sorted should start with 1", actual)
}

func Test_C33_34_SortAsc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(5).Add(2).Add(8)
	coredynamic.SortAsc(c)

	// Act
	actual := args.Map{"result": c.At(0) != 2 || c.At(2) != 8}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected ascending", actual)
}

func Test_C33_35_SortDesc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(5).Add(2).Add(8)
	coredynamic.SortDesc(c)

	// Act
	actual := args.Map{"result": c.At(0) != 8 || c.At(2) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected descending", actual)
}

func Test_C33_36_SortAscLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(5).Add(2)
	coredynamic.SortAscLock(c)

	// Act
	actual := args.Map{"result": c.At(0) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected asc", actual)
}

func Test_C33_37_SortDescLock(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(2).Add(5)
	coredynamic.SortDescLock(c)

	// Act
	actual := args.Map{"result": c.At(0) != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected desc", actual)
}

func Test_C33_38_SortedAsc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1)
	s := coredynamic.SortedAsc(c)

	// Act
	actual := args.Map{"result": c.At(0) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "original unchanged", actual)
	actual := args.Map{"result": s.At(0) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_C33_39_SortedDesc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(3)
	s := coredynamic.SortedDesc(c)

	// Act
	actual := args.Map{"result": s.At(0) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected desc", actual)
}

func Test_C33_40_IsSorted_Asc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"result": c.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_C33_41_IsSorted_NotSorted(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)

	// Act
	actual := args.Map{"result": c.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not sorted", actual)
}

func Test_C33_42_IsSorted_SingleItem(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1)

	// Act
	actual := args.Map{"result": c.IsSorted(func(a, b int) bool { return a < b })}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "single item should be sorted", actual)
}

func Test_C33_43_IsSortedAsc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)

	// Act
	actual := args.Map{"result": coredynamic.IsSortedAsc(c)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected asc sorted", actual)
}

func Test_C33_44_IsSortedDesc(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1)

	// Act
	actual := args.Map{"result": coredynamic.IsSortedDesc(c)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected desc sorted", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionTypes — Factory shortcuts
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_45_NewStringCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewStringCollection(4)
	c.Add("hello")

	// Act
	actual := args.Map{"result": c.Length() != 1 || c.At(0) != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C33_46_EmptyStringCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyStringCollection()

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C33_47_NewIntCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewIntCollection(4)
	c.Add(42)

	// Act
	actual := args.Map{"result": c.At(0) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C33_48_EmptyIntCollection(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyIntCollection()

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C33_49_NewInt64Collection(t *testing.T) {
	// Arrange
	c := coredynamic.NewInt64Collection(4)
	c.Add(int64(100))

	// Act
	actual := args.Map{"result": c.At(0) != 100}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C33_50_NewByteCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewByteCollection(4)
	c.Add(byte(0xFF))

	// Act
	actual := args.Map{"result": c.At(0) != 0xFF}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C33_51_NewBoolCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewBoolCollection(4)
	c.Add(true).Add(false)

	// Act
	actual := args.Map{"result": c.At(0) != true || c.At(1) != false}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C33_52_NewFloat64Collection(t *testing.T) {
	// Arrange
	c := coredynamic.NewFloat64Collection(4)
	c.Add(3.14)

	// Act
	actual := args.Map{"result": c.At(0) != 3.14}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_C33_53_NewAnyMapCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewAnyMapCollection(4)
	c.Add(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_54_NewStringMapCollection(t *testing.T) {
	// Arrange
	c := coredynamic.NewStringMapCollection(4)
	c.Add(map[string]string{"a": "b"})

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicReflect — ReflectValue, ReflectKind, Loop, LoopMap, Filter
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_55_Dynamic_ReflectValue(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)
	rv := d.ReflectValue()

	// Act
	actual := args.Map{"result": rv.Kind() != reflect.Int}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected int kind", actual)
}

func Test_C33_56_Dynamic_ReflectKind(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("hello")

	// Act
	actual := args.Map{"result": d.ReflectKind() != reflect.String}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string", actual)
}

func Test_C33_57_Dynamic_ReflectTypeName(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)
	name := d.ReflectTypeName()

	// Act
	actual := args.Map{"result": name == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty type name", actual)
}

func Test_C33_58_Dynamic_ReflectType(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid("test")
	rt := d.ReflectType()

	// Act
	actual := args.Map{"result": rt != reflect.TypeOf("")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string type", actual)
}

func Test_C33_59_Dynamic_IsReflectTypeOf(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)

	// Act
	actual := args.Map{"result": d.IsReflectTypeOf(reflect.TypeOf(0))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
	actual := args.Map{"result": d.IsReflectTypeOf(reflect.TypeOf(""))}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}

func Test_C33_60_Dynamic_IsReflectKind(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)

	// Act
	actual := args.Map{"result": d.IsReflectKind(reflect.Int)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C33_61_Dynamic_ItemReflectValueUsingIndex(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{10, 20, 30})
	rv := d.ItemReflectValueUsingIndex(1)

	// Act
	actual := args.Map{"result": rv.Int() != 20}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 20", actual)
}

func Test_C33_62_Dynamic_ItemUsingIndex(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	v := d.ItemUsingIndex(0)

	// Act
	actual := args.Map{"result": v != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
}

func Test_C33_63_Dynamic_ItemReflectValueUsingKey(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"x": 42})
	rv := d.ItemReflectValueUsingKey("x")

	// Act
	actual := args.Map{"result": rv.Int() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_64_Dynamic_ItemUsingKey(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]string{"k": "v"})
	v := d.ItemUsingKey("k")

	// Act
	actual := args.Map{"result": v != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected v", actual)
}

func Test_C33_65_Dynamic_ReflectSetTo(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)
	var target int
	err := d.ReflectSetTo(&target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
}

func Test_C33_66_Dynamic_ReflectSetTo_Nil(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil receiver", actual)
}

func Test_C33_67_Dynamic_Loop_Slice(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{10, 20, 30})
	var sum int
	called := d.Loop(func(i int, item any) bool {
		sum += item.(int)
		return false
	})

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected called", actual)
	actual := args.Map{"result": sum != 60}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 60", actual)
}

func Test_C33_68_Dynamic_Loop_Break(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return i == 1
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_69_Dynamic_Loop_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()
	called := d.Loop(func(i int, item any) bool { return false })

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not called for invalid", actual)
}

func Test_C33_70_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4, 5})
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		v := item.Value().(int)
		return v%2 == 0, false // take evens
	})

	// Act
	actual := args.Map{"result": result.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 evens", actual)
}

func Test_C33_71_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4})
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 1 // break after index 1
	})

	// Act
	actual := args.Map{"result": result.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_72_Dynamic_FilterAsDynamicCollection_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})

	// Act
	actual := args.Map{"result": result.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid", actual)
}

func Test_C33_73_Dynamic_LoopMap(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})
	count := 0
	called := d.LoopMap(func(i int, k, v any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected called", actual)
	actual := args.Map{"result": count != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_74_Dynamic_LoopMap_Break(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2, "c": 3})
	count := 0
	d.LoopMap(func(i int, k, v any) bool {
		count++
		return true // break immediately
	})

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_75_Dynamic_LoopMap_Invalid(t *testing.T) {
	// Arrange
	d := coredynamic.InvalidDynamic()
	called := d.LoopMap(func(i int, k, v any) bool { return false })

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not called", actual)
}

func Test_C33_76_Dynamic_MapToKeyVal(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(map[string]int{"x": 10})
	kv, err := d.MapToKeyVal()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": kv == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_77_Dynamic_ConvertUsingFunc(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamicValid(42)
	converter := func(in any, typeMust reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(in)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(0))

	// Act
	actual := args.Map{"result": result == nil || !result.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid result", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectKindValidation, ReflectTypeValidation
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_78_ReflectKindValidation_Match(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.Int, 42)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_79_ReflectKindValidation_Mismatch(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectKindValidation(reflect.String, 42)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_80_ReflectTypeValidation_Match(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(0), 42)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_81_ReflectTypeValidation_Mismatch(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_82_ReflectTypeValidation_NilNotAllowed(t *testing.T) {
	// Arrange
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(0), nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil when not allowed", actual)
}

func Test_C33_83_ReflectTypeValidation_NilAllowed(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(0), nil)
	if err == nil {
		// nil type != int type, so still an error
	}
}

// ═══════════════════════════════════════════════════════════════════════
// PointerOrNonPointer, PointerOrNonPointerUsingReflectValue
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_84_PointerOrNonPointer_NonPointerOutput(t *testing.T) {
	// Arrange
	val := 42
	out, rv := coredynamic.PointerOrNonPointer(false, &val)

	// Act
	actual := args.Map{"result": out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	actual := args.Map{"result": rv.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid reflect value", actual)
}

func Test_C33_85_PointerOrNonPointer_StructDirect(t *testing.T) {
	// Arrange
	val := 42
	out, rv := coredynamic.PointerOrNonPointer(false, val)

	// Act
	actual := args.Map{"result": out != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	_ = rv
}

// ═══════════════════════════════════════════════════════════════════════
// IsAnyTypesOf, TypesIndexOf
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_86_IsAnyTypesOf_Found(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")

	// Act
	actual := args.Map{"result": coredynamic.IsAnyTypesOf(intType, intType, strType)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found", actual)
}

func Test_C33_87_IsAnyTypesOf_NotFound(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	boolType := reflect.TypeOf(true)

	// Act
	actual := args.Map{"result": coredynamic.IsAnyTypesOf(boolType, intType, strType)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C33_88_TypesIndexOf_Found(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	idx := coredynamic.TypesIndexOf(strType, intType, strType)

	// Act
	actual := args.Map{"result": idx != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_89_TypesIndexOf_NotFound(t *testing.T) {
	// Arrange
	intType := reflect.TypeOf(0)
	boolType := reflect.TypeOf(true)
	idx := coredynamic.TypesIndexOf(boolType, intType)

	// Act
	actual := args.Map{"result": idx != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// Type, TypeSameStatus, TypeNotEqualErr, TypeMustBeSame
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_90_Type(t *testing.T) {
	// Arrange
	rt := coredynamic.Type(42)

	// Act
	actual := args.Map{"result": rt != reflect.TypeOf(0)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected int type", actual)
}

func Test_C33_91_TypeSameStatus_Same(t *testing.T) {
	// Arrange
	st := coredynamic.TypeSameStatus(42, 100)

	// Act
	actual := args.Map{"result": st.IsSame}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected same", actual)
}

func Test_C33_92_TypeSameStatus_Different(t *testing.T) {
	// Arrange
	st := coredynamic.TypeSameStatus(42, "hello")

	// Act
	actual := args.Map{"result": st.IsSame}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected different", actual)
}

func Test_C33_93_TypeSameStatus_NilLeft(t *testing.T) {
	// Arrange
	st := coredynamic.TypeSameStatus(nil, 42)

	// Act
	actual := args.Map{"result": st.IsLeftUnknownNull}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected left null", actual)
}

func Test_C33_94_TypeSameStatus_NilRight(t *testing.T) {
	// Arrange
	st := coredynamic.TypeSameStatus(42, nil)

	// Act
	actual := args.Map{"result": st.IsRightUnknownNull}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected right null", actual)
}

func Test_C33_95_TypeSameStatus_Pointers(t *testing.T) {
	// Arrange
	val := 42
	st := coredynamic.TypeSameStatus(&val, 42)

	// Act
	actual := args.Map{"result": st.IsLeftPointer}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected left pointer", actual)
	actual := args.Map{"result": st.IsRightPointer}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected right non-pointer", actual)
}

func Test_C33_96_TypeNotEqualErr_Same(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr(42, 100)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_97_TypeNotEqualErr_Different(t *testing.T) {
	// Arrange
	err := coredynamic.TypeNotEqualErr(42, "hello")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_98_TypeMustBeSame_Same(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		actual := args.Map{"result": r := recover(); r != nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not panic for same types", actual)
	}()
	coredynamic.TypeMustBeSame(42, 100)
}

func Test_C33_99_TypeMustBeSame_Different(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		actual := args.Map{"result": r := recover(); r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic for different types", actual)
	}()
	coredynamic.TypeMustBeSame(42, "hello")
}

// ═══════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr, MustBeAcceptedTypes
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_100_NotAcceptedTypesErr_Accepted(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr(42, reflect.TypeOf(0), reflect.TypeOf(""))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_101_NotAcceptedTypesErr_Rejected(t *testing.T) {
	// Arrange
	err := coredynamic.NotAcceptedTypesErr(42, reflect.TypeOf(""), reflect.TypeOf(true))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_102_MustBeAcceptedTypes_Accepted(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		actual := args.Map{"result": r := recover(); r != nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not panic", actual)
	}()
	coredynamic.MustBeAcceptedTypes(42, reflect.TypeOf(0))
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectInterfaceVal, AnyToReflectVal
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_103_ReflectInterfaceVal_NonPointer(t *testing.T) {
	// Arrange
	v := coredynamic.ReflectInterfaceVal(42)

	// Act
	actual := args.Map{"result": v != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_104_ReflectInterfaceVal_Pointer(t *testing.T) {
	// Arrange
	val := 42
	v := coredynamic.ReflectInterfaceVal(&val)

	// Act
	actual := args.Map{"result": v != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_105_AnyToReflectVal(t *testing.T) {
	// Arrange
	rv := coredynamic.AnyToReflectVal("hello")

	// Act
	actual := args.Map{"result": rv.Kind() != reflect.String}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected string kind", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// ZeroSet, SafeZeroSet
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_106_ZeroSet(t *testing.T) {
	// Arrange
	type sample struct{ X int }
	s := &sample{X: 42}
	rv := reflect.ValueOf(s)
	coredynamic.ZeroSet(rv)

	// Act
	actual := args.Map{"result": s.X != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_107_SafeZeroSet(t *testing.T) {
	// Arrange
	type sample struct{ X int }
	s := &sample{X: 42}
	rv := reflect.ValueOf(s)
	coredynamic.SafeZeroSet(rv)

	// Act
	actual := args.Map{"result": s.X != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_108_SafeZeroSet_NilReflectValue(t *testing.T) {
	// pass zero reflect.Value — should not panic
	var rt reflect.Type
	coredynamic.SafeZeroSet(reflect.ValueOf(rt))
}

// ═══════════════════════════════════════════════════════════════════════
// LengthOfReflect
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_109_LengthOfReflect_Slice(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C33_110_LengthOfReflect_Array(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf([3]int{1, 2, 3})

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C33_111_LengthOfReflect_Map(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_112_LengthOfReflect_NonCollection(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act
	actual := args.Map{"result": coredynamic.LengthOfReflect(rv) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// SimpleRequest — constructors and methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_113_InvalidSimpleRequestNoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequestNoMessage()

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": r.Message() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty message", actual)
}

func Test_C33_114_InvalidSimpleRequest(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequest("bad input")

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": r.Message() != "bad input"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C33_115_NewSimpleRequest(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequest("data", true, "ok")

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual := args.Map{"result": r.Request() != "data"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected data", actual)
}

func Test_C33_116_NewSimpleRequestValid(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid(42)

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual := args.Map{"result": r.Value() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_117_SimpleRequest_NilReceiver(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"result": r.Message() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": r.Request() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": r.Value() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_118_SimpleRequest_GetErrorOnTypeMismatch_Match(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_119_SimpleRequest_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_120_SimpleRequest_GetErrorOnTypeMismatch_WithMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequest(42, true, "extra info")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_121_SimpleRequest_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_C33_122_SimpleRequest_IsReflectKind(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid(42)

	// Act
	actual := args.Map{"result": r.IsReflectKind(reflect.Int)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C33_123_SimpleRequest_IsReflectKind_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"result": r.IsReflectKind(reflect.Int)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C33_124_SimpleRequest_IsPointer_NonPtr(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleRequestValid(42)

	// Act
	actual := args.Map{"result": r.IsPointer()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C33_125_SimpleRequest_IsPointer_Ptr(t *testing.T) {
	// Arrange
	val := 42
	r := coredynamic.NewSimpleRequestValid(&val)

	// Act
	actual := args.Map{"result": r.IsPointer()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C33_126_SimpleRequest_IsPointer_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest

	// Act
	actual := args.Map{"result": r.IsPointer()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
}

func Test_C33_127_SimpleRequest_InvalidError_WithMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequest("some error")
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_128_SimpleRequest_InvalidError_EmptyMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequestNoMessage()
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty message", actual)
}

func Test_C33_129_SimpleRequest_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleRequest
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_C33_130_SimpleRequest_InvalidError_CachedErr(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleRequest("cached")
	err1 := r.InvalidError()
	err2 := r.InvalidError()

	// Act
	actual := args.Map{"result": err1 != err2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same cached error instance", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// SimpleResult — constructors and methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_131_InvalidSimpleResultNoMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResultNoMessage()

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": r.Message != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C33_132_InvalidSimpleResult(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResult("error msg")

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": r.Message != "error msg"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected message", actual)
}

func Test_C33_133_NewSimpleResultValid(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid(42)

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
	actual := args.Map{"result": r.Result != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_134_NewSimpleResult(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult("data", true, "")

	// Act
	actual := args.Map{"result": r.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)
}

func Test_C33_135_SimpleResult_GetErrorOnTypeMismatch_Match(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_136_SimpleResult_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_137_SimpleResult_GetErrorOnTypeMismatch_WithMessage(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResult(42, true, "info")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error with message", actual)
}

func Test_C33_138_SimpleResult_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_C33_139_SimpleResult_InvalidError_WithMessage(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResult("bad")
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_140_SimpleResult_InvalidError_Empty(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResultNoMessage()
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C33_141_SimpleResult_InvalidError_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	err := r.InvalidError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil receiver", actual)
}

func Test_C33_142_SimpleResult_InvalidError_Cached(t *testing.T) {
	// Arrange
	r := coredynamic.InvalidSimpleResult("cached")
	e1 := r.InvalidError()
	e2 := r.InvalidError()

	// Act
	actual := args.Map{"result": e1 != e2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same cached error", actual)
}

func Test_C33_143_SimpleResult_Clone(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid(42)
	cloned := r.Clone()

	// Act
	actual := args.Map{"result": cloned.Result != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C33_144_SimpleResult_Clone_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	cloned := r.Clone()

	// Act
	actual := args.Map{"result": cloned.Result != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil result", actual)
}

func Test_C33_145_SimpleResult_ClonePtr(t *testing.T) {
	// Arrange
	r := coredynamic.NewSimpleResultValid("data")
	cloned := r.ClonePtr()

	// Act
	actual := args.Map{"result": cloned == nil || cloned.Result != "data"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cloned ptr", actual)
}

func Test_C33_146_SimpleResult_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var r *coredynamic.SimpleResult
	cloned := r.ClonePtr()

	// Act
	actual := args.Map{"result": cloned != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection — extended methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_147_DynamicCollection_Skip(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	dc.Add(coredynamic.NewDynamicValid(3))
	skipped := dc.Skip(1)

	// Act
	actual := args.Map{"result": len(skipped) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_148_DynamicCollection_SkipCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	tsc := dc.SkipCollection(1)

	// Act
	actual := args.Map{"result": tsc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_149_DynamicCollection_SkipDynamic(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("a"))
	dc.Add(coredynamic.NewDynamicValid("b"))
	v := dc.SkipDynamic(1)

	// Act
	actual := args.Map{"result": v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_150_DynamicCollection_Take(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	dc.Add(coredynamic.NewDynamicValid(3))
	taken := dc.Take(2)

	// Act
	actual := args.Map{"result": len(taken) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_151_DynamicCollection_TakeCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	tc := dc.TakeCollection(1)

	// Act
	actual := args.Map{"result": tc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_152_DynamicCollection_TakeDynamic(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("x"))
	v := dc.TakeDynamic(1)

	// Act
	actual := args.Map{"result": v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_153_DynamicCollection_LimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	lc := dc.LimitCollection(1)

	// Act
	actual := args.Map{"result": lc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_154_DynamicCollection_SafeLimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	lc := dc.SafeLimitCollection(100) // larger than length

	// Act
	actual := args.Map{"result": lc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_155_DynamicCollection_LimitDynamic(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	v := dc.LimitDynamic(1)

	// Act
	actual := args.Map{"result": v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_156_DynamicCollection_Limit(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	l := dc.Limit(1)

	// Act
	actual := args.Map{"result": len(l) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_157_DynamicCollection_AddAnyNonNull(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyNonNull(nil, true) // should skip
	dc.AddAnyNonNull(42, true)

	// Act
	actual := args.Map{"result": dc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_158_DynamicCollection_AddAnyMany(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyMany("a", "b", "c")

	// Act
	actual := args.Map{"result": dc.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C33_159_DynamicCollection_AddAnyMany_Nil(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyMany(nil)
	// nil is a valid item, length should be 1
}

func Test_C33_160_DynamicCollection_AddPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	d := coredynamic.NewDynamicValid(42)
	dc.AddPtr(&d)

	// Act
	actual := args.Map{"result": dc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_161_DynamicCollection_AddPtr_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddPtr(nil) // should skip

	// Act
	actual := args.Map{"result": dc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_162_DynamicCollection_AddManyPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	d1 := coredynamic.NewDynamicValid(1)
	d2 := coredynamic.NewDynamicValid(2)
	dc.AddManyPtr(&d1, nil, &d2)

	// Act
	actual := args.Map{"result": dc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_163_DynamicCollection_AddManyPtr_NilSlice(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddManyPtr(nil)
	// nil variadic should skip
}

func Test_C33_164_DynamicCollection_AnyItems(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid("x"))
	items := dc.AnyItems()

	// Act
	actual := args.Map{"result": len(items) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_165_DynamicCollection_AnyItems_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	items := dc.AnyItems()

	// Act
	actual := args.Map{"result": len(items) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_166_DynamicCollection_AnyItemsCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	ac := dc.AnyItemsCollection()

	// Act
	actual := args.Map{"result": ac.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_167_DynamicCollection_AnyItemsCollection_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	ac := dc.AnyItemsCollection()

	// Act
	actual := args.Map{"result": ac.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_168_DynamicCollection_AddAnyWithTypeValidation_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(0), 42)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": dc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_169_DynamicCollection_AddAnyWithTypeValidation_Invalid(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(0), "string")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_170_DynamicCollection_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(0), 1, "bad", 3)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// should have added valid items
	actual := args.Map{"result": dc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_171_DynamicCollection_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0), 1, "bad", 3)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual := args.Map{"result": dc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_172_DynamicCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty items", actual)
}

func Test_C33_173_DynamicCollection_GetPagesSize(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	pages := dc.GetPagesSize(3)

	// Act
	actual := args.Map{"result": pages != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
}

func Test_C33_174_DynamicCollection_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	pages := dc.GetPagesSize(0)

	// Act
	actual := args.Map{"result": pages != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_175_DynamicCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	page := dc.GetSinglePageCollection(3, 1)

	// Act
	actual := args.Map{"result": page.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C33_176_DynamicCollection_GetSinglePageCollection_SmallerThanPage(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	page := dc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same collection", actual)
}

func Test_C33_177_DynamicCollection_JsonString(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	s, err := dc.JsonString()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty json", actual)
}

func Test_C33_178_DynamicCollection_JsonStringMust(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("hello"))
	s := dc.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C33_179_DynamicCollection_JsonModel(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	model := dc.JsonModel()

	// Act
	actual := args.Map{"result": len(model.Items) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1 item in model", actual)
}

func Test_C33_180_DynamicCollection_JsonModelAny(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	v := dc.JsonModelAny()

	// Act
	actual := args.Map{"result": v == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_181_DynamicCollection_Json(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	j := dc.Json()

	// Act
	actual := args.Map{"result": j.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error in json result", actual)
}

func Test_C33_182_DynamicCollection_JsonPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	jp := dc.JsonPtr()

	// Act
	actual := args.Map{"result": jp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_183_DynamicCollection_Strings(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid("hello"))
	strs := dc.Strings()

	// Act
	actual := args.Map{"result": len(strs) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C33_184_DynamicCollection_Strings_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	strs := dc.Strings()

	// Act
	actual := args.Map{"result": len(strs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_185_DynamicCollection_String(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	s := dc.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C33_186_DynamicCollection_ListStrings(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	strs := dc.ListStrings()

	// Act
	actual := args.Map{"result": len(strs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_187_DynamicCollection_ListStringsPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("x"))
	strs := dc.ListStringsPtr()

	// Act
	actual := args.Map{"result": len(strs) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_188_DynamicCollection_RemoveAt_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	ok := dc.RemoveAt(0)

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
	actual := args.Map{"result": dc.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_189_DynamicCollection_RemoveAt_Invalid(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	ok := dc.RemoveAt(99)

	// Act
	actual := args.Map{"result": ok}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C33_190_DynamicCollection_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	jsonResult := dc.JsonPtr()

	dc2 := coredynamic.NewDynamicCollection(4)
	result, err := dc2.ParseInjectUsingJson(jsonResult)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": result == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_191_DynamicCollection_JsonParseSelfInject(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	jsonResult := dc.JsonPtr()

	dc2 := coredynamic.NewDynamicCollection(4)
	err := dc2.JsonParseSelfInject(jsonResult)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
}

func Test_C33_192_DynamicCollection_ParseInjectUsingJson_BadJson(t *testing.T) {
	// Arrange
	badJson := corejson.NewPtr("not a dynamic collection")
	dc := coredynamic.NewDynamicCollection(4)
	_, err := dc.ParseInjectUsingJson(badJson)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for bad json", actual)
}

func Test_C33_193_DynamicCollection_JsonResultsCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	rc := dc.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_194_DynamicCollection_JsonResultsCollection_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	rc := dc.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil even when empty", actual)
}

func Test_C33_195_DynamicCollection_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	rc := dc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C33_196_DynamicCollection_JsonResultsPtrCollection_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	rc := dc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil even when empty", actual)
}

func Test_C33_197_DynamicCollection_AddAnySliceFromSingleItem(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnySliceFromSingleItem(true, []int{1, 2, 3})

	// Act
	actual := args.Map{"result": dc.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C33_198_DynamicCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnySliceFromSingleItem(true, nil)

	// Act
	actual := args.Map{"result": dc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C33_199_DynamicCollection_MarshalJSON(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	bytes, err := dc.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": len(bytes) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty bytes", actual)
}

func Test_C33_200_DynamicCollection_UnmarshalJSON(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	bytes, _ := dc.MarshalJSON()

	dc2 := coredynamic.NewDynamicCollection(4)
	err := dc2.UnmarshalJSON(bytes)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for Dynamic item payload without typed destination", actual)
}

func Test_C33_201_DynamicCollection_UnmarshalJSON_Bad(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.UnmarshalJSON([]byte("not json"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C33_202_DynamicCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	pages := dc.GetPagedCollection(3)

	// Act
	actual := args.Map{"result": len(pages) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pages", actual)
}

func Test_C33_203_DynamicCollection_GetPagedCollection_SmallerThanPage(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	pages := dc.GetPagedCollection(10)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C33_204_DynamicCollection_GetPagingInfo(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	info := dc.GetPagingInfo(3, 1)

	// Act
	actual := args.Map{"result": info.TotalPages != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 total pages", actual)
}
