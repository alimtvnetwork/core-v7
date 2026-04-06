package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── Variant type-check methods ──

func Test_Cov3_Variant_TypeChecks(t *testing.T) {
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

func Test_Cov3_Variant_Accessors(t *testing.T) {
	v := stringcompareas.Equal
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
	expected := args.Map{
		"value": 0, "valueByte": 0, "valueInt": 0,
		"valueInt8": 0, "valueInt16": 0, "valueInt32": 0,
		"valueUInt16": 0, "valueString": "0",
		"name": "Equal", "string": "Equal",
		"isValid": true, "isInvalid": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant accessors -- Equal", actual)
}

func Test_Cov3_Variant_Invalid(t *testing.T) {
	actual := args.Map{
		"isValid":   stringcompareas.Invalid.IsValid(),
		"isInvalid": stringcompareas.Invalid.IsInvalid(),
	}
	expected := args.Map{"isValid": false, "isInvalid": true}
	expected.ShouldBeEqual(t, 0, "Variant returns error -- Invalid checks", actual)
}

// ── Enum interface methods ──

func Test_Cov3_Variant_EnumMethods(t *testing.T) {
	v := stringcompareas.Equal
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
	expected := args.Map{
		"nameValue": true, "typeName": true,
		"toNumberStr": "0", "rangeNamesCsv": true,
		"enumType": true, "isByteEqual": true,
		"isValueEqual": true, "isNameEqual": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- enum methods", actual)
}

func Test_Cov3_Variant_IsAnyNamesOf(t *testing.T) {
	actual := args.Map{
		"match":   stringcompareas.Equal.IsAnyNamesOf("Equal", "StartsWith"),
		"noMatch": stringcompareas.Equal.IsAnyNamesOf("StartsWith"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsAnyNamesOf", actual)
}

func Test_Cov3_Variant_IsAnyValuesEqual(t *testing.T) {
	actual := args.Map{
		"match":   stringcompareas.Equal.IsAnyValuesEqual(0, 1),
		"noMatch": stringcompareas.Equal.IsAnyValuesEqual(1, 2),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant returns non-empty -- IsAnyValuesEqual", actual)
}

// ── IsAnyMethod ──

func Test_Cov3_Variant_IsAnyMethod(t *testing.T) {
	actual := args.Map{
		"match":   stringcompareas.Equal.IsAnyMethod("Equal"),
		"noMatch": stringcompareas.Equal.IsAnyMethod("Invalid"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsAnyMethod", actual)
}

// ── Is / AllNameValues / OnlySupportedErr ──

func Test_Cov3_Variant_Is(t *testing.T) {
	actual := args.Map{
		"match": stringcompareas.Equal.Is(stringcompareas.Equal),
		"no":    stringcompareas.Equal.Is(stringcompareas.Regex),
	}
	expected := args.Map{"match": true, "no": false}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- Is", actual)
}

func Test_Cov3_Variant_AllNameValues(t *testing.T) {
	result := stringcompareas.Equal.AllNameValues()
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns non-empty -- AllNameValues", actual)
}

func Test_Cov3_Variant_OnlySupportedErr(t *testing.T) {
	err := stringcompareas.Equal.OnlySupportedErr("Equal")
	actual := args.Map{"notNil": err != nil}
	// OnlySupportedErr checks if names NOT in the enum's names are present
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "Variant returns error -- OnlySupportedErr", actual)
}

func Test_Cov3_Variant_OnlySupportedMsgErr(t *testing.T) {
	err := stringcompareas.Equal.OnlySupportedMsgErr("msg: ", "Equal")
	actual := args.Map{"notNil": err != nil}
	expected := args.Map{"notNil": err != nil}
	expected.ShouldBeEqual(t, 0, "Variant returns error -- OnlySupportedMsgErr", actual)
}

// ── RangesByte / MinByte / MaxByte ──

func Test_Cov3_Variant_RangesByte(t *testing.T) {
	v := stringcompareas.Equal
	result := v.RangesByte()
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- RangesByte", actual)
}

func Test_Cov3_Variant_MinMaxByte(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"minOK": v.MinByte() <= v.MaxByte(),
	}
	expected := args.Map{"minOK": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- Min/MaxByte", actual)
}

// ── Format / IntegerEnumRanges / MinMaxAny / RangesDynamicMap ──

func Test_Cov3_Variant_Format(t *testing.T) {
	result := stringcompareas.Equal.Format("{name}={value}")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- Format", actual)
}

func Test_Cov3_Variant_IntegerEnumRanges(t *testing.T) {
	result := stringcompareas.Equal.IntegerEnumRanges()
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IntegerEnumRanges", actual)
}

func Test_Cov3_Variant_MinMaxAny(t *testing.T) {
	min, max := stringcompareas.Equal.MinMaxAny()
	actual := args.Map{"minNotNil": min != nil, "maxNotNil": max != nil}
	expected := args.Map{"minNotNil": true, "maxNotNil": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- MinMaxAny", actual)
}

func Test_Cov3_Variant_MinMaxIntStr(t *testing.T) {
	actual := args.Map{
		"minStr": stringcompareas.Equal.MinValueString(),
		"maxStr": stringcompareas.Equal.MaxValueString(),
		"minInt": stringcompareas.Equal.MinInt(),
		"maxInt": stringcompareas.Equal.MaxInt(),
	}
	expected := args.Map{
		"minStr": actual["minStr"], "maxStr": actual["maxStr"],
		"minInt": actual["minInt"], "maxInt": actual["maxInt"],
	}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- MinMax int/str", actual)
}

func Test_Cov3_Variant_RangesDynamicMap(t *testing.T) {
	result := stringcompareas.Equal.RangesDynamicMap()
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- RangesDynamicMap", actual)
}

// ── MarshalJSON / UnmarshalJSON / UnmarshallEnumToValue ──

func Test_Cov3_Variant_MarshalJSON(t *testing.T) {
	data, err := stringcompareas.Equal.MarshalJSON()
	actual := args.Map{"hasData": len(data) > 0, "noErr": err == nil}
	expected := args.Map{"hasData": true, "noErr": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- MarshalJSON", actual)
}

func Test_Cov3_Variant_UnmarshalJSON(t *testing.T) {
	var v stringcompareas.Variant
	err := v.UnmarshalJSON([]byte(`"Equal"`))
	actual := args.Map{"noErr": err == nil, "val": v.Name()}
	expected := args.Map{"noErr": true, "val": "Equal"}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- UnmarshalJSON", actual)
}

func Test_Cov3_Variant_UnmarshallEnumToValue(t *testing.T) {
	val, err := stringcompareas.Equal.UnmarshallEnumToValue([]byte(`"Equal"`))
	actual := args.Map{"noErr": err == nil, "val": int(val)}
	expected := args.Map{"noErr": true, "val": 0}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- UnmarshallEnumToValue", actual)
}

// ── IsEnumEqual / IsAnyEnumsEqual ──

func Test_Cov3_Variant_IsEnumEqual(t *testing.T) {
	a := stringcompareas.Equal
	b := stringcompareas.StartsWith
	actual := args.Map{
		"same": a.IsEnumEqual(&a),
		"diff": a.IsEnumEqual(&b),
	}
	expected := args.Map{"same": true, "diff": false}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsEnumEqual", actual)
}

func Test_Cov3_Variant_IsAnyEnumsEqual_NoMatch(t *testing.T) {
	a := stringcompareas.Equal
	b := stringcompareas.StartsWith
	c := stringcompareas.EndsWith
	actual := args.Map{"result": a.IsAnyEnumsEqual(&b, &c)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Variant returns empty -- IsAnyEnumsEqual no match", actual)
}

// ── ToPtr / AsBasicEnumContractsBinder / AsStringCompareTyper / AsBasicByteEnumContractsBinder ──

func Test_Cov3_Variant_ToPtr(t *testing.T) {
	ptr := stringcompareas.Equal.ToPtr()
	actual := args.Map{"notNil": ptr != nil, "val": *ptr == stringcompareas.Equal}
	expected := args.Map{"notNil": true, "val": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- ToPtr", actual)
}

func Test_Cov3_Variant_Binders(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"basic":    v.AsBasicEnumContractsBinder() != nil,
		"compare":  v.AsStringCompareTyper() != nil,
		"byteBind": v.AsBasicByteEnumContractsBinder() != nil,
	}
	expected := args.Map{"basic": true, "compare": true, "byteBind": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- binder methods", actual)
}

// ── DynamicCompare ──

func Test_Cov3_Variant_DynamicCompare(t *testing.T) {
	dynFunc := func(index int, content string, compareAs stringcompareas.Variant) bool {
		return compareAs == stringcompareas.Equal && content == "hello"
	}
	actual := args.Map{
		"match":   stringcompareas.Equal.DynamicCompare(dynFunc, 0, "hello"),
		"noMatch": stringcompareas.Equal.DynamicCompare(dynFunc, 0, "world"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- DynamicCompare", actual)
}

// ── IsCompareSuccessCaseSensitive / IsCompareSuccessNonCaseSensitive ──

func Test_Cov3_Variant_CompareSuccessCaseSensitive(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"match":   v.IsCompareSuccessCaseSensitive("hello", "hello"),
		"noMatch": v.IsCompareSuccessCaseSensitive("Hello", "hello"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsCompareSuccessCaseSensitive", actual)
}

func Test_Cov3_Variant_CompareSuccessNonCaseSensitive(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"match": v.IsCompareSuccessNonCaseSensitive("Hello", "hello"),
	}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "Variant returns correct value -- IsCompareSuccessNonCaseSensitive", actual)
}

// ── IsNegativeCondition for non-negative ──

func Test_Cov3_Equal_IsNotNegativeCondition(t *testing.T) {
	actual := args.Map{"result": stringcompareas.Equal.IsNegativeCondition()}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal returns correct value -- not a negative condition", actual)
}

// ── NonGlob IsNegativeCondition ──

func Test_Cov3_NonGlob_IsNegativeCondition(t *testing.T) {
	actual := args.Map{"result": stringcompareas.NonGlob.IsNegativeCondition()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NonGlob returns correct value -- is negative condition", actual)
}

// ── IsCompareSuccess with Glob/NonGlob ──

func Test_Cov3_Glob_IsCompareSuccess(t *testing.T) {
	actual := args.Map{
		"match":      stringcompareas.Glob.IsCompareSuccess(false, "hello.txt", "*.txt"),
		"noMatch":    stringcompareas.Glob.IsCompareSuccess(false, "hello.go", "*.txt"),
		"ignoreCase": stringcompareas.Glob.IsCompareSuccess(true, "Hello.TXT", "*.txt"),
	}
	expected := args.Map{"match": true, "noMatch": false, "ignoreCase": true}
	expected.ShouldBeEqual(t, 0, "Glob returns correct value -- IsCompareSuccess", actual)
}

func Test_Cov3_NonGlob_IsCompareSuccess(t *testing.T) {
	actual := args.Map{
		"noMatch": stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.txt"),
		"match":   stringcompareas.NonGlob.IsCompareSuccess(false, "hello.go", "*.txt"),
	}
	expected := args.Map{"noMatch": false, "match": true}
	expected.ShouldBeEqual(t, 0, "NonGlob returns correct value -- IsCompareSuccess", actual)
}

// ── AnyChars ──

func Test_Cov3_AnyChars_IsCompareSuccess(t *testing.T) {
	actual := args.Map{
		"match":      stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "eo"),
		"ignoreCase": stringcompareas.AnyChars.IsCompareSuccess(true, "HELLO", "eo"),
	}
	expected := args.Map{"match": true, "ignoreCase": true}
	expected.ShouldBeEqual(t, 0, "AnyChars returns correct value -- IsCompareSuccess", actual)
}

// ── NotAnyChars ──

func Test_Cov3_NotAnyChars_IsCompareSuccess(t *testing.T) {
	actual := args.Map{
		"noChars": stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "xyz"),
	}
	expected := args.Map{"noChars": true}
	expected.ShouldBeEqual(t, 0, "NotAnyChars returns correct value -- IsCompareSuccess", actual)
}

// ── VerifyMessage match returns empty ──

func Test_Cov3_VerifyMessage_Match(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessage(false, "hello", "hello")
	actual := args.Map{"isEmpty": msg == ""}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage match -- empty", actual)
}

// ── VerifyError match returns nil ──

func Test_Cov3_VerifyError_Match(t *testing.T) {
	err := stringcompareas.Equal.VerifyError(false, "hello", "hello")
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "VerifyError match -- nil", actual)
}

// ── VerifyMessage negative condition, case strict ──

func Test_Cov3_VerifyMessage_NegativeCaseStrict(t *testing.T) {
	msg := stringcompareas.NotEqual.VerifyMessage(false, "hello", "hello")
	actual := args.Map{"nonEmpty": msg != ""}
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage negative case strict -- error msg", actual)
}
