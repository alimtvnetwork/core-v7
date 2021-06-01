package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsInt32CompareResultChecker interface {
	IsInt32CompareResult(comparingValue int32, compare corecomparator.Compare) bool
}
