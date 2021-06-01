package coreinterface

type IsDynamicValidRangeUsingArgsChecker interface {
	IsDynamicValidRange(val, max, min interface{}) bool
}
