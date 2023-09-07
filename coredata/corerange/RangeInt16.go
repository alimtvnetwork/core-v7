package corerange

import (
	"math"

	"gitlab.com/auk-go/core/constants"
)

type RangeInt16 struct {
	*BaseRange
	Start, End int16
}

func NewRangeInt16MinMax(
	rawString, separator string,
	min, max int16,
) *RangeInt16 {
	minMax := MinMaxInt{
		Min: int(min),
		Max: int(max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMax).
		CreateRangeInt16()
}

func NewRangeInt16(
	rawString, separator string,
	minMaxInt16 *MinMaxInt16,
) *RangeInt16 {
	if minMaxInt16 == nil {
		minMax := &MinMaxInt{
			Min: math.MinInt16,
			Max: math.MaxInt16,
		}

		rangeInt := NewRangeInt(
			rawString,
			separator,
			minMax)

		return rangeInt.CreateRangeInt16()
	}

	minMax := minMaxInt16.CreateMinMaxInt()

	rangeInt := NewRangeInt(
		rawString,
		separator,
		minMax)

	return rangeInt.CreateRangeInt16()
}

func (r *RangeInt16) Difference() int16 {
	return r.End - r.Start
}

func (r *RangeInt16) DifferenceAbsolute() int16 {
	diff := r.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (r *RangeInt16) RangeLength() int16 {
	return r.DifferenceAbsolute() + 1
}

// RangesInt16 returns empty ints if IsInvalid
// return range int values
func (r *RangeInt16) RangesInt16() *[]int16 {
	return r.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (r *RangeInt16) Ranges() *[]int16 {
	if r.IsInvalid() {
		return &[]int16{}
	}

	length := r.RangeLength()
	start := r.Start
	slice := make([]int16, constants.Zero, length)

	var i int16

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

func (r *RangeInt16) String() string {
	return r.BaseRange.String(r.Start, r.End)
}

// IsWithinRange r.Start >= value && value <= r.End
func (r *RangeInt16) IsWithinRange(value int16) bool {
	return r.Start >= value && value <= r.End
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (r *RangeInt16) IsValidPlusWithinRange(value int16) bool {
	return r.IsValid && r.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (r *RangeInt16) IsInvalidValue(value int16) bool {
	return !r.IsValid || !r.IsWithinRange(value)
}
