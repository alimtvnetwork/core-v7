package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// SliceValidator.IsValid
// ==========================================

func Test_SliceValidator_IsValid_ExactMatch(t *testing.T) {
	// Arrange
	tc := svIsValidExactMatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b", "c"}, ExpectedLines: []string{"a", "b", "c"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_Mismatch(t *testing.T) {
	// Arrange
	tc := svIsValidMismatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b"}, ExpectedLines: []string{"a", "x"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_LengthMismatch(t *testing.T) {
	// Arrange
	tc := svIsValidLengthMismatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a"}, ExpectedLines: []string{"a", "b"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_BothNil(t *testing.T) {
	// Arrange
	tc := svIsValidBothNilTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: nil, ExpectedLines: nil,
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_OneNil(t *testing.T) {
	// Arrange
	tc := svIsValidOneNilTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: nil, ExpectedLines: []string{"a"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver tests migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_IsValid_BothEmpty(t *testing.T) {
	// Arrange
	tc := svIsValidBothEmptyTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: []string{}, ExpectedLines: []string{},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.IsValid — with Trim
// ==========================================

func Test_SliceValidator_IsValid_TrimMatch(t *testing.T) {
	// Arrange
	tc := svIsValidTrimMatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultTrimCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"  hello  ", " world "}, ExpectedLines: []string{"hello", "world"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.IsValid — Contains
// ==========================================

func Test_SliceValidator_IsValid_Contains(t *testing.T) {
	// Arrange
	tc := svIsValidContainsTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Contains,
		ActualLines: []string{"hello world", "foo bar"}, ExpectedLines: []string{"ello", "bar"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator — helper methods
// ==========================================

func Test_SliceValidator_ActualLinesLength(t *testing.T) {
	// Arrange
	tc := svActualLinesLengthTestCase
	v := corevalidator.SliceValidator{ActualLines: []string{"a", "b"}}

	// Act
	actual := args.Map{"length": v.ActualLinesLength()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_ExpectingLinesLength(t *testing.T) {
	// Arrange
	tc := svExpectingLinesLengthTestCase
	v := corevalidator.SliceValidator{ExpectedLines: []string{"a", "b", "c"}}

	// Act
	actual := args.Map{"length": v.ExpectingLinesLength()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsUsedAlready_False(t *testing.T) {
	// Arrange
	tc := svIsUsedAlreadyFalseTestCase
	v := corevalidator.SliceValidator{ExpectedLines: []string{"a"}}

	// Act
	actual := args.Map{"isUsed": v.IsUsedAlready()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsUsedAlready_TrueAfterComparing(t *testing.T) {
	// Arrange
	tc := svIsUsedAlreadyTrueTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"}, ActualLines: []string{"a"},
	}
	_ = v.ComparingValidators()

	// Act
	actual := args.Map{"isUsed": v.IsUsedAlready()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_MethodName(t *testing.T) {
	// Arrange
	tc := svMethodNameTestCase
	v := corevalidator.SliceValidator{CompareAs: stringcompareas.Contains}

	// Act
	actual := args.Map{"name": v.MethodName()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.SetActual / SetActualVsExpected
// ==========================================

func Test_SliceValidator_SetActual(t *testing.T) {
	// Arrange
	tc := svSetActualTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"},
	}
	v.SetActual([]string{"a"})

	// Act
	actual := args.Map{"length": v.ActualLinesLength()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_SetActualVsExpected(t *testing.T) {
	// Arrange
	tc := svSetActualVsExpectedTestCase
	v := corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	v.SetActualVsExpected([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{
		"actualLen":   v.ActualLinesLength(),
		"expectedLen": v.ExpectingLinesLength(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.IsValidOtherLines
// ==========================================

func Test_SliceValidator_IsValidOtherLines_Match(t *testing.T) {
	// Arrange
	tc := svIsValidOtherLinesMatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValidOtherLines(true, []string{"a", "b"})}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValidOtherLines_Mismatch(t *testing.T) {
	// Arrange
	tc := svIsValidOtherLinesMismatchTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValidOtherLines(true, []string{"a", "x"})}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.AllVerifyError
// ==========================================

func Test_SliceValidator_AllVerifyError_Pass(t *testing.T) {
	// Arrange
	tc := svAllVerifyErrorPassTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b"}, ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}

	// Act
	actual := args.Map{"hasError": v.AllVerifyError(params) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_AllVerifyError_Fail(t *testing.T) {
	// Arrange
	tc := svAllVerifyErrorFailTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "x"}, ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "test", IsCaseSensitive: true}

	// Act
	actual := args.Map{"hasError": v.AllVerifyError(params) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

func Test_SliceValidator_AllVerifyError_SkipEmpty(t *testing.T) {
	// Arrange
	tc := svAllVerifyErrorSkipEmptyTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{}, ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, IsSkipCompareOnActualEmpty: true}

	// Act
	actual := args.Map{"hasError": v.AllVerifyError(params) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// SliceValidator.VerifyFirstError
// ==========================================

func Test_SliceValidator_VerifyFirstError_Pass(t *testing.T) {
	// Arrange
	tc := svVerifyFirstErrorPassTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"a", "b"}, ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, IsCaseSensitive: true}

	// Act
	actual := args.Map{"hasError": v.VerifyFirstError(params) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator.Dispose
// ==========================================

func Test_SliceValidator_Dispose(t *testing.T) {
	// Arrange
	tc := svDisposeTestCase
	v := corevalidator.SliceValidator{
		CompareAs: stringcompareas.Equal, ActualLines: []string{"a"}, ExpectedLines: []string{"a"},
	}
	v.Dispose()

	// Act
	actual := args.Map{
		"actualNil":   v.ActualLines == nil,
		"expectedNil": v.ExpectedLines == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to SliceValidator_NilReceiver_testcases.go)

// ==========================================
// SliceValidator — case insensitive
// ==========================================

func Test_SliceValidator_IsValid_CaseInsensitive(t *testing.T) {
	// Arrange
	tc := svCaseInsensitiveTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"Hello", "WORLD"}, ExpectedLines: []string{"hello", "world"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(false)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_SliceValidator_IsValid_CaseSensitiveFail(t *testing.T) {
	// Arrange
	tc := svCaseSensitiveFailTestCase
	v := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultDisabledCoreCondition, CompareAs: stringcompareas.Equal,
		ActualLines: []string{"Hello"}, ExpectedLines: []string{"hello"},
	}

	// Act
	actual := args.Map{"isValid": v.IsValid(true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// NewSliceValidatorUsingErr
// ==========================================

func Test_NewSliceValidatorUsingErr_NilError(t *testing.T) {
	// Arrange
	tc := svNewUsingErrNilTestCase
	v := corevalidator.NewSliceValidatorUsingErr(
		nil, "expected\nlines",
		true, false, false,
		stringcompareas.Equal,
	)

	// Act
	actual := args.Map{
		"isNotNil":  v != nil,
		"actualLen": v.ActualLinesLength(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
