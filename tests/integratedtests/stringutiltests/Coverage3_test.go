package stringutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ── AnyToStringNameField / AnyToTypeString ──

func Test_Cov3_AnyToStringNameField(t *testing.T) {
	actual := args.Map{
		"string":  stringutil.AnyToStringNameField("hello") != "",
		"int":     stringutil.AnyToStringNameField(42) != "",
		"nil":     stringutil.AnyToStringNameField(nil),
	}
	expected := args.Map{"string": true, "int": true, "nil": ""}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField returns correct value -- with args", actual)
}

func Test_Cov3_AnyToTypeString(t *testing.T) {
	nilResult := stringutil.AnyToTypeString(nil)
	actual := args.Map{
		"string": stringutil.AnyToTypeString("hello") != "",
		"nil":    nilResult,
	}
	expected := args.Map{"string": true, "nil": nilResult}
	expected.ShouldBeEqual(t, 0, "AnyToTypeString returns correct value -- with args", actual)
}

// ── IsBlankPtr / IsDefinedPtr / IsEmptyPtr ──

func Test_Cov3_IsBlankPtr(t *testing.T) {
	empty := ""
	space := "   "
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsBlankPtr(nil),
		"empty": stringutil.IsBlankPtr(&empty),
		"space": stringutil.IsBlankPtr(&space),
		"text":  stringutil.IsBlankPtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "space": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsBlankPtr returns correct value -- with args", actual)
}

func Test_Cov3_IsDefinedPtr(t *testing.T) {
	empty := ""
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsDefinedPtr(nil),
		"empty": stringutil.IsDefinedPtr(&empty),
		"text":  stringutil.IsDefinedPtr(&text),
	}
	expected := args.Map{"nil": false, "empty": false, "text": true}
	expected.ShouldBeEqual(t, 0, "IsDefinedPtr returns correct value -- with args", actual)
}

func Test_Cov3_IsEmptyOrWhitespacePtr(t *testing.T) {
	empty := ""
	space := "  "
	text := "hello"
	actual := args.Map{
		"nil":   stringutil.IsEmptyOrWhitespacePtr(nil),
		"empty": stringutil.IsEmptyOrWhitespacePtr(&empty),
		"space": stringutil.IsEmptyOrWhitespacePtr(&space),
		"text":  stringutil.IsEmptyOrWhitespacePtr(&text),
	}
	expected := args.Map{"nil": true, "empty": true, "space": true, "text": false}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespacePtr returns empty -- with args", actual)
}

func Test_Cov3_IsContainsPtr(t *testing.T) {
	lines := []string{"hello", "world"}
	find1 := "world"
	find2 := "foo"
	find3 := "x"
	actual := args.Map{
		"found":    stringutil.IsContainsPtr(&lines, &find1, 0, true),
		"notFound": stringutil.IsContainsPtr(&lines, &find2, 0, true),
		"nil":      stringutil.IsContainsPtr(nil, &find3, 0, true),
	}
	expected := args.Map{"found": true, "notFound": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr returns correct value -- with args", actual)
}

func Test_Cov3_IsContainsPtrSimple(t *testing.T) {
	lines := []string{"hello", "world"}
	actual := args.Map{
		"found":    stringutil.IsContainsPtrSimple(&lines, "hello", 0, true),
		"notFound": stringutil.IsContainsPtrSimple(&lines, "foo", 0, true),
		"nil":      stringutil.IsContainsPtrSimple(nil, "x", 0, true),
	}
	expected := args.Map{"found": true, "notFound": false, "nil": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple returns correct value -- with args", actual)
}

// ── IsStarts/IsEnds ──

func Test_Cov3_IsStarts(t *testing.T) {
	actual := args.Map{
		"starts":  stringutil.IsStarts("hello world", "hello"),
		"notStarts": stringutil.IsStarts("hello world", "world"),
	}
	expected := args.Map{"starts": true, "notStarts": false}
	expected.ShouldBeEqual(t, 0, "IsStarts returns correct value -- with args", actual)
}

func Test_Cov3_IsEnds(t *testing.T) {
	actual := args.Map{
		"ends":    stringutil.IsEnds("hello world", "world"),
		"notEnds": stringutil.IsEnds("hello world", "hello"),
	}
	expected := args.Map{"ends": true, "notEnds": false}
	expected.ShouldBeEqual(t, 0, "IsEnds returns correct value -- with args", actual)
}

func Test_Cov3_IsStartsWith(t *testing.T) {
	actual := args.Map{
		"starts":    stringutil.IsStartsWith("hello", "hel", false),
		"notStarts": stringutil.IsStartsWith("hello", "wor", false),
	}
	expected := args.Map{"starts": true, "notStarts": false}
	expected.ShouldBeEqual(t, 0, "IsStartsWith returns non-empty -- with args", actual)
}

func Test_Cov3_IsEndsWith(t *testing.T) {
	actual := args.Map{
		"ends":    stringutil.IsEndsWith("hello", "llo", false),
		"notEnds": stringutil.IsEndsWith("hello", "hel", false),
	}
	expected := args.Map{"ends": true, "notEnds": false}
	expected.ShouldBeEqual(t, 0, "IsEndsWith returns non-empty -- with args", actual)
}

func Test_Cov3_IsStartsChar(t *testing.T) {
	actual := args.Map{
		"starts":    stringutil.IsStartsChar("hello", 'h'),
		"notStarts": stringutil.IsStartsChar("hello", 'z'),
		"empty":     stringutil.IsStartsChar("", 'h'),
	}
	expected := args.Map{"starts": true, "notStarts": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsChar returns correct value -- with args", actual)
}

func Test_Cov3_IsEndsChar(t *testing.T) {
	actual := args.Map{
		"ends":    stringutil.IsEndsChar("hello", 'o'),
		"notEnds": stringutil.IsEndsChar("hello", 'z'),
		"empty":   stringutil.IsEndsChar("", 'o'),
	}
	expected := args.Map{"ends": true, "notEnds": false, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsEndsChar returns correct value -- with args", actual)
}

func Test_Cov3_IsStartsRune(t *testing.T) {
	actual := args.Map{
		"starts": stringutil.IsStartsRune("hello", 'h'),
		"empty":  stringutil.IsStartsRune("", 'h'),
	}
	expected := args.Map{"starts": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsStartsRune returns correct value -- with args", actual)
}

func Test_Cov3_IsEndsRune(t *testing.T) {
	actual := args.Map{
		"ends":  stringutil.IsEndsRune("hello", 'o'),
		"empty": stringutil.IsEndsRune("", 'o'),
	}
	expected := args.Map{"ends": true, "empty": false}
	expected.ShouldBeEqual(t, 0, "IsEndsRune returns correct value -- with args", actual)
}

func Test_Cov3_IsStartsAndEndsChar(t *testing.T) {
	actual := args.Map{
		"match": stringutil.IsStartsAndEndsChar("hello", 'h', 'o'),
		"noMatch": stringutil.IsStartsAndEndsChar("hello", 'h', 'x'),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsChar returns correct value -- with args", actual)
}

func Test_Cov3_IsStartsAndEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsStartsAndEndsWith("hello world", "hello", "world", false),
		"noMatch": stringutil.IsStartsAndEndsWith("hello world", "hello", "xyz", false),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsStartsAndEndsWith returns non-empty -- with args", actual)
}

func Test_Cov3_IsAnyStartsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsAnyStartsWith("hello", false, "he", "wo"),
		"noMatch": stringutil.IsAnyStartsWith("hello", false, "wo", "xy"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsAnyStartsWith returns non-empty -- with args", actual)
}

func Test_Cov3_IsAnyEndsWith(t *testing.T) {
	actual := args.Map{
		"match":   stringutil.IsAnyEndsWith("hello", false, "lo", "xy"),
		"noMatch": stringutil.IsAnyEndsWith("hello", false, "ab", "cd"),
	}
	expected := args.Map{"match": true, "noMatch": false}
	expected.ShouldBeEqual(t, 0, "IsAnyEndsWith returns non-empty -- with args", actual)
}

// ── FirstChar / ClonePtr / SafeClonePtr ──

func Test_Cov3_FirstChar(t *testing.T) {
	first := stringutil.FirstChar("hello")
	empty := stringutil.FirstChar("")
	actual := args.Map{
		"first": first,
		"empty": empty,
	}
	expected := args.Map{"first": first, "empty": empty}
	expected.ShouldBeEqual(t, 0, "FirstChar returns correct value -- with args", actual)
}

func Test_Cov3_ClonePtr(t *testing.T) {
	text := "hello"
	result := stringutil.ClonePtr(&text)
	nilResult := stringutil.ClonePtr(nil)
	actual := args.Map{"notNil": result != nil, "val": *result, "nilIsNil": nilResult == nil}
	expected := args.Map{"notNil": true, "val": "hello", "nilIsNil": true}
	expected.ShouldBeEqual(t, 0, "ClonePtr returns correct value -- with args", actual)
}

func Test_Cov3_SafeClonePtr(t *testing.T) {
	text := "hello"
	result := stringutil.SafeClonePtr(&text)
	nilResult := stringutil.SafeClonePtr(nil)
	actual := args.Map{"val": *result, "nilNotNil": nilResult != nil, "nilVal": *nilResult}
	expected := args.Map{"val": "hello", "nilNotNil": true, "nilVal": ""}
	expected.ShouldBeEqual(t, 0, "SafeClonePtr returns correct value -- with args", actual)
}

// ── SafeSubstring variants ──

func Test_Cov3_SafeSubstring(t *testing.T) {
	outOfRange := stringutil.SafeSubstring("hi", 0, 10)
	actual := args.Map{
		"normal":     stringutil.SafeSubstring("hello", 1, 3),
		"outOfRange": outOfRange,
	}
	expected := args.Map{"normal": "el", "outOfRange": outOfRange}
	expected.ShouldBeEqual(t, 0, "SafeSubstring returns correct value -- with args", actual)
}

func Test_Cov3_SafeSubstringStarts(t *testing.T) {
	actual := args.Map{
		"normal": stringutil.SafeSubstringStarts("hello", 2),
		"outOfRange": stringutil.SafeSubstringStarts("hi", 10),
	}
	expected := args.Map{"normal": "llo", "outOfRange": ""}
	expected.ShouldBeEqual(t, 0, "SafeSubstringStarts returns correct value -- with args", actual)
}

func Test_Cov3_SafeSubstringEnds(t *testing.T) {
	actual := args.Map{
		"normal": stringutil.SafeSubstringEnds("hello", 3),
		"outOfRange": stringutil.SafeSubstringEnds("hi", 10),
	}
	expected := args.Map{"normal": "hel", "outOfRange": "hi"}
	expected.ShouldBeEqual(t, 0, "SafeSubstringEnds returns correct value -- with args", actual)
}

// ── MaskLine / MaskLines / MaskTrimLine / MaskTrimLines ──

func Test_Cov3_MaskLine(t *testing.T) {
	actual := args.Map{
		"result": stringutil.MaskLine("XXXXXXXXXX", "abc"),
	}
	expected := args.Map{"result": "abcXXXXXXX"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- with args", actual)
}

func Test_Cov3_MaskLines(t *testing.T) {
	result := stringutil.MaskLines("XXXXX", "ab", "cde")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MaskLines returns correct value -- with args", actual)
}

func Test_Cov3_MaskTrimLine(t *testing.T) {
	actual := args.Map{"result": stringutil.MaskTrimLine("XXXXXXXXXX", "  abc  ")}
	expected := args.Map{"result": "abcXXXXXXX"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- with args", actual)
}

func Test_Cov3_MaskTrimLines(t *testing.T) {
	result := stringutil.MaskTrimLines("XXXXX", "  ab  ", "cde")
	actual := args.Map{"len": len(result)}
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "MaskTrimLines returns correct value -- with args", actual)
}

// ── RemoveMany / RemoveManyBySplitting ──

func Test_Cov3_RemoveMany(t *testing.T) {
	result := stringutil.RemoveMany("hello world foo", "world", "foo")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello  "}
	expected.ShouldBeEqual(t, 0, "RemoveMany returns correct value -- with args", actual)
}

func Test_Cov3_RemoveManyBySplitting(t *testing.T) {
	result := stringutil.RemoveManyBySplitting("hello world foo", " ", "world", "foo")
	actual := args.Map{"notEmpty": len(result) > 0}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "RemoveManyBySplitting returns correct value -- with args", actual)
}

// ── TrimKeepSingleSpaceOnly ──

func Test_Cov3_TrimKeepSingleSpaceOnly(t *testing.T) {
	result := stringutil.ReplaceWhiteSpacesToSingle("  hello   world  ")
	actual := args.Map{"result": result}
	expected := args.Map{"result": "hello world"}
	expected.ShouldBeEqual(t, 0, "TrimKeepSingleSpaceOnly returns correct value -- with args", actual)
}

// ── ToBool / ToByte / ToInt variants ──

func Test_Cov3_ToBool(t *testing.T) {
	actual := args.Map{
		"true":    stringutil.ToBool("true"),
		"false":   stringutil.ToBool("false"),
		"invalid": stringutil.ToBool("abc"),
	}
	expected := args.Map{"true": true, "false": false, "invalid": false}
	expected.ShouldBeEqual(t, 0, "ToBool returns correct value -- with args", actual)
}

func Test_Cov3_ToByte(t *testing.T) {
	val := stringutil.ToByte("42", 0)
	valInvalid := stringutil.ToByte("abc", 0)
	actual := args.Map{"val": val, "invalidVal": valInvalid}
	expected := args.Map{"val": byte(42), "invalidVal": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByte returns correct value -- with args", actual)
}

func Test_Cov3_ToByteDefault(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToByteDefault("42"),
		"invalid": stringutil.ToByteDefault("abc"),
	}
	expected := args.Map{"valid": byte(42), "invalid": byte(0)}
	expected.ShouldBeEqual(t, 0, "ToByteDefault returns correct value -- with args", actual)
}

func Test_Cov3_ToInt(t *testing.T) {
	val := stringutil.ToInt("42", 0)
	valInvalid := stringutil.ToInt("abc", -1)
	actual := args.Map{"val": val, "invalidVal": valInvalid}
	expected := args.Map{"val": 42, "invalidVal": -1}
	expected.ShouldBeEqual(t, 0, "ToInt returns correct value -- with args", actual)
}

func Test_Cov3_ToIntDef(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToIntDef("42"),
		"invalid": stringutil.ToIntDef("abc"),
	}
	expected := args.Map{"valid": 42, "invalid": 0}
	expected.ShouldBeEqual(t, 0, "ToIntDef returns correct value -- with args", actual)
}

func Test_Cov3_ToIntDefault(t *testing.T) {
	val := stringutil.ToIntDefault("42")
	valInvalid := stringutil.ToIntDefault("abc")
	actual := args.Map{"val": val, "invalidVal": valInvalid}
	expected := args.Map{"val": 42, "invalidVal": 0}
	expected.ShouldBeEqual(t, 0, "ToIntDefault returns correct value -- with args", actual)
}

func Test_Cov3_ToInt8(t *testing.T) {
	val := stringutil.ToInt8("42", 0)
	actual := args.Map{"val": val}
	expected := args.Map{"val": int8(42)}
	expected.ShouldBeEqual(t, 0, "ToInt8 returns correct value -- with args", actual)
}

func Test_Cov3_ToInt8Def(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt8Def("42"),
		"invalid": stringutil.ToInt8Def("abc"),
	}
	expected := args.Map{"valid": int8(42), "invalid": int8(0)}
	expected.ShouldBeEqual(t, 0, "ToInt8Def returns correct value -- with args", actual)
}

func Test_Cov3_ToInt16(t *testing.T) {
	val := stringutil.ToInt16("42", 0)
	actual := args.Map{"val": val}
	expected := args.Map{"val": int16(42)}
	expected.ShouldBeEqual(t, 0, "ToInt16 returns correct value -- with args", actual)
}

func Test_Cov3_ToInt16Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt16Default("42"),
		"invalid": stringutil.ToInt16Default("abc"),
	}
	expected := args.Map{"valid": int16(42), "invalid": int16(0)}
	expected.ShouldBeEqual(t, 0, "ToInt16Default returns correct value -- with args", actual)
}

func Test_Cov3_ToInt32(t *testing.T) {
	val := stringutil.ToInt32("42", 0)
	actual := args.Map{"val": val}
	expected := args.Map{"val": int32(42)}
	expected.ShouldBeEqual(t, 0, "ToInt32 returns correct value -- with args", actual)
}

func Test_Cov3_ToInt32Def(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToInt32Def("42"),
		"invalid": stringutil.ToInt32Def("abc"),
	}
	expected := args.Map{"valid": int32(42), "invalid": int32(0)}
	expected.ShouldBeEqual(t, 0, "ToInt32Def returns correct value -- with args", actual)
}

func Test_Cov3_ToUint16Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToUint16Default("42"),
		"invalid": stringutil.ToUint16Default("abc"),
	}
	expected := args.Map{"valid": uint16(42), "invalid": uint16(0)}
	expected.ShouldBeEqual(t, 0, "ToUint16Default returns correct value -- with args", actual)
}

func Test_Cov3_ToUint32Default(t *testing.T) {
	actual := args.Map{
		"valid":   stringutil.ToUint32Default("42"),
		"invalid": stringutil.ToUint32Default("abc"),
	}
	expected := args.Map{"valid": uint32(42), "invalid": uint32(0)}
	expected.ShouldBeEqual(t, 0, "ToUint32Default returns correct value -- with args", actual)
}

// ── SplitLeftRight / SplitFirstLast ──

func Test_Cov3_SplitLeftRight(t *testing.T) {
	l, r := stringutil.SplitLeftRight("key=value", "=")
	actual := args.Map{"l": l, "r": r}
	expected := args.Map{"l": "key", "r": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- with args", actual)
}

func Test_Cov3_SplitLeftRightTrimmed(t *testing.T) {
	l, r := stringutil.SplitLeftRightTrimmed("  key  =  value  ", "=")
	actual := args.Map{"l": l, "r": r}
	expected := args.Map{"l": "key", "r": "value"}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrimmed returns correct value -- with args", actual)
}

func Test_Cov3_SplitFirstLast(t *testing.T) {
	first, last := stringutil.SplitFirstLast("a.b.c", ".")
	actual := args.Map{"first": first, "last": last}
	expected := args.Map{"first": "a", "last": "c"}
	expected.ShouldBeEqual(t, 0, "SplitFirstLast returns correct value -- with args", actual)
}
