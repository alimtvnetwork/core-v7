package chmodhelper

import (
	"errors"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/messages"
)

var (
	rwxInstructionNilErr = errcore.
				CannotBeNilType.
				Error(
			"rwx (...) - parsing failed",
			"rwxInstruction / rwxOwnerGroupOther - given as nil")

	failedToCompileVarWrapperToWrapper = errcore.
						FailedToExecuteType.
						Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			constants.EmptyString)

	errHyphenedRwxLength          = errors.New("length should be " + HyphenedRwxLengthString)
	errFullRwxLengthWithoutHyphen = errors.New("length should be " + FullRwxLengthWithoutHyphenString)
)
