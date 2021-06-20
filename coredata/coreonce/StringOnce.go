package coreonce

import (
	"encoding/json"
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/internal/strutilinternal"
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

func (receiver *StringOnce) MarshalJSON() ([]byte, error) {

	return json.Marshal(receiver.Value())
}

func (receiver *StringOnce) UnmarshalJSON(data []byte) error {
	receiver.isInitialized = issetter.True

	return json.Unmarshal(data, &receiver.innerData)
}

func (receiver *StringOnce) ValuePtr() *string {
	val := receiver.Value()

	return &val
}

func (receiver *StringOnce) Value() string {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

func (receiver *StringOnce) IsEqual(equalString string) bool {
	return receiver.Value() == equalString
}

func (receiver *StringOnce) IsContains(equalString string) bool {
	return strings.Contains(receiver.Value(), equalString)
}

func (receiver *StringOnce) IsEmpty() bool {
	return receiver.Value() == ""
}

func (receiver *StringOnce) IsEmptyOrWhitespace() bool {
	return strutilinternal.IsEmptyOrWhitespace(receiver.Value())
}

func (receiver *StringOnce) Bytes() []byte {
	return []byte(receiver.Value())
}

func (receiver *StringOnce) Error() error {
	return errors.New(receiver.Value())
}

func (receiver *StringOnce) String() string {
	return receiver.Value()
}
