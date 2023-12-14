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
		RangeSegmentsValidator: corevalidator.RangeSegmentsValidator{
			VerifierSegments: []corevalidator.RangesSegment{

				{
					RangeInt:      corerange.RangeInt{},
					ExpectedLines: nil,
					CompareAs:     stringcompareas.Equal,
					Condition:     corevalidator.DefaultDisabledCoreCondition,
				},
			},
		},
	},
}
