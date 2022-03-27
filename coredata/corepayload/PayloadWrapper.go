package corepayload

import (
	"bytes"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/entityinf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/payloadinf"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/errcore"
)

type PayloadWrapper struct {
	Name           string      `json:"Name,omitempty"`
	Identifier     string      `json:"Identifier,omitempty"`
	TaskTypeName   string      `json:"TaskTypeName,omitempty"`
	EntityType     string      `json:"EntityType,omitempty"`
	CategoryName   string      `json:"CategoryName,omitempty"`
	HasManyRecords bool        `json:"HasManyRecords,omitempty"`
	Payloads       []byte      `json:"Payloads,omitempty"`
	Attributes     *Attributes `json:"AnyAttributes,omitempty"`
}

func (it *PayloadWrapper) HasSafeItems() bool {
	if it.IsEmpty() || it.HasError() {
		return false
	}

	return true
}

func (it *PayloadWrapper) DynamicPayloads() []byte {
	if it == nil {
		return []byte{}
	}

	return it.Payloads
}

func (it *PayloadWrapper) SetDynamicPayloads(payloads []byte) error {
	if it == nil {
		return defaulterr.NilResult
	}

	it.Payloads = payloads

	return nil
}

func (it *PayloadWrapper) AttrAsBinder() payloadinf.AttributesBinder {
	return it.Attributes
}

func (it *PayloadWrapper) InitializeAttributesOnNull() payloadinf.AttributesBinder {
	if it.Attributes == nil {
		it.Attributes = New.Attributes.Empty()
	}

	return it.Attributes
}

func (it *PayloadWrapper) BasicError() errcoreinf.BasicErrWrapper {
	if it.HasError() {
		return it.Attributes.BasicErrWrapper
	}

	return nil
}

func (it *PayloadWrapper) PayloadDeserializeToPayloadBinder() (payloadinf.PayloadsBinder, error) {
	if it.IsNull() {
		return nil, defaulterr.NilResult
	}

	if it.HasError() {
		return nil, it.Attributes.BasicErrWrapper.CompiledError()
	}

	return it.DeserializePayloadsToPayloadWrapper()
}

func (it *PayloadWrapper) All() (id, name, entity, category string, dynamicPayloads []byte) {
	return it.Identifier, it.Name, it.EntityType, it.CategoryName, it.Payloads
}

func (it *PayloadWrapper) AllSafe() (id, name, entity, category string, dynamicPayloads []byte) {
	if it.IsNull() {
		return "", "", "", "", []byte{}
	}

	return it.All()
}

func (it *PayloadWrapper) PayloadName() string {
	return it.Name
}

func (it *PayloadWrapper) PayloadCategory() string {
	return it.CategoryName
}

func (it *PayloadWrapper) PayloadTaskType() string {
	return it.TaskTypeName
}

func (it *PayloadWrapper) PayloadEntityType() string {
	return it.EntityType
}

func (it *PayloadWrapper) PayloadDynamic() []byte {
	return it.Payloads
}

func (it *PayloadWrapper) PayloadProperties() payloadinf.PayloadPropertiesDefiner {
	return &payloadProperties{it}
}

func (it *PayloadWrapper) HandleError() {
	if it.HasError() {
		it.BasicError().HandleError()
	}
}

func (it *PayloadWrapper) ReflectSetTo(
	toPointer interface{},
) error {
	return coredynamic.ReflectSetFromTo(
		it,
		toPointer)
}

func (it *PayloadWrapper) AnyAttributes() interface{} {
	return it.Attributes
}

func (it *PayloadWrapper) ReflectSetAttributes(
	toPointer interface{},
) error {
	return coredynamic.ReflectSetFromTo(
		it.Attributes,
		toPointer)
}

func (it *PayloadWrapper) IdString() string {
	return it.Identifier
}

func (it *PayloadWrapper) IdInteger() int {
	return it.IdentifierInteger()
}

func (it *PayloadWrapper) IsStandardTaskEntityEqual(
	entity entityinf.StandardTaskEntityDefiner,
) bool {
	another, isSuccess := entity.(*PayloadWrapper)

	if !isSuccess {
		return false
	}

	return it.IsEqual(another)
}

func (it *PayloadWrapper) ValueReflectSet(
	setterPtr interface{},
) error {
	return coredynamic.ReflectSetFromTo(
		it.Payloads,
		setterPtr)
}

func (it *PayloadWrapper) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *PayloadWrapper) SerializeMust() []byte {
	json := it.Json()
	json.HandleError()

	return json.Bytes
}

func (it *PayloadWrapper) Username() string {
	if it.IsEmptyAttributes() {
		return ""
	}

	virtualUser := it.Attributes.VirtualUser()

	if virtualUser == nil {
		return ""
	}

	return virtualUser.Name
}

func (it *PayloadWrapper) Value() interface{} {
	return it.Payloads
}

func (it *PayloadWrapper) Error() error {
	if it.IsEmptyError() {
		return nil
	}

	return it.Attributes.Error()
}

func (it *PayloadWrapper) IsEqual(right *PayloadWrapper) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it == right {
		return true
	}

	if it.Name != right.Name {
		return false
	}
	if it.Identifier != right.Identifier {
		return false
	}

	if it.TaskTypeName != right.TaskTypeName {
		return false
	}

	if it.EntityType != right.EntityType {
		return false
	}

	if it.CategoryName != right.CategoryName {
		return false
	}

	if it.HasManyRecords != right.HasManyRecords {
		return false
	}

	if !bytes.Equal(it.Payloads, right.Payloads) {
		return false
	}

	if !it.Attributes.IsEqual(right.Attributes) {
		return false
	}

	return true
}

func (it *PayloadWrapper) IsPayloadsEqual(nextPayloads []byte) bool {
	return it != nil && bytes.Equal(it.Payloads, nextPayloads)
}

func (it *PayloadWrapper) IsName(name string) bool {
	return it != nil && it.Name == name
}

func (it *PayloadWrapper) IsIdentifier(id string) bool {
	return it != nil && it.Name == id
}

func (it *PayloadWrapper) IsTaskTypeName(taskType string) bool {
	return it != nil && it.TaskTypeName == taskType
}

func (it *PayloadWrapper) IsEntityType(entityType string) bool {
	return it != nil && it.EntityType == entityType
}

func (it *PayloadWrapper) IsCategory(category string) bool {
	return it != nil && it.CategoryName == category
}

func (it PayloadWrapper) String() string {
	return it.JsonString()
}

func (it PayloadWrapper) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
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

func (it *PayloadWrapper) HasIssuesOrEmpty() bool {
	return it == nil ||
		it.
			Attributes.
			HasError() ||
		it.Length() == 0
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

// IdentifierInteger
//
// Invalid value returns constants.InvalidValue
func (it *PayloadWrapper) IdentifierInteger() int {
	if it.Identifier == "" {
		return constants.InvalidValue
	}

	idInt, _ := converters.StringToIntegerWithDefault(
		it.Identifier,
		constants.InvalidValue)

	return idInt
}

// IdentifierUnsignedInteger
//
// Invalid value returns constants.Zero
func (it *PayloadWrapper) IdentifierUnsignedInteger() uint {
	idInt := it.IdentifierInteger()

	if idInt < 0 {
		return constants.Zero
	}

	return uint(idInt)
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
	err := corejson.
		Deserialize.
		UsingBytes(
			it.Payloads,
			unmarshallingPointer)

	if err != nil {
		panic(err)
	}
}

func (it *PayloadWrapper) DeserializePayloadsToPayloadsCollection() (
	payloadsCollection *PayloadsCollection, err error,
) {
	return New.
		PayloadsCollection.
		Deserialize(it.Payloads)
}

func (it *PayloadWrapper) DeserializePayloadsToPayloadWrapper() (
	payloadWrapper *PayloadWrapper, err error,
) {
	return New.
		PayloadWrapper.
		Deserialize(
			it.Payloads)
}

func (it *PayloadWrapper) DeserializePayloadsToPayloadWrapperMust() (
	payloadWrapper *PayloadWrapper,
) {
	rs, err := New.
		PayloadWrapper.
		Deserialize(it.Payloads)
	errcore.HandleErr(err)

	return rs
}

func (it *PayloadWrapper) JsonModel() PayloadWrapper {
	return it.NonPtr()
}

func (it *PayloadWrapper) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *PayloadWrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it *PayloadWrapper) JsonPtr() *corejson.Result {
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

func (it *PayloadWrapper) Clear() {
	if it == nil {
		return
	}

	it.Payloads = it.Payloads[:0]
	it.Attributes.Clear()
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

func (it *PayloadWrapper) Clone(
	isDeepClone bool,
) (PayloadWrapper, error) {
	clonedPtr, err := it.ClonePtr(isDeepClone)

	if err != nil {
		return PayloadWrapper{}, err
	}

	if clonedPtr == nil {
		return PayloadWrapper{}, defaulterr.NilResult
	}

	return clonedPtr.NonPtr(), nil
}

func (it *PayloadWrapper) ClonePtr(
	isDeepClone bool,
) (*PayloadWrapper, error) {
	if it == nil {
		return nil, nil
	}

	attrCloned, err := it.
		Attributes.
		ClonePtr(isDeepClone)

	if err != nil {
		return nil, err
	}

	if isDeepClone {
		return &PayloadWrapper{
			Name:           it.Name,
			Identifier:     it.Identifier,
			TaskTypeName:   it.TaskTypeName,
			EntityType:     it.EntityType,
			CategoryName:   it.CategoryName,
			HasManyRecords: it.HasManyRecords,
			Payloads: corejson.BytesDeepClone(
				it.Payloads),
			Attributes: attrCloned,
		}, nil
	}

	// NOT deep clone
	return &PayloadWrapper{
		Name:           it.Name,
		Identifier:     it.Identifier,
		TaskTypeName:   it.TaskTypeName,
		EntityType:     it.EntityType,
		CategoryName:   it.CategoryName,
		HasManyRecords: it.HasManyRecords,
		Payloads:       it.Payloads,
		Attributes:     attrCloned,
	}, nil
}

func (it *PayloadWrapper) NonPtr() PayloadWrapper {
	if it == nil {
		return PayloadWrapper{}
	}

	return *it
}

// ToPtr
//
// can panic if nil
func (it PayloadWrapper) ToPtr() *PayloadWrapper {
	return &it
}

func (it PayloadWrapper) AsStandardTaskEntityDefinerContractsBinder() entityinf.StandardTaskEntityDefinerContractsBinder {
	return &it
}

func (it PayloadWrapper) AsPayloadsBinder() payloadinf.PayloadsBinder {
	return &it
}
