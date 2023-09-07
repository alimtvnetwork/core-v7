package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreimpl/enumimpl/enumtype"
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/core/errcore"
)

type BasicInt32 struct {
	numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int32 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int32][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int32]string // contains name without double quotes
	minVal, maxVal                           int32
}

func (it BasicInt32) IsAnyNamesOf(
	value int32,
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

func (it BasicInt32) IsAnyOf(value int32, checkingItems ...int32) bool {
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

func (it BasicInt32) Max() int32 {
	return it.maxVal
}

func (it BasicInt32) Min() int32 {
	return it.minVal
}

func (it BasicInt32) GetValueByString(valueString string) int32 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it BasicInt32) GetValueByName(name string) (int32, error) {
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

func (it BasicInt32) GetStringValue(input int32) string {
	return it.StringRanges()[input]
}

func (it BasicInt32) ExpectingEnumValueError(
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

func (it BasicInt32) Ranges() []int32 {
	return it.actualValueRanges.([]int32)
}

func (it BasicInt32) Hashmap() map[string]int32 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicInt32) HashmapPtr() *map[string]int32 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it BasicInt32) IsValidRange(value int32) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it BasicInt32) ToEnumJsonBytes(value int32) ([]byte, error) {
	jsonBytes, has := it.valueToJsonDoubleQuoteStringBytesHashmap[value]

	if has {
		return jsonBytes, nil
	}

	return []byte{}, it.notFoundJsonBytesError(value)
}

func (it BasicInt32) ToEnumString(value int32) string {
	return it.valueNameHashmap[value]
}

func (it BasicInt32) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal int32,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it BasicInt32) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()

}

func (it BasicInt32) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it BasicInt32) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (int32, error) {
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

func (it BasicInt32) EnumType() enumtype.Variant {
	return enumtype.Integer32
}
