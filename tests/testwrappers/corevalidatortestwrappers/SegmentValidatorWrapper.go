package corevalidatortestwrappers

import "gitlab.com/auk-go/core/corevalidator"

type SegmentValidatorWrapper struct {
	Header                string
	IsSkipOnContentsEmpty bool
	IsCaseSensitive       bool
	corevalidator.SimpleSliceRangeValidator
}
