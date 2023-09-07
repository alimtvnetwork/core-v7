package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

type BasicInt16 struct {
	numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int16 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int16][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int16]string // contains name without double quotes
	minVal, maxVal                           int16
}

func (it BasicInt16) IsAnyNamesOf(
	value int16,
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

func (it BasicInt16) IsAnyOf(value int16, checkingItems ...int16) bool {
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

func (it BasicInt16) Max() int16 {
	return it.maxVal
}

func (it BasicInt16) Min() int16 {
	return it.minVal
}

func (it BasicInt16) GetValueByString(valueString string) int16 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it BasicInt16) GetValueByName(
	name string,
) (int16, error) {
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

func (it BasicInt16) GetStringValue(input int16) string {
	return it.StringRanges()[input]
}

func (it BasicInt16) ExpectingEnumValueError(
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

func (it BasicInt16) Ranges() []int16 {
	return it.actualValueRanges.([]int16)
}

func (it BasicInt16) Hashmap() map[string]int16 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicInt16) HashmapPtr() *map[string]int16 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicInt16) IsValidRange(value int16) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it BasicInt16) ToEnumJsonBytes(value int16) ([]byte, error) {
	jsonBytes, has := it.valueToJsonDoubleQuoteStringBytesHashmap[value]

	if has {
		return jsonBytes, nil
	}

	return []byte{}, it.notFoundJsonBytesError(value)
}

func (it BasicInt16) ToEnumString(value int16) string {
	return it.valueNameHashmap[value]
}

func (it BasicInt16) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal int16,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it BasicInt16) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it BasicInt16) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it BasicInt16) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (int16, error) {
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

func (it BasicInt16) EnumType() enumtype.Variant {
	return enumtype.Integer16
}
