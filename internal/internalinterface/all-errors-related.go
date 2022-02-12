package internalinterface

type BasicErrWrapper interface {
	ErrorHandler

	IsReferencesEmpty() bool
	HasReferences() bool

	IsNull() bool
	IsAnyNull() bool
	IsEmptyErrorChecker
	HasErrorChecker
	IsSuccessValidator
	IsEmptyChecker
	HasCurrentError() bool

	TypeNameCodeMessage() string
	RawErrorTypeValue() uint16
	CodeTypeName() string

	String() string

	ErrorMessageHandler
	ErrorValueGetter
	FullStringSplitByNewLine() []string
	FullStringWithoutReferences() string
	RawErrorTypeName() string

	StringIf(isWithRef bool) string
	FullStringWithTracesGetter
	FullStringWithTracesIfGetter
	ReferencesCompiledString() string
	FullString() string
	CompiledError() error
	CompiledErrorWithStackTracesGetter
	CompiledStackTracesStringGetter
	FullOrErrorMessage(
		isErrorMessage,
		isWithRef bool,
	) string

	JsonModelAnyGetter

	// SerializeWithoutTraces
	//
	//  Stack traces will be SKIPPED from the json bytes
	SerializeWithoutTraces() ([]byte, error)
	// Serialize
	//
	//  Should include stack traces
	Serialize() ([]byte, error)
	SerializeMust() []byte
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	Dispose()

	IsErrorEqualsChecker
}

type AddErrorer interface {
	AddError(err error)
}

type IsErrorsCollected interface {
	IsErrorsCollected(errorsItems ...error) bool
}

type BaseRawErrCollectionDefiner interface {
	Add(err error)
	AddErrorer
	IsErrorsCollected
	IsSuccessValidator
	IsValidChecker
	IsInvalidChecker
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
	HasErrorOrHasAnyErrorChecker
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
	FirstOrDefault()
	FirstOrDefaultCompiledError()
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

	IsSuccess() bool
	IsValid() bool
	IsFailed() bool
	GetAsError() error

	ErrWrapperLogger

	ErrorHandler

	// HandleWithMsg Skip if no error.
	HandleWithMsg(msg string)
	// Dispose After dispose nothing will work, everything be removed from memory.
	Dispose()
	JsonModelAny() interface{}
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
