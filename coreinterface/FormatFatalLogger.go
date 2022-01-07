package coreinterface

type FormatFatalLogger interface {
	// FatalF logs a message at Fatal level
	// and process will exit with status set to 1.
	FatalF(format string, args ...interface{})
}
