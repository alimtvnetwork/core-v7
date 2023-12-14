package coretestcases

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

// CaseV1
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
//   - Will verify type using VerifyTypeOf
type CaseV1 coretests.BaseTestCase

func (it CaseV1) Input() interface{} {
	return it.ArrangeInput
}

func (it CaseV1) Expected() interface{} {
	return it.ExpectedInput
}

func (it CaseV1) ArrangeTypeName() string {
	return reflectinternal.TypeName(it.ArrangeInput)
}

// Actual
//
// Must SetActual first.
func (it CaseV1) Actual() interface{} {
	return it.ActualInput
}

func (it CaseV1) AsSimpleTestCaseWrapper() coretests.SimpleTestCaseWrapper {
	return it
}

func (it CaseV1) SetActual(actual interface{}) {
	it.ActualInput = actual
}

func (it CaseV1) CaseTitle() string {
	return it.Title
}

func (it CaseV1) SetExpected(expected interface{}) {
	it.ExpectedInput = expected
}

func (it CaseV1) VerifyAllEqual(
	caseIndex int,
	actualElements ...string,
) error {
	return it.VerifyAll(
		caseIndex,
		stringcompareas.Equal,
		actualElements,
	)
}

func (it CaseV1) VerifyAllEqualCondition(
	caseIndex int,
	condition corevalidator.Condition,
	actualElements ...string,
) error {
	return it.VerifyAllCondition(
		caseIndex,
		stringcompareas.Equal,
		condition,
		actualElements,
	)
}

func (it CaseV1) SliceValidator(
	compareAs stringcompareas.Variant,
	actualElements []string,
) corevalidator.SliceValidator {
	return it.SliceValidatorCondition(
		compareAs,
		corevalidator.DefaultDisabledCoreCondition,
		actualElements,
	)
}

func (it CaseV1) SliceValidatorCondition(
	compareAs stringcompareas.Variant,
	condition corevalidator.Condition,
	actualElements []string,
) corevalidator.SliceValidator {
	it.SetActual(actualElements)

	sliceValidator := corevalidator.SliceValidator{
		Condition:     condition,
		CompareAs:     compareAs,
		ActualLines:   actualElements,
		ExpectedLines: it.ExpectedInput.([]string),
	}

	return sliceValidator
}

func (it CaseV1) VerifyAll(
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElements []string,
) error {
	it.SetActual(actualElements)

	sliceValidator := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     compareAs,
		ActualLines:   actualElements,
		ExpectedLines: it.ExpectedInput.([]string),
	}

	finalErr := it.VerifyAllSliceValidator(
		caseIndex,
		sliceValidator,
	)

	sliceValidator.Dispose()

	return finalErr
}

func (it CaseV1) VerifyAllCondition(
	caseIndex int,
	compareAs stringcompareas.Variant,
	condition corevalidator.Condition,
	actualElements []string,
) error {
	it.SetActual(actualElements)

	sliceValidator := corevalidator.SliceValidator{
		Condition:     condition,
		CompareAs:     compareAs,
		ActualLines:   actualElements,
		ExpectedLines: it.ExpectedInput.([]string),
	}

	finalErr := it.VerifyAllSliceValidator(
		caseIndex,
		sliceValidator,
	)

	sliceValidator.Dispose()

	return finalErr
}

func (it CaseV1) VerifyFirst(
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElements []string,
) error {
	it.SetActual(actualElements)

	sliceValidator := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     compareAs,
		ActualLines:   actualElements,
		ExpectedLines: it.ExpectedInput.([]string),
	}

	param := corevalidator.Parameter{
		CaseIndex:          caseIndex,
		Header:             it.Title,
		IsAttachUserInputs: true,
		IsCaseSensitive:    true,
	}

	return sliceValidator.VerifyFirstError(&param)
}

func (it CaseV1) VerifyAllSliceValidator(
	caseIndex int,
	validator corevalidator.SliceValidator,
) error {
	param := corevalidator.Parameter{
		CaseIndex:          caseIndex,
		Header:             it.Title,
		IsAttachUserInputs: true,
		IsCaseSensitive:    true,
	}

	return validator.AllVerifyError(&param)
}

func (it CaseV1) VerifyError(
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElements ...string,
) error {
	toBaseTestCase := it.AsBaseTestCase()
	validationFinalError := it.VerifyAll(
		caseIndex,
		compareAs,
		actualElements,
	)

	if toBaseTestCase.IsTypeInvalidOrSkipVerify() {
		return validationFinalError
	}

	typeVerifyErr := toBaseTestCase.TypeValidationError()

	return errcore.MergeErrors(
		validationFinalError,
		typeVerifyErr,
	)
}

func (it CaseV1) ShouldBe(
	t *testing.T,
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElements ...string,
) error {
	return it.ShouldBeUsingCondition(
		t,
		caseIndex,
		compareAs,
		corevalidator.DefaultDisabledCoreCondition,
		actualElements...,
	)
}

func (it CaseV1) ShouldBeUsingCondition(
	t *testing.T,
	caseIndex int,
	compareAs stringcompareas.Variant,
	condition corevalidator.Condition,
	actualElements ...string,
) error {
	toBaseTestCase := it.AsBaseTestCase()
	validationFinalError := it.VerifyAllCondition(
		caseIndex,
		compareAs,
		condition,
		actualElements,
	)

	convey.Convey(
		toBaseTestCase.Title, t, func() {
			convey.So(
				validationFinalError,
				should.BeNil,
			)
		},
	)

	if toBaseTestCase.IsTypeInvalidOrSkipVerify() {
		return validationFinalError
	}

	typeVerifyErr := it.TypeShouldMatch(t)

	return errcore.MergeErrors(
		validationFinalError,
		typeVerifyErr,
	)
}

// TypeShouldMatch
//
// Assert along with returns the error.
func (it CaseV1) TypeShouldMatch(
	t *testing.T,
) error {
	baseCase := it.AsBaseTestCase()
	typeVerifyErr := baseCase.TypeValidationError()
	typeVerifyTitle := fmt.Sprintf(
		typeVerifyTitleFormat,
		it.Title,
	)

	convey.Convey(
		typeVerifyTitle, t, func() {
			convey.So(
				typeVerifyErr,
				should.BeNil,
			)
		},
	)

	return typeVerifyErr
}

func (it CaseV1) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.Equal,
		actualElements...,
	)
}

func (it CaseV1) ShouldBeTrimEqual(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBeUsingCondition(
		t,
		caseIndex,
		stringcompareas.Equal,
		corevalidator.DefaultTrimCoreCondition,
		actualElements...,
	)
}

func (it CaseV1) ShouldBeSortedEqual(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBeUsingCondition(
		t,
		caseIndex,
		stringcompareas.Equal,
		corevalidator.DefaultSortTrimCoreCondition,
		actualElements...,
	)
}

func (it CaseV1) ShouldContains(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.Contains,
		actualElements...,
	)
}

func (it CaseV1) ShouldStartsWith(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.StartsWith,
		actualElements...,
	)
}

func (it CaseV1) ShouldEndsWith(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.EndsWith,
		actualElements...,
	)
}

func (it CaseV1) ShouldBeNotEqual(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.NotEqual,
		actualElements...,
	)
}

// ShouldBeRegex
//
// Each expectation line acts as a regex to
// be validated against the actual line.
func (it CaseV1) ShouldBeRegex(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBe(
		t,
		caseIndex,
		stringcompareas.Regex,
		actualElements...,
	)
}

// ShouldBeRegex
//
// Each expectation line acts as a regex to
// be validated against the actual line.
func (it CaseV1) ShouldBeTrimRegex(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.ShouldBeUsingCondition(
		t,
		caseIndex,
		stringcompareas.Regex,
		corevalidator.DefaultTrimCoreCondition,
		actualElements...,
	)
}

func (it CaseV1) ShouldHaveNoError(
	t *testing.T,
	additionalTitle string,
	caseIndex int,
	err error,
) {
	finalTitle := fmt.Sprintf(
		"%d - %s - %s",
		caseIndex,
		it.CaseTitle(),
		additionalTitle,
	)

	convey.Convey(
		finalTitle, t, func() {
			convey.So(
				err,
				should.BeNil,
			)
		},
	)
}

// AssertDirectly
//
// Assert directly using convey.Convey
func (it CaseV1) AssertDirectly(
	t *testing.T,
	additionalTitle string,
	msg string,
	caseIndex int,
	actual interface{},
	assertion convey.Assertion,
	expectation interface{},
) {
	finalTitle := it.PrepareTitle(
		caseIndex,
		additionalTitle,
	)

	convey.Convey(
		finalTitle, t, func() {
			convey.SoMsg(
				msg,
				actual,
				assertion,
				expectation,
			)
		},
	)
}

func (it CaseV1) PrepareTitle(
	caseIndex int,
	additionalTitle string,
) string {
	return fmt.Sprintf(
		"%d - %s - %s",
		caseIndex,
		it.CaseTitle(),
		additionalTitle,
	)
}

func (it CaseV1) AsBaseTestCase() coretests.BaseTestCase {
	return coretests.BaseTestCase(it)
}

func (it CaseV1) AsSimpleTestCaseWrapperContractsBinder() coretests.SimpleTestCaseWrapperContractsBinder {
	return &it
}
