package corestr

import (
	"github.com/alimtvnetwork/core/constants"
)

type newCollectionsOfCollectionCreator struct{}

func (it *newCollectionsOfCollectionCreator) Cap(
	capacity int,
) *CollectionsOfCollection {
	collection := make([]*Collection, 0, capacity)

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
	stringItems ...[]string,
) *CollectionsOfCollection {
	length := len(stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddsStringsOfStrings(isMakeClone, stringItems...)
}

func (it *newCollectionsOfCollectionCreator) SpreadStrings(
	isMakeClone bool,
	stringItems ...string,
) *CollectionsOfCollection {
	length := len(
		stringItems,
	)

	return it.LenCap(
		constants.Zero,
		length,
	).AddStrings(
		isMakeClone,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) CloneStrings(
	stringItems []string,
) *CollectionsOfCollection {
	return it.StringsOption(
		true,
		0,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) Strings(
	stringItems []string,
) *CollectionsOfCollection {
	length := len(
		stringItems,
	)
	collection := it.Cap(
		length,
	)

	return collection.AddStrings(
		false,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) StringsOption(
	isMakeClone bool,
	capacity int,
	stringItems []string,
) *CollectionsOfCollection {
	length := len(stringItems)
	collection := it.Cap(
		length + capacity,
	)

	return collection.AddStrings(
		isMakeClone,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) StringsOptions(
	isMakeClone bool,
	capacity int,
	stringItems []string,
) *CollectionsOfCollection {
	length := len(stringItems)
	collection := it.Cap(length + capacity)

	return collection.AddStrings(
		isMakeClone,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) LenCap(
	length,
	capacity int,
) *CollectionsOfCollection {
	collection := make(
		[]*Collection,
		length,
		capacity,
	)

	return &CollectionsOfCollection{
		items: collection,
	}
}
