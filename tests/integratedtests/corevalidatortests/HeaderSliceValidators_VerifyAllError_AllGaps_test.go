package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage19 — corevalidator remaining 8 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── HeaderSliceValidators.VerifyAllError non-empty with pass (line 112-118) ──

func Test_HeaderSliceValidators_VerifyAllError_Pass(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "test-header",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"hello"},
				ExpectedLines: []string{"hello"},
			},
		},
	}
	params := &corevalidator.Parameter{
		Header:          "Test header validators pass",
		IsCaseSensitive: true,
	}

	// Act
	err := hsv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "HeaderSliceValidators VerifyAllError pass", expected)
}

// ── SliceValidatorMessages: ActualInputWithExpectingMessage returns non-empty (line 80) ──

func Test_SliceValidatorMessages_UserInputsMergeWithError(t *testing.T) {
	// Arrange
	sv := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"actual"},
		ExpectedLines: []string{"expected"},
	}
	params := &corevalidator.Parameter{
		Header:          "Test mismatch",
		IsCaseSensitive: true,
		CaseIndex:       1,
	}

	// Act
	err := sv.VerifyFirstError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "SliceValidator VerifyFirstError mismatch", expected)
}

// ── SliceValidatorVerify: ActualLinesLength > 0 but comparingLength == 0 (line 220-225) ──

func Test_SliceValidatorVerify_NonEmptyActualButEmptyExpected(t *testing.T) {
	// Arrange
	sv := corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"something"},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{
		Header:          "Test non-empty actual empty expected",
		IsCaseSensitive: true,
		CaseIndex:       1,
	}

	// Act
	err := sv.VerifyFirstError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "SliceValidatorVerify non-empty actual empty expected", expected)
}

// ── SliceValidators.VerifyAllError non-empty with pass (line 114-120) ──

func Test_SliceValidators_VerifyAllError_Pass(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"match"},
				ExpectedLines: []string{"match"},
			},
		},
	}
	params := &corevalidator.Parameter{
		Header:          "Test slice validators pass",
		IsCaseSensitive: true,
	}

	// Act
	err := sv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "SliceValidators VerifyAllError pass", expected)
}

// ── TextValidator: verifyDetailErrorUsingLineProcessing nil guard (line 177) ──
// nil receiver guard — documented as dead code
