package loggerinf

import (
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/pathextendinf"
	"gitlab.com/evatix-go/core/internal/internalinterface"
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
