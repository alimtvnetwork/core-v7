package stringutiltests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ── IsEmpty / IsNotEmpty / IsDefined / IsBlank ──

func Test_Cov4_IsEmpty(t *testing.T) {
	actual := args.Map{"empty": stringutil.IsEmpty(""), "notEmpty": stringutil.IsEmpty("x")}
	expected := args.Map{"empty": true, "notEmpty": false}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_Cov4_IsNotEmpty(t *testing.T) {
	actual := args.Map{"notEmpty": stringutil.IsNotEmpty("x"), "empty": stringutil.IsNotEmpty("")}
	expected := args.Map{"notEmpty": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- with args", actual)
}

func Test_Cov4_IsEmptyPtr(t *testing.T) {
	empty := ""
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsEmptyPtr(nil),
		"empty": stringutil.IsEmptyPtr(&empty),
		"text":  stringutil.IsEmptyPtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyPtr returns empty -- with args", actual)
}

func Test_Cov4_IsNullOrEmptyPtr(t *testing.T) {
	empty := ""
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsNullOrEmptyPtr(nil),
		"empty": stringutil.IsNullOrEmptyPtr(&empty),
		"text":  stringutil.IsNullOrEmptyPtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- with args", actual)
}

func Test_Cov4_IsBlank(t *testing.T) {
	actual := args.Map{
		"empty":  stringutil.IsBlank(""),
		"space":  stringutil.IsBlank(" "),
		"nl":     stringutil.IsBlank("\n"),
		"text":   stringutil.IsBlank("x"),
		"tabs":   stringutil.IsBlank("\t  "),
	}
	expected := args.Map{"empty": true, "space": true, "nl": true, "text": false, "tabs": true}
	expected.ShouldBeEqual(t, 0, "IsBlank returns correct value -- with args", actual)
}

func Test_Cov4_IsEmptyOrWhitespace(t *testing.T) {
	actual := args.Map{
		"empty": stringutil.IsEmptyOrWhitespace(""),
		"space": stringutil.IsEmptyOrWhitespace("  "),
		"text":  stringutil.IsEmptyOrWhitespace("x"),
	}
	expected := args.Map{"empty": true, "space": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_Cov4_IsContains(t *testing.T) {
	lines := []string{"hello", "world"}
	actual := args.Map{
		"found":    stringutil.IsContains(lines, "world", 0, true),
		"notFound": stringutil.IsContains(lines, "foo", 0, true),
	}
	expected := args.Map{"found": true, "notFound": false}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- with args", actual)
}

// ── AnyToString ──

func Test_Cov4_AnyToString(t *testing.T) {
	actual := args.Map{
		"nil":    stringutil.AnyToString(nil),
		"string": stringutil.AnyToString("hello"),
		"int":    stringutil.AnyToString(42) != "",
	}
	expected := args.Map{"nil": "", "string": "hello", "int": true}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- with args", actual)
}

// ── SplitLeftRightType / SplitLeftRightTypeTrimmed / SplitLeftRightsTrims ──

func Test_Cov4_SplitLeftRightType(t *testing.T) {
	result := stringutil.SplitLeftRightType("key=value", "=")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightType returns correct value -- with args", actual)
}

func Test_Cov4_SplitLeftRightTypeTrimmed(t *testing.T) {
	result := stringutil.SplitLeftRightTypeTrimmed(" key = value ", "=")
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTypeTrimmed returns correct value -- with args", actual)
}

func Test_Cov4_SplitLeftRightsTrims(t *testing.T) {
	result := stringutil.SplitLeftRightsTrims("=", "a=1", "b=2")
	emptyResult := stringutil.SplitLeftRightsTrims("=")
	actual := args.Map{"len": len(result), "emptyLen": len(emptyResult)}
	expected := args.Map{"len": 2, "emptyLen": 0}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightsTrims returns correct value -- with args", actual)
}

// ── SplitContentsByWhitespaceConditions ──

func Test_Cov4_SplitContentsByWhitespaceConditions(t *testing.T) {
	result1 := stringutil.SplitContentsByWhitespaceConditions("hello world", true, true, true, false, false)
	result2 := stringutil.SplitContentsByWhitespaceConditions("hello world hello", false, true, false, true, true)
	result3 := stringutil.SplitContentsByWhitespaceConditions("  a  b  ", false, false, false, false, false)
	actual := args.Map{
		"trimSorted":    len(result1) > 0,
		"uniqueLower":   len(result2) > 0,
		"noFlags":       len(result3) > 0,
	}
	expected := args.Map{"trimSorted": true, "uniqueLower": true, "noFlags": true}
	expected.ShouldBeEqual(t, 0, "SplitContentsByWhitespaceConditions returns correct value -- with args", actual)
}

// ── ToIntUsingRegexMatch ──

func Test_Cov4_ToIntUsingRegexMatch(t *testing.T) {
	re := regexp.MustCompile(`^\d+$`)
	actual := args.Map{
		"valid":    stringutil.ToIntUsingRegexMatch(re, "42"),
		"invalid":  stringutil.ToIntUsingRegexMatch(re, "abc"),
		"nilRegex": stringutil.ToIntUsingRegexMatch(nil, "42"),
	}
	expected := args.Map{"valid": 42, "invalid": 0, "nilRegex": 0}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch returns correct value -- with args", actual)
}

// ── ReplaceTemplate ──

func Test_Cov4_ReplaceTemplate_CurlyOne(t *testing.T) {
	result := stringutil.ReplaceWhiteSpacesToSingle("Hello  World   !")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "Hello World !"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- with args", actual)
}

// Removed: stringutil.Replace var does not exist in source.
// Coverage for ReplaceWhiteSpacesToSingle is in Coverage3.
