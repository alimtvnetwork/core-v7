package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/issetter"
)

// ========================================
// S17: SimpleStringOnce core methods
//   Value, Set, Get, numeric conversions,
//   Boolean, IsSetter, comparison, state
// ========================================

func Test_C77_SSO_Value_IsInitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Value_IsInitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act & Assert
		if sso.Value() != "hello" {
			t.Errorf("expected 'hello', got '%s'", sso.Value())
		}
		if !sso.IsInitialized() {
			t.Error("expected initialized")
		}
	})
}

func Test_C77_SSO_IsDefined(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsDefined", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		if !sso.IsDefined() {
			t.Error("expected defined")
		}
	})
}

func Test_C77_SSO_IsUninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act & Assert
		if !sso.IsUninitialized() {
			t.Error("expected uninitialized")
		}
	})
}

func Test_C77_SSO_Invalidate(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Invalidate", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		sso.Invalidate()

		// Assert
		if sso.IsInitialized() || sso.Value() != "" {
			t.Error("expected invalidated")
		}
	})
}

func Test_C77_SSO_Reset(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Reset", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		sso.Reset()

		// Assert
		if sso.IsInitialized() {
			t.Error("expected reset")
		}
	})
}

func Test_C77_SSO_IsInvalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsInvalid", func() {
		// Arrange
		uninit := corestr.New.SimpleStringOnce.Empty()
		initEmpty := corestr.New.SimpleStringOnce.Init("")
		valid := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		if !uninit.IsInvalid() {
			t.Error("expected invalid for uninitialized")
		}
		if !initEmpty.IsInvalid() {
			t.Error("expected invalid for empty value")
		}
		if valid.IsInvalid() {
			t.Error("expected valid")
		}
	})
}

func Test_C77_SSO_IsInvalid_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsInvalid_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act & Assert
		if !sso.IsInvalid() {
			t.Error("expected invalid for nil")
		}
	})
}

func Test_C77_SSO_ValueBytes(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueBytes", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		result := sso.ValueBytes()

		// Assert
		if string(result) != "abc" {
			t.Error("bytes mismatch")
		}
	})
}

func Test_C77_SSO_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueBytesPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		result := sso.ValueBytesPtr()

		// Assert
		if string(result) != "xyz" {
			t.Error("bytes mismatch")
		}
	})
}

func Test_C77_SSO_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		err := sso.SetOnUninitialized("val")

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if sso.Value() != "val" {
			t.Error("value not set")
		}
	})
}

func Test_C77_SSO_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		err := sso.SetOnUninitialized("new")

		// Assert
		if err == nil {
			t.Error("expected error for already initialized")
		}
		if sso.Value() != "existing" {
			t.Error("value should not change")
		}
	})
}

func Test_C77_SSO_GetSetOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetSetOnce_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetSetOnce("first")

		// Assert
		if result != "first" {
			t.Errorf("expected 'first', got '%s'", result)
		}
	})
}

func Test_C77_SSO_GetSetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetSetOnce_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		result := sso.GetSetOnce("new")

		// Assert
		if result != "existing" {
			t.Errorf("expected 'existing', got '%s'", result)
		}
	})
}

func Test_C77_SSO_GetOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnce_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetOnce()

		// Assert
		if result != "" {
			t.Error("expected empty string")
		}
		if !sso.IsInitialized() {
			t.Error("should be initialized after GetOnce")
		}
	})
}

func Test_C77_SSO_GetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnce_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		result := sso.GetOnce()

		// Assert
		if result != "val" {
			t.Errorf("expected 'val'")
		}
	})
}

func Test_C77_SSO_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnceFunc", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetOnceFunc(func() string { return "computed" })

		// Assert
		if result != "computed" {
			t.Errorf("expected 'computed', got '%s'", result)
		}
	})
}

func Test_C77_SSO_GetOnceFunc_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnceFunc_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		result := sso.GetOnceFunc(func() string { return "new" })

		// Assert
		if result != "existing" {
			t.Errorf("expected 'existing'")
		}
	})
}

func Test_C77_SSO_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnceIfUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		isSet := sso.SetOnceIfUninitialized("val")

		// Assert
		if !isSet {
			t.Error("expected true")
		}
	})
}

func Test_C77_SSO_SetOnceIfUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnceIfUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		isSet := sso.SetOnceIfUninitialized("new")

		// Assert
		if isSet {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetInitialize_SetUnInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		sso.SetInitialize()

		// Assert
		if !sso.IsInitialized() {
			t.Error("expected initialized")
		}

		// Act
		sso.SetUnInit()

		// Assert
		if sso.IsInitialized() {
			t.Error("expected uninitialized")
		}
	})
}

func Test_C77_SSO_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ConcatNew", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		result := sso.ConcatNew(" world")

		// Assert
		if result.Value() != "hello world" {
			t.Errorf("expected 'hello world', got '%s'", result.Value())
		}
	})
}

func Test_C77_SSO_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ConcatNewUsingStrings", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a")

		// Act
		result := sso.ConcatNewUsingStrings("-", "b", "c")

		// Assert
		if result.Value() != "a-b-c" {
			t.Errorf("expected 'a-b-c', got '%s'", result.Value())
		}
	})
}

func Test_C77_SSO_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsEmpty_IsWhitespace", func() {
		// Arrange
		empty := corestr.New.SimpleStringOnce.Init("")
		ws := corestr.New.SimpleStringOnce.Init("  ")
		val := corestr.New.SimpleStringOnce.Init("x")

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

func Test_C77_SSO_Trim(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Trim", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" hello ")

		// Act & Assert
		if sso.Trim() != "hello" {
			t.Error("trim mismatch")
		}
	})
}

func Test_C77_SSO_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_HasValidNonEmpty", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		empty := corestr.New.SimpleStringOnce.Init("")

		// Act & Assert
		if !valid.HasValidNonEmpty() {
			t.Error("expected true")
		}
		if empty.HasValidNonEmpty() {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C77_SSO_HasValidNonWhitespace", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		ws := corestr.New.SimpleStringOnce.Init("  ")

		// Act & Assert
		if !valid.HasValidNonWhitespace() {
			t.Error("expected true")
		}
		if ws.HasValidNonWhitespace() {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_IsValueBool(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsValueBool", func() {
		// Arrange
		ssoFalse := corestr.New.SimpleStringOnce.Init("false")
		ssoTrue := corestr.New.SimpleStringOnce.Init("true")

		// Act & Assert
		if ssoFalse.IsValueBool() {
			t.Error("expected false")
		}
		if !ssoTrue.IsValueBool() {
			t.Error("expected true")
		}
	})
}

func Test_C77_SSO_SafeValue(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SafeValue", func() {
		// Arrange
		init := corestr.New.SimpleStringOnce.Init("val")
		uninit := corestr.New.SimpleStringOnce.Empty()

		// Act & Assert
		if init.SafeValue() != "val" {
			t.Error("expected 'val'")
		}
		if uninit.SafeValue() != "" {
			t.Error("expected empty for uninitialized")
		}
	})
}

func Test_C77_SSO_Int(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Int", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if sso.Int() != 42 {
			t.Error("expected 42")
		}
		if invalid.Int() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C77_SSO_Byte(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Byte", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("200")
		overflow := corestr.New.SimpleStringOnce.Init("300")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if valid.Byte() != 200 {
			t.Error("expected 200")
		}
		if overflow.Byte() != 0 {
			t.Error("expected 0 for overflow")
		}
		if invalid.Byte() != 0 {
			t.Error("expected 0 for invalid")
		}
	})
}

func Test_C77_SSO_Int16(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Int16", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("1000")
		overflow := corestr.New.SimpleStringOnce.Init("40000")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if valid.Int16() != 1000 {
			t.Error("expected 1000")
		}
		if overflow.Int16() != 0 {
			t.Error("expected 0 for overflow")
		}
		if invalid.Int16() != 0 {
			t.Error("expected 0 for invalid")
		}
	})
}

func Test_C77_SSO_Int32(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Int32", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("100000")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if valid.Int32() != 100000 {
			t.Error("expected 100000")
		}
		if invalid.Int32() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C77_SSO_Uint16(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Uint16", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("500")

		// Act
		val, inRange := valid.Uint16()

		// Assert
		if !inRange || val != 500 {
			t.Errorf("expected 500 in range, got %d, %v", val, inRange)
		}
	})
}

func Test_C77_SSO_Uint32(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Uint32", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("100000")

		// Act
		val, inRange := valid.Uint32()

		// Assert
		if !inRange || val != 100000 {
			t.Errorf("expected 100000, got %d", val)
		}
	})
}

func Test_C77_SSO_WithinRange_InRange(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_InRange", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		if !inRange || val != 50 {
			t.Error("expected 50 in range")
		}
	})
}

func Test_C77_SSO_WithinRange_BelowMin_WithBoundary(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_BelowMin_WithBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("-5")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		if inRange {
			t.Error("expected out of range")
		}
		if val != 0 {
			t.Errorf("expected boundary min 0, got %d", val)
		}
	})
}

func Test_C77_SSO_WithinRange_AboveMax_WithBoundary(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_AboveMax_WithBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		if inRange {
			t.Error("expected out of range")
		}
		if val != 100 {
			t.Errorf("expected boundary max 100, got %d", val)
		}
	})
}

func Test_C77_SSO_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_NoBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")

		// Act
		val, inRange := sso.WithinRange(false, 0, 100)

		// Assert
		if inRange {
			t.Error("expected out of range")
		}
		if val != 200 {
			t.Errorf("expected raw value 200, got %d", val)
		}
	})
}

func Test_C77_SSO_WithinRange_Invalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		if inRange {
			t.Error("expected false for invalid")
		}
		if val != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C77_SSO_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRangeDefault", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")

		// Act
		val, inRange := sso.WithinRangeDefault(0, 100)

		// Assert
		if !inRange || val != 50 {
			t.Error("expected 50 in range")
		}
	})
}

func Test_C77_SSO_Boolean_True_Values(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_True_Values", func() {
		// Arrange
		tests := []string{"true", "yes", "y", "1", "YES", "Y"}

		for _, v := range tests {
			sso := corestr.New.SimpleStringOnce.Init(v)

			// Act & Assert
			if !sso.Boolean(false) {
				t.Errorf("expected true for '%s'", v)
			}
		}
	})
}

func Test_C77_SSO_Boolean_False_Values(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_False_Values", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("false")

		// Act & Assert
		if sso.Boolean(false) {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_Boolean_Invalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act & Assert
		if sso.Boolean(false) {
			t.Error("expected false for invalid")
		}
	})
}

func Test_C77_SSO_Boolean_ConsiderInit_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_ConsiderInit_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")

		// Act & Assert
		if sso.Boolean(true) {
			t.Error("expected false for uninitialized with considerInit")
		}
	})
}

func Test_C77_SSO_BooleanDefault(t *testing.T) {
	safeTest(t, "Test_C77_SSO_BooleanDefault", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("yes")

		// Act & Assert
		if !sso.BooleanDefault() {
			t.Error("expected true")
		}
	})
}

func Test_C77_SSO_IsSetter_True(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_True", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("yes")

		// Act
		result := sso.IsSetter(false)

		// Assert
		if result != issetter.True {
			t.Errorf("expected True, got %v", result)
		}
	})
}

func Test_C77_SSO_IsSetter_False(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_False", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("false")

		// Act
		result := sso.IsSetter(false)

		// Assert
		if result != issetter.False {
			t.Errorf("expected False, got %v", result)
		}
	})
}

func Test_C77_SSO_IsSetter_Invalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		result := sso.IsSetter(false)

		// Assert
		if result != issetter.Uninitialized {
			t.Errorf("expected Uninitialized, got %v", result)
		}
	})
}

func Test_C77_SSO_IsSetter_ConsiderInit_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_ConsiderInit_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")

		// Act
		result := sso.IsSetter(true)

		// Assert
		if result != issetter.False {
			t.Errorf("expected False, got %v", result)
		}
	})
}

func Test_C77_SSO_ValueInt(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueInt", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if sso.ValueInt(0) != 42 {
			t.Error("expected 42")
		}
		if invalid.ValueInt(99) != 99 {
			t.Error("expected 99")
		}
	})
}

func Test_C77_SSO_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueDefInt", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("10")
		invalid := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		if sso.ValueDefInt() != 10 {
			t.Error("expected 10")
		}
		if invalid.ValueDefInt() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C77_SSO_ValueByte(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueByte", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("100")
		overflow := corestr.New.SimpleStringOnce.Init("300")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if sso.ValueByte(0) != 100 {
			t.Error("expected 100")
		}
		if overflow.ValueByte(5) != 5 {
			t.Error("expected 5 for overflow")
		}
		if invalid.ValueByte(7) != 7 {
			t.Error("expected 7 for invalid")
		}
	})
}

func Test_C77_SSO_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueDefByte", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")
		overflow := corestr.New.SimpleStringOnce.Init("999")

		// Act & Assert
		if sso.ValueDefByte() != 50 {
			t.Error("expected 50")
		}
		if overflow.ValueDefByte() != 0 {
			t.Error("expected 0")
		}
	})
}

func Test_C77_SSO_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueFloat64", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if sso.ValueFloat64(0) != 3.14 {
			t.Error("expected 3.14")
		}
		if invalid.ValueFloat64(1.5) != 1.5 {
			t.Error("expected 1.5")
		}
	})
}

func Test_C77_SSO_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueDefFloat64", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("2.5")

		// Act & Assert
		if sso.ValueDefFloat64() != 2.5 {
			t.Error("expected 2.5")
		}
	})
}

func Test_C77_SSO_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_C77_SSO_NonPtr_Ptr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		nonPtr := sso.NonPtr()
		ptr := sso.Ptr()

		// Assert
		if nonPtr.Value() != "x" {
			t.Error("nonPtr mismatch")
		}
		if ptr == nil {
			t.Error("ptr nil")
		}
	})
}

func Test_C77_SSO_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_HasSafeNonEmpty", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		empty := corestr.New.SimpleStringOnce.Init("")

		// Act & Assert
		if !valid.HasSafeNonEmpty() {
			t.Error("expected true")
		}
		if empty.HasSafeNonEmpty() {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_Is(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Is", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act & Assert
		if !sso.Is("hello") {
			t.Error("expected true")
		}
		if sso.Is("world") {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyOf", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("b")

		// Act & Assert
		if !sso.IsAnyOf("a", "b", "c") {
			t.Error("expected true")
		}
		if sso.IsAnyOf("x", "y") {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_IsAnyOf_Empty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyOf_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert — empty values returns true
		if !sso.IsAnyOf() {
			t.Error("expected true for empty values")
		}
	})
}

func Test_C77_SSO_IsContains(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsContains", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello world")

		// Act & Assert
		if !sso.IsContains("world") {
			t.Error("expected true")
		}
		if sso.IsContains("xyz") {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyContains", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello world")

		// Act & Assert
		if !sso.IsAnyContains("xyz", "world") {
			t.Error("expected true")
		}
		if sso.IsAnyContains("abc", "def") {
			t.Error("expected false")
		}
	})
}

func Test_C77_SSO_IsAnyContains_Empty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyContains_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		if !sso.IsAnyContains() {
			t.Error("expected true for empty")
		}
	})
}

func Test_C77_SSO_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsEqualNonSensitive", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("Hello")

		// Act & Assert
		if !sso.IsEqualNonSensitive("hello") {
			t.Error("expected true")
		}
	})
}

func Test_C77_SSO_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsRegexMatches", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		if !sso.IsRegexMatches(re) {
			t.Error("expected true")
		}
		if sso.IsRegexMatches(nil) {
			t.Error("expected false for nil regex")
		}
	})
}

func Test_C77_SSO_RegexFindString(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindString", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc123def")
		re := regexp.MustCompile(`\d+`)

		// Act
		result := sso.RegexFindString(re)

		// Assert
		if result != "123" {
			t.Errorf("expected '123', got '%s'", result)
		}
	})
}

func Test_C77_SSO_RegexFindString_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindString_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		if sso.RegexFindString(nil) != "" {
			t.Error("expected empty for nil regex")
		}
	})
}

func Test_C77_SSO_RegexFindAllStringsWithFlag(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindAllStringsWithFlag", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a1b2c3")
		re := regexp.MustCompile(`\d`)

		// Act
		items, hasAny := sso.RegexFindAllStringsWithFlag(re, -1)

		// Assert
		if !hasAny || len(items) != 3 {
			t.Errorf("expected 3 matches, got %d", len(items))
		}
	})
}

func Test_C77_SSO_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		items, hasAny := sso.RegexFindAllStringsWithFlag(nil, -1)

		// Assert
		if hasAny || len(items) != 0 {
			t.Error("expected empty for nil regex")
		}
	})
}

func Test_C77_SSO_RegexFindAllStrings(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindAllStrings", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a1b2")
		re := regexp.MustCompile(`\d`)

		// Act
		items := sso.RegexFindAllStrings(re, -1)

		// Assert
		if len(items) != 2 {
			t.Errorf("expected 2, got %d", len(items))
		}
	})
}

func Test_C77_SSO_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindAllStrings_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		items := sso.RegexFindAllStrings(nil, -1)

		// Assert
		if len(items) != 0 {
			t.Error("expected empty")
		}
	})
}
