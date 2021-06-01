package corerange

import "gitlab.com/evatix-go/core/constants"

type MinMaxInt struct {
	Min, Max int
}

func (receiver *MinMaxInt) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		receiver)
}

func (receiver *MinMaxInt) Difference() int {
	return receiver.Max - receiver.Min
}

func (receiver *MinMaxInt) DifferenceAbsolute() int {
	diff := receiver.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (receiver *MinMaxInt) RangeLength() int {
	return receiver.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxInt) RangesInt() *[]int {
	return receiver.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (receiver *MinMaxInt) Ranges() *[]int {
	length := receiver.RangeLength()
	start := receiver.Min
	slice := make([]int, constants.Zero, length)

	for i := 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (receiver *MinMaxInt) IsWithinRange(value int) bool {
	return receiver.Min >= value && value <= receiver.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (receiver *MinMaxInt) IsInvalidValue(value int) bool {
	return !receiver.IsWithinRange(value)
}
