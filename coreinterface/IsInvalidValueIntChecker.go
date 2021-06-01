package coreinterface

type IsInvalidValueIntChecker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int) bool
}
