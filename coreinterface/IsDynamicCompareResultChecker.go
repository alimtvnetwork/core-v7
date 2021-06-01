package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsDynamicCompareResultChecker interface {
	IsDynamicCompareResult(input interface{}, compare corecomparator.Compare) bool
}
