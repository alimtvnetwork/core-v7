package chmodhelper

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

var (
	rwxInstructionNilErr = msgtype.
				CannotBeNilMessage.
				Error(
			"rwx (...) - parsing failed",
			"rwxInstruction / rwxOwnerGroupOther - given as nil")

	failedToCompileVarWrapperToWrapper = msgtype.
						FailedToExecute.
						Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			constants.EmptyString)

	errHyphenedRwxLength          = errors.New("length should be " + HyphenedRwxLengthString)
	errFullRwxLengthWithoutHyphen = errors.New("length should be " + FullRwxLengthWithoutHyphenString)
)
