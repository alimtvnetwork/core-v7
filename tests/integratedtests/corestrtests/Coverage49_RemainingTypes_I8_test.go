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
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.Value != "hello" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Empty", func() {
		vv := corestr.NewValidValueEmpty()
		actual := args.Map{"result": vv.IsEmpty() || !vv.IsValid}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Invalid", func() {
		vv := corestr.InvalidValidValue("err")
		actual := args.Map{"result": vv.IsValid || vv.Message != "err"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_InvalidNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		actual := args.Map{"result": vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_ValidValue_UsingAny(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_UsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, "test")
		actual := args.Map{"result": vv.Value == ""}
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
		vv := corestr.NewValidValue("abc")
		b := vv.ValueBytesOnce()
		actual := args.Map{"result": len(b) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		b2 := vv.ValueBytesOnce() // cached
		actual := args.Map{"result": len(b2) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cached 3", actual)
	})
}

func Test_C49_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBytesOncePtr", func() {
		vv := corestr.NewValidValue("ab")
		b := vv.ValueBytesOncePtr()
		actual := args.Map{"result": len(b) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsWhitespace", func() {
		vv := corestr.NewValidValue("   ")
		actual := args.Map{"result": vv.IsWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
	})
}

func Test_C49_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Trim", func() {
		vv := corestr.NewValidValue("  x  ")
		actual := args.Map{"result": vv.Trim() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_C49_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		actual := args.Map{"result": vv.HasValidNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasValidNonWhitespace", func() {
		vv := corestr.NewValidValue("x")
		actual := args.Map{"result": vv.HasValidNonWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasSafeNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		actual := args.Map{"result": vv.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		actual := args.Map{"result": vv.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		vv2 := corestr.NewValidValue("invalid")
		actual := args.Map{"result": vv2.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		vv3 := corestr.NewValidValue("")
		actual := args.Map{"result": vv3.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C49_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueInt", func() {
		vv := corestr.NewValidValue("42")
		actual := args.Map{"result": vv.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		vv2 := corestr.NewValidValue("bad")
		actual := args.Map{"result": vv2.ValueInt(99) != 99}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C49_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefInt", func() {
		vv := corestr.NewValidValue("10")
		actual := args.Map{"result": vv.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C49_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueByte", func() {
		vv := corestr.NewValidValue("100")
		b := vv.ValueByte(0)
		actual := args.Map{"result": b != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		vv2 := corestr.NewValidValue("300")
		b2 := vv2.ValueByte(0)
		actual := args.Map{"result": b2 != 255}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255 for overflow", actual)
		vv3 := corestr.NewValidValue("-1")
		b3 := vv3.ValueByte(0)
		actual := args.Map{"result": b3 != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for negative", actual)
	})
}

func Test_C49_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefByte", func() {
		vv := corestr.NewValidValue("50")
		actual := args.Map{"result": vv.ValueDefByte() != 50}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_C49_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		f := vv.ValueFloat64(0)
		actual := args.Map{"result": f < 3.13 || f > 3.15}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected ~3.14", actual)
	})
}

func Test_C49_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefFloat64", func() {
		vv := corestr.NewValidValue("bad")
		actual := args.Map{"result": vv.ValueDefFloat64() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C49_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Is", func() {
		vv := corestr.NewValidValue("x")
		actual := args.Map{"result": vv.Is("x")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
	})
}

func Test_C49_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		actual := args.Map{"result": vv.IsAnyOf("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": vv.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty values should return true", actual)
		actual := args.Map{"result": vv.IsAnyOf("x")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not match", actual)
	})
}

func Test_C49_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		actual := args.Map{"result": vv.IsContains("world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
	})
}

func Test_C49_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello")
		actual := args.Map{"result": vv.IsAnyContains("ell")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should contain", actual)
		actual := args.Map{"result": vv.IsAnyContains()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty should return true", actual)
	})
}

func Test_C49_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		actual := args.Map{"result": vv.IsEqualNonSensitive("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match case-insensitive", actual)
	})
}

func Test_C49_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsRegexMatches", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": vv.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should match", actual)
		actual := args.Map{"result": vv.IsRegexMatches(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_C49_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindString", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": vv.RegexFindString(re) != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual := args.Map{"result": vv.RegexFindString(nil) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return empty", actual)
	})
}

func Test_C49_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindAllStrings", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		result := vv.RegexFindAllStrings(re, -1)
		actual := args.Map{"result": len(result) != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		result2 := vv.RegexFindAllStrings(nil, -1)
		actual := args.Map{"result": len(result2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C49_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindAllStringsWithFlag", func() {
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, has := vv.RegexFindAllStringsWithFlag(re, -1)
		actual := args.Map{"result": has || len(items) != 2}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		_, has2 := vv.RegexFindAllStringsWithFlag(nil, -1)
		actual := args.Map{"result": has2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil regex should return false", actual)
	})
}

func Test_C49_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")
		actual := args.Map{"result": len(parts) != 3}
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
		vv := corestr.NewValidValue("x")
		cloned := vv.Clone()
		actual := args.Map{"result": cloned.Value != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone failed", actual)
	})
}

func Test_C49_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clone_Nil", func() {
		var vv *corestr.ValidValue
		cloned := vv.Clone()
		actual := args.Map{"result": cloned != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil clone should be nil", actual)
	})
}

func Test_C49_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_String", func() {
		vv := corestr.NewValidValue("test")
		actual := args.Map{"result": vv.String() != "test"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected test", actual)
	})
}

func Test_C49_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_String_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"result": vv.String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil string should be empty", actual)
	})
}

func Test_C49_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("x")
		s := vv.FullString()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_FullString_Nil", func() {
		var vv *corestr.ValidValue
		actual := args.Map{"result": vv.FullString() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should return empty", actual)
	})
}

func Test_C49_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clear", func() {
		vv := corestr.NewValidValue("x")
		vv.Clear()
		actual := args.Map{"result": vv.Value != ""}
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
		vv := corestr.NewValidValue("x")
		j := vv.Json()
		actual := args.Map{"result": j.Error}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "j.Error", actual)
	})
}

func Test_C49_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_JsonPtr", func() {
		vv := corestr.NewValidValue("x")
		j := vv.JsonPtr()
		actual := args.Map{"result": j == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C49_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("x")
		_, err := vv.Serialize()
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C49_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ParseInjectUsingJson", func() {
		vv := corestr.NewValidValue("x")
		j := vv.Json()
		vv2 := corestr.NewValidValueEmpty()
		_, err := vv2.ParseInjectUsingJson(&j)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

// ===================== ValidValues =====================

func Test_C49_ValidValues_New(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_New", func() {
		vvs := corestr.NewValidValues(5)
		actual := args.Map{"result": vvs.IsEmpty() || vvs.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Empty", func() {
		vvs := corestr.EmptyValidValues()
		actual := args.Map{"result": vvs.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Add", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		actual := args.Map{"result": vvs.Count() != 2 || vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddFull", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "x", "msg")
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_UsingValues(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_UsingValues", func() {
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vvs := corestr.NewValidValuesUsingValues(v1)
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_UsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_UsingValues_Empty", func() {
		vvs := corestr.NewValidValuesUsingValues()
		actual := args.Map{"result": vvs.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		actual := args.Map{"result": vvs.SafeValueAt(0) != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
		actual := args.Map{"result": vvs.SafeValueAt(99) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for out of range", actual)
	})
}

func Test_C49_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValidValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		actual := args.Map{"result": vvs.SafeValidValueAt(0) != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_C49_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		result := vvs.SafeValuesAtIndexes(0, 2)
		actual := args.Map{"result": len(result) != 2 || result[0] != "a" || result[1] != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValidValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.SafeValidValuesAtIndexes(0)
		actual := args.Map{"result": len(result) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.Strings()
		actual := args.Map{"result": len(s) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_FullStrings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.FullStrings()
		actual := args.Map{"result": len(s) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_String", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_HasIndex", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		actual := args.Map{"result": vvs.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have index 0", actual)
		actual := args.Map{"result": vvs.HasIndex(1)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should not have index 1", actual)
	})
}

func Test_C49_ValidValues_LastIndex(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_LastIndex", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		actual := args.Map{"result": vvs.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Find", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})
		actual := args.Map{"result": len(found) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Find_Break(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Find_Break", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, true
		})
		actual := args.Map{"result": len(found) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 (break after first)", actual)
	})
}

func Test_C49_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew", func() {
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew_EmptyClone", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.ConcatNew(true)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew_EmptyNoClone", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.ConcatNew(false)
		actual := args.Map{"result": result != vvs}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same ref", actual)
	})
}

func Test_C49_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddValidValues", func() {
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)
		actual := args.Map{"result": vvs1.Length() != 2}
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
		vvs := corestr.NewValidValues(5)
		vvs.Adds(corestr.ValidValue{Value: "a"})
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddsPtr", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddsPtr(corestr.NewValidValue("a"))
		actual := args.Map{"result": vvs.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashsetMap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"result": vvs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashset", func() {
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)
		actual := args.Map{"result": vvs.Length() != 1}
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
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "key", "val")
		hm := vvs.Hashmap()
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Map", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "key", "val")
		m := vvs.Map()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ===================== TextWithLineNumber =====================

func Test_C49_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_HasLineNumber", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		actual := args.Map{"result": tln.HasLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsInvalidLineNumber", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1}
		actual := args.Map{"result": tln.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C49_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsInvalidLineNumber_Nil", func() {
		var tln *corestr.TextWithLineNumber
		actual := args.Map{"result": tln.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil should be invalid", actual)
	})
}

func Test_C49_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_Length", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}
		actual := args.Map{"result": tln.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C49_TextWithLineNumber_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_Length_Nil", func() {
		var tln *corestr.TextWithLineNumber
		actual := args.Map{"result": tln.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil length should be 0", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmpty", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": tln.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmpty_Nil", func() {
		var tln *corestr.TextWithLineNumber
		actual := args.Map{"result": tln.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "nil should be empty", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmptyText", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		actual := args.Map{"result": tln.IsEmptyText()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty text", actual)
	})
}

func Test_C49_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": tln.IsEmptyTextLineBoth()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// ===================== ValueStatus =====================

func Test_C49_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("msg")
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("msg")
		cloned := vs.Clone()
		actual := args.Map{"result": cloned.ValueValid.Message != "msg"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone failed", actual)
	})
}

// ===================== LeftMiddleRight =====================

func Test_C49_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}

func Test_C49_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		actual := args.Map{"result": lmr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{"result": lmr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be invalid", actual)
	})
}

func Test_C49_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Bytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": len(lmr.LeftBytes()) != 1 || len(lmr.MiddleBytes()) != 1 || len(lmr.RightBytes()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 byte each", actual)
	})
}

func Test_C49_LeftMiddleRight_Trim(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Trim", func() {
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
		actual := args.Map{"result": lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "trim failed", actual)
	})
}

func Test_C49_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("", "b", "")
		actual := args.Map{"result": lmr.IsLeftEmpty() || lmr.IsMiddleEmpty() || !lmr.IsRightEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "empty checks failed", actual)
	})
}

func Test_C49_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")
		actual := args.Map{"result": lmr.IsLeftWhitespace() || !lmr.IsMiddleWhitespace() || !lmr.IsRightWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "whitespace checks failed", actual)
	})
}

func Test_C49_LeftMiddleRight_HasValid(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_HasValid", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_HasSafeNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.IsAll("a", "b", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_Is(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Is", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.Is("a", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C49_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()
		actual := args.Map{"result": cloned.Left != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "clone failed", actual)
	})
}

func Test_C49_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "c"}
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
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		actual := args.Map{"result": coc.IsEmpty() || coc.HasItems() || coc.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_CollectionsOfCollection_Add(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Add", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(col)
		actual := args.Map{"result": coc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AddStrings", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		actual := args.Map{"result": coc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AllIndividualItemsLength", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		actual := args.Map{"result": coc.AllIndividualItemsLength() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_CollectionsOfCollection_Items(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Items", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(col)
		actual := args.Map{"result": len(coc.Items()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_List(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_List", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		list := coc.List(0)
		actual := args.Map{"result": len(list) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_CollectionsOfCollection_ToCollection(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_ToCollection", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		col := coc.ToCollection()
		actual := args.Map{"result": col.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_CollectionsOfCollection_String(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_String", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		s := coc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_CollectionsOfCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_JSON", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		data, err := json.Marshal(coc)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		coc2 := corestr.New.CollectionsOfCollection.Cap(5)
		err = json.Unmarshal(data, coc2)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
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
		hsc := corestr.Empty.HashsetsCollection()
		actual := args.Map{"result": hsc.IsEmpty() || hsc.HasItems() || hsc.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Add", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(hs)
		actual := args.Map{"result": hsc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_AddNonNil(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddNonNil", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddNonNil(nil)
		actual := args.Map{"result": hsc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nil should not add", actual)
		hsc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))
		actual := args.Map{"result": hsc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddNonEmpty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddNonEmpty(corestr.Empty.Hashset())
		actual := args.Map{"result": hsc.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should not add", actual)
	})
}

func Test_C49_HashsetsCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Adds", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Adds(hs)
		actual := args.Map{"result": hsc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_StringsList(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_StringsList", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(hs)
		list := hsc.StringsList()
		actual := args.Map{"result": len(list) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_HasAll(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_HasAll", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(hs)
		actual := args.Map{"result": hsc.HasAll("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should have all", actual)
	})
}

func Test_C49_HashsetsCollection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_HasAll_Empty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		actual := args.Map{"result": hsc.HasAll("a")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "empty should return false", actual)
	})
}

func Test_C49_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_IsEqual", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(hs)
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(hs)
		actual := args.Map{"result": hsc1.IsEqualPtr(hsc2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be equal", actual)
	})
}

func Test_C49_HashsetsCollection_IsEqual_SameRef(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_IsEqual_SameRef", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		actual := args.Map{"result": hsc.IsEqualPtr(hsc)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "same ref", actual)
	})
}

func Test_C49_HashsetsCollection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ConcatNew", func() {
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		result := hsc1.ConcatNew(hsc2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C49_HashsetsCollection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ConcatNew_Empty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hsc.ConcatNew()
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddHashsetsCollection", func() {
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		hsc1.AddHashsetsCollection(hsc2)
		actual := args.Map{"result": hsc1.Length() != 2}
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
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		actual := args.Map{"result": hsc.LastIndex() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C49_HashsetsCollection_ListPtr(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ListPtr", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		p := hsc.ListPtr()
		actual := args.Map{"result": p == nil || len(*p) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollection_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ListDirectPtr", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		p := hsc.ListDirectPtr()
		actual := args.Map{"result": p == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C49_HashsetsCollection_String(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_String", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		s := hsc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_HashsetsCollection_String_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_String_Empty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		s := hsc.String()
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected NoElements", actual)
	})
}

func Test_C49_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Join", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		j := hsc.Join(",")
		actual := args.Map{"result": j == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C49_HashsetsCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_JSON", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := json.Marshal(hsc)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
		hsc2 := corestr.Empty.HashsetsCollection()
		err = json.Unmarshal(data, hsc2)
		actual := args.Map{"result": err}
		expected := args.Map{"result": nil}
		expected.ShouldBeEqual(t, 0, "err", actual)
	})
}

func Test_C49_HashsetsCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Serialize", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		_, err := hsc.Serialize()
		actual := args.Map{"result": err}
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
		sso := corestr.Empty.SimpleStringOnce()
		actual := args.Map{"result": sso.IsInitialized() || sso.IsDefined() || !sso.IsUninitialized() || !sso.IsInvalid()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_C49_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_SetOnUninitialized", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		err := sso.SetOnUninitialized("hello")
		actual := args.Map{"result": err != nil || sso.Value() != "hello" || !sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "set failed", actual)
		err2 := sso.SetOnUninitialized("world")
		actual := args.Map{"result": err2 == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should error on already initialized", actual)
	})
}

func Test_C49_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetSetOnce", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetSetOnce("first")
		actual := args.Map{"result": v != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected first", actual)
		v2 := sso.GetSetOnce("second")
		actual := args.Map{"result": v2 != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return first (already set)", actual)
	})
}

func Test_C49_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetOnce", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetOnce()
		actual := args.Map{"result": v != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be initialized", actual)
	})
}

func Test_C49_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetOnceFunc", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetOnceFunc(func() string { return "computed" })
		actual := args.Map{"result": v != "computed"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected computed", actual)
		v2 := sso.GetOnceFunc(func() string { return "other" })
		actual := args.Map{"result": v2 != "computed"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should return first value", actual)
	})
}

func Test_C49_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Invalidate", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		sso.SetOnUninitialized("x")
		sso.Invalidate()
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be uninitialized", actual)
	})
}

func Test_C49_SimpleStringOnce_Reset(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Reset", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		sso.SetOnUninitialized("x")
		sso.Reset()
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "should be uninitialized", actual)
	})
}

func Test_C49_SimpleStringOnce_Boolean(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Boolean", func() {
		sso := corestr.New.SimpleStringOnce.Init("true")
		actual := args.Map{"result": sso.Boolean(false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		sso2 := corestr.New.SimpleStringOnce.Init("yes")
		actual := args.Map{"result": sso2.Boolean(false)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for yes", actual)
	})
}

func Test_C49_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Int", func() {
		sso := corestr.New.SimpleStringOnce.Init("42")
		actual := args.Map{"result": sso.Int() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_C49_SimpleStringOnce_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_IsEmpty", func() {
		sso := corestr.Empty.SimpleStringOnce()
		actual := args.Map{"result": sso.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C49_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_ConcatNew", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		result := sso.ConcatNew(" world")
		actual := args.Map{"result": result.Value() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)
	})
}

func Test_C49_SimpleStringOnce_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_ConcatNewUsingStrings", func() {
		sso := corestr.New.SimpleStringOnce.Init("a")
		result := sso.ConcatNewUsingStrings(",", "b", "c")
		actual := args.Map{"result": result.Value() != "a,b,c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b,c", actual)
	})
}

func Test_C49_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_WithinRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("50")
		val, inRange := sso.WithinRange(true, 0, 100)
		actual := args.Map{"result": inRange || val != 50}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected in range", actual)
		val2, inRange2 := sso.WithinRange(true, 60, 100)
		actual := args.Map{"result": inRange2 || val2 != 60}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary min", actual)
		sso3 := corestr.New.SimpleStringOnce.Init("200")
		val3, inRange3 := sso3.WithinRange(true, 0, 100)
		actual := args.Map{"result": inRange3 || val3 != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary max", actual)
	})
}

func Test_C49_SimpleStringOnce_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_WithinRange_NoBoundary", func() {
		sso := corestr.New.SimpleStringOnce.Init("200")
		val, inRange := sso.WithinRange(false, 0, 100)
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		_ = val
	})
}

// ===================== DataModels =====================

func Test_C49_HashmapDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashmapDataModel", func() {
		hm := corestr.New.Hashmap.Cap(5)
		hm.AddOrUpdate("k", "v")
		dm := corestr.NewHashmapsDataModelUsing(hm)
		hm2 := corestr.NewHashmapUsingDataModel(dm)
		actual := args.Map{"result": hm2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashsetDataModel", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		dm := corestr.NewHashsetsDataModelUsing(hs)
		hs2 := corestr.NewHashsetUsingDataModel(dm)
		actual := args.Map{"result": hs2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C49_HashsetsCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollectionDataModel", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)
		actual := args.Map{"result": hsc2.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

// ===================== AddsUsingProcessorAsync =====================

func Test_C49_LinkedCollections_AddsUsingProcessorAsync(t *testing.T) {
	safeTest(t, "Test_C49_LinkedCollections_AddsUsingProcessorAsync", func() {
		lc := corestr.New.LinkedCollection.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		processor := func(any any, index int) *corestr.Collection {
			return corestr.New.Collection.Strings([]string{any.(string)})
		}
		lc.AddsUsingProcessorAsync(wg, processor, true, "hello")
		wg.Wait()
		actual := args.Map{"result": lc.LengthLock() != 1}
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
		rb := corestr.ReturningBool{IsBreak: true, IsKeep: false}
		actual := args.Map{"result": rb.IsBreak || rb.IsKeep}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "unexpected", actual)
	})
}
