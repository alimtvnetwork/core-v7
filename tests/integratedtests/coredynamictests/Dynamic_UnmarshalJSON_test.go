package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Dynamic.UnmarshalJSON on nil receiver (line 54) ──

func Test_Dynamic_UnmarshalJSON_NilReceiver_I29(t *testing.T) {
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
//
// Source API chain:
//   ParseInjectUsingJsonMust → ParseInjectUsingJson → jsonResult.Unmarshal(it)
//   → json.Unmarshal(bytes, *Dynamic) → Dynamic.UnmarshalJSON(bytes)
//   → corejson.Deserialize.UsingBytes(bytes, it.innerData)
//
// For success: innerData must be a pointer that json.Unmarshal can write to.
// We serialize the raw map, then unmarshal it into a Dynamic whose innerData
// is a *map[string]any.

func Test_Dynamic_ParseInjectUsingJsonMust_Valid_I29(t *testing.T) {
	// Arrange
	innerMap := map[string]any{
		"key": "value",
	}
	// Serialize the RAW map — NOT a Dynamic wrapper.
	jsonResult := corejson.New(innerMap)

	// Create target Dynamic with pointer-to-map as innerData
	// so UnmarshalJSON → json.Unmarshal(bytes, *map[string]any) succeeds.
	targetMap := map[string]any{}
	target := coredynamic.NewDynamicPtr(&targetMap, false)

	// Act
	result := target.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	actual := args.Map{
		"notNil": result != nil,
		"hasKey": targetMap["key"] == "value",
	}
	expected := args.Map{
		"notNil": true,
		"hasKey": true,
	}
	actual.ShouldBeEqual(t, 1, "Dynamic ParseInjectUsingJsonMust valid", expected)
}

// ── Dynamic.JsonStringMust valid (line 149-163) ──

func Test_Dynamic_JsonStringMust_Valid_I29(t *testing.T) {
	// Arrange
	d := coredynamic.NewDynamic("hello", true)

	// Act
	result := d.JsonStringMust()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "Dynamic JsonStringMust valid", expected)
}
