package coreutilstests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

// ==========================================
// IsNotEmpty
// ==========================================

func Test_IsNotEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsNotEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsNotEmpty(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsDefined
// ==========================================

func Test_IsDefined_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsDefinedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsDefined(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsStarts
// ==========================================

func Test_IsStarts_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsStartsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startsWith, _ := input.GetAsString("startsWith")

		// Act
		result := stringutil.IsStarts(content, startsWith)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsEnds
// ==========================================

func Test_IsEnds_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsEndsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		endsWith, _ := input.GetAsString("endsWith")

		// Act
		result := stringutil.IsEnds(content, endsWith)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsStartsChar
// ==========================================

func Test_IsStartsChar_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsStartsCharTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		charRaw, _ := input.Get("char")
		char := charRaw.(byte)

		// Act
		result := stringutil.IsStartsChar(content, char)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsEndsChar
// ==========================================

func Test_IsEndsChar_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsEndsCharTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		charRaw, _ := input.Get("char")
		char := charRaw.(byte)

		// Act
		result := stringutil.IsEndsChar(content, char)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsStartsRune
// ==========================================

func Test_IsStartsRune_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsStartsRuneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		runeRaw, _ := input.Get("rune")
		r := runeRaw.(rune)

		// Act
		result := stringutil.IsStartsRune(content, r)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsEndsRune
// ==========================================

func Test_IsEndsRune_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsEndsRuneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		runeRaw, _ := input.Get("rune")
		r := runeRaw.(rune)

		// Act
		result := stringutil.IsEndsRune(content, r)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsStartsAndEndsChar
// ==========================================

func Test_IsStartsAndEndsChar_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsStartsAndEndsCharTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startCharRaw, _ := input.Get("startChar")
		endCharRaw, _ := input.Get("endChar")
		startChar := startCharRaw.(byte)
		endChar := endCharRaw.(byte)

		// Act
		result := stringutil.IsStartsAndEndsChar(content, startChar, endChar)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsStartsAndEndsWith
// ==========================================

func Test_IsStartsAndEndsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsStartsAndEndsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startsWith, _ := input.GetAsString("startsWith")
		endsWith, _ := input.GetAsString("endsWith")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true

		// Act
		result := stringutil.IsStartsAndEndsWith(content, startsWith, endsWith, isIgnoreCase)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsStartsAndEnds
// ==========================================

func Test_IsStartsAndEnds_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsStartsAndEndsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startsWith, _ := input.GetAsString("startsWith")
		endsWith, _ := input.GetAsString("endsWith")

		// Act
		result := stringutil.IsStartsAndEnds(content, startsWith, endsWith)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsAnyStartsWith
// ==========================================

func Test_IsAnyStartsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsAnyStartsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true
		termsRaw, _ := input.Get("terms")
		terms := termsRaw.([]string)

		// Act
		result := stringutil.IsAnyStartsWith(content, isIgnoreCase, terms...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsAnyEndsWith
// ==========================================

func Test_IsAnyEndsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsAnyEndsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true
		termsRaw, _ := input.Get("terms")
		terms := termsRaw.([]string)

		// Act
		result := stringutil.IsAnyEndsWith(content, isIgnoreCase, terms...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// FirstCharOrDefault
// ==========================================

func Test_FirstCharOrDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range extFirstCharTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.FirstCharOrDefault(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// LastCharOrDefault
// ==========================================

func Test_LastCharOrDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range extLastCharTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.LastCharOrDefault(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// ClonePtr
// ==========================================

func Test_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilRaw, _ := input.Get("isNil")
		isNil := isNilRaw == true

		var result *string
		if isNil {
			// Act
			result = stringutil.ClonePtr(nil)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result == nil))
		} else {
			value, _ := input.GetAsString("value")
			// Act
			result = stringutil.ClonePtr(&value)

			// Assert
			testCase.ShouldBeEqual(t, caseIndex, *result)
		}
	}
}

// ==========================================
// SafeClonePtr
// ==========================================

func Test_SafeClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range extSafeClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilRaw, _ := input.Get("isNil")
		isNil := isNilRaw == true

		var result *string
		if isNil {
			// Act
			result = stringutil.SafeClonePtr(nil)
		} else {
			value, _ := input.GetAsString("value")
			// Act
			result = stringutil.SafeClonePtr(&value)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, *result)
	}
}

// ==========================================
// Ptr functions
// ==========================================

func Test_PtrFunctions_Verification(t *testing.T) {
	for caseIndex, testCase := range extPtrFunctionsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		funcName, _ := input.GetAsString("func")
		isNilRaw, _ := input.Get("isNil")
		isNil := isNilRaw == true

		var actual bool
		if isNil {
			// Act
			switch funcName {
			case "IsEmptyPtr":
				actual = stringutil.IsEmptyPtr(nil)
			case "IsBlankPtr":
				actual = stringutil.IsBlankPtr(nil)
			case "IsEmptyOrWhitespacePtr":
				actual = stringutil.IsEmptyOrWhitespacePtr(nil)
			case "IsNullOrEmptyPtr":
				actual = stringutil.IsNullOrEmptyPtr(nil)
			case "IsDefinedPtr":
				actual = stringutil.IsDefinedPtr(nil)
			}
		} else {
			value, _ := input.GetAsString("value")
			// Act
			switch funcName {
			case "IsEmptyPtr":
				actual = stringutil.IsEmptyPtr(&value)
			case "IsBlankPtr":
				actual = stringutil.IsBlankPtr(&value)
			case "IsEmptyOrWhitespacePtr":
				actual = stringutil.IsEmptyOrWhitespacePtr(&value)
			case "IsNullOrEmptyPtr":
				actual = stringutil.IsNullOrEmptyPtr(&value)
			case "IsDefinedPtr":
				actual = stringutil.IsDefinedPtr(&value)
			}
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", actual))
	}
}

// ==========================================
// ToBool
// ==========================================

func Test_ToBool_Verification(t *testing.T) {
	for caseIndex, testCase := range extToBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ToBool(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// ToByte
// ==========================================

func Test_ToByte_Verification(t *testing.T) {
	for caseIndex, testCase := range extToByteTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		defRaw, _ := input.Get("def")
		defVal := defRaw.(byte)

		// Act
		result := stringutil.ToByte(inputStr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// ToByteDefault
// ==========================================

func Test_ToByteDefault_Verification(t *testing.T) {
	for caseIndex, testCase := range extToByteDefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ToByteDefault(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// ToInt
// ==========================================

func Test_ToInt_Verification(t *testing.T) {
	for caseIndex, testCase := range extToIntTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		defRaw, _ := input.Get("def")
		defVal := defRaw.(int)

		// Act
		result := stringutil.ToInt(inputStr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// ToIntDef / ToIntDefault
// ==========================================

func Test_ToIntDef_Verification(t *testing.T) {
	// Arrange
	// Act
	result1 := stringutil.ToIntDef("42")
	result2 := stringutil.ToIntDef("abc")

	// Assert
	actual := args.Map{"result": result1 != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntDef(42) expected 42", actual)
	actual = args.Map{"result": result2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntDef(abc) expected 0", actual)
}

func Test_ToIntDefault_Verification(t *testing.T) {
	// Arrange
	// Act
	result1 := stringutil.ToIntDefault("42")
	result2 := stringutil.ToIntDefault("abc")

	// Assert
	actual := args.Map{"result": result1 != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntDefault(42) expected 42", actual)
	actual = args.Map{"result": result2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntDefault(abc) expected 0", actual)
}

// ==========================================
// ToInt8 / ToInt8Def
// ==========================================

func Test_ToInt8_Verification(t *testing.T) {
	for caseIndex, testCase := range extToInt8TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		defRaw, _ := input.Get("def")
		defVal := defRaw.(int8)

		// Act
		result := stringutil.ToInt8(inputStr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

func Test_ToInt8Def_Verification(t *testing.T) {
	// Arrange
	// Act
	result1 := stringutil.ToInt8Def("50")
	result2 := stringutil.ToInt8Def("abc")

	// Assert
	actual := args.Map{"result": result1 != 50}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt8Def(50) expected 50", actual)
	actual = args.Map{"result": result2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt8Def(abc) expected 0", actual)
}

// ==========================================
// ToInt16 / ToInt16Default
// ==========================================

func Test_ToInt16_Verification(t *testing.T) {
	for caseIndex, testCase := range extToInt16TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		defRaw, _ := input.Get("def")
		defVal := defRaw.(int16)

		// Act
		result := stringutil.ToInt16(inputStr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

func Test_ToInt16Default_Verification(t *testing.T) {
	// Arrange
	// Act
	result1 := stringutil.ToInt16Default("1000")
	result2 := stringutil.ToInt16Default("abc")
	result3 := stringutil.ToInt16Default("40000") // overflow

	// Assert
	actual := args.Map{"result": result1 != 1000}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt16Default(1000) expected 1000", actual)
	actual = args.Map{"result": result2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt16Default(abc) expected 0", actual)
	actual = args.Map{"result": result3 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt16Default(40000) expected 0", actual)
}

// ==========================================
// ToInt32 / ToInt32Def
// ==========================================

func Test_ToInt32_Verification(t *testing.T) {
	for caseIndex, testCase := range extToInt32TestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		defRaw, _ := input.Get("def")
		defVal := defRaw.(int32)

		// Act
		result := stringutil.ToInt32(inputStr, defVal)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

func Test_ToInt32Def_Verification(t *testing.T) {
	// Arrange
	// Act
	result1 := stringutil.ToInt32Def("65536")
	result2 := stringutil.ToInt32Def("abc")

	// Assert
	actual := args.Map{"result": result1 != 65536}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt32Def(65536) expected 65536", actual)
	actual = args.Map{"result": result2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToInt32Def(abc) expected 0", actual)
}

// ==========================================
// ToUint16Default
// ==========================================

func Test_ToUint16Default_Verification(t *testing.T) {
	for caseIndex, testCase := range extToUint16DefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ToUint16Default(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// ToUint32Default
// ==========================================

func Test_ToUint32Default_Verification(t *testing.T) {
	for caseIndex, testCase := range extToUint32DefaultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ToUint32Default(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%d", result))
	}
}

// ==========================================
// AnyToString
// ==========================================

func Test_AnyToString_Verification(t *testing.T) {
	for caseIndex, testCase := range extAnyToStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, _ := input.Get("input")

		// Act
		result := stringutil.AnyToString(inputRaw)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// AnyToStringNameField
// ==========================================

func Test_AnyToStringNameField_Verification(t *testing.T) {
	// Arrange
	// Act
	resultNil := stringutil.AnyToStringNameField(nil)
	resultVal := stringutil.AnyToStringNameField(42)

	// Assert
	actual := args.Map{"result": resultNil != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField(nil) expected empty, got ''", actual)
	actual = args.Map{"result": resultVal == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyToStringNameField(42) should not be empty", actual)
}

// ==========================================
// AnyToTypeString
// ==========================================

func Test_AnyToTypeString_Verification(t *testing.T) {
	// Arrange
	// Act
	result := stringutil.AnyToTypeString(42)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyToTypeString(42) should not be empty", actual)
}

// ==========================================
// MaskLine
// ==========================================

func Test_MaskLine_Verification(t *testing.T) {
	for caseIndex, testCase := range extMaskLineTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mask, _ := input.GetAsString("mask")
		line, _ := input.GetAsString("line")

		// Act
		result := stringutil.MaskLine(mask, line)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// MaskTrimLine
// ==========================================

func Test_MaskTrimLine_Verification(t *testing.T) {
	// Arrange
	// Act
	result1 := stringutil.MaskTrimLine("----------", "  hi  ")
	result2 := stringutil.MaskTrimLine("----------", "  ")
	result3 := stringutil.MaskTrimLine("---", "hello world")

	// Assert
	actual := args.Map{"result": result1 != "hi--------"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine expected 'hi--------', got ''", actual)
	actual = args.Map{"result": result2 != "----------"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine whitespace expected mask, got ''", actual)
	actual = args.Map{"result": result3 != "hello world"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MaskTrimLine long line expected 'hello world', got ''", actual)
}

// ==========================================
// MaskLines
// ==========================================

func Test_MaskLines_Verification(t *testing.T) {
	for caseIndex, testCase := range extMaskLinesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mask, _ := input.GetAsString("mask")
		linesRaw, _ := input.Get("lines")
		lines := linesRaw.([]string)

		// Act
		result := stringutil.MaskLines(mask, lines...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, strings.Join(result, ","))
	}
}

// ==========================================
// MaskTrimLines
// ==========================================

func Test_MaskTrimLines_Verification(t *testing.T) {
	for caseIndex, testCase := range extMaskTrimLinesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		mask, _ := input.GetAsString("mask")
		linesRaw, _ := input.Get("lines")
		lines := linesRaw.([]string)

		// Act
		result := stringutil.MaskTrimLines(mask, lines...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, strings.Join(result, ","))
	}
}

// ==========================================
// IsContains
// ==========================================

func Test_IsContains_Verification(t *testing.T) {
	for caseIndex, testCase := range extIsContainsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		linesRaw, _ := input.Get("lines")
		find, _ := input.GetAsString("find")
		startRaw, _ := input.Get("start")
		start := startRaw.(int)
		caseSensitiveRaw, _ := input.Get("caseSensitive")
		caseSensitive := caseSensitiveRaw == true

		var lines []string
		if linesRaw != nil {
			lines = linesRaw.([]string)
		}

		// Act
		result := stringutil.IsContains(lines, find, start, caseSensitive)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

// ==========================================
// IsContainsPtr / IsContainsPtrSimple
// ==========================================

func Test_IsContainsPtr_Verification(t *testing.T) {
	// Arrange
	lines := []string{"Hello", "World"}
	find := "hello"

	// Act - case insensitive
	result1 := stringutil.IsContainsPtr(&lines, &find, 0, false)

	// Assert
	actual := args.Map{"result": result1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr case-insensitive should find 'hello'", actual)

	// Act - case sensitive
	result2 := stringutil.IsContainsPtr(&lines, &find, 0, true)

	// Assert
	actual = args.Map{"result": result2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr case-sensitive should not find 'hello'", actual)

	// Act - nil lines
	result3 := stringutil.IsContainsPtr(nil, &find, 0, true)

	// Assert
	actual = args.Map{"result": result3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtr nil should return false", actual)
}

func Test_IsContainsPtrSimple_Verification(t *testing.T) {
	// Arrange
	lines := []string{"Hello", "World"}

	// Act - case insensitive
	result1 := stringutil.IsContainsPtrSimple(&lines, "hello", 0, false)

	// Assert
	actual := args.Map{"result": result1}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple case-insensitive should find 'hello'", actual)

	// Act - nil
	result2 := stringutil.IsContainsPtrSimple(nil, "x", 0, true)

	// Assert
	actual = args.Map{"result": result2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "IsContainsPtrSimple nil should return false", actual)
}

// ==========================================
// SplitLeftRightTrimmed
// ==========================================

func Test_SplitLeftRightTrimmed_Verification(t *testing.T) {
	for caseIndex, testCase := range extSplitLeftRightTrimmedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		separator, _ := input.GetAsString("separator")

		// Act
		left, right := stringutil.SplitLeftRightTrimmed(inputStr, separator)
		actual := args.Map{
			"left":  left,
			"right": right,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// SplitFirstLast
// ==========================================

func Test_SplitFirstLast_Verification(t *testing.T) {
	for caseIndex, testCase := range extSplitFirstLastTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		separator, _ := input.GetAsString("separator")

		// Act
		first, last := stringutil.SplitFirstLast(inputStr, separator)
		actual := args.Map{
			"first": first,
			"last":  last,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================
// SafeSubstringStarts
// ==========================================

func Test_SafeSubstringStarts_Verification(t *testing.T) {
	for caseIndex, testCase := range extSafeSubstringStartsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startRaw, _ := input.Get("start")
		start := startRaw.(int)

		// Act
		result := stringutil.SafeSubstringStarts(content, start)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// SafeSubstringEnds
// ==========================================

func Test_SafeSubstringEnds_Verification(t *testing.T) {
	for caseIndex, testCase := range extSafeSubstringEndsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		endRaw, _ := input.Get("end")
		end := endRaw.(int)

		// Act
		result := stringutil.SafeSubstringEnds(content, end)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// RemoveManyBySplitting
// ==========================================

func Test_RemoveManyBySplitting_Verification(t *testing.T) {
	for caseIndex, testCase := range extRemoveManyBySplittingTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		splitsBy, _ := input.GetAsString("splitsBy")
		removesRaw, _ := input.Get("removes")
		removes := removesRaw.([]string)

		// Act
		result := stringutil.RemoveManyBySplitting(content, splitsBy, removes...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, strings.Join(result, ","))
	}
}

// ==========================================
// ReplaceTemplate.CurlyOne
// ==========================================

func Test_ReplaceTemplate_CurlyOne_Verification(t *testing.T) {
	for caseIndex, testCase := range extReplaceTemplateCurlyOneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		format, _ := input.GetAsString("format")
		key, _ := input.GetAsString("key")
		valueRaw, _ := input.Get("value")

		// Act
		result := stringutil.ReplaceTemplate.CurlyOne(format, key, valueRaw)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// ReplaceTemplate.CurlyTwo
// ==========================================

func Test_ReplaceTemplate_CurlyTwo_Verification(t *testing.T) {
	for caseIndex, testCase := range extReplaceTemplateCurlyTwoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		format, _ := input.GetAsString("format")
		key1, _ := input.GetAsString("key1")
		val1Raw, _ := input.Get("val1")
		key2, _ := input.GetAsString("key2")
		val2Raw, _ := input.Get("val2")

		// Act
		result := stringutil.ReplaceTemplate.CurlyTwo(format, key1, val1Raw, key2, val2Raw)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// ReplaceTemplate.ReplaceWhiteSpaces
// ==========================================

func Test_ReplaceTemplate_ReplaceWhiteSpaces_Verification(t *testing.T) {
	for caseIndex, testCase := range extReplaceWhiteSpacesTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ReplaceTemplate.ReplaceWhiteSpaces(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

// ==========================================
// ReplaceTemplate additional methods
// ==========================================

func Test_ReplaceTemplate_DirectOne_Verification(t *testing.T) {
	// Arrange
	format := "Hello NAME!"

	// Act
	result := stringutil.ReplaceTemplate.DirectOne(format, "NAME", "World")

	// Assert
	actual := args.Map{"result": result != "Hello World!"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectOne expected 'Hello World!', got ''", actual)

	// Act - empty
	result2 := stringutil.ReplaceTemplate.DirectOne("", "NAME", "World")

	// Assert
	actual = args.Map{"result": result2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectOne empty format expected '', got ''", actual)
}

func Test_ReplaceTemplate_DirectTwoItem_Verification(t *testing.T) {
	// Arrange
	format := "NAME is AGE years old"

	// Act
	result := stringutil.ReplaceTemplate.DirectTwoItem(format, "NAME", "Alice", "AGE", 30)

	// Assert
	actual := args.Map{"result": result != "Alice is 30 years old"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectTwoItem expected 'Alice is 30 years old', got ''", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingMap_Verification(t *testing.T) {
	// Arrange
	format := "Hello NAME from PLACE"
	m := map[string]string{"NAME": "Alice", "PLACE": "NYC"}

	// Act
	result := stringutil.ReplaceTemplate.DirectKeyUsingMap(format, m)

	// Assert
	actual := args.Map{"result": result != "Hello Alice from NYC"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap expected 'Hello Alice from NYC', got ''", actual)

	// Act - empty map
	result2 := stringutil.ReplaceTemplate.DirectKeyUsingMap(format, map[string]string{})

	// Assert
	actual = args.Map{"result": result2 != format}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMap empty map should return format unchanged", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingMapTrim_Verification(t *testing.T) {
	// Arrange
	format := "  Hello NAME  "
	m := map[string]string{"NAME": "World"}

	// Act
	result := stringutil.ReplaceTemplate.DirectKeyUsingMapTrim(format, m)

	// Assert
	actual := args.Map{"result": result != "Hello World"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingMapTrim expected 'Hello World', got ''", actual)
}

func Test_ReplaceTemplate_DirectKeyUsingKeyVal_Verification(t *testing.T) {
	// Arrange
	format := "Hello NAME from PLACE"

	// Act
	result := stringutil.ReplaceTemplate.DirectKeyUsingKeyVal(
		format,
		stringutil.KeyValReplacer{Key: "NAME", Value: "Alice"},
		stringutil.KeyValReplacer{Key: "PLACE", Value: "NYC"},
	)

	// Assert
	actual := args.Map{"result": result != "Hello Alice from NYC"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal expected 'Hello Alice from NYC', got ''", actual)

	// Act - empty
	result2 := stringutil.ReplaceTemplate.DirectKeyUsingKeyVal(format)

	// Assert
	actual = args.Map{"result": result2 != format}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DirectKeyUsingKeyVal empty should return format unchanged", actual)
}

func Test_ReplaceTemplate_CurlyKeyUsingMap_Verification(t *testing.T) {
	// Arrange
	format := "{greeting} {name}!"
	m := map[string]string{"greeting": "Hi", "name": "Bob"}

	// Act
	result := stringutil.ReplaceTemplate.CurlyKeyUsingMap(format, m)

	// Assert
	actual := args.Map{"result": result != "Hi Bob!"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap expected 'Hi Bob!', got ''", actual)
}

func Test_ReplaceTemplate_UsingMapOptions_Verification(t *testing.T) {
	// Arrange
	format := "{key1} and KEY2"
	m := map[string]string{"key1": "A", "KEY2": "B"}

	// Act - curly mode
	result1 := stringutil.ReplaceTemplate.UsingMapOptions(true, format, m)

	// Assert
	actual := args.Map{"result": strings.Contains(result1, "A")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions curly should replace {key1}", actual)

	// Act - direct mode
	result2 := stringutil.ReplaceTemplate.UsingMapOptions(false, format, m)

	// Assert
	actual = args.Map{"result": strings.Contains(result2, "B")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions direct should replace KEY2", actual)

	// Act - empty
	result3 := stringutil.ReplaceTemplate.UsingMapOptions(true, "", m)

	// Assert
	actual = args.Map{"result": result3 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "UsingMapOptions empty format should return empty", actual)
}

func Test_ReplaceTemplate_ReplaceWhiteSpacesToSingle_Verification(t *testing.T) {
	// Arrange
	// Act
	result := stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("  hello   world  ")

	// Assert
	actual := args.Map{"result": result != "hello world"}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle expected 'hello world', got ''", actual)
}

// ==========================================
// SplitLeftRightType / SplitLeftRightTypeTrimmed
// ==========================================

func Test_SplitLeftRightType_Verification(t *testing.T) {
	// Arrange
	// Act
	result := stringutil.SplitLeftRightType("key=value", "=")

	actual := args.Map{
		"left":  result.Left,
		"right": result.Right,
	}
	expected := args.Map{
		"left":  "key",
		"right": "value",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightType returns correct parts -- key=value", actual)
}

func Test_SplitLeftRightTypeTrimmed_Verification(t *testing.T) {
	// Arrange
	// Act
	result := stringutil.SplitLeftRightTypeTrimmed(" key = value ", "=")

	actual := args.Map{
		"left":  result.Left,
		"right": result.Right,
	}
	expected := args.Map{
		"left":  "key",
		"right": "value",
	}
	expected.ShouldBeEqual(t, 0, "SplitLeftRightTypeTrimmed returns correct parts -- trimmed key=value", actual)
}

// ==========================================
// ToIntUsingRegexMatch
// ==========================================

func Test_ToIntUsingRegexMatch_Verification(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^\d+$`)

	// Act
	result1 := stringutil.ToIntUsingRegexMatch(re, "42")
	result2 := stringutil.ToIntUsingRegexMatch(re, "abc")
	result3 := stringutil.ToIntUsingRegexMatch(nil, "42")

	// Assert
	actual := args.Map{"result": result1 != 42}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch(42) expected 42", actual)
	actual = args.Map{"result": result2 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch(abc) expected 0", actual)
	actual = args.Map{"result": result3 != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "ToIntUsingRegexMatch(nil re) expected 0", actual)
}

// ==========================================
// SplitContentsByWhitespaceConditions
// ==========================================

func Test_SplitContentsByWhitespaceConditions_Verification(t *testing.T) {
	// Arrange
	input := "  Hello   World  hello  "

	// Act - basic split
	result1 := stringutil.SplitContentsByWhitespaceConditions(input, false, false, false, false, false)

	// Assert
	actual := args.Map{"result": len(result1) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Basic split expected 3 items", actual)

	// Act - sorted
	result2 := stringutil.SplitContentsByWhitespaceConditions(input, false, false, true, false, false)

	// Assert
	actual = args.Map{"result": len(result2) < 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Sorted split should have items", actual)

	// Act - unique + lowercase
	result3 := stringutil.SplitContentsByWhitespaceConditions(input, false, false, false, true, true)

	// Assert
	actual = args.Map{"result": len(result3) != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Unique+lowercase expected 2 items", actual)
}
