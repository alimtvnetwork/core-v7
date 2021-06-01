package corerange

import (
	"math"

	"gitlab.com/evatix-go/core/constants"
)

type RangeInt8 struct {
	*BaseRange
	Start, End int8
}

func NewRangeInt8MinMax(
	rawString, separator string,
	min, max int8,
) *RangeInt8 {
	minMax := MinMaxInt{
		Min: int(min),
		Max: int(max),
	}

	return NewRangeInt(
		rawString,
		separator,
		&minMax).
		CreateRangeInt8()
}

func NewRangeInt8(
	rawString, separator string,
	minMaxInt8 *MinMaxInt8,
) *RangeInt8 {
	if minMaxInt8 == nil {
		minMax := &MinMaxInt{
			Min: math.MinInt8,
			Max: math.MaxInt8,
		}

		rangeInt := NewRangeInt(
			rawString,
			separator,
			minMax)

		return rangeInt.CreateRangeInt8()
	}

	minMax := minMaxInt8.CreateMinMaxInt()

	rangeInt := NewRangeInt(
		rawString,
		separator,
		minMax)

	return rangeInt.CreateRangeInt8()
}

func (r *RangeInt8) Difference() int8 {
	return r.End - r.Start
}

func (r *RangeInt8) DifferenceAbsolute() int8 {
	diff := r.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (r *RangeInt8) RangeLength() int8 {
	return r.DifferenceAbsolute() + 1
}

// RangesInt8 returns empty ints if IsInvalid
// return range int values
func (r *RangeInt8) RangesInt8() *[]int8 {
	return r.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (r *RangeInt8) Ranges() *[]int8 {
	if r.IsInvalid() {
		return &[]int8{}
	}

	length := r.RangeLength()
	start := r.Start
	slice := make([]int8, constants.Zero, length)

	var i int8

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

func (r *RangeInt8) String() string {
	return r.BaseRange.String(r.Start, r.End)
}

// IsWithinRange r.Start >= value && value <= r.End
func (r *RangeInt8) IsWithinRange(value int8) bool {
	return r.Start >= value && value <= r.End
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (r *RangeInt8) IsValidPlusWithinRange(value int8) bool {
	return r.IsValid && r.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (r *RangeInt8) IsInvalidValue(value int8) bool {
	return !r.IsValid || !r.IsWithinRange(value)
}
