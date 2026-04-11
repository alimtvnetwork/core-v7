package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── Variant type-check methods ──

func Test_Variant_TypeChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"isEqual":         stringcompareas.Equal.IsEqual(),
		"isStartsWith":    stringcompareas.StartsWith.IsStartsWith(),
		"isEndsWith":      stringcompareas.EndsWith.IsEndsWith(),
		"isAnywhere":      stringcompareas.Anywhere.IsAnywhere(),
		"isContains":      stringcompareas.Contains.IsContains(),
		"isAnyChars":      stringcompareas.AnyChars.IsAnyChars(),
		"isRegex":         stringcompareas.Regex.IsRegex(),
		"isNotEqual":      stringcompareas.NotEqual.IsNotEqual(),
		"isNotStartsWith": stringcompareas.NotStartsWith.IsNotStartsWith(),
		"isNotEndsWith":   stringcompareas.NotEndsWith.IsNotEndsWith(),
		"isNotContains":   stringcompareas.NotContains.IsNotContains(),
		"isNotMatchRegex": stringcompareas.NotMatchRegex.IsNotMatchRegex(),
		"isGlob":          stringcompareas.Glob.IsGlob(),
		"isNonGlob":       stringcompareas.NonGlob.IsNonGlob(),
	}

	// Assert
	expected := args.Map{
		"isEqual": true, "isStartsWith": true, "isEndsWith": true,
		"isAnywhere": true, "isContains": true, "isAnyChars": true,
		"isRegex": true, "isNotEqual": true, "isNotStartsWith": true,
		"isNotEndsWith": true, "isNotContains": true, "isNotMatchRegex": true,
		"isGlob": true, "isNonGlob": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant type checks -- all true", actual)
}

// ── Variant enum accessors ──

func Test_Variant_Accessors(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{
		"value":      int(v.Value()),
		"valueByte":  int(v.ValueByte()),
		"valueInt":   v.ValueInt(),
		"valueInt8":  int(v.ValueInt8()),
		"valueInt16": int(v.ValueInt16()),
		"valueInt32": int(v.ValueInt32()),
		"valueUInt16": int(v.ValueUInt16()),
		"valueString": v.ValueString(),
		"name":       v.Name(),
		"string":     v.String(),
		"isValid":    v.IsValid(),
		"isInvalid":  v.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"value": 0, "valueByte": 0, "valueInt": 0,
		"valueInt8": 0, "valueInt16": 0, "valueInt32": 0,
		"valueUInt16": 0, "valueString": "0",
		"name": "Equal", "string": "Equal",
		"isValid": true, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant accessors -- Equal", actual)
}

func Test_Variant_Invalid(t *testing.T) {
	// Act
	actual := args.Map{
		"isValid":   stringcompareas.Invalid.IsValid(),
		"isInvalid": stringcompareas.Invalid.IsInvalid(),
	}

	// Assert
	expected := args.Map{
		"isValid": false,
		"isInvalid": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns error -- Invalid checks", actual)
}

// ── Enum interface methods ──

func Test_Variant_EnumMethods(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{
		"nameValue":     v.NameValue() != "",
		"typeName":      v.TypeName() != "",
		"toNumberStr":   v.ToNumberString(),
		"rangeNamesCsv": v.RangeNamesCsv() != "",
		"enumType":      v.EnumType() != nil,
		"isByteEqual":   v.IsByteValueEqual(0),
		"isValueEqual":  v.IsValueEqual(0),
		"isNameEqual":   v.IsNameEqual("Equal"),
	}

	// Assert
	expected := args.Map{
		"nameValue": true, "typeName": true,
		"toNumberStr": "0", "rangeNamesCsv": true,
		"enumType": true, "isByteEqual": true,
		"isValueEqual": true, "isNameEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- enum methods", actual)
}

func Test_Variant_IsAnyNamesOf_FromVariantTypeChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringcompareas.Equal.IsAnyNamesOf("Equal", "StartsWith"),
		"noMatch": stringcompareas.Equal.IsAnyNamesOf("StartsWith"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsAnyNamesOf", actual)
}

func Test_Variant_IsAnyValuesEqual_FromVariantTypeChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringcompareas.Equal.IsAnyValuesEqual(0, 1),
		"noMatch": stringcompareas.Equal.IsAnyValuesEqual(1, 2),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns non-empty -- IsAnyValuesEqual", actual)
}

// ── IsAnyMethod ──

func Test_Variant_IsAnyMethod_FromVariantTypeChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringcompareas.Equal.IsAnyMethod("Equal"),
		"noMatch": stringcompareas.Equal.IsAnyMethod("Invalid"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsAnyMethod", actual)
}

// ── Is / AllNameValues / OnlySupportedErr ──

func Test_Variant_Is_FromVariantTypeChecks(t *testing.T) {
	// Act
	actual := args.Map{
		"match": stringcompareas.Equal.Is(stringcompareas.Equal),
		"no":    stringcompareas.Equal.Is(stringcompareas.Regex),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"no": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- Is", actual)
}

func Test_Variant_AllNameValues_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	result := stringcompareas.Equal.AllNameValues()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns non-empty -- AllNameValues", actual)
}

func Test_Variant_OnlySupportedErr_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.OnlySupportedErr("Equal")

	// Act
	actual := args.Map{"notNil": err != nil}
	// OnlySupportedErr checks if names NOT in the enum's names are present

	// Assert
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "Variant returns error -- OnlySupportedErr", actual)
}

func Test_Variant_OnlySupportedMsgErr(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.OnlySupportedMsgErr("msg: ", "Equal")

	// Act
	actual := args.Map{"notNil": err != nil}

	// Assert
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "Variant returns error -- OnlySupportedMsgErr", actual)
}

// ── RangesByte / MinByte / MaxByte ──

func Test_Variant_RangesByte_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal
	result := v.RangesByte()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- RangesByte", actual)
}

func Test_Variant_MinMaxByte(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{
		"minOK": v.MinByte() <= v.MaxByte(),
	}

	// Assert
	expected := args.Map{"minOK": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- Min/MaxByte", actual)
}

// ── Format / IntegerEnumRanges / MinMaxAny / RangesDynamicMap ──

func Test_Variant_Format_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	result := stringcompareas.Equal.Format("{name}={value}")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- Format", actual)
}

func Test_Variant_IntegerEnumRanges_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	result := stringcompareas.Equal.IntegerEnumRanges()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IntegerEnumRanges", actual)
}

func Test_Variant_MinMaxAny_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	min, max := stringcompareas.Equal.MinMaxAny()

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
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- MinMaxAny", actual)
}

func Test_Variant_MinMaxIntStr(t *testing.T) {
	// Act
	actual := args.Map{
		"minStr": stringcompareas.Equal.MinValueString(),
		"maxStr": stringcompareas.Equal.MaxValueString(),
		"minInt": stringcompareas.Equal.MinInt(),
		"maxInt": stringcompareas.Equal.MaxInt(),
	}

	// Assert
	expected := args.Map{
		"minStr": actual["minStr"], "maxStr": actual["maxStr"],
		"minInt": actual["minInt"], "maxInt": actual["maxInt"],
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- MinMax int/str", actual)
}

func Test_Variant_RangesDynamicMap_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	result := stringcompareas.Equal.RangesDynamicMap()

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- RangesDynamicMap", actual)
}

// ── MarshalJSON / UnmarshalJSON / UnmarshallEnumToValue ──

func Test_Variant_MarshalJSON_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	data, err := stringcompareas.Equal.MarshalJSON()

	// Act
	actual := args.Map{
		"hasData": len(data) > 0,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"hasData": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- MarshalJSON", actual)
}

func Test_Variant_UnmarshalJSON_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	var v stringcompareas.Variant
	err := v.UnmarshalJSON([]byte(`"Equal"`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": v.Name(),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": "Equal",
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- UnmarshalJSON", actual)
}

func Test_Variant_UnmarshallEnumToValue_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	val, err := stringcompareas.Equal.UnmarshallEnumToValue([]byte(`"Equal"`))

	// Act
	actual := args.Map{
		"noErr": err == nil,
		"val": int(val),
	}

	// Assert
	expected := args.Map{
		"noErr": true,
		"val": 0,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- UnmarshallEnumToValue", actual)
}

// ── IsEnumEqual / IsAnyEnumsEqual ──

func Test_Variant_IsEnumEqual_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	a := stringcompareas.Equal
	b := stringcompareas.StartsWith

	// Act
	actual := args.Map{
		"same": a.IsEnumEqual(&a),
		"diff": a.IsEnumEqual(&b),
	}

	// Assert
	expected := args.Map{
		"same": true,
		"diff": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsEnumEqual", actual)
}

func Test_Variant_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	// Arrange
	a := stringcompareas.Equal
	b := stringcompareas.StartsWith
	c := stringcompareas.EndsWith

	// Act
	actual := args.Map{"result": a.IsAnyEnumsEqual(&b, &c)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Variant returns empty -- IsAnyEnumsEqual no match", actual)
}

// ── ToPtr / AsBasicEnumContractsBinder / AsStringCompareTyper / AsBasicByteEnumContractsBinder ──

func Test_Variant_ToPtr_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	ptr := stringcompareas.Equal.ToPtr()

	// Act
	actual := args.Map{
		"notNil": ptr != nil,
		"val": *ptr == stringcompareas.Equal,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToPtr", actual)
}

func Test_Variant_Binders(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{
		"basic":    v.AsBasicEnumContractsBinder() != nil,
		"compare":  v.AsStringCompareTyper() != nil,
		"byteBind": v.AsBasicByteEnumContractsBinder() != nil,
	}

	// Assert
	expected := args.Map{
		"basic": true,
		"compare": true,
		"byteBind": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- binder methods", actual)
}

// ── DynamicCompare ──

func Test_Variant_DynamicCompare_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	dynFunc := func(index int, content string, compareAs stringcompareas.Variant) bool {
		return compareAs == stringcompareas.Equal && content == "hello"
	}

	// Act
	actual := args.Map{
		"match":   stringcompareas.Equal.DynamicCompare(dynFunc, 0, "hello"),
		"noMatch": stringcompareas.Equal.DynamicCompare(dynFunc, 0, "world"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- DynamicCompare", actual)
}

// ── IsCompareSuccessCaseSensitive / IsCompareSuccessNonCaseSensitive ──

func Test_Variant_CompareSuccessCaseSensitive_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{
		"match":   v.IsCompareSuccessCaseSensitive("hello", "hello"),
		"noMatch": v.IsCompareSuccessCaseSensitive("Hello", "hello"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsCompareSuccessCaseSensitive", actual)
}

func Test_Variant_CompareSuccessNonCaseSensitive_FromVariantTypeChecks(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{
		"match": v.IsCompareSuccessNonCaseSensitive("Hello", "hello"),
	}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsCompareSuccessNonCaseSensitive", actual)
}

// ── IsNegativeCondition for non-negative ──

func Test_Equal_IsNotNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal returns correct value -- not a negative condition", actual)
}

// ── NonGlob IsNegativeCondition ──

func Test_NonGlob_IsNegativeCondition_FromVariantTypeChecks(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NonGlob.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NonGlob returns correct value -- is negative condition", actual)
}

// ── IsCompareSuccess with Glob/NonGlob ──

func Test_Glob_IsCompareSuccess(t *testing.T) {
	// Act
	actual := args.Map{
		"match":      stringcompareas.Glob.IsCompareSuccess(false, "hello.txt", "*.txt"),
		"noMatch":    stringcompareas.Glob.IsCompareSuccess(false, "hello.go", "*.txt"),
		"ignoreCase": stringcompareas.Glob.IsCompareSuccess(true, "Hello.TXT", "*.txt"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
		"ignoreCase": true,
	}
	expected.ShouldBeEqual(t, 0, "Glob returns correct value -- IsCompareSuccess", actual)
}

func Test_NonGlob_IsCompareSuccess(t *testing.T) {
	// Act
	actual := args.Map{
		"noMatch": stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.txt"),
		"match":   stringcompareas.NonGlob.IsCompareSuccess(false, "hello.go", "*.txt"),
	}

	// Assert
	expected := args.Map{
		"noMatch": false,
		"match": true,
	}
	expected.ShouldBeEqual(t, 0, "NonGlob returns correct value -- IsCompareSuccess", actual)
}

// ── AnyChars ──

func Test_AnyChars_IsCompareSuccess(t *testing.T) {
	// Act
	actual := args.Map{
		"match":      stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "eo"),
		"ignoreCase": stringcompareas.AnyChars.IsCompareSuccess(true, "HELLO", "eo"),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"ignoreCase": true,
	}
	expected.ShouldBeEqual(t, 0, "AnyChars returns correct value -- IsCompareSuccess", actual)
}

// ── NotAnyChars ──

func Test_NotAnyChars_IsCompareSuccess(t *testing.T) {
	// Act
	actual := args.Map{
		"noChars": stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "xyz"),
	}

	// Assert
	expected := args.Map{"noChars": true}
	expected.ShouldBeEqual(t, 0, "NotAnyChars returns correct value -- IsCompareSuccess", actual)
}

// ── VerifyMessage match returns empty ──

func Test_VerifyMessage_Match(t *testing.T) {
	// Arrange
	msg := stringcompareas.Equal.VerifyMessage(false, "hello", "hello")

	// Act
	actual := args.Map{"isEmpty": msg == ""}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage match -- empty", actual)
}

// ── VerifyError match returns nil ──

func Test_VerifyError_Match(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.VerifyError(false, "hello", "hello")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "VerifyError match -- nil", actual)
}

// ── VerifyMessage negative condition, case strict ──

func Test_VerifyMessage_NegativeCaseStrict(t *testing.T) {
	// Arrange
	msg := stringcompareas.NotEqual.VerifyMessage(false, "hello", "hello")

	// Act
	actual := args.Map{"nonEmpty": msg != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage negative case strict -- error msg", actual)
}
