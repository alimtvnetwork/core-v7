package corecomparatortests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Cover remaining Compare methods not hit by existing tests

func Test_Compare_IsLess_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftLess.IsLess()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftLess should IsLess", actual)
	actual := args.Map{"result": corecomparator.Equal.IsLess()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not IsLess", actual)
}

func Test_Compare_IsLessEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftLess.IsLessEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftLess should IsLessEqual", actual)
	actual := args.Map{"result": corecomparator.Equal.IsLessEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should IsLessEqual", actual)
	actual := args.Map{"result": corecomparator.LeftGreater.IsLessEqual()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "LeftGreater should not IsLessEqual", actual)
}

func Test_Compare_IsGreater_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftGreater.IsGreater()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftGreater should IsGreater", actual)
	actual := args.Map{"result": corecomparator.Equal.IsGreater()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not IsGreater", actual)
}

func Test_Compare_IsGreaterEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftGreater.IsGreaterEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftGreater should IsGreaterEqual", actual)
	actual := args.Map{"result": corecomparator.Equal.IsGreaterEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should IsGreaterEqual", actual)
}

func Test_Compare_IsNameEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsNameEqual("Equal")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match Equal name", actual)
	actual := args.Map{"result": corecomparator.Equal.IsNameEqual("NotEqual")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match NotEqual", actual)
}

func Test_Compare_ToNumberString_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.ToNumberString() != "0"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Compare_IsDefined_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsDefined()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should be defined", actual)
	actual := args.Map{"result": corecomparator.Inconclusive.IsDefined()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Inconclusive should not be defined", actual)
}

func Test_Compare_IsValid_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsValid()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should be valid", actual)
}

func Test_Compare_IsEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected equal", actual)
}

func Test_Compare_IsNotEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.NotEqual.IsNotEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected not equal", actual)
}

func Test_Compare_IsNotEqualLogically_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsNotEqualLogically()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not be logically not-equal", actual)
	actual := args.Map{"result": corecomparator.LeftGreater.IsNotEqualLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftGreater should be logically not-equal", actual)
}

func Test_Compare_IsLeftLess_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftLess.IsLeftLess()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Compare_IsLeftLessEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftLessEqual.IsLeftLessEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Compare_IsLeftLessEqualLogically_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftLess.IsLeftLessEqualLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.LeftLessEqual.IsLeftLessEqualLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.Equal.IsLeftLessEqualLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Compare_IsLeftGreaterEqualLogically_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftGreater.IsLeftGreaterEqualLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.LeftGreaterEqual.IsLeftGreaterEqualLogically()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Compare_IsLeftGreaterOrGreaterEqualOrEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsLeftGreaterOrGreaterEqualOrEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.LeftGreater.IsLeftGreaterOrGreaterEqualOrEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.LeftGreaterEqual.IsLeftGreaterOrGreaterEqualOrEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Compare_IsInconclusiveOrNotEqual_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Inconclusive.IsInconclusiveOrNotEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.NotEqual.IsInconclusiveOrNotEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual := args.Map{"result": corecomparator.Equal.IsInconclusiveOrNotEqual()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Compare_IsDefinedProperly_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsDefinedProperly()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_Compare_IsAnyOf_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsAnyOf()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty should return true", actual)
	actual := args.Map{"result": corecomparator.Equal.IsAnyOf(corecomparator.NotEqual, corecomparator.Equal)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find Equal", actual)
	actual := args.Map{"result": corecomparator.Equal.IsAnyOf(corecomparator.NotEqual)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not find Equal in [NotEqual]", actual)
}

func Test_Compare_NameValue_Cov4(t *testing.T) {
	r := corecomparator.Equal.NameValue()
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_Compare_CsvStrings_Cov4(t *testing.T) {
	r := corecomparator.Equal.CsvStrings()
	actual := args.Map{"result": len(r) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty args should return empty slice", actual)
	r = corecomparator.Equal.CsvStrings(corecomparator.Equal, corecomparator.NotEqual)
	actual := args.Map{"result": len(r) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_Compare_CsvString_Cov4(t *testing.T) {
	r := corecomparator.Equal.CsvString()
	actual := args.Map{"result": r != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty args should return empty", actual)
	r = corecomparator.Equal.CsvString(corecomparator.Equal)
	actual := args.Map{"result": r == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_Compare_MarshalJSON_Cov4(t *testing.T) {
	data, err := json.Marshal(corecomparator.Equal)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not error", actual)
	actual := args.Map{"result": string(data) == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_Compare_Value_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.Value() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Compare_ValueByte_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.ValueByte() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Compare_ValueInt_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.ValueInt() != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_Compare_OperatorSymbol_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.OperatorSymbol() != "="}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected =", actual)
}

func Test_Compare_OperatorShortForm_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.OperatorShortForm() != "eq"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected eq", actual)
}

func Test_Compare_NumberJsonString_Cov4(t *testing.T) {
	r := corecomparator.Equal.NumberJsonString()
	actual := args.Map{"result": r != "\"0\""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected quoted 0", actual)
}

func Test_Compare_IsAnyNamesOf_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.IsAnyNamesOf("NotEqual", "Equal")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should find Equal", actual)
	actual := args.Map{"result": corecomparator.Equal.IsAnyNamesOf("NotEqual")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not find", actual)
}

func Test_Compare_IsCompareEqualLogically_Cov4(t *testing.T) {
	// it == expectedCompare
	actual := args.Map{"result": corecomparator.Equal.IsCompareEqualLogically(corecomparator.Equal)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// expectedCompare == NotEqual
	actual := args.Map{"result": corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.NotEqual)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "LeftGreater is logically not-equal", actual)
	// expectedCompare.IsLeftGreaterEqualLogically
	actual := args.Map{"result": corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.LeftGreaterEqual)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// expectedCompare.IsLeftLessEqualLogically
	actual := args.Map{"result": corecomparator.LeftLess.IsCompareEqualLogically(corecomparator.LeftLessEqual)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	// fallthrough false
	actual := args.Map{"result": corecomparator.Inconclusive.IsCompareEqualLogically(corecomparator.LeftGreater)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_Compare_OnlySupportedErr_Cov4(t *testing.T) {
	// with message, supported
	err := corecomparator.Equal.OnlySupportedErr("test", corecomparator.Equal)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be nil", actual)
	// with message, not supported
	err = corecomparator.LeftGreater.OnlySupportedErr("test", corecomparator.Equal)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should error", actual)
	// empty message delegates to OnlySupportedDirectErr
	err = corecomparator.Equal.OnlySupportedErr("", corecomparator.Equal)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be nil", actual)
}

func Test_Compare_OnlySupportedDirectErr_Cov4(t *testing.T) {
	err := corecomparator.Equal.OnlySupportedDirectErr(corecomparator.Equal)
	actual := args.Map{"result": err != nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should be nil", actual)
	err = corecomparator.LeftGreater.OnlySupportedDirectErr(corecomparator.Equal)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should error", actual)
}

func Test_Min_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Min() != corecomparator.Equal}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
}

func Test_Max_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.Max() != corecomparator.NotEqual}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected NotEqual", actual)
}

func Test_MinLength_Cov4(t *testing.T) {
	actual := args.Map{"result": corecomparator.MinLength(3, 5) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	actual := args.Map{"result": corecomparator.MinLength(5, 3) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_Ranges_Cov4(t *testing.T) {
	r := corecomparator.Ranges()
	actual := args.Map{"result": len(r) == 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}
