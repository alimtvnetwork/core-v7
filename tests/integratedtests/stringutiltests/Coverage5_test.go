package stringutiltests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ═══════════════════════════════════════════
// FirstChar / LastChar
// ═══════════════════════════════════════════

func Test_Cov5_FirstChar_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringutil.FirstChar("hello")}
	expected := args.Map{"result": byte('h')}
	expected.ShouldBeEqual(t, 0, "FirstChar returns h -- hello", actual)
}

func Test_Cov5_FirstChar_Empty(t *testing.T) {
	actual := args.Map{"result": stringutil.FirstChar("")}
	expected := args.Map{"result": byte(0)}
	expected.ShouldBeEqual(t, 0, "FirstChar returns 0 -- empty", actual)
}

func Test_Cov5_FirstCharOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringutil.FirstCharOrDefault("abc")}
	expected := args.Map{"result": byte('a')}
	expected.ShouldBeEqual(t, 0, "FirstCharOrDefault returns a -- abc", actual)
}

func Test_Cov5_FirstCharOrDefault_Empty(t *testing.T) {
	actual := args.Map{"result": stringutil.FirstCharOrDefault("")}
	expected := args.Map{"result": byte(0)}
	expected.ShouldBeEqual(t, 0, "FirstCharOrDefault returns 0 -- empty", actual)
}

func Test_Cov5_LastChar_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringutil.LastChar("hello")}
	expected := args.Map{"result": byte('o')}
	expected.ShouldBeEqual(t, 0, "LastChar returns o -- hello", actual)
}

func Test_Cov5_LastCharOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringutil.LastCharOrDefault("hello")}
	expected := args.Map{"result": byte('o')}
	expected.ShouldBeEqual(t, 0, "LastCharOrDefault returns o -- hello", actual)
}

func Test_Cov5_LastCharOrDefault_Empty(t *testing.T) {
	actual := args.Map{"result": stringutil.LastCharOrDefault("")}
	expected := args.Map{"result": byte(0)}
	expected.ShouldBeEqual(t, 0, "LastCharOrDefault returns 0 -- empty", actual)
}

// ═══════════════════════════════════════════
// IsBlank / IsBlankPtr
// ═══════════════════════════════════════════

func Test_Cov5_IsBlankPtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsBlankPtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns true -- nil", actual)
}

func Test_Cov5_IsBlankPtr_NonBlank(t *testing.T) {
	s := "hello"
	actual := args.Map{"result": stringutil.IsBlankPtr(&s)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns false -- non-blank", actual)
}

// ═══════════════════════════════════════════
// IsContains
// ═══════════════════════════════════════════

func Test_Cov5_IsContains_CaseSensitive_Found(t *testing.T) {
	actual := args.Map{"result": stringutil.IsContains([]string{"a", "b", "c"}, "b", 0, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContains returns true -- case sensitive found", actual)
}

func Test_Cov5_IsContains_CaseSensitive_NotFound(t *testing.T) {
	actual := args.Map{"result": stringutil.IsContains([]string{"a", "b"}, "B", 0, true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns false -- case sensitive not found", actual)
}

func Test_Cov5_IsContains_CaseInsensitive_Found(t *testing.T) {
	actual := args.Map{"result": stringutil.IsContains([]string{"a", "B"}, "b", 0, false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContains returns true -- case insensitive found", actual)
}

func Test_Cov5_IsContains_NilSlice(t *testing.T) {
	actual := args.Map{"result": stringutil.IsContains(nil, "a", 0, true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns false -- nil slice", actual)
}

func Test_Cov5_IsContains_EmptySlice(t *testing.T) {
	actual := args.Map{"result": stringutil.IsContains([]string{}, "a", 0, true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns false -- empty slice", actual)
}

// ═══════════════════════════════════════════
// IsContainsPtr / IsContainsPtrSimple
// ═══════════════════════════════════════════

func Test_Cov5_IsContainsPtr_Found(t *testing.T) {
	lines := []string{"hello", "world"}
	find := "world"
	actual := args.Map{"result": stringutil.IsContainsPtr(&lines, &find, 0, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns true -- found", actual)
}

func Test_Cov5_IsContainsPtr_NilLines(t *testing.T) {
	find := "x"
	actual := args.Map{"result": stringutil.IsContainsPtr(nil, &find, 0, true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns false -- nil lines", actual)
}

func Test_Cov5_IsContainsPtrSimple_Found(t *testing.T) {
	lines := []string{"hello", "world"}
	actual := args.Map{"result": stringutil.IsContainsPtrSimple(&lines, "world", 0, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns true -- found", actual)
}

func Test_Cov5_IsContainsPtrSimple_CaseInsensitive(t *testing.T) {
	lines := []string{"Hello"}
	actual := args.Map{"result": stringutil.IsContainsPtrSimple(&lines, "hello", 0, false)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns true -- case insensitive", actual)
}

// ═══════════════════════════════════════════
// IsEmpty / IsEmptyPtr / IsNotEmpty / IsDefined / IsDefinedPtr
// ═══════════════════════════════════════════

func Test_Cov5_IsEmptyPtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmptyPtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns true -- nil", actual)
}

func Test_Cov5_IsEmptyPtr_Empty(t *testing.T) {
	s := ""
	actual := args.Map{"result": stringutil.IsEmptyPtr(&s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns true -- empty string", actual)
}

func Test_Cov5_IsNotEmpty(t *testing.T) {
	actual := args.Map{
		"empty":    stringutil.IsNotEmpty(""),
		"nonEmpty": stringutil.IsNotEmpty("hello"),
	}
	expected := args.Map{"empty": false, "nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns bool -- both cases", actual)
}

func Test_Cov5_IsDefined(t *testing.T) {
	actual := args.Map{
		"defined":   stringutil.IsDefined("hello"),
		"empty":     stringutil.IsDefined(""),
		"space":     stringutil.IsDefined(" "),
		"whitespace": stringutil.IsDefined("\n"),
	}
	expected := args.Map{"defined": true, "empty": false, "space": false, "whitespace": false}
	expected.ShouldBeEqual(t, 0, "IsDefined returns bool -- various inputs", actual)
}

func Test_Cov5_IsDefinedPtr(t *testing.T) {
	s := "hello"
	actual := args.Map{
		"defined": stringutil.IsDefinedPtr(&s),
		"nil":     stringutil.IsDefinedPtr(nil),
	}
	expected := args.Map{"defined": true, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns bool -- both cases", actual)
}

func Test_Cov5_IsEmptyOrWhitespacePtr(t *testing.T) {
	s := " "
	actual := args.Map{
		"nil":   stringutil.IsEmptyOrWhitespacePtr(nil),
		"space": stringutil.IsEmptyOrWhitespacePtr(&s),
	}
	expected := args.Map{"nil": true, "space": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr returns true -- nil and space", actual)
}

func Test_Cov5_IsNullOrEmptyPtr(t *testing.T) {
	s := ""
	actual := args.Map{
		"nil":   stringutil.IsNullOrEmptyPtr(nil),
		"empty": stringutil.IsNullOrEmptyPtr(&s),
	}
	expected := args.Map{"nil": true, "empty": true}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns true -- nil and empty", actual)
}

// ═══════════════════════════════════════════
// IsStarts / IsEnds / IsStartsWith / IsEndsWith
// ═══════════════════════════════════════════

func Test_Cov5_IsStarts(t *testing.T) {
	actual := args.Map{
		"yes": stringutil.IsStarts("hello world", "hello"),
		"no":  stringutil.IsStarts("hello world", "world"),
	}
	expected := args.Map{"yes": true, "no": false}
	expected.ShouldBeEqual(t, 0, "IsStarts returns bool -- both cases", actual)
}

func Test_Cov5_IsEnds(t *testing.T) {
	actual := args.Map{
		"yes": stringutil.IsEnds("hello world", "world"),
		"no":  stringutil.IsEnds("hello world", "hello"),
	}
	expected := args.Map{"yes": true, "no": false}
	expected.ShouldBeEqual(t, 0, "IsEnds returns bool -- both cases", actual)
}

func Test_Cov5_IsStartsWith_IgnoreCase(t *testing.T) {
	actual := args.Map{
		"ignoreCase": stringutil.IsStartsWith("Hello World", "hello", true),
		"emptyStart": stringutil.IsStartsWith("hello", "", false),
		"emptyContent": stringutil.IsStartsWith("", "hello", false),
		"longer": stringutil.IsStartsWith("hi", "hello", false),
	}
	expected := args.Map{"ignoreCase": true, "emptyStart": true, "emptyContent": false, "longer": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns bool -- edge cases", actual)
}

func Test_Cov5_IsEndsWith_IgnoreCase(t *testing.T) {
	actual := args.Map{
		"ignoreCase":   stringutil.IsEndsWith("Hello WORLD", "world", true),
		"emptyEnd":     stringutil.IsEndsWith("hello", "", false),
		"emptyContent": stringutil.IsEndsWith("", "hello", false),
		"longer":       stringutil.IsEndsWith("hi", "hello", false),
	}
	expected := args.Map{"ignoreCase": true, "emptyEnd": true, "emptyContent": false, "longer": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns bool -- edge cases", actual)
}

// ═══════════════════════════════════════════
// IsAnyStartsWith / IsAnyEndsWith
// ═══════════════════════════════════════════

func Test_Cov5_IsAnyStartsWith_Found(t *testing.T) {
	actual := args.Map{
		"found":      stringutil.IsAnyStartsWith("hello world", false, "hello", "goodbye"),
		"notFound":   stringutil.IsAnyStartsWith("hello world", false, "goodbye", "bye"),
		"emptyBoth":  stringutil.IsAnyStartsWith("", false),
		"emptyTerms": stringutil.IsAnyStartsWith("hello", false),
	}
	expected := args.Map{"found": true, "notFound": false, "emptyBoth": true, "emptyTerms": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns bool -- various cases", actual)
}

func Test_Cov5_IsAnyEndsWith_Found(t *testing.T) {
	actual := args.Map{
		"found":      stringutil.IsAnyEndsWith("hello world", false, "world", "earth"),
		"notFound":   stringutil.IsAnyEndsWith("hello world", false, "earth", "mars"),
		"emptyBoth":  stringutil.IsAnyEndsWith("", false),
		"emptyTerms": stringutil.IsAnyEndsWith("hello", false),
	}
	expected := args.Map{"found": true, "notFound": false, "emptyBoth": true, "emptyTerms": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns bool -- various cases", actual)
}

// ═══════════════════════════════════════════
// IsStartsChar / IsEndsChar / IsStartsRune / IsEndsRune
// ═══════════════════════════════════════════

func Test_Cov5_IsStartsChar(t *testing.T) {
	actual := args.Map{
		"yes":   stringutil.IsStartsChar("hello", 'h'),
		"no":    stringutil.IsStartsChar("hello", 'x'),
		"empty": stringutil.IsStartsChar("", 'h'),
	}
	expected := args.Map{"yes": true, "no": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsChar returns bool -- various cases", actual)
}

func Test_Cov5_IsEndsChar(t *testing.T) {
	actual := args.Map{
		"yes":   stringutil.IsEndsChar("hello", 'o'),
		"no":    stringutil.IsEndsChar("hello", 'x'),
		"empty": stringutil.IsEndsChar("", 'o'),
	}
	expected := args.Map{"yes": true, "no": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns bool -- various cases", actual)
}

func Test_Cov5_IsStartsRune(t *testing.T) {
	actual := args.Map{
		"yes":   stringutil.IsStartsRune("hello", 'h'),
		"no":    stringutil.IsStartsRune("hello", 'x'),
		"empty": stringutil.IsStartsRune("", 'h'),
	}
	expected := args.Map{"yes": true, "no": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsRune returns bool -- various cases", actual)
}

func Test_Cov5_IsEndsRune(t *testing.T) {
	actual := args.Map{
		"yes":   stringutil.IsEndsRune("hello", 'o'),
		"no":    stringutil.IsEndsRune("hello", 'x'),
		"empty": stringutil.IsEndsRune("", 'o'),
	}
	expected := args.Map{"yes": true, "no": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsEndsRune returns bool -- various cases", actual)
}

// ═══════════════════════════════════════════
// IsStartsAndEndsWith / IsStartsAndEnds / IsStartsAndEndsChar
// ═══════════════════════════════════════════

func Test_Cov5_IsStartsAndEndsWith(t *testing.T) {
	actual := args.Map{
		"both":    stringutil.IsStartsAndEndsWith("[hello]", "[", "]", false),
		"neither": stringutil.IsStartsAndEndsWith("hello", "[", "]", false),
	}
	expected := args.Map{"both": true, "neither": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith returns bool -- both cases", actual)
}

func Test_Cov5_IsStartsAndEnds(t *testing.T) {
	actual := args.Map{
		"both":    stringutil.IsStartsAndEnds("{hello}", "{", "}"),
		"neither": stringutil.IsStartsAndEnds("hello", "{", "}"),
	}
	expected := args.Map{"both": true, "neither": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEnds returns bool -- both cases", actual)
}

func Test_Cov5_IsStartsAndEndsChar(t *testing.T) {
	actual := args.Map{
		"both":  stringutil.IsStartsAndEndsChar("[x]", '[', ']'),
		"no":    stringutil.IsStartsAndEndsChar("hello", '[', ']'),
		"empty": stringutil.IsStartsAndEndsChar("", '[', ']'),
	}
	expected := args.Map{"both": true, "no": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns bool -- various cases", actual)
}

// ═══════════════════════════════════════════
// ClonePtr / SafeClonePtr
// ═══════════════════════════════════════════

func Test_Cov5_ClonePtr_NonNil(t *testing.T) {
	s := "hello"
	result := stringutil.ClonePtr(&s)
	actual := args.Map{"notNil": result != nil, "value": *result, "notSamePtr": result != &s}
	expected := args.Map{"notNil": true, "value": "hello", "notSamePtr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns clone -- non-nil", actual)
}

func Test_Cov5_ClonePtr_Nil(t *testing.T) {
	result := stringutil.ClonePtr(nil)
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil input", actual)
}

func Test_Cov5_SafeClonePtr_NonNil(t *testing.T) {
	s := "hello"
	result := stringutil.SafeClonePtr(&s)
	actual := args.Map{"notNil": result != nil, "value": *result}
	expected := args.Map{"notNil": true, "value": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns clone -- non-nil", actual)
}

func Test_Cov5_SafeClonePtr_Nil(t *testing.T) {
	result := stringutil.SafeClonePtr(nil)
	actual := args.Map{"notNil": result != nil, "value": *result}
	expected := args.Map{"notNil": true, "value": ""}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns empty clone -- nil input", actual)
}

// ═══════════════════════════════════════════
// AnyToString / AnyToTypeString
// ═══════════════════════════════════════════

func Test_Cov5_AnyToString(t *testing.T) {
	actual := args.Map{
		"string": stringutil.AnyToString("hello"),
		"int":    stringutil.AnyToString(42) != "",
		"nil":    stringutil.AnyToString(nil),
	}
	expected := args.Map{"string": "hello", "int": true, "nil": ""}
	expected.ShouldBeEqual(t, 0, "AnyToString returns string -- various types", actual)
}

func Test_Cov5_AnyToTypeString(t *testing.T) {
	result := stringutil.AnyToTypeString("hello")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypeString returns non-empty -- string input", actual)
}

// ═══════════════════════════════════════════
// ToBool
// ═══════════════════════════════════════════

func Test_Cov5_ToBool(t *testing.T) {
	actual := args.Map{
		"yes":   stringutil.ToBool("yes"),
		"Yes":   stringutil.ToBool("Yes"),
		"y":     stringutil.ToBool("y"),
		"1":     stringutil.ToBool("1"),
		"true":  stringutil.ToBool("true"),
		"no":    stringutil.ToBool("no"),
		"0":     stringutil.ToBool("0"),
		"false": stringutil.ToBool("false"),
		"empty": stringutil.ToBool(""),
		"bad":   stringutil.ToBool("invalid"),
	}
	expected := args.Map{
		"yes": true, "Yes": true, "y": true, "1": true, "true": true,
		"no": false, "0": false, "false": false, "empty": false, "bad": false,
	}
	expected.ShouldBeEqual(t, 0, "ToBool returns bool -- various inputs", actual)
}

// ═══════════════════════════════════════════
// ToInt / ToInt8 / ToInt16 / ToInt32 / ToIntDef / ToIntDefault
// ═══════════════════════════════════════════

func Test_Cov5_ToInt(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt("42", -1),
		"invalid": stringutil.ToInt("abc", -1),
	}
	expected := args.Map{"valid": 42, "invalid": -1}
	expected.ShouldBeEqual(t, 0, "ToInt returns int -- valid and invalid", actual)
}

func Test_Cov5_ToInt8(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt8("10", -1),
		"invalid": stringutil.ToInt8("abc", -1),
	}
	expected := args.Map{"valid": int8(10), "invalid": int8(-1)}
	expected.ShouldBeEqual(t, 0, "ToInt8 returns int8 -- valid and invalid", actual)
}

func Test_Cov5_ToInt8Def(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt8Def("10"),
		"invalid": stringutil.ToInt8Def("abc"),
	}
	expected := args.Map{"valid": int8(10), "invalid": int8(0)}
	expected.ShouldBeEqual(t, 0, "ToInt8Def returns int8 -- valid and invalid", actual)
}

func Test_Cov5_ToInt16(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt16("100", -1),
		"invalid": stringutil.ToInt16("abc", -1),
	}
	expected := args.Map{"valid": int16(100), "invalid": int16(-1)}
	expected.ShouldBeEqual(t, 0, "ToInt16 returns int16 -- valid and invalid", actual)
}

func Test_Cov5_ToInt32(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt32("100", -1),
		"invalid": stringutil.ToInt32("abc", -1),
	}
	expected := args.Map{"valid": int32(100), "invalid": int32(-1)}
	expected.ShouldBeEqual(t, 0, "ToInt32 returns int32 -- valid and invalid", actual)
}

func Test_Cov5_ToInt32Def(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt32Def("100"),
		"invalid": stringutil.ToInt32Def("abc"),
	}
	expected := args.Map{"valid": int32(100), "invalid": int32(0)}
	expected.ShouldBeEqual(t, 0, "ToInt32Def returns int32 -- valid and invalid", actual)
}

func Test_Cov5_ToIntDef(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToIntDef("42"),
		"invalid": stringutil.ToIntDef("abc"),
	}
	expected := args.Map{"valid": 42, "invalid": 0}
	expected.ShouldBeEqual(t, 0, "ToIntDef returns int -- valid and invalid", actual)
}

func Test_Cov5_ToIntDefault(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToIntDefault("42"),
		"invalid": stringutil.ToIntDefault("abc"),
	}
	expected := args.Map{"valid": 42, "invalid": 0}
	expected.ShouldBeEqual(t, 0, "ToIntDefault returns int -- valid and invalid", actual)
}

func Test_Cov5_ToInt16Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt16Default("100"),
		"invalid": stringutil.ToInt16Default("abc"),
		"outMax":  stringutil.ToInt16Default("40000"),
	}
	expected := args.Map{"valid": int16(100), "invalid": int16(0), "outMax": int16(0)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns int16 -- various inputs", actual)
}

func Test_Cov5_ToUint16Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToUint16Default("100"),
		"invalid": stringutil.ToUint16Default("abc"),
		"outMax":  stringutil.ToUint16Default("70000"),
	}
	expected := args.Map{"valid": uint16(100), "invalid": uint16(0), "outMax": uint16(0)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns uint16 -- various inputs", actual)
}

func Test_Cov5_ToUint32Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToUint32Default("100"),
		"invalid": stringutil.ToUint32Default("abc"),
		"negative": stringutil.ToUint32Default("-1"),
	}
	expected := args.Map{"valid": uint32(100), "invalid": uint32(0), "negative": uint32(0)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns uint32 -- various inputs", actual)
}

func Test_Cov5_ToByte(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToByte("100", 0),
		"invalid": stringutil.ToByte("abc", 42),
		"outMax":  stringutil.ToByte("300", 42),
	}
	expected := args.Map{"valid": byte(100), "invalid": byte(42), "outMax": byte(42)}
	expected.ShouldBeEqual(t, 0, "ToByte returns byte -- various inputs", actual)
}

func Test_Cov5_ToByteDefault(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToByteDefault("100"),
		"invalid": stringutil.ToByteDefault("abc"),
		"outMax":  stringutil.ToByteDefault("300"),
	}
	expected := args.Map{"valid": byte(100), "invalid": byte(0), "outMax": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns byte -- various inputs", actual)
}

// ═══════════════════════════════════════════
// ToIntUsingRegexMatch
// ═══════════════════════════════════════════

func Test_Cov5_ToIntUsingRegexMatch(t *testing.T) {
	digitRe := regexp.MustCompile(`^\d+$`)
	actual := args.Map{
		"valid":      stringutil.ToIntUsingRegexMatch(digitRe, "42"),
		"noMatch":    stringutil.ToIntUsingRegexMatch(digitRe, "abc"),
		"nilRegex":   stringutil.ToIntUsingRegexMatch(nil, "42"),
	}
	expected := args.Map{"valid": 42, "noMatch": 0, "nilRegex": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns int -- various inputs", actual)
}

// ═══════════════════════════════════════════
// MaskLine / MaskTrimLine / MaskLines / MaskTrimLines
// ═══════════════════════════════════════════

func Test_Cov5_MaskLine(t *testing.T) {
	actual := args.Map{
		"normal":    stringutil.MaskLine("XXXXXXX", "he"),
		"empty":     stringutil.MaskLine("XXXXXXX", ""),
		"longer":    stringutil.MaskLine("XX", "hello"),
		"emptyMask": stringutil.MaskLine("", "hello"),
	}
	expected := args.Map{"normal": "heXXXXX", "empty": "XXXXXXX", "longer": "hello", "emptyMask": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns masked -- various inputs", actual)
}

func Test_Cov5_MaskTrimLine(t *testing.T) {
	actual := args.Map{
		"normal":  stringutil.MaskTrimLine("XXXXXXX", "  he  "),
		"empty":   stringutil.MaskTrimLine("XXXXXXX", "  "),
		"longer":  stringutil.MaskTrimLine("XX", "  hello  "),
	}
	expected := args.Map{"normal": "heXXXXX", "empty": "XXXXXXX", "longer": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns masked -- various inputs", actual)
}

func Test_Cov5_MaskLines(t *testing.T) {
	result := stringutil.MaskLines("XXXXXXX", "he", "hel")
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "heXXXXX", "second": "helXXXX"}
	expected.ShouldBeEqual(t, 0, "MaskLines returns masked slice -- two items", actual)
}

func Test_Cov5_MaskLines_Empty(t *testing.T) {
	result := stringutil.MaskLines("XXXXXXX")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskLines returns empty -- no lines", actual)
}

func Test_Cov5_MaskTrimLines(t *testing.T) {
	result := stringutil.MaskTrimLines("XXXXXXX", "  he  ", "  hel  ")
	actual := args.Map{"len": len(result), "first": result[0], "second": result[1]}
	expected := args.Map{"len": 2, "first": "heXXXXX", "second": "helXXXX"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns masked slice -- two items", actual)
}

// ═══════════════════════════════════════════
// RemoveManyBySplitting
// ═══════════════════════════════════════════

func Test_Cov5_RemoveManyBySplitting(t *testing.T) {
	result := stringutil.RemoveManyBySplitting("hello--world--test", "--", "")
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 3, "first": "hello"}
	expected.ShouldBeEqual(t, 0, "RemoveManyBySplitting returns split -- with separator", actual)
}

// ═══════════════════════════════════════════
// SplitFirstLast
// ═══════════════════════════════════════════

func Test_Cov5_SplitFirstLast(t *testing.T) {
	first, last := stringutil.SplitFirstLast("a.b.c", ".")
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns first and last -- dot separator", actual)
}

func Test_Cov5_SplitFirstLast_NoSeparator(t *testing.T) {
	first, last := stringutil.SplitFirstLast("hello", ".")
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "hello", "last": ""}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns first and empty -- no separator", actual)
}

// ═══════════════════════════════════════════
// SplitLeftRight / SplitLeftRightTrimmed / SplitLeftRightType / SplitLeftRightTypeTrimmed / SplitLeftRightsTrims
// ═══════════════════════════════════════════

func Test_Cov5_SplitLeftRightTrimmed(t *testing.T) {
	left, right := stringutil.SplitLeftRightTrimmed(" key = value ", "=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "key", "right": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns trimmed -- key=value", actual)
}

func Test_Cov5_SplitLeftRightTrimmed_NoSeparator(t *testing.T) {
	left, right := stringutil.SplitLeftRightTrimmed(" hello ", "=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "hello", "right": ""}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns left only -- no separator", actual)
}

func Test_Cov5_SplitLeftRightType(t *testing.T) {
	result := stringutil.SplitLeftRightType("key=value", "=")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightType returns LeftRight -- key=value", actual)
}

func Test_Cov5_SplitLeftRightTypeTrimmed(t *testing.T) {
	result := stringutil.SplitLeftRightTypeTrimmed(" key = value ", "=")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTypeTrimmed returns LeftRight -- trimmed", actual)
}

func Test_Cov5_SplitLeftRightsTrims(t *testing.T) {
	result := stringutil.SplitLeftRightsTrims("=", " a=1 ", " b=2 ")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims returns slice -- two items", actual)
}
