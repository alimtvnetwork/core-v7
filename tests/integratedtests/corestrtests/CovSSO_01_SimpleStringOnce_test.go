package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// SimpleStringOnce — Segment 16+17 Part 1
// ══════════════════════════════════════════════════════════════════════════════

func newSSO(val string) *corestr.SimpleStringOnce {
	sso := &corestr.SimpleStringOnce{}
	sso.SetOnUninitialized(val)
	return sso
}

func Test_CovSSO_01_Value_IsInitialized_IsDefined(t *testing.T) {
	safeTest(t, "Test_CovSSO_01_Value_IsInitialized_IsDefined", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
		actual = args.Map{"result": sso.IsDefined()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected undefined", actual)
		actual = args.Map{"result": sso.IsUninitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
		sso.SetOnUninitialized("hello")
		actual = args.Map{"result": sso.Value() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
		actual = args.Map{"result": sso.IsDefined()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected defined", actual)
	})
}

func Test_CovSSO_02_SetOnUninitialized_AlreadyInit(t *testing.T) {
	safeTest(t, "Test_CovSSO_02_SetOnUninitialized_AlreadyInit", func() {
		// Arrange
		sso := newSSO("first")
		err := sso.SetOnUninitialized("second")

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovSSO_03_GetSetOnce(t *testing.T) {
	safeTest(t, "Test_CovSSO_03_GetSetOnce", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		v := sso.GetSetOnce("hello")

		// Act
		actual := args.Map{"result": v != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		// already init
		v2 := sso.GetSetOnce("world")
		actual = args.Map{"result": v2 != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovSSO_04_GetOnce(t *testing.T) {
	safeTest(t, "Test_CovSSO_04_GetOnce", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		v := sso.GetOnce()

		// Act
		actual := args.Map{"result": v != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		// already init
		v2 := sso.GetOnce()
		actual = args.Map{"result": v2 != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_CovSSO_05_GetOnceFunc(t *testing.T) {
	safeTest(t, "Test_CovSSO_05_GetOnceFunc", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		v := sso.GetOnceFunc(func() string { return "computed" })

		// Act
		actual := args.Map{"result": v != "computed"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected computed", actual)
		// already init
		v2 := sso.GetOnceFunc(func() string { return "other" })
		actual = args.Map{"result": v2 != "computed"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected computed", actual)
	})
}

func Test_CovSSO_06_SetOnceIfUninitialized(t *testing.T) {
	safeTest(t, "Test_CovSSO_06_SetOnceIfUninitialized", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		ok := sso.SetOnceIfUninitialized("hello")

		// Act
		actual := args.Map{"result": ok}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		ok2 := sso.SetOnceIfUninitialized("world")
		actual = args.Map{"result": ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovSSO_07_Invalidate_Reset(t *testing.T) {
	safeTest(t, "Test_CovSSO_07_Invalidate_Reset", func() {
		// Arrange
		sso := newSSO("hello")
		sso.Invalidate()

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
		sso.SetOnUninitialized("new")
		sso.Reset()
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_CovSSO_08_SetInitialize_SetUnInit(t *testing.T) {
	safeTest(t, "Test_CovSSO_08_SetInitialize_SetUnInit", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}
		sso.SetInitialize()

		// Act
		actual := args.Map{"result": sso.IsInitialized()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected initialized", actual)
		sso.SetUnInit()
		actual = args.Map{"result": sso.IsInitialized()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected uninitialized", actual)
	})
}

func Test_CovSSO_09_IsInvalid(t *testing.T) {
	safeTest(t, "Test_CovSSO_09_IsInvalid", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": sso.IsInvalid()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid", actual)
		sso.SetOnUninitialized("")
		actual = args.Map{"result": sso.IsInvalid()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected invalid for empty value", actual)
		sso2 := newSSO("hello")
		actual = args.Map{"result": sso2.IsInvalid()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid", actual)
	})
}

func Test_CovSSO_10_IsEmpty_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_CovSSO_10_IsEmpty_IsWhitespace", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": sso.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": sso.IsWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected whitespace", actual)
		sso2 := newSSO("hello")
		actual = args.Map{"result": sso2.IsEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_CovSSO_11_Trim(t *testing.T) {
	safeTest(t, "Test_CovSSO_11_Trim", func() {
		// Arrange
		sso := newSSO("  hello  ")

		// Act
		actual := args.Map{"result": sso.Trim() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected trimmed", actual)
	})
}

func Test_CovSSO_12_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty(t *testing.T) {
	safeTest(t, "Test_CovSSO_12_HasValidNonEmpty_HasValidNonWhitespace_HasSafeNonEmpty", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": sso.HasValidNonEmpty()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": sso.HasValidNonWhitespace()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": sso.HasSafeNonEmpty()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		sso2 := newSSO("hello")
		actual = args.Map{"result": sso2.HasValidNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso2.HasValidNonWhitespace()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso2.HasSafeNonEmpty()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSSO_13_SafeValue(t *testing.T) {
	safeTest(t, "Test_CovSSO_13_SafeValue", func() {
		// Arrange
		sso := &corestr.SimpleStringOnce{}

		// Act
		actual := args.Map{"result": sso.SafeValue() != ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		sso2 := newSSO("hello")
		actual = args.Map{"result": sso2.SafeValue() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovSSO_14_ValueBytes_ValueBytesPtr(t *testing.T) {
	safeTest(t, "Test_CovSSO_14_ValueBytes_ValueBytesPtr", func() {
		// Arrange
		sso := newSSO("hi")

		// Act
		actual := args.Map{"result": len(sso.ValueBytes()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(sso.ValueBytesPtr()) != 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_CovSSO_15_Int_ValueInt_ValueDefInt(t *testing.T) {
	safeTest(t, "Test_CovSSO_15_Int_ValueInt_ValueDefInt", func() {
		// Arrange
		sso := newSSO("42")

		// Act
		actual := args.Map{"result": sso.Int() != 42}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": sso.ValueInt(0) != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		actual = args.Map{"result": sso.ValueDefInt() != 42}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 42", actual)
		// invalid
		sso2 := newSSO("abc")
		actual = args.Map{"result": sso2.Int() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": sso2.ValueInt(99) != 99}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 99", actual)
		actual = args.Map{"result": sso2.ValueDefInt() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSSO_16_Byte_ValueByte_ValueDefByte(t *testing.T) {
	safeTest(t, "Test_CovSSO_16_Byte_ValueByte_ValueDefByte", func() {
		// Arrange
		sso := newSSO("100")

		// Act
		actual := args.Map{"result": sso.Byte() != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual = args.Map{"result": sso.ValueByte(0) != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual = args.Map{"result": sso.ValueDefByte() != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		// out of range
		sso2 := newSSO("999")
		actual = args.Map{"result": sso2.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": sso2.ValueByte(5) != 5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 5", actual)
		// invalid
		sso3 := newSSO("abc")
		actual = args.Map{"result": sso3.Byte() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": sso3.ValueByte(7) != 7}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 7", actual)
	})
}

func Test_CovSSO_17_Int16_Int32(t *testing.T) {
	safeTest(t, "Test_CovSSO_17_Int16_Int32", func() {
		// Arrange
		sso := newSSO("100")

		// Act
		actual := args.Map{"result": sso.Int16() != 100}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		actual = args.Map{"result": sso.Int32() != 100}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100", actual)
		// out of range for int16
		sso2 := newSSO("99999")
		actual = args.Map{"result": sso2.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// invalid
		sso3 := newSSO("abc")
		actual = args.Map{"result": sso3.Int16() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		actual = args.Map{"result": sso3.Int32() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_CovSSO_18_Uint16_Uint32(t *testing.T) {
	safeTest(t, "Test_CovSSO_18_Uint16_Uint32", func() {
		// Arrange
		sso := newSSO("100")
		v16, ok16 := sso.Uint16()

		// Act
		actual := args.Map{"result": v16 != 100 || !ok16}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100 in range", actual)
		v32, ok32 := sso.Uint32()
		actual = args.Map{"result": v32 != 100 || !ok32}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 100 in range", actual)
	})
}

func Test_CovSSO_19_WithinRange_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_CovSSO_19_WithinRange_WithinRangeDefault", func() {
		// Arrange
		sso := newSSO("50")
		v, ok := sso.WithinRange(true, 0, 100)

		// Act
		actual := args.Map{"result": v != 50 || !ok}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50 in range", actual)
		// out of range with boundary
		v2, ok2 := sso.WithinRange(true, 60, 100)
		actual = args.Map{"result": v2 != 60 || ok2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected min boundary", actual)
		// out of range, above max
		sso2 := newSSO("200")
		v3, ok3 := sso2.WithinRange(true, 0, 100)
		actual = args.Map{"result": v3 != 100 || ok3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected max boundary", actual)
		// no boundary
		v4, ok4 := sso2.WithinRange(false, 0, 100)
		actual = args.Map{"result": v4 != 200 || ok4}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 200 out of range", actual)
		// invalid
		sso3 := newSSO("abc")
		v5, ok5 := sso3.WithinRange(true, 0, 100)
		actual = args.Map{"result": v5 != 0 || ok5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		// WithinRangeDefault
		v6, ok6 := sso.WithinRangeDefault(0, 100)
		actual = args.Map{"result": v6 != 50 || !ok6}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 50", actual)
	})
}

func Test_CovSSO_20_ValueFloat64_ValueDefFloat64(t *testing.T) {
	safeTest(t, "Test_CovSSO_20_ValueFloat64_ValueDefFloat64", func() {
		// Arrange
		sso := newSSO("3.14")

		// Act
		actual := args.Map{"result": sso.ValueFloat64(0) != 3.14}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		actual = args.Map{"result": sso.ValueDefFloat64() != 3.14}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3.14", actual)
		// invalid
		sso2 := newSSO("abc")
		actual = args.Map{"result": sso2.ValueFloat64(1.5) != 1.5}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1.5", actual)
	})
}

func Test_CovSSO_21_Boolean_BooleanDefault_IsValueBool(t *testing.T) {
	safeTest(t, "Test_CovSSO_21_Boolean_BooleanDefault_IsValueBool", func() {
		// Arrange
		cases := []struct {
			val    string
			expect bool
		}{
			{"yes", true}, {"y", true}, {"1", true}, {"YES", true}, {"Y", true},
			{"true", true}, {"false", false}, {"", false}, {"abc", false},
		}
		for _, c := range cases {
			sso := newSSO(c.val)

		// Act
			actual := args.Map{"result": sso.Boolean(false) != c.expect}

		// Assert
			expected := args.Map{"result": false}
			expected.ShouldBeEqual(t, 0, "Boolean() expected", actual)
		}
		// with consider init
		uninit := &corestr.SimpleStringOnce{}
		actual = args.Map{"result": uninit.Boolean(true)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for uninitialized", actual)
		// BooleanDefault
		sso := newSSO("true")
		actual = args.Map{"result": sso.BooleanDefault()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		// IsValueBool
		actual = args.Map{"result": sso.IsValueBool()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSSO_22_IsSetter(t *testing.T) {
	safeTest(t, "Test_CovSSO_22_IsSetter", func() {
		// Arrange
		sso := newSSO("yes")
		v := sso.IsSetter(false)

		// Act
		actual := args.Map{"result": v.String() != "True"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected True", actual)
		// invalid
		sso2 := newSSO("abc")
		v2 := sso2.IsSetter(false)
		actual = args.Map{"result": v2.String() != "Uninitialized"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected Uninitialized", actual)
		// consider init, uninitialized
		uninit := &corestr.SimpleStringOnce{}
		v3 := uninit.IsSetter(true)
		actual = args.Map{"result": v3.String() != "False"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected False", actual)
		// false value
		sso3 := newSSO("false")
		v4 := sso3.IsSetter(false)
		actual = args.Map{"result": v4.String() != "False"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected False", actual)
	})
}

func Test_CovSSO_23_Is_IsAnyOf(t *testing.T) {
	safeTest(t, "Test_CovSSO_23_Is_IsAnyOf", func() {
		// Arrange
		sso := newSSO("hello")

		// Act
		actual := args.Map{"result": sso.Is("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.Is("world")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": sso.IsAnyOf("a", "hello")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyOf("a", "b")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		// empty values → true
		actual = args.Map{"result": sso.IsAnyOf()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSSO_24_IsContains_IsAnyContains(t *testing.T) {
	safeTest(t, "Test_CovSSO_24_IsContains_IsAnyContains", func() {
		// Arrange
		sso := newSSO("hello world")

		// Act
		actual := args.Map{"result": sso.IsContains("world")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyContains("xyz", "world")}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsAnyContains("xyz", "abc")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": sso.IsAnyContains()}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
	})
}

func Test_CovSSO_25_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_CovSSO_25_IsEqualNonSensitive", func() {
		// Arrange
		sso := newSSO("Hello")

		// Act
		actual := args.Map{"result": sso.IsEqualNonSensitive("hello")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_CovSSO_26_IsRegexMatches_RegexFind(t *testing.T) {
	safeTest(t, "Test_CovSSO_26_IsRegexMatches_RegexFind", func() {
		// Arrange
		sso := newSSO("hello123")
		re := regexp.MustCompile(`\d+`)

		// Act
		actual := args.Map{"result": sso.IsRegexMatches(re)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": sso.IsRegexMatches(nil)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for nil", actual)
		found := sso.RegexFindString(re)
		actual = args.Map{"result": found != "123"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 123", actual)
		actual = args.Map{"result": sso.RegexFindString(nil) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		items := sso.RegexFindAllStrings(re, -1)
		actual = args.Map{"result": len(items) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": len(sso.RegexFindAllStrings(nil, -1)) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
		items2, has := sso.RegexFindAllStringsWithFlag(re, -1)
		actual = args.Map{"result": has || len(items2) != 1}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected found", actual)
		_, has2 := sso.RegexFindAllStringsWithFlag(nil, -1)
		actual = args.Map{"result": has2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_CovSSO_27_Split_SplitLeftRight_SplitLeftRightTrim(t *testing.T) {
	safeTest(t, "Test_CovSSO_27_Split_SplitLeftRight_SplitLeftRightTrim", func() {
		// Arrange
		sso := newSSO("key=value")
		parts := sso.Split("=")

		// Act
		actual := args.Map{"result": len(parts) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		l, r := sso.SplitLeftRight("=")
		actual = args.Map{"result": l != "key" || r != "value"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key,value", actual)
		// no right
		sso2 := newSSO("onlykey")
		l2, r2 := sso2.SplitLeftRight("=")
		actual = args.Map{"result": l2 != "onlykey" || r2 != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected onlykey,empty", actual)
		// trim
		sso3 := newSSO(" key = value ")
		l3, r3 := sso3.SplitLeftRightTrim("=")
		actual = args.Map{"result": l3 != "key" || r3 != "value"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected key,value got '',''", actual)
		// no right trim
		sso4 := newSSO("onlykey")
		l4, r4 := sso4.SplitLeftRightTrim("=")
		actual = args.Map{"result": l4 != "onlykey" || r4 != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected onlykey,empty", actual)
	})
}

func Test_CovSSO_28_SplitNonEmpty_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_CovSSO_28_SplitNonEmpty_SplitTrimNonWhitespace", func() {
		// Arrange
		sso := newSSO("a,,b")
		parts := sso.SplitNonEmpty(",")

		// Act
		actual := args.Map{"result": len(parts) < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
		sso2 := newSSO("a, ,b")
		parts2 := sso2.SplitTrimNonWhitespace(",")
		actual = args.Map{"result": len(parts2) < 2}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_CovSSO_29_LinesSimpleSlice_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_CovSSO_29_LinesSimpleSlice_SimpleSlice", func() {
		// Arrange
		sso := newSSO("a\nb")
		ss := sso.LinesSimpleSlice()

		// Act
		actual := args.Map{"result": ss.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ss2 := sso.SimpleSlice(",")
		actual = args.Map{"result": ss2.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_CovSSO_30_ConcatNew_ConcatNewUsingStrings(t *testing.T) {
	safeTest(t, "Test_CovSSO_30_ConcatNew_ConcatNewUsingStrings", func() {
		// Arrange
		sso := newSSO("hello")
		c := sso.ConcatNew(" world")

		// Act
		actual := args.Map{"result": c.Value() != "hello world"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello world", actual)
		c2 := sso.ConcatNewUsingStrings("-", "a", "b")
		actual = args.Map{"result": c2.Value() != "hello-a-b"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello-a-b', got ''", actual)
	})
}

func Test_CovSSO_31_NonPtr_Ptr(t *testing.T) {
	safeTest(t, "Test_CovSSO_31_NonPtr_Ptr", func() {
		// Arrange
		sso := newSSO("hello")
		np := sso.NonPtr()

		// Act
		actual := args.Map{"result": np.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		p := sso.Ptr()
		actual = args.Map{"result": p.Value() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovSSO_32_Clone_ClonePtr_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_CovSSO_32_Clone_ClonePtr_CloneUsingNewVal", func() {
		// Arrange
		sso := newSSO("hello")
		c := sso.Clone()

		// Act
		actual := args.Map{"result": c.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		cp := sso.ClonePtr()
		actual = args.Map{"result": cp.Value() != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		cv := sso.CloneUsingNewVal("new")
		actual = args.Map{"result": cv.Value() != "new"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected new", actual)
	})
}

func Test_CovSSO_33_String_StringPtr(t *testing.T) {
	safeTest(t, "Test_CovSSO_33_String_StringPtr", func() {
		// Arrange
		sso := newSSO("hello")

		// Act
		actual := args.Map{"result": sso.String() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		sp := sso.StringPtr()
		actual = args.Map{"result": *sp != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_CovSSO_34_Dispose(t *testing.T) {
	safeTest(t, "Test_CovSSO_34_Dispose", func() {
		sso := newSSO("hello")
		sso.Dispose()
	})
}

func Test_CovSSO_35_JsonModel_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_CovSSO_35_JsonModel_MarshalUnmarshal", func() {
		// Arrange
		sso := newSSO("hello")
		_ = sso.JsonModel()
		_ = sso.JsonModelAny()
		data, err := sso.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		sso2 := &corestr.SimpleStringOnce{}
		err2 := sso2.UnmarshalJSON(data)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		// invalid
		err3 := sso2.UnmarshalJSON([]byte("bad"))
		actual = args.Map{"result": err3 == nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_CovSSO_36_Json_JsonPtr_ParseInject(t *testing.T) {
	safeTest(t, "Test_CovSSO_36_Json_JsonPtr_ParseInject", func() {
		// Arrange
		sso := newSSO("hello")
		_ = sso.Json()
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		r, err := sso2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"result": err != nil || r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovSSO_37_ParseInjectUsingJsonMust(t *testing.T) {
	safeTest(t, "Test_CovSSO_37_ParseInjectUsingJsonMust", func() {
		// Arrange
		sso := newSSO("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		r := sso2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"result": r == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_CovSSO_38_JsonParseSelfInject(t *testing.T) {
	safeTest(t, "Test_CovSSO_38_JsonParseSelfInject", func() {
		// Arrange
		sso := newSSO("hello")
		jr := sso.JsonPtr()
		sso2 := &corestr.SimpleStringOnce{}
		err := sso2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}

func Test_CovSSO_39_AsInterfaces(t *testing.T) {
	safeTest(t, "Test_CovSSO_39_AsInterfaces", func() {
		sso := newSSO("hello")
		_ = sso.AsJsonContractsBinder()
		_ = sso.AsJsoner()
		_ = sso.AsJsonParseSelfInjector()
		_ = sso.AsJsonMarshaller()
	})
}

func Test_CovSSO_40_Serialize_Deserialize(t *testing.T) {
	safeTest(t, "Test_CovSSO_40_Serialize_Deserialize", func() {
		// Arrange
		sso := newSSO("hello")
		_, err := sso.Serialize()

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
		target := &corestr.SimpleStringOnce{}
		err2 := sso.Deserialize(target)
		actual = args.Map{"result": err2 != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	})
}
