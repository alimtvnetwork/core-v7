package coreinterface

type FormatInfoLogger interface {
	InfoF(format string, args ...interface{}) // Info logs a message at Info level.
}
