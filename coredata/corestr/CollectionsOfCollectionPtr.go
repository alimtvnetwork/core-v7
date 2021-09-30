package corestr

import (
	"encoding/json"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type CollectionsOfCollectionPtr struct {
	items *[]*CollectionPtr
}

func (it *CollectionsOfCollectionPtr) IsEmpty() bool {
	return it.items == nil || len(*it.items) == 0
}

func (it *CollectionsOfCollectionPtr) HasItems() bool {
	return it.items != nil && len(*it.items) > 0
}

func (it *CollectionsOfCollectionPtr) Length() int {
	if it.items == nil {
		return 0
	}

	return len(*it.items)
}

func (it *CollectionsOfCollectionPtr) AllIndividualItemsLength() int {
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

func (it *CollectionsOfCollectionPtr) ItemsPtr() *[]*CollectionPtr {
	return it.items
}

func (it *CollectionsOfCollectionPtr) Items() []*CollectionPtr {
	return *it.items
}

func (it *CollectionsOfCollectionPtr) ListPtr(
	additionalCapacity int,
) *[]string {
	allLength := it.AllIndividualItemsLength()
	list := make([]string, 0, allLength+additionalCapacity)

	if allLength == 0 {
		return &list
	}

	for _, collection := range *it.items {
		for _, s := range *collection.ListPtr() {
			list = append(list, *s)
		}
	}

	return &list
}

func (it *CollectionsOfCollectionPtr) ToCollection() *Collection {
	list := it.ListPtr(0)

	return NewCollectionUsingStringsPlusCapPtr(list, 0)
}

func (it *CollectionsOfCollectionPtr) AddStringsPtr(
	stringsItems *[]string,
	addCapacity int,
) *CollectionsOfCollectionPtr {
	if stringsItems == nil {
		return it
	}

	return it.Adds(
		NewCollectionPtrUsingStrings(
			stringsItems,
			addCapacity))
}

func (it *CollectionsOfCollectionPtr) AddPointerStringsPtr(
	pointerStringsItems *[]*string,
) *CollectionsOfCollectionPtr {
	if pointerStringsItems == nil {
		return it
	}

	return it.Adds(
		NewCollectionPtrUsingPointerStrings(
			pointerStringsItems,
			0))
}

func (it *CollectionsOfCollectionPtr) AddsStringsOfStrings(
	addCapacity int,
	stringsOfPointerStrings ...*[]string,
) *CollectionsOfCollectionPtr {
	if stringsOfPointerStrings == nil {
		return it
	}

	for _, stringsPointer := range stringsOfPointerStrings {
		it.AddStringsPtr(
			stringsPointer,
			addCapacity)
	}

	return it
}

func (it *CollectionsOfCollectionPtr) AddsStringsOfPointerStrings(
	addCapacity int,
	stringsOfPointerStrings *[]*[]string,
) *CollectionsOfCollectionPtr {
	if stringsOfPointerStrings == nil {
		return it
	}

	for _, stringsPointer := range *stringsOfPointerStrings {
		it.AddStringsPtr(
			stringsPointer,
			addCapacity)
	}

	return it
}

func (it *CollectionsOfCollectionPtr) Adds(
	collections ...*CollectionPtr,
) *CollectionsOfCollectionPtr {
	if collections == nil {
		return it
	}

	return it.AddCollections(&collections)
}

func (it *CollectionsOfCollectionPtr) AddCollections(
	collections *[]*CollectionPtr,
) *CollectionsOfCollectionPtr {
	if collections == nil {
		return it
	}

	for i := range *collections {
		*it.items = append(
			*it.items,
			(*collections)[i])
	}

	return it
}

func (it *CollectionsOfCollectionPtr) String() string {
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

func (it *CollectionsOfCollectionPtr) JsonModel() *CollectionsOfCollectionPtrModel {
	return &CollectionsOfCollectionPtrModel{
		Items: it.items,
	}
}

func (it *CollectionsOfCollectionPtr) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *CollectionsOfCollectionPtr) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *CollectionsOfCollectionPtr) UnmarshalJSON(data []byte) error {
	var dataModel CollectionsOfCollectionPtrModel

	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.items = dataModel.Items
	}

	return err
}

func (it CollectionsOfCollectionPtr) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it CollectionsOfCollectionPtr) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *CollectionsOfCollectionPtr) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*CollectionsOfCollectionPtr, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return EmptyCollectionsOfCollectionPtr(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *CollectionsOfCollectionPtr) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *CollectionsOfCollectionPtr {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *CollectionsOfCollectionPtr) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *CollectionsOfCollectionPtr) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CollectionsOfCollectionPtr) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *CollectionsOfCollectionPtr) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
