package corestr

import (
	"gitlab.com/evatix-go/core/constants"
)

type newCollectionsOfCollectionCreator struct{}

func (it *newCollectionsOfCollectionCreator) Cap(
	capacity int,
) *CollectionsOfCollection {
	collection := make([]*Collection, constants.Zero, capacity)

	return &CollectionsOfCollection{
		items: collection,
	}
}

func (it *newCollectionsOfCollectionCreator) Empty() *CollectionsOfCollection {
	collection := make([]*Collection, constants.Zero)

	return &CollectionsOfCollection{
		items: collection,
	}
}

func (it *newCollectionsOfCollectionCreator) StringsOfStrings(
	isMakeClone bool,
	stringItems ...*[]string,
) *CollectionsOfCollection {
	length := LengthOfStringsOfPointerStrings(&stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(isMakeClone, &stringItems)
}

func (it *newCollectionsOfCollectionCreator) StringsOfStringsPtrPtr(
	isMakeClone bool,
	stringItems *[]*[]string,
) *CollectionsOfCollection {
	length := LengthOfStringsOfPointerStrings(
		stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(
		isMakeClone,
		stringItems)
}

func (it *newCollectionsOfCollectionCreator) SpreadStrings(
	isMakeClone bool,
	stringItems ...string,
) *CollectionsOfCollection {
	length := len(
		stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddStringsPtr(
		isMakeClone,
		&stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) CloneStrings(
	stringItems []string,
) *CollectionsOfCollection {
	return it.StringsOption(
		true,
		0,
		stringItems)
}

func (it *newCollectionsOfCollectionCreator) Strings(
	stringItems []string,
) *CollectionsOfCollection {
	length := len(
		stringItems)
	collection := it.Cap(
		length)

	return collection.AddStringsPtr(
		false,
		&stringItems)
}

func (it *newCollectionsOfCollectionCreator) StringsOption(
	isMakeClone bool,
	capacity int,
	stringItems []string,
) *CollectionsOfCollection {
	length := len(
		stringItems)
	collection := it.Cap(
		length + capacity)

	return collection.AddStringsPtr(
		isMakeClone,
		&stringItems)
}

func (it *newCollectionsOfCollectionCreator) StringsPtrOption(
	isMakeClone bool,
	capacity int,
	stringItems *[]string,
) *CollectionsOfCollection {
	length := LengthOfStringsPtr(
		stringItems)
	collection := it.Cap(
		length + capacity)

	return collection.AddStringsPtr(
		isMakeClone,
		stringItems)
}

func (it *newCollectionsOfCollectionCreator) PointerStringsPtrOption(
	capacity int,
	stringItems *[]*string,
) *CollectionsOfCollection {
	length := LengthOfPointerStrings(
		stringItems)
	collection := it.Cap(
		length + capacity)

	return collection.AddPointerStringsPtr(
		stringItems)
}

func (it *newCollectionsOfCollectionCreator) LenCap(
	length,
	capacity int,
) *CollectionsOfCollection {
	collection := make(
		[]*Collection,
		length,
		capacity)

	return &CollectionsOfCollection{
		items: collection,
	}
}
