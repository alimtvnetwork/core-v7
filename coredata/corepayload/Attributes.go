package corepayload

import (
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/defaulterr"
)

type Attributes struct {
	ErrorMessage     string                   `json:"ErrorMessage,omitempty"`
	PagingInfo       *PagingInfo              `json:"PagingInfo,omitempty"`
	KeyValuePairs    *corestr.Hashmap         `json:"KeyValuePairs,omitempty"`
	AnyKeyValuePairs *coredynamic.MapAnyItems `json:"AnyKeyValuePairs,omitempty"`
	DynamicPayloads  []byte                   `json:"DynamicPayloads,omitempty"`
}

func (it *Attributes) JsonString() string {
	return it.Json().JsonString()
}

func (it *Attributes) JsonStringMust() string {
	return it.Json().JsonString()
}

func (it *Attributes) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *Attributes) Count() int {
	return it.Length()
}

func (it *Attributes) Capacity() int {
	return it.Length()
}

func (it *Attributes) Length() int {
	if it == nil {
		return 0
	}

	return len(it.DynamicPayloads)
}

func (it *Attributes) HasPagingInfo() bool {
	return it != nil && it.PagingInfo != nil
}

func (it *Attributes) HasKeyValuePairs() bool {
	return it != nil && it.KeyValuePairs.HasAnyItem()
}

func (it *Attributes) IsValid() bool {
	return it != nil && it.ErrorMessage == ""
}

func (it *Attributes) IsInvalid() bool {
	return it == nil || it.ErrorMessage != ""
}

func (it *Attributes) HasError() bool {
	return it != nil && it.ErrorMessage != ""
}

func (it *Attributes) Error() error {
	if it.IsEmptyError() {
		return nil
	}

	return errors.New(it.ErrorMessage)
}

func (it *Attributes) MustBeEmptyError() {
	if it.IsEmptyError() {
		return
	}

	panic(it.ErrorMessage)
}

// DeserializeErrorMessage
//
// Expectation Attributes.ErrorMessage needs to
// be in json format and unmarshalToPointer
// should match reflection types
func (it *Attributes) DeserializeErrorMessage(
	unmarshalToPointer interface{},
) error {
	if it.IsEmptyError() {
		return nil
	}

	return corejson.
		Deserialize.
		UsingString(
			it.ErrorMessage,
			unmarshalToPointer)
}

func (it *Attributes) DeserializeDynamicPayloads(
	unmarshalToPointer interface{},
) error {
	return corejson.
		Deserialize.
		UsingBytes(
			it.DynamicPayloads,
			unmarshalToPointer)
}

func (it *Attributes) DeserializeDynamicPayloadsToAttributes() (newAttr *Attributes, err error) {
	newAttr = New.Attributes.Empty()
	err = corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		newAttr)

	if err != nil {
		newAttr.AttachOrAppendErrorMessage(err.Error())
	}

	return newAttr, err
}

func (it *Attributes) DeserializeDynamicPayloadsToPayloadWrapper() (payloadWrapper *PayloadWrapper, err error) {
	payloadWrapper = New.PayloadWrapper.Empty()
	err = corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		payloadWrapper)

	if err != nil {
		payloadWrapper.Attributes.AttachOrAppendErrorMessage(
			err.Error())
	}

	return payloadWrapper, err
}

func (it *Attributes) DeserializeDynamicPayloadsMust(
	unmarshalToPointer interface{},
) {
	corejson.Deserialize.
		UsingBytesMust(
			it.DynamicPayloads,
			unmarshalToPointer)
}

func (it *Attributes) IsEmptyError() bool {
	return it == nil || it.ErrorMessage == ""
}

func (it *Attributes) DynamicBytesLength() int {
	if it == nil {
		return 0
	}

	return len(it.DynamicPayloads)
}

func (it *Attributes) StringKeyValuePairsLength() int {
	if it == nil {
		return 0
	}

	return it.KeyValuePairs.Length()
}

func (it *Attributes) AnyKeyValuePairsLength() int {
	if it == nil {
		return 0
	}

	return it.AnyKeyValuePairs.Length()
}

func (it *Attributes) IsEmpty() bool {
	return it == nil ||
		it.DynamicBytesLength() == 0 &&
			it.StringKeyValuePairsLength() == 0 &&
			it.AnyKeyValuePairsLength() == 0
}

func (it *Attributes) HasItems() bool {
	return !it.IsEmpty()
}

func (it *Attributes) HasStringKeyValuePairs() bool {
	return it.StringKeyValuePairsLength() > 0
}

func (it *Attributes) HasAnyKeyValuePairs() bool {
	return it.AnyKeyValuePairsLength() > 0
}

func (it *Attributes) HasDynamicPayloads() bool {
	return it.DynamicBytesLength() > 0
}

func (it *Attributes) DynamicPayloadsDeserialize(
	unmarshallingPointer interface{},
) error {
	if it == nil {
		return defaulterr.AttributeNull
	}

	return corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		unmarshallingPointer)
}

func (it *Attributes) DynamicPayloadsDeserializeMust(
	unmarshallingPointer interface{},
) {
	err := corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		unmarshallingPointer)

	if err != nil {
		panic(err)
	}
}

func (it *Attributes) AddOrUpdateString(
	key, value string,
) (isNewlyAdded bool) {
	return it.
		KeyValuePairs.
		AddOrUpdate(key, value)
}

func (it *Attributes) AddOrUpdateAnyItem(
	key string,
	anyItem interface{},
) (isNewlyAdded bool) {
	return it.
		AnyKeyValuePairs.
		Add(key, anyItem)
}

func (it *Attributes) JsonModel() *Attributes {
	return it
}

func (it *Attributes) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Attributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Attributes) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, it)

	return err
}

func (it Attributes) Json() corejson.Result {
	return corejson.New(it)
}

func (it Attributes) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Attributes) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Attributes, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return &Attributes{}, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *Attributes) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Attributes {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Attributes) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Attributes) AttachOrAppendError(
	err error,
) *Attributes {
	if err == nil {
		return it
	}

	return it.AttachOrAppendErrorMessage(
		err.Error())
}

func (it *Attributes) AttachOrAppendErrorMessage(
	appendingNewErrorMessage string,
) *Attributes {
	if it == nil {
		return New.
			Attributes.
			UsingErrMsg(appendingNewErrorMessage)
	}

	if it.IsEmptyError() {
		it.ErrorMessage = appendingNewErrorMessage

		return it
	}

	it.ErrorMessage = fmt.Sprintf(
		constants.SprintFormatWithNewLine,
		it.ErrorMessage,
		appendingNewErrorMessage)

	return it
}

func (it *Attributes) Clear() {
	if it == nil {
		return
	}

	it.KeyValuePairs.Clear()
	it.AnyKeyValuePairs.Clear()
	it.DynamicPayloads = []byte{}
}

func (it *Attributes) Dispose() {
	it.Clear()
}

func (it *Attributes) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Attributes) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}
