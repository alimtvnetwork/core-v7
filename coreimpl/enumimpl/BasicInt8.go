package enumimpl

import "gitlab.com/evatix-go/core/converters"

type BasicInt8 struct {
	*numberEnumBase
	hashMap          map[string]int8
	jsonBytesHashmap map[int8][]byte
	minVal, maxVal   int8
}

func NewBasicInt8(
	actualValueRanges []int8,
	stringRanges []string,
	min, max int8,
) *BasicInt8 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	hashMap := make(map[string]int8, len(actualValueRanges))
	jsonBytesHashmap := make(map[int8][]byte, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		hashMap[key] = actualVal
		jsonBytesHashmap[actualVal] = []byte(key)
	}

	return &BasicInt8{
		numberEnumBase:   enumBase,
		minVal:           min,
		maxVal:           max,
		hashMap:          hashMap,
		jsonBytesHashmap: jsonBytesHashmap,
	}
}

func (receiver *BasicInt8) Max() int8 {
	return receiver.maxVal
}

func (receiver *BasicInt8) Min() int8 {
	return receiver.minVal
}

func (receiver *BasicInt8) GetValueByString(valueString string) int8 {
	return receiver.hashMap[valueString]
}

func (receiver *BasicInt8) GetStringValue(input int8) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicInt8) Ranges() []int8 {
	return receiver.actualValueRanges.([]int8)
}

func (receiver *BasicInt8) Hashmap() map[string]int8 {
	return receiver.hashMap
}

func (receiver *BasicInt8) HashmapPtr() *map[string]int8 {
	return &receiver.hashMap
}

func (receiver *BasicInt8) IsValidRange(value int8) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

func (receiver *BasicInt8) ToEnumJsonBytes(value int8) []byte {
	return receiver.jsonBytesHashmap[value]
}

func (receiver *BasicInt8) ToEnumString(value int8) string {
	return *converters.UnsafeBytesToStringPtr(receiver.jsonBytesHashmap[value])
}
