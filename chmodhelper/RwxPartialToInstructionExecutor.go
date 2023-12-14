package chmodhelper

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/errcore"
)

// RwxPartialToInstructionExecutor
//
// rwxPartial can be any length in
// between 0-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func RwxPartialToInstructionExecutor(
	rwxPartial string,
	condition *chmodins.Condition,
) (*RwxInstructionExecutor, error) {
	if condition == nil {
		return nil, errcore.CannotBeNilOrEmptyType.
			ErrorNoRefs("condition")
	}

	ownerGroupOther, err := chmodins.ExpandRwxFullStringToOwnerGroupOtherByFixingFirst(
		rwxPartial)

	if err != nil {
		return nil, err
	}

	rwxInstruction := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *ownerGroupOther,
		Condition:          *condition,
	}

	return ParseRwxInstructionToExecutor(rwxInstruction)
}
