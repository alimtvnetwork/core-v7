package corerange

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type BaseRange struct {
	RawInput         string
	Separator        string
	IsValid          bool
	HasStart, HasEnd bool
}

func (receiver *BaseRange) CreateRangeInt(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(
		receiver.RawInput,
		receiver.Separator,
		minMax)
}

func (receiver *BaseRange) IsInvalid() bool {
	return !receiver.IsValid
}

func (receiver *BaseRange) BaseRangeClone() *BaseRange {
	return &BaseRange{
		RawInput:  receiver.RawInput,
		Separator: receiver.Separator,
		IsValid:   receiver.IsValid,
		HasStart:  receiver.HasStart,
		HasEnd:    receiver.HasEnd,
	}
}

func (receiver *BaseRange) String(start, end interface{}) string {
	format := constants.SprintValueFormat +
		receiver.Separator +
		constants.SprintValueFormat

	return fmt.Sprint(format, start, end)
}
