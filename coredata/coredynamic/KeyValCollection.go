package coredynamic

import (
	"encoding/json"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

type KeyValCollection struct {
	items *[]*KeyVal
}

func EmptyKeyValCollection() *KeyValCollection {
	return NewKeyValCollection(constants.Zero)
}

func NewKeyValCollection(capacity int) *KeyValCollection {
	slice := make([]*KeyVal, 0, capacity)

	return &KeyValCollection{items: &slice}
}

func (receiver *KeyValCollection) Length() int {
	return len(*receiver.items)
}

func (receiver *KeyValCollection) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *KeyValCollection) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *KeyValCollection) AddPtr(
	keyVal *KeyVal,
) *KeyValCollection {
	*receiver.items = append(*receiver.items, keyVal)

	return receiver
}

func (receiver *KeyValCollection) AddMany(
	keyValues ...*KeyVal,
) *KeyValCollection {
	if keyValues == nil || len(keyValues) == 0 {
		return receiver
	}

	for _, keyVal := range keyValues {
		*receiver.items = append(
			*receiver.items, keyVal)
	}

	return receiver
}

func (receiver *KeyValCollection) Items() *[]*KeyVal {
	return receiver.items
}

func (receiver *KeyValCollection) String() string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		*receiver.items)
}

func (receiver *KeyValCollection) StringJson() (jsonString string, err error) {
	toBytes, err := json.Marshal(receiver.items)

	if err != nil {
		return constants.EmptyString, nil
	}

	return string(toBytes), nil
}

func (receiver *KeyValCollection) StringJsonMust() string {
	toString, err := receiver.StringJson()

	if err != nil {
		msgtype.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), receiver.items)
	}

	return toString
}
