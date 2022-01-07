package coreinterface

type TaskEntityDefinerContractsBinder interface {
	TaskEntityDefiner
	AsTaskEntityDefinerContractsBinder() TaskEntityDefinerContractsBinder
}
