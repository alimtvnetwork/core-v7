package entityinf

type TaskEntityDefinerContractsBinder interface {
	TaskEntityDefiner
	AsTaskEntityDefinerContractsBinder() TaskEntityDefinerContractsBinder
}
