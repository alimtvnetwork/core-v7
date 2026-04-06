package errcore

import (
	"errors"
	"fmt"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/csvinternal"
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
	NotFound                                   RawErrorType = "not found"
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
	CannotBeNegativeIndexType                  RawErrorType = "Invalid operation or index: index cannot be negative, operations canceled."
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
	PathExist                                  RawErrorType = "file path exist but expect to be missing or clear"
	ParsingFailed                              RawErrorType = "parsing failed"
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
	EnumValuesOutOfRangeType                   RawErrorType = "Out of Range or Invalid Range: Enum values are are not within the range as per the expectation."
	SearchInputEmptyType                       RawErrorType = "Search Input is either null or empty."
	SearchInputOrSearchTermEmptyType           RawErrorType = "Search Input or search term either null or empty."
	EmptyResultCannotMakeJsonType              RawErrorType = "Empty result, cannot make json out of it."
	MarshallingFailedType                      RawErrorType = "Failed to marshal or serialize."
	UnMarshallingFailedType                    RawErrorType = "Failed to unmarshal or deserialize."
	Serialize                                  RawErrorType = "Failed to serialize or marshal convert to bytes."
	Deserialize                                RawErrorType = "Failed to deserialize or unmarshal convert to object from bytes."
	ParsingFailedType                          RawErrorType = "Failed to parse."
	TypeMismatchType                           RawErrorType = "TypeMismatchType: Type is not as expected."
	NotImplementedType                         RawErrorType = "Not Implemented: Feature or method is not implemented yet."
	NotSupportedType                           RawErrorType = "Not Supported: Feature or method is not supported yet."
	RangesOnlySupportedType                    RawErrorType = "Only Ranges: Only selected ranges supported for the function or feature."
	PathsMissingOrHavingIssuesType             RawErrorType = "Path missing or having other access issues!"
	BytesAreNilOrEmptyType                     RawErrorType = "Bytes data either nil or empty."
	// Deprecated: Use ValidationFailedType instead.
	ValidataionFailedType RawErrorType = "Validation failed!"
	ValidationFailedType  RawErrorType = "Validation failed!"
	LengthIssueType       RawErrorType = "Length Issue!"
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
	reference any,
) RawErrorType {
	return RawErrorType(
		CombineWithMsgTypeNoStack(
			it,
			otherMsg+constants.NewLineUnix+another.String(),
			reference,
		),
	)
}

func (it RawErrorType) Combine(
	otherMsg string, reference any,
) string {
	return CombineWithMsgTypeNoStack(it, otherMsg, reference)
}

func (it RawErrorType) TypesAttach(
	otherMsg string,
	reflectionTypes ...any,
) string {
	return CombineWithMsgTypeNoStack(
		it,
		otherMsg,
		typesNamesString(
			reflectionTypes...,
		),
	)
}

func (it RawErrorType) TypesAttachErr(
	otherMsg string,
	reflectionTypes ...any,
) error {
	message := it.TypesAttach(otherMsg, reflectionTypes...)

	return errors.New(message)
}

func (it RawErrorType) SrcDestination(
	otherMsg string,
	srcName string, srcValue any,
	destinationName string, destinationValue any,
) string {
	reference := VarTwoNoType(
		srcName, srcValue,
		destinationName, destinationValue,
	)

	return CombineWithMsgTypeNoStack(it, otherMsg, reference)
}

func (it RawErrorType) SrcDestinationErr(
	otherMsg string,
	srcName string, srcValue any,
	destinationName string, destinationValue any,
) error {
	wholeMessage := it.SrcDestination(
		otherMsg,
		srcName, srcValue,
		destinationName, destinationValue,
	)

	return errors.New(wholeMessage)
}

func (it RawErrorType) Error(otherMsg string, reference any) error {
	msg := CombineWithMsgTypeNoStack(it, otherMsg, reference)

	return StackEnhance.MsgToErrSkip(1, msg)
}

func (it RawErrorType) ErrorSkip(skipStack int, otherMsg string, reference any) error {
	msg := CombineWithMsgTypeNoStack(it, otherMsg, reference)

	return StackEnhance.MsgToErrSkip(1+skipStack, msg)
}

func (it RawErrorType) Fmt(
	format string,
	v ...any,
) error {
	if format == "" && len(v) == 0 {
		return it.ErrorRefOnly(nil)
	}

	compiledMessage := fmt.Sprintf(
		format,
		v...,
	)

	msg := CombineWithMsgTypeNoStack(it, compiledMessage, nil)

	return StackEnhance.MsgToErrSkip(1, msg)
}

func (it RawErrorType) FmtIf(
	isError bool,
	format string,
	v ...any,
) error {
	isNoError := !isError

	if isNoError {
		return nil
	}

	return it.Fmt(format, v...)
}

func (it RawErrorType) MergeError(
	err error,
) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", it.String(), err)
}

func (it RawErrorType) MergeErrorWithMessage(
	err error,
	message string,
) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s %s: %w", it.String(), message, err)
}

func (it RawErrorType) MergeErrorWithMessageRef(
	err error,
	message string,
	reference any,
) error {
	if err == nil {
		return nil
	}

	refMsg := CombineWithMsgTypeNoStack(it, message, reference)

	return fmt.Errorf("%s: %w", refMsg, err)
}

func (it RawErrorType) MergeErrorWithRef(
	err error,
	reference any,
) error {
	if err == nil {
		return nil
	}

	refMsg := CombineWithMsgTypeNoStack(it, "", reference)

	return fmt.Errorf("%s: %w", refMsg, err)
}

func (it RawErrorType) MsgCsvRef(
	otherMsg string,
	csvReferenceItems ...any,
) string {
	if len(csvReferenceItems) == 0 {
		return it.NoRef(otherMsg)
	}

	csvString := csvinternal.AnyItemsToStringDefault(
		csvReferenceItems...,
	)

	if otherMsg == "" {
		return fmt.Sprintf(
			messageWithRefWithoutQuoteFormat,
			it.String(),
			csvString,
		)
	}

	return fmt.Sprintf(
		messageWithOtherMsgWithRefWithoutQuoteFormat,
		it.String(),
		otherMsg,
		csvString,
	)
}

func (it RawErrorType) MsgCsvRefError(
	otherMsg string,
	csvReferenceItems ...any,
) error {
	msg := it.MsgCsvRef(otherMsg, csvReferenceItems...)
	msg = StackEnhance.MsgSkip(1, msg)

	return errors.New(msg)
}

func (it RawErrorType) ErrorRefOnly(reference any) error {
	msg := CombineWithMsgTypeStackTrace(it, constants.EmptyString, reference)

	return errors.New(msg)
}

func (it RawErrorType) Expecting(expecting, actual any) error {
	msg := Expecting(
		it.String(),
		expecting,
		actual,
	)

	msg = StackEnhance.MsgSkip(1, msg)

	return errors.New(msg)
}

func (it RawErrorType) NoRef(otherMsg string) string {
	if otherMsg == "" {
		return StackEnhance.MsgSkip(1, it.String())
	}

	msg := CombineWithMsgTypeStackTrace(it, otherMsg, nil)

	return msg
}

func (it RawErrorType) ErrorNoRefs(otherMsg string) error {
	return it.ErrorNoRefsSkip(1, otherMsg)
}

func (it RawErrorType) ErrorNoRefsSkip(stack int, otherMsg string) error {
	if otherMsg == "" {
		return StackEnhance.MsgToErrSkip(1+stack, it.String())
	}

	msg := CombineWithMsgTypeNoStack(it, otherMsg, nil)

	return StackEnhance.MsgToErrSkip(1+stack, msg)
}

func (it RawErrorType) HandleUsingPanic(otherMsg string, reference any) {
	msg := it.Combine(otherMsg, reference)

	panic(msg)
}
