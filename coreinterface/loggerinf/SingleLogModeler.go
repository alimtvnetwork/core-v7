package loggerinf

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/serializerinf"
	"gitlab.com/evatix-go/core/internal/internalinterface"
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
