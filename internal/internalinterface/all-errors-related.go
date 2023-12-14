package internalinterface

import (
	"fmt"
)

type IsReferencesEmptyChecker interface {
	IsReferencesEmpty() bool
}

type HasReferencesChecker interface {
	HasReferences() bool
}

type StringCompiler interface {
	Compile() string
}

type HasCurrentErrorChecker interface {
	HasCurrentError() bool
}

type FullStringer interface {
	FullString() string
}

type TypeNamer interface {
	TypeName() string
}

type IsNullOrAnyNullChecker interface {
	IsNull() bool
	IsAnyNull() bool
}

type CodeTypeNamer interface {
	CodeTypeName() string
}

type TypeCodeNameStringer interface {
	TypeCodeNameString() string
}

// SerializeWithoutTracesGetter
//
//	Stack traces will be SKIPPED from the json bytes
type SerializeWithoutTracesGetter interface {
	SerializeWithoutTraces() ([]byte, error)
}

type FullOrErrorMessageGetter interface {
	FullOrErrorMessage(
		isErrorMessage,
		isWithRef bool,
	) string
}

type ReferencesCompiledStringGetter interface {
	ReferencesCompiledString() string
}

type ErrorStringGetter interface {
	ErrorString() string
}

type BaseErrorOrCollectionWrapper interface {
	ErrorHandler
	IsNullOrAnyNullChecker
	HasErrorChecker
	IsEmptyChecker
	ErrorStringGetter
	IsSuccessValidator
	IsInvalidChecker
	HasErrorOrHasAnyErrorChecker
	HasAnyIssues() bool
	IsDefined() bool

	// StringCompiler
	//
	//  error wrapper compiles to string with traces.
	StringCompiler
	ErrorHandler
	ErrorMessageHandler

	FullStringer
	CompiledErrorGetter
	FullStringWithTracesGetter
	FullStringWithTracesIfGetter
	ReferencesCompiledStringGetter
	CompiledErrorWithStackTracesGetter
	CompiledStackTracesStringGetter

	CompiledJsonErrorWithStackTracesGetter
	CompiledJsonStringWithStackTracesGetter

	FullStringSplitByNewLine() []string
	FullStringWithoutReferences() string

	MustBeEmptyError()
	MustBeSafe()

	// SerializeWithoutTracesGetter
	//
	//  Stack traces will be SKIPPED from the json bytes
	SerializeWithoutTracesGetter
	// Serialize
	//
	//  Should include stack traces
	Serialize() ([]byte, error)
	SerializeMust() []byte
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	ErrorValueGetter

	Dispose()

	CompiledVoidLogger
	IsCollectionTyper

	ReflectSetTo(toPtr interface{}) error

	fmt.Stringer
}

// IsCollectionTyper
//
//	returns true if current type is collection
type IsCollectionTyper interface {
	// IsCollectionType
	//
	//  returns true if current type is collection
	IsCollectionType() bool
}

// BasicErrWrapper
//
// IsEmpty:
//
//	Refers to no error for print or doesn't treat this as error.
//
//	Conditions (true):
//	    - if Wrapper nil, Or,
//	    - if Wrapper is StaticEmptyPtr, Or,
//	    - if Wrapper .errorType is IsNoError(), Or,
//	    - if Wrapper .currentError NOT nil and Wrapper .references.IsEmpty()
type BasicErrWrapper interface {
	BaseErrorOrCollectionWrapper
	IsReferencesEmptyChecker
	HasReferencesChecker
	IsEmptyErrorChecker
	HasCurrentErrorChecker

	TypeNameCodeMessage() string
	TypeNameWithCustomMessage(customMessage string) string
	RawErrorTypeValue() uint16
	TypeNamer
	CodeTypeNamer
	TypeCodeNameStringer

	IsErrorMessageEqual(msg string) bool
	// IsErrorMessage
	//
	// If error IsEmpty then returns false regardless
	IsErrorMessage(msg string, isCaseSensitive bool) bool
	ErrorValueGetter
	StringIf(isWithRef bool) string

	FullOrErrorMessageGetter
	JsonModelAnyGetter
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error

	IsErrorEqualsChecker
}

type AddErrorer interface {
	AddError(err error)
}

type IsErrorsCollected interface {
	IsErrorsCollected(errorsItems ...error) bool
}

type BaseRawErrCollectionDefiner interface {
	BaseErrorOrCollectionWrapper
	Add(err error)
	AddMessages(messages ...string)
	AddMsg(message string)
	AddErrorWithMessage(err error, message string)
	AddErrorWithMessageRef(err error, message string, reference interface{})
	Fmt(format string, v ...interface{})
	FmtIf(isAdd bool, format string, v ...interface{})
	References(message string, v ...interface{})
	AddErrorer
	IsErrorsCollected
	AddWithTraceRef(
		err error,
		traces []string,
		referenceItem interface{},
	)
	AddWithCompiledTraceRef(
		err error,
		compiledTrace string,
		referenceItem interface{},
	)
	AddWithRef(
		err error,
		referenceItem interface{},
	)
	AddManyErrorer
	ConditionalErrorAdder
	// AddString
	//
	//  Empty string will be ignored
	AddString(
		message string,
	)
	AddStringSliceAsErr(
		errSliceStrings ...string,
	)
	CommonSliceDefiner
	StringUsingJoiner
	StringUsingJoinerAdditional(joiner, additionalMessage string) string
	CompiledErrorGetter
	CompiledErrorUsingJoiner(joiner string) error
	CompiledErrorUsingJoinerAdditionalMessage(joiner, additionalMessage string) error
	CompiledErrorUsingStackTraces(joiner string, stackTraces []string) error
	StringWithAdditionalMessage(additionalMessage string) string
}

type DyanmicLinqer interface {
	FirstDynamic() interface{}
	LastDynamic() interface{}
	FirstOrDefaultError() error
	FirstOrDefaultFullMessage() string
	LastOrDefaultCompiledError() error
	LastOrDefaultError() error
	LastOrDefaultFullMessage() string
	FirstOrDefaultDynamic() interface{}
	LastOrDefaultDynamic() interface{}
	SkipDynamic(skippingItemsCount int) interface{}
	TakeDynamic(takeDynamicItems int) interface{}
	LimitDynamic(limit int) interface{}
}

type AddManyErrorer interface {
	// AddErrors no error then skip adding
	AddErrors(errs ...error)
}

type AddManyPointerErrorer interface {
	// AddErrorsPtr no error then skip adding
	AddErrorsPtr(errs *[]error)
}

type ConditionalErrorAdder interface {
	// ConditionalAddError adds error if isAdd and error not nil.
	ConditionalAddError(
		isAdd bool,
		err error,
	)
}

type BaseErrorWrapperCollectionDefiner interface {
	BaseErrorOrCollectionWrapper
	DyanmicLinqer
	CommonSliceDefiner
	LastIndex() int
	HasIndex(index int) bool

	AddErrorer
	AddManyErrorer
	AddManyPointerErrorer
	ConditionalErrorAdder

	HasError() bool
	IsEmpty() bool
	Length() int

	ToString(
		isIncludeStakeTraces,
		isIncludeHeader bool,
	) string
	ToStrings(
		isIncludeStakeTraces,
		isIncludeHeader bool,
	) []string

	Strings(isIncludeStakeTraces bool) []string

	String() string
	StringIf(isIncludeTraces bool) string
	StringStackTracesWithoutHeader() string
	DisplayStringWithTraces() string

	DisplayStringWithLimitTraces(limit int) string
	LogDisplayStringWithLimitTraces(limit int)
	FullStringWithTracesIfGetter

	StringWithoutHeader() string
	StringsWithoutHeader() []string

	LinesIf(
		isIncludeReferences bool,
	) []string

	StringsWithoutReferencePlusHeader() []string
	StringsIf(isIncludeStakeTraces bool) []string

	FullStrings() []string
	FullStringsWithTraces() []string
	FullStringsWithLimitTraces(limit int) []string

	Errors() []error
	CompiledErrors() []error
	CompiledErrorsWithStackTraces() []error

	GetAsError() error

	// HandleWithMsg Skip if no error.
	HandleWithMsg(msg string)
}
