package coreinstruction

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/reqtype"
)

type LineIdentifier struct {
	LineNumber   int             `json:"LineNumber,omitempty"`
	LineModifyAs reqtype.Request `json:"LineModifyAs,omitempty"`
}

func (it *LineIdentifier) IsInvalidLineNumber() bool {
	return it == nil || it.LineNumber < 0
}

func (it *LineIdentifier) IsInvalidLineNumberUsingLastLineNumber(lastLineNumber int) bool {
	return it == nil || it.LineNumber < 0 || it.LineNumber > lastLineNumber
}

func (it *LineIdentifier) HasLineNumber() bool {
	return it != nil && it.LineNumber > constants.InvalidValue
}

func (it *LineIdentifier) IsNewLineRequest() bool {
	return it.LineModifyAs.IsCreate()
}

func (it *LineIdentifier) IsDeleteLineRequest() bool {
	return it.HasLineNumber() &&
		it.LineModifyAs.IsDelete() ||
		it.LineModifyAs.IsDrop()
}

func (it *LineIdentifier) IsModifyLineRequest() bool {
	return it.HasLineNumber() &&
		it.LineModifyAs.IsUpdate()
}

func (it *LineIdentifier) IsAddNewOrModifyLineRequest() bool {
	return it.IsNewLineRequest() || it.IsModifyLineRequest()
}

func (it *LineIdentifier) ToBaseLineIdentifier() *BaseLineIdentifier {
	if it == nil {
		return nil
	}

	return NewBaseLineIdentifier(it.LineNumber, it.LineModifyAs)
}

func (it *LineIdentifier) Clone() *LineIdentifier {
	if it == nil {
		return nil
	}

	return &LineIdentifier{
		LineNumber:   it.LineNumber,
		LineModifyAs: it.LineModifyAs,
	}
}
