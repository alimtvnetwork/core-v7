package corevalidator

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type SliceValidators struct {
	Validators []SliceValidator
}

func (it *SliceValidators) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Validators)
}

func (it *SliceValidators) IsEmpty() bool {
	return it == nil || len(it.Validators) == 0
}

func (it *SliceValidators) SetActualOnAll(actualLines ...string) {
	if it.IsEmpty() {
		return
	}

	for _, sliceValidator := range it.Validators {
		sliceValidator.SetActual(actualLines)
	}
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
	params *Parameter,
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

func (it *SliceValidators) AssertVerifyAll(
	t *testing.T,
	params *Parameter,
) {
	if it.IsEmpty() {
		return
	}

	finalError := it.VerifyAllError(params)

	convey.Convey(params.Header, t, func() {
		convey.So(
			finalError,
			should.BeNil)
	})
}

func (it *SliceValidators) VerifyAllError(
	params *Parameter,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it.Validators {
		err := sliceValidator.AllVerifyError(params)

		errs.AddError(err)
	}

	header := params.Header

	errs.InsertAt(0, header)

	return errs.AsDefaultError()
}

func (it *SliceValidators) AssertVerifyAllUsingActual(
	t *testing.T,
	params *Parameter,
	actualLines ...string,
) {
	if it.IsEmpty() {
		return
	}

	finalError := it.VerifyAllErrorUsingActual(
		params,
		actualLines...)

	convey.Convey(params.Header, t, func() {
		convey.So(
			finalError,
			should.BeNil)
	})
}

func (it *SliceValidators) VerifyAllErrorUsingActual(
	params *Parameter,
	actualLines ...string,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it.Validators {
		sliceValidator.SetActual(actualLines)
		err := sliceValidator.AllVerifyError(params)

		errs.AddError(err)
	}

	header := params.Header

	errs.InsertAt(0, header)

	return errs.AsDefaultError()
}

// VerifyFirst
//
// Only collect using the SliceValidator.VerifyFirstError
func (it *SliceValidators) VerifyFirst(
	params *Parameter,
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

	header := params.Header

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
	params *Parameter,
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

	errs.InsertAt(0, params.Header)
	err := errs.AsDefaultError()

	if isPrintErr {
		fmt.Println(err)
	}

	return err
}
