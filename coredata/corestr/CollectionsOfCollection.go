package corestr

import (
	"encoding/json"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type CollectionsOfCollection struct {
	items *[]*Collection
	sync.Mutex
}

func (it *CollectionsOfCollection) IsEmpty() bool {
	return it.items == nil || len(*it.items) == 0
}

func (it *CollectionsOfCollection) HasItems() bool {
	return it.items != nil && len(*it.items) > 0
}

func (it *CollectionsOfCollection) Length() int {
	if it.items == nil {
		return 0
	}

	return len(*it.items)
}

func (it *CollectionsOfCollection) AllIndividualItemsLength() int {
	if it.IsEmpty() {
		return 0
	}

	allLength := 0

	for _, collection := range *it.items {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		allLength += collection.Length()
	}

	return allLength
}

func (it *CollectionsOfCollection) ItemsPtr() *[]*Collection {
	return it.items
}

func (it *CollectionsOfCollection) Items() []*Collection {
	return *it.items
}

func (it *CollectionsOfCollection) ListPtr(additionalCapacity int) *[]string {
	allLength := it.AllIndividualItemsLength()
	list := make([]string, 0, allLength+additionalCapacity)

	if allLength == 0 {
		return &list
	}

	for _, collection := range *it.items {

		for _, s := range *collection.ListPtr() {
			list = append(list, s)
		}
	}

	return &list
}

func (it *CollectionsOfCollection) ToCollection() *Collection {
	list := it.ListPtr(0)

	return NewCollectionUsingStringsPlusCapPtr(list, 0)
}

func (it *CollectionsOfCollection) AddStringsPtr(
	stringsItems *[]string, isCloneAdd bool,
) *CollectionsOfCollection {
	if stringsItems == nil {
		return it
	}

	return it.Adds(NewCollectionUsingStrings(*stringsItems, isCloneAdd))
}

func (it *CollectionsOfCollection) AddPointerStringsPtr(
	pointerStringsItems *[]*string,
) *CollectionsOfCollection {
	if pointerStringsItems == nil {
		return it
	}

	stringsItems := converters.PointerStringsToStrings(pointerStringsItems)

	return it.Adds(NewCollectionUsingStrings(*stringsItems, false))
}

func (it *CollectionsOfCollection) AddsStringsOfStrings(
	isMakeClone bool,
	stringsOfPointerStrings ...*[]string,
) *CollectionsOfCollection {
	if stringsOfPointerStrings == nil {
		return it
	}

	for _, stringsPointer := range stringsOfPointerStrings {
		it.AddStringsPtr(stringsPointer, isMakeClone)
	}

	return it
}

func (it *CollectionsOfCollection) AddsStringsOfPointerStrings(
	isMakeClone bool,
	stringsOfPointerStrings *[]*[]string,
) *CollectionsOfCollection {
	if stringsOfPointerStrings == nil {
		return it
	}

	for _, stringsPointer := range *stringsOfPointerStrings {
		it.AddStringsPtr(stringsPointer, isMakeClone)
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
		it.AddStringsPtr(
			&items,
			isMakeClone)
		it.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return it
}

// AddAsyncFuncItemsPointer must add all the lengths to the wg
func (it *CollectionsOfCollection) AddAsyncFuncItemsPointer(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() *[]string,
) *CollectionsOfCollection {
	if asyncFunctions == nil {
		return it
	}

	asyncFuncWrap := func(asyncFunc func() *[]string) {
		items := asyncFunc()

		if items == nil || len(*items) == 0 {
			wg.Done()

			return
		}

		it.Lock()
		it.AddStringsPtr(
			items,
			isMakeClone)
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
	collections ...*Collection,
) *CollectionsOfCollection {
	if collections == nil {
		return it
	}

	return it.AddCollections(&collections)
}

func (it *CollectionsOfCollection) AddCollections(
	collections *[]*Collection,
) *CollectionsOfCollection {
	if collections == nil {
		return it
	}

	for i := range *collections {
		*it.items = append(*it.items, (*collections)[i])
	}

	return it
}

func (it *CollectionsOfCollection) String() string {
	list := make(
		[]string,
		0,
		it.Length())

	for i, collection := range *it.items {
		list = append(
			list,
			collection.SummaryString(i+1))
	}

	return strings.Join(
		list,
		constants.DoubleNewLine)
}

func (it *CollectionsOfCollection) JsonModel() *CollectionsOfCollectionModel {
	return &CollectionsOfCollectionModel{
		Items: it.items,
	}
}

func (it *CollectionsOfCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *CollectionsOfCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
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
	return corejson.NewFromAny(it)
}

func (it CollectionsOfCollection) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *CollectionsOfCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CollectionsOfCollection, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return EmptyCollectionsOfCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
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
