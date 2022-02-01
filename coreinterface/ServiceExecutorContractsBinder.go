package coreinterface

type ServiceExecutorContractsBinder interface {
	ServiceExecutor
	AsServiceExecutorContractsBinder() ServiceExecutorContractsBinder
}
