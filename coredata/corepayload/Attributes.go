package corepayload

import (
	"bytes"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/payloadinf"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/errcore"
)

type Attributes struct {
	BasicErrWrapper  errcoreinf.BasicErrWrapper `json:"BasicErrWrapper,omitempty"`
	AuthInfo         *AuthInfo                  `json:"AuthInfo,omitempty"`
	PagingInfo       *PagingInfo                `json:"PagingInfo,omitempty"`
	KeyValuePairs    *corestr.Hashmap           `json:"KeyValuePairs,omitempty"`
	AnyKeyValuePairs *coredynamic.MapAnyItems   `json:"AnyKeyValuePairs,omitempty"`
	FromTo           *coreinstruction.FromTo    `json:"FromTo,omitempty"`
	DynamicPayloads  []byte                     `json:"DynamicPayloads,omitempty"`
}

func (it *Attributes) IsNull() bool {
	return it == nil
}

func (it *Attributes) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Attributes) HandleErr() {
	if it.HasError() {
		it.BasicErrWrapper.HandleError()
	}
}

func (it *Attributes) HasStringKey(key string) bool {
	if it.HasKeyValuePairs() {
		return it.KeyValuePairs.Has(key)
	}

	return false
}

func (it *Attributes) HasAnyKey(key string) bool {
	if it.HasAnyKeyValuePairs() {
		return it.AnyKeyValuePairs.HasKey(key)
	}

	return false
}

func (it *Attributes) AddNewStringKeyValueOnly(key, value string) (isAdded bool) {
	if it == nil || it.KeyValuePairs == nil {
		return false
	}

	it.KeyValuePairs.AddOrUpdate(key, value)

	return true
}

func (it *Attributes) AddNewAnyKeyValueOnly(key string, value interface{}) (isAdded bool) {
	if it == nil || it.AnyKeyValuePairs == nil {
		return false
	}

	it.AnyKeyValuePairs.Add(key, value)

	return true
}

func (it *Attributes) GetStringKeyValue(key string) (value string, isFound bool) {
	if it == nil || it.KeyValuePairs == nil {
		return "", false
	}

	return it.KeyValuePairs.Get(key)
}

func (it *Attributes) GetAnyKeyValue(key string) (valueAny interface{}, isFound bool) {
	if it == nil || it.KeyValuePairs == nil {
		return nil, false
	}

	return it.AnyKeyValuePairs.Get(key)
}

func (it *Attributes) AnyKeyReflectSetTo(key string, toPtr interface{}) error {
	if it == nil || it.KeyValuePairs == nil {
		return errcore.
			CannotBeNilOrEmptyType.ErrorNoRefs(
			"KeyValuePairs is nil")
	}

	return it.AnyKeyValuePairs.ReflectSetTo(key, toPtr)
}

func (it *Attributes) HandleError() {
	if it.HasError() {
		it.BasicErrWrapper.HandleError()
	}
}

func (it *Attributes) ReflectSetTo(toPointer interface{}) error {
	return coredynamic.ReflectSetFromTo(it, toPointer)
}

func (it *Attributes) Payloads() []byte {
	if it.IsEmpty() {
		return []byte{}
	}

	return it.DynamicPayloads
}

func (it *Attributes) AnyKeyValMap() map[string]interface{} {
	if it.IsEmpty() {
		return map[string]interface{}{}
	}

	return it.AnyKeyValuePairs.Items
}

func (it *Attributes) Hashmap() map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	return it.KeyValuePairs.Items()
}

func (it *Attributes) CompiledError() error {
	return it.Error()
}

func (it *Attributes) HasIssuesOrEmpty() bool {
	return it.IsEmpty() ||
		!it.IsValid() ||
		it.BasicErrWrapper != nil &&
			it.BasicErrWrapper.HasError()
}

func (it *Attributes) IsSafeValid() bool {
	return it.HasIssuesOrEmpty()
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

func (it *Attributes) HasFromTo() bool {
	return it != nil && it.FromTo != nil
}

func (it *Attributes) IsValid() bool {
	return it != nil &&
		it.IsEmptyError()
}

func (it *Attributes) IsInvalid() bool {
	return it == nil || it.HasIssuesOrEmpty()
}

func (it *Attributes) HasError() bool {
	return it != nil &&
		it.BasicErrWrapper != nil &&
		it.BasicErrWrapper.HasError()
}

func (it *Attributes) Error() error {
	if it.IsEmptyError() {
		return nil
	}

	return it.
		BasicErrWrapper.
		CompiledErrorWithStackTraces()
}

func (it *Attributes) MustBeEmptyError() {
	if it.IsEmptyError() {
		return
	}

	panic(it.Error())
}

// BasicErrorDeserializedTo
//
// Expectation Attributes.ErrorMessage needs to
// be in json format and unmarshalToPointer
// should match reflection types
func (it *Attributes) BasicErrorDeserializedTo(
	unmarshalToPointer interface{},
) error {
	if it.IsEmptyError() {
		return nil
	}

	return corejson.
		Deserialize.
		UsingBytes(
			it.BasicErrWrapper.SerializeMust(),
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

func (it *Attributes) DeserializeDynamicPayloadsToAttributes() (
	newAttr *Attributes, err error,
) {
	newAttr = &Attributes{}
	err = corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		newAttr)

	return newAttr, err
}

func (it *Attributes) DeserializeDynamicPayloadsToPayloadWrapper() (
	payloadWrapper *PayloadWrapper, err error,
) {
	payloadWrapper = New.PayloadWrapper.Empty()
	err = corejson.Deserialize.UsingBytes(
		it.DynamicPayloads,
		payloadWrapper)

	return payloadWrapper, err
}

func (it *Attributes) DeserializeDynamicPayloadsToPayloadWrappersCollection() (
	payloadsCollection *PayloadsCollection, err error,
) {
	return New.
		PayloadsCollection.
		Deserialize(
			it.DynamicPayloads)
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
	return it == nil ||
		it.BasicErrWrapper == nil ||
		it.BasicErrWrapper.IsEmpty()
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

func (it *Attributes) IsPagingInfoEmpty() bool {
	return it == nil ||
		it.PagingInfo.IsEmpty()
}

func (it *Attributes) IsKeyValuePairsEmpty() bool {
	return it == nil ||
		it.KeyValuePairs.IsEmpty()
}

func (it *Attributes) IsAnyKeyValuePairsEmpty() bool {
	return it == nil ||
		it.AnyKeyValuePairs.IsEmpty()
}

func (it *Attributes) IsUserInfoEmpty() bool {
	return it == nil ||
		it.AuthInfo.IsUserInfoEmpty()
}

func (it *Attributes) VirtualUser() *User {
	if it.IsUserInfoEmpty() {
		return nil
	}

	return it.AuthInfo.UserInfo.User
}

func (it *Attributes) SystemUser() *User {
	if it.IsUserInfoEmpty() {
		return nil
	}

	return it.AuthInfo.UserInfo.SystemUser
}

func (it *Attributes) SessionUser() *User {
	if it.IsSessionInfoEmpty() {
		return nil
	}

	return it.AuthInfo.SessionInfo.User
}

func (it *Attributes) IsAuthInfoEmpty() bool {
	return it == nil ||
		it.AuthInfo.IsEmpty()
}

func (it *Attributes) IsSessionInfoEmpty() bool {
	return it == nil ||
		it.AuthInfo.IsSessionInfoEmpty()
}

func (it *Attributes) HasUserInfo() bool {
	return !it.IsUserInfoEmpty()
}

func (it *Attributes) HasAuthInfo() bool {
	return !it.IsAuthInfoEmpty()
}

func (it *Attributes) HasSessionInfo() bool {
	return !it.IsSessionInfoEmpty()
}

func (it *Attributes) SessionInfo() *SessionInfo {
	if it.IsSessionInfoEmpty() {
		return nil
	}

	return it.AuthInfo.SessionInfo
}

func (it *Attributes) AuthType() string {
	if it.IsAuthInfoEmpty() {
		return ""
	}

	return it.AuthInfo.ActionType
}

func (it *Attributes) ResourceName() string {
	if it.IsAuthInfoEmpty() {
		return ""
	}

	return it.AuthInfo.ResourceName
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

func (it Attributes) String() string {
	return it.JsonString()
}

func (it Attributes) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
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

func (it *Attributes) SetBasicErr(
	basicErr errcoreinf.BasicErrWrapper,
) payloadinf.AttributesBinder {
	it.BasicErrWrapper = basicErr

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

func (it *Attributes) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Attributes) IsEqual(attributes *Attributes) bool {
	if it == nil && attributes == nil {
		return true
	}

	if it == nil || attributes == nil {
		return false
	}

	if it == attributes {
		return true
	}

	if it.IsErrorDifferent(attributes.BasicErrWrapper) {
		return false
	}

	if !it.PagingInfo.IsEqual(attributes.PagingInfo) {
		return false
	}

	if !it.KeyValuePairs.IsEqualPtr(attributes.KeyValuePairs) {
		return false
	}

	if !bytes.Equal(
		it.DynamicPayloads,
		attributes.DynamicPayloads) {
		return false
	}

	if !it.AnyKeyValuePairs.IsEqual(attributes.AnyKeyValuePairs) {
		return false
	}

	return true
}

func (it *Attributes) Clone(
	isDeepClone bool,
) (Attributes, error) {
	clonedPtr, err := it.ClonePtr(isDeepClone)

	if err != nil {
		return Attributes{}, err
	}

	if clonedPtr == nil {
		return Attributes{}, nil
	}

	return clonedPtr.NonPtr(), nil
}

func (it *Attributes) ClonePtr(
	isDeepClone bool,
) (*Attributes, error) {
	if it == nil {
		return nil, nil
	}

	if isDeepClone {
		return it.deepClonePtr()
	}

	// NOT deep clone
	return New.
		Attributes.
		All(
			it.AuthInfo,
			it.KeyValuePairs,
			it.AnyKeyValuePairs,
			it.PagingInfo,
			it.DynamicPayloads,
			it.FromTo,
			it.BasicErrWrapper,
		), nil
}

func (it *Attributes) deepClonePtr() (*Attributes, error) {
	anyMap, err := it.AnyKeyValuePairs.ClonePtr()

	if err != nil {
		return nil, err
	}

	var basicErr errcoreinf.BasicErrWrapper

	if it.HasError() {
		basicErr = it.BasicErrWrapper.CloneInterface()
	}

	return New.
		Attributes.
		All(
			it.AuthInfo.ClonePtr(),
			it.KeyValuePairs.ClonePtr(),
			anyMap,
			it.PagingInfo.ClonePtr(),
			corejson.BytesDeepClone(it.DynamicPayloads),
			it.FromTo.ClonePtr(),
			basicErr), nil
}

func (it Attributes) NonPtr() Attributes {
	return it
}

func (it Attributes) AsAttributesBinder() payloadinf.AttributesBinder {
	return &it
}

func (it *Attributes) IsErrorDifferent(basicErr errcoreinf.BasicErrWrapper) bool {
	return !it.IsErrorEqual(basicErr)
}

func (it *Attributes) IsErrorEqual(basicErr errcoreinf.BasicErrWrapper) bool {
	if it.IsEmptyError() || basicErr == nil || basicErr.IsEmpty() {
		return true
	}

	return it.BasicErrWrapper.IsBasicErrEqual(basicErr)
}
