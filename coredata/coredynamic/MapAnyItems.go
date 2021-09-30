package coredynamic

import (
	"encoding/json"
	"reflect"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/msgtype"
)

type MapAnyItems struct {
	Items map[string]interface{}
}

func EmptyMapAnyItems() *MapAnyItems {
	return NewMapAnyItems(constants.Zero)
}

func NewMapAnyItems(capacity int) *MapAnyItems {
	slice := make(map[string]interface{}, capacity)

	return &MapAnyItems{Items: slice}
}

func (it *MapAnyItems) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *MapAnyItems) IsEmpty() bool {
	return it.Length() == 0
}

func (it *MapAnyItems) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *MapAnyItems) HasKey(key string) bool {
	_, has := it.Items[key]

	return has
}

func (it *MapAnyItems) Get(
	key string,
) (any interface{}, has bool) {
	valInf, has := it.Items[key]

	if has {
		return valInf, has
	}

	return nil, false
}

func (it *MapAnyItems) GetUsingUnmarshallAt(
	key string,
	unmarshalRef interface{},
) error {
	valInf, has := it.Items[key]

	if !has {
		return msgtype.
			KeyNotExistInMap.
			ErrorRefOnly(key)
	}

	rawBytes, err := json.Marshal(valInf)

	if err != nil {
		ref := msgtype.Var2NoType(
			"key", key,
			"type", TypeName(valInf))

		return msgtype.MarshallingFailed.ErrorRefOnly(ref)
	}

	unmarshalErr := json.Unmarshal(rawBytes, unmarshalRef)

	if unmarshalErr != nil {
		ref := msgtype.Var3NoType(
			"key", key,
			"StoreType", TypeName(valInf),
			"RequestedType", TypeName(unmarshalRef))

		return msgtype.UnMarshallingFailed.ErrorRefOnly(ref)
	}

	return nil
}

func (it *MapAnyItems) GetUsingUnmarshallManyAt(
	keyAnys ...corejson.KeyAny,
) error {
	for _, keyAny := range keyAnys {
		err := it.GetUsingUnmarshallAt(
			keyAny.Key,
			keyAny.AnyInf)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *MapAnyItems) GetManyItemsRefs(
	keyAnys ...corejson.KeyAny,
) error {
	if len(keyAnys) == 0 {
		return nil
	}

	for _, keyAny := range keyAnys {
		err := it.GetItemRef(
			keyAny.Key,
			keyAny.AnyInf)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *MapAnyItems) GetItemRef(
	key string,
	referenceOut interface{},
) error {
	valInf, has := it.Items[key]

	if !has {
		return msgtype.
			KeyNotExistInMap.
			Error("key", key)
	}

	if referenceOut == nil {
		reference := msgtype.Var2NoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut))

		return msgtype.
			CannotBeNilMessage.
			Error(
				"referenceOut cannot be nil",
				reference)
	}

	outInfRv := reflect.ValueOf(referenceOut)
	foundItemRv := reflect.ValueOf(valInf)

	if outInfRv.Kind() != reflect.Ptr {
		reference := msgtype.Var2NoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut))

		return msgtype.
			ShouldBePointer.
			Error(
				"referenceOut is not a pointer!",
				reference)
	}

	if outInfRv.IsNil() || foundItemRv.IsNil() {
		reference := msgtype.Var3NoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut),
			"foundItemType", TypeName(valInf))

		return msgtype.
			CannotBeNilMessage.
			Error(
				"referenceOut or found item is nil",
				reference)
	}

	foundTypeName := foundItemRv.Type().String()
	refOutTypeName := outInfRv.Type().String()
	isTypeNotEqual := foundTypeName != refOutTypeName
	if isTypeNotEqual {
		reference := msgtype.Var3NoType(
			"key", key,
			"referenceOutType", refOutTypeName,
			"foundItemType", foundTypeName)

		return msgtype.
			TypeMismatch.
			Error(
				"Use UnmarshalAt method to get generic data to specific type.",
				reference)
	}

	if foundItemRv.Kind() != reflect.Ptr {
		outInfRv.Elem().Set(foundItemRv)

		return nil
	}

	if foundItemRv.Kind() == reflect.Ptr {
		outInfRv.Elem().Set(foundItemRv.Elem())

		return nil
	}

	reference := msgtype.Var3NoType(
		"key", key,
		"referenceOutType", TypeName(referenceOut),
		"foundItemType", TypeName(valInf))

	return msgtype.
		UnexpectedValueErrorMessage.
		Error(
			"unknown error",
			reference)
}

func (it *MapAnyItems) Add(
	key string,
	valInf interface{},
) *MapAnyItems {
	it.Items[key] = valInf

	return it
}

func (it *MapAnyItems) AddKeyAny(
	keyAny corejson.KeyAny,
) *MapAnyItems {
	return it.Add(
		keyAny.Key,
		keyAny.AnyInf)
}

func (it *MapAnyItems) AddKeyAnyWithValidation(
	typeVerify reflect.Type,
	keyAny corejson.KeyAny,
) error {
	actualTypeOf := reflect.TypeOf(keyAny.AnyInf)
	if actualTypeOf != typeVerify {
		return msgtype.
			TypeMismatch.
			Expecting(
				typeVerify.String(),
				actualTypeOf.String())
	}

	it.AddKeyAny(keyAny)

	return nil
}

func (it *MapAnyItems) AddWithValidation(
	typeVerify reflect.Type,
	key string,
	anyInf interface{},
) error {
	actualTypeOf := reflect.TypeOf(anyInf)
	if actualTypeOf != typeVerify {
		return msgtype.
			TypeMismatch.
			Expecting(
				typeVerify.String(),
				actualTypeOf.String())
	}

	it.Add(key, anyInf)

	return nil
}

func (it *MapAnyItems) AddJsonResultPtr(
	key string,
	jsonResult *corejson.Result,
) *MapAnyItems {
	if jsonResult == nil {
		return it
	}

	it.Items[key] = jsonResult

	return it
}

func (it *MapAnyItems) JsonString() (jsonString string, err error) {
	toBytes, err := json.Marshal(it.Items)

	if err != nil {
		return constants.EmptyString, err
	}

	return string(toBytes), err
}

func (it *MapAnyItems) JsonStringMust() string {
	toString, err := it.JsonString()

	if err != nil {
		msgtype.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), it.Items)
	}

	return toString
}

func (it *MapAnyItems) JsonResultOfKey(key string) *corejson.Result {
	item, has := it.Get(key)

	if has {
		return corejson.NewFromAnyPtr(item)
	}

	return corejson.EmptyWithErrorPtr(
		msgtype.KeyNotExistInMap.Error("Key", key))
}

func (it *MapAnyItems) JsonResultOfKeys(
	keys ...string,
) *corejson.MapResults {
	mapResults := corejson.NewMapResultsUsingCap(len(keys))

	if len(keys) == 0 {
		return mapResults
	}

	for _, key := range keys {
		mapResults.AddPtr(
			key,
			it.JsonResultOfKey(key))
	}

	return mapResults
}

func (it *MapAnyItems) AllKeys() []string {
	if it.IsEmpty() {
		return []string{}
	}

	keys := make([]string, it.Length())

	index := 0
	for key := range it.Items {
		keys[index] = key
		index++
	}

	return keys
}

func (it *MapAnyItems) AllValues() []interface{} {
	if it.IsEmpty() {
		return []interface{}{}
	}

	values := make([]interface{}, it.Length())

	index := 0
	for _, result := range it.Items {
		values[index] = result
		index++
	}

	return values
}

func (it *MapAnyItems) JsonMapResults() *corejson.MapResults {
	mapResults := corejson.NewMapResultsUsingCap(it.Length())

	if it.IsEmpty() {
		return mapResults
	}

	for key, anyInf := range it.Items {
		mapResults.AddAny(
			key,
			anyInf)
	}

	return mapResults
}

func (it *MapAnyItems) JsonModel() *corejson.MapResults {
	mapResults := corejson.NewMapResultsUsingCap(
		it.Length() +
			constants.Capacity3)

	if it.IsEmpty() {
		return mapResults
	}

	for key, anyInf := range it.Items {
		mapResults.AddAnyNonEmpty(key, anyInf)
	}

	return mapResults
}

func (it *MapAnyItems) JsonModelAny() interface{} {
	return it.JsonModel()
}

//goland:noinspection GoLinterLocal
func (it *MapAnyItems) Json() *corejson.Result {
	if it.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(it)

	return corejson.NewPtr(jsonBytes, err)
}

//goland:noinspection GoLinterLocal
func (it *MapAnyItems) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*MapAnyItems, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return EmptyMapAnyItems(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *MapAnyItems) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *MapAnyItems {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *MapAnyItems) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *MapAnyItems) Strings() []string {
	return msgtype.VarMapStrings(it.Items)
}

func (it *MapAnyItems) String() string {
	return msgtype.VarMap(it.Items)
}
