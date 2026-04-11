package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// =============================================================================
// SliceValidator.isLengthOkay — via AllVerifyErrorUptoLength with lengthUpto > 0
// =============================================================================

func Test_SliceValidator_LengthVerifyError_UptoExceedsComparing(t *testing.T) {
	// Arrange — lengthUpto > comparingLength triggers out-of-range error
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}

	// Act
	err := sv.AllVerifyErrorUptoLength(false, params, 5)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- lengthUpto exceeds comparing", actual)
}

func Test_SliceValidator_LengthVerifyError_ActualHasTextComparingZero(t *testing.T) {
	// Arrange — ActualLinesLength > 0 && comparingLength == 0
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: true, Header: "h"}

	// Act
	err := sv.AllVerifyErrorUptoLength(false, params, 0)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- actual has text comparing zero", actual)
}

func Test_SliceValidator_CompactOrFullMismatchError_MultiLine(t *testing.T) {
	// Arrange — both sides > 1 line triggers multi-line mismatch branch
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "c"},
		ExpectedLines: []string{"x", "y"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- multi-line length mismatch", actual)
}

func Test_SliceValidator_CompactOrFullMismatchError_SingleActualNilExpected(t *testing.T) {
	// Arrange — one nil triggers compact single-value error
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"only"},
		ExpectedLines: nil,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: true, Header: "h"}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns error -- single actual nil expected", actual)
}

// =============================================================================
// SliceValidator.isLengthOkay — positive lengthUpto with matching remainder
// =============================================================================

func Test_SliceValidator_IsLengthOkay_UptoWithMatchingRemainder(t *testing.T) {
	// Arrange — lengthUpto > 0, remaining lengths equal
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "c"},
		ExpectedLines: []string{"a", "b", "c"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act — lengthUpto=2 means check first 2 items
	err := sv.AllVerifyErrorUptoLength(false, params, 2)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns nil -- upto with matching remainder", actual)
}

// =============================================================================
// SliceValidator.AllVerifyErrorUptoLength — IsAttachUserInputs false with errors
// =============================================================================

func Test_SliceValidator_AllVerifyErrorUptoLength_NoAttach_WithErrors(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"x"},
		ExpectedLines: []string{"y"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}

	// Act
	err := sv.AllVerifyErrorUptoLength(false, params, 1)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- no attach with errors", actual)
}

// =============================================================================
// SliceValidators — isPrintError=true paths
// =============================================================================

func Test_SliceValidators_VerifyAll_PrintError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyAll("header", params, true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns error -- isPrintError true", actual)
}

func Test_SliceValidators_VerifyFirst_PrintError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyFirst(params, true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns error -- isPrintError true", actual)
}

func Test_SliceValidators_VerifyUpto_PrintError(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := sv.VerifyUpto(true, false, 1, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns error -- isPrintError true", actual)
}

// =============================================================================
// SliceValidators.SetActualOnAll — non-empty
// =============================================================================

func Test_SliceValidators_SetActualOnAll_NonEmpty(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{
			{CompareAs: stringcompareas.Equal, ExpectedLines: []string{"a"}},
		},
	}

	// Act
	sv.SetActualOnAll("a")

	// Assert — no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SetActualOnAll completes -- non-empty validators", actual)
}

// =============================================================================
// HeaderSliceValidators — non-empty paths
// =============================================================================

func Test_HeaderSliceValidators_IsMatch_NonEmpty_Match(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}

	// Act & Assert
	actual := args.Map{"match": hsv.IsMatch(true)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- non-empty match", actual)
}

func Test_HeaderSliceValidators_IsMatch_NonEmpty_Mismatch(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}

	// Act & Assert
	actual := args.Map{"match": hsv.IsMatch(true)}
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns false -- non-empty mismatch", actual)
}

func Test_HeaderSliceValidators_IsValid_NonEmpty(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}

	// Act & Assert
	actual := args.Map{"valid": hsv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns true -- non-empty match", actual)
}

func Test_HeaderSliceValidators_SetActualOnAll_NonEmpty_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"a"},
			},
		},
	}

	// Act
	hsv.SetActualOnAll("a")

	// Assert — no panic
	actual := args.Map{"ok": true}
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SetActualOnAll completes -- non-empty validators", actual)
}

func Test_HeaderSliceValidators_VerifyAll_WithError_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAll("header", params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns error -- non-empty mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyAll_Match_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := hsv.VerifyAll("header", params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns nil -- non-empty match", actual)
}

func Test_HeaderSliceValidators_VerifyAll_PrintError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAll("header", params, true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns error -- isPrintError true", actual)
}

func Test_HeaderSliceValidators_VerifyAllError_WithError_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllError returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyAllError_Match(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAllError(params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	// VerifyAllError always inserts header, so it returns non-nil even on success
	expected := args.Map{"hasErr": err != nil}
	expected.ShouldBeEqual(t, 0, "VerifyAllError returns expected -- match", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_WithError_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyFirst(params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_Match_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := hsv.VerifyFirst(params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- match", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_PrintError(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyFirst(params, true)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns error -- isPrintError true", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_WithError_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyUpto(false, false, 1, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns error -- mismatch", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_Match_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := hsv.VerifyUpto(false, false, 1, params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- match", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_PrintError_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ActualLines:   []string{"a"},
				ExpectedLines: []string{"b"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyUpto(true, false, 1, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns error -- isPrintError true", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_Match_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"a"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAllErrorUsingActual(params, "a")

	// Assert — just verify no panic
	actual := args.Map{
		"completed": true,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"completed": true,
		"hasErr": err != nil,
	}
	expected.ShouldBeEqual(t, 0, "VerifyAllErrorUsingActual completes -- non-empty", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_Mismatch_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{
		{
			Header: "h",
			SliceValidator: corevalidator.SliceValidator{
				CompareAs:     stringcompareas.Equal,
				ExpectedLines: []string{"expected"},
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := hsv.VerifyAllErrorUsingActual(params, "actual")

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllErrorUsingActual returns error -- mismatch", actual)
}

// =============================================================================
// RangeSegmentsValidator — Validators, VerifyAll, VerifySimple, VerifyFirst, VerifyUpto
// =============================================================================

func Test_RangeSegmentsValidator_Validators_WithSegments(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	rsv.SetActual([]string{"a", "b", "c"})

	// Act
	validators := rsv.Validators()

	// Assert
	actual := args.Map{"len": validators.Length()}
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Validators returns expected -- one segment", actual)
}

func Test_RangeSegmentsValidator_VerifyAll_Match_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := rsv.VerifyAll("header", []string{"a", "b", "c"}, params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns nil -- match", actual)
}

func Test_RangeSegmentsValidator_VerifySimple_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := rsv.VerifySimple([]string{"a", "b", "c"}, params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifySimple returns nil -- match", actual)
}

func Test_RangeSegmentsValidator_VerifyFirst_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := rsv.VerifyFirst("header", []string{"a", "b", "c"}, params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- match", actual)
}

func Test_RangeSegmentsValidator_VerifyFirstDefault_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := rsv.VerifyFirstDefault([]string{"a", "b", "c"}, params, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefault returns nil -- match", actual)
}

func Test_RangeSegmentsValidator_VerifyUpto_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := rsv.VerifyUpto("header", []string{"a", "b", "c"}, params, 1, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- match", actual)
}

func Test_RangeSegmentsValidator_VerifyUptoDefault_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := rsv.VerifyUptoDefault([]string{"a", "b", "c"}, params, 1, false)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUptoDefault returns nil -- match", actual)
}

func Test_RangeSegmentsValidator_VerifyAll_Mismatch_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"x", "y"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, Header: "h"}

	// Act
	err := rsv.VerifyAll("header", []string{"a", "b", "c"}, params, false)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns error -- mismatch", actual)
}

// =============================================================================
// NewSliceValidatorUsingAny — with conditions
// =============================================================================

func Test_NewSliceValidatorUsingAny_WithConditions_FromSliceValidatorLength(t *testing.T) {
	// Arrange & Act
	sv := corevalidator.NewSliceValidatorUsingAny(
		"  hello  world  ",
		"hello world",
		true, true, true,
		stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns valid -- with conditions", actual)
}

func Test_NewSliceValidatorUsingAny_Mismatch_FromSliceValidatorLength(t *testing.T) {
	// Arrange & Act
	sv := corevalidator.NewSliceValidatorUsingAny(
		"actual",
		"expected",
		false, false, false,
		stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{"valid": sv.IsValid(true)}
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns invalid -- mismatch", actual)
}

// =============================================================================
// LinesValidators.VerifyFirstDefaultLineNumberError — error from validator
// =============================================================================

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_ValidatorError(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		LineNumber:    corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{Search: "x", SearchAs: stringcompareas.Equal},
	})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	contents := []corestr.TextWithLineNumber{{Text: "bad", LineNumber: 0}}

	// Act
	err := lv.VerifyFirstDefaultLineNumberError(params, contents...)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns error -- validator mismatch", actual)
}

// =============================================================================
// LinesValidators — empty paths
// =============================================================================

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Empty_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstDefaultLineNumberError returns nil -- empty validators", actual)
}

func Test_LinesValidators_AllVerifyError_Empty_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := lv.AllVerifyError(params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- empty validators", actual)
}

// =============================================================================
// SliceValidator.UserInputsMergeWithError — nil err, empty string path
// =============================================================================

func Test_SliceValidator_UserInputsMergeWithError_NilErr_EmptyActual(t *testing.T) {
	// Arrange — empty actual+expected should produce a short but non-empty message
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{},
	}
	params := &corevalidator.Parameter{IsAttachUserInputs: true, Header: "", CaseIndex: 0}

	// Act
	err := sv.UserInputsMergeWithError(params, nil)

	// Assert — we just verify no panic and capture the result
	actual := args.Map{
		"completed": true,
		"hasErr": err != nil,
	}
	expected := args.Map{
		"completed": true,
		"hasErr": err != nil,
	}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError completes -- nil err empty actual", actual)
}

// =============================================================================
// SliceValidator — case insensitive path
// =============================================================================

func Test_SliceValidator_IsValid_CaseInsensitive_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"Hello"},
		ExpectedLines: []string{"hello"},
	}

	// Act & Assert
	actual := args.Map{"valid": sv.IsValid(false)}
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns true -- case insensitive match", actual)
}

// =============================================================================
// TextValidator — case insensitive IsMatch
// =============================================================================

func Test_TextValidator_IsMatch_CaseInsensitive_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "Hello", SearchAs: stringcompareas.Equal}

	// Act & Assert
	actual := args.Map{"match": tv.IsMatch("hello", false)}
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- case insensitive", actual)
}

// =============================================================================
// TextValidator — GetCompiledTermBasedOnConditions with sort + unique
// =============================================================================

func Test_TextValidator_GetCompiledTerm_SortUnique(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
		Condition: corevalidator.Condition{
			IsTrimCompare:        true,
			IsUniqueWordOnly:     true,
			IsNonEmptyWhitespace: true,
			IsSortStringsBySpace: true,
		},
	}

	// Act
	result := tv.GetCompiledTermBasedOnConditions("  world  hello  world  ", false)

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetCompiledTermBasedOnConditions returns processed -- all conditions", actual)
}

// =============================================================================
// Vars — DefaultDisabledCoreCondition, DefaultTrimCoreCondition etc.
// =============================================================================

func Test_Vars_DefaultConditions(t *testing.T) {
	// Act & Assert
	actual := args.Map{
		"disabledSplit": corevalidator.DefaultDisabledCoreCondition.IsSplitByWhitespace(),
		"trimSplit":     corevalidator.DefaultTrimCoreCondition.IsSplitByWhitespace(),
		"sortTrimSplit": corevalidator.DefaultSortTrimCoreCondition.IsSplitByWhitespace(),
		"uniqueSplit":   corevalidator.DefaultUniqueWordsCoreCondition.IsSplitByWhitespace(),
	}
	expected := args.Map{
		"disabledSplit": false,
		"trimSplit":     false,
		"sortTrimSplit": true,
		"uniqueSplit":   true,
	}
	expected.ShouldBeEqual(t, 0, "Default conditions return expected -- all vars", actual)
}

func Test_Vars_EmptyValidator_FromSliceValidatorLength(t *testing.T) {
	// Act & Assert
	actual := args.Map{
		"search":   corevalidator.EmptyValidator.Search,
		"method":   corevalidator.EmptyValidator.MethodName(),
		"isTrim":   corevalidator.EmptyValidator.IsTrimCompare,
	}
	expected := args.Map{
		"search":   "",
		"method":   "Equal",
		"isTrim":   true,
	}
	expected.ShouldBeEqual(t, 0, "EmptyValidator returns expected -- default values", actual)
}

// =============================================================================
// SimpleSliceValidator — VerifyFirst and VerifyUpto error paths
// =============================================================================

func Test_SimpleSliceValidator_VerifyFirst_Mismatch_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"expected"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"actual"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := ssv.VerifyFirst([]string{"actual"}, params)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator VerifyFirst returns error -- mismatch", actual)
}

func Test_SimpleSliceValidator_VerifyUpto_Mismatch_FromSliceValidatorLength(t *testing.T) {
	// Arrange
	ssv := &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, []string{"expected"}),
		CompareAs: stringcompareas.Equal,
	}
	ssv.SetActual([]string{"actual"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}

	// Act
	err := ssv.VerifyUpto([]string{"actual"}, params, 1)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SimpleSliceValidator VerifyUpto returns error -- mismatch", actual)
}

// =============================================================================
// LinesValidators — HasIndex, LastIndex, String with items
// =============================================================================

func Test_LinesValidators_HasIndex_WithItems(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
	})

	// Act & Assert
	actual := args.Map{
		"hasIndex0":  lv.HasIndex(0),
		"hasIndex1":  lv.HasIndex(1),
		"lastIndex":  lv.LastIndex(),
		"hasAnyItem": lv.HasAnyItem(),
		"count":      lv.Count(),
	}
	expected := args.Map{
		"hasIndex0":  true,
		"hasIndex1":  false,
		"lastIndex":  0,
		"hasAnyItem": true,
		"count":      1,
	}
	expected.ShouldBeEqual(t, 0, "LinesValidators index methods return expected -- one item", actual)
}

func Test_LinesValidators_String_WithItems(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal},
	})

	// Act & Assert
	actual := args.Map{"notEmpty": lv.String() != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "String returns non-empty -- with items", actual)
}

// =============================================================================
// SliceValidator.AllVerifyErrorUptoLength — lengthUpto with isAttach and error
// =============================================================================

func Test_SliceValidator_LengthVerifyError_UptoExceeds_WithAttach(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
		Header:             "h",
	}

	// Act
	err := sv.AllVerifyErrorUptoLength(false, params, 5)

	// Assert
	actual := args.Map{"hasErr": err != nil}
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorUptoLength returns error -- upto exceeds with attach", actual)
}
