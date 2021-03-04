package corestr

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type CollectionsOfCollectionPtr struct {
	items *[]*CollectionPtr
}

func (cc *CollectionsOfCollectionPtr) IsEmpty() bool {
	return cc.items == nil || len(*cc.items) == 0
}

func (cc *CollectionsOfCollectionPtr) HasItems() bool {
	return cc.items != nil && len(*cc.items) > 0
}

func (cc *CollectionsOfCollectionPtr) Length() int {
	if cc.items == nil {
		return 0
	}

	return len(*cc.items)
}

func (cc *CollectionsOfCollectionPtr) AllIndividualItemsLength() int {
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

func (cc *CollectionsOfCollectionPtr) ItemsPtr() *[]*CollectionPtr {
	return cc.items
}

func (cc *CollectionsOfCollectionPtr) Items() []*CollectionPtr {
	return *cc.items
}

func (cc *CollectionsOfCollectionPtr) ListPtr(
	additionalCapacity int,
) *[]string {
	allLength := cc.AllIndividualItemsLength()
	list := make([]string, 0, allLength+additionalCapacity)

	if allLength == 0 {
		return &list
	}

	for _, collection := range *cc.items {
		for _, s := range *collection.ListPtr() {
			list = append(list, *s)
		}
	}

	return &list
}

func (cc *CollectionsOfCollectionPtr) ToCollection() *Collection {
	list := cc.ListPtr(0)

	return NewCollectionUsingStringsPlusCap(list, 0)
}

func (cc *CollectionsOfCollectionPtr) AddStringsPtr(
	stringsItems *[]string,
	addCapacity int,
) *CollectionsOfCollectionPtr {
	if stringsItems == nil {
		return cc
	}

	return cc.Adds(
		NewCollectionPtrUsingStrings(
			stringsItems,
			addCapacity))
}

func (cc *CollectionsOfCollectionPtr) AddPointerStringsPtr(
	pointerStringsItems *[]*string,
) *CollectionsOfCollectionPtr {
	if pointerStringsItems == nil {
		return cc
	}

	return cc.Adds(
		NewCollectionPtrUsingPointerStrings(
			pointerStringsItems,
			0))
}

func (cc *CollectionsOfCollectionPtr) AddsStringsOfStrings(
	addCapacity int,
	stringsOfPointerStrings ...*[]string,
) *CollectionsOfCollectionPtr {
	if stringsOfPointerStrings == nil {
		return cc
	}

	for _, stringsPointer := range stringsOfPointerStrings {
		cc.AddStringsPtr(
			stringsPointer,
			addCapacity)
	}

	return cc
}

func (cc *CollectionsOfCollectionPtr) AddsStringsOfPointerStrings(
	addCapacity int,
	stringsOfPointerStrings *[]*[]string,
) *CollectionsOfCollectionPtr {
	if stringsOfPointerStrings == nil {
		return cc
	}

	for _, stringsPointer := range *stringsOfPointerStrings {
		cc.AddStringsPtr(
			stringsPointer,
			addCapacity)
	}

	return cc
}

func (cc *CollectionsOfCollectionPtr) Adds(
	collections ...*CollectionPtr,
) *CollectionsOfCollectionPtr {
	if collections == nil {
		return cc
	}

	return cc.AddCollections(&collections)
}

func (cc *CollectionsOfCollectionPtr) AddCollections(
	collections *[]*CollectionPtr,
) *CollectionsOfCollectionPtr {
	if collections == nil {
		return cc
	}

	for i := range *collections {
		*cc.items = append(
			*cc.items,
			(*collections)[i])
	}

	return cc
}

func (cc *CollectionsOfCollectionPtr) String() string {
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
