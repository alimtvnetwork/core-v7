package enumimpl

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/simplewrap"
)

type BasicString struct {
	*numberEnumBase
	jsonDoubleQuoteNameToValueHashMap        map[string]bool   // contains names double quotes to value
	valueToJsonDoubleQuoteStringBytesHashmap map[string][]byte // contains value to string bytes with double quotes
	minVal, maxVal                           string
}

func NewBasicString(
	typeName string,
	stringRanges []string,
	min, max string,
) *BasicString {
	enumBase := newNumberEnumBase(
		typeName,
		stringRanges,
		stringRanges,
		min,
		max)

	jsonDoubleQuoteNameToValueHashMap := make(
		map[string]string,
		len(stringRanges))
	valueToJsonDoubleQuoteStringBytesHashmap := make(
		map[string][]byte,
		len(stringRanges))

	for i, actualVal := range stringRanges {
		key := stringRanges[i]
		jsonName := simplewrap.WithDoubleQuote(key)
		jsonDoubleQuoteNameToValueHashMap[jsonName] = actualVal
		valueToJsonDoubleQuoteStringBytesHashmap[key] = []byte(jsonName)
	}

	return &BasicString{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		jsonDoubleQuoteNameToValueHashMap: *converters.
			StringsToMap(&stringRanges),
		valueToJsonDoubleQuoteStringBytesHashmap: valueToJsonDoubleQuoteStringBytesHashmap,
	}
}

func (receiver *BasicString) IsAnyOf(value string, checkingItems ...string) bool {
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

func (receiver *BasicString) Max() string {
	return receiver.maxVal
}

func (receiver *BasicString) Min() string {
	return receiver.minVal
}

func (receiver *BasicString) Ranges() []string {
	return receiver.actualValueRanges.([]string)
}

func (receiver *BasicString) Hashset() map[string]bool {
	return receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicString) HashsetPtr() *map[string]bool {
	return &receiver.jsonDoubleQuoteNameToValueHashMap
}

func (receiver *BasicString) IsValidRange(value string) bool {
	return receiver.jsonDoubleQuoteNameToValueHashMap[value]
}

// ToEnumJsonBytes used for MarshalJSON from map
func (receiver *BasicString) ToEnumJsonBytes(value string) []byte {
	return receiver.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (receiver *BasicString) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (string, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.EmptyString,
			defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return receiver.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty && (str == "" || str == `""`) {
		return receiver.minVal, nil
	}

	return str, nil
}
