package corestr

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

// --------- Hashset starts ----------

func EmptyHashset() *Hashset {
	return NewHashset(0)
}

func NewHashset(length int) *Hashset {
	hashset := make(map[string]bool, length)

	return &Hashset{
		items:         &hashset,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    true,
		Mutex:         sync.Mutex{},
	}
}

// addCapacity will not work if it is not a clone.
func NewHashsetWithValues(
	addCapacity int,
	isMakeClone bool,
	items ...string,
) *Hashset {
	if items == nil {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(
		&items,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingStringPointersArray(
	inputArray *[]*string,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return NewHashset(defaultHashsetItems)
	}

	maps := converters.StringsPointersToStringBoolMap(inputArray)

	return NewHashsetUsingMap(
		maps,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingCollection(
	collection *Collection,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(
		collection.items,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingStrings(
	inputArray *[]string,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return EmptyHashset()
	}

	maps := converters.StringsToMap(inputArray)

	return NewHashsetUsingMap(
		maps,
		addCapacity,
		isMakeClone)
}

// addCapacity will not work if it is not a clone.
func NewHashsetUsingMap(
	itemsMap *map[string]bool,
	addCapacity int,
	isMakeClone bool,
) *Hashset {
	if itemsMap == nil || *itemsMap == nil {
		return NewHashset(defaultHashsetItems)
	}

	length := len(*itemsMap)
	if isMakeClone {
		hashset := NewHashset(length + addCapacity)

		hashset.AddItemsMap(itemsMap)

		return hashset
	}

	return &Hashset{
		items:         itemsMap,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
		Mutex:         sync.Mutex{},
	}
}

// --------- Collection starts ----------

func NewCollection(capacity int) *Collection {
	collection := make([]string, 0, capacity)

	return &Collection{
		items: &collection,
	}
}

func EmptyCollection() *Collection {
	collection := make([]string, 0, 0)

	return &Collection{
		items: &collection,
	}
}

func NewCollectionUsingStrings(stringItems *[]string) *Collection {
	return &Collection{
		items: stringItems,
	}
}

func NewCollectionUsingLength(len, capacity int) *Collection {
	collection := make([]string, len, capacity)

	return &Collection{
		items: &collection,
	}
}

// --------- HashsetsCollection starts ----------

func EmptyHashsetsCollection() *HashsetsCollection {
	collection := make([]*Hashset, 0, 0)

	return &HashsetsCollection{
		items: &collection,
	}
}

func NewHashsetsCollection(hashsets *[]Hashset) *HashsetsCollection {
	if hashsets == nil ||
		*hashsets == nil {
		return EmptyHashsetsCollection()
	}

	length := len(*hashsets)
	collection := make(
		[]*Hashset,
		length,
		length+constants.ArbitraryCapacity10)

	for i, hashset := range *hashsets {
		collection[i] = &hashset
	}

	return &HashsetsCollection{
		items: &collection,
	}
}

func NewHashsetsCollectionUsingPointerHashsets(hashsets *[]*Hashset) *HashsetsCollection {
	if hashsets == nil ||
		*hashsets == nil {
		return EmptyHashsetsCollection()
	}

	return &HashsetsCollection{
		items: hashsets,
	}
}

func NewHashsetsCollectionUsingLength(len, capacity int) *HashsetsCollection {
	collection := make([]*Hashset, len, capacity)

	return &HashsetsCollection{
		items: &collection,
	}
}

// --------- Hashmap starts ----------

func EmptyHashmap() *Hashmap {
	return NewHashmap(0)
}

func NewHashmap(length int) *Hashmap {
	hashset := make(map[string]string, length)

	return &Hashmap{
		items:         &hashset,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    true,
		Mutex:         sync.Mutex{},
	}
}

func NewHashmapUsingKeyAnyValues(keyAnyValues *[]KeyAnyValuePair) *Hashmap {
	if keyAnyValues == nil || *keyAnyValues == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*keyAnyValues)
	hashMap := NewHashmap(length + constants.ArbitraryCapacity10)
	hashMap.AddOrUpdateKeyAnyValsPtr(keyAnyValues)

	return hashMap
}

func NewHashmapUsingKeyValues(keyValues *[]KeyValuePair) *Hashmap {
	if keyValues == nil || *keyValues == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*keyValues)
	hashMap := NewHashmap(length + constants.ArbitraryCapacity10)
	hashMap.AddOrUpdateKeyValsPtr(keyValues)

	return hashMap
}

func NewHashmapUsingCollection(keys, values *Collection) *Hashmap {
	if keys == nil || keys.IsEmpty() {
		return EmptyHashmap()
	}

	itemsMap := converters.KeysValuesStringsToMap(
		keys.items,
		values.items)

	return NewHashmapUsingMap(
		itemsMap,
		0,
		false)
}

func NewHashmapUsingStrings(keys, values *[]string) *Hashmap {
	if keys == nil || *keys == nil {
		return EmptyHashmap()
	}

	itemsMap := converters.KeysValuesStringsToMap(
		keys,
		values)

	return NewHashmapUsingMap(
		itemsMap,
		0,
		false)
}

// isMakeClone : copies itemsMap or else use the same one as pointer assign.
func NewHashmapUsingMap(
	itemsMap *map[string]string,
	addCapacity int,
	isMakeClone bool,
) *Hashmap {
	if itemsMap == nil || *itemsMap == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*itemsMap)

	if isMakeClone {
		hashMap := NewHashmap(length + addCapacity)
		hashMap.AddOrUpdateMap(itemsMap)

		return hashMap
	}

	return &Hashmap{
		items:         itemsMap,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
	}
}

// always returns the clone of the items.
func NewHashmapUsingMapUsingAddCapacity(
	itemsMap *map[string]string,
	addCapacity int,
) *Hashmap {
	if itemsMap == nil || *itemsMap == nil {
		return NewHashmap(defaultHashsetItems)
	}

	length := len(*itemsMap)
	hashMap := NewHashmap(length + addCapacity)
	hashMap.AddOrUpdateMap(itemsMap)

	return hashMap
}
