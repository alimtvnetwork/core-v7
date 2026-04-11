package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// TextValidator.IsMatch — Equal
// ==========================================

func Test_TextValidator_IsMatch_ExactEqual(t *testing.T) {
	// Arrange
	tc := tvIsMatchExactEqualTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_ExactNotEqual(t *testing.T) {
	// Arrange
	tc := tvIsMatchExactNotEqualTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_CaseInsensitive_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvIsMatchCaseInsensitiveTestCase
	v := corevalidator.TextValidator{
		Search: "Hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello", false)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_CaseSensitiveFail(t *testing.T) {
	// Arrange
	tc := tvIsMatchCaseSensitiveFailTestCase
	v := corevalidator.TextValidator{
		Search: "Hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — with Trim
// ==========================================

func Test_TextValidator_IsMatch_TrimMatch(t *testing.T) {
	// Arrange
	tc := tvIsMatchTrimTestCase
	v := corevalidator.TextValidator{
		Search: "  hello  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_TrimBothSides(t *testing.T) {
	// Arrange
	tc := tvIsMatchTrimBothTestCase
	v := corevalidator.TextValidator{
		Search: "  hello  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("  hello  ", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — Contains
// ==========================================

func Test_TextValidator_IsMatch_Contains_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvIsMatchContainsTestCase
	v := corevalidator.TextValidator{
		Search: "ell", SearchAs: stringcompareas.Contains,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_ContainsMissing(t *testing.T) {
	// Arrange
	tc := tvIsMatchContainsMissingTestCase
	v := corevalidator.TextValidator{
		Search: "xyz", SearchAs: stringcompareas.Contains,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — NotEqual
// ==========================================

func Test_TextValidator_IsMatch_NotEqual_Different(t *testing.T) {
	// Arrange
	tc := tvIsMatchNotEqualDifferentTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.NotEqual,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_NotEqual_Same(t *testing.T) {
	// Arrange
	tc := tvIsMatchNotEqualSameTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.NotEqual,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — UniqueWords + Sort
// ==========================================

func Test_TextValidator_IsMatch_UniqueWordsSorted(t *testing.T) {
	// Arrange
	tc := tvIsMatchUniqueWordsSortedTestCase
	v := corevalidator.TextValidator{
		Search: "  banana  apple  apple  cherry  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultUniqueWordsCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("cherry banana apple", false)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatch — Empty strings
// ==========================================

func Test_TextValidator_IsMatch_EmptySearchEmptyContent(t *testing.T) {
	// Arrange
	tc := tvIsMatchEmptyBothTestCase
	v := corevalidator.TextValidator{
		Search: "", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatch_EmptySearchNonEmptyContent(t *testing.T) {
	// Arrange
	tc := tvIsMatchEmptySearchNonEmptyTestCase
	v := corevalidator.TextValidator{
		Search: "", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.IsMatchMany
// ==========================================

func Test_TextValidator_IsMatchMany_AllMatch_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvIsMatchManyAllTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(false, true, "hello", "hello", "hello")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatchMany_OneFails_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvIsMatchManyOneFailsTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(false, true, "hello", "world", "hello")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_IsMatchMany_EmptySkip_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvIsMatchManyEmptySkipTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(true, true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.VerifyDetailError
// ==========================================

func Test_TextValidator_VerifyDetailError_Match_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvVerifyDetailMatchTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyDetailError(params, "hello") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_VerifyDetailError_Mismatch_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvVerifyDetailMismatchTestCase
	v := corevalidator.TextValidator{
		Search: "hello", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyDetailError(params, "world") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to TextValidator_NilReceiver_testcases.go)

// ==========================================
// TextValidator.VerifyMany
// ==========================================

func Test_TextValidator_VerifyMany_FirstOnly_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvVerifyManyFirstOnlyTestCase
	v := corevalidator.TextValidator{
		Search: "x", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyMany(false, params, "a", "b", "c") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_VerifyMany_AllErrors(t *testing.T) {
	// Arrange
	tc := tvVerifyManyAllErrorsTestCase
	v := corevalidator.TextValidator{
		Search: "x", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0, Header: "test", IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyMany(true, params, "a", "b") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidator_VerifyFirstError_EmptySkip_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvVerifyFirstEmptySkipTestCase
	v := corevalidator.TextValidator{
		Search: "x", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	}
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyFirstError(params) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidator.SearchTextFinalized — caching
// ==========================================

func Test_TextValidator_SearchTextFinalized_Cached_FromTextValidator(t *testing.T) {
	// Arrange
	tc := tvSearchTextFinalizedTestCase
	v := corevalidator.TextValidator{
		Search: "  hello  ", SearchAs: stringcompareas.Equal,
		Condition: corevalidator.DefaultTrimCoreCondition,
	}
	first := v.SearchTextFinalized()
	second := v.SearchTextFinalized()

	// Act
	actual := args.Map{
		"isCached": first == second,
		"value":    first,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// EmptyValidator preset
// ==========================================

func Test_EmptyValidator_MatchesEmpty(t *testing.T) {
	// Arrange
	tc := tvEmptyMatchesEmptyTestCase

	// Act
	actual := args.Map{"isMatch": corevalidator.EmptyValidator.IsMatch("", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyValidator_MatchesTrimmedEmpty(t *testing.T) {
	// Arrange
	tc := tvEmptyMatchesTrimmedTestCase

	// Act
	actual := args.Map{"isMatch": corevalidator.EmptyValidator.IsMatch("   ", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_EmptyValidator_NoMatchNonEmpty(t *testing.T) {
	// Arrange
	tc := tvEmptyNoMatchNonEmptyTestCase

	// Act
	actual := args.Map{"isMatch": corevalidator.EmptyValidator.IsMatch("hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
