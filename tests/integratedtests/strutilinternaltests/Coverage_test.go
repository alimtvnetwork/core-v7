package strutilinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

func Test_IsNullOrEmpty(t *testing.T) {
	// Arrange
	empty := ""
	hello := "hello"

	// Act
	actual := args.Map{
		"nil":      strutilinternal.IsNullOrEmpty(nil),
		"empty":    strutilinternal.IsNullOrEmpty(&empty),
		"nonEmpty": strutilinternal.IsNullOrEmpty(&hello),
	}
	expected := args.Map{
		"nil":      true,
		"empty":    true,
		"nonEmpty": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmpty returns empty -- with args", actual)
}

func Test_IsNullOrEmptyOrWhitespace(t *testing.T) {
	// Arrange
	space := " "
	hello := "hello"

	// Act
	actual := args.Map{
		"nil":   strutilinternal.IsNullOrEmptyOrWhitespace(nil),
		"space": strutilinternal.IsNullOrEmptyOrWhitespace(&space),
		"text":  strutilinternal.IsNullOrEmptyOrWhitespace(&hello),
	}
	expected := args.Map{
		"nil":   true,
		"space": true,
		"text":  false,
	}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_IsEmptyOrWhitespace(t *testing.T) {
	// Act
	actual := args.Map{
		"empty":     strutilinternal.IsEmptyOrWhitespace(""),
		"space":     strutilinternal.IsEmptyOrWhitespace(" "),
		"newline":   strutilinternal.IsEmptyOrWhitespace("\n"),
		"tabs":      strutilinternal.IsEmptyOrWhitespace("\t"),
		"text":      strutilinternal.IsEmptyOrWhitespace("hello"),
	}
	expected := args.Map{
		"empty":     true,
		"space":     true,
		"newline":   true,
		"tabs":      true,
		"text":      false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_NonEmptySlice(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.NonEmptySlice([]string{"a", "", "b", "", "c"})
	emptyResult := strutilinternal.NonEmptySlice([]string{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(emptyResult),
	}
	expected := args.Map{
		"resultLen": 3,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlice returns empty -- with args", actual)
}

func Test_NonEmptySlicePtr(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.NonEmptySlicePtr([]string{"a", "", "b"})
	emptyResult := strutilinternal.NonEmptySlicePtr([]string{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(emptyResult),
	}
	expected := args.Map{
		"resultLen": 2,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr returns empty -- with args", actual)
}

func Test_NonEmptyJoin(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.NonEmptyJoin([]string{"a", "", "b"}, ",")
	emptyResult := strutilinternal.NonEmptyJoin([]string{}, ",")

	// Assert
	actual := args.Map{
		"result": result,
		"empty":  emptyResult,
	}
	expected := args.Map{
		"result": "a,b",
		"empty":  "",
	}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- with args", actual)
}

func Test_NonWhitespaceSlice(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.NonWhitespaceSlice([]string{"a", " ", "b", "\n", "c"})
	nilResult := strutilinternal.NonWhitespaceSlice(nil)
	emptyResult := strutilinternal.NonWhitespaceSlice([]string{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"nilLen":    len(nilResult),
		"emptyLen":  len(emptyResult),
	}
	expected := args.Map{
		"resultLen": 3,
		"nilLen":    0,
		"emptyLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceSlice returns correct value -- with args", actual)
}

func Test_NonWhitespaceTrimSlice(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.NonWhitespaceTrimSlice([]string{" a ", " ", " b "})
	nilResult := strutilinternal.NonWhitespaceTrimSlice(nil)
	emptyResult := strutilinternal.NonWhitespaceTrimSlice([]string{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"nilLen":    len(nilResult),
		"emptyLen":  len(emptyResult),
		"firstItem": result[0],
	}
	expected := args.Map{
		"resultLen": 2,
		"nilLen":    0,
		"emptyLen":  0,
		"firstItem": "a",
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceTrimSlice returns correct value -- with args", actual)
}

func Test_NonWhitespaceJoin(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.NonWhitespaceJoin([]string{"a", " ", "b"}, ",")
	nilResult := strutilinternal.NonWhitespaceJoin(nil, ",")
	emptyResult := strutilinternal.NonWhitespaceJoin([]string{}, ",")

	// Assert
	actual := args.Map{
		"result": result,
		"nil":    nilResult,
		"empty":  emptyResult,
	}
	expected := args.Map{
		"result": "a,b",
		"nil":    "",
		"empty":  "",
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns correct value -- with args", actual)
}

func Test_Clone(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.Clone([]string{"a", "b", "c"})
	emptyResult := strutilinternal.Clone([]string{})

	// Assert
	actual := args.Map{
		"resultLen": len(result),
		"emptyLen":  len(emptyResult),
		"firstItem": result[0],
	}
	expected := args.Map{
		"resultLen": 3,
		"emptyLen":  0,
		"firstItem": "a",
	}
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- with args", actual)
}

func Test_CurlyWrapIf(t *testing.T) {
	// Arrange & Act
	curly := strutilinternal.CurlyWrapIf(true, "hello")
	noCurly := strutilinternal.CurlyWrapIf(false, "hello")

	// Assert
	actual := args.Map{
		"curlyNotEmpty":   curly != "",
		"noCurlyNotEmpty": noCurly != "",
	}
	expected := args.Map{
		"curlyNotEmpty":   true,
		"noCurlyNotEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf returns correct value -- with args", actual)
}

func Test_AnyToString(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.AnyToString("hello")
	nilResult := strutilinternal.AnyToString(nil)
	intResult := strutilinternal.AnyToString(42)

	// Assert
	actual := args.Map{
		"result":    result,
		"nilResult": nilResult,
		"intResult": intResult,
	}
	expected := args.Map{
		"result":    "hello",
		"nilResult": "",
		"intResult": "42",
	}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- with args", actual)
}

func Test_AnyToFieldNameString(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.AnyToFieldNameString("hello")
	nilResult := strutilinternal.AnyToFieldNameString(nil)

	// Assert
	actual := args.Map{
		"resultNotEmpty": result != "",
		"nilResult":      nilResult,
	}
	expected := args.Map{
		"resultNotEmpty": true,
		"nilResult":      "",
	}
	expected.ShouldBeEqual(t, 0, "AnyToFieldNameString returns correct value -- with args", actual)
}

func Test_AnyToStringUsing(t *testing.T) {
	// Arrange & Act
	withFields := strutilinternal.AnyToStringUsing(true, "hello")
	withoutFields := strutilinternal.AnyToStringUsing(false, "hello")
	nilResult := strutilinternal.AnyToStringUsing(true, nil)

	// Assert
	actual := args.Map{
		"withFieldsNotEmpty":   withFields != "",
		"withoutFieldsNotEmpty": withoutFields != "",
		"nilResult":            nilResult,
	}
	expected := args.Map{
		"withFieldsNotEmpty":   true,
		"withoutFieldsNotEmpty": true,
		"nilResult":            "",
	}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing returns correct value -- with args", actual)
}

func Test_MaskLine(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.MaskLine("XXXXXXXXXX", "abc")
	emptyLine := strutilinternal.MaskLine("XXXXXXXXXX", "")
	longLine := strutilinternal.MaskLine("XXX", "abcdef")
	emptyMask := strutilinternal.MaskLine("", "abc")

	// Assert
	actual := args.Map{
		"result":    result,
		"emptyLine": emptyLine,
		"longLine":  longLine,
		"emptyMask": emptyMask,
	}
	expected := args.Map{
		"result":    "abcXXXXXXX",
		"emptyLine": "XXXXXXXXXX",
		"longLine":  "abcdef",
		"emptyMask": "abc",
	}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- with args", actual)
}

func Test_MaskTrimLine(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.MaskTrimLine("XXXXXXXXXX", "  abc  ")
	emptyLine := strutilinternal.MaskTrimLine("XXXXXXXXXX", "   ")

	// Assert
	actual := args.Map{
		"result":    result,
		"emptyLine": emptyLine,
	}
	expected := args.Map{
		"result":    "abcXXXXXXX",
		"emptyLine": "XXXXXXXXXX",
	}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- with args", actual)
}

func Test_SplitLeftRight(t *testing.T) {
	// Arrange & Act
	l, r := strutilinternal.SplitLeftRight("=", "key=value")
	l2, r2 := strutilinternal.SplitLeftRight("=", "noequals")
	l3, r3 := strutilinternal.SplitLeftRightTrim("=", "  key  =  value  ")

	// Assert
	actual := args.Map{
		"l":  l,
		"r":  r,
		"l2": l2,
		"r2": r2,
		"l3": l3,
		"r3": r3,
	}
	expected := args.Map{
		"l":  "key",
		"r":  "value",
		"l2": "noequals",
		"r2": "",
		"l3": "key",
		"r3": "value",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- with args", actual)
}

func Test_ReplaceTemplateMap(t *testing.T) {
	// Arrange & Act
	result := strutilinternal.ReplaceTemplateMap(
		true,
		"Hello {name}, you are {age}",
		map[string]string{"name": "Alice", "age": "30"},
	)
	emptyTemplate := strutilinternal.ReplaceTemplateMap(true, "", map[string]string{"a": "b"})
	emptyMap := strutilinternal.ReplaceTemplateMap(true, "hello", map[string]string{})

	// Assert
	actual := args.Map{
		"result":        result,
		"emptyTemplate": emptyTemplate,
		"emptyMap":      emptyMap,
	}
	expected := args.Map{
		"result":        "Hello Alice, you are 30",
		"emptyTemplate": "",
		"emptyMap":      "hello",
	}
	expected.ShouldBeEqual(t, 0, "ReplaceTemplateMap returns correct value -- with args", actual)
}

func Test_ReflectInterfaceVal(t *testing.T) {
	// Arrange
	val := 42
	ptrVal := &val

	// Act
	result := strutilinternal.ReflectInterfaceVal(val)
	ptrResult := strutilinternal.ReflectInterfaceVal(ptrVal)

	// Assert
	actual := args.Map{
		"result":    result,
		"ptrResult": ptrResult,
	}
	expected := args.Map{
		"result":    42,
		"ptrResult": 42,
	}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_Basic(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"a", "b", "c"}
	var nilConv *strutilinternal.SliceToMapConverter

	// Act
	actual := args.Map{
		"length":     conv.Length(),
		"isEmpty":    conv.IsEmpty(),
		"hasAny":     conv.HasAnyItem(),
		"lastIndex":  conv.LastIndex(),
		"nilLength":  nilConv.Length(),
		"nilIsEmpty": nilConv.IsEmpty(),
		"stringsLen": len(conv.Strings()),
		"safeLen":    len(conv.SafeStrings()),
	}
	expected := args.Map{
		"length":     3,
		"isEmpty":    false,
		"hasAny":     true,
		"lastIndex":  2,
		"nilLength":  0,
		"nilIsEmpty": true,
		"stringsLen": 3,
		"safeLen":    3,
	}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter_Basic returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_Hashset(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"a", "b", "c"}

	// Act
	hashset := conv.Hashset()

	// Assert
	actual := args.Map{
		"hasA": hashset["a"],
		"hasB": hashset["b"],
		"hasD": hashset["d"],
	}
	expected := args.Map{
		"hasA": true,
		"hasB": true,
		"hasD": false,
	}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter_Hashset returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineSplitMap(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"key1=val1", "key2=val2"}

	// Act
	splitMap := conv.LineSplitMap("=")
	splitTrimMap := conv.LineSplitMapTrim("=")
	splitOpts := conv.LineSplitMapOptions(false, "=")
	splitOptsT := conv.LineSplitMapOptions(true, "=")
	emptyConv := strutilinternal.SliceToMapConverter{}
	emptyMap := emptyConv.LineSplitMap("=")

	// Assert
	actual := args.Map{
		"mapLen":     len(splitMap),
		"trimMapLen": len(splitTrimMap),
		"optsLen":    len(splitOpts),
		"optsTLen":   len(splitOptsT),
		"emptyLen":   len(emptyMap),
		"val1":       splitMap["key1"],
	}
	expected := args.Map{
		"mapLen":     2,
		"trimMapLen": 2,
		"optsLen":    2,
		"optsTLen":   2,
		"emptyLen":   0,
		"val1":       "val1",
	}
	expected.ShouldBeEqual(t, 0, "LineSplitMap returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineProcessorMap(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello", "world"}
	processor := func(line string) (string, string) { return line, line + "!" }

	// Act
	noTrim := conv.LineProcessorMapOptions(false, processor)
	withTrim := conv.LineProcessorMapOptions(true, processor)
	nilProc := conv.LineProcessorMapOptions(false, nil)
	emptyConv := strutilinternal.SliceToMapConverter{}
	emptyResult := emptyConv.LineProcessorMapOptions(false, processor)

	// Assert
	actual := args.Map{
		"noTrimLen":   len(noTrim),
		"withTrimLen": len(withTrim),
		"nilProcLen":  len(nilProc),
		"emptyLen":    len(emptyResult),
	}
	expected := args.Map{
		"noTrimLen":   2,
		"withTrimLen": 2,
		"nilProcLen":  0,
		"emptyLen":    0,
	}
	expected.ShouldBeEqual(t, 0, "LineProcessorMap returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineProcessorMapStringInteger(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello", "world"}
	processor := func(line string) (string, int) { return line, len(line) }

	// Act
	noTrim := conv.LineProcessorMapStringIntegerOptions(false, processor)
	withTrim := conv.LineProcessorMapStringIntegerTrim(processor)
	nilProc := conv.LineProcessorMapStringIntegerOptions(false, nil)

	// Assert
	actual := args.Map{
		"noTrimLen":   len(noTrim),
		"withTrimLen": len(withTrim),
		"nilProcLen":  len(nilProc),
	}
	expected := args.Map{
		"noTrimLen":   2,
		"withTrimLen": 2,
		"nilProcLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringInteger returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineProcessorMapStringAny(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello", "world"}
	processor := func(line string) (string, any) { return line, len(line) }

	// Act
	noTrim := conv.LineProcessorMapStringAnyOptions(false, processor)
	withTrim := conv.LineProcessorMapStringAnyTrim(processor)
	nilProc := conv.LineProcessorMapStringAnyOptions(false, nil)

	// Assert
	actual := args.Map{
		"noTrimLen":   len(noTrim),
		"withTrimLen": len(withTrim),
		"nilProcLen":  len(nilProc),
	}
	expected := args.Map{
		"noTrimLen":   2,
		"withTrimLen": 2,
		"nilProcLen":  0,
	}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringAny returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_EmptySafeStrings(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{}

	// Act
	actual := args.Map{
		"safeLen": len(conv.SafeStrings()),
	}
	expected := args.Map{
		"safeLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "EmptySafeStrings returns empty -- with args", actual)
}
