package coremath

import (
	"math"

	"gitlab.com/evatix-go/core/constants"
)

type integer64Within struct{}

func (it integer64Within) ToByte(value int64) bool {
	return value >= 0 && value <= 255
}

func (it integer64Within) ToUnsignedInt16(value int64) bool {
	return value >= 0 && value <= math.MaxUint16
}

func (it integer64Within) ToUnsignedInt32(value int64) bool {
	return value >= 0 && value <= math.MaxUint32
}

func (it integer64Within) ToUnsignedInt64(value int64) bool {
	return value >= 0 && value <= math.MaxUint64
}

func (it integer64Within) ToInt8(value int64) bool {
	return value >= math.MinInt8 && value <= math.MaxInt8
}

func (it integer64Within) ToInt16(value int64) bool {
	return value >= math.MinInt16 && value <= math.MaxInt16
}

func (it integer64Within) ToInt32(value int64) bool {
	return value >= math.MinInt32 && value <= math.MaxInt32
}

func (it integer64Within) ToInt(value int64) bool {
	return value >= int64(constants.MinInt) && value <= int64(constants.MaxInt)
}
