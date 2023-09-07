package corestr

import "gitlab.com/auk-go/core/constants"

type newHashsetsCollectionCreator struct{}

func (it *newHashsetsCollectionCreator) Empty() *HashsetsCollection {
	collection := make(
		[]*Hashset,
		constants.Zero,
		constants.Zero)

	return &HashsetsCollection{
		items: collection,
	}
}

func (it *newHashsetsCollectionCreator) UsingHashsets(
	hashsets ...Hashset,
) *HashsetsCollection {
	length := len(hashsets)

	if length == 0 {
		return it.Empty()
	}

	collection := make(
		[]*Hashset,
		length,
		length+constants.ArbitraryCapacity10)

	for i := range hashsets {
		collection[i] = &hashsets[i]
	}

	return &HashsetsCollection{
		items: collection,
	}
}

func (it *newHashsetsCollectionCreator) UsingHashsetsPointers(
	hashsets ...*Hashset,
) *HashsetsCollection {
	if len(hashsets) == 0 {
		return it.Empty()
	}

	return &HashsetsCollection{
		items: hashsets,
	}
}

func (it *newHashsetsCollectionCreator) LenCap(
	length, capacity int,
) *HashsetsCollection {
	collection := make([]*Hashset, length, capacity)

	return &HashsetsCollection{
		items: collection,
	}
}

func (it *newHashsetsCollectionCreator) Cap(
	capacity int,
) *HashsetsCollection {
	collection := make([]*Hashset, constants.Zero, capacity)

	return &HashsetsCollection{
		items: collection,
	}
}
