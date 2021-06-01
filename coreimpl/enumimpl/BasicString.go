package enumimpl

import "gitlab.com/evatix-go/core/converters"

type BasicString struct {
	*numberEnumBase
	hashset        map[string]bool
	minVal, maxVal string
}

func NewBasicString(
	stringRanges *[]string,
	min, max string,
) *BasicString {
	enumBase := newNumberEnumBase(
		stringRanges,
		stringRanges,
		min,
		max)

	return &BasicString{
		numberEnumBase: enumBase,
		minVal:         min,
		maxVal:         max,
		hashset: *converters.
			StringsToMap(stringRanges),
	}
}

func (receiver *BasicString) Max() string {
	return receiver.maxVal
}

func (receiver *BasicString) Min() string {
	return receiver.minVal
}

func (receiver *BasicString) Ranges() *[]string {
	return receiver.actualValueRanges.(*[]string)
}

func (receiver *BasicString) IsValidRange(value string) bool {
	return receiver.hashset[value]
}
