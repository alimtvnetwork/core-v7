package corerange

import "gitlab.com/auk-go/core/constants"

type MinMaxInt struct {
	Min, Max int
}

func (it *MinMaxInt) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		it)
}

func (it *MinMaxInt) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		it).
		CreateRangeInt8()
}

func (it *MinMaxInt) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		it).
		CreateRangeInt16()
}

func (it *MinMaxInt) Difference() int {
	return it.Max - it.Min
}

func (it *MinMaxInt) DifferenceAbsolute() int {
	diff := it.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

func (it *MinMaxInt) IsMinEqual(val int) bool {
	return it != nil && it.Min == val
}

func (it *MinMaxInt) IsMinAboveEqual(val int) bool {
	return it != nil && it.Min >= val
}

func (it *MinMaxInt) IsMinAbove(val int) bool {
	return it != nil && it.Min > val
}

func (it *MinMaxInt) IsMinLess(val int) bool {
	return it != nil && it.Min < val
}

func (it *MinMaxInt) IsMinLessEqual(val int) bool {
	return it != nil && it.Min <= val
}

func (it *MinMaxInt) IsMaxEqual(val int) bool {
	return it != nil && it.Max == val
}

func (it *MinMaxInt) IsMaxAboveEqual(val int) bool {
	return it != nil && it.Max >= val
}

func (it *MinMaxInt) IsMaxAbove(val int) bool {
	return it != nil && it.Max > val
}

func (it *MinMaxInt) IsMaxLess(val int) bool {
	return it != nil && it.Max < val
}

func (it *MinMaxInt) IsMaxLessEqual(val int) bool {
	return it != nil && it.Max <= val
}

func (it *MinMaxInt) RangeLengthInt() int {
	return it.RangeLength()
}

// RangeLength (5 - 3 = 2) + 1
func (it *MinMaxInt) RangeLength() int {
	return it.DifferenceAbsolute() + 1
}

// RangesInt
//
//  returns empty integers if IsInvalid
//  return range int values
func (it *MinMaxInt) RangesInt() []int {
	actualRanges := it.Ranges()
	rangesIntegers := make(
		[]int,
		it.RangeLengthInt())

	for i, actualValue := range actualRanges {
		rangesIntegers[i] = actualValue
	}

	return rangesIntegers
}

// Ranges
//
//  returns empty integers if IsInvalid
//  return range int values
func (it *MinMaxInt) Ranges() []int {
	length := it.RangeLength()
	start := it.Min
	slice := make([]int, constants.Zero, length)

	for i := 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (it *MinMaxInt) IsWithinRange(value int) bool {
	return it.Min >= value && value <= it.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (it *MinMaxInt) IsInvalidValue(value int) bool {
	return !it.IsWithinRange(value)
}

func (it MinMaxInt) IsOutOfRange(value int) bool {
	return !it.IsWithinRange(value)
}

func (it *MinMaxInt) ClonePtr() *MinMaxInt {
	if it == nil {
		return nil
	}

	return &MinMaxInt{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt) Clone() MinMaxInt {
	return MinMaxInt{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt) IsEqual(right *MinMaxInt) bool {
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
