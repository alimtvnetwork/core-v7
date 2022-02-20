package coremath

import (
	"math"
)

type integer16Within struct{}

func (it integer16Within) ToByte(value int16) bool {
	return value >= 0 && value <= 255
}

func (it integer16Within) ToUnsignedInt16(value int16) bool {
	return value >= 0
}

func (it integer16Within) ToUnsignedInt32(value int16) bool {
	return value >= 0
}

func (it integer16Within) ToUnsignedInt64(value int16) bool {
	return value >= 0 && value <= math.MaxUint64
}

func (it integer16Within) ToInt8(value int16) bool {
	return value >= math.MinInt8 && value <= math.MaxInt8
}
