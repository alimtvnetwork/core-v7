package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage17 — corevalidator remaining gaps (28 uncovered lines)
// ══════════════════════════════════════════════════════════════════════════════

// --- SliceValidators.AssertVerifyAll with empty validators ---

func Test_SliceValidators_AssertVerifyAll_Empty(t *testing.T) {
	// Arrange
	validators := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{
		Header: "test header",
	}

	// Act & Assert — should return early without error
	convey.Convey("SliceValidators.AssertVerifyAll with empty validators skips", t, func() {
		convey.So(func() {
			validators.AssertVerifyAll(t, params)
		}, convey.ShouldNotPanic)
	})
}

// --- SliceValidators.AssertVerifyAllUsingActual with non-empty ---

func Test_SliceValidators_AssertVerifyAllUsingActual_Matching(t *testing.T) {
	// Arrange
	validator := corevalidator.SliceValidator{
		Condition: corevalidator.DefaultTrimCoreCondition,
		CompareAs: stringcompareas.Equal,
		ExpectedLines: []string{
			"line one",
			"line two",
		},
	}
	validators := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{validator},
	}
	params := &corevalidator.Parameter{
		Header:          "verify matching lines",
		IsCaseSensitive: true,
	}

	// Act & Assert
	validators.AssertVerifyAllUsingActual(
		t,
		params,
		"line one",
		"line two",
	)
}

// --- HeaderSliceValidators.AssertVerifyAll with empty ---

func Test_HeaderSliceValidators_AssertVerifyAll_Empty(t *testing.T) {
	// Arrange
	validators := corevalidator.HeaderSliceValidators{}
	params := &corevalidator.Parameter{
		Header: "test header",
	}

	// Act & Assert
	convey.Convey("HeaderSliceValidators.AssertVerifyAll empty skips", t, func() {
		convey.So(func() {
			validators.AssertVerifyAll(t, params)
		}, convey.ShouldNotPanic)
	})
}

// --- HeaderSliceValidators.AssertVerifyAllUsingActual ---

func Test_HeaderSliceValidators_AssertVerifyAllUsingActual_Matching(t *testing.T) {
	// Arrange
	validator := corevalidator.HeaderSliceValidator{
		Header: "subheader",
		SliceValidator: corevalidator.SliceValidator{
			Condition: corevalidator.DefaultTrimCoreCondition,
			CompareAs: stringcompareas.Equal,
			ExpectedLines: []string{
				"hello",
			},
		},
	}
	validators := corevalidator.HeaderSliceValidators{validator}
	params := &corevalidator.Parameter{
		Header:          "verify header matching",
		IsCaseSensitive: true,
	}

	// Act & Assert
	validators.AssertVerifyAllUsingActual(
		t,
		params,
		"hello",
	)
}

// --- SliceValidatorVerify branches ---

func Test_SliceValidator_VerifyError_ActualWithNoExpected(t *testing.T) {
	// Arrange — actual lines present but no expected lines set
	validator := &corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"line one"},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{
		Header:             "test actual with no expected",
		IsAttachUserInputs: true,
		IsCaseSensitive:    true,
	}

	// Act
	err := validator.AllVerifyError(params)

	// Assert
	convey.Convey("VerifyError returns error when actual has lines but expected is empty", t, func() {
		convey.So(err, convey.ShouldNotBeNil)
	})
}

// --- SliceValidatorMessages.UserInputsMergeWithError ---

func Test_SliceValidator_UserInputsMergeWithError_NilErr(t *testing.T) {
	// Arrange
	validator := &corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{
		Header:             "test",
		IsAttachUserInputs: true,
		IsCaseSensitive:    true,
	}

	// Act
	err := validator.AllVerifyError(params)

	// Assert
	convey.Convey("AllVerifyError with empty actual and expected returns nil", t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

// Coverage note: Remaining uncovered lines:
// - TextValidator.verifyDetailErrorUsingLineProcessing nil receiver (line 177) —
//   defensive dead code, method only called from non-nil receivers
// - SliceValidatorMessages line 80 (err==nil && len(toStr)==0) — requires
//   specific state where IsAttachUserInputs is true but no actual/expected mismatch
