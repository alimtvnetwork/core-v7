package corejson

type JsonContractsBinder interface {
	SimpleJsoner
	AsJsonContractsBinder() JsonContractsBinder
}
