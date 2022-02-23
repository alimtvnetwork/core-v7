package coreonce

import (
	"encoding/json"
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type AnyOnce struct {
	innerData       interface{}
	initializerFunc func() interface{}
	isInitialized   bool
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

func (it *AnyOnce) Value() interface{} {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *AnyOnce) IsNull() bool {
	return it.Value() == nil
}

func (it *AnyOnce) IsStringEmpty() bool {
	return it.String() == ""
}

func (it *AnyOnce) IsStringEmptyOrWhitespace() bool {
	return strings.TrimSpace(it.String()) == ""
}

func (it *AnyOnce) String() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, it.Value())
}

func (it *AnyOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
