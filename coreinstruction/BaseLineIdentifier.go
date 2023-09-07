package coreinstruction

import (
	"gitlab.com/auk-go/core/reqtype"
)

type BaseLineIdentifier struct {
	LineIdentifier
}

func NewBaseLineIdentifier(lineNumber int, modifyAs reqtype.Request) *BaseLineIdentifier {
	return &BaseLineIdentifier{
		LineIdentifier{
			LineNumber:   lineNumber,
			LineModifyAs: modifyAs,
		},
	}
}

func (it *BaseLineIdentifier) ToNewLineIdentifier() *LineIdentifier {
	if it == nil {
		return nil
	}

	return &LineIdentifier{
		LineNumber:   it.LineNumber,
		LineModifyAs: it.LineModifyAs,
	}
}

func (it *BaseLineIdentifier) Clone() *BaseLineIdentifier {
	if it == nil {
		return nil
	}

	return &BaseLineIdentifier{
		LineIdentifier{
			LineNumber:   it.LineNumber,
			LineModifyAs: it.LineModifyAs,
		},
	}
}
