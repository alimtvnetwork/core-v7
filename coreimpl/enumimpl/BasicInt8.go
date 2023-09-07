package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

type BasicInt8 struct {
	numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int8 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int8][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int8]string // contains name without double quotes
	minVal, maxVal                           int8
}

func (it BasicInt8) IsAnyNamesOf(
	value int8,
	names ...string,
) bool {
	currentName := it.ToEnumString(value)

	for _, name := range names {
		if name == currentName {
			return true
		}
	}

	return false
}

func (it BasicInt8) IsAnyOf(
	value int8,
	checkingItems ...int8,
) bool {
	if len(checkingItems) == 0 {
		return true
	}

	for _, givenByte := range checkingItems {
		if value == givenByte {
			return true
		}
	}

	return false
}

func (it BasicInt8) Max() int8 {
	return it.maxVal
}

func (it BasicInt8) Min() int8 {
	return it.minVal
}

func (it BasicInt8) GetValueByString(valueString string) int8 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it BasicInt8) GetValueByName(name string) (int8, error) {
	v, has := it.jsonDoubleQuoteNameToValueHashMap[name]

	if has {
		return v, nil
	}

	wrapped := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name)

	nextVal, isFoundByWrapped := it.jsonDoubleQuoteNameToValueHashMap[wrapped]

	if isFoundByWrapped {
		return nextVal, nil
	}

	// has error
	return constants.InvalidValue, enumUnmarshallingMappingFailedError(
		it.TypeName(),
		name,
		it.RangeNamesCsv())
}

func (it BasicInt8) ExpectingEnumValueError(
	rawString string,
	expectedEnum interface{},
) error {
	expectedEnumName := it.ToName(expectedEnum)
	expectedValue := it.GetValueByString(expectedEnumName)
	convValue, err := it.GetValueByName(rawString)

	if err != nil {
		return errcore.ExpectingErrorSimpleNoType(
			"Expecting enum: "+expectedEnumName,
			expectedEnumName,
			rawString+err.Error())
	}

	if convValue == expectedValue {
		return nil
	}

	return errcore.ExpectingErrorSimpleNoType(
		"Expecting enum: "+expectedEnumName,
		expectedEnumName,
		rawString+it.RangesInvalidMessage())
}

func (it BasicInt8) GetStringValue(input int8) string {
	return it.StringRanges()[input]
}

func (it BasicInt8) Ranges() []int8 {
	return it.actualValueRanges.([]int8)
}

func (it BasicInt8) Hashmap() map[string]int8 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicInt8) HashmapPtr() *map[string]int8 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicInt8) IsValidRange(value int8) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes
//
//  used for MarshalJSON from map
func (it BasicInt8) ToEnumJsonBytes(value int8) ([]byte, error) {
	jsonBytes, has := it.valueToJsonDoubleQuoteStringBytesHashmap[value]

	if has {
		return jsonBytes, nil
	}

	return []byte{}, it.notFoundJsonBytesError(value)
}

func (it BasicInt8) ToEnumString(value int8) string {
	return it.valueNameHashmap[value]
}

func (it BasicInt8) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal int8,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it BasicInt8) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it BasicInt8) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it BasicInt8) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (int8, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.Zero,
			defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return it.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty &&
		(str == constants.EmptyString || str == constants.DoubleQuotationStartEnd) {
		return it.minVal, nil
	}

	return it.GetValueByName(str)
}

func (it BasicInt8) EnumType() enumtype.Variant {
	return enumtype.Integer8
}
