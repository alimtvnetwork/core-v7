package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

type RwxInstructionExecutor struct {
	rwxInstruction *chmodins.RwxInstruction
	varWrapper     *RwxVariableWrapper
}

// IsVarWrapper if it has any wildcard symbol in it
func (receiver *RwxInstructionExecutor) IsVarWrapper() bool {
	return !receiver.varWrapper.IsFixedType()
}

// IsFixedWrapper true indicates no wildcard symbol
func (receiver *RwxInstructionExecutor) IsFixedWrapper() bool {
	return receiver.varWrapper.IsFixedType()
}

func (receiver *RwxInstructionExecutor) CompiledWrapper(mode os.FileMode) (*RwxWrapper, error) {
	if receiver.IsFixedWrapper() {
		return receiver.
			varWrapper.
			ToCompileWrapperPtr(nil), nil
	}

	if receiver.IsVarWrapper() {
		fixedWrapper := NewUsingFileMode(mode)

		return receiver.
			varWrapper.
			ToCompileWrapperPtr(&fixedWrapper), nil
	}

	return nil, failedToCompileVarWrapperToWrapper
}

func (receiver *RwxInstructionExecutor) CompiledRwxWrapperUsingFixedRwxWrapper(
	wrapper *RwxWrapper,
) (*RwxWrapper, error) {
	if receiver.IsFixedWrapper() {
		return receiver.
			varWrapper.
			ToCompileWrapperPtr(nil), nil
	}

	if receiver.IsVarWrapper() {
		return receiver.
			varWrapper.
			ToCompileWrapperPtr(wrapper), nil
	}

	return nil, msgtype.
		FailedToExecute.
		Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			wrapper.String())
}

func (receiver *RwxInstructionExecutor) ApplyOnPath(location string) error {
	existingRwxFileModWrapper, err := GetExistingChmodRwxWrapperPtr(location)

	if err != nil {
		return msgtype.PathErrorMessage.Error(messages.FailedToGetFileModeRwx, location)
	}

	compiledWrapper, err2 := receiver.CompiledRwxWrapperUsingFixedRwxWrapper(existingRwxFileModWrapper)

	if err2 != nil {
		funcWithLoc := "ApplyOnPath" + constants.HypenAngelRight + location

		return msgtype.
			MeaningFulError(
				msgtype.PathErrorMessage, funcWithLoc, err2)
	}

	if receiver.rwxInstruction.IsRecursive {
		return compiledWrapper.LinuxApplyRecursive(
			receiver.rwxInstruction.IsSkipOnNonExist,
			location)
	}

	return compiledWrapper.ApplyChmod(
		receiver.rwxInstruction.IsSkipOnNonExist,
		location,
	)
}

func (receiver *RwxInstructionExecutor) ApplyOnPathsDirect(locations ...string) error {
	if len(locations) == 0 {
		return nil
	}

	return receiver.ApplyOnPathsPtr(&locations)
}

func (receiver *RwxInstructionExecutor) ApplyOnPaths(locations []string) error {
	if len(locations) == 0 {
		return nil
	}

	return receiver.ApplyOnPathsPtr(&locations)
}

func (receiver *RwxInstructionExecutor) ApplyOnPathsPtr(locations *[]string) error {
	if locations == nil {
		return nil
	}

	isContinueOnError := receiver.
		rwxInstruction.
		IsContinueOnError

	if !isContinueOnError {
		return receiver.applyOnPaths(locations)
	}

	return receiver.applyOnPathsContinueOnError(locations)
}

func (receiver *RwxInstructionExecutor) applyOnPaths(locations *[]string) error {
	for _, location := range *locations {
		err := receiver.ApplyOnPath(location)

		if err != nil {
			return err
		}
	}

	return nil
}

func (receiver *RwxInstructionExecutor) applyOnPathsContinueOnError(locations *[]string) error {
	errorSlice := make([]string, constants.Zero)

	for _, location := range *locations {
		err := receiver.ApplyOnPath(location)

		if err != nil {
			errorSlice = append(errorSlice, err.Error())
		}
	}

	return msgtype.SliceToErrorPtr(&errorSlice)
}
