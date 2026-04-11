package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════
// BasicInt8 – all methods
// ══════════════════════════════════════════

func Test_BasicInt8_Default(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
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
		"min":          int8(0),
		"max":          int8(2),
		"isValidRange": true,
		"outOfRange":   false,
		"toString":     "Active",
		"typeName":     "enumimpltests.myInt8",
		"length":       3,
		"enumType":     "Integer8",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_Default returns correct value -- with args", actual)
}

func Test_BasicInt8_IsAnyOf(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt8_IsAnyOf returns correct value -- with args", actual)
}

func Test_BasicInt8_IsAnyNamesOf(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt8_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_BasicInt8_GetValueByString(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"byName": bi.GetValueByString("Active"),
	}

	// Assert
	expected := args.Map{
		"byName": int8(1),
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_GetValueByString returns correct value -- with args", actual)
}

func Test_BasicInt8_GetValueByName(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
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
		"val":      int8(1),
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_GetValueByName returns correct value -- with args", actual)
}

func Test_BasicInt8_GetStringValue(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"val": bi.GetStringValue(0),
	}

	// Assert
	expected := args.Map{
		"val": "Invalid",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_GetStringValue returns correct value -- with args", actual)
}

func Test_BasicInt8_Ranges(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"rangesLen":  len(bi.Ranges()),
		"hmLen":      len(bi.Hashmap()) > 0,
		"hmPtrNotNil": bi.HashmapPtr() != nil,
	}

	// Assert
	expected := args.Map{
		"rangesLen":  2,
		"hmLen":      true,
		"hmPtrNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_Ranges returns correct value -- with args", actual)
}

func Test_BasicInt8_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt8_ToEnumJsonBytes returns correct value -- with args", actual)
}

func Test_BasicInt8_AppendPrependJoin(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_AppendPrependJoin returns correct value -- with args", actual)
}

func Test_BasicInt8_ToNumberString(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"notEmpty": bi.ToNumberString(1) != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_ToNumberString returns correct value -- with args", actual)
}

func Test_BasicInt8_UnmarshallToValue(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	val1, err1 := bi.UnmarshallToValue(true, nil)
	_, err2 := bi.UnmarshallToValue(false, nil)
	val3, err3 := bi.UnmarshallToValue(true, []byte(""))
	val4, err4 := bi.UnmarshallToValue(true, []byte(`""`))
	val5, err5 := bi.UnmarshallToValue(false, []byte("Active"))

	// Act
	actual := args.Map{
		"nilMapped":      val1,
		"nilMappedErr":   err1 == nil,
		"nilNoMapErr":    err2 != nil,
		"emptyMapped":    val3,
		"emptyNoErr":     err3 == nil,
		"dqMapped":       val4,
		"dqNoErr":        err4 == nil,
		"validVal":       val5,
		"validNoErr":     err5 == nil,
	}

	// Assert
	expected := args.Map{
		"nilMapped":      int8(0),
		"nilMappedErr":   true,
		"nilNoMapErr":    true,
		"emptyMapped":    int8(0),
		"emptyNoErr":     true,
		"dqMapped":       int8(0),
		"dqNoErr":        true,
		"validVal":       int8(1),
		"validNoErr":     true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_UnmarshallToValue returns correct value -- with args", actual)
}

func Test_BasicInt8_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bi.ExpectingEnumValueError("Active", int8(1))
	hasErr := bi.ExpectingEnumValueError("Invalid", int8(1))
	parseErr := bi.ExpectingEnumValueError("NotExist", int8(1))

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
	expected.ShouldBeEqual(t, 0, "BasicInt8_ExpectingEnumValueError returns error -- with args", actual)
}

// ── newBasicInt8Creator paths ──

func Test_BasicInt8_CreateUsingMap(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.CreateUsingMap(
		"testInt8Enum",
		map[int8]string{0: "Invalid", 1: "Active"},
	)

	// Act
	actual := args.Map{
		"typeName": bi.TypeName(),
		"length":   bi.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testInt8Enum",
		"length":   2,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_CreateUsingMap returns correct value -- with args", actual)
}

func Test_BasicInt8_WithAliasMap(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.DefaultWithAliasMap(
		myInt8(0),
		[]string{"Invalid", "Active"},
		map[string]int8{"on": 1},
	)

	val, err := bi.GetValueByName("on")

	// Act
	actual := args.Map{
		"aliasVal": val,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"aliasVal": int8(1),
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_WithAliasMap returns non-empty -- with args", actual)
}

func Test_BasicInt8_DefaultAllCases(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.DefaultAllCases(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	val, err := bi.GetValueByName("active")

	// Act
	actual := args.Map{
		"lowerVal": val,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"lowerVal": int8(1),
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_DefaultAllCases returns correct value -- with args", actual)
}

func Test_BasicInt8_DefaultWithAliasMapAllCases(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.DefaultWithAliasMapAllCases(
		myInt8(0),
		[]string{"Invalid", "Active"},
		map[string]int8{"enabled": 1},
	)

	val, err := bi.GetValueByName("ENABLED")

	// Act
	actual := args.Map{
		"upperAlias": val,
		"noErr":      err == nil,
	}

	// Assert
	expected := args.Map{
		"upperAlias": int8(1),
		"noErr":      true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_DefaultWithAliasMapAllCases returns non-empty -- with args", actual)
}

func Test_BasicInt8_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.CreateUsingMapPlusAliasMapOptions(
		true,
		myInt8(0),
		map[int8]string{0: "Invalid", 1: "Active"},
		map[string]int8{"on": 1},
	)

	val, err := bi.GetValueByName("ON")

	// Act
	actual := args.Map{
		"upperAlias": val,
		"noErr":      err == nil,
	}

	// Assert
	expected := args.Map{
		"upperAlias": int8(1),
		"noErr":      true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_CreateUsingMapPlusAliasMapOptions returns correct value -- with args", actual)
}

func Test_BasicInt8_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.UsingFirstItemSliceAliasMap(
		myInt8(0),
		[]string{"Invalid", "Active"},
		map[string]int8{"on": 1},
	)

	// Act
	actual := args.Map{
		"typeName": bi.TypeName(),
	}

	// Assert
	expected := args.Map{
		"typeName": "enumimpltests.myInt8",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_UsingFirstItemSliceAliasMap returns correct value -- with args", actual)
}

func Test_BasicInt8_UsingTypeSlice(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice(
		"testInt8",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{
		"typeName": bi.TypeName(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testInt8",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_UsingTypeSlice returns correct value -- with args", actual)
}

func Test_BasicInt8_AppendPrependJoinNamer(t *testing.T) {
	// Arrange
	type myInt8 int8
	bi := enumimpl.New.BasicInt8.Default(
		myInt8(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinNamer(".", mockNamer4{"B"}, mockNamer4{"A"})

	// Act
	actual := args.Map{
		"result": result,
	}

	// Assert
	expected := args.Map{
		"result": "A.B",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt8_AppendPrependJoinNamer returns correct value -- with args", actual)
}

type mockNamer4 struct{ name string }

func (m mockNamer4) Name() string { return m.name }
