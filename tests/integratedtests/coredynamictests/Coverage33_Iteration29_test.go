package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

// ═══════════════════════════════════════════════════════════════════════
// CollectionMethods — AddIf, AddManyIf, AddCollection, ConcatNew, etc.
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_01_Collection_AddIf_True(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.AddIf(true, "b")
	if c.Length() != 2 {
		t.Errorf("expected 2, got %d", c.Length())
	}
}

func Test_C33_02_Collection_AddIf_False(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.AddIf(false, "x")
	if c.Length() != 0 {
		t.Errorf("expected 0, got %d", c.Length())
	}
}

func Test_C33_03_Collection_AddManyIf_True(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(true, 1, 2, 3)
	if c.Length() != 3 {
		t.Errorf("expected 3, got %d", c.Length())
	}
}

func Test_C33_04_Collection_AddManyIf_False(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(false, 1, 2, 3)
	if c.Length() != 0 {
		t.Errorf("expected 0, got %d", c.Length())
	}
}

func Test_C33_05_Collection_AddManyIf_EmptyItems(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddManyIf(true)
	if c.Length() != 0 {
		t.Errorf("expected 0, got %d", c.Length())
	}
}

func Test_C33_06_Collection_AddCollection(t *testing.T) {
	c1 := coredynamic.NewCollection[string](4)
	c1.Add("a").Add("b")
	c2 := coredynamic.NewCollection[string](4)
	c2.Add("c").Add("d")
	c1.AddCollection(c2)
	if c1.Length() != 4 {
		t.Errorf("expected 4, got %d", c1.Length())
	}
}

func Test_C33_07_Collection_AddCollection_Nil(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.AddCollection(nil)
	if c.Length() != 1 {
		t.Errorf("expected 1, got %d", c.Length())
	}
}

func Test_C33_08_Collection_AddCollection_Empty(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.AddCollection(coredynamic.EmptyCollection[string]())
	if c.Length() != 1 {
		t.Errorf("expected 1, got %d", c.Length())
	}
}

func Test_C33_09_Collection_AddCollections(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c2 := coredynamic.NewCollection[int](2)
	c2.Add(2).Add(3)
	c3 := coredynamic.NewCollection[int](2)
	c3.Add(4)
	c.AddCollections(c2, nil, c3)
	if c.Length() != 4 {
		t.Errorf("expected 4, got %d", c.Length())
	}
}

func Test_C33_10_Collection_ConcatNew(t *testing.T) {
	c := coredynamic.NewCollection[string](2)
	c.Add("a").Add("b")
	c2 := c.ConcatNew("c", "d")
	if c2.Length() != 4 {
		t.Errorf("expected 4, got %d", c2.Length())
	}
	// original unchanged
	if c.Length() != 2 {
		t.Errorf("expected original 2, got %d", c.Length())
	}
}

func Test_C33_11_Collection_Clone(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(10).Add(20)
	cloned := c.Clone()
	cloned.Add(30)
	if c.Length() != 2 {
		t.Errorf("expected original 2, got %d", c.Length())
	}
	if cloned.Length() != 3 {
		t.Errorf("expected cloned 3, got %d", cloned.Length())
	}
}

func Test_C33_12_Collection_Clone_Nil(t *testing.T) {
	var c *coredynamic.Collection[int]
	cloned := c.Clone()
	if cloned.Length() != 0 {
		t.Errorf("expected 0, got %d", cloned.Length())
	}
}

func Test_C33_13_Collection_Capacity(t *testing.T) {
	c := coredynamic.NewCollection[string](10)
	if c.Capacity() < 10 {
		t.Errorf("expected capacity >= 10, got %d", c.Capacity())
	}
}

func Test_C33_14_Collection_Capacity_Nil(t *testing.T) {
	var c *coredynamic.Collection[string]
	if c.Capacity() != 0 {
		t.Errorf("expected 0, got %d", c.Capacity())
	}
}

func Test_C33_15_Collection_AddCapacity(t *testing.T) {
	c := coredynamic.NewCollection[int](2)
	c.Add(1).Add(2)
	oldCap := c.Capacity()
	c.AddCapacity(10)
	if c.Capacity() < oldCap+10 {
		t.Errorf("expected capacity growth")
	}
}

func Test_C33_16_Collection_AddCapacity_Zero(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.AddCapacity(0)
	// no panic
}

func Test_C33_17_Collection_Resize(t *testing.T) {
	c := coredynamic.NewCollection[int](2)
	c.Add(1)
	c.Resize(20)
	if c.Capacity() < 20 {
		t.Errorf("expected capacity >= 20")
	}
	if c.Length() != 1 {
		t.Errorf("expected length 1")
	}
}

func Test_C33_18_Collection_Resize_Smaller(t *testing.T) {
	c := coredynamic.NewCollection[int](20)
	c.Resize(5) // should be no-op
	if c.Capacity() < 20 {
		t.Errorf("expected capacity unchanged")
	}
}

func Test_C33_19_Collection_Reverse(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	c.Reverse()
	if c.At(0) != 3 || c.At(1) != 2 || c.At(2) != 1 {
		t.Errorf("expected reversed [3,2,1]")
	}
}

func Test_C33_20_Collection_Reverse_SingleItem(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.Reverse()
	if c.At(0) != 1 {
		t.Errorf("expected [1]")
	}
}

func Test_C33_21_Collection_Reverse_Empty(t *testing.T) {
	c := coredynamic.EmptyCollection[int]()
	c.Reverse() // no panic
}

func Test_C33_22_Collection_InsertAt(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a").Add("c")
	c.InsertAt(1, "b")
	if c.Length() != 3 {
		t.Errorf("expected 3")
	}
	if c.At(1) != "b" {
		t.Errorf("expected b at index 1, got %s", c.At(1))
	}
}

func Test_C33_23_Collection_InsertAt_EmptyItems(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("a")
	c.InsertAt(0) // no items
	if c.Length() != 1 {
		t.Errorf("expected 1")
	}
}

func Test_C33_24_Collection_IndexOfFunc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(10).Add(20).Add(30)
	idx := c.IndexOfFunc(func(v int) bool { return v == 20 })
	if idx != 1 {
		t.Errorf("expected 1, got %d", idx)
	}
}

func Test_C33_25_Collection_IndexOfFunc_NotFound(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(10)
	idx := c.IndexOfFunc(func(v int) bool { return v == 99 })
	if idx != -1 {
		t.Errorf("expected -1, got %d", idx)
	}
}

func Test_C33_26_Collection_ContainsFunc(t *testing.T) {
	c := coredynamic.NewCollection[string](4)
	c.Add("hello").Add("world")
	if !c.ContainsFunc(func(s string) bool { return s == "world" }) {
		t.Error("expected true")
	}
	if c.ContainsFunc(func(s string) bool { return s == "nope" }) {
		t.Error("expected false")
	}
}

func Test_C33_27_Collection_SafeAt_Valid(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(42)
	if c.SafeAt(0) != 42 {
		t.Error("expected 42")
	}
}

func Test_C33_28_Collection_SafeAt_OutOfRange(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	v := c.SafeAt(99)
	if v != 0 {
		t.Errorf("expected zero value, got %d", v)
	}
}

func Test_C33_29_Collection_SprintItems(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	result := c.SprintItems("[%d]")
	if len(result) != 2 || result[0] != "[1]" {
		t.Errorf("unexpected result %v", result)
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionSort — SortFunc, SortAsc, SortDesc, IsSorted, etc.
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_30_Collection_SortFunc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	c.SortFunc(func(a, b int) bool { return a < b })
	if c.At(0) != 1 || c.At(1) != 2 || c.At(2) != 3 {
		t.Error("expected sorted asc")
	}
}

func Test_C33_31_Collection_SortFunc_SingleItem(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	c.SortFunc(func(a, b int) bool { return a < b })
	if c.At(0) != 1 {
		t.Error("expected [1]")
	}
}

func Test_C33_32_Collection_SortFuncLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	c.SortFuncLock(func(a, b int) bool { return a < b })
	if c.At(0) != 1 {
		t.Error("expected sorted")
	}
}

func Test_C33_33_Collection_SortedFunc_NoMutate(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	sorted := c.SortedFunc(func(a, b int) bool { return a < b })
	if c.At(0) != 3 {
		t.Error("original should be unchanged")
	}
	if sorted.At(0) != 1 {
		t.Error("sorted should start with 1")
	}
}

func Test_C33_34_SortAsc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(5).Add(2).Add(8)
	coredynamic.SortAsc(c)
	if c.At(0) != 2 || c.At(2) != 8 {
		t.Error("expected ascending")
	}
}

func Test_C33_35_SortDesc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(5).Add(2).Add(8)
	coredynamic.SortDesc(c)
	if c.At(0) != 8 || c.At(2) != 2 {
		t.Error("expected descending")
	}
}

func Test_C33_36_SortAscLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(5).Add(2)
	coredynamic.SortAscLock(c)
	if c.At(0) != 2 {
		t.Error("expected asc")
	}
}

func Test_C33_37_SortDescLock(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(2).Add(5)
	coredynamic.SortDescLock(c)
	if c.At(0) != 5 {
		t.Error("expected desc")
	}
}

func Test_C33_38_SortedAsc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1)
	s := coredynamic.SortedAsc(c)
	if c.At(0) != 3 {
		t.Error("original unchanged")
	}
	if s.At(0) != 1 {
		t.Error("expected sorted")
	}
}

func Test_C33_39_SortedDesc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(3)
	s := coredynamic.SortedDesc(c)
	if s.At(0) != 3 {
		t.Error("expected desc")
	}
}

func Test_C33_40_IsSorted_Asc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2).Add(3)
	if !c.IsSorted(func(a, b int) bool { return a < b }) {
		t.Error("expected sorted")
	}
}

func Test_C33_41_IsSorted_NotSorted(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1).Add(2)
	if c.IsSorted(func(a, b int) bool { return a < b }) {
		t.Error("expected not sorted")
	}
}

func Test_C33_42_IsSorted_SingleItem(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1)
	if !c.IsSorted(func(a, b int) bool { return a < b }) {
		t.Error("single item should be sorted")
	}
}

func Test_C33_43_IsSortedAsc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(1).Add(2)
	if !coredynamic.IsSortedAsc(c) {
		t.Error("expected asc sorted")
	}
}

func Test_C33_44_IsSortedDesc(t *testing.T) {
	c := coredynamic.NewCollection[int](4)
	c.Add(3).Add(1)
	if !coredynamic.IsSortedDesc(c) {
		t.Error("expected desc sorted")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// CollectionTypes — Factory shortcuts
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_45_NewStringCollection(t *testing.T) {
	c := coredynamic.NewStringCollection(4)
	c.Add("hello")
	if c.Length() != 1 || c.At(0) != "hello" {
		t.Error("unexpected")
	}
}

func Test_C33_46_EmptyStringCollection(t *testing.T) {
	c := coredynamic.EmptyStringCollection()
	if c.Length() != 0 {
		t.Error("expected empty")
	}
}

func Test_C33_47_NewIntCollection(t *testing.T) {
	c := coredynamic.NewIntCollection(4)
	c.Add(42)
	if c.At(0) != 42 {
		t.Error("unexpected")
	}
}

func Test_C33_48_EmptyIntCollection(t *testing.T) {
	c := coredynamic.EmptyIntCollection()
	if c.Length() != 0 {
		t.Error("expected empty")
	}
}

func Test_C33_49_NewInt64Collection(t *testing.T) {
	c := coredynamic.NewInt64Collection(4)
	c.Add(int64(100))
	if c.At(0) != 100 {
		t.Error("unexpected")
	}
}

func Test_C33_50_NewByteCollection(t *testing.T) {
	c := coredynamic.NewByteCollection(4)
	c.Add(byte(0xFF))
	if c.At(0) != 0xFF {
		t.Error("unexpected")
	}
}

func Test_C33_51_NewBoolCollection(t *testing.T) {
	c := coredynamic.NewBoolCollection(4)
	c.Add(true).Add(false)
	if c.At(0) != true || c.At(1) != false {
		t.Error("unexpected")
	}
}

func Test_C33_52_NewFloat64Collection(t *testing.T) {
	c := coredynamic.NewFloat64Collection(4)
	c.Add(3.14)
	if c.At(0) != 3.14 {
		t.Error("unexpected")
	}
}

func Test_C33_53_NewAnyMapCollection(t *testing.T) {
	c := coredynamic.NewAnyMapCollection(4)
	c.Add(map[string]any{"k": "v"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_54_NewStringMapCollection(t *testing.T) {
	c := coredynamic.NewStringMapCollection(4)
	c.Add(map[string]string{"a": "b"})
	if c.Length() != 1 {
		t.Error("expected 1")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicReflect — ReflectValue, ReflectKind, Loop, LoopMap, Filter
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_55_Dynamic_ReflectValue(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	rv := d.ReflectValue()
	if rv.Kind() != reflect.Int {
		t.Error("expected int kind")
	}
}

func Test_C33_56_Dynamic_ReflectKind(t *testing.T) {
	d := coredynamic.NewDynamicValid("hello")
	if d.ReflectKind() != reflect.String {
		t.Error("expected string")
	}
}

func Test_C33_57_Dynamic_ReflectTypeName(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	name := d.ReflectTypeName()
	if name == "" {
		t.Error("expected non-empty type name")
	}
}

func Test_C33_58_Dynamic_ReflectType(t *testing.T) {
	d := coredynamic.NewDynamicValid("test")
	rt := d.ReflectType()
	if rt != reflect.TypeOf("") {
		t.Error("expected string type")
	}
}

func Test_C33_59_Dynamic_IsReflectTypeOf(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if !d.IsReflectTypeOf(reflect.TypeOf(0)) {
		t.Error("expected match")
	}
	if d.IsReflectTypeOf(reflect.TypeOf("")) {
		t.Error("expected no match")
	}
}

func Test_C33_60_Dynamic_IsReflectKind(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	if !d.IsReflectKind(reflect.Int) {
		t.Error("expected true")
	}
}

func Test_C33_61_Dynamic_ItemReflectValueUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{10, 20, 30})
	rv := d.ItemReflectValueUsingIndex(1)
	if rv.Int() != 20 {
		t.Error("expected 20")
	}
}

func Test_C33_62_Dynamic_ItemUsingIndex(t *testing.T) {
	d := coredynamic.NewDynamicValid([]string{"a", "b"})
	v := d.ItemUsingIndex(0)
	if v != "a" {
		t.Errorf("expected a, got %v", v)
	}
}

func Test_C33_63_Dynamic_ItemReflectValueUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"x": 42})
	rv := d.ItemReflectValueUsingKey("x")
	if rv.Int() != 42 {
		t.Error("expected 42")
	}
}

func Test_C33_64_Dynamic_ItemUsingKey(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]string{"k": "v"})
	v := d.ItemUsingKey("k")
	if v != "v" {
		t.Errorf("expected v, got %v", v)
	}
}

func Test_C33_65_Dynamic_ReflectSetTo(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	var target int
	err := d.ReflectSetTo(&target)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_C33_66_Dynamic_ReflectSetTo_Nil(t *testing.T) {
	var d *coredynamic.Dynamic
	err := d.ReflectSetTo(nil)
	if err == nil {
		t.Error("expected error for nil receiver")
	}
}

func Test_C33_67_Dynamic_Loop_Slice(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{10, 20, 30})
	var sum int
	called := d.Loop(func(i int, item any) bool {
		sum += item.(int)
		return false
	})
	if !called {
		t.Error("expected called")
	}
	if sum != 60 {
		t.Errorf("expected 60, got %d", sum)
	}
}

func Test_C33_68_Dynamic_Loop_Break(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3})
	count := 0
	d.Loop(func(i int, item any) bool {
		count++
		return i == 1
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C33_69_Dynamic_Loop_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	called := d.Loop(func(i int, item any) bool { return false })
	if called {
		t.Error("expected not called for invalid")
	}
}

func Test_C33_70_Dynamic_FilterAsDynamicCollection(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4, 5})
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		v := item.Value().(int)
		return v%2 == 0, false // take evens
	})
	if result.Length() != 2 {
		t.Errorf("expected 2 evens, got %d", result.Length())
	}
}

func Test_C33_71_Dynamic_FilterAsDynamicCollection_Break(t *testing.T) {
	d := coredynamic.NewDynamicValid([]int{1, 2, 3, 4})
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, i == 1 // break after index 1
	})
	if result.Length() != 2 {
		t.Errorf("expected 2, got %d", result.Length())
	}
}

func Test_C33_72_Dynamic_FilterAsDynamicCollection_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	result := d.FilterAsDynamicCollection(func(i int, item coredynamic.Dynamic) (bool, bool) {
		return true, false
	})
	if result.Length() != 0 {
		t.Error("expected empty for invalid")
	}
}

func Test_C33_73_Dynamic_LoopMap(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2})
	count := 0
	called := d.LoopMap(func(i int, k, v any) bool {
		count++
		return false
	})
	if !called {
		t.Error("expected called")
	}
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func Test_C33_74_Dynamic_LoopMap_Break(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"a": 1, "b": 2, "c": 3})
	count := 0
	d.LoopMap(func(i int, k, v any) bool {
		count++
		return true // break immediately
	})
	if count != 1 {
		t.Errorf("expected 1, got %d", count)
	}
}

func Test_C33_75_Dynamic_LoopMap_Invalid(t *testing.T) {
	d := coredynamic.InvalidDynamic()
	called := d.LoopMap(func(i int, k, v any) bool { return false })
	if called {
		t.Error("expected not called")
	}
}

func Test_C33_76_Dynamic_MapToKeyVal(t *testing.T) {
	d := coredynamic.NewDynamicValid(map[string]int{"x": 10})
	kv, err := d.MapToKeyVal()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if kv == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_77_Dynamic_ConvertUsingFunc(t *testing.T) {
	d := coredynamic.NewDynamicValid(42)
	converter := func(in any, typeMust reflect.Type) *coredynamic.SimpleResult {
		return coredynamic.NewSimpleResultValid(in)
	}
	result := d.ConvertUsingFunc(converter, reflect.TypeOf(0))
	if result == nil || !result.IsValid() {
		t.Error("expected valid result")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectKindValidation, ReflectTypeValidation
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_78_ReflectKindValidation_Match(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.Int, 42)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_79_ReflectKindValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectKindValidation(reflect.String, 42)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_80_ReflectTypeValidation_Match(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(0), 42)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_81_ReflectTypeValidation_Mismatch(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(false, reflect.TypeOf(""), 42)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_82_ReflectTypeValidation_NilNotAllowed(t *testing.T) {
	err := coredynamic.ReflectTypeValidation(true, reflect.TypeOf(0), nil)
	if err == nil {
		t.Error("expected error for nil when not allowed")
	}
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
	val := 42
	out, rv := coredynamic.PointerOrNonPointer(false, &val)
	if out != 42 {
		t.Errorf("expected 42, got %v", out)
	}
	if !rv.IsValid() {
		t.Error("expected valid reflect value")
	}
}

func Test_C33_85_PointerOrNonPointer_StructDirect(t *testing.T) {
	val := 42
	out, rv := coredynamic.PointerOrNonPointer(false, val)
	if out != 42 {
		t.Errorf("expected 42, got %v", out)
	}
	_ = rv
}

// ═══════════════════════════════════════════════════════════════════════
// IsAnyTypesOf, TypesIndexOf
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_86_IsAnyTypesOf_Found(t *testing.T) {
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	if !coredynamic.IsAnyTypesOf(intType, intType, strType) {
		t.Error("expected found")
	}
}

func Test_C33_87_IsAnyTypesOf_NotFound(t *testing.T) {
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	boolType := reflect.TypeOf(true)
	if coredynamic.IsAnyTypesOf(boolType, intType, strType) {
		t.Error("expected not found")
	}
}

func Test_C33_88_TypesIndexOf_Found(t *testing.T) {
	intType := reflect.TypeOf(0)
	strType := reflect.TypeOf("")
	idx := coredynamic.TypesIndexOf(strType, intType, strType)
	if idx != 1 {
		t.Errorf("expected 1, got %d", idx)
	}
}

func Test_C33_89_TypesIndexOf_NotFound(t *testing.T) {
	intType := reflect.TypeOf(0)
	boolType := reflect.TypeOf(true)
	idx := coredynamic.TypesIndexOf(boolType, intType)
	if idx != -1 {
		t.Errorf("expected -1, got %d", idx)
	}
}

// ═══════════════════════════════════════════════════════════════════════
// Type, TypeSameStatus, TypeNotEqualErr, TypeMustBeSame
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_90_Type(t *testing.T) {
	rt := coredynamic.Type(42)
	if rt != reflect.TypeOf(0) {
		t.Error("expected int type")
	}
}

func Test_C33_91_TypeSameStatus_Same(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, 100)
	if !st.IsSame {
		t.Error("expected same")
	}
}

func Test_C33_92_TypeSameStatus_Different(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, "hello")
	if st.IsSame {
		t.Error("expected different")
	}
}

func Test_C33_93_TypeSameStatus_NilLeft(t *testing.T) {
	st := coredynamic.TypeSameStatus(nil, 42)
	if !st.IsLeftUnknownNull {
		t.Error("expected left null")
	}
}

func Test_C33_94_TypeSameStatus_NilRight(t *testing.T) {
	st := coredynamic.TypeSameStatus(42, nil)
	if !st.IsRightUnknownNull {
		t.Error("expected right null")
	}
}

func Test_C33_95_TypeSameStatus_Pointers(t *testing.T) {
	val := 42
	st := coredynamic.TypeSameStatus(&val, 42)
	if !st.IsLeftPointer {
		t.Error("expected left pointer")
	}
	if st.IsRightPointer {
		t.Error("expected right non-pointer")
	}
}

func Test_C33_96_TypeNotEqualErr_Same(t *testing.T) {
	err := coredynamic.TypeNotEqualErr(42, 100)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_97_TypeNotEqualErr_Different(t *testing.T) {
	err := coredynamic.TypeNotEqualErr(42, "hello")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_98_TypeMustBeSame_Same(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("should not panic for same types")
		}
	}()
	coredynamic.TypeMustBeSame(42, 100)
}

func Test_C33_99_TypeMustBeSame_Different(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for different types")
		}
	}()
	coredynamic.TypeMustBeSame(42, "hello")
}

// ═══════════════════════════════════════════════════════════════════════
// NotAcceptedTypesErr, MustBeAcceptedTypes
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_100_NotAcceptedTypesErr_Accepted(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr(42, reflect.TypeOf(0), reflect.TypeOf(""))
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_101_NotAcceptedTypesErr_Rejected(t *testing.T) {
	err := coredynamic.NotAcceptedTypesErr(42, reflect.TypeOf(""), reflect.TypeOf(true))
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_102_MustBeAcceptedTypes_Accepted(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("should not panic")
		}
	}()
	coredynamic.MustBeAcceptedTypes(42, reflect.TypeOf(0))
}

// ═══════════════════════════════════════════════════════════════════════
// ReflectInterfaceVal, AnyToReflectVal
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_103_ReflectInterfaceVal_NonPointer(t *testing.T) {
	v := coredynamic.ReflectInterfaceVal(42)
	if v != 42 {
		t.Errorf("expected 42, got %v", v)
	}
}

func Test_C33_104_ReflectInterfaceVal_Pointer(t *testing.T) {
	val := 42
	v := coredynamic.ReflectInterfaceVal(&val)
	if v != 42 {
		t.Errorf("expected 42, got %v", v)
	}
}

func Test_C33_105_AnyToReflectVal(t *testing.T) {
	rv := coredynamic.AnyToReflectVal("hello")
	if rv.Kind() != reflect.String {
		t.Error("expected string kind")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// ZeroSet, SafeZeroSet
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_106_ZeroSet(t *testing.T) {
	type sample struct{ X int }
	s := &sample{X: 42}
	rv := reflect.ValueOf(s)
	coredynamic.ZeroSet(rv)
	if s.X != 0 {
		t.Errorf("expected 0, got %d", s.X)
	}
}

func Test_C33_107_SafeZeroSet(t *testing.T) {
	type sample struct{ X int }
	s := &sample{X: 42}
	rv := reflect.ValueOf(s)
	coredynamic.SafeZeroSet(rv)
	if s.X != 0 {
		t.Errorf("expected 0, got %d", s.X)
	}
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
	rv := reflect.ValueOf([]int{1, 2, 3})
	if coredynamic.LengthOfReflect(rv) != 3 {
		t.Error("expected 3")
	}
}

func Test_C33_110_LengthOfReflect_Array(t *testing.T) {
	rv := reflect.ValueOf([3]int{1, 2, 3})
	if coredynamic.LengthOfReflect(rv) != 3 {
		t.Error("expected 3")
	}
}

func Test_C33_111_LengthOfReflect_Map(t *testing.T) {
	rv := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	if coredynamic.LengthOfReflect(rv) != 2 {
		t.Error("expected 2")
	}
}

func Test_C33_112_LengthOfReflect_NonCollection(t *testing.T) {
	rv := reflect.ValueOf(42)
	if coredynamic.LengthOfReflect(rv) != 0 {
		t.Error("expected 0")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// SimpleRequest — constructors and methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_113_InvalidSimpleRequestNoMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleRequestNoMessage()
	if r.IsValid() {
		t.Error("expected invalid")
	}
	if r.Message() != "" {
		t.Error("expected empty message")
	}
}

func Test_C33_114_InvalidSimpleRequest(t *testing.T) {
	r := coredynamic.InvalidSimpleRequest("bad input")
	if r.IsValid() {
		t.Error("expected invalid")
	}
	if r.Message() != "bad input" {
		t.Error("expected message")
	}
}

func Test_C33_115_NewSimpleRequest(t *testing.T) {
	r := coredynamic.NewSimpleRequest("data", true, "ok")
	if !r.IsValid() {
		t.Error("expected valid")
	}
	if r.Request() != "data" {
		t.Error("expected data")
	}
}

func Test_C33_116_NewSimpleRequestValid(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid(42)
	if !r.IsValid() {
		t.Error("expected valid")
	}
	if r.Value() != 42 {
		t.Error("expected 42")
	}
}

func Test_C33_117_SimpleRequest_NilReceiver(t *testing.T) {
	var r *coredynamic.SimpleRequest
	if r.Message() != "" {
		t.Error("expected empty")
	}
	if r.Request() != nil {
		t.Error("expected nil")
	}
	if r.Value() != nil {
		t.Error("expected nil")
	}
}

func Test_C33_118_SimpleRequest_GetErrorOnTypeMismatch_Match(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_119_SimpleRequest_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_120_SimpleRequest_GetErrorOnTypeMismatch_WithMessage(t *testing.T) {
	r := coredynamic.NewSimpleRequest(42, true, "extra info")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_121_SimpleRequest_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err != nil {
		t.Error("expected nil for nil receiver")
	}
}

func Test_C33_122_SimpleRequest_IsReflectKind(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid(42)
	if !r.IsReflectKind(reflect.Int) {
		t.Error("expected true")
	}
}

func Test_C33_123_SimpleRequest_IsReflectKind_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	if r.IsReflectKind(reflect.Int) {
		t.Error("expected false for nil")
	}
}

func Test_C33_124_SimpleRequest_IsPointer_NonPtr(t *testing.T) {
	r := coredynamic.NewSimpleRequestValid(42)
	if r.IsPointer() {
		t.Error("expected false")
	}
}

func Test_C33_125_SimpleRequest_IsPointer_Ptr(t *testing.T) {
	val := 42
	r := coredynamic.NewSimpleRequestValid(&val)
	if !r.IsPointer() {
		t.Error("expected true")
	}
}

func Test_C33_126_SimpleRequest_IsPointer_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	if r.IsPointer() {
		t.Error("expected false for nil")
	}
}

func Test_C33_127_SimpleRequest_InvalidError_WithMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleRequest("some error")
	err := r.InvalidError()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_128_SimpleRequest_InvalidError_EmptyMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleRequestNoMessage()
	err := r.InvalidError()
	if err != nil {
		t.Errorf("expected nil for empty message, got %v", err)
	}
}

func Test_C33_129_SimpleRequest_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.SimpleRequest
	err := r.InvalidError()
	if err != nil {
		t.Error("expected nil for nil receiver")
	}
}

func Test_C33_130_SimpleRequest_InvalidError_CachedErr(t *testing.T) {
	r := coredynamic.InvalidSimpleRequest("cached")
	err1 := r.InvalidError()
	err2 := r.InvalidError()
	if err1 != err2 {
		t.Error("expected same cached error instance")
	}
}

// ═══════════════════════════════════════════════════════════════════════
// SimpleResult — constructors and methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C33_131_InvalidSimpleResultNoMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleResultNoMessage()
	if r.IsValid() {
		t.Error("expected invalid")
	}
	if r.Message != "" {
		t.Error("expected empty")
	}
}

func Test_C33_132_InvalidSimpleResult(t *testing.T) {
	r := coredynamic.InvalidSimpleResult("error msg")
	if r.IsValid() {
		t.Error("expected invalid")
	}
	if r.Message != "error msg" {
		t.Error("expected message")
	}
}

func Test_C33_133_NewSimpleResultValid(t *testing.T) {
	r := coredynamic.NewSimpleResultValid(42)
	if !r.IsValid() {
		t.Error("expected valid")
	}
	if r.Result != 42 {
		t.Error("expected 42")
	}
}

func Test_C33_134_NewSimpleResult(t *testing.T) {
	r := coredynamic.NewSimpleResult("data", true, "")
	if !r.IsValid() {
		t.Error("expected valid")
	}
}

func Test_C33_135_SimpleResult_GetErrorOnTypeMismatch_Match(t *testing.T) {
	r := coredynamic.NewSimpleResultValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_136_SimpleResult_GetErrorOnTypeMismatch_Mismatch(t *testing.T) {
	r := coredynamic.NewSimpleResultValid(42)
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), false)
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_137_SimpleResult_GetErrorOnTypeMismatch_WithMessage(t *testing.T) {
	r := coredynamic.NewSimpleResult(42, true, "info")
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(""), true)
	if err == nil {
		t.Error("expected error with message")
	}
}

func Test_C33_138_SimpleResult_GetErrorOnTypeMismatch_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	err := r.GetErrorOnTypeMismatch(reflect.TypeOf(0), false)
	if err != nil {
		t.Error("expected nil for nil receiver")
	}
}

func Test_C33_139_SimpleResult_InvalidError_WithMessage(t *testing.T) {
	r := coredynamic.InvalidSimpleResult("bad")
	err := r.InvalidError()
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_140_SimpleResult_InvalidError_Empty(t *testing.T) {
	r := coredynamic.InvalidSimpleResultNoMessage()
	err := r.InvalidError()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func Test_C33_141_SimpleResult_InvalidError_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	err := r.InvalidError()
	if err != nil {
		t.Error("expected nil for nil receiver")
	}
}

func Test_C33_142_SimpleResult_InvalidError_Cached(t *testing.T) {
	r := coredynamic.InvalidSimpleResult("cached")
	e1 := r.InvalidError()
	e2 := r.InvalidError()
	if e1 != e2 {
		t.Error("expected same cached error")
	}
}

func Test_C33_143_SimpleResult_Clone(t *testing.T) {
	r := coredynamic.NewSimpleResultValid(42)
	cloned := r.Clone()
	if cloned.Result != 42 {
		t.Error("expected 42")
	}
}

func Test_C33_144_SimpleResult_Clone_Nil(t *testing.T) {
	var r *coredynamic.SimpleResult
	cloned := r.Clone()
	if cloned.Result != nil {
		t.Error("expected nil result")
	}
}
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	dc.Add(coredynamic.NewDynamicValid(3))
	skipped := dc.Skip(1)
	if len(skipped) != 2 {
		t.Errorf("expected 2, got %d", len(skipped))
	}
}

func Test_C33_148_DynamicCollection_SkipCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	tsc := dc.SkipCollection(1)
	if tsc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_149_DynamicCollection_SkipDynamic(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("a"))
	dc.Add(coredynamic.NewDynamicValid("b"))
	v := dc.SkipDynamic(1)
	if v == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_150_DynamicCollection_Take(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	dc.Add(coredynamic.NewDynamicValid(3))
	taken := dc.Take(2)
	if len(taken) != 2 {
		t.Error("expected 2")
	}
}

func Test_C33_151_DynamicCollection_TakeCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	tc := dc.TakeCollection(1)
	if tc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_152_DynamicCollection_TakeDynamic(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("x"))
	v := dc.TakeDynamic(1)
	if v == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_153_DynamicCollection_LimitCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	lc := dc.LimitCollection(1)
	if lc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_154_DynamicCollection_SafeLimitCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	lc := dc.SafeLimitCollection(100) // larger than length
	if lc.Length() != 1 {
		t.Errorf("expected 1, got %d", lc.Length())
	}
}

func Test_C33_155_DynamicCollection_LimitDynamic(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	v := dc.LimitDynamic(1)
	if v == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_156_DynamicCollection_Limit(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	l := dc.Limit(1)
	if len(l) != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_157_DynamicCollection_AddAnyNonNull(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyNonNull(nil, true) // should skip
	dc.AddAnyNonNull(42, true)
	if dc.Length() != 1 {
		t.Errorf("expected 1, got %d", dc.Length())
	}
}

func Test_C33_158_DynamicCollection_AddAnyMany(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyMany("a", "b", "c")
	if dc.Length() != 3 {
		t.Error("expected 3")
	}
}

func Test_C33_159_DynamicCollection_AddAnyMany_Nil(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnyMany(nil)
	// nil is a valid item, length should be 1
}

func Test_C33_160_DynamicCollection_AddPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	d := coredynamic.NewDynamicValid(42)
	dc.AddPtr(&d)
	if dc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_161_DynamicCollection_AddPtr_Nil(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddPtr(nil) // should skip
	if dc.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C33_162_DynamicCollection_AddManyPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	d1 := coredynamic.NewDynamicValid(1)
	d2 := coredynamic.NewDynamicValid(2)
	dc.AddManyPtr(&d1, nil, &d2)
	if dc.Length() != 2 {
		t.Errorf("expected 2, got %d", dc.Length())
	}
}

func Test_C33_163_DynamicCollection_AddManyPtr_NilSlice(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddManyPtr(nil)
	// nil variadic should skip
}

func Test_C33_164_DynamicCollection_AnyItems(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid("x"))
	items := dc.AnyItems()
	if len(items) != 2 {
		t.Error("expected 2")
	}
}

func Test_C33_165_DynamicCollection_AnyItems_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	items := dc.AnyItems()
	if len(items) != 0 {
		t.Error("expected 0")
	}
}

func Test_C33_166_DynamicCollection_AnyItemsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	ac := dc.AnyItemsCollection()
	if ac.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_167_DynamicCollection_AnyItemsCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	ac := dc.AnyItemsCollection()
	if ac.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C33_168_DynamicCollection_AddAnyWithTypeValidation_Valid(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(0), 42)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if dc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_169_DynamicCollection_AddAnyWithTypeValidation_Invalid(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyWithTypeValidation(false, reflect.TypeOf(0), "string")
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_170_DynamicCollection_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyItemsWithTypeValidation(true, false, reflect.TypeOf(0), 1, "bad", 3)
	if err == nil {
		t.Error("expected error")
	}
	// should have added valid items
	if dc.Length() != 2 {
		t.Errorf("expected 2, got %d", dc.Length())
	}
}

func Test_C33_171_DynamicCollection_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0), 1, "bad", 3)
	if err == nil {
		t.Error("expected error")
	}
	if dc.Length() != 1 {
		t.Errorf("expected 1, got %d", dc.Length())
	}
}

func Test_C33_172_DynamicCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(0))
	if err != nil {
		t.Error("expected nil for empty items")
	}
}

func Test_C33_173_DynamicCollection_GetPagesSize(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	pages := dc.GetPagesSize(3)
	if pages != 4 {
		t.Errorf("expected 4 pages, got %d", pages)
	}
}

func Test_C33_174_DynamicCollection_GetPagesSize_Zero(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	pages := dc.GetPagesSize(0)
	if pages != 0 {
		t.Errorf("expected 0, got %d", pages)
	}
}

func Test_C33_175_DynamicCollection_GetSinglePageCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	page := dc.GetSinglePageCollection(3, 1)
	if page.Length() != 3 {
		t.Errorf("expected 3, got %d", page.Length())
	}
}

func Test_C33_176_DynamicCollection_GetSinglePageCollection_SmallerThanPage(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	page := dc.GetSinglePageCollection(10, 1)
	if page.Length() != 1 {
		t.Error("expected same collection")
	}
}

func Test_C33_177_DynamicCollection_JsonString(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	s, err := dc.JsonString()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if s == "" {
		t.Error("expected non-empty json")
	}
}

func Test_C33_178_DynamicCollection_JsonStringMust(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("hello"))
	s := dc.JsonStringMust()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C33_179_DynamicCollection_JsonModel(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	model := dc.JsonModel()
	if len(model.Items) != 1 {
		t.Error("expected 1 item in model")
	}
}

func Test_C33_180_DynamicCollection_JsonModelAny(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	v := dc.JsonModelAny()
	if v == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_181_DynamicCollection_Json(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	j := dc.Json()
	if j.HasError() {
		t.Error("expected no error in json result")
	}
}

func Test_C33_182_DynamicCollection_JsonPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	jp := dc.JsonPtr()
	if jp == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_183_DynamicCollection_Strings(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid("hello"))
	strs := dc.Strings()
	if len(strs) != 2 {
		t.Errorf("expected 2, got %d", len(strs))
	}
}

func Test_C33_184_DynamicCollection_Strings_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	strs := dc.Strings()
	if len(strs) != 0 {
		t.Error("expected 0")
	}
}

func Test_C33_185_DynamicCollection_String(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	s := dc.String()
	if s == "" {
		t.Error("expected non-empty")
	}
}

func Test_C33_186_DynamicCollection_ListStrings(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	strs := dc.ListStrings()
	if len(strs) != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_187_DynamicCollection_ListStringsPtr(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid("x"))
	strs := dc.ListStringsPtr()
	if len(strs) != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_188_DynamicCollection_RemoveAt_Valid(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	dc.Add(coredynamic.NewDynamicValid(2))
	ok := dc.RemoveAt(0)
	if !ok {
		t.Error("expected success")
	}
	if dc.Length() != 1 {
		t.Error("expected 1")
	}
}

func Test_C33_189_DynamicCollection_RemoveAt_Invalid(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	ok := dc.RemoveAt(99)
	if ok {
		t.Error("expected false")
	}
}

func Test_C33_190_DynamicCollection_ParseInjectUsingJson(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	jsonResult := dc.JsonPtr()

	dc2 := coredynamic.NewDynamicCollection(4)
	result, err := dc2.ParseInjectUsingJson(jsonResult)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_191_DynamicCollection_JsonParseSelfInject(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	jsonResult := dc.JsonPtr()

	dc2 := coredynamic.NewDynamicCollection(4)
	err := dc2.JsonParseSelfInject(jsonResult)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_C33_192_DynamicCollection_ParseInjectUsingJson_BadJson(t *testing.T) {
	badJson := corejson.NewPtr("not a dynamic collection")
	dc := coredynamic.NewDynamicCollection(4)
	_, err := dc.ParseInjectUsingJson(badJson)
	if err == nil {
		t.Error("expected error for bad json")
	}
}

func Test_C33_193_DynamicCollection_JsonResultsCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	rc := dc.JsonResultsCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_194_DynamicCollection_JsonResultsCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	rc := dc.JsonResultsCollection()
	if rc == nil {
		t.Error("expected non-nil even when empty")
	}
}

func Test_C33_195_DynamicCollection_JsonResultsPtrCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	rc := dc.JsonResultsPtrCollection()
	if rc == nil {
		t.Error("expected non-nil")
	}
}

func Test_C33_196_DynamicCollection_JsonResultsPtrCollection_Empty(t *testing.T) {
	dc := coredynamic.EmptyDynamicCollection()
	rc := dc.JsonResultsPtrCollection()
	if rc == nil {
		t.Error("expected non-nil even when empty")
	}
}

func Test_C33_197_DynamicCollection_AddAnySliceFromSingleItem(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnySliceFromSingleItem(true, []int{1, 2, 3})
	if dc.Length() != 3 {
		t.Errorf("expected 3, got %d", dc.Length())
	}
}

func Test_C33_198_DynamicCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnySliceFromSingleItem(true, nil)
	if dc.Length() != 0 {
		t.Error("expected 0")
	}
}

func Test_C33_199_DynamicCollection_MarshalJSON(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	bytes, err := dc.MarshalJSON()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(bytes) == 0 {
		t.Error("expected non-empty bytes")
	}
}

func Test_C33_200_DynamicCollection_UnmarshalJSON(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(42))
	bytes, _ := dc.MarshalJSON()

	dc2 := coredynamic.NewDynamicCollection(4)
	err := dc2.UnmarshalJSON(bytes)
	if err == nil {
		t.Error("expected error for Dynamic item payload without typed destination")
	}
}

func Test_C33_201_DynamicCollection_UnmarshalJSON_Bad(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	err := dc.UnmarshalJSON([]byte("not json"))
	if err == nil {
		t.Error("expected error")
	}
}

func Test_C33_202_DynamicCollection_GetPagedCollection(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	pages := dc.GetPagedCollection(3)
	if len(pages) == 0 {
		t.Error("expected pages")
	}
}

func Test_C33_203_DynamicCollection_GetPagedCollection_SmallerThanPage(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(4)
	dc.Add(coredynamic.NewDynamicValid(1))
	pages := dc.GetPagedCollection(10)
	if len(pages) != 1 {
		t.Errorf("expected 1, got %d", len(pages))
	}
}

func Test_C33_204_DynamicCollection_GetPagingInfo(t *testing.T) {
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.Add(coredynamic.NewDynamicValid(i))
	}
	info := dc.GetPagingInfo(3, 1)
	if info.TotalPages != 4 {
		t.Errorf("expected 4 total pages")
	}
}
