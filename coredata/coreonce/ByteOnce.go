package coreonce

import (
	"encoding/json"
	"strconv"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/internal/messages"
	"gitlab.com/evatix-go/core/issetter"
)

type ByteOnce struct {
	innerData       byte
	initializerFunc func() byte
	isInitialized   issetter.Value
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

func (receiver *ByteOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(receiver.Value())
}

func (receiver *ByteOnce) UnmarshalJSON(data []byte) error {
	receiver.isInitialized = issetter.True

	return json.Unmarshal(data, &receiver.innerData)
}

func (receiver *ByteOnce) Value() byte {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

func (receiver *ByteOnce) Int() int {
	return int(receiver.Value())
}

// IsEmpty returns true if zero
func (receiver *ByteOnce) IsEmpty() bool {
	return receiver.Value() == 0
}

func (receiver *ByteOnce) IsZero() bool {
	return receiver.Value() == 0
}

func (receiver *ByteOnce) IsNegative() bool {
	return receiver.Value() < 0
}

func (receiver *ByteOnce) IsPositive() bool {
	return receiver.Value() > 0
}

func (receiver *ByteOnce) IsByteCompareResult(
	valueCompare byte,
	compare corecomparator.Compare,
) bool {
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

func (receiver *ByteOnce) String() string {
	return strconv.Itoa(int(receiver.Value()))
}
