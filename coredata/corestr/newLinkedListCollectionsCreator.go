package corestr

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

	return linkedList.
		AddPointerStringsPtr(stringItems)
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
			&collections)
}

func (it *newLinkedListCollectionsCreator) StringsPtr(
	isMakeClone bool,
	stringItems *[]string,
) *LinkedCollections {
	if stringItems == nil {
		return &LinkedCollections{}
	}

	linkedList := it.Create()

	return linkedList.
		AddStringsPtr(isMakeClone, stringItems)
}

func (it *newLinkedListCollectionsCreator) StringsOptions(
	isClone bool,
	stringItems []string,
) *LinkedCollections {
	linkedList := &LinkedCollections{}

	if len(stringItems) == 0 {
		return linkedList
	}

	return linkedList.
		AddStringsPtr(isClone, &stringItems)
}

func (it *newLinkedListCollectionsCreator) StringsPtrOptions(
	isClone bool,
	stringItems *[]string,
) *LinkedCollections {
	linkedList := &LinkedCollections{}

	if stringItems == nil || len(*stringItems) == 0 {
		return linkedList
	}

	return linkedList.
		AddStringsPtr(isClone, stringItems)
}

func (it *newLinkedListCollectionsCreator) Strings(
	stringItems []string,
) *LinkedCollections {
	linkedList := &LinkedCollections{}

	if len(stringItems) == 0 {
		return linkedList
	}

	return linkedList.
		AddStringsPtr(
			false,
			&stringItems)
}
