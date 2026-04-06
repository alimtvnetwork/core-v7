package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
)

// TestValue_AllNameValues verifies AllNameValues returns all enum names.
func TestValue_AllNameValues(t *testing.T) {
	// Arrange
	val := issetter.Uninitialized

	// Act
	result := val.AllNameValues()

	// Assert
	actual := args.Map{"result": len(result) != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6 name values", actual)
}

// TestValue_OnlySupportedErr verifies unsupported name detection.
func TestValue_OnlySupportedErr(t *testing.T) {
	for _, tc := range onlySupportedErrCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			val := issetter.Uninitialized

			// Act
			err := val.OnlySupportedErr(tc.names...)

			// Assert
			actual := args.Map{"result": tc.expectErr && err == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected error but got nil", actual)
			actual := args.Map{"result": tc.expectErr && err != nil}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected nil but got error:", actual)
		})
	}
}

// TestValue_OnlySupportedMsgErr verifies message-prefixed error.
func TestValue_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	val := issetter.Uninitialized

	// Act
	err := val.OnlySupportedMsgErr("prefix: ", "True", "False", "Uninitialized", "Set", "Unset", "Wildcard")

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil but got error:", actual)
}

// TestValue_IntegerEnumRanges verifies integer ranges.
func TestValue_IntegerEnumRanges(t *testing.T) {
	// Arrange
	val := issetter.Uninitialized

	// Act
	ranges := val.IntegerEnumRanges()

	// Assert
	actual := args.Map{"result": len(ranges) != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6 ranges", actual)
}

// TestValue_MinMaxAny verifies min/max.
func TestValue_MinMaxAny(t *testing.T) {
	// Arrange
	val := issetter.Uninitialized

	// Act
	minVal, maxVal := val.MinMaxAny()

	// Assert
	actual := args.Map{"result": minVal != issetter.Uninitialized}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Uninitialized min", actual)
	actual := args.Map{"result": maxVal != issetter.Wildcard}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Wildcard max", actual)
}

// TestValue_Format verifies format string replacement.
func TestValue_Format(t *testing.T) {
	// Arrange
	val := issetter.True

	// Act
	result := val.Format("{name}={value}")

	// Assert
	actual := args.Map{"result": result != "True=1"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'True=1', got ''", actual)
}

// TestValue_Conversions verifies value type conversions.
func TestValue_Conversions(t *testing.T) {
	for _, tc := range conversionCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange & Act & Assert
			actual := args.Map{"result": tc.val.ValueByte() != tc.expectedByte}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "ValueByte: expected", actual)
			actual := args.Map{"result": tc.val.ValueInt() != tc.expectedInt}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "ValueInt: expected", actual)
		})
	}
}

// TestValue_LogicalChecks verifies logical boolean checks.
func TestValue_LogicalChecks(t *testing.T) {
	for _, tc := range logicalCheckCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange & Act & Assert
			actual := args.Map{"result": tc.val.IsOn() != tc.isOn}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "IsOn: expected", actual)
			actual := args.Map{"result": tc.val.IsOff() != tc.isOff}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "IsOff: expected", actual)
			actual := args.Map{"result": tc.val.IsAsk() != tc.isAsk}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "IsAsk: expected", actual)
			actual := args.Map{"result": tc.val.IsAccept() != tc.isAccept}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "IsAccept: expected", actual)
			actual := args.Map{"result": tc.val.IsReject() != tc.isReject}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "IsReject: expected", actual)
		})
	}
}

// TestValue_Names verifies name variants.
func TestValue_Names(t *testing.T) {
	for _, tc := range nameCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act & Assert
			actual := args.Map{"result": tc.val.YesNoName() != tc.yesNo}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "YesNoName: expected '', got ''", actual)
			actual := args.Map{"result": tc.val.OnOffName() != tc.onOff}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "OnOffName: expected '', got ''", actual)
			actual := args.Map{"result": tc.val.TrueFalseName() != tc.trueFalse}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "TrueFalseName: expected '', got ''", actual)
		})
	}
}

// TestValue_MarshalUnmarshalJSON verifies JSON round-trip.
func TestValue_MarshalUnmarshalJSON(t *testing.T) {
	for _, tc := range jsonCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			bytes, err := tc.val.MarshalJSON()

			// Assert
			actual := args.Map{"result": err != nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)

			var result issetter.Value
			err = result.UnmarshalJSON(bytes)
			actual := args.Map{"result": err != nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
			actual := args.Map{"result": result != tc.val}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "round-trip: expected", actual)
		})
	}
}

// TestValue_UnmarshalJSON_Invalid verifies error on invalid JSON.
func TestValue_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var v issetter.Value

	// Act
	err := v.UnmarshalJSON([]byte("invalid"))

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid JSON", actual)
}

// TestValue_UnmarshalJSON_Nil verifies error on nil input.
func TestValue_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	var v issetter.Value

	// Act
	err := v.UnmarshalJSON(nil)

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil data", actual)
}

// TestValue_ToBooleanValue verifies conversion from Set/Unset to True/False.
func TestValue_ToBooleanValue(t *testing.T) {
	for _, tc := range toBooleanCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := tc.input.ToBooleanValue()

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestValue_ToSetUnsetValue verifies conversion from True/False to Set/Unset.
func TestValue_ToSetUnsetValue(t *testing.T) {
	for _, tc := range toSetUnsetCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := tc.input.ToSetUnsetValue()

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestValue_WildcardApply verifies wildcard application logic.
func TestValue_WildcardApply(t *testing.T) {
	for _, tc := range wildcardApplyCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := tc.val.WildcardApply(tc.input)

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestValue_OrBool verifies OrBool logic.
func TestValue_OrBool(t *testing.T) {
	// Arrange & Act & Assert
	actual := args.Map{"result": issetter.True.OrBool(false) != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.OrBool(false) should be true", actual)
	actual := args.Map{"result": issetter.False.OrBool(true) != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False.OrBool(true) should be true", actual)
	actual := args.Map{"result": issetter.Wildcard.OrBool(true) != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard.OrBool(true) should be true", actual)
	actual := args.Map{"result": issetter.Wildcard.OrBool(false) != false}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard.OrBool(false) should be false", actual)
}

// TestValue_AndBool verifies AndBool logic.
func TestValue_AndBool(t *testing.T) {
	actual := args.Map{"result": issetter.True.AndBool(true) != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.AndBool(true) should be true", actual)
	actual := args.Map{"result": issetter.True.AndBool(false) != false}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.AndBool(false) should be false", actual)
	actual := args.Map{"result": issetter.Wildcard.AndBool(true) != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard.AndBool(true) should be true", actual)
}

// TestValue_And verifies And logic.
func TestValue_And(t *testing.T) {
	result := issetter.True.And(issetter.True)
	actual := args.Map{"result": result != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.And(True) should be True", actual)
	result = issetter.True.And(issetter.False)
	actual := args.Map{"result": result != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.And(False) should be False", actual)
	result = issetter.Wildcard.And(issetter.True)
	actual := args.Map{"result": result != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard.And(True) should be True", actual)
}

// TestValue_IsCompareResult verifies comparison operations.
func TestValue_IsCompareResult(t *testing.T) {
	for _, tc := range compareResultCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := tc.val.IsCompareResult(tc.n, tc.compare)

			// Assert
			actual := args.Map{"result": result != tc.expected}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestValue_IsBetween verifies range check.
func TestValue_IsBetween(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsBetween(0, 5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be between 0 and 5", actual)
	actual := args.Map{"result": issetter.True.IsBetween(2, 5)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True(1) should not be between 2 and 5", actual)
}

// TestValue_IsBetweenInt verifies int range check.
func TestValue_IsBetweenInt(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsBetweenInt(0, 5)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be between 0 and 5", actual)
}

// TestValue_Add verifies arithmetic Add.
func TestValue_Add(t *testing.T) {
	result := issetter.True.Add(1)
	actual := args.Map{"result": result != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.Add(1) should be False(2)", actual)
}

// TestValue_GetSetBoolOnInvalid verifies lazy boolean getter/setter.
func TestValue_GetSetBoolOnInvalid(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized

	// Act
	result := v.GetSetBoolOnInvalid(true)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": v != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True", actual)
}

// TestValue_GetSetBoolOnInvalidFunc verifies lazy func-based boolean getter/setter.
func TestValue_GetSetBoolOnInvalidFunc(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized

	// Act
	result := v.GetSetBoolOnInvalidFunc(func() bool { return false })

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

// TestValue_LazyEvaluateBool verifies lazy evaluate.
func TestValue_LazyEvaluateBool(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := false

	// Act
	isCalled := v.LazyEvaluateBool(func() { called = true })

	// Assert
	actual := args.Map{"result": isCalled || !called}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected evaluator to be called", actual)
}

// TestValue_LazyEvaluateSet verifies lazy set evaluate.
func TestValue_LazyEvaluateSet(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := false

	// Act
	isCalled := v.LazyEvaluateSet(func() { called = true })

	// Assert
	actual := args.Map{"result": isCalled || !called}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected evaluator to be called", actual)
}

// TestValue_ToByteCondition verifies byte condition mapping.
func TestValue_ToByteCondition(t *testing.T) {
	actual := args.Map{"result": issetter.True.ToByteCondition(10, 20, 30) != 10}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should return trueVal", actual)
	actual := args.Map{"result": issetter.False.ToByteCondition(10, 20, 30) != 20}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False should return falseVal", actual)
	actual := args.Map{"result": issetter.Uninitialized.ToByteCondition(10, 20, 30) != 30}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should return invalid", actual)
}

// TestValue_ToByteConditionWithWildcard verifies wildcard byte condition.
func TestValue_ToByteConditionWithWildcard(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.ToByteConditionWithWildcard(5, 10, 20, 30) != 5}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard should return wildcard val", actual)
}

// TestValue_Deserialize verifies Deserialize round-trip.
func TestValue_Deserialize(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized

	// Act
	result, err := v.Deserialize([]byte(`"True"`))

	// Assert
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": result != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True", actual)
}

// TestValue_Deserialize_Invalid verifies Deserialize error.
func TestValue_Deserialize_Invalid(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized

	// Act
	_, err := v.Deserialize([]byte("garbage"))

	// Assert
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// TestValue_TypeName verifies TypeName.
func TestValue_TypeName(t *testing.T) {
	actual := args.Map{"result": issetter.True.TypeName() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeName should not be empty", actual)
}

// TestValue_RangeNamesCsv verifies CSV ranges.
func TestValue_RangeNamesCsv(t *testing.T) {
	actual := args.Map{"result": issetter.True.RangeNamesCsv() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
}

// TestValue_MaxByte_MinByte verifies max/min byte.
func TestValue_MaxByte_MinByte(t *testing.T) {
	actual := args.Map{"result": issetter.True.MaxByte() != issetter.Wildcard.ValueByte()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxByte mismatch", actual)
	actual := args.Map{"result": issetter.True.MinByte() != issetter.Uninitialized.ValueByte()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinByte mismatch", actual)
}

// TestValue_ToPtr verifies pointer conversion.
func TestValue_ToPtr(t *testing.T) {
	v := issetter.True
	ptr := v.ToPtr()
	actual := args.Map{"result": *ptr != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToPtr value mismatch", actual)
}

// TestValue_IsAnyValuesEqual verifies multi-value comparison.
func TestValue_IsAnyValuesEqual(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyValuesEqual(0, 1, 2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) should match 1 in list", actual)
	actual := args.Map{"result": issetter.True.IsAnyValuesEqual(0, 2, 3)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True(1) should not match 0,2,3", actual)
}

// TestValue_IsAnyNamesOf verifies name matching.
func TestValue_IsAnyNamesOf(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("True", "False")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should match 'True'", actual)
	actual := args.Map{"result": issetter.True.IsAnyNamesOf("False", "Set")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should not match 'False','Set'", actual)
}

// TestGetBool verifies GetBool helper.
func TestGetBool(t *testing.T) {
	actual := args.Map{"result": issetter.GetBool(true) != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetBool(true) should be True", actual)
	actual := args.Map{"result": issetter.GetBool(false) != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetBool(false) should be False", actual)
}

// TestGetSet verifies GetSet helper.
func TestGetSet(t *testing.T) {
	actual := args.Map{"result": issetter.GetSet(true, issetter.Set, issetter.Unset) != issetter.Set}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSet(true) should return trueValue", actual)
	actual := args.Map{"result": issetter.GetSet(false, issetter.Set, issetter.Unset) != issetter.Unset}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSet(false) should return falseValue", actual)
}

// TestGetSetByte verifies GetSetByte helper.
func TestGetSetByte(t *testing.T) {
	r := issetter.GetSetByte(true, 1, 2)
	actual := args.Map{"result": r != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True(1)", actual)
}

// TestGetSetUnset verifies GetSetUnset helper.
func TestGetSetUnset(t *testing.T) {
	actual := args.Map{"result": issetter.GetSetUnset(true) != issetter.Set}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSetUnset(true) should be Set", actual)
	actual := args.Map{"result": issetter.GetSetUnset(false) != issetter.Unset}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "GetSetUnset(false) should be Unset", actual)
}

// TestNewBool verifies NewBool helper.
func TestNewBool(t *testing.T) {
	actual := args.Map{"result": issetter.NewBool(true) != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NewBool(true) should be True", actual)
}

// TestCombinedBooleans verifies CombinedBooleans helper.
func TestCombinedBooleans(t *testing.T) {
	actual := args.Map{"result": issetter.CombinedBooleans(true, true) != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all true should be True", actual)
	actual := args.Map{"result": issetter.CombinedBooleans(true, false) != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "any false should be False", actual)
}

// TestNew verifies New from string.
func TestNew(t *testing.T) {
	for _, tc := range newFromStringCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			val, err := issetter.New(tc.input)

			// Assert
			actual := args.Map{"result": tc.expectErr && err == nil}
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "expected error", actual)
			actual := args.Map{"result": tc.expectErr && err != nil}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
			actual := args.Map{"result": tc.expectErr && val != tc.expected}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected", actual)
		})
	}
}

// TestGetSetterByComparing verifies GetSetterByComparing helper.
func TestGetSetterByComparing(t *testing.T) {
	// Act
	r := issetter.GetSetterByComparing(issetter.True, issetter.False, 5, 1, 3, 5)

	// Assert
	actual := args.Map{"result": r != issetter.True}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected True when value matches range", actual)

	r = issetter.GetSetterByComparing(issetter.True, issetter.False, 7, 1, 3, 5)
	actual := args.Map{"result": r != issetter.False}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected False when value not in range", actual)
}

// TestIsOutOfRange verifies IsOutOfRange.
func TestIsOutOfRange(t *testing.T) {
	actual := args.Map{"result": issetter.IsOutOfRange(1)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1 should not be out of range", actual)
}

// TestValue_YesNoMappedValue verifies YesNoMappedValue.
func TestValue_YesNoMappedValue(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.YesNoMappedValue() != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should return empty", actual)
	actual := args.Map{"result": issetter.True.YesNoMappedValue() != "yes"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True should return 'yes', got ''", actual)
	actual := args.Map{"result": issetter.False.YesNoMappedValue() != "no"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "False should return 'no', got ''", actual)
}

// TestValue_GetErrorOnOutOfRange verifies error on out of range.
func TestValue_GetErrorOnOutOfRange(t *testing.T) {
	// In range
	err := issetter.True.GetErrorOnOutOfRange(1, "out of range")
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "1 should not be out of range", actual)
}

// TestValue_RangesDynamicMap verifies dynamic ranges map.
func TestValue_RangesDynamicMap(t *testing.T) {
	m := issetter.True.RangesDynamicMap()
	actual := args.Map{"result": len(m) != 6}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 6 entries", actual)
}

// TestValue_MinMaxValueString verifies min/max value strings.
func TestValue_MinMaxValueString(t *testing.T) {
	v := issetter.Uninitialized
	actual := args.Map{"result": v.MinValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)
	actual := args.Map{"result": v.MaxValueString() == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)
}

// TestValue_MaxMinInt verifies MaxInt/MinInt.
func TestValue_MaxMinInt(t *testing.T) {
	v := issetter.Uninitialized
	actual := args.Map{"result": v.MaxInt() != issetter.Wildcard.ValueInt()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxInt mismatch", actual)
	actual := args.Map{"result": v.MinInt() != issetter.Uninitialized.ValueInt()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinInt mismatch", actual)
}

// TestValue_EnumType verifies EnumType.
func TestValue_EnumType(t *testing.T) {
	actual := args.Map{"result": issetter.True.EnumType() == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EnumType should not be nil", actual)
}

// TestValue_Serialize verifies Serialize.
func TestValue_Serialize(t *testing.T) {
	b, err := issetter.True.Serialize()
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": len(b) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Serialize should return bytes", actual)
}

// TestValue_WildcardOrBool verifies IsWildcardOrBool.
func TestValue_WildcardOrBool(t *testing.T) {
	actual := args.Map{"result": issetter.Wildcard.IsWildcardOrBool(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard.IsWildcardOrBool should be true regardless", actual)
	actual := args.Map{"result": issetter.False.IsWildcardOrBool(false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False.IsWildcardOrBool(false) should be true per formula (!isBool && IsFalse())", actual)
}

// TestValue_OrValue verifies OrValue.
func TestValue_OrValue(t *testing.T) {
	actual := args.Map{"result": issetter.True.OrValue(issetter.False)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.OrValue(False) should be true", actual)
	actual := args.Map{"result": issetter.Wildcard.OrValue(issetter.True)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard.OrValue(True) should be true", actual)
}

// TestValue_WildcardValueApply verifies WildcardValueApply.
func TestValue_WildcardValueApply(t *testing.T) {
	actual := args.Map{"result": issetter.True.WildcardValueApply(issetter.False) != true}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True.WildcardValueApply should return True", actual)
	actual := args.Map{"result": issetter.Wildcard.WildcardValueApply(issetter.False) != false}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard should delegate to input", actual)
}

// TestValue_IsNot verifies IsNot.
func TestValue_IsNot(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsNot(False) should be true", actual)
}

// TestValue_Negative_Positive verifies IsNegative/IsPositive.
func TestValue_Negative_Positive(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.IsNegative()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Uninitialized should be negative", actual)
	actual := args.Map{"result": issetter.True.IsPositive()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be positive", actual)
}

// Test remaining helper functions for completeness
func TestValue_ComparisonHelpers(t *testing.T) {
	v := issetter.True
	actual := args.Map{"result": v.IsGreater(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) > 0", actual)
	actual := args.Map{"result": v.IsGreaterEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) >= 1", actual)
	actual := args.Map{"result": v.IsLess(2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) < 2", actual)
	actual := args.Map{"result": v.IsLessEqual(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) <= 1", actual)
	actual := args.Map{"result": v.IsGreaterInt(0)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) > int(0)", actual)
	actual := args.Map{"result": v.IsGreaterEqualInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) >= int(1)", actual)
	actual := args.Map{"result": v.IsLessInt(2)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) < int(2)", actual)
	actual := args.Map{"result": v.IsLessEqualInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) <= int(1)", actual)
	actual := args.Map{"result": v.IsEqualInt(1)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True(1) == int(1)", actual)
}

// TestValue_InitChecks verifies Init/InitBoolean/InitSet/InitSetWild checks.
func TestValue_InitChecks(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsInitBoolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be InitBoolean", actual)
	actual := args.Map{"result": issetter.False.IsInitBoolean()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False should be InitBoolean", actual)
	actual := args.Map{"result": issetter.Set.IsInitBoolean()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Set should not be InitBoolean", actual)
	actual := args.Map{"result": issetter.Set.IsInitSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set should be InitSet", actual)
	actual := args.Map{"result": issetter.Wildcard.IsInitBooleanWild()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard should be InitBooleanWild", actual)
	actual := args.Map{"result": issetter.Wildcard.IsInitSetWild()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Wildcard should be InitSetWild", actual)
}

// TestValue_LowercaseNames verifies lowercase name variants.
func TestValue_LowercaseNames(t *testing.T) {
	actual := args.Map{"result": issetter.True.YesNoLowercaseName() != "yes"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True YesNoLowercaseName should be 'yes'", actual)
	actual := args.Map{"result": issetter.True.OnOffLowercaseName() != "on"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True OnOffLowercaseName should be 'on'", actual)
	actual := args.Map{"result": issetter.True.TrueFalseLowercaseName() != "true"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True TrueFalseLowercaseName should be 'true'", actual)
	actual := args.Map{"result": issetter.True.SetUnsetLowercaseName() != "set"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "True SetUnsetLowercaseName should be 'set'", actual)
}

// TestValue_IsOnLogically verifies logical on/off.
func TestValue_IsOnLogically(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsOnLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.IsOnLogically should be true", actual)
	actual := args.Map{"result": issetter.False.IsOffLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "False.IsOffLogically should be true", actual)
	actual := args.Map{"result": issetter.Uninitialized.IsOnLogically()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized should not be on logically", actual)
}

// TestValue_IsDefinedLogically verifies defined/undefined.
func TestValue_IsDefinedLogically(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsDefinedLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True should be defined logically", actual)
	actual := args.Map{"result": issetter.Wildcard.IsDefinedLogically()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Wildcard should be undefined logically", actual)
}

// Test remaining statefulness
func TestValue_HasInitialized(t *testing.T) {
	actual := args.Map{"result": issetter.Uninitialized.HasInitialized()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Uninitialized.HasInitialized should be false", actual)
	actual := args.Map{"result": issetter.True.HasInitialized()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.HasInitialized should be true", actual)
	actual := args.Map{"result": issetter.True.HasInitializedAndTrue()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "True.HasInitializedAndTrue should be true", actual)
	actual := args.Map{"result": issetter.Set.HasInitializedAndTrue()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Set.HasInitializedAndTrue should be false", actual)
	actual := args.Map{"result": issetter.Set.HasInitializedAndSet()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Set.HasInitializedAndSet should be true", actual)
}

// TestValue_UnmarshallEnumToValue verifies enum unmarshal.
func TestValue_UnmarshallEnumToValue(t *testing.T) {
	v := issetter.Uninitialized
	b, err := v.UnmarshallEnumToValue([]byte(`"True"`))
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
	actual := args.Map{"result": b != 1}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected byte 1", actual)
}

// Unused but important coverage
func TestValue_Misc(t *testing.T) {
	_ = issetter.True.NameValue()
	_ = issetter.True.ToNumberString()
	_ = issetter.True.ValueString()
	_ = issetter.True.StringValue()
	_ = issetter.True.String()
	_ = issetter.True.IsValueEqual(1)
	_ = issetter.True.IsByteValueEqual(1)
	_ = issetter.True.Is(issetter.True)
	_ = issetter.True.IsEqual(1)
	_ = issetter.True.Boolean()
	_ = issetter.True.IsYes()
	_ = issetter.True.IsLater()
	_ = issetter.True.IsNo()
	_ = issetter.True.IsFailed()
	_ = issetter.True.IsSuccess()
	_ = issetter.True.IsSkip()
	_ = issetter.True.IsIndeterminate()
	_ = issetter.True.IsAccepted()
	_ = issetter.True.IsRejected()
	_ = issetter.True.ValueUInt16()
	_ = issetter.True.ValueInt8()
	_ = issetter.True.ValueInt16()
	_ = issetter.True.ValueInt32()
	_ = issetter.True.IsNameEqual("True")
	_ = issetter.True.IsUninitialized()
	_ = issetter.True.IsInitialized()
	_ = issetter.True.IsUnSetOrUninitialized()
	_ = issetter.True.IsValid()
	_ = issetter.True.IsInvalid()
	_ = issetter.True.IsWildcard()
	_ = issetter.True.IsInit()
	_ = issetter.True.IsDefinedBoolean()
	_ = issetter.True.IsTrue()
	_ = issetter.True.IsFalse()
	_ = issetter.True.IsTrueOrSet()
	_ = issetter.True.IsSet()
	_ = issetter.True.IsUnset()
	_ = issetter.Max()
	_ = issetter.Min()
	_ = issetter.MaxByte()
	_ = issetter.MinByte()
	_ = issetter.RangeNamesCsv()
	_ = issetter.IntegerEnumRanges()
}
