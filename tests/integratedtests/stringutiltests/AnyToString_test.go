package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

func Test_AnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"string": stringutil.AnyToString("hello"),
		"nil":    stringutil.AnyToString(nil),
		"int":    stringutil.AnyToString(42),
	}
	expected := args.Map{
		"string": "hello",
		"nil":    "",
		"int":    "42",
	}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- with args", actual)
}

func Test_IsEmpty_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    stringutil.IsEmpty(""),
		"nonEmpty": stringutil.IsEmpty("hello"),
	}
	expected := args.Map{
		"empty":    true,
		"nonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns empty -- with args", actual)
}

func Test_IsNotEmpty_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    stringutil.IsNotEmpty(""),
		"nonEmpty": stringutil.IsNotEmpty("hello"),
	}
	expected := args.Map{
		"empty":    false,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "IsNotEmpty returns empty -- with args", actual)
}

func Test_IsBlank(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    stringutil.IsBlank(""),
		"space":    stringutil.IsBlank("   "),
		"text":     stringutil.IsBlank("hello"),
	}
	expected := args.Map{
		"empty":    true,
		"space":    true,
		"text":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsBlank returns correct value -- with args", actual)
}

func Test_IsEmptyOrWhitespace_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    stringutil.IsEmptyOrWhitespace(""),
		"space":    stringutil.IsEmptyOrWhitespace("  "),
		"tab":      stringutil.IsEmptyOrWhitespace("\t"),
		"text":     stringutil.IsEmptyOrWhitespace("hello"),
	}
	expected := args.Map{
		"empty":    true,
		"space":    true,
		"tab":      true,
		"text":     false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_IsDefined(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":    stringutil.IsDefined(""),
		"nonEmpty": stringutil.IsDefined("hello"),
	}
	expected := args.Map{
		"empty":    false,
		"nonEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "IsDefined returns correct value -- with args", actual)
}

func Test_IsContains(t *testing.T) {
	// Act
	actual := args.Map{
		"found":    stringutil.IsContains([]string{"hello", "world"}, "world", 0, true),
		"notFound": stringutil.IsContains([]string{"hello", "world"}, "foo", 0, true),
	}
	expected := args.Map{
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsContains returns correct value -- with args", actual)
}

func Test_IsStarts_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"starts": stringutil.IsStarts("hello world", "hello"),
		"not":    stringutil.IsStarts("hello world", "world"),
	}
	expected := args.Map{
		"starts": true,
		"not":    false,
	}
	expected.ShouldBeEqual(t, 0, "IsStarts returns correct value -- with args", actual)
}

func Test_IsEnds_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"ends": stringutil.IsEnds("hello world", "world"),
		"not":  stringutil.IsEnds("hello world", "hello"),
	}
	expected := args.Map{
		"ends": true,
		"not":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsEnds returns correct value -- with args", actual)
}

func Test_FirstChar_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"first": stringutil.FirstChar("hello"),
		"empty": stringutil.FirstChar(""),
	}
	expected := args.Map{
		"first": byte('h'),
		"empty": byte(0),
	}
	expected.ShouldBeEqual(t, 0, "FirstChar returns correct value -- with args", actual)
}

func Test_ToInt_FromAnyToString(t *testing.T) {
	// Act
	val := stringutil.ToInt("42", -1)
	valBad := stringutil.ToInt("abc", -1)

	actual := args.Map{
		"val":    val,
		"valBad": valBad,
	}
	expected := args.Map{
		"val":    42,
		"valBad": -1,
	}
	expected.ShouldBeEqual(t, 0, "ToInt returns correct value -- with args", actual)
}

func Test_ToIntDefault_FromAnyToString(t *testing.T) {
	// Act
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

func Test_ToBool_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"true":    stringutil.ToBool("true"),
		"false":   stringutil.ToBool("false"),
		"invalid": stringutil.ToBool("abc"),
	}
	expected := args.Map{
		"true":    true,
		"false":   false,
		"invalid": false,
	}
	expected.ShouldBeEqual(t, 0, "ToBool returns correct value -- with args", actual)
}

func Test_SplitLeftRight_FromAnyToString(t *testing.T) {
	// Act
	left, right := stringutil.SplitLeftRight("key=value", "=")
	left2, right2 := stringutil.SplitLeftRight("noequals", "=")

	actual := args.Map{
		"left":   left,
		"right":  right,
		"left2":  left2,
		"right2": right2,
	}
	expected := args.Map{
		"left":   "key",
		"right":  "value",
		"left2":  "noequals",
		"right2": "",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- with args", actual)
}

func Test_MaskLine_FromAnyToString(t *testing.T) {
	// Act
	result := stringutil.MaskLine("XXXXXXXXXX", "abc")

	actual := args.Map{
		"result": result,
	}
	expected := args.Map{
		"result": "abcXXXXXXX",
	}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- with args", actual)
}

func Test_SafeSubstring_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"normal": stringutil.SafeSubstring("hello world", 0, 5),
		"over":   stringutil.SafeSubstring("hi", 0, 10),
		"empty":  stringutil.SafeSubstring("", 0, 5),
	}
	expected := args.Map{
		"normal": "hello",
		"over":   "",
		"empty":  "",
	}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- with args", actual)
}

func Test_IsNullOrEmptyPtr_FromAnyToString(t *testing.T) {
	// Arrange
	empty := ""
	hello := "hello"

	// Act
	actual := args.Map{
		"nil":      stringutil.IsNullOrEmptyPtr(nil),
		"empty":    stringutil.IsNullOrEmptyPtr(&empty),
		"nonEmpty": stringutil.IsNullOrEmptyPtr(&hello),
	}
	expected := args.Map{
		"nil":      true,
		"empty":    true,
		"nonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyPtr returns empty -- with args", actual)
}

func Test_IsAnyStartsWith_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"found":    stringutil.IsAnyStartsWith("hello", false, "he", "wo"),
		"notFound": stringutil.IsAnyStartsWith("hello", false, "wo", "fo"),
	}
	expected := args.Map{
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns non-empty -- with args", actual)
}

func Test_IsAnyEndsWith_FromAnyToString(t *testing.T) {
	// Act
	actual := args.Map{
		"found":    stringutil.IsAnyEndsWith("hello", false, "lo", "wo"),
		"notFound": stringutil.IsAnyEndsWith("hello", false, "wo", "fo"),
	}
	expected := args.Map{
		"found":    true,
		"notFound": false,
	}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns non-empty -- with args", actual)
}

func Test_RemoveMany_FromAnyToString(t *testing.T) {
	// Act
	result := stringutil.RemoveMany("hello world foo", "world", "foo")

	actual := args.Map{
		"result": result,
	}
	expected := args.Map{
		"result": "hello  ",
	}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns correct value -- with args", actual)
}
