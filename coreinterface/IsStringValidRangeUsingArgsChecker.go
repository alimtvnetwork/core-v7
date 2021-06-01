package coreinterface

type IsStringValidRangeUsingArgsChecker interface {
	IsStringValidRange(val, max, min string) bool
}
