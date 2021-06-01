package corerange

import "gitlab.com/evatix-go/core/constants"

type MinMaxByte struct {
	Min, Max byte
}

func (receiver *MinMaxByte) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(receiver.Min),
		Max: int(receiver.Min),
	}
}

func (receiver *MinMaxByte) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt())
}

func (receiver *MinMaxByte) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (receiver *MinMaxByte) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (receiver *MinMaxByte) Difference() byte {
	return receiver.Max - receiver.Min
}

func (receiver *MinMaxByte) DifferenceAbsolute() byte {
	diff := receiver.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (receiver *MinMaxByte) RangeLength() byte {
	return receiver.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxByte) RangesInt() *[]byte {
	return receiver.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxByte) Ranges() *[]byte {
	length := receiver.RangeLength()
	start := receiver.Min
	slice := make([]byte, constants.Zero, length)
	var i byte

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (receiver *MinMaxByte) IsWithinRange(value byte) bool {
	return receiver.Min >= value && value <= receiver.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (receiver *MinMaxByte) IsInvalidValue(value byte) bool {
	return !receiver.IsWithinRange(value)
}
