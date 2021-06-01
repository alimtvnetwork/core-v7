package chmodhelper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodinstruction"
)

func ParseRwxInstructionToExecutor(
	rwxInstruction *chmodinstruction.RwxInstruction,
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
