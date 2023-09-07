package coredynamic

import "gitlab.com/auk-go/core/constants"

type ValueStatus struct {
	IsValid bool
	Message string
	Index   int
	Value   interface{}
}

func InvalidValueStatusNoMessage() *ValueStatus {
	return InvalidValueStatus(constants.EmptyString)
}

func InvalidValueStatus(message string) *ValueStatus {
	return &ValueStatus{
		IsValid: false,
		Message: message,
		Index:   constants.InvalidNotFoundCase,
		Value:   nil,
	}
}
