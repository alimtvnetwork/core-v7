package coreonce

import (
	"encoding/json"
)

type BytesOnce struct {
	innerData       []byte
	initializerFunc func() []byte
	isInitialized   bool
}

func NewBytesOnce(initializerFunc func() []byte) BytesOnce {
	return BytesOnce{
		initializerFunc: initializerFunc,
	}
}

func NewBytesOncePtr(initializerFunc func() []byte) *BytesOnce {
	return &BytesOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *BytesOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *BytesOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *BytesOnce) Value() []byte {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *BytesOnce) Execute() []byte {
	return it.Value()
}

// IsEmpty returns true if zero
func (it *BytesOnce) IsEmpty() bool {
	return it == nil || it.initializerFunc == nil || len(it.Value()) == 0
}

func (it *BytesOnce) String() string {
	return string(it.Value())
}

func (it *BytesOnce) Length() int {
	if it == nil || it.initializerFunc == nil {
		return 0
	}

	return len(it.Value())
}

func (it *BytesOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
