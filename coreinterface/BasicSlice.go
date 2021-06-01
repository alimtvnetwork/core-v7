package coreinterface

import "gitlab.com/evatix-go/core/coredata/corejson"

type BasicSlicer interface {
	LengthGetter
	CountGetter
	EmptyChecker
	LastIndexGetter
	HasIndexChecker
	ListStringsGetter
	JsonCombineStringer
	ItemAtRemover
	HasAnyItemChecker
	corejson.JsonContractsBinder
}
