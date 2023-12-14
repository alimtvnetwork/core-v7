package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

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
		return diff * -1
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
//	returns empty integers if IsInvalid
//	return range int values
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
//	returns empty integers if IsInvalid
//	return range int values
func (it *MinMaxInt64) Ranges() []int64 {
	length := it.RangeLength()
	start := it.Min
	slice := make(
		[]int64,
		length)

	var i int64

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *MinMaxInt64) CreateRanges(minMaxRanges ...MinMaxInt64) []int64 {
	if len(minMaxRanges) == 0 {
		return it.Ranges()
	}

	firstRanges := it.Ranges()
	totalPossible := len(firstRanges)
	for _, maxRange := range minMaxRanges {
		totalPossible += int(maxRange.DifferenceAbsolute())
	}

	slice := make([]int64, 0, totalPossible)
	slice = append(slice, firstRanges...)
	for _, maxRange := range minMaxRanges {
		slice = append(slice, maxRange.Ranges()...)
	}

	return slice
}

// RangesExcept
//
// Returns ranges only without the except items
func (it *MinMaxInt64) RangesExcept(exceptItems ...int) []int64 {
	length := it.RangeLength()
	start := it.Min
	slice := make([]int64, 0, length)
	toHashmap := convertinteranl.
		Integers.
		ToMapBool(exceptItems...)

	for i := 0; i < int(length); i++ {
		id := start + int64(i)
		if toHashmap[int(id)] {
			continue
		}

		// add not exist
		slice = append(slice, id)
	}

	return slice
}

// IsWithinRange it.Min <= value && value <= it.Max
func (it *MinMaxInt64) IsWithinRange(value int64) bool {
	return it != nil && it.Min <= value && value <= it.Max
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

func (it MinMaxInt64) String() string {
	return fmt.Sprintf(
		constants.SprintFormatNumberWithHyphen,
		it.Min,
		it.Max)
}
