package chmodhelper

import "gitlab.com/evatix-go/core/chmodhelper/chmodinstruction"

func ParseRwxInstructionToVarWrapper(
	rwxInstruction *chmodinstruction.RwxInstruction,
) (
	*VarWrapper, error,
) {
	if rwxInstruction == nil {
		return nil, rwxInstructionNilErr
	}

	ownerVarAttr, ownerErr := ParseRwxToVarAttribute(rwxInstruction.Owner)

	if ownerErr != nil {
		return nil, ownerErr
	}

	groupVarAttr, groupErr := ParseRwxToVarAttribute(rwxInstruction.Group)

	if groupErr != nil {
		return nil, groupErr
	}

	otherVarAttr, otherErr := ParseRwxToVarAttribute(rwxInstruction.Other)

	if otherErr != nil {
		return nil, otherErr
	}

	rawInput := ParseRwxInstructionToStringRwx(
		rwxInstruction,
		false)

	isFixedType := ownerVarAttr.IsFixedType() &&
		groupVarAttr.IsFixedType() &&
		otherVarAttr.IsFixedType()

	return &VarWrapper{
		rawInput:    rawInput,
		isFixedType: isFixedType,
		Owner:       *ownerVarAttr,
		Group:       *groupVarAttr,
		Other:       *otherVarAttr,
	}, nil
}
