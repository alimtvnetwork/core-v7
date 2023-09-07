package internalinterface

type LogTypeCheckerBinder interface {
	LogTypeChecker
	AsLogTypeChecker() LogTypeChecker
}
