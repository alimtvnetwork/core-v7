package corevalidator

import (
	"errors"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/messages"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

type LinesValidators struct {
	Items []LineValidator
}

func NewLinesValidators(capacity int) *LinesValidators {
	linesValidators := LinesValidators{
		Items: make(
			[]LineValidator,
			0,
			capacity),
	}

	return &linesValidators
}

func (it *LinesValidators) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}

func (it *LinesValidators) Length() int {
	if it == nil {
		return constants.Zero
	}

	return len(it.Items)
}

func (it *LinesValidators) Count() int {
	return it.Length()
}

func (it *LinesValidators) IsEmpty() bool {
	return it.Length() == 0
}

func (it *LinesValidators) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *LinesValidators) LastIndex() int {
	return it.Length() - 1
}

func (it *LinesValidators) AddPtr(
	validator *LineValidator,
) *LinesValidators {
	if validator == nil {
		return it
	}

	it.Items = append(
		it.Items,
		*validator)

	return it
}

func (it *LinesValidators) Add(
	validator LineValidator,
) *LinesValidators {
	it.Items = append(
		it.Items,
		validator)

	return it
}

func (it *LinesValidators) Adds(
	validators ...LineValidator,
) *LinesValidators {
	for _, validator := range validators {
		it.Items = append(
			it.Items,
			validator)
	}

	return it
}

func (it *LinesValidators) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *LinesValidators) String() string {
	return strutilinternal.AnyToFieldNameString(it.Items)
}

func (it *LinesValidators) IsMatchText(
	text string,
	isCaseSensitive bool,
) bool {
	if it.IsEmpty() {
		return true
	}

	for _, validator := range it.Items {
		if !validator.IsMatch(
			constants.InvalidLineNumber,
			text,
			isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it *LinesValidators) IsMatch(
	isSkipValidationOnNoContents,
	isCaseSensitive bool,
	contentsWithLines ...corestr.TextWithLineNumber,
) bool {
	if it.IsEmpty() {
		return true
	}

	length := len(contentsWithLines)
	if length == 0 && isSkipValidationOnNoContents {
		return true
	} else if length == 0 && !isSkipValidationOnNoContents {
		return false
	}

	for _, validator := range it.Items {
		isNotMatches := !validator.IsMatchMany(
			isSkipValidationOnNoContents,
			isCaseSensitive,
			contentsWithLines...)

		if isNotMatches {
			return false
		}
	}

	return true
}

// VerifyFirstDefaultLineNumberError index considers to be the line number
func (it *LinesValidators) VerifyFirstDefaultLineNumberError(
	params *Parameter,
	contentsWithLines ...corestr.TextWithLineNumber,
) error {
	if it.IsEmpty() {
		return nil
	}

	length := len(contentsWithLines)
	const funcName = "VerifyFirstDefaultLineNumberError"
	if length == 0 && params.IsSkipCompareOnActualEmpty {
		return nil
	} else if length == 0 && !params.IsSkipCompareOnActualEmpty {
		return errcore.MeaningfulErrorWithData(
			errcore.ValidataionFailedType,
			funcName,
			errors.New(messages.CannotVerifyEmptyContentsWhereValidatorsArePresent),
			it.Items)
	}

	for _, validator := range it.Items {
		err := validator.VerifyFirstError(
			params,
			contentsWithLines...)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *LinesValidators) AllVerifyError(
	params *Parameter,
	contentsWithLines ...corestr.TextWithLineNumber,
) error {
	if it.IsEmpty() {
		return nil
	}

	length := len(contentsWithLines)
	const funcName = "AllVerifyError"
	if length == 0 && params.IsSkipCompareOnActualEmpty {
		return nil
	} else if length == 0 && !params.IsSkipCompareOnActualEmpty {
		return errcore.MeaningfulErrorWithData(
			errcore.ValidataionFailedType,
			funcName,
			errors.New(messages.CannotVerifyEmptyContentsWhereValidatorsArePresent),
			it.Items)
	}

	var sliceErr []string

	for _, validator := range it.Items {
		err := validator.AllVerifyError(
			params,
			contentsWithLines...)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error())
		}
	}

	return errcore.SliceToError(sliceErr)
}
