package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage75 — coredynamic remaining 29 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── AnyCollection.JsonStringMust valid (line 495) ──

func Test_Cov75_AnyCollection_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewAnyCollection(3)
	coll.Add("hello")
	coll.Add(42)

	// Act
	jsonStr := coll.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(jsonStr) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "AnyCollection JsonStringMust valid", expected)
}

// ── Collection.JsonString valid (line 355) ──

func Test_Cov75_Collection_JsonString_Valid_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](3)
	coll.Add("a")
	coll.Add("b")

	// Act
	jsonStr, err := coll.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "Collection JsonString valid", expected)
}

// ── Collection.JsonStringMust valid (line 364) ──

func Test_Cov75_Collection_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](2)
	coll.Add("test")

	// Act
	result := coll.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "Collection JsonStringMust valid", expected)
}

// ── CollectionLock.LengthLock (line 15) ──

func Test_Cov75_CollectionLock_LengthLock_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[int](5)
	coll.Add(1)
	coll.Add(2)

	// Act
	length := coll.LengthLock()

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 2}
	actual.ShouldBeEqual(t, 1, "CollectionLock LengthLock", expected)
}

// ── CollectionLock.RemoveAtLock invalid index (line 125) ──

func Test_Cov75_CollectionLock_RemoveAtLock_InvalidIndex_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](2)
	coll.Add("a")

	// Act
	removed := coll.RemoveAtLock(99)

	// Assert
	actual := args.Map{"removed": removed}
	expected := args.Map{"removed": false}
	actual.ShouldBeEqual(t, 1, "CollectionLock RemoveAtLock invalid index", expected)
}

// ── DynamicCollection.JsonStringMust valid (line 426) ──

func Test_Cov75_DynamicCollection_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	dc := coredynamic.NewDynamicCollection(3)
	dc.AddAny("hello", true)

	// Act
	jsonStr := dc.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(jsonStr) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "DynamicCollection JsonStringMust valid", expected)
}

// ── Dynamic.UnmarshalJSON on nil receiver (line 54) ──

func Test_Cov75_Dynamic_UnmarshalJSON_NilReceiver_I29(t *testing.T) {
	// Arrange
	var d *coredynamic.Dynamic

	// Act
	err := d.UnmarshalJSON([]byte(`"hello"`))

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "Dynamic UnmarshalJSON nil receiver", expected)
}

// ── Dynamic.ParseInjectUsingJsonMust valid (line 123) ──

func Test_Cov75_Dynamic_ParseInjectUsingJsonMust_Valid_I29(t *testing.T) {
	// Arrange — create a Dynamic with a map so it round-trips cleanly
	innerMap := map[string]any{"key": "value"}
	d := coredynamic.NewDynamic(innerMap, true)
	jsonResult := corejson.New(d)

	// Act
	target := coredynamic.NewDynamic(map[string]any{}, false)
	result := target.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	actual.ShouldBeEqual(t, 1, "Dynamic ParseInjectUsingJsonMust valid", expected)
}

// ── Dynamic.JsonStringMust valid (line 149-163) ──

func Test_Cov75_Dynamic_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	result := d.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "Dynamic JsonStringMust valid", expected)
}

// ── KeyVal.CastKeyVal nil receiver (line 134) ──

func Test_Cov75_KeyVal_CastKeyVal_NilReceiver_I29(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	var k string
	var v int
	err := kv.CastKeyVal(&k, &v)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "KeyVal CastKeyVal nil receiver", expected)
}

// ── KeyVal.JsonParseSelfInject (line 300) ──

func Test_Cov75_KeyVal_JsonParseSelfInject_I29(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "a", Value: "b"}
	jsonResult := corejson.New(kv)

	// Act
	err := kv.JsonParseSelfInject(&jsonResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "KeyVal JsonParseSelfInject", expected)
}

// ── KeyValCollection operations ──

func Test_Cov75_KeyValCollection_ParseInjectUsingJson_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k1", Value: "v1"})
	jsonResult := corejson.New(kvc)

	// Act
	result, err := kvc.ParseInjectUsingJson(&jsonResult)

	// Assert
	actual := args.Map{
		"notNil":   result != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "KeyValCollection ParseInjectUsingJson valid", expected)
}

func Test_Cov75_KeyValCollection_JsonParseSelfInject_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "x", Value: "y"})
	jsonResult := corejson.New(kvc)

	// Act
	err := kvc.JsonParseSelfInject(&jsonResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "KeyValCollection JsonParseSelfInject", expected)
}

func Test_Cov75_KeyValCollection_Serialize_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "b"})

	// Act
	bytes, err := kvc.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "KeyValCollection Serialize valid", expected)
}

func Test_Cov75_KeyValCollection_JsonString_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "b"})

	// Act
	jsonStr, err := kvc.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "KeyValCollection JsonString valid", expected)
}

func Test_Cov75_KeyValCollection_ParseInjectUsingJsonMust_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jsonResult := corejson.New(kvc)

	// Act
	result := kvc.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	actual.ShouldBeEqual(t, 1, "KeyValCollection ParseInjectUsingJsonMust valid", expected)
}

// ── MapAnyItems.JsonMapResults (line 903) ──

func Test_Cov75_MapAnyItems_JsonMapResults_Valid_I29(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(3)
	m.Add("key1", "val1")
	m.Add("key2", 42)

	// Act
	results, err := m.JsonMapResults()

	// Assert
	actual := args.Map{
		"notNil":   results != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "MapAnyItems JsonMapResults valid", expected)
}

// ── MapAnyItems.GetUsingUnmarshallAt type mismatch (line 350) ──

func Test_Cov75_MapAnyItems_GetUsingUnmarshallAt_TypeMismatch_I29(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(3)
	m.Add("key1", "stringValue")

	var target int

	// Act
	err := m.GetUsingUnmarshallAt("key1", &target)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "MapAnyItems GetUsingUnmarshallAt type mismatch", expected)
}

// ── ReflectInterfaceVal non-ptr (line 20) ──

func Test_Cov75_ReflectInterfaceVal_NonPtr_I29(t *testing.T) {
	// Arrange / Act
	result := coredynamic.ReflectInterfaceVal(42)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	actual.ShouldBeEqual(t, 1, "ReflectInterfaceVal non-ptr", expected)
}

// ── ReflectSetFromTo: []byte → struct (line 159-167) ──

func Test_Cov75_ReflectSetFromTo_BytesToStruct_I29(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	data, _ := json.Marshal(sample{Name: "test"})
	var target sample

	// Act
	err := coredynamic.ReflectSetFromTo(data, &target)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
		"name":     target.Name,
	}
	expected := args.Map{
		"hasError": false,
		"name":     "test",
	}
	actual.ShouldBeEqual(t, 1, "ReflectSetFromTo bytes to struct", expected)
}

// ── ReflectSetFromTo: struct → *[]byte (line 174-180) ──

func Test_Cov75_ReflectSetFromTo_StructToBytes_I29(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	from := sample{Name: "test"}
	var target []byte

	// Act
	err := coredynamic.ReflectSetFromTo(from, &target)

	// Assert
	actual := args.Map{
		"hasError":   err != nil,
		"hasContent": len(target) > 0,
	}
	expected := args.Map{
		"hasError":   false,
		"hasContent": true,
	}
	actual.ShouldBeEqual(t, 1, "ReflectSetFromTo struct to bytes", expected)
}

// ── SafeZeroSet non-pointer (line 18) ──

func Test_Cov75_SafeZeroSet_NonPointer_I29(t *testing.T) {
	// Arrange
	val := reflect.ValueOf(42)

	// Act — should not panic on non-pointer
	coredynamic.SafeZeroSet(val)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	actual.ShouldBeEqual(t, 1, "SafeZeroSet non-pointer", expected)
}

// ── TypedDynamic.JsonString valid (line 117) ──

func Test_Cov75_TypedDynamic_JsonString_Valid_I29(t *testing.T) {
	// Arrange
	td := coredynamic.NewTypedDynamic[string]("hello", true)

	// Act
	jsonStr, err := td.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "TypedDynamic JsonString valid", expected)
}
