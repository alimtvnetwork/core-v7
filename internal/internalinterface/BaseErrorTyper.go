package internalinterface

type BaseErrorTyper interface {
	NameWithNameEqualer
	NameValue() string
	IsValid() bool
	IsInvalid() bool
	IsRawValue(rawValue uint16) bool
	IsNoError() bool
	IsEmptyError() bool
	HasError() bool
	Combine(
		additionalMessage,
		varName string,
		val interface{},
	) string
	CombineNoRefs(
		additionalMessage string,
	) string
	Error(
		additionalMessage,
		varName string,
		val interface{},
	) error
	ErrorReferences(
		additionalMessage string,
		references ...interface{},
	) error
	ErrorNoRefs(
		additionalMessage string,
	) error
	Panic(
		additionalMessage,
		varName string,
		val interface{},
	)
	PanicNoRefs(
		additionalMessage string,
	)
	ErrTypeDetailDefiner() ErrTypeDetailDefiner
	// CodeWithTypeName
	//
	// 	errconsts.ErrorCodeHyphenTypeNameFormat  = "(#%d - %s)"
	CodeWithTypeName() string
	TypeName() string
	CodeTypeNameWithCustomMessage(
		customMessage string,
	) string
	ReferencesCsv(
		additionalMessage string,
		references ...interface{},
	) string
	ReferencesLines(
		additionalMessage string,
		referencesLines ...interface{},
	) string
	ReferencesLinesError(
		additionalMessage string,
		referencesLines ...interface{},
	) error
	ReferencesCsvError(
		additionalMessage string,
		references ...interface{},
	) error
	ShortReferencesCsv(
		references ...interface{},
	) string
	ShortReferencesCsvError(
		references ...interface{},
	) error
	// ExplicitCodeCodeValueTypeName
	//
	// 	errconsts.ErrorCodeWithTypeNameFormat = "(Code - #%d) : %s"
	ExplicitCodeCodeValueTypeName() string
	// CodeTypeNameWithReference
	//
	// 	errconsts.ErrorCodeHyphenTypeNameWithReferencesFormat = "(#%d - %s - {%v})"
	CodeTypeNameWithReference(
		referenceLine string,
	) string

	// CodeTypeNameWithReferences
	//
	// 	errconsts.ErrorCodeHyphenTypeNameWithReferencesFormat = "(#%d - %s - {%v})"
	CodeTypeNameWithReferences(
		references ...string,
	) string
	RawValue() uint16
	Value() int16
	ValueInt16() int16
	ValueInt() int
	ValueUInt() uint
}
