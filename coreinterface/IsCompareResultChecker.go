package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsCompareResultChecker interface {
	IsCompareResult(compare corecomparator.Compare) bool
}
