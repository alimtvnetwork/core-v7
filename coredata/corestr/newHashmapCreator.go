package corestr

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
)

type newHashmapCreator struct{}

func (it *newHashmapCreator) Empty() *Hashmap {
	return it.Cap(constants.Zero)
}

func (it *newHashmapCreator) Cap(length int) *Hashmap {
	hashset := make(map[string]string, length)

	return &Hashmap{
		items:         hashset,
		hasMapUpdated: false,
		length:        length,
		isEmptySet:    true,
	}
}

func (it *newHashmapCreator) KeyAnyValuesPtr(
	keyAnyValues *[]KeyAnyValuePair,
) *Hashmap {
	if keyAnyValues == nil || *keyAnyValues == nil {
		return it.Cap(defaultHashsetItems)
	}

	length := len(*keyAnyValues)
	hashMap := it.Cap(length + constants.ArbitraryCapacity10)

	return hashMap.AddOrUpdateKeyAnyValsPtr(keyAnyValues)
}

func (it *newHashmapCreator) KeyValuesPtr(
	keyValues *[]KeyValuePair,
) *Hashmap {
	if keyValues == nil || *keyValues == nil {
		return it.Cap(defaultHashsetItems)
	}

	length := len(*keyValues)
	hashMap := it.Cap(length + constants.ArbitraryCapacity10)

	return hashMap.AddOrUpdateKeyValsPtr(keyValues)
}

func (it *newHashmapCreator) KeyValuesCollection(
	keys, values *Collection,
) *Hashmap {
	if keys == nil || keys.IsEmpty() {
		return it.Empty()
	}

	itemsMap := converters.KeysValuesStringsToMapPtr(
		keys.ListPtr(),
		values.ListPtr())

	return it.UsingMap(
		*itemsMap)
}

func (it *newHashmapCreator) KeyValuesStrings(
	keys, values []string,
) *Hashmap {
	if len(keys) == 0 {
		return it.Empty()
	}

	itemsMap := converters.KeysValuesStringsToMapPtr(
		&keys,
		&values)

	return it.UsingMap(
		*itemsMap)
}

func (it *newHashmapCreator) UsingMap(
	itemsMap map[string]string,
) *Hashmap {
	length := len(itemsMap)

	return &Hashmap{
		items:      itemsMap,
		length:     length,
		isEmptySet: length == constants.Zero,
	}
}

// UsingMapOptions
// isMakeClone : copies itemsMap or else use the same one as pointer assign.
func (it *newHashmapCreator) UsingMapOptions(
	isMakeClone bool,
	addCapacity int,
	itemsMap map[string]string,
) *Hashmap {
	if len(itemsMap) == 0 {
		return it.Cap(addCapacity)
	}

	length := len(itemsMap)

	if isMakeClone {
		hashMap := it.Cap(length + addCapacity)

		return hashMap.AddOrUpdateMap(itemsMap)
	}

	// no clone
	return &Hashmap{
		items:      itemsMap,
		length:     length,
		isEmptySet: length == constants.Zero,
	}
}

// MapWithCap always returns the clone of the items.
func (it *newHashmapCreator) MapWithCap(
	addCapacity int,
	itemsMap map[string]string,
) *Hashmap {
	if len(itemsMap) == 0 {
		return it.Cap(addCapacity)
	}

	if addCapacity == 0 {
		return it.UsingMap(itemsMap)
	}

	length := len(itemsMap)
	hashMap := it.Cap(length + addCapacity)

	return hashMap.AddOrUpdateMap(itemsMap)
}
