package coreonce

import (
	"encoding/json"
	"errors"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/issetter"
)

type ErrorOnce struct {
	innerData       error
	initializerFunc func() error
	isInitialized   issetter.Value
}

func NewErrorOnce(initializerFunc func() error) ErrorOnce {
	return ErrorOnce{
		initializerFunc: initializerFunc,
	}
}

func NewErrorOncePtr(initializerFunc func() error) *ErrorOnce {
	return &ErrorOnce{
		initializerFunc: initializerFunc,
	}
}

func (receiver *ErrorOnce) MarshalJSON() ([]byte, error) {
	if receiver.IsNullOrEmpty() {
		return json.Marshal("")
	}

	return json.Marshal(receiver.Value().Error())
}

func (receiver *ErrorOnce) UnmarshalJSON(data []byte) error {
	receiver.isInitialized = issetter.True
	var str string

	err := json.Unmarshal(data, &str)
	receiver.innerData = errors.New(str)

	return err
}

func (receiver *ErrorOnce) HasError() bool {
	return !receiver.IsNullOrEmpty()
}

func (receiver *ErrorOnce) IsNull() bool {
	return receiver.Value() == nil
}

func (receiver *ErrorOnce) IsNullOrEmpty() bool {
	err := receiver.Value()

	return err == nil || err.Error() == ""
}

func (receiver *ErrorOnce) Message() string {
	if receiver.IsNull() {
		return constants.EmptyString
	}

	return receiver.Value().Error()
}

func (receiver *ErrorOnce) IsMessageEqual(msg string) bool {
	if receiver.IsNull() {
		return false
	}

	return receiver.Message() == msg
}

// HandleError with panic if error exist or else skip
//
// Skip if no error type (NoError).
func (receiver *ErrorOnce) HandleError() {
	if receiver.IsNullOrEmpty() {
		return
	}

	panic(receiver.Value())
}

// HandleErrorWith by concatenating message and then panic if error exist or else skip
//
// Skip if no error type (NoError).
func (receiver *ErrorOnce) HandleErrorWith(messages ...string) {
	if receiver.IsNullOrEmpty() {
		return
	}

	panic(receiver.ConcatNewString(messages...))
}

func (receiver *ErrorOnce) ConcatNewString(messages ...string) string {
	additionalMessages :=
		converters.StringsToCsv(
			&messages,
			false)

	if receiver.IsNullOrEmpty() {
		return additionalMessages
	}

	return receiver.Value().Error() +
		constants.NewLineUnix +
		additionalMessages
}

func (receiver *ErrorOnce) ConcatNew(messages ...string) error {
	return errors.New(receiver.ConcatNewString(messages...))
}

func (receiver *ErrorOnce) Value() error {
	if receiver.isInitialized.IsTrue() {
		return receiver.innerData
	}

	receiver.innerData = receiver.initializerFunc()
	receiver.isInitialized = issetter.True

	return receiver.innerData
}

func (receiver *ErrorOnce) String() string {
	return receiver.Value().Error()
}
