package chmodhelper

import "gitlab.com/auk-go/core/chmodhelper/chmodins"

func ParseRwxInstructionsToExecutors(
	rwxInstructions []chmodins.RwxInstruction,
) (
	*RwxInstructionExecutors, error,
) {
	if rwxInstructions == nil {
		return NewRwxInstructionExecutors(0), rwxInstructionNilErr
	}

	length := len(rwxInstructions)
	executors := NewRwxInstructionExecutors(length)

	if length == 0 {
		return executors, nil
	}

	for _, instruction := range rwxInstructions {
		executor, err := ParseRwxInstructionToExecutor(&instruction)

		if err != nil {
			return executors, err
		}

		executors.Add(executor)
	}

	return executors, nil
}
