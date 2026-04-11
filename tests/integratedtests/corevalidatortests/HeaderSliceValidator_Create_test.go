package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// =============================================================================
// HeaderSliceValidator — uncovered branches
// =============================================================================

func Test_HeaderSliceValidator_Create(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidator{
		Header: "test-header",
		SliceValidator: corevalidator.SliceValidator{
			CompareAs:     stringcompareas.Equal,
			ExpectedLines: []string{"line1", "line2"},
		},
	}

	// Act
	actual := args.Map{
		"header": hsv.Header,
		"linesLen": len(hsv.ExpectedLines),
	}

	// Assert
	expected := args.Map{
		"header": "test-header",
		"linesLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "HeaderSliceValidator returns expected -- valid input", actual)
}

// =============================================================================
// SliceValidator — additional branch coverage
// =============================================================================

func Test_SliceValidator_SetActualVsExpected_Mismatch(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	sv.SetActualVsExpected([]string{"a"}, []string{"b"})

	// Act
	actual := args.Map{"isValid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"isValid": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator SetActualVsExpected returns invalid -- mismatch", actual)
}

func Test_SliceValidator_VerifyFirstError_Valid(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.VerifyFirstError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator VerifyFirstError returns nil -- matching", actual)
}

func Test_SliceValidator_AllVerifyError_Valid(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a"},
	}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator AllVerifyError returns nil -- matching", actual)
}

func Test_SliceValidator_ActualLines_Valid(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"a", "b"},
	}
	sv.SetActualVsExpected([]string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{
		"actualLen":    sv.ActualLinesLength(),
		"expectedLen":  sv.ExpectingLinesLength(),
		"actualStr":    sv.ActualLinesString() != "",
		"expectedStr":  sv.ExpectingLinesString() != "",
		"isUsed":       sv.IsUsedAlready(),
	}

	// Assert
	expected := args.Map{
		"actualLen": 2, "expectedLen": 2,
		"actualStr": true, "expectedStr": true, "isUsed": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator ActualLines returns expected -- 2 lines", actual)
}

// =============================================================================
// TextValidator — additional branch coverage
// =============================================================================

func Test_TextValidator_IsMatch_Equal_True(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- equal", actual)
}

func Test_TextValidator_IsMatch_Equal_False(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hello", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"isMatch": tv.IsMatch("world", true)}

	// Assert
	expected := args.Map{"isMatch": false}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns false -- not equal", actual)
}

func Test_TextValidator_IsMatch_StartsWith(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "hel", SearchAs: stringcompareas.StartsWith}

	// Act
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- starts with", actual)
}

func Test_TextValidator_IsMatch_EndsWith(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "llo", SearchAs: stringcompareas.EndsWith}

	// Act
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- ends with", actual)
}

func Test_TextValidator_IsMatch_Contains_FromHeaderSliceValidator(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{Search: "ell", SearchAs: stringcompareas.Contains}

	// Act
	actual := args.Map{"isMatch": tv.IsMatch("hello", true)}

	// Assert
	expected := args.Map{"isMatch": true}
	expected.ShouldBeEqual(t, 0, "TextValidator IsMatch returns true -- contains", actual)
}

func Test_TextValidator_ToString_SingleLine_FromHeaderSliceValidator(t *testing.T) {
	// Arrange
	tv := corevalidator.TextValidator{Search: "test", SearchAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"notEmpty": tv.ToString(true) != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "TextValidator ToString returns non-empty -- single line", actual)
}

// =============================================================================
// TextValidators — additional branch coverage
// =============================================================================

func Test_TextValidators_IsMatch_WithItems(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{
		"matchTrue":  tvs.IsMatch("hello", true),
		"matchFalse": tvs.IsMatch("world", true),
	}

	// Assert
	expected := args.Map{
		"matchTrue": true,
		"matchFalse": false,
	}
	expected.ShouldBeEqual(t, 0, "TextValidators IsMatch returns expected -- with validator", actual)
}

func Test_TextValidators_VerifyFirstError_WithMatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.VerifyFirstError(0, "hello", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TextValidators VerifyFirstError returns nil -- matching", actual)
}

func Test_TextValidators_AllVerifyError_WithMatch(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.AllVerifyError(0, "hello", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": false}
	expected.ShouldBeEqual(t, 0, "TextValidators AllVerifyError returns nil -- matching", actual)
}

// =============================================================================
// RangeSegmentsValidator — uncovered branches
// =============================================================================

func Test_RangeSegmentsValidator_Create(t *testing.T) {
	// Arrange
	rsv := &corevalidator.RangeSegmentsValidator{
		Title: "test-range",
	}

	// Act
	actual := args.Map{
		"title": rsv.Title,
		"segLen": rsv.LengthOfVerifierSegments(),
	}

	// Assert
	expected := args.Map{
		"title": "test-range",
		"segLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "RangeSegmentsValidator returns expected -- basic", actual)
}

// =============================================================================
// Condition — additional branches
// =============================================================================

func Test_Condition_AllTrue(t *testing.T) {
	// Arrange
	c := corevalidator.Condition{
		IsUniqueWordOnly:     true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
	}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	expected := args.Map{"isSplit": true}
	expected.ShouldBeEqual(t, 0, "Condition IsSplitByWhitespace returns true -- all true", actual)
}

// =============================================================================
// Parameter — additional branches
// =============================================================================

func Test_Parameter_SkipOnEmpty(t *testing.T) {
	// Arrange
	p := corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}

	// Act
	actual := args.Map{"isSkip": p.IsSkipCompareOnActualEmpty}

	// Assert
	expected := args.Map{"isSkip": true}
	expected.ShouldBeEqual(t, 0, "Parameter IsSkipCompareOnActualEmpty returns true -- set", actual)
}
