package corestr

import "gitlab.com/auk-go/core/constants"

type newCharCollectionMapCreator struct{}

// CapSelfCap CharHashsetMap.eachCollectionCapacity,
// capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the Empty items definition.
func (it *newCharCollectionMapCreator) CapSelfCap(
	capacity, selfCollectionCapacity int,
) *CharCollectionMap {
	if capacity < charCollectionDefaultCapacity {
		capacity = charCollectionDefaultCapacity
	}

	mapElements := make(map[byte]*Collection, capacity)

	if selfCollectionCapacity < charCollectionDefaultCapacity {
		selfCollectionCapacity = charCollectionDefaultCapacity
	}

	return &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: selfCollectionCapacity,
	}
}

// Empty eachCollectionCapacity = 0
func (it *newCharCollectionMapCreator) Empty() *CharCollectionMap {
	mapElements := make(map[byte]*Collection, constants.Zero)

	return &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: defaultEachCollectionCapacity,
	}
}

func (it *newCharCollectionMapCreator) Items(
	items []string,
) *CharCollectionMap {
	if len(items) == 0 {
		return it.Empty()
	}

	mapElements := make(map[byte]*Collection, len(items))
	charCollectionMap := &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: constants.Zero,
	}

	charCollectionMap.AddStrings(items...)

	return charCollectionMap
}

func (it *newCharCollectionMapCreator) ItemsPtrWithCap(
	additionalCapacity int,
	eachCollectionCap int,
	items []string,
) *CharCollectionMap {
	length := len(items)
	isDefined := length > 0
	if isDefined {
		additionalCapacity += length
	}

	mapElements := make(
		map[byte]*Collection,
		additionalCapacity,
	)

	charCollectionMap := &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: eachCollectionCap,
	}

	if !isDefined || length == 0 {
		return charCollectionMap
	}

	return charCollectionMap.AddStrings(items...)
}
