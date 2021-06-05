package enumimpl

import "gitlab.com/evatix-go/core/converters"

type BasicInt16 struct {
	*numberEnumBase
	hashMap          map[string]int16
	jsonBytesHashmap map[int16][]byte
	minVal, maxVal   int16
}

func NewBasicInt16(
	actualValueRanges []int16,
	stringRanges []string,
	min, max int16,
) *BasicInt16 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	hashMap := make(map[string]int16, len(actualValueRanges))
	jsonBytesHashmap := make(map[int16][]byte, len(actualValueRanges))

	for i, actualVal := range actualValueRanges {
		key := stringRanges[i]
		hashMap[key] = actualVal
		jsonBytesHashmap[actualVal] = []byte(key)
	}

	return &BasicInt16{
		numberEnumBase:   enumBase,
		minVal:           min,
		maxVal:           max,
		hashMap:          hashMap,
		jsonBytesHashmap: jsonBytesHashmap,
	}
}

func (receiver *BasicInt16) Max() int16 {
	return receiver.maxVal
}

func (receiver *BasicInt16) Min() int16 {
	return receiver.minVal
}

func (receiver *BasicInt16) GetValueByString(valueString string) int16 {
	return receiver.hashMap[valueString]
}

func (receiver *BasicInt16) GetStringValue(input int16) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicInt16) Ranges() []int16 {
	return receiver.actualValueRanges.([]int16)
}

func (receiver *BasicInt16) Hashmap() map[string]int16 {
	return receiver.hashMap
}

func (receiver *BasicInt16) HashmapPtr() *map[string]int16 {
	return &receiver.hashMap
}

func (receiver *BasicInt16) IsValidRange(value int16) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

func (receiver *BasicInt16) ToEnumJsonBytes(value int16) []byte {
	return receiver.jsonBytesHashmap[value]
}

func (receiver *BasicInt16) ToEnumString(value int16) string {
	return *converters.UnsafeBytesToStringPtr(receiver.jsonBytesHashmap[value])
}
