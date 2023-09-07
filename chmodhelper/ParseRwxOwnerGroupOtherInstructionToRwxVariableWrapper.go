package chmodhelper

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

func ParseRwxOwnerGroupOtherToRwxVariableWrapper(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (
	*RwxVariableWrapper, error,
) {
	if rwxOwnerGroupOther == nil {
		return nil, rwxInstructionNilErr
	}

	ownerVarAttr, ownerErr := ParseRwxToVarAttribute(rwxOwnerGroupOther.Owner)

	if ownerErr != nil {
		return nil, ownerErr
	}

	groupVarAttr, groupErr := ParseRwxToVarAttribute(rwxOwnerGroupOther.Group)

	if groupErr != nil {
		return nil, groupErr
	}

	otherVarAttr, otherErr := ParseRwxToVarAttribute(rwxOwnerGroupOther.Other)

	if otherErr != nil {
		return nil, otherErr
	}

	rawInput := ParseRwxInstructionToStringRwx(
		rwxOwnerGroupOther,
		false)

	isFixedType := ownerVarAttr.IsFixedType() &&
		groupVarAttr.IsFixedType() &&
		otherVarAttr.IsFixedType()

	return &RwxVariableWrapper{
		rawInput:    rawInput,
		isFixedType: isFixedType,
		Owner:       *ownerVarAttr,
		Group:       *groupVarAttr,
		Other:       *otherVarAttr,
	}, nil
}
