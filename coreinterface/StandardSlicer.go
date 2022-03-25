package coreinterface

import "gitlab.com/evatix-go/core/coredata/corejson"

type StandardSlicer interface {
	BasicSlicer
	ListStringsGetter
	JsonCombineStringer
	corejson.JsonContractsBinder
}
