package errcoreinf

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/internal/internalinterface"
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

type GetAsBasicWrapperGetter interface {
	GetAsBasicWrapper() BasicErrWrapper
}

type GetAsBasicWrapperUsingTyperGetter interface {
	GetAsBasicWrapperUsingTyper(errorTyper BasicErrorTyper) BasicErrWrapper
}

type StackTracer interface {
	StackTracesJsonResult() *corejson.Result
	StackTraces() string

	// NewStackTraces
	//
	// returns new code stack (only) traces string
	NewStackTraces(stackSkip int) string
	// NewDefaultStackTraces
	//
	// returns new code stack (only) traces string
	NewDefaultStackTraces() string
	// NewStackTracesJsonResult
	//
	// returns new code stack (only) traces json
	NewStackTracesJsonResult(stackSkip int) *corejson.Result
	// NewDefaultStackTracesJsonResult
	//
	// returns new code stack (only) traces json
	NewDefaultStackTracesJsonResult() *corejson.Result
}

type BaseErrorOrCollectionWrapper interface {
	internalinterface.BaseErrorOrCollectionWrapper

	// JsonResultWithoutTraces
	//
	// Mostly used for testing cases
	JsonResultWithoutTraces() *corejson.Result

	IsCollect(another BaseErrorOrCollectionWrapper) bool
	IsCollectedAny(anotherItems ...BaseErrorOrCollectionWrapper) bool
	IsCollectOn(isCollect bool, another BaseErrorOrCollectionWrapper) bool
	IsEmptyAll(anotherItems ...BaseErrorOrCollectionWrapper) bool

	GetAsBasicWrapperGetter
	GetAsBasicWrapperUsingTyperGetter
	StackTracer

	coreinterface.ReflectSetter
	corejson.Jsoner
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

// FullOrErrorMessageGetter
//
//	isErrorMessage : true will return only the error or else full string
//	isWithRef : refers to include reference or not
type FullOrErrorMessageGetter interface {
	internalinterface.FullOrErrorMessageGetter
}

type ErrorStringGetter interface {
	internalinterface.ErrorStringGetter
}

type ReferencesCompiledStringGetter interface {
	internalinterface.ReferencesCompiledStringGetter
}

// ExplicitCodeValueNamer
//
//	returns string in format "(Code - #%d) : %s"
type ExplicitCodeValueNamer interface {
	// ExplicitCodeValueName
	//
	// 	returns string in format "(Code - #%d) : %s"
	ExplicitCodeValueName() string
}

type CodeTypeNameWithReferencer interface {
	// CodeTypeNameWithReference
	//
	// 	returns string in format
	// 	- "(#%d - %s - {%v})" : (error-code - name - referenceLine)
	CodeTypeNameWithReference(
		referenceLine string,
	) string
}

type JsonModelAnyGetter interface {
	JsonModelAny() interface{}
}

type CategoryNamer interface {
	CategoryName() string
}

type BaseErrorTyper interface {
	internalinterface.BaseErrorTyper
	ExplicitCodeValueNamer
	CodeTypeNameWithReferencer
	JsonModelAnyGetter
	CategoryNamer
	coreinterface.AllSerializer
	IsErrorTyperEqual(errTyper BaseErrorTyper) bool
}

type BaseErrorTypeGetter interface {
	BaseErrorTyper() BaseErrorTyper
}

type BasicErrorTyperGetter interface {
	BasicErrorTyper() BasicErrorTyper
}

type ErrTypeDetailDefiner interface {
	TypenameString() string
	TypeMessage() string
	BaseErrorTypeGetter
}

type BasicErrorTyper interface {
	BaseErrorTyper
	enuminf.BasicEnumer

	ErrTypeDetailDefiner() ErrTypeDetailDefiner
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

type Referencer interface {
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
	coreinterface.AllSerializer

	IsEqualReferencer(ref Referencer) bool

	coreinterface.ReflectSetter
}

type StringsGetter interface {
	Strings() []string
}

type ReferencesListGetter interface {
	ReferencesList() []Referencer
}

type ReferenceCollectionDefiner interface {
	ReferencerCollection() []Referencer
	coreinterface.HasAnyItemChecker
	coreinterface.IsEmptyChecker
	coreinterface.LengthGetter
	coreinterface.CountGetter

	AddVarVal(varName string, val interface{}) ReferenceCollectionDefiner
	AddReferencer(ref Referencer) ReferenceCollectionDefiner
	AddReferences(references ...Referencer) ReferenceCollectionDefiner

	coreinterface.MapStringAnyGetter
	coreinterface.MapStringStringGetter
	coreinterface.AllSerializer
	corejson.Jsoner

	StringsGetter
	fmt.Stringer
	Compile() string

	coreinterface.ReflectSetter
	ReferencesListGetter

	CloneNewDefiner() ReferenceCollectionDefiner
}

type BasicErrWrapper interface {
	internalinterface.BasicErrWrapper
	BaseErrorOrCollectionWrapper
	ErrorTypeAsBasicErrorTyper() BasicErrorTyper
	ReferencesCollection() ReferenceCollectionDefiner
	coreinterface.ReflectSetter
	ReferencesListGetter

	IsBasicErrEqual(another BasicErrWrapper) bool
	MergeNewErrInf(right BaseErrorOrCollectionWrapper) BaseErrorOrCollectionWrapper
	MergeNewMessage(newMessage string) BaseErrorOrCollectionWrapper
	CloneInterface() BasicErrWrapper
}

type CompiledBasicErrWrapper interface {
	CompiledToGenericBasicErrWrapper() BasicErrWrapper
	CompiledToBasicErrWrapper(errType BasicErrorTyper) BasicErrWrapper

	CompiledToErrorWithTraces(errType BasicErrorTyper) error
}

type BaseErrorWrapperCollectionDefiner interface {
	BaseErrorOrCollectionWrapper
	internalinterface.BaseErrorWrapperCollectionDefiner

	// LinesWithoutTraces
	//
	//  returns lines without traces
	LinesWithoutTraces() []string

	CompiledBasicErrWrapper
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

type ShouldBeErrorVerifier interface {
	ShouldBeError(right interface{}) error
}

type ShouldBeMessageVerifier interface {
	ShouldBe(right interface{}) string
}

type LeftShouldBeMessageVerifier interface {
	ShouldBe(left, right interface{}) string
}

type LeftShouldBeErrorVerifier interface {
	ShouldBeError(left, right interface{}) error
}

type CompleteSuccesser interface {
	CompleteSuccess() BaseErrorOrCollectionWrapper
}

type MustCompleteSuccesser interface {
	CompleteSuccessMust()
}

type CompleteFailurer interface {
	CompleteFailure() BaseErrorOrCollectionWrapper
}

type MustCompleteFailurer interface {
	CompleteFailureMust()
}

type GenericErrorCompiler interface {
	BaseErrorTypeGetter

	CompiledMessage() string
	JsonString() string

	Length() int
	IsEmpty() bool
	HasAnyItem() bool
	HasAnyIssues() bool

	CompileString() string
	Serialize() ([]byte, error)

	Format(format string) string
	CompileError() error

	corejson.Jsoner

	ToGenericErr() BaseErrorOrCollectionWrapper
	MustBeEmptier
	CompiledVoidLogger

	fmt.Stringer
}

type MustBeEmptier interface {
	MustBeSuccess() bool
	MustBeEmpty()
	HandleError()
}

type CompleteSuccessJsoner interface {
	CompleteSuccessJson() *corejson.Result
}

type CompleteFailureJsoner interface {
	CompleteFailureJson() *corejson.Result
}

type ErrorCompleter interface {
	CompleteReceiveError(completionTyper enuminf.CompletionStateTyper) error
	CompleteUsingErrReceiveError(completionTyper enuminf.CompletionStateTyper) error

	CompleteSuccesser
	CompleteFailurer
	MustCompleteSuccesser
	MustCompleteFailurer

	CompleteSuccessJsoner
	CompleteFailureJsoner

	Complete(completionTyper enuminf.CompletionStateTyper) BaseErrorOrCollectionWrapper
	CompleteUsingErr(err error) BaseErrorOrCollectionWrapper
	CompleteUsingErrWithTitle(title string, err error) BaseErrorOrCollectionWrapper
	CompleteUsingBaseErrOrCollection(baseErrOrCollection BaseErrorOrCollectionWrapper) BaseErrorOrCollectionWrapper
	CompleteUsingBasicErrWrapper(basicErrWrapper BasicErrWrapper) BaseErrorOrCollectionWrapper
	CompleteUsingBaseErrorWrapperCollectionDefiner(baseErrOrCollection BaseErrorWrapperCollectionDefiner) BaseErrorOrCollectionWrapper
	CompleteUsingBaseRawErrCollectionDefiner(baseErrOrCollection BaseRawErrCollectionDefiner) BaseErrorOrCollectionWrapper
}
