package coreasync

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

func ParallelTaskWithErrorFunctionsWaited(
	isIncludeStackTrace bool,
	isContinueOnError bool,
	codeStackSkip int,
	tasks ...errcore.TaskWithErrFunc,
) error {
	length := len(tasks)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return tasks[0]()
	}

	rawErrCollection := ParallelTasksWithErrorToRawErrCollectionWaited(
		isContinueOnError,
		tasks...)

	if !isIncludeStackTrace && rawErrCollection.HasError() {
		return rawErrCollection.CompiledError()
	} else if rawErrCollection.HasError() {
		codeStacks := codestack.NewStacksDefaultCount(codeStackSkip + codestack.Skip1)

		return rawErrCollection.CompiledErrorUsingStackTraces(
			constants.NewLineUnix,
			codeStacks.ShortStrings())
	}

	return nil
}
