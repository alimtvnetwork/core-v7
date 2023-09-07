package corestr

import "gitlab.com/auk-go/core/constants"

type TextWithLineNumber struct {
	LineNumber int
	Text       string
}

func (it *TextWithLineNumber) HasLineNumber() bool {
	return it != nil && it.LineNumber > constants.InvalidValue
}

func (it *TextWithLineNumber) IsInvalidLineNumber() bool {
	return it == nil || it.LineNumber == constants.InvalidValue
}

func (it *TextWithLineNumber) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Text)
}

func (it *TextWithLineNumber) IsEmpty() bool {
	if it == nil {
		return true
	}

	return it.IsEmptyText() || it.IsInvalidLineNumber()
}

func (it *TextWithLineNumber) IsEmptyText() bool {
	if it == nil {
		return true
	}

	return len(it.Text) == 0
}

func (it *TextWithLineNumber) IsEmptyTextLineBoth() bool {
	return it.IsEmpty()
}
