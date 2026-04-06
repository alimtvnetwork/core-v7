package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ── IsEmpty / IsNotEmpty ──

func Test_Cov2_IsEmpty_True(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmpty("")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- empty string", actual)
}

func Test_Cov2_IsEmpty_False(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmpty("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- non-empty", actual)
}

func Test_Cov2_IsNotEmpty_True(t *testing.T) {
	actual := args.Map{"result": stringutil.IsNotEmpty("hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- non-empty", actual)
}

func Test_Cov2_IsNotEmpty_False(t *testing.T) {
	actual := args.Map{"result": stringutil.IsNotEmpty("")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- empty", actual)
}

// ── IsBlank ──

func Test_Cov2_IsBlank_Empty(t *testing.T) {
	actual := args.Map{"result": stringutil.IsBlank("")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlank returns empty -- empty", actual)
}

func Test_Cov2_IsBlank_Whitespace(t *testing.T) {
	actual := args.Map{"result": stringutil.IsBlank("   ")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlank returns correct value -- whitespace only", actual)
}

func Test_Cov2_IsBlank_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringutil.IsBlank("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsBlank returns empty -- non-empty", actual)
}

// ── IsEmptyPtr / IsBlankPtr ──

func Test_Cov2_IsEmptyPtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmptyPtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns nil -- nil", actual)
}

func Test_Cov2_IsEmptyPtr_Empty(t *testing.T) {
	s := ""
	actual := args.Map{"result": stringutil.IsEmptyPtr(&s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- empty", actual)
}

func Test_Cov2_IsEmptyPtr_NonEmpty(t *testing.T) {
	s := "hello"
	actual := args.Map{"result": stringutil.IsEmptyPtr(&s)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- non-empty", actual)
}

func Test_Cov2_IsBlankPtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsBlankPtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns nil -- nil", actual)
}

func Test_Cov2_IsBlankPtr_Whitespace(t *testing.T) {
	s := "   "
	actual := args.Map{"result": stringutil.IsBlankPtr(&s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns correct value -- whitespace", actual)
}

// ── IsEmptyOrWhitespace / IsEmptyOrWhitespacePtr ──

func Test_Cov2_IsEmptyOrWhitespace_Empty(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespace("")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- empty", actual)
}

func Test_Cov2_IsEmptyOrWhitespace_Whitespace(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespace("   ")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- whitespace", actual)
}

func Test_Cov2_IsEmptyOrWhitespace_NonEmpty(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespace("hello")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- non-empty", actual)
}

func Test_Cov2_IsEmptyOrWhitespacePtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespacePtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr returns nil -- nil", actual)
}

// ── IsDefined / IsDefinedPtr ──

func Test_Cov2_IsDefined_True(t *testing.T) {
	actual := args.Map{"result": stringutil.IsDefined("hello")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDefined returns empty -- non-empty", actual)
}

func Test_Cov2_IsDefined_False(t *testing.T) {
	actual := args.Map{"result": stringutil.IsDefined("")}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsDefined returns empty -- empty", actual)
}

func Test_Cov2_IsDefinedPtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsDefinedPtr(nil)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns nil -- nil", actual)
}

func Test_Cov2_IsDefinedPtr_NonEmpty(t *testing.T) {
	s := "hello"
	actual := args.Map{"result": stringutil.IsDefinedPtr(&s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns empty -- non-empty", actual)
}

// ── IsNullOrEmptyPtr ──

func Test_Cov2_IsNullOrEmptyPtr_Nil(t *testing.T) {
	actual := args.Map{"result": stringutil.IsNullOrEmptyPtr(nil)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns nil -- nil", actual)
}

func Test_Cov2_IsNullOrEmptyPtr_Empty(t *testing.T) {
	s := ""
	actual := args.Map{"result": stringutil.IsNullOrEmptyPtr(&s)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- empty", actual)
}

func Test_Cov2_IsNullOrEmptyPtr_NonEmpty(t *testing.T) {
	s := "hello"
	actual := args.Map{"result": stringutil.IsNullOrEmptyPtr(&s)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- non-empty", actual)
}

// ── IsStarts / IsEnds / IsContains ──

func Test_Cov2_IsStarts(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStarts("hello world", "hello"),
		"noMatch": stringutil.IsStarts("hello world", "world"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStarts returns correct value -- with args", actual)
}

func Test_Cov2_IsEnds(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsEnds("hello world", "world"),
		"noMatch": stringutil.IsEnds("hello world", "hello"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEnds returns correct value -- with args", actual)
}

func Test_Cov2_IsContains_Slice(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsContains([]string{"hello", "world"}, "world", 0, true),
		"noMatch": stringutil.IsContains([]string{"hello", "world"}, "xyz", 0, true),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- with args", actual)
}

// ── IsStartsWith / IsEndsWith ──

func Test_Cov2_IsStartsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsWith("hello", "hel", false),
		"noMatch": stringutil.IsStartsWith("hello", "xyz", false),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- with args", actual)
}

func Test_Cov2_IsEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsEndsWith("hello", "llo", false),
		"noMatch": stringutil.IsEndsWith("hello", "xyz", false),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- with args", actual)
}

// ── IsStartsChar / IsEndsChar / IsStartsRune / IsEndsRune ──

func Test_Cov2_IsStartsChar(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsChar("hello", 'h'),
		"noMatch": stringutil.IsStartsChar("hello", 'x'),
		"empty":   stringutil.IsStartsChar("", 'h'),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"empty":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsChar returns correct value -- with args", actual)
}

func Test_Cov2_IsEndsChar(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsEndsChar("hello", 'o'),
		"noMatch": stringutil.IsEndsChar("hello", 'x'),
		"empty":   stringutil.IsEndsChar("", 'o'),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"empty":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns correct value -- with args", actual)
}

func Test_Cov2_IsStartsRune(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsRune("hello", 'h'),
		"noMatch": stringutil.IsStartsRune("hello", 'x'),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsRune returns correct value -- with args", actual)
}

func Test_Cov2_IsEndsRune(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsEndsRune("hello", 'o'),
		"noMatch": stringutil.IsEndsRune("hello", 'x'),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsRune returns correct value -- with args", actual)
}

// ── IsStartsAndEndsChar / IsStartsAndEndsWith ──

func Test_Cov2_IsStartsAndEndsChar(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsAndEndsChar("[hello]", '[', ']'),
		"noMatch": stringutil.IsStartsAndEndsChar("hello", '[', ']'),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns correct value -- with args", actual)
}

func Test_Cov2_IsStartsAndEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsAndEndsWith("<<hello>>", "<<", ">>", false),
		"noMatch": stringutil.IsStartsAndEndsWith("hello", "<<", ">>", false),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith returns non-empty -- with args", actual)
}

// ── IsAnyStartsWith / IsAnyEndsWith ──

func Test_Cov2_IsAnyStartsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsAnyStartsWith("hello", false, "xyz", "hel"),
		"noMatch": stringutil.IsAnyStartsWith("hello", false, "xyz", "abc"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns non-empty -- with args", actual)
}

func Test_Cov2_IsAnyEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsAnyEndsWith("hello", false, "xyz", "llo"),
		"noMatch": stringutil.IsAnyEndsWith("hello", false, "xyz", "abc"),
	}
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns non-empty -- with args", actual)
}

// ── IsContainsPtr / IsContainsPtrSimple ──

func Test_Cov2_IsContainsPtr_Nil(t *testing.T) {
	find := "hello"
	actual := args.Map{"result": stringutil.IsContainsPtr(nil, &find, 0, true)}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns nil -- nil", actual)
}

func Test_Cov2_IsContainsPtr_Match(t *testing.T) {
	s := []string{"hello", "world"}
	find := "world"
	actual := args.Map{"result": stringutil.IsContainsPtr(&s, &find, 0, true)}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- match", actual)
}

// ── ClonePtr / SafeClonePtr ──

func Test_Cov2_ClonePtr_Nil(t *testing.T) {
	actual := args.Map{"isNil": stringutil.ClonePtr(nil) == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov2_ClonePtr_NonNil(t *testing.T) {
	s := "hello"
	cloned := stringutil.ClonePtr(&s)
	actual := args.Map{"value": *cloned, "diffPtr": cloned != &s}
	expected := args.Map{"value": "hello", "diffPtr": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- non-nil", actual)
}

func Test_Cov2_SafeClonePtr_Nil(t *testing.T) {
	result := stringutil.SafeClonePtr(nil)
	actual := args.Map{"notNil": result != nil, "value": *result}
	expected := args.Map{"notNil": true, "value": ""}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns nil -- nil", actual)
}

// ── FirstChar ──

func Test_Cov2_FirstChar(t *testing.T) {
	actual := args.Map{
		"normal": stringutil.FirstChar("hello"),
		"empty":  stringutil.FirstChar(""),
	}
	expected := args.Map{
		"normal": byte('h'),
		"empty":  byte(0),
	}
	expected.ShouldBeEqual(t, 0, "FirstChar returns correct value -- with args", actual)
}

// ── SafeSubstring ──

func Test_Cov2_SafeSubstring(t *testing.T) {
	actual := args.Map{
		"normal":    stringutil.SafeSubstring("hello", 1, 3),
		"outOfBound": stringutil.SafeSubstring("hello", 0, 100),
	}
	expected := args.Map{
		"normal":    "el",
		"outOfBound": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- with args", actual)
}

// ── ToBool ──

func Test_Cov2_ToBool(t *testing.T) {
	actual := args.Map{
		"true":    stringutil.ToBool("true"),
		"false":   stringutil.ToBool("false"),
		"empty":   stringutil.ToBool(""),
		"invalid": stringutil.ToBool("abc"),
	}
	expected := args.Map{
		"true":    true,
		"false":   false,
		"empty":   false,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "ToBool returns correct value -- with args", actual)
}

// ── ToInt / ToIntDef / ToIntDefault ──

func Test_Cov2_ToInt(t *testing.T) {
	result := stringutil.ToInt("42", -1)
	actual := args.Map{"value": result}
	expected := args.Map{"value": 42}
	expected.ShouldBeEqual(t, 0, "ToInt returns non-empty -- valid", actual)
}

func Test_Cov2_ToInt_Invalid(t *testing.T) {
	result := stringutil.ToInt("abc", -1)
	actual := args.Map{"value": result}
	expected := args.Map{"value": -1}
	expected.ShouldBeEqual(t, 0, "ToInt returns error -- invalid", actual)
}

func Test_Cov2_ToIntDefault(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToIntDefault("42"),
		"invalid": stringutil.ToIntDefault("abc"),
	}
	expected := args.Map{
		"valid":   42,
		"invalid": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToIntDefault returns correct value -- with args", actual)
}

// ── AnyToString ──

func Test_Cov2_AnyToString(t *testing.T) {
	actual := args.Map{
		"int":    stringutil.AnyToString(42),
		"string": stringutil.AnyToString("hello"),
		"bool":   stringutil.AnyToString(true),
	}
	expected := args.Map{
		"int":    "42",
		"string": "hello",
		"bool":   "true",
	}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- with args", actual)
}

// ── SplitLeftRight ──

func Test_Cov2_SplitLeftRight(t *testing.T) {
	left, right := stringutil.SplitLeftRight("hello:world", ":")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "hello", "right": "world"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- with args", actual)
}

func Test_Cov2_SplitLeftRight_NoSep(t *testing.T) {
	left, right := stringutil.SplitLeftRight("hello", ":")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "hello", "right": ""}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns empty -- no sep", actual)
}

// ── SplitLeftRightTrimmed ──

func Test_Cov2_SplitLeftRightTrimmed(t *testing.T) {
	left, right := stringutil.SplitLeftRightTrimmed(" hello : world ", ":")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "hello", "right": "world"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns correct value -- with args", actual)
}

// ── RemoveMany ──

func Test_Cov2_RemoveMany(t *testing.T) {
	result := stringutil.RemoveMany("hello world foo", "world", "foo")
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns correct value -- with args", actual)
}
