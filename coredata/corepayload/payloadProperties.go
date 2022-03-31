package corepayload

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/payloadinf"
)

type payloadProperties struct {
	payloadWrapper *PayloadWrapper
}

func (it *payloadProperties) SetBasicError(basicError errcoreinf.BasicErrWrapper) {
	it.payloadWrapper.InitializeAttributesOnNull()
	it.payloadWrapper.Attributes.SetBasicErr(basicError)
}

func (it *payloadProperties) BasicError() errcoreinf.BasicErrWrapper {
	return it.payloadWrapper.BasicError()
}

func (it *payloadProperties) ReflectSetTo(toPointer interface{}) error {
	return coredynamic.ReflectSetFromTo(it.payloadWrapper, toPointer)
}

func (it payloadProperties) AllSafe() (id, name, entity, category string, dynamicPayloads []byte) {
	return it.payloadWrapper.AllSafe()
}

func (it payloadProperties) All() (id, name, entity, category string, dynamicPayloads []byte) {
	return it.payloadWrapper.All()
}

func (it payloadProperties) Name() string {
	return it.payloadWrapper.Name
}

func (it *payloadProperties) SetName(name string) error {
	it.payloadWrapper.Name = name

	return nil
}

func (it *payloadProperties) SetNameMust(name string) {
	it.payloadWrapper.Name = name
}

func (it *payloadProperties) IdInteger() int {
	return it.payloadWrapper.IdInteger()
}

func (it *payloadProperties) IdUnsignedInteger() uint {
	return it.payloadWrapper.IdentifierUnsignedInteger()
}

func (it payloadProperties) IdString() string {
	return it.payloadWrapper.Identifier
}

func (it *payloadProperties) SetIdString(id string) error {
	it.payloadWrapper.Identifier = id

	return nil
}

func (it *payloadProperties) SetIdStringMust(id string) {
	it.payloadWrapper.Identifier = id
}

func (it payloadProperties) Category() string {
	return it.payloadWrapper.CategoryName
}

func (it *payloadProperties) SetCategory(category string) error {
	it.payloadWrapper.CategoryName = category

	return nil
}

func (it *payloadProperties) SetCategoryMust(category string) {
	it.payloadWrapper.CategoryName = category
}

func (it payloadProperties) EntityType() string {
	return it.payloadWrapper.EntityType
}

func (it *payloadProperties) SetEntityType(entityName string) error {
	it.payloadWrapper.EntityType = entityName

	return nil
}

func (it *payloadProperties) SetEntityTypeMust(entityName string) {
	it.payloadWrapper.EntityType = entityName
}

func (it *payloadProperties) HasManyRecord() bool {
	return it.payloadWrapper.HasManyRecords
}

func (it payloadProperties) HasSingleRecordOnly() bool {
	return !it.HasManyRecord()
}

func (it *payloadProperties) SetSingleRecordFlag() {
	it.payloadWrapper.HasManyRecords = false
}

func (it *payloadProperties) SetManyRecordFlag() {
	it.payloadWrapper.HasManyRecords = true
}

func (it payloadProperties) DynamicPayloads() []byte {
	return it.payloadWrapper.Payloads
}

func (it *payloadProperties) SetDynamicPayloads(dynamicPayloads []byte) error {
	it.payloadWrapper.Payloads = dynamicPayloads

	return nil
}

func (it *payloadProperties) DynamicPayloadsDeserializedTo(toPtr interface{}) error {
	return it.payloadWrapper.Deserialize(toPtr)
}

func (it *payloadProperties) SetDynamicPayloadsMust(dynamicPayloads []byte) {
	it.payloadWrapper.Payloads = dynamicPayloads
}

func (it payloadProperties) Json() corejson.Result {
	return it.payloadWrapper.Json()
}

func (it payloadProperties) JsonPtr() *corejson.Result {
	return it.payloadWrapper.JsonPtr()
}

func (it payloadProperties) AsPayloadPropertiesDefiner() payloadinf.PayloadPropertiesDefiner {
	return &it
}
