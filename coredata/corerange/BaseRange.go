package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

type BaseRange struct {
	RawInput         string
	Separator        string
	IsValid          bool
	HasStart, HasEnd bool
}

func (it *BaseRange) CreateRangeInt(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(
		it.RawInput,
		it.Separator,
		minMax)
}

func (it *BaseRange) IsInvalid() bool {
	return !it.IsValid
}

func (it *BaseRange) BaseRangeClone() *BaseRange {
	return &BaseRange{
		RawInput:  it.RawInput,
		Separator: it.Separator,
		IsValid:   it.IsValid,
		HasStart:  it.HasStart,
		HasEnd:    it.HasEnd,
	}
}

func (it *BaseRange) String(start, end interface{}) string {
	format := constants.SprintValueFormat +
		it.Separator +
		constants.SprintValueFormat

	return fmt.Sprintf(format, start, end)
}
