package corejson

import (
	"errors"
	"math"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

type MapResults struct {
	Items map[string]Result `json:"JsonResultsMap"`
}

func (it *MapResults) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *MapResults) LastIndex() int {
	return it.Length() - 1
}

func (it *MapResults) IsEmpty() bool {
	return it.Length() == 0
}

func (it *MapResults) HasAnyItem() bool {
	return it.Length() > 0
}

// AddSkipOnNil skip on nil
func (it *MapResults) AddSkipOnNil(
	key string,
	result *Result,
) *MapResults {
	if result == nil {
		return it
	}

	it.Items[key] = *result

	return it
}

func (it *MapResults) GetByKey(
	key string,
) *Result {
	r, has := it.Items[key]

	if has {
		return &r
	}

	return nil
}

// HasError has any error
func (it *MapResults) HasError() bool {
	for _, result := range it.Items {
		if result.HasError() {
			return true
		}
	}

	return false
}

func (it *MapResults) AllErrors() (
	errListPtr []error,
	hasAnyError bool,
) {
	length := it.Length()
	errList := make(
		[]error,
		0,
		length)

	if length == 0 {
		return errList, hasAnyError
	}

	for key, val := range it.Items {
		err := val.Error

		if err != nil {
			hasAnyError = true
			errList = append(
				errList,
				errors.New(key+constants.HyphenAngelRight+err.Error()))
		}
	}

	return errList, hasAnyError
}

func (it *MapResults) GetErrorsStrings() []string {
	length := it.Length()
	errStrList := make(
		[]string,
		0,
		length)

	if length == 0 {
		return errStrList
	}

	for key, result := range it.Items {
		if result.IsEmptyError() {
			continue
		}

		errStrList = append(
			errStrList,
			key+constants.HyphenAngelRight+result.Error.Error())
	}

	return errStrList
}

func (it *MapResults) GetErrorsStringsPtr() *[]string {
	errStrList := it.GetErrorsStrings()

	return &errStrList
}

func (it *MapResults) GetErrorsAsSingleString() string {
	errStrList := it.GetErrorsStrings()

	return strings.Join(
		errStrList,
		constants.DefaultLine)
}

func (it *MapResults) GetErrorsAsSingle() error {
	errorString := it.GetErrorsAsSingleString()

	return errors.New(errorString)
}

func (it *MapResults) Unmarshal(
	key string,
	any interface{},
) error {
	result, has := it.Items[key]

	if has {
		return errcore.
			KeyNotExistInMapType.
			Error("Given key not found!", key)
	}

	if result.IsEmptyJsonBytes() {
		return errcore.
			EmptyResultCannotMakeJsonType.
			Error("Cannot make json of empty bytes!", key)
	}

	return result.Unmarshal(
		any)
}

func (it *MapResults) Deserialize(
	key string,
	any interface{},
) error {
	return it.Unmarshal(key, any)
}

func (it *MapResults) DeserializeMust(
	key string,
	any interface{},
) *MapResults {
	err := it.Unmarshal(key, any)
	errcore.MustBeEmpty(err)

	return it
}

func (it *MapResults) UnmarshalMany(
	keyAnyItems ...KeyAny,
) error {
	if len(keyAnyItems) == 0 {
		return nil
	}

	for _, keyAny := range keyAnyItems {
		err := it.Unmarshal(
			keyAny.Key,
			keyAny.AnyInf)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *MapResults) UnmarshalManySafe(
	keyAnyItems ...KeyAny,
) error {
	if len(keyAnyItems) == 0 {
		return nil
	}

	for _, keyAny := range keyAnyItems {
		err := it.SafeUnmarshal(
			keyAny.Key,
			keyAny.AnyInf)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *MapResults) SafeUnmarshal(
	key string,
	any interface{},
) error {
	result, has := it.Items[key]

	if has || result.IsEmptyJsonBytes() {
		return nil
	}

	return result.Unmarshal(
		any)
}

func (it *MapResults) SafeDeserialize(
	key string,
	any interface{},
) error {
	return it.SafeUnmarshal(
		key,
		any)
}

func (it *MapResults) SafeDeserializeMust(
	key string,
	any interface{},
) *MapResults {
	err := it.SafeUnmarshal(
		key,
		any)
	errcore.MustBeEmpty(err)

	return it
}

func (it *MapResults) InjectIntoAt(
	key string,
	injector JsonParseSelfInjector,
) error {
	return injector.JsonParseSelfInject(
		it.GetByKey(key))
}

func (it *MapResults) Add(
	key string,
	result Result,
) *MapResults {
	it.Items[key] = result

	return it
}

func (it *MapResults) AddPtr(
	key string,
	result *Result,
) *MapResults {
	if result == nil {
		return it
	}

	it.Items[key] = *result

	return it
}

// AddAny returns error if any during marshalling it.
func (it *MapResults) AddAny(
	key string,
	item interface{},
) error {
	if item == nil {
		return errcore.MarshallingFailedType.Error(
			errcore.CannotBeNilType.String(),
			key)
	}

	jsonResult := NewResult.Any(
		item)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulError()
	}

	it.Add(key, jsonResult)

	return nil
}

// AddAnySkipOnNil returns error if any during marshalling it.
func (it *MapResults) AddAnySkipOnNil(
	key string,
	item interface{},
) error {
	if item == nil {
		return nil
	}

	jsonResult := NewResult.Any(item)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulError()
	}

	it.Add(key, jsonResult)

	return nil
}

func (it *MapResults) AddAnyNonEmptyNonError(
	key string,
	item interface{},
) *MapResults {
	if item == nil {
		return it
	}

	return it.AddNonEmptyNonErrorPtr(
		key,
		NewResult.AnyPtr(item))
}

func (it *MapResults) AddAnyNonEmpty(
	key string,
	item interface{},
) *MapResults {
	if item == nil {
		return it
	}

	return it.Add(
		key,
		NewResult.Any(item))
}

func (it *MapResults) AddKeyWithResult(
	result KeyWithResult,
) *MapResults {
	return it.AddPtr(result.Key, &result.Result)
}

func (it *MapResults) AddKeyWithResultPtr(
	result *KeyWithResult,
) *MapResults {
	if result == nil {
		return it
	}

	return it.AddPtr(result.Key, &result.Result)
}

func (it *MapResults) AddKeysWithResultsPtr(
	results ...*KeyWithResult,
) *MapResults {
	if len(results) == 0 {
		return it
	}

	for _, result := range results {
		it.AddKeyWithResultPtr(result)
	}

	return it
}

func (it *MapResults) AddKeysWithResults(
	results ...KeyWithResult,
) *MapResults {
	if len(results) == 0 {
		return it
	}

	for _, result := range results {
		it.AddKeyWithResult(result)
	}

	return it
}

func (it *MapResults) AddKeyAnyInf(
	result KeyAny,
) *MapResults {
	return it.AddAnyNonEmpty(
		result.Key,
		result.AnyInf)
}

func (it *MapResults) AddKeyAnyInfPtr(
	result *KeyAny,
) *MapResults {
	if result == nil {
		return it
	}

	return it.AddAnyNonEmpty(
		result.Key,
		result.AnyInf)
}

func (it *MapResults) AddKeyAnyItems(
	results ...KeyAny,
) *MapResults {
	if results == nil {
		return it
	}

	for _, result := range results {
		it.AddKeyAnyInf(result)
	}

	return it
}

func (it *MapResults) AddKeyAnyItemsPtr(
	results ...*KeyAny,
) *MapResults {
	if results == nil {
		return it
	}

	for _, result := range results {
		it.AddKeyAnyInfPtr(result)
	}

	return it
}

func (it *MapResults) AddNonEmptyNonErrorPtr(
	key string,
	result *Result,
) *MapResults {
	if result == nil || result.HasError() {
		return it
	}

	it.Items[key] = *result

	return it
}

func (it *MapResults) AddMapResults(
	mapResults *MapResults,
) *MapResults {
	if mapResults == nil || mapResults.IsEmpty() {
		return it
	}

	for key := range mapResults.Items {
		it.Items[key] = mapResults.Items[key]
	}

	return it
}

func (it *MapResults) AddMapAnyItems(
	addOrUpdateMap map[string]interface{},
) *MapResults {
	if len(addOrUpdateMap) == 0 {
		return it
	}

	for key := range addOrUpdateMap {
		it.Items[key] = NewResult.Any(addOrUpdateMap[key])
	}

	return it
}

func (it *MapResults) AllKeys() []string {
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

func (it *MapResults) AllKeysSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	keys := it.AllKeys()
	sort.Strings(keys)

	return keys
}

func (it *MapResults) AllValues() []Result {
	if it.IsEmpty() {
		return []Result{}
	}

	values := make([]Result, it.Length())

	index := 0
	for _, result := range it.Items {
		values[index] = result
		index++
	}

	return values
}

func (it *MapResults) AllResultsCollection() *ResultsCollection {
	if it.IsEmpty() {
		return Empty.ResultsCollection()
	}

	resultsCollection := NewResultsCollection.UsingCap(
		it.Length())

	index := 0
	for _, result := range it.Items {
		resultsCollection.Add(result)
		index++
	}

	return resultsCollection
}

func (it *MapResults) AllResults() []Result {
	return it.AllValues()
}

func (it *MapResults) GetStrings() []string {
	length := it.Length()
	list := make([]string, length)

	if length == 0 {
		return list
	}

	index := 0
	for _, result := range it.Items {
		list[index] = *result.JsonStringPtr()
		index++
	}

	return list
}

func (it *MapResults) GetStringsPtr() *[]string {
	stringsItems := it.GetStrings()

	return &stringsItems
}

// AddJsoner skip on nil
func (it *MapResults) AddJsoner(
	key string,
	jsoner Jsoner,
) *MapResults {
	if jsoner == nil {
		return it
	}

	return it.AddPtr(key, jsoner.JsonPtr())
}

func (it *MapResults) AddKeyWithJsoner(
	keyWithJsoner KeyWithJsoner,
) *MapResults {
	return it.AddJsoner(
		keyWithJsoner.Key,
		keyWithJsoner.Jsoner)
}

func (it *MapResults) AddKeysWithJsoners(
	keysWithJsoners ...KeyWithJsoner,
) *MapResults {
	if keysWithJsoners == nil {
		return nil
	}

	for _, jsoner := range keysWithJsoners {
		it.AddKeyWithJsoner(jsoner)
	}

	return it
}

func (it *MapResults) AddKeyWithJsonerPtr(
	keyWithJsoner *KeyWithJsoner,
) *MapResults {
	if keyWithJsoner == nil || keyWithJsoner.Jsoner == nil {
		return it
	}

	return it.AddJsoner(
		keyWithJsoner.Key,
		keyWithJsoner.Jsoner)
}

func (it *MapResults) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *MapResults) GetPagedCollection(
	eachPageSize int,
) []*MapResults {
	length := it.Length()

	if length < eachPageSize {
		return []*MapResults{
			it,
		}
	}

	allKeys := it.AllKeysSorted()
	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*MapResults, pagesPossibleCeiling)

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

func (it *MapResults) AddMapResultsUsingCloneOption(
	isClone, isDeepClone bool,
	mapResults map[string]Result,
) *MapResults {
	if len(mapResults) == 0 {
		return it
	}

	if !isClone && !isDeepClone {
		for key, result := range mapResults {
			it.Items[key] = result
		}

		return it
	}

	for key, result := range mapResults {
		cloned := result.CloneIf(
			isClone,
			isDeepClone)

		it.Items[key] = cloned
	}

	return it
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *MapResults) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
	allKeys []string,
) *MapResults {
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

func (it *MapResults) GetNewMapUsingKeys(
	isPanicOnMissing bool,
	keys ...string,
) *MapResults {
	if len(keys) == 0 {
		return Empty.MapResults()
	}

	mapResults := make(
		map[string]Result,
		len(keys))

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

	return &MapResults{Items: mapResults}
}

func (it *MapResults) ResultCollection() *ResultsCollection {
	if it.IsEmpty() {
		return Empty.ResultsCollection()
	}

	results := NewResultsCollection.UsingCap(
		it.Length())

	return results.AddRawMapResults(
		it.Items)
}

//goland:noinspection GoLinterLocal
func (it *MapResults) JsonModel() *MapResults {
	return it
}

//goland:noinspection GoLinterLocal
func (it *MapResults) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *MapResults) Clear() *MapResults {
	if it == nil {
		return it
	}

	temp := it.Items
	clearFunc := func() {
		for _, result := range temp {
			result.Dispose()
		}
	}

	go clearFunc()
	it.Items = map[string]Result{}

	return it
}

func (it *MapResults) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it MapResults) Json() Result {
	return NewResult.Any(it)
}

func (it MapResults) JsonPtr() *Result {
	return NewResult.AnyPtr(it)
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *MapResults) ParseInjectUsingJson(
	jsonResult *Result,
) (*MapResults, error) {
	err := jsonResult.Unmarshal(
		&it,
	)

	if err != nil {
		return Empty.MapResults(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *MapResults) ParseInjectUsingJsonMust(
	jsonResult *Result,
) *MapResults {
	resultCollection, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return resultCollection
}

func (it *MapResults) AsJsonContractsBinder() JsonContractsBinder {
	return it
}

func (it *MapResults) AsJsoner() Jsoner {
	return it
}

func (it *MapResults) JsonParseSelfInject(
	jsonResult *Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *MapResults) AsJsonParseSelfInjector() JsonParseSelfInjector {
	return it
}
