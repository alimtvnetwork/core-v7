package coreinterface

type DynamicLinqContractsBinder interface {
	DynamicLinq
	AsDynamicLinqContractsBinder() DynamicLinqContractsBinder
}
