package chmodhelper

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
)

func ParseRwxInstructionToExecutor(
	rwxInstruction *chmodins.RwxInstruction,
) (
	*RwxInstructionExecutor, error,
) {
	if rwxInstruction == nil {
		return nil, rwxInstructionNilErr
	}

	varWrapper, err := ParseRwxInstructionToVarWrapper(rwxInstruction)

	return &RwxInstructionExecutor{
		rwxInstruction: rwxInstruction,
		varWrapper:     varWrapper,
	}, err
}
