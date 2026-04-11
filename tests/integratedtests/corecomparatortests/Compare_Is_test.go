package corecomparatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── Compare methods ──

func Test_Compare_Is(t *testing.T) {
	// Act
	actual := args.Map{
		"same": corecomparator.Equal.Is(corecomparator.Equal),
		"diff": corecomparator.Equal.Is(corecomparator.LeftGreater),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare.Is returns correct value -- with args", actual)
}

func Test_Compare_LessEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"less":     corecomparator.LeftLess.IsLessEqual(),
		"equal":    corecomparator.Equal.IsLessEqual(),
		"greater":  corecomparator.LeftGreater.IsLessEqual(),
	}

	// Assert
	expected := args.Map{
		"less": true,
		"equal": true,
		"greater": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- LessEqual", actual)
}

func Test_Compare_GreaterEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"greater": corecomparator.LeftGreater.IsGreaterEqual(),
		"equal":   corecomparator.Equal.IsGreaterEqual(),
		"less":    corecomparator.LeftLess.IsGreaterEqual(),
	}

	// Assert
	expected := args.Map{
		"greater": true,
		"equal": true,
		"less": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- GreaterEqual", actual)
}

func Test_Compare_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   corecomparator.Equal.IsNameEqual("Equal"),
		"noMatch": corecomparator.Equal.IsNameEqual("NotEqual"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsNameEqual", actual)
}

func Test_Compare_ToNumberString(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Equal.ToNumberString()}

	// Assert
	expected := args.Map{"result": "0"}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- ToNumberString", actual)
}

func Test_Compare_IsLeftLessOrLessEqualOrEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"less":    corecomparator.LeftLess.IsLeftLessOrLessEqualOrEqual(),
		"le":      corecomparator.LeftLessEqual.IsLeftLessOrLessEqualOrEqual(),
		"eq":      corecomparator.Equal.IsLeftLessOrLessEqualOrEqual(),
		"greater": corecomparator.LeftGreater.IsLeftLessOrLessEqualOrEqual(),
	}

	// Assert
	expected := args.Map{
		"less": true,
		"le": true,
		"eq": true,
		"greater": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsLeftLessOrLessEqualOrEqual", actual)
}

func Test_Compare_IsLeftGreaterOrGreaterEqualOrEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"greater": corecomparator.LeftGreater.IsLeftGreaterOrGreaterEqualOrEqual(),
		"ge":      corecomparator.LeftGreaterEqual.IsLeftGreaterOrGreaterEqualOrEqual(),
		"eq":      corecomparator.Equal.IsLeftGreaterOrGreaterEqualOrEqual(),
		"less":    corecomparator.LeftLess.IsLeftGreaterOrGreaterEqualOrEqual(),
	}

	// Assert
	expected := args.Map{
		"greater": true,
		"ge": true,
		"eq": true,
		"less": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsLeftGreaterOrGreaterEqualOrEqual", actual)
}

func Test_Compare_IsNotEqualLogically(t *testing.T) {
	// Act
	actual := args.Map{
		"notEq": corecomparator.LeftGreater.IsNotEqualLogically(),
		"eq":    corecomparator.Equal.IsNotEqualLogically(),
	}

	// Assert
	expected := args.Map{
		"notEq": true,
		"eq": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsNotEqualLogically", actual)
}

func Test_Compare_IsDefinedPlus(t *testing.T) {
	// Act
	actual := args.Map{
		"match":     corecomparator.Equal.IsDefinedPlus(corecomparator.Equal),
		"incon":     corecomparator.Inconclusive.IsDefinedPlus(corecomparator.Equal),
		"noMatch":   corecomparator.Equal.IsDefinedPlus(corecomparator.LeftGreater),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"incon": false,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsDefinedPlus", actual)
}

func Test_Compare_IsInconclusiveOrNotEqual(t *testing.T) {
	// Act
	actual := args.Map{
		"incon":  corecomparator.Inconclusive.IsInconclusiveOrNotEqual(),
		"notEq":  corecomparator.NotEqual.IsInconclusiveOrNotEqual(),
		"eq":     corecomparator.Equal.IsInconclusiveOrNotEqual(),
	}

	// Assert
	expected := args.Map{
		"incon": true,
		"notEq": true,
		"eq": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsInconclusiveOrNotEqual", actual)
}

func Test_Compare_IsAnyOf(t *testing.T) {
	// Act
	actual := args.Map{
		"match":  corecomparator.Equal.IsAnyOf(corecomparator.LeftGreater, corecomparator.Equal),
		"empty":  corecomparator.Equal.IsAnyOf(),
		"noMatch": corecomparator.Equal.IsAnyOf(corecomparator.LeftGreater),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"empty": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsAnyOf", actual)
}

func Test_Compare_NameValue(t *testing.T) {
	// Arrange
	result := corecomparator.Equal.NameValue()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare NameValue -- not empty", actual)
}

func Test_Compare_CsvString(t *testing.T) {
	// Arrange
	result := corecomparator.Equal.CsvString(corecomparator.Equal, corecomparator.LeftGreater)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- CsvString", actual)
}

func Test_Compare_CsvStrings(t *testing.T) {
	// Arrange
	result := corecomparator.Equal.CsvStrings(corecomparator.Equal)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- CsvStrings", actual)
}

func Test_Compare_CsvStrings_Empty(t *testing.T) {
	// Arrange
	result := corecomparator.Equal.CsvStrings()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Compare returns empty -- CsvStrings empty", actual)
}

func Test_Compare_CsvString_Empty(t *testing.T) {
	// Arrange
	result := corecomparator.Equal.CsvString()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Compare returns empty -- CsvString empty", actual)
}

func Test_Compare_OperatorSymbol(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": corecomparator.Equal.OperatorSymbol(),
		"gt": corecomparator.LeftGreater.OperatorSymbol(),
	}

	// Assert
	expected := args.Map{
		"eq": "=",
		"gt": ">",
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- OperatorSymbol", actual)
}

func Test_Compare_OperatorShortForm(t *testing.T) {
	// Act
	actual := args.Map{
		"eq": corecomparator.Equal.OperatorShortForm(),
		"gt": corecomparator.LeftGreater.OperatorShortForm(),
	}

	// Assert
	expected := args.Map{
		"eq": "eq",
		"gt": "gt",
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- OperatorShortForm", actual)
}

func Test_Compare_SqlOperatorSymbol(t *testing.T) {
	// Act
	actual := args.Map{
		"eq":    corecomparator.Equal.SqlOperatorSymbol(),
		"notEq": corecomparator.NotEqual.SqlOperatorSymbol(),
	}

	// Assert
	expected := args.Map{
		"eq": "=",
		"notEq": "<>",
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- SqlOperatorSymbol", actual)
}

func Test_Compare_NumberJsonString(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Equal.NumberJsonString()}

	// Assert
	expected := args.Map{"result": "\"0\""}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- NumberJsonString", actual)
}

func Test_Compare_IsAnyNamesOf(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   corecomparator.Equal.IsAnyNamesOf("Equal", "NotEqual"),
		"noMatch": corecomparator.Equal.IsAnyNamesOf("LeftGreater"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsAnyNamesOf", actual)
}

func Test_Compare_ValueAccessors(t *testing.T) {
	// Act
	actual := args.Map{
		"valueByte":  corecomparator.Equal.ValueByte(),
		"valueInt":   corecomparator.Equal.ValueInt(),
		"valueInt8":  corecomparator.Equal.ValueInt8(),
		"valueInt16": corecomparator.Equal.ValueInt16(),
		"valueInt32": corecomparator.Equal.ValueInt32(),
		"valueStr":   corecomparator.Equal.ValueString(),
	}

	// Assert
	expected := args.Map{
		"valueByte": byte(0), "valueInt": 0, "valueInt8": int8(0),
		"valueInt16": int16(0), "valueInt32": int32(0), "valueStr": "0",
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- value accessors", actual)
}

func Test_Compare_Value(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.LeftGreater.Value()}

	// Assert
	expected := args.Map{"result": byte(1)}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- Value", actual)
}

func Test_Compare_IsCompareEqualLogically(t *testing.T) {
	// Act
	actual := args.Map{
		"same":    corecomparator.Equal.IsCompareEqualLogically(corecomparator.Equal),
		"neLogic": corecomparator.LeftGreater.IsCompareEqualLogically(corecomparator.NotEqual),
		"geLogic": corecomparator.Equal.IsCompareEqualLogically(corecomparator.LeftGreaterEqual),
		"leLogic": corecomparator.Equal.IsCompareEqualLogically(corecomparator.LeftLessEqual),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"neLogic": true,
		"geLogic": true,
		"leLogic": true,
	}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- IsCompareEqualLogically", actual)
}

func Test_Compare_OnlySupportedErr_FromCompareIs(t *testing.T) {
	// Arrange
	err := corecomparator.Inconclusive.OnlySupportedErr("msg", corecomparator.Equal)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare OnlySupportedErr -- error", actual)
}

func Test_Compare_OnlySupportedErr_NoMsg(t *testing.T) {
	// Arrange
	err := corecomparator.Inconclusive.OnlySupportedErr("", corecomparator.Equal)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare OnlySupportedErr no msg -- error", actual)
}

func Test_Compare_OnlySupportedErr_Matching(t *testing.T) {
	// Arrange
	err := corecomparator.Equal.OnlySupportedErr("msg", corecomparator.Equal)

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "Compare OnlySupportedErr matching -- nil", actual)
}

func Test_Compare_MarshalUnmarshalJSON(t *testing.T) {
	// Arrange
	data, _ := corecomparator.Equal.MarshalJSON()

	// Act
	actual := args.Map{"notEmpty": len(data) > 0}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Compare returns correct value -- MarshalJSON", actual)
}

func Test_Compare_UnmarshalJSON_Valid(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	err := c.UnmarshalJSON([]byte("Equal"))

	// Act
	actual := args.Map{
		"isNil": err == nil,
		"value": c.Name(),
	}

	// Assert
	expected := args.Map{
		"isNil": true,
		"value": "Equal",
	}
	expected.ShouldBeEqual(t, 0, "Compare returns non-empty -- UnmarshalJSON valid", actual)
}

func Test_Compare_UnmarshalJSON_Nil(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	err := c.UnmarshalJSON(nil)

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare UnmarshalJSON nil -- error", actual)
}

func Test_Compare_UnmarshalJSON_Invalid(t *testing.T) {
	// Arrange
	var c corecomparator.Compare
	err := c.UnmarshalJSON([]byte("invalid"))

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "Compare UnmarshalJSON invalid -- error", actual)
}

// ── Min / Max / MinLength / RangeNamesCsv / Ranges ──

func Test_Min(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Min().Name()}

	// Assert
	expected := args.Map{"result": "Equal"}
	expected.ShouldBeEqual(t, 0, "Min -- Equal", actual)
}

func Test_Max(t *testing.T) {
	// Act
	actual := args.Map{"result": corecomparator.Max().Name()}

	// Assert
	expected := args.Map{"result": "NotEqual"}
	expected.ShouldBeEqual(t, 0, "Max -- NotEqual", actual)
}

func Test_MinLength(t *testing.T) {
	// Act
	actual := args.Map{
		"leftSmaller": corecomparator.MinLength(2, 5),
		"rightSmaller": corecomparator.MinLength(5, 3),
	}

	// Assert
	expected := args.Map{
		"leftSmaller": 2,
		"rightSmaller": 3,
	}
	expected.ShouldBeEqual(t, 0, "MinLength returns correct value -- with args", actual)
}

func Test_RangeNamesCsv(t *testing.T) {
	// Arrange
	result := corecomparator.RangeNamesCsv()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv -- not empty", actual)
}

func Test_Ranges(t *testing.T) {
	// Arrange
	result := corecomparator.Ranges()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 7}
	expected.ShouldBeEqual(t, 0, "Ranges -- 7 items", actual)
}

// ── BaseIsCaseSensitive / BaseIsIgnoreCase ──

func Test_BaseIsCaseSensitive(t *testing.T) {
	// Arrange
	b := corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}

	// Act
	actual := args.Map{
		"isIgnoreCase": b.IsIgnoreCase(),
		"toIgnore":     b.BaseIsIgnoreCase().IsIgnoreCase,
	}

	// Assert
	expected := args.Map{
		"isIgnoreCase": false,
		"toIgnore": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns correct value -- with args", actual)
}

func Test_BaseIsCaseSensitive_Clone(t *testing.T) {
	// Arrange
	b := corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}
	cloned := b.Clone()

	// Act
	actual := args.Map{"isCaseSensitive": cloned.IsCaseSensitive}

	// Assert
	expected := args.Map{"isCaseSensitive": true}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns correct value -- Clone", actual)
}

func Test_BaseIsCaseSensitive_ClonePtr(t *testing.T) {
	// Arrange
	b := &corecomparator.BaseIsCaseSensitive{IsCaseSensitive: true}
	cloned := b.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"val": cloned.IsCaseSensitive,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns correct value -- ClonePtr", actual)
}

func Test_BaseIsCaseSensitive_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var b *corecomparator.BaseIsCaseSensitive
	cloned := b.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseIsCaseSensitive returns nil -- ClonePtr nil", actual)
}

func Test_BaseIsIgnoreCase(t *testing.T) {
	// Arrange
	b := corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}

	// Act
	actual := args.Map{
		"isCaseSensitive": b.IsCaseSensitive(),
		"toSensitive":     b.BaseIsCaseSensitive().IsCaseSensitive,
	}

	// Assert
	expected := args.Map{
		"isCaseSensitive": false,
		"toSensitive": false,
	}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns correct value -- with args", actual)
}

func Test_BaseIsIgnoreCase_Clone(t *testing.T) {
	// Arrange
	b := corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}
	cloned := b.Clone()

	// Act
	actual := args.Map{"isIgnoreCase": cloned.IsIgnoreCase}

	// Assert
	expected := args.Map{"isIgnoreCase": true}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns correct value -- Clone", actual)
}

func Test_BaseIsIgnoreCase_ClonePtr(t *testing.T) {
	// Arrange
	b := &corecomparator.BaseIsIgnoreCase{IsIgnoreCase: true}
	cloned := b.ClonePtr()

	// Act
	actual := args.Map{
		"notNil": cloned != nil,
		"val": cloned.IsIgnoreCase,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns correct value -- ClonePtr", actual)
}

func Test_BaseIsIgnoreCase_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var b *corecomparator.BaseIsIgnoreCase
	cloned := b.ClonePtr()

	// Act
	actual := args.Map{"isNil": cloned == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "BaseIsIgnoreCase returns nil -- ClonePtr nil", actual)
}
