package enumimpl

type BasicInt16 struct {
	*numberEnumBase
	minVal, maxVal int16
}

func NewBasicInt16(
	actualValueRanges *[]int16,
	stringRanges *[]string,
	min, max int16,
) *BasicInt16 {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	return &BasicInt16{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
	}
}

func (receiver *BasicInt16) Max() int16 {
	return receiver.maxVal
}

func (receiver *BasicInt16) Min() int16 {
	return receiver.minVal
}

func (receiver *BasicInt16) Ranges() *[]int16 {
	return receiver.actualValueRanges.(*[]int16)
}

func (receiver *BasicInt16) IsValidRange(value int16) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}
