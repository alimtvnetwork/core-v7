package corerange

import (
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type RangeInt struct {
	*BaseRange
	Start, End int
}

func NewRangeIntUsingValues(
	start, end int,
	isValid bool,
) *RangeInt {
	return &RangeInt{
		BaseRange: &BaseRange{
			RawInput:  strconv.Itoa(start) + defaultSeparator + strconv.Itoa(end),
			Separator: defaultSeparator,
			IsValid:   isValid,
			HasStart:  start > 0,
			HasEnd:    end > 0,
		},
		Start: start,
		End:   end,
	}
}

func NewRangeIntMinMax(
	rawString, separator string,
	min, max int,
) *RangeInt {
	minMax := MinMaxInt{
		Min: min,
		Max: max,
	}

	return NewRangeInt(rawString, separator, &minMax)
}

// NewRangeInt : MinMaxInt represent no validation on start and end range.
func NewRangeInt(
	rawString, separator string,
	minMax *MinMaxInt,
) *RangeInt {
	ranges := strings.Split(rawString, separator)
	length := len(ranges)
	hasStart := length >= 1
	hasEnd := length >= 2
	isValid := false
	var start, end int

	if hasStart {
		start, isValid = converters.StringTo.IntegerWithDefault(
			ranges[coreindexes.First],
			constants.MaxInt,
		)
	}

	if hasEnd {
		end, isValid = converters.StringTo.IntegerWithDefault(
			ranges[coreindexes.Second],
			constants.MinInt,
		)
	}

	isValid = isValid &&
		length == 2 &&
		hasStart &&
		hasEnd &&
		end > start

	if minMax != nil {
		isValid = isValid &&
			start >= minMax.Min &&
			end <= minMax.Max

		return &RangeInt{
			BaseRange: &BaseRange{
				RawInput:  rawString,
				Separator: separator,
				IsValid:   isValid,
				HasStart:  hasStart,
				HasEnd:    hasEnd,
			},
			Start: start,
			End:   end,
		}
	}

	return &RangeInt{
		BaseRange: &BaseRange{
			RawInput:  rawString,
			Separator: separator,
			IsValid:   isValid,
			HasStart:  hasStart,
			HasEnd:    hasEnd,
		},
		Start: start,
		End:   end,
	}
}

func (it *RangeInt) Difference() int {
	return it.End - it.Start
}

func (it *RangeInt) DifferenceAbsolute() int {
	diff := it.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (it *RangeInt) RangeLength() int {
	return it.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (it *RangeInt) RangesInt() []int {
	return it.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *RangeInt) Ranges() []int {
	if it.IsInvalid() {
		return []int{}
	}

	length := it.RangeLength()
	start := it.Start
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		slice[i] = start + i
	}

	return slice
}

func (it *RangeInt) CreateRanges(minMaxRanges ...MinMaxInt) []int {
	if len(minMaxRanges) == 0 {
		return it.Ranges()
	}

	firstRanges := it.Ranges()
	totalPossible := len(firstRanges)
	for _, maxRange := range minMaxRanges {
		totalPossible += maxRange.DifferenceAbsolute()
	}

	slice := make([]int, 0, totalPossible)
	slice = append(slice, firstRanges...)
	for _, maxRange := range minMaxRanges {
		slice = append(slice, maxRange.Ranges()...)
	}

	return slice
}

// RangesExcept
//
// Returns ranges only without the except items
func (it *RangeInt) RangesExcept(exceptItems ...int) []int {
	length := it.RangeLength()
	start := it.Start
	slice := make([]int, 0, length)
	toHashmap := convertinteranl.
		Integers.
		ToMapBool(exceptItems...)

	for i := 0; i < length; i++ {
		id := start + i
		if toHashmap[id] {
			continue
		}

		// add not exist
		slice = append(slice, id)
	}

	return slice
}

func (it *RangeInt) String() string {
	return it.BaseRange.String(it.Start, it.End)
}

func (it *RangeInt) CreateStartEnd() *StartEndInt {
	return &StartEndInt{
		Start: it.Start,
		End:   it.End,
	}
}

func (it *RangeInt) CreateRangeInt8() *RangeInt8 {
	return &RangeInt8{
		BaseRange: it.BaseRangeClone(),
		Start:     int8(it.Start),
		End:       int8(it.End),
	}
}

func (it *RangeInt) CreateRangeByte() *RangeByte {
	return &RangeByte{
		BaseRange: it.BaseRangeClone(),
		Start:     byte(it.Start),
		End:       byte(it.End),
	}
}

func (it *RangeInt) CreateRangeInt16() *RangeInt16 {
	return &RangeInt16{
		BaseRange: it.BaseRangeClone(),
		Start:     int16(it.Start),
		End:       int16(it.End),
	}
}

func (it *RangeInt) ShallowCreateRangeInt16() *RangeInt16 {
	return &RangeInt16{
		BaseRange: it.BaseRange,
		Start:     int16(it.Start),
		End:       int16(it.End),
	}
}

func (it *RangeInt) ShallowCreateRangeInt8() *RangeInt8 {
	return &RangeInt8{
		BaseRange: it.BaseRange,
		Start:     int8(it.Start),
		End:       int8(it.End),
	}
}

func (it *RangeInt) ShallowCreateRangeByte() *RangeByte {
	return &RangeByte{
		BaseRange: it.BaseRange,
		Start:     byte(it.Start),
		End:       byte(it.End),
	}
}

// IsWithinRange it.End >= value && value >= it.Start
func (it *RangeInt) IsWithinRange(value int) bool {
	return it.End >= value && value >= it.Start
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (it *RangeInt) IsValidPlusWithinRange(value int) bool {
	return it.IsValid && it.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (it *RangeInt) IsInvalidValue(value int) bool {
	return !it.IsValid || !it.IsWithinRange(value)
}
