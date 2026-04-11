package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/issetter"
	"github.com/alimtvnetwork/core/coretests/args"
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
		actual := args.Map{"result": sso.Value() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello', got ''", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
	})
}

func Test_C77_SSO_IsDefined(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsDefined", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": sso.IsDefined()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected defined", actual)
	})
}

func Test_C77_SSO_IsUninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act & Assert
		actual := args.Map{"result": sso.IsUninitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_C77_SSO_Invalidate(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Invalidate", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		sso.Invalidate()

		// Assert
		actual := args.Map{"result": sso.IsInitialized() || sso.Value() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected invalidated", actual)
	})
}

func Test_C77_SSO_Reset(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Reset", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		sso.Reset()

		// Assert
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected reset", actual)
	})
}

func Test_C77_SSO_IsInvalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsInvalid", func() {
		// Arrange
		uninit := corestr.New.SimpleStringOnce.Empty()
		initEmpty := corestr.New.SimpleStringOnce.Init("")
		valid := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": uninit.IsInvalid()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for uninitialized", actual)
		actual = args.Map{"result": initEmpty.IsInvalid()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for empty value", actual)
		actual = args.Map{"result": valid.IsInvalid()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
	})
}

func Test_C77_SSO_IsInvalid_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsInvalid_Nil", func() {
		// Arrange
		var sso *corestr.SimpleStringOnce

		// Act & Assert
		actual := args.Map{"result": sso.IsInvalid()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for nil", actual)
	})
}

func Test_C77_SSO_ValueBytes(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueBytes", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		result := sso.ValueBytes()

		// Assert
		actual := args.Map{"result": string(result) != "abc"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "bytes mismatch", actual)
	})
}

func Test_C77_SSO_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueBytesPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		result := sso.ValueBytesPtr()

		// Assert
		actual := args.Map{"result": string(result) != "xyz"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "bytes mismatch", actual)
	})
}

func Test_C77_SSO_SetOnUninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		err := sso.SetOnUninitialized("val")

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error:", actual)
		actual = args.Map{"result": sso.Value() != "val"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value not set", actual)
	})
}

func Test_C77_SSO_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		err := sso.SetOnUninitialized("new")

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error for already initialized", actual)
		actual = args.Map{"result": sso.Value() != "existing"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "value should not change", actual)
	})
}

func Test_C77_SSO_GetSetOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetSetOnce_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetSetOnce("first")

		// Assert
		actual := args.Map{"result": result != "first"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'first', got ''", actual)
	})
}

func Test_C77_SSO_GetSetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetSetOnce_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		result := sso.GetSetOnce("new")

		// Assert
		actual := args.Map{"result": result != "existing"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'existing', got ''", actual)
	})
}

func Test_C77_SSO_GetOnce_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnce_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetOnce()

		// Assert
		actual := args.Map{"result": result != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty string", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "should be initialized after GetOnce", actual)
	})
}

func Test_C77_SSO_GetOnce_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnce_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("val")

		// Act
		result := sso.GetOnce()

		// Assert
		actual := args.Map{"result": result != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'val'", actual)
	})
}

func Test_C77_SSO_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnceFunc", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		result := sso.GetOnceFunc(func() string { return "computed" })

		// Assert
		actual := args.Map{"result": result != "computed"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'computed', got ''", actual)
	})
}

func Test_C77_SSO_GetOnceFunc_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_GetOnceFunc_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("existing")

		// Act
		result := sso.GetOnceFunc(func() string { return "new" })

		// Assert
		actual := args.Map{"result": result != "existing"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'existing'", actual)
	})
}

func Test_C77_SSO_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnceIfUninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		isSet := sso.SetOnceIfUninitialized("val")

		// Assert
		actual := args.Map{"result": isSet}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C77_SSO_SetOnceIfUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetOnceIfUninitialized_AlreadyInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		isSet := sso.SetOnceIfUninitialized("new")

		// Assert
		actual := args.Map{"result": isSet}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SetInitialize_SetUnInit", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Empty()

		// Act
		sso.SetInitialize()

		// Assert
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)

		// Act
		sso.SetUnInit()

		// Assert
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_C77_SSO_ConcatNew(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ConcatNew", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		result := sso.ConcatNew(" world")

		// Assert
		actual := args.Map{"result": result.Value() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello world', got ''", actual)
	})
}

func Test_C77_SSO_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ConcatNewUsingStrings", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("a")

		// Act
		result := sso.ConcatNewUsingStrings("-", "b", "c")

		// Assert
		actual := args.Map{"result": result.Value() != "a-b-c"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a-b-c', got ''", actual)
	})
}

func Test_C77_SSO_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsEmpty_IsWhitespace", func() {
		// Arrange
		empty := corestr.New.SimpleStringOnce.Init("")
		ws := corestr.New.SimpleStringOnce.Init("  ")
		val := corestr.New.SimpleStringOnce.Init("x")

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

func Test_C77_SSO_Trim(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Trim", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init(" hello ")

		// Act & Assert
		actual := args.Map{"result": sso.Trim() != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "trim mismatch", actual)
	})
}

func Test_C77_SSO_HasValidNonEmpty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_HasValidNonEmpty", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		empty := corestr.New.SimpleStringOnce.Init("")

		// Act & Assert
		actual := args.Map{"result": valid.HasValidNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": empty.HasValidNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_HasValidNonWhitespace(t *testing.T) {
	safeTest(t, "Test_C77_SSO_HasValidNonWhitespace", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		ws := corestr.New.SimpleStringOnce.Init("  ")

		// Act & Assert
		actual := args.Map{"result": valid.HasValidNonWhitespace()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ws.HasValidNonWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_IsValueBool(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsValueBool", func() {
		// Arrange
		ssoFalse := corestr.New.SimpleStringOnce.Init("false")
		ssoTrue := corestr.New.SimpleStringOnce.Init("true")

		// Act & Assert
		actual := args.Map{"result": ssoFalse.IsValueBool()}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": ssoTrue.IsValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C77_SSO_SafeValue(t *testing.T) {
	safeTest(t, "Test_C77_SSO_SafeValue", func() {
		// Arrange
		init := corestr.New.SimpleStringOnce.Init("val")
		uninit := corestr.New.SimpleStringOnce.Empty()

		// Act & Assert
		actual := args.Map{"result": init.SafeValue() != "val"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'val'", actual)
		actual = args.Map{"result": uninit.SafeValue() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for uninitialized", actual)
	})
}

func Test_C77_SSO_Int(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Int", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.Int() != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": invalid.Int() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C77_SSO_Byte(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Byte", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("200")
		overflow := corestr.New.SimpleStringOnce.Init("300")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": valid.Byte() != 200}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200", actual)
		actual = args.Map{"result": overflow.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for overflow", actual)
		actual = args.Map{"result": invalid.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for invalid", actual)
	})
}

func Test_C77_SSO_Int16(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Int16", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("1000")
		overflow := corestr.New.SimpleStringOnce.Init("40000")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": valid.Int16() != 1000}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1000", actual)
		actual = args.Map{"result": overflow.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for overflow", actual)
		actual = args.Map{"result": invalid.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0 for invalid", actual)
	})
}

func Test_C77_SSO_Int32(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Int32", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("100000")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": valid.Int32() != 100000}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100000", actual)
		actual = args.Map{"result": invalid.Int32() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C77_SSO_Uint16(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Uint16", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("500")

		// Act
		val, inRange := valid.Uint16()

		// Assert
		actual := args.Map{"result": inRange || val != 500}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 500 in range, got", actual)
	})
}

func Test_C77_SSO_Uint32(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Uint32", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("100000")

		// Act
		val, inRange := valid.Uint32()

		// Assert
		actual := args.Map{"result": inRange || val != 100000}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 100000", actual)
	})
}

func Test_C77_SSO_WithinRange_InRange(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_InRange", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange || val != 50}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50 in range", actual)
	})
}

func Test_C77_SSO_WithinRange_BelowMin_WithBoundary(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_BelowMin_WithBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("-5")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		actual = args.Map{"result": val != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary min 0", actual)
	})
}

func Test_C77_SSO_WithinRange_AboveMax_WithBoundary(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_AboveMax_WithBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		actual = args.Map{"result": val != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected boundary max 100", actual)
	})
}

func Test_C77_SSO_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_NoBoundary", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("200")

		// Act
		val, inRange := sso.WithinRange(false, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected out of range", actual)
		actual = args.Map{"result": val != 200}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected raw value 200", actual)
	})
}

func Test_C77_SSO_WithinRange_Invalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRange_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		val, inRange := sso.WithinRange(true, 0, 100)

		// Assert
		actual := args.Map{"result": inRange}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
		actual = args.Map{"result": val != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C77_SSO_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_C77_SSO_WithinRangeDefault", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")

		// Act
		val, inRange := sso.WithinRangeDefault(0, 100)

		// Assert
		actual := args.Map{"result": inRange || val != 50}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 50 in range", actual)
	})
}

func Test_C77_SSO_Boolean_True_Values(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_True_Values", func() {
		// Arrange
		tests := []string{"true", "yes", "y", "1", "YES", "Y"}

		for _, v := range tests {
			sso := corestr.New.SimpleStringOnce.Init(v)

			// Act & Assert
			actual := args.Map{"result": sso.Boolean(false)}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected true for ''", actual)
		}
	})
}

func Test_C77_SSO_Boolean_False_Values(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_False_Values", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("false")

		// Act & Assert
		actual := args.Map{"result": sso.Boolean(false)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_Boolean_Invalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act & Assert
		actual := args.Map{"result": sso.Boolean(false)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for invalid", actual)
	})
}

func Test_C77_SSO_Boolean_ConsiderInit_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Boolean_ConsiderInit_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")

		// Act & Assert
		actual := args.Map{"result": sso.Boolean(true)}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for uninitialized with considerInit", actual)
	})
}

func Test_C77_SSO_BooleanDefault(t *testing.T) {
	safeTest(t, "Test_C77_SSO_BooleanDefault", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("yes")

		// Act & Assert
		actual := args.Map{"result": sso.BooleanDefault()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C77_SSO_IsSetter_True(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_True", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("yes")

		// Act
		result := sso.IsSetter(false)

		// Assert
		actual := args.Map{"result": result != issetter.True}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected True", actual)
	})
}

func Test_C77_SSO_IsSetter_False(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_False", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("false")

		// Act
		result := sso.IsSetter(false)

		// Assert
		actual := args.Map{"result": result != issetter.False}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected False", actual)
	})
}

func Test_C77_SSO_IsSetter_Invalid(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_Invalid", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("xyz")

		// Act
		result := sso.IsSetter(false)

		// Assert
		actual := args.Map{"result": result != issetter.Uninitialized}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected Uninitialized", actual)
	})
}

func Test_C77_SSO_IsSetter_ConsiderInit_Uninitialized(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsSetter_ConsiderInit_Uninitialized", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")

		// Act
		result := sso.IsSetter(true)

		// Assert
		actual := args.Map{"result": result != issetter.False}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected False", actual)
	})
}

func Test_C77_SSO_ValueInt(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueInt", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("42")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.ValueInt(0) != 42}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": invalid.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
	})
}

func Test_C77_SSO_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueDefInt", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("10")
		invalid := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": sso.ValueDefInt() != 10}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 10", actual)
		actual = args.Map{"result": invalid.ValueDefInt() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C77_SSO_ValueByte(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueByte", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("100")
		overflow := corestr.New.SimpleStringOnce.Init("300")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.ValueByte(0) != 100}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual = args.Map{"result": overflow.ValueByte(5) != 5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5 for overflow", actual)
		actual = args.Map{"result": invalid.ValueByte(7) != 7}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 7 for invalid", actual)
	})
}

func Test_C77_SSO_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueDefByte", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("50")
		overflow := corestr.New.SimpleStringOnce.Init("999")

		// Act & Assert
		actual := args.Map{"result": sso.ValueDefByte() != 50}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
		actual = args.Map{"result": overflow.ValueDefByte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_C77_SSO_ValueFloat64(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueFloat64", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		invalid := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.ValueFloat64(0) != 3.14}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual = args.Map{"result": invalid.ValueFloat64(1.5) != 1.5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	})
}

func Test_C77_SSO_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_C77_SSO_ValueDefFloat64", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("2.5")

		// Act & Assert
		actual := args.Map{"result": sso.ValueDefFloat64() != 2.5}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2.5", actual)
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
		actual := args.Map{"result": nonPtr.Value() != "x"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "nonPtr mismatch", actual)
		actual = args.Map{"result": ptr == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "ptr nil", actual)
	})
}

func Test_C77_SSO_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_HasSafeNonEmpty", func() {
		// Arrange
		valid := corestr.New.SimpleStringOnce.Init("x")
		empty := corestr.New.SimpleStringOnce.Init("")

		// Act & Assert
		actual := args.Map{"result": valid.HasSafeNonEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": empty.HasSafeNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_Is(t *testing.T) {
	safeTest(t, "Test_C77_SSO_Is", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act & Assert
		actual := args.Map{"result": sso.Is("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.Is("world")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyOf", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("b")

		// Act & Assert
		actual := args.Map{"result": sso.IsAnyOf("a", "b", "c")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyOf("x", "y")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_IsAnyOf_Empty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyOf_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert — empty values returns true
		actual := args.Map{"result": sso.IsAnyOf()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty values", actual)
	})
}

func Test_C77_SSO_IsContains(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsContains", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello world")

		// Act & Assert
		actual := args.Map{"result": sso.IsContains("world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsContains("xyz")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyContains", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello world")

		// Act & Assert
		actual := args.Map{"result": sso.IsAnyContains("xyz", "world")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyContains("abc", "def")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_C77_SSO_IsAnyContains_Empty(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsAnyContains_Empty", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act & Assert
		actual := args.Map{"result": sso.IsAnyContains()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_C77_SSO_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsEqualNonSensitive", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("Hello")

		// Act & Assert
		actual := args.Map{"result": sso.IsEqualNonSensitive("hello")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_C77_SSO_IsRegexMatches(t *testing.T) {
	safeTest(t, "Test_C77_SSO_IsRegexMatches", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)

		// Act & Assert
		actual := args.Map{"result": sso.IsRegexMatches(re)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil regex", actual)
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
		actual := args.Map{"result": result != "123"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected '123', got ''", actual)
	})
}

func Test_C77_SSO_RegexFindString_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindString_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act & Assert
		actual := args.Map{"result": sso.RegexFindString(nil) != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil regex", actual)
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
		actual := args.Map{"result": hasAny || len(items) != 3}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected 3 matches", actual)
	})
}

func Test_C77_SSO_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindAllStringsWithFlag_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		items, hasAny := sso.RegexFindAllStringsWithFlag(nil, -1)

		// Assert
		actual := args.Map{"result": hasAny || len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty for nil regex", actual)
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
		actual := args.Map{"result": len(items) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_C77_SSO_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_C77_SSO_RegexFindAllStrings_Nil", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("abc")

		// Act
		items := sso.RegexFindAllStrings(nil, -1)

		// Assert
		actual := args.Map{"result": len(items) != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}
