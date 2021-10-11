package coredynamic

import (
	"encoding/json"
	"math"
	"reflect"
	"sort"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
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
		return errcore.
			KeyNotExistInMap.
			ErrorRefOnly(key)
	}

	rawBytes, err := json.Marshal(valInf)

	if err != nil {
		ref := errcore.Var2NoType(
			"key", key,
			"type", TypeName(valInf))

		return errcore.MarshallingFailed.ErrorRefOnly(ref)
	}

	unmarshalErr := json.Unmarshal(rawBytes, unmarshalRef)

	if unmarshalErr != nil {
		ref := errcore.Var3NoType(
			"key", key,
			"StoreType", TypeName(valInf),
			"RequestedType", TypeName(unmarshalRef))

		return errcore.UnMarshallingFailed.ErrorRefOnly(ref)
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
		return errcore.
			KeyNotExistInMap.
			Error("key", key)
	}

	if referenceOut == nil {
		reference := errcore.Var2NoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut))

		return errcore.
			CannotBeNilMessage.
			Error(
				"referenceOut cannot be nil",
				reference)
	}

	outInfRv := reflect.ValueOf(referenceOut)
	foundItemRv := reflect.ValueOf(valInf)

	if outInfRv.Kind() != reflect.Ptr {
		reference := errcore.Var2NoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut))

		return errcore.
			ShouldBePointer.
			Error(
				"referenceOut is not a pointer!",
				reference)
	}

	if outInfRv.IsNil() || foundItemRv.IsNil() {
		reference := errcore.Var3NoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut),
			"foundItemType", TypeName(valInf))

		return errcore.
			CannotBeNilMessage.
			Error(
				"referenceOut or found item is nil",
				reference)
	}

	foundTypeName := foundItemRv.Type().String()
	refOutTypeName := outInfRv.Type().String()
	isTypeNotEqual := foundTypeName != refOutTypeName
	if isTypeNotEqual {
		reference := errcore.Var3NoType(
			"key", key,
			"referenceOutType", refOutTypeName,
			"foundItemType", foundTypeName)

		return errcore.
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

	reference := errcore.Var3NoType(
		"key", key,
		"referenceOutType", TypeName(referenceOut),
		"foundItemType", TypeName(valInf))

	return errcore.
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
		return errcore.
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
		return errcore.
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

func (it *MapAnyItems) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *MapAnyItems) GetPagedCollection(
	eachPageSize int,
) []*MapAnyItems {
	length := it.Length()

	if length < eachPageSize {
		return []*MapAnyItems{
			it,
		}
	}

	allKeys := it.AllKeysSorted()
	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*MapAnyItems, pagesPossibleCeiling)

	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
			allKeys)

		collectionOfCollection[oneBasedPageIndex-1] = pagedCollection

		wg.Done()
	}

	wg.Add(pagesPossibleCeiling)
	for i := 1; i <= pagesPossibleCeiling; i++ {
		go addPagedItemsFunc(i)
	}

	wg.Wait()

	return collectionOfCollection
}

func (it *MapAnyItems) AddMapResultsUsingCloneOption(
	mapResults map[string]interface{},
) *MapAnyItems {
	if len(mapResults) == 0 {
		return it
	}

	for key, result := range mapResults {
		it.Items[key] = result
	}

	return it
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *MapAnyItems) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
	allKeys []string,
) *MapAnyItems {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	if length != len(allKeys) {
		reference := errcore.Var2NoType(
			"MapLength", it.Length(),
			"AllKeysLength", len(allKeys))

		errcore.
			LengthShouldBeEqualToMessage.
			HandleUsingPanic(
				"allKeys length should be exact same as the map length, "+
					"use AllKeys method to get the keys.",
				reference)
	}

	/**
	 * eachPageItems = 10
	 * pageIndex = 4
	 * skipItems = 10 * (4 - 1) = 30
	 */
	skipItems := eachPageSize * (pageIndex - 1)
	if skipItems < 0 {
		errcore.
			CannotBeNegativeIndex.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				pageIndex)
	}

	endingIndex := skipItems + eachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	list := allKeys[skipItems:endingIndex]

	return it.GetNewMapUsingKeys(
		true,
		list...)
}

func (it *MapAnyItems) GetNewMapUsingKeys(
	isPanicOnMissing bool,
	keys ...string,
) *MapAnyItems {
	if len(keys) == 0 {
		return EmptyMapAnyItems()
	}

	mapResults := make(map[string]interface{}, len(keys))

	for _, key := range keys {
		item, has := it.Items[key]

		if isPanicOnMissing && !has {
			errcore.
				KeyNotExistInMap.
				HandleUsingPanic(
					"given key is not found in the map, key ="+key,
					it.AllKeys())
		}

		if has {
			mapResults[key] = item
		}
	}

	return &MapAnyItems{Items: mapResults}
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
		errcore.
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
		errcore.KeyNotExistInMap.Error("Key", key))
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

func (it *MapAnyItems) AllKeysSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	keys := it.AllKeys()
	sort.Strings(keys)

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

func (it *MapAnyItems) JsonResultsCollection() *corejson.ResultsCollection {
	jsonResultsCollection := corejson.NewResultsCollection(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, anyInf := range it.Items {
		jsonResultsCollection.AddAny(
			anyInf)
	}

	return jsonResultsCollection
}

func (it *MapAnyItems) JsonResultsPtrCollection() *corejson.ResultsPtrCollection {
	jsonResultsCollection := corejson.NewResultsPtrCollection(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, anyInf := range it.Items {
		jsonResultsCollection.AddAny(
			anyInf)
	}

	return jsonResultsCollection
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

func (it MapAnyItems) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it MapAnyItems) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *MapAnyItems) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*MapAnyItems, error) {
	err := jsonResult.Unmarshal(it)

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
	return errcore.VarMapStrings(it.Items)
}

func (it *MapAnyItems) String() string {
	return errcore.VarMap(it.Items)
}
