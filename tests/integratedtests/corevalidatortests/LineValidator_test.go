package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// LineValidator.IsMatch
// ==========================================

func Test_LineValidator_IsMatch_BothMatch(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchBothTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch(0, "hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_IsMatch_LineNumberMismatch(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchLineMismatchTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch(0, "hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_IsMatch_TextMismatch_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchTextMismatchTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch(0, "world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_IsMatch_SkipLineNumber(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchSkipLineTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatch(99, "hello", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// LineValidator.IsMatchMany
// ==========================================

func Test_LineValidator_IsMatchMany_AllMatch_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchManyAllTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(false, true, items...)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_IsMatchMany_OneFails_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchManyOneFailsTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "nope", LineNumber: 1},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(false, true, items...)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_IsMatchMany_EmptySkip_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorIsMatchManyEmptySkipTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(true, true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver test migrated to LineValidator_NilReceiver_testcases.go)

// ==========================================
// LineValidator.VerifyError
// ==========================================

func Test_LineValidator_VerifyError_Match_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorVerifyErrorMatchTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyError(params, 0, "hello") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_VerifyError_LineNumberMismatch(t *testing.T) {
	// Arrange
	tc := lineValidatorVerifyErrorLineMismatchTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 5},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyError(params, 0, "hello") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_VerifyError_TextMismatch_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorVerifyErrorTextMismatchTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "hello",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "test",
		IsCaseSensitive: true,
	}

	// Act
	actual := args.Map{"hasError": v.VerifyError(params, 0, "world") != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// LineValidator.VerifyMany
// ==========================================

func Test_LineValidator_VerifyMany_ContinueOnError_FromLineValidator(t *testing.T) {
	// Arrange
	tc := lineValidatorVerifyManyContinueTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
	}

	// Act
	actual := args.Map{"hasError": v.VerifyMany(true, params, items...) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_LineValidator_VerifyMany_FirstOnly(t *testing.T) {
	// Arrange
	tc := lineValidatorVerifyManyFirstOnlyTestCase
	v := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
		{Text: "also bad", LineNumber: 1},
	}

	// Act
	actual := args.Map{"hasError": v.VerifyMany(false, params, items...) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
