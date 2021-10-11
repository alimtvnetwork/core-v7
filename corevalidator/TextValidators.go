package corevalidator

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type TextValidators struct {
	Items []TextValidator
}

func NewTextValidators(capacity int) *TextValidators {
	slice := make([]TextValidator, 0, capacity)

	return &TextValidators{
		Items: slice,
	}
}

func (it *TextValidators) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}

func (it *TextValidators) Length() int {
	if it == nil {
		return constants.Zero
	}

	return len(it.Items)
}

func (it *TextValidators) Count() int {
	return it.LastIndex()
}

func (it *TextValidators) IsEmpty() bool {
	return it.Length() == 0
}

func (it *TextValidators) Add(
	validator TextValidator,
) *TextValidators {
	it.Items = append(
		it.Items,
		validator)

	return it
}

func (it *TextValidators) Adds(
	validators ...TextValidator,
) *TextValidators {
	if len(validators) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		validators...)

	return it
}

func (it *TextValidators) AddSimple(
	searchTerm string,
	compareAs stringcompareas.Variant,
) *TextValidators {
	return it.Add(TextValidator{
		Search:   searchTerm,
		SearchAs: compareAs,
	})
}

func (it *TextValidators) AddSimpleAllTrue(
	searchTerm string,
	compareAs stringcompareas.Variant,
) *TextValidators {
	coreCondition := ValidatorCoreCondition{
		IsTrimCompare:        true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
		IsUniqueWordOnly:     true,
	}

	return it.Add(
		TextValidator{
			Search:                 searchTerm,
			SearchAs:               compareAs,
			ValidatorCoreCondition: coreCondition,
		})
}

func (it *TextValidators) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *TextValidators) LastIndex() int {
	return it.Length() - 1
}

func (it *TextValidators) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *TextValidators) String() string {
	return utilstringinternal.AnyToFieldNameString(
		it.Items)
}

func (it *TextValidators) IsMatch(
	content string,
	isCaseSensitive bool,
) bool {
	if it.IsEmpty() {
		return true
	}

	for _, validator := range it.Items {
		if !validator.IsMatch(
			content,
			isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it *TextValidators) IsMatchMany(
	isSkipOnContentsEmpty,
	isCaseSensitive bool,
	contents ...string,
) bool {
	if it.IsEmpty() {
		return true
	}

	for _, validator := range it.Items {
		isNotMatched := !validator.IsMatchMany(
			isSkipOnContentsEmpty,
			isCaseSensitive,
			contents...)

		if isNotMatched {
			return isNotMatched
		}
	}

	return true
}

func (it *TextValidators) VerifyFirstError(
	caseIndex int,
	content string,
	isCaseSensitive bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	params := ValidatorParamsBase{
		CaseIndex:                         caseIndex,
		IsIgnoreCompareOnActualInputEmpty: false,
		IsAttachUserInputs:                false,
		IsCaseSensitive:                   isCaseSensitive,
	}

	for _, validator := range it.Items {
		err := validator.VerifyDetailError(
			&params,
			content,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *TextValidators) VerifyErrorMany(
	isContinueOnError bool,
	params *ValidatorParamsBase,
	contents ...string,
) error {
	if it == nil {
		return nil
	}

	if isContinueOnError {
		return it.AllVerifyErrorMany(
			params,
			contents...)
	}

	return it.VerifyFirstErrorMany(
		params,
		contents...)
}

func (it *TextValidators) VerifyFirstErrorMany(
	params *ValidatorParamsBase,
	contents ...string,
) error {
	if it.IsEmpty() {
		return nil
	}

	for _, item := range it.Items {
		err := item.AllVerifyError(
			params,
			contents...,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *TextValidators) AllVerifyErrorMany(
	params *ValidatorParamsBase,
	contents ...string,
) error {
	if it.IsEmpty() {
		return nil
	}

	capacity := defaultcapacity.OfSearch(it.Length())
	errorSlice := make(
		[]string,
		0,
		capacity)

	for _, item := range it.Items {
		err := item.AllVerifyError(
			params,
			contents...,
		)

		if err != nil {
			errorSlice = append(
				errorSlice,
				err.Error())
		}
	}

	return errcore.SliceToError(
		errorSlice)
}

func (it *TextValidators) AllVerifyError(
	caseIndex int,
	content string,
	isCaseSensitive bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	capacity := defaultcapacity.OfSearch(it.Length())
	errorSlice := make(
		[]string,
		0,
		capacity)

	params := ValidatorParamsBase{
		CaseIndex:                         caseIndex,
		IsIgnoreCompareOnActualInputEmpty: false,
		IsAttachUserInputs:                false,
		IsCaseSensitive:                   isCaseSensitive,
	}

	for _, item := range it.Items {
		err := item.VerifyDetailError(
			&params,
			content)

		if err != nil {
			errorSlice = append(errorSlice, err.Error())
		}
	}

	return errcore.SliceToError(errorSlice)
}
