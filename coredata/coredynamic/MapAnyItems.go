package coredynamic

import (
	"encoding/json"
	"math"
	"reflect"
	"sort"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/mapdiffinternal"
	"gitlab.com/evatix-go/core/internal/reflectinternal"
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

func NewMapAnyItemsUsingAnyTypeMap(
	anyTypeOfMap interface{},
) (*MapAnyItems, error) {
	if reflectinternal.IsNull(anyTypeOfMap) {
		return EmptyMapAnyItems(), errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("given any map was nil, cannot process it.")
	}

	rv := reflect.ValueOf(anyTypeOfMap)
	convertedMap, err := AnyTypeMapToMapStringAny(rv)

	if err != nil {
		return EmptyMapAnyItems(), err
	}

	return &MapAnyItems{Items: convertedMap}, nil
}

func NewMapAnyItemsUsingItems(
	itemsMap map[string]interface{},
) *MapAnyItems {
	if len(itemsMap) == 0 {
		return EmptyMapAnyItems()
	}

	return &MapAnyItems{Items: itemsMap}
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

func (it *MapAnyItems) ReflectSetTo(
	key string,
	toPointerOrBytes interface{},
) error {
	valInf, has := it.Items[key]

	if !has {
		return errcore.ErrorWithRefToError(
			defaulterr.KeyNotExistInMap,
			it.AllKeysSorted())
	}

	return ReflectSetFromTo(
		valInf,
		toPointerOrBytes)
}

func (it *MapAnyItems) ReflectSetToMust(
	key string,
	toPointerOrBytes interface{},
) {
	err := it.ReflectSetTo(key, toPointerOrBytes)
	errcore.HandleErr(err)
}

func (it *MapAnyItems) GetValue(
	key string,
) (any interface{}) {
	valInf, has := it.Items[key]

	if has {
		return valInf
	}

	return nil
}

func (it *MapAnyItems) GetFieldsMap(
	key string,
) (
	fieldMap map[string]interface{},
	parsingErr error,
	isFound bool,
) {
	valInf, has := it.Items[key]

	if has {
		fieldsMap, parsingErr := corejson.
			Deserialize.
			AnyToFieldsMap(valInf)

		return fieldsMap, parsingErr, true
	}

	return nil, nil, false
}

// GetSafeFieldsMap
//
// Warning:
//  Swallows the parsing err if any
func (it *MapAnyItems) GetSafeFieldsMap(
	key string,
) (
	fieldMap map[string]interface{},
	isFound bool,
) {
	fieldMap, _, isFound = it.GetFieldsMap(key)

	return fieldMap, isFound
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

func (it *MapAnyItems) Deserialize(
	key string,
	toPointer interface{},
) error {
	return it.GetUsingUnmarshallAt(
		key,
		toPointer)
}

func (it *MapAnyItems) DeserializeMust(
	key string,
	toPointer interface{},
) {
	err := it.GetUsingUnmarshallAt(
		key,
		toPointer)
	errcore.HandleErr(err)
}

func (it *MapAnyItems) GetUsingUnmarshallAt(
	key string,
	unmarshalRef interface{},
) error {
	valInf, has := it.Items[key]

	if !has {
		return errcore.
			KeyNotExistInMapType.
			ErrorRefOnly(key)
	}

	rawBytes, err := json.Marshal(valInf)

	if err != nil {
		ref := errcore.VarTwoNoType(
			"key", key,
			"type", TypeName(valInf))

		return errcore.MarshallingFailedType.ErrorRefOnly(ref)
	}

	unmarshalErr := json.Unmarshal(rawBytes, unmarshalRef)

	if unmarshalErr != nil {
		ref := errcore.VarThreeNoType(
			"key", key,
			"StoreType", TypeName(valInf),
			"RequestedType", TypeName(unmarshalRef))

		return errcore.UnMarshallingFailedType.ErrorRefOnly(ref)
	}

	return nil
}

func (it *MapAnyItems) GetUsingUnmarshallManyAt(
	keyAnyItems ...corejson.KeyAny,
) error {
	for _, keyAny := range keyAnyItems {
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
	keyAnyItems ...corejson.KeyAny,
) error {
	if len(keyAnyItems) == 0 {
		return nil
	}

	for _, keyAny := range keyAnyItems {
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
			KeyNotExistInMapType.
			Error("key", key)
	}

	if referenceOut == nil {
		reference := errcore.VarTwoNoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut))

		return errcore.
			CannotBeNilType.
			Error(
				"referenceOut cannot be nil",
				reference)
	}

	outInfRv := reflect.ValueOf(referenceOut)
	foundItemRv := reflect.ValueOf(valInf)

	if outInfRv.Kind() != reflect.Ptr {
		reference := errcore.VarTwoNoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut))

		return errcore.
			ShouldBePointerType.
			Error(
				"referenceOut is not a pointer!",
				reference)
	}

	if outInfRv.IsNil() || foundItemRv.IsNil() {
		reference := errcore.VarThreeNoType(
			"key", key,
			"referenceOutType", TypeName(referenceOut),
			"foundItemType", TypeName(valInf))

		return errcore.
			CannotBeNilType.
			Error(
				"referenceOut or found item is nil",
				reference)
	}

	foundTypeName := foundItemRv.Type().String()
	refOutTypeName := outInfRv.Type().String()
	isTypeNotEqual := foundTypeName != refOutTypeName
	if isTypeNotEqual {
		reference := errcore.VarThreeNoType(
			"key", key,
			"referenceOutType", refOutTypeName,
			"foundItemType", foundTypeName)

		return errcore.
			TypeMismatchType.
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

	reference := errcore.VarThreeNoType(
		"key", key,
		"referenceOutType", TypeName(referenceOut),
		"foundItemType", TypeName(valInf))

	return errcore.
		UnexpectedValueType.
		Error(
			"unknown error",
			reference)
}

func (it *MapAnyItems) Add(
	key string,
	valInf interface{},
) (isNewlyAdded bool) {
	_, isAlreadyExist := it.Items[key]

	it.Items[key] = valInf

	return !isAlreadyExist
}

func (it *MapAnyItems) Set(
	key string,
	valInf interface{},
) (isNewlyAdded bool) {
	_, isAlreadyExist := it.Items[key]

	it.Items[key] = valInf

	return !isAlreadyExist
}

func (it *MapAnyItems) AddKeyAny(
	keyAny corejson.KeyAny,
) (isNewlyAdded bool) {
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
			TypeMismatchType.
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
			TypeMismatchType.
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

// AddMapResult
//
//  apply override on existing result
func (it *MapAnyItems) AddMapResult(
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

func (it *MapAnyItems) AddMapResultOption(
	isOverride bool,
	mapResults map[string]interface{},
) *MapAnyItems {
	if len(mapResults) == 0 {
		return it
	}

	if isOverride {
		return it.AddMapResult(mapResults)
	}

	// no override
	for key, result := range mapResults {
		_, isFound := it.Items[key]

		if !isFound {
			continue
		}

		it.Items[key] = result
	}

	return it
}

func (it *MapAnyItems) AddManyMapResultsUsingOption(
	isOverridingExisting bool,
	mapsOfMapsResults ...map[string]interface{},
) *MapAnyItems {
	if len(mapsOfMapsResults) == 0 {
		return it
	}

	for _, mapResult := range mapsOfMapsResults {
		it.AddMapResultOption(
			isOverridingExisting,
			mapResult)
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
		reference := errcore.VarTwoNoType(
			"MapLength", it.Length(),
			"AllKeysLength", len(allKeys))

		errcore.
			LengthShouldBeEqualToType.
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
			CannotBeNegativeIndexType.
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
				KeyNotExistInMapType.
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
			MarshallingFailedType.
			HandleUsingPanic(err.Error(), it.Items)
	}

	return toString
}

func (it *MapAnyItems) JsonResultOfKey(
	key string,
) *corejson.Result {
	item, has := it.Get(key)

	if has {
		return corejson.NewPtr(item)
	}

	err := errcore.
		KeyNotExistInMapType.
		Error("Key", key)

	return corejson.
		Empty.
		ResultPtrWithErr(
			reflectinternal.TypeName(it),
			err)
}

func (it *MapAnyItems) JsonResultOfKeys(
	keys ...string,
) *corejson.MapResults {
	mapResults := corejson.NewMapResults.UsingCap(len(keys))

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

func (it *MapAnyItems) DiffRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) map[string]interface{} {
	mapDiffer := mapdiffinternal.MapStringAnyDiff(
		rightMap)

	return mapDiffer.DiffRaw(
		isRegardlessType,
		rightMap)
}

func (it *MapAnyItems) Diff(
	isRegardlessType bool,
	rightMap *MapAnyItems,
) *MapAnyItems {
	rawMap := it.DiffRaw(
		isRegardlessType,
		rightMap.Items)

	return NewMapAnyItemsUsingItems(rawMap)
}

func (it *MapAnyItems) IsRawEqual(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	differ := it.RawMapStringAnyDiff()

	return differ.
		IsRawEqual(
			isRegardlessType,
			rightMap)
}

func (it *MapAnyItems) HashmapDiffUsingRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) MapAnyItemDiff {
	diffMap := it.DiffRaw(
		isRegardlessType,
		rightMap)

	if len(diffMap) == 0 {
		return map[string]interface{}{}
	}

	return diffMap
}

func (it *MapAnyItems) MapAnyItems() *MapAnyItems {
	return it
}

func (it *MapAnyItems) HasAnyChanges(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	return !it.IsRawEqual(
		isRegardlessType,
		rightMap)
}

func (it *MapAnyItems) MapStringAnyDiff() mapdiffinternal.MapStringAnyDiff {
	return it.Items
}

func (it *MapAnyItems) DiffJsonMessage(
	isRegardlessType bool,
	rightMap map[string]interface{},
) string {
	differ := it.RawMapStringAnyDiff()

	return differ.DiffJsonMessage(
		isRegardlessType,
		rightMap)
}

func (it *MapAnyItems) ToStringsSliceOfDiffMap(
	diffMap map[string]interface{},
) (diffSlice []string) {
	differ := it.RawMapStringAnyDiff()

	return differ.ToStringsSliceOfDiffMap(
		diffMap)
}

func (it *MapAnyItems) ShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) string {
	differ := it.RawMapStringAnyDiff()

	return differ.ShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)
}

func (it *MapAnyItems) LogShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) (diffMessage string) {
	differ := it.RawMapStringAnyDiff()

	return differ.LogShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)
}

func (it *MapAnyItems) JsonMapResults() (*corejson.MapResults, error) {
	mapResults := corejson.NewMapResults.UsingCap(it.Length())

	if it.IsEmpty() {
		return mapResults, nil
	}

	for key, anyInf := range it.Items {
		err := mapResults.AddAny(
			key,
			anyInf)

		if err != nil {
			return mapResults, err
		}
	}

	return mapResults, nil
}

func (it *MapAnyItems) JsonResultsCollection() *corejson.ResultsCollection {
	jsonResultsCollection := corejson.
		NewResultsCollection.
		UsingCap(it.Length())

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
	jsonResultsCollection := corejson.NewResultsPtrCollection.UsingCap(it.Length())

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
	mapResults := corejson.NewMapResults.UsingCap(
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
	return corejson.New(it)
}

func (it MapAnyItems) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
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

func (it *MapAnyItems) Clear() {
	if it == nil {
		return
	}

	it.Items = map[string]interface{}{}
}

func (it *MapAnyItems) DeepClear() {
	if it == nil {
		return
	}

	tempItems := it.Items

	tempClearFunc := func() {
		for key := range tempItems {
			delete(tempItems, key)
		}
	}

	go tempClearFunc()

	it.Items = map[string]interface{}{}
}

func (it *MapAnyItems) Dispose() {
	if it == nil {
		return
	}

	it.DeepClear()
	it.Items = nil
}

func (it *MapAnyItems) String() string {
	return errcore.VarMap(it.Items)
}

func (it *MapAnyItems) IsEqualRaw(
	rightMappedItems map[string]interface{},
) bool {
	if it == nil && rightMappedItems == nil {
		return true
	}

	if it == nil || rightMappedItems == nil {
		return false
	}

	leftLength := it.Length()
	rightLength := len(rightMappedItems)

	if leftLength != rightLength {
		return false
	}

	for key := range it.Items {
		rightElem, has := rightMappedItems[key]

		if !has {
			return false
		}

		leftElem := it.Items[key]
		if !reflectinternal.IsAnyEqual(leftElem, rightElem) {
			return false
		}
	}

	return true
}

func (it *MapAnyItems) IsEqual(
	right *MapAnyItems,
) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	leftLength := it.Length()
	rightLength := right.Length()

	if leftLength != rightLength {
		return false
	}

	return it.IsEqualRaw(right.Items)
}

func (it *MapAnyItems) ClonePtr() (*MapAnyItems, error) {
	if it == nil {
		return nil, defaulterr.NilResult
	}

	jsonResult := it.Json()
	if jsonResult.HasError() {
		return EmptyMapAnyItems(), jsonResult.MeaningfulError()
	}

	bytesConv := NewBytesConverter(
		jsonResult.Bytes)

	return bytesConv.ToMapAnyItems()
}

func (it *MapAnyItems) RawMapStringAnyDiff() mapdiffinternal.MapStringAnyDiff {
	if it == nil {
		return map[string]interface{}{}
	}

	return it.Items
}
