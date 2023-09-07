package coreinterface

import "gitlab.com/auk-go/core/coredata/corejson"

type StandardSlicer interface {
	BasicSlicer
	ListStringsGetter
	JsonCombineStringer
	corejson.JsonContractsBinder
}
