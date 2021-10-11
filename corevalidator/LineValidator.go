package corevalidator

import (
	"errors"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
)

type LineValidator struct {
	LineNumber
	TextValidator
}

// IsMatch
//
// lineNumber == -1 mean no checking in line number,
//
// having LineValidator.LineNumber = -1 is also means the same
func (it *LineValidator) IsMatch(
	lineNumber int,
	content string,
	isCaseSensitive bool,
) bool {
	if !it.LineNumber.IsMatch(lineNumber) {
		return false
	}

	return it.TextValidator.IsMatch(
		content,
		isCaseSensitive)
}

func (it *LineValidator) IsMatchMany(
	isSkipOnContentsEmpty,
	isCaseSensitive bool,
	contentsWithLine ...corestr.TextWithLineNumber,
) bool {
	if it == nil {
		return true
	}

	if len(contentsWithLine) == 0 && isSkipOnContentsEmpty {
		return true
	}

	for _, textWithLine := range contentsWithLine {
		if !it.IsMatch(
			textWithLine.LineNumber,
			textWithLine.Text,
			isCaseSensitive) {
			return false
		}
	}

	return true
}

// VerifyError
//
// lineNumber == -1 mean no checking in line number,
//
// having LineValidator.LineNumber = -1 is also means the same
func (it *LineValidator) VerifyError(
	params *ValidatorParamsBase,
	processingLineNumber int,
	content string,
) error {
	if !it.LineNumber.IsMatch(processingLineNumber) {
		msg := errcore.GetSearchLineNumberExpectationMessage(
			params.CaseIndex,
			it.LineNumber.LineNumber,
			processingLineNumber,
			content,
			it.Search,
			*it)

		return errors.New(msg)
	}

	return it.TextValidator.verifyDetailErrorUsingLineProcessing(
		processingLineNumber,
		params,
		content)
}

func (it *LineValidator) VerifyMany(
	isContinueOnError bool,
	params *ValidatorParamsBase,
	contentsWithLine ...corestr.TextWithLineNumber,
) error {
	if isContinueOnError {
		return it.AllVerifyError(
			params,
			contentsWithLine...)
	}

	return it.VerifyFirstError(
		params,
		contentsWithLine...)
}

func (it *LineValidator) VerifyFirstError(
	params *ValidatorParamsBase,
	contentsWithLine ...corestr.TextWithLineNumber,
) error {
	if it == nil {
		return nil
	}

	length := len(contentsWithLine)
	if length == 0 && params.IsIgnoreCompareOnActualInputEmpty {
		return nil
	}

	for _, textWithLine := range contentsWithLine {
		err := it.VerifyError(
			params,
			textWithLine.LineNumber,
			textWithLine.Text,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *LineValidator) AllVerifyError(
	params *ValidatorParamsBase,
	contentsWithLine ...corestr.TextWithLineNumber,
) error {
	if it == nil {
		return nil
	}

	length := len(contentsWithLine)
	if length == 0 && params.IsIgnoreCompareOnActualInputEmpty {
		return nil
	}

	var sliceErr []string

	for _, textWithLine := range contentsWithLine {
		err := it.VerifyError(
			params,
			textWithLine.LineNumber,
			textWithLine.Text,
		)

		if err != nil {
			return err
		}
	}

	return errcore.SliceToError(
		sliceErr)
}
