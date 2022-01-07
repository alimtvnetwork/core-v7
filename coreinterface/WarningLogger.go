package coreinterface

type WarningLogger interface {
	Warn(args ...interface{}) // Warn logs a message at Warning level.
}
