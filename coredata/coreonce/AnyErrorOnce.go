package coreonce

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type AnyErrorOnce struct {
	innerData       interface{}
	err             error
	initializerFunc func() (interface{}, error)
	compiledString  *string
	isInitialized   bool
}

func NewAnyErrorOnce(initializerFunc func() (interface{}, error)) AnyErrorOnce {
	return AnyErrorOnce{
		initializerFunc: initializerFunc,
	}
}

func NewAnyErrorOncePtr(initializerFunc func() (interface{}, error)) *AnyErrorOnce {
	return &AnyErrorOnce{
		initializerFunc: initializerFunc,
	}
}

// Error
//
//	Runs the execution and returns the error.
func (it *AnyErrorOnce) Error() error {
	if it.isInitialized {
		return it.err
	}

	_, err := it.Value()

	return err
}

func (it *AnyErrorOnce) HasError() bool {
	return it.Error() != nil
}

func (it *AnyErrorOnce) IsEmptyError() bool {
	return it.Error() == nil
}

// IsEmpty
//
//	represent values and error both empty
func (it *AnyErrorOnce) IsEmpty() bool {
	return it == nil || it.IsNull() && it.IsEmptyError()
}

func (it *AnyErrorOnce) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *AnyErrorOnce) IsDefined() bool {
	return !it.IsEmpty()
}

// IsInvalid
//
//	represents has error
func (it *AnyErrorOnce) IsInvalid() bool {
	return it.HasError()
}

// IsValid
//
//	represents empty error
func (it *AnyErrorOnce) IsValid() bool {
	return it.IsEmptyError()
}

// IsSuccess
//
//	represents empty error
func (it *AnyErrorOnce) IsSuccess() bool {
	return it.IsEmptyError()
}

// IsFailed
//
//	represents has error
func (it *AnyErrorOnce) IsFailed() bool {
	return !it.IsEmptyError()
}

func (it *AnyErrorOnce) ValueWithError() (interface{}, error) {
	return it.Value()
}

func (it *AnyErrorOnce) Execute() (interface{}, error) {
	return it.Value()
}

func (it *AnyErrorOnce) ExecuteMust() interface{} {
	val, err := it.Value()

	if err != nil {
		panic(err)
	}

	return val
}

func (it *AnyErrorOnce) ValueMust() interface{} {
	val, err := it.Value()

	if err != nil {
		panic(err)
	}

	return val
}

func (it *AnyErrorOnce) Value() (interface{}, error) {
	if it.isInitialized {
		return it.innerData, it.err
	}

	it.innerData, it.err = it.initializerFunc()
	it.isInitialized = true

	return it.innerData, it.err
}

// ValueStringOnly
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
func (it *AnyErrorOnce) ValueStringOnly() (val string) {
	val, _ = it.ValueString()

	return val
}

// SafeString
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
func (it *AnyErrorOnce) SafeString() (val string) {
	return it.ValueStringOnly()
}

// ValueStringMust
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
//
//	Panic if error exist.
func (it *AnyErrorOnce) ValueStringMust() (val string) {
	val, err := it.ValueString()

	if err != nil {
		panic(err)
	}

	return val
}

// ValueString
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
func (it *AnyErrorOnce) ValueString() (val string, err error) {
	if it.compiledString != nil {
		return *it.compiledString, it.err
	}

	valInf, err := it.Value()

	if valInf == nil {
		return constants.NilAngelBracket, err
	}

	toString := fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		valInf)

	it.compiledString = &toString

	return *it.compiledString, err
}

func (it *AnyErrorOnce) CastValueString() (
	val string, err error, isSuccess bool,
) {
	valInf, err := it.Execute()
	toString, isSuccess := valInf.(string)

	return toString, err, isSuccess
}

func (it *AnyErrorOnce) CastValueStrings() (
	valueStrings []string, err error, isSuccess bool,
) {
	valInf, err := it.Execute()
	toStrings, isSuccess := valInf.([]string)

	return toStrings, err, isSuccess
}

func (it *AnyErrorOnce) CastValueHashmapMap() (
	valueMap map[string]string,
	err error,
	isSuccess bool,
) {
	valInf, err := it.Execute()
	toStrings, isSuccess := valInf.(map[string]string)

	return toStrings, err, isSuccess
}

func (it *AnyErrorOnce) CastValueMapStringAnyMap() (
	valueMap map[string]interface{},
	err error,
	isSuccess bool,
) {
	valInf, err := it.Execute()
	toStrings, isSuccess := valInf.(map[string]interface{})

	return toStrings, err, isSuccess
}

func (it *AnyErrorOnce) CastValueBytes() (
	rawBytes []byte,
	err error,
	isSuccess bool,
) {
	valInf, err := it.Execute()
	toStrings, isSuccess := valInf.([]byte)

	return toStrings, err, isSuccess
}

func (it *AnyErrorOnce) ValueOnly() interface{} {
	if it.isInitialized {
		return it.innerData
	}

	val, _ := it.Value()

	return val
}

func (it *AnyErrorOnce) IsInitialized() bool {
	return it.isInitialized
}

func (it *AnyErrorOnce) IsNull() bool {
	return it.ValueOnly() == nil
}

func (it *AnyErrorOnce) IsStringEmpty() bool {
	return it.String() == ""
}

func (it *AnyErrorOnce) IsStringEmptyOrWhitespace() bool {
	return strings.TrimSpace(it.String()) == ""
}

func (it *AnyErrorOnce) String() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.ValueOnly())
}

func (it *AnyErrorOnce) Deserialize(toPtr interface{}) error {
	allBytes, err := it.Serialize()

	if err != nil {
		return err
	}

	unmarshallErr := json.Unmarshal(allBytes, toPtr)

	if err == nil {
		return err
	}

	var safeString string
	if len(allBytes) > 0 {
		safeString = string(allBytes)
	}

	var typeSafeName string
	if toPtr != nil {
		typeSafeName = reflect.TypeOf(toPtr).String()
	}

	message :=
		"deserializing failed: " + unmarshallErr.Error() +
			", json payload:" + safeString +
			", type: " + typeSafeName

	// has err
	return errors.New(message)
}

func (it *AnyErrorOnce) Serialize() ([]byte, error) {
	value, err := it.Value()

	if err != nil {
		return nil, errors.New(
			"cannot serialize on existing error, " + err.Error() +
				"value string : " + it.SafeString())
	}

	allBytes, marshalErr := json.Marshal(value)

	if marshalErr == nil {
		return allBytes, nil
	}

	return nil, errors.New(
		"unmarshalling error, " + marshalErr.Error() +
			"value string : " + it.SafeString())
}

func (it *AnyErrorOnce) SerializeSkipExistingError() ([]byte, error) {
	return json.Marshal(it.ValueOnly())
}

func (it *AnyErrorOnce) SerializeMust() []byte {
	allBytes, err := it.Serialize()

	if err != nil {
		panic(err)
	}

	return allBytes
}
