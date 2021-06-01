package enumimpl

type BasicInt32 struct {
	*numberEnumBase
	minVal, maxVal int32
}

func NewBasicInt32(
	actualValueRanges *[]int32,
	stringRanges *[]string,
	min, max int32,
) *BasicInt32 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	return &BasicInt32{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
	}
}

func (receiver *BasicInt32) Max() int32 {
	return receiver.maxVal
}

func (receiver *BasicInt32) Min() int32 {
	return receiver.minVal
}

func (receiver *BasicInt32) Ranges() *[]int32 {
	return receiver.actualValueRanges.(*[]int32)
}

func (receiver *BasicInt32) IsValidRange(value int32) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}
