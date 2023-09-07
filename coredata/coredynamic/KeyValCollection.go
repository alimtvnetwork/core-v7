package coredynamic

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/pagingutil"
)

type KeyValCollection struct {
	items []KeyVal
}

func EmptyKeyValCollection() *KeyValCollection {
	return NewKeyValCollection(constants.Zero)
}

func NewKeyValCollection(capacity int) *KeyValCollection {
	slice := make([]KeyVal, 0, capacity)

	return &KeyValCollection{items: slice}
}

func (it *KeyValCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *KeyValCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *KeyValCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *KeyValCollection) Add(
	keyVal KeyVal,
) *KeyValCollection {
	it.items = append(it.items, keyVal)

	return it
}

func (it *KeyValCollection) AddPtr(
	keyVal *KeyVal,
) *KeyValCollection {
	if keyVal == nil {
		return it
	}

	it.items = append(it.items, *keyVal)

	return it
}

func (it *KeyValCollection) AddMany(
	keyValues ...KeyVal,
) *KeyValCollection {
	if keyValues == nil || len(keyValues) == 0 {
		return it
	}

	for _, keyVal := range keyValues {
		it.items = append(
			it.items,
			keyVal)
	}

	return it
}

func (it *KeyValCollection) AddManyPtr(
	keyValues ...*KeyVal,
) *KeyValCollection {
	if keyValues == nil || len(keyValues) == 0 {
		return it
	}

	for _, keyVal := range keyValues {
		if keyVal == nil {
			continue
		}

		it.items = append(
			it.items,
			*keyVal)
	}

	return it
}

func (it *KeyValCollection) Items() []KeyVal {
	return it.items
}

func (it *KeyValCollection) MapAnyItems() *MapAnyItems {
	if it.IsEmpty() {
		return EmptyMapAnyItems()
	}

	mapItems := make(map[string]interface{}, it.Length())
	for _, keyVal := range it.items {
		mapItems[keyVal.KeyString()] = keyVal.Value
	}

	return &MapAnyItems{Items: mapItems}
}

func (it *KeyValCollection) JsonMapResults() (*corejson.MapResults, error) {
	mapResults := corejson.
		NewMapResults.
		UsingCap(it.Length())

	if it.IsEmpty() {
		return mapResults, nil
	}

	for _, keyVal := range it.items {
		err := mapResults.AddAny(
			keyVal.KeyString(),
			keyVal.Value)

		if err != nil {
			return mapResults, err
		}
	}

	return mapResults, nil
}

func (it *KeyValCollection) JsonResultsCollection() *corejson.ResultsCollection {
	jsonResultsCollection := corejson.NewResultsCollection.UsingCap(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, keyVal := range it.items {
		jsonResultsCollection.AddAny(
			keyVal.Value)
	}

	return jsonResultsCollection
}

func (it *KeyValCollection) JsonResultsPtrCollection() *corejson.ResultsPtrCollection {
	jsonResultsCollection := corejson.NewResultsPtrCollection.UsingCap(it.Length())

	if it.IsEmpty() {
		return jsonResultsCollection
	}

	for _, keyVal := range it.items {
		jsonResultsCollection.AddAny(
			keyVal.Value)
	}

	return jsonResultsCollection
}

func (it *KeyValCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *KeyValCollection) GetPagedCollection(
	eachPageSize int,
) []*KeyValCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*KeyValCollection{
			it,
		}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make(
		[]*KeyValCollection,
		pagesPossibleCeiling)

	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
		)

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

func (it *KeyValCollection) GetPagingInfo(
	eachPageSize int,
	pageIndex int,
) pagingutil.PagingInfo {
	return pagingutil.GetPagingInfo(pagingutil.PagingRequest{
		Length:       it.Length(),
		PageIndex:    pageIndex,
		EachPageSize: eachPageSize,
	})
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *KeyValCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *KeyValCollection {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	pageInfo := it.GetPagingInfo(
		eachPageSize,
		pageIndex)

	list := it.items[pageInfo.SkipItems:pageInfo.EndingLength]

	return &KeyValCollection{
		items: list,
	}
}

func (it *KeyValCollection) AllKeys() []string {
	if it.IsEmpty() {
		return []string{}
	}

	keys := make([]string, it.Length())

	for i, keyVal := range it.items {
		keys[i] = keyVal.KeyString()
	}

	return keys
}

func (it *KeyValCollection) AllKeysSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	keys := it.AllKeys()
	sort.Strings(keys)

	return keys
}

func (it *KeyValCollection) AllValues() []interface{} {
	if it.IsEmpty() {
		return []interface{}{}
	}

	values := make([]interface{}, it.Length())

	for i, result := range it.items {
		values[i] = result.Value
	}

	return values
}

func (it *KeyValCollection) String() string {
	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.items)
}

func (it KeyValCollection) JsonModel() interface{} {
	return it
}

func (it KeyValCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it KeyValCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it KeyValCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *KeyValCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*KeyValCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *KeyValCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *KeyValCollection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *KeyValCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *KeyValCollection) Serialize() (jsonBytesPtr []byte, err error) {
	jsonResult := it.Json()

	if jsonResult.HasError() {
		return []byte{}, jsonResult.MeaningfulError()
	}

	return jsonResult.SafeBytes(), nil
}

func (it *KeyValCollection) JsonString() (jsonString string, err error) {
	jsonResult := it.Json()

	if jsonResult.HasError() {
		return constants.EmptyString, jsonResult.MeaningfulError()
	}

	return jsonResult.JsonString(), nil
}

func (it *KeyValCollection) JsonStringMust() string {
	jsonResult := it.Json()
	jsonResult.HandleError()

	return jsonResult.JsonString()
}

func (it KeyValCollection) Clone() KeyValCollection {
	keyValCollection := NewKeyValCollection(it.Length())
	keyValCollection.AddMany(it.items...)

	return keyValCollection.NonPtr()
}

func (it *KeyValCollection) ClonePtr() *KeyValCollection {
	if it == nil {
		return nil
	}

	cloned := it.Clone()

	return cloned.Ptr()
}

func (it KeyValCollection) NonPtr() KeyValCollection {
	return it
}

func (it *KeyValCollection) Ptr() *KeyValCollection {
	return it
}
