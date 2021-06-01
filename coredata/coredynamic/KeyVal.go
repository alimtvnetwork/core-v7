package coredynamic

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type KeyVal struct {
	Key   interface{}
	Value interface{}
}

func (receiver *KeyVal) KeyDynamic() *Dynamic {
	return NewDynamicPtr(receiver.Key, true)
}

func (receiver *KeyVal) ValueDynamic() *Dynamic {
	return NewDynamicPtr(receiver.Value, true)
}

func (receiver *KeyVal) String() string {
	return fmt.Sprintf(
		constants.KeyValuuePariSimpleFormat,
		receiver.Key,
		receiver.Key,
		receiver.Value,
		receiver.Value)
}
