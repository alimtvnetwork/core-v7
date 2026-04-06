package stringcompareastests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// Variant — enum methods
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_SC01_Variant_Value(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{"val": int(v.Value())}
	expected := args.Map{"val": 0}
	expected.ShouldBeEqual(t, 0, "Variant_Value returns correct value -- with args", actual)
}

func Test_I5_SC02_Variant_IsAnyMethod(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"matchEqual":   v.IsAnyMethod("Equal"),
		"noMatchOther": v.IsAnyMethod("NotEqual"),
	}
	expected := args.Map{"matchEqual": true, "noMatchOther": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsAnyMethod returns correct value -- with args", actual)
}

func Test_I5_SC03_Variant_AllNameValues(t *testing.T) {
	v := stringcompareas.Equal
	names := v.AllNameValues()
	actual := args.Map{"hasItems": len(names) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant_AllNameValues returns non-empty -- with args", actual)
}

func Test_I5_SC04_Variant_OnlySupportedErr(t *testing.T) {
	v := stringcompareas.Equal
	err := v.OnlySupportedErr("Equal", "StartsWith")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variant_OnlySupportedErr returns error -- with args", actual)
}

func Test_I5_SC05_Variant_OnlySupportedMsgErr(t *testing.T) {
	v := stringcompareas.Equal
	err := v.OnlySupportedMsgErr("test msg", "Equal")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "Variant_OnlySupportedMsgErr returns error -- with args", actual)
}

func Test_I5_SC06_Variant_ValueUInt16(t *testing.T) {
	v := stringcompareas.StartsWith
	actual := args.Map{"val": int(v.ValueUInt16())}
	expected := args.Map{"val": 1}
	expected.ShouldBeEqual(t, 0, "Variant_ValueUInt16 returns correct value -- with args", actual)
}

func Test_I5_SC07_Variant_IntegerEnumRanges(t *testing.T) {
	v := stringcompareas.Equal
	ranges := v.IntegerEnumRanges()
	actual := args.Map{"hasItems": len(ranges) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant_IntegerEnumRanges returns correct value -- with args", actual)
}

func Test_I5_SC08_Variant_MinMaxAny(t *testing.T) {
	v := stringcompareas.Equal
	min, max := v.MinMaxAny()
	actual := args.Map{"minNotNil": min != nil, "maxNotNil": max != nil}
	expected := args.Map{"minNotNil": true, "maxNotNil": true}
	expected.ShouldBeEqual(t, 0, "Variant_MinMaxAny returns correct value -- with args", actual)
}

func Test_I5_SC09_Variant_MinMaxValueString(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"minNotEmpty": v.MinValueString() != "",
		"maxNotEmpty": v.MaxValueString() != "",
	}
	expected := args.Map{"minNotEmpty": true, "maxNotEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant_MinMaxValueString returns non-empty -- with args", actual)
}

func Test_I5_SC10_Variant_MaxInt_MinInt(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{"maxGe0": v.MaxInt() >= 0, "minGe0": v.MinInt() >= 0}
	expected := args.Map{"maxGe0": true, "minGe0": true}
	expected.ShouldBeEqual(t, 0, "Variant_MaxInt_MinInt returns correct value -- with args", actual)
}

func Test_I5_SC11_Variant_RangesDynamicMap(t *testing.T) {
	v := stringcompareas.Equal
	m := v.RangesDynamicMap()
	actual := args.Map{"hasItems": len(m) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "Variant_RangesDynamicMap returns correct value -- with args", actual)
}

func Test_I5_SC12_Variant_IsByteValueEqual(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"matchSelf":  v.IsByteValueEqual(byte(stringcompareas.Equal)),
		"noMatchOth": v.IsByteValueEqual(byte(stringcompareas.Regex)),
	}
	expected := args.Map{"matchSelf": true, "noMatchOth": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsByteValueEqual returns correct value -- with args", actual)
}

func Test_I5_SC13_Variant_Format(t *testing.T) {
	v := stringcompareas.Equal
	s := v.Format("{name}")
	actual := args.Map{"notEmpty": s != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant_Format returns correct value -- with args", actual)
}

func Test_I5_SC14_Variant_IsEnumEqual(t *testing.T) {
	v := stringcompareas.Equal
	v2 := stringcompareas.Equal
	v3 := stringcompareas.Regex
	actual := args.Map{
		"same": v.IsEnumEqual(&v2),
		"diff": v.IsEnumEqual(&v3),
	}
	expected := args.Map{"same": true, "diff": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsEnumEqual returns correct value -- with args", actual)
}

func Test_I5_SC15_Variant_IsAnyEnumsEqual(t *testing.T) {
	v := stringcompareas.Equal
	v2 := stringcompareas.Regex
	v3 := stringcompareas.Equal
	actual := args.Map{
		"matchAny":   v.IsAnyEnumsEqual(&v2, &v3),
		"noMatchAny": v.IsAnyEnumsEqual(&v2),
	}
	expected := args.Map{"matchAny": true, "noMatchAny": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsAnyEnumsEqual returns correct value -- with args", actual)
}

func Test_I5_SC16_Variant_IsNameEqual(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{"match": v.IsNameEqual("Equal"), "noMatch": v.IsNameEqual("Regex")}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsNameEqual returns correct value -- with args", actual)
}

func Test_I5_SC17_Variant_IsAnyNamesOf(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"match":   v.IsAnyNamesOf("Regex", "Equal"),
		"noMatch": v.IsAnyNamesOf("Regex", "StartsWith"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsAnyNamesOf returns correct value -- with args", actual)
}

func Test_I5_SC18_Variant_IsValueEqual(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"match":   v.IsValueEqual(byte(stringcompareas.Equal)),
		"noMatch": v.IsValueEqual(byte(stringcompareas.Regex)),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsValueEqual returns correct value -- with args", actual)
}

func Test_I5_SC19_Variant_IsAnyValuesEqual(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"match":   v.IsAnyValuesEqual(byte(stringcompareas.Regex), byte(stringcompareas.Equal)),
		"noMatch": v.IsAnyValuesEqual(byte(stringcompareas.Regex)),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Variant_IsAnyValuesEqual returns non-empty -- with args", actual)
}

func Test_I5_SC20_Variant_ValueInt_Variants(t *testing.T) {
	v := stringcompareas.StartsWith
	actual := args.Map{
		"int":   v.ValueInt(),
		"int8":  int(v.ValueInt8()),
		"int16": int(v.ValueInt16()),
		"int32": int(v.ValueInt32()),
	}
	expected := args.Map{"int": 1, "int8": 1, "int16": 1, "int32": 1}
	expected.ShouldBeEqual(t, 0, "Variant_ValueInt_Variants returns correct value -- with args", actual)
}

func Test_I5_SC21_Variant_ValueString(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{"notEmpty": v.ValueString() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant_ValueString returns non-empty -- with args", actual)
}

func Test_I5_SC22_Variant_IsValid_IsInvalid(t *testing.T) {
	v := stringcompareas.Equal
	inv := stringcompareas.Invalid
	actual := args.Map{
		"valid":      v.IsValid(),
		"notInvalid": !v.IsInvalid(),
		"invalid":    inv.IsInvalid(),
		"notValid":   !inv.IsValid(),
	}
	expected := args.Map{"valid": true, "notInvalid": true, "invalid": true, "notValid": true}
	expected.ShouldBeEqual(t, 0, "Variant_IsValid_IsInvalid returns error -- with args", actual)
}

func Test_I5_SC23_Variant_Name_NameValue_TypeName(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"name":         v.Name(),
		"nameValue":    v.NameValue() != "",
		"typeName":     v.TypeName() != "",
		"toString":     v.ToNumberString() != "",
		"stringMethod": v.String(),
	}
	expected := args.Map{
		"name":         "Equal",
		"nameValue":    true,
		"typeName":     true,
		"toString":     true,
		"stringMethod": "Equal",
	}
	expected.ShouldBeEqual(t, 0, "Variant_Name_NameValue_TypeName returns correct value -- with args", actual)
}

func Test_I5_SC24_Variant_Is(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{"isSelf": v.Is(stringcompareas.Equal), "isOther": v.Is(stringcompareas.Regex)}
	expected := args.Map{"isSelf": true, "isOther": false}
	expected.ShouldBeEqual(t, 0, "Variant_Is returns correct value -- with args", actual)
}

func Test_I5_SC25_Variant_IsBoolean_Methods(t *testing.T) {
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
	expected.ShouldBeEqual(t, 0, "Variant_IsBoolean_Methods returns correct value -- with args", actual)
}

func Test_I5_SC26_Variant_IsNegativeCondition(t *testing.T) {
	actual := args.Map{
		"notEqual":      stringcompareas.NotEqual.IsNegativeCondition(),
		"notStartsWith": stringcompareas.NotStartsWith.IsNegativeCondition(),
		"notEndsWith":   stringcompareas.NotEndsWith.IsNegativeCondition(),
		"notContains":   stringcompareas.NotContains.IsNegativeCondition(),
		"notAnyChars":   stringcompareas.NotAnyChars.IsNegativeCondition(),
		"notMatchRegex": stringcompareas.NotMatchRegex.IsNegativeCondition(),
		"nonGlob":       stringcompareas.NonGlob.IsNegativeCondition(),
		"equal":         stringcompareas.Equal.IsNegativeCondition(),
	}
	expected := args.Map{
		"notEqual": true, "notStartsWith": true, "notEndsWith": true,
		"notContains": true, "notAnyChars": true, "notMatchRegex": true,
		"nonGlob": true, "equal": false,
	}
	expected.ShouldBeEqual(t, 0, "Variant_IsNegativeCondition returns correct value -- with args", actual)
}

func Test_I5_SC27_Variant_MarshalJSON(t *testing.T) {
	v := stringcompareas.Equal
	b, err := v.MarshalJSON()
	actual := args.Map{"hasBytes": len(b) > 0, "nilErr": err == nil}
	expected := args.Map{"hasBytes": true, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Variant_MarshalJSON returns correct value -- with args", actual)
}

func Test_I5_SC28_Variant_UnmarshalJSON(t *testing.T) {
	v := stringcompareas.Invalid
	err := v.UnmarshalJSON([]byte(`"Equal"`))
	actual := args.Map{"nilErr": err == nil, "isEqual": v.IsEqual()}
	expected := args.Map{"nilErr": true, "isEqual": true}
	expected.ShouldBeEqual(t, 0, "Variant_UnmarshalJSON returns correct value -- with args", actual)
}

func Test_I5_SC29_Variant_RangeNamesCsv(t *testing.T) {
	v := stringcompareas.Equal
	csv := v.RangeNamesCsv()
	actual := args.Map{"notEmpty": csv != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Variant_RangeNamesCsv returns correct value -- with args", actual)
}

func Test_I5_SC30_Variant_MaxByte_MinByte_ValueByte_RangesByte(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"maxGe0":     int(v.MaxByte()) >= 0,
		"minGe0":     int(v.MinByte()) >= 0,
		"valueByte0": int(v.ValueByte()) == 0,
		"rangesLen":  len(v.RangesByte()) > 0,
	}
	expected := args.Map{"maxGe0": true, "minGe0": true, "valueByte0": true, "rangesLen": true}
	expected.ShouldBeEqual(t, 0, "Variant_MaxByte_MinByte_ValueByte_RangesByte returns correct value -- with args", actual)
}

func Test_I5_SC31_Variant_UnmarshallEnumToValue(t *testing.T) {
	v := stringcompareas.Equal
	val, err := v.UnmarshallEnumToValue([]byte(`"Equal"`))
	actual := args.Map{"val": int(val), "nilErr": err == nil}
	expected := args.Map{"val": 0, "nilErr": true}
	expected.ShouldBeEqual(t, 0, "Variant_UnmarshallEnumToValue returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsLineCompareFunc, DynamicCompare, IsCompareSuccess
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_SC32_Variant_IsLineCompareFunc(t *testing.T) {
	fn := stringcompareas.Equal.IsLineCompareFunc()
	actual := args.Map{"notNil": fn != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Variant_IsLineCompareFunc returns correct value -- with args", actual)
}

func Test_I5_SC33_Variant_DynamicCompare(t *testing.T) {
	v := stringcompareas.Equal
	fn := func(index int, content string, compareAs stringcompareas.Variant) bool {
		return compareAs == stringcompareas.Equal && content == "test"
	}
	actual := args.Map{"result": v.DynamicCompare(fn, 0, "test")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Variant_DynamicCompare returns correct value -- with args", actual)
}

func Test_I5_SC34_Variant_IsCompareSuccess_AllVariants(t *testing.T) {
	actual := args.Map{
		"equal":         stringcompareas.Equal.IsCompareSuccess(false, "hello", "hello"),
		"startsWith":    stringcompareas.StartsWith.IsCompareSuccess(false, "hello world", "hello"),
		"endsWith":      stringcompareas.EndsWith.IsCompareSuccess(false, "hello world", "world"),
		"anywhere":      stringcompareas.Anywhere.IsCompareSuccess(false, "hello world", "lo wo"),
		"contains":      stringcompareas.Contains.IsCompareSuccess(false, "hello world", "lo wo"),
		"anyChars":      stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "h"),
		"regex":         stringcompareas.Regex.IsCompareSuccess(false, "hello", "^hel"),
		"notEqual":      stringcompareas.NotEqual.IsCompareSuccess(false, "hello", "world"),
		"notStartsWith": stringcompareas.NotStartsWith.IsCompareSuccess(false, "hello", "world"),
		"notEndsWith":   stringcompareas.NotEndsWith.IsCompareSuccess(false, "hello", "world"),
		"notContains":   stringcompareas.NotContains.IsCompareSuccess(false, "hello", "xyz"),
		"notAnyChars":   stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "xyz"),
		"notMatchRegex": stringcompareas.NotMatchRegex.IsCompareSuccess(false, "hello", "^xyz"),
		"glob":          stringcompareas.Glob.IsCompareSuccess(false, "hello.txt", "*.txt"),
		"nonGlob":       stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.csv"),
	}
	expected := args.Map{
		"equal": true, "startsWith": true, "endsWith": true,
		"anywhere": true, "contains": true, "anyChars": true,
		"regex": true, "notEqual": true, "notStartsWith": true,
		"notEndsWith": true, "notContains": true, "notAnyChars": true,
		"notMatchRegex": true, "glob": true, "nonGlob": true,
	}
	expected.ShouldBeEqual(t, 0, "IsCompareSuccess_AllVariants returns correct value -- with args", actual)
}

func Test_I5_SC35_Variant_IsCompareSuccess_IgnoreCase(t *testing.T) {
	actual := args.Map{
		"equalIgnore":     stringcompareas.Equal.IsCompareSuccess(true, "Hello", "hello"),
		"startsWithIgn":   stringcompareas.StartsWith.IsCompareSuccess(true, "Hello World", "hello"),
		"endsWithIgn":     stringcompareas.EndsWith.IsCompareSuccess(true, "Hello World", "WORLD"),
		"anywhereIgn":     stringcompareas.Anywhere.IsCompareSuccess(true, "Hello World", "LO WO"),
		"anyCharsIgn":     stringcompareas.AnyChars.IsCompareSuccess(true, "Hello", "H"),
		"notAnyCharsIgn":  stringcompareas.NotAnyChars.IsCompareSuccess(true, "hello", "XYZ"),
		"notContainsIgn":  stringcompareas.NotContains.IsCompareSuccess(true, "hello", "XYZ"),
		"globIgn":         stringcompareas.Glob.IsCompareSuccess(true, "Hello.TXT", "*.txt"),
		"nonGlobIgn":      stringcompareas.NonGlob.IsCompareSuccess(true, "Hello.TXT", "*.csv"),
	}
	expected := args.Map{
		"equalIgnore": true, "startsWithIgn": true, "endsWithIgn": true,
		"anywhereIgn": true, "anyCharsIgn": true, "notAnyCharsIgn": true,
		"notContainsIgn": true, "globIgn": true, "nonGlobIgn": true,
	}
	expected.ShouldBeEqual(t, 0, "IsCompareSuccess_IgnoreCase returns correct value -- with args", actual)
}

func Test_I5_SC36_Variant_IsCompareSuccessCaseSensitive(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"match":   v.IsCompareSuccessCaseSensitive("hello", "hello"),
		"noMatch": v.IsCompareSuccessCaseSensitive("Hello", "hello"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsCompareSuccessCaseSensitive returns correct value -- with args", actual)
}

func Test_I5_SC37_Variant_IsCompareSuccessNonCaseSensitive(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{"match": v.IsCompareSuccessNonCaseSensitive("Hello", "hello")}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsCompareSuccessNonCaseSensitive returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// VerifyMessage, VerifyError, VerifyMessageCaseSensitive, VerifyErrorCaseSensitive
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_SC38_Variant_VerifyMessage_Match(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessage(false, "hello", "hello")
	actual := args.Map{"empty": msg == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage_Match returns correct value -- with args", actual)
}

func Test_I5_SC39_Variant_VerifyMessage_Mismatch_Positive(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessage(false, "hello", "world")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage_Mismatch_Positive returns correct value -- with args", actual)
}

func Test_I5_SC40_Variant_VerifyMessage_Mismatch_IgnoreCase(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessage(true, "Hello", "world")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage_Mismatch_IgnoreCase returns correct value -- with args", actual)
}

func Test_I5_SC41_Variant_VerifyMessage_Mismatch_Negative(t *testing.T) {
	msg := stringcompareas.NotEqual.VerifyMessage(false, "hello", "hello")
	actual := args.Map{"notEmpty": msg != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage_Mismatch_Negative returns correct value -- with args", actual)
}

func Test_I5_SC42_Variant_VerifyError_Match(t *testing.T) {
	err := stringcompareas.Equal.VerifyError(false, "hello", "hello")
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError_Match returns error -- with args", actual)
}

func Test_I5_SC43_Variant_VerifyError_Mismatch(t *testing.T) {
	err := stringcompareas.Equal.VerifyError(false, "hello", "world")
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError_Mismatch returns error -- with args", actual)
}

func Test_I5_SC44_Variant_VerifyMessageCaseSensitive(t *testing.T) {
	msg := stringcompareas.Equal.VerifyMessageCaseSensitive("hello", "hello")
	actual := args.Map{"empty": msg == ""}
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessageCaseSensitive returns correct value -- with args", actual)
}

func Test_I5_SC45_Variant_VerifyErrorCaseSensitive(t *testing.T) {
	err := stringcompareas.Equal.VerifyErrorCaseSensitive("hello", "hello")
	actual := args.Map{"nilErr": err == nil}
	expected := args.Map{"nilErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorCaseSensitive returns error -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// EnumType, As* interface casts, ToPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_SC46_Variant_EnumType(t *testing.T) {
	v := stringcompareas.Equal
	et := v.EnumType()
	actual := args.Map{"notNil": et != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "Variant_EnumType returns correct value -- with args", actual)
}

func Test_I5_SC47_Variant_AsInterfaces(t *testing.T) {
	v := stringcompareas.Equal
	actual := args.Map{
		"basicEnumContracts":     v.AsBasicEnumContractsBinder() != nil,
		"stringCompareTyper":     v.AsStringCompareTyper() != nil,
		"basicByteEnumContracts": v.AsBasicByteEnumContractsBinder() != nil,
	}
	expected := args.Map{
		"basicEnumContracts":     true,
		"stringCompareTyper":     true,
		"basicByteEnumContracts": true,
	}
	expected.ShouldBeEqual(t, 0, "Variant_AsInterfaces returns correct value -- with args", actual)
}

func Test_I5_SC48_Variant_ToPtr(t *testing.T) {
	v := stringcompareas.Equal
	p := v.ToPtr()
	actual := args.Map{"notNil": p != nil, "isEqual": p.IsEqual()}
	expected := args.Map{"notNil": true, "isEqual": true}
	expected.ShouldBeEqual(t, 0, "Variant_ToPtr returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Compare funcs — AnyChars, NotAnyChars, Glob, NonGlob edge cases
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_SC49_AnyChars_CaseSensitive(t *testing.T) {
	actual := args.Map{
		"match":   stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "h"),
		"noMatch": stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "H"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "AnyChars_CaseSensitive returns correct value -- with args", actual)
}

func Test_I5_SC50_NotAnyChars_CaseSensitive(t *testing.T) {
	actual := args.Map{
		"notFound": stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "xyz"),
		"found":    stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "h"),
	}
	expected := args.Map{"notFound": true, "found": false}
	expected.ShouldBeEqual(t, 0, "NotAnyChars_CaseSensitive returns correct value -- with args", actual)
}

func Test_I5_SC51_Glob_InvalidPattern(t *testing.T) {
	// Invalid glob pattern
	actual := args.Map{"result": stringcompareas.Glob.IsCompareSuccess(false, "hello", "[")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Glob_InvalidPattern returns error -- with args", actual)
}

func Test_I5_SC52_Glob_CaseSensitive(t *testing.T) {
	actual := args.Map{
		"match":   stringcompareas.Glob.IsCompareSuccess(false, "hello.txt", "*.txt"),
		"noMatch": stringcompareas.Glob.IsCompareSuccess(false, "Hello.TXT", "*.txt"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "Glob_CaseSensitive returns correct value -- with args", actual)
}

func Test_I5_SC53_NonGlob_CaseSensitive(t *testing.T) {
	actual := args.Map{
		"noMatch": stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.csv"),
		"match":   stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.txt"),
	}
	expected := args.Map{"noMatch": true, "match": false}
	expected.ShouldBeEqual(t, 0, "NonGlob_CaseSensitive returns correct value -- with args", actual)
}

func Test_I5_SC54_NonGlob_IgnoreCase(t *testing.T) {
	actual := args.Map{
		"noMatch": stringcompareas.NonGlob.IsCompareSuccess(true, "Hello.TXT", "*.csv"),
	}
	expected := args.Map{"noMatch": true}
	expected.ShouldBeEqual(t, 0, "NonGlob_IgnoreCase returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BasicEnumImpl var
// ══════════════════════════════════════════════════════════════════════════════

func Test_I5_SC55_BasicEnumImpl(t *testing.T) {
	actual := args.Map{"notNil": stringcompareas.BasicEnumImpl != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BasicEnumImpl returns correct value -- with args", actual)
}
