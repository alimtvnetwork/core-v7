package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsIntCompareResultChecker interface {
	IsIntCompareResult(comparingValue int, compare corecomparator.Compare) bool
}
