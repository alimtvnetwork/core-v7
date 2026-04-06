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
		v := corestr.NewValidValue("hello")
		actual := args.Map{"result": v.Value != "hello" || !v.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello, valid", actual)
	})
}

func Test_C27_02_ValidValue_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_02_ValidValue_NewValidValueEmpty", func() {
		v := corestr.NewValidValueEmpty()
		actual := args.Map{"result": v.Value != "" || !v.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty, valid", actual)
	})
}

func Test_C27_03_ValidValue_InvalidValidValue(t *testing.T) {
	safeTest(t, "Test_C27_03_ValidValue_InvalidValidValue", func() {
		v := corestr.InvalidValidValue("err")
		actual := args.Map{"result": v.IsValid || v.Message != "err"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid with err", actual)
	})
}

func Test_C27_04_ValidValue_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_04_ValidValue_InvalidValidValueNoMessage", func() {
		v := corestr.InvalidValidValueNoMessage()
		actual := args.Map{"result": v.IsValid || v.Message != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid no message", actual)
	})
}

func Test_C27_05_ValidValue_NewValidValueUsingAny(t *testing.T) {
	safeTest(t, "Test_C27_05_ValidValue_NewValidValueUsingAny", func() {
		v := corestr.NewValidValueUsingAny(false, true, 42)
		actual := args.Map{"result": v.IsValid || v.Value == ""}
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
		v := corestr.NewValidValue("hello")
		b := v.ValueBytesOnce()
		actual := args.Map{"result": string(b) != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		// second call should return same
		b2 := v.ValueBytesOnce()
		actual := args.Map{"result": string(b2) != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected cached hello", actual)
	})
}

func Test_C27_08_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C27_08_ValidValue_ValueBytesOncePtr", func() {
		v := corestr.NewValidValue("hi")
		b := v.ValueBytesOncePtr()
		actual := args.Map{"result": string(b) != "hi"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}

func Test_C27_09_ValidValue_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_09_ValidValue_IsEmpty", func() {
		v := corestr.NewValidValueEmpty()
		actual := args.Map{"result": v.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		v2 := corestr.NewValidValue("x")
		actual := args.Map{"result": v2.IsEmpty()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_C27_10_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_10_ValidValue_IsWhitespace", func() {
		v := corestr.NewValidValue("   ")
		actual := args.Map{"result": v.IsWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
	})
}

func Test_C27_11_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C27_11_ValidValue_Trim", func() {
		v := corestr.NewValidValue("  hello  ")
		actual := args.Map{"result": v.Trim() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_12_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_12_ValidValue_HasValidNonEmpty", func() {
		v := corestr.NewValidValue("x")
		actual := args.Map{"result": v.HasValidNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_13_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_13_ValidValue_HasValidNonWhitespace", func() {
		v := corestr.NewValidValue("x")
		actual := args.Map{"result": v.HasValidNonWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_14_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_14_ValidValue_ValueBool", func() {
		v := corestr.NewValidValue("true")
		actual := args.Map{"result": v.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		v2 := corestr.NewValidValue("nope")
		actual := args.Map{"result": v2.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		v3 := corestr.NewValidValue("")
		actual := args.Map{"result": v3.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_C27_15_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C27_15_ValidValue_ValueInt", func() {
		v := corestr.NewValidValue("42")
		actual := args.Map{"result": v.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		v2 := corestr.NewValidValue("bad")
		actual := args.Map{"result": v2.ValueInt(99) != 99}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default 99", actual)
	})
}

func Test_C27_16_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C27_16_ValidValue_ValueDefInt", func() {
		v := corestr.NewValidValue("10")
		actual := args.Map{"result": v.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C27_17_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_17_ValidValue_ValueByte", func() {
		v := corestr.NewValidValue("200")
		actual := args.Map{"result": v.ValueByte(0) != 200}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
		v2 := corestr.NewValidValue("300")
		actual := args.Map{"result": v2.ValueByte(0) != 255}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 255 for overflow", actual)
		v3 := corestr.NewValidValue("-1")
		actual := args.Map{"result": v3.ValueByte(0) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for negative", actual)
	})
}

func Test_C27_18_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C27_18_ValidValue_ValueDefByte", func() {
		v := corestr.NewValidValue("100")
		actual := args.Map{"result": v.ValueDefByte() != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_C27_19_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_19_ValidValue_ValueFloat64", func() {
		v := corestr.NewValidValue("3.14")
		actual := args.Map{"result": v.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		v2 := corestr.NewValidValue("bad")
		actual := args.Map{"result": v2.ValueFloat64(1.0) != 1.0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected default", actual)
	})
}

func Test_C27_20_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_20_ValidValue_ValueDefFloat64", func() {
		v := corestr.NewValidValue("2.5")
		actual := args.Map{"result": v.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C27_21_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_21_ValidValue_HasSafeNonEmpty", func() {
		v := corestr.NewValidValue("x")
		actual := args.Map{"result": v.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_22_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C27_22_ValidValue_Is", func() {
		v := corestr.NewValidValue("hello")
		actual := args.Map{"result": v.Is("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_23_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C27_23_ValidValue_IsAnyOf", func() {
		v := corestr.NewValidValue("b")
		actual := args.Map{"result": v.IsAnyOf("a", "b", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": v.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty values", actual)
		actual := args.Map{"result": v.IsAnyOf("x", "y")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_24_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C27_24_ValidValue_IsContains", func() {
		v := corestr.NewValidValue("hello world")
		actual := args.Map{"result": v.IsContains("world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_25_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C27_25_ValidValue_IsAnyContains", func() {
		v := corestr.NewValidValue("hello world")
		actual := args.Map{"result": v.IsAnyContains("xyz", "world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": v.IsAnyContains()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
		actual := args.Map{"result": v.IsAnyContains("xyz", "abc")}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_26_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C27_26_ValidValue_IsEqualNonSensitive", func() {
		v := corestr.NewValidValue("Hello")
		actual := args.Map{"result": v.IsEqualNonSensitive("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_27_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C27_27_ValidValue_IsRegexMatches", func() {
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": v.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": v.IsRegexMatches(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil regex", actual)
	})
}

func Test_C27_28_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C27_28_ValidValue_RegexFindString", func() {
		v := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": v.RegexFindString(re) != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual := args.Map{"result": v.RegexFindString(nil) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
	})
}

func Test_C27_29_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C27_29_ValidValue_RegexFindAllStringsWithFlag", func() {
		v := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, has := v.RegexFindAllStringsWithFlag(re, -1)
		actual := args.Map{"result": has || len(items) != 3}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
		_, has2 := v.RegexFindAllStringsWithFlag(nil, -1)
		actual := args.Map{"result": has2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
	})
}

func Test_C27_30_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C27_30_ValidValue_RegexFindAllStrings", func() {
		v := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items := v.RegexFindAllStrings(re, -1)
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		items2 := v.RegexFindAllStrings(nil, -1)
		actual := args.Map{"result": len(items2) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_31_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C27_31_ValidValue_Split", func() {
		v := corestr.NewValidValue("a,b,c")
		s := v.Split(",")
		actual := args.Map{"result": len(s) != 3}
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
		v := corestr.NewValidValue("hello")
		c := v.Clone()
		actual := args.Map{"result": c.Value != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_35_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C27_35_ValidValue_Clone_Nil", func() {
		var v *corestr.ValidValue
		actual := args.Map{"result": v.Clone() != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_C27_36_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C27_36_ValidValue_String", func() {
		v := corestr.NewValidValue("hi")
		actual := args.Map{"result": v.String() != "hi"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hi", actual)
	})
}

func Test_C27_37_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C27_37_ValidValue_String_Nil", func() {
		var v *corestr.ValidValue
		actual := args.Map{"result": v.String() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_38_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C27_38_ValidValue_FullString", func() {
		v := corestr.NewValidValue("hi")
		actual := args.Map{"result": v.FullString() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_39_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C27_39_ValidValue_FullString_Nil", func() {
		var v *corestr.ValidValue
		actual := args.Map{"result": v.FullString() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_40_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_C27_40_ValidValue_Clear", func() {
		v := corestr.NewValidValue("hi")
		v.Clear()
		actual := args.Map{"result": v.Value != "" || v.IsValid}
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
		v := corestr.NewValidValue("hi")
		actual := args.Map{"result": v.JsonPtr() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_C27_46_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C27_46_ValidValue_Serialize", func() {
		v := corestr.NewValidValue("hi")
		b, err := v.Serialize()
		actual := args.Map{"result": err != nil || len(b) == 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	})
}

func Test_C27_47_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C27_47_ValidValue_ParseInjectUsingJson", func() {
		v := corestr.NewValidValue("hi")
		jp := v.JsonPtr()
		v2 := &corestr.ValidValue{}
		_, err := v2.ParseInjectUsingJson(jp)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

func Test_C27_48_ValidValue_Deserialize(t *testing.T) {
	safeTest(t, "Test_C27_48_ValidValue_Deserialize", func() {
		v := corestr.NewValidValue("hi")
		var target corestr.ValidValue
		err := v.Deserialize(&target)
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected:", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// ValidValues
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_49_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_49_ValidValues_Empty", func() {
		vv := corestr.EmptyValidValues()
		actual := args.Map{"result": vv.IsEmpty() || vv.Length() != 0}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_50_ValidValues_NewValidValues(t *testing.T) {
	safeTest(t, "Test_C27_50_ValidValues_NewValidValues", func() {
		vv := corestr.NewValidValues(4)
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_51_ValidValues_NewValidValuesUsingValues(t *testing.T) {
	safeTest(t, "Test_C27_51_ValidValues_NewValidValuesUsingValues", func() {
		vv := corestr.NewValidValuesUsingValues(
			corestr.ValidValue{Value: "a", IsValid: true},
			corestr.ValidValue{Value: "b", IsValid: true},
		)
		actual := args.Map{"result": vv.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_52_ValidValues_NewValidValuesUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C27_52_ValidValues_NewValidValuesUsingValues_Empty", func() {
		vv := corestr.NewValidValuesUsingValues()
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_53_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C27_53_ValidValues_Add", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("hello")
		actual := args.Map{"result": vv.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_54_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C27_54_ValidValues_AddFull", func() {
		vv := corestr.NewValidValues(4)
		vv.AddFull(false, "val", "msg")
		actual := args.Map{"result": vv.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_55_ValidValues_Count(t *testing.T) {
	safeTest(t, "Test_C27_55_ValidValues_Count", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		actual := args.Map{"result": vv.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_56_ValidValues_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_C27_56_ValidValues_HasAnyItem", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		actual := args.Map{"result": vv.HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_57_ValidValues_LastIndex(t *testing.T) {
	safeTest(t, "Test_C27_57_ValidValues_LastIndex", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv.Add("b")
		actual := args.Map{"result": vv.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_58_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_C27_58_ValidValues_HasIndex", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		actual := args.Map{"result": vv.HasIndex(0)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": vv.HasIndex(5)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_59_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C27_59_ValidValues_SafeValueAt", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("hello")
		actual := args.Map{"result": vv.SafeValueAt(0) != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		actual := args.Map{"result": vv.SafeValueAt(99) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_60_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C27_60_ValidValues_SafeValidValueAt", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("hello")
		actual := args.Map{"result": vv.SafeValidValueAt(0) != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_C27_61_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_61_ValidValues_SafeValuesAtIndexes", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv.Add("b")
		vals := vv.SafeValuesAtIndexes(0, 1)
		actual := args.Map{"result": len(vals) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_62_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C27_62_ValidValues_SafeValidValuesAtIndexes", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vals := vv.SafeValidValuesAtIndexes(0)
		actual := args.Map{"result": len(vals) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_63_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C27_63_ValidValues_Strings", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		s := vv.Strings()
		actual := args.Map{"result": len(s) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_64_ValidValues_Strings_Empty(t *testing.T) {
	safeTest(t, "Test_C27_64_ValidValues_Strings_Empty", func() {
		vv := corestr.EmptyValidValues()
		s := vv.Strings()
		actual := args.Map{"result": len(s) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_65_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C27_65_ValidValues_FullStrings", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		s := vv.FullStrings()
		actual := args.Map{"result": len(s) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_66_ValidValues_FullStrings_Empty(t *testing.T) {
	safeTest(t, "Test_C27_66_ValidValues_FullStrings_Empty", func() {
		vv := corestr.EmptyValidValues()
		s := vv.FullStrings()
		actual := args.Map{"result": len(s) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_67_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C27_67_ValidValues_String", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		actual := args.Map{"result": vv.String() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_68_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C27_68_ValidValues_Find", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv.Add("b")
		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, v.Value == "b", false
		})
		actual := args.Map{"result": len(found) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_69_ValidValues_Find_Empty(t *testing.T) {
	safeTest(t, "Test_C27_69_ValidValues_Find_Empty", func() {
		vv := corestr.EmptyValidValues()
		found := vv.Find(func(i int, v *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return v, true, false
		})
		actual := args.Map{"result": len(found) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_70_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C27_70_ValidValues_ConcatNew", func() {
		vv1 := corestr.NewValidValues(4)
		vv1.Add("a")
		vv2 := corestr.NewValidValues(4)
		vv2.Add("b")
		result := vv1.ConcatNew(false, vv2)
		actual := args.Map{"result": result.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_71_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C27_71_ValidValues_ConcatNew_EmptyClone", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		result := vv.ConcatNew(true)
		actual := args.Map{"result": result.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_72_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C27_72_ValidValues_ConcatNew_EmptyNoClone", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		result := vv.ConcatNew(false)
		actual := args.Map{"result": result != vv}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same pointer", actual)
	})
}

func Test_C27_73_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C27_73_ValidValues_AddValidValues", func() {
		vv := corestr.NewValidValues(4)
		vv.Add("a")
		vv2 := corestr.NewValidValues(4)
		vv2.Add("b")
		vv.AddValidValues(vv2)
		actual := args.Map{"result": vv.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_74_ValidValues_AddValidValues_Nil(t *testing.T) {
	safeTest(t, "Test_C27_74_ValidValues_AddValidValues_Nil", func() {
		vv := corestr.NewValidValues(4)
		vv.AddValidValues(nil)
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_75_ValidValues_Adds(t *testing.T) {
	safeTest(t, "Test_C27_75_ValidValues_Adds", func() {
		vv := corestr.NewValidValues(4)
		vv.Adds(corestr.ValidValue{Value: "a"}, corestr.ValidValue{Value: "b"})
		actual := args.Map{"result": vv.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_76_ValidValues_Adds_Empty(t *testing.T) {
	safeTest(t, "Test_C27_76_ValidValues_Adds_Empty", func() {
		vv := corestr.NewValidValues(4)
		vv.Adds()
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_77_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C27_77_ValidValues_AddsPtr", func() {
		vv := corestr.NewValidValues(4)
		v := corestr.NewValidValue("a")
		vv.AddsPtr(v)
		actual := args.Map{"result": vv.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_78_ValidValues_AddsPtr_Empty(t *testing.T) {
	safeTest(t, "Test_C27_78_ValidValues_AddsPtr_Empty", func() {
		vv := corestr.NewValidValues(4)
		vv.AddsPtr()
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_79_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C27_79_ValidValues_AddHashsetMap", func() {
		vv := corestr.NewValidValues(4)
		vv.AddHashsetMap(map[string]bool{"a": true, "b": false})
		actual := args.Map{"result": vv.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_80_ValidValues_AddHashsetMap_Nil(t *testing.T) {
	safeTest(t, "Test_C27_80_ValidValues_AddHashsetMap_Nil", func() {
		vv := corestr.NewValidValues(4)
		vv.AddHashsetMap(nil)
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_81_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C27_81_ValidValues_AddHashset", func() {
		hs := corestr.New.Hashset.StringsSpreadItems("a", "b")
		vv := corestr.NewValidValues(4)
		vv.AddHashset(hs)
		actual := args.Map{"result": vv.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_82_ValidValues_AddHashset_Nil(t *testing.T) {
	safeTest(t, "Test_C27_82_ValidValues_AddHashset_Nil", func() {
		vv := corestr.NewValidValues(4)
		vv.AddHashset(nil)
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_83_ValidValues_Hashmap(t *testing.T) {
	safeTest(t, "Test_C27_83_ValidValues_Hashmap", func() {
		vv := corestr.NewValidValues(4)
		vv.AddFull(true, "k", "v")
		hm := vv.Hashmap()
		actual := args.Map{"result": hm.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_84_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C27_84_ValidValues_Map", func() {
		vv := corestr.NewValidValues(4)
		vv.AddFull(true, "k", "v")
		m := vv.Map()
		actual := args.Map{"result": len(m) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_C27_85_ValidValues_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C27_85_ValidValues_Length_Nil", func() {
		var vv *corestr.ValidValues
		actual := args.Map{"result": vv.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// ValueStatus
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_86_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_86_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		actual := args.Map{"result": vs.ValueValid.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_87_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_87_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("err")
		actual := args.Map{"result": vs.ValueValid.Message != "err"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected err", actual)
	})
}

func Test_C27_88_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C27_88_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("err")
		c := vs.Clone()
		actual := args.Map{"result": c.ValueValid.Message != "err"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected err", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// LeftRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_89_LeftRight_NewLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_89_LeftRight_NewLeftRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b" || !lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, valid", actual)
	})
}

func Test_C27_90_LeftRight_InvalidLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_90_LeftRight_InvalidLeftRight", func() {
		lr := corestr.InvalidLeftRight("err")
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_91_LeftRight_InvalidLeftRightNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_91_LeftRight_InvalidLeftRightNoMessage", func() {
		lr := corestr.InvalidLeftRightNoMessage()
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_92_LeftRight_LeftRightUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_92_LeftRight_LeftRightUsingSlice", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a", "b"})
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b", actual)
	})
}

func Test_C27_93_LeftRight_LeftRightUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_93_LeftRight_LeftRightUsingSlice_Empty", func() {
		lr := corestr.LeftRightUsingSlice([]string{})
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_94_LeftRight_LeftRightUsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_94_LeftRight_LeftRightUsingSlice_Single", func() {
		lr := corestr.LeftRightUsingSlice([]string{"a"})
		actual := args.Map{"result": lr.Left != "a" || lr.Right != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, empty", actual)
	})
}

func Test_C27_95_LeftRight_LeftRightUsingSlicePtr(t *testing.T) {
	safeTest(t, "Test_C27_95_LeftRight_LeftRightUsingSlicePtr", func() {
		lr := corestr.LeftRightUsingSlicePtr([]string{"a", "b"})
		actual := args.Map{"result": lr.Left != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_96_LeftRight_LeftRightTrimmedUsingSlice(t *testing.T) {
	safeTest(t, "Test_C27_96_LeftRight_LeftRightTrimmedUsingSlice", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a ", " b "})
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_97_LeftRight_LeftRightTrimmedUsingSlice_Nil(t *testing.T) {
	safeTest(t, "Test_C27_97_LeftRight_LeftRightTrimmedUsingSlice_Nil", func() {
		lr := corestr.LeftRightTrimmedUsingSlice(nil)
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_98_LeftRight_LeftRightTrimmedUsingSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_98_LeftRight_LeftRightTrimmedUsingSlice_Empty", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{})
		actual := args.Map{"result": lr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single(t *testing.T) {
	safeTest(t, "Test_C27_99_LeftRight_LeftRightTrimmedUsingSlice_Single", func() {
		lr := corestr.LeftRightTrimmedUsingSlice([]string{" a "})
		actual := args.Map{"result": lr.Left != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_100_LeftRight_LeftBytes(t *testing.T) {
	safeTest(t, "Test_C27_100_LeftRight_LeftBytes", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": string(lr.LeftBytes()) != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_101_LeftRight_RightBytes(t *testing.T) {
	safeTest(t, "Test_C27_101_LeftRight_RightBytes", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": string(lr.RightBytes()) != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C27_102_LeftRight_LeftTrim(t *testing.T) {
	safeTest(t, "Test_C27_102_LeftRight_LeftTrim", func() {
		lr := corestr.NewLeftRight(" a ", "b")
		actual := args.Map{"result": lr.LeftTrim() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_103_LeftRight_RightTrim(t *testing.T) {
	safeTest(t, "Test_C27_103_LeftRight_RightTrim", func() {
		lr := corestr.NewLeftRight("a", " b ")
		actual := args.Map{"result": lr.RightTrim() != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_C27_104_LeftRight_IsLeftEmpty(t *testing.T) {
	safeTest(t, "Test_C27_104_LeftRight_IsLeftEmpty", func() {
		lr := corestr.NewLeftRight("", "b")
		actual := args.Map{"result": lr.IsLeftEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_105_LeftRight_IsRightEmpty(t *testing.T) {
	safeTest(t, "Test_C27_105_LeftRight_IsRightEmpty", func() {
		lr := corestr.NewLeftRight("a", "")
		actual := args.Map{"result": lr.IsRightEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_106_LeftRight_IsLeftWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_106_LeftRight_IsLeftWhitespace", func() {
		lr := corestr.NewLeftRight("   ", "b")
		actual := args.Map{"result": lr.IsLeftWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_107_LeftRight_IsRightWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_107_LeftRight_IsRightWhitespace", func() {
		lr := corestr.NewLeftRight("a", "   ")
		actual := args.Map{"result": lr.IsRightWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_108_LeftRight_HasValidNonEmptyLeft(t *testing.T) {
	safeTest(t, "Test_C27_108_LeftRight_HasValidNonEmptyLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.HasValidNonEmptyLeft()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_109_LeftRight_HasValidNonEmptyRight(t *testing.T) {
	safeTest(t, "Test_C27_109_LeftRight_HasValidNonEmptyRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.HasValidNonEmptyRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_110_LeftRight_HasValidNonWhitespaceLeft(t *testing.T) {
	safeTest(t, "Test_C27_110_LeftRight_HasValidNonWhitespaceLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.HasValidNonWhitespaceLeft()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_111_LeftRight_HasValidNonWhitespaceRight(t *testing.T) {
	safeTest(t, "Test_C27_111_LeftRight_HasValidNonWhitespaceRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.HasValidNonWhitespaceRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_112_LeftRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_112_LeftRight_HasSafeNonEmpty", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_113_LeftRight_NonPtr(t *testing.T) {
	safeTest(t, "Test_C27_113_LeftRight_NonPtr", func() {
		lr := corestr.NewLeftRight("a", "b")
		np := lr.NonPtr()
		actual := args.Map{"result": np.Left != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_C27_114_LeftRight_Ptr(t *testing.T) {
	safeTest(t, "Test_C27_114_LeftRight_Ptr", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.Ptr() != lr}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected same", actual)
	})
}

func Test_C27_115_LeftRight_IsLeftRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_115_LeftRight_IsLeftRegexMatch", func() {
		lr := corestr.NewLeftRight("abc123", "b")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": lr.IsLeftRegexMatch(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": lr.IsLeftRegexMatch(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_116_LeftRight_IsRightRegexMatch(t *testing.T) {
	safeTest(t, "Test_C27_116_LeftRight_IsRightRegexMatch", func() {
		lr := corestr.NewLeftRight("a", "abc123")
		re := regexp.MustCompile(`\d+`)
		actual := args.Map{"result": lr.IsRightRegexMatch(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual := args.Map{"result": lr.IsRightRegexMatch(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_117_LeftRight_IsLeft(t *testing.T) {
	safeTest(t, "Test_C27_117_LeftRight_IsLeft", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.IsLeft("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_118_LeftRight_IsRight(t *testing.T) {
	safeTest(t, "Test_C27_118_LeftRight_IsRight", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.IsRight("b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_119_LeftRight_Is(t *testing.T) {
	safeTest(t, "Test_C27_119_LeftRight_Is", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.Is("a", "b")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_120_LeftRight_IsEqual(t *testing.T) {
	safeTest(t, "Test_C27_120_LeftRight_IsEqual", func() {
		lr1 := corestr.NewLeftRight("a", "b")
		lr2 := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr1.IsEqual(lr2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_121_LeftRight_IsEqual_BothNil(t *testing.T) {
	safeTest(t, "Test_C27_121_LeftRight_IsEqual_BothNil", func() {
		var lr1, lr2 *corestr.LeftRight
		actual := args.Map{"result": lr1.IsEqual(lr2)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_122_LeftRight_IsEqual_OneNil(t *testing.T) {
	safeTest(t, "Test_C27_122_LeftRight_IsEqual_OneNil", func() {
		lr := corestr.NewLeftRight("a", "b")
		actual := args.Map{"result": lr.IsEqual(nil)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_123_LeftRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_123_LeftRight_Clone", func() {
		lr := corestr.NewLeftRight("a", "b")
		c := lr.Clone()
		actual := args.Map{"result": c.Left != "a" || c.Right != "b"}
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
		lr := corestr.LeftRightFromSplit("key=value", "=")
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key, value", actual)
	})
}

func Test_C27_129_LeftRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_129_LeftRightFromSplitTrimmed", func() {
		lr := corestr.LeftRightFromSplitTrimmed(" key = value ", "=")
		actual := args.Map{"result": lr.Left != "key" || lr.Right != "value"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed key, value", actual)
	})
}

func Test_C27_130_LeftRightFromSplitFull(t *testing.T) {
	safeTest(t, "Test_C27_130_LeftRightFromSplitFull", func() {
		lr := corestr.LeftRightFromSplitFull("a:b:c:d", ":")
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b:c:d"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b:c:d", actual)
	})
}

func Test_C27_131_LeftRightFromSplitFullTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_131_LeftRightFromSplitFullTrimmed", func() {
		lr := corestr.LeftRightFromSplitFullTrimmed(" a : b : c ", ":")
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "b : c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// LeftMiddleRight
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_132_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C27_132_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_133_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C27_133_LeftMiddleRight_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		actual := args.Map{"result": lmr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_134_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C27_134_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		actual := args.Map{"result": lmr.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
	})
}

func Test_C27_135_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_C27_135_LeftMiddleRight_Bytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": string(lmr.LeftBytes()) != "a" || string(lmr.MiddleBytes()) != "b" || string(lmr.RightBytes()) != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_136_LeftMiddleRight_Trims(t *testing.T) {
	safeTest(t, "Test_C27_136_LeftMiddleRight_Trims", func() {
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
		actual := args.Map{"result": lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_C27_137_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_137_LeftMiddleRight_IsEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("", "", "")
		actual := args.Map{"result": lmr.IsLeftEmpty() || !lmr.IsMiddleEmpty() || !lmr.IsRightEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all empty", actual)
	})
}

func Test_C27_138_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C27_138_LeftMiddleRight_IsWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")
		actual := args.Map{"result": lmr.IsLeftWhitespace() || !lmr.IsMiddleWhitespace() || !lmr.IsRightWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all whitespace", actual)
	})
}

func Test_C27_139_LeftMiddleRight_HasValid(t *testing.T) {
	safeTest(t, "Test_C27_139_LeftMiddleRight_HasValid", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all valid non-empty", actual)
		actual := args.Map{"result": lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected all non-whitespace", actual)
	})
}

func Test_C27_140_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C27_140_LeftMiddleRight_HasSafeNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_141_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_C27_141_LeftMiddleRight_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.IsAll("a", "b", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_142_LeftMiddleRight_Is(t *testing.T) {
	safeTest(t, "Test_C27_142_LeftMiddleRight_Is", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		actual := args.Map{"result": lmr.Is("a", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_143_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C27_143_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		c := lmr.Clone()
		actual := args.Map{"result": c.Left != "a" || c.Middle != "b" || c.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_144_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C27_144_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		actual := args.Map{"result": lr.Left != "a" || lr.Right != "c"}
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
		lmr := corestr.LeftMiddleRightFromSplit("a.b.c", ".")
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c", actual)
	})
}

func Test_C27_150_LeftMiddleRightFromSplitTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_150_LeftMiddleRightFromSplitTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitTrimmed(" a . b . c ", ".")
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed a, b, c", actual)
	})
}

func Test_C27_151_LeftMiddleRightFromSplitN(t *testing.T) {
	safeTest(t, "Test_C27_151_LeftMiddleRightFromSplitN", func() {
		lmr := corestr.LeftMiddleRightFromSplitN("a:b:c:d:e", ":")
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c:d:e"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a, b, c:d:e", actual)
	})
}

func Test_C27_152_LeftMiddleRightFromSplitNTrimmed(t *testing.T) {
	safeTest(t, "Test_C27_152_LeftMiddleRightFromSplitNTrimmed", func() {
		lmr := corestr.LeftMiddleRightFromSplitNTrimmed(" a : b : c : d ", ":")
		actual := args.Map{"result": lmr.Left != "a" || lmr.Middle != "b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// TextWithLineNumber
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_153_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_153_TextWithLineNumber_HasLineNumber", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		actual := args.Map{"result": tl.HasLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_154_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C27_154_TextWithLineNumber_IsInvalidLineNumber", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		actual := args.Map{"result": tl.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_155_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C27_155_TextWithLineNumber_IsInvalidLineNumber_Nil", func() {
		var tl *corestr.TextWithLineNumber
		actual := args.Map{"result": tl.IsInvalidLineNumber()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_156_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C27_156_TextWithLineNumber_Length", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		actual := args.Map{"result": tl.Length() != 5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
	})
}

func Test_C27_157_TextWithLineNumber_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C27_157_TextWithLineNumber_Length_Nil", func() {
		var tl *corestr.TextWithLineNumber
		actual := args.Map{"result": tl.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_158_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C27_158_TextWithLineNumber_IsEmpty", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": tl.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_159_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C27_159_TextWithLineNumber_IsEmpty_Nil", func() {
		var tl *corestr.TextWithLineNumber
		actual := args.Map{"result": tl.IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_C27_160_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C27_160_TextWithLineNumber_IsEmptyText", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		actual := args.Map{"result": tl.IsEmptyText()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty text", actual)
	})
}

func Test_C27_161_TextWithLineNumber_IsEmptyText_Nil(t *testing.T) {
	safeTest(t, "Test_C27_161_TextWithLineNumber_IsEmptyText_Nil", func() {
		var tl *corestr.TextWithLineNumber
		actual := args.Map{"result": tl.IsEmptyText()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_162_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C27_162_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		tl := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		actual := args.Map{"result": tl.IsEmptyTextLineBoth()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// CloneSlice, CloneSliceIf
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_163_CloneSlice_Basic(t *testing.T) {
	safeTest(t, "Test_C27_163_CloneSlice_Basic", func() {
		s := corestr.CloneSlice([]string{"a", "b"})
		actual := args.Map{"result": len(s) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_164_CloneSlice_Empty(t *testing.T) {
	safeTest(t, "Test_C27_164_CloneSlice_Empty", func() {
		s := corestr.CloneSlice(nil)
		actual := args.Map{"result": len(s) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C27_165_CloneSliceIf_Clone(t *testing.T) {
	safeTest(t, "Test_C27_165_CloneSliceIf_Clone", func() {
		s := corestr.CloneSliceIf(true, "a", "b")
		actual := args.Map{"result": len(s) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_166_CloneSliceIf_NoClone(t *testing.T) {
	safeTest(t, "Test_C27_166_CloneSliceIf_NoClone", func() {
		s := corestr.CloneSliceIf(false, "a", "b")
		actual := args.Map{"result": len(s) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C27_167_CloneSliceIf_Empty(t *testing.T) {
	safeTest(t, "Test_C27_167_CloneSliceIf_Empty", func() {
		s := corestr.CloneSliceIf(true)
		actual := args.Map{"result": len(s) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// AnyToString
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_168_AnyToString_WithFieldName(t *testing.T) {
	safeTest(t, "Test_C27_168_AnyToString_WithFieldName", func() {
		s := corestr.AnyToString(true, 42)
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_169_AnyToString_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_C27_169_AnyToString_WithoutFieldName", func() {
		s := corestr.AnyToString(false, 42)
		actual := args.Map{"result": s == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_C27_170_AnyToString_Empty(t *testing.T) {
	safeTest(t, "Test_C27_170_AnyToString_Empty", func() {
		s := corestr.AnyToString(false, "")
		actual := args.Map{"result": s != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_171_AllIndividualStringsOfStringsLength(t *testing.T) {
	safeTest(t, "Test_C27_171_AllIndividualStringsOfStringsLength", func() {
		items := [][]string{{"a", "b"}, {"c"}}
		l := corestr.AllIndividualStringsOfStringsLength(&items)
		actual := args.Map{"result": l != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_172_AllIndividualStringsOfStringsLength_Nil(t *testing.T) {
	safeTest(t, "Test_C27_172_AllIndividualStringsOfStringsLength_Nil", func() {
		l := corestr.AllIndividualStringsOfStringsLength(nil)
		actual := args.Map{"result": l != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// AllIndividualsLengthOfSimpleSlices
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_173_AllIndividualsLengthOfSimpleSlices(t *testing.T) {
	safeTest(t, "Test_C27_173_AllIndividualsLengthOfSimpleSlices", func() {
		s1 := corestr.New.SimpleSlice.Create([]string{"a", "b"})
		s2 := corestr.New.SimpleSlice.Create([]string{"c"})
		l := corestr.AllIndividualsLengthOfSimpleSlices(s1, s2)
		actual := args.Map{"result": l != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_C27_174_AllIndividualsLengthOfSimpleSlices_Nil(t *testing.T) {
	safeTest(t, "Test_C27_174_AllIndividualsLengthOfSimpleSlices_Nil", func() {
		l := corestr.AllIndividualsLengthOfSimpleSlices()
		actual := args.Map{"result": l != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// Utils
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_175_Utils_WrapDoubleIfMissing(t *testing.T) {
	safeTest(t, "Test_C27_175_Utils_WrapDoubleIfMissing", func() {
		u := corestr.StringUtils
		actual := args.Map{"result": u.WrapDoubleIfMissing("hello") != `"hello"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual := args.Map{"result": u.WrapDoubleIfMissing(`"hello"`) != `"hello"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected already wrapped", actual)
		actual := args.Map{"result": u.WrapDoubleIfMissing("") != `""`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty wrapped", actual)
	})
}

func Test_C27_176_Utils_WrapSingleIfMissing(t *testing.T) {
	safeTest(t, "Test_C27_176_Utils_WrapSingleIfMissing", func() {
		u := corestr.StringUtils
		actual := args.Map{"result": u.WrapSingleIfMissing("hello") != "'hello'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
		actual := args.Map{"result": u.WrapSingleIfMissing("'hello'") != "'hello'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected already wrapped", actual)
		actual := args.Map{"result": u.WrapSingleIfMissing("") != "''"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty wrapped", actual)
	})
}

func Test_C27_177_Utils_WrapDouble(t *testing.T) {
	safeTest(t, "Test_C27_177_Utils_WrapDouble", func() {
		u := corestr.StringUtils
		actual := args.Map{"result": u.WrapDouble("hi") != `"hi"`}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_C27_178_Utils_WrapSingle(t *testing.T) {
	safeTest(t, "Test_C27_178_Utils_WrapSingle", func() {
		u := corestr.StringUtils
		actual := args.Map{"result": u.WrapSingle("hi") != "'hi'"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

func Test_C27_179_Utils_WrapTilda(t *testing.T) {
	safeTest(t, "Test_C27_179_Utils_WrapTilda", func() {
		u := corestr.StringUtils
		actual := args.Map{"result": u.WrapTilda("hi") != "`hi`"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected wrapped", actual)
	})
}

// ═══════════════════════════════════════════════════════════════════════
// KeyValuePair — comprehensive
// ═══════════════════════════════════════════════════════════════════════

func Test_C27_180_KeyValuePair_Basic(t *testing.T) {
	safeTest(t, "Test_C27_180_KeyValuePair_Basic", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.KeyName() != "k" || kv.VariableName() != "k" || kv.ValueString() != "v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k, v", actual)
	})
}

func Test_C27_181_KeyValuePair_IsVariableNameEqual(t *testing.T) {
	safeTest(t, "Test_C27_181_KeyValuePair_IsVariableNameEqual", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.IsVariableNameEqual("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_182_KeyValuePair_IsValueEqual(t *testing.T) {
	safeTest(t, "Test_C27_182_KeyValuePair_IsValueEqual", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.IsValueEqual("v")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_183_KeyValuePair_ValueBool(t *testing.T) {
	safeTest(t, "Test_C27_183_KeyValuePair_ValueBool", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "true"}
		actual := args.Map{"result": kv.ValueBool()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		kv2 := corestr.KeyValuePair{Key: "k", Value: ""}
		actual := args.Map{"result": kv2.ValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C27_184_KeyValuePair_ValueInt(t *testing.T) {
	safeTest(t, "Test_C27_184_KeyValuePair_ValueInt", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "42"}
		actual := args.Map{"result": kv.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
	})
}

func Test_C27_185_KeyValuePair_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C27_185_KeyValuePair_ValueDefInt", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "10"}
		actual := args.Map{"result": kv.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
	})
}

func Test_C27_186_KeyValuePair_ValueByte(t *testing.T) {
	safeTest(t, "Test_C27_186_KeyValuePair_ValueByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "100"}
		actual := args.Map{"result": kv.ValueByte(0) != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
	})
}

func Test_C27_187_KeyValuePair_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C27_187_KeyValuePair_ValueDefByte", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "50"}
		actual := args.Map{"result": kv.ValueDefByte() != 50}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_C27_188_KeyValuePair_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C27_188_KeyValuePair_ValueFloat64", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "3.14"}
		actual := args.Map{"result": kv.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
	})
}

func Test_C27_189_KeyValuePair_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C27_189_KeyValuePair_ValueDefFloat64", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "2.5"}
		actual := args.Map{"result": kv.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
	})
}

func Test_C27_190_KeyValuePair_ValueValid(t *testing.T) {
	safeTest(t, "Test_C27_190_KeyValuePair_ValueValid", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValid()
		actual := args.Map{"result": vv.Value != "v" || !vv.IsValid}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v, valid", actual)
	})
}

func Test_C27_191_KeyValuePair_ValueValidOptions(t *testing.T) {
	safeTest(t, "Test_C27_191_KeyValuePair_ValueValidOptions", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		vv := kv.ValueValidOptions(false, "msg")
		actual := args.Map{"result": vv.IsValid || vv.Message != "msg"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false, msg", actual)
	})
}

func Test_C27_192_KeyValuePair_IsKeyEmpty(t *testing.T) {
	safeTest(t, "Test_C27_192_KeyValuePair_IsKeyEmpty", func() {
		kv := corestr.KeyValuePair{Key: "", Value: "v"}
		actual := args.Map{"result": kv.IsKeyEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_193_KeyValuePair_IsValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_193_KeyValuePair_IsValueEmpty", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: ""}
		actual := args.Map{"result": kv.IsValueEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_194_KeyValuePair_HasKey(t *testing.T) {
	safeTest(t, "Test_C27_194_KeyValuePair_HasKey", func() {
		kv := corestr.KeyValuePair{Key: "k"}
		actual := args.Map{"result": kv.HasKey()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_195_KeyValuePair_HasValue(t *testing.T) {
	safeTest(t, "Test_C27_195_KeyValuePair_HasValue", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.HasValue()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_196_KeyValuePair_IsKeyValueEmpty(t *testing.T) {
	safeTest(t, "Test_C27_196_KeyValuePair_IsKeyValueEmpty", func() {
		kv := corestr.KeyValuePair{}
		actual := args.Map{"result": kv.IsKeyValueEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_197_KeyValuePair_TrimKey(t *testing.T) {
	safeTest(t, "Test_C27_197_KeyValuePair_TrimKey", func() {
		kv := corestr.KeyValuePair{Key: " k "}
		actual := args.Map{"result": kv.TrimKey() != "k"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected k", actual)
	})
}

func Test_C27_198_KeyValuePair_TrimValue(t *testing.T) {
	safeTest(t, "Test_C27_198_KeyValuePair_TrimValue", func() {
		kv := corestr.KeyValuePair{Value: " v "}
		actual := args.Map{"result": kv.TrimValue() != "v"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected v", actual)
	})
}

func Test_C27_199_KeyValuePair_Is(t *testing.T) {
	safeTest(t, "Test_C27_199_KeyValuePair_Is", func() {
		kv := corestr.KeyValuePair{Key: "k", Value: "v"}
		actual := args.Map{"result": kv.Is("k", "v")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C27_200_KeyValuePair_IsKey(t *testing.T) {
	safeTest(t, "Test_C27_200_KeyValuePair_IsKey", func() {
		kv := corestr.KeyValuePair{Key: "k"}
		actual := args.Map{"result": kv.IsKey("k")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}
