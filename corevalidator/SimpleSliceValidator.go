package corevalidator

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
)

type SimpleSliceValidator struct {
	Expected *corestr.SimpleSlice
	actual   *corestr.SimpleSlice
	ValidatorCoreCondition
	CompareAs stringcompareas.Variant
}

func (it *SimpleSliceValidator) SetActual(lines []string) *SimpleSliceValidator {
	it.actual = corestr.NewSimpleSliceUsing(
		false,
		lines...)

	return it
}

func (it *SimpleSliceValidator) SliceValidator() *SliceValidator {
	sliceValidator := SliceValidator{
		CompareAs:              it.CompareAs,
		ValidatorCoreCondition: it.ValidatorCoreCondition,
		ActualLines:            it.actual.Items,
		ExpectedLines:          it.Expected.Items,
	}

	return &sliceValidator
}

func (it *SimpleSliceValidator) VerifyAll(
	actual []string,
	params *ValidatorParamsBase,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.AllVerifyError(params)
}

func (it *SimpleSliceValidator) VerifyFirst(
	actual []string,
	params *ValidatorParamsBase,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.VerifyFirstError(params)
}

func (it *SimpleSliceValidator) VerifyUpto(
	actual []string,
	params *ValidatorParamsBase,
	length int,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.AllVerifyErrorUptoLength(
		false,
		params,
		length)
}
