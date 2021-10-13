package coredynamic

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type KeyVal struct {
	Key   interface{}
	Value interface{}
}

func (it *KeyVal) KeyDynamic() Dynamic {
	return NewDynamic(it.Key, true)
}

func (it *KeyVal) ValueDynamic() Dynamic {
	return NewDynamic(it.Value, true)
}

func (it *KeyVal) KeyDynamicPtr() *Dynamic {
	return NewDynamicPtr(it.Key, true)
}

func (it *KeyVal) ValueDynamicPtr() *Dynamic {
	return NewDynamicPtr(it.Value, true)
}

func (it *KeyVal) String() string {
	return fmt.Sprintf(
		constants.KeyValuePariSimpleFormat,
		it.Key,
		it.Key,
		it.Value,
		it.Value)
}

func (it *KeyVal) KeyString() string {
	if it == nil || it.Key == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Key,
	)
}

func (it *KeyVal) ValueString() string {
	if it == nil || it.Value == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.Value,
	)
}
