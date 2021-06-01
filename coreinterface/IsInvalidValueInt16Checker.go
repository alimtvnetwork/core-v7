package coreinterface

type IsInvalidValueInt16Checker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int16) bool
}
