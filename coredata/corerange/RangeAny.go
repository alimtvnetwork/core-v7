package corerange

import "gitlab.com/auk-go/core/internal/strutilinternal"

type RangeAny struct {
	*BaseRange
	RawInput   interface{}
	Start, End interface{}
}

func (r *RangeAny) RawInputString() string {
	return strutilinternal.AnyToString(r.RawInput)
}

func (r *RangeAny) StartString() string {
	return strutilinternal.AnyToString(r.Start)
}

func (r *RangeAny) EndString() string {
	return strutilinternal.AnyToString(r.End)
}

func (r *RangeAny) CreateRangeInt() *RangeInt {
	return NewRangeInt(
		r.RawInputString(),
		r.Separator,
		nil)
}

func (r *RangeAny) CreateRangeIntMinMax(minMax *MinMaxInt) *RangeInt {
	return NewRangeInt(
		r.RawInputString(),
		r.Separator,
		minMax)
}

func (r *RangeAny) CreateRangeString() *RangeString {
	return &RangeString{
		StartEndString: r.CreateStartEndString(),
	}
}

func (r *RangeAny) CreateStartEndString() *StartEndString {
	return &StartEndString{
		BaseRange: r.BaseRangeClone(),
		Start:     r.StartString(),
		End:       r.EndString(),
	}
}

func (r *RangeAny) String() string {
	return r.BaseRange.String(r.Start, r.End)
}
