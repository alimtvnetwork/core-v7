package corerange

import (
	"gitlab.com/evatix-go/core/constants"
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

func (r *RangeByte) Difference() byte {
	return r.End - r.Start
}

func (r *RangeByte) DifferenceAbsolute() byte {
	diff := r.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (r *RangeByte) RangeLength() byte {
	return r.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (r *RangeByte) RangesInt() *[]byte {
	return r.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (r *RangeByte) Ranges() *[]byte {
	if r.IsInvalid() {
		return &[]byte{}
	}

	length := r.RangeLength()
	start := r.Start
	slice := make([]byte, constants.Zero, length)
	var i byte

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

// IsWithinRange r.Start >= value && value <= r.End
func (r *RangeByte) IsWithinRange(value byte) bool {
	return r.Start >= value && value <= r.End
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (r *RangeByte) IsValidPlusWithinRange(value byte) bool {
	return r.IsValid && r.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (r *RangeByte) IsInvalidValue(value byte) bool {
	return !r.IsValid || !r.IsWithinRange(value)
}

func (r *RangeByte) String() string {
	return r.BaseRange.String(r.Start, r.End)
}
