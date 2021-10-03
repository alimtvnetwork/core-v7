package corestr

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type KeyAnyValuePair struct {
	Key         string
	valueString SimpleStringOnce
	Value       interface{}
}

func (it *KeyAnyValuePair) IsValueNull() bool {
	return it.Value == nil
}

func (it *KeyAnyValuePair) HasNonNull() bool {
	return it.Value != nil
}

func (it *KeyAnyValuePair) ValueString() string {
	if it.valueString.IsInitialized() {
		return it.valueString.String()
	}

	if it.HasNonNull() {
		valueString := fmt.Sprintf(constants.SprintValueFormat, it.Value)

		return it.
			valueString.
			GetPlusSetOnUninitialized(valueString)
	}

	return it.
		valueString.
		GetPlusSetEmptyOnUninitialized()
}
