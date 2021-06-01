package coredynamic

import (
	"encoding/json"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

type Collection struct {
	items *[]*Dynamic
}

func EmptyDynamicCollection() *Collection {
	return NewDynamicCollection(constants.Zero)
}

func NewDynamicCollection(capacity int) *Collection {
	slice := make([]*Dynamic, 0, capacity)

	return &Collection{items: &slice}
}

func (receiver *Collection) At(index int) *Dynamic {
	return (*receiver.items)[index]
}

func (receiver *Collection) Items() *[]*Dynamic {
	return receiver.items
}

func (receiver *Collection) Length() int {
	return len(*receiver.items)
}

func (receiver *Collection) Count() int {
	return len(*receiver.items)
}

func (receiver *Collection) IsEmpty() bool {
	return len(*receiver.items) == 0
}

func (receiver *Collection) HasAnyItem() bool {
	return len(*receiver.items) > 0
}

func (receiver *Collection) LastIndex() int {
	return len(*receiver.items) - 1
}

func (receiver *Collection) HasIndex(index int) bool {
	return receiver.LastIndex() >= index
}

func (receiver *Collection) ListStringsPtr() *[]string {
	slice := make([]string, constants.Zero, receiver.Length()+1)

	for _, dynamic := range *receiver.items {
		slice = append(slice, dynamic.StringJsonMust())
	}

	return &slice
}

func (receiver *Collection) ListStrings() []string {
	return *receiver.ListStringsPtr()
}

func (receiver *Collection) String() string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		*receiver.items)
}

func (receiver *Collection) StringJson() (jsonString string, err error) {
	toBytes, err := json.Marshal(receiver.items)

	if err != nil {
		return constants.EmptyString, nil
	}

	return string(toBytes), nil
}

func (receiver *Collection) StringJsonMust() string {
	toString, err := receiver.StringJson()

	if err != nil {
		msgtype.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), receiver.items)
	}

	return toString
}

func (receiver *Collection) RemoveAt(index int) (isSuccess bool) {
	if !receiver.HasIndex(index) {
		return false
	}

	items := *receiver.items
	*receiver.items = append(
		items[:index],
		items[index+constants.One:]...)

	return true
}

func (receiver *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.items)
}

func (receiver *Collection) UnmarshalJSON(data []byte) error {
	return msgtype.
		NotImplemented.
		Error(msgtype.UnMarshallingFailed.String(), data)
}

func (receiver *Collection) AddAny(anyItem interface{}, isValid bool) *Collection {
	*receiver.items = append(
		*receiver.items,
		NewDynamicPtr(anyItem, isValid))

	return receiver
}

func (receiver *Collection) AddAnyNonNull(anyItem interface{}, isValid bool) *Collection {
	if anyItem == nil {
		return receiver
	}

	*receiver.items = append(
		*receiver.items,
		NewDynamicPtr(anyItem, isValid))

	return receiver
}

func (receiver *Collection) AddAnyMany(anyItems ...interface{}) *Collection {
	if anyItems == nil {
		return receiver
	}

	for _, item := range anyItems {
		*receiver.items = append(
			*receiver.items,
			NewDynamicPtr(item, true))
	}

	return receiver
}

func (receiver *Collection) Add(dynamic Dynamic) *Collection {
	*receiver.items = append(*receiver.items, &dynamic)

	return receiver
}

func (receiver *Collection) AddPtr(dynamic *Dynamic) *Collection {
	*receiver.items = append(*receiver.items, dynamic)

	return receiver
}

func (receiver *Collection) AddManyPtr(dynamicItems ...*Dynamic) *Collection {
	if dynamicItems == nil {
		return receiver
	}

	for _, item := range dynamicItems {
		*receiver.items = append(*receiver.items, item)
	}

	return receiver
}

func (receiver *Collection) AddNonEmpty(dynamic *Dynamic) *Collection {
	if dynamic == nil {
		return receiver
	}

	*receiver.items = append(*receiver.items, dynamic)

	return receiver
}
