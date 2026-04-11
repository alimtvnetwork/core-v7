package strutilinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

// ── AnyToFieldNameString ──

func Test_AnyToFieldNameString_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.AnyToFieldNameString(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToFieldNameString returns nil -- nil", actual)
}

func Test_AnyToFieldNameString_Value(t *testing.T) {
	// Arrange
	result := strutilinternal.AnyToFieldNameString("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToFieldNameString returns correct value -- value", actual)
}

// ── AnyToString ──

func Test_AnyToString_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.AnyToString(nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToString returns nil -- nil", actual)
}

func Test_AnyToString_Value(t *testing.T) {
	// Arrange
	result := strutilinternal.AnyToString("hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- value", actual)
}

func Test_AnyToString_Ptr(t *testing.T) {
	// Arrange
	v := "hello"
	result := strutilinternal.AnyToString(&v)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToString returns correct value -- ptr", actual)
}

// ── AnyToStringUsing ──

func Test_AnyToStringUsing_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.AnyToStringUsing(true, nil)}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing returns nil -- nil", actual)
}

func Test_AnyToStringUsing_IncludeFields(t *testing.T) {
	// Arrange
	result := strutilinternal.AnyToStringUsing(true, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing returns correct value -- include fields", actual)
}

func Test_AnyToStringUsing_NoFields(t *testing.T) {
	// Arrange
	result := strutilinternal.AnyToStringUsing(false, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyToStringUsing returns empty -- no fields", actual)
}

// ── MaskLine ──

func Test_MaskLine_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.MaskLine("****", "")}

	// Assert
	expected := args.Map{"result": "****"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns empty -- empty line", actual)
}

func Test_MaskLine_LongerThanMask(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.MaskLine("**", "hello")}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- longer than mask", actual)
}

func Test_MaskLine_EmptyMask(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.MaskLine("", "hello")}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskLine returns empty -- empty mask", actual)
}

func Test_MaskLine_Partial(t *testing.T) {
	// Arrange
	result := strutilinternal.MaskLine("********", "hi")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 8}
	expected.ShouldBeEqual(t, 0, "MaskLine returns correct value -- partial", actual)
}

// ── MaskTrimLine ──

func Test_MaskTrimLine_Whitespace(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.MaskTrimLine("****", "   ")}

	// Assert
	expected := args.Map{"result": "****"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- whitespace", actual)
}

func Test_MaskTrimLine_Longer(t *testing.T) {
	// Act
	actual := args.Map{"result": strutilinternal.MaskTrimLine("**", " hello ")}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- longer than mask", actual)
}

func Test_MaskTrimLine_Partial(t *testing.T) {
	// Arrange
	result := strutilinternal.MaskTrimLine("********", " hi ")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 8}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine returns correct value -- partial", actual)
}

// ── SplitLeftRight / SplitLeftRightTrim ──

func Test_SplitLeftRight_TwoParts(t *testing.T) {
	// Arrange
	l, r := strutilinternal.SplitLeftRight("=", "key=value")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "value",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns correct value -- two parts", actual)
}

func Test_SplitLeftRight_NoParts(t *testing.T) {
	// Arrange
	l, r := strutilinternal.SplitLeftRight("=", "noequals")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": "noequals",
		"right": "",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRight returns empty -- no parts", actual)
}

func Test_SplitLeftRightTrim(t *testing.T) {
	// Arrange
	l, r := strutilinternal.SplitLeftRightTrim("=", " key = value ")

	// Act
	actual := args.Map{
		"left": l,
		"right": r,
	}

	// Assert
	expected := args.Map{
		"left": "key",
		"right": "value",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTrim returns correct value -- with args", actual)
}

// ── CurlyWrapIf ──

func Test_CurlyWrapIf_True(t *testing.T) {
	// Arrange
	result := strutilinternal.CurlyWrapIf(true, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf returns non-empty -- true", actual)
}

func Test_CurlyWrapIf_False(t *testing.T) {
	// Arrange
	result := strutilinternal.CurlyWrapIf(false, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CurlyWrapIf returns non-empty -- false", actual)
}

// ── Clone ──

func Test_Clone_Empty(t *testing.T) {
	// Arrange
	result := strutilinternal.Clone([]string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Clone returns empty -- empty", actual)
}

func Test_Clone_Items(t *testing.T) {
	// Arrange
	result := strutilinternal.Clone([]string{"a", "b"})

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
	expected.ShouldBeEqual(t, 0, "Clone returns correct value -- items", actual)
}

// ── ReflectInterfaceVal ──

func Test_ReflectInterfaceVal_NonPtr(t *testing.T) {
	// Arrange
	result := strutilinternal.ReflectInterfaceVal("hello")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns non-empty -- non-ptr", actual)
}

func Test_ReflectInterfaceVal_Ptr(t *testing.T) {
	// Arrange
	v := "hello"
	result := strutilinternal.ReflectInterfaceVal(&v)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "ReflectInterfaceVal returns correct value -- ptr", actual)
}

// ── NonEmpty / NonEmptyJoin ──

func Test_NonEmpty(t *testing.T) {
	// Arrange
	result := strutilinternal.NonEmptySlice([]string{"a", "", "b", ""})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmpty returns empty -- with args", actual)
}

func Test_NonEmptyJoin(t *testing.T) {
	// Arrange
	result := strutilinternal.NonEmptyJoin([]string{"a", "", "b"}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonEmptyJoin returns empty -- with args", actual)
}

// ── NonEmptySlicePtr ──

func Test_NonEmptySlicePtr(t *testing.T) {
	// Arrange
	result := strutilinternal.NonEmptySlicePtr([]string{"a", "", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonEmptySlicePtr returns empty -- with args", actual)
}

// ── NonWhitespaceSlice ──

func Test_NonWhitespaceSlice(t *testing.T) {
	// Arrange
	result := strutilinternal.NonWhitespaceSlice([]string{"a", " ", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceSlice returns correct value -- with args", actual)
}

// ── NonWhitespaceTrimSlice ──

func Test_NonWhitespaceTrimSlice(t *testing.T) {
	// Arrange
	result := strutilinternal.NonWhitespaceTrimSlice([]string{" a ", " "})

	// Act
	actual := args.Map{
		"len": len(result),
		"first": result[0],
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"first": "a",
	}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceTrimSlice returns correct value -- with args", actual)
}

// ── NonWhitespaceJoin ──

func Test_NonWhitespaceJoin(t *testing.T) {
	// Arrange
	result := strutilinternal.NonWhitespaceJoin([]string{"a", " ", "b"}, ",")

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "a,b"}
	expected.ShouldBeEqual(t, 0, "NonWhitespaceJoin returns correct value -- with args", actual)
}

// ── IsEmptyOrWhitespace / IsNullOrEmpty / IsNullOrEmptyOrWhitespace ──

func Test_IsEmptyOrWhitespace(t *testing.T) {
	// Act
	actual := args.Map{
		"empty": strutilinternal.IsEmptyOrWhitespace(""),
		"ws":    strutilinternal.IsEmptyOrWhitespace("  "),
		"val":   strutilinternal.IsEmptyOrWhitespace("a"),
	}

	// Assert
	expected := args.Map{
		"empty": true,
		"ws": true,
		"val": false,
	}
	expected.ShouldBeEqual(t, 0, "IsEmptyOrWhitespace returns empty -- with args", actual)
}

func Test_IsNullOrEmpty(t *testing.T) {
	// Act
	actual := args.Map{
		"nil":   strutilinternal.IsNullOrEmpty(nil),
		"empty": strutilinternal.IsNullOrEmpty(ptrStr("")),
		"val":   strutilinternal.IsNullOrEmpty(ptrStr("a")),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"empty": true,
		"val": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmpty returns empty -- with args", actual)
}

func Test_IsNullOrEmptyOrWhitespace(t *testing.T) {
	// Act
	actual := args.Map{
		"nil": strutilinternal.IsNullOrEmptyOrWhitespace(nil),
		"ws":  strutilinternal.IsNullOrEmptyOrWhitespace(ptrStr("  ")),
		"val": strutilinternal.IsNullOrEmptyOrWhitespace(ptrStr("a")),
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"ws": true,
		"val": false,
	}
	expected.ShouldBeEqual(t, 0, "IsNullOrEmptyOrWhitespace returns empty -- with args", actual)
}

func ptrStr(s string) *string { return &s }

// ── ReplaceTemplateMap curly ──

func Test_ReplaceTemplateMap_Curly(t *testing.T) {
	// Arrange
	result := strutilinternal.ReplaceTemplateMap(
		true,
		"Hello {name}, you are {age}",
		map[string]string{"name": "Alice", "age": "30"},
	)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "Hello Alice, you are 30"}
	expected.ShouldBeEqual(t, 0, "ReplaceTemplateMap returns correct value -- curly", actual)
}
