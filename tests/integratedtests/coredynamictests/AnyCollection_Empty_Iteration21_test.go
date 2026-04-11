package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — constructors, accessors, navigation
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyCollection_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{
		"empty": ac.IsEmpty(),
		"len": ac.Length(),
		"count": ac.Count(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"len": 0,
		"count": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Empty", actual)
}

func Test_AnyCollection_New(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(5)

	// Act
	actual := args.Map{"empty": ac.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- New", actual)
}

func Test_AnyCollection_Add(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{
		"len": ac.Length(),
		"hasAny": ac.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Add", actual)
}

func Test_AnyCollection_AddAny(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("x", true)

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAny", actual)
}

func Test_AnyCollection_AddMany(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddMany("a", "b", nil, "c")

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 3} // nil skipped
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddMany", actual)
}

func Test_AnyCollection_AddMany_Nil(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddMany()

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddMany nil", actual)
}

func Test_AnyCollection_AddNonNull(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddNonNull("a").AddNonNull(nil)

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddNonNull", actual)
}

func Test_AnyCollection_AddNonNullDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddNonNullDynamic("a", true).AddNonNullDynamic(nil, false)

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddNonNullDynamic", actual)
}

func Test_AnyCollection_AddAnyManyDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnyManyDynamic("a", "b")

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnyManyDynamic", actual)
}

func Test_AnyCollection_AddAnyManyDynamic_Nil(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnyManyDynamic()

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddAnyManyDynamic nil", actual)
}

func Test_AnyCollection_At(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{"val": ac.At(1)}

	// Assert
	expected := args.Map{"val": "b"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- At", actual)
}

func Test_AnyCollection_AtAsDynamic(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true)
	d := ac.AtAsDynamic(0)

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AtAsDynamic", actual)
}

func Test_AnyCollection_Items(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{"len": len(ac.Items())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Items", actual)
}

func Test_AnyCollection_Items_Nil(t *testing.T) {
	// Arrange
	var ac *coredynamic.AnyCollection
	items := ac.Items()

	// Act
	actual := args.Map{
		"nil": items == nil,
		"len": len(items),
	}

	// Assert
	expected := args.Map{
		"nil": false,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- Items nil", actual)
}

func Test_AnyCollection_DynamicItems(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	di := ac.DynamicItems()

	// Act
	actual := args.Map{"len": len(di)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- DynamicItems", actual)
}

func Test_AnyCollection_DynamicCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true)
	dc := ac.DynamicCollection()

	// Act
	actual := args.Map{"notNil": dc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- DynamicCollection", actual)
}

func Test_AnyCollection_First(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{
		"first": ac.First(),
		"last": ac.Last(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "b",
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- First/Last", actual)
}

func Test_AnyCollection_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"nil": ac.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- FirstOrDefault empty", actual)
}

func Test_AnyCollection_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"nil": ac.LastOrDefault() == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- LastOrDefault empty", actual)
}

func Test_AnyCollection_FirstOrDefault_HasItem(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("x")

	// Act
	actual := args.Map{"val": ac.FirstOrDefault()}

	// Assert
	expected := args.Map{"val": "x"}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- FirstOrDefault has item", actual)
}

func Test_AnyCollection_Skip(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	skipped := ac.Skip(1)

	// Act
	actual := args.Map{"len": len(skipped)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Skip", actual)
}

func Test_AnyCollection_SkipCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	sc := ac.SkipCollection(2)

	// Act
	actual := args.Map{"len": sc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- SkipCollection", actual)
}

func Test_AnyCollection_Take(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	taken := ac.Take(2)

	// Act
	actual := args.Map{"len": len(taken)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Take", actual)
}

func Test_AnyCollection_TakeCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	tc := ac.TakeCollection(2)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- TakeCollection", actual)
}

func Test_AnyCollection_LimitCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	lc := ac.LimitCollection(2)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LimitCollection", actual)
}

func Test_AnyCollection_SafeLimitCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	lc := ac.SafeLimitCollection(10)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- SafeLimitCollection", actual)
}

func Test_AnyCollection_SafeLimitCollection_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	lc := ac.SafeLimitCollection(10)

	// Act
	actual := args.Map{"empty": lc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- SafeLimitCollection empty", actual)
}

func Test_AnyCollection_LastIndex(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")

	// Act
	actual := args.Map{"idx": ac.LastIndex()}

	// Assert
	expected := args.Map{"idx": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LastIndex", actual)
}

func Test_AnyCollection_HasIndex(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")

	// Act
	actual := args.Map{
		"has0": ac.HasIndex(0),
		"has1": ac.HasIndex(1),
	}

	// Assert
	expected := args.Map{
		"has0": true,
		"has1": false,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- HasIndex", actual)
}

func Test_AnyCollection_RemoveAt(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	ok := ac.RemoveAt(1)

	// Act
	actual := args.Map{
		"ok": ok,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- RemoveAt", actual)
}

func Test_AnyCollection_RemoveAt_Invalid(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ok := ac.RemoveAt(5)

	// Act
	actual := args.Map{"ok": ok}

	// Assert
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns error -- RemoveAt invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — Loop, LoopDynamic
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyCollection_Loop_Sync(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop sync", actual)
}

func Test_AnyCollection_Loop_Break(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return i == 0
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop break", actual)
}

func Test_AnyCollection_Loop_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	count := 0
	ac.Loop(false, func(i int, item any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Loop empty", actual)
}

func Test_AnyCollection_Loop_Async(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	ac.Loop(true, func(i int, item any) bool {
		return false
	})

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Loop async", actual)
}

func Test_AnyCollection_LoopDynamic_Sync(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic sync", actual)
}

func Test_AnyCollection_LoopDynamic_Break(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	count := 0
	ac.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return true
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic break", actual)
}

func Test_AnyCollection_LoopDynamic_Async(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("a", true).AddAny("b", true)
	ac.LoopDynamic(true, func(i int, item coredynamic.Dynamic) bool {
		return false
	})

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- LoopDynamic async", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — Type validation
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyCollection_AddAnyWithTypeValidation_Match(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyWithTypeValidation match", actual)
}

func Test_AnyCollection_AddAnyWithTypeValidation_Mismatch(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyWithTypeValidation(false, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyWithTypeValidation mismatch", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_StopOnErr(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyItemsWithTypeValidation(
		false, false,
		reflect.TypeOf(""),
		"a", 42, "c",
	)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation stop", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_ContinueOnErr(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyItemsWithTypeValidation(
		true, false,
		reflect.TypeOf(""),
		"a", 42, "c",
	)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns non-empty -- AddAnyItemsWithTypeValidation continue", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.AddAnyItemsWithTypeValidation(false, false, reflect.TypeOf(""))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- AddAnyItemsWithTypeValidation empty", actual)
}

func Test_AnyCollection_AddAnySliceFromSingleItem(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnySliceFromSingleItem([]string{"a", "b"})

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- AddAnySliceFromSingleItem", actual)
}

func Test_AnyCollection_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAnySliceFromSingleItem(nil)

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- AddAnySliceFromSingleItem nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — Paging
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyCollection_GetPagesSize(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b").Add("c")

	// Act
	actual := args.Map{"pages": ac.GetPagesSize(2)}

	// Assert
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagesSize", actual)
}

func Test_AnyCollection_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"pages": ac.GetPagesSize(0)}

	// Assert
	expected := args.Map{"pages": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagesSize zero", actual)
}

func Test_AnyCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	for i := 0; i < 5; i++ {
		ac.Add(i)
	}
	pages := ac.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagedCollection", actual)
}

func Test_AnyCollection_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	pages := ac.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagedCollection small", actual)
}

func Test_AnyCollection_GetSinglePageCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	page := ac.GetSinglePageCollection(3, 2)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetSinglePageCollection", actual)
}

func Test_AnyCollection_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	page := ac.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"len": page.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetSinglePageCollection small", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyCollection — JSON methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyCollection_JsonString(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	s, err := ac.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonString", actual)
}

func Test_AnyCollection_JsonStringMust(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	s := ac.JsonStringMust()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonStringMust", actual)
}

func Test_AnyCollection_MarshalJSON(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	b, err := ac.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasBytes": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- MarshalJSON", actual)
}

func Test_AnyCollection_UnmarshalJSON(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.UnmarshalJSON([]byte(`["a","b"]`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": ac.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- UnmarshalJSON", actual)
}

func Test_AnyCollection_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	err := ac.UnmarshalJSON([]byte(`not json`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns error -- UnmarshalJSON invalid", actual)
}

func Test_AnyCollection_JsonResultsCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	rc := ac.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonResultsCollection", actual)
}

func Test_AnyCollection_JsonResultsCollection_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	rc := ac.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- JsonResultsCollection empty", actual)
}

func Test_AnyCollection_JsonResultsPtrCollection(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	rc := ac.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": rc != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonResultsPtrCollection", actual)
}

func Test_AnyCollection_JsonModel(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	model := ac.JsonModel()

	// Act
	actual := args.Map{"len": len(model)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonModel", actual)
}

func Test_AnyCollection_JsonModelAny(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")

	// Act
	actual := args.Map{"notNil": ac.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonModelAny", actual)
}

func Test_AnyCollection_Json(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	jr := ac.Json()

	// Act
	actual := args.Map{"hasErr": jr.HasError()}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Json", actual)
}

func Test_AnyCollection_JsonPtr(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")
	jr := ac.JsonPtr()

	// Act
	actual := args.Map{"notNil": jr != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonPtr", actual)
}

func Test_AnyCollection_ParseInjectUsingJson(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	jr := corejson.NewPtr([]any{"a", "b"})
	result, err := ac.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ParseInjectUsingJson", actual)
}

func Test_AnyCollection_JsonParseSelfInject(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	jr := corejson.NewPtr([]any{"x"})
	err := ac.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- JsonParseSelfInject", actual)
}

func Test_AnyCollection_Strings(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a").Add("b")
	strs := ac.Strings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- Strings", actual)
}

func Test_AnyCollection_Strings_Empty(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	strs := ac.Strings()

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns empty -- Strings empty", actual)
}

func Test_AnyCollection_String(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("a")

	// Act
	actual := args.Map{"notEmpty": ac.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- String", actual)
}

func Test_AnyCollection_ListStringsPtr(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("hello", true)
	strs := ac.ListStringsPtr(false)

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ListStringsPtr", actual)
}

func Test_AnyCollection_ListStrings(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.AddAny("hello", true)
	strs := ac.ListStrings(true)

	// Act
	actual := args.Map{"len": len(strs)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ListStrings", actual)
}

func Test_AnyCollection_ReflectSetAt(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	ac.Add("hello")
	var target string
	err := ac.ReflectSetAt(0, &target)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": target,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- ReflectSetAt", actual)
}

func Test_AnyCollection_GetPagingInfo(t *testing.T) {
	// Arrange
	ac := coredynamic.EmptyAnyCollection()
	for i := 0; i < 10; i++ {
		ac.Add(i)
	}
	info := ac.GetPagingInfo(3, 2)

	// Act
	actual := args.Map{"hasSkip": info.SkipItems > 0}

	// Assert
	expected := args.Map{"hasSkip": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns correct value -- GetPagingInfo", actual)
}

func Test_AnyCollection_Nil_Length(t *testing.T) {
	// Arrange
	var ac *coredynamic.AnyCollection

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- nil Length", actual)
}

func Test_AnyCollection_Nil_IsEmpty(t *testing.T) {
	// Arrange
	var ac *coredynamic.AnyCollection

	// Act
	actual := args.Map{"empty": ac.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection returns nil -- nil IsEmpty", actual)
}
