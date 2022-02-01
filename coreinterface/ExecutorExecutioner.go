package coreinterface

type ExecutorExecutioner interface {
	Run(executors ...Executor) error
	MustRun(executors ...Executor)
}
