package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── KeyVal.CastKeyVal nil receiver (line 134) ──

func Test_KeyVal_CastKeyVal_NilReceiver_I29(t *testing.T) {
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

func Test_KeyVal_JsonParseSelfInject_I29(t *testing.T) {
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

func Test_KeyValCollection_ParseInjectUsingJson_Valid_I29(t *testing.T) {
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

func Test_KeyValCollection_JsonParseSelfInject_I29(t *testing.T) {
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

func Test_KeyValCollection_Serialize_Valid_I29(t *testing.T) {
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

func Test_KeyValCollection_JsonString_Valid_I29(t *testing.T) {
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

func Test_KeyValCollection_ParseInjectUsingJsonMust_Valid_I29(t *testing.T) {
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
