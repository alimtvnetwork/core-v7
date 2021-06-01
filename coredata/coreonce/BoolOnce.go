package coreonce

import (
	"encoding/json"
	"strconv"

	"gitlab.com/evatix-go/core/issetter"
)

type BoolOnce struct {
	innerData       bool
	initializerFunc func() bool
	isInitialized   issetter.Value
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

func (receiver *BoolOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.Value())
}

func (receiver *BoolOnce) UnmarshalJSON(data []byte) error {
	receiver.isInitialized = issetter.True

	return json.Unmarshal(data, &receiver.innerData)
}

func (receiver *BoolOnce) Value() bool {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

func (receiver *BoolOnce) String() string {
	return strconv.FormatBool(receiver.Value())
}
