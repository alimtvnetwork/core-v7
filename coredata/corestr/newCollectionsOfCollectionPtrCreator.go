package corestr

import "gitlab.com/auk-go/core/constants"

type newCollectionsOfCollectionPtrCreator struct{}

func (it *newCollectionsOfCollectionPtrCreator) Cap(
	capacity int,
) *CollectionsOfCollectionPtr {
	collection := make(
		[]*CollectionPtr,
		constants.Zero,
		capacity)

	return &CollectionsOfCollectionPtr{
		items: collection,
	}
}

func (it *newCollectionsOfCollectionPtrCreator) Empty() *CollectionsOfCollectionPtr {
	collection := make([]*CollectionPtr, constants.Zero)

	return &CollectionsOfCollectionPtr{
		items: collection,
	}
}

func (it *newCollectionsOfCollectionPtrCreator) StringsOfStringsPointer(
	stringItems ...*[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStringsOfPointerStrings(
		&stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(
		constants.Zero,
		&stringItems)
}

func (it *newCollectionsOfCollectionPtrCreator) StringsOfPointerStrings(
	stringItems *[]*[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStringsOfPointerStrings(
		stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddsStringsOfPointerStrings(
		constants.Zero,
		stringItems)
}

func (it *newCollectionsOfCollectionPtrCreator) Strings(
	stringItems []string,
) *CollectionsOfCollectionPtr {
	length := len(stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddStringsPtr(constants.Zero, &stringItems)
}

func (it *newCollectionsOfCollectionPtrCreator) StringsPtr(
	stringItems *[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStringsPtr(stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddStringsPtr(constants.Zero, stringItems)
}

func (it *newCollectionsOfCollectionPtrCreator) CapStrings(
	addCapacity int,
	stringItems ...string,
) *CollectionsOfCollectionPtr {
	length := len(
		stringItems)

	if length == 0 && addCapacity == 0 {
		return it.Empty()
	}

	if length == 0 && addCapacity > 0 {
		return it.Cap(length)
	}

	collection := it.Cap(
		length + addCapacity)

	return collection.AddStringsPtr(
		constants.Zero,
		&stringItems,
	)
}

func (it *newCollectionsOfCollectionPtrCreator) CapStringsPtr(
	addCapacity int,
	stringItems *[]string,
) *CollectionsOfCollectionPtr {
	length := LengthOfStringsPtr(
		stringItems)

	if length == 0 && addCapacity == 0 {
		return it.Empty()
	}

	if length == 0 && addCapacity > 0 {
		return it.Cap(length)
	}

	collection := it.Cap(
		length + addCapacity)

	return collection.AddStringsPtr(
		constants.Zero,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionPtrCreator) PointerStrings(
	stringItems []*string,
) *CollectionsOfCollectionPtr {
	return it.CapPointerStrings(
		0,
		stringItems...)
}

func (it *newCollectionsOfCollectionPtrCreator) CapPointerStrings(
	capacity int,
	stringItems ...*string,
) *CollectionsOfCollectionPtr {
	length := len(stringItems)

	if length == 0 && capacity == 0 {
		return it.Empty()
	}

	if length == 0 && capacity > 0 {
		return it.Cap(length)
	}

	collection := it.Cap(
		length + capacity)

	return collection.AddPointerStrings(
		stringItems...)
}

func (it *newCollectionsOfCollectionPtrCreator) LenCap(
	length,
	capacity int,
) *CollectionsOfCollectionPtr {
	collection := make(
		[]*CollectionPtr,
		length,
		capacity)

	return &CollectionsOfCollectionPtr{
		items: collection,
	}
}
