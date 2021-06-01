package corerange

import (
	"strings"

	"gitlab.com/evatix-go/core/coreindexes"
)

type StartEndString struct {
	*BaseRange
	Start, End string
}

func NewStartEndString(raw, sep string) *StartEndString {
	ranges := strings.Split(raw, sep)
	length := len(ranges)
	hasStart := length >= 1
	hasEnd := length >= 2
	isValid := false
	isValid = isValid &&
		length == 2 &&
		hasStart &&
		hasEnd

	var start, end string

	if hasStart {
		start = ranges[coreindexes.First]
	}

	if hasEnd {
		end = ranges[coreindexes.Second]
	}

	return &StartEndString{
		BaseRange: &BaseRange{
			RawInput:  raw,
			Separator: sep,
			IsValid:   isValid,
			HasStart:  hasStart,
			HasEnd:    hasEnd,
		},
		Start: start,
		End:   end,
	}
}

func (r *StartEndString) CreateRangeString() *RangeString {
	return &RangeString{
		StartEndString: NewStartEndString(
			r.RawInput,
			r.Separator),
	}
}

func (r *StartEndString) String() string {
	return r.BaseRange.String(r.Start, r.End)
}
