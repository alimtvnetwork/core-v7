package corestrtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/issetter"
)

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Split methods
// ═══════════════════════════════════════════════════════════════

func Test_SSO_SplitLeftRight_HasBoth(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SplitLeftRight_HasBoth", func() {
		sso := corestr.New.SimpleStringOnce.Init("key=val")
		left, right := sso.SplitLeftRight("=")
		tc := caseV1Compat{Name: "SplitLeftRight left", Expected: "key", Actual: left}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "SplitLeftRight right", Expected: "val", Actual: right}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_SplitLeftRight_NoSep(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SplitLeftRight_NoSep", func() {
		sso := corestr.New.SimpleStringOnce.Init("nosep")
		left, right := sso.SplitLeftRight("=")
		tc := caseV1Compat{Name: "SplitLeftRight no sep left", Expected: "nosep", Actual: left}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "SplitLeftRight no sep right", Expected: "", Actual: right}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_SplitLeftRightTrim_HasBoth(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SplitLeftRightTrim_HasBoth", func() {
		sso := corestr.New.SimpleStringOnce.Init("  key = val  ")
		left, right := sso.SplitLeftRightTrim("=")
		tc := caseV1Compat{Name: "SplitLeftRightTrim left", Expected: "key", Actual: left}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "SplitLeftRightTrim right", Expected: "val", Actual: right}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_SplitLeftRightTrim_NoSep(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SplitLeftRightTrim_NoSep", func() {
		sso := corestr.New.SimpleStringOnce.Init("  nosep  ")
		left, right := sso.SplitLeftRightTrim("=")
		tc := caseV1Compat{Name: "SplitLeftRightTrim no sep left", Expected: "nosep", Actual: left}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "SplitLeftRightTrim no sep right", Expected: "", Actual: right}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_LinesSimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_LinesSimpleSlice", func() {
		sso := corestr.New.SimpleStringOnce.Init("a\nb\nc")
		ss := sso.LinesSimpleSlice()
		tc := caseV1Compat{Name: "LinesSimpleSlice", Expected: 3, Actual: ss.Length()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SimpleSlice(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SimpleSlice", func() {
		sso := corestr.New.SimpleStringOnce.Init("a,b,c")
		ss := sso.SimpleSlice(",")
		tc := caseV1Compat{Name: "SimpleSlice", Expected: 3, Actual: ss.Length()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Split(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Split", func() {
		sso := corestr.New.SimpleStringOnce.Init("a-b-c")
		result := sso.Split("-")
		tc := caseV1Compat{Name: "Split", Expected: 3, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SplitNonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SplitNonEmpty", func() {
		sso := corestr.New.SimpleStringOnce.Init("a,,b")
		result := sso.SplitNonEmpty(",")
		tc := caseV1Compat{Name: "SplitNonEmpty", Expected: true, Actual: len(result) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SplitTrimNonWhitespace(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SplitTrimNonWhitespace", func() {
		sso := corestr.New.SimpleStringOnce.Init("a, ,b")
		result := sso.SplitTrimNonWhitespace(",")
		tc := caseV1Compat{Name: "SplitTrimNonWhitespace", Expected: true, Actual: len(result) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — String/Match methods
// ═══════════════════════════════════════════════════════════════

func Test_SSO_Is(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Is", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		tc := caseV1Compat{Name: "Is match", Expected: true, Actual: sso.Is("hello")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsAnyOf_Found(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsAnyOf_Found", func() {
		sso := corestr.New.SimpleStringOnce.Init("b")
		tc := caseV1Compat{Name: "IsAnyOf found", Expected: true, Actual: sso.IsAnyOf("a", "b", "c")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsAnyOf_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsAnyOf_NotFound", func() {
		sso := corestr.New.SimpleStringOnce.Init("z")
		tc := caseV1Compat{Name: "IsAnyOf not found", Expected: false, Actual: sso.IsAnyOf("a", "b")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsAnyOf_Empty(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsAnyOf_Empty", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "IsAnyOf empty returns true", Expected: true, Actual: sso.IsAnyOf()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsContains(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsContains", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello world")
		tc := caseV1Compat{Name: "IsContains", Expected: true, Actual: sso.IsContains("world")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsAnyContains_Found(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsAnyContains_Found", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello world")
		tc := caseV1Compat{Name: "IsAnyContains found", Expected: true, Actual: sso.IsAnyContains("xyz", "world")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsAnyContains_NotFound(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsAnyContains_NotFound", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")
		tc := caseV1Compat{Name: "IsAnyContains not found", Expected: false, Actual: sso.IsAnyContains("xyz", "abc")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsAnyContains_Empty(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsAnyContains_Empty", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		tc := caseV1Compat{Name: "IsAnyContains empty", Expected: true, Actual: sso.IsAnyContains()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsEqualNonSensitive(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsEqualNonSensitive", func() {
		sso := corestr.New.SimpleStringOnce.Init("Hello")
		tc := caseV1Compat{Name: "IsEqualNonSensitive", Expected: true, Actual: sso.IsEqualNonSensitive("hello")}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsRegexMatches_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsRegexMatches_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc123")
		re := regexp.MustCompile(`\d+`)
		tc := caseV1Compat{Name: "IsRegexMatches valid", Expected: true, Actual: sso.IsRegexMatches(re)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsRegexMatches_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsRegexMatches_Nil", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		tc := caseV1Compat{Name: "IsRegexMatches nil", Expected: false, Actual: sso.IsRegexMatches(nil)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_RegexFindString_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_RegexFindString_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc123def")
		re := regexp.MustCompile(`\d+`)
		tc := caseV1Compat{Name: "RegexFindString", Expected: "123", Actual: sso.RegexFindString(re)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_RegexFindString_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_RegexFindString_Nil", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		tc := caseV1Compat{Name: "RegexFindString nil", Expected: "", Actual: sso.RegexFindString(nil)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_RegexFindAllStrings_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_RegexFindAllStrings_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("a1b2c3")
		re := regexp.MustCompile(`\d`)
		result := sso.RegexFindAllStrings(re, -1)
		tc := caseV1Compat{Name: "RegexFindAllStrings", Expected: 3, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_RegexFindAllStrings_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_RegexFindAllStrings_Nil", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		result := sso.RegexFindAllStrings(nil, -1)
		tc := caseV1Compat{Name: "RegexFindAllStrings nil", Expected: 0, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_RegexFindAllStringsWithFlag_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_RegexFindAllStringsWithFlag_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("a1b2c3")
		re := regexp.MustCompile(`\d`)
		items, hasAny := sso.RegexFindAllStringsWithFlag(re, -1)
		tc := caseV1Compat{Name: "RegexFindAllStringsWithFlag hasAny", Expected: true, Actual: hasAny}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "RegexFindAllStringsWithFlag count", Expected: 3, Actual: len(items)}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_RegexFindAllStringsWithFlag_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_RegexFindAllStringsWithFlag_Nil", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		_, hasAny := sso.RegexFindAllStringsWithFlag(nil, -1)
		tc := caseV1Compat{Name: "RegexFindAllStringsWithFlag nil", Expected: false, Actual: hasAny}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Numeric conversions
// ═══════════════════════════════════════════════════════════════

func Test_SSO_Int16_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Int16_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("100")
		tc := caseV1Compat{Name: "Int16 valid", Expected: int16(100), Actual: sso.Int16()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int16_OutOfRange(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Int16_OutOfRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("99999")
		tc := caseV1Compat{Name: "Int16 out of range", Expected: int16(0), Actual: sso.Int16()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int16_Error(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Int16_Error", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		tc := caseV1Compat{Name: "Int16 error", Expected: int16(0), Actual: sso.Int16()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int32_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Int32_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("5000")
		tc := caseV1Compat{Name: "Int32 valid", Expected: int32(5000), Actual: sso.Int32()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Int32_Error(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Int32_Error", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		tc := caseV1Compat{Name: "Int32 error", Expected: int32(0), Actual: sso.Int32()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Uint16_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Uint16_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("500")
		val, isInRange := sso.Uint16()
		tc := caseV1Compat{Name: "Uint16 valid", Expected: uint16(500), Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "Uint16 inRange", Expected: true, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_Uint32_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Uint32_Valid", func() {
		sso := corestr.New.SimpleStringOnce.Init("1000")
		val, isInRange := sso.Uint32()
		tc := caseV1Compat{Name: "Uint32 valid", Expected: uint32(1000), Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "Uint32 inRange", Expected: true, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_InRange(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_WithinRange_InRange", func() {
		sso := corestr.New.SimpleStringOnce.Init("5")
		val, isInRange := sso.WithinRange(true, 0, 10)
		tc := caseV1Compat{Name: "WithinRange inRange", Expected: 5, Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "WithinRange isInRange", Expected: true, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_BelowMin(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_WithinRange_BelowMin", func() {
		sso := corestr.New.SimpleStringOnce.Init("-5")
		val, isInRange := sso.WithinRange(true, 0, 10)
		tc := caseV1Compat{Name: "WithinRange below min", Expected: 0, Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "WithinRange below isInRange", Expected: false, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_AboveMax(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_WithinRange_AboveMax", func() {
		sso := corestr.New.SimpleStringOnce.Init("15")
		val, isInRange := sso.WithinRange(true, 0, 10)
		tc := caseV1Compat{Name: "WithinRange above max", Expected: 10, Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "WithinRange above isInRange", Expected: false, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_NoBoundary(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_WithinRange_NoBoundary", func() {
		sso := corestr.New.SimpleStringOnce.Init("-5")
		val, isInRange := sso.WithinRange(false, 0, 10)
		tc := caseV1Compat{Name: "WithinRange noBoundary val", Expected: -5, Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "WithinRange noBoundary isInRange", Expected: false, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRange_ParseErr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_WithinRange_ParseErr", func() {
		sso := corestr.New.SimpleStringOnce.Init("abc")
		val, isInRange := sso.WithinRange(true, 0, 10)
		tc := caseV1Compat{Name: "WithinRange parseErr", Expected: 0, Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "WithinRange parseErr isInRange", Expected: false, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_WithinRangeDefault(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_WithinRangeDefault", func() {
		sso := corestr.New.SimpleStringOnce.Init("5")
		val, isInRange := sso.WithinRangeDefault(0, 10)
		tc := caseV1Compat{Name: "WithinRangeDefault", Expected: 5, Actual: val}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "WithinRangeDefault isInRange", Expected: true, Actual: isInRange}
		tc2.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Boolean / IsSetter
// ═══════════════════════════════════════════════════════════════

func Test_SSO_BooleanDefault_True(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_BooleanDefault_True", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		tc := caseV1Compat{Name: "BooleanDefault true", Expected: true, Actual: sso.BooleanDefault()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Boolean_ConsiderInit_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Boolean_ConsiderInit_Uninit", func() {
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")
		tc := caseV1Compat{Name: "Boolean considerInit uninit", Expected: false, Actual: sso.Boolean(true)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Boolean_YesValues(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Boolean_YesValues", func() {
		for _, val := range []string{"yes", "y", "1", "YES", "Y"} {
			sso := corestr.New.SimpleStringOnce.Init(val)
			tc := caseV1Compat{Name: "Boolean " + val, Expected: true, Actual: sso.Boolean(false)}

		// Assert
			tc.ShouldBeEqual(t)
		}
	})
}

func Test_SSO_Boolean_ParseTrue(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Boolean_ParseTrue", func() {
		sso := corestr.New.SimpleStringOnce.Init("true")
		tc := caseV1Compat{Name: "Boolean parse true", Expected: true, Actual: sso.Boolean(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Boolean_ParseErr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Boolean_ParseErr", func() {
		sso := corestr.New.SimpleStringOnce.Init("notbool")
		tc := caseV1Compat{Name: "Boolean parse err", Expected: false, Actual: sso.Boolean(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsValueBool(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsValueBool", func() {
		sso := corestr.New.SimpleStringOnce.Uninitialized("true")
		tc := caseV1Compat{Name: "IsValueBool", Expected: true, Actual: sso.IsValueBool()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsSetter_True(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsSetter_True", func() {
		sso := corestr.New.SimpleStringOnce.Init("yes")
		tc := caseV1Compat{Name: "IsSetter true", Expected: issetter.True, Actual: sso.IsSetter(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsSetter_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsSetter_Uninit", func() {
		sso := corestr.New.SimpleStringOnce.Uninitialized("yes")
		tc := caseV1Compat{Name: "IsSetter uninit", Expected: issetter.False, Actual: sso.IsSetter(true)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsSetter_ParseErr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsSetter_ParseErr", func() {
		sso := corestr.New.SimpleStringOnce.Init("notbool")
		tc := caseV1Compat{Name: "IsSetter parseErr", Expected: issetter.Uninitialized, Actual: sso.IsSetter(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsSetter_ParseTrue(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsSetter_ParseTrue", func() {
		sso := corestr.New.SimpleStringOnce.Init("true")
		tc := caseV1Compat{Name: "IsSetter parse true", Expected: issetter.True, Actual: sso.IsSetter(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_IsSetter_ParseFalse(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_IsSetter_ParseFalse", func() {
		sso := corestr.New.SimpleStringOnce.Init("false")
		tc := caseV1Compat{Name: "IsSetter parse false", Expected: issetter.False, Actual: sso.IsSetter(false)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// SimpleStringOnce — Clone / String / Dispose / Json
// ═══════════════════════════════════════════════════════════════

func Test_SSO_CloneUsingNewVal(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_CloneUsingNewVal", func() {
		sso := corestr.New.SimpleStringOnce.Init("old")
		cloned := sso.CloneUsingNewVal("new")
		tc := caseV1Compat{Name: "CloneUsingNewVal value", Expected: "new", Actual: cloned.Value()}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "CloneUsingNewVal isInit", Expected: true, Actual: cloned.IsInitialized()}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_Clone(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Clone", func() {
		sso := corestr.New.SimpleStringOnce.Init("val")
		cloned := sso.Clone()
		tc := caseV1Compat{Name: "Clone", Expected: "val", Actual: cloned.Value()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_ClonePtr_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_ClonePtr_Valid", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("val")
		cloned := sso.ClonePtr()
		tc := caseV1Compat{Name: "ClonePtr", Expected: "val", Actual: cloned.Value()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_ClonePtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_ClonePtr_Nil", func() {
		var sso *corestr.SimpleStringOnce
		cloned := sso.ClonePtr()
		tc := caseV1Compat{Name: "ClonePtr nil", Expected: true, Actual: cloned == nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_StringPtr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_StringPtr", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("hello")
		ptr := sso.StringPtr()
		tc := caseV1Compat{Name: "StringPtr", Expected: "hello", Actual: *ptr}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_StringPtr_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_StringPtr_Nil", func() {
		var sso *corestr.SimpleStringOnce
		ptr := sso.StringPtr()
		tc := caseV1Compat{Name: "StringPtr nil", Expected: "", Actual: *ptr}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Dispose(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Dispose", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("val")
		sso.Dispose()
		tc := caseV1Compat{Name: "Dispose value empty", Expected: "", Actual: sso.Value()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Dispose_Nil(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Dispose_Nil", func() {
		var sso *corestr.SimpleStringOnce
		sso.Dispose() // should not panic
		tc := caseV1Compat{Name: "Dispose nil no panic", Expected: true, Actual: true}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_NonPtr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_NonPtr", func() {
		sso := corestr.New.SimpleStringOnce.Init("val")
		np := sso.NonPtr()
		tc := caseV1Compat{Name: "NonPtr", Expected: "val", Actual: np.Value()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Ptr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Ptr", func() {
		sso := corestr.New.SimpleStringOnce.Init("val")
		p := sso.Ptr()
		tc := caseV1Compat{Name: "Ptr not nil", Expected: true, Actual: p != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SafeValue_Init(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SafeValue_Init", func() {
		sso := corestr.New.SimpleStringOnce.Init("test")
		tc := caseV1Compat{Name: "SafeValue init", Expected: "test", Actual: sso.SafeValue()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_SafeValue_Uninit(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_SafeValue_Uninit", func() {
		sso := corestr.New.SimpleStringOnce.Uninitialized("test")
		tc := caseV1Compat{Name: "SafeValue uninit", Expected: "", Actual: sso.SafeValue()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_Json(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Json", func() {
		sso := corestr.New.SimpleStringOnce.Init("jsonval")
		jsonResult := sso.Json()
		tc := caseV1Compat{Name: "Json not empty", Expected: true, Actual: jsonResult.HasAnyItem()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_JsonPtr(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_JsonPtr", func() {
		sso := corestr.New.SimpleStringOnce.Init("jsonval")
		ptr := sso.JsonPtr()
		tc := caseV1Compat{Name: "JsonPtr not nil", Expected: true, Actual: ptr != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_MarshalUnmarshal(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_MarshalUnmarshal", func() {
		sso := corestr.New.SimpleStringOnce.Init("marshal")
		data, err := sso.MarshalJSON()
		tc := caseV1Compat{Name: "MarshalJSON no err", Expected: true, Actual: err == nil}

		// Assert
		tc.ShouldBeEqual(t)

		sso2 := corestr.New.SimpleStringOnce.Empty()
		err2 := sso2.UnmarshalJSON(data)
		tc2 := caseV1Compat{Name: "UnmarshalJSON no err", Expected: true, Actual: err2 == nil}
		tc2.ShouldBeEqual(t)
		tc3 := caseV1Compat{Name: "UnmarshalJSON value", Expected: "marshal", Actual: sso2.Value()}
		tc3.ShouldBeEqual(t)
	})
}

func Test_SSO_Serialize(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_Serialize", func() {
		sso := corestr.New.SimpleStringOnce.Init("serval")
		data, err := sso.Serialize()
		tc := caseV1Compat{Name: "Serialize no err", Expected: true, Actual: err == nil}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "Serialize has data", Expected: true, Actual: len(data) > 0}
		tc2.ShouldBeEqual(t)
	})
}

func Test_SSO_AsJsonContractsBinder(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_AsJsonContractsBinder", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("v")
		binder := sso.AsJsonContractsBinder()
		tc := caseV1Compat{Name: "AsJsonContractsBinder", Expected: true, Actual: binder != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_AsJsoner(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_AsJsoner", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("v")
		j := sso.AsJsoner()
		tc := caseV1Compat{Name: "AsJsoner", Expected: true, Actual: j != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_AsJsonParseSelfInjector(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_AsJsonParseSelfInjector", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("v")
		inj := sso.AsJsonParseSelfInjector()
		tc := caseV1Compat{Name: "AsJsonParseSelfInjector", Expected: true, Actual: inj != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_SSO_AsJsonMarshaller(t *testing.T) {
	safeTest(t, "Test_Cov44_SSO_AsJsonMarshaller", func() {
		sso := corestr.New.SimpleStringOnce.InitPtr("v")
		m := sso.AsJsonMarshaller()
		tc := caseV1Compat{Name: "AsJsonMarshaller", Expected: true, Actual: m != nil}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// ValidValue — Factory functions
// ═══════════════════════════════════════════════════════════════

func Test_NewValidValueUsingAny_WithFieldName(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValueUsingAny_WithFieldName", func() {
		vv := corestr.NewValidValueUsingAny(true, true, "hello")
		tc := caseV1Compat{Name: "NewValidValueUsingAny with field name", Expected: true, Actual: vv.IsValid}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "NewValidValueUsingAny value not empty", Expected: true, Actual: len(vv.Value) > 0}
		tc2.ShouldBeEqual(t)
	})
}

func Test_NewValidValueUsingAny_WithoutFieldName(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValueUsingAny_WithoutFieldName", func() {
		vv := corestr.NewValidValueUsingAny(false, false, 42)
		tc := caseV1Compat{Name: "NewValidValueUsingAny without field name", Expected: false, Actual: vv.IsValid}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_NewValidValueUsingAnyAutoValid_NonEmpty(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValueUsingAnyAutoValid_NonEmpty", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "hello")
		tc := caseV1Compat{Name: "NewValidValueUsingAnyAutoValid non-empty", Expected: true, Actual: len(vv.Value) > 0}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_NewValidValueUsingAnyAutoValid_Empty(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValueUsingAnyAutoValid_Empty", func() {
		vv := corestr.NewValidValueUsingAnyAutoValid(false, "")
		tc := caseV1Compat{Name: "NewValidValueUsingAnyAutoValid empty", Expected: true, Actual: vv.IsValid}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_NewValidValue(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValue", func() {
		vv := corestr.NewValidValue("test")
		tc := caseV1Compat{Name: "NewValidValue", Expected: "test", Actual: vv.Value}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "NewValidValue isValid", Expected: true, Actual: vv.IsValid}
		tc2.ShouldBeEqual(t)
	})
}

func Test_NewValidValueEmpty(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValueEmpty", func() {
		vv := corestr.NewValidValueEmpty()
		tc := caseV1Compat{Name: "NewValidValueEmpty", Expected: "", Actual: vv.Value}

		// Assert
		tc.ShouldBeEqual(t)
		tc2 := caseV1Compat{Name: "NewValidValueEmpty isValid", Expected: true, Actual: vv.IsValid}
		tc2.ShouldBeEqual(t)
	})
}

func Test_InvalidValidValueNoMessage(t *testing.T) {
	safeTest(t, "Test_Cov44_InvalidValidValueNoMessage", func() {
		vv := corestr.InvalidValidValueNoMessage()
		tc := caseV1Compat{Name: "InvalidValidValueNoMessage", Expected: false, Actual: vv.IsValid}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_ValidValue_ValueBytesOncePtr(t *testing.T) {
	safeTest(t, "Test_Cov44_ValidValue_ValueBytesOncePtr", func() {
		vv := corestr.NewValidValue("bytes")
		result := vv.ValueBytesOncePtr()
		tc := caseV1Compat{Name: "ValueBytesOncePtr", Expected: 5, Actual: len(result)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

// ═══════════════════════════════════════════════════════════════
// ValidValues — factory
// ═══════════════════════════════════════════════════════════════

func Test_NewValidValuesUsingValues_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValuesUsingValues_Valid", func() {
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		v2 := corestr.ValidValue{Value: "b", IsValid: true}
		vv := corestr.NewValidValuesUsingValues(v1, v2)
		tc := caseV1Compat{Name: "NewValidValuesUsingValues", Expected: 2, Actual: vv.Length()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_NewValidValuesUsingValues_Empty(t *testing.T) {
	safeTest(t, "Test_Cov44_NewValidValuesUsingValues_Empty", func() {
		vv := corestr.NewValidValuesUsingValues()
		tc := caseV1Compat{Name: "NewValidValuesUsingValues empty", Expected: 0, Actual: vv.Length()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_ValidValues_Count(t *testing.T) {
	safeTest(t, "Test_Cov44_ValidValues_Count", func() {
		vv := corestr.NewValidValues(5)
		tc := caseV1Compat{Name: "Count", Expected: 0, Actual: vv.Count()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_ValidValues_HasAnyItem(t *testing.T) {
	safeTest(t, "Test_Cov44_ValidValues_HasAnyItem", func() {
		vv := corestr.EmptyValidValues()
		tc := caseV1Compat{Name: "HasAnyItem empty", Expected: false, Actual: vv.HasAnyItem()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_ValidValues_LastIndex(t *testing.T) {
	safeTest(t, "Test_Cov44_ValidValues_LastIndex", func() {
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vv := corestr.NewValidValuesUsingValues(v1)
		tc := caseV1Compat{Name: "LastIndex", Expected: 0, Actual: vv.LastIndex()}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_ValidValues_HasIndex_Valid(t *testing.T) {
	safeTest(t, "Test_Cov44_ValidValues_HasIndex_Valid", func() {
		v1 := corestr.ValidValue{Value: "a", IsValid: true}
		vv := corestr.NewValidValuesUsingValues(v1)
		tc := caseV1Compat{Name: "HasIndex valid", Expected: true, Actual: vv.HasIndex(0)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}

func Test_ValidValues_HasIndex_Invalid(t *testing.T) {
	safeTest(t, "Test_Cov44_ValidValues_HasIndex_Invalid", func() {
		vv := corestr.EmptyValidValues()
		tc := caseV1Compat{Name: "HasIndex invalid", Expected: false, Actual: vv.HasIndex(0)}

		// Assert
		tc.ShouldBeEqual(t)
	})
}
