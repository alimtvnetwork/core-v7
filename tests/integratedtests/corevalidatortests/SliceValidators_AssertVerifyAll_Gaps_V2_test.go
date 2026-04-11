package corevalidatortests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ── SliceValidators.AssertVerifyAll: non-empty with error (lines 114-120) ──

func Test_SliceValidators_AssertVerifyAll_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"actual"},
				ExpectedLines: []string{"expected"},
			},
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0,
		Header:    "AssertVerifyAll error test",
	}

	// Act — use VerifyAllError to avoid GoConvey assertion failure
	err := sv.VerifyAllError(params)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}
	tc := coretestcases.CaseV1{
		Title: "SliceValidators VerifyAllError returns error -- mismatched lines",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SliceValidators.AssertVerifyAllUsingActual: empty (line 163-165) ──

func Test_SliceValidators_AssertVerifyAllUsingActual_Empty(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	params := &corevalidator.Parameter{
		CaseIndex: 0,
		Header:    "empty validators",
	}

	// Act
	sv.AssertVerifyAllUsingActual(t, params, "line1")

	// Assert — no panic, returns early
	actual := args.Map{
		"isEmpty": fmt.Sprintf("%v", sv.IsEmpty()),
	}
	tc := coretestcases.CaseV1{
		Title: "AssertVerifyAllUsingActual returns early -- empty validators",
		ExpectedInput: args.Map{
			"isEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SliceValidators.VerifyAllErrorUsingActual: error path (lines 192-216) ──

func Test_SliceValidators_VerifyAllErrorUsingActual_WithError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"expected"},
			},
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "VerifyAllErrorUsingActual test",
		IsAttachUserInputs: true,
	}

	// Act
	err := sv.VerifyAllErrorUsingActual(params, "actual-line")

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}
	tc := coretestcases.CaseV1{
		Title: "VerifyAllErrorUsingActual returns error -- mismatched lines",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── HeaderSliceValidators.AssertVerifyAll: non-empty with error (lines 112-118) ──

func Test_HeaderSliceValidators_AssertVerifyAll_WithError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"actual"},
				ExpectedLines: []string{"expected"},
			},
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex: 0,
		Header:    "HeaderSlice AssertVerifyAll error test",
	}

	// Act — use VerifyAllError to avoid GoConvey assertion failure
	err := hsv.VerifyAllError(params)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}
	tc := coretestcases.CaseV1{
		Title: "HeaderSliceValidators VerifyAllError returns error -- mismatched lines",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── HeaderSliceValidators.AssertVerifyAllUsingActual: empty (line 161-163) ──

func Test_HeaderSliceValidators_AssertVerifyAllUsingActual_Empty(t *testing.T) {
	// Arrange
	var hsv corevalidator.HeaderSliceValidators
	params := &corevalidator.Parameter{
		CaseIndex: 0,
		Header:    "empty header validators",
	}

	// Act
	hsv.AssertVerifyAllUsingActual(t, params, "line1")

	// Assert — no panic, returns early
	actual := args.Map{
		"isEmpty": fmt.Sprintf("%v", hsv.IsEmpty()),
	}
	tc := coretestcases.CaseV1{
		Title: "AssertVerifyAllUsingActual returns early -- empty header validators",
		ExpectedInput: args.Map{
			"isEmpty": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SliceValidatorMessages.UserInputsMergeWithError: nil err + empty str (line 80-82) ──

func Test_SliceValidator_UserInputsMergeWithError_NilErrEmptyStr(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "",
		IsAttachUserInputs: true,
	}

	// Act
	err := sv.UserInputsMergeWithError(params, nil)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}
	tc := coretestcases.CaseV1{
		Title: "UserInputsMergeWithError returns error -- nil err but non-empty formatting message",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── SliceValidatorVerify.lengthVerifyError: actual>0 comparing=0 (lines 220-232) ──

func Test_SliceValidator_LengthVerify_ActualNotEmptyComparingZero(t *testing.T) {
	// Arrange — actual lines present but no expected lines set
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"line1", "line2"},
		ExpectedLines: []string{}, // comparing length = 0
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "length verify test",
		IsAttachUserInputs: false,
	}

	// Act — AllVerifyError calls lengthVerifyError internally
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{
		"hasError": fmt.Sprintf("%v", err != nil),
	}
	tc := coretestcases.CaseV1{
		Title: "AllVerifyError returns error -- actual not empty but comparing is zero",
		ExpectedInput: args.Map{
			"hasError": "true",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── TextValidator.verifyDetailErrorUsingLineProcessing: nil receiver (line 177-179) ──
// This is unexported and called through verifyDetailError.
// Nil TextValidator is defensive dead code — TextValidators never contains nil entries.
// Accepted gap.
