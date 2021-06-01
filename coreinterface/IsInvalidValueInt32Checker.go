package coreinterface

type IsInvalidValueInt32Checker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int32) bool
}
