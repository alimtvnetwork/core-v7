package corestr

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
)

type newHashsetCreator struct{}

func (it *newHashsetCreator) Empty() *Hashset {
	return it.Cap(constants.Zero)
}

func (it *newHashsetCreator) Cap(length int) *Hashset {
	hashset := make(map[string]bool, length)

	return &Hashset{
		items:         hashset,
		hasMapUpdated: false,
		length:        length,
		isEmptySet:    true,
	}
}

// StringsOption addCapacity will not work if it is not a clone.
//
//goland:noinspection ALL
func (it *newHashsetCreator) StringsOption(
	addCapacity int,
	isMakeClone bool,
	items ...string,
) *Hashset {
	if items == nil && addCapacity == 0 {
		return it.Empty()
	}

	if items == nil && addCapacity > 0 {
		return it.Cap(addCapacity)
	}

	return it.Strings(
		items,
	)
}

func (it *newHashsetCreator) PointerStrings(
	inputArray []*string,
) *Hashset {
	if len(inputArray) == 0 {
		return it.Empty()
	}

	maps := converters.StringsTo.PtrOfPtrToMapStringBool(&inputArray)

	return it.UsingMap(maps)
}

// PointerStringsPtrOption addCapacity will not work if it is not a clone.
func (it *newHashsetCreator) PointerStringsPtrOption(
	addCapacity int,
	isMakeClone bool,
	inputArray *[]*string,
) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return it.Cap(addCapacity)
	}

	maps := converters.StringsTo.PtrOfPtrToMapStringBool(inputArray)

	return it.UsingMapOption(
		addCapacity,
		isMakeClone,
		maps,
	)
}

// UsingCollection addCapacity will not work if it is not a clone.
func (it *newHashsetCreator) UsingCollection(
	collection *Collection,
) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return it.Empty()
	}

	return it.Strings(
		collection.items,
	)
}

func (it *newHashsetCreator) Strings(
	inputArray []string,
) *Hashset {
	if len(inputArray) == 0 {
		return it.Empty()
	}

	maps := converters.StringsTo.Hashset(inputArray)

	return it.UsingMap(
		maps,
	)
}

func (it *newHashsetCreator) SimpleSlice(
	simpleSlice *SimpleSlice,
) *Hashset {
	if simpleSlice.IsEmpty() {
		return it.Empty()
	}

	maps := converters.StringsTo.Hashset(simpleSlice.Strings())

	return it.UsingMap(
		maps,
	)
}

func (it *newHashsetCreator) StringsSpreadItems(
	inputArray ...string,
) *Hashset {
	if len(inputArray) == 0 {
		return it.Empty()
	}

	maps := converters.StringsTo.Hashset(inputArray)

	return it.UsingMapOption(
		constants.Zero,
		false,
		maps,
	)
}

// UsingMapOption addCapacity will not work if it is not a clone.
func (it *newHashsetCreator) UsingMapOption(
	addCapacity int,
	isMakeClone bool,
	itemsMap map[string]bool,
) *Hashset {
	if len(itemsMap) == 0 {
		return it.Cap(addCapacity)
	}

	length := len(itemsMap)

	if isMakeClone {
		hashset := it.Cap(length + addCapacity)

		return hashset.AddItemsMap(itemsMap)
	}

	return &Hashset{
		items:      itemsMap,
		length:     length,
		isEmptySet: length == constants.Zero,
	}
}

func (it *newHashsetCreator) UsingMap(
	itemsMap map[string]bool,
) *Hashset {
	if len(itemsMap) == 0 {
		return it.Cap(0)
	}

	length := len(itemsMap)
	hashset := it.Cap(length)

	return hashset.AddItemsMap(itemsMap)
}
