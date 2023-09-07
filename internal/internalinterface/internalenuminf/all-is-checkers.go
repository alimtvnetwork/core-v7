package internalenuminf

type IsValidChecker interface {
	// IsValid similar or alias for IsSuccessChecker
	IsValid() bool
}

type IsInvalidChecker interface {
	IsInvalid() bool
}

type IsValidInvalidChecker interface {
	IsValidChecker
	IsInvalidChecker
}

type IsNameEqualer interface {
	IsNameEqual(name string) bool
}

// IsAnyNameOfChecker
//
//  Returns true if any of the name matches.
type IsAnyNameOfChecker interface {
	// IsAnyNamesOf
	//
	//  Returns true if any of the name matches.
	IsAnyNamesOf(names ...string) bool
}

type IsAnyValueByteEqualer interface {
	IsAnyValuesEqual(anyByteValues ...byte) bool
}

type IsAnyValueIntegerEqualer interface {
	IsAnyValuesEqual(anyValues ...int) bool
}

type IsAnyValueInteger8Equaler interface {
	IsAnyValuesEqual(anyValues ...int8) bool
}

type IsAnyValueInteger16Equaler interface {
	IsAnyValuesEqual(anyValues ...int16) bool
}

type IsAnyValueInteger32Equaler interface {
	IsAnyValuesEqual(anyValues ...int32) bool
}

type IsByteValueEqualer interface {
	IsByteValueEqual(value byte) bool
}

type IsValueIntegerEqualer interface {
	IsValueEqual(value int) bool
}

type IsValueInteger8Equaler interface {
	IsValueEqual(value int8) bool
}

type IsValueInteger16Equaler interface {
	IsValueEqual(value int16) bool
}

type IsValueInteger32Equaler interface {
	IsValueEqual(value int32) bool
}

type RangeValidateChecker interface {
	// RangesInvalidMessage get invalid message
	RangesInvalidMessage() string
	// RangesInvalidErr get invalid message error
	RangesInvalidErr() error
	// IsValidRange Is with in the range as expected.
	IsValidRange() bool
	// IsInvalidRange Is out of the ranges expected.
	IsInvalidRange() bool
}

type IsEnumEqualer interface {
	IsEnumEqual(enum BasicEnumer) bool
}

type IsAnyEnumsEqualer interface {
	IsAnyEnumsEqual(enums ...BasicEnumer) bool
}
