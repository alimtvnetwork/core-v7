package coredynamictests

import (
	"reflect"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// AnyCollection — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_01_AnyCollection_Basic(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("a").Add("b").Add("c")

	// Act
	actual := args.Map{"result": ac.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": ac.First() != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a", actual)
	actual := args.Map{"result": ac.Last() != "c"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected c", actual)
	actual := args.Map{"result": ac.At(1) != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected b", actual)
}

func Test_C32_02_AnyCollection_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"result": ac.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": ac.HasAnyItem()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no items", actual)
	actual := args.Map{"result": ac.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": ac.LastOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_03_AnyCollection_NilReceiver(t *testing.T) {
	// Arrange
	var ac *coredynamic.AnyCollection

	// Act
	actual := args.Map{"result": ac.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": ac.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C32_04_AnyCollection_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")

	// Act
	actual := args.Map{"result": ac.FirstOrDefault() != "x"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x", actual)
	actual := args.Map{"result": ac.LastOrDefault() != "x"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected x", actual)
	actual := args.Map{"result": ac.FirstOrDefaultDynamic() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": ac.LastOrDefaultDynamic() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_05_AnyCollection_FirstDynamic_LastDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{"result": ac.FirstDynamic() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": ac.LastDynamic() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_06_AnyCollection_Skip_Take(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3).Add(4)

	// Act
	actual := args.Map{"result": len(ac.Skip(2)) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(ac.Take(2)) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(ac.Limit(3)) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": ac.SkipDynamic(1) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": ac.TakeDynamic(2) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": ac.LimitDynamic(2) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_07_AnyCollection_SkipCollection_TakeCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3).Add(4)
	sc := ac.SkipCollection(2)

	// Act
	actual := args.Map{"result": sc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	tc := ac.TakeCollection(2)
	actual := args.Map{"result": tc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	lc := ac.LimitCollection(3)
	actual := args.Map{"result": lc.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	slc := ac.SafeLimitCollection(100)
	actual := args.Map{"result": slc.Length() != 4}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
}

func Test_C32_08_AnyCollection_Count_LastIndex_HasIndex(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{"result": ac.Count() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": ac.LastIndex() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": ac.HasIndex(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": ac.HasIndex(2)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C32_09_AnyCollection_RemoveAt(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add("a").Add("b").Add("c")

	// Act
	actual := args.Map{"result": ac.RemoveAt(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected success", actual)
	actual := args.Map{"result": ac.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": ac.RemoveAt(99)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
}

func Test_C32_10_AnyCollection_Items_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	items := ac.Items()

	// Act
	actual := args.Map{"result": len(items) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_11_AnyCollection_Items_NonEmpty(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)

	// Act
	actual := args.Map{"result": len(ac.Items()) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_12_AnyCollection_DynamicItems(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x").Add("y")
	di := ac.DynamicItems()

	// Act
	actual := args.Map{"result": len(di) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	empty := coredynamic.EmptyAnyCollection()
	actual := args.Map{"result": len(empty.DynamicItems()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_13_AnyCollection_DynamicCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	dc := ac.DynamicCollection()

	// Act
	actual := args.Map{"result": dc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyAnyCollection()
	actual := args.Map{"result": empty.DynamicCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_14_AnyCollection_AtAsDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(42)
	d := ac.AtAsDynamic(0)

	// Act
	actual := args.Map{"result": d.ValueInt() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C32_15_AnyCollection_ReflectSetAt(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("hello")
	var target string
	err := ac.ReflectSetAt(0, &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
}

func Test_C32_16_AnyCollection_Loop_Sync(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C32_17_AnyCollection_Loop_SyncBreak(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return i == 1
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_18_AnyCollection_Loop_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	ac.Loop(true, func(i int, item any) bool {
		return false
	})
}

func Test_C32_19_AnyCollection_Loop_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Loop(false, func(i int, item any) bool {

	// Act
		actual := args.Map{"result": false}

	// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
		return false
	})
}

func Test_C32_20_AnyCollection_LoopDynamic_Sync(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x").Add("y")
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_21_AnyCollection_LoopDynamic_SyncBreak(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2).Add(3)
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return i == 0
	})

	// Act
	actual := args.Map{"result": count != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_22_AnyCollection_LoopDynamic_Async(t *testing.T) {
	ac := coredynamic.NewAnyCollection(4)
	ac.Add(1).Add(2)
	ac.LoopDynamic(true, func(i int, item coredynamic.Dynamic) bool {
		return false
	})
}

func Test_C32_23_AnyCollection_LoopDynamic_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {

	// Act
		actual := args.Map{"result": false}

	// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should not be called", actual)
		return false
	})
}

func Test_C32_24_AnyCollection_AddAny(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.AddAny("val", true)

	// Act
	actual := args.Map{"result": ac.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_25_AnyCollection_AddNonNull(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.AddNonNull(nil)

	// Act
	actual := args.Map{"result": ac.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	ac.AddNonNull("x")
	actual := args.Map{"result": ac.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_26_AnyCollection_AddNonNullDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.AddNonNullDynamic(nil, true)

	// Act
	actual := args.Map{"result": ac.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	ac.AddNonNullDynamic("x", true)
	actual := args.Map{"result": ac.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_27_AnyCollection_AddAnyManyDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnyManyDynamic("a", "b")

	// Act
	actual := args.Map{"result": ac.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_28_AnyCollection_AddMany(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.AddMany("a", nil, "b")

	// Act
	actual := args.Map{"result": ac.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 (nil skipped)", actual)
}

func Test_C32_29_AnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	ac.AddAnySliceFromSingleItem([]string{"a", "b"})

	// Act
	actual := args.Map{"result": ac.Length() < 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items added", actual)
	ac.AddAnySliceFromSingleItem(nil)
}

func Test_C32_30_AnyCollection_ListStrings(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(42).Add("hello")
	strs := ac.ListStrings(false)

	// Act
	actual := args.Map{"result": len(strs) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	strsWithName := ac.ListStrings(true)
	actual := args.Map{"result": len(strsWithName) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_31_AnyCollection_Strings_String(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("a").Add("b")
	strs := ac.Strings()

	// Act
	actual := args.Map{"result": len(strs) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	s := ac.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	empty := coredynamic.EmptyAnyCollection()
	actual := args.Map{"result": len(empty.Strings()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_32_AnyCollection_JsonString(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	s, err := ac.JsonString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error or empty string", actual)
}

func Test_C32_33_AnyCollection_JsonStringMust(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	s := ac.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_34_AnyCollection_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("hello")
	b, err := ac.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "marshal failed", actual)
	ac2 := coredynamic.EmptyAnyCollection()
	err = ac2.UnmarshalJSON(b)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unmarshal failed:", actual)
}

func Test_C32_35_AnyCollection_JsonModel_JsonModelAny(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	m := ac.JsonModel()

	// Act
	actual := args.Map{"result": len(m) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	ma := ac.JsonModelAny()
	actual := args.Map{"result": ma == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_36_AnyCollection_Json_JsonPtr(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	j := ac.Json()
	_ = j
	jp := ac.JsonPtr()

	// Act
	actual := args.Map{"result": jp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_37_AnyCollection_JsonResultsCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	rc := ac.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyAnyCollection()
	actual := args.Map{"result": empty.JsonResultsCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_38_AnyCollection_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add("x")
	rc := ac.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyAnyCollection()
	actual := args.Map{"result": empty.JsonResultsPtrCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_39_AnyCollection_Paging(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(10)
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}

	// Act
	actual := args.Map{"result": ac.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": ac.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	page := ac.GetSinglePageCollection(3, 2)
	actual := args.Map{"result": page.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C32_40_AnyCollection_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1).Add(2)
	page := ac.GetSinglePageCollection(5, 1)

	// Act
	actual := args.Map{"result": page.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_41_AnyCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(7)
	for i := 0; i < 7; i++ {
		ac.Add(i)
	}
	pages := ac.GetPagedCollection(3)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
}

func Test_C32_42_AnyCollection_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	pages := ac.GetPagedCollection(5)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_43_AnyCollection_AddAnyItemsWithTypeValidation(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(4)
	strType := reflect.TypeOf("")
	err := ac.AddAnyItemsWithTypeValidation(false, false, strType, "a", "b")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	// type mismatch - stop on first error
	err = ac.AddAnyItemsWithTypeValidation(false, false, strType, 42)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
	// continue on error
	err = ac.AddAnyItemsWithTypeValidation(true, false, strType, "ok", 42)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	// empty items
	err = ac.AddAnyItemsWithTypeValidation(false, false, strType)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// MapAnyItems — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_44_MapAnyItems_Basic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)

	// Act
	actual := args.Map{"result": m.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": m.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	actual := args.Map{"result": m.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has any", actual)
	actual := args.Map{"result": m.HasKey("a")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has key a", actual)
	actual := args.Map{"result": m.HasKey("z")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for z", actual)
}

func Test_C32_45_MapAnyItems_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems

	// Act
	actual := args.Map{"result": m.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": m.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": m.HasKey("x")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C32_46_MapAnyItems_GetValue(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": m.GetValue("a") != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	actual := args.Map{"result": m.GetValue("z") != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_47_MapAnyItems_Get(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": "val"})
	v, has := m.Get("a")

	// Act
	actual := args.Map{"result": has || v != "val"}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected val, has=true", actual)
	v2, has2 := m.Get("z")
	actual := args.Map{"result": has2 || v2 != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil, false", actual)
}

func Test_C32_48_MapAnyItems_EmptyItems(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(nil)

	// Act
	actual := args.Map{"result": m.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{})
	actual := args.Map{"result": m2.IsEmpty()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C32_49_MapAnyItems_Add_Set(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.Add("x", 1)

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected newly added", actual)
	isNew2 := m.Add("x", 2)
	actual := args.Map{"result": isNew2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not new (override)", actual)
	isNew3 := m.Set("y", 3)
	actual := args.Map{"result": isNew3}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected newly added", actual)
}

func Test_C32_50_MapAnyItems_AddKeyAny(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	isNew := m.AddKeyAny(corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
}

func Test_C32_51_MapAnyItems_AddKeyAnyWithValidation(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	strType := reflect.TypeOf("")
	err := m.AddKeyAnyWithValidation(strType, corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	err = m.AddKeyAnyWithValidation(strType, corejson.KeyAny{Key: "k2", AnyInf: 42})
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected type mismatch error", actual)
}

func Test_C32_52_MapAnyItems_AddWithValidation(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	strType := reflect.TypeOf("")
	err := m.AddWithValidation(strType, "k", "v")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	err = m.AddWithValidation(strType, "k2", 42)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C32_53_MapAnyItems_AddJsonResultPtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	jr := corejson.NewPtr("hello")
	m.AddJsonResultPtr("k", jr)

	// Act
	actual := args.Map{"result": m.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	m.AddJsonResultPtr("k2", nil) // should skip
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 1", actual)
}

func Test_C32_54_MapAnyItems_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("b", 2)
	m.Add("a", 1)
	keys := m.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sortedKeys := m.AllKeysSorted()
	actual := args.Map{"result": sortedKeys[0] != "a" || sortedKeys[1] != "b"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
	vals := m.AllValues()
	actual := args.Map{"result": len(vals) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	empty := coredynamic.EmptyMapAnyItems()
	actual := args.Map{"result": len(empty.AllKeys()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": len(empty.AllKeysSorted()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": len(empty.AllValues()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_55_MapAnyItems_ReflectSetTo(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "hello")
	var target string
	err := m.ReflectSetTo("k", &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	err = m.ReflectSetTo("missing", &target)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for missing key", actual)
}

func Test_C32_56_MapAnyItems_GetUsingUnmarshallAt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("name", "John")
	var result string
	err := m.GetUsingUnmarshallAt("name", &result)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	err = m.GetUsingUnmarshallAt("missing", &result)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for missing key", actual)
}

func Test_C32_57_MapAnyItems_Deserialize(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("val", 42)
	var result int
	err := m.Deserialize("val", &result)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
}

func Test_C32_58_MapAnyItems_GetUsingUnmarshallManyAt(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", "hello")
	m.Add("b", 42)
	var s string
	var n int
	err := m.GetUsingUnmarshallManyAt(
		corejson.KeyAny{Key: "a", AnyInf: &s},
		corejson.KeyAny{Key: "b", AnyInf: &n},
	)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
}

func Test_C32_59_MapAnyItems_GetFieldsMap(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("obj", map[string]any{"x": 1})
	fm, err, found := m.GetFieldsMap("obj")

	// Act
	actual := args.Map{"result": found || err != nil}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found=true, err=nil, got found= err=", actual)
	_ = fm
	_, _, found2 := m.GetFieldsMap("missing")
	actual := args.Map{"result": found2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_C32_60_MapAnyItems_GetSafeFieldsMap(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("obj", map[string]any{"x": 1})
	fm, found := m.GetSafeFieldsMap("obj")

	// Act
	actual := args.Map{"result": found}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected found", actual)
	_ = fm
}

func Test_C32_61_MapAnyItems_GetItemRef(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	val := "hello"
	m.Add("k", &val)
	var target string
	err := m.GetItemRef("k", &target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	// missing key
	err = m.GetItemRef("missing", &target)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for missing key", actual)
	// nil referenceOut
	err = m.GetItemRef("k", nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil ref", actual)
	// non-pointer referenceOut
	err = m.GetItemRef("k", "not-a-pointer")
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for non-pointer", actual)
}

func Test_C32_62_MapAnyItems_GetManyItemsRefs(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	val := "hello"
	m.Add("k", &val)
	var target string
	err := m.GetManyItemsRefs(
		corejson.KeyAny{Key: "k", AnyInf: &target},
	)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	// empty
	err = m.GetManyItemsRefs()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty", actual)
}

func Test_C32_63_MapAnyItems_AddMapResult(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.AddMapResult(map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{"result": m.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	m.AddMapResult(nil)
	actual := args.Map{"result": m.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected still 2", actual)
}

func Test_C32_64_MapAnyItems_AddMapResultOption(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.AddMapResultOption(true, map[string]any{"a": 99, "b": 2})

	// Act
	actual := args.Map{"result": m.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	m.AddMapResultOption(false, map[string]any{"a": 100})
	m.AddMapResultOption(false, nil)
}

func Test_C32_65_MapAnyItems_AddManyMapResultsUsingOption(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.AddManyMapResultsUsingOption(true,
		map[string]any{"a": 1},
		map[string]any{"b": 2},
	)

	// Act
	actual := args.Map{"result": m.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	m.AddManyMapResultsUsingOption(true)
}

func Test_C32_66_MapAnyItems_GetNewMapUsingKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	sub := m.GetNewMapUsingKeys(false, "a", "c")

	// Act
	actual := args.Map{"result": sub.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	empty := m.GetNewMapUsingKeys(false)
	actual := args.Map{"result": empty.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	// not panic on missing when isPanicOnMissing=false
	sub2 := m.GetNewMapUsingKeys(false, "a", "missing")
	actual := args.Map{"result": sub2.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_67_MapAnyItems_JsonString(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s, err := m.JsonString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected json string", actual)
}

func Test_C32_68_MapAnyItems_JsonStringMust(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	s := m.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_69_MapAnyItems_JsonResultOfKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("k", "v")
	jr := m.JsonResultOfKey("k")

	// Act
	actual := args.Map{"result": jr == nil || jr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid result", actual)
	jr2 := m.JsonResultOfKey("missing")
	actual := args.Map{"result": jr2 == nil || !jr2.HasError()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for missing key", actual)
}

func Test_C32_70_MapAnyItems_JsonResultOfKeys(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	mr := m.JsonResultOfKeys("a", "b")

	// Act
	actual := args.Map{"result": mr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	mr2 := m.JsonResultOfKeys()
	_ = mr2
}

func Test_C32_71_MapAnyItems_Paging(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(10)
	for i := 0; i < 10; i++ {
		m.Add("k"+string(rune('a'+i)), i)
	}

	// Act
	actual := args.Map{"result": m.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": m.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_72_MapAnyItems_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(2)
	m.Add("a", 1)
	pages := m.GetPagedCollection(5)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_73_MapAnyItems_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(2)
	m.Add("a", 1)
	page := m.GetSinglePageCollection(5, 1, m.AllKeysSorted())

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_74_MapAnyItems_IsEqualRaw(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)

	// Act
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"a": 1, "b": 2})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"a": 1, "b": 3})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"a": 1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal (different length)", actual)
	actual := args.Map{"result": m.IsEqualRaw(map[string]any{"a": 1, "c": 2})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal (missing key)", actual)
	var nilM *coredynamic.MapAnyItems
	actual := args.Map{"result": nilM.IsEqualRaw(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for both nil", actual)
	actual := args.Map{"result": nilM.IsEqualRaw(map[string]any{"a": 1})}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil vs non-nil", actual)
}

func Test_C32_75_MapAnyItems_IsEqual(t *testing.T) {
	// Arrange
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": m1.IsEqual(m2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
	m3 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 2})
	actual := args.Map{"result": m1.IsEqual(m3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	var nilM *coredynamic.MapAnyItems
	actual := args.Map{"result": nilM.IsEqual(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": nilM.IsEqual(m1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	m4 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	actual := args.Map{"result": m1.IsEqual(m4)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal (different length)", actual)
}

func Test_C32_76_MapAnyItems_Clear_DeepClear_Dispose(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Clear()

	// Act
	actual := args.Map{"result": m.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	m.Add("b", 2)
	m.DeepClear()
	actual := args.Map{"result": m.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	m.Dispose()
	var nilM *coredynamic.MapAnyItems
	nilM.Clear()
	nilM.DeepClear()
	nilM.Dispose()
}

func Test_C32_77_MapAnyItems_Strings_String(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	strs := m.Strings()

	// Act
	actual := args.Map{"result": len(strs) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	s := m.String()
	actual := args.Map{"result": s == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_78_MapAnyItems_MapAnyItems_Self(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)

	// Act
	actual := args.Map{"result": m.MapAnyItems() != m}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected self reference", actual)
}

func Test_C32_79_MapAnyItems_MapStringAnyDiff(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	diff := m.MapStringAnyDiff()

	// Act
	actual := args.Map{"result": diff == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_80_MapAnyItems_RawMapStringAnyDiff_NilReceiver(t *testing.T) {
	// Arrange
	var m *coredynamic.MapAnyItems
	diff := m.RawMapStringAnyDiff()

	// Act
	actual := args.Map{"result": len(diff) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map for nil receiver", actual)
}

func Test_C32_81_MapAnyItems_IsRawEqual(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)

	// Act
	actual := args.Map{"result": m.IsRawEqual(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_C32_82_MapAnyItems_HasAnyChanges(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)

	// Act
	actual := args.Map{"result": m.HasAnyChanges(false, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no changes", actual)
	actual := args.Map{"result": m.HasAnyChanges(false, map[string]any{"a": 2})}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected changes", actual)
}

func Test_C32_83_MapAnyItems_Json_JsonPtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	j := m.Json()
	_ = j
	jp := m.JsonPtr()

	// Act
	actual := args.Map{"result": jp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_84_MapAnyItems_JsonModel_JsonModelAny(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	model := m.JsonModel()

	// Act
	actual := args.Map{"result": model == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": m.JsonModelAny() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	empty := coredynamic.EmptyMapAnyItems()
	em := empty.JsonModel()
	_ = em
}

func Test_C32_85_MapAnyItems_JsonMapResults(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	mr, err := m.JsonMapResults()

	// Act
	actual := args.Map{"result": err != nil || mr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	empty := coredynamic.EmptyMapAnyItems()
	mr2, _ := empty.JsonMapResults()
	_ = mr2
}

func Test_C32_86_MapAnyItems_JsonResultsCollection(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	rc := m.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyMapAnyItems()
	actual := args.Map{"result": empty.JsonResultsCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_87_MapAnyItems_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	rc := m.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyMapAnyItems()
	actual := args.Map{"result": empty.JsonResultsPtrCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_88_MapAnyItems_ClonePtr(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	cloned, err := m.ClonePtr()

	// Act
	actual := args.Map{"result": err != nil || cloned == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error or nil:", actual)
	var nilM *coredynamic.MapAnyItems
	_, err = nilM.ClonePtr()
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil receiver", actual)
}

func Test_C32_89_MapAnyItems_NewUsingAnyTypeMap(t *testing.T) {
	// Arrange
	m, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(map[string]int{"a": 1})

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": m.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	_, err = coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_90_ValueStatus(t *testing.T) {
	// Arrange
	vs := coredynamic.InvalidValueStatus("test")

	// Act
	actual := args.Map{"result": vs.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": vs.Message != "test"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)
	vs2 := coredynamic.InvalidValueStatusNoMessage()
	actual := args.Map{"result": vs2.Message != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// KeyVal — additional coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_91_KeyVal_KeyDynamic_ValueDynamic(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "name", Value: 42}
	kd := kv.KeyDynamic()

	// Act
	actual := args.Map{"result": kd.ValueString() != "name"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected name", actual)
	vd := kv.ValueDynamic()
	actual := args.Map{"result": vd.ValueInt() != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C32_92_KeyVal_KeyDynamicPtr_ValueDynamicPtr(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"result": kv.KeyDynamicPtr() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": kv.ValueDynamicPtr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	var nilKV *coredynamic.KeyVal
	actual := args.Map{"result": nilKV.KeyDynamicPtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": nilKV.ValueDynamicPtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_93_KeyVal_IsKeyNull_IsValueNull(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "a", Value: nil}

	// Act
	actual := args.Map{"result": kv.IsKeyNull()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": kv.IsValueNull()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_C32_94_KeyVal_IsKeyNullOrEmptyString(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "", Value: nil}

	// Act
	actual := args.Map{"result": kv.IsKeyNullOrEmptyString()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	kv2 := coredynamic.KeyVal{Key: "x", Value: nil}
	actual := args.Map{"result": kv2.IsKeyNullOrEmptyString()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_C32_95_KeyVal_ValueInt_ValueUInt_ValueBool_ValueInt64_ValueStrings(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}

	// Act
	actual := args.Map{"result": kv.ValueInt() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
	kv2 := coredynamic.KeyVal{Key: "k", Value: uint(10)}
	actual := args.Map{"result": kv2.ValueUInt() != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 10", actual)
	kv3 := coredynamic.KeyVal{Key: "k", Value: true}
	actual := args.Map{"result": kv3.ValueBool()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	kv4 := coredynamic.KeyVal{Key: "k", Value: int64(99)}
	actual := args.Map{"result": kv4.ValueInt64() != 99}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 99", actual)
	kv5 := coredynamic.KeyVal{Key: "k", Value: []string{"a"}}
	actual := args.Map{"result": len(kv5.ValueStrings()) != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	// mismatches
	kvBad := coredynamic.KeyVal{Key: "k", Value: "str"}
	actual := args.Map{"result": kvBad.ValueInt() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": kvBad.ValueUInt() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": kvBad.ValueBool()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	actual := args.Map{"result": kvBad.ValueInt64() != -1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	actual := args.Map{"result": kvBad.ValueStrings() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_96_KeyVal_String_NilReceiver(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	actual := args.Map{"result": kv.String() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C32_97_KeyVal_String(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "name", Value: "val"}
	s := kv.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_98_KeyVal_KeyString_ValueString(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}

	// Act
	actual := args.Map{"result": kv.KeyString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	actual := args.Map{"result": kv.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	var nilKV *coredynamic.KeyVal
	actual := args.Map{"result": nilKV.KeyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual := args.Map{"result": nilKV.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	kvNilKey := &coredynamic.KeyVal{Key: nil, Value: nil}
	actual := args.Map{"result": kvNilKey.KeyString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil key", actual)
	actual := args.Map{"result": kvNilKey.ValueString() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil value", actual)
}

func Test_C32_99_KeyVal_ValueReflectValue(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: 42}
	rv := kv.ValueReflectValue()

	// Act
	actual := args.Map{"result": rv.Int() != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_C32_100_KeyVal_ValueNullErr_KeyNullErr(t *testing.T) {
	// Arrange
	var nilKV *coredynamic.KeyVal

	// Act
	actual := args.Map{"result": nilKV.ValueNullErr() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual := args.Map{"result": nilKV.KeyNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	kvNullVal := &coredynamic.KeyVal{Key: "k", Value: nil}
	actual := args.Map{"result": kvNullVal.ValueNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for null value", actual)
	kvNullKey := &coredynamic.KeyVal{Key: nil, Value: "v"}
	actual := args.Map{"result": kvNullKey.KeyNullErr() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for null key", actual)
	kvOk := &coredynamic.KeyVal{Key: "k", Value: "v"}
	actual := args.Map{"result": kvOk.ValueNullErr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": kvOk.KeyNullErr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_101_KeyVal_CastKeyVal(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var k, v string
	err := kv.CastKeyVal(&k, &v)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	var nilKV *coredynamic.KeyVal
	actual := args.Map{"result": nilKV.CastKeyVal(&k, &v) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C32_102_KeyVal_ReflectSetKey(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var target string
	err := kv.ReflectSetKey(&target)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	var nilKV *coredynamic.KeyVal
	actual := args.Map{"result": nilKV.ReflectSetKey(&target) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C32_103_KeyVal_KeyReflectSet_ValueReflectSet_ReflectSetTo(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "hello", Value: "world"}
	var k, v string

	// Act
	actual := args.Map{"result": kv.KeyReflectSet(&k) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	actual := args.Map{"result": kv.ValueReflectSet(&v) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	actual := args.Map{"result": kv.ReflectSetTo(&v) != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	var nilKV *coredynamic.KeyVal
	actual := args.Map{"result": nilKV.KeyReflectSet(&k) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual := args.Map{"result": nilKV.ValueReflectSet(&v) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	actual := args.Map{"result": nilKV.ReflectSetTo(&v) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_C32_104_KeyVal_ReflectSetToMust(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	var target string
	kv.ReflectSetToMust(&target)

	// Act
	actual := args.Map{"result": target != "v"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected v", actual)
}

func Test_C32_105_KeyVal_Json(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "k", Value: "v"}
	j := kv.Json()
	_ = j
	jp := kv.JsonPtr()

	// Act
	actual := args.Map{"result": jp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": kv.JsonModel() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": kv.JsonModelAny() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_106_KeyVal_Serialize(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "k", Value: "v"}
	b, err := kv.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "serialize failed", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValCollection — additional coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_107_KeyValCollection_Basic(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})

	// Act
	actual := args.Map{"result": kvc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": kvc.IsEmpty()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	actual := args.Map{"result": kvc.HasAnyItem()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected has any", actual)
}

func Test_C32_108_KeyValCollection_AddPtr(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(4)
	kv := &coredynamic.KeyVal{Key: "a", Value: 1}
	kvc.AddPtr(kv)
	kvc.AddPtr(nil)

	// Act
	actual := args.Map{"result": kvc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_109_KeyValCollection_AddMany_AddManyPtr(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.AddMany(
		coredynamic.KeyVal{Key: "a", Value: 1},
		coredynamic.KeyVal{Key: "b", Value: 2},
	)

	// Act
	actual := args.Map{"result": kvc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	kv := &coredynamic.KeyVal{Key: "c", Value: 3}
	kvc.AddManyPtr(kv, nil)
	actual := args.Map{"result": kvc.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C32_110_KeyValCollection_Items_NilReceiver(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"result": kvc.Items() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": kvc.Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_111_KeyValCollection_AllKeys_AllKeysSorted_AllValues(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: 2})
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	keys := kvc.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	sorted := kvc.AllKeysSorted()
	actual := args.Map{"result": sorted[0] != "a"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
	vals := kvc.AllValues()
	actual := args.Map{"result": len(vals) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	empty := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"result": len(empty.AllKeys()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": len(empty.AllKeysSorted()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	actual := args.Map{"result": len(empty.AllValues()) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_112_KeyValCollection_MapAnyItems(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(4)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	m := kvc.MapAnyItems()

	// Act
	actual := args.Map{"result": m.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"result": empty.MapAnyItems().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_113_KeyValCollection_String_NilReceiver(t *testing.T) {
	// Arrange
	var kvc *coredynamic.KeyValCollection

	// Act
	actual := args.Map{"result": kvc.String() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_C32_114_KeyValCollection_String(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_115_KeyValCollection_Clone_ClonePtr(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	c := kvc.Clone()

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	cp := kvc.ClonePtr()
	actual := args.Map{"result": cp.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	var nilKVC *coredynamic.KeyValCollection
	actual := args.Map{"result": nilKVC.ClonePtr() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_116_KeyValCollection_NonPtr_Ptr(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	np := kvc.NonPtr()
	_ = np
	p := kvc.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_117_KeyValCollection_JsonString(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s, err := kvc.JsonString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "json string failed", actual)
}

func Test_C32_118_KeyValCollection_JsonStringMust(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	s := kvc.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_119_KeyValCollection_Serialize(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	b, err := kvc.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "serialize failed", actual)
}

func Test_C32_120_KeyValCollection_Paging(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(10)
	for i := 0; i < 10; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}

	// Act
	actual := args.Map{"result": kvc.GetPagesSize(3) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4", actual)
	actual := args.Map{"result": kvc.GetPagesSize(0) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_121_KeyValCollection_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	page := kvc.GetSinglePageCollection(5, 1)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_122_KeyValCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(7)
	for i := 0; i < 7; i++ {
		kvc.Add(coredynamic.KeyVal{Key: "k", Value: i})
	}
	pages := kvc.GetPagedCollection(3)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 pages", actual)
}

func Test_C32_123_KeyValCollection_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	pages := kvc.GetPagedCollection(5)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_124_KeyValCollection_JsonMapResults(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	mr, err := kvc.JsonMapResults()

	// Act
	actual := args.Map{"result": err != nil || mr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "error:", actual)
	empty := coredynamic.EmptyKeyValCollection()
	mr2, _ := empty.JsonMapResults()
	_ = mr2
}

func Test_C32_125_KeyValCollection_JsonResultsCollection(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	rc := kvc.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"result": empty.JsonResultsCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_126_KeyValCollection_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	rc := kvc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyKeyValCollection()
	actual := args.Map{"result": empty.JsonResultsPtrCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_127_KeyValCollection_Json_JsonPtr(t *testing.T) {
	// Arrange
	kvc := coredynamic.KeyValCollection{}
	j := kvc.Json()
	_ = j
	jp := kvc.JsonPtr()

	// Act
	actual := args.Map{"result": jp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": kvc.JsonModel() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": kvc.JsonModelAny() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// newCreator factories
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_128_NewCreator_Collection_String(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.String.Empty()

	// Act
	actual := args.Map{"result": c.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	c2 := coredynamic.New.Collection.String.Cap(5)
	actual := args.Map{"result": c2.Capacity() < 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cap >= 5", actual)
	c3 := coredynamic.New.Collection.String.From([]string{"a", "b"})
	actual := args.Map{"result": c3.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	c4 := coredynamic.New.Collection.String.Clone([]string{"x"})
	actual := args.Map{"result": c4.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c5 := coredynamic.New.Collection.String.Items("a", "b", "c")
	actual := args.Map{"result": c5.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	c6 := coredynamic.New.Collection.String.Create([]string{"x"})
	actual := args.Map{"result": c6.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c7 := coredynamic.New.Collection.String.LenCap(3, 10)
	actual := args.Map{"result": c7.Length() != 3 || c7.Capacity() < 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected len=3, cap>=10", actual)
}

func Test_C32_129_NewCreator_Collection_Int(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int.Empty()
	c.Add(1)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c2 := coredynamic.New.Collection.Int.LenCap(2, 5)
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_130_NewCreator_Collection_Int64(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Int64.Cap(5)
	c.Add(int64(1))

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c2 := coredynamic.New.Collection.Int64.LenCap(2, 5)
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_131_NewCreator_Collection_Byte(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Byte.Cap(5)
	c.Add(byte(1))

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c2 := coredynamic.New.Collection.Byte.LenCap(2, 5)
	actual := args.Map{"result": c2.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_132_NewCreator_Collection_Any(t *testing.T) {
	// Arrange
	c := coredynamic.New.Collection.Any.Empty()
	c.Add(42)

	// Act
	actual := args.Map{"result": c.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	c2 := coredynamic.New.Collection.Any.Items("a", 1, true)
	actual := args.Map{"result": c2.Length() != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C32_133_NewCreator_Collection_Others(t *testing.T) {
	// Arrange
	bs := coredynamic.New.Collection.ByteSlice.Empty()
	bs.Add([]byte{1})

	// Act
	actual := args.Map{"result": bs.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	b := coredynamic.New.Collection.Bool.Cap(5)
	b.Add(true)
	actual := args.Map{"result": b.Length() != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	f32 := coredynamic.New.Collection.Float32.Empty()
	f32.Add(1.0)
	f64 := coredynamic.New.Collection.Float64.Empty()
	f64.Add(2.0)
	am := coredynamic.New.Collection.AnyMap.Empty()
	am.Add(map[string]any{"a": 1})
	sm := coredynamic.New.Collection.StringMap.Empty()
	sm.Add(map[string]string{"a": "b"})
	im := coredynamic.New.Collection.IntMap.Empty()
	im.Add(map[string]int{"a": 1})
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection — additional methods
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_134_DynCol_AddAnyItemsWithTypeValidation(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	strType := reflect.TypeOf("")
	err := dc.AddAnyItemsWithTypeValidation(false, false, strType, "a", "b")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	err = dc.AddAnyItemsWithTypeValidation(false, false, strType, 42)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err = dc.AddAnyItemsWithTypeValidation(true, false, strType, "ok", 42)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err = dc.AddAnyItemsWithTypeValidation(false, false, strType)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_135_DynCol_AddAnySliceFromSingleItem(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAnySliceFromSingleItem(true, []string{"a", "b"})

	// Act
	actual := args.Map{"result": dc.Length() < 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected items", actual)
	dc.AddAnySliceFromSingleItem(true, nil)
}

func Test_C32_136_DynCol_JsonResultsCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	rc := dc.JsonResultsCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"result": empty.JsonResultsCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_137_DynCol_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)
	rc := dc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"result": rc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	empty := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"result": empty.JsonResultsPtrCollection().Length() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_C32_138_DynCol_JsonString(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(42, true)
	s, err := dc.JsonString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "json string failed", actual)
}

func Test_C32_139_DynCol_JsonStringMust(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	s := dc.JsonStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_C32_140_DynCol_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("hello", true)
	b, err := dc.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "marshal failed", actual)
}

func Test_C32_141_DynCol_SafeLimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	slc := dc.SafeLimitCollection(100)

	// Act
	actual := args.Map{"result": slc.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C32_142_DynCol_At_First_Last_Accessors(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny("a", true).AddAny("b", true).AddAny("c", true)
	atVal := dc.At(1)

	// Act
	actual := args.Map{"result": atVal.ValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	f := dc.First()
	actual := args.Map{"result": f.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	l := dc.Last()
	actual := args.Map{"result": l.ValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	fd := dc.FirstDynamic()
	actual := args.Map{"result": fd == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	ld := dc.LastDynamic()
	actual := args.Map{"result": ld == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_143_DynCol_FirstOrDefault_LastOrDefault(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("x", true)

	// Act
	actual := args.Map{"result": dc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": dc.LastOrDefault() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": dc.FirstOrDefaultDynamic() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": dc.LastOrDefaultDynamic() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	empty := coredynamic.EmptyDynamicCollection()
	actual := args.Map{"result": empty.FirstOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual := args.Map{"result": empty.LastOrDefault() != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_C32_144_DynCol_Skip_Take_Limit(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true).AddAny(4, true)

	// Act
	actual := args.Map{"result": len(dc.Skip(2)) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(dc.Take(2)) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual := args.Map{"result": len(dc.Limit(3)) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": dc.SkipDynamic(1) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": dc.TakeDynamic(2) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual := args.Map{"result": dc.LimitDynamic(2) == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_C32_145_DynCol_SkipCollection_TakeCollection_LimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(4)
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	sc := dc.SkipCollection(1)

	// Act
	actual := args.Map{"result": sc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	tc := dc.TakeCollection(2)
	actual := args.Map{"result": tc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	lc := dc.LimitCollection(2)
	actual := args.Map{"result": lc.Length() != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_C32_146_DynCol_GetPagedCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(10)
	for i := 0; i < 10; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagedCollection(3)

	// Act
	actual := args.Map{"result": len(pages) != 4}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 4 pages", actual)
}

func Test_C32_147_DynCol_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	pages := dc.GetPagedCollection(5)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_C32_148_DynCol_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny(1, true)
	page := dc.GetSinglePageCollection(5, 1)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

// ═══════════════════════════════════════════════════════════════════════
// DynamicCollection AddWithWgLock equivalent not exists;
// Additional AddAny with sync test
// ═══════════════════════════════════════════════════════════════════════

func Test_C32_149_Collection_AddWithWgLock_Proper(t *testing.T) {
	// Arrange
	c := coredynamic.NewCollection[int](10)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go c.AddWithWgLock(wg, 1)
	go c.AddWithWgLock(wg, 2)
	go c.AddWithWgLock(wg, 3)
	wg.Wait()

	// Act
	actual := args.Map{"result": c.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_C32_150_MapAnyItems_DiffRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	m.Add("b", 2)
	diff := m.DiffRaw(false, map[string]any{"a": 1, "b": 3})
	_ = diff
}

func Test_C32_151_MapAnyItems_Diff(t *testing.T) {
	m1 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 2})
	m2 := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1, "b": 3})
	diff := m1.Diff(false, m2)
	_ = diff
}

func Test_C32_152_MapAnyItems_HashmapDiffUsingRaw(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	diff := m.HashmapDiffUsingRaw(false, map[string]any{"a": 1})
	_ = diff
	diff2 := m.HashmapDiffUsingRaw(false, map[string]any{"a": 2})
	_ = diff2
}

func Test_C32_153_MapAnyItems_DiffJsonMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	msg := m.DiffJsonMessage(false, map[string]any{"a": 2})
	_ = msg
}

func Test_C32_154_MapAnyItems_ToStringsSliceOfDiffMap(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	strs := m.ToStringsSliceOfDiffMap(map[string]any{"a": 2})
	_ = strs
}

func Test_C32_155_MapAnyItems_ShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	msg := m.ShouldDiffMessage(false, "test", map[string]any{"a": 2})
	_ = msg
}

func Test_C32_156_MapAnyItems_LogShouldDiffMessage(t *testing.T) {
	m := coredynamic.NewMapAnyItems(4)
	m.Add("a", 1)
	msg := m.LogShouldDiffMessage(false, "test", map[string]any{"a": 2})
	_ = msg
}
