package corerange

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type StartEndSimpleString struct {
	Start string `json:"Start"`
	End   string `json:"End"`
}

func (it *StartEndSimpleString) IsInvalidStart() bool {
	return it == nil || it.Start == ""
}

func (it *StartEndSimpleString) IsStartEndBothDefined() bool {
	return it != nil && it.HasStart() && it.HasEnd()
}

func (it *StartEndSimpleString) IsInvalidStartEndBoth() bool {
	return it.IsInvalidStart() && it.IsInvalidEnd()
}

func (it *StartEndSimpleString) IsInvalidAnyStartEnd() bool {
	return it.IsInvalidStart() || it.IsInvalidEnd()
}

func (it *StartEndSimpleString) StartValidValue() *corestr.ValidValue {
	if it == nil {
		return nil
	}

	return corestr.NewValidValue(it.Start)
}

func (it *StartEndSimpleString) EndValidValue() *corestr.ValidValue {
	if it == nil {
		return nil
	}

	return corestr.NewValidValue(it.End)
}

func (it *StartEndSimpleString) StartEndString() *StartEndString {
	if it == nil {
		return nil
	}

	hasStart := it.HasStart()
	hasEnd := it.HasEnd()

	return &StartEndString{
		BaseRange: &BaseRange{
			RawInput:  it.StringHyphen(),
			Separator: constants.Hyphen,
			IsValid:   hasStart == hasEnd && hasStart == true,
			HasStart:  hasEnd,
			HasEnd:    hasStart,
		},
		Start: it.Start,
		End:   it.End,
	}
}

func (it *StartEndSimpleString) HasStart() bool {
	return it != nil && it.Start != ""
}

func (it *StartEndSimpleString) IsInvalidEnd() bool {
	return it == nil || it.End == ""
}

func (it *StartEndSimpleString) HasEnd() bool {
	return it != nil && it.End != ""
}

func (it *StartEndSimpleString) StringUsingFormat(format string) string {
	return fmt.Sprintf(format, it.Start, it.End)
}

func (it *StartEndSimpleString) StringSpace() string {
	return fmt.Sprintf("%s %s", it.Start, it.End)
}

func (it *StartEndSimpleString) StringHyphen() string {
	return fmt.Sprintf("%s-%s", it.Start, it.End)
}

func (it *StartEndSimpleString) StringColon() string {
	return fmt.Sprintf("%s:%s", it.Start, it.End)
}

func (it *StartEndSimpleString) RangeInt(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(
		it.StringColon(),
		constants.Colon,
		minMax)
}

func (it *StartEndSimpleString) RangeInt16(minMax *MinMaxInt16) *RangeInt16 {
	return NewRangeInt16(it.StringColon(), constants.Colon, minMax)
}

func (it *StartEndSimpleString) RangeInt8(minMax *MinMaxInt8) *RangeInt8 {
	return NewRangeInt8(
		it.StringColon(),
		constants.Colon,
		minMax)
}
