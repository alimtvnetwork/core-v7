package corecmptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════
// args.Map — coverage for Map methods
// ══════════════════════════════════════════════════════════════════

func Test_ArgsMap_Length(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Length", actual)
}

func Test_ArgsMap_ArgsCount_NoFuncNoExpect(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"count": m.ArgsCount()}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- ArgsCount no func no expect", actual)
}

func Test_ArgsMap_Has_Present(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"has": m.Has("key")}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Has present", actual)
}

func Test_ArgsMap_Has_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"has": m.Has("missing")}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Has missing", actual)
}

func Test_ArgsMap_Has_NilMap(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"has": m.Has("key")}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- Has nil map", actual)
}

func Test_ArgsMap_HasDefined_Present(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"has": m.HasDefined("key")}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- HasDefined present", actual)
}

func Test_ArgsMap_HasDefined_NilMap(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"has": m.HasDefined("key")}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- HasDefined nil map", actual)
}

func Test_ArgsMap_HasDefinedAll_Present(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{"has": m.HasDefinedAll("a", "b")}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- HasDefinedAll present", actual)
}

func Test_ArgsMap_HasDefinedAll_NilMap(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"has": m.HasDefinedAll("a")}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- HasDefinedAll nil map", actual)
}

func Test_ArgsMap_HasDefinedAll_NoNames(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{"has": m.HasDefinedAll()}

	// Assert
	expected := args.Map{"has": false}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- HasDefinedAll no names", actual)
}

func Test_ArgsMap_IsKeyInvalid_NilMap(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"invalid": m.IsKeyInvalid("key")}

	// Assert
	expected := args.Map{"invalid": false}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- IsKeyInvalid nil map", actual)
}

func Test_ArgsMap_IsKeyMissing_NilMap(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"missing": m.IsKeyMissing("key")}

	// Assert
	expected := args.Map{"missing": false}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- IsKeyMissing nil map", actual)
}

func Test_ArgsMap_IsKeyMissing_Present(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"missing": m.IsKeyMissing("key")}

	// Assert
	expected := args.Map{"missing": false}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- IsKeyMissing present", actual)
}

func Test_ArgsMap_IsKeyMissing_Absent(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}

	// Act
	actual := args.Map{"missing": m.IsKeyMissing("other")}

	// Assert
	expected := args.Map{"missing": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- IsKeyMissing absent", actual)
}

func Test_ArgsMap_Get_NilMap(t *testing.T) {
	// Arrange
	var m args.Map
	_, isValid := m.Get("key")

	// Act
	actual := args.Map{"isValid": isValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "Map returns nil -- Get nil map", actual)
}

func Test_ArgsMap_Get_Present(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	item, isValid := m.Get("key")

	// Act
	actual := args.Map{
		"item": item,
		"isValid": isValid,
	}

	// Assert
	expected := args.Map{
		"item": "val",
		"isValid": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Get present", actual)
}

func Test_ArgsMap_Get_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	_, isValid := m.Get("missing")

	// Act
	actual := args.Map{"isValid": isValid}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Get missing", actual)
}

func Test_ArgsMap_GetAsInt(t *testing.T) {
	// Arrange
	m := args.Map{"num": 42}
	val, ok := m.GetAsInt("num")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsInt", actual)
}

func Test_ArgsMap_GetAsInt_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}
	val, ok := m.GetAsInt("num")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": 0,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsInt missing", actual)
}

func Test_ArgsMap_GetAsIntDefault(t *testing.T) {
	// Arrange
	m := args.Map{"num": 42}

	// Act
	actual := args.Map{"val": m.GetAsIntDefault("num", 99)}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsIntDefault found", actual)
}

func Test_ArgsMap_GetAsIntDefault_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	actual := args.Map{"val": m.GetAsIntDefault("num", 99)}

	// Assert
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsIntDefault missing", actual)
}

func Test_ArgsMap_GetAsBool(t *testing.T) {
	// Arrange
	m := args.Map{"flag": true}
	val, ok := m.GetAsBool("flag")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsBool", actual)
}

func Test_ArgsMap_GetAsBool_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}
	val, ok := m.GetAsBool("flag")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": false,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsBool missing", actual)
}

func Test_ArgsMap_GetAsBoolDefault(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	actual := args.Map{"val": m.GetAsBoolDefault("flag", true)}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsBoolDefault fallback", actual)
}

func Test_ArgsMap_GetAsString(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}
	val, ok := m.GetAsString("name")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsString", actual)
}

func Test_ArgsMap_GetAsStringDefault(t *testing.T) {
	// Arrange
	m := args.Map{}

	// Act
	actual := args.Map{"val": m.GetAsStringDefault("name")}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- GetAsStringDefault empty", actual)
}

func Test_ArgsMap_GetAsStrings(t *testing.T) {
	// Arrange
	m := args.Map{"items": []string{"a", "b"}}
	items, ok := m.GetAsStrings("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsStrings", actual)
}

func Test_ArgsMap_GetAsStrings_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}
	items, ok := m.GetAsStrings("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsStrings missing", actual)
}

func Test_ArgsMap_GetAsAnyItems(t *testing.T) {
	// Arrange
	m := args.Map{"items": []any{1, "two"}}
	items, ok := m.GetAsAnyItems("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsAnyItems", actual)
}

func Test_ArgsMap_GetAsAnyItems_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}
	items, ok := m.GetAsAnyItems("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"ok": false,
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetAsAnyItems missing", actual)
}

func Test_ArgsMap_GetDirectLower(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}

	// Act
	actual := args.Map{"val": m.GetDirectLower("NAME")}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetDirectLower", actual)
}

func Test_ArgsMap_GetDirectLower_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}
	isNil := m.GetDirectLower("MISSING") == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetDirectLower missing", actual)
}

func Test_ArgsMap_When(t *testing.T) {
	// Arrange
	m := args.Map{"when": "now"}

	// Act
	actual := args.Map{"val": m.When()}

	// Assert
	expected := args.Map{"val": "now"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- When", actual)
}

func Test_ArgsMap_Title(t *testing.T) {
	// Arrange
	m := args.Map{"title": "test"}

	// Act
	actual := args.Map{"val": m.Title()}

	// Assert
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Title", actual)
}

func Test_ArgsMap_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual("result")

	// Act
	actual := args.Map{"val": m.Actual()}

	// Assert
	expected := args.Map{"val": "result"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SetActual", actual)
}

func Test_ArgsMap_GetFirstOfNames_Found(t *testing.T) {
	// Arrange
	m := args.Map{"p2": "val"}

	// Act
	actual := args.Map{"val": m.GetFirstOfNames("p1", "p2", "p3")}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetFirstOfNames found", actual)
}

func Test_ArgsMap_GetFirstOfNames_Empty(t *testing.T) {
	// Arrange
	m := args.Map{"x": "val"}
	isNil := m.GetFirstOfNames() == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- GetFirstOfNames empty", actual)
}

func Test_ArgsMap_SortedKeys_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{
		"len": len(keys),
		"isNil": err == nil,
	}

	// Assert
	expected := args.Map{
		"len": 0,
		"isNil": true,
	}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- SortedKeys empty", actual)
}

func Test_ArgsMap_SortedKeys_NonEmpty(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, _ := m.SortedKeys()

	// Act
	actual := args.Map{
		"first": keys[0],
		"second": keys[1],
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SortedKeys sorted", actual)
}

func Test_ArgsMap_CompileToStrings(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	lines := m.CompileToStrings()

	// Act
	actual := args.Map{
		"len": len(lines),
		"first": lines[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a : 1",
	}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- CompileToStrings", actual)
}

func Test_ArgsMap_CompileToStrings_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	lines := m.CompileToStrings()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- CompileToStrings empty", actual)
}

func Test_ArgsMap_CompileToString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.CompileToString()

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "a : 1"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- CompileToString", actual)
}

func Test_ArgsMap_GoLiteralLines(t *testing.T) {
	// Arrange
	m := args.Map{"name": "test"}
	lines := m.GoLiteralLines()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GoLiteralLines", actual)
}

func Test_ArgsMap_GoLiteralLines_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	lines := m.GoLiteralLines()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- GoLiteralLines empty", actual)
}

func Test_ArgsMap_GoLiteralString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.GoLiteralString()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GoLiteralString", actual)
}

func Test_ArgsMap_GetByIndex(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	result := m.GetByIndex(0)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Map returns non-empty -- GetByIndex valid", actual)
}

func Test_ArgsMap_GetByIndex_OOB(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.GetByIndex(10)

	// Act
	actual := args.Map{"isNil": result == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- GetByIndex OOB", actual)
}

func Test_ArgsMap_Slice(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	slice := m.Slice()

	// Act
	actual := args.Map{"len": len(slice)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Slice", actual)
}

func Test_ArgsMap_String_NonEmpty(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	result := m.String()

	// Act
	actual := args.Map{"hasMap": strings.Contains(result, "Map")}

	// Assert
	expected := args.Map{"hasMap": true}
	expected.ShouldBeEqual(t, 0, "Map returns empty -- String non-empty", actual)
}

func Test_ArgsMap_Raw(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	raw := m.Raw()

	// Act
	actual := args.Map{"len": len(raw)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Raw", actual)
}

func Test_ArgsMap_Args(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	result := m.Args("a", "b")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Args", actual)
}

func Test_ArgsMap_Expect(t *testing.T) {
	// Arrange
	m := args.Map{"expect": "val"}

	// Act
	actual := args.Map{"val": m.Expect()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Expect", actual)
}

func Test_ArgsMap_Arrange(t *testing.T) {
	// Arrange
	m := args.Map{"arrange": "data"}

	// Act
	actual := args.Map{"val": m.Arrange()}

	// Assert
	expected := args.Map{"val": "data"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Arrange", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.One — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsOne_All(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}

	// Act
	actual := args.Map{
		"first":     one.FirstItem(),
		"expected":  one.Expected(),
		"hasFirst":  one.HasFirst(),
		"hasExpect": one.HasExpect(),
		"argsCount": one.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first":     "hello",
		"expected":  42,
		"hasFirst":  true,
		"hasExpect": true,
		"argsCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- all methods", actual)
}

func Test_ArgsOne_Slice(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: 42}
	slice := one.Slice()

	// Act
	actual := args.Map{"len": len(slice)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- Slice", actual)
}

func Test_ArgsOne_SliceCached(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}
	_ = one.Slice()
	slice := one.Slice() // cached

	// Act
	actual := args.Map{"len": len(slice)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- Slice cached", actual)
}

func Test_ArgsOne_String(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}
	result := one.String()

	// Act
	actual := args.Map{"hasOne": strings.Contains(result, "One")}

	// Assert
	expected := args.Map{"hasOne": true}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- String", actual)
}

func Test_ArgsOne_GetByIndex(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}

	// Act
	actual := args.Map{"val": one.GetByIndex(0)}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- GetByIndex", actual)
}

func Test_ArgsOne_GetByIndex_OOB(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}
	isNil := one.GetByIndex(10) == nil

	// Act
	actual := args.Map{"isNil": isNil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- GetByIndex OOB", actual)
}

func Test_ArgsOne_Args(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}
	result := one.Args(1)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- Args upTo 1", actual)
}

func Test_ArgsOne_Args_Zero(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}
	result := one.Args(0)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- Args upTo 0", actual)
}

func Test_ArgsOne_ValidArgs(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello"}
	result := one.ValidArgs()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One returns non-empty -- ValidArgs", actual)
}

func Test_ArgsOne_LeftRight(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: "exp"}
	lr := one.LeftRight()

	// Act
	actual := args.Map{
		"left": lr.Left,
		"expect": lr.Expect,
	}

	// Assert
	expected := args.Map{
		"left": "hello",
		"expect": "exp",
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- LeftRight", actual)
}

func Test_ArgsOne_ArgTwo(t *testing.T) {
	// Arrange
	one := args.OneAny{First: "hello", Expect: "exp"}
	two := one.ArgTwo()

	// Act
	actual := args.Map{
		"first": two.First,
		"expect": two.Expect,
	}

	// Assert
	expected := args.Map{
		"first": "hello",
		"expect": "exp",
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- ArgTwo", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Two — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsTwo_All(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b", Expect: 1}

	// Act
	actual := args.Map{
		"first":     two.FirstItem(),
		"second":    two.SecondItem(),
		"expected":  two.Expected(),
		"hasFirst":  two.HasFirst(),
		"hasSecond": two.HasSecond(),
		"argsCount": two.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "expected": 1,
		"hasFirst": true, "hasSecond": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- all methods", actual)
}

func Test_ArgsTwo_Slice(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{"len": len(two.Slice())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- Slice", actual)
}

func Test_ArgsTwo_String(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}
	result := two.String()

	// Act
	actual := args.Map{"hasTwo": strings.Contains(result, "Two")}

	// Assert
	expected := args.Map{"hasTwo": true}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- String", actual)
}

func Test_ArgsTwo_Args(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{
		"args0": len(two.Args(0)),
		"args1": len(two.Args(1)),
		"args2": len(two.Args(2)),
	}

	// Assert
	expected := args.Map{
		"args0": 0,
		"args1": 1,
		"args2": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- Args", actual)
}

func Test_ArgsTwo_LeftRight(t *testing.T) {
	// Arrange
	two := args.TwoAny{First: "a", Second: "b"}
	lr := two.LeftRight()

	// Act
	actual := args.Map{
		"left": lr.Left,
		"right": lr.Right,
	}

	// Assert
	expected := args.Map{
		"left": "a",
		"right": "b",
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Three — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsThree_All(t *testing.T) {
	// Arrange
	three := args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: 1}

	// Act
	actual := args.Map{
		"first":     three.FirstItem(),
		"second":    three.SecondItem(),
		"third":     three.ThirdItem(),
		"hasThird":  three.HasThird(),
		"argsCount": three.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"hasThird": true, "argsCount": 3,
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- all methods", actual)
}

func Test_ArgsThree_ArgTwo(t *testing.T) {
	// Arrange
	three := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	two := three.ArgTwo()

	// Act
	actual := args.Map{
		"first": two.First,
		"second": two.Second,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- ArgTwo", actual)
}

func Test_ArgsThree_ArgThree(t *testing.T) {
	// Arrange
	three := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	copy := three.ArgThree()

	// Act
	actual := args.Map{"third": copy.Third}

	// Assert
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- ArgThree", actual)
}

func Test_ArgsThree_Args(t *testing.T) {
	// Arrange
	three := args.ThreeAny{First: "a", Second: "b", Third: "c"}

	// Act
	actual := args.Map{"args3": len(three.Args(3))}

	// Assert
	expected := args.Map{"args3": 3}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- Args", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Four — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsFour_All(t *testing.T) {
	// Arrange
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}

	// Act
	actual := args.Map{
		"fourth":    four.FourthItem(),
		"hasFourth": four.HasFourth(),
		"argsCount": four.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"fourth": "d",
		"hasFourth": true,
		"argsCount": 4,
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- all methods", actual)
}

func Test_ArgsFour_Args(t *testing.T) {
	// Arrange
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}

	// Act
	actual := args.Map{"args4": len(four.Args(4))}

	// Assert
	expected := args.Map{"args4": 4}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- Args", actual)
}

func Test_ArgsFour_ArgTwo(t *testing.T) {
	// Arrange
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	two := four.ArgTwo()

	// Act
	actual := args.Map{"first": two.First}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- ArgTwo", actual)
}

func Test_ArgsFour_ArgThree(t *testing.T) {
	// Arrange
	four := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	three := four.ArgThree()

	// Act
	actual := args.Map{"third": three.Third}

	// Assert
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- ArgThree", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Five — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsFive_All(t *testing.T) {
	// Arrange
	five := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}

	// Act
	actual := args.Map{
		"fifth":    five.FifthItem(),
		"hasFifth": five.HasFifth(),
		"count":    five.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"fifth": "e",
		"hasFifth": true,
		"count": 5,
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- all methods", actual)
}

func Test_ArgsFive_Args(t *testing.T) {
	// Arrange
	five := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}

	// Act
	actual := args.Map{"args5": len(five.Args(5))}

	// Assert
	expected := args.Map{"args5": 5}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- Args", actual)
}

func Test_ArgsFive_ArgFour(t *testing.T) {
	// Arrange
	five := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	four := five.ArgFour()

	// Act
	actual := args.Map{"fourth": four.Fourth}

	// Assert
	expected := args.Map{"fourth": "d"}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- ArgFour", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Six — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsSix_All(t *testing.T) {
	// Arrange
	six := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}

	// Act
	actual := args.Map{
		"sixth":    six.SixthItem(),
		"hasSixth": six.HasSixth(),
		"count":    six.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"sixth": "f",
		"hasSixth": true,
		"count": 6,
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- all methods", actual)
}

func Test_ArgsSix_Args(t *testing.T) {
	// Arrange
	six := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}

	// Act
	actual := args.Map{"args6": len(six.Args(6))}

	// Assert
	expected := args.Map{"args6": 6}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- Args", actual)
}

func Test_ArgsSix_ArgFive(t *testing.T) {
	// Arrange
	six := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	five := six.ArgFive()

	// Act
	actual := args.Map{"fifth": five.Fifth}

	// Assert
	expected := args.Map{"fifth": "e"}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- ArgFive", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.LeftRight — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsLeftRight_All(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b", Expect: 1}

	// Act
	actual := args.Map{
		"left":      lr.FirstItem(),
		"right":     lr.SecondItem(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasFirst":  lr.HasFirst(),
		"hasSecond": lr.HasSecond(),
		"count":     lr.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"left": "a", "right": "b",
		"hasLeft": true, "hasRight": true,
		"hasFirst": true, "hasSecond": true,
		"count": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- all methods", actual)
}

func Test_ArgsLeftRight_Clone(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "a", Right: "b"}
	cloned := lr.Clone()

	// Act
	actual := args.Map{
		"left": cloned.Left,
		"right": cloned.Right,
	}

	// Assert
	expected := args.Map{
		"left": "a",
		"right": "b",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clone", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.String — coverage
// ══════════════════════════════════════════════════════════════════

func Test_ArgsString_Methods(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{
		"string":    s.String(),
		"length":    s.Length(),
		"count":     s.Count(),
		"isEmpty":   s.IsEmpty(),
		"isDefined": s.IsDefined(),
		"hasCh":     s.HasCharacter(),
		"asciiLen":  s.AscIILength(),
	}

	// Assert
	expected := args.Map{
		"string": "hello", "length": 5, "count": 5,
		"isEmpty": false, "isDefined": true, "hasCh": true,
		"asciiLen": 5,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- basic methods", actual)
}

func Test_ArgsString_IsEmptyTrue(t *testing.T) {
	// Arrange
	s := args.String("")

	// Act
	actual := args.Map{
		"isEmpty": s.IsEmpty(),
		"isEW": s.IsEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"isEmpty": true,
		"isEW": true,
	}
	expected.ShouldBeEqual(t, 0, "String returns empty -- empty", actual)
}

func Test_ArgsString_TrimSpace(t *testing.T) {
	// Arrange
	s := args.String("  hello  ")

	// Act
	actual := args.Map{"val": s.TrimSpace().String()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- TrimSpace", actual)
}

func Test_ArgsString_ReplaceAll(t *testing.T) {
	// Arrange
	s := args.String("hello world")

	// Act
	actual := args.Map{"val": s.ReplaceAll("world", "go").String()}

	// Assert
	expected := args.Map{"val": "hello go"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- ReplaceAll", actual)
}

func Test_ArgsString_Concat(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{"val": s.Concat(" ", "world").String()}

	// Assert
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Concat", actual)
}

func Test_ArgsString_Join(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{"val": s.Join(",", "a", "b").String()}

	// Assert
	expected := args.Map{"val": "hello,a,b"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Join", actual)
}

func Test_ArgsString_Split(t *testing.T) {
	// Arrange
	s := args.String("a,b,c")
	result := s.Split(",")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Split", actual)
}

func Test_ArgsString_Bytes(t *testing.T) {
	// Arrange
	s := args.String("hi")

	// Act
	actual := args.Map{"len": len(s.Bytes())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Bytes", actual)
}

func Test_ArgsString_Runes(t *testing.T) {
	// Arrange
	s := args.String("hi")

	// Act
	actual := args.Map{"len": len(s.Runes())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Runes", actual)
}

func Test_ArgsString_Substring(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{"val": s.Substring(1, 4).String()}

	// Assert
	expected := args.Map{"val": "ell"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Substring", actual)
}

func Test_ArgsString_DoubleQuote(t *testing.T) {
	// Arrange
	s := args.String("hi")
	result := s.DoubleQuote().String()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- DoubleQuote", actual)
}

func Test_ArgsString_DoubleQuoteQ(t *testing.T) {
	// Arrange
	s := args.String("hi")
	result := s.DoubleQuoteQ().String()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- DoubleQuoteQ", actual)
}

func Test_ArgsString_SingleQuote(t *testing.T) {
	// Arrange
	s := args.String("hi")
	result := s.SingleQuote().String()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- SingleQuote", actual)
}

func Test_ArgsString_ValueDoubleQuote(t *testing.T) {
	// Arrange
	s := args.String("hi")
	result := s.ValueDoubleQuote().String()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- ValueDoubleQuote", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncWrap — coverage
// ══════════════════════════════════════════════════════════════════

func sampleFunc(s string) int { return len(s) }

func Test_NewFuncWrap_Default(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{
		"isValid":   fw.IsValid(),
		"argsCount": fw.ArgsCount(),
		"retLen":    fw.ReturnLength(),
	}

	// Assert
	expected := args.Map{
		"isValid": true,
		"argsCount": 1,
		"retLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- Default", actual)
}

func Test_NewFuncWrap_Default_Nil(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{"isValid": fw.IsValid()}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- Default nil", actual)
}

func Test_NewFuncWrap_Default_NotFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default("not a func")

	// Act
	actual := args.Map{
		"isValid": fw.IsValid(),
		"isInvalid": fw.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- Default not func", actual)
}

func Test_NewFuncWrap_Invalid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Invalid()

	// Act
	actual := args.Map{"isInvalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- Invalid", actual)
}

func Test_FuncWrap_Invoke(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	results, err := fw.Invoke("hello")

	// Act
	actual := args.Map{
		"result": results[0],
		"err": err == nil,
	}

	// Assert
	expected := args.Map{
		"result": 5,
		"err": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- Invoke", actual)
}

func Test_FuncWrap_InvokeMust(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	results := fw.InvokeMust("hi")

	// Act
	actual := args.Map{"result": results[0]}

	// Assert
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeMust", actual)
}

func Test_FuncWrap_GetFuncName(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"name": fw.GetFuncName()}

	// Assert
	expected := args.Map{"name": "sampleFunc"}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetFuncName", actual)
}

func Test_FuncWrap_GetFuncName_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{"name": fw.GetFuncName()}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- GetFuncName nil", actual)
}

func Test_FuncWrap_HasValidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"valid": fw.HasValidFunc()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns non-empty -- HasValidFunc", actual)
}

func Test_FuncWrap_ValidationError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"isNil": fw.ValidationError() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- ValidationError valid", actual)
}

func Test_FuncWrap_ValidationError_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{"hasErr": fw.ValidationError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- ValidationError nil", actual)
}

func Test_FuncWrap_ValidationError_Invalid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default("not a func")

	// Act
	actual := args.Map{"hasErr": fw.ValidationError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- ValidationError invalid", actual)
}

func Test_FuncWrap_InvalidError_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{"hasErr": fw.InvalidError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- InvalidError nil", actual)
}

func Test_FuncWrap_InvalidError_Valid(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"isNil": fw.InvalidError() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns error -- InvalidError valid", actual)
}

func Test_FuncWrap_GetInArgsTypes(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	types := fw.GetInArgsTypes()

	// Act
	actual := args.Map{"len": len(types)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetInArgsTypes", actual)
}

func Test_FuncWrap_GetOutArgsTypes(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	types := fw.GetOutArgsTypes()

	// Act
	actual := args.Map{"len": len(types)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetOutArgsTypes", actual)
}

func Test_FuncWrap_GetInArgsTypesNames(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.GetInArgsTypesNames()

	// Act
	actual := args.Map{
		"len": len(names),
		"first": names[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "string",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetInArgsTypesNames", actual)
}

func Test_FuncWrap_GetOutArgsTypesNames(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.GetOutArgsTypesNames()

	// Act
	actual := args.Map{
		"len": len(names),
		"first": names[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "int",
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetOutArgsTypesNames", actual)
}

func Test_FuncWrap_InArgNames(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.InArgNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InArgNames", actual)
}

func Test_FuncWrap_OutArgNames(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	names := fw.OutArgNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- OutArgNames", actual)
}

func Test_FuncWrap_IsStringFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "" })

	// Act
	actual := args.Map{"isString": fw.IsStringFunc()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsStringFunc", actual)
}

func Test_FuncWrap_IsBoolFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() bool { return false })

	// Act
	actual := args.Map{"isBool": fw.IsBoolFunc()}

	// Assert
	expected := args.Map{"isBool": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsBoolFunc", actual)
}

func Test_FuncWrap_IsVoidFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() {})

	// Act
	actual := args.Map{"isVoid": fw.IsVoidFunc()}

	// Assert
	expected := args.Map{"isVoid": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsVoidFunc", actual)
}

func Test_FuncWrap_IsAnyFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"isAny": fw.IsAnyFunc()}

	// Assert
	expected := args.Map{"isAny": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsAnyFunc", actual)
}

func Test_FuncWrap_InvokeAsBool(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() bool { return true })
	val, err := fw.InvokeAsBool()

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeAsBool", actual)
}

func Test_FuncWrap_InvokeAsString(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() string { return "hi" })
	val, err := fw.InvokeAsString()

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": "hi",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeAsString", actual)
}

func Test_FuncWrap_InvokeAsAny(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(func() int { return 42 })
	val, err := fw.InvokeAsAny()

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- InvokeAsAny", actual)
}

func Test_FuncWrap_VoidCall(t *testing.T) {
	// Arrange
	called := false
	fw := args.NewFuncWrap.Default(func() { called = true })
	_, err := fw.VoidCall()

	// Act
	actual := args.Map{
		"called": called,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- VoidCall", actual)
}

func Test_FuncWrap_IsEqual_Same(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"equal": fw.IsEqual(fw)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsEqual same", actual)
}

func Test_FuncWrap_IsNotEqual(t *testing.T) {
	// Arrange
	fw1 := args.NewFuncWrap.Default(sampleFunc)
	fw2 := args.NewFuncWrap.Default(func() {})

	// Act
	actual := args.Map{"notEqual": fw1.IsNotEqual(fw2)}

	// Assert
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsNotEqual", actual)
}

func Test_FuncWrap_PkgPath(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	path := fw.PkgPath()

	// Act
	actual := args.Map{"hasPath": len(path) > 0}

	// Assert
	expected := args.Map{"hasPath": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- PkgPath", actual)
}

func Test_FuncWrap_PkgNameOnly(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	name := fw.PkgNameOnly()

	// Act
	actual := args.Map{"hasName": len(name) > 0}

	// Assert
	expected := args.Map{"hasName": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- PkgNameOnly", actual)
}

func Test_FuncWrap_FuncDirectInvokeName(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)
	name := fw.FuncDirectInvokeName()

	// Act
	actual := args.Map{"hasName": len(name) > 0}

	// Assert
	expected := args.Map{"hasName": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- FuncDirectInvokeName", actual)
}

func Test_FuncWrap_GetType(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"notNil": fw.GetType() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetType", actual)
}

func Test_FuncWrap_IsInTypeMatches(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"matches": fw.IsInTypeMatches("hello")}

	// Assert
	expected := args.Map{"matches": true}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- IsInTypeMatches", actual)
}

func Test_FuncWrap_GetPascalCaseFuncName(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(sampleFunc)

	// Act
	actual := args.Map{"name": fw.GetPascalCaseFuncName()}

	// Assert
	expected := args.Map{"name": "SampleFunc"}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- GetPascalCaseFuncName", actual)
}

func Test_FuncWrap_GetPascalCaseFuncName_Nil(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny

	// Act
	actual := args.Map{"name": fw.GetPascalCaseFuncName()}

	// Assert
	expected := args.Map{"name": ""}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns nil -- GetPascalCaseFuncName nil", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.NewFuncWrap — Map, Many, Single
// ══════════════════════════════════════════════════════════════════

func Test_NewFuncWrap_Map(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"len": fm.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap returns correct value -- Map", actual)
}

func Test_NewFuncWrap_Map_Empty(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map()

	// Act
	actual := args.Map{"len": fm.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap returns empty -- Map empty", actual)
}

func Test_NewFuncWrap_Many(t *testing.T) {
	// Arrange
	wraps := args.NewFuncWrap.Many(sampleFunc)

	// Act
	actual := args.Map{"len": len(wraps)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap returns correct value -- Many", actual)
}

func Test_NewFuncWrap_Many_Empty(t *testing.T) {
	// Arrange
	wraps := args.NewFuncWrap.Many()

	// Act
	actual := args.Map{"len": len(wraps)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap returns empty -- Many empty", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Empty — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Empty_Map(t *testing.T) {
	// Arrange
	m := args.Empty.Map()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- Map", actual)
}

func Test_Empty_FuncWrap(t *testing.T) {
	// Arrange
	fw := args.Empty.FuncWrap()

	// Act
	actual := args.Map{"isInvalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"isInvalid": true}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- FuncWrap", actual)
}

func Test_Empty_FuncMap(t *testing.T) {
	// Arrange
	fm := args.Empty.FuncMap()

	// Act
	actual := args.Map{"isEmpty": fm.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- FuncMap", actual)
}

func Test_Empty_Holder(t *testing.T) {
	// Arrange
	h := args.Empty.Holder()

	// Act
	actual := args.Map{"count": h.ArgsCount()}

	// Assert
	expected := args.Map{"count": 7}
	expected.ShouldBeEqual(t, 0, "Empty returns empty -- Holder", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncMap — coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncMap_Has(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{
		"has":        fm.Has("sampleFunc"),
		"hasMissing": fm.Has("missing"),
		"isContains": fm.IsContains("sampleFunc"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"hasMissing": false,
		"isContains": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Has", actual)
}

func Test_FuncMap_Get(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)
	f := fm.Get("sampleFunc")

	// Act
	actual := args.Map{"notNil": f != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Get", actual)
}

func Test_FuncMap_Get_Missing(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)
	f := fm.Get("missing")

	// Act
	actual := args.Map{"isNil": f == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Get missing", actual)
}

func Test_FuncMap_IsValidFuncOf(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"valid": fm.IsValidFuncOf("sampleFunc")}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns non-empty -- IsValidFuncOf", actual)
}

func Test_FuncMap_IsInvalidFunc(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"invalid": fm.IsInvalidFunc("missing")}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- IsInvalidFunc missing", actual)
}

func Test_FuncMap_HasAnyItem(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"hasAny": fm.HasAnyItem()}

	// Assert
	expected := args.Map{"hasAny": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- HasAnyItem", actual)
}

func Test_FuncMap_ArgsCount(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"count": fm.ArgsCount("sampleFunc")}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- ArgsCount", actual)
}

func Test_FuncMap_ArgsCount_Missing(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"count": fm.ArgsCount("missing")}

	// Assert
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- ArgsCount missing", actual)
}

func Test_FuncMap_ReturnLength(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"retLen": fm.ReturnLength("sampleFunc")}

	// Assert
	expected := args.Map{"retLen": 1}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- ReturnLength", actual)
}

func Test_FuncMap_ReturnLength_Missing(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"retLen": fm.ReturnLength("missing")}

	// Assert
	expected := args.Map{"retLen": 0}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- ReturnLength missing", actual)
}

func Test_FuncMap_InvalidError_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	actual := args.Map{"hasErr": fm.InvalidError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns empty -- InvalidError empty", actual)
}

func Test_FuncMap_InvalidError_NonEmpty(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(sampleFunc)

	// Act
	actual := args.Map{"isNil": fm.InvalidError() == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns empty -- InvalidError non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncDetector — coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncDetector_GetFuncWrap_Func(t *testing.T) {
	// Arrange
	fw := args.FuncDetector.GetFuncWrap(sampleFunc)

	// Act
	actual := args.Map{"isValid": fw.IsValid()}

	// Assert
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector returns correct value -- from func", actual)
}

func Test_FuncDetector_GetFuncWrap_FuncWrap(t *testing.T) {
	// Arrange
	original := args.NewFuncWrap.Default(sampleFunc)
	fw := args.FuncDetector.GetFuncWrap(original)

	// Act
	actual := args.Map{"same": fw == original}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector returns correct value -- from FuncWrap", actual)
}

func Test_FuncDetector_GetFuncWrap_Map(t *testing.T) {
	// Arrange
	m := args.Map{"func": sampleFunc}
	fw := args.FuncDetector.GetFuncWrap(m)

	// Act
	actual := args.Map{"isValid": fw.IsValid()}

	// Assert
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector returns correct value -- from Map", actual)
}
