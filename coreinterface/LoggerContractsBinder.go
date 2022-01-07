package coreinterface

type LoggerContractsBinder interface {
	Logger
	AsLogger() Logger
}
