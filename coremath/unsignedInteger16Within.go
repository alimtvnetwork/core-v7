package coremath

import (
	"math"
)

type unsignedInteger16Within struct{}

func (it unsignedInteger16Within) ToByte(value uint16) bool {
	return value <= 255
}

func (it unsignedInteger16Within) ToInt8(value uint16) bool {
	return value <= math.MaxInt8
}
