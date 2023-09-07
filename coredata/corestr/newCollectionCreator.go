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

	return collection.AddStringsPtr(&stringItems)
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

func (it *newCollectionCreator) StringsPtr(stringItems *[]string) *Collection {
	if stringItems == nil {
		return it.Empty()
	}

	return &Collection{
		items: *stringItems,
	}
}

func (it *newCollectionCreator) StringsOptions(isMakeClone bool, stringItems []string) *Collection {
	if isMakeClone {
		return it.CloneStrings(stringItems)
	}

	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) StringsPtrOption(isMakeClone bool, stringItems *[]string) *Collection {
	if isMakeClone {
		length := LengthOfStringsPtr(stringItems)
		slice := make([]string, 0, length+constants.Capacity4)

		collection := &Collection{
			items: slice,
		}

		return collection.AddStringsPtr(stringItems)
	}

	if stringItems == nil {
		return it.Empty()
	}

	return &Collection{
		items: *stringItems,
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

func (it *newCollectionCreator) StringsPtrPlusCap(
	additionalCap int,
	stringItems *[]string,
) *Collection {
	if additionalCap == 0 {
		return it.StringsPtrOption(
			false,
			stringItems)
	}

	length := LengthOfStringsPtr(stringItems)
	collection := it.Cap(length + additionalCap)

	return collection.AddStringsPtr(stringItems)
}

func (it *newCollectionCreator) PointerStrings(
	stringItems []*string,
) *Collection {
	if len(stringItems) == 0 {
		return it.Empty()
	}

	length := LengthOfPointerStrings(&stringItems)
	collection := it.Cap(length)

	return collection.AddPointerStringsPtr(&stringItems)
}

func (it *newCollectionCreator) PointerStringsPtrUsingCap(
	capacity int,
	stringItems *[]*string,
) *Collection {
	length := LengthOfPointerStrings(stringItems)
	collection := it.Cap(length + capacity)

	return collection.AddPointerStringsPtr(stringItems)
}

func (it *newCollectionCreator) LenCap(length, capacity int) *Collection {
	collection := make([]string, length, capacity)

	return &Collection{
		items: collection,
	}
}
