package coreinterface

type IsInvalidValueInt8Checker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int8) bool
}
