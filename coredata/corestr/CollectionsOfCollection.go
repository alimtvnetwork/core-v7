package corestr

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

type CollectionsOfCollection struct {
	items []*Collection
	sync.RWMutex
}

func (it *CollectionsOfCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *CollectionsOfCollection) IsEmpty() bool {
	return it == nil || it.items == nil || len(it.items) == 0
}

func (it *CollectionsOfCollection) HasItems() bool {
	return it != nil && it.items != nil && len(it.items) > 0
}

func (it *CollectionsOfCollection) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *CollectionsOfCollection) AllIndividualItemsLength() int {
	if it.IsEmpty() {
		return 0
	}

	allLength := 0

	for _, collection := range it.items {
		if collection.IsEmpty() {
			continue
		}

		allLength += collection.Length()
	}

	return allLength
}

func (it *CollectionsOfCollection) Items() []*Collection {
	return it.items
}

func (it *CollectionsOfCollection) List(additionalCapacity int) []string {
	allLength := it.AllIndividualItemsLength()
	list := make([]string, 0, allLength+additionalCapacity)

	if allLength == 0 {
		return list
	}

	for _, collection := range it.items {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		for _, s := range collection.List() {
			list = append(list, s)
		}
	}

	return list
}

func (it *CollectionsOfCollection) ToCollection() *Collection {
	list := it.List(0)

	return New.Collection.Strings(list)
}

func (it *CollectionsOfCollection) AddStrings(
	isCloneAdd bool,
	stringsItems []string,
) *CollectionsOfCollection {
	if len(stringsItems) == 0 {
		return it
	}

	return it.Adds(*New.Collection.StringsOptions(isCloneAdd, stringsItems))
}

func (it *CollectionsOfCollection) AddsStringsOfStrings(
	isMakeClone bool,
	stringsOfPointerStrings ...[]string,
) *CollectionsOfCollection {
	if stringsOfPointerStrings == nil {
		return it
	}

	for _, stringsPointer := range stringsOfPointerStrings {
		it.AddStrings(isMakeClone, stringsPointer)
	}

	return it
}

// AddAsyncFuncItems must add all the lengths to the wg
func (it *CollectionsOfCollection) AddAsyncFuncItems(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() []string,
) *CollectionsOfCollection {
	if asyncFunctions == nil {
		return it
	}

	asyncFuncWrap := func(asyncFunc func() []string) {
		items := asyncFunc()

		if len(items) == 0 {
			wg.Done()

			return
		}

		it.Lock()

		it.AddStrings(
			isMakeClone,
			items,
		)

		it.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return it
}

func (it *CollectionsOfCollection) Adds(
	collections ...Collection,
) *CollectionsOfCollection {
	if collections == nil {
		return it
	}

	return it.AddCollections(collections...)
}

func (it *CollectionsOfCollection) AddCollections(
	collections ...Collection,
) *CollectionsOfCollection {
	if collections == nil {
		return it
	}

	for _, item := range collections {
		it.items = append(it.items, &item)
	}

	return it
}

func (it *CollectionsOfCollection) Add(
	collection *Collection,
) *CollectionsOfCollection {
	if collection.IsEmpty() {
		return it
	}

	it.items = append(it.items, collection)

	return it
}

func (it *CollectionsOfCollection) String() string {
	list := make(
		[]string,
		0,
		it.Length(),
	)

	for i, collection := range it.items {
		list = append(
			list,
			collection.SummaryString(i+1),
		)
	}

	return strings.Join(
		list,
		constants.DoubleNewLine,
	)
}

func (it *CollectionsOfCollection) JsonModel() CollectionsOfCollectionModel {
	return CollectionsOfCollectionModel{
		Items: it.items,
	}
}

func (it *CollectionsOfCollection) JsonModelAny() any {
	return it.JsonModel()
}

func (it *CollectionsOfCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *CollectionsOfCollection) UnmarshalJSON(data []byte) error {
	var dataModel CollectionsOfCollectionModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
	}

	return err
}

func (it CollectionsOfCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it CollectionsOfCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *CollectionsOfCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CollectionsOfCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return Empty.CollectionsOfCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//
//goland:noinspection GoLinterLocal
func (it *CollectionsOfCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CollectionsOfCollection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *CollectionsOfCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *CollectionsOfCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CollectionsOfCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *CollectionsOfCollection) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
