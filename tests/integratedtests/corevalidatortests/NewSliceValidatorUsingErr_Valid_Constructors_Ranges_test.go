package corevalidatortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corerange"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func createSimpleSliceValidator(expectedLines []string) *corevalidator.SimpleSliceValidator {
	return &corevalidator.SimpleSliceValidator{
		Expected:  corestr.New.SimpleSlice.Direct(false, expectedLines),
		CompareAs: stringcompareas.Equal,
	}
}

// ── NewSliceValidatorUsingErr ──

func Test_NewSliceValidatorUsingErr_Valid(t *testing.T) {
	// Arrange
	err := errors.New("line1\nline2")
	expected := "line1\nline2"

	// Act
	sv := corevalidator.NewSliceValidatorUsingErr(
		err, expected, false, false, false, stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{
		"actualLen":   sv.ActualLinesLength(),
		"expectedLen": sv.ExpectingLinesLength(),
	}
	exp := args.Map{
		"actualLen":   2,
		"expectedLen": 2,
	}
	exp.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr creates validator -- 2 lines", actual)
}

func Test_NewSliceValidatorUsingErr_NilErr(t *testing.T) {
	// Arrange & Act
	sv := corevalidator.NewSliceValidatorUsingErr(
		nil, "", false, false, false, stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{"actualLen": sv.ActualLinesLength()}
	exp := args.Map{"actualLen": 0}
	exp.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr creates empty -- nil err", actual)
}

// ── NewSliceValidatorUsingAny ──

func Test_NewSliceValidatorUsingAny_Valid(t *testing.T) {
	// Arrange & Act
	sv := corevalidator.NewSliceValidatorUsingAny(
		"hello\nworld", "hello\nworld", false, false, false, stringcompareas.Equal,
	)

	// Assert
	actual := args.Map{
		"actualLen":   sv.ActualLinesLength(),
		"expectedLen": sv.ExpectingLinesLength(),
	}
	exp := args.Map{
		"actualLen":   2,
		"expectedLen": 2,
	}
	exp.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny creates validator -- 2 lines", actual)
}

// ── RangeSegmentsValidator ──

func Test_RangeSegmentsValidator_VerifySimple_FromNewSliceValidatorUsi(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "Test Range",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	actual := []string{"a", "b", "c"}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := rsv.VerifySimple(actual, params, false)

	// Assert
	result := args.Map{"noErr": err == nil}
	exp := args.Map{"noErr": true}
	exp.ShouldBeEqual(t, 0, "VerifySimple returns nil -- matching range", result)
}

func Test_RangeSegmentsValidator_VerifyFirstDefault_FromNewSliceValidatorUsi(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "Test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 1},
				ExpectedLines: []string{"x"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := rsv.VerifyFirstDefault([]string{"x", "y"}, params, false)

	// Assert
	result := args.Map{"noErr": err == nil}
	exp := args.Map{"noErr": true}
	exp.ShouldBeEqual(t, 0, "VerifyFirstDefault returns nil -- matching", result)
}

func Test_RangeSegmentsValidator_VerifyUptoDefault_FromNewSliceValidatorUsi(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "Test",
		VerifierSegments: []corevalidator.RangesSegment{
			{
				RangeInt:      corerange.RangeInt{Start: 0, End: 2},
				ExpectedLines: []string{"a", "b"},
				CompareAs:     stringcompareas.Equal,
			},
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := rsv.VerifyUptoDefault([]string{"a", "b", "c"}, params, 1, false)

	// Assert
	result := args.Map{"noErr": err == nil}
	exp := args.Map{"noErr": true}
	exp.ShouldBeEqual(t, 0, "VerifyUptoDefault returns nil -- matching upto 1", result)
}

func Test_RangeSegmentsValidator_LengthOfVerifierSegments_FromNewSliceValidatorUsi(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title:            "Test",
		VerifierSegments: []corevalidator.RangesSegment{{}, {}},
	}

	// Act
	result := rsv.LengthOfVerifierSegments()

	// Assert
	actual := args.Map{"len": result}
	exp := args.Map{"len": 2}
	exp.ShouldBeEqual(t, 0, "LengthOfVerifierSegments returns 2 -- 2 segments", actual)
}

// ── SimpleSliceValidator ──

func Test_SimpleSliceValidator_VerifyUpto_FromNewSliceValidatorUsi(t *testing.T) {
	// Arrange
	sv := createSimpleSliceValidator(
		[]string{"a", "b", "c"},
	)
	sv.SetActual([]string{"a", "b", "c"})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := sv.VerifyUpto([]string{"a", "b", "c"}, params, 2)

	// Assert
	actual := args.Map{"noErr": err == nil}
	exp := args.Map{"noErr": true}
	exp.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- matching first 2", actual)
}

func Test_SimpleSliceValidator_VerifyFirst_FromNewSliceValidatorUsi(t *testing.T) {
	// Arrange
	sv := createSimpleSliceValidator(
		[]string{"hello", "world"},
	)
	sv.SetActual([]string{"hello", "world"})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}

	// Act
	err := sv.VerifyFirst([]string{"hello", "world"}, params)

	// Assert
	actual := args.Map{"noErr": err == nil}
	exp := args.Map{"noErr": true}
	exp.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- matching", actual)
}
