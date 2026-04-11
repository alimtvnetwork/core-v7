package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ── New ──

func Test_New_ValidName(t *testing.T) {
	// Arrange
	v, err := issetter.New("True")

	// Act
	actual := args.Map{
		"val": v,
		"isNilErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"val": issetter.True,
		"isNilErr": true,
	}
	expected.ShouldBeEqual(t, 0, "New returns non-empty -- valid name", actual)
}

func Test_New_InvalidName(t *testing.T) {
	// Arrange
	_, err := issetter.New("bogus")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New returns error -- invalid name", actual)
}

// ── NewBool ──

func Test_NewBool_True(t *testing.T) {
	// Arrange
	v := issetter.NewBool(true)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "NewBool returns non-empty -- true", actual)
}

func Test_NewBool_False(t *testing.T) {
	// Arrange
	v := issetter.NewBool(false)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "NewBool returns non-empty -- false", actual)
}

// ── NewBooleans ──

func Test_NewBooleans_AllTrue(t *testing.T) {
	// Arrange
	v := issetter.NewBooleans(true, true)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "NewBooleans returns non-empty -- all true", actual)
}

func Test_NewBooleans_AnyFalse(t *testing.T) {
	// Arrange
	v := issetter.NewBooleans(true, false)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "NewBooleans returns non-empty -- any false", actual)
}

// ── CombinedBooleans ──

func Test_CombinedBooleans_AllTrue(t *testing.T) {
	// Arrange
	v := issetter.CombinedBooleans(true, true, true)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans returns non-empty -- all true", actual)
}

func Test_CombinedBooleans_HasFalse(t *testing.T) {
	// Arrange
	v := issetter.CombinedBooleans(true, false, true)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans returns non-empty -- has false", actual)
}

// ── GetBool ──

func Test_GetBool_True(t *testing.T) {
	// Arrange
	v := issetter.GetBool(true)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "GetBool returns non-empty -- true", actual)
}

func Test_GetBool_False(t *testing.T) {
	// Arrange
	v := issetter.GetBool(false)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "GetBool returns non-empty -- false", actual)
}

// ── GetSet ──

func Test_GetSet_True(t *testing.T) {
	// Arrange
	v := issetter.GetSet(true, issetter.True, issetter.False)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "GetSet returns non-empty -- true", actual)
}

func Test_GetSet_False(t *testing.T) {
	// Arrange
	v := issetter.GetSet(false, issetter.True, issetter.False)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "GetSet returns non-empty -- false", actual)
}

// ── GetSetUnset ──

func Test_GetSetUnset_True(t *testing.T) {
	// Arrange
	v := issetter.GetSetUnset(true)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.Set}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns non-empty -- true", actual)
}

func Test_GetSetUnset_False(t *testing.T) {
	// Arrange
	v := issetter.GetSetUnset(false)

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.Unset}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns non-empty -- false", actual)
}

// ── GetSetterByComparing ──

func Test_GetSetterByComparing_Match(t *testing.T) {
	// Arrange
	v := issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "a", "b")

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing returns correct value -- match", actual)
}

func Test_GetSetterByComparing_NoMatch(t *testing.T) {
	// Arrange
	v := issetter.GetSetterByComparing(issetter.True, issetter.False, "x", "a", "b")

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.False}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing returns empty -- no match", actual)
}

// ── Value methods ──

func Test_Value_IsTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsTrue()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsTrue returns non-empty -- with args", actual)
}

func Test_Value_IsFalse(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsFalse()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsFalse returns non-empty -- with args", actual)
}

func Test_Value_IsSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsSet returns correct value -- with args", actual)
}

func Test_Value_IsUnset(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Unset.IsUnset()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsUnset returns correct value -- with args", actual)
}

func Test_Value_IsWildcard(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsWildcard()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcard returns correct value -- with args", actual)
}

func Test_Value_HasInitialized_FromNewValidName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.HasInitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.HasInitialized returns correct value -- with args", actual)
}

func Test_Value_IsInvalid(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsInvalid returns error -- with args", actual)
}

func Test_Value_IsValid_FromNewValidName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsValid returns non-empty -- with args", actual)
}

func Test_Value_IsOn(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsOn()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsOn returns correct value -- with args", actual)
}

func Test_Value_IsOff(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsOff()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsOff returns correct value -- with args", actual)
}

func Test_Value_Boolean(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Boolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.Boolean returns correct value -- with args", actual)
}

func Test_Value_String_FromNewValidName(t *testing.T) {
	// Arrange
	result := issetter.True.String()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.String returns correct value -- with args", actual)
}

func Test_Value_ValueByte(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueByte()}

	// Assert
	expected := args.Map{"result": byte(1)}
	expected.ShouldBeEqual(t, 0, "Value.ValueByte returns correct value -- with args", actual)
}

func Test_Value_ValueInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ValueInt()}

	// Assert
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "Value.ValueInt returns correct value -- with args", actual)
}

func Test_Value_IsNot_FromNewValidName(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsNot returns correct value -- with args", actual)
}

func Test_Value_IsAccept(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsAccept()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsAccept returns correct value -- with args", actual)
}

func Test_Value_IsReject(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Unset.IsReject()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsReject returns correct value -- with args", actual)
}

func Test_Value_IsYes(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsYes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsYes returns correct value -- with args", actual)
}

func Test_Value_IsTrueOrSet_FromNewValidName(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.IsTrueOrSet(),
		"set": issetter.Set.IsTrueOrSet(),
		"false": issetter.False.IsTrueOrSet(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"set": true,
		"false": false,
	}
	expected.ShouldBeEqual(t, 0, "Value.IsTrueOrSet returns non-empty -- with args", actual)
}

func Test_Value_IsInitBoolean(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.IsInitBoolean(),
		"set": issetter.Set.IsInitBoolean(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"set": false,
	}
	expected.ShouldBeEqual(t, 0, "Value.IsInitBoolean returns correct value -- with args", actual)
}

func Test_Value_IsDefinedLogically(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.IsDefinedLogically(),
		"wild": issetter.Wildcard.IsDefinedLogically(),
	}

	// Assert
	expected := args.Map{
		"true": true,
		"wild": false,
	}
	expected.ShouldBeEqual(t, 0, "Value.IsDefinedLogically returns correct value -- with args", actual)
}

func Test_Value_IsUndefinedLogically(t *testing.T) {
	// Act
	actual := args.Map{
		"wild": issetter.Wildcard.IsUndefinedLogically(),
		"true": issetter.True.IsUndefinedLogically(),
	}

	// Assert
	expected := args.Map{
		"wild": true,
		"true": false,
	}
	expected.ShouldBeEqual(t, 0, "Value.IsUndefinedLogically returns correct value -- with args", actual)
}

func Test_Value_ToBooleanValue(t *testing.T) {
	// Arrange
	v := issetter.Set.ToBooleanValue()

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.True}
	expected.ShouldBeEqual(t, 0, "Value.ToBooleanValue returns correct value -- with args", actual)
}

func Test_Value_ToSetUnsetValue(t *testing.T) {
	// Arrange
	v := issetter.True.ToSetUnsetValue()

	// Act
	actual := args.Map{"val": v}

	// Assert
	expected := args.Map{"val": issetter.Set}
	expected.ShouldBeEqual(t, 0, "Value.ToSetUnsetValue returns correct value -- with args", actual)
}

func Test_Value_NameValue(t *testing.T) {
	// Arrange
	result := issetter.True.NameValue()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.NameValue returns correct value -- with args", actual)
}

func Test_Value_ToNumberString(t *testing.T) {
	// Arrange
	result := issetter.True.ToNumberString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "Value.ToNumberString returns correct value -- with args", actual)
}

func Test_Value_ValueString(t *testing.T) {
	// Arrange
	result := issetter.True.ValueString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "1"}
	expected.ShouldBeEqual(t, 0, "Value.ValueString returns non-empty -- with args", actual)
}

func Test_Value_IsWildcardOrBool_Wildcard(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsWildcardOrBool(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcardOrBool returns correct value -- wildcard", actual)
}

func Test_Value_IsWildcardOrBool_True(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsWildcardOrBool(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcardOrBool returns non-empty -- true", actual)
}

func Test_Value_IsWildcardOrBool_False(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.False.IsWildcardOrBool(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsWildcardOrBool returns non-empty -- false", actual)
}

func Test_Value_ToByteCondition(t *testing.T) {
	// Act
	actual := args.Map{
		"true": issetter.True.ToByteCondition(1, 0, 255),
		"false": issetter.False.ToByteCondition(1, 0, 255),
		"uninit": issetter.Uninitialized.ToByteCondition(1, 0, 255),
	}

	// Assert
	expected := args.Map{
		"true": byte(1),
		"false": byte(0),
		"uninit": byte(255),
	}
	expected.ShouldBeEqual(t, 0, "Value.ToByteCondition returns correct value -- with args", actual)
}

func Test_Value_Format_FromNewValidName(t *testing.T) {
	// Arrange
	result := issetter.True.Format("{name}={value}")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.Format returns correct value -- with args", actual)
}

func Test_Value_IsNameEqual_FromNewValidName(t *testing.T) {
	// Arrange
	name := issetter.True.String()

	// Act
	actual := args.Map{"result": issetter.True.IsNameEqual(name)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsNameEqual returns correct value -- with args", actual)
}

func Test_Value_IsAnyNamesOf_Match(t *testing.T) {
	// Arrange
	name := issetter.True.String()

	// Act
	actual := args.Map{"result": issetter.True.IsAnyNamesOf(name, "other")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value.IsAnyNamesOf returns correct value -- match", actual)
}

func Test_Value_IsAnyNamesOf_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("bogus")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Value.IsAnyNamesOf returns empty -- no match", actual)
}

func Test_Value_AllNameValues_FromNewValidName(t *testing.T) {
	// Arrange
	result := issetter.True.AllNameValues()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Value.AllNameValues returns non-empty -- with args", actual)
}

func Test_Value_RangeNamesCsv_FromNewValidName(t *testing.T) {
	// Arrange
	result := issetter.True.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value.RangeNamesCsv returns correct value -- with args", actual)
}
