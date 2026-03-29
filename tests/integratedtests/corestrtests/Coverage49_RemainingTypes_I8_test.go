package corestrtests

import (
	"encoding/json"
	"regexp"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
)

// ===================== ValidValue =====================

func Test_C49_ValidValue_New(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_New", func() {
		vv := corestr.NewValidValue("hello")
		if vv.Value != "hello" || !vv.IsValid {
			t.Fatal("unexpected")
		}
	})
}

func Test_C49_ValidValue_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Empty", func() {
		vv := corestr.NewValidValueEmpty()
		if !vv.IsEmpty() || !vv.IsValid {
			t.Fatal("unexpected")
		}
	})
}

func Test_C49_ValidValue_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Invalid", func() {
		vv := corestr.InvalidValidValue("err")
		if vv.IsValid || vv.Message != "err" {
			t.Fatal("unexpected")
		}
	})
}

func Test_C49_ValidValue_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_InvalidNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		if vv.IsValid {
			t.Fatal("should be invalid")
		}
	})
}

func Test_C49_ValidValue_UsingAny(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_UsingAny", func() {
		vv := corestr.NewValidValueUsingAny(false, true, "test")
		if vv.Value == "" {
			t.Fatal("expected non-empty")
		}
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
		if len(b) != 3 {
			t.Fatal("expected 3")
		}
		b2 := vv.ValueBytesOnce() // cached
		if len(b2) != 3 {
			t.Fatal("expected cached 3")
		}
	})
}

func Test_C49_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBytesOncePtr", func() {
		vv := corestr.NewValidValue("ab")
		b := vv.ValueBytesOncePtr()
		if len(b) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C49_ValidValue_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsWhitespace", func() {
		vv := corestr.NewValidValue("   ")
		if !vv.IsWhitespace() {
			t.Fatal("expected whitespace")
		}
	})
}

func Test_C49_ValidValue_Trim(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Trim", func() {
		vv := corestr.NewValidValue("  x  ")
		if vv.Trim() != "x" {
			t.Fatal("expected x")
		}
	})
}

func Test_C49_ValidValue_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasValidNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		if !vv.HasValidNonEmpty() {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_ValidValue_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasValidNonWhitespace", func() {
		vv := corestr.NewValidValue("x")
		if !vv.HasValidNonWhitespace() {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_ValidValue_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_HasSafeNonEmpty", func() {
		vv := corestr.NewValidValue("x")
		if !vv.HasSafeNonEmpty() {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_ValidValue_ValueBool(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueBool", func() {
		vv := corestr.NewValidValue("true")
		if !vv.ValueBool() {
			t.Fatal("expected true")
		}
		vv2 := corestr.NewValidValue("invalid")
		if vv2.ValueBool() {
			t.Fatal("expected false")
		}
		vv3 := corestr.NewValidValue("")
		if vv3.ValueBool() {
			t.Fatal("expected false for empty")
		}
	})
}

func Test_C49_ValidValue_ValueInt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueInt", func() {
		vv := corestr.NewValidValue("42")
		if vv.ValueInt(0) != 42 {
			t.Fatal("expected 42")
		}
		vv2 := corestr.NewValidValue("bad")
		if vv2.ValueInt(99) != 99 {
			t.Fatal("expected default")
		}
	})
}

func Test_C49_ValidValue_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefInt", func() {
		vv := corestr.NewValidValue("10")
		if vv.ValueDefInt() != 10 {
			t.Fatal("expected 10")
		}
	})
}

func Test_C49_ValidValue_ValueByte(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueByte", func() {
		vv := corestr.NewValidValue("100")
		b := vv.ValueByte(0)
		if b != 100 {
			t.Fatalf("expected 100, got %d", b)
		}
		vv2 := corestr.NewValidValue("300")
		b2 := vv2.ValueByte(0)
		if b2 != 255 {
			t.Fatal("expected 255 for overflow")
		}
		vv3 := corestr.NewValidValue("-1")
		b3 := vv3.ValueByte(0)
		if b3 != 0 {
			t.Fatal("expected 0 for negative")
		}
	})
}

func Test_C49_ValidValue_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefByte", func() {
		vv := corestr.NewValidValue("50")
		if vv.ValueDefByte() != 50 {
			t.Fatal("expected 50")
		}
	})
}

func Test_C49_ValidValue_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueFloat64", func() {
		vv := corestr.NewValidValue("3.14")
		f := vv.ValueFloat64(0)
		if f < 3.13 || f > 3.15 {
			t.Fatal("expected ~3.14")
		}
	})
}

func Test_C49_ValidValue_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ValueDefFloat64", func() {
		vv := corestr.NewValidValue("bad")
		if vv.ValueDefFloat64() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C49_ValidValue_Is(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Is", func() {
		vv := corestr.NewValidValue("x")
		if !vv.Is("x") {
			t.Fatal("should match")
		}
	})
}

func Test_C49_ValidValue_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsAnyOf", func() {
		vv := corestr.NewValidValue("b")
		if !vv.IsAnyOf("a", "b") {
			t.Fatal("should match")
		}
		if !vv.IsAnyOf() {
			t.Fatal("empty values should return true")
		}
		if vv.IsAnyOf("x") {
			t.Fatal("should not match")
		}
	})
}

func Test_C49_ValidValue_IsContains(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsContains", func() {
		vv := corestr.NewValidValue("hello world")
		if !vv.IsContains("world") {
			t.Fatal("should contain")
		}
	})
}

func Test_C49_ValidValue_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsAnyContains", func() {
		vv := corestr.NewValidValue("hello")
		if !vv.IsAnyContains("ell") {
			t.Fatal("should contain")
		}
		if !vv.IsAnyContains() {
			t.Fatal("empty should return true")
		}
	})
}

func Test_C49_ValidValue_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsEqualNonSensitive", func() {
		vv := corestr.NewValidValue("Hello")
		if !vv.IsEqualNonSensitive("hello") {
			t.Fatal("should match case-insensitive")
		}
	})
}

func Test_C49_ValidValue_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_IsRegexMatches", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		if !vv.IsRegexMatches(re) {
			t.Fatal("should match")
		}
		if vv.IsRegexMatches(nil) {
			t.Fatal("nil regex should return false")
		}
	})
}

func Test_C49_ValidValue_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindString", func() {
		vv := corestr.NewValidValue("abc123")
		re := regexp.MustCompile(`\d+`)
		if vv.RegexFindString(re) != "123" {
			t.Fatal("expected 123")
		}
		if vv.RegexFindString(nil) != "" {
			t.Fatal("nil regex should return empty")
		}
	})
}

func Test_C49_ValidValue_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindAllStrings", func() {
		vv := corestr.NewValidValue("a1b2c3")
		re := regexp.MustCompile(`\d`)
		result := vv.RegexFindAllStrings(re, -1)
		if len(result) != 3 {
			t.Fatal("expected 3")
		}
		result2 := vv.RegexFindAllStrings(nil, -1)
		if len(result2) != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C49_ValidValue_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_RegexFindAllStringsWithFlag", func() {
		vv := corestr.NewValidValue("a1b2")
		re := regexp.MustCompile(`\d`)
		items, has := vv.RegexFindAllStringsWithFlag(re, -1)
		if !has || len(items) != 2 {
			t.Fatal("expected 2")
		}
		_, has2 := vv.RegexFindAllStringsWithFlag(nil, -1)
		if has2 {
			t.Fatal("nil regex should return false")
		}
	})
}

func Test_C49_ValidValue_Split(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Split", func() {
		vv := corestr.NewValidValue("a,b,c")
		parts := vv.Split(",")
		if len(parts) != 3 {
			t.Fatal("expected 3")
		}
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
		if cloned.Value != "x" {
			t.Fatal("clone failed")
		}
	})
}

func Test_C49_ValidValue_Clone_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clone_Nil", func() {
		var vv *corestr.ValidValue
		cloned := vv.Clone()
		if cloned != nil {
			t.Fatal("nil clone should be nil")
		}
	})
}

func Test_C49_ValidValue_String(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_String", func() {
		vv := corestr.NewValidValue("test")
		if vv.String() != "test" {
			t.Fatal("expected test")
		}
	})
}

func Test_C49_ValidValue_String_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_String_Nil", func() {
		var vv *corestr.ValidValue
		if vv.String() != "" {
			t.Fatal("nil string should be empty")
		}
	})
}

func Test_C49_ValidValue_FullString(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_FullString", func() {
		vv := corestr.NewValidValue("x")
		s := vv.FullString()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C49_ValidValue_FullString_Nil(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_FullString_Nil", func() {
		var vv *corestr.ValidValue
		if vv.FullString() != "" {
			t.Fatal("nil should return empty")
		}
	})
}

func Test_C49_ValidValue_Clear(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Clear", func() {
		vv := corestr.NewValidValue("x")
		vv.Clear()
		if vv.Value != "" {
			t.Fatal("expected empty after clear")
		}
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
		if j.Error != nil {
			t.Fatal(j.Error)
		}
	})
}

func Test_C49_ValidValue_JsonPtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_JsonPtr", func() {
		vv := corestr.NewValidValue("x")
		j := vv.JsonPtr()
		if j == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_C49_ValidValue_Serialize(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_Serialize", func() {
		vv := corestr.NewValidValue("x")
		_, err := vv.Serialize()
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_C49_ValidValue_ParseInjectUsingJson(t *testing.T) {
	safeTest(t, "Test_C49_ValidValue_ParseInjectUsingJson", func() {
		vv := corestr.NewValidValue("x")
		j := vv.Json()
		vv2 := corestr.NewValidValueEmpty()
		_, err := vv2.ParseInjectUsingJson(&j)
		if err != nil {
			t.Fatal(err)
		}
	})
}

// ===================== ValidValues =====================

func Test_C49_ValidValues_New(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_New", func() {
		vvs := corestr.NewValidValues(5)
		if !vvs.IsEmpty() || vvs.HasAnyItem() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_ValidValues_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Empty", func() {
		vvs := corestr.EmptyValidValues()
		if !vvs.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_ValidValues_Add(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Add", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		if vvs.Count() != 2 || vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C49_ValidValues_AddFull(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddFull", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "x", "msg")
		if vvs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_UsingValues(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_UsingValues", func() {
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vvs := corestr.NewValidValuesUsingValues(v1)
		if vvs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_UsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_UsingValues_Empty", func() {
		vvs := corestr.NewValidValuesUsingValues()
		if !vvs.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_ValidValues_SafeValueAt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		if vvs.SafeValueAt(0) != "x" {
			t.Fatal("expected x")
		}
		if vvs.SafeValueAt(99) != "" {
			t.Fatal("expected empty for out of range")
		}
	})
}

func Test_C49_ValidValues_SafeValidValueAt(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValidValueAt", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("x")
		if vvs.SafeValidValueAt(0) != "x" {
			t.Fatal("expected x")
		}
	})
}

func Test_C49_ValidValues_SafeValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		result := vvs.SafeValuesAtIndexes(0, 2)
		if len(result) != 2 || result[0] != "a" || result[1] != "c" {
			t.Fatal("unexpected")
		}
	})
}

func Test_C49_ValidValues_SafeValidValuesAtIndexes(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_SafeValidValuesAtIndexes", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.SafeValidValuesAtIndexes(0)
		if len(result) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_Strings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Strings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.Strings()
		if len(s) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_FullStrings(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_FullStrings", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.FullStrings()
		if len(s) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_String(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_String", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		s := vvs.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C49_ValidValues_HasIndex(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_HasIndex", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		if !vvs.HasIndex(0) {
			t.Fatal("should have index 0")
		}
		if vvs.HasIndex(1) {
			t.Fatal("should not have index 1")
		}
	})
}

func Test_C49_ValidValues_LastIndex(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_LastIndex", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		if vvs.LastIndex() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_Find(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Find", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b").Add("c")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, vv.Value == "b", false
		})
		if len(found) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_Find_Break(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Find_Break", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a").Add("b")
		found := vvs.Find(func(i int, vv *corestr.ValidValue) (*corestr.ValidValue, bool, bool) {
			return vv, true, true
		})
		if len(found) != 1 {
			t.Fatal("expected 1 (break after first)")
		}
	})
}

func Test_C49_ValidValues_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew", func() {
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		result := vvs1.ConcatNew(false, vvs2)
		if result.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C49_ValidValues_ConcatNew_EmptyClone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew_EmptyClone", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.ConcatNew(true)
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_ConcatNew_EmptyNoClone(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_ConcatNew_EmptyNoClone", func() {
		vvs := corestr.NewValidValues(5)
		vvs.Add("a")
		result := vvs.ConcatNew(false)
		if result != vvs {
			t.Fatal("expected same ref")
		}
	})
}

func Test_C49_ValidValues_AddValidValues(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddValidValues", func() {
		vvs1 := corestr.NewValidValues(5)
		vvs1.Add("a")
		vvs2 := corestr.NewValidValues(5)
		vvs2.Add("b")
		vvs1.AddValidValues(vvs2)
		if vvs1.Length() != 2 {
			t.Fatal("expected 2")
		}
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
		if vvs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_AddsPtr(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddsPtr", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddsPtr(corestr.NewValidValue("a"))
		if vvs.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_AddHashsetMap(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashsetMap", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddHashsetMap(map[string]bool{"a": true, "b": false})
		if vvs.Length() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C49_ValidValues_AddHashset(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_AddHashset", func() {
		vvs := corestr.NewValidValues(5)
		hs := corestr.New.Hashset.Strings([]string{"a"})
		vvs.AddHashset(hs)
		if vvs.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if hm.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_ValidValues_Map(t *testing.T) {
	safeTest(t, "Test_C49_ValidValues_Map", func() {
		vvs := corestr.NewValidValues(5)
		vvs.AddFull(true, "key", "val")
		m := vvs.Map()
		if len(m) != 1 {
			t.Fatal("expected 1")
		}
	})
}

// ===================== TextWithLineNumber =====================

func Test_C49_TextWithLineNumber_HasLineNumber(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_HasLineNumber", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}
		if !tln.HasLineNumber() {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_TextWithLineNumber_IsInvalidLineNumber(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsInvalidLineNumber", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1}
		if !tln.IsInvalidLineNumber() {
			t.Fatal("expected invalid")
		}
	})
}

func Test_C49_TextWithLineNumber_IsInvalidLineNumber_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsInvalidLineNumber_Nil", func() {
		var tln *corestr.TextWithLineNumber
		if !tln.IsInvalidLineNumber() {
			t.Fatal("nil should be invalid")
		}
	})
}

func Test_C49_TextWithLineNumber_Length(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_Length", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 1, Text: "abc"}
		if tln.Length() != 3 {
			t.Fatal("expected 3")
		}
	})
}

func Test_C49_TextWithLineNumber_Length_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_Length_Nil", func() {
		var tln *corestr.TextWithLineNumber
		if tln.Length() != 0 {
			t.Fatal("nil length should be 0")
		}
	})
}

func Test_C49_TextWithLineNumber_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmpty", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		if !tln.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_TextWithLineNumber_IsEmpty_Nil(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmpty_Nil", func() {
		var tln *corestr.TextWithLineNumber
		if !tln.IsEmpty() {
			t.Fatal("nil should be empty")
		}
	})
}

func Test_C49_TextWithLineNumber_IsEmptyText(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmptyText", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		if !tln.IsEmptyText() {
			t.Fatal("expected empty text")
		}
	})
}

func Test_C49_TextWithLineNumber_IsEmptyTextLineBoth(t *testing.T) {
	safeTest(t, "Test_C49_TextWithLineNumber_IsEmptyTextLineBoth", func() {
		tln := &corestr.TextWithLineNumber{LineNumber: -1, Text: ""}
		if !tln.IsEmptyTextLineBoth() {
			t.Fatal("expected true")
		}
	})
}

// ===================== ValueStatus =====================

func Test_C49_ValueStatus_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_Invalid", func() {
		vs := corestr.InvalidValueStatus("msg")
		if vs.ValueValid.IsValid {
			t.Fatal("should be invalid")
		}
	})
}

func Test_C49_ValueStatus_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_InvalidNoMessage", func() {
		vs := corestr.InvalidValueStatusNoMessage()
		if vs.ValueValid.IsValid {
			t.Fatal("should be invalid")
		}
	})
}

func Test_C49_ValueStatus_Clone(t *testing.T) {
	safeTest(t, "Test_C49_ValueStatus_Clone", func() {
		vs := corestr.InvalidValueStatus("msg")
		cloned := vs.Clone()
		if cloned.ValueValid.Message != "msg" {
			t.Fatal("clone failed")
		}
	})
}

// ===================== LeftMiddleRight =====================

func Test_C49_LeftMiddleRight_New(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_New", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		if lmr.Left != "a" || lmr.Middle != "b" || lmr.Right != "c" {
			t.Fatal("unexpected")
		}
	})
}

func Test_C49_LeftMiddleRight_Invalid(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Invalid", func() {
		lmr := corestr.InvalidLeftMiddleRight("err")
		if lmr.IsValid {
			t.Fatal("should be invalid")
		}
	})
}

func Test_C49_LeftMiddleRight_InvalidNoMessage(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_InvalidNoMessage", func() {
		lmr := corestr.InvalidLeftMiddleRightNoMessage()
		if lmr.IsValid {
			t.Fatal("should be invalid")
		}
	})
}

func Test_C49_LeftMiddleRight_Bytes(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Bytes", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		if len(lmr.LeftBytes()) != 1 || len(lmr.MiddleBytes()) != 1 || len(lmr.RightBytes()) != 1 {
			t.Fatal("expected 1 byte each")
		}
	})
}

func Test_C49_LeftMiddleRight_Trim(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Trim", func() {
		lmr := corestr.NewLeftMiddleRight(" a ", " b ", " c ")
		if lmr.LeftTrim() != "a" || lmr.MiddleTrim() != "b" || lmr.RightTrim() != "c" {
			t.Fatal("trim failed")
		}
	})
}

func Test_C49_LeftMiddleRight_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("", "b", "")
		if !lmr.IsLeftEmpty() || lmr.IsMiddleEmpty() || !lmr.IsRightEmpty() {
			t.Fatal("empty checks failed")
		}
	})
}

func Test_C49_LeftMiddleRight_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsWhitespace", func() {
		lmr := corestr.NewLeftMiddleRight("  ", "  ", "  ")
		if !lmr.IsLeftWhitespace() || !lmr.IsMiddleWhitespace() || !lmr.IsRightWhitespace() {
			t.Fatal("whitespace checks failed")
		}
	})
}

func Test_C49_LeftMiddleRight_HasValid(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_HasValid", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		if !lmr.HasValidNonEmptyLeft() || !lmr.HasValidNonEmptyMiddle() || !lmr.HasValidNonEmptyRight() {
			t.Fatal("expected true")
		}
		if !lmr.HasValidNonWhitespaceLeft() || !lmr.HasValidNonWhitespaceMiddle() || !lmr.HasValidNonWhitespaceRight() {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_LeftMiddleRight_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_HasSafeNonEmpty", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		if !lmr.HasSafeNonEmpty() {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_LeftMiddleRight_IsAll(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_IsAll", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		if !lmr.IsAll("a", "b", "c") {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_LeftMiddleRight_Is(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Is", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		if !lmr.Is("a", "c") {
			t.Fatal("expected true")
		}
	})
}

func Test_C49_LeftMiddleRight_Clone(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_Clone", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		cloned := lmr.Clone()
		if cloned.Left != "a" {
			t.Fatal("clone failed")
		}
	})
}

func Test_C49_LeftMiddleRight_ToLeftRight(t *testing.T) {
	safeTest(t, "Test_C49_LeftMiddleRight_ToLeftRight", func() {
		lmr := corestr.NewLeftMiddleRight("a", "b", "c")
		lr := lmr.ToLeftRight()
		if lr.Left != "a" || lr.Right != "c" {
			t.Fatal("conversion failed")
		}
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
		if !coc.IsEmpty() || coc.HasItems() || coc.Length() != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_CollectionsOfCollection_Add(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Add", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(col)
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_CollectionsOfCollection_AddStrings(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AddStrings", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		if coc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_CollectionsOfCollection_AllIndividualItemsLength(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_AllIndividualItemsLength", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		if coc.AllIndividualItemsLength() != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C49_CollectionsOfCollection_Items(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_Items", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		col := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(col)
		if len(coc.Items()) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_CollectionsOfCollection_List(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_List", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a", "b"})
		list := coc.List(0)
		if len(list) != 2 {
			t.Fatal("expected 2")
		}
	})
}

func Test_C49_CollectionsOfCollection_ToCollection(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_ToCollection", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		col := coc.ToCollection()
		if col.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_CollectionsOfCollection_String(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_String", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		s := coc.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C49_CollectionsOfCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C49_CollectionsOfCollection_JSON", func() {
		coc := corestr.New.CollectionsOfCollection.Cap(5)
		coc.AddStrings(false, []string{"a"})
		data, err := json.Marshal(coc)
		if err != nil {
			t.Fatal(err)
		}
		coc2 := corestr.New.CollectionsOfCollection.Cap(5)
		err = json.Unmarshal(data, coc2)
		if err != nil {
			t.Fatal(err)
		}
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
		if !hsc.IsEmpty() || hsc.HasItems() || hsc.Length() != 0 {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_HashsetsCollection_Add(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Add", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(hs)
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollection_AddNonNil(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddNonNil", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddNonNil(nil)
		if hsc.Length() != 0 {
			t.Fatal("nil should not add")
		}
		hsc.AddNonNil(corestr.New.Hashset.Strings([]string{"a"}))
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollection_AddNonEmpty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddNonEmpty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.AddNonEmpty(corestr.Empty.Hashset())
		if hsc.Length() != 0 {
			t.Fatal("empty should not add")
		}
	})
}

func Test_C49_HashsetsCollection_Adds(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Adds", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Adds(hs)
		if hsc.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollection_StringsList(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_StringsList", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(hs)
		list := hsc.StringsList()
		if len(list) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollection_HasAll(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_HasAll", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hs := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(hs)
		if !hsc.HasAll("a", "b") {
			t.Fatal("should have all")
		}
	})
}

func Test_C49_HashsetsCollection_HasAll_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_HasAll_Empty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		if hsc.HasAll("a") {
			t.Fatal("empty should return false")
		}
	})
}

func Test_C49_HashsetsCollection_IsEqual(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_IsEqual", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(hs)
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(hs)
		if !hsc1.IsEqualPtr(hsc2) {
			t.Fatal("should be equal")
		}
	})
}

func Test_C49_HashsetsCollection_IsEqual_SameRef(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_IsEqual_SameRef", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		if !hsc.IsEqualPtr(hsc) {
			t.Fatal("same ref")
		}
	})
}

func Test_C49_HashsetsCollection_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ConcatNew", func() {
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		result := hsc1.ConcatNew(hsc2)
		if result.Length() != 2 {
			t.Fatalf("expected 2, got %d", result.Length())
		}
	})
}

func Test_C49_HashsetsCollection_ConcatNew_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ConcatNew_Empty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		result := hsc.ConcatNew()
		if result.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollection_AddHashsetsCollection(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_AddHashsetsCollection", func() {
		hsc1 := corestr.Empty.HashsetsCollection()
		hsc1.Add(corestr.New.Hashset.Strings([]string{"a"}))
		hsc2 := corestr.Empty.HashsetsCollection()
		hsc2.Add(corestr.New.Hashset.Strings([]string{"b"}))
		hsc1.AddHashsetsCollection(hsc2)
		if hsc1.Length() != 2 {
			t.Fatal("expected 2")
		}
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
		if hsc.LastIndex() != 0 {
			t.Fatal("expected 0")
		}
	})
}

func Test_C49_HashsetsCollection_ListPtr(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ListPtr", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		p := hsc.ListPtr()
		if p == nil || len(*p) != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollection_ListDirectPtr(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_ListDirectPtr", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		p := hsc.ListDirectPtr()
		if p == nil {
			t.Fatal("expected non-nil")
		}
	})
}

func Test_C49_HashsetsCollection_String(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_String", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		s := hsc.String()
		if s == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C49_HashsetsCollection_String_Empty(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_String_Empty", func() {
		hsc := corestr.Empty.HashsetsCollection()
		s := hsc.String()
		if s == "" {
			t.Fatal("expected NoElements")
		}
	})
}

func Test_C49_HashsetsCollection_Join(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Join", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		j := hsc.Join(",")
		if j == "" {
			t.Fatal("expected non-empty")
		}
	})
}

func Test_C49_HashsetsCollection_JSON(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_JSON", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		data, err := json.Marshal(hsc)
		if err != nil {
			t.Fatal(err)
		}
		hsc2 := corestr.Empty.HashsetsCollection()
		err = json.Unmarshal(data, hsc2)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func Test_C49_HashsetsCollection_Serialize(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollection_Serialize", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		_, err := hsc.Serialize()
		if err != nil {
			t.Fatal(err)
		}
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
		if sso.IsInitialized() || sso.IsDefined() || !sso.IsUninitialized() || !sso.IsInvalid() {
			t.Fatal("expected uninitialized")
		}
	})
}

func Test_C49_SimpleStringOnce_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_SetOnUninitialized", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		err := sso.SetOnUninitialized("hello")
		if err != nil || sso.Value() != "hello" || !sso.IsInitialized() {
			t.Fatal("set failed")
		}
		err2 := sso.SetOnUninitialized("world")
		if err2 == nil {
			t.Fatal("should error on already initialized")
		}
	})
}

func Test_C49_SimpleStringOnce_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetSetOnce", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetSetOnce("first")
		if v != "first" {
			t.Fatal("expected first")
		}
		v2 := sso.GetSetOnce("second")
		if v2 != "first" {
			t.Fatal("should return first (already set)")
		}
	})
}

func Test_C49_SimpleStringOnce_GetOnce(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetOnce", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetOnce()
		if v != "" {
			t.Fatal("expected empty")
		}
		if !sso.IsInitialized() {
			t.Fatal("should be initialized")
		}
	})
}

func Test_C49_SimpleStringOnce_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_GetOnceFunc", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		v := sso.GetOnceFunc(func() string { return "computed" })
		if v != "computed" {
			t.Fatal("expected computed")
		}
		v2 := sso.GetOnceFunc(func() string { return "other" })
		if v2 != "computed" {
			t.Fatal("should return first value")
		}
	})
}

func Test_C49_SimpleStringOnce_Invalidate(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Invalidate", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		sso.SetOnUninitialized("x")
		sso.Invalidate()
		if sso.IsInitialized() {
			t.Fatal("should be uninitialized")
		}
	})
}

func Test_C49_SimpleStringOnce_Reset(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Reset", func() {
		sso := corestr.Empty.SimpleStringOncePtr()
		sso.SetOnUninitialized("x")
		sso.Reset()
		if sso.IsInitialized() {
			t.Fatal("should be uninitialized")
		}
	})
}

func Test_C49_SimpleStringOnce_Boolean(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Boolean", func() {
		sso := corestr.New.SimpleStringOnce.Init("true")
		if !sso.Boolean(false) {
			t.Fatal("expected true")
		}
		sso2 := corestr.New.SimpleStringOnce.Init("yes")
		if !sso2.Boolean(false) {
			t.Fatal("expected true for yes")
		}
	})
}

func Test_C49_SimpleStringOnce_Int(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_Int", func() {
		sso := corestr.New.SimpleStringOnce.Init("42")
		if sso.Int() != 42 {
			t.Fatal("expected 42")
		}
	})
}

func Test_C49_SimpleStringOnce_IsEmpty(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_IsEmpty", func() {
		sso := corestr.Empty.SimpleStringOnce()
		if !sso.IsEmpty() {
			t.Fatal("expected empty")
		}
	})
}

func Test_C49_SimpleStringOnce_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_ConcatNew", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		result := sso.ConcatNew(" world")
		if result.Value() != "hello world" {
			t.Fatal("expected hello world")
		}
	})
}

func Test_C49_SimpleStringOnce_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_ConcatNewUsingStrings", func() {
		sso := corestr.New.SimpleStringOnce.Init("a")
		result := sso.ConcatNewUsingStrings(",", "b", "c")
		if result.Value() != "a,b,c" {
			t.Fatalf("expected a,b,c got %s", result.Value())
		}
	})
}

func Test_C49_SimpleStringOnce_WithinRange(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_WithinRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("50")
		val, inRange := sso.WithinRange(true, 0, 100)
		if !inRange || val != 50 {
			t.Fatal("expected in range")
		}
		val2, inRange2 := sso.WithinRange(true, 60, 100)
		if inRange2 || val2 != 60 {
			t.Fatal("expected boundary min")
		}
		sso3 := corestr.New.SimpleStringOnce.Init("200")
		val3, inRange3 := sso3.WithinRange(true, 0, 100)
		if inRange3 || val3 != 100 {
			t.Fatal("expected boundary max")
		}
	})
}

func Test_C49_SimpleStringOnce_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_C49_SimpleStringOnce_WithinRange_NoBoundary", func() {
		sso := corestr.New.SimpleStringOnce.Init("200")
		val, inRange := sso.WithinRange(false, 0, 100)
		if inRange {
			t.Fatal("expected out of range")
		}
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
		if hm2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashsetDataModel", func() {
		hs := corestr.New.Hashset.Strings([]string{"a"})
		dm := corestr.NewHashsetsDataModelUsing(hs)
		hs2 := corestr.NewHashsetUsingDataModel(dm)
		if hs2.Length() != 1 {
			t.Fatal("expected 1")
		}
	})
}

func Test_C49_HashsetsCollectionDataModel(t *testing.T) {
	safeTest(t, "Test_C49_HashsetsCollectionDataModel", func() {
		hsc := corestr.Empty.HashsetsCollection()
		hsc.Add(corestr.New.Hashset.Strings([]string{"a"}))
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)
		hsc2 := corestr.NewHashsetsCollectionUsingDataModel(dm)
		if hsc2.Length() != 1 {
			t.Fatal("expected 1")
		}
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
		if lc.LengthLock() != 1 {
			t.Fatal("expected 1")
		}
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
		if !rb.IsBreak || rb.IsKeep {
			t.Fatal("unexpected")
		}
	})
}
