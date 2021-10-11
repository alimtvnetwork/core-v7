package coreasync

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/errcore"
)

type waited struct{}

func (it *waited) ParallelTasks(
	isIncludeStackTraces,
	isContinueOnError bool,
	tasks ...errcore.TaskWithErrFunc,
) error {
	if len(tasks) == 0 {
		return nil
	}

	if len(tasks) == 1 {
		return tasks[0]()
	}

	return ParallelTaskWithErrorFunctionsWaited(
		isIncludeStackTraces,
		isContinueOnError,
		codestack.Skip1,
		tasks...)
}

func (it *waited) ParallelVoidTasks(
	tasks ...VoidTask,
) {
	if len(tasks) == 0 {
		return
	}

	if len(tasks) == 1 {
		tasks[0]()

		return
	}

	ParallelTasksWait(
		tasks...)
}
