package corestrtests

import (
	"encoding/json"
	"regexp"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===================== ValidValue =====================

func Test_C49_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_New", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.Value != "hello" || !vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Empty", func() {
		// Arrange
		vv := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{"result": vv.IsEmpty() || !vv.IsValid}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Invalid", func() {
		// Arrange
		vv := corestr.InvalidValidValue("err")

		// Act
		actual := args.Map{"result": vv.IsValid || vv.Message != "err"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_InvalidNoMessage", func() {
		// Arrange
		vv := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{"result": vv.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_ValidValue_UsingAny(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_UsingAny", func() {
		// Arrange
		vv := corestr.NewValidValueUsingAny(false, true, "test")

		// Act
		actual := args.Map{"result": vv.Value == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_ValidValue_UsingAnyAutoValid(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_UsingAnyAutoValid", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "test")
		_ = vv
	})
}

func Test_C49_ValidValue_ValueBytesOnce(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBytesOnce", func() {
		// Arrange
		vv := corestr.NewValidValue("abc")
		b := vv.ValueBytesOnce()

		// Act
		actual := args.Map{"result": len(b) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		b2 := vv.ValueBytesOnce() // cached
		actual = args.Map{"result": len(b2) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cached 3", actual)
	})
}

func Test_C49_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBytesOncePtr", func() {
		// Arrange
		vv := corestr.NewValidValue("ab")
		b := vv.ValueBytesOncePtr()

		// Act
		actual := args.Map{"result": len(b) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("   ")

		// Act
		actual := args.Map{"result": vv.IsWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
	})
}

func Test_C49_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Trim", func() {
		// Arrange
		vv := corestr.NewValidValue("  x  ")

		// Act
		actual := args.Map{"result": vv.Trim() != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_C49_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasValidNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": vv.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasValidNonWhitespace", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": vv.HasValidNonWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasSafeNonEmpty", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": vv.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBool", func() {
		// Arrange
		vv := corestr.NewValidValue("true")

		// Act
		actual := args.Map{"result": vv.ValueBool()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		vv2 := corestr.NewValidValue("invalid")
		actual = args.Map{"result": vv2.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		vv3 := corestr.NewValidValue("")
		actual = args.Map{"result": vv3.ValueBool()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C49_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueInt", func() {
		// Arrange
		vv := corestr.NewValidValue("42")

		// Act
		actual := args.Map{"result": vv.ValueInt(0) != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		vv2 := corestr.NewValidValue("bad")
		actual = args.Map{"result": vv2.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C49_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefInt", func() {
		// Arrange
		vv := corestr.NewValidValue("10")

		// Act
		actual := args.Map{"result": vv.ValueDefInt() != 10}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C49_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueByte", func() {
		// Arrange
		vv := corestr.NewValidValue("100")
		b := vv.ValueByte(0)

		// Act
		actual := args.Map{"result": b != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		vv2 := corestr.NewValidValue("300")
		b2 := vv2.ValueByte(0)
		actual = args.Map{"result": b2 != 255}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255 for overflow", actual)
		vv3 := corestr.NewValidValue("-1")
		b3 := vv3.ValueByte(0)
		actual = args.Map{"result": b3 != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for negative", actual)
	})
}

func Test_C49_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefByte", func() {
		// Arrange
		vv := corestr.NewValidValue("50")

		// Act
		actual := args.Map{"result": vv.ValueDefByte() != 50}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_C49_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("3.14")
		f := vv.ValueFloat64(0)

		// Act
		actual := args.Map{"result": f < 3.13 || f > 3.15}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected ~3.14", actual)
	})
}

func Test_C49_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefFloat64", func() {
		// Arrange
		vv := corestr.NewValidValue("bad")

		// Act
		actual := args.Map{"result": vv.ValueDefFloat64() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C49_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Is", func() {
		// Arrange
		vv := corestr.NewValidValue("x")

		// Act
		actual := args.Map{"result": vv.Is("x")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C49_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsAnyOf", func() {
		// Arrange
		vv := corestr.NewValidValue("b")

		// Act
		actual := args.Map{"result": vv.IsAnyOf("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual = args.Map{"result": vv.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty values should return true", actual)
		actual = args.Map{"result": vv.IsAnyOf("x")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not match", actual)
	})
}

func Test_C49_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello world")

		// Act
		actual := args.Map{"result": vv.IsContains("world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
	})
}

func Test_C49_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsAnyContains", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{"result": vv.IsAnyContains("ell")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
		actual = args.Map{"result": vv.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty should return true", actual)
	})
}

func Test_C49_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsEqualNonSensitive", func() {
		// Arrange
		vv := corestr.NewValidValue("Hello")

		// Act
		actual := args.Map{"result": vv.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match case-insensitive", actual)
	})
}

func Test_C49_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsRegexMatches", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": vv.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual = args.Map{"result": vv.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_C49_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindString", func() {
		// Arrange
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": vv.RegexFindString(re) != "123"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual = args.Map{"result": vv.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return empty", actual)
	})
}

func Test_C49_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindAllStrings", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		result := vv.RegexFindAllStrings(re, -1)

		// Act
		actual := args.Map{"result": len(result) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		result2 := vv.RegexFindAllStrings(nil, -1)
		actual = args.Map{"result": len(result2) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C49_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindAllStringsWithFlag", func() {
		// Arrange
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, has := vv.RegexFindAllStringsWithFlag(re, -1)

		// Act
		actual := args.Map{"result": has || len(items) != 2}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_, has2 := vv.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": has2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_C49_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Split", func() {
		// Arrange
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")

		// Act
		actual := args.Map{"result": len(parts) != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C49_ValidValue_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_SplitNonEmpty", func() {
		vv := corestr.NewValidValue("a,,b")
		parts := vv.SplitNonEmpty(",")
		_ = parts
	})
}

func Test_C49_ValidValue_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_SplitTrimNonWhitespace", func() {
		vv := corestr.NewValidValue(" a , , b ")
		parts := vv.SplitTrimNonWhitespace(",")
		_ = parts
	})
}

func Test_C49_ValidValue_Clone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clone", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		cloned := vv.Clone()

		// Act
		actual := args.Map{"result": cloned.Value != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone failed", actual)
	})
}

func Test_C49_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clone_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue
		cloned := vv.Clone()

		// Act
		actual := args.Map{"result": cloned != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil clone should be nil", actual)
	})
}

func Test_C49_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_String", func() {
		// Arrange
		vv := corestr.NewValidValue("test")

		// Act
		actual := args.Map{"result": vv.String() != "test"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected test", actual)
	})
}

func Test_C49_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_String_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"result": vv.String() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil string should be empty", actual)
	})
}

func Test_C49_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_FullString", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		s := vv.FullString()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_FullString_Nil", func() {
		// Arrange
		var vv *corestr.ValidValue

		// Act
		actual := args.Map{"result": vv.FullString() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	})
}

func Test_C49_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clear", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		vv.Clear()

		// Act
		actual := args.Map{"result": vv.Value != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty after clear", actual)
	})
}

func Test_C49_ValidValue_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clear_Nil", func() {
		var vv *corestr.ValidValue
		vv.Clear()
	})
}

func Test_C49_ValidValue_Dispose(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Dispose", func() {
		vv := corestr.NewValidValue("x")
		vv.Dispose()
	})
}

func Test_C49_ValidValue_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Dispose_Nil", func() {
		var vv *corestr.ValidValue
		vv.Dispose()
	})
}

func Test_C49_ValidValue_Json(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Json", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		j := vv.Json()

		// Act
		actual := args.Map{"result": j.Error}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C49_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_JsonPtr", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		j := vv.JsonPtr()

		// Act
		actual := args.Map{"result": j == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C49_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Serialize", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		_, err := vv.Serialize()

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C49_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ParseInjectUsingJson", func() {
		// Arrange
		vv := corestr.NewValidValue("x")
		j := vv.Json()
		vv2 := corestr.NewValidValueEmpty()
		_, err := vv2.ParseInjectUsingJson(&j)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ===================== ValidValues =====================

func Test_C49_ValidValues_New(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_New", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)

		// Act
		actual := args.Map{"result": vvs.IsEmpty() || vvs.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Empty", func() {
		// Arrange
		vvs := corestr.EmptyValidValues()

		// Act
		actual := args.Map{"result": vvs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Add", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{"result": vvs.Count() != 2 || vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddFull", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "x", "msg")

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_UsingValues(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_UsingValues", func() {
		// Arrange
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vvs := corestr.NewValidValuesUsingValues(v1)

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_UsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_UsingValues_Empty", func() {
		// Arrange
		vvs := corestr.NewValidValuesUsingValues()

		// Act
		actual := args.Map{"result": vvs.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")

		// Act
		actual := args.Map{"result": vvs.SafeValueAt(0) != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
		actual = args.Map{"result": vvs.SafeValueAt(99) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for out of range", actual)
	})
}

func Test_C49_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValidValueAt", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")

		// Act
		actual := args.Map{"result": vvs.SafeValidValueAt(0) != "x"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_C49_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		result := vvs.SafeValuesAtIndexes(0, 2)

		// Act
		actual := args.Map{"result": len(result) != 2 || result[0] != "a" || result[1] != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValidValuesAtIndexes", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.SafeValidValuesAtIndexes(0)

		// Act
		actual := args.Map{"result": len(result) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Strings", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.Strings()

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_FullStrings", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.FullStrings()

		// Act
		actual := args.Map{"result": len(s) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_String", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_HasIndex", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")

		// Act
		actual := args.Map{"result": vvs.HasIndex(0)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have index 0", actual)
		actual = args.Map{"result": vvs.HasIndex(1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have index 1", actual)
	})
}

func Test_C49_ValidValues_LastIndex(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_LastIndex", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")

		// Act
		actual := args.Map{"result": vvs.LastIndex() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Find", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Find_Break(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Find_Break", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, true
		})

		// Act
		actual := args.Map{"result": len(found) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break after first)", actual)
	})
}

func Test_C49_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew", func() {
		// Arrange
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew_EmptyClone", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.ConcatNew(true)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew_EmptyNoClone", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.ConcatNew(false)

		// Act
		actual := args.Map{"result": result != vvs}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ref", actual)
	})
}

func Test_C49_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddValidValues", func() {
		// Arrange
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)

		// Act
		actual := args.Map{"result": vvs1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddValidValues_Nil", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddValidValues(nil)
	})
}

func Test_C49_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Adds", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.Adds(corestr.ValidValue{Value: "a"})

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddsPtr", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddsPtr(corestr.NewValidValue("a"))

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashsetMap", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": vvs.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashset", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)

		// Act
		actual := args.Map{"result": vvs.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashset_Nil", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddHashset(nil)
	})
}

func Test_C49_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Hashmap", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "key", "val")
		hm := vvs.Hashmap()

		// Act
		actual := args.Map{"result": hm.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Map", func() {
		// Arrange
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "key", "val")
		m := vvs.Map()

		// Act
		actual := args.Map{"result": len(m) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ===================== TextWithLineNumber =====================

func Test_C49_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_HasLineNumber", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{"result": tln.HasLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsInvalidLineNumber", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: -1}

		// Act
		actual := args.Map{"result": tln.IsInvalidLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C49_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsInvalidLineNumber_Nil", func() {
		// Arrange
		var tln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tln.IsInvalidLineNumber()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
	})
}

func Test_C49_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_Length", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}

		// Act
		actual := args.Map{"result": tln.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C49_TextWithLineNumber_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_Length_Nil", func() {
		// Arrange
		var tln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tln.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil length should be 0", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmpty", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": tln.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmpty_Nil", func() {
		// Arrange
		var tln *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"result": tln.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmptyText", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{"result": tln.IsEmptyText()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty text", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		// Arrange
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}

		// Act
		actual := args.Map{"result": tln.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// ===================== ValueStatus =====================

func Test_C49_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_Invalid", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("msg")

		// Act
		actual := args.Map{"result": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_InvalidNoMessage", func() {
		// Arrange
		vs := corestr.InvalidValueStatusNoMessage()

		// Act
		actual := args.Map{"result": vs.ValueValid.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_Clone", func() {
		// Arrange
		vs := corestr.InvalidValueStatus("msg")
		cloned := vs.Clone()

		// Act
		actual := args.Map{"result": cloned.ValueValid.Message != "msg"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone failed", actual)
	})
}

// ===================== LeftMiddleRight =====================

func Test_C49_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_New", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Invalid", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRight("err")

		// Act
		actual := args.Map{"result": lmr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_InvalidNoMessage", func() {
		// Arrange
		lmr := corestr.InvalidLeftMiddleRightNoMessage()

		// Act
		actual := args.Map{"result": lmr.IsValid}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Bytes", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": len(lmr.LeftBytes()) != 1 || len(lmr.MiddleBytes()) != 1 || len(lmr.RightBytes()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 byte each", actual)
	})
}

func Test_C49_LeftMiddleRight_Trim(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Trim", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")

		// Act
		actual := args.Map{"result": lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "trim failed", actual)
	})
}

func Test_C49_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("", "b", "")

		// Act
		actual := args.Map{"result": lmr.IsLeftEmpty() || lmr.IsMiddleEmpty() || !lmr.IsRightEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty checks failed", actual)
	})
}

func Test_C49_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsWhitespace", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")

		// Act
		actual := args.Map{"result": lmr.IsLeftWhitespace() || !lmr.IsMiddleWhitespace() || !lmr.IsRightWhitespace()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "whitespace checks failed", actual)
	})
}

func Test_C49_LeftMiddleRight_HasValid(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_HasValid", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_HasSafeNonEmpty", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsAll", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.IsAll("a", "b", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_Is(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Is", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")

		// Act
		actual := args.Map{"result": lmr.Is("a", "c")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Clone", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()

		// Act
		actual := args.Map{"result": cloned.Left != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone failed", actual)
	})
}

func Test_C49_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_ToLeftRight", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()

		// Act
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "conversion failed", actual)
	})
}

func Test_C49_LeftMiddleRight_Clear(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Clear", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Clear()
	})
}

func Test_C49_LeftMiddleRight_Clear_Nil(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Clear_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Clear()
	})
}

func Test_C49_LeftMiddleRight_Dispose(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Dispose", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lmr.Dispose()
	})
}

func Test_C49_LeftMiddleRight_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Dispose_Nil", func() {
		var lmr *corestr.LeftMiddleRight
		lmr.Dispose()
	})
}

// ===================== CollectionsOfCollection =====================

func Test_C49_CollectionsOfCollection_Basic(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Basic", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)

		// Act
		actual := args.Map{"result": coc.IsEmpty() || coc.HasItems() || coc.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_CollectionsOfCollection_Add(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Add", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(col)

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AddStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})

		// Act
		actual := args.Map{"result": coc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AllIndividualItemsLength", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})

		// Act
		actual := args.Map{"result": coc.AllIndividualItemsLength() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_CollectionsOfCollection_Items(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Items", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(col)

		// Act
		actual := args.Map{"result": len(coc.Items()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_List(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_List", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		list := coc.List(0)

		// Act
		actual := args.Map{"result": len(list) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_CollectionsOfCollection_ToCollection(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_ToCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		col := coc.ToCollection()

		// Act
		actual := args.Map{"result": col.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_String(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_String", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		s := coc.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_CollectionsOfCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_JSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		data, err := json.Marshal(coc)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		coc2 := corestr.New.CollectionsOfCollection.Cap(5)
		err = json.Unmarshal(data, coc2)
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C49_CollectionsOfCollection_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AsInterfaces", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		_ = coc.AsJsonContractsBinder()
		_ = coc.AsJsoner()
		_ = coc.AsJsonMarshaller()
		_ = coc.AsJsonParseSelfInjector()
	})
}

// ===================== HashsetsCollection =====================

func Test_C49_HashsetsCollection_Basic(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Basic", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc.IsEmpty() || hsc.HasItems() || hsc.Length() != 0}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Add", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(hs)

		// Act
		actual := args.Map{"result": hsc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_AddNonNil(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddNonNil", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddNonNil(nil)

		// Act
		actual := args.Map{"result": hsc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not add", actual)
		hsc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))
		actual = args.Map{"result": hsc.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddNonEmpty", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddNonEmpty(corestr.Empty.Hashset())

		// Act
		actual := args.Map{"result": hsc.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not add", actual)
	})
}

func Test_C49_HashsetsCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Adds", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Adds(hs)

		// Act
		actual := args.Map{"result": hsc.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_StringsList(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_StringsList", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(hs)
		list := hsc.StringsList()

		// Act
		actual := args.Map{"result": len(list) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_HasAll(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_HasAll", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(hs)

		// Act
		actual := args.Map{"result": hsc.HasAll("a", "b")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have all", actual)
	})
}

func Test_C49_HashsetsCollection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_HasAll_Empty", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()

		// Act
		actual := args.Map{"result": hsc.HasAll("a")}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should return false", actual)
	})
}

func Test_C49_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_IsEqual", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(hs)
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(hs)

		// Act
		actual := args.Map{"result": hsc1.IsEqualPtr(hsc2)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C49_HashsetsCollection_IsEqual_SameRef(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_IsEqual_SameRef", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hsc.IsEqualPtr(hsc)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref", actual)
	})
}

func Test_C49_HashsetsCollection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ConcatNew", func() {
		// Arrange
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		result := hsc1.ConcatNew(hsc2)

		// Act
		actual := args.Map{"result": result.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_HashsetsCollection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ConcatNew_Empty", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hsc.ConcatNew()

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddHashsetsCollection", func() {
		// Arrange
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		hsc1.AddHashsetsCollection(hsc2)

		// Act
		actual := args.Map{"result": hsc1.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_HashsetsCollection_AddHashsetsCollection_Nil(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddHashsetsCollection_Nil", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddHashsetsCollection(nil)
	})
}

func Test_C49_HashsetsCollection_LastIndex(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_LastIndex", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))

		// Act
		actual := args.Map{"result": hsc.LastIndex() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C49_HashsetsCollection_ListPtr(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ListPtr", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		p := hsc.ListPtr()

		// Act
		actual := args.Map{"result": p == nil || len(*p) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ListDirectPtr", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		p := hsc.ListDirectPtr()

		// Act
		actual := args.Map{"result": p == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C49_HashsetsCollection_String(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_String", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		s := hsc.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_HashsetsCollection_String_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_String_Empty", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		s := hsc.String()

		// Act
		actual := args.Map{"result": s == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C49_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Join", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		j := hsc.Join(",")

		// Act
		actual := args.Map{"result": j == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_HashsetsCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_JSON", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := json.Marshal(hsc)

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		hsc2 := corestr.Empty.HashsetsCollection()
		err = json.Unmarshal(data, hsc2)
		actual = args.Map{"result": err}
		expected = args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C49_HashsetsCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Serialize", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		_, err := hsc.Serialize()

		// Act
		actual := args.Map{"result": err}

		// Assert
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C49_HashsetsCollection_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AsInterfaces", func() {
		hsc := corestr.Empty.HashsetsCollection()
		_ = hsc.AsJsonContractsBinder()
		_ = hsc.AsJsoner()
		_ = hsc.AsJsonMarshaller()
		_ = hsc.AsJsonParseSelfInjector()
	})
}

// ===================== SimpleStringOnce (key methods) =====================

func Test_C49_SimpleStringOnce_Basic(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Basic", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOnce()

		// Act
		actual := args.Map{"result": sso.IsInitialized() || sso.IsDefined() || !sso.IsUninitialized() || !sso.IsInvalid()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_C49_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_SetOnUninitialized", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOncePtr()
		err := sso.SetOnUninitialized("hello")

		// Act
		actual := args.Map{"result": err != nil || sso.Value() != "hello" || !sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "set failed", actual)
		err2 := sso.SetOnUninitialized("world")
		actual = args.Map{"result": err2 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should error on already initialized", actual)
	})
}

func Test_C49_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetSetOnce", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetSetOnce("first")

		// Act
		actual := args.Map{"result": v != "first"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first", actual)
		v2 := sso.GetSetOnce("second")
		actual = args.Map{"result": v2 != "first"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return first (already set)", actual)
	})
}

func Test_C49_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetOnce", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetOnce()

		// Act
		actual := args.Map{"result": v != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be initialized", actual)
	})
}

func Test_C49_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetOnceFunc", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetOnceFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"result": v != "computed"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected computed", actual)
		v2 := sso.GetOnceFunc(func() string { return "other" })
		actual = args.Map{"result": v2 != "computed"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return first value", actual)
	})
}

func Test_C49_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Invalidate", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOncePtr()
		sso.SetOnUninitialized("x")
		sso.Invalidate()

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be uninitialized", actual)
	})
}

func Test_C49_SimpleStringOnce_Reset(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Reset", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOncePtr()
		sso.SetOnUninitialized("x")
		sso.Reset()

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be uninitialized", actual)
	})
}

func Test_C49_SimpleStringOnce_Boolean(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Boolean", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("true")

		// Act
		actual := args.Map{"result": sso.Boolean(false)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		sso2 := corestr.New.SimpleStringOnce.Init("yes")
		actual = args.Map{"result": sso2.Boolean(false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for yes", actual)
	})
}

func Test_C49_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Int", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")

		// Act
		actual := args.Map{"result": sso.Int() != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_C49_SimpleStringOnce_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_IsEmpty", func() {
		// Arrange
		sso := corestr.Empty.SimpleStringOnce()

		// Act
		actual := args.Map{"result": sso.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_ConcatNew", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		result := sso.ConcatNew(" world")

		// Act
		actual := args.Map{"result": result.Value() != "hello world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)
	})
}

func Test_C49_SimpleStringOnce_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_ConcatNewUsingStrings", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a")
		result := sso.ConcatNewUsingStrings(",", "b", "c")

		// Act
		actual := args.Map{"result": result.Value() != "a,b,c"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_C49_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_WithinRange", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")
		val, inRange := sso.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": inRange || val != 50}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected in range", actual)
		val2, inRange2 := sso.WithinRange(true, 60, 100)
		actual = args.Map{"result": inRange2 || val2 != 60}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary min", actual)
		sso3 := corestr.New.SimpleStringOnce.Init("200")
		val3, inRange3 := sso3.WithinRange(true, 0, 100)
		actual = args.Map{"result": inRange3 || val3 != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary max", actual)
	})
}

func Test_C49_SimpleStringOnce_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_WithinRange_NoBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")
		val, inRange := sso.WithinRange(false, 0, 100)

		// Act
		actual := args.Map{"result": inRange}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		_ = val
	})
}

// ===================== DataModels =====================

func Test_C49_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashmapDataModel", func() {
		// Arrange
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		dm := corestr.NewHashmapsDataModelUsing(hm)
		hm2 := corestr.NewHashmapUsingDataModel(dm)

		// Act
		actual := args.Map{"result": hm2.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashsetDataModel", func() {
		// Arrange
		hs := corestr.New.Hashset.Strings([]string{"a"})
		dm := corestr.NewHashsetsDataModelUsing(hs)
		hs2 := corestr.NewHashsetUsingDataModel(dm)

		// Act
		actual := args.Map{"result": hs2.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollectionDataModel", func() {
		// Arrange
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)

		// Act
		actual := args.Map{"result": hsc2.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ===================== AddsUsingProcessorAsync =====================

func Test_C49_LinkedCollections_AddsUsingProcessorAsync(t *testing.T) {
	safeTest(t, "Test_C49_LinkedCollections_AddsUsingProcessorAsync", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		processor := func(any any, index int) *corestr.Collection {
			return corestr.New.Collection.Strings([]string{any.(string)})
		}
		lc.AddsUsingProcessorAsync(wg, processor, true, "hello")
		wg.Wait()

		// Act
		actual := args.Map{"result": lc.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_LinkedCollections_AddsUsingProcessorAsync_NilSkip(t *testing.T) {
	safeTest(t, "Test_C49_LinkedCollections_AddsUsingProcessorAsync_NilSkip", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		processor := func(any any, index int) *corestr.Collection {
			return nil
		}
		lc.AddsUsingProcessorAsync(wg, processor, true)
		wg.Wait()
	})
}

// ===================== Funcs types coverage =====================

func Test_C49_ReturningBool(t *testing.T) {
	safeTest(t, "Test_C49_ReturningBool", func() {
		// Arrange
		rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}

		// Act
		actual := args.Map{"result": rb.IsBreak || rb.IsKeep}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}
