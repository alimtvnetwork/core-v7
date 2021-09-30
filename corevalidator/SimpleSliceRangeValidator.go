package corevalidator

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type SimpleSliceRangeValidator struct {
	actual           *corestr.SimpleSlice
	VerifierSegments []RangesSegment
}

func (it *SimpleSliceRangeValidator) LengthOfVerifierSegments() int {
	return len(it.VerifierSegments)
}

func (it *SimpleSliceRangeValidator) SetActual(
	lines []string,
) *SimpleSliceRangeValidator {
	it.actual = corestr.NewSimpleSliceUsing(
		false,
		lines...)

	return it
}

func (it *SimpleSliceRangeValidator) SliceValidators() *SliceValidators {
	validators := make([]*SliceValidator, it.LengthOfVerifierSegments())

	for _, segment := range it.VerifierSegments {
		expectedSegments := segment.ExpectedLines
		actualSegments := it.actual.Items[segment.RangeInt.Start:segment.RangeInt.End]

		sliceValidator := SliceValidator{
			CompareAs:              segment.CompareAs,
			ValidatorCoreCondition: segment.ValidatorCoreCondition,
			ActualLines:            actualSegments,
			ExpectedLines:          expectedSegments,
		}

		validators = append(validators, &sliceValidator)
	}

	return &SliceValidators{
		Validators: validators,
	}
}

func (it *SimpleSliceRangeValidator) VerifyAll(
	header string,
	actual []string,
	params *ValidatorParamsBase,
	isPrintError bool,
) error {
	it.SetActual(actual)

	return it.SliceValidators().VerifyAll(
		header,
		params,
		isPrintError)
}

func (it *SimpleSliceRangeValidator) VerifyFirst(
	header string,
	actual []string,
	params *ValidatorParamsBase,
	isPrintError bool,
) error {
	it.SetActual(actual)

	return it.SliceValidators().VerifyFirst(
		header,
		params,
		isPrintError)
}

func (it *SimpleSliceRangeValidator) VerifyUpto(
	header string,
	actual []string,
	params *ValidatorParamsBase,
	length int,
	isPrintError bool,
) error {
	it.SetActual(actual)

	return it.SliceValidators().VerifyUpto(
		isPrintError,
		false,
		length,
		header,
		params,
	)
}
