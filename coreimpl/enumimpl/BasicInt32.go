package enumimpl

import (
	"fmt"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicInt32 struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]int32 // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[int32][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[int32]string // contains name without double quotes
	minVal, maxVal                           int32
}

func NewBasicInt32(
	typeName string,
	actualValueRanges []int32,
	stringRanges []string,
	min, max int32,
) *BasicInt32 {
	enumBase := newNumberEnumBase(
		typeName,
		actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]int32, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[int32][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[int32]string, len(actualValueRanges))

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

	return &BasicInt32{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func (it *BasicInt32) IsAnyOf(value int32, checkingItems ...int32) bool {
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

func (it *BasicInt32) Max() int32 {
	return it.maxVal
}

func (it *BasicInt32) Min() int32 {
	return it.minVal
}

func (it *BasicInt32) GetValueByString(valueString string) int32 {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it *BasicInt32) GetStringValue(input int32) string {
	return it.StringRanges()[input]
}

func (it *BasicInt32) Ranges() []int32 {
	return it.actualValueRanges.([]int32)
}

func (it *BasicInt32) Hashmap() map[string]int32 {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicInt32) HashmapPtr() *map[string]int32 {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicInt32) IsValidRange(value int32) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it *BasicInt32) ToEnumJsonBytes(value int32) []byte {
	return it.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (it *BasicInt32) ToEnumString(value int32) string {
	return it.valueNameHashmap[value]
}

func (it *BasicInt32) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal int32,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it *BasicInt32) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()

}

func (it *BasicInt32) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it *BasicInt32) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (int32, error) {
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
