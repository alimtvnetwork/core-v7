package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
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
		if vv.Value != "hello" || !vv.IsValid {
			t.Error("NewValidValue mismatch")
		}
	})
}

func Test_C79_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValueEmpty", func() {
		// Arrange & Act
		vv := corestr.NewValidValueEmpty()

		// Assert
		if vv.Value != "" || !vv.IsValid {
			t.Error("expected empty valid value")
		}
	})
}

func Test_C79_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_InvalidValidValue", func() {
		// Arrange & Act
		vv := corestr.InvalidValidValue("err msg")

		// Assert
		if vv.IsValid || vv.Message != "err msg" {
			t.Error("expected invalid with message")
		}
	})
}

func Test_C79_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_InvalidValidValueNoMessage", func() {
		// Arrange & Act
		vv := corestr.InvalidValidValueNoMessage()

		// Assert
		if vv.IsValid {
			t.Error("expected invalid")
		}
	})
}

func Test_C79_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValueUsingAny", func() {
		// Arrange & Act
		vv := corestr.NewValidValueUsingAny(false, true, 42)

		// Assert
		if vv.Value != "42" || !vv.IsValid {
			t.Errorf("expected '42' valid, got '%s' %v", vv.Value, vv.IsValid)
		}
	})
}

func Test_C79_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_NewValidValueUsingAnyAutoValid", func() {
		// Arrange & Act
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")

		// Assert
		if vv.Value == "" {
			t.Error("expected non-empty value")
		}
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
		if string(bytes1) != "abc" || string(bytes2) != "abc" {
			t.Error("bytes mismatch")
		}
	})
}
func Test_C79_ValidValue_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsEmpty_IsWhitespace", func() {
		// Arrange
		empty := corestr.NewValidValue("")
		ws := corestr.NewValidValue("  ")
		val := corestr.NewValidValue("x")

		// Act & Assert
		if !empty.IsEmpty() {
			t.Error("expected empty")
		}
		if !ws.IsWhitespace() {
			t.Error("expected whitespace")
		}
		if val.IsEmpty() {
			t.Error("expected not empty")
		}
	})
}

func Test_C79_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue(" hello ")

		// Act & Assert
		if vv.Trim() != "hello" {
			t.Error("trim mismatch")
		}
	})
}

func Test_C79_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		valid := corestr.NewValidValue("x")
		invalid := corestr.InvalidValidValue("err")

		// Act & Assert
		if !valid.HasValidNonEmpty() {
			t.Error("expected true")
		}
		if invalid.HasValidNonEmpty() {
			t.Error("expected false")
		}
	})
}

func Test_C79_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		valid := corestr.NewValidValue("x")
		ws := corestr.NewValidValue("  ")

		// Act & Assert
		if !valid.HasValidNonWhitespace() {
			t.Error("expected true")
		}
		if ws.HasValidNonWhitespace() {
			t.Error("expected false for whitespace")
		}
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
		if !trueVV.ValueBool() {
			t.Error("expected true")
		}
		if falseVV.ValueBool() {
			t.Error("expected false")
		}
		if invalidVV.ValueBool() {
			t.Error("expected false for invalid")
		}
		if emptyVV.ValueBool() {
			t.Error("expected false for empty")
		}
	})
}

func Test_C79_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueInt", func() {
		// Arrange
		vv := corestr.NewValidValue("42")
		invalid := corestr.NewValidValue("abc")

		// Act & Assert
		if vv.ValueInt(0) != 42 {
			t.Error("expected 42")
		}
		if invalid.ValueInt(99) != 99 {
			t.Error("expected 99")
		}
	})
}

func Test_C79_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueDefInt", func() {
		// Arrange
		vv := corestr.NewValidValue("10")
		invalid := corestr.NewValidValue("x")

		// Act & Assert
		if vv.ValueDefInt() != 10 {
			t.Error("expected 10")
		}
		if invalid.ValueDefInt() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C79_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueByte", func() {
		// Arrange
		valid := corestr.NewValidValue("100")
		overflow := corestr.NewValidValue("300")
		negative := corestr.NewValidValue("-5")
		invalid := corestr.NewValidValue("abc")

		// Act & Assert
		if valid.ValueByte(0) != 100 {
			t.Error("expected 100")
		}
		if overflow.ValueByte(0) != 255 { // MaxUnit8
			t.Errorf("expected 255 for overflow, got %d", overflow.ValueByte(0))
		}
		if negative.ValueByte(0) != 0 {
			t.Error("expected 0 for negative")
		}
		if invalid.ValueByte(0) != 0 {
			t.Error("expected 0 for invalid")
		}
	})
}

func Test_C79_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueDefByte", func() {
		// Arrange
		valid := corestr.NewValidValue("50")
		overflow := corestr.NewValidValue("999")
		negative := corestr.NewValidValue("-1")

		// Act & Assert
		if valid.ValueDefByte() != 50 {
			t.Error("expected 50")
		}
		if overflow.ValueDefByte() != 255 {
			t.Errorf("expected 255, got %d", overflow.ValueDefByte())
		}
		if negative.ValueDefByte() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C79_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")
		invalid := corestr.NewValidValue("abc")

		// Act & Assert
		if vv.ValueFloat64(0) != 3.14 {
			t.Error("expected 3.14")
		}
		if invalid.ValueFloat64(1.5) != 1.5 {
			t.Error("expected 1.5")
		}
	})
}

func Test_C79_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_ValueDefFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("2.5")

		// Act & Assert
		if vv.ValueDefFloat64() != 2.5 {
			t.Error("expected 2.5")
		}
	})
}

func Test_C79_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		valid := corestr.NewValidValue("x")
		empty := corestr.NewValidValue("")

		// Act & Assert
		if !valid.HasSafeNonEmpty() {
			t.Error("expected true")
		}
		if empty.HasSafeNonEmpty() {
			t.Error("expected false")
		}
	})
}

func Test_C79_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Is", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act & Assert
		if !vv.Is("hello") {
			t.Error("expected true")
		}
		if vv.Is("world") {
			t.Error("expected false")
		}
	})
}

func Test_C79_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act & Assert
		if !vv.IsAnyOf("a", "b") {
			t.Error("expected true")
		}
		if vv.IsAnyOf("x", "y") {
			t.Error("expected false")
		}
		if !vv.IsAnyOf() {
			t.Error("expected true for empty")
		}
	})
}

func Test_C79_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act & Assert
		if !vv.IsContains("world") {
			t.Error("expected true")
		}
	})
}

func Test_C79_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act & Assert
		if !vv.IsAnyContains("xyz", "world") {
			t.Error("expected true")
		}
		if vv.IsAnyContains("abc") {
			t.Error("expected false")
		}
		if !vv.IsAnyContains() {
			t.Error("expected true for empty")
		}
	})
}

func Test_C79_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act & Assert
		if !vv.IsEqualNonSensitive("hello") {
			t.Error("expected true")
		}
	})
}

func Test_C79_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_IsRegexMatches", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		if !vv.IsRegexMatches(re) {
			t.Error("expected true")
		}
		if vv.IsRegexMatches(nil) {
			t.Error("expected false for nil")
		}
	})
}

func Test_C79_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindString", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		if vv.RegexFindString(re) != "123" {
			t.Error("expected '123'")
		}
		if vv.RegexFindString(nil) != "" {
			t.Error("expected empty for nil")
		}
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
		if !hasAny || len(items) != 3 {
			t.Error("expected 3 matches")
		}
	})
}

func Test_C79_ValidValue_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act
		items, hasAny := vv.RegexFindAllStringsWithFlag(nil, -1)

		// Assert
		if hasAny || len(items) != 0 {
			t.Error("expected empty for nil")
		}
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
		if len(items) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C79_ValidValue_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_RegexFindAllStrings_Nil", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")

		// Act & Assert
		if len(vv.RegexFindAllStrings(nil, -1)) != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C79_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")

		// Act
		result := vv.Split(",")

		// Assert
		if len(result) != 3 {
			t.Errorf("expected 3, got %d", len(result))
		}
	})
}

func Test_C79_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_SplitNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("a::b")

		// Act
		result := vv.SplitNonEmpty("::")

		// Assert
		if len(result) < 2 {
			t.Errorf("expected at least 2, got %d", len(result))
		}
	})
}

func Test_C79_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_SplitTrimNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("a , , b")

		// Act
		result := vv.SplitTrimNonWhitespace(",")

		// Assert
		if len(result) < 2 {
			t.Errorf("expected at least 2, got %d", len(result))
		}
	})
}

func Test_C79_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		cloned := vv.Clone()

		// Assert
		if cloned == nil || cloned.Value != "x" {
			t.Error("clone mismatch")
		}
	})
}

func Test_C79_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Clone_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		cloned := vv.Clone()

		// Assert
		if cloned != nil {
			t.Error("expected nil")
		}
	})
}

func Test_C79_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_String", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act & Assert
		if vv.String() != "hello" {
			t.Error("String mismatch")
		}
	})
}

func Test_C79_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_String_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act & Assert
		if vv.String() != "" {
			t.Error("expected empty for nil")
		}
	})
}

func Test_C79_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		result := vv.FullString()

		// Assert
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C79_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_FullString_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act & Assert
		if vv.FullString() != "" {
			t.Error("expected empty for nil")
		}
	})
}

func Test_C79_ValidValue_Clear_Dispose(t *testing.T) {
	safeTest(t, "Test_C79_ValidValue_Clear_Dispose", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		vv.Clear()

		// Assert
		if vv.Value != "" || vv.IsValid {
			t.Error("expected cleared")
		}
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
		if jsonResult.HasError() {
			t.Error("json error")
		}
		if err != nil || len(bytes) == 0 {
			t.Error("serialize error")
		}
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
		if err != nil {
			t.Errorf("error: %v", err)
		}
		if result.Value != "hello" {
			t.Error("value mismatch")
		}
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
		if err != nil {
			t.Errorf("error: %v", err)
		}
	})
}

// --- ValidValues ---

func Test_C79_ValidValues_NewValidValues(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_NewValidValues", func() {
		// Arrange & Act
		vvs := corestr.NewValidValues(5)

		// Assert
		if vvs == nil || vvs.Length() != 0 {
			t.Error("expected empty")
		}
	})
}

func Test_C79_ValidValues_EmptyValidValues(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_EmptyValidValues", func() {
		// Arrange & Act
		vvs := corestr.EmptyValidValues()

		// Assert
		if vvs == nil || !vvs.IsEmpty() {
			t.Error("expected empty")
		}
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
		if vvs.Length() != 2 {
			t.Errorf("expected 2, got %d", vvs.Length())
		}
	})
}

func Test_C79_ValidValues_NewValidValuesUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_NewValidValuesUsingValues_Empty", func() {
		// Arrange & Act
		vvs := corestr.NewValidValuesUsingValues()

		// Assert
		if vvs.Length() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C79_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Add", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.Add("a").Add("b")

		// Assert
		if vvs.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_C79_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddFull", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddFull(false, "err", "msg")

		// Assert
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C79_ValidValues_Count_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Count_HasAnyItem", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")

		// Act & Assert
		if vvs.Count() != 1 {
			t.Error("expected 1")
		}
		if !vvs.HasAnyItem() {
			t.Error("expected true")
		}
	})
}

func Test_C79_ValidValues_LastIndex_HasIndex(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_LastIndex_HasIndex", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")

		// Act & Assert
		if vvs.LastIndex() != 1 {
			t.Error("expected 1")
		}
		if !vvs.HasIndex(0) {
			t.Error("expected true for 0")
		}
		if vvs.HasIndex(5) {
			t.Error("expected false for 5")
		}
	})
}

func Test_C79_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_SafeValueAt", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a").Add("b")

		// Act & Assert
		if vvs.SafeValueAt(0) != "a" {
			t.Error("expected 'a'")
		}
		if vvs.SafeValueAt(5) != "" {
			t.Error("expected empty for out of range")
		}
	})
}

func Test_C79_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_SafeValidValueAt", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()
		vvs.Add("a")
		vvs.AddFull(false, "b", "err")

		// Act & Assert
		if vvs.SafeValidValueAt(0) != "a" {
			t.Error("expected 'a'")
		}
		if vvs.SafeValidValueAt(1) != "" {
			t.Error("expected empty for invalid value")
		}
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
		if len(result) != 2 || result[0] != "a" || result[1] != "c" {
			t.Errorf("expected [a, c], got %v", result)
		}
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
		if len(result) != 2 || result[0] != "a" || result[1] != "" {
			t.Errorf("unexpected result: %v", result)
		}
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
		if len(result) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C79_ValidValues_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Strings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act & Assert
		if len(vvs.Strings()) != 0 {
			t.Error("expected 0")
		}
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
		if len(result) != 1 || result[0] == "" {
			t.Error("expected non-empty full string")
		}
	})
}

func Test_C79_ValidValues_FullStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_FullStrings_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act & Assert
		if len(vvs.FullStrings()) != 0 {
			t.Error("expected 0")
		}
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
		if result == "" {
			t.Error("expected non-empty")
		}
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
		if len(found) != 1 {
			t.Errorf("expected 1, got %d", len(found))
		}
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
		if len(found) != 1 {
			t.Errorf("expected 1 due to break, got %d", len(found))
		}
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
		if len(found) != 0 {
			t.Error("expected 0")
		}
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
		if vvs.Length() != 2 {
			t.Errorf("expected 2")
		}
	})
}

func Test_C79_ValidValues_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_Adds_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.Adds()

		// Assert
		if vvs.Length() != 0 {
			t.Error("expected 0")
		}
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
		if vvs.Length() != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C79_ValidValues_AddsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddsPtr_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddsPtr()

		// Assert
		if vvs.Length() != 0 {
			t.Error("expected 0")
		}
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
		if vvs1.Length() != 2 {
			t.Errorf("expected 2, got %d", vvs1.Length())
		}
	})
}

func Test_C79_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddValidValues_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddValidValues(nil)

		// Assert
		if vvs.Length() != 0 {
			t.Error("expected 0")
		}
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
		if result.Length() != 2 {
			t.Errorf("expected 2, got %d", result.Length())
		}
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
		if result.Length() != 1 {
			t.Errorf("expected 1, got %d", result.Length())
		}
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
		if result.Length() != 1 {
			t.Error("expected same instance")
		}
	})
}

func Test_C79_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashsetMap", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddHashsetMap(map[string]bool{"x": true, "y": false})

		// Assert
		if vvs.Length() != 2 {
			t.Errorf("expected 2, got %d", vvs.Length())
		}
	})
}

func Test_C79_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashsetMap_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddHashsetMap(nil)

		// Assert
		if vvs.Length() != 0 {
			t.Error("expected 0")
		}
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
		if vvs.Length() != 2 {
			t.Errorf("expected 2, got %d", vvs.Length())
		}
	})
}

func Test_C79_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C79_ValidValues_AddHashset_Nil", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		vvs.AddHashset(nil)

		// Assert
		if vvs.Length() != 0 {
			t.Error("expected 0")
		}
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
		if hm.Length() != 1 {
			t.Error("expected 1")
		}
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
		if len(m) != 1 {
			t.Error("expected 1")
		}
	})
}

// --- ValueStatus ---

func Test_C79_ValueStatus_InvalidValueStatusNoMessage(t *testing.T) {
	safeTest(t, "Test_C79_ValueStatus_InvalidValueStatusNoMessage", func() {
		// Arrange & Act
		vs := corestr.InvalidValueStatusNoMessage()

		// Assert
		if vs.ValueValid.IsValid {
			t.Error("expected invalid")
		}
		if vs.Index != -1 {
			t.Errorf("expected -1, got %d", vs.Index)
		}
	})
}

func Test_C79_ValueStatus_InvalidValueStatus(t *testing.T) {
	safeTest(t, "Test_C79_ValueStatus_InvalidValueStatus", func() {
		// Arrange & Act
		vs := corestr.InvalidValueStatus("err")

		// Assert
		if vs.ValueValid.IsValid {
			t.Error("expected invalid")
		}
		if vs.ValueValid.Message != "err" {
			t.Error("message mismatch")
		}
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
		if cloned.Index != 5 || cloned.ValueValid.Value != "x" {
			t.Error("clone mismatch")
		}
	})
}

// --- TextWithLineNumber ---

func Test_C79_TWLN_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_HasLineNumber", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		invalid := &corestr.TextWithLineNumber{LineNumber: -1, Text: "x"}

		// Act & Assert
		if !twln.HasLineNumber() {
			t.Error("expected true")
		}
		if invalid.HasLineNumber() {
			t.Error("expected false for -1")
		}
	})
}

func Test_C79_TWLN_HasLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_HasLineNumber_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		if twln.HasLineNumber() {
			t.Error("expected false for nil")
		}
	})
}

func Test_C79_TWLN_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsInvalidLineNumber", func() {
		// Arrange
		valid := &corestr.TextWithLineNumber{LineNumber: 1, Text: "x"}
		invalid := &corestr.TextWithLineNumber{LineNumber: -1, Text: "x"}

		// Act & Assert
		if valid.IsInvalidLineNumber() {
			t.Error("expected false")
		}
		if !invalid.IsInvalidLineNumber() {
			t.Error("expected true")
		}
	})
}

func Test_C79_TWLN_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsInvalidLineNumber_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		if !twln.IsInvalidLineNumber() {
			t.Error("expected true for nil")
		}
	})
}

func Test_C79_TWLN_Length(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_Length", func() {
		// Arrange
		twln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act & Assert
		if twln.Length() != 5 {
			t.Errorf("expected 5, got %d", twln.Length())
		}
	})
}

func Test_C79_TWLN_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_Length_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		if twln.Length() != 0 {
			t.Error("expected 0 for nil")
		}
	})
}

func Test_C79_TWLN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmpty", func() {
		// Arrange
		empty := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		valid := &corestr.TextWithLineNumber{LineNumber: 1, Text: "x"}

		// Act & Assert
		if !empty.IsEmpty() {
			t.Error("expected empty")
		}
		if valid.IsEmpty() {
			t.Error("expected not empty")
		}
	})
}

func Test_C79_TWLN_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmpty_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		if !twln.IsEmpty() {
			t.Error("expected true for nil")
		}
	})
}

func Test_C79_TWLN_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmptyText", func() {
		// Arrange
		empty := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		valid := &corestr.TextWithLineNumber{LineNumber: 1, Text: "x"}

		// Act & Assert
		if !empty.IsEmptyText() {
			t.Error("expected true")
		}
		if valid.IsEmptyText() {
			t.Error("expected false")
		}
	})
}

func Test_C79_TWLN_IsEmptyText_Nil(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmptyText_Nil", func() {
		// Arrange
		var twln *corestr.TextWithLineNumber

		// Act & Assert
		if !twln.IsEmptyText() {
			t.Error("expected true for nil")
		}
	})
}

func Test_C79_TWLN_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C79_TWLN_IsEmptyTextLineBoth", func() {
		// Arrange
		empty := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act & Assert
		if !empty.IsEmptyTextLineBoth() {
			t.Error("expected true")
		}
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
		if len(result) != 2 || result[0] != "a" {
			t.Error("clone mismatch")
		}
	})
}

func Test_C79_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C79_CloneSlice_Empty", func() {
		// Arrange & Act
		result := corestr.CloneSlice([]string{})

		// Assert
		if len(result) != 0 {
			t.Error("expected empty")
		}
	})
}

// --- CloneSliceIf ---

func Test_C79_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_C79_CloneSliceIf_Clone", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(true, "a", "b")

		// Assert
		if len(result) != 2 {
			t.Error("expected 2")
		}
	})
}

func Test_C79_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_C79_CloneSliceIf_NoClone", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(false, "a")

		// Assert
		if len(result) != 1 {
			t.Error("expected 1")
		}
	})
}

func Test_C79_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_C79_CloneSliceIf_Empty", func() {
		// Arrange & Act
		result := corestr.CloneSliceIf(true)

		// Assert
		if len(result) != 0 {
			t.Error("expected 0")
		}
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
		if result == "" {
			t.Error("expected non-empty")
		}
	})
}

func Test_C79_AnyToString_WithoutFieldNames(t *testing.T) {
	safeTest(t, "Test_C79_AnyToString_WithoutFieldNames", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, 42)

		// Assert
		if result != "42" {
			t.Errorf("expected '42', got '%s'", result)
		}
	})
}

func Test_C79_AnyToString_EmptyString(t *testing.T) {
	safeTest(t, "Test_C79_AnyToString_EmptyString", func() {
		// Arrange & Act
		result := corestr.AnyToString(false, "")

		// Assert
		if result != "" {
			t.Error("expected empty")
		}
	})
}

// --- utils (StringUtils) ---

func Test_C79_Utils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDouble", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDouble("x")

		// Assert
		if result != `"x"` {
			t.Errorf("expected '\"x\"', got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingle", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingle("x")

		// Assert
		if result != "'x'" {
			t.Errorf("expected \"'x'\", got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapTilda", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapTilda("x")

		// Assert
		if result != "`x`" {
			t.Errorf("expected \"`x`\", got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapDoubleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDoubleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing(`"hello"`)

		// Assert
		if result != `"hello"` {
			t.Errorf("expected no change, got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapDoubleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDoubleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("hello")

		// Assert
		if result != `"hello"` {
			t.Errorf("expected '\"hello\"', got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapDoubleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapDoubleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapDoubleIfMissing("")

		// Assert
		if result != `""` {
			t.Errorf("expected '\"\"', got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapSingleIfMissing_AlreadyWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingleIfMissing_AlreadyWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("'hello'")

		// Assert
		if result != "'hello'" {
			t.Errorf("expected no change, got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapSingleIfMissing_NotWrapped(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingleIfMissing_NotWrapped", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("hello")

		// Assert
		if result != "'hello'" {
			t.Errorf("expected \"'hello'\", got '%s'", result)
		}
	})
}

func Test_C79_Utils_WrapSingleIfMissing_Empty(t *testing.T) {
	safeTest(t, "Test_C79_Utils_WrapSingleIfMissing_Empty", func() {
		// Arrange & Act
		result := corestr.StringUtils.WrapSingleIfMissing("")

		// Assert
		if result != "''" {
			t.Errorf("expected \"''\", got '%s'", result)
		}
	})
}
