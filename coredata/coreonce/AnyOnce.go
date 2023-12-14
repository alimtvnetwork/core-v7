package coreonce

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type AnyOnce struct {
	innerData       interface{}
	initializerFunc func() interface{}
	compiledString  *string
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

func (it *AnyOnce) Execute() interface{} {
	return it.Value()
}

// ValueStringOnly
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
func (it *AnyOnce) ValueStringOnly() (val string) {
	return it.ValueString()
}

// SafeString
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
func (it *AnyOnce) SafeString() (val string) {
	return it.ValueStringOnly()
}

// ValueStringMust
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
//
//	Panic if error exist.
func (it *AnyOnce) ValueStringMust() (val string) {
	return it.ValueString()
}

// ValueString
//
//	Usages SPrintf to get the string,
//	mostly use the String() func to get the value
func (it *AnyOnce) ValueString() (val string) {
	if it.compiledString != nil {
		return *it.compiledString
	}

	valInf := it.Value()

	if valInf == nil {
		return constants.NilAngelBracket
	}

	toString := fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		valInf)

	it.compiledString = &toString

	return *it.compiledString
}

func (it *AnyOnce) CastValueString() (
	val string, isSuccess bool,
) {
	valInf := it.Execute()
	toString, isSuccess := valInf.(string)

	return toString, isSuccess
}

func (it *AnyOnce) CastValueStrings() (
	valueStrings []string, isSuccess bool,
) {
	valInf := it.Execute()
	toStrings, isSuccess := valInf.([]string)

	return toStrings, isSuccess
}

func (it *AnyOnce) CastValueHashmapMap() (
	valueMap map[string]string,
	isSuccess bool,
) {
	valInf := it.Execute()
	toStrings, isSuccess := valInf.(map[string]string)

	return toStrings, isSuccess
}

func (it *AnyOnce) CastValueMapStringAnyMap() (
	valueMap map[string]interface{},
	isSuccess bool,
) {
	valInf := it.Execute()
	toStrings, isSuccess := valInf.(map[string]interface{})

	return toStrings, isSuccess
}

func (it *AnyOnce) CastValueBytes() (
	rawBytes []byte,
	isSuccess bool,
) {
	valInf := it.Execute()
	toStrings, isSuccess := valInf.([]byte)

	return toStrings, isSuccess
}

func (it *AnyOnce) ValueOnly() interface{} {
	return it.Value()
}

func (it *AnyOnce) IsInitialized() bool {
	return it.isInitialized
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

func (it *AnyOnce) Deserialize(toPtr interface{}) error {
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

func (it *AnyOnce) Serialize() ([]byte, error) {
	value := it.Value()
	allBytes, marshalErr := json.Marshal(value)

	if marshalErr == nil {
		return allBytes, nil
	}

	return nil, errors.New(
		"unmarshalling error, " + marshalErr.Error() +
			"value string : " + it.SafeString())
}

func (it *AnyOnce) SerializeSkipExistingError() ([]byte, error) {
	return it.Serialize()
}

func (it *AnyOnce) SerializeMust() []byte {
	value := it.Value()

	jsonByes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	return jsonByes
}
