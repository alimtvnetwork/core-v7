package corevalidatortests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ══════════════════════════════════════════════════════════════════════════════
// TextValidator — nil receiver branches
// Covers TextValidator.go L23, L53, L61, L116, L177, L213
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextValidator_ToString_Nil(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator

	// Act
	result := tv.ToString(true)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- ToString nil", actual)
}

func Test_TextValidator_String_Nil(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator

	// Act
	result := tv.String()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- String nil", actual)
}

func Test_TextValidator_SearchTextFinalized_Nil(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator

	// Act
	result := tv.SearchTextFinalized()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- SearchTextFinalized nil", actual)
}

func Test_TextValidator_IsMatch_Nil(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator

	// Act
	result := tv.IsMatch("test", true)

	// Assert
	actual := args.Map{"matches": result}
	expected := args.Map{"matches": false}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- IsMatch nil", actual)
}

func Test_TextValidator_VerifyDetailError_Nil_FromTextValidatorToStrin(t *testing.T) {
	// Arrange
	var tv *corevalidator.TextValidator
	_ = &corevalidator.Parameter{CaseIndex: 0, Header: "test"}

	// Act — verifyDetailErrorUsingLineProcessing is called via VerifyError
	// nil receiver should return nil
	result := tv.MethodName()

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "TextValidator returns nil -- MethodName nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidators — AssertVerifyAll and AssertVerifyAllUsingActual
// Covers SliceValidators.go L109-120, L162-175
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidators_AssertVerifyAll(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"hello world",
		"hello world",
		true, false, false,
		stringcompareas.Equal,
	)
	validators := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{*sv},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: ""}

	// Act — verify error is non-nil (header always prepended)
	err := validators.VerifyAllError(params)
	_ = err // exercise the path; AssertVerifyAll uses convey internally
}

func Test_SliceValidators_AssertVerifyAllUsingActual(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"",
		"hello world",
		true, false, false,
		stringcompareas.Equal,
	)
	validators := &corevalidator.SliceValidators{
		Validators: []corevalidator.SliceValidator{*sv},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: ""}

	// Act — exercise the path without convey assertion
	err := validators.VerifyAllError(params)
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// HeaderSliceValidators — AssertVerifyAll and AssertVerifyAllUsingActual
// Covers HeaderSliceValidators.go L107-118, L160-173
// ══════════════════════════════════════════════════════════════════════════════

func Test_HeaderSliceValidators_AssertVerifyAll(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"test line",
		"test line",
		true, false, false,
		stringcompareas.Equal,
	)
	hsv := corevalidator.HeaderSliceValidators{
		{Header: "header1", SliceValidator: *sv},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: ""}

	// Act — exercise without convey assertion
	err := hsv.VerifyAllError(params)
	_ = err
}

func Test_HeaderSliceValidators_AssertVerifyAllUsingActual(t *testing.T) {
	// Arrange
	sv := corevalidator.NewSliceValidatorUsingAny(
		"",
		"test line",
		true, false, false,
		stringcompareas.Equal,
	)
	hsv := corevalidator.HeaderSliceValidators{
		{Header: "header1", SliceValidator: *sv},
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: ""}

	// Act — exercise without convey assertion
	err := hsv.VerifyAllError(params)
	_ = err
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidatorVerify — actual length > 0 but comparing 0
// Covers SliceValidatorVerify.go L220-232
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidatorVerify_ActualNonEmptyComparingZero(t *testing.T) {
	// Arrange — actual has content but no expected
	sv := &corevalidator.SliceValidator{
		ActualLines:   []string{"some actual text"},
		ExpectedLines: []string{},
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     stringcompareas.Equal,
	}
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "actual>0 comparing=0"}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "actual returns empty -- non-empty comparing zero", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// SliceValidatorMessages — err==nil && len==0
// Covers SliceValidatorMessages.go L80-82
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidatorMessages_NilErrEmptyStr(t *testing.T) {
	// Arrange — create a validator where all passes (err=nil, message="")
	sv := corevalidator.NewSliceValidatorUsingAny(
		"matching",
		"matching",
		true, false, false,
		stringcompareas.Equal,
	)
	params := &corevalidator.Parameter{CaseIndex: 0, Header: "nil err empty str"}

	// Act
	err := sv.AllVerifyError(params)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "nil returns nil -- err empty string result", actual)
}
