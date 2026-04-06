package stringutiltests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ══════════════════════════════════════════════════════════════════════════════
// AnyToString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_AnyToString_Nil(t *testing.T) {
	actual := args.Map{"v": stringutil.AnyToString(nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "AnyToString returns nil -- nil", actual)
}

func Test_Cov6_AnyToString_Value(t *testing.T) {
	actual := args.Map{"notEmpty": stringutil.AnyToString(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToStringNameField
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_AnyToStringNameField_Nil(t *testing.T) {
	actual := args.Map{"v": stringutil.AnyToStringNameField(nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField returns nil -- nil", actual)
}

func Test_Cov6_AnyToStringNameField_Struct(t *testing.T) {
	type s struct{ X int }
	actual := args.Map{"notEmpty": stringutil.AnyToStringNameField(s{X: 1}) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField returns correct value -- struct", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToTypeString
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_AnyToTypeString(t *testing.T) {
	actual := args.Map{"notEmpty": stringutil.AnyToTypeString(42) != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypeString returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ClonePtr / SafeClonePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ClonePtr_Nil(t *testing.T) {
	actual := args.Map{"nil": stringutil.ClonePtr(nil) == nil}
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_Cov6_ClonePtr_Value(t *testing.T) {
	s := "hello"
	c := stringutil.ClonePtr(&s)
	actual := args.Map{"val": *c, "diff": c != &s}
	expected := args.Map{"val": "hello", "diff": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- value", actual)
}

func Test_Cov6_SafeClonePtr_Nil(t *testing.T) {
	c := stringutil.SafeClonePtr(nil)
	actual := args.Map{"notNil": c != nil, "val": *c}
	expected := args.Map{"notNil": true, "val": ""}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns nil -- nil", actual)
}

func Test_Cov6_SafeClonePtr_Value(t *testing.T) {
	s := "hello"
	c := stringutil.SafeClonePtr(&s)
	actual := args.Map{"val": *c}
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns correct value -- value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FirstChar / LastChar variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_FirstChar_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.FirstChar("")}
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "FirstChar returns empty -- empty", actual)
}

func Test_Cov6_FirstChar_NonEmpty(t *testing.T) {
	actual := args.Map{"v": stringutil.FirstChar("abc")}
	expected := args.Map{"v": byte('a')}
	expected.ShouldBeEqual(t, 0, "FirstChar returns empty -- non-empty", actual)
}

func Test_Cov6_FirstCharOrDefault_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.FirstCharOrDefault("")}
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "FirstCharOrDefault returns empty -- empty", actual)
}

func Test_Cov6_FirstCharOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"v": stringutil.FirstCharOrDefault("xyz")}
	expected := args.Map{"v": byte('x')}
	expected.ShouldBeEqual(t, 0, "FirstCharOrDefault returns empty -- non-empty", actual)
}

func Test_Cov6_LastChar(t *testing.T) {
	actual := args.Map{"v": stringutil.LastChar("abc")}
	expected := args.Map{"v": byte('c')}
	expected.ShouldBeEqual(t, 0, "LastChar returns correct value -- with args", actual)
}

func Test_Cov6_LastCharOrDefault_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.LastCharOrDefault("")}
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "LastCharOrDefault returns empty -- empty", actual)
}

func Test_Cov6_LastCharOrDefault_NonEmpty(t *testing.T) {
	actual := args.Map{"v": stringutil.LastCharOrDefault("abc")}
	expected := args.Map{"v": byte('c')}
	expected.ShouldBeEqual(t, 0, "LastCharOrDefault returns empty -- non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsAnyEndsWith / IsAnyStartsWith — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsAnyEndsWith_ContentNoTerms(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyEndsWith("abc", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns empty -- content no terms", actual)
}

func Test_Cov6_IsAnyEndsWith_EmptyBoth(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyEndsWith("", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns empty -- empty both", actual)
}

func Test_Cov6_IsAnyEndsWith_Match(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyEndsWith("hello.txt", false, ".csv", ".txt")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns non-empty -- match", actual)
}

func Test_Cov6_IsAnyEndsWith_NoMatch(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyEndsWith("hello.txt", false, ".csv", ".json")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns empty -- no match", actual)
}

func Test_Cov6_IsAnyStartsWith_ContentNoTerms(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyStartsWith("abc", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns empty -- content no terms", actual)
}

func Test_Cov6_IsAnyStartsWith_EmptyBoth(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyStartsWith("", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns empty -- empty both", actual)
}

func Test_Cov6_IsAnyStartsWith_Match(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyStartsWith("hello world", false, "hi", "hello")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns non-empty -- match", actual)
}

func Test_Cov6_IsAnyStartsWith_NoMatch(t *testing.T) {
	actual := args.Map{"v": stringutil.IsAnyStartsWith("hello", false, "hi", "hey")}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns empty -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsBlank / IsBlankPtr / IsDefined / IsDefinedPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsBlank_All(t *testing.T) {
	actual := args.Map{
		"empty":   stringutil.IsBlank(""),
		"space":   stringutil.IsBlank(" "),
		"newline": stringutil.IsBlank("\n"),
		"tabs":    stringutil.IsBlank("\t  \t"),
		"text":    stringutil.IsBlank("x"),
	}
	expected := args.Map{"empty": true, "space": true, "newline": true, "tabs": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsBlank returns correct value -- all", actual)
}

func Test_Cov6_IsBlankPtr_Nil(t *testing.T) {
	actual := args.Map{"v": stringutil.IsBlankPtr(nil)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns nil -- nil", actual)
}

func Test_Cov6_IsBlankPtr_Value(t *testing.T) {
	s := "hello"
	actual := args.Map{"v": stringutil.IsBlankPtr(&s)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns correct value -- value", actual)
}

func Test_Cov6_IsDefined_All(t *testing.T) {
	actual := args.Map{"empty": stringutil.IsDefined(""), "text": stringutil.IsDefined("x")}
	expected := args.Map{"empty": false, "text": true}
	expected.ShouldBeEqual(t, 0, "IsDefined returns correct value -- with args", actual)
}

func Test_Cov6_IsDefinedPtr_Nil(t *testing.T) {
	actual := args.Map{"v": stringutil.IsDefinedPtr(nil)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns nil -- nil", actual)
}

func Test_Cov6_IsDefinedPtr_Value(t *testing.T) {
	s := "x"
	actual := args.Map{"v": stringutil.IsDefinedPtr(&s)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns correct value -- value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsContains / IsContainsPtr / IsContainsPtrSimple — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsContains_NilLines(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContains(nil, "a", 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns nil -- nil", actual)
}

func Test_Cov6_IsContains_EmptyLines(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContains([]string{}, "a", 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns empty -- empty", actual)
}

func Test_Cov6_IsContains_CaseSensitiveFound(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello", "World"}, "Hello", 0, true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-sensitive found", actual)
}

func Test_Cov6_IsContains_CaseSensitiveNotFound(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello"}, "hello", 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-sensitive not found", actual)
}

func Test_Cov6_IsContains_CaseInsensitiveFound(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello"}, "hello", 0, false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-insensitive found", actual)
}

func Test_Cov6_IsContains_CaseInsensitiveNotFound(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello"}, "xyz", 0, false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-insensitive not found", actual)
}

func Test_Cov6_IsContainsPtr_NilLines(t *testing.T) {
	f := "a"
	actual := args.Map{"v": stringutil.IsContainsPtr(nil, &f, 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns nil -- nil", actual)
}

func Test_Cov6_IsContainsPtr_EmptyLines(t *testing.T) {
	lines := []string{}
	f := "a"
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns empty -- empty", actual)
}

func Test_Cov6_IsContainsPtr_CaseSensitiveFound(t *testing.T) {
	lines := []string{"Hello"}
	f := "Hello"
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- sensitive found", actual)
}

func Test_Cov6_IsContainsPtr_CaseSensitiveNotFound(t *testing.T) {
	lines := []string{"Hello"}
	f := "hello"
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- sensitive not found", actual)
}

func Test_Cov6_IsContainsPtr_CaseInsensitiveFound(t *testing.T) {
	lines := []string{"Hello"}
	f := "hello"
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- insensitive found", actual)
}

func Test_Cov6_IsContainsPtr_CaseInsensitiveNotFound(t *testing.T) {
	lines := []string{"Hello"}
	f := "xyz"
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- insensitive not found", actual)
}

func Test_Cov6_IsContainsPtrSimple_NilLines(t *testing.T) {
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(nil, "a", 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns nil -- nil", actual)
}

func Test_Cov6_IsContainsPtrSimple_Empty(t *testing.T) {
	lines := []string{}
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "a", 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns empty -- empty", actual)
}

func Test_Cov6_IsContainsPtrSimple_SensitiveFound(t *testing.T) {
	lines := []string{"Hello"}
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "Hello", 0, true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- sensitive found", actual)
}

func Test_Cov6_IsContainsPtrSimple_SensitiveNotFound(t *testing.T) {
	lines := []string{"Hello"}
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "hello", 0, true)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- sensitive not found", actual)
}

func Test_Cov6_IsContainsPtrSimple_InsensitiveFound(t *testing.T) {
	lines := []string{"Hello"}
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "hello", 0, false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- insensitive found", actual)
}

func Test_Cov6_IsContainsPtrSimple_InsensitiveNotFound(t *testing.T) {
	lines := []string{"Hello"}
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "xyz", 0, false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- insensitive not found", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEmpty / IsNotEmpty / IsEmptyPtr / IsNullOrEmptyPtr / IsEmptyOrWhitespace / IsEmptyOrWhitespacePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsEmpty(t *testing.T) {
	actual := args.Map{"empty": stringutil.IsEmpty(""), "text": stringutil.IsEmpty("x")}
	expected := args.Map{"empty": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_Cov6_IsNotEmpty(t *testing.T) {
	actual := args.Map{"empty": stringutil.IsNotEmpty(""), "text": stringutil.IsNotEmpty("x")}
	expected := args.Map{"empty": false, "text": true}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- with args", actual)
}

func Test_Cov6_IsEmptyPtr(t *testing.T) {
	s := ""
	s2 := "x"
	actual := args.Map{"nil": stringutil.IsEmptyPtr(nil), "empty": stringutil.IsEmptyPtr(&s), "text": stringutil.IsEmptyPtr(&s2)}
	expected := args.Map{"nil": true, "empty": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- with args", actual)
}

func Test_Cov6_IsNullOrEmptyPtr(t *testing.T) {
	s := "x"
	actual := args.Map{"nil": stringutil.IsNullOrEmptyPtr(nil), "text": stringutil.IsNullOrEmptyPtr(&s)}
	expected := args.Map{"nil": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- with args", actual)
}

func Test_Cov6_IsEmptyOrWhitespace(t *testing.T) {
	actual := args.Map{
		"empty": stringutil.IsEmptyOrWhitespace(""),
		"space": stringutil.IsEmptyOrWhitespace(" "),
		"nl":    stringutil.IsEmptyOrWhitespace("\n"),
		"tabs":  stringutil.IsEmptyOrWhitespace("\t"),
		"text":  stringutil.IsEmptyOrWhitespace("x"),
	}
	expected := args.Map{"empty": true, "space": true, "nl": true, "tabs": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_Cov6_IsEmptyOrWhitespacePtr(t *testing.T) {
	s := "x"
	actual := args.Map{"nil": stringutil.IsEmptyOrWhitespacePtr(nil), "text": stringutil.IsEmptyOrWhitespacePtr(&s)}
	expected := args.Map{"nil": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEnds / IsEndsChar / IsEndsRune / IsEndsWith — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsEnds(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEnds("hello", "lo")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEnds returns correct value -- with args", actual)
}

func Test_Cov6_IsEndsChar_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsChar("", 'x')}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns empty -- empty", actual)
}

func Test_Cov6_IsEndsChar_Match(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsChar("abc", 'c')}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns correct value -- match", actual)
}

func Test_Cov6_IsEndsRune(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsRune("abc", 'c')}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsRune returns correct value -- with args", actual)
}

func Test_Cov6_IsEndsWith_EmptyEndsWith(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("abc", "", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns empty -- empty endsWith", actual)
}

func Test_Cov6_IsEndsWith_EmptyBase(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("", "x", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns empty -- empty base", actual)
}

func Test_Cov6_IsEndsWith_Equal(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("abc", "abc", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- equal", actual)
}

func Test_Cov6_IsEndsWith_EndsLonger(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("ab", "abcd", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- ends longer", actual)
}

func Test_Cov6_IsEndsWith_IgnoreCaseSameLen(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("ABC", "abc", true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- ignore case same len", actual)
}

func Test_Cov6_IsEndsWith_CaseSensitiveSuffix(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("hello world", "world", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- case-sensitive suffix", actual)
}

func Test_Cov6_IsEndsWith_CaseInsensitiveSuffix(t *testing.T) {
	actual := args.Map{"v": stringutil.IsEndsWith("hello WORLD", "world", true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- case-insensitive suffix", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStarts / IsStartsChar / IsStartsRune / IsStartsWith — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsStarts(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStarts("hello", "hel")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStarts returns correct value -- with args", actual)
}

func Test_Cov6_IsStartsChar(t *testing.T) {
	actual := args.Map{"empty": stringutil.IsStartsChar("", 'x'), "match": stringutil.IsStartsChar("abc", 'a')}
	expected := args.Map{"empty": false, "match": true}
	expected.ShouldBeEqual(t, 0, "IsStartsChar returns correct value -- with args", actual)
}

func Test_Cov6_IsStartsRune(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsRune("abc", 'a')}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsRune returns correct value -- with args", actual)
}

func Test_Cov6_IsStartsWith_EmptyStartsWith(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("abc", "", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns empty -- empty", actual)
}

func Test_Cov6_IsStartsWith_EmptyContent(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("", "x", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns empty -- empty content", actual)
}

func Test_Cov6_IsStartsWith_Equal(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("abc", "abc", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- equal", actual)
}

func Test_Cov6_IsStartsWith_StartsLonger(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("ab", "abcd", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- starts longer", actual)
}

func Test_Cov6_IsStartsWith_IgnoreCaseSameLen(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("ABC", "abc", true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- ignore case same len", actual)
}

func Test_Cov6_IsStartsWith_CaseSensitivePrefix(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("hello world", "hello", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- case-sensitive", actual)
}

func Test_Cov6_IsStartsWith_CaseInsensitivePrefix(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsWith("HELLO world", "hello", true)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- case-insensitive", actual)
}

func Test_Cov6_IsStartsWith_BaseLenEqualsStartsLen_NoCaseFold(t *testing.T) {
	// basePathLength <= startsWithLength branch when not equal and same len, case sensitive
	actual := args.Map{"v": stringutil.IsStartsWith("abc", "xyz", false)}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns empty -- same len no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStartsAndEndsChar / IsStartsAndEndsWith / IsStartsAndEnds
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_IsStartsAndEndsChar_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsAndEndsChar("", '{', '}')}
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns empty -- empty", actual)
}

func Test_Cov6_IsStartsAndEndsChar_Match(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsAndEndsChar("{hello}", '{', '}')}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns correct value -- match", actual)
}

func Test_Cov6_IsStartsAndEndsWith(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsAndEndsWith("hello world", "hello", "world", false)}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith returns non-empty -- with args", actual)
}

func Test_Cov6_IsStartsAndEnds(t *testing.T) {
	actual := args.Map{"v": stringutil.IsStartsAndEnds("hello world", "hello", "world")}
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEnds returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MaskLine / MaskLines / MaskTrimLine / MaskTrimLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_MaskLine_EmptyLine(t *testing.T) {
	actual := args.Map{"v": stringutil.MaskLine("----", "")}
	expected := args.Map{"v": "----"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns empty -- empty", actual)
}

func Test_Cov6_MaskLine_LineLongerThanMask(t *testing.T) {
	actual := args.Map{"v": stringutil.MaskLine("--", "abcde")}
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- line > mask", actual)
}

func Test_Cov6_MaskLine_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.MaskLine("--------", "abc")}
	expected := args.Map{"v": "abc-----"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- normal", actual)
}

func Test_Cov6_MaskLines_Empty(t *testing.T) {
	actual := args.Map{"len": len(stringutil.MaskLines("---"))}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskLines returns empty -- empty", actual)
}

func Test_Cov6_MaskLines_Normal(t *testing.T) {
	result := stringutil.MaskLines("--------", "abc", "de")
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "abc-----", "second": "de------"}
	expected.ShouldBeEqual(t, 0, "MaskLines returns correct value -- normal", actual)
}

func Test_Cov6_MaskLines_LineLongerThanMask(t *testing.T) {
	result := stringutil.MaskLines("--", "abcde")
	actual := args.Map{"v": result[0]}
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskLines returns correct value -- line > mask", actual)
}

func Test_Cov6_MaskTrimLine_EmptyAfterTrim(t *testing.T) {
	actual := args.Map{"v": stringutil.MaskTrimLine("----", "   ")}
	expected := args.Map{"v": "----"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns empty -- empty after trim", actual)
}

func Test_Cov6_MaskTrimLine_LineLongerThanMask(t *testing.T) {
	actual := args.Map{"v": stringutil.MaskTrimLine("--", "abcde")}
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- line > mask", actual)
}

func Test_Cov6_MaskTrimLine_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.MaskTrimLine("--------", " ab ")}
	expected := args.Map{"v": "ab------"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- normal", actual)
}

func Test_Cov6_MaskTrimLines_Empty(t *testing.T) {
	actual := args.Map{"len": len(stringutil.MaskTrimLines("---"))}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns empty -- empty", actual)
}

func Test_Cov6_MaskTrimLines_Normal(t *testing.T) {
	result := stringutil.MaskTrimLines("--------", " ab ", " c ")
	actual := args.Map{"first": result[0], "second": result[1]}
	expected := args.Map{"first": "ab------", "second": "c-------"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns correct value -- normal", actual)
}

func Test_Cov6_MaskTrimLines_LineLongerThanMask(t *testing.T) {
	result := stringutil.MaskTrimLines("--", "abcde")
	actual := args.Map{"v": result[0]}
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns correct value -- line > mask", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RemoveMany / RemoveManyBySplitting
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_RemoveMany_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.RemoveMany("")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns empty -- empty", actual)
}

func Test_Cov6_RemoveMany_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.RemoveMany("hello world foo", "world ", "foo")}
	expected := args.Map{"v": "hello "}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns correct value -- normal", actual)
}

func Test_Cov6_RemoveManyBySplitting(t *testing.T) {
	result := stringutil.RemoveManyBySplitting("a=1,b=2", ",", "=")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RemoveManyBySplitting returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeSubstring / SafeSubstringEnds / SafeSubstringStarts
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_SafeSubstring_BothMinusOne(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstring("hello", -1, -1)}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- both -1", actual)
}

func Test_Cov6_SafeSubstring_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstring("", 0, 3)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns empty -- empty", actual)
}

func Test_Cov6_SafeSubstring_StartMinusOne(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstring("hello", -1, 3)}
	expected := args.Map{"v": "hel"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- start -1", actual)
}

func Test_Cov6_SafeSubstring_EndMinusOne(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstring("hello", 2, -1)}
	expected := args.Map{"v": "llo"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- end -1", actual)
}

func Test_Cov6_SafeSubstring_ValidRange(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstring("hello", 1, 4)}
	expected := args.Map{"v": "ell"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns non-empty -- valid range", actual)
}

func Test_Cov6_SafeSubstring_OutOfRange(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstring("hi", 5, 10)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- out of range", actual)
}

func Test_Cov6_SafeSubstringEnds_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringEnds("", 3)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns empty -- empty", actual)
}

func Test_Cov6_SafeSubstringEnds_LenShorterThanEnd(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringEnds("hi", 10)}
	expected := args.Map{"v": "hi"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- len shorter", actual)
}

func Test_Cov6_SafeSubstringEnds_MinusOne(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringEnds("hello", -1)}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- -1", actual)
}

func Test_Cov6_SafeSubstringEnds_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringEnds("hello", 3)}
	expected := args.Map{"v": "hel"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- normal", actual)
}

func Test_Cov6_SafeSubstringStarts_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringStarts("", 0)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns empty -- empty", actual)
}

func Test_Cov6_SafeSubstringStarts_MinusOne(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringStarts("hello", -1)}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- -1", actual)
}

func Test_Cov6_SafeSubstringStarts_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringStarts("hello", 2)}
	expected := args.Map{"v": "llo"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- normal", actual)
}

func Test_Cov6_SafeSubstringStarts_BeyondLen(t *testing.T) {
	actual := args.Map{"v": stringutil.SafeSubstringStarts("hi", 10)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- beyond", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SplitContentsByWhitespaceConditions — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_SplitContentsByWhitespace_Basic(t *testing.T) {
	result := stringutil.SplitContentsByWhitespaceConditions("a b c", false, false, false, false, false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "SplitContents returns correct value -- basic", actual)
}

func Test_Cov6_SplitContentsByWhitespace_TrimAndNonEmpty(t *testing.T) {
	result := stringutil.SplitContentsByWhitespaceConditions("  a  b  ", true, true, false, false, false)
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitContents returns empty -- trim+nonEmpty", actual)
}

func Test_Cov6_SplitContentsByWhitespace_NonEmptyNoTrim(t *testing.T) {
	result := stringutil.SplitContentsByWhitespaceConditions("a b", false, true, false, false, false)
	actual := args.Map{"hasItems": len(result) > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "SplitContents returns empty -- nonEmpty no trim", actual)
}

func Test_Cov6_SplitContentsByWhitespace_UniqueAndSort(t *testing.T) {
	result := stringutil.SplitContentsByWhitespaceConditions("b a b a", false, false, true, true, true)
	actual := args.Map{"len": len(result), "first": result[0]}
	expected := args.Map{"len": 2, "first": "a"}
	expected.ShouldBeEqual(t, 0, "SplitContents returns correct value -- unique+sort", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SplitFirstLast / SplitLeftRight / SplitLeftRightTrimmed / SplitLeftRightType / SplitLeftRightTypeTrimmed / SplitLeftRightsTrims
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_SplitFirstLast_WithSep(t *testing.T) {
	first, last := stringutil.SplitFirstLast("a.b.c", ".")
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns non-empty -- with sep", actual)
}

func Test_Cov6_SplitFirstLast_NoSep(t *testing.T) {
	first, last := stringutil.SplitFirstLast("abc", ".")
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "abc", "last": ""}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns empty -- no sep", actual)
}

func Test_Cov6_SplitLeftRight_WithSep(t *testing.T) {
	left, right := stringutil.SplitLeftRight("key=val", "=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "key", "right": "val"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns non-empty -- with sep", actual)
}

func Test_Cov6_SplitLeftRight_NoSep(t *testing.T) {
	left, right := stringutil.SplitLeftRight("key", "=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "key", "right": ""}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns empty -- no sep", actual)
}

func Test_Cov6_SplitLeftRightTrimmed_WithSep(t *testing.T) {
	left, right := stringutil.SplitLeftRightTrimmed(" key = val ", "=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "key", "right": "val"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns non-empty -- with sep", actual)
}

func Test_Cov6_SplitLeftRightTrimmed_NoSep(t *testing.T) {
	left, right := stringutil.SplitLeftRightTrimmed(" key ", "=")
	actual := args.Map{"left": left, "right": right}
	expected := args.Map{"left": "key", "right": ""}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns empty -- no sep", actual)
}

func Test_Cov6_SplitLeftRightType(t *testing.T) {
	result := stringutil.SplitLeftRightType("key=val", "=")
	actual := args.Map{"left": result.Left, "right": result.Right}
	expected := args.Map{"left": "key", "right": "val"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightType returns correct value -- with args", actual)
}

func Test_Cov6_SplitLeftRightTypeTrimmed(t *testing.T) {
	result := stringutil.SplitLeftRightTypeTrimmed(" key = val ", "=")
	actual := args.Map{"left": result.Left, "right": result.Right}
	expected := args.Map{"left": "key", "right": "val"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTypeTrimmed returns correct value -- with args", actual)
}

func Test_Cov6_SplitLeftRightsTrims_Empty(t *testing.T) {
	result := stringutil.SplitLeftRightsTrims("=")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims returns empty -- empty", actual)
}

func Test_Cov6_SplitLeftRightsTrims_Items(t *testing.T) {
	result := stringutil.SplitLeftRightsTrims("=", " a = b ", " c = d ")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims returns correct value -- items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToBool — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ToBool_All(t *testing.T) {
	actual := args.Map{
		"empty":   stringutil.ToBool(""),
		"yes":     stringutil.ToBool("yes"),
		"Yes":     stringutil.ToBool("Yes"),
		"YES":     stringutil.ToBool("YES"),
		"y":       stringutil.ToBool("y"),
		"1":       stringutil.ToBool("1"),
		"no":      stringutil.ToBool("no"),
		"NO":      stringutil.ToBool("NO"),
		"No":      stringutil.ToBool("No"),
		"0":       stringutil.ToBool("0"),
		"n":       stringutil.ToBool("n"),
		"true":    stringutil.ToBool("true"),
		"false":   stringutil.ToBool("false"),
		"invalid": stringutil.ToBool("abc"),
	}
	expected := args.Map{
		"empty": false, "yes": true, "Yes": true, "YES": true, "y": true, "1": true,
		"no": false, "NO": false, "No": false, "0": false, "n": false,
		"true": true, "false": false, "invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "ToBool returns correct value -- all", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToByte / ToByteDefault — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ToByte_Valid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToByte("42", 0)}
	expected := args.Map{"v": byte(42)}
	expected.ShouldBeEqual(t, 0, "ToByte returns non-empty -- valid", actual)
}

func Test_Cov6_ToByte_Invalid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToByte("abc", 99)}
	expected := args.Map{"v": byte(99)}
	expected.ShouldBeEqual(t, 0, "ToByte returns error -- invalid", actual)
}

func Test_Cov6_ToByte_OutOfRange(t *testing.T) {
	actual := args.Map{"v": stringutil.ToByte("999", 77)}
	expected := args.Map{"v": byte(77)}
	expected.ShouldBeEqual(t, 0, "ToByte returns correct value -- out of range", actual)
}

func Test_Cov6_ToByteDefault_Valid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToByteDefault("42")}
	expected := args.Map{"v": byte(42)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns non-empty -- valid", actual)
}

func Test_Cov6_ToByteDefault_Invalid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToByteDefault("abc")}
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns error -- invalid", actual)
}

func Test_Cov6_ToByteDefault_OutOfRange(t *testing.T) {
	actual := args.Map{"v": stringutil.ToByteDefault("999")}
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns correct value -- out of range", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToInt / ToIntDef / ToIntDefault / ToInt8 / ToInt8Def / ToInt16 / ToInt16Default / ToInt32 / ToInt32Def
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ToInt(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToInt("42", -1), "invalid": stringutil.ToInt("abc", -1)}
	expected := args.Map{"valid": 42, "invalid": -1}
	expected.ShouldBeEqual(t, 0, "ToInt returns correct value -- with args", actual)
}

func Test_Cov6_ToIntDef(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToIntDef("42"), "invalid": stringutil.ToIntDef("abc")}
	expected := args.Map{"valid": 42, "invalid": 0}
	expected.ShouldBeEqual(t, 0, "ToIntDef returns correct value -- with args", actual)
}

func Test_Cov6_ToIntDefault(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToIntDefault("42"), "invalid": stringutil.ToIntDefault("abc")}
	expected := args.Map{"valid": 42, "invalid": 0}
	expected.ShouldBeEqual(t, 0, "ToIntDefault returns correct value -- with args", actual)
}

func Test_Cov6_ToInt8(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToInt8("42", -1), "invalid": stringutil.ToInt8("abc", -1)}
	expected := args.Map{"valid": int8(42), "invalid": int8(-1)}
	expected.ShouldBeEqual(t, 0, "ToInt8 returns correct value -- with args", actual)
}

func Test_Cov6_ToInt8Def(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToInt8Def("42"), "invalid": stringutil.ToInt8Def("abc")}
	expected := args.Map{"valid": int8(42), "invalid": int8(0)}
	expected.ShouldBeEqual(t, 0, "ToInt8Def returns correct value -- with args", actual)
}

func Test_Cov6_ToInt16(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToInt16("42", -1), "invalid": stringutil.ToInt16("abc", -1)}
	expected := args.Map{"valid": int16(42), "invalid": int16(-1)}
	expected.ShouldBeEqual(t, 0, "ToInt16 returns correct value -- with args", actual)
}

func Test_Cov6_ToInt16Default_Valid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToInt16Default("100")}
	expected := args.Map{"v": int16(100)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns non-empty -- valid", actual)
}

func Test_Cov6_ToInt16Default_Invalid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToInt16Default("abc")}
	expected := args.Map{"v": int16(0)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns error -- invalid", actual)
}

func Test_Cov6_ToInt16Default_OutOfRange(t *testing.T) {
	actual := args.Map{"v": stringutil.ToInt16Default("999999")}
	expected := args.Map{"v": int16(0)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns correct value -- out of range", actual)
}

func Test_Cov6_ToInt32(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToInt32("42", -1), "invalid": stringutil.ToInt32("abc", -1)}
	expected := args.Map{"valid": int32(42), "invalid": int32(-1)}
	expected.ShouldBeEqual(t, 0, "ToInt32 returns correct value -- with args", actual)
}

func Test_Cov6_ToInt32Def(t *testing.T) {
	actual := args.Map{"valid": stringutil.ToInt32Def("42"), "invalid": stringutil.ToInt32Def("abc")}
	expected := args.Map{"valid": int32(42), "invalid": int32(0)}
	expected.ShouldBeEqual(t, 0, "ToInt32Def returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToUint16Default / ToUint32Default — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ToUint16Default_Valid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToUint16Default("100")}
	expected := args.Map{"v": uint16(100)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns non-empty -- valid", actual)
}

func Test_Cov6_ToUint16Default_Invalid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToUint16Default("abc")}
	expected := args.Map{"v": uint16(0)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns error -- invalid", actual)
}

func Test_Cov6_ToUint16Default_OutOfRange(t *testing.T) {
	actual := args.Map{"v": stringutil.ToUint16Default("999999")}
	expected := args.Map{"v": uint16(0)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns correct value -- out of range", actual)
}

func Test_Cov6_ToUint32Default_Valid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToUint32Default("100")}
	expected := args.Map{"v": uint32(100)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns non-empty -- valid", actual)
}

func Test_Cov6_ToUint32Default_Invalid(t *testing.T) {
	actual := args.Map{"v": stringutil.ToUint32Default("abc")}
	expected := args.Map{"v": uint32(0)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns error -- invalid", actual)
}

func Test_Cov6_ToUint32Default_OutOfRange(t *testing.T) {
	// MaxInt32+1 would exceed the range
	actual := args.Map{"v": stringutil.ToUint32Default("99999999999")}
	expected := args.Map{"v": uint32(0)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns correct value -- out of range", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToIntUsingRegexMatch — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ToIntUsingRegexMatch_NilRegex(t *testing.T) {
	actual := args.Map{"v": stringutil.ToIntUsingRegexMatch(nil, "42")}
	expected := args.Map{"v": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns nil -- nil", actual)
}

func Test_Cov6_ToIntUsingRegexMatch_NoMatch(t *testing.T) {
	re := regexp.MustCompile(`^\d+$`)
	actual := args.Map{"v": stringutil.ToIntUsingRegexMatch(re, "abc")}
	expected := args.Map{"v": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns empty -- no match", actual)
}

func Test_Cov6_ToIntUsingRegexMatch_Valid(t *testing.T) {
	re := regexp.MustCompile(`^\d+$`)
	actual := args.Map{"v": stringutil.ToIntUsingRegexMatch(re, "42")}
	expected := args.Map{"v": 42}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns non-empty -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpacesToSingle (standalone)
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ReplaceWhiteSpacesToSingle_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceWhiteSpacesToSingle("")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns empty -- empty", actual)
}

func Test_Cov6_ReplaceWhiteSpacesToSingle_Whitespace(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceWhiteSpacesToSingle("  hello   world  \t foo  ")}
	expected := args.Map{"v": "hello world foo"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- ws", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceTemplate — all methods and branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_ReplaceTemplate_CurlyOne_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyOne("", "k", "v")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyOne returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyOne_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyOne("hello {name}!", "name", "world")}
	expected := args.Map{"v": "hello world!"}
	expected.ShouldBeEqual(t, 0, "CurlyOne returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_Curly_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.Curly("", map[string]string{"k": "v"})}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "Curly returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_Curly_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.Curly("{a}-{b}", map[string]string{"a": "1", "b": "2"})}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "Curly returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyTwo_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwo("", "a", 1, "b", 2)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyTwo returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyTwo_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwo("{a}-{b}", "a", 1, "b", 2)}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyTwo returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_DirectOne_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectOne("", "k", "v")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectOne returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_DirectOne_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectOne("hello KEY!", "KEY", "world")}
	expected := args.Map{"v": "hello world!"}
	expected.ShouldBeEqual(t, 0, "DirectOne returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_DirectTwoItem_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectTwoItem("", "a", 1, "b", 2)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectTwoItem returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_DirectTwoItem_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectTwoItem("A-B", "A", 1, "B", 2)}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectTwoItem returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyTwoItem_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwoItem("", "a", 1, "b", 2)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyTwoItem returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyTwoItem_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwoItem("{a}-{b}", "a", 1, "b", 2)}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyTwoItem returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_DirectKeyUsingMap_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMap("", map[string]string{"k": "v"})}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_DirectKeyUsingMap_EmptyMap(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMap("hello", map[string]string{})}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap returns empty -- empty map", actual)
}

func Test_Cov6_ReplaceTemplate_DirectKeyUsingMap_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMap("A-B", map[string]string{"A": "1", "B": "2"})}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_DirectKeyUsingKeyVal_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingKeyVal("")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_DirectKeyUsingKeyVal_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingKeyVal("A-B",
		stringutil.KeyValReplacer{Key: "A", Value: "1"},
		stringutil.KeyValReplacer{Key: "B", Value: "2"},
	)}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_DirectKeyUsingMapTrim(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMapTrim(" A-B ", map[string]string{"A": "1", "B": "2"})}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMapTrim returns correct value -- with args", actual)
}

func Test_Cov6_ReplaceTemplate_ReplaceWhiteSpaces_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_ReplaceWhiteSpaces_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("  hello  world  ")}
	expected := args.Map{"v": "helloworld"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_ReplaceWhiteSpacesToSingle_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_ReplaceWhiteSpacesToSingle_Normal(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("  hello   world  ")}
	expected := args.Map{"v": "hello world"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyKeyUsingMap_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("", map[string]string{"k": "v"})}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_CurlyKeyUsingMap_EmptyMap(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("hello", map[string]string{})}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap returns empty -- empty map", actual)
}

func Test_Cov6_ReplaceTemplate_UsingMapOptions_NonCurly(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingMapOptions(false, "A-B", map[string]string{"A": "1", "B": "2"})}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions returns non-empty -- non-curly", actual)
}

func Test_Cov6_ReplaceTemplate_UsingMapOptions_Curly(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingMapOptions(true, "{a}-{b}", map[string]string{"a": "1", "b": "2"})}
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions returns correct value -- curly", actual)
}

func Test_Cov6_ReplaceTemplate_UsingMapOptions_EmptyMap(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingMapOptions(true, "hello", map[string]string{})}
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions returns empty -- empty map", actual)
}

type testNamer struct{ name string }

func (n testNamer) Name() string { return n.name }

func Test_Cov6_ReplaceTemplate_UsingNamerMapOptions_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingNamerMapOptions(true, "", nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingNamerMapOptions returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_UsingStringerMapOptions_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingStringerMapOptions(true, "", nil)}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingStringerMapOptions returns empty -- empty", actual)
}

type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_Cov6_ReplaceTemplate_UsingStringerMapOptions_Curly(t *testing.T) {
	m := map[fmt.Stringer]string{testStringer{"a"}: "1"}
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingStringerMapOptions(true, "{a}", m)}
	expected := args.Map{"v": "1"}
	expected.ShouldBeEqual(t, 0, "UsingStringerMapOptions returns correct value -- curly", actual)
}

func Test_Cov6_ReplaceTemplate_UsingStringerMapOptions_NonCurly(t *testing.T) {
	m := map[fmt.Stringer]string{testStringer{"A"}: "1"}
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingStringerMapOptions(false, "A-B", m)}
	expected := args.Map{"v": "1-B"}
	expected.ShouldBeEqual(t, 0, "UsingStringerMapOptions returns non-empty -- non-curly", actual)
}

func Test_Cov6_ReplaceTemplate_UsingWrappedTemplate_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingWrappedTemplate("", "x")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingWrappedTemplate returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_UsingWrappedTemplate_Normal(t *testing.T) {
	actual := args.Map{"has": stringutil.ReplaceTemplate.UsingWrappedTemplate("{wrapped}", "x") != ""}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingWrappedTemplate returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_UsingBracketsWrappedTemplate_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("", "x")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingBracketsWrappedTemplate returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_UsingQuotesWrappedTemplate_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate("", "x")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingQuotesWrappedTemplate returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_UsingValueTemplate_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingValueTemplate("", "x")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingValueTemplate returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_UsingValueTemplate_Normal(t *testing.T) {
	actual := args.Map{"has": stringutil.ReplaceTemplate.UsingValueTemplate("test {value} end", "X") != ""}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingValueTemplate returns correct value -- normal", actual)
}

func Test_Cov6_ReplaceTemplate_UsingValueWithFieldsTemplate_Empty(t *testing.T) {
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingValueWithFieldsTemplate("", "x")}
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingValueWithFieldsTemplate returns empty -- empty", actual)
}

func Test_Cov6_ReplaceTemplate_UsingValueWithFieldsTemplate_Normal(t *testing.T) {
	actual := args.Map{"has": stringutil.ReplaceTemplate.UsingValueWithFieldsTemplate("test {value-fields} end", "X") != ""}
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingValueWithFieldsTemplate returns non-empty -- normal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValReplacer — struct
// ══════════════════════════════════════════════════════════════════════════════

func Test_Cov6_KeyValReplacer(t *testing.T) {
	kvr := stringutil.KeyValReplacer{Key: "k", Value: "v"}
	actual := args.Map{"key": kvr.Key, "val": kvr.Value}
	expected := args.Map{"key": "k", "val": "v"}
	expected.ShouldBeEqual(t, 0, "KeyValReplacer returns correct value -- with args", actual)
}
