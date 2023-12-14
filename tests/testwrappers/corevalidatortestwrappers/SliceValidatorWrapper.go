package corevalidatortestwrappers

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/corevalidator"
)

type SliceValidatorWrapper struct {
	Case      coretestcases.CaseV1
	Validator corevalidator.SliceValidator
}
