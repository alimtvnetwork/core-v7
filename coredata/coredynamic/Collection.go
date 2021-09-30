package coredynamic

import (
	"encoding/json"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

type Collection struct {
	items []Dynamic
}

func EmptyDynamicCollection() *Collection {
	return NewDynamicCollection(constants.Zero)
}

func NewDynamicCollection(capacity int) *Collection {
	slice := make([]Dynamic, 0, capacity)

	return &Collection{items: slice}
}

func (it *Collection) At(index int) Dynamic {
	return it.items[index]
}

func (it *Collection) Items() []Dynamic {
	return it.items
}

func (it *Collection) Length() int {
	return len(it.items)
}

func (it *Collection) Count() int {
	return len(it.items)
}

func (it *Collection) IsEmpty() bool {
	return len(it.items) == 0
}

func (it *Collection) HasAnyItem() bool {
	return len(it.items) > 0
}

func (it *Collection) LastIndex() int {
	return len(it.items) - 1
}

func (it *Collection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *Collection) ListStringsPtr() *[]string {
	slice := make([]string, constants.Zero, it.Length()+1)

	for _, dynamic := range it.items {
		slice = append(slice, dynamic.JsonStringMust())
	}

	return &slice
}

func (it *Collection) ListStrings() []string {
	return *it.ListStringsPtr()
}

func (it *Collection) String() string {
	return strings.Join(it.ListStrings(), constants.NewLineUnix)
}

func (it *Collection) JsonString() (jsonString string, err error) {
	toBytes, err := json.Marshal(it.items)

	if err != nil {
		return constants.EmptyString, nil
	}

	return string(toBytes), nil
}

func (it *Collection) JsonStringMust() string {
	toString, err := it.JsonString()

	if err != nil {
		msgtype.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), it.items)
	}

	return toString
}

func (it *Collection) RemoveAt(index int) (isSuccess bool) {
	if !it.HasIndex(index) {
		return false
	}

	items := it.items
	it.items = append(
		items[:index],
		items[index+constants.One:]...)

	return true
}

func (it *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.items)
}

func (it *Collection) UnmarshalJSON(data []byte) error {
	return msgtype.
		NotImplemented.
		Error(msgtype.UnMarshallingFailed.String(), data)
}

func (it *Collection) AddAny(anyItem interface{}, isValid bool) *Collection {
	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid))

	return it
}

func (it *Collection) AddAnyNonNull(anyItem interface{}, isValid bool) *Collection {
	if anyItem == nil {
		return it
	}

	it.items = append(
		it.items,
		NewDynamic(anyItem, isValid))

	return it
}

func (it *Collection) AddAnyMany(anyItems ...interface{}) *Collection {
	if anyItems == nil {
		return it
	}

	for _, item := range anyItems {
		it.items = append(
			it.items,
			NewDynamic(item, true))
	}

	return it
}

func (it *Collection) Add(dynamic Dynamic) *Collection {
	it.items = append(it.items, dynamic)

	return it
}

func (it *Collection) AddPtr(dynamic *Dynamic) *Collection {
	if dynamic == nil {
		return it
	}

	it.items = append(it.items, *dynamic)

	return it
}

func (it *Collection) AddManyPtr(dynamicItems ...*Dynamic) *Collection {
	if dynamicItems == nil {
		return it
	}

	for _, item := range dynamicItems {
		if item == nil {
			continue
		}

		it.items = append(it.items, *item)
	}

	return it
}
