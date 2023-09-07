package internalinterface

type LoggerTyper interface {
	SimpleEnumer
	LogTypeChecker
	IsLogNameEqual(name string) bool
	IsLogValueEqual(logVal byte) bool
}
