package issettertests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/issetter"
)

func Test_New_Valid_Cov2(t *testing.T) {
	v, err := issetter.New("True")
	actual := args.Map{"hasErr": err != nil, "isTrue": v == issetter.True}
	expected := args.Map{"hasErr": false, "isTrue": true}
	expected.ShouldBeEqual(t, 0, "New_Valid returns non-empty -- with args", actual)
}

func Test_New_Invalid_Cov2(t *testing.T) {
	_, err := issetter.New("invalid_name_xyz")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "New_Invalid returns error -- with args", actual)
}

func Test_NewMust_Valid_Cov2(t *testing.T) {
	actual := args.Map{"isTrue": issetter.NewMust("True") == issetter.True}
	expected := args.Map{"isTrue": true}
	expected.ShouldBeEqual(t, 0, "NewMust_Valid returns non-empty -- with args", actual)
}

func Test_NewBool_Cov2(t *testing.T) {
	actual := args.Map{
		"true":  fmt.Sprintf("%v", issetter.NewBool(true)),
		"false": fmt.Sprintf("%v", issetter.NewBool(false)),
	}
	expected := args.Map{
		"true":  fmt.Sprintf("%v", issetter.True),
		"false": fmt.Sprintf("%v", issetter.False),
	}
	expected.ShouldBeEqual(t, 0, "NewBool returns correct value -- with args", actual)
}

func Test_NewBooleans_Cov2(t *testing.T) {
	actual := args.Map{
		"allTrue":  fmt.Sprintf("%v", issetter.NewBooleans(true, true)),
		"oneFalse": fmt.Sprintf("%v", issetter.NewBooleans(true, false)),
	}
	expected := args.Map{
		"allTrue":  fmt.Sprintf("%v", issetter.True),
		"oneFalse": fmt.Sprintf("%v", issetter.False),
	}
	expected.ShouldBeEqual(t, 0, "NewBooleans returns correct value -- with args", actual)
}

func Test_CombinedBooleans_Cov2(t *testing.T) {
	actual := args.Map{
		"empty":    fmt.Sprintf("%v", issetter.CombinedBooleans()),
		"allTrue":  fmt.Sprintf("%v", issetter.CombinedBooleans(true, true, true)),
		"oneFalse": fmt.Sprintf("%v", issetter.CombinedBooleans(true, false)),
	}
	expected := args.Map{
		"empty":    fmt.Sprintf("%v", issetter.True),
		"allTrue":  fmt.Sprintf("%v", issetter.True),
		"oneFalse": fmt.Sprintf("%v", issetter.False),
	}
	expected.ShouldBeEqual(t, 0, "CombinedBooleans returns correct value -- with args", actual)
}

func Test_GetSetByte_Cov2(t *testing.T) {
	actual := args.Map{
		"true":  fmt.Sprintf("%v", issetter.GetSetByte(true, byte(issetter.Set), byte(issetter.Unset))),
		"false": fmt.Sprintf("%v", issetter.GetSetByte(false, byte(issetter.Set), byte(issetter.Unset))),
	}
	expected := args.Map{
		"true":  fmt.Sprintf("%v", issetter.Set),
		"false": fmt.Sprintf("%v", issetter.Unset),
	}
	expected.ShouldBeEqual(t, 0, "GetSetByte returns correct value -- with args", actual)
}

func Test_GetSetUnset_Cov2(t *testing.T) {
	actual := args.Map{
		"true":  fmt.Sprintf("%v", issetter.GetSetUnset(true)),
		"false": fmt.Sprintf("%v", issetter.GetSetUnset(false)),
	}
	expected := args.Map{
		"true":  fmt.Sprintf("%v", issetter.Set),
		"false": fmt.Sprintf("%v", issetter.Unset),
	}
	expected.ShouldBeEqual(t, 0, "GetSetUnset returns correct value -- with args", actual)
}

func Test_GetSetterByComparing_Cov2(t *testing.T) {
	actual := args.Map{
		"match":   fmt.Sprintf("%v", issetter.GetSetterByComparing(issetter.True, issetter.False, 42, 1, 42, 100)),
		"noMatch": fmt.Sprintf("%v", issetter.GetSetterByComparing(issetter.True, issetter.False, 42, 1, 2, 3)),
	}
	expected := args.Map{
		"match":   fmt.Sprintf("%v", issetter.True),
		"noMatch": fmt.Sprintf("%v", issetter.False),
	}
	expected.ShouldBeEqual(t, 0, "GetSetterByComparing returns correct value -- with args", actual)
}

func Test_IsCompareResult_Cov2(t *testing.T) {
	actual := args.Map{
		"equal":        issetter.True.IsCompareResult(1, corecomparator.Equal),
		"leftGreater":  issetter.Set.IsCompareResult(1, corecomparator.LeftGreater),
		"leftGtEqual":  issetter.True.IsCompareResult(1, corecomparator.LeftGreaterEqual),
		"leftLess":     issetter.True.IsCompareResult(2, corecomparator.LeftLess),
		"leftLtEqual":  issetter.True.IsCompareResult(1, corecomparator.LeftLessEqual),
		"notEqual":     issetter.True.IsCompareResult(2, corecomparator.NotEqual),
	}
	expected := args.Map{
		"equal":        true,
		"leftGreater":  true,
		"leftGtEqual":  true,
		"leftLess":     true,
		"leftLtEqual":  true,
		"notEqual":     true,
	}
	expected.ShouldBeEqual(t, 0, "IsCompareResult returns correct value -- with args", actual)
}

func Test_IsOutOfRange_Cov2(t *testing.T) {
	actual := args.Map{"inRange": issetter.IsOutOfRange(1), "outOfRange": issetter.IsOutOfRange(255)}
	expected := args.Map{"inRange": false, "outOfRange": true}
	expected.ShouldBeEqual(t, 0, "IsOutOfRange returns correct value -- with args", actual)
}

func Test_Value_Methods_Cov2(t *testing.T) {
	actual := args.Map{
		"is":              issetter.True.Is(issetter.True),
		"isEqual":         issetter.True.IsEqual(1),
		"isGreater":       issetter.Set.IsGreater(1),
		"isGreaterEqual":  issetter.True.IsGreaterEqual(1),
		"isLess":          issetter.True.IsLess(2),
		"isLessEqual":     issetter.True.IsLessEqual(1),
		"isEqualInt":      issetter.True.IsEqualInt(1),
		"isGreaterInt":    issetter.Set.IsGreaterInt(1),
		"isGreaterEqualInt": issetter.True.IsGreaterEqualInt(1),
		"isLessInt":       issetter.True.IsLessInt(2),
		"isLessEqualInt":  issetter.True.IsLessEqualInt(1),
		"isBetween_yes":   issetter.True.IsBetween(0, 5),
		"isBetween_no":    issetter.True.IsBetween(2, 5),
		"isBetweenInt":    issetter.True.IsBetweenInt(0, 5),
	}
	expected := args.Map{
		"is":              true,
		"isEqual":         true,
		"isGreater":       true,
		"isGreaterEqual":  true,
		"isLess":          true,
		"isLessEqual":     true,
		"isEqualInt":      true,
		"isGreaterInt":    true,
		"isGreaterEqualInt": true,
		"isLessInt":       true,
		"isLessEqualInt":  true,
		"isBetween_yes":   true,
		"isBetween_no":    false,
		"isBetweenInt":    true,
	}
	expected.ShouldBeEqual(t, 0, "Value_Methods returns correct value -- with args", actual)
}

func Test_Value_Add_Cov2(t *testing.T) {
	actual := args.Map{"result": fmt.Sprintf("%v", issetter.True.Add(1))}
	expected := args.Map{"result": fmt.Sprintf("%v", issetter.False)}
	expected.ShouldBeEqual(t, 0, "Value_Add returns correct value -- with args", actual)
}

func Test_Value_IsNegative_Cov2(t *testing.T) {
	actual := args.Map{
		"uninitialized": issetter.Uninitialized.IsNegative(),
		"false":         issetter.False.IsNegative(),
		"unset":         issetter.Unset.IsNegative(),
		"true":          issetter.True.IsNegative(),
	}
	expected := args.Map{
		"uninitialized": true,
		"false":         true,
		"unset":         true,
		"true":          false,
	}
	expected.ShouldBeEqual(t, 0, "Value_IsNegative returns correct value -- with args", actual)
}

func Test_Value_IsPositive_Cov2(t *testing.T) {
	actual := args.Map{"true": issetter.True.IsPositive(), "set": issetter.Set.IsPositive()}
	expected := args.Map{"true": true, "set": true}
	expected.ShouldBeEqual(t, 0, "Value_IsPositive returns correct value -- with args", actual)
}

func Test_Value_GetErrorOnOutOfRange_Cov2(t *testing.T) {
	actual := args.Map{
		"inRange":    issetter.True.GetErrorOnOutOfRange(1, "test") == nil,
		"outOfRange": issetter.True.GetErrorOnOutOfRange(255, "test") != nil,
	}
	expected := args.Map{"inRange": true, "outOfRange": true}
	expected.ShouldBeEqual(t, 0, "Value_GetErrorOnOutOfRange returns error -- with args", actual)
}

func Test_Value_NameMethods_Cov2(t *testing.T) {
	actual := args.Map{
		"yesNo":            issetter.True.YesNoMappedValue(),
		"yesNoFalse":       issetter.False.YesNoMappedValue(),
		"yesNoUninit":      issetter.Uninitialized.YesNoMappedValue(),
		"yesNoLower":       issetter.True.YesNoLowercaseName(),
		"yesNoName":        issetter.True.YesNoName(),
		"trueFalseName":    issetter.True.TrueFalseName(),
		"onOffLower":       issetter.True.OnOffLowercaseName(),
		"onOffName":        issetter.True.OnOffName(),
		"trueFalseLower":   issetter.True.TrueFalseLowercaseName(),
		"setUnsetLower":    issetter.True.SetUnsetLowercaseName(),
	}
	expected := args.Map{
		"yesNo":            "yes",
		"yesNoFalse":       "no",
		"yesNoUninit":      "",
		"yesNoLower":       "yes",
		"yesNoName":        "Yes",
		"trueFalseName":    "True",
		"onOffLower":       "on",
		"onOffName":        "On",
		"trueFalseLower":   "true",
		"setUnsetLower":    "set",
	}
	expected.ShouldBeEqual(t, 0, "Value_NameMethods returns correct value -- with args", actual)
}

func Test_Value_Serialize_Cov2(t *testing.T) {
	data, err := issetter.True.Serialize()
	actual := args.Map{"hasErr": err != nil, "hasData": len(data) > 0}
	expected := args.Map{"hasErr": false, "hasData": true}
	expected.ShouldBeEqual(t, 0, "Value_Serialize returns correct value -- with args", actual)
}

func Test_Value_TypeName_Cov2(t *testing.T) {
	actual := args.Map{"notEmpty": issetter.True.TypeName() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Value_TypeName returns correct value -- with args", actual)
}

func Test_Value_IsAnyValuesEqual_Cov2(t *testing.T) {
	actual := args.Map{
		"found":    issetter.True.IsAnyValuesEqual(0, 1, 2),
		"notFound": issetter.True.IsAnyValuesEqual(0, 2, 3),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "Value_IsAnyValuesEqual returns non-empty -- with args", actual)
}

func Test_Value_UnmarshallEnumToValue_Cov2(t *testing.T) {
	val, err := issetter.Uninitialized.UnmarshallEnumToValue([]byte(`"True"`))
	actual := args.Map{"hasErr": err != nil, "val": int(val)}
	expected := args.Map{"hasErr": false, "val": 1}
	expected.ShouldBeEqual(t, 0, "Value_UnmarshallEnumToValue returns correct value -- with args", actual)
}

func Test_Value_Deserialize_Cov2(t *testing.T) {
	v, err := issetter.Uninitialized.Deserialize([]byte(`"True"`))
	actual := args.Map{"hasErr": err != nil, "isTrue": v == issetter.True}
	expected := args.Map{"hasErr": false, "isTrue": true}
	expected.ShouldBeEqual(t, 0, "Value_Deserialize_Valid returns non-empty -- with args", actual)
}

func Test_Value_Deserialize_Invalid_Cov2(t *testing.T) {
	_, err := issetter.Uninitialized.Deserialize([]byte(`"INVALID_XYZ"`))
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Value_Deserialize_Invalid returns error -- with args", actual)
}

func Test_Value_UnmarshalJSON_Cov2(t *testing.T) {
	var v issetter.Value
	errNil := v.UnmarshalJSON(nil)
	var v2 issetter.Value
	errInvalid := v2.UnmarshalJSON([]byte(`"UNKNOWN_XYZ"`))
	actual := args.Map{"nilErr": errNil != nil, "invalidErr": errInvalid != nil}
	expected := args.Map{"nilErr": true, "invalidErr": true}
	expected.ShouldBeEqual(t, 0, "Value_UnmarshalJSON returns correct value -- with args", actual)
}

func Test_Value_MinMaxByte_Cov2(t *testing.T) {
	actual := args.Map{
		"maxByte": int(issetter.True.MaxByte()),
		"minByte": int(issetter.True.MinByte()),
	}
	expected := args.Map{
		"maxByte": int(issetter.Wildcard.ValueByte()),
		"minByte": 0,
	}
	expected.ShouldBeEqual(t, 0, "Value_MinMaxByte returns correct value -- with args", actual)
}

func Test_Value_ToPtr_Cov2(t *testing.T) {
	ptr := issetter.True.ToPtr()
	actual := args.Map{"notNil": ptr != nil, "isTrue": *ptr == issetter.True}
	expected := args.Map{"notNil": true, "isTrue": true}
	expected.ShouldBeEqual(t, 0, "Value_ToPtr returns correct value -- with args", actual)
}

func Test_Value_ValueUInt16_Cov2(t *testing.T) {
	actual := args.Map{"result": int(issetter.True.ValueUInt16())}
	expected := args.Map{"result": 1}
	expected.ShouldBeEqual(t, 0, "Value_ValueUInt16 returns correct value -- with args", actual)
}

func Test_Value_IsNo_Cov2(t *testing.T) {
	actual := args.Map{"false": issetter.False.IsNo(), "true": issetter.True.IsNo()}
	expected := args.Map{"false": true, "true": false}
	expected.ShouldBeEqual(t, 0, "Value_IsNo returns correct value -- with args", actual)
}

func Test_Value_IsWildcardOrBool_False_Cov2(t *testing.T) {
	actual := args.Map{"result": issetter.True.IsWildcardOrBool(false)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Value_IsWildcardOrBool_False returns non-empty -- with args", actual)
}

func Test_PackageLevelFuncs_Cov2(t *testing.T) {
	actual := args.Map{
		"min":           fmt.Sprintf("%v", issetter.Min()),
		"max":           fmt.Sprintf("%v", issetter.Max()),
		"minByte":       int(issetter.MinByte()),
		"maxByte":       int(issetter.MaxByte()),
		"rangeNotEmpty": issetter.RangeNamesCsv() != "",
	}
	expected := args.Map{
		"min":           fmt.Sprintf("%v", issetter.Uninitialized),
		"max":           fmt.Sprintf("%v", issetter.Wildcard),
		"minByte":       0,
		"maxByte":       int(issetter.Set.Value()),
		"rangeNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "PackageLevelFuncs returns correct value -- with args", actual)
}

func Test_Value_OnlySupportedErr_Empty_Cov2(t *testing.T) {
	actual := args.Map{"hasErr": issetter.True.OnlySupportedErr() != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Value_OnlySupportedErr_Empty returns empty -- with args", actual)
}

func Test_Value_OnlySupportedMsgErr_Nil_Cov2(t *testing.T) {
	names := []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}
	actual := args.Map{"hasErr": issetter.True.OnlySupportedMsgErr("prefix: ", names...) != nil}
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "Value_OnlySupportedMsgErr_Nil returns nil -- with args", actual)
}
