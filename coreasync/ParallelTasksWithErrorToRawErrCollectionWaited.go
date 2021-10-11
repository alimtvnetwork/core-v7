package coreasync

import (
	"sync"

	"gitlab.com/evatix-go/core/errcore"
)

func ParallelTasksWithErrorToRawErrCollectionWaited(
	isContinueOnError bool,
	tasks ...errcore.TaskWithErrFunc,
) *errcore.RawErrCollection {
	length := len(tasks)
	if length == 0 {
		return nil
	}

	rawErr := errcore.RawErrCollection{}

	if length == 1 {
		err := tasks[0]()
		rawErr.Add(err)

		return &rawErr
	}

	isErrorFound := false
	wg := sync.WaitGroup{}
	locker := sync.Mutex{}

	runWrapper := func(index int) {
		defer wg.Done()

		if !isContinueOnError && isErrorFound {
			return
		}

		err := tasks[index]()

		if err != nil {
			locker.Lock()

			rawErr.Add(err)

			isErrorFound = true

			locker.Unlock()
		}
	}

	for i := 0; i < length; i++ {
		if !isContinueOnError && isErrorFound {
			break
		}

		wg.Add(1)
		go runWrapper(i)

		if !isContinueOnError && isErrorFound {
			break
		}
	}

	wg.Wait()

	if rawErr.HasError() {
		return &rawErr
	}

	return nil
}
