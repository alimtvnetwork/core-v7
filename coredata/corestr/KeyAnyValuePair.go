package corestr

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type KeyAnyValuePair struct {
	Key         string
	valueString *string
	Value       interface{}
}

func (keyAnyValuePair *KeyAnyValuePair) IsValueNull() bool {
	return keyAnyValuePair.Value == nil
}

func (keyAnyValuePair *KeyAnyValuePair) HasNonNull() bool {
	return keyAnyValuePair.Value != nil
}

func (keyAnyValuePair *KeyAnyValuePair) ValueString() string {
	return *keyAnyValuePair.ValueStringPtr()
}

func (keyAnyValuePair *KeyAnyValuePair) ValueStringPtr() *string {
	if keyAnyValuePair.valueString == nil && keyAnyValuePair.HasNonNull() {
		valueString := fmt.Sprintf(constants.SprintValueFormat, keyAnyValuePair.Value)
		keyAnyValuePair.valueString = &valueString
	} else if keyAnyValuePair.valueString == nil {
		valueString := ""
		keyAnyValuePair.valueString = &valueString
	}

	return keyAnyValuePair.valueString
}
