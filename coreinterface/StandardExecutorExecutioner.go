package coreinterface

type StandardExecutorExecutioner interface {
	Run(executors ...StandardExecutor) error
	MustRun(executors ...StandardExecutor)
}
