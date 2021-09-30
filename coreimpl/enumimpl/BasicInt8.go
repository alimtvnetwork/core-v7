package enumimpl

import (
	"fmt"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicInt8 struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int8 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int8][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int8]string // contains name without double quotes
	minVal, maxVal                           int8
}

func NewBasicInt8(
	typeName string,
	actualValueRanges []int8,
	stringRanges []string,
	min, max int8,
) *BasicInt8 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int8, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int8][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int8]string, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		indexJson := simplewrap.WithDoubleQuoteAny(i)
		indexString := strconv.Itoa(i)
		jsonName := simplewrap.WithDoubleQuote(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		jsonDoubleQuoteNameToValueHashMap[indexJson] = actualVal
		jsonDoubleQuoteNameToValueHashMap[indexString] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[actualVal] = []byte(jsonName)
		valueNameHashmap[actualVal] = key
	}

	return &BasicInt8{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func NewBasicInt8UsingIndexedSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicInt8 {
	min := constants.Zero
	max := len(indexedSliceWithValues)

	actualValues := make([]int8, max)
	for i := range indexedSliceWithValues {
		actualValues[i] = int8(i)
	}

	return NewBasicInt8(
		typeName,
		actualValues,
		indexedSliceWithValues,
		int8(min),
		int8(max))
}

func (it *BasicInt8) IsAnyOf(value int8, checkingItems ...int8) bool {
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

func (it *BasicInt8) Max() int8 {
	return it.maxVal
}

func (it *BasicInt8) Min() int8 {
	return it.minVal
}

func (it *BasicInt8) GetValueByString(valueString string) int8 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it *BasicInt8) GetStringValue(input int8) string {
	return it.StringRanges()[input]
}

func (it *BasicInt8) Ranges() []int8 {
	return it.actualValueRanges.([]int8)
}

func (it *BasicInt8) Hashmap() map[string]int8 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicInt8) HashmapPtr() *map[string]int8 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicInt8) IsValidRange(value int8) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it *BasicInt8) ToEnumJsonBytes(value int8) []byte {
	return it.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (it *BasicInt8) ToEnumString(value int8) string {
	return it.valueNameHashmap[value]
}

func (it *BasicInt8) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal int8,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it *BasicInt8) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it *BasicInt8) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it *BasicInt8) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (int8, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.Zero,
			defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return it.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty &&
		(str == constants.EmptyString || str == constants.DoubleQuotationStartEnd) {
		return it.minVal, nil
	}

	v, has := it.jsonDoubleQuoteNameToValueHashMap[str]

	if !has {
		return constants.Zero, enumUnmarshallingMappingFailedError(
			it.TypeName(),
			str,
			it.RangeNamesCsv())
	}

	return v, nil
}
