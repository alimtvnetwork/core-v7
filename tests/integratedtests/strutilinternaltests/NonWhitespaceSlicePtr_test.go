package strutilinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

// ── NonWhitespaceSlicePtr ──

func Test_NonWhitespaceSlicePtr(t *testing.T) {
	// Arrange
	result := strutilinternal.NonWhitespaceSlicePtr([]string{"a", " ", "b"})
	emptyResult := strutilinternal.NonWhitespaceSlicePtr([]string{})

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceSlicePtr returns correct value -- with args", actual)
}

// ── NonWhitespaceTrimSlicePtr ──

func Test_NonWhitespaceTrimSlicePtr(t *testing.T) {
	// Arrange
	result := strutilinternal.NonWhitespaceTrimSlicePtr([]string{" a ", " "})
	emptyResult := strutilinternal.NonWhitespaceTrimSlicePtr([]string{})

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(emptyResult),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"emptyLen": 0,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceTrimSlicePtr returns correct value -- with args", actual)
}

// ── NonWhitespaceJoinPtr ──

func Test_NonWhitespaceJoinPtr(t *testing.T) {
	// Arrange
	result := strutilinternal.NonWhitespaceJoinPtr([]string{"a", " ", "b"}, ",")
	emptyResult := strutilinternal.NonWhitespaceJoinPtr([]string{}, ",")

	// Act
	actual := args.Map{
		"result": result,
		"empty": emptyResult,
	}

	// Assert
	expected := args.Map{
		"result": "a,b",
		"empty": "",
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoinPtr returns correct value -- with args", actual)
}

// ── SliceToMapConverter extended ──

func Test_SliceToMapConverter_Hashset_FromNonWhitespaceSlicePt(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"a", "b", "a"}
	hs := conv.Hashset()

	// Act
	actual := args.Map{
		"len": len(hs),
		"hasA": hs["a"],
		"hasB": hs["b"],
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"hasA": true,
		"hasB": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter returns correct value -- Hashset", actual)
}

func Test_SliceToMapConverter_LineSplitMap_FromNonWhitespaceSlicePt(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"key=value", "a=b"}
	result := conv.LineSplitMap("=")

	// Act
	actual := args.Map{
		"key": result["key"],
		"a": result["a"],
	}

	// Assert
	expected := args.Map{
		"key": "value",
		"a": "b",
	}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter returns correct value -- LineSplitMap", actual)
}

func Test_SliceToMapConverter_LineSplitMapTrim(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"  key  =  value  ", "  "}
	result := conv.LineSplitMapTrim("=")

	// Act
	actual := args.Map{"hasKey": result["key"] != ""}

	// Assert
	expected := args.Map{"hasKey": true}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter returns correct value -- LineSplitMapTrim", actual)
}

func Test_SliceToMapConverter_LineSplitMapOptions(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"key=val"}
	trimResult := conv.LineSplitMapOptions(true, "=")
	noTrimResult := conv.LineSplitMapOptions(false, "=")

	// Act
	actual := args.Map{
		"trimLen":   len(trimResult),
		"noTrimLen": len(noTrimResult),
	}

	// Assert
	expected := args.Map{
		"trimLen": 1,
		"noTrimLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter returns correct value -- LineSplitMapOptions", actual)
}

func Test_SliceToMapConverter_LineSplitMap_Empty(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{}
	result := conv.LineSplitMap("=")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter returns empty -- LineSplitMap empty", actual)
}

func Test_SliceToMapConverter_LineProcessorMapOptions(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello", "world", "  "}
	fn := func(line string) (string, string) { return line, "v" }
	trimResult := conv.LineProcessorMapOptions(true, fn)
	noTrimResult := conv.LineProcessorMapOptions(false, fn)
	nilResult := conv.LineProcessorMapOptions(true, nil)

	// Act
	actual := args.Map{
		"trimLen":   len(trimResult),
		"noTrimLen": len(noTrimResult),
		"nilLen":    len(nilResult),
	}

	// Assert
	expected := args.Map{
		"trimLen": 2,
		"noTrimLen": 3,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "SliceToMapConverter returns correct value -- LineProcessorMapOptions", actual)
}

func Test_SliceToMapConverter_LineProcessorMapStringIntegerTrim(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello", "  "}
	fn := func(line string) (string, int) { return line, 1 }
	result := conv.LineProcessorMapStringIntegerTrim(fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringIntegerTrim returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineProcessorMapStringIntegerOptions(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello"}
	fn := func(line string) (string, int) { return line, 1 }
	trimResult := conv.LineProcessorMapStringIntegerOptions(true, fn)
	noTrimResult := conv.LineProcessorMapStringIntegerOptions(false, fn)
	nilResult := conv.LineProcessorMapStringIntegerOptions(true, nil)

	// Act
	actual := args.Map{
		"trimLen":   len(trimResult),
		"noTrimLen": len(noTrimResult),
		"nilLen":    len(nilResult),
	}

	// Assert
	expected := args.Map{
		"trimLen": 1,
		"noTrimLen": 1,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringIntegerOptions returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineProcessorMapStringAnyTrim(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello", "  "}
	fn := func(line string) (string, any) { return line, 1 }
	result := conv.LineProcessorMapStringAnyTrim(fn)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringAnyTrim returns correct value -- with args", actual)
}

func Test_SliceToMapConverter_LineProcessorMapStringAnyOptions(t *testing.T) {
	// Arrange
	conv := strutilinternal.SliceToMapConverter{"hello"}
	fn := func(line string) (string, any) { return line, "v" }
	trimResult := conv.LineProcessorMapStringAnyOptions(true, fn)
	noTrimResult := conv.LineProcessorMapStringAnyOptions(false, fn)
	nilResult := conv.LineProcessorMapStringAnyOptions(true, nil)

	// Act
	actual := args.Map{
		"trimLen":   len(trimResult),
		"noTrimLen": len(noTrimResult),
		"nilLen":    len(nilResult),
	}

	// Assert
	expected := args.Map{
		"trimLen": 1,
		"noTrimLen": 1,
		"nilLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "LineProcessorMapStringAnyOptions returns correct value -- with args", actual)
}

// ── ReplaceTemplateMap non-curly ──

func Test_ReplaceTemplateMap_NonCurly(t *testing.T) {
	// Arrange
	result := strutilinternal.ReplaceTemplateMap(
		false,
		"Hello name, you are age",
		map[string]string{"name": "Alice", "age": "30"},
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "Hello Alice, you are 30"}
	expected.ShouldBeEqual(t, 0, "ReplaceTemplateMap returns non-empty -- non-curly", actual)
}
