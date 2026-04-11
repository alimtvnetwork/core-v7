package enumimpltests

import (
	"testing"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── BasicInt8 — uncovered branches ──

func Test_BasicInt8_IsAnyOf_Empty(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B", "C"})

	// Act
	actual := args.Map{"anyOfEmpty": bi.IsAnyOf(0)}

	// Assert
	expected := args.Map{"anyOfEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- empty checkingItems", actual)
}

func Test_BasicInt8_IsAnyOf_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyOf(0, 5, 6)}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns false -- value not in list", actual)
}

func Test_BasicInt8_IsAnyOf_Found(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyOf(0, 0, 1)}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- value in list", actual)
}

func Test_BasicInt8_IsAnyNamesOf_Found(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyNamesOf(0, "A", "B")}

	// Assert
	expected := args.Map{"found": true}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns true -- name in list", actual)
}

func Test_BasicInt8_IsAnyNamesOf_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyNamesOf(0, "X", "Y")}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns false -- name not in list", actual)
}

func Test_BasicInt8_GetValueByName_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	_, err := bi.GetValueByName("UNKNOWN")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns error -- unknown name", actual)
}

func Test_BasicInt8_GetValueByName_Found(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	val, err := bi.GetValueByName("A")

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int8(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns value -- known name", actual)
}

func Test_BasicInt8_ToEnumJsonBytes_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	_, err := bi.ToEnumJsonBytes(99)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns error -- value not in map", actual)
}

func Test_BasicInt8_ToEnumJsonBytes_Found(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	b, err := bi.ToEnumJsonBytes(0)

	// Act
	actual := args.Map{
		"hasBytes": len(b) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns bytes -- valid value", actual)
}

func Test_BasicInt8_ExpectingEnumValueError_Mismatch(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	err := bi.ExpectingEnumValueError("B", int8(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- mismatch", actual)
}

func Test_BasicInt8_ExpectingEnumValueError_Unknown(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	err := bi.ExpectingEnumValueError("UNKNOWN", int8(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- unknown input", actual)
}

func Test_BasicInt8_UnmarshallToValue_NilNotMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	_, err := bi.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns error -- nil not mapped", actual)
}

func Test_BasicInt8_UnmarshallToValue_NilMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	val, err := bi.UnmarshallToValue(true, nil)

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int8(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped to first", actual)
}

func Test_BasicInt8_UnmarshallToValue_EmptyMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	val, err := bi.UnmarshallToValue(true, []byte(`""`))

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int8(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty string mapped", actual)
}

func Test_BasicInt8_MinMax(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B", "C"})

	// Act
	actual := args.Map{
		"min": bi.Min(),
		"max": bi.Max(),
	}

	// Assert
	expected := args.Map{
		"min": int8(0),
		"max": int8(2),
	}
	expected.ShouldBeEqual(t, 0, "MinMax returns correct -- three items", actual)
}

func Test_BasicInt8_Ranges_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"len": len(bi.Ranges())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Ranges returns correct length -- two items", actual)
}

func Test_BasicInt8_IsValidRange(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{
		"valid":   bi.IsValidRange(0),
		"invalid": bi.IsValidRange(5),
	}

	// Assert
	expected := args.Map{
		"valid":   true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidRange returns correct -- valid and invalid", actual)
}

func Test_BasicInt8_EnumType(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A"})

	// Act
	actual := args.Map{"notZero": bi.EnumType() != 0}

	// Assert
	expected := args.Map{"notZero": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-zero -- Integer8", actual)
}

func Test_BasicInt8_AppendPrependJoinValue(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})
	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinValue returns non-empty -- valid values", actual)
}

func Test_BasicInt8_GetStringValue_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"val": bi.GetStringValue(0)}

	// Assert
	expected := args.Map{"val": "A"}
	expected.ShouldBeEqual(t, 0, "GetStringValue returns name -- index 0", actual)
}

func Test_BasicInt8_Hashmap(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"hasItems": len(bi.Hashmap()) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Hashmap returns non-empty -- two items", actual)
}

func Test_BasicInt8_HashmapPtr(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt8.UsingTypeSlice("TestI8", []string{"A", "B"})

	// Act
	actual := args.Map{"notNil": bi.HashmapPtr() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "HashmapPtr returns non-nil -- two items", actual)
}

// ── BasicInt16 — uncovered branches ──

func Test_BasicInt16_IsAnyOf_Empty(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})

	// Act
	actual := args.Map{"anyOfEmpty": bi.IsAnyOf(0)}

	// Assert
	expected := args.Map{"anyOfEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- empty checkingItems", actual)
}

func Test_BasicInt16_IsAnyOf_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyOf(0, 5, 6)}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns false -- value not in list", actual)
}

func Test_BasicInt16_IsAnyNamesOf_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyNamesOf(0, "X")}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns false -- name not in list", actual)
}

func Test_BasicInt16_GetValueByName_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	_, err := bi.GetValueByName("UNKNOWN")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns error -- unknown name", actual)
}

func Test_BasicInt16_ToEnumJsonBytes_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	_, err := bi.ToEnumJsonBytes(99)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns error -- value not in map", actual)
}

func Test_BasicInt16_ExpectingEnumValueError_Mismatch(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	err := bi.ExpectingEnumValueError("B", int16(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- mismatch", actual)
}

func Test_BasicInt16_ExpectingEnumValueError_Unknown(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	err := bi.ExpectingEnumValueError("UNKNOWN", int16(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- unknown input", actual)
}

func Test_BasicInt16_UnmarshallToValue_NilNotMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	_, err := bi.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns error -- nil not mapped", actual)
}

func Test_BasicInt16_UnmarshallToValue_NilMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	val, err := bi.UnmarshallToValue(true, nil)

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int16(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped to first", actual)
}

func Test_BasicInt16_UnmarshallToValue_EmptyMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	val, err := bi.UnmarshallToValue(true, []byte(`""`))

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int16(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty string mapped", actual)
}

func Test_BasicInt16_MinMax(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})

	// Act
	actual := args.Map{
		"min": bi.Min(),
		"max": bi.Max(),
	}

	// Assert
	expected := args.Map{
		"min": int16(0),
		"max": int16(1),
	}
	expected.ShouldBeEqual(t, 0, "MinMax returns correct -- two items", actual)
}

func Test_BasicInt16_Ranges_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})

	// Act
	actual := args.Map{"len": len(bi.Ranges())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Ranges returns correct length -- two items", actual)
}

func Test_BasicInt16_IsValidRange(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})

	// Act
	actual := args.Map{
		"valid": bi.IsValidRange(0),
		"invalid": bi.IsValidRange(5),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidRange returns correct -- valid and invalid", actual)
}

func Test_BasicInt16_EnumType(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A"})

	// Act
	actual := args.Map{"notZero": bi.EnumType() != 0}

	// Assert
	expected := args.Map{"notZero": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-zero -- Integer16", actual)
}

func Test_BasicInt16_AppendPrependJoinValue_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt16.UsingTypeSlice("TestI16", []string{"A", "B"})
	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinValue returns non-empty -- valid values", actual)
}

// ── BasicInt32 — uncovered branches ──

func Test_BasicInt32_IsAnyOf_Empty(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})

	// Act
	actual := args.Map{"anyOfEmpty": bi.IsAnyOf(0)}

	// Assert
	expected := args.Map{"anyOfEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- empty checkingItems", actual)
}

func Test_BasicInt32_IsAnyOf_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyOf(0, 5, 6)}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns false -- value not in list", actual)
}

func Test_BasicInt32_IsAnyNamesOf_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bi.IsAnyNamesOf(0, "X")}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns false -- name not in list", actual)
}

func Test_BasicInt32_GetValueByName_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	_, err := bi.GetValueByName("UNKNOWN")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns error -- unknown name", actual)
}

func Test_BasicInt32_ToEnumJsonBytes_NotFound(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	_, err := bi.ToEnumJsonBytes(99)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns error -- value not in map", actual)
}

func Test_BasicInt32_ExpectingEnumValueError_Mismatch(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	err := bi.ExpectingEnumValueError("B", int32(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- mismatch", actual)
}

func Test_BasicInt32_ExpectingEnumValueError_Unknown(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	err := bi.ExpectingEnumValueError("UNKNOWN", int32(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- unknown input", actual)
}

func Test_BasicInt32_UnmarshallToValue_NilNotMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	_, err := bi.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns error -- nil not mapped", actual)
}

func Test_BasicInt32_UnmarshallToValue_NilMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	val, err := bi.UnmarshallToValue(true, nil)

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int32(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped to first", actual)
}

func Test_BasicInt32_UnmarshallToValue_EmptyMapped(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	val, err := bi.UnmarshallToValue(true, []byte(`""`))

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": int32(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty string mapped", actual)
}

func Test_BasicInt32_MinMax(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})

	// Act
	actual := args.Map{
		"min": bi.Min(),
		"max": bi.Max(),
	}

	// Assert
	expected := args.Map{
		"min": int32(0),
		"max": int32(1),
	}
	expected.ShouldBeEqual(t, 0, "MinMax returns correct -- two items", actual)
}

func Test_BasicInt32_Ranges_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})

	// Act
	actual := args.Map{"len": len(bi.Ranges())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Ranges returns correct length -- two items", actual)
}

func Test_BasicInt32_IsValidRange(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})

	// Act
	actual := args.Map{
		"valid": bi.IsValidRange(0),
		"invalid": bi.IsValidRange(5),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidRange returns correct -- valid and invalid", actual)
}

func Test_BasicInt32_EnumType(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A"})

	// Act
	actual := args.Map{"notZero": bi.EnumType() != 0}

	// Assert
	expected := args.Map{"notZero": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-zero -- Integer32", actual)
}

func Test_BasicInt32_AppendPrependJoinValue_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bi := enumimpl.New.BasicInt32.UsingTypeSlice("TestI32", []string{"A", "B"})
	result := bi.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinValue returns non-empty -- valid values", actual)
}

// ── BasicUInt16 — uncovered branches ──

func Test_BasicUInt16_IsAnyOf_Empty(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})

	// Act
	actual := args.Map{"anyOfEmpty": bu.IsAnyOf(0)}

	// Assert
	expected := args.Map{"anyOfEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns true -- empty checkingItems", actual)
}

func Test_BasicUInt16_IsAnyOf_NotFound(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bu.IsAnyOf(0, 5, 6)}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyOf returns false -- value not in list", actual)
}

func Test_BasicUInt16_IsAnyNamesOf_NotFound(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})

	// Act
	actual := args.Map{"found": bu.IsAnyNamesOf(0, "X")}

	// Assert
	expected := args.Map{"found": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf returns false -- name not in list", actual)
}

func Test_BasicUInt16_GetValueByName_NotFound(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	_, err := bu.GetValueByName("UNKNOWN")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetValueByName returns error -- unknown name", actual)
}

func Test_BasicUInt16_ToEnumJsonBytes_NotFound(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	_, err := bu.ToEnumJsonBytes(99)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ToEnumJsonBytes returns error -- value not in map", actual)
}

func Test_BasicUInt16_ExpectingEnumValueError_Mismatch(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	err := bu.ExpectingEnumValueError("B", uint16(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- mismatch", actual)
}

func Test_BasicUInt16_ExpectingEnumValueError_Unknown(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	err := bu.ExpectingEnumValueError("UNKNOWN", uint16(0))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "ExpectingEnumValueError returns error -- unknown input", actual)
}

func Test_BasicUInt16_UnmarshallToValue_NilNotMapped(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	_, err := bu.UnmarshallToValue(false, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns error -- nil not mapped", actual)
}

func Test_BasicUInt16_UnmarshallToValue_NilMapped(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	val, err := bu.UnmarshallToValue(true, nil)

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": uint16(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- nil mapped to first", actual)
}

func Test_BasicUInt16_UnmarshallToValue_EmptyMapped(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	val, err := bu.UnmarshallToValue(true, []byte(`""`))

	// Act
	actual := args.Map{
		"val": val,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": uint16(0),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "UnmarshallToValue returns min -- empty string mapped", actual)
}

func Test_BasicUInt16_MinMax(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})

	// Act
	actual := args.Map{
		"min": bu.Min(),
		"max": bu.Max(),
	}

	// Assert
	expected := args.Map{
		"min": uint16(0),
		"max": uint16(1),
	}
	expected.ShouldBeEqual(t, 0, "MinMax returns correct -- two items", actual)
}

func Test_BasicUInt16_Ranges_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})

	// Act
	actual := args.Map{"len": len(bu.Ranges())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Ranges returns correct length -- two items", actual)
}

func Test_BasicUInt16_IsValidRange(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})

	// Act
	actual := args.Map{
		"valid": bu.IsValidRange(0),
		"invalid": bu.IsValidRange(5),
	}

	// Assert
	expected := args.Map{
		"valid": true,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValidRange returns correct -- valid and invalid", actual)
}

func Test_BasicUInt16_EnumType(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A"})

	// Act
	actual := args.Map{"notZero": bu.EnumType() != 0}

	// Assert
	expected := args.Map{"notZero": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-zero -- UnsignedInteger16", actual)
}

func Test_BasicUInt16_AppendPrependJoinValue_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bu := enumimpl.New.BasicUInt16.UsingTypeSlice("TestU16", []string{"A", "B"})
	result := bu.AppendPrependJoinValue(".", 1, 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AppendPrependJoinValue returns non-empty -- valid values", actual)
}

// ── Creator methods — Int8 / Int16 / Int32 / UInt16 ──

func Test_BasicInt8Creator_CreateUsingMap(t *testing.T) {
	// Arrange
	m := map[int8]string{0: "Off", 1: "On"}
	bi := enumimpl.New.BasicInt8.CreateUsingMap("TestI8", m)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMap returns enum -- two entries", actual)
}

func Test_BasicInt8Creator_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	// Arrange
	type testEnum int8
	m := map[int8]string{0: "Off", 1: "On"}
	aliases := map[string]int8{"off": 0}
	bi := enumimpl.New.BasicInt8.CreateUsingMapPlusAliasMapOptions(true, testEnum(0), m, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMapPlusAliasMapOptions returns enum -- with aliases", actual)
}

func Test_BasicInt8Creator_DefaultAllCases(t *testing.T) {
	// Arrange
	type testEnum int8
	bi := enumimpl.New.BasicInt8.DefaultAllCases(testEnum(0), []string{"Off", "On"})

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAllCases returns enum -- two entries all cases", actual)
}

func Test_BasicInt8Creator_DefaultWithAliasMapAllCases(t *testing.T) {
	// Arrange
	type testEnum int8
	aliases := map[string]int8{"off": 0}
	bi := enumimpl.New.BasicInt8.DefaultWithAliasMapAllCases(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMapAllCases returns enum -- with alias all cases", actual)
}

func Test_BasicInt8Creator_Default(t *testing.T) {
	// Arrange
	type testEnum int8
	bi := enumimpl.New.BasicInt8.Default(testEnum(0), []string{"Off", "On"})

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Default returns enum -- two entries", actual)
}

func Test_BasicInt8Creator_DefaultWithAliasMap(t *testing.T) {
	// Arrange
	type testEnum int8
	aliases := map[string]int8{"off": 0}
	bi := enumimpl.New.BasicInt8.DefaultWithAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMap returns enum -- with alias", actual)
}

func Test_BasicInt8Creator_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type testEnum int8
	aliases := map[string]int8{"off": 0}
	bi := enumimpl.New.BasicInt8.UsingFirstItemSliceAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UsingFirstItemSliceAliasMap returns enum -- with alias", actual)
}

func Test_BasicInt8Creator_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	// Arrange
	type testEnum int8
	aliases := map[string]int8{"off": 0}
	bi := enumimpl.New.BasicInt8.CreateUsingSlicePlusAliasMapOptions(true, testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingSlicePlusAliasMapOptions returns enum -- with aliases", actual)
}

func Test_BasicInt16Creator_CreateUsingMap(t *testing.T) {
	// Arrange
	m := map[int16]string{0: "Off", 1: "On"}
	bi := enumimpl.New.BasicInt16.CreateUsingMap("TestI16", m)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMap returns enum -- two entries", actual)
}

func Test_BasicInt16Creator_CreateUsingMapPlusAliasMapOptions(t *testing.T) {
	// Arrange
	type testEnum int16
	m := map[int16]string{0: "Off", 1: "On"}
	aliases := map[string]int16{"off": 0}
	bi := enumimpl.New.BasicInt16.CreateUsingMapPlusAliasMapOptions(true, testEnum(0), m, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMapPlusAliasMapOptions returns enum -- with aliases", actual)
}

func Test_BasicInt16Creator_DefaultAllCases(t *testing.T) {
	// Arrange
	type testEnum int16
	bi := enumimpl.New.BasicInt16.DefaultAllCases(testEnum(0), []string{"Off", "On"})

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultAllCases returns enum -- two entries all cases", actual)
}

func Test_BasicInt16Creator_DefaultWithAliasMapAllCases(t *testing.T) {
	// Arrange
	type testEnum int16
	aliases := map[string]int16{"off": 0}
	bi := enumimpl.New.BasicInt16.DefaultWithAliasMapAllCases(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMapAllCases returns enum -- with alias all cases", actual)
}

func Test_BasicInt16Creator_Default(t *testing.T) {
	// Arrange
	type testEnum int16
	bi := enumimpl.New.BasicInt16.Default(testEnum(0), []string{"Off", "On"})

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Default returns enum -- two entries", actual)
}

func Test_BasicInt16Creator_DefaultWithAliasMap(t *testing.T) {
	// Arrange
	type testEnum int16
	aliases := map[string]int16{"off": 0}
	bi := enumimpl.New.BasicInt16.DefaultWithAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMap returns enum -- with alias", actual)
}

func Test_BasicInt16Creator_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type testEnum int16
	aliases := map[string]int16{"off": 0}
	bi := enumimpl.New.BasicInt16.UsingFirstItemSliceAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UsingFirstItemSliceAliasMap returns enum -- with alias", actual)
}

func Test_BasicInt16Creator_CreateUsingSlicePlusAliasMapOptions(t *testing.T) {
	// Arrange
	type testEnum int16
	aliases := map[string]int16{"off": 0}
	bi := enumimpl.New.BasicInt16.CreateUsingSlicePlusAliasMapOptions(true, testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingSlicePlusAliasMapOptions returns enum -- with aliases", actual)
}

func Test_BasicInt32Creator_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type testEnum int32
	aliases := map[string]int32{"off": 0}
	bi := enumimpl.New.BasicInt32.UsingFirstItemSliceAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UsingFirstItemSliceAliasMap returns enum -- with alias", actual)
}

func Test_BasicInt32Creator_Default(t *testing.T) {
	// Arrange
	type testEnum int32
	bi := enumimpl.New.BasicInt32.Default(testEnum(0), []string{"Off", "On"})

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Default returns enum -- two entries", actual)
}

func Test_BasicInt32Creator_DefaultWithAliasMap(t *testing.T) {
	// Arrange
	type testEnum int32
	aliases := map[string]int32{"off": 0}
	bi := enumimpl.New.BasicInt32.DefaultWithAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bi.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMap returns enum -- with alias", actual)
}

func Test_BasicUInt16Creator_CreateUsingMap(t *testing.T) {
	// Arrange
	m := map[uint16]string{0: "Off", 1: "On"}
	bu := enumimpl.New.BasicUInt16.CreateUsingMap("TestU16", m)

	// Act
	actual := args.Map{"len": bu.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "CreateUsingMap returns enum -- two entries", actual)
}

func Test_BasicUInt16Creator_UsingFirstItemSliceAliasMap(t *testing.T) {
	// Arrange
	type testEnum uint16
	aliases := map[string]uint16{"off": 0}
	bu := enumimpl.New.BasicUInt16.UsingFirstItemSliceAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bu.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UsingFirstItemSliceAliasMap returns enum -- with alias", actual)
}

func Test_BasicUInt16Creator_Default(t *testing.T) {
	// Arrange
	type testEnum uint16
	bu := enumimpl.New.BasicUInt16.Default(testEnum(0), []string{"Off", "On"})

	// Act
	actual := args.Map{"len": bu.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Default returns enum -- two entries", actual)
}

func Test_BasicUInt16Creator_DefaultWithAliasMap(t *testing.T) {
	// Arrange
	type testEnum uint16
	aliases := map[string]uint16{"off": 0}
	bu := enumimpl.New.BasicUInt16.DefaultWithAliasMap(testEnum(0), []string{"Off", "On"}, aliases)

	// Act
	actual := args.Map{"len": bu.Length()}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DefaultWithAliasMap returns enum -- with alias", actual)
}

// ── differCheckerImpl — uncovered branches ──

func Test_DifferCheckerImpl_GetSingleDiffResult_Left(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(true, "left", "right")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "left"}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns left -- isLeft true", actual)
}

func Test_DifferCheckerImpl_GetSingleDiffResult_Right(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetSingleDiffResult(false, "left", "right")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "right"}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns right -- isLeft false", actual)
}

func Test_DifferCheckerImpl_GetResultOnKeyMissing(t *testing.T) {
	// Arrange
	result := enumimpl.DefaultDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("k", "v")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": "v"}
	expected.ShouldBeEqual(t, 0, "GetResultOnKeyMissing returns lVal -- always returns left value", actual)
}

func Test_DifferCheckerImpl_IsEqual_Regardless(t *testing.T) {
	// Act
	actual := args.Map{
		"sameRegardless": enumimpl.DefaultDiffCheckerImpl.IsEqual(true, 1, int64(1)),
		"sameStrict":     enumimpl.DefaultDiffCheckerImpl.IsEqual(false, 1, 1),
		"diffStrict":     enumimpl.DefaultDiffCheckerImpl.IsEqual(false, 1, int64(1)),
	}

	// Assert
	expected := args.Map{
		"sameRegardless": true,
		"sameStrict":     true,
		"diffStrict":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsEqual returns correct -- regardless and strict modes", actual)
}

// ── leftRightDiffCheckerImpl — uncovered branches ──

func Test_LeftRightDiffChecker_GetSingleDiffResult_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetSingleDiffResult(true, "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetSingleDiffResult returns non-empty -- diff string", actual)
}

func Test_LeftRightDiffChecker_GetResultOnKeyMissing_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	result := enumimpl.LeftRightDiffCheckerImpl.GetResultOnKeyMissingInRightExistInLeft("k", "v")
	resultStr, ok := result.(string)

	// Act
	actual := args.Map{
		"ok": ok,
		"contains": len(resultStr) > 0,
	}

	// Assert
	expected := args.Map{
		"ok": true,
		"contains": true,
	}
	expected.ShouldBeEqual(t, 0, "GetResultOnKeyMissing returns formatted string -- key missing", actual)
}

func Test_LeftRightDiffChecker_IsEqual_FromBasicInt8IsAnyOf(t *testing.T) {
	// Act
	actual := args.Map{"equal": enumimpl.LeftRightDiffCheckerImpl.IsEqual(true, 1, int64(1))}

	// Assert
	expected := args.Map{"equal": true}
	expected.ShouldBeEqual(t, 0, "IsEqual returns true -- delegates to default regardless", actual)
}

// ── KeyAnyVal — uncovered branches ──

func Test_KeyAnyVal_Methods(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "TestKey", AnyValue: 42}

	// Act
	actual := args.Map{
		"key":        kav.KeyString(),
		"anyVal":     kav.AnyVal(),
		"anyValStr":  kav.AnyValString() != "",
		"wrapKey":    kav.WrapKey() != "",
		"wrapVal":    kav.WrapValue() != "",
		"isString":   kav.IsString(),
		"valInt":     kav.ValInt(),
		"stringVal":  kav.String() != "",
	}

	// Assert
	expected := args.Map{
		"key":        "TestKey",
		"anyVal":     42,
		"anyValStr":  true,
		"wrapKey":    true,
		"wrapVal":    true,
		"isString":   false,
		"valInt":     42,
		"stringVal":  true,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal methods return correct -- integer value", actual)
}

func Test_KeyAnyVal_IsString_True(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "TestKey", AnyValue: "hello"}

	// Act
	actual := args.Map{"isString": kav.IsString()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "IsString returns true -- string value", actual)
}

func Test_KeyAnyVal_IsString_MinInt(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "TestKey", AnyValue: constants.MinInt}

	// Act
	actual := args.Map{"isString": kav.IsString()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "IsString returns true -- MinInt value", actual)
}

func Test_KeyAnyVal_String_IsString(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "TestKey", AnyValue: "hello"}

	// Act
	actual := args.Map{"notEmpty": kav.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns formatted -- string value", actual)
}

func Test_KeyAnyVal_KeyValInteger(t *testing.T) {
	// Arrange
	kav := enumimpl.KeyAnyVal{Key: "TestKey", AnyValue: 42}
	kvi := kav.KeyValInteger()

	// Act
	actual := args.Map{
		"key": kvi.Key,
		"val": kvi.ValueInteger,
	}

	// Assert
	expected := args.Map{
		"key": "TestKey",
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger returns correct -- key and value", actual)
}

// ── KeyValInteger — uncovered branches ──

func Test_KeyValInteger_Methods(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "TestKey", ValueInteger: 42}

	// Act
	actual := args.Map{
		"wrapKey":   kvi.WrapKey() != "",
		"wrapVal":   kvi.WrapValue() != "",
		"isString":  kvi.IsString(),
		"stringVal": kvi.String() != "",
	}

	// Assert
	expected := args.Map{
		"wrapKey":   true,
		"wrapVal":   true,
		"isString":  false,
		"stringVal": true,
	}
	expected.ShouldBeEqual(t, 0, "KeyValInteger methods return correct -- integer value", actual)
}

func Test_KeyValInteger_IsString_MinInt(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "TestKey", ValueInteger: constants.MinInt}

	// Act
	actual := args.Map{"isString": kvi.IsString()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "IsString returns true -- MinInt value", actual)
}

func Test_KeyValInteger_String_IsString(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "TestKey", ValueInteger: constants.MinInt}

	// Act
	actual := args.Map{"notEmpty": kvi.String() != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns formatted -- string value at MinInt", actual)
}

func Test_KeyValInteger_KeyAnyVal(t *testing.T) {
	// Arrange
	kvi := enumimpl.KeyValInteger{Key: "TestKey", ValueInteger: 42}
	kav := kvi.KeyAnyVal()

	// Act
	actual := args.Map{
		"key": kav.Key,
		"val": kav.AnyValue,
	}

	// Assert
	expected := args.Map{
		"key": "TestKey",
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyVal returns correct -- key and value", actual)
}

// ── KeyAnyValues — uncovered branches ──

func Test_KeyAnyValues_Empty(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{}, []byte{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns empty -- empty names", actual)
}

func Test_KeyAnyValues_Items(t *testing.T) {
	// Arrange
	result := enumimpl.KeyAnyValues([]string{"A", "B"}, []byte{10, 20})

	// Act
	actual := args.Map{
		"len": len(result),
		"firstKey": result[0].Key,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"firstKey": "A",
	}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns mapped -- two items", actual)
}

// ── AllNameValues — uncovered branches ──

func Test_AllNameValues_Items(t *testing.T) {
	// Arrange
	result := enumimpl.AllNameValues([]string{"A", "B"}, []byte{10, 20})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns formatted -- two items", actual)
}

// ── Format — uncovered branches ──

func Test_Format_Template(t *testing.T) {
	// Arrange
	result := enumimpl.Format("MyEnum", "Active", "1", "Enum of {type-name} - {name} - {value}")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"contains": len(result) > 10,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"contains": true,
	}
	expected.ShouldBeEqual(t, 0, "Format returns formatted -- template with replacements", actual)
}

// ── IntegersRangesOfAnyVal ──

func Test_IntegersRangesOfAnyVal_ByteSlice(t *testing.T) {
	// Arrange
	result := enumimpl.IntegersRangesOfAnyVal([]byte{3, 1, 2})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
	}
	expected.ShouldBeEqual(t, 0, "IntegersRangesOfAnyVal returns sorted -- byte slice", actual)
}

// ── PrependJoin / JoinPrependUsingDot ──

func Test_PrependJoin_Items(t *testing.T) {
	// Arrange
	result := enumimpl.PrependJoin(".", "prefix", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "PrependJoin returns joined -- with prepend", actual)
}

func Test_JoinPrependUsingDot(t *testing.T) {
	// Arrange
	result := enumimpl.JoinPrependUsingDot("prefix", "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinPrependUsingDot returns joined -- with dot", actual)
}

// ── OnlySupportedErr / UnsupportedNames ──

func Test_OnlySupportedErr_AllSupported(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"A", "B"}, "A", "B")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- all supported", actual)
}

func Test_OnlySupportedErr_HasUnsupported(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{"A", "B", "C"}, "A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns error -- has unsupported", actual)
}

func Test_OnlySupportedErr_EmptyAll(t *testing.T) {
	// Arrange
	err := enumimpl.OnlySupportedErr(0, []string{}, "A")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- empty allNames", actual)
}

func Test_UnsupportedNames_Items(t *testing.T) {
	// Arrange
	result := enumimpl.UnsupportedNames([]string{"A", "B", "C"}, "A")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "UnsupportedNames returns unsupported -- two not in supported", actual)
}

// ── numberEnumBase methods — uncovered branches ──

func Test_NumberEnumBase_MinMaxAny_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	min, max := bb.MinMaxAny()

	// Act
	actual := args.Map{
		"minNotNil": min != nil,
		"maxNotNil": max != nil,
	}

	// Assert
	expected := args.Map{
		"minNotNil": true,
		"maxNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "MinMaxAny returns non-nil -- byte enum", actual)
}

func Test_NumberEnumBase_MinInt(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})

	// Act
	actual := args.Map{"min": bb.MinInt()}

	// Assert
	expected := args.Map{"min": 0}
	expected.ShouldBeEqual(t, 0, "MinInt returns 0 -- byte enum", actual)
}

func Test_NumberEnumBase_MaxInt(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})

	// Act
	actual := args.Map{"max": bb.MaxInt()}

	// Assert
	expected := args.Map{"max": 1}
	expected.ShouldBeEqual(t, 0, "MaxInt returns 1 -- byte enum two items", actual)
}

func Test_NumberEnumBase_AllNameValues_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	result := bb.AllNameValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns formatted -- two items", actual)
}

func Test_NumberEnumBase_RangesMap_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	m := bb.RangesMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesMap returns map -- two items", actual)
}

func Test_NumberEnumBase_OnlySupportedErr_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B", "C"})
	err := bb.OnlySupportedErr("A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns error -- unsupported exist", actual)
}

func Test_NumberEnumBase_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B", "C"})
	err := bb.OnlySupportedMsgErr("test msg", "A")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr returns error -- with message", actual)
}

func Test_NumberEnumBase_IntegerEnumRanges_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	result := bb.IntegerEnumRanges()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges returns ranges -- two items", actual)
}

func Test_NumberEnumBase_Count(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})

	// Act
	actual := args.Map{"count": bb.Count()}

	// Assert
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "Count returns 2 -- two items", actual)
}

func Test_NumberEnumBase_RangesDynamicMap_Cached(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	_ = bb.RangesDynamicMap()
	m := bb.RangesDynamicMap() // cached

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap returns cached -- second call", actual)
}

func Test_NumberEnumBase_DynamicMap_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	dm := bb.DynamicMap()

	// Act
	actual := args.Map{"len": len(dm)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "DynamicMap returns map -- two items", actual)
}

func Test_NumberEnumBase_RangesIntegerStringMap(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	m := bb.RangesIntegerStringMap()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangesIntegerStringMap returns map -- two items", actual)
}

func Test_NumberEnumBase_KeyAnyValues_Cached(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	_ = bb.KeyAnyValues()
	result := bb.KeyAnyValues() // cached

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyAnyValues returns cached -- second call", actual)
}

func Test_NumberEnumBase_KeyValIntegers(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	result := bb.KeyValIntegers()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "KeyValIntegers returns slice -- two items", actual)
}

func Test_NumberEnumBase_Loop_Break(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B", "C"})
	count := 0
	bb.Loop(func(index int, name string, anyVal any) bool {
		count++
		return index == 0 // break on first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "Loop breaks early -- break on first", actual)
}

func Test_NumberEnumBase_LoopInteger_Break(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B", "C"})
	count := 0
	bb.LoopInteger(func(index int, name string, anyVal int) bool {
		count++
		return index == 0 // break on first
	})

	// Act
	actual := args.Map{"count": count}

	// Assert
	expected := args.Map{"count": 1}
	expected.ShouldBeEqual(t, 0, "LoopInteger breaks early -- break on first", actual)
}

func Test_NumberEnumBase_TypeName(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("MyEnum", []string{"A"})

	// Act
	actual := args.Map{"name": bb.TypeName()}

	// Assert
	expected := args.Map{"name": "MyEnum"}
	expected.ShouldBeEqual(t, 0, "TypeName returns correct -- MyEnum", actual)
}

func Test_NumberEnumBase_ValueString(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A"})
	result := bb.ValueString(byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ValueString returns formatted -- byte 0", actual)
}

func Test_NumberEnumBase_Format_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A"})
	result := bb.Format("Enum {type-name} {name} {value}", byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Format returns formatted -- template", actual)
}

func Test_NumberEnumBase_RangeNamesCsv(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	result := bb.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv returns csv -- two items", actual)
}

func Test_NumberEnumBase_RangesInvalidMessage(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	result := bb.RangesInvalidMessage()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangesInvalidMessage returns message -- two items", actual)
}

func Test_NumberEnumBase_RangesInvalidErr(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})
	err := bb.RangesInvalidErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "RangesInvalidErr returns error -- always has message", actual)
}

func Test_NumberEnumBase_StringRangesPtr(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A", "B"})

	// Act
	actual := args.Map{"len": len(bb.StringRangesPtr())}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "StringRangesPtr returns slice -- two items", actual)
}

func Test_NumberEnumBase_JsonString_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A"})
	result := bb.JsonString(byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns formatted -- byte 0", actual)
}

func Test_NumberEnumBase_ToEnumString(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A"})
	result := bb.ToEnumString(byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToEnumString returns formatted -- byte 0", actual)
}

func Test_NumberEnumBase_ToName(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A"})
	result := bb.ToName(byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToName returns formatted -- byte 0", actual)
}

func Test_NumberEnumBase_NameWithValue_FromBasicInt8IsAnyOf(t *testing.T) {
	// Arrange
	bb := enumimpl.New.BasicByte.UsingTypeSlice("Test", []string{"A"})
	result := bb.NameWithValue(byte(0))

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "NameWithValue returns formatted -- byte 0", actual)
}

// ── DynamicMap — additional uncovered branches ──

func Test_DynamicMap_ConvMapIntegerString(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	m := dm.ConvMapIntegerString()

	// Act
	actual := args.Map{"len": len(m)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ConvMapIntegerString returns map -- two entries", actual)
}

func Test_DynamicMap_IsValueString_Int(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}

	// Act
	actual := args.Map{"isString": dm.IsValueString()}

	// Assert
	expected := args.Map{"isString": false}
	expected.ShouldBeEqual(t, 0, "IsValueString returns false -- int value", actual)
}

func Test_DynamicMap_IsValueString_String(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "x"}

	// Act
	actual := args.Map{"isString": dm.IsValueString()}

	// Assert
	expected := args.Map{"isString": true}
	expected.ShouldBeEqual(t, 0, "IsValueString returns true -- string value", actual)
}

func Test_DynamicMap_SortedKeyValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": 2, "a": 1}
	result := dm.SortedKeyValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyValues returns sorted -- two entries", actual)
}

func Test_DynamicMap_SortedKeyAnyValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"b": "y", "a": "x"}
	result := dm.SortedKeyAnyValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SortedKeyAnyValues returns sorted -- two entries", actual)
}

func Test_DynamicMap_String_Number(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	result := dm.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns formatted -- number map", actual)
}

func Test_DynamicMap_String_StringValues(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": "x"}
	result := dm.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns formatted -- string map", actual)
}

func Test_DynamicMap_DiffRaw_Same(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1}
	other := map[string]any{"a": 1}
	result := dm.DiffRaw(false, other)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns empty -- same maps", actual)
}

func Test_DynamicMap_DiffRaw_Different(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	other := map[string]any{"a": 1, "b": 3}
	result := dm.DiffRaw(false, other)

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns diffs -- different values", actual)
}

func Test_DynamicMap_DiffRaw_MissingKey(t *testing.T) {
	// Arrange
	dm := enumimpl.DynamicMap{"a": 1, "b": 2}
	other := map[string]any{"a": 1}
	result := dm.DiffRaw(false, other)

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "DiffRaw returns diffs -- missing key in right", actual)
}

// ── DiffLeftRight — additional uncovered branches ──

func Test_DiffLeftRight_DiffString_Different(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	result := dlr.DiffString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "DiffString returns non-empty -- different values", actual)
}

func Test_DiffLeftRight_JsonString_NonNil(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}
	result := dlr.JsonString()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JsonString returns non-empty -- non-nil pointer", actual)
}

func Test_DiffLeftRight_HasMismatch_SameType(t *testing.T) {
	// Arrange
	dlr := &enumimpl.DiffLeftRight{Left: "a", Right: "b"}

	// Act
	actual := args.Map{
		"mismatchStrict":     dlr.HasMismatch(false),
		"mismatchRegardless": dlr.HasMismatch(true),
	}

	// Assert
	expected := args.Map{
		"mismatchStrict":     true,
		"mismatchRegardless": true,
	}
	expected.ShouldBeEqual(t, 0, "HasMismatch returns true -- same type different values", actual)
}
