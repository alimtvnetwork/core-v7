package enumimpltests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════
// BasicInt32 – all methods
// ══════════════════════════════════════════

func Test_BasicInt32_Default(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
		"min":          int32(0),
		"max":          int32(2),
		"isValidRange": true,
		"outOfRange":   false,
		"toString":     "Active",
		"typeName":     "enumimpltests.myInt32",
		"length":       3,
		"enumType":     "Integer32",
	}
	expected.ShouldBeEqual(t, 0, "BasicInt32_Default returns correct value -- with args", actual)
}

func Test_BasicInt32_IsAnyOf(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt32_IsAnyOf returns correct value -- with args", actual)
}

func Test_BasicInt32_IsAnyNamesOf(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt32_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_BasicInt32_GetValueByName(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
		"val":      int32(1),
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt32_GetValueByName returns correct value -- with args", actual)
}

func Test_BasicInt32_GetStringValue(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"val": bi.GetStringValue(0)}

	// Assert
	expected := args.Map{"val": "Invalid"}
	expected.ShouldBeEqual(t, 0, "BasicInt32_GetStringValue returns correct value -- with args", actual)
}

func Test_BasicInt32_Ranges(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt32_Ranges returns correct value -- with args", actual)
}

func Test_BasicInt32_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
	expected.ShouldBeEqual(t, 0, "BasicInt32_ToEnumJsonBytes returns correct value -- with args", actual)
}

func Test_BasicInt32_AppendPrependJoinValue(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BasicInt32_AppendPrependJoinValue returns correct value -- with args", actual)
}

func Test_BasicInt32_AppendPrependJoinNamer(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinNamer(".", mockNamer6{"B"}, mockNamer6{"A"})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "A.B"}
	expected.ShouldBeEqual(t, 0, "BasicInt32_AppendPrependJoinNamer returns correct value -- with args", actual)
}

func Test_BasicInt32_ToNumberString(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"notEmpty": bi.ToNumberString(1) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BasicInt32_ToNumberString returns correct value -- with args", actual)
}

func Test_BasicInt32_UnmarshallToValue(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
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
		"nilMapped":   int32(0),
		"nilErr":      true,
		"nilNoMapErr": true,
		"emptyVal":    int32(0),
		"emptyErr":    true,
		"dqVal":       int32(0),
		"dqErr":       true,
		"validVal":    int32(1),
		"validErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt32_UnmarshallToValue returns correct value -- with args", actual)
}

func Test_BasicInt32_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bi.ExpectingEnumValueError("Active", int32(1))
	hasErr := bi.ExpectingEnumValueError("Invalid", int32(1))
	parseErr := bi.ExpectingEnumValueError("NotExist", int32(1))

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
	expected.ShouldBeEqual(t, 0, "BasicInt32_ExpectingEnumValueError returns error -- with args", actual)
}

// ── newBasicInt32Creator paths ──

func Test_BasicInt32_CreateUsingMap(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.CreateUsingMap(
		"testInt32Enum",
		map[int32]string{0: "Invalid", 1: "Active"},
	)

	// Act
	actual := args.Map{
		"typeName": bi.TypeName(),
		"length": bi.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testInt32Enum",
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt32_CreateUsingMap returns correct value -- with args", actual)
}

func Test_BasicInt32_WithAliasMap(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.DefaultWithAliasMap(
		myInt32(0),
		[]string{"Invalid", "Active"},
		map[string]int32{"on": 1},
	)

	val, err := bi.GetValueByName("on")

	// Act
	actual := args.Map{
		"aliasVal": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"aliasVal": int32(1),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicInt32_WithAliasMap returns non-empty -- with args", actual)
}

func Test_BasicInt32_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.UsingFirstItemSliceAliasMap(
		myInt32(0),
		[]string{"Invalid", "Active"},
		map[string]int32{"on": 1},
	)

	// Act
	actual := args.Map{"typeName": bi.TypeName()}

	// Assert
	expected := args.Map{"typeName": "enumimpltests.myInt32"}
	expected.ShouldBeEqual(t, 0, "BasicInt32_UsingFirstItemSliceAliasMap returns correct value -- with args", actual)
}

func Test_BasicInt32_UsingTypeSlice(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice(
		"testInt32",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"typeName": bi.TypeName()}

	// Assert
	expected := args.Map{"typeName": "testInt32"}
	expected.ShouldBeEqual(t, 0, "BasicInt32_UsingTypeSlice returns correct value -- with args", actual)
}

func Test_BasicInt32_GetValueByString(t *testing.T) {
	// Arrange
	type myInt32 int32
	bi := enumimpl.New.BasicInt32.Default(
		myInt32(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"byName": bi.GetValueByString("Active")}

	// Assert
	expected := args.Map{"byName": int32(1)}
	expected.ShouldBeEqual(t, 0, "BasicInt32_GetValueByString returns correct value -- with args", actual)
}

// ══════════════════════════════════════════
// BasicUInt16 – all methods
// ══════════════════════════════════════════

func Test_BasicUInt16_Default(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
		"min":          uint16(0),
		"max":          uint16(2),
		"isValidRange": true,
		"outOfRange":   false,
		"toString":     "Active",
		"typeName":     "enumimpltests.myUInt16",
		"length":       3,
		"enumType":     "UnsignedInteger16",
	}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_Default returns correct value -- with args", actual)
}

func Test_BasicUInt16_IsAnyOf(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
	expected.ShouldBeEqual(t, 0, "BasicUInt16_IsAnyOf returns correct value -- with args", actual)
}

func Test_BasicUInt16_IsAnyNamesOf(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
	expected.ShouldBeEqual(t, 0, "BasicUInt16_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_BasicUInt16_GetValueByName(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
		"val":      uint16(1),
		"noErr":    true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_GetValueByName returns correct value -- with args", actual)
}

func Test_BasicUInt16_GetStringValue(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"val": bi.GetStringValue(0)}

	// Assert
	expected := args.Map{"val": "Invalid"}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_GetStringValue returns correct value -- with args", actual)
}

func Test_BasicUInt16_Ranges(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
	expected.ShouldBeEqual(t, 0, "BasicUInt16_Ranges returns correct value -- with args", actual)
}

func Test_BasicUInt16_ToEnumJsonBytes(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
	expected.ShouldBeEqual(t, 0, "BasicUInt16_ToEnumJsonBytes returns correct value -- with args", actual)
}

func Test_BasicUInt16_AppendPrependJoinValue(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_AppendPrependJoinValue returns correct value -- with args", actual)
}

func Test_BasicUInt16_AppendPrependJoinNamer(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
		[]string{"Invalid", "Active"},
	)

	result := bi.AppendPrependJoinNamer(".", mockNamer6{"B"}, mockNamer6{"A"})

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "A.B"}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_AppendPrependJoinNamer returns correct value -- with args", actual)
}

func Test_BasicUInt16_ToNumberString(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"notEmpty": bi.ToNumberString(1) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_ToNumberString returns correct value -- with args", actual)
}

func Test_BasicUInt16_UnmarshallToValue(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
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
		"nilMapped":   uint16(0),
		"nilErr":      true,
		"nilNoMapErr": true,
		"emptyVal":    uint16(0),
		"emptyErr":    true,
		"dqVal":       uint16(0),
		"dqErr":       true,
		"validVal":    uint16(1),
		"validErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_UnmarshallToValue returns correct value -- with args", actual)
}

func Test_BasicUInt16_ExpectingEnumValueError(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
		[]string{"Invalid", "Active"},
	)

	noErr := bi.ExpectingEnumValueError("Active", uint16(1))
	hasErr := bi.ExpectingEnumValueError("Invalid", uint16(1))
	parseErr := bi.ExpectingEnumValueError("NotExist", uint16(1))

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
	expected.ShouldBeEqual(t, 0, "BasicUInt16_ExpectingEnumValueError returns error -- with args", actual)
}

// ── newBasicUInt16Creator paths ──

func Test_BasicUInt16_CreateUsingMap(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicUInt16.CreateUsingMap(
		"testUInt16Enum",
		map[uint16]string{0: "Invalid", 1: "Active"},
	)

	// Act
	actual := args.Map{
		"typeName": bi.TypeName(),
		"length": bi.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testUInt16Enum",
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_CreateUsingMap returns correct value -- with args", actual)
}

func Test_BasicUInt16_WithAliasMap(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.DefaultWithAliasMap(
		myUInt16(0),
		[]string{"Invalid", "Active"},
		map[string]uint16{"on": 1},
	)

	val, err := bi.GetValueByName("on")

	// Act
	actual := args.Map{
		"aliasVal": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"aliasVal": uint16(1),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_WithAliasMap returns non-empty -- with args", actual)
}

func Test_BasicUInt16_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.UsingFirstItemSliceAliasMap(
		myUInt16(0),
		[]string{"Invalid", "Active"},
		map[string]uint16{"on": 1},
	)

	// Act
	actual := args.Map{"typeName": bi.TypeName()}

	// Assert
	expected := args.Map{"typeName": "enumimpltests.myUInt16"}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_UsingFirstItemSliceAliasMap returns correct value -- with args", actual)
}

func Test_BasicUInt16_UsingTypeSlice(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicUInt16.UsingTypeSlice(
		"testUInt16",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"typeName": bi.TypeName()}

	// Assert
	expected := args.Map{"typeName": "testUInt16"}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_UsingTypeSlice returns correct value -- with args", actual)
}

func Test_BasicUInt16_GetValueByString(t *testing.T) {
	// Arrange
	type myUInt16 uint16
	bi := enumimpl.New.BasicUInt16.Default(
		myUInt16(0),
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"byName": bi.GetValueByString("Active")}

	// Assert
	expected := args.Map{"byName": uint16(1)}
	expected.ShouldBeEqual(t, 0, "BasicUInt16_GetValueByString returns correct value -- with args", actual)
}

// ══════════════════════════════════════════
// FormatUsingFmt
// ══════════════════════════════════════════

type testFormatter struct {
	typeName, name, valueStr string
}

func (f testFormatter) TypeName() string   { return f.typeName }
func (f testFormatter) Name() string       { return f.name }
func (f testFormatter) ValueString() string { return f.valueStr }

func Test_FormatUsingFmt_FromBasicInt32Default(t *testing.T) {
	// Arrange
	f := testFormatter{typeName: "MyEnum", name: "Active", valueStr: "1"}
	result := enumimpl.FormatUsingFmt(f, "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{
		"notEmpty":    result != "",
		"containsAll": result != "",
	}

	// Assert
	expected := args.Map{
		"notEmpty":    true,
		"containsAll": true,
	}
	expected.ShouldBeEqual(t, 0, "FormatUsingFmt returns correct value -- with args", actual)
}

// ══════════════════════════════════════════
// DynamicMap – extended coverage
// ══════════════════════════════════════════

func Test_DynamicMap_Set(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.Set("b", 2)
	isUpdate := dm.Set("a", 99)

	// Act
	actual := args.Map{
		"isNew":    isNew,
		"isUpdate": isUpdate,
		"len":      dm.Length(),
	}

	// Assert
	expected := args.Map{
		"isNew":    true,
		"isUpdate": false,
		"len":      2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_Set returns correct value -- with args", actual)
}

func Test_DynamicMap_AddNewOnly(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	added := dm.AddNewOnly("b", 2)
	notAdded := dm.AddNewOnly("a", 99)

	// Act
	actual := args.Map{
		"added":    added,
		"notAdded": notAdded,
		"aVal":     dm["a"],
	}

	// Assert
	expected := args.Map{
		"added":    true,
		"notAdded": false,
		"aVal":     1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_AddNewOnly returns correct value -- with args", actual)
}

func Test_DynamicMap_Add(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.Add("b", 2)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"len":    dm.Length(),
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"len":    2,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_Add returns correct value -- with args", actual)
}

func Test_DynamicMap_LastIndex(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{
		"lastIndex": dm.LastIndex(),
		"hasIndex":  dm.HasIndex(0),
		"noIndex":   dm.HasIndex(5),
	}

	// Assert
	expected := args.Map{
		"lastIndex": 1,
		"hasIndex":  true,
		"noIndex":   false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_LastIndex returns correct value -- with args", actual)
}

func Test_DynamicMap_AllValuesIntegers_FromBasicInt32Default(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	vals := dm.AllValuesIntegers()

	// Act
	actual := args.Map{"len": len(vals)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicMap_AllValuesIntegers returns non-empty -- with args", actual)
}

func Test_DynamicMap_KeyValue(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	val, found := dm.KeyValue("a")
	_, notFound := dm.KeyValue("z")

	// Act
	actual := args.Map{
		"val":      val,
		"found":    found,
		"notFound": notFound,
	}

	// Assert
	expected := args.Map{
		"val":      1,
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_KeyValue returns correct value -- with args", actual)
}

func Test_DynamicMap_KeyValueString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "hello"}
	val, found := dm.KeyValueString("a")
	_, notFound := dm.KeyValueString("z")

	// Act
	actual := args.Map{
		"val":      val,
		"found":    found,
		"notFound": notFound,
	}

	// Assert
	expected := args.Map{
		"val":      "hello",
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_KeyValueString returns non-empty -- with args", actual)
}

func Test_DynamicMap_KeyValueInt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 42, "b": "notInt"}
	val, found, failed := dm.KeyValueInt("a")
	_, notFound, _ := dm.KeyValueInt("z")

	// Act
	actual := args.Map{
		"val":      val,
		"found":    found,
		"failed":   failed,
		"notFound": notFound,
	}

	// Assert
	expected := args.Map{
		"val":      42,
		"found":    true,
		"failed":   false,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_KeyValueInt returns correct value -- with args", actual)
}

func Test_DynamicMap_KeyValueByte(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(5)}
	val, found, failed := dm.KeyValueByte("a")
	_, notFound, _ := dm.KeyValueByte("z")

	// Act
	actual := args.Map{
		"val":      val,
		"found":    found,
		"failed":   failed,
		"notFound": notFound,
	}

	// Assert
	expected := args.Map{
		"val":      byte(5),
		"found":    true,
		"failed":   false,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_KeyValueByte returns correct value -- with args", actual)
}

func Test_DynamicMap_KeyValueIntDefault(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 42}
	val := dm.KeyValueIntDefault("a")
	missing := dm.KeyValueIntDefault("z")

	// Act
	actual := args.Map{
		"val": val,
		"missing": missing < 0,
	}

	// Assert
	expected := args.Map{
		"val": 42,
		"missing": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_KeyValueIntDefault returns correct value -- with args", actual)
}

func Test_DynamicMap_IsValueTypeOf(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "hello"}
	isStr := dm.IsValueTypeOf(reflect.TypeOf(""))

	// Act
	actual := args.Map{"checked": isStr}

	// Assert
	expected := args.Map{"checked": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsValueTypeOf returns correct value -- with args", actual)
}

func Test_DynamicMap_IsRawEqual(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	same := dm.IsRawEqual(true, map[string]any{"a": 1})
	diff := dm.IsRawEqual(false, map[string]any{"a": 2})
	diffLen := dm.IsRawEqual(false, map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{
		"same": same,
		"diff": diff,
		"diffLen": diffLen,
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
		"diffLen": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsRawEqual returns correct value -- with args", actual)
}

func Test_DynamicMap_IsRawMismatch(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	mismatch := dm.IsRawMismatch(false, map[string]any{"a": 2})

	// Act
	actual := args.Map{"mismatch": mismatch}

	// Assert
	expected := args.Map{"mismatch": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsRawMismatch returns correct value -- with args", actual)
}

func Test_DynamicMap_IsKeysEqualOnly(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1, "b": 2}
	same := dm.IsKeysEqualOnly(map[string]any{"a": 99, "b": 88})
	diff := dm.IsKeysEqualOnly(map[string]any{"a": 1, "c": 2})
	diffLen := dm.IsKeysEqualOnly(map[string]any{"a": 1})

	var nilDm *enumimpl.DynamicMap
	bothNil := nilDm.IsKeysEqualOnly(nil)
	leftNil := nilDm.IsKeysEqualOnly(map[string]any{"a": 1})

	// Act
	actual := args.Map{
		"same":    same,
		"diff":    diff,
		"diffLen": diffLen,
		"bothNil": bothNil,
		"leftNil": leftNil,
	}

	// Assert
	expected := args.Map{
		"same":    true,
		"diff":    false,
		"diffLen": false,
		"bothNil": true,
		"leftNil": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsKeysEqualOnly returns correct value -- with args", actual)
}

func Test_DynamicMap_DiffRaw(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1, "b": 2}
	diff := dm.DiffRaw(false, map[string]any{"a": 1, "b": 3})

	// Act
	actual := args.Map{"hasDiff": diff.Length() > 0}

	// Assert
	expected := args.Map{"hasDiff": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffRaw returns correct value -- with args", actual)
}

func Test_DynamicMap_DiffRawNilCases(t *testing.T) {
	// Arrange
	var nilDm *enumimpl.DynamicMap
	// both nil
	d1 := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	// left nil
	d2 := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	// right nil
	dm := &enumimpl.DynamicMap{"a": 1}
	d3 := dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{
		"bothNilLen": d1.Length(),
		"leftNilLen": d2.Length(),
		"rightNil":   d3.Length(),
	}

	// Assert
	expected := args.Map{
		"bothNilLen": 0,
		"leftNilLen": 1,
		"rightNil":   1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffRawNilCases returns nil -- with args", actual)
}

func Test_DynamicMap_DiffJsonMessage(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noDiff := dm.DiffJsonMessage(false, map[string]any{"a": 1})
	hasDiff := dm.DiffJsonMessage(false, map[string]any{"a": 2})

	// Act
	actual := args.Map{
		"noDiff":  noDiff == "",
		"hasDiff": hasDiff != "",
	}

	// Assert
	expected := args.Map{
		"noDiff":  true,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffJsonMessage returns correct value -- with args", actual)
}

func Test_DynamicMap_ShouldDiffMessage(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noDiff := dm.ShouldDiffMessage(false, "title", map[string]any{"a": 1})
	hasDiff := dm.ShouldDiffMessage(false, "title", map[string]any{"a": 2})

	// Act
	actual := args.Map{
		"noDiff":  noDiff == "",
		"hasDiff": hasDiff != "",
	}

	// Assert
	expected := args.Map{
		"noDiff":  true,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_ShouldDiffMessage returns correct value -- with args", actual)
}

func Test_DynamicMap_ExpectingMessage(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noMsg := dm.ExpectingMessage("title", map[string]any{"a": 1})
	hasMsg := dm.ExpectingMessage("title", map[string]any{"a": 2})

	// Act
	actual := args.Map{
		"noMsg":  noMsg == "",
		"hasMsg": hasMsg != "",
	}

	// Assert
	expected := args.Map{
		"noMsg":  true,
		"hasMsg": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_ExpectingMessage returns correct value -- with args", actual)
}

func Test_DynamicMap_Strings(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	strs := dm.Strings()
	str := dm.String()

	emptyDm := enumimpl.DynamicMap{}
	emptyStrs := emptyDm.Strings()

	// Act
	actual := args.Map{
		"len":      len(strs),
		"notEmpty": str != "",
		"emptyLen": len(emptyStrs),
	}

	// Assert
	expected := args.Map{
		"len":      2,
		"notEmpty": true,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_Strings returns correct value -- with args", actual)
}

func Test_DynamicMap_StringsUsingFmt(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	strs := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return fmt.Sprintf("%s=%v", key, val)
	})

	emptyDm := enumimpl.DynamicMap{}
	emptyStrs := emptyDm.StringsUsingFmt(func(index int, key string, val any) string {
		return ""
	})

	// Act
	actual := args.Map{
		"len":      len(strs),
		"emptyLen": len(emptyStrs),
	}

	// Assert
	expected := args.Map{
		"len":      1,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_StringsUsingFmt returns correct value -- with args", actual)
}

func Test_DynamicMap_IsStringEqual(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	str := dm.String()

	// Act
	actual := args.Map{
		"same": dm.IsStringEqual(str),
		"diff": dm.IsStringEqual("other"),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_IsStringEqual returns correct value -- with args", actual)
}

func Test_DynamicMap_Serialize(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	bytes, err := dm.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"noErr":    err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr":    true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_Serialize returns correct value -- with args", actual)
}

func Test_DynamicMap_ConcatNew(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	other := enumimpl.DynamicMap{"b": 2, "a": 99}

	overridden := dm.ConcatNew(true, other)
	notOverridden := dm.ConcatNew(false, other)
	bothEmpty := (enumimpl.DynamicMap{}).ConcatNew(false, enumimpl.DynamicMap{})

	// Act
	actual := args.Map{
		"overriddenA":  overridden["a"],
		"overriddenB":  overridden["b"],
		"notOverA":     notOverridden["a"],
		"bothEmptyLen": bothEmpty.Length(),
	}

	// Assert
	expected := args.Map{
		"overriddenA":  99,
		"overriddenB":  2,
		"notOverA":     1,
		"bothEmptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_ConcatNew returns correct value -- with args", actual)
}

func Test_DynamicMap_ConvMaps(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}

	convSI := dm.ConvMapStringInteger()
	convIS := dm.ConvMapIntegerString()
	convBS := dm.ConvMapByteString()
	convI8 := dm.ConvMapInt8String()
	convI16 := dm.ConvMapInt16String()
	convI32 := dm.ConvMapInt32String()
	convU16 := dm.ConvMapUInt16String()
	convI64 := dm.ConvMapInt64String()
	convSS := dm.ConvMapStringString()

	emptyDm := enumimpl.DynamicMap{}
	emptyBS := emptyDm.ConvMapByteString()
	emptyI8 := emptyDm.ConvMapInt8String()
	emptyI16 := emptyDm.ConvMapInt16String()
	emptyI32 := emptyDm.ConvMapInt32String()
	emptyU16 := emptyDm.ConvMapUInt16String()
	emptyI64 := emptyDm.ConvMapInt64String()
	emptySS := emptyDm.ConvMapStringString()
	emptySI := emptyDm.ConvMapStringInteger()
	emptyIS := emptyDm.ConvMapIntegerString()

	// Act
	actual := args.Map{
		"siLen":      len(convSI),
		"isLen":      len(convIS),
		"bsLen":      len(convBS),
		"i8Len":      len(convI8),
		"i16Len":     len(convI16),
		"i32Len":     len(convI32),
		"u16Len":     len(convU16),
		"i64Len":     len(convI64),
		"ssLen":      len(convSS),
		"emptyBS":    len(emptyBS),
		"emptyI8":    len(emptyI8),
		"emptyI16":   len(emptyI16),
		"emptyI32":   len(emptyI32),
		"emptyU16":   len(emptyU16),
		"emptyI64":   len(emptyI64),
		"emptySS":    len(emptySS),
		"emptySI":    len(emptySI),
		"emptyIS":    len(emptyIS),
	}

	// Assert
	expected := args.Map{
		"siLen":      2,
		"isLen":      2,
		"bsLen":      len(convBS),
		"i8Len":      len(convI8),
		"i16Len":     len(convI16),
		"i32Len":     len(convI32),
		"u16Len":     len(convU16),
		"i64Len":     len(convI64),
		"ssLen":      len(convSS),
		"emptyBS":    0,
		"emptyI8":    0,
		"emptyI16":   0,
		"emptyI32":   0,
		"emptyU16":   0,
		"emptyI64":   0,
		"emptySS":    0,
		"emptySI":    0,
		"emptyIS":    0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_ConvMaps returns correct value -- with args", actual)
}

func Test_DynamicMap_BasicFromDynamicMap(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}

	bb := dm.BasicByte("testByte")
	bi8 := dm.BasicInt8("testInt8")
	bi16 := dm.BasicInt16("testInt16")
	bi32 := dm.BasicInt32("testInt32")
	bu16 := dm.BasicUInt16("testUInt16")
	bs := dm.BasicString("testString")

	// Act
	actual := args.Map{
		"bbNotNil":  bb != nil,
		"bi8NotNil": bi8 != nil,
		"bi16NotNil": bi16 != nil,
		"bi32NotNil": bi32 != nil,
		"bu16NotNil": bu16 != nil,
		"bsNotNil":  bs != nil,
	}

	// Assert
	expected := args.Map{
		"bbNotNil":  true,
		"bi8NotNil": true,
		"bi16NotNil": true,
		"bi32NotNil": true,
		"bu16NotNil": true,
		"bsNotNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_BasicFromDynamicMap returns correct value -- with args", actual)
}

func Test_DynamicMap_BasicWithAliasFromDynamicMap(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}

	bb := dm.BasicByteUsingAliasMap("t", map[string]byte{"on": 1})
	bi8 := dm.BasicInt8UsingAliasMap("t", map[string]int8{"on": 1})
	bi16 := dm.BasicInt16UsingAliasMap("t", map[string]int16{"on": 1})
	bi32 := dm.BasicInt32UsingAliasMap("t", map[string]int32{"on": 1})
	bu16 := dm.BasicUInt16UsingAliasMap("t", map[string]uint16{"on": 1})
	bs := dm.BasicStringUsingAliasMap("t", map[string]string{"on": "Active"})

	// Act
	actual := args.Map{
		"bbNotNil":  bb != nil,
		"bi8NotNil": bi8 != nil,
		"bi16NotNil": bi16 != nil,
		"bi32NotNil": bi32 != nil,
		"bu16NotNil": bu16 != nil,
		"bsNotNil":  bs != nil,
	}

	// Assert
	expected := args.Map{
		"bbNotNil":  true,
		"bi8NotNil": true,
		"bi16NotNil": true,
		"bi32NotNil": true,
		"bu16NotNil": true,
		"bsNotNil":  true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_BasicWithAliasFromDynamicMap returns non-empty -- with args", actual)
}

func Test_DynamicMap_DiffLeftRight(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1, "b": 2}
	lDiff, rDiff := dm.DiffRawLeftRightUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl,
		false,
		map[string]any{"a": 1, "c": 3},
	)

	// Act
	actual := args.Map{
		"lDiffLen": lDiff.Length(),
		"rDiffLen": rDiff.Length(),
	}

	// Assert
	expected := args.Map{
		"lDiffLen": 1,
		"rDiffLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffLeftRight returns correct value -- with args", actual)
}

func Test_DynamicMap_DiffLeftRightNilCases(t *testing.T) {
	// Arrange
	var nilDm *enumimpl.DynamicMap
	l1, r1 := nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	l2, r2 := nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	dm := &enumimpl.DynamicMap{"a": 1}
	l3, r3 := dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{
		"bothNilL": l1.Length(),
		"bothNilR": r1.Length(),
		"leftNilL": l2.Length(),
		"leftNilR": r2.Length(),
		"rightNilL": l3.Length(),
		"rightNilR": r3.Length(),
	}

	// Assert
	expected := args.Map{
		"bothNilL": 0,
		"bothNilR": 0,
		"leftNilL": 1,
		"leftNilR": 0,
		"rightNilL": 1,
		"rightNilR": 0,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffLeftRightNilCases returns nil -- with args", actual)
}

func Test_DynamicMap_DiffJsonMessageLeftRight(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noDiff := dm.DiffJsonMessageLeftRight(false, map[string]any{"a": 1})
	hasDiff := dm.DiffJsonMessageLeftRight(false, map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{
		"noDiff":  noDiff == "",
		"hasDiff": hasDiff != "",
	}

	// Assert
	expected := args.Map{
		"noDiff":  true,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffJsonMessageLeftRight returns correct value -- with args", actual)
}

func Test_DynamicMap_ShouldDiffUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noDiff := dm.ShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 1})
	hasDiff := dm.ShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 2})

	// Act
	actual := args.Map{
		"noDiff": noDiff == "",
		"hasDiff": hasDiff != "",
	}

	// Assert
	expected := args.Map{
		"noDiff": true,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_ShouldDiffUsingDifferChecker returns correct value -- with args", actual)
}

func Test_DynamicMap_ShouldDiffLeftRightUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noDiff := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 1})
	hasDiff := dm.ShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{
		"noDiff": noDiff == "",
		"hasDiff": hasDiff != "",
	}

	// Assert
	expected := args.Map{
		"noDiff": true,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_ShouldDiffLeftRightUsingDifferChecker returns correct value -- with args", actual)
}

func Test_DynamicMap_MapIntegerString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}
	rm, sorted := dm.MapIntegerString()

	emptyDm := enumimpl.DynamicMap{}
	emptyRm, emptySorted := emptyDm.MapIntegerString()

	strDm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	strRm, strSorted := strDm.MapIntegerString()

	// Act
	actual := args.Map{
		"rmLen":       len(rm),
		"sortedLen":   len(sorted),
		"emptyRmLen":  len(emptyRm),
		"emptySorted": len(emptySorted),
		"strRmLen":    len(strRm),
		"strSorted":   len(strSorted),
	}

	// Assert
	expected := args.Map{
		"rmLen":       len(rm),
		"sortedLen":   len(sorted),
		"emptyRmLen":  0,
		"emptySorted": 0,
		"strRmLen":    len(strRm),
		"strSorted":   len(strSorted),
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_MapIntegerString returns correct value -- with args", actual)
}

// ── BasicString creator paths ──

func Test_BasicString_CreateUsingNamesSpread(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingNamesSpread(
		"testStrEnum",
		"Invalid", "Active",
	)

	// Act
	actual := args.Map{
		"typeName": bs.TypeName(),
		"length":   bs.Length(),
	}

	// Assert
	expected := args.Map{
		"typeName": "testStrEnum",
		"length":   2,
	}
	expected.ShouldBeEqual(t, 0, "BasicString_CreateUsingNamesSpread returns correct value -- with args", actual)
}

func Test_BasicString_CreateUsingNamesMinMax(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateUsingNamesMinMax(
		"testStrEnum",
		[]string{"Invalid", "Active"},
		"Active",
		"Invalid",
	)

	// Act
	actual := args.Map{
		"min": bs.Min(),
		"max": bs.Max(),
	}

	// Assert
	expected := args.Map{
		"min": "Active",
		"max": "Invalid",
	}
	expected.ShouldBeEqual(t, 0, "BasicString_CreateUsingNamesMinMax returns correct value -- with args", actual)
}

func Test_BasicString_UsingFirstItemSliceAllCases(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.UsingFirstItemSliceAllCases(
		"testItem",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"typeName": bs.TypeName()}

	// Assert
	expected := args.Map{"typeName": "string"}
	expected.ShouldBeEqual(t, 0, "BasicString_UsingFirstItemSliceAllCases returns correct value -- with args", actual)
}

func Test_BasicString_UsingFirstItemSliceCaseOptions(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.UsingFirstItemSliceCaseOptions(
		false,
		"testItem",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"typeName": bs.TypeName()}

	// Assert
	expected := args.Map{"typeName": "string"}
	expected.ShouldBeEqual(t, 0, "BasicString_UsingFirstItemSliceCaseOptions returns correct value -- with args", actual)
}

func Test_BasicString_CreateDefault(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.CreateDefault(
		"testItem",
		[]string{"Invalid", "Active"},
	)

	// Act
	actual := args.Map{"typeName": bs.TypeName()}

	// Assert
	expected := args.Map{"typeName": "string"}
	expected.ShouldBeEqual(t, 0, "BasicString_CreateDefault returns correct value -- with args", actual)
}

func Test_BasicString_ToEnumString(t *testing.T) {
	// Arrange
	bs := enumimpl.New.BasicString.Create(
		"testStringEnum",
		[]string{"Invalid", "Active"},
	)

	result := bs.ToEnumString("Active")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "Active"}
	expected.ShouldBeEqual(t, 0, "BasicString_ToEnumString_passthrough returns correct value -- with args", actual)
}

// ── LeftRightDiffChecker – same values ──

func Test_LeftRightDiffChecker_SameValues(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetSingleDiffResult(true, "same", "same")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "LeftRightDiffChecker_SameValues returns non-empty -- with args", actual)
}

// ── DynamicMap LogExpectingMessage (no crash) ──

func Test_DynamicMap_LogExpectingMessage(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("title", map[string]any{"a": 1})
	dm.LogExpectingMessage("title", map[string]any{"a": 2})

	// Act
	actual := args.Map{"noCrash": true}

	// Assert
	expected := args.Map{"noCrash": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap_LogExpectingMessage returns correct value -- with args", actual)
}

func Test_DynamicMap_LogShouldDiff(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	dm.LogShouldDiffMessage(false, "title", map[string]any{"a": 1})
	dm.LogShouldDiffMessage(false, "title", map[string]any{"a": 2})
	dm.LogShouldDiffLeftRightMessage(false, "title", map[string]any{"a": 1})
	dm.LogShouldDiffLeftRightMessage(false, "title", map[string]any{"a": 1, "b": 2})
	dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 1})
	dm.LogShouldDiffMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 2})
	dm.LogShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 1})
	dm.LogShouldDiffLeftRightMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, "title", map[string]any{"a": 1, "b": 2})

	// Act
	actual := args.Map{"noCrash": true}

	// Assert
	expected := args.Map{"noCrash": true}
	expected.ShouldBeEqual(t, 0, "DynamicMap_LogShouldDiff returns correct value -- with args", actual)
}

// ── DynamicMap DiffJsonMessageUsingDifferChecker ──

func Test_DynamicMap_DiffJsonMessageUsingDifferChecker(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"a": 1}
	noDiff := dm.DiffJsonMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	hasDiff := dm.DiffJsonMessageUsingDifferChecker(
		enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 2})

	// Act
	actual := args.Map{
		"noDiff": noDiff == "",
		"hasDiff": hasDiff != "",
	}

	// Assert
	expected := args.Map{
		"noDiff": true,
		"hasDiff": true,
	}
	expected.ShouldBeEqual(t, 0, "DynamicMap_DiffJsonMessageUsingDifferChecker returns correct value -- with args", actual)
}

// ── DynamicMap with string values (SortedKeyAnyValues string path) ──

func Test_DynamicMap_SortedKeyAnyValues_StringValues(t *testing.T) {
	// Arrange
	dm := &enumimpl.DynamicMap{"b": "world", "a": "hello"}
	result := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicMap_SortedKeyAnyValues_StringValues returns non-empty -- with args", actual)
}

type mockNamer6 struct{ name string }

func (m mockNamer6) Name() string { return m.name }
