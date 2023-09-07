package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
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
