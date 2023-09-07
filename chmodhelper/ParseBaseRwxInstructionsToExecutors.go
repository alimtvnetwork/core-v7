package chmodhelper

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

func ParseBaseRwxInstructionsToExecutors(
	baseRwxInstructions *chmodins.BaseRwxInstructions,
) (
	*RwxInstructionExecutors, error,
) {
	if baseRwxInstructions == nil || baseRwxInstructions.RwxInstructions == nil {
		return NewRwxInstructionExecutors(0), rwxInstructionNilErr
	}

	return ParseRwxInstructionsToExecutors(
		baseRwxInstructions.RwxInstructions)
}
