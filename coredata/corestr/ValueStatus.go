package corestr

import "gitlab.com/auk-go/core/constants"

type ValueStatus struct {
	ValueValid *ValidValue
	Index      int
}

func InvalidValueStatusNoMessage() *ValueStatus {
	return InvalidValueStatus(constants.EmptyString)
}

func InvalidValueStatus(message string) *ValueStatus {
	return &ValueStatus{
		ValueValid: InvalidValidValue(message),
		Index:      constants.InvalidNotFoundCase,
	}
}

func (v *ValueStatus) Clone() *ValueStatus {
	return &ValueStatus{
		ValueValid: v.ValueValid.Clone(),
		Index:      v.Index,
	}
}
