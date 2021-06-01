package coreinterface

type IsInvalidValueByteChecker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value byte) bool
}
