package coreonce

import (
	"encoding/json"
	"errors"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
)

type ErrorOnce struct {
	innerData       error
	initializerFunc func() error
	isInitialized   bool
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

func (it *ErrorOnce) MarshalJSON() ([]byte, error) {
	if it.IsNullOrEmpty() {
		return json.Marshal("")
	}

	return json.Marshal(it.Value().Error())
}

func (it *ErrorOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true
	var str string

	err := json.Unmarshal(data, &str)
	it.innerData = errors.New(str)

	return err
}

func (it *ErrorOnce) HasError() bool {
	return !it.IsNullOrEmpty()
}

func (it *ErrorOnce) IsEmpty() bool {
	return it.IsNullOrEmpty()
}

func (it *ErrorOnce) IsEmptyError() bool {
	return it.IsNullOrEmpty()
}

func (it *ErrorOnce) HasAnyItem() bool {
	return !it.IsNullOrEmpty()
}

func (it *ErrorOnce) IsDefined() bool {
	return !it.IsNullOrEmpty()
}

// IsInvalid
//
//  represents has error
func (it *ErrorOnce) IsInvalid() bool {
	return !it.IsNullOrEmpty()
}

// IsValid
//
//  represents empty error
func (it *ErrorOnce) IsValid() bool {
	return it.IsNullOrEmpty()
}

// IsSuccess
//
//  represents empty error
func (it *ErrorOnce) IsSuccess() bool {
	return it.IsNullOrEmpty()
}

// IsFailed
//
//  represents has error
func (it *ErrorOnce) IsFailed() bool {
	return !it.IsNullOrEmpty()
}

func (it *ErrorOnce) IsNull() bool {
	return it.Value() == nil
}

func (it *ErrorOnce) IsNullOrEmpty() bool {
	err := it.Value()

	return err == nil || err.Error() == ""
}

func (it *ErrorOnce) Message() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return it.Value().Error()
}

func (it *ErrorOnce) IsMessageEqual(msg string) bool {
	if it.IsNull() {
		return false
	}

	return it.Message() == msg
}

// HandleError with panic if error exist or else skip
//
// Skip if no error type (NoError).
func (it *ErrorOnce) HandleError() {
	if it.IsNullOrEmpty() {
		return
	}

	panic(it.Value())
}

// HandleErrorWith by concatenating message and then panic if error exist or else skip
//
// Skip if no error type (NoError).
func (it *ErrorOnce) HandleErrorWith(messages ...string) {
	if it.IsNullOrEmpty() {
		return
	}

	panic(it.ConcatNewString(messages...))
}

func (it *ErrorOnce) ConcatNewString(messages ...string) string {
	additionalMessages :=
		converters.StringsTo.Csv(
			false,
			messages...,
		)

	if it.IsNullOrEmpty() {
		return additionalMessages
	}

	return it.Value().Error() +
		constants.NewLineUnix +
		additionalMessages
}

func (it *ErrorOnce) ConcatNew(messages ...string) error {
	return errors.New(it.ConcatNewString(messages...))
}

func (it *ErrorOnce) Value() error {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *ErrorOnce) Execute() error {
	return it.Value()
}

func (it *ErrorOnce) String() string {
	return it.Value().Error()
}

func (it *ErrorOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
