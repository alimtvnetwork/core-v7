package msgtype

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
)

type Variation string

//goland:noinspection ALL
const (
	InvalidRequest                         Variation = "Invalid : request, cannot process it."
	InvalidNullPointer                     Variation = "Invalid : null pointer, cannot process it."
	InvalidEmptyValue                      Variation = "Invalid : empty value given, cannot process it."
	OutOfRange                             Variation = "Out of range : given value, cannot process it."
	OutOfRangeLength                       Variation = "Out of range : given data length, cannot process it."
	InvalidEmptyPathErrorMessage           Variation = "Invalid : empty path given, cannot process it."
	InvalidStringErrorMessage              Variation = "Invalid : string cannot process it."
	InvalidIntegerErrorMessage             Variation = "Invalid : integer cannot process it."
	InvalidFloatErrorMessage               Variation = "Invalid : float cannot process it."
	InvalidTypeErrorMessage                Variation = "Invalid : type cannot process it."
	InvalidPointerErrorMessage             Variation = "Invalid : pointer cannot process it."
	InvalidValueErrorMessage               Variation = "Invalid : value cannot process it."
	InvalidCharErrorMessage                Variation = "Invalid : character cannot process it."
	InvalidArgumentsErrorMessage           Variation = "Invalid : arguments or argument cannot process it."
	InvalidAnyPathEmptyErrorMessage        Variation = "Invalid : any of the given path was empty, thus cannot process it."
	UnsupportedOperatingSystem             Variation = "Unsupported : given operating system is not supported by the executable or system!"
	UnsupportedArchitecture                Variation = "Unsupported : given operating system architecture is not supported by the executable or system!"
	UnsupportedCategory                    Variation = "Unsupported : given category or type or variant is not supported by the executable or system!"
	UnsupportedVersion                     Variation = "Unsupported : given version request is not supported by the executable or system!"
	UnsupportedInLinux                     Variation = "Unsupported : given request is not supported in Linux!"
	UnsupportedInUnix                      Variation = "Unsupported : given request is not supported in any of Unix (including Linux, macOs, CentOS etc) operating versions!"
	UnsupportedInWindows                   Variation = "Unsupported : given request is not supported in any of Windows operating system versions!"
	FailedToExecute                        Variation = "Failed : request failed to execute!"
	FailedToCreateCmd                      Variation = "Failed : To create cmd, command process call. Nil pointer! Cannot proceed further."
	FailedToParse                          Variation = "Failed : request failed to parse!"
	FailedToConvert                        Variation = "Failed : request failed to convert!"
	CannotRemoveIndexesFromEmptyCollection Variation = "Invalid operation: cannot remove indexes (either indexes are nil) or cannot remove indexes from the empty collection."
	CannotBeNegativeIndex                  Variation = "Invalid operation / index: index cannot be negative, operations canceled."
	CannotBeNegativeMessage                Variation = "Values or value cannot be negative value."
	CannotBeNilOrEmptyMessage              Variation = "Values or value cannot be nil or null or empty."
	AlreadyInitialized                     Variation = "Value is already initialized."
	KeyNotExistInMap                       Variation = "Key doesn't exist in map."
	CannotBeNilMessage                     Variation = "Values or value cannot be nil or null."
	ShouldBePointer                        Variation = "Reference or Input needs to be a pointer!"
	CannotConvertToRwxWhereVarRwxPossible  Variation = "Cannot convert Rwx, it had wildcards in type. It can only be converted to VarRwx."
	ShouldBeNilMessage                     Variation = "Values or value should be nil or null."
	ShouldBeLessThanMessage                Variation = "Values or value should be less than the reference."
	ShouldBeGreaterThanMessage             Variation = "Values or value should be greater than the reference."
	ShouldBeLessThanEqualMessage           Variation = "Values or value should be less or equal to the reference."
	ShouldBeEqualToMessage                 Variation = "Values or value should be equal to the reference."
	LengthShouldBeEqualToMessage           Variation = "Values' or value's length should be equal to the reference."
	EmptyStatusMessage                     Variation = "Empty status found."
	NullResultMessage                      Variation = "Null or null or nil pointer, which is unexpected."
	EmptyArrayMessage                      Variation = "Empty array, which is unexpected."
	EmptyItemsMessage                      Variation = "Empty items, which is unexpected."
	PathErrorMessage                       Variation = "Path error, which is unexpected."
	PathRemoveFailedMessage                Variation = "Path remove failed."
	PathCreateFailedMessage                Variation = "Path create failed."
	FileCloseFailedMessage                 Variation = "File close failed."
	PathExpandFailedMessage                Variation = "Path expand failed."
	PathChmodMismatchErrorMessage          Variation = "Path chmod doesn't match as expected. IsMatchesExpectation mismatch error."
	PathInvalidErrorMessage                Variation = "Path is missing or have permission issues in the location given."
	PathChmodApplyMessage                  Variation = "Path chmod apply error."
	PathChmodConvertFailedMessage          Variation = "Path chmod convert failed to octal."
	UnexpectedValueErrorMessage            Variation = "Unexpected value error, which is unexpected."
	UnexpectedTypeErrorMessage             Variation = "Unexpected type error, which is unexpected."
	IntegerOutOfRangeMessage               Variation = "Integer out of range. Range, which is unexpected."
	FloatOutOfRangeMessage                 Variation = "Float out of range. Range, which is unexpected."
	StringOutOfRangeMessage                Variation = "ToFileModeString out of range. Range, which is unexpected."
	ShouldBeGreaterThanEqualMessage        Variation = "Values or value should be greater or equal to the reference."
	UnixIgnoreMessage                      Variation = "Windows tests ignored in Unix."
	WindowsIgnoreMessage                   Variation = "Unix tests ignored in Windows."
	ComparatorShouldBeWithinRange          Variation = "Comparator should be within the range."
	CannotModifyCompleteResource           Variation = "Cannot modify complete or frozen resource."
	EnumValuesOutOfRange                   Variation = "Out of Range / Invalid Range: Enum values are are not within the range as per the expectation."
	SearchInputEmpty                       Variation = "Search Input is either null or empty."
	SearchInputOrSearchTermEmpty           Variation = "Search Input or search term either null or empty."
	EmptyResultCannotMakeJson              Variation = "Empty result, cannot make json out of it."
	MarshallingFailed                      Variation = "Failed to marshal / parse / serialize."
	UnMarshallingFailed                    Variation = "Failed to unmarshal / parse / deserialize."
	ParsingFailed                          Variation = "Failed to parse."
	TypeMismatch                           Variation = "TypeMismatch: Type is not as expected."
	NotImplemented                         Variation = "Not Implemented: Feature / method is not implemented yet."
	NotSupported                           Variation = "Not Supported: Feature / method is not supported yet."
	RangesOnlySupported                    Variation = "Only Ranges: Only selected ranges supported for the function / feature."
	PathsMissingOrHavingIssues             Variation = "Path missing or having other access issues!"
	JsonResultBytesAreNilOrEmpty           Variation = "Json Result: Bytes data either nil or empty, cannot process the data for the given resource."
	ValidataionFailed                      Variation = "Validation failed!"
	LengthIssue                            Variation = "Length Issue!"
)

func GetSet(
	isCondition bool,
	trueValue Variation,
	falseValue Variation,
) Variation {
	if isCondition {
		return trueValue
	}

	return falseValue
}

func GetSetVariant(
	isCondition bool,
	trueValue string,
	falseValue string,
) Variation {
	if isCondition {
		return Variation(trueValue)
	}

	return Variation(falseValue)
}

func (variation Variation) String() string {
	return string(variation)
}

func (variation Variation) CombineWithAnother(
	another Variation,
	otherMsg string,
	reference interface{},
) Variation {
	return Variation(CombineWithMsgType(
		variation,
		otherMsg+constants.NewLineUnix+another.String(),
		reference))
}

func (variation Variation) Combine(otherMsg string, reference interface{}) string {
	return CombineWithMsgType(variation, otherMsg, reference)
}

func (variation Variation) Error(otherMsg string, reference interface{}) error {
	msg := CombineWithMsgType(variation, otherMsg, reference)

	return errors.New(msg)
}

func (variation Variation) ErrorRefOnly(reference interface{}) error {
	msg := CombineWithMsgType(variation, constants.EmptyString, reference)

	return errors.New(msg)
}

func (variation Variation) Expecting(expecting, actual interface{}) error {
	msg := Expecting(
		variation.String(),
		expecting,
		actual)

	return errors.New(msg)
}

func (variation Variation) ErrorNoRefs(otherMsg string) error {
	msg := CombineWithMsgType(variation, otherMsg, nil)

	return errors.New(msg)
}

func (variation Variation) HandleUsingPanic(otherMsg string, reference interface{}) {
	msg := CombineWithMsgType(variation, otherMsg, reference)

	panic(msg)
}
