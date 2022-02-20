package coremath

import (
	"math"

	"gitlab.com/evatix-go/core/constants"
)

type integer32Within struct{}

func (it integer32Within) ToByte(value int32) bool {
	return value >= 0 && value <= 255
}

func (it integer32Within) ToUnsignedInt16(value int32) bool {
	return value >= 0 && value <= math.MaxUint16
}

func (it integer32Within) ToUnsignedInt32(value int32) bool {
	return value >= 0 && value <= math.MaxInt32
}

func (it integer32Within) ToUnsignedInt64(value int32) bool {
	return value >= 0 && value <= math.MaxUint64
}

func (it integer32Within) ToInt8(value int32) bool {
	return value >= math.MinInt8 && value <= math.MaxInt8
}

func (it integer32Within) ToInt16(value int32) bool {
	return value >= math.MinInt16 && value <= math.MaxInt16
}

func (it integer32Within) ToInt32(value int32) bool {
	return value >= math.MinInt32 && value <= math.MaxInt32
}

func (it integer32Within) ToInt(value int32) bool {
	return value >= int32(constants.MinInt) && value <= int32(constants.MaxInt)
}
