package coredynamic

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/pagingutil"
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

func (it *KeyValCollection) JsonMapResults() *corejson.MapResults {
	mapResults := corejson.NewMapResultsUsingCap(it.Length())

	if it.IsEmpty() {
		return mapResults
	}

	for _, keyVal := range it.items {
		mapResults.AddAny(
			keyVal.KeyString(),
			keyVal.Value)
	}

	return mapResults
}

func (it *KeyValCollection) JsonResultsCollection() *corejson.ResultsCollection {
	jsonResultsCollection := corejson.NewResultsCollection(it.Length())

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
	jsonResultsCollection := corejson.NewResultsPtrCollection(it.Length())

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

func (it *KeyValCollection) JsonString() (jsonString string, err error) {
	toBytes, err := json.Marshal(it.items)

	if err != nil {
		return constants.EmptyString, err
	}

	return string(toBytes), err
}

func (it *KeyValCollection) JsonStringMust() string {
	toString, err := it.JsonString()

	if err != nil {
		errcore.
			MarshallingFailed.
			HandleUsingPanic(err.Error(), it.items)
	}

	return toString
}
