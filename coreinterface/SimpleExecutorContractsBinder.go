package coreinterface

type SimpleExecutorContractsBinder interface {
	SimpleExecutor
	AsSimpleExecutorContractsBinder() SimpleExecutorContractsBinder
}
