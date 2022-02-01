package coreinterface

type SimpleExecutor interface {
	NameGetter
	TypeNameGetter
	Executor
	IsApplyFuncBinder
}
