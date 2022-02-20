package coremath

import (
	"math"

	"gitlab.com/evatix-go/core/constants"
)

type integerOutOfRange struct{}

func (it integerOutOfRange) ToByte(value int) bool {
	return !(value >= 0 && value <= 255)
}

func (it integerOutOfRange) ToUnsignedInt16(value int) bool {
	return !(value >= 0 && value <= math.MaxUint16)
}

func (it integerOutOfRange) ToUnsignedInt32(value int) bool {
	return !(value >= 0 && value <= math.MaxUint32)
}

func (it integerOutOfRange) ToUnsignedInt64(value int) bool {
	return !(value >= 0 && value <= math.MaxUint64)
}

func (it integerOutOfRange) ToInt8(value int) bool {
	return !(value >= math.MinInt8 && value <= math.MaxInt8)
}

func (it integerOutOfRange) ToInt16(value int) bool {
	return !(value >= math.MinInt16 && value <= math.MaxInt16)
}

func (it integerOutOfRange) ToInt32(value int) bool {
	return !(value >= math.MinInt32 && value <= math.MaxInt32)
}

func (it integerOutOfRange) ToInt(value int) bool {
	return !(value >= constants.MinInt && value <= constants.MaxInt)
}
