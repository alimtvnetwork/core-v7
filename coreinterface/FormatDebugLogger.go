package coreinterface

type FormatDebugLogger interface {
	DebugF(formatter string, args ...interface{}) // Debug logs a message at Debug level.
}
