package enumimpl

import "gitlab.com/evatix-go/core/converters"

type BasicString struct {
	*numberEnumBase
	hashset          map[string]bool
	jsonBytesHashmap map[string][]byte
	minVal, maxVal   string
}

func NewBasicString(
	stringRanges []string,
	min, max string,
) *BasicString {
	enumBase := newNumberEnumBase(
		stringRanges,
		stringRanges,
		min,
		max)

	jsonBytesHashmap := make(
		map[string][]byte,
		len(stringRanges))

	for _, actualVal := range stringRanges {
		jsonBytesHashmap[actualVal] = []byte(actualVal)
	}

	return &BasicString{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		hashset: *converters.
			StringsToMap(&stringRanges),
		jsonBytesHashmap: jsonBytesHashmap,
	}
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
	return receiver.hashset
}

func (receiver *BasicString) HashsetPtr() *map[string]bool {
	return &receiver.hashset
}

func (receiver *BasicString) IsValidRange(value string) bool {
	return receiver.hashset[value]
}

func (receiver *BasicString) ToEnumJsonBytes(value string) []byte {
	return receiver.jsonBytesHashmap[value]
}
