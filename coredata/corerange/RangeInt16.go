package corerange

import (
	"math"
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

func (it *RangeInt16) Difference() int16 {
	return it.End - it.Start
}

func (it *RangeInt16) DifferenceAbsolute() int16 {
	diff := it.Difference()

	if diff < 0 {
		return diff * -1
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (it *RangeInt16) RangeLength() int16 {
	return it.DifferenceAbsolute() + 1
}

// RangesInt16 returns empty ints if IsInvalid
// return range int values
func (it *RangeInt16) RangesInt16() []int16 {
	return it.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *RangeInt16) Ranges() []int16 {
	if it.IsInvalid() {
		return []int16{}
	}

	length := it.RangeLength()
	start := it.Start
	slice := make([]int16, length)

	var i int16

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *RangeInt16) String() string {
	return it.BaseRange.String(it.Start, it.End)
}

// IsWithinRange it.End >= value && value >= it.Start
func (it *RangeInt16) IsWithinRange(value int16) bool {
	return it.End >= value && value >= it.Start
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (it *RangeInt16) IsValidPlusWithinRange(value int16) bool {
	return it.IsValid && it.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (it *RangeInt16) IsInvalidValue(value int16) bool {
	return !it.IsValid || !it.IsWithinRange(value)
}
