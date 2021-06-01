package corerange

import "gitlab.com/evatix-go/core/constants"

type MinMaxInt8 struct {
	Min, Max int8
}

func (receiver *MinMaxInt8) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(receiver.Min),
		Max: int(receiver.Min),
	}
}

func (receiver *MinMaxInt8) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt())
}

func (receiver *MinMaxInt8) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (receiver *MinMaxInt8) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (receiver *MinMaxInt8) Difference() int8 {
	return receiver.Max - receiver.Min
}

func (receiver *MinMaxInt8) DifferenceAbsolute() int8 {
	diff := receiver.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (receiver *MinMaxInt8) RangeLength() int8 {
	return receiver.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxInt8) RangesInt() *[]int8 {
	return receiver.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxInt8) Ranges() *[]int8 {
	length := receiver.RangeLength()
	start := receiver.Min
	slice := make([]int8, constants.Zero, length)
	var i int8

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (receiver *MinMaxInt8) IsWithinRange(value int8) bool {
	return receiver.Min >= value && value <= receiver.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (receiver *MinMaxInt8) IsInvalidValue(value int8) bool {
	return !receiver.IsWithinRange(value)
}
