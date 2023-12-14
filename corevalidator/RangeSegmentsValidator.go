package corevalidator

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corestr"
)

type RangeSegmentsValidator struct {
	actual           *corestr.SimpleSlice
	Title            string
	VerifierSegments []RangesSegment
}

func (it *RangeSegmentsValidator) LengthOfVerifierSegments() int {
	return len(it.VerifierSegments)
}

func (it *RangeSegmentsValidator) SetActual(
	lines []string,
) *RangeSegmentsValidator {
	it.actual = corestr.New.SimpleSlice.Direct(
		false,
		lines,
	)

	return it
}

func (it *RangeSegmentsValidator) Validators() HeaderSliceValidators {
	validators := make([]HeaderSliceValidator, it.LengthOfVerifierSegments())

	for _, segment := range it.VerifierSegments {
		expectedSegments := segment.ExpectedLines
		start := segment.RangeInt.Start
		end := segment.RangeInt.End
		actualSegments := it.actual.Strings()[start:end]
		totalItems := end - start + 1
		header := fmt.Sprintf(
			"%s - validate for range %d to %d (total: %d lines)",
			it.Title,
			start,
			end,
			totalItems,
		)
		validator := HeaderSliceValidator{
			Header: header,
			SliceValidator: SliceValidator{
				Condition:     segment.Condition,
				CompareAs:     segment.CompareAs,
				ActualLines:   actualSegments,
				ExpectedLines: expectedSegments,
			},
		}

		validators = append(validators, validator)
	}

	return validators
}

func (it *RangeSegmentsValidator) VerifyAll(
	header string,
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	it.SetActual(actual)

	return it.Validators().VerifyAll(
		header,
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifySimple(
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	return it.VerifyAll(
		it.Title,
		actual,
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifyFirst(
	header string,
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	params.Header = header
	it.SetActual(actual)

	return it.Validators().VerifyFirst(
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifyUpto(
	header string,
	actual []string,
	params *Parameter,
	length int,
	isPrintError bool,
) error {
	params.Header = header
	it.SetActual(actual)

	return it.Validators().VerifyUpto(
		isPrintError,
		false,
		length,
		params,
	)
}

func (it *RangeSegmentsValidator) VerifyFirstDefault(
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	return it.VerifyFirst(
		it.Title,
		actual,
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifyUptoDefault(
	actual []string,
	params *Parameter,
	length int,
	isPrintError bool,
) error {
	return it.VerifyUpto(
		it.Title,
		actual,
		params,
		length,
		isPrintError,
	)
}
