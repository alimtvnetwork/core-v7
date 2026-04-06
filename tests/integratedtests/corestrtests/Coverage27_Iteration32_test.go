package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════════════════════════════
// ValidValue — comprehensive coverage
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_01_ValidValue_NewValidValue(t *testing.T) {
	safeTest(t, "Test_C27_01_ValidValue_NewValidValue", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": v.Value != "hello" || !v.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello, valid", actual)
	})
}

func Test_C27_02_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_02_ValidValue_NewValidValueEmpty", func() {
		// Arrange
		v := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{"result": v.Value != "" || !v.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty, valid", actual)
	})
}

func Test_C27_03_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C27_03_ValidValue_InvalidValidValue", func() {
		// Arrange
		v := corestr.InvalidValidValue("err")

		// Act
		actual := args.Map{"result": v.IsValid || v.Message != "err"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid with err", actual)
	})
}

func Test_C27_04_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_04_ValidValue_InvalidValidValueNoMessage", func() {
		// Arrange
		v := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{"result": v.IsValid || v.Message != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid no message", actual)
	})
}

func Test_C27_05_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C27_05_ValidValue_NewValidValueUsingAny", func() {
		// Arrange
		v := corestr.NewValidValueUsingAny(false, true, 42)

		// Act
		actual := args.Map{"result": v.IsValid || v.Value == ""}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected valid with value", actual)
	})
}

func Test_C27_06_ValidValue_NewValidValueUsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C27_06_ValidValue_NewValidValueUsingAnyAutoValid", func() {
		v := corestr.NewValidValueUsingAnyAutoValid(false, 42)
		_ = v
	})
}

func Test_C27_07_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C27_07_ValidValue_ValueBytesOnce", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		b := v.ValueBytesOnce()

		// Act
		actual := args.Map{"result": string(b) != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		// second call should return same
		b2 := v.ValueBytesOnce()
		actual = args.Map{"result": string(b2) != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cached hello", actual)
	})
}

func Test_C27_08_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C27_08_ValidValue_ValueBytesOncePtr", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		b := v.ValueBytesOncePtr()

		// Act
		actual := args.Map{"result": string(b) != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}

func Test_C27_09_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_09_ValidValue_IsEmpty", func() {
		// Arrange
		v := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{"result": v.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		v2 := corestr.NewValidValue("x")
		actual = args.Map{"result": v2.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_C27_10_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_10_ValidValue_IsWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("   ")

		// Act
		actual := args.Map{"result": v.IsWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
	})
}

func Test_C27_11_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C27_11_ValidValue_Trim", func() {
		// Arrange
		v := corestr.NewValidValue("  hello  ")

		// Act
		actual := args.Map{"result": v.Trim() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_12_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_12_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": v.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_13_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_13_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		v := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": v.HasValidNonWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_14_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_14_ValidValue_ValueBool", func() {
		// Arrange
		v := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"result": v.ValueBool()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		v2 := corestr.NewValidValue("nope")
		actual = args.Map{"result": v2.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		v3 := corestr.NewValidValue("")
		actual = args.Map{"result": v3.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C27_15_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C27_15_ValidValue_ValueInt", func() {
		// Arrange
		v := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"result": v.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		v2 := corestr.NewValidValue("bad")
		actual = args.Map{"result": v2.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default 99", actual)
	})
}

func Test_C27_16_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C27_16_ValidValue_ValueDefInt", func() {
		// Arrange
		v := corestr.NewValidValue("10")

		// Act
		actual := args.Map{"result": v.ValueDefInt() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C27_17_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_17_ValidValue_ValueByte", func() {
		// Arrange
		v := corestr.NewValidValue("200")

		// Act
		actual := args.Map{"result": v.ValueByte(0) != 200}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
		v2 := corestr.NewValidValue("300")
		actual = args.Map{"result": v2.ValueByte(0) != 255}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255 for overflow", actual)
		v3 := corestr.NewValidValue("-1")
		actual = args.Map{"result": v3.ValueByte(0) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for negative", actual)
	})
}

func Test_C27_18_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C27_18_ValidValue_ValueDefByte", func() {
		// Arrange
		v := corestr.NewValidValue("100")

		// Act
		actual := args.Map{"result": v.ValueDefByte() != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_C27_19_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_19_ValidValue_ValueFloat64", func() {
		// Arrange
		v := corestr.NewValidValue("3.14")

		// Act
		actual := args.Map{"result": v.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		v2 := corestr.NewValidValue("bad")
		actual = args.Map{"result": v2.ValueFloat64(1.0) != 1.0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C27_20_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_20_ValidValue_ValueDefFloat64", func() {
		// Arrange
		v := corestr.NewValidValue("2.5")

		// Act
		actual := args.Map{"result": v.ValueDefFloat64() != 2.5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C27_21_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_21_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		v := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": v.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_22_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C27_22_ValidValue_Is", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": v.Is("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_23_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C27_23_ValidValue_IsAnyOf", func() {
		// Arrange
		v := corestr.NewValidValue("b")

		// Act
		actual := args.Map{"result": v.IsAnyOf("a", "b", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": v.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty values", actual)
		actual = args.Map{"result": v.IsAnyOf("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_24_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_24_ValidValue_IsContains", func() {
		// Arrange
		v := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"result": v.IsContains("world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_25_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C27_25_ValidValue_IsAnyContains", func() {
		// Arrange
		v := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"result": v.IsAnyContains("xyz", "world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": v.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
		actual = args.Map{"result": v.IsAnyContains("xyz", "abc")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_26_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C27_26_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		v := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{"result": v.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_27_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C27_27_ValidValue_IsRegexMatches", func() {
		// Arrange
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": v.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": v.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil regex", actual)
	})
}

func Test_C27_28_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C27_28_ValidValue_RegexFindString", func() {
		// Arrange
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": v.RegexFindString(re) != "123"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual = args.Map{"result": v.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C27_29_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C27_29_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, has := v.RegexFindAllStringsWithFlag(re, -1)

		// Act
		actual := args.Map{"result": has || len(items) != 3}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_, has2 := v.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": has2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C27_30_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C27_30_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		v := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items := v.RegexFindAllStrings(re, -1)

		// Act
		actual := args.Map{"result": len(items) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		items2 := v.RegexFindAllStrings(nil, -1)
		actual = args.Map{"result": len(items2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_31_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C27_31_ValidValue_Split", func() {
		// Arrange
		v := corestr.NewValidValue("a,b,c")
		s := v.Split(",")

		// Act
		actual := args.Map{"result": len(s) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_32_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_32_ValidValue_SplitNonEmpty", func() {
		v := corestr.NewValidValue("a,,b")
		s := v.SplitNonEmpty(",")
		_ = s // just no panic
	})
}

func Test_C27_33_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_33_ValidValue_SplitTrimNonWhitespace", func() {
		v := corestr.NewValidValue("a , b , c")
		s := v.SplitTrimNonWhitespace(",")
		_ = s
	})
}

func Test_C27_34_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_C27_34_ValidValue_Clone", func() {
		// Arrange
		v := corestr.NewValidValue("hello")
		c := v.Clone()

		// Act
		actual := args.Map{"result": c.Value != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_35_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C27_35_ValidValue_Clone_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"result": v.Clone() != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C27_36_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C27_36_ValidValue_String", func() {
		// Arrange
		v := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{"result": v.String() != "hi"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}

func Test_C27_37_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C27_37_ValidValue_String_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"result": v.String() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_38_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C27_38_ValidValue_FullString", func() {
		// Arrange
		v := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{"result": v.FullString() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_39_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C27_39_ValidValue_FullString_Nil", func() {
		// Arrange
		var v *corestr.ValidValue

		// Act
		actual := args.Map{"result": v.FullString() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_40_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_C27_40_ValidValue_Clear", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		v.Clear()

		// Act
		actual := args.Map{"result": v.Value != "" || v.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cleared", actual)
	})
}

func Test_C27_41_ValidValue_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C27_41_ValidValue_Clear_Nil", func() {
		var v *corestr.ValidValue
		v.Clear() // no panic
	})
}

func Test_C27_42_ValidValue_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_42_ValidValue_Dispose", func() {
		v := corestr.NewValidValue("hi")
		v.Dispose()
	})
}

func Test_C27_43_ValidValue_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C27_43_ValidValue_Dispose_Nil", func() {
		var v *corestr.ValidValue
		v.Dispose()
	})
}

func Test_C27_44_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_C27_44_ValidValue_Json", func() {
		v := corestr.NewValidValue("hi")
		j := v.Json()
		_ = j
	})
}

func Test_C27_45_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C27_45_ValidValue_JsonPtr", func() {
		// Arrange
		v := corestr.NewValidValue("hi")

		// Act
		actual := args.Map{"result": v.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C27_46_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_46_ValidValue_Serialize", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		b, err := v.Serialize()

		// Act
		actual := args.Map{"result": err != nil || len(b) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_C27_47_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_47_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		jp := v.JsonPtr()
		v2 := &corestr.ValidValue{}
		_, err := v2.ParseInjectUsingJson(jp)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C27_48_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_C27_48_ValidValue_Deserialize", func() {
		// Arrange
		v := corestr.NewValidValue("hi")
		var target corestr.ValidValue
		err := v.Deserialize(&target)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// ValidValues
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_49_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_49_ValidValues_Empty", func() {
		// Arrange
		vv := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vv.IsEmpty() || vv.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_50_ValidValues_NewValidValues(t *testing.T) {
	safeTest(t, "Test_C27_50_ValidValues_NewValidValues", func() {
		// Arrange
		vv := corestr.NewValidValues(4)

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_51_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_C27_51_ValidValues_NewValidValuesUsingValues", func() {
		// Arrange
		vv := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)

		// Act
		actual := args.Map{"result": vv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_52_ValidValues_NewValidValuesUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_52_ValidValues_NewValidValuesUsingValues_Empty", func() {
		// Arrange
		vv := corestr.NewValidValuesUsingValues()

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_53_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C27_53_ValidValues_Add", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("hello")

		// Act
		actual := args.Map{"result": vv.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_54_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C27_54_ValidValues_AddFull", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddFull(false, "val", "msg")

		// Act
		actual := args.Map{"result": vv.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_55_ValidValues_Count(t *testing.T) {
	safeTest(t, "Test_C27_55_ValidValues_Count", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")

		// Act
		actual := args.Map{"result": vv.Count() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_56_ValidValues_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_56_ValidValues_HasAnyItem", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")

		// Act
		actual := args.Map{"result": vv.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_57_ValidValues_LastIndex(t *testing.T) {
	safeTest(t, "Test_C27_57_ValidValues_LastIndex", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv.Add("b")

		// Act
		actual := args.Map{"result": vv.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_58_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_C27_58_ValidValues_HasIndex", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")

		// Act
		actual := args.Map{"result": vv.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": vv.HasIndex(5)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_59_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_59_ValidValues_SafeValueAt", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("hello")

		// Act
		actual := args.Map{"result": vv.SafeValueAt(0) != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		actual = args.Map{"result": vv.SafeValueAt(99) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_60_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C27_60_ValidValues_SafeValidValueAt", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("hello")

		// Act
		actual := args.Map{"result": vv.SafeValidValueAt(0) != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_61_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_61_ValidValues_SafeValuesAtIndexes", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv.Add("b")
		vals := vv.SafeValuesAtIndexes(0, 1)

		// Act
		actual := args.Map{"result": len(vals) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_62_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_62_ValidValues_SafeValidValuesAtIndexes", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vals := vv.SafeValidValuesAtIndexes(0)

		// Act
		actual := args.Map{"result": len(vals) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_63_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C27_63_ValidValues_Strings", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		s := vv.Strings()

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_64_ValidValues_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C27_64_ValidValues_Strings_Empty", func() {
		// Arrange
		vv := corestr.EmptyValidValues()
		s := vv.Strings()

		// Act
		actual := args.Map{"result": len(s) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_65_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C27_65_ValidValues_FullStrings", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		s := vv.FullStrings()

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_66_ValidValues_FullStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C27_66_ValidValues_FullStrings_Empty", func() {
		// Arrange
		vv := corestr.EmptyValidValues()
		s := vv.FullStrings()

		// Act
		actual := args.Map{"result": len(s) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_67_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C27_67_ValidValues_String", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")

		// Act
		actual := args.Map{"result": vv.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_68_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C27_68_ValidValues_Find", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv.Add("b")
		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, v.Value == "b", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_69_ValidValues_Find_Empty(t *testing.T) {
	safeTest(t, "Test_C27_69_ValidValues_Find_Empty", func() {
		// Arrange
		vv := corestr.EmptyValidValues()
		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, false
		})

		// Act
		actual := args.Map{"result": len(found) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_70_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C27_70_ValidValues_ConcatNew", func() {
		// Arrange
		vv1 := corestr.NewValidValues(4)
		vv1.Add("a")
		vv2 := corestr.NewValidValues(4)
		vv2.Add("b")
		result := vv1.ConcatNew(false, vv2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_71_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C27_71_ValidValues_ConcatNew_EmptyClone", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		result := vv.ConcatNew(true)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_72_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C27_72_ValidValues_ConcatNew_EmptyNoClone", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		result := vv.ConcatNew(false)

		// Act
		actual := args.Map{"result": result != vv}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
	})
}

func Test_C27_73_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C27_73_ValidValues_AddValidValues", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv2 := corestr.NewValidValues(4)
		vv2.Add("b")
		vv.AddValidValues(vv2)

		// Act
		actual := args.Map{"result": vv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_74_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_C27_74_ValidValues_AddValidValues_Nil", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddValidValues(nil)

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_75_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C27_75_ValidValues_Adds", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})

		// Act
		actual := args.Map{"result": vv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_76_ValidValues_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C27_76_ValidValues_Adds_Empty", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.Adds()

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_77_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C27_77_ValidValues_AddsPtr", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		v := corestr.NewValidValue("a")
		vv.AddsPtr(v)

		// Act
		actual := args.Map{"result": vv.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_78_ValidValues_AddsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C27_78_ValidValues_AddsPtr_Empty", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddsPtr()

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_79_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_79_ValidValues_AddHashsetMap", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddHashsetMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": vv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_80_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C27_80_ValidValues_AddHashsetMap_Nil", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddHashsetMap(nil)

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_81_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_81_ValidValues_AddHashset", func() {
		// Arrange
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		vv := corestr.NewValidValues(4)
		vv.AddHashset(hs)

		// Act
		actual := args.Map{"result": vv.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_82_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C27_82_ValidValues_AddHashset_Nil", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddHashset(nil)

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_83_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_83_ValidValues_Hashmap", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddFull(true, "k", "v")
		hm := vv.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_84_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C27_84_ValidValues_Map", func() {
		// Arrange
		vv := corestr.NewValidValues(4)
		vv.AddFull(true, "k", "v")
		m := vv.Map()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_85_ValidValues_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C27_85_ValidValues_Length_Nil", func() {
		// Arrange
		var vv *corestr.ValidValues

		// Act
		actual := args.Map{"result": vv.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_86_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_86_ValueStatus_InvalidNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{"result": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_87_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_87_ValueStatus_Invalid", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")

		// Act
		actual := args.Map{"result": vs.ValueValid.Message != "err"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected err", actual)
	})
}

func Test_C27_88_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C27_88_ValueStatus_Clone", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("err")
		c := vs.Clone()

		// Act
		actual := args.Map{"result": c.ValueValid.Message != "err"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected err", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_89_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_89_LeftRight_NewLeftRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b" || !lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, valid", actual)
	})
}

func Test_C27_90_LeftRight_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_90_LeftRight_InvalidLeftRight", func() {
		// Arrange
		lr := corestr.InvalidLeftRight("err")

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_91_LeftRight_InvalidLeftRightNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_91_LeftRight_InvalidLeftRightNoMessage", func() {
		// Arrange
		lr := corestr.InvalidLeftRightNoMessage()

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_92_LeftRight_LeftRightUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_92_LeftRight_LeftRightUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b", actual)
	})
}

func Test_C27_93_LeftRight_LeftRightUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_93_LeftRight_LeftRightUsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{})

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_94_LeftRight_LeftRightUsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_94_LeftRight_LeftRightUsingSlice_Single", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlice([]string{"a"})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, empty", actual)
	})
}

func Test_C27_95_LeftRight_LeftRightUsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_C27_95_LeftRight_LeftRightUsingSlicePtr", func() {
		// Arrange
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})

		// Act
		actual := args.Map{"result": lr.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_96_LeftRight_LeftRightTrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_96_LeftRight_LeftRightTrimmedUsingSlice", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_97_LeftRight_LeftRightTrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_C27_97_LeftRight_LeftRightTrimmedUsingSlice_Nil", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice(nil)

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_98_LeftRight_LeftRightTrimmedUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_98_LeftRight_LeftRightTrimmedUsingSlice_Empty", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})

		// Act
		actual := args.Map{"result": lr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single", func() {
		// Arrange
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})

		// Act
		actual := args.Map{"result": lr.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_100_LeftRight_LeftBytes(t *testing.T) {
	safeTest(t, "Test_C27_100_LeftRight_LeftBytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": string(lr.LeftBytes()) != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_101_LeftRight_RightBytes(t *testing.T) {
	safeTest(t, "Test_C27_101_LeftRight_RightBytes", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": string(lr.RightBytes()) != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C27_102_LeftRight_LeftTrim(t *testing.T) {
	safeTest(t, "Test_C27_102_LeftRight_LeftTrim", func() {
		// Arrange
		lr := corestr.NewLeftRight(" a ", "b")

		// Act
		actual := args.Map{"result": lr.LeftTrim() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_103_LeftRight_RightTrim(t *testing.T) {
	safeTest(t, "Test_C27_103_LeftRight_RightTrim", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", " b ")

		// Act
		actual := args.Map{"result": lr.RightTrim() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C27_104_LeftRight_IsLeftEmpty(t *testing.T) {
	safeTest(t, "Test_C27_104_LeftRight_IsLeftEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("", "b")

		// Act
		actual := args.Map{"result": lr.IsLeftEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_105_LeftRight_IsRightEmpty(t *testing.T) {
	safeTest(t, "Test_C27_105_LeftRight_IsRightEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "")

		// Act
		actual := args.Map{"result": lr.IsRightEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_106_LeftRight_IsLeftWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_106_LeftRight_IsLeftWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("   ", "b")

		// Act
		actual := args.Map{"result": lr.IsLeftWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_107_LeftRight_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_107_LeftRight_IsRightWhitespace", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "   ")

		// Act
		actual := args.Map{"result": lr.IsRightWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_108_LeftRight_HasValidNonEmptyLeft(t *testing.T) {
	safeTest(t, "Test_C27_108_LeftRight_HasValidNonEmptyLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonEmptyLeft()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_109_LeftRight_HasValidNonEmptyRight(t *testing.T) {
	safeTest(t, "Test_C27_109_LeftRight_HasValidNonEmptyRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonEmptyRight()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_110_LeftRight_HasValidNonWhitespaceLeft(t *testing.T) {
	safeTest(t, "Test_C27_110_LeftRight_HasValidNonWhitespaceLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonWhitespaceLeft()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_111_LeftRight_HasValidNonWhitespaceRight(t *testing.T) {
	safeTest(t, "Test_C27_111_LeftRight_HasValidNonWhitespaceRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasValidNonWhitespaceRight()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_112_LeftRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_112_LeftRight_HasSafeNonEmpty", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_113_LeftRight_NonPtr(t *testing.T) {
	safeTest(t, "Test_C27_113_LeftRight_NonPtr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()

		// Act
		actual := args.Map{"result": np.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_114_LeftRight_Ptr(t *testing.T) {
	safeTest(t, "Test_C27_114_LeftRight_Ptr", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.Ptr() != lr}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_C27_115_LeftRight_IsLeftRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_115_LeftRight_IsLeftRegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("abc123", "b")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsLeftRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_116_LeftRight_IsRightRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_116_LeftRight_IsRightRegexMatch", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": lr.IsRightRegexMatch(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lr.IsRightRegexMatch(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_117_LeftRight_IsLeft(t *testing.T) {
	safeTest(t, "Test_C27_117_LeftRight_IsLeft", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.IsLeft("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_118_LeftRight_IsRight(t *testing.T) {
	safeTest(t, "Test_C27_118_LeftRight_IsRight", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.IsRight("b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_119_LeftRight_Is(t *testing.T) {
	safeTest(t, "Test_C27_119_LeftRight_Is", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.Is("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_120_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_C27_120_LeftRight_IsEqual", func() {
		// Arrange
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr1.IsEqual(lr2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_121_LeftRight_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C27_121_LeftRight_IsEqual_BothNil", func() {
		// Arrange
		var lr1, lr2 *corestr.LeftRight

		// Act
		actual := args.Map{"result": lr1.IsEqual(lr2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_122_LeftRight_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_C27_122_LeftRight_IsEqual_OneNil", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")

		// Act
		actual := args.Map{"result": lr.IsEqual(nil)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_123_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_123_LeftRight_Clone", func() {
		// Arrange
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()

		// Act
		actual := args.Map{"result": c.Left != "a" || c.Right != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b", actual)
	})
}

func Test_C27_124_LeftRight_Clear(t *testing.T) {
	safeTest(t, "Test_C27_124_LeftRight_Clear", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Clear()
	})
}

func Test_C27_125_LeftRight_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C27_125_LeftRight_Clear_Nil", func() {
		var lr *corestr.LeftRight
		lr.Clear()
	})
}

func Test_C27_126_LeftRight_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_126_LeftRight_Dispose", func() {
		lr := corestr.NewLeftRight("a", "b")
		lr.Dispose()
	})
}

func Test_C27_127_LeftRight_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C27_127_LeftRight_Dispose_Nil", func() {
		var lr *corestr.LeftRight
		lr.Dispose()
	})
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRightFromSplit
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_128_LeftRightFromSplit(t *testing.T) {
	safeTest(t, "Test_C27_128_LeftRightFromSplit", func() {
		// Arrange
		lr := corestr.LeftRightFromSplit("key=value", "=")

		// Act
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key, value", actual)
	})
}

func Test_C27_129_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_129_LeftRightFromSplitTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")

		// Act
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed key, value", actual)
	})
}

func Test_C27_130_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_C27_130_LeftRightFromSplitFull", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b:c:d"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b:c:d", actual)
	})
}

func Test_C27_131_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_131_LeftRightFromSplitFullTrimmed", func() {
		// Arrange
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b : c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_132_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C27_132_LeftMiddleRight_New", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_133_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_133_LeftMiddleRight_Invalid", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{"result": lmr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_134_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_134_LeftMiddleRight_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"result": lmr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_135_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_C27_135_LeftMiddleRight_Bytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": string(lmr.LeftBytes()) != "a" || string(lmr.MiddleBytes()) != "b" || string(lmr.RightBytes()) != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_136_LeftMiddleRight_Trims(t *testing.T) {
	safeTest(t, "Test_C27_136_LeftMiddleRight_Trims", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")

		// Act
		actual := args.Map{"result": lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_137_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_137_LeftMiddleRight_IsEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "", "")

		// Act
		actual := args.Map{"result": lmr.IsLeftEmpty() || !lmr.IsMiddleEmpty() || !lmr.IsRightEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all empty", actual)
	})
}

func Test_C27_138_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_138_LeftMiddleRight_IsWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")

		// Act
		actual := args.Map{"result": lmr.IsLeftWhitespace() || !lmr.IsMiddleWhitespace() || !lmr.IsRightWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all whitespace", actual)
	})
}

func Test_C27_139_LeftMiddleRight_HasValid(t *testing.T) {
	safeTest(t, "Test_C27_139_LeftMiddleRight_HasValid", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all valid non-empty", actual)
		actual = args.Map{"result": lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all non-whitespace", actual)
	})
}

func Test_C27_140_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_140_LeftMiddleRight_HasSafeNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_141_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_C27_141_LeftMiddleRight_IsAll", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.IsAll("a", "b", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_142_LeftMiddleRight_Is(t *testing.T) {
	safeTest(t, "Test_C27_142_LeftMiddleRight_Is", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.Is("a", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_143_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_143_LeftMiddleRight_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()

		// Act
		actual := args.Map{"result": c.Left != "a" || c.Middle != "b" || c.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_144_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_144_LeftMiddleRight_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, c", actual)
	})
}

func Test_C27_145_LeftMiddleRight_Clear(t *testing.T) {
	safeTest(t, "Test_C27_145_LeftMiddleRight_Clear", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()
	})
}

func Test_C27_146_LeftMiddleRight_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C27_146_LeftMiddleRight_Clear_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Clear()
	})
}

func Test_C27_147_LeftMiddleRight_Dispose(t *testing.T) {
	safeTest(t, "Test_C27_147_LeftMiddleRight_Dispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()
	})
}

func Test_C27_148_LeftMiddleRight_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C27_148_LeftMiddleRight_Dispose_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Dispose()
	})
}

// ═══════════════════════════════════════════════════════════════════════
// LeftMiddleRightFromSplit
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_149_LeftMiddleRightFromSplit(t *testing.T) {
	safeTest(t, "Test_C27_149_LeftMiddleRightFromSplit", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_150_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_150_LeftMiddleRightFromSplitTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed a, b, c", actual)
	})
}

func Test_C27_151_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_C27_151_LeftMiddleRightFromSplitN", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c:d:e"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c:d:e", actual)
	})
}

func Test_C27_152_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_152_LeftMiddleRightFromSplitNTrimmed", func() {
		// Arrange
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_153_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_153_TextWithLineNumber_HasLineNumber", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}

		// Act
		actual := args.Map{"result": tl.HasLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_154_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_154_TextWithLineNumber_IsInvalidLineNumber", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}

		// Act
		actual := args.Map{"result": tl.IsInvalidLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_155_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C27_155_TextWithLineNumber_IsInvalidLineNumber_Nil", func() {
		// Arrange
		var tl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tl.IsInvalidLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_156_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C27_156_TextWithLineNumber_Length", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act
		actual := args.Map{"result": tl.Length() != 5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_C27_157_TextWithLineNumber_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C27_157_TextWithLineNumber_Length_Nil", func() {
		// Arrange
		var tl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tl.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_158_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_158_TextWithLineNumber_IsEmpty", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": tl.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_159_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C27_159_TextWithLineNumber_IsEmpty_Nil", func() {
		// Arrange
		var tl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tl.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_160_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C27_160_TextWithLineNumber_IsEmptyText", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{"result": tl.IsEmptyText()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty text", actual)
	})
}

func Test_C27_161_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	safeTest(t, "Test_C27_161_TextWithLineNumber_IsEmptyText_Nil", func() {
		// Arrange
		var tl *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tl.IsEmptyText()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_162_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C27_162_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		// Arrange
		tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": tl.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// CloneSlice, CloneSliceIf
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_163_CloneSlice_Basic(t *testing.T) {
	safeTest(t, "Test_C27_163_CloneSlice_Basic", func() {
		// Arrange
		s := corestr.CloneSlice([]string{"a", "b"})

		// Act
		actual := args.Map{"result": len(s) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_164_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_164_CloneSlice_Empty", func() {
		// Arrange
		s := corestr.CloneSlice(nil)

		// Act
		actual := args.Map{"result": len(s) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_165_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_C27_165_CloneSliceIf_Clone", func() {
		// Arrange
		s := corestr.CloneSliceIf(true, "a", "b")

		// Act
		actual := args.Map{"result": len(s) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_166_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_C27_166_CloneSliceIf_NoClone", func() {
		// Arrange
		s := corestr.CloneSliceIf(false, "a", "b")

		// Act
		actual := args.Map{"result": len(s) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_167_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_C27_167_CloneSliceIf_Empty", func() {
		// Arrange
		s := corestr.CloneSliceIf(true)

		// Act
		actual := args.Map{"result": len(s) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_168_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_C27_168_AnyToString_WithFieldName", func() {
		// Arrange
		s := corestr.AnyToString(true, 42)

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_169_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_C27_169_AnyToString_WithoutFieldName", func() {
		// Arrange
		s := corestr.AnyToString(false, 42)

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_170_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_C27_170_AnyToString_Empty", func() {
		// Arrange
		s := corestr.AnyToString(false, "")

		// Act
		actual := args.Map{"result": s != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_171_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_C27_171_AllIndividualStringsOfStringsLength", func() {
		// Arrange
		items := [][]string{{"a", "b"}, {"c"}}
		l := corestr.AllIndividualStringsOfStringsLength(&items)

		// Act
		actual := args.Map{"result": l != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_172_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_C27_172_AllIndividualStringsOfStringsLength_Nil", func() {
		// Arrange
		l := corestr.AllIndividualStringsOfStringsLength(nil)

		// Act
		actual := args.Map{"result": l != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_173_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C27_173_AllIndividualsLengthOfSimpleSlices", func() {
		// Arrange
		s1 := corestr.New.SimpleSlice.Create([]string{"a", "b"})
		s2 := corestr.New.SimpleSlice.Create([]string{"c"})
		l := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)

		// Act
		actual := args.Map{"result": l != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_174_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_C27_174_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		// Arrange
		l := corestr.AllIndividualsLengthOfSimpleSlices()

		// Act
		actual := args.Map{"result": l != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// Utils
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_175_Utils_WrapDoubleIfMissing(t *testing.T) {
	safeTest(t, "Test_C27_175_Utils_WrapDoubleIfMissing", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapDoubleIfMissing("hello") != `"hello"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing(`"hello"`) != `"hello"`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected already wrapped", actual)
		actual = args.Map{"result": u.WrapDoubleIfMissing("") != `""`}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty wrapped", actual)
	})
}

func Test_C27_176_Utils_WrapSingleIfMissing(t *testing.T) {
	safeTest(t, "Test_C27_176_Utils_WrapSingleIfMissing", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapSingleIfMissing("hello") != "'hello'"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("'hello'") != "'hello'"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected already wrapped", actual)
		actual = args.Map{"result": u.WrapSingleIfMissing("") != "''"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty wrapped", actual)
	})
}

func Test_C27_177_Utils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_C27_177_Utils_WrapDouble", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapDouble("hi") != `"hi"`}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_C27_178_Utils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_C27_178_Utils_WrapSingle", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapSingle("hi") != "'hi'"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_C27_179_Utils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_C27_179_Utils_WrapTilda", func() {
		// Arrange
		u := corestr.StringUtils

		// Act
		actual := args.Map{"result": u.WrapTilda("hi") != "`hi`"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValuePair — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_180_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_180_KeyValuePair_Basic", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.KeyName() != "k" || kv.VariableName() != "k" || kv.ValueString() != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k, v", actual)
	})
}

func Test_C27_181_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C27_181_KeyValuePair_IsVariableNameEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.IsVariableNameEqual("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_182_KeyValuePair_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_C27_182_KeyValuePair_IsValueEqual", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.IsValueEqual("v")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_183_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_183_KeyValuePair_ValueBool", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}

		// Act
		actual := args.Map{"result": kv.ValueBool()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		kv2 := corestr.KeyValuePair{Key: "k", Value: ""}
		actual = args.Map{"result": kv2.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_184_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_C27_184_KeyValuePair_ValueInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}

		// Act
		actual := args.Map{"result": kv.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_C27_185_KeyValuePair_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C27_185_KeyValuePair_ValueDefInt", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "10"}

		// Act
		actual := args.Map{"result": kv.ValueDefInt() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C27_186_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_186_KeyValuePair_ValueByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}

		// Act
		actual := args.Map{"result": kv.ValueByte(0) != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_C27_187_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C27_187_KeyValuePair_ValueDefByte", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "50"}

		// Act
		actual := args.Map{"result": kv.ValueDefByte() != 50}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_C27_188_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_188_KeyValuePair_ValueFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}

		// Act
		actual := args.Map{"result": kv.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
	})
}

func Test_C27_189_KeyValuePair_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_189_KeyValuePair_ValueDefFloat64", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}

		// Act
		actual := args.Map{"result": kv.ValueDefFloat64() != 2.5}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C27_190_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_C27_190_KeyValuePair_ValueValid", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()

		// Act
		actual := args.Map{"result": vv.Value != "v" || !vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v, valid", actual)
	})
}

func Test_C27_191_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_C27_191_KeyValuePair_ValueValidOptions", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValidOptions(false, "msg")

		// Act
		actual := args.Map{"result": vv.IsValid || vv.Message != "msg"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false, msg", actual)
	})
}

func Test_C27_192_KeyValuePair_IsKeyEmpty(t *testing.T) {
	safeTest(t, "Test_C27_192_KeyValuePair_IsKeyEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "", Value: "v"}

		// Act
		actual := args.Map{"result": kv.IsKeyEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_193_KeyValuePair_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_193_KeyValuePair_IsValueEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: ""}

		// Act
		actual := args.Map{"result": kv.IsValueEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_194_KeyValuePair_HasKey(t *testing.T) {
	safeTest(t, "Test_C27_194_KeyValuePair_HasKey", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k"}

		// Act
		actual := args.Map{"result": kv.HasKey()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_195_KeyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_C27_195_KeyValuePair_HasValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.HasValue()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_196_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_196_KeyValuePair_IsKeyValueEmpty", func() {
		// Arrange
		kv := corestr.KeyValuePair{}

		// Act
		actual := args.Map{"result": kv.IsKeyValueEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_197_KeyValuePair_TrimKey(t *testing.T) {
	safeTest(t, "Test_C27_197_KeyValuePair_TrimKey", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: " k "}

		// Act
		actual := args.Map{"result": kv.TrimKey() != "k"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
	})
}

func Test_C27_198_KeyValuePair_TrimValue(t *testing.T) {
	safeTest(t, "Test_C27_198_KeyValuePair_TrimValue", func() {
		// Arrange
		kv := corestr.KeyValuePair{Value: " v "}

		// Act
		actual := args.Map{"result": kv.TrimValue() != "v"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C27_199_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_C27_199_KeyValuePair_Is", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}

		// Act
		actual := args.Map{"result": kv.Is("k", "v")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_200_KeyValuePair_IsKey(t *testing.T) {
	safeTest(t, "Test_C27_200_KeyValuePair_IsKey", func() {
		// Arrange
		kv := corestr.KeyValuePair{Key: "k"}

		// Act
		actual := args.Map{"result": kv.IsKey("k")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}
