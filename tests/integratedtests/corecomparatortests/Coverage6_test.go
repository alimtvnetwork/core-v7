package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Compare methods ──

func Test_Cov6_Compare_Is(t *testing.T) {
	actual := args.Map{
		"same": corecomparator.Equal.Is(corecomparator.Equal),
		"diff": corecomparator.Equal.Is(corecomparator.LeftGreater),
	}
	expected := args.Map{"same": true, "diff": false}
	expected.ShouldBeEqual(t, 0, "Compare.Is returns correct value -- with args", actual)
}

func Test_Cov6_Compare_LessEqual(t *testing.T) {
	actual := args.Map{
		"less":     corecomparator.LeftLess.IsLessEqual(),
		"equal":    corecomparator.Equal.IsLessEqual(),
		"greater":  corecomparator.LeftGreater.IsLessEqual(),
	}
	expected := args.Map{"less": true, "equal": true, "greater": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- LessEqual", actual)
}

func Test_Cov6_Compare_GreaterEqual(t *testing.T) {
	actual := args.Map{
		"greater": corecomparator.LeftGreater.IsGreaterEqual(),
		"equal":   corecomparator.Equal.IsGreaterEqual(),
		"less":    corecomparator.LeftLess.IsGreaterEqual(),
	}
	expected := args.Map{"greater": true, "equal": true, "less": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- GreaterEqual", actual)
}

func Test_Cov6_Compare_IsNameEqual(t *testing.T) {
	actual := args.Map{
		"match":   corecomparator.Equal.IsNameEqual("Equal"),
		"noMatch": corecomparator.Equal.IsNameEqual("NotEqual"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsNameEqual", actual)
}

func Test_Cov6_Compare_ToNumberString(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.ToNumberString()}
	expected := args.Map{"result": "0"}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- ToNumberString", actual)
}

func Test_Cov6_Compare_IsLeftLessOrLessEqualOrEqual(t *testing.T) {
	actual := args.Map{
		"less":    corecomparator.LeftLess.IsLeftLessOrLessEqualOrEqual(),
		"le":      corecomparator.LeftLessEqual.IsLeftLessOrLessEqualOrEqual(),
		"eq":      corecomparator.Equal.IsLeftLessOrLessEqualOrEqual(),
		"greater": corecomparator.LeftGreater.IsLeftLessOrLessEqualOrEqual(),
	}
	expected := args.Map{"less": true, "le": true, "eq": true, "greater": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsLeftLessOrLessEqualOrEqual", actual)
}

func Test_Cov6_Compare_IsLeftGreaterOrGreaterEqualOrEqual(t *testing.T) {
	actual := args.Map{
		"greater": corecomparator.LeftGreater.IsLeftGreaterOrGreaterEqualOrEqual(),
		"ge":      corecomparator.LeftGreaterEqual.IsLeftGreaterOrGreaterEqualOrEqual(),
		"eq":      corecomparator.Equal.IsLeftGreaterOrGreaterEqualOrEqual(),
		"less":    corecomparator.LeftLess.IsLeftGreaterOrGreaterEqualOrEqual(),
	}
	expected := args.Map{"greater": true, "ge": true, "eq": true, "less": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsLeftGreaterOrGreaterEqualOrEqual", actual)
}

func Test_Cov6_Compare_IsNotEqualLogically(t *testing.T) {
	actual := args.Map{
		"notEq": corecomparator.LeftGreater.IsNotEqualLogically(),
		"eq":    corecomparator.Equal.IsNotEqualLogically(),
	}
	expected := args.Map{"notEq": true, "eq": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsNotEqualLogically", actual)
}

func Test_Cov6_Compare_IsDefinedPlus(t *testing.T) {
	actual := args.Map{
		"match":     corecomparator.Equal.IsDefinedPlus(corecomparator.Equal),
		"incon":     corecomparator.Inconclusive.IsDefinedPlus(corecomparator.Equal),
		"noMatch":   corecomparator.Equal.IsDefinedPlus(corecomparator.LeftGreater),
	}
	expected := args.Map{"match": true, "incon": false, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsDefinedPlus", actual)
}

func Test_Cov6_Compare_IsInconclusiveOrNotEqual(t *testing.T) {
	actual := args.Map{
		"incon":  corecomparator.Inconclusive.IsInconclusiveOrNotEqual(),
		"notEq":  corecomparator.NotEqual.IsInconclusiveOrNotEqual(),
		"eq":     corecomparator.Equal.IsInconclusiveOrNotEqual(),
	}
	expected := args.Map{"incon": true, "notEq": true, "eq": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsInconclusiveOrNotEqual", actual)
}

func Test_Cov6_Compare_IsAnyOf(t *testing.T) {
	actual := args.Map{
		"match":  corecomparator.Equal.IsAnyOf(corecomparator.LeftGreater, corecomparator.Equal),
		"empty":  corecomparator.Equal.IsAnyOf(),
		"noMatch": corecomparator.Equal.IsAnyOf(corecomparator.LeftGreater),
	}
	expected := args.Map{"match": true, "empty": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsAnyOf", actual)
}

func Test_Cov6_Compare_NameValue(t *testing.T) {
	result := corecomparator.Equal.NameValue()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare NameValue -- not empty", actual)
}

func Test_Cov6_Compare_CsvString(t *testing.T) {
	result := corecomparator.Equal.CsvString(corecomparator.Equal, corecomparator.LeftGreater)
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- CsvString", actual)
}

func Test_Cov6_Compare_CsvStrings(t *testing.T) {
	result := corecomparator.Equal.CsvStrings(corecomparator.Equal)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- CsvStrings", actual)
}

func Test_Cov6_Compare_CsvStrings_Empty(t *testing.T) {
	result := corecomparator.Equal.CsvStrings()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Compare returns empty -- CsvStrings empty", actual)
}

func Test_Cov6_Compare_CsvString_Empty(t *testing.T) {
	result := corecomparator.Equal.CsvString()
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Compare returns empty -- CsvString empty", actual)
}

func Test_Cov6_Compare_OperatorSymbol(t *testing.T) {
	actual := args.Map{
		"eq": corecomparator.Equal.OperatorSymbol(),
		"gt": corecomparator.LeftGreater.OperatorSymbol(),
	}
	expected := args.Map{"eq": "=", "gt": ">"}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- OperatorSymbol", actual)
}

func Test_Cov6_Compare_OperatorShortForm(t *testing.T) {
	actual := args.Map{
		"eq": corecomparator.Equal.OperatorShortForm(),
		"gt": corecomparator.LeftGreater.OperatorShortForm(),
	}
	expected := args.Map{"eq": "eq", "gt": "gt"}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- OperatorShortForm", actual)
}

func Test_Cov6_Compare_SqlOperatorSymbol(t *testing.T) {
	actual := args.Map{
		"eq":    corecomparator.Equal.SqlOperatorSymbol(),
		"notEq": corecomparator.NotEqual.SqlOperatorSymbol(),
	}
	expected := args.Map{"eq": "=", "notEq": "<>"}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- SqlOperatorSymbol", actual)
}

func Test_Cov6_Compare_NumberJsonString(t *testing.T) {
	actual := args.Map{"result": corecomparator.Equal.NumberJsonString()}
	expected := args.Map{"result": "\"0\""}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- NumberJsonString", actual)
}

func Test_Cov6_Compare_IsAnyNamesOf(t *testing.T) {
	actual := args.Map{
		"match":   corecomparator.Equal.IsAnyNamesOf("Equal", "NotEqual"),
		"noMatch": corecomparator.Equal.IsAnyNamesOf("LeftGreater"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsAnyNamesOf", actual)
}

func Test_Cov6_Compare_ValueAccessors(t *testing.T) {
	actual := args.Map{
		"valueByte":  corecomparator.Equal.ValueByte(),
		"valueInt":   corecomparator.Equal.ValueInt(),
		"valueInt8":  corecomparator.Equal.ValueInt8(),
		"valueInt16": corecomparator.Equal.ValueInt16(),
		"valueInt32": corecomparator.Equal.ValueInt32(),
		"valueStr":   corecomparator.Equal.ValueString(),
	}
	expected := args.Map{
		"valueByte": byte(0), "valueInt": 0, "valueInt8": int8(0),
		"valueInt16": int16(0), "valueInt32": int32(0), "valueStr": "0",
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- value accessors", actual)
}

func Test_Cov6_Compare_Value(t *testing.T) {
	actual := args.Map{"result": corecomparator.LeftGreater.Value()}
	expected := args.Map{"result": byte(1)}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- Value", actual)
}

func Test_Cov6_Compare_IsCompareEqualLogically(t *testing.T) {
	actual := args.Map{
		"same":    corecomparator.Equal.IsCompareEqualLogically(corecomparator.Equal),
		"neLogic": corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.NotEqual),
		"geLogic": corecomparator.Equal.IsCompareEqualLogically(corecomparator.LeftGreaterEqual),
		"leLogic": corecomparator.Equal.IsCompareEqualLogically(corecomparator.LeftLessEqual),
	}
	expected := args.Map{"same": true, "neLogic": true, "geLogic": true, "leLogic": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsCompareEqualLogically", actual)
}

func Test_Cov6_Compare_OnlySupportedErr(t *testing.T) {
	err := corecomparator.Inconclusive.OnlySupportedErr("msg", corecomparator.Equal)
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare OnlySupportedErr -- error", actual)
}

func Test_Cov6_Compare_OnlySupportedErr_NoMsg(t *testing.T) {
	err := corecomparator.Inconclusive.OnlySupportedErr("", corecomparator.Equal)
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare OnlySupportedErr no msg -- error", actual)
}

func Test_Cov6_Compare_OnlySupportedErr_Matching(t *testing.T) {
	err := corecomparator.Equal.OnlySupportedErr("msg", corecomparator.Equal)
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Compare OnlySupportedErr matching -- nil", actual)
}

func Test_Cov6_Compare_MarshalUnmarshalJSON(t *testing.T) {
	data, _ := corecomparator.Equal.MarshalJSON()
	actual := args.Map{"notEmpty": len(data) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- MarshalJSON", actual)
}

func Test_Cov6_Compare_UnmarshalJSON_Valid(t *testing.T) {
	var c corecomparator.Compare
	err := c.UnmarshalJSON([]byte("Equal"))
	actual := args.Map{"isNil": err == nil, "value": c.Name()}
	expected := args.Map{"isNil": true, "value": "Equal"}
	expected.ShouldBeEqual(t, 0, "Compare returns non-empty -- UnmarshalJSON valid", actual)
}

func Test_Cov6_Compare_UnmarshalJSON_Nil(t *testing.T) {
	var c corecomparator.Compare
	err := c.UnmarshalJSON(nil)
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare UnmarshalJSON nil -- error", actual)
}

func Test_Cov6_Compare_UnmarshalJSON_Invalid(t *testing.T) {
	var c corecomparator.Compare
	err := c.UnmarshalJSON([]byte("invalid"))
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare UnmarshalJSON invalid -- error", actual)
}

// ── Min / Max / MinLength / RangeNamesCsv / Ranges ──

func Test_Cov6_Min(t *testing.T) {
	actual := args.Map{"result": corecomparator.Min().Name()}
	expected := args.Map{"result": "Equal"}
	expected.ShouldBeEqual(t, 0, "Min -- Equal", actual)
}

func Test_Cov6_Max(t *testing.T) {
	actual := args.Map{"result": corecomparator.Max().Name()}
	expected := args.Map{"result": "NotEqual"}
	expected.ShouldBeEqual(t, 0, "Max -- NotEqual", actual)
}

func Test_Cov6_MinLength(t *testing.T) {
	actual := args.Map{
		"leftSmaller": corecomparator.MinLength(2, 5),
		"rightSmaller": corecomparator.MinLength(5, 3),
	}
	expected := args.Map{"leftSmaller": 2, "rightSmaller": 3}
	expected.ShouldBeEqual(t, 0, "MinLength returns correct value -- with args", actual)
}

func Test_Cov6_RangeNamesCsv(t *testing.T) {
	result := corecomparator.RangeNamesCsv()
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv -- not empty", actual)
}

func Test_Cov6_Ranges(t *testing.T) {
	result := corecomparator.Ranges()
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "Ranges -- 7 items", actual)
}

// ── BaseIsCaseSensitive / BaseIsIgnoreCase ──

func Test_Cov6_BaseIsCaseSensitive(t *testing.T) {
	b := corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}
	actual := args.Map{
		"isIgnoreCase": b.IsIgnoreCase(),
		"toIgnore":     b.BaseIsIgnoreCase().IsIgnoreCase,
	}
	expected := args.Map{"isIgnoreCase": false, "toIgnore": false}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns correct value -- with args", actual)
}

func Test_Cov6_BaseIsCaseSensitive_Clone(t *testing.T) {
	b := corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}
	cloned := b.Clone()
	actual := args.Map{"isCaseSensitive": cloned.IsCaseSensitive}
	expected := args.Map{"isCaseSensitive": true}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns correct value -- Clone", actual)
}

func Test_Cov6_BaseIsCaseSensitive_ClonePtr(t *testing.T) {
	b := &corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}
	cloned := b.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "val": cloned.IsCaseSensitive}
	expected := args.Map{"notNil": true, "val": true}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns correct value -- ClonePtr", actual)
}

func Test_Cov6_BaseIsCaseSensitive_ClonePtr_Nil(t *testing.T) {
	var b *corecomparator.BaseIsCaseSensitive
	cloned := b.ClonePtr()
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns nil -- ClonePtr nil", actual)
}

func Test_Cov6_BaseIsIgnoreCase(t *testing.T) {
	b := corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}
	actual := args.Map{
		"isCaseSensitive": b.IsCaseSensitive(),
		"toSensitive":     b.BaseIsCaseSensitive().IsCaseSensitive,
	}
	expected := args.Map{"isCaseSensitive": false, "toSensitive": false}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns correct value -- with args", actual)
}

func Test_Cov6_BaseIsIgnoreCase_Clone(t *testing.T) {
	b := corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}
	cloned := b.Clone()
	actual := args.Map{"isIgnoreCase": cloned.IsIgnoreCase}
	expected := args.Map{"isIgnoreCase": true}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns correct value -- Clone", actual)
}

func Test_Cov6_BaseIsIgnoreCase_ClonePtr(t *testing.T) {
	b := &corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}
	cloned := b.ClonePtr()
	actual := args.Map{"notNil": cloned != nil, "val": cloned.IsIgnoreCase}
	expected := args.Map{"notNil": true, "val": true}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns correct value -- ClonePtr", actual)
}

func Test_Cov6_BaseIsIgnoreCase_ClonePtr_Nil(t *testing.T) {
	var b *corecomparator.BaseIsIgnoreCase
	cloned := b.ClonePtr()
	actual := args.Map{"isNil": cloned == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns nil -- ClonePtr nil", actual)
}
