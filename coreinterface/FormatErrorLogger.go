package coreinterface

type FormatErrorLogger interface {
	ErrorF(format string, args ...interface{}) // Error logs a message at Error level.
}
