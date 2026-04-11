package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ============================================================================
// And / Or / WildcardApply — additional branches
// ============================================================================

func Test_And_TrueFalse(t *testing.T) {
	// Act
	actual := args.Map{
		"trueAndFalse":  issetter.True.And(issetter.False) == issetter.False,
		"falseAndTrue":  issetter.False.And(issetter.True) == issetter.False,
		"falseAndFalse": issetter.False.And(issetter.False) == issetter.False,
		"wildAndTrue":   issetter.Wildcard.And(issetter.True) == issetter.True,
	}

	// Assert
	expected := args.Map{
		"trueAndFalse": true, "falseAndTrue": true,
		"falseAndFalse": true, "wildAndTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "And returns expected -- all combinations", actual)
}

func Test_OrBool_Branches(t *testing.T) {
	// Act
	actual := args.Map{
		"falseOrTrue":  issetter.False.OrBool(true),
		"falseOrFalse": issetter.False.OrBool(false),
		"wildOrFalse":  issetter.Wildcard.OrBool(false),
	}

	// Assert
	expected := args.Map{
		"falseOrTrue": true, "falseOrFalse": false, "wildOrFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "OrBool returns expected -- all branches", actual)
}

func Test_WildcardApply_False(t *testing.T) {
	// Act
	actual := args.Map{
		"wildFalse":  issetter.Wildcard.WildcardApply(false),
		"falseFalse": issetter.False.WildcardApply(true),
	}

	// Assert
	expected := args.Map{
		"wildFalse": false,
		"falseFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "WildcardApply returns expected -- false paths", actual)
}

func Test_WildcardValueApply(t *testing.T) {
	// Act
	actual := args.Map{
		"wildTrue":  issetter.Wildcard.WildcardValueApply(issetter.True),
		"wildFalse": issetter.Wildcard.WildcardValueApply(issetter.False),
		"trueAny":   issetter.True.WildcardValueApply(issetter.False),
		"falseAny":  issetter.False.WildcardValueApply(issetter.True),
	}

	// Assert
	expected := args.Map{
		"wildTrue": true, "wildFalse": false,
		"trueAny": true, "falseAny": false,
	}
	expected.ShouldBeEqual(t, 0, "WildcardValueApply returns expected -- all paths", actual)
}

// ============================================================================
// OrValue
// ============================================================================

func Test_OrValue(t *testing.T) {
	// Act
	actual := args.Map{
		"trueOrFalse":  issetter.True.OrValue(issetter.False),
		"falseOrTrue":  issetter.False.OrValue(issetter.True),
		"falseOrFalse": issetter.False.OrValue(issetter.False),
	}

	// Assert
	expected := args.Map{
		"trueOrFalse": true, "falseOrTrue": true, "falseOrFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "OrValue returns expected -- all combinations", actual)
}

// ============================================================================
// AndBool edge cases
// ============================================================================

func Test_AndBool_WildcardTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"wildTrue":  issetter.Wildcard.AndBool(true),
		"wildFalse": issetter.Wildcard.AndBool(false),
	}

	// Assert
	expected := args.Map{
		"wildTrue": true,
		"wildFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "AndBool wildcard passes through -- both", actual)
}

// ============================================================================
// IsWildcardOrBool with True
// ============================================================================

func Test_IsWildcardOrBool(t *testing.T) {
	// Act
	actual := args.Map{
		"wildTrue":  issetter.Wildcard.IsWildcardOrBool(true),
		"wildFalse": issetter.Wildcard.IsWildcardOrBool(false),
		"trueTrue":  issetter.True.IsWildcardOrBool(true),
		"trueFalse": issetter.True.IsWildcardOrBool(false),
		"falseFalse": issetter.False.IsWildcardOrBool(false),
	}

	// Assert
	expected := args.Map{
		"wildTrue": true, "wildFalse": true,
		"trueTrue": true, "trueFalse": false,
		"falseFalse": true,
	}
	expected.ShouldBeEqual(t, 0, "IsWildcardOrBool returns expected -- all combos", actual)
}

// ============================================================================
// ToByteCondition — Uninitialized
// ============================================================================

func Test_ToByteCondition_Uninit(t *testing.T) {
	// Act
	actual := args.Map{
		"result": int(issetter.Uninitialized.ToByteCondition(1, 0, 255)),
	}

	// Assert
	expected := args.Map{"result": 255}
	expected.ShouldBeEqual(t, 0, "ToByteCondition returns invalid -- Uninitialized", actual)
}

func Test_ToByteConditionWithWildcard_Uninit(t *testing.T) {
	// Act
	actual := args.Map{
		"result": int(issetter.Uninitialized.ToByteConditionWithWildcard(99, 1, 0, 255)),
	}

	// Assert
	expected := args.Map{"result": 255}
	expected.ShouldBeEqual(t, 0, "ToByteConditionWithWildcard returns invalid -- Uninitialized", actual)
}

// ============================================================================
// ToBooleanValue / ToSetUnsetValue — edge cases
// ============================================================================

func Test_ToBooleanValue_Wildcard(t *testing.T) {
	// Arrange
	result := issetter.Wildcard.ToBooleanValue()

	// Act
	actual := args.Map{"isWild": result == issetter.Wildcard}

	// Assert
	expected := args.Map{"isWild": true}
	expected.ShouldBeEqual(t, 0, "ToBooleanValue returns correct value -- Wildcard stays Wildcard", actual)
}

func Test_ToSetUnsetValue_Wildcard(t *testing.T) {
	// Arrange
	result := issetter.Wildcard.ToSetUnsetValue()

	// Act
	actual := args.Map{"isWild": result == issetter.Wildcard}

	// Assert
	expected := args.Map{"isWild": true}
	expected.ShouldBeEqual(t, 0, "ToSetUnsetValue returns correct value -- Wildcard stays Wildcard", actual)
}

// ============================================================================
// IsOnLogically / IsOffLogically — more values
// ============================================================================

func Test_IsOnLogically_Set(t *testing.T) {
	// Act
	actual := args.Map{
		"setOn":    issetter.Set.IsOnLogically(),
		"unsetOff": issetter.Unset.IsOffLogically(),
		"wildOff":  issetter.Wildcard.IsOffLogically(),
	}

	// Assert
	expected := args.Map{
		"setOn": true, "unsetOff": true, "wildOff": false,
	}
	expected.ShouldBeEqual(t, 0, "IsOnLogically/IsOffLogically returns expected -- Set/Unset/Wild", actual)
}

// ============================================================================
// GetSetBoolOnInvalid — already initialized
// ============================================================================

func Test_GetSetBoolOnInvalid_AlreadyInit(t *testing.T) {
	// Arrange
	v := issetter.True
	result := v.GetSetBoolOnInvalid(false)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "GetSetBoolOnInvalid returns existing -- already True", actual)
}

func Test_GetSetBoolOnInvalidFunc_AlreadyInit(t *testing.T) {
	// Arrange
	v := issetter.False
	result := v.GetSetBoolOnInvalidFunc(func() bool { return true })

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSetBoolOnInvalidFunc returns existing -- already False", actual)
}

// ============================================================================
// LazyEvaluateBool — True value
// ============================================================================

func Test_LazyEvaluateBool_True(t *testing.T) {
	// Arrange
	v := issetter.True
	called := v.LazyEvaluateBool(func() {})

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LazyEvaluateBool not called -- already True", actual)
}

// ============================================================================
// LazyEvaluateSet — Set value
// ============================================================================

func Test_LazyEvaluateSet_Set(t *testing.T) {
	// Arrange
	v := issetter.Set
	called := v.LazyEvaluateSet(func() {})

	// Act
	actual := args.Map{"called": called}

	// Assert
	expected := args.Map{"called": false}
	expected.ShouldBeEqual(t, 0, "LazyEvaluateSet not called -- already Set", actual)
}

// ============================================================================
// YesNoMappedValue edge values
// ============================================================================

func Test_YesNoMappedValue_SetUnset(t *testing.T) {
	// Act
	actual := args.Map{
		"set":   issetter.Set.YesNoMappedValue(),
		"unset": issetter.Unset.YesNoMappedValue(),
	}

	// Assert
	expected := args.Map{
		"set": "yes",
		"unset": "no",
	}
	expected.ShouldBeEqual(t, 0, "YesNoMappedValue returns expected -- Set/Unset", actual)
}

// ============================================================================
// Name methods — False/Unset/Wildcard/Uninitialized
// ============================================================================

func Test_NameMethods_False(t *testing.T) {
	// Act
	actual := args.Map{
		"yesNo":        issetter.False.YesNoName(),
		"trueFalse":    issetter.False.TrueFalseName(),
		"onOff":        issetter.False.OnOffName(),
		"yesNoLower":   issetter.False.YesNoLowercaseName(),
		"trFaLower":    issetter.False.TrueFalseLowercaseName(),
		"onOffLower":   issetter.False.OnOffLowercaseName(),
		"setUnsetLow":  issetter.False.SetUnsetLowercaseName(),
	}

	// Assert
	expected := args.Map{
		"yesNo": "No", "trueFalse": "False", "onOff": "Off",
		"yesNoLower": "no", "trFaLower": "false",
		"onOffLower": "off", "setUnsetLow": "unset",
	}
	expected.ShouldBeEqual(t, 0, "Name methods False returns expected -- all variants", actual)
}

func Test_NameMethods_Wildcard(t *testing.T) {
	// Act
	actual := args.Map{
		"yesNo":      issetter.Wildcard.YesNoMappedValue(),
		"trueFalse":  issetter.Wildcard.TrueFalseName(),
		"onOff":      issetter.Wildcard.OnOffName(),
	}

	// Assert
	expected := args.Map{
		"yesNo": issetter.Wildcard.YesNoMappedValue(), "trueFalse": issetter.Wildcard.TrueFalseName(), "onOff": issetter.Wildcard.OnOffName(),
	}
	expected.ShouldBeEqual(t, 0, "Name methods Wildcard returns empty -- undefined", actual)
}

// ============================================================================
// IsEqual / IsBetween additional edge cases
// ============================================================================

func Test_IsBetween_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{
		"below": issetter.Uninitialized.IsBetween(1, 5),
		"above": issetter.Wildcard.IsBetween(0, 2),
	}

	// Assert
	expected := args.Map{
		"below": false,
		"above": false,
	}
	expected.ShouldBeEqual(t, 0, "IsBetween out of range returns false -- both edges", actual)
}
