package corerange

import "gitlab.com/evatix-go/core/constants"

type MinMaxInt16 struct {
	Min, Max int16
}

func (receiver *MinMaxInt16) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(receiver.Min),
		Max: int(receiver.Min),
	}
}

func (receiver *MinMaxInt16) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt())
}

func (receiver *MinMaxInt16) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (receiver *MinMaxInt16) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		receiver.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (receiver *MinMaxInt16) Difference() int16 {
	return receiver.Max - receiver.Min
}

func (receiver *MinMaxInt16) DifferenceAbsolute() int16 {
	diff := receiver.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (receiver *MinMaxInt16) RangeLength() int16 {
	return receiver.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxInt16) RangesInt() *[]int16 {
	return receiver.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxInt16) Ranges() *[]int16 {
	length := receiver.RangeLength()
	start := receiver.Min
	slice := make([]int16, constants.Zero, length)
	var i int16

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (receiver *MinMaxInt16) IsWithinRange(value int16) bool {
	return receiver.Min >= value && value <= receiver.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (receiver *MinMaxInt16) IsInvalidValue(value int16) bool {
	return !receiver.IsWithinRange(value)
}
