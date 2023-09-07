package corevalidator

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type SliceValidators struct {
	Validators []*SliceValidator
}

func (it SliceValidators) Length() int {
	return len(it.Validators)
}

func (it SliceValidators) IsEmpty() bool {
	return len(it.Validators) == 0
}

func (it *SliceValidators) IsValid(
	isCaseSensitive bool,
) bool {
	return it.IsMatch(isCaseSensitive)
}

func (it *SliceValidators) IsMatch(
	isCaseSensitive bool,
) bool {
	if it.IsEmpty() {
		return true
	}

	for _, sliceValidator := range it.Validators {
		if !sliceValidator.IsValid(isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it *SliceValidators) VerifyAll(
	header string,
	params *ValidatorParamsBase,
	isPrintError bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it.Validators {
		err := sliceValidator.AllVerifyError(params)

		errs.AddError(err)
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintError {
		fmt.Println(err)
	}

	return err
}

func (it *SliceValidators) VerifyFirst(
	header string,
	params *ValidatorParamsBase,
	isPrintError bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it.Validators {
		err := sliceValidator.VerifyFirstError(params)

		errs.AddError(err)
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintError {
		fmt.Println(err)
	}

	return err
}

func (it *SliceValidators) VerifyUpto(
	isPrintErr,
	isFirstOnly bool,
	length int,
	header string,
	params *ValidatorParamsBase,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it.Validators {
		err := sliceValidator.AllVerifyErrorUptoLength(
			isFirstOnly,
			params,
			length)

		errs.AddError(err)
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintErr {
		fmt.Println(err)
	}

	return err
}
