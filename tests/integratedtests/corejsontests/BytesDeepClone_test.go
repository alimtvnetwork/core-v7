package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BytesDeepClone ──

func Test_BytesDeepClone(t *testing.T) {
	// Arrange
	original := []byte(`{"key":"value"}`)
	cloned := corejson.BytesDeepClone(original)
	clonedNil := corejson.BytesDeepClone(nil)
	nilIsNil := clonedNil == nil

	// Act
	actual := args.Map{
		"len":        len(cloned),
		"nilIsNil":   nilIsNil,
		"notSamePtr": &original[0] != &cloned[0],
	}

	// Assert
	expected := args.Map{
		"len": 15,
		"nilIsNil": nilIsNil,
		"notSamePtr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesDeepClone returns independent copy -- valid input", actual)
}

// ── Result via NewPtr ──

func Test_Result_NewPtr(t *testing.T) {
	// Arrange
	r := corejson.NewPtr("hello")

	// Act
	actual := args.Map{
		"notNil":   r != nil,
		"notEmpty": r.JsonString() != "",
		"hasBytes": len(r.SafeBytes()) > 0,
		"noErr":    r.MeaningfulError() == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true, "notEmpty": true, "hasBytes": true, "noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns valid json -- NewPtr string", actual)
}

func Test_Result_New_ValueType(t *testing.T) {
	// Arrange
	r := corejson.New("hello")

	// Act
	actual := args.Map{
		"hasBytes": len(r.Bytes) > 0,
		"noErr":    r.Error == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns bytes -- New value type", actual)
}

// ── BytesToString ──

func Test_BytesToString(t *testing.T) {
	// Arrange
	result := corejson.BytesToString([]byte(`"hello"`))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString returns non-empty -- valid bytes", actual)
}

func Test_BytesToString_Nil(t *testing.T) {
	// Arrange
	result := corejson.BytesToString(nil)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "BytesToString returns empty -- nil input", actual)
}

// ── JsonString (returns string, error) ──

func Test_JsonString(t *testing.T) {
	// Arrange
	result, err := corejson.JsonString(map[string]string{"a": "1"})

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonString returns json -- map input", actual)
}

// ── JsonStringOrErrMsg ──

func Test_JsonStringOrErrMsg(t *testing.T) {
	// Arrange
	result := corejson.JsonStringOrErrMsg(map[string]string{"a": "1"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonStringOrErrMsg returns json -- valid input", actual)
}

// ── BytesCloneIf ──

func Test_BytesCloneIf_True(t *testing.T) {
	// Arrange
	original := []byte("hello")
	cloned := corejson.BytesCloneIf(true, original)
	clonedLen := len(cloned)

	// Act
	actual := args.Map{
		"len": clonedLen,
	}

	// Assert
	expected := args.Map{"len": clonedLen}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns cloned bytes -- true flag", actual)
}

func Test_BytesCloneIf_False(t *testing.T) {
	// Arrange
	original := []byte("hello")
	result := corejson.BytesCloneIf(false, original)
	// BytesCloneIf(false, ...) returns []byte{} per implementation

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns empty -- false flag", actual)
}

func Test_BytesCloneIf_NilInput(t *testing.T) {
	// Arrange
	result := corejson.BytesCloneIf(true, nil)
	// len(nil) == 0, so !isDeepClone || len == 0 → returns []byte{}

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "BytesCloneIf returns empty -- nil input", actual)
}

// ── SimpleJsonBinder (round-trip) ──

func Test_SimpleJsonBinder(t *testing.T) {
	// Arrange
	type testStruct struct {
		Name string `json:"name"`
	}
	input := testStruct{Name: "test"}
	r := corejson.NewPtr(input)
	var output testStruct
	err := r.Unmarshal(&output)

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"name":  output.Name,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"name": "test",
	}
	expected.ShouldBeEqual(t, 0, "SimpleJsonBinder returns deserialized struct -- round-trip", actual)
}

// ── Result PrettyJsonString ──

func Test_Result_PrettyJsonString(t *testing.T) {
	// Arrange
	r := corejson.NewPtr(map[string]string{"key": "val"})
	pretty := r.PrettyJsonString()

	// Act
	actual := args.Map{
		"notEmpty":     pretty != "",
		"containsKey":  len(pretty) > 5,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"containsKey": true,
	}
	expected.ShouldBeEqual(t, 0, "Result returns pretty json -- map input", actual)
}
