package issettertests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── GetBool / GetSet ──

func Test_GetBool_Coverage(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.GetBool(true) != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "true should return True", actual)
	actual = args.Map{"result": issetter.GetBool(false) != issetter.False}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "false should return False", actual)
}

func Test_GetSet_Coverage(t *testing.T) {
	// Arrange
	result := issetter.GetSet(true, issetter.Set, issetter.Unset)

	// Act
	actual := args.Map{"result": result != issetter.Set}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "true condition should return Set", actual)

	result = issetter.GetSet(false, issetter.Set, issetter.Unset)
	actual = args.Map{"result": result != issetter.Unset}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "false condition should return Unset", actual)
}

// ── Value constants and basic checks ──

func Test_Value_Constants(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be 0", actual)
	actual = args.Map{"result": issetter.True != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should be 1", actual)
	actual = args.Map{"result": issetter.False != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False should be 2", actual)
}

func Test_Value_BoolChecks(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsTrue()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsTrue", actual)
	actual = args.Map{"result": issetter.False.IsFalse()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False.IsFalse", actual)
	actual = args.Map{"result": issetter.Set.IsSet()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set.IsSet", actual)
	actual = args.Map{"result": issetter.Unset.IsUnset()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Unset.IsUnset", actual)
	actual = args.Map{"result": issetter.Wildcard.IsWildcard()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard.IsWildcard", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsUninitialized()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized.IsUninitialized", actual)
}

func Test_Value_LogicalGroups(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsOn()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be On", actual)
	actual = args.Map{"result": issetter.Set.IsOn()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set should be On", actual)
	actual = args.Map{"result": issetter.False.IsOff()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be Off", actual)
	actual = args.Map{"result": issetter.Unset.IsOff()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Unset should be Off", actual)
	actual = args.Map{"result": issetter.True.IsSuccess()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be Success", actual)
	actual = args.Map{"result": issetter.False.IsFailed()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be Failed", actual)
	actual = args.Map{"result": issetter.True.IsAccept()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be Accept", actual)
	actual = args.Map{"result": issetter.False.IsReject()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be Reject", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsLater()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be Later", actual)
	actual = args.Map{"result": issetter.Wildcard.IsAsk()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard should be Ask", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsIndeterminate()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be Indeterminate", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsSkip()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be Skip", actual)
}

func Test_Value_DefinedUndefined(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsDefinedLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be defined", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsDefinedLogically()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should not be defined", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsUndefinedLogically()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be undefined", actual)
}

func Test_Value_Conversions(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Value() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.Value() should be 1", actual)
	actual = args.Map{"result": issetter.True.ValueByte() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.ValueByte() should be 1", actual)
	actual = args.Map{"result": issetter.True.ValueInt() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.ValueInt() should be 1", actual)
	actual = args.Map{"result": issetter.True.ValueInt8() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.ValueInt8() should be 1", actual)
	actual = args.Map{"result": issetter.True.ValueInt16() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.ValueInt16() should be 1", actual)
	actual = args.Map{"result": issetter.True.ValueInt32() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.ValueInt32() should be 1", actual)
	actual = args.Map{"result": issetter.True.ValueUInt16() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.ValueUInt16() should be 1", actual)
}

func Test_Value_String(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.String() != "True"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.String() should be True", actual)
	actual = args.Map{"result": issetter.True.Name() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Name should not be empty", actual)
	actual = args.Map{"result": issetter.True.NameValue() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameValue should not be empty", actual)
	actual = args.Map{"result": issetter.True.ValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueString should not be empty", actual)
	actual = args.Map{"result": issetter.True.StringValue() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "StringValue should not be empty", actual)
	actual = args.Map{"result": issetter.True.ToNumberString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToNumberString should not be empty", actual)
}

func Test_Value_IsNot(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True is not False", actual)
	actual = args.Map{"result": issetter.True.IsNot(issetter.True)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True is True", actual)
}

func Test_Value_Init(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsInit()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be Init", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsInit()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should not be Init", actual)
	actual = args.Map{"result": issetter.True.IsInitBoolean()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be InitBoolean", actual)
	actual = args.Map{"result": issetter.False.IsInitBoolean()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be InitBoolean", actual)
	actual = args.Map{"result": issetter.True.IsInitBooleanWild()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be InitBooleanWild", actual)
	actual = args.Map{"result": issetter.Set.IsInitSet()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set should be InitSet", actual)
	actual = args.Map{"result": issetter.Set.IsInitSetWild()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set should be InitSetWild", actual)
}

func Test_Value_ToBooleanSetUnset(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.ToBooleanValue() != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Set should convert to True", actual)
	actual = args.Map{"result": issetter.Unset.ToBooleanValue() != issetter.False}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Unset should convert to False", actual)
	actual = args.Map{"result": issetter.True.ToSetUnsetValue() != issetter.Set}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should convert to Set", actual)
	actual = args.Map{"result": issetter.False.ToSetUnsetValue() != issetter.Unset}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False should convert to Unset", actual)
}

func Test_Value_BooleanOp(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Boolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.Boolean should be true", actual)
	actual = args.Map{"result": issetter.False.Boolean()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False.Boolean should be false", actual)
	actual = args.Map{"result": issetter.True.IsYes()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsYes should be true", actual)
}

func Test_Value_WildcardApply_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.WildcardApply(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard.WildcardApply(true) should return true", actual)
	actual = args.Map{"result": issetter.True.WildcardApply(false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.WildcardApply should return true", actual)
}

func Test_Value_OrBool(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.OrBool(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.OrBool(false) should be true", actual)
	actual = args.Map{"result": issetter.Wildcard.OrBool(true)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard.OrBool(true) should return true", actual)
}

func Test_Value_AndBool(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.AndBool(true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.AndBool(true) should be true", actual)
	actual = args.Map{"result": issetter.True.AndBool(false)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.AndBool(false) should be false", actual)
}

func Test_Value_And(t *testing.T) {
	// Arrange
	result := issetter.True.And(issetter.True)

	// Act
	actual := args.Map{"result": result != issetter.True}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.And(True) should be True", actual)
	result = issetter.Wildcard.And(issetter.False)
	actual = args.Map{"result": result != issetter.False}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard.And(False) should be False", actual)
}

func Test_Value_ToByteCondition_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.ToByteCondition(1, 0, 255) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should return trueVal", actual)
	actual = args.Map{"result": issetter.False.ToByteCondition(1, 0, 255) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False should return falseVal", actual)
	actual = args.Map{"result": issetter.Wildcard.ToByteCondition(1, 0, 255) != 255}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard should return invalid", actual)
}

func Test_Value_ToByteConditionWithWildcard_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.ToByteConditionWithWildcard(99, 1, 0, 255) != 99}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard should return wildcard val", actual)
	actual = args.Map{"result": issetter.True.ToByteConditionWithWildcard(99, 1, 0, 255) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should return trueVal", actual)
}

func Test_Value_IsValid(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be valid", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be invalid", actual)
}

func Test_Value_Format(t *testing.T) {
	// Arrange
	result := issetter.True.Format("{name}={value}")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should not be empty", actual)
}

func Test_Value_EnumType(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.EnumType() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EnumType should not be nil", actual)
}

func Test_Value_AllNameValues(t *testing.T) {
	// Arrange
	names := issetter.True.AllNameValues()

	// Act
	actual := args.Map{"result": len(names) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNameValues should not be empty", actual)
}

func Test_Value_RangeNamesCsv(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.RangeNamesCsv() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
}

func Test_Value_MinMaxAny(t *testing.T) {
	// Arrange
	min, max := issetter.True.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinMaxAny should not return nil", actual)
}

func Test_Value_MinMaxStrings(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MinValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)
	actual = args.Map{"result": issetter.True.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)
}

func Test_Value_IntegerEnumRanges(t *testing.T) {
	// Arrange
	ranges := issetter.True.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(ranges) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges should not be empty", actual)
}

func Test_Value_RangesDynamicMap(t *testing.T) {
	// Arrange
	m := issetter.True.RangesDynamicMap()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap should not be empty", actual)
}

func Test_Value_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsNameEqual("True")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsNameEqual(True) should be true", actual)
}

func Test_Value_IsAnyNamesOf(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("False", "True")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find True", actual)
}

func Test_Value_JSON(t *testing.T) {
	// Arrange
	data, err := json.Marshal(issetter.True)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)
	actual = args.Map{"result": len(data) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON should not be empty", actual)

	var v issetter.Value
	err = json.Unmarshal([]byte(`"True"`), &v)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
	actual = args.Map{"result": v != issetter.True}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should unmarshal to True", actual)
}

func Test_Value_LazyEvaluateBool_Cov(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := v.LazyEvaluateBool(func() {})

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be called on Uninitialized", actual)

	called = v.LazyEvaluateBool(func() {})
	actual = args.Map{"result": called}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be called again", actual)
}

func Test_Value_LazyEvaluateSet_Cov(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := v.LazyEvaluateSet(func() {})

	// Act
	actual := args.Map{"result": called}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be called on Uninitialized", actual)

	called = v.LazyEvaluateSet(func() {})
	actual = args.Map{"result": called}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be called again on Set", actual)
}

func Test_Value_GetSetBoolOnInvalid_Cov(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	result := v.GetSetBoolOnInvalid(true)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should return true after setting", actual)
}

func Test_Value_GetSetBoolOnInvalidFunc(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	result := v.GetSetBoolOnInvalidFunc(func() bool { return false })

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return false after setting", actual)
}

func Test_Value_IsTrueOrSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsTrueOrSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsTrueOrSet should be true", actual)
	actual = args.Map{"result": issetter.Set.IsTrueOrSet()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set.IsTrueOrSet should be true", actual)
	actual = args.Map{"result": issetter.False.IsTrueOrSet()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False.IsTrueOrSet should be false", actual)
}

func Test_Value_HasInitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.HasInitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be initialized", actual)
	actual = args.Map{"result": issetter.Uninitialized.HasInitialized()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should not be initialized", actual)
}

func Test_Value_HasInitializedAndSet(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.HasInitializedAndSet()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set should be initialized and set", actual)
}

func Test_Value_HasInitializedAndTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.HasInitializedAndTrue()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be initialized and true", actual)
}

func Test_Value_IsOnOffLogically(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsOnLogically()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be on logically", actual)
	actual = args.Map{"result": issetter.False.IsOffLogically()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be off logically", actual)
}

func Test_Value_IsAcceptedRejected(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsAccepted()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be accepted", actual)
	actual = args.Map{"result": issetter.False.IsRejected()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be rejected", actual)
}

func Test_Value_IsUnSetOrUninitialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Uninitialized.IsUnSetOrUninitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be unset or uninitialized", actual)
	actual = args.Map{"result": issetter.Unset.IsUnSetOrUninitialized()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Unset should be unset or uninitialized", actual)
	actual = args.Map{"result": issetter.True.IsUnSetOrUninitialized()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should not be unset or uninitialized", actual)
}

func Test_Value_WildcardValueApply(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.WildcardValueApply(issetter.True)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard should pass through input", actual)
}

func Test_Value_OrValue(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.OrValue(issetter.False)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.OrValue(False) should be true", actual)
}

func Test_Value_IsWildcardOrBool_Cov(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Wildcard.IsWildcardOrBool(false)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard should always return true", actual)
}

func Test_Value_IsDefinedBoolean(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsDefinedBoolean()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be defined boolean", actual)
	actual = args.Map{"result": issetter.Set.IsDefinedBoolean()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Set should not be defined boolean", actual)
}

func Test_Value_IsValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsValueEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsValueEqual(1) should be true", actual)
}

func Test_Value_IsByteValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsByteValueEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsByteValueEqual(1) should be true", actual)
}

func Test_Value_OnlySupportedErr(t *testing.T) {
	// Arrange
	names := []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}
	err := issetter.True.OnlySupportedErr(names...)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all supported should return nil, got:", actual)
}

func Test_Value_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("test: ", "NonExistent")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unsupported should return error", actual)
}

func Test_Value_MinMaxInt(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.MaxInt() < issetter.True.MinInt()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxInt should be >= MinInt", actual)
}

func Test_Value_Initialized(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsInitialized()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be initialized", actual)
	actual = args.Map{"result": issetter.Uninitialized.IsInitialized()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should not be initialized", actual)
}
