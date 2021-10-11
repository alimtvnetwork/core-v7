package coreasync

import "sync"

// ParallelTasksWait
//
// Wait on all tasks to be finished
func ParallelTasksWait(tasks ...VoidTask) {
	length := len(tasks)
	if length == 0 {
		return
	}

	if length == 1 {
		tasks[0]()

		return
	}

	wg := sync.WaitGroup{}

	runWrapper := func(index int) {
		tasks[index]()

		wg.Done()
	}

	wg.Add(length)

	for i := 0; i < length; i++ {
		go runWrapper(i)
	}

	wg.Wait()
}
