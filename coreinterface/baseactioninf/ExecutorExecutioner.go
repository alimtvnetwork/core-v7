package baseactioninf

type ExecutorExecutioner interface {
	Run(executors ...Executor) error
	MustRun(executors ...Executor)
}
