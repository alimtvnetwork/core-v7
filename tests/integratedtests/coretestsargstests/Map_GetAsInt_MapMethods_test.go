package coretestsargstests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Map.GetAsInt / GetAsIntDefault ──

func Test_Map_GetAsInt_Found(t *testing.T) {
	// Arrange
	m := args.Map{"count": 42}

	// Act
	val, ok := m.GetAsInt("count")

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": 42,
		"ok":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns value -- key found", actual)
}

func Test_Map_GetAsInt_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"other": "x"}

	// Act
	val, ok := m.GetAsInt("count")

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": 0,
		"ok":  false,
	}
	expected.ShouldBeEqual(t, 0, "GetAsInt returns zero -- key missing", actual)
}

func Test_Map_GetAsIntDefault_Fallback(t *testing.T) {
	// Arrange
	m := args.Map{"other": "x"}

	// Act
	val := m.GetAsIntDefault("count", 99)

	// Assert
	actual := args.Map{"val": val}
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault returns default -- missing key", actual)
}

func Test_Map_GetAsIntDefault_Found(t *testing.T) {
	// Arrange
	m := args.Map{"count": 10}

	// Act
	val := m.GetAsIntDefault("count", 99)

	// Assert
	actual := args.Map{"val": val}
	expected := args.Map{"val": 10}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault returns value -- key found", actual)
}

// ── Map.GetAsBool / GetAsBoolDefault ──

func Test_Map_GetAsBool_Found(t *testing.T) {
	// Arrange
	m := args.Map{"flag": true}

	// Act
	val, ok := m.GetAsBool("flag")

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": true,
		"ok":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsBool returns value -- key found", actual)
}

func Test_Map_GetAsBool_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	val, ok := m.GetAsBool("flag")

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": false,
		"ok":  false,
	}
	expected.ShouldBeEqual(t, 0, "GetAsBool returns false -- missing key", actual)
}

func Test_Map_GetAsBoolDefault_Fallback(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	val := m.GetAsBoolDefault("flag", true)

	// Assert
	actual := args.Map{"val": val}
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "GetAsBoolDefault returns default -- missing", actual)
}

// ── Map.GetAsString / GetAsStringDefault ──

func Test_Map_GetAsString_Found(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}

	// Act
	val, ok := m.GetAsString("name")

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": "hello",
		"ok":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsString returns value -- key found", actual)
}

func Test_Map_GetAsStringDefault_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	val := m.GetAsStringDefault("name")

	// Assert
	actual := args.Map{"val": val}
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "GetAsStringDefault returns empty -- missing", actual)
}

// ── Map.GetAsStrings / GetAsAnyItems ──

func Test_Map_GetAsStrings_Found(t *testing.T) {
	// Arrange
	m := args.Map{"items": []string{"a", "b"}}

	// Act
	val, ok := m.GetAsStrings("items")

	// Assert
	actual := args.Map{
		"length": len(val),
		"ok":     ok,
	}
	expected := args.Map{
		"length": 2,
		"ok":     true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsStrings returns slice -- key found", actual)
}

func Test_Map_GetAsStrings_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	val, ok := m.GetAsStrings("items")

	// Assert
	actual := args.Map{
		"length": len(val),
		"ok":     ok,
	}
	expected := args.Map{
		"length": 0,
		"ok":     false,
	}
	expected.ShouldBeEqual(t, 0, "GetAsStrings returns empty -- missing", actual)
}

func Test_Map_GetAsAnyItems_Found(t *testing.T) {
	// Arrange
	m := args.Map{"items": []any{1, "two"}}

	// Act
	val, ok := m.GetAsAnyItems("items")

	// Assert
	actual := args.Map{
		"length": len(val),
		"ok":     ok,
	}
	expected := args.Map{
		"length": 2,
		"ok":     true,
	}
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems returns slice -- key found", actual)
}

// ── Map.HasDefinedAll / IsKeyInvalid / IsKeyMissing ──

func Test_Map_HasDefinedAll_AllPresent(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "two",
	}

	// Act
	result := m.HasDefinedAll("a", "b")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns true -- all defined", actual)
}

func Test_Map_HasDefinedAll_OneMissing(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	result := m.HasDefinedAll("a", "b")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns false -- one missing", actual)
}

func Test_Map_HasDefinedAll_Empty(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	result := m.HasDefinedAll()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns false -- no names", actual)
}

func Test_Map_IsKeyMissing(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	missing := m.IsKeyMissing("b")
	present := m.IsKeyMissing("a")

	// Assert
	actual := args.Map{
		"missing": missing,
		"present": present,
	}
	expected := args.Map{
		"missing": true,
		"present": false,
	}
	expected.ShouldBeEqual(t, 0, "IsKeyMissing returns correct -- key check", actual)
}

// ── Map.GetLowerCase / GetDirectLower / Expect / Actual / Arrange / SetActual ──

func Test_Map_GetLowerCase(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}

	// Act
	val, ok := m.GetLowerCase("NAME")

	// Assert
	actual := args.Map{
		"val": val,
		"ok":  ok,
	}
	expected := args.Map{
		"val": "hello",
		"ok":  true,
	}
	expected.ShouldBeEqual(t, 0, "GetLowerCase returns value -- uppercase key", actual)
}

func Test_Map_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	m.SetActual("result")

	// Assert
	actual := args.Map{"val": m.Actual()}
	expected := args.Map{"val": "result"}
	expected.ShouldBeEqual(t, 0, "SetActual sets value -- then Actual retrieves", actual)
}

func Test_Map_Arrange(t *testing.T) {
	// Arrange
	m := args.Map{"arrange": "setup"}

	// Act
	result := m.Arrange()

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": "setup"}
	expected.ShouldBeEqual(t, 0, "Arrange returns value -- key present", actual)
}

// ── Map.CompileToStrings / CompileToString ──

func Test_Map_CompileToStrings(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": "two",
	}

	// Act
	lines := m.CompileToStrings()

	// Assert
	actual := args.Map{"length": len(lines)}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "CompileToStrings returns sorted lines -- 2 keys", actual)
}

func Test_Map_CompileToStrings_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	lines := m.CompileToStrings()

	// Assert
	actual := args.Map{"length": len(lines)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "CompileToStrings returns empty -- no keys", actual)
}

func Test_Map_CompileToString(t *testing.T) {
	// Arrange
	m := args.Map{"x": 5}

	// Act
	result := m.CompileToString()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "CompileToString returns string -- single key", actual)
}

// ── Map.GoLiteralLines / GoLiteralString ──

func Test_Map_GoLiteralLines(t *testing.T) {
	// Arrange
	m := args.Map{
		"name":  "hello",
		"count": 5,
	}

	// Act
	lines := m.GoLiteralLines()

	// Assert
	actual := args.Map{"length": len(lines)}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "GoLiteralLines returns lines -- 2 keys", actual)
}

func Test_Map_GoLiteralLines_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	lines := m.GoLiteralLines()

	// Assert
	actual := args.Map{"length": len(lines)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "GoLiteralLines returns empty -- no keys", actual)
}

func Test_Map_GoLiteralString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	result := m.GoLiteralString()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "GoLiteralString returns string -- single key", actual)
}

// ── Map.Seventh ──

func Test_Map_Seventh(t *testing.T) {
	// Arrange
	m := args.Map{"seventh": "val7"}

	// Act
	result := m.Seventh()

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": "val7"}
	expected.ShouldBeEqual(t, 0, "Seventh returns value -- key present", actual)
}

// ── Map nil checks ──

func Test_Map_NilMap_Has(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	result := m.Has("x")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Has returns false -- nil map", actual)
}

func Test_Map_NilMap_HasDefined(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	result := m.HasDefined("x")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasDefined returns false -- nil map", actual)
}

func Test_Map_NilMap_Get(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	val, ok := m.Get("x")

	// Assert
	actual := args.Map{
		"isNil": val == nil,
		"ok":    ok,
	}
	expected := args.Map{
		"isNil": true,
		"ok":    false,
	}
	expected.ShouldBeEqual(t, 0, "Get returns nil -- nil map", actual)
}

func Test_Map_NilMap_IsKeyInvalid(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	result := m.IsKeyInvalid("x")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsKeyInvalid returns false -- nil map", actual)
}

func Test_Map_NilMap_IsKeyMissing(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	result := m.IsKeyMissing("x")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsKeyMissing returns false -- nil map", actual)
}

func Test_Map_NilMap_HasDefinedAll(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	result := m.HasDefinedAll("x")

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns false -- nil map", actual)
}
