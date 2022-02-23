package coreonce

import (
	"encoding/json"
	"strconv"
)

type BoolOnce struct {
	innerData       bool
	initializerFunc func() bool
	isInitialized   bool
}

func NewBoolOnce(initializerFunc func() bool) BoolOnce {
	return BoolOnce{
		initializerFunc: initializerFunc,
	}
}

func NewBoolOncePtr(initializerFunc func() bool) *BoolOnce {
	return &BoolOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *BoolOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *BoolOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *BoolOnce) Value() bool {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *BoolOnce) String() string {
	return strconv.FormatBool(it.Value())
}

func (it *BoolOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
