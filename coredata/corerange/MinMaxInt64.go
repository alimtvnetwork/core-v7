package corerange

import "gitlab.com/auk-go/core/constants"

type MinMaxInt64 struct {
	Min, Max int64
}

func (it *MinMaxInt64) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(it.Min),
		Max: int(it.Min),
	}
}

func (it *MinMaxInt64) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt())
}

func (it *MinMaxInt64) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (it *MinMaxInt64) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (it *MinMaxInt64) Difference() int64 {
	return it.Max - it.Min
}

func (it *MinMaxInt64) DifferenceAbsolute() int64 {
	diff := it.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

func (it *MinMaxInt64) IsMinEqual(val int64) bool {
	return it != nil && it.Min == val
}

func (it *MinMaxInt64) IsMinAboveEqual(val int64) bool {
	return it != nil && it.Min >= val
}

func (it *MinMaxInt64) IsMinAbove(val int64) bool {
	return it != nil && it.Min > val
}

func (it *MinMaxInt64) IsMinLess(val int64) bool {
	return it != nil && it.Min < val
}

func (it *MinMaxInt64) IsMinLessEqual(val int64) bool {
	return it != nil && it.Min <= val
}

func (it *MinMaxInt64) IsMaxEqual(val int64) bool {
	return it != nil && it.Max == val
}

func (it *MinMaxInt64) IsMaxAboveEqual(val int64) bool {
	return it != nil && it.Max >= val
}

func (it *MinMaxInt64) IsMaxAbove(val int64) bool {
	return it != nil && it.Max > val
}

func (it *MinMaxInt64) IsMaxLess(val int64) bool {
	return it != nil && it.Max < val
}

func (it *MinMaxInt64) IsMaxLessEqual(val int64) bool {
	return it != nil && it.Max <= val
}

func (it *MinMaxInt64) RangeLengthInt() int {
	return int(it.RangeLength())
}

// RangeLength (5 - 3 = 2) + 1
func (it *MinMaxInt64) RangeLength() int64 {
	return it.DifferenceAbsolute() + 1
}

// RangesInt
//
//  returns empty integers if IsInvalid
//  return range int values
func (it *MinMaxInt64) RangesInt() []int {
	actualRanges := it.Ranges()
	rangesIntegers := make(
		[]int,
		it.RangeLengthInt())

	for i, actualValue := range actualRanges {
		rangesIntegers[i] = int(actualValue)
	}

	return rangesIntegers
}

// Ranges
//
//  returns empty integers if IsInvalid
//  return range int values
func (it *MinMaxInt64) Ranges() []int64 {
	length := it.RangeLength()
	start := it.Min
	slice := make(
		[]int64,
		constants.Zero,
		length)

	var i int64

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (it *MinMaxInt64) IsWithinRange(value int64) bool {
	return it != nil && it.Min >= value && value <= it.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (it MinMaxInt64) IsInvalidValue(value int64) bool {
	return !it.IsWithinRange(value)
}

func (it MinMaxInt64) IsOutOfRange(value int64) bool {
	return !it.IsWithinRange(value)
}

func (it *MinMaxInt64) ClonePtr() *MinMaxInt64 {
	if it == nil {
		return nil
	}

	return &MinMaxInt64{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt64) Clone() MinMaxInt64 {
	return MinMaxInt64{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt64) IsEqual(right *MinMaxInt64) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return true
	}

	if it == right {
		return true
	}

	return it.Max == right.Max &&
		it.Min == right.Min
}
