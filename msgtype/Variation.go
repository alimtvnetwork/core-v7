package msgtype

import (
	"errors"
	"strings"
)

type Variation string

//goland:noinspection ALL
const (
	InvalidRequest                  Variation = "Invalid : request, cannot process it."
	InvalidNullPointer              Variation = "Invalid : null pointer, cannot process it."
	InvalidEmptyValue               Variation = "Invalid : empty value given, cannot process it."
	OutOfRange                      Variation = "Out of range : given value, cannot process it."
	OutOfRangeLength                Variation = "Out of range : given data length, cannot process it."
	InvalidEmptyPathErrorMessage    Variation = "Invalid : empty path given, cannot process it."
	InvalidStringErrorMessage       Variation = "Invalid : string cannot process it."
	InvalidIntegerErrorMessage      Variation = "Invalid : integer cannot process it."
	InvalidFloatErrorMessage        Variation = "Invalid : float cannot process it."
	InvalidTypeErrorMessage         Variation = "Invalid : type cannot process it."
	InvalidPointerErrorMessage      Variation = "Invalid : pointer cannot process it."
	InvalidValueErrorMessage        Variation = "Invalid : value cannot process it."
	InvalidCharErrorMessage         Variation = "Invalid : character cannot process it."
	InvalidArgumentsErrorMessage    Variation = "Invalid : arguments or argument cannot process it."
	InvalidAnyPathEmptyErrorMessage Variation = "Invalid : any of the given path was empty, thus cannot process it."
	UnsupportedOperatingSystem      Variation = "Unsupported : given operating system is not supported by the executable or system!"
	UnsupportedArchitecture         Variation = "Unsupported : given operating system architecture is not supported by the executable or system!"
	UnsupportedCategory             Variation = "Unsupported : given category or type or variant is not supported by the executable or system!"
	UnsupportedVersion              Variation = "Unsupported : given version request is not supported by the executable or system!"
	UnsupportedInLinux              Variation = "Unsupported : given request is not supported in Linux!"
	UnsupportedInUnix               Variation = "Unsupported : given request is not supported in any of Unix (including Linux, macOs, CentOS etc) operating versions!"
	UnsupportedInWindows            Variation = "Unsupported : given request is not supported in any of Windows operating system versions!"
	FailedToExecute                 Variation = "Failed : request failed to execute!"
	FailedToParse                   Variation = "Failed : request failed to parse!"
	FailedToConvert                 Variation = "Failed : request failed to convert!"
	CannotBeNegativeMessage         Variation = "Values or value cannot be negative value."
	CannotBeNilOrEmptyMessage       Variation = "Values or value cannot be nil or null or empty."
	CannotBeNilMessage              Variation = "Values or value cannot be nil or null."
	ShouldBeLessThanMessage         Variation = "Values or value should be less than the reference."
	ShouldBeGreaterThanMessage      Variation = "Values or value should be greater than the reference."
	ShouldBeLessThanEqualMessage    Variation = "Values or value should be less or equal to the reference."
	ShouldBeEqualToMessage          Variation = "Values or value should be equal to the reference."
	LengthShouldBeEqualToMessage    Variation = "Values' or value's length should be equal to the reference."
	EmptyStatusMessage              Variation = "Empty status found."
	NullResultMessage               Variation = "Null or null or nil pointer, which is unexpected."
	EmptyArrayMessage               Variation = "Empty array, which is unexpected."
	EmptyItemsMessage               Variation = "Empty items, which is unexpected."
	FileErrorMessage                Variation = "File error, which is unexpected."
	UnexpectedValueErrorMessage     Variation = "Unexpected value error, which is unexpected."
	UnexpectedTypeErrorMessage      Variation = "Unexpected type error, which is unexpected."
	IntegerOutOfRangeMessage        Variation = "Integer out of range. Range, which is unexpected."
	FloatOutOfRangeMessage          Variation = "Float out of range. Range, which is unexpected."
	StringOutOfRangeMessage         Variation = "ToFileModeString out of range. Range, which is unexpected."
	ShouldBeGreaterThanEqualMessage Variation = "Values or value should be greater or equal to the reference."
	UnixIgnoreMessage               Variation = "Windows tests ignored in Unix."
	WindowsIgnoreMessage            Variation = "Unix tests ignored in Windows."
	ComparatorShouldBeWithinRanghe  Variation = "Comparator should be within the range."
)

func (variation Variation) String() string {
	return string(variation)
}

func (variation Variation) Combine(otherMsg string, reference interface{}) string {
	return CombineWithMsgType(variation, otherMsg, reference)
}

func (variation Variation) Error(otherMsg string, reference interface{}) error {
	msg := CombineWithMsgType(variation, otherMsg, reference)

	return errors.New(strings.ToLower(msg))
}

func (variation Variation) HandleUsingPanic(otherMsg string, reference interface{}) {
	msg := CombineWithMsgType(variation, otherMsg, reference)

	panic(msg)
}
