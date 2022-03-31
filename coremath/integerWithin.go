package coremath

import (
	"math"

	"gitlab.com/evatix-go/core/osconsts"
)

type integerWithin struct{}

func (it integerWithin) ToByte(value int) bool {
	return value >= 0 && value <= 255
}

func (it integerWithin) ToUnsignedInt16(value int) bool {
	return value >= 0 && value <= math.MaxUint16
}

func (it integerWithin) ToUnsignedInt32(value int) bool {
	if osconsts.IsX32Architecture {
		return it.isUint32On32BitArch(value)
	}

	return it.isUint32(value)
}

func (it integerWithin) isUint32(value int) bool {
	return value >= 0 && value <= math.MaxUint32
}

func (it integerWithin) isUint32On32BitArch(value int) bool {
	return value >= 0 && value <= math.MaxInt32
}

func (it integerWithin) ToUnsignedInt64(value int) bool {
	return value >= 0
}

func (it integerWithin) ToInt8(value int) bool {
	return value >= math.MinInt8 && value <= math.MaxInt8
}

func (it integerWithin) ToInt16(value int) bool {
	return value >= math.MinInt16 && value <= math.MaxInt16
}

func (it integerWithin) ToInt32(value int) bool {
	return value >= math.MinInt32 && value <= math.MaxInt32
}
