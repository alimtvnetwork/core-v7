package errcore

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/csvinternal"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
)

type RawErrorType string

//goland:noinspection ALL
const (
	InvalidRequestType                         RawErrorType = "Invalid : request, cannot process it."
	InvalidNullPointerType                     RawErrorType = "Invalid : null pointer, cannot process it."
	InvalidEmptyValueType                      RawErrorType = "Invalid : empty value given, cannot process it."
	OutOfRangeType                             RawErrorType = "Out of range : given value, cannot process it."
	OutOfRangeLengthType                       RawErrorType = "Out of range : given data length, cannot process it."
	InvalidEmptyPathType                       RawErrorType = "Invalid : empty path given, cannot process it."
	InvalidStringType                          RawErrorType = "Invalid : string cannot process it."
	InvalidIntegerType                         RawErrorType = "Invalid : integer cannot process it."
	InvalidFloatType                           RawErrorType = "Invalid : float cannot process it."
	InvalidType                                RawErrorType = "Invalid : type cannot process it."
	InvalidPointerType                         RawErrorType = "Invalid : pointer cannot process it."
	InvalidValueType                           RawErrorType = "Invalid : value cannot process it."
	InvalidCharType                            RawErrorType = "Invalid : character cannot process it."
	InvalidArgumentsType                       RawErrorType = "Invalid : arguments or argument cannot process it."
	InvalidAnyPathEmptyType                    RawErrorType = "Invalid : any of the given path was empty, thus cannot process it."
	UnsupportedOperatingSystemType             RawErrorType = "Unsupported : given operating system is not supported by the executable or system!"
	UnsupportedArchitectureType                RawErrorType = "Unsupported : given operating system architecture is not supported by the executable or system!"
	UnsupportedCategoryType                    RawErrorType = "Unsupported : given category or type or variant is not supported by the executable or system!"
	UnsupportedVersionType                     RawErrorType = "Unsupported : given version request is not supported by the executable or system!"
	UnsupportedInLinuxType                     RawErrorType = "Unsupported : given request is not supported in Linux!"
	UnsupportedInUnixType                      RawErrorType = "Unsupported : given request is not supported in any of Unix (including Linux, macOs, CentOS etc) operating versions!"
	UnsupportedInWindowsType                   RawErrorType = "Unsupported : given request is not supported in any of Windows operating system versions!"
	FailedToExecuteType                        RawErrorType = "Failed : request failed to execute!"
	FailedToCreateCmdType                      RawErrorType = "Failed : To create cmd, command process call. Nil pointer! Cannot proceed further."
	FailedToParseType                          RawErrorType = "Failed : request failed to parse!"
	FailedToConvertType                        RawErrorType = "Failed : request failed to convert!"
	CannotRemoveIndexesFromEmptyCollectionType RawErrorType = "Invalid operation: cannot remove indexes (either indexes are nil) or cannot remove indexes from the empty collection."
	CannotBeNegativeIndexType                  RawErrorType = "Invalid operation / index: index cannot be negative, operations canceled."
	CannotBeNegativeType                       RawErrorType = "Values or value cannot be negative value."
	CannotBeNilOrEmptyType                     RawErrorType = "Values or value cannot be nil or null or empty."
	AlreadyInitializedType                     RawErrorType = "Value is already initialized."
	KeyNotExistInMapType                       RawErrorType = "Key doesn't exist in map."
	CannotBeNilType                            RawErrorType = "Values or value cannot be nil or null."
	ShouldBePointerType                        RawErrorType = "Reference or Input needs to be a pointer!"
	CannotConvertToRwxWhereVarRwxPossibleType  RawErrorType = "Cannot convert Rwx, it had wildcards in type. It can only be converted to VarRwx."
	ShouldBeNilType                            RawErrorType = "Values or value should be nil or null."
	ShouldBeLessThanType                       RawErrorType = "Values or value should be less than the reference."
	ShouldBeGreaterThanType                    RawErrorType = "Values or value should be greater than the reference."
	ShouldBeLessThanEqualType                  RawErrorType = "Values or value should be less or equal to the reference."
	ShouldBeEqualToType                        RawErrorType = "Values or value should be equal to the reference."
	LengthShouldBeEqualToType                  RawErrorType = "Values' or value's length should be equal to the reference."
	EmptyStatusType                            RawErrorType = "Empty status found."
	NullResultType                             RawErrorType = "Null or null or nil pointer, which is unexpected."
	EmptyArrayType                             RawErrorType = "Empty array, which is unexpected."
	EmptyItemsType                             RawErrorType = "Empty items, which is unexpected."
	PathErrorType                              RawErrorType = "Path error, which is unexpected."
	PathRemoveFailedType                       RawErrorType = "Path remove failed."
	PathCreateFailedType                       RawErrorType = "Path create failed."
	FileCloseFailedType                        RawErrorType = "File close failed."
	PathExpandFailedType                       RawErrorType = "Path expand failed."
	PathChmodMismatchErrorType                 RawErrorType = "Path chmod doesn't match as expected. IsMatchesExpectation mismatch error."
	PathInvalidErrorType                       RawErrorType = "Path is missing or have permission issues in the location given."
	PathChmodApplyType                         RawErrorType = "Path chmod apply error."
	PathChmodConvertFailedType                 RawErrorType = "Path chmod convert failed to octal."
	UnexpectedValueType                        RawErrorType = "Unexpected value error, which is unexpected."
	UnexpectedType                             RawErrorType = "Unexpected type error, which is unexpected."
	UnsupportedType                            RawErrorType = "Unsupported type, none of the type matches."
	IntegerOutOfRangeType                      RawErrorType = "Integer out of range. Range, which is unexpected."
	FloatOutOfRangeType                        RawErrorType = "Float out of range. Range, which is unexpected."
	StringOutOfRangeType                       RawErrorType = "ToFileModeString out of range. Range, which is unexpected."
	ShouldBeGreaterThanEqualType               RawErrorType = "Values or value should be greater or equal to the reference."
	UnixIgnoreType                             RawErrorType = "Windows tests ignored in Unix."
	WindowsIgnoreType                          RawErrorType = "Unix tests ignored in Windows."
	ComparatorShouldBeWithinRangeType          RawErrorType = "Comparator should be within the range."
	CannotModifyCompleteResourceType           RawErrorType = "Cannot modify complete or frozen resource."
	EnumValuesOutOfRangeType                   RawErrorType = "Out of Range / Invalid Range: Enum values are are not within the range as per the expectation."
	SearchInputEmptyType                       RawErrorType = "Search Input is either null or empty."
	SearchInputOrSearchTermEmptyType           RawErrorType = "Search Input or search term either null or empty."
	EmptyResultCannotMakeJsonType              RawErrorType = "Empty result, cannot make json out of it."
	MarshallingFailedType                      RawErrorType = "Failed to marshal / parse / serialize."
	UnMarshallingFailedType                    RawErrorType = "Failed to unmarshal / parse / deserialize."
	Serialize                                  RawErrorType = "Failed to serialize or marshal convert to bytes."
	Deserialize                                RawErrorType = "Failed to deserialize or unmarshal convert to object from bytes."
	ParsingFailedType                          RawErrorType = "Failed to parse."
	TypeMismatchType                           RawErrorType = "TypeMismatchType: Type is not as expected."
	NotImplementedType                         RawErrorType = "Not Implemented: Feature / method is not implemented yet."
	NotSupportedType                           RawErrorType = "Not Supported: Feature / method is not supported yet."
	RangesOnlySupportedType                    RawErrorType = "Only Ranges: Only selected ranges supported for the function / feature."
	PathsMissingOrHavingIssuesType             RawErrorType = "Path missing or having other access issues!"
	BytesAreNilOrEmptyType                     RawErrorType = "Bytes data either nil or empty."
	ValidataionFailedType                      RawErrorType = "Validation failed!"
	LengthIssueType                            RawErrorType = "Length Issue!"
)

func GetSet(
	isCondition bool,
	trueValue RawErrorType,
	falseValue RawErrorType,
) RawErrorType {
	if isCondition {
		return trueValue
	}

	return falseValue
}

func GetSetVariant(
	isCondition bool,
	trueValue string,
	falseValue string,
) RawErrorType {
	if isCondition {
		return RawErrorType(trueValue)
	}

	return RawErrorType(falseValue)
}

func (it RawErrorType) String() string {
	return string(it)
}

func (it RawErrorType) CombineWithAnother(
	another RawErrorType,
	otherMsg string,
	reference interface{},
) RawErrorType {
	return RawErrorType(CombineWithMsgType(
		it,
		otherMsg+constants.NewLineUnix+another.String(),
		reference))
}

func (it RawErrorType) Combine(
	otherMsg string, reference interface{},
) string {
	return CombineWithMsgType(it, otherMsg, reference)
}

func (it RawErrorType) TypesAttach(
	otherMsg string,
	reflectionTypes ...interface{},
) string {
	return CombineWithMsgType(
		it,
		otherMsg,
		reflectinternal.TypeNamesString(
			true,
			reflectionTypes...))
}

func (it RawErrorType) TypesAttachErr(
	otherMsg string,
	reflectionTypes ...interface{},
) error {
	message := it.TypesAttach(otherMsg, reflectionTypes...)

	return errors.New(message)
}

func (it RawErrorType) SrcDestination(
	otherMsg string,
	srcName string, srcValue interface{},
	destinationName string, destinationValue interface{},
) string {
	reference := VarTwoNoType(
		srcName, srcValue,
		destinationName, destinationValue)

	return CombineWithMsgType(it, otherMsg, reference)
}

func (it RawErrorType) SrcDestinationErr(
	otherMsg string,
	srcName string, srcValue interface{},
	destinationName string, destinationValue interface{},
) error {
	wholeMessage := it.SrcDestination(
		otherMsg,
		srcName, srcValue,
		destinationName, destinationValue)

	return errors.New(wholeMessage)
}

func (it RawErrorType) Error(otherMsg string, reference interface{}) error {
	msg := CombineWithMsgType(it, otherMsg, reference)

	return errors.New(msg)
}

func (it RawErrorType) MsgCsvRef(
	otherMsg string,
	csvReferenceItems ...interface{},
) string {
	if len(csvReferenceItems) == 0 {
		return it.NoRef(otherMsg)
	}

	csvString := csvinternal.AnyItemsToStringDefault(
		csvReferenceItems...)

	if otherMsg == "" {
		return fmt.Sprintf(
			messageWithRefWithoutQuoteFormat,
			it.String(),
			csvString)
	}

	return fmt.Sprintf(
		messageWithOtherMsgWithRefWithoutQuoteFormat,
		it.String(),
		otherMsg,
		csvString)
}

func (it RawErrorType) MsgCsvRefError(
	otherMsg string,
	csvReferenceItems ...interface{},
) error {
	msg := it.MsgCsvRef(otherMsg, csvReferenceItems...)

	return errors.New(msg)
}

func (it RawErrorType) ErrorRefOnly(reference interface{}) error {
	msg := CombineWithMsgType(it, constants.EmptyString, reference)

	return errors.New(msg)
}

func (it RawErrorType) Expecting(expecting, actual interface{}) error {
	msg := Expecting(
		it.String(),
		expecting,
		actual)

	return errors.New(msg)
}

func (it RawErrorType) NoRef(otherMsg string) string {
	if otherMsg == "" {
		return it.String()
	}

	msg := CombineWithMsgType(it, otherMsg, nil)

	return msg
}

func (it RawErrorType) ErrorNoRefs(otherMsg string) error {
	if otherMsg == "" {
		return errors.New(it.String())
	}

	msg := CombineWithMsgType(it, otherMsg, nil)

	return errors.New(msg)
}

func (it RawErrorType) HandleUsingPanic(otherMsg string, reference interface{}) {
	msg := it.Combine(otherMsg, reference)

	panic(msg)
}
