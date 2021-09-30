package enumimpl

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coreinterface"
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

func (it *BasicString) IsAnyOf(value string, checkingItems ...string) bool {
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

func (it *BasicString) Max() string {
	return it.maxVal
}

func (it *BasicString) Min() string {
	return it.minVal
}

func (it *BasicString) Ranges() []string {
	return it.actualValueRanges.([]string)
}

func (it *BasicString) Hashset() map[string]bool {
	return it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicString) HashsetPtr() *map[string]bool {
	return &it.jsonDoubleQuoteNameToValueHashMap
}

func (it *BasicString) IsValidRange(value string) bool {
	return it.jsonDoubleQuoteNameToValueHashMap[value]
}

func (it *BasicString) AppendPrependJoinValue(
	joiner string,
	appendVal, prependVal string,
) string {
	return it.ToEnumString(prependVal) +
		joiner +
		it.ToEnumString(appendVal)
}

func (it *BasicString) AppendPrependJoinNamer(
	joiner string,
	appendVal, prependVal coreinterface.ToNamer,
) string {
	return prependVal.Name() +
		joiner +
		appendVal.Name()
}

// ToEnumJsonBytes used for MarshalJSON from map
func (it *BasicString) ToEnumJsonBytes(value string) []byte {
	return it.valueToJsonDoubleQuoteStringBytesHashmap[value]
}

// UnmarshallToValue Mostly used for UnmarshalJSON
//
// Given bytes string enum value and transpile to exact enum raw value using map
func (it *BasicString) UnmarshallToValue(
	isMappedToFirstIfEmpty bool,
	jsonUnmarshallingValue []byte,
) (string, error) {
	if !isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return constants.EmptyString,
			defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	if isMappedToFirstIfEmpty && jsonUnmarshallingValue == nil {
		return it.minVal, nil
	}

	str := string(jsonUnmarshallingValue)
	if isMappedToFirstIfEmpty && (str == "" || str == `""`) {
		return it.minVal, nil
	}

	return str, nil
}
