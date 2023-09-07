package coreonce

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type BytesErrorOnce struct {
	innerData       []byte
	err             error
	initializerFunc func() ([]byte, error)
	isInitialized   bool
}

func NewBytesErrorOnce(initializerFunc func() ([]byte, error)) BytesErrorOnce {
	return BytesErrorOnce{
		initializerFunc: initializerFunc,
	}
}

func NewBytesErrorOncePtr(initializerFunc func() ([]byte, error)) *BytesErrorOnce {
	return &BytesErrorOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *BytesErrorOnce) HandleError() {
	err := it.Error()

	if err != nil {
		panic(err)
	}
}

func (it *BytesErrorOnce) MustBeEmptyError() {
	err := it.Error()

	if err != nil {
		panic(err)
	}
}

func (it *BytesErrorOnce) MustHaveSafeItems() {
	err := it.Error()

	if err != nil {
		panic(err)
	}

	if it.IsBytesEmpty() {
		panic("values cannot be null or empty!")
	}
}

// Error
//
//  Runs the execution and returns the error.
func (it *BytesErrorOnce) Error() error {
	if it.isInitialized {
		return it.err
	}

	_, err := it.Value()

	return err
}

func (it *BytesErrorOnce) HasError() bool {
	return it.Error() != nil
}

func (it *BytesErrorOnce) HasIssuesOrEmpty() bool {
	if it == nil {
		return true
	}

	val, err := it.Value()

	return err != nil || len(val) == 0
}

func (it *BytesErrorOnce) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *BytesErrorOnce) IsEmptyError() bool {
	return it.Error() == nil
}

// IsEmpty
//
//  represent values and error both empty
func (it *BytesErrorOnce) IsEmpty() bool {
	return it == nil || it.IsNull() && it.IsEmptyError()
}

// IsEmptyBytes
//
//  represent values
func (it *BytesErrorOnce) IsEmptyBytes() bool {
	return it == nil || it.Length() == 0
}

func (it *BytesErrorOnce) Length() int {
	return len(it.ValueOnly())
}

func (it *BytesErrorOnce) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *BytesErrorOnce) IsDefined() bool {
	return !it.IsEmpty()
}

// IsInvalid
//
//  represents has error
func (it *BytesErrorOnce) IsInvalid() bool {
	return it.HasError()
}

// IsValid
//
//  represents empty error
func (it *BytesErrorOnce) IsValid() bool {
	return it.IsEmptyError()
}

// IsSuccess
//
//  represents empty error
func (it *BytesErrorOnce) IsSuccess() bool {
	return it.IsEmptyError()
}

// IsFailed
//
//  represents has error
func (it *BytesErrorOnce) IsFailed() bool {
	return !it.IsEmptyError()
}

func (it *BytesErrorOnce) Execute() ([]byte, error) {
	return it.Value()
}

func (it *BytesErrorOnce) Deserialize(
	toPtr interface{},
) error {
	rawBytes, err := it.Value()
	var valString string

	if len(rawBytes) > 0 {
		valString = string(rawBytes)
	}

	var typeNameString string

	if toPtr != nil {
		// actual nil-ness
		// not required
		typeNameString =
			reflect.
				TypeOf(toPtr).
				String()
	}

	if err != nil {
		return errors.New(
			"existing error cannot deserialize, " +
				err.Error() +
				", payload : " + valString + "," +
				" to type:" + typeNameString)
	}

	jsonUnmarshalErr := json.Unmarshal(rawBytes, toPtr)

	if err == nil {
		return nil
	}

	// has error
	return errors.New(
		"deserialize failed, " +
			jsonUnmarshalErr.Error() +
			", payload : " + valString + "," +
			" to type:" + typeNameString)
}

func (it *BytesErrorOnce) DeserializeMust(
	toPtr interface{},
) {
	err := it.Deserialize(toPtr)

	if err != nil {
		panic(err)
	}
}

func (it *BytesErrorOnce) ValueWithError() ([]byte, error) {
	return it.Value()
}

func (it *BytesErrorOnce) Value() ([]byte, error) {
	if it.isInitialized {
		return it.innerData, it.err
	}

	it.innerData, it.err = it.initializerFunc()
	it.isInitialized = true

	return it.innerData, it.err
}

func (it *BytesErrorOnce) ValueOnly() []byte {
	if it.isInitialized {
		return it.innerData
	}

	val, _ := it.Value()

	return val
}

func (it *BytesErrorOnce) IsInitialized() bool {
	return it.isInitialized
}

func (it *BytesErrorOnce) IsBytesEmpty() bool {
	return it.Length() == 0
}

func (it *BytesErrorOnce) IsNull() bool {
	return it.ValueOnly() == nil
}

func (it *BytesErrorOnce) IsStringEmpty() bool {
	return it.String() == ""
}

func (it *BytesErrorOnce) IsStringEmptyOrWhitespace() bool {
	return strings.TrimSpace(it.String()) == ""
}

func (it *BytesErrorOnce) String() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return string(it.ValueOnly())
}

func (it *BytesErrorOnce) MarshalJSON() ([]byte, error) {
	return it.Value()
}

func (it *BytesErrorOnce) Serialize() ([]byte, error) {
	return it.Value()
}

func (it *BytesErrorOnce) SerializeMust() []byte {
	allBytes, err := it.Serialize()

	if err != nil {
		panic(err)
	}

	return allBytes
}
