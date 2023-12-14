package corerange

import (
	"gitlab.com/auk-go/core/constants"
)

type RangeByte struct {
	*BaseRange
	Start, End byte
}

func NewRangeByteMinMax(
	rawString, separator string,
	min, max byte,
) *RangeByte {
	minMax := MinMaxInt{
		Min: int(min),
		Max: int(max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMax).
		CreateRangeByte()
}

func NewRangeByte(
	rawString, separator string,
	minMax *MinMaxByte,
) *RangeByte {
	if minMax == nil {
		minMaxInt := MinMaxInt{
			Min: constants.Zero,
			Max: constants.MaxUnit8AsInt,
		}

		return NewRangeInt(
			rawString,
			separator,
			&minMaxInt).
			CreateRangeByte()
	}

	minMaxInt := MinMaxInt{
		Min: int(minMax.Min),
		Max: int(minMax.Max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMaxInt).
		CreateRangeByte()
}

// Difference
//
// Checks comparison wise which one is bigger than does the diff.
func (it *RangeByte) Difference() byte {
	if it.End > it.Start {
		return it.End - it.Start
	}

	return it.Start - it.End
}

func (it *RangeByte) DifferenceAbsolute() byte {
	return it.Difference()
}

// RangeLength (5 - 3 = 2) + 1
func (it *RangeByte) RangeLength() byte {
	return it.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (it *RangeByte) RangesInt() *[]byte {
	return it.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *RangeByte) Ranges() *[]byte {
	if it.IsInvalid() {
		return &[]byte{}
	}

	length := it.RangeLength()
	start := it.Start
	slice := make([]byte, constants.Zero, length)
	var i byte

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

// IsWithinRange it.Start <= value && value <= it.End
func (it *RangeByte) IsWithinRange(value byte) bool {
	return it.Start <= value && value <= it.End
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (it *RangeByte) IsValidPlusWithinRange(value byte) bool {
	return it.IsValid && it.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (it *RangeByte) IsInvalidValue(value byte) bool {
	return !it.IsValid || !it.IsWithinRange(value)
}

func (it *RangeByte) String() string {
	return it.BaseRange.String(it.Start, it.End)
}
