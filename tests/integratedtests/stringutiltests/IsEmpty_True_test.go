package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ── IsEmpty / IsNotEmpty ──

func Test_IsEmpty_True(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmpty("")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- empty string", actual)
}

func Test_IsEmpty_False(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmpty("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- non-empty", actual)
}

func Test_IsNotEmpty_True(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsNotEmpty("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- non-empty", actual)
}

func Test_IsNotEmpty_False(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsNotEmpty("")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- empty", actual)
}

// ── IsBlank ──

func Test_IsBlank_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsBlank("")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlank returns empty -- empty", actual)
}

func Test_IsBlank_Whitespace(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsBlank("   ")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlank returns correct value -- whitespace only", actual)
}

func Test_IsBlank_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsBlank("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsBlank returns empty -- non-empty", actual)
}

// ── IsEmptyPtr / IsBlankPtr ──

func Test_IsEmptyPtr_Nil_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmptyPtr(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns nil -- nil", actual)
}

func Test_IsEmptyPtr_Empty_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	s := ""

	// Act
	actual := args.Map{"result": stringutil.IsEmptyPtr(&s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- empty", actual)
}

func Test_IsEmptyPtr_NonEmpty(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	actual := args.Map{"result": stringutil.IsEmptyPtr(&s)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- non-empty", actual)
}

func Test_IsBlankPtr_Nil_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsBlankPtr(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns nil -- nil", actual)
}

func Test_IsBlankPtr_Whitespace(t *testing.T) {
	// Arrange
	s := "   "

	// Act
	actual := args.Map{"result": stringutil.IsBlankPtr(&s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns correct value -- whitespace", actual)
}

// ── IsEmptyOrWhitespace / IsEmptyOrWhitespacePtr ──

func Test_IsEmptyOrWhitespace_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespace("")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- empty", actual)
}

func Test_IsEmptyOrWhitespace_Whitespace(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespace("   ")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- whitespace", actual)
}

func Test_IsEmptyOrWhitespace_NonEmpty(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespace("hello")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- non-empty", actual)
}

func Test_IsEmptyOrWhitespacePtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsEmptyOrWhitespacePtr(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr returns nil -- nil", actual)
}

// ── IsDefined / IsDefinedPtr ──

func Test_IsDefined_True(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsDefined("hello")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDefined returns empty -- non-empty", actual)
}

func Test_IsDefined_False(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsDefined("")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsDefined returns empty -- empty", actual)
}

func Test_IsDefinedPtr_Nil_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsDefinedPtr(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns nil -- nil", actual)
}

func Test_IsDefinedPtr_NonEmpty(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	actual := args.Map{"result": stringutil.IsDefinedPtr(&s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns empty -- non-empty", actual)
}

// ── IsNullOrEmptyPtr ──

func Test_IsNullOrEmptyPtr_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": stringutil.IsNullOrEmptyPtr(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns nil -- nil", actual)
}

func Test_IsNullOrEmptyPtr_Empty(t *testing.T) {
	// Arrange
	s := ""

	// Act
	actual := args.Map{"result": stringutil.IsNullOrEmptyPtr(&s)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- empty", actual)
}

func Test_IsNullOrEmptyPtr_NonEmpty(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	actual := args.Map{"result": stringutil.IsNullOrEmptyPtr(&s)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- non-empty", actual)
}

// ── IsStarts / IsEnds / IsContains ──

func Test_IsStarts_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsStarts("hello world", "hello"),
		"noMatch": stringutil.IsStarts("hello world", "world"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStarts returns correct value -- with args", actual)
}

func Test_IsEnds_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsEnds("hello world", "world"),
		"noMatch": stringutil.IsEnds("hello world", "hello"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEnds returns correct value -- with args", actual)
}

func Test_IsContains_Slice(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsContains([]string{"hello", "world"}, "world", 0, true),
		"noMatch": stringutil.IsContains([]string{"hello", "world"}, "xyz", 0, true),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- with args", actual)
}

// ── IsStartsWith / IsEndsWith ──

func Test_IsStartsWith_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsStartsWith("hello", "hel", false),
		"noMatch": stringutil.IsStartsWith("hello", "xyz", false),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- with args", actual)
}

func Test_IsEndsWith_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsEndsWith("hello", "llo", false),
		"noMatch": stringutil.IsEndsWith("hello", "xyz", false),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- with args", actual)
}

// ── IsStartsChar / IsEndsChar / IsStartsRune / IsEndsRune ──

func Test_IsStartsChar_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsStartsChar("hello", 'h'),
		"noMatch": stringutil.IsStartsChar("hello", 'x'),
		"empty":   stringutil.IsStartsChar("", 'h'),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"empty":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsChar returns correct value -- with args", actual)
}

func Test_IsEndsChar_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsEndsChar("hello", 'o'),
		"noMatch": stringutil.IsEndsChar("hello", 'x'),
		"empty":   stringutil.IsEndsChar("", 'o'),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
		"empty":   false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns correct value -- with args", actual)
}

func Test_IsStartsRune_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsStartsRune("hello", 'h'),
		"noMatch": stringutil.IsStartsRune("hello", 'x'),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsRune returns correct value -- with args", actual)
}

func Test_IsEndsRune_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsEndsRune("hello", 'o'),
		"noMatch": stringutil.IsEndsRune("hello", 'x'),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEndsRune returns correct value -- with args", actual)
}

// ── IsStartsAndEndsChar / IsStartsAndEndsWith ──

func Test_IsStartsAndEndsChar_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsStartsAndEndsChar("[hello]", '[', ']'),
		"noMatch": stringutil.IsStartsAndEndsChar("hello", '[', ']'),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns correct value -- with args", actual)
}

func Test_IsStartsAndEndsWith_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsStartsAndEndsWith("<<hello>>", "<<", ">>", false),
		"noMatch": stringutil.IsStartsAndEndsWith("hello", "<<", ">>", false),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith returns non-empty -- with args", actual)
}

// ── IsAnyStartsWith / IsAnyEndsWith ──

func Test_IsAnyStartsWith_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsAnyStartsWith("hello", false, "xyz", "hel"),
		"noMatch": stringutil.IsAnyStartsWith("hello", false, "xyz", "abc"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns non-empty -- with args", actual)
}

func Test_IsAnyEndsWith_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"match":   stringutil.IsAnyEndsWith("hello", false, "xyz", "llo"),
		"noMatch": stringutil.IsAnyEndsWith("hello", false, "xyz", "abc"),
	}

	// Assert
	expected := args.Map{
		"match":   true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns non-empty -- with args", actual)
}

// ── IsContainsPtr / IsContainsPtrSimple ──

func Test_IsContainsPtr_Nil(t *testing.T) {
	// Arrange
	find := "hello"

	// Act
	actual := args.Map{"result": stringutil.IsContainsPtr(nil, &find, 0, true)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns nil -- nil", actual)
}

func Test_IsContainsPtr_Match(t *testing.T) {
	// Arrange
	s := []string{"hello", "world"}
	find := "world"

	// Act
	actual := args.Map{"result": stringutil.IsContainsPtr(&s, &find, 0, true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- match", actual)
}

// ── ClonePtr / SafeClonePtr ──

func Test_ClonePtr_Nil_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{"isNil": stringutil.ClonePtr(nil) == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- nil", actual)
}

func Test_ClonePtr_NonNil_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	s := "hello"
	cloned := stringutil.ClonePtr(&s)

	// Act
	actual := args.Map{
		"value": *cloned,
		"diffPtr": cloned != &s,
	}

	// Assert
	expected := args.Map{
		"value": "hello",
		"diffPtr": true,
	}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns nil -- non-nil", actual)
}

func Test_SafeClonePtr_Nil_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	result := stringutil.SafeClonePtr(nil)

	// Act
	actual := args.Map{
		"notNil": result != nil,
		"value": *result,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"value": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns nil -- nil", actual)
}

// ── FirstChar ──

func Test_FirstChar_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"normal": stringutil.FirstChar("hello"),
		"empty":  stringutil.FirstChar(""),
	}

	// Assert
	expected := args.Map{
		"normal": byte('h'),
		"empty":  byte(0),
	}
	expected.ShouldBeEqual(t, 0, "FirstChar returns correct value -- with args", actual)
}

// ── SafeSubstring ──

func Test_SafeSubstring_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"normal":    stringutil.SafeSubstring("hello", 1, 3),
		"outOfBound": stringutil.SafeSubstring("hello", 0, 100),
	}

	// Assert
	expected := args.Map{
		"normal":    "el",
		"outOfBound": "",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- with args", actual)
}

// ── ToBool ──

func Test_ToBool_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"true":    stringutil.ToBool("true"),
		"false":   stringutil.ToBool("false"),
		"empty":   stringutil.ToBool(""),
		"invalid": stringutil.ToBool("abc"),
	}

	// Assert
	expected := args.Map{
		"true":    true,
		"false":   false,
		"empty":   false,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "ToBool returns correct value -- with args", actual)
}

// ── ToInt / ToIntDef / ToIntDefault ──

func Test_ToInt_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	result := stringutil.ToInt("42", -1)

	// Act
	actual := args.Map{"value": result}

	// Assert
	expected := args.Map{"value": 42}
	expected.ShouldBeEqual(t, 0, "ToInt returns non-empty -- valid", actual)
}

func Test_ToInt_Invalid(t *testing.T) {
	// Arrange
	result := stringutil.ToInt("abc", -1)

	// Act
	actual := args.Map{"value": result}

	// Assert
	expected := args.Map{"value": -1}
	expected.ShouldBeEqual(t, 0, "ToInt returns error -- invalid", actual)
}

func Test_ToIntDefault_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"valid":   stringutil.ToIntDefault("42"),
		"invalid": stringutil.ToIntDefault("abc"),
	}

	// Assert
	expected := args.Map{
		"valid":   42,
		"invalid": 0,
	}
	expected.ShouldBeEqual(t, 0, "ToIntDefault returns correct value -- with args", actual)
}

// ── AnyToString ──

func Test_AnyToString_FromIsEmptyTrue(t *testing.T) {
	// Act
	actual := args.Map{
		"int":    stringutil.AnyToString(42),
		"string": stringutil.AnyToString("hello"),
		"bool":   stringutil.AnyToString(true),
	}

	// Assert
	expected := args.Map{
		"int":    "42",
		"string": "hello",
		"bool":   "true",
	}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- with args", actual)
}

// ── SplitLeftRight ──

func Test_SplitLeftRight_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRight("hello:world", ":")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "hello",
		"right": "world",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- with args", actual)
}

func Test_SplitLeftRight_NoSep_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRight("hello", ":")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "hello",
		"right": "",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns empty -- no sep", actual)
}

// ── SplitLeftRightTrimmed ──

func Test_SplitLeftRightTrimmed_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	left, right := stringutil.SplitLeftRightTrimmed(" hello : world ", ":")

	// Act
	actual := args.Map{
		"left": left,
		"right": right,
	}

	// Assert
	expected := args.Map{
		"left": "hello",
		"right": "world",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns correct value -- with args", actual)
}

// ── RemoveMany ──

func Test_RemoveMany_FromIsEmptyTrue(t *testing.T) {
	// Arrange
	result := stringutil.RemoveMany("hello world foo", "world", "foo")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns correct value -- with args", actual)
}
