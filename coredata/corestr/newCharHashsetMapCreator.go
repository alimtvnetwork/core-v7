package corestr

import (
	"gitlab.com/auk-go/core/constants"
)

type newCharHashsetMapCreator struct{}

// Cap
//
// CharHashsetMap.eachHashsetCapacity,
// capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the Empty hashset definition.
func (it *newCharHashsetMapCreator) Cap(
	capacity, selfHashsetCapacity int,
) *CharHashsetMap {
	const limit = constants.ArbitraryCapacity10

	if capacity < limit {
		capacity = limit
	}

	mapElements := make(
		map[byte]*Hashset,
		capacity)

	if selfHashsetCapacity < limit {
		selfHashsetCapacity = limit
	}

	return &CharHashsetMap{
		items:               mapElements,
		eachHashsetCapacity: selfHashsetCapacity,
	}
}

func (it *newCharHashsetMapCreator) CapItemsPtr(
	capacity, selfHashsetCapacity int,
	items *[]string,
) *CharHashsetMap {
	charHashsetMap := it.Cap(
		capacity, selfHashsetCapacity)

	return charHashsetMap.AddStringsPtr(items)
}

func (it *newCharHashsetMapCreator) CapItems(
	capacity, selfHashsetCapacity int,
	items []string,
) *CharHashsetMap {
	charHashsetMap := it.Cap(
		capacity, selfHashsetCapacity)

	return charHashsetMap.AddStrings(items...)
}

func (it *newCharHashsetMapCreator) Strings(
	selfHashsetCapacity int,
	items []string,
) *CharHashsetMap {
	if items == nil {
		return it.Cap(
			constants.ArbitraryCapacity5,
			selfHashsetCapacity)
	}

	length := len(items)
	charHashsetMap := it.Cap(
		length,
		selfHashsetCapacity)

	charHashsetMap.AddStrings(items...)

	return charHashsetMap
}

func (it *newCharHashsetMapCreator) StringsPtr(
	selfHashsetCapacity int,
	items *[]string,
) *CharHashsetMap {
	if items == nil {
		return it.Cap(
			constants.ArbitraryCapacity5,
			selfHashsetCapacity)
	}

	length := len(*items)
	charHashsetMap := it.Cap(
		length,
		selfHashsetCapacity)

	return charHashsetMap.AddStringsPtr(items)
}
