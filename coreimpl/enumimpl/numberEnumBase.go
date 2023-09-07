package enumimpl

import (
	"errors"
	"fmt"
	"sort"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/coreonce"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/csvinternal"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type numberEnumBase struct {
	actualValueRanges      interface{}
	stringRanges           []string
	rangesCsvString        coreonce.StringOnce
	rangesInvalidMessage   coreonce.StringOnce
	invalidError           coreonce.ErrorOnce
	integerEnumRangesOnce  coreonce.IntegersOnce
	typeName               string
	minAny, maxAny         interface{}
	minStr, maxStr         string
	keyAnyValues           []KeyAnyVal
	rangesDynamicMap       map[string]interface{}
	rangesIntegerStringMap map[int]string
}

// newNumberEnumBase
//
//  @actualRangesAnyType : []Byte, []int, []int8... not pointer
//
//  Lengths must match stringRanges and actualRangesAnyType
func newNumberEnumBase(
	typeName string,
	actualRangesAnyType interface{},
	nameRanges []string,
	min, max interface{},
) numberEnumBase {
	if nameRanges == nil {
		errcore.MeaningfulErrorHandle(
			errcore.CannotBeNilType,
			"newNumberEnumBase",
			errors.New("StringRanges cannot be nil"))
	}

	integerEnumRangesOnce := coreonce.NewIntegersOnce(func() []int {
		return IntegersRangesOfAnyVal(actualRangesAnyType)
	})

	_, isString := actualRangesAnyType.([]string)

	rangesToCsvOnce := coreonce.NewStringOnce(func() string {
		if isString {
			clonedList := strutilinternal.Clone(nameRanges)
			sort.Strings(clonedList)

			return csvinternal.StringsToStringDefaultNoQuotations(
				clonedList...)
		}

		allKeyValues := KeyAnyValues(
			nameRanges,
			actualRangesAnyType)
		length := len(allKeyValues)
		newMap := make(map[int]string, length)
		integersSlice := make([]int, length)

		for i, keyAnyVal := range allKeyValues {
			valueInt := keyAnyVal.ValInt()
			newMap[valueInt] = keyAnyVal.String()
			integersSlice[i] = valueInt
		}

		sort.Ints(integersSlice)

		newSortedSlice := make([]string, length)

		for i, valueInt := range integersSlice {
			nameValue := newMap[valueInt]
			newSortedSlice[i] = nameValue
		}

		return csvinternal.StringsToStringDefaultNoQuotations(
			newSortedSlice...)
	})

	invalidMessageOnce := coreonce.NewStringOnce(func() string {
		msg := errcore.EnumRangeNotMeet(
			min,
			max,
			rangesToCsvOnce.Value())

		return msg
	})

	return numberEnumBase{
		actualValueRanges:    actualRangesAnyType,
		stringRanges:         nameRanges,
		rangesCsvString:      rangesToCsvOnce,
		rangesInvalidMessage: invalidMessageOnce,
		invalidError: coreonce.NewErrorOnce(func() error {
			return errors.New(invalidMessageOnce.Value())
		}),
		integerEnumRangesOnce: integerEnumRangesOnce,
		typeName:              typeName,
		minAny:                min,
		maxAny:                max,
	}
}

func (it numberEnumBase) MinMaxAny() (min, max interface{}) {
	return it.minAny, it.maxAny
}

func (it *numberEnumBase) MinValueString() string {
	if it.minStr != "" {
		return it.minStr
	}

	it.minStr = convAnyValToString(it.minAny)

	return it.minStr
}

func (it numberEnumBase) MinInt() int {
	return ConvEnumAnyValToInteger(it.minAny)
}

func (it numberEnumBase) MaxInt() int {
	return ConvEnumAnyValToInteger(it.maxAny)
}

func (it numberEnumBase) AllNameValues() []string {
	return AllNameValues(
		it.StringRanges(),
		it.actualValueRanges)
}

func (it numberEnumBase) RangesMap() map[int]string {
	return it.DynamicMap().ConvMapIntegerString()
}

func (it numberEnumBase) OnlySupportedErr(supportedNames ...string) error {
	return OnlySupportedErr(
		it.StringRanges(),
		supportedNames...)
}

func (it numberEnumBase) OnlySupportedMsgErr(errMessage string, supportedNames ...string) error {
	return errcore.ConcatMessageWithErr(
		errMessage,
		it.OnlySupportedErr(supportedNames...))
}

func (it *numberEnumBase) MaxValueString() string {
	if it.maxStr != "" {
		return it.maxStr
	}

	it.maxStr = convAnyValToString(it.maxAny)

	return it.maxStr
}

func (it *numberEnumBase) IntegerEnumRanges() []int {
	return it.integerEnumRangesOnce.Values()
}

func (it numberEnumBase) Length() int {
	return len(it.StringRanges())
}

func (it numberEnumBase) Count() int {
	return len(it.StringRanges())
}

func (it *numberEnumBase) RangesDynamicMap() map[string]interface{} {
	if it.rangesDynamicMap != nil {
		return it.rangesDynamicMap
	}

	newMap := make(
		map[string]interface{},
		len(it.stringRanges)+1)

	for _, keyAnyVal := range it.KeyAnyValues() {
		newMap[keyAnyVal.Key] = keyAnyVal.AnyValue
	}

	it.rangesDynamicMap = newMap

	return newMap
}

func (it *numberEnumBase) DynamicMap() DynamicMap {
	return it.RangesDynamicMap()
}

func (it *numberEnumBase) notFoundJsonBytesError(
	currentValueInf interface{},
) error {
	compiledMessage := fmt.Sprintf(
		currentValueNotFoundInJsonMapFormat,
		currentValueInf,
		it.RangesInvalidMessage())

	return errors.New(compiledMessage)
}

func (it *numberEnumBase) RangesIntegerStringMap() map[int]string {
	if it.rangesDynamicMap != nil {
		return it.rangesIntegerStringMap
	}

	newMap := make(
		map[int]string,
		len(it.stringRanges)+1)

	for _, keyAnyVal := range it.KeyAnyValues() {
		newMap[keyAnyVal.ValInt()] = keyAnyVal.Key
	}

	it.rangesIntegerStringMap = newMap

	return newMap
}

func (it *numberEnumBase) KeyAnyValues() []KeyAnyVal {
	if it.keyAnyValues != nil {
		return it.keyAnyValues
	}

	it.keyAnyValues = KeyAnyValues(
		it.StringRanges(),
		it.actualValueRanges)

	return it.keyAnyValues
}

func (it numberEnumBase) KeyValIntegers() []KeyValInteger {
	slice := make([]KeyValInteger, it.Length())

	it.LoopInteger(func(index int, name string, valInteger int) (isBreak bool) {
		slice[index] = KeyValInteger{
			Key:          name,
			ValueInteger: valInteger,
		}

		return false
	})

	return slice
}

func (it numberEnumBase) Loop(looperFunc LooperFunc) {
	for i, keyAnyVal := range it.KeyAnyValues() {
		isBreak := looperFunc(i, keyAnyVal.Key, keyAnyVal.AnyValue)

		if isBreak {
			return
		}
	}
}

func (it numberEnumBase) LoopInteger(looperFunc LooperIntegerFunc) {
	for i, keyAnyVal := range it.KeyAnyValues() {
		isBreak := looperFunc(
			i,
			keyAnyVal.Key,
			keyAnyVal.ValInt())

		if isBreak {
			return
		}
	}
}

func (it numberEnumBase) TypeName() string {
	return it.typeName
}

// NameWithValueOption
//
// Warning :
//
// Make sure non ptr is called +
// String should also be attached with non ptr.
func (it numberEnumBase) NameWithValueOption(
	value interface{},
	isIncludeQuotation bool,
) string {
	if isIncludeQuotation {
		return fmt.Sprintf(
			constants.EnumDoubleQuoteNameValueFormat,
			value,
			value)
	}

	return NameWithValue(value)
}

// NameWithValue
//
// Warning :
//
// Make sure non ptr is called +
// String should also be attached with non ptr.
func (it numberEnumBase) NameWithValue(
	value interface{},
) string {
	return NameWithValue(value)
}

func (it numberEnumBase) ValueString(
	value interface{},
) string {
	return fmt.Sprintf(
		constants.SprintNumberFormat,
		value,
	)
}

// Format
//
//  Outputs name and
//  value by given format.
//
// sample-format :
//  - "Enum of {type-name} - {name} - {value}"
//
// sample-format-output :
//  - "Enum of EnumFullName - Invalid - 0"
//
// Key-Meaning :
//  - {type-name} : represents type-name string
//  - {name}      : represents name string
//  - {value}     : represents value string
func (it numberEnumBase) Format(
	format string,
	value interface{},
) string {
	return Format(
		it.TypeName(),
		it.ToName(value),
		it.ValueString(value),
		format,
	)
}

func (it *numberEnumBase) RangeNamesCsv() string {
	return it.rangesCsvString.Value()
}

func (it *numberEnumBase) RangesInvalidMessage() string {
	return it.rangesInvalidMessage.Value()
}

func (it *numberEnumBase) RangesInvalidErr() error {
	return it.invalidError.Value()
}

func (it numberEnumBase) StringRangesPtr() *[]string {
	return &it.stringRanges
}

func (it numberEnumBase) StringRanges() []string {
	return it.stringRanges
}

func (it numberEnumBase) NamesHashset() map[string]bool {
	if it.Length() == 0 {
		return map[string]bool{}
	}

	return toHashset(it.StringRanges()...)
}

func (it numberEnumBase) JsonString(input interface{}) string {
	return it.ToEnumString(input)
}

func (it numberEnumBase) ToEnumString(
	input interface{},
) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		input)
}

func (it numberEnumBase) ToName(
	input interface{},
) string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		input)
}
