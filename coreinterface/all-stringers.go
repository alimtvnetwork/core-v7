package coreinterface

type AllKeysStringer interface {
	AllKeys() []string
}

type AllKeysSortedStringer interface {
	AllKeysSorted() []string
}

type JsonCombineStringer interface {
	JsonStringer
	// MustJsonStringer panic if any error
	MustJsonStringer
}

type BuildStringer interface {
	Build() string
}

type ByteToStringer interface {
	String(input byte) string
}

type MustJsonStringer interface {
	JsonStringMust() string
}

type NameValueStringer interface {
	NameValue() string
}

type FullStringer interface {
	FullString() string
}

type FullStringWithTracer interface {
	FullStringWithTraces() string
}

type ToValueStringer interface {
	Value() string
}

// ToNumberStringer
//
//  It returns string number value.
//
// Examples:
//  - ToNumberString() -> "1"  if the value is 1
//  - ToNumberString() -> "10" if the value is 10
type ToNumberStringer interface {
	// ToNumberString
	//
	//  It returns string number value.
	//
	// Examples:
	//  - ToNumberString() -> "1"  if the value is 1
	//  - ToNumberString() -> "10" if the value is 10
	ToNumberString() string
}

type ValidationCheckerWithStringer interface {
	Stringer
	IsInvalidChecker
	IsValidChecker
}

type SafeStringer interface {
	SafeString() string
}
