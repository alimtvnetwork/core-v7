package coremath

import (
	"math"

	"gitlab.com/evatix-go/core/constants"
)

type integer64OutOfRange struct{}

func (it integer64OutOfRange) Byte(value int64) bool {
	return !(value >= 0 && value <= 255)
}

func (it integer64OutOfRange) UnsignedInt16(value int64) bool {
	return !(value >= 0 && value <= int64(math.MaxUint16))
}

func (it integer64OutOfRange) UnsignedInt32(value int64) bool {
	return !(value >= 0 && value <= int64(math.MaxUint32))
}

func (it integer64OutOfRange) UnsignedInt64(value int64) bool {
	return !(value >= 0 && value <= int64(math.MaxUint64))
}

func (it integer64OutOfRange) Int8(value int64) bool {
	return !(value >= int64(math.MinInt8) && value <= int64(math.MaxInt8))
}

func (it integer64OutOfRange) Int16(value int64) bool {
	return !(value >= int64(math.MinInt16) && value <= int64(math.MaxInt16))
}

func (it integer64OutOfRange) Int32(value int64) bool {
	return !(value >= int64(math.MinInt32) && value <= int64(math.MaxInt32))
}

func (it integer64OutOfRange) Int(value int64) bool {
	return !(value >= int64(constants.MinInt) && value <= int64(constants.MaxInt))
}
