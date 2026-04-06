package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ========================================
// S19: ValidValue, ValidValues, ValueStatus,
//   TextWithLineNumber, utils, CloneSlice/If,
//   AnyToString
// ========================================

// --- ValidValue ---

func Test_C79_ValidValue_NewValidValue(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValue", func() {
		// Arrange & Act
		vv := corestr.NewValidValue("hello")

		// Assert
		actual := args.Map{"result": vv.Value != "hello" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewValidValue mismatch", actual)
	})
}

func Test_C79_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValueEmpty", func() {
		// Arrange & Act
		vv := corestr.NewValidValueEmpty()

		// Assert
		actual := args.Map{"result": vv.Value != "" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty valid value", actual)
	})
}

func Test_C79_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_InvalidValidValue", func() {
		// Arrange & Act
		vv := corestr.InvalidValidValue("err msg")

		// Assert
		actual := args.Map{"result": vv.IsValid || vv.Message != "err msg"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid with message", actual)
	})
}

func Test_C79_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_InvalidValidValueNoMessage", func() {
		// Arrange & Act
		vv := corestr.InvalidValidValueNoMessage()

		// Assert
		actual := args.Map{"result": vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C79_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValueUsingAny", func() {
		// Arrange & Act
		vv := corestr.NewValidValueUsingAny(false, true, 42)

		// Assert
		actual := args.Map{"result": vv.Value != "42" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '42' valid, got ''", actual)
	})
}

func Test_C79_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValueUsingAnyAutoValid", func() {
		// Arrange & Act
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")

		// Assert
		actual := args.Map{"result": vv.Value == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty value", actual)
	})
}

func Test_C79_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		bytes1 := vv.ValueBytesOnce()
		bytes2 := vv.ValueBytesOnce() // should reuse cached

		// Assert
		actual := args.Map{"result": string(bytes1) != "abc" || string(bytes2) != "abc"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "bytes mismatch", actual)
	})
}

func Test_C79_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueBytesOncePtr", func() {
		// Arrange
		vv := corestr.NewValidValue("xyz")

		// Act
		result := vv.ValueBytesOncePtr()

		// Assert
		actual := args.Map{"result": string(result) != "xyz"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "bytes mismatch", actual)
	})
}

func Test_C79_ValidValue_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsEmpty_IsWhitespace", func() {
		// Arrange
		empty := corestr.NewValidValue("")
		ws := corestr.NewValidValue("  ")
		val := corestr.NewValidValue("x")

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ws.IsWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
		actual = args.Map{"result": val.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_C79_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue(" hello ")

		// Act & Assert
		actual := args.Map{"result": vv.Trim() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "trim mismatch", actual)
	})
}

func Test_C79_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		valid := corestr.NewValidValue("x")
		invalid := corestr.InvalidValidValue("err")

		// Act & Assert
		actual := args.Map{"result": valid.HasValidNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": invalid.HasValidNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C79_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		valid := corestr.NewValidValue("x")
		ws := corestr.NewValidValue("  ")

		// Act & Assert
		actual := args.Map{"result": valid.HasValidNonWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ws.HasValidNonWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for whitespace", actual)
	})
}

func Test_C79_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueBool", func() {
		// Arrange
		trueVV := corestr.NewValidValue("true")
		falseVV := corestr.NewValidValue("false")
		invalidVV := corestr.NewValidValue("xyz")
		emptyVV := corestr.NewValidValue("")

		// Act & Assert
		actual := args.Map{"result": trueVV.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": falseVV.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": invalidVV.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
		actual = args.Map{"result": emptyVV.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C79_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueInt", func() {
		// Arrange
		vv := corestr.NewValidValue("42")
		invalid := corestr.NewValidValue("abc")

		// Act & Assert
		actual := args.Map{"result": vv.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": invalid.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_C79_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueDefInt", func() {
		// Arrange
		vv := corestr.NewValidValue("10")
		invalid := corestr.NewValidValue("x")

		// Act & Assert
		actual := args.Map{"result": vv.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		actual = args.Map{"result": invalid.ValueDefInt() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueByte", func() {
		// Arrange
		valid := corestr.NewValidValue("100")
		overflow := corestr.NewValidValue("300")
		negative := corestr.NewValidValue("-5")
		invalid := corestr.NewValidValue("abc")

		// Act
		actual := args.Map{
			"valid":    valid.ValueByte(0),
			"overflow": overflow.ValueByte(0),
			"negative": negative.ValueByte(0),
			"invalid":  invalid.ValueByte(0),
		}

		// Assert
		expected := args.Map{
			"valid":    byte(100),
			"overflow": byte(255),
			"negative": byte(0),
			"invalid":  byte(0),
		}
		expected.ShouldBeEqual(t, 0, "ValueByte returns correct byte -- various inputs", actual)
	})
}

func Test_C79_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueDefByte", func() {
		// Arrange
		valid := corestr.NewValidValue("50")
		overflow := corestr.NewValidValue("999")
		negative := corestr.NewValidValue("-1")

		// Act & Assert
		actual := args.Map{"result": valid.ValueDefByte() != 50}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
		actual = args.Map{"result": overflow.ValueDefByte() != 255}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255", actual)
		actual = args.Map{"result": negative.ValueDefByte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")
		invalid := corestr.NewValidValue("abc")

		// Act & Assert
		actual := args.Map{"result": vv.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual = args.Map{"result": invalid.ValueFloat64(1.5) != 1.5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	})
}

func Test_C79_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueDefFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("2.5")

		// Act & Assert
		actual := args.Map{"result": vv.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C79_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		valid := corestr.NewValidValue("x")
		empty := corestr.NewValidValue("")

		// Act & Assert
		actual := args.Map{"result": valid.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": empty.HasSafeNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C79_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Is", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act & Assert
		actual := args.Map{"result": vv.Is("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.Is("world")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C79_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act & Assert
		actual := args.Map{"result": vv.IsAnyOf("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsAnyOf("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": vv.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_C79_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act & Assert
		actual := args.Map{"result": vv.IsContains("world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C79_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act & Assert
		actual := args.Map{"result": vv.IsAnyContains("xyz", "world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsAnyContains("abc")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": vv.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_C79_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act & Assert
		actual := args.Map{"result": vv.IsEqualNonSensitive("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C79_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsRegexMatches", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		actual := args.Map{"result": vv.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C79_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindString", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		actual := args.Map{"result": vv.RegexFindString(re) != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '123'", actual)
		actual = args.Map{"result": vv.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C79_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)

		// Act
		items, hasAny := vv.RegexFindAllStringsWithFlag(re, -1)

		// Assert
		actual := args.Map{"result": hasAny || len(items) != 3}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 3 matches", actual)
	})
}

func Test_C79_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		items, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)

		// Assert
		actual := args.Map{"result": hasAny || len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C79_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)

		// Act
		items := vv.RegexFindAllStrings(re, -1)

		// Assert
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindAllStrings_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act & Assert
		actual := args.Map{"result": len(vv.RegexFindAllStrings(nil, -1)) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C79_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")

		// Act
		result := vv.Split(",")

		// Assert
		actual := args.Map{"result": len(result) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C79_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_SplitNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("a::b")

		// Act
		result := vv.SplitNonEmpty("::")

		// Assert
		actual := args.Map{"result": len(result) < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C79_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_SplitTrimNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("a , , b")

		// Act
		result := vv.SplitTrimNonWhitespace(",")

		// Assert
		actual := args.Map{"result": len(result) < 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_C79_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		cloned := vv.Clone()

		// Assert
		actual := args.Map{"result": cloned == nil || cloned.Value != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

func Test_C79_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Clone_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		cloned := vv.Clone()

		// Assert
		actual := args.Map{"result": cloned != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C79_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_String", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act & Assert
		actual := args.Map{"result": vv.String() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "String mismatch", actual)
	})
}

func Test_C79_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_String_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act & Assert
		actual := args.Map{"result": vv.String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C79_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		result := vv.FullString()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C79_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_FullString_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act & Assert
		actual := args.Map{"result": vv.FullString() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C79_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Clear_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		vv.Clear()

		// Assert
		actual := args.Map{"result": vv.Value != "" || vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cleared", actual)
	})
}

func Test_C79_ValidValue_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Dispose_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act — should not panic
		vv.Clear()
		vv.Dispose()
	})
}

func Test_C79_ValidValue_Json_Serialize(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Json_Serialize", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		jsonResult := vv.Json()
		bytes, err := vv.Serialize()

		// Assert
		actual := args.Map{"result": jsonResult.HasError()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "json error", actual)
		actual = args.Map{"result": err != nil || len(bytes) == 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "serialize error", actual)
	})
}

func Test_C79_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		jsonResult := vv.JsonPtr()
		target := &corestr.ValidValue{}

		// Act
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
		actual = args.Map{"result": result.Value != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value mismatch", actual)
	})
}

func Test_C79_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Deserialize", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		var target corestr.ValidValue
		err := vv.Deserialize(&target)

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "error:", actual)
	})
}

// --- ValidValues ---

func Test_C79_ValidValues_NewValidValues(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_NewValidValues", func() {
		// Arrange & Act
		vvs := corestr.NewValidValues(5)

		// Assert
		actual := args.Map{"result": vvs == nil || vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C79_ValidValues_EmptyValidValues(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_EmptyValidValues", func() {
		// Arrange & Act
		vvs := corestr.EmptyValidValues()

		// Assert
		actual := args.Map{"result": vvs == nil || !vvs.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C79_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_NewValidValuesUsingValues", func() {
		// Arrange & Act
		vvs := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)

		// Assert
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_NewValidValuesUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_NewValidValuesUsingValues_Empty", func() {
		// Arrange & Act
		vvs := corestr.NewValidValuesUsingValues()

		// Assert
		actual := args.Map{"result": vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Add", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.Add("a").Add("b")

		// Assert
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddFull", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddFull(false, "err", "msg")

		// Assert
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_ValidValues_Count_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Count_HasAnyItem", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")

		// Act & Assert
		actual := args.Map{"result": vvs.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": vvs.HasAnyItem()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C79_ValidValues_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_LastIndex_HasIndex", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")

		// Act & Assert
		actual := args.Map{"result": vvs.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": vvs.HasIndex(0)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for 0", actual)
		actual = args.Map{"result": vvs.HasIndex(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for 5", actual)
	})
}

func Test_C79_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_SafeValueAt", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")

		// Act & Assert
		actual := args.Map{"result": vvs.SafeValueAt(0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		actual = args.Map{"result": vvs.SafeValueAt(5) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for out of range", actual)
	})
}

func Test_C79_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_SafeValidValueAt", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		vvs.AddFull(false, "b", "err")

		// Act & Assert
		actual := args.Map{"result": vvs.SafeValidValueAt(0) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a'", actual)
		actual = args.Map{"result": vvs.SafeValidValueAt(1) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for invalid value", actual)
	})
}

func Test_C79_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_SafeValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")

		// Act
		result := vvs.SafeValuesAtIndexes(0, 2)

		// Assert
		actual := args.Map{"result": len(result) != 2 || result[0] != "a" || result[1] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected [a, c]", actual)
	})
}

func Test_C79_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_SafeValidValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").AddFull(false, "b", "err")

		// Act
		result := vvs.SafeValidValuesAtIndexes(0, 1)

		// Assert
		actual := args.Map{"result": len(result) != 2 || result[0] != "a" || result[1] != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected result:", actual)
	})
}

func Test_C79_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Strings", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("x")

		// Act
		result := vvs.Strings()

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_ValidValues_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Strings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act & Assert
		actual := args.Map{"result": len(vvs.Strings()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_FullStrings", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("x")

		// Act
		result := vvs.FullStrings()

		// Assert
		actual := args.Map{"result": len(result) != 1 || result[0] == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty full string", actual)
	})
}

func Test_C79_ValidValues_FullStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_FullStrings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act & Assert
		actual := args.Map{"result": len(vvs.FullStrings()) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_String", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("x")

		// Act
		result := vvs.String()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C79_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Find", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")

		// Act
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})

		// Assert
		actual := args.Map{"result": len(found) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_ValidValues_Find_WithBreak(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Find_WithBreak", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b").Add("c")

		// Act
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, i == 0
		})

		// Assert
		actual := args.Map{"result": len(found) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 due to break", actual)
	})
}

func Test_C79_ValidValues_Find_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Find_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, false
		})

		// Assert
		actual := args.Map{"result": len(found) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Adds", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.Adds(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)

		// Assert
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Adds_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.Adds()

		// Assert
		actual := args.Map{"result": vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddsPtr", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		v := corestr.NewValidValue("x")

		// Act
		vvs.AddsPtr(v)

		// Assert
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_ValidValues_AddsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddsPtr_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddsPtr()

		// Assert
		actual := args.Map{"result": vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddValidValues", func() {
		// Arrange
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")

		// Act
		vvs1.AddValidValues(vvs2)

		// Assert
		actual := args.Map{"result": vvs1.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddValidValues_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddValidValues(nil)

		// Assert
		actual := args.Map{"result": vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_ConcatNew", func() {
		// Arrange
		vvs1 := corestr.EmptyValidValues()
		vvs1.Add("a")
		vvs2 := corestr.EmptyValidValues()
		vvs2.Add("b")

		// Act
		result := vvs1.ConcatNew(false, vvs2)

		// Assert
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_ConcatNew_EmptyWithClone(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_ConcatNew_EmptyWithClone", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")

		// Act
		result := vvs.ConcatNew(true)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_ConcatNew_EmptyNoClone", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")

		// Act
		result := vvs.ConcatNew(false)

		// Assert
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same instance", actual)
	})
}

func Test_C79_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashsetMap", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddHashsetMap(map[string]bool{"x": true, "y": false})

		// Assert
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashsetMap_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddHashsetMap(nil)

		// Assert
		actual := args.Map{"result": vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashset", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})

		// Act
		vvs.AddHashset(hs)

		// Assert
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashset_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddHashset(nil)

		// Assert
		actual := args.Map{"result": vvs.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C79_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Hashmap", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("k")

		// Act
		hm := vvs.Hashmap()

		// Assert
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Map", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("k")

		// Act
		m := vvs.Map()

		// Assert
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// --- ValueStatus ---

func Test_C79_ValueStatus_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_C79_ValueStatus_InvalidValueStatusNoMessage", func() {
		// Arrange & Act
		vs := corestr.InvalidValueStatusNoMessage()

		// Assert
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		actual = args.Map{"result": vs.Index != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
	})
}

func Test_C79_ValueStatus_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_C79_ValueStatus_InvalidValueStatus", func() {
		// Arrange & Act
		vs := corestr.InvalidValueStatus("err")

		// Assert
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		actual = args.Map{"result": vs.ValueValid.Message != "err"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "message mismatch", actual)
	})
}

func Test_C79_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C79_ValueStatus_Clone", func() {
		// Arrange
		vs := &corestr.ValueStatus{
			ValueValid: corestr.NewValidValue("x"),
			Index:      5,
		}

		// Act
		cloned := vs.Clone()

		// Assert
		actual := args.Map{"result": cloned.Index != 5 || cloned.ValueValid.Value != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

// --- TextWithLineNumber ---

func Test_C79_TWLN_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_HasLineNumber", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		invalid := &corestr.TextWithLineNumber{LineNumber: -1, Text: "x"}

		// Act & Assert
		actual := args.Map{"result": twln.HasLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": invalid.HasLineNumber()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for -1", actual)
	})
}

func Test_C79_TWLN_HasLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_HasLineNumber_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		actual := args.Map{"result": twln.HasLineNumber()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C79_TWLN_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsInvalidLineNumber", func() {
		// Arrange
		valid := &corestr.TextWithLineNumber{LineNumber: 1, Text: "x"}
		invalid := &corestr.TextWithLineNumber{LineNumber: -1, Text: "x"}

		// Act & Assert
		actual := args.Map{"result": valid.IsInvalidLineNumber()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": invalid.IsInvalidLineNumber()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C79_TWLN_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsInvalidLineNumber_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		actual := args.Map{"result": twln.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	})
}

func Test_C79_TWLN_Length(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_Length", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act & Assert
		actual := args.Map{"result": twln.Length() != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_C79_TWLN_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_Length_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		actual := args.Map{"result": twln.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for nil", actual)
	})
}

func Test_C79_TWLN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmpty", func() {
		// Arrange
		empty := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		valid := &corestr.TextWithLineNumber{LineNumber: 1, Text: "x"}

		// Act & Assert
		actual := args.Map{"result": empty.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": valid.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_C79_TWLN_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmpty_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		actual := args.Map{"result": twln.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	})
}

func Test_C79_TWLN_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmptyText", func() {
		// Arrange
		empty := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		valid := &corestr.TextWithLineNumber{LineNumber: 1, Text: "x"}

		// Act & Assert
		actual := args.Map{"result": empty.IsEmptyText()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": valid.IsEmptyText()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C79_TWLN_IsEmptyText_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmptyText_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		actual := args.Map{"result": twln.IsEmptyText()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for nil", actual)
	})
}

func Test_C79_TWLN_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmptyTextLineBoth", func() {
		// Arrange
		empty := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act & Assert
		actual := args.Map{"result": empty.IsEmptyTextLineBoth()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// --- CloneSlice ---

func Test_C79_CloneSlice(t *testing.T) {
	safeTest(t, "Test_C79_CloneSlice", func() {
		// Arrange
		input := []string{"a", "b"}

		// Act
		result := corestr.CloneSlice(input)

		// Assert
		actual := args.Map{"result": len(result) != 2 || result[0] != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone mismatch", actual)
	})
}

func Test_C79_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C79_CloneSlice_Empty", func() {
		// Arrange & Act
		result := corestr.CloneSlice([]string{})

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// --- CloneSliceIf ---

func Test_C79_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_C79_CloneSliceIf_Clone", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(true, "a", "b")

		// Assert
		actual := args.Map{"result": len(result) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C79_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_C79_CloneSliceIf_NoClone", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(false, "a")

		// Assert
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C79_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_C79_CloneSliceIf_Empty", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(true)

		// Assert
		actual := args.Map{"result": len(result) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// --- AnyToString ---

func Test_C79_AnyToString_WithFieldNames(t *testing.T) {
	safeTest(t, "Test_C79_AnyToString_WithFieldNames", func() {
		// Arrange
		type testStruct struct{ Name string }

		// Act
		result := corestr.AnyToString(true, testStruct{Name: "test"})

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C79_AnyToString_WithoutFieldNames(t *testing.T) {
	safeTest(t, "Test_C79_AnyToString_WithoutFieldNames", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, 42)

		// Assert
		actual := args.Map{"result": result != "42"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '42', got ''", actual)
	})
}

func Test_C79_AnyToString_EmptyString(t *testing.T) {
	safeTest(t, "Test_C79_AnyToString_EmptyString", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "")

		// Assert
		actual := args.Map{"result": result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// --- utils (StringUtils) ---

func Test_C79_Utils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDouble", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDouble("x")

		// Assert
		actual := args.Map{"result": result != `"x"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '\"x\"', got ''", actual)
	})
}

func Test_C79_Utils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingle", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingle("x")

		// Assert
		actual := args.Map{"result": result != "'x'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected \"'x'\"", actual)
	})
}

func Test_C79_Utils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapTilda", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapTilda("x")

		// Assert
		actual := args.Map{"result": result != "`x`"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected \"`x`\"", actual)
	})
}

func Test_C79_Utils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing(`"hello"`)

		// Assert
		actual := args.Map{"result": result != `"hello"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no change, got ''", actual)
	})
}

func Test_C79_Utils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDoubleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("hello")

		// Assert
		actual := args.Map{"result": result != `"hello"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '\"hello\"', got ''", actual)
	})
}

func Test_C79_Utils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDoubleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("")

		// Assert
		actual := args.Map{"result": result != `""`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '\"\"', got ''", actual)
	})
}

func Test_C79_Utils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("'hello'")

		// Assert
		actual := args.Map{"result": result != "'hello'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no change, got ''", actual)
	})
}

func Test_C79_Utils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("hello")

		// Assert
		actual := args.Map{"result": result != "'hello'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected \"'hello'\"", actual)
	})
}

func Test_C79_Utils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("")

		// Assert
		actual := args.Map{"result": result != "''"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected \"''\"", actual)
	})
}
