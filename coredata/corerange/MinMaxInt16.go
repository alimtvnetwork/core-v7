package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type MinMaxInt16 struct {
	Min, Max int16
}

func (it *MinMaxInt16) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(it.Min),
		Max: int(it.Min),
	}
}

func (it *MinMaxInt16) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt())
}

func (it *MinMaxInt16) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (it *MinMaxInt16) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (it *MinMaxInt16) Difference() int16 {
	return it.Max - it.Min
}

func (it *MinMaxInt16) DifferenceAbsolute() int16 {
	diff := it.Difference()

	if diff < 0 {
		return diff * -1
	}

	return diff
}

func (it *MinMaxInt16) IsMinEqual(val int16) bool {
	return it != nil && it.Min == val
}

func (it *MinMaxInt16) IsMinAboveEqual(val int16) bool {
	return it != nil && it.Min >= val
}

func (it *MinMaxInt16) IsMinAbove(val int16) bool {
	return it != nil && it.Min > val
}

func (it *MinMaxInt16) IsMinLess(val int16) bool {
	return it != nil && it.Min < val
}

func (it *MinMaxInt16) IsMinLessEqual(val int16) bool {
	return it != nil && it.Min <= val
}

func (it *MinMaxInt16) IsMaxEqual(val int16) bool {
	return it != nil && it.Max == val
}

func (it *MinMaxInt16) IsMaxAboveEqual(val int16) bool {
	return it != nil && it.Max >= val
}

func (it *MinMaxInt16) IsMaxAbove(val int16) bool {
	return it != nil && it.Max > val
}

func (it *MinMaxInt16) IsMaxLess(val int16) bool {
	return it != nil && it.Max < val
}

func (it *MinMaxInt16) IsMaxLessEqual(val int16) bool {
	return it != nil && it.Max <= val
}

func (it *MinMaxInt16) RangeLengthInt() int {
	return int(it.RangeLength())
}

// RangeLength (5 - 3 = 2) + 1
func (it *MinMaxInt16) RangeLength() int16 {
	return it.DifferenceAbsolute() + 1
}

// RangesInt
//
//	returns empty integers if IsInvalid
//	return range int values
func (it *MinMaxInt16) RangesInt() []int {
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

func (it *MinMaxInt16) Ranges() []int16 {
	length := it.RangeLength()
	start := it.Min
	slice := make([]int16, length)
	var i int16

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *MinMaxInt16) CreateRanges(minMaxRanges ...MinMaxInt16) []int16 {
	if len(minMaxRanges) == 0 {
		return it.Ranges()
	}

	firstRanges := it.Ranges()
	totalPossible := len(firstRanges)
	for _, maxRange := range minMaxRanges {
		totalPossible += int(maxRange.DifferenceAbsolute())
	}

	slice := make([]int16, 0, totalPossible)
	slice = append(slice, firstRanges...)
	for _, maxRange := range minMaxRanges {
		slice = append(slice, maxRange.Ranges()...)
	}

	return slice
}

// RangesExcept
//
// Returns ranges only without the except items
func (it *MinMaxInt16) RangesExcept(exceptItems ...int) []int16 {
	length := it.RangeLength()
	start := it.Min
	slice := make([]int16, 0, length)
	toHashmap := convertinteranl.
		Integers.
		ToMapBool(exceptItems...)

	for i := 0; i < int(length); i++ {
		id := start + int16(i)
		if toHashmap[int(id)] {
			continue
		}

		// add not exist
		slice = append(slice, id)
	}

	return slice
}

// IsWithinRange it.Min <= value && value <= it.Max
func (it *MinMaxInt16) IsWithinRange(value int16) bool {
	return it.Min <= value && value <= it.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (it *MinMaxInt16) IsInvalidValue(value int16) bool {
	return !it.IsWithinRange(value)
}

func (it MinMaxInt16) IsOutOfRange(value int16) bool {
	return !it.IsWithinRange(value)
}

func (it *MinMaxInt16) ClonePtr() *MinMaxInt16 {
	if it == nil {
		return nil
	}

	return &MinMaxInt16{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt16) Clone() MinMaxInt16 {
	return MinMaxInt16{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt16) IsEqual(right *MinMaxInt16) bool {
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

func (it MinMaxInt16) String() string {
	return fmt.Sprintf(
		constants.SprintFormatNumberWithHyphen,
		it.Min,
		it.Max)
}
