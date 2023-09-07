package loggerinf

import (
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/internal/internalinterface"
)

type ModelBytesGetter interface {
	ModelBytes() []byte
}

type LogMessageGetter interface {
	LogMessage() string
}

type CompiledAttributesGetter interface {
	CompiledAttributes() string
}

type ErrorAsBasicErrWrapperGetter interface {
	ErrorAsBasicErrWrapper() errcoreinf.BasicErrWrapper
}

type CallerGetter interface {
	Caller() Caller
}

type PersistentIdGetter interface {
	PersistentId() uint
}

type ParentPersistentIdGetter interface {
	ParentPersistentId() uint
}

type SpecificValuerGetter interface {
	SpecificValuer() SpecificValuer
}

type RawPayloadsGetter interface {
	internalinterface.RawPayloadsGetter
}

type FilterTextGetter interface {
	FilterText() string
}
