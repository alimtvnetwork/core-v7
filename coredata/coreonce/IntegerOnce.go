package coreonce

import (
	"encoding/json"
	"strconv"
)

type IntegerOnce struct {
	innerData       int
	initializerFunc func() int
	isInitialized   bool
}

func NewIntegerOnce(initializerFunc func() int) IntegerOnce {
	return IntegerOnce{
		initializerFunc: initializerFunc,
	}
}

func NewIntegerOncePtr(initializerFunc func() int) *IntegerOnce {
	return &IntegerOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *IntegerOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *IntegerOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *IntegerOnce) Value() int {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

// IsEmpty returns true if zero
func (it *IntegerOnce) IsEmpty() bool {
	return it.Value() == 0
}

func (it *IntegerOnce) IsZero() bool {
	return it.Value() == 0
}

func (it *IntegerOnce) IsNegative() bool {
	return it.Value() < 0
}

func (it *IntegerOnce) IsPositive() bool {
	return it.Value() > 0
}

func (it *IntegerOnce) String() string {
	return strconv.Itoa(it.Value())
}

func (it *IntegerOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
