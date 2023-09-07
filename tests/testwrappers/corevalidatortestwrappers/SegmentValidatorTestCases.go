package corevalidatortestwrappers

import (
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

var SegmentValidatorTestCases = []SegmentValidatorWrapper{
	{
		Header:                "",
		IsSkipOnContentsEmpty: false,
		IsCaseSensitive:       true,
		SimpleSliceRangeValidator: corevalidator.SimpleSliceRangeValidator{
			VerifierSegments: []corevalidator.RangesSegment{
				{
					RangeInt:      corerange.RangeInt{},
					ExpectedLines: nil,
					CompareAs:     stringcompareas.Equal,
					ValidatorCoreCondition: corevalidator.ValidatorCoreCondition{
						IsTrimCompare:        false,
						IsUniqueWordOnly:     false,
						IsNonEmptyWhitespace: false,
						IsSortStringsBySpace: false,
					},
				},
			},
		},
	},
}
