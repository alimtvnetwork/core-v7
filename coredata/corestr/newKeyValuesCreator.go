package corestr

import "gitlab.com/auk-go/core/constants"

type newKeyValuesCreator struct{}

func (it *newKeyValuesCreator) Cap(
	capacity int,
) *KeyValueCollection {
	collection := make(
		[]*KeyValuePair,
		constants.Zero,
		capacity)

	return &KeyValueCollection{
		KeyValuePairs: collection,
	}
}

func (it *newKeyValuesCreator) Empty() *KeyValueCollection {
	collection := make(
		[]*KeyValuePair,
		constants.Zero,
		constants.Zero)

	return &KeyValueCollection{
		KeyValuePairs: collection,
	}
}

func (it *newKeyValuesCreator) UsingMap(
	input map[string]string,
) *KeyValueCollection {
	length := len(input)
	keyValCollection := it.Cap(length)

	if length == 0 {
		return keyValCollection
	}

	return keyValCollection.AddMap(input)
}

func (it *newKeyValuesCreator) UsingKeyValuePairs(
	keyValPairs ...KeyValuePair,
) *KeyValueCollection {
	length := len(keyValPairs)
	keyValCollection := it.Cap(length)

	if length == 0 {
		return keyValCollection
	}

	return keyValCollection.Adds(keyValPairs...)
}

func (it *newKeyValuesCreator) UsingKeyValueStrings(
	keys, values []string,
) *KeyValueCollection {
	length := len(keys)
	keyValCollection := it.Cap(length)

	if length == 0 {
		return keyValCollection
	}

	for i, key := range keys {
		keyValCollection.Add(key, values[i])
	}

	return keyValCollection
}
