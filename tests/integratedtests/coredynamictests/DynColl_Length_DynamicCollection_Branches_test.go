package coredynamictests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// DynamicCollection — nil/empty branches
// =============================================================================

func Test_DynColl_Length_Nil(t *testing.T) {
	// Arrange
	var dc *coredynamic.DynamicCollection

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Length nil", actual)
}

func Test_DynColl_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var dc *coredynamic.DynamicCollection

	// Act
	actual := args.Map{"r": dc.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection IsEmpty nil", actual)
}

func Test_DynColl_IsEmpty_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"r": dc.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection IsEmpty empty", actual)
}

func Test_DynColl_HasAnyItem_False(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"r": dc.HasAnyItem()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasAnyItem false", actual)
}

func Test_DynColl_HasAnyItem_True(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("x", true)

	// Act
	actual := args.Map{"r": dc.HasAnyItem()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasAnyItem true", actual)
}

func Test_DynColl_Count(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)

	// Act
	actual := args.Map{"r": dc.Count()}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Count", actual)
}

func Test_DynColl_LastIndex(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)

	// Act
	actual := args.Map{"r": dc.LastIndex()}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastIndex", actual)
}

func Test_DynColl_HasIndex_True(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)

	// Act
	actual := args.Map{"r": dc.HasIndex(0)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasIndex true", actual)
}

func Test_DynColl_HasIndex_False(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"r": dc.HasIndex(0)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection HasIndex false", actual)
}

// =============================================================================
// DynamicCollection — Items branches
// =============================================================================

func Test_DynColl_Items_Nil(t *testing.T) {
	// Arrange
	var dc *coredynamic.DynamicCollection

	// Act
	actual := args.Map{"len": len(dc.Items())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Items nil", actual)
}

func Test_DynColl_Items_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)

	// Act
	actual := args.Map{"len": len(dc.Items())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Items valid", actual)
}

// =============================================================================
// DynamicCollection — First / Last / OrDefault
// =============================================================================

func Test_DynColl_FirstOrDefault_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"isNil": dc.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection FirstOrDefault empty", actual)
}

func Test_DynColl_FirstOrDefault_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("first", true)

	// Act
	actual := args.Map{"notNil": dc.FirstOrDefault() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection FirstOrDefault valid", actual)
}

func Test_DynColl_FirstOrDefaultDynamic_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"isNil": dc.FirstOrDefaultDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection FirstOrDefaultDynamic empty", actual)
}

func Test_DynColl_LastOrDefault_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"isNil": dc.LastOrDefault() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastOrDefault empty", actual)
}

func Test_DynColl_LastOrDefault_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true).AddAny("last", true)

	// Act
	actual := args.Map{"notNil": dc.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastOrDefault valid", actual)
}

func Test_DynColl_LastOrDefaultDynamic_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"isNil": dc.LastOrDefaultDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LastOrDefaultDynamic empty", actual)
}

// =============================================================================
// DynamicCollection — Skip / Take / Limit
// =============================================================================

func Test_DynColl_Skip(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)

	// Act
	actual := args.Map{"len": len(dc.Skip(1))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Skip", actual)
}

func Test_DynColl_SkipCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	sc := dc.SkipCollection(2)

	// Act
	actual := args.Map{"len": sc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection SkipCollection", actual)
}

func Test_DynColl_Take(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)

	// Act
	actual := args.Map{"len": len(dc.Take(2))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Take", actual)
}

func Test_DynColl_TakeCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	tc := dc.TakeCollection(2)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection TakeCollection", actual)
}

func Test_DynColl_LimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	lc := dc.LimitCollection(1)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LimitCollection", actual)
}

func Test_DynColl_SafeLimitCollection(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	lc := dc.SafeLimitCollection(100)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection SafeLimitCollection", actual)
}

func Test_DynColl_LimitDynamic(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	r := dc.LimitDynamic(1)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection LimitDynamic", actual)
}

func Test_DynColl_Limit(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)

	// Act
	actual := args.Map{"len": len(dc.Limit(1))}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Limit", actual)
}

func Test_DynColl_SkipDynamic(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	r := dc.SkipDynamic(1)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection SkipDynamic", actual)
}

func Test_DynColl_TakeDynamic(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	r := dc.TakeDynamic(1)

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection TakeDynamic", actual)
}

// =============================================================================
// DynamicCollection — RemoveAt
// =============================================================================

func Test_DynColl_RemoveAt_Invalid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"r": dc.RemoveAt(0)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection RemoveAt invalid", actual)
}

func Test_DynColl_RemoveAt_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	ok := dc.RemoveAt(1)

	// Act
	actual := args.Map{
		"ok": ok,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection RemoveAt valid", actual)
}

// =============================================================================
// DynamicCollection — Loop
// =============================================================================

func Test_DynColl_Loop_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	called := false
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool { called = true; return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Loop empty", actual)
}

func Test_DynColl_Loop_Break(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return true
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Loop break", actual)
}

func Test_DynColl_Loop_All(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true)
	count := 0
	dc.Loop(func(i int, d *coredynamic.Dynamic) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Loop all", actual)
}

// =============================================================================
// DynamicCollection — Add variants
// =============================================================================

func Test_DynColl_AddAnyNonNull_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull(nil, false)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyNonNull nil", actual)
}

func Test_DynColl_AddAnyNonNull_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyNonNull("a", true)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyNonNull valid", actual)
}

func Test_DynColl_AddAnyMany_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany(nil...)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyMany nil", actual)
}

func Test_DynColl_AddAnyMany_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnyMany("a", "b")

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyMany valid", actual)
}

func Test_DynColl_Add(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamic("x", true)
	dc.Add(d)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Add", actual)
}

func Test_DynColl_AddPtr_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddPtr(nil)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddPtr nil", actual)
}

func Test_DynColl_AddPtr_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	d := coredynamic.NewDynamicPtr("x", true)
	dc.AddPtr(d)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddPtr valid", actual)
}

func Test_DynColl_AddManyPtr_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddManyPtr(nil...)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddManyPtr nil", actual)
}

func Test_DynColl_AddManyPtr_WithNils(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	d1 := coredynamic.NewDynamicPtr("a", true)
	dc.AddManyPtr(d1, nil)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddManyPtr with nils", actual)
}

// =============================================================================
// DynamicCollection — AnyItems / AnyItemsCollection / AddAnySliceFromSingleItem
// =============================================================================

func Test_DynColl_AnyItems_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"len": len(dc.AnyItems())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItems empty", actual)
}

func Test_DynColl_AnyItems_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)

	// Act
	actual := args.Map{"len": len(dc.AnyItems())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItems valid", actual)
}

func Test_DynColl_AnyItemsCollection_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	ac := dc.AnyItemsCollection()

	// Act
	actual := args.Map{"empty": ac.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItemsCollection empty", actual)
}

func Test_DynColl_AnyItemsCollection_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	ac := dc.AnyItemsCollection()

	// Act
	actual := args.Map{"len": ac.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AnyItemsCollection valid", actual)
}

func Test_DynColl_AddAnySliceFromSingleItem_Nil(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnySliceFromSingleItem(true, nil)

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnySliceFromSingleItem nil", actual)
}

func Test_DynColl_AddAnySliceFromSingleItem_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAnySliceFromSingleItem(true, []int{1, 2, 3})

	// Act
	actual := args.Map{"len": dc.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnySliceFromSingleItem valid", actual)
}

// =============================================================================
// DynamicCollection — Type validation
// =============================================================================

func Test_DynColl_AddAnyWithTypeValidation_Error(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(true, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyWithTypeValidation error", actual)
}

func Test_DynColl_AddAnyWithTypeValidation_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyWithTypeValidation(true, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyWithTypeValidation valid", actual)
}

func Test_DynColl_AddAnyItemsWithTypeValidation_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, true, reflect.TypeOf(""))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyItemsWithTypeValidation empty", actual)
}

func Test_DynColl_AddAnyItemsWithTypeValidation_ContinueOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(true, true, reflect.TypeOf(""), "a", 42, "b")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyItemsWithTypeValidation continue on error", actual)
}

func Test_DynColl_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.AddAnyItemsWithTypeValidation(false, true, reflect.TypeOf(""), "a", 42, "b")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": dc.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection AddAnyItemsWithTypeValidation stop on error", actual)
}

// =============================================================================
// DynamicCollection — JSON
// =============================================================================

func Test_DynColl_JsonString_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	s, err := dc.JsonString()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nonEmpty": len(s) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonString valid", actual)
}

func Test_DynColl_JsonStringMust_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	s := dc.JsonStringMust()

	// Act
	actual := args.Map{"nonEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonStringMust valid", actual)
}

func Test_DynColl_MarshalJSON(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	b, err := dc.MarshalJSON()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"nonEmpty": len(b) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicCollection MarshalJSON", actual)
}

func Test_DynColl_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	err := dc.UnmarshalJSON([]byte(`not json`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection UnmarshalJSON invalid", actual)
}

func Test_DynColl_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := dc.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ParseInjectUsingJson error", actual)
}

func Test_DynColl_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		dc.ParseInjectUsingJsonMust(jr)
	}()

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ParseInjectUsingJsonMust panics", actual)
}

func Test_DynColl_JsonParseSelfInject_Error(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	err := dc.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonParseSelfInject error", actual)
}

func Test_DynColl_Json(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Json", actual)
}

func Test_DynColl_JsonPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.JsonPtr()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonPtr", actual)
}

// =============================================================================
// DynamicCollection — JsonResultsCollection / JsonResultsPtrCollection
// =============================================================================

func Test_DynColl_JsonResultsCollection_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	r := dc.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsCollection empty", actual)
}

func Test_DynColl_JsonResultsCollection_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsCollection valid", actual)
}

func Test_DynColl_JsonResultsPtrCollection_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	r := dc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsPtrCollection empty", actual)
}

func Test_DynColl_JsonResultsPtrCollection_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonResultsPtrCollection valid", actual)
}

// =============================================================================
// DynamicCollection — Paging
// =============================================================================

func Test_DynColl_GetPagesSize_Zero(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"r": dc.GetPagesSize(0)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagesSize zero", actual)
}

func Test_DynColl_GetPagesSize_Negative(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"r": dc.GetPagesSize(-1)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagesSize negative", actual)
}

func Test_DynColl_GetPagesSize_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true).AddAny(2, true).AddAny(3, true)

	// Act
	actual := args.Map{"r": dc.GetPagesSize(2)}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagesSize valid", actual)
}

func Test_DynColl_GetPagedCollection_SmallData(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	pages := dc.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagedCollection small", actual)
}

func Test_DynColl_GetPagedCollection_MultiPage(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	for i := 0; i < 5; i++ {
		dc.AddAny(i, true)
	}
	pages := dc.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetPagedCollection multi", actual)
}

func Test_DynColl_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	r := dc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection GetSinglePageCollection small", actual)
}

// =============================================================================
// DynamicCollection — Misc
// =============================================================================

func Test_DynColl_Strings_Empty(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"len": len(dc.Strings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Strings empty", actual)
}

func Test_DynColl_Strings_Valid(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)

	// Act
	actual := args.Map{"len": len(dc.Strings())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "DynamicCollection Strings valid", actual)
}

func Test_DynColl_String(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("a", true)

	// Act
	actual := args.Map{"nonEmpty": len(dc.String()) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection String", actual)
}

func Test_DynColl_JsonModel(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny(1, true)
	m := dc.JsonModel()

	// Act
	actual := args.Map{"hasItems": len(m.Items) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonModel", actual)
}

func Test_DynColl_JsonModelAny(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()

	// Act
	actual := args.Map{"notNil": dc.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection JsonModelAny", actual)
}

func Test_DynColl_ListStrings(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	r := dc.ListStrings()

	// Act
	actual := args.Map{"nonEmpty": len(r) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ListStrings", actual)
}

func Test_DynColl_ListStringsPtr(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("hello", true)
	r := dc.ListStringsPtr()

	// Act
	actual := args.Map{"nonEmpty": len(r) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection ListStringsPtr", actual)
}

func Test_DynColl_At(t *testing.T) {
	// Arrange
	dc := coredynamic.EmptyDynamicCollection()
	dc.AddAny("x", true)
	d := dc.At(0)

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "DynamicCollection At", actual)
}
