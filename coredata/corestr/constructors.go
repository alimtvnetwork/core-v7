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

func NewHashsetWithValues(items ...string) *Hashset {
	if items == nil {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(&items)
}

func NewUsingStringPointersArray(inputArray *[]*string) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return NewHashset(defaultHashsetItems)
	}

	maps := converters.StringsPointersToStringBoolMap(inputArray)

	return NewHashsetUsingMap(maps)
}

func NewHashsetUsingCollection(collection *Collection) *Hashset {
	if collection == nil || collection.IsEmpty() {
		return EmptyHashset()
	}

	return NewHashsetUsingStrings(collection.items)
}

func NewHashsetUsingStrings(inputArray *[]string) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return EmptyHashset()
	}

	maps := converters.StringsToMap(inputArray)

	return NewHashsetUsingMap(maps)
}

func NewHashsetUsingMap(mapString *map[string]bool) *Hashset {
	if mapString == nil || *mapString == nil {
		return NewHashset(defaultHashsetItems)
	}

	length := len(*mapString)

	return &Hashset{
		items:         mapString,
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
