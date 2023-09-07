package baseactioninf

type SimpleExecutorContractsBinder interface {
	SimpleExecutor
	AsSimpleExecutorContractsBinder() SimpleExecutorContractsBinder
}
