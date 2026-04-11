package stringcompareastests

import (
	"encoding/json"
	"testing"

	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Variant_Name_Verification(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.Name() != "Equal"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Equal", actual)
	actual = args.Map{"result": stringcompareas.StartsWith.Name() != "StartsWith"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected StartsWith", actual)
	actual = args.Map{"result": stringcompareas.EndsWith.Name() != "EndsWith"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected EndsWith", actual)
	actual = args.Map{"result": stringcompareas.Anywhere.Name() != "Anywhere"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Anywhere", actual)
	actual = args.Map{"result": stringcompareas.Regex.Name() != "Regex"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected Regex", actual)
}

func Test_Variant_Is_Methods(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should be IsEqual", actual)
	actual = args.Map{"result": stringcompareas.StartsWith.IsStartsWith()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StartsWith should be IsStartsWith", actual)
	actual = args.Map{"result": stringcompareas.EndsWith.IsEndsWith()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "EndsWith should be IsEndsWith", actual)
	actual = args.Map{"result": stringcompareas.Anywhere.IsAnywhere()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Anywhere should be IsAnywhere", actual)
	actual = args.Map{"result": stringcompareas.Contains.IsContains()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Contains should be IsContains", actual)
	actual = args.Map{"result": stringcompareas.AnyChars.IsAnyChars()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyChars should be IsAnyChars", actual)
	actual = args.Map{"result": stringcompareas.Regex.IsRegex()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Regex should be IsRegex", actual)
	actual = args.Map{"result": stringcompareas.Glob.IsGlob()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Glob should be IsGlob", actual)
	actual = args.Map{"result": stringcompareas.NonGlob.IsNonGlob()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NonGlob should be IsNonGlob", actual)
}

func Test_Variant_Not_Methods(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NotEqual.IsNotEqual()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEqual should be IsNotEqual", actual)
	actual = args.Map{"result": stringcompareas.NotStartsWith.IsNotStartsWith()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotStartsWith should be IsNotStartsWith", actual)
	actual = args.Map{"result": stringcompareas.NotEndsWith.IsNotEndsWith()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEndsWith should be IsNotEndsWith", actual)
	actual = args.Map{"result": stringcompareas.NotContains.IsNotContains()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotContains should be IsNotContains", actual)
	actual = args.Map{"result": stringcompareas.NotMatchRegex.IsNotMatchRegex()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotMatchRegex should be IsNotMatchRegex", actual)
}

func Test_Variant_IsNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NotEqual.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEqual should be negative", actual)
	actual = args.Map{"result": stringcompareas.NotStartsWith.IsNegativeCondition()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotStartsWith should be negative", actual)
	actual = args.Map{"result": stringcompareas.NonGlob.IsNegativeCondition()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NonGlob should be negative", actual)
	actual = args.Map{"result": stringcompareas.Equal.IsNegativeCondition()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not be negative", actual)
}

func Test_Variant_ValidInvalid(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsValid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should be valid", actual)
	actual = args.Map{"result": stringcompareas.Equal.IsInvalid()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not be invalid", actual)
	actual = args.Map{"result": stringcompareas.Invalid.IsInvalid()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Invalid should be invalid", actual)
}

func Test_Variant_IsCompareSuccess(t *testing.T) {
	// Equal
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsCompareSuccess(false, "hello", "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should match", actual)
	actual = args.Map{"result": stringcompareas.Equal.IsCompareSuccess(false, "hello", "world")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not match different", actual)

	// StartsWith
	actual = args.Map{"result": stringcompareas.StartsWith.IsCompareSuccess(false, "hello world", "hello")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "StartsWith should match", actual)

	// EndsWith
	actual = args.Map{"result": stringcompareas.EndsWith.IsCompareSuccess(false, "hello world", "world")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "EndsWith should match", actual)

	// Anywhere/Contains
	actual = args.Map{"result": stringcompareas.Anywhere.IsCompareSuccess(false, "hello world", "lo wo")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Anywhere should match", actual)

	// NotEqual
	actual = args.Map{"result": stringcompareas.NotEqual.IsCompareSuccess(false, "hello", "world")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEqual should match different strings", actual)

	// Case insensitive
	actual = args.Map{"result": stringcompareas.Equal.IsCompareSuccess(true, "Hello", "hello")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal ignore case should match", actual)
}

func Test_Variant_CompareSuccessCaseSensitive(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{"result": v.IsCompareSuccessCaseSensitive("hello", "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "case sensitive should match", actual)
	actual = args.Map{"result": v.IsCompareSuccessCaseSensitive("Hello", "hello")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "case sensitive should not match different case", actual)
}

func Test_Variant_CompareSuccessNonCaseSensitive(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{"result": v.IsCompareSuccessNonCaseSensitive("Hello", "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "case insensitive should match", actual)
}

func Test_Variant_VerifyMessage(t *testing.T) {
	// Arrange
	msg := stringcompareas.Equal.VerifyMessage(false, "hello", "hello")

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should return empty message", actual)

	msg = stringcompareas.Equal.VerifyMessage(false, "hello", "world")
	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-matching should return error message", actual)

	// negative case
	msg = stringcompareas.NotEqual.VerifyMessage(false, "hello", "hello")
	actual = args.Map{"result": msg == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqual same values should return error message", actual)
}

func Test_Variant_VerifyError(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.VerifyError(false, "hello", "hello")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should return nil error", actual)

	err = stringcompareas.Equal.VerifyError(false, "hello", "world")
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "non-matching should return error", actual)
}

func Test_Variant_VerifyMessageCaseSensitive(t *testing.T) {
	// Arrange
	msg := stringcompareas.Equal.VerifyMessageCaseSensitive("hello", "hello")

	// Act
	actual := args.Map{"result": msg != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should return empty", actual)
}

func Test_Variant_VerifyErrorCaseSensitive(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.VerifyErrorCaseSensitive("hello", "hello")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "matching should return nil", actual)
}

func Test_Variant_MarshalJSON(t *testing.T) {
	// Arrange
	bytes, err := stringcompareas.Equal.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON error:", actual)
	actual = args.Map{"result": len(bytes) == 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MarshalJSON should return bytes", actual)
}

func Test_Variant_UnmarshalJSON(t *testing.T) {
	// Arrange
	v := stringcompareas.Invalid
	data, _ := json.Marshal("Equal")
	err := v.UnmarshalJSON(data)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshalJSON error:", actual)
}

func Test_Variant_ValueMethods(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{"result": v.Value() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should be 0", actual)
	actual = args.Map{"result": v.ValueInt() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt should be 0", actual)
	actual = args.Map{"result": v.ValueInt8() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt8 should be 0", actual)
	actual = args.Map{"result": v.ValueInt16() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt16 should be 0", actual)
	actual = args.Map{"result": v.ValueInt32() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueInt32 should be 0", actual)
	actual = args.Map{"result": v.ValueUInt16() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueUInt16 should be 0", actual)
}

func Test_Variant_String(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.String() != "Equal"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should be Equal", actual)
}

func Test_Variant_Is(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.Is(stringcompareas.Equal)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal should Is Equal", actual)
	actual = args.Map{"result": stringcompareas.Equal.Is(stringcompareas.StartsWith)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Equal should not Is StartsWith", actual)
}

func Test_Variant_NameValue(t *testing.T) {
	// Arrange
	nv := stringcompareas.Equal.NameValue()

	// Act
	actual := args.Map{"result": nv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NameValue should not be empty", actual)
}

func Test_Variant_ValueString(t *testing.T) {
	// Arrange
	vs := stringcompareas.Equal.ValueString()

	// Act
	actual := args.Map{"result": vs == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ValueString should not be empty", actual)
}

func Test_Variant_ToNumberString(t *testing.T) {
	// Arrange
	ns := stringcompareas.Equal.ToNumberString()

	// Act
	actual := args.Map{"result": ns == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToNumberString should not be empty", actual)
}

func Test_Variant_RangeNamesCsv(t *testing.T) {
	// Arrange
	csv := stringcompareas.Equal.RangeNamesCsv()

	// Act
	actual := args.Map{"result": csv == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsv should not be empty", actual)
}

func Test_Variant_IsAnyMethod(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsAnyMethod("Equal")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match Equal", actual)
	actual = args.Map{"result": stringcompareas.Equal.IsAnyMethod("StartsWith")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match StartsWith", actual)
}

func Test_Variant_IsNameEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsNameEqual("Equal")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should be name equal", actual)
}

func Test_Variant_IsAnyNamesOf(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsAnyNamesOf("Equal", "StartsWith")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match Equal", actual)
	actual = args.Map{"result": stringcompareas.Equal.IsAnyNamesOf("StartsWith", "EndsWith")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match", actual)
}

func Test_Variant_IsValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsValueEqual(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal value should be 0", actual)
}

func Test_Variant_IsAnyValuesEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsAnyValuesEqual(0, 1)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match 0", actual)
	actual = args.Map{"result": stringcompareas.Equal.IsAnyValuesEqual(1, 2)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not match", actual)
}

func Test_Variant_IsByteValueEqual(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.IsByteValueEqual(0)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match", actual)
}

func Test_Variant_MaxMinByte(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal

	// Act
	actual := args.Map{"result": v.MaxByte() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "max should not be 0", actual)
	actual = args.Map{"result": v.MinByte() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "min should be 0", actual)
}

func Test_Variant_RangesByte(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal
	ranges := v.RangesByte()

	// Act
	actual := args.Map{"result": len(ranges) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ranges should not be empty", actual)
}

func Test_Variant_AllNameValues(t *testing.T) {
	// Arrange
	nv := stringcompareas.Equal.AllNameValues()

	// Act
	actual := args.Map{"result": len(nv) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllNameValues should not be empty", actual)
}

func Test_Variant_IntegerEnumRanges(t *testing.T) {
	// Arrange
	r := stringcompareas.Equal.IntegerEnumRanges()

	// Act
	actual := args.Map{"result": len(r) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IntegerEnumRanges should not be empty", actual)
}

func Test_Variant_MinMaxAny(t *testing.T) {
	// Arrange
	min, max := stringcompareas.Equal.MinMaxAny()

	// Act
	actual := args.Map{"result": min == nil || max == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinMaxAny should not be nil", actual)
}

func Test_Variant_MinMaxValueString(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.MinValueString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MinValueString should not be empty", actual)
	actual = args.Map{"result": stringcompareas.Equal.MaxValueString() == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxValueString should not be empty", actual)
}

func Test_Variant_MaxMinInt(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.MaxInt() == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaxInt should not be 0", actual)
}

func Test_Variant_RangesDynamicMap(t *testing.T) {
	// Arrange
	m := stringcompareas.Equal.RangesDynamicMap()

	// Act
	actual := args.Map{"result": len(m) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RangesDynamicMap should not be empty", actual)
}

func Test_Variant_Format(t *testing.T) {
	// Arrange
	f := stringcompareas.Equal.Format("type: %s")

	// Act
	actual := args.Map{"result": f == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Format should not be empty", actual)
}

func Test_Variant_ToPtr(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal
	ptr := v.ToPtr()

	// Act
	actual := args.Map{"result": ptr == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToPtr should not be nil", actual)
}

func Test_Variant_TypeName(t *testing.T) {
	// Arrange
	tn := stringcompareas.Equal.TypeName()

	// Act
	actual := args.Map{"result": tn == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "TypeName should not be empty", actual)
}

func Test_Variant_EnumType(t *testing.T) {
	// Arrange
	et := stringcompareas.Equal.EnumType()

	// Act
	actual := args.Map{"result": et == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "EnumType should not be nil", actual)
}

func Test_Variant_AsInterfaces(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Equal.AsBasicEnumContractsBinder() == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsBasicEnumContractsBinder should not be nil", actual)
	actual = args.Map{"result": stringcompareas.Equal.AsStringCompareTyper() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsStringCompareTyper should not be nil", actual)
	actual = args.Map{"result": stringcompareas.Equal.AsBasicByteEnumContractsBinder() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AsBasicByteEnumContractsBinder should not be nil", actual)
}

func Test_Variant_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	data, _ := json.Marshal("Equal")
	val, err := stringcompareas.Equal.UnmarshallEnumToValue(data)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UnmarshallEnumToValue error:", actual)
	actual = args.Map{"result": val != byte(stringcompareas.Equal)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should return Equal value", actual)
}

func Test_Variant_IsEnumEqual(t *testing.T) {
	// Arrange
	a := stringcompareas.Equal
	b := stringcompareas.Equal

	// Act
	actual := args.Map{"result": a.IsEnumEqual(&b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "same enum should be equal", actual)
}

func Test_Variant_IsAnyEnumsEqual(t *testing.T) {
	// Arrange
	a := stringcompareas.Equal
	b := stringcompareas.StartsWith
	c := stringcompareas.Equal

	// Act
	actual := args.Map{"result": a.IsAnyEnumsEqual(&b, &c)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should match Equal", actual)
}

func Test_Variant_OnlySupportedErr(t *testing.T) {
	// Arrange
	// Passing all names as supported → no unsupported → nil error
	allNames := []string{
		"Equal", "StartsWith", "EndsWith", "Anywhere",
		"IsContains", "AnyChars", "Regex",
		"NotEqual", "NotStartsWith", "NotEndsWith",
		"NotContains", "NotAnyChars", "NotMatchRegex",
		"Glob", "NonGlob", "Invalid",
	}
	err := stringcompareas.Equal.OnlySupportedErr(allNames...)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "all names supported should not error, got:", actual)

	// Passing only "Equal" → all others unsupported → error expected
	err2 := stringcompareas.Equal.OnlySupportedErr("Equal")
	actual = args.Map{"result": err2 == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "partial support should return error for unsupported names", actual)
}

func Test_Variant_DynamicCompare(t *testing.T) {
	// Arrange
	v := stringcompareas.Equal
	dynFunc := func(index int, content string, compareAs stringcompareas.Variant) bool {
		return compareAs == stringcompareas.Equal && content == "hello"
	}

	// Act
	actual := args.Map{"result": v.DynamicCompare(dynFunc, 0, "hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "dynamic compare should return true", actual)
}

func Test_Variant_IsLineCompareFunc(t *testing.T) {
	// Arrange
	fn := stringcompareas.Equal.IsLineCompareFunc()

	// Act
	actual := args.Map{"result": fn == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsLineCompareFunc should not be nil", actual)
	actual = args.Map{"result": fn("hello", "hello", false)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal compare func should match same strings", actual)
}

func Test_Variant_GlobCompare(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Glob.IsCompareSuccess(false, "hello.txt", "*.txt")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Glob should match *.txt", actual)
	actual = args.Map{"result": stringcompareas.NonGlob.IsCompareSuccess(false, "hello.txt", "*.txt")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NonGlob should not match *.txt", actual)
}

func Test_Variant_AnyCharsCompare(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.AnyChars.IsCompareSuccess(false, "hello", "hlo")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyChars should match when chars exist", actual)
	actual = args.Map{"result": stringcompareas.NotAnyChars.IsCompareSuccess(false, "hello", "hlo")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotAnyChars should not match when chars exist", actual)
}

func Test_Variant_RegexCompare(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Regex.IsCompareSuccess(false, "hello123", `\d+`)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Regex should match digits", actual)
}
