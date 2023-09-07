package coremath

import (
	"math"
)

type integer32Within struct{}

func (it integer32Within) ToByte(value int32) bool {
	return value >= 0 && value <= 255
}

func (it integer32Within) ToUnsignedInt16(value int32) bool {
	return value >= 0 && value <= math.MaxUint16
}

func (it integer32Within) ToUnsignedInt32(value int32) bool {
	return value >= 0
}

func (it integer32Within) ToUnsignedInt64(value int32) bool {
	return value >= 0
}

func (it integer32Within) ToInt8(value int32) bool {
	return value >= math.MinInt8 && value <= math.MaxInt8
}

func (it integer32Within) ToInt16(value int32) bool {
	return value >= math.MinInt16 && value <= math.MaxInt16
}

func (it integer32Within) ToInt(value int) bool {
	// it would be different based on architecture.
	return value >= math.MinInt32 && value <= math.MaxInt32
}
