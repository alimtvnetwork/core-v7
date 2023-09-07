package corevalidatortestwrappers

import (
	"gitlab.com/auk-go/core/corevalidator"
)

type TextValidatorsWrapper struct {
	Header string
	// ComparingLines is actually the actual data from
	// test but here it is the test cases for the expectation
	ComparingLines        []string
	Validators            corevalidator.TextValidators
	IsSkipOnContentsEmpty bool
	IsCaseSensitive       bool
	ExpectationLines      []string
}
