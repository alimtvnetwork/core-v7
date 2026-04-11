package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── deserializeFromResultTo: Byte, ByteMust ──

func Test_ResultTo_Byte_Valid(t *testing.T) {
	// Arrange
	jsonResult := corejson.Serialize.Apply(byte(65))

	// Act
	result, err := corejson.Deserialize.ResultTo.Byte(jsonResult)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": byte(65),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo.Byte returns 65 -- valid", actual)
}

func Test_ResultTo_ByteMust_Valid(t *testing.T) {
	// Arrange
	jsonResult := corejson.Serialize.Apply(byte(42))

	// Act
	result := corejson.Deserialize.ResultTo.ByteMust(jsonResult)

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": byte(42)}
	expected.ShouldBeEqual(t, 0, "ResultTo.ByteMust returns 42 -- valid", actual)
}

// ── deserializeFromResultTo: MapStringString, MapStringStringMust ──

func Test_ResultTo_MapStringString_Valid(t *testing.T) {
	// Arrange
	m := map[string]string{"hello": "world"}
	jsonResult := corejson.Serialize.Apply(m)

	// Act
	result, err := corejson.Deserialize.ResultTo.MapStringString(jsonResult)

	// Assert
	actual := args.Map{
		"hello": result["hello"],
		"noErr": err == nil,
	}
	expected := args.Map{
		"hello": "world",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo.MapStringString returns map -- valid", actual)
}

func Test_ResultTo_MapStringStringMust_Valid(t *testing.T) {
	// Arrange
	m := map[string]string{"a": "b"}
	jsonResult := corejson.Serialize.Apply(m)

	// Act
	result := corejson.Deserialize.ResultTo.MapStringStringMust(jsonResult)

	// Assert
	actual := args.Map{"a": result["a"]}
	expected := args.Map{"a": "b"}
	expected.ShouldBeEqual(t, 0, "ResultTo.MapStringStringMust returns map -- valid", actual)
}

// ── deserializeFromResultTo: ResultPtr, ResultPtrMust ──

func Test_ResultTo_ResultPtr_Valid(t *testing.T) {
	// Arrange
	inner := corejson.New("hello")
	outerResult := corejson.Serialize.Apply(inner)

	// Act
	result, err := corejson.Deserialize.ResultTo.ResultPtr(outerResult)

	// Assert
	actual := args.Map{
		"notNil": result != nil,
		"noErr": err == nil,
	}
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo.ResultPtr returns ptr -- valid", actual)
}

func Test_ResultTo_ResultPtrMust_Valid(t *testing.T) {
	// Arrange
	inner := corejson.New("hello")
	outerResult := corejson.Serialize.Apply(inner)

	// Act
	result := corejson.Deserialize.ResultTo.ResultPtrMust(outerResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultTo.ResultPtrMust returns ptr -- valid", actual)
}

// ── deserializeFromResultTo: Bytes, BytesMust ──

func Test_ResultTo_Bytes_Valid(t *testing.T) {
	// Arrange
	inner := corejson.New("data")
	outerResult := corejson.Serialize.Apply(inner)

	// Act
	result, err := corejson.Deserialize.ResultTo.Bytes(outerResult)

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo.Bytes returns bytes -- valid", actual)
}

func Test_ResultTo_BytesMust_Valid(t *testing.T) {
	// Arrange
	inner := corejson.New("data")
	outerResult := corejson.Serialize.Apply(inner)

	// Act
	result := corejson.Deserialize.ResultTo.BytesMust(outerResult)

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "ResultTo.BytesMust returns bytes -- valid", actual)
}

// ── deserializeFromResultTo: ResultsPtrCollection, ResultsPtrCollectionMust ──

func Test_ResultTo_ResultsPtrCollection_Valid(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b")
	jsonResult := corejson.Serialize.Apply(coll)

	// Act
	result, err := corejson.Deserialize.ResultTo.ResultsPtrCollection(jsonResult)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo.ResultsPtrCollection returns coll -- valid", actual)
}

func Test_ResultTo_ResultsPtrCollectionMust_Valid(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("x")
	jsonResult := corejson.Serialize.Apply(coll)

	// Act
	result := corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultTo.ResultsPtrCollectionMust returns coll -- valid", actual)
}

// ── deserializeFromResultTo: MapResults, MapResultsMust ──

func Test_ResultTo_MapResults_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	jsonResult := corejson.Serialize.Apply(mr)

	// Act
	result, err := corejson.Deserialize.ResultTo.MapResults(jsonResult)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ResultTo.MapResults returns map -- valid", actual)
}

func Test_ResultTo_MapResultsMust_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	jsonResult := corejson.Serialize.Apply(mr)

	// Act
	result := corejson.Deserialize.ResultTo.MapResultsMust(jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ResultTo.MapResultsMust returns map -- valid", actual)
}
