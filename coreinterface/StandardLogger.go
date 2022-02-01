package coreinterface

type StandardLogger interface {
	LoggerWithFormatLogger
	ConditionalLogger
}
