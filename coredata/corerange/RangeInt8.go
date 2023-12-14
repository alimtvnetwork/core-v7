package corerange

import (
	"math"
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

func (it *RangeInt8) Difference() int8 {
	return it.End - it.Start
}

func (it *RangeInt8) DifferenceAbsolute() int8 {
	diff := it.Difference()

	if diff < 0 {
		return diff * -1
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (it *RangeInt8) RangeLength() int8 {
	return it.DifferenceAbsolute() + 1
}

// RangesInt8 returns empty ints if IsInvalid
// return range int values
func (it *RangeInt8) RangesInt8() []int8 {
	return it.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *RangeInt8) Ranges() []int8 {
	if it.IsInvalid() {
		return []int8{}
	}

	length := it.RangeLength()
	start := it.Start
	slice := make([]int8, length)

	var i int8

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *RangeInt8) String() string {
	return it.BaseRange.String(it.Start, it.End)
}

// IsWithinRange it.End >= value && value >= it.Start
func (it *RangeInt8) IsWithinRange(value int8) bool {
	return it.End >= value && value >= it.Start
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (it *RangeInt8) IsValidPlusWithinRange(value int8) bool {
	return it.IsValid && it.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (it *RangeInt8) IsInvalidValue(value int8) bool {
	return !it.IsValid || !it.IsWithinRange(value)
}
