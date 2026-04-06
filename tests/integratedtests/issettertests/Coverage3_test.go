package issettertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ============================================================================
// GetSetByte
// ============================================================================

func Test_GetSetByte_True_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", issetter.GetSetByte(true, 1, 2))}

	// Assert
	expected := args.Map{"result": "True"}
	expected.ShouldBeEqual(t, 0, "GetSetByte returns Value(trueValue) -- true", actual)
}

func Test_GetSetByte_False_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": fmt.Sprintf("%v", issetter.GetSetByte(false, 1, 2))}

	// Assert
	expected := args.Map{"result": "False"}
	expected.ShouldBeEqual(t, 0, "GetSetByte returns Value(falseValue) -- false", actual)
}

// ============================================================================
// GetSetUnset
// ============================================================================

func Test_GetSetUnset_True_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetSetUnset(true) == issetter.Set}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns Set -- true", actual)
}

func Test_GetSetUnset_False_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetSetUnset(false) == issetter.Unset}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns Unset -- false", actual)
}

// ============================================================================
// GetSetterByComparing
// ============================================================================

func Test_GetSetterByComparing_Match_Cov3(t *testing.T) {
	// Arrange
	result := issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "x", "a", "b")

	// Act
	actual := args.Map{"result": result == issetter.True}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing returns trueVal -- match found", actual)
}

func Test_GetSetterByComparing_NoMatch_Cov3(t *testing.T) {
	// Arrange
	result := issetter.GetSetterByComparing(issetter.True, issetter.False, "z", "x", "a", "b")

	// Act
	actual := args.Map{"result": result == issetter.False}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing returns falseVal -- no match", actual)
}

// ============================================================================
// Value methods — logical checks
// ============================================================================

func Test_Value_IsLater_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"uninit":  issetter.Uninitialized.IsLater(),
		"true":    issetter.True.IsLater(),
		"wild":    issetter.Wildcard.IsLater(),
	}

	// Assert
	expected := args.Map{
		"uninit": true,
		"true": false,
		"wild": true,
	}
	expected.ShouldBeEqual(t, 0, "IsLater returns true for undefined -- Uninitialized/Wildcard", actual)
}

func Test_Value_IsNot_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"trueNotFalse": issetter.True.IsNot(issetter.False),
		"trueSame":     issetter.True.IsNot(issetter.True),
	}

	// Assert
	expected := args.Map{
		"trueNotFalse": true,
		"trueSame": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNot returns correct -- comparison", actual)
}

func Test_Value_IsNo_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"false": issetter.False.IsNo(),
		"unset": issetter.Unset.IsNo(),
		"true":  issetter.True.IsNo(),
	}

	// Assert
	expected := args.Map{
		"false": true,
		"unset": true,
		"true": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNo returns true for False/Unset -- logical no", actual)
}

func Test_Value_IsAsk_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"uninit": issetter.Uninitialized.IsAsk(),
		"wild":   issetter.Wildcard.IsAsk(),
		"true":   issetter.True.IsAsk(),
	}

	// Assert
	expected := args.Map{
		"uninit": true,
		"wild": true,
		"true": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAsk returns true for undefined -- Uninitialized/Wildcard", actual)
}

func Test_Value_IsIndeterminate_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"uninit": issetter.Uninitialized.IsIndeterminate(),
		"false":  issetter.False.IsIndeterminate(),
	}

	// Assert
	expected := args.Map{
		"uninit": true,
		"false": false,
	}
	expected.ShouldBeEqual(t, 0, "IsIndeterminate same as IsAsk -- Uninitialized", actual)
}

func Test_Value_IsAccept_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.IsAccept(),
		"set":  issetter.Set.IsAccept(),
		"unset": issetter.Unset.IsAccept(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"set": true,
		"unset": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAccept returns true for True/Set -- logical accept", actual)
}

func Test_Value_IsReject_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"false": issetter.False.IsReject(),
		"true":  issetter.True.IsReject(),
	}

	// Assert
	expected := args.Map{
		"false": true,
		"true": false,
	}
	expected.ShouldBeEqual(t, 0, "IsReject returns true for False -- logical reject", actual)
}

func Test_Value_IsFailed_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsFailed()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsFailed returns true for False -- failure", actual)
}

func Test_Value_IsSuccess_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsSuccess()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsSuccess returns true for True -- success", actual)
}

func Test_Value_IsSkip_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"wild":   issetter.Wildcard.IsSkip(),
		"uninit": issetter.Uninitialized.IsSkip(),
		"set":    issetter.Set.IsSkip(),
	}

	// Assert
	expected := args.Map{
		"wild": true,
		"uninit": true,
		"set": false,
	}
	expected.ShouldBeEqual(t, 0, "IsSkip returns true for undefined -- Wildcard/Uninitialized", actual)
}

// ============================================================================
// Value methods — name/string
// ============================================================================

func Test_Value_NameValue_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.NameValue()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "NameValue returns non-empty -- True", actual)
}

func Test_Value_IsNameEqual_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   issetter.True.IsNameEqual("True"),
		"noMatch": issetter.True.IsNameEqual("False"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNameEqual matches name -- True", actual)
}

func Test_Value_IsAnyNamesOf_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   issetter.True.IsAnyNamesOf("False", "True"),
		"noMatch": issetter.True.IsAnyNamesOf("False", "Unset"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyNamesOf checks multiple names -- True", actual)
}

func Test_Value_ToNumberString_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ToNumberString()}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "ToNumberString returns 1 -- True", actual)
}

// ============================================================================
// Value methods — type conversions
// ============================================================================

func Test_Value_ValueByte_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueByte()}

	// Assert
	expected := args.Map{"result": byte(1)}
	expected.ShouldBeEqual(t, 0, "ValueByte returns 1 -- True", actual)
}

func Test_Value_ValueInt_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.ValueInt()}

	// Assert
	expected := args.Map{"result": 2}
	expected.ShouldBeEqual(t, 0, "ValueInt returns 2 -- False", actual)
}

func Test_Value_ValueInt8_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt8()}

	// Assert
	expected := args.Map{"result": int8(1)}
	expected.ShouldBeEqual(t, 0, "ValueInt8 returns 1 -- True", actual)
}

func Test_Value_ValueInt16_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt16()}

	// Assert
	expected := args.Map{"result": int16(1)}
	expected.ShouldBeEqual(t, 0, "ValueInt16 returns 1 -- True", actual)
}

func Test_Value_ValueInt32_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt32()}

	// Assert
	expected := args.Map{"result": int32(1)}
	expected.ShouldBeEqual(t, 0, "ValueInt32 returns 1 -- True", actual)
}

func Test_Value_ValueString_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueString()}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "ValueString returns 1 -- True", actual)
}

func Test_Value_ValueUInt16_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueUInt16()}

	// Assert
	expected := args.Map{"result": uint16(1)}
	expected.ShouldBeEqual(t, 0, "ValueUInt16 returns 1 -- True", actual)
}

func Test_Value_StringValue_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.StringValue()}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "StringValue returns 1 -- True", actual)
}

func Test_Value_String_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.String()}

	// Assert
	expected := args.Map{"result": "True"}
	expected.ShouldBeEqual(t, 0, "String returns True -- True", actual)
}

// ============================================================================
// Value methods — state checks
// ============================================================================

func Test_Value_IsTrueOrSet_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.IsTrueOrSet(),
		"set":  issetter.Set.IsTrueOrSet(),
		"false": issetter.False.IsTrueOrSet(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"set": true,
		"false": false,
	}
	expected.ShouldBeEqual(t, 0, "IsTrueOrSet returns true -- True/Set", actual)
}

func Test_Value_HasInitialized_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"true":   issetter.True.HasInitialized(),
		"uninit": issetter.Uninitialized.HasInitialized(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"uninit": false,
	}
	expected.ShouldBeEqual(t, 0, "HasInitialized returns false for Uninitialized -- check", actual)
}

func Test_Value_HasInitializedAndSet_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"set":  issetter.Set.HasInitializedAndSet(),
		"true": issetter.True.HasInitializedAndSet(),
	}

	// Assert
	expected := args.Map{
		"set": true,
		"true": false,
	}
	expected.ShouldBeEqual(t, 0, "HasInitializedAndSet checks Set only -- Set vs True", actual)
}

func Test_Value_HasInitializedAndTrue_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.HasInitializedAndTrue(),
		"set":  issetter.Set.HasInitializedAndTrue(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"set": false,
	}
	expected.ShouldBeEqual(t, 0, "HasInitializedAndTrue checks True only -- True vs Set", actual)
}

// ============================================================================
// Value methods — enum support
// ============================================================================

func Test_Value_AllNameValues_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.AllNameValues()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "AllNameValues returns non-empty -- enum list", actual)
}

func Test_Value_IntegerEnumRanges_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.IntegerEnumRanges()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges returns non-empty -- ranges", actual)
}

func Test_Value_MinMaxAny_Cov3(t *testing.T) {
	// Arrange
	min, max := issetter.True.MinMaxAny()

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
	expected.ShouldBeEqual(t, 0, "MinMaxAny returns non-nil -- min/max", actual)
}

func Test_Value_MinValueString_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MinValueString()}

	// Assert
	expected := args.Map{"result": issetter.Uninitialized.StringValue()}
	expected.ShouldBeEqual(t, 0, "MinValueString returns Uninitialized value -- min", actual)
}

func Test_Value_MaxValueString_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.MaxValueString()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "MaxValueString returns non-empty -- max", actual)
}

func Test_Value_MaxInt_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MaxInt()}

	// Assert
	expected := args.Map{"result": issetter.Wildcard.ValueInt()}
	expected.ShouldBeEqual(t, 0, "MaxInt returns Wildcard int -- max", actual)
}

func Test_Value_MinInt_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MinInt()}

	// Assert
	expected := args.Map{"result": 0}
	expected.ShouldBeEqual(t, 0, "MinInt returns 0 -- min", actual)
}

func Test_Value_RangesDynamicMap_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.RangesDynamicMap()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap returns non-empty -- dynamic map", actual)
}

func Test_Value_IsValueEqual_Cov3(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   issetter.True.IsValueEqual(1),
		"noMatch": issetter.True.IsValueEqual(2),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsValueEqual compares byte value -- True", actual)
}

func Test_Value_RangeNamesCsv_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.RangeNamesCsv()

	// Act
	actual := args.Map{"hasContent": len(result) > 0}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv returns non-empty -- csv", actual)
}

func Test_Value_IsByteValueEqual_Cov3(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsByteValueEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsByteValueEqual returns true -- matching byte", actual)
}

func Test_Value_Format_Cov3(t *testing.T) {
	// Arrange
	result := issetter.True.Format("{name}={value}")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "True=1"}
	expected.ShouldBeEqual(t, 0, "Format replaces placeholders -- {name}={value}", actual)
}

func Test_Value_EnumType_Cov3(t *testing.T) {
	// Arrange
	et := issetter.True.EnumType()

	// Act
	actual := args.Map{"notNil": et != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "EnumType returns non-nil -- Byte enum type", actual)
}

func Test_Value_OnlySupportedErr_Cov3(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr("True", "False", "Uninitialized", "Unset", "Set", "Wildcard")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- all supported", actual)
}

func Test_Value_OnlySupportedErr_Missing_Cov3(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr("True")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns error -- missing names", actual)
}

func Test_Value_OnlySupportedErr_Empty_Cov3(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "OnlySupportedErr returns nil -- empty names", actual)
}

func Test_Value_OnlySupportedMsgErr_Cov3(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr returns error with prefix -- missing names", actual)
}

func Test_Value_OnlySupportedMsgErr_NoErr_Cov3(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True", "False", "Uninitialized", "Unset", "Set", "Wildcard")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "OnlySupportedMsgErr returns nil -- all supported", actual)
}
