package chmodhelper

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodinstruction"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/msgtype"
)

type RwxInstructionExecutor struct {
	rwxInstruction *chmodinstruction.RwxInstruction
	varWrapper     *VarWrapper
}

func (receiver *RwxInstructionExecutor) IsVarWrapper() bool {
	return !receiver.varWrapper.IsFixedType()
}

func (receiver *RwxInstructionExecutor) IsFixedWrapper() bool {
	return receiver.varWrapper.IsFixedType()
}

func (receiver *RwxInstructionExecutor) CompiledWrapper(mode os.FileMode) (*Wrapper, error) {
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

func (receiver *RwxInstructionExecutor) CompiledWrapperUsingWrapper(
	wrapper *Wrapper,
) (*Wrapper, error) {
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
	fileModWrapper, err := GetExistingChmodWrapperPtr(location)

	if err != nil {
		return msgtype.FileErrorMessage.Error(messages.FailedToGetFileModeRwx, location)
	}

	compiledWrapper, err2 := receiver.CompiledWrapperUsingWrapper(fileModWrapper)

	if err2 != nil {
		funcWithLoc := "ApplyOnPath" + constants.HypenAngelRight + location

		return msgtype.
			MeaningFulError(
				msgtype.FileErrorMessage, funcWithLoc, err2)
	}

	if receiver.rwxInstruction.IsRecursive {
		return compiledWrapper.UnixApplyRecursive(
			location,
			receiver.rwxInstruction.IsSkipOnNonExist)
	}

	return compiledWrapper.ApplyChmod(
		location,
		receiver.rwxInstruction.IsSkipOnNonExist)
}

func (receiver *RwxInstructionExecutor) ApplyOnPaths(locations *[]string) error {
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

	return msgtype.SliceToError(&errorSlice)
}
