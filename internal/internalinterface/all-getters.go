package internalinterface

type IdentifierGetter interface {
	Identifier() string
}

type IdStringGetter interface {
	Id() string
}

type IntegerIdGetter interface {
	Id() int
}

type IdUnsignedIntegerGetter interface {
	Id() uint
}

type IdStringerWithNamer interface {
	IdAsStringer
	ToNamer
}

type IdAsStringer interface {
	IdString() string
}

type IdentifierIntegerGetter interface {
	IdentifierInt() int
}

type IdIntegerGetter interface {
	IdInteger() int
}

type UsernameGetter interface {
	Username() string
}

type CategoryNameGetter interface {
	CategoryName() string
}

type TypeNameGetter interface {
	TypeName() string
}

type TypenameStringGetter interface {
	TypenameString() string
}

type ErrorGetter interface {
	Error() error
}

type AnyValueGetter interface {
	Value() interface{}
}

type AnyAttributesGetter interface {
	AnyAttributes() interface{}
}

type AnyAttributesReflectSetter interface {
	ReflectSetAttributes(toPointer interface{}) error
}

type RawPayloadsGetter interface {
	RawPayloads() (payloads []byte, err error)
	RawPayloadsMust() (payloads []byte)
}

type ValueInt64Getter interface {
	Value() int64
}

type ValueIntegerGetter interface {
	Value() int
}

type ValueReflectSetter interface {
	ValueReflectSet(setterPtr interface{}) error
}

type ValueStringGetter interface {
	Value() string
}

type ValueStringsGetter interface {
	Value() []string
}

type ErrTypeDetailDefiner interface {
	TypenameString() string
	TypeMessage() string
	BaseErrorTypeGetter
}

type BaseErrorTypeGetter interface {
	BaseErrorTyper() BaseErrorTyper
}

type ErrorValueGetter interface {
	Value() error
}

type CompiledStackTracesStringGetter interface {
	CompiledStackTracesString() string
}

type CompiledErrorWithStackTracesGetter interface {
	CompiledErrorWithStackTraces() error
}

type FullStringWithTracesGetter interface {
	FullStringWithTraces() string
}

// FullStringWithTracesIfGetter
//
//  Returns full string with stack traces if given as true
//  Or, else just FullString returns
type FullStringWithTracesIfGetter interface {
	// FullStringWithTracesIf
	//
	//  Returns full string with stack traces if given as true
	//  Or, else just FullString returns
	FullStringWithTracesIf(isStackTraces bool) string
}

type JsonModelAnyGetter interface {
	JsonModelAny() interface{}
}

type CompiledErrorGetter interface {
	CompiledError() error
}
