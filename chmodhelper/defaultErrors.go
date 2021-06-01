package chmodhelper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

var (
	rwxInstructionNilErr = msgtype.
				CannotBeNilMessage.
				Error(
			"rwx (...) - parsing failed",
			"rwxInstruction - given as nil")

	failedToCompileVarWrapperToWrapper = msgtype.
						FailedToExecute.
						Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			constants.EmptyString)
)
