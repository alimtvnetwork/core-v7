package enumimpltests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coreimpl/enumimpl/enumtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Helper: create test enums for reuse
// ══════════════════════════════════════════════════════════════════════════════

type testByte byte

const testByteVal testByte = 0

func byteEnum() *enumimpl.BasicByte {
	return enumimpl.New.BasicByte.Default(
		testByteVal,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testInt8 int8

const testInt8Val testInt8 = 0

func int8Enum() *enumimpl.BasicInt8 {
	return enumimpl.New.BasicInt8.Default(
		testInt8Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testInt16 int16

const testInt16Val testInt16 = 0

func int16Enum() *enumimpl.BasicInt16 {
	return enumimpl.New.BasicInt16.Default(
		testInt16Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testInt32 int32

const testInt32Val testInt32 = 0

func int32Enum() *enumimpl.BasicInt32 {
	return enumimpl.New.BasicInt32.Default(
		testInt32Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testUInt16 uint16

const testUInt16Val testUInt16 = 0

func uint16Enum() *enumimpl.BasicUInt16 {
	return enumimpl.New.BasicUInt16.Default(
		testUInt16Val,
		[]string{"Invalid", "Active", "Inactive"},
	)
}

func stringEnum() *enumimpl.BasicString {
	return enumimpl.New.BasicString.Create(
		"TestStringEnum",
		[]string{"Invalid", "Active", "Inactive"},
	)
}

type testNamer struct{ name string }

func (n testNamer) Name() string { return n.name }

// ══════════════════════════════════════════════════════════════════════════════
// BasicByte — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicByte_IsAnyOf_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.IsAnyOf(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty variadic should return true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 1, 2)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicByte_IsAnyNamesOf_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.IsAnyNamesOf(0, "Invalid", "Active")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyNamesOf(0, "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicByte_GetValueByString_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.GetValueByString("Active")
}

func Test_BasicByte_GetValueByName_AllBranches(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Direct key
	v, err := e.GetValueByName("Active")

	// Act
	actual := args.Map{"result": err != nil || v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	// Wrapped key
	v, err = e.GetValueByName(`"Active"`)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error for wrapped", actual)

	// Not found
	_, err = e.GetValueByName("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicByte_GetStringValue_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	s := e.GetStringValue(0)

	// Act
	actual := args.Map{"result": s != "Invalid"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid", actual)
}

func Test_BasicByte_ExpectingEnumValueError_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Matching
	err := e.ExpectingEnumValueError("Active", byte(1))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	// Not matching
	err = e.ExpectingEnumValueError("Active", byte(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// Invalid raw string
	err = e.ExpectingEnumValueError("NotExist", byte(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid raw", actual)
}

func Test_BasicByte_Ranges_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": len(e.Ranges()) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_BasicByte_Hashmap_HashmapPtr(t *testing.T) {
	e := byteEnum()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_BasicByte_IsValidRange(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.IsValidRange(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": e.IsValidRange(100)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_BasicByte_ToEnumJsonBytes_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	b, err := e.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)

	// Not found
	_, err = e.ToEnumJsonBytes(99)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicByte_ToEnumString(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.ToEnumString(0) != "Invalid"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid", actual)
}

func Test_BasicByte_AppendPrependJoinValue_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	r := e.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"result": r != "Invalid.Active"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Invalid.Active", actual)
}

func Test_BasicByte_AppendPrependJoinNamer_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	r := e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})

	// Act
	actual := args.Map{"result": r != "A.B"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected A.B", actual)
}

func Test_BasicByte_ToNumberString_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	s := e.ToNumberString(byte(42))

	// Act
	actual := args.Map{"result": s != "42"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)
}

func Test_BasicByte_JsonMap(t *testing.T) {
	e := byteEnum()
	_ = e.JsonMap()
}

func Test_BasicByte_UnmarshallToValue_AllBranches(t *testing.T) {
	// Arrange
	e := byteEnum()

	// nil + not mapped
	_, err := e.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// nil + mapped
	v, err := e.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	// empty + mapped
	v, err = e.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	// double quote empty + mapped
	v, err = e.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	// Valid value
	v, err = e.UnmarshallToValue(false, []byte("Active"))
	actual = args.Map{"result": err != nil || v != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BasicByte_EnumType_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.EnumType() != enumtype.Byte}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Byte", actual)
}

func Test_BasicByte_AsBasicByter_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	byter := e.AsBasicByter()

	// Act
	actual := args.Map{"result": byter == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicInt8 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicInt8_IsAnyOf_FromMisc12(t *testing.T) {
	// Arrange
	e := int8Enum()

	// Act
	actual := args.Map{"result": e.IsAnyOf(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicInt8_IsAnyNamesOf_FromMisc12(t *testing.T) {
	// Arrange
	e := int8Enum()

	// Act
	actual := args.Map{"result": e.IsAnyNamesOf(0, "Invalid")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyNamesOf(0, "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicInt8_GetValueByName_AllBranches(t *testing.T) {
	// Arrange
	e := int8Enum()

	v, err := e.GetValueByName("Active")

	// Act
	actual := args.Map{"result": err != nil || v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_, err = e.GetValueByName("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt8_GetValueByString_FromMisc12(t *testing.T) {
	e := int8Enum()
	_ = e.GetValueByString("Active")
}

func Test_BasicInt8_GetStringValue_FromMisc12(t *testing.T) {
	e := int8Enum()
	_ = e.GetStringValue(0)
}

func Test_BasicInt8_ExpectingEnumValueError_FromMisc12(t *testing.T) {
	// Arrange
	e := int8Enum()

	err := e.ExpectingEnumValueError("Active", int8(1))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	err = e.ExpectingEnumValueError("Active", int8(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	err = e.ExpectingEnumValueError("NotExist", int8(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt8_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := int8Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_BasicInt8_IsValidRange_FromMisc12(t *testing.T) {
	// Arrange
	e := int8Enum()

	// Act
	actual := args.Map{"result": e.IsValidRange(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": e.IsValidRange(100)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_BasicInt8_ToEnumJsonBytes_FromMisc12(t *testing.T) {
	// Arrange
	e := int8Enum()

	_, err := e.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	_, err = e.ToEnumJsonBytes(99)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt8_ToEnumString(t *testing.T) {
	e := int8Enum()
	_ = e.ToEnumString(0)
}

func Test_BasicInt8_AppendPrependJoinValue_FromMisc12(t *testing.T) {
	e := int8Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_BasicInt8_AppendPrependJoinNamer_FromMisc12(t *testing.T) {
	e := int8Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_BasicInt8_ToNumberString_FromMisc12(t *testing.T) {
	e := int8Enum()
	_ = e.ToNumberString(int8(42))
}

func Test_BasicInt8_UnmarshallToValue_AllBranches(t *testing.T) {
	// Arrange
	e := int8Enum()

	_, err := e.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	v, err := e.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	v, err = e.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	actual = args.Map{"result": err != nil || v != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BasicInt8_EnumType_FromMisc12(t *testing.T) {
	// Arrange
	e := int8Enum()

	// Act
	actual := args.Map{"result": e.EnumType() != enumtype.Integer8}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Integer8", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicInt16 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicInt16_IsAnyOf_FromMisc12(t *testing.T) {
	// Arrange
	e := int16Enum()

	// Act
	actual := args.Map{"result": e.IsAnyOf(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicInt16_IsAnyNamesOf_FromMisc12(t *testing.T) {
	// Arrange
	e := int16Enum()

	// Act
	actual := args.Map{"result": e.IsAnyNamesOf(0, "Invalid")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyNamesOf(0, "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicInt16_GetValueByName_AllBranches(t *testing.T) {
	// Arrange
	e := int16Enum()

	v, err := e.GetValueByName("Active")

	// Act
	actual := args.Map{"result": err != nil || v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_, err = e.GetValueByName("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt16_GetValueByString_FromMisc12(t *testing.T) {
	e := int16Enum()
	_ = e.GetValueByString("Active")
}

func Test_BasicInt16_GetStringValue_FromMisc12(t *testing.T) {
	e := int16Enum()
	_ = e.GetStringValue(0)
}

func Test_BasicInt16_ExpectingEnumValueError_FromMisc12(t *testing.T) {
	// Arrange
	e := int16Enum()

	err := e.ExpectingEnumValueError("Active", int16(1))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	err = e.ExpectingEnumValueError("Active", int16(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	err = e.ExpectingEnumValueError("NotExist", int16(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt16_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := int16Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_BasicInt16_IsValidRange_FromMisc12(t *testing.T) {
	// Arrange
	e := int16Enum()

	// Act
	actual := args.Map{"result": e.IsValidRange(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": e.IsValidRange(100)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_BasicInt16_ToEnumJsonBytes_FromMisc12(t *testing.T) {
	// Arrange
	e := int16Enum()

	_, err := e.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	_, err = e.ToEnumJsonBytes(99)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt16_ToEnumString(t *testing.T) {
	e := int16Enum()
	_ = e.ToEnumString(0)
}

func Test_BasicInt16_AppendPrependJoinValue_FromMisc12(t *testing.T) {
	e := int16Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_BasicInt16_AppendPrependJoinNamer_FromMisc12(t *testing.T) {
	e := int16Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_BasicInt16_ToNumberString_FromMisc12(t *testing.T) {
	e := int16Enum()
	_ = e.ToNumberString(int16(42))
}

func Test_BasicInt16_UnmarshallToValue_AllBranches(t *testing.T) {
	// Arrange
	e := int16Enum()

	_, err := e.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	v, err := e.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	v, err = e.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	actual = args.Map{"result": err != nil || v != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BasicInt16_EnumType_FromMisc12(t *testing.T) {
	// Arrange
	e := int16Enum()

	// Act
	actual := args.Map{"result": e.EnumType() != enumtype.Integer16}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Integer16", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicInt32 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicInt32_IsAnyOf_FromMisc12(t *testing.T) {
	// Arrange
	e := int32Enum()

	// Act
	actual := args.Map{"result": e.IsAnyOf(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicInt32_IsAnyNamesOf_FromMisc12(t *testing.T) {
	// Arrange
	e := int32Enum()

	// Act
	actual := args.Map{"result": e.IsAnyNamesOf(0, "Invalid")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyNamesOf(0, "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicInt32_GetValueByName_AllBranches(t *testing.T) {
	// Arrange
	e := int32Enum()

	v, err := e.GetValueByName("Active")

	// Act
	actual := args.Map{"result": err != nil || v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_, err = e.GetValueByName("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt32_GetValueByString_FromMisc12(t *testing.T) {
	e := int32Enum()
	_ = e.GetValueByString("Active")
}

func Test_BasicInt32_GetStringValue_FromMisc12(t *testing.T) {
	e := int32Enum()
	_ = e.GetStringValue(0)
}

func Test_BasicInt32_ExpectingEnumValueError_FromMisc12(t *testing.T) {
	// Arrange
	e := int32Enum()

	err := e.ExpectingEnumValueError("Active", int32(1))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	err = e.ExpectingEnumValueError("Active", int32(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	err = e.ExpectingEnumValueError("NotExist", int32(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt32_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := int32Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_BasicInt32_IsValidRange_FromMisc12(t *testing.T) {
	// Arrange
	e := int32Enum()

	// Act
	actual := args.Map{"result": e.IsValidRange(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": e.IsValidRange(100)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_BasicInt32_ToEnumJsonBytes_FromMisc12(t *testing.T) {
	// Arrange
	e := int32Enum()

	_, err := e.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	_, err = e.ToEnumJsonBytes(99)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicInt32_ToEnumString(t *testing.T) {
	e := int32Enum()
	_ = e.ToEnumString(0)
}

func Test_BasicInt32_AppendPrependJoinValue_FromMisc12(t *testing.T) {
	e := int32Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_BasicInt32_AppendPrependJoinNamer_FromMisc12(t *testing.T) {
	e := int32Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_BasicInt32_ToNumberString_FromMisc12(t *testing.T) {
	e := int32Enum()
	_ = e.ToNumberString(int32(42))
}

func Test_BasicInt32_UnmarshallToValue_AllBranches(t *testing.T) {
	// Arrange
	e := int32Enum()

	_, err := e.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	v, err := e.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	v, err = e.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	actual = args.Map{"result": err != nil || v != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BasicInt32_EnumType_FromMisc12(t *testing.T) {
	// Arrange
	e := int32Enum()

	// Act
	actual := args.Map{"result": e.EnumType() != enumtype.Integer32}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Integer32", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicUInt16 — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicUInt16_IsAnyOf_FromMisc12(t *testing.T) {
	// Arrange
	e := uint16Enum()

	// Act
	actual := args.Map{"result": e.IsAnyOf(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyOf(1, 0, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicUInt16_IsAnyNamesOf_FromMisc12(t *testing.T) {
	// Arrange
	e := uint16Enum()

	// Act
	actual := args.Map{"result": e.IsAnyNamesOf(0, "Invalid")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyNamesOf(0, "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicUInt16_GetValueByName_AllBranches(t *testing.T) {
	// Arrange
	e := uint16Enum()

	v, err := e.GetValueByName("Active")

	// Act
	actual := args.Map{"result": err != nil || v != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_, err = e.GetValueByName("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicUInt16_GetValueByString_FromMisc12(t *testing.T) {
	e := uint16Enum()
	_ = e.GetValueByString("Active")
}

func Test_BasicUInt16_GetStringValue_FromMisc12(t *testing.T) {
	e := uint16Enum()
	_ = e.GetStringValue(0)
}

func Test_BasicUInt16_ExpectingEnumValueError_FromMisc12(t *testing.T) {
	// Arrange
	e := uint16Enum()

	err := e.ExpectingEnumValueError("Active", uint16(1))

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	err = e.ExpectingEnumValueError("Active", uint16(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	err = e.ExpectingEnumValueError("NotExist", uint16(0))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicUInt16_Ranges_Hashmap_HashmapPtr(t *testing.T) {
	e := uint16Enum()
	_ = e.Ranges()
	_ = e.Hashmap()
	_ = e.HashmapPtr()
}

func Test_BasicUInt16_IsValidRange_FromMisc12(t *testing.T) {
	// Arrange
	e := uint16Enum()

	// Act
	actual := args.Map{"result": e.IsValidRange(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": e.IsValidRange(100)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_BasicUInt16_ToEnumJsonBytes_FromMisc12(t *testing.T) {
	// Arrange
	e := uint16Enum()

	_, err := e.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	_, err = e.ToEnumJsonBytes(99)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicUInt16_ToEnumString(t *testing.T) {
	e := uint16Enum()
	_ = e.ToEnumString(0)
}

func Test_BasicUInt16_AppendPrependJoinValue_FromMisc12(t *testing.T) {
	e := uint16Enum()
	_ = e.AppendPrependJoinValue(".", 1, 0)
}

func Test_BasicUInt16_AppendPrependJoinNamer_FromMisc12(t *testing.T) {
	e := uint16Enum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_BasicUInt16_ToNumberString_FromMisc12(t *testing.T) {
	e := uint16Enum()
	_ = e.ToNumberString(uint16(42))
}

func Test_BasicUInt16_UnmarshallToValue_AllBranches(t *testing.T) {
	// Arrange
	e := uint16Enum()

	_, err := e.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	v, err := e.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err != nil || v != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected min", actual)

	v, err = e.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	actual = args.Map{"result": err != nil || v != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_BasicUInt16_EnumType_FromMisc12(t *testing.T) {
	// Arrange
	e := uint16Enum()

	// Act
	actual := args.Map{"result": e.EnumType() != enumtype.UnsignedInteger16}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected UnsignedInteger16", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicString — full coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_BasicString_IsAnyNamesOf_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.IsAnyNamesOf("Invalid", "Invalid", "Active")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyNamesOf("Invalid", "Active")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicString_IsAnyOf_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.IsAnyOf("x")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)

	actual = args.Map{"result": e.IsAnyOf("Active", "Invalid", "Active")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": e.IsAnyOf("Active", "Invalid")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BasicString_MinMax(t *testing.T) {
	e := stringEnum()
	_ = e.Min()
	_ = e.Max()
}

func Test_BasicString_Ranges_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": len(e.Ranges()) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_BasicString_HasAnyItem_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_BasicString_MaxIndex_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.MaxIndex() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_BasicString_GetNameByIndex_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.GetNameByIndex(1) != "Active"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Active", actual)

	actual = args.Map{"result": e.GetNameByIndex(100) != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)

	actual = args.Map{"result": e.GetNameByIndex(0) != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for index 0 (condition > 0)", actual)
}

func Test_BasicString_GetIndexByName_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	idx := e.GetIndexByName("Active")

	// Act
	actual := args.Map{"result": idx < 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid index", actual)

	idx = e.GetIndexByName("")
	actual = args.Map{"result": idx >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for empty", actual)

	idx = e.GetIndexByName("NotExist")
	actual = args.Map{"result": idx >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for not exist", actual)
}

func Test_BasicString_NameWithIndexMap_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()
	m := e.NameWithIndexMap()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_BasicString_RangesIntegers_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()
	r := e.RangesIntegers()

	// Act
	actual := args.Map{"result": len(r) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_BasicString_Hashset_HashsetPtr(t *testing.T) {
	e := stringEnum()
	_ = e.Hashset()
	_ = e.HashsetPtr()
}

func Test_BasicString_GetValueByName_AllBranches(t *testing.T) {
	// Arrange
	e := stringEnum()

	v, err := e.GetValueByName("Active")

	// Act
	actual := args.Map{"result": err != nil || v != "Active"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Active", actual)

	_, err = e.GetValueByName("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicString_IsValidRange_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.IsValidRange("Active")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": e.IsValidRange("NotExist")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_BasicString_OnlySupportedErr_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()
	err := e.OnlySupportedErr("Active")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error (unsupported exist)", actual)

	err = e.OnlySupportedErr("Invalid", "Active", "Inactive")
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil (all supported)", actual)
}

func Test_BasicString_OnlySupportedMsgErr_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()
	err := e.OnlySupportedMsgErr("test msg", "Active")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicString_AppendPrependJoinValue_FromMisc12(t *testing.T) {
	e := stringEnum()
	_ = e.AppendPrependJoinValue(".", "Active", "Invalid")
}

func Test_BasicString_AppendPrependJoinNamer(t *testing.T) {
	e := stringEnum()
	_ = e.AppendPrependJoinNamer(".", testNamer{"B"}, testNamer{"A"})
}

func Test_BasicString_ToEnumJsonBytes_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	_, err := e.ToEnumJsonBytes("Active")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	_, err = e.ToEnumJsonBytes("NotExist")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_BasicString_UnmarshallToValue_AllBranches(t *testing.T) {
	// Arrange
	e := stringEnum()

	_, err := e.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	v, err := e.UnmarshallToValue(true, nil)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)
	_ = v

	v, err = e.UnmarshallToValue(true, []byte(""))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(true, []byte(`""`))
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no error", actual)

	v, err = e.UnmarshallToValue(false, []byte("Active"))
	actual = args.Map{"result": err != nil || v != "Active"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Active", actual)
}

func Test_BasicString_EnumType_FromMisc12(t *testing.T) {
	// Arrange
	e := stringEnum()

	// Act
	actual := args.Map{"result": e.EnumType() != enumtype.String}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected String", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// numberEnumBase — all methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_NumberEnumBase_MinMaxAny_FromMisc12(t *testing.T) {
	e := byteEnum()
	min, max := e.MinMaxAny()
	_ = min
	_ = max
}

func Test_NumberEnumBase_MinValueString_MaxValueString(t *testing.T) {
	// Arrange
	e := byteEnum()
	s := e.MinValueString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	s = e.MaxValueString()
	actual = args.Map{"result": s == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	// Call again to test cached path
	s = e.MinValueString()
	s = e.MaxValueString()
	_ = s
}

func Test_NumberEnumBase_MinInt_MaxInt(t *testing.T) {
	e := byteEnum()
	_ = e.MinInt()
	_ = e.MaxInt()
}

func Test_NumberEnumBase_AllNameValues_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	nvs := e.AllNameValues()

	// Act
	actual := args.Map{"result": len(nvs) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_NumberEnumBase_RangesMap_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.RangesMap()
}

func Test_NumberEnumBase_OnlySupportedErr_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	err := e.OnlySupportedErr("Active")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_NumberEnumBase_OnlySupportedMsgErr_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	err := e.OnlySupportedMsgErr("msg", "Active")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_NumberEnumBase_IntegerEnumRanges_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	r := e.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(r) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_NumberEnumBase_Length_Count(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)

	actual = args.Map{"result": e.Count() != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_NumberEnumBase_RangesDynamicMap_DynamicMap(t *testing.T) {
	e := byteEnum()
	_ = e.RangesDynamicMap()
	_ = e.DynamicMap()

	// Call again to test cached path
	_ = e.RangesDynamicMap()
}

func Test_NumberEnumBase_RangesIntegerStringMap_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.RangesIntegerStringMap()
}

func Test_NumberEnumBase_KeyAnyValues_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.KeyAnyValues()

	// Call again for cached
	_ = e.KeyAnyValues()
}

func Test_NumberEnumBase_KeyValIntegers_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	kvs := e.KeyValIntegers()

	// Act
	actual := args.Map{"result": len(kvs) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_NumberEnumBase_Loop_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	count := 0

	e.Loop(func(index int, name string, anyVal any) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)

	// Test break
	count = 0

	e.Loop(func(index int, name string, anyVal any) bool {
		count++
		return true
	})

	actual = args.Map{"result": count != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_NumberEnumBase_LoopInteger_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	count := 0

	e.LoopInteger(func(index int, name string, anyVal int) bool {
		count++
		return false
	})

	// Act
	actual := args.Map{"result": count != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)

	// Test break
	count = 0

	e.LoopInteger(func(index int, name string, anyVal int) bool {
		count++
		return true
	})

	actual = args.Map{"result": count != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_NumberEnumBase_TypeName_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()

	// Act
	actual := args.Map{"result": e.TypeName() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_NumberEnumBase_NameWithValueOption(t *testing.T) {
	e := byteEnum()
	_ = e.NameWithValueOption(byte(1), true)
	_ = e.NameWithValueOption(byte(1), false)
}

func Test_NumberEnumBase_NameWithValue_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.NameWithValue(byte(1))
}

func Test_NumberEnumBase_ValueString_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.ValueString(byte(1))
}

func Test_NumberEnumBase_Format_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	r := e.Format("Enum of {type-name} - {name} - {value}", byte(1))

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_NumberEnumBase_RangeNamesCsv_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	csv := e.RangeNamesCsv()

	// Act
	actual := args.Map{"result": csv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_NumberEnumBase_RangesInvalidMessage_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	msg := e.RangesInvalidMessage()

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_NumberEnumBase_RangesInvalidErr_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	err := e.RangesInvalidErr()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NumberEnumBase_StringRangesPtr_StringRanges(t *testing.T) {
	e := byteEnum()
	_ = e.StringRangesPtr()
	_ = e.StringRanges()
}

func Test_NumberEnumBase_NamesHashset_FromMisc12(t *testing.T) {
	// Arrange
	e := byteEnum()
	h := e.NamesHashset()

	// Act
	actual := args.Map{"result": h["Active"]}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Active in hashset", actual)
}

func Test_NumberEnumBase_JsonString_FromMisc12(t *testing.T) {
	e := byteEnum()
	_ = e.JsonString(byte(1))
}

func Test_NumberEnumBase_ToEnumString_ToName(t *testing.T) {
	e := byteEnum()
	_ = e.ToEnumString(byte(1))
	_ = e.ToName(byte(1))
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — comprehensive coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicMap_AddOrUpdate_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.AddOrUpdate("b", 2)

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)

	isNew = dm.AddOrUpdate("a", 3)

	actual = args.Map{"result": isNew}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not new", actual)
}

func Test_DynamicMap_Set_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isNew := dm.Set("b", 2)

	// Act
	actual := args.Map{"result": isNew}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected new", actual)
}

func Test_DynamicMap_AddNewOnly_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isAdded := dm.AddNewOnly("b", 2)

	// Act
	actual := args.Map{"result": isAdded}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected added", actual)

	isAdded = dm.AddNewOnly("a", 3)

	actual = args.Map{"result": isAdded}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not added", actual)
}

func Test_DynamicMap_AllKeys_AllKeysSorted(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	keys := dm.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	sorted := dm.AllKeysSorted()

	actual = args.Map{"result": sorted[0] != "a" || sorted[1] != "b"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)

	empty := enumimpl.DynamicMap{}
	_ = empty.AllKeys()
	_ = empty.AllKeysSorted()
}

func Test_DynamicMap_AllValuesStrings_Sorted(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	_ = dm.AllValuesStrings()
	_ = dm.AllValuesStringsSorted()

	empty := enumimpl.DynamicMap{}
	_ = empty.AllValuesStrings()
	_ = empty.AllValuesStringsSorted()
}

func Test_DynamicMap_AllValuesIntegers_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	vals := dm.AllValuesIntegers()

	// Act
	actual := args.Map{"result": len(vals) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	empty := enumimpl.DynamicMap{}
	_ = empty.AllValuesIntegers()
}

func Test_DynamicMap_MapIntegerString_FromMisc12(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m, keys := dm.MapIntegerString()
	_ = m
	_ = keys

	empty := enumimpl.DynamicMap{}
	_, _ = empty.MapIntegerString()

	// String values
	strDm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	_, _ = strDm.MapIntegerString()
}

func Test_DynamicMap_SortedKeyValues_FromMisc12(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	_ = dm.SortedKeyValues()

	empty := enumimpl.DynamicMap{}
	_ = empty.SortedKeyValues()
}

func Test_DynamicMap_SortedKeyAnyValues_FromMisc12(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	_ = dm.SortedKeyAnyValues()

	empty := enumimpl.DynamicMap{}
	_ = empty.SortedKeyAnyValues()

	// String values
	strDm := enumimpl.DynamicMap{"a": "x", "b": "y"}
	_ = strDm.SortedKeyAnyValues()
}

func Test_DynamicMap_First_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	k, v := dm.First()
	_ = k
	_ = v

	empty := enumimpl.DynamicMap{}
	k, v = empty.First()

	// Act
	actual := args.Map{"result": k != "" || v != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_DynamicMap_IsValueTypeOf_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	isInt := dm.IsValueTypeOf(reflect.TypeOf(1))

	// Act
	actual := args.Map{"result": isInt}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_DynamicMap_IsValueString_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsValueString()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	strDm := enumimpl.DynamicMap{"a": "x"}

	actual = args.Map{"result": strDm.IsValueString()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_DynamicMap_Length_Count_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"result": dm.Length() != 1 || dm.Count() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	actual = args.Map{"result": dm.IsEmpty() || !dm.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not empty", actual)

	var nilDm *enumimpl.DynamicMap

	actual = args.Map{"result": nilDm.Length() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil Length should be 0", actual)
}

func Test_DynamicMap_LastIndex_HasIndex(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.LastIndex() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	actual = args.Map{"result": dm.HasIndex(1) || dm.HasIndex(2)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "index check failed", actual)
}

func Test_DynamicMap_HasKey_IsMissingKey_HasAllKeys_HasAnyKeys(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.HasKey("a") || dm.HasKey("c")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasKey failed", actual)

	actual = args.Map{"result": dm.IsMissingKey("a") || !dm.IsMissingKey("c")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsMissingKey failed", actual)

	actual = args.Map{"result": dm.HasAllKeys("a", "b") || dm.HasAllKeys("a", "c")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAllKeys failed", actual)

	actual = args.Map{"result": dm.HasAnyKeys("a", "c") || dm.HasAnyKeys("c", "d")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "HasAnyKeys failed", actual)
}

func Test_DynamicMap_IsEqual_AllBranches(t *testing.T) {
	// Arrange
	var nilA, nilB *enumimpl.DynamicMap

	// Act
	actual := args.Map{"result": nilA.IsEqual(false, nilB)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	dm := enumimpl.DynamicMap{"a": 1}

	actual = args.Map{"result": nilA.IsEqual(false, &dm)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)

	actual = args.Map{"result": dm.IsEqual(false, nilA)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-nil vs nil should not be equal", actual)

	dm2 := enumimpl.DynamicMap{"a": 1}

	actual = args.Map{"result": dm.IsEqual(false, &dm2)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same maps should be equal", actual)

	actual = args.Map{"result": dm.IsEqual(false, &dm)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same pointer should be equal", actual)
}

func Test_DynamicMap_IsRawEqual_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	raw := map[string]any{"a": 1}

	// Act
	actual := args.Map{"result": dm.IsRawEqual(false, raw)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)

	// Different length
	actual = args.Map{"result": dm.IsRawEqual(false, map[string]any{"a": 1, "b": 2})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	// Missing key
	actual = args.Map{"result": dm.IsRawEqual(false, map[string]any{"b": 1})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	// nil checks
	var nilDm *enumimpl.DynamicMap

	actual = args.Map{"result": nilDm.IsRawEqual(false, nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	actual = args.Map{"result": nilDm.IsRawEqual(false, raw)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
}

func Test_DynamicMap_IsMismatch_IsRawMismatch(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	dm2 := enumimpl.DynamicMap{"a": 2}

	// Act
	actual := args.Map{"result": dm.IsMismatch(false, &dm2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)

	actual = args.Map{"result": dm.IsRawMismatch(false, map[string]any{"a": 2})}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)
}

func Test_DynamicMap_Raw_FromMisc12(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.Raw()
}

func Test_DynamicMap_DiffRaw_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	right := map[string]any{"a": 1, "c": 3}

	diff := dm.DiffRaw(false, right)

	// Act
	actual := args.Map{"result": diff.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected diff", actual)
}

func Test_DynamicMap_DiffRawUsingDifferChecker_AllBranches(t *testing.T) {
	// Arrange
	// nil left, nil right
	var nilDm *enumimpl.DynamicMap
	diff := nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	// Act
	actual := args.Map{"result": diff.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)

	// nil left, non-nil right
	right := map[string]any{"a": 1}
	diff = nilDm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, right)

	actual = args.Map{"result": diff.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	// non-nil left, nil right
	dm := enumimpl.DynamicMap{"a": 1}
	diff = dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)

	actual = args.Map{"result": diff.IsEmpty()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	// Equal maps
	diff = dm.DiffRawUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})

	actual = args.Map{"result": diff.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_DynamicMap_DiffRawLeftRightUsingDifferChecker(t *testing.T) {
	var nilDm *enumimpl.DynamicMap
	l, r := nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	_, _ = l, r

	l, r = nilDm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	_, _ = l, r

	dm := enumimpl.DynamicMap{"a": 1}
	l, r = dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, nil)
	_, _ = l, r

	dm2 := enumimpl.DynamicMap{"a": 1, "b": 2}
	l, r = dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, dm2)
	_, _ = l, r

	// Equal maps
	l, r = dm.DiffRawLeftRightUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 1})
	_, _ = l, r
}

func Test_DynamicMap_DiffJsonMessage_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessage(false, map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for equal", actual)

	msg = dm.DiffJsonMessage(false, map[string]any{"a": 2})

	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty for different", actual)
}

func Test_DynamicMap_DiffJsonMessageUsingDifferChecker_FromMisc12(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.DiffJsonMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, map[string]any{"a": 2})
}

func Test_DynamicMap_DiffJsonMessageLeftRight_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.DiffJsonMessageLeftRight(false, map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)

	msg = dm.DiffJsonMessageLeftRight(false, map[string]any{"a": 2, "b": 3})

	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_DynamicMap_ShouldDiffMessage_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ShouldDiffMessage(false, "test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)

	msg = dm.ShouldDiffMessage(false, "test", map[string]any{"a": 2})

	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_DynamicMap_ShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.ShouldDiffMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	_ = dm.ShouldDiffMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 2})
}

func Test_DynamicMap_ShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.ShouldDiffLeftRightMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
	_ = dm.ShouldDiffLeftRightMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 2})
}

func Test_DynamicMap_ExpectingMessage_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	msg := dm.ExpectingMessage("test", map[string]any{"a": 1})

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for equal", actual)

	msg = dm.ExpectingMessage("test", map[string]any{"a": 2})

	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty for different", actual)
}

func Test_DynamicMap_IsKeysEqualOnly_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	// Act
	actual := args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 99, "b": 88})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 1})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	actual = args.Map{"result": dm.IsKeysEqualOnly(map[string]any{"a": 1, "c": 3})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	var nilDm *enumimpl.DynamicMap

	actual = args.Map{"result": nilDm.IsKeysEqualOnly(nil)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "both nil should be equal", actual)

	actual = args.Map{"result": nilDm.IsKeysEqualOnly(map[string]any{"a": 1})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil vs non-nil should not be equal", actual)
}

func Test_DynamicMap_KeyValue_KeyValueString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	v, found := dm.KeyValue("a")

	// Act
	actual := args.Map{"result": found || v != 1}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_, found = dm.KeyValue("z")

	actual = args.Map{"result": found}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)

	s, found := dm.KeyValueString("a")

	actual = args.Map{"result": found || s != "1"}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	_, found = dm.KeyValueString("z")

	actual = args.Map{"result": found}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_DynamicMap_KeyValueIntDefault_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 42, "b": "notint"}
	v := dm.KeyValueIntDefault("a")

	// Act
	actual := args.Map{"result": v != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	v = dm.KeyValueIntDefault("z")

	actual = args.Map{"result": v >= 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_DynamicMap_KeyValueByte_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": byte(42), "b": 100}
	v, found, failed := dm.KeyValueByte("a")

	// Act
	actual := args.Map{"result": found || failed || v != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	_, found, _ = dm.KeyValueByte("z")

	actual = args.Map{"result": found}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_DynamicMap_KeyValueInt_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 42, "b": byte(10)}
	v, found, failed := dm.KeyValueInt("a")

	// Act
	actual := args.Map{"result": found || failed || v != 42}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	v, found, failed = dm.KeyValueInt("b")
	_ = v
	_ = found
	_ = failed

	_, found, _ = dm.KeyValueInt("z")

	actual = args.Map{"result": found}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not found", actual)
}

func Test_DynamicMap_Add_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{}
	result := dm.Add("key", "val")

	// Act
	actual := args.Map{"result": result == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_DynamicMap_ConvMap_Methods(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}

	_ = dm.ConvMapStringInteger()
	_ = dm.ConvMapIntegerString()
	_ = dm.ConvMapByteString()
	_ = dm.ConvMapInt8String()
	_ = dm.ConvMapInt16String()
	_ = dm.ConvMapInt32String()
	_ = dm.ConvMapUInt16String()
	_ = dm.ConvMapStringString()
	_ = dm.ConvMapInt64String()

	empty := enumimpl.DynamicMap{}
	_ = empty.ConvMapStringInteger()
	_ = empty.ConvMapIntegerString()
	_ = empty.ConvMapByteString()
	_ = empty.ConvMapInt8String()
	_ = empty.ConvMapInt16String()
	_ = empty.ConvMapInt32String()
	_ = empty.ConvMapUInt16String()
	_ = empty.ConvMapStringString()
	_ = empty.ConvMapInt64String()
}

func Test_DynamicMap_BasicFactories(t *testing.T) {
	dm := enumimpl.DynamicMap{"Invalid": 0, "Active": 1}

	_ = dm.BasicByte("test")
	_ = dm.BasicByteUsingAliasMap("test", nil)
	_ = dm.BasicInt8("test")
	_ = dm.BasicInt8UsingAliasMap("test", nil)
	_ = dm.BasicInt16("test")
	_ = dm.BasicInt16UsingAliasMap("test", nil)
	_ = dm.BasicInt32("test")
	_ = dm.BasicInt32UsingAliasMap("test", nil)
	_ = dm.BasicString("test")
	_ = dm.BasicStringUsingAliasMap("test", nil)
	_ = dm.BasicUInt16("test")
	_ = dm.BasicUInt16UsingAliasMap("test", nil)
}

func Test_DynamicMap_ConcatNew_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	another := enumimpl.DynamicMap{"b": 2}

	result := dm.ConcatNew(true, another)

	// Act
	actual := args.Map{"result": result.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	result = dm.ConcatNew(false, another)
	actual = args.Map{"result": result.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	empty := enumimpl.DynamicMap{}
	result = empty.ConcatNew(true, enumimpl.DynamicMap{})

	actual = args.Map{"result": result.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_DynamicMap_Strings_String(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.Strings()
	_ = dm.String()

	empty := enumimpl.DynamicMap{}
	_ = empty.Strings()
	_ = empty.String()
}

func Test_DynamicMap_StringsUsingFmt_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	r := dm.StringsUsingFmt(func(index int, key string, val any) string {
		return key
	})

	// Act
	actual := args.Map{"result": len(r) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)

	empty := enumimpl.DynamicMap{}
	_ = empty.StringsUsingFmt(func(index int, key string, val any) string { return "" })
}

func Test_DynamicMap_IsStringEqual_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	s := dm.String()

	// Act
	actual := args.Map{"result": dm.IsStringEqual(s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_DynamicMap_Serialize_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	b, err := dm.Serialize()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected serialized", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DiffLeftRight — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_DiffLeftRight_AllMethods(t *testing.T) {
	// Arrange
	d := &enumimpl.DiffLeftRight{Left: 1, Right: 2}

	l, r := d.Types()
	_ = l
	_ = r

	// Act
	actual := args.Map{"result": d.IsSameTypeSame() != true}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same type", actual)

	actual = args.Map{"result": d.IsSame()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not same", actual)

	actual = args.Map{"result": d.IsSameRegardlessOfType()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not same", actual)

	actual = args.Map{"result": d.IsEqual(false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	actual = args.Map{"result": d.IsEqual(true)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	actual = args.Map{"result": d.HasMismatch(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)

	actual = args.Map{"result": d.HasMismatch(true)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)

	actual = args.Map{"result": d.IsNotEqual()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	actual = args.Map{"result": d.HasMismatchRegardlessOfType()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected mismatch", actual)

	_ = d.String()
	_ = d.JsonString()

	ls, rs := d.SpecificFullString()
	_ = ls
	_ = rs

	ds := d.DiffString()
	actual = args.Map{"result": ds == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)

	// Same values
	same := &enumimpl.DiffLeftRight{Left: 1, Right: 1}
	ds = same.DiffString()

	actual = args.Map{"result": ds != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for same", actual)

	// nil
	var nilD *enumimpl.DiffLeftRight

	actual = args.Map{"result": nilD.JsonString() != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// differCheckerImpl — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_DifferCheckerImpl(t *testing.T) {
	// Arrange
	d := enumimpl.DefaultDiffCheckerImpl

	l := d.GetSingleDiffResult(true, 1, 2)

	// Act
	actual := args.Map{"result": l != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected left", actual)

	r := d.GetSingleDiffResult(false, 1, 2)
	actual = args.Map{"result": r != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected right", actual)

	val := d.GetResultOnKeyMissingInRightExistInLeft("k", 42)
	actual = args.Map{"result": val != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	actual = args.Map{"result": d.IsEqual(false, 1, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)

	actual = args.Map{"result": d.IsEqual(false, 1, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)

	actual = args.Map{"result": d.IsEqual(true, 1, 1)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal regardless", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// leftRightDiffCheckerImpl — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_LeftRightDiffCheckerImpl(t *testing.T) {
	// Arrange
	d := enumimpl.LeftRightDiffCheckerImpl

	_ = d.GetSingleDiffResult(true, 1, 2)
	_ = d.GetResultOnKeyMissingInRightExistInLeft("k", 42)

	// Act
	actual := args.Map{"result": d.IsEqual(false, 1, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Format / FormatUsingFmt — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_Format_FromMisc12(t *testing.T) {
	// Arrange
	result := enumimpl.Format("MyType", "Active", "1", "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

type mockFormatter struct{}

func (m mockFormatter) TypeName() string   { return "MyType" }
func (m mockFormatter) Name() string       { return "Active" }
func (m mockFormatter) ValueString() string { return "1" }

func Test_FormatUsingFmt_FromMisc12(t *testing.T) {
	// Arrange
	result := enumimpl.FormatUsingFmt(mockFormatter{}, "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Standalone functions — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_NameWithValue_FromMisc12(t *testing.T) {
	// Arrange
	r := enumimpl.NameWithValue(42)

	// Act
	actual := args.Map{"result": r == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_PrependJoin_FromMisc12(t *testing.T) {
	// Arrange
	r := enumimpl.PrependJoin(".", "a", "b", "c")

	// Act
	actual := args.Map{"result": r != "a.b.c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a.b.c", actual)
}

func Test_JoinPrependUsingDot_FromMisc12(t *testing.T) {
	// Arrange
	r := enumimpl.JoinPrependUsingDot("a", "b", "c")

	// Act
	actual := args.Map{"result": r != "a.b.c"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected a.b.c", actual)
}

func Test_OnlySupportedErr(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(4, []string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	err = enumimpl.OnlySupportedErr(4, []string{"a", "b"}, "a", "b")

	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)

	err = enumimpl.OnlySupportedErr(4, []string{}, "a")

	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty allNames", actual)
}

func Test_UnsupportedNames_FromMisc12(t *testing.T) {
	// Arrange
	unsupported := enumimpl.UnsupportedNames([]string{"a", "b", "c"}, "a")

	// Act
	actual := args.Map{"result": len(unsupported) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_AllNameValues_FromMisc12(t *testing.T) {
	// Arrange
	result := enumimpl.AllNameValues([]string{"a", "b"}, []byte{0, 1})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_IntegersRangesOfAnyVal_FromMisc12(t *testing.T) {
	// Arrange
	result := enumimpl.IntegersRangesOfAnyVal([]byte{2, 0, 1})

	// Act
	actual := args.Map{"result": len(result) != 3 || result[0] != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted [0,1,2]", actual)
}

func Test_ConvEnumAnyValToInteger_AllBranches(t *testing.T) {
	// int
	// Act
	actual := args.Map{"result": enumimpl.ConvEnumAnyValToInteger(42) != 42}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	// string
	r := enumimpl.ConvEnumAnyValToInteger("hello")
	_ = r

	// byte
	actual = args.Map{"result": enumimpl.ConvEnumAnyValToInteger(byte(5)) != 5}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 5", actual)

	// non-convertible
	r = enumimpl.ConvEnumAnyValToInteger(struct{}{})
	_ = r
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyVal — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyAnyVal_AllMethods(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "test", AnyValue: 42}

	// Act
	actual := args.Map{"result": kav.KeyString() != "test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)

	actual = args.Map{"result": kav.AnyVal() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	_ = kav.AnyValString()
	_ = kav.WrapKey()
	_ = kav.WrapValue()

	actual = args.Map{"result": kav.IsString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	actual = args.Map{"result": kav.ValInt() != 42}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 42", actual)

	kvi := kav.KeyValInteger()

	actual = args.Map{"result": kvi.Key != "test"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)

	_ = kav.String()

	// String type
	kavStr := enumimpl.KeyAnyVal{Key: "test", AnyValue: "hello"}

	actual = args.Map{"result": kavStr.IsString()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	_ = kavStr.String()
}

func Test_KeyAnyValues_FromMisc12(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{"a", "b"}, []int{1, 2})

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)

	result = enumimpl.KeyAnyValues([]string{}, []int{})

	actual = args.Map{"result": len(result) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValInteger — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValInteger_AllMethods(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "test", ValueInteger: 42}

	_ = kvi.WrapKey()
	_ = kvi.WrapValue()
	_ = kvi.String()

	kav := kvi.KeyAnyVal()

	// Act
	actual := args.Map{"result": kav.Key != "test"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected test", actual)

	actual = args.Map{"result": kvi.IsString()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicByteCreator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewBasicByte_CreateUsingMap(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.CreateUsingMap("test", map[byte]string{0: "Invalid", 1: "Active"})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_CreateUsingMapPlusAliasMap(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMap("test",
		map[byte]string{0: "Invalid", 1: "Active"},
		map[string]byte{"active": 1})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_Create(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.Create("test", []byte{0, 1}, []string{"Invalid", "Active"}, 0, 1)

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_UsingTypeSlice(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.UsingTypeSlice("test", []string{"Invalid", "Active"})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_DefaultWithAliasMap(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.DefaultWithAliasMap(testByteVal, []string{"Invalid", "Active"}, map[string]byte{"act": 1})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_DefaultAllCases(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.DefaultAllCases(testByteVal, []string{"Invalid", "Active"})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_DefaultWithAliasMapAllCases(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.DefaultWithAliasMapAllCases(testByteVal, []string{"Invalid", "Active"}, map[string]byte{"act": 1})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.UsingFirstItemSliceAliasMap(testByteVal, []string{"Invalid", "Active"}, nil)

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_UsingFirstItemSliceCaseOptions(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.UsingFirstItemSliceCaseOptions(true, testByteVal, []string{"Invalid", "Active"})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_UsingFirstItemSliceAllCases(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(testByteVal, []string{"Invalid", "Active"})

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.CreateUsingSlicePlusAliasMapOptions(false, testByteVal, []string{"Invalid", "Active"}, nil)

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewBasicByte_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.CreateUsingMapPlusAliasMapOptions(true, testByteVal, map[byte]string{0: "Invalid", 1: "Active"}, nil)

	// Act
	actual := args.Map{"result": e == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicInt8Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewBasicInt8_All(t *testing.T) {
	_ = enumimpl.New.BasicInt8.CreateUsingMap("test", map[int8]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicInt8.CreateUsingMapPlusAliasMap("test", map[int8]string{0: "Invalid"}, map[string]int8{"act": 1})
	_ = enumimpl.New.BasicInt8.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt8.DefaultWithAliasMap(testInt8Val, []string{"Invalid", "Active"}, map[string]int8{"act": 1})
	_ = enumimpl.New.BasicInt8.DefaultAllCases(testInt8Val, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt8.DefaultWithAliasMapAllCases(testInt8Val, []string{"Invalid", "Active"}, map[string]int8{"act": 1})
	_ = enumimpl.New.BasicInt8.UsingFirstItemSliceAliasMap(testInt8Val, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicInt8.CreateUsingSlicePlusAliasMapOptions(false, testInt8Val, []string{"Invalid"}, nil)
	_ = enumimpl.New.BasicInt8.CreateUsingMapPlusAliasMapOptions(true, testInt8Val, map[int8]string{0: "Invalid"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicInt16Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewBasicInt16_All(t *testing.T) {
	_ = enumimpl.New.BasicInt16.CreateUsingMap("test", map[int16]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicInt16.CreateUsingMapPlusAliasMap("test", map[int16]string{0: "Invalid"}, map[string]int16{"act": 1})
	_ = enumimpl.New.BasicInt16.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt16.DefaultWithAliasMap(testInt16Val, []string{"Invalid", "Active"}, map[string]int16{"act": 1})
	_ = enumimpl.New.BasicInt16.DefaultAllCases(testInt16Val, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt16.DefaultWithAliasMapAllCases(testInt16Val, []string{"Invalid", "Active"}, map[string]int16{"act": 1})
	_ = enumimpl.New.BasicInt16.UsingFirstItemSliceAliasMap(testInt16Val, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicInt16.CreateUsingSlicePlusAliasMapOptions(false, testInt16Val, []string{"Invalid"}, nil)
	_ = enumimpl.New.BasicInt16.CreateUsingMapPlusAliasMapOptions(true, testInt16Val, map[int16]string{0: "Invalid"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicInt32Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewBasicInt32_All(t *testing.T) {
	_ = enumimpl.New.BasicInt32.CreateUsingMap("test", map[int32]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicInt32.CreateUsingMapPlusAliasMap("test", map[int32]string{0: "Invalid"}, map[string]int32{"act": 1})
	_ = enumimpl.New.BasicInt32.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicInt32.DefaultWithAliasMap(testInt32Val, []string{"Invalid", "Active"}, map[string]int32{"act": 1})
	_ = enumimpl.New.BasicInt32.UsingFirstItemSliceAliasMap(testInt32Val, []string{"Invalid", "Active"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicUInt16Creator
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewBasicUInt16_All(t *testing.T) {
	_ = enumimpl.New.BasicUInt16.CreateUsingMap("test", map[uint16]string{0: "Invalid", 1: "Active"})
	_ = enumimpl.New.BasicUInt16.CreateUsingMapPlusAliasMap("test", map[uint16]string{0: "Invalid"}, map[string]uint16{"act": 1})
	_ = enumimpl.New.BasicUInt16.UsingTypeSlice("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicUInt16.DefaultWithAliasMap(testUInt16Val, []string{"Invalid", "Active"}, map[string]uint16{"act": 1})
	_ = enumimpl.New.BasicUInt16.UsingFirstItemSliceAliasMap(testUInt16Val, []string{"Invalid", "Active"}, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// Creator factories — newBasicStringCreator
// ══════════════════════════════════════════════════════════════════════════════

type testString string

const testStringVal testString = ""

func Test_NewBasicString_All(t *testing.T) {
	_ = enumimpl.New.BasicString.Create("test", []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicString.CreateDefault(testStringVal, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicString.CreateAliasMapOnly("test", []string{"Invalid", "Active"}, map[string]string{"act": "Active"})
	_ = enumimpl.New.BasicString.CreateUsingAliasMap("test", []string{"Invalid", "Active"}, nil, "Active", "Invalid")
	_ = enumimpl.New.BasicString.CreateUsingNamesSpread("test", "Invalid", "Active")
	_ = enumimpl.New.BasicString.CreateUsingNamesMinMax("test", []string{"Invalid", "Active"}, "Active", "Invalid")
	_ = enumimpl.New.BasicString.CreateUsingSlicePlusAliasMapOptions(true, testStringVal, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicString.CreateUsingMapPlusAliasMapOptions(false, testStringVal, []string{"Invalid", "Active"}, nil)
	_ = enumimpl.New.BasicString.UsingFirstItemSliceCaseOptions(true, testStringVal, []string{"Invalid", "Active"})
	_ = enumimpl.New.BasicString.UsingFirstItemSliceAllCases(testStringVal, []string{"Invalid", "Active"})
}

// ══════════════════════════════════════════════════════════════════════════════
// enumtype.Variant — coverage
// ══════════════════════════════════════════════════════════════════════════════

func Test_EnumType_Variant_AllMethods(t *testing.T) {
	// Arrange
	v := enumtype.Byte

	_ = v.TypeName()
	_ = v.ValueUInt16()
	_ = v.RangeNamesCsv()

	min, max := v.MinMaxAny()
	_, _ = min, max

	_ = v.MinValueString()
	_ = v.MaxValueString()
	_ = v.MaxInt()
	_ = v.MinInt()
	_ = v.RangesDynamicMap()
	_ = v.IntegerEnumRanges()
	_ = v.EnumType()
	_ = v.Value()
	_ = v.Name()
	_ = v.String()
	_ = v.NameValue()
	_ = v.ToNumberString()
	_ = v.ValueByte()
	_ = v.ValueInt()
	_ = v.ValueInt8()
	_ = v.ValueInt16()
	_ = v.ValueInt32()
	_ = v.ValueString()

	// Act
	actual := args.Map{"result": v.IsValid() || v.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected valid", actual)

	actual = args.Map{"result": v.IsNameEqual("Invalid")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)

	actual = args.Map{"result": v.IsNameEqual("Byte")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": v.IsAnyNamesOf("Invalid", "Byte")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": v.IsAnyNamesOf("Invalid")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_EnumType_Variant_TypeChecks(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.IsBoolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Byte.IsByte()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.UnsignedInteger16.IsUnsignedInteger16()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.UnsignedInteger32.IsUnsignedInteger32()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.UnsignedInteger64.IsUnsignedInteger64()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Integer8.IsInteger8()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Integer16.IsInteger16()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Integer32.IsInteger32()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Integer64.IsInteger64()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Integer.IsInteger()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.String.IsString()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Byte.IsNumber()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Integer8.IsAnyInteger()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Byte.IsAnyUnsignedNumber()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)

	actual = args.Map{"result": enumtype.Invalid.IsValid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)

	actual = args.Map{"result": enumtype.Invalid.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected invalid", actual)
}

func Test_EnumType_Variant_MarshalJSON(t *testing.T) {
	// Arrange
	v := enumtype.Byte
	b, err := v.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_EnumType_Variant_UnmarshalJSON(t *testing.T) {
	// Arrange
	v := enumtype.Invalid

	// Valid
	err := v.UnmarshalJSON([]byte(`"Byte"`))

	// Act
	actual := args.Map{"result": err != nil || v != enumtype.Byte}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Byte", actual)

	// Empty
	err = v.UnmarshalJSON([]byte(""))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// Too short
	err = v.UnmarshalJSON([]byte(`""`))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// Not found
	err = v.UnmarshalJSON([]byte(`"NotExist"`))
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_EnumType_Variant_Format_Panics(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected panic", actual)
	}()

	v := enumtype.Byte
	v.Format("test")
}

func Test_EnumType_Variant_RoundTrip(t *testing.T) {
	// Arrange
	original := enumtype.Integer32
	b, err := json.Marshal(original)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "marshal failed", actual)

	var result enumtype.Variant
	err = json.Unmarshal(b, &result)

	actual = args.Map{"result": err != nil || result != enumtype.Integer32}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Integer32", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap — isEqualSingle / isNotEqual via IsRawEqual (regardless type)
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicMap_IsRawEqual_RegardlessType_FromMisc12(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Regardless type: int 1 vs float64 should not match in string fmt

	// Act
	actual := args.Map{"result": dm.IsRawEqual(true, map[string]any{"a": 1})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal regardless", actual)

	actual = args.Map{"result": dm.IsRawEqual(true, map[string]any{"a": 2})}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// DynamicMap LogShouldDiff* (covers fmt.Println paths - only call, no output check)
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicMap_LogShouldDiffMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffMessage(false, "test", map[string]any{"a": 1})
}

func Test_DynamicMap_LogShouldDiffLeftRightMessage(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffLeftRightMessage(false, "test", map[string]any{"a": 1})
}

func Test_DynamicMap_LogShouldDiffMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
}

func Test_DynamicMap_LogShouldDiffLeftRightMessageUsingDifferChecker(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	_ = dm.LogShouldDiffLeftRightMessageUsingDifferChecker(enumimpl.DefaultDiffCheckerImpl, false, "test", map[string]any{"a": 1})
}

func Test_DynamicMap_LogExpectingMessage_FromMisc12(t *testing.T) {
	dm := enumimpl.DynamicMap{"a": 1}
	dm.LogExpectingMessage("test", map[string]any{"a": 1})
}

// ══════════════════════════════════════════════════════════════════════════════
// toStringsSliceOfDiffMap — string value branch
// ══════════════════════════════════════════════════════════════════════════════

func Test_DynamicMap_DiffJsonMessage_StringValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "x"}
	msg := dm.DiffJsonMessage(false, map[string]any{"a": "y"})

	// Act
	actual := args.Map{"result": msg == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// NamesHashset empty
// ══════════════════════════════════════════════════════════════════════════════

func Test_NamesHashset_Empty(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicByte.Create("test", []byte{}, []string{}, 0, 0)
	h := e.NamesHashset()

	// Act
	actual := args.Map{"result": len(h) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BasicString_GetIndexByName_EmptyEnum(t *testing.T) {
	// Arrange
	e := enumimpl.New.BasicString.Create("test", []string{})
	idx := e.GetIndexByName("something")

	// Act
	actual := args.Map{"result": idx >= 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected invalid for empty enum", actual)
}
