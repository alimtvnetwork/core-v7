package corepayload

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type PayloadWrapper struct {
	Name, Identifier         string
	TaskTypeName, EntityType string

	CategoryName   string
	HasManyRecords bool
	Payloads       []byte
	Attributes     *Attributes `json:"Attributes,omitempty"`
}

func (it *PayloadWrapper) JsonString() string {
	return it.Json().JsonString()
}

func (it *PayloadWrapper) JsonStringMust() string {
	return it.Json().JsonString()
}

func (it *PayloadWrapper) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *PayloadWrapper) HasError() bool {
	return it != nil && it.Attributes.HasError()
}

func (it *PayloadWrapper) IsEmptyError() bool {
	return it == nil || it.Attributes.IsEmptyError()
}

func (it *PayloadWrapper) HasAttributes() bool {
	return it != nil && it.Attributes != nil
}

func (it *PayloadWrapper) IsEmptyAttributes() bool {
	return it == nil || it.Attributes == nil
}

func (it *PayloadWrapper) HasSingleRecord() bool {
	return it != nil && !it.HasManyRecords
}

func (it *PayloadWrapper) IsNull() bool {
	return it == nil
}

func (it *PayloadWrapper) HasAnyNil() bool {
	return it == nil
}

func (it *PayloadWrapper) Count() int {
	return it.Length()
}

func (it *PayloadWrapper) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Payloads)
}

func (it *PayloadWrapper) IsEmpty() bool {
	return it.Length() == 0
}

func (it *PayloadWrapper) HasItems() bool {
	return it.Length() > 0
}

func (it *PayloadWrapper) BytesConverter() *coredynamic.BytesConverter {
	return coredynamic.NewBytesConverter(it.Payloads)
}

func (it *PayloadWrapper) Deserialize(
	unmarshallingPointer interface{},
) error {
	return corejson.
		Deserialize.
		UsingBytes(
			it.Payloads,
			unmarshallingPointer)
}

func (it *PayloadWrapper) DeserializeMust(
	unmarshallingPointer interface{},
) {
	corejson.
		Deserialize.
		UsingBytesMust(
			it.Payloads,
			unmarshallingPointer)
}

func (it *PayloadWrapper) PayloadDeserialize(
	unmarshallingPointer interface{},
) error {
	return corejson.Deserialize.UsingBytes(
		it.Payloads,
		unmarshallingPointer)
}

func (it *PayloadWrapper) PayloadDeserializeMust(
	unmarshallingPointer interface{},
) {
	err := corejson.Deserialize.UsingBytes(
		it.Payloads,
		unmarshallingPointer)

	if err != nil {
		panic(err)
	}
}

func (it *PayloadWrapper) DeserializePayloadsToPayloadWrapper() (
	payloadWrapper *PayloadWrapper, err error,
) {
	payloadWrapper = New.PayloadWrapper.Empty()
	err = corejson.Deserialize.UsingBytes(
		it.Payloads,
		payloadWrapper)

	if err != nil {
		payloadWrapper.Attributes.AttachOrAppendErrorMessage(
			err.Error())
	}

	return payloadWrapper, err
}

func (it PayloadWrapper) JsonModel() PayloadWrapper {
	return it
}

func (it *PayloadWrapper) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it PayloadWrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it PayloadWrapper) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *PayloadWrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*PayloadWrapper, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *PayloadWrapper) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *PayloadWrapper {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *PayloadWrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *PayloadWrapper) Clear() *PayloadWrapper {
	if it == nil {
		return nil
	}

	it.Payloads = it.Payloads[:0]

	return it
}

func (it *PayloadWrapper) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Attributes = nil
}

func (it *PayloadWrapper) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}
