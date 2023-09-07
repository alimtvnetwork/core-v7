package loggerinf

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/coreinterface/serializerinf"
	"gitlab.com/auk-go/core/internal/internalinterface"
)

type SingleLogModeler interface {
	internalinterface.IdentifierGetter
	PersistentIdGetter
	ParentPersistentIdGetter

	HasParentChecker
	HasModelChecker
	hasErrorChecker

	coreinterface.CategoryRevealer
	coreinterface.CategoryRevealerGetter

	FilterTyper() enuminf.BasicEnumer
	LevelTyper() enuminf.LogLevelTyper
	LogTyper() enuminf.LoggerTyper
	BasicErrorTyper() errcoreinf.BasicErrorTyper
	ModelTyper() enuminf.BasicEnumer
	EntityTypeName() string

	ModelBytesGetter

	LogMessageGetter
	CompiledAttributesGetter
	CallerGetter

	SpecificValuerGetter
	ErrorAsBasicErrWrapperGetter
	ReflectSetter
	serializerinf.Deserializer

	corejson.JsonContractsBinder
}
