package corestr

type newLinkedListCreator struct{}

func (it *newLinkedListCreator) Create() *LinkedList {
	return &LinkedList{}
}

func (it *newLinkedListCreator) Empty() *LinkedList {
	return &LinkedList{}
}

func (it *newLinkedListCreator) PointerStringsPtr(
	stringItems *[]*string,
) *LinkedList {
	if stringItems == nil || len(*stringItems) == 0 {
		return &LinkedList{}
	}

	linkedList := it.Empty()

	return linkedList.
		AddPointerStringsPtr(stringItems)
}

func (it *newLinkedListCreator) Strings(
	stringItems []string,
) *LinkedList {
	if len(stringItems) == 0 {
		return &LinkedList{}
	}

	linkedList := it.Create()

	return linkedList.
		Adds(stringItems...)
}

func (it *newLinkedListCreator) SpreadStrings(
	stringItems ...string,
) *LinkedList {
	if len(stringItems) == 0 {
		return &LinkedList{}
	}

	linkedList := it.Create()

	return linkedList.
		Adds(stringItems...)
}

func (it *newLinkedListCreator) UsingMap(
	itemsMap map[string]bool,
) *LinkedList {
	if itemsMap == nil {
		return &LinkedList{}
	}

	linkedList := it.Create()

	return linkedList.
		AddItemsMap(itemsMap)
}
