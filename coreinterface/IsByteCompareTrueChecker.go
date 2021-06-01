package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsByteCompareResultChecker interface {
	IsByteCompareResult(comparingValue byte, compare corecomparator.Compare) bool
}
