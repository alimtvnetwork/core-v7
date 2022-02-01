package coreinterface

type StandardExecutorContractsBinder interface {
	StandardExecutor
	AsStandardExecutorContractsBinder() StandardExecutorContractsBinder
}
