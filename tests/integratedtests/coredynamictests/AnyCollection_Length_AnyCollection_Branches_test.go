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
// AnyCollection — nil/empty branches
// =============================================================================

func Test_AnyCollection_Length_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.AnyCollection

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection Length nil", actual)
}

func Test_AnyCollection_IsEmpty_Nil(t *testing.T) {
	// Arrange
	var c *coredynamic.AnyCollection

	// Act
	actual := args.Map{"r": c.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection IsEmpty nil", actual)
}

func Test_AnyCollection_IsEmpty_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"r": c.IsEmpty()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection IsEmpty empty", actual)
}

func Test_AnyCollection_HasAnyItem_False(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"r": c.HasAnyItem()}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection HasAnyItem false", actual)
}

func Test_AnyCollection_HasAnyItem_True(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)

	// Act
	actual := args.Map{"r": c.HasAnyItem()}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection HasAnyItem true", actual)
}

func Test_AnyCollection_LastIndex_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("a").Add("b")

	// Act
	actual := args.Map{"r": c.LastIndex()}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection LastIndex", actual)
}

func Test_AnyCollection_HasIndex_True(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("a")

	// Act
	actual := args.Map{"r": c.HasIndex(0)}

	// Assert
	expected := args.Map{"r": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection HasIndex true", actual)
}

func Test_AnyCollection_HasIndex_False(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"r": c.HasIndex(0)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection HasIndex false", actual)
}

func Test_AnyCollection_Count(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("a")

	// Act
	actual := args.Map{"r": c.Count()}

	// Assert
	expected := args.Map{"r": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection Count", actual)
}

// =============================================================================
// AnyCollection — Items / DynamicItems / DynamicCollection
// =============================================================================

func Test_AnyCollection_Items_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"len": len(c.Items())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection Items empty", actual)
}

func Test_AnyCollection_Items_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2)

	// Act
	actual := args.Map{"len": len(c.Items())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection Items valid", actual)
}

func Test_AnyCollection_DynamicItems_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"len": len(c.DynamicItems())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection DynamicItems empty", actual)
}

func Test_AnyCollection_DynamicItems_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)

	// Act
	actual := args.Map{"len": len(c.DynamicItems())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection DynamicItems valid", actual)
}

func Test_AnyCollection_DynamicCollection_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	dc := c.DynamicCollection()

	// Act
	actual := args.Map{"empty": dc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection DynamicCollection empty", actual)
}

func Test_AnyCollection_DynamicCollection_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	dc := c.DynamicCollection()

	// Act
	actual := args.Map{"empty": dc.IsEmpty()}

	// Assert
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection DynamicCollection valid", actual)
}

// =============================================================================
// AnyCollection — First / Last / FirstOrDefault / LastOrDefault
// =============================================================================

func Test_AnyCollection_FirstOrDefault_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"isNil": c.FirstOrDefault() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection FirstOrDefault empty", actual)
}

func Test_AnyCollection_FirstOrDefault_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("first")

	// Act
	actual := args.Map{"r": c.FirstOrDefault()}

	// Assert
	expected := args.Map{"r": "first"}
	expected.ShouldBeEqual(t, 0, "AnyCollection FirstOrDefault valid", actual)
}

func Test_AnyCollection_FirstOrDefaultDynamic_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"isNil": c.FirstOrDefaultDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection FirstOrDefaultDynamic empty", actual)
}

func Test_AnyCollection_LastOrDefault_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"isNil": c.LastOrDefault() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection LastOrDefault empty", actual)
}

func Test_AnyCollection_LastOrDefault_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("a").Add("last")

	// Act
	actual := args.Map{"r": c.LastOrDefault()}

	// Assert
	expected := args.Map{"r": "last"}
	expected.ShouldBeEqual(t, 0, "AnyCollection LastOrDefault valid", actual)
}

func Test_AnyCollection_LastOrDefaultDynamic_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"isNil": c.LastOrDefaultDynamic() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection LastOrDefaultDynamic empty", actual)
}

// =============================================================================
// AnyCollection — Skip / Take / Limit
// =============================================================================

func Test_AnyCollection_Skip_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"len": len(c.Skip(1))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection Skip", actual)
}

func Test_AnyCollection_SkipCollection_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)
	sc := c.SkipCollection(2)

	// Act
	actual := args.Map{"len": sc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection SkipCollection", actual)
}

func Test_AnyCollection_Take_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"len": len(c.Take(2))}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection Take", actual)
}

func Test_AnyCollection_TakeCollection_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)
	tc := c.TakeCollection(2)

	// Act
	actual := args.Map{"len": tc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection TakeCollection", actual)
}

func Test_AnyCollection_LimitCollection_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)
	lc := c.LimitCollection(1)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection LimitCollection", actual)
}

func Test_AnyCollection_SafeLimitCollection_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2)
	lc := c.SafeLimitCollection(100)

	// Act
	actual := args.Map{"len": lc.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection SafeLimitCollection", actual)
}

// =============================================================================
// AnyCollection — RemoveAt
// =============================================================================

func Test_AnyCollection_RemoveAt_Invalid_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"r": c.RemoveAt(0)}

	// Assert
	expected := args.Map{"r": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection RemoveAt invalid", actual)
}

func Test_AnyCollection_RemoveAt_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)
	ok := c.RemoveAt(1)

	// Act
	actual := args.Map{
		"ok": ok,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection RemoveAt valid", actual)
}

// =============================================================================
// AnyCollection — Loop (sync and async)
// =============================================================================

func Test_AnyCollection_Loop_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	called := false
	c.Loop(false, func(i int, item any) bool { called = true; return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection Loop empty", actual)
}

func Test_AnyCollection_Loop_Sync_Break(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)
	count := 0
	c.Loop(false, func(i int, item any) bool {
		count++
		return i == 0
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection Loop sync break", actual)
}

func Test_AnyCollection_Loop_Async_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)
	c.Loop(true, func(i int, item any) bool { return false })

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection Loop async", actual)
}

func Test_AnyCollection_LoopDynamic_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	called := false
	c.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool { called = true; return false })

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "AnyCollection LoopDynamic empty", actual)
}

func Test_AnyCollection_LoopDynamic_Sync_Break(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2)
	count := 0
	c.LoopDynamic(false, func(i int, item coredynamic.Dynamic) bool {
		count++
		return true
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection LoopDynamic sync break", actual)
}

func Test_AnyCollection_LoopDynamic_Async_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2)
	c.LoopDynamic(true, func(i int, item coredynamic.Dynamic) bool { return false })

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection LoopDynamic async", actual)
}

// =============================================================================
// AnyCollection — Add variants
// =============================================================================

func Test_AnyCollection_AddNonNull_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddNonNull(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddNonNull nil", actual)
}

func Test_AnyCollection_AddNonNull_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddNonNull("a")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddNonNull valid", actual)
}

func Test_AnyCollection_AddNonNullDynamic_Nil(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddNonNullDynamic(nil, false)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddNonNullDynamic nil", actual)
}

func Test_AnyCollection_AddNonNullDynamic_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddNonNullDynamic("a", true)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddNonNullDynamic valid", actual)
}

func Test_AnyCollection_AddAnyManyDynamic_Nil_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddAnyManyDynamic(nil...)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyManyDynamic nil", actual)
}

func Test_AnyCollection_AddAnyManyDynamic_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddAnyManyDynamic("a", "b")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyManyDynamic valid", actual)
}

func Test_AnyCollection_AddMany_Nil_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddMany(nil...)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddMany nil", actual)
}

func Test_AnyCollection_AddMany_WithNils(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddMany("a", nil, "b")

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddMany with nils", actual)
}

func Test_AnyCollection_AddAnySliceFromSingleItem_Nil_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddAnySliceFromSingleItem(nil)

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnySliceFromSingleItem nil", actual)
}

func Test_AnyCollection_AddAnySliceFromSingleItem_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.AddAnySliceFromSingleItem([]int{1, 2, 3})

	// Act
	actual := args.Map{"len": c.Length()}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnySliceFromSingleItem valid", actual)
}

// =============================================================================
// AnyCollection — Type validation
// =============================================================================

func Test_AnyCollection_AddAnyWithTypeValidation_Error(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.AddAnyWithTypeValidation(true, reflect.TypeOf(""), 42)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyWithTypeValidation error", actual)
}

func Test_AnyCollection_AddAnyWithTypeValidation_Valid_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.AddAnyWithTypeValidation(true, reflect.TypeOf(""), "hello")

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyWithTypeValidation valid", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.AddAnyItemsWithTypeValidation(false, true, reflect.TypeOf(""))

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyItemsWithTypeValidation empty", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_ContinueOnError_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.AddAnyItemsWithTypeValidation(true, true, reflect.TypeOf(""), "a", 42, "b")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyItemsWithTypeValidation continue on error", actual)
}

func Test_AnyCollection_AddAnyItemsWithTypeValidation_StopOnError(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.AddAnyItemsWithTypeValidation(false, true, reflect.TypeOf(""), "a", 42, "b")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"hasErr": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection AddAnyItemsWithTypeValidation stop on error", actual)
}

// =============================================================================
// AnyCollection — JSON branches
// =============================================================================

func Test_AnyCollection_JsonString_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	s, err := c.JsonString()

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
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonString valid", actual)
}

func Test_AnyCollection_JsonStringMust_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	s := c.JsonStringMust()

	// Act
	actual := args.Map{"nonEmpty": len(s) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonStringMust valid", actual)
}

func Test_AnyCollection_MarshalJSON_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	b, err := c.MarshalJSON()

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
	expected.ShouldBeEqual(t, 0, "AnyCollection MarshalJSON", actual)
}

func Test_AnyCollection_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.UnmarshalJSON([]byte(`[1,2,3]`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": c.Length(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "AnyCollection UnmarshalJSON valid", actual)
}

func Test_AnyCollection_UnmarshalJSON_Invalid_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	err := c.UnmarshalJSON([]byte(`not json`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection UnmarshalJSON invalid", actual)
}

func Test_AnyCollection_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	_, err := c.ParseInjectUsingJson(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection ParseInjectUsingJson error", actual)
}

func Test_AnyCollection_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		c.ParseInjectUsingJsonMust(jr)
	}()

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection ParseInjectUsingJsonMust panics", actual)
}

func Test_AnyCollection_JsonParseSelfInject_Error(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	jr := &corejson.Result{Error: errors.New("fail")}
	err := c.JsonParseSelfInject(jr)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonParseSelfInject error", actual)
}

func Test_AnyCollection_Json_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	r := c.Json()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection Json", actual)
}

func Test_AnyCollection_JsonPtr_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	r := c.JsonPtr()

	// Act
	actual := args.Map{"noErr": !r.HasError()}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonPtr", actual)
}

// =============================================================================
// AnyCollection — JsonResultsCollection / JsonResultsPtrCollection
// =============================================================================

func Test_AnyCollection_JsonResultsCollection_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	r := c.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonResultsCollection empty", actual)
}

func Test_AnyCollection_JsonResultsCollection_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	r := c.JsonResultsCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonResultsCollection valid", actual)
}

func Test_AnyCollection_JsonResultsPtrCollection_Empty(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	r := c.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonResultsPtrCollection empty", actual)
}

func Test_AnyCollection_JsonResultsPtrCollection_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	r := c.JsonResultsPtrCollection()

	// Act
	actual := args.Map{"notNil": r != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonResultsPtrCollection valid", actual)
}

// =============================================================================
// AnyCollection — Paging
// =============================================================================

func Test_AnyCollection_GetPagesSize_Zero_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"r": c.GetPagesSize(0)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection GetPagesSize zero", actual)
}

func Test_AnyCollection_GetPagesSize_Negative(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"r": c.GetPagesSize(-1)}

	// Assert
	expected := args.Map{"r": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection GetPagesSize negative", actual)
}

func Test_AnyCollection_GetPagesSize_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1).Add(2).Add(3)

	// Act
	actual := args.Map{"r": c.GetPagesSize(2)}

	// Assert
	expected := args.Map{"r": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection GetPagesSize valid", actual)
}

func Test_AnyCollection_GetPagedCollection_SmallData(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	pages := c.GetPagedCollection(10)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection GetPagedCollection small data", actual)
}

func Test_AnyCollection_GetPagedCollection_MultiPage(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	for i := 0; i < 5; i++ {
		c.Add(i)
	}
	pages := c.GetPagedCollection(2)

	// Act
	actual := args.Map{"pages": len(pages)}

	// Assert
	expected := args.Map{"pages": 3}
	expected.ShouldBeEqual(t, 0, "AnyCollection GetPagedCollection multi page", actual)
}

func Test_AnyCollection_GetSinglePageCollection_Small_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)
	r := c.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"len": r.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection GetSinglePageCollection small", actual)
}

// =============================================================================
// AnyCollection — Misc
// =============================================================================

func Test_AnyCollection_Strings_Empty_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"len": len(c.Strings())}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyCollection Strings empty", actual)
}

func Test_AnyCollection_Strings_Valid(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("a").Add(1)

	// Act
	actual := args.Map{"len": len(c.Strings())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AnyCollection Strings valid", actual)
}

func Test_AnyCollection_String_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("a")

	// Act
	actual := args.Map{"nonEmpty": len(c.String()) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection String", actual)
}

func Test_AnyCollection_JsonModel_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(1)

	// Act
	actual := args.Map{"notNil": c.JsonModel() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonModel", actual)
}

func Test_AnyCollection_JsonModelAny_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()

	// Act
	actual := args.Map{"notNil": c.JsonModelAny() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection JsonModelAny", actual)
}

func Test_AnyCollection_ListStrings_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("hello")
	r := c.ListStrings(false)

	// Act
	actual := args.Map{"len": len(r)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyCollection ListStrings", actual)
}

func Test_AnyCollection_ListStringsPtr_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("hello")
	r := c.ListStringsPtr(true)

	// Act
	actual := args.Map{"nonEmpty": len(r) > 0}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection ListStringsPtr", actual)
}

func Test_AnyCollection_At_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add("x")

	// Act
	actual := args.Map{"r": c.At(0)}

	// Assert
	expected := args.Map{"r": "x"}
	expected.ShouldBeEqual(t, 0, "AnyCollection At", actual)
}

func Test_AnyCollection_AtAsDynamic_FromAnyCollectionLengthA(t *testing.T) {
	// Arrange
	c := coredynamic.EmptyAnyCollection()
	c.Add(42)
	d := c.AtAsDynamic(0)

	// Act
	actual := args.Map{"valid": d.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "AnyCollection AtAsDynamic", actual)
}
