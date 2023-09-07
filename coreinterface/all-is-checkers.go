package coreinterface

import (
	"reflect"

	"gitlab.com/auk-go/core/internal/internalinterface"
)

type IsReflectionTypeChecker interface {
	IsManyReflectionOfType(typeOf reflect.Type, dynamicItems ...interface{}) bool
	IsReflectionOfType(dynamic interface{}, typeOf reflect.Type) bool
	IsReflectionOfTypeName(dynamic interface{}, typeOfName string) bool
}

type EmptyChecker interface {
	IsEmpty() bool
	HasAnyItemChecker
}

type DynamicDataHasChecker interface {
	HasDynamic(searchItem interface{}) bool
	HasDynamicAll(searchTerms ...interface{}) bool
	HasDynamicAny(searchTerms ...interface{}) bool
}

type BooleanChecker interface {
	IsAnyByOrder(booleans ...bool) bool
	HasAll(searchTerms ...string) bool
	HasAny(searchTerms ...string) bool
	HasItemsWithoutIssues() bool
}

type IsAnyByOrder interface {
	IsAnyByOrder(booleans ...bool) bool
}

type StringHasAllChecker interface {
	HasAll(searchTerms ...string) bool
}

type StringHasAnyChecker interface {
	HasAny(searchTerms ...string) bool
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

type StringHasChecker interface {
	Has(search string) bool
}

type StringHasCombineChecker interface {
	StringHasChecker
	StringHasAllChecker
	StringHasAnyChecker
	HasAnyItemChecker
}

type SimpleValidInvalidChecker interface {
	IsValidChecker
	IsInvalidChecker
	InvalidMessageGetter
}

type IsValidInvalidChecker interface {
	IsValidChecker
	IsInvalidChecker
}

type SimpleValidatorIssueChecker interface {
	SimpleValidInvalidChecker
	HasAnyItemChecker
	InvalidDirectErrorGetter
}

type StringIsAnyOfChecker interface {
	IsAnyOf(value string, checkingItems ...string) bool
}

// IsAnyNullChecker
//
// Returns true if self is null or values is null
// Values have to be null to have true return.
// False: Any empty slice will return false.
type IsAnyNullChecker interface {
	// IsAnyNull
	//
	// Returns true if self is null or values is null
	// Values have to be null to have true return.
	// False: Any empty slice will return false.
	IsAnyNull() bool
}

type IsApplyFuncBinder interface {
	IsApply() (isSuccess bool)
}

type IsByteValidRangeUsingArgsChecker interface {
	IsByteValidRange(val byte) bool
}

type IsByteValueValidChecker interface {
	IsByteValueValid(value byte) bool
}

type IsDynamicContainsChecker interface {
	IsDynamicContains(item interface{}) bool
}

type IsDynamicContainsInCollectionChecker interface {
	IsDynamicContainsInCollection(collection, item interface{}) bool
}

type IsDynamicItemValidChecker interface {
	IsDynamicItemValid(item interface{}) bool
	IsDynamicItemsValid(items ...interface{}) bool
}

type IsDynamicNullChecker interface {
	// IsDynamicNull may check using reflection that data is nil.
	IsDynamicNull(dynamic interface{}) bool
}

type IsDynamicValidRangeUsingArgsChecker interface {
	IsDynamicValidRange(val, max, min interface{}) bool
}

type IsDynamicValueValidChecker interface {
	IsDynamicValueValid(value interface{}) bool
}

type IsEmptyChecker interface {
	internalinterface.IsEmptyChecker
}

type IsDefinedChecker interface {
	IsDefined() bool
}

type IsEmptyErrorChecker interface {
	IsEmptyError() bool
}

type IsEmptyOrWhitespaceChecker interface {
	IsEmptyOrWhitespace() bool
}

type IsFailedChecker interface {
	// IsFailed has error or any other issues, or alias for HasIssues or HasError
	IsFailed() bool
}

type IsInt8ValidRangeUsingArgsChecker interface {
	IsInt8ValidRange(val int8) bool
}

type IsInt8ValueValidChecker interface {
	IsInt8ValueValid(value int8) bool
}

type IsInt16ValidRangeUsingArgsChecker interface {
	IsInt16ValidRange(val int16) bool
}

type IsInt16ValueValidChecker interface {
	IsInt16ValueValid(value int16) bool
}

type IsInt32ValidRangeUsingArgsChecker interface {
	IsInt32ValidRange(val int32) bool
}

type IsInt32ValueValidChecker interface {
	IsInt32ValueValid(value int32) bool
}

type IsInt64ValueValidChecker interface {
	IsInt64ValueValid(value int64) bool
}

type IsIntValidRangeUsingArgsChecker interface {
	IsIntValidRange(val int) bool
}

type IsValidChecker interface {
	// IsValid similar or alias for IsSuccessChecker
	IsValid() bool
}

type IsInvalidChecker interface {
	IsInvalid() bool
}

type IsInvalidValueByteChecker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value byte) bool
}

type IsInvalidValueInt8Checker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int8) bool
}

type IsInvalidValueInt16Checker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int16) bool
}

type IsInvalidValueInt32Checker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int32) bool
}

type IsInvalidValueIntChecker interface {
	// IsInvalidValue delegates and acts as !IsWithinRange(value)
	IsInvalidValue(value int) bool
}

type IsNilChecker interface {
	IsNil() bool
}

type IsNullChecker interface {
	IsNull() bool
}

type IsNullOrEmptyChecker interface {
	IsNullOrEmpty() bool
}

type IsOutOfRangeByteChecker interface {
	IsOutOfRange(n byte)
}

type IsPointerChecker interface {
	IsPointer() bool
}

type IsReflectKindChecker interface {
	IsReflectKind(checkingKind reflect.Kind) bool
}

type IsReflectTypeOfChecker interface {
	IsReflectTypeOf(typeRequest reflect.Type) bool
}

type IsStringContainsChecker interface {
	IsContains(contains string) bool
}

type IsStringEqualChecker interface {
	IsEqual(equalString string) bool
}

type IsStringValidRangeUsingArgsChecker interface {
	IsStringValidRange(val, max, min string) bool
}

type IsSuccessChecker interface {
	// IsSuccess No error
	IsSuccess() bool
}

type IsSuccessValidator interface {
	IsValidChecker
	IsSuccessChecker
	IsFailedChecker
}

type IsWithinRangeByteChecker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value byte) bool
}

type IsWithinRangeInt8Checker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int8) bool
}

type IsWithinRangeInt16Checker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int16) bool
}

type IsWithinRangeInt32Checker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int32) bool
}

type IsWithinRangeIntChecker interface {
	// IsWithinRange r.Min >= value && value <= r.Max
	//
	// Or, r.Start >= value && value <= r.End
	IsWithinRange(value int) bool
}

type Int16IsAnyOfChecker interface {
	IsAnyOf(value int16, checkingItems ...int16) bool
}

type ByteIsAnyOfChecker interface {
	IsAnyOf(value byte, checkingItems ...byte) bool
}

type IsEnabledChecker interface {
	IsEnabled() bool
}

type IsDisabledChecker interface {
	IsDisabled() bool
}

type IsEnableAllChecker interface {
	IsEnableAll() bool
}

type IsEnableAnyChecker interface {
	IsEnableAny() bool
}

type IsEnableAnyByNamesChecker interface {
	IsEnableAnyByNames(enabledNames ...string) bool
}

type IsDisableAllChecker interface {
	IsDisableAll() bool
}

type IsDisableAnyChecker interface {
	IsDisableAny() bool
}

type IsDisableAnyByNamesChecker interface {
	IsDisableAnyByNames(disabledNames ...string) bool
}

type IsFlagsEnabledByNamesChecker interface {
	IsFlagsEnabledByNames(enabledNames ...string) bool
}

type IsFlagsDisabledByNamesChecker interface {
	IsFlagsDisabledByNames(disabledNames ...string) bool
}

type IsEnableDisableConditionChecker interface {
	IsEnableAllChecker
	IsEnableAnyChecker
	IsEnableAnyByNamesChecker

	IsDisableAllChecker
	IsDisableAnyChecker
	IsDisableAnyByNamesChecker
}

type IsKeyMissingChecker interface {
	IsMissingKey(key string) bool
}

type IsCompletedChecker interface {
	IsCompleted() bool
}

type IsCompletedLockChecker interface {
	IsCompletedLock() bool
}

type IsCompletedLockUnlockChecker interface {
	IsCompletedChecker
	IsCompletedLockChecker
}

type IsMissingKeyChecker interface {
	IsMissingKey(key string) bool
}

type IsValueStringChecker interface {
	IsValueString() bool
}

type IsValueTypeOfChecker interface {
	IsValueTypeOf(rfType reflect.Type) bool
}
