package coreinterface

type CompareInt16Enumer interface {
	IsEqual(comparingValue int16) bool
	IsLess(comparingValue int16) bool
	IsLessEqual(comparingValue int16) bool
	IsGreater(comparingValue int16) bool
	IsGreaterEqual(comparingValue int16) bool
	IsNotEqual(comparingValue int16) bool
	IsInt16CompareResultChecker
}
