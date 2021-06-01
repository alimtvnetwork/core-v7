package coreinterface

type CompareInt8Enumer interface {
	IsEqual(comparingValue int8) bool
	IsLess(comparingValue int8) bool
	IsLessEqual(comparingValue int8) bool
	IsGreater(comparingValue int8) bool
	IsGreaterEqual(comparingValue int8) bool
	IsNotEqual(comparingValue int8) bool
	IsInt8CompareResultChecker
}
