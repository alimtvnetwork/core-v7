package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsInt16CompareResultChecker interface {
	IsInt16CompareResult(comparingValue int16, compare corecomparator.Compare) bool
}
