package coreonce

import (
	"encoding/json"
	"strconv"
)

type ByteOnce struct {
	innerData       byte
	initializerFunc func() byte
	isInitialized   bool
}

func NewByteOnce(initializerFunc func() byte) ByteOnce {
	return ByteOnce{
		initializerFunc: initializerFunc,
	}
}

func NewByteOncePtr(initializerFunc func() byte) *ByteOnce {
	return &ByteOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *ByteOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *ByteOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *ByteOnce) Value() byte {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *ByteOnce) Execute() byte {
	return it.Value()
}

func (it *ByteOnce) Int() int {
	return int(it.Value())
}

// IsEmpty returns true if zero
func (it *ByteOnce) IsEmpty() bool {
	return it.Value() == 0
}

func (it *ByteOnce) IsZero() bool {
	return it.Value() == 0
}

func (it *ByteOnce) IsNegative() bool {
	return it.Value() < 0
}

func (it *ByteOnce) IsPositive() bool {
	return it.Value() > 0
}

func (it *ByteOnce) String() string {
	return strconv.Itoa(int(it.Value()))
}

func (it *ByteOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
