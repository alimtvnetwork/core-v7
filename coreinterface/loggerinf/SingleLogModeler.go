package loggerinf

import (
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/internal/internalinterface"
)

type SingleLogModeler interface {
	internalinterface.IdentifierGetter
	PersistentId() uint
	ParentPersistentId() uint
	HasParent() bool
	FilterTyper() enuminf.BasicEnumer
	LevelTyper() enuminf.LogLevelTyper
	LogTyper() enuminf.LoggerTyper
	BasicErrorTyper() errcoreinf.BasicErrorTyper
	ModelTyper() enuminf.BasicEnumer
	ModelBytes() []byte
	HasModel() bool
	LogMessage() string
	CompiledAttributes() string
	Caller() Caller
	HasError() bool
	SpecificValuer() SpecificValuer
	ErrorAsBasicErrWrapper() errcoreinf.BasicErrWrapper
}
