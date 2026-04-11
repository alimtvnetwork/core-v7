package corevalidatortests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — nil receiver paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_NilReceiver_IsUsedAlready(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"used": sv.IsUsedAlready()}

	// Assert
	expected := args.Map{"used": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns false -- nil IsUsedAlready", actual)
}

func Test_SliceValidator_NilReceiver_ActualLinesLength(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"len": sv.ActualLinesLength()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns 0 -- nil ActualLinesLength", actual)
}

func Test_SliceValidator_NilReceiver_ActualLinesString(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"str": sv.ActualLinesString()}

	// Assert
	expected := args.Map{"str": ""}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns empty -- nil ActualLinesString", actual)
}

func Test_SliceValidator_NilReceiver_ExpectingLinesString(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"str": sv.ExpectingLinesString()}

	// Assert
	expected := args.Map{"str": ""}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns empty -- nil ExpectingLinesString", actual)
}

func Test_SliceValidator_NilReceiver_ExpectingLinesLength(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"len": sv.ExpectingLinesLength()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns 0 -- nil ExpectingLinesLength", actual)
}

func Test_SliceValidator_NilReceiver_IsValid(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns true -- nil IsValid", actual)
}

func Test_SliceValidator_NilReceiver_Dispose(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	sv.Dispose() // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator no panic -- nil Dispose", actual)
}

func Test_SliceValidator_NilReceiver_VerifyFirstError(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.VerifyFirstError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil VerifyFirstError", actual)
}

func Test_SliceValidator_NilReceiver_AllVerifyError(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil AllVerifyError", actual)
}

func Test_SliceValidator_NilReceiver_AllVerifyErrorQuick(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorQuick(0, "header", "line1")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil AllVerifyErrorQuick", actual)
}

func Test_SliceValidator_NilReceiver_AllVerifyErrorTestCase(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorTestCase(0, "header", true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil AllVerifyErrorTestCase", actual)
}

func Test_SliceValidator_NilReceiver_AllVerifyErrorExceptLast(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorExceptLast(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil AllVerifyErrorExceptLast", actual)
}

func Test_SliceValidator_NilReceiver_AllVerifyErrorUptoLength(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorUptoLength(false, &corevalidator.Parameter{}, 5)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil AllVerifyErrorUptoLength", actual)
}

func Test_SliceValidator_NilReceiver_VerifyFirstLengthUptoError(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.VerifyFirstLengthUptoError(&corevalidator.Parameter{}, 5)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- nil VerifyFirstLengthUptoError", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — SetActual, SetActualVsExpected
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_SetActual_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"hello"},
	}
	sv.SetActual([]string{"hello"})

	// Act
	actual := args.Map{
		"used": sv.IsUsedAlready(),
		"len": sv.ActualLinesLength(),
	}

	// Assert
	expected := args.Map{
		"used": true,
		"len": 1,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns correct value -- SetActual", actual)
}

func Test_SliceValidator_SetActualVsExpected_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}
	sv.SetActualVsExpected([]string{"a"}, []string{"a"})

	// Act
	actual := args.Map{
		"used": sv.IsUsedAlready(),
		"aLen": sv.ActualLinesLength(),
		"eLen": sv.ExpectingLinesLength(),
	}

	// Assert
	expected := args.Map{
		"used": true,
		"aLen": 1,
		"eLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns correct value -- SetActualVsExpected", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — IsValid and IsValidOtherLines
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_IsValid_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"hello", "world"},
		ExpectedLines: []string{"hello", "world"},
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns true -- valid match", actual)
}

func Test_SliceValidator_IsValid_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"world"},
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns false -- mismatch", actual)
}

func Test_SliceValidator_IsValid_LengthDiff(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a"},
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns false -- length diff", actual)
}

func Test_SliceValidator_IsValid_BothNil_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   nil,
		ExpectedLines: nil,
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns true -- both nil", actual)
}

func Test_SliceValidator_IsValid_OneNil_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: nil,
	}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": false}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns false -- one nil", actual)
}

func Test_SliceValidator_IsValidOtherLines_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"ok"},
	}

	// Act
	actual := args.Map{"valid": sv.IsValidOtherLines(true, []string{"ok"})}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns true -- IsValidOtherLines match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — ComparingValidators (lazy creation & caching)
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_ComparingValidators_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"line1", "line2"},
	}
	v1 := sv.ComparingValidators()
	v2 := sv.ComparingValidators() // cached

	// Act
	actual := args.Map{
		"len": v1.Length(),
		"same": v1 == v2,
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"same": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns correct value -- ComparingValidators cached", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — Dispose with validators
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_Dispose_WithValidators_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	_ = sv.ComparingValidators() // create validators
	sv.Dispose()

	// Act
	actual := args.Map{
		"actualNil": sv.ActualLines == nil,
		"expectedNil": sv.ExpectedLines == nil,
	}

	// Assert
	expected := args.Map{
		"actualNil": true,
		"expectedNil": true,
	}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- Dispose clears all", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — AllVerifyErrorQuick success path
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_AllVerifyErrorQuick_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"hello"},
	}
	err := sv.AllVerifyErrorQuick(0, "test", "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- AllVerifyErrorQuick match", actual)
}

func Test_SliceValidator_AllVerifyErrorQuick_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ExpectedLines: []string{"hello"},
	}
	err := sv.AllVerifyErrorQuick(0, "test", "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- AllVerifyErrorQuick mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — AllVerifyErrorTestCase success/fail
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_AllVerifyErrorTestCase_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"hello"},
	}
	err := sv.AllVerifyErrorTestCase(0, "test match", true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- AllVerifyErrorTestCase match", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"hello"},
		ExpectedLines: []string{"world"},
	}
	err := sv.AllVerifyErrorTestCase(0, "test mismatch", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- AllVerifyErrorTestCase mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — AllVerifyErrorExceptLast
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_AllVerifyErrorExceptLast_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "DIFFERS"},
		ExpectedLines: []string{"a", "b", "c"},
	}
	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "except last",
		IsAttachUserInputs: false,
		IsCaseSensitive:    true,
	}
	err := sv.AllVerifyErrorExceptLast(params)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- AllVerifyErrorExceptLast skips last", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — VerifyFirstError (first error only)
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_VerifyFirstError_Match(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b"},
		ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := sv.VerifyFirstError(params)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- VerifyFirstError match", actual)
}

func Test_SliceValidator_VerifyFirstError_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "WRONG"},
		ExpectedLines: []string{"a", "b"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: true, Header: "first err"}
	err := sv.VerifyFirstError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- VerifyFirstError mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — isEmptyIgnoreCase
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_EmptyIgnoreCase(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{},
		ExpectedLines: []string{"expected"},
	}
	params := &corevalidator.Parameter{
		IsSkipCompareOnActualEmpty: true,
		IsCaseSensitive:            true,
	}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns nil -- empty ignore case skips", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — initialVerifyError nil paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_InitialVerify_AnyNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   nil,
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- actual nil expected not", actual)
}

func Test_SliceValidator_InitialVerify_ExpectedNil(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: nil,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- expected nil actual not", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — compactOrFullMismatchError single vs multi
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_CompactMismatch_SingleValue(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"actual"},
		ExpectedLines: nil,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- compact single mismatch", actual)
}

func Test_SliceValidator_CompactMismatch_MultiValue(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a", "b", "c"},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, IsAttachUserInputs: false}
	err := sv.AllVerifyError(params)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns error -- multi value mismatch", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — UserInputsMergeWithError paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_UserInputsMerge_NoAttach(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}
	params := &corevalidator.Parameter{IsAttachUserInputs: false}
	err := sv.UserInputsMergeWithError(params, errors.New("original"))

	// Act
	actual := args.Map{"msg": err.Error()}

	// Assert
	expected := args.Map{"msg": "original"}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns error -- no attach", actual)
}

func Test_SliceValidator_UserInputsMerge_WithAttach_NilErr(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{IsAttachUserInputs: true, CaseIndex: 1, Header: "test"}
	err := sv.UserInputsMergeWithError(params, nil)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns error -- nil err with inputs", actual)
}

func Test_SliceValidator_UserInputsMerge_WithAttach_WithErr(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{IsAttachUserInputs: true, CaseIndex: 1, Header: "test"}
	err := sv.UserInputsMergeWithError(params, errors.New("base"))

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns error -- with err and inputs", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidator — MethodName
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidator_MethodName_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidator{CompareAs: stringcompareas.Equal}

	// Act
	actual := args.Map{"name": sv.MethodName()}

	// Assert
	expected := args.Map{"name": stringcompareas.Equal.Name()}
	expected.ShouldBeEqual(t, 0, "SliceValidator returns correct value -- MethodName", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidatorConstructors — NewSliceValidatorUsingErr, NewSliceValidatorUsingAny
// ══════════════════════════════════════════════════════════════════════════════

func Test_NewSliceValidatorUsingErr_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingErr(
		errors.New("line1\nline2"),
		"line1\nline2",
		true, true, true,
		stringcompareas.Equal,
	)

	// Act
	actual := args.Map{
		"aLen": sv.ActualLinesLength(),
		"eLen": sv.ExpectingLinesLength(),
	}

	// Assert
	expected := args.Map{
		"aLen": 2,
		"eLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingErr returns correct value -- with args", actual)
}

func Test_NewSliceValidatorUsingAny_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"line1\nline2",
		"line1\nline2",
		true, true, true,
		stringcompareas.Equal,
	)

	// Act
	actual := args.Map{
		"aLen": sv.ActualLinesLength(),
		"eLen": sv.ExpectingLinesLength(),
	}

	// Assert
	expected := args.Map{
		"aLen": 2,
		"eLen": 2,
	}
	expected.ShouldBeEqual(t, 0, "NewSliceValidatorUsingAny returns correct value -- with args", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidator — various method paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidator_ToString_SingleLine_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "test",
		SearchAs: stringcompareas.Equal,
		Condition: corevalidator.Condition{
			IsTrimCompare:        true,
			IsNonEmptyWhitespace: true,
			IsSortStringsBySpace: true,
			IsUniqueWordOnly:     true,
		},
	}
	s := tv.ToString(true)
	ml := tv.ToString(false)

	// Act
	actual := args.Map{
		"single": s != "",
		"multi": ml != "",
	}

	// Assert
	expected := args.Map{
		"single": true,
		"multi": true,
	}
	expected.ShouldBeEqual(t, 0, "TextValidator returns correct value -- ToString formats", actual)
}

func Test_TextValidator_IsMatch_CaseInsensitive_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "Hello",
		SearchAs: stringcompareas.Equal,
	}

	// Act
	actual := args.Map{
		"sensitive":   tv.IsMatch("Hello", true),
		"insensitive": tv.IsMatch("hello", false),
		"mismatch":    tv.IsMatch("world", true),
	}

	// Assert
	expected := args.Map{
		"sensitive": true,
		"insensitive": true,
		"mismatch": false,
	}
	expected.ShouldBeEqual(t, 0, "TextValidator returns correct value -- IsMatch case sensitivity", actual)
}

func Test_TextValidator_IsMatchMany(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}

	// Act
	actual := args.Map{
		"allMatch":  tv.IsMatchMany(false, true, "hello", "hello"),
		"oneFails":  tv.IsMatchMany(false, true, "hello", "world"),
		"emptySkip": tv.IsMatchMany(true, true),
	}

	// Assert
	expected := args.Map{
		"allMatch": true,
		"oneFails": false,
		"emptySkip": true,
	}
	expected.ShouldBeEqual(t, 0, "TextValidator returns correct value -- IsMatchMany", actual)
}

func Test_TextValidator_IsMatchMany_Nil_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator

	// Act
	actual := args.Map{"nil": tv.IsMatchMany(true, true, "a")}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns true -- nil IsMatchMany", actual)
}

func Test_TextValidator_VerifyMany_ContinueOnError_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(true, params, "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- VerifyMany continue success", actual)
}

func Test_TextValidator_VerifyMany_FirstOnly(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyMany(false, params, "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- VerifyMany first only success", actual)
}

func Test_TextValidator_VerifyFirstError_SkipEmpty(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	err := tv.VerifyFirstError(params)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- VerifyFirstError skip empty", actual)
}

func Test_TextValidator_AllVerifyError_SkipEmpty(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsSkipCompareOnActualEmpty: true}
	err := tv.AllVerifyError(params)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- AllVerifyError skip empty", actual)
}

func Test_TextValidator_AllVerifyError_WithErrors_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "expected",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.AllVerifyError(params, "wrong1", "wrong2")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns error -- AllVerifyError with errors", actual)
}

func Test_TextValidator_VerifyDetailError_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifyDetailError(params, "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- VerifyDetailError match", actual)
}

func Test_TextValidator_VerifyDetailError_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, CaseIndex: 1, Header: "test"}
	err := tv.VerifyDetailError(params, "world")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns error -- VerifyDetailError mismatch", actual)
}

func Test_TextValidator_VerifySimpleError_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "ok",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tv.VerifySimpleError(0, params, "ok")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- VerifySimpleError match", actual)
}

func Test_TextValidator_VerifySimpleError_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "ok",
		SearchAs: stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{IsCaseSensitive: true, CaseIndex: 0}
	err := tv.VerifySimpleError(0, params, "not ok")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns error -- VerifySimpleError mismatch", actual)
}

func Test_TextValidator_MethodName_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{SearchAs: stringcompareas.Contains}
	var nilTv *corevalidator.TextValidator

	// Act
	actual := args.Map{
		"name": tv.MethodName(),
		"nil": nilTv.MethodName(),
	}

	// Assert
	expected := args.Map{
		"name": stringcompareas.Contains.Name(),
		"nil": "",
	}
	expected.ShouldBeEqual(t, 0, "TextValidator returns correct value -- MethodName", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidator — SearchTextFinalized with conditions (trim, whitespace split)
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidator_SearchTextFinalized_Trim(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "  hello  ",
		SearchAs: stringcompareas.Equal,
		Condition: corevalidator.Condition{
			IsTrimCompare: true,
		},
	}
	result := tv.SearchTextFinalized()

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": "hello"}
	expected.ShouldBeEqual(t, 0, "TextValidator returns correct value -- SearchTextFinalized trim", actual)
}

func Test_TextValidator_SearchTextFinalized_Cached_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tv := &corevalidator.TextValidator{
		Search:   "hello",
		SearchAs: stringcompareas.Equal,
	}
	r1 := tv.SearchTextFinalized()
	r2 := tv.SearchTextFinalized() // cached

	// Act
	actual := args.Map{"same": r1 == r2}

	// Assert
	expected := args.Map{"same": true}
	expected.ShouldBeEqual(t, 0, "TextValidator returns correct value -- SearchTextFinalized cached", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidators — various paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidators_NilLength_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	var tvs *corevalidator.TextValidators

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TextValidators returns 0 -- nil Length", actual)
}

func Test_TextValidators_AddSimple_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("test", stringcompareas.Equal)

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "TextValidators returns correct value -- AddSimple", actual)
}

func Test_TextValidators_AddSimpleAllTrue_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimpleAllTrue("test", stringcompareas.Equal)

	// Act
	actual := args.Map{
		"len": tvs.Length(),
		"hasAny": tvs.HasAnyItem(),
	}

	// Assert
	expected := args.Map{
		"len": 1,
		"hasAny": true,
	}
	expected.ShouldBeEqual(t, 0, "TextValidators returns correct value -- AddSimpleAllTrue", actual)
}

func Test_TextValidators_Adds_Empty_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.Adds()

	// Act
	actual := args.Map{"len": tvs.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "TextValidators returns empty -- Adds empty", actual)
}

func Test_TextValidators_IsMatch_Empty_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(0)

	// Act
	actual := args.Map{"match": tvs.IsMatch("anything", true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns true -- empty IsMatch", actual)
}

func Test_TextValidators_IsMatch_Success(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{
		"match": tvs.IsMatch("hello", true),
		"noMatch": tvs.IsMatch("world", true),
	}

	// Assert
	expected := args.Map{
		"match": true,
		"noMatch": false,
	}
	expected.ShouldBeEqual(t, 0, "TextValidators returns correct value -- IsMatch", actual)
}

func Test_TextValidators_IsMatchMany_Empty_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(0)

	// Act
	actual := args.Map{"match": tvs.IsMatchMany(true, true, "a")}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns true -- empty IsMatchMany", actual)
}

func Test_TextValidators_IsMatchMany_Fail(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)

	// Act
	actual := args.Map{"match": tvs.IsMatchMany(false, true, "world")}

	// Assert
	expected := args.Map{"match": false}
	expected.ShouldBeEqual(t, 0, "TextValidators returns false -- IsMatchMany fail", actual)
}

func Test_TextValidators_VerifyFirstError_Empty_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(0)
	err := tvs.VerifyFirstError(0, "content", true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- empty VerifyFirstError", actual)
}

func Test_TextValidators_VerifyFirstError_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.VerifyFirstError(0, "hello", true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- VerifyFirstError match", actual)
}

func Test_TextValidators_VerifyFirstError_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.VerifyFirstError(0, "world", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- VerifyFirstError mismatch", actual)
}

func Test_TextValidators_AllVerifyError_Match_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.AllVerifyError(0, "hello", true)

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- AllVerifyError match", actual)
}

func Test_TextValidators_AllVerifyError_Mismatch_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	err := tvs.AllVerifyError(0, "world", true)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns error -- AllVerifyError mismatch", actual)
}

func Test_TextValidators_VerifyErrorMany_Nil_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	var tvs *corevalidator.TextValidators
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyErrorMany(true, params, "a")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- nil VerifyErrorMany", actual)
}

func Test_TextValidators_VerifyErrorMany_Continue_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyErrorMany(true, params, "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- VerifyErrorMany continue match", actual)
}

func Test_TextValidators_VerifyErrorMany_FirstOnly_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("hello", stringcompareas.Equal)
	params := &corevalidator.Parameter{IsCaseSensitive: true}
	err := tvs.VerifyErrorMany(false, params, "hello")

	// Act
	actual := args.Map{"nil": err == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- VerifyErrorMany first only match", actual)
}

func Test_TextValidators_Dispose_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("test", stringcompareas.Equal)
	tvs.Dispose()

	// Act
	actual := args.Map{"nil": tvs.Items == nil}

	// Assert
	expected := args.Map{"nil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns nil -- Dispose", actual)
}

func Test_TextValidators_Dispose_Nil_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	var tvs *corevalidator.TextValidators
	tvs.Dispose() // no panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "TextValidators no panic -- nil Dispose", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Condition — IsSplitByWhitespace
// ══════════════════════════════════════════════════════════════════════════════

func Test_Condition_IsSplitByWhitespace_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	c1 := corevalidator.Condition{}
	c2 := corevalidator.Condition{IsUniqueWordOnly: true}
	c3 := corevalidator.Condition{IsNonEmptyWhitespace: true}
	c4 := corevalidator.Condition{IsSortStringsBySpace: true}

	// Act
	actual := args.Map{
		"none": c1.IsSplitByWhitespace(),
		"unique": c2.IsSplitByWhitespace(),
		"nonEmpty": c3.IsSplitByWhitespace(),
		"sort": c4.IsSplitByWhitespace(),
	}

	// Assert
	expected := args.Map{
		"none": false,
		"unique": true,
		"nonEmpty": true,
		"sort": true,
	}
	expected.ShouldBeEqual(t, 0, "Condition returns correct value -- IsSplitByWhitespace", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Parameter — IsIgnoreCase
// ══════════════════════════════════════════════════════════════════════════════

func Test_Parameter_IsIgnoreCase_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	p1 := corevalidator.Parameter{IsCaseSensitive: true}
	p2 := corevalidator.Parameter{IsCaseSensitive: false}

	// Act
	actual := args.Map{
		"p1": p1.IsIgnoreCase(),
		"p2": p2.IsIgnoreCase(),
	}

	// Assert
	expected := args.Map{
		"p1": false,
		"p2": true,
	}
	expected.ShouldBeEqual(t, 0, "Parameter returns correct value -- IsIgnoreCase", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidators — String, HasIndex, LastIndex, Count
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidators_QueryMethods(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	tvs.AddSimple("a", stringcompareas.Equal)
	tvs.AddSimple("b", stringcompareas.Contains)

	// Act
	actual := args.Map{
		"len":      tvs.Length(),
		"count":    tvs.Count(),
		"lastIdx":  tvs.LastIndex(),
		"hasIdx0":  tvs.HasIndex(0),
		"hasIdx5":  tvs.HasIndex(5),
		"hasAny":   tvs.HasAnyItem(),
		"empty":    tvs.IsEmpty(),
		"strNotE":  tvs.String() != "",
	}

	// Assert
	expected := args.Map{
		"len": 2, "count": 1, "lastIdx": 1,
		"hasIdx0": true, "hasIdx5": false,
		"hasAny": true, "empty": false, "strNotE": true,
	}
	expected.ShouldBeEqual(t, 0, "TextValidators returns correct value -- query methods", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// TextValidators — AsBasicSliceContractsBinder
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidators_AsBasicSliceContractsBinder_FromSliceValidatorNilRec(t *testing.T) {
	// Arrange
	tvs := corevalidator.NewTextValidators(5)
	binder := tvs.AsBasicSliceContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "TextValidators returns correct value -- AsBasicSliceContractsBinder", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// Vars — predefined conditions and validator
// ══════════════════════════════════════════════════════════════════════════════

func Test_Vars_DefaultConditions_FromSliceValidatorNilRec(t *testing.T) {
	// Act
	actual := args.Map{
		"disabled":  corevalidator.DefaultDisabledCoreCondition.IsTrimCompare,
		"trim":      corevalidator.DefaultTrimCoreCondition.IsTrimCompare,
		"sortTrim":  corevalidator.DefaultSortTrimCoreCondition.IsSortStringsBySpace,
		"unique":    corevalidator.DefaultUniqueWordsCoreCondition.IsUniqueWordOnly,
		"emptyName": corevalidator.EmptyValidator.SearchAs.Name(),
	}

	// Assert
	expected := args.Map{
		"disabled": false, "trim": true, "sortTrim": true,
		"unique": true, "emptyName": stringcompareas.Equal.Name(),
	}
	expected.ShouldBeEqual(t, 0, "Vars returns correct value -- default conditions", actual)
}
