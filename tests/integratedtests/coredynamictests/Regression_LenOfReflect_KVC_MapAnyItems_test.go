package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Regression: LengthOfReflect — pointer inputs
// Prevents: panic on reflect.Indirect of nil/zero/ptr values
// ══════════════════════════════════════════════════════════════════════════════

func Test_Reg_LengthOfReflect_PtrToSlice(t *testing.T) {
	// Arrange
	s := []int{1, 2, 3}
	rv := reflect.ValueOf(&s)

	// Act
	length := coredynamic.LengthOfReflect(rv)

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 3 -- pointer to slice", actual)
}

func Test_Reg_LengthOfReflect_PtrToMap(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1, "b": 2}
	rv := reflect.ValueOf(&m)

	// Act
	length := coredynamic.LengthOfReflect(rv)

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 2 -- pointer to map", actual)
}

func Test_Reg_LengthOfReflect_PtrToArray(t *testing.T) {
	// Arrange
	a := [4]int{1, 2, 3, 4}
	rv := reflect.ValueOf(&a)

	// Act
	length := coredynamic.LengthOfReflect(rv)

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 4}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 4 -- pointer to array", actual)
}

func Test_Reg_LengthOfReflect_NilPtr(t *testing.T) {
	// Arrange
	var s *[]int
	rv := reflect.ValueOf(s)

	// Act
	length := coredynamic.LengthOfReflect(rv)

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 0 -- nil pointer", actual)
}

func Test_Reg_LengthOfReflect_ZeroValue(t *testing.T) {
	// Arrange
	rv := reflect.Value{}

	// Act
	length := coredynamic.LengthOfReflect(rv)

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 0 -- zero reflect.Value", actual)
}

func Test_Reg_LengthOfReflect_NonCollection(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act
	length := coredynamic.LengthOfReflect(rv)

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "LengthOfReflect returns 0 -- non-collection int", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Regression: SafeZeroSet — zero/nil/non-ptr inputs must not panic
// ══════════════════════════════════════════════════════════════════════════════

func Test_Reg_SafeZeroSet_ZeroReflectValue(t *testing.T) {
	// Arrange
	rv := reflect.Value{}

	// Act — must not panic
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- zero reflect.Value", actual)
}

func Test_Reg_SafeZeroSet_NonPtr(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)

	// Act — must not panic
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- non-pointer int", actual)
}

func Test_Reg_SafeZeroSet_NilPtr(t *testing.T) {
	// Arrange
	var s *string
	rv := reflect.ValueOf(s)

	// Act — must not panic
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- nil pointer", actual)
}

func Test_Reg_SafeZeroSet_ValidPtr(t *testing.T) {
	// Arrange
	x := 42
	rv := reflect.ValueOf(&x)

	// Act
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"zeroed": x == 0}
	expected := args.Map{"zeroed": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet zeroes value -- valid pointer", actual)
}

func Test_Reg_SafeZeroSet_NilInterface(t *testing.T) {
	// Arrange
	var rt reflect.Type
	rv := reflect.ValueOf(rt)

	// Act — must not panic
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- nil interface", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Regression: MapAnyItems — NewUsingAnyTypeMap
// Prevents: wrong conversion logic or nil handling
// ══════════════════════════════════════════════════════════════════════════════

func Test_Reg_MapAnyItems_NewUsingAnyTypeMap_StringInt(t *testing.T) {
	// Arrange
	m := map[string]int{"a": 1, "b": 2}

	// Act
	items, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(m)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"len": items.Length(),
	}
	expected := args.Map{
		"noErr": true,
		"len": 2,
	}
	expected.ShouldBeEqual(t, 0, "NewUsingAnyTypeMap returns map -- map[string]int", actual)
}

func Test_Reg_MapAnyItems_NewUsingAnyTypeMap_Nil(t *testing.T) {
	// Arrange — nil input

	// Act
	items, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(nil)

	// Assert
	actual := args.Map{
		"hasErr": err != nil,
		"empty": items.IsEmpty(),
	}
	expected := args.Map{
		"hasErr": true,
		"empty": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUsingAnyTypeMap returns error -- nil input", actual)
}

func Test_Reg_MapAnyItems_NewUsingAnyTypeMap_StringAny(t *testing.T) {
	// Arrange
	m := map[string]any{"key": "val"}

	// Act
	items, err := coredynamic.NewMapAnyItemsUsingAnyTypeMap(m)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"hasKey": items.HasKey("key"),
	}
	expected := args.Map{
		"noErr": true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "NewUsingAnyTypeMap returns map -- map[string]any", actual)
}

func Test_Reg_MapAnyItems_GetFieldsMap_MapValue(t *testing.T) {
	// Arrange
	inner := map[string]any{"x": 1}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": inner})

	// Act
	fm, err, found := m.GetFieldsMap("k")

	// Assert
	actual := args.Map{
		"found": found,
		"noErr": err == nil,
		"notNil": fm != nil,
	}
	expected := args.Map{
		"found": true,
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetFieldsMap returns map -- map value", actual)
}

func Test_Reg_MapAnyItems_GetSafeFieldsMap_MapValue(t *testing.T) {
	// Arrange
	inner := map[string]any{"y": 2}
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"k": inner})

	// Act
	fm, found := m.GetSafeFieldsMap("k")

	// Assert
	actual := args.Map{
		"found": found,
		"notNil": fm != nil,
	}
	expected := args.Map{
		"found": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSafeFieldsMap returns map -- map value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Regression: KeyValCollection — JSON serialization round-trip
// Prevents: false expectations about unexported fields causing empty JSON
// ══════════════════════════════════════════════════════════════════════════════

func Test_Reg_KVC_Json_HasBytes(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})

	// Act
	j := kvc.Json()

	// Assert
	actual := args.Map{
		"hasBytes": j.HasBytes(),
		"noErr": !j.HasError(),
	}
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC Json returns bytes -- non-empty collection", actual)
}

func Test_Reg_KVC_JsonString_NonEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})

	// Act
	s, err := kvc.JsonString()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"notEmpty": s != "",
	}
	expected := args.Map{
		"noErr": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC JsonString returns non-empty -- non-empty collection", actual)
}

func Test_Reg_KVC_Serialize_NonEmpty(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(1)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})

	// Act
	b, err := kvc.Serialize()

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"hasData": len(b) > 0,
	}
	expected := args.Map{
		"noErr": true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC Serialize returns bytes -- non-empty collection", actual)
}

func Test_Reg_KVC_JsonRoundTrip(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	kvc.Add(coredynamic.KeyVal{Key: "b", Value: "two"})

	// Act
	jr := kvc.JsonPtr()
	kvc2 := coredynamic.EmptyKeyValCollection()
	_, parseErr := kvc2.ParseInjectUsingJson(jr)

	// Assert
	actual := args.Map{
		"noParseErr": parseErr == nil,
		"sameLen":    kvc2.Length() == kvc.Length(),
	}
	expected := args.Map{
		"noParseErr": true,
		"sameLen": true,
	}
	expected.ShouldBeEqual(t, 0, "KVC round-trip returns same length -- serialize then parse", actual)
}

func Test_Reg_KVC_Empty_Json(t *testing.T) {
	// Arrange
	kvc := coredynamic.EmptyKeyValCollection()

	// Act
	j := kvc.Json()

	// Assert
	actual := args.Map{"noErr": !j.HasError()}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "KVC Json returns no error -- empty collection", actual)
}