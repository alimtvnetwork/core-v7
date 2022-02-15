package errcoreinf

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/internal/internalinterface"
)

type IsReferencesEmptyChecker interface {
	internalinterface.IsReferencesEmptyChecker
}

type HasReferencesChecker interface {
	internalinterface.HasReferencesChecker
}

type StringCompiler interface {
	internalinterface.StringCompiler
}

type HasCurrentErrorChecker interface {
	internalinterface.HasCurrentErrorChecker
}

type FullStringer interface {
	internalinterface.FullStringer
}

type TypeNamer interface {
	internalinterface.TypeNamer
}

type CodeTypeNamer interface {
	internalinterface.CodeTypeNamer
}

type TypeCodeNameStringer interface {
	internalinterface.TypeCodeNameStringer
}

type IsNullOrAnyNullChecker interface {
	internalinterface.IsNullOrAnyNullChecker
}

type BaseErrorOrCollectionWrapper interface {
	internalinterface.BaseErrorOrCollectionWrapper
}

type AddErrorer interface {
	internalinterface.AddErrorer
}

type IsErrorsCollected interface {
	internalinterface.IsErrorsCollected
}

type BaseRawErrCollectionDefiner interface {
	internalinterface.BaseRawErrCollectionDefiner
}

type FullStringWithTracesGetter interface {
	internalinterface.FullStringWithTracesGetter
}

type FullStringWithTracesIfGetter interface {
	internalinterface.FullStringWithTracesIfGetter
}

type FullOrErrorMessageGetter interface {
	internalinterface.FullOrErrorMessageGetter
}

type ErrorStringGetter interface {
	internalinterface.ErrorStringGetter
}

type ReferencesCompiledStringGetter interface {
	internalinterface.ReferencesCompiledStringGetter
}

type BaseErrorTyper interface {
	internalinterface.BaseErrorTyper
}

type BasicErrorTyper interface {
	BaseErrorTyper
	ErrorTypeAsBasicEnum() enuminf.BasicEnumer
}

type DyanmicLinqer interface {
	internalinterface.DyanmicLinqer
}

type AddManyErrorer interface {
	internalinterface.AddManyErrorer
}

type AddManyPointerErrorer interface {
	internalinterface.AddManyPointerErrorer
}

type ConditionalErrorAdder interface {
	internalinterface.ConditionalErrorAdder
}

type VarNamer interface {
	VarName() string
}

type ErrWrapperLogger interface {
	internalinterface.CompiledVoidLogger
}

type ValueDynamicGetter interface {
	ValueDynamic() interface{}
}

type ValueStringGetter interface {
	ValueString() string
}

type VariableValueStringGetter interface {
	VariableValueString() (varName, value string)
}

type VariableValueDynamicGetter interface {
	VariableValueDynamic() (varName string, value interface{})
}
type StringWithoutTyper interface {
	StringWithoutType() string
}

type VariableValueRefer interface {
	VarNamer
	ValueDynamicGetter
	VariableValueStringGetter
	VariableValueDynamicGetter
	ValueStringGetter
	StringCompiler
	StringWithoutTyper
	FullStringer
	fmt.Stringer
	corejson.Jsoner
	Serialize() ([]byte, error)
}

type BasicErrWrapper interface {
	internalinterface.BasicErrWrapper
	ErrorTypeAsBasicErrorTyper() BasicErrorTyper
	Referencer() VariableValueRefer
}

type BaseErrorWrapperCollectionDefiner interface {
	BaseErrorOrCollectionWrapper
	internalinterface.BaseErrorWrapperCollectionDefiner

	CompiledBasicErrWrapper(errType BasicErrorTyper) BasicErrWrapper
	AddErrorUsingBasicType(errType BasicErrorTyper, err error) BaseErrorWrapperCollectionDefiner
	AddBasicErrWrapper(basicErrWrapper BasicErrWrapper) BaseErrorWrapperCollectionDefiner
}

type VoidLogger interface {
	internalinterface.VoidLogger
}

type VoidTracesLogger interface {
	internalinterface.VoidTracesLogger
}

type FatalVoidLogger interface {
	internalinterface.FatalVoidLogger
}

type FatalTracesVoidLogger interface {
	internalinterface.FatalTracesVoidLogger
}

type VoidIfLogger interface {
	internalinterface.VoidIfLogger
}

type CompiledVoidLogger interface {
	internalinterface.CompiledVoidLogger
}
