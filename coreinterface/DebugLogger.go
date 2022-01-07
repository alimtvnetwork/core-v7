package coreinterface

type DebugLogger interface {
	Debug(args ...interface{}) // Debug logs a message at Debug level.
}
