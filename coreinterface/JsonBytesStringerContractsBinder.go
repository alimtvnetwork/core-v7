package coreinterface

type JsonBytesStringerContractsBinder interface {
	JsonByter
	JsonCombineStringer
	AsJsonBytesStringerContractsBinder() JsonBytesStringerContractsBinder
}
