package coredynamictests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MapAnyItems: GetItemRef branches ──

func Test_MapAnyItems_GetItemRef_MissingKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	var out int

	// Act
	err := m.GetItemRef("missing", &out)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetItemRef returns error -- missing key", actual)
}

func Test_MapAnyItems_GetItemRef_NilRef_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	err := m.GetItemRef("a", nil)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetItemRef returns error -- nil reference", actual)
}

func Test_MapAnyItems_GetItemRef_NonPointer(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	err := m.GetItemRef("a", 42) // non-pointer

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetItemRef returns error -- non-pointer reference", actual)
}

func Test_MapAnyItems_GetItemRef_TypeMismatch_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	val := "hello"
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": &val})
	var out int

	// Act
	err := m.GetItemRef("a", &out)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetItemRef returns error -- type mismatch", actual)
}

func Test_MapAnyItems_GetItemRef_MatchingPtrType(t *testing.T) {
	// Arrange
	val := "hello"
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": &val})
	var out string

	// Act
	err := m.GetItemRef("a", &out)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"val":   out,
	}
	expected := args.Map{
		"noErr": true,
		"val":   "hello",
	}
	expected.ShouldBeEqual(t, 0, "GetItemRef returns correct -- matching ptr type", actual)
}

// ── MapAnyItems: GetUsingUnmarshallAt error path ──

func Test_MapAnyItems_GetUsingUnmarshallAt_MissingKey(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})
	var out int

	// Act
	err := m.GetUsingUnmarshallAt("missing", &out)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetUsingUnmarshallAt returns error -- missing key", actual)
}

// ── MapAnyItems: GetNewMapUsingKeys ──

func Test_MapAnyItems_GetNewMapUsingKeys_Empty_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	result := m.GetNewMapUsingKeys(false)

	// Assert
	actual := args.Map{"isEmpty": result.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetNewMapUsingKeys returns empty -- no keys", actual)
}

func Test_MapAnyItems_GetNewMapUsingKeys_NotPanic(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{
		"a": 1,
		"b": 2,
	})

	// Act
	result := m.GetNewMapUsingKeys(false, "a", "missing")

	// Assert
	actual := args.Map{"len": result.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "GetNewMapUsingKeys returns partial -- missing key no panic", actual)
}

// ── MapAnyItems: DiffChangedOnlyRaw ──

func Test_MapAnyItems_DiffChangedOnlyRaw_NoDiff(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{"a": 1})

	// Act
	diff := m.DiffRaw(false, map[string]any{"a": 1})

	// Assert
	actual := args.Map{"isEmpty": len(diff) == 0}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffChangedOnlyRaw returns empty -- no diff", actual)
}

// ── MapAnyItems: ToMapResults error path ──

func Test_MapAnyItems_ToMapResults_Valid(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{
		"key": "value",
	})

	// Act
	mr, err := m.JsonMapResults()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"notNil": mr != nil,
	}
	expected := args.Map{
		"noErr":  true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ToMapResults returns correct -- valid items", actual)
}

// ── MapAnyItems: CloneUsingJson ──

func Test_MapAnyItems_CloneUsingJson(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItemsUsingItems(map[string]any{
		"key": "value",
	})

	// Act
	cloned, err := m.ClonePtr()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"notNil": cloned != nil,
	}
	expected := args.Map{
		"noErr":  true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "CloneUsingJson returns correct -- valid items", actual)
}

// ── AnyCollection: JsonString ──

func Test_AnyCollection_JsonString_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(3)
	ac.Add(1)
	ac.Add(2)
	ac.Add(3)

	// Act
	str, err := ac.JsonString()

	// Assert
	actual := args.Map{
		"noErr":   err == nil,
		"hasStr":  len(str) > 0,
	}
	expected := args.Map{
		"noErr":   true,
		"hasStr":  true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct -- AnyCollection", actual)
}

func Test_AnyCollection_JsonStringMust_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(1)
	ac.Add("a")

	// Act
	str := ac.JsonStringMust()

	// Assert
	actual := args.Map{"hasStr": len(str) > 0}
	expected := args.Map{"hasStr": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct -- AnyCollection", actual)
}

// ── DynamicCollection: JsonString ──

func Test_DynamicCollection_JsonString_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(2)
	dc.AddAny("hello", true)

	// Act
	str, err := dc.JsonString()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"noErr":  true,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct -- DynamicCollection", actual)
}

func Test_DynamicCollection_JsonStringMust_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(1)
	dc.AddAny("test", true)

	// Act
	str := dc.JsonStringMust()

	// Assert
	actual := args.Map{"hasStr": len(str) > 0}
	expected := args.Map{"hasStr": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct -- DynamicCollection", actual)
}

// ── Collection[T]: JsonString, LengthLock, ItemsLock ──

func Test_Collection_JsonString(t *testing.T) {
	// Arrange
	col := coredynamic.NewCollection[int](3)
	col.Add(1)
	col.Add(2)

	// Act
	str, err := col.JsonString()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"noErr":  true,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct -- Collection[int]", actual)
}

func Test_Collection_LengthLock(t *testing.T) {
	// Arrange
	col := coredynamic.NewCollection[string](2)
	col.Add("a")

	// Act
	length := col.LengthLock()

	// Assert
	actual := args.Map{"len": length}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LengthLock returns correct -- Collection", actual)
}

func Test_Collection_ItemsLock(t *testing.T) {
	// Arrange
	col := coredynamic.NewCollection[string](2)
	col.Add("a")
	col.Add("b")

	// Act
	items := col.ItemsLock()

	// Assert
	actual := args.Map{"len": len(items)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ItemsLock returns correct -- Collection", actual)
}

// ── Dynamic: JsonString, MarshalJSON error path, JsonBytes ──

func Test_Dynamic_JsonString_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]any{"k": "v"}, true)

	// Act
	str, err := d.JsonString()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"noErr":  true,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct -- Dynamic", actual)
}

func Test_Dynamic_JsonStringMust_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]any{"k": "v"}, true)

	// Act
	str := d.JsonStringMust()

	// Assert
	actual := args.Map{"hasStr": len(str) > 0}
	expected := args.Map{"hasStr": true}
	expected.ShouldBeEqual(t, 0, "JsonStringMust returns correct -- Dynamic", actual)
}

func Test_Dynamic_JsonBytes_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]any{"k": "v"}, true)

	// Act
	bytes, err := d.JsonBytes()

	// Assert
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonBytes returns correct -- Dynamic", actual)
}

func Test_Dynamic_MarshalJSON_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic(map[string]any{"k": "v"}, true)

	// Act
	bytes, err := d.MarshalJSON()

	// Assert
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "MarshalJSON returns correct -- Dynamic", actual)
}

// ── KeyVal: SetTo error path, Serialize ──

func Test_KeyVal_SetTo_Error(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "key", Value: "value"}
	var wrongType int

	// Act
	err := kv.ReflectSetTo(&wrongType)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SetTo returns error -- type mismatch", actual)
}

func Test_KeyVal_Serialize_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	kv := coredynamic.KeyVal{Key: "key", Value: "value"}

	// Act
	bytes, err := kv.Serialize()

	// Assert
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct -- valid KeyVal", actual)
}

// ── KeyValCollection: Serialize, JsonString, ToMapResults ──

func Test_KeyValCollection_Serialize_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	bytes, err := kvc.Serialize()

	// Assert
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(bytes) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "Serialize returns correct -- KeyValCollection", actual)
}

func Test_KeyValCollection_JsonString_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	str, err := kvc.JsonString()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"noErr":  true,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct -- KeyValCollection", actual)
}

func Test_KeyValCollection_ToMapResults(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})

	// Act
	mr, err := kvc.JsonMapResults()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"notNil": mr != nil,
	}
	expected := args.Map{
		"noErr":  true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ToMapResults returns correct -- KeyValCollection", actual)
}

// ── TypedDynamic: JsonString ──

func Test_TypedDynamic_JsonString(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[map[string]any](map[string]any{"k": "v"}, true)

	// Act
	str, err := td.JsonString()

	// Assert
	actual := args.Map{
		"noErr":  err == nil,
		"hasStr": len(str) > 0,
	}
	expected := args.Map{
		"noErr":  true,
		"hasStr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns correct -- TypedDynamic", actual)
}

// ── ReflectSetFromTo: byte conversion paths ──

func Test_ReflectSetFromTo_BytesToStruct_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	jsonBytes := []byte(`{"name":"test"}`)
	var out sample

	// Act
	err := coredynamic.ReflectSetFromTo(jsonBytes, &out)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"name":  out.Name,
	}
	expected := args.Map{
		"noErr": true,
		"name":  "test",
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct -- bytes to struct", actual)
}

func Test_ReflectSetFromTo_StructToBytes_FromMapAnyItemsGetItemRe(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	from := sample{Name: "test"}
	var out []byte

	// Act
	err := coredynamic.ReflectSetFromTo(from, &out)

	// Assert
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(out) > 0,
	}
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "ReflectSetFromTo returns correct -- struct to bytes", actual)
}

// ── PointerOrNonPointerUsingReflectValue: pointer output ──

func Test_PointerOrNonPointer_PointerOutput(t *testing.T) {
	// Arrange
	val := 42
	rv := reflect.ValueOf(val)

	// Act
	result, resultRv := coredynamic.PointerOrNonPointerUsingReflectValue(true, rv)

	// Assert
	actual := args.Map{
		"notNil":    result != nil,
		"rvValid":   resultRv.IsValid(),
	}
	expected := args.Map{
		"notNil":    true,
		"rvValid":   true,
	}
	expected.ShouldBeEqual(t, 0, "PointerOrNonPointer returns ptr -- pointer output requested", actual)
}

// ── SafeZeroSet: non-pointer ──

func Test_SafeZeroSet_NonPointer(t *testing.T) {
	// Arrange
	val := 42
	rv := reflect.ValueOf(val)

	// Act — should not panic, just return
	coredynamic.SafeZeroSet(rv)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	expected.ShouldBeEqual(t, 0, "SafeZeroSet returns safely -- non-pointer", actual)
}

// ── AnyCollection: ParseInjectUsingJsonMust ──

func Test_AnyCollection_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	ac := coredynamic.NewAnyCollection(2)
	ac.Add(1)
	ac.Add(2)
	jsonResult := corejson.New(ac)

	// Act
	result := ac.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust returns correct -- valid json", actual)
}

// ── KeyValCollection: ParseInjectUsingJsonMust ──

func Test_KeyValCollection_ParseInjectUsingJsonMust(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: 1})
	jsonResult := corejson.New(kvc)

	// Act
	result := kvc.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust returns correct -- valid json", actual)
}

// ── CastTo: null non-pointer path ──

func Test_CastTo_NullNonPointer(t *testing.T) {
	// Arrange
	var ptr *int
	rv := reflect.ValueOf(ptr)

	// Act
	result := coredynamic.CastTo(
		false, // non-pointer output
		rv,
	)

	// Assert
	actual := args.Map{"hasErr": result.Error != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CastTo returns error -- null non-pointer", actual)
}
