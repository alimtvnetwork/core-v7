package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════
// BasicInt16 – all methods
// ══════════════════════════════════════════

func Test_BasicInt16_Default(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"min":          bi.Min(),
		"max":          bi.Max(),
		"isValidRange": bi.IsValidRange(1),
		"outOfRange":   bi.IsValidRange(5),
		"toString":     bi.ToEnumString(1),
		"typeName":     bi.TypeName(),
		"length":       bi.Length(),
		"enumType":     bi.EnumType().String(),
	}

	// Assert
	expected := args.Map{
		"min":          int16(0),
		"max":          int16(2),
		"isValidRange": true,
		"outOfRange":   false,
		"toString":     "Active",
		"typeName":     "enumimpltests.myInt16",
		"length":       3,
		"enumType":     "Integer16",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_Default returns correct value -- with args", actual)
}

func Test_BasicInt16_IsAnyOf(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active", "Inactive"},
	)

	// Act
	actual := args.Map{
		"empty":   bi.IsAnyOf(1),
		"match":   bi.IsAnyOf(1, 0, 1, 2),
		"noMatch": bi.IsAnyOf(1, 0, 2),
	}

	// Assert
	expected := args.Map{
		"empty":   true,
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_IsAnyOf returns correct value -- with args", actual)
}

func Test_BasicInt16_IsAnyNamesOf(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"match":   bi.IsAnyNamesOf(1, "Active"),
		"noMatch": bi.IsAnyNamesOf(1, "Invalid"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_BasicInt16_GetValueByName(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bi.GetValueByName("Active")
	_, errNotFound := bi.GetValueByName("NotExist")

	// Act
	actual := args.Map{
		"val":      val,
		"noErr":    err == nil,
		"hasError": errNotFound != nil,
	}

	// Assert
	expected := args.Map{
		"val":      int16(1),
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_GetValueByName returns correct value -- with args", actual)
}

func Test_BasicInt16_GetStringValue(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"val": bi.GetStringValue(0)}

	// Assert
	expected := args.Map{"val": "Invalid"}
	expected.ShouldBeEqual(t, 0, "BasicInt16_GetStringValue returns correct value -- with args", actual)
}

func Test_BasicInt16_Ranges(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"rangesLen":   len(bi.Ranges()),
		"hmLen":       len(bi.Hashmap()) > 0,
		"hmPtrNotNil": bi.HashmapPtr() != nil,
	}

	// Assert
	expected := args.Map{
		"rangesLen":   2,
		"hmLen":       true,
		"hmPtrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_Ranges returns correct value -- with args", actual)
}

func Test_BasicInt16_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	jsonBytes, err := bi.ToEnumJsonBytes(0)
	_, errNotFound := bi.ToEnumJsonBytes(99)

	// Act
	actual := args.Map{
		"hasBytes": len(jsonBytes) > 0,
		"noErr":    err == nil,
		"notFound": errNotFound != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr":    true,
		"notFound": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_ToEnumJsonBytes returns correct value -- with args", actual)
}

func Test_BasicInt16_AppendPrependJoinValue(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BasicInt16_AppendPrependJoinValue returns correct value -- with args", actual)
}

func Test_BasicInt16_AppendPrependJoinNamer(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinNamer(".", mockNamer5{"B"}, mockNamer5{"A"})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "A.B"}
	expected.ShouldBeEqual(t, 0, "BasicInt16_AppendPrependJoinNamer returns correct value -- with args", actual)
}

func Test_BasicInt16_ToNumberString(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"notEmpty": bi.ToNumberString(1) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BasicInt16_ToNumberString returns correct value -- with args", actual)
}

func Test_BasicInt16_UnmarshallToValue(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	val1, err1 := bi.UnmarshallToValue(true, nil)
	_, err2 := bi.UnmarshallToValue(false, nil)
	val3, err3 := bi.UnmarshallToValue(true, []byte(""))
	val4, err4 := bi.UnmarshallToValue(true, []byte(`""`))
	val5, err5 := bi.UnmarshallToValue(false, []byte("Active"))

	// Act
	actual := args.Map{
		"nilMapped":   val1,
		"nilErr":      err1 == nil,
		"nilNoMapErr": err2 != nil,
		"emptyVal":    val3,
		"emptyErr":    err3 == nil,
		"dqVal":       val4,
		"dqErr":       err4 == nil,
		"validVal":    val5,
		"validErr":    err5 == nil,
	}

	// Assert
	expected := args.Map{
		"nilMapped":   int16(0),
		"nilErr":      true,
		"nilNoMapErr": true,
		"emptyVal":    int16(0),
		"emptyErr":    true,
		"dqVal":       int16(0),
		"dqErr":       true,
		"validVal":    int16(1),
		"validErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_UnmarshallToValue returns correct value -- with args", actual)
}

func Test_BasicInt16_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bi.ExpectingEnumValueError("Active", int16(1))
	hasErr := bi.ExpectingEnumValueError("Invalid", int16(1))
	parseErr := bi.ExpectingEnumValueError("NotExist", int16(1))

	// Act
	actual := args.Map{
		"matchNoErr":  noErr == nil,
		"mismatchErr": hasErr != nil,
		"parseErr":    parseErr != nil,
	}

	// Assert
	expected := args.Map{
		"matchNoErr":  true,
		"mismatchErr": true,
		"parseErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_ExpectingEnumValueError returns error -- with args", actual)
}

// ── newBasicInt16Creator paths ──

func Test_BasicInt16_CreateUsingMap(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.CreateUsingMap(
		"testInt16Enum",
		map[int16]string{0: "Invalid", 1: "Active"},
	)

	// Act
	actual := args.Map{
		"typeName": bi.TypeName(),
		"length": bi.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testInt16Enum",
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_CreateUsingMap returns correct value -- with args", actual)
}

func Test_BasicInt16_WithAliasMap(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.DefaultWithAliasMap(
		myInt16(0),
		[]string{"Invalid", "Active"},
		map[string]int16{"on": 1},
	)

	val, err := bi.GetValueByName("on")

	// Act
	actual := args.Map{
		"aliasVal": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"aliasVal": int16(1),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_WithAliasMap returns non-empty -- with args", actual)
}

func Test_BasicInt16_DefaultAllCases(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.DefaultAllCases(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bi.GetValueByName("active")

	// Act
	actual := args.Map{
		"lowerVal": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"lowerVal": int16(1),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_DefaultAllCases returns correct value -- with args", actual)
}

func Test_BasicInt16_DefaultWithAliasMapAllCases(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.DefaultWithAliasMapAllCases(
		myInt16(0),
		[]string{"Invalid", "Active"},
		map[string]int16{"enabled": 1},
	)

	val, err := bi.GetValueByName("ENABLED")

	// Act
	actual := args.Map{
		"upperAlias": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"upperAlias": int16(1),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_DefaultWithAliasMapAllCases returns non-empty -- with args", actual)
}

func Test_BasicInt16_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.CreateUsingMapPlusAliasMapOptions(
		true,
		myInt16(0),
		map[int16]string{0: "Invalid", 1: "Active"},
		map[string]int16{"on": 1},
	)

	val, err := bi.GetValueByName("ON")

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int16(1),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt16_CreateUsingMapPlusAliasMapOptions returns correct value -- with args", actual)
}

func Test_BasicInt16_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.UsingFirstItemSliceAliasMap(
		myInt16(0),
		[]string{"Invalid", "Active"},
		map[string]int16{"on": 1},
	)

	// Act
	actual := args.Map{"typeName": bi.TypeName()}

	// Assert
	expected := args.Map{"typeName": "enumimpltests.myInt16"}
	expected.ShouldBeEqual(t, 0, "BasicInt16_UsingFirstItemSliceAliasMap returns correct value -- with args", actual)
}

func Test_BasicInt16_UsingTypeSlice(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice(
		"testInt16",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"typeName": bi.TypeName()}

	// Assert
	expected := args.Map{"typeName": "testInt16"}
	expected.ShouldBeEqual(t, 0, "BasicInt16_UsingTypeSlice returns correct value -- with args", actual)
}

func Test_BasicInt16_GetValueByString(t *testing.T) {
	// Arrange
	type myInt16 int16
	bi := enumimpl.New.BasicInt16.Default(
		myInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"byName": bi.GetValueByString("Active")}

	// Assert
	expected := args.Map{"byName": int16(1)}
	expected.ShouldBeEqual(t, 0, "BasicInt16_GetValueByString returns correct value -- with args", actual)
}

type mockNamer5 struct{ name string }

func (m mockNamer5) Name() string { return m.name }
