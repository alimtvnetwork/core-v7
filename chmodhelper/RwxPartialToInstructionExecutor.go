package chmodhelper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/errcore"
)

// RwxPartialToInstructionExecutor
// rwxPartial can be any length in
// between 0-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//  - "-rwx" will be "-rwx******"
//  - "-rwxr-x" will be "-rwxr-x***"
//  - "-rwxr-x" will be "-rwxr-x***"
func RwxPartialToInstructionExecutor(
	rwxPartial string,
	condition *chmodins.Condition,
) (*RwxInstructionExecutor, error) {
	if condition == nil {
		return nil, errcore.CannotBeNilOrEmptyMessage.
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
