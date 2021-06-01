package coreinterface

type CompareByteEnumer interface {
	IsEqual(comparingValue byte) bool
	IsLess(comparingValue byte) bool
	IsLessEqual(comparingValue byte) bool
	IsGreater(comparingValue byte) bool
	IsGreaterEqual(comparingValue byte) bool
	IsNotEqual(comparingValue byte) bool
	IsByteCompareResultChecker
}
