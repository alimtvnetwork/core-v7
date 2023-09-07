package loggerinf

import (
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/serializerinf"
)

type StandardLoggerGetter interface {
	StandardLogger() StandardLogger
}

type BaseLogDefiner interface {
	coreinterface.TypeNameGetter
	coreinterface.MessageGetter
}

type LogDefiner interface {
	BaseLogDefiner
	coreinterface.RawMessageBytesGetter
	serializerinf.FieldBytesToPointerDeserializer
}

type LogDefinerWriter interface {
	LogWrite(logEntity LogDefiner) error
	LogWriteMust(logEntity LogDefiner)
	LogWriteUsingStackSkip(
		stackSkipIndex int,
		logEntity LogDefiner,
	) error
}

type LogTypeWriter interface {
	LogTypeWrite(logType string, v ...interface{}) error
	LogTypeWriteMust(logType string, v ...interface{})

	LogTypeWriteStackSkip(
		stackSkipIndex int,
		logType string,
		v ...interface{},
	) error

	LogTypeWriteStackSkipMust(
		stackSkipIndex int,
		logType string,
		v ...interface{},
	)
}
