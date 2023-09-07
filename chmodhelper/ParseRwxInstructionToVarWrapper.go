package chmodhelper

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

func ParseRwxInstructionToVarWrapper(
	rwxInstruction *chmodins.RwxInstruction,
) (
	*RwxVariableWrapper, error,
) {
	if rwxInstruction == nil {
		return nil, rwxInstructionNilErr
	}

	return ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&rwxInstruction.RwxOwnerGroupOther)
}
