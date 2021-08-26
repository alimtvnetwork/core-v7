package coreonce

import (
	"encoding/json"
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/internal/utilstringinternal"
	"gitlab.com/evatix-go/core/issetter"
)

type StringOnce struct {
	innerData       string
	initializerFunc func() string
	isInitialized   issetter.Value
}

func NewStringOnce(initializerFunc func() string) StringOnce {
	return StringOnce{
		initializerFunc: initializerFunc,
	}
}

func NewStringOncePtr(initializerFunc func() string) *StringOnce {
	return &StringOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *StringOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *StringOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = issetter.True

	return json.Unmarshal(data, &it.innerData)
}

func (it *StringOnce) ValuePtr() *string {
	val := it.Value()

	return &val
}

func (it *StringOnce) Value() string {
	if it.isInitialized.IsTrue() {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = issetter.True

	return it.innerData
}

func (it *StringOnce) IsEqual(equalString string) bool {
	return it.Value() == equalString
}

func (it *StringOnce) IsContains(equalString string) bool {
	return strings.Contains(it.Value(), equalString)
}

func (it *StringOnce) IsEmpty() bool {
	return it.Value() == ""
}

func (it *StringOnce) IsEmptyOrWhitespace() bool {
	return utilstringinternal.IsEmptyOrWhitespace(it.Value())
}

func (it *StringOnce) Bytes() []byte {
	return []byte(it.Value())
}

func (it *StringOnce) Error() error {
	return errors.New(it.Value())
}

func (it *StringOnce) String() string {
	return it.Value()
}
