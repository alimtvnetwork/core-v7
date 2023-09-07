package loggerinf

import (
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/pathextendinf"
	"gitlab.com/auk-go/core/internal/internalinterface"
)

type AttrPersistentLogger interface {
	internalinterface.IdStringerWithNamer
	LogPathInfo() pathextendinf.PathInfoer
	IsRotating() bool
	IsDbLogger() bool
	IsFileLogger() bool

	DynamicConfig() interface{}
	ConfigReflectSetTo(toPointer interface{}) error

	// PersistentLoggerTyper
	//
	//  Which type of persistent logger
	PersistentLoggerTyper() enuminf.BasicEnumer
}
