package coreasync

// ParallelTasks
//
// no waiting on any tasks
func ParallelTasks(tasks ...VoidTask) {
	length := len(tasks)
	if length == 0 {
		return
	}

	for i := 0; i < length; i++ {
		go tasks[i]()
	}
}
