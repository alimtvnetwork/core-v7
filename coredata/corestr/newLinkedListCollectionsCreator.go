package corestr

import (
	"gitlab.com/auk-go/core/converters"
)

type newLinkedListCollectionsCreator struct{}

func (it *newLinkedListCollectionsCreator) Create() *LinkedCollections {
	return &LinkedCollections{}
}

func (it *newLinkedListCollectionsCreator) Empty() *LinkedCollections {
	return &LinkedCollections{}
}

func (it *newLinkedListCollectionsCreator) PointerStringsPtr(
	stringItems *[]*string,
) *LinkedCollections {
	if stringItems == nil {
		return &LinkedCollections{}
	}

	linkedList := it.Create()
	slice := converters.
		StringsTo.
		PtrOfPtrToPtrStrings(stringItems)

	return linkedList.AddStrings(*slice...)
}

func (it *newLinkedListCollectionsCreator) UsingCollections(
	collections ...*Collection,
) *LinkedCollections {
	if collections == nil {
		return &LinkedCollections{}
	}

	linkedList := it.Create()

	return linkedList.
		AppendCollectionsPointers(
			true,
			&collections,
		)
}

func (it *newLinkedListCollectionsCreator) Strings(
	stringItems ...string,
) *LinkedCollections {
	linkedList := &LinkedCollections{}

	if len(stringItems) == 0 {
		return linkedList
	}

	return linkedList.AddStrings(stringItems...)
}
