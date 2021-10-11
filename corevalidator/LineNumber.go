package corevalidator

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

type LineNumber struct {
	LineNumber int `json:"LineNumber,omitempty"` // -1 means no checking in line
}

func (it *LineNumber) HasLineNumber() bool {
	return it.LineNumber > constants.InvalidValue
}

func (it *LineNumber) IsMatch(lineNumber int) bool {
	if lineNumber == constants.InvalidValue ||
		it.LineNumber == constants.InvalidValue {
		return true
	}

	return it.LineNumber == lineNumber
}

func (it *LineNumber) VerifyError(
	lineNumber int,
) error {
	if it.IsMatch(lineNumber) {
		return nil
	}

	msg := errcore.Expecting(
		"Line Number didn't match",
		it.LineNumber,
		lineNumber)

	return errors.New(msg)
}
