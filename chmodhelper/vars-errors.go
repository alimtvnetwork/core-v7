package chmodhelper

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/messages"
)

var (
	rwxInstructionNilErr = errcore.
				CannotBeNilMessage.
				Error(
			"rwx (...) - parsing failed",
			"rwxInstruction / rwxOwnerGroupOther - given as nil")

	failedToCompileVarWrapperToWrapper = errcore.
						FailedToExecute.
						Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			constants.EmptyString)

	errHyphenedRwxLength          = errors.New("length should be " + HyphenedRwxLengthString)
	errFullRwxLengthWithoutHyphen = errors.New("length should be " + FullRwxLengthWithoutHyphenString)
)
