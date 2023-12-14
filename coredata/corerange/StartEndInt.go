package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type StartEndInt struct {
	Start int `json:"Start"`
	End   int `json:"End"`
}

func (it *StartEndInt) IsInvalidStart() bool {
	return it == nil || it.Start <= 0
}

func (it *StartEndInt) IsStartEndBothDefined() bool {
	return it != nil && it.HasStart() && it.HasEnd()
}

func (it *StartEndInt) IsInvalidStartEndBoth() bool {
	return it.IsInvalidStart() && it.IsInvalidEnd()
}

func (it *StartEndInt) IsInvalidAnyStartEnd() bool {
	return it.IsInvalidStart() || it.IsInvalidEnd()
}

func (it *StartEndInt) IsStartGraterThan(val int) bool {
	return it != nil && it.Start > val
}

func (it *StartEndInt) IsEndGraterThan(val int) bool {
	return it != nil && it.End > val
}

func (it *StartEndInt) HasStart() bool {
	return it != nil && it.Start > 0
}

func (it *StartEndInt) IsInvalidEnd() bool {
	return it == nil || it.End <= 0
}

func (it *StartEndInt) HasEnd() bool {
	return it != nil && it.End > 0
}

func (it *StartEndInt) StringUsingFormat(format string) string {
	return fmt.Sprintf(format, it.Start, it.End)
}

func (it *StartEndInt) String() string {
	return fmt.Sprintf("%d-%d", it.Start, it.End)
}

func (it *StartEndInt) StringSpace() string {
	return fmt.Sprintf("%d %d", it.Start, it.End)
}

func (it *StartEndInt) StringHyphen() string {
	return fmt.Sprintf("%d-%d", it.Start, it.End)
}

func (it *StartEndInt) StringColon() string {
	return fmt.Sprintf("%d:%d", it.Start, it.End)
}

func (it *StartEndInt) RangeInt(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(it.StringColon(), constants.Colon, minMax)
}

func (it *StartEndInt) RangeInt16(minMax *MinMaxInt16) *RangeInt16 {
	return NewRangeInt16(it.StringColon(), constants.Colon, minMax)
}

func (it *StartEndInt) RangeInt8(minMax *MinMaxInt8) *RangeInt8 {
	return NewRangeInt8(it.StringColon(), constants.Colon, minMax)
}

// Ranges returns empty ints if IsInvalid
// return range int values
func (it *StartEndInt) Ranges() []int {
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

func (it *StartEndInt) CreateRanges(startEndRanges ...StartEndInt) []int {
	if len(startEndRanges) == 0 {
		return it.Ranges()
	}

	firstRanges := it.Ranges()
	totalPossible := len(firstRanges)
	for _, maxRange := range startEndRanges {
		totalPossible += maxRange.DifferenceAbsolute()
	}

	slice := make([]int, 0, totalPossible)
	slice = append(slice, firstRanges...)
	for _, maxRange := range startEndRanges {
		slice = append(slice, maxRange.Ranges()...)
	}

	return slice
}

func (it *StartEndInt) IsInvalid() bool {
	return it == nil
}

// RangesExcept
//
// Returns ranges only without the except items
func (it *StartEndInt) RangesExcept(exceptItems ...int) []int {
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

func (it *StartEndInt) Diff() int {
	return it.End - it.Start
}

func (it *StartEndInt) DifferenceAbsolute() int {
	diff := it.Diff()

	if diff < 0 {
		return diff * -1
	}

	return diff
}

func (it *StartEndInt) RangeLength() int {
	return it.DifferenceAbsolute() + 1
}
