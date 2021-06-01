package corerange

type RangeString struct {
	*StartEndString
}

func NewRangeString(raw, sep string) *RangeString {
	return &RangeString{
		StartEndString: NewStartEndString(raw, sep),
	}
}

func (r *RangeString) String() string {
	return r.BaseRange.String(r.Start, r.End)
}
