package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
)

// ── SliceValidators — uncovered branches ──

func Test_SliceValidators_Length_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidators

	// Act
	actual := args.Map{"len": sv.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil receiver", actual)
}

func Test_SliceValidators_IsEmpty_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidators

	// Act
	actual := args.Map{"isEmpty": sv.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil receiver", actual)
}

func Test_SliceValidators_IsEmpty_EmptyValidators(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{Validators: []corevalidator.SliceValidator{}}

	// Act
	actual := args.Map{"isEmpty": sv.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- empty validators slice", actual)
}

func Test_SliceValidators_IsMatch_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}

	// Act
	actual := args.Map{"match": sv.IsMatch(true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- empty validators", actual)
}

func Test_SliceValidators_IsValid_Empty(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}

	// Act
	actual := args.Map{"valid": sv.IsValid(true)}

	// Assert
	expected := args.Map{"valid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns true -- empty validators", actual)
}

func Test_SliceValidators_VerifyAll_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	err := sv.VerifyAll("header", &corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns nil -- empty validators", actual)
}

func Test_SliceValidators_VerifyAllError_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	err := sv.VerifyAllError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllError returns nil -- empty validators", actual)
}

func Test_SliceValidators_VerifyFirst_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	err := sv.VerifyFirst(&corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- empty validators", actual)
}

func Test_SliceValidators_VerifyUpto_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	err := sv.VerifyUpto(false, false, 10, &corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- empty validators", actual)
}

func Test_SliceValidators_SetActualOnAll_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	sv.SetActualOnAll("line1") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SetActualOnAll does not panic -- empty validators", actual)
}

func Test_SliceValidators_VerifyAllErrorUsingActual_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	sv := &corevalidator.SliceValidators{}
	err := sv.VerifyAllErrorUsingActual(&corevalidator.Parameter{}, "a", "b")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllErrorUsingActual returns nil -- empty validators", actual)
}

// ── HeaderSliceValidators — uncovered branches ──

func Test_HeaderSliceValidators_Length_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var hsv corevalidator.HeaderSliceValidators

	// Act
	actual := args.Map{"len": hsv.Length()}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "Length returns 0 -- nil slice", actual)
}

func Test_HeaderSliceValidators_IsEmpty_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var hsv corevalidator.HeaderSliceValidators

	// Act
	actual := args.Map{"isEmpty": hsv.IsEmpty()}

	// Assert
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "IsEmpty returns true -- nil slice", actual)
}

func Test_HeaderSliceValidators_IsMatch_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}

	// Act
	actual := args.Map{"match": hsv.IsMatch(true)}

	// Assert
	expected := args.Map{"match": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns true -- empty slice", actual)
}

func Test_HeaderSliceValidators_VerifyAll_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}
	err := hsv.VerifyAll("header", &corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAll returns nil -- empty slice", actual)
}

func Test_HeaderSliceValidators_VerifyAllError_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}
	err := hsv.VerifyAllError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllError returns nil -- empty slice", actual)
}

func Test_HeaderSliceValidators_VerifyFirst_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}
	err := hsv.VerifyFirst(&corevalidator.Parameter{}, false)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirst returns nil -- empty slice", actual)
}

func Test_HeaderSliceValidators_VerifyUpto_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}
	err := hsv.VerifyUpto(false, false, 10, &corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyUpto returns nil -- empty slice", actual)
}

func Test_HeaderSliceValidators_SetActualOnAll_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}
	hsv.SetActualOnAll("a") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "SetActualOnAll does not panic -- empty slice", actual)
}

func Test_HeaderSliceValidators_VerifyAllErrorUsingActual_Empty_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	hsv := corevalidator.HeaderSliceValidators{}
	err := hsv.VerifyAllErrorUsingActual(&corevalidator.Parameter{}, "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllErrorUsingActual returns nil -- empty slice", actual)
}

// ── SliceValidatorMessages — uncovered branches ──

func Test_SliceValidatorMessages_UserInputsMergeWithError_NoAttach(t *testing.T) {
	// Arrange
	sv := corevalidator.SliceValidator{
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"b"},
	}
	params := &corevalidator.Parameter{
		IsAttachUserInputs: false,
		CaseIndex:          0,
		Header:             "test",
	}
	err := sv.UserInputsMergeWithError(params, nil)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "UserInputsMergeWithError returns nil -- attach disabled nil error", actual)
}

// ── SliceValidatorAssertions — nil receiver ──

func Test_SliceValidator_AssertAllQuick_Nil(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	sv.AssertAllQuick(t, 0, "nil receiver test", "a") // should not panic

	// Act
	actual := args.Map{"ok": true}

	// Assert
	expected := args.Map{"ok": true}
	expected.ShouldBeEqual(t, 0, "AssertAllQuick does not panic -- nil receiver", actual)
}

// ── SliceValidatorVerify — nil receiver paths ──

func Test_SliceValidator_VerifyFirstError_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.VerifyFirstError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyFirstError returns nil -- nil receiver", actual)
}

func Test_SliceValidator_AllVerifyError_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyError(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyError returns nil -- nil receiver", actual)
}

func Test_SliceValidator_AllVerifyErrorExceptLast_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorExceptLast(&corevalidator.Parameter{})

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorExceptLast returns nil -- nil receiver", actual)
}

func Test_SliceValidator_AllVerifyErrorQuick_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorQuick(0, "header", "a")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorQuick returns nil -- nil receiver", actual)
}

func Test_SliceValidator_AllVerifyErrorTestCase_Nil_FromSliceValidatorsLengt(t *testing.T) {
	// Arrange
	var sv *corevalidator.SliceValidator
	err := sv.AllVerifyErrorTestCase(0, "header", true)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorTestCase returns nil -- nil receiver", actual)
}
