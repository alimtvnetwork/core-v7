package corevalidator

import (
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

type RangesSegment struct {
	corerange.RangeInt
	ExpectedLines []string
	CompareAs     stringcompareas.Variant
	Condition
}
