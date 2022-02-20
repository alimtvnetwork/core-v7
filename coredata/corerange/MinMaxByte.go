package corerange

import "gitlab.com/evatix-go/core/constants"

type MinMaxByte struct {
	Min, Max byte
}

func (it *MinMaxByte) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(it.Min),
		Max: int(it.Min),
	}
}

func (it *MinMaxByte) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt())
}

func (it *MinMaxByte) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (it *MinMaxByte) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (it *MinMaxByte) Difference() byte {
	return it.Max - it.Min
}

func (it *MinMaxByte) DifferenceAbsolute() byte {
	diff := it.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

func (it *MinMaxByte) IsMinEqual(val byte) bool {
	return it != nil && it.Min == val
}

func (it *MinMaxByte) IsMinAboveEqual(val byte) bool {
	return it != nil && it.Min >= val
}

func (it *MinMaxByte) IsMinAbove(val byte) bool {
	return it != nil && it.Min > val
}

func (it *MinMaxByte) IsMinLess(val byte) bool {
	return it != nil && it.Min < val
}

func (it *MinMaxByte) IsMinLessEqual(val byte) bool {
	return it != nil && it.Min <= val
}

func (it *MinMaxByte) IsMaxEqual(val byte) bool {
	return it != nil && it.Max == val
}

func (it *MinMaxByte) IsMaxAboveEqual(val byte) bool {
	return it != nil && it.Max >= val
}

func (it *MinMaxByte) IsMaxAbove(val byte) bool {
	return it != nil && it.Max > val
}

func (it *MinMaxByte) IsMaxLess(val byte) bool {
	return it != nil && it.Max < val
}

func (it *MinMaxByte) IsMaxLessEqual(val byte) bool {
	return it != nil && it.Max <= val
}

func (it *MinMaxByte) RangeLengthInt() int {
	return int(it.RangeLength())
}

// RangeLength (5 - 3 = 2) + 1
func (it *MinMaxByte) RangeLength() byte {
	return it.DifferenceAbsolute() + 1
}

// RangesInt
//
//  returns empty integers if IsInvalid
//  return range int values
func (it *MinMaxByte) RangesInt() []int {
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
//  returns empty bytes if IsInvalid
//  return range int values
func (it *MinMaxByte) Ranges() []byte {
	length := it.RangeLength()
	start := it.Min
	slice := make(
		[]byte,
		constants.Zero,
		length)

	var i byte

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

// IsWithinRange r.Min >= value && value <= r.Max
func (it *MinMaxByte) IsWithinRange(value byte) bool {
	return it.Min >= value && value <= it.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (it *MinMaxByte) IsInvalidValue(value byte) bool {
	return !it.IsWithinRange(value)
}

func (it MinMaxByte) IsOutOfRange(value byte) bool {
	return !it.IsWithinRange(value)
}

func (it *MinMaxByte) ClonePtr() *MinMaxByte {
	if it == nil {
		return nil
	}

	return &MinMaxByte{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxByte) Clone() MinMaxByte {
	return MinMaxByte{
		Min: it.Min,
		Max: it.Max,
	}
}
