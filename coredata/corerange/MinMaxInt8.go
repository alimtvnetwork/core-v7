package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type MinMaxInt8 struct {
	Min, Max int8
}

func (it *MinMaxInt8) CreateMinMaxInt() *MinMaxInt {
	return &MinMaxInt{
		Min: int(it.Min),
		Max: int(it.Min),
	}
}

func (it *MinMaxInt8) CreateRangeInt(rawString, separator string) *RangeInt {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt())
}

func (it *MinMaxInt8) CreateRangeInt8(rawString, separator string) *RangeInt8 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt8()
}

func (it *MinMaxInt8) CreateRangeInt16(rawString, separator string) *RangeInt16 {
	return NewRangeInt(
		rawString,
		separator,
		it.CreateMinMaxInt()).
		CreateRangeInt16()
}

func (it *MinMaxInt8) Difference() int8 {
	return it.Max - it.Min
}

func (it *MinMaxInt8) DifferenceAbsolute() int8 {
	diff := it.Difference()

	if diff < 0 {
		return diff * -1
	}

	return diff
}

func (it *MinMaxInt8) IsMinEqual(val int8) bool {
	return it != nil && it.Min == val
}

func (it *MinMaxInt8) IsMinAboveEqual(val int8) bool {
	return it != nil && it.Min >= val
}

func (it *MinMaxInt8) IsMinAbove(val int8) bool {
	return it != nil && it.Min > val
}

func (it *MinMaxInt8) IsMinLess(val int8) bool {
	return it != nil && it.Min < val
}

func (it *MinMaxInt8) IsMinLessEqual(val int8) bool {
	return it != nil && it.Min <= val
}

func (it *MinMaxInt8) IsMaxEqual(val int8) bool {
	return it != nil && it.Max == val
}

func (it *MinMaxInt8) IsMaxAboveEqual(val int8) bool {
	return it != nil && it.Max >= val
}

func (it *MinMaxInt8) IsMaxAbove(val int8) bool {
	return it != nil && it.Max > val
}

func (it *MinMaxInt8) IsMaxLess(val int8) bool {
	return it != nil && it.Max < val
}

func (it *MinMaxInt8) IsMaxLessEqual(val int8) bool {
	return it != nil && it.Max <= val
}

func (it *MinMaxInt8) RangeLengthInt() int {
	return int(it.RangeLength())
}

// RangeLength (5 - 3 = 2) + 1
func (it *MinMaxInt8) RangeLength() int8 {
	return it.DifferenceAbsolute() + 1
}

// RangesInt
//
//	returns empty integers if IsInvalid
//	return range int values
func (it *MinMaxInt8) RangesInt() []int {
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
func (it *MinMaxInt8) Ranges() []int8 {
	length := it.RangeLength()
	start := it.Min
	slice := make([]int8, length)
	var i int8

	for i = 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *MinMaxInt8) CreateRanges(minMaxRanges ...MinMaxInt8) []int8 {
	if len(minMaxRanges) == 0 {
		return it.Ranges()
	}

	firstRanges := it.Ranges()
	totalPossible := len(firstRanges)
	for _, maxRange := range minMaxRanges {
		totalPossible += int(maxRange.DifferenceAbsolute())
	}

	slice := make([]int8, 0, totalPossible)
	slice = append(slice, firstRanges...)
	for _, maxRange := range minMaxRanges {
		slice = append(slice, maxRange.Ranges()...)
	}

	return slice
}

// RangesExcept
//
// Returns ranges only without the except items
func (it *MinMaxInt8) RangesExcept(exceptItems ...int8) []int8 {
	length := it.RangeLength()
	start := it.Min
	slice := make([]int8, 0, length)
	toHashmap := convertinteranl.
		Integers.
		Int8ToMapBool(exceptItems...)

	for i := 0; i < int(length); i++ {
		id := start + int8(i)
		if toHashmap[id] {
			continue
		}

		// add not exist
		slice = append(slice, id)
	}

	return slice
}

// IsWithinRange it.Min <= value && value <= it.Max
func (it *MinMaxInt8) IsWithinRange(value int8) bool {
	return it.Min <= value && value <= it.Max
}

// IsInvalidValue  !r.IsWithinRange(value)
func (it *MinMaxInt8) IsInvalidValue(value int8) bool {
	return !it.IsWithinRange(value)
}

func (it MinMaxInt8) IsOutOfRange(value int8) bool {
	return !it.IsWithinRange(value)
}

func (it *MinMaxInt8) ClonePtr() *MinMaxInt8 {
	if it == nil {
		return nil
	}

	return &MinMaxInt8{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt8) Clone() MinMaxInt8 {
	return MinMaxInt8{
		Min: it.Min,
		Max: it.Max,
	}
}

func (it *MinMaxInt8) IsEqual(right *MinMaxInt8) bool {
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

func (it MinMaxInt8) String() string {
	return fmt.Sprintf(
		constants.SprintFormatNumberWithHyphen,
		it.Min,
		it.Max)
}
