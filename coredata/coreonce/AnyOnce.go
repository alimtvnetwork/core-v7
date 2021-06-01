package coreonce

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/internal/stringutil"
	"gitlab.com/evatix-go/core/issetter"
)

type AnyOnce struct {
	innerData       interface{}
	initializerFunc func() interface{}
	isInitialized   issetter.Value
}

func NewAnyOnce(initializerFunc func() interface{}) AnyOnce {
	return AnyOnce{
		initializerFunc: initializerFunc,
	}
}

func NewAnyOncePtr(initializerFunc func() interface{}) *AnyOnce {
	return &AnyOnce{
		initializerFunc: initializerFunc,
	}
}

func (receiver *AnyOnce) Value() interface{} {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

func (receiver *AnyOnce) IsNull() bool {
	return receiver.Value() == nil
}

func (receiver *AnyOnce) IsStringEmpty() bool {
	return receiver.String() == ""
}

func (receiver *AnyOnce) IsStringEmptyOrWhitespace() bool {
	return stringutil.IsEmptyOrWhitespace(receiver.String())
}

func (receiver *AnyOnce) String() string {
	if receiver.IsNull() {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, receiver.Value())
}
