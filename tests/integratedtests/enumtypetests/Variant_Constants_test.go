package enumtypetests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl/enumtype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Variant_Constants(t *testing.T) {
	// Assert
	actual := args.Map{"result": enumtype.Invalid != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invalid should be 0", actual)
	actual = args.Map{"result": enumtype.Boolean != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Boolean should be 1", actual)
	actual = args.Map{"result": enumtype.String != 11}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should be 11", actual)
}

func Test_Variant_Name(t *testing.T) {
	// Act & Assert
	actual := args.Map{"result": enumtype.Boolean.Name() != "Boolean"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Boolean", actual)
	actual = args.Map{"result": enumtype.Integer.String() != "Integer"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Integer String mismatch", actual)
	actual = args.Map{"result": enumtype.Byte.NameValue() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameValue should not be empty", actual)
}

func Test_Variant_TypeChecks(t *testing.T) {
	// Assert
	actual := args.Map{"result": enumtype.Boolean.IsBoolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Boolean.IsBoolean should be true", actual)
	actual = args.Map{"result": enumtype.Byte.IsByte()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte.IsByte should be true", actual)
	actual = args.Map{"result": enumtype.UnsignedInteger16.IsUnsignedInteger16()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "UnsignedInteger16 check failed", actual)
	actual = args.Map{"result": enumtype.UnsignedInteger32.IsUnsignedInteger32()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "UnsignedInteger32 check failed", actual)
	actual = args.Map{"result": enumtype.UnsignedInteger64.IsUnsignedInteger64()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "UnsignedInteger64 check failed", actual)
	actual = args.Map{"result": enumtype.Integer8.IsInteger8()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer8 check failed", actual)
	actual = args.Map{"result": enumtype.Integer16.IsInteger16()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer16 check failed", actual)
	actual = args.Map{"result": enumtype.Integer32.IsInteger32()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer32 check failed", actual)
	actual = args.Map{"result": enumtype.Integer64.IsInteger64()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer64 check failed", actual)
	actual = args.Map{"result": enumtype.Integer.IsInteger()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer check failed", actual)
	actual = args.Map{"result": enumtype.String.IsString()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "String check failed", actual)
}

func Test_Variant_IsNumber(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Integer.IsNumber()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer should be number", actual)
	actual = args.Map{"result": enumtype.Boolean.IsNumber()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Boolean should not be number", actual)
	actual = args.Map{"result": enumtype.String.IsNumber()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be number", actual)
}

func Test_Variant_IsAnyInteger(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Integer.IsAnyInteger()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Integer should be any integer", actual)
	actual = args.Map{"result": enumtype.Byte.IsAnyInteger()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Byte should not be any integer", actual)
}

func Test_Variant_IsAnyUnsignedNumber(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Byte.IsAnyUnsignedNumber()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Byte should be unsigned", actual)
	actual = args.Map{"result": enumtype.Integer.IsAnyUnsignedNumber()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Integer should not be unsigned", actual)
}

func Test_Variant_ValidInvalid(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Invalid.IsValid()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Invalid should not be valid", actual)
	actual = args.Map{"result": enumtype.Invalid.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be invalid", actual)
	actual = args.Map{"result": enumtype.Boolean.IsValid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Boolean should be valid", actual)
	actual = args.Map{"result": enumtype.Boolean.IsInvalid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Boolean should not be invalid", actual)
}

func Test_Variant_ValueConversions(t *testing.T) {
	// Arrange
	v := enumtype.Integer // 10

	// Act
	actual := args.Map{"result": v.Value() != 10}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Value mismatch", actual)
	actual = args.Map{"result": v.ValueByte() != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueByte mismatch", actual)
	actual = args.Map{"result": v.ValueInt() != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt mismatch", actual)
	actual = args.Map{"result": v.ValueInt8() != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt8 mismatch", actual)
	actual = args.Map{"result": v.ValueInt16() != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt16 mismatch", actual)
	actual = args.Map{"result": v.ValueInt32() != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt32 mismatch", actual)
	actual = args.Map{"result": v.ValueUInt16() != 10}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueUInt16 mismatch", actual)
	actual = args.Map{"result": v.ValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueString should not be empty", actual)
	actual = args.Map{"result": v.ToNumberString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToNumberString should not be empty", actual)
}

func Test_Variant_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.IsNameEqual("Boolean")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNameEqual should be true for Boolean", actual)
	actual = args.Map{"result": enumtype.Boolean.IsNameEqual("String")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsNameEqual should be false", actual)
}

func Test_Variant_IsAnyNamesOf(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.IsAnyNamesOf("String", "Boolean")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf should find Boolean", actual)
	actual = args.Map{"result": enumtype.Boolean.IsAnyNamesOf("String", "Integer")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf should not find Boolean", actual)
}

func Test_Variant_TypeName(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.TypeName() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeName should not be empty", actual)
}

func Test_Variant_RangeNamesCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.RangeNamesCsv() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
}

func Test_Variant_MinMaxAny(t *testing.T) {
	// Arrange
	min, max := enumtype.Boolean.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinMaxAny should not return nil", actual)
}

func Test_Variant_MinMaxStrings(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.MinValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)
	actual = args.Map{"result": enumtype.Boolean.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)
}

func Test_Variant_MinMaxInt(t *testing.T) {
	// Act
	actual := args.Map{"result": enumtype.Boolean.MaxInt() != enumtype.String.ValueInt()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxInt mismatch", actual)
	actual = args.Map{"result": enumtype.Boolean.MinInt() != enumtype.Invalid.ValueInt()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinInt mismatch", actual)
}

func Test_Variant_RangesDynamicMap(t *testing.T) {
	// Arrange
	m := enumtype.Boolean.RangesDynamicMap()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap should not be empty", actual)
}

func Test_Variant_IntegerEnumRanges(t *testing.T) {
	// Arrange
	ranges := enumtype.Boolean.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(ranges) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges should not be empty", actual)
}

func Test_Variant_EnumType(t *testing.T) {
	// Arrange
	et := enumtype.Boolean.EnumType()

	// Act
	actual := args.Map{"result": et == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EnumType should not be nil", actual)
}

func Test_Variant_MarshalJSON(t *testing.T) {
	// Arrange
	data, err := json.Marshal(enumtype.Boolean)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)
	actual = args.Map{"result": len(data) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON should not be empty", actual)
}

func Test_Variant_UnmarshalJSON(t *testing.T) {
	// Arrange
	var v enumtype.Variant
	err := json.Unmarshal([]byte(`"Boolean"`), &v)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
	actual = args.Map{"result": v != enumtype.Boolean}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON should parse to Boolean", actual)
}

func Test_Variant_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var v enumtype.Variant
	err := json.Unmarshal([]byte(`""`), &v)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON should error on empty", actual)

	err = json.Unmarshal([]byte(`"NonExistent"`), &v)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON should error on nonexistent", actual)
}

func Test_Variant_Format_Panics(t *testing.T) {
	// Arrange
	defer func() {

	// Act
		r := recover()
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "Format should panic", actual)
	}()

	enumtype.Boolean.Format("{name}")
}
