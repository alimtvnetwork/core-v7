package enumimpl

import (
	"fmt"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicByte struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]byte // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[byte][]byte // contains value to string bytes with double quotes
	valueNameHashmap                         map[byte]string // contains name without double quotes
	minVal, maxVal                           byte
}

func NewBasicByte(
	typeName string,
	actualValueRanges []byte,
	stringRanges []string,
	min, max byte,
) *BasicByte {
	enumBase := newNumberEnumBase(
		typeName,
		&actualValueRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(map[string]byte, len(actualValueRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(map[byte][]byte, len(actualValueRanges))
	valueNameHashmap := make(map[byte]string, len(actualValueRanges))

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

	return &BasicByte{
		numberEnumBase:                           enumBase,
		minVal:                                   min,
		maxVal:                                   max,
		jsonDoubleQuoteNameToValueHashMap:        jsonDoubleQuoteNameToValueHashMap,
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
		valueNameHashmap:                         valueNameHashmap,
	}
}

func NewBasicByteUsingIndexedSlice(
	typeName string,
	indexedSliceWithValues []string,
) *BasicByte {
	min := constants.Zero
	max := len(indexedSliceWithValues)

	actualValues := make([]byte, max)
	for i := range indexedSliceWithValues {
		actualValues[i] = byte(i)
	}

	return NewBasicByte(
		typeName,
		actualValues,
		indexedSliceWithValues,
		byte(min),
		byte(max))
}

func (it *BasicByte) IsAnyOf(value byte, givenBytes ...byte) bool {
	if len(givenBytes) == 0 {
		return true
	}

	for _, givenByte := range givenBytes {
		if value == givenByte {
			return true
		}
	}

	return false
}

func (it *BasicByte) Max() byte {
	return it.maxVal
}

func (it *BasicByte) Min() byte {
	return it.minVal
}

func (it *BasicByte) GetValueByString(valueString string) byte {
	return it.jsonDoubleQuoteNameToValueHashMap[valueString]
}

func (it *BasicByte) GetStringValue(input byte) string {
	return it.StringRanges()[input]
}

func (it *BasicByte) Ranges() []byte {
	return it.actualValueRanges.([]byte)
}

func (it *BasicByte) Hashmap() map[string]byte {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicByte) HashmapPtr() *map[string]byte {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicByte) IsValidRange(value byte) bool {
	return value >= it.minVal && value <= it.maxVal
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it *BasicByte) ToEnumJsonBytes(value byte) []byte {
	return it.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

func (it *BasicByte) ToEnumString(value byte) string {
	return it.valueNameHashmap[value]
}

func (it *BasicByte) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal byte,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it *BasicByte) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

func (it *BasicByte) ToNumberString(valueInRawFormat interface{}) string {
	return fmt.Sprintf(constants.SprintValueFormat, valueInRawFormat)
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it *BasicByte) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (byte, error) {
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
