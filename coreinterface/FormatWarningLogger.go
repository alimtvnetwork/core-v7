package coreinterface

type FormatWarningLogger interface {
	WarnF(format string, args ...interface{}) // Warn logs a message at Warning level.
}
