package enumimpl

import (
	"fmt"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicInt16 struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int16 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int16][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int16]string // contains name without double quotes
	minVal, maxVal                           int16
}

func NewBasicInt16(
	typeName string,
	actualValueRanges []int16,
	stringRanges []string,
	min, max int16,
) *BasicInt16 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int16, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int16][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int16]string, len(actualValueRanges))

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

	return &BasicInt16{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func NewBasicInt16UsingIndexedSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicInt16 {
	min := constants.Zero
	max := len(indexedSliceWithValues)

	actualValues := make([]int16, max)
	for i := range indexedSliceWithValues {
		actualValues[i] = int16(i)
	}

	return NewBasicInt16(
		typeName,
		actualValues,
		indexedSliceWithValues,
		int16(min),
		int16(max))
}

func (it *BasicInt16) IsAnyOf(value int16, checkingItems ...int16) bool {
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

func (it *BasicInt16) Max() int16 {
	return it.maxVal
}

func (it *BasicInt16) Min() int16 {
	return it.minVal
}

func (it *BasicInt16) GetValueByString(valueString string) int16 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it *BasicInt16) GetStringValue(input int16) string {
	return it.StringRanges()[input]
}

func (it *BasicInt16) Ranges() []int16 {
	return it.actualValueRanges.([]int16)
}

func (it *BasicInt16) Hashmap() map[string]int16 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicInt16) HashmapPtr() *map[string]int16 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicInt16) IsValidRange(value int16) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it *BasicInt16) ToEnumJsonBytes(value int16) []byte {
	return it.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (it *BasicInt16) ToEnumString(value int16) string {
	return it.valueNameHashmap[value]
}

func (it *BasicInt16) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal int16,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it *BasicInt16) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it *BasicInt16) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it *BasicInt16) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (int16, error) {
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
