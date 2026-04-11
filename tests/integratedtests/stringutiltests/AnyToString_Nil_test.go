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

func Test_AnyToString_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.AnyToString(nil)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "AnyToString returns nil -- nil", actual)
}

func Test_AnyToString_Value(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": stringutil.AnyToString(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToStringNameField
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyToStringNameField_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.AnyToStringNameField(nil)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField returns nil -- nil", actual)
}

func Test_AnyToStringNameField_Struct(t *testing.T) {
	// Arrange
	type s struct{ X int }

	// Act
	actual := args.Map{"notEmpty": stringutil.AnyToStringNameField(s{X: 1}) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField returns correct value -- struct", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// AnyToTypeString
// ══════════════════════════════════════════════════════════════════════════════

func Test_AnyToTypeString_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{"notEmpty": stringutil.AnyToTypeString(42) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToTypeString returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ClonePtr / SafeClonePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_ClonePtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"nil": stringutil.ClonePtr(nil) == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_ClonePtr_Value(t *testing.T) {
	// Arrange
	s := "hello"
	c := stringutil.ClonePtr(&s)

	// Act
	actual := args.Map{
		"val": *c,
		"diff": c != &s,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"diff": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- value", actual)
}

func Test_SafeClonePtr_Nil(t *testing.T) {
	// Arrange
	c := stringutil.SafeClonePtr(nil)

	// Act
	actual := args.Map{
		"notNil": c != nil,
		"val": *c,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"val": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns nil -- nil", actual)
}

func Test_SafeClonePtr_Value(t *testing.T) {
	// Arrange
	s := "hello"
	c := stringutil.SafeClonePtr(&s)

	// Act
	actual := args.Map{"val": *c}

	// Assert
	expected := args.Map{"val": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns correct value -- value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// FirstChar / LastChar variants
// ══════════════════════════════════════════════════════════════════════════════

func Test_FirstChar_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.FirstChar("")}

	// Assert
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "FirstChar returns empty -- empty", actual)
}

func Test_FirstChar_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.FirstChar("abc")}

	// Assert
	expected := args.Map{"v": byte('a')}
	expected.ShouldBeEqual(t, 0, "FirstChar returns empty -- non-empty", actual)
}

func Test_FirstCharOrDefault_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.FirstCharOrDefault("")}

	// Assert
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "FirstCharOrDefault returns empty -- empty", actual)
}

func Test_FirstCharOrDefault_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.FirstCharOrDefault("xyz")}

	// Assert
	expected := args.Map{"v": byte('x')}
	expected.ShouldBeEqual(t, 0, "FirstCharOrDefault returns empty -- non-empty", actual)
}

func Test_LastChar(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.LastChar("abc")}

	// Assert
	expected := args.Map{"v": byte('c')}
	expected.ShouldBeEqual(t, 0, "LastChar returns correct value -- with args", actual)
}

func Test_LastCharOrDefault_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.LastCharOrDefault("")}

	// Assert
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "LastCharOrDefault returns empty -- empty", actual)
}

func Test_LastCharOrDefault_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.LastCharOrDefault("abc")}

	// Assert
	expected := args.Map{"v": byte('c')}
	expected.ShouldBeEqual(t, 0, "LastCharOrDefault returns empty -- non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsAnyEndsWith / IsAnyStartsWith — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsAnyEndsWith_ContentNoTerms(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyEndsWith("abc", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns empty -- content no terms", actual)
}

func Test_IsAnyEndsWith_EmptyBoth(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyEndsWith("", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns empty -- empty both", actual)
}

func Test_IsAnyEndsWith_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyEndsWith("hello.txt", false, ".csv", ".txt")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns non-empty -- match", actual)
}

func Test_IsAnyEndsWith_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyEndsWith("hello.txt", false, ".csv", ".json")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns empty -- no match", actual)
}

func Test_IsAnyStartsWith_ContentNoTerms(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyStartsWith("abc", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns empty -- content no terms", actual)
}

func Test_IsAnyStartsWith_EmptyBoth(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyStartsWith("", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns empty -- empty both", actual)
}

func Test_IsAnyStartsWith_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyStartsWith("hello world", false, "hi", "hello")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns non-empty -- match", actual)
}

func Test_IsAnyStartsWith_NoMatch(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsAnyStartsWith("hello", false, "hi", "hey")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns empty -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsBlank / IsBlankPtr / IsDefined / IsDefinedPtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsBlank_All(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":   stringutil.IsBlank(""),
		"space":   stringutil.IsBlank(" "),
		"newline": stringutil.IsBlank("\n"),
		"tabs":    stringutil.IsBlank("\t  \t"),
		"text":    stringutil.IsBlank("x"),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"space": true,
		"newline": true,
		"tabs": true,
		"text": false,
	}
	expected.ShouldBeEqual(t, 0, "IsBlank returns correct value -- all", actual)
}

func Test_IsBlankPtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsBlankPtr(nil)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns nil -- nil", actual)
}

func Test_IsBlankPtr_Value(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	actual := args.Map{"v": stringutil.IsBlankPtr(&s)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns correct value -- value", actual)
}

func Test_IsDefined_All(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": stringutil.IsDefined(""),
		"text": stringutil.IsDefined("x"),
	}

	// Assert
	expected := args.Map{
		"empty": false,
		"text": true,
	}
	expected.ShouldBeEqual(t, 0, "IsDefined returns correct value -- with args", actual)
}

func Test_IsDefinedPtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsDefinedPtr(nil)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns nil -- nil", actual)
}

func Test_IsDefinedPtr_Value(t *testing.T) {
	// Arrange
	s := "x"

	// Act
	actual := args.Map{"v": stringutil.IsDefinedPtr(&s)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns correct value -- value", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsContains / IsContainsPtr / IsContainsPtrSimple — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsContains_NilLines(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContains(nil, "a", 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns nil -- nil", actual)
}

func Test_IsContains_EmptyLines(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContains([]string{}, "a", 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns empty -- empty", actual)
}

func Test_IsContains_CaseSensitiveFound(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello", "World"}, "Hello", 0, true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-sensitive found", actual)
}

func Test_IsContains_CaseSensitiveNotFound(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello"}, "hello", 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-sensitive not found", actual)
}

func Test_IsContains_CaseInsensitiveFound(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello"}, "hello", 0, false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-insensitive found", actual)
}

func Test_IsContains_CaseInsensitiveNotFound(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContains([]string{"Hello"}, "xyz", 0, false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- case-insensitive not found", actual)
}

func Test_IsContainsPtr_NilLines(t *testing.T) {
	// Arrange
	f := "a"

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtr(nil, &f, 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns nil -- nil", actual)
}

func Test_IsContainsPtr_EmptyLines(t *testing.T) {
	// Arrange
	lines := []string{}
	f := "a"

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns empty -- empty", actual)
}

func Test_IsContainsPtr_CaseSensitiveFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}
	f := "Hello"

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- sensitive found", actual)
}

func Test_IsContainsPtr_CaseSensitiveNotFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}
	f := "hello"

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- sensitive not found", actual)
}

func Test_IsContainsPtr_CaseInsensitiveFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}
	f := "hello"

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- insensitive found", actual)
}

func Test_IsContainsPtr_CaseInsensitiveNotFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}
	f := "xyz"

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtr(&lines, &f, 0, false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- insensitive not found", actual)
}

func Test_IsContainsPtrSimple_NilLines(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(nil, "a", 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns nil -- nil", actual)
}

func Test_IsContainsPtrSimple_Empty(t *testing.T) {
	// Arrange
	lines := []string{}

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "a", 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns empty -- empty", actual)
}

func Test_IsContainsPtrSimple_SensitiveFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "Hello", 0, true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- sensitive found", actual)
}

func Test_IsContainsPtrSimple_SensitiveNotFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "hello", 0, true)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- sensitive not found", actual)
}

func Test_IsContainsPtrSimple_InsensitiveFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "hello", 0, false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- insensitive found", actual)
}

func Test_IsContainsPtrSimple_InsensitiveNotFound(t *testing.T) {
	// Arrange
	lines := []string{"Hello"}

	// Act
	actual := args.Map{"v": stringutil.IsContainsPtrSimple(&lines, "xyz", 0, false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- insensitive not found", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEmpty / IsNotEmpty / IsEmptyPtr / IsNullOrEmptyPtr / IsEmptyOrWhitespace / IsEmptyOrWhitespacePtr
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsEmpty(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": stringutil.IsEmpty(""),
		"text": stringutil.IsEmpty("x"),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"text": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_IsNotEmpty(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": stringutil.IsNotEmpty(""),
		"text": stringutil.IsNotEmpty("x"),
	}

	// Assert
	expected := args.Map{
		"empty": false,
		"text": true,
	}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- with args", actual)
}

func Test_IsEmptyPtr(t *testing.T) {
	// Arrange
	s := ""
	s2 := "x"

	// Act
	actual := args.Map{
		"nil": stringutil.IsEmptyPtr(nil),
		"empty": stringutil.IsEmptyPtr(&s),
		"text": stringutil.IsEmptyPtr(&s2),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"empty": true,
		"text": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- with args", actual)
}

func Test_IsNullOrEmptyPtr(t *testing.T) {
	// Arrange
	s := "x"

	// Act
	actual := args.Map{
		"nil": stringutil.IsNullOrEmptyPtr(nil),
		"text": stringutil.IsNullOrEmptyPtr(&s),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"text": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- with args", actual)
}

func Test_IsEmptyOrWhitespace(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": stringutil.IsEmptyOrWhitespace(""),
		"space": stringutil.IsEmptyOrWhitespace(" "),
		"nl":    stringutil.IsEmptyOrWhitespace("\n"),
		"tabs":  stringutil.IsEmptyOrWhitespace("\t"),
		"text":  stringutil.IsEmptyOrWhitespace("x"),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"space": true,
		"nl": true,
		"tabs": true,
		"text": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_IsEmptyOrWhitespacePtr_FromAnyToStringNil(t *testing.T) {
	// Arrange
	s := "x"

	// Act
	actual := args.Map{
		"nil": stringutil.IsEmptyOrWhitespacePtr(nil),
		"text": stringutil.IsEmptyOrWhitespacePtr(&s),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"text": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr returns empty -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsEnds / IsEndsChar / IsEndsRune / IsEndsWith — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsEnds_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEnds("hello", "lo")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEnds returns correct value -- with args", actual)
}

func Test_IsEndsChar_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsChar("", 'x')}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns empty -- empty", actual)
}

func Test_IsEndsChar_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsChar("abc", 'c')}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns correct value -- match", actual)
}

func Test_IsEndsRune_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsRune("abc", 'c')}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsRune returns correct value -- with args", actual)
}

func Test_IsEndsWith_EmptyEndsWith(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("abc", "", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns empty -- empty endsWith", actual)
}

func Test_IsEndsWith_EmptyBase(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("", "x", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns empty -- empty base", actual)
}

func Test_IsEndsWith_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("abc", "abc", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- equal", actual)
}

func Test_IsEndsWith_EndsLonger(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("ab", "abcd", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- ends longer", actual)
}

func Test_IsEndsWith_IgnoreCaseSameLen(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("ABC", "abc", true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- ignore case same len", actual)
}

func Test_IsEndsWith_CaseSensitiveSuffix(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("hello world", "world", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- case-sensitive suffix", actual)
}

func Test_IsEndsWith_CaseInsensitiveSuffix(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsEndsWith("hello WORLD", "world", true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- case-insensitive suffix", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStarts / IsStartsChar / IsStartsRune / IsStartsWith — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStarts_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStarts("hello", "hel")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStarts returns correct value -- with args", actual)
}

func Test_IsStartsChar_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": stringutil.IsStartsChar("", 'x'),
		"match": stringutil.IsStartsChar("abc", 'a'),
	}

	// Assert
	expected := args.Map{
		"empty": false,
		"match": true,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsChar returns correct value -- with args", actual)
}

func Test_IsStartsRune_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsRune("abc", 'a')}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsRune returns correct value -- with args", actual)
}

func Test_IsStartsWith_EmptyStartsWith(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("abc", "", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns empty -- empty", actual)
}

func Test_IsStartsWith_EmptyContent(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("", "x", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns empty -- empty content", actual)
}

func Test_IsStartsWith_Equal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("abc", "abc", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- equal", actual)
}

func Test_IsStartsWith_StartsLonger(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("ab", "abcd", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- starts longer", actual)
}

func Test_IsStartsWith_IgnoreCaseSameLen(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("ABC", "abc", true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- ignore case same len", actual)
}

func Test_IsStartsWith_CaseSensitivePrefix(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("hello world", "hello", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- case-sensitive", actual)
}

func Test_IsStartsWith_CaseInsensitivePrefix(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("HELLO world", "hello", true)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- case-insensitive", actual)
}

func Test_IsStartsWith_BaseLenEqualsStartsLen_NoCaseFold(t *testing.T) {
	// basePathLength <= startsWithLength branch when not equal and same len, case sensitive
	// Act
	actual := args.Map{"v": stringutil.IsStartsWith("abc", "xyz", false)}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns empty -- same len no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// IsStartsAndEndsChar / IsStartsAndEndsWith / IsStartsAndEnds
// ══════════════════════════════════════════════════════════════════════════════

func Test_IsStartsAndEndsChar_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsAndEndsChar("", '{', '}')}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns empty -- empty", actual)
}

func Test_IsStartsAndEndsChar_Match(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsAndEndsChar("{hello}", '{', '}')}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns correct value -- match", actual)
}

func Test_IsStartsAndEndsWith_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsAndEndsWith("hello world", "hello", "world", false)}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith returns non-empty -- with args", actual)
}

func Test_IsStartsAndEnds(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.IsStartsAndEnds("hello world", "hello", "world")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEnds returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MaskLine / MaskLines / MaskTrimLine / MaskTrimLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_MaskLine_EmptyLine(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.MaskLine("----", "")}

	// Assert
	expected := args.Map{"v": "----"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns empty -- empty", actual)
}

func Test_MaskLine_LineLongerThanMask(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.MaskLine("--", "abcde")}

	// Assert
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- line > mask", actual)
}

func Test_MaskLine_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.MaskLine("--------", "abc")}

	// Assert
	expected := args.Map{"v": "abc-----"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- normal", actual)
}

func Test_MaskLines_Empty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(stringutil.MaskLines("---"))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskLines returns empty -- empty", actual)
}

func Test_MaskLines_Normal(t *testing.T) {
	// Arrange
	result := stringutil.MaskLines("--------", "abc", "de")

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "abc-----",
		"second": "de------",
	}
	expected.ShouldBeEqual(t, 0, "MaskLines returns correct value -- normal", actual)
}

func Test_MaskLines_LineLongerThanMask(t *testing.T) {
	// Arrange
	result := stringutil.MaskLines("--", "abcde")

	// Act
	actual := args.Map{"v": result[0]}

	// Assert
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskLines returns correct value -- line > mask", actual)
}

func Test_MaskTrimLine_EmptyAfterTrim(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.MaskTrimLine("----", "   ")}

	// Assert
	expected := args.Map{"v": "----"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns empty -- empty after trim", actual)
}

func Test_MaskTrimLine_LineLongerThanMask(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.MaskTrimLine("--", "abcde")}

	// Assert
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- line > mask", actual)
}

func Test_MaskTrimLine_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.MaskTrimLine("--------", " ab ")}

	// Assert
	expected := args.Map{"v": "ab------"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- normal", actual)
}

func Test_MaskTrimLines_Empty(t *testing.T) {
	// Act
	actual := args.Map{"len": len(stringutil.MaskTrimLines("---"))}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns empty -- empty", actual)
}

func Test_MaskTrimLines_Normal(t *testing.T) {
	// Arrange
	result := stringutil.MaskTrimLines("--------", " ab ", " c ")

	// Act
	actual := args.Map{
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"first": "ab------",
		"second": "c-------",
	}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns correct value -- normal", actual)
}

func Test_MaskTrimLines_LineLongerThanMask(t *testing.T) {
	// Arrange
	result := stringutil.MaskTrimLines("--", "abcde")

	// Act
	actual := args.Map{"v": result[0]}

	// Assert
	expected := args.Map{"v": "abcde"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns correct value -- line > mask", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// RemoveMany / RemoveManyBySplitting
// ══════════════════════════════════════════════════════════════════════════════

func Test_RemoveMany_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.RemoveMany("")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns empty -- empty", actual)
}

func Test_RemoveMany_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.RemoveMany("hello world foo", "world ", "foo")}

	// Assert
	expected := args.Map{"v": "hello "}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns correct value -- normal", actual)
}

func Test_RemoveManyBySplitting_FromAnyToStringNil(t *testing.T) {
	// Arrange
	result := stringutil.RemoveManyBySplitting("a=1,b=2", ",", "=")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RemoveManyBySplitting returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeSubstring / SafeSubstringEnds / SafeSubstringStarts
// ══════════════════════════════════════════════════════════════════════════════

func Test_SafeSubstring_BothMinusOne(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstring("hello", -1, -1)}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- both -1", actual)
}

func Test_SafeSubstring_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstring("", 0, 3)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns empty -- empty", actual)
}

func Test_SafeSubstring_StartMinusOne(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstring("hello", -1, 3)}

	// Assert
	expected := args.Map{"v": "hel"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- start -1", actual)
}

func Test_SafeSubstring_EndMinusOne(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstring("hello", 2, -1)}

	// Assert
	expected := args.Map{"v": "llo"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- end -1", actual)
}

func Test_SafeSubstring_ValidRange(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstring("hello", 1, 4)}

	// Assert
	expected := args.Map{"v": "ell"}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns non-empty -- valid range", actual)
}

func Test_SafeSubstring_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstring("hi", 5, 10)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- out of range", actual)
}

func Test_SafeSubstringEnds_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringEnds("", 3)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns empty -- empty", actual)
}

func Test_SafeSubstringEnds_LenShorterThanEnd(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringEnds("hi", 10)}

	// Assert
	expected := args.Map{"v": "hi"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- len shorter", actual)
}

func Test_SafeSubstringEnds_MinusOne(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringEnds("hello", -1)}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- -1", actual)
}

func Test_SafeSubstringEnds_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringEnds("hello", 3)}

	// Assert
	expected := args.Map{"v": "hel"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- normal", actual)
}

func Test_SafeSubstringStarts_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringStarts("", 0)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns empty -- empty", actual)
}

func Test_SafeSubstringStarts_MinusOne(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringStarts("hello", -1)}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- -1", actual)
}

func Test_SafeSubstringStarts_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringStarts("hello", 2)}

	// Assert
	expected := args.Map{"v": "llo"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- normal", actual)
}

func Test_SafeSubstringStarts_BeyondLen(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.SafeSubstringStarts("hi", 10)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- beyond", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SplitContentsByWhitespaceConditions — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_SplitContentsByWhitespace_Basic(t *testing.T) {
	// Arrange
	result := stringutil.SplitContentsByWhitespaceConditions("a b c", false, false, false, false, false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "SplitContents returns correct value -- basic", actual)
}

func Test_SplitContentsByWhitespace_TrimAndNonEmpty(t *testing.T) {
	// Arrange
	result := stringutil.SplitContentsByWhitespaceConditions("  a  b  ", true, true, false, false, false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitContents returns empty -- trim+nonEmpty", actual)
}

func Test_SplitContentsByWhitespace_NonEmptyNoTrim(t *testing.T) {
	// Arrange
	result := stringutil.SplitContentsByWhitespaceConditions("a b", false, true, false, false, false)

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "SplitContents returns empty -- nonEmpty no trim", actual)
}

func Test_SplitContentsByWhitespace_UniqueAndSort(t *testing.T) {
	// Arrange
	result := stringutil.SplitContentsByWhitespaceConditions("b a b a", false, false, true, true, true)

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "SplitContents returns correct value -- unique+sort", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SplitFirstLast / SplitLeftRight / SplitLeftRightTrimmed / SplitLeftRightType / SplitLeftRightTypeTrimmed / SplitLeftRightsTrims
// ══════════════════════════════════════════════════════════════════════════════

func Test_SplitFirstLast_WithSep(t *testing.T) {
	// Arrange
	first, last := stringutil.SplitFirstLast("a.b.c", ".")

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"last": "c",
	}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns non-empty -- with sep", actual)
}

func Test_SplitFirstLast_NoSep(t *testing.T) {
	// Arrange
	first, last := stringutil.SplitFirstLast("abc", ".")

	// Act
	actual := args.Map{
		"first": first,
		"last": last,
	}

	// Assert
	expected := args.Map{
		"first": "abc",
		"last": "",
	}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns empty -- no sep", actual)
}

func Test_SplitLeftRight_WithSep(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRight("key=val", "=")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "val",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns non-empty -- with sep", actual)
}

func Test_SplitLeftRight_NoSep(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRight("key", "=")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns empty -- no sep", actual)
}

func Test_SplitLeftRightTrimmed_WithSep(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRightTrimmed(" key = val ", "=")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "val",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns non-empty -- with sep", actual)
}

func Test_SplitLeftRightTrimmed_NoSep(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRightTrimmed(" key ", "=")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns empty -- no sep", actual)
}

func Test_SplitLeftRightType(t *testing.T) {
	// Arrange
	result := stringutil.SplitLeftRightType("key=val", "=")

	// Act
	actual := args.Map{
		"left": result.Left,
		"right": result.Right,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "val",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightType returns correct value -- with args", actual)
}

func Test_SplitLeftRightTypeTrimmed(t *testing.T) {
	// Arrange
	result := stringutil.SplitLeftRightTypeTrimmed(" key = val ", "=")

	// Act
	actual := args.Map{
		"left": result.Left,
		"right": result.Right,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "val",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTypeTrimmed returns correct value -- with args", actual)
}

func Test_SplitLeftRightsTrims_Empty(t *testing.T) {
	// Arrange
	result := stringutil.SplitLeftRightsTrims("=")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims returns empty -- empty", actual)
}

func Test_SplitLeftRightsTrims_Items(t *testing.T) {
	// Arrange
	result := stringutil.SplitLeftRightsTrims("=", " a = b ", " c = d ")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims returns correct value -- items", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToBool — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ToBool_All(t *testing.T) {
	// Act
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

	// Assert
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

func Test_ToByte_Valid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToByte("42", 0)}

	// Assert
	expected := args.Map{"v": byte(42)}
	expected.ShouldBeEqual(t, 0, "ToByte returns non-empty -- valid", actual)
}

func Test_ToByte_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToByte("abc", 99)}

	// Assert
	expected := args.Map{"v": byte(99)}
	expected.ShouldBeEqual(t, 0, "ToByte returns error -- invalid", actual)
}

func Test_ToByte_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToByte("999", 77)}

	// Assert
	expected := args.Map{"v": byte(77)}
	expected.ShouldBeEqual(t, 0, "ToByte returns correct value -- out of range", actual)
}

func Test_ToByteDefault_Valid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToByteDefault("42")}

	// Assert
	expected := args.Map{"v": byte(42)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns non-empty -- valid", actual)
}

func Test_ToByteDefault_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToByteDefault("abc")}

	// Assert
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns error -- invalid", actual)
}

func Test_ToByteDefault_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToByteDefault("999")}

	// Assert
	expected := args.Map{"v": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns correct value -- out of range", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToInt / ToIntDef / ToIntDefault / ToInt8 / ToInt8Def / ToInt16 / ToInt16Default / ToInt32 / ToInt32Def
// ══════════════════════════════════════════════════════════════════════════════

func Test_ToInt_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToInt("42", -1),
		"invalid": stringutil.ToInt("abc", -1),
	}

	// Assert
	expected := args.Map{
		"valid": 42,
		"invalid": -1,
	}
	expected.ShouldBeEqual(t, 0, "ToInt returns correct value -- with args", actual)
}

func Test_ToIntDef_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToIntDef("42"),
		"invalid": stringutil.ToIntDef("abc"),
	}

	// Assert
	expected := args.Map{
		"valid": 42,
		"invalid": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToIntDef returns correct value -- with args", actual)
}

func Test_ToIntDefault_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToIntDefault("42"),
		"invalid": stringutil.ToIntDefault("abc"),
	}

	// Assert
	expected := args.Map{
		"valid": 42,
		"invalid": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToIntDefault returns correct value -- with args", actual)
}

func Test_ToInt8_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToInt8("42", -1),
		"invalid": stringutil.ToInt8("abc", -1),
	}

	// Assert
	expected := args.Map{
		"valid": int8(42),
		"invalid": int8(-1),
	}
	expected.ShouldBeEqual(t, 0, "ToInt8 returns correct value -- with args", actual)
}

func Test_ToInt8Def_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToInt8Def("42"),
		"invalid": stringutil.ToInt8Def("abc"),
	}

	// Assert
	expected := args.Map{
		"valid": int8(42),
		"invalid": int8(0),
	}
	expected.ShouldBeEqual(t, 0, "ToInt8Def returns correct value -- with args", actual)
}

func Test_ToInt16_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToInt16("42", -1),
		"invalid": stringutil.ToInt16("abc", -1),
	}

	// Assert
	expected := args.Map{
		"valid": int16(42),
		"invalid": int16(-1),
	}
	expected.ShouldBeEqual(t, 0, "ToInt16 returns correct value -- with args", actual)
}

func Test_ToInt16Default_Valid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToInt16Default("100")}

	// Assert
	expected := args.Map{"v": int16(100)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns non-empty -- valid", actual)
}

func Test_ToInt16Default_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToInt16Default("abc")}

	// Assert
	expected := args.Map{"v": int16(0)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns error -- invalid", actual)
}

func Test_ToInt16Default_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToInt16Default("999999")}

	// Assert
	expected := args.Map{"v": int16(0)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns correct value -- out of range", actual)
}

func Test_ToInt32_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToInt32("42", -1),
		"invalid": stringutil.ToInt32("abc", -1),
	}

	// Assert
	expected := args.Map{
		"valid": int32(42),
		"invalid": int32(-1),
	}
	expected.ShouldBeEqual(t, 0, "ToInt32 returns correct value -- with args", actual)
}

func Test_ToInt32Def_FromAnyToStringNil(t *testing.T) {
	// Act
	actual := args.Map{
		"valid": stringutil.ToInt32Def("42"),
		"invalid": stringutil.ToInt32Def("abc"),
	}

	// Assert
	expected := args.Map{
		"valid": int32(42),
		"invalid": int32(0),
	}
	expected.ShouldBeEqual(t, 0, "ToInt32Def returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToUint16Default / ToUint32Default — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ToUint16Default_Valid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToUint16Default("100")}

	// Assert
	expected := args.Map{"v": uint16(100)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns non-empty -- valid", actual)
}

func Test_ToUint16Default_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToUint16Default("abc")}

	// Assert
	expected := args.Map{"v": uint16(0)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns error -- invalid", actual)
}

func Test_ToUint16Default_OutOfRange(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToUint16Default("999999")}

	// Assert
	expected := args.Map{"v": uint16(0)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns correct value -- out of range", actual)
}

func Test_ToUint32Default_Valid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToUint32Default("100")}

	// Assert
	expected := args.Map{"v": uint32(100)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns non-empty -- valid", actual)
}

func Test_ToUint32Default_Invalid(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToUint32Default("abc")}

	// Assert
	expected := args.Map{"v": uint32(0)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns error -- invalid", actual)
}

func Test_ToUint32Default_OutOfRange(t *testing.T) {
	// MaxInt32+1 would exceed the range
	// Act
	actual := args.Map{"v": stringutil.ToUint32Default("99999999999")}

	// Assert
	expected := args.Map{"v": uint32(0)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns correct value -- out of range", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ToIntUsingRegexMatch — all branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ToIntUsingRegexMatch_NilRegex(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ToIntUsingRegexMatch(nil, "42")}

	// Assert
	expected := args.Map{"v": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns nil -- nil", actual)
}

func Test_ToIntUsingRegexMatch_NoMatch(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^\d+$`)

	// Act
	actual := args.Map{"v": stringutil.ToIntUsingRegexMatch(re, "abc")}

	// Assert
	expected := args.Map{"v": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns empty -- no match", actual)
}

func Test_ToIntUsingRegexMatch_Valid(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^\d+$`)

	// Act
	actual := args.Map{"v": stringutil.ToIntUsingRegexMatch(re, "42")}

	// Assert
	expected := args.Map{"v": 42}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns non-empty -- valid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpacesToSingle (standalone)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReplaceWhiteSpacesToSingle_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceWhiteSpacesToSingle("")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns empty -- empty", actual)
}

func Test_ReplaceWhiteSpacesToSingle_Whitespace(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceWhiteSpacesToSingle("  hello   world  \t foo  ")}

	// Assert
	expected := args.Map{"v": "hello world foo"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- ws", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceTemplate — all methods and branches
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReplaceTemplate_CurlyOne_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyOne("", "k", "v")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyOne returns empty -- empty", actual)
}

func Test_ReplaceTemplate_CurlyOne_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyOne("hello {name}!", "name", "world")}

	// Assert
	expected := args.Map{"v": "hello world!"}
	expected.ShouldBeEqual(t, 0, "CurlyOne returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_Curly_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.Curly("", map[string]string{"k": "v"})}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "Curly returns empty -- empty", actual)
}

func Test_ReplaceTemplate_Curly_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.Curly("{a}-{b}", map[string]string{"a": "1", "b": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "Curly returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_CurlyTwo_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwo("", "a", 1, "b", 2)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyTwo returns empty -- empty", actual)
}

func Test_ReplaceTemplate_CurlyTwo_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwo("{a}-{b}", "a", 1, "b", 2)}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyTwo returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_DirectOne_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectOne("", "k", "v")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectOne returns empty -- empty", actual)
}

func Test_ReplaceTemplate_DirectOne_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectOne("hello KEY!", "KEY", "world")}

	// Assert
	expected := args.Map{"v": "hello world!"}
	expected.ShouldBeEqual(t, 0, "DirectOne returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_DirectTwoItem_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectTwoItem("", "a", 1, "b", 2)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectTwoItem returns empty -- empty", actual)
}

func Test_ReplaceTemplate_DirectTwoItem_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectTwoItem("A-B", "A", 1, "B", 2)}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectTwoItem returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_CurlyTwoItem_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwoItem("", "a", 1, "b", 2)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyTwoItem returns empty -- empty", actual)
}

func Test_ReplaceTemplate_CurlyTwoItem_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyTwoItem("{a}-{b}", "a", 1, "b", 2)}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyTwoItem returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingMap_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMap("", map[string]string{"k": "v"})}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap returns empty -- empty", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingMap_EmptyMap(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMap("hello", map[string]string{})}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap returns empty -- empty map", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingMap_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMap("A-B", map[string]string{"A": "1", "B": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingKeyVal_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingKeyVal("")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal returns empty -- empty", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingKeyVal_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingKeyVal("A-B",
		stringutil.KeyValReplacer{Key: "A", Value: "1"},
		stringutil.KeyValReplacer{Key: "B", Value: "2"},
	)}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingMapTrim(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.DirectKeyUsingMapTrim(" A-B ", map[string]string{"A": "1", "B": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMapTrim returns correct value -- with args", actual)
}

func Test_ReplaceTemplate_ReplaceWhiteSpaces_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces returns empty -- empty", actual)
}

func Test_ReplaceTemplate_ReplaceWhiteSpaces_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("  hello  world  ")}

	// Assert
	expected := args.Map{"v": "helloworld"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_ReplaceWhiteSpacesToSingle_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns empty -- empty", actual)
}

func Test_ReplaceTemplate_ReplaceWhiteSpacesToSingle_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("  hello   world  ")}

	// Assert
	expected := args.Map{"v": "hello world"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_CurlyKeyUsingMap_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("", map[string]string{"k": "v"})}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap returns empty -- empty", actual)
}

func Test_ReplaceTemplate_CurlyKeyUsingMap_EmptyMap(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("hello", map[string]string{})}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap returns empty -- empty map", actual)
}

func Test_ReplaceTemplate_UsingMapOptions_NonCurly(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingMapOptions(false, "A-B", map[string]string{"A": "1", "B": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions returns non-empty -- non-curly", actual)
}

func Test_ReplaceTemplate_UsingMapOptions_Curly(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingMapOptions(true, "{a}-{b}", map[string]string{"a": "1", "b": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions returns correct value -- curly", actual)
}

func Test_ReplaceTemplate_UsingMapOptions_EmptyMap(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingMapOptions(true, "hello", map[string]string{})}

	// Assert
	expected := args.Map{"v": "hello"}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions returns empty -- empty map", actual)
}

type testNamer struct{ name string }

func (n testNamer) Name() string { return n.name }

func Test_ReplaceTemplate_UsingNamerMapOptions_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingNamerMapOptions(true, "", nil)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingNamerMapOptions returns empty -- empty", actual)
}

func Test_ReplaceTemplate_UsingStringerMapOptions_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingStringerMapOptions(true, "", nil)}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingStringerMapOptions returns empty -- empty", actual)
}

type testStringer struct{ val string }

func (s testStringer) String() string { return s.val }

func Test_ReplaceTemplate_UsingStringerMapOptions_Curly(t *testing.T) {
	// Arrange
	m := map[fmt.Stringer]string{testStringer{"a"}: "1"}

	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingStringerMapOptions(true, "{a}", m)}

	// Assert
	expected := args.Map{"v": "1"}
	expected.ShouldBeEqual(t, 0, "UsingStringerMapOptions returns correct value -- curly", actual)
}

func Test_ReplaceTemplate_UsingStringerMapOptions_NonCurly(t *testing.T) {
	// Arrange
	m := map[fmt.Stringer]string{testStringer{"A"}: "1"}

	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingStringerMapOptions(false, "A-B", m)}

	// Assert
	expected := args.Map{"v": "1-B"}
	expected.ShouldBeEqual(t, 0, "UsingStringerMapOptions returns non-empty -- non-curly", actual)
}

func Test_ReplaceTemplate_UsingWrappedTemplate_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingWrappedTemplate("", "x")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingWrappedTemplate returns empty -- empty", actual)
}

func Test_ReplaceTemplate_UsingWrappedTemplate_Normal(t *testing.T) {
	// Act
	actual := args.Map{"has": stringutil.ReplaceTemplate.UsingWrappedTemplate("{wrapped}", "x") != ""}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingWrappedTemplate returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_UsingBracketsWrappedTemplate_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("", "x")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingBracketsWrappedTemplate returns empty -- empty", actual)
}

func Test_ReplaceTemplate_UsingQuotesWrappedTemplate_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate("", "x")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingQuotesWrappedTemplate returns empty -- empty", actual)
}

func Test_ReplaceTemplate_UsingValueTemplate_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingValueTemplate("", "x")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingValueTemplate returns empty -- empty", actual)
}

func Test_ReplaceTemplate_UsingValueTemplate_Normal(t *testing.T) {
	// Act
	actual := args.Map{"has": stringutil.ReplaceTemplate.UsingValueTemplate("test {value} end", "X") != ""}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingValueTemplate returns correct value -- normal", actual)
}

func Test_ReplaceTemplate_UsingValueWithFieldsTemplate_Empty(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.UsingValueWithFieldsTemplate("", "x")}

	// Assert
	expected := args.Map{"v": ""}
	expected.ShouldBeEqual(t, 0, "UsingValueWithFieldsTemplate returns empty -- empty", actual)
}

func Test_ReplaceTemplate_UsingValueWithFieldsTemplate_Normal(t *testing.T) {
	// Act
	actual := args.Map{"has": stringutil.ReplaceTemplate.UsingValueWithFieldsTemplate("test {value-fields} end", "X") != ""}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingValueWithFieldsTemplate returns non-empty -- normal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValReplacer — struct
// ══════════════════════════════════════════════════════════════════════════════

func Test_KeyValReplacer(t *testing.T) {
	// Arrange
	kvr := stringutil.KeyValReplacer{Key: "k", Value: "v"}

	// Act
	actual := args.Map{
		"key": kvr.Key,
		"val": kvr.Value,
	}

	// Assert
	expected := args.Map{
		"key": "k",
		"val": "v",
	}
	expected.ShouldBeEqual(t, 0, "KeyValReplacer returns correct value -- with args", actual)
}
