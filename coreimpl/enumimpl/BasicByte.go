package enumimpl

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

type BasicByte struct {
	*numberEnumBase
	hashMap          map[string]byte
	jsonBytesHashmap map[byte][]byte
	minVal, maxVal   byte
}

func NewBasicByte(
	actualValueRanges []byte,
	stringRanges []string,
	min, max byte,
) *BasicByte {
	enumBase := newNumberEnumBase(
		&actualValueRanges,
		stringRanges,
		min,
		max)

	hashMap := make(map[string]byte, len(actualValueRanges))
	jsonBytesHashmap := make(map[byte][]byte, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		hashMap[key] = actualVal
		jsonBytesHashmap[actualVal] = []byte(key)
	}

	return &BasicByte{
		numberEnumBase:   enumBase,
		minVal:           min,
		maxVal:           max,
		hashMap:          hashMap,
		jsonBytesHashmap: jsonBytesHashmap,
	}
}

func NewBasicByteUsingIndexedSlice(
	indexedSliceWithValues []string,
) *BasicByte {
	min := constants.Zero
	max := len(indexedSliceWithValues)

	actualValues := make([]byte, max)
	for i := range indexedSliceWithValues {
		actualValues[i] = byte(i)
	}

	return NewBasicByte(
		actualValues,
		indexedSliceWithValues,
		byte(min),
		byte(max))
}

func (receiver *BasicByte) Max() byte {
	return receiver.maxVal
}

func (receiver *BasicByte) Min() byte {
	return receiver.minVal
}

func (receiver *BasicByte) GetValueByString(valueString string) byte {
	return receiver.hashMap[valueString]
}

func (receiver *BasicByte) GetStringValue(input byte) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicByte) Ranges() []byte {
	return receiver.actualValueRanges.([]byte)
}

func (receiver *BasicByte) Hashmap() map[string]byte {
	return receiver.hashMap
}

func (receiver *BasicByte) HashmapPtr() *map[string]byte {
	return &receiver.hashMap
}

func (receiver *BasicByte) IsValidRange(value byte) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

func (receiver *BasicByte) ToEnumJsonBytes(value byte) []byte {
	return receiver.jsonBytesHashmap[value]
}

func (receiver *BasicByte) ToEnumString(value byte) string {
	return *converters.UnsafeBytesToStringPtr(receiver.jsonBytesHashmap[value])
}
