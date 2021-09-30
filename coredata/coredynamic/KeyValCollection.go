package coredynamic

import (
	"encoding/json"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

type KeyValCollection struct {
	items []*KeyVal
}

func EmptyKeyValCollection() *KeyValCollection {
	return NewKeyValCollection(constants.Zero)
}

func NewKeyValCollection(capacity int) *KeyValCollection {
	slice := make([]*KeyVal, 0, capacity)

	return &KeyValCollection{items: slice}
}

func (it *KeyValCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *KeyValCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *KeyValCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *KeyValCollection) AddPtr(
	keyVal *KeyVal,
) *KeyValCollection {
	it.items = append(it.items, keyVal)

	return it
}

func (it *KeyValCollection) AddMany(
	keyValues ...*KeyVal,
) *KeyValCollection {
	if keyValues == nil || len(keyValues) == 0 {
		return it
	}

	for _, keyVal := range keyValues {
		it.items = append(
			it.items, keyVal)
	}

	return it
}

func (it *KeyValCollection) Items() []*KeyVal {
	return it.items
}

func (it *KeyValCollection) String() string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.items)
}

func (it *KeyValCollection) JsonString() (jsonString string, err error) {
	toBytes, err := json.Marshal(it.items)

	if err != nil {
		return constants.EmptyString, err
	}

	return string(toBytes), err
}

func (it *KeyValCollection) JsonStringMust() string {
	toString, err := it.JsonString()

	if err != nil {
		msgtype.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), it.items)
	}

	return toString
}
