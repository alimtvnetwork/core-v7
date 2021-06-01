package coreinterface

import "gitlab.com/evatix-go/core/corecomparator"

type IsInt8CompareResultChecker interface {
	IsInt8CompareResult(val, comparingValue int8, compare corecomparator.Compare) bool
}
