package corevalidator

import (
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

type SimpleSliceValidator struct {
	Expected *corestr.SimpleSlice
	actual   *corestr.SimpleSlice
	Condition
	CompareAs stringcompareas.Variant
}

func (it *SimpleSliceValidator) SetActual(lines []string) *SimpleSliceValidator {
	it.actual = corestr.New.SimpleSlice.Direct(
		false,
		lines,
	)

	return it
}

func (it *SimpleSliceValidator) SliceValidator() *SliceValidator {
	sliceValidator := SliceValidator{
		CompareAs:     it.CompareAs,
		Condition:     it.Condition,
		ActualLines:   it.actual.Strings(),
		ExpectedLines: it.Expected.Strings(),
	}

	return &sliceValidator
}

func (it *SimpleSliceValidator) VerifyAll(
	actual []string,
	params *Parameter,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.AllVerifyError(params)
}

func (it *SimpleSliceValidator) VerifyFirst(
	actual []string,
	params *Parameter,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.VerifyFirstError(params)
}

func (it *SimpleSliceValidator) VerifyUpto(
	actual []string,
	params *Parameter,
	length int,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.AllVerifyErrorUptoLength(
		false,
		params,
		length,
	)
}
