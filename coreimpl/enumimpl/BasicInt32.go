package enumimpl

import "gitlab.com/evatix-go/core/converters"

type BasicInt32 struct {
	*numberEnumBase
	hashMap          map[string]int32
	jsonBytesHashmap map[int32][]byte
	minVal, maxVal   int32
}

func NewBasicInt32(
	actualValueRanges []int32,
	stringRanges []string,
	min, max int32,
) *BasicInt32 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	hashMap := make(map[string]int32, len(actualValueRanges))
	for i, actual := range actualValueRanges {
		key := stringRanges[i]
		hashMap[key] = actual
	}

	return &BasicInt32{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		hashMap:        hashMap,
	}
}

func (receiver *BasicInt32) Max() int32 {
	return receiver.maxVal
}

func (receiver *BasicInt32) Min() int32 {
	return receiver.minVal
}

func (receiver *BasicInt32) GetValueByString(valueString string) int32 {
	return receiver.hashMap[valueString]
}

func (receiver *BasicInt32) GetStringValue(input int32) string {
	return receiver.StringRanges()[input]
}

func (receiver *BasicInt32) Ranges() []int32 {
	return receiver.actualValueRanges.([]int32)
}

func (receiver *BasicInt32) Hashmap() map[string]int32 {
	return receiver.hashMap
}

func (receiver *BasicInt32) HashmapPtr() *map[string]int32 {
	return &receiver.hashMap
}

func (receiver *BasicInt32) IsValidRange(value int32) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}

func (receiver *BasicInt32) ToEnumJsonBytes(value int32) []byte {
	return receiver.jsonBytesHashmap[value]
}

func (receiver *BasicInt32) ToEnumString(value int32) string {
	return *converters.UnsafeBytesToStringPtr(receiver.jsonBytesHashmap[value])
}
