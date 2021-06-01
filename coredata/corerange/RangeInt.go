package corerange

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coreindexes"
)

type RangeInt struct {
	*BaseRange
	Start, End int
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
		start, isValid = converters.StringToIntegerWithDefault(
			ranges[coreindexes.First],
			constants.MaxInt)
	}

	if hasEnd {
		end, isValid = converters.StringToIntegerWithDefault(
			ranges[coreindexes.Second],
			constants.MinInt)
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

func (r *RangeInt) Difference() int {
	return r.End - r.Start
}

func (r *RangeInt) DifferenceAbsolute() int {
	diff := r.Difference()

	if diff < 0 {
		return diff
	}

	return diff
}

// RangeLength (5 - 3 = 2) + 1
func (r *RangeInt) RangeLength() int {
	return r.DifferenceAbsolute() + 1
}

// RangesInt returns empty ints if IsInvalid
// return range int values
func (r *RangeInt) RangesInt() *[]int {
	return r.Ranges()
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (r *RangeInt) Ranges() *[]int {
	if r.IsInvalid() {
		return &[]int{}
	}

	length := r.RangeLength()
	start := r.Start
	slice := make([]int, constants.Zero, length)

	for i := 0; i < length; i++ {
		slice[i] = start + i
	}

	return &slice
}

func (r *RangeInt) String() string {
	return r.BaseRange.String(r.Start, r.End)
}

func (r *RangeInt) CreateRangeInt8() *RangeInt8 {
	return &RangeInt8{
		BaseRange: r.BaseRangeClone(),
		Start:     int8(r.Start),
		End:       int8(r.End),
	}
}

func (r *RangeInt) CreateRangeByte() *RangeByte {
	return &RangeByte{
		BaseRange: r.BaseRangeClone(),
		Start:     byte(r.Start),
		End:       byte(r.End),
	}
}

func (r *RangeInt) CreateRangeInt16() *RangeInt16 {
	return &RangeInt16{
		BaseRange: r.BaseRangeClone(),
		Start:     int16(r.Start),
		End:       int16(r.End),
	}
}

func (r *RangeInt) ShallowCreateRangeInt16() *RangeInt16 {
	return &RangeInt16{
		BaseRange: r.BaseRange,
		Start:     int16(r.Start),
		End:       int16(r.End),
	}
}

func (r *RangeInt) ShallowCreateRangeInt8() *RangeInt8 {
	return &RangeInt8{
		BaseRange: r.BaseRange,
		Start:     int8(r.Start),
		End:       int8(r.End),
	}
}

func (r *RangeInt) ShallowCreateRangeByte() *RangeByte {
	return &RangeByte{
		BaseRange: r.BaseRange,
		Start:     byte(r.Start),
		End:       byte(r.End),
	}
}

// IsWithinRange r.Start >= value && value <= r.End
func (r *RangeInt) IsWithinRange(value int) bool {
	return r.Start >= value && value <= r.End
}

// IsValidPlusWithinRange r.IsValid && r.IsWithinRange(value)
func (r *RangeInt) IsValidPlusWithinRange(value int) bool {
	return r.IsValid && r.IsWithinRange(value)
}

// IsInvalidValue !r.IsValid || !r.IsWithinRange(value)
func (r *RangeInt) IsInvalidValue(value int) bool {
	return !r.IsValid || !r.IsWithinRange(value)
}
