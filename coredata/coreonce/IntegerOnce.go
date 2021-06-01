package coreonce

import (
	"encoding/json"
	"strconv"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/issetter"
)

type IntegerOnce struct {
	innerData       int
	initializerFunc func() int
	isInitialized   issetter.Value
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

func (receiver *IntegerOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.Value())
}

func (receiver *IntegerOnce) UnmarshalJSON(data []byte) error {
	receiver.isInitialized = issetter.True

	return json.Unmarshal(data, &receiver.innerData)
}

func (receiver *IntegerOnce) Value() int {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

// IsEmpty returns true if zero
func (receiver *IntegerOnce) IsEmpty() bool {
	return receiver.Value() == 0
}

func (receiver *IntegerOnce) IsZero() bool {
	return receiver.Value() == 0
}

func (receiver *IntegerOnce) IsNegative() bool {
	return receiver.Value() < 0
}

func (receiver *IntegerOnce) IsPositive() bool {
	return receiver.Value() > 0
}

func (receiver *IntegerOnce) IsIntCompareResult(valueCompare int, compare corecomparator.Compare) bool {
	currentValue := receiver.Value()
	switch compare {
	case corecomparator.Equal:
		return currentValue == valueCompare
	case corecomparator.LeftGreater:
		return currentValue > valueCompare
	case corecomparator.LeftGreaterEqual:
		return currentValue >= valueCompare
	case corecomparator.LeftLess:
		return currentValue < valueCompare
	case corecomparator.LeftLessEqual:
		return currentValue <= valueCompare
	case corecomparator.NotEqual:
		return currentValue != valueCompare
	default:
		panic(messages.ComparatorOutOfRangeMessage)
	}
}

func (receiver *IntegerOnce) String() string {
	return strconv.Itoa(receiver.Value())
}
