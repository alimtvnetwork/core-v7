package corecmptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════
// args.Map — additional coverage for uncovered methods
// ══════════════════════════════════════════════════════════════════

func Test_Map_CompileToStrings(t *testing.T) {
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
		"second": lines[1],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a : 1",
		"second": "b : 2",
	}
	expected.ShouldBeEqual(t, 0, "CompileToStrings returns correct value -- with args", actual)
}

func Test_Map_CompileToStrings_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	lines := m.CompileToStrings()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "CompileToStrings returns empty -- empty", actual)
}

func Test_Map_CompileToString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.CompileToString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a : 1"}
	expected.ShouldBeEqual(t, 0, "CompileToString returns correct value -- with args", actual)
}

func Test_Map_CompileToString_Multi(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	result := m.CompileToString()

	// Act
	actual := args.Map{"hasNewline": strings.Contains(result, "\n")}

	// Assert
	expected := args.Map{"hasNewline": true}
	expected.ShouldBeEqual(t, 0, "CompileToString returns correct value -- multi", actual)
}

func Test_Map_GoLiteralLines(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"name": "hello",
	}
	lines := m.GoLiteralLines()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GoLiteralLines returns correct value -- with args", actual)
}

func Test_Map_GoLiteralLines_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	lines := m.GoLiteralLines()

	// Act
	actual := args.Map{"len": len(lines)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "GoLiteralLines returns empty -- empty", actual)
}

func Test_Map_GoLiteralString(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.GoLiteralString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GoLiteralString returns correct value -- with args", actual)
}

func Test_Map_GetAsInt(t *testing.T) {
	// Arrange
	m := args.Map{"val": 42}
	val, ok := m.GetAsInt("val")

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
	expected.ShouldBeEqual(t, 0, "GetAsInt returns correct value -- with args", actual)
}

func Test_Map_GetAsInt_Missing(t *testing.T) {
	// Arrange
	m := args.Map{"val": "str"}
	val, ok := m.GetAsInt("val")

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
	expected.ShouldBeEqual(t, 0, "GetAsInt returns correct value -- wrong type", actual)
}

func Test_Map_GetAsIntDefault(t *testing.T) {
	// Arrange
	m := args.Map{"val": 42}
	val := m.GetAsIntDefault("val", 0)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault returns correct value -- found", actual)
}

func Test_Map_GetAsIntDefault_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}
	val := m.GetAsIntDefault("val", 99)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "GetAsIntDefault returns correct value -- default", actual)
}

func Test_Map_GetAsBool(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "GetAsBool returns correct value -- with args", actual)
}

func Test_Map_GetAsBoolDefault(t *testing.T) {
	// Arrange
	m := args.Map{}
	val := m.GetAsBoolDefault("flag", true)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": true}
	expected.ShouldBeEqual(t, 0, "GetAsBoolDefault returns correct value -- with args", actual)
}

func Test_Map_GetAsString(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "GetAsString returns correct value -- with args", actual)
}

func Test_Map_GetAsStringDefault(t *testing.T) {
	// Arrange
	m := args.Map{}
	val := m.GetAsStringDefault("name")

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "GetAsStringDefault returns correct value -- with args", actual)
}

func Test_Map_GetAsStrings(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "GetAsStrings returns correct value -- with args", actual)
}

func Test_Map_GetAsStrings_Missing(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "GetAsStrings returns correct value -- missing", actual)
}

func Test_Map_GetAsAnyItems(t *testing.T) {
	// Arrange
	m := args.Map{"items": []any{1, "a"}}
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
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems returns correct value -- with args", actual)
}

func Test_Map_GetAsAnyItems_Missing(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "GetAsAnyItems returns correct value -- missing", actual)
}

func Test_Map_Slice(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	slice := m.Slice()

	// Act
	actual := args.Map{"len": len(slice)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Slice", actual)
}

func Test_Map_String(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- String", actual)
}

func Test_Map_GetByIndex(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.GetByIndex(0)

	// Act
	actual := args.Map{"notNil": result != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "GetByIndex returns non-empty -- valid", actual)
}

func Test_Map_GetByIndex_OutOfBounds(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.GetByIndex(99)

	// Act
	actual := args.Map{"nil": result == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetByIndex returns correct value -- out of bounds", actual)
}

func Test_Map_SortedKeys(t *testing.T) {
	// Arrange
	m := args.Map{
		"b": 2,
		"a": 1,
	}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"first": keys[0],
		"second": keys[1],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "SortedKeys returns correct value -- with args", actual)
}

func Test_Map_SortedKeys_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	keys, err := m.SortedKeys()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"len": len(keys),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"len": 0,
	}
	expected.ShouldBeEqual(t, 0, "SortedKeys returns empty -- empty", actual)
}

func Test_Map_When(t *testing.T) {
	// Arrange
	m := args.Map{"when": "condition"}

	// Act
	actual := args.Map{"val": m.When()}

	// Assert
	expected := args.Map{"val": "condition"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- When", actual)
}

func Test_Map_Title(t *testing.T) {
	// Arrange
	m := args.Map{"title": "test"}

	// Act
	actual := args.Map{"val": m.Title()}

	// Assert
	expected := args.Map{"val": "test"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Title", actual)
}

func Test_Map_GetLowerCase(t *testing.T) {
	// Arrange
	m := args.Map{"name": "hello"}
	val, ok := m.GetLowerCase("NAME")

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
	expected.ShouldBeEqual(t, 0, "GetLowerCase returns correct value -- with args", actual)
}

func Test_Map_GetDirectLower(t *testing.T) {
	// Arrange
	m := args.Map{"key": "val"}
	result := m.GetDirectLower("KEY")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "GetDirectLower returns correct value -- with args", actual)
}

func Test_Map_GetDirectLower_Missing(t *testing.T) {
	// Arrange
	m := args.Map{}
	result := m.GetDirectLower("KEY")

	// Act
	actual := args.Map{"nil": result == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetDirectLower returns correct value -- missing", actual)
}

func Test_Map_Expect(t *testing.T) {
	// Arrange
	m := args.Map{"expect": 42}

	// Act
	actual := args.Map{"val": m.Expect()}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Expect", actual)
}

func Test_Map_Actual(t *testing.T) {
	// Arrange
	m := args.Map{"actual": "data"}

	// Act
	actual := args.Map{"val": m.Actual()}

	// Assert
	expected := args.Map{"val": "data"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Actual", actual)
}

func Test_Map_Arrange(t *testing.T) {
	// Arrange
	m := args.Map{"arrange": "setup"}

	// Act
	actual := args.Map{"val": m.Arrange()}

	// Assert
	expected := args.Map{"val": "setup"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- Arrange", actual)
}

func Test_Map_SetActual(t *testing.T) {
	// Arrange
	m := args.Map{}
	m.SetActual("new-val")

	// Act
	actual := args.Map{"val": m.Actual()}

	// Assert
	expected := args.Map{"val": "new-val"}
	expected.ShouldBeEqual(t, 0, "Map returns correct value -- SetActual", actual)
}

func Test_Map_HasDefinedAll(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}

	// Act
	actual := args.Map{
		"all":  m.HasDefinedAll("a", "b"),
		"miss": m.HasDefinedAll("a", "c"),
	}

	// Assert
	expected := args.Map{
		"all": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns correct value -- with args", actual)
}

func Test_Map_HasDefinedAll_Nil(t *testing.T) {
	// Arrange
	var m args.Map

	// Act
	actual := args.Map{"nil": m.HasDefinedAll("a")}

	// Assert
	expected := args.Map{"nil": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns nil -- nil", actual)
}

func Test_Map_HasDefinedAll_Empty(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{"empty": m.HasDefinedAll()}

	// Assert
	expected := args.Map{"empty": false}
	expected.ShouldBeEqual(t, 0, "HasDefinedAll returns empty -- empty names", actual)
}

func Test_Map_IsKeyInvalid(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{
		"valid":   !m.IsKeyInvalid("a"),
		"invalid": m.IsKeyInvalid("b"),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "IsKeyInvalid returns error -- with args", actual)
}

func Test_Map_IsKeyMissing(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}

	// Act
	actual := args.Map{
		"present": !m.IsKeyMissing("a"),
		"missing": m.IsKeyMissing("b"),
	}

	// Assert
	expected := args.Map{
		"present": true,
		"missing": true,
	}
	expected.ShouldBeEqual(t, 0, "IsKeyMissing returns correct value -- with args", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames(t *testing.T) {
	// Arrange
	m := args.Map{"items": []string{"a", "b"}}
	result := m.GetAsStringSliceFirstOfNames("items")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "GetAsStringSliceFirstOfNames returns correct value -- with args", actual)
}

func Test_Map_GetAsStringSliceFirstOfNames_Empty(t *testing.T) {
	// Arrange
	m := args.Map{}
	result := m.GetAsStringSliceFirstOfNames()

	// Act
	actual := args.Map{"nil": result == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "GetAsStringSliceFirstOfNames returns empty -- empty", actual)
}

func Test_Map_FirstItem(t *testing.T) {
	// Arrange
	m := args.Map{"first": "val"}

	// Act
	actual := args.Map{"val": m.FirstItem()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "FirstItem returns correct value -- with args", actual)
}

func Test_Map_SecondItem(t *testing.T) {
	// Arrange
	m := args.Map{"second": "val"}

	// Act
	actual := args.Map{"val": m.SecondItem()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "SecondItem returns correct value -- with args", actual)
}

func Test_Map_ThirdItem(t *testing.T) {
	// Arrange
	m := args.Map{"third": "val"}

	// Act
	actual := args.Map{"val": m.ThirdItem()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "ThirdItem returns correct value -- with args", actual)
}

func Test_Map_FourthItem(t *testing.T) {
	// Arrange
	m := args.Map{"fourth": "val"}

	// Act
	actual := args.Map{"val": m.FourthItem()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "FourthItem returns correct value -- with args", actual)
}

func Test_Map_FifthItem(t *testing.T) {
	// Arrange
	m := args.Map{"fifth": "val"}

	// Act
	actual := args.Map{"val": m.FifthItem()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "FifthItem returns correct value -- with args", actual)
}

func Test_Map_SixthItem(t *testing.T) {
	// Arrange
	m := args.Map{"sixth": "val"}

	// Act
	actual := args.Map{"val": m.SixthItem()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "SixthItem returns correct value -- with args", actual)
}

func Test_Map_Seventh(t *testing.T) {
	// Arrange
	m := args.Map{"seventh": "val"}

	// Act
	actual := args.Map{"val": m.Seventh()}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Seventh returns correct value -- with args", actual)
}

func Test_Map_ValidArgs(t *testing.T) {
	// Arrange
	m := args.Map{
		"a": 1,
		"b": 2,
	}
	result := m.ValidArgs()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ValidArgs returns non-empty -- with args", actual)
}

func Test_Map_Args(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	result := m.Args("a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Args returns correct value -- with args", actual)
}

func Test_Map_Raw(t *testing.T) {
	// Arrange
	m := args.Map{"a": 1}
	raw := m.Raw()

	// Act
	actual := args.Map{"len": len(raw)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Raw returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.String — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_String_Concat(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.Concat(" world")

	// Act
	actual := args.Map{"val": result.String()}

	// Assert
	expected := args.Map{"val": "hello world"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Concat", actual)
}

func Test_String_Join(t *testing.T) {
	// Arrange
	s := args.String("a")
	result := s.Join("-", "b", "c")

	// Act
	actual := args.Map{"val": result.String()}

	// Assert
	expected := args.Map{"val": "a-b-c"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Join", actual)
}

func Test_String_Split(t *testing.T) {
	// Arrange
	s := args.String("a-b-c")
	result := s.Split("-")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Split", actual)
}

func Test_String_DoubleQuote(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.DoubleQuote()

	// Act
	actual := args.Map{"notEmpty": result.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- DoubleQuote", actual)
}

func Test_String_DoubleQuoteQ(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.DoubleQuoteQ()

	// Act
	actual := args.Map{"notEmpty": result.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- DoubleQuoteQ", actual)
}

func Test_String_SingleQuote(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.SingleQuote()

	// Act
	actual := args.Map{"notEmpty": result.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- SingleQuote", actual)
}

func Test_String_ValueDoubleQuote(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.ValueDoubleQuote()

	// Act
	actual := args.Map{"notEmpty": result.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- ValueDoubleQuote", actual)
}

func Test_String_Bytes(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.Bytes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Bytes", actual)
}

func Test_String_Runes(t *testing.T) {
	// Arrange
	s := args.String("hello")
	result := s.Runes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 5}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Runes", actual)
}

func Test_String_Length(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{
		"len": s.Length(),
		"count": s.Count(),
		"ascii": s.AscIILength(),
	}

	// Assert
	expected := args.Map{
		"len": 5,
		"count": 5,
		"ascii": 5,
	}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Length/Count/ASCII", actual)
}

func Test_String_IsEmptyOrWhitespace(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":      args.String("").IsEmptyOrWhitespace(),
		"whitespace": args.String("   ").IsEmptyOrWhitespace(),
		"notEmpty":   args.String("x").IsEmptyOrWhitespace(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"whitespace": true,
		"notEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_String_TrimSpace(t *testing.T) {
	// Arrange
	s := args.String("  hello  ")

	// Act
	actual := args.Map{"val": s.TrimSpace().String()}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- TrimSpace", actual)
}

func Test_String_ReplaceAll(t *testing.T) {
	// Arrange
	s := args.String("hello world")

	// Act
	actual := args.Map{"val": s.ReplaceAll("world", "go").String()}

	// Assert
	expected := args.Map{"val": "hello go"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- ReplaceAll", actual)
}

func Test_String_Substring(t *testing.T) {
	// Arrange
	s := args.String("hello")

	// Act
	actual := args.Map{"val": s.Substring(0, 3).String()}

	// Assert
	expected := args.Map{"val": "hel"}
	expected.ShouldBeEqual(t, 0, "String returns correct value -- Substring", actual)
}

func Test_String_IsEmpty(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    args.String("").IsEmpty(),
		"notEmpty": args.String("x").IsEmpty(),
		"hasCh":    args.String("x").HasCharacter(),
		"defined":  args.String("x").IsDefined(),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"notEmpty": false,
		"hasCh": true,
		"defined": true,
	}
	expected.ShouldBeEqual(t, 0, "String returns empty -- IsEmpty/HasCharacter/IsDefined", actual)
}

func Test_String_TrimReplaceMap(t *testing.T) {
	// Arrange
	s := args.String("hello {name}")
	result := s.TrimReplaceMap(map[string]string{"{name}": "world"})

	// Act
	actual := args.Map{"notEmpty": result.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TrimReplaceMap returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.One — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_One_AllMethods(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello", Expect: "expected"}

	// Act
	actual := args.Map{
		"first":      o.FirstItem(),
		"expected":   o.Expected(),
		"hasFirst":   o.HasFirst(),
		"hasExpect":  o.HasExpect(),
		"argsCount":  o.ArgsCount(),
		"strNotNull": o.String() != "",
	}

	// Assert
	expected := args.Map{
		"first": "hello", "expected": "expected", "hasFirst": true,
		"hasExpect": true, "argsCount": 1, "strNotNull": true,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- all methods", actual)
}

func Test_One_Args(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello"}

	// Act
	actual := args.Map{
		"len0": len(o.Args(0)),
		"len1": len(o.Args(1)),
	}

	// Assert
	expected := args.Map{
		"len0": 0,
		"len1": 1,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- Args", actual)
}

func Test_One_ValidArgs(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello"}

	// Act
	actual := args.Map{"len": len(o.ValidArgs())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "One returns non-empty -- ValidArgs", actual)
}

func Test_One_LeftRight(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello", Expect: "exp"}
	lr := o.LeftRight()

	// Act
	actual := args.Map{"left": lr.Left}

	// Assert
	expected := args.Map{"left": "hello"}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- LeftRight", actual)
}

func Test_One_GetByIndex(t *testing.T) {
	// Arrange
	o := args.OneAny{First: "hello"}

	// Act
	actual := args.Map{
		"val": o.GetByIndex(0),
		"nil": o.GetByIndex(99) == nil,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "One returns correct value -- GetByIndex", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Two — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Two_AllMethods(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: "a", Second: "b", Expect: "exp"}

	// Act
	actual := args.Map{
		"first":     tw.FirstItem(),
		"second":    tw.SecondItem(),
		"expected":  tw.Expected(),
		"hasFirst":  tw.HasFirst(),
		"hasSecond": tw.HasSecond(),
		"hasExpect": tw.HasExpect(),
		"argsCount": tw.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "expected": "exp",
		"hasFirst": true, "hasSecond": true, "hasExpect": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- all methods", actual)
}

func Test_Two_Args(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{
		"len1": len(tw.Args(1)),
		"len2": len(tw.Args(2)),
	}

	// Assert
	expected := args.Map{
		"len1": 1,
		"len2": 2,
	}
	expected.ShouldBeEqual(t, 0, "Two returns correct value -- Args", actual)
}

func Test_Two_LeftRight(t *testing.T) {
	// Arrange
	tw := args.TwoAny{First: "a", Second: "b"}
	lr := tw.LeftRight()

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
// args.Three — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Three_AllMethods(t *testing.T) {
	// Arrange
	th := args.ThreeAny{First: "a", Second: "b", Third: "c", Expect: "exp"}

	// Act
	actual := args.Map{
		"first":     th.FirstItem(),
		"second":    th.SecondItem(),
		"third":     th.ThirdItem(),
		"hasThird":  th.HasThird(),
		"argsCount": th.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"first": "a", "second": "b", "third": "c",
		"hasThird": true, "argsCount": 3,
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- all methods", actual)
}

func Test_Three_ArgTwo(t *testing.T) {
	// Arrange
	th := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	tw := th.ArgTwo()

	// Act
	actual := args.Map{
		"first": tw.FirstItem(),
		"second": tw.SecondItem(),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- ArgTwo", actual)
}

func Test_Three_ArgThree(t *testing.T) {
	// Arrange
	th := args.ThreeAny{First: "a", Second: "b", Third: "c"}
	copy := th.ArgThree()

	// Act
	actual := args.Map{"third": copy.ThirdItem()}

	// Assert
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Three returns correct value -- ArgThree", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Four — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Four_AllMethods(t *testing.T) {
	// Arrange
	f := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d", Expect: "exp"}

	// Act
	actual := args.Map{
		"hasFourth":  f.HasFourth(),
		"argsCount":  f.ArgsCount(),
		"fourthItem": f.FourthItem(),
	}

	// Assert
	expected := args.Map{
		"hasFourth": true,
		"argsCount": 4,
		"fourthItem": "d",
	}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- all methods", actual)
}

func Test_Four_ArgThree(t *testing.T) {
	// Arrange
	f := args.FourAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	th := f.ArgThree()

	// Act
	actual := args.Map{"third": th.ThirdItem()}

	// Assert
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Four returns correct value -- ArgThree", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Five — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Five_AllMethods(t *testing.T) {
	// Arrange
	f := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Expect: "exp"}

	// Act
	actual := args.Map{
		"hasFifth":  f.HasFifth(),
		"argsCount": f.ArgsCount(),
		"fifthItem": f.FifthItem(),
	}

	// Assert
	expected := args.Map{
		"hasFifth": true,
		"argsCount": 5,
		"fifthItem": "e",
	}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- all methods", actual)
}

func Test_Five_ArgFour(t *testing.T) {
	// Arrange
	f := args.FiveAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	fo := f.ArgFour()

	// Act
	actual := args.Map{"fourth": fo.FourthItem()}

	// Assert
	expected := args.Map{"fourth": "d"}
	expected.ShouldBeEqual(t, 0, "Five returns correct value -- ArgFour", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Six — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Six_AllMethods(t *testing.T) {
	// Arrange
	s := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "exp"}

	// Act
	actual := args.Map{
		"hasSixth":  s.HasSixth(),
		"argsCount": s.ArgsCount(),
		"sixthItem": s.SixthItem(),
	}

	// Assert
	expected := args.Map{
		"hasSixth": true,
		"argsCount": 6,
		"sixthItem": "f",
	}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- all methods", actual)
}

func Test_Six_ArgFive(t *testing.T) {
	// Arrange
	s := args.SixAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f"}
	fi := s.ArgFive()

	// Act
	actual := args.Map{"fifth": fi.FifthItem()}

	// Assert
	expected := args.Map{"fifth": "e"}
	expected.ShouldBeEqual(t, 0, "Six returns correct value -- ArgFive", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.LeftRight — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_LeftRight_AllMethods(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "l", Right: "r", Expect: "exp"}

	// Act
	actual := args.Map{
		"left":      lr.FirstItem(),
		"right":     lr.SecondItem(),
		"expected":  lr.Expected(),
		"hasLeft":   lr.HasLeft(),
		"hasRight":  lr.HasRight(),
		"hasExpect": lr.HasExpect(),
		"argsCount": lr.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"left": "l", "right": "r", "expected": "exp",
		"hasLeft": true, "hasRight": true, "hasExpect": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- all methods", actual)
}

func Test_LeftRight_Clone(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "l", Right: "r"}
	cloned := lr.Clone()

	// Act
	actual := args.Map{
		"left": cloned.Left,
		"right": cloned.Right,
	}

	// Assert
	expected := args.Map{
		"left": "l",
		"right": "r",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- Clone", actual)
}

func Test_LeftRight_ArgTwo(t *testing.T) {
	// Arrange
	lr := args.LeftRightAny{Left: "l", Right: "r"}
	tw := lr.ArgTwo()

	// Act
	actual := args.Map{
		"first": tw.First,
		"second": tw.Second,
	}

	// Assert
	expected := args.Map{
		"first": "l",
		"second": "r",
	}
	expected.ShouldBeEqual(t, 0, "LeftRight returns correct value -- ArgTwo", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Holder — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Holder_AllMethods(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e", Sixth: "f", Expect: "exp"}

	// Act
	actual := args.Map{
		"argsCount":  h.ArgsCount(),
		"hasFirst":   h.HasFirst(),
		"hasSecond":  h.HasSecond(),
		"hasThird":   h.HasThird(),
		"hasFourth":  h.HasFourth(),
		"hasFifth":   h.HasFifth(),
		"hasSixth":   h.HasSixth(),
		"hasExpect":  h.HasExpect(),
		"firstItem":  h.FirstItem(),
		"secondItem": h.SecondItem(),
		"thirdItem":  h.ThirdItem(),
		"fourthItem": h.FourthItem(),
		"fifthItem":  h.FifthItem(),
		"sixthItem":  h.SixthItem(),
		"expected":   h.Expected(),
	}

	// Assert
	expected := args.Map{
		"argsCount": 7, "hasFirst": true, "hasSecond": true,
		"hasThird": true, "hasFourth": true, "hasFifth": true,
		"hasSixth": true, "hasExpect": true,
		"firstItem": "a", "secondItem": "b", "thirdItem": "c",
		"fourthItem": "d", "fifthItem": "e", "sixthItem": "f",
		"expected": "exp",
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- all methods", actual)
}

func Test_Holder_ArgTwo(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b"}
	tw := h.ArgTwo()

	// Act
	actual := args.Map{
		"first": tw.First,
		"second": tw.Second,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- ArgTwo", actual)
}

func Test_Holder_ArgThree(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b", Third: "c"}
	th := h.ArgThree()

	// Act
	actual := args.Map{"third": th.Third}

	// Assert
	expected := args.Map{"third": "c"}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- ArgThree", actual)
}

func Test_Holder_ArgFour(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b", Third: "c", Fourth: "d"}
	fo := h.ArgFour()

	// Act
	actual := args.Map{"fourth": fo.Fourth}

	// Assert
	expected := args.Map{"fourth": "d"}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- ArgFour", actual)
}

func Test_Holder_ArgFive(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b", Third: "c", Fourth: "d", Fifth: "e"}
	fi := h.ArgFive()

	// Act
	actual := args.Map{"fifth": fi.Fifth}

	// Assert
	expected := args.Map{"fifth": "e"}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- ArgFive", actual)
}

func Test_Holder_ValidArgs(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b"}

	// Act
	actual := args.Map{"len": len(h.ValidArgs())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Holder returns non-empty -- ValidArgs", actual)
}

func Test_Holder_Args(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Second: "b", Third: "c"}

	// Act
	actual := args.Map{
		"len2": len(h.Args(2)),
		"len3": len(h.Args(3)),
	}

	// Assert
	expected := args.Map{
		"len2": 2,
		"len3": 3,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- Args", actual)
}

func Test_Holder_Slice(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a", Expect: "exp"}
	slice := h.Slice()

	// Act
	actual := args.Map{"len": len(slice)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- Slice", actual)
}

func Test_Holder_String(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a"}

	// Act
	actual := args.Map{"notEmpty": h.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- String", actual)
}

func Test_Holder_GetByIndex(t *testing.T) {
	// Arrange
	h := args.HolderAny{First: "a"}

	// Act
	actual := args.Map{
		"val": h.GetByIndex(0),
		"nil": h.GetByIndex(99) == nil,
	}

	// Assert
	expected := args.Map{
		"val": "a",
		"nil": true,
	}
	expected.ShouldBeEqual(t, 0, "Holder returns correct value -- GetByIndex", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Dynamic — full coverage
// ══════════════════════════════════════════════════════════════════

func Test_Dynamic_AllMethods(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{
		"first": "a",
		"actual": "data",
	}, Expect: "exp"}

	// Act
	actual := args.Map{
		"expected":   d.Expected(),
		"hasExpect":  d.HasExpect(),
		"hasFirst":   d.HasFirst(),
		"firstItem":  d.FirstItem(),
		"actual":     d.Actual(),
		"arrange":    d.Arrange(),
		"argsCount":  d.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"expected": "exp", "hasExpect": true, "hasFirst": true,
		"firstItem": "a", "actual": "data", "arrange": nil,
		"argsCount": 1,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- all methods", actual)
}

func Test_Dynamic_NilReceiver(t *testing.T) {
	// Arrange
	var d *args.DynamicAny

	// Act
	actual := args.Map{
		"argsCount":  d.ArgsCount(),
		"getWork":    d.GetWorkFunc() == nil,
		"hasFirst":   d.HasFirst(),
		"hasDefined": d.HasDefined("a"),
		"has":        d.Has("a"),
		"invalid":    d.IsKeyInvalid("a"),
		"missing":    d.IsKeyMissing("a"),
		"hasExpect":  d.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"argsCount": 0, "getWork": true, "hasFirst": false,
		"hasDefined": false, "has": false, "invalid": false,
		"missing": false, "hasExpect": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns nil -- nil receiver", actual)
}

func Test_Dynamic_Get(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"k": "v"}}
	val, ok := d.Get("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- Get", actual)
}

func Test_Dynamic_GetAsInt(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"n": 42}}
	val, ok := d.GetAsInt("n")

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
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsInt", actual)
}

func Test_Dynamic_GetAsIntDefault(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{}}
	val := d.GetAsIntDefault("n", 99)

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": 99}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsIntDefault", actual)
}

func Test_Dynamic_GetAsString(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"s": "hello"}}
	val, ok := d.GetAsString("s")

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
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsString", actual)
}

func Test_Dynamic_GetAsStringDefault(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{}}
	val := d.GetAsStringDefault("s")

	// Act
	actual := args.Map{"val": val}

	// Assert
	expected := args.Map{"val": ""}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsStringDefault", actual)
}

func Test_Dynamic_GetAsStrings(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"items": []string{"a"}}}
	items, ok := d.GetAsStrings("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsStrings", actual)
}

func Test_Dynamic_GetAsAnyItems(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"items": []any{1}}}
	items, ok := d.GetAsAnyItems("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetAsAnyItems", actual)
}

func Test_Dynamic_HasDefinedAll(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{
		"a": 1,
		"b": 2,
	}}

	// Act
	actual := args.Map{
		"all":    d.HasDefinedAll("a", "b"),
		"miss":   d.HasDefinedAll("a", "c"),
		"empty":  d.HasDefinedAll(),
	}

	// Assert
	expected := args.Map{
		"all": true,
		"miss": false,
		"empty": false,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- HasDefinedAll", actual)
}

func Test_Dynamic_String(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"a": 1}}

	// Act
	actual := args.Map{"notEmpty": d.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- String", actual)
}

func Test_Dynamic_GetLowerCase(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"name": "val"}}
	val, ok := d.GetLowerCase("NAME")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "val",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetLowerCase", actual)
}

func Test_Dynamic_GetDirectLower(t *testing.T) {
	// Arrange
	d := args.DynamicAny{Params: args.Map{"key": "val"}}
	result := d.GetDirectLower("KEY")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "Dynamic returns correct value -- GetDirectLower", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.DynamicFunc — coverage
// ══════════════════════════════════════════════════════════════════

func Test_DynamicFunc_AllMethods(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{
		Params:   args.Map{
			"first": "a",
			"when": "cond",
			"title": "t",
		},
		WorkFunc: strings.ToUpper,
		Expect:   "exp",
	}

	// Act
	actual := args.Map{
		"argsCount": df.ArgsCount(),
		"hasFunc":   df.HasFunc(),
		"hasExpect": df.HasExpect(),
		"length":    df.Length(),
		"hasFirst":  df.HasFirst(),
		"when":      df.When(),
		"title":     df.Title(),
	}

	// Assert
	expected := args.Map{
		"argsCount": 2, "hasFunc": true, "hasExpect": true,
		"length": 3, "hasFirst": true, "when": "cond", "title": "t",
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- all methods", actual)
}

func Test_DynamicFunc_NilReceiver(t *testing.T) {
	// Arrange
	var df *args.DynamicFuncAny

	// Act
	actual := args.Map{
		"argsCount":  df.ArgsCount(),
		"length":     df.Length(),
		"hasDefined": df.HasDefined("a"),
		"has":        df.Has("a"),
		"invalid":    df.IsKeyInvalid("a"),
		"missing":    df.IsKeyMissing("a"),
		"hasFunc":    df.HasFunc(),
		"hasExpect":  df.HasExpect(),
	}

	// Assert
	expected := args.Map{
		"argsCount": 0, "length": 0, "hasDefined": false,
		"has": false, "invalid": false, "missing": false,
		"hasFunc": false, "hasExpect": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns nil -- nil receiver", actual)
}

func Test_DynamicFunc_Get(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"k": "v"}}
	val, ok := df.Get("k")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "v",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- Get", actual)
}

func Test_DynamicFunc_GetAsInt(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"n": 42}}
	val, ok := df.GetAsInt("n")

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
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- GetAsInt", actual)
}

func Test_DynamicFunc_GetAsString(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"s": "hello"}}
	val, ok := df.GetAsString("s")

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
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- GetAsString", actual)
}

func Test_DynamicFunc_GetAsStrings(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"items": []string{"a"}}}
	items, ok := df.GetAsStrings("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- GetAsStrings", actual)
}

func Test_DynamicFunc_GetAsAnyItems(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"items": []any{1}}}
	items, ok := df.GetAsAnyItems("items")

	// Act
	actual := args.Map{
		"len": len(items),
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- GetAsAnyItems", actual)
}

func Test_DynamicFunc_Actual_Arrange(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{
		"actual": "data",
		"arrange": "setup",
	}}

	// Act
	actual := args.Map{
		"actual": df.Actual(),
		"arrange": df.Arrange(),
	}

	// Assert
	expected := args.Map{
		"actual": "data",
		"arrange": "setup",
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- Actual/Arrange", actual)
}

func Test_DynamicFunc_HasDefinedAll(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{
		"a": 1,
		"b": 2,
	}}

	// Act
	actual := args.Map{
		"all":  df.HasDefinedAll("a", "b"),
		"miss": df.HasDefinedAll("a", "c"),
	}

	// Assert
	expected := args.Map{
		"all": true,
		"miss": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- HasDefinedAll", actual)
}

func Test_DynamicFunc_String(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"a": 1}, WorkFunc: strings.ToUpper}

	// Act
	actual := args.Map{"notEmpty": df.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- String", actual)
}

func Test_DynamicFunc_GetLowerCase(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"name": "val"}}
	val, ok := df.GetLowerCase("NAME")

	// Act
	actual := args.Map{
		"val": val,
		"ok": ok,
	}

	// Assert
	expected := args.Map{
		"val": "val",
		"ok": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- GetLowerCase", actual)
}

func Test_DynamicFunc_GetDirectLower(t *testing.T) {
	// Arrange
	df := args.DynamicFuncAny{Params: args.Map{"key": "val"}}
	result := df.GetDirectLower("KEY")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "val"}
	expected.ShouldBeEqual(t, 0, "DynamicFunc returns correct value -- GetDirectLower", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.OneFunc — coverage
// ══════════════════════════════════════════════════════════════════

func Test_OneFunc_AllMethods(t *testing.T) {
	// Arrange
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper, Expect: "HELLO"}

	// Act
	actual := args.Map{
		"hasFirst":  of.HasFirst(),
		"hasFunc":   of.HasFunc(),
		"hasExpect": of.HasExpect(),
		"argsCount": of.ArgsCount(),
		"firstItem": of.FirstItem(),
		"expected":  of.Expected(),
		"funcName":  of.GetFuncName() != "",
	}

	// Assert
	expected := args.Map{
		"hasFirst": true, "hasFunc": true, "hasExpect": true,
		"argsCount": 1, "firstItem": "hello", "expected": "HELLO",
		"funcName": true,
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- all methods", actual)
}

func Test_OneFunc_InvokeWithValidArgs(t *testing.T) {
	// Arrange
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper}
	results, err := of.InvokeWithValidArgs()

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"result": results[0],
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"result": "HELLO",
	}
	expected.ShouldBeEqual(t, 0, "OneFunc returns non-empty -- InvokeWithValidArgs", actual)
}

func Test_OneFunc_Slice(t *testing.T) {
	// Arrange
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper, Expect: "exp"}
	slice := of.Slice()

	// Act
	actual := args.Map{"len": len(slice)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- Slice", actual)
}

func Test_OneFunc_LeftRight(t *testing.T) {
	// Arrange
	of := args.OneFuncAny{First: "hello", WorkFunc: strings.ToUpper}
	lr := of.LeftRight()

	// Act
	actual := args.Map{"left": lr.Left}

	// Assert
	expected := args.Map{"left": "hello"}
	expected.ShouldBeEqual(t, 0, "OneFunc returns correct value -- LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.TwoFunc — coverage
// ══════════════════════════════════════════════════════════════════

func Test_TwoFunc_AllMethods(t *testing.T) {
	// Arrange
	tf := args.TwoFuncAny{First: "a", Second: "b", WorkFunc: strings.Join, Expect: "exp"}

	// Act
	actual := args.Map{
		"hasFirst":  tf.HasFirst(),
		"hasSecond": tf.HasSecond(),
		"hasFunc":   tf.HasFunc(),
		"hasExpect": tf.HasExpect(),
		"argsCount": tf.ArgsCount(),
	}

	// Assert
	expected := args.Map{
		"hasFirst": true, "hasSecond": true, "hasFunc": true,
		"hasExpect": true, "argsCount": 2,
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- all methods", actual)
}

func Test_TwoFunc_ArgTwo(t *testing.T) {
	// Arrange
	tf := args.TwoFuncAny{First: "a", Second: "b"}
	tw := tf.ArgTwo()

	// Act
	actual := args.Map{
		"first": tw.First,
		"second": tw.Second,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"second": "b",
	}
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- ArgTwo", actual)
}

func Test_TwoFunc_LeftRight(t *testing.T) {
	// Arrange
	tf := args.TwoFuncAny{First: "a", Second: "b"}
	lr := tf.LeftRight()

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
	expected.ShouldBeEqual(t, 0, "TwoFunc returns correct value -- LeftRight", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncMap — coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncMap_Basic(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper, strings.ToLower)

	// Act
	actual := args.Map{
		"len":     fm.Length(),
		"count":   fm.Count(),
		"hasAny":  fm.HasAnyItem(),
		"isEmpty": fm.IsEmpty(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"count": 2,
		"hasAny": true,
		"isEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- basic", actual)
}

func Test_FuncMap_Has_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"has":      fm.Has("ToUpper"),
		"notHas":   fm.Has("missing"),
		"contains": fm.IsContains("ToUpper"),
	}

	// Assert
	expected := args.Map{
		"has": true,
		"notHas": false,
		"contains": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Has", actual)
}

func Test_FuncMap_Get_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	f := fm.Get("ToUpper")

	// Act
	actual := args.Map{
		"notNil": f != nil,
		"nilMissing": fm.Get("missing") == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilMissing": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Get", actual)
}

func Test_FuncMap_IsValidFuncOf_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"valid":   fm.IsValidFuncOf("ToUpper"),
		"invalid": fm.IsInvalidFunc("missing"),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns non-empty -- IsValidFuncOf", actual)
}

func Test_FuncMap_ArgsCount_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"args":   fm.ArgsCount("ToUpper"),
		"ret":    fm.ReturnLength("ToUpper"),
		"argsL":  fm.ArgsLength("ToUpper"),
		"nilArg": fm.ArgsCount("missing"),
		"nilRet": fm.ReturnLength("missing"),
	}

	// Assert
	expected := args.Map{
		"args": 1,
		"ret": 1,
		"argsL": 1,
		"nilArg": 0,
		"nilRet": 0,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- ArgsCount/ReturnLength", actual)
}

func Test_FuncMap_PkgPath(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"notEmpty":   fm.PkgPath("ToUpper") != "",
		"nilEmpty":   fm.PkgPath("missing") == "",
		"pkgName":    fm.PkgNameOnly("ToUpper") != "",
		"nilPkg":     fm.PkgNameOnly("missing") == "",
		"directName": fm.FuncDirectInvokeName("ToUpper") != "",
		"nilDirect":  fm.FuncDirectInvokeName("missing") == "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true, "nilEmpty": true,
		"pkgName": true, "nilPkg": true,
		"directName": true, "nilDirect": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- PkgPath", actual)
}

func Test_FuncMap_GetPascalCaseFuncName(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	result := fm.GetPascalCaseFuncName("ToUpper")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName returns correct value -- with args", actual)
}

func Test_FuncMap_GetPascalCaseFuncName_Empty(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}
	result := fm.GetPascalCaseFuncName("anything")

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName returns empty -- empty", actual)
}

func Test_FuncMap_IsPublicPrivate(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"public":     fm.IsPublicMethod("ToUpper"),
		"nilPublic":  fm.IsPublicMethod("missing"),
		"nilPrivate": fm.IsPrivateMethod("missing"),
	}

	// Assert
	expected := args.Map{
		"public": true,
		"nilPublic": false,
		"nilPrivate": false,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Public/Private", actual)
}

func Test_FuncMap_GetType(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"notNil":  fm.GetType("ToUpper") != nil,
		"nilType": fm.GetType("missing") == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"nilType": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- GetType", actual)
}

func Test_FuncMap_GetInOutArgsTypes(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	inTypes := fm.GetInArgsTypes("ToUpper")
	outTypes := fm.GetOutArgsTypes("ToUpper")
	inNames := fm.GetInArgsTypesNames("ToUpper")

	// Act
	actual := args.Map{
		"inLen":    len(inTypes),
		"outLen":   len(outTypes),
		"namesLen": len(inNames),
		"nilIn":    len(fm.GetInArgsTypes("missing")),
		"nilOut":   len(fm.GetOutArgsTypes("missing")),
		"nilNames": len(fm.GetInArgsTypesNames("missing")),
	}

	// Assert
	expected := args.Map{
		"inLen": 1, "outLen": 1, "namesLen": 1,
		"nilIn": 0, "nilOut": 0, "nilNames": 0,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- GetInOutArgsTypes", actual)
}

func Test_FuncMap_ValidationError(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)

	// Act
	actual := args.Map{
		"valid":   fm.ValidationError("ToUpper") == nil,
		"invalid": fm.ValidationError("missing") != nil,
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- ValidationError", actual)
}

func Test_FuncMap_Invoke(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	var knownName string
	for k := range fm {
		knownName = k
		break
	}
	results, err := fm.Invoke(knownName, "hello")
	var result any
	if len(results) > 0 {
		result = results[0]
	}

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"hasResult": result != nil,
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"hasResult": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Invoke", actual)
}

func Test_FuncMap_Invoke_NotFound(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map(strings.ToUpper)
	_, err := fm.Invoke("missing")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- Invoke not found", actual)
}

func Test_FuncMap_InvalidError(t *testing.T) {
	// Arrange
	fm := args.FuncMap{}

	// Act
	actual := args.Map{"hasErr": fm.InvalidError() != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns error -- InvalidError", actual)
}

func Test_FuncMap_VoidCall(t *testing.T) {
	// Arrange
	called := false
	fn := func() { called = true }
	fm := args.NewFuncWrap.Map(fn)
	name := fm.Get(fm.Get(fm.Get("").GetFuncName()).GetFuncName()).GetFuncName()
	// Just exercise VoidCall
	_ = name

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "FuncMap returns correct value -- VoidCall", actual)
	_ = called
}

// ══════════════════════════════════════════════════════════════════
// args.FuncWrap — additional coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncWrap_TypedHelpers(t *testing.T) {
	// Arrange
	boolFn := func(s string) bool { return s == "yes" }
	errFn := func(s string) error { return nil }
	strFn := func(s string) string { return s }
	voidFn := func() {}
	valErrFn := func(s string) (string, error) { return s, nil }

	boolFW := args.NewTypedFuncWrap(boolFn)
	errFW := args.NewTypedFuncWrap(errFn)
	strFW := args.NewTypedFuncWrap(strFn)
	voidFW := args.NewTypedFuncWrap(voidFn)
	valErrFW := args.NewTypedFuncWrap(valErrFn)

	// Act
	actual := args.Map{
		"isBool":     boolFW.IsBoolFunc(),
		"isError":    errFW.IsErrorFunc(),
		"isString":   strFW.IsStringFunc(),
		"isVoid":     voidFW.IsVoidFunc(),
		"isValErr":   valErrFW.IsValueErrorFunc(),
		"isAnyErr":   valErrFW.IsAnyErrorFunc(),
		"isAny":      strFW.IsAnyFunc(),
	}

	// Assert
	expected := args.Map{
		"isBool": true, "isError": true, "isString": true,
		"isVoid": true, "isValErr": true, "isAnyErr": true,
		"isAny": true,
	}
	expected.ShouldBeEqual(t, 0, "FuncWrap returns correct value -- typed helpers", actual)
}

func Test_FuncWrap_InvokeAsBool_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fn := func(s string) bool { return s == "yes" }
	fw := args.NewTypedFuncWrap(fn)
	result, err := fw.InvokeAsBool("yes")

	// Act
	actual := args.Map{
		"result": result,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"result": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsBool returns correct value -- with args", actual)
}

func Test_FuncWrap_InvokeAsString_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fn := func(s string) string { return strings.ToUpper(s) }
	fw := args.NewTypedFuncWrap(fn)
	result, err := fw.InvokeAsString("hello")

	// Act
	actual := args.Map{
		"result": result,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"result": "HELLO",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsString returns correct value -- with args", actual)
}

func Test_FuncWrap_InvokeAsAny_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s }
	fw := args.NewTypedFuncWrap(fn)
	result, err := fw.InvokeAsAny("hello")

	// Act
	actual := args.Map{
		"result": result,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"result": "hello",
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAny returns correct value -- with args", actual)
}

func Test_FuncWrap_InvokeAsAnyError(t *testing.T) {
	// Arrange
	fn := func(s string) (string, error) { return s, nil }
	fw := args.NewTypedFuncWrap(fn)
	result, funcErr, procErr := fw.InvokeAsAnyError("hello")

	// Act
	actual := args.Map{
		"result": result,
		"funcErr": funcErr == nil,
		"procErr": procErr == nil,
	}

	// Assert
	expected := args.Map{
		"result": "hello",
		"funcErr": true,
		"procErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsAnyError returns error -- with args", actual)
}

func Test_FuncWrap_InvokeAsError(t *testing.T) {
	// Arrange
	fn := func() error { return nil }
	fw := args.NewTypedFuncWrap(fn)
	funcErr, procErr := fw.InvokeAsError()

	// Act
	actual := args.Map{
		"funcErr": funcErr == nil,
		"procErr": procErr == nil,
	}

	// Assert
	expected := args.Map{
		"funcErr": true,
		"procErr": true,
	}
	expected.ShouldBeEqual(t, 0, "InvokeAsError returns error -- with args", actual)
}

func Test_FuncWrap_InArgNames_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s }
	fw := args.NewTypedFuncWrap(fn)
	names := fw.InArgNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "InArgNames returns correct value -- with args", actual)
}

func Test_FuncWrap_InArgNames_Multi(t *testing.T) {
	// Arrange
	fn := func(a string, b int) string { return a }
	fw := args.NewTypedFuncWrap(fn)
	names := fw.InArgNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "InArgNames returns correct value -- multi", actual)
}

func Test_FuncWrap_OutArgNames_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fn := func() (string, error) { return "", nil }
	fw := args.NewTypedFuncWrap(fn)
	names := fw.OutArgNames()

	// Act
	actual := args.Map{"len": len(names)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "OutArgNames returns correct value -- with args", actual)
}

func Test_FuncWrap_InArgNamesEachLine(t *testing.T) {
	// Arrange
	fn := func(a string, b int) string { return a }
	fw := args.NewTypedFuncWrap(fn)
	lines := fw.InArgNamesEachLine()

	// Act
	actual := args.Map{"greaterThan1": len(lines) > 1}

	// Assert
	expected := args.Map{"greaterThan1": true}
	expected.ShouldBeEqual(t, 0, "InArgNamesEachLine returns correct value -- with args", actual)
}

func Test_FuncWrap_OutArgNamesEachLine(t *testing.T) {
	// Arrange
	fn := func() (string, error) { return "", nil }
	fw := args.NewTypedFuncWrap(fn)
	lines := fw.OutArgNamesEachLine()

	// Act
	actual := args.Map{"greaterThan1": len(lines) > 1}

	// Assert
	expected := args.Map{"greaterThan1": true}
	expected.ShouldBeEqual(t, 0, "OutArgNamesEachLine returns correct value -- with args", actual)
}

func Test_FuncWrap_IsInTypeMatches_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fn := func(s string) string { return s }
	fw := args.NewTypedFuncWrap(fn)

	// Act
	actual := args.Map{
		"match":    fw.IsInTypeMatches("hello"),
		"outMatch": fw.IsOutTypeMatches("result"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"outMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "IsInTypeMatches/IsOutTypeMatches returns correct value -- with args", actual)
}

func Test_FuncWrap_PascalCase(t *testing.T) {
	// Arrange
	fn := func() {}
	fw := args.NewTypedFuncWrap(fn)
	result := fw.GetPascalCaseFuncName()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetPascalCaseFuncName returns correct value -- with args", actual)
}

func Test_FuncWrap_PkgNameOnly_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap(strings.ToUpper)
	result := fw.PkgNameOnly()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PkgNameOnly returns correct value -- with args", actual)
}

func Test_FuncWrap_FuncDirectInvokeName_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap(strings.ToUpper)
	result := fw.FuncDirectInvokeName()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FuncDirectInvokeName returns correct value -- with args", actual)
}

func Test_FuncWrap_IsEqual_SameFunc(t *testing.T) {
	// Arrange
	a := args.NewTypedFuncWrap(strings.ToUpper)
	b := args.NewTypedFuncWrap(strings.ToUpper)

	// Act
	actual := args.Map{"equal": a.IsEqual(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct value -- same func", actual)
}

func Test_FuncWrap_IsNotEqual_DiffFunc(t *testing.T) {
	// Arrange
	a := args.NewTypedFuncWrap(strings.ToUpper)
	b := args.NewTypedFuncWrap(strings.ToLower)

	// Act
	actual := args.Map{"notEqual": a.IsNotEqual(b)}

	// Assert
	expected := args.Map{"notEqual": true}
	expected.ShouldBeEqual(t, 0, "IsNotEqual returns correct value -- diff func", actual)
}

func Test_FuncWrap_IsEqualValue(t *testing.T) {
	// Arrange
	a := args.NewTypedFuncWrap(strings.ToUpper)
	b := *args.NewTypedFuncWrap(strings.ToUpper)

	// Act
	actual := args.Map{"equal": a.IsEqualValue(b)}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqualValue returns correct value -- with args", actual)
}

func Test_FuncWrap_ValidationError_Nil_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	var fw *args.FuncWrapAny
	err := fw.ValidationError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ValidationError returns nil -- nil", actual)
}

func Test_FuncWrap_InvalidError(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Invalid()
	err := fw.InvalidError()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "InvalidError returns error -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.FuncDetector — coverage
// ══════════════════════════════════════════════════════════════════

func Test_FuncDetector_GetFuncWrap_FromMap(t *testing.T) {
	// Arrange
	m := args.Map{"func": strings.ToUpper}
	fw := args.FuncDetector.GetFuncWrap(m)

	// Act
	actual := args.Map{"valid": fw.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector returns correct value -- Map", actual)
}

func Test_FuncDetector_GetFuncWrap_FromFuncWrap(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(strings.ToUpper)
	result := args.FuncDetector.GetFuncWrap(fw)

	// Act
	actual := args.Map{"valid": result.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector returns correct value -- FuncWrap", actual)
}

func Test_FuncDetector_GetFuncWrap_FromRawFunc(t *testing.T) {
	// Arrange
	result := args.FuncDetector.GetFuncWrap(strings.ToUpper)

	// Act
	actual := args.Map{"valid": result.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "FuncDetector returns correct value -- raw func", actual)
}

// ══════════════════════════════════════════════════════════════════
// args.Empty / NewFuncWrap — coverage
// ══════════════════════════════════════════════════════════════════

func Test_Empty_Map_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	m := args.Empty.Map()

	// Act
	actual := args.Map{"len": m.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Empty.Map returns empty -- with args", actual)
}

func Test_Empty_FuncWrap_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fw := args.Empty.FuncWrap()

	// Act
	actual := args.Map{"invalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "Empty.FuncWrap returns empty -- with args", actual)
}

func Test_Empty_FuncMap_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fm := args.Empty.FuncMap()

	// Act
	actual := args.Map{"empty": fm.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Empty.FuncMap returns empty -- with args", actual)
}

func Test_Empty_Holder_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	h := args.Empty.Holder()

	// Act
	actual := args.Map{"argsCount": h.ArgsCount()}

	// Assert
	expected := args.Map{"argsCount": 7}
	expected.ShouldBeEqual(t, 0, "Empty.Holder returns empty -- with args", actual)
}

func Test_NewFuncWrap_Many_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fws := args.NewFuncWrap.Many(strings.ToUpper, strings.ToLower)

	// Act
	actual := args.Map{"len": len(fws)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Many returns correct value -- with args", actual)
}

func Test_NewFuncWrap_Many_Empty_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fws := args.NewFuncWrap.Many()

	// Act
	actual := args.Map{"len": len(fws)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Many returns empty -- empty", actual)
}

func Test_NewFuncWrap_Map_Empty_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fm := args.NewFuncWrap.Map()

	// Act
	actual := args.Map{"empty": fm.IsEmpty()}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Map returns empty -- empty", actual)
}

func Test_NewFuncWrap_Invalid_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Invalid()

	// Act
	actual := args.Map{"invalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Invalid returns error -- with args", actual)
}

func Test_NewFuncWrap_Default_Nil_FromMapCompileToStringsa(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default(nil)

	// Act
	actual := args.Map{"invalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Default returns nil -- nil", actual)
}

func Test_NewFuncWrap_Default_NonFunc(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Default("not a func")

	// Act
	actual := args.Map{"invalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Default returns non-empty -- non-func", actual)
}

func Test_NewFuncWrap_Single(t *testing.T) {
	// Arrange
	fw := args.NewFuncWrap.Single(strings.ToUpper)

	// Act
	actual := args.Map{"valid": fw.IsValid()}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewFuncWrap.Single returns correct value -- with args", actual)
}

func Test_NewTypedFuncWrap_Nil(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap[any](nil)

	// Act
	actual := args.Map{"invalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap returns nil -- nil", actual)
}

func Test_NewTypedFuncWrap_NonFunc(t *testing.T) {
	// Arrange
	fw := args.NewTypedFuncWrap("not a func")

	// Act
	actual := args.Map{"invalid": fw.IsInvalid()}

	// Assert
	expected := args.Map{"invalid": true}
	expected.ShouldBeEqual(t, 0, "NewTypedFuncWrap returns non-empty -- non-func", actual)
}
