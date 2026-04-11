package issettertests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

// ============================================================================
// Value — comprehensive method coverage
// ============================================================================

func Test_Value_BooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"trueIsOn":      issetter.True.IsOn(),
		"falseIsOff":    issetter.False.IsOff(),
		"trueIsAccept":  issetter.True.IsAccept(),
		"falseIsReject": issetter.False.IsReject(),
		"trueIsSuccess": issetter.True.IsSuccess(),
		"falseIsFailed": issetter.False.IsFailed(),
		"trueIsYes":     issetter.True.IsYes(),
		"trueBoolean":   issetter.True.Boolean(),
		"falseIsNo":     issetter.False.IsNo(),
		"uninitIsAsk":   issetter.Uninitialized.IsAsk(),
		"wildIsSkip":    issetter.Wildcard.IsSkip(),
		"uninitIsLater": issetter.Uninitialized.IsLater(),
	}

	// Assert
	expected := args.Map{
		"trueIsOn": true, "falseIsOff": true,
		"trueIsAccept": true, "falseIsReject": true,
		"trueIsSuccess": true, "falseIsFailed": true,
		"trueIsYes": true, "trueBoolean": true,
		"falseIsNo": true, "uninitIsAsk": true,
		"wildIsSkip": true, "uninitIsLater": true,
	}
	expected.ShouldBeEqual(t, 0, "Value boolean checks -- all", actual)
}

func Test_Value_StateChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"trueHasInit":      issetter.True.HasInitialized(),
		"trueHasInitTrue":  issetter.True.HasInitializedAndTrue(),
		"setHasInitSet":    issetter.Set.HasInitializedAndSet(),
		"trueIsInit":       issetter.True.IsInit(),
		"trueIsInitBool":   issetter.True.IsInitBoolean(),
		"trueIsDefBool":    issetter.True.IsDefinedBoolean(),
		"trueIsInitBoolW":  issetter.True.IsInitBooleanWild(),
		"wildIsInitBoolW":  issetter.Wildcard.IsInitBooleanWild(),
		"setIsInitSet":     issetter.Set.IsInitSet(),
		"wildIsInitSetW":   issetter.Wildcard.IsInitSetWild(),
		"trueIsTrueOrSet":  issetter.True.IsTrueOrSet(),
		"setIsTrueOrSet":   issetter.Set.IsTrueOrSet(),
		"uninitIsNeg":      issetter.Uninitialized.IsNegative(),
		"trueIsPos":        issetter.True.IsPositive(),
		"setIsPos":         issetter.Set.IsPositive(),
		"trueIsDefLogic":   issetter.True.IsDefinedLogically(),
		"uninitIsUndefLog": issetter.Uninitialized.IsUndefinedLogically(),
		"trueIsAccepted":   issetter.True.IsAccepted(),
		"falseIsRejected":  issetter.False.IsRejected(),
		"uninitIsIndet":    issetter.Uninitialized.IsIndeterminate(),
		"uninitIsInvalid":  issetter.Uninitialized.IsInvalid(),
		"trueIsValid":      issetter.True.IsValid(),
	}

	// Assert
	expected := args.Map{
		"trueHasInit": true, "trueHasInitTrue": true,
		"setHasInitSet": true, "trueIsInit": true,
		"trueIsInitBool": true, "trueIsDefBool": true,
		"trueIsInitBoolW": true, "wildIsInitBoolW": true,
		"setIsInitSet": true, "wildIsInitSetW": true,
		"trueIsTrueOrSet": true, "setIsTrueOrSet": true,
		"uninitIsNeg": true, "trueIsPos": true, "setIsPos": true,
		"trueIsDefLogic": true, "uninitIsUndefLog": true,
		"trueIsAccepted": true, "falseIsRejected": true,
		"uninitIsIndet": true, "uninitIsInvalid": true, "trueIsValid": true,
	}
	expected.ShouldBeEqual(t, 0, "Value state checks -- all", actual)
}

func Test_Value_ConversionMethods(t *testing.T) {
	// Act
	actual := args.Map{
		"valueUInt16":  int(issetter.True.ValueUInt16()),
		"valueByte":    int(issetter.True.ValueByte()),
		"valueInt":     issetter.True.ValueInt(),
		"valueInt8":    int(issetter.True.ValueInt8()),
		"valueInt16":   int(issetter.True.ValueInt16()),
		"valueInt32":   int(issetter.True.ValueInt32()),
		"valueString":  issetter.True.ValueString(),
		"stringValue":  issetter.True.StringValue(),
		"toNumberStr":  issetter.True.ToNumberString(),
		"name":         issetter.True.Name(),
		"string":       issetter.True.String(),
		"nameValue":    issetter.True.NameValue() != "",
		"typeName":     issetter.True.TypeName() != "",
	}

	// Assert
	expected := args.Map{
		"valueUInt16": 1, "valueByte": 1, "valueInt": 1,
		"valueInt8": 1, "valueInt16": 1, "valueInt32": 1,
		"valueString": "1", "stringValue": "1", "toNumberStr": "1",
		"name": "True", "string": "True",
		"nameValue": true, "typeName": true,
	}
	expected.ShouldBeEqual(t, 0, "Value conversion methods -- True", actual)
}

func Test_Value_Format_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	result := issetter.True.Format("{type-name}:{name}={value}")

	// Act
	actual := args.Map{"hasContent": result != ""}

	// Assert
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "Value Format -- template", actual)
}

func Test_Value_AllNameValues_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	result := issetter.True.AllNameValues()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 6}
	expected.ShouldBeEqual(t, 0, "Value AllNameValues -- 6 values", actual)
}

func Test_Value_IsNameEqual_FromValueBooleanChecks(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "Value IsNameEqual -- True", actual)
}

func Test_Value_IsAnyNamesOf_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   issetter.True.IsAnyNamesOf("True", "False"),
		"noMatch": issetter.True.IsAnyNamesOf("False", "Set"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Value IsAnyNamesOf -- True", actual)
}

func Test_Value_IsNot_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsNot(issetter.False)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value IsNot -- True != False", actual)
}

func Test_Value_IsValueEqual_FromValueBooleanChecks(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "Value IsValueEqual -- True==1", actual)
}

func Test_Value_IsByteValueEqual_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsByteValueEqual(1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value IsByteValueEqual -- True==1", actual)
}

func Test_Value_Is(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.Is(issetter.True)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Value Is -- True==True", actual)
}

func Test_Value_Add(t *testing.T) {
	// Arrange
	result := issetter.True.Add(1)

	// Act
	actual := args.Map{"val": int(result)}

	// Assert
	expected := args.Map{"val": 2}
	expected.ShouldBeEqual(t, 0, "Value Add -- True+1=2", actual)
}

// ============================================================================
// Comparison methods (int variants)
// ============================================================================

func Test_Value_IntComparisons(t *testing.T) {
	// Arrange
	v := issetter.Set // value 4

	// Act
	actual := args.Map{
		"equalInt":      v.IsEqualInt(4),
		"greaterInt":    v.IsGreaterInt(3),
		"greaterEqInt":  v.IsGreaterEqualInt(4),
		"lessInt":       v.IsLessInt(5),
		"lessEqInt":     v.IsLessEqualInt(4),
		"betweenInt":    v.IsBetweenInt(3, 5),
		"notBetweenInt": v.IsBetweenInt(0, 2),
	}

	// Assert
	expected := args.Map{
		"equalInt": true, "greaterInt": true,
		"greaterEqInt": true, "lessInt": true,
		"lessEqInt": true, "betweenInt": true, "notBetweenInt": false,
	}
	expected.ShouldBeEqual(t, 0, "Value int comparisons -- Set(4)", actual)
}

func Test_Value_IsAnyValuesEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   issetter.True.IsAnyValuesEqual(0, 1, 2),
		"noMatch": issetter.True.IsAnyValuesEqual(3, 4, 5),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Value IsAnyValuesEqual -- True in 0,1,2", actual)
}

// ============================================================================
// MarshalJSON / UnmarshalJSON / Serialize / Deserialize
// ============================================================================

func Test_Value_MarshalUnmarshal(t *testing.T) {
	// Arrange
	mb, _ := issetter.True.MarshalJSON()
	var v issetter.Value
	err := v.UnmarshalJSON(mb)

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"val": v == issetter.True,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "Value Marshal/Unmarshal roundtrip -- True", actual)
}

func Test_Value_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	var v issetter.Value
	err := v.UnmarshalJSON(nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Value UnmarshalJSON nil -- error", actual)
}

func Test_Value_UnmarshalJSON_Unknown(t *testing.T) {
	// Arrange
	var v issetter.Value
	err := v.UnmarshalJSON([]byte(`"garbage_xyz"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Value UnmarshalJSON unknown -- error", actual)
}

func Test_Value_Serialize(t *testing.T) {
	// Arrange
	bytes, err := issetter.True.Serialize()

	// Act
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"hasBytes": true,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Value Serialize -- True", actual)
}

func Test_Value_Deserialize(t *testing.T) {
	// Arrange
	v, err := issetter.Uninitialized.Deserialize([]byte(`"True"`))

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"isTrue": v == issetter.True,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"isTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "Value Deserialize -- True", actual)
}

func Test_Value_Deserialize_Invalid(t *testing.T) {
	// Arrange
	_, err := issetter.Uninitialized.Deserialize([]byte(`"garbage"`))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Value Deserialize invalid -- error", actual)
}

func Test_Value_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	val, err := issetter.Uninitialized.UnmarshallEnumToValue([]byte(`"Set"`))

	// Act
	actual := args.Map{
		"val": int(val),
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"val": 4,
		"hasErr": false,
	}
	expected.ShouldBeEqual(t, 0, "Value UnmarshallEnumToValue -- Set", actual)
}

// ============================================================================
// EnumType, MinMaxAny, Ranges
// ============================================================================

func Test_Value_EnumType_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{"notNil": issetter.True.EnumType() != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Value EnumType -- not nil", actual)
}

func Test_Value_MinMaxAny_FromValueBooleanChecks(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "Value MinMaxAny -- not nil", actual)
}

func Test_Value_MinMaxStrings_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"minStr":  issetter.True.MinValueString(),
		"maxStr":  issetter.True.MaxValueString(),
		"maxInt":  issetter.True.MaxInt(),
		"minInt":  issetter.True.MinInt(),
		"maxByte": int(issetter.True.MaxByte()),
		"minByte": int(issetter.True.MinByte()),
	}

	// Assert
	expected := args.Map{
		"minStr": "0", "maxStr": "5",
		"maxInt": 5, "minInt": 0,
		"maxByte": 5, "minByte": 0,
	}
	expected.ShouldBeEqual(t, 0, "Value min/max strings and ints -- True", actual)
}

func Test_Value_RangesDynamicMap_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	rdm := issetter.True.RangesDynamicMap()

	// Act
	actual := args.Map{"len": len(rdm)}

	// Assert
	expected := args.Map{"len": 6}
	expected.ShouldBeEqual(t, 0, "Value RangesDynamicMap -- 6 entries", actual)
}

func Test_Value_IntegerEnumRanges_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	ranges := issetter.True.IntegerEnumRanges()

	// Act
	actual := args.Map{"len": len(ranges)}

	// Assert
	expected := args.Map{"len": 6}
	expected.ShouldBeEqual(t, 0, "Value IntegerEnumRanges -- 6", actual)
}

func Test_Value_RangeNamesCsv_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	csv := issetter.True.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": csv != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value RangeNamesCsv -- not empty", actual)
}

func Test_Value_ToPtr(t *testing.T) {
	// Arrange
	ptr := issetter.True.ToPtr()

	// Act
	actual := args.Map{
		"notNil": ptr != nil,
		"val": *ptr == issetter.True,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "Value ToPtr -- True", actual)
}

// ============================================================================
// OnlySupportedErr
// ============================================================================

func Test_Value_OnlySupportedErr_Empty(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedErr()

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Value OnlySupportedErr empty -- nil", actual)
}

func Test_Value_OnlySupportedMsgErr_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	err := issetter.True.OnlySupportedMsgErr("prefix: ", "True", "False", "Uninitialized", "Set", "Unset", "Wildcard")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Value OnlySupportedMsgErr all -- nil", actual)
}

// ============================================================================
// GetSetBoolOnInvalid / GetSetBoolOnInvalidFunc — uninit paths
// ============================================================================

func Test_GetSetBoolOnInvalid_Uninit(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	result := v.GetSetBoolOnInvalid(true)

	// Act
	actual := args.Map{
		"result": result,
		"isTrue": v == issetter.True,
	}

	// Assert
	expected := args.Map{
		"result": true,
		"isTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSetBoolOnInvalid uninit sets -- true", actual)
}

func Test_GetSetBoolOnInvalidFunc_Uninit(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	result := v.GetSetBoolOnInvalidFunc(func() bool { return false })

	// Act
	actual := args.Map{
		"result": result,
		"isFalse": v == issetter.False,
	}

	// Assert
	expected := args.Map{
		"result": false,
		"isFalse": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSetBoolOnInvalidFunc uninit sets -- false", actual)
}

// ============================================================================
// LazyEvaluateBool / LazyEvaluateSet — uninit paths
// ============================================================================

func Test_LazyEvaluateBool_Uninit(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := false
	result := v.LazyEvaluateBool(func() { called = true })

	// Act
	actual := args.Map{
		"called": called,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyEvaluateBool uninit calls -- called", actual)
}

func Test_LazyEvaluateSet_Uninit(t *testing.T) {
	// Arrange
	v := issetter.Uninitialized
	called := false
	result := v.LazyEvaluateSet(func() { called = true })

	// Act
	actual := args.Map{
		"called": called,
		"result": result,
	}

	// Assert
	expected := args.Map{
		"called": true,
		"result": true,
	}
	expected.ShouldBeEqual(t, 0, "LazyEvaluateSet uninit calls -- called", actual)
}

// ============================================================================
// GetErrorOnOutOfRange
// ============================================================================

func Test_GetErrorOnOutOfRange_InRange(t *testing.T) {
	// Arrange
	err := issetter.True.GetErrorOnOutOfRange(1, "out of range")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "GetErrorOnOutOfRange in range -- nil", actual)
}

func Test_GetErrorOnOutOfRange_OutOfRange(t *testing.T) {
	// Arrange
	err := issetter.True.GetErrorOnOutOfRange(255, "out of range")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "GetErrorOnOutOfRange out of range -- error", actual)
}

// ============================================================================
// IsCompareResult
// ============================================================================

func Test_IsCompareResult_Equal(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.Equal)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult Equal -- True==1", actual)
}

func Test_IsCompareResult_LeftGreater(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.Set.IsCompareResult(1, corecomparator.LeftGreater)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult LeftGreater -- Set>1", actual)
}

func Test_IsCompareResult_LeftLess(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(5, corecomparator.LeftLess)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult LeftLess -- True<5", actual)
}

func Test_IsCompareResult_NotEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(2, corecomparator.NotEqual)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult NotEqual -- True!=2", actual)
}

func Test_IsCompareResult_LeftGreaterEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.LeftGreaterEqual)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult LeftGreaterEqual -- True>=1", actual)
}

func Test_IsCompareResult_LeftLessEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": issetter.True.IsCompareResult(1, corecomparator.LeftLessEqual)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsCompareResult LeftLessEqual -- True<=1", actual)
}

// ============================================================================
// Package-level functions
// ============================================================================

func Test_IsOutOfRange(t *testing.T) {
	// Act
	actual := args.Map{
		"inRange":  issetter.IsOutOfRange(1),
		"outRange": issetter.IsOutOfRange(255),
	}

	// Assert
	expected := args.Map{
		"inRange": false,
		"outRange": true,
	}
	expected.ShouldBeEqual(t, 0, "IsOutOfRange -- 1 and 255", actual)
}

func Test_CombinedBooleans(t *testing.T) {
	// Act
	actual := args.Map{
		"allTrue":  issetter.CombinedBooleans(true, true) == issetter.True,
		"oneFalse": issetter.CombinedBooleans(true, false) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"allTrue": true,
		"oneFalse": true,
	}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans -- various", actual)
}

func Test_GetSetByte_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"trueVal":  int(issetter.GetSetByte(true, 1, 2)),
		"falseVal": int(issetter.GetSetByte(false, 1, 2)),
	}

	// Assert
	expected := args.Map{
		"trueVal": 1,
		"falseVal": 2,
	}
	expected.ShouldBeEqual(t, 0, "GetSetByte -- true/false", actual)
}

func Test_GetSetterByComparing(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   issetter.GetSetterByComparing(issetter.True, issetter.False, "a", "a", "b") == issetter.True,
		"noMatch": issetter.GetSetterByComparing(issetter.True, issetter.False, "c", "a", "b") == issetter.False,
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing -- match and no match", actual)
}

func Test_New_Valid_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	v, err := issetter.New("true")

	// Act
	actual := args.Map{
		"hasErr": err != nil,
		"isTrue": v == issetter.True,
	}

	// Assert
	expected := args.Map{
		"hasErr": false,
		"isTrue": true,
	}
	expected.ShouldBeEqual(t, 0, "New valid -- true", actual)
}

func Test_New_Invalid_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	_, err := issetter.New("garbage")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New invalid -- error", actual)
}

func Test_NewMust_Valid_FromValueBooleanChecks(t *testing.T) {
	// Arrange
	v := issetter.NewMust("false")

	// Act
	actual := args.Map{"isFalse": v == issetter.False}

	// Assert
	expected := args.Map{"isFalse": true}
	expected.ShouldBeEqual(t, 0, "NewMust valid -- false", actual)
}

func Test_NewBool_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"true":  issetter.NewBool(true) == issetter.True,
		"false": issetter.NewBool(false) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"true": true,
		"false": true,
	}
	expected.ShouldBeEqual(t, 0, "NewBool -- true/false", actual)
}

func Test_NewBooleans(t *testing.T) {
	// Act
	actual := args.Map{
		"allTrue":  issetter.NewBooleans(true, true) == issetter.True,
		"oneFalse": issetter.NewBooleans(true, false) == issetter.False,
	}

	// Assert
	expected := args.Map{
		"allTrue": true,
		"oneFalse": true,
	}
	expected.ShouldBeEqual(t, 0, "NewBooleans -- various", actual)
}

func Test_GetSetUnset_FromValueBooleanChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"trueIsSet":    issetter.GetSetUnset(true) == issetter.Set,
		"falseIsUnset": issetter.GetSetUnset(false) == issetter.Unset,
	}

	// Assert
	expected := args.Map{
		"trueIsSet": true,
		"falseIsUnset": true,
	}
	expected.ShouldBeEqual(t, 0, "GetSetUnset -- true/false", actual)
}

func Test_Max_Min(t *testing.T) {
	// Act
	actual := args.Map{
		"max":     issetter.Max() == issetter.Wildcard,
		"min":     issetter.Min() == issetter.Uninitialized,
		"maxByte": int(issetter.MaxByte()),
		"minByte": int(issetter.MinByte()),
	}

	// Assert
	expected := args.Map{
		"max": true,
		"min": true,
		"maxByte": 4,
		"minByte": 0,
	}
	expected.ShouldBeEqual(t, 0, "Max/Min -- expected", actual)
}

func Test_RangeNamesCsv(t *testing.T) {
	// Arrange
	csv := issetter.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": csv != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv -- not empty", actual)
}

func Test_IntegerEnumRanges(t *testing.T) {
	// Arrange
	ranges := issetter.IntegerEnumRanges()

	// Act
	actual := args.Map{"len": len(ranges)}

	// Assert
	expected := args.Map{"len": 6}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges -- 6", actual)
}

// ============================================================================
// ToBooleanValue / ToSetUnsetValue — more values
// ============================================================================

func Test_ToBooleanValue_SetUnset(t *testing.T) {
	// Act
	actual := args.Map{
		"setToTrue":   issetter.Set.ToBooleanValue() == issetter.True,
		"unsetToFalse": issetter.Unset.ToBooleanValue() == issetter.False,
		"uninitStays":  issetter.Uninitialized.ToBooleanValue() == issetter.Uninitialized,
	}

	// Assert
	expected := args.Map{
		"setToTrue": true,
		"unsetToFalse": true,
		"uninitStays": true,
	}
	expected.ShouldBeEqual(t, 0, "ToBooleanValue Set/Unset -- expected", actual)
}

func Test_ToSetUnsetValue_TrueFalse(t *testing.T) {
	// Act
	actual := args.Map{
		"trueToSet":   issetter.True.ToSetUnsetValue() == issetter.Set,
		"falseToUnset": issetter.False.ToSetUnsetValue() == issetter.Unset,
		"uninitStays":  issetter.Uninitialized.ToSetUnsetValue() == issetter.Uninitialized,
	}

	// Assert
	expected := args.Map{
		"trueToSet": true,
		"falseToUnset": true,
		"uninitStays": true,
	}
	expected.ShouldBeEqual(t, 0, "ToSetUnsetValue True/False -- expected", actual)
}
