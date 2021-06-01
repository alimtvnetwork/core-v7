package enumimpl

type BasicByte struct {
	*numberEnumBase
	minVal, maxVal byte
}

func NewBasicByte(
	actualValueRanges *[]byte,
	stringRanges *[]string,
	min, max byte,
) *BasicByte {
	enumBase := newNumberEnumBase(
		actualValueRanges,
		stringRanges,
		min,
		max)

	return &BasicByte{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
	}
}
func (receiver *BasicByte) Max() byte {
	return receiver.maxVal
}

func (receiver *BasicByte) Min() byte {
	return receiver.minVal
}

func (receiver *BasicByte) Ranges() *[]byte {
	return receiver.actualValueRanges.(*[]byte)
}

func (receiver *BasicByte) IsValidRange(value byte) bool {
	return value >= receiver.minVal && value <= receiver.maxVal
}
