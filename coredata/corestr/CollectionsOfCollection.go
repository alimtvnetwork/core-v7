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

func (cc *CollectionsOfCollection) IsEmpty() bool {
	return cc.items == nil || len(*cc.items) == 0
}

func (cc *CollectionsOfCollection) HasItems() bool {
	return cc.items != nil && len(*cc.items) > 0
}

func (cc *CollectionsOfCollection) Length() int {
	if cc.items == nil {
		return 0
	}

	return len(*cc.items)
}

func (cc *CollectionsOfCollection) AllIndividualItemsLength() int {
	if cc.IsEmpty() {
		return 0
	}

	allLength := 0

	for _, collection := range *cc.items {
		if collection == nil || collection.IsEmpty() {
			continue
		}

		allLength += collection.Length()
	}

	return allLength
}

func (cc *CollectionsOfCollection) ItemsPtr() *[]*Collection {
	return cc.items
}

func (cc *CollectionsOfCollection) Items() []*Collection {
	return *cc.items
}

func (cc *CollectionsOfCollection) ListPtr(additionalCapacity int) *[]string {
	allLength := cc.AllIndividualItemsLength()
	list := make([]string, 0, allLength+additionalCapacity)

	if allLength == 0 {
		return &list
	}

	for _, collection := range *cc.items {

		for _, s := range *collection.ListPtr() {
			list = append(list, s)
		}
	}

	return &list
}

func (cc *CollectionsOfCollection) ToCollection() *Collection {
	list := cc.ListPtr(0)

	return NewCollectionUsingStringsPlusCap(list, 0)
}

func (cc *CollectionsOfCollection) AddStringsPtr(
	stringsItems *[]string, isCloneAdd bool,
) *CollectionsOfCollection {
	if stringsItems == nil {
		return cc
	}

	return cc.Adds(NewCollectionUsingStrings(stringsItems, isCloneAdd))
}

func (cc *CollectionsOfCollection) AddPointerStringsPtr(
	pointerStringsItems *[]*string,
) *CollectionsOfCollection {
	if pointerStringsItems == nil {
		return cc
	}

	stringsItems := converters.PointerStringsToStrings(pointerStringsItems)

	return cc.Adds(NewCollectionUsingStrings(stringsItems, false))
}

func (cc *CollectionsOfCollection) AddsStringsOfStrings(
	isMakeClone bool,
	stringsOfPointerStrings ...*[]string,
) *CollectionsOfCollection {
	if stringsOfPointerStrings == nil {
		return cc
	}

	for _, stringsPointer := range stringsOfPointerStrings {
		cc.AddStringsPtr(stringsPointer, isMakeClone)
	}

	return cc
}

func (cc *CollectionsOfCollection) AddsStringsOfPointerStrings(
	isMakeClone bool,
	stringsOfPointerStrings *[]*[]string,
) *CollectionsOfCollection {
	if stringsOfPointerStrings == nil {
		return cc
	}

	for _, stringsPointer := range *stringsOfPointerStrings {
		cc.AddStringsPtr(stringsPointer, isMakeClone)
	}

	return cc
}

// AddAsyncFuncItems must add all the lengths to the wg
func (cc *CollectionsOfCollection) AddAsyncFuncItems(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() []string,
) *CollectionsOfCollection {
	if asyncFunctions == nil {
		return cc
	}

	asyncFuncWrap := func(asyncFunc func() []string) {
		items := asyncFunc()

		if len(items) == 0 {
			wg.Done()

			return
		}

		cc.Lock()
		cc.AddStringsPtr(
			&items,
			isMakeClone)
		cc.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return cc
}

// AddAsyncFuncItemsPointer must add all the lengths to the wg
func (cc *CollectionsOfCollection) AddAsyncFuncItemsPointer(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() *[]string,
) *CollectionsOfCollection {
	if asyncFunctions == nil {
		return cc
	}

	asyncFuncWrap := func(asyncFunc func() *[]string) {
		items := asyncFunc()

		if items == nil || len(*items) == 0 {
			wg.Done()

			return
		}

		cc.Lock()
		cc.AddStringsPtr(
			items,
			isMakeClone)
		cc.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return cc
}

func (cc *CollectionsOfCollection) Adds(
	collections ...*Collection,
) *CollectionsOfCollection {
	if collections == nil {
		return cc
	}

	return cc.AddCollections(&collections)
}

func (cc *CollectionsOfCollection) AddCollections(
	collections *[]*Collection,
) *CollectionsOfCollection {
	if collections == nil {
		return cc
	}

	for i := range *collections {
		*cc.items = append(*cc.items, (*collections)[i])
	}

	return cc
}

func (cc *CollectionsOfCollection) String() string {
	list := make(
		[]string,
		0,
		cc.Length())

	for i, collection := range *cc.items {
		list = append(
			list,
			collection.SummaryString(i+1))
	}

	return strings.Join(
		list,
		constants.DoubleNewLine)
}

func (cc *CollectionsOfCollection) JsonModel() *CollectionsOfCollectionModel {
	return &CollectionsOfCollectionModel{
		Items: cc.items,
	}
}

func (cc *CollectionsOfCollection) JsonModelAny() interface{} {
	return cc.JsonModel()
}

func (cc *CollectionsOfCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(*cc.JsonModel())
}

func (cc *CollectionsOfCollection) UnmarshalJSON(data []byte) error {
	var dataModel CollectionsOfCollectionModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		cc.items = dataModel.Items
	}

	return err
}

//goland:noinspection GoLinterLocal
func (cc *CollectionsOfCollection) Json() *corejson.Result {
	if cc.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(cc)

	return corejson.NewPtr(jsonBytes, err)
}

//goland:noinspection GoLinterLocal
func (cc *CollectionsOfCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CollectionsOfCollection, error) {
	err := jsonResult.Unmarshal(&cc)

	if err != nil {
		return EmptyCollectionsOfCollection(), err
	}

	return cc, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (cc *CollectionsOfCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CollectionsOfCollection {
	newUsingJson, err :=
		cc.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (cc *CollectionsOfCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := cc.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (cc *CollectionsOfCollection) AsJsoner() corejson.Jsoner {
	return cc
}

func (cc *CollectionsOfCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return cc
}

func (cc *CollectionsOfCollection) AsJsonMarshaller() corejson.JsonMarshaller {
	return cc
}
