package corestr

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

type CollectionsOfCollection struct {
	items *[]*Collection
}

func (cc *CollectionsOfCollection) IsEmpty() bool {
	return cc.items == nil || len(*cc.items) == 0
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

func (cc *CollectionsOfCollection) AddStringsPtr(stringsItems *[]string, isCloneAdd bool) *CollectionsOfCollection {
	if stringsItems == nil {
		return cc
	}

	return cc.Adds(NewCollectionUsingStrings(stringsItems, isCloneAdd))
}

func (cc *CollectionsOfCollection) AddPointerStringsPtr(pointerStringsItems *[]*string) *CollectionsOfCollection {
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

func (cc *CollectionsOfCollection) Adds(collections ...*Collection) *CollectionsOfCollection {
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
	list := make([]string, 0, cc.Length())

	for i, collection := range *cc.items {
		list = append(list, collection.SummaryString(i+1))
	}

	return strings.Join(list, constants.DoubleNewLine)
}
