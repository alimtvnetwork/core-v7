package corestr

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type newCollectionCreator struct{}

func (it *newCollectionCreator) Empty() *Collection {
	return &Collection{
		items: []string{},
	}
}

func (it *newCollectionCreator) Cap(capacity int) *Collection {
	collection := make([]string, constants.Zero, capacity)

	return &Collection{
		items: collection,
	}
}

func (it *newCollectionCreator) CloneStrings(stringItems []string) *Collection {
	length := len(stringItems)
	slice := make([]string, 0, length+constants.Capacity4)

	collection := &Collection{
		items: slice,
	}

	return collection.AddStrings(stringItems)
}

func (it *newCollectionCreator) Create(stringItems []string) *Collection {
	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) Strings(stringItems []string) *Collection {
	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) StringsOptions(
	isMakeClone bool,
	stringItems []string,
) *Collection {
	if isMakeClone {
		length := len(stringItems)
		slice := make([]string, 0, length+constants.Capacity4)

		collection := &Collection{
			items: slice,
		}

		return collection.AddStrings(stringItems)
	}

	if len(stringItems) == 0 {
		return it.Empty()
	}

	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) LineUsingSep(sep, line string) *Collection {
	lines := strings.Split(line, sep)

	return &Collection{
		items: lines,
	}
}

func (it *newCollectionCreator) LineDefault(compiledLine string) *Collection {
	lines := strings.Split(compiledLine, constants.DefaultLine)

	return &Collection{
		items: lines,
	}
}

func (it *newCollectionCreator) StringsPlusCap(
	additionalCapacity int,
	stringItems []string,
) *Collection {
	if additionalCapacity == 0 {
		return it.Strings(stringItems)
	}

	length := len(stringItems)
	collection := it.Cap(length + additionalCapacity)

	return collection.Adds(stringItems...)
}

func (it *newCollectionCreator) CapStrings(
	additionalCap int,
	stringItems []string,
) *Collection {
	if additionalCap == 0 {
		return it.StringsOptions(
			false,
			stringItems,
		)
	}

	length := len(stringItems)
	collection := it.Cap(length + additionalCap)

	return collection.AddStrings(stringItems)
}

func (it *newCollectionCreator) LenCap(length, capacity int) *Collection {
	collection := make([]string, length, capacity)

	return &Collection{
		items: collection,
	}
}
