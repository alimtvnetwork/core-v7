package corejsontests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Result creation ──

func Test_NewResult_Serialize(t *testing.T) {
	// Arrange
	result := corejson.NewResult.Serialize(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"hasBytes": result.HasBytes(),
		"noErr":    result.IsEmptyError(),
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "NewResult.Serialize produces valid result -- map input", actual)
}

func Test_Result_Methods(t *testing.T) {
	// Arrange
	result := corejson.NewResult.Serialize("hello")

	// Act
	actual := args.Map{
		"isEmpty":    result.IsEmpty(),
		"hasBytes":   result.HasBytes(),
		"jsonString": len(result.JsonString()) > 0,
		"string":     len(result.String()) > 0,
		"hasError":   result.HasError(),
		"emptyError": result.IsEmptyError(),
	}

	// Assert
	expected := args.Map{
		"isEmpty":    false,
		"hasBytes":   true,
		"jsonString": true,
		"string":     true,
		"hasError":   false,
		"emptyError": true,
	}
	expected.ShouldBeEqual(t, 0, "Result has no error -- valid input", actual)
}

func Test_Result_Clone(t *testing.T) {
	// Arrange
	result := corejson.NewResult.Serialize("hello")
	cloned := result.Clone(false)

	// Act
	actual := args.Map{
		"sameJson": cloned.JsonString() == result.JsonString(),
	}

	// Assert
	expected := args.Map{
		"sameJson": true,
	}
	expected.ShouldBeEqual(t, 0, "Result.Clone produces equal json -- valid", actual)
}

func Test_Result_ClonePtr(t *testing.T) {
	// Arrange
	result := corejson.NewResult.Serialize("hello")
	cloned := result.ClonePtr(false)

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Result.ClonePtr returns non-nil -- valid", actual)
}

func Test_Result_Nil_ClonePtr(t *testing.T) {
	// Arrange
	var result *corejson.Result
	cloned := result.ClonePtr(false)

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Result.ClonePtr returns nil -- nil receiver", actual)
}

// ── Serialize/Deserialize roundtrip ──

func Test_SerializeDeserialize_Roundtrip(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string
		Age  int
	}
	original := testStruct{Name: "test", Age: 25}

	rawBytes, err := corejson.Serialize.Raw(original)

	// Act
	actual1 := args.Map{
		"noErr": err == nil,
		"hasBytes": len(rawBytes) > 0,
	}

	// Assert
	expected1 := args.Map{
		"noErr": true,
		"hasBytes": true,
	}
	expected1.ShouldBeEqual(t, 0, "Serialize.Raw succeeds -- struct input", actual1)

	var deserialized testStruct
	err = corejson.Deserialize.UsingBytes(rawBytes, &deserialized)
	actual2 := args.Map{
		"noErr":    err == nil,
		"sameName": deserialized.Name == original.Name,
		"sameAge":  deserialized.Age == original.Age,
	}
	expected2 := args.Map{
		"noErr":    true,
		"sameName": true,
		"sameAge":  true,
	}
	expected2.ShouldBeEqual(t, 1, "Deserialize.UsingBytes roundtrip -- struct", actual2)
}

// ── AnyTo ──

func Test_AnyTo_SerializedJsonResult(t *testing.T) {
	// Arrange
	result := corejson.AnyTo.SerializedJsonResult(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"notNil":  result != nil,
		"noErr":   result.IsEmptyError(),
		"hasData": result.HasBytes(),
	}

	// Assert
	expected := args.Map{
		"notNil":  true,
		"noErr":   true,
		"hasData": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedJsonResult returns valid -- map input", actual)
}

func Test_AnyTo_SerializedRaw_FromNewResultSerialize(t *testing.T) {
	// Arrange
	rawBytes, err := corejson.AnyTo.SerializedRaw(map[string]string{"k": "v"})

	// Act
	actual := args.Map{
		"noErr":    err == nil,
		"hasBytes": len(rawBytes) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":    true,
		"hasBytes": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedRaw returns bytes -- map input", actual)
}

func Test_AnyTo_SerializedString_FromNewResultSerialize(t *testing.T) {
	// Arrange
	s, err := corejson.AnyTo.SerializedString(map[string]int{"a": 1})

	// Act
	actual := args.Map{
		"noErr":      err == nil,
		"hasContent": len(s) > 0,
	}

	// Assert
	expected := args.Map{
		"noErr":      true,
		"hasContent": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyTo.SerializedString returns json string -- map input", actual)
}

func Test_AnyTo_SafeJsonString_FromNewResultSerialize(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonString returns non-empty -- map input", actual)
}

func Test_AnyTo_SafeJsonPrettyString_FromNewResultSerialize(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.SafeJsonPrettyString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.SafeJsonPrettyString returns non-empty -- map input", actual)
}

func Test_AnyTo_JsonString_FromNewResultSerialize(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonString(map[string]int{"a": 1})

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonString returns non-empty -- map input", actual)
}

func Test_AnyTo_JsonStringMust_FromNewResultSerialize(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.JsonStringMust(map[string]int{"a": 1})

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.JsonStringMust returns non-empty -- map input", actual)
}

func Test_AnyTo_PrettyStringMust_FromNewResultSerialize(t *testing.T) {
	// Arrange
	s := corejson.AnyTo.PrettyStringMust(map[string]int{"a": 1})

	// Act
	actual := args.Map{"hasContent": len(s) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "AnyTo.PrettyStringMust returns non-empty -- map input", actual)
}

// ── Deserialize from bytes (BytesTo, not FromBytesTo) ──

func Test_DeserializeFromBytes_String(t *testing.T) {
	// Arrange
	b, _ := json.Marshal("hello")
	s, err := corejson.Deserialize.BytesTo.String(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": s,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "hello",
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.String roundtrip -- hello", actual)
}

func Test_DeserializeFromBytes_Integer(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(42)
	val, err := corejson.Deserialize.BytesTo.Integer(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Integer roundtrip -- 42", actual)
}

func Test_DeserializeFromBytes_Integer64(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(int64(999))
	val, err := corejson.Deserialize.BytesTo.Integer64(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": val,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": int64(999),
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Integer64 roundtrip -- 999", actual)
}

func Test_DeserializeFromBytes_MapAnyItem(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(map[string]any{"k": "v"})
	m, err := corejson.Deserialize.BytesTo.MapAnyItem(b)

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"hasKey": m["k"] == "v",
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.MapAnyItem roundtrip -- map", actual)
}

func Test_DeserializeFromBytes_MapStringString(t *testing.T) {
	// Arrange
	b, _ := json.Marshal(map[string]string{"k": "v"})
	m, err := corejson.Deserialize.BytesTo.MapStringString(b)

	// Act
	actual := args.Map{
		"noErr":  err == nil,
		"hasKey": m["k"] == "v",
	}

	// Assert
	expected := args.Map{
		"noErr":  true,
		"hasKey": true,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.MapStringString roundtrip -- map", actual)
}

func Test_DeserializeFromBytes_Bytes(t *testing.T) {
	// Arrange
	original := []byte{1, 2, 3}
	b, _ := json.Marshal(original)
	result, err := corejson.Deserialize.BytesTo.Bytes(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len":   len(result),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len":   3,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Bytes roundtrip -- 3 bytes", actual)
}

func Test_DeserializeFromBytes_Integers(t *testing.T) {
	// Arrange
	b, _ := json.Marshal([]int{1, 2, 3})
	val, err := corejson.Deserialize.BytesTo.Integers(b)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(val),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "DeserializeFromBytes.Integers roundtrip -- 3 ints", actual)
}

// ── Empty creators ──

func Test_Empty_Result_FromNewResultSerialize(t *testing.T) {
	// Arrange
	r := corejson.Empty.Result()

	// Act
	actual := args.Map{"isEmpty": r.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty.Result returns empty -- no data", actual)
}

func Test_Empty_ResultPtr_FromNewResultSerialize(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultPtr()

	// Act
	actual := args.Map{
		"notNil": r != nil,
		"isEmpty": r.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"isEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Empty.ResultPtr returns empty ptr -- no data", actual)
}
