package chmodhelper

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/issetter"
)

func ParseRwxToVarAttribute(rwx string) (varAttribute *VarAttribute, err error) {
	length := len(rwx)

	if length != SingleRwxLength {
		return nil, GetRwxLengthError(rwx)
	}

	r, w, x := ExpandCharRwx(rwx)

	// any is true
	isRead := issetter.GetBool(r == ReadChar)
	isWrite := issetter.GetBool(w == WriteChar)
	isExecute := issetter.GetBool(x == ExecuteChar)

	// is any has '*' wildcard
	isReadWildcard := r == constants.WildcardChar
	isWriteWildcard := w == constants.WildcardChar
	isExecuteWildcard := x == constants.WildcardChar

	isVarType := isReadWildcard ||
		isWriteWildcard ||
		isExecuteWildcard

	if isVarType {
		readVal := issetter.GetSet(
			isReadWildcard,
			issetter.Wildcard,
			isRead)

		writeVal := issetter.GetSet(
			isWriteWildcard,
			issetter.Wildcard,
			isWrite)

		execVal := issetter.GetSet(
			isExecuteWildcard,
			issetter.Wildcard,
			isExecute)

		return &VarAttribute{
			rawInput:    rwx,
			isFixedType: !isVarType,
			isRead:      readVal,
			isWrite:     writeVal,
			isExecute:   execVal,
		}, nil
	}

	return &VarAttribute{
		rawInput:    rwx,
		isFixedType: !isVarType,
		isRead:      isRead,
		isWrite:     isWrite,
		isExecute:   isExecute,
	}, nil
}
