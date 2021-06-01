package enumimpl

type BasicInt8 struct {
	*numberEnumBase
	minVal, maxVal int8
}

func NewBasicInt8(
	actualValueRanges *[]int8,
	stringRanges *[]string,
	min, max int8,
) *BasicInt8 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	return &BasicInt8{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
	}
}

func (receiver *BasicInt8) Max() int8 {
	return receiver.maxVal
}

func (receiver *BasicInt8) Min() int8 {
	return receiver.minVal
}

func (receiver *BasicInt8) Ranges() *[]int8 {
	return receiver.actualValueRanges.(*[]int8)
}

func (receiver *BasicInt8) IsValidRange(value int8) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}
