package corestr

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
)

type newCollectionPtrCreator struct{}

func (it *newCollectionPtrCreator) Empty() *CollectionPtr {
	return &CollectionPtr{
		items: []*string{},
	}
}

func (it *newCollectionPtrCreator) Cap(capacity int) *CollectionPtr {
	collection := make([]*string, constants.Zero, capacity)

	return &CollectionPtr{
		items: collection,
	}
}

func (it *newCollectionPtrCreator) StringsOptions(
	capacity int,
	stringItems []string,
) *CollectionPtr {
	length := len(stringItems)
	slice := make([]*string, 0, length+capacity)

	collection := &CollectionPtr{
		items: slice,
	}

	return collection.AddStringsPtr(&stringItems)
}

func (it *newCollectionPtrCreator) Strings(
	stringItems []string,
) *CollectionPtr {
	return &CollectionPtr{
		items: *converters.StringsTo.PointerStrings(&stringItems),
	}
}

func (it *newCollectionPtrCreator) Default(
	stringItems []*string,
) *CollectionPtr {
	return &CollectionPtr{
		items: stringItems,
	}
}

func (it *newCollectionPtrCreator) StringsPtr(
	stringItems *[]string,
) *CollectionPtr {
	if stringItems == nil {
		return it.Empty()
	}

	collection := &CollectionPtr{
		items: *converters.StringsTo.PointerStrings(stringItems),
	}

	return collection.AddStringsPtr(stringItems)
}

func (it *newCollectionPtrCreator) LineUsingSep(sep, line string) *CollectionPtr {
	lines := strings.Split(line, sep)

	return &CollectionPtr{
		items: *converters.StringsTo.PointerStrings(&lines),
	}
}

func (it *newCollectionPtrCreator) LineDefault(compiledLine string) *CollectionPtr {
	lines := strings.Split(
		compiledLine,
		constants.DefaultLine)

	return &CollectionPtr{
		items: *converters.StringsTo.PointerStrings(&lines),
	}
}

func (it *newCollectionPtrCreator) StringsPlusCap(
	additionalCapacity int,
	stringItems []string,
) *CollectionPtr {
	if additionalCapacity == 0 {
		return it.Strings(stringItems)
	}

	length := len(stringItems)
	collection := it.Cap(length + additionalCapacity)

	return collection.Adds(stringItems...)
}

func (it *newCollectionPtrCreator) StringsPtrPlusCap(
	additionalCap int,
	stringItems *[]string,
) *CollectionPtr {
	if additionalCap == 0 {
		return it.StringsPtr(
			stringItems)
	}

	length := LengthOfStringsPtr(stringItems)
	collection := it.Cap(length + additionalCap)

	return collection.AddStringsPtr(stringItems)
}

func (it *newCollectionPtrCreator) PointerStrings(
	stringItems []*string,
) *CollectionPtr {
	length := len(stringItems)
	collection := it.Cap(length)

	return collection.AddPointerStrings(stringItems...)
}

func (it *newCollectionPtrCreator) PointerStringsPtrUsingCap(
	capacity int,
	stringItems *[]*string,
) *CollectionPtr {
	length := LengthOfPointerStrings(stringItems)
	collection := it.Cap(length + capacity)

	return collection.AddPointerStrings(*stringItems...)
}

func (it *newCollectionPtrCreator) LenCap(length, capacity int) *CollectionPtr {
	collection := make([]*string, length, capacity)

	return &CollectionPtr{
		items: collection,
	}
}
