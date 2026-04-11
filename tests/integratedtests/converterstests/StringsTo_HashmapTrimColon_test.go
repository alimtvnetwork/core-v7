package converterstests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── stringsTo coverage ──

func Test_StringsTo_HashmapTrimColon_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapTrimColon("a:1", "b:2")

	// Act
	actual := args.Map{
		"len":    len(result),
		"hasA":   result["a"] == "1",
	}

	// Assert
	expected := args.Map{
		"len":    2,
		"hasA":   true,
	}
	expected.ShouldBeEqual(t, 0, "HashmapTrimColon returns correct value -- with args", actual)
}

func Test_StringsTo_HashmapTrimHyphen_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapTrimHyphen("a-1", "b-2")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "HashmapTrimHyphen returns correct value -- with args", actual)
}

func Test_StringsTo_HashmapOptions_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapOptions(true, "=", "a = 1", "b = 2")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "HashmapOptions returns correct value -- with args", actual)
}

func Test_StringsTo_HashmapTrim_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapTrim(":", []string{"a:1"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapTrim returns correct value -- with args", actual)
}

func Test_StringsTo_HashmapUsingFuncOptions(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapUsingFuncOptions(true, func(line string) (string, string) {
		return "k", "v"
	}, "line1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapUsingFuncOptions returns correct value -- with args", actual)
}

func Test_StringsTo_HashmapUsingFuncTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.HashmapUsingFuncTrim(func(line string) (string, string) {
		return "k", "v"
	}, "line1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "HashmapUsingFuncTrim returns correct value -- with args", actual)
}

func Test_StringsTo_MapStringIntegerUsingFunc(t *testing.T) {
	// Arrange
	result := converters.StringsTo.MapStringIntegerUsingFunc(true, func(line string) (string, int) {
		return "k", 42
	}, "line1")

	// Act
	actual := args.Map{
		"len": len(result),
		"val": result["k"],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "MapStringIntegerUsingFunc returns correct value -- with args", actual)
}

func Test_StringsTo_MapStringAnyUsingFunc(t *testing.T) {
	// Arrange
	result := converters.StringsTo.MapStringAnyUsingFunc(false, func(line string) (string, any) {
		return "k", "val"
	}, "line1")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "MapStringAnyUsingFunc returns correct value -- with args", actual)
}

func Test_StringsTo_MapConverter_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	mc := converters.StringsTo.MapConverter("a:1", "b:2")

	// Act
	actual := args.Map{
		"len": mc.Length(),
		"hasAny": mc.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "MapConverter returns correct value -- with args", actual)
}

func Test_StringsTo_IntegersSkipMapAndDefaultValue(t *testing.T) {
	// Arrange
	skip := map[string]bool{"skip": true}
	result := converters.StringsTo.IntegersSkipMapAndDefaultValue(-1, skip, "1", "skip", "abc")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
		"second": result[1],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
		"second": 0,
	}
	expected.ShouldBeEqual(t, 0, "IntegersSkipMapAndDefaultValue returns correct value -- with args", actual)
}

func Test_StringsTo_IntegersSkipAndDefaultValue_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.IntegersSkipAndDefaultValue(-1, "skip", "1", "skip", "abc")

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 3,
		"first": 1,
	}
	expected.ShouldBeEqual(t, 0, "IntegersSkipAndDefaultValue returns correct value -- with args", actual)
}

func Test_StringsTo_BytesConditional(t *testing.T) {
	// Arrange
	result := converters.StringsTo.BytesConditional(func(in string) (byte, bool, bool) {
		return in[0], true, false
	}, []string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "BytesConditional returns correct value -- with args", actual)
}

func Test_StringsTo_BytesWithDefaults(t *testing.T) {
	// Arrange
	result := converters.StringsTo.BytesWithDefaults(0, "1", "abc", "300")

	// Act
	actual := args.Map{
		"len":      result.Length(),
		"hasError": result.HasError(),
		"first":    int(result.Values[0]),
	}

	// Assert
	expected := args.Map{
		"len":      3,
		"hasError": true,
		"first":    1,
	}
	expected.ShouldBeEqual(t, 0, "BytesWithDefaults returns non-empty -- with args", actual)
}

func Test_StringsTo_Csv_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Csv(false, "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Csv returns correct value -- with args", actual)
}

func Test_StringsTo_CsvUsingPtrStrings(t *testing.T) {
	// Arrange
	nilResult := converters.StringsTo.CsvUsingPtrStrings(false, nil)
	items := []string{"a", "b"}
	result := converters.StringsTo.CsvUsingPtrStrings(false, &items)

	// Act
	actual := args.Map{
		"nilEmpty": nilResult == "",
		"notEmpty": result != "",
	}

	// Assert
	expected := args.Map{
		"nilEmpty": true,
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "CsvUsingPtrStrings returns correct value -- with args", actual)
}

func Test_StringsTo_CsvWithIndexes(t *testing.T) {
	// Arrange
	result := converters.StringsTo.CsvWithIndexes([]string{"a", "b"})

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CsvWithIndexes returns non-empty -- with args", actual)
}

func Test_StringsTo_Float64sMust(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Float64sMust("1.5", "2.5")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Float64sMust returns correct value -- with args", actual)
}

func Test_StringsTo_Float64sConditional(t *testing.T) {
	// Arrange
	result := converters.StringsTo.Float64sConditional(func(in string) (float64, bool, bool) {
		return 1.0, true, false
	}, []string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "Float64sConditional returns correct value -- with args", actual)
}

func Test_StringsTo_PtrOfPtrToPtrStrings(t *testing.T) {
	// Arrange
	nilResult := converters.StringsTo.PtrOfPtrToPtrStrings(nil)

	// Act
	actual := args.Map{"nilEmpty": len(*nilResult) == 0}

	// Assert
	expected := args.Map{"nilEmpty": true}
	expected.ShouldBeEqual(t, 0, "PtrOfPtrToPtrStrings_nil returns nil -- with args", actual)

	s1 := "a"
	s2 := "b"
	var nilStr *string
	items := []*string{&s1, &s2, nilStr}
	result := converters.StringsTo.PtrOfPtrToPtrStrings(&items)
	actual2 := args.Map{
		"len": len(*result),
		"third": (*result)[2],
	}
	expected2 := args.Map{
		"len": 3,
		"third": "",
	}
	expected2.ShouldBeEqual(t, 1, "PtrOfPtrToPtrStrings_values returns non-empty -- with args", actual2)
}

func Test_StringsTo_PtrOfPtrToMapStringBool(t *testing.T) {
	// Arrange
	nilResult := converters.StringsTo.PtrOfPtrToMapStringBool(nil)

	// Act
	actual := args.Map{"nilLen": len(nilResult)}

	// Assert
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "PtrOfPtrToMapStringBool_nil returns nil -- with args", actual)

	s1 := "a"
	var nilStr *string
	items := []*string{&s1, nilStr}
	result := converters.StringsTo.PtrOfPtrToMapStringBool(&items)
	actual2 := args.Map{
		"len": len(result),
		"hasA": result["a"],
	}
	expected2 := args.Map{
		"len": 1,
		"hasA": true,
	}
	expected2.ShouldBeEqual(t, 1, "PtrOfPtrToMapStringBool_values returns non-empty -- with args", actual2)
}

func Test_StringsTo_CloneIf(t *testing.T) {
	// Arrange
	original := []string{"a", "b"}
	cloned := converters.StringsTo.CloneIf(true, original...)
	notCloned := converters.StringsTo.CloneIf(false, original...)
	emptyClone := converters.StringsTo.CloneIf(true)

	// Act
	actual := args.Map{
		"clonedLen": len(cloned),
		"notClonedLen": len(notCloned),
		"emptyLen": len(emptyClone),
	}

	// Assert
	expected := args.Map{
		"clonedLen": 2,
		"notClonedLen": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "CloneIf returns correct value -- with args", actual)
}

// ── StringsToMapConverter coverage ──

func Test_StringsToMapConverter_SafeStrings(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{"a:1"})

	// Act
	actual := args.Map{"len": len(mc.SafeStrings())}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "SafeStrings returns correct value -- with args", actual)

	var emptyMc converters.StringsToMapConverter
	actual2 := args.Map{"len": len(emptyMc.SafeStrings())}
	expected2 := args.Map{"len": 0}
	expected2.ShouldBeEqual(t, 1, "SafeStrings_empty returns empty -- with args", actual2)
}

func Test_StringsToMapConverter_LineSplitMapOptions(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{"a:1"})
	trimResult := mc.LineSplitMapOptions(true, ":")
	noTrimResult := mc.LineSplitMapOptions(false, ":")

	// Act
	actual := args.Map{
		"trimLen":   len(trimResult),
		"noTrimLen": len(noTrimResult),
	}

	// Assert
	expected := args.Map{
		"trimLen":   1,
		"noTrimLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "LineSplitMapOptions returns correct value -- with args", actual)
}

func Test_StringsToMapConverter_LineProcessorMapOptions(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapOptions(true, func(line string) (string, string) {
		return "k", "v"
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapOptions returns correct value -- with args", actual)
}

func Test_StringsToMapConverter_LineProcessorMapStringIntegerTrim_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapStringIntegerTrim(func(line string) (string, int) {
		return "k", 42
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"val": result["k"],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"val": 42,
	}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringIntegerTrim returns correct value -- with args", actual)
}

func Test_StringsToMapConverter_LineProcessorMapStringAnyTrim_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	mc := converters.StringsToMapConverter([]string{"line1"})
	result := mc.LineProcessorMapStringAnyTrim(func(line string) (string, any) {
		return "k", "val"
	})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringAnyTrim returns correct value -- with args", actual)
}

func Test_StringsToMapConverter_NilReceiver(t *testing.T) {
	// Arrange
	var mc *converters.StringsToMapConverter

	// Act
	actual := args.Map{
		"length":    mc.Length(),
		"isEmpty":   mc.IsEmpty(),
		"hasAny":    mc.HasAnyItem(),
		"lastIndex": mc.LastIndex(),
	}

	// Assert
	expected := args.Map{
		"length":    0,
		"isEmpty":   true,
		"hasAny":    false,
		"lastIndex": -1,
	}
	expected.ShouldBeEqual(t, 0, "NilReceiver returns nil -- with args", actual)
}

// ── anyItemConverter extra coverage ──

func Test_AnyTo_ToStringsUsingProcessor(t *testing.T) {
	// Arrange
	nilResult := converters.AnyTo.ToStringsUsingProcessor(true, func(index int, in any) (string, bool, bool) {
		return "x", true, false
	}, nil)

	// Act
	actual := args.Map{"nilLen": len(nilResult)}

	// Assert
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingProcessor_nil returns nil -- with args", actual)
}

func Test_AnyTo_ToStringsUsingSimpleProcessor(t *testing.T) {
	// Arrange
	nilResult := converters.AnyTo.ToStringsUsingSimpleProcessor(true, func(index int, in any) string {
		return "x"
	}, nil)

	// Act
	actual := args.Map{"nilLen": len(nilResult)}

	// Assert
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsUsingSimpleProcessor_nil returns nil -- with args", actual)
}

func Test_AnyTo_ToNonNullItems(t *testing.T) {
	// Arrange
	nilResult := converters.AnyTo.ToNonNullItems(true, nil)

	// Act
	actual := args.Map{"nilLen": len(nilResult)}

	// Assert
	expected := args.Map{"nilLen": 0}
	expected.ShouldBeEqual(t, 0, "ToNonNullItems_nil returns nil -- with args", actual)
}

func Test_AnyTo_SmartStringsJoiner_FromStringsToHashmapTrim(t *testing.T) {
	// Arrange
	result := converters.AnyTo.SmartStringsJoiner(",", "a", "b")
	emptyResult := converters.AnyTo.SmartStringsJoiner(",")

	// Act
	actual := args.Map{
		"notEmpty": result != "",
		"empty": emptyResult,
	}

	// Assert
	expected := args.Map{
		"notEmpty": true,
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "SmartStringsJoiner returns correct value -- with args", actual)
}

// ── StringTo IntegerMust ──

func Test_StringTo_IntegerMust(t *testing.T) {
	// Arrange
	result := converters.StringTo.IntegerMust("42")

	// Act
	actual := args.Map{"val": result}

	// Assert
	expected := args.Map{"val": 42}
	expected.ShouldBeEqual(t, 0, "IntegerMust returns correct value -- with args", actual)
}

func Test_StringTo_Float64Must(t *testing.T) {
	// Arrange
	result := converters.StringTo.Float64Must("3.14")

	// Act
	actual := args.Map{"gt3": result > 3.0}

	// Assert
	expected := args.Map{"gt3": true}
	expected.ShouldBeEqual(t, 0, "Float64Must returns correct value -- with args", actual)
}
